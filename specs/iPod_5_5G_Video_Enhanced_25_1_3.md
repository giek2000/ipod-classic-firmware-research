# iPod 5.5G Video Enhanced - RetailOS 1.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.3 |
| **IPSW** | iPod_25.1.3.ipsw |
| **Device** | iPod 5.5G Video Enhanced (2007, 30/80GB HDD, 320x240 2.5in color, Video Playback) |
| **UpdaterFamilyID** | 25 |
| **Binary Size** | 13,903,872 bytes (13.26 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 13,903,872 bytes |
| **Total Strings (>=4)** | 77,126 |
| **Function Prologues** | 22,824 (ARM: 12,890, Thumb: 9,934) |
| **DRAM References** | 77,952 |
| **Peripheral Refs** | 8,976 |
| **Build** | Unknown |
| **SoC** | PortalPlayer PP5022C |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Codename** | P112 |
| **DFU PID** | N/A |
| **SHA-256** | `7830d1345aa2313db154e06ae93f2b5961e1cb04e8edeaae27dadc303e0d9fb3` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00B28868 | `TCC_Create_Task` | Known | Controller |
| 0x00B28878 | `TCC_Current_Task_Pointer` | Known | Controller |
| 0x00B28891 | `TCC_Delete_HISR` | Known | Controller |
| 0x00B288A1 | `TCC_Delete_Task` | Known | Controller |
| 0x00B288B1 | `TCC_Relinquish` | Known | Controller |
| 0x00B288C0 | `TCC_Reset_Task` | Known | Controller |
| 0x00B288CF | `TCC_Resume_Service` | Known | Controller |
| 0x00B288E2 | `TCC_Task_Sleep` | Known | Controller |
| 0x00B288F1 | `TCC_Terminate_Task` | Known | Controller |
| 0x00B28904 | `TCF_Task_Information` | Known | Controller |
| 0x00B28919 | `TCS_Change_Preemption` | Known | Controller |
| 0x00B2892F | `TCS_Change_Priority` | Known | Controller |
| 0x00B28943 | `TCT_Activate_HISR` | Known | Controller |
| 0x00B28955 | `TCT_Control_Interrupts` | Known | Controller |
| 0x00B2896C | `TCT_Local_Control_Interrupts` | Known | Controller |
| 0x00B90110 | `TCC_Reset_Task` | Known | Controller |
| 0x00B9011F | `TCC_Resume_Service` | Known | Controller |
| 0x00B90147 | `TCC_Create_Task` | Known | Controller |
| 0x00B902A2 | `TCC_Delete_HISR` | Known | Controller |
| 0x00B902B2 | `TCC_Terminate_Task` | Known | Controller |
| 0x00B902C5 | `TCC_Delete_Task` | Known | Controller |
| 0x00B902F1 | `TCT_Activate_HISR` | Known | Controller |
| 0x00B9031B | `TCT_Control_Interrupts` | Known | Controller |
| 0x00B904E5 | `TCC_Current_Task_Pointer` | Known | Controller |
| 0x00B904FE | `TCS_Change_Priority` | Known | Controller |
| 0x00B9056C | `TCC_Task_Sleep` | Known | Controller |
| 0x00B906A5 | `TCF_Task_Information` | Known | Controller |
| 0x00B9EC0D | `TCC_Create_Task` | Known | Controller |
| 0x00B9EC1D | `TCC_Relinquish` | Known | Controller |
| 0x00B9EC57 | `TCC_Terminate_Task` | Known | Controller |
| 0x00B9EC6A | `TCC_Delete_Task` | Known | Controller |
| 0x00B9ED33 | `TCT_Local_Control_Interrupts` | Known | Controller |
| 0x00B9EDD3 | `TCC_Task_Sleep` | Known | Controller |
| 0x00B9EEA1 | `TCS_Change_Preemption` | Known | Controller |
| 0x00D2CD18 | `TCFRL` | Known | Controller |

---

## 2. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00216934 | `Channel UnitTests` | Hidden | Developer Tool |

---

## 3. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017A888 | `AudioCodecs` | Known | Audio system |

---

## 4. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x004B8C04 | `GotoBackToIdleCommand` | Known | Navigation |

---

## 5. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00099174 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x000C9B20 | `USB Secondary Interrupt Task` | Known | RTOS task thread |
| 0x000E334C | `ICAPTPCameraIOTask` | Known | RTOS task thread |
| 0x000EE528 | `Task` | Known | RTOS task thread |
| 0x0011DBA8 | `USBStatusTask` | Known | RTOS task thread |
| 0x0011DBC4 | `USBTaskTimeTask` | Known | RTOS task thread |
| 0x00169A2C | `VCUpdateTask` | Known | RTOS task thread |
| 0x0016F904 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x0017587C | `USBDeviceTask` | Known | RTOS task thread |
| 0x0017BFB4 | `DiskReaderTask` | Known | RTOS task thread |
| 0x0018B3BC | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0018B3D0 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0019E7F8 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x00210664 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x00216B28 | `Channel DiskReaderTask` | Known | RTOS task thread |
| 0x00261E3C | `FirewireTask` | Known | RTOS task thread |
| 0x00261E54 | `OptoTask` | Known | RTOS task thread |
| 0x00261E64 | `SerialOptoTask` | Known | RTOS task thread |
| 0x00261E78 | `BacklightTask` | Known | RTOS task thread |
| 0x00261E8C | `CNATask` | Known | RTOS task thread |
| 0x00261EAC | `DiskMgrTask` | Known | RTOS task thread |
| 0x00261EBC | `HoldSwitchTask` | Known | RTOS task thread |
| 0x00261ED0 | `TopPlugTask` | Known | RTOS task thread |
| 0x00261EE0 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x00261EF4 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x00261F0C | `AccessoryDetectTask` | Known | RTOS task thread |
| 0x00261F34 | `AlarmTask` | Known | RTOS task thread |
| 0x00261F44 | `WatchdogTask` | Known | RTOS task thread |
| 0x00261FBC | `USBAudioTask` | Known | RTOS task thread |
| 0x002AE410 | `HostOSTask` | Known | RTOS task thread |
| 0x002AEFFC | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x0050535C | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x00505374 | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x00505498 | `VideoDaisyTask` | Known | RTOS task thread |

---

## 6. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002166F4 | `Channel Reserved` | Known | Logging channel |
| 0x00216708 | `Channel AppBoot` | Known | Logging channel |
| 0x00216718 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x00216734 | `Channel PrefsWriting` | Known | Logging channel |
| 0x0021674C | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x0021676C | `Channel PlayFromDisk` | Known | Logging channel |
| 0x00216784 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x002167A0 | `Channel TestLogging` | Known | Logging channel |
| 0x002167B4 | `Channel AppFileLoading` | Known | Logging channel |
| 0x002167CC | `Channel VCardReading` | Known | Logging channel |
| 0x002167E4 | `Channel LongSongScanning` | Known | Logging channel |
| 0x00216858 | `Channel VoiceRecording` | Known | Logging channel |
| 0x00216870 | `Channel VoiceRecordingNewFileSegment` | Known | Logging channel |
| 0x00216898 | `Channel PhotoBrowse` | Known | Logging channel |
| 0x002168AC | `Channel PhotoImporting` | Known | Logging channel |
| 0x002168C4 | `Channel Notes` | Known | Logging channel |
| 0x002168D4 | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x002168F0 | `Channel DiskModeChannel` | Known | Logging channel |
| 0x00216908 | `Channel FirewireChannel` | Known | Logging channel |
| 0x00216920 | `Channel USBChannel` | Known | Logging channel |
| 0x00216948 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x002169C0 | `Channel OnTheGoFileMgmt` | Known | Logging channel |
| 0x002169D8 | `Channel SlideShow` | Known | Logging channel |
| 0x002169EC | `Channel ImageCache` | Known | Logging channel |
| 0x00216A00 | `Channel AlbumArtReading` | Known | Logging channel |
| 0x00216A18 | `Channel Video` | Known | Logging channel |
| 0x00216A28 | `Channel DiskImage` | Known | Logging channel |
| 0x00216A3C | `Channel ResourceAccess` | Known | Logging channel |
| 0x00216A54 | `Channel VideoCoreBoot` | Known | Logging channel |
| 0x00216A6C | `Channel DiskFormatConvert` | Known | Logging channel |
| 0x00216A88 | `Channel StreamCacheAddFile` | Known | Logging channel |
| 0x00216AA4 | `Channel FontFileAccess` | Known | Logging channel |
| 0x00216ABC | `Channel ScreenLock` | Known | Logging channel |
| 0x00216B40 | `Channel ProfilerAccess` | Known | Logging channel |
| 0x00216B58 | `Channel eAppAccess` | Known | Logging channel |
| 0x00216B6C | `Channel eAppWriteBackCache` | Known | Logging channel |
| 0x00216B88 | `Channel TrainerFileAccess` | Known | Logging channel |
| 0x00216BA4 | `Channel IapStorage` | Known | Logging channel |
| 0x00216BB8 | `Channel XMLParsing` | Known | Logging channel |
| 0x00216BCC | `Channel AudioPrompt` | Known | Logging channel |
| 0x00216BE0 | `Channel AudioPromptXML` | Known | Logging channel |
| 0x00216BF8 | `Channel StreamCacheSeek` | Known | Logging channel |
| 0x00216C10 | `Channel PredictiveCacheSpinup` | Known | Logging channel |

---

## 7. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002AE43C | `games_RO` | Known | Game system |
| 0x002AE448 | `gamedata_RW` | Known | Game system |
| 0x002AE464 | `gamedata_ShareRW` | Known | Game system |
| 0x0067D0BB | `iPod_Control/games_RO/` | Known | Game system |

---

## 8. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E6304 | `ksidksdrmrdc` | Known | DRM system |
| 0x0017A85C | `AppleDRMVersion` | Known | DRM system |
| 0x0017A8FC | `AppleDRM` | Known | DRM system |
| 0x0017B93C | `AppleVideoDRM` | Known | DRM system |
| 0x00180AF8 | `drmsp608mp4aesdsmp4v` | Known | DRM system |
| 0x0067D58E | `DRMLevel` | Known | DRM system |
| 0x00CC0D10 | `Sdrm` | Known | DRM system |

---

## 9. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000B95E0 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000B95F8 | `iTunesDB` | Known | iTunes database |
| 0x000B9620 | `System_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x000C0874 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000F06C8 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000F803C | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000F9A88 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000F9B84 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x002AE1FC | `iTunesDB` | Known | iTunes database |
| 0x002AE208 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x002C0CA3 | ` auf Ihrem iPod. Weitere Anleitungen finden Sie im iPod Handbuch, in der iTunes ` | Known | iTunes database |
| 0x0067D4B8 | `iPod_Control/iTunes/` | Known | iTunes database |

---

## 10. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017AD94 | `FireWireGUID` | Known | FireWire |
| 0x0017ADA4 | `FireWireVersion` | Known | FireWire |
| 0x0017B390 | `FireWire` | Known | FireWire |
| 0x002B3B37 | `es FireWire nen` | Known | FireWire |
| 0x002B5AB0 | `FireWire p` | Known | FireWire |
| 0x002BA000 | `FireWire-forbindelser underst` | Known | FireWire |
| 0x002BBD20 | `FireWire tilsluttet` | Known | FireWire |
| 0x002C08A8 | `FireWire wird nicht unterst` | Known | FireWire |
| 0x002C292E | `ber FireWire verbunden` | Known | FireWire |
| 0x002C84C6 | ` FireWire. ` | Known | FireWire |
| 0x002CBA66 | ` FireWire` | Known | FireWire |
| 0x002D01B9 | `de canciones o archivos no son posibles con FireWire :` | Known | FireWire |
| 0x002D20E8 | `FireWire conectado` | Known | FireWire |
| 0x002D6718 | `FireWire-tiedonsiirtoa ei tueta. Siirt` | Known | FireWire |
| 0x002D8418 | `FireWire liitetty` | Known | FireWire |
| 0x002DD6E5 | `s via FireWire : connectez l` | Known | FireWire |
| 0x002DF908 | `FireWire Connect` | Known | FireWire |
| 0x002E4114 | `A FireWire kapcsolat nem t` | Known | FireWire |
| 0x002E6324 | `FireWire csatlakozik` | Known | FireWire |
| 0x002EA98C | `Connessioni di dati via FireWire non sono supportate. Per trasferire brani o dat` | Known | FireWire |
| 0x002EC7DC | `FireWire connesso` | Known | FireWire |
| 0x002F17AC | `FireWire ` | Known | FireWire |
| 0x002F3EC8 | `FireWire ` | Known | FireWire |
| 0x002F8AE8 | `FireWire ` | Known | FireWire |
| 0x002FABC4 | `FireWire ` | Known | FireWire |
| 0x002FF77A | `ren via FireWire, maar alleen via de meegeleverde USB-kabel.` | Known | FireWire |
| 0x00301588 | `FireWire aangesloten` | Known | FireWire |
| 0x00305ACF | `ring via FireWire st` | Known | FireWire |
| 0x00307770 | `Koblet til via FireWire` | Known | FireWire |
| 0x0030BE7F | `czenie FireWire nie jest wspierane. By przes` | Known | FireWire |
| 0x0030DDCF | `czony przez Firewire` | Known | FireWire |
| 0x003124CF | `es FireWire n` | Known | FireWire |
| 0x003143D0 | `FireWire ligado` | Known | FireWire |
| 0x00319D95 | ` FireWire ` | Known | FireWire |
| 0x0031D1F7 | ` FireWire` | Known | FireWire |
| 0x00321824 | `FireWire-` | Known | FireWire |
| 0x00323540 | `FireWire anslutet` | Known | FireWire |
| 0x00327B7C | `FireWire ba` | Known | FireWire |
| 0x00329A88 | `FireWire Ba` | Known | FireWire |
| 0x0032E251 | ` FireWire ` | Known | FireWire |
| 0x0032FE94 | `FireWire ` | Known | FireWire |
| 0x003346A5 | ` FireWire ` | Known | FireWire |
| 0x003363CC | `FireWire ` | Known | FireWire |
| 0x004B6E08 | `FireWire connections are not supported. To transfer songs, connect the USB cable` | Known | FireWire |
| 0x004B8AB0 | `FireWire Connected` | Known | FireWire |
| 0x006CCDC1 | `USBCompositeDevice1.6` | Known | USB |
| 0x006CCE19 | `USBCompositeDevice1.6` | Known | USB |

---

## 11. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002BF660 | `Radio-Region` | Known | FM Radio |
| 0x002C1C38 | `Radio-Region` | Known | FM Radio |
| 0x004B5738 | `Radio Region` | Known | FM Radio |
| 0x004B7E6C | `Radio Region` | Known | FM Radio |

---

## 12. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x006CD0AD | `ICAType4CameraDriver` | Known | Camera |
| 0x006CD105 | `PTPCameraDriver` | Known | Camera |
| 0x00B13ED0 | `camera_control` | Known | Camera |

---

## 13. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000986D0 | `iPod_Control` | Filesystem Path |  |
| 0x000986FC | `iPod_Control\Device` | Filesystem Path |  |
| 0x000A4F28 | `iPod_Control\Device` | Filesystem Path |  |
| 0x000A6908 | `iPod_Control` | Filesystem Path |  |
| 0x000A6F48 | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x000BC108 | `iPod_Control\Music\` | Filesystem Path |  |
| 0x000C06F4 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x000E3894 | `iPod_Control/%s%s%s` | Filesystem Path |  |
| 0x000E38A8 | `iPod_Control/%s/%s%s%s` | Filesystem Path |  |
| 0x001AC3AC | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001ACD40 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path |  |
| 0x001ACD68 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001ACED4 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001D7F64 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D81F0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D82A8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D8404 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D8524 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D85F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D8754 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D883C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D88F8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D89A8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D8A9C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D8B40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D8BF4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D8CB0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D8DE4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D8F54 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9018 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D90C8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9204 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D92D0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D939C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9464 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9508 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D95D0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9680 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9730 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D97F8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D98C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9980 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9A30 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9AE0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9B90 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9C40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9CF0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9DC4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9EC4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D9FA4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001DA0AC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001DA198 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0024723C | `iPod_Control/Device` | Filesystem Path |  |
| 0x00247250 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00247744 | `Resources/Fonts` | Filesystem Path |  |
| 0x002AE89C | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x002AF04A | `iPod_Control/Device` | Filesystem Path |  |
| 0x002AF060 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x002AF522 | `iPod_Control/Device` | Filesystem Path |  |
| 0x0067D171 | `iPod_Control/Device` | Filesystem Path |  |
| 0x0067D185 | `iPod_Control/Device/accessories` | Filesystem Path |  |
| 0x0067D5F6 | `/Resources/VideoCore` | Filesystem Path |  |

---

## 14. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002B2F98 | `Acoustic` | EQ Preset |  |
| 0x002B2FA4 | `Bass Booster` | EQ Preset |  |
| 0x002B2FC4 | `Classical` | EQ Preset |  |
| 0x002B2FD0 | `Dance` | EQ Preset |  |
| 0x002B2FE0 | `Electronic` | EQ Preset |  |
| 0x002B2FF4 | `Hip Hop` | EQ Preset |  |
| 0x002B2FFC | `Jazz` | EQ Preset |  |
| 0x002B3004 | `Latin` | EQ Preset |  |
| 0x002B300C | `Loudness` | EQ Preset |  |
| 0x002B3018 | `Lounge` | EQ Preset |  |
| 0x002B3020 | `Piano` | EQ Preset |  |
| 0x002B3034 | `Rock` | EQ Preset |  |
| 0x002B303C | `Small Speakers` | EQ Preset |  |
| 0x002B304C | `Spoken Word` | EQ Preset |  |
| 0x002B3058 | `Treble Booster` | EQ Preset |  |
| 0x002B3078 | `Vocal Booster` | EQ Preset |  |
| 0x002B94A4 | `Acoustic` | EQ Preset |  |
| 0x002B94B0 | `Bass Booster` | EQ Preset |  |
| 0x002B94D0 | `Classical` | EQ Preset |  |
| 0x002B94DC | `Dance` | EQ Preset |  |
| 0x002B94EC | `Electronic` | EQ Preset |  |
| 0x002B9500 | `Hip Hop` | EQ Preset |  |
| 0x002B9508 | `Jazz` | EQ Preset |  |
| 0x002B9510 | `Latin` | EQ Preset |  |
| 0x002B9518 | `Loudness` | EQ Preset |  |
| 0x002B9524 | `Lounge` | EQ Preset |  |
| 0x002B952C | `Piano` | EQ Preset |  |
| 0x002B9540 | `Rock` | EQ Preset |  |
| 0x002B9548 | `Small Speakers` | EQ Preset |  |
| 0x002B9558 | `Spoken Word` | EQ Preset |  |
| 0x002B9564 | `Treble Booster` | EQ Preset |  |
| 0x002B9584 | `Vocal Booster` | EQ Preset |  |
| 0x002BFD48 | `Acoustic` | EQ Preset |  |
| 0x002BFD78 | `Dance` | EQ Preset |  |
| 0x002BFD88 | `Electronic` | EQ Preset |  |
| 0x002BFD9C | `Hip Hop` | EQ Preset |  |
| 0x002BFDA4 | `Jazz` | EQ Preset |  |
| 0x002BFDAC | `Latin` | EQ Preset |  |
| 0x002BFDB4 | `Loudness` | EQ Preset |  |
| 0x002BFDC8 | `Piano` | EQ Preset |  |
| 0x002BFDDC | `Rock` | EQ Preset |  |
| 0x002C70D4 | `Hip Hop` | EQ Preset |  |
| 0x002C70DC | `Jazz` | EQ Preset |  |
| 0x002C70E4 | `Latin` | EQ Preset |  |
| 0x002C70EC | `Loudness` | EQ Preset |  |
| 0x002C70F8 | `Lounge` | EQ Preset |  |
| 0x002C7100 | `Piano` | EQ Preset |  |
| 0x002C7114 | `Rock` | EQ Preset |  |
| 0x002CF674 | `Dance` | EQ Preset |  |
| 0x002CF69C | `Hip Hop` | EQ Preset |  |
| 0x002CF6A4 | `Jazz` | EQ Preset |  |
| 0x002CF6B4 | `Loudness` | EQ Preset |  |
| 0x002CF6C0 | `Lounge` | EQ Preset |  |
| 0x002CF6C8 | `Piano` | EQ Preset |  |
| 0x002CF6DC | `Rock` | EQ Preset |  |
| 0x002D5BF4 | `Jazz` | EQ Preset |  |
| 0x002D5BFC | `Latin` | EQ Preset |  |
| 0x002D5C10 | `Lounge` | EQ Preset |  |
| 0x002D5C18 | `Piano` | EQ Preset |  |
| 0x002D5C2C | `Rock` | EQ Preset |  |
| 0x002DCB34 | `Hip Hop` | EQ Preset |  |
| 0x002DCB3C | `Jazz` | EQ Preset |  |
| 0x002DCB64 | `Lounge` | EQ Preset |  |
| 0x002DCB6C | `Piano` | EQ Preset |  |
| 0x002DCB84 | `Rock` | EQ Preset |  |
| 0x002E35B4 | `Latin` | EQ Preset |  |
| 0x002E35E0 | `Rock` | EQ Preset |  |
| 0x002E9E30 | `Dance` | EQ Preset |  |
| 0x002E9E54 | `Hip Hop` | EQ Preset |  |
| 0x002E9E5C | `Jazz` | EQ Preset |  |
| 0x002E9E6C | `Loudness` | EQ Preset |  |
| 0x002E9E78 | `Lounge` | EQ Preset |  |
| 0x002E9E80 | `Piano` | EQ Preset |  |
| 0x002E9E94 | `Rock` | EQ Preset |  |
| 0x002F0634 | `Acoustic` | EQ Preset |  |
| 0x002F0640 | `Bass Booster` | EQ Preset |  |
| 0x002F0660 | `Classical` | EQ Preset |  |
| 0x002F066C | `Dance` | EQ Preset |  |
| 0x002F067C | `Electronic` | EQ Preset |  |
| 0x002F0690 | `Hip Hop` | EQ Preset |  |
| 0x002F0698 | `Jazz` | EQ Preset |  |
| 0x002F06A0 | `Latin` | EQ Preset |  |
| 0x002F06A8 | `Loudness` | EQ Preset |  |
| 0x002F06B4 | `Lounge` | EQ Preset |  |
| 0x002F06BC | `Piano` | EQ Preset |  |
| 0x002F06D0 | `Rock` | EQ Preset |  |
| 0x002F06D8 | `Small Speakers` | EQ Preset |  |
| 0x002F06E8 | `Spoken Word` | EQ Preset |  |
| 0x002F06F4 | `Treble Booster` | EQ Preset |  |
| 0x002F0714 | `Vocal Booster` | EQ Preset |  |
| 0x002F7C04 | `Acoustic` | EQ Preset |  |
| 0x002F7C10 | `Bass Booster` | EQ Preset |  |
| 0x002F7C30 | `Classical` | EQ Preset |  |
| 0x002F7C3C | `Dance` | EQ Preset |  |
| 0x002F7C4C | `Electronic` | EQ Preset |  |
| 0x002F7C60 | `Hip Hop` | EQ Preset |  |
| 0x002F7C68 | `Jazz` | EQ Preset |  |
| 0x002F7C70 | `Latin` | EQ Preset |  |
| 0x002F7C78 | `Loudness` | EQ Preset |  |
| 0x002F7C84 | `Lounge` | EQ Preset |  |
| 0x002F7C8C | `Piano` | EQ Preset |  |
| 0x002F7CA0 | `Rock` | EQ Preset |  |
| 0x002F7CA8 | `Small Speakers` | EQ Preset |  |
| 0x002F7CB8 | `Spoken Word` | EQ Preset |  |
| 0x002F7CC4 | `Treble Booster` | EQ Preset |  |
| 0x002F7CE4 | `Vocal Booster` | EQ Preset |  |
| 0x002FEC20 | `Dance` | EQ Preset |  |
| 0x002FEC54 | `Jazz` | EQ Preset |  |
| 0x002FEC5C | `Latin` | EQ Preset |  |
| 0x002FEC64 | `Loudness` | EQ Preset |  |
| 0x002FEC70 | `Lounge` | EQ Preset |  |
| 0x002FEC78 | `Piano` | EQ Preset |  |
| 0x002FEC8C | `Rock` | EQ Preset |  |
| 0x00304F94 | `Dance` | EQ Preset |  |
| 0x00304FC0 | `Jazz` | EQ Preset |  |
| 0x00304FD0 | `Loudness` | EQ Preset |  |
| 0x00304FDC | `Lounge` | EQ Preset |  |
| 0x00304FE4 | `Piano` | EQ Preset |  |
| 0x00304FF8 | `Rock` | EQ Preset |  |
| 0x0030B30C | `Hip Hop` | EQ Preset |  |
| 0x0030B314 | `Jazz` | EQ Preset |  |
| 0x0030B340 | `Lounge` | EQ Preset |  |
| 0x0030B348 | `Piano` | EQ Preset |  |
| 0x0030B35C | `Rock` | EQ Preset |  |
| 0x00311968 | `Dance` | EQ Preset |  |
| 0x00311990 | `Hip Hop` | EQ Preset |  |
| 0x00311998 | `Jazz` | EQ Preset |  |
| 0x003119A8 | `Loudness` | EQ Preset |  |
| 0x003119B4 | `Lounge` | EQ Preset |  |
| 0x003119BC | `Piano` | EQ Preset |  |
| 0x003119D0 | `Rock` | EQ Preset |  |
| 0x00320CDC | `Acoustic` | EQ Preset |  |
| 0x00320CE8 | `Bass Booster` | EQ Preset |  |
| 0x00320D08 | `Classical` | EQ Preset |  |
| 0x00320D14 | `Dance` | EQ Preset |  |
| 0x00320D24 | `Electronic` | EQ Preset |  |
| 0x00320D38 | `Hip Hop` | EQ Preset |  |
| 0x00320D40 | `Jazz` | EQ Preset |  |
| 0x00320D48 | `Latin` | EQ Preset |  |
| 0x00320D50 | `Loudness` | EQ Preset |  |
| 0x00320D5C | `Lounge` | EQ Preset |  |
| 0x00320D64 | `Piano` | EQ Preset |  |
| 0x00320D78 | `Rock` | EQ Preset |  |
| 0x00320D80 | `Small Speakers` | EQ Preset |  |
| 0x00320D90 | `Spoken Word` | EQ Preset |  |
| 0x00320D9C | `Treble Booster` | EQ Preset |  |
| 0x00320DBC | `Vocal Booster` | EQ Preset |  |
| 0x00327088 | `Hip Hop` | EQ Preset |  |
| 0x00327094 | `Latin` | EQ Preset |  |
| 0x0032709C | `Loudness` | EQ Preset |  |
| 0x003270A8 | `Lounge` | EQ Preset |  |
| 0x003270C4 | `Rock` | EQ Preset |  |
| 0x0032D49C | `Acoustic` | EQ Preset |  |
| 0x0032D4A8 | `Bass Booster` | EQ Preset |  |
| 0x0032D4C8 | `Classical` | EQ Preset |  |
| 0x0032D4D4 | `Dance` | EQ Preset |  |
| 0x0032D4E4 | `Electronic` | EQ Preset |  |
| 0x0032D4F8 | `Hip Hop` | EQ Preset |  |
| 0x0032D500 | `Jazz` | EQ Preset |  |
| 0x0032D508 | `Latin` | EQ Preset |  |
| 0x0032D510 | `Loudness` | EQ Preset |  |
| 0x0032D51C | `Lounge` | EQ Preset |  |
| 0x0032D524 | `Piano` | EQ Preset |  |
| 0x0032D538 | `Rock` | EQ Preset |  |
| 0x0032D540 | `Small Speakers` | EQ Preset |  |
| 0x0032D550 | `Spoken Word` | EQ Preset |  |
| 0x0032D55C | `Treble Booster` | EQ Preset |  |
| 0x0032D57C | `Vocal Booster` | EQ Preset |  |
| 0x00333914 | `Acoustic` | EQ Preset |  |
| 0x00333920 | `Bass Booster` | EQ Preset |  |
| 0x00333940 | `Classical` | EQ Preset |  |
| 0x0033394C | `Dance` | EQ Preset |  |
| 0x0033395C | `Electronic` | EQ Preset |  |
| 0x00333970 | `Hip Hop` | EQ Preset |  |
| 0x00333978 | `Jazz` | EQ Preset |  |
| 0x00333980 | `Latin` | EQ Preset |  |
| 0x00333988 | `Loudness` | EQ Preset |  |
| 0x00333994 | `Lounge` | EQ Preset |  |
| 0x0033399C | `Piano` | EQ Preset |  |
| 0x003339B0 | `Rock` | EQ Preset |  |
| 0x003339B8 | `Small Speakers` | EQ Preset |  |
| 0x003339C8 | `Spoken Word` | EQ Preset |  |
| 0x003339D4 | `Treble Booster` | EQ Preset |  |
| 0x003339F4 | `Vocal Booster` | EQ Preset |  |
| 0x004B5F1C | `Acoustic` | EQ Preset |  |
| 0x004B5F28 | `Bass Booster` | EQ Preset |  |
| 0x004B5F48 | `Classical` | EQ Preset |  |
| 0x004B5F54 | `Dance` | EQ Preset |  |
| 0x004B5F64 | `Electronic` | EQ Preset |  |
| 0x004B5F78 | `Hip Hop` | EQ Preset |  |
| 0x004B5F80 | `Jazz` | EQ Preset |  |
| 0x004B5F88 | `Latin` | EQ Preset |  |
| 0x004B5F90 | `Loudness` | EQ Preset |  |
| 0x004B5F9C | `Lounge` | EQ Preset |  |
| 0x004B5FA4 | `Piano` | EQ Preset |  |
| 0x004B5FB8 | `Rock` | EQ Preset |  |
| 0x004B5FC0 | `Small Speakers` | EQ Preset |  |
| 0x004B5FD0 | `Spoken Word` | EQ Preset |  |
| 0x004B5FDC | `Treble Booster` | EQ Preset |  |
| 0x004B5FFC | `Vocal Booster` | EQ Preset |  |

---
