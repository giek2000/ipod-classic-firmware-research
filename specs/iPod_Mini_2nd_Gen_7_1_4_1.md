# iPod Mini 2nd Generation - RetailOS 1.4.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.4.1 |
| **IPSW** | iPod_7.1.4.1.ipsw |
| **Device** | iPod Mini 2nd Generation (2005, Click Wheel, Anodized Aluminum) |
| **UpdaterFamilyID** | 7 |
| **Binary Size** | 4,506,624 bytes (4.30 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 4,506,624 bytes |
| **Total Strings (>=6)** | 9,856 |
| **Function Prologues** | 10,514 (ARM: 7,561, Thumb: 2,953) |
| **SoC** | PortalPlayer PP5020 |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Encrypted** | No |
| **SHA-256** | `56becf109e5233a4de6e774fb87aa24ef6e8de8f44f30938a0beb6ed486e00f4` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168608 | `MP3ExampleTask` | Hidden | Hidden Test |

---

## 2. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168204 | `WatchdogTask` | Known | RTOS task thread |
| 0x00168214 | `AlarmTask` | Known | RTOS task thread |
| 0x0016822C | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x00168240 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x00168250 | `TopPlugTask` | Known | RTOS task thread |
| 0x0016825C | `HoldSwitchTask` | Known | RTOS task thread |
| 0x0016826C | `PlayBtnTask` | Known | RTOS task thread |
| 0x00168278 | `PrvBtnTask` | Known | RTOS task thread |
| 0x00168284 | `NextBtnTask` | Known | RTOS task thread |
| 0x00168290 | `ActionBtnTask` | Known | RTOS task thread |
| 0x001682A0 | `MenuBtnTask` | Known | RTOS task thread |
| 0x001682AC | `DiskMgrTask` | Known | RTOS task thread |
| 0x001682C8 | `CNATask` | Known | RTOS task thread |
| 0x001682D0 | `BacklightTask` | Known | RTOS task thread |
| 0x001682E0 | `SerialOptoTask` | Known | RTOS task thread |
| 0x001682F0 | `OptoTask` | Known | RTOS task thread |
| 0x001682FC | `FirewireTask` | Known | RTOS task thread |
| 0x001685CC | `HostOSTask` | Known | RTOS task thread |
| 0x0016D763 | `5RunTestsTask` | Known | RTOS task thread |
| 0x001BD464 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x001CEA88 | `FWInterruptHandlerTask` | Known | RTOS task thread |
| 0x00314684 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x00314698 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x003146AC | `SBP2CommandTask` | Known | RTOS task thread |

---

## 3. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016FC3A | `.net codec i dette produkt bruges i henhold til en licensaftale fra VoiceAge Cor` | Known | Audio system |
| 0x001730B1 | `.net Codec in diesem Produkt wird mit der Lizenz der VoiceAge Corporation verwen` | Known | Audio system |
| 0x00178E36 | `.net codec t` | Known | Audio system |
| 0x0017C8A8 | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x0017F756 | ` utilizzato su licenza da VoiceAge Corporation. Il codec ACELP` | Known | Audio system |
| 0x0018606E | `.net codec` | Known | Audio system |
| 0x001894F4 | `.net-codec in dit product wordt gebruikt in licentie van VoiceAge Corporation. G` | Known | Audio system |
| 0x001A4641 | `.net codec in this product is used under license from VoiceAge Corporation. Port` | Known | Audio system |
| 0x00313DB4 | `AudioCodecs` | Known | Audio system |
| 0x00315F63 | `msCodeCom` | Known | Audio system |

---

## 4. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016FAE8 | `Audible-softwaren i dette produkt bruges i henhold til en licensaftale fra Audib` | Known | Audible audiobook format |
| 0x0016FB48 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x00172F68 | `Die Audible Software in diesem Produkt wird in Lizenz von Audible verwendet. Cop` | Known | Audible audiobook format |
| 0x00172FC1 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x00175EAC | `El software Audible incluido en este producto se usa bajo licencia de Audible. C` | Known | Audible audiobook format |
| 0x00175F07 | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x00178D2A | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright ` | Known | Audible audiobook format |
| 0x00178D64 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x0017C798 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x0017C7E2 | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x0017C7F7 | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x0017F68C | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x0017F6B5 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0017F6E4 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x00182B3D | ` Audible ` | Known | Audible audiobook format |
| 0x00182B5E | `Audible ` | Known | Audible audiobook format |
| 0x00182BB7 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x00185F1F | ` Audible ` | Known | Audible audiobook format |
| 0x00185F3A | ` Audible` | Known | Audible audiobook format |
| 0x00185F7E | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x001893AC | `De Audible-software in dit product wordt gebruikt in licentie van Audible. Copyr` | Known | Audible audiobook format |
| 0x00189403 | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x0018C210 | `Audible-programvaren i dette produktet brukes under lisens fra Audible. Copyrigh` | Known | Audible audiobook format |
| 0x0018C264 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0018F538 | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x0018F567 | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x0018F57E | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x00192553 | ` Audible ` | Known | Audible audiobook format |
| 0x00192565 | ` Audible ` | Known | Audible audiobook format |
| 0x00192589 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x0019551C | `Audible ` | Known | Audible audiobook format |
| 0x00195530 | ` Audible ` | Known | Audible audiobook format |
| 0x0019555A | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x001A4508 | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x001A455D | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x00313CDC | `Audible` | Known | Audible audiobook format |

---

## 5. Audio/Codec - Apple Lossless

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00313D50 | `AppleLossless` | Known | Apple Lossless codec |

---

## 6. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016FCE4 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunhofer IIS og THOMSON multi` | Known | MP3 codec |
| 0x00173177 | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und THOMSON Multimedia.` | Known | MP3 codec |
| 0x001760A9 | `n de audio MPEG Layer-3 utilizada bajo licencia de Fraunhofer IIS y THOMSON mult` | Known | MP3 codec |
| 0x00178ECC | `MPEG Layer-3 -` | Known | MP3 codec |
| 0x00178EDE | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | MP3 codec |
| 0x0017C9B4 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | MP3 codec |
| 0x0017F82C | `Tecnologia di codifica audio MPEG Layer-3 su licenza da Fraunhofer IIS e THOMSON` | Known | MP3 codec |
| 0x00182D6C | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00182DB8 | `Fraunhofer IIS ` | Known | MP3 codec |
| 0x00186108 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x0018612F | ` Fraunhofer IIS` | Known | MP3 codec |
| 0x00189590 | `Technologie voor codering van MPEG Layer-3-audio in licentie van Fraunhofer IIS ` | Known | MP3 codec |
| 0x0018C3E0 | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fraunhofer IIS og THOMSON m` | Known | MP3 codec |
| 0x0018F718 | `Ljudkodningstekniken MPEG Layer-3 ` | Known | MP3 codec |
| 0x0018F74E | `n Fraunhofer IIS och THOMSON multimedia.` | Known | MP3 codec |
| 0x001926E0 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00192702 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x001956B8 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x001956DD | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x001A46D4 | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |
| 0x001CE613 | `Boot time (MP3PlayerExampleApp constructor)` | Known | MP3 codec |

---

## 7. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00313CAC | `AppleDRMVersion` | Known | DRM system |
| 0x00313CE4 | `AppleDRM` | Known | DRM system |

---

## 8. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168CD9 | `#!#iTunesDB` | Known | iTunes database |

---

## 9. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001681E8 | `FirewireHandler` | Known | FireWire interface |
| 0x00168548 | `FirewireGuid` | Known | FireWire interface |
| 0x00170594 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x00172A17 | `ffnen Sie das Adressbuch, Microsoft Entourage oder Palm Desktop und exportieren ` | Known | FireWire interface |
| 0x00173A12 | `ber FireWire verbunden` | Known | FireWire interface |
| 0x00176980 | `FireWire conectado` | Known | FireWire interface |
| 0x001797D0 | `FireWire liitetty` | Known | FireWire interface |
| 0x0017C23D | `utiliser comme disque FireWire. Puis glissez les vCards dans le dossier Contacts` | Known | FireWire interface |
| 0x0017D294 | `FireWire Connect` | Known | FireWire interface |
| 0x00180078 | `FireWire Connesso` | Known | FireWire interface |
| 0x00183878 | `FireWire ` | Known | FireWire interface |
| 0x00186A04 | `FireWire ` | Known | FireWire interface |
| 0x00188D1C | `Op de iPod kunt u adres- en agendagegevens opslaan. Als u met iSync werkt (Mac O` | Known | FireWire interface |
| 0x00189E5C | `FireWire aangesloten` | Known | FireWire interface |
| 0x0018CC48 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x0018F1F6 | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x0018FFC0 | `FireWire anslutet` | Known | FireWire interface |
| 0x00192F20 | `FireWire ` | Known | FireWire interface |
| 0x00195F34 | `FireWire ` | Known | FireWire interface |
| 0x001A4F74 | `FireWire Connected` | Known | FireWire interface |
| 0x00313E04 | `FireWire` | Known | FireWire interface |
| 0x00313E74 | `FireWireVersion` | Known | FireWire interface |
| 0x0031D439 | `FireWire` | Known | FireWire interface |
| 0x0031D74F | `<key>FireWireGUID</key>` | Known | FireWire interface |

---

## 10. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168488 | `diskModeImageRev` | Known | Hardware interface |
| 0x001A3E8C | `Disk Mode` | Known | Hardware interface |

---

## 11. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001681F8 | `PCFPowerMgr` | Known | Power management |
| 0x001682B8 | `USBPowerSense` | Known | Power management |
| 0x0016884C | `PowerManager` | Known | Power management |
| 0x001A3A40 | `Charging` | Known | Power management |
| 0x001A4FC0 | `Low Battery` | Known | Power management |
| 0x00313E20 | `PowerInformation` | Known | Power management |

---

## 12. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016F900 | `Alarmer` | Known | UI element |
| 0x001704E4 | `Alarmer` | Known | UI element |
| 0x00172F37 | `hlen" den Alarm beenden` | Known | UI element |
| 0x00175088 | `Calendario` | Known | UI element |
| 0x00175094 | `Calendarios` | Known | UI element |
| 0x00175C50 | `Calendario` | Known | UI element |
| 0x00175C5C | `Calendarios` | Known | UI element |
| 0x00175C78 | `Alarmas` | Known | UI element |
| 0x00175E90 | `Alarma` | Known | UI element |
| 0x001763FC | `Alarma` | Known | UI element |
| 0x0017654E | `Calendario` | Known | UI element |
| 0x00176798 | `Alarma` | Known | UI element |
| 0x001767C0 | `Calendarios` | Known | UI element |
| 0x00176890 | `Alarma` | Known | UI element |
| 0x001768D8 | `Alarmas` | Known | UI element |
| 0x0017C56C | `Alarmes` | Known | UI element |
| 0x0017C774 | `Alarme` | Known | UI element |
| 0x0017CCBC | `Alarme` | Known | UI element |
| 0x0017CD30 | `Alarme` | Known | UI element |
| 0x0017D078 | `Alarme` | Known | UI element |
| 0x0017D174 | `Alarme` | Known | UI element |
| 0x0017D1D4 | `Alarmes` | Known | UI element |
| 0x0017E928 | `Calendario` | Known | UI element |
| 0x0017E934 | `Calendari` | Known | UI element |
| 0x0017F444 | `Calendario` | Known | UI element |
| 0x0017F450 | `Calendari` | Known | UI element |
| 0x0017FC56 | `Calendario` | Known | UI element |
| 0x0017FEAC | `Calendari` | Known | UI element |
| 0x0018C008 | `Alarmer` | Known | UI element |
| 0x0018CA50 | `Alarmtidspunkt` | Known | UI element |
| 0x0018CB94 | `Alarmer` | Known | UI element |
| 0x0018FDD4 | `Alarmtid` | Known | UI element |
| 0x001A3454 | `Calendar` | Known | UI element |
| 0x001A3460 | `Calendars` | Known | UI element |
| 0x001A42DC | `Calendar` | Known | UI element |
| 0x001A42E8 | `Calendars` | Known | UI element |
| 0x001A4300 | `Alarms` | Known | UI element |
| 0x001A4A10 | `Alarm Clock` | Known | UI element |
| 0x001A4AFA | `Calendar` | Known | UI element |
| 0x001A4D28 | `Alarm Time` | Known | UI element |
| 0x001A4D34 | `Alarm Clock` | Known | UI element |
| 0x001A4D84 | `Calendars` | Known | UI element |
| 0x001A4E50 | `Alarm Clock` | Known | UI element |
| 0x001A4EB8 | `Alarms` | Known | UI element |
| 0x001CD27F | `Calendars` | Known | UI element |
| 0x001CD289 | `Calendars\` | Known | UI element |

---

## 13. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016EE6C | `Podcasts` | Known | Menu item |
| 0x001701CA | `Podcasts` | Known | Menu item |
| 0x00170324 | `Podcasts` | Known | Menu item |
| 0x00170510 | `Podcasts` | Known | Menu item |
| 0x00172170 | `Extras` | Known | Menu item |
| 0x001721C8 | `Podcasts` | Known | Menu item |
| 0x00173608 | `Extras` | Known | Menu item |
| 0x00173642 | `Podcasts` | Known | Menu item |
| 0x001737A8 | `Podcasts` | Known | Menu item |
| 0x0017398C | `Podcasts` | Known | Menu item |
| 0x00173A78 | `Extras` | Known | Menu item |
| 0x001750B4 | `Extras` | Known | Menu item |
| 0x00175110 | `Podcasts` | Known | Menu item |
| 0x00176570 | `Extras` | Known | Menu item |
| 0x001765A2 | `Podcasts` | Known | Menu item |
| 0x00176724 | `Podcasts` | Known | Menu item |
| 0x00176900 | `Podcasts` | Known | Menu item |
| 0x001769E4 | `Extras` | Known | Menu item |
| 0x0017B898 | `Albums` | Known | Menu item |
| 0x0017B8A0 | `Genres` | Known | Menu item |
| 0x0017B910 | `Extras` | Known | Menu item |
| 0x0017B964 | `Podcasts` | Known | Menu item |
| 0x0017B9A0 | `Albums` | Known | Menu item |
| 0x0017CE40 | `Extras` | Known | Menu item |
| 0x0017CE6A | `Genres` | Known | Menu item |
| 0x0017CE76 | `Podcasts` | Known | Menu item |
| 0x0017CE8E | `Albums` | Known | Menu item |
| 0x0017D000 | `Genres` | Known | Menu item |
| 0x0017D008 | `Podcasts` | Known | Menu item |
| 0x0017D020 | `Albums` | Known | Menu item |
| 0x0017D1E8 | `Genres` | Known | Menu item |
| 0x0017D1FC | `Podcasts` | Known | Menu item |
| 0x0017D208 | `Albums` | Known | Menu item |
| 0x0017D318 | `Extras` | Known | Menu item |
| 0x001885C4 | `Albums` | Known | Menu item |
| 0x001885CC | `Genres` | Known | Menu item |
| 0x0018868C | `Podcasts` | Known | Menu item |
| 0x001886C4 | `Albums` | Known | Menu item |
| 0x00189A82 | `Genres` | Known | Menu item |
| 0x00189A8E | `Podcasts` | Known | Menu item |
| 0x00189AA6 | `Albums` | Known | Menu item |
| 0x00189BFC | `Genres` | Known | Menu item |
| 0x00189C04 | `Podcasts` | Known | Menu item |
| 0x00189C18 | `Albums` | Known | Menu item |
| 0x00189DC8 | `Genres` | Known | Menu item |
| 0x00189DDC | `Podcasts` | Known | Menu item |
| 0x00189DE8 | `Albums` | Known | Menu item |
| 0x001A33F8 | `Now Playing` | Known | Menu item |
| 0x001A3404 | `Artists` | Known | Menu item |
| 0x001A341C | `Albums` | Known | Menu item |
| 0x001A3424 | `Genres` | Known | Menu item |
| 0x001A342C | `Composers` | Known | Menu item |
| 0x001A3488 | `Extras` | Known | Menu item |
| 0x001A3490 | `Playlists` | Known | Menu item |
| 0x001A349C | `Audiobooks` | Known | Menu item |
| 0x001A34CC | `Shuffle Songs` | Known | Menu item |
| 0x001A34DC | `Podcasts` | Known | Menu item |
| 0x001A3514 | `Albums` | Known | Menu item |
| 0x001A3E58 | `Now Playing` | Known | Menu item |
| 0x001A4AD8 | `Shuffle Songs` | Known | Menu item |
| 0x001A4B18 | `Extras` | Known | Menu item |
| 0x001A4B22 | `Audiobooks` | Known | Menu item |
| 0x001A4B32 | `Composers` | Known | Menu item |
| 0x001A4B3E | `Genres` | Known | Menu item |
| 0x001A4B4A | `Podcasts` | Known | Menu item |
| 0x001A4B5E | `Albums` | Known | Menu item |
| 0x001A4B6A | `Artists` | Known | Menu item |
| 0x001A4B76 | `Playlists` | Known | Menu item |
| 0x001A4C64 | `Main Menu` | Known | Menu item |
| 0x001A4CBC | `Audiobooks` | Known | Menu item |
| 0x001A4CC8 | `Composers` | Known | Menu item |
| 0x001A4CD4 | `Genres` | Known | Menu item |
| 0x001A4CDC | `Podcasts` | Known | Menu item |
| 0x001A4CF0 | `Albums` | Known | Menu item |
| 0x001A4CF8 | `Artists` | Known | Menu item |
| 0x001A4D00 | `Playlists` | Known | Menu item |
| 0x001A4D0C | `Settings` | Known | Menu item |
| 0x001A4E74 | `Audiobooks` | Known | Menu item |
| 0x001A4EC0 | `Settings` | Known | Menu item |
| 0x001A4ECC | `Genres` | Known | Menu item |
| 0x001A4ED4 | `Artists` | Known | Menu item |
| 0x001A4EDC | `Podcasts` | Known | Menu item |
| 0x001A4EE8 | `Albums` | Known | Menu item |
| 0x001A4EF0 | `Composers` | Known | Menu item |
| 0x001A4F0C | `Audiobooks` | Known | Menu item |
| 0x001A4F48 | `Playlists` | Known | Menu item |
| 0x001A4FD4 | `Extras` | Known | Menu item |
| 0x001A5024 | `Main Menu` | Known | Menu item |

---

## 14. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001684EC | `iPod_Control\Device` | Filesystem Path | |
| 0x00168500 | `iPod_Control` | Filesystem Path | |
| 0x00168510 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x00168C78 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x00168CE8 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x00168D10 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x00168D40 | `iPod_Control\Device\` | Filesystem Path | |
| 0x00168D58 | `System_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x00168D7C | `iPod_Control\Device` | Filesystem Path | |
| 0x00168D9C | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x00168DE4 | `iPod_Control\Music\` | Filesystem Path | |
| 0x001BD414 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001CD2AD | `iPod_Control\iTunes\` | Filesystem Path | |

---

## 15. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x003BF345 | `DgFBWA` | Build Path | |

---

## 16. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016F1D8 | `Acoustic` | EQ Preset | |
| 0x0016F1E4 | `Bass Booster` | EQ Preset | |
| 0x0016F204 | `Classical` | EQ Preset | |
| 0x0016F220 | `Electronic` | EQ Preset | |
| 0x0016F234 | `Hip Hop` | EQ Preset | |
| 0x0016F24C | `Loudness` | EQ Preset | |
| 0x0016F258 | `Lounge` | EQ Preset | |
| 0x0016F27C | `Small Speakers` | EQ Preset | |
| 0x0016F28C | `Spoken Word` | EQ Preset | |
| 0x0016F298 | `Treble Booster` | EQ Preset | |
| 0x0016F2B8 | `Vocal Booster` | EQ Preset | |
| 0x00172584 | `Acoustic` | EQ Preset | |
| 0x001725C4 | `Electronic` | EQ Preset | |
| 0x001725D8 | `Hip Hop` | EQ Preset | |
| 0x001725F0 | `Loudness` | EQ Preset | |
| 0x001754FC | `Hip Hop` | EQ Preset | |
| 0x0017550C | `Latina` | EQ Preset | |
| 0x00175514 | `Loudness` | EQ Preset | |
| 0x00175520 | `Lounge` | EQ Preset | |
| 0x001784BC | `Hip Hop` | EQ Preset | |
| 0x001784E0 | `Lounge` | EQ Preset | |
| 0x0017BD78 | `Hip Hop` | EQ Preset | |
| 0x0017BD90 | `Loudness` | EQ Preset | |
| 0x0017ED94 | `Hip Hop` | EQ Preset | |
| 0x0017EDA4 | `Latina` | EQ Preset | |
| 0x0017EDAC | `Loudness` | EQ Preset | |
| 0x0017EDB8 | `Lounge` | EQ Preset | |
| 0x00181DE8 | `Acoustic` | EQ Preset | |
| 0x00181DF4 | `Bass Booster` | EQ Preset | |
| 0x00181E14 | `Classical` | EQ Preset | |
| 0x00181E30 | `Electronic` | EQ Preset | |
| 0x00181E44 | `Hip Hop` | EQ Preset | |
| 0x00181E5C | `Loudness` | EQ Preset | |
| 0x00181E68 | `Lounge` | EQ Preset | |
| 0x00181E8C | `Small Speakers` | EQ Preset | |
| 0x00181E9C | `Spoken Word` | EQ Preset | |
| 0x00181EA8 | `Treble Booster` | EQ Preset | |
| 0x00181EC8 | `Vocal Booster` | EQ Preset | |
| 0x00185498 | `Acoustic` | EQ Preset | |
| 0x001854A4 | `Bass Booster` | EQ Preset | |
| 0x001854C4 | `Classical` | EQ Preset | |
| 0x001854E0 | `Electronic` | EQ Preset | |
| 0x001854F4 | `Hip Hop` | EQ Preset | |
| 0x0018550C | `Loudness` | EQ Preset | |
| 0x00185518 | `Lounge` | EQ Preset | |
| 0x0018553C | `Small Speakers` | EQ Preset | |
| 0x0018554C | `Spoken Word` | EQ Preset | |
| 0x00185558 | `Treble Booster` | EQ Preset | |
| 0x00185578 | `Vocal Booster` | EQ Preset | |
| 0x00188A70 | `Loudness` | EQ Preset | |
| 0x00188A7C | `Lounge` | EQ Preset | |
| 0x0018B900 | `Hip Hop` | EQ Preset | |
| 0x0018B910 | `Latino` | EQ Preset | |
| 0x0018B918 | `Loudness` | EQ Preset | |
| 0x0018B924 | `Lounge` | EQ Preset | |
| 0x0018EBEC | `Acoustic` | EQ Preset | |
| 0x0018EBF8 | `Bass Booster` | EQ Preset | |
| 0x0018EC18 | `Classical` | EQ Preset | |
| 0x0018EC34 | `Electronic` | EQ Preset | |
| 0x0018EC48 | `Hip Hop` | EQ Preset | |
| 0x0018EC60 | `Loudness` | EQ Preset | |
| 0x0018EC6C | `Lounge` | EQ Preset | |
| 0x0018EC90 | `Small Speakers` | EQ Preset | |
| 0x0018ECA0 | `Spoken Word` | EQ Preset | |
| 0x0018ECAC | `Treble Booster` | EQ Preset | |
| 0x0018ECCC | `Vocal Booster` | EQ Preset | |
| 0x00191BE4 | `Acoustic` | EQ Preset | |
| 0x00191BF0 | `Bass Booster` | EQ Preset | |
| 0x00191C10 | `Classical` | EQ Preset | |
| 0x00191C2C | `Electronic` | EQ Preset | |
| 0x00191C40 | `Hip Hop` | EQ Preset | |
| 0x00191C58 | `Loudness` | EQ Preset | |
| 0x00191C64 | `Lounge` | EQ Preset | |
| 0x00191C88 | `Small Speakers` | EQ Preset | |
| 0x00191C98 | `Spoken Word` | EQ Preset | |
| 0x00191CA4 | `Treble Booster` | EQ Preset | |
| 0x00191CC4 | `Vocal Booster` | EQ Preset | |
| 0x00194B54 | `Acoustic` | EQ Preset | |
| 0x00194B60 | `Bass Booster` | EQ Preset | |
| 0x00194B80 | `Classical` | EQ Preset | |
| 0x00194B9C | `Electronic` | EQ Preset | |
| 0x00194BB0 | `Hip Hop` | EQ Preset | |
| 0x00194BC8 | `Loudness` | EQ Preset | |
| 0x00194BD4 | `Lounge` | EQ Preset | |
| 0x00194BF8 | `Small Speakers` | EQ Preset | |
| 0x00194C08 | `Spoken Word` | EQ Preset | |
| 0x00194C14 | `Treble Booster` | EQ Preset | |
| 0x00194C34 | `Vocal Booster` | EQ Preset | |
| 0x001A38BC | `Acoustic` | EQ Preset | |
| 0x001A38C8 | `Bass Booster` | EQ Preset | |
| 0x001A38E8 | `Classical` | EQ Preset | |
| 0x001A3904 | `Electronic` | EQ Preset | |
| 0x001A3918 | `Hip Hop` | EQ Preset | |
| 0x001A3930 | `Loudness` | EQ Preset | |
| 0x001A393C | `Lounge` | EQ Preset | |
| 0x001A3960 | `Small Speakers` | EQ Preset | |
| 0x001A3970 | `Spoken Word` | EQ Preset | |
| 0x001A397C | `Treble Booster` | EQ Preset | |
| 0x001A399C | `Vocal Booster` | EQ Preset | |

---

## 17. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017FF40 | `Errore` | Diagnostic | |
| 0x00180130 | `Errore` | Diagnostic | |
| 0x001CE3B0 | `%s Error in file %s.` | Diagnostic | |
| 0x001CEE97 | `Error loading operating system. Setup cannot continue.` | Diagnostic | |
| 0x00318174 | `setct-ErrorTBS` | Diagnostic | |

---

## 18. Assertions

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00017BC0 | `*** assertion failed: %s, file %s, line %d` | Assertion | |
| 0x00314254 | `%s(%d): OpenSSL internal error, assertion failed: %s` | Assertion | |

---
