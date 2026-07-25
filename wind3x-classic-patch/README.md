# wInd3x: iPod Classic (6th Generation) Support

## Summary

This patch adds iPod Classic (6th generation, all hardware revisions 2007-2014) support to wInd3x. The iPod Classic shares the S5L8702 SoC and BootROM with the Nano 3G, making it compatible with the same DFU exploit path (DFUProtoVersion1). Only a device descriptor addition was required.

## Confirmed Working

| Command | Status | Notes |
|---------|--------|-------|
| `haxdfu` | ✅ | BootROM exploit succeeds immediately |
| `decrypt` | ✅ | Decrypts IMG1 OSOS via hardware AES engine (~2 hours) |
| `mse extract` | ✅ | Parses Classic IPSW firmware into osos/rsrc/aupd/hash |
| `spew` | ✅ | Device detected as "Classic 6G" |

## Tested Hardware

- **Model:** iPod Classic 160GB (MC293, Late 2009 / commonly called "7th gen")
- **Firmware:** RetailOS 2.0.4 (UpdaterFamilyID 35)
- **SoC:** Samsung S5L8702 (ARM926EJ-S)
- **Host:** Windows 11 + WSL2 Ubuntu 26.04 + usbipd-win 5.3.0

## USB Product IDs

All iPod Classic hardware revisions share the DFU PID `0x1223`. WTF PIDs vary by revision:

| Revision | Release | WTF PID | Model Numbers |
|----------|---------|---------|---------------|
| Initial (80/160GB) | Sep 2007 | `0x1241` | MB029, MB147, MB145, MB150 |
| Rev A (120GB) | Sep 2008 | `0x1245` | MB562, MB565 |
| Rev B (160GB) | Sep 2009 | `0x1247` | MC293, MC297 (tested ✅) |
| Rev C (160GB) | Oct 2012 | `0x1250` | MD717, MD718 (untested) |

Source: [TheAppleWiki USB Product IDs](https://theapplewiki.com/wiki/USB_Product_IDs)

## DFU Mode Entry

1. Connect iPod Classic to computer via 30-pin USB cable
2. Hold **Menu + Select (center button)** simultaneously
3. The device resets — **keep holding both buttons through the reboot**
4. After ~8 seconds total, release both buttons
5. Screen will be completely blank (backlight on, no icons/text)
6. Device enumerates as VID:PID `05AC:1223`

**Requirement:** The iPod must be running stock Apple firmware (restored via iTunes/Finder). Custom bootloaders (EmCORE, Rockbox bootloader) intercept the boot process differently. If this happens, restore the iPod to stock firmware first using iTunes, then enter DFU mode.

Source: [TheAppleWiki DFU Mode](https://theapplewiki.com/wiki/DFU_Mode) — section "iPod classic (6th generation)"

## Firmware Structure

The iPod Classic IPSW contains a file named `Firmware-35.X.X.X` which is in MSE format:

```
$ ./wInd3x mse extract Firmware-35.9.0.4 -o ./out/
2026/07/23 11:57:17 Parsing MSE for (guessed) generation: Nano 3G
2026/07/23 11:57:17 File 0: rsrc, offset 6000, len 4e00000, prefix: false
2026/07/23 11:57:17 File 1: osos, offset 4e07000, len a1ba53, prefix: false
2026/07/23 11:57:18 File 2: aupd, offset 5824000, len 11c8b3, prefix: false
2026/07/23 11:57:18 File 3: hash, offset 5942000, len 1000, prefix: false
```

| Section | Size | Description |
|---------|------|-------------|
| `rsrc` | 78 MB | FAT16 resource filesystem (fonts, games, training data) — unencrypted |
| `osos` | 10.1 MB | RetailOS binary — AES encrypted IMG1 |
| `aupd` | 1.1 MB | Firmware updater — AES encrypted IMG1 |
| `hash` | 4 KB | Integrity verification data |

## Changes

Only `pkg/devices/devices.go` is modified:

1. Added `Classic6G Kind = "classic6g"` constant
2. Added `"Classic 6G"` to `String()` switch
3. Added `Classic6G` to `SoCCode()` — returns `"8702"` (grouped with Nano3)
4. Added `Classic6G` to `DFUVersion()` — returns `DFUProtoVersion1`
5. Added device description entry:
   - DFU PID: `0x1223`
   - WTF PID: `0x1247`
   - Disk PID: `0x1261`
   - UpdaterFamilyID: `35`

## Notes for Future Work

- Only the Late 2009 WTF PID (`0x1247`) is registered. To support all revisions in WTF mode, additional entries (or entries per revision) would be needed for PIDs `0x1241`, `0x1245`, `0x1250`.
- `nor read` and `nand` commands are untested on Classic.
- The Classic firmware partition layout differs slightly from Nano — the MSE parser currently guesses "Nano 3G" generation but parses correctly regardless.
