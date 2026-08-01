# iPod IPSW Container Format

## Overview

An IPSW (iPod Software) file is a **standard ZIP archive** containing Apple iPod firmware update packages. iTunes downloads these from Apple's servers to flash the iPod's firmware. All 53 IPSWs examined across 18 generations follow one of two internal formats.

Analysis performed using Ghidra and Capstone disassembly.

---

## IPSW Container Properties

| Property | Value |
|----------|-------|
| Container | Standard ZIP (PK header, DEFLATE) |
| ZIP Encryption | None — no password on any examined IPSW |
| Compression | Standard DEFLATE (~35–65% ratio) |
| Total IPSWs Cataloged | 53 |
| Generations Covered | 18 (1st Gen through Nano 7G) |

---

## Two IPSW Formats

### Format A: Legacy (PortalPlayer era + Classic/Nano 2–3)

Contains exactly **2 files**:
- `Firmware-{UpdaterFamilyID}.{version}` — Monolithic firmware image
- `manifest.plist` — XML metadata

### Format B: Samsung S5L (Nano 4–7)

Contains exactly **3 files**:
- `{Board}.bootloader.release.rb3` — Second-stage bootloader
- `Firmware.MSE` — Multi-Section Encrypted firmware image
- `manifest.plist` — XML metadata

---

## Firmware Image Internal Structure

Both Format A and Format B firmware images share the same internal layout:

### Header Structure

| Offset | Size | Content |
|--------|------|---------|
| 0x000–0x0FE | 255 B | ASCII art copyright warning (`{{~~` border) |
| 0x0FF–0x10B | 13 B | Binary header: `]ih[` magic + directory offset + version |
| 0x10C–0x3FFF | ~16 KB | Zero-filled alignment padding |
| 0x4000 | Variable | Section Directory (`!ATA` table) |

### Binary Header at 0x0FF

```
Byte 0x100: 0x5D 0x69 0x68 0x5B  →  "]ih[" (image header tag)
Byte 0x104: 0x00 0x40 0x00 0x00  →  Directory offset = 0x4000
Byte 0x108: 0x0C 0x01 0x03 0x00  →  Version 3 (all post-1st-gen)
```

---

## Section Directory (`!ATA` Table at 0x4000)

Each directory entry is 40 bytes:

| Offset | Size | Field |
|--------|------|-------|
| 0x00 | 4 | Magic: `!ATA` (0x21415441) |
| 0x04 | 4 | Section tag (4 chars, little-endian) |
| 0x08 | 4 | Reserved (zeros) |
| 0x0C | 4 | Section offset in image |
| 0x10 | 4 | Section length |
| 0x14 | 4 | Device offset (flash target) |
| 0x18 | 4 | Reserved |
| 0x1C | 4 | Checksum |
| 0x20 | 4 | Version/flags |
| 0x24 | 4 | Alignment/load address |

### Known Section Tags

| Tag (Big-Endian) | Tag (Little-Endian bytes) | Name | Purpose |
|-----------------|--------------------------|------|---------|
| `OSOS` | `soso` | Operating System | Main firmware (RetailOS) |
| `AUPD` | `dpua` | Apple Updater | Disk-mode updater firmware |
| `RSRC` | `crsr` | Resources | UI graphics, fonts, games, training data |
| `HIBE` | `ebih` | Hibernate | Sleep state resume code |
| `OSBK` | `kbso` | OS Backup | Recovery fallback OS |
| `HASH` | `hsah` | Hash | Integrity verification data |

### Section Sizes (iPod Classic 7G, FamilyID 35)

| Section | Offset | Length | Encrypted |
|---------|--------|--------|-----------|
| rsrc | 0x6000 | 78 MB | No (plaintext FAT16) |
| osos | 0x4E07000 | 10.1 MB | Yes (AES-128-CBC) |
| aupd | 0x5824000 | 1.1 MB | Yes (AES-128-CBC) |
| hash | 0x5942000 | 4 KB | No |

---

## IMG1 Container Format (OSOS Wrapper)

The OSOS section is wrapped in Apple's IMG1 format:

| Offset | Size | Field |
|--------|------|-------|
| 0x000 | 4 | Magic: `8702` (identifies S5L8702 target) |
| 0x004 | 4 | Version string: `1.0\x02` |
| 0x008 | 4 | Format/flags |
| 0x00C | 4 | Encrypted body length |
| 0x010 | 4 | Plaintext length (after decryption) |
| 0x014 | 4 | Unknown/CRC |
| 0x018 | 4 | Load address |
| ... | ... | Header extends to 0x800 |
| 0x800 | N | Encrypted ARM code body (AES-128-CBC) |

Total header: 2,048 bytes (0x800). Code begins immediately after.

---

## Encryption Status by Generation

| Generation | SoC | Firmware Encrypted? | Key |
|-----------|-----|--------------------|----|
| 1st Gen | PP5002 | No | — |
| 3rd Gen | PP5002 | No | — |
| Mini 1G/2G | PP5020/5022 | No | — |
| 4th Gen Mono/Color | PP5020 | No | — |
| 5th Gen Video | PP5021 | No | — |
| 5.5G Enhanced | PP5022C | No | — |
| Nano 1G | PP5022 | No | — |
| **Nano 2G** | **S5L8701** | **Yes** | Hardware AES (GID) |
| **Nano 3G** | **S5L8702** | **Yes** | Hardware AES (GID) |
| **Classic 6G/7G** | **S5L8702** | **Yes** | Hardware AES (GID) |
| Nano 4G | S5L8720 | Yes | Hardware AES (GID) |
| Nano 5G | S5L8730 | Yes | Hardware AES (GID) |
| Nano 6G | S5L8723 | Yes | Hardware AES (GID) |
| Nano 7G | S5L8740 | Yes | Hardware AES (GID) |

The transition to encryption occurred with the Samsung S5L SoC family (Nano 2G, 2006).

---

## Firmware Size Growth Across Generations

| Generation | Raw FW Size | IPSW Size | Compression |
|-----------|------------|-----------|-------------|
| 1st Gen (PP5002) | 4.83 MB | 2.00 MB | 41% |
| 3rd Gen (PP5002) | 4.35 MB | 1.92 MB | 44% |
| Mini (PP5020) | 4.30 MB | 2.78 MB | 65% |
| 4th Gen (PP5020) | 4.39–6.21 MB | 2.82–3.65 MB | 59–64% |
| 5G Video (PP5021) | 13.19 MB | 6.11 MB | 46% |
| Nano 1G (PP5022) | 21.85 MB | 16.88 MB | 77% |
| **Nano 3G (S5L8702)** | **89.51 MB** | **58.53 MB** | **65%** |
| **Classic 7G (S5L8702)** | **89.27–91.30 MB** | **58.29–60.57 MB** | **65%** |
| Nano 4G (S5L8720) | 87–88 MB | 57–58 MB | 66% |
| Nano 5G (S5L8730) | 117–130 MB | 75–86 MB | 64% |
| Nano 6G (S5L8723) | 150–168 MB | 102–107 MB | 64% |
| Nano 7G (S5L8740) | 167–184 MB | 105–116 MB | 63% |

Firmware grew **40×** from 1st Gen (4.8 MB) to Nano 7G (184 MB).

---

## Bootloader Files (Format B Only)

| Filename | Board ID | Device | SoC |
|----------|----------|--------|-----|
| N58s.bootloader.release.rb3 | N58s | Nano 4G | S5L8720 |
| N33.bootloader.release.rb3 | N33 | Nano 5G | S5L8730 |
| N20.bootloader.release.rb3 | N20 | Nano 6G | S5L8723 |
| N31.bootloader.rb3 | N31 | Nano 7G | S5L8740 |

These bootloaders are **not encrypted** (header starts with version string) but are signed. Size: ~130–150 KB each.

---

## UpdaterFamilyID Reference

| ID | Device | SoC |
|----|--------|-----|
| 1 | iPod 1st Gen | PP5002 |
| 2 | iPod 3rd Gen | PP5002 |
| 4/10 | iPod 4th Gen Mono | PP5020 |
| 5/11 | iPod 4th Gen Color | PP5020 |
| 6/3 | iPod Mini 1G | PP5020 |
| 7 | iPod Mini 2G | PP5022 |
| 13/20 | iPod 5th Gen Video | PP5021 |
| 14/17 | iPod Nano 1G | PP5022 |
| 25 | iPod 5.5G Enhanced | PP5022C |
| 19/29 | iPod Nano 2G | S5L8701 |
| **26** | **iPod Nano 3G** | **S5L8702** |
| 24 | iPod Classic 6G | S5L8702 |
| 33 | iPod Classic 6.5G | S5L8702 |
| 35/38 | iPod Classic 7G | S5L8702 |
| 31 | iPod Nano 4G | S5L8720 |
| 34 | iPod Nano 5G | S5L8730 |
| 36 | iPod Nano 6G | S5L8723 |
| 37/39 | iPod Nano 7G | S5L8740 |

---

## Manifest.plist Format

### Format A (Legacy)
```xml
<key>FirmwareName</key>
<string>Firmware-35.9.0.4</string>
<key>UpdaterFamilyID</key>
<integer>35</integer>
<key>FamilyID</key>
<integer>11</integer>
```

### Format B (S5L era)
```xml
<key>FirmwareName</key>
<string>Firmware.MSE</string>
<key>BootloaderName</key>
<string>N31.bootloader.rb3</string>
<key>BuildVersion</key>
<string>37A10002</string>
<key>ProductVersion</key>
<string>1.0.1</string>
```

---

## Identical Binary Observations

| Finding | Details |
|---------|---------|
| iPod 4th Gen Mono | FamilyID 4 (`iPod_4.3.1.1`) and FamilyID 10 (`iPod_10.3.1.1`) are byte-identical |
| iPod 4th Gen Color | Color (`iPod_11.1.2.1`), Photo (`iPod_5.1.2.1`) — byte-identical |
| iPod Nano 1G | 2GB (FamilyID 14) and 4GB (FamilyID 17) — byte-identical |

Apple used different UpdaterFamilyIDs to route updates to specific hardware variants, but the firmware binaries themselves are often shared across capacity variants of the same model.

---

## Sources

- 53 IPSW files downloaded from Apple/ipsw.me
- wInd3x `mse extract` for MSE container parsing
- Binary header analysis via hex editor and Ghidra
