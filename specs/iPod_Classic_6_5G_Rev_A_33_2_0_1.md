# iPod Classic 6.5G (Rev A 120GB) — RetailOS 2.0.1 Firmware Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 2.0.1 |
| **IPSW** | iPod_33.2.0.1.ipsw |
| **Device** | iPod Classic 6.5G 120GB (2008, thinner single-platter) |
| **UpdaterFamilyID** | 33 |
| **Binary Size** | 10,514,000 bytes (10.03 MB) |
| **ARM Code Start** | 0x800 (encrypted) |
| **ARM Code Size** | 10,511,952 bytes |
| **SoC** | Samsung S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Encrypted** | Yes (HW AES-128-CBC) |
| **Decryption Status** | ❌ NOT DECRYPTED |
| **Code Entropy** | 7.9970 bits/byte (confirms encryption) |
| **SHA-256** | `d057642cdda966e367cfba008c27cc80892dc97f29959e79f538bf04b03da59a` |

---

## Encryption Status

This firmware binary remains **AES-encrypted**. The code section (offset 0x800 onwards) has entropy of 7.997 bits/byte — indistinguishable from random data. No code analysis, string extraction, or feature detection is possible until the binary is decrypted using the S5L8702 GID key via the wInd3x exploit on a physical iPod Classic.

### Evidence of Encryption

| Check | Result | Expected if Decrypted |
|-------|--------|----------------------|
| First instruction (0x800) | 0xA7B08B77 (invalid) | ARM branch (0xEAxxxxxx) |
| Entropy (0x800 + 64KB) | 7.9970 | ~5.5–6.0 |
| Meaningful strings | 0 | ~50,000+ |
| ARM function prologues | 0 identifiable | ~17,000+ |

---

## IMG1 Header (Unencrypted)

| Field | Value |
|-------|-------|
| Magic | `8702` (ASCII) |
| Format Version | `1.0` |
| Entry Point | 0x00A06650 |
| Entry Point (repeated) | 0x00A06650 × 3 |
| Header size | 2,048 bytes (0x800) |

The entry point (0x00A06650 = ~10.02 MB offset) places the kernel initialization near the end of the binary — consistent with the pattern observed in all S5L8702 firmware where RTXC startup code is linked last.

---

## Hardware Information

| Attribute | Value |
|-----------|-------|
| Release date | September 9, 2008 |
| Model numbers | MB562 (120GB Silver), MB565 (120GB Black) |
| Codename | N25A |
| SoC | Samsung S5L8702 (same as 6G) |
| CPU | ARM926EJ-S @ 200 MHz |
| RAM | 64 MB SDRAM |
| Storage | 120GB 1.8" Toshiba HDD (single-platter, thinner) |
| Display | 2.5" TFT LCD, 320×240 (QVGA) |
| Battery | 580 mAh |
| Connectivity | USB 2.0, 30-pin dock |
| DFU USB PID | 0x1223 |

---

## Key Differences from 6G

| Aspect | 6G (FamilyID 24) | 6.5G Rev A (FamilyID 33) |
|--------|-------------------|--------------------------|
| Release | September 2007 | September 2008 |
| Capacity | 80GB / 160GB | 120GB |
| Form factor | Thick (dual-platter 160GB) | Thinner (single-platter) |
| Firmware base | 1.x | 2.x |
| Genius support | No | **Yes** (added with iTunes 8) |
| Binary size | 9.47 MB | 10.03 MB (+5.9%) |
| Code growth | — | +587,472 bytes |

The 5.9% growth from 6G→6.5G is the largest single-generation code increase in the Classic lineage, likely due to the addition of Genius and related infrastructure.

---

## Expected Features (Inferred)

Since code analysis is blocked by encryption, features are inferred from device capabilities and firmware version:

| Feature | Expected | Source |
|---------|----------|--------|
| RTXC RTOS | ✅ | Same platform as decrypted 7G |
| Silver UI Framework | ✅ | Same platform as decrypted 7G |
| Cover Flow | ✅ | Standard since 6G launch |
| FairPlay DRM | ✅ | iTunes Store compatibility |
| MeCCA Codec Framework | ✅ | AAC/MP3/ALAC/H.264 support |
| SQLite | ✅ | Media database engine |
| FreeType2 | ✅ | Font rendering |
| Nike+ iPod | ✅ | Standard since 6G |
| **Genius** | **✅** | **Added with FW 2.0 / iTunes 8 (Sept 2008)** |
| Games (Click Wheel) | ✅ | Standard feature |
| Photos | ✅ | Standard feature |
| Video Playback | ✅ | H.264/MPEG-4 playback |
| Voice Memos | ✅ | Standard Extras feature |
| Demo Mode | ✅ | Apple Store kiosk mode |
| Debug Menu | ✅ | Present in all S5L firmware |
| Disk Mode | ✅ | Standard feature |
| USB Audio | ✅ | Standard feature |
| FM Radio | ✅ | Likely (supported with accessory) |
| EU Volume Limit | ❓ | Uncertain for 2.0.1 |
| Genius Mixes | ❓ | May require 2.0.3+ |

---

## Firmware Size in Context

| Model | FamilyID | FW Ver | Size | Entry Point | Δ |
|-------|----------|--------|------|-------------|---|
| Classic 6G | 24 | 1.1.2 | 9.47 MB | 0x00976F80 | — |
| **Classic 6.5G** | **33** | **2.0.1** | **10.03 MB** | **0x00A06650** | **+5.9%** |
| Classic 7G Rev B | 35 | 2.0.4 | 10.11 MB | 0x00A1B5F0 | +0.8% |
| Classic 7G Rev C | 38 | 2.0.5 | 10.14 MB | — | +0.3% |

---

## Decryption Path

To produce a full analysis matching the 7G specs:

1. Obtain iPod Classic 6.5G hardware (or 7G Rev B which shares SoC/GID)
2. Put device in DFU mode (USB PID 0x1223)
3. Run `wInd3x haxdfu` to exploit BootROM
4. Stream-decrypt the OSOS section using the device's AES engine
5. Verify decryption: entropy should drop to ~5.5, ARM branches at 0x800

**Cross-device decryption:** All S5L8702-based Classics (6G/6.5G/7G Rev A/B) with DFU PID 0x1223 share the same GID key. A 7G Rev B should be able to decrypt this firmware.

---

## Firmware Version History

| Version | Date | Notes |
|---------|------|-------|
| 2.0 | Sept 2008 | Ship firmware (Genius support) |
| **2.0.1** | **Oct 2008** | **← This firmware** |
| 2.0.2 | Nov 2008 | Final for FamilyID 33 |
