# iPod Classic 7G — RTXC RTOS Internals

## RTOS Overview

| Field | Value |
|-------|-------|
| RTOS | RTXC (Real-Time eXecutive with Classes) |
| Scheduling | Preemptive, priority-based |
| Kernel Location | SRAM at 0x22000000–0x2200FFFF (64 KB) |
| Kernel API Functions | 264 thunk functions bridging OSOS to SRAM |
| Identified Tasks | 11 |
| IPC Mechanisms | Semaphores, message queues, timers |
| Devices | iPod Classic 6G/7G, iPod Nano 3G |

Analysis performed using Ghidra and Capstone disassembly.

---

## RTXC Task Table

| Task Name | Address | Size | Callers | Callees | Description |
|-----------|---------|------|---------|---------|-------------|
| HostOSTask | 0x000E8F30 | 140 B | 1 | 7 | Main OS supervisor — coordinates all tasks |
| MP3ExampleTask | 0x0014434C | 76 B | 1 | 3 | MP3 decode test (development leftover) |
| USBDeviceTask | 0x001497D8 | 132 B | 1 | 5 | USB device mode handler |
| DiskReaderTask | 0x00153B48 | 76 B | 0 | 3 | Asynchronous disk I/O reader |
| ATAWorkLoopTask | 0x00163CDC | 128 B | 1 | 4 | ATA/CE-ATA command queue processor |
| GeniusMixesTask | 0x0019C9D4 | 168 B | 1 | 5 | Genius playlist generation (background) |
| TMusicLoadingTask | 0x001B999C | 92 B | 1 | 3 | Music library loader/scanner |
| MeCCAIOTask | 0x001F5354 | 276 B | 1 | 8 | Audio codec I/O thread |
| StreamCacheTimeoutTask | 0x00228A70 | 116 B | 1 | 4 | Stream buffer timeout handler |
| StreamCacheReadTask | 0x00228C2C | 52 B | 1 | 2 | Stream buffer reader |
| FirewireTask | 0x002AD4B8 | 864 B | 0 | 4 | 30-pin dock/charging handler |

---

## Task Descriptions

### HostOSTask (0x000E8F30)

The main supervisor task — first "application level" code after kernel init:
- Hardware driver initialization (ATA, USB, Audio, LCD, FireWire)
- Filesystem mount (HDD + NOR NVRAM)
- Database open (SQLite iTunesDB)
- UI launch → Main Menu via Silver framework

### USBDeviceTask (0x001497D8)

All USB device-mode operations:
- Mass Storage Class (Disk Mode, PID 0x1261)
- iTunes sync protocol (proprietary Apple accessory protocol)
- USB Audio Class
- Coordinates with FirewireTask for charging detection

### DiskReaderTask (0x00153B48)

Asynchronous disk I/O feeding the audio pipeline:
- Reads track data from storage
- Fills StreamCache buffers
- Works with ATAWorkLoopTask for serialized disk access

### ATAWorkLoopTask (0x00163CDC)

Serializes all ATA/CE-ATA commands to the hard disk:
- Processes command queue (reads, writes, identify)
- Manages disk spin-up/spin-down for power management
- Classic-only (not present on Nano 3G which uses NAND)

### GeniusMixesTask (0x0019C9D4)

Background Genius playlist generation:
- Computes similarity metrics between tracks
- Builds Genius Mix playlists
- Runs at low priority to avoid impacting playback

### MeCCAIOTask (0x001F5354)

The codec I/O thread — heart of the audio pipeline:
- Receives decoded PCM from codec registry
- Feeds audio data to DAC via I2S interface (0x38400000)
- Manages double-buffering for gapless playback
- Handles EQ processing chain

### FirewireTask (0x002AD4B8)

Largest task by code size (864 bytes):
- 30-pin dock connector detection
- FireWire charging protocol
- Dock accessory communication
- Power negotiation

---

## Scheduling Model

RTXC uses **preemptive priority-based scheduling**:
- Higher-priority tasks preempt lower-priority tasks immediately
- Tasks of equal priority run to completion or yield (no time-slicing)
- Tasks block on semaphores, message queues, or timers
- The IRQ dispatcher at 0x00003880 handles preemption on interrupt return

---

## Kernel API (264 Thunk Functions)

The OSOS accesses kernel services through 264 thunk functions at 0x000385xx. Each thunk is a single branch instruction to the corresponding SRAM kernel entry point.

### Top 10 Most-Called Kernel Functions

| Rank | Thunk Address | SRAM Target | Calls | Purpose |
|------|---------------|-------------|-------|---------|
| 1 | 0x000385B0 | 0x22000020 | 381 | malloc (memory allocate) |
| 2 | 0x000385B8 | 0x2200027C | 269 | free (memory release) |
| 3 | 0x000385F8 | 0x22000188 | 218 | Semaphore wait (block) |
| 4 | 0x000385C8 | 0x220002D4 | 89 | Semaphore signal (wake) |
| 5 | 0x00038620 | 0x22001EDC | 40 | Send message to queue |
| 6 | 0x00038608 | 0x22003FD0 | 30 | Receive message (block) |
| 7 | 0x00038610 | 0x220042B4 | 29 | Timer start |
| 8 | 0x000386B8 | 0x22001EE8 | 29 | Task suspend |
| 9 | 0x00038788 | 0x22005018 | 27 | DMA transfer initiate |
| 10 | 0x00038600 | 0x220000D4 | 26 | Task create |

### Additional Kernel Entry Points

| Thunk Address | SRAM Target | Purpose |
|---------------|-------------|---------|
| 0x000386B0 | 0x22003C28 | Kernel utility |
| 0x000387D8 | 0x22006E88 | Kernel utility |
| 0x00038860 | 0x22007470 | Kernel utility |
| 0x000388D0 | 0x22007AEC | Kernel utility |
| 0x00038978 | 0x22007E38 | Kernel utility |
| 0x00038980 | 0x22005448 | Kernel utility |
| 0x00038988 | 0x220072CC | Kernel utility |

---

## Interrupt Handler Table

### ARM Exception Vectors (at 0x00000800)

| Vector | Address | Handler Target | Purpose |
|--------|---------|---------------|---------|
| Reset | 0x00000800 | → 0x000090C4 | Boot entry point |
| Undefined Instruction | 0x00000804 | → 0x00003C28 | Fault handler |
| SWI | 0x00000808 | → 0x00003C30 | System call dispatcher |
| Prefetch Abort | 0x0000080C | → 0x00003C38 | Code fetch fault |
| Data Abort | 0x00000810 | → 0x00003C40 | Data access fault |
| Reserved | 0x00000814 | → 0x00003C48 | Unused |
| IRQ | 0x00000818 | → 0x00003880 | Hardware interrupt dispatcher |
| FIQ | 0x0000081C | → 0x00003C50 | Fast interrupt |

### IRQ Dispatcher (0x00003880)

The main interrupt routing function:
1. Saves context (all registers to stack)
2. Reads interrupt controller status (0x3C000000)
3. Looks up handler in registered handler table
4. Calls appropriate device handler
5. Checks if task preemption is needed (higher-priority task became ready)
6. Restores context and returns (`SUBS PC, LR, #4`)

---

## RTXCbug — Built-in RTOS Debugger

The firmware contains `RTXCbug`, an interactive debugger for inspecting RTOS state:

| Offset | String | Purpose |
|--------|--------|---------|
| 0x002BFA1D | `** RTXCbug - ` | Debugger banner |
| 0x002BFA60 | `  X - Exit RTXCbug` | Exit command |
| 0x002BFA75 | `RTXCbug> ` | Command prompt |
| 0x002C0451 | `RTXCbug - RTXC Objects> ` | Object inspection mode |
| 0x003947E1 | `Re-entering RTXCbug mode` | Re-entry notification |
| 0x003F3175 | `S_RTXCBUG` | RTXCbug state identifier |

RTXCbug allows inspection of:
- Task states and priorities
- Semaphore ownership and wait queues
- Message queue contents
- Timer status
- Memory allocation state

---

## Memory Management

| Metric | Value |
|--------|-------|
| malloc calls in OSOS | 381 |
| free calls in OSOS | 269 |
| Heap location | DRAM (after static sections) |
| Managed by | RTXC kernel at SRAM 0x22000020 / 0x2200027C |

---

## IPC Patterns Observed

### Semaphore Usage (Most Common)
- 218 wait calls + 89 signal calls = primary synchronization mechanism
- Used for: buffer ready signals, hardware completion, task coordination

### Message Queues
- 40 send + 30 receive calls
- Used for: inter-task commands, event notifications

### Timers
- 29 timer start calls
- Used for: timeouts, periodic polling, power management delays

---

## Sources

- Firmware binary analysis (Ghidra, ARM926EJ-S)
- Cross-reference counting of thunk functions
- String analysis for RTXCbug and task names
- RTXC RTOS documentation (public domain references)
