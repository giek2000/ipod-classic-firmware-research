# iPod Mini 1st Generation - RetailOS 1.4.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.4.1 |
| **IPSW** | iPod_6.1.4.1.ipsw |
| **Device** | iPod Mini 1st Generation (2004, Click Wheel, Anodized Aluminum) |
| **UpdaterFamilyID** | 6 |
| **Binary Size** | 4,506,624 bytes (4.30 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 4,506,624 bytes |
| **Total Strings (>=6)** | 9,944 |
| **Function Prologues** | 10,416 (ARM: 7,563, Thumb: 2,853) |
| **SoC** | PortalPlayer PP5020 |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Encrypted** | No |
| **SHA-256** | `a69031d594a0b54649c0a6cc087241808463b9d94a9e45793cedb7f02abd357f` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168950 | `MP3ExampleTask` | Hidden | Hidden Test |

---

## 2. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001684D4 | `WatchdogTask` | Known | RTOS task thread |
| 0x001684E4 | `AlarmTask` | Known | RTOS task thread |
| 0x001684FC | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x00168510 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x00168520 | `TopPlugTask` | Known | RTOS task thread |
| 0x0016852C | `HoldSwitchTask` | Known | RTOS task thread |
| 0x0016853C | `PlayBtnTask` | Known | RTOS task thread |
| 0x00168548 | `PrvBtnTask` | Known | RTOS task thread |
| 0x00168554 | `NextBtnTask` | Known | RTOS task thread |
| 0x00168560 | `ActionBtnTask` | Known | RTOS task thread |
| 0x00168570 | `MenuBtnTask` | Known | RTOS task thread |
| 0x0016857C | `DiskMgrTask` | Known | RTOS task thread |
| 0x00168598 | `CNATask` | Known | RTOS task thread |
| 0x001685A0 | `BacklightTask` | Known | RTOS task thread |
| 0x001685B0 | `SerialOptoTask` | Known | RTOS task thread |
| 0x001685C0 | `OptoTask` | Known | RTOS task thread |
| 0x001685CC | `FirewireTask` | Known | RTOS task thread |
| 0x00168914 | `HostOSTask` | Known | RTOS task thread |
| 0x0016DAAB | `5RunTestsTask` | Known | RTOS task thread |
| 0x001BD794 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x001CEDB8 | `FWInterruptHandlerTask` | Known | RTOS task thread |
| 0x00314A94 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x00314AA8 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x00314ABC | `SBP2CommandTask` | Known | RTOS task thread |

---

## 3. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016FF82 | `.net codec i dette produkt bruges i henhold til en licensaftale fra VoiceAge Cor` | Known | Audio system |
| 0x001733F9 | `.net Codec in diesem Produkt wird mit der Lizenz der VoiceAge Corporation verwen` | Known | Audio system |
| 0x0017917E | `.net codec t` | Known | Audio system |
| 0x0017CBF0 | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x0017FA9E | ` utilizzato su licenza da VoiceAge Corporation. Il codec ACELP` | Known | Audio system |
| 0x001863B6 | `.net codec` | Known | Audio system |
| 0x0018983C | `.net-codec in dit product wordt gebruikt in licentie van VoiceAge Corporation. G` | Known | Audio system |
| 0x001A4989 | `.net codec in this product is used under license from VoiceAge Corporation. Port` | Known | Audio system |
| 0x003141C4 | `AudioCodecs` | Known | Audio system |
| 0x003162A7 | `msCodeCom` | Known | Audio system |

---

## 4. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016FE30 | `Audible-softwaren i dette produkt bruges i henhold til en licensaftale fra Audib` | Known | Audible audiobook format |
| 0x0016FE90 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x001732B0 | `Die Audible Software in diesem Produkt wird in Lizenz von Audible verwendet. Cop` | Known | Audible audiobook format |
| 0x00173309 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x001761F4 | `El software Audible incluido en este producto se usa bajo licencia de Audible. C` | Known | Audible audiobook format |
| 0x0017624F | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x00179072 | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright ` | Known | Audible audiobook format |
| 0x001790AC | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x0017CAE0 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x0017CB2A | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x0017CB3F | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x0017F9D4 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x0017F9FD | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0017FA2C | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x00182E85 | ` Audible ` | Known | Audible audiobook format |
| 0x00182EA6 | `Audible ` | Known | Audible audiobook format |
| 0x00182EFF | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x00186267 | ` Audible ` | Known | Audible audiobook format |
| 0x00186282 | ` Audible` | Known | Audible audiobook format |
| 0x001862C6 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x001896F4 | `De Audible-software in dit product wordt gebruikt in licentie van Audible. Copyr` | Known | Audible audiobook format |
| 0x0018974B | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x0018C558 | `Audible-programvaren i dette produktet brukes under lisens fra Audible. Copyrigh` | Known | Audible audiobook format |
| 0x0018C5AC | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0018F880 | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x0018F8AF | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x0018F8C6 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x0019289B | ` Audible ` | Known | Audible audiobook format |
| 0x001928AD | ` Audible ` | Known | Audible audiobook format |
| 0x001928D1 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x00195864 | `Audible ` | Known | Audible audiobook format |
| 0x00195878 | ` Audible ` | Known | Audible audiobook format |
| 0x001958A2 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x001A4850 | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x001A48A5 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x003140EC | `Audible` | Known | Audible audiobook format |

---

## 5. Audio/Codec - Apple Lossless

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00314160 | `AppleLossless` | Known | Apple Lossless codec |

---

## 6. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017002C | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunhofer IIS og THOMSON multi` | Known | MP3 codec |
| 0x001734BF | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und THOMSON Multimedia.` | Known | MP3 codec |
| 0x001763F1 | `n de audio MPEG Layer-3 utilizada bajo licencia de Fraunhofer IIS y THOMSON mult` | Known | MP3 codec |
| 0x00179214 | `MPEG Layer-3 -` | Known | MP3 codec |
| 0x00179226 | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | MP3 codec |
| 0x0017CCFC | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | MP3 codec |
| 0x0017FB74 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da Fraunhofer IIS e THOMSON` | Known | MP3 codec |
| 0x001830B4 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00183100 | `Fraunhofer IIS ` | Known | MP3 codec |
| 0x00186450 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00186477 | ` Fraunhofer IIS` | Known | MP3 codec |
| 0x001898D8 | `Technologie voor codering van MPEG Layer-3-audio in licentie van Fraunhofer IIS ` | Known | MP3 codec |
| 0x0018C728 | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fraunhofer IIS og THOMSON m` | Known | MP3 codec |
| 0x0018FA60 | `Ljudkodningstekniken MPEG Layer-3 ` | Known | MP3 codec |
| 0x0018FA96 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | MP3 codec |
| 0x00192A28 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00192A4A | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x00195A00 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00195A25 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x001A4A1C | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |
| 0x001CE943 | `Boot time (MP3PlayerExampleApp constructor)` | Known | MP3 codec |

---

## 7. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x003140BC | `AppleDRMVersion` | Known | DRM system |
| 0x003140F4 | `AppleDRM` | Known | DRM system |

---

## 8. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00169021 | `#!#iTunesDB` | Known | iTunes database |

---

## 9. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001684B8 | `FirewireHandler` | Known | FireWire interface |
| 0x00168890 | `FirewireGuid` | Known | FireWire interface |
| 0x001708DC | `FireWire tilsluttet` | Known | FireWire interface |
| 0x00172D5F | `ffnen Sie das Adressbuch, Microsoft Entourage oder Palm Desktop und exportieren ` | Known | FireWire interface |
| 0x00173D5A | `ber FireWire verbunden` | Known | FireWire interface |
| 0x00176CC8 | `FireWire conectado` | Known | FireWire interface |
| 0x00179B18 | `FireWire liitetty` | Known | FireWire interface |
| 0x0017C585 | `utiliser comme disque FireWire. Puis glissez les vCards dans le dossier Contacts` | Known | FireWire interface |
| 0x0017D5DC | `FireWire Connect` | Known | FireWire interface |
| 0x001803C0 | `FireWire Connesso` | Known | FireWire interface |
| 0x00183BC0 | `FireWire ` | Known | FireWire interface |
| 0x00186D4C | `FireWire ` | Known | FireWire interface |
| 0x00189064 | `Op de iPod kunt u adres- en agendagegevens opslaan. Als u met iSync werkt (Mac O` | Known | FireWire interface |
| 0x0018A1A4 | `FireWire aangesloten` | Known | FireWire interface |
| 0x0018CF90 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x0018F53E | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x00190308 | `FireWire anslutet` | Known | FireWire interface |
| 0x00193268 | `FireWire ` | Known | FireWire interface |
| 0x0019627C | `FireWire ` | Known | FireWire interface |
| 0x001A52BC | `FireWire Connected` | Known | FireWire interface |
| 0x00314214 | `FireWire` | Known | FireWire interface |
| 0x00314284 | `FireWireVersion` | Known | FireWire interface |
| 0x0031D7AD | `FireWire` | Known | FireWire interface |
| 0x0031DB03 | `<key>FireWireGUID</key>` | Known | FireWire interface |

---

## 10. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001687D0 | `diskModeImageRev` | Known | Hardware interface |
| 0x001A41D4 | `Disk Mode` | Known | Hardware interface |

---

## 11. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001684C8 | `PCFPowerMgr` | Known | Power management |
| 0x00168588 | `USBPowerSense` | Known | Power management |
| 0x00168B94 | `PowerManager` | Known | Power management |
| 0x001A3D88 | `Charging` | Known | Power management |
| 0x001A5308 | `Low Battery` | Known | Power management |
| 0x00314230 | `PowerInformation` | Known | Power management |

---

## 12. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016FC48 | `Alarmer` | Known | UI element |
| 0x0017082C | `Alarmer` | Known | UI element |
| 0x0017327F | `hlen" den Alarm beenden` | Known | UI element |
| 0x001753D0 | `Calendario` | Known | UI element |
| 0x001753DC | `Calendarios` | Known | UI element |
| 0x00175F98 | `Calendario` | Known | UI element |
| 0x00175FA4 | `Calendarios` | Known | UI element |
| 0x00175FC0 | `Alarmas` | Known | UI element |
| 0x001761D8 | `Alarma` | Known | UI element |
| 0x00176744 | `Alarma` | Known | UI element |
| 0x00176896 | `Calendario` | Known | UI element |
| 0x00176AE0 | `Alarma` | Known | UI element |
| 0x00176B08 | `Calendarios` | Known | UI element |
| 0x00176BD8 | `Alarma` | Known | UI element |
| 0x00176C20 | `Alarmas` | Known | UI element |
| 0x0017C8B4 | `Alarmes` | Known | UI element |
| 0x0017CABC | `Alarme` | Known | UI element |
| 0x0017D004 | `Alarme` | Known | UI element |
| 0x0017D078 | `Alarme` | Known | UI element |
| 0x0017D3C0 | `Alarme` | Known | UI element |
| 0x0017D4BC | `Alarme` | Known | UI element |
| 0x0017D51C | `Alarmes` | Known | UI element |
| 0x0017EC70 | `Calendario` | Known | UI element |
| 0x0017EC7C | `Calendari` | Known | UI element |
| 0x0017F78C | `Calendario` | Known | UI element |
| 0x0017F798 | `Calendari` | Known | UI element |
| 0x0017FF9E | `Calendario` | Known | UI element |
| 0x001801F4 | `Calendari` | Known | UI element |
| 0x0018C350 | `Alarmer` | Known | UI element |
| 0x0018CD98 | `Alarmtidspunkt` | Known | UI element |
| 0x0018CEDC | `Alarmer` | Known | UI element |
| 0x0019011C | `Alarmtid` | Known | UI element |
| 0x001A379C | `Calendar` | Known | UI element |
| 0x001A37A8 | `Calendars` | Known | UI element |
| 0x001A4624 | `Calendar` | Known | UI element |
| 0x001A4630 | `Calendars` | Known | UI element |
| 0x001A4648 | `Alarms` | Known | UI element |
| 0x001A4D58 | `Alarm Clock` | Known | UI element |
| 0x001A4E42 | `Calendar` | Known | UI element |
| 0x001A5070 | `Alarm Time` | Known | UI element |
| 0x001A507C | `Alarm Clock` | Known | UI element |
| 0x001A50CC | `Calendars` | Known | UI element |
| 0x001A5198 | `Alarm Clock` | Known | UI element |
| 0x001A5200 | `Alarms` | Known | UI element |
| 0x001CD5AF | `Calendars` | Known | UI element |
| 0x001CD5B9 | `Calendars\` | Known | UI element |

---

## 13. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016F1B4 | `Podcasts` | Known | Menu item |
| 0x00170512 | `Podcasts` | Known | Menu item |
| 0x0017066C | `Podcasts` | Known | Menu item |
| 0x00170858 | `Podcasts` | Known | Menu item |
| 0x001724B8 | `Extras` | Known | Menu item |
| 0x00172510 | `Podcasts` | Known | Menu item |
| 0x00173950 | `Extras` | Known | Menu item |
| 0x0017398A | `Podcasts` | Known | Menu item |
| 0x00173AF0 | `Podcasts` | Known | Menu item |
| 0x00173CD4 | `Podcasts` | Known | Menu item |
| 0x00173DC0 | `Extras` | Known | Menu item |
| 0x001753FC | `Extras` | Known | Menu item |
| 0x00175458 | `Podcasts` | Known | Menu item |
| 0x001768B8 | `Extras` | Known | Menu item |
| 0x001768EA | `Podcasts` | Known | Menu item |
| 0x00176A6C | `Podcasts` | Known | Menu item |
| 0x00176C48 | `Podcasts` | Known | Menu item |
| 0x00176D2C | `Extras` | Known | Menu item |
| 0x0017BBE0 | `Albums` | Known | Menu item |
| 0x0017BBE8 | `Genres` | Known | Menu item |
| 0x0017BC58 | `Extras` | Known | Menu item |
| 0x0017BCAC | `Podcasts` | Known | Menu item |
| 0x0017BCE8 | `Albums` | Known | Menu item |
| 0x0017D188 | `Extras` | Known | Menu item |
| 0x0017D1B2 | `Genres` | Known | Menu item |
| 0x0017D1BE | `Podcasts` | Known | Menu item |
| 0x0017D1D6 | `Albums` | Known | Menu item |
| 0x0017D348 | `Genres` | Known | Menu item |
| 0x0017D350 | `Podcasts` | Known | Menu item |
| 0x0017D368 | `Albums` | Known | Menu item |
| 0x0017D530 | `Genres` | Known | Menu item |
| 0x0017D544 | `Podcasts` | Known | Menu item |
| 0x0017D550 | `Albums` | Known | Menu item |
| 0x0017D660 | `Extras` | Known | Menu item |
| 0x0018890C | `Albums` | Known | Menu item |
| 0x00188914 | `Genres` | Known | Menu item |
| 0x001889D4 | `Podcasts` | Known | Menu item |
| 0x00188A0C | `Albums` | Known | Menu item |
| 0x00189DCA | `Genres` | Known | Menu item |
| 0x00189DD6 | `Podcasts` | Known | Menu item |
| 0x00189DEE | `Albums` | Known | Menu item |
| 0x00189F44 | `Genres` | Known | Menu item |
| 0x00189F4C | `Podcasts` | Known | Menu item |
| 0x00189F60 | `Albums` | Known | Menu item |
| 0x0018A110 | `Genres` | Known | Menu item |
| 0x0018A124 | `Podcasts` | Known | Menu item |
| 0x0018A130 | `Albums` | Known | Menu item |
| 0x001A3740 | `Now Playing` | Known | Menu item |
| 0x001A374C | `Artists` | Known | Menu item |
| 0x001A3764 | `Albums` | Known | Menu item |
| 0x001A376C | `Genres` | Known | Menu item |
| 0x001A3774 | `Composers` | Known | Menu item |
| 0x001A37D0 | `Extras` | Known | Menu item |
| 0x001A37D8 | `Playlists` | Known | Menu item |
| 0x001A37E4 | `Audiobooks` | Known | Menu item |
| 0x001A3814 | `Shuffle Songs` | Known | Menu item |
| 0x001A3824 | `Podcasts` | Known | Menu item |
| 0x001A385C | `Albums` | Known | Menu item |
| 0x001A41A0 | `Now Playing` | Known | Menu item |
| 0x001A4E20 | `Shuffle Songs` | Known | Menu item |
| 0x001A4E60 | `Extras` | Known | Menu item |
| 0x001A4E6A | `Audiobooks` | Known | Menu item |
| 0x001A4E7A | `Composers` | Known | Menu item |
| 0x001A4E86 | `Genres` | Known | Menu item |
| 0x001A4E92 | `Podcasts` | Known | Menu item |
| 0x001A4EA6 | `Albums` | Known | Menu item |
| 0x001A4EB2 | `Artists` | Known | Menu item |
| 0x001A4EBE | `Playlists` | Known | Menu item |
| 0x001A4FAC | `Main Menu` | Known | Menu item |
| 0x001A5004 | `Audiobooks` | Known | Menu item |
| 0x001A5010 | `Composers` | Known | Menu item |
| 0x001A501C | `Genres` | Known | Menu item |
| 0x001A5024 | `Podcasts` | Known | Menu item |
| 0x001A5038 | `Albums` | Known | Menu item |
| 0x001A5040 | `Artists` | Known | Menu item |
| 0x001A5048 | `Playlists` | Known | Menu item |
| 0x001A5054 | `Settings` | Known | Menu item |
| 0x001A51BC | `Audiobooks` | Known | Menu item |
| 0x001A5208 | `Settings` | Known | Menu item |
| 0x001A5214 | `Genres` | Known | Menu item |
| 0x001A521C | `Artists` | Known | Menu item |
| 0x001A5224 | `Podcasts` | Known | Menu item |
| 0x001A5230 | `Albums` | Known | Menu item |
| 0x001A5238 | `Composers` | Known | Menu item |
| 0x001A5254 | `Audiobooks` | Known | Menu item |
| 0x001A5290 | `Playlists` | Known | Menu item |
| 0x001A531C | `Extras` | Known | Menu item |
| 0x001A536C | `Main Menu` | Known | Menu item |

---

## 14. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168834 | `iPod_Control\Device` | Filesystem Path | |
| 0x00168848 | `iPod_Control` | Filesystem Path | |
| 0x00168858 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x00168FC0 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x00169030 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x00169058 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x00169088 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001690A0 | `System_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x001690C4 | `iPod_Control\Device` | Filesystem Path | |
| 0x001690E4 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x0016912C | `iPod_Control\Music\` | Filesystem Path | |
| 0x001BD744 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001CD5DD | `iPod_Control\iTunes\` | Filesystem Path | |

---

## 15. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016F520 | `Acoustic` | EQ Preset | |
| 0x0016F52C | `Bass Booster` | EQ Preset | |
| 0x0016F54C | `Classical` | EQ Preset | |
| 0x0016F568 | `Electronic` | EQ Preset | |
| 0x0016F57C | `Hip Hop` | EQ Preset | |
| 0x0016F594 | `Loudness` | EQ Preset | |
| 0x0016F5A0 | `Lounge` | EQ Preset | |
| 0x0016F5C4 | `Small Speakers` | EQ Preset | |
| 0x0016F5D4 | `Spoken Word` | EQ Preset | |
| 0x0016F5E0 | `Treble Booster` | EQ Preset | |
| 0x0016F600 | `Vocal Booster` | EQ Preset | |
| 0x001728CC | `Acoustic` | EQ Preset | |
| 0x0017290C | `Electronic` | EQ Preset | |
| 0x00172920 | `Hip Hop` | EQ Preset | |
| 0x00172938 | `Loudness` | EQ Preset | |
| 0x00175844 | `Hip Hop` | EQ Preset | |
| 0x00175854 | `Latina` | EQ Preset | |
| 0x0017585C | `Loudness` | EQ Preset | |
| 0x00175868 | `Lounge` | EQ Preset | |
| 0x00178804 | `Hip Hop` | EQ Preset | |
| 0x00178828 | `Lounge` | EQ Preset | |
| 0x0017C0C0 | `Hip Hop` | EQ Preset | |
| 0x0017C0D8 | `Loudness` | EQ Preset | |
| 0x0017F0DC | `Hip Hop` | EQ Preset | |
| 0x0017F0EC | `Latina` | EQ Preset | |
| 0x0017F0F4 | `Loudness` | EQ Preset | |
| 0x0017F100 | `Lounge` | EQ Preset | |
| 0x00182130 | `Acoustic` | EQ Preset | |
| 0x0018213C | `Bass Booster` | EQ Preset | |
| 0x0018215C | `Classical` | EQ Preset | |
| 0x00182178 | `Electronic` | EQ Preset | |
| 0x0018218C | `Hip Hop` | EQ Preset | |
| 0x001821A4 | `Loudness` | EQ Preset | |
| 0x001821B0 | `Lounge` | EQ Preset | |
| 0x001821D4 | `Small Speakers` | EQ Preset | |
| 0x001821E4 | `Spoken Word` | EQ Preset | |
| 0x001821F0 | `Treble Booster` | EQ Preset | |
| 0x00182210 | `Vocal Booster` | EQ Preset | |
| 0x001857E0 | `Acoustic` | EQ Preset | |
| 0x001857EC | `Bass Booster` | EQ Preset | |
| 0x0018580C | `Classical` | EQ Preset | |
| 0x00185828 | `Electronic` | EQ Preset | |
| 0x0018583C | `Hip Hop` | EQ Preset | |
| 0x00185854 | `Loudness` | EQ Preset | |
| 0x00185860 | `Lounge` | EQ Preset | |
| 0x00185884 | `Small Speakers` | EQ Preset | |
| 0x00185894 | `Spoken Word` | EQ Preset | |
| 0x001858A0 | `Treble Booster` | EQ Preset | |
| 0x001858C0 | `Vocal Booster` | EQ Preset | |
| 0x00188DB8 | `Loudness` | EQ Preset | |
| 0x00188DC4 | `Lounge` | EQ Preset | |
| 0x0018BC48 | `Hip Hop` | EQ Preset | |
| 0x0018BC58 | `Latino` | EQ Preset | |
| 0x0018BC60 | `Loudness` | EQ Preset | |
| 0x0018BC6C | `Lounge` | EQ Preset | |
| 0x0018EF34 | `Acoustic` | EQ Preset | |
| 0x0018EF40 | `Bass Booster` | EQ Preset | |
| 0x0018EF60 | `Classical` | EQ Preset | |
| 0x0018EF7C | `Electronic` | EQ Preset | |
| 0x0018EF90 | `Hip Hop` | EQ Preset | |
| 0x0018EFA8 | `Loudness` | EQ Preset | |
| 0x0018EFB4 | `Lounge` | EQ Preset | |
| 0x0018EFD8 | `Small Speakers` | EQ Preset | |
| 0x0018EFE8 | `Spoken Word` | EQ Preset | |
| 0x0018EFF4 | `Treble Booster` | EQ Preset | |
| 0x0018F014 | `Vocal Booster` | EQ Preset | |
| 0x00191F2C | `Acoustic` | EQ Preset | |
| 0x00191F38 | `Bass Booster` | EQ Preset | |
| 0x00191F58 | `Classical` | EQ Preset | |
| 0x00191F74 | `Electronic` | EQ Preset | |
| 0x00191F88 | `Hip Hop` | EQ Preset | |
| 0x00191FA0 | `Loudness` | EQ Preset | |
| 0x00191FAC | `Lounge` | EQ Preset | |
| 0x00191FD0 | `Small Speakers` | EQ Preset | |
| 0x00191FE0 | `Spoken Word` | EQ Preset | |
| 0x00191FEC | `Treble Booster` | EQ Preset | |
| 0x0019200C | `Vocal Booster` | EQ Preset | |
| 0x00194E9C | `Acoustic` | EQ Preset | |
| 0x00194EA8 | `Bass Booster` | EQ Preset | |
| 0x00194EC8 | `Classical` | EQ Preset | |
| 0x00194EE4 | `Electronic` | EQ Preset | |
| 0x00194EF8 | `Hip Hop` | EQ Preset | |
| 0x00194F10 | `Loudness` | EQ Preset | |
| 0x00194F1C | `Lounge` | EQ Preset | |
| 0x00194F40 | `Small Speakers` | EQ Preset | |
| 0x00194F50 | `Spoken Word` | EQ Preset | |
| 0x00194F5C | `Treble Booster` | EQ Preset | |
| 0x00194F7C | `Vocal Booster` | EQ Preset | |
| 0x001A3C04 | `Acoustic` | EQ Preset | |
| 0x001A3C10 | `Bass Booster` | EQ Preset | |
| 0x001A3C30 | `Classical` | EQ Preset | |
| 0x001A3C4C | `Electronic` | EQ Preset | |
| 0x001A3C60 | `Hip Hop` | EQ Preset | |
| 0x001A3C78 | `Loudness` | EQ Preset | |
| 0x001A3C84 | `Lounge` | EQ Preset | |
| 0x001A3CA8 | `Small Speakers` | EQ Preset | |
| 0x001A3CB8 | `Spoken Word` | EQ Preset | |
| 0x001A3CC4 | `Treble Booster` | EQ Preset | |
| 0x001A3CE4 | `Vocal Booster` | EQ Preset | |

---

## 16. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00180288 | `Errore` | Diagnostic | |
| 0x00180478 | `Errore` | Diagnostic | |
| 0x001CE6E0 | `%s Error in file %s.` | Diagnostic | |
| 0x001CF1C7 | `Error loading operating system. Setup cannot continue.` | Diagnostic | |
| 0x003184B8 | `setct-ErrorTBS` | Diagnostic | |

---

## 17. Assertions

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00017740 | `*** assertion failed: %s, file %s, line %d` | Assertion | |
| 0x00314664 | `%s(%d): OpenSSL internal error, assertion failed: %s` | Assertion | |

---
