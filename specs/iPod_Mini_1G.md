# iPod Mini 1st Generation - RetailOS 6.1.4.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 6.1.4.1 |
| **IPSW** | iPod_6.1.4.1.ipsw |
| **Device** | iPod Mini 1st Generation (2004, 4GB Microdrive) |
| **Binary Size** | 4,506,624 bytes (4.30 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 4,506,624 bytes |
| **Total Strings (>=6)** | 9,944 |
| **Function Prologues** | 7,563 |
| **SoC** | PortalPlayer PP5020 |
| **Architecture** | ARM7TDMI (ARMv4T), dual-core |
| **Encrypted** | No |
| **Decryption Method** | Extract from IPSW (unencrypted) |
| **SHA-256** | `a69031d594a0b54649c0a6cc087241808463b9d94a9e45793cedb7f02abd357f` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168950 | `MP3ExampleTask` | Hidden | Hidden Test |

---

## 2. Discovered Features

### EQ Preset

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
| 0x0017263C | `USA/Rockies (NZ)` | EQ Preset | |
| 0x00172650 | `USA/Rockies (SZ)` | EQ Preset | |
| 0x00175854 | `Latina` | EQ Preset | |
| 0x0018BC58 | `Latino` | EQ Preset | |
| 0x0018C982 | ` navigeringsflaten) for ` | EQ Preset | |
| 0x001CD9D5 | `LATIN-1` | EQ Preset | |
| 0x001CD9DD | `LATIN1` | EQ Preset | |
| 0x002C17B7 | `~ BR&B$"` | EQ Preset | |
| 0x003180AE | `Secure Electronic Transactions` | EQ Preset | |

### Localization

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016D9B0 | `x-mac-japanese` | Localization | |
| 0x001A3874 | `English` | Localization | |
| 0x001A38AC | `Italiano` | Localization | |
| 0x001CDA40 | `X-MAC-JAPANESE` | Localization | |
| 0x001CDA4F | `MAC-JAPANESE` | Localization | |
| 0x001CDA5C | `MACJAPANESE` | Localization | |
| 0x001CDA88 | `X-MAC-CHINESETRAD` | Localization | |
| 0x001CDA9A | `MAC-CHINESETRAD` | Localization | |
| 0x001CDAB7 | `X-MAC-CHINESESIMP` | Localization | |
| 0x001CDAC9 | `MAC-CHINESESIMP` | Localization | |
| 0x001CDAE9 | `X-MAC-KOREAN` | Localization | |
| 0x001CDAF6 | `MAC-KOREAN` | Localization | |

### Filesystem Path

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168834 | `iPod_Control\Device` | Filesystem Path | |
| 0x00168848 | `iPod_Control` | Filesystem Path | |
| 0x00168858 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x00168FC0 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x00169058 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x00169088 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001690E4 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x0016912C | `iPod_Control\Music\` | Filesystem Path | |

### Assertion

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00017740 | `*** assertion failed: %s, file %s, line %d` | Assertion | |
| 0x00314664 | `%s(%d): OpenSSL internal error, assertion failed: %s` | Assertion | |

---

## 3. Known User-Facing Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001688F0 | `KeyRepeatTimer` | Known | UI element |
| 0x0016F9FB | `k dine vCards til mappen Contacts p` | Known | UI element |
| 0x0016FB91 | `kke vCard-arkiverne til mappen "Contacts". Arkiverne bl...` | Known | UI element |
| 0x0016FC48 | `Alarmer` | Known | UI element |
| 0x00170474 | `Nulstil hovedmenu` | Known | Menu item |
| 0x00170614 | `Shuffle` | Known | UI element |
| 0x0017061C | `Hovedmenu` | Known | Menu item |
| 0x001707A4 | `Menuer` | Known | Menu item |
| 0x001724B8 | `Extras` | Known | UI element |
| 0x00172FA0 | `ffnen Sie ihn durch Doppelklicken in das iPod Symbol un...` | Known | UI element |
| 0x00173127 | `nnen sich hier Textdateien anzeigen lassen, indem Sie d...` | Known | UI element |
| 0x0017327F | `hlen" den Alarm beenden` | Known | UI element |
| 0x001753D0 | `Calendario` | Known | UI element |
| 0x001753DC | `Calendarios` | Known | UI element |
| 0x00175B28 | `El iPod puede almacenar contactos y eventos de calendar...` | Known | UI element |
| 0x00175D0A | `gido y arrastrar los archivos vCard a la carpeta Contac...` | Known | UI element |
| 0x00175E9D | `n de usar el iPod como disco y hacer doble clic en el i...` | Known | UI element |
| 0x00175FC0 | `Alarmas` | Known | UI element |
| 0x001760AD | `gido y arrastrar los archivos de texto a la carpeta Not...` | Known | UI element |
| 0x001761D8 | `Alarma` | Known | UI element |
| 0x001767B8 | `Reloj con alarma` | Known | UI element |
| 0x00176964 | `Contraste` | Known | UI element |
| 0x00176AD4 | `Hora alarma` | Known | UI element |
| 0x00178C4A | ` vCardit iPodin Contacts-kansioon. Lis` | Known | UI element |
| 0x00178DC2 | ` vCardit Contacts-kansioon. T` | Known | UI element |
| 0x00178F5A | ` tekstitiedostot iPodin Notes-kansioon. Lis` | Known | UI element |
| 0x0017BC10 | `Contacts` | Known | UI element |
| 0x0017C3DC | `Votre iPod peut stocker des contacts et des ` | Known | UI element |
| 0x0017C478 | `lectionnez Appareils > Ajouter un appareil. Puis choisi...` | Known | UI element |
| 0x0017C510 | `adresses, Microsoft Entourage ou Palm Desktop et export...` | Known | UI element |
| 0x0017C688 | ` mille contacts en plus de votre musique. Les applicati...` | Known | UI element |
| 0x0017C8B4 | `Alarmes` | Known | UI element |
| 0x0017C984 | `utilisation du disque, puis glissez ces fichiers dans l...` | Known | UI element |
| 0x0017CA80 | `Chargement des notes.` | Known | UI element |
| 0x0017CABC | `Alarme` | Known | UI element |
| 0x0017CAC4 | `Chargement des contacts.` | Known | UI element |
| 0x0017D113 | `init. menu p.` | Known | Menu item |
| 0x0017D2E4 | `Menu principal` | Known | Menu item |
| 0x0017D3B4 | `H. alarme` | Known | UI element |
| 0x0017D6B4 | `Menu princ.` | Known | Menu item |
| 0x0017EC7C | `Calendari` | Known | UI element |
| 0x0017F3B5 | ` archiviare contatti ed eventi di calendari. Se utilizz...` | Known | UI element |
| 0x0017FF4C | `Ripr. Menu Princ.` | Known | Menu item |
| 0x00180060 | `Contrasto` | Known | UI element |
| 0x00185D77 | ` Contacts ` | Known | UI element |
| 0x00185F59 | ` "Contacts" ` | Known | UI element |
| 0x0018612C | ` Notes ` | Known | UI element |
| 0x001889C4 | `Shuffle nummers` | Known | UI element |
| 0x00189280 | `De iPod biedt ruimte voor maar liefst duizend adressen ...` | Known | UI element |
| 0x00189544 | `Om tekstbestanden te bekijken, stelt u de iPod in als h...` | Known | UI element |
| 0x00189D2C | `Herstel menu` | Known | Menu item |
| 0x00189D58 | `Shuffle nrs.` | Known | UI element |
| 0x00189E50 | `Contrast` | Known | UI element |
| 0x00189EE8 | `Hoofdmenu` | Known | Menu item |
| 0x0018A070 | `Menu's` | Known | Menu item |
| 0x0018C268 | ` iPod-symbolet, og flytt vCard-filene inn i Contacts-ma...` | Known | UI element |
| 0x0018CD98 | `Alarmtidspunkt` | Known | UI element |
| 0x0018F3EB | `refter drar du in vCard-filerna i mappen "Contacts" i i...` | Known | UI element |
| 0x0018F597 | ` skrivbordet och drar in vCard-filerna i mappen "Contac...` | Known | UI element |
| 0x0018F759 | `ge och drar sedan in textfilerna i mappen "Notes" p` | Known | UI element |
| 0x0019011C | `Alarmtid` | Known | UI element |
| 0x001A3740 | `Now Playing` | Known | UI element |
| 0x001A379C | `Calendar` | Known | UI element |
| 0x001A37A8 | `Calendars` | Known | UI element |
| 0x001A37BC | `Backlight` | Known | UI element |
| 0x001A3814 | `Shuffle Songs` | Known | UI element |
| 0x001A4648 | `Alarms` | Known | UI element |
| 0x001A47D0 | `Notes loading.` | Known | UI element |
| 0x001A483C | `Contacts loading.` | Known | UI element |
| 0x001A4D4C | `Sleep Timer` | Known | UI element |
| 0x001A4D58 | `Alarm Clock` | Known | UI element |
| 0x001A4DFC | `Reset Main Menu` | Known | Menu item |
| 0x001A4EE0 | `Reset All Settings` | Known | User setting |
| 0x001A4F8C | `Backlight Timer` | Known | UI element |
| 0x001A4F9C | `Repeat` | Known | UI element |
| 0x001A4FAC | `Main Menu` | Known | Menu item |
| 0x001A5054 | `Settings` | Known | User setting |
| 0x001A5070 | `Alarm Time` | Known | UI element |
| 0x001A51A8 | `Reset All` | Known | UI element |
| 0x001CD5A5 | `Contacts\` | Known | UI element |
| 0x001CD5B9 | `Calendars\` | Known | UI element |
| 0x001CD5CA | `Notes\` | Known | UI element |
| 0x001CD610 | `vcalendar` | Known | UI element |
| 0x001CD6C4 | `dalarm` | Known | UI element |
| 0x001CD83C | `valarm` | Known | UI element |
| 0x001CE62C | `To check song links, set the preference NotesOnly to tr...` | Known | UI element |
| 0x001CE778 | `The NotesOnly pref can only be set globally in the Pref...` | Known | UI element |
| 0x001CE854 | `Warning: Preferences file must be in the Notes folder, ...` | Known | UI element |
| 0x001CEC60 | `Memory full. %d notes loaded, some notes not loaded. (3...` | Known | UI element |
| 0x003194B6 | `Illegal instruction` | Known | UI element |
| 0x003194E4 | `Illegal address` | Known | UI element |
| 0x003195EA | `NotesOnly` | Known | UI element |

---

## 4. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168414 | `RtcTaskClass` | Known | RTOS task thread |
| 0x00168424 | `TimerTaskClass` | Known | RTOS task thread |
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
| 0x00168950 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x0016D770 | `LoadDataTasks` | Known | RTOS task thread |
| 0x0016DAAB | `5RunTestsTask` | Known | RTOS task thread |
| 0x001BD794 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x00314A94 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x00314AA8 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x00314ABC | `SBP2CommandTask` | Known | RTOS task thread |

---

## 5. Audio/Codec

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0010B9F8 | `RIFFWAVEfmt data` | Known | PCM audio format |
| 0x0015B39C | `RIFFWAVEfmt dataD` | Known | PCM audio format |
| 0x00168960 | `pcmWrite.wav` | Known | PCM audio format |
| 0x0016FE30 | `Audible-softwaren i dette produkt bruges i henhold til ...` | Known | Audible audiobook format |
| 0x0016FE90 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x0016FF82 | `.net codec i dette produkt bruges i henhold til en lice...` | Known | Audio system |
| 0x0017002C | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunh...` | Known | Audio system |
| 0x001732B0 | `Die Audible Software in diesem Produkt wird in Lizenz v...` | Known | Audible audiobook format |
| 0x00173309 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x001733F9 | `.net Codec in diesem Produkt wird mit der Lizenz der Vo...` | Known | Audio system |
| 0x001734BF | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und ...` | Known | Audio system |
| 0x001761F4 | `El software Audible incluido en este producto se usa ba...` | Known | Audible audiobook format |
| 0x0017624F | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x001763F1 | `n de audio MPEG Layer-3 utilizada bajo licencia de Frau...` | Known | Audio system |
| 0x00179072 | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright...` | Known | Audible audiobook format |
| 0x001790AC | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x0017917E | `.net codec t` | Known | Audio system |
| 0x00179214 | `MPEG Layer-3 -` | Known | Audio system |
| 0x00179226 | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | Audio system |
| 0x0017CAE0 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x0017CB2A | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x0017CB3F | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x0017CBF0 | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x0017CCC4 | `La technologie de codage audio MPEG Layer 3 est utilis` | Known | Audio system |
| 0x0017CCFC | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | Audio system |
| 0x0017F9D4 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x0017F9FD | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0017FA2C | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x0017FA9E | ` utilizzato su licenza da VoiceAge Corporation. Il code...` | Known | Audio system |
| 0x0017FB74 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da...` | Known | Audio system |
| 0x00182E85 | ` Audible ` | Known | Audible audiobook format |
| 0x00182EA6 | `Audible ` | Known | Audible audiobook format |
| 0x00182EFF | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x001830B4 | `MPEG Layer-3 ` | Known | Audio system |
| 0x00183100 | `Fraunhofer IIS ` | Known | Audio system |
| 0x00186282 | ` Audible` | Known | Audible audiobook format |
| 0x001862C6 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x001863B6 | `.net codec` | Known | Audio system |
| 0x00186477 | ` Fraunhofer IIS` | Known | Audio system |
| 0x001896F4 | `De Audible-software in dit product wordt gebruikt in li...` | Known | Audible audiobook format |
| 0x0018974B | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x0018983C | `.net-codec in dit product wordt gebruikt in licentie va...` | Known | Audio system |
| 0x001898D8 | `Technologie voor codering van MPEG Layer-3-audio in lic...` | Known | Audio system |
| 0x0018C558 | `Audible-programvaren i dette produktet brukes under lis...` | Known | Audible audiobook format |
| 0x0018C5AC | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0018C728 | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fr...` | Known | Audio system |
| 0x0018F880 | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x0018F8AF | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x0018F8C6 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x0018FA60 | `Ljudkodningstekniken MPEG Layer-3 ` | Known | Audio system |
| 0x0018FA96 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | Audio system |
| 0x00192A4A | ` Fraunhofer IIS ` | Known | Audio system |
| 0x001A4850 | `The Audible software in this product is used under lice...` | Known | Audible audiobook format |
| 0x001A4989 | `.net codec in this product is used under license from V...` | Known | Audio system |
| 0x001A4A1C | `MPEG Layer-3 audio coding technology licensed from Frau...` | Known | Audio system |
| 0x001CE48D | `&Aacute` | Known | AAC codec |
| 0x001CE553 | `&aacute` | Known | AAC codec |
| 0x001CE943 | `Boot time (MP3PlayerExampleApp constructor)` | Known | MP3 codec |
| 0x0030EDA1 | ``0aLaCfDf` | Known | Apple Lossless codec |
| 0x003140EC | `Audible` | Known | Audible audiobook format |
| 0x003141C4 | `AudioCodecs` | Known | Audio system |
| 0x00314354 | `mp4_aacdec_sync` | Known | AAC codec |
| 0x00314364 | `mp3dec_sync` | Known | MP3 codec |
| 0x003162A7 | `msCodeCom` | Known | Audio system |
| 0x00317027 | `aaControls` | Known | AAC codec |

---

## 6. Storage/Hardware

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00003A00 | `!ATAsoso` | Known | ATA/disk interface |
| 0x00003A28 | `!ATAdpua` | Known | ATA/disk interface |
| 0x0004CF24 | `atad8@-` | Known | ATA/disk interface |
| 0x000932B8 | `atadmrts\|@-` | Known | ATA/disk interface |
| 0x000B32B0 | `atadmhdp` | Known | ATA/disk interface |
| 0x000B4A5C | `atadmhbddbhmmhsd>@-` | Known | ATA/disk interface |
| 0x000B4E74 | `atadmhpo` | Known | ATA/disk interface |
| 0x000E1824 | `nutiatad` | Known | ATA/disk interface |
| 0x001682EC | `data abort` | Known | ATA/disk interface |
| 0x001684B8 | `FirewireHandler` | Known | FireWire interface |
| 0x001685FD | `diskmode` | Known | Hardware interface |
| 0x00168606 | `diskscan` | Known | Hardware interface |
| 0x001687D0 | `diskModeImageRev` | Known | Hardware interface |
| 0x00168890 | `FirewireGuid` | Known | FireWire interface |
| 0x0016F0D0 | `Spiller nu` | Known | Hardware interface |
| 0x0016F170 | `Spillelister` | Known | Hardware interface |
| 0x0016F18C | `Genoptag spil` | Known | Hardware interface |
| 0x0016F630 | `Slet spilleliste` | Known | Hardware interface |
| 0x0016F644 | `Arkiver spilleliste` | Known | Hardware interface |
| 0x0016F6E4 | `Ny spilleliste %lu` | Known | Hardware interface |
| 0x0016F768 | `Harddisk` | Known | Hardware interface |
| 0x0016F9D6 | ` den kan bruges som harddisk, og tr` | Known | Hardware interface |
| 0x0016FB3B | `r du har tilsluttet iPod som disk, skal du dobbeltklikk...` | Known | Hardware interface |
| 0x0016FD02 | ` den kan bruges som disk og anbringe tekstarkiver i map...` | Known | Hardware interface |
| 0x00170293 | `je den til spillelisten On-The-Go. Hold knappen nede, n` | Known | Hardware interface |
| 0x001702CC | `r en spilleliste, kunstner eller et album er valgt for ...` | Known | Hardware interface |
| 0x00170309 | `je alle sangene til spillelisten On-The-Go. Spilleliste...` | Known | Hardware interface |
| 0x00170410 | `Nyt spil` | Known | Hardware interface |
| 0x00170538 | ` Spillelister` | Known | Hardware interface |
| 0x001708DC | `FireWire tilsluttet` | Known | FireWire interface |
| 0x0017246C | `Spiele` | Known | Hardware interface |
| 0x001724E0 | `Weiterspielen` | Known | Hardware interface |
| 0x00172D5F | `ffnen Sie das Adressbuch, Microsoft Entourage oder Palm...` | Known | FireWire interface |
| 0x0017386C | `Neues Spiel` | Known | Hardware interface |
| 0x00173D5A | `ber FireWire verbunden` | Known | FireWire interface |
| 0x00176CC8 | `FireWire conectado` | Known | FireWire interface |
| 0x00178868 | `Diskanttivahv.` | Known | Hardware interface |
| 0x00178878 | `Diskanttiheikenn.` | Known | Hardware interface |
| 0x0017893C | `Ladataan` | Known | ATA/disk interface |
| 0x00179008 | `Muistiinpanoja ladataan.` | Known | ATA/disk interface |
| 0x00179044 | `Yhteystietoja ladataan.` | Known | ATA/disk interface |
| 0x00179B18 | `FireWire liitetty` | Known | FireWire interface |
| 0x0017C585 | `utiliser comme disque FireWire. Puis glissez les vCards...` | Known | FireWire interface |
| 0x0017D5DC | `FireWire Connect` | Known | FireWire interface |
| 0x0017FE6C | `Data & Ora` | Known | ATA/disk interface |
| 0x0017FF10 | `Imposta Data & Ora` | Known | ATA/disk interface |
| 0x001803C0 | `FireWire Connesso` | Known | FireWire interface |
| 0x00183BC0 | `FireWire ` | Known | FireWire interface |
| 0x00189064 | `Op de iPod kunt u adres- en agendagegevens opslaan. Als...` | Known | FireWire interface |
| 0x0018A1A4 | `FireWire aangesloten` | Known | FireWire interface |
| 0x0018B7AC | `Spilles n` | Known | Hardware interface |
| 0x0018B860 | `Fortsett spill` | Known | Hardware interface |
| 0x0018BCB4 | `Mer diskant` | Known | Hardware interface |
| 0x0018BCC0 | `Mindre diskant` | Known | Hardware interface |
| 0x0018BD04 | `Slett spilleliste` | Known | Hardware interface |
| 0x0018BE58 | `Diskmodus` | Known | Hardware interface |
| 0x0018C044 | `pner du Adressebok, Microsoft Entourage eller Palm Desk...` | Known | Hardware interface |
| 0x0018C19C | `ringer i tillegg til musikken din. Microsoft Outlook, M...` | Known | Hardware interface |
| 0x0018C3DC | `Hvis du vil vise tekstfiler her, aktiverer du iPod for ...` | Known | Hardware interface |
| 0x0018C99C | ` legge den til i On-The-Go-spillelisten. Du kan legge t...` | Known | Hardware interface |
| 0x0018CAC4 | `Nytt spill` | Known | Hardware interface |
| 0x0018CF90 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x0018F3E0 | `rddisk. D` | Known | Hardware interface |
| 0x0018F53E | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x0018F55E | `rddisk. Sedan dubbelklickar du bara p` | Known | Hardware interface |
| 0x0018F750 | `rddiskl` | Known | Hardware interface |
| 0x00190308 | `FireWire anslutet` | Known | FireWire interface |
| 0x001A41D4 | `Disk Mode` | Known | Hardware interface |
| 0x001A42A8 | `Your iPod can store contacts and calendar events. If yo...` | Known | Hardware interface |
| 0x001A4470 | `Your iPod can store up to one thousand contacts right a...` | Known | Hardware interface |
| 0x001A46D0 | `To view text files here, enable iPod for disk use, then...` | Known | Hardware interface |
| 0x001A52BC | `FireWire Connected` | Known | FireWire interface |
| 0x001A52D0 | `No battery power remains. Please connect iPod to power.` | Known | Power management |
| 0x001A5308 | `Low Battery` | Known | Power management |
| 0x001BD82C | `TROMResourceDB - unknown header version! (Try regenerat...` | Known | ATA/disk interface |
| 0x001CDD94 | `Bad Data` | Known | ATA/disk interface |
| 0x001CE740 | `Bad meta data, name not a recognized preference. (27)` | Known | ATA/disk interface |
| 0x001CE7F8 | `Bad meta data, name not found. (23)` | Known | ATA/disk interface |
| 0x001CE81C | `Bad meta data, name termination quote not found. (24)` | Known | ATA/disk interface |
| 0x001CE8A4 | `Bad meta data, content not found. (25)` | Known | ATA/disk interface |
| 0x001CE8CC | `Bad meta data, content termination quote not found. (26...` | Known | ATA/disk interface |
| 0x001CE9F8 | `Bad meta data, encoding not found. (8)` | Known | ATA/disk interface |
| 0x001CEA20 | `Bad meta data, encoding termination quote not found. (9...` | Known | ATA/disk interface |
| 0x001CEA5C | `Bad meta data, unknown encoding. (17)` | Known | ATA/disk interface |
| 0x001CEC18 | `Bad data. (32)` | Known | ATA/disk interface |
| 0x001CEDAC | `USB MSC` | Known | USB interface |
| 0x00303BC3 | `ataTaza[aea;ajaaaVa)b'b+b+dMd[d]dtdvdrdsd}dudfd` | Known | ATA/disk interface |
| 0x003087AD | `aGa>a(a'aJa?a<a,a4a=aBaDasawaXaYaZakataoaeaqa_a]aSaua` | Known | ATA/disk interface |
| 0x003141A8 | `MaximumDataRate` | Known | ATA/disk interface |
| 0x00314214 | `FireWire` | Known | FireWire interface |
| 0x00314284 | `FireWireVersion` | Known | FireWire interface |
| 0x00314454 | `MEMDISK` | Known | Hardware interface |
| 0x003144EA | `ex_data` | Known | ATA/disk interface |
| 0x003146BC | `c:\buildtools\MWSF2\Q22Firmware.proj\projectfiles\sandb...` | Known | ATA/disk interface |
| 0x003156FB | `RSA Data Security, Inc.` | Known | ATA/disk interface |
| 0x00315718 | `RSA Data Security, Inc. PKCS` | Known | ATA/disk interface |
| 0x00315879 | `pkcs7-data` | Known | ATA/disk interface |
| 0x00315884 | `pkcs7-signedData` | Known | ATA/disk interface |
| 0x00315895 | `pkcs7-envelopedData` | Known | ATA/disk interface |
| 0x003158A9 | `pkcs7-signedAndEnvelopedData` | Known | ATA/disk interface |
| 0x003158C6 | `pkcs7-digestData` | Known | ATA/disk interface |
| 0x003158D7 | `pkcs7-encryptedData` | Known | ATA/disk interface |
| 0x00315B30 | `nsDataType` | Known | ATA/disk interface |
| 0x00315B3B | `Netscape Data Type` | Known | ATA/disk interface |
| 0x003168D5 | `id-smime-ct-authData` | Known | ATA/disk interface |
| 0x00316942 | `id-smime-ct-DVCSRequestData` | Known | ATA/disk interface |
| 0x0031695E | `id-smime-ct-DVCSResponseData` | Known | ATA/disk interface |
| 0x00317376 | `id-cmc-dataReturn` | Known | ATA/disk interface |
| 0x003174C9 | `id-on-personalData` | Known | ATA/disk interface |
| 0x003175CA | `id-cct-PKIData` | Known | ATA/disk interface |
| 0x00317D69 | `qualityLabelledData` | Known | ATA/disk interface |
| 0x00318144 | `setct-PANData` | Known | ATA/disk interface |
| 0x0031816F | `setct-OIData` | Known | ATA/disk interface |
| 0x00318185 | `setct-PIData` | Known | ATA/disk interface |
| 0x00318192 | `setct-PIDataUnsigned` | Known | ATA/disk interface |
| 0x0031820D | `setct-PInitResData` | Known | ATA/disk interface |
| 0x0031822D | `setct-PResData` | Known | ATA/disk interface |
| 0x00318283 | `setct-CapTokenData` | Known | ATA/disk interface |
| 0x003182D1 | `setct-AuthRevResData` | Known | ATA/disk interface |
| 0x0031831B | `setct-CapResData` | Known | ATA/disk interface |
| 0x00318353 | `setct-CapRevResData` | Known | ATA/disk interface |
| 0x0031838A | `setct-CredResData` | Known | ATA/disk interface |
| 0x003183C5 | `setct-CredRevResData` | Known | ATA/disk interface |
| 0x003183DA | `setct-PCertReqData` | Known | ATA/disk interface |
| 0x003183FF | `setct-BatchAdminReqData` | Known | ATA/disk interface |
| 0x00318417 | `setct-BatchAdminResData` | Known | ATA/disk interface |
| 0x0031846F | `setct-CertReqData` | Known | ATA/disk interface |
| 0x00318492 | `setct-CertResData` | Known | ATA/disk interface |
| 0x0031881E | `setCext-merchData` | Known | ATA/disk interface |
| 0x003188A9 | `setCext-Track2Data` | Known | ATA/disk interface |
| 0x00318A97 | `set-brand-IATA-ATA` | Known | ATA/disk interface |
| 0x0031DB03 | `<key>FireWireGUID</key>` | Known | FireWire interface |

---

## 7. Error Messages

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00017740 | `*** assertion failed: %s, file %s, line %d` | Known | Error/assertion message |
| 0x0001DD6C | `Invalid Operation` | Known | Error/assertion message |
| 0x001CE690 | `Too many errors, further errors discarded. (31)` | Known | Error/assertion message |
| 0x001CE6E0 | `%s Error in file %s.` | Known | Error/assertion message |
| 0x001CEAA4 | `Cannot link to a .link file. (29)` | Known | Error/assertion message |
| 0x001CF196 | `tInvalid partition table. Setup cannot continue.` | Known | Error/assertion message |
| 0x001CF1C7 | `Error loading operating system. Setup cannot continue.` | Known | Error/assertion message |
| 0x00314664 | `%s(%d): OpenSSL internal error, assertion failed: %s` | Known | Error/assertion message |
| 0x003152F0 | `error:%08lX:%s:%s:%s` | Known | Error/assertion message |
| 0x00315594 | `internal error: list index %ld out of range` | Known | Error/assertion message |

---

## 8. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016D27C | ` !"#$%&'()*+,-./0123456789:;<=>?@abcdefghijklmnopqrstuv...` | Known | Filesystem path |
| 0x0016D37C | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x0016D57C | `url;type=work:apple.com/support/ipod` | Known | Filesystem path |
| 0x0016FDA8 | `Point: %d (%d/%d)` | Known | Filesystem path |
| 0x0016FF03 | `rende VoiceAge Corporation i USA og/eller andre lande o...` | Known | Filesystem path |
| 0x001725D4 | `USA/Hawaii (NZ)` | Known | Filesystem path |
| 0x001725E4 | `USA/Hawaii (SZ)` | Known | Filesystem path |
| 0x001725F4 | `USA/Alaska (NZ)` | Known | Filesystem path |
| 0x00172604 | `USA/Alaska (SZ)` | Known | Filesystem path |
| 0x00172614 | `USA/Pazifik (NZ)` | Known | Filesystem path |
| 0x00172628 | `USA/Pazifik (SZ)` | Known | Filesystem path |
| 0x0017263C | `USA/Rockies (NZ)` | Known | Filesystem path |
| 0x00172650 | `USA/Rockies (SZ)` | Known | Filesystem path |
| 0x00172664 | `USA/Zentral (NZ)` | Known | Filesystem path |
| 0x00172678 | `USA/Zentral (SZ)` | Known | Filesystem path |
| 0x0017268C | `USA/Ost (NZ)` | Known | Filesystem path |
| 0x0017269C | `USA/Ost (SZ)` | Known | Filesystem path |
| 0x001729A0 | `Vorn./Nachn.` | Known | Filesystem path |
| 0x001729B0 | `Nachn./Vorn.` | Known | Filesystem path |
| 0x00173220 | `Punkte: %d (%d/%d)` | Known | Filesystem path |
| 0x00173343 | ` ist entweder eine eingetragene Marke oder eine Marke d...` | Known | Filesystem path |
| 0x0017616C | `Result.: %d (%d/%d)` | Known | Filesystem path |
| 0x0017628F | ` es una marca registrada o una marca comercial de Voice...` | Known | Filesystem path |
| 0x00176C14 | `Fecha/hora` | Known | Filesystem path |
| 0x00178934 | `%d / %d` | Known | Filesystem path |
| 0x00178FD4 | `Tulos: %d (%d/%d)` | Known | Filesystem path |
| 0x0017911D | `ity tavaramerkki Yhdysvalloissa ja/tai muissa maissa, j...` | Known | Filesystem path |
| 0x0017CA4B | `sult. : %d (%d/%d)` | Known | Filesystem path |
| 0x0017CBC2 | `tats-Unis et/ou dans d` | Known | Filesystem path |
| 0x0017D0CF | `gler date/heure` | Known | Filesystem path |
| 0x0017F944 | `Punti: %d (%d/%d)` | Known | Filesystem path |
| 0x001819AC | `%b/%-d %-I:%M %2p` | Known | Filesystem path |
| 0x001819C0 | `%-m/%-d` | Known | Filesystem path |
| 0x001819E4 | `%y/%-m/%d` | Known | Filesystem path |
| 0x001819F0 | `%Y/%b/%-d` | Known | Filesystem path |
| 0x00182DB4 | ` %d (%d/%d)` | Known | Filesystem path |
| 0x00185220 | `%Y/%B/%d %A  %I:%M:%S %p` | Known | Filesystem path |
| 0x0018523C | `%Y/%B/%d` | Known | Filesystem path |
| 0x00185254 | `%-m/%-d %-I:%M %p` | Known | Filesystem path |
| 0x00185288 | `%Y/%-m/%d` | Known | Filesystem path |
| 0x001861D2 | `: %d (%d/%d)` | Known | Filesystem path |
| 0x00189678 | `Score: %d (%d/%d)` | Known | Filesystem path |
| 0x00189783 | ` is een gedeponeerd handelsmerk of een handelsmerk van ...` | Known | Filesystem path |
| 0x00189CF0 | `Stel datum/tijd in` | Known | Filesystem path |
| 0x0018A0E4 | `Datum/tijd` | Known | Filesystem path |
| 0x0018C4C4 | `Poeng: %d (%d/%d)` | Known | Filesystem path |
| 0x0018C5EB | ` er enten et registrert varemerke eller et varemerke fo...` | Known | Filesystem path |
| 0x0018EA84 | `%-d/%-m` | Known | Filesystem path |
| 0x0018F7F8 | `ng: %d (%d/%d)` | Known | Filesystem path |
| 0x0018F94C | `r VoiceAge Corporation i USA och/eller andra l` | Known | Filesystem path |
| 0x0018FDC4 | `Datum/tid` | Known | Filesystem path |
| 0x0018FE60 | `ll in datum/tid` | Known | Filesystem path |
| 0x00192801 | `%d (%d/%d)` | Known | Filesystem path |
| 0x001A368C | `%-m/%d/%y` | Known | Filesystem path |
| 0x001A48DB | ` is either registered trademark or trademark of VoiceAg...` | Known | Filesystem path |
| 0x001CDB19 | `file://` | Known | Filesystem path |
| 0x001CDB21 | `image://` | Known | Filesystem path |
| 0x001CDD00 | `</TITLE>` | Known | Filesystem path |
| 0x001CDD10 | `</BODY>` | Known | Filesystem path |
| 0x001CDD3E | `</ROT13>` | Known | Filesystem path |
| 0x001CEAE8 | `Bad link, no matching </a> for anchor tag. (20)` | Known | Filesystem path |
| 0x001E6A17 | `W/}lE>q` | Known | Filesystem path |
| 0x00215659 | `H."0*Bx/` | Known | Filesystem path |
| 0x0021CFB6 | `U/~RERT` | Known | Filesystem path |
| 0x00221D1E | `TUOPT/\|` | Known | Filesystem path |
| 0x00228FB3 | `HuGZp/$j` | Known | Filesystem path |
| 0x0022F4C3 | `(bJ)b"(b2""z/` | Known | Filesystem path |
| 0x00231C33 | `JUAPDD(/` | Known | Filesystem path |
| 0x00237F42 | `/B\|$BD'` | Known | Filesystem path |
| 0x00238EBF | `$Bd$BT/` | Known | Filesystem path |
| 0x0023F197 | `/" +J\|!` | Known | Filesystem path |
| 0x00246016 | `Fb""")/` | Known | Filesystem path |
| 0x002471CD | `/RyO(UIH` | Known | Filesystem path |
| 0x0024827D | `~$Bh'"~$Bz/` | Known | Filesystem path |
| 0x0024BAE7 | `$B +BZ/` | Known | Filesystem path |
| 0x00253505 | `0c(HBP/` | Known | Filesystem path |
| 0x00257A9B | `$B~("\|/` | Known | Filesystem path |
| 0x0026E805 | `T/DDDDD` | Known | Filesystem path |
| 0x0026EA6F | `"~UeB /` | Known | Filesystem path |
| 0x00271521 | `$B((B /` | Known | Filesystem path |
| 0x00279604 | ` "\|$B~/` | Known | Filesystem path |
| 0x0027C780 | `@$B\|$"(/` | Known | Filesystem path |
| 0x0027D880 | `)"8/B""` | Known | Filesystem path |
| 0x0027DFC8 | `r4c6 bN/` | Known | Filesystem path |
| 0x002833C9 | `RDT%B(/` | Known | Filesystem path |
| 0x00284555 | `RBHUE\|/` | Known | Filesystem path |
| 0x0028BBD5 | `]B""B</` | Known | Filesystem path |
| 0x0028F452 | `,B\|RED/` | Known | Filesystem path |
| 0x00294ED5 | `$BT). /` | Known | Filesystem path |
| 0x002960D5 | `#"TUB(/` | Known | Filesystem path |
| 0x002A7949 | `/" %BD"` | Known | Filesystem path |
| 0x002AE26C | `ODD""(/` | Known | Filesystem path |
| 0x002AF4A7 | `B"$R%"B$" /` | Known | Filesystem path |
| 0x002AFE14 | `bG\|jG\|/` | Known | Filesystem path |
| 0x002B19FE | `$E$$BR/` | Known | Filesystem path |
| 0x002B1A9F | `dRB~RA$/` | Known | Filesystem path |
| 0x002B23F0 | `TT&T%B(/` | Known | Filesystem path |
| 0x002C10E4 | `)'>$B8/` | Known | Filesystem path |
| 0x002C36F2 | `$B\|%EV/` | Known | Filesystem path |
| 0x002CAC46 | `BDU!BJ ""/` | Known | Filesystem path |
| 0x002CBBAA | `Z-bD("(%B>/` | Known | Filesystem path |
| 0x002D42F9 | `@(/ Q\|f` | Known | Filesystem path |
| 0x002FB50D | `% %!%"%#%$%%%&%'%(%)%*%+%,%-%.%/%0%1%2%3%4%5%6%7%8%9%:%...` | Known | Filesystem path |
| 0x002FBCED | `qWlIl/Ymg*` | Known | Filesystem path |
| 0x002FD55B | `X)W,W*W3W9W.W/W\W;WBWiW` | Known | Filesystem path |
| 0x002FDC29 | `n/o6oKoto*o` | Known | Filesystem path |
| 0x002FE1C1 | `q/q1qsq\qhqEqrqJqxqzq` | Known | Filesystem path |
| 0x002FF449 | `b6bKbNb/e` | Known | Filesystem path |
| 0x002FF5FF | `V3W0W(W-W,W/W)W` | Known | Filesystem path |
| 0x002FF6D3 | `NMOOOGOWO^O4O[OUO0OPOQO=O:O8OCOTO<OFOcO\O`O/ONO6OYO]OHO...` | Known | Filesystem path |
| 0x002FFD31 | `e%f-f f'f/f` | Known | Filesystem path |
| 0x00300033 | `bNc>c/cUcBcFcOcIc:cPc=c*c+c(cMcLcHeIe` | Known | Filesystem path |
| 0x003002E5 | `fFUjUfUDU^UaUCUJU1UVUOUUU/UdU8U.U\U,UcU3UAUWU` | Known | Filesystem path |
| 0x0030054B | `\|F}C}q}.}9}<}@}0}3}D}/}B}2}1}=` | Known | Filesystem path |
| 0x003006D3 | `W/X*X4X$X0X1X!X` | Known | Filesystem path |
| 0x003007E9 | `k.l/l,l/n8nTn!n2ngnJn n%n#n` | Known | Filesystem path |
| 0x0030084B | `r6s%s4s)s:t*t3t"t%t5t6t4t/t` | Known | Filesystem path |
| 0x003010C9 | `S.V;V9V2V?V4V)VSVNVWVtV6V/V0V` | Known | Filesystem path |
| 0x0030170F | `\|.~>~F~7~2~C~+~=~1~E~A~4~9~H~5~?~/~D` | Known | Filesystem path |
| 0x003026FF | `P P'P5P/P1P` | Known | Filesystem path |
| 0x00302851 | `h5h+h-h/hNhDh4h` | Known | Filesystem path |
| 0x00302865 | `h&h(h.hMh:h%h h,k/k-k1k4kmk` | Known | Filesystem path |
| 0x00302971 | `w"w'w#x,x"x5x/x(x.x+x!x)x3x*x1xTy[yOy\ySyRyQy` | Known | Filesystem path |
| 0x00302DFB | `v9w/w-w1w2w4w3w=w%w;w5wHxRxIxMxJxLx&xExPxdygyiyjycykyay` | Known | Filesystem path |
| 0x00302E57 | `{1{+{-{/{2{8{` | Known | Filesystem path |
| 0x003032EB | `t/uoulu` | Known | Filesystem path |
| 0x00303651 | ``IaJa+aEa6a2a.aFa/aOa)a@a bh` | Known | Filesystem path |
| 0x00303697 | `d&d0d(dAd5d/d` | Known | Filesystem path |
| 0x003036FF | `hxi4iii@ioiDiviXiAitiLi;iKi7i\iOiQi2iRi/i{i<iFkEkCkBkHk...` | Known | Filesystem path |
| 0x00303D69 | `y+zJz0z/z(z&z` | Known | Filesystem path |
| 0x0030464D | `u/v-v1v=v3v<v5v2v0v` | Known | Filesystem path |
| 0x00304943 | `X-[%[2[#[,['[&[/[.[{[` | Known | Filesystem path |
| 0x00305271 | `j<p5p/p7p4p1pBp8p?p:p9p@p;p3pAp` | Known | Filesystem path |
| 0x00305473 | `w-y1y/yT\|S\|` | Known | Filesystem path |
| 0x0030683F | `%#%3%+%;%K% %/%(%7%?%` | Known | Filesystem path |
| 0x003069BE | `02*2+2,2-2.2/2@272B2C292:212>24222;2623252<2=2?282` | Known | Filesystem path |
| 0x00306EE9 | `\7_J_/`P`m`` | Known | Filesystem path |
| 0x003070C5 | `OHSIT>T/Z` | Known | Filesystem path |
| 0x0030713B | `i_l*mim/n` | Known | Filesystem path |
| 0x0030762D | `N,p]u/f` | Known | Filesystem path |
| 0x00308339 | `S#S/S1S3S8S@SFSES` | Known | Filesystem path |
| 0x00308439 | `q4V6V2V8VkVdV/VlVjV` | Known | Filesystem path |
| 0x003086E1 | `_)_-_8_A_H_L_N_/_Q_V_W_Y_a_m_s_w_` | Known | Filesystem path |
| 0x00308DA7 | `s4s/s)s%s>sNsOs` | Known | Filesystem path |
| 0x0030CBAB | `h>kLp/t` | Known | Filesystem path |
| 0x0030D235 | `o;v/}7~` | Known | Filesystem path |
| 0x0030DFF9 | `e1f/h\q6z` | Known | Filesystem path |
| 0x0030E645 | `UuX/c"dIfKfmh` | Known | Filesystem path |
| 0x00313FE8 | `%s<key>%s</key>` | Known | Filesystem path |
| 0x00313FF8 | `%s<%s/>` | Known | Filesystem path |
| 0x00314020 | `%s</dict>` | Known | Filesystem path |
| 0x0031403C | `%s<string>%s</string>` | Known | Filesystem path |
| 0x00314064 | `%s<integer>%s</integer>` | Known | Filesystem path |
| 0x00314090 | `%s<integer>%d</integer>` | Known | Filesystem path |
| 0x003143F4 | ` ,;=+[]*?<>\|":/\` | Known | Filesystem path |
| 0x00314405 | `\/:*?"<>\|` | Known | Filesystem path |
| 0x00315428 | `You need to read the OpenSSL FAQ, http://www.openssl.or...` | Known | Filesystem path |
| 0x003165F5 | `S/MIME Capabilities` | Known | Filesystem path |
| 0x0031DA64 | `<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1....` | Known | Filesystem path |
| 0x0031DAF2 | `</dict>` | Known | Filesystem path |
| 0x0031DAFA | `</plist>` | Known | Filesystem path |
| 0x0031DB1B | `<string>%08lX%08lX</string>` | Known | Filesystem path |
| 0x0031F2B2 | `/ 0 1 2 3 4 5 6 7 8 9 : ; < = > ? @ A B C D E F G H I J...` | Known | Filesystem path |
| 0x0031F493 | `! !!!"!#!$!%!&!'!(!)!*!+!,!-!.!/!0!1!2!3!4!5!6!7!8!9!:!...` | Known | Filesystem path |
| 0x003362DC | `/1f;{1Q` | Known | Filesystem path |
| 0x0034549F | `p/^s{eL` | Known | Filesystem path |
| 0x0034DF3C | `4/f$-M<` | Known | Filesystem path |
| 0x0035093A | `rm#OQ\|/Ho` | Known | Filesystem path |
| 0x00353858 | `/:/JtlW5` | Known | Filesystem path |
| 0x00354611 | `7/A\*2D` | Known | Filesystem path |
| 0x0035EA5C | `vf:/Yq6F` | Known | Filesystem path |
| 0x0036156A | `&b/EN56` | Known | Filesystem path |
| 0x0036E32D | `SPR./(o` | Known | Filesystem path |
| 0x0037228D | `E,[//H/` | Known | Filesystem path |
| 0x003801BC | `/,w>.s!` | Known | Filesystem path |
| 0x0038524B | `/wi-9Ak` | Known | Filesystem path |
| 0x00385D80 | `c*Y//2b"` | Known | Filesystem path |
| 0x003864E1 | `E,6(A9/` | Known | Filesystem path |
| 0x0038EA04 | `oRnN/BB` | Known | Filesystem path |
| 0x00390C7B | `/}NVsvYf` | Known | Filesystem path |
| 0x00393A02 | `HJD:/0i` | Known | Filesystem path |
| 0x00398F26 | `>/FEnj~` | Known | Filesystem path |
| 0x0039D7BA | `/W."4s@` | Known | Filesystem path |
| 0x003A051E | `*x/;yCZ` | Known | Filesystem path |
| 0x003A897F | `/V~48-J` | Known | Filesystem path |
| 0x003B351D | `?/5k*hF` | Known | Filesystem path |
| 0x003CA305 | `V(,/0O)` | Known | Filesystem path |
| 0x003DC1EF | `qy/n5Jr` | Known | Filesystem path |
| 0x003DF960 | `7/]hwF,` | Known | Filesystem path |
| 0x003E9383 | `Q/IFt-8x` | Known | Filesystem path |
| 0x003ED3DC | `cBvFSKNZ/` | Known | Filesystem path |
| 0x003FBDC9 | `>xX/o\|X*` | Known | Filesystem path |
| 0x00400ED9 | `Pk*X/.l` | Known | Filesystem path |
| 0x00406793 | `^"X(E/L` | Known | Filesystem path |
| 0x0040869E | `H^3 2/]` | Known | Filesystem path |
| 0x00409466 | `!S<-D]`/` | Known | Filesystem path |
| 0x00423329 | `6E[)*/p` | Known | Filesystem path |
| 0x00429B1B | `=8\|'/[y$` | Known | Filesystem path |
| 0x0042C1D1 | `)V.&).R/` | Known | Filesystem path |
| 0x00432308 | `Vc//pP<J` | Known | Filesystem path |
| 0x0043A272 | `/7$HTr*.` | Known | Filesystem path |
| 0x004488EB | `CVUyi'/` | Known | Filesystem path |
| 0x00449C97 | `eX$9*\|=/*` | Known | Filesystem path |
| 0x00449D59 | `/{%S`f1u` | Known | Filesystem path |

---

## 9. Binary Structure (for Ghidra/IDA)

| Parameter | Value |
|-----------|-------|
| **Architecture** | ARM (32-bit, little-endian) |
| **Processor** | ARM7TDMI (ARMv4T), dual-core |
| **Base Address** | 0x00000000 |
| **Entry Point** | 0x00000000 |
| **Endianness** | Little-Endian |
| **File Size** | 4,506,624 bytes |

