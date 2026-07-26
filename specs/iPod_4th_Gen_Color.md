# iPod 4th Generation (Color/Photo) - RetailOS 11.1.2.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 11.1.2.1 |
| **IPSW** | iPod_11.1.2.1.ipsw |
| **Device** | iPod 4th Gen Color / Photo (2004, 20-60GB, Color LCD) |
| **Binary Size** | 6,514,176 bytes (6.21 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 6,514,176 bytes |
| **Total Strings (>=6)** | 12,965 |
| **Function Prologues** | 9,777 |
| **SoC** | PortalPlayer PP5021 |
| **Architecture** | ARM7TDMI (ARMv4T), dual-core |
| **Encrypted** | No |
| **Decryption Method** | Extract from IPSW (unencrypted) |
| **SHA-256** | `55845b4694263be104e8bfded72f11d1b1d5b9cbeec64f9ffaced80b0bcdc2f5` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001C9E78 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x004FE144 | `Root Hub Driver Internal Error unused case in hub handl...` | Hidden | Undocumented UI |

---

## 2. Discovered Features

### EQ Preset

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001D2A08 | `Acoustic` | EQ Preset | |
| 0x001D2A14 | `Bass Booster` | EQ Preset | |
| 0x001D2A34 | `Classical` | EQ Preset | |
| 0x001D2A50 | `Electronic` | EQ Preset | |
| 0x001D2A64 | `Hip Hop` | EQ Preset | |
| 0x001D2A7C | `Loudness` | EQ Preset | |
| 0x001D2A88 | `Lounge` | EQ Preset | |
| 0x001D2AAC | `Small Speakers` | EQ Preset | |
| 0x001D2ABC | `Spoken Word` | EQ Preset | |
| 0x001D2AC8 | `Treble Booster` | EQ Preset | |
| 0x001D2AE8 | `Vocal Booster` | EQ Preset | |
| 0x001D6E7C | `USA/Rockies (NZ)` | EQ Preset | |
| 0x001D6E90 | `USA/Rockies (SZ)` | EQ Preset | |
| 0x001DB524 | `Latina` | EQ Preset | |
| 0x001E45EC | `Latino` | EQ Preset | |
| 0x002A1581 | `LATIN-1` | EQ Preset | |
| 0x002A1589 | `LATIN1` | EQ Preset | |
| 0x004A92AF | `~ BR&B$"` | EQ Preset | |

### Localization

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001D0458 | `x-mac-japanese` | Localization | |
| 0x002710A0 | `English` | Localization | |
| 0x002710D8 | `Italiano` | Localization | |
| 0x002A15EC | `X-MAC-JAPANESE` | Localization | |
| 0x002A15FB | `MAC-JAPANESE` | Localization | |
| 0x002A1608 | `MACJAPANESE` | Localization | |
| 0x002A1634 | `X-MAC-CHINESETRAD` | Localization | |
| 0x002A1646 | `MAC-CHINESETRAD` | Localization | |
| 0x002A1663 | `X-MAC-CHINESESIMP` | Localization | |
| 0x002A1675 | `MAC-CHINESESIMP` | Localization | |
| 0x002A1695 | `X-MAC-KOREAN` | Localization | |
| 0x002A16A2 | `MAC-KOREAN` | Localization | |

### Filesystem Path

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001C9D54 | `iPod_Control\Device` | Filesystem Path | |
| 0x001C9D68 | `iPod_Control` | Filesystem Path | |
| 0x001C9D78 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x001CA518 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001CA5B0 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x001CA604 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CA6A0 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x001CA6EC | `iPod_Control\Music\` | Filesystem Path | |
| 0x002A0EB1 | `iPod_Control/iTunes/` | Filesystem Path | |

### Assertion

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00005280 | `*** assertion failed: %s, file %s, line %d` | Assertion | |

---

## 3. Known User-Facing Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001C9E10 | `KeyRepeatTimer` | Known | UI element |
| 0x001D2ED2 | `k derefter vCards til mappen Contacts p` | Known | UI element |
| 0x001D3079 | `kke vCard-arkiverne til mappen "Contacts". Arkiverne bl...` | Known | UI element |
| 0x001D3250 | `Alarmer` | Known | UI element |
| 0x001D4050 | `Nulstil menu` | Known | Menu item |
| 0x001D4240 | `Hovedmenu` | Known | Menu item |
| 0x001D4800 | `Menuer` | Known | Menu item |
| 0x001D6CEC | `Extras` | Known | UI element |
| 0x001D7664 | `Contacts` | Known | UI element |
| 0x001D77E0 | `ffnen Sie ihn durch Doppelklicken in das iPod Symbol un...` | Known | UI element |
| 0x001D7B57 | `nnen sich hier Textdateien anzeigen lassen, indem Sie d...` | Known | UI element |
| 0x001D8133 | `hlen" beendet Alarm` | Known | UI element |
| 0x001DB090 | `Calendario` | Known | UI element |
| 0x001DB09C | `Calendarios` | Known | UI element |
| 0x001DB7F4 | `El iPod puede almacenar contactos y eventos de calendar...` | Known | UI element |
| 0x001DB9D6 | `gido y arrastrar los archivos vCard a la carpeta Contac...` | Known | UI element |
| 0x001DBB69 | `n de usar el iPod como disco y hacer doble clic en el i...` | Known | UI element |
| 0x001DBDC4 | `Alarmas` | Known | UI element |
| 0x001DBF41 | `gido y arrastrar los archivos de texto a la carpeta Not...` | Known | UI element |
| 0x001DC524 | `Alarma` | Known | UI element |
| 0x001DD26C | `Hora alarma` | Known | UI element |
| 0x001DD58C | `Contraste` | Known | UI element |
| 0x001DFE0E | ` vCardit iPodin Contacts-kansioon. Lis` | Known | UI element |
| 0x001DFF7C | ` vCardit iPodin Contacts-kansioon. Ne haetaan automaatt...` | Known | UI element |
| 0x001E02DA | ` tekstitiedostot iPodin Notes-kansioon. Lis` | Known | UI element |
| 0x001E48F4 | `Votre iPod peut stocker des contacts et des ` | Known | UI element |
| 0x001E4990 | `lectionnez Appareils > Ajouter un appareil. Puis choisi...` | Known | UI element |
| 0x001E4A28 | `adresses, Microsoft Entourage ou Palm Desktop et export...` | Known | UI element |
| 0x001E4A9D | `utiliser comme disque dur. Puis faites glisser les vCar...` | Known | UI element |
| 0x001E4BAC | ` mille contacts. Les applications Microsoft Outlook, Mi...` | Known | UI element |
| 0x001E4EF8 | `Alarmes` | Known | UI element |
| 0x001E505E | `utilisation du disque, puis faites glisser ces fichiers...` | Known | UI element |
| 0x001E565C | `Chargement des notes.` | Known | UI element |
| 0x001E5698 | `Alarme` | Known | UI element |
| 0x001E56A0 | `Chargement des contacts.` | Known | UI element |
| 0x001E5E5B | `init. menu p.` | Known | Menu item |
| 0x001E60B8 | `Menu princ.` | Known | Menu item |
| 0x001E6430 | `H. alarme` | Known | UI element |
| 0x001E872C | `Calendari` | Known | UI element |
| 0x001E8E6D | ` archiviare contatti ed eventi di calendari. Se utilizz...` | Known | UI element |
| 0x001EA1CC | `Ripr. Menu Princ.` | Known | Menu item |
| 0x001EAA28 | `Contrasto` | Known | UI element |
| 0x001F243B | ` Contacts ` | Known | UI element |
| 0x001F261D | ` "Contacts" ` | Known | UI element |
| 0x001F2A08 | ` Notes ` | Known | UI element |
| 0x001F65EC | `Shuffle nummers` | Known | UI element |
| 0x001F6F10 | `De iPod biedt ruimte voor de gegevens van maar liefst d...` | Known | UI element |
| 0x001F73E0 | `Om tekstbestanden te bekijken, stelt u de iPod in als h...` | Known | UI element |
| 0x001F8004 | `Shuffle foto's` | Known | UI element |
| 0x001F8144 | `Herstel menu` | Known | Menu item |
| 0x001F834C | `Shuffle` | Known | UI element |
| 0x001F8354 | `Hoofdmenu` | Known | Menu item |
| 0x001F8938 | `Menu's` | Known | Menu item |
| 0x001F8940 | `Contrast` | Known | UI element |
| 0x001FB1A4 | `ringene til Contacts-mappen p` | Known | UI element |
| 0x001FB328 | ` iPod-symbolet, og flytt vCard-filene inn i Contacts-ma...` | Known | UI element |
| 0x001FC7B8 | `Alarmtidspunkt` | Known | UI element |
| 0x001FF96F | ` skrivbordet och drar in vCard-filerna i mappen "Contac...` | Known | UI element |
| 0x001FFCDD | `ge och drar sedan in textfilerna i mappen "Notes" p` | Known | UI element |
| 0x00200EA0 | `Alarmtid` | Known | UI element |
| 0x00270F30 | `Now Playing` | Known | UI element |
| 0x00270FC0 | `Calendar` | Known | UI element |
| 0x00270FCC | `Calendars` | Known | UI element |
| 0x00270FE0 | `Backlight` | Known | UI element |
| 0x00271038 | `Shuffle Songs` | Known | UI element |
| 0x00271FA4 | `Alarms` | Known | UI element |
| 0x0027239C | `Slideshow Settings` | Known | User setting |
| 0x002725C0 | `Notes loading.` | Known | UI element |
| 0x0027262C | `Contacts loading.` | Known | UI element |
| 0x00272BB4 | `Shuffle Photos` | Known | UI element |
| 0x00272BC4 | `Repeat` | Known | UI element |
| 0x00272C38 | `Sleep Timer` | Known | UI element |
| 0x00272C44 | `Alarm Clock` | Known | UI element |
| 0x00272D1C | `Reset Main Menu` | Known | Menu item |
| 0x00272E78 | `Reset All Settings` | Known | User setting |
| 0x00272F34 | `Backlight Timer` | Known | UI element |
| 0x00272F54 | `Main Menu` | Known | Menu item |
| 0x00272FFC | `Settings` | Known | User setting |
| 0x00273230 | `Alarm Time` | Known | UI element |
| 0x0027353C | `Reset All` | Known | UI element |
| 0x002A0EE4 | `vcalendar` | Known | UI element |
| 0x002A0F98 | `dalarm` | Known | UI element |
| 0x002A1110 | `valarm` | Known | UI element |
| 0x002A21BC | `To check song links, set the preference NotesOnly to tr...` | Known | UI element |
| 0x002A2308 | `The NotesOnly pref can only be set globally in the Pref...` | Known | UI element |
| 0x002A23E4 | `Warning: Preferences file must be in the Notes folder, ...` | Known | UI element |
| 0x002A2704 | `Memory full. %d notes loaded, some notes not loaded. (3...` | Known | UI element |
| 0x004FF032 | `Illegal instruction` | Known | UI element |
| 0x004FF060 | `Illegal address` | Known | UI element |
| 0x004FF166 | `NotesOnly` | Known | UI element |

---

## 4. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001C99A4 | `RtcTaskClass` | Known | RTOS task thread |
| 0x001C99B4 | `TimerTaskClass` | Known | RTOS task thread |
| 0x001C9A6C | `WatchdogTask` | Known | RTOS task thread |
| 0x001C9A7C | `AlarmTask` | Known | RTOS task thread |
| 0x001C9A94 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x001C9AA8 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x001C9AB8 | `TopPlugTask` | Known | RTOS task thread |
| 0x001C9AC4 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x001C9AD4 | `PlayBtnTask` | Known | RTOS task thread |
| 0x001C9AE0 | `PrvBtnTask` | Known | RTOS task thread |
| 0x001C9AEC | `NextBtnTask` | Known | RTOS task thread |
| 0x001C9AF8 | `ActionBtnTask` | Known | RTOS task thread |
| 0x001C9B08 | `MenuBtnTask` | Known | RTOS task thread |
| 0x001C9B14 | `DiskMgrTask` | Known | RTOS task thread |
| 0x001C9B30 | `CNATask` | Known | RTOS task thread |
| 0x001C9B38 | `BacklightTask` | Known | RTOS task thread |
| 0x001C9B48 | `SerialOptoTask` | Known | RTOS task thread |
| 0x001C9B58 | `OptoTask` | Known | RTOS task thread |
| 0x001C9B64 | `FirewireTask` | Known | RTOS task thread |
| 0x001C9E30 | `LcdUpdateTask` | Known | RTOS task thread |
| 0x001C9E44 | `HostOSTask` | Known | RTOS task thread |
| 0x001C9E78 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x001D0218 | `LoadDataTasks` | Known | RTOS task thread |
| 0x001D0500 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x001D07FF | `5RunTestsTask` | Known | RTOS task thread |
| 0x002953B4 | `USBDeviceTask` | Known | RTOS task thread |
| 0x00295574 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x002A1504 | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x002A151C | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x002C1A58 | `FWStateClearSet_CSR_Task` | Known | RTOS task thread |
| 0x002C1F60 | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x002C1F98 | `FX_DisplayTask` | Known | RTOS task thread |
| 0x002C1FA8 | `FX_RenderTask` | Known | RTOS task thread |
| 0x004FBD40 | `USBStatusTask` | Known | RTOS task thread |
| 0x004FBD50 | `USBTaskTimeTask` | Known | RTOS task thread |
| 0x004FD99C | `USB Secondary Interrupt Task` | Known | RTOS task thread |
| 0x004FE0A0 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x004FE0B4 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x004FE1D8 | `ICAPTPCameraIOTask` | Known | RTOS task thread |
| 0x004FE1F4 | `SBP2CommandTask` | Known | RTOS task thread |

---

## 5. Audio/Codec

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0014DCC8 | `RIFFWAVEfmt data` | Known | PCM audio format |
| 0x001B3194 | `RIFFWAVEfmt dataD` | Known | PCM audio format |
| 0x001D3928 | `Audible-softwaren i dette produkt bruges i henhold til ...` | Known | Audible audiobook format |
| 0x001D3988 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x001D3A7A | `.net codec i dette produkt bruges i henhold til en lice...` | Known | Audio system |
| 0x001D3B24 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunh...` | Known | Audio system |
| 0x001D8160 | `Die Audible Software in diesem Produkt wird in Lizenz v...` | Known | Audible audiobook format |
| 0x001D81B9 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x001D82A9 | `.net Codec in diesem Produkt wird mit der Lizenz der Vo...` | Known | Audio system |
| 0x001D836F | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und ...` | Known | Audio system |
| 0x001DC540 | `El software Audible incluido en este producto se usa ba...` | Known | Audible audiobook format |
| 0x001DC59B | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x001DC739 | `n de audio MPEG Layer-3 utilizada bajo licencia de Frau...` | Known | Audio system |
| 0x001E084E | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright...` | Known | Audible audiobook format |
| 0x001E0888 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x001E09F4 | `MPEG Layer-3 -` | Known | Audio system |
| 0x001E0A06 | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | Audio system |
| 0x001E56BC | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x001E5706 | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x001E571B | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x001E57CC | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x001E58A0 | `La technologie de codage audio MPEG Layer 3 est utilis` | Known | Audio system |
| 0x001E58D8 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | Audio system |
| 0x001E9B04 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x001E9B2D | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x001E9B5C | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x001E9BCE | ` utilizzato su licenza da VoiceAge Corporation. Il code...` | Known | Audio system |
| 0x001E9CA4 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da...` | Known | Audio system |
| 0x001EE58D | ` Audible ` | Known | Audible audiobook format |
| 0x001EE5AE | `Audible ` | Known | Audible audiobook format |
| 0x001EE607 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x001EE7BC | `MPEG Layer-3 ` | Known | Audio system |
| 0x001EE808 | `Fraunhofer IIS ` | Known | Audio system |
| 0x001F302E | ` Audible` | Known | Audible audiobook format |
| 0x001F3072 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x001F3162 | `.net codec` | Known | Audio system |
| 0x001F3223 | ` Fraunhofer IIS` | Known | Audio system |
| 0x001F79E4 | `De Audible-software in dit product wordt gebruikt in li...` | Known | Audible audiobook format |
| 0x001F7A3B | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x001F7B2C | `.net-codec in dit product wordt gebruikt in licentie va...` | Known | Audio system |
| 0x001F7BC8 | `Technologie voor codering van MPEG Layer-3-audio in lic...` | Known | Audio system |
| 0x001FBC04 | `Audible-programvaren i dette produktet brukes under lis...` | Known | Audible audiobook format |
| 0x001FBC58 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x001FBDD4 | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fr...` | Known | Audio system |
| 0x0020025C | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x0020028B | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x002002A2 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x0020043C | `Ljudkodningstekniken MPEG Layer-3 ` | Known | Audio system |
| 0x00200472 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | Audio system |
| 0x002047F6 | ` Fraunhofer IIS ` | Known | Audio system |
| 0x00272640 | `The Audible software in this product is used under lice...` | Known | Audible audiobook format |
| 0x00272779 | `.net codec in this product is used under license from V...` | Known | Audio system |
| 0x0027280C | `MPEG Layer-3 audio coding technology licensed from Frau...` | Known | Audio system |
| 0x00272BA0 | `TV Out` | Known | Audio system |
| 0x002A201D | `&Aacute` | Known | AAC codec |
| 0x002A20E3 | `&aacute` | Known | AAC codec |
| 0x004F7D41 | ``0aLaCfDf` | Known | Apple Lossless codec |
| 0x004FD23C | `Audible` | Known | Audible audiobook format |
| 0x004FD314 | `AudioCodecs` | Known | Audio system |
| 0x004FDA18 | `mp4_aacdec_sync` | Known | AAC codec |
| 0x004FDA28 | `mp3dec_sync` | Known | MP3 codec |

---

## 6. Storage/Hardware

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00003A00 | `!ATAsoso` | Known | ATA/disk interface |
| 0x00003A28 | `!ATAdpua` | Known | ATA/disk interface |
| 0x0000509D | `diskmode` | Known | Hardware interface |
| 0x000050A6 | `diskscan` | Known | Hardware interface |
| 0x000AD79C | `atadmrts\|@-` | Known | ATA/disk interface |
| 0x000D4B18 | `atadmhdp` | Known | ATA/disk interface |
| 0x000D65A8 | `atadmhbddbhmmhsd>@-` | Known | ATA/disk interface |
| 0x000D736C | `atadmhfddfhmmhsd` | Known | ATA/disk interface |
| 0x000D79E4 | `atadmhpo` | Known | ATA/disk interface |
| 0x00116748 | `nutiatad` | Known | ATA/disk interface |
| 0x001C9878 | `data abort` | Known | ATA/disk interface |
| 0x001C9A3C | `FirewireInitiator` | Known | FireWire interface |
| 0x001C9A50 | `FirewireHandler` | Known | FireWire interface |
| 0x001C9CF0 | `diskModeImageRev` | Known | Hardware interface |
| 0x001C9DB0 | `FirewireGuid` | Known | FireWire interface |
| 0x001CA148 | `I2C write Error` | Known | Hardware interface |
| 0x001CA15C | `I2C read Error %02x` | Known | Hardware interface |
| 0x001CA61C | `Photos\Photo Database` | Known | ATA/disk interface |
| 0x001CA678 | `Photo Database` | Known | ATA/disk interface |
| 0x001D04E8 | `Photo Import Database` | Known | ATA/disk interface |
| 0x001D2574 | `Spiller nu` | Known | Hardware interface |
| 0x001D2648 | `Spillelister` | Known | Hardware interface |
| 0x001D2664 | `Genoptag spil` | Known | Hardware interface |
| 0x001D2B18 | `Slet spilleliste` | Known | Hardware interface |
| 0x001D2B2C | `Arkiver spilleliste` | Known | Hardware interface |
| 0x001D2BCC | `Ny spilleliste %lu` | Known | Hardware interface |
| 0x001D2C58 | `Harddisk` | Known | Hardware interface |
| 0x001D2E50 | `bne Adressebog, Microsoft Entourage eller Palm Desktop ...` | Known | Hardware interface |
| 0x001D3023 | `r du har tilsluttet iPod som disk, skal du dobbeltklikk...` | Known | Hardware interface |
| 0x001D339A | ` den kan bruges som disk og anbringe tekstarkiver i map...` | Known | Hardware interface |
| 0x001D376F | ` afspilningsknappen p` | Known | Hardware interface |
| 0x001D3DAB | `je den til spillelisten On-The-Go. Hold knappen nede, n` | Known | Hardware interface |
| 0x001D3DE4 | `r en spilleliste, kunstner eller et album er valgt for ...` | Known | Hardware interface |
| 0x001D3E21 | `je alle sangene til spillelisten On-The-Go.` | Known | Hardware interface |
| 0x001D3FDC | `Nyt spil` | Known | Hardware interface |
| 0x001D4300 | `Dette mediearkiv kan ikke vises eller afspilles p` | Known | Hardware interface |
| 0x001D4944 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x001D6CA0 | `Spiele` | Known | Hardware interface |
| 0x001D6D14 | `Weiterspielen` | Known | Hardware interface |
| 0x001D78B8 | `Beispiel` | Known | Hardware interface |
| 0x001D78D0 | `Beispielfirma GmbH` | Known | Hardware interface |
| 0x001D78E4 | `Dieses Beispiel zeigt, welche Infos Sie bei einem Konta...` | Known | Hardware interface |
| 0x001D7980 | `Beispielstadt` | Known | Hardware interface |
| 0x001D79A4 | `Beispielstra` | Known | Hardware interface |
| 0x001D79C0 | `Beispielort` | Known | Hardware interface |
| 0x001D8848 | `Neues Spiel` | Known | Hardware interface |
| 0x001D8BB0 | `Die Mediendatei kann nicht auf dem iPod angezeigt oder ...` | Known | Hardware interface |
| 0x001D9246 | `ber FireWire verbunden` | Known | FireWire interface |
| 0x001DD6B0 | `FireWire conectado` | Known | FireWire interface |
| 0x001DFA30 | `Diskanttivahv.` | Known | Hardware interface |
| 0x001DFA40 | `Diskanttiheik.` | Known | Hardware interface |
| 0x001DFB00 | `Ladataan` | Known | ATA/disk interface |
| 0x001E07E4 | `Muistiinpanoja ladataan.` | Known | ATA/disk interface |
| 0x001E0820 | `Yhteystietoja ladataan.` | Known | ATA/disk interface |
| 0x001E15D4 | `nityksen jatkamiseen ei ole tarpeeksi vapaata levytilaa...` | Known | ATA/disk interface |
| 0x001E1614 | `nityksen aloittamiseen ei ole tarpeeksi vapaata levytil...` | Known | ATA/disk interface |
| 0x001E1874 | `FireWire liitetty` | Known | FireWire interface |
| 0x001E68BC | `FireWire Connect` | Known | FireWire interface |
| 0x001E9778 | `Durata diapositiva` | Known | ATA/disk interface |
| 0x001E9F98 | `Spazzata dal centro` | Known | ATA/disk interface |
| 0x001E9FAC | `Spazzata verso il basso` | Known | ATA/disk interface |
| 0x001E9FC4 | `Spazzata diagonale` | Known | ATA/disk interface |
| 0x001E9FD8 | `Spinta verso il basso` | Known | Hardware interface |
| 0x001E9FF0 | `Spinta diagonale` | Known | Hardware interface |
| 0x001EA0C0 | `Data & Ora` | Known | ATA/disk interface |
| 0x001EA190 | `Imposta Data & Ora` | Known | ATA/disk interface |
| 0x001EAB5C | `FireWire Connesso` | Known | FireWire interface |
| 0x001EFAA4 | `FireWire ` | Known | FireWire interface |
| 0x001F6CB0 | `Op de iPod kunt u contact- en agendagegevens bewaren. A...` | Known | Hardware interface |
| 0x001F76D4 | `Handmatig` | Known | Hardware interface |
| 0x001F8A88 | `FireWire aangesloten` | Known | FireWire interface |
| 0x001FA834 | `Spilles n` | Known | Hardware interface |
| 0x001FA914 | `Fortsett spill` | Known | Hardware interface |
| 0x001FAD7C | `Mer diskant` | Known | Hardware interface |
| 0x001FAD88 | `Mindre diskant` | Known | Hardware interface |
| 0x001FADD0 | `Slett spilleliste` | Known | Hardware interface |
| 0x001FAF24 | `Diskmodus` | Known | Hardware interface |
| 0x001FAFD0 | `Privatadresse` | Known | ATA/disk interface |
| 0x001FB107 | `pner du Adressebok, Microsoft Entourage eller Palm Desk...` | Known | Hardware interface |
| 0x001FB25C | `ringer i tillegg til musikken din. Microsoft Outlook, M...` | Known | Hardware interface |
| 0x001FB644 | `Hvis du vil vise tekstfiler her, aktiverer du iPod for ...` | Known | Hardware interface |
| 0x001FB9C4 | `r bildene til datamaskinen, og synkroniser dem via iTun...` | Known | ATA/disk interface |
| 0x001FBA52 | ` avspillingsknappen p` | Known | Hardware interface |
| 0x001FC037 | ` legge den til i On-The-Go-spillelisten. Spillelister, ...` | Known | Hardware interface |
| 0x001FC24C | `Nytt spill` | Known | Hardware interface |
| 0x001FC584 | `Denne mediefilen kan ikke vises eller spilles p` | Known | Hardware interface |
| 0x001FC5CB | ` datamaskinen ved hjelp av QuickTime.` | Known | ATA/disk interface |
| 0x001FC62C | `r importerte bilder til datamaskinen, og synkroniser vi...` | Known | ATA/disk interface |
| 0x001FC934 | `Det er ikke nok ledig diskplass til ` | Known | Hardware interface |
| 0x001FCBA0 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x001FF7C6 | `rddisk och drar in vCard-filerna i kontaktmappen i iPod...` | Known | Hardware interface |
| 0x001FF916 | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x001FF936 | `rddisk. Sedan dubbelklickar du bara p` | Known | Hardware interface |
| 0x001FFCD4 | `rddiskl` | Known | Hardware interface |
| 0x001FFDA4 | `inget kort inmatat` | Known | ATA/disk interface |
| 0x00200688 | `Stort bildmaterial` | Known | Hardware interface |
| 0x002012B0 | `FireWire anslutet` | Known | FireWire interface |
| 0x00271A00 | `Disk Mode` | Known | Hardware interface |
| 0x00271AD4 | `Your iPod can store contacts and calendar events. If yo...` | Known | Hardware interface |
| 0x00271C9C | `Your iPod can store up to one thousand contacts right a...` | Known | Hardware interface |
| 0x002720A8 | `To view text files here, enable iPod for disk use, then...` | Known | Hardware interface |
| 0x00272C90 | `Disk Browser` | Known | Hardware interface |
| 0x0027339C | `There is not enough free disk space to continue recordi...` | Known | Hardware interface |
| 0x002733D8 | `There is not enough free disk space to start recording.` | Known | Hardware interface |
| 0x00273638 | `FireWire Connected` | Known | FireWire interface |
| 0x0027364C | `No battery power remains. Please connect iPod to power.` | Known | Power management |
| 0x00273684 | `Low Battery` | Known | Power management |
| 0x00295610 | `TROMResourceDB - unknown header version! (Try regenerat...` | Known | ATA/disk interface |
| 0x002A1927 | `Bad Data` | Known | ATA/disk interface |
| 0x002A22D0 | `Bad meta data, name not a recognized preference. (27)` | Known | ATA/disk interface |
| 0x002A2388 | `Bad meta data, name not found. (23)` | Known | ATA/disk interface |
| 0x002A23AC | `Bad meta data, name termination quote not found. (24)` | Known | ATA/disk interface |
| 0x002A2434 | `Bad meta data, content not found. (25)` | Known | ATA/disk interface |
| 0x002A245C | `Bad meta data, content termination quote not found. (26...` | Known | ATA/disk interface |
| 0x002A2498 | `Bad meta data, encoding not found. (8)` | Known | ATA/disk interface |
| 0x002A24C0 | `Bad meta data, encoding termination quote not found. (9...` | Known | ATA/disk interface |
| 0x002A24FC | `Bad meta data, unknown encoding. (17)` | Known | ATA/disk interface |
| 0x002A26B8 | `Bad data. (32)` | Known | ATA/disk interface |
| 0x002C1F34 | `MEMDISK` | Known | Hardware interface |
| 0x004ECB63 | `ataTaza[aea;ajaaaVa)b'b+b+dMd[d]dtdvdrdsd}dudfd` | Known | ATA/disk interface |
| 0x004F174D | `aGa>a(a'aJa?a<a,a4a=aBaDasawaXaYaZakataoaeaqa_a]aSaua` | Known | ATA/disk interface |
| 0x004FBC58 | `USB MSC` | Known | USB interface |
| 0x004FD2F8 | `MaximumDataRate` | Known | ATA/disk interface |
| 0x004FD320 | `FireWireVersion` | Known | FireWire interface |
| 0x004FD460 | `FireWire` | Known | FireWire interface |
| 0x00505BCB | `<key>FireWireGUID</key>` | Known | FireWire interface |

---

## 7. Error Messages

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00005280 | `*** assertion failed: %s, file %s, line %d` | Known | Error/assertion message |
| 0x0000C904 | `Invalid Operation` | Known | Error/assertion message |
| 0x001DC434 | `Error durante la importaci` | Known | Error/assertion message |
| 0x001E99E1 | ` verificato un errore durante l'importazione` | Known | Error/assertion message |
| 0x00272174 | `connection failed` | Known | Error/assertion message |
| 0x002723C0 | `Imported photos cannot be viewed on TV. Transfer photos...` | Known | Error/assertion message |
| 0x00272514 | `An error occurred while importing` | Known | Error/assertion message |
| 0x00273008 | `This file cannot be viewed on iPod.` | Known | Error/assertion message |
| 0x0027302C | `This media file cannot be viewed or played on iPod. Use...` | Known | Error/assertion message |
| 0x00273124 | `This photo format cannot be viewed on iPod. Transfer im...` | Known | Error/assertion message |
| 0x00273364 | `Cannot record because there is no microphone attached.` | Known | Error/assertion message |
| 0x002A2220 | `Too many errors, further errors discarded. (31)` | Known | Error/assertion message |
| 0x002A2270 | `%s Error in file %s.` | Known | Error/assertion message |
| 0x002A2544 | `Cannot link to a .link file. (29)` | Known | Error/assertion message |
| 0x004FE070 | `error:%08lX:%s:%s:%s` | Known | Error/assertion message |
| 0x004FE0F0 | `Root hub Error Calling Add Device` | Known | Error/assertion message |
| 0x004FE144 | `Root Hub Driver Internal Error unused case in hub handl...` | Known | Error/assertion message |
| 0x004FE634 | `internal error: list index %ld out of range` | Known | Error/assertion message |

---

## 8. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001CFCD0 | ` !"#$%&'()*+,-./0123456789:;<=>?@abcdefghijklmnopqrstuv...` | Known | Filesystem path |
| 0x001CFDD0 | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x001D311C | `Eksempelfirma A/S` | Known | Filesystem path |
| 0x001D3188 | `apple.com/dk/support/ipod` | Known | Filesystem path |
| 0x001D35C4 | `Point: %d (%d/%d)` | Known | Filesystem path |
| 0x001D39FB | `rende VoiceAge Corporation i USA og/eller andre lande o...` | Known | Filesystem path |
| 0x001D6E14 | `USA/Hawaii (NZ)` | Known | Filesystem path |
| 0x001D6E24 | `USA/Hawaii (SZ)` | Known | Filesystem path |
| 0x001D6E34 | `USA/Alaska (NZ)` | Known | Filesystem path |
| 0x001D6E44 | `USA/Alaska (SZ)` | Known | Filesystem path |
| 0x001D6E54 | `USA/Pazifik (NZ)` | Known | Filesystem path |
| 0x001D6E68 | `USA/Pazifik (SZ)` | Known | Filesystem path |
| 0x001D6E7C | `USA/Rockies (NZ)` | Known | Filesystem path |
| 0x001D6E90 | `USA/Rockies (SZ)` | Known | Filesystem path |
| 0x001D6EA4 | `USA/Zentral (NZ)` | Known | Filesystem path |
| 0x001D6EB8 | `USA/Zentral (SZ)` | Known | Filesystem path |
| 0x001D6ECC | `USA/Ost (NZ)` | Known | Filesystem path |
| 0x001D6EDC | `USA/Ost (SZ)` | Known | Filesystem path |
| 0x001D71E0 | `Vorn./Nachn.` | Known | Filesystem path |
| 0x001D71F0 | `Nachn./Vorn.` | Known | Filesystem path |
| 0x001D7940 | `apple.com/de/support/ipod` | Known | Filesystem path |
| 0x001D79E0 | `089/9 87 65 43 21` | Known | Filesystem path |
| 0x001D79F4 | `089/9 87 65 43 20` | Known | Filesystem path |
| 0x001D7A08 | `0171/9 87 65 43 21` | Known | Filesystem path |
| 0x001D7DF8 | `Punkte: %d (%d/%d)` | Known | Filesystem path |
| 0x001D7F84 | `ber die Start/Pause-Taste von jedem ausgew` | Known | Filesystem path |
| 0x001D81F3 | ` ist entweder eine eingetragene Marke oder eine Marke d...` | Known | Filesystem path |
| 0x001DBCEC | `apple.com/es/support/ipod` | Known | Filesystem path |
| 0x001DC1AC | `Result.: %d (%d/%d)` | Known | Filesystem path |
| 0x001DC5DB | ` es una marca registrada o una marca comercial de Voice...` | Known | Filesystem path |
| 0x001DCBC8 | `Fecha/hora` | Known | Filesystem path |
| 0x001DFAF8 | `%d / %d` | Known | Filesystem path |
| 0x001E0084 | `apple.com/fi/support/ipod` | Known | Filesystem path |
| 0x001E039C | `%s / %s` | Known | Filesystem path |
| 0x001E03C8 | `%d / %d valokuvaa tuotu` | Known | Filesystem path |
| 0x001E04D4 | `Tulos: %d (%d/%d)` | Known | Filesystem path |
| 0x001E08C3 | ` on VoiceAge Corporationin Yhdysvalloissa ja/tai muissa...` | Known | Filesystem path |
| 0x001E4E1C | `apple.com/fr/support/ipod` | Known | Filesystem path |
| 0x001E52E3 | `sult. : %d (%d/%d)` | Known | Filesystem path |
| 0x001E579E | `tats-Unis et/ou dans d` | Known | Filesystem path |
| 0x001E5E17 | `gler date/heure` | Known | Filesystem path |
| 0x001E92C8 | `apple.com/it/support/ipod` | Known | Filesystem path |
| 0x001E975C | `Punti: %d (%d/%d)` | Known | Filesystem path |
| 0x001EC8D4 | `%b/%-d %-I:%M %2p` | Known | Filesystem path |
| 0x001EC8E8 | `%-m/%-d` | Known | Filesystem path |
| 0x001EC90C | `%y/%-m/%d` | Known | Filesystem path |
| 0x001EC918 | `%Y/%b/%-d` | Known | Filesystem path |
| 0x001EDB9C | `apple.com/jp/support/ipod` | Known | Filesystem path |
| 0x001EE110 | ` %d (%d/%d)` | Known | Filesystem path |
| 0x001F1890 | `%Y/%B/%d %A  %I:%M:%S %p` | Known | Filesystem path |
| 0x001F18AC | `%Y/%B/%d` | Known | Filesystem path |
| 0x001F18C4 | `%-m/%-d %-I:%M %p` | Known | Filesystem path |
| 0x001F18F8 | `%Y/%-m/%d` | Known | Filesystem path |
| 0x001F2738 | `apple.co.kr/support/ipod` | Known | Filesystem path |
| 0x001F2B14 | `%d / %d ` | Known | Filesystem path |
| 0x001F2C76 | `: %d (%d/%d)` | Known | Filesystem path |
| 0x001F71CC | `apple.com/nl/support/ipod` | Known | Filesystem path |
| 0x001F769C | `Score: %d (%d/%d)` | Known | Filesystem path |
| 0x001F7A73 | ` is een gedeponeerd handelsmerk of een handelsmerk van ...` | Known | Filesystem path |
| 0x001F8108 | `Stel datum/tijd in` | Known | Filesystem path |
| 0x001FB45C | `apple.com/no/support/ipod` | Known | Filesystem path |
| 0x001FB8A4 | `Poeng: %d (%d/%d)` | Known | Filesystem path |
| 0x001FBC97 | ` er enten et registrert varemerke eller et varemerke fo...` | Known | Filesystem path |
| 0x001FEE14 | `%-d/%-m` | Known | Filesystem path |
| 0x001FFA90 | `apple.com/support/se/ipod` | Known | Filesystem path |
| 0x001FFEF8 | `ng: %d (%d/%d)` | Known | Filesystem path |
| 0x00200328 | `r VoiceAge Corporation i USA och/eller andra l` | Known | Filesystem path |
| 0x00200968 | `ll in datum/tid` | Known | Filesystem path |
| 0x00203030 | `%Y/%m/%d` | Known | Filesystem path |
| 0x00203E50 | `apple.com.cn/support/ipod` | Known | Filesystem path |
| 0x002041DD | ` %d/%d ` | Known | Filesystem path |
| 0x0020869D | `%d (%d/%d)` | Known | Filesystem path |
| 0x00270DF8 | `%-m/%d/%y` | Known | Filesystem path |
| 0x00271ED0 | `apple.com/support/ipod` | Known | Filesystem path |
| 0x002726CB | ` is either registered trademark or trademark of VoiceAg...` | Known | Filesystem path |
| 0x00294EF0 | `Created: %d/%d/%4d %d:%02d:%02d %s` | Known | Filesystem path |
| 0x00294F14 | `Last Accessed: %d/%d/%4d %2d:%02d:%02d %s` | Known | Filesystem path |
| 0x00294F40 | `Modified: %d/%d/%4d %2d:%02d:%02d %s` | Known | Filesystem path |
| 0x002A0E79 | `Contacts/` | Known | Filesystem path |
| 0x002A0E8D | `Calendars/` | Known | Filesystem path |
| 0x002A0EA5 | `Recordings/` | Known | Filesystem path |
| 0x002A0EB1 | `iPod_Control/iTunes/` | Known | Filesystem path |
| 0x002A16C5 | `file://` | Known | Filesystem path |
| 0x002A16CD | `image://` | Known | Filesystem path |
| 0x002A189C | `</TITLE>` | Known | Filesystem path |
| 0x002A18AC | `</BODY>` | Known | Filesystem path |
| 0x002A18DA | `</ROT13>` | Known | Filesystem path |
| 0x002A2588 | `Bad link, no matching </a> for anchor tag. (20)` | Known | Filesystem path |
| 0x002C1ED3 | `` ,;=+[]*?<>\|":/\` | Known | Filesystem path |
| 0x002C1EE5 | `\/:*?"<>\|` | Known | Filesystem path |
| 0x002D979F | `W/}lE>q` | Known | Filesystem path |
| 0x0033AC8D | `H."0*Bx/` | Known | Filesystem path |
| 0x003425EA | `U/~RERT` | Known | Filesystem path |
| 0x00347352 | `TUOPT/\|` | Known | Filesystem path |
| 0x0034E5E7 | `HuGZp/$j` | Known | Filesystem path |
| 0x00354AF7 | `(bJ)b"(b2""z/` | Known | Filesystem path |
| 0x00357267 | `JUAPDD(/` | Known | Filesystem path |
| 0x0035D576 | `/B\|$BD'` | Known | Filesystem path |
| 0x0035E4F3 | `$Bd$BT/` | Known | Filesystem path |
| 0x003647CB | `/" +J\|!` | Known | Filesystem path |
| 0x0036B64A | `Fb""")/` | Known | Filesystem path |
| 0x0036C801 | `/RyO(UIH` | Known | Filesystem path |
| 0x0036D8B1 | `~$Bh'"~$Bz/` | Known | Filesystem path |
| 0x0037111B | `$B +BZ/` | Known | Filesystem path |
| 0x00378B39 | `0c(HBP/` | Known | Filesystem path |
| 0x0037D0CF | `$B~("\|/` | Known | Filesystem path |
| 0x00393E39 | `T/DDDDD` | Known | Filesystem path |
| 0x003940A3 | `"~UeB /` | Known | Filesystem path |
| 0x00396B55 | `$B((B /` | Known | Filesystem path |
| 0x0039EC38 | ` "\|$B~/` | Known | Filesystem path |
| 0x003A1DB4 | `@$B\|$"(/` | Known | Filesystem path |
| 0x003A2EB4 | `)"8/B""` | Known | Filesystem path |
| 0x003A35FC | `r4c6 bN/` | Known | Filesystem path |
| 0x003A89FD | `RDT%B(/` | Known | Filesystem path |
| 0x003A9B89 | `RBHUE\|/` | Known | Filesystem path |
| 0x003B1209 | `]B""B</` | Known | Filesystem path |
| 0x003B4A86 | `,B\|RED/` | Known | Filesystem path |
| 0x003BA509 | `$BT). /` | Known | Filesystem path |
| 0x003BB709 | `#"TUB(/` | Known | Filesystem path |
| 0x003CB807 | `   ! " # $ % & ' ( ) * + , - . / 0 1 2 3 4 5 6 7 8 9 : ...` | Known | Filesystem path |
| 0x003CBA07 | `! !!!"!#!$!%!&!'!(!)!*!+!,!-!.!/!0!1!2!3!4!5!6!7!8!9!:!...` | Known | Filesystem path |
| 0x003CBC07 | `" "!"""#"$"%"&"'"(")"*"+","-"."/"0"1"2"3"4"5"6"7"8"9":"...` | Known | Filesystem path |
| 0x00413725 | `4 84444/` | Known | Filesystem path |
| 0x0041398F | `Up8)Ut4)Up8/` | Known | Filesystem path |
| 0x0041895B | `+/h4\|40` | Known | Filesystem path |
| 0x004254F7 | `/EU@~5V@y5V` | Known | Filesystem path |
| 0x004414B0 | `UW5\U\/` | Known | Filesystem path |
| 0x0046035A | `-)WF=0/X` | Known | Filesystem path |
| 0x0046B5E4 | `5UU\5UU\5UU\5UU\5UU\5UU\5UU\5UU\5UU\/` | Known | Filesystem path |
| 0x0048F441 | `/" %BD"` | Known | Filesystem path |
| 0x00495D64 | `ODD""(/` | Known | Filesystem path |
| 0x00496F9F | `B"$R%"B$" /` | Known | Filesystem path |
| 0x0049790C | `bG\|jG\|/` | Known | Filesystem path |
| 0x004994F6 | `$E$$BR/` | Known | Filesystem path |
| 0x00499597 | `dRB~RA$/` | Known | Filesystem path |
| 0x00499EE8 | `TT&T%B(/` | Known | Filesystem path |
| 0x004A8BDC | `)'>$B8/` | Known | Filesystem path |
| 0x004AB1EA | `$B\|%EV/` | Known | Filesystem path |
| 0x004B273E | `BDU!BJ ""/` | Known | Filesystem path |
| 0x004B36A2 | `Z-bD("(%B>/` | Known | Filesystem path |
| 0x004E44AD | `% %!%"%#%$%%%&%'%(%)%*%+%,%-%.%/%0%1%2%3%4%5%6%7%8%9%:%...` | Known | Filesystem path |
| 0x004E4C8D | `qWlIl/Ymg*` | Known | Filesystem path |
| 0x004E64FB | `X)W,W*W3W9W.W/W\W;WBWiW` | Known | Filesystem path |
| 0x004E6BC9 | `n/o6oKoto*o` | Known | Filesystem path |
| 0x004E7161 | `q/q1qsq\qhqEqrqJqxqzq` | Known | Filesystem path |
| 0x004E83E9 | `b6bKbNb/e` | Known | Filesystem path |
| 0x004E859F | `V3W0W(W-W,W/W)W` | Known | Filesystem path |
| 0x004E8673 | `NMOOOGOWO^O4O[OUO0OPOQO=O:O8OCOTO<OFOcO\O`O/ONO6OYO]OHO...` | Known | Filesystem path |
| 0x004E8CD1 | `e%f-f f'f/f` | Known | Filesystem path |
| 0x004E8FD3 | `bNc>c/cUcBcFcOcIc:cPc=c*c+c(cMcLcHeIe` | Known | Filesystem path |
| 0x004E9285 | `fFUjUfUDU^UaUCUJU1UVUOUUU/UdU8U.U\U,UcU3UAUWU` | Known | Filesystem path |
| 0x004E94EB | `\|F}C}q}.}9}<}@}0}3}D}/}B}2}1}=` | Known | Filesystem path |
| 0x004E9673 | `W/X*X4X$X0X1X!X` | Known | Filesystem path |
| 0x004E9789 | `k.l/l,l/n8nTn!n2ngnJn n%n#n` | Known | Filesystem path |
| 0x004E97EB | `r6s%s4s)s:t*t3t"t%t5t6t4t/t` | Known | Filesystem path |
| 0x004EA069 | `S.V;V9V2V?V4V)VSVNVWVtV6V/V0V` | Known | Filesystem path |
| 0x004EA6AF | `\|.~>~F~7~2~C~+~=~1~E~A~4~9~H~5~?~/~D` | Known | Filesystem path |
| 0x004EB69F | `P P'P5P/P1P` | Known | Filesystem path |
| 0x004EB7F1 | `h5h+h-h/hNhDh4h` | Known | Filesystem path |
| 0x004EB805 | `h&h(h.hMh:h%h h,k/k-k1k4kmk` | Known | Filesystem path |
| 0x004EB911 | `w"w'w#x,x"x5x/x(x.x+x!x)x3x*x1xTy[yOy\ySyRyQy` | Known | Filesystem path |
| 0x004EBD9B | `v9w/w-w1w2w4w3w=w%w;w5wHxRxIxMxJxLx&xExPxdygyiyjycykyay` | Known | Filesystem path |
| 0x004EBDF7 | `{1{+{-{/{2{8{` | Known | Filesystem path |
| 0x004EC28B | `t/uoulu` | Known | Filesystem path |
| 0x004EC5F1 | ``IaJa+aEa6a2a.aFa/aOa)a@a bh` | Known | Filesystem path |
| 0x004EC637 | `d&d0d(dAd5d/d` | Known | Filesystem path |
| 0x004EC69F | `hxi4iii@ioiDiviXiAitiLi;iKi7i\iOiQi2iRi/i{i<iFkEkCkBkHk...` | Known | Filesystem path |
| 0x004ECD09 | `y+zJz0z/z(z&z` | Known | Filesystem path |
| 0x004ED5ED | `u/v-v1v=v3v<v5v2v0v` | Known | Filesystem path |
| 0x004ED8E3 | `X-[%[2[#[,['[&[/[.[{[` | Known | Filesystem path |
| 0x004EE211 | `j<p5p/p7p4p1pBp8p?p:p9p@p;p3pAp` | Known | Filesystem path |
| 0x004EE413 | `w-y1y/yT\|S\|` | Known | Filesystem path |
| 0x004EF7DF | `%#%3%+%;%K% %/%(%7%?%` | Known | Filesystem path |
| 0x004EF95E | `02*2+2,2-2.2/2@272B2C292:212>24222;2623252<2=2?282` | Known | Filesystem path |
| 0x004EFE89 | `\7_J_/`P`m`` | Known | Filesystem path |
| 0x004F0065 | `OHSIT>T/Z` | Known | Filesystem path |
| 0x004F00DB | `i_l*mim/n` | Known | Filesystem path |
| 0x004F05CD | `N,p]u/f` | Known | Filesystem path |
| 0x004F12D9 | `S#S/S1S3S8S@SFSES` | Known | Filesystem path |
| 0x004F13D9 | `q4V6V2V8VkVdV/VlVjV` | Known | Filesystem path |
| 0x004F1681 | `_)_-_8_A_H_L_N_/_Q_V_W_Y_a_m_s_w_` | Known | Filesystem path |
| 0x004F1D47 | `s4s/s)s%s>sNsOs` | Known | Filesystem path |
| 0x004F5B4B | `h>kLp/t` | Known | Filesystem path |
| 0x004F61D5 | `o;v/}7~` | Known | Filesystem path |
| 0x004F6F99 | `e1f/h\q6z` | Known | Filesystem path |
| 0x004F75E5 | `UuX/c"dIfKfmh` | Known | Filesystem path |
| 0x004FD018 | `%s<key>%s</key>` | Known | Filesystem path |
| 0x004FD028 | `%s<%s/>` | Known | Filesystem path |
| 0x004FD050 | `%s<key>%d</key>` | Known | Filesystem path |
| 0x004FD078 | `%s</dict>` | Known | Filesystem path |
| 0x004FD0A0 | `%s</array>` | Known | Filesystem path |
| 0x004FD0BC | `%s<string>%s</string>` | Known | Filesystem path |
| 0x004FD0E4 | `%s<integer>%s</integer>` | Known | Filesystem path |
| 0x004FD110 | `%s<real>%s</real>` | Known | Filesystem path |
| 0x004FD134 | `%s<integer>%d</integer>` | Known | Filesystem path |
| 0x004FDA8E | `  !"##$%&&'())*+,-../01234556789:;<=>?@ABCDEFGHIJKMNOPQ...` | Known | Filesystem path |
| 0x004FDCB4 | ` !""#$%&''()*+,-./0123456789:;<>?@ABDEFGIJKMNOQRTUVXY[\...` | Known | Filesystem path |
| 0x0050446A | `/ 0 1 2 3 4 5 6 7 8 9 : ; < = > ? @ A B C D E F G H I J...` | Known | Filesystem path |
| 0x0050464B | `! !!!"!#!$!%!&!'!(!)!*!+!,!-!.!/!0!1!2!3!4!5!6!7!8!9!:!...` | Known | Filesystem path |
| 0x00505B2C | `<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1....` | Known | Filesystem path |
| 0x00505BBA | `</dict>` | Known | Filesystem path |
| 0x00505BC2 | `</plist>` | Known | Filesystem path |
| 0x00505BE3 | `<string>%08lX%08lX</string>` | Known | Filesystem path |
| 0x00506EE4 | `/B'2N6REMQLEVJ\|aViu\J]lLm` | Known | Filesystem path |
| 0x0051AB64 | `/1f;{1Q` | Known | Filesystem path |
| 0x00527B93 | `eJo:~/5pJ` | Known | Filesystem path |
| 0x00529138 | `/2S,ZS{` | Known | Filesystem path |
| 0x0052F204 | `pN/NLZ0` | Known | Filesystem path |
| 0x0052F5F0 | `Iy/w217v` | Known | Filesystem path |
| 0x00535DFA | `f>./okQ` | Known | Filesystem path |
| 0x00536640 | `/H.&LC3` | Known | Filesystem path |
| 0x0053C245 | `Gmb[hD0/1` | Known | Filesystem path |
| 0x00551D39 | `GH*/q?Nv` | Known | Filesystem path |
| 0x0055B6EF | `,1(#/3"'` | Known | Filesystem path |
| 0x0055D743 | `3/&t9UG` | Known | Filesystem path |
| 0x00565813 | `/e9i!Aa` | Known | Filesystem path |
| 0x0056B9D8 | `)$>\:/4V` | Known | Filesystem path |
| 0x0056DE25 | `~UYICO*/]A` | Known | Filesystem path |
| 0x00570F30 | `/gncThYd` | Known | Filesystem path |
| 0x0057339C | `(IZ?:c/`$` | Known | Filesystem path |
| 0x00578501 | `l/_Ky^kv` | Known | Filesystem path |
| 0x0057AC9F | `pM>}=Z/.` | Known | Filesystem path |
| 0x0057BECB | `u4/\wPyWa` | Known | Filesystem path |
| 0x0057D35F | `(,%4`/m` | Known | Filesystem path |
| 0x00586C47 | `wEK1/mJ$` | Known | Filesystem path |
| 0x0058843F | `%/~\|j(&` | Known | Filesystem path |
| 0x005897B7 | `/m'DXbS` | Known | Filesystem path |
| 0x00589873 | `ZR/.wpYA` | Known | Filesystem path |
| 0x00590BCC | `ZLs}=\jX/` | Known | Filesystem path |
| 0x0059EB21 | `/5$PRJ+` | Known | Filesystem path |
| 0x005A286B | `^,,uE./` | Known | Filesystem path |
| 0x005A7FF0 | `u/>@X9"` | Known | Filesystem path |
| 0x005AE346 | `E>j/[Ce` | Known | Filesystem path |
| 0x005C28D7 | `?hy;/,q` | Known | Filesystem path |
| 0x005C50B0 | `/FFni_GF` | Known | Filesystem path |
| 0x005CD930 | `/GOae{K` | Known | Filesystem path |
| 0x005D75F4 | `H2?;7a/` | Known | Filesystem path |
| 0x005D7B0B | `P&A/"M>` | Known | Filesystem path |
| 0x005DA917 | ``QC/~!;` | Known | Filesystem path |
| 0x005E21FC | `J/%/M%\|` | Known | Filesystem path |
| 0x005E401A | `cieOa;/` | Known | Filesystem path |
| 0x005E7CFA | `lt!z;/c` | Known | Filesystem path |
| 0x005E9CC5 | `)Y4.B/`` | Known | Filesystem path |
| 0x005F2101 | `[D/k @W!!` | Known | Filesystem path |
| 0x005F5D3C | `bpA/c-IG` | Known | Filesystem path |
| 0x005FD9AA | `S=>!e/6` | Known | Filesystem path |
| 0x005FE1B5 | `e%t/XEg` | Known | Filesystem path |
| 0x005FE714 | `7oG4?p4T&h7Y/&` | Known | Filesystem path |
| 0x006000F2 | `dv5/@%hc` | Known | Filesystem path |
| 0x00604BB9 | `b];1Y/W` | Known | Filesystem path |
| 0x00607F27 | `g`y)/sC,` | Known | Filesystem path |
| 0x0060C927 | `7$i{</1` | Known | Filesystem path |
| 0x0060D14D | `?5$wR/~e` | Known | Filesystem path |
| 0x006106AB | `/c@G@sA` | Known | Filesystem path |
| 0x00618FF5 | `HWEg/!v[_` | Known | Filesystem path |
| 0x0061BD10 | `Mb.#aj/p` | Known | Filesystem path |
| 0x00628C9A | `/}T2gQ[` | Known | Filesystem path |
| 0x00629CFE | `/wI("9N` | Known | Filesystem path |
| 0x00630381 | `7!Ii/QZ*A` | Known | Filesystem path |
| 0x00631D99 | `{ztbs(/` | Known | Filesystem path |
| 0x00632FC3 | `~qtbo/O` | Known | Filesystem path |
| 0x00634307 | `]\|NI\|Jk/` | Known | Filesystem path |
| 0x00635B6A | `/%W_76\` | Known | Filesystem path |

---

## 9. Video Playback

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00272BA0 | `TV Out` | Known | Video playback |

---

## 10. Binary Structure (for Ghidra/IDA)

| Parameter | Value |
|-----------|-------|
| **Architecture** | ARM (32-bit, little-endian) |
| **Processor** | ARM7TDMI (ARMv4T), dual-core |
| **Base Address** | 0x00000000 |
| **Entry Point** | 0x00000000 |
| **Endianness** | Little-Endian |
| **File Size** | 6,514,176 bytes |

