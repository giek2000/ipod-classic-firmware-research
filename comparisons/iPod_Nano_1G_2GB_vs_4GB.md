# iPod Nano 1st Generation — 2GB vs 4GB Comparison

## Key Finding

**The 2GB (FamilyID 14) and 4GB (FamilyID 17) firmware binaries are byte-for-byte identical.**

Both produce identical SHA-256 hashes, confirming Apple shipped the exact same firmware binary regardless of flash capacity.

## Verification

| Metric | 2GB (FamilyID 14) | 4GB (FamilyID 17) |
|--------|-------------------|-------------------|
| Binary Size | 22,905,856 bytes | 22,905,856 bytes |
| Strings | 54,684 | 54,684 |
| Functions | 11,306 | 11,306 |
| SHA-256 | (identical) | (identical) |

## Implications

- Storage capacity differences are handled purely in hardware/config, not in the firmware binary
- The UpdaterFamilyID distinction is for Apple's update mechanism to route the correct IPSW to each model
- For research purposes, only one copy of this firmware needs to be analyzed

## Firmware Details

- **Size:** 22,905,856 bytes (21.84 MB)
- **Strings:** 54,684
- **Functions:** 11,306
- **SoC:** PortalPlayer PP5021C
- **Encrypted:** No

## Features

- ✅ MeCCA (Codec Framework)
- ✅ FreeType2 Fonts
- ✅ Nike+ iPod
- ✅ Voice Memos
- ✅ Photos
- ✅ Debug Menu
- ✅ Disk Mode
- ✅ USB Audio
- ✅ Pedometer
- ✅ PortalPlayer

### Unique to Nano 1G (vs iPod Video)

- **Pedometer** — step counting capability (earliest iPod with this feature)
- Smaller binary despite more strings (efficient flash-optimized code)

