# iPod 5.5G (Video Enhanced 80GB) - RetailOS 1.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.3 |
| **IPSW** | iPod_25.1.3.ipsw |
| **Device** | iPod 5.5G (Video Enhanced 80GB) (2006, Click Wheel, Search, Brighter Display) |
| **UpdaterFamilyID** | 25 |
| **Binary Size** | 13,903,872 bytes (13.26 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 13,903,872 bytes |
| **Total Strings (>=6)** | 30,181 |
| **Function Prologues** | 22,824 (ARM: 12,890, Thumb: 9,934) |
| **SoC** | PortalPlayer PP5022C |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Encrypted** | No |
| **SHA-256** | `7830d1345aa2313db154e06ae93f2b5961e1cb04e8edeaae27dadc303e0d9fb3` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016F904 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x00216934 | `Channel UnitTests` | Hidden | Developer Tool |

---

## 2. Controllers (TSilver/TC Classes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00B28878 | `TCC_Current_Task_Pointer` | Known | UI controller |
| 0x00B28891 | `TCC_Delete_HISR` | Known | UI controller |
| 0x00B288B1 | `TCC_Relinquish` | Known | UI controller |
| 0x00B288CF | `TCC_Resume_Service` | Known | UI controller |
| 0x00B288E2 | `TCC_Task_Sleep` | Known | UI controller |
| 0x00B28904 | `TCF_Task_Information` | Known | UI controller |
| 0x00B28919 | `TCS_Change_Preemption` | Known | UI controller |
| 0x00B2892F | `TCS_Change_Priority` | Known | UI controller |
| 0x00B28943 | `TCT_Activate_HISR` | Known | UI controller |
| 0x00B28955 | `TCT_Control_Interrupts` | Known | UI controller |
| 0x00B2896C | `TCT_Local_Control_Interrupts` | Known | UI controller |
| 0x00B9011F | `TCC_Resume_Service` | Known | UI controller |
| 0x00B902A2 | `TCC_Delete_HISR` | Known | UI controller |
| 0x00B902F1 | `TCT_Activate_HISR` | Known | UI controller |
| 0x00B9031B | `TCT_Control_Interrupts` | Known | UI controller |
| 0x00B904E5 | `TCC_Current_Task_Pointer` | Known | UI controller |
| 0x00B904FE | `TCS_Change_Priority` | Known | UI controller |
| 0x00B9056C | `TCC_Task_Sleep` | Known | UI controller |
| 0x00B906A5 | `TCF_Task_Information` | Known | UI controller |
| 0x00B9EC1D | `TCC_Relinquish` | Known | UI controller |
| 0x00B9ED33 | `TCT_Local_Control_Interrupts` | Known | UI controller |
| 0x00B9EDD3 | `TCC_Task_Sleep` | Known | UI controller |
| 0x00B9EEA1 | `TCS_Change_Preemption` | Known | UI controller |

---

## 3. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00099174 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x000C9B20 | `USB Secondary Interrupt Task` | Known | RTOS task thread |
| 0x000E334C | `ICAPTPCameraIOTask` | Known | RTOS task thread |
| 0x0011DBA8 | `USBStatusTask` | Known | RTOS task thread |
| 0x0011DBC4 | `USBTaskTimeTask` | Known | RTOS task thread |
| 0x00169A2C | `VCUpdateTask` | Known | RTOS task thread |
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
| 0x00B28868 | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B288A1 | `TCC_Delete_Task` | Known | RTOS task thread |
| 0x00B288C0 | `TCC_Reset_Task` | Known | RTOS task thread |
| 0x00B288F1 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B90110 | `TCC_Reset_Task` | Known | RTOS task thread |
| 0x00B90147 | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B902B2 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B902C5 | `TCC_Delete_Task` | Known | RTOS task thread |
| 0x00B9EC0D | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B9EC57 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B9EC6A | `TCC_Delete_Task` | Known | RTOS task thread |

---

## 4. Logging Channels

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

## 5. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017A888 | `AudioCodecs` | Known | Audio system |
| 0x0017B928 | `VideoCodecs` | Known | Audio system |
| 0x002BAAFA | `.net codec i dette produkt bruges i henhold til en licensaftale fra VoiceAge Cor` | Known | Audio system |
| 0x002C15C9 | `.net Codec in diesem Produkt wird in Lizenz der VoiceAge Corporation verwendet. ` | Known | Audio system |
| 0x002C9774 | `.net codec ` | Known | Audio system |
| 0x002DE2E8 | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x002E4D81 | `.net codec a VoiceAge Coporation c` | Known | Audio system |
| 0x002EB492 | ` utilizzato su licenza da VoiceAge Corporation. Il codec ACELP` | Known | Audio system |
| 0x002F9742 | `.net codec` | Known | Audio system |
| 0x00300274 | `.net-codec in dit product wordt gebruikt in licentie van VoiceAge Corporation. G` | Known | Audio system |
| 0x00312F5D | `a da VoiceAge Corporation. O codec ACELP` | Known | Audio system |
| 0x00328709 | `.net codec'i VoiceAge Corporation lisans` | Known | Audio system |
| 0x004B78B5 | `.net codec in this product is used under license from VoiceAge Corporation. Port` | Known | Audio system |
| 0x0068048E | `msCodeCom` | Known | Audio system |
| 0x00B8FB9D | `codec_string` | Known | Audio system |
| 0x00B8FBAA | `codec_name` | Known | Audio system |
| 0x00B9E4C9 | `codec_string` | Known | Audio system |
| 0x00B9E4D6 | `codec_name` | Known | Audio system |

---

## 6. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017A964 | `Audible` | Known | Audible audiobook format |
| 0x002B451D | ` Audible v` | Known | Audible audiobook format |
| 0x002B456F | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x002B4585 | ` 2002 Audible, Inc. V` | Known | Audible audiobook format |
| 0x002BA9A8 | `Audible-softwaren i dette produkt bruges i henhold til en licensaftale fra Audib` | Known | Audible audiobook format |
| 0x002BAA08 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x002C1484 | `Die Audible Software in diesem Produkt wird in Lizenz der Audible verwendet. Cop` | Known | Audible audiobook format |
| 0x002C14DD | ` 2002 Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x002C9563 | ` Audible ` | Known | Audible audiobook format |
| 0x002C95C0 | ` Audible. ` | Known | Audible audiobook format |
| 0x002C95F6 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002D0BA8 | `El software Audible incluido en este producto se usa bajo licencia de Audible. C` | Known | Audible audiobook format |
| 0x002D0C03 | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x002D7096 | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright ` | Known | Audible audiobook format |
| 0x002D70D0 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x002DE1D8 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x002DE222 | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x002DE237 | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x002E4C42 | ` Audible szoftver az Audible licence alatt van haszn` | Known | Audible audiobook format |
| 0x002E4C8C | ` 2002, Audible, Inc. Minden jog fenntartva.` | Known | Audible audiobook format |
| 0x002EB3C8 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x002EB3F1 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x002EB420 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x002F2541 | ` Audible ` | Known | Audible audiobook format |
| 0x002F2562 | `Audible ` | Known | Audible audiobook format |
| 0x002F25BB | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002F95F3 | ` Audible ` | Known | Audible audiobook format |
| 0x002F960E | ` Audible` | Known | Audible audiobook format |
| 0x002F9652 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x0030012C | `De Audible-software in dit product wordt gebruikt in licentie van Audible. Copyr` | Known | Audible audiobook format |
| 0x00300183 | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x0030645C | `Audible-programvaren i dette produktet brukes under lisens fra Audible. Copyrigh` | Known | Audible audiobook format |
| 0x003064B0 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0030C864 | `Oprogramowanie Audible w tym produkcie jest wykorzystywane na podstawie licencji` | Known | Audible audiobook format |
| 0x0030C8D0 | ` 2002 Audible, Inc. Wszystkie prawa zastrze` | Known | Audible audiobook format |
| 0x00312E4C | `O software Audible ` | Known | Audible audiobook format |
| 0x00312E82 | `a da Audible. Copyright ` | Known | Audible audiobook format |
| 0x00312E9C | ` 2002 da Audible, Inc. Reservados todos os direitos.` | Known | Audible audiobook format |
| 0x0031AEAD | ` Audible ` | Known | Audible audiobook format |
| 0x0031AEFF | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x0031AF15 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x003221AC | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x003221DB | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x003221F2 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x003285C0 | `ndeki Audible yaz` | Known | Audible audiobook format |
| 0x003285D9 | ` Audible lisans` | Known | Audible audiobook format |
| 0x0032860E | ` 2002, Audible, Inc. T` | Known | Audible audiobook format |
| 0x0032EBDB | ` Audible ` | Known | Audible audiobook format |
| 0x0032EBED | ` Audible ` | Known | Audible audiobook format |
| 0x0032EC11 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x0033504C | `Audible ` | Known | Audible audiobook format |
| 0x00335060 | ` Audible ` | Known | Audible audiobook format |
| 0x0033508A | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x004B777C | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x004B77D1 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |

---

## 7. Audio/Codec - Apple Lossless

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017A938 | `AppleLossless` | Known | Apple Lossless codec |
| 0x002E519C | `l alacsony.` | Known | Apple Lossless codec |

---

## 8. Audio/Codec - AAC

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00B39860 | `AACDEC  VLL ` | Known | AAC codec |
| 0x00B42FA8 | `MPEG4 AAC LC Decoder` | Known | AAC codec |
| 0x00B4540D | `AACDecoderGetMem` | Known | AAC codec |
| 0x00B4541E | `AACDecoderInit` | Known | AAC codec |
| 0x00B4542D | `AACDecoderGetConfig` | Known | AAC codec |
| 0x00B45441 | `AACDecoderSetConfig` | Known | AAC codec |
| 0x00B45455 | `AACHeaderDecode` | Known | AAC codec |
| 0x00B45465 | `AACDecode` | Known | AAC codec |
| 0x00B4546F | `AACDecoderInit_Ittiam` | Known | AAC codec |
| 0x00B45485 | `AACDecoderGetConfig_Ittiam` | Known | AAC codec |
| 0x00B454A0 | `AACDecoderSetConfig_Ittiam` | Known | AAC codec |
| 0x00B454BB | `AACHeaderDecode_Ittiam` | Known | AAC codec |
| 0x00B454D2 | `AACDecode_Ittiam` | Known | AAC codec |

---

## 9. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002B4762 | ` zvuku MPEG Layer-3 byla poskytnuta spole` | Known | MP3 codec |
| 0x002B478D | `nostmi Fraunhofer IIS a` | Known | MP3 codec |
| 0x002BABA4 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunhofer IIS og THOMSON multi` | Known | MP3 codec |
| 0x002C1687 | `r MPEG Layer-3 wurde lizenziert von Fraunhofer IIS und THOMSON Multimedia.` | Known | MP3 codec |
| 0x002C98BB | ` MPEG Layer-3 ` | Known | MP3 codec |
| 0x002C98F9 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x002D0DA1 | `n de audio MPEG Layer-3 utilizada bajo licencia de Fraunhofer IIS y THOMSON mult` | Known | MP3 codec |
| 0x002D723C | `MPEG Layer-3 -` | Known | MP3 codec |
| 0x002D724E | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | MP3 codec |
| 0x002DE3F4 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | MP3 codec |
| 0x002E4E14 | `Az MPEG Layer-3 hangk` | Known | MP3 codec |
| 0x002E4E3C | `gia a Fraunhofer IIS ` | Known | MP3 codec |
| 0x002EB568 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da Fraunhofer IIS e THOMSON` | Known | MP3 codec |
| 0x002F2770 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002F27BC | `Fraunhofer IIS ` | Known | MP3 codec |
| 0x002F97DC | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002F9803 | ` Fraunhofer IIS` | Known | MP3 codec |
| 0x00300310 | `Technologie voor codering van MPEG Layer-3-audio in licentie van Fraunhofer IIS ` | Known | MP3 codec |
| 0x0030662C | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fraunhofer IIS og THOMSON m` | Known | MP3 codec |
| 0x0030CA78 | `Technologia kodowania audio MPEG Layer-3 licencjonowana od Fraunhofer IIS oraz T` | Known | MP3 codec |
| 0x0031303E | `udio MPEG Layer-3 licenciada pela Fraunhofer IIS e THOMSON multimedia.` | Known | MP3 codec |
| 0x0031B134 | `MPEG Layer-3: ` | Known | MP3 codec |
| 0x0031B18D | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x0032238C | `Ljudkodningstekniken MPEG Layer-3 ` | Known | MP3 codec |
| 0x003223C2 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | MP3 codec |
| 0x00328798 | `MPEG Layer-3 ses kodlama teknolojisi Fraunhofer IIS ve THOMSON multimedia'dan li` | Known | MP3 codec |
| 0x0032ED68 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x0032ED8A | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x003351E8 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x0033520D | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x004B7948 | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |

---

## 10. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E7819 | `;=1sinffniscpsap@-` | Known | DRM system |
| 0x0017A85C | `AppleDRMVersion` | Known | DRM system |
| 0x0017A8FC | `AppleDRM` | Known | DRM system |
| 0x0017B93C | `AppleVideoDRM` | Known | DRM system |
| 0x00180AF8 | `drmsp608mp4aesdsmp4v` | Known | DRM system |
| 0x001CDD5C | `tkhdedtselstmdiamdhdminfstblstsdstcoco64stscstszsttsstssdrmidrmsx` | Known | DRM system |
| 0x0067D58E | `DRMLevel` | Known | DRM system |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002AE43C | `games_RO` | Known | Game system |
| 0x002AE448 | `gamedata_RW` | Known | Game system |
| 0x002AE464 | `gamedata_ShareRW` | Known | Game system |

---

## 12. Photo System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00098974 | `Photo Database` | Known | Photo system |
| 0x000B9604 | `Photos\Photo Database` | Known | Photo system |
| 0x000C0E2C | `Photo Database` | Known | Photo system |
| 0x0019BABC | `23iUPhoto Database` | Known | Photo system |
| 0x0019DB14 | `Photo Database` | Known | Photo system |
| 0x0019DE78 | `Photo Database` | Known | Photo system |
| 0x0019E124 | `Photo Import Database` | Known | Photo system |
| 0x002187E0 | `Photo Database Size` | Known | Photo system |

---

## 13. Video System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00B53EE0 | `H.264 Video Decoder` | Known | Video system |

---

## 14. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000B95F8 | `iTunesDB` | Known | iTunes database |
| 0x001C31C4 | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x0020E8C4 | `iTunes Image DB` | Known | iTunes database |
| 0x002AE1FC | `iTunesDB` | Known | iTunes database |
| 0x002B3CFD | ` z iTunes nebo vCards. ` | Known | iTunes database |
| 0x002B4017 | `ipojte iPod k iTunes a instalujte hru znovu.` | Known | iTunes database |
| 0x002B40A7 | `ipojte iPod k iTunes a zkop` | Known | iTunes database |
| 0x002B413F | `i a program iTunes jej odemkne.` | Known | iTunes database |
| 0x002B4265 | `m iTunes.` | Known | iTunes database |
| 0x002B4BCD | `m iTunes.` | Known | iTunes database |
| 0x002B5543 | `es iTunes.` | Known | iTunes database |
| 0x002B55E9 | `es iTunes.` | Known | iTunes database |
| 0x002BA19C | `iPod kan opbevare og vise kontaktoplysninger importeret fra iTunes eller vCards.` | Known | iTunes database |
| 0x002BA524 | `Slut iPod til iTunes, og installer spillet igen.` | Known | iTunes database |
| 0x002BA5BC | `Slut iPod til iTunes, og overf` | Known | iTunes database |
| 0x002BA653 | `slutte iPod til computeren, hvorefter iTunes l` | Known | iTunes database |
| 0x002BA708 | `r fotografier til computeren, og synkroniser via iTunes for at vise dem p` | Known | iTunes database |
| 0x002BAF50 | `%s er for gammel til denne iPod. Slut iPod til computeren, og start iTunes for a` | Known | iTunes database |
| 0x002BB7C6 | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x002BB855 | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x002C0ADD | `hlen Sie Ihren iPod in iTunes aus der Liste ` | Known | iTunes database |
| 0x002C0B4B | `iTunes` | Known | iTunes database |
| 0x002C0F90 | `Verbinden Sie Ihren iPod mit iTunes und installieren Sie das Spiel erneut.` | Known | iTunes database |
| 0x002C1054 | `Verbinden Sie Ihren iPod mit iTunes und laden Sie die aktuelle Version.` | Known | iTunes database |
| 0x002C10FB | `en Sie Ihren iPod an Ihren Computer an und iTunes deaktiviert die Anzeigensperre` | Known | iTunes database |
| 0x002C11A4 | `Importierte Fotos werden nicht auf dem TV angezeigt. Senden Sie sie erst an den ` | Known | iTunes database |
| 0x002C1A89 | `en Sie den iPod an den Computer an und starten Sie iTunes, um %s auf die aktuell` | Known | iTunes database |
| 0x002C23A1 | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x002C2442 | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x002C87A8 | ` iTunes ` | Known | iTunes database |
| 0x002C8CF1 | ` iTunes ` | Known | iTunes database |
| 0x002C8E19 | ` iTunes ` | Known | iTunes database |
| 0x002C8F42 | ` iTunes ` | Known | iTunes database |
| 0x002C90D7 | ` iTunes ` | Known | iTunes database |
| 0x002C9FCA | ` iTunes ` | Known | iTunes database |
| 0x002CB0DA | ` iTunes ` | Known | iTunes database |
| 0x002CB203 | ` iTunes ` | Known | iTunes database |
| 0x002D0379 | `n importada de iTunes o de tarjetas virtuales (vCards). ` | Known | iTunes database |
| 0x002D0704 | `Conecte el iPod a iTunes y reinstale el juego.` | Known | iTunes database |
| 0x002D079C | `Conecte el iPod a iTunes y descargue la versi` | Known | iTunes database |
| 0x002D0824 | `n, conecte el iPod al ordenador y iTunes lo desbloquear` | Known | iTunes database |
| 0x002D0900 | `celas con iTunes para verlas en la TV.` | Known | iTunes database |
| 0x002D1198 | `%s es demasiado antiguo para ejecutarse en este iPod. Conecte el iPod al ordenad` | Known | iTunes database |
| 0x002D1AD8 | `Esta foto es demasiado grande para mostrarla en el iPod. Transfiera las fotos im` | Known | iTunes database |
| 0x002D1B7C | `Este formato de foto no puede visualizarse en el iPod. Transfiera las fotos impo` | Known | iTunes database |
| 0x002D68D8 | ` iTunesista tai vCardeina tuotua tietoa. ` | Known | iTunes database |
| 0x002D6C3E | ` iPod iTunesiin ja asenna peli uudelleen.` | Known | iTunes database |
| 0x002D6CCE | ` iPod iTunesiin ja hae uusin versio.` | Known | iTunes database |
| 0x002D6D43 | `tietokoneeseen, niin iTunes avaa lukituksen.` | Known | iTunes database |
| 0x002D6DE0 | ` kuvat tietokoneelle ja synkronoi ne iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002D762E | ` %s uusimpaan versioon avaamalla iTunes.` | Known | iTunes database |
| 0x002D7EA6 | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002D7F2D | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002DD8E2 | `iTunes ou de vCards. ` | Known | iTunes database |
| 0x002DDC88 | `Connectez votre iPod avec iTunes et r` | Known | iTunes database |
| 0x002DDD30 | `Connectez votre iPod avec iTunes et t` | Known | iTunes database |
| 0x002DDDF1 | ` votre ordinateur et iTunes le d` | Known | iTunes database |
| 0x002DDED4 | `rez-les sur votre ordinateur puis synchronisez-les avec iTunes.` | Known | iTunes database |
| 0x002DE7FC | `ordinateur et lancez iTunes pour mettre ` | Known | iTunes database |
| 0x002DF2A6 | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x002DF363 | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x002E42CC | `Az iPod, az iTunes programb` | Known | iTunes database |
| 0x002E46E5 | `t az iTunes programhoz, ` | Known | iTunes database |
| 0x002E47B5 | `t az iTunes programhoz ` | Known | iTunes database |
| 0x002E4863 | `phez, hogy az iTunes feloldja a z` | Known | iTunes database |
| 0x002E4971 | `ljon az iTunes haszn` | Known | iTunes database |
| 0x002E5295 | `s az iTunes futtat` | Known | iTunes database |
| 0x002E5D07 | `lja azokat az iTunes programmal.` | Known | iTunes database |
| 0x002E5DDE | `lja azokat az iTunes programmal.` | Known | iTunes database |
| 0x002EAB40 | ` memorizzare e visualizzare informazioni importanti da iTunes o vCards. ` | Known | iTunes database |
| 0x002EAF00 | `Collega iPod a iTunes e reinstalla il gioco.` | Known | iTunes database |
| 0x002EAFA4 | `Collega  iPod a iTunes ed esegui il download dell'ultima versione.` | Known | iTunes database |
| 0x002EB01C | `Se dimentichi la combinazione, collega iPod al computer e iTunes sar` | Known | iTunes database |
| 0x002EB0BC | `Le foto importate non possono visualizzarsi in TV. Trasferisci le foto sul compu` | Known | iTunes database |
| 0x002EB8F1 | ` troppo vecchio per funzionare con questo iPod. Collega iPod al computer ed eseg` | Known | iTunes database |
| 0x002EC1E5 | ` troppo grande per essere visualizzato in iPod. Trasferisci le foto sul computer` | Known | iTunes database |
| 0x002EC28F | ` essere visualizzato in iPod. Trasferisci le foto sul computer e sincronizzale t` | Known | iTunes database |
| 0x002F1A2F | `iTunes ` | Known | iTunes database |
| 0x002F1EC1 | `iTunes ` | Known | iTunes database |
| 0x002F1FCD | `iTunes ` | Known | iTunes database |
| 0x002F20CD | `iTunes` | Known | iTunes database |
| 0x002F21F8 | `iTunes ` | Known | iTunes database |
| 0x002F2C5F | `iTunes ` | Known | iTunes database |
| 0x002F37D3 | `iTunes ` | Known | iTunes database |
| 0x002F3899 | `iTunes ` | Known | iTunes database |
| 0x002F8CDA | ` iTunes ` | Known | iTunes database |
| 0x002F90B3 | ` iTunes` | Known | iTunes database |
| 0x002F917B | ` iTunes` | Known | iTunes database |
| 0x002F9252 | ` iTunes` | Known | iTunes database |
| 0x002F9335 | ` iTunes` | Known | iTunes database |
| 0x002F9C23 | ` iTunes` | Known | iTunes database |
| 0x002FA5FA | ` iTunes` | Known | iTunes database |
| 0x002FA69F | ` iTunes` | Known | iTunes database |
| 0x002FF92D | `mporteerd uit iTunes of vCards. ` | Known | iTunes database |
| 0x002FFCBC | `Verbind de iPod met iTunes en installeer het spel opnieuw.` | Known | iTunes database |
| 0x002FFD5C | `Verbind de iPod met iTunes en download de nieuwste versie.` | Known | iTunes database |
| 0x002FFDD0 | `Als u de combinatie bent vergeten, verbind iPod met uw computer en iTunes zal he` | Known | iTunes database |
| 0x002FFE81 | `mporteerde foto's op tv niet mogelijk. Kopieer foto's naar de computer en synchr` | Known | iTunes database |
| 0x00300718 | `%s is te oud om op deze iPod te worden gebruikt. Sluit de iPod aan op de compute` | Known | iTunes database |
| 0x00300FF3 | `mporteerde foto's naar de computer en synchroniseer ze met iTunes voor weergave ` | Known | iTunes database |
| | *...and 73 more* | | |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002BF660 | `Radio-Region` | Known | FM Radio |
| 0x002C1C38 | `Radio-Region` | Known | FM Radio |
| 0x004B5738 | `Radio Region` | Known | FM Radio |
| 0x004B7E6C | `Radio Region` | Known | FM Radio |
| 0x004B8438 | `Radio Settings` | Known | FM Radio |

---

## 16. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017AD94 | `FireWireGUID` | Known | FireWire interface |
| 0x0017ADA4 | `FireWireVersion` | Known | FireWire interface |
| 0x0017B390 | `FireWire` | Known | FireWire interface |
| 0x002B3B37 | `es FireWire nen` | Known | FireWire interface |
| 0x002B5AB0 | `FireWire p` | Known | FireWire interface |
| 0x002BA000 | `FireWire-forbindelser underst` | Known | FireWire interface |
| 0x002BBD20 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x002C08A8 | `FireWire wird nicht unterst` | Known | FireWire interface |
| 0x002C292E | `ber FireWire verbunden` | Known | FireWire interface |
| 0x002C84C6 | ` FireWire. ` | Known | FireWire interface |
| 0x002CBA66 | ` FireWire` | Known | FireWire interface |
| 0x002D01B9 | `de canciones o archivos no son posibles con FireWire :` | Known | FireWire interface |
| 0x002D20E8 | `FireWire conectado` | Known | FireWire interface |
| 0x002D6718 | `FireWire-tiedonsiirtoa ei tueta. Siirt` | Known | FireWire interface |
| 0x002D8418 | `FireWire liitetty` | Known | FireWire interface |
| 0x002DD6E5 | `s via FireWire : connectez l` | Known | FireWire interface |
| 0x002DF908 | `FireWire Connect` | Known | FireWire interface |
| 0x002E4114 | `A FireWire kapcsolat nem t` | Known | FireWire interface |
| 0x002E6324 | `FireWire csatlakozik` | Known | FireWire interface |
| 0x002EA98C | `Connessioni di dati via FireWire non sono supportate. Per trasferire brani o dat` | Known | FireWire interface |
| 0x002EC7DC | `FireWire connesso` | Known | FireWire interface |
| 0x002F17AC | `FireWire ` | Known | FireWire interface |
| 0x002F3EC8 | `FireWire ` | Known | FireWire interface |
| 0x002F8AE8 | `FireWire ` | Known | FireWire interface |
| 0x002FABC4 | `FireWire ` | Known | FireWire interface |
| 0x002FF77A | `ren via FireWire, maar alleen via de meegeleverde USB-kabel.` | Known | FireWire interface |
| 0x00301588 | `FireWire aangesloten` | Known | FireWire interface |
| 0x00305ACF | `ring via FireWire st` | Known | FireWire interface |
| 0x00307770 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x0030BE7F | `czenie FireWire nie jest wspierane. By przes` | Known | FireWire interface |
| 0x0030DDCF | `czony przez Firewire` | Known | FireWire interface |
| 0x003124CF | `es FireWire n` | Known | FireWire interface |
| 0x003143D0 | `FireWire ligado` | Known | FireWire interface |
| 0x00319D95 | ` FireWire ` | Known | FireWire interface |
| 0x0031D1F7 | ` FireWire` | Known | FireWire interface |
| 0x00321824 | `FireWire-` | Known | FireWire interface |
| 0x00323540 | `FireWire anslutet` | Known | FireWire interface |
| 0x00327B7C | `FireWire ba` | Known | FireWire interface |
| 0x00329A88 | `FireWire Ba` | Known | FireWire interface |
| 0x0032E251 | ` FireWire ` | Known | FireWire interface |
| 0x0032FE94 | `FireWire ` | Known | FireWire interface |
| 0x003346A5 | ` FireWire ` | Known | FireWire interface |
| 0x003363CC | `FireWire ` | Known | FireWire interface |
| 0x004B6E08 | `FireWire connections are not supported. To transfer songs, connect the USB cable` | Known | FireWire interface |
| 0x004B8AB0 | `FireWire Connected` | Known | FireWire interface |

---

## 17. USB

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x006CCDC1 | `USBCompositeDevice1.6` | Known | USB interface |
| 0x006CCE19 | `USBCompositeDevice1.6` | Known | USB interface |

---

## 18. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F5E00 | `LCD Module could not be determined.` | Known | Hardware interface |
| 0x0017B838 | `ForcedDiskMode` | Known | Hardware interface |
| 0x00218740 | `Enter Disk Mode` | Known | Hardware interface |
| 0x00218750 | `Exit Disk Mode` | Known | Hardware interface |
| 0x004B6DFC | `Disk Mode` | Known | Hardware interface |
| 0x00500850 | `I2C write Error` | Known | Hardware interface |
| 0x00500864 | `I2C read Error %02x` | Known | Hardware interface |
| 0x00505524 | `Sub-LCD` | Known | Hardware interface |
| 0x0067D51A | `Sub-LCD` | Known | Hardware interface |
| 0x0067DF63 | `MonoHope-LCD` | Known | Hardware interface |
| 0x00682379 | `OCSP_RESPID` | Known | Hardware interface |
| 0x00A9182A | `Sub-LCDRegularFONTLAB30:TTEXPORTSub-LCDRegularSub-LCD` | Known | Hardware interface |
| 0x00AB1696 | `MonoHope-LCDRegularFONTLAB30:TTEXPORTMonoHope-LCDRegularMonoHope-LCD` | Known | Hardware interface |
| 0x00AB766E | `MonoHope-LCDRegularFONTLAB30:TTEXPORTMonoHope-LCDRegularMonoHope-LCD` | Known | Hardware interface |

---

## 19. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0014C95C | `PowerManager` | Known | Power management |
| 0x0017B36C | `PowerInformation` | Known | Power management |
| 0x002187AC | `Begin Charging` | Known | Power management |
| 0x002187BC | `Stop Charging` | Known | Power management |
| 0x00261E98 | `USBPowerSense` | Known | Power management |
| 0x00261F58 | `PCFPowerMgr` | Known | Power management |
| 0x004B69A4 | `Charging` | Known | Power management |
| 0x004B8AFC | `Low Battery` | Known | Power management |

---

## 20. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002B88B8 | `Alarmer` | Known | UI element |
| 0x002BBC64 | `Alarmer` | Known | UI element |
| 0x002CE8AC | `Calendario` | Known | UI element |
| 0x002CE8B8 | `Calendarios` | Known | UI element |
| 0x002CE8C4 | `Calendarios` | Known | UI element |
| 0x002CE900 | `Alarmas` | Known | UI element |
| 0x002CF430 | `Calendario` | Known | UI element |
| 0x002CF43C | `Calendarios` | Known | UI element |
| 0x002D0674 | `Alarma` | Known | UI element |
| 0x002D1290 | `Alarma` | Known | UI element |
| 0x002D1300 | `Alarma` | Known | UI element |
| 0x002D169A | `Calendario` | Known | UI element |
| 0x002D1938 | `Alarma` | Known | UI element |
| 0x002D1FEC | `Alarma` | Known | UI element |
| 0x002D2030 | `Alarmas` | Known | UI element |
| 0x002DBCA0 | `Alarmes` | Known | UI element |
| 0x002DDBCC | `Alarme` | Known | UI element |
| 0x002DE8D0 | `Alarme` | Known | UI element |
| 0x002DE938 | `Alarme` | Known | UI element |
| 0x002DF020 | `Alarme` | Known | UI element |
| 0x002DF7C8 | `Alarme` | Known | UI element |
| 0x002DF830 | `Alarmes` | Known | UI element |
| 0x002E90E4 | `Calendario` | Known | UI element |
| 0x002E90F0 | `Calendari` | Known | UI element |
| 0x002E90FC | `Calendari` | Known | UI element |
| 0x002E9C20 | `Calendario` | Known | UI element |
| 0x002E9C2C | `Calendari` | Known | UI element |
| 0x002EBDB3 | `Calendario` | Known | UI element |
| 0x00304378 | `Alarmer` | Known | UI element |
| 0x00307058 | `Alarmtidspunkt` | Known | UI element |
| 0x003076AC | `Alarmer` | Known | UI element |
| 0x0030A594 | `Alarmy` | Known | UI element |
| 0x0030A954 | `Gotowe` | Known | UI element |
| 0x0030AB38 | `Gotowe` | Known | UI element |
| 0x0030DCEC | `Alarmy` | Known | UI element |
| 0x00310C10 | `Alarmes` | Known | UI element |
| 0x00312950 | `Alarme` | Known | UI element |
| 0x00313500 | `Alarme` | Known | UI element |
| 0x00314310 | `Alarmes` | Known | UI element |
| 0x00322E28 | `Alarmtid` | Known | UI element |
| 0x0032631C | `Alarmlar` | Known | UI element |
| 0x003292E8 | `Alarm Zaman` | Known | UI element |
| 0x003299C8 | `Alarmlar` | Known | UI element |
| 0x004B5178 | `Calendar` | Known | UI element |
| 0x004B5184 | `Calendars` | Known | UI element |
| 0x004B5190 | `Calendars` | Known | UI element |
| 0x004B51C4 | `Alarms` | Known | UI element |
| 0x004B5C6C | `Calendar` | Known | UI element |
| 0x004B5C78 | `Calendars` | Known | UI element |
| 0x004B7DF4 | `Alarm Clock` | Known | UI element |
| 0x004B811E | `Calendar` | Known | UI element |
| 0x004B8390 | `Alarm Time` | Known | UI element |
| 0x004B839C | `Alarm Clock` | Known | UI element |
| 0x004B89A4 | `Alarm Clock` | Known | UI element |
| 0x004B8A04 | `Alarms` | Known | UI element |
| 0x004B8C04 | `GotoBackToIdleCommand` | Known | UI element |
| 0x00504484 | `TCalendarCntlr_Alarm` | Known | UI element |
| 0x0067D4E0 | `Calendars/` | Known | UI element |
| 0x0067D4FB | `Calendars` | Known | UI element |

---

## 21. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0007E490 | `Settings` | Known | Menu item |
| 0x002B2C9C | `Podcasts` | Known | Menu item |
| 0x002B2DFC | `Podcasts` | Known | Menu item |
| 0x002B511A | `Podcasts` | Known | Menu item |
| 0x002B5A20 | `Podcasts` | Known | Menu item |
| 0x002B9208 | `Podcasts` | Known | Menu item |
| 0x002B9358 | `Podcasts` | Known | Menu item |
| 0x002BB46D | `Podcasts` | Known | Menu item |
| 0x002BBC90 | `Podcasts` | Known | Menu item |
| 0x002BFA80 | `Podcasts` | Known | Menu item |
| 0x002BFB80 | `Extras` | Known | Menu item |
| 0x002BFBB0 | `Videos` | Known | Menu item |
| 0x002BFBE8 | `Podcasts` | Known | Menu item |
| 0x002C1864 | `Videos` | Known | Menu item |
| 0x002C1F3C | `Extras` | Known | Menu item |
| 0x002C1F44 | `Videos` | Known | Menu item |
| 0x002C1FA7 | `Podcasts` | Known | Menu item |
| 0x002C2268 | `Videos` | Known | Menu item |
| 0x002C2898 | `Podcasts` | Known | Menu item |
| 0x002C2994 | `Extras` | Known | Menu item |
| 0x002C6B28 | `Podcasts` | Known | Menu item |
| 0x002C6DE8 | `Podcasts` | Known | Menu item |
| 0x002CA93D | `Podcasts` | Known | Menu item |
| 0x002CB940 | `Podcasts` | Known | Menu item |
| 0x002CF360 | `Podcasts` | Known | Menu item |
| 0x002CF45C | `Extras` | Known | Menu item |
| 0x002CF4C0 | `Podcasts` | Known | Menu item |
| 0x002D16F0 | `Extras` | Known | Menu item |
| 0x002D1762 | `Podcasts` | Known | Menu item |
| 0x002D2058 | `Podcasts` | Known | Menu item |
| 0x002D214C | `Extras` | Known | Menu item |
| 0x002DC7A8 | `Podcasts` | Known | Menu item |
| 0x002DC7F0 | `Albums` | Known | Menu item |
| 0x002DC808 | `Genres` | Known | Menu item |
| 0x002DC848 | `Photos` | Known | Menu item |
| 0x002DC8C8 | `Extras` | Known | Menu item |
| 0x002DC938 | `Podcasts` | Known | Menu item |
| 0x002DCA40 | `Albums` | Known | Menu item |
| 0x002DE02C | `Photos` | Known | Menu item |
| 0x002DE0D8 | `Photos` | Known | Menu item |
| 0x002DE5B8 | `Photos` | Known | Menu item |
| 0x002DED8C | `Extras` | Known | Menu item |
| 0x002DEDC0 | `Photos` | Known | Menu item |
| 0x002DEDEE | `Genres` | Known | Menu item |
| 0x002DEE12 | `Podcasts` | Known | Menu item |
| 0x002DEE46 | `Albums` | Known | Menu item |
| 0x002DEFC4 | `Genres` | Known | Menu item |
| 0x002DEFD8 | `Albums` | Known | Menu item |
| 0x002DF38C | `Photos` | Known | Menu item |
| 0x002DF844 | `Genres` | Known | Menu item |
| 0x002DF858 | `Podcasts` | Known | Menu item |
| 0x002DF874 | `Albums` | Known | Menu item |
| 0x002DF990 | `Extras` | Known | Menu item |
| 0x002FE934 | `Podcasts` | Known | Menu item |
| 0x002FE978 | `Albums` | Known | Menu item |
| 0x002FE98C | `Genres` | Known | Menu item |
| 0x002FEAA0 | `Podcasts` | Known | Menu item |
| 0x002FEB78 | `Albums` | Known | Menu item |
| 0x00300C37 | `Genres` | Known | Menu item |
| 0x00300C57 | `Podcasts` | Known | Menu item |
| 0x00300C7F | `Albums` | Known | Menu item |
| 0x00300DC0 | `Genres` | Known | Menu item |
| 0x00300DD0 | `Albums` | Known | Menu item |
| 0x003014DC | `Genres` | Known | Menu item |
| 0x003014F0 | `Podcasts` | Known | Menu item |
| 0x00301508 | `Albums` | Known | Menu item |
| 0x0031161C | `Podcasts` | Known | Menu item |
| 0x00311740 | `Extras` | Known | Menu item |
| 0x003117A4 | `Podcasts` | Known | Menu item |
| 0x00313994 | `Extras` | Known | Menu item |
| 0x00313A0E | `Podcasts` | Known | Menu item |
| 0x00314348 | `Podcasts` | Known | Menu item |
| 0x0031441C | `Extras` | Known | Menu item |
| 0x00336340 | `Podcasts` | Known | Menu item |
| 0x004B5AB8 | `Podcasts` | Known | Menu item |
| 0x004B5BC8 | `Now Playing` | Known | Menu item |
| 0x004B5BD4 | `Artists` | Known | Menu item |
| 0x004B5BEC | `Albums` | Known | Menu item |
| 0x004B5C04 | `Genres` | Known | Menu item |
| 0x004B5C0C | `Composers` | Known | Menu item |
| 0x004B5C38 | `Photos` | Known | Menu item |
| 0x004B5CA0 | `Extras` | Known | Menu item |
| 0x004B5CA8 | `Playlists` | Known | Menu item |
| 0x004B5CB4 | `Audiobooks` | Known | Menu item |
| 0x004B5CC8 | `Videos` | Known | Menu item |
| 0x004B5CF4 | `Shuffle Songs` | Known | Menu item |
| 0x004B5D04 | `Podcasts` | Known | Menu item |
| 0x004B5DC4 | `Albums` | Known | Menu item |
| 0x004B6DC8 | `Now Playing` | Known | Menu item |
| 0x004B6E78 | `Audiobooks` | Known | Menu item |
| 0x004B7B20 | `Photos` | Known | Menu item |
| 0x004B7B28 | `Videos` | Known | Menu item |
| 0x004B80DC | `Shuffle Songs` | Known | Menu item |
| 0x004B8168 | `Extras` | Known | Menu item |
| 0x004B8170 | `Videos` | Known | Menu item |
| 0x004B818C | `Photos` | Known | Menu item |
| 0x004B81A6 | `Composers` | Known | Menu item |
| 0x004B81B6 | `Genres` | Known | Menu item |
| 0x004B81C6 | `Audiobooks` | Known | Menu item |
| 0x004B81DA | `Podcasts` | Known | Menu item |
| | *...and 24 more* | | |

---

## 22. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000986D0 | `iPod_Control` | Filesystem Path | |
| 0x000986FC | `iPod_Control\Device` | Filesystem Path | |
| 0x000A4F28 | `iPod_Control\Device` | Filesystem Path | |
| 0x000A6908 | `iPod_Control` | Filesystem Path | |
| 0x000A6F48 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x000B95E0 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000B9620 | `System_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x000BC108 | `iPod_Control\Music\` | Filesystem Path | |
| 0x000C06F4 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x000C0874 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000E3894 | `iPod_Control/%s%s%s` | Filesystem Path | |
| 0x000E38A8 | `iPod_Control/%s/%s%s%s` | Filesystem Path | |
| 0x000F06C8 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000F803C | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000F9A88 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000F9B84 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001AC3AC | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001ACD40 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path | |
| 0x001ACD68 | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001ACED4 | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001D7F64 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D81F0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D82A8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8404 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8524 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D85F4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8754 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D883C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D88F8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D89A8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8A9C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8B40 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8BF4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8CB0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8DE4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D8F54 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D9018 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D90C8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D9204 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D92D0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D939C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D9464 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D9508 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D95D0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D9680 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D9730 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D97F8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D98C4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D9980 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D9A30 | `iPod_Control\Device\` | Filesystem Path | |

---

## 23. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002B2F98 | `Acoustic` | EQ Preset | |
| 0x002B2FA4 | `Bass Booster` | EQ Preset | |
| 0x002B2FC4 | `Classical` | EQ Preset | |
| 0x002B2FE0 | `Electronic` | EQ Preset | |
| 0x002B2FF4 | `Hip Hop` | EQ Preset | |
| 0x002B300C | `Loudness` | EQ Preset | |
| 0x002B3018 | `Lounge` | EQ Preset | |
| 0x002B303C | `Small Speakers` | EQ Preset | |
| 0x002B304C | `Spoken Word` | EQ Preset | |
| 0x002B3058 | `Treble Booster` | EQ Preset | |
| 0x002B3078 | `Vocal Booster` | EQ Preset | |
| 0x002B94A4 | `Acoustic` | EQ Preset | |
| 0x002B94B0 | `Bass Booster` | EQ Preset | |
| 0x002B94D0 | `Classical` | EQ Preset | |
| 0x002B94EC | `Electronic` | EQ Preset | |
| 0x002B9500 | `Hip Hop` | EQ Preset | |
| 0x002B9518 | `Loudness` | EQ Preset | |
| 0x002B9524 | `Lounge` | EQ Preset | |
| 0x002B9548 | `Small Speakers` | EQ Preset | |
| 0x002B9558 | `Spoken Word` | EQ Preset | |
| 0x002B9564 | `Treble Booster` | EQ Preset | |
| 0x002B9584 | `Vocal Booster` | EQ Preset | |
| 0x002BFD48 | `Acoustic` | EQ Preset | |
| 0x002BFD88 | `Electronic` | EQ Preset | |
| 0x002BFD9C | `Hip Hop` | EQ Preset | |
| 0x002BFDB4 | `Loudness` | EQ Preset | |
| 0x002C70D4 | `Hip Hop` | EQ Preset | |
| 0x002C70EC | `Loudness` | EQ Preset | |
| 0x002C70F8 | `Lounge` | EQ Preset | |
| 0x002CF69C | `Hip Hop` | EQ Preset | |
| 0x002CF6AC | `Latina` | EQ Preset | |
| 0x002CF6B4 | `Loudness` | EQ Preset | |
| 0x002CF6C0 | `Lounge` | EQ Preset | |
| 0x002D5C10 | `Lounge` | EQ Preset | |
| 0x002DCB34 | `Hip Hop` | EQ Preset | |
| 0x002DCB64 | `Lounge` | EQ Preset | |
| 0x002E9E54 | `Hip Hop` | EQ Preset | |
| 0x002E9E64 | `Latina` | EQ Preset | |
| 0x002E9E6C | `Loudness` | EQ Preset | |
| 0x002E9E78 | `Lounge` | EQ Preset | |
| 0x002F0634 | `Acoustic` | EQ Preset | |
| 0x002F0640 | `Bass Booster` | EQ Preset | |
| 0x002F0660 | `Classical` | EQ Preset | |
| 0x002F067C | `Electronic` | EQ Preset | |
| 0x002F0690 | `Hip Hop` | EQ Preset | |
| 0x002F06A8 | `Loudness` | EQ Preset | |
| 0x002F06B4 | `Lounge` | EQ Preset | |
| 0x002F06D8 | `Small Speakers` | EQ Preset | |
| 0x002F06E8 | `Spoken Word` | EQ Preset | |
| 0x002F06F4 | `Treble Booster` | EQ Preset | |
| 0x002F0714 | `Vocal Booster` | EQ Preset | |
| 0x002F7C04 | `Acoustic` | EQ Preset | |
| 0x002F7C10 | `Bass Booster` | EQ Preset | |
| 0x002F7C30 | `Classical` | EQ Preset | |
| 0x002F7C4C | `Electronic` | EQ Preset | |
| 0x002F7C60 | `Hip Hop` | EQ Preset | |
| 0x002F7C78 | `Loudness` | EQ Preset | |
| 0x002F7C84 | `Lounge` | EQ Preset | |
| 0x002F7CA8 | `Small Speakers` | EQ Preset | |
| 0x002F7CB8 | `Spoken Word` | EQ Preset | |
| 0x002F7CC4 | `Treble Booster` | EQ Preset | |
| 0x002F7CE4 | `Vocal Booster` | EQ Preset | |
| 0x002FEC64 | `Loudness` | EQ Preset | |
| 0x002FEC70 | `Lounge` | EQ Preset | |
| 0x00304FC8 | `Latino` | EQ Preset | |
| 0x00304FD0 | `Loudness` | EQ Preset | |
| 0x00304FDC | `Lounge` | EQ Preset | |
| 0x0030B30C | `Hip Hop` | EQ Preset | |
| 0x0030B340 | `Lounge` | EQ Preset | |
| 0x00311990 | `Hip Hop` | EQ Preset | |
| 0x003119A0 | `Latina` | EQ Preset | |
| 0x003119A8 | `Loudness` | EQ Preset | |
| 0x003119B4 | `Lounge` | EQ Preset | |
| 0x00320CDC | `Acoustic` | EQ Preset | |
| 0x00320CE8 | `Bass Booster` | EQ Preset | |
| 0x00320D08 | `Classical` | EQ Preset | |
| 0x00320D24 | `Electronic` | EQ Preset | |
| 0x00320D38 | `Hip Hop` | EQ Preset | |
| 0x00320D50 | `Loudness` | EQ Preset | |
| 0x00320D5C | `Lounge` | EQ Preset | |
| 0x00320D80 | `Small Speakers` | EQ Preset | |
| 0x00320D90 | `Spoken Word` | EQ Preset | |
| 0x00320D9C | `Treble Booster` | EQ Preset | |
| 0x00320DBC | `Vocal Booster` | EQ Preset | |
| 0x00327088 | `Hip Hop` | EQ Preset | |
| 0x0032709C | `Loudness` | EQ Preset | |
| 0x003270A8 | `Lounge` | EQ Preset | |
| 0x0032D49C | `Acoustic` | EQ Preset | |
| 0x0032D4A8 | `Bass Booster` | EQ Preset | |
| 0x0032D4C8 | `Classical` | EQ Preset | |
| 0x0032D4E4 | `Electronic` | EQ Preset | |
| 0x0032D4F8 | `Hip Hop` | EQ Preset | |
| 0x0032D510 | `Loudness` | EQ Preset | |
| 0x0032D51C | `Lounge` | EQ Preset | |
| 0x0032D540 | `Small Speakers` | EQ Preset | |
| 0x0032D550 | `Spoken Word` | EQ Preset | |
| 0x0032D55C | `Treble Booster` | EQ Preset | |
| 0x0032D57C | `Vocal Booster` | EQ Preset | |
| 0x00333914 | `Acoustic` | EQ Preset | |
| 0x00333920 | `Bass Booster` | EQ Preset | |
| 0x00333940 | `Classical` | EQ Preset | |
| 0x0033395C | `Electronic` | EQ Preset | |
| 0x00333970 | `Hip Hop` | EQ Preset | |
| 0x00333988 | `Loudness` | EQ Preset | |
| 0x00333994 | `Lounge` | EQ Preset | |
| 0x003339B8 | `Small Speakers` | EQ Preset | |
| 0x003339C8 | `Spoken Word` | EQ Preset | |
| 0x003339D4 | `Treble Booster` | EQ Preset | |
| 0x003339F4 | `Vocal Booster` | EQ Preset | |
| 0x004B5F1C | `Acoustic` | EQ Preset | |
| 0x004B5F28 | `Bass Booster` | EQ Preset | |
| 0x004B5F48 | `Classical` | EQ Preset | |
| 0x004B5F64 | `Electronic` | EQ Preset | |
| 0x004B5F78 | `Hip Hop` | EQ Preset | |
| 0x004B5F90 | `Loudness` | EQ Preset | |
| 0x004B5F9C | `Lounge` | EQ Preset | |
| 0x004B5FC0 | `Small Speakers` | EQ Preset | |
| 0x004B5FD0 | `Spoken Word` | EQ Preset | |
| 0x004B5FDC | `Treble Booster` | EQ Preset | |
| 0x004B5FFC | `Vocal Booster` | EQ Preset | |

---

## 24. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00106B94 | `Error-SDriver` | Diagnostic | |
| 0x00106BA4 | `Error-AClient` | Diagnostic | |
| 0x00107680 | `Root Hub Driver Internal Error unused case in hub handler` | Diagnostic | |
| 0x001076BC | `Root hub Error Calling Add Device` | Diagnostic | |
| 0x0010C33C | `Error inside %s` | Diagnostic | |
| 0x00142CE0 | `%s Error in file %s.` | Diagnostic | |
| 0x00270D70 | `Error inside %s` | Diagnostic | |
| 0x00270E00 | `Error inside %s` | Diagnostic | |
| 0x00270E84 | `Error inside %s` | Diagnostic | |
| 0x0027133C | `Error inside %s` | Diagnostic | |
| 0x00271400 | `Error inside %s` | Diagnostic | |
| 0x002714CC | `Error inside %s` | Diagnostic | |
| 0x00271788 | `Error inside %s` | Diagnostic | |
| 0x00271978 | `Error inside %s` | Diagnostic | |
| 0x002719DC | `Error inside %s` | Diagnostic | |
| 0x00271B10 | `Error inside %s` | Diagnostic | |
| 0x00271B68 | `Error inside %s` | Diagnostic | |
| 0x00271BB8 | `Error inside %s` | Diagnostic | |
| 0x00271C88 | `Error inside %s` | Diagnostic | |
| 0x00271CD8 | `Error inside %s` | Diagnostic | |
| 0x002720CC | `Error inside %s` | Diagnostic | |
| 0x0027213C | `Error inside %s` | Diagnostic | |
| 0x002725BC | `Error inside %s` | Diagnostic | |
| 0x00272A24 | `Error inside %s` | Diagnostic | |
| 0x00272C28 | `Error inside %s` | Diagnostic | |
| 0x00272C98 | `Error inside %s` | Diagnostic | |
| 0x00272D98 | `Error inside %s` | Diagnostic | |
| 0x00272EA0 | `Error inside %s` | Diagnostic | |
| 0x00272F14 | `Error inside %s` | Diagnostic | |
| 0x00272F60 | `Error inside %s` | Diagnostic | |

---
