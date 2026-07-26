# iPod 3rd Generation - RetailOS 2.2.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 2.2.3 |
| **IPSW** | iPod_2.2.3.ipsw |
| **Device** | iPod 3rd Generation (2003, Dock Connector, Touch Buttons) |
| **Binary Size** | 4,561,920 bytes (4.35 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 4,561,920 bytes |
| **Total Strings (>=6)** | 9,755 |
| **Function Prologues** | 7,192 |
| **SoC** | PortalPlayer PP5002 |
| **Architecture** | ARM7TDMI (ARMv4T), dual-core |
| **Encrypted** | No |
| **Decryption Method** | Extract from IPSW (unencrypted) |
| **SHA-256** | `d4a95ed35add9f001058ca66ce515e535833cfb2e9de2c6c9377ba9ef56ea0ce` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00151224 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x00355E94 | ` <- Error:Internal Error` | Hidden | Undocumented UI |
| 0x0035F114 | `BTM Debug Zones %s (0x%08X)` | Hidden | Debug/Diagnostic |
| 0x00362884 | `Retail mode` | Hidden | Demo/Retail Mode |
| 0x00362894 | `Debug mode` | Hidden | Debug/Diagnostic |
| 0x0036425C | `SP_ERR_ASSERTION` | Hidden | Developer Tool |

---

## 2. Discovered Features

### Diagnostic

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0036285C | `Diag mode` | Diagnostic | |

### EQ Preset

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
| 0x0015CAE8 | `USA/Rockies (NZ)` | EQ Preset | |
| 0x0015CAFC | `USA/Rockies (SZ)` | EQ Preset | |
| 0x00160444 | `Latina` | EQ Preset | |
| 0x00179910 | `Latino` | EQ Preset | |
| 0x001AFF04 | `iPod_Control\Device\_short_deepsleep` | EQ Preset | |
| 0x001AFF2C | `iPod_Control\Device\_no_deepsleep` | EQ Preset | |
| 0x001C32A4 | `LATIN-1` | EQ Preset | |
| 0x001C32AC | `LATIN1` | EQ Preset | |
| 0x001F0EEB | `~ BR&B$"` | EQ Preset | |
| 0x00315A72 | `Secure Electronic Transactions` | EQ Preset | |
| 0x0035E9AB | `Switch to deep sleep mode (off)` | EQ Preset | |

### Localization

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001514EC | `ChargerManager` | Localization | |
| 0x00193B0C | `English` | Localization | |
| 0x00193B44 | `Italiano` | Localization | |
| 0x001AFB68 | `x-mac-japanese` | Localization | |
| 0x001C330F | `X-MAC-JAPANESE` | Localization | |
| 0x001C331E | `MAC-JAPANESE` | Localization | |
| 0x001C332B | `MACJAPANESE` | Localization | |
| 0x001C3357 | `X-MAC-CHINESETRAD` | Localization | |
| 0x001C3369 | `MAC-CHINESETRAD` | Localization | |
| 0x001C3386 | `X-MAC-CHINESESIMP` | Localization | |
| 0x001C3398 | `MAC-CHINESESIMP` | Localization | |
| 0x001C33B8 | `X-MAC-KOREAN` | Localization | |
| 0x001C33C5 | `MAC-KOREAN` | Localization | |

### Filesystem Path

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00151360 | `iPod_Control\Device` | Filesystem Path | |
| 0x00151374 | `iPod_Control` | Filesystem Path | |
| 0x00151384 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x001AFE3C | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001AFED4 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x001AFF50 | `iPod_Control\Device\_show_voltage` | Filesystem Path | |
| 0x001AFF74 | `iPod_Control\Device\_disable_cache` | Filesystem Path | |
| 0x001AFFDC | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x001C45D0 | `iPod_Control\Testing\` | Filesystem Path | |
| 0x001C45E8 | `iPod_Control\Testing\\TestLog.txt` | Filesystem Path | |
| 0x001C460C | `iPod_Control\Testing\\Tests.Lock` | Filesystem Path | |

### Assertion

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0001BB3C | `*** assertion failed: %s, file %s, line %d` | Assertion | |
| 0x0030F07C | `%s(%d): OpenSSL internal error, assertion failed: %s` | Assertion | |
| 0x0036425C | `SP_ERR_ASSERTION` | Assertion | |

---

## 3. Known User-Facing Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001511DC | `KeyRepeatTimer` | Known | UI element |
| 0x0015989D | `k dine vCards til mappen Contacts p` | Known | UI element |
| 0x00159A32 | `kke vCard-arkiverne til mappen "Contacts". Arkiverne bl...` | Known | UI element |
| 0x00159AE8 | `Alarmer` | Known | UI element |
| 0x0015A51C | `Nulstil menu` | Known | Menu item |
| 0x0015A6B4 | `Hovedmenu` | Known | Menu item |
| 0x0015AA18 | `Menuer` | Known | Menu item |
| 0x0015C9B0 | `Extras` | Known | UI element |
| 0x0015D2CA | `ffnen Sie ihn durch Doppelklicken in das iPod Symbol un...` | Known | UI element |
| 0x0015D4C4 | `ffnen Sie ihn durch Doppelklicken in das iPod Symbol un...` | Known | UI element |
| 0x0015D6BB | `nnen sich hier Textdateien anzeigen lassen, indem Sie d...` | Known | UI element |
| 0x0015DA1B | `hlen" den Alarm beenden` | Known | UI element |
| 0x0016000C | `Calendario` | Known | UI element |
| 0x00160018 | `Calendarios` | Known | UI element |
| 0x0016071A | `El iPod puede almacenar contactos y eventos de calendar...` | Known | UI element |
| 0x00160A9B | `n de usar el iPod como disco y hacer doble clic en el i...` | Known | UI element |
| 0x00160BBC | `Alarmas` | Known | UI element |
| 0x00160D21 | `gido y arrastrar los archivos de texto a la carpeta Not...` | Known | UI element |
| 0x00161078 | `Alarma` | Known | UI element |
| 0x0016163C | `Reloj con alarma` | Known | UI element |
| 0x00161828 | `Contraste` | Known | UI element |
| 0x00161974 | `Hora alarma` | Known | UI element |
| 0x00163F75 | ` vCardit iPodin Contacts-kansioon. Lis` | Known | UI element |
| 0x001640F0 | ` vCardit Contacts-kansioon. T` | Known | UI element |
| 0x00164316 | ` tekstitiedostot iPodin Notes-kansioon. Lis` | Known | UI element |
| 0x00167638 | `Contacts` | Known | UI element |
| 0x00167DB2 | `Votre iPod peut stocker des contacts et des ` | Known | UI element |
| 0x00168056 | ` mille contacts en plus de votre musique. Les applicati...` | Known | UI element |
| 0x0016817B | `iPod et de glisser ces vCards dans le dossier "Contacts...` | Known | UI element |
| 0x00168278 | `Alarmes` | Known | UI element |
| 0x0016838C | `Pour visualiser les fichiers au format texte, activez l...` | Known | UI element |
| 0x00168700 | `Chargement des notes.` | Known | UI element |
| 0x0016873C | `Alarme` | Known | UI element |
| 0x00168744 | `Chargement des contacts.` | Known | UI element |
| 0x00168DA3 | `init. menu p.` | Known | Menu item |
| 0x00168F78 | `Menu principal` | Known | Menu item |
| 0x0016903C | `H. alarme` | Known | UI element |
| 0x00169520 | `Menu princ.` | Known | Menu item |
| 0x0016ADBC | `Calendari` | Known | UI element |
| 0x0016C2BC | `Ripr. Menu Princ.` | Known | Menu item |
| 0x0016C3E4 | `Contrasto` | Known | UI element |
| 0x00172C0E | ` Contacts ` | Known | UI element |
| 0x00172DF2 | ` "Contacts" ` | Known | UI element |
| 0x00173051 | ` Notes ` | Known | UI element |
| 0x00175FE4 | `Shuffle nummers` | Known | UI element |
| 0x001768AA | `De iPod biedt ruimte voor maar liefst duizend adressen ...` | Known | UI element |
| 0x00176BD8 | `Om tekstbestanden te bekijken, stelt u de iPod in als h...` | Known | UI element |
| 0x001775C0 | `Herstel menu` | Known | Menu item |
| 0x001776F8 | `Contrast` | Known | UI element |
| 0x00177770 | `Shuffle` | Known | UI element |
| 0x00177778 | `Hoofdmenu` | Known | Menu item |
| 0x00177AD8 | `Menu's` | Known | Menu item |
| 0x00179F33 | ` iPod-symbolet, og flytt vCard-filene inn i Contacts-ma...` | Known | UI element |
| 0x0017AC9C | `Alarmtidspunkt` | Known | UI element |
| 0x0017D722 | `refter drar du in vCard-filerna i mappen "Contacts" i i...` | Known | UI element |
| 0x0017D8D1 | ` skrivbordet och drar in vCard-filerna i mappen "Contac...` | Known | UI element |
| 0x0017DB01 | `ge och drar sedan in textfilerna i mappen "Notes" p` | Known | UI element |
| 0x0017E688 | `Alarmtid` | Known | UI element |
| 0x001939E0 | `Now Playing` | Known | UI element |
| 0x00193A58 | `Calendar` | Known | UI element |
| 0x00193A64 | `Calendars` | Known | UI element |
| 0x00193A78 | `Backlight` | Known | UI element |
| 0x00193AD0 | `Shuffle Songs` | Known | UI element |
| 0x001948D0 | `Alarms` | Known | UI element |
| 0x00194C9C | `Notes loading.` | Known | UI element |
| 0x00194D08 | `Contacts loading.` | Known | UI element |
| 0x00195200 | `Sleep Timer` | Known | UI element |
| 0x0019520C | `Alarm Clock` | Known | UI element |
| 0x001952CC | `Reset Main Menu` | Known | Menu item |
| 0x001953C4 | `Reset All Settings` | Known | User setting |
| 0x00195458 | `Backlight Timer` | Known | UI element |
| 0x00195468 | `Repeat` | Known | UI element |
| 0x00195478 | `Main Menu` | Known | Menu item |
| 0x00195514 | `Settings` | Known | User setting |
| 0x00195530 | `Alarm Time` | Known | UI element |
| 0x00195834 | `Reset All` | Known | UI element |
| 0x001C2DFC | `vcalendar` | Known | UI element |
| 0x001C2EB0 | `dalarm` | Known | UI element |
| 0x001C3028 | `valarm` | Known | UI element |
| 0x001C3181 | `Contacts\` | Known | UI element |
| 0x001C3195 | `Calendars\` | Known | UI element |
| 0x001C31A6 | `Notes\` | Known | UI element |
| 0x001C3EF8 | `To check song links, set the preference NotesOnly to tr...` | Known | UI element |
| 0x001C4044 | `The NotesOnly pref can only be set globally in the Pref...` | Known | UI element |
| 0x001C4120 | `Warning: Preferences file must be in the Notes folder, ...` | Known | UI element |
| 0x001C452C | `Memory full. %d notes loaded, some notes not loaded. (3...` | Known | UI element |
| 0x0030B0DA | `BacklightTimer` | Known | UI element |
| 0x0030B0F1 | `RepeatMode` | Known | UI element |
| 0x0030B10D | `BacklightState` | Known | UI element |
| 0x0030B1D6 | `Is24HrClock` | Known | UI element |
| 0x0030B1F0 | `CurAlarm` | Known | UI element |
| 0x0030B1F9 | `CurAlarmText` | Known | UI element |
| 0x0030B206 | `ContactsDisplay` | Known | UI element |
| 0x0030B216 | `ContactsSort` | Known | UI element |
| 0x0030B223 | `MenuEnabled` | Known | Menu item |
| 0x003174EE | `Illegal instruction` | Known | UI element |
| 0x0031751C | `Illegal address` | Known | UI element |
| 0x00317622 | `NotesOnly` | Known | UI element |
| 0x0035EC7B | ` Set volume and other settings on audio chip.` | Known | User setting |
| 0x00360A03 | ` Display help menu` | Known | Menu item |
| 0x00361D41 | `<FN_NAME='%s' spFn_t=0x%08X MENU_OPTION=%d>` | Known | Menu item |
| 0x0041422C | ` CONTRAST` | Known | UI element |
| 0x0041772D | ` The contrast value is ,` | Known | UI element |
| 0x0041B9E4 | `N.CONTRAST` | Known | UI element |

---

## 4. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00151210 | `HostOSTask` | Known | RTOS task thread |
| 0x00151224 | `MP3ExampleTask` | Known | RTOS task thread |
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
| 0x00151638 | `RTXC v3.2fpp for ARM and Thumb - ARM ADS 1.0 Jul-08-00 ...` | Known | RTOS task thread |
| 0x001AF6F4 | `RunTestsTask` | Known | RTOS task thread |
| 0x001AFA28 | `LoadDataTasks` | Known | RTOS task thread |
| 0x001AFDB4 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x001C4588 | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x001C45A0 | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x001C4668 | `TrackCacheReadTask` | Known | RTOS task thread |

---

## 5. Audio/Codec

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0011DEE8 | `RIFFWAVEfmt data` | Known | PCM audio format |
| 0x00144BF8 | `RIFFWAVEfmt dataD` | Known | PCM audio format |
| 0x0015150C | `SerialAccsryMgr` | Known | Apple Lossless codec |
| 0x00159F28 | `Audible-softwaren i dette produkt bruges i henhold til ...` | Known | Audible audiobook format |
| 0x00159F88 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x0015A07A | `.net codec i dette produkt bruges i henhold til en lice...` | Known | Audio system |
| 0x0015A124 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunh...` | Known | Audio system |
| 0x0015DA4C | `Die Audible Software in diesem Produkt wird in Lizenz v...` | Known | Audible audiobook format |
| 0x0015DAA5 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x0015DB95 | `.net Codec in diesem Produkt wird mit der Lizenz der Vo...` | Known | Audio system |
| 0x0015DC5B | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und ...` | Known | Audio system |
| 0x00161094 | `El software Audible incluido en este producto se usa ba...` | Known | Audible audiobook format |
| 0x001610EF | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x00161291 | `n de audio MPEG Layer-3 utilizada bajo licencia de Frau...` | Known | Audio system |
| 0x0016461A | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright...` | Known | Audible audiobook format |
| 0x00164654 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x00164726 | `.net codec t` | Known | Audio system |
| 0x001647BC | `MPEG Layer-3 -` | Known | Audio system |
| 0x001647CE | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | Audio system |
| 0x00168760 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x001687AA | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x001687BF | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x0016886E | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x00168940 | `La technologie de codage audio MPEG Layer 3 est utilis` | Known | Audio system |
| 0x00168978 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | Audio system |
| 0x0016BD48 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x0016BD71 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0016BDA0 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x0016BE12 | ` utilizzato su licenza da VoiceAge Corporation. Il code...` | Known | Audio system |
| 0x0016BEE8 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da...` | Known | Audio system |
| 0x0016F8BD | ` Audible ` | Known | Audible audiobook format |
| 0x0016F8DE | `Audible ` | Known | Audible audiobook format |
| 0x0016F938 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x0016FAEC | `MPEG Layer-3 ` | Known | Audio system |
| 0x0016FB38 | `Fraunhofer IIS ` | Known | Audio system |
| 0x00173406 | ` Audible` | Known | Audible audiobook format |
| 0x0017344A | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x0017353A | `.net codec` | Known | Audio system |
| 0x001735FB | ` Fraunhofer IIS` | Known | Audio system |
| 0x00176F88 | `De Audible-software in dit product wordt gebruikt in li...` | Known | Audible audiobook format |
| 0x00176FDF | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x001770D0 | `.net-codec in dit product wordt gebruikt in licentie va...` | Known | Audio system |
| 0x0017716C | `Technologie voor codering van MPEG Layer-3-audio in lic...` | Known | Audio system |
| 0x0017A47C | `Audible-programvaren i dette produktet brukes under lis...` | Known | Audible audiobook format |
| 0x0017A4D0 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0017A64C | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fr...` | Known | Audio system |
| 0x0017DE10 | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x0017DE3F | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x0017DE56 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x0017DFF0 | `Ljudkodningstekniken MPEG Layer-3 ` | Known | Audio system |
| 0x0017E026 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | Audio system |
| 0x001816C6 | ` Fraunhofer IIS ` | Known | Audio system |
| 0x00194D1C | `The Audible software in this product is used under lice...` | Known | Audible audiobook format |
| 0x00194E55 | `.net codec in this product is used under license from V...` | Known | Audio system |
| 0x00194EE8 | `MPEG Layer-3 audio coding technology licensed from Frau...` | Known | Audio system |
| 0x001C3D59 | `&Aacute` | Known | AAC codec |
| 0x001C3E1F | `&aacute` | Known | AAC codec |
| 0x001C420F | `Boot time (MP3PlayerExampleApp constructor)` | Known | MP3 codec |
| 0x0030423D | ``0aLaCfDf` | Known | Apple Lossless codec |
| 0x00313C6B | `msCodeCom` | Known | Audio system |
| 0x003149EB | `aaControls` | Known | AAC codec |
| 0x0041F458 | `D1CODEC2.8` | Known | Audio system |
| 0x0041F47C | `D1 CODEC 2.8V` | Known | Audio system |

---

## 6. Storage/Hardware

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00003800 | `!ATAsoso` | Known | ATA/disk interface |
| 0x00003828 | `!ATAdpua` | Known | ATA/disk interface |
| 0x00057708 | `atad8@-` | Known | ATA/disk interface |
| 0x000A49E0 | `atadmrtsp@-` | Known | ATA/disk interface |
| 0x000C9370 | `atadmhbdmhsd8@-` | Known | ATA/disk interface |
| 0x000C9868 | `atadmhpo` | Known | ATA/disk interface |
| 0x000E8098 | `atadmhdp8@-` | Known | ATA/disk interface |
| 0x000F895C | `nutiatad` | Known | ATA/disk interface |
| 0x000F8BD0 | `atadmhdp` | Known | ATA/disk interface |
| 0x00151008 | `data abort` | Known | ATA/disk interface |
| 0x001512FC | `diskModeImageRev` | Known | Hardware interface |
| 0x001513BC | `FirewireGuid` | Known | FireWire interface |
| 0x00151461 | `diskmode` | Known | Hardware interface |
| 0x0015146A | `diskscan` | Known | Hardware interface |
| 0x00151528 | `FirewireInitiator` | Known | FireWire interface |
| 0x0015153C | `FirewireInterrupt` | Known | FireWire interface |
| 0x00151550 | `FirewireHandling` | Known | FireWire interface |
| 0x00158F94 | `Spiller nu` | Known | Hardware interface |
| 0x0015904C | `Spillelister` | Known | Hardware interface |
| 0x00159068 | `Genoptag spil` | Known | Hardware interface |
| 0x001594C0 | `Slet spilleliste` | Known | Hardware interface |
| 0x001594D4 | `Arkiver spilleliste` | Known | Hardware interface |
| 0x00159574 | `Ny spilleliste %lu` | Known | Hardware interface |
| 0x00159604 | `Harddisk` | Known | Hardware interface |
| 0x00159873 | ` den kan bruges som FireWire-disk, og tr` | Known | FireWire interface |
| 0x001599DC | `r du har tilsluttet iPod som disk, skal du dobbeltklikk...` | Known | Hardware interface |
| 0x00159C0A | ` den kan bruges som disk og anbringe tekstarkiver i map...` | Known | Hardware interface |
| 0x0015A37F | `je den til spillelisten On-The-Go. Hold knappen nede, n` | Known | Hardware interface |
| 0x0015A3B8 | `r en spilleliste, kunstner eller et album er valgt for ...` | Known | Hardware interface |
| 0x0015A3F5 | `je alle sangene til spillelisten On-The-Go.` | Known | Hardware interface |
| 0x0015A4AC | `Nyt spil` | Known | Hardware interface |
| 0x0015A5F0 | ` Spillelister` | Known | Hardware interface |
| 0x0015AB40 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x0015C964 | `Spiele` | Known | Hardware interface |
| 0x0015C9D8 | `Weiterspielen` | Known | Hardware interface |
| 0x0015D20B | `ffnen Sie das Addressbuch, Microsoft Entourage oder Pal...` | Known | FireWire interface |
| 0x0015DFF0 | `Neues Spiel` | Known | Hardware interface |
| 0x0015E706 | `ber FireWire verbunden` | Known | FireWire interface |
| 0x001608FC | `gido FireWire y arrastrar los archivos vCard a la carpe...` | Known | FireWire interface |
| 0x00161DA0 | `FireWire conectado` | Known | FireWire interface |
| 0x00163B88 | `Diskanttivahv.` | Known | Hardware interface |
| 0x00163B98 | `Diskanttiheikenn.` | Known | Hardware interface |
| 0x00163C5C | `Ladataan` | Known | ATA/disk interface |
| 0x00163EEE | `sin, avaa Osoitekirja, Microsoft Entourage tai Palm Des...` | Known | FireWire interface |
| 0x001645B0 | `Muistiinpanoja ladataan.` | Known | ATA/disk interface |
| 0x001645EC | `Yhteystietoja ladataan.` | Known | ATA/disk interface |
| 0x00165020 | `nityksen jatkamiseen ei ole tarpeeksi vapaata levytilaa...` | Known | ATA/disk interface |
| 0x00165060 | `nityksen aloittamiseen ei ole tarpeeksi vapaata levytil...` | Known | ATA/disk interface |
| 0x001652A0 | `FireWire liitetty` | Known | FireWire interface |
| 0x00167E4F | `lectionnez Appareils > Ajouter un appareil. Puis choisi...` | Known | FireWire interface |
| 0x0016948C | `FireWire Connect` | Known | FireWire interface |
| 0x0016B49F | ` archiviare contatti ed eventi di calendari. Se utilizz...` | Known | FireWire interface |
| 0x0016C1B8 | `Data & Ora` | Known | ATA/disk interface |
| 0x0016C280 | `Imposta Data & Ora` | Known | ATA/disk interface |
| 0x0016C444 | `Disattivata` | Known | ATA/disk interface |
| 0x0016C948 | `FireWire Connesso` | Known | FireWire interface |
| 0x0016EEF4 | ` FireWire ` | Known | FireWire interface |
| 0x0017083C | `FireWire ` | Known | FireWire interface |
| 0x00172BAB | `  FireWire ` | Known | FireWire interface |
| 0x00176652 | `Op de iPod kunt u adres- en agendagegevens opslaan. Als...` | Known | FireWire interface |
| 0x00177C00 | `FireWire aangesloten` | Known | FireWire interface |
| 0x00179470 | `Spilles n` | Known | Hardware interface |
| 0x00179554 | `Fortsett spill` | Known | Hardware interface |
| 0x0017996C | `Mer diskant` | Known | Hardware interface |
| 0x00179978 | `Mindre diskant` | Known | Hardware interface |
| 0x001799BC | `Slett spilleliste` | Known | Hardware interface |
| 0x001799D0 | `Arkiver spillelister` | Known | Hardware interface |
| 0x00179B14 | `Diskmodus` | Known | Hardware interface |
| 0x00179CFE | `pner du Adressebok, Microsoft Entourage eller Palm Desk...` | Known | FireWire interface |
| 0x00179E5E | `ringer i tillegg til musikken din. Microsoft Outlook, M...` | Known | FireWire interface |
| 0x0017A114 | `Hvis du vil vise tekstfiler her, aktiverer du iPod for ...` | Known | FireWire interface |
| 0x0017A88F | ` legge den til i On-The-Go-spillelisten. Du kan legge t...` | Known | Hardware interface |
| 0x0017A9B4 | `Nytt spill` | Known | Hardware interface |
| 0x0017AE1C | `Det er ikke nok ledig diskplass til ` | Known | Hardware interface |
| 0x0017B084 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x0017D6F2 | `ll sedan in din iPod som FireWire-h` | Known | FireWire interface |
| 0x0017D717 | `rddisk. D` | Known | Hardware interface |
| 0x0017D878 | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x0017D898 | `rddisk. Sedan dubbelklickar du bara p` | Known | Hardware interface |
| 0x0017DAF8 | `rddiskl` | Known | Hardware interface |
| 0x0017DB98 | `inget kort inmatat` | Known | ATA/disk interface |
| 0x0017EA7C | `FireWire anslutet` | Known | FireWire interface |
| 0x0019444C | `Disk Mode` | Known | Hardware interface |
| 0x00194522 | `Your iPod can store contacts and calendar events. If yo...` | Known | FireWire interface |
| 0x001946F6 | `Your iPod can store up to one thousand contacts right a...` | Known | Hardware interface |
| 0x001949BC | `To view text files here, enable iPod for disk use, then...` | Known | Hardware interface |
| 0x00195694 | `There is not enough free disk space to continue recordi...` | Known | Hardware interface |
| 0x001956D0 | `There is not enough free disk space to start recording.` | Known | Hardware interface |
| 0x00195930 | `FireWire Connected` | Known | FireWire interface |
| 0x00195944 | `No battery power remains. Please connect iPod to power.` | Known | Power management |
| 0x0019597C | `Low Battery` | Known | Power management |
| 0x001C0FB4 | `TROMResourceDB - unknown header version! (Try regenerat...` | Known | ATA/disk interface |
| 0x001C3660 | `Bad Data` | Known | ATA/disk interface |
| 0x001C400C | `Bad meta data, name not a recognized preference. (27)` | Known | ATA/disk interface |
| 0x001C40C4 | `Bad meta data, name not found. (23)` | Known | ATA/disk interface |
| 0x001C40E8 | `Bad meta data, name termination quote not found. (24)` | Known | ATA/disk interface |
| 0x001C4170 | `Bad meta data, content not found. (25)` | Known | ATA/disk interface |
| 0x001C4198 | `Bad meta data, content termination quote not found. (26...` | Known | ATA/disk interface |
| 0x001C42C4 | `Bad meta data, encoding not found. (8)` | Known | ATA/disk interface |
| 0x001C42EC | `Bad meta data, encoding termination quote not found. (9...` | Known | ATA/disk interface |
| 0x001C4328 | `Bad meta data, unknown encoding. (17)` | Known | ATA/disk interface |
| 0x001C44E4 | `Bad data. (32)` | Known | ATA/disk interface |
| 0x001C46E4 | `C:\iPod\tagged_checkout\sobek_build\Q14\Sources\Service...` | Known | Hardware interface |
| 0x001C4734 | `offset < fDataLength` | Known | ATA/disk interface |
| 0x002F2849 | `aGa>a(a'aJa?a<a,a4a=aBaDasawaXaYaZakataoaeaqa_a]aSaua` | Known | ATA/disk interface |
| 0x002FD357 | `ataTaza[aea;ajaaaVa)b'b+b+dMd[d]dtdvdrdsd}dudfd` | Known | ATA/disk interface |
| 0x0030AE98 | `data!=NULL` | Known | ATA/disk interface |
| 0x0030B134 | `BatteryLevel` | Known | Power management |
| 0x0030B141 | `FirewirePower` | Known | FireWire interface |
| 0x0030EF22 | `ex_data` | Known | ATA/disk interface |
| 0x0030F0D4 | `C:\iPod\tagged_checkout\sobek_build\Q14\Sources\Service...` | Known | ATA/disk interface |
| 0x003130BF | `RSA Data Security, Inc.` | Known | ATA/disk interface |
| 0x003130DC | `RSA Data Security, Inc. PKCS` | Known | ATA/disk interface |
| 0x0031323D | `pkcs7-data` | Known | ATA/disk interface |
| 0x00313248 | `pkcs7-signedData` | Known | ATA/disk interface |
| 0x00313259 | `pkcs7-envelopedData` | Known | ATA/disk interface |
| 0x0031326D | `pkcs7-signedAndEnvelopedData` | Known | ATA/disk interface |
| 0x0031328A | `pkcs7-digestData` | Known | ATA/disk interface |
| 0x0031329B | `pkcs7-encryptedData` | Known | ATA/disk interface |
| 0x003134F4 | `nsDataType` | Known | ATA/disk interface |
| 0x003134FF | `Netscape Data Type` | Known | ATA/disk interface |
| 0x00314299 | `id-smime-ct-authData` | Known | ATA/disk interface |
| 0x00314306 | `id-smime-ct-DVCSRequestData` | Known | ATA/disk interface |
| 0x00314322 | `id-smime-ct-DVCSResponseData` | Known | ATA/disk interface |
| 0x00314D3A | `id-cmc-dataReturn` | Known | ATA/disk interface |
| 0x00314E8D | `id-on-personalData` | Known | ATA/disk interface |
| 0x00314F8E | `id-cct-PKIData` | Known | ATA/disk interface |
| 0x0031572D | `qualityLabelledData` | Known | ATA/disk interface |
| 0x00315B08 | `setct-PANData` | Known | ATA/disk interface |
| 0x00315B33 | `setct-OIData` | Known | ATA/disk interface |
| 0x00315B49 | `setct-PIData` | Known | ATA/disk interface |
| 0x00315B56 | `setct-PIDataUnsigned` | Known | ATA/disk interface |
| 0x00315BD1 | `setct-PInitResData` | Known | ATA/disk interface |
| 0x00315BF1 | `setct-PResData` | Known | ATA/disk interface |
| 0x00315C47 | `setct-CapTokenData` | Known | ATA/disk interface |
| 0x00315C95 | `setct-AuthRevResData` | Known | ATA/disk interface |
| 0x00315CDF | `setct-CapResData` | Known | ATA/disk interface |
| 0x00315D17 | `setct-CapRevResData` | Known | ATA/disk interface |
| 0x00315D4E | `setct-CredResData` | Known | ATA/disk interface |
| 0x00315D89 | `setct-CredRevResData` | Known | ATA/disk interface |
| 0x00315D9E | `setct-PCertReqData` | Known | ATA/disk interface |
| 0x00315DC3 | `setct-BatchAdminReqData` | Known | ATA/disk interface |
| 0x00315DDB | `setct-BatchAdminResData` | Known | ATA/disk interface |
| 0x00315E33 | `setct-CertReqData` | Known | ATA/disk interface |
| 0x00315E56 | `setct-CertResData` | Known | ATA/disk interface |
| 0x003161E2 | `setCext-merchData` | Known | ATA/disk interface |
| 0x0031626D | `setCext-Track2Data` | Known | ATA/disk interface |
| 0x0031645B | `set-brand-IATA-ATA` | Known | ATA/disk interface |
| 0x00343224 | `Predictor Data Present, invalid output!` | Known | ATA/disk interface |
| 0x00356D85 | `cmp_sct_data:  adrx = %d, Wtbuf[adrx] = %x, Rdbuf[adrx]...` | Known | ATA/disk interface |
| 0x00357A80 | ` Data error sts=%d` | Known | ATA/disk interface |
| 0x0035E700 | `Running on FireWire power.` | Known | FireWire interface |
| 0x0035E720 | `Running on battery power.` | Known | Power management |
| 0x0035E8DF | `@Battery reading is %d (%x) E%4cF` | Known | Power management |
| 0x0035E963 | `Switch to USB disk mode` | Known | USB interface |
| 0x0035E9D7 | `Turn on/off hard disk` | Known | Hardware interface |
| 0x0035EA17 | `Turn on/off firewire` | Known | FireWire interface |
| 0x0035EA37 | `Turn on/off battery charging` | Known | Power management |
| 0x0035F874 | `!ATADisk is FDISK format.` | Known | ATA/disk interface |
| 0x0035F890 | `Disk is Mac format.` | Known | Hardware interface |
| 0x0035F8A8 | `Disk contains %d interesting partitions.` | Known | Hardware interface |
| 0x0035FAA4 | `Write a boot record when formatting FDISK.` | Known | Hardware interface |
| 0x00361B8D | `<BATTERY MEASUREMENT [ %d.%d , 0x%X ]>` | Known | Power management |
| 0x00361BD5 | `<DISK-MODE  V%X.%X.%X>` | Known | Hardware interface |
| 0x00362868 | `Disk mode` | Known | Hardware interface |
| 0x00362874 | `Disk Scan mode` | Known | Hardware interface |
| 0x00363E30 | `SP_ERR_NULL_DATAINFUNC` | Known | ATA/disk interface |
| 0x00363E48 | `SP_ERR_NULL_DATAOUTFUNC` | Known | ATA/disk interface |
| 0x00363FAC | `SP_ERR_UNSUPPORTED_DATACLASS` | Known | ATA/disk interface |
| 0x00363FCC | `SP_ERR_UNSUPPORTED_DATATYPE` | Known | ATA/disk interface |
| 0x00363FE8 | `SP_ERR_UNSUPPORTED_DATASUBTYPE` | Known | ATA/disk interface |
| 0x0036414C | `SP_ERR_OTHER_DATAINOUT_ERR` | Known | ATA/disk interface |
| 0x00415A3F | `@diskmode` | Known | Hardware interface |
| 0x00419618 | `USB DISK` | Known | USB interface |
| 0x00419624 | `FW DISK` | Known | Hardware interface |
| 0x0041B174 | `DISK MODE` | Known | Hardware interface |
| 0x0041B9CC | `L.USB DISK` | Known | USB interface |
| 0x0041BB40 | `F.FIREWIRE` | Known | FireWire interface |
| 0x0041FD2C | `"Apple   iPod Disk Drive` | Known | Hardware interface |
| 0x0044E624 | `Stand-alone Disk-Mode 2.0 running` | Known | Hardware interface |

---

## 7. Error Messages

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0001BB3C | `*** assertion failed: %s, file %s, line %d` | Known | Error/assertion message |
| 0x00021CC8 | `Invalid Operation` | Known | Error/assertion message |
| 0x00160F7C | `Error durante la  importanci` | Known | Error/assertion message |
| 0x0016BC19 | ` verificato un errore durante l'importazione` | Known | Error/assertion message |
| 0x00194BE0 | `An error occurred while importing` | Known | Error/assertion message |
| 0x0019565C | `Cannot record because there is no microphone attached.` | Known | Error/assertion message |
| 0x001C3F5C | `Too many errors, further errors discarded. (31)` | Known | Error/assertion message |
| 0x001C3FAC | `%s Error in file %s.` | Known | Error/assertion message |
| 0x001C4370 | `Cannot link to a .link file. (29)` | Known | Error/assertion message |
| 0x001C4647 | `load failed (%d)` | Known | Error/assertion message |
| 0x001C4882 | `tInvalid partition table. Setup cannot continue.` | Known | Error/assertion message |
| 0x001C48B3 | `Error loading operating system. Setup cannot continue.` | Known | Error/assertion message |
| 0x0030F07C | `%s(%d): OpenSSL internal error, assertion failed: %s` | Known | Error/assertion message |
| 0x0030FE88 | `error:%08lX:%s:%s:%s` | Known | Error/assertion message |
| 0x00316BFC | `internal error: list index %ld out of range` | Known | Error/assertion message |
| 0x00355E7B | `( <- Error:Unsupported` | Known | Error/assertion message |
| 0x00355E94 | ` <- Error:Internal Error` | Known | Error/assertion message |
| 0x00355EBF | `(Invalid Key Value arguments` | Known | Error/assertion message |
| 0x003566AC | `***Error r_reg Unknown Ide Reg(%X)` | Known | Error/assertion message |
| 0x00356788 | `***Error w_reg Unknown Ide Reg(%X)` | Known | Error/assertion message |
| 0x00357001 | `command:  head No. error` | Known | Error/assertion message |
| 0x00357021 | `command:  DRVH reg. compare error` | Known | Error/assertion message |
| 0x00357049 | `command:  CYL_high reg. compare error` | Known | Error/assertion message |
| 0x00357075 | `command:  CYL_low reg. compare error` | Known | Error/assertion message |
| 0x003570A1 | `command:  SECT_num reg. compare error` | Known | Error/assertion message |
| 0x003570CD | `command:  SECT_cnt reg. compare error` | Known | Error/assertion message |
| 0x00357C4D | `xfer_command: command failed.  sts=%x` | Known | Error/assertion message |
| 0x00357F25 | `FW transfer error status_reg=%x` | Known | Error/assertion message |
| 0x00357F75 | `Check FW error status_reg=%x` | Known | Error/assertion message |
| 0x00358111 | `FW Write error status_reg=%x` | Known | Error/assertion message |
| 0x00358131 | `DoToshiba: stand-by immediate command error status_reg=...` | Known | Error/assertion message |
| 0x003581A9 | `Write Copyright error status_reg=%x` | Known | Error/assertion message |
| 0x003581F9 | `Restart error status_reg=%x` | Known | Error/assertion message |
| 0x003582E5 | `doToshiba:  sts error.  sts = %x` | Known | Error/assertion message |
| 0x0035B4DD | `get_fw_rev:  identify command failed.  cannot get FW Re...` | Known | Error/assertion message |
| 0x0035B529 | `check_fw:  buffer compare error.  i = %x, Wtbuf[i] = %x...` | Known | Error/assertion message |
| 0x0035E73C | `Power status is invalid.` | Known | Error/assertion message |
| 0x0035F85C | `Error opening device.` | Known | Error/assertion message |
| 0x0035F9F0 | `Device %s error reading sectors.` | Known | Error/assertion message |
| 0x00363BE4 | `SP_ERR_INIT_FAILED` | Known | Error/assertion message |
| 0x00363EB0 | `SP_ERR_NULL_ERRORFUNC` | Known | Error/assertion message |
| 0x003640CC | `SP_ERR_INVALID_CHANNEL_ID` | Known | Error/assertion message |
| 0x003640E8 | `SP_ERR_INVALID_COMMAND` | Known | Error/assertion message |
| 0x0036428C | `SP_ERR_HW_RESET_FAILED` | Known | Error/assertion message |
| 0x0041FAA2 | `XMODEM/XMODEM-1K transfer failed.` | Known | Error/assertion message |
| 0x0044E990 | `btGrInitialize: display info invalid` | Known | Error/assertion message |

---

## 8. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00153790 | ` !"#$%&'()*+,-./0123456789:;<=>?@abcdefghijklmnopqrstuv...` | Known | Filesystem path |
| 0x00153890 | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x00159DDC | `Point: %d (%d/%d)` | Known | Filesystem path |
| 0x00159FFB | `rende VoiceAge Corporation i USA og/eller andre lande o...` | Known | Filesystem path |
| 0x0015CA80 | `USA/Hawaii (NZ)` | Known | Filesystem path |
| 0x0015CA90 | `USA/Hawaii (SZ)` | Known | Filesystem path |
| 0x0015CAA0 | `USA/Alaska (NZ)` | Known | Filesystem path |
| 0x0015CAB0 | `USA/Alaska (SZ)` | Known | Filesystem path |
| 0x0015CAC0 | `USA/Pazifik (NZ)` | Known | Filesystem path |
| 0x0015CAD4 | `USA/Pazifik (SZ)` | Known | Filesystem path |
| 0x0015CAE8 | `USA/Rockies (NZ)` | Known | Filesystem path |
| 0x0015CAFC | `USA/Rockies (SZ)` | Known | Filesystem path |
| 0x0015CB10 | `USA/Zentral (NZ)` | Known | Filesystem path |
| 0x0015CB24 | `USA/Zentral (SZ)` | Known | Filesystem path |
| 0x0015CB38 | `USA/Ost (NZ)` | Known | Filesystem path |
| 0x0015CB48 | `USA/Ost (SZ)` | Known | Filesystem path |
| 0x0015CE4C | `Vorn./Nachn.` | Known | Filesystem path |
| 0x0015CE5C | `Nachn./Vorn.` | Known | Filesystem path |
| 0x0015D8F0 | `Punkte: %d (%d/%d)` | Known | Filesystem path |
| 0x0015DADF | ` ist entweder eine eingetragene Marke oder eine Marke d...` | Known | Filesystem path |
| 0x00160F44 | `Result.: %d (%d/%d)` | Known | Filesystem path |
| 0x0016112F | ` es una marca registrada o una marca comercial de Voice...` | Known | Filesystem path |
| 0x00161CE0 | `Fecha/hora` | Known | Filesystem path |
| 0x00163C54 | `%d / %d` | Known | Filesystem path |
| 0x00164398 | `%s / %s` | Known | Filesystem path |
| 0x001643C4 | `%d / %d valokuvaa tuotu` | Known | Filesystem path |
| 0x001644C0 | `Tulos: %d (%d/%d)` | Known | Filesystem path |
| 0x001646C5 | `ity tavaramerkki Yhdysvalloissa ja/tai muissa maissa, j...` | Known | Filesystem path |
| 0x001685EF | `sult. : %d (%d/%d)` | Known | Filesystem path |
| 0x00168842 | `tats-Unis et/ou dans d'autres pays, utilis` | Known | Filesystem path |
| 0x00168D5F | `gler date/heure` | Known | Filesystem path |
| 0x001693C4 | `Date/heure` | Known | Filesystem path |
| 0x0016BBE0 | `Punti: %d (%d/%d)` | Known | Filesystem path |
| 0x0016E198 | `%b/%-d %-I:%M %2p` | Known | Filesystem path |
| 0x0016E1AC | `%-m/%-d` | Known | Filesystem path |
| 0x0016E1D4 | `%y/%-m/%-d` | Known | Filesystem path |
| 0x0016E1E0 | `%Y/%b/%-d` | Known | Filesystem path |
| 0x0016F6F8 | ` %d (%d/%d)` | Known | Filesystem path |
| 0x001720D4 | `%Y/%B/%d %A  %I:%M:%S %p` | Known | Filesystem path |
| 0x001720F0 | `%Y/%B/%d` | Known | Filesystem path |
| 0x00172108 | `%-m/%-d %-I:%M %p` | Known | Filesystem path |
| 0x0017213C | `%Y/-%-m/-%-d` | Known | Filesystem path |
| 0x00173124 | `%d / %d ` | Known | Filesystem path |
| 0x0017327A | `: %d (%d/%d)` | Known | Filesystem path |
| 0x00176E48 | `Score: %d (%d/%d)` | Known | Filesystem path |
| 0x00177017 | ` is een gedeponeerd handelsmerk of een handelsmerk van ...` | Known | Filesystem path |
| 0x00177584 | `Stel datum/tijd in` | Known | Filesystem path |
| 0x0017A330 | `Poeng: %d (%d/%d)` | Known | Filesystem path |
| 0x0017A50F | ` er enten et registrert varemerke eller et varemerke fo...` | Known | Filesystem path |
| 0x0017CDEC | `%-d/%-m` | Known | Filesystem path |
| 0x0017DCCC | `ng: %d (%d/%d)` | Known | Filesystem path |
| 0x0017DEDC | `r VoiceAge Corporation i USA och/eller andra l` | Known | Filesystem path |
| 0x0017E3F0 | `ll in datum/tid` | Known | Filesystem path |
| 0x001802DC | `%Y/%-m/%-d` | Known | Filesystem path |
| 0x00181289 | ` %d/%d ` | Known | Filesystem path |
| 0x00181395 | `%d (%d/%d)` | Known | Filesystem path |
| 0x0019393C | `%-m/%-d/%y` | Known | Filesystem path |
| 0x00194DA7 | ` is either registered trademark or trademark of VoiceAg...` | Known | Filesystem path |
| 0x001AF834 | `url;type=work:apple.com/support/ipod` | Known | Filesystem path |
| 0x001C33E5 | `file://` | Known | Filesystem path |
| 0x001C33ED | `image://` | Known | Filesystem path |
| 0x001C35CC | `</TITLE>` | Known | Filesystem path |
| 0x001C35DC | `</BODY>` | Known | Filesystem path |
| 0x001C360A | `</ROT13>` | Known | Filesystem path |
| 0x001C43B4 | `Bad link, no matching </a> for anchor tag. (20)` | Known | Filesystem path |
| 0x001D707D | `/" %BD"` | Known | Filesystem path |
| 0x001DD9A0 | `ODD""(/` | Known | Filesystem path |
| 0x001DEBDB | `B"$R%"B$" /` | Known | Filesystem path |
| 0x001DF548 | `bG\|jG\|/` | Known | Filesystem path |
| 0x001E1132 | `$E$$BR/` | Known | Filesystem path |
| 0x001E11D3 | `dRB~RA$/` | Known | Filesystem path |
| 0x001E1B24 | `TT&T%B(/` | Known | Filesystem path |
| 0x001F0818 | `)'>$B8/` | Known | Filesystem path |
| 0x001F2E26 | `$B\|%EV/` | Known | Filesystem path |
| 0x001FA37A | `BDU!BJ ""/` | Known | Filesystem path |
| 0x001FB2DE | `Z-bD("(%B>/` | Known | Filesystem path |
| 0x00203A2D | `@(/ Q\|f` | Known | Filesystem path |
| 0x0023459D | `T/DDDDD` | Known | Filesystem path |
| 0x00234807 | `"~UeB /` | Known | Filesystem path |
| 0x002372B9 | `$B((B /` | Known | Filesystem path |
| 0x0023F39C | ` "\|$B~/` | Known | Filesystem path |
| 0x00242518 | `@$B\|$"(/` | Known | Filesystem path |
| 0x00243618 | `)"8/B""` | Known | Filesystem path |
| 0x00243D60 | `r4c6 bN/` | Known | Filesystem path |
| 0x00249161 | `RDT%B(/` | Known | Filesystem path |
| 0x0024A2ED | `RBHUE\|/` | Known | Filesystem path |
| 0x0025196D | `]B""B</` | Known | Filesystem path |
| 0x002551EA | `,B\|RED/` | Known | Filesystem path |
| 0x0025AC6D | `$BT). /` | Known | Filesystem path |
| 0x0025BE6D | `#"TUB(/` | Known | Filesystem path |
| 0x002610E1 | `H."0*Bx/` | Known | Filesystem path |
| 0x00268A3E | `U/~RERT` | Known | Filesystem path |
| 0x0026D7A6 | `TUOPT/\|` | Known | Filesystem path |
| 0x00274A3B | `HuGZp/$j` | Known | Filesystem path |
| 0x0027AF4B | `(bJ)b"(b2""z/` | Known | Filesystem path |
| 0x0027D6BB | `JUAPDD(/` | Known | Filesystem path |
| 0x002839CA | `/B\|$BD'` | Known | Filesystem path |
| 0x00284947 | `$Bd$BT/` | Known | Filesystem path |
| 0x0028AC1F | `/" +J\|!` | Known | Filesystem path |
| 0x00291A9E | `Fb""")/` | Known | Filesystem path |
| 0x00292C55 | `/RyO(UIH` | Known | Filesystem path |
| 0x00293D05 | `~$Bh'"~$Bz/` | Known | Filesystem path |
| 0x0029756F | `$B +BZ/` | Known | Filesystem path |
| 0x0029EF8D | `0c(HBP/` | Known | Filesystem path |
| 0x002A3523 | `$B~("\|/` | Known | Filesystem path |
| 0x002C0333 | `W/}lE>q` | Known | Filesystem path |
| 0x002F08DB | `%#%3%+%;%K% %/%(%7%?%` | Known | Filesystem path |
| 0x002F0A5A | `02*2+2,2-2.2/2@272B2C292:212>24222;2623252<2=2?282` | Known | Filesystem path |
| 0x002F0F85 | `\7_J_/`P`m`` | Known | Filesystem path |
| 0x002F1161 | `OHSIT>T/Z` | Known | Filesystem path |
| 0x002F11D7 | `i_l*mim/n` | Known | Filesystem path |
| 0x002F16C9 | `N,p]u/f` | Known | Filesystem path |
| 0x002F23D5 | `S#S/S1S3S8S@SFSES` | Known | Filesystem path |
| 0x002F24D5 | `q4V6V2V8VkVdV/VlVjV` | Known | Filesystem path |
| 0x002F277D | `_)_-_8_A_H_L_N_/_Q_V_W_Y_a_m_s_w_` | Known | Filesystem path |
| 0x002F2E43 | `s4s/s)s%s>sNsOs` | Known | Filesystem path |
| 0x002F4CA1 | `% %!%"%#%$%%%&%'%(%)%*%+%,%-%.%/%0%1%2%3%4%5%6%7%8%9%:%...` | Known | Filesystem path |
| 0x002F5481 | `qWlIl/Ymg*` | Known | Filesystem path |
| 0x002F6CEF | `X)W,W*W3W9W.W/W\W;WBWiW` | Known | Filesystem path |
| 0x002F73BD | `n/o6oKoto*o` | Known | Filesystem path |
| 0x002F7955 | `q/q1qsq\qhqEqrqJqxqzq` | Known | Filesystem path |
| 0x002F8BDD | `b6bKbNb/e` | Known | Filesystem path |
| 0x002F8D93 | `V3W0W(W-W,W/W)W` | Known | Filesystem path |
| 0x002F8E67 | `NMOOOGOWO^O4O[OUO0OPOQO=O:O8OCOTO<OFOcO\O`O/ONO6OYO]OHO...` | Known | Filesystem path |
| 0x002F94C5 | `e%f-f f'f/f` | Known | Filesystem path |
| 0x002F97C7 | `bNc>c/cUcBcFcOcIc:cPc=c*c+c(cMcLcHeIe` | Known | Filesystem path |
| 0x002F9A79 | `fFUjUfUDU^UaUCUJU1UVUOUUU/UdU8U.U\U,UcU3UAUWU` | Known | Filesystem path |
| 0x002F9CDF | `\|F}C}q}.}9}<}@}0}3}D}/}B}2}1}=` | Known | Filesystem path |
| 0x002F9E67 | `W/X*X4X$X0X1X!X` | Known | Filesystem path |
| 0x002F9F7D | `k.l/l,l/n8nTn!n2ngnJn n%n#n` | Known | Filesystem path |
| 0x002F9FDF | `r6s%s4s)s:t*t3t"t%t5t6t4t/t` | Known | Filesystem path |
| 0x002FA85D | `S.V;V9V2V?V4V)VSVNVWVtV6V/V0V` | Known | Filesystem path |
| 0x002FAEA3 | `\|.~>~F~7~2~C~+~=~1~E~A~4~9~H~5~?~/~D` | Known | Filesystem path |
| 0x002FBE93 | `P P'P5P/P1P` | Known | Filesystem path |
| 0x002FBFE5 | `h5h+h-h/hNhDh4h` | Known | Filesystem path |
| 0x002FBFF9 | `h&h(h.hMh:h%h h,k/k-k1k4kmk` | Known | Filesystem path |
| 0x002FC105 | `w"w'w#x,x"x5x/x(x.x+x!x)x3x*x1xTy[yOy\ySyRyQy` | Known | Filesystem path |
| 0x002FC58F | `v9w/w-w1w2w4w3w=w%w;w5wHxRxIxMxJxLx&xExPxdygyiyjycykyay` | Known | Filesystem path |
| 0x002FC5EB | `{1{+{-{/{2{8{` | Known | Filesystem path |
| 0x002FCA7F | `t/uoulu` | Known | Filesystem path |
| 0x002FCDE5 | ``IaJa+aEa6a2a.aFa/aOa)a@a bh` | Known | Filesystem path |
| 0x002FCE2B | `d&d0d(dAd5d/d` | Known | Filesystem path |
| 0x002FCE93 | `hxi4iii@ioiDiviXiAitiLi;iKi7i\iOiQi2iRi/i{i<iFkEkCkBkHk...` | Known | Filesystem path |
| 0x002FD4FD | `y+zJz0z/z(z&z` | Known | Filesystem path |
| 0x002FDDE1 | `u/v-v1v=v3v<v5v2v0v` | Known | Filesystem path |
| 0x002FE0D7 | `X-[%[2[#[,['[&[/[.[{[` | Known | Filesystem path |
| 0x002FEA05 | `j<p5p/p7p4p1pBp8p?p:p9p@p;p3pAp` | Known | Filesystem path |
| 0x002FEC07 | `w-y1y/yT\|S\|` | Known | Filesystem path |
| 0x00302047 | `h>kLp/t` | Known | Filesystem path |
| 0x003026D1 | `o;v/}7~` | Known | Filesystem path |
| 0x00303495 | `e1f/h\q6z` | Known | Filesystem path |
| 0x00303AE1 | `UuX/c"dIfKfmh` | Known | Filesystem path |
| 0x00307FC8 | `SASUC -D5S00A/` | Known | Filesystem path |
| 0x0030F764 | `You need to read the OpenSSL FAQ, http://www.openssl.or...` | Known | Filesystem path |
| 0x00313FB9 | `S/MIME Capabilities` | Known | Filesystem path |
| 0x0031A9AA | `/ 0 1 2 3 4 5 6 7 8 9 : ; < = > ? @ A B C D E F G H I J...` | Known | Filesystem path |
| 0x0031AB8B | `! !!!"!#!$!%!&!'!(!)!*!+!,!-!.!/!0!1!2!3!4!5!6!7!8!9!:!...` | Known | Filesystem path |
| 0x00337688 | `/1f;{1Q` | Known | Filesystem path |
| 0x0035E268 | `Syntax: %s [FourCharKeyName] [val0] [val1 val2 val3] [/...` | Known | Filesystem path |
| 0x0035E9F7 | `Turn on/off backlight` | Known | Filesystem path |
| 0x0035EA58 | `/u[012]` | Known | Filesystem path |
| 0x0035EA60 | `Set audio to off/standby/on` | Known | Filesystem path |
| 0x0035EB31 | `/w warm reboot instead of cold` | Known | Filesystem path |
| 0x0035ECAC | `/v[0-9]` | Known | Filesystem path |
| 0x0035F344 | `Syntax : %s [+\|-\|HexAddress] [HexSize] [/b\|/w\|/d]` | Known | Filesystem path |
| 0x0035F50C | `Syntax : %s HexAddress HexVal [/b\|/w\|/d]` | Known | Filesystem path |
| 0x0035FED3 | `List available images. [/f with full info]` | Known | Filesystem path |
| 0x00360198 | `%x bytes read/written.` | Known | Filesystem path |
| 0x003601CF | ` Read/write files on the host.` | Known | Filesystem path |
| 0x003609C7 | `/? option gives command specific help (i.e., dir /?)` | Known | Filesystem path |
| 0x00361220 | `</MM_NODE>` | Known | Filesystem path |
| 0x003617E8 | `</HANDLE` | Known | Filesystem path |
| 0x00361844 | `</APIMGR>` | Known | Filesystem path |
| 0x00361A1D | `<BOARD RTC ADJUST         %c%d.%02d SEC/DAY>` | Known | Filesystem path |
| 0x00361C45 | `</ADDRESS RANGES>` | Known | Filesystem path |
| 0x00361C58 | `</SYSINFO>` | Known | Filesystem path |
| 0x00361D70 | `</EXPORTS>` | Known | Filesystem path |
| 0x00361DD0 | `</IMPORTS>` | Known | Filesystem path |
| 0x0036D158 | `{{~~  /-----\   {{~~ /       \  {{~~\|         \| {{~~\...` | Known | Filesystem path |
| 0x004163F6 | `Begin XMODEM/XMODEM-1K upload of IRAM image file..........` | Known | Filesystem path |
| 0x0041FAC5 | `Restart XMODEM/XMODEM-1K upload of image file...` | Known | Filesystem path |

---

## 9. Binary Structure (for Ghidra/IDA)

| Parameter | Value |
|-----------|-------|
| **Architecture** | ARM (32-bit, little-endian) |
| **Processor** | ARM7TDMI (ARMv4T), dual-core |
| **Base Address** | 0x00000000 |
| **Entry Point** | 0x00000000 |
| **Endianness** | Little-Endian |
| **File Size** | 4,561,920 bytes |

