# iPod Mini 2nd Generation - RetailOS 7.1.4.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 7.1.4.1 |
| **IPSW** | iPod_7.1.4.1.ipsw |
| **Device** | iPod Mini 2nd Generation (2005, 4/6GB Microdrive) |
| **Binary Size** | 4,506,624 bytes (4.30 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 4,506,624 bytes |
| **Total Strings (>=6)** | 9,856 |
| **Function Prologues** | 7,561 |
| **SoC** | PortalPlayer PP5020 |
| **Architecture** | ARM7TDMI (ARMv4T), dual-core |
| **Encrypted** | No |
| **Decryption Method** | Extract from IPSW (unencrypted) |
| **SHA-256** | `56becf109e5233a4de6e774fb87aa24ef6e8de8f44f30938a0beb6ed486e00f4` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168608 | `MP3ExampleTask` | Hidden | Hidden Test |

---

## 2. Discovered Features

### EQ Preset

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
| 0x001722F4 | `USA/Rockies (NZ)` | EQ Preset | |
| 0x00172308 | `USA/Rockies (SZ)` | EQ Preset | |
| 0x0017550C | `Latina` | EQ Preset | |
| 0x0018B910 | `Latino` | EQ Preset | |
| 0x0018C63A | ` navigeringsflaten) for ` | EQ Preset | |
| 0x001CD6A5 | `LATIN-1` | EQ Preset | |
| 0x001CD6AD | `LATIN1` | EQ Preset | |
| 0x002C1473 | `~ BR&B$"` | EQ Preset | |
| 0x00317D6A | `Secure Electronic Transactions` | EQ Preset | |

### Localization

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016D668 | `x-mac-japanese` | Localization | |
| 0x001A352C | `English` | Localization | |
| 0x001A3564 | `Italiano` | Localization | |
| 0x001CD710 | `X-MAC-JAPANESE` | Localization | |
| 0x001CD71F | `MAC-JAPANESE` | Localization | |
| 0x001CD72C | `MACJAPANESE` | Localization | |
| 0x001CD758 | `X-MAC-CHINESETRAD` | Localization | |
| 0x001CD76A | `MAC-CHINESETRAD` | Localization | |
| 0x001CD787 | `X-MAC-CHINESESIMP` | Localization | |
| 0x001CD799 | `MAC-CHINESESIMP` | Localization | |
| 0x001CD7B9 | `X-MAC-KOREAN` | Localization | |
| 0x001CD7C6 | `MAC-KOREAN` | Localization | |

### Filesystem Path

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001684EC | `iPod_Control\Device` | Filesystem Path | |
| 0x00168500 | `iPod_Control` | Filesystem Path | |
| 0x00168510 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x00168C78 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x00168D10 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x00168D40 | `iPod_Control\Device\` | Filesystem Path | |
| 0x00168D9C | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x00168DE4 | `iPod_Control\Music\` | Filesystem Path | |

### Assertion

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00017BC0 | `*** assertion failed: %s, file %s, line %d` | Assertion | |
| 0x00314254 | `%s(%d): OpenSSL internal error, assertion failed: %s` | Assertion | |

---

## 3. Known User-Facing Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001685A8 | `KeyRepeatTimer` | Known | UI element |
| 0x0016F6B3 | `k dine vCards til mappen Contacts p` | Known | UI element |
| 0x0016F849 | `kke vCard-arkiverne til mappen "Contacts". Arkiverne bl...` | Known | UI element |
| 0x0016F900 | `Alarmer` | Known | UI element |
| 0x0017012C | `Nulstil hovedmenu` | Known | Menu item |
| 0x001702CC | `Shuffle` | Known | UI element |
| 0x001702D4 | `Hovedmenu` | Known | Menu item |
| 0x0017045C | `Menuer` | Known | Menu item |
| 0x00172170 | `Extras` | Known | UI element |
| 0x00172C58 | `ffnen Sie ihn durch Doppelklicken in das iPod Symbol un...` | Known | UI element |
| 0x00172DDF | `nnen sich hier Textdateien anzeigen lassen, indem Sie d...` | Known | UI element |
| 0x00172F37 | `hlen" den Alarm beenden` | Known | UI element |
| 0x00175088 | `Calendario` | Known | UI element |
| 0x00175094 | `Calendarios` | Known | UI element |
| 0x001757E0 | `El iPod puede almacenar contactos y eventos de calendar...` | Known | UI element |
| 0x001759C2 | `gido y arrastrar los archivos vCard a la carpeta Contac...` | Known | UI element |
| 0x00175B55 | `n de usar el iPod como disco y hacer doble clic en el i...` | Known | UI element |
| 0x00175C78 | `Alarmas` | Known | UI element |
| 0x00175D65 | `gido y arrastrar los archivos de texto a la carpeta Not...` | Known | UI element |
| 0x00175E90 | `Alarma` | Known | UI element |
| 0x00176470 | `Reloj con alarma` | Known | UI element |
| 0x0017661C | `Contraste` | Known | UI element |
| 0x0017678C | `Hora alarma` | Known | UI element |
| 0x00178902 | ` vCardit iPodin Contacts-kansioon. Lis` | Known | UI element |
| 0x00178A7A | ` vCardit Contacts-kansioon. T` | Known | UI element |
| 0x00178C12 | ` tekstitiedostot iPodin Notes-kansioon. Lis` | Known | UI element |
| 0x0017B8C8 | `Contacts` | Known | UI element |
| 0x0017C094 | `Votre iPod peut stocker des contacts et des ` | Known | UI element |
| 0x0017C130 | `lectionnez Appareils > Ajouter un appareil. Puis choisi...` | Known | UI element |
| 0x0017C1C8 | `adresses, Microsoft Entourage ou Palm Desktop et export...` | Known | UI element |
| 0x0017C340 | ` mille contacts en plus de votre musique. Les applicati...` | Known | UI element |
| 0x0017C56C | `Alarmes` | Known | UI element |
| 0x0017C63C | `utilisation du disque, puis glissez ces fichiers dans l...` | Known | UI element |
| 0x0017C738 | `Chargement des notes.` | Known | UI element |
| 0x0017C774 | `Alarme` | Known | UI element |
| 0x0017C77C | `Chargement des contacts.` | Known | UI element |
| 0x0017CDCB | `init. menu p.` | Known | Menu item |
| 0x0017CF9C | `Menu principal` | Known | Menu item |
| 0x0017D06C | `H. alarme` | Known | UI element |
| 0x0017D36C | `Menu princ.` | Known | Menu item |
| 0x0017E934 | `Calendari` | Known | UI element |
| 0x0017F06D | ` archiviare contatti ed eventi di calendari. Se utilizz...` | Known | UI element |
| 0x0017FC04 | `Ripr. Menu Princ.` | Known | Menu item |
| 0x0017FD18 | `Contrasto` | Known | UI element |
| 0x00185A2F | ` Contacts ` | Known | UI element |
| 0x00185C11 | ` "Contacts" ` | Known | UI element |
| 0x00185DE4 | ` Notes ` | Known | UI element |
| 0x0018867C | `Shuffle nummers` | Known | UI element |
| 0x00188F38 | `De iPod biedt ruimte voor maar liefst duizend adressen ...` | Known | UI element |
| 0x001891FC | `Om tekstbestanden te bekijken, stelt u de iPod in als h...` | Known | UI element |
| 0x001899E4 | `Herstel menu` | Known | Menu item |
| 0x00189A10 | `Shuffle nrs.` | Known | UI element |
| 0x00189B08 | `Contrast` | Known | UI element |
| 0x00189BA0 | `Hoofdmenu` | Known | Menu item |
| 0x00189D28 | `Menu's` | Known | Menu item |
| 0x0018BF20 | ` iPod-symbolet, og flytt vCard-filene inn i Contacts-ma...` | Known | UI element |
| 0x0018CA50 | `Alarmtidspunkt` | Known | UI element |
| 0x0018F0A3 | `refter drar du in vCard-filerna i mappen "Contacts" i i...` | Known | UI element |
| 0x0018F24F | ` skrivbordet och drar in vCard-filerna i mappen "Contac...` | Known | UI element |
| 0x0018F411 | `ge och drar sedan in textfilerna i mappen "Notes" p` | Known | UI element |
| 0x0018FDD4 | `Alarmtid` | Known | UI element |
| 0x001A33F8 | `Now Playing` | Known | UI element |
| 0x001A3454 | `Calendar` | Known | UI element |
| 0x001A3460 | `Calendars` | Known | UI element |
| 0x001A3474 | `Backlight` | Known | UI element |
| 0x001A34CC | `Shuffle Songs` | Known | UI element |
| 0x001A4300 | `Alarms` | Known | UI element |
| 0x001A4488 | `Notes loading.` | Known | UI element |
| 0x001A44F4 | `Contacts loading.` | Known | UI element |
| 0x001A4A04 | `Sleep Timer` | Known | UI element |
| 0x001A4A10 | `Alarm Clock` | Known | UI element |
| 0x001A4AB4 | `Reset Main Menu` | Known | Menu item |
| 0x001A4B98 | `Reset All Settings` | Known | User setting |
| 0x001A4C44 | `Backlight Timer` | Known | UI element |
| 0x001A4C54 | `Repeat` | Known | UI element |
| 0x001A4C64 | `Main Menu` | Known | Menu item |
| 0x001A4D0C | `Settings` | Known | User setting |
| 0x001A4D28 | `Alarm Time` | Known | UI element |
| 0x001A4E60 | `Reset All` | Known | UI element |
| 0x001CD275 | `Contacts\` | Known | UI element |
| 0x001CD289 | `Calendars\` | Known | UI element |
| 0x001CD29A | `Notes\` | Known | UI element |
| 0x001CD2E0 | `vcalendar` | Known | UI element |
| 0x001CD394 | `dalarm` | Known | UI element |
| 0x001CD50C | `valarm` | Known | UI element |
| 0x001CE2FC | `To check song links, set the preference NotesOnly to tr...` | Known | UI element |
| 0x001CE448 | `The NotesOnly pref can only be set globally in the Pref...` | Known | UI element |
| 0x001CE524 | `Warning: Preferences file must be in the Notes folder, ...` | Known | UI element |
| 0x001CE930 | `Memory full. %d notes loaded, some notes not loaded. (3...` | Known | UI element |
| 0x00319172 | `Illegal instruction` | Known | UI element |
| 0x003191A0 | `Illegal address` | Known | UI element |
| 0x003192A6 | `NotesOnly` | Known | UI element |

---

## 4. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168144 | `RtcTaskClass` | Known | RTOS task thread |
| 0x00168154 | `TimerTaskClass` | Known | RTOS task thread |
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
| 0x00168608 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x0016D428 | `LoadDataTasks` | Known | RTOS task thread |
| 0x0016D763 | `5RunTestsTask` | Known | RTOS task thread |
| 0x001BD464 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x00314684 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x00314698 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x003146AC | `SBP2CommandTask` | Known | RTOS task thread |

---

## 5. Audio/Codec

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0010BB68 | `RIFFWAVEfmt data` | Known | PCM audio format |
| 0x0015B0CC | `RIFFWAVEfmt dataD` | Known | PCM audio format |
| 0x00168618 | `pcmWrite.wav` | Known | PCM audio format |
| 0x0016FAE8 | `Audible-softwaren i dette produkt bruges i henhold til ...` | Known | Audible audiobook format |
| 0x0016FB48 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x0016FC3A | `.net codec i dette produkt bruges i henhold til en lice...` | Known | Audio system |
| 0x0016FCE4 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunh...` | Known | Audio system |
| 0x00172F68 | `Die Audible Software in diesem Produkt wird in Lizenz v...` | Known | Audible audiobook format |
| 0x00172FC1 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x001730B1 | `.net Codec in diesem Produkt wird mit der Lizenz der Vo...` | Known | Audio system |
| 0x00173177 | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und ...` | Known | Audio system |
| 0x00175EAC | `El software Audible incluido en este producto se usa ba...` | Known | Audible audiobook format |
| 0x00175F07 | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x001760A9 | `n de audio MPEG Layer-3 utilizada bajo licencia de Frau...` | Known | Audio system |
| 0x00178D2A | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright...` | Known | Audible audiobook format |
| 0x00178D64 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x00178E36 | `.net codec t` | Known | Audio system |
| 0x00178ECC | `MPEG Layer-3 -` | Known | Audio system |
| 0x00178EDE | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | Audio system |
| 0x0017C798 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x0017C7E2 | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x0017C7F7 | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x0017C8A8 | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x0017C97C | `La technologie de codage audio MPEG Layer 3 est utilis` | Known | Audio system |
| 0x0017C9B4 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | Audio system |
| 0x0017F68C | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x0017F6B5 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0017F6E4 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x0017F756 | ` utilizzato su licenza da VoiceAge Corporation. Il code...` | Known | Audio system |
| 0x0017F82C | `Tecnologia di codifica audio MPEG Layer-3 su licenza da...` | Known | Audio system |
| 0x00182B3D | ` Audible ` | Known | Audible audiobook format |
| 0x00182B5E | `Audible ` | Known | Audible audiobook format |
| 0x00182BB7 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x00182D6C | `MPEG Layer-3 ` | Known | Audio system |
| 0x00182DB8 | `Fraunhofer IIS ` | Known | Audio system |
| 0x00185F3A | ` Audible` | Known | Audible audiobook format |
| 0x00185F7E | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x0018606E | `.net codec` | Known | Audio system |
| 0x0018612F | ` Fraunhofer IIS` | Known | Audio system |
| 0x001893AC | `De Audible-software in dit product wordt gebruikt in li...` | Known | Audible audiobook format |
| 0x00189403 | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x001894F4 | `.net-codec in dit product wordt gebruikt in licentie va...` | Known | Audio system |
| 0x00189590 | `Technologie voor codering van MPEG Layer-3-audio in lic...` | Known | Audio system |
| 0x0018C210 | `Audible-programvaren i dette produktet brukes under lis...` | Known | Audible audiobook format |
| 0x0018C264 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0018C3E0 | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fr...` | Known | Audio system |
| 0x0018F538 | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x0018F567 | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x0018F57E | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x0018F718 | `Ljudkodningstekniken MPEG Layer-3 ` | Known | Audio system |
| 0x0018F74E | `n Fraunhofer IIS och THOMSON multimedia.` | Known | Audio system |
| 0x00192702 | ` Fraunhofer IIS ` | Known | Audio system |
| 0x001A4508 | `The Audible software in this product is used under lice...` | Known | Audible audiobook format |
| 0x001A4641 | `.net codec in this product is used under license from V...` | Known | Audio system |
| 0x001A46D4 | `MPEG Layer-3 audio coding technology licensed from Frau...` | Known | Audio system |
| 0x001CE15D | `&Aacute` | Known | AAC codec |
| 0x001CE223 | `&aacute` | Known | AAC codec |
| 0x001CE613 | `Boot time (MP3PlayerExampleApp constructor)` | Known | MP3 codec |
| 0x0030EA5D | ``0aLaCfDf` | Known | Apple Lossless codec |
| 0x00313CDC | `Audible` | Known | Audible audiobook format |
| 0x00313DB4 | `AudioCodecs` | Known | Audio system |
| 0x00313F44 | `mp4_aacdec_sync` | Known | AAC codec |
| 0x00313F54 | `mp3dec_sync` | Known | MP3 codec |
| 0x00315F63 | `msCodeCom` | Known | Audio system |
| 0x00316CE3 | `aaControls` | Known | AAC codec |

---

## 6. Storage/Hardware

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00003A00 | `!ATAsoso` | Known | ATA/disk interface |
| 0x00003A28 | `!ATAdpua` | Known | ATA/disk interface |
| 0x00005081 | `diskmode` | Known | Hardware interface |
| 0x0000508A | `diskscan` | Known | Hardware interface |
| 0x0004D37C | `atad8@-` | Known | ATA/disk interface |
| 0x000935EC | `atadmrts\|@-` | Known | ATA/disk interface |
| 0x000B35C4 | `atadmhdp` | Known | ATA/disk interface |
| 0x000B4D70 | `atadmhbddbhmmhsd>@-` | Known | ATA/disk interface |
| 0x000B5188 | `atadmhpo` | Known | ATA/disk interface |
| 0x000E1984 | `nutiatad` | Known | ATA/disk interface |
| 0x000E2728 | `atadtG` | Known | ATA/disk interface |
| 0x0016801C | `data abort` | Known | ATA/disk interface |
| 0x001681E8 | `FirewireHandler` | Known | FireWire interface |
| 0x00168488 | `diskModeImageRev` | Known | Hardware interface |
| 0x00168548 | `FirewireGuid` | Known | FireWire interface |
| 0x0016ED88 | `Spiller nu` | Known | Hardware interface |
| 0x0016EE28 | `Spillelister` | Known | Hardware interface |
| 0x0016EE44 | `Genoptag spil` | Known | Hardware interface |
| 0x0016F2E8 | `Slet spilleliste` | Known | Hardware interface |
| 0x0016F2FC | `Arkiver spilleliste` | Known | Hardware interface |
| 0x0016F39C | `Ny spilleliste %lu` | Known | Hardware interface |
| 0x0016F420 | `Harddisk` | Known | Hardware interface |
| 0x0016F68E | ` den kan bruges som harddisk, og tr` | Known | Hardware interface |
| 0x0016F7F3 | `r du har tilsluttet iPod som disk, skal du dobbeltklikk...` | Known | Hardware interface |
| 0x0016F9BA | ` den kan bruges som disk og anbringe tekstarkiver i map...` | Known | Hardware interface |
| 0x0016FF4B | `je den til spillelisten On-The-Go. Hold knappen nede, n` | Known | Hardware interface |
| 0x0016FF84 | `r en spilleliste, kunstner eller et album er valgt for ...` | Known | Hardware interface |
| 0x0016FFC1 | `je alle sangene til spillelisten On-The-Go. Spilleliste...` | Known | Hardware interface |
| 0x001700C8 | `Nyt spil` | Known | Hardware interface |
| 0x001701F0 | ` Spillelister` | Known | Hardware interface |
| 0x00170594 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x00172124 | `Spiele` | Known | Hardware interface |
| 0x00172198 | `Weiterspielen` | Known | Hardware interface |
| 0x00172A17 | `ffnen Sie das Adressbuch, Microsoft Entourage oder Palm...` | Known | FireWire interface |
| 0x00173524 | `Neues Spiel` | Known | Hardware interface |
| 0x00173A12 | `ber FireWire verbunden` | Known | FireWire interface |
| 0x00176980 | `FireWire conectado` | Known | FireWire interface |
| 0x00178520 | `Diskanttivahv.` | Known | Hardware interface |
| 0x00178530 | `Diskanttiheikenn.` | Known | Hardware interface |
| 0x001785F4 | `Ladataan` | Known | ATA/disk interface |
| 0x00178CC0 | `Muistiinpanoja ladataan.` | Known | ATA/disk interface |
| 0x00178CFC | `Yhteystietoja ladataan.` | Known | ATA/disk interface |
| 0x001797D0 | `FireWire liitetty` | Known | FireWire interface |
| 0x0017C23D | `utiliser comme disque FireWire. Puis glissez les vCards...` | Known | FireWire interface |
| 0x0017D294 | `FireWire Connect` | Known | FireWire interface |
| 0x0017FB24 | `Data & Ora` | Known | ATA/disk interface |
| 0x0017FBC8 | `Imposta Data & Ora` | Known | ATA/disk interface |
| 0x00180078 | `FireWire Connesso` | Known | FireWire interface |
| 0x00183878 | `FireWire ` | Known | FireWire interface |
| 0x00188D1C | `Op de iPod kunt u adres- en agendagegevens opslaan. Als...` | Known | FireWire interface |
| 0x00189E5C | `FireWire aangesloten` | Known | FireWire interface |
| 0x0018B464 | `Spilles n` | Known | Hardware interface |
| 0x0018B518 | `Fortsett spill` | Known | Hardware interface |
| 0x0018B96C | `Mer diskant` | Known | Hardware interface |
| 0x0018B978 | `Mindre diskant` | Known | Hardware interface |
| 0x0018B9BC | `Slett spilleliste` | Known | Hardware interface |
| 0x0018BB10 | `Diskmodus` | Known | Hardware interface |
| 0x0018BCFC | `pner du Adressebok, Microsoft Entourage eller Palm Desk...` | Known | Hardware interface |
| 0x0018BE54 | `ringer i tillegg til musikken din. Microsoft Outlook, M...` | Known | Hardware interface |
| 0x0018C094 | `Hvis du vil vise tekstfiler her, aktiverer du iPod for ...` | Known | Hardware interface |
| 0x0018C654 | ` legge den til i On-The-Go-spillelisten. Du kan legge t...` | Known | Hardware interface |
| 0x0018C77C | `Nytt spill` | Known | Hardware interface |
| 0x0018CC48 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x0018F098 | `rddisk. D` | Known | Hardware interface |
| 0x0018F1F6 | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x0018F216 | `rddisk. Sedan dubbelklickar du bara p` | Known | Hardware interface |
| 0x0018F408 | `rddiskl` | Known | Hardware interface |
| 0x0018FFC0 | `FireWire anslutet` | Known | FireWire interface |
| 0x001A3E8C | `Disk Mode` | Known | Hardware interface |
| 0x001A3F60 | `Your iPod can store contacts and calendar events. If yo...` | Known | Hardware interface |
| 0x001A4128 | `Your iPod can store up to one thousand contacts right a...` | Known | Hardware interface |
| 0x001A4388 | `To view text files here, enable iPod for disk use, then...` | Known | Hardware interface |
| 0x001A4F74 | `FireWire Connected` | Known | FireWire interface |
| 0x001A4F88 | `No battery power remains. Please connect iPod to power.` | Known | Power management |
| 0x001A4FC0 | `Low Battery` | Known | Power management |
| 0x001BD4FC | `TROMResourceDB - unknown header version! (Try regenerat...` | Known | ATA/disk interface |
| 0x001CDA64 | `Bad Data` | Known | ATA/disk interface |
| 0x001CE410 | `Bad meta data, name not a recognized preference. (27)` | Known | ATA/disk interface |
| 0x001CE4C8 | `Bad meta data, name not found. (23)` | Known | ATA/disk interface |
| 0x001CE4EC | `Bad meta data, name termination quote not found. (24)` | Known | ATA/disk interface |
| 0x001CE574 | `Bad meta data, content not found. (25)` | Known | ATA/disk interface |
| 0x001CE59C | `Bad meta data, content termination quote not found. (26...` | Known | ATA/disk interface |
| 0x001CE6C8 | `Bad meta data, encoding not found. (8)` | Known | ATA/disk interface |
| 0x001CE6F0 | `Bad meta data, encoding termination quote not found. (9...` | Known | ATA/disk interface |
| 0x001CE72C | `Bad meta data, unknown encoding. (17)` | Known | ATA/disk interface |
| 0x001CE8E8 | `Bad data. (32)` | Known | ATA/disk interface |
| 0x001CEA7C | `USB MSC` | Known | USB interface |
| 0x0030387F | `ataTaza[aea;ajaaaVa)b'b+b+dMd[d]dtdvdrdsd}dudfd` | Known | ATA/disk interface |
| 0x00308469 | `aGa>a(a'aJa?a<a,a4a=aBaDasawaXaYaZakataoaeaqa_a]aSaua` | Known | ATA/disk interface |
| 0x00313D98 | `MaximumDataRate` | Known | ATA/disk interface |
| 0x00313E04 | `FireWire` | Known | FireWire interface |
| 0x00313E74 | `FireWireVersion` | Known | FireWire interface |
| 0x00314044 | `MEMDISK` | Known | Hardware interface |
| 0x003140DA | `ex_data` | Known | ATA/disk interface |
| 0x003142AC | `c:\buildtools\MWSF2\Q22Firmware.proj\projectfiles\sandb...` | Known | ATA/disk interface |
| 0x003153B7 | `RSA Data Security, Inc.` | Known | ATA/disk interface |
| 0x003153D4 | `RSA Data Security, Inc. PKCS` | Known | ATA/disk interface |
| 0x00315535 | `pkcs7-data` | Known | ATA/disk interface |
| 0x00315540 | `pkcs7-signedData` | Known | ATA/disk interface |
| 0x00315551 | `pkcs7-envelopedData` | Known | ATA/disk interface |
| 0x00315565 | `pkcs7-signedAndEnvelopedData` | Known | ATA/disk interface |
| 0x00315582 | `pkcs7-digestData` | Known | ATA/disk interface |
| 0x00315593 | `pkcs7-encryptedData` | Known | ATA/disk interface |
| 0x003157EC | `nsDataType` | Known | ATA/disk interface |
| 0x003157F7 | `Netscape Data Type` | Known | ATA/disk interface |
| 0x00316591 | `id-smime-ct-authData` | Known | ATA/disk interface |
| 0x003165FE | `id-smime-ct-DVCSRequestData` | Known | ATA/disk interface |
| 0x0031661A | `id-smime-ct-DVCSResponseData` | Known | ATA/disk interface |
| 0x00317032 | `id-cmc-dataReturn` | Known | ATA/disk interface |
| 0x00317185 | `id-on-personalData` | Known | ATA/disk interface |
| 0x00317286 | `id-cct-PKIData` | Known | ATA/disk interface |
| 0x00317A25 | `qualityLabelledData` | Known | ATA/disk interface |
| 0x00317E00 | `setct-PANData` | Known | ATA/disk interface |
| 0x00317E2B | `setct-OIData` | Known | ATA/disk interface |
| 0x00317E41 | `setct-PIData` | Known | ATA/disk interface |
| 0x00317E4E | `setct-PIDataUnsigned` | Known | ATA/disk interface |
| 0x00317EC9 | `setct-PInitResData` | Known | ATA/disk interface |
| 0x00317EE9 | `setct-PResData` | Known | ATA/disk interface |
| 0x00317F3F | `setct-CapTokenData` | Known | ATA/disk interface |
| 0x00317F8D | `setct-AuthRevResData` | Known | ATA/disk interface |
| 0x00317FD7 | `setct-CapResData` | Known | ATA/disk interface |
| 0x0031800F | `setct-CapRevResData` | Known | ATA/disk interface |
| 0x00318046 | `setct-CredResData` | Known | ATA/disk interface |
| 0x00318081 | `setct-CredRevResData` | Known | ATA/disk interface |
| 0x00318096 | `setct-PCertReqData` | Known | ATA/disk interface |
| 0x003180BB | `setct-BatchAdminReqData` | Known | ATA/disk interface |
| 0x003180D3 | `setct-BatchAdminResData` | Known | ATA/disk interface |
| 0x0031812B | `setct-CertReqData` | Known | ATA/disk interface |
| 0x0031814E | `setct-CertResData` | Known | ATA/disk interface |
| 0x003184DA | `setCext-merchData` | Known | ATA/disk interface |
| 0x00318565 | `setCext-Track2Data` | Known | ATA/disk interface |
| 0x00318753 | `set-brand-IATA-ATA` | Known | ATA/disk interface |
| 0x0031D74F | `<key>FireWireGUID</key>` | Known | FireWire interface |

---

## 7. Error Messages

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00017BC0 | `*** assertion failed: %s, file %s, line %d` | Known | Error/assertion message |
| 0x0001E1EC | `Invalid Operation` | Known | Error/assertion message |
| 0x001CE360 | `Too many errors, further errors discarded. (31)` | Known | Error/assertion message |
| 0x001CE3B0 | `%s Error in file %s.` | Known | Error/assertion message |
| 0x001CE774 | `Cannot link to a .link file. (29)` | Known | Error/assertion message |
| 0x001CEE66 | `tInvalid partition table. Setup cannot continue.` | Known | Error/assertion message |
| 0x001CEE97 | `Error loading operating system. Setup cannot continue.` | Known | Error/assertion message |
| 0x00314254 | `%s(%d): OpenSSL internal error, assertion failed: %s` | Known | Error/assertion message |
| 0x00314EE0 | `error:%08lX:%s:%s:%s` | Known | Error/assertion message |
| 0x00315184 | `internal error: list index %ld out of range` | Known | Error/assertion message |

---

## 8. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016CF34 | ` !"#$%&'()*+,-./0123456789:;<=>?@abcdefghijklmnopqrstuv...` | Known | Filesystem path |
| 0x0016D034 | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x0016D234 | `url;type=work:apple.com/support/ipod` | Known | Filesystem path |
| 0x0016FA60 | `Point: %d (%d/%d)` | Known | Filesystem path |
| 0x0016FBBB | `rende VoiceAge Corporation i USA og/eller andre lande o...` | Known | Filesystem path |
| 0x0017228C | `USA/Hawaii (NZ)` | Known | Filesystem path |
| 0x0017229C | `USA/Hawaii (SZ)` | Known | Filesystem path |
| 0x001722AC | `USA/Alaska (NZ)` | Known | Filesystem path |
| 0x001722BC | `USA/Alaska (SZ)` | Known | Filesystem path |
| 0x001722CC | `USA/Pazifik (NZ)` | Known | Filesystem path |
| 0x001722E0 | `USA/Pazifik (SZ)` | Known | Filesystem path |
| 0x001722F4 | `USA/Rockies (NZ)` | Known | Filesystem path |
| 0x00172308 | `USA/Rockies (SZ)` | Known | Filesystem path |
| 0x0017231C | `USA/Zentral (NZ)` | Known | Filesystem path |
| 0x00172330 | `USA/Zentral (SZ)` | Known | Filesystem path |
| 0x00172344 | `USA/Ost (NZ)` | Known | Filesystem path |
| 0x00172354 | `USA/Ost (SZ)` | Known | Filesystem path |
| 0x00172658 | `Vorn./Nachn.` | Known | Filesystem path |
| 0x00172668 | `Nachn./Vorn.` | Known | Filesystem path |
| 0x00172ED8 | `Punkte: %d (%d/%d)` | Known | Filesystem path |
| 0x00172FFB | ` ist entweder eine eingetragene Marke oder eine Marke d...` | Known | Filesystem path |
| 0x00175E24 | `Result.: %d (%d/%d)` | Known | Filesystem path |
| 0x00175F47 | ` es una marca registrada o una marca comercial de Voice...` | Known | Filesystem path |
| 0x001768CC | `Fecha/hora` | Known | Filesystem path |
| 0x001785EC | `%d / %d` | Known | Filesystem path |
| 0x00178C8C | `Tulos: %d (%d/%d)` | Known | Filesystem path |
| 0x00178DD5 | `ity tavaramerkki Yhdysvalloissa ja/tai muissa maissa, j...` | Known | Filesystem path |
| 0x0017C703 | `sult. : %d (%d/%d)` | Known | Filesystem path |
| 0x0017C87A | `tats-Unis et/ou dans d` | Known | Filesystem path |
| 0x0017CD87 | `gler date/heure` | Known | Filesystem path |
| 0x0017F5FC | `Punti: %d (%d/%d)` | Known | Filesystem path |
| 0x00181664 | `%b/%-d %-I:%M %2p` | Known | Filesystem path |
| 0x00181678 | `%-m/%-d` | Known | Filesystem path |
| 0x0018169C | `%y/%-m/%d` | Known | Filesystem path |
| 0x001816A8 | `%Y/%b/%-d` | Known | Filesystem path |
| 0x00182A6C | ` %d (%d/%d)` | Known | Filesystem path |
| 0x00184ED8 | `%Y/%B/%d %A  %I:%M:%S %p` | Known | Filesystem path |
| 0x00184EF4 | `%Y/%B/%d` | Known | Filesystem path |
| 0x00184F0C | `%-m/%-d %-I:%M %p` | Known | Filesystem path |
| 0x00184F40 | `%Y/%-m/%d` | Known | Filesystem path |
| 0x00185E8A | `: %d (%d/%d)` | Known | Filesystem path |
| 0x00189330 | `Score: %d (%d/%d)` | Known | Filesystem path |
| 0x0018943B | ` is een gedeponeerd handelsmerk of een handelsmerk van ...` | Known | Filesystem path |
| 0x001899A8 | `Stel datum/tijd in` | Known | Filesystem path |
| 0x00189D9C | `Datum/tijd` | Known | Filesystem path |
| 0x0018C17C | `Poeng: %d (%d/%d)` | Known | Filesystem path |
| 0x0018C2A3 | ` er enten et registrert varemerke eller et varemerke fo...` | Known | Filesystem path |
| 0x0018E73C | `%-d/%-m` | Known | Filesystem path |
| 0x0018F4B0 | `ng: %d (%d/%d)` | Known | Filesystem path |
| 0x0018F604 | `r VoiceAge Corporation i USA och/eller andra l` | Known | Filesystem path |
| 0x0018FA7C | `Datum/tid` | Known | Filesystem path |
| 0x0018FB18 | `ll in datum/tid` | Known | Filesystem path |
| 0x001924B9 | `%d (%d/%d)` | Known | Filesystem path |
| 0x001A3344 | `%-m/%d/%y` | Known | Filesystem path |
| 0x001A4593 | ` is either registered trademark or trademark of VoiceAg...` | Known | Filesystem path |
| 0x001CD7E9 | `file://` | Known | Filesystem path |
| 0x001CD7F1 | `image://` | Known | Filesystem path |
| 0x001CD9D0 | `</TITLE>` | Known | Filesystem path |
| 0x001CD9E0 | `</BODY>` | Known | Filesystem path |
| 0x001CDA0E | `</ROT13>` | Known | Filesystem path |
| 0x001CE7B8 | `Bad link, no matching </a> for anchor tag. (20)` | Known | Filesystem path |
| 0x001E66D3 | `W/}lE>q` | Known | Filesystem path |
| 0x00215315 | `H."0*Bx/` | Known | Filesystem path |
| 0x0021CC72 | `U/~RERT` | Known | Filesystem path |
| 0x002219DA | `TUOPT/\|` | Known | Filesystem path |
| 0x00228C6F | `HuGZp/$j` | Known | Filesystem path |
| 0x0022F17F | `(bJ)b"(b2""z/` | Known | Filesystem path |
| 0x002318EF | `JUAPDD(/` | Known | Filesystem path |
| 0x00237BFE | `/B\|$BD'` | Known | Filesystem path |
| 0x00238B7B | `$Bd$BT/` | Known | Filesystem path |
| 0x0023EE53 | `/" +J\|!` | Known | Filesystem path |
| 0x00245CD2 | `Fb""")/` | Known | Filesystem path |
| 0x00246E89 | `/RyO(UIH` | Known | Filesystem path |
| 0x00247F39 | `~$Bh'"~$Bz/` | Known | Filesystem path |
| 0x0024B7A3 | `$B +BZ/` | Known | Filesystem path |
| 0x002531C1 | `0c(HBP/` | Known | Filesystem path |
| 0x00257757 | `$B~("\|/` | Known | Filesystem path |
| 0x0026E4C1 | `T/DDDDD` | Known | Filesystem path |
| 0x0026E72B | `"~UeB /` | Known | Filesystem path |
| 0x002711DD | `$B((B /` | Known | Filesystem path |
| 0x002792C0 | ` "\|$B~/` | Known | Filesystem path |
| 0x0027C43C | `@$B\|$"(/` | Known | Filesystem path |
| 0x0027D53C | `)"8/B""` | Known | Filesystem path |
| 0x0027DC84 | `r4c6 bN/` | Known | Filesystem path |
| 0x00283085 | `RDT%B(/` | Known | Filesystem path |
| 0x00284211 | `RBHUE\|/` | Known | Filesystem path |
| 0x0028B891 | `]B""B</` | Known | Filesystem path |
| 0x0028F10E | `,B\|RED/` | Known | Filesystem path |
| 0x00294B91 | `$BT). /` | Known | Filesystem path |
| 0x00295D91 | `#"TUB(/` | Known | Filesystem path |
| 0x002A7605 | `/" %BD"` | Known | Filesystem path |
| 0x002ADF28 | `ODD""(/` | Known | Filesystem path |
| 0x002AF163 | `B"$R%"B$" /` | Known | Filesystem path |
| 0x002AFAD0 | `bG\|jG\|/` | Known | Filesystem path |
| 0x002B16BA | `$E$$BR/` | Known | Filesystem path |
| 0x002B175B | `dRB~RA$/` | Known | Filesystem path |
| 0x002B20AC | `TT&T%B(/` | Known | Filesystem path |
| 0x002C0DA0 | `)'>$B8/` | Known | Filesystem path |
| 0x002C33AE | `$B\|%EV/` | Known | Filesystem path |
| 0x002CA902 | `BDU!BJ ""/` | Known | Filesystem path |
| 0x002CB866 | `Z-bD("(%B>/` | Known | Filesystem path |
| 0x002D3FB5 | `@(/ Q\|f` | Known | Filesystem path |
| 0x002FB1C9 | `% %!%"%#%$%%%&%'%(%)%*%+%,%-%.%/%0%1%2%3%4%5%6%7%8%9%:%...` | Known | Filesystem path |
| 0x002FB9A9 | `qWlIl/Ymg*` | Known | Filesystem path |
| 0x002FD217 | `X)W,W*W3W9W.W/W\W;WBWiW` | Known | Filesystem path |
| 0x002FD8E5 | `n/o6oKoto*o` | Known | Filesystem path |
| 0x002FDE7D | `q/q1qsq\qhqEqrqJqxqzq` | Known | Filesystem path |
| 0x002FF105 | `b6bKbNb/e` | Known | Filesystem path |
| 0x002FF2BB | `V3W0W(W-W,W/W)W` | Known | Filesystem path |
| 0x002FF38F | `NMOOOGOWO^O4O[OUO0OPOQO=O:O8OCOTO<OFOcO\O`O/ONO6OYO]OHO...` | Known | Filesystem path |
| 0x002FF9ED | `e%f-f f'f/f` | Known | Filesystem path |
| 0x002FFCEF | `bNc>c/cUcBcFcOcIc:cPc=c*c+c(cMcLcHeIe` | Known | Filesystem path |
| 0x002FFFA1 | `fFUjUfUDU^UaUCUJU1UVUOUUU/UdU8U.U\U,UcU3UAUWU` | Known | Filesystem path |
| 0x00300207 | `\|F}C}q}.}9}<}@}0}3}D}/}B}2}1}=` | Known | Filesystem path |
| 0x0030038F | `W/X*X4X$X0X1X!X` | Known | Filesystem path |
| 0x003004A5 | `k.l/l,l/n8nTn!n2ngnJn n%n#n` | Known | Filesystem path |
| 0x00300507 | `r6s%s4s)s:t*t3t"t%t5t6t4t/t` | Known | Filesystem path |
| 0x00300D85 | `S.V;V9V2V?V4V)VSVNVWVtV6V/V0V` | Known | Filesystem path |
| 0x003013CB | `\|.~>~F~7~2~C~+~=~1~E~A~4~9~H~5~?~/~D` | Known | Filesystem path |
| 0x003023BB | `P P'P5P/P1P` | Known | Filesystem path |
| 0x0030250D | `h5h+h-h/hNhDh4h` | Known | Filesystem path |
| 0x00302521 | `h&h(h.hMh:h%h h,k/k-k1k4kmk` | Known | Filesystem path |
| 0x0030262D | `w"w'w#x,x"x5x/x(x.x+x!x)x3x*x1xTy[yOy\ySyRyQy` | Known | Filesystem path |
| 0x00302AB7 | `v9w/w-w1w2w4w3w=w%w;w5wHxRxIxMxJxLx&xExPxdygyiyjycykyay` | Known | Filesystem path |
| 0x00302B13 | `{1{+{-{/{2{8{` | Known | Filesystem path |
| 0x00302FA7 | `t/uoulu` | Known | Filesystem path |
| 0x0030330D | ``IaJa+aEa6a2a.aFa/aOa)a@a bh` | Known | Filesystem path |
| 0x00303353 | `d&d0d(dAd5d/d` | Known | Filesystem path |
| 0x003033BB | `hxi4iii@ioiDiviXiAitiLi;iKi7i\iOiQi2iRi/i{i<iFkEkCkBkHk...` | Known | Filesystem path |
| 0x00303A25 | `y+zJz0z/z(z&z` | Known | Filesystem path |
| 0x00304309 | `u/v-v1v=v3v<v5v2v0v` | Known | Filesystem path |
| 0x003045FF | `X-[%[2[#[,['[&[/[.[{[` | Known | Filesystem path |
| 0x00304F2D | `j<p5p/p7p4p1pBp8p?p:p9p@p;p3pAp` | Known | Filesystem path |
| 0x0030512F | `w-y1y/yT\|S\|` | Known | Filesystem path |
| 0x003064FB | `%#%3%+%;%K% %/%(%7%?%` | Known | Filesystem path |
| 0x0030667A | `02*2+2,2-2.2/2@272B2C292:212>24222;2623252<2=2?282` | Known | Filesystem path |
| 0x00306BA5 | `\7_J_/`P`m`` | Known | Filesystem path |
| 0x00306D81 | `OHSIT>T/Z` | Known | Filesystem path |
| 0x00306DF7 | `i_l*mim/n` | Known | Filesystem path |
| 0x003072E9 | `N,p]u/f` | Known | Filesystem path |
| 0x00307FF5 | `S#S/S1S3S8S@SFSES` | Known | Filesystem path |
| 0x003080F5 | `q4V6V2V8VkVdV/VlVjV` | Known | Filesystem path |
| 0x0030839D | `_)_-_8_A_H_L_N_/_Q_V_W_Y_a_m_s_w_` | Known | Filesystem path |
| 0x00308A63 | `s4s/s)s%s>sNsOs` | Known | Filesystem path |
| 0x0030C867 | `h>kLp/t` | Known | Filesystem path |
| 0x0030CEF1 | `o;v/}7~` | Known | Filesystem path |
| 0x0030DCB5 | `e1f/h\q6z` | Known | Filesystem path |
| 0x0030E301 | `UuX/c"dIfKfmh` | Known | Filesystem path |
| 0x00313BD8 | `%s<key>%s</key>` | Known | Filesystem path |
| 0x00313BE8 | `%s<%s/>` | Known | Filesystem path |
| 0x00313C10 | `%s</dict>` | Known | Filesystem path |
| 0x00313C2C | `%s<string>%s</string>` | Known | Filesystem path |
| 0x00313C54 | `%s<integer>%s</integer>` | Known | Filesystem path |
| 0x00313C80 | `%s<integer>%d</integer>` | Known | Filesystem path |
| 0x00313FE4 | ` ,;=+[]*?<>\|":/\` | Known | Filesystem path |
| 0x00313FF5 | `\/:*?"<>\|` | Known | Filesystem path |
| 0x00315018 | `You need to read the OpenSSL FAQ, http://www.openssl.or...` | Known | Filesystem path |
| 0x003162B1 | `S/MIME Capabilities` | Known | Filesystem path |
| 0x0031D6B0 | `<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1....` | Known | Filesystem path |
| 0x0031D73E | `</dict>` | Known | Filesystem path |
| 0x0031D746 | `</plist>` | Known | Filesystem path |
| 0x0031D767 | `<string>%08lX%08lX</string>` | Known | Filesystem path |
| 0x0031EEFE | `/ 0 1 2 3 4 5 6 7 8 9 : ; < = > ? @ A B C D E F G H I J...` | Known | Filesystem path |
| 0x0031F0DF | `! !!!"!#!$!%!&!'!(!)!*!+!,!-!.!/!0!1!2!3!4!5!6!7!8!9!:!...` | Known | Filesystem path |
| 0x00336214 | `/1f;{1Q` | Known | Filesystem path |
| 0x0034D462 | `Tfp9/deM5` | Known | Filesystem path |
| 0x0034DB64 | `UGa</J]m` | Known | Filesystem path |
| 0x00362BDB | `\uc&/5:` | Known | Filesystem path |
| 0x00378E5E | `)A/AK54A` | Known | Filesystem path |
| 0x0037FFBB | `]Z&/f6y` | Known | Filesystem path |
| 0x00383249 | `P3/i8E_` | Known | Filesystem path |
| 0x00384C18 | `$q9/a_A` | Known | Filesystem path |
| 0x003934FE | ``*/gI(I.` | Known | Filesystem path |
| 0x0039705B | `/g%9j-]` | Known | Filesystem path |
| 0x0039A3CE | `HV7v/Wh/` | Known | Filesystem path |
| 0x0039BE07 | `c?/LZtlQV` | Known | Filesystem path |
| 0x003ACBD8 | `~/@Ndi8` | Known | Filesystem path |
| 0x003AF99D | `/_72>Hx;` | Known | Filesystem path |
| 0x003B5103 | `%&1AwCz/c25` | Known | Filesystem path |
| 0x003B9C23 | `l[H/Af^` | Known | Filesystem path |
| 0x003BA83B | `xXNSLg/` | Known | Filesystem path |
| 0x003BF02D | `y 5{\{/-` | Known | Filesystem path |
| 0x003C048A | `P/C*h~X` | Known | Filesystem path |
| 0x003C1B5F | `tRd#p/_` | Known | Filesystem path |
| 0x003CBE4F | `4%5K/$?h` | Known | Filesystem path |
| 0x003CBF46 | `o/y^p#h` | Known | Filesystem path |
| 0x003D7967 | `+_BRx/wH~` | Known | Filesystem path |
| 0x003E62D3 | `/]/5!83` | Known | Filesystem path |
| 0x003EFD08 | `@n6_/he` | Known | Filesystem path |
| 0x003F731A | `1/6v:Kt` | Known | Filesystem path |
| 0x003FC4FF | `3)486/F` | Known | Filesystem path |
| 0x003FD677 | `/&A<Gixq` | Known | Filesystem path |
| 0x003FF519 | `OB/nbUi` | Known | Filesystem path |
| 0x00403F0F | `9/6RI5T lC6` | Known | Filesystem path |
| 0x004088B0 | `/tB.uUk0` | Known | Filesystem path |
| 0x0040EE4F | `JC\|U/kb^` | Known | Filesystem path |
| 0x00418360 | `S$+Sl/3A9V` | Known | Filesystem path |
| 0x00419556 | `/ 0`+a!` | Known | Filesystem path |
| 0x0041B22D | `/QK'/8w` | Known | Filesystem path |
| 0x0041D909 | `]J%w)/]F` | Known | Filesystem path |
| 0x0041EDCE | `0)/HkmQ)` | Known | Filesystem path |
| 0x00426394 | `D8@/6;%` | Known | Filesystem path |
| 0x00429004 | `0lm8/FEn` | Known | Filesystem path |
| 0x004327BE | `/XxkVTT#)` | Known | Filesystem path |
| 0x00436DF1 | `E?/EO2,` | Known | Filesystem path |
| 0x0043FDEA | `*~k.[)'/` | Known | Filesystem path |
| 0x00443A4C | `Y}/L!TA*J` | Known | Filesystem path |
| 0x0044937C | `>3:rMS/` | Known | Filesystem path |

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

