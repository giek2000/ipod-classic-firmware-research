# iPod Firmware Decryption & Extraction Guide — All Models

## Overview

This guide covers extracting and decrypting iPod RetailOS (OSOS) firmware binaries for all iPod models. The process differs based on whether the firmware uses hardware AES encryption.

### Encryption Categories

| Category | Models | Method |
|----------|--------|--------|
| **Unencrypted** | iPod 1G–5G, Mini 1G/2G, Nano 1G | Extract from IPSW — no decryption needed |
| **Hardware AES** | iPod Classic 6G/7G, Nano 2G–7G | Exploit + decrypt via wInd3x hardware AES engine |

---

## Category 1: Unencrypted Models (1G–5G, Mini, Nano 1G)

These iPods store their firmware as raw ARM binaries inside the IPSW. No decryption is required.

### Supported Models

| Device | UpdaterFamilyID | SoC | IPSW Pattern |
|--------|----------------|-----|--------------|
| iPod 1st Gen | 1 | PP5002 | `iPod_1.x.x.ipsw` |
| iPod 3rd Gen | 2 | PP5002 | `iPod_2.x.x.ipsw` |
| iPod 4th Gen Mono | 4, 10 | PP5020 | `iPod_4.x.x.ipsw` / `iPod_10.x.x.ipsw` |
| iPod 4th Gen Color | 5, 11 | PP5021 | `iPod_5.x.x.ipsw` / `iPod_11.x.x.ipsw` |
| iPod 5th Gen Video | 13 | PP5021C | `iPod_13.x.x.ipsw` |
| iPod 5G Late 2006 | 20 | PP5021C | `iPod_20.x.x.ipsw` |
| iPod 5.5G Enhanced | 25 | PP5022C | `iPod_25.x.x.ipsw` |
| iPod Mini 1G | 6 | PP5020 | `iPod_6.x.x.ipsw` |
| iPod Mini 2G | 7 | PP5020 | `iPod_7.x.x.ipsw` |
| iPod Nano 1G (2GB) | 14 | PP5021C | `iPod_14.x.x.ipsw` |
| iPod Nano 1G (4GB) | 17 | PP5021C | `iPod_17.x.x.ipsw` |

### Extraction Steps

```bash
# 1. Download the IPSW from Apple or ipsw.me
# 2. The IPSW is a ZIP file — extract it
unzip iPod_13.1.3.ipsw

# 3. The extracted file is in MSE format (e.g., "Firmware-13.1.3")
#    Use wInd3x to parse and extract sections:
./wInd3x mse extract Firmware-13.1.3 -o ./extracted/

# Expected output:
# File 0: rsrc (resource filesystem)
# File 1: osos (operating system binary — this is what you want)
# File 2: aupd (firmware updater)
# File 3: hash (integrity check)

# 4. The extracted 'osos' file is the raw firmware binary — ready to analyze
mv ./extracted/osos osos_decrypted_iPod_5th_Gen_Video_13.1.3.bin
```

For these models, the `osos` section is NOT encrypted. It's a raw ARM binary you can load directly into Ghidra or IDA.

### Verification

```bash
# Check for recognizable strings
strings ./extracted/osos | head -50

# You should see iPod UI strings, function names, etc.
# If you see garbage/random bytes, the firmware IS encrypted (wrong category)
```

---

## Category 2: Hardware AES Encrypted Models (Classic, Nano 2G+)

These iPods use the SoC's hardware AES engine to encrypt the OSOS section. The encryption key is burned into the silicon and cannot be extracted — decryption must be performed ON the device itself.

### Supported Models (Confirmed)

| Device | SoC | DFU PID | Decryption Time | Status |
|--------|-----|---------|-----------------|--------|
| iPod Classic 6G/7G (all revisions) | S5L8702 | `0x1223` | ~2 hours | ✅ Confirmed |
| iPod Nano 3G | S5L8702 | `0x1231` | ~30 min | ✅ Native wInd3x support |

### Models Expected to Work (Same exploit family)

| Device | SoC | DFU PID | Notes |
|--------|-----|---------|-------|
| iPod Nano 4G | S5L8720 | `0x1225` | Native wInd3x support |
| iPod Nano 5G | S5L8730 | `0x1234` | Native wInd3x support |
| iPod Nano 6G | S5L8723 | `0x1242` | Native wInd3x support |
| iPod Nano 7G | S5L8942 | `0x1263` | Native wInd3x support |

### Prerequisites

| Requirement | Details |
|-------------|---------|
| **Target iPod** | Must be running stock Apple firmware |
| **USB cable** | 30-pin (Classic/Nano 1-6G) or Lightning (Nano 7G) |
| **Computer** | Linux (native or WSL2 on Windows) |
| **wInd3x** | Built from source (with Classic patch for Classic models) |
| **IPSW file** | Downloaded from Apple or ipsw.me |
| **Time** | ~2 hours for Classic, ~30 min for Nano |

### USB Product IDs — Complete Reference

| Device | DFU PID | WTF PID | Disk PID |
|--------|---------|---------|----------|
| iPod Classic Initial (80/160GB, 2007) | `0x1223` | `0x1241` | `0x1261` |
| iPod Classic Rev A (120GB, 2008) | `0x1223` | `0x1245` | `0x1261` |
| iPod Classic Rev B (160GB, 2009) | `0x1223` | `0x1247` | `0x1261` |
| iPod Classic Rev C (160GB, 2012) | `0x1250` | `0x1250` | `0x1261` |
| iPod Nano 2G | `0x1227` | `0x1232` | — |
| iPod Nano 3G | `0x1229` | `0x1231` | — |
| iPod Nano 4G | `0x1225` | `0x1243` | — |
| iPod Nano 5G | `0x1234` | `0x1246` | — |
| iPod Nano 6G | `0x1242` | `0x1249` | — |
| iPod Nano 7G | `0x1263` | `0x1267` | — |

Source: [TheAppleWiki USB Product IDs](https://theapplewiki.com/wiki/USB_Product_IDs)

---

### Step 1: Build wInd3x

For **iPod Nano 3G–7G**, wInd3x has native support — build from source directly:

```bash
sudo apt install golang-go libusb-1.0-0-dev git
git clone https://github.com/freemyipod/wInd3x.git
cd wInd3x
go build ./cmd/wInd3x
```

For **iPod Classic**, apply the device descriptor patch first:

```bash
git clone https://github.com/freemyipod/wInd3x.git
cd wInd3x
# Replace pkg/devices/devices.go with the patched version from wind3x-patch/
cp /path/to/wind3x-patch/devices.go pkg/devices/devices.go
go build ./cmd/wInd3x
```

See [`wind3x-patch/BUILD.md`](wind3x-patch/BUILD.md) for detailed build instructions.

---

### Step 2: Extract OSOS from IPSW

```bash
# Extract firmware from IPSW ZIP
unzip iPod_35.2.0.4.ipsw
# Produces: Firmware-35.9.0.4

# Parse MSE format and extract sections
./wInd3x mse extract Firmware-35.9.0.4 -o ./extracted/
```

The `extracted/osos` file is the encrypted IMG1 binary.

---

### Step 3: Enter DFU Mode

DFU mode entry varies slightly by model:

#### iPod Classic (all revisions)
1. Connect via 30-pin USB cable
2. Hold **Menu + Select (center button)** simultaneously
3. Keep holding through the reboot (~8 seconds total)
4. Release — screen stays completely blank (backlight on)
5. Device enumerates as `05AC:1223`

#### iPod Nano 3G–6G
1. Connect via 30-pin USB cable
2. Hold **Menu + Select** simultaneously
3. Keep holding through reboot (~6 seconds)
4. Release — screen goes blank
5. Device enumerates as its DFU PID (see table above)

#### iPod Nano 7G
1. Connect via Lightning cable
2. Hold **Sleep/Wake + Volume Down** simultaneously
3. Keep holding ~8 seconds
4. Release — screen blank
5. Device enumerates as `05AC:1263`

#### Verification

```bash
# Linux:
lsusb | grep Apple
# Expected: Bus XXX Device XXX: ID 05ac:XXXX Apple, Inc.

# The PID should match the DFU PID for your model
```

**Important:** The iPod must be running stock Apple firmware. If Rockbox or EmCORE is installed, restore via iTunes first.

---

### Step 4: Run the Exploit

```bash
sudo ./wInd3x haxdfu
```

Expected output:
```
INFO Generating payload...
INFO Running rce....
INFO Haxed DFU running!
```

This exploits the BootROM to gain code execution. The exploit is:
- **Non-destructive** — no permanent changes to the device
- **Non-persistent** — a hard reset returns to normal operation

---

### Step 5: Decrypt

```bash
sudo ./wInd3x decrypt ./extracted/osos ./osos_decrypted.bin -r /tmp/recovery.dat
```

| Flag | Purpose |
|------|---------|
| `-r /tmp/recovery.dat` | Resumable decryption — if interrupted, re-run to continue |
| `-v` | Verbose — shows progress percentage |

### Timing

The tool sends encrypted blocks to the device's AES engine over USB and reads back decrypted output. Each block requires a full USB control transfer round-trip.

| Model | Firmware Size | Approximate Time |
|-------|--------------|-----------------|
| iPod Classic | ~10.1 MB | ~2 hours |
| iPod Nano 3G | ~3-5 MB | ~30-60 min |
| iPod Nano 4G-7G | ~5-30 MB | ~1-3 hours |

### During Decryption

- `handle_events: error: libusb: interrupted` messages are **normal** (USB timeout retries)
- Progress is saved continuously to the recovery file
- **Do NOT disconnect the iPod**
- If interrupted, simply re-run the same command — it resumes automatically

---

### Step 6: Verify

```bash
# Check file size matches expected OSOS size
ls -la osos_decrypted.bin

# Look for recognizable strings
strings osos_decrypted.bin | grep "Shuffle Songs"
strings osos_decrypted.bin | grep -i "menu"

# If you see iPod UI strings, decryption succeeded
```

---

## Loading in Ghidra/IDA

### PortalPlayer Models (1G–5G, Mini, Nano 1G)

```
Architecture: ARM 32-bit, Little-Endian
Processor:    ARM7TDMI / ARM926EJ-S (ARMv4T/ARMv5TEJ)
Load Address: 0x00000000
Entry Point:  0x00000000 (reset vector at start)
Language:     ARM:LE:32:v5t (Ghidra)
```

### iPod Classic (S5L8702)

```
Architecture: ARM 32-bit, Little-Endian
Processor:    ARM926EJ-S (ARMv5TEJ)
Load Address: 0x00000000
Entry Point:  0x00000800 (skip 2048-byte IMG1 header)
Language:     ARM:LE:32:v5t (Ghidra)
```

---

## WSL2 Setup (Windows)

If you don't have a native Linux machine, WSL2 works with USB passthrough:

```powershell
# Windows (Admin PowerShell):
winget install usbipd

# After putting iPod in DFU mode:
usbipd list                          # Find Apple DFU device
usbipd bind --busid <BUSID>
usbipd attach --wsl --busid <BUSID>
```

Then in WSL:
```bash
sudo ./wInd3x haxdfu
sudo ./wInd3x decrypt ./extracted/osos ./osos_decrypted.bin -r /tmp/recovery.dat
```

**Note:** USB/IP adds latency — expect slower decryption. The `-r` flag is essential for resuming.

---

## Troubleshooting

| Problem | Solution |
|---------|----------|
| "unknown device" | Apply the Classic patch to devices.go and rebuild |
| Exploit fails | Ensure DFU mode (blank screen, correct PID), use `sudo`, try USB 2.0 port |
| Decryption stalls | Check USB connection, re-run with `-r` flag |
| Wrong output size | Re-run `haxdfu` first, then `decrypt` with `-r` |
| libusb errors | Normal — tool retries automatically |
| Device reboots | Hardware watchdog — re-enter DFU, re-run `haxdfu`, resume with `-r` |

---

## References

- [freemyipod/wInd3x](https://github.com/freemyipod/wInd3x) — BootROM exploit and firmware tool
- [TheAppleWiki - DFU Mode](https://theapplewiki.com/wiki/DFU_Mode)
- [TheAppleWiki - USB Product IDs](https://theapplewiki.com/wiki/USB_Product_IDs)
- [usbipd-win](https://github.com/dorssel/usbipd-win) — USB/IP for WSL2
