# iPod 3rd Generation - RetailOS 2.2.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 2.2.3 |
| **IPSW** | iPod_2.2.2.3.ipsw |
| **Device** | iPod 3rd Generation (2003, Touch Wheel + Touch Buttons) |
| **UpdaterFamilyID** | 2 |
| **Binary Size** | 4,561,920 bytes (4.35 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 4,561,920 bytes |
| **Total Strings (>=6)** | 9,755 |
| **Function Prologues** | 8,113 (ARM: 7,192, Thumb: 921) |
| **SoC** | PortalPlayer PP5002 |
| **Architecture** | ARM7TDMI (ARMv4T), dual-core |
| **Encrypted** | No |
| **SHA-256** | `d4a95ed35add9f001058ca66ce515e535833cfb2e9de2c6c9377ba9ef56ea0ce` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00151224 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x0035F114 | `BTM Debug Zones %s (0x%08X)` | Hidden | Debug/Diagnostic |
| 0x00362884 | `Retail mode` | Hidden | Demo/Retail Mode |
| 0x00362894 | `Debug mode` | Hidden | Debug/Diagnostic |

---

## 2. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00151210 | `HostOSTask` | Known | RTOS task thread |
| 0x001514B8 | `USBI2CSlaveTask` | Known | RTOS task thread |
| 0x00151570 | `WatchdogTask` | Known | RTOS task thread |
| 0x00151580 | `AlarmTask` | Known | RTOS task thread |
| 0x00151598 | `BattPowerTask` | Known | RTOS task thread |
| 0x001515A8 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x001515B8 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x001515C8 | `PlayBtnTask` | Known | RTOS task thread |
| 0x001515D4 | `PrvBtnTask` | Known | RTOS task thread |
| 0x001515E0 | `NextBtnTask` | Known | RTOS task thread |
| 0x001515EC | `ActionBtnTask` | Known | RTOS task thread |
| 0x001515FC | `MenuBtnTask` | Known | RTOS task thread |
| 0x00151608 | `DiskMgrTask` | Known | RTOS task thread |
| 0x00151614 | `CNATask` | Known | RTOS task thread |
| 0x0015161C | `OptoTask` | Known | RTOS task thread |
| 0x00151628 | `FirewireTask` | Known | RTOS task thread |
| 0x00151638 | `RTXC v3.2fpp for ARM and Thumb - ARM ADS 1.0 Jul-08-00 Key: 24104` | Known | RTOS task thread |
| 0x001AF6F4 | `RunTestsTask` | Known | RTOS task thread |
| 0x001AFDB4 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x001C4588 | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x001C45A0 | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x001C4668 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x0044EB54 | `USBI2CSlaveTask` | Known | RTOS task thread |
| 0x0044EBAC | `CNATask` | Known | RTOS task thread |
| 0x0044EBB4 | `BattPowerTask` | Known | RTOS task thread |

---

## 3. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0015A07A | `.net codec i dette produkt bruges i henhold til en licensaftale fra VoiceAge Cor` | Known | Audio system |
| 0x0015DB95 | `.net Codec in diesem Produkt wird mit der Lizenz der VoiceAge Corporation verwen` | Known | Audio system |
| 0x00164726 | `.net codec t` | Known | Audio system |
| 0x0016886E | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x0016BE12 | ` utilizzato su licenza da VoiceAge Corporation. Il codec ACELP` | Known | Audio system |
| 0x0017353A | `.net codec` | Known | Audio system |
| 0x001770D0 | `.net-codec in dit product wordt gebruikt in licentie van VoiceAge Corporation. G` | Known | Audio system |
| 0x00194E55 | `.net codec in this product is used under license from VoiceAge Corporation. Port` | Known | Audio system |
| 0x00313C6B | `msCodeCom` | Known | Audio system |
| 0x0041F458 | `D1CODEC2.8` | Known | Audio system |
| 0x0041F47C | `D1 CODEC 2.8V` | Known | Audio system |

---

## 4. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00159F28 | `Audible-softwaren i dette produkt bruges i henhold til en licensaftale fra Audib` | Known | Audible audiobook format |
| 0x00159F88 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x0015DA4C | `Die Audible Software in diesem Produkt wird in Lizenz von Audible verwendet. Cop` | Known | Audible audiobook format |
| 0x0015DAA5 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x00161094 | `El software Audible incluido en este producto se usa bajo licencia de Audible. C` | Known | Audible audiobook format |
| 0x001610EF | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x0016461A | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright ` | Known | Audible audiobook format |
| 0x00164654 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x00168760 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x001687AA | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x001687BF | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x0016BD48 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x0016BD71 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0016BDA0 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x0016F8BD | ` Audible ` | Known | Audible audiobook format |
| 0x0016F8DE | `Audible ` | Known | Audible audiobook format |
| 0x0016F938 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x001733EB | ` Audible ` | Known | Audible audiobook format |
| 0x00173406 | ` Audible` | Known | Audible audiobook format |
| 0x0017344A | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x00176F88 | `De Audible-software in dit product wordt gebruikt in licentie van Audible. Copyr` | Known | Audible audiobook format |
| 0x00176FDF | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x0017A47C | `Audible-programvaren i dette produktet brukes under lisens fra Audible. Copyrigh` | Known | Audible audiobook format |
| 0x0017A4D0 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0017DE10 | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x0017DE3F | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x0017DE56 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x001814FB | ` Audible ` | Known | Audible audiobook format |
| 0x0018150D | ` Audible ` | Known | Audible audiobook format |
| 0x00181532 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x00184BC0 | `Audible ` | Known | Audible audiobook format |
| 0x00184BD4 | ` Audible ` | Known | Audible audiobook format |
| 0x00184BFE | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x00194D1C | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x00194D71 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |

---

## 5. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0015A124 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunhofer IIS og THOMSON multi` | Known | MP3 codec |
| 0x0015DC5B | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und THOMSON Multimedia.` | Known | MP3 codec |
| 0x00161291 | `n de audio MPEG Layer-3 utilizada bajo licencia de Fraunhofer IIS y THOMSON mult` | Known | MP3 codec |
| 0x001647BC | `MPEG Layer-3 -` | Known | MP3 codec |
| 0x001647CE | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | MP3 codec |
| 0x00168978 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | MP3 codec |
| 0x0016BEE8 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da Fraunhofer IIS e THOMSON` | Known | MP3 codec |
| 0x0016FAEC | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x0016FB38 | `Fraunhofer IIS ` | Known | MP3 codec |
| 0x001735D4 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x001735FB | ` Fraunhofer IIS` | Known | MP3 codec |
| 0x0017716C | `Technologie voor codering van MPEG Layer-3-audio in licentie van Fraunhofer IIS ` | Known | MP3 codec |
| 0x0017A64C | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fraunhofer IIS og THOMSON m` | Known | MP3 codec |
| 0x0017DFF0 | `Ljudkodningstekniken MPEG Layer-3 ` | Known | MP3 codec |
| 0x0017E026 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | MP3 codec |
| 0x001816A4 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x001816C6 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x00184D5C | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00184D81 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x00194EE8 | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |
| 0x001C420F | `Boot time (MP3PlayerExampleApp constructor)` | Known | MP3 codec |

---

## 6. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001AFE9D | `#!#iTunesDB` | Known | iTunes database |

---

## 7. Clock/Alarms

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0030B1D6 | `Is24HrClock` | Known | Clock system |

---

## 8. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001513BC | `FirewireGuid` | Known | FireWire interface |
| 0x00151528 | `FirewireInitiator` | Known | FireWire interface |
| 0x0015153C | `FirewireInterrupt` | Known | FireWire interface |
| 0x00151550 | `FirewireHandling` | Known | FireWire interface |
| 0x00159873 | ` den kan bruges som FireWire-disk, og tr` | Known | FireWire interface |
| 0x0015AB40 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x0015D20B | `ffnen Sie das Addressbuch, Microsoft Entourage oder Palm Desktop und exportieren` | Known | FireWire interface |
| 0x0015E706 | `ber FireWire verbunden` | Known | FireWire interface |
| 0x001608FC | `gido FireWire y arrastrar los archivos vCard a la carpeta Contacts del iPod. Par` | Known | FireWire interface |
| 0x00161DA0 | `FireWire conectado` | Known | FireWire interface |
| 0x00163EEE | `sin, avaa Osoitekirja, Microsoft Entourage tai Palm Desktop ja vie yhteystiedot ` | Known | FireWire interface |
| 0x001652A0 | `FireWire liitetty` | Known | FireWire interface |
| 0x00167E4F | `lectionnez Appareils > Ajouter un appareil. Puis choisissez iPod et cliquez sur ` | Known | FireWire interface |
| 0x0016948C | `FireWire Connect` | Known | FireWire interface |
| 0x0016B49F | ` archiviare contatti ed eventi di calendari. Se utilizzi iSync (con Mac OS X v10` | Known | FireWire interface |
| 0x0016C948 | `FireWire Connesso` | Known | FireWire interface |
| 0x0016EEF4 | ` FireWire ` | Known | FireWire interface |
| 0x0017083C | `FireWire ` | Known | FireWire interface |
| 0x00172BAB | `  FireWire ` | Known | FireWire interface |
| 0x001740E4 | `FireWire ` | Known | FireWire interface |
| 0x00176652 | `Op de iPod kunt u adres- en agendagegevens opslaan. Als u met iSync werkt (Mac O` | Known | FireWire interface |
| 0x00177C00 | `FireWire aangesloten` | Known | FireWire interface |
| 0x00179CFE | `pner du Adressebok, Microsoft Entourage eller Palm Desktop og eksporterer kontak` | Known | FireWire interface |
| 0x00179E5E | `ringer i tillegg til musikken din. Microsoft Outlook, Microsoft Outlook Express ` | Known | FireWire interface |
| 0x0017A114 | `Hvis du vil vise tekstfiler her, aktiverer du iPod for bruk som FireWire-disk og` | Known | FireWire interface |
| 0x0017B084 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x0017D6F2 | `ll sedan in din iPod som FireWire-h` | Known | FireWire interface |
| 0x0017D878 | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x0017EA7C | `FireWire anslutet` | Known | FireWire interface |
| 0x00180DD2 | ` FireWire ` | Known | FireWire interface |
| 0x001820B8 | `FireWire ` | Known | FireWire interface |
| 0x00184449 | ` FireWire ` | Known | FireWire interface |
| 0x001857B8 | `FireWire ` | Known | FireWire interface |
| 0x00194522 | `Your iPod can store contacts and calendar events. If you're using iSync (with Ma` | Known | FireWire interface |
| 0x00195930 | `FireWire Connected` | Known | FireWire interface |
| 0x0030B141 | `FirewirePower` | Known | FireWire interface |
| 0x0035E700 | `Running on FireWire power.` | Known | FireWire interface |
| 0x0035EA17 | `Turn on/off firewire` | Known | FireWire interface |
| 0x0044EB64 | `FirewireInterrupt` | Known | FireWire interface |
| 0x0044EB78 | `FirewireHandling` | Known | FireWire interface |
| 0x0044EBC4 | `FirewirePower` | Known | FireWire interface |

---

## 9. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001512FC | `diskModeImageRev` | Known | Hardware interface |
| 0x0019444C | `Disk Mode` | Known | Hardware interface |
| 0x001C46E4 | `C:\iPod\tagged_checkout\sobek_build\Q14\Sources\Services\I2C\I2CResponder.cpp` | Known | Hardware interface |
| 0x00360C73 | ` Send a command to the LCD` | Known | Hardware interface |
| 0x0041F470 | `D3LCD2.8` | Known | Hardware interface |
| 0x0041F49C | `D3 LCD   2.8V` | Known | Hardware interface |
| 0x0044EDD0 | `C:\iPod\tagged_checkout\sobek_build\Q14\Sources\Services\I2C\I2CResponder.cpp` | Known | Hardware interface |

---

## 10. Storage (ATA/Disk)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0035F874 | `!ATADisk is FDISK format.` | Known | ATA/disk interface |

---

## 11. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001514C8 | `USBPowerSense` | Known | Power management |
| 0x001514D8 | `PowerStateControl` | Known | Power management |
| 0x00151564 | `PCFPowerMgr` | Known | Power management |
| 0x00194000 | `Charging` | Known | Power management |
| 0x0019597C | `Low Battery` | Known | Power management |
| 0x0030B134 | `BatteryLevel` | Known | Power management |
| 0x0035E73C | `Power status is invalid.` | Known | Power management |
| 0x0035E8DF | `@Battery reading is %d (%x) E%4cF` | Known | Power management |
| 0x0044EB8C | `PowerStateControl` | Known | Power management |
| 0x0044EBA0 | `USBPower` | Known | Power management |

---

## 12. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00159AE8 | `Alarmer` | Known | UI element |
| 0x0015AA94 | `Alarmer` | Known | UI element |
| 0x0015DA1B | `hlen" den Alarm beenden` | Known | UI element |
| 0x0016000C | `Calendario` | Known | UI element |
| 0x00160018 | `Calendarios` | Known | UI element |
| 0x00160B94 | `Calendario` | Known | UI element |
| 0x00160BA0 | `Calendarios` | Known | UI element |
| 0x00160BBC | `Alarmas` | Known | UI element |
| 0x00161078 | `Alarma` | Known | UI element |
| 0x001615C8 | `Alarma` | Known | UI element |
| 0x00161766 | `Calendario` | Known | UI element |
| 0x00161980 | `Alarma` | Known | UI element |
| 0x00161BE0 | `Calendarios` | Known | UI element |
| 0x00161CB0 | `Alarma` | Known | UI element |
| 0x00161CEC | `Alarmas` | Known | UI element |
| 0x00168278 | `Alarmes` | Known | UI element |
| 0x0016873C | `Alarme` | Known | UI element |
| 0x00168C68 | `Alarme` | Known | UI element |
| 0x00168CDC | `Alarme` | Known | UI element |
| 0x00169048 | `Alarme` | Known | UI element |
| 0x00169384 | `Alarme` | Known | UI element |
| 0x001693D0 | `Alarmes` | Known | UI element |
| 0x0016ADB0 | `Calendario` | Known | UI element |
| 0x0016ADBC | `Calendari` | Known | UI element |
| 0x0016B880 | `Calendario` | Known | UI element |
| 0x0016B88C | `Calendari` | Known | UI element |
| 0x0016C32E | `Calendario` | Known | UI element |
| 0x0016C788 | `Calendari` | Known | UI element |
| 0x0017A01C | `Alarmer` | Known | UI element |
| 0x0017AC9C | `Alarmtidspunkt` | Known | UI element |
| 0x0017AFD4 | `Alarmer` | Known | UI element |
| 0x0017E688 | `Alarmtid` | Known | UI element |
| 0x00193A58 | `Calendar` | Known | UI element |
| 0x00193A64 | `Calendars` | Known | UI element |
| 0x001948AC | `Calendar` | Known | UI element |
| 0x001948B8 | `Calendars` | Known | UI element |
| 0x001948D0 | `Alarms` | Known | UI element |
| 0x0019520C | `Alarm Clock` | Known | UI element |
| 0x00195332 | `Calendar` | Known | UI element |
| 0x00195530 | `Alarm Time` | Known | UI element |
| 0x0019553C | `Alarm Clock` | Known | UI element |
| 0x00195758 | `Calendars` | Known | UI element |
| 0x00195824 | `Alarm Clock` | Known | UI element |
| 0x00195878 | `Alarms` | Known | UI element |
| 0x001C318B | `Calendars` | Known | UI element |
| 0x001C3195 | `Calendars\` | Known | UI element |
| 0x0030B1F0 | `CurAlarm` | Known | UI element |
| 0x0030B1F9 | `CurAlarmText` | Known | UI element |

---

## 13. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0015C9B0 | `Extras` | Known | Menu item |
| 0x0015E10C | `Extras` | Known | Menu item |
| 0x0015E76C | `Extras` | Known | Menu item |
| 0x00160038 | `Extras` | Known | Menu item |
| 0x00161788 | `Extras` | Known | Menu item |
| 0x00161E04 | `Extras` | Known | Menu item |
| 0x001675E8 | `Albums` | Known | Menu item |
| 0x001675F0 | `Genres` | Known | Menu item |
| 0x00167680 | `Extras` | Known | Menu item |
| 0x001676F0 | `Albums` | Known | Menu item |
| 0x00168E3C | `Extras` | Known | Menu item |
| 0x00168E66 | `Genres` | Known | Menu item |
| 0x00168E7E | `Albums` | Known | Menu item |
| 0x00168FDC | `Genres` | Known | Menu item |
| 0x00168FF0 | `Albums` | Known | Menu item |
| 0x00169194 | `Photos` | Known | Menu item |
| 0x001693E4 | `Genres` | Known | Menu item |
| 0x001693F8 | `Albums` | Known | Menu item |
| 0x00169510 | `Extras` | Known | Menu item |
| 0x00175F04 | `Albums` | Known | Menu item |
| 0x00175F0C | `Genres` | Known | Menu item |
| 0x00176008 | `Albums` | Known | Menu item |
| 0x0017767E | `Genres` | Known | Menu item |
| 0x00177696 | `Albums` | Known | Menu item |
| 0x001777D4 | `Genres` | Known | Menu item |
| 0x001777E4 | `Albums` | Known | Menu item |
| 0x00177B70 | `Genres` | Known | Menu item |
| 0x00177B84 | `Albums` | Known | Menu item |
| 0x001939E0 | `Now Playing` | Known | Menu item |
| 0x001939EC | `Artists` | Known | Menu item |
| 0x00193A04 | `Albums` | Known | Menu item |
| 0x00193A0C | `Genres` | Known | Menu item |
| 0x00193A14 | `Composers` | Known | Menu item |
| 0x00193A8C | `Extras` | Known | Menu item |
| 0x00193A94 | `Playlists` | Known | Menu item |
| 0x00193AA0 | `Audiobooks` | Known | Menu item |
| 0x00193AD0 | `Shuffle Songs` | Known | Menu item |
| 0x00193AF4 | `Albums` | Known | Menu item |
| 0x00194418 | `Now Playing` | Known | Menu item |
| 0x001952F0 | `Shuffle Songs` | Known | Menu item |
| 0x00195350 | `Extras` | Known | Menu item |
| 0x0019535A | `Audiobooks` | Known | Menu item |
| 0x0019536A | `Composers` | Known | Menu item |
| 0x00195376 | `Genres` | Known | Menu item |
| 0x0019538A | `Albums` | Known | Menu item |
| 0x00195396 | `Artists` | Known | Menu item |
| 0x001953A2 | `Playlists` | Known | Menu item |
| 0x00195478 | `Main Menu` | Known | Menu item |
| 0x001954D0 | `Audiobooks` | Known | Menu item |
| 0x001954DC | `Composers` | Known | Menu item |
| 0x001954E8 | `Genres` | Known | Menu item |
| 0x001954F8 | `Albums` | Known | Menu item |
| 0x00195500 | `Artists` | Known | Menu item |
| 0x00195508 | `Playlists` | Known | Menu item |
| 0x00195514 | `Settings` | Known | Menu item |
| 0x00195630 | `Photos` | Known | Menu item |
| 0x00195880 | `Settings` | Known | Menu item |
| 0x0019588C | `Genres` | Known | Menu item |
| 0x00195894 | `Artists` | Known | Menu item |
| 0x0019589C | `Albums` | Known | Menu item |
| 0x001958A4 | `Composers` | Known | Menu item |
| 0x001958C0 | `Audiobooks` | Known | Menu item |
| 0x001958FC | `Playlists` | Known | Menu item |
| 0x00195990 | `Extras` | Known | Menu item |
| 0x001959A0 | `Main Menu` | Known | Menu item |

---

## 14. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00151360 | `iPod_Control\Device` | Filesystem Path | |
| 0x00151374 | `iPod_Control` | Filesystem Path | |
| 0x00151384 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x001AFE3C | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001AFEAC | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001AFED4 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x001AFF04 | `iPod_Control\Device\_short_deepsleep` | Filesystem Path | |
| 0x001AFF2C | `iPod_Control\Device\_no_deepsleep` | Filesystem Path | |
| 0x001AFF50 | `iPod_Control\Device\_show_voltage` | Filesystem Path | |
| 0x001AFF74 | `iPod_Control\Device\_disable_cache` | Filesystem Path | |
| 0x001AFF98 | `System_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x001AFFBC | `iPod_Control\Device` | Filesystem Path | |
| 0x001AFFDC | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x001C31B9 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001C45D0 | `iPod_Control\Testing\` | Filesystem Path | |
| 0x001C45E8 | `iPod_Control\Testing\\TestLog.txt` | Filesystem Path | |
| 0x001C460C | `iPod_Control\Testing\\Tests.Lock` | Filesystem Path | |
| 0x001C467C | `iPod_Control\iTunes\` | Filesystem Path | |

---

## 15. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001593B0 | `Acoustic` | EQ Preset | |
| 0x001593BC | `Bass Booster` | EQ Preset | |
| 0x001593DC | `Classical` | EQ Preset | |
| 0x001593F8 | `Electronic` | EQ Preset | |
| 0x0015940C | `Hip Hop` | EQ Preset | |
| 0x00159424 | `Loudness` | EQ Preset | |
| 0x00159430 | `Lounge` | EQ Preset | |
| 0x00159454 | `Small Speakers` | EQ Preset | |
| 0x00159464 | `Spoken Word` | EQ Preset | |
| 0x00159470 | `Treble Booster` | EQ Preset | |
| 0x00159490 | `Vocal Booster` | EQ Preset | |
| 0x0015CD78 | `Acoustic` | EQ Preset | |
| 0x0015CDB8 | `Electronic` | EQ Preset | |
| 0x0015CDCC | `Hip Hop` | EQ Preset | |
| 0x0015CDE4 | `Loudness` | EQ Preset | |
| 0x00160434 | `Hip Hop` | EQ Preset | |
| 0x00160444 | `Latina` | EQ Preset | |
| 0x0016044C | `Loudness` | EQ Preset | |
| 0x00160458 | `Lounge` | EQ Preset | |
| 0x00163B48 | `Lounge` | EQ Preset | |
| 0x00167A94 | `Hip Hop` | EQ Preset | |
| 0x00167AAC | `Loudness` | EQ Preset | |
| 0x0016B1C4 | `Hip Hop` | EQ Preset | |
| 0x0016B1D4 | `Latina` | EQ Preset | |
| 0x0016B1DC | `Loudness` | EQ Preset | |
| 0x0016B1E8 | `Lounge` | EQ Preset | |
| 0x0016E8CC | `Acoustic` | EQ Preset | |
| 0x0016E8D8 | `Bass Booster` | EQ Preset | |
| 0x0016E8F8 | `Classical` | EQ Preset | |
| 0x0016E914 | `Electronic` | EQ Preset | |
| 0x0016E928 | `Hip Hop` | EQ Preset | |
| 0x0016E940 | `Loudness` | EQ Preset | |
| 0x0016E94C | `Lounge` | EQ Preset | |
| 0x0016E970 | `Small Speakers` | EQ Preset | |
| 0x0016E980 | `Spoken Word` | EQ Preset | |
| 0x0016E98C | `Treble Booster` | EQ Preset | |
| 0x0016E9AC | `Vocal Booster` | EQ Preset | |
| 0x0017266C | `Acoustic` | EQ Preset | |
| 0x00172678 | `Bass Booster` | EQ Preset | |
| 0x00172698 | `Classical` | EQ Preset | |
| 0x001726B4 | `Electronic` | EQ Preset | |
| 0x001726C8 | `Hip Hop` | EQ Preset | |
| 0x001726E0 | `Loudness` | EQ Preset | |
| 0x001726EC | `Lounge` | EQ Preset | |
| 0x00172710 | `Small Speakers` | EQ Preset | |
| 0x00172720 | `Spoken Word` | EQ Preset | |
| 0x0017272C | `Treble Booster` | EQ Preset | |
| 0x0017274C | `Vocal Booster` | EQ Preset | |
| 0x00176390 | `Loudness` | EQ Preset | |
| 0x0017639C | `Lounge` | EQ Preset | |
| 0x00179900 | `Hip Hop` | EQ Preset | |
| 0x00179910 | `Latino` | EQ Preset | |
| 0x00179918 | `Loudness` | EQ Preset | |
| 0x00179924 | `Lounge` | EQ Preset | |
| 0x0017D260 | `Acoustic` | EQ Preset | |
| 0x0017D26C | `Bass Booster` | EQ Preset | |
| 0x0017D28C | `Classical` | EQ Preset | |
| 0x0017D2A8 | `Electronic` | EQ Preset | |
| 0x0017D2BC | `Hip Hop` | EQ Preset | |
| 0x0017D2D4 | `Loudness` | EQ Preset | |
| 0x0017D2E0 | `Lounge` | EQ Preset | |
| 0x0017D304 | `Small Speakers` | EQ Preset | |
| 0x0017D314 | `Spoken Word` | EQ Preset | |
| 0x0017D320 | `Treble Booster` | EQ Preset | |
| 0x0017D340 | `Vocal Booster` | EQ Preset | |
| 0x001808F0 | `Acoustic` | EQ Preset | |
| 0x001808FC | `Bass Booster` | EQ Preset | |
| 0x0018091C | `Classical` | EQ Preset | |
| 0x00180938 | `Electronic` | EQ Preset | |
| 0x0018094C | `Hip Hop` | EQ Preset | |
| 0x00180964 | `Loudness` | EQ Preset | |
| 0x00180970 | `Lounge` | EQ Preset | |
| 0x00180994 | `Small Speakers` | EQ Preset | |
| 0x001809A4 | `Spoken Word` | EQ Preset | |
| 0x001809B0 | `Treble Booster` | EQ Preset | |
| 0x001809D0 | `Vocal Booster` | EQ Preset | |
| 0x00183F40 | `Acoustic` | EQ Preset | |
| 0x00183F4C | `Bass Booster` | EQ Preset | |
| 0x00183F6C | `Classical` | EQ Preset | |
| 0x00183F88 | `Electronic` | EQ Preset | |
| 0x00183F9C | `Hip Hop` | EQ Preset | |
| 0x00183FB4 | `Loudness` | EQ Preset | |
| 0x00183FC0 | `Lounge` | EQ Preset | |
| 0x00183FE4 | `Small Speakers` | EQ Preset | |
| 0x00183FF4 | `Spoken Word` | EQ Preset | |
| 0x00184000 | `Treble Booster` | EQ Preset | |
| 0x00184020 | `Vocal Booster` | EQ Preset | |
| 0x00193E7C | `Acoustic` | EQ Preset | |
| 0x00193E88 | `Bass Booster` | EQ Preset | |
| 0x00193EA8 | `Classical` | EQ Preset | |
| 0x00193EC4 | `Electronic` | EQ Preset | |
| 0x00193ED8 | `Hip Hop` | EQ Preset | |
| 0x00193EF0 | `Loudness` | EQ Preset | |
| 0x00193EFC | `Lounge` | EQ Preset | |
| 0x00193F20 | `Small Speakers` | EQ Preset | |
| 0x00193F30 | `Spoken Word` | EQ Preset | |
| 0x00193F3C | `Treble Booster` | EQ Preset | |
| 0x00193F5C | `Vocal Booster` | EQ Preset | |

---

## 16. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00160F7C | `Error durante la  importanci` | Diagnostic | |
| 0x0016C9B8 | `Errore` | Diagnostic | |
| 0x001C3FAC | `%s Error in file %s.` | Diagnostic | |
| 0x001C48B3 | `Error loading operating system. Setup cannot continue.` | Diagnostic | |
| 0x00315E7C | `setct-ErrorTBS` | Diagnostic | |
| 0x00355E7B | `( <- Error:Unsupported` | Diagnostic | |
| 0x00355E94 | ` <- Error:Internal Error` | Diagnostic | |
| 0x00355EDC | ` <- Error: %s` | Diagnostic | |
| 0x003566AC | `***Error r_reg Unknown Ide Reg(%X)` | Diagnostic | |
| 0x00356788 | `***Error w_reg Unknown Ide Reg(%X)` | Diagnostic | |
| 0x0035F85C | `Error opening device.` | Diagnostic | |
| 0x003639CC | `SP_ERR_COMPATIBILITY` | Diagnostic | |
| 0x003639E4 | `SP_ERR_MAJOR_VERSION` | Diagnostic | |
| 0x003639FC | `SP_ERR_COMP_VERSION` | Diagnostic | |
| 0x00363A10 | `SP_ERR_BAD_MODULE_ID` | Diagnostic | |
| 0x00363A28 | `SP_ERR_BAD_UNIT_NUMBER` | Diagnostic | |
| 0x00363A40 | `SP_ERR_BAD_INSTANCE` | Diagnostic | |
| 0x00363A54 | `SP_ERR_BAD_HANDLE` | Diagnostic | |
| 0x00363A68 | `SP_ERR_BAD_INDEX` | Diagnostic | |
| 0x00363A7C | `SP_ERR_BAD_PARAMETER` | Diagnostic | |
| 0x00363A94 | `SP_ERR_NO_INSTANCES` | Diagnostic | |
| 0x00363AA8 | `SP_ERR_NO_COMPONENT` | Diagnostic | |
| 0x00363B24 | `SP_ERR_NO_RESOURCES` | Diagnostic | |
| 0x00363B38 | `SP_ERR_INSTANCE_IN_USE` | Diagnostic | |
| 0x00363B50 | `SP_ERR_RESOURCE_OWNED` | Diagnostic | |
| 0x00363B68 | `SP_ERR_RESOURCE_NOT_OWNED` | Diagnostic | |
| 0x00363B84 | `SP_ERR_INCONSISTENT_PARAMS` | Diagnostic | |
| 0x00363BA0 | `SP_ERR_NOT_INITIALIZED` | Diagnostic | |
| 0x00363BB8 | `SP_ERR_NOT_ENABLED` | Diagnostic | |
| 0x00363BCC | `SP_ERR_NOT_SUPPORTED` | Diagnostic | |

---

## 17. Assertions

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0001BB3C | `*** assertion failed: %s, file %s, line %d` | Assertion | |
| 0x0030F07C | `%s(%d): OpenSSL internal error, assertion failed: %s` | Assertion | |
| 0x0036425C | `SP_ERR_ASSERTION` | Assertion | |
| 0x00442F7C | `*** assertion failed: %s, file %s, line %d` | Assertion | |

---
