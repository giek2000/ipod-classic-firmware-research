# iPod 5.5G Video Enhanced - RetailOS 1.2.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.2.3 |
| **IPSW** | iPod_25.1.2.3.ipsw |
| **Device** | iPod 5.5G Video Enhanced (2006, 30/80GB HDD, 320x240 2.5in color, Video Playback) |
| **UpdaterFamilyID** | 25 |
| **Binary Size** | 13,893,632 bytes (13.25 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 13,893,632 bytes |
| **Total Strings (>=4)** | 76,584 |
| **Function Prologues** | 23,090 (ARM: 12,745, Thumb: 10,345) |
| **DRAM References** | 77,952 |
| **Peripheral Refs** | 8,934 |
| **Build** | Unknown |
| **SoC** | PortalPlayer PP5022C |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Codename** | P112 |
| **DFU PID** | N/A |
| **SHA-256** | `13b8e5cff4f4d4771a6b9bd8f13257763b65b51e88f46347fda8694899a5dc61` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00AE5068 | `TCC_Create_Task` | Known | Controller |
| 0x00AE5078 | `TCC_Current_Task_Pointer` | Known | Controller |
| 0x00AE5091 | `TCC_Delete_HISR` | Known | Controller |
| 0x00AE50A1 | `TCC_Delete_Task` | Known | Controller |
| 0x00AE50B1 | `TCC_Relinquish` | Known | Controller |
| 0x00AE50C0 | `TCC_Reset_Task` | Known | Controller |
| 0x00AE50CF | `TCC_Resume_Service` | Known | Controller |
| 0x00AE50E2 | `TCC_Task_Sleep` | Known | Controller |
| 0x00AE50F1 | `TCC_Terminate_Task` | Known | Controller |
| 0x00AE5104 | `TCF_Task_Information` | Known | Controller |
| 0x00AE5119 | `TCS_Change_Preemption` | Known | Controller |
| 0x00AE512F | `TCS_Change_Priority` | Known | Controller |
| 0x00AE5143 | `TCT_Activate_HISR` | Known | Controller |
| 0x00AE5155 | `TCT_Control_Interrupts` | Known | Controller |
| 0x00AE516C | `TCT_Local_Control_Interrupts` | Known | Controller |
| 0x00B4C910 | `TCC_Reset_Task` | Known | Controller |
| 0x00B4C91F | `TCC_Resume_Service` | Known | Controller |
| 0x00B4C947 | `TCC_Create_Task` | Known | Controller |
| 0x00B4CAA2 | `TCC_Delete_HISR` | Known | Controller |
| 0x00B4CAB2 | `TCC_Terminate_Task` | Known | Controller |
| 0x00B4CAC5 | `TCC_Delete_Task` | Known | Controller |
| 0x00B4CAF1 | `TCT_Activate_HISR` | Known | Controller |
| 0x00B4CB1B | `TCT_Control_Interrupts` | Known | Controller |
| 0x00B4CCE5 | `TCC_Current_Task_Pointer` | Known | Controller |
| 0x00B4CCFE | `TCS_Change_Priority` | Known | Controller |
| 0x00B4CD6C | `TCC_Task_Sleep` | Known | Controller |
| 0x00B4CEA5 | `TCF_Task_Information` | Known | Controller |
| 0x00B5B40D | `TCC_Create_Task` | Known | Controller |
| 0x00B5B41D | `TCC_Relinquish` | Known | Controller |
| 0x00B5B457 | `TCC_Terminate_Task` | Known | Controller |
| 0x00B5B46A | `TCC_Delete_Task` | Known | Controller |
| 0x00B5B533 | `TCT_Local_Control_Interrupts` | Known | Controller |
| 0x00B5B5D3 | `TCC_Task_Sleep` | Known | Controller |
| 0x00B5B6A1 | `TCS_Change_Preemption` | Known | Controller |

---

## 2. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0020EA48 | `Channel UnitTests` | Hidden | Developer Tool |

---

## 3. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001758D0 | `AudioCodecs` | Known | Audio system |

---

## 4. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x004B71A4 | `GotoBackToIdleCommand` | Known | Navigation |

---

## 5. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009908C | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x000C94F8 | `USB Secondary Interrupt Task` | Known | RTOS task thread |
| 0x000E20A8 | `ICAPTPCameraIOTask` | Known | RTOS task thread |
| 0x000ED1F4 | `Task` | Known | RTOS task thread |
| 0x0011C548 | `USBStatusTask` | Known | RTOS task thread |
| 0x0011C564 | `USBTaskTimeTask` | Known | RTOS task thread |
| 0x00165034 | `VCUpdateTask` | Known | RTOS task thread |
| 0x0016AEF0 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x00170AC0 | `USBDeviceTask` | Known | RTOS task thread |
| 0x00176FD4 | `DiskReaderTask` | Known | RTOS task thread |
| 0x00186244 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x00186258 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x00199514 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x002088F8 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x00258FDC | `FirewireTask` | Known | RTOS task thread |
| 0x00258FF4 | `OptoTask` | Known | RTOS task thread |
| 0x00259004 | `SerialOptoTask` | Known | RTOS task thread |
| 0x00259018 | `BacklightTask` | Known | RTOS task thread |
| 0x0025902C | `CNATask` | Known | RTOS task thread |
| 0x0025904C | `DiskMgrTask` | Known | RTOS task thread |
| 0x0025905C | `HoldSwitchTask` | Known | RTOS task thread |
| 0x00259070 | `TopPlugTask` | Known | RTOS task thread |
| 0x00259080 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x00259094 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x002590AC | `AccessoryDetectTask` | Known | RTOS task thread |
| 0x002590D4 | `AlarmTask` | Known | RTOS task thread |
| 0x002590E4 | `WatchdogTask` | Known | RTOS task thread |
| 0x0025915C | `USBAudioTask` | Known | RTOS task thread |
| 0x002A511C | `HostOSTask` | Known | RTOS task thread |
| 0x002A5CCC | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x005033F8 | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x00503410 | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x00503428 | `VideoDaisyTask` | Known | RTOS task thread |

---

## 6. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0020E854 | `Channel Reserved` | Known | Logging channel |
| 0x0020E868 | `Channel AppBoot` | Known | Logging channel |
| 0x0020E878 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x0020E894 | `Channel PrefsWriting` | Known | Logging channel |
| 0x0020E8AC | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x0020E8CC | `Channel PlayFromDisk` | Known | Logging channel |
| 0x0020E8E4 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x0020E900 | `Channel TestLogging` | Known | Logging channel |
| 0x0020E914 | `Channel AppFileLoading` | Known | Logging channel |
| 0x0020E92C | `Channel VCardReading` | Known | Logging channel |
| 0x0020E944 | `Channel LongSongScanning` | Known | Logging channel |
| 0x0020E9B8 | `Channel VoiceRecording` | Known | Logging channel |
| 0x0020E9D0 | `Channel PhotoImporting` | Known | Logging channel |
| 0x0020E9E8 | `Channel Notes` | Known | Logging channel |
| 0x0020E9F8 | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x0020EA14 | `Channel DiskMode` | Known | Logging channel |
| 0x0020EA28 | `Channel Firewire` | Known | Logging channel |
| 0x0020EA3C | `Channel USB` | Known | Logging channel |
| 0x0020EA5C | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x0020EA74 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 7. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E25D4 | `games_RO` | Known | Game system |
| 0x000E25E0 | `gamedata_RW` | Known | Game system |
| 0x0067A90B | `iPod_Control/games_RO/` | Known | Game system |

---

## 8. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E505C | `ksidksdrmrdc` | Known | DRM system |
| 0x000E6544 | `drmsdrmisinffniscpsap@-` | Known | DRM system |
| 0x001758A4 | `AppleDRMVersion` | Known | DRM system |
| 0x00175944 | `AppleDRM` | Known | DRM system |
| 0x00176970 | `AppleVideoDRM` | Known | DRM system |
| 0x0017BA3C | `drmsmp4aesdsmp4v` | Known | DRM system |
| 0x0067ADCF | `DRMLevel` | Known | DRM system |
| 0x00CBE510 | `Sdrm` | Known | DRM system |

---

## 9. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000B90C4 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000B90DC | `iTunesDB` | Known | iTunes database |
| 0x000B9104 | `System_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x000C0314 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000EF350 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000F6CC4 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000F872C | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000F8828 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x002A4ED4 | `iTunesDB` | Known | iTunes database |
| 0x002A4EE0 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0067AD08 | `iPod_Control/iTunes/` | Known | iTunes database |

---

## 10. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00175DDC | `FireWireGUID` | Known | FireWire |
| 0x00175DEC | `FireWireVersion` | Known | FireWire |
| 0x001763D8 | `FireWire` | Known | FireWire |
| 0x002AADAF | ` FireWire nen` | Known | FireWire |
| 0x002ACDD8 | `FireWire p` | Known | FireWire |
| 0x002B1828 | `FireWire-forbindelser underst` | Known | FireWire |
| 0x002B3574 | `FireWire tilsluttet` | Known | FireWire |
| 0x002B8620 | `FireWire wird nicht unterst` | Known | FireWire |
| 0x002BA522 | `ber FireWire verbunden` | Known | FireWire |
| 0x002C079A | ` FireWire. ` | Known | FireWire |
| 0x002C3DA2 | ` FireWire` | Known | FireWire |
| 0x002C89FD | `de canciones o archivos no son posibles con FireWire :` | Known | FireWire |
| 0x002CA954 | `FireWire conectado` | Known | FireWire |
| 0x002CF474 | `FireWire-tiedonsiirtoa ei tueta. Siirt` | Known | FireWire |
| 0x002D11A0 | `FireWire liitetty` | Known | FireWire |
| 0x002D6999 | `s via FireWire : connectez l` | Known | FireWire |
| 0x002D8BE8 | `FireWire Connect` | Known | FireWire |
| 0x002DD910 | `A FireWire kapcsolat nem t` | Known | FireWire |
| 0x002DFB44 | `FireWire csatlakozik` | Known | FireWire |
| 0x002E46A8 | `Connessioni di dati via FireWire non sono supportate. Per trasferire brani o dat` | Known | FireWire |
| 0x002E6534 | `FireWire connesso` | Known | FireWire |
| 0x002EBBE8 | `FireWire ` | Known | FireWire |
| 0x002EE340 | `FireWire ` | Known | FireWire |
| 0x002F34DC | `FireWire ` | Known | FireWire |
| 0x002F55DC | `FireWire ` | Known | FireWire |
| 0x002FA69E | `ren via FireWire, maar alleen via de meegeleverde USB-kabel.` | Known | FireWire |
| 0x002FC4D8 | `FireWire aangesloten` | Known | FireWire |
| 0x00300F3F | `ring via FireWire st` | Known | FireWire |
| 0x00302C08 | `Koblet til via FireWire` | Known | FireWire |
| 0x003077EB | `czenie FireWire nie jest wspierane. By przes` | Known | FireWire |
| 0x0030973B | `czone przez FireWire` | Known | FireWire |
| 0x0030E31F | `es FireWire n` | Known | FireWire |
| 0x0031024C | `FireWire ligado` | Known | FireWire |
| 0x00316315 | ` FireWire ` | Known | FireWire |
| 0x0031975B | ` FireWire` | Known | FireWire |
| 0x0031E278 | `FireWire-` | Known | FireWire |
| 0x0031FFBC | `FireWire anslutet` | Known | FireWire |
| 0x00324AA8 | `FireWire ba` | Known | FireWire |
| 0x003269EC | `FireWire Ba` | Known | FireWire |
| 0x0032B819 | ` FireWire ` | Known | FireWire |
| 0x0032D47C | `FireWire ` | Known | FireWire |
| 0x003322C9 | ` FireWire ` | Known | FireWire |
| 0x00334024 | `FireWire ` | Known | FireWire |
| 0x004B5380 | `FireWire connections are not supported. To transfer songs, connect the USB cable` | Known | FireWire |
| 0x004B7050 | `FireWire Connected` | Known | FireWire |
| 0x006CA5BD | `USBCompositeDevice1.6` | Known | USB |
| 0x006CA615 | `USBCompositeDevice1.6` | Known | USB |

---

## 11. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002B710C | `Radio-Region` | Known | FM Radio |
| 0x004B3A0C | `Radio Region` | Known | FM Radio |
| 0x004B6418 | `Radio Region` | Known | FM Radio |

---

## 12. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x006CA8A9 | `ICAType4CameraDriver` | Known | Camera |
| 0x006CA901 | `PTPCameraDriver` | Known | Camera |
| 0x00AD06D0 | `camera_control` | Known | Camera |

---

## 13. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000985E8 | `iPod_Control` | Filesystem Path |  |
| 0x00098614 | `iPod_Control\Device` | Filesystem Path |  |
| 0x000A4B1C | `iPod_Control\Device` | Filesystem Path |  |
| 0x000A6508 | `iPod_Control` | Filesystem Path |  |
| 0x000A6B60 | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x000BBBF0 | `iPod_Control\Music\` | Filesystem Path |  |
| 0x000C0194 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x000E2600 | `iPod_Control/%s/%s%s%s` | Filesystem Path |  |
| 0x001A66CC | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001A7060 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path |  |
| 0x001A7088 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001A71F4 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001D032C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D05B8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D0670 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D07C0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D08E0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D09B0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D0B48 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D0C04 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D0CB4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D0DA8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D0E4C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D0F00 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D0FBC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D10F0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1260 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1324 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D13D4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1510 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D15DC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D16A8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1770 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1814 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D18DC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D198C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1A3C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1B04 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1BC4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1C74 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1D24 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1DD4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1E84 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1F58 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2058 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2138 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2240 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D232C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0023E7FC | `iPod_Control/Device` | Filesystem Path |  |
| 0x0023E810 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x0023ED04 | `Resources/Fonts` | Filesystem Path |  |
| 0x002A5568 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x002A5D1A | `iPod_Control/Device` | Filesystem Path |  |
| 0x002A5D30 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x002A61F2 | `iPod_Control/Device` | Filesystem Path |  |
| 0x0067A9C1 | `iPod_Control/Device` | Filesystem Path |  |
| 0x0067A9D5 | `iPod_Control/Device/accessories` | Filesystem Path |  |
| 0x0067AE37 | `/Resources/VideoCore` | Filesystem Path |  |

---

## 14. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002AA1FC | `Acoustic` | EQ Preset |  |
| 0x002AA208 | `Bass Booster` | EQ Preset |  |
| 0x002AA228 | `Classical` | EQ Preset |  |
| 0x002AA234 | `Dance` | EQ Preset |  |
| 0x002AA244 | `Electronic` | EQ Preset |  |
| 0x002AA258 | `Hip Hop` | EQ Preset |  |
| 0x002AA260 | `Jazz` | EQ Preset |  |
| 0x002AA268 | `Latin` | EQ Preset |  |
| 0x002AA270 | `Loudness` | EQ Preset |  |
| 0x002AA27C | `Lounge` | EQ Preset |  |
| 0x002AA284 | `Piano` | EQ Preset |  |
| 0x002AA298 | `Rock` | EQ Preset |  |
| 0x002AA2A0 | `Small Speakers` | EQ Preset |  |
| 0x002AA2B0 | `Spoken Word` | EQ Preset |  |
| 0x002AA2BC | `Treble Booster` | EQ Preset |  |
| 0x002AA2DC | `Vocal Booster` | EQ Preset |  |
| 0x002B0CCC | `Acoustic` | EQ Preset |  |
| 0x002B0CD8 | `Bass Booster` | EQ Preset |  |
| 0x002B0CF8 | `Classical` | EQ Preset |  |
| 0x002B0D04 | `Dance` | EQ Preset |  |
| 0x002B0D14 | `Electronic` | EQ Preset |  |
| 0x002B0D28 | `Hip Hop` | EQ Preset |  |
| 0x002B0D30 | `Jazz` | EQ Preset |  |
| 0x002B0D38 | `Latin` | EQ Preset |  |
| 0x002B0D40 | `Loudness` | EQ Preset |  |
| 0x002B0D4C | `Lounge` | EQ Preset |  |
| 0x002B0D54 | `Piano` | EQ Preset |  |
| 0x002B0D68 | `Rock` | EQ Preset |  |
| 0x002B0D70 | `Small Speakers` | EQ Preset |  |
| 0x002B0D80 | `Spoken Word` | EQ Preset |  |
| 0x002B0D8C | `Treble Booster` | EQ Preset |  |
| 0x002B0DAC | `Vocal Booster` | EQ Preset |  |
| 0x002B7AC0 | `Acoustic` | EQ Preset |  |
| 0x002B7AF0 | `Dance` | EQ Preset |  |
| 0x002B7B00 | `Electronic` | EQ Preset |  |
| 0x002B7B14 | `Hip Hop` | EQ Preset |  |
| 0x002B7B1C | `Jazz` | EQ Preset |  |
| 0x002B7B24 | `Latin` | EQ Preset |  |
| 0x002B7B2C | `Loudness` | EQ Preset |  |
| 0x002B7B40 | `Piano` | EQ Preset |  |
| 0x002B7B54 | `Rock` | EQ Preset |  |
| 0x002BF398 | `Hip Hop` | EQ Preset |  |
| 0x002BF3A0 | `Jazz` | EQ Preset |  |
| 0x002BF3A8 | `Latin` | EQ Preset |  |
| 0x002BF3B0 | `Loudness` | EQ Preset |  |
| 0x002BF3BC | `Lounge` | EQ Preset |  |
| 0x002BF3C4 | `Piano` | EQ Preset |  |
| 0x002BF3D8 | `Rock` | EQ Preset |  |
| 0x002C7EB8 | `Dance` | EQ Preset |  |
| 0x002C7EE0 | `Hip Hop` | EQ Preset |  |
| 0x002C7EE8 | `Jazz` | EQ Preset |  |
| 0x002C7EF8 | `Loudness` | EQ Preset |  |
| 0x002C7F04 | `Lounge` | EQ Preset |  |
| 0x002C7F0C | `Piano` | EQ Preset |  |
| 0x002C7F20 | `Rock` | EQ Preset |  |
| 0x002CE950 | `Jazz` | EQ Preset |  |
| 0x002CE958 | `Latin` | EQ Preset |  |
| 0x002CE96C | `Lounge` | EQ Preset |  |
| 0x002CE974 | `Piano` | EQ Preset |  |
| 0x002CE988 | `Rock` | EQ Preset |  |
| 0x002D5DE8 | `Hip Hop` | EQ Preset |  |
| 0x002D5DF0 | `Jazz` | EQ Preset |  |
| 0x002D5E18 | `Lounge` | EQ Preset |  |
| 0x002D5E20 | `Piano` | EQ Preset |  |
| 0x002D5E38 | `Rock` | EQ Preset |  |
| 0x002DCD68 | `Dance` | EQ Preset |  |
| 0x002DCD98 | `Jazz` | EQ Preset |  |
| 0x002DCDA0 | `Latin` | EQ Preset |  |
| 0x002DCDD0 | `Rock` | EQ Preset |  |
| 0x002E3B4C | `Dance` | EQ Preset |  |
| 0x002E3B70 | `Hip Hop` | EQ Preset |  |
| 0x002E3B78 | `Jazz` | EQ Preset |  |
| 0x002E3B88 | `Loudness` | EQ Preset |  |
| 0x002E3B94 | `Lounge` | EQ Preset |  |
| 0x002E3B9C | `Piano` | EQ Preset |  |
| 0x002E3BB0 | `Rock` | EQ Preset |  |
| 0x002EAA70 | `Acoustic` | EQ Preset |  |
| 0x002EAA7C | `Bass Booster` | EQ Preset |  |
| 0x002EAA9C | `Classical` | EQ Preset |  |
| 0x002EAAA8 | `Dance` | EQ Preset |  |
| 0x002EAAB8 | `Electronic` | EQ Preset |  |
| 0x002EAACC | `Hip Hop` | EQ Preset |  |
| 0x002EAAD4 | `Jazz` | EQ Preset |  |
| 0x002EAADC | `Latin` | EQ Preset |  |
| 0x002EAAE4 | `Loudness` | EQ Preset |  |
| 0x002EAAF0 | `Lounge` | EQ Preset |  |
| 0x002EAAF8 | `Piano` | EQ Preset |  |
| 0x002EAB0C | `Rock` | EQ Preset |  |
| 0x002EAB14 | `Small Speakers` | EQ Preset |  |
| 0x002EAB24 | `Spoken Word` | EQ Preset |  |
| 0x002EAB30 | `Treble Booster` | EQ Preset |  |
| 0x002EAB50 | `Vocal Booster` | EQ Preset |  |
| 0x002F25F8 | `Acoustic` | EQ Preset |  |
| 0x002F2604 | `Bass Booster` | EQ Preset |  |
| 0x002F2624 | `Classical` | EQ Preset |  |
| 0x002F2630 | `Dance` | EQ Preset |  |
| 0x002F2640 | `Electronic` | EQ Preset |  |
| 0x002F2654 | `Hip Hop` | EQ Preset |  |
| 0x002F265C | `Jazz` | EQ Preset |  |
| 0x002F2664 | `Latin` | EQ Preset |  |
| 0x002F266C | `Loudness` | EQ Preset |  |
| 0x002F2678 | `Lounge` | EQ Preset |  |
| 0x002F2680 | `Piano` | EQ Preset |  |
| 0x002F2694 | `Rock` | EQ Preset |  |
| 0x002F269C | `Small Speakers` | EQ Preset |  |
| 0x002F26AC | `Spoken Word` | EQ Preset |  |
| 0x002F26B8 | `Treble Booster` | EQ Preset |  |
| 0x002F26D8 | `Vocal Booster` | EQ Preset |  |
| 0x002F9B44 | `Dance` | EQ Preset |  |
| 0x002F9B78 | `Jazz` | EQ Preset |  |
| 0x002F9B80 | `Latin` | EQ Preset |  |
| 0x002F9B88 | `Loudness` | EQ Preset |  |
| 0x002F9B94 | `Lounge` | EQ Preset |  |
| 0x002F9B9C | `Piano` | EQ Preset |  |
| 0x002F9BB0 | `Rock` | EQ Preset |  |
| 0x00300404 | `Dance` | EQ Preset |  |
| 0x00300430 | `Jazz` | EQ Preset |  |
| 0x00300440 | `Loudness` | EQ Preset |  |
| 0x0030044C | `Lounge` | EQ Preset |  |
| 0x00300454 | `Piano` | EQ Preset |  |
| 0x00300468 | `Rock` | EQ Preset |  |
| 0x00306C78 | `Hip Hop` | EQ Preset |  |
| 0x00306C80 | `Jazz` | EQ Preset |  |
| 0x00306CAC | `Lounge` | EQ Preset |  |
| 0x00306CB4 | `Piano` | EQ Preset |  |
| 0x00306CC8 | `Rock` | EQ Preset |  |
| 0x0030D7B8 | `Dance` | EQ Preset |  |
| 0x0030D7E0 | `Hip Hop` | EQ Preset |  |
| 0x0030D7E8 | `Jazz` | EQ Preset |  |
| 0x0030D7F8 | `Loudness` | EQ Preset |  |
| 0x0030D804 | `Lounge` | EQ Preset |  |
| 0x0030D80C | `Piano` | EQ Preset |  |
| 0x0030D820 | `Rock` | EQ Preset |  |
| 0x0031D730 | `Acoustic` | EQ Preset |  |
| 0x0031D73C | `Bass Booster` | EQ Preset |  |
| 0x0031D75C | `Classical` | EQ Preset |  |
| 0x0031D768 | `Dance` | EQ Preset |  |
| 0x0031D778 | `Electronic` | EQ Preset |  |
| 0x0031D78C | `Hip Hop` | EQ Preset |  |
| 0x0031D794 | `Jazz` | EQ Preset |  |
| 0x0031D79C | `Latin` | EQ Preset |  |
| 0x0031D7A4 | `Loudness` | EQ Preset |  |
| 0x0031D7B0 | `Lounge` | EQ Preset |  |
| 0x0031D7B8 | `Piano` | EQ Preset |  |
| 0x0031D7CC | `Rock` | EQ Preset |  |
| 0x0031D7D4 | `Small Speakers` | EQ Preset |  |
| 0x0031D7E4 | `Spoken Word` | EQ Preset |  |
| 0x0031D7F0 | `Treble Booster` | EQ Preset |  |
| 0x0031D810 | `Vocal Booster` | EQ Preset |  |
| 0x00323FCC | `Hip Hop` | EQ Preset |  |
| 0x00323FD8 | `Latin` | EQ Preset |  |
| 0x00323FE0 | `Loudness` | EQ Preset |  |
| 0x00323FEC | `Lounge` | EQ Preset |  |
| 0x00324008 | `Rock` | EQ Preset |  |
| 0x0032AA68 | `Acoustic` | EQ Preset |  |
| 0x0032AA74 | `Bass Booster` | EQ Preset |  |
| 0x0032AA94 | `Classical` | EQ Preset |  |
| 0x0032AAA0 | `Dance` | EQ Preset |  |
| 0x0032AAB0 | `Electronic` | EQ Preset |  |
| 0x0032AAC4 | `Hip Hop` | EQ Preset |  |
| 0x0032AACC | `Jazz` | EQ Preset |  |
| 0x0032AAD4 | `Latin` | EQ Preset |  |
| 0x0032AADC | `Loudness` | EQ Preset |  |
| 0x0032AAE8 | `Lounge` | EQ Preset |  |
| 0x0032AAF0 | `Piano` | EQ Preset |  |
| 0x0032AB04 | `Rock` | EQ Preset |  |
| 0x0032AB0C | `Small Speakers` | EQ Preset |  |
| 0x0032AB1C | `Spoken Word` | EQ Preset |  |
| 0x0032AB28 | `Treble Booster` | EQ Preset |  |
| 0x0032AB48 | `Vocal Booster` | EQ Preset |  |
| 0x00331538 | `Acoustic` | EQ Preset |  |
| 0x00331544 | `Bass Booster` | EQ Preset |  |
| 0x00331564 | `Classical` | EQ Preset |  |
| 0x00331570 | `Dance` | EQ Preset |  |
| 0x00331580 | `Electronic` | EQ Preset |  |
| 0x00331594 | `Hip Hop` | EQ Preset |  |
| 0x0033159C | `Jazz` | EQ Preset |  |
| 0x003315A4 | `Latin` | EQ Preset |  |
| 0x003315AC | `Loudness` | EQ Preset |  |
| 0x003315B8 | `Lounge` | EQ Preset |  |
| 0x003315C0 | `Piano` | EQ Preset |  |
| 0x003315D4 | `Rock` | EQ Preset |  |
| 0x003315DC | `Small Speakers` | EQ Preset |  |
| 0x003315EC | `Spoken Word` | EQ Preset |  |
| 0x003315F8 | `Treble Booster` | EQ Preset |  |
| 0x00331618 | `Vocal Booster` | EQ Preset |  |
| 0x004B4494 | `Acoustic` | EQ Preset |  |
| 0x004B44A0 | `Bass Booster` | EQ Preset |  |
| 0x004B44C0 | `Classical` | EQ Preset |  |
| 0x004B44CC | `Dance` | EQ Preset |  |
| 0x004B44DC | `Electronic` | EQ Preset |  |
| 0x004B44F0 | `Hip Hop` | EQ Preset |  |
| 0x004B44F8 | `Jazz` | EQ Preset |  |
| 0x004B4500 | `Latin` | EQ Preset |  |
| 0x004B4508 | `Loudness` | EQ Preset |  |
| 0x004B4514 | `Lounge` | EQ Preset |  |
| 0x004B451C | `Piano` | EQ Preset |  |
| 0x004B4530 | `Rock` | EQ Preset |  |
| 0x004B4538 | `Small Speakers` | EQ Preset |  |
| 0x004B4548 | `Spoken Word` | EQ Preset |  |
| 0x004B4554 | `Treble Booster` | EQ Preset |  |
| 0x004B4574 | `Vocal Booster` | EQ Preset |  |

---
