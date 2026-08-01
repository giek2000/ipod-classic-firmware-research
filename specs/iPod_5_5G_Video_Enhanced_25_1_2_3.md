# iPod 5.5G (Video Enhanced 80GB) - RetailOS 1.2.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.2.3 |
| **IPSW** | iPod_25.1.2.3.ipsw |
| **Device** | iPod 5.5G (Video Enhanced 80GB) (2006, Click Wheel, Search, Brighter Display) |
| **UpdaterFamilyID** | 25 |
| **Binary Size** | 13,893,632 bytes (13.25 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 13,893,632 bytes |
| **Total Strings (>=6)** | 30,023 |
| **Function Prologues** | 23,090 (ARM: 12,745, Thumb: 10,345) |
| **SoC** | PortalPlayer PP5022C |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Encrypted** | No |
| **SHA-256** | `13b8e5cff4f4d4771a6b9bd8f13257763b65b51e88f46347fda8694899a5dc61` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016AEF0 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x0020EA48 | `Channel UnitTests` | Hidden | Developer Tool |

---

## 2. Controllers (TSilver/TC Classes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00AE5078 | `TCC_Current_Task_Pointer` | Known | UI controller |
| 0x00AE5091 | `TCC_Delete_HISR` | Known | UI controller |
| 0x00AE50B1 | `TCC_Relinquish` | Known | UI controller |
| 0x00AE50CF | `TCC_Resume_Service` | Known | UI controller |
| 0x00AE50E2 | `TCC_Task_Sleep` | Known | UI controller |
| 0x00AE5104 | `TCF_Task_Information` | Known | UI controller |
| 0x00AE5119 | `TCS_Change_Preemption` | Known | UI controller |
| 0x00AE512F | `TCS_Change_Priority` | Known | UI controller |
| 0x00AE5143 | `TCT_Activate_HISR` | Known | UI controller |
| 0x00AE5155 | `TCT_Control_Interrupts` | Known | UI controller |
| 0x00AE516C | `TCT_Local_Control_Interrupts` | Known | UI controller |
| 0x00B4C91F | `TCC_Resume_Service` | Known | UI controller |
| 0x00B4CAA2 | `TCC_Delete_HISR` | Known | UI controller |
| 0x00B4CAF1 | `TCT_Activate_HISR` | Known | UI controller |
| 0x00B4CB1B | `TCT_Control_Interrupts` | Known | UI controller |
| 0x00B4CCE5 | `TCC_Current_Task_Pointer` | Known | UI controller |
| 0x00B4CCFE | `TCS_Change_Priority` | Known | UI controller |
| 0x00B4CD6C | `TCC_Task_Sleep` | Known | UI controller |
| 0x00B4CEA5 | `TCF_Task_Information` | Known | UI controller |
| 0x00B5B41D | `TCC_Relinquish` | Known | UI controller |
| 0x00B5B533 | `TCT_Local_Control_Interrupts` | Known | UI controller |
| 0x00B5B5D3 | `TCC_Task_Sleep` | Known | UI controller |
| 0x00B5B6A1 | `TCS_Change_Preemption` | Known | UI controller |

---

## 3. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009908C | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x000C94F8 | `USB Secondary Interrupt Task` | Known | RTOS task thread |
| 0x000E20A8 | `ICAPTPCameraIOTask` | Known | RTOS task thread |
| 0x0011C548 | `USBStatusTask` | Known | RTOS task thread |
| 0x0011C564 | `USBTaskTimeTask` | Known | RTOS task thread |
| 0x00165034 | `VCUpdateTask` | Known | RTOS task thread |
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
| 0x00AE5068 | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00AE50A1 | `TCC_Delete_Task` | Known | RTOS task thread |
| 0x00AE50C0 | `TCC_Reset_Task` | Known | RTOS task thread |
| 0x00AE50F1 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B4C910 | `TCC_Reset_Task` | Known | RTOS task thread |
| 0x00B4C947 | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B4CAB2 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B4CAC5 | `TCC_Delete_Task` | Known | RTOS task thread |
| 0x00B5B40D | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B5B457 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B5B46A | `TCC_Delete_Task` | Known | RTOS task thread |

---

## 4. Logging Channels

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

## 5. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001758D0 | `AudioCodecs` | Known | Audio system |
| 0x0017695C | `VideoCodecs` | Known | Audio system |
| 0x002B2322 | `.net codec i dette produkt bruges i henhold til en licensaftale fra VoiceAge Cor` | Known | Audio system |
| 0x002B9191 | `.net Codec in diesem Produkt wird in Lizenz der VoiceAge Corporation verwendet. ` | Known | Audio system |
| 0x002C1A10 | `.net codec ` | Known | Audio system |
| 0x002D759C | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x002DE57D | `.net codec a VoiceAge Coporation c` | Known | Audio system |
| 0x002E51B6 | ` utilizzato su licenza da VoiceAge Corporation. Il codec ACELP` | Known | Audio system |
| 0x002F4132 | `.net codec` | Known | Audio system |
| 0x002FB198 | `.net-codec in dit product wordt gebruikt in licentie van VoiceAge Corporation. G` | Known | Audio system |
| 0x0030EDAD | `a da VoiceAge Corporation. O codec ACELP` | Known | Audio system |
| 0x00325671 | `.net codec'i VoiceAge Corporation lisans` | Known | Audio system |
| 0x004B5E2D | `.net codec in this product is used under license from VoiceAge Corporation. Port` | Known | Audio system |
| 0x0067DCB2 | `msCodeCom` | Known | Audio system |
| 0x00B4C39D | `codec_string` | Known | Audio system |
| 0x00B4C3AA | `codec_name` | Known | Audio system |
| 0x00B5ACC9 | `codec_string` | Known | Audio system |
| 0x00B5ACD6 | `codec_name` | Known | Audio system |

---

## 6. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001759AC | `Audible` | Known | Audible audiobook format |
| 0x002AB791 | ` Audible v` | Known | Audible audiobook format |
| 0x002AB7E3 | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x002AB7F9 | ` 2002 Audible, Inc. V` | Known | Audible audiobook format |
| 0x002B21D0 | `Audible-softwaren i dette produkt bruges i henhold til en licensaftale fra Audib` | Known | Audible audiobook format |
| 0x002B2230 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x002B904C | `Die Audible Software in diesem Produkt wird in Lizenz der Audible verwendet. Cop` | Known | Audible audiobook format |
| 0x002B90A5 | ` 2002 Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x002C17FF | ` Audible ` | Known | Audible audiobook format |
| 0x002C185C | ` Audible. ` | Known | Audible audiobook format |
| 0x002C1892 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002C93EC | `El software Audible incluido en este producto se usa bajo licencia de Audible. C` | Known | Audible audiobook format |
| 0x002C9447 | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x002CFDF2 | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright ` | Known | Audible audiobook format |
| 0x002CFE2C | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x002D748C | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x002D74D6 | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x002D74EB | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x002DE43E | ` Audible szoftver az Audible licence alatt van haszn` | Known | Audible audiobook format |
| 0x002DE488 | ` 2002, Audible, Inc. Minden jog fenntartva.` | Known | Audible audiobook format |
| 0x002E50EC | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x002E5115 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x002E5144 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x002EC97D | ` Audible ` | Known | Audible audiobook format |
| 0x002EC99E | `Audible ` | Known | Audible audiobook format |
| 0x002EC9F7 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002F3FE3 | ` Audible ` | Known | Audible audiobook format |
| 0x002F3FFE | ` Audible` | Known | Audible audiobook format |
| 0x002F4042 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002FB050 | `De Audible-software in dit product wordt gebruikt in licentie van Audible. Copyr` | Known | Audible audiobook format |
| 0x002FB0A7 | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x003018CC | `Audible-programvaren i dette produktet brukes under lisens fra Audible. Copyrigh` | Known | Audible audiobook format |
| 0x00301920 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x003081C4 | `Oprogramowanie Audible w tym produkcie jest wykorzystywane na podstawie licencji` | Known | Audible audiobook format |
| 0x00308230 | ` 2002 Audible, Inc. Wszystkie prawa zastrze` | Known | Audible audiobook format |
| 0x0030EC9C | `O software Audible ` | Known | Audible audiobook format |
| 0x0030ECD2 | `a da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0030ECEC | ` 2002 da Audible, Inc. Reservados todos os direitos.` | Known | Audible audiobook format |
| 0x00317425 | ` Audible ` | Known | Audible audiobook format |
| 0x00317477 | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x0031748D | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x0031EC00 | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x0031EC2F | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x0031EC46 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x00325528 | `ndeki Audible yaz` | Known | Audible audiobook format |
| 0x00325541 | ` Audible lisans` | Known | Audible audiobook format |
| 0x00325576 | ` 2002, Audible, Inc. T` | Known | Audible audiobook format |
| 0x0032C1A3 | ` Audible ` | Known | Audible audiobook format |
| 0x0032C1B5 | ` Audible ` | Known | Audible audiobook format |
| 0x0032C1D9 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x00332C70 | `Audible ` | Known | Audible audiobook format |
| 0x00332C84 | ` Audible ` | Known | Audible audiobook format |
| 0x00332CAE | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x004B5CF4 | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x004B5D49 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |

---

## 7. Audio/Codec - Apple Lossless

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00175980 | `AppleLossless` | Known | Apple Lossless codec |
| 0x002DE9B8 | `l alacsony.` | Known | Apple Lossless codec |

---

## 8. Audio/Codec - AAC

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00AF6060 | `AACDEC  VLL ` | Known | AAC codec |
| 0x00AFF7A8 | `MPEG4 AAC LC Decoder` | Known | AAC codec |
| 0x00B01C0D | `AACDecoderGetMem` | Known | AAC codec |
| 0x00B01C1E | `AACDecoderInit` | Known | AAC codec |
| 0x00B01C2D | `AACDecoderGetConfig` | Known | AAC codec |
| 0x00B01C41 | `AACDecoderSetConfig` | Known | AAC codec |
| 0x00B01C55 | `AACHeaderDecode` | Known | AAC codec |
| 0x00B01C65 | `AACDecode` | Known | AAC codec |
| 0x00B01C6F | `AACDecoderInit_Ittiam` | Known | AAC codec |
| 0x00B01C85 | `AACDecoderGetConfig_Ittiam` | Known | AAC codec |
| 0x00B01CA0 | `AACDecoderSetConfig_Ittiam` | Known | AAC codec |
| 0x00B01CBB | `AACHeaderDecode_Ittiam` | Known | AAC codec |
| 0x00B01CD2 | `AACDecode_Ittiam` | Known | AAC codec |

---

## 9. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002AB9D6 | ` zvuku MPEG Layer-3 byla poskytnuta spole` | Known | MP3 codec |
| 0x002ABA01 | `nostmi Fraunhofer IIS a` | Known | MP3 codec |
| 0x002B23CC | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunhofer IIS og THOMSON multi` | Known | MP3 codec |
| 0x002B924F | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und THOMSON Multimedia.` | Known | MP3 codec |
| 0x002C1B57 | ` MPEG Layer-3 ` | Known | MP3 codec |
| 0x002C1B95 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x002C95E5 | `n de audio MPEG Layer-3 utilizada bajo licencia de Fraunhofer IIS y THOMSON mult` | Known | MP3 codec |
| 0x002CFF98 | `MPEG Layer-3 -` | Known | MP3 codec |
| 0x002CFFAA | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | MP3 codec |
| 0x002D76A8 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | MP3 codec |
| 0x002DE610 | `Az MPEG Layer-3 hangk` | Known | MP3 codec |
| 0x002DE638 | `gia a Fraunhofer IIS ` | Known | MP3 codec |
| 0x002E528C | `Tecnologia di codifica audio MPEG Layer-3 su licenza da Fraunhofer IIS e THOMSON` | Known | MP3 codec |
| 0x002ECBAC | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002ECBF8 | `Fraunhofer IIS ` | Known | MP3 codec |
| 0x002F41CC | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002F41F3 | ` Fraunhofer IIS` | Known | MP3 codec |
| 0x002FB234 | `Technologie voor codering van MPEG Layer-3-audio in licentie van Fraunhofer IIS ` | Known | MP3 codec |
| 0x00301A9C | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fraunhofer IIS og THOMSON m` | Known | MP3 codec |
| 0x003083D8 | `Technologia kodowania audio MPEG Layer-3 licencjonowana od Fraunhofer IIS oraz T` | Known | MP3 codec |
| 0x0030EE8E | `udio MPEG Layer-3 licenciada pela Fraunhofer IIS e THOMSON multimedia.` | Known | MP3 codec |
| 0x003176AC | `MPEG Layer-3: ` | Known | MP3 codec |
| 0x00317705 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x0031EDE0 | `Ljudkodningstekniken MPEG Layer-3 ` | Known | MP3 codec |
| 0x0031EE16 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | MP3 codec |
| 0x00325700 | `MPEG Layer-3 ses kodlama teknolojisi Fraunhofer IIS ve THOMSON multimedia'dan li` | Known | MP3 codec |
| 0x0032C330 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x0032C352 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x00332E0C | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00332E31 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x004B5EC0 | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |

---

## 10. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E6544 | `drmsdrmisinffniscpsap@-` | Known | DRM system |
| 0x001758A4 | `AppleDRMVersion` | Known | DRM system |
| 0x00175944 | `AppleDRM` | Known | DRM system |
| 0x00176970 | `AppleVideoDRM` | Known | DRM system |
| 0x0017BA3C | `drmsmp4aesdsmp4v` | Known | DRM system |
| 0x001C7624 | `tkhdedtselstmdiamdhdminfstblstsdstcoco64stscstszsttsstssdrmidrms` | Known | DRM system |
| 0x0067ADCF | `DRMLevel` | Known | DRM system |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E25D4 | `games_RO` | Known | Game system |
| 0x000E25E0 | `gamedata_RW` | Known | Game system |

---

## 12. Photo System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009888C | `Photo Database` | Known | Photo system |
| 0x000B90E8 | `Photos\Photo Database` | Known | Photo system |
| 0x000C08CC | `Photo Database` | Known | Photo system |
| 0x001967D8 | `23iUPhoto Database` | Known | Photo system |
| 0x00198830 | `Photo Database` | Known | Photo system |
| 0x00198B94 | `Photo Database` | Known | Photo system |
| 0x00198E40 | `Photo Import Database` | Known | Photo system |
| 0x00210634 | `Photo Database Size` | Known | Photo system |

---

## 13. Video System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00B10700 | `H.264 Video Decoder` | Known | Video system |

---

## 14. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000B90DC | `iTunesDB` | Known | iTunes database |
| 0x001BC9F4 | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x00206B58 | `iTunes Image DB` | Known | iTunes database |
| 0x002A4ED4 | `iTunesDB` | Known | iTunes database |
| 0x002AAF71 | ` z iTunes nebo vCards. ` | Known | iTunes database |
| 0x002AB28B | `ipojte iPod k iTunes a instalujte hru znovu.` | Known | iTunes database |
| 0x002AB31B | `ipojte iPod k iTunes a zkop` | Known | iTunes database |
| 0x002AB3BC | `i a iTunes ho odemkne.` | Known | iTunes database |
| 0x002AB4D8 | `m iTunes.` | Known | iTunes database |
| 0x002ABE61 | `m iTunes.` | Known | iTunes database |
| 0x002AC81F | `es iTunes.` | Known | iTunes database |
| 0x002AC8C5 | `es iTunes.` | Known | iTunes database |
| 0x002B19C4 | `iPod kan opbevare og vise kontaktoplysninger importeret fra iTunes eller vCards.` | Known | iTunes database |
| 0x002B1D4C | `Slut iPod til iTunes, og installer spillet igen.` | Known | iTunes database |
| 0x002B1DE4 | `Slut iPod til iTunes, og overf` | Known | iTunes database |
| 0x002B1E7B | `slutte iPod til computeren, hvorefter iTunes l` | Known | iTunes database |
| 0x002B1F30 | `r fotografier til computeren, og synkroniser via iTunes for at vise dem p` | Known | iTunes database |
| 0x002B2798 | `%s er for gammel til denne iPod. Slut iPod til computeren, og start iTunes for a` | Known | iTunes database |
| 0x002B301A | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x002B30A9 | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x002B87C7 | `nnen Kontakte (mit iTunes importiert oder vCards) auf Ihrem iPod sichern und anz` | Known | iTunes database |
| 0x002B8B58 | `Verbinden Sie Ihren iPod mit iTunes und installieren Sie das Spiel erneut.` | Known | iTunes database |
| 0x002B8C1C | `Verbinden Sie Ihren iPod mit iTunes und laden Sie die aktuelle Version.` | Known | iTunes database |
| 0x002B8CC3 | `en Sie Ihren iPod an Ihren Computer an und iTunes deaktiviert die Anzeigensperre` | Known | iTunes database |
| 0x002B8D6C | `Importierte Fotos werden nicht auf dem TV angezeigt. Senden Sie sie erst an den ` | Known | iTunes database |
| 0x002B9671 | `en Sie den iPod an den Computer an und starten Sie iTunes, um %s auf die aktuell` | Known | iTunes database |
| 0x002B9F9D | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x002BA03E | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x002C0A78 | ` iTunes ` | Known | iTunes database |
| 0x002C0FC1 | ` iTunes ` | Known | iTunes database |
| 0x002C10E9 | ` iTunes ` | Known | iTunes database |
| 0x002C1212 | ` iTunes ` | Known | iTunes database |
| 0x002C13A7 | ` iTunes ` | Known | iTunes database |
| 0x002C2286 | ` iTunes ` | Known | iTunes database |
| 0x002C33FA | ` iTunes ` | Known | iTunes database |
| 0x002C3523 | ` iTunes ` | Known | iTunes database |
| 0x002C8BBD | `n importada de iTunes o de tarjetas virtuales (vCards). ` | Known | iTunes database |
| 0x002C8F48 | `Conecte el iPod a iTunes y reinstale el juego.` | Known | iTunes database |
| 0x002C8FE0 | `Conecte el iPod a iTunes y descargue la versi` | Known | iTunes database |
| 0x002C9068 | `n, conecte el iPod al ordenador y iTunes lo desbloquear` | Known | iTunes database |
| 0x002C9144 | `celas con iTunes para verlas en la TV.` | Known | iTunes database |
| 0x002C99F8 | `%s es demasiado antiguo para ejecutarse en este iPod. Conecte el iPod al ordenad` | Known | iTunes database |
| 0x002CA348 | `Esta foto es demasiado grande para mostrarla en el iPod. Transfiera las fotos im` | Known | iTunes database |
| 0x002CA3EC | `Este formato de foto no puede visualizarse en el iPod. Transfiera las fotos impo` | Known | iTunes database |
| 0x002CF634 | ` iTunesista tai vCardeina tuotua tietoa. ` | Known | iTunes database |
| 0x002CF99A | ` iPod iTunesiin ja asenna peli uudelleen.` | Known | iTunes database |
| 0x002CFA2A | ` iPod iTunesiin ja hae uusin versio.` | Known | iTunes database |
| 0x002CFA9F | `tietokoneeseen, niin iTunes avaa lukituksen.` | Known | iTunes database |
| 0x002CFB3C | ` kuvat tietokoneelle ja synkronoi ne iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002D03A6 | ` %s uusimpaan versioon avaamalla iTunes.` | Known | iTunes database |
| 0x002D0C32 | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002D0CB9 | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x002D6B96 | `iTunes ou de vCards. ` | Known | iTunes database |
| 0x002D6F44 | `Connectez votre iPod avec iTunes et r` | Known | iTunes database |
| 0x002D6FEC | `Connectez votre iPod avec iTunes et t` | Known | iTunes database |
| 0x002D70AD | ` votre ordinateur et iTunes le d` | Known | iTunes database |
| 0x002D7190 | `rez-les sur votre ordinateur puis synchronisez-les avec iTunes.` | Known | iTunes database |
| 0x002D7ACC | `ordinateur et lancez iTunes pour mettre ` | Known | iTunes database |
| 0x002D8586 | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x002D8643 | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x002DDAC8 | `Az iPod, az iTunes programb` | Known | iTunes database |
| 0x002DDEE1 | `t az iTunes programhoz, ` | Known | iTunes database |
| 0x002DDFB1 | `t az iTunes programhoz ` | Known | iTunes database |
| 0x002DE05F | `phez, hogy az iTunes feloldja a z` | Known | iTunes database |
| 0x002DE16D | `ljon az iTunes haszn` | Known | iTunes database |
| 0x002DEAB0 | `s az iTunes futtat` | Known | iTunes database |
| 0x002DF537 | `lja azokat az iTunes programmal.` | Known | iTunes database |
| 0x002DF60E | `lja azokat az iTunes programmal.` | Known | iTunes database |
| 0x002E485C | ` memorizzare e visualizzare informazioni importate da iTunes o vCards. ` | Known | iTunes database |
| 0x002E4C24 | `Collega iPod a iTunes e reinstalla il gioco.` | Known | iTunes database |
| 0x002E4CC8 | `Collega  iPod a iTunes ed esegui il download dell'ultima versione.` | Known | iTunes database |
| 0x002E4D40 | `Se dimentichi la combinazione, collega iPod al computer e iTunes sar` | Known | iTunes database |
| 0x002E4DE0 | `Le foto importate non possono visualizzarsi in TV. Trasferisci le foto sul compu` | Known | iTunes database |
| 0x002E5635 | ` troppo vecchio per funzionare con questo iPod. Collega iPod al computer ed eseg` | Known | iTunes database |
| 0x002E5F41 | ` troppo grande per essere visualizzato in iPod. Trasferisci le foto sul computer` | Known | iTunes database |
| 0x002E5FEB | ` essere visualizzato in iPod. Trasferisci le foto sul computer e sincronizzale t` | Known | iTunes database |
| 0x002EBE6B | `iTunes ` | Known | iTunes database |
| 0x002EC2FD | `iTunes ` | Known | iTunes database |
| 0x002EC409 | `iTunes ` | Known | iTunes database |
| 0x002EC509 | `iTunes` | Known | iTunes database |
| 0x002EC634 | `iTunes ` | Known | iTunes database |
| 0x002ED0BB | `iTunes ` | Known | iTunes database |
| 0x002EDC4F | `iTunes ` | Known | iTunes database |
| 0x002EDD15 | `iTunes ` | Known | iTunes database |
| 0x002F36CA | ` iTunes ` | Known | iTunes database |
| 0x002F3AA3 | ` iTunes` | Known | iTunes database |
| 0x002F3B6B | ` iTunes` | Known | iTunes database |
| 0x002F3C42 | ` iTunes` | Known | iTunes database |
| 0x002F3D25 | ` iTunes` | Known | iTunes database |
| 0x002F462F | ` iTunes` | Known | iTunes database |
| 0x002F5016 | ` iTunes` | Known | iTunes database |
| 0x002F50BB | ` iTunes` | Known | iTunes database |
| 0x002FA851 | `mporteerd uit iTunes of vCards. ` | Known | iTunes database |
| 0x002FABE0 | `Verbind de iPod met iTunes en installeer het spel opnieuw.` | Known | iTunes database |
| 0x002FAC80 | `Verbind de iPod met iTunes en download de nieuwste versie.` | Known | iTunes database |
| 0x002FACF4 | `Als u de combinatie bent vergeten, verbind iPod met uw computer en iTunes zal he` | Known | iTunes database |
| 0x002FADA5 | `mporteerde foto's op tv niet mogelijk. Kopieer foto's naar de computer en synchr` | Known | iTunes database |
| 0x002FB65C | `%s is te oud om op deze iPod te worden gebruikt. Sluit de iPod aan op de compute` | Known | iTunes database |
| 0x002FBF47 | `mporteerde foto's naar de computer en synchroniseer ze met iTunes voor weergave ` | Known | iTunes database |
| 0x002FBFDE | `mporteerde foto's naar de computer en synchroniseer ze met iTunes voor weergave ` | Known | iTunes database |
| | *...and 72 more* | | |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002B710C | `Radio-Region` | Known | FM Radio |
| 0x004B3A0C | `Radio Region` | Known | FM Radio |
| 0x004B6418 | `Radio Region` | Known | FM Radio |
| 0x004B69DC | `Radio Settings` | Known | FM Radio |

---

## 16. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00175DDC | `FireWireGUID` | Known | FireWire interface |
| 0x00175DEC | `FireWireVersion` | Known | FireWire interface |
| 0x001763D8 | `FireWire` | Known | FireWire interface |
| 0x002AADAF | ` FireWire nen` | Known | FireWire interface |
| 0x002ACDD8 | `FireWire p` | Known | FireWire interface |
| 0x002B1828 | `FireWire-forbindelser underst` | Known | FireWire interface |
| 0x002B3574 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x002B8620 | `FireWire wird nicht unterst` | Known | FireWire interface |
| 0x002BA522 | `ber FireWire verbunden` | Known | FireWire interface |
| 0x002C079A | ` FireWire. ` | Known | FireWire interface |
| 0x002C3DA2 | ` FireWire` | Known | FireWire interface |
| 0x002C89FD | `de canciones o archivos no son posibles con FireWire :` | Known | FireWire interface |
| 0x002CA954 | `FireWire conectado` | Known | FireWire interface |
| 0x002CF474 | `FireWire-tiedonsiirtoa ei tueta. Siirt` | Known | FireWire interface |
| 0x002D11A0 | `FireWire liitetty` | Known | FireWire interface |
| 0x002D6999 | `s via FireWire : connectez l` | Known | FireWire interface |
| 0x002D8BE8 | `FireWire Connect` | Known | FireWire interface |
| 0x002DD910 | `A FireWire kapcsolat nem t` | Known | FireWire interface |
| 0x002DFB44 | `FireWire csatlakozik` | Known | FireWire interface |
| 0x002E46A8 | `Connessioni di dati via FireWire non sono supportate. Per trasferire brani o dat` | Known | FireWire interface |
| 0x002E6534 | `FireWire connesso` | Known | FireWire interface |
| 0x002EBBE8 | `FireWire ` | Known | FireWire interface |
| 0x002EE340 | `FireWire ` | Known | FireWire interface |
| 0x002F34DC | `FireWire ` | Known | FireWire interface |
| 0x002F55DC | `FireWire ` | Known | FireWire interface |
| 0x002FA69E | `ren via FireWire, maar alleen via de meegeleverde USB-kabel.` | Known | FireWire interface |
| 0x002FC4D8 | `FireWire aangesloten` | Known | FireWire interface |
| 0x00300F3F | `ring via FireWire st` | Known | FireWire interface |
| 0x00302C08 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x003077EB | `czenie FireWire nie jest wspierane. By przes` | Known | FireWire interface |
| 0x0030973B | `czone przez FireWire` | Known | FireWire interface |
| 0x0030E31F | `es FireWire n` | Known | FireWire interface |
| 0x0031024C | `FireWire ligado` | Known | FireWire interface |
| 0x00316315 | ` FireWire ` | Known | FireWire interface |
| 0x0031975B | ` FireWire` | Known | FireWire interface |
| 0x0031E278 | `FireWire-` | Known | FireWire interface |
| 0x0031FFBC | `FireWire anslutet` | Known | FireWire interface |
| 0x00324AA8 | `FireWire ba` | Known | FireWire interface |
| 0x003269EC | `FireWire Ba` | Known | FireWire interface |
| 0x0032B819 | ` FireWire ` | Known | FireWire interface |
| 0x0032D47C | `FireWire ` | Known | FireWire interface |
| 0x003322C9 | ` FireWire ` | Known | FireWire interface |
| 0x00334024 | `FireWire ` | Known | FireWire interface |
| 0x004B5380 | `FireWire connections are not supported. To transfer songs, connect the USB cable` | Known | FireWire interface |
| 0x004B7050 | `FireWire Connected` | Known | FireWire interface |

---

## 17. USB

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x006CA5BD | `USBCompositeDevice1.6` | Known | USB interface |
| 0x006CA615 | `USBCompositeDevice1.6` | Known | USB interface |

---

## 18. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F4A88 | `LCD Module could not be determined.` | Known | Hardware interface |
| 0x00176880 | `ForcedDiskMode` | Known | Hardware interface |
| 0x00210594 | `Enter Disk Mode` | Known | Hardware interface |
| 0x002105A4 | `Exit Disk Mode` | Known | Hardware interface |
| 0x004B5374 | `Disk Mode` | Known | Hardware interface |
| 0x004FE8F0 | `I2C write Error` | Known | Hardware interface |
| 0x004FE904 | `I2C read Error %02x` | Known | Hardware interface |
| 0x0067FB9D | `OCSP_RESPID` | Known | Hardware interface |

---

## 19. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0014B0D4 | `PowerManager` | Known | Power management |
| 0x001763B4 | `PowerInformation` | Known | Power management |
| 0x00210600 | `Begin Charging` | Known | Power management |
| 0x00210610 | `Stop Charging` | Known | Power management |
| 0x00259038 | `USBPowerSense` | Known | Power management |
| 0x002590F8 | `PCFPowerMgr` | Known | Power management |
| 0x004B4F1C | `Charging` | Known | Power management |
| 0x004B709C | `Low Battery` | Known | Power management |

---

## 20. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002AFE38 | `Alarmer` | Known | UI element |
| 0x002B34B8 | `Alarmer` | Known | UI element |
| 0x002C6E40 | `Calendario` | Known | UI element |
| 0x002C6E4C | `Calendarios` | Known | UI element |
| 0x002C6E58 | `Calendarios` | Known | UI element |
| 0x002C6E94 | `Alarmas` | Known | UI element |
| 0x002C79B8 | `Calendario` | Known | UI element |
| 0x002C79C4 | `Calendarios` | Known | UI element |
| 0x002C8EB8 | `Alarma` | Known | UI element |
| 0x002C9B04 | `Alarma` | Known | UI element |
| 0x002C9B74 | `Alarma` | Known | UI element |
| 0x002C9EFA | `Calendario` | Known | UI element |
| 0x002CA198 | `Alarma` | Known | UI element |
| 0x002CA84C | `Alarma` | Known | UI element |
| 0x002CA89C | `Alarmas` | Known | UI element |
| 0x002D4C80 | `Alarmes` | Known | UI element |
| 0x002D6E88 | `Alarme` | Known | UI element |
| 0x002D7BB0 | `Alarme` | Known | UI element |
| 0x002D7C18 | `Alarme` | Known | UI element |
| 0x002D82F0 | `Alarme` | Known | UI element |
| 0x002D8A98 | `Alarme` | Known | UI element |
| 0x002D8B10 | `Alarmes` | Known | UI element |
| 0x002E2B5C | `Calendario` | Known | UI element |
| 0x002E2B68 | `Calendari` | Known | UI element |
| 0x002E2B74 | `Calendari` | Known | UI element |
| 0x002E3690 | `Calendario` | Known | UI element |
| 0x002E369C | `Calendari` | Known | UI element |
| 0x002E5AF7 | `Calendario` | Known | UI element |
| 0x002FF520 | `Alarmer` | Known | UI element |
| 0x003024E4 | `Alarmtidspunkt` | Known | UI element |
| 0x00302B44 | `Alarmer` | Known | UI element |
| 0x00305C84 | `Alarmy` | Known | UI element |
| 0x0030601C | `Gotowe` | Known | UI element |
| 0x00306200 | `Gotowe` | Known | UI element |
| 0x00309658 | `Alarmy` | Known | UI element |
| 0x0030C7D4 | `Alarmes` | Known | UI element |
| 0x0030E7A0 | `Alarme` | Known | UI element |
| 0x0030F37C | `Alarme` | Known | UI element |
| 0x0031018C | `Alarmes` | Known | UI element |
| 0x0031F89C | `Alarmtid` | Known | UI element |
| 0x00323008 | `Alarmlar` | Known | UI element |
| 0x00326234 | `Alarm Zaman` | Known | UI element |
| 0x00326930 | `Alarmlar` | Known | UI element |
| 0x004B3458 | `Calendar` | Known | UI element |
| 0x004B3464 | `Calendars` | Known | UI element |
| 0x004B3470 | `Calendars` | Known | UI element |
| 0x004B34A4 | `Alarms` | Known | UI element |
| 0x004B3F40 | `Calendar` | Known | UI element |
| 0x004B3F4C | `Calendars` | Known | UI element |
| 0x004B63A0 | `Alarm Clock` | Known | UI element |
| 0x004B66B2 | `Calendar` | Known | UI element |
| 0x004B6924 | `Alarm Time` | Known | UI element |
| 0x004B6930 | `Alarm Clock` | Known | UI element |
| 0x004B6F38 | `Alarm Clock` | Known | UI element |
| 0x004B6FA4 | `Alarms` | Known | UI element |
| 0x004B71A4 | `GotoBackToIdleCommand` | Known | UI element |
| 0x00502520 | `TCalendarCntlr_Alarm` | Known | UI element |
| 0x0067AD30 | `Calendars/` | Known | UI element |
| 0x0067AD4B | `Calendars` | Known | UI element |

---

## 21. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0007E490 | `Settings` | Known | Menu item |
| 0x002A9C38 | `Podcasts` | Known | Menu item |
| 0x002A9DA4 | `Podcasts` | Known | Menu item |
| 0x002AC3C6 | `Podcasts` | Known | Menu item |
| 0x002ACD38 | `Podcasts` | Known | Menu item |
| 0x002B077C | `Podcasts` | Known | Menu item |
| 0x002B08CC | `Podcasts` | Known | Menu item |
| 0x002B2CB5 | `Podcasts` | Known | Menu item |
| 0x002B34E4 | `Podcasts` | Known | Menu item |
| 0x002B752C | `Podcasts` | Known | Menu item |
| 0x002B762C | `Extras` | Known | Menu item |
| 0x002B765C | `Videos` | Known | Menu item |
| 0x002B7694 | `Podcasts` | Known | Menu item |
| 0x002B944C | `Videos` | Known | Menu item |
| 0x002B9B24 | `Extras` | Known | Menu item |
| 0x002B9B2C | `Videos` | Known | Menu item |
| 0x002B9B8F | `Podcasts` | Known | Menu item |
| 0x002B9E64 | `Videos` | Known | Menu item |
| 0x002BA48C | `Podcasts` | Known | Menu item |
| 0x002BA588 | `Extras` | Known | Menu item |
| 0x002BE970 | `Podcasts` | Known | Menu item |
| 0x002BEC40 | `Podcasts` | Known | Menu item |
| 0x002C2C31 | `Podcasts` | Known | Menu item |
| 0x002C3C7C | `Podcasts` | Known | Menu item |
| 0x002C78E8 | `Podcasts` | Known | Menu item |
| 0x002C79E4 | `Extras` | Known | Menu item |
| 0x002C7A48 | `Podcasts` | Known | Menu item |
| 0x002C9F50 | `Extras` | Known | Menu item |
| 0x002C9FC2 | `Podcasts` | Known | Menu item |
| 0x002CA8C4 | `Podcasts` | Known | Menu item |
| 0x002CA9B8 | `Extras` | Known | Menu item |
| 0x002D577C | `Podcasts` | Known | Menu item |
| 0x002D57C4 | `Albums` | Known | Menu item |
| 0x002D57DC | `Genres` | Known | Menu item |
| 0x002D581C | `Photos` | Known | Menu item |
| 0x002D589C | `Extras` | Known | Menu item |
| 0x002D590C | `Podcasts` | Known | Menu item |
| 0x002D5A14 | `Albums` | Known | Menu item |
| 0x002D72E8 | `Photos` | Known | Menu item |
| 0x002D7394 | `Photos` | Known | Menu item |
| 0x002D7888 | `Photos` | Known | Menu item |
| 0x002D805C | `Extras` | Known | Menu item |
| 0x002D8090 | `Photos` | Known | Menu item |
| 0x002D80BE | `Genres` | Known | Menu item |
| 0x002D80E2 | `Podcasts` | Known | Menu item |
| 0x002D8116 | `Albums` | Known | Menu item |
| 0x002D8294 | `Genres` | Known | Menu item |
| 0x002D82A8 | `Albums` | Known | Menu item |
| 0x002D866C | `Photos` | Known | Menu item |
| 0x002D8B24 | `Genres` | Known | Menu item |
| 0x002D8B38 | `Podcasts` | Known | Menu item |
| 0x002D8B54 | `Albums` | Known | Menu item |
| 0x002D8C70 | `Extras` | Known | Menu item |
| 0x002DC764 | `Podcasts` | Known | Menu item |
| 0x002DC8DC | `Podcasts` | Known | Menu item |
| 0x002DF05E | `Podcasts` | Known | Menu item |
| 0x002DFAA4 | `Podcasts` | Known | Menu item |
| 0x002F959C | `Podcasts` | Known | Menu item |
| 0x002F95E0 | `Albums` | Known | Menu item |
| 0x002F95F4 | `Genres` | Known | Menu item |
| 0x002F9708 | `Podcasts` | Known | Menu item |
| 0x002F97E0 | `Albums` | Known | Menu item |
| 0x002FBB7B | `Genres` | Known | Menu item |
| 0x002FBB9B | `Podcasts` | Known | Menu item |
| 0x002FBBC3 | `Albums` | Known | Menu item |
| 0x002FBD04 | `Genres` | Known | Menu item |
| 0x002FBD14 | `Albums` | Known | Menu item |
| 0x002FC42C | `Genres` | Known | Menu item |
| 0x002FC440 | `Podcasts` | Known | Menu item |
| 0x002FC458 | `Albums` | Known | Menu item |
| 0x00306628 | `Podcasts` | Known | Menu item |
| 0x00308D4A | `Podcasts` | Known | Menu item |
| 0x00309680 | `Podcasts` | Known | Menu item |
| 0x0030D1D0 | `Podcasts` | Known | Menu item |
| 0x0030D2F4 | `Extras` | Known | Menu item |
| 0x0030D358 | `Podcasts` | Known | Menu item |
| 0x0030F800 | `Extras` | Known | Menu item |
| 0x0030F87A | `Podcasts` | Known | Menu item |
| 0x003101C4 | `Podcasts` | Known | Menu item |
| 0x00310298 | `Extras` | Known | Menu item |
| 0x00314508 | `Podcasts` | Known | Menu item |
| 0x00314780 | `Podcasts` | Known | Menu item |
| 0x00318689 | `Podcasts` | Known | Menu item |
| 0x00319640 | `Podcasts` | Known | Menu item |
| 0x00323A20 | `Podcasts` | Known | Menu item |
| 0x00323B8C | `Podcasts` | Known | Menu item |
| 0x00326072 | `Podcasts` | Known | Menu item |
| 0x0032695C | `Podcasts` | Known | Menu item |
| 0x00333F98 | `Podcasts` | Known | Menu item |
| 0x004B3D8C | `Podcasts` | Known | Menu item |
| 0x004B3E9C | `Now Playing` | Known | Menu item |
| 0x004B3EA8 | `Artists` | Known | Menu item |
| 0x004B3EC0 | `Albums` | Known | Menu item |
| 0x004B3ED8 | `Genres` | Known | Menu item |
| 0x004B3EE0 | `Composers` | Known | Menu item |
| 0x004B3F0C | `Photos` | Known | Menu item |
| 0x004B3F74 | `Extras` | Known | Menu item |
| 0x004B3F7C | `Playlists` | Known | Menu item |
| 0x004B3F88 | `Audiobooks` | Known | Menu item |
| 0x004B3F9C | `Videos` | Known | Menu item |
| | *...and 39 more* | | |

---

## 22. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000985E8 | `iPod_Control` | Filesystem Path | |
| 0x00098614 | `iPod_Control\Device` | Filesystem Path | |
| 0x000A4B1C | `iPod_Control\Device` | Filesystem Path | |
| 0x000A6508 | `iPod_Control` | Filesystem Path | |
| 0x000A6B60 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x000B90C4 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000B9104 | `System_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x000BBBF0 | `iPod_Control\Music\` | Filesystem Path | |
| 0x000C0194 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x000C0314 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000E2600 | `iPod_Control/%s/%s%s%s` | Filesystem Path | |
| 0x000EF350 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000F6CC4 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000F872C | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000F8828 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001A66CC | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001A7060 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path | |
| 0x001A7088 | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001A71F4 | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001D032C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D05B8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0670 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D07C0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D08E0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D09B0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0B48 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0C04 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0CB4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0DA8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0E4C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0F00 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D0FBC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D10F0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D1260 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D1324 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D13D4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D1510 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D15DC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D16A8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D1770 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D1814 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D18DC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D198C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D1A3C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D1B04 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D1BC4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D1C74 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D1D24 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D1DD4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D1E84 | `iPod_Control\Device\` | Filesystem Path | |

---

## 23. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002AA1FC | `Acoustic` | EQ Preset | |
| 0x002AA208 | `Bass Booster` | EQ Preset | |
| 0x002AA228 | `Classical` | EQ Preset | |
| 0x002AA244 | `Electronic` | EQ Preset | |
| 0x002AA258 | `Hip Hop` | EQ Preset | |
| 0x002AA270 | `Loudness` | EQ Preset | |
| 0x002AA27C | `Lounge` | EQ Preset | |
| 0x002AA2A0 | `Small Speakers` | EQ Preset | |
| 0x002AA2B0 | `Spoken Word` | EQ Preset | |
| 0x002AA2BC | `Treble Booster` | EQ Preset | |
| 0x002AA2DC | `Vocal Booster` | EQ Preset | |
| 0x002B0CCC | `Acoustic` | EQ Preset | |
| 0x002B0CD8 | `Bass Booster` | EQ Preset | |
| 0x002B0CF8 | `Classical` | EQ Preset | |
| 0x002B0D14 | `Electronic` | EQ Preset | |
| 0x002B0D28 | `Hip Hop` | EQ Preset | |
| 0x002B0D40 | `Loudness` | EQ Preset | |
| 0x002B0D4C | `Lounge` | EQ Preset | |
| 0x002B0D70 | `Small Speakers` | EQ Preset | |
| 0x002B0D80 | `Spoken Word` | EQ Preset | |
| 0x002B0D8C | `Treble Booster` | EQ Preset | |
| 0x002B0DAC | `Vocal Booster` | EQ Preset | |
| 0x002B7AC0 | `Acoustic` | EQ Preset | |
| 0x002B7B00 | `Electronic` | EQ Preset | |
| 0x002B7B14 | `Hip Hop` | EQ Preset | |
| 0x002B7B2C | `Loudness` | EQ Preset | |
| 0x002BF398 | `Hip Hop` | EQ Preset | |
| 0x002BF3B0 | `Loudness` | EQ Preset | |
| 0x002BF3BC | `Lounge` | EQ Preset | |
| 0x002C7EE0 | `Hip Hop` | EQ Preset | |
| 0x002C7EF0 | `Latina` | EQ Preset | |
| 0x002C7EF8 | `Loudness` | EQ Preset | |
| 0x002C7F04 | `Lounge` | EQ Preset | |
| 0x002CE96C | `Lounge` | EQ Preset | |
| 0x002D5DE8 | `Hip Hop` | EQ Preset | |
| 0x002D5E18 | `Lounge` | EQ Preset | |
| 0x002E3B70 | `Hip Hop` | EQ Preset | |
| 0x002E3B80 | `Latina` | EQ Preset | |
| 0x002E3B88 | `Loudness` | EQ Preset | |
| 0x002E3B94 | `Lounge` | EQ Preset | |
| 0x002EAA70 | `Acoustic` | EQ Preset | |
| 0x002EAA7C | `Bass Booster` | EQ Preset | |
| 0x002EAA9C | `Classical` | EQ Preset | |
| 0x002EAAB8 | `Electronic` | EQ Preset | |
| 0x002EAACC | `Hip Hop` | EQ Preset | |
| 0x002EAAE4 | `Loudness` | EQ Preset | |
| 0x002EAAF0 | `Lounge` | EQ Preset | |
| 0x002EAB14 | `Small Speakers` | EQ Preset | |
| 0x002EAB24 | `Spoken Word` | EQ Preset | |
| 0x002EAB30 | `Treble Booster` | EQ Preset | |
| 0x002EAB50 | `Vocal Booster` | EQ Preset | |
| 0x002F25F8 | `Acoustic` | EQ Preset | |
| 0x002F2604 | `Bass Booster` | EQ Preset | |
| 0x002F2624 | `Classical` | EQ Preset | |
| 0x002F2640 | `Electronic` | EQ Preset | |
| 0x002F2654 | `Hip Hop` | EQ Preset | |
| 0x002F266C | `Loudness` | EQ Preset | |
| 0x002F2678 | `Lounge` | EQ Preset | |
| 0x002F269C | `Small Speakers` | EQ Preset | |
| 0x002F26AC | `Spoken Word` | EQ Preset | |
| 0x002F26B8 | `Treble Booster` | EQ Preset | |
| 0x002F26D8 | `Vocal Booster` | EQ Preset | |
| 0x002F9B88 | `Loudness` | EQ Preset | |
| 0x002F9B94 | `Lounge` | EQ Preset | |
| 0x00300438 | `Latino` | EQ Preset | |
| 0x00300440 | `Loudness` | EQ Preset | |
| 0x0030044C | `Lounge` | EQ Preset | |
| 0x00306C78 | `Hip Hop` | EQ Preset | |
| 0x00306CAC | `Lounge` | EQ Preset | |
| 0x0030D7E0 | `Hip Hop` | EQ Preset | |
| 0x0030D7F0 | `Latina` | EQ Preset | |
| 0x0030D7F8 | `Loudness` | EQ Preset | |
| 0x0030D804 | `Lounge` | EQ Preset | |
| 0x0031D730 | `Acoustic` | EQ Preset | |
| 0x0031D73C | `Bass Booster` | EQ Preset | |
| 0x0031D75C | `Classical` | EQ Preset | |
| 0x0031D778 | `Electronic` | EQ Preset | |
| 0x0031D78C | `Hip Hop` | EQ Preset | |
| 0x0031D7A4 | `Loudness` | EQ Preset | |
| 0x0031D7B0 | `Lounge` | EQ Preset | |
| 0x0031D7D4 | `Small Speakers` | EQ Preset | |
| 0x0031D7E4 | `Spoken Word` | EQ Preset | |
| 0x0031D7F0 | `Treble Booster` | EQ Preset | |
| 0x0031D810 | `Vocal Booster` | EQ Preset | |
| 0x00323FCC | `Hip Hop` | EQ Preset | |
| 0x00323FE0 | `Loudness` | EQ Preset | |
| 0x00323FEC | `Lounge` | EQ Preset | |
| 0x0032AA68 | `Acoustic` | EQ Preset | |
| 0x0032AA74 | `Bass Booster` | EQ Preset | |
| 0x0032AA94 | `Classical` | EQ Preset | |
| 0x0032AAB0 | `Electronic` | EQ Preset | |
| 0x0032AAC4 | `Hip Hop` | EQ Preset | |
| 0x0032AADC | `Loudness` | EQ Preset | |
| 0x0032AAE8 | `Lounge` | EQ Preset | |
| 0x0032AB0C | `Small Speakers` | EQ Preset | |
| 0x0032AB1C | `Spoken Word` | EQ Preset | |
| 0x0032AB28 | `Treble Booster` | EQ Preset | |
| 0x0032AB48 | `Vocal Booster` | EQ Preset | |
| 0x00331538 | `Acoustic` | EQ Preset | |
| 0x00331544 | `Bass Booster` | EQ Preset | |
| 0x00331564 | `Classical` | EQ Preset | |
| 0x00331580 | `Electronic` | EQ Preset | |
| 0x00331594 | `Hip Hop` | EQ Preset | |
| 0x003315AC | `Loudness` | EQ Preset | |
| 0x003315B8 | `Lounge` | EQ Preset | |
| 0x003315DC | `Small Speakers` | EQ Preset | |
| 0x003315EC | `Spoken Word` | EQ Preset | |
| 0x003315F8 | `Treble Booster` | EQ Preset | |
| 0x00331618 | `Vocal Booster` | EQ Preset | |
| 0x004B4494 | `Acoustic` | EQ Preset | |
| 0x004B44A0 | `Bass Booster` | EQ Preset | |
| 0x004B44C0 | `Classical` | EQ Preset | |
| 0x004B44DC | `Electronic` | EQ Preset | |
| 0x004B44F0 | `Hip Hop` | EQ Preset | |
| 0x004B4508 | `Loudness` | EQ Preset | |
| 0x004B4514 | `Lounge` | EQ Preset | |
| 0x004B4538 | `Small Speakers` | EQ Preset | |
| 0x004B4548 | `Spoken Word` | EQ Preset | |
| 0x004B4554 | `Treble Booster` | EQ Preset | |
| 0x004B4574 | `Vocal Booster` | EQ Preset | |

---

## 24. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00105720 | `Error-SDriver` | Diagnostic | |
| 0x00105730 | `Error-AClient` | Diagnostic | |
| 0x0010620C | `Root Hub Driver Internal Error unused case in hub handler` | Diagnostic | |
| 0x00106248 | `Root hub Error Calling Add Device` | Diagnostic | |
| 0x0010AEC8 | `Error inside %s` | Diagnostic | |
| 0x00141474 | `%s Error in file %s.` | Diagnostic | |
| 0x00267E68 | `Error inside %s` | Diagnostic | |
| 0x00267EF8 | `Error inside %s` | Diagnostic | |
| 0x00267F7C | `Error inside %s` | Diagnostic | |
| 0x00268434 | `Error inside %s` | Diagnostic | |
| 0x002684F8 | `Error inside %s` | Diagnostic | |
| 0x002685C4 | `Error inside %s` | Diagnostic | |
| 0x00268880 | `Error inside %s` | Diagnostic | |
| 0x00268A70 | `Error inside %s` | Diagnostic | |
| 0x00268AD4 | `Error inside %s` | Diagnostic | |
| 0x00268C08 | `Error inside %s` | Diagnostic | |
| 0x00268C60 | `Error inside %s` | Diagnostic | |
| 0x00268CB0 | `Error inside %s` | Diagnostic | |
| 0x00268D80 | `Error inside %s` | Diagnostic | |
| 0x00268DD0 | `Error inside %s` | Diagnostic | |
| 0x002691C4 | `Error inside %s` | Diagnostic | |
| 0x00269234 | `Error inside %s` | Diagnostic | |
| 0x002696B4 | `Error inside %s` | Diagnostic | |
| 0x00269B1C | `Error inside %s` | Diagnostic | |
| 0x00269D20 | `Error inside %s` | Diagnostic | |
| 0x00269D90 | `Error inside %s` | Diagnostic | |
| 0x00269E90 | `Error inside %s` | Diagnostic | |
| 0x00269F98 | `Error inside %s` | Diagnostic | |
| 0x0026A00C | `Error inside %s` | Diagnostic | |
| 0x0026A058 | `Error inside %s` | Diagnostic | |

---
