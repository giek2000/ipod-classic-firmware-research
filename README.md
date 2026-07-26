# iPod Firmware Research — Complete Analysis Archive

## Overview

This repository documents the decryption and analysis of **20 iPod firmware binaries** spanning every PortalPlayer-based iPod generation (2001–2009) plus the iPod Classic 7th Generation (S5L8702).

All firmware binaries were extracted from official Apple IPSW files and — where applicable — decrypted using the device's hardware AES engine via the [wInd3x](https://github.com/freemyipod/wInd3x) BootROM exploit.

## Analyzed Models

| Generation | Model | Firmware | Encrypted | Spec |
|-----------|-------|----------|-----------|------|
| 1st Gen (2001) | iPod 5/10GB | 1.1.5 | No | [Spec](specs/iPod_1st_Gen.md) |
| 3rd Gen (2003) | iPod Dock Connector | 2.2.3 | No | [Spec](specs/iPod_3rd_Gen.md) |
| Mini 1G (2004) | iPod Mini 4GB | 6.1.4.1 | No | [Spec](specs/iPod_Mini_1G.md) |
| Mini 2G (2005) | iPod Mini 4/6GB | 7.1.4.1 | No | [Spec](specs/iPod_Mini_2G.md) |
| 4th Gen Mono (2004) | Click Wheel 20/40GB | 4.3.1.1 | No | [Spec](specs/iPod_4th_Gen_Mono.md) |
| 4th Gen Color (2004) | Color/Photo 20-60GB | 11.1.2.1 | No | [Spec](specs/iPod_4th_Gen_Color.md) |
| 5th Gen Video (2005) | iPod Video 30/60GB | 13.1.2.1, 13.1.3 | No | [Spec](specs/iPod_5th_Gen_Video.md) |
| 5G Late 2006 | iPod Video 30/80GB | 20.1.2.1, 20.1.3 | No | [Spec](specs/iPod_5th_Gen_Video.md) |
| 5.5G Enhanced (2006) | iPod Video Enhanced | 25.1.2.1–25.1.3 | No | [Spec](specs/iPod_5_5G_Enhanced.md) |
| Nano 1G (2005) | iPod Nano 2/4GB | 14.1.3.1, 17.1.3.1 | No | [Spec](specs/iPod_Nano_1G.md) |
| Classic 7G (2009) | iPod Classic 160GB | 2.0.4 | Yes (AES) | [Spec](specs/iPod_Classic_7G_204.md) |
| Classic 7G (2012) | iPod Classic 160GB | 2.0.5 | Yes (AES) | [Spec](specs/iPod_Classic_7G_205.md) |

## Quick-Reference Feature Matrix

| Feature | 1G | 3G | Mini | 4G | 5G/5.5G | Nano 1G | Classic 7G |
|---------|----|----|------|----|---------|---------|------------|
| RTXC RTOS | ✅ | ✅ | — | — | — | — | ✅ |
| Pixo OS | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| FreeType2 | — | — | — | — | — | — | ✅ |
| MeCCA Codecs | — | — | — | — | — | — | ✅ |
| SQLite | — | — | — | — | — | — | ✅ |
| FairPlay DRM | — | — | — | — | — | — | ✅ |
| Cover Flow | — | — | — | — | — | — | ✅ |
| Nike+ iPod | — | — | — | — | — | ✅ | ✅ |
| Voice Memos | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Photos | — | ✅ | — | ✅ | ✅ | ✅ | ✅ |
| Games | — | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Video Playback | — | — | — | — | ✅ | — | ✅ |
| USB | — | — | — | — | ✅ | ✅ | ✅ |
| Debug Menu | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Disk Mode | — | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| FM Radio | — | — | — | — | — | — | ✅ |
| Genius | — | — | — | — | — | — | ✅ |
| Demo Mode | — | — | — | — | ✅ | — | ✅ |
| EQ | — | — | — | — | — | — | ✅ |

## Documentation Structure

```
├── README.md                     ← You are here
├── DECRYPTION_GUIDE.md           All-model decryption/extraction guide
├── LICENSE                       MIT License
├── wind3x-patch/                 wInd3x device descriptor patch
│   ├── README.md
│   ├── devices.go
│   ├── devices.go.patch
│   ├── BUILD.md
│   └── TESTED_OUTPUT.md
├── specs/                        Per-device feature specifications
│   ├── iPod_1st_Gen.md
│   ├── iPod_3rd_Gen.md
│   ├── iPod_Mini_1G.md
│   ├── iPod_Mini_2G.md
│   ├── iPod_4th_Gen_Mono.md
│   ├── iPod_4th_Gen_Color.md
│   ├── iPod_5th_Gen_Video.md
│   ├── iPod_5_5G_Enhanced.md
│   ├── iPod_Nano_1G.md
│   ├── iPod_Classic_7G_204.md
│   ├── iPod_Classic_7G_205.md
│   └── MASTER_COMPARISON.md
└── comparisons/                  Version-to-version diffs
    ├── Classic_204_vs_205.md
    ├── iPod_4th_Gen_Mono_vs_Color.md
    ├── iPod_5th_Gen_versions.md
    └── iPod_Nano_1G_2GB_vs_4GB.md
```

## Key Findings

### Architecture
- **PortalPlayer era (2001–2006):** PP5002/PP5020/PP5021/PP5022 SoCs running Pixo OS
- **Samsung era (2007–2014):** S5L8702 SoC running RTXC with "Silver" UI framework
- All iPods use ARM cores (ARM7TDMI → ARM926EJ-S)
- The Classic 7G is the only model with hardware AES firmware encryption

### Identical Binaries Discovered
- **iPod 4th Gen Mono:** FamilyID 4 and FamilyID 10 are byte-identical
- **iPod 4th Gen Color/Photo:** Color, Color FamilyID 11, and Photo FamilyID 5 are byte-identical
- **iPod Nano 1G:** 2GB (FamilyID 14) and 4GB (FamilyID 17) are byte-identical

Apple used different UpdaterFamilyIDs to route updates to specific hardware, but the firmware binaries themselves are often shared across capacity variants.

### Hidden Features (Classic 7G)
- **Debug Menu** — full internal debug interface
- **Demo Mode** — Apple Store kiosk mode (TCDemoMode controller)
- **RTXCbug** — interactive RTOS debugger
- **22 Logging Channels** — disk, USB, playback, photo import telemetry
- **Unit Test Framework** — built-in test suite

## Decryption Guide

See [`DECRYPTION_GUIDE.md`](DECRYPTION_GUIDE.md) for complete instructions covering all iPod models:
- **Unencrypted models (1G–5G, Nano 1G, Mini):** Simple IPSW extraction
- **Encrypted models (Classic 6G/7G):** Hardware AES decryption via wInd3x

## wInd3x Patch

See [`wind3x-patch/`](wind3x-patch/) for the device descriptor patch that adds iPod Classic support to wInd3x. Only `pkg/devices/devices.go` needs modification.

**Quick reference:**
```
DFU PID: 0x1223 (Rev A/B) / 0x1250 (Rev C)
WTF PID: 0x1247 (Rev B) / 0x1250 (Rev C)
SoC:     S5L8702 (same as Nano 3G)
Protocol: DFUProtoVersion1
```

## Naming Convention

Firmware binary filenames follow this pattern:
```
osos_decrypted_<Device>_<FirmwareVersion>.bin
```

Where:
- `osos` = Operating System section of the IPSW
- `<Device>` = iPod model identifier (e.g., `iPod_5th_Gen_Video`)
- `<FirmwareVersion>` = Apple's internal version (e.g., `13.1.3`)

The first number in the firmware version corresponds to the UpdaterFamilyID.

## References

- [freemyipod/wInd3x](https://github.com/freemyipod/wInd3x) — BootROM exploit and firmware tool
- [TheAppleWiki - USB Product IDs](https://theapplewiki.com/wiki/USB_Product_IDs)
- [TheAppleWiki - DFU Mode](https://theapplewiki.com/wiki/DFU_Mode)
- [TheAppleWiki - iPod Firmware](https://theapplewiki.com/wiki/Firmware)

## Legal

This research is conducted for educational and interoperability purposes under fair use principles. All analyzed iPod models are discontinued products. No copyrighted Apple code is distributed — only analysis results and documentation.

## License

MIT License — see [LICENSE](LICENSE).
