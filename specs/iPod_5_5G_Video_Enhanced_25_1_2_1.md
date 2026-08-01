# iPod 5.5G (Video Enhanced 80GB) - RetailOS 1.2.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.2.1 |
| **IPSW** | iPod_25.1.2.1.ipsw |
| **Device** | iPod 5.5G (Video Enhanced 80GB) (2006, Click Wheel, Search, Brighter Display) |
| **UpdaterFamilyID** | 25 |
| **Binary Size** | 13,844,480 bytes (13.20 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 13,844,480 bytes |
| **Total Strings (>=6)** | 29,898 |
| **Function Prologues** | 22,390 (ARM: 12,640, Thumb: 9,750) |
| **SoC** | PortalPlayer PP5022C |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Encrypted** | No |
| **SHA-256** | `60c5c5e66972da3334b0be38be4c1c0a62d4de035ee0704ede9221baeae70d2b` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016A4A4 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x0020D828 | `Channel UnitTests` | Hidden | Developer Tool |

---

## 2. Controllers (TSilver/TC Classes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00AD9078 | `TCC_Current_Task_Pointer` | Known | UI controller |
| 0x00AD9091 | `TCC_Delete_HISR` | Known | UI controller |
| 0x00AD90B1 | `TCC_Relinquish` | Known | UI controller |
| 0x00AD90CF | `TCC_Resume_Service` | Known | UI controller |
| 0x00AD90E2 | `TCC_Task_Sleep` | Known | UI controller |
| 0x00AD9104 | `TCF_Task_Information` | Known | UI controller |
| 0x00AD9119 | `TCS_Change_Preemption` | Known | UI controller |
| 0x00AD912F | `TCS_Change_Priority` | Known | UI controller |
| 0x00AD9143 | `TCT_Activate_HISR` | Known | UI controller |
| 0x00AD9155 | `TCT_Control_Interrupts` | Known | UI controller |
| 0x00AD916C | `TCT_Local_Control_Interrupts` | Known | UI controller |
| 0x00B4091F | `TCC_Resume_Service` | Known | UI controller |
| 0x00B40AA2 | `TCC_Delete_HISR` | Known | UI controller |
| 0x00B40AF1 | `TCT_Activate_HISR` | Known | UI controller |
| 0x00B40B1B | `TCT_Control_Interrupts` | Known | UI controller |
| 0x00B40CE5 | `TCC_Current_Task_Pointer` | Known | UI controller |
| 0x00B40CFE | `TCS_Change_Priority` | Known | UI controller |
| 0x00B40D6C | `TCC_Task_Sleep` | Known | UI controller |
| 0x00B40EA5 | `TCF_Task_Information` | Known | UI controller |
| 0x00B4F41D | `TCC_Relinquish` | Known | UI controller |
| 0x00B4F533 | `TCT_Local_Control_Interrupts` | Known | UI controller |
| 0x00B4F5D3 | `TCC_Task_Sleep` | Known | UI controller |
| 0x00B4F6A1 | `TCS_Change_Preemption` | Known | UI controller |

---

## 3. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009B000 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x000C97E8 | `USB Secondary Interrupt Task` | Known | RTOS task thread |
| 0x000E2218 | `ICAPTPCameraIOTask` | Known | RTOS task thread |
| 0x0011C6B4 | `USBStatusTask` | Known | RTOS task thread |
| 0x0011C6D0 | `USBTaskTimeTask` | Known | RTOS task thread |
| 0x001645E8 | `VCUpdateTask` | Known | RTOS task thread |
| 0x00170074 | `USBDeviceTask` | Known | RTOS task thread |
| 0x00176588 | `DiskReaderTask` | Known | RTOS task thread |
| 0x001857F8 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0018580C | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x00198A44 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x002076D8 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x00257DBC | `FirewireTask` | Known | RTOS task thread |
| 0x00257DD4 | `OptoTask` | Known | RTOS task thread |
| 0x00257DE4 | `SerialOptoTask` | Known | RTOS task thread |
| 0x00257DF8 | `BacklightTask` | Known | RTOS task thread |
| 0x00257E0C | `CNATask` | Known | RTOS task thread |
| 0x00257E2C | `DiskMgrTask` | Known | RTOS task thread |
| 0x00257E3C | `HoldSwitchTask` | Known | RTOS task thread |
| 0x00257E50 | `TopPlugTask` | Known | RTOS task thread |
| 0x00257E60 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x00257E74 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x00257E8C | `AccessoryDetectTask` | Known | RTOS task thread |
| 0x00257EB4 | `AlarmTask` | Known | RTOS task thread |
| 0x00257EC4 | `WatchdogTask` | Known | RTOS task thread |
| 0x00257F3C | `USBAudioTask` | Known | RTOS task thread |
| 0x002A4270 | `HostOSTask` | Known | RTOS task thread |
| 0x002A4E1C | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x005028B8 | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x005028D0 | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x005028E8 | `VideoDaisyTask` | Known | RTOS task thread |
| 0x00AD9068 | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00AD90A1 | `TCC_Delete_Task` | Known | RTOS task thread |
| 0x00AD90C0 | `TCC_Reset_Task` | Known | RTOS task thread |
| 0x00AD90F1 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B40910 | `TCC_Reset_Task` | Known | RTOS task thread |
| 0x00B40947 | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B40AB2 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B40AC5 | `TCC_Delete_Task` | Known | RTOS task thread |
| 0x00B4F40D | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B4F457 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B4F46A | `TCC_Delete_Task` | Known | RTOS task thread |

---

## 4. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0020D634 | `Channel Reserved` | Known | Logging channel |
| 0x0020D648 | `Channel AppBoot` | Known | Logging channel |
| 0x0020D658 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x0020D674 | `Channel PrefsWriting` | Known | Logging channel |
| 0x0020D68C | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x0020D6AC | `Channel PlayFromDisk` | Known | Logging channel |
| 0x0020D6C4 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x0020D6E0 | `Channel TestLogging` | Known | Logging channel |
| 0x0020D6F4 | `Channel AppFileLoading` | Known | Logging channel |
| 0x0020D70C | `Channel VCardReading` | Known | Logging channel |
| 0x0020D724 | `Channel LongSongScanning` | Known | Logging channel |
| 0x0020D798 | `Channel VoiceRecording` | Known | Logging channel |
| 0x0020D7B0 | `Channel PhotoImporting` | Known | Logging channel |
| 0x0020D7C8 | `Channel Notes` | Known | Logging channel |
| 0x0020D7D8 | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x0020D7F4 | `Channel DiskMode` | Known | Logging channel |
| 0x0020D808 | `Channel Firewire` | Known | Logging channel |
| 0x0020D81C | `Channel USB` | Known | Logging channel |
| 0x0020D83C | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x0020D854 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 5. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174E84 | `AudioCodecs` | Known | Audio system |
| 0x00175F10 | `VideoCodecs` | Known | Audio system |
| 0x002B146A | `.net codec i dette produkt bruges i henhold til en licensaftale fra VoiceAge Cor` | Known | Audio system |
| 0x002B82D9 | `.net Codec in diesem Produkt wird in Lizenz der VoiceAge Corporation verwendet. ` | Known | Audio system |
| 0x002C0B30 | `.net codec ` | Known | Audio system |
| 0x002D66E4 | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x002DD6C1 | `.net codec a VoiceAge Coporation c` | Known | Audio system |
| 0x002E42FE | ` utilizzato su licenza da VoiceAge Corporation. Il codec ACELP` | Known | Audio system |
| 0x002F3276 | `.net codec` | Known | Audio system |
| 0x002FA2E0 | `.net-codec in dit product wordt gebruikt in licentie van VoiceAge Corporation. G` | Known | Audio system |
| 0x0030DEF5 | `a da VoiceAge Corporation. O codec ACELP` | Known | Audio system |
| 0x003247B9 | `.net codec'i VoiceAge Corporation lisans` | Known | Audio system |
| 0x004B4F75 | `.net codec in this product is used under license from VoiceAge Corporation. Port` | Known | Audio system |
| 0x0067C39C | `msCodeCom` | Known | Audio system |
| 0x00B4039D | `codec_string` | Known | Audio system |
| 0x00B403AA | `codec_name` | Known | Audio system |
| 0x00B4ECC9 | `codec_string` | Known | Audio system |
| 0x00B4ECD6 | `codec_name` | Known | Audio system |

---

## 6. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174F60 | `Audible` | Known | Audible audiobook format |
| 0x002AA8D9 | ` Audible v` | Known | Audible audiobook format |
| 0x002AA92B | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x002AA941 | ` 2002 Audible, Inc. V` | Known | Audible audiobook format |
| 0x002B1318 | `Audible-softwaren i dette produkt bruges i henhold til en licensaftale fra Audib` | Known | Audible audiobook format |
| 0x002B1378 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x002B8194 | `Die Audible Software in diesem Produkt wird in Lizenz der Audible verwendet. Cop` | Known | Audible audiobook format |
| 0x002B81ED | ` 2002 Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x002C091F | ` Audible ` | Known | Audible audiobook format |
| 0x002C097C | ` Audible. ` | Known | Audible audiobook format |
| 0x002C09B2 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002C8530 | `El software Audible incluido en este producto se usa bajo licencia de Audible. C` | Known | Audible audiobook format |
| 0x002C858B | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x002CEF36 | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright ` | Known | Audible audiobook format |
| 0x002CEF70 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x002D65D4 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x002D661E | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x002D6633 | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x002DD582 | ` Audible szoftver az Audible licence alatt van haszn` | Known | Audible audiobook format |
| 0x002DD5CC | ` 2002, Audible, Inc. Minden jog fenntartva.` | Known | Audible audiobook format |
| 0x002E4234 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x002E425D | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x002E428C | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x002EBABD | ` Audible ` | Known | Audible audiobook format |
| 0x002EBADE | `Audible ` | Known | Audible audiobook format |
| 0x002EBB37 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002F3127 | ` Audible ` | Known | Audible audiobook format |
| 0x002F3142 | ` Audible` | Known | Audible audiobook format |
| 0x002F3186 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002FA198 | `De Audible-software in dit product wordt gebruikt in licentie van Audible. Copyr` | Known | Audible audiobook format |
| 0x002FA1EF | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x00300A14 | `Audible-programvaren i dette produktet brukes under lisens fra Audible. Copyrigh` | Known | Audible audiobook format |
| 0x00300A68 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0030730C | `Oprogramowanie Audible w tym produkcie jest wykorzystywane na podstawie licencji` | Known | Audible audiobook format |
| 0x00307378 | ` 2002 Audible, Inc. Wszystkie prawa zastrze` | Known | Audible audiobook format |
| 0x0030DDE4 | `O software Audible ` | Known | Audible audiobook format |
| 0x0030DE1A | `a da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0030DE34 | ` 2002 da Audible, Inc. Reservados todos os direitos.` | Known | Audible audiobook format |
| 0x0031655D | ` Audible ` | Known | Audible audiobook format |
| 0x003165AF | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x003165C5 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x0031DD48 | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x0031DD77 | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x0031DD8E | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x00324670 | `ndeki Audible yaz` | Known | Audible audiobook format |
| 0x00324689 | ` Audible lisans` | Known | Audible audiobook format |
| 0x003246BE | ` 2002, Audible, Inc. T` | Known | Audible audiobook format |
| 0x0032B2EB | ` Audible ` | Known | Audible audiobook format |
| 0x0032B2FD | ` Audible ` | Known | Audible audiobook format |
| 0x0032B321 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x00331DB8 | `Audible ` | Known | Audible audiobook format |
| 0x00331DCC | ` Audible ` | Known | Audible audiobook format |
| 0x00331DF6 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x004B4E3C | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x004B4E91 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |

---

## 7. Audio/Codec - Apple Lossless

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174F34 | `AppleLossless` | Known | Apple Lossless codec |
| 0x002DDAFC | `l alacsony.` | Known | Apple Lossless codec |

---

## 8. Audio/Codec - AAC

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00AEA060 | `AACDEC  VLL ` | Known | AAC codec |
| 0x00AF37A8 | `MPEG4 AAC LC Decoder` | Known | AAC codec |
| 0x00AF5C0D | `AACDecoderGetMem` | Known | AAC codec |
| 0x00AF5C1E | `AACDecoderInit` | Known | AAC codec |
| 0x00AF5C2D | `AACDecoderGetConfig` | Known | AAC codec |
| 0x00AF5C41 | `AACDecoderSetConfig` | Known | AAC codec |
| 0x00AF5C55 | `AACHeaderDecode` | Known | AAC codec |
| 0x00AF5C65 | `AACDecode` | Known | AAC codec |
| 0x00AF5C6F | `AACDecoderInit_Ittiam` | Known | AAC codec |
| 0x00AF5C85 | `AACDecoderGetConfig_Ittiam` | Known | AAC codec |
| 0x00AF5CA0 | `AACDecoderSetConfig_Ittiam` | Known | AAC codec |
| 0x00AF5CBB | `AACHeaderDecode_Ittiam` | Known | AAC codec |
| 0x00AF5CD2 | `AACDecode_Ittiam` | Known | AAC codec |

---

## 9. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002AAB1E | ` zvuku MPEG Layer-3 byla poskytnuta spole` | Known | MP3 codec |
| 0x002AAB49 | `nostmi Fraunhofer IIS a` | Known | MP3 codec |
| 0x002B1514 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunhofer IIS og THOMSON multi` | Known | MP3 codec |
| 0x002B8397 | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und THOMSON Multimedia.` | Known | MP3 codec |
| 0x002C0C77 | ` MPEG Layer-3 ` | Known | MP3 codec |
| 0x002C0CB5 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x002C8729 | `n de audio MPEG Layer-3 utilizada bajo licencia de Fraunhofer IIS y THOMSON mult` | Known | MP3 codec |
| 0x002CF0DC | `MPEG Layer-3 -` | Known | MP3 codec |
| 0x002CF0EE | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | MP3 codec |
| 0x002D67F0 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | MP3 codec |
| 0x002DD754 | `Az MPEG Layer-3 hangk` | Known | MP3 codec |
| 0x002DD77C | `gia a Fraunhofer IIS ` | Known | MP3 codec |
| 0x002E43D4 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da Fraunhofer IIS e THOMSON` | Known | MP3 codec |
| 0x002EBCEC | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002EBD38 | `Fraunhofer IIS ` | Known | MP3 codec |
| 0x002F3310 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002F3337 | ` Fraunhofer IIS` | Known | MP3 codec |
| 0x002FA37C | `Technologie voor codering van MPEG Layer-3-audio in licentie van Fraunhofer IIS ` | Known | MP3 codec |
| 0x00300BE4 | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fraunhofer IIS og THOMSON m` | Known | MP3 codec |
| 0x00307520 | `Technologia kodowania audio MPEG Layer-3 licencjonowana od Fraunhofer IIS oraz T` | Known | MP3 codec |
| 0x0030DFD6 | `udio MPEG Layer-3 licenciada pela Fraunhofer IIS e THOMSON multimedia.` | Known | MP3 codec |
| 0x003167E4 | `MPEG Layer-3: ` | Known | MP3 codec |
| 0x0031683D | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x0031DF28 | `Ljudkodningstekniken MPEG Layer-3 ` | Known | MP3 codec |
| 0x0031DF5E | `n Fraunhofer IIS och THOMSON multimedia.` | Known | MP3 codec |
| 0x00324848 | `MPEG Layer-3 ses kodlama teknolojisi Fraunhofer IIS ve THOMSON multimedia'dan li` | Known | MP3 codec |
| 0x0032B478 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x0032B49A | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x00331F54 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00331F79 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x004B5008 | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |

---

## 10. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E6680 | `drmsdrmisinffniscpsap@-` | Known | DRM system |
| 0x00174E58 | `AppleDRMVersion` | Known | DRM system |
| 0x00174EF8 | `AppleDRM` | Known | DRM system |
| 0x00175F24 | `AppleVideoDRM` | Known | DRM system |
| 0x0017AFF0 | `drmsmp4aesdsmp4v` | Known | DRM system |
| 0x001C63CC | `tkhdedtselstmdiamdhdminfstblstsdstcoco64stscstszsttsstssdrmidrms` | Known | DRM system |
| 0x006795EB | `DRMLevel` | Known | DRM system |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E2744 | `games_RO` | Known | Game system |
| 0x000E2750 | `gamedata_RW` | Known | Game system |

---

## 12. Photo System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009A800 | `Photo Database` | Known | Photo system |
| 0x000B950C | `Photos\Photo Database` | Known | Photo system |
| 0x000C0CF0 | `Photo Database` | Known | Photo system |
| 0x00195D08 | `23iUPhoto Database` | Known | Photo system |
| 0x00197D60 | `Photo Database` | Known | Photo system |
| 0x001980C4 | `Photo Database` | Known | Photo system |
| 0x00198370 | `Photo Import Database` | Known | Photo system |
| 0x0020F414 | `Photo Database Size` | Known | Photo system |

---

## 13. Video System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00B04700 | `H.264 Video Decoder` | Known | Video system |

---

## 14. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000B9500 | `iTunesDB` | Known | iTunes database |
| 0x001BB79C | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x00205938 | `iTunes Image DB` | Known | iTunes database |
| 0x002A4028 | `iTunesDB` | Known | iTunes database |
| 0x002AA0B9 | ` z iTunes nebo vCards. ` | Known | iTunes database |
| 0x002AA3D3 | `ipojte iPod k iTunes a instalujte hru znovu.` | Known | iTunes database |
| 0x002AA463 | `ipojte iPod k iTunes a zkop` | Known | iTunes database |
| 0x002AA504 | `i a iTunes ho odemkne.` | Known | iTunes database |
| 0x002AA620 | `m iTunes.` | Known | iTunes database |
| 0x002AAFA9 | `m iTunes.` | Known | iTunes database |
| 0x002AB96F | `es iTunes.` | Known | iTunes database |
| 0x002ABA15 | `es iTunes.` | Known | iTunes database |
| 0x002B0B0C | `iPod kan opbevare og vise kontaktoplysninger importeret fra iTunes eller vCards.` | Known | iTunes database |
| 0x002B0E94 | `Slut iPod til iTunes, og installer spillet igen.` | Known | iTunes database |
| 0x002B0F2C | `Slut iPod til iTunes, og overf` | Known | iTunes database |
| 0x002B0FC3 | `slutte iPod til computeren, hvorefter iTunes l` | Known | iTunes database |
| 0x002B1078 | `r fotografier til computeren, og synkroniser via iTunes for at vise dem p` | Known | iTunes database |
| 0x002B18E0 | `%s er for gammel til denne iPod. Slut iPod til computeren, og start iTunes for a` | Known | iTunes database |
| 0x002B216A | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x002B21F9 | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x002B790F | `nnen Kontakte (mit iTunes importiert oder vCards) auf Ihrem iPod sichern und anz` | Known | iTunes database |
| 0x002B7CA0 | `Verbinden Sie Ihren iPod mit iTunes und installieren Sie das Spiel erneut.` | Known | iTunes database |
| 0x002B7D64 | `Verbinden Sie Ihren iPod mit iTunes und laden Sie die aktuelle Version.` | Known | iTunes database |
| 0x002B7E0B | `en Sie Ihren iPod an Ihren Computer an und iTunes deaktiviert die Anzeigensperre` | Known | iTunes database |
| 0x002B7EB4 | `Importierte Fotos werden nicht auf dem TV angezeigt. Senden Sie sie erst an den ` | Known | iTunes database |
| 0x002B87B9 | `en Sie den iPod an den Computer an und starten Sie iTunes, um %s auf die aktuell` | Known | iTunes database |
| 0x002B90ED | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x002B918E | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x002BFB98 | ` iTunes ` | Known | iTunes database |
| 0x002C00E1 | ` iTunes ` | Known | iTunes database |
| 0x002C0209 | ` iTunes ` | Known | iTunes database |
| 0x002C0332 | ` iTunes ` | Known | iTunes database |
| 0x002C04C7 | ` iTunes ` | Known | iTunes database |
| 0x002C13A6 | ` iTunes ` | Known | iTunes database |
| 0x002C254A | ` iTunes ` | Known | iTunes database |
| 0x002C2673 | ` iTunes ` | Known | iTunes database |
| 0x002C7D01 | `n importada de iTunes o de tarjetas virtuales (vCards). ` | Known | iTunes database |
| 0x002C808C | `Conecte el iPod a iTunes y reinstale el juego.` | Known | iTunes database |
| 0x002C8124 | `Conecte el iPod a iTunes y descargue la versi` | Known | iTunes database |
| 0x002C81AC | `n, conecte el iPod al ordenador y iTunes lo desbloquear` | Known | iTunes database |
| 0x002C8288 | `celas con iTunes para verlas en la TV.` | Known | iTunes database |
| 0x002C8B3C | `%s es demasiado antiguo para ejecutarse en este iPod. Conecte el iPod al ordenad` | Known | iTunes database |
| 0x002C9498 | `Esta foto es demasiado grande para mostrarla en el iPod. Transfiera las fotos im` | Known | iTunes database |
| 0x002C953C | `Este formato de foto no puede visualizarse en el iPod. Transfiera las fotos impo` | Known | iTunes database |
| 0x002CE778 | ` iTunesista tai vCardeina tuotua tietoa. ` | Known | iTunes database |
| 0x002CEADE | ` iPod iTunesiin ja asenna peli uudelleen.` | Known | iTunes database |
| 0x002CEB6E | ` iPod iTunesiin ja hae uusin versio.` | Known | iTunes database |
| 0x002CEBE3 | `tietokoneeseen, niin iTunes avaa lukituksen.` | Known | iTunes database |
| 0x002CEC80 | ` kuvat tietokoneelle ja synkronoi ne iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002CF4EA | ` %s uusimpaan versioon avaamalla iTunes.` | Known | iTunes database |
| 0x002CFD82 | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002CFE09 | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002D5CDE | `iTunes ou de vCards. ` | Known | iTunes database |
| 0x002D608C | `Connectez votre iPod avec iTunes et r` | Known | iTunes database |
| 0x002D6134 | `Connectez votre iPod avec iTunes et t` | Known | iTunes database |
| 0x002D61F5 | ` votre ordinateur et iTunes le d` | Known | iTunes database |
| 0x002D62D8 | `rez-les sur votre ordinateur puis synchronisez-les avec iTunes.` | Known | iTunes database |
| 0x002D6C14 | `ordinateur et lancez iTunes pour mettre ` | Known | iTunes database |
| 0x002D76D6 | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x002D7793 | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x002DCC0C | `Az iPod, az iTunes programb` | Known | iTunes database |
| 0x002DD025 | `t az iTunes programhoz, ` | Known | iTunes database |
| 0x002DD0F5 | `t az iTunes programhoz ` | Known | iTunes database |
| 0x002DD1A3 | `phez, hogy az iTunes feloldja a z` | Known | iTunes database |
| 0x002DD2B1 | `ljon az iTunes haszn` | Known | iTunes database |
| 0x002DDBF4 | `s az iTunes futtat` | Known | iTunes database |
| 0x002DE687 | `lja azokat az iTunes programmal.` | Known | iTunes database |
| 0x002DE75E | `lja azokat az iTunes programmal.` | Known | iTunes database |
| 0x002E39A4 | ` memorizzare e visualizzare informazioni importate da iTunes o vCards. ` | Known | iTunes database |
| 0x002E3D6C | `Collega iPod a iTunes e reinstalla il gioco.` | Known | iTunes database |
| 0x002E3E10 | `Collega  iPod a iTunes ed esegui il download dell'ultima versione.` | Known | iTunes database |
| 0x002E3E88 | `Se dimentichi la combinazione, collega iPod al computer e iTunes sar` | Known | iTunes database |
| 0x002E3F28 | `Le foto importate non possono visualizzarsi in TV. Trasferisci le foto sul compu` | Known | iTunes database |
| 0x002E477D | ` troppo vecchio per funzionare con questo iPod. Collega iPod al computer ed eseg` | Known | iTunes database |
| 0x002E5091 | ` troppo grande per essere visualizzato in iPod. Trasferisci le foto sul computer` | Known | iTunes database |
| 0x002E513B | ` essere visualizzato in iPod. Trasferisci le foto sul computer e sincronizzale t` | Known | iTunes database |
| 0x002EAFAB | `iTunes ` | Known | iTunes database |
| 0x002EB43D | `iTunes ` | Known | iTunes database |
| 0x002EB549 | `iTunes ` | Known | iTunes database |
| 0x002EB649 | `iTunes` | Known | iTunes database |
| 0x002EB774 | `iTunes ` | Known | iTunes database |
| 0x002EC1FB | `iTunes ` | Known | iTunes database |
| 0x002ECD9F | `iTunes ` | Known | iTunes database |
| 0x002ECE65 | `iTunes ` | Known | iTunes database |
| 0x002F280E | ` iTunes ` | Known | iTunes database |
| 0x002F2BE7 | ` iTunes` | Known | iTunes database |
| 0x002F2CAF | ` iTunes` | Known | iTunes database |
| 0x002F2D86 | ` iTunes` | Known | iTunes database |
| 0x002F2E69 | ` iTunes` | Known | iTunes database |
| 0x002F3773 | ` iTunes` | Known | iTunes database |
| 0x002F4166 | ` iTunes` | Known | iTunes database |
| 0x002F420B | ` iTunes` | Known | iTunes database |
| 0x002F9999 | `mporteerd uit iTunes of vCards. ` | Known | iTunes database |
| 0x002F9D28 | `Verbind de iPod met iTunes en installeer het spel opnieuw.` | Known | iTunes database |
| 0x002F9DC8 | `Verbind de iPod met iTunes en download de nieuwste versie.` | Known | iTunes database |
| 0x002F9E3C | `Als u de combinatie bent vergeten, verbind iPod met uw computer en iTunes zal he` | Known | iTunes database |
| 0x002F9EED | `mporteerde foto's op tv niet mogelijk. Kopieer foto's naar de computer en synchr` | Known | iTunes database |
| 0x002FA7A4 | `%s is te oud om op deze iPod te worden gebruikt. Sluit de iPod aan op de compute` | Known | iTunes database |
| 0x002FB097 | `mporteerde foto's naar de computer en synchroniseer ze met iTunes voor weergave ` | Known | iTunes database |
| 0x002FB12E | `mporteerde foto's naar de computer en synchroniseer ze met iTunes voor weergave ` | Known | iTunes database |
| | *...and 72 more* | | |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002B6254 | `Radio-Region` | Known | FM Radio |
| 0x004B2B54 | `Radio Region` | Known | FM Radio |
| 0x004B5568 | `Radio Region` | Known | FM Radio |
| 0x004B5B2C | `Radio Settings` | Known | FM Radio |

---

## 16. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00175390 | `FireWireGUID` | Known | FireWire interface |
| 0x001753A0 | `FireWireVersion` | Known | FireWire interface |
| 0x0017598C | `FireWire` | Known | FireWire interface |
| 0x002A9EF7 | ` FireWire nen` | Known | FireWire interface |
| 0x002ABF28 | `FireWire p` | Known | FireWire interface |
| 0x002B0970 | `FireWire-forbindelser underst` | Known | FireWire interface |
| 0x002B26C4 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x002B7768 | `FireWire wird nicht unterst` | Known | FireWire interface |
| 0x002B9672 | `ber FireWire verbunden` | Known | FireWire interface |
| 0x002BF8BA | ` FireWire. ` | Known | FireWire interface |
| 0x002C2EF2 | ` FireWire` | Known | FireWire interface |
| 0x002C7B41 | `de canciones o archivos no son posibles con FireWire :` | Known | FireWire interface |
| 0x002C9AA4 | `FireWire conectado` | Known | FireWire interface |
| 0x002CE5B8 | `FireWire-tiedonsiirtoa ei tueta. Siirt` | Known | FireWire interface |
| 0x002D02F0 | `FireWire liitetty` | Known | FireWire interface |
| 0x002D5AE1 | `s via FireWire : connectez l` | Known | FireWire interface |
| 0x002D7D38 | `FireWire Connect` | Known | FireWire interface |
| 0x002DCA54 | `A FireWire kapcsolat nem t` | Known | FireWire interface |
| 0x002DEC94 | `FireWire csatlakozik` | Known | FireWire interface |
| 0x002E37F0 | `Connessioni di dati via FireWire non sono supportate. Per trasferire brani o dat` | Known | FireWire interface |
| 0x002E5684 | `FireWire connesso` | Known | FireWire interface |
| 0x002EAD28 | `FireWire ` | Known | FireWire interface |
| 0x002ED490 | `FireWire ` | Known | FireWire interface |
| 0x002F2620 | `FireWire ` | Known | FireWire interface |
| 0x002F472C | `FireWire ` | Known | FireWire interface |
| 0x002F97E6 | `ren via FireWire, maar alleen via de meegeleverde USB-kabel.` | Known | FireWire interface |
| 0x002FB628 | `FireWire aangesloten` | Known | FireWire interface |
| 0x00300087 | `ring via FireWire st` | Known | FireWire interface |
| 0x00301D58 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x00306933 | `czenie FireWire nie jest wspierane. By przes` | Known | FireWire interface |
| 0x0030888B | `czone przez FireWire` | Known | FireWire interface |
| 0x0030D467 | `es FireWire n` | Known | FireWire interface |
| 0x0030F39C | `FireWire ligado` | Known | FireWire interface |
| 0x0031544D | ` FireWire ` | Known | FireWire interface |
| 0x003188AB | ` FireWire` | Known | FireWire interface |
| 0x0031D3C0 | `FireWire-` | Known | FireWire interface |
| 0x0031F10C | `FireWire anslutet` | Known | FireWire interface |
| 0x00323BF0 | `FireWire ba` | Known | FireWire interface |
| 0x00325B3C | `FireWire Ba` | Known | FireWire interface |
| 0x0032A961 | ` FireWire ` | Known | FireWire interface |
| 0x0032C5CC | `FireWire ` | Known | FireWire interface |
| 0x00331411 | ` FireWire ` | Known | FireWire interface |
| 0x00333174 | `FireWire ` | Known | FireWire interface |
| 0x004B44C8 | `FireWire connections are not supported. To transfer songs, connect the USB cable` | Known | FireWire interface |
| 0x004B61A0 | `FireWire Connected` | Known | FireWire interface |

---

## 17. USB

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x006BE6D9 | `USBCompositeDevice1.6` | Known | USB interface |
| 0x006BE731 | `USBCompositeDevice1.6` | Known | USB interface |

---

## 18. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F48B8 | `LCD Module could not be determined.` | Known | Hardware interface |
| 0x00175E34 | `ForcedDiskMode` | Known | Hardware interface |
| 0x0020F374 | `Enter Disk Mode` | Known | Hardware interface |
| 0x0020F384 | `Exit Disk Mode` | Known | Hardware interface |
| 0x004B44BC | `Disk Mode` | Known | Hardware interface |
| 0x004FDDB0 | `I2C write Error` | Known | Hardware interface |
| 0x004FDDC4 | `I2C read Error %02x` | Known | Hardware interface |
| 0x0067E3A3 | `OCSP_RESPID` | Known | Hardware interface |

---

## 19. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0014A548 | `PowerManager` | Known | Power management |
| 0x00175968 | `PowerInformation` | Known | Power management |
| 0x0020F3E0 | `Begin Charging` | Known | Power management |
| 0x0020F3F0 | `Stop Charging` | Known | Power management |
| 0x00257E18 | `USBPowerSense` | Known | Power management |
| 0x00257ED8 | `PCFPowerMgr` | Known | Power management |
| 0x004B4064 | `Charging` | Known | Power management |
| 0x004B61EC | `Low Battery` | Known | Power management |

---

## 20. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002AEF88 | `Alarmer` | Known | UI element |
| 0x002B2608 | `Alarmer` | Known | UI element |
| 0x002C5F90 | `Calendario` | Known | UI element |
| 0x002C5F9C | `Calendarios` | Known | UI element |
| 0x002C5FA8 | `Calendarios` | Known | UI element |
| 0x002C5FE4 | `Alarmas` | Known | UI element |
| 0x002C6AFC | `Calendario` | Known | UI element |
| 0x002C6B08 | `Calendarios` | Known | UI element |
| 0x002C7FFC | `Alarma` | Known | UI element |
| 0x002C8C48 | `Alarma` | Known | UI element |
| 0x002C8CB8 | `Alarma` | Known | UI element |
| 0x002C904A | `Calendario` | Known | UI element |
| 0x002C92E8 | `Alarma` | Known | UI element |
| 0x002C999C | `Alarma` | Known | UI element |
| 0x002C99EC | `Alarmas` | Known | UI element |
| 0x002D3DD0 | `Alarmes` | Known | UI element |
| 0x002D5FD0 | `Alarme` | Known | UI element |
| 0x002D6CF8 | `Alarme` | Known | UI element |
| 0x002D6D60 | `Alarme` | Known | UI element |
| 0x002D7440 | `Alarme` | Known | UI element |
| 0x002D7BE8 | `Alarme` | Known | UI element |
| 0x002D7C60 | `Alarmes` | Known | UI element |
| 0x002E1CAC | `Calendario` | Known | UI element |
| 0x002E1CB8 | `Calendari` | Known | UI element |
| 0x002E1CC4 | `Calendari` | Known | UI element |
| 0x002E27D8 | `Calendario` | Known | UI element |
| 0x002E27E4 | `Calendari` | Known | UI element |
| 0x002E4C47 | `Calendario` | Known | UI element |
| 0x002FE670 | `Alarmer` | Known | UI element |
| 0x00301634 | `Alarmtidspunkt` | Known | UI element |
| 0x00301C94 | `Alarmer` | Known | UI element |
| 0x00304DD4 | `Alarmy` | Known | UI element |
| 0x0030516C | `Gotowe` | Known | UI element |
| 0x00305348 | `Gotowe` | Known | UI element |
| 0x003087A8 | `Alarmy` | Known | UI element |
| 0x0030B924 | `Alarmes` | Known | UI element |
| 0x0030D8E8 | `Alarme` | Known | UI element |
| 0x0030E4C4 | `Alarme` | Known | UI element |
| 0x0030F2DC | `Alarmes` | Known | UI element |
| 0x0031E9EC | `Alarmtid` | Known | UI element |
| 0x00322158 | `Alarmlar` | Known | UI element |
| 0x00325384 | `Alarm Zaman` | Known | UI element |
| 0x00325A80 | `Alarmlar` | Known | UI element |
| 0x004B25A8 | `Calendar` | Known | UI element |
| 0x004B25B4 | `Calendars` | Known | UI element |
| 0x004B25C0 | `Calendars` | Known | UI element |
| 0x004B25F4 | `Alarms` | Known | UI element |
| 0x004B3088 | `Calendar` | Known | UI element |
| 0x004B3094 | `Calendars` | Known | UI element |
| 0x004B54E8 | `Alarm Clock` | Known | UI element |
| 0x004B5802 | `Calendar` | Known | UI element |
| 0x004B5A74 | `Alarm Time` | Known | UI element |
| 0x004B5A80 | `Alarm Clock` | Known | UI element |
| 0x004B6088 | `Alarm Clock` | Known | UI element |
| 0x004B60F4 | `Alarms` | Known | UI element |
| 0x004B62F4 | `GotoBackToIdleCommand` | Known | UI element |
| 0x005019E0 | `TCalendarCntlr_Alarm` | Known | UI element |
| 0x0067954C | `Calendars/` | Known | UI element |
| 0x00679567 | `Calendars` | Known | UI element |

---

## 21. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00080DD8 | `Settings` | Known | Menu item |
| 0x002A8D80 | `Podcasts` | Known | Menu item |
| 0x002A8EEC | `Podcasts` | Known | Menu item |
| 0x002AB516 | `Podcasts` | Known | Menu item |
| 0x002ABE88 | `Podcasts` | Known | Menu item |
| 0x002AF8C4 | `Podcasts` | Known | Menu item |
| 0x002AFA14 | `Podcasts` | Known | Menu item |
| 0x002B1E05 | `Podcasts` | Known | Menu item |
| 0x002B2634 | `Podcasts` | Known | Menu item |
| 0x002B6674 | `Podcasts` | Known | Menu item |
| 0x002B6774 | `Extras` | Known | Menu item |
| 0x002B67A4 | `Videos` | Known | Menu item |
| 0x002B67DC | `Podcasts` | Known | Menu item |
| 0x002B8594 | `Videos` | Known | Menu item |
| 0x002B8C74 | `Extras` | Known | Menu item |
| 0x002B8C7C | `Videos` | Known | Menu item |
| 0x002B8CDF | `Podcasts` | Known | Menu item |
| 0x002B8FB4 | `Videos` | Known | Menu item |
| 0x002B95DC | `Podcasts` | Known | Menu item |
| 0x002B96D8 | `Extras` | Known | Menu item |
| 0x002BDA90 | `Podcasts` | Known | Menu item |
| 0x002BDD60 | `Podcasts` | Known | Menu item |
| 0x002C1D81 | `Podcasts` | Known | Menu item |
| 0x002C2DCC | `Podcasts` | Known | Menu item |
| 0x002C6A2C | `Podcasts` | Known | Menu item |
| 0x002C6B28 | `Extras` | Known | Menu item |
| 0x002C6B8C | `Podcasts` | Known | Menu item |
| 0x002C90A0 | `Extras` | Known | Menu item |
| 0x002C9112 | `Podcasts` | Known | Menu item |
| 0x002C9A14 | `Podcasts` | Known | Menu item |
| 0x002C9B08 | `Extras` | Known | Menu item |
| 0x002D48C4 | `Podcasts` | Known | Menu item |
| 0x002D490C | `Albums` | Known | Menu item |
| 0x002D4924 | `Genres` | Known | Menu item |
| 0x002D4964 | `Photos` | Known | Menu item |
| 0x002D49E4 | `Extras` | Known | Menu item |
| 0x002D4A54 | `Podcasts` | Known | Menu item |
| 0x002D4B5C | `Albums` | Known | Menu item |
| 0x002D6430 | `Photos` | Known | Menu item |
| 0x002D64DC | `Photos` | Known | Menu item |
| 0x002D69D0 | `Photos` | Known | Menu item |
| 0x002D71AC | `Extras` | Known | Menu item |
| 0x002D71E0 | `Photos` | Known | Menu item |
| 0x002D720E | `Genres` | Known | Menu item |
| 0x002D7232 | `Podcasts` | Known | Menu item |
| 0x002D7266 | `Albums` | Known | Menu item |
| 0x002D73E4 | `Genres` | Known | Menu item |
| 0x002D73F8 | `Albums` | Known | Menu item |
| 0x002D77BC | `Photos` | Known | Menu item |
| 0x002D7C74 | `Genres` | Known | Menu item |
| 0x002D7C88 | `Podcasts` | Known | Menu item |
| 0x002D7CA4 | `Albums` | Known | Menu item |
| 0x002D7DC0 | `Extras` | Known | Menu item |
| 0x002DB8A8 | `Podcasts` | Known | Menu item |
| 0x002DBA20 | `Podcasts` | Known | Menu item |
| 0x002DE1AE | `Podcasts` | Known | Menu item |
| 0x002DEBF4 | `Podcasts` | Known | Menu item |
| 0x002F86E4 | `Podcasts` | Known | Menu item |
| 0x002F8728 | `Albums` | Known | Menu item |
| 0x002F873C | `Genres` | Known | Menu item |
| 0x002F8850 | `Podcasts` | Known | Menu item |
| 0x002F8928 | `Albums` | Known | Menu item |
| 0x002FACCB | `Genres` | Known | Menu item |
| 0x002FACEB | `Podcasts` | Known | Menu item |
| 0x002FAD13 | `Albums` | Known | Menu item |
| 0x002FAE54 | `Genres` | Known | Menu item |
| 0x002FAE64 | `Albums` | Known | Menu item |
| 0x002FB57C | `Genres` | Known | Menu item |
| 0x002FB590 | `Podcasts` | Known | Menu item |
| 0x002FB5A8 | `Albums` | Known | Menu item |
| 0x00305770 | `Podcasts` | Known | Menu item |
| 0x00307E9A | `Podcasts` | Known | Menu item |
| 0x003087D0 | `Podcasts` | Known | Menu item |
| 0x0030C318 | `Podcasts` | Known | Menu item |
| 0x0030C43C | `Extras` | Known | Menu item |
| 0x0030C4A0 | `Podcasts` | Known | Menu item |
| 0x0030E950 | `Extras` | Known | Menu item |
| 0x0030E9CA | `Podcasts` | Known | Menu item |
| 0x0030F314 | `Podcasts` | Known | Menu item |
| 0x0030F3E8 | `Extras` | Known | Menu item |
| 0x00313640 | `Podcasts` | Known | Menu item |
| 0x003138B8 | `Podcasts` | Known | Menu item |
| 0x003177D9 | `Podcasts` | Known | Menu item |
| 0x00318790 | `Podcasts` | Known | Menu item |
| 0x00322B68 | `Podcasts` | Known | Menu item |
| 0x00322CD4 | `Podcasts` | Known | Menu item |
| 0x003251C2 | `Podcasts` | Known | Menu item |
| 0x00325AAC | `Podcasts` | Known | Menu item |
| 0x003330E8 | `Podcasts` | Known | Menu item |
| 0x004B2ED4 | `Podcasts` | Known | Menu item |
| 0x004B2FE4 | `Now Playing` | Known | Menu item |
| 0x004B2FF0 | `Artists` | Known | Menu item |
| 0x004B3008 | `Albums` | Known | Menu item |
| 0x004B3020 | `Genres` | Known | Menu item |
| 0x004B3028 | `Composers` | Known | Menu item |
| 0x004B3054 | `Photos` | Known | Menu item |
| 0x004B30BC | `Extras` | Known | Menu item |
| 0x004B30C4 | `Playlists` | Known | Menu item |
| 0x004B30D0 | `Audiobooks` | Known | Menu item |
| 0x004B30E4 | `Videos` | Known | Menu item |
| | *...and 39 more* | | |

---

## 22. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009A55C | `iPod_Control` | Filesystem Path | |
| 0x0009A588 | `iPod_Control\Device` | Filesystem Path | |
| 0x0009CEAC | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000A6180 | `iPod_Control\Device` | Filesystem Path | |
| 0x000A7B6C | `iPod_Control` | Filesystem Path | |
| 0x000A81C4 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x000B94E8 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000B9528 | `System_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x000BC014 | `iPod_Control\Music\` | Filesystem Path | |
| 0x000C05B8 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x000C0738 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000E2770 | `iPod_Control/%s/%s%s%s` | Filesystem Path | |
| 0x000EF180 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000F6BA4 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000F8634 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000F8730 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001A5AA4 | `iPod_Control/Accessories` | Filesystem Path | |
| 0x001A6004 | `iPod_Control/Accessories` | Filesystem Path | |
| 0x001CF124 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF3B0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF468 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF5B8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF6D8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF7A8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF940 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF9FC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFAAC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFBA0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFC44 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFCF8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFDB4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFEE8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0058 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D011C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D01CC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0308 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D03D4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D04A0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0568 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D060C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D06D4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0784 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0834 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D08FC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D09BC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0A6C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0B1C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0BCC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0C7C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0D50 | `iPod_Control\Device\` | Filesystem Path | |

---

## 23. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002A9344 | `Acoustic` | EQ Preset | |
| 0x002A9350 | `Bass Booster` | EQ Preset | |
| 0x002A9370 | `Classical` | EQ Preset | |
| 0x002A938C | `Electronic` | EQ Preset | |
| 0x002A93A0 | `Hip Hop` | EQ Preset | |
| 0x002A93B8 | `Loudness` | EQ Preset | |
| 0x002A93C4 | `Lounge` | EQ Preset | |
| 0x002A93E8 | `Small Speakers` | EQ Preset | |
| 0x002A93F8 | `Spoken Word` | EQ Preset | |
| 0x002A9404 | `Treble Booster` | EQ Preset | |
| 0x002A9424 | `Vocal Booster` | EQ Preset | |
| 0x002AFE14 | `Acoustic` | EQ Preset | |
| 0x002AFE20 | `Bass Booster` | EQ Preset | |
| 0x002AFE40 | `Classical` | EQ Preset | |
| 0x002AFE5C | `Electronic` | EQ Preset | |
| 0x002AFE70 | `Hip Hop` | EQ Preset | |
| 0x002AFE88 | `Loudness` | EQ Preset | |
| 0x002AFE94 | `Lounge` | EQ Preset | |
| 0x002AFEB8 | `Small Speakers` | EQ Preset | |
| 0x002AFEC8 | `Spoken Word` | EQ Preset | |
| 0x002AFED4 | `Treble Booster` | EQ Preset | |
| 0x002AFEF4 | `Vocal Booster` | EQ Preset | |
| 0x002B6C08 | `Acoustic` | EQ Preset | |
| 0x002B6C48 | `Electronic` | EQ Preset | |
| 0x002B6C5C | `Hip Hop` | EQ Preset | |
| 0x002B6C74 | `Loudness` | EQ Preset | |
| 0x002BE4B8 | `Hip Hop` | EQ Preset | |
| 0x002BE4D0 | `Loudness` | EQ Preset | |
| 0x002BE4DC | `Lounge` | EQ Preset | |
| 0x002C7024 | `Hip Hop` | EQ Preset | |
| 0x002C7034 | `Latina` | EQ Preset | |
| 0x002C703C | `Loudness` | EQ Preset | |
| 0x002C7048 | `Lounge` | EQ Preset | |
| 0x002CDAB0 | `Lounge` | EQ Preset | |
| 0x002D4F30 | `Hip Hop` | EQ Preset | |
| 0x002D4F60 | `Lounge` | EQ Preset | |
| 0x002E2CB8 | `Hip Hop` | EQ Preset | |
| 0x002E2CC8 | `Latina` | EQ Preset | |
| 0x002E2CD0 | `Loudness` | EQ Preset | |
| 0x002E2CDC | `Lounge` | EQ Preset | |
| 0x002E9BB0 | `Acoustic` | EQ Preset | |
| 0x002E9BBC | `Bass Booster` | EQ Preset | |
| 0x002E9BDC | `Classical` | EQ Preset | |
| 0x002E9BF8 | `Electronic` | EQ Preset | |
| 0x002E9C0C | `Hip Hop` | EQ Preset | |
| 0x002E9C24 | `Loudness` | EQ Preset | |
| 0x002E9C30 | `Lounge` | EQ Preset | |
| 0x002E9C54 | `Small Speakers` | EQ Preset | |
| 0x002E9C64 | `Spoken Word` | EQ Preset | |
| 0x002E9C70 | `Treble Booster` | EQ Preset | |
| 0x002E9C90 | `Vocal Booster` | EQ Preset | |
| 0x002F173C | `Acoustic` | EQ Preset | |
| 0x002F1748 | `Bass Booster` | EQ Preset | |
| 0x002F1768 | `Classical` | EQ Preset | |
| 0x002F1784 | `Electronic` | EQ Preset | |
| 0x002F1798 | `Hip Hop` | EQ Preset | |
| 0x002F17B0 | `Loudness` | EQ Preset | |
| 0x002F17BC | `Lounge` | EQ Preset | |
| 0x002F17E0 | `Small Speakers` | EQ Preset | |
| 0x002F17F0 | `Spoken Word` | EQ Preset | |
| 0x002F17FC | `Treble Booster` | EQ Preset | |
| 0x002F181C | `Vocal Booster` | EQ Preset | |
| 0x002F8CD0 | `Loudness` | EQ Preset | |
| 0x002F8CDC | `Lounge` | EQ Preset | |
| 0x002FF580 | `Latino` | EQ Preset | |
| 0x002FF588 | `Loudness` | EQ Preset | |
| 0x002FF594 | `Lounge` | EQ Preset | |
| 0x00305DC0 | `Hip Hop` | EQ Preset | |
| 0x00305DF4 | `Lounge` | EQ Preset | |
| 0x0030C928 | `Hip Hop` | EQ Preset | |
| 0x0030C938 | `Latina` | EQ Preset | |
| 0x0030C940 | `Loudness` | EQ Preset | |
| 0x0030C94C | `Lounge` | EQ Preset | |
| 0x0031C878 | `Acoustic` | EQ Preset | |
| 0x0031C884 | `Bass Booster` | EQ Preset | |
| 0x0031C8A4 | `Classical` | EQ Preset | |
| 0x0031C8C0 | `Electronic` | EQ Preset | |
| 0x0031C8D4 | `Hip Hop` | EQ Preset | |
| 0x0031C8EC | `Loudness` | EQ Preset | |
| 0x0031C8F8 | `Lounge` | EQ Preset | |
| 0x0031C91C | `Small Speakers` | EQ Preset | |
| 0x0031C92C | `Spoken Word` | EQ Preset | |
| 0x0031C938 | `Treble Booster` | EQ Preset | |
| 0x0031C958 | `Vocal Booster` | EQ Preset | |
| 0x00323114 | `Hip Hop` | EQ Preset | |
| 0x00323128 | `Loudness` | EQ Preset | |
| 0x00323134 | `Lounge` | EQ Preset | |
| 0x00329BB0 | `Acoustic` | EQ Preset | |
| 0x00329BBC | `Bass Booster` | EQ Preset | |
| 0x00329BDC | `Classical` | EQ Preset | |
| 0x00329BF8 | `Electronic` | EQ Preset | |
| 0x00329C0C | `Hip Hop` | EQ Preset | |
| 0x00329C24 | `Loudness` | EQ Preset | |
| 0x00329C30 | `Lounge` | EQ Preset | |
| 0x00329C54 | `Small Speakers` | EQ Preset | |
| 0x00329C64 | `Spoken Word` | EQ Preset | |
| 0x00329C70 | `Treble Booster` | EQ Preset | |
| 0x00329C90 | `Vocal Booster` | EQ Preset | |
| 0x00330680 | `Acoustic` | EQ Preset | |
| 0x0033068C | `Bass Booster` | EQ Preset | |
| 0x003306AC | `Classical` | EQ Preset | |
| 0x003306C8 | `Electronic` | EQ Preset | |
| 0x003306DC | `Hip Hop` | EQ Preset | |
| 0x003306F4 | `Loudness` | EQ Preset | |
| 0x00330700 | `Lounge` | EQ Preset | |
| 0x00330724 | `Small Speakers` | EQ Preset | |
| 0x00330734 | `Spoken Word` | EQ Preset | |
| 0x00330740 | `Treble Booster` | EQ Preset | |
| 0x00330760 | `Vocal Booster` | EQ Preset | |
| 0x004B35DC | `Acoustic` | EQ Preset | |
| 0x004B35E8 | `Bass Booster` | EQ Preset | |
| 0x004B3608 | `Classical` | EQ Preset | |
| 0x004B3624 | `Electronic` | EQ Preset | |
| 0x004B3638 | `Hip Hop` | EQ Preset | |
| 0x004B3650 | `Loudness` | EQ Preset | |
| 0x004B365C | `Lounge` | EQ Preset | |
| 0x004B3680 | `Small Speakers` | EQ Preset | |
| 0x004B3690 | `Spoken Word` | EQ Preset | |
| 0x004B369C | `Treble Booster` | EQ Preset | |
| 0x004B36BC | `Vocal Booster` | EQ Preset | |

---

## 24. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0010559C | `Error-SDriver` | Diagnostic | |
| 0x001055AC | `Error-AClient` | Diagnostic | |
| 0x00106088 | `Root Hub Driver Internal Error unused case in hub handler` | Diagnostic | |
| 0x001060C4 | `Root hub Error Calling Add Device` | Diagnostic | |
| 0x0010AD44 | `Error inside %s` | Diagnostic | |
| 0x001408E8 | `%s Error in file %s.` | Diagnostic | |
| 0x00267194 | `Error inside %s` | Diagnostic | |
| 0x00267224 | `Error inside %s` | Diagnostic | |
| 0x002672A8 | `Error inside %s` | Diagnostic | |
| 0x00267760 | `Error inside %s` | Diagnostic | |
| 0x00267824 | `Error inside %s` | Diagnostic | |
| 0x002678F0 | `Error inside %s` | Diagnostic | |
| 0x00267BAC | `Error inside %s` | Diagnostic | |
| 0x00267D9C | `Error inside %s` | Diagnostic | |
| 0x00267E00 | `Error inside %s` | Diagnostic | |
| 0x00267F34 | `Error inside %s` | Diagnostic | |
| 0x00267F8C | `Error inside %s` | Diagnostic | |
| 0x00267FDC | `Error inside %s` | Diagnostic | |
| 0x002680AC | `Error inside %s` | Diagnostic | |
| 0x002680FC | `Error inside %s` | Diagnostic | |
| 0x002684F0 | `Error inside %s` | Diagnostic | |
| 0x00268560 | `Error inside %s` | Diagnostic | |
| 0x002689E0 | `Error inside %s` | Diagnostic | |
| 0x00268E48 | `Error inside %s` | Diagnostic | |
| 0x0026904C | `Error inside %s` | Diagnostic | |
| 0x002690BC | `Error inside %s` | Diagnostic | |
| 0x002691BC | `Error inside %s` | Diagnostic | |
| 0x002692C4 | `Error inside %s` | Diagnostic | |
| 0x00269338 | `Error inside %s` | Diagnostic | |
| 0x00269384 | `Error inside %s` | Diagnostic | |

---
