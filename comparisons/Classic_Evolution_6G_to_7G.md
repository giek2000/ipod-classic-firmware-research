# iPod Classic Firmware Evolution: 6G (1.0.1) through 7G (2.0.5)

## Overview

This document tracks the complete firmware evolution of the iPod Classic across
8 firmware versions spanning 5 years (2007–2012), all on the same S5L8702 SoC.
The iPod Nano 3G is included as a reference point — it shares the same SoC and launched simultaneously.

## Firmware Size & Function Growth

| Version | FamilyID | Year | Binary Size | Strings | ARM Funcs | Thumb Funcs | Total Funcs |
|---------|----------|------|-------------|---------|-----------|-------------|-------------|
| 6G 1.0.1 | 24 | 2007 | 9,177,008 | 22,969 | 15,410 | 4,555 | 19,965 |
| 6G 1.0.3 | 24 | 2007 | 9,493,328 | 24,446 | 15,735 | 4,722 | 20,457 |
| 6G 1.1 | 24 | 2008 | 9,865,856 | 25,550 | 15,950 | 5,107 | 21,057 |
| 6G 1.1.2 | 24 | 2008 | 9,926,528 | 25,696 | 16,096 | 5,107 | 21,203 |
| 6.5G 2.0 | 33 | 2008 | 10,509,968 | 27,317 | 17,394 | 7,054 | 24,448 |
| 6.5G 2.0.1 | 33 | 2008 | 10,514,000 | 27,560 | 17,413 | 5,397 | 22,810 |
| 7G 2.0.4 | 35 | 2009 | 10,599,920 | 27,754 | 17,721 | 5,315 | 23,036 |
| 7G 2.0.5 | 38 | 2012 | 10,634,528 | 28,050 | 17,762 | 5,402 | 23,164 |
| *Nano 3G 1.1.3* | *26* | *2007* | *10,792,304* | *27,446* | *17,473* | *5,347* | *22,820* |

---

## Version-to-Version Growth

| Transition | Size Delta | % Growth | Func Delta | Key Changes |
|-----------|-----------|---------|-----------|-------------|
| 6G 1.0.1 → 6G 1.0.3 | +316,320 | +3% | +492 | Bug fixes |
| 6G 1.0.3 → 6G 1.1 | +372,528 | +3% | +600 | Feature additions |
| 6G 1.1 → 6G 1.1.2 | +60,672 | +0% | +146 | Minor update |
| 6G 1.1.2 → 6.5G 2.0 | +583,440 | +5% | +3,245 | **Genius added**, new hardware revision |
| 6.5G 2.0 → 6.5G 2.0.1 | +4,032 | +0% | +-1,638 | Bug fixes |
| 6.5G 2.0.1 → 7G 2.0.4 | +85,920 | +0% | +226 | New hardware revision (Rev B) |
| 7G 2.0.4 → 7G 2.0.5 | +34,608 | +0% | +128 | **EU Volume Limit**, FreeType2 update |

---

## Feature Timeline

| Feature | 6G 1.0.1 | 6G 1.0.3 | 6G 1.1 | 6G 1.1.2 | 6.5G 2.0 | 6.5G 2.0.1 | 7G 2.0.4 | 7G 2.0.5 |
|---------|----------|----------|--------|----------|----------|-----------|----------|----------|
| Cover Flow | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Genius | — | — | — | — | ✅ | ✅ | ✅ | ✅ |
| Genius Mixes | — | — | — | — | — | — | ✅ | ✅ |
| FM Radio | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Games | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Video Playback | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Voice Memos | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Nike+ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| FairPlay DRM | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| EU Volume Limit | — | — | — | — | — | — | — | ✅ |
| Demo Mode | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| SQLite | — | — | — | — | ✅ | ✅ | ✅ | ✅ |

---

## String Overlap Between All Versions

Strings common to ALL 8 Classic firmware versions: **21,064**

| Comparison | Common | Unique to Earlier | Unique to Later |
|-----------|--------|------------------|----------------|
| 6G 1.0.1 vs 6G 1.0.3 | 21,625 | 1,344 | 2,821 |
| 6G 1.0.3 vs 6G 1.1 | 23,782 | 664 | 1,768 |
| 6G 1.1 vs 6G 1.1.2 | 25,366 | 184 | 330 |
| 6G 1.1.2 vs 6.5G 2.0 | 25,104 | 592 | 2,213 |
| 6.5G 2.0 vs 6.5G 2.0.1 | 27,066 | 251 | 494 |
| 6.5G 2.0.1 vs 7G 2.0.4 | 26,857 | 703 | 897 |
| 7G 2.0.4 vs 7G 2.0.5 | 27,261 | 493 | 789 |

---

## Technical Observations

### 1. Consistent Growth Pattern
- Total growth from first to last: 1,457,520 bytes (+15%)
- Function count grew from 19,965 to 23,164 (+3,199)
- The 6G→6.5G transition is the largest single jump (Genius + SQLite)

### 2. Same SoC Throughout
All versions use the S5L8702 (ARM926EJ-S, ARMv5TEJ). No architecture changes.
Hardware revisions only affect HDD capacity and minor board components.

### 3. Build Codename: N25C
Every Classic firmware uses the N25C codename. The build numbers increment:
- 6G: N25CFirmwareWin-XX (earliest builds)
- 6.5G: N25CFirmwareWin-XX (higher build numbers)
- 7G 2.0.4: N25CFirmwareWin-75
- 7G 2.0.5: N25CFirmwareWin-247
