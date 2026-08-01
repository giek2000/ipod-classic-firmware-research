# iPod 4th Generation (Monochrome) - RetailOS 3.1.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 3.1.1 |
| **IPSW** | iPod_10.3.1.1.ipsw |
| **Device** | iPod 4th Generation (Monochrome) (2004, Click Wheel (first)) |
| **UpdaterFamilyID** | 10 |
| **Binary Size** | 4,605,440 bytes (4.39 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 4,605,440 bytes |
| **Total Strings (>=6)** | 10,743 |
| **Function Prologues** | 10,908 (ARM: 7,824, Thumb: 3,084) |
| **SoC** | PortalPlayer PP5020 |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Encrypted** | No |
| **SHA-256** | `9c959c47dbc17f78a65936c5461749c71af39f36744033881c981ce39cb872ca` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174294 | `MP3ExampleTask` | Hidden | Hidden Test |

---

## 2. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00173E18 | `WatchdogTask` | Known | RTOS task thread |
| 0x00173E28 | `AlarmTask` | Known | RTOS task thread |
| 0x00173E40 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x00173E54 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x00173E64 | `TopPlugTask` | Known | RTOS task thread |
| 0x00173E70 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x00173E80 | `PlayBtnTask` | Known | RTOS task thread |
| 0x00173E8C | `PrvBtnTask` | Known | RTOS task thread |
| 0x00173E98 | `NextBtnTask` | Known | RTOS task thread |
| 0x00173EA4 | `ActionBtnTask` | Known | RTOS task thread |
| 0x00173EB4 | `MenuBtnTask` | Known | RTOS task thread |
| 0x00173EC0 | `DiskMgrTask` | Known | RTOS task thread |
| 0x00173EDC | `CNATask` | Known | RTOS task thread |
| 0x00173EE4 | `BacklightTask` | Known | RTOS task thread |
| 0x00173EF4 | `SerialOptoTask` | Known | RTOS task thread |
| 0x00173F04 | `OptoTask` | Known | RTOS task thread |
| 0x00173F10 | `FirewireTask` | Known | RTOS task thread |
| 0x00174258 | `HostOSTask` | Known | RTOS task thread |
| 0x00179368 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x00179483 | `5RunTestsTask` | Known | RTOS task thread |
| 0x001D55B4 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x001E5AA8 | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x001E5AC0 | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x001E6F08 | `FWInterruptHandlerTask` | Known | RTOS task thread |
| 0x0032CBE4 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0032CBF8 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0032CC0C | `SBP2CommandTask` | Known | RTOS task thread |

---

## 3. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017BF6A | `.net codec i dette produkt bruges i henhold til en licensaftale fra VoiceAge Cor` | Known | Audio system |
| 0x0017FBD5 | `.net Codec in diesem Produkt wird mit der Lizenz der VoiceAge Corporation verwen` | Known | Audio system |
| 0x00186AC2 | `.net codec t` | Known | Audio system |
| 0x0018ADDC | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x0018E55A | ` utilizzato su licenza da VoiceAge Corporation. Il codec ACELP` | Known | Audio system |
| 0x00196182 | `.net codec` | Known | Audio system |
| 0x00199F14 | `.net-codec in dit product wordt gebruikt in licentie van VoiceAge Corporation. G` | Known | Audio system |
| 0x001B88B9 | `.net codec in this product is used under license from VoiceAge Corporation. Port` | Known | Audio system |
| 0x0032C314 | `AudioCodecs` | Known | Audio system |
| 0x0032E3F7 | `msCodeCom` | Known | Audio system |

---

## 4. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017BE18 | `Audible-softwaren i dette produkt bruges i henhold til en licensaftale fra Audib` | Known | Audible audiobook format |
| 0x0017BE78 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x0017FA8C | `Die Audible Software in diesem Produkt wird in Lizenz von Audible verwendet. Cop` | Known | Audible audiobook format |
| 0x0017FAE5 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x0018329C | `El software Audible incluido en este producto se usa bajo licencia de Audible. C` | Known | Audible audiobook format |
| 0x001832F7 | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x001869B6 | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright ` | Known | Audible audiobook format |
| 0x001869F0 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x0018ACCC | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x0018AD16 | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x0018AD2B | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x0018E490 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x0018E4B9 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0018E4E8 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x00192295 | ` Audible ` | Known | Audible audiobook format |
| 0x001922B6 | `Audible ` | Known | Audible audiobook format |
| 0x0019230F | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x00196033 | ` Audible ` | Known | Audible audiobook format |
| 0x0019604E | ` Audible` | Known | Audible audiobook format |
| 0x00196092 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x00199DCC | `De Audible-software in dit product wordt gebruikt in licentie van Audible. Copyr` | Known | Audible audiobook format |
| 0x00199E23 | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x0019D470 | `Audible-programvaren i dette produktet brukes under lisens fra Audible. Copyrigh` | Known | Audible audiobook format |
| 0x0019D4C4 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x001A0F8C | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x001A0FBB | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x001A0FD2 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x001A485F | ` Audible ` | Known | Audible audiobook format |
| 0x001A4871 | ` Audible ` | Known | Audible audiobook format |
| 0x001A4895 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x001A80AC | `Audible ` | Known | Audible audiobook format |
| 0x001A80C0 | ` Audible ` | Known | Audible audiobook format |
| 0x001A80EA | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x001B8780 | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x001B87D5 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x0032C23C | `Audible` | Known | Audible audiobook format |

---

## 5. Audio/Codec - Apple Lossless

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0032C2B0 | `AppleLossless` | Known | Apple Lossless codec |

---

## 6. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017C014 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunhofer IIS og THOMSON multi` | Known | MP3 codec |
| 0x0017FC9B | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und THOMSON Multimedia.` | Known | MP3 codec |
| 0x00183499 | `n de audio MPEG Layer-3 utilizada bajo licencia de Fraunhofer IIS y THOMSON mult` | Known | MP3 codec |
| 0x00186B58 | `MPEG Layer-3 -` | Known | MP3 codec |
| 0x00186B6A | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | MP3 codec |
| 0x0018AEE8 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | MP3 codec |
| 0x0018E630 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da Fraunhofer IIS e THOMSON` | Known | MP3 codec |
| 0x001924C4 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00192510 | `Fraunhofer IIS ` | Known | MP3 codec |
| 0x0019621C | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00196243 | ` Fraunhofer IIS` | Known | MP3 codec |
| 0x00199FB0 | `Technologie voor codering van MPEG Layer-3-audio in licentie van Fraunhofer IIS ` | Known | MP3 codec |
| 0x0019D640 | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fraunhofer IIS og THOMSON m` | Known | MP3 codec |
| 0x001A116C | `Ljudkodningstekniken MPEG Layer-3 ` | Known | MP3 codec |
| 0x001A11A2 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | MP3 codec |
| 0x001A49EC | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x001A4A0E | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x001A8248 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x001A826D | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x001B894C | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |
| 0x001E6A93 | `Boot time (MP3PlayerExampleApp constructor)` | Known | MP3 codec |

---

## 7. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0032C20C | `AppleDRMVersion` | Known | DRM system |
| 0x0032C244 | `AppleDRM` | Known | DRM system |

---

## 8. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174965 | `#!#iTunesDB` | Known | iTunes database |

---

## 9. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00173DDC | `FirewireInitiator` | Known | FireWire interface |
| 0x00173DFC | `FirewireHandler` | Known | FireWire interface |
| 0x001741D4 | `FirewireGuid` | Known | FireWire interface |
| 0x0017B762 | ` den kan bruges som FireWire-disk, og tr` | Known | FireWire interface |
| 0x0017CA9C | `FireWire tilsluttet` | Known | FireWire interface |
| 0x0017F2C7 | `ffnen Sie das Adressbuch, Microsoft Entourage oder Palm Desktop und exportieren ` | Known | FireWire interface |
| 0x001807AE | `ber FireWire verbunden` | Known | FireWire interface |
| 0x00183FF4 | `FireWire conectado` | Known | FireWire interface |
| 0x00186306 | ` FireWire-levyn` | Known | FireWire interface |
| 0x00187694 | `FireWire liitetty` | Known | FireWire interface |
| 0x0018A4A9 | `utiliser comme disque FireWire. Puis glissez les vCards dans le dossier Contacts` | Known | FireWire interface |
| 0x0018BA54 | `FireWire Connect` | Known | FireWire interface |
| 0x0018F0FC | `FireWire Connesso` | Known | FireWire interface |
| 0x001932E0 | `FireWire ` | Known | FireWire interface |
| 0x00196DB4 | `FireWire ` | Known | FireWire interface |
| 0x00199498 | `Op de iPod kunt u adres- en agendagegevens opslaan. Als u met iSync werkt (Mac O` | Known | FireWire interface |
| 0x0019AAA8 | `FireWire aangesloten` | Known | FireWire interface |
| 0x0019E09C | `Koblet til via FireWire` | Known | FireWire interface |
| 0x001A086C | `ll sedan in din iPod som FireWire-h` | Known | FireWire interface |
| 0x001A09F2 | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x001A1C78 | `FireWire anslutet` | Known | FireWire interface |
| 0x001A5454 | `FireWire ` | Known | FireWire interface |
| 0x001A8CEC | `FireWire ` | Known | FireWire interface |
| 0x001B93F4 | `FireWire Connected` | Known | FireWire interface |
| 0x0032C364 | `FireWire` | Known | FireWire interface |
| 0x0032C3D4 | `FireWireVersion` | Known | FireWire interface |
| 0x003359D9 | `FireWire` | Known | FireWire interface |
| 0x00335D2F | `<key>FireWireGUID</key>` | Known | FireWire interface |

---

## 10. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174114 | `diskModeImageRev` | Known | Hardware interface |
| 0x001B7EC0 | `Disk Mode` | Known | Hardware interface |

---

## 11. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00173E0C | `PCFPowerMgr` | Known | Power management |
| 0x00173ECC | `USBPowerSense` | Known | Power management |
| 0x001744D8 | `PowerManager` | Known | Power management |
| 0x001B7A74 | `Charging` | Known | Power management |
| 0x001B9440 | `Low Battery` | Known | Power management |
| 0x0032C380 | `PowerInformation` | Known | Power management |

---

## 12. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017B9D8 | `Alarmer` | Known | UI element |
| 0x0017C9EC | `Alarmer` | Known | UI element |
| 0x0017FA5B | `hlen" den Alarm beenden` | Known | UI element |
| 0x001821D4 | `Calendario` | Known | UI element |
| 0x001821E0 | `Calendarios` | Known | UI element |
| 0x00182D9C | `Calendario` | Known | UI element |
| 0x00182DA8 | `Calendarios` | Known | UI element |
| 0x00182DC4 | `Alarmas` | Known | UI element |
| 0x00183280 | `Alarma` | Known | UI element |
| 0x001837EC | `Alarma` | Known | UI element |
| 0x0018398A | `Calendario` | Known | UI element |
| 0x00183BD4 | `Alarma` | Known | UI element |
| 0x00183E34 | `Calendarios` | Known | UI element |
| 0x00183F04 | `Alarma` | Known | UI element |
| 0x00183F4C | `Alarmas` | Known | UI element |
| 0x0018A7D8 | `Alarmes` | Known | UI element |
| 0x0018ACA8 | `Alarme` | Known | UI element |
| 0x0018B1F0 | `Alarme` | Known | UI element |
| 0x0018B264 | `Alarme` | Known | UI element |
| 0x0018B5F4 | `Alarme` | Known | UI element |
| 0x0018B934 | `Alarme` | Known | UI element |
| 0x0018B994 | `Alarmes` | Known | UI element |
| 0x0018D498 | `Calendario` | Known | UI element |
| 0x0018D4A4 | `Calendari` | Known | UI element |
| 0x0018DFB4 | `Calendario` | Known | UI element |
| 0x0018DFC0 | `Calendari` | Known | UI element |
| 0x0018EAA2 | `Calendario` | Known | UI element |
| 0x0018EF30 | `Calendari` | Known | UI element |
| 0x0019D018 | `Alarmer` | Known | UI element |
| 0x0019DCC4 | `Alarmtidspunkt` | Known | UI element |
| 0x0019DFE8 | `Alarmer` | Known | UI element |
| 0x001A1874 | `Alarmtid` | Known | UI element |
| 0x001B7488 | `Calendar` | Known | UI element |
| 0x001B7494 | `Calendars` | Known | UI element |
| 0x001B8310 | `Calendar` | Known | UI element |
| 0x001B831C | `Calendars` | Known | UI element |
| 0x001B8334 | `Alarms` | Known | UI element |
| 0x001B8C88 | `Alarm Clock` | Known | UI element |
| 0x001B8DAE | `Calendar` | Known | UI element |
| 0x001B8FDC | `Alarm Time` | Known | UI element |
| 0x001B8FE8 | `Alarm Clock` | Known | UI element |
| 0x001B9204 | `Calendars` | Known | UI element |
| 0x001B92D0 | `Alarm Clock` | Known | UI element |
| 0x001B9338 | `Alarms` | Known | UI element |
| 0x001E5647 | `Calendars` | Known | UI element |
| 0x001E5651 | `Calendars\` | Known | UI element |

---

## 13. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017AF34 | `Podcasts` | Known | Menu item |
| 0x0017C4EA | `Podcasts` | Known | Menu item |
| 0x0017C644 | `Podcasts` | Known | Menu item |
| 0x0017CA18 | `Podcasts` | Known | Menu item |
| 0x0017EA20 | `Extras` | Known | Menu item |
| 0x0017EA78 | `Podcasts` | Known | Menu item |
| 0x00180174 | `Extras` | Known | Menu item |
| 0x001801AE | `Podcasts` | Known | Menu item |
| 0x00180314 | `Podcasts` | Known | Menu item |
| 0x00180728 | `Podcasts` | Known | Menu item |
| 0x00180814 | `Extras` | Known | Menu item |
| 0x00182200 | `Extras` | Known | Menu item |
| 0x0018225C | `Podcasts` | Known | Menu item |
| 0x001839AC | `Extras` | Known | Menu item |
| 0x001839DE | `Podcasts` | Known | Menu item |
| 0x00183B60 | `Podcasts` | Known | Menu item |
| 0x00183F74 | `Podcasts` | Known | Menu item |
| 0x00184058 | `Extras` | Known | Menu item |
| 0x00189AE4 | `Albums` | Known | Menu item |
| 0x00189AEC | `Genres` | Known | Menu item |
| 0x00189B7C | `Extras` | Known | Menu item |
| 0x00189BD0 | `Podcasts` | Known | Menu item |
| 0x00189C0C | `Albums` | Known | Menu item |
| 0x0018B3C4 | `Extras` | Known | Menu item |
| 0x0018B3EE | `Genres` | Known | Menu item |
| 0x0018B3FA | `Podcasts` | Known | Menu item |
| 0x0018B412 | `Albums` | Known | Menu item |
| 0x0018B57C | `Genres` | Known | Menu item |
| 0x0018B584 | `Podcasts` | Known | Menu item |
| 0x0018B59C | `Albums` | Known | Menu item |
| 0x0018B744 | `Photos` | Known | Menu item |
| 0x0018B9A8 | `Genres` | Known | Menu item |
| 0x0018B9BC | `Podcasts` | Known | Menu item |
| 0x0018B9C8 | `Albums` | Known | Menu item |
| 0x0018BAD8 | `Extras` | Known | Menu item |
| 0x00198D04 | `Albums` | Known | Menu item |
| 0x00198D0C | `Genres` | Known | Menu item |
| 0x00198DF4 | `Podcasts` | Known | Menu item |
| 0x00198E2C | `Albums` | Known | Menu item |
| 0x0019A4DE | `Genres` | Known | Menu item |
| 0x0019A4EA | `Podcasts` | Known | Menu item |
| 0x0019A502 | `Albums` | Known | Menu item |
| 0x0019A658 | `Genres` | Known | Menu item |
| 0x0019A660 | `Podcasts` | Known | Menu item |
| 0x0019A674 | `Albums` | Known | Menu item |
| 0x0019AA14 | `Genres` | Known | Menu item |
| 0x0019AA28 | `Podcasts` | Known | Menu item |
| 0x0019AA34 | `Albums` | Known | Menu item |
| 0x001B7410 | `Now Playing` | Known | Menu item |
| 0x001B741C | `Artists` | Known | Menu item |
| 0x001B7434 | `Albums` | Known | Menu item |
| 0x001B743C | `Genres` | Known | Menu item |
| 0x001B7444 | `Composers` | Known | Menu item |
| 0x001B74BC | `Extras` | Known | Menu item |
| 0x001B74C4 | `Playlists` | Known | Menu item |
| 0x001B74D0 | `Audiobooks` | Known | Menu item |
| 0x001B7500 | `Shuffle Songs` | Known | Menu item |
| 0x001B7510 | `Podcasts` | Known | Menu item |
| 0x001B7548 | `Albums` | Known | Menu item |
| 0x001B7E8C | `Now Playing` | Known | Menu item |
| 0x001B8D6C | `Shuffle Songs` | Known | Menu item |
| 0x001B8DCC | `Extras` | Known | Menu item |
| 0x001B8DD6 | `Audiobooks` | Known | Menu item |
| 0x001B8DE6 | `Composers` | Known | Menu item |
| 0x001B8DF2 | `Genres` | Known | Menu item |
| 0x001B8DFE | `Podcasts` | Known | Menu item |
| 0x001B8E12 | `Albums` | Known | Menu item |
| 0x001B8E1E | `Artists` | Known | Menu item |
| 0x001B8E2A | `Playlists` | Known | Menu item |
| 0x001B8F18 | `Main Menu` | Known | Menu item |
| 0x001B8F70 | `Audiobooks` | Known | Menu item |
| 0x001B8F7C | `Composers` | Known | Menu item |
| 0x001B8F88 | `Genres` | Known | Menu item |
| 0x001B8F90 | `Podcasts` | Known | Menu item |
| 0x001B8FA4 | `Albums` | Known | Menu item |
| 0x001B8FAC | `Artists` | Known | Menu item |
| 0x001B8FB4 | `Playlists` | Known | Menu item |
| 0x001B8FC0 | `Settings` | Known | Menu item |
| 0x001B90DC | `Photos` | Known | Menu item |
| 0x001B92F4 | `Audiobooks` | Known | Menu item |
| 0x001B9340 | `Settings` | Known | Menu item |
| 0x001B934C | `Genres` | Known | Menu item |
| 0x001B9354 | `Artists` | Known | Menu item |
| 0x001B935C | `Podcasts` | Known | Menu item |
| 0x001B9368 | `Albums` | Known | Menu item |
| 0x001B9370 | `Composers` | Known | Menu item |
| 0x001B938C | `Audiobooks` | Known | Menu item |
| 0x001B93C8 | `Playlists` | Known | Menu item |
| 0x001B9454 | `Extras` | Known | Menu item |
| 0x001B94A4 | `Main Menu` | Known | Menu item |

---

## 14. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174178 | `iPod_Control\Device` | Filesystem Path | |
| 0x0017418C | `iPod_Control` | Filesystem Path | |
| 0x0017419C | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x00174904 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x00174974 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x0017499C | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x001749CC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001749E4 | `System_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x00174A08 | `iPod_Control\Device` | Filesystem Path | |
| 0x00174A28 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x00174A70 | `iPod_Control\Music\` | Filesystem Path | |
| 0x001D5564 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001E5675 | `iPod_Control\iTunes\` | Filesystem Path | |

---

## 15. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017B2A0 | `Acoustic` | EQ Preset | |
| 0x0017B2AC | `Bass Booster` | EQ Preset | |
| 0x0017B2CC | `Classical` | EQ Preset | |
| 0x0017B2E8 | `Electronic` | EQ Preset | |
| 0x0017B2FC | `Hip Hop` | EQ Preset | |
| 0x0017B314 | `Loudness` | EQ Preset | |
| 0x0017B320 | `Lounge` | EQ Preset | |
| 0x0017B344 | `Small Speakers` | EQ Preset | |
| 0x0017B354 | `Spoken Word` | EQ Preset | |
| 0x0017B360 | `Treble Booster` | EQ Preset | |
| 0x0017B380 | `Vocal Booster` | EQ Preset | |
| 0x0017EE34 | `Acoustic` | EQ Preset | |
| 0x0017EE74 | `Electronic` | EQ Preset | |
| 0x0017EE88 | `Hip Hop` | EQ Preset | |
| 0x0017EEA0 | `Loudness` | EQ Preset | |
| 0x00182648 | `Hip Hop` | EQ Preset | |
| 0x00182658 | `Latina` | EQ Preset | |
| 0x00182660 | `Loudness` | EQ Preset | |
| 0x0018266C | `Lounge` | EQ Preset | |
| 0x00185EFC | `Lounge` | EQ Preset | |
| 0x00189FE4 | `Hip Hop` | EQ Preset | |
| 0x00189FFC | `Loudness` | EQ Preset | |
| 0x0018D904 | `Hip Hop` | EQ Preset | |
| 0x0018D914 | `Latina` | EQ Preset | |
| 0x0018D91C | `Loudness` | EQ Preset | |
| 0x0018D928 | `Lounge` | EQ Preset | |
| 0x00191224 | `Acoustic` | EQ Preset | |
| 0x00191230 | `Bass Booster` | EQ Preset | |
| 0x00191250 | `Classical` | EQ Preset | |
| 0x0019126C | `Electronic` | EQ Preset | |
| 0x00191280 | `Hip Hop` | EQ Preset | |
| 0x00191298 | `Loudness` | EQ Preset | |
| 0x001912A4 | `Lounge` | EQ Preset | |
| 0x001912C8 | `Small Speakers` | EQ Preset | |
| 0x001912D8 | `Spoken Word` | EQ Preset | |
| 0x001912E4 | `Treble Booster` | EQ Preset | |
| 0x00191304 | `Vocal Booster` | EQ Preset | |
| 0x001952B8 | `Acoustic` | EQ Preset | |
| 0x001952C4 | `Bass Booster` | EQ Preset | |
| 0x001952E4 | `Classical` | EQ Preset | |
| 0x00195300 | `Electronic` | EQ Preset | |
| 0x00195314 | `Hip Hop` | EQ Preset | |
| 0x0019532C | `Loudness` | EQ Preset | |
| 0x00195338 | `Lounge` | EQ Preset | |
| 0x0019535C | `Small Speakers` | EQ Preset | |
| 0x0019536C | `Spoken Word` | EQ Preset | |
| 0x00195378 | `Treble Booster` | EQ Preset | |
| 0x00195398 | `Vocal Booster` | EQ Preset | |
| 0x001991D8 | `Loudness` | EQ Preset | |
| 0x001991E4 | `Lounge` | EQ Preset | |
| 0x0019C910 | `Hip Hop` | EQ Preset | |
| 0x0019C920 | `Latino` | EQ Preset | |
| 0x0019C928 | `Loudness` | EQ Preset | |
| 0x0019C934 | `Lounge` | EQ Preset | |
| 0x001A03DC | `Acoustic` | EQ Preset | |
| 0x001A03E8 | `Bass Booster` | EQ Preset | |
| 0x001A0408 | `Classical` | EQ Preset | |
| 0x001A0424 | `Electronic` | EQ Preset | |
| 0x001A0438 | `Hip Hop` | EQ Preset | |
| 0x001A0450 | `Loudness` | EQ Preset | |
| 0x001A045C | `Lounge` | EQ Preset | |
| 0x001A0480 | `Small Speakers` | EQ Preset | |
| 0x001A0490 | `Spoken Word` | EQ Preset | |
| 0x001A049C | `Treble Booster` | EQ Preset | |
| 0x001A04BC | `Vocal Booster` | EQ Preset | |
| 0x001A3C64 | `Acoustic` | EQ Preset | |
| 0x001A3C70 | `Bass Booster` | EQ Preset | |
| 0x001A3C90 | `Classical` | EQ Preset | |
| 0x001A3CAC | `Electronic` | EQ Preset | |
| 0x001A3CC0 | `Hip Hop` | EQ Preset | |
| 0x001A3CD8 | `Loudness` | EQ Preset | |
| 0x001A3CE4 | `Lounge` | EQ Preset | |
| 0x001A3D08 | `Small Speakers` | EQ Preset | |
| 0x001A3D18 | `Spoken Word` | EQ Preset | |
| 0x001A3D24 | `Treble Booster` | EQ Preset | |
| 0x001A3D44 | `Vocal Booster` | EQ Preset | |
| 0x001A743C | `Acoustic` | EQ Preset | |
| 0x001A7448 | `Bass Booster` | EQ Preset | |
| 0x001A7468 | `Classical` | EQ Preset | |
| 0x001A7484 | `Electronic` | EQ Preset | |
| 0x001A7498 | `Hip Hop` | EQ Preset | |
| 0x001A74B0 | `Loudness` | EQ Preset | |
| 0x001A74BC | `Lounge` | EQ Preset | |
| 0x001A74E0 | `Small Speakers` | EQ Preset | |
| 0x001A74F0 | `Spoken Word` | EQ Preset | |
| 0x001A74FC | `Treble Booster` | EQ Preset | |
| 0x001A751C | `Vocal Booster` | EQ Preset | |
| 0x001B78F0 | `Acoustic` | EQ Preset | |
| 0x001B78FC | `Bass Booster` | EQ Preset | |
| 0x001B791C | `Classical` | EQ Preset | |
| 0x001B7938 | `Electronic` | EQ Preset | |
| 0x001B794C | `Hip Hop` | EQ Preset | |
| 0x001B7964 | `Loudness` | EQ Preset | |
| 0x001B7970 | `Lounge` | EQ Preset | |
| 0x001B7994 | `Small Speakers` | EQ Preset | |
| 0x001B79A4 | `Spoken Word` | EQ Preset | |
| 0x001B79B0 | `Treble Booster` | EQ Preset | |
| 0x001B79D0 | `Vocal Booster` | EQ Preset | |

---

## 16. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00183184 | `Error durante la importaci` | Diagnostic | |
| 0x0018EFC4 | `Errore` | Diagnostic | |
| 0x0018F1B4 | `Errore` | Diagnostic | |
| 0x001E6830 | `%s Error in file %s.` | Diagnostic | |
| 0x001E7317 | `Error loading operating system. Setup cannot continue.` | Diagnostic | |
| 0x00330608 | `setct-ErrorTBS` | Diagnostic | |

---

## 17. Assertions

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00017740 | `*** assertion failed: %s, file %s, line %d` | Assertion | |
| 0x0032C7B4 | `%s(%d): OpenSSL internal error, assertion failed: %s` | Assertion | |

---
