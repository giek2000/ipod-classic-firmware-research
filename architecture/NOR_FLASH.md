# iPod Classic 7G — NOR Flash Memory Map

## Physical Characteristics

| Field | Value |
|-------|-------|
| Total Size | 1,048,576 bytes (1 MB) |
| Mapped Address | 0x20000000 |
| Technology | Parallel NOR Flash (memory-mapped) |
| Manufacturer | SST (ID: 0xBF) |
| Device | SST39VF3201B (ID: 0x272F) |
| Sector Size | 4,096 bytes (erase unit) |
| Access | Direct memory-mapped reads at 0x20000000 |
| Functions Accessing | 51 functions (1,009 total references) |

Analysis performed using Ghidra and Capstone disassembly.

---

## Complete Memory Map

| Offset | End | Size | Region | Description |
|--------|-----|------|--------|-------------|
| 0x00000 | 0x000FF | 256 B | SysCfg | Device identity & configuration |
| 0x00100 | 0x07FFF | ~31 KB | Reserved | Erased (0xFF) |
| 0x08000 | 0x21FFF | 104 KB | Boot Image 1 | IMG1 bootloader (primary) |
| 0x22000 | 0x3CFFF | 108 KB | Boot Image 2 | IMG1 bootloader (backup) |
| 0x3D000 | 0x41CFF | ~19 KB | Reserved | Erased (0xFF) |
| 0x41D00 | 0x7BFFF | ~234 KB | Extended | Bootloader data / disk mode image |
| 0x7C000 | 0x7DFFF | 8 KB | NVRAM Primary | Persistent settings |
| 0x7E000 | 0x7FFFF | 8 KB | NVRAM Backup | Copy of NVRAM |

---

## SysCfg (System Configuration) — Offset 0x000

The SysCfg region contains device identity, calibration, and boot configuration data.

### Header Format (16 bytes)

| Offset | Size | Field | Value |
|--------|------|-------|-------|
| 0x00 | 4 | Magic | `gfCS` (= `SCfg` big-endian) |
| 0x04 | 4 | Total size | 0x000000CC (204 bytes) |
| 0x08 | 4 | Max entries | 0x00000020 (32 slots) |
| 0x0C | 4 | Entry count | 9 entries used |

### Entry Format (20 bytes each)

| Offset | Size | Field |
|--------|------|-------|
| 0x00 | 4 | Tag (4-byte ASCII, big-endian stored) |
| 0x04 | 16 | Data (format depends on tag) |

### Parsed SysCfg Entries (iPod Classic 160GB, Rev C)

| # | Offset | Tag (Raw) | Tag (Decoded) | Description | Value |
|---|--------|-----------|---------------|-------------|-------|
| 0 | 0x000 | `gfCS` | `SCfg` | Header/Magic | Size=204, Max=32, Count=9 |
| 1 | 0x014 | `mNrS` | `SrNm` | Serial Number | `8K20257G9ZS` |
| 2 | 0x028 | `dIwF` | `FwId` | Firmware ID | 12 bytes (build identifier) |
| 3 | 0x03C | `dIwH` | `HwId` | Hardware ID | `rC` (Revision C) |
| 4 | 0x050 | `rVwH` | `HwVr` | Hardware Version | 8 bytes (version data) |
| 5 | 0x064 | `cdoC` | `Codc` | Audio Codec | `SB` (Cirrus Logic) |
| 6 | 0x078 | `rVwS` | `SwVr` | Software Version | `2.0.4` |
| 7 | 0x08C | `NBLM` | `MLBN` | MLB Serial Number | `CD2110P02283` |
| 8 | 0x0A0 | `doMM` | `MMod` | Machine Model | `MC293` (160GB Late 2009) |
| 9 | 0x0B4 | `ngeR` | `Regn` | Region Code | 8 bytes (region data) |

---

## Known SysCfg Tag Reference

| Tag (BE) | Tag (LE dump) | Size | Description |
|----------|---------------|------|-------------|
| `SCfg` | `gfCS` | 4 | Magic identifier (header) |
| `SrNm` | `mNrS` | 16 | Device serial number |
| `FwId` | `dIwF` | 12 | Firmware build identifier |
| `HwId` | `dIwH` | 4+ | Hardware identifier + revision |
| `HwVr` | `rVwH` | 8 | Hardware version/revision code |
| `SwVr` | `rVwS` | 16 | Software version string |
| `Codc` | `cdoC` | 4 | Audio codec chip identifier |
| `MLBN` | `NBLM` | 16 | Main Logic Board Serial Number |
| `MMod` | `doMM` | 16 | Machine Model (e.g., `MC293`) |
| `Regn` | `ngeR` | 8 | Region code |
| `BoCr` | `rCoB` | 4 | Board/case color |
| `rPID` | `DIPr` | 4 | USB Product ID (DFU mode) |
| `Batt` | `ttaB` | 8+ | Battery calibration data |
| `ICon` | `noCI` | 4 | Device icon/type identifier |
| `FmID` | `DImF` | 4 | UpdaterFamilyID (35=Rev B, 38=Rev C) |

---

## Boot Images (IMG1 Containers)

### IMG1 Header Format

| Offset | Size | Field |
|--------|------|-------|
| 0x000 | 4 | Magic: `8702` (S5L8702 identifier) |
| 0x004 | 4 | Version: `1.0\x02` |
| 0x008 | 4 | Format/flags |
| 0x00C | 4 | Encrypted body length |
| 0x010 | 4 | Plaintext length (post-decryption) |
| 0x014 | 4 | CRC/checksum |
| 0x018 | 4 | Load address |
| ... | ... | Header extends to 0x800 |
| 0x800 | N | Encrypted ARM code body |

### Boot Images Found in NOR Dump

| # | NOR Offset | Magic | Version | Body Length | Purpose |
|---|-----------|-------|---------|-------------|---------|
| 1 | 0x08000 | `8702` | `1.0` | 102,336 B | Primary bootloader |
| 2 | 0x22000 | `8702` | `1.0` | 129,024 B | Backup bootloader |

---

## NVRAM (Persistent Settings)

Located at NOR offset 0x7C000 (primary) and 0x7E000 (backup copy).

| Property | Value |
|----------|-------|
| Primary Offset | 0x7C000 |
| Backup Offset | 0x7E000 |
| Size | 8 KB each |
| Purpose | EQ settings, volume, backlight, language, etc. |

The OSOS reads NVRAM for user settings on boot. EQ processing parameters, volume limit, and other persistent state are stored here.

---

## NOR Flash Hardware Interface

| Property | Value |
|----------|-------|
| Base Address | 0x20000000 |
| Access Control Register | 0x70000030 |
| Ready/Busy Bit | Bit 27 of 0x70000030 |
| Write Enable Bit | Bit 30 of 0x70000030 |
| Cache Control | 0xF000F040 |

### NOR Flash Chip Identification

| Field | Value |
|-------|-------|
| Manufacturer ID | 0xBF (SST) |
| Device ID | 0x272F (SST39VF3201B) |
| Interface | CFI-compliant parallel NOR |
| Word Size | 16-bit |

---

## Usage in Firmware

### Access Patterns Observed

| Access Type | Count | Context |
|------------|-------|---------|
| SysCfg reads | Frequent | Device identity at boot |
| NVRAM reads | On boot | Load persistent settings |
| NVRAM writes | On settings change | Save user preferences |
| Boot image access | Once per boot | Bootloader execution |

### Functions Referencing NOR (0x20000000)

- 51 distinct functions access addresses in the 0x20000000–0x2007FFFF range
- 1,009 total references to NOR addresses observed in the OSOS binary
- NOR is used exclusively for **read** operations from the OSOS (bootloader handles writes)

---

## Sources

- NOR flash dump analysis via hex editor
- Firmware cross-reference analysis (Ghidra)
- SST39VF3201B datasheet for CFI protocol reference
- Rockbox NOR documentation
