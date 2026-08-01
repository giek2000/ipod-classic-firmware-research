# iPod Classic 7G — RSRC Resource Partition

## Overview

The RSRC partition is one of the four sections within the iPod's firmware image. Unlike OSOS and AUPD which are AES encrypted, **RSRC is stored as plaintext** — neither encrypted nor signed at the filesystem level.

On S5L8702 devices (iPod Classic, Nano 3G), RSRC contains a FAT16 filesystem with UI resources, game bundles, workout definitions, and fonts.

Analysis performed using Ghidra and Capstone disassembly.

---

## RSRC in the Firmware Image

| Section | Offset | Size | Encrypted |
|---------|--------|------|-----------|
| rsrc | 0x6000 | 78 MB | **No** (plaintext) |
| osos | 0x4E07000 | 10.1 MB | Yes (AES-128-CBC) |
| aupd | 0x5824000 | 1.1 MB | Yes (AES-128-CBC) |
| hash | 0x5942000 | 4 KB | No |

---

## FAT16 Filesystem Parameters

| Parameter | Value |
|-----------|-------|
| Volume Label | `IPODRESOURC` |
| OEM ID | `MTOOL399` |
| Bytes/sector | 512 |
| Sectors/cluster | 4 (cluster size: 2,048 bytes) |
| Total sectors | 159,744 |
| Total size | 81,788,928 bytes (78 MB) |
| FAT copies | 2 |
| FAT size | 156 sectors each |
| Root entries | 512 |
| Media type | 0xF0 |

### Filesystem Memory Layout

| Offset within RSRC | Content |
|--------------------|---------|
| 0x00000000 | Boot Sector (FAT16 BPB) |
| 0x00000200 | FAT Table 1 (79,872 bytes) |
| 0x00013C00 | FAT Table 2 (copy) |
| 0x00027800 | Root Directory (16,384 bytes) |
| 0x0002B800 | Data Region Start |

---

## Directory Structure

```
RESOURC/
├── FONTS/
│   ├── CJK11.TTF              — CJK character font (11pt, ~1.9 MB)
│   ├── CJK16.TTF              — CJK character font (16pt, ~2.9 MB)
│   ├── HELVETICA*.TTF         — UI system fonts (Helvetica family)
│   ├── MONOHOPE*.TTF          — Monospace font (MonoHope-LCD)
│   └── MONOHORIZONTAL*.SBI    — Bitmap fonts
│
├── GAMES/GAMES_RO/
│   ├── 11004/                 — iQuiz
│   │   ├── EXECUTABLES/       — ARM game binary (~500 KB)
│   │   ├── MANIFEST.PLIST     — SHA1 digest manifest
│   │   ├── MANIFEST.P7B       — PKCS#7 signature (3,175 bytes)
│   │   ├── IC-INFO.SID        — Game identification
│   │   ├── SOUNDS/            — Sound effects
│   │   ├── FONTS/             — Game-specific fonts
│   │   └── UI/                — UI textures (.IPD format)
│   │
│   ├── 11010/                 — Klondike
│   │   ├── EXECUTABLES/       — ARM game binary (~632 KB)
│   │   ├── AUDIO/             — Sound effects (WAV)
│   │   ├── TEXTURES/          — Card/UI graphics
│   │   └── G.M4A              — Background music
│   │
│   └── 12347/                 — Vortex
│       ├── EXECUTABLES/       — ARM game binary (~640 KB)
│       ├── MEDIA/SFX/         — Sound effects
│       └── TEX/               — Game textures
│
└── TRAINER/
    ├── 100CAL*.XML through 800CAL*.XML  — Calorie goals
    ├── 10K.XML, 5K.XML, 3K.XML         — Distance goals
    ├── *MINUTES*.XML                    — Time-based workouts
    ├── ADJUST*.XML                      — Calibration data
    └── SPORTS*.XML                      — Sport-specific (~38 KB)
```

---

## Content Categories

| Category | Contents | Approximate Size |
|----------|----------|-----------------|
| UI Fonts | Helvetica family, CJK11/16, MonoHope-LCD, bitmap SBI | ~8 MB |
| Game Bundles | iQuiz, Klondike, Vortex, others | ~50 MB |
| Training Data | Nike+ workout XMLs (calorie, distance, time) | ~2 MB |
| Photo Resources | Photo browsing support, artwork placeholders | ~1 MB |
| Game Audio | M4A and WAV sound effects/music per game | ~12 MB |
| Localization | Game-specific localization resources | ~5 MB |

---

## Game Bundle Structure

Each game in `GAMES_RO/{ID}/` follows a consistent layout:

| File/Directory | Purpose |
|---------------|---------|
| `EXECUTABLES/*.BIN` | ARM game binary (500–640 KB typical) |
| `MANIFEST.PLIST` | SHA1 digests of all game files |
| `MANIFEST.P7B` | PKCS#7 detached signature over manifest |
| `IC-INFO.SID` | Game identification/licensing (356 bytes) |
| `ITUNESARTWORK` | Album art placeholder |
| `SOUNDS/` or `AUDIO/` | Sound effects and music |
| `FONTS/` | Game-specific font files |
| `UI/` or `TEX/` or `GRAPHICS/` | Visual assets |
| `LOCALIZATION/` | Language-specific strings |
| `DATA/` | Game data files |

---

## Firmware String References to RSRC

| Offset | String | Context |
|--------|--------|---------|
| 0x000987D8 | `games_RO` | Read-only game mount point |
| 0x000987B4 | `gamestats_WO` | Write-only game stats |
| 0x000987C4 | `gamedata_ShareRW` | Shared game data |
| 0x000A27E0 | `Resources/Games` | RSRC games directory |
| 0x009BC3D1 | `Resources/Games/games_RO/` | Full games path |

---

## Security Model

### What IS Protected in RSRC

| Component | Protection | Mechanism |
|-----------|-----------|-----------|
| Game executables (.BIN) | PKCS#7 code signing | MANIFEST.P7B signature |
| Game assets (all files) | Integrity check | SHA1 digest in MANIFEST.PLIST |
| Game manifest itself | Digital signature | PKCS#7 over MANIFEST.PLIST |

### What Is NOT Protected in RSRC

| Component | Status |
|-----------|--------|
| RSRC filesystem structure | Plaintext, no integrity check |
| Font files (TTF/SBI) | Plaintext, no signature |
| Workout/training XMLs | Plaintext, no signature |
| Game audio (M4A/WAV) | Covered only by manifest SHA1 |
| Game textures/UI | Covered only by manifest SHA1 |
| iTunesArtwork placeholders | Plaintext |

### Security Observations

1. **Fonts are unprotected** — Any TTF can be placed in the FONTS directory; the firmware loads them via FreeType2 without verification
2. **Training data is unprotected** — XML workout definitions can be freely modified
3. **Game executables are signed** — PKCS#7 (RSA-SHA1) prevents unauthorized code execution
4. **Game assets require valid manifest** — Modifying any file invalidates the SHA1 in MANIFEST.PLIST, which is itself signed

---

## RSRC Size by Device

| Device | SoC | RSRC Size | Format |
|--------|-----|-----------|--------|
| iPod Classic 6G/7G | S5L8702 | 78 MB | FAT16 |
| iPod Nano 3G | S5L8702 | ~16 MB | FAT16 (presumed) |
| iPod 5.5G Video | PP5021 | ~5 MB | Raw blob |
| iPod Nano 1G | PP5022 | ~16 MB | Raw blob |

On PortalPlayer-era devices (pre-S5L), the RSRC section is a raw resource blob rather than a FAT16 filesystem. Resources are referenced by offset rather than filesystem path.

---

## Encryption Status Summary

| Section | iPod Classic (S5L8702) | iPod 5G (PP5021) |
|---------|----------------------|-------------------|
| RSRC | **Unencrypted** (plaintext FAT16) | **Unencrypted** (raw blob) |
| OSOS | AES-128-CBC (GID key) | Unencrypted |
| AUPD | AES-128-CBC (GID key) | Unencrypted |
| HASH | Unencrypted | N/A |

RSRC is **always unencrypted** regardless of device generation. Only OSOS and AUPD are AES encrypted on S5L-era devices.

---

## Sources

- Firmware section extraction and FAT16 filesystem analysis
- OSOS string references to RSRC paths
- Game bundle structure inspection
- Security model inference from PKCS#7 and XMLDSig presence
