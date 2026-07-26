# iPod Firmware — Cross-Generation Master Comparison

## All Analyzed Firmware Binaries

| Device | Firmware | Size | Strings | Functions | SoC | Encrypted |
|--------|----------|------|---------|-----------|-----|----------|
| Classic 7G204 | Unknown | 10.11 MB | 55,243 | 17,721 | S5L8702 | Yes (AES) |
| Classic 7G205 | Unknown | 10.14 MB | 55,712 | 17,762 | S5L8702 | Yes (AES) |
| iPod 1st Gen 1.1.5 | 1.1.5 | 4.83 MB | 26,630 | 6,885 | PortalPlayer | No |
| iPod 3rd Gen 2.2.3 | 2.2.3 | 4.35 MB | 9,755 | 7,192 | PortalPlayer | No |
| iPod 4th Gen Color | Unknown | 6.21 MB | 12,965 | 9,777 | PortalPlayer | No |
| iPod 4th Gen Color 11.1.2.1 | 11.1.2.1 | 6.21 MB | 12,965 | 9,777 | PortalPlayer | No |
| iPod 4th Gen Mono 10.3.1.1 | 10.3.1.1 | 4.39 MB | 10,743 | 7,824 | PortalPlayer | No |
| iPod 4th Gen Mono 4.3.1.1 | 4.3.1.1 | 4.39 MB | 10,743 | 7,824 | PortalPlayer | No |
| iPod 4th Gen Photo 5.1.2.1 | 5.1.2.1 | 6.21 MB | 12,965 | 9,777 | PortalPlayer | No |
| iPod 5 5G Video Enhanced 25.1.2.1 | 25.1.2.1 | 13.20 MB | 29,898 | 12,640 | PortalPlayer | No |
| iPod 5 5G Video Enhanced 25.1.2.3 | 25.1.2.3 | 13.25 MB | 30,023 | 12,745 | PortalPlayer | No |
| iPod 5 5G Video Enhanced 25.1.3 | 25.1.3 | 13.26 MB | 30,181 | 12,890 | PortalPlayer | No |
| iPod 5th Gen Video 13.1.2.1 | 13.1.2.1 | 13.19 MB | 29,921 | 12,639 | PortalPlayer | No |
| iPod 5th Gen Video 13.1.3 | 13.1.3 | 13.25 MB | 30,182 | 12,890 | PortalPlayer | No |
| iPod 5th Gen Video Late 20.1.2.1 | 20.1.2.1 | 13.19 MB | 29,921 | 12,639 | PortalPlayer | No |
| iPod 5th Gen Video Late 20.1.3 | 20.1.3 | 13.25 MB | 30,182 | 12,890 | PortalPlayer | No |
| iPod Mini 1st Gen 6.1.4.1 | 6.1.4.1 | 4.30 MB | 9,944 | 7,563 | PortalPlayer | No |
| iPod Mini 2nd Gen 7.1.4.1 | 7.1.4.1 | 4.30 MB | 9,856 | 7,561 | PortalPlayer | No |
| iPod Nano 1st Gen 2GB 14.1.3.1 | 14.1.3.1 | 21.84 MB | 54,684 | 11,306 | PortalPlayer | No |
| iPod Nano 1st Gen 4GB 17.1.3.1 | 17.1.3.1 | 21.84 MB | 54,684 | 11,306 | PortalPlayer | No |

## Feature Matrix

| Feature | 1G | 3G | Mini | 4G Mono | 4G Color | 5G Video | 5.5G | Nano 1G | Classic 7G |
|---------|----|----|------|---------|----------|----------|------|---------|------------|
| Cover Flow | — | — | — | — | — | — | — | — | ✅ |
| Debug Menu | ✅ | ✅ | ✅ | ✅ | — | ✅ | ✅ | ✅ | ✅ |
| Demo Mode | — | — | — | — | — | — | — | — | ✅ |
| Disk Mode | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| EU Volume Limit | — | — | — | — | — | — | — | — | ✅ |
| FM Radio | — | — | — | — | — | — | — | — | ✅ |
| FairPlay DRM | — | — | — | — | — | — | — | — | ✅ |
| FreeType2 Fonts | — | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Games | — | — | — | — | — | ✅ | ✅ | — | ✅ |
| Genius | — | — | — | — | — | — | — | — | ✅ |
| MeCCA (Codec Framework) | — | — | — | — | — | ✅ | ✅ | ✅ | ✅ |
| Nike+ iPod | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Pedometer | — | — | — | — | — | — | — | ✅ | — |
| Photos | — | — | — | — | ✅ | ✅ | ✅ | ✅ | ✅ |
| Pixo OS | ✅ | — | — | — | — | — | — | — | ✅ |
| PortalPlayer | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| RTXC (RTOS) | ✅ | ✅ | — | — | — | — | — | — | ✅ |
| SQLite Database | — | — | — | — | — | — | — | — | ✅ |
| Silver UI Framework | — | — | — | — | — | — | — | — | ✅ |
| USB Audio | — | — | — | — | — | ✅ | ✅ | ✅ | ✅ |
| Video Playback | — | — | — | — | — | — | — | — | ✅ |
| Voice Memos | — | — | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |

## Evolution Notes

### Architecture Progression

| Generation | SoC | RTOS | Display | Storage |
|-----------|-----|------|---------|--------|
| 1st Gen (2001) | PortalPlayer PP5002 | RTXC | Mono LCD | 5/10GB HDD |
| 3rd Gen (2003) | PortalPlayer PP5002 | RTXC | Mono LCD | 10-40GB HDD |
| Mini (2004-2005) | PortalPlayer PP5020 | Pixo | Mono LCD | 4/6GB Microdrive |
| 4th Gen Mono (2004) | PortalPlayer PP5020 | Pixo | Mono LCD | 20/40GB HDD |
| 4th Gen Color (2004) | PortalPlayer PP5021 | Pixo | Color LCD | 20/40/60GB HDD |
| 5th Gen Video (2005) | PortalPlayer PP5021C | Pixo | Color LCD | 30/60GB HDD |
| 5.5G Enhanced (2006) | PortalPlayer PP5022C | Pixo | Color LCD | 30/80GB HDD |
| Nano 1G (2005) | PortalPlayer PP5021C | Pixo | Color LCD | 1/2/4GB Flash |
| Classic 7G (2009) | Samsung S5L8702 | RTXC | Color LCD | 120/160GB HDD |

### Key Milestones

- **Color Display:** Introduced with iPod Photo/4th Gen Color (2004)
- **Video Playback:** 5th Gen Video (2005) — first with MeCCA codec framework
- **Games:** 5th Gen Video (2005) — ARM binaries with PKCS#7 signing
- **Hardware AES Encryption:** Classic 6G/7G (2007+) — S5L8702 SoC
- **USB Audio:** 5th Gen Video onwards
- **Nike+ Integration:** Present in all analyzed firmware
- **FairPlay DRM:** Classic 7G only (integrated into RTXC OS)
- **Cover Flow:** Classic 7G only
- **Genius:** Classic 7G only
- **FM Radio:** Classic 7G only (hardware accessory)

