# iPod Classic 7th Generation — Firmware Research

## Overview

This repository contains research into the iPod Classic 7th Generation (Late 2009, 160GB, S5L8702 SoC). We have successfully decrypted and analyzed the RetailOS 2.0.4 firmware binary using the device's hardware AES engine via the [wInd3x](https://github.com/freemyipod/wInd3x) BootROM exploit — believed to be the first public documentation of this process for iPod Classic.

## Key Findings

### Firmware Architecture
- **SoC:** Samsung S5L8702 (ARM926EJ-S, ARMv5TEJ, ~200MHz)
- **RTOS:** RTXC (real-time operating system with built-in debugger "RTXCbug")
- **UI Framework:** "Silver" (TSilver* controller classes, data-driven layouts)
- **Codec Framework:** MeCCA (Apple's proprietary media codec architecture)
- **Database:** SQLite (embedded, with custom iPod VFS)
- **Font Engine:** FreeType2

### Firmware Statistics
| Metric | Value |
|--------|-------|
| Binary size | 10.6 MB |
| Functions | 17,721 |
| Strings | 55,243 |
| Languages | 20+ |
| ARM code start | Offset 0x800 (IMG1 header) |

### Hidden/Disabled Features Found
- **Debug Menu** — full debug interface accessible from Extras
- **Demo Mode** — Apple Store kiosk/demo mode (TCDemoMode controller)
- **Unit Test Framework** — built-in test suite with logging channels
- **RTXCbug Debugger** — interactive RTOS debugger with task inspection, stack analysis, semaphore/mailbox/queue monitoring
- **22 Logging Channels** — internal telemetry (disk activity, USB, playback, photo import, etc.)
- **Usage Analytics** — tracks listening time, playback sessions, backlight usage, disk spinup count

### Audio System
- **Supported codecs:** MP3 (Fraunhofer), AAC, Protected AAC (FairPlay), ALAC, WAV/AIFF, Audible
- **EQ presets:** 22 built-in (Bass Booster, Treble Booster, Rock, Pop, Jazz, Classical, Electronic, etc.)
- **DRM:** FairPlay with X.509 certificate chain, PKCS#7 game signing

### Game System
- ARM binaries loaded from RSRC FAT16 partition
- PKCS#7 signature verification (SHA1 + RSA)
- Defined error handling (signing error, version error, memory error screens)
- Platform versioning (GamesPlatformID, GamesPlatformVersion)

### Security
- FairPlay DRM certificates embedded
- Game code signing via PKCS#7 (Apple FairPlay CA)
- OpenSSL library embedded (RSA-SHA1, SSL2/3)
- Screen lock with PIN/passkey system
- Serial/resistor verification for accessories

## IPSW Structure

The iPod Classic IPSW (`iPod_35.2.0.4.ipsw`) is a ZIP containing a single MSE-format firmware image:

```
Firmware-35.9.0.4 (93.6 MB, MSE format)
├── rsrc (78 MB) — FAT16 resource filesystem
│   ├── Fonts (CJK, Helvetica, Monospace)
│   ├── Games (iQuiz, Klondike, Vortex — ARM binaries)
│   └── Trainer (Nike+ workout XML definitions)
├── osos (10.1 MB) — RetailOS binary (IMG1 encrypted)
├── aupd (1.1 MB) — Firmware updater (IMG1 encrypted)
└── hash (4 KB) — Integrity verification
```

## wInd3x iPod Classic Support

We confirmed and documented that the [wInd3x](https://github.com/freemyipod/wInd3x) tool works on iPod Classic (all hardware revisions) after applying our device descriptor patch:

| Command | Status | Notes |
|---------|--------|-------|
| `haxdfu` | ✅ Working | BootROM exploit succeeds |
| `decrypt` | ✅ Working | Full 10.1MB OSOS decryption via hardware AES (~2 hours) |
| `mse extract` | ✅ Working | Parses Classic IPSW into osos/rsrc/aupd/hash |
| `spew` | ✅ Working | Device detected as Classic |

### Device Descriptor Patch

Only `pkg/devices/devices.go` needs modification to add Classic support:

```go
Classic6G Kind = "classic6g"
// DFU PID: 0x1223, WTF PID: 0x1247, Disk PID: 0x1261
// UpdaterFamilyID: 35, SoCCode: "8702", DFUProtoVersion1
```

See `wind3x-classic-patch/` for the complete patch, build instructions, and test evidence.

### USB Product IDs (from TheAppleWiki)

| Revision | Release | DFU PID | WTF PID |
|----------|---------|---------|---------|
| All Classics | — | `0x1223` | varies |
| Rev B (160GB) | Sep 2009 | `0x1223` | `0x1247` |
| Rev C (160GB) | Oct 2012 | `0x1223` | `0x1250` |

### DFU Mode Entry

1. Connect iPod via USB (30-pin cable)
2. Hold **Menu + Select (center)** simultaneously
3. Keep holding through reboot (~8 seconds total)
4. Release — screen stays completely blank
5. Device enumerates as `05AC:1223`

**Note:** Requires stock Apple firmware. Custom bootloaders (EmCORE) intercept and enter WTF mode instead.

## Detailed Firmware Feature Specification

See [`RETAILOS_FEATURE_SPEC.md`](RETAILOS_FEATURE_SPEC.md) for the full annotated specification — every controller, handler, screen, setting, and hidden feature found in the firmware, with hex offsets into the binary.

## Decryption Guide

See [`DECRYPTION_GUIDE.md`](DECRYPTION_GUIDE.md) for a complete step-by-step guide to decrypting the iPod Classic OSOS firmware using wInd3x and the device's hardware AES engine.

## References

- [freemyipod/wInd3x](https://github.com/freemyipod/wInd3x) — BootROM exploit and firmware tool
- [TheAppleWiki - USB Product IDs](https://theapplewiki.com/wiki/USB_Product_IDs)
- [TheAppleWiki - DFU Mode](https://theapplewiki.com/wiki/DFU_Mode)
- [Olsro/reddit-ipod-guides](https://github.com/Olsro/reddit-ipod-guides) — iPod Classic modding guides
- [DavidBuchanan314/classic-ipod-tools](https://github.com/DavidBuchanan314/classic-ipod-tools) — NOR flash tools

## Legal

This research is conducted for educational and interoperability purposes under fair use principles. The iPod Classic is a discontinued product (2014). No copyrighted Apple code is distributed in this repository — only analysis results and documentation.

## License

MIT License
