# iPod Classic 7th Generation — OSOS Firmware Decryption Guide

## Overview

The iPod Classic RetailOS (OSOS) firmware is AES-encrypted in IMG1 format. Decryption requires the device's hardware AES engine — the key is burned into the SoC and cannot be extracted. This guide documents how to decrypt the firmware using the open-source [wInd3x](https://github.com/freemyipod/wInd3x) tool with an iPod Classic device descriptor patch.

## Prerequisites

| Requirement | Details |
|-------------|---------|
| **iPod Classic** | Any 6th/7th generation (2007-2014), 80/120/160GB |
| **USB cable** | 30-pin to USB-A |
| **Computer** | Linux (native or WSL2 on Windows) |
| **wInd3x** | Built from source with Classic patch applied |
| **IPSW file** | `iPod_35.2.0.4.ipsw` (download from Apple or ipsw.me) |
| **Stock firmware** | iPod must be running original Apple firmware (not Rockbox/EmCORE) |

### Time Estimate

The decryption processes approximately 48 bytes per USB transaction. For the full 10.1MB OSOS binary, expect approximately **2 hours** depending on USB bus speed and system load. The process is resumable — if interrupted, it continues from where it left off.

---

## Step 1: Build wInd3x with iPod Classic Support

wInd3x does not natively support iPod Classic. You need to patch `pkg/devices/devices.go` to add the Classic's USB identifiers.

### What to add

The iPod Classic uses the S5L8702 SoC (same as Nano 3G) and shares the same DFU protocol version. You need to:

1. Add a `Classic6G` device kind constant
2. Group it with Nano3 in the `SoCCode()` function (returns `"8702"`)
3. Group it with Nano3 in the `DFUVersion()` function (returns `DFUProtoVersion1`)
4. Add a device description entry with these USB PIDs:
   - **DFU PID:** `0x1223`
   - **WTF PID:** `0x1247`
   - **Disk PID:** `0x1261`
   - **UpdaterFamilyID:** `35`

### Build instructions

```bash
# Install dependencies
sudo apt install golang-go libusb-1.0-0-dev git

# Clone wInd3x
git clone https://github.com/freemyipod/wInd3x.git
cd wInd3x

# Edit pkg/devices/devices.go — add Classic6G device (see above)
# Then build:
go build ./cmd/wInd3x

# Verify
./wInd3x --help
```

---

## Step 2: Extract OSOS from IPSW

The IPSW is a ZIP file containing a single MSE-format firmware image. Extract the OSOS section:

```bash
# Extract the firmware binary from the IPSW ZIP
unzip iPod_35.2.0.4.ipsw
# This produces a file named "Firmware-35.9.0.4" (93.6 MB)

# Use wInd3x to parse the MSE format and extract sections
./wInd3x mse extract Firmware-35.9.0.4 -o ./extracted/
```

Expected output:
```
Parsing MSE for (guessed) generation: Nano 3G
File 0: rsrc, offset 6000, len 4e00000, prefix: false
File 1: osos, offset 4e07000, len a1ba53, prefix: false
File 2: aupd, offset 5824000, len 11c8b3, prefix: false
File 3: hash, offset 5942000, len 1000, prefix: false
```

The `extracted/osos` file (10.1 MB) is the encrypted RetailOS binary in IMG1 format.

**Note:** The parser reports "Nano 3G" because iPod Classic shares the same SoC code. This is expected and correct.

---

## Step 3: Enter DFU Mode

1. Connect iPod to computer via 30-pin USB cable
2. Hold **Menu + Select (center button)** simultaneously
3. The device resets — **keep holding both buttons through the reboot**
4. After ~8 seconds total, release both buttons
5. Screen will be completely blank (backlight on, no icons/text)
6. Device enumerates as USB VID:PID `05AC:1223`

Verify with `lsusb`:
```bash
$ lsusb | grep Apple
Bus 001 Device 005: ID 05ac:1223 Apple, Inc.
```

**Important:** The iPod must be running stock Apple firmware. If you have Rockbox or EmCORE installed, restore via iTunes first.

---

## Step 4: Run the Exploit

```bash
sudo ./wInd3x haxdfu
```

Expected output:
```
INFO Generating payload...
INFO Running rce....
INFO Haxed DFU running!
```

This exploits the S5L8702 BootROM to gain code execution in DFU mode. The exploit is non-destructive and non-persistent — a hard reset returns the device to normal.

If you see "Device already running haxed DFU", the exploit was already active from a previous run.

---

## Step 5: Decrypt the OSOS

```bash
sudo ./wInd3x decrypt ./extracted/osos ./osos_decrypted.bin -r /tmp/recovery.dat
```

| Flag | Purpose |
|------|---------|
| `-r /tmp/recovery.dat` | Enables resumable decryption. If interrupted, re-run the same command to continue from where it left off. |
| `-v` | (Optional) Verbose output showing progress percentage |

### What happens during decryption

The tool sends 48-byte blocks to the device's AES engine over USB and reads back the decrypted output. This is slow because each block requires a full USB control transfer round-trip.

### Expected behavior

- You will see periodic `handle_events: error: libusb: interrupted` messages — these are normal USB timeout retries
- The tool automatically retries on transient USB errors
- Progress is saved continuously to the recovery file
- **Do NOT disconnect the iPod during decryption**
- If the process is interrupted (USB disconnect, system sleep, etc.), simply re-run the same command — it resumes automatically

### Completion

When finished, you'll have `osos_decrypted.bin` — the full unencrypted RetailOS binary (10,599,920 bytes / ~10.1 MB).

---

## Step 6: Verify the Decrypted Binary

Quick sanity checks:

```bash
# File size should be ~10.1 MB
ls -la osos_decrypted.bin
# Expected: 10,599,920 bytes

# Check for ARM vector table at offset 0x800
hexdump -C -s 0x800 -n 32 osos_decrypted.bin

# Search for known strings
strings osos_decrypted.bin | grep "Shuffle Songs"
strings osos_decrypted.bin | grep "RTXCbug"
strings osos_decrypted.bin | grep "MeCCADecode"
```

If you see recognizable iPod UI strings and function names, decryption was successful.

---

## Using WSL2 on Windows

If you don't have a native Linux machine, WSL2 works with USB passthrough via [usbipd-win](https://github.com/dorssel/usbipd-win):

### Setup

```powershell
# Windows (Admin PowerShell):
winget install usbipd
```

### Workflow

```powershell
# 1. Put iPod in DFU mode (Menu+Select 8 sec)

# 2. List USB devices and find the iPod
usbipd list
# Look for: 05ac:1223  USB DFU Device

# 3. Share and attach to WSL
usbipd bind --busid <BUSID>
usbipd attach --wsl --busid <BUSID>
```

Then in WSL:
```bash
# Run wInd3x normally
sudo ./wInd3x haxdfu
sudo ./wInd3x decrypt ./extracted/osos ./osos_decrypted.bin -r /tmp/recovery.dat
```

### WSL Caveats

- USB/IP adds latency — decryption will be slower (closer to 2+ hours)
- You may see more timeout/retry messages — this is normal
- If the iPod disconnects from WSL, re-attach with `usbipd attach --wsl --busid <BUSID>` and re-run (the `-r` flag handles resuming)
- Make sure WSL is running BEFORE attaching the device

---

## USB Product IDs Reference

All iPod Classic hardware revisions share DFU PID `0x1223`:

| Revision | Release | WTF PID | Model Numbers |
|----------|---------|---------|---------------|
| Initial (80/160GB) | Sep 2007 | `0x1241` | MB029, MB147, MB145, MB150 |
| Rev A (120GB) | Sep 2008 | `0x1245` | MB562, MB565 |
| Rev B (160GB) | Sep 2009 | `0x1247` | MC293, MC297 |
| Rev C (160GB) | Oct 2012 | `0x1250` | MD717, MD718 |

Source: [TheAppleWiki USB Product IDs](https://theapplewiki.com/wiki/USB_Product_IDs)

---

## Troubleshooting

### "unknown device" or wInd3x doesn't recognize iPod
You haven't applied the Classic device descriptor patch to `devices.go`. Rebuild wInd3x with the patch.

### Exploit fails / no "Haxed DFU running"
- Ensure iPod is in DFU mode (blank screen, PID 0x1223)
- Ensure you're running with `sudo` (USB access requires root)
- Try a different USB port (USB 2.0 ports are more reliable than USB 3.0 hubs)

### Decryption stalls or produces wrong-size output
- Check USB connection is stable
- Re-run with `-r` flag to resume
- If output is consistently wrong, the device may not be in haxed DFU state — re-run `haxdfu` first

### "libusb: interrupted" errors
These are normal. The tool retries automatically. Only be concerned if the progress percentage stops advancing for more than 5 minutes.

### Device reboots during decryption
The iPod has a hardware watchdog. If the process takes too long between USB transactions, the device may reboot. Put it back in DFU mode, re-run `haxdfu`, then re-run `decrypt` with the `-r` flag.

---

## What You Get

The decrypted `osos_decrypted.bin` is a raw ARM binary (ARM926EJ-S, ARMv5TEJ, little-endian) that can be loaded into disassemblers like Ghidra or IDA Pro:

- **Load address:** 0x00000000
- **Entry point:** Vector table at offset 0x800 (skip IMG1 header)
- **Architecture:** ARM 32-bit, little-endian
- **Processor:** ARM926EJ-S (select ARMv5TEJ in Ghidra)

See `RETAILOS_FEATURE_SPEC.md` for a complete catalog of discovered features, functions, and hidden capabilities with their hex offsets.

---

## References

- [freemyipod/wInd3x](https://github.com/freemyipod/wInd3x) — BootROM exploit and firmware tool
- [TheAppleWiki - DFU Mode](https://theapplewiki.com/wiki/DFU_Mode)
- [TheAppleWiki - USB Product IDs](https://theapplewiki.com/wiki/USB_Product_IDs)
- [usbipd-win](https://github.com/dorssel/usbipd-win) — USB/IP for WSL2
