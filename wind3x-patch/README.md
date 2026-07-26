# wInd3x: iPod Classic Device Support Patch

## Summary

This patch adds full iPod Classic (6th/7th generation, all hardware revisions 2007–2014) support to [wInd3x](https://github.com/freemyipod/wInd3x). The iPod Classic uses the S5L8702 SoC (same as Nano 3G), making it compatible with the same DFU exploit path.

Only `pkg/devices/devices.go` needs to be modified.

## Confirmed Working

| Command | Status | Notes |
|---------|--------|-------|
| `haxdfu` | ✅ | BootROM exploit succeeds immediately |
| `decrypt` | ✅ | Decrypts IMG1 OSOS via hardware AES (~2 hours) |
| `mse extract` | ✅ | Parses Classic IPSW into osos/rsrc/aupd/hash |
| `spew` | ✅ | Device detected correctly |

## What This Patch Adds

All four iPod Classic hardware revisions:

| Revision | Release | DFU PID | WTF PID | Models |
|----------|---------|---------|---------|--------|
| Initial (80/160GB) | Sep 2007 | `0x1223` | `0x1241` | MB029, MB147, MB145, MB150 |
| Rev A (120GB) | Sep 2008 | `0x1223` | `0x1245` | MB562, MB565 |
| Rev B (160GB) | Sep 2009 | `0x1223` | `0x1247` | MC293, MC297 ✅ tested |
| Rev C (160GB) | Oct 2012 | `0x1250` | `0x1250` | MD717, MD718 ✅ tested |

## Code Changes

1. Added `Classic6G`, `Classic6GA`, `Classic6GB`, `Classic6GC` Kind constants
2. Grouped all Classic variants with Nano3 in `SoCCode()` → returns `"8702"`
3. Grouped all Classic variants with Nano3 in `DFUVersion()` → returns `DFUProtoVersion1`
4. Added four device Description entries (one per hardware revision)

## Usage

```bash
# Replace the original devices.go
cp devices.go /path/to/wInd3x/pkg/devices/devices.go

# Or apply the patch
cd /path/to/wInd3x
git apply /path/to/devices.go.patch

# Build
go build ./cmd/wInd3x
```

See [BUILD.md](BUILD.md) for detailed build instructions.
See [TESTED_OUTPUT.md](TESTED_OUTPUT.md) for test evidence.

## DFU Mode Entry

1. Connect iPod Classic via 30-pin USB cable
2. Hold **Menu + Select (center button)** simultaneously
3. Keep holding through reboot (~8 seconds)
4. Release — screen completely blank (backlight on)
5. Device enumerates as `05AC:1223` (or `05AC:1250` for Rev C)

**Requirement:** Stock Apple firmware. Restore via iTunes if running Rockbox/EmCORE.

## Notes

- The MSE parser reports "Nano 3G" generation — this is expected (same SoC)
- The `-r` flag is essential for the ~2 hour decrypt process
- USB 2.0 ports are more reliable than USB 3.0 hubs for the exploit
