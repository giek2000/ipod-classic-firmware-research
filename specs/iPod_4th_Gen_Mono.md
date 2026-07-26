# iPod 4th Generation (Monochrome) - RetailOS 4.3.1.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 4.3.1.1 |
| **IPSW** | iPod_4.3.1.1.ipsw |
| **Device** | iPod 4th Generation Click Wheel (2004, 20/40GB, Grayscale) |
| **Binary Size** | 4,605,440 bytes (4.39 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 4,605,440 bytes |
| **Total Strings (>=6)** | 10,743 |
| **Function Prologues** | 7,824 |
| **SoC** | PortalPlayer PP5020 |
| **Architecture** | ARM7TDMI (ARMv4T), dual-core |
| **Encrypted** | No |
| **Decryption Method** | Extract from IPSW (unencrypted) |
| **SHA-256** | `9c959c47dbc17f78a65936c5461749c71af39f36744033881c981ce39cb872ca` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174294 | `MP3ExampleTask` | Hidden | Hidden Test |

---

## 2. Discovered Features

### EQ Preset

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
| 0x0017EBA4 | `USA/Rockies (NZ)` | EQ Preset | |
| 0x0017EBB8 | `USA/Rockies (SZ)` | EQ Preset | |
| 0x00182658 | `Latina` | EQ Preset | |
| 0x0019C920 | `Latino` | EQ Preset | |
| 0x001E5B25 | `LATIN-1` | EQ Preset | |
| 0x001E5B2D | `LATIN1` | EQ Preset | |
| 0x002D9907 | `~ BR&B$"` | EQ Preset | |
| 0x003301FE | `Secure Electronic Transactions` | EQ Preset | |

### Localization

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001792F4 | `x-mac-japanese` | Localization | |
| 0x001B7560 | `English` | Localization | |
| 0x001B7598 | `Italiano` | Localization | |
| 0x001E5B90 | `X-MAC-JAPANESE` | Localization | |
| 0x001E5B9F | `MAC-JAPANESE` | Localization | |
| 0x001E5BAC | `MACJAPANESE` | Localization | |
| 0x001E5BD8 | `X-MAC-CHINESETRAD` | Localization | |
| 0x001E5BEA | `MAC-CHINESETRAD` | Localization | |
| 0x001E5C07 | `X-MAC-CHINESESIMP` | Localization | |
| 0x001E5C19 | `MAC-CHINESESIMP` | Localization | |
| 0x001E5C39 | `X-MAC-KOREAN` | Localization | |
| 0x001E5C46 | `MAC-KOREAN` | Localization | |

### Filesystem Path

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174178 | `iPod_Control\Device` | Filesystem Path | |
| 0x0017418C | `iPod_Control` | Filesystem Path | |
| 0x0017419C | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x00174904 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x0017499C | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x001749CC | `iPod_Control\Device\` | Filesystem Path | |
| 0x00174A28 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x00174A70 | `iPod_Control\Music\` | Filesystem Path | |

### Assertion

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00017740 | `*** assertion failed: %s, file %s, line %d` | Assertion | |
| 0x0032C7B4 | `%s(%d): OpenSSL internal error, assertion failed: %s` | Assertion | |

---

## 3. Known User-Facing Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00174234 | `KeyRepeatTimer` | Known | UI element |
| 0x0017B78C | `k dine vCards til mappen Contacts p` | Known | UI element |
| 0x0017B921 | `kke vCard-arkiverne til mappen "Contacts". Arkiverne bl...` | Known | UI element |
| 0x0017B9D8 | `Alarmer` | Known | UI element |
| 0x0017C430 | `Nulstil menu` | Known | Menu item |
| 0x0017C5F4 | `Hovedmenu` | Known | Menu item |
| 0x0017C964 | `Menuer` | Known | Menu item |
| 0x0017EA20 | `Extras` | Known | UI element |
| 0x0017F508 | `ffnen Sie ihn durch Doppelklicken in das iPod Symbol un...` | Known | UI element |
| 0x0017F6FF | `nnen sich hier Textdateien anzeigen lassen, indem Sie d...` | Known | UI element |
| 0x0017FA5B | `hlen" den Alarm beenden` | Known | UI element |
| 0x001821D4 | `Calendario` | Known | UI element |
| 0x001821E0 | `Calendarios` | Known | UI element |
| 0x0018292C | `El iPod puede almacenar contactos y eventos de calendar...` | Known | UI element |
| 0x00182B0E | `gido y arrastrar los archivos vCard a la carpeta Contac...` | Known | UI element |
| 0x00182CA1 | `n de usar el iPod como disco y hacer doble clic en el i...` | Known | UI element |
| 0x00182DC4 | `Alarmas` | Known | UI element |
| 0x00182F29 | `gido y arrastrar los archivos de texto a la carpeta Not...` | Known | UI element |
| 0x00183280 | `Alarma` | Known | UI element |
| 0x00183860 | `Reloj con alarma` | Known | UI element |
| 0x00183A58 | `Contraste` | Known | UI element |
| 0x00183BC8 | `Hora alarma` | Known | UI element |
| 0x0018631E | ` sitten vCardit iPodin Contacts-kansioon. Tarkemmat tie...` | Known | UI element |
| 0x001864A0 | ` vCardit iPodin Contacts-kansioon. Ne haetaan automaatt...` | Known | UI element |
| 0x001866B2 | ` tekstitiedostot iPodin Notes-kansioon. Lis` | Known | UI element |
| 0x00189B34 | `Contacts` | Known | UI element |
| 0x0018A300 | `Votre iPod peut stocker des contacts et des ` | Known | UI element |
| 0x0018A39C | `lectionnez Appareils > Ajouter un appareil. Puis choisi...` | Known | UI element |
| 0x0018A434 | `adresses, Microsoft Entourage ou Palm Desktop et export...` | Known | UI element |
| 0x0018A5AC | ` mille contacts en plus de votre musique. Les applicati...` | Known | UI element |
| 0x0018A7D8 | `Alarmes` | Known | UI element |
| 0x0018A934 | `utilisation du disque, puis glissez ces fichiers dans l...` | Known | UI element |
| 0x0018AC6C | `Chargement des notes.` | Known | UI element |
| 0x0018ACA8 | `Alarme` | Known | UI element |
| 0x0018ACB0 | `Chargement des contacts.` | Known | UI element |
| 0x0018B32B | `init. menu p.` | Known | Menu item |
| 0x0018B520 | `Menu principal` | Known | Menu item |
| 0x0018B5E8 | `H. alarme` | Known | UI element |
| 0x0018BB2C | `Menu princ.` | Known | Menu item |
| 0x0018D4A4 | `Calendari` | Known | UI element |
| 0x0018DBDD | ` archiviare contatti ed eventi di calendari. Se utilizz...` | Known | UI element |
| 0x0018EA2C | `Ripr. Menu Princ.` | Known | Menu item |
| 0x0018EB64 | `Contrasto` | Known | UI element |
| 0x0019584F | ` Contacts ` | Known | UI element |
| 0x00195A31 | ` "Contacts" ` | Known | UI element |
| 0x00195C90 | ` Notes ` | Known | UI element |
| 0x00198DE4 | `Shuffle nummers` | Known | UI element |
| 0x001996F0 | `De iPod biedt ruimte voor maar liefst duizend adressen ...` | Known | UI element |
| 0x00199A1C | `Om tekstbestanden te bekijken, stelt u de iPod in als h...` | Known | UI element |
| 0x0019A420 | `Herstel menu` | Known | Menu item |
| 0x0019A564 | `Contrast` | Known | UI element |
| 0x0019A5F4 | `Shuffle` | Known | UI element |
| 0x0019A5FC | `Hoofdmenu` | Known | Menu item |
| 0x0019A968 | `Menu's` | Known | Menu item |
| 0x0019CF30 | ` iPod-symbolet, og flytt vCard-filene inn i Contacts-ma...` | Known | UI element |
| 0x0019DCC4 | `Alarmtidspunkt` | Known | UI element |
| 0x001A089C | `refter drar du in vCard-filerna i mappen "Contacts" i i...` | Known | UI element |
| 0x001A0A4B | ` skrivbordet och drar in vCard-filerna i mappen "Contac...` | Known | UI element |
| 0x001A0C7D | `ge och drar sedan in textfilerna i mappen "Notes" p` | Known | UI element |
| 0x001A1874 | `Alarmtid` | Known | UI element |
| 0x001B7410 | `Now Playing` | Known | UI element |
| 0x001B7488 | `Calendar` | Known | UI element |
| 0x001B7494 | `Calendars` | Known | UI element |
| 0x001B74A8 | `Backlight` | Known | UI element |
| 0x001B7500 | `Shuffle Songs` | Known | UI element |
| 0x001B8334 | `Alarms` | Known | UI element |
| 0x001B8700 | `Notes loading.` | Known | UI element |
| 0x001B876C | `Contacts loading.` | Known | UI element |
| 0x001B8C7C | `Sleep Timer` | Known | UI element |
| 0x001B8C88 | `Alarm Clock` | Known | UI element |
| 0x001B8D48 | `Reset Main Menu` | Known | Menu item |
| 0x001B8E4C | `Reset All Settings` | Known | User setting |
| 0x001B8EF8 | `Backlight Timer` | Known | UI element |
| 0x001B8F08 | `Repeat` | Known | UI element |
| 0x001B8F18 | `Main Menu` | Known | Menu item |
| 0x001B8FC0 | `Settings` | Known | User setting |
| 0x001B8FDC | `Alarm Time` | Known | UI element |
| 0x001B92E0 | `Reset All` | Known | UI element |
| 0x001E563D | `Contacts\` | Known | UI element |
| 0x001E5651 | `Calendars\` | Known | UI element |
| 0x001E5662 | `Notes\` | Known | UI element |
| 0x001E56A8 | `vcalendar` | Known | UI element |
| 0x001E575C | `dalarm` | Known | UI element |
| 0x001E58D4 | `valarm` | Known | UI element |
| 0x001E677C | `To check song links, set the preference NotesOnly to tr...` | Known | UI element |
| 0x001E68C8 | `The NotesOnly pref can only be set globally in the Pref...` | Known | UI element |
| 0x001E69A4 | `Warning: Preferences file must be in the Notes folder, ...` | Known | UI element |
| 0x001E6DB0 | `Memory full. %d notes loaded, some notes not loaded. (3...` | Known | UI element |
| 0x00331606 | `Illegal instruction` | Known | UI element |
| 0x00331634 | `Illegal address` | Known | UI element |
| 0x0033173A | `NotesOnly` | Known | UI element |

---

## 4. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00173D44 | `RtcTaskClass` | Known | RTOS task thread |
| 0x00173D54 | `TimerTaskClass` | Known | RTOS task thread |
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
| 0x00174294 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x001790B4 | `LoadDataTasks` | Known | RTOS task thread |
| 0x00179368 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x00179483 | `5RunTestsTask` | Known | RTOS task thread |
| 0x001D55B4 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x001E5AA8 | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x001E5AC0 | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x0032CBE4 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0032CBF8 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0032CC0C | `SBP2CommandTask` | Known | RTOS task thread |

---

## 5. Audio/Codec

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00114C6C | `RIFFWAVEfmt data` | Known | PCM audio format |
| 0x00164634 | `RIFFWAVEfmt dataD` | Known | PCM audio format |
| 0x001742A4 | `pcmWrite.wav` | Known | PCM audio format |
| 0x0017BE18 | `Audible-softwaren i dette produkt bruges i henhold til ...` | Known | Audible audiobook format |
| 0x0017BE78 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x0017BF6A | `.net codec i dette produkt bruges i henhold til en lice...` | Known | Audio system |
| 0x0017C014 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunh...` | Known | Audio system |
| 0x0017FA8C | `Die Audible Software in diesem Produkt wird in Lizenz v...` | Known | Audible audiobook format |
| 0x0017FAE5 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x0017FBD5 | `.net Codec in diesem Produkt wird mit der Lizenz der Vo...` | Known | Audio system |
| 0x0017FC9B | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und ...` | Known | Audio system |
| 0x0018329C | `El software Audible incluido en este producto se usa ba...` | Known | Audible audiobook format |
| 0x001832F7 | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x00183499 | `n de audio MPEG Layer-3 utilizada bajo licencia de Frau...` | Known | Audio system |
| 0x001869B6 | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright...` | Known | Audible audiobook format |
| 0x001869F0 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x00186AC2 | `.net codec t` | Known | Audio system |
| 0x00186B58 | `MPEG Layer-3 -` | Known | Audio system |
| 0x00186B6A | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | Audio system |
| 0x0018ACCC | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x0018AD16 | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x0018AD2B | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x0018ADDC | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x0018AEB0 | `La technologie de codage audio MPEG Layer 3 est utilis` | Known | Audio system |
| 0x0018AEE8 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | Audio system |
| 0x0018E490 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x0018E4B9 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0018E4E8 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x0018E55A | ` utilizzato su licenza da VoiceAge Corporation. Il code...` | Known | Audio system |
| 0x0018E630 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da...` | Known | Audio system |
| 0x00192295 | ` Audible ` | Known | Audible audiobook format |
| 0x001922B6 | `Audible ` | Known | Audible audiobook format |
| 0x0019230F | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x001924C4 | `MPEG Layer-3 ` | Known | Audio system |
| 0x00192510 | `Fraunhofer IIS ` | Known | Audio system |
| 0x0019604E | ` Audible` | Known | Audible audiobook format |
| 0x00196092 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x00196182 | `.net codec` | Known | Audio system |
| 0x00196243 | ` Fraunhofer IIS` | Known | Audio system |
| 0x00199DCC | `De Audible-software in dit product wordt gebruikt in li...` | Known | Audible audiobook format |
| 0x00199E23 | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x00199F14 | `.net-codec in dit product wordt gebruikt in licentie va...` | Known | Audio system |
| 0x00199FB0 | `Technologie voor codering van MPEG Layer-3-audio in lic...` | Known | Audio system |
| 0x0019D470 | `Audible-programvaren i dette produktet brukes under lis...` | Known | Audible audiobook format |
| 0x0019D4C4 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0019D640 | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fr...` | Known | Audio system |
| 0x001A0F8C | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x001A0FBB | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x001A0FD2 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x001A116C | `Ljudkodningstekniken MPEG Layer-3 ` | Known | Audio system |
| 0x001A11A2 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | Audio system |
| 0x001A4A0E | ` Fraunhofer IIS ` | Known | Audio system |
| 0x001B8780 | `The Audible software in this product is used under lice...` | Known | Audible audiobook format |
| 0x001B88B9 | `.net codec in this product is used under license from V...` | Known | Audio system |
| 0x001B894C | `MPEG Layer-3 audio coding technology licensed from Frau...` | Known | Audio system |
| 0x001E65DD | `&Aacute` | Known | AAC codec |
| 0x001E66A3 | `&aacute` | Known | AAC codec |
| 0x001E6A93 | `Boot time (MP3PlayerExampleApp constructor)` | Known | MP3 codec |
| 0x00326EF1 | ``0aLaCfDf` | Known | Apple Lossless codec |
| 0x0032C23C | `Audible` | Known | Audible audiobook format |
| 0x0032C314 | `AudioCodecs` | Known | Audio system |
| 0x0032C4A4 | `mp4_aacdec_sync` | Known | AAC codec |
| 0x0032C4B4 | `mp3dec_sync` | Known | MP3 codec |
| 0x0032E3F7 | `msCodeCom` | Known | Audio system |
| 0x0032F177 | `aaControls` | Known | AAC codec |

---

## 6. Storage/Hardware

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00003A00 | `!ATAsoso` | Known | ATA/disk interface |
| 0x00003A28 | `!ATAdpua` | Known | ATA/disk interface |
| 0x0004E894 | `atad8@-` | Known | ATA/disk interface |
| 0x00098560 | `atadmrts\|@-` | Known | ATA/disk interface |
| 0x000B98F0 | `atadmhdp` | Known | ATA/disk interface |
| 0x000BB24C | `atadmhbddbhmmhsd>@-` | Known | ATA/disk interface |
| 0x000BB73C | `atadmhpo` | Known | ATA/disk interface |
| 0x000E93C4 | `nutiatad` | Known | ATA/disk interface |
| 0x00173C1C | `data abort` | Known | ATA/disk interface |
| 0x00173DDC | `FirewireInitiator` | Known | FireWire interface |
| 0x00173DFC | `FirewireHandler` | Known | FireWire interface |
| 0x00173F41 | `diskmode` | Known | Hardware interface |
| 0x00173F4A | `diskscan` | Known | Hardware interface |
| 0x00174114 | `diskModeImageRev` | Known | Hardware interface |
| 0x001741D4 | `FirewireGuid` | Known | FireWire interface |
| 0x0017AE38 | `Spiller nu` | Known | Hardware interface |
| 0x0017AEF0 | `Spillelister` | Known | Hardware interface |
| 0x0017AF0C | `Genoptag spil` | Known | Hardware interface |
| 0x0017B3B0 | `Slet spilleliste` | Known | Hardware interface |
| 0x0017B3C4 | `Arkiver spilleliste` | Known | Hardware interface |
| 0x0017B464 | `Ny spilleliste %lu` | Known | Hardware interface |
| 0x0017B4F4 | `Harddisk` | Known | Hardware interface |
| 0x0017B762 | ` den kan bruges som FireWire-disk, og tr` | Known | FireWire interface |
| 0x0017B8CB | `r du har tilsluttet iPod som disk, skal du dobbeltklikk...` | Known | Hardware interface |
| 0x0017BAFA | ` den kan bruges som disk og anbringe tekstarkiver i map...` | Known | Hardware interface |
| 0x0017C27B | `je den til spillelisten On-The-Go. Hold knappen nede, n` | Known | Hardware interface |
| 0x0017C2B4 | `r en spilleliste, kunstner eller et album er valgt for ...` | Known | Hardware interface |
| 0x0017C2F1 | `je alle sangene til spillelisten On-The-Go.` | Known | Hardware interface |
| 0x0017C3C0 | `Nyt spil` | Known | Hardware interface |
| 0x0017C510 | ` Spillelister` | Known | Hardware interface |
| 0x0017CA9C | `FireWire tilsluttet` | Known | FireWire interface |
| 0x0017E9D4 | `Spiele` | Known | Hardware interface |
| 0x0017EA48 | `Weiterspielen` | Known | Hardware interface |
| 0x0017F2C7 | `ffnen Sie das Adressbuch, Microsoft Entourage oder Palm...` | Known | FireWire interface |
| 0x00180058 | `Neues Spiel` | Known | Hardware interface |
| 0x001807AE | `ber FireWire verbunden` | Known | FireWire interface |
| 0x00183FF4 | `FireWire conectado` | Known | FireWire interface |
| 0x00185F3C | `Diskanttivahv.` | Known | Hardware interface |
| 0x00185F4C | `Diskanttiheikenn.` | Known | Hardware interface |
| 0x00186010 | `Ladataan` | Known | ATA/disk interface |
| 0x00186306 | ` FireWire-levyn` | Known | FireWire interface |
| 0x0018694C | `Muistiinpanoja ladataan.` | Known | ATA/disk interface |
| 0x00186988 | `Yhteystietoja ladataan.` | Known | ATA/disk interface |
| 0x00187410 | `nityksen jatkamiseen ei ole tarpeeksi vapaata levytilaa...` | Known | ATA/disk interface |
| 0x00187450 | `nityksen aloittamiseen ei ole tarpeeksi vapaata levytil...` | Known | ATA/disk interface |
| 0x00187694 | `FireWire liitetty` | Known | FireWire interface |
| 0x0018A4A9 | `utiliser comme disque FireWire. Puis glissez les vCards...` | Known | FireWire interface |
| 0x0018BA54 | `FireWire Connect` | Known | FireWire interface |
| 0x0018E928 | `Data & Ora` | Known | ATA/disk interface |
| 0x0018E9F0 | `Imposta Data & Ora` | Known | ATA/disk interface |
| 0x0018F0FC | `FireWire Connesso` | Known | FireWire interface |
| 0x001932E0 | `FireWire ` | Known | FireWire interface |
| 0x00199498 | `Op de iPod kunt u adres- en agendagegevens opslaan. Als...` | Known | FireWire interface |
| 0x0019AAA8 | `FireWire aangesloten` | Known | FireWire interface |
| 0x0019C44C | `Spilles n` | Known | Hardware interface |
| 0x0019C518 | `Fortsett spill` | Known | Hardware interface |
| 0x0019C974 | `Mer diskant` | Known | Hardware interface |
| 0x0019C980 | `Mindre diskant` | Known | Hardware interface |
| 0x0019C9C4 | `Slett spilleliste` | Known | Hardware interface |
| 0x0019CB18 | `Diskmodus` | Known | Hardware interface |
| 0x0019CD0C | `pner du Adressebok, Microsoft Entourage eller Palm Desk...` | Known | Hardware interface |
| 0x0019CE64 | `ringer i tillegg til musikken din. Microsoft Outlook, M...` | Known | Hardware interface |
| 0x0019D110 | `Hvis du vil vise tekstfiler her, aktiverer du iPod for ...` | Known | Hardware interface |
| 0x0019D88F | ` legge den til i On-The-Go-spillelisten. Du kan legge t...` | Known | Hardware interface |
| 0x0019D9C4 | `Nytt spill` | Known | Hardware interface |
| 0x0019DE3C | `Det er ikke nok ledig diskplass til ` | Known | Hardware interface |
| 0x0019E09C | `Koblet til via FireWire` | Known | FireWire interface |
| 0x001A086C | `ll sedan in din iPod som FireWire-h` | Known | FireWire interface |
| 0x001A0891 | `rddisk. D` | Known | Hardware interface |
| 0x001A09F2 | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x001A0A12 | `rddisk. Sedan dubbelklickar du bara p` | Known | Hardware interface |
| 0x001A0C74 | `rddiskl` | Known | Hardware interface |
| 0x001A0D14 | `inget kort inmatat` | Known | ATA/disk interface |
| 0x001A1C78 | `FireWire anslutet` | Known | FireWire interface |
| 0x001B7EC0 | `Disk Mode` | Known | Hardware interface |
| 0x001B7F94 | `Your iPod can store contacts and calendar events. If yo...` | Known | Hardware interface |
| 0x001B815C | `Your iPod can store up to one thousand contacts right a...` | Known | Hardware interface |
| 0x001B8420 | `To view text files here, enable iPod for disk use, then...` | Known | Hardware interface |
| 0x001B9140 | `There is not enough free disk space to continue recordi...` | Known | Hardware interface |
| 0x001B917C | `There is not enough free disk space to start recording.` | Known | Hardware interface |
| 0x001B93F4 | `FireWire Connected` | Known | FireWire interface |
| 0x001B9408 | `No battery power remains. Please connect iPod to power.` | Known | Power management |
| 0x001B9440 | `Low Battery` | Known | Power management |
| 0x001D564C | `TROMResourceDB - unknown header version! (Try regenerat...` | Known | ATA/disk interface |
| 0x001E5EE4 | `Bad Data` | Known | ATA/disk interface |
| 0x001E6890 | `Bad meta data, name not a recognized preference. (27)` | Known | ATA/disk interface |
| 0x001E6948 | `Bad meta data, name not found. (23)` | Known | ATA/disk interface |
| 0x001E696C | `Bad meta data, name termination quote not found. (24)` | Known | ATA/disk interface |
| 0x001E69F4 | `Bad meta data, content not found. (25)` | Known | ATA/disk interface |
| 0x001E6A1C | `Bad meta data, content termination quote not found. (26...` | Known | ATA/disk interface |
| 0x001E6B48 | `Bad meta data, encoding not found. (8)` | Known | ATA/disk interface |
| 0x001E6B70 | `Bad meta data, encoding termination quote not found. (9...` | Known | ATA/disk interface |
| 0x001E6BAC | `Bad meta data, unknown encoding. (17)` | Known | ATA/disk interface |
| 0x001E6D68 | `Bad data. (32)` | Known | ATA/disk interface |
| 0x001E6EFC | `USB MSC` | Known | USB interface |
| 0x0031BD13 | `ataTaza[aea;ajaaaVa)b'b+b+dMd[d]dtdvdrdsd}dudfd` | Known | ATA/disk interface |
| 0x003208FD | `aGa>a(a'aJa?a<a,a4a=aBaDasawaXaYaZakataoaeaqa_a]aSaua` | Known | ATA/disk interface |
| 0x0032C2F8 | `MaximumDataRate` | Known | ATA/disk interface |
| 0x0032C364 | `FireWire` | Known | FireWire interface |
| 0x0032C3D4 | `FireWireVersion` | Known | FireWire interface |
| 0x0032C5A4 | `MEMDISK` | Known | Hardware interface |
| 0x0032C63A | `ex_data` | Known | ATA/disk interface |
| 0x0032C80C | `c:\buildtools\MWSF2\Q22Firmware.proj\projectfiles\sandb...` | Known | ATA/disk interface |
| 0x0032D84B | `RSA Data Security, Inc.` | Known | ATA/disk interface |
| 0x0032D868 | `RSA Data Security, Inc. PKCS` | Known | ATA/disk interface |
| 0x0032D9C9 | `pkcs7-data` | Known | ATA/disk interface |
| 0x0032D9D4 | `pkcs7-signedData` | Known | ATA/disk interface |
| 0x0032D9E5 | `pkcs7-envelopedData` | Known | ATA/disk interface |
| 0x0032D9F9 | `pkcs7-signedAndEnvelopedData` | Known | ATA/disk interface |
| 0x0032DA16 | `pkcs7-digestData` | Known | ATA/disk interface |
| 0x0032DA27 | `pkcs7-encryptedData` | Known | ATA/disk interface |
| 0x0032DC80 | `nsDataType` | Known | ATA/disk interface |
| 0x0032DC8B | `Netscape Data Type` | Known | ATA/disk interface |
| 0x0032EA25 | `id-smime-ct-authData` | Known | ATA/disk interface |
| 0x0032EA92 | `id-smime-ct-DVCSRequestData` | Known | ATA/disk interface |
| 0x0032EAAE | `id-smime-ct-DVCSResponseData` | Known | ATA/disk interface |
| 0x0032F4C6 | `id-cmc-dataReturn` | Known | ATA/disk interface |
| 0x0032F619 | `id-on-personalData` | Known | ATA/disk interface |
| 0x0032F71A | `id-cct-PKIData` | Known | ATA/disk interface |
| 0x0032FEB9 | `qualityLabelledData` | Known | ATA/disk interface |
| 0x00330294 | `setct-PANData` | Known | ATA/disk interface |
| 0x003302BF | `setct-OIData` | Known | ATA/disk interface |
| 0x003302D5 | `setct-PIData` | Known | ATA/disk interface |
| 0x003302E2 | `setct-PIDataUnsigned` | Known | ATA/disk interface |
| 0x0033035D | `setct-PInitResData` | Known | ATA/disk interface |
| 0x0033037D | `setct-PResData` | Known | ATA/disk interface |
| 0x003303D3 | `setct-CapTokenData` | Known | ATA/disk interface |
| 0x00330421 | `setct-AuthRevResData` | Known | ATA/disk interface |
| 0x0033046B | `setct-CapResData` | Known | ATA/disk interface |
| 0x003304A3 | `setct-CapRevResData` | Known | ATA/disk interface |
| 0x003304DA | `setct-CredResData` | Known | ATA/disk interface |
| 0x00330515 | `setct-CredRevResData` | Known | ATA/disk interface |
| 0x0033052A | `setct-PCertReqData` | Known | ATA/disk interface |
| 0x0033054F | `setct-BatchAdminReqData` | Known | ATA/disk interface |
| 0x00330567 | `setct-BatchAdminResData` | Known | ATA/disk interface |
| 0x003305BF | `setct-CertReqData` | Known | ATA/disk interface |
| 0x003305E2 | `setct-CertResData` | Known | ATA/disk interface |
| 0x0033096E | `setCext-merchData` | Known | ATA/disk interface |
| 0x003309F9 | `setCext-Track2Data` | Known | ATA/disk interface |
| 0x00330BE7 | `set-brand-IATA-ATA` | Known | ATA/disk interface |
| 0x00335D2F | `<key>FireWireGUID</key>` | Known | FireWire interface |

---

## 7. Error Messages

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00017740 | `*** assertion failed: %s, file %s, line %d` | Known | Error/assertion message |
| 0x0001DDA8 | `Invalid Operation` | Known | Error/assertion message |
| 0x00183184 | `Error durante la importaci` | Known | Error/assertion message |
| 0x0018E361 | ` verificato un errore durante l'importazione` | Known | Error/assertion message |
| 0x001B8644 | `An error occurred while importing` | Known | Error/assertion message |
| 0x001B9108 | `Cannot record because there is no microphone attached.` | Known | Error/assertion message |
| 0x001E67E0 | `Too many errors, further errors discarded. (31)` | Known | Error/assertion message |
| 0x001E6830 | `%s Error in file %s.` | Known | Error/assertion message |
| 0x001E6BF4 | `Cannot link to a .link file. (29)` | Known | Error/assertion message |
| 0x001E72E6 | `tInvalid partition table. Setup cannot continue.` | Known | Error/assertion message |
| 0x001E7317 | `Error loading operating system. Setup cannot continue.` | Known | Error/assertion message |
| 0x0032C7B4 | `%s(%d): OpenSSL internal error, assertion failed: %s` | Known | Error/assertion message |
| 0x0032D440 | `error:%08lX:%s:%s:%s` | Known | Error/assertion message |
| 0x0032D6E4 | `internal error: list index %ld out of range` | Known | Error/assertion message |

---

## 8. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00178BC0 | ` !"#$%&'()*+,-./0123456789:;<=>?@abcdefghijklmnopqrstuv...` | Known | Filesystem path |
| 0x00178CC0 | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x00178EC0 | `url;type=work:apple.com/support/ipod` | Known | Filesystem path |
| 0x0017BCCC | `Point: %d (%d/%d)` | Known | Filesystem path |
| 0x0017BEEB | `rende VoiceAge Corporation i USA og/eller andre lande o...` | Known | Filesystem path |
| 0x0017EB3C | `USA/Hawaii (NZ)` | Known | Filesystem path |
| 0x0017EB4C | `USA/Hawaii (SZ)` | Known | Filesystem path |
| 0x0017EB5C | `USA/Alaska (NZ)` | Known | Filesystem path |
| 0x0017EB6C | `USA/Alaska (SZ)` | Known | Filesystem path |
| 0x0017EB7C | `USA/Pazifik (NZ)` | Known | Filesystem path |
| 0x0017EB90 | `USA/Pazifik (SZ)` | Known | Filesystem path |
| 0x0017EBA4 | `USA/Rockies (NZ)` | Known | Filesystem path |
| 0x0017EBB8 | `USA/Rockies (SZ)` | Known | Filesystem path |
| 0x0017EBCC | `USA/Zentral (NZ)` | Known | Filesystem path |
| 0x0017EBE0 | `USA/Zentral (SZ)` | Known | Filesystem path |
| 0x0017EBF4 | `USA/Ost (NZ)` | Known | Filesystem path |
| 0x0017EC04 | `USA/Ost (SZ)` | Known | Filesystem path |
| 0x0017EF08 | `Vorn./Nachn.` | Known | Filesystem path |
| 0x0017EF18 | `Nachn./Vorn.` | Known | Filesystem path |
| 0x0017F930 | `Punkte: %d (%d/%d)` | Known | Filesystem path |
| 0x0017FB1F | ` ist entweder eine eingetragene Marke oder eine Marke d...` | Known | Filesystem path |
| 0x0018314C | `Result.: %d (%d/%d)` | Known | Filesystem path |
| 0x00183337 | ` es una marca registrada o una marca comercial de Voice...` | Known | Filesystem path |
| 0x00183F40 | `Fecha/hora` | Known | Filesystem path |
| 0x00186008 | `%d / %d` | Known | Filesystem path |
| 0x00186734 | `%s / %s` | Known | Filesystem path |
| 0x00186760 | `%d / %d valokuvaa tuotu` | Known | Filesystem path |
| 0x0018685C | `Tulos: %d (%d/%d)` | Known | Filesystem path |
| 0x00186A61 | `ity tavaramerkki Yhdysvalloissa ja/tai muissa maissa, j...` | Known | Filesystem path |
| 0x0018AB57 | `sult. : %d (%d/%d)` | Known | Filesystem path |
| 0x0018ADAE | `tats-Unis et/ou dans d` | Known | Filesystem path |
| 0x0018B2E7 | `gler date/heure` | Known | Filesystem path |
| 0x0018E328 | `Punti: %d (%d/%d)` | Known | Filesystem path |
| 0x00190A78 | `%b/%-d %-I:%M %2p` | Known | Filesystem path |
| 0x00190A8C | `%-m/%-d` | Known | Filesystem path |
| 0x00190AB0 | `%y/%-m/%d` | Known | Filesystem path |
| 0x00190ABC | `%Y/%b/%-d` | Known | Filesystem path |
| 0x001920C4 | ` %d (%d/%d)` | Known | Filesystem path |
| 0x00194CD0 | `%Y/%B/%d %A  %I:%M:%S %p` | Known | Filesystem path |
| 0x00194CEC | `%Y/%B/%d` | Known | Filesystem path |
| 0x00194D04 | `%-m/%-d %-I:%M %p` | Known | Filesystem path |
| 0x00194D38 | `%Y/%-m/%d` | Known | Filesystem path |
| 0x00195D64 | `%d / %d ` | Known | Filesystem path |
| 0x00195EBA | `: %d (%d/%d)` | Known | Filesystem path |
| 0x00199C8C | `Score: %d (%d/%d)` | Known | Filesystem path |
| 0x00199E5B | ` is een gedeponeerd handelsmerk of een handelsmerk van ...` | Known | Filesystem path |
| 0x0019A3E4 | `Stel datum/tijd in` | Known | Filesystem path |
| 0x0019D324 | `Poeng: %d (%d/%d)` | Known | Filesystem path |
| 0x0019D503 | ` er enten et registrert varemerke eller et varemerke fo...` | Known | Filesystem path |
| 0x0019FF14 | `%-d/%-m` | Known | Filesystem path |
| 0x001A0E48 | `ng: %d (%d/%d)` | Known | Filesystem path |
| 0x001A1058 | `r VoiceAge Corporation i USA och/eller andra l` | Known | Filesystem path |
| 0x001A1594 | `ll in datum/tid` | Known | Filesystem path |
| 0x001A45ED | ` %d/%d ` | Known | Filesystem path |
| 0x001A46F9 | `%d (%d/%d)` | Known | Filesystem path |
| 0x001B735C | `%-m/%d/%y` | Known | Filesystem path |
| 0x001B880B | ` is either registered trademark or trademark of VoiceAg...` | Known | Filesystem path |
| 0x001E5C69 | `file://` | Known | Filesystem path |
| 0x001E5C71 | `image://` | Known | Filesystem path |
| 0x001E5E50 | `</TITLE>` | Known | Filesystem path |
| 0x001E5E60 | `</BODY>` | Known | Filesystem path |
| 0x001E5E8E | `</ROT13>` | Known | Filesystem path |
| 0x001E6C38 | `Bad link, no matching </a> for anchor tag. (20)` | Known | Filesystem path |
| 0x001FEB67 | `W/}lE>q` | Known | Filesystem path |
| 0x0022D7A9 | `H."0*Bx/` | Known | Filesystem path |
| 0x00235106 | `U/~RERT` | Known | Filesystem path |
| 0x00239E6E | `TUOPT/\|` | Known | Filesystem path |
| 0x00241103 | `HuGZp/$j` | Known | Filesystem path |
| 0x00247613 | `(bJ)b"(b2""z/` | Known | Filesystem path |
| 0x00249D83 | `JUAPDD(/` | Known | Filesystem path |
| 0x00250092 | `/B\|$BD'` | Known | Filesystem path |
| 0x0025100F | `$Bd$BT/` | Known | Filesystem path |
| 0x002572E7 | `/" +J\|!` | Known | Filesystem path |
| 0x0025E166 | `Fb""")/` | Known | Filesystem path |
| 0x0025F31D | `/RyO(UIH` | Known | Filesystem path |
| 0x002603CD | `~$Bh'"~$Bz/` | Known | Filesystem path |
| 0x00263C37 | `$B +BZ/` | Known | Filesystem path |
| 0x0026B655 | `0c(HBP/` | Known | Filesystem path |
| 0x0026FBEB | `$B~("\|/` | Known | Filesystem path |
| 0x00286955 | `T/DDDDD` | Known | Filesystem path |
| 0x00286BBF | `"~UeB /` | Known | Filesystem path |
| 0x00289671 | `$B((B /` | Known | Filesystem path |
| 0x00291754 | ` "\|$B~/` | Known | Filesystem path |
| 0x002948D0 | `@$B\|$"(/` | Known | Filesystem path |
| 0x002959D0 | `)"8/B""` | Known | Filesystem path |
| 0x00296118 | `r4c6 bN/` | Known | Filesystem path |
| 0x0029B519 | `RDT%B(/` | Known | Filesystem path |
| 0x0029C6A5 | `RBHUE\|/` | Known | Filesystem path |
| 0x002A3D25 | `]B""B</` | Known | Filesystem path |
| 0x002A75A2 | `,B\|RED/` | Known | Filesystem path |
| 0x002AD025 | `$BT). /` | Known | Filesystem path |
| 0x002AE225 | `#"TUB(/` | Known | Filesystem path |
| 0x002BFA99 | `/" %BD"` | Known | Filesystem path |
| 0x002C63BC | `ODD""(/` | Known | Filesystem path |
| 0x002C75F7 | `B"$R%"B$" /` | Known | Filesystem path |
| 0x002C7F64 | `bG\|jG\|/` | Known | Filesystem path |
| 0x002C9B4E | `$E$$BR/` | Known | Filesystem path |
| 0x002C9BEF | `dRB~RA$/` | Known | Filesystem path |
| 0x002CA540 | `TT&T%B(/` | Known | Filesystem path |
| 0x002D9234 | `)'>$B8/` | Known | Filesystem path |
| 0x002DB842 | `$B\|%EV/` | Known | Filesystem path |
| 0x002E2D96 | `BDU!BJ ""/` | Known | Filesystem path |
| 0x002E3CFA | `Z-bD("(%B>/` | Known | Filesystem path |
| 0x002EC449 | `@(/ Q\|f` | Known | Filesystem path |
| 0x0031365D | `% %!%"%#%$%%%&%'%(%)%*%+%,%-%.%/%0%1%2%3%4%5%6%7%8%9%:%...` | Known | Filesystem path |
| 0x00313E3D | `qWlIl/Ymg*` | Known | Filesystem path |
| 0x003156AB | `X)W,W*W3W9W.W/W\W;WBWiW` | Known | Filesystem path |
| 0x00315D79 | `n/o6oKoto*o` | Known | Filesystem path |
| 0x00316311 | `q/q1qsq\qhqEqrqJqxqzq` | Known | Filesystem path |
| 0x00317599 | `b6bKbNb/e` | Known | Filesystem path |
| 0x0031774F | `V3W0W(W-W,W/W)W` | Known | Filesystem path |
| 0x00317823 | `NMOOOGOWO^O4O[OUO0OPOQO=O:O8OCOTO<OFOcO\O`O/ONO6OYO]OHO...` | Known | Filesystem path |
| 0x00317E81 | `e%f-f f'f/f` | Known | Filesystem path |
| 0x00318183 | `bNc>c/cUcBcFcOcIc:cPc=c*c+c(cMcLcHeIe` | Known | Filesystem path |
| 0x00318435 | `fFUjUfUDU^UaUCUJU1UVUOUUU/UdU8U.U\U,UcU3UAUWU` | Known | Filesystem path |
| 0x0031869B | `\|F}C}q}.}9}<}@}0}3}D}/}B}2}1}=` | Known | Filesystem path |
| 0x00318823 | `W/X*X4X$X0X1X!X` | Known | Filesystem path |
| 0x00318939 | `k.l/l,l/n8nTn!n2ngnJn n%n#n` | Known | Filesystem path |
| 0x0031899B | `r6s%s4s)s:t*t3t"t%t5t6t4t/t` | Known | Filesystem path |
| 0x00319219 | `S.V;V9V2V?V4V)VSVNVWVtV6V/V0V` | Known | Filesystem path |
| 0x0031985F | `\|.~>~F~7~2~C~+~=~1~E~A~4~9~H~5~?~/~D` | Known | Filesystem path |
| 0x0031A84F | `P P'P5P/P1P` | Known | Filesystem path |
| 0x0031A9A1 | `h5h+h-h/hNhDh4h` | Known | Filesystem path |
| 0x0031A9B5 | `h&h(h.hMh:h%h h,k/k-k1k4kmk` | Known | Filesystem path |
| 0x0031AAC1 | `w"w'w#x,x"x5x/x(x.x+x!x)x3x*x1xTy[yOy\ySyRyQy` | Known | Filesystem path |
| 0x0031AF4B | `v9w/w-w1w2w4w3w=w%w;w5wHxRxIxMxJxLx&xExPxdygyiyjycykyay` | Known | Filesystem path |
| 0x0031AFA7 | `{1{+{-{/{2{8{` | Known | Filesystem path |
| 0x0031B43B | `t/uoulu` | Known | Filesystem path |
| 0x0031B7A1 | ``IaJa+aEa6a2a.aFa/aOa)a@a bh` | Known | Filesystem path |
| 0x0031B7E7 | `d&d0d(dAd5d/d` | Known | Filesystem path |
| 0x0031B84F | `hxi4iii@ioiDiviXiAitiLi;iKi7i\iOiQi2iRi/i{i<iFkEkCkBkHk...` | Known | Filesystem path |
| 0x0031BEB9 | `y+zJz0z/z(z&z` | Known | Filesystem path |
| 0x0031C79D | `u/v-v1v=v3v<v5v2v0v` | Known | Filesystem path |
| 0x0031CA93 | `X-[%[2[#[,['[&[/[.[{[` | Known | Filesystem path |
| 0x0031D3C1 | `j<p5p/p7p4p1pBp8p?p:p9p@p;p3pAp` | Known | Filesystem path |
| 0x0031D5C3 | `w-y1y/yT\|S\|` | Known | Filesystem path |
| 0x0031E98F | `%#%3%+%;%K% %/%(%7%?%` | Known | Filesystem path |
| 0x0031EB0E | `02*2+2,2-2.2/2@272B2C292:212>24222;2623252<2=2?282` | Known | Filesystem path |
| 0x0031F039 | `\7_J_/`P`m`` | Known | Filesystem path |
| 0x0031F215 | `OHSIT>T/Z` | Known | Filesystem path |
| 0x0031F28B | `i_l*mim/n` | Known | Filesystem path |
| 0x0031F77D | `N,p]u/f` | Known | Filesystem path |
| 0x00320489 | `S#S/S1S3S8S@SFSES` | Known | Filesystem path |
| 0x00320589 | `q4V6V2V8VkVdV/VlVjV` | Known | Filesystem path |
| 0x00320831 | `_)_-_8_A_H_L_N_/_Q_V_W_Y_a_m_s_w_` | Known | Filesystem path |
| 0x00320EF7 | `s4s/s)s%s>sNsOs` | Known | Filesystem path |
| 0x00324CFB | `h>kLp/t` | Known | Filesystem path |
| 0x00325385 | `o;v/}7~` | Known | Filesystem path |
| 0x00326149 | `e1f/h\q6z` | Known | Filesystem path |
| 0x00326795 | `UuX/c"dIfKfmh` | Known | Filesystem path |
| 0x0032C138 | `%s<key>%s</key>` | Known | Filesystem path |
| 0x0032C148 | `%s<%s/>` | Known | Filesystem path |
| 0x0032C170 | `%s</dict>` | Known | Filesystem path |
| 0x0032C18C | `%s<string>%s</string>` | Known | Filesystem path |
| 0x0032C1B4 | `%s<integer>%s</integer>` | Known | Filesystem path |
| 0x0032C1E0 | `%s<integer>%d</integer>` | Known | Filesystem path |
| 0x0032C544 | ` ,;=+[]*?<>\|":/\` | Known | Filesystem path |
| 0x0032C555 | `\/:*?"<>\|` | Known | Filesystem path |
| 0x0032D578 | `You need to read the OpenSSL FAQ, http://www.openssl.or...` | Known | Filesystem path |
| 0x0032E745 | `S/MIME Capabilities` | Known | Filesystem path |
| 0x00335C90 | `<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1....` | Known | Filesystem path |
| 0x00335D1E | `</dict>` | Known | Filesystem path |
| 0x00335D26 | `</plist>` | Known | Filesystem path |
| 0x00335D47 | `<string>%08lX%08lX</string>` | Known | Filesystem path |
| 0x003361A1 | `]7/[P{!:` | Known | Filesystem path |
| 0x003374DE | `/ 0 1 2 3 4 5 6 7 8 9 : ; < = > ? @ A B C D E F G H I J...` | Known | Filesystem path |
| 0x003376BF | `! !!!"!#!$!%!&!'!(!)!*!+!,!-!.!/!0!1!2!3!4!5!6!7!8!9!:!...` | Known | Filesystem path |
| 0x0034E508 | `/1f;{1Q` | Known | Filesystem path |
| 0x0035BD9F | `^r+@2Y/Fo/` | Known | Filesystem path |
| 0x0035D43C | `l</gT;T` | Known | Filesystem path |
| 0x0035D80F | `;5_b}B/` | Known | Filesystem path |
| 0x003630A7 | `2//FOH!` | Known | Filesystem path |
| 0x0036777F | `zx/W4<,` | Known | Filesystem path |
| 0x0036F0BA | `q;UT/5Z` | Known | Filesystem path |
| 0x00371D3C | `/f?=}9A` | Known | Filesystem path |
| 0x0037624D | `{m:/C.#.` | Known | Filesystem path |
| 0x00378DA6 | `txhQG[/` | Known | Filesystem path |
| 0x00379465 | `*a%mT/!` | Known | Filesystem path |
| 0x0037AE4B | `/KdeD6<` | Known | Filesystem path |
| 0x00389E88 | `KS/W=Iv((` | Known | Filesystem path |
| 0x0038D02E | `C/X!6u\|` | Known | Filesystem path |
| 0x0038DD01 | `u/R9 aW` | Known | Filesystem path |
| 0x0038E44C | `ie/V>N D` | Known | Filesystem path |
| 0x00395535 | `wJl#L/ROdF` | Known | Filesystem path |
| 0x0039F1A8 | `/3}6*DP` | Known | Filesystem path |
| 0x003A184C | `i>o>7Y/` | Known | Filesystem path |
| 0x003ADB13 | ``1T/u1"` | Known | Filesystem path |
| 0x003B1C85 | `p<DW/7k^` | Known | Filesystem path |
| 0x003B2752 | `&/Dn}\$` | Known | Filesystem path |
| 0x003B62F2 | `n=uL/gZ` | Known | Filesystem path |
| 0x003BAB1D | `,`i; ;/` | Known | Filesystem path |
| 0x003C6AAA | `~n%.tT/` | Known | Filesystem path |
| 0x003C7436 | `}hVgV\M/` | Known | Filesystem path |
| 0x003D0C1B | `+2U"0o/K` | Known | Filesystem path |
| 0x003D2040 | `- oPlgl/` | Known | Filesystem path |
| 0x003DB13B | `*6E/k}O` | Known | Filesystem path |
| 0x003E2663 | `UE/\nxM` | Known | Filesystem path |
| 0x003E6097 | `/XEf{/l`` | Known | Filesystem path |
| 0x003E638F | `kX+f/@1P` | Known | Filesystem path |
| 0x003EB3B1 | `1^{/S\|^Q@` | Known | Filesystem path |
| 0x003EB753 | `TNZFe/q` | Known | Filesystem path |
| 0x003F1613 | `B/P]@&u` | Known | Filesystem path |
| 0x003F2997 | `/n,5!QI+` | Known | Filesystem path |
| 0x00408063 | `N/ecSB\` | Known | Filesystem path |
| 0x0040847E | `!rd}aj8/[&` | Known | Filesystem path |
| 0x00410DAA | `M/k,zE8` | Known | Filesystem path |
| 0x00412993 | `%?]7/G+` | Known | Filesystem path |
| 0x0041365E | `QC/Lcxr` | Known | Filesystem path |
| 0x00417D09 | `4jTrD/a7` | Known | Filesystem path |
| 0x004180ED | `h/SzH.v` | Known | Filesystem path |
| 0x00422BE6 | `oQ1`}`/` | Known | Filesystem path |
| 0x00423A9D | `{_JT/_zP` | Known | Filesystem path |
| 0x00429282 | `mM\G/3=` | Known | Filesystem path |
| 0x0042D405 | `;DUR:y/~>\` | Known | Filesystem path |
| 0x0043074E | `5/rs9QI` | Known | Filesystem path |
| 0x004321E9 | `#Mwg/N:7` | Known | Filesystem path |
| 0x00435D7F | `u3/eS>=J;"` | Known | Filesystem path |
| 0x0043656C | `'`F1hD/` | Known | Filesystem path |
| 0x0043A47B | `t:_1$/k+` | Known | Filesystem path |
| 0x0044124D | `#8/wk8l~P` | Known | Filesystem path |
| 0x0044146C | `@~/#k>h` | Known | Filesystem path |
| 0x00447692 | `E;Iq/d\|` | Known | Filesystem path |
| 0x00448B46 | `7,dM;.I/` | Known | Filesystem path |
| 0x0044BFE2 | `mX/)?Q'` | Known | Filesystem path |
| 0x004539FE | `mk6]/HN)` | Known | Filesystem path |
| 0x004564C2 | `/\r)G>2^` | Known | Filesystem path |
| 0x004598BE | `!\|n<P"D0/` | Known | Filesystem path |
| 0x00461E99 | `+'HjldU./` | Known | Filesystem path |

---

## 9. Binary Structure (for Ghidra/IDA)

| Parameter | Value |
|-----------|-------|
| **Architecture** | ARM (32-bit, little-endian) |
| **Processor** | ARM7TDMI (ARMv4T), dual-core |
| **Base Address** | 0x00000000 |
| **Entry Point** | 0x00000000 |
| **Endianness** | Little-Endian |
| **File Size** | 4,605,440 bytes |

