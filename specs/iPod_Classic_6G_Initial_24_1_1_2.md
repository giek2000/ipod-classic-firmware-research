# iPod Classic 6G (Initial) — RetailOS 1.1.2 Firmware Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.1.2 |
| **IPSW** | iPod_24.1.1.2.ipsw |
| **Device** | iPod Classic 6G 80/160GB (2007, first Click Wheel Classic) |
| **UpdaterFamilyID** | 24 |
| **Binary Size** | 9,926,528 bytes (9.47 MB) |
| **ARM Code Start** | 0x800 (encrypted) |
| **ARM Code Size** | 9,924,480 bytes |
| **SoC** | Samsung S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Encrypted** | Yes (HW AES-128-CBC) |
| **Decryption Status** | ❌ NOT DECRYPTED |
| **Code Entropy** | 7.9969 bits/byte (confirms encryption) |
| **SHA-256** | `e3093a571ba005a1220ce5ee89a91a0440c4778b7d2a974eac32f61fc673d6f2` |

---

## Encryption Status

This firmware binary remains **AES-encrypted**. The code section (offset 0x800 onwards) has entropy of 7.997 bits/byte — indistinguishable from random data. No code analysis, string extraction, or feature detection is possible until the binary is decrypted using the S5L8702 GID key via the wInd3x exploit on a physical iPod Classic.

### Evidence of Encryption

| Check | Result | Expected if Decrypted |
|-------|--------|----------------------|
| First instruction (0x800) | 0xFA41DA2C (invalid) | ARM branch (0xEAxxxxxx) |
| Entropy (0x800 + 64KB) | 7.9969 | ~5.5–6.0 |
| Meaningful strings | 0 | ~50,000+ |
| ARM function prologues | 0 identifiable | ~17,000+ |

---

## IMG1 Header (Unencrypted)

| Field | Value |
|-------|-------|
| Magic | `8702` (ASCII) |
| Format Version | `1.0` |
| Entry Point | 0x00976F80 |
| Entry Point (repeated) | 0x00976F80 × 3 |
| Header size | 2,048 bytes (0x800) |

The entry point (0x00976F80 = ~9.46 MB offset) is near the end of the binary, consistent with RTXC kernel initialization code being linked last — matching the pattern seen in decrypted 7G firmware.

---

## Hardware Information

| Attribute | Value |
|-----------|-------|
| Release date | September 5, 2007 |
| Model numbers | MB029 (80GB Silver), MB145 (160GB Silver), MB147 (160GB Black) |
| Codename | N73 (80GB) / N25 (160GB) |
| SoC | Samsung S5L8702 |
| CPU | ARM926EJ-S @ 200 MHz |
| RAM | 32 MB (80GB model) / 64 MB (160GB model) |
| Storage | 80GB or 160GB 1.8" Toshiba HDD |
| Display | 2.5" TFT LCD, 320×240 (QVGA) |
| Battery | 580 mAh (80GB) / 850 mAh (160GB) |
| Connectivity | USB 2.0, 30-pin dock |
| DFU USB PID | 0x1223 |

---

## Expected Features (Inferred)

Since code analysis is blocked by encryption, features are inferred from the device's documented capabilities at launch:

| Feature | Expected | Source |
|---------|----------|--------|
| RTXC RTOS | ✅ | Same platform as decrypted 7G |
| Silver UI Framework | ✅ | Same platform as decrypted 7G |
| Cover Flow | ✅ | Flagship marketing feature at 6G launch |
| FairPlay DRM | ✅ | iTunes Store compatibility |
| MeCCA Codec Framework | ✅ | AAC/MP3/ALAC/H.264 support |
| SQLite | ✅ | Media database engine |
| FreeType2 | ✅ | Font rendering |
| Nike+ iPod | ✅ | Documented at launch |
| Games (Click Wheel) | ✅ | Supported at launch |
| Photos | ✅ | Supported at launch |
| Video Playback | ✅ | H.264/MPEG-4 playback |
| Voice Memos | ✅ | Standard Extras feature |
| Demo Mode | ✅ | Apple Store kiosk mode |
| Debug Menu | ✅ | Present in all S5L firmware |
| Disk Mode | ✅ | Standard feature |
| USB Audio | ✅ | Standard feature |
| Genius | ❌ | Not introduced until Sept 2008 |
| EU Volume Limit | ❓ | Unknown for 1.1.2 |
| FM Radio | ❓ | May require accessory support |

---

## Decryption Path

To produce a full analysis matching the 7G specs:

1. Obtain iPod Classic 6G hardware (or 7G Rev B which shares SoC/GID)
2. Put device in DFU mode (USB PID 0x1223)
3. Run `wInd3x haxdfu` to exploit BootROM
4. Stream-decrypt the OSOS section using the device's AES engine
5. Verify decryption: entropy should drop to ~5.5, ARM branches should appear at 0x800

**Cross-device decryption note:** Since all S5L8702-based iPods (6G, 6.5G, 7G Rev A/B) share the same GID key, a 7G Rev B device should be able to decrypt 6G firmware.

---

## Firmware Version History

| Version | Date | Notes |
|---------|------|-------|
| 1.0 | Sept 2007 | Ship firmware |
| 1.0.1 | Oct 2007 | Bug fixes |
| 1.0.2 | Nov 2007 | Stability |
| 1.0.3 | Jan 2008 | Bug fixes |
| 1.1 | Jan 2008 | Feature update |
| 1.1.1 | Feb 2008 | Bug fixes |
| **1.1.2** | **Jul 2008** | **← This firmware (final for FamilyID 24)** |
