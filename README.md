# iPod Firmware Research — Complete Analysis Archive

## Overview

This repository documents the decryption and analysis of **22 iPod firmware binaries** spanning every PortalPlayer and Samsung S5L-based iPod generation (2001–2014). This represents the most comprehensive publicly available firmware analysis of Apple's iPod product line.

All firmware binaries were extracted from official Apple IPSW files and — where applicable — decrypted using the device's hardware AES engine via the [wInd3x](https://github.com/freemyipod/wInd3x) BootROM exploit. Analysis performed using Ghidra and Capstone disassembly.

---

## Analyzed Models

| # | Generation | Model | Firmware | SoC | Encrypted | Spec |
|---|-----------|-------|----------|-----|-----------|------|
| 1 | 1st Gen (2001) | iPod 5/10GB | 1.1.5 | PP5002 | No | [Spec](specs/iPod_1st_Gen_1_1_5.md) |
| 2 | 3rd Gen (2003) | iPod Dock Connector | 2.2.3 | PP5002 | No | [Spec](specs/iPod_3rd_Gen_2_2_3.md) |
| 3 | Mini 1G (2004) | iPod Mini 4GB | 6.1.4.1 | PP5020 | No | [Spec](specs/iPod_Mini_1st_Gen_6_1_4_1.md) |
| 4 | Mini 2G (2005) | iPod Mini 4/6GB | 7.1.4.1 | PP5022 | No | [Spec](specs/iPod_Mini_2nd_Gen_7_1_4_1.md) |
| 5 | 4th Gen Mono (2004) | Click Wheel 20/40GB | 4.3.1.1 | PP5020 | No | [Spec](specs/iPod_4th_Gen_Mono_4_3_1_1.md) |
| 6 | 4th Gen Color (2004) | Color/Photo 20-60GB | 11.1.2.1 | PP5020 | No | [Spec](specs/iPod_4th_Gen_Color_11_1_2_1.md) |
| 7 | 5th Gen Video (2005) | iPod Video 30/60GB | 13.1.2.1 | PP5021 | No | [Spec](specs/iPod_5th_Gen_Video_13_1_2_1.md) |
| 8 | 5G Late 2006 | iPod Video 30/80GB | 20.1.2.1 | PP5021 | No | [Spec](specs/iPod_5th_Gen_Video_Late_20_1_2_1.md) |
| 9 | 5.5G Enhanced (2006) | iPod Video Enhanced | 25.1.2.1 | PP5022C | No | [Spec](specs/iPod_5_5G_Video_Enhanced_25_1_2_1.md) |
| 10 | Nano 1G (2005) | iPod Nano 2GB | 14.1.3.1 | PP5022 | No | [Spec](specs/iPod_Nano_1st_Gen_2GB_14_1_3_1.md) |
| 11 | Nano 1G (2005) | iPod Nano 4GB | 17.1.3.1 | PP5022 | No | [Spec](specs/iPod_Nano_1st_Gen_4GB_17_1_3_1.md) |
| 12 | **Nano 3G (2007)** | iPod Nano 4/8GB | 26.1.1.3 | **S5L8702** | **Yes (AES)** | [Spec](specs/iPod_Nano_3G.md) |
| 13 | Classic 7G (2009) | iPod Classic 160GB | 2.0.4 | S5L8702 | Yes (AES) | [Spec](specs/204.md) |
| 14 | Classic 7G (2012) | iPod Classic 160GB | 2.0.5 | S5L8702 | Yes (AES) | [Spec](specs/205.md) |

**Additional firmware dumps analyzed:** iPod 4th Gen Mono (FamilyID 10), 4th Gen Color (FamilyID 5), 5G Video (13.1.3, 20.1.3), 5.5G (25.1.2.3, 25.1.3), Classic 7G (2.0.4 no-DRM variant).

---

## Quick-Reference Feature Matrix

| Feature | 1G | 3G | Mini | 4G Mono | 4G Color | 5G/5.5G | Nano 1G | **Nano 3G** | Classic 7G |
|---------|----|----|------|---------|----------|---------|---------|-------------|------------|
| RTXC RTOS | ✅ | ✅ | — | — | — | — | — | **✅** | ✅ |
| Pixo OS | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| FreeType2 | — | — | — | — | — | — | — | **✅** | ✅ |
| MeCCA Codecs | — | — | — | — | — | — | — | **✅** | ✅ |
| SQLite | — | — | — | — | — | — | — | **✅** | ✅ |
| FairPlay DRM | — | — | — | — | — | — | — | **✅** | ✅ |
| Cover Flow | — | — | — | — | — | — | — | **✅** | ✅ |
| Nike+ iPod | — | — | — | — | — | — | ✅ | **✅** | ✅ |
| Voice Memos | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | **✅** | ✅ |
| Photos | — | ✅ | — | — | ✅ | ✅ | ✅ | **✅** | ✅ |
| Games | — | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | **✅** | ✅ |
| Video Playback | — | — | — | — | — | ✅ | — | **✅** | ✅ |
| USB | — | — | — | — | — | ✅ | ✅ | **✅** | ✅ |
| Debug Menu | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | **✅** | ✅ |
| Disk Mode | — | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — | ✅ |
| EQ Presets | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | **✅** | ✅ |
| FM Radio | — | — | — | — | — | — | — | — | ✅ |
| Genius | — | — | — | — | — | — | — | — | ✅ |
| Demo Mode | — | — | — | — | — | ✅ | — | **✅** | ✅ |
| HW AES Encryption | — | — | — | — | — | — | — | **✅** | ✅ |

---

## Architecture Deep Dives

| Document | Description |
|----------|-------------|
| [Boot Chain](architecture/BOOT_CHAIN.md) | Full boot sequence from BootROM → NOR → Disk → OSOS |
| [Memory Map](architecture/MEMORY_MAP.md) | S5L8702 physical memory map, peripheral registers, DRAM layout |
| [IPSW Format](architecture/IPSW_FORMAT.md) | IPSW container structure, MSE sections, IMG1 header format |
| [NOR Flash](architecture/NOR_FLASH.md) | NOR flash memory map, SysCfg structure, boot image locations |
| [RTOS Internals](architecture/RTOS_INTERNALS.md) | RTXC task table, scheduling model, kernel API |
| [Codec Framework](architecture/CODEC_FRAMEWORK.md) | MeCCA pipeline architecture, audio/video decode flow |
| [DRM Architecture](architecture/DRM_ARCHITECTURE.md) | FairPlay key hierarchy, AES hardware, game signing |
| [Cover Flow Engine](architecture/COVERFLOW_ENGINE.md) | Rendering pipeline, fixed-point math, DMA double-buffering |
| [RSRC Resources](architecture/RSRC_RESOURCES.md) | Resource partition format, contents, security model |

---

## IPSW Structure Overview

All iPod firmware is distributed as IPSW files (standard ZIP archives). Two formats exist:

| Format | Era | Contents | Encryption |
|--------|-----|----------|------------|
| **Format A** | PortalPlayer (2001–2006) + Classic/Nano 2-3 | `Firmware-{ID}.{ver}` + `manifest.plist` | None (PP) / AES (S5L) |
| **Format B** | Samsung S5L (Nano 4–7) | `{Board}.bootloader.rb3` + `Firmware.MSE` + `manifest.plist` | AES-128-CBC |

**53 IPSWs** cataloged across **18 generations**. See [architecture/IPSW_FORMAT.md](architecture/IPSW_FORMAT.md) for complete inventory.

---

## Boot Chain Overview

```
Power On → BootROM (silicon, 64KB) → NOR Bootloader (IMG1) → AES Decrypt → RetailOS (OSOS) → RTXC Kernel → Application Layer
```

The S5L8702 boot chain has no post-decryption hash verification — encryption serves as the sole authentication mechanism. See [architecture/BOOT_CHAIN.md](architecture/BOOT_CHAIN.md).

---

## NOR Flash Overview

| Region | Offset | Size | Content |
|--------|--------|------|---------|
| SysCfg | 0x00000 | 256 B | Device identity (serial, model, HW revision) |
| Boot Image 1 | 0x08000 | 104 KB | Primary IMG1 bootloader |
| Boot Image 2 | 0x22000 | 108 KB | Backup bootloader |
| Extended | 0x41D00 | 762 KB | Bootloader data / disk mode |

SysCfg tags: `SCfg`, `SrNm`, `FwId`, `HwId`, `HwVr`, `SwVr`, `MLBN`, `MMod`, `Regn`, `Codc`. See [architecture/NOR_FLASH.md](architecture/NOR_FLASH.md).

---

## Memory Map Overview (S5L8702)

| Address Range | Size | Description |
|--------------|------|-------------|
| 0x00000000–0x0000FFFF | 64 KB | BootROM (mirrored) |
| 0x08000000–0x0BFFFFFF | 64 MB | DRAM (main working memory) |
| 0x20000000–0x200FFFFF | 1 MB | NOR Flash (bootloader, SysCfg) |
| 0x22000000–0x2200FFFF | 64 KB | SRAM (RTXC kernel, DFU buffer) |
| 0x38000000–0x3FFFFFFF | 128 MB | Hardware Peripherals (MMIO) |

Peripheral reference counts from firmware analysis: GPIO (1,155 refs), LCD (0x38200000), Clock/Power (1,313 refs), Timer/IRQ (871 refs), AES Engine (866 refs). See [architecture/MEMORY_MAP.md](architecture/MEMORY_MAP.md).

---

## Documentation Structure

```
├── README.md                          ← You are here
├── DECRYPTION_GUIDE.md                All-model decryption/extraction guide
├── LICENSE                            MIT License
├── wind3x-patch/                      wInd3x device descriptor patch (code)
│   ├── README.md
│   ├── devices.go
│   ├── devices.go.patch
│   ├── BUILD.md
│   └── TESTED_OUTPUT.md
├── specs/                             Per-device firmware feature specifications
│   ├── iPod_1st_Gen_1_1_5.md
│   ├── iPod_3rd_Gen_2_2_3.md
│   ├── iPod_Mini_1st_Gen_6_1_4_1.md
│   ├── iPod_Mini_2nd_Gen_7_1_4_1.md
│   ├── iPod_4th_Gen_Mono_4_3_1_1.md
│   ├── iPod_4th_Gen_Color_11_1_2_1.md
│   ├── iPod_5th_Gen_Video_13_1_2_1.md
│   ├── iPod_5_5G_Video_Enhanced_25_1_2_1.md
│   ├── iPod_Nano_1st_Gen_2GB_14_1_3_1.md
│   ├── iPod_Nano_1st_Gen_4GB_17_1_3_1.md
│   ├── iPod_Nano_3G.md               ← NEW: Nano 3rd Generation
│   ├── 204.md                         iPod Classic 7G (FW 2.0.4)
│   └── 205.md                         iPod Classic 7G (FW 2.0.5)
├── comparisons/                       Version-to-version diffs
│   ├── Classic_204_vs_205.md
│   ├── Nano3G_vs_Classic.md           ← NEW: Three-way comparison
│   ├── iPod_4th_Gen_Mono_vs_Color.md
│   ├── iPod_5th_Gen_versions.md
│   └── iPod_Nano_1G_2GB_vs_4GB.md
└── architecture/                      ← NEW: Deep dive documents
    ├── BOOT_CHAIN.md
    ├── MEMORY_MAP.md
    ├── IPSW_FORMAT.md
    ├── NOR_FLASH.md
    ├── RTOS_INTERNALS.md
    ├── CODEC_FRAMEWORK.md
    ├── DRM_ARCHITECTURE.md
    ├── COVERFLOW_ENGINE.md
    └── RSRC_RESOURCES.md
```

---

## Key Findings

### Architecture
- **PortalPlayer era (2001–2006):** PP5002/PP5020/PP5021/PP5022 SoCs running Pixo OS
- **Samsung era (2007–2014):** S5L8702 SoC running RTXC with "Silver" UI framework
- All iPods use ARM cores (ARM7TDMI → ARM926EJ-S)
- The Nano 3G and Classic share the **same S5L8702 SoC** — first encrypted iPod with a hard disk
- Classic 7G firmware contains 17,721 ARM functions + 5,312 Thumb functions (23,033 total)

### Identical Binaries Discovered
- **iPod 4th Gen Mono:** FamilyID 4 and FamilyID 10 are byte-identical
- **iPod 4th Gen Color/Photo:** Color, Color FamilyID 11, and Photo FamilyID 5 are byte-identical
- **iPod Nano 1G:** 2GB (FamilyID 14) and 4GB (FamilyID 17) are byte-identical

### Hidden Features (S5L8702 Devices)
- **Debug Menu** — Full internal debug interface with disk browser, memory viewer, unit tests
- **Demo Mode** — Apple Store kiosk mode (TCDemoMode controller)
- **RTXCbug** — Interactive RTOS debugger with task/object inspection
- **22 Logging Channels** — Disk, USB, playback, photo import telemetry
- **Unit Test Framework** — Built-in test suite (TUnitTestSuiteCntlr)
- **MockupMode** — Developer UI prototyping tool

### Cover Flow — No GPU
The S5L8702 has **no GPU or hardware 2D/3D accelerator**. Cover Flow is achieved through highly optimized fixed-point ARM assembly rendering at ~30 FPS on a 200MHz CPU. Column-by-column affine texture mapping with DMA double-buffering.

---

## Decryption Guide

See [`DECRYPTION_GUIDE.md`](DECRYPTION_GUIDE.md) for complete instructions covering all iPod models:
- **Unencrypted models (1G–5G, Nano 1G, Mini):** Simple IPSW extraction
- **Encrypted models (Classic 6G/7G, Nano 3G+):** Hardware AES decryption via wInd3x

## wInd3x Patch

See [`wind3x-patch/`](wind3x-patch/) for the device descriptor patch that adds iPod Classic support to wInd3x.

**Quick reference:**
```
DFU PID: 0x1223 (Rev A/B) / 0x1250 (Rev C)
WTF PID: 0x1247 (Rev B) / 0x1250 (Rev C)
SoC:     S5L8702 (same as Nano 3G)
Protocol: DFUProtoVersion1
```

---

## Build Codenames

| Codename | Device | SoC |
|----------|--------|-----|
| N46 | iPod Nano 3rd Generation | S5L8702 |
| N25C | iPod Classic 7G (Rev C) | S5L8702 |
| N58s | iPod Nano 4th Generation | S5L8720 |
| N33 | iPod Nano 5th Generation | S5L8730 |
| N20 | iPod Nano 6th Generation | S5L8723 |
| N31 | iPod Nano 7th Generation | S5L8740 |

---

## References

- [freemyipod/wInd3x](https://github.com/freemyipod/wInd3x) — BootROM exploit and firmware tool
- [Olsro/reddit-ipod-guides](https://github.com/Olsro/reddit-ipod-guides) — iPod Classic modding guides
- [DavidBuchanan314/classic-ipod-tools](https://github.com/DavidBuchanan314/classic-ipod-tools) — NOR flash tools
- [TheAppleWiki - USB Product IDs](https://theapplewiki.com/wiki/USB_Product_IDs)
- [TheAppleWiki - DFU Mode](https://theapplewiki.com/wiki/DFU_Mode)
- [TheAppleWiki - iPod Firmware](https://theapplewiki.com/wiki/Firmware)

## Legal

This research is conducted for educational and interoperability purposes under fair use principles. All analyzed iPod models are discontinued products. No copyrighted Apple code is distributed — only analysis results and documentation.

## License

MIT License — see [LICENSE](LICENSE).
