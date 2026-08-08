# iPod Nano 3G vs iPod Classic 6G — Same SoC, Same Year

## Overview

The iPod Nano 3rd Generation and iPod Classic 6th Generation were released simultaneously
in September 2007 and share the same Samsung S5L8702 SoC (ARM926EJ-S). This document
compares these sibling devices to understand how Apple differentiated the same silicon
for two very different form factors.

## Hardware Comparison

| Specification | Nano 3G | Classic 6G |
|--------------|---------|------------|
| SoC | S5L8702 | S5L8702 |
| CPU | ARM926EJ-S (ARMv5TEJ) | ARM926EJ-S (ARMv5TEJ) |
| RAM | 64 MB | 64 MB |
| Storage | 4/8GB NAND Flash | 80/160GB CE-ATA HDD |
| Display | 320×240 2.0" | 320×240 2.5" |
| Connector | 30-pin | 30-pin |
| Form Factor | Flash player (52g) | Hard disk player (140g) |
| DFU PID | 0x1229 | 0x1223 |
| Codename | N46 | N25C |
| Release | September 2007 | September 2007 |

---

## Firmware Comparison

| Metric | Nano 3G (1.1.3) | Classic 6G (1.1.2) | Delta |
|--------|----------------|-------------------|-------|
| Binary Size | 10,792,304 | 9,865,904 | Nano is +926,400 (+9%) |
| Strings (>=6) | 27,446 | 25,540 | Nano +1,906 |
| ARM Functions | 17,473 | 15,950 | Nano +1,523 |
| Thumb Functions | 5,347 | 5,098 | Nano +249 |
| Total Functions | 22,820 | 21,048 | Nano +1,772 |

---

## String Overlap

| Metric | Count |
|--------|-------|
| Common to both | 25,139 |
| Only in Nano 3G | 2,307 |
| Only in Classic 6G | 401 |
| Overlap percentage | 91% of Nano / 98% of Classic |

---

## Feature Comparison

| Feature | Nano 3G | Classic 6G | Notes |
|---------|---------|------------|-------|
| RTXC RTOS | ✅ | ✅ | Same kernel |
| Silver UI Framework | ✅ | ✅ | Same framework |
| Cover Flow | ✅ | ✅ | Same rendering engine |
| MeCCA Codecs | ✅ | ✅ | Same codec pipeline |
| FreeType2 | ✅ | ✅ | Same font engine |
| FairPlay DRM | ✅ | ✅ | Same DRM system |
| Nike+ iPod | ✅ | ✅ | Extended on Nano (workout types) |
| Voice Memos | ✅ | ✅ | Same |
| Games | ✅ | ✅ | Same game engine |
| Video Playback | ✅ | ✅ | Same |
| Photos | ✅ | ✅ | Same |
| Demo Mode | ✅ | ✅ | Same |
| FM Radio | ✅ | ✅ | Both have it |
| Disk Mode | ✅ | ✅ | Classic only (HDD) |
| Clock/Alarms | ✅ | ✅ | Nano 3G has extended clock UI |
| CE-ATA HDD | — | ✅ | Classic only |
| NAND Flash | ✅ | — | Nano only |

---

## Key Differences

### Storage Architecture
The primary firmware difference is the storage subsystem:
- **Nano 3G:** Samsung Whimory 2.1 NAND FTL (flash translation layer)
- **Classic 6G:** CE-ATA hard disk with ATAWorkLoopTask for command queuing

### Nike+ Implementation
The Nano 3G has a significantly expanded Nike+ system:
- 71 unique workout handlers (calibration, calorie/distance/timed goals)
- 241 workout-specific screens
- The Classic has basic Nike+ only

### Clock/Alarm System
The Nano 3G has a full alarm clock UI (TCClock, TCAlarmMenu, multiple alarms)
The Classic 6G has basic time display only.

### Binary Size
The Nano 3G binary is **larger** despite having no hard disk driver, because:
- More Nike+ workout screens and handlers
- Full clock/alarm system
- NAND Whimory FTL driver (replaces smaller CE-ATA driver)
- Later firmware version (1.1.3 vs 1.1.2) with more bug fixes

### Same GID Key
Both devices share the same S5L8702 GID AES key. The wInd3x exploit
works identically on both — a single exploit covers the entire S5L8702 family.

---

## Conclusion

The Nano 3G and Classic 6G are **sibling firmwares** built from the same source tree
with different feature flags and storage drivers. They share ~80% of their code.
Understanding one directly aids reverse engineering of the other.
