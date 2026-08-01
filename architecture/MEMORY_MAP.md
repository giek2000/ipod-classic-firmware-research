# S5L8702 Physical Memory Map & Peripheral Registers

## Overview

The Samsung S5L8702 SoC (ARM926EJ-S) used in the iPod Classic 6G/7G and iPod Nano 3G has the following physical memory layout. All data derived from firmware binary analysis using Ghidra and Capstone.

---

## Physical Memory Map

| Address Range | Size | Description |
|--------------|------|-------------|
| 0x00000000–0x0000FFFF | 64 KB | BootROM (mirrored from 0x20000000) |
| 0x08000000–0x0BFFFFFF | 64 MB | DRAM (main working memory) |
| 0x20000000–0x200FFFFF | 1 MB | NOR Flash (bootloader, SysCfg, NVRAM) |
| 0x22000000–0x2200FFFF | 64 KB | SRAM (RTXC kernel, DFU buffer) |
| 0x38000000–0x3FFFFFFF | 128 MB | Hardware Peripheral Registers (MMIO) |

---

## Peripheral Register Base Addresses

| Base Address | Peripheral | References in FW | Unique Addrs | Notes |
|-------------|-----------|-----------------|--------------|-------|
| 0x38000000 | Clock/Power (CLKCON) | 1,313 | 647+ | Clock gating, PLL, power management |
| 0x38200000 | LCD Controller | — | — | Display output, framebuffer DMA |
| 0x38400000 | Audio (I2S/Codec) | — | — | I2S interface to Cirrus Logic DAC |
| 0x38600000 | DMA Controller | — | — | DMA for audio, LCD, disk |
| 0x38700000 | SPI Controller | — | — | SPI bus (accessories) |
| 0x38800000 | NAND Controller | — | — | NAND flash interface (Nano 3G) |
| 0x39000000 | GPIO Controller | 1,155 | 556+ | Buttons, clickwheel, hold switch, dock |
| 0x3A000000 | Unknown Peripheral A | 2,815 | 820+ | Most-referenced — likely display/memory ctrl |
| 0x3C000000 | Timer/Interrupt Ctrl | 871 | 465+ | Hardware timers, IRQ routing |
| 0x3C400000 | I2C Controller | — | — | Audio codec config, PMU |
| 0x3C500000 | SPI Controller | — | — | NOR flash, dock accessories |
| 0x3C600000 | UART | — | — | Debug serial, dock serial |
| 0x3CC00000 | USB OTG Controller | — | — | USB device/host |
| 0x3CE00000 | ATA/CE-ATA Controller | — | — | Hard disk interface (Classic only) |
| 0x3D000000 | AES/SHA Engine | 866 | 520+ | Hardware crypto accelerator |

---

## Peripheral Reference Counts

Total hardware register references observed in the Classic 7G OSOS binary: **8,029**

| Peripheral | References | Percentage | Primary Use |
|-----------|-----------|------------|-------------|
| Unknown 0x3A (display/mem?) | 2,815 | 35.1% | Highest usage — likely framebuffer/display |
| Clock/Power CLKCON | 1,313 | 16.4% | Clock gating for power management |
| GPIO Controller | 1,155 | 14.4% | Button input, clickwheel, hold switch |
| NOR Flash | 1,009 | 12.6% | SysCfg reads, NVRAM, bootloader config |
| Timer/Interrupt | 871 | 10.8% | RTXC scheduling tick, IRQ routing |
| AES/SHA Engine | 866 | 10.8% | FairPlay DRM, firmware decrypt |

---

## NOR Flash Memory Map (0x20000000)

| Offset | Size | Content |
|--------|------|---------|
| 0x20000000 | 16 KB | Config Data (IMG1 header, image pointers) |
| 0x20004000 | 16 KB | SysCfg (device identity, calibration) |
| 0x20008000 | varies | Images Section (bootloader, WTF, Disk Mode) |
| 0x2007C000 | 8 KB | NVRAM primary (persistent settings) |
| 0x2007E000 | 8 KB | NVRAM backup copy |

**Total NOR size:** 512 KB (0x80000)

---

## SRAM Layout (64 KB at 0x22000000)

The SRAM hosts the RTXC kernel, loaded by the NOR bootloader before OSOS starts:

| Address | Purpose |
|---------|---------|
| 0x22000020 | malloc entry point |
| 0x220000D4 | Task create |
| 0x22000188 | Semaphore wait |
| 0x2200027C | free entry point |
| 0x220002D4 | Semaphore signal |
| 0x22001EDC | Send message |
| 0x22001EE8 | Task suspend |
| 0x22003C28 | Kernel utility |
| 0x22003FD0 | Receive message |
| 0x220042B4 | Timer start |
| 0x22005018 | DMA transfer |
| 0x22028220 | DFU buffer address |

---

## DRAM Layout (64 MB at 0x08000000)

The 64 MB DRAM serves as working memory for the entire system:

| Use | Approximate Size |
|-----|-----------------|
| Decrypted OSOS firmware | ~10.1 MB |
| RTXC task stacks | Variable per task |
| Heap (malloc/free) | Dynamic |
| Audio decode buffers (StreamCache) | ~1–4 MB |
| Video decode buffers | ~2 MB |
| LCD framebuffer(s) | 153,600 B × 2 (double-buffer) |
| Album art cache | Variable |
| SQLite working set | Variable |

---

## OSOS Binary Internal Layout

The RetailOS firmware after decryption:

| Offset | Size | Content |
|--------|------|---------|
| 0x00000000–0x000007FF | 2,048 B | IMG1 Header (skip when loading in analysis tools) |
| 0x00000800–0x00000820 | 32 B | ARM Vector Table (8 vectors × 4 bytes) |
| 0x00000820–0x003ED38F | ~4.1 MB | Executable Code (17,721 ARM + 5,312 Thumb functions) |
| 0x003ED390–0x003EE9C0 | ~6.2 KB | Data tables / lookup structures |
| 0x003EE9C1–0x003F2D0B | ~17 KB | Largest contiguous free space block |
| 0x003F0000–0x00720000 | ~3.3 MB | Data Tables / Pointer Arrays |
| 0x00730000–0x007E0000 | ~720 KB | UI Layout Definitions (Silver screen XMLs) |
| 0x007F0000–0x009F0000 | ~2.0 MB | String Tables (55,243 strings, 20+ languages) |
| 0x00980000–0x009E0000 | ~384 KB | C++ RTTI Data (class names, vtable pointers) |
| 0x009F0000–0x00A1BA52 | ~144 KB | Remaining data + padding |

### Key Statistics

| Metric | Value |
|--------|-------|
| Total binary size | 10,599,920 bytes (10.11 MB) |
| Total functions | 23,033 |
| BL call instructions | 86,263 |
| Unique call targets | 15,437 |
| Hardware register references | 8,029 |
| Strings | 55,243 (20+ languages) |
| Free space (available) | 934,250 bytes (8.8% of firmware) |

---

## Clock/Power Controller Details (0x38000000)

The most-accessed controller after the unknown 0x3A peripheral. Controls all clock gating and power domains:

| Register | Function |
|----------|----------|
| 0x38000000 | Base control register |
| 0x38000001 | Clock enable/disable |
| 0x38000003 | PLL configuration |
| 0x38000004 | Clock divider |
| 0x38000012 | Power domain control |
| 0x38000016 | Peripheral clock gate |
| 0x38000017 | Peripheral clock gate 2 |
| 0x38000020 | PLL lock status |

---

## GPIO Controller Details (0x39000000)

Handles all discrete I/O for user interaction:

| Function | GPIO Use |
|----------|----------|
| Clickwheel | Rotary input via GPIO interrupt |
| Hold Switch | Digital input — disables all buttons |
| Buttons | Menu, Select, Play/Pause, Forward, Back |
| Dock Connector | Pin sensing for accessory detection |
| Headphone Jack | Insertion detection |

---

## AES/SHA Engine Details (0x3D000000)

| Parameter | Value |
|-----------|-------|
| Base Address | 0x3D000000 |
| Register Range | 0x3D000000–0x3DFFFFFF |
| Functions Accessing | 17 |
| Total Register Accesses | 866 |
| Algorithm | AES-128-CBC |
| Key Sources | GID (silicon fuses) or software-provided |
| Uses | Boot decryption, FairPlay DRM, game key unwrap |

---

## Sources

- Firmware binary analysis (Ghidra, ARM926EJ-S processor)
- Peripheral reference counting via cross-reference analysis
- Rockbox S5L8702 hardware documentation
- wInd3x memory dump capabilities
