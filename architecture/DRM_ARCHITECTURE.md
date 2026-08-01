# iPod Classic 7G — DRM & Content Protection Architecture

## Overview

The iPod Classic implements multiple layers of content protection. This document describes the observed architecture from firmware analysis — what structures exist, where they are located, and what they protect.

Analysis performed using Ghidra and Capstone disassembly.

---

## Protection Layers

| Layer | Mechanism | Purpose |
|-------|-----------|---------|
| Firmware Encryption | AES-128-CBC (GID key) | Protects OSOS binary on storage |
| FairPlay DRM | RSA + AES key hierarchy | Protects purchased music/video |
| Game Code Signing | PKCS#7 (RSA-SHA1) | Authenticates game executables |
| Game Manifest Signing | XML Digital Signatures | Validates game asset integrity |
| iTunesDB Signing | PKCS#7 | Authenticates music database |

---

## FairPlay Key Hierarchy

```
┌─────────────────────┐
│  GID Key (silicon)  │  Burnt into S5L8702 fuses — never leaves chip
│  AES-128 master key │  Used by hardware AES engine at 0x3D000000
└─────────┬───────────┘
          │
          │  Used to decrypt firmware and in FairPlay key operations
          ▼
┌─────────────────────┐
│  Device Key         │  Derived from GID + device-specific data
│  (per-device)       │  Stored in NOR SysCfg area
└─────────┬───────────┘
          │
          │  Used to decrypt per-content keys from SC_Info
          ▼
┌─────────────────────┐
│  Content Key        │  Per-track AES key stored in SC_Info files
│  (per-track)        │  Encrypted with device key
└─────────┬───────────┘
          │
          │  Decrypts actual audio/video content blocks
          ▼
┌─────────────────────┐
│  Decrypted Content  │  Playable audio/video data
│  (in DRAM)          │  Fed to MeCCA codec pipeline
└─────────────────────┘
```

---

## AES Hardware Engine

| Parameter | Value |
|-----------|-------|
| Base Address | 0x3D000000 |
| Register Range | 0x3D000000–0x3DFFFFFF |
| Functions Accessing | 17 |
| Total Register Accesses | 866 |
| Algorithm | AES-128-CBC |
| Key Sources | GID (silicon fuses) or software-provided |

### AES Usage Contexts

1. **Boot decryption** — NOR bootloader decrypts OSOS from disk
2. **FairPlay content** — OSOS decrypts protected audio/video blocks
3. **Key unwrapping** — Device key operations for content authorization

---

## FairPlay Functions in Firmware

| Function | Address | Size | Purpose |
|----------|---------|------|---------|
| DRM Header/Key Handler | 0x000E6D64 | 360 B | Processes "HeaderKey" and "EncryptedBlocks" |
| DRM Capability/Version | 0x00150AB0 | 6,028 B | DRM version and feature query |

### FairPlay String References

| Offset | String | Context |
|--------|--------|---------|
| — | `HeaderKey` | Key embedded in DRM-protected file header |
| — | `EncryptedBlocks` | Encrypted content segments |
| — | `AppleDRMVersion` | DRM protocol version check |
| — | `AppleDRM` | General DRM identifier |
| — | `AppleVideoDRM` | Video-specific DRM variant |
| — | `KeyTypeSupportVersion` | Key type negotiation |

### DRM Content Structure (Protected M4A/M4V)

Protected files contain standard MP4 atoms plus DRM-specific atoms:
- `drm ` — DRM metadata
- `idrm` — iTunes DRM info
- `sinf` — Sample encryption info
- Encrypted sample data blocks
- Content key encrypted with device key in SC_Info

---

## Game DRM — Code Signing

### Verification Flow

```
Game Selected → Game Loader → Mount Filesystem → Read Manifest.p7b → PKCS#7 Verify → Execute
```

### Game Signing Functions

| Function | Address | Size | Purpose |
|----------|---------|------|---------|
| Game Filesystem Mount | 0x00098600 | 72 B | Mounts games_RO, gamedata_RW, gamestats_WO |
| PKCS#7 Signature Verify | 0x0009829C | 824 B | Reads `.p7b`, verifies RSA-SHA1 |
| X.509 Certificate Parser | 0x0006FE08 | 792 B | Parses certificate chain |
| XML DSig Verifier | 0x00273350 | 2,040 B | Validates XML digital signatures |
| Platform ID Check | 0x00150AB0 | 6,028 B | Checks GamesPlatformID/Version |

### Game Mount Points (from firmware strings)

| Offset | String | Purpose |
|--------|--------|---------|
| 0x000987D8 | `games_RO` | Read-only game assets (RSRC partition) |
| 0x000987B4 | `gamestats_WO` | Write-only game statistics |
| 0x000987C4 | `gamedata_ShareRW` | Shared read-write game data |

### Game Verification Sequence

1. Game selected from TCGamesMenu
2. Game Loader mounts filesystem paths
3. Reads `Manifest.plist.p7b` — PKCS#7 signed manifest
4. Verifies signature against Apple certificate chain
5. Checks `GamesPlatformID` and `GamesPlatformVersion` match device
6. If verification passes → loads and executes ARM game binary
7. If verification fails → displays error screen

### Game Error Screens

| Screen | String | Displayed When |
|--------|--------|----------------|
| Game_Signing_Error | `Game_Signing_Error_Screen` | Signature verification fails |
| Game_Version_Error | `Game_Version_Error_Screen` | Platform version mismatch |
| Game_Memory_Error | `Game_Memory_Error_Screen` | Insufficient memory |

---

## XML Digital Signatures

The firmware implements W3C XML Digital Signature (XMLDSig) verification:

### Supported Algorithms

| Algorithm URI | Purpose |
|--------------|---------|
| `xmldsig#rsa-sha1` | RSA signature with SHA-1 hash |
| `xmldsig#enveloped-signature` | Signature within signed document |
| `xmldsig#sha1` | SHA-1 digest method |

### XMLDSig Verifier (0x00273350)

| Property | Value |
|----------|-------|
| Address | 0x00273350 |
| Size | 2,040 B |
| Called by | FUN_00273C60 (112 B wrapper) |
| Purpose | Validates XML documents with `<Signature>` elements |

---

## OpenSSL in Firmware

The firmware contains a statically-linked OpenSSL library:

| Function | Address | Size | Evidence |
|----------|---------|------|----------|
| OpenSSL Signature Check | 0x00063860 | 520 B | "signature has problems" string |
| SSL Hash Init | 0x0005FC30 | 116 B | "ssl2-md5", "ssl3-sha1" strings |
| OpenSSL Config | 0x002D97E8 | 68 B | "../../apps/openssl.cnf" path |

### Cryptographic Algorithms Present

| Algorithm | Evidence |
|-----------|----------|
| RSA-SHA1 | Used in PKCS#7 and XML DSig |
| SHA-1 | Digest for manifests and signatures |
| MD5 | SSL2 compatibility (legacy) |
| AES-128-CBC | Hardware-accelerated (0x3D000000) |
| X.509 | Certificate chain parsing ("NO X509_NAME") |

---

## What IS Authenticated

| Content | Mechanism | Verification Function |
|---------|-----------|----------------------|
| Game executables (.p7b) | PKCS#7/CMS (RSA-SHA1) | 0x0009829C |
| Game manifests | XML Digital Signature | 0x00273350 |
| iTunesDB | PKCS#7 (iTunesDB.p7b) | 0x0009829C |
| FairPlay content keys | RSA encryption | 0x000E6D64 |
| Firmware (OSOS) | AES-128-CBC encryption | NOR bootloader |

## What Is NOT Authenticated

| Content | Observation |
|---------|-------------|
| OSOS after decryption | No hash verification between decrypt and execute |
| Individual OSOS modules | No per-section signing (monolithic binary) |
| NOR flash contents | No measured boot / secure boot fuse observed |
| RSRC partition | Plaintext, no integrity check on filesystem |
| Fonts (TTF) in RSRC | Replaceable without signature |
| Training data (XML) | Modifiable without signature |
| Game audio/textures | Covered by manifest SHA1 only |

---

## Firmware Update Validation

| Offset | String | Purpose |
|--------|--------|---------|
| — | `WillFlash` | Firmware update in progress flag |
| — | `AutoRebootAfterFirmwareUpdate` | Auto-reboot behavior |
| — | `CanFlashBacklight` | Backlight during update |

---

## Security Architecture Summary

| Mechanism | Scope | Notes |
|-----------|-------|-------|
| AES-128-CBC (GID) | Firmware storage | Encryption = sole authentication |
| PKCS#7 code signing | Game executables + iTunesDB | RSA-SHA1, Apple CA chain |
| XML Digital Signatures | Game manifests | W3C XMLDSig standard |
| FairPlay DRM | Purchased content | Per-device key hierarchy |
| X.509 certificates | Trust chain for signing | Apple root CA embedded |
| No anti-rollback | Firmware versioning | Any version installable |
| No runtime integrity | Post-boot verification | No self-checking during execution |
| No TrustZone | Hardware level | ARM926EJ-S lacks TrustZone |

---

## Sources

- Firmware binary analysis (Ghidra, ARM926EJ-S)
- Cross-reference analysis for crypto functions
- String extraction for DRM identifiers
- W3C XMLDSig specification (for algorithm URI identification)
