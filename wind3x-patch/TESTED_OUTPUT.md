# Test Evidence — iPod Classic on wInd3x

## Device Identification

```
$ usbipd list
BUSID  VID:PID    DEVICE                                        STATE
2-1    05ac:1223  Apple Mobile Device USB Device                Shared
```

## haxdfu Output

```
$ sudo ./wInd3x haxdfu -v
2026/07/23 13:39:02 INFO Generating payload...
2026/07/23 13:39:02 INFO Running rce....
2026/07/23 13:39:02 INFO Haxed DFU running!
```

## decrypt Output (sample — in progress)

```
$ sudo ./wInd3x decrypt /tmp/mse_out/osos ./osos_decrypted.bin -v -r /tmp/recov.dat
2026/07/23 13:59:34 INFO Parsed image. kind="Nano 3G"
2026/07/23 13:59:34 INFO Decrypting ... len=10597864
2026/07/23 13:59:34 INFO Using recovery buffer... path=/tmp/recov.dat
2026/07/23 13:59:34 INFO Decrypting ix=15360 percent=0.14
2026/07/23 13:59:39 INFO Decrypting ix=22272 percent=0.21
[...]
```

Decryption proceeds at approximately 48 bytes per USB transaction. Full 10.1MB OSOS decrypts in approximately **2 hours** via USB/IP (WSL2) or slightly less on native Linux.

## mse extract Output

```
$ ./wInd3x mse extract Firmware-35.9.0.4 -o /tmp/mse_out/
2026/07/23 11:57:17 Parsing MSE for (guessed) generation: Nano 3G
2026/07/23 11:57:17 File 0: rsrc, offset 6000, len 4e00000, prefix: false
2026/07/23 11:57:17 File 1: osos, offset 4e07000, len a1ba53, prefix: false
2026/07/23 11:57:18 File 2: aupd, offset 5824000, len 11c8b3, prefix: false
2026/07/23 11:57:18 File 3: hash, offset 5942000, len 1000, prefix: false
2026/07/23 11:57:18 INFO Extracting ... path=/tmp/mse_out/rsrc
2026/07/23 11:57:18 INFO Extracting ... path=/tmp/mse_out/osos
2026/07/23 11:57:18 INFO Extracting ... path=/tmp/mse_out/aupd
2026/07/23 11:57:18 INFO Extracting ... path=/tmp/mse_out/hash
```

## Second haxdfu Attempt (device already exploited)

```
$ sudo ./wInd3x haxdfu -v
2026/07/23 13:57:29 INFO Device already running haxed DFU
```

This confirms the exploit persists across multiple wInd3x invocations until the device is power-cycled.

## Decrypted Binary Verification

```
$ ls -la osos_decrypted.bin
-rw-r--r-- 1 root root 10599920 Jul 23 14:25 osos_decrypted.bin

$ strings osos_decrypted.bin | grep "Shuffle Songs"
Shuffle Songs

$ strings osos_decrypted.bin | grep "RTXCbug"
** RTXCbug -
RTXCbug>
```

## Notes

- The parser identifies the firmware as "Nano 3G" generation because iPod Classic shares the S5L8702 SoC. This is expected.
- The `handle_events: error: libusb: interrupted` messages during decryption are normal USB timeout retries — not errors.
- The `-r` flag is essential for the ~2 hour decrypt process. Without it, any USB interruption means starting over.
