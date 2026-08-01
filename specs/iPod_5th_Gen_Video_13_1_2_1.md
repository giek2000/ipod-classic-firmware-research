# iPod 5th Generation (Video 30GB) - RetailOS 1.2.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.2.1 |
| **IPSW** | iPod_13.1.2.1.ipsw |
| **Device** | iPod 5th Generation (Video 30GB) (2005, Click Wheel, Video Playback) |
| **UpdaterFamilyID** | 13 |
| **Binary Size** | 13,834,752 bytes (13.19 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 13,834,752 bytes |
| **Total Strings (>=6)** | 29,921 |
| **Function Prologues** | 22,319 (ARM: 12,639, Thumb: 9,680) |
| **SoC** | PortalPlayer PP5021C |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Encrypted** | No |
| **SHA-256** | `d16ffe10d63012e4d9880018cae6c75e4a957a81e47586b4cf0a4686a8a0df5a` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00169AA4 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x0020CE28 | `Channel UnitTests` | Hidden | Developer Tool |

---

## 2. Controllers (TSilver/TC Classes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00AD7678 | `TCC_Current_Task_Pointer` | Known | UI controller |
| 0x00AD7691 | `TCC_Delete_HISR` | Known | UI controller |
| 0x00AD76B1 | `TCC_Relinquish` | Known | UI controller |
| 0x00AD76CF | `TCC_Resume_Service` | Known | UI controller |
| 0x00AD76E2 | `TCC_Task_Sleep` | Known | UI controller |
| 0x00AD7704 | `TCF_Task_Information` | Known | UI controller |
| 0x00AD7719 | `TCS_Change_Preemption` | Known | UI controller |
| 0x00AD772F | `TCS_Change_Priority` | Known | UI controller |
| 0x00AD7743 | `TCT_Activate_HISR` | Known | UI controller |
| 0x00AD7755 | `TCT_Control_Interrupts` | Known | UI controller |
| 0x00AD776C | `TCT_Local_Control_Interrupts` | Known | UI controller |
| 0x00B3EF1F | `TCC_Resume_Service` | Known | UI controller |
| 0x00B3F0A2 | `TCC_Delete_HISR` | Known | UI controller |
| 0x00B3F0F1 | `TCT_Activate_HISR` | Known | UI controller |
| 0x00B3F11B | `TCT_Control_Interrupts` | Known | UI controller |
| 0x00B3F2E5 | `TCC_Current_Task_Pointer` | Known | UI controller |
| 0x00B3F2FE | `TCS_Change_Priority` | Known | UI controller |
| 0x00B3F36C | `TCC_Task_Sleep` | Known | UI controller |
| 0x00B3F4A5 | `TCF_Task_Information` | Known | UI controller |
| 0x00B4DA1D | `TCC_Relinquish` | Known | UI controller |
| 0x00B4DB33 | `TCT_Local_Control_Interrupts` | Known | UI controller |
| 0x00B4DBD3 | `TCC_Task_Sleep` | Known | UI controller |
| 0x00B4DCA1 | `TCS_Change_Preemption` | Known | UI controller |

---

## 3. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009A600 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x000C8DE8 | `USB Secondary Interrupt Task` | Known | RTOS task thread |
| 0x000E1818 | `ICAPTPCameraIOTask` | Known | RTOS task thread |
| 0x0011BCB4 | `USBStatusTask` | Known | RTOS task thread |
| 0x0011BCD0 | `USBTaskTimeTask` | Known | RTOS task thread |
| 0x00163BE8 | `VCUpdateTask` | Known | RTOS task thread |
| 0x0016F674 | `USBDeviceTask` | Known | RTOS task thread |
| 0x00175B88 | `DiskReaderTask` | Known | RTOS task thread |
| 0x00184DF8 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x00184E0C | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x00198044 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x00206CD8 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x002573BC | `FirewireTask` | Known | RTOS task thread |
| 0x002573D4 | `OptoTask` | Known | RTOS task thread |
| 0x002573E4 | `SerialOptoTask` | Known | RTOS task thread |
| 0x002573F8 | `BacklightTask` | Known | RTOS task thread |
| 0x0025740C | `CNATask` | Known | RTOS task thread |
| 0x0025742C | `DiskMgrTask` | Known | RTOS task thread |
| 0x0025743C | `HoldSwitchTask` | Known | RTOS task thread |
| 0x00257450 | `TopPlugTask` | Known | RTOS task thread |
| 0x00257460 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x00257474 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x0025748C | `AccessoryDetectTask` | Known | RTOS task thread |
| 0x002574B4 | `AlarmTask` | Known | RTOS task thread |
| 0x002574C4 | `WatchdogTask` | Known | RTOS task thread |
| 0x0025753C | `USBAudioTask` | Known | RTOS task thread |
| 0x002A3870 | `HostOSTask` | Known | RTOS task thread |
| 0x002A441C | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x00501EB8 | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x00501ED0 | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x00501EE8 | `VideoDaisyTask` | Known | RTOS task thread |
| 0x00AD7668 | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00AD76A1 | `TCC_Delete_Task` | Known | RTOS task thread |
| 0x00AD76C0 | `TCC_Reset_Task` | Known | RTOS task thread |
| 0x00AD76F1 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B3EF10 | `TCC_Reset_Task` | Known | RTOS task thread |
| 0x00B3EF47 | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B3F0B2 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B3F0C5 | `TCC_Delete_Task` | Known | RTOS task thread |
| 0x00B4DA0D | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B4DA57 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B4DA6A | `TCC_Delete_Task` | Known | RTOS task thread |

---

## 4. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0020CC34 | `Channel Reserved` | Known | Logging channel |
| 0x0020CC48 | `Channel AppBoot` | Known | Logging channel |
| 0x0020CC58 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x0020CC74 | `Channel PrefsWriting` | Known | Logging channel |
| 0x0020CC8C | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x0020CCAC | `Channel PlayFromDisk` | Known | Logging channel |
| 0x0020CCC4 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x0020CCE0 | `Channel TestLogging` | Known | Logging channel |
| 0x0020CCF4 | `Channel AppFileLoading` | Known | Logging channel |
| 0x0020CD0C | `Channel VCardReading` | Known | Logging channel |
| 0x0020CD24 | `Channel LongSongScanning` | Known | Logging channel |
| 0x0020CD98 | `Channel VoiceRecording` | Known | Logging channel |
| 0x0020CDB0 | `Channel PhotoImporting` | Known | Logging channel |
| 0x0020CDC8 | `Channel Notes` | Known | Logging channel |
| 0x0020CDD8 | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x0020CDF4 | `Channel DiskMode` | Known | Logging channel |
| 0x0020CE08 | `Channel Firewire` | Known | Logging channel |
| 0x0020CE1C | `Channel USB` | Known | Logging channel |
| 0x0020CE3C | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x0020CE54 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 5. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174484 | `AudioCodecs` | Known | Audio system |
| 0x00175510 | `VideoCodecs` | Known | Audio system |
| 0x002B0A6A | `.net codec i dette produkt bruges i henhold til en licensaftale fra VoiceAge Cor` | Known | Audio system |
| 0x002B78D9 | `.net Codec in diesem Produkt wird in Lizenz der VoiceAge Corporation verwendet. ` | Known | Audio system |
| 0x002C0130 | `.net codec ` | Known | Audio system |
| 0x002D5CE4 | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x002DCCC1 | `.net codec a VoiceAge Coporation c` | Known | Audio system |
| 0x002E38FE | ` utilizzato su licenza da VoiceAge Corporation. Il codec ACELP` | Known | Audio system |
| 0x002F2876 | `.net codec` | Known | Audio system |
| 0x002F98E0 | `.net-codec in dit product wordt gebruikt in licentie van VoiceAge Corporation. G` | Known | Audio system |
| 0x0030D4F5 | `a da VoiceAge Corporation. O codec ACELP` | Known | Audio system |
| 0x00323DB9 | `.net codec'i VoiceAge Corporation lisans` | Known | Audio system |
| 0x004B4575 | `.net codec in this product is used under license from VoiceAge Corporation. Port` | Known | Audio system |
| 0x0067B99C | `msCodeCom` | Known | Audio system |
| 0x00B3E99D | `codec_string` | Known | Audio system |
| 0x00B3E9AA | `codec_name` | Known | Audio system |
| 0x00B4D2C9 | `codec_string` | Known | Audio system |
| 0x00B4D2D6 | `codec_name` | Known | Audio system |

---

## 6. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174560 | `Audible` | Known | Audible audiobook format |
| 0x002A9ED9 | ` Audible v` | Known | Audible audiobook format |
| 0x002A9F2B | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x002A9F41 | ` 2002 Audible, Inc. V` | Known | Audible audiobook format |
| 0x002B0918 | `Audible-softwaren i dette produkt bruges i henhold til en licensaftale fra Audib` | Known | Audible audiobook format |
| 0x002B0978 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x002B7794 | `Die Audible Software in diesem Produkt wird in Lizenz der Audible verwendet. Cop` | Known | Audible audiobook format |
| 0x002B77ED | ` 2002 Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x002BFF1F | ` Audible ` | Known | Audible audiobook format |
| 0x002BFF7C | ` Audible. ` | Known | Audible audiobook format |
| 0x002BFFB2 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002C7B30 | `El software Audible incluido en este producto se usa bajo licencia de Audible. C` | Known | Audible audiobook format |
| 0x002C7B8B | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x002CE536 | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright ` | Known | Audible audiobook format |
| 0x002CE570 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x002D5BD4 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x002D5C1E | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x002D5C33 | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x002DCB82 | ` Audible szoftver az Audible licence alatt van haszn` | Known | Audible audiobook format |
| 0x002DCBCC | ` 2002, Audible, Inc. Minden jog fenntartva.` | Known | Audible audiobook format |
| 0x002E3834 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x002E385D | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x002E388C | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x002EB0BD | ` Audible ` | Known | Audible audiobook format |
| 0x002EB0DE | `Audible ` | Known | Audible audiobook format |
| 0x002EB137 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002F2727 | ` Audible ` | Known | Audible audiobook format |
| 0x002F2742 | ` Audible` | Known | Audible audiobook format |
| 0x002F2786 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002F9798 | `De Audible-software in dit product wordt gebruikt in licentie van Audible. Copyr` | Known | Audible audiobook format |
| 0x002F97EF | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x00300014 | `Audible-programvaren i dette produktet brukes under lisens fra Audible. Copyrigh` | Known | Audible audiobook format |
| 0x00300068 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0030690C | `Oprogramowanie Audible w tym produkcie jest wykorzystywane na podstawie licencji` | Known | Audible audiobook format |
| 0x00306978 | ` 2002 Audible, Inc. Wszystkie prawa zastrze` | Known | Audible audiobook format |
| 0x0030D3E4 | `O software Audible ` | Known | Audible audiobook format |
| 0x0030D41A | `a da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0030D434 | ` 2002 da Audible, Inc. Reservados todos os direitos.` | Known | Audible audiobook format |
| 0x00315B5D | ` Audible ` | Known | Audible audiobook format |
| 0x00315BAF | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x00315BC5 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x0031D348 | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x0031D377 | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x0031D38E | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x00323C70 | `ndeki Audible yaz` | Known | Audible audiobook format |
| 0x00323C89 | ` Audible lisans` | Known | Audible audiobook format |
| 0x00323CBE | ` 2002, Audible, Inc. T` | Known | Audible audiobook format |
| 0x0032A8EB | ` Audible ` | Known | Audible audiobook format |
| 0x0032A8FD | ` Audible ` | Known | Audible audiobook format |
| 0x0032A921 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x003313B8 | `Audible ` | Known | Audible audiobook format |
| 0x003313CC | ` Audible ` | Known | Audible audiobook format |
| 0x003313F6 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x004B443C | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x004B4491 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |

---

## 7. Audio/Codec - Apple Lossless

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174534 | `AppleLossless` | Known | Apple Lossless codec |
| 0x002DD0FC | `l alacsony.` | Known | Apple Lossless codec |

---

## 8. Audio/Codec - AAC

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00AE8660 | `AACDEC  VLL ` | Known | AAC codec |
| 0x00AF1DA8 | `MPEG4 AAC LC Decoder` | Known | AAC codec |
| 0x00AF420D | `AACDecoderGetMem` | Known | AAC codec |
| 0x00AF421E | `AACDecoderInit` | Known | AAC codec |
| 0x00AF422D | `AACDecoderGetConfig` | Known | AAC codec |
| 0x00AF4241 | `AACDecoderSetConfig` | Known | AAC codec |
| 0x00AF4255 | `AACHeaderDecode` | Known | AAC codec |
| 0x00AF4265 | `AACDecode` | Known | AAC codec |
| 0x00AF426F | `AACDecoderInit_Ittiam` | Known | AAC codec |
| 0x00AF4285 | `AACDecoderGetConfig_Ittiam` | Known | AAC codec |
| 0x00AF42A0 | `AACDecoderSetConfig_Ittiam` | Known | AAC codec |
| 0x00AF42BB | `AACHeaderDecode_Ittiam` | Known | AAC codec |
| 0x00AF42D2 | `AACDecode_Ittiam` | Known | AAC codec |

---

## 9. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002AA11E | ` zvuku MPEG Layer-3 byla poskytnuta spole` | Known | MP3 codec |
| 0x002AA149 | `nostmi Fraunhofer IIS a` | Known | MP3 codec |
| 0x002B0B14 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunhofer IIS og THOMSON multi` | Known | MP3 codec |
| 0x002B7997 | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und THOMSON Multimedia.` | Known | MP3 codec |
| 0x002C0277 | ` MPEG Layer-3 ` | Known | MP3 codec |
| 0x002C02B5 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x002C7D29 | `n de audio MPEG Layer-3 utilizada bajo licencia de Fraunhofer IIS y THOMSON mult` | Known | MP3 codec |
| 0x002CE6DC | `MPEG Layer-3 -` | Known | MP3 codec |
| 0x002CE6EE | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | MP3 codec |
| 0x002D5DF0 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | MP3 codec |
| 0x002DCD54 | `Az MPEG Layer-3 hangk` | Known | MP3 codec |
| 0x002DCD7C | `gia a Fraunhofer IIS ` | Known | MP3 codec |
| 0x002E39D4 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da Fraunhofer IIS e THOMSON` | Known | MP3 codec |
| 0x002EB2EC | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002EB338 | `Fraunhofer IIS ` | Known | MP3 codec |
| 0x002F2910 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002F2937 | ` Fraunhofer IIS` | Known | MP3 codec |
| 0x002F997C | `Technologie voor codering van MPEG Layer-3-audio in licentie van Fraunhofer IIS ` | Known | MP3 codec |
| 0x003001E4 | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fraunhofer IIS og THOMSON m` | Known | MP3 codec |
| 0x00306B20 | `Technologia kodowania audio MPEG Layer-3 licencjonowana od Fraunhofer IIS oraz T` | Known | MP3 codec |
| 0x0030D5D6 | `udio MPEG Layer-3 licenciada pela Fraunhofer IIS e THOMSON multimedia.` | Known | MP3 codec |
| 0x00315DE4 | `MPEG Layer-3: ` | Known | MP3 codec |
| 0x00315E3D | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x0031D528 | `Ljudkodningstekniken MPEG Layer-3 ` | Known | MP3 codec |
| 0x0031D55E | `n Fraunhofer IIS och THOMSON multimedia.` | Known | MP3 codec |
| 0x00323E48 | `MPEG Layer-3 ses kodlama teknolojisi Fraunhofer IIS ve THOMSON multimedia'dan li` | Known | MP3 codec |
| 0x0032AA78 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x0032AA9A | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x00331554 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00331579 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x004B4608 | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |

---

## 10. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E5C80 | `drmsdrmisinffniscpsap@-` | Known | DRM system |
| 0x00174458 | `AppleDRMVersion` | Known | DRM system |
| 0x001744F8 | `AppleDRM` | Known | DRM system |
| 0x00175524 | `AppleVideoDRM` | Known | DRM system |
| 0x0017A5F0 | `drmsmp4aesdsmp4v` | Known | DRM system |
| 0x001C59CC | `tkhdedtselstmdiamdhdminfstblstsdstcoco64stscstszsttsstssdrmidrms` | Known | DRM system |
| 0x00678BEB | `DRMLevel` | Known | DRM system |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E1D44 | `games_RO` | Known | Game system |
| 0x000E1D50 | `gamedata_RW` | Known | Game system |

---

## 12. Photo System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00099E00 | `Photo Database` | Known | Photo system |
| 0x000B8B0C | `Photos\Photo Database` | Known | Photo system |
| 0x000C02F0 | `Photo Database` | Known | Photo system |
| 0x00195308 | `23iUPhoto Database` | Known | Photo system |
| 0x00197360 | `Photo Database` | Known | Photo system |
| 0x001976C4 | `Photo Database` | Known | Photo system |
| 0x00197970 | `Photo Import Database` | Known | Photo system |
| 0x0020EA14 | `Photo Database Size` | Known | Photo system |

---

## 13. Video System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00B02D00 | `H.264 Video Decoder` | Known | Video system |

---

## 14. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000B8B00 | `iTunesDB` | Known | iTunes database |
| 0x001BAD9C | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x00204F38 | `iTunes Image DB` | Known | iTunes database |
| 0x002A3628 | `iTunesDB` | Known | iTunes database |
| 0x002A96B9 | ` z iTunes nebo vCards. ` | Known | iTunes database |
| 0x002A99D3 | `ipojte iPod k iTunes a instalujte hru znovu.` | Known | iTunes database |
| 0x002A9A63 | `ipojte iPod k iTunes a zkop` | Known | iTunes database |
| 0x002A9B04 | `i a iTunes ho odemkne.` | Known | iTunes database |
| 0x002A9C20 | `m iTunes.` | Known | iTunes database |
| 0x002AA5A9 | `m iTunes.` | Known | iTunes database |
| 0x002AAF6F | `es iTunes.` | Known | iTunes database |
| 0x002AB015 | `es iTunes.` | Known | iTunes database |
| 0x002B010C | `iPod kan opbevare og vise kontaktoplysninger importeret fra iTunes eller vCards.` | Known | iTunes database |
| 0x002B0494 | `Slut iPod til iTunes, og installer spillet igen.` | Known | iTunes database |
| 0x002B052C | `Slut iPod til iTunes, og overf` | Known | iTunes database |
| 0x002B05C3 | `slutte iPod til computeren, hvorefter iTunes l` | Known | iTunes database |
| 0x002B0678 | `r fotografier til computeren, og synkroniser via iTunes for at vise dem p` | Known | iTunes database |
| 0x002B0EE0 | `%s er for gammel til denne iPod. Slut iPod til computeren, og start iTunes for a` | Known | iTunes database |
| 0x002B176A | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x002B17F9 | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x002B6F0F | `nnen Kontakte (mit iTunes importiert oder vCards) auf Ihrem iPod sichern und anz` | Known | iTunes database |
| 0x002B72A0 | `Verbinden Sie Ihren iPod mit iTunes und installieren Sie das Spiel erneut.` | Known | iTunes database |
| 0x002B7364 | `Verbinden Sie Ihren iPod mit iTunes und laden Sie die aktuelle Version.` | Known | iTunes database |
| 0x002B740B | `en Sie Ihren iPod an Ihren Computer an und iTunes deaktiviert die Anzeigensperre` | Known | iTunes database |
| 0x002B74B4 | `Importierte Fotos werden nicht auf dem TV angezeigt. Senden Sie sie erst an den ` | Known | iTunes database |
| 0x002B7DB9 | `en Sie den iPod an den Computer an und starten Sie iTunes, um %s auf die aktuell` | Known | iTunes database |
| 0x002B86ED | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x002B878E | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x002BF198 | ` iTunes ` | Known | iTunes database |
| 0x002BF6E1 | ` iTunes ` | Known | iTunes database |
| 0x002BF809 | ` iTunes ` | Known | iTunes database |
| 0x002BF932 | ` iTunes ` | Known | iTunes database |
| 0x002BFAC7 | ` iTunes ` | Known | iTunes database |
| 0x002C09A6 | ` iTunes ` | Known | iTunes database |
| 0x002C1B4A | ` iTunes ` | Known | iTunes database |
| 0x002C1C73 | ` iTunes ` | Known | iTunes database |
| 0x002C7301 | `n importada de iTunes o de tarjetas virtuales (vCards). ` | Known | iTunes database |
| 0x002C768C | `Conecte el iPod a iTunes y reinstale el juego.` | Known | iTunes database |
| 0x002C7724 | `Conecte el iPod a iTunes y descargue la versi` | Known | iTunes database |
| 0x002C77AC | `n, conecte el iPod al ordenador y iTunes lo desbloquear` | Known | iTunes database |
| 0x002C7888 | `celas con iTunes para verlas en la TV.` | Known | iTunes database |
| 0x002C813C | `%s es demasiado antiguo para ejecutarse en este iPod. Conecte el iPod al ordenad` | Known | iTunes database |
| 0x002C8A98 | `Esta foto es demasiado grande para mostrarla en el iPod. Transfiera las fotos im` | Known | iTunes database |
| 0x002C8B3C | `Este formato de foto no puede visualizarse en el iPod. Transfiera las fotos impo` | Known | iTunes database |
| 0x002CDD78 | ` iTunesista tai vCardeina tuotua tietoa. ` | Known | iTunes database |
| 0x002CE0DE | ` iPod iTunesiin ja asenna peli uudelleen.` | Known | iTunes database |
| 0x002CE16E | ` iPod iTunesiin ja hae uusin versio.` | Known | iTunes database |
| 0x002CE1E3 | `tietokoneeseen, niin iTunes avaa lukituksen.` | Known | iTunes database |
| 0x002CE280 | ` kuvat tietokoneelle ja synkronoi ne iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002CEAEA | ` %s uusimpaan versioon avaamalla iTunes.` | Known | iTunes database |
| 0x002CF382 | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002CF409 | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002D52DE | `iTunes ou de vCards. ` | Known | iTunes database |
| 0x002D568C | `Connectez votre iPod avec iTunes et r` | Known | iTunes database |
| 0x002D5734 | `Connectez votre iPod avec iTunes et t` | Known | iTunes database |
| 0x002D57F5 | ` votre ordinateur et iTunes le d` | Known | iTunes database |
| 0x002D58D8 | `rez-les sur votre ordinateur puis synchronisez-les avec iTunes.` | Known | iTunes database |
| 0x002D6214 | `ordinateur et lancez iTunes pour mettre ` | Known | iTunes database |
| 0x002D6CD6 | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x002D6D93 | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x002DC20C | `Az iPod, az iTunes programb` | Known | iTunes database |
| 0x002DC625 | `t az iTunes programhoz, ` | Known | iTunes database |
| 0x002DC6F5 | `t az iTunes programhoz ` | Known | iTunes database |
| 0x002DC7A3 | `phez, hogy az iTunes feloldja a z` | Known | iTunes database |
| 0x002DC8B1 | `ljon az iTunes haszn` | Known | iTunes database |
| 0x002DD1F4 | `s az iTunes futtat` | Known | iTunes database |
| 0x002DDC87 | `lja azokat az iTunes programmal.` | Known | iTunes database |
| 0x002DDD5E | `lja azokat az iTunes programmal.` | Known | iTunes database |
| 0x002E2FA4 | ` memorizzare e visualizzare informazioni importate da iTunes o vCards. ` | Known | iTunes database |
| 0x002E336C | `Collega iPod a iTunes e reinstalla il gioco.` | Known | iTunes database |
| 0x002E3410 | `Collega  iPod a iTunes ed esegui il download dell'ultima versione.` | Known | iTunes database |
| 0x002E3488 | `Se dimentichi la combinazione, collega iPod al computer e iTunes sar` | Known | iTunes database |
| 0x002E3528 | `Le foto importate non possono visualizzarsi in TV. Trasferisci le foto sul compu` | Known | iTunes database |
| 0x002E3D7D | ` troppo vecchio per funzionare con questo iPod. Collega iPod al computer ed eseg` | Known | iTunes database |
| 0x002E4691 | ` troppo grande per essere visualizzato in iPod. Trasferisci le foto sul computer` | Known | iTunes database |
| 0x002E473B | ` essere visualizzato in iPod. Trasferisci le foto sul computer e sincronizzale t` | Known | iTunes database |
| 0x002EA5AB | `iTunes ` | Known | iTunes database |
| 0x002EAA3D | `iTunes ` | Known | iTunes database |
| 0x002EAB49 | `iTunes ` | Known | iTunes database |
| 0x002EAC49 | `iTunes` | Known | iTunes database |
| 0x002EAD74 | `iTunes ` | Known | iTunes database |
| 0x002EB7FB | `iTunes ` | Known | iTunes database |
| 0x002EC39F | `iTunes ` | Known | iTunes database |
| 0x002EC465 | `iTunes ` | Known | iTunes database |
| 0x002F1E0E | ` iTunes ` | Known | iTunes database |
| 0x002F21E7 | ` iTunes` | Known | iTunes database |
| 0x002F22AF | ` iTunes` | Known | iTunes database |
| 0x002F2386 | ` iTunes` | Known | iTunes database |
| 0x002F2469 | ` iTunes` | Known | iTunes database |
| 0x002F2D73 | ` iTunes` | Known | iTunes database |
| 0x002F3766 | ` iTunes` | Known | iTunes database |
| 0x002F380B | ` iTunes` | Known | iTunes database |
| 0x002F8F99 | `mporteerd uit iTunes of vCards. ` | Known | iTunes database |
| 0x002F9328 | `Verbind de iPod met iTunes en installeer het spel opnieuw.` | Known | iTunes database |
| 0x002F93C8 | `Verbind de iPod met iTunes en download de nieuwste versie.` | Known | iTunes database |
| 0x002F943C | `Als u de combinatie bent vergeten, verbind iPod met uw computer en iTunes zal he` | Known | iTunes database |
| 0x002F94ED | `mporteerde foto's op tv niet mogelijk. Kopieer foto's naar de computer en synchr` | Known | iTunes database |
| 0x002F9DA4 | `%s is te oud om op deze iPod te worden gebruikt. Sluit de iPod aan op de compute` | Known | iTunes database |
| 0x002FA697 | `mporteerde foto's naar de computer en synchroniseer ze met iTunes voor weergave ` | Known | iTunes database |
| 0x002FA72E | `mporteerde foto's naar de computer en synchroniseer ze met iTunes voor weergave ` | Known | iTunes database |
| | *...and 72 more* | | |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002B5854 | `Radio-Region` | Known | FM Radio |
| 0x004B2154 | `Radio Region` | Known | FM Radio |
| 0x004B4B68 | `Radio Region` | Known | FM Radio |
| 0x004B512C | `Radio Settings` | Known | FM Radio |

---

## 16. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174990 | `FireWireGUID` | Known | FireWire interface |
| 0x001749A0 | `FireWireVersion` | Known | FireWire interface |
| 0x00174F8C | `FireWire` | Known | FireWire interface |
| 0x002A94F7 | ` FireWire nen` | Known | FireWire interface |
| 0x002AB528 | `FireWire p` | Known | FireWire interface |
| 0x002AFF70 | `FireWire-forbindelser underst` | Known | FireWire interface |
| 0x002B1CC4 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x002B6D68 | `FireWire wird nicht unterst` | Known | FireWire interface |
| 0x002B8C72 | `ber FireWire verbunden` | Known | FireWire interface |
| 0x002BEEBA | ` FireWire. ` | Known | FireWire interface |
| 0x002C24F2 | ` FireWire` | Known | FireWire interface |
| 0x002C7141 | `de canciones o archivos no son posibles con FireWire :` | Known | FireWire interface |
| 0x002C90A4 | `FireWire conectado` | Known | FireWire interface |
| 0x002CDBB8 | `FireWire-tiedonsiirtoa ei tueta. Siirt` | Known | FireWire interface |
| 0x002CF8F0 | `FireWire liitetty` | Known | FireWire interface |
| 0x002D50E1 | `s via FireWire : connectez l` | Known | FireWire interface |
| 0x002D7338 | `FireWire Connect` | Known | FireWire interface |
| 0x002DC054 | `A FireWire kapcsolat nem t` | Known | FireWire interface |
| 0x002DE294 | `FireWire csatlakozik` | Known | FireWire interface |
| 0x002E2DF0 | `Connessioni di dati via FireWire non sono supportate. Per trasferire brani o dat` | Known | FireWire interface |
| 0x002E4C84 | `FireWire connesso` | Known | FireWire interface |
| 0x002EA328 | `FireWire ` | Known | FireWire interface |
| 0x002ECA90 | `FireWire ` | Known | FireWire interface |
| 0x002F1C20 | `FireWire ` | Known | FireWire interface |
| 0x002F3D2C | `FireWire ` | Known | FireWire interface |
| 0x002F8DE6 | `ren via FireWire, maar alleen via de meegeleverde USB-kabel.` | Known | FireWire interface |
| 0x002FAC28 | `FireWire aangesloten` | Known | FireWire interface |
| 0x002FF687 | `ring via FireWire st` | Known | FireWire interface |
| 0x00301358 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x00305F33 | `czenie FireWire nie jest wspierane. By przes` | Known | FireWire interface |
| 0x00307E8B | `czone przez FireWire` | Known | FireWire interface |
| 0x0030CA67 | `es FireWire n` | Known | FireWire interface |
| 0x0030E99C | `FireWire ligado` | Known | FireWire interface |
| 0x00314A4D | ` FireWire ` | Known | FireWire interface |
| 0x00317EAB | ` FireWire` | Known | FireWire interface |
| 0x0031C9C0 | `FireWire-` | Known | FireWire interface |
| 0x0031E70C | `FireWire anslutet` | Known | FireWire interface |
| 0x003231F0 | `FireWire ba` | Known | FireWire interface |
| 0x0032513C | `FireWire Ba` | Known | FireWire interface |
| 0x00329F61 | ` FireWire ` | Known | FireWire interface |
| 0x0032BBCC | `FireWire ` | Known | FireWire interface |
| 0x00330A11 | ` FireWire ` | Known | FireWire interface |
| 0x00332774 | `FireWire ` | Known | FireWire interface |
| 0x004B3AC8 | `FireWire connections are not supported. To transfer songs, connect the USB cable` | Known | FireWire interface |
| 0x004B57A0 | `FireWire Connected` | Known | FireWire interface |

---

## 17. USB

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x006BDCD9 | `USBCompositeDevice1.6` | Known | USB interface |
| 0x006BDD31 | `USBCompositeDevice1.6` | Known | USB interface |

---

## 18. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F3EB8 | `LCD Module could not be determined.` | Known | Hardware interface |
| 0x00175434 | `ForcedDiskMode` | Known | Hardware interface |
| 0x0020E974 | `Enter Disk Mode` | Known | Hardware interface |
| 0x0020E984 | `Exit Disk Mode` | Known | Hardware interface |
| 0x004B3ABC | `Disk Mode` | Known | Hardware interface |
| 0x004FD3B0 | `I2C write Error` | Known | Hardware interface |
| 0x004FD3C4 | `I2C read Error %02x` | Known | Hardware interface |
| 0x0067D9A3 | `OCSP_RESPID` | Known | Hardware interface |

---

## 19. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00149B48 | `PowerManager` | Known | Power management |
| 0x00174F68 | `PowerInformation` | Known | Power management |
| 0x0020E9E0 | `Begin Charging` | Known | Power management |
| 0x0020E9F0 | `Stop Charging` | Known | Power management |
| 0x00257418 | `USBPowerSense` | Known | Power management |
| 0x002574D8 | `PCFPowerMgr` | Known | Power management |
| 0x004B3664 | `Charging` | Known | Power management |
| 0x004B57EC | `Low Battery` | Known | Power management |

---

## 20. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002AE588 | `Alarmer` | Known | UI element |
| 0x002B1C08 | `Alarmer` | Known | UI element |
| 0x002C5590 | `Calendario` | Known | UI element |
| 0x002C559C | `Calendarios` | Known | UI element |
| 0x002C55A8 | `Calendarios` | Known | UI element |
| 0x002C55E4 | `Alarmas` | Known | UI element |
| 0x002C60FC | `Calendario` | Known | UI element |
| 0x002C6108 | `Calendarios` | Known | UI element |
| 0x002C75FC | `Alarma` | Known | UI element |
| 0x002C8248 | `Alarma` | Known | UI element |
| 0x002C82B8 | `Alarma` | Known | UI element |
| 0x002C864A | `Calendario` | Known | UI element |
| 0x002C88E8 | `Alarma` | Known | UI element |
| 0x002C8F9C | `Alarma` | Known | UI element |
| 0x002C8FEC | `Alarmas` | Known | UI element |
| 0x002D33D0 | `Alarmes` | Known | UI element |
| 0x002D55D0 | `Alarme` | Known | UI element |
| 0x002D62F8 | `Alarme` | Known | UI element |
| 0x002D6360 | `Alarme` | Known | UI element |
| 0x002D6A40 | `Alarme` | Known | UI element |
| 0x002D71E8 | `Alarme` | Known | UI element |
| 0x002D7260 | `Alarmes` | Known | UI element |
| 0x002E12AC | `Calendario` | Known | UI element |
| 0x002E12B8 | `Calendari` | Known | UI element |
| 0x002E12C4 | `Calendari` | Known | UI element |
| 0x002E1DD8 | `Calendario` | Known | UI element |
| 0x002E1DE4 | `Calendari` | Known | UI element |
| 0x002E4247 | `Calendario` | Known | UI element |
| 0x002FDC70 | `Alarmer` | Known | UI element |
| 0x00300C34 | `Alarmtidspunkt` | Known | UI element |
| 0x00301294 | `Alarmer` | Known | UI element |
| 0x003043D4 | `Alarmy` | Known | UI element |
| 0x0030476C | `Gotowe` | Known | UI element |
| 0x00304948 | `Gotowe` | Known | UI element |
| 0x00307DA8 | `Alarmy` | Known | UI element |
| 0x0030AF24 | `Alarmes` | Known | UI element |
| 0x0030CEE8 | `Alarme` | Known | UI element |
| 0x0030DAC4 | `Alarme` | Known | UI element |
| 0x0030E8DC | `Alarmes` | Known | UI element |
| 0x0031DFEC | `Alarmtid` | Known | UI element |
| 0x00321758 | `Alarmlar` | Known | UI element |
| 0x00324984 | `Alarm Zaman` | Known | UI element |
| 0x00325080 | `Alarmlar` | Known | UI element |
| 0x004B1BA8 | `Calendar` | Known | UI element |
| 0x004B1BB4 | `Calendars` | Known | UI element |
| 0x004B1BC0 | `Calendars` | Known | UI element |
| 0x004B1BF4 | `Alarms` | Known | UI element |
| 0x004B2688 | `Calendar` | Known | UI element |
| 0x004B2694 | `Calendars` | Known | UI element |
| 0x004B4AE8 | `Alarm Clock` | Known | UI element |
| 0x004B4E02 | `Calendar` | Known | UI element |
| 0x004B5074 | `Alarm Time` | Known | UI element |
| 0x004B5080 | `Alarm Clock` | Known | UI element |
| 0x004B5688 | `Alarm Clock` | Known | UI element |
| 0x004B56F4 | `Alarms` | Known | UI element |
| 0x004B58F4 | `GotoBackToIdleCommand` | Known | UI element |
| 0x00500FE0 | `TCalendarCntlr_Alarm` | Known | UI element |
| 0x00678B4C | `Calendars/` | Known | UI element |
| 0x00678B67 | `Calendars` | Known | UI element |

---

## 21. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000803D8 | `Settings` | Known | Menu item |
| 0x002A8380 | `Podcasts` | Known | Menu item |
| 0x002A84EC | `Podcasts` | Known | Menu item |
| 0x002AAB16 | `Podcasts` | Known | Menu item |
| 0x002AB488 | `Podcasts` | Known | Menu item |
| 0x002AEEC4 | `Podcasts` | Known | Menu item |
| 0x002AF014 | `Podcasts` | Known | Menu item |
| 0x002B1405 | `Podcasts` | Known | Menu item |
| 0x002B1C34 | `Podcasts` | Known | Menu item |
| 0x002B5C74 | `Podcasts` | Known | Menu item |
| 0x002B5D74 | `Extras` | Known | Menu item |
| 0x002B5DA4 | `Videos` | Known | Menu item |
| 0x002B5DDC | `Podcasts` | Known | Menu item |
| 0x002B7B94 | `Videos` | Known | Menu item |
| 0x002B8274 | `Extras` | Known | Menu item |
| 0x002B827C | `Videos` | Known | Menu item |
| 0x002B82DF | `Podcasts` | Known | Menu item |
| 0x002B85B4 | `Videos` | Known | Menu item |
| 0x002B8BDC | `Podcasts` | Known | Menu item |
| 0x002B8CD8 | `Extras` | Known | Menu item |
| 0x002BD090 | `Podcasts` | Known | Menu item |
| 0x002BD360 | `Podcasts` | Known | Menu item |
| 0x002C1381 | `Podcasts` | Known | Menu item |
| 0x002C23CC | `Podcasts` | Known | Menu item |
| 0x002C602C | `Podcasts` | Known | Menu item |
| 0x002C6128 | `Extras` | Known | Menu item |
| 0x002C618C | `Podcasts` | Known | Menu item |
| 0x002C86A0 | `Extras` | Known | Menu item |
| 0x002C8712 | `Podcasts` | Known | Menu item |
| 0x002C9014 | `Podcasts` | Known | Menu item |
| 0x002C9108 | `Extras` | Known | Menu item |
| 0x002D3EC4 | `Podcasts` | Known | Menu item |
| 0x002D3F0C | `Albums` | Known | Menu item |
| 0x002D3F24 | `Genres` | Known | Menu item |
| 0x002D3F64 | `Photos` | Known | Menu item |
| 0x002D3FE4 | `Extras` | Known | Menu item |
| 0x002D4054 | `Podcasts` | Known | Menu item |
| 0x002D415C | `Albums` | Known | Menu item |
| 0x002D5A30 | `Photos` | Known | Menu item |
| 0x002D5ADC | `Photos` | Known | Menu item |
| 0x002D5FD0 | `Photos` | Known | Menu item |
| 0x002D67AC | `Extras` | Known | Menu item |
| 0x002D67E0 | `Photos` | Known | Menu item |
| 0x002D680E | `Genres` | Known | Menu item |
| 0x002D6832 | `Podcasts` | Known | Menu item |
| 0x002D6866 | `Albums` | Known | Menu item |
| 0x002D69E4 | `Genres` | Known | Menu item |
| 0x002D69F8 | `Albums` | Known | Menu item |
| 0x002D6DBC | `Photos` | Known | Menu item |
| 0x002D7274 | `Genres` | Known | Menu item |
| 0x002D7288 | `Podcasts` | Known | Menu item |
| 0x002D72A4 | `Albums` | Known | Menu item |
| 0x002D73C0 | `Extras` | Known | Menu item |
| 0x002DAEA8 | `Podcasts` | Known | Menu item |
| 0x002DB020 | `Podcasts` | Known | Menu item |
| 0x002DD7AE | `Podcasts` | Known | Menu item |
| 0x002DE1F4 | `Podcasts` | Known | Menu item |
| 0x002F7CE4 | `Podcasts` | Known | Menu item |
| 0x002F7D28 | `Albums` | Known | Menu item |
| 0x002F7D3C | `Genres` | Known | Menu item |
| 0x002F7E50 | `Podcasts` | Known | Menu item |
| 0x002F7F28 | `Albums` | Known | Menu item |
| 0x002FA2CB | `Genres` | Known | Menu item |
| 0x002FA2EB | `Podcasts` | Known | Menu item |
| 0x002FA313 | `Albums` | Known | Menu item |
| 0x002FA454 | `Genres` | Known | Menu item |
| 0x002FA464 | `Albums` | Known | Menu item |
| 0x002FAB7C | `Genres` | Known | Menu item |
| 0x002FAB90 | `Podcasts` | Known | Menu item |
| 0x002FABA8 | `Albums` | Known | Menu item |
| 0x00304D70 | `Podcasts` | Known | Menu item |
| 0x0030749A | `Podcasts` | Known | Menu item |
| 0x00307DD0 | `Podcasts` | Known | Menu item |
| 0x0030B918 | `Podcasts` | Known | Menu item |
| 0x0030BA3C | `Extras` | Known | Menu item |
| 0x0030BAA0 | `Podcasts` | Known | Menu item |
| 0x0030DF50 | `Extras` | Known | Menu item |
| 0x0030DFCA | `Podcasts` | Known | Menu item |
| 0x0030E914 | `Podcasts` | Known | Menu item |
| 0x0030E9E8 | `Extras` | Known | Menu item |
| 0x00312C40 | `Podcasts` | Known | Menu item |
| 0x00312EB8 | `Podcasts` | Known | Menu item |
| 0x00316DD9 | `Podcasts` | Known | Menu item |
| 0x00317D90 | `Podcasts` | Known | Menu item |
| 0x00322168 | `Podcasts` | Known | Menu item |
| 0x003222D4 | `Podcasts` | Known | Menu item |
| 0x003247C2 | `Podcasts` | Known | Menu item |
| 0x003250AC | `Podcasts` | Known | Menu item |
| 0x003326E8 | `Podcasts` | Known | Menu item |
| 0x004B24D4 | `Podcasts` | Known | Menu item |
| 0x004B25E4 | `Now Playing` | Known | Menu item |
| 0x004B25F0 | `Artists` | Known | Menu item |
| 0x004B2608 | `Albums` | Known | Menu item |
| 0x004B2620 | `Genres` | Known | Menu item |
| 0x004B2628 | `Composers` | Known | Menu item |
| 0x004B2654 | `Photos` | Known | Menu item |
| 0x004B26BC | `Extras` | Known | Menu item |
| 0x004B26C4 | `Playlists` | Known | Menu item |
| 0x004B26D0 | `Audiobooks` | Known | Menu item |
| 0x004B26E4 | `Videos` | Known | Menu item |
| | *...and 39 more* | | |

---

## 22. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00099B5C | `iPod_Control` | Filesystem Path | |
| 0x00099B88 | `iPod_Control\Device` | Filesystem Path | |
| 0x0009C4AC | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000A5780 | `iPod_Control\Device` | Filesystem Path | |
| 0x000A716C | `iPod_Control` | Filesystem Path | |
| 0x000A77C4 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x000B8AE8 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000B8B28 | `System_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x000BB614 | `iPod_Control\Music\` | Filesystem Path | |
| 0x000BFBB8 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x000BFD38 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000E1D70 | `iPod_Control/%s/%s%s%s` | Filesystem Path | |
| 0x000EE780 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000F61A4 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000F7C34 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000F7D30 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001A50A4 | `iPod_Control/Accessories` | Filesystem Path | |
| 0x001A5604 | `iPod_Control/Accessories` | Filesystem Path | |
| 0x001CE724 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CE9B0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CEA68 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CEBB8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CECD8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CEDA8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CEF40 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CEFFC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF0AC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF1A0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF244 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF2F8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF3B4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF4E8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF658 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF71C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF7CC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF908 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CF9D4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFAA0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFB68 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFC0C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFCD4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFD84 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFE34 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFEFC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CFFBC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D006C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D011C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D01CC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D027C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0350 | `iPod_Control\Device\` | Filesystem Path | |

---

## 23. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002A8944 | `Acoustic` | EQ Preset | |
| 0x002A8950 | `Bass Booster` | EQ Preset | |
| 0x002A8970 | `Classical` | EQ Preset | |
| 0x002A898C | `Electronic` | EQ Preset | |
| 0x002A89A0 | `Hip Hop` | EQ Preset | |
| 0x002A89B8 | `Loudness` | EQ Preset | |
| 0x002A89C4 | `Lounge` | EQ Preset | |
| 0x002A89E8 | `Small Speakers` | EQ Preset | |
| 0x002A89F8 | `Spoken Word` | EQ Preset | |
| 0x002A8A04 | `Treble Booster` | EQ Preset | |
| 0x002A8A24 | `Vocal Booster` | EQ Preset | |
| 0x002AF414 | `Acoustic` | EQ Preset | |
| 0x002AF420 | `Bass Booster` | EQ Preset | |
| 0x002AF440 | `Classical` | EQ Preset | |
| 0x002AF45C | `Electronic` | EQ Preset | |
| 0x002AF470 | `Hip Hop` | EQ Preset | |
| 0x002AF488 | `Loudness` | EQ Preset | |
| 0x002AF494 | `Lounge` | EQ Preset | |
| 0x002AF4B8 | `Small Speakers` | EQ Preset | |
| 0x002AF4C8 | `Spoken Word` | EQ Preset | |
| 0x002AF4D4 | `Treble Booster` | EQ Preset | |
| 0x002AF4F4 | `Vocal Booster` | EQ Preset | |
| 0x002B6208 | `Acoustic` | EQ Preset | |
| 0x002B6248 | `Electronic` | EQ Preset | |
| 0x002B625C | `Hip Hop` | EQ Preset | |
| 0x002B6274 | `Loudness` | EQ Preset | |
| 0x002BDAB8 | `Hip Hop` | EQ Preset | |
| 0x002BDAD0 | `Loudness` | EQ Preset | |
| 0x002BDADC | `Lounge` | EQ Preset | |
| 0x002C6624 | `Hip Hop` | EQ Preset | |
| 0x002C6634 | `Latina` | EQ Preset | |
| 0x002C663C | `Loudness` | EQ Preset | |
| 0x002C6648 | `Lounge` | EQ Preset | |
| 0x002CD0B0 | `Lounge` | EQ Preset | |
| 0x002D4530 | `Hip Hop` | EQ Preset | |
| 0x002D4560 | `Lounge` | EQ Preset | |
| 0x002E22B8 | `Hip Hop` | EQ Preset | |
| 0x002E22C8 | `Latina` | EQ Preset | |
| 0x002E22D0 | `Loudness` | EQ Preset | |
| 0x002E22DC | `Lounge` | EQ Preset | |
| 0x002E91B0 | `Acoustic` | EQ Preset | |
| 0x002E91BC | `Bass Booster` | EQ Preset | |
| 0x002E91DC | `Classical` | EQ Preset | |
| 0x002E91F8 | `Electronic` | EQ Preset | |
| 0x002E920C | `Hip Hop` | EQ Preset | |
| 0x002E9224 | `Loudness` | EQ Preset | |
| 0x002E9230 | `Lounge` | EQ Preset | |
| 0x002E9254 | `Small Speakers` | EQ Preset | |
| 0x002E9264 | `Spoken Word` | EQ Preset | |
| 0x002E9270 | `Treble Booster` | EQ Preset | |
| 0x002E9290 | `Vocal Booster` | EQ Preset | |
| 0x002F0D3C | `Acoustic` | EQ Preset | |
| 0x002F0D48 | `Bass Booster` | EQ Preset | |
| 0x002F0D68 | `Classical` | EQ Preset | |
| 0x002F0D84 | `Electronic` | EQ Preset | |
| 0x002F0D98 | `Hip Hop` | EQ Preset | |
| 0x002F0DB0 | `Loudness` | EQ Preset | |
| 0x002F0DBC | `Lounge` | EQ Preset | |
| 0x002F0DE0 | `Small Speakers` | EQ Preset | |
| 0x002F0DF0 | `Spoken Word` | EQ Preset | |
| 0x002F0DFC | `Treble Booster` | EQ Preset | |
| 0x002F0E1C | `Vocal Booster` | EQ Preset | |
| 0x002F82D0 | `Loudness` | EQ Preset | |
| 0x002F82DC | `Lounge` | EQ Preset | |
| 0x002FEB80 | `Latino` | EQ Preset | |
| 0x002FEB88 | `Loudness` | EQ Preset | |
| 0x002FEB94 | `Lounge` | EQ Preset | |
| 0x003053C0 | `Hip Hop` | EQ Preset | |
| 0x003053F4 | `Lounge` | EQ Preset | |
| 0x0030BF28 | `Hip Hop` | EQ Preset | |
| 0x0030BF38 | `Latina` | EQ Preset | |
| 0x0030BF40 | `Loudness` | EQ Preset | |
| 0x0030BF4C | `Lounge` | EQ Preset | |
| 0x0031BE78 | `Acoustic` | EQ Preset | |
| 0x0031BE84 | `Bass Booster` | EQ Preset | |
| 0x0031BEA4 | `Classical` | EQ Preset | |
| 0x0031BEC0 | `Electronic` | EQ Preset | |
| 0x0031BED4 | `Hip Hop` | EQ Preset | |
| 0x0031BEEC | `Loudness` | EQ Preset | |
| 0x0031BEF8 | `Lounge` | EQ Preset | |
| 0x0031BF1C | `Small Speakers` | EQ Preset | |
| 0x0031BF2C | `Spoken Word` | EQ Preset | |
| 0x0031BF38 | `Treble Booster` | EQ Preset | |
| 0x0031BF58 | `Vocal Booster` | EQ Preset | |
| 0x00322714 | `Hip Hop` | EQ Preset | |
| 0x00322728 | `Loudness` | EQ Preset | |
| 0x00322734 | `Lounge` | EQ Preset | |
| 0x003291B0 | `Acoustic` | EQ Preset | |
| 0x003291BC | `Bass Booster` | EQ Preset | |
| 0x003291DC | `Classical` | EQ Preset | |
| 0x003291F8 | `Electronic` | EQ Preset | |
| 0x0032920C | `Hip Hop` | EQ Preset | |
| 0x00329224 | `Loudness` | EQ Preset | |
| 0x00329230 | `Lounge` | EQ Preset | |
| 0x00329254 | `Small Speakers` | EQ Preset | |
| 0x00329264 | `Spoken Word` | EQ Preset | |
| 0x00329270 | `Treble Booster` | EQ Preset | |
| 0x00329290 | `Vocal Booster` | EQ Preset | |
| 0x0032FC80 | `Acoustic` | EQ Preset | |
| 0x0032FC8C | `Bass Booster` | EQ Preset | |
| 0x0032FCAC | `Classical` | EQ Preset | |
| 0x0032FCC8 | `Electronic` | EQ Preset | |
| 0x0032FCDC | `Hip Hop` | EQ Preset | |
| 0x0032FCF4 | `Loudness` | EQ Preset | |
| 0x0032FD00 | `Lounge` | EQ Preset | |
| 0x0032FD24 | `Small Speakers` | EQ Preset | |
| 0x0032FD34 | `Spoken Word` | EQ Preset | |
| 0x0032FD40 | `Treble Booster` | EQ Preset | |
| 0x0032FD60 | `Vocal Booster` | EQ Preset | |
| 0x004B2BDC | `Acoustic` | EQ Preset | |
| 0x004B2BE8 | `Bass Booster` | EQ Preset | |
| 0x004B2C08 | `Classical` | EQ Preset | |
| 0x004B2C24 | `Electronic` | EQ Preset | |
| 0x004B2C38 | `Hip Hop` | EQ Preset | |
| 0x004B2C50 | `Loudness` | EQ Preset | |
| 0x004B2C5C | `Lounge` | EQ Preset | |
| 0x004B2C80 | `Small Speakers` | EQ Preset | |
| 0x004B2C90 | `Spoken Word` | EQ Preset | |
| 0x004B2C9C | `Treble Booster` | EQ Preset | |
| 0x004B2CBC | `Vocal Booster` | EQ Preset | |

---

## 24. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00104B9C | `Error-SDriver` | Diagnostic | |
| 0x00104BAC | `Error-AClient` | Diagnostic | |
| 0x00105688 | `Root Hub Driver Internal Error unused case in hub handler` | Diagnostic | |
| 0x001056C4 | `Root hub Error Calling Add Device` | Diagnostic | |
| 0x0010A344 | `Error inside %s` | Diagnostic | |
| 0x0013FEE8 | `%s Error in file %s.` | Diagnostic | |
| 0x00266794 | `Error inside %s` | Diagnostic | |
| 0x00266824 | `Error inside %s` | Diagnostic | |
| 0x002668A8 | `Error inside %s` | Diagnostic | |
| 0x00266D60 | `Error inside %s` | Diagnostic | |
| 0x00266E24 | `Error inside %s` | Diagnostic | |
| 0x00266EF0 | `Error inside %s` | Diagnostic | |
| 0x002671AC | `Error inside %s` | Diagnostic | |
| 0x0026739C | `Error inside %s` | Diagnostic | |
| 0x00267400 | `Error inside %s` | Diagnostic | |
| 0x00267534 | `Error inside %s` | Diagnostic | |
| 0x0026758C | `Error inside %s` | Diagnostic | |
| 0x002675DC | `Error inside %s` | Diagnostic | |
| 0x002676AC | `Error inside %s` | Diagnostic | |
| 0x002676FC | `Error inside %s` | Diagnostic | |
| 0x00267AF0 | `Error inside %s` | Diagnostic | |
| 0x00267B60 | `Error inside %s` | Diagnostic | |
| 0x00267FE0 | `Error inside %s` | Diagnostic | |
| 0x00268448 | `Error inside %s` | Diagnostic | |
| 0x0026864C | `Error inside %s` | Diagnostic | |
| 0x002686BC | `Error inside %s` | Diagnostic | |
| 0x002687BC | `Error inside %s` | Diagnostic | |
| 0x002688C4 | `Error inside %s` | Diagnostic | |
| 0x00268938 | `Error inside %s` | Diagnostic | |
| 0x00268984 | `Error inside %s` | Diagnostic | |

---
