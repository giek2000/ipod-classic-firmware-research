# Building wInd3x with iPod Classic Support

## Building From Source

### Requirements
- Go 1.21+ (`go version`)
- libusb development headers (`sudo apt install libusb-1.0-0-dev`)
- Git

### Steps

```bash
# Clone wInd3x
git clone https://github.com/freemyipod/wInd3x.git
cd wInd3x

# Apply patch (replace devices.go)
cp /path/to/this/devices.go pkg/devices/devices.go

# Build
go build ./cmd/wInd3x

# Verify
./wInd3x --help
```

### Building on Raspberry Pi (ARM)

For Raspberry Pi (recommended by freemyipod for USB reliability):

```bash
sudo apt install golang-go libusb-1.0-0-dev git
git clone https://github.com/freemyipod/wInd3x.git
cd wInd3x
cp /path/to/devices.go pkg/devices/devices.go
go build ./cmd/wInd3x
```

The native ARM build works because CGo compiles against the local libusb.

### WSL2 Usage (Windows)

wInd3x works from WSL2 with `usbipd-win` for USB passthrough:

```powershell
# Windows (Admin PowerShell):
winget install usbipd
usbipd list                          # Find the Apple DFU device (05AC:1223)
usbipd bind --busid <BUSID>
usbipd attach --wsl --busid <BUSID>

# WSL must be running first (open a WSL terminal before attaching)
```

Then in WSL:
```bash
sudo ./wInd3x haxdfu
sudo ./wInd3x decrypt ./extracted/osos ./osos_decrypted.bin -r /tmp/recovery.dat
```

**Note:** USB/IP adds latency. The `decrypt` command will take approximately 2 hours and may report timeout errors that are automatically retried. The `-r` flag enables resumable decryption in case of disconnection.
