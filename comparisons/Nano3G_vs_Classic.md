# iPod Nano 3G vs Classic 7G — Three-Way Firmware Comparison

## Overview

The iPod Nano 3G and iPod Classic 7G share the same Samsung S5L8702 SoC, the same RTXC RTOS, and the same Silver UI framework. This document compares the Nano 3G (FW 1.1.3) against the Classic 7G in both firmware versions (2.0.4 and 2.0.5).

Analysis performed using Ghidra and Capstone disassembly.

---

## Hardware Comparison

| Specification | Nano 3G | Classic 7G (2.0.4) | Classic 7G (2.0.5) |
|--------------|---------|--------------------|--------------------|
| SoC | S5L8702 | S5L8702 | S5L8702 |
| CPU | ARM926EJ-S | ARM926EJ-S | ARM926EJ-S |
| RAM | 64 MB | 64 MB | 64 MB |
| Storage | 4/8GB NAND | 160GB HDD (CE-ATA) | 160GB HDD (CE-ATA) |
| Display | 320×240 (2.0") | 320×240 (2.5") | 320×240 (2.5") |
| Connector | 30-pin | 30-pin | 30-pin |
| DFU PID | 0x1229 | 0x1223 / 0x1250 | 0x1223 / 0x1250 |
| Form Factor | Flash player | Hard disk player | Hard disk player |
| Release | September 2007 | September 2009 | October 2012 |

---

## Firmware Size Comparison

| Metric | Nano 3G (1.1.3) | Classic (2.0.4) | Classic (2.0.5) |
|--------|-----------------|-----------------|-----------------|
| OSOS Binary Size | 10,792,304 B | 10,599,920 B | 10,861,904 B |
| ARM Functions | 17,473 | 17,721 | ~17,900 |
| Thumb Functions | 5,347 | 5,315 | ~5,350 |
| Total Functions | 22,820 | 23,036 | ~23,250 |
| Total Strings | 53,545 | 55,243 | ~56,000 |
| IPSW Size (compressed) | 58.53 MB | 58.29 MB | 60.57 MB |
| Firmware Image (raw) | 89.51 MB | 89.27 MB | 91.30 MB |
| UpdaterFamilyID | 26 | 35 | 38 |

---

## Feature Comparison

| Feature | Nano 3G | Classic 2.0.4 | Classic 2.0.5 |
|---------|---------|---------------|---------------|
| Cover Flow | ✅ | ✅ | ✅ |
| Video Playback | ✅ | ✅ | ✅ |
| Photos/Slideshow | ✅ | ✅ | ✅ |
| Games (signed) | ✅ | ✅ | ✅ |
| FairPlay DRM | ✅ | ✅ | ✅ |
| Nike+ iPod | ✅ | ✅ | ✅ |
| Voice Memos | ✅ | ✅ | ✅ |
| Podcasts | ✅ | ✅ | ✅ |
| Audiobooks | ✅ | ✅ | ✅ |
| TV Shows | ✅ | ✅ | ✅ |
| Movie Rentals | ✅ | ✅ | ✅ |
| Calendar/Contacts | ✅ | ✅ | ✅ |
| World Clock | ✅ | ✅ | ✅ |
| Genius Playlists | ❌ | ✅ | ✅ |
| Genius Mixes | ❌ | ✅ | ✅ |
| FM Radio | ❌ | ✅ | ✅ |
| Disk Mode (USB MSC) | ❌ | ✅ | ✅ |
| VoiceOver | ❌ | ❌ | ✅ |
| iTunes U | ❌ | ✅ | ✅ |

---

## RTOS Comparison

| RTXC Metric | Nano 3G | Classic 2.0.4 | Notes |
|-------------|---------|---------------|-------|
| RTOS | RTXC | RTXC | Identical kernel |
| Kernel Location | 0x22000000 (SRAM) | 0x22000000 (SRAM) | Same address |
| Kernel API Functions | ~264 thunks | 264 thunks | Same API |
| Task Count | ~10 tasks | 11 tasks | Classic adds GeniusMixesTask |
| Scheduling | Preemptive priority | Preemptive priority | Identical model |
| RTXCbug present | ✅ | ✅ | Same debug tool |

### Task Differences

| Task | Nano 3G | Classic 7G | Notes |
|------|---------|------------|-------|
| HostOSTask | ✅ | ✅ | Main supervisor |
| MP3ExampleTask | ✅ | ✅ | Debug leftover |
| USBDeviceTask | ✅ | ✅ | USB handler |
| DiskReaderTask | ✅ | ✅ | Storage I/O (NAND vs HDD) |
| ATAWorkLoopTask | ❌ | ✅ | Classic only — HDD command queue |
| GeniusMixesTask | ❌ | ✅ | Classic only — Genius computation |
| TMusicLoadingTask | ✅ | ✅ | Library loader |
| MeCCAIOTask | ✅ | ✅ | Audio codec I/O |
| StreamCacheReadTask | ✅ | ✅ | Audio buffer fill |
| StreamCacheTimeoutTask | ✅ | ✅ | Buffer timeout handler |
| FirewireTask | ✅ | ✅ | Dock/charging handler |

---

## DRM & Security Comparison

| Security Feature | Nano 3G | Classic 2.0.4 | Classic 2.0.5 |
|-----------------|---------|---------------|---------------|
| Firmware Encryption | AES-128-CBC (GID) | AES-128-CBC (GID) | AES-128-CBC (GID) |
| Game Signing (PKCS#7) | ✅ | ✅ | ✅ |
| Game Manifest (XMLDSig) | ✅ | ✅ | ✅ |
| iTunesDB Signing | ✅ | ✅ | ✅ |
| FairPlay Content Keys | ✅ | ✅ | ✅ |
| Post-decrypt Hash Check | ❌ | ❌ | ❌ |
| Anti-rollback Fuses | ❌ | ❌ | ❌ |

---

## UI Controller Comparison (Notable Differences)

| Controller | Nano 3G | Classic 2.0.4 | Classic 2.0.5 |
|-----------|---------|---------------|---------------|
| TSilverMediaListCntlr_Genius | ❌ | ✅ | ✅ |
| TSilverMediaListCntlr_GeniusMixes | ❌ | ✅ | ✅ |
| TCSettings_RadioRegions | ❌ | ✅ | ✅ |
| TSilverMediaListCntlr_iTunesU | ❌ | ✅ | ✅ |
| TSilverTrainerCntlr | ✅ | ✅ | ✅ |
| TCVoiceMemos | ✅ | ✅ | ✅ |
| TCSportTimer | ✅ | ✅ | ✅ |
| TCDemoMode | ✅ | ✅ | ✅ |
| RTXCbug | ✅ | ✅ | ✅ |

---

## Storage Architecture Differences

| Aspect | Nano 3G | Classic 7G |
|--------|---------|------------|
| Storage Type | NAND Flash | CE-ATA Hard Disk |
| Controller | NAND FTL | ATA/CE-ATA (0x3CE00000) |
| Spin-up Time | None (instant) | ~1-2 seconds |
| Power Mgmt | Static | Spin-down after idle timeout |
| ATAWorkLoopTask | Not present | Present (128B, serializes disk access) |
| Firmware Partition | NAND partition | Hidden disk partition 1 (~40MB) |
| User Storage | Transparent NAND | FAT32 partition 2 |
| Disk Mode | Not available | Available (USB PID 0x1261) |

---

## Codec Pipeline (Identical Architecture)

Both devices use the same MeCCA codec framework:

| Component | Address (Nano 3G) | Address (Classic) | Notes |
|-----------|-------------------|-------------------|-------|
| Codec Registry | ~0x00150xxx | 0x00150AB0 | Same structure, offset varies |
| MeCCAIOTask | Present | 0x001F5354 | Same architecture |
| Supported Codecs | MP3, AAC, ALAC, WAV | MP3, AAC, ALAC, WAV | Identical |
| FairPlay Decode | ✅ | ✅ | Same DRM pipeline |
| I2S Output | 0x38400000 | 0x38400000 | Same peripheral |
| Cirrus Logic DAC | ✅ | ✅ | Same audio path |

---

## Key Insight: Shared Codebase

The Nano 3G and Classic firmwares are clearly built from the **same source tree** with conditional compilation:
- Same function naming conventions (TSilver*, TC*, Handle*)
- Same RTOS and kernel API (RTXC, 264 thunks at 0x22000000)
- Same UI framework (Silver, data-driven XML layouts)
- Same codec framework (MeCCA)
- Same DRM implementation (FairPlay, PKCS#7, XMLDSig)
- Classic adds: Genius, FM Radio, ATAWorkLoopTask, Disk Mode
- Nano adds: Nike+ training focus, NAND FTL layer

The firmware differences are primarily:
1. **Storage driver layer** — NAND FTL vs CE-ATA
2. **Feature flags** — Genius, Radio enabled/disabled at build time
3. **Hardware peripherals** — Different GPIO mappings for different form factor
