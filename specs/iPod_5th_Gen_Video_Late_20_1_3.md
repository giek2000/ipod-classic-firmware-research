# iPod 5th Gen Late 2005 (Video 60GB) - RetailOS 1.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.3 |
| **IPSW** | iPod_20.1.3.ipsw |
| **Device** | iPod 5th Gen Late 2005 (Video 60GB) (2005, Click Wheel, Video Playback) |
| **UpdaterFamilyID** | 20 |
| **Binary Size** | 13,893,632 bytes (13.25 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 13,893,632 bytes |
| **Total Strings (>=6)** | 30,182 |
| **Function Prologues** | 22,648 (ARM: 12,890, Thumb: 9,758) |
| **SoC** | PortalPlayer PP5021C |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Encrypted** | No |
| **SHA-256** | `182632b68c54103693ef8cfe8b248fe23f8b733973ee3b8a756bc1c3f9aa2c88` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016EF04 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x00215F34 | `Channel UnitTests` | Hidden | Developer Tool |

---

## 2. Controllers (TSilver/TC Classes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00B26C78 | `TCC_Current_Task_Pointer` | Known | UI controller |
| 0x00B26C91 | `TCC_Delete_HISR` | Known | UI controller |
| 0x00B26CB1 | `TCC_Relinquish` | Known | UI controller |
| 0x00B26CCF | `TCC_Resume_Service` | Known | UI controller |
| 0x00B26CE2 | `TCC_Task_Sleep` | Known | UI controller |
| 0x00B26D04 | `TCF_Task_Information` | Known | UI controller |
| 0x00B26D19 | `TCS_Change_Preemption` | Known | UI controller |
| 0x00B26D2F | `TCS_Change_Priority` | Known | UI controller |
| 0x00B26D43 | `TCT_Activate_HISR` | Known | UI controller |
| 0x00B26D55 | `TCT_Control_Interrupts` | Known | UI controller |
| 0x00B26D6C | `TCT_Local_Control_Interrupts` | Known | UI controller |
| 0x00B8E51F | `TCC_Resume_Service` | Known | UI controller |
| 0x00B8E6A2 | `TCC_Delete_HISR` | Known | UI controller |
| 0x00B8E6F1 | `TCT_Activate_HISR` | Known | UI controller |
| 0x00B8E71B | `TCT_Control_Interrupts` | Known | UI controller |
| 0x00B8E8E5 | `TCC_Current_Task_Pointer` | Known | UI controller |
| 0x00B8E8FE | `TCS_Change_Priority` | Known | UI controller |
| 0x00B8E96C | `TCC_Task_Sleep` | Known | UI controller |
| 0x00B8EAA5 | `TCF_Task_Information` | Known | UI controller |
| 0x00B9D01D | `TCC_Relinquish` | Known | UI controller |
| 0x00B9D133 | `TCT_Local_Control_Interrupts` | Known | UI controller |
| 0x00B9D1D3 | `TCC_Task_Sleep` | Known | UI controller |
| 0x00B9D2A1 | `TCS_Change_Preemption` | Known | UI controller |

---

## 3. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00098774 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x000C9120 | `USB Secondary Interrupt Task` | Known | RTOS task thread |
| 0x000E294C | `ICAPTPCameraIOTask` | Known | RTOS task thread |
| 0x0011D1A8 | `USBStatusTask` | Known | RTOS task thread |
| 0x0011D1C4 | `USBTaskTimeTask` | Known | RTOS task thread |
| 0x0016902C | `VCUpdateTask` | Known | RTOS task thread |
| 0x00174E7C | `USBDeviceTask` | Known | RTOS task thread |
| 0x0017B5B4 | `DiskReaderTask` | Known | RTOS task thread |
| 0x0018A9BC | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0018A9D0 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0019DDF8 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x0020FC64 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x00216128 | `Channel DiskReaderTask` | Known | RTOS task thread |
| 0x0026143C | `FirewireTask` | Known | RTOS task thread |
| 0x00261454 | `OptoTask` | Known | RTOS task thread |
| 0x00261464 | `SerialOptoTask` | Known | RTOS task thread |
| 0x00261478 | `BacklightTask` | Known | RTOS task thread |
| 0x0026148C | `CNATask` | Known | RTOS task thread |
| 0x002614AC | `DiskMgrTask` | Known | RTOS task thread |
| 0x002614BC | `HoldSwitchTask` | Known | RTOS task thread |
| 0x002614D0 | `TopPlugTask` | Known | RTOS task thread |
| 0x002614E0 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x002614F4 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x0026150C | `AccessoryDetectTask` | Known | RTOS task thread |
| 0x00261534 | `AlarmTask` | Known | RTOS task thread |
| 0x00261544 | `WatchdogTask` | Known | RTOS task thread |
| 0x002615BC | `USBAudioTask` | Known | RTOS task thread |
| 0x002ADA10 | `HostOSTask` | Known | RTOS task thread |
| 0x002AE5FC | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x0050495C | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x00504974 | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x00504A98 | `VideoDaisyTask` | Known | RTOS task thread |
| 0x00B26C68 | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B26CA1 | `TCC_Delete_Task` | Known | RTOS task thread |
| 0x00B26CC0 | `TCC_Reset_Task` | Known | RTOS task thread |
| 0x00B26CF1 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B8E510 | `TCC_Reset_Task` | Known | RTOS task thread |
| 0x00B8E547 | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B8E6B2 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B8E6C5 | `TCC_Delete_Task` | Known | RTOS task thread |
| 0x00B9D00D | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B9D057 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B9D06A | `TCC_Delete_Task` | Known | RTOS task thread |

---

## 4. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00215CF4 | `Channel Reserved` | Known | Logging channel |
| 0x00215D08 | `Channel AppBoot` | Known | Logging channel |
| 0x00215D18 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x00215D34 | `Channel PrefsWriting` | Known | Logging channel |
| 0x00215D4C | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x00215D6C | `Channel PlayFromDisk` | Known | Logging channel |
| 0x00215D84 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x00215DA0 | `Channel TestLogging` | Known | Logging channel |
| 0x00215DB4 | `Channel AppFileLoading` | Known | Logging channel |
| 0x00215DCC | `Channel VCardReading` | Known | Logging channel |
| 0x00215DE4 | `Channel LongSongScanning` | Known | Logging channel |
| 0x00215E58 | `Channel VoiceRecording` | Known | Logging channel |
| 0x00215E70 | `Channel VoiceRecordingNewFileSegment` | Known | Logging channel |
| 0x00215E98 | `Channel PhotoBrowse` | Known | Logging channel |
| 0x00215EAC | `Channel PhotoImporting` | Known | Logging channel |
| 0x00215EC4 | `Channel Notes` | Known | Logging channel |
| 0x00215ED4 | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x00215EF0 | `Channel DiskModeChannel` | Known | Logging channel |
| 0x00215F08 | `Channel FirewireChannel` | Known | Logging channel |
| 0x00215F20 | `Channel USBChannel` | Known | Logging channel |
| 0x00215F48 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x00215FC0 | `Channel OnTheGoFileMgmt` | Known | Logging channel |
| 0x00215FD8 | `Channel SlideShow` | Known | Logging channel |
| 0x00215FEC | `Channel ImageCache` | Known | Logging channel |
| 0x00216000 | `Channel AlbumArtReading` | Known | Logging channel |
| 0x00216018 | `Channel Video` | Known | Logging channel |
| 0x00216028 | `Channel DiskImage` | Known | Logging channel |
| 0x0021603C | `Channel ResourceAccess` | Known | Logging channel |
| 0x00216054 | `Channel VideoCoreBoot` | Known | Logging channel |
| 0x0021606C | `Channel DiskFormatConvert` | Known | Logging channel |
| 0x00216088 | `Channel StreamCacheAddFile` | Known | Logging channel |
| 0x002160A4 | `Channel FontFileAccess` | Known | Logging channel |
| 0x002160BC | `Channel ScreenLock` | Known | Logging channel |
| 0x00216140 | `Channel ProfilerAccess` | Known | Logging channel |
| 0x00216158 | `Channel eAppAccess` | Known | Logging channel |
| 0x0021616C | `Channel eAppWriteBackCache` | Known | Logging channel |
| 0x00216188 | `Channel TrainerFileAccess` | Known | Logging channel |
| 0x002161A4 | `Channel IapStorage` | Known | Logging channel |
| 0x002161B8 | `Channel XMLParsing` | Known | Logging channel |
| 0x002161CC | `Channel AudioPrompt` | Known | Logging channel |
| 0x002161E0 | `Channel AudioPromptXML` | Known | Logging channel |
| 0x002161F8 | `Channel StreamCacheSeek` | Known | Logging channel |
| 0x00216210 | `Channel PredictiveCacheSpinup` | Known | Logging channel |

---

## 5. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00179E88 | `AudioCodecs` | Known | Audio system |
| 0x0017AF28 | `VideoCodecs` | Known | Audio system |
| 0x002BA0FA | `.net codec i dette produkt bruges i henhold til en licensaftale fra VoiceAge Cor` | Known | Audio system |
| 0x002C0BC9 | `.net Codec in diesem Produkt wird in Lizenz der VoiceAge Corporation verwendet. ` | Known | Audio system |
| 0x002C8D74 | `.net codec ` | Known | Audio system |
| 0x002DD8E8 | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x002E4381 | `.net codec a VoiceAge Coporation c` | Known | Audio system |
| 0x002EAA92 | ` utilizzato su licenza da VoiceAge Corporation. Il codec ACELP` | Known | Audio system |
| 0x002F8D42 | `.net codec` | Known | Audio system |
| 0x002FF874 | `.net-codec in dit product wordt gebruikt in licentie van VoiceAge Corporation. G` | Known | Audio system |
| 0x0031255D | `a da VoiceAge Corporation. O codec ACELP` | Known | Audio system |
| 0x00327D09 | `.net codec'i VoiceAge Corporation lisans` | Known | Audio system |
| 0x004B6EB5 | `.net codec in this product is used under license from VoiceAge Corporation. Port` | Known | Audio system |
| 0x0067FA8E | `msCodeCom` | Known | Audio system |
| 0x00B8DF9D | `codec_string` | Known | Audio system |
| 0x00B8DFAA | `codec_name` | Known | Audio system |
| 0x00B9C8C9 | `codec_string` | Known | Audio system |
| 0x00B9C8D6 | `codec_name` | Known | Audio system |

---

## 6. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00179F64 | `Audible` | Known | Audible audiobook format |
| 0x002B3B1D | ` Audible v` | Known | Audible audiobook format |
| 0x002B3B6F | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x002B3B85 | ` 2002 Audible, Inc. V` | Known | Audible audiobook format |
| 0x002B9FA8 | `Audible-softwaren i dette produkt bruges i henhold til en licensaftale fra Audib` | Known | Audible audiobook format |
| 0x002BA008 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x002C0A84 | `Die Audible Software in diesem Produkt wird in Lizenz der Audible verwendet. Cop` | Known | Audible audiobook format |
| 0x002C0ADD | ` 2002 Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x002C8B63 | ` Audible ` | Known | Audible audiobook format |
| 0x002C8BC0 | ` Audible. ` | Known | Audible audiobook format |
| 0x002C8BF6 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002D01A8 | `El software Audible incluido en este producto se usa bajo licencia de Audible. C` | Known | Audible audiobook format |
| 0x002D0203 | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x002D6696 | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright ` | Known | Audible audiobook format |
| 0x002D66D0 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x002DD7D8 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x002DD822 | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x002DD837 | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x002E4242 | ` Audible szoftver az Audible licence alatt van haszn` | Known | Audible audiobook format |
| 0x002E428C | ` 2002, Audible, Inc. Minden jog fenntartva.` | Known | Audible audiobook format |
| 0x002EA9C8 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x002EA9F1 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x002EAA20 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x002F1B41 | ` Audible ` | Known | Audible audiobook format |
| 0x002F1B62 | `Audible ` | Known | Audible audiobook format |
| 0x002F1BBB | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002F8BF3 | ` Audible ` | Known | Audible audiobook format |
| 0x002F8C0E | ` Audible` | Known | Audible audiobook format |
| 0x002F8C52 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002FF72C | `De Audible-software in dit product wordt gebruikt in licentie van Audible. Copyr` | Known | Audible audiobook format |
| 0x002FF783 | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x00305A5C | `Audible-programvaren i dette produktet brukes under lisens fra Audible. Copyrigh` | Known | Audible audiobook format |
| 0x00305AB0 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0030BE64 | `Oprogramowanie Audible w tym produkcie jest wykorzystywane na podstawie licencji` | Known | Audible audiobook format |
| 0x0030BED0 | ` 2002 Audible, Inc. Wszystkie prawa zastrze` | Known | Audible audiobook format |
| 0x0031244C | `O software Audible ` | Known | Audible audiobook format |
| 0x00312482 | `a da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0031249C | ` 2002 da Audible, Inc. Reservados todos os direitos.` | Known | Audible audiobook format |
| 0x0031A4AD | ` Audible ` | Known | Audible audiobook format |
| 0x0031A4FF | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x0031A515 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x003217AC | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x003217DB | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x003217F2 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x00327BC0 | `ndeki Audible yaz` | Known | Audible audiobook format |
| 0x00327BD9 | ` Audible lisans` | Known | Audible audiobook format |
| 0x00327C0E | ` 2002, Audible, Inc. T` | Known | Audible audiobook format |
| 0x0032E1DB | ` Audible ` | Known | Audible audiobook format |
| 0x0032E1ED | ` Audible ` | Known | Audible audiobook format |
| 0x0032E211 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x0033464C | `Audible ` | Known | Audible audiobook format |
| 0x00334660 | ` Audible ` | Known | Audible audiobook format |
| 0x0033468A | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x004B6D7C | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x004B6DD1 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |

---

## 7. Audio/Codec - Apple Lossless

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00179F38 | `AppleLossless` | Known | Apple Lossless codec |
| 0x002E479C | `l alacsony.` | Known | Apple Lossless codec |

---

## 8. Audio/Codec - AAC

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00B37C60 | `AACDEC  VLL ` | Known | AAC codec |
| 0x00B413A8 | `MPEG4 AAC LC Decoder` | Known | AAC codec |
| 0x00B4380D | `AACDecoderGetMem` | Known | AAC codec |
| 0x00B4381E | `AACDecoderInit` | Known | AAC codec |
| 0x00B4382D | `AACDecoderGetConfig` | Known | AAC codec |
| 0x00B43841 | `AACDecoderSetConfig` | Known | AAC codec |
| 0x00B43855 | `AACHeaderDecode` | Known | AAC codec |
| 0x00B43865 | `AACDecode` | Known | AAC codec |
| 0x00B4386F | `AACDecoderInit_Ittiam` | Known | AAC codec |
| 0x00B43885 | `AACDecoderGetConfig_Ittiam` | Known | AAC codec |
| 0x00B438A0 | `AACDecoderSetConfig_Ittiam` | Known | AAC codec |
| 0x00B438BB | `AACHeaderDecode_Ittiam` | Known | AAC codec |
| 0x00B438D2 | `AACDecode_Ittiam` | Known | AAC codec |

---

## 9. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002B3D62 | ` zvuku MPEG Layer-3 byla poskytnuta spole` | Known | MP3 codec |
| 0x002B3D8D | `nostmi Fraunhofer IIS a` | Known | MP3 codec |
| 0x002BA1A4 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunhofer IIS og THOMSON multi` | Known | MP3 codec |
| 0x002C0C87 | `r MPEG Layer-3 wurde lizenziert von Fraunhofer IIS und THOMSON Multimedia.` | Known | MP3 codec |
| 0x002C8EBB | ` MPEG Layer-3 ` | Known | MP3 codec |
| 0x002C8EF9 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x002D03A1 | `n de audio MPEG Layer-3 utilizada bajo licencia de Fraunhofer IIS y THOMSON mult` | Known | MP3 codec |
| 0x002D683C | `MPEG Layer-3 -` | Known | MP3 codec |
| 0x002D684E | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | MP3 codec |
| 0x002DD9F4 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | MP3 codec |
| 0x002E4414 | `Az MPEG Layer-3 hangk` | Known | MP3 codec |
| 0x002E443C | `gia a Fraunhofer IIS ` | Known | MP3 codec |
| 0x002EAB68 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da Fraunhofer IIS e THOMSON` | Known | MP3 codec |
| 0x002F1D70 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002F1DBC | `Fraunhofer IIS ` | Known | MP3 codec |
| 0x002F8DDC | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002F8E03 | ` Fraunhofer IIS` | Known | MP3 codec |
| 0x002FF910 | `Technologie voor codering van MPEG Layer-3-audio in licentie van Fraunhofer IIS ` | Known | MP3 codec |
| 0x00305C2C | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fraunhofer IIS og THOMSON m` | Known | MP3 codec |
| 0x0030C078 | `Technologia kodowania audio MPEG Layer-3 licencjonowana od Fraunhofer IIS oraz T` | Known | MP3 codec |
| 0x0031263E | `udio MPEG Layer-3 licenciada pela Fraunhofer IIS e THOMSON multimedia.` | Known | MP3 codec |
| 0x0031A734 | `MPEG Layer-3: ` | Known | MP3 codec |
| 0x0031A78D | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x0032198C | `Ljudkodningstekniken MPEG Layer-3 ` | Known | MP3 codec |
| 0x003219C2 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | MP3 codec |
| 0x00327D98 | `MPEG Layer-3 ses kodlama teknolojisi Fraunhofer IIS ve THOMSON multimedia'dan li` | Known | MP3 codec |
| 0x0032E368 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x0032E38A | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x003347E8 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x0033480D | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x004B6F48 | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |

---

## 10. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E6E19 | `;=1sinffniscpsap@-` | Known | DRM system |
| 0x00179E5C | `AppleDRMVersion` | Known | DRM system |
| 0x00179EFC | `AppleDRM` | Known | DRM system |
| 0x0017AF3C | `AppleVideoDRM` | Known | DRM system |
| 0x001800F8 | `drmsp608mp4aesdsmp4v` | Known | DRM system |
| 0x001CD35C | `tkhdedtselstmdiamdhdminfstblstsdstcoco64stscstszsttsstssdrmidrmsx` | Known | DRM system |
| 0x0067CB8E | `DRMLevel` | Known | DRM system |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002ADA3C | `games_RO` | Known | Game system |
| 0x002ADA48 | `gamedata_RW` | Known | Game system |
| 0x002ADA64 | `gamedata_ShareRW` | Known | Game system |

---

## 12. Photo System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00097F74 | `Photo Database` | Known | Photo system |
| 0x000B8C04 | `Photos\Photo Database` | Known | Photo system |
| 0x000C042C | `Photo Database` | Known | Photo system |
| 0x0019B0BC | `23iUPhoto Database` | Known | Photo system |
| 0x0019D114 | `Photo Database` | Known | Photo system |
| 0x0019D478 | `Photo Database` | Known | Photo system |
| 0x0019D724 | `Photo Import Database` | Known | Photo system |
| 0x00217DE0 | `Photo Database Size` | Known | Photo system |

---

## 13. Video System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00B522E0 | `H.264 Video Decoder` | Known | Video system |

---

## 14. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000B8BF8 | `iTunesDB` | Known | iTunes database |
| 0x001C27C4 | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x0020DEC4 | `iTunes Image DB` | Known | iTunes database |
| 0x002AD7FC | `iTunesDB` | Known | iTunes database |
| 0x002B32FD | ` z iTunes nebo vCards. ` | Known | iTunes database |
| 0x002B3617 | `ipojte iPod k iTunes a instalujte hru znovu.` | Known | iTunes database |
| 0x002B36A7 | `ipojte iPod k iTunes a zkop` | Known | iTunes database |
| 0x002B373F | `i a program iTunes jej odemkne.` | Known | iTunes database |
| 0x002B3865 | `m iTunes.` | Known | iTunes database |
| 0x002B41CD | `m iTunes.` | Known | iTunes database |
| 0x002B4B43 | `es iTunes.` | Known | iTunes database |
| 0x002B4BE9 | `es iTunes.` | Known | iTunes database |
| 0x002B979C | `iPod kan opbevare og vise kontaktoplysninger importeret fra iTunes eller vCards.` | Known | iTunes database |
| 0x002B9B24 | `Slut iPod til iTunes, og installer spillet igen.` | Known | iTunes database |
| 0x002B9BBC | `Slut iPod til iTunes, og overf` | Known | iTunes database |
| 0x002B9C53 | `slutte iPod til computeren, hvorefter iTunes l` | Known | iTunes database |
| 0x002B9D08 | `r fotografier til computeren, og synkroniser via iTunes for at vise dem p` | Known | iTunes database |
| 0x002BA550 | `%s er for gammel til denne iPod. Slut iPod til computeren, og start iTunes for a` | Known | iTunes database |
| 0x002BADC6 | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x002BAE55 | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x002C00DD | `hlen Sie Ihren iPod in iTunes aus der Liste ` | Known | iTunes database |
| 0x002C014B | `iTunes` | Known | iTunes database |
| 0x002C0590 | `Verbinden Sie Ihren iPod mit iTunes und installieren Sie das Spiel erneut.` | Known | iTunes database |
| 0x002C0654 | `Verbinden Sie Ihren iPod mit iTunes und laden Sie die aktuelle Version.` | Known | iTunes database |
| 0x002C06FB | `en Sie Ihren iPod an Ihren Computer an und iTunes deaktiviert die Anzeigensperre` | Known | iTunes database |
| 0x002C07A4 | `Importierte Fotos werden nicht auf dem TV angezeigt. Senden Sie sie erst an den ` | Known | iTunes database |
| 0x002C1089 | `en Sie den iPod an den Computer an und starten Sie iTunes, um %s auf die aktuell` | Known | iTunes database |
| 0x002C19A1 | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x002C1A42 | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x002C7DA8 | ` iTunes ` | Known | iTunes database |
| 0x002C82F1 | ` iTunes ` | Known | iTunes database |
| 0x002C8419 | ` iTunes ` | Known | iTunes database |
| 0x002C8542 | ` iTunes ` | Known | iTunes database |
| 0x002C86D7 | ` iTunes ` | Known | iTunes database |
| 0x002C95CA | ` iTunes ` | Known | iTunes database |
| 0x002CA6DA | ` iTunes ` | Known | iTunes database |
| 0x002CA803 | ` iTunes ` | Known | iTunes database |
| 0x002CF979 | `n importada de iTunes o de tarjetas virtuales (vCards). ` | Known | iTunes database |
| 0x002CFD04 | `Conecte el iPod a iTunes y reinstale el juego.` | Known | iTunes database |
| 0x002CFD9C | `Conecte el iPod a iTunes y descargue la versi` | Known | iTunes database |
| 0x002CFE24 | `n, conecte el iPod al ordenador y iTunes lo desbloquear` | Known | iTunes database |
| 0x002CFF00 | `celas con iTunes para verlas en la TV.` | Known | iTunes database |
| 0x002D0798 | `%s es demasiado antiguo para ejecutarse en este iPod. Conecte el iPod al ordenad` | Known | iTunes database |
| 0x002D10D8 | `Esta foto es demasiado grande para mostrarla en el iPod. Transfiera las fotos im` | Known | iTunes database |
| 0x002D117C | `Este formato de foto no puede visualizarse en el iPod. Transfiera las fotos impo` | Known | iTunes database |
| 0x002D5ED8 | ` iTunesista tai vCardeina tuotua tietoa. ` | Known | iTunes database |
| 0x002D623E | ` iPod iTunesiin ja asenna peli uudelleen.` | Known | iTunes database |
| 0x002D62CE | ` iPod iTunesiin ja hae uusin versio.` | Known | iTunes database |
| 0x002D6343 | `tietokoneeseen, niin iTunes avaa lukituksen.` | Known | iTunes database |
| 0x002D63E0 | ` kuvat tietokoneelle ja synkronoi ne iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002D6C2E | ` %s uusimpaan versioon avaamalla iTunes.` | Known | iTunes database |
| 0x002D74A6 | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002D752D | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002DCEE2 | `iTunes ou de vCards. ` | Known | iTunes database |
| 0x002DD288 | `Connectez votre iPod avec iTunes et r` | Known | iTunes database |
| 0x002DD330 | `Connectez votre iPod avec iTunes et t` | Known | iTunes database |
| 0x002DD3F1 | ` votre ordinateur et iTunes le d` | Known | iTunes database |
| 0x002DD4D4 | `rez-les sur votre ordinateur puis synchronisez-les avec iTunes.` | Known | iTunes database |
| 0x002DDDFC | `ordinateur et lancez iTunes pour mettre ` | Known | iTunes database |
| 0x002DE8A6 | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x002DE963 | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x002E38CC | `Az iPod, az iTunes programb` | Known | iTunes database |
| 0x002E3CE5 | `t az iTunes programhoz, ` | Known | iTunes database |
| 0x002E3DB5 | `t az iTunes programhoz ` | Known | iTunes database |
| 0x002E3E63 | `phez, hogy az iTunes feloldja a z` | Known | iTunes database |
| 0x002E3F71 | `ljon az iTunes haszn` | Known | iTunes database |
| 0x002E4895 | `s az iTunes futtat` | Known | iTunes database |
| 0x002E5307 | `lja azokat az iTunes programmal.` | Known | iTunes database |
| 0x002E53DE | `lja azokat az iTunes programmal.` | Known | iTunes database |
| 0x002EA140 | ` memorizzare e visualizzare informazioni importanti da iTunes o vCards. ` | Known | iTunes database |
| 0x002EA500 | `Collega iPod a iTunes e reinstalla il gioco.` | Known | iTunes database |
| 0x002EA5A4 | `Collega  iPod a iTunes ed esegui il download dell'ultima versione.` | Known | iTunes database |
| 0x002EA61C | `Se dimentichi la combinazione, collega iPod al computer e iTunes sar` | Known | iTunes database |
| 0x002EA6BC | `Le foto importate non possono visualizzarsi in TV. Trasferisci le foto sul compu` | Known | iTunes database |
| 0x002EAEF1 | ` troppo vecchio per funzionare con questo iPod. Collega iPod al computer ed eseg` | Known | iTunes database |
| 0x002EB7E5 | ` troppo grande per essere visualizzato in iPod. Trasferisci le foto sul computer` | Known | iTunes database |
| 0x002EB88F | ` essere visualizzato in iPod. Trasferisci le foto sul computer e sincronizzale t` | Known | iTunes database |
| 0x002F102F | `iTunes ` | Known | iTunes database |
| 0x002F14C1 | `iTunes ` | Known | iTunes database |
| 0x002F15CD | `iTunes ` | Known | iTunes database |
| 0x002F16CD | `iTunes` | Known | iTunes database |
| 0x002F17F8 | `iTunes ` | Known | iTunes database |
| 0x002F225F | `iTunes ` | Known | iTunes database |
| 0x002F2DD3 | `iTunes ` | Known | iTunes database |
| 0x002F2E99 | `iTunes ` | Known | iTunes database |
| 0x002F82DA | ` iTunes ` | Known | iTunes database |
| 0x002F86B3 | ` iTunes` | Known | iTunes database |
| 0x002F877B | ` iTunes` | Known | iTunes database |
| 0x002F8852 | ` iTunes` | Known | iTunes database |
| 0x002F8935 | ` iTunes` | Known | iTunes database |
| 0x002F9223 | ` iTunes` | Known | iTunes database |
| 0x002F9BFA | ` iTunes` | Known | iTunes database |
| 0x002F9C9F | ` iTunes` | Known | iTunes database |
| 0x002FEF2D | `mporteerd uit iTunes of vCards. ` | Known | iTunes database |
| 0x002FF2BC | `Verbind de iPod met iTunes en installeer het spel opnieuw.` | Known | iTunes database |
| 0x002FF35C | `Verbind de iPod met iTunes en download de nieuwste versie.` | Known | iTunes database |
| 0x002FF3D0 | `Als u de combinatie bent vergeten, verbind iPod met uw computer en iTunes zal he` | Known | iTunes database |
| 0x002FF481 | `mporteerde foto's op tv niet mogelijk. Kopieer foto's naar de computer en synchr` | Known | iTunes database |
| 0x002FFD18 | `%s is te oud om op deze iPod te worden gebruikt. Sluit de iPod aan op de compute` | Known | iTunes database |
| 0x003005F3 | `mporteerde foto's naar de computer en synchroniseer ze met iTunes voor weergave ` | Known | iTunes database |
| | *...and 73 more* | | |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002BEC60 | `Radio-Region` | Known | FM Radio |
| 0x002C1238 | `Radio-Region` | Known | FM Radio |
| 0x004B4D38 | `Radio Region` | Known | FM Radio |
| 0x004B746C | `Radio Region` | Known | FM Radio |
| 0x004B7A38 | `Radio Settings` | Known | FM Radio |

---

## 16. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017A394 | `FireWireGUID` | Known | FireWire interface |
| 0x0017A3A4 | `FireWireVersion` | Known | FireWire interface |
| 0x0017A990 | `FireWire` | Known | FireWire interface |
| 0x002B3137 | `es FireWire nen` | Known | FireWire interface |
| 0x002B50B0 | `FireWire p` | Known | FireWire interface |
| 0x002B9600 | `FireWire-forbindelser underst` | Known | FireWire interface |
| 0x002BB320 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x002BFEA8 | `FireWire wird nicht unterst` | Known | FireWire interface |
| 0x002C1F2E | `ber FireWire verbunden` | Known | FireWire interface |
| 0x002C7AC6 | ` FireWire. ` | Known | FireWire interface |
| 0x002CB066 | ` FireWire` | Known | FireWire interface |
| 0x002CF7B9 | `de canciones o archivos no son posibles con FireWire :` | Known | FireWire interface |
| 0x002D16E8 | `FireWire conectado` | Known | FireWire interface |
| 0x002D5D18 | `FireWire-tiedonsiirtoa ei tueta. Siirt` | Known | FireWire interface |
| 0x002D7A18 | `FireWire liitetty` | Known | FireWire interface |
| 0x002DCCE5 | `s via FireWire : connectez l` | Known | FireWire interface |
| 0x002DEF08 | `FireWire Connect` | Known | FireWire interface |
| 0x002E3714 | `A FireWire kapcsolat nem t` | Known | FireWire interface |
| 0x002E5924 | `FireWire csatlakozik` | Known | FireWire interface |
| 0x002E9F8C | `Connessioni di dati via FireWire non sono supportate. Per trasferire brani o dat` | Known | FireWire interface |
| 0x002EBDDC | `FireWire connesso` | Known | FireWire interface |
| 0x002F0DAC | `FireWire ` | Known | FireWire interface |
| 0x002F34C8 | `FireWire ` | Known | FireWire interface |
| 0x002F80E8 | `FireWire ` | Known | FireWire interface |
| 0x002FA1C4 | `FireWire ` | Known | FireWire interface |
| 0x002FED7A | `ren via FireWire, maar alleen via de meegeleverde USB-kabel.` | Known | FireWire interface |
| 0x00300B88 | `FireWire aangesloten` | Known | FireWire interface |
| 0x003050CF | `ring via FireWire st` | Known | FireWire interface |
| 0x00306D70 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x0030B47F | `czenie FireWire nie jest wspierane. By przes` | Known | FireWire interface |
| 0x0030D3CF | `czony przez Firewire` | Known | FireWire interface |
| 0x00311ACF | `es FireWire n` | Known | FireWire interface |
| 0x003139D0 | `FireWire ligado` | Known | FireWire interface |
| 0x00319395 | ` FireWire ` | Known | FireWire interface |
| 0x0031C7F7 | ` FireWire` | Known | FireWire interface |
| 0x00320E24 | `FireWire-` | Known | FireWire interface |
| 0x00322B40 | `FireWire anslutet` | Known | FireWire interface |
| 0x0032717C | `FireWire ba` | Known | FireWire interface |
| 0x00329088 | `FireWire Ba` | Known | FireWire interface |
| 0x0032D851 | ` FireWire ` | Known | FireWire interface |
| 0x0032F494 | `FireWire ` | Known | FireWire interface |
| 0x00333CA5 | ` FireWire ` | Known | FireWire interface |
| 0x003359CC | `FireWire ` | Known | FireWire interface |
| 0x004B6408 | `FireWire connections are not supported. To transfer songs, connect the USB cable` | Known | FireWire interface |
| 0x004B80B0 | `FireWire Connected` | Known | FireWire interface |

---

## 17. USB

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x006CC3C1 | `USBCompositeDevice1.6` | Known | USB interface |
| 0x006CC419 | `USBCompositeDevice1.6` | Known | USB interface |

---

## 18. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F5400 | `LCD Module could not be determined.` | Known | Hardware interface |
| 0x0017AE38 | `ForcedDiskMode` | Known | Hardware interface |
| 0x00217D40 | `Enter Disk Mode` | Known | Hardware interface |
| 0x00217D50 | `Exit Disk Mode` | Known | Hardware interface |
| 0x004B63FC | `Disk Mode` | Known | Hardware interface |
| 0x004FFE50 | `I2C write Error` | Known | Hardware interface |
| 0x004FFE64 | `I2C read Error %02x` | Known | Hardware interface |
| 0x00504B24 | `Sub-LCD` | Known | Hardware interface |
| 0x0067CB1A | `Sub-LCD` | Known | Hardware interface |
| 0x0067D563 | `MonoHope-LCD` | Known | Hardware interface |
| 0x00681979 | `OCSP_RESPID` | Known | Hardware interface |
| 0x00A8FC2A | `Sub-LCDRegularFONTLAB30:TTEXPORTSub-LCDRegularSub-LCD` | Known | Hardware interface |
| 0x00AAFA96 | `MonoHope-LCDRegularFONTLAB30:TTEXPORTMonoHope-LCDRegularMonoHope-LCD` | Known | Hardware interface |
| 0x00AB5A6E | `MonoHope-LCDRegularFONTLAB30:TTEXPORTMonoHope-LCDRegularMonoHope-LCD` | Known | Hardware interface |

---

## 19. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0014BF5C | `PowerManager` | Known | Power management |
| 0x0017A96C | `PowerInformation` | Known | Power management |
| 0x00217DAC | `Begin Charging` | Known | Power management |
| 0x00217DBC | `Stop Charging` | Known | Power management |
| 0x00261498 | `USBPowerSense` | Known | Power management |
| 0x00261558 | `PCFPowerMgr` | Known | Power management |
| 0x004B5FA4 | `Charging` | Known | Power management |
| 0x004B80FC | `Low Battery` | Known | Power management |

---

## 20. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002B7EB8 | `Alarmer` | Known | UI element |
| 0x002BB264 | `Alarmer` | Known | UI element |
| 0x002CDEAC | `Calendario` | Known | UI element |
| 0x002CDEB8 | `Calendarios` | Known | UI element |
| 0x002CDEC4 | `Calendarios` | Known | UI element |
| 0x002CDF00 | `Alarmas` | Known | UI element |
| 0x002CEA30 | `Calendario` | Known | UI element |
| 0x002CEA3C | `Calendarios` | Known | UI element |
| 0x002CFC74 | `Alarma` | Known | UI element |
| 0x002D0890 | `Alarma` | Known | UI element |
| 0x002D0900 | `Alarma` | Known | UI element |
| 0x002D0C9A | `Calendario` | Known | UI element |
| 0x002D0F38 | `Alarma` | Known | UI element |
| 0x002D15EC | `Alarma` | Known | UI element |
| 0x002D1630 | `Alarmas` | Known | UI element |
| 0x002DB2A0 | `Alarmes` | Known | UI element |
| 0x002DD1CC | `Alarme` | Known | UI element |
| 0x002DDED0 | `Alarme` | Known | UI element |
| 0x002DDF38 | `Alarme` | Known | UI element |
| 0x002DE620 | `Alarme` | Known | UI element |
| 0x002DEDC8 | `Alarme` | Known | UI element |
| 0x002DEE30 | `Alarmes` | Known | UI element |
| 0x002E86E4 | `Calendario` | Known | UI element |
| 0x002E86F0 | `Calendari` | Known | UI element |
| 0x002E86FC | `Calendari` | Known | UI element |
| 0x002E9220 | `Calendario` | Known | UI element |
| 0x002E922C | `Calendari` | Known | UI element |
| 0x002EB3B3 | `Calendario` | Known | UI element |
| 0x00303978 | `Alarmer` | Known | UI element |
| 0x00306658 | `Alarmtidspunkt` | Known | UI element |
| 0x00306CAC | `Alarmer` | Known | UI element |
| 0x00309B94 | `Alarmy` | Known | UI element |
| 0x00309F54 | `Gotowe` | Known | UI element |
| 0x0030A138 | `Gotowe` | Known | UI element |
| 0x0030D2EC | `Alarmy` | Known | UI element |
| 0x00310210 | `Alarmes` | Known | UI element |
| 0x00311F50 | `Alarme` | Known | UI element |
| 0x00312B00 | `Alarme` | Known | UI element |
| 0x00313910 | `Alarmes` | Known | UI element |
| 0x00322428 | `Alarmtid` | Known | UI element |
| 0x0032591C | `Alarmlar` | Known | UI element |
| 0x003288E8 | `Alarm Zaman` | Known | UI element |
| 0x00328FC8 | `Alarmlar` | Known | UI element |
| 0x004B4778 | `Calendar` | Known | UI element |
| 0x004B4784 | `Calendars` | Known | UI element |
| 0x004B4790 | `Calendars` | Known | UI element |
| 0x004B47C4 | `Alarms` | Known | UI element |
| 0x004B526C | `Calendar` | Known | UI element |
| 0x004B5278 | `Calendars` | Known | UI element |
| 0x004B73F4 | `Alarm Clock` | Known | UI element |
| 0x004B771E | `Calendar` | Known | UI element |
| 0x004B7990 | `Alarm Time` | Known | UI element |
| 0x004B799C | `Alarm Clock` | Known | UI element |
| 0x004B7FA4 | `Alarm Clock` | Known | UI element |
| 0x004B8004 | `Alarms` | Known | UI element |
| 0x004B8204 | `GotoBackToIdleCommand` | Known | UI element |
| 0x00503A84 | `TCalendarCntlr_Alarm` | Known | UI element |
| 0x0067CAE0 | `Calendars/` | Known | UI element |
| 0x0067CAFB | `Calendars` | Known | UI element |

---

## 21. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0007DA90 | `Settings` | Known | Menu item |
| 0x002B229C | `Podcasts` | Known | Menu item |
| 0x002B23FC | `Podcasts` | Known | Menu item |
| 0x002B471A | `Podcasts` | Known | Menu item |
| 0x002B5020 | `Podcasts` | Known | Menu item |
| 0x002B8808 | `Podcasts` | Known | Menu item |
| 0x002B8958 | `Podcasts` | Known | Menu item |
| 0x002BAA6D | `Podcasts` | Known | Menu item |
| 0x002BB290 | `Podcasts` | Known | Menu item |
| 0x002BF080 | `Podcasts` | Known | Menu item |
| 0x002BF180 | `Extras` | Known | Menu item |
| 0x002BF1B0 | `Videos` | Known | Menu item |
| 0x002BF1E8 | `Podcasts` | Known | Menu item |
| 0x002C0E64 | `Videos` | Known | Menu item |
| 0x002C153C | `Extras` | Known | Menu item |
| 0x002C1544 | `Videos` | Known | Menu item |
| 0x002C15A7 | `Podcasts` | Known | Menu item |
| 0x002C1868 | `Videos` | Known | Menu item |
| 0x002C1E98 | `Podcasts` | Known | Menu item |
| 0x002C1F94 | `Extras` | Known | Menu item |
| 0x002C6128 | `Podcasts` | Known | Menu item |
| 0x002C63E8 | `Podcasts` | Known | Menu item |
| 0x002C9F3D | `Podcasts` | Known | Menu item |
| 0x002CAF40 | `Podcasts` | Known | Menu item |
| 0x002CE960 | `Podcasts` | Known | Menu item |
| 0x002CEA5C | `Extras` | Known | Menu item |
| 0x002CEAC0 | `Podcasts` | Known | Menu item |
| 0x002D0CF0 | `Extras` | Known | Menu item |
| 0x002D0D62 | `Podcasts` | Known | Menu item |
| 0x002D1658 | `Podcasts` | Known | Menu item |
| 0x002D174C | `Extras` | Known | Menu item |
| 0x002DBDA8 | `Podcasts` | Known | Menu item |
| 0x002DBDF0 | `Albums` | Known | Menu item |
| 0x002DBE08 | `Genres` | Known | Menu item |
| 0x002DBE48 | `Photos` | Known | Menu item |
| 0x002DBEC8 | `Extras` | Known | Menu item |
| 0x002DBF38 | `Podcasts` | Known | Menu item |
| 0x002DC040 | `Albums` | Known | Menu item |
| 0x002DD62C | `Photos` | Known | Menu item |
| 0x002DD6D8 | `Photos` | Known | Menu item |
| 0x002DDBB8 | `Photos` | Known | Menu item |
| 0x002DE38C | `Extras` | Known | Menu item |
| 0x002DE3C0 | `Photos` | Known | Menu item |
| 0x002DE3EE | `Genres` | Known | Menu item |
| 0x002DE412 | `Podcasts` | Known | Menu item |
| 0x002DE446 | `Albums` | Known | Menu item |
| 0x002DE5C4 | `Genres` | Known | Menu item |
| 0x002DE5D8 | `Albums` | Known | Menu item |
| 0x002DE98C | `Photos` | Known | Menu item |
| 0x002DEE44 | `Genres` | Known | Menu item |
| 0x002DEE58 | `Podcasts` | Known | Menu item |
| 0x002DEE74 | `Albums` | Known | Menu item |
| 0x002DEF90 | `Extras` | Known | Menu item |
| 0x002FDF34 | `Podcasts` | Known | Menu item |
| 0x002FDF78 | `Albums` | Known | Menu item |
| 0x002FDF8C | `Genres` | Known | Menu item |
| 0x002FE0A0 | `Podcasts` | Known | Menu item |
| 0x002FE178 | `Albums` | Known | Menu item |
| 0x00300237 | `Genres` | Known | Menu item |
| 0x00300257 | `Podcasts` | Known | Menu item |
| 0x0030027F | `Albums` | Known | Menu item |
| 0x003003C0 | `Genres` | Known | Menu item |
| 0x003003D0 | `Albums` | Known | Menu item |
| 0x00300ADC | `Genres` | Known | Menu item |
| 0x00300AF0 | `Podcasts` | Known | Menu item |
| 0x00300B08 | `Albums` | Known | Menu item |
| 0x00310C1C | `Podcasts` | Known | Menu item |
| 0x00310D40 | `Extras` | Known | Menu item |
| 0x00310DA4 | `Podcasts` | Known | Menu item |
| 0x00312F94 | `Extras` | Known | Menu item |
| 0x0031300E | `Podcasts` | Known | Menu item |
| 0x00313948 | `Podcasts` | Known | Menu item |
| 0x00313A1C | `Extras` | Known | Menu item |
| 0x00335940 | `Podcasts` | Known | Menu item |
| 0x004B50B8 | `Podcasts` | Known | Menu item |
| 0x004B51C8 | `Now Playing` | Known | Menu item |
| 0x004B51D4 | `Artists` | Known | Menu item |
| 0x004B51EC | `Albums` | Known | Menu item |
| 0x004B5204 | `Genres` | Known | Menu item |
| 0x004B520C | `Composers` | Known | Menu item |
| 0x004B5238 | `Photos` | Known | Menu item |
| 0x004B52A0 | `Extras` | Known | Menu item |
| 0x004B52A8 | `Playlists` | Known | Menu item |
| 0x004B52B4 | `Audiobooks` | Known | Menu item |
| 0x004B52C8 | `Videos` | Known | Menu item |
| 0x004B52F4 | `Shuffle Songs` | Known | Menu item |
| 0x004B5304 | `Podcasts` | Known | Menu item |
| 0x004B53C4 | `Albums` | Known | Menu item |
| 0x004B63C8 | `Now Playing` | Known | Menu item |
| 0x004B6478 | `Audiobooks` | Known | Menu item |
| 0x004B7120 | `Photos` | Known | Menu item |
| 0x004B7128 | `Videos` | Known | Menu item |
| 0x004B76DC | `Shuffle Songs` | Known | Menu item |
| 0x004B7768 | `Extras` | Known | Menu item |
| 0x004B7770 | `Videos` | Known | Menu item |
| 0x004B778C | `Photos` | Known | Menu item |
| 0x004B77A6 | `Composers` | Known | Menu item |
| 0x004B77B6 | `Genres` | Known | Menu item |
| 0x004B77C6 | `Audiobooks` | Known | Menu item |
| 0x004B77DA | `Podcasts` | Known | Menu item |
| | *...and 24 more* | | |

---

## 22. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00097CD0 | `iPod_Control` | Filesystem Path | |
| 0x00097CFC | `iPod_Control\Device` | Filesystem Path | |
| 0x000A4528 | `iPod_Control\Device` | Filesystem Path | |
| 0x000A5F08 | `iPod_Control` | Filesystem Path | |
| 0x000A6548 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x000B8BE0 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000B8C20 | `System_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x000BB708 | `iPod_Control\Music\` | Filesystem Path | |
| 0x000BFCF4 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x000BFE74 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000E2E94 | `iPod_Control/%s%s%s` | Filesystem Path | |
| 0x000E2EA8 | `iPod_Control/%s/%s%s%s` | Filesystem Path | |
| 0x000EFCC8 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000F763C | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000F9088 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000F9184 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001AB9AC | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001AC340 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path | |
| 0x001AC368 | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001AC4D4 | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001D7564 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D77F0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D78A8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D7A04 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D7B24 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D7BF4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D7D54 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D7E3C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D7EF8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D7FA8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D809C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8140 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D81F4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D82B0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D83E4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8554 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8618 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D86C8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8804 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D88D0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D899C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8A64 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8B08 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8BD0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8C80 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8D30 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8DF8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8EC4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8F80 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D9030 | `iPod_Control\Device\` | Filesystem Path | |

---

## 23. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002B2598 | `Acoustic` | EQ Preset | |
| 0x002B25A4 | `Bass Booster` | EQ Preset | |
| 0x002B25C4 | `Classical` | EQ Preset | |
| 0x002B25E0 | `Electronic` | EQ Preset | |
| 0x002B25F4 | `Hip Hop` | EQ Preset | |
| 0x002B260C | `Loudness` | EQ Preset | |
| 0x002B2618 | `Lounge` | EQ Preset | |
| 0x002B263C | `Small Speakers` | EQ Preset | |
| 0x002B264C | `Spoken Word` | EQ Preset | |
| 0x002B2658 | `Treble Booster` | EQ Preset | |
| 0x002B2678 | `Vocal Booster` | EQ Preset | |
| 0x002B8AA4 | `Acoustic` | EQ Preset | |
| 0x002B8AB0 | `Bass Booster` | EQ Preset | |
| 0x002B8AD0 | `Classical` | EQ Preset | |
| 0x002B8AEC | `Electronic` | EQ Preset | |
| 0x002B8B00 | `Hip Hop` | EQ Preset | |
| 0x002B8B18 | `Loudness` | EQ Preset | |
| 0x002B8B24 | `Lounge` | EQ Preset | |
| 0x002B8B48 | `Small Speakers` | EQ Preset | |
| 0x002B8B58 | `Spoken Word` | EQ Preset | |
| 0x002B8B64 | `Treble Booster` | EQ Preset | |
| 0x002B8B84 | `Vocal Booster` | EQ Preset | |
| 0x002BF348 | `Acoustic` | EQ Preset | |
| 0x002BF388 | `Electronic` | EQ Preset | |
| 0x002BF39C | `Hip Hop` | EQ Preset | |
| 0x002BF3B4 | `Loudness` | EQ Preset | |
| 0x002C66D4 | `Hip Hop` | EQ Preset | |
| 0x002C66EC | `Loudness` | EQ Preset | |
| 0x002C66F8 | `Lounge` | EQ Preset | |
| 0x002CEC9C | `Hip Hop` | EQ Preset | |
| 0x002CECAC | `Latina` | EQ Preset | |
| 0x002CECB4 | `Loudness` | EQ Preset | |
| 0x002CECC0 | `Lounge` | EQ Preset | |
| 0x002D5210 | `Lounge` | EQ Preset | |
| 0x002DC134 | `Hip Hop` | EQ Preset | |
| 0x002DC164 | `Lounge` | EQ Preset | |
| 0x002E9454 | `Hip Hop` | EQ Preset | |
| 0x002E9464 | `Latina` | EQ Preset | |
| 0x002E946C | `Loudness` | EQ Preset | |
| 0x002E9478 | `Lounge` | EQ Preset | |
| 0x002EFC34 | `Acoustic` | EQ Preset | |
| 0x002EFC40 | `Bass Booster` | EQ Preset | |
| 0x002EFC60 | `Classical` | EQ Preset | |
| 0x002EFC7C | `Electronic` | EQ Preset | |
| 0x002EFC90 | `Hip Hop` | EQ Preset | |
| 0x002EFCA8 | `Loudness` | EQ Preset | |
| 0x002EFCB4 | `Lounge` | EQ Preset | |
| 0x002EFCD8 | `Small Speakers` | EQ Preset | |
| 0x002EFCE8 | `Spoken Word` | EQ Preset | |
| 0x002EFCF4 | `Treble Booster` | EQ Preset | |
| 0x002EFD14 | `Vocal Booster` | EQ Preset | |
| 0x002F7204 | `Acoustic` | EQ Preset | |
| 0x002F7210 | `Bass Booster` | EQ Preset | |
| 0x002F7230 | `Classical` | EQ Preset | |
| 0x002F724C | `Electronic` | EQ Preset | |
| 0x002F7260 | `Hip Hop` | EQ Preset | |
| 0x002F7278 | `Loudness` | EQ Preset | |
| 0x002F7284 | `Lounge` | EQ Preset | |
| 0x002F72A8 | `Small Speakers` | EQ Preset | |
| 0x002F72B8 | `Spoken Word` | EQ Preset | |
| 0x002F72C4 | `Treble Booster` | EQ Preset | |
| 0x002F72E4 | `Vocal Booster` | EQ Preset | |
| 0x002FE264 | `Loudness` | EQ Preset | |
| 0x002FE270 | `Lounge` | EQ Preset | |
| 0x003045C8 | `Latino` | EQ Preset | |
| 0x003045D0 | `Loudness` | EQ Preset | |
| 0x003045DC | `Lounge` | EQ Preset | |
| 0x0030A90C | `Hip Hop` | EQ Preset | |
| 0x0030A940 | `Lounge` | EQ Preset | |
| 0x00310F90 | `Hip Hop` | EQ Preset | |
| 0x00310FA0 | `Latina` | EQ Preset | |
| 0x00310FA8 | `Loudness` | EQ Preset | |
| 0x00310FB4 | `Lounge` | EQ Preset | |
| 0x003202DC | `Acoustic` | EQ Preset | |
| 0x003202E8 | `Bass Booster` | EQ Preset | |
| 0x00320308 | `Classical` | EQ Preset | |
| 0x00320324 | `Electronic` | EQ Preset | |
| 0x00320338 | `Hip Hop` | EQ Preset | |
| 0x00320350 | `Loudness` | EQ Preset | |
| 0x0032035C | `Lounge` | EQ Preset | |
| 0x00320380 | `Small Speakers` | EQ Preset | |
| 0x00320390 | `Spoken Word` | EQ Preset | |
| 0x0032039C | `Treble Booster` | EQ Preset | |
| 0x003203BC | `Vocal Booster` | EQ Preset | |
| 0x00326688 | `Hip Hop` | EQ Preset | |
| 0x0032669C | `Loudness` | EQ Preset | |
| 0x003266A8 | `Lounge` | EQ Preset | |
| 0x0032CA9C | `Acoustic` | EQ Preset | |
| 0x0032CAA8 | `Bass Booster` | EQ Preset | |
| 0x0032CAC8 | `Classical` | EQ Preset | |
| 0x0032CAE4 | `Electronic` | EQ Preset | |
| 0x0032CAF8 | `Hip Hop` | EQ Preset | |
| 0x0032CB10 | `Loudness` | EQ Preset | |
| 0x0032CB1C | `Lounge` | EQ Preset | |
| 0x0032CB40 | `Small Speakers` | EQ Preset | |
| 0x0032CB50 | `Spoken Word` | EQ Preset | |
| 0x0032CB5C | `Treble Booster` | EQ Preset | |
| 0x0032CB7C | `Vocal Booster` | EQ Preset | |
| 0x00332F14 | `Acoustic` | EQ Preset | |
| 0x00332F20 | `Bass Booster` | EQ Preset | |
| 0x00332F40 | `Classical` | EQ Preset | |
| 0x00332F5C | `Electronic` | EQ Preset | |
| 0x00332F70 | `Hip Hop` | EQ Preset | |
| 0x00332F88 | `Loudness` | EQ Preset | |
| 0x00332F94 | `Lounge` | EQ Preset | |
| 0x00332FB8 | `Small Speakers` | EQ Preset | |
| 0x00332FC8 | `Spoken Word` | EQ Preset | |
| 0x00332FD4 | `Treble Booster` | EQ Preset | |
| 0x00332FF4 | `Vocal Booster` | EQ Preset | |
| 0x004B551C | `Acoustic` | EQ Preset | |
| 0x004B5528 | `Bass Booster` | EQ Preset | |
| 0x004B5548 | `Classical` | EQ Preset | |
| 0x004B5564 | `Electronic` | EQ Preset | |
| 0x004B5578 | `Hip Hop` | EQ Preset | |
| 0x004B5590 | `Loudness` | EQ Preset | |
| 0x004B559C | `Lounge` | EQ Preset | |
| 0x004B55C0 | `Small Speakers` | EQ Preset | |
| 0x004B55D0 | `Spoken Word` | EQ Preset | |
| 0x004B55DC | `Treble Booster` | EQ Preset | |
| 0x004B55FC | `Vocal Booster` | EQ Preset | |

---

## 24. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00106194 | `Error-SDriver` | Diagnostic | |
| 0x001061A4 | `Error-AClient` | Diagnostic | |
| 0x00106C80 | `Root Hub Driver Internal Error unused case in hub handler` | Diagnostic | |
| 0x00106CBC | `Root hub Error Calling Add Device` | Diagnostic | |
| 0x0010B93C | `Error inside %s` | Diagnostic | |
| 0x001422E0 | `%s Error in file %s.` | Diagnostic | |
| 0x00270370 | `Error inside %s` | Diagnostic | |
| 0x00270400 | `Error inside %s` | Diagnostic | |
| 0x00270484 | `Error inside %s` | Diagnostic | |
| 0x0027093C | `Error inside %s` | Diagnostic | |
| 0x00270A00 | `Error inside %s` | Diagnostic | |
| 0x00270ACC | `Error inside %s` | Diagnostic | |
| 0x00270D88 | `Error inside %s` | Diagnostic | |
| 0x00270F78 | `Error inside %s` | Diagnostic | |
| 0x00270FDC | `Error inside %s` | Diagnostic | |
| 0x00271110 | `Error inside %s` | Diagnostic | |
| 0x00271168 | `Error inside %s` | Diagnostic | |
| 0x002711B8 | `Error inside %s` | Diagnostic | |
| 0x00271288 | `Error inside %s` | Diagnostic | |
| 0x002712D8 | `Error inside %s` | Diagnostic | |
| 0x002716CC | `Error inside %s` | Diagnostic | |
| 0x0027173C | `Error inside %s` | Diagnostic | |
| 0x00271BBC | `Error inside %s` | Diagnostic | |
| 0x00272024 | `Error inside %s` | Diagnostic | |
| 0x00272228 | `Error inside %s` | Diagnostic | |
| 0x00272298 | `Error inside %s` | Diagnostic | |
| 0x00272398 | `Error inside %s` | Diagnostic | |
| 0x002724A0 | `Error inside %s` | Diagnostic | |
| 0x00272514 | `Error inside %s` | Diagnostic | |
| 0x00272560 | `Error inside %s` | Diagnostic | |

---
