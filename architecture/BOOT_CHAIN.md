# iPod Classic / Nano 3G — Boot Chain Architecture

## Target Platform

| Field | Value |
|-------|-------|
| SoC | Samsung S5L8702 |
| CPU | ARM926EJ-S (ARMv5TEJ), ~200 MHz |
| Applies to | iPod Classic 6G/7G, iPod Nano 3G |
| RTOS | RTXC (preemptive, priority-based) |
| Analysis | Performed using Ghidra and Capstone disassembly |

---

## Full Boot Sequence

```
Power On (PMU triggers)
       │
       ▼
┌─────────────────┐
│   BootROM       │  Fixed in silicon, ~64KB at 0x20000000
│   (S5L8702)     │  Contains GID AES key (burnt into chip)
│                 │  Contains DFU mode implementation
└────────┬────────┘
         │
         │  Reads NOR flash
         │  Checks SysCfg (HwVr, FwId)
         │  Button check → DFU mode if requested
         ▼
┌─────────────────┐
│  NOR Bootloader │  IMG1 format container
│                 │  51 functions reference NOR addresses
│                 │  Contains SysCfg at NOR offset 0x4000
│                 │  Decides: Boot OSOS / DFU / Disk Mode
└────────┬────────┘
         │
         │  Reads firmware from storage
         │  Decrypts via hardware AES engine (0x3D000000)
         │  AES-128-CBC, key from silicon GID
         ▼
┌─────────────────┐
│  AES Decryption │  Hardware engine at 0x3D000000
│  (full OSOS)    │  866 total register accesses observed
│                 │  Decrypts entire OSOS IMG1 blob
└────────┬────────┘
         │
         │  Copies decrypted binary to DRAM
         │  Jumps to offset 0x800 (vector table)
         ▼
┌─────────────────┐
│  RetailOS       │  The main firmware binary
│  (OSOS)         │  Vector table at 0x800
│                 │  RTXC kernel init → HostOSTask → full OS
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  RTXC Kernel    │  264 kernel API functions at 0x22000000+
│  Tasks Running  │  11 identified tasks
│                 │  Semaphores, message queues, timers
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Application    │  Silver UI framework
│  Layer          │  MeCCA codec framework
│                 │  SQLite, FreeType2, FairPlay DRM
└─────────────────┘
```

---

## Stage 1: BootROM (Silicon — Unmodifiable)

The BootROM is fixed in hardware mask ROM and cannot be modified.

| Property | Value |
|----------|-------|
| Location | 0x20000000 (also mirrored at 0x00000000) |
| Size | ~64 KB |
| Modifiable | No — mask ROM |
| Contains | GID AES key, DFU implementation, AES decrypt function |
| AES Function Address | 0x20001F04 |

**Responsibilities:**
- Read NOR flash to locate bootloader
- Check button state (MENU+SELECT → DFU mode)
- Enumerate as USB DFU device if triggered (USB PID varies by model)
- On normal boot: load and verify NOR bootloader, jump to it

---

## Stage 2: NOR Flash Bootloader

The NOR bootloader is stored in IMG1 container format within the NOR flash.

| Property | Value |
|----------|-------|
| Format | IMG1 container (2KB header + encrypted payload) |
| IMG1 Magic | `8702` (identifies S5L8702 target) |
| Primary Location | NOR offset 0x08000 (104 KB) |
| Backup Location | NOR offset 0x22000 (108 KB) |
| Functions referencing NOR | 51 (1,009 total references) |

**Responsibilities:**
- Read SysCfg for device identity (serial, HwVr, FwId, model)
- Decide boot path: normal OSOS boot, DFU, or Disk Mode
- Read firmware image from storage (hidden partition on Classic, NAND on Nano)
- Program AES hardware engine with GID key
- Decrypt full OSOS into DRAM
- Jump to offset 0x800 within decrypted payload

---

## Stage 3: AES Decryption

| Parameter | Value |
|-----------|-------|
| Algorithm | AES-128-CBC |
| Key Source | GID (silicon fuse, never leaves chip) |
| IV | Initial (zeros) |
| Engine Address | 0x3D000000 |
| Payload Size | ~10.1 MB (Classic) / ~10.3 MB (Nano 3G) |
| OSOS register accesses | 866 to AES engine |
| Functions touching AES | 17 |

---

## Stage 4: RetailOS Execution

### Vector Table at 0x00000800

The first 2KB of the decrypted OSOS is an IMG1 header (skipped). The ARM vector table begins at offset 0x800:

| Vector | Address | Target | Purpose |
|--------|---------|--------|---------|
| Reset | 0x00000800 | → 0x000090C4 | Boot entry point |
| Undefined Instruction | 0x00000804 | → 0x00003C28 | Fault handler |
| SWI | 0x00000808 | → 0x00003C30 | System call dispatcher |
| Prefetch Abort | 0x0000080C | → 0x00003C38 | Code fetch fault |
| Data Abort | 0x00000810 | → 0x00003C40 | Data access fault |
| Reserved | 0x00000814 | → 0x00003C48 | Unused |
| IRQ | 0x00000818 | → 0x00003880 | Hardware interrupt dispatcher |
| FIQ | 0x0000081C | → 0x00003C50 | Fast interrupt |

### Reset Handler (0x000090C4)

The first code to execute after boot:
1. Initialize RTXC kernel
2. Configure clock tree (CLKCON at 0x38000000, 1,313 references)
3. Set up GPIO (0x39000000, 1,155 references)
4. Initialize timer/interrupt controller (0x3C000000, 871 references)
5. Configure memory controller and ARM9 cache
6. Clear BSS, copy initialized data
7. Initialize stacks for each CPU mode (IRQ, FIQ, SVC, USR)
8. Launch HostOSTask — the main supervisor

### IRQ Dispatcher (0x00003880)

Installed at vector 0x818, handles all hardware interrupts:
1. Save CPU context (all registers)
2. Read interrupt controller status (0x3C000000)
3. Look up handler in registered handler table
4. Call appropriate device handler
5. Check if RTXC task preemption is needed
6. Restore context and return

---

## Stage 5: RTXC Kernel and Application Layer

Once HostOSTask launches:
- Hardware drivers initialized (ATA/NAND, USB, Audio, LCD, dock)
- Filesystem mounted
- SQLite database opened (iTunesDB)
- UI launched via Silver framework → Main Menu

---

## DFU Mode

### Entry

Hold MENU + SELECT simultaneously during boot (~8 seconds):
- Screen goes completely blank (backlight on, no icons/text)
- Device enumerates as USB DFU class device
- This executes from BootROM — not from any flash

### USB Parameters

| Model | DFU PID | WTF PID | Disk Mode PID |
|-------|---------|---------|---------------|
| Classic Initial (2007) | 0x1223 | 0x1241 | 0x1261 |
| Classic Rev A (2008) | 0x1223 | 0x1245 | 0x1261 |
| Classic Rev B (2009) | 0x1223 | 0x1247 | 0x1261 |
| Classic Rev C (2012) | 0x1250 | 0x1250 | 0x1261 |
| Nano 3G | 0x1229 | 0x1231 | — |

---

## Authentication Analysis

### What IS Authenticated (Encryption = Authentication)

| Layer | Mechanism |
|-------|-----------|
| Firmware on storage | AES-128-CBC with GID key |
| Game executables | PKCS#7 code signing |
| Game manifests | XML Digital Signatures |
| iTunesDB | PKCS#7 signature |

### What Is NOT Authenticated

| Gap | Observation |
|-----|-------------|
| Post-decrypt integrity | No hash check between decryption and execution |
| OSOS modules | No per-section signing — monolithic binary |
| NOR contents | No secure boot chain fuse/verification observed |
| Anti-rollback | No fuse counting — any firmware version accepted |
| Runtime integrity | No self-verification during execution |

### Key Observation

The NOR bootloader decrypts the OSOS and jumps to it unconditionally. Encryption serves as the **sole** authentication mechanism — the assumption is that only the holder of the GID key can produce valid encrypted firmware. This design means that once an exploit provides access to the hardware AES engine, arbitrary code can be loaded.

---

## Component Address Classification

| Component | Address Range | Notes |
|-----------|--------------|-------|
| BootROM | 0x20000000–0x2000FFFF | Mask ROM, unmodifiable |
| NOR Bootloader | 0x20004000–0x2007FFFF | Writable NOR flash |
| RTXC Kernel (SRAM) | 0x22000000–0x2200FFFF | Loaded by bootloader |
| RetailOS (OSOS) | 0x00000800–0x00A1BA52 | Primary analysis target |
| Hardware Registers | 0x38000000–0x3FFFFFFF | MMIO peripherals |

---

## Sources

- Firmware analysis via Ghidra (ARM926EJ-S, ARMv5TEJ)
- wInd3x project documentation (freemyipod/wInd3x)
- DavidBuchanan314 classic-ipod-tools
- Rockbox bootloader-ipodclassic documentation
- TheAppleWiki
