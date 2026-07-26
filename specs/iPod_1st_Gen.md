# iPod 1st Generation - RetailOS 1.1.5 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.1.5 |
| **IPSW** | iPod_1.1.5.ipsw |
| **Device** | iPod 1st Generation (2001, 5/10GB, Scroll Wheel) |
| **Binary Size** | 5,066,752 bytes (4.83 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 5,066,752 bytes |
| **Total Strings (>=6)** | 26,630 |
| **Function Prologues** | 6,885 |
| **SoC** | PortalPlayer PP5002 |
| **Architecture** | ARM7TDMI (ARMv4T), dual-core |
| **Encrypted** | No |
| **Decryption Method** | Extract from IPSW (unencrypted) |
| **SHA-256** | `297498df5c42a3a85f4b3d153bdf0d99546ae6655090d35dd5bb30ee0002cd3e` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000AFB0 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x0004A502 | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x0004A69C | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x0004A81B | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x0004A85C | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x0004A876 | `RTXCbug> ` | Hidden | Developer Tool |
| 0x0004AF4A | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x0004B01E | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x00147C26 | `Returning from RTXCbug` | Hidden | Developer Tool |
| 0x0031AAE0 | `SP_ERR_ASSERTION` | Hidden | Developer Tool |
| 0x0031E1DC | ` <- Error:Internal Error` | Hidden | Undocumented UI |
| 0x00326A94 | `BTM Debug Zones %s (0x%08X)` | Hidden | Debug/Diagnostic |
| 0x0032A634 | `Retail mode` | Hidden | Demo/Retail Mode |
| 0x0032A644 | `Debug mode` | Hidden | Debug/Diagnostic |

---

## 2. Discovered Features

### Diagnostic

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0014A33A | `DIAGNOSTIC FAILURE ON COMPONENT NN (80H-FFH)` | Diagnostic | |
| 0x0032A60C | `Diag mode` | Diagnostic | |

### Factory/Calibration

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0014AA3D | `POWER CALIBRATION AREA ALMOST FULL` | Factory/Calibration | |
| 0x0014AA60 | `POWER CALIBRATION AREA IS FULL` | Factory/Calibration | |
| 0x0014AA7F | `POWER CALIBRATION AREA ERROR` | Factory/Calibration | |

### EQ Preset

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001509EA | `Secure Electronic Transactions` | EQ Preset | |
| 0x00291F08 | `Acoustic` | EQ Preset | |
| 0x00291F14 | `Bass Booster` | EQ Preset | |
| 0x00291F34 | `Classical` | EQ Preset | |
| 0x00291F50 | `Electronic` | EQ Preset | |
| 0x00291F64 | `Hip Hop` | EQ Preset | |
| 0x00291F7C | `Loudness` | EQ Preset | |
| 0x00291F88 | `Lounge` | EQ Preset | |
| 0x00291FAC | `Small Speakers` | EQ Preset | |
| 0x00291FBC | `Spoken Word` | EQ Preset | |
| 0x00291FC8 | `Treble Booster` | EQ Preset | |
| 0x00291FE8 | `Vocal Booster` | EQ Preset | |
| 0x00295EB8 | `USA/Rockies (DST)` | EQ Preset | |
| 0x00295ECC | `USA/Rockies` | EQ Preset | |
| 0x002977B0 | `Latina` | EQ Preset | |
| 0x0029D678 | `Latino` | EQ Preset | |

### Localization

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00145EE0 | `x-mac-japanese` | Localization | |
| 0x002C6720 | `English` | Localization | |
| 0x002C6758 | `Italiano` | Localization | |

### Filesystem Path

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000351A4 | `iPod_Control\Device` | Filesystem Path | |
| 0x000351F4 | `iPod_Control` | Filesystem Path | |
| 0x000352FC | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x00038D64 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x00039098 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x0006D00C | `iPod_Control\Testing\TestLog.txt` | Filesystem Path | |
| 0x0006D030 | `\iPod_Control\Testing\TestLog.txt` | Filesystem Path | |
| 0x0006D364 | `iPod_Control\Testing` | Filesystem Path | |
| 0x0006D37C | `\iPod_Control\Testing` | Filesystem Path | |
| 0x0006D3A0 | `iPod_Control\Testing\Tests.Lock` | Filesystem Path | |
| 0x0006D3C0 | `\iPod_Control\Testing\Tests.Lock` | Filesystem Path | |
| 0x0007CAB8 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x0007CF18 | `iPod_Control\Device\Limit` | Filesystem Path | |
| 0x000DB014 | ` </MusicItem>` | Filesystem Path | |

### Assertion

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F60F8 | `%s(%d): OpenSSL internal error, assertion failed: %s` | Assertion | |
| 0x0011D9C4 | `*** assertion failed: %s, file %s, line %d` | Assertion | |
| 0x0031AAE0 | `SP_ERR_ASSERTION` | Assertion | |

---

## 3. Known User-Facing Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000750C | `KeyRepeatTimer` | Known | UI element |
| 0x00048D78 | `Clock rate is %d Hz, Tick interval is %d ms, ` | Known | UI element |
| 0x0004A680 | `U - Return To Main Menu` | Known | Menu item |
| 0x00148670 | `vcalendar` | Known | UI element |
| 0x00148724 | `dalarm` | Known | UI element |
| 0x00148890 | `valarm` | Known | UI element |
| 0x00148990 | `\Contacts` | Known | UI element |
| 0x0014899A | `Contacts` | Known | UI element |
| 0x001489A3 | `\Calendars` | Known | UI element |
| 0x001489AE | `Calendars` | Known | UI element |
| 0x0014A7A8 | `ILLEGAL MODE FOR THIS TRACK` | Known | UI element |
| 0x0014B0AE | `Contrast` | Known | UI element |
| 0x0014B0C2 | `BacklightTimer` | Known | UI element |
| 0x0014B0D1 | `Shuffle` | Known | UI element |
| 0x0014B0D9 | `RepeatMode` | Known | UI element |
| 0x0014B0F5 | `BacklightState` | Known | UI element |
| 0x0014B17A | `ContactsDisplay` | Known | UI element |
| 0x0014B18A | `ContactsSort` | Known | UI element |
| 0x0014BA96 | `<tracknotes>` | Known | UI element |
| 0x00152046 | `Illegal instruction` | Known | UI element |
| 0x00152074 | `Illegal address` | Known | UI element |
| 0x0029273B | `lpemenuen.` | Known | Menu item |
| 0x002930B0 | `Nulstil menu` | Known | Menu item |
| 0x002931F8 | `Alarmer` | Known | UI element |
| 0x00293278 | `Hovedmenu` | Known | Menu item |
| 0x00293330 | `Menuer` | Known | Menu item |
| 0x00294CAC | `Extras` | Known | UI element |
| 0x0029552C | `ffnen Sie ihn durch Doppelklicken in das iPod Symbol un...` | Known | UI element |
| 0x00295737 | `ffnen Sie ihn durch Doppelklicken in das iPod Symbol un...` | Known | UI element |
| 0x00295882 | `hlen", um den Alarm zu beenden.` | Known | UI element |
| 0x00297684 | `Calendario` | Known | UI element |
| 0x00297690 | `Calendarios` | Known | UI element |
| 0x0029823C | `Alarma` | Known | UI element |
| 0x00298B90 | `Alarmas` | Known | UI element |
| 0x00298B98 | `Contraste` | Known | UI element |
| 0x0029A9C6 | ` vCardit Contacts-kansioon. T` | Known | UI element |
| 0x0029DD44 | ` mille contacts en plus de votre musique. Les applicati...` | Known | UI element |
| 0x0029DD95 | `adresses, Microsoft Entourage et Palm Desktop peuvent t...` | Known | UI element |
| 0x0029DE74 | `e sur votre bureau, puis glissez ces vCards dans le dos...` | Known | UI element |
| 0x0029DF1C | `es, consultez l'Aide iPod dans le menu Aide iTunes.` | Known | Menu item |
| 0x0029DF74 | ` mille contacts en plus de votre musique. Les applicati...` | Known | UI element |
| 0x0029E09C | `e sur votre bureau, puis glissez ces vCard dans le doss...` | Known | UI element |
| 0x0029E1D8 | `Alarme` | Known | UI element |
| 0x0029E1E0 | `Chargement contacts.` | Known | UI element |
| 0x0029E92F | `init. menu p.` | Known | Menu item |
| 0x0029EAA8 | `Alarmes` | Known | UI element |
| 0x0029EB44 | `Menu principal` | Known | Menu item |
| 0x0029ED7C | `Menu princ.` | Known | Menu item |
| 0x002A004C | `Calendari` | Known | UI element |
| 0x002A0981 | ` dettagliate, fai clic su "Aiuto iPod" disponibile dal ...` | Known | Menu item |
| 0x002A12C0 | `Ripr. Menu Princ.` | Known | Menu item |
| 0x002A1418 | `Contrasto` | Known | UI element |
| 0x002A6424 | ` "Contacts" ` | Known | UI element |
| 0x002A8BCC | `Shuffle nummers` | Known | UI element |
| 0x002A9F1C | `Herstel menu` | Known | Menu item |
| 0x002AA10C | `Hoofdmenu` | Known | Menu item |
| 0x002AA1D4 | `Menu's` | Known | Menu item |
| 0x002AA20C | `Backlight` | Known | UI element |
| 0x002AE764 | `es virtuais para a pasta "Contacts" do iPod. Eles ser` | Known | UI element |
| 0x002AE80D | ` no menu Ajuda do iTunes.` | Known | Menu item |
| 0x002AE97F | `es virtuais (vCards) para a pasta "Contacts" do iPod. E...` | Known | UI element |
| 0x002AF2E4 | `Redef. Menu Princ.` | Known | Menu item |
| 0x002AF4D0 | `Menu Principal` | Known | Menu item |
| 0x002AF6E8 | `Menu Prin.` | Known | Menu item |
| 0x002B11B3 | ` skrivbordet och drar in vCard-filerna i mappen "Contac...` | Known | UI element |
| 0x002B3AF9 | ` Contacts ` | Known | UI element |
| 0x002C663C | `Now Playing` | Known | UI element |
| 0x002C6698 | `Calendar` | Known | UI element |
| 0x002C66E4 | `Shuffle Songs` | Known | UI element |
| 0x002C6E16 | `Your iPod can store up to one thousand contacts right a...` | Known | UI element |
| 0x002C7258 | `Contacts loading.` | Known | UI element |
| 0x002C795C | `Reset Main Menu` | Known | Menu item |
| 0x002C7A2C | `Reset All Settings` | Known | User setting |
| 0x002C7AA8 | `Sleep Timer` | Known | UI element |
| 0x002C7AB4 | `Alarms` | Known | UI element |
| 0x002C7B08 | `Backlight Timer` | Known | UI element |
| 0x002C7B24 | `Repeat` | Known | UI element |
| 0x002C7B34 | `Main Menu` | Known | Menu item |
| 0x002C7BA8 | `Settings` | Known | User setting |
| 0x002C7C38 | `Reset All` | Known | UI element |
| 0x002ECA33 | `S_MENU` | Known | Menu item |
| 0x002ECABB | `S_ALARM` | Known | UI element |
| 0x002F01C8 | `T_MENUBUTTON` | Known | Menu item |
| 0x002F0250 | `T_ALARM` | Known | UI element |
| 0x003261BB | `Set volume and other settings on audio chip.` | Known | User setting |
| 0x003285D3 | ` Display help menu` | Known | Menu item |
| 0x003296CD | `<FN_NAME='%s' spFn_t=0x%08X MENU_OPTION=%d>` | Known | Menu item |
| 0x003F3710 | ` CONTRAST` | Known | UI element |
| 0x003F6085 | ` The contrast value is ,` | Known | UI element |

---

## 4. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00009CE8 | `HostOSTask` | Known | RTOS task thread |
| 0x0000AFB0 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x0002C88C | `RunTestsTask` | Known | RTOS task thread |
| 0x0002DB44 | `LoadDataTasks` | Known | RTOS task thread |
| 0x00048886 | `Task (# or name)> ` | Known | RTOS task thread |
| 0x00048A8B | `** Stack Snapshot **` | Known | RTOS task thread |
| 0x00048B0E | `RTXC Kernel           %08p %5d %5d %5d` | Known | RTOS task thread |
| 0x00048D5B | `** Clock Snapshot **` | Known | RTOS task thread |
| 0x00049197 | `** Mailbox Snapshot **` | Known | RTOS task thread |
| 0x000497CF | `** Semaphore Snapshot **` | Known | RTOS task thread |
| 0x00049A3F | `** Queue Snapshot **` | Known | RTOS task thread |
| 0x00049E77 | `** Task Snapshot **` | Known | RTOS task thread |
| 0x0004A502 | `RTXCbug - RTXC Objects> ` | Known | RTOS task thread |
| 0x0004A55C | `T - Tasks` | Known | RTOS task thread |
| 0x0004A61C | `$ - Enter Task Manager Mode` | Known | RTOS task thread |
| 0x0004A63C | `# - Task Registers` | Known | RTOS task thread |
| 0x0004A69C | `X - Exit RTXCbug` | Known | RTOS task thread |
| 0x0004A81B | `** RTXCbug - ` | Known | RTOS task thread |
| 0x0004A82C | `  K - RTXC` | Known | RTOS task thread |
| 0x0004A85C | `  X - Exit RTXCbug` | Known | RTOS task thread |
| 0x0004A876 | `RTXCbug> ` | Known | RTOS task thread |
| 0x0004AAC8 | `Undefined Task` | Known | RTOS task thread |
| 0x0004AB0C | `Task Register set is undefined` | Known | RTOS task thread |
| 0x0004AB32 | `** Task Register Snapshot **` | Known | RTOS task thread |
| 0x0004AF4A | `$RTXCbug> ` | Known | RTOS task thread |
| 0x0004AFFC | `X - Exit Task Manager Mode` | Known | RTOS task thread |
| 0x0004B01E | `Re-entering RTXCbug mode` | Known | RTOS task thread |
| 0x0007C554 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x00145284 | `RTXC v3.2fpp for ARM and Thumb - ARM ADS 1.0 Jul-08-00 ...` | Known | RTOS task thread |
| 0x00147C26 | `Returning from RTXCbug` | Known | RTOS task thread |
| 0x002EC47D | `S_RTXCBUG` | Known | RTOS task thread |
| 0x002F0085 | `T_RTXCBUG` | Known | RTOS task thread |

---

## 5. Audio/Codec

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000C55C4 | `$$$tag_temp.mp3` | Known | MP3 codec |
| 0x000D7598 | `default_source_file.wav` | Known | PCM audio format |
| 0x0014B9F8 | `<codec>` | Known | Audio system |
| 0x0014BA00 | `</codec>` | Known | Audio system |
| 0x0014EBE3 | `msCodeCom` | Known | Audio system |
| 0x0014F963 | `aaControls` | Known | AAC codec |
| 0x002929A8 | `Audible-softwaren i dette produkt bruges i henhold til ...` | Known | Audible audiobook format |
| 0x00292A08 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x00292AFA | `.net codec i dette produkt bruges i henhold til en lice...` | Known | Audio system |
| 0x00292BA4 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunh...` | Known | Audio system |
| 0x002958BC | `Dies Audible Software in diesem Produkt wird in Lizenz ...` | Known | Audible audiobook format |
| 0x00295916 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x00295A05 | `.net Codec in diesem Produkt wird mit der Lizenz der Vo...` | Known | Audio system |
| 0x00295ACB | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und ...` | Known | Audio system |
| 0x00298258 | `El software Audible incorporado en este producto se usa...` | Known | Audible audiobook format |
| 0x002982B6 | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x00298455 | `n de audio MPEG Layer-3 utilizada bajo licencia de Frau...` | Known | Audio system |
| 0x0029ACC8 | `Audible-ohjelmistoa t` | Known | Audible audiobook format |
| 0x0029ACFA | `n Audiblen lisenssill` | Known | Audible audiobook format |
| 0x0029AD1F | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x0029ADF1 | `.net codec t` | Known | Audio system |
| 0x0029AE88 | `MPEG Layer-3 -` | Known | Audio system |
| 0x0029AE9A | `nikoodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | Audio system |
| 0x0029E1F8 | `Le logiciel Audible de ce produit est utilis` | Known | Audible audiobook format |
| 0x0029E226 | ` sous licence Audible. Copyright ` | Known | Audible audiobook format |
| 0x0029E249 | ` 2002 par Audible, Inc. Tous drois r` | Known | Audible audiobook format |
| 0x0029E2F0 | ` sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x0029E3C8 | `encodage audio MPEG Layer-3 sous licence de Fraunhofer ...` | Known | Audio system |
| 0x002A0C40 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x002A0C69 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x002A0C98 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x002A0D0A | ` utilizzato su licenza da VoiceAge Corporation. Il code...` | Known | Audio system |
| 0x002A0DE0 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da...` | Known | Audio system |
| 0x002A38E1 | ` Audible ` | Known | Audible audiobook format |
| 0x002A3902 | `Audible ` | Known | Audible audiobook format |
| 0x002A394F | ` Copyright 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002A3B20 | `MPEG Layer-3 ` | Known | Audio system |
| 0x002A3B6C | `Fraunhofer IIS ` | Known | Audio system |
| 0x002A67AE | ` Audible` | Known | Audible audiobook format |
| 0x002A67F1 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002A68E2 | `.net codec` | Known | Audio system |
| 0x002A69A3 | ` Fraunhofer IIS` | Known | Audio system |
| 0x002A97DC | `De Audible-software in dit product wordt gebruikt in li...` | Known | Audible audiobook format |
| 0x002A9833 | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x002A992F | `.net-codec in dit product wordt gebruikt in licentie va...` | Known | Audio system |
| 0x002A99CC | `Technologie voor codering van MPEG Layer-3-audio in lic...` | Known | Audio system |
| 0x002AC0D8 | `Audible-programvaren i dette produktet brukes p` | Known | Audible audiobook format |
| 0x002AC109 | ` lisens fra Audible. Copyright ` | Known | Audible audiobook format |
| 0x002AC12A | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x002AC2B4 | `MPEG Layer-3 lydkodingsteknologi lisensiert fra Fraunho...` | Known | Audio system |
| 0x002AEAD8 | `O software Audible deste produto ` | Known | Audible audiobook format |
| 0x002AEB0D | `a da Audible. Copyright ` | Known | Audible audiobook format |
| 0x002AEB27 | ` 2002 por Audible, Inc. Todos os direitos reservados.` | Known | Audible audiobook format |
| 0x002AECDB | `udio MPEG Layer-3 ` | Known | Audio system |
| 0x002AECEF | ` licenciada da Fraunhofer IIS e THOMSON multimedia.` | Known | Audio system |
| 0x002B14C0 | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x002B14EF | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x002B1506 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x002B16A0 | `Ljudkodningstekniken MPEG Layer-3 ` | Known | Audio system |
| 0x002B16D6 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | Audio system |
| 0x002B3E0D | ` 2002 by Audible, Inc.` | Known | Audible audiobook format |
| 0x002B3F92 | ` Fraunhofer IIS ` | Known | Audio system |
| 0x002B6932 | `Copyright (C) 2002 by Audible, Inc. All rights reserved...` | Known | Audible audiobook format |
| 0x002C726C | `The Audible software in this product is used under lice...` | Known | Audible audiobook format |
| 0x002C72C1 | ` 2004 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002C73A5 | `.net codec in this product is used under license from V...` | Known | Audio system |
| 0x002C7438 | `MPEG Layer-3 audio coding technology licensed from Frau...` | Known | Audio system |
| 0x002E9A51 | ``0aLaCfDf` | Known | Apple Lossless codec |
| 0x002EC604 | `S_CSA_CODECMSG` | Known | Audio system |
| 0x0043F004 | `S21400AAC018470000A0380500203905008438050026` | Known | AAC codec |
| 0x00443E10 | `S21400C680F1E76989081AAC49C97A8842F6DBAD4DEC` | Known | AAC codec |
| 0x0046C004 | `S21401AAC0233405007365080077650800FFFF000062` | Known | AAC codec |
| 0x004835BA | `S214022FA0287813494143C90B0906090E6722511AAC` | Known | AAC codec |
| 0x004895B4 | `S2140251C01200001AACC39FE500B0DCE504005BE306` | Known | AAC codec |
| 0x0048963B | `S2140251F000B0DCE504B01BE20400001AACC11FE5F7` | Known | AAC codec |
| 0x00496CDC | `S214029E409403C1FA6A64B900AACF7E64E40193A6B9` | Known | AAC codec |
| 0x00496DBD | `S214029E908A6662559403C1F17365B900AAC58665E0` | Known | AAC codec |
| 0x0049762D | `S21402A1908AACA756A083B62AB6664B019404CB5661` | Known | AAC codec |
| 0x0049862C | `S21402A740AB039EAACB50CF02AD67DA50F301BC260C` | Known | AAC codec |
| 0x00499004 | `S21402AAC096AB5E5357039B5A6E5312839F0B7F536C` | Known | AAC codec |
| 0x0049D05A | `S21402C1A045AAAC3279C6459EB83279C64592C532A2` | Known | AAC codec |
| 0x004C2B75 | `S2140398107F3A8E3A9D3AAC3ABA3AC93AD83AE63AD9` | Known | AAC codec |
| 0x004C6004 | `S21403AAC0F90E120314FA15F012E500408E088A3ABE` | Known | AAC codec |

---

## 6. Storage/Hardware

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00003800 | `!ATAsoso` | Known | ATA/disk interface |
| 0x00003828 | `!ATAdpua` | Known | ATA/disk interface |
| 0x0000493C | `data abort` | Known | ATA/disk interface |
| 0x00035140 | `diskModeImageRev` | Known | Hardware interface |
| 0x00035318 | `atadpszBoardHwName` | Known | ATA/disk interface |
| 0x0003533C | `pu8FirewireGuid` | Known | FireWire interface |
| 0x00038D84 | `atad\|=` | Known | ATA/disk interface |
| 0x0003AC08 | `atadL?` | Known | ATA/disk interface |
| 0x000773F8 | `atadmhbdmhsd` | Known | ATA/disk interface |
| 0x00147BC9 | `diskmode` | Known | Hardware interface |
| 0x00147BD2 | `diskscan` | Known | Hardware interface |
| 0x001485AC | `TROMResourceDB - unknown header version! (Try regenerat...` | Known | ATA/disk interface |
| 0x0014956A | `SPINDLE SERVO FAILURE` | Known | Hardware interface |
| 0x001497AF | `RECOVERED DATA WITH NO ERROR CORRECTION APPLIED` | Known | ATA/disk interface |
| 0x001497DF | `RECOVERED DATA WITH RETRIES` | Known | ATA/disk interface |
| 0x001497FB | `RECOVERED DATA WITH POSITIVE HEAD OFFSET` | Known | ATA/disk interface |
| 0x00149824 | `RECOVERED DATA WITH NEGATIVE HEAD OFFSET` | Known | ATA/disk interface |
| 0x0014984D | `RECOVERED DATA WITH RETRIES AND/OR CIRC APPLIED` | Known | ATA/disk interface |
| 0x0014987D | `RECOVERED DATA USING PREVIOUS SECTOR ID` | Known | ATA/disk interface |
| 0x001498A5 | `RECOVERED DATA WITHOUT ECC - RECOMMEND REASSIGNMENT` | Known | ATA/disk interface |
| 0x001498D9 | `RECOVERED DATA WITHOUT ECC - RECOMMEND REWRITE` | Known | ATA/disk interface |
| 0x00149908 | `RECOVERED DATA WITHOUT ECC - DATA REWRITTEN` | Known | ATA/disk interface |
| 0x00149934 | `RECOVERED DATA WITH ERROR CORRECTION APPLIED` | Known | ATA/disk interface |
| 0x00149961 | `RECOVERED DATA WITH ERROR CORR. & RETRIES APPLIED` | Known | ATA/disk interface |
| 0x00149993 | `RECOVERED DATA - DATA AUTO-REALLOCATED` | Known | ATA/disk interface |
| 0x001499BA | `RECOVERED DATA WITH CIRC` | Known | ATA/disk interface |
| 0x001499D3 | `RECOVERED DATA WITH L-EC` | Known | ATA/disk interface |
| 0x001499EC | `RECOVERED DATA ` | Known | ATA/disk interface |
| 0x00149A53 | `SYNCHRONOUS DATA TRANSFER ERROR` | Known | ATA/disk interface |
| 0x0014A321 | `INQUIRY DATA HAS CHANGED` | Known | ATA/disk interface |
| 0x0014A430 | `DATA PHASE ERROR` | Known | ATA/disk interface |
| 0x0014AD78 | `DATAPLAY` | Known | ATA/disk interface |
| 0x0014B11C | `BatteryLevel` | Known | Power management |
| 0x0014B129 | `FirewirePower` | Known | FireWire interface |
| 0x0014D322 | `ex_data` | Known | ATA/disk interface |
| 0x0014D470 | `C:\archive\Sobek_172\src\P68\Sources\Services\OpenSSL\c...` | Known | ATA/disk interface |
| 0x0014E037 | `RSA Data Security, Inc.` | Known | ATA/disk interface |
| 0x0014E054 | `RSA Data Security, Inc. PKCS` | Known | ATA/disk interface |
| 0x0014E1B5 | `pkcs7-data` | Known | ATA/disk interface |
| 0x0014E1C0 | `pkcs7-signedData` | Known | ATA/disk interface |
| 0x0014E1D1 | `pkcs7-envelopedData` | Known | ATA/disk interface |
| 0x0014E1E5 | `pkcs7-signedAndEnvelopedData` | Known | ATA/disk interface |
| 0x0014E202 | `pkcs7-digestData` | Known | ATA/disk interface |
| 0x0014E213 | `pkcs7-encryptedData` | Known | ATA/disk interface |
| 0x0014E46C | `nsDataType` | Known | ATA/disk interface |
| 0x0014E477 | `Netscape Data Type` | Known | ATA/disk interface |
| 0x0014F211 | `id-smime-ct-authData` | Known | ATA/disk interface |
| 0x0014F27E | `id-smime-ct-DVCSRequestData` | Known | ATA/disk interface |
| 0x0014F29A | `id-smime-ct-DVCSResponseData` | Known | ATA/disk interface |
| 0x0014FCB2 | `id-cmc-dataReturn` | Known | ATA/disk interface |
| 0x0014FE05 | `id-on-personalData` | Known | ATA/disk interface |
| 0x0014FF06 | `id-cct-PKIData` | Known | ATA/disk interface |
| 0x001506A5 | `qualityLabelledData` | Known | ATA/disk interface |
| 0x00150A80 | `setct-PANData` | Known | ATA/disk interface |
| 0x00150AAB | `setct-OIData` | Known | ATA/disk interface |
| 0x00150AC1 | `setct-PIData` | Known | ATA/disk interface |
| 0x00150ACE | `setct-PIDataUnsigned` | Known | ATA/disk interface |
| 0x00150B49 | `setct-PInitResData` | Known | ATA/disk interface |
| 0x00150B69 | `setct-PResData` | Known | ATA/disk interface |
| 0x00150BBF | `setct-CapTokenData` | Known | ATA/disk interface |
| 0x00150C0D | `setct-AuthRevResData` | Known | ATA/disk interface |
| 0x00150C57 | `setct-CapResData` | Known | ATA/disk interface |
| 0x00150C8F | `setct-CapRevResData` | Known | ATA/disk interface |
| 0x00150CC6 | `setct-CredResData` | Known | ATA/disk interface |
| 0x00150D01 | `setct-CredRevResData` | Known | ATA/disk interface |
| 0x00150D16 | `setct-PCertReqData` | Known | ATA/disk interface |
| 0x00150D3B | `setct-BatchAdminReqData` | Known | ATA/disk interface |
| 0x00150D53 | `setct-BatchAdminResData` | Known | ATA/disk interface |
| 0x00150DAB | `setct-CertReqData` | Known | ATA/disk interface |
| 0x00150DCE | `setct-CertResData` | Known | ATA/disk interface |
| 0x0015115A | `setCext-merchData` | Known | ATA/disk interface |
| 0x001511E5 | `setCext-Track2Data` | Known | ATA/disk interface |
| 0x001513D3 | `set-brand-IATA-ATA` | Known | ATA/disk interface |
| 0x00291E10 | `Spiller nu` | Known | Hardware interface |
| 0x00291EA0 | `Spillelister` | Known | Hardware interface |
| 0x0029263D | `r du har tilsluttet iPod som FireWire-disk, skal du dob...` | Known | FireWire interface |
| 0x0029315C | ` Spillelister` | Known | Hardware interface |
| 0x002954A9 | `nnen die enthaltenen Kontakte im Standard-Format "vCard...` | Known | FireWire interface |
| 0x002956B0 | `nnen die enthaltenen Kontakte im Standard-Format "vCard...` | Known | FireWire interface |
| 0x00297EC1 | `gido FireWire y hacer doble clic en el icono del iPod e...` | Known | FireWire interface |
| 0x0029A2DC | `Diskanttivahv.` | Known | Hardware interface |
| 0x0029A2EC | `Diskanttiheikenn.` | Known | Hardware interface |
| 0x0029A93A | ` tietonsa vakiintuneessa "vCards"-muodossa. Kun iPodin ...` | Known | FireWire interface |
| 0x0029ACB0 | `Yhteystietoja ladataan.` | Known | ATA/disk interface |
| 0x0029B7AC | `Ladataan` | Known | ATA/disk interface |
| 0x0029DE19 | ` votre iPod comme disque dur FireWire, double-cliquez s...` | Known | FireWire interface |
| 0x002A07CB | ` contenere, insieme alla musica, fino a un migliaio di ...` | Known | FireWire interface |
| 0x002A09DB | ` contenere, insieme alla musica, fino a un migliaio di ...` | Known | FireWire interface |
| 0x002A1294 | `Data & Ora` | Known | ATA/disk interface |
| 0x002A1464 | `Disattivata` | Known | ATA/disk interface |
| 0x002A330A | ` FireWire ` | Known | FireWire interface |
| 0x002A9314 | `De iPod biedt ruimte voor maar liefst duizend adressen ...` | Known | FireWire interface |
| 0x002A9544 | `De iPod biedt ruimte voor maar liefst duizend adressen ...` | Known | FireWire interface |
| 0x002AB56C | `Spilles n` | Known | Hardware interface |
| 0x002AB734 | `Mer diskant` | Known | Hardware interface |
| 0x002AB740 | `Mindre diskant` | Known | Hardware interface |
| 0x002ABD97 | `rst aktiverer du iPod som en FireWire-disk, deretter do...` | Known | FireWire interface |
| 0x002AE708 | `gido FireWire, basta clicar duas vezes no ` | Known | FireWire interface |
| 0x002AF2A4 | `Ajustar Data e Hora` | Known | ATA/disk interface |
| 0x002AF3D8 | `Data e Hora` | Known | ATA/disk interface |
| 0x002B115A | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x002B117A | `rddisk. Sedan dubbelklickar du bara p` | Known | Hardware interface |
| 0x002C6E6B | `X's Address Book, Microsoft Entourage, and Palm Desktop...` | Known | FireWire interface |
| 0x002C6FFA | `Your iPod can store up to one thousand contacts right a...` | Known | FireWire interface |
| 0x002C7CF4 | `No battery power remains. Please connect iPod to power.` | Known | Power management |
| 0x002C7D2C | `Low Battery` | Known | Power management |
| 0x002D98AD | `aGa>a(a'aJa?a<a,a4a=aBaDasawaXaYaZakataoaeaqa_a]aSaua` | Known | ATA/disk interface |
| 0x002E3525 | `ataTaza[aea;ajaaaVa)b'b+b+dMd[d]dtdvdrdsd}dudfd` | Known | ATA/disk interface |
| 0x002EA61F | `M_DISKMGR` | Known | Hardware interface |
| 0x002EC516 | `S_GPIO_IR_TMOUT` | Known | GPIO hardware |
| 0x002EC56B | `S_GPIO` | Known | GPIO hardware |
| 0x002EC57C | `S_GPIO_CD_A` | Known | GPIO hardware |
| 0x002EC58D | `S_GPIO_CD_B` | Known | GPIO hardware |
| 0x002EC59E | `S_GPIO_ACK` | Known | GPIO hardware |
| 0x002EC5C0 | `S_I2C_SCAN` | Known | Hardware interface |
| 0x002ECA22 | `S_DISKMGRQ` | Known | Hardware interface |
| 0x002F0184 | `T_FIREWIRE` | Known | FireWire interface |
| 0x002F01B7 | `T_DISKMGR` | Known | Hardware interface |
| 0x002F021D | `T_HOLDSWITCH` | Known | Hardware interface |
| 0x0030FDB8 | `Predictor Data Present, invalid output!` | Known | ATA/disk interface |
| 0x0031A6B4 | `SP_ERR_NULL_DATAINFUNC` | Known | ATA/disk interface |
| 0x0031A6CC | `SP_ERR_NULL_DATAOUTFUNC` | Known | ATA/disk interface |
| 0x0031A830 | `SP_ERR_UNSUPPORTED_DATACLASS` | Known | ATA/disk interface |
| 0x0031A850 | `SP_ERR_UNSUPPORTED_DATATYPE` | Known | ATA/disk interface |
| 0x0031A86C | `SP_ERR_UNSUPPORTED_DATASUBTYPE` | Known | ATA/disk interface |
| 0x0031A9D0 | `SP_ERR_OTHER_DATAINOUT_ERR` | Known | ATA/disk interface |
| 0x0031FB2C | ` Data error sts=%d` | Known | ATA/disk interface |
| 0x0032040D | `cmp_sct_data:  adrx = %d, Wtbuf[adrx] = %x, Rdbuf[adrx]...` | Known | ATA/disk interface |
| 0x00325D34 | `Running on FireWire power.` | Known | FireWire interface |
| 0x00325D54 | `Running on battery power.` | Known | Power management |
| 0x00325E1F | `@Battery reading is %d (%x) E%4cF` | Known | Power management |
| 0x00325EAF | `Switch to disk mode` | Known | Hardware interface |
| 0x00325F17 | `Turn on/off hard disk` | Known | Hardware interface |
| 0x00325F57 | `Turn on/off firewire` | Known | FireWire interface |
| 0x00325F77 | `Turn on/off battery charging` | Known | Power management |
| 0x003271D8 | `!ATADisk is FDISK format.` | Known | ATA/disk interface |
| 0x003271F4 | `Disk is Mac format.` | Known | Hardware interface |
| 0x0032720C | `Disk contains %d interesting partitions.` | Known | Hardware interface |
| 0x00327408 | `Write a boot record when formatting FDISK.` | Known | Hardware interface |
| 0x00327531 | `Spin down hard drive now.` | Known | Hardware interface |
| 0x00329AC1 | `<BATTERY MEASUREMENT [ %d.%d , 0x%X ]>` | Known | Power management |
| 0x00329B09 | `<DISK-MODE  V%X.%X.%X>` | Known | Hardware interface |
| 0x0032A618 | `Disk mode` | Known | Hardware interface |
| 0x0032A624 | `Disk Scan mode` | Known | Hardware interface |
| 0x003F328F | `@diskmode` | Known | Hardware interface |
| 0x003F8C70 | `F.FIREWIRE` | Known | FireWire interface |
| 0x0040CC5C | `Stand-alone Disk-Mode running` | Known | Hardware interface |
| 0x00412D2C | `DMA ABRT` | Known | Hardware interface |
| 0x00412D38 | `DMA ABR1` | Known | Hardware interface |
| 0x00412D44 | `DMA ABR2` | Known | Hardware interface |
| 0x00412D50 | `DMA ABR3` | Known | Hardware interface |
| 0x004149D3 | `@BUSRDMA` | Known | Hardware interface |
| 0x004149DC | `BUSWDMA` | Known | Hardware interface |
| 0x004149EC | `BUSDMA2` | Known | Hardware interface |
| 0x004149F4 | `BUSDMA3` | Known | Hardware interface |
| 0x004149FC | `BUSDMA4` | Known | Hardware interface |
| 0x00417F20 | `SDMACKST` | Known | Hardware interface |
| 0x00418004 | `ATACKST1` | Known | ATA/disk interface |
| 0x00418010 | `ATACKST2` | Known | ATA/disk interface |
| 0x0041801F | `@ATACKSTE` | Known | ATA/disk interface |
| 0x0041802F | `@ATACKST3` | Known | ATA/disk interface |
| 0x0041820C | `RDMACBAK` | Known | Hardware interface |
| 0x004184EC | `WDMACBAK` | Known | Hardware interface |

---

## 7. Error Messages

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009EF94 | `Device failed read of mode page 2a.` | Known | Error/assertion message |
| 0x0009FED8 | `UNKNOWN ERROR (%06x)` | Known | Error/assertion message |
| 0x000A017C | `Invalid session type for open disc.` | Known | Error/assertion message |
| 0x000A01A0 | `Invalid session type.` | Known | Error/assertion message |
| 0x000A01B8 | `Error reading write parameters page from drive.` | Known | Error/assertion message |
| 0x000A01E8 | `Error sending write parameters page to drive.` | Known | Error/assertion message |
| 0x000A0240 | `Close session failed` | Known | Error/assertion message |
| 0x000C6ABC | `SetTAOWriteMode: Error reading write parameters page.` | Known | Error/assertion message |
| 0x000C6B18 | `SetTAOWriteMode: Error sending write parameters page.` | Known | Error/assertion message |
| 0x000C6D84 | `Invalid track number.` | Known | Error/assertion message |
| 0x000F60F8 | `%s(%d): OpenSSL internal error, assertion failed: %s` | Known | Error/assertion message |
| 0x001050D4 | `error:%08lX:%s:%s:%s` | Known | Error/assertion message |
| 0x0011AA30 | `internal error: list index %ld out of range` | Known | Error/assertion message |
| 0x0011D9C4 | `*** assertion failed: %s, file %s, line %d` | Known | Error/assertion message |
| 0x00123D2C | `Invalid Operation` | Known | Error/assertion message |
| 0x001489C3 | `load failed (%d)` | Known | Error/assertion message |
| 0x00148E40 | `tInvalid partition table. Setup cannot continue.` | Known | Error/assertion message |
| 0x00148E71 | `Error loading operating system. Setup cannot continue.` | Known | Error/assertion message |
| 0x00149138 | `PSM_Test() Error` | Known | Error/assertion message |
| 0x00149294 | `AUDIO PLAY OPERATION STOPPED DUE TO ERROR` | Known | Error/assertion message |
| 0x00149501 | `LOGICAL UNIT COMMUNICATION PARITY ERROR` | Known | Error/assertion message |
| 0x00149529 | `TRACK FOLLOWING ERROR` | Known | Error/assertion message |
| 0x00149592 | `ERROR LOG OVERFLOW` | Known | Error/assertion message |
| 0x001495FF | `WRITE ERROR - RECOVERY NEEDED` | Known | Error/assertion message |
| 0x0014961D | `WRITE ERROR - RECOVERY FAILED` | Known | Error/assertion message |
| 0x0014963B | `WRITE ERROR - LOSS OF STREAMING` | Known | Error/assertion message |
| 0x0014965B | `WRITE ERROR - PADDING BLOCKS ADDED` | Known | Error/assertion message |
| 0x0014967E | `UNRECOVERED READ ERROR` | Known | Error/assertion message |
| 0x001496AC | `ERROR TOO LONG TO CORRECT` | Known | Error/assertion message |
| 0x001496C6 | `L-EC UNCORRECTABLE ERROR` | Known | Error/assertion message |
| 0x001496DF | `CIRC UNRECOVERED ERROR` | Known | Error/assertion message |
| 0x001496F6 | `ERROR READING UPC/EAN NUMBER` | Known | Error/assertion message |
| 0x00149713 | `ERROR READING ISRC NUMBER` | Known | Error/assertion message |
| 0x0014972D | `READ ERROR - LOSS OF STREAMING` | Known | Error/assertion message |
| 0x0014974C | `RANDOM POSITIONING ERROR` | Known | Error/assertion message |
| 0x00149765 | `MECHANICAL POSITIONING ERROR` | Known | Error/assertion message |
| 0x00149782 | `POSITIONING ERROR DETECTED BY READ OF MEDIUM` | Known | Error/assertion message |
| 0x00149A37 | `PARAMETER LIST LENGTH ERROR` | Known | Error/assertion message |
| 0x00149A96 | `INVALID COMMAND OPERATION CODE` | Known | Error/assertion message |
| 0x00149AD8 | `INVALID ELEMENT ADDRESS` | Known | Error/assertion message |
| 0x00149AF0 | `INVALID FIELD IN CDB` | Known | Error/assertion message |
| 0x00149B2A | `INVALID FIELD IN PARAMETER LIST` | Known | Error/assertion message |
| 0x00149B62 | `PARAMETER VALUE INVALID` | Known | Error/assertion message |
| 0x00149B9D | `INVALID RELEASE OF ACTIVE PERSISTENT RESERVATION` | Known | Error/assertion message |
| 0x00149DA6 | `COPY CANNOT EXECUTE SINCE INITIATOR CANNOT DISCONNECT` | Known | Error/assertion message |
| 0x00149DDC | `COMMAND SEQUENCE ERROR` | Known | Error/assertion message |
| 0x00149E77 | `CANNOT READ MEDIUM ` | Known | Error/assertion message |
| 0x00149EE1 | `CANNOT WRITE MEDIUM ` | Known | Error/assertion message |
| 0x00149F30 | `CANNOT FORMAT MEDIUM ` | Known | Error/assertion message |
| 0x00149FD4 | `FORMAT COMMAND FAILED` | Known | Error/assertion message |
| 0x00149FEA | `ZONED FORMATTING FAILED DUE TO SPARE LINKING` | Known | Error/assertion message |
| 0x0014A221 | `MECHANICAL POSITIONING OR CHANGER ERROR` | Known | Error/assertion message |
| 0x0014A249 | `INVALID BITS IN IDENTIFY MESSAGE` | Known | Error/assertion message |
| 0x0014A3C0 | `SCSI PARITY ERROR` | Known | Error/assertion message |
| 0x0014A3D2 | `INITIATOR DETECTED ERROR MESSAGE RECEIVED` | Known | Error/assertion message |
| 0x0014A3FC | `INVALID MESSAGE ERROR.BSR NCITS` | Known | Error/assertion message |
| 0x0014A41C | `COMMAND PHASE ERROR` | Known | Error/assertion message |
| 0x0014A441 | `LOGICAL UNIT FAILED SELF-CONFIGURATION` | Known | Error/assertion message |
| 0x0014A4C0 | `MEDIA LOAD OR EJECT FAILED` | Known | Error/assertion message |
| 0x0014A7C4 | `INVALID PACKET SIZE` | Known | Error/assertion message |
| 0x0014A906 | `LOGICAL UNITREGION MUST BE PERMANENT/REGION RESET COUNT...` | Known | Error/assertion message |
| 0x0014A944 | `SESSION FIXATION ERROR` | Known | Error/assertion message |
| 0x0014A95B | `SESSION FIXATION ERROR WRITING LEAD-IN` | Known | Error/assertion message |
| 0x0014A982 | `SESSION FIXATION ERROR WRITING LEAD-OUT` | Known | Error/assertion message |
| 0x0014A9AA | `SESSION FIXATION ERROR ` | Known | Error/assertion message |
| 0x0014AA2C | `CD CONTROL ERROR` | Known | Error/assertion message |
| 0x0014AA7F | `POWER CALIBRATION AREA ERROR` | Known | Error/assertion message |
| 0x0014AB40 | `INVALID MESSAGE ERROR` | Known | Error/assertion message |
| 0x0014ABA9 | `CD CONTROL ERROR.BSR NCITS` | Known | Error/assertion message |
| 0x0031A468 | `SP_ERR_INIT_FAILED` | Known | Error/assertion message |
| 0x0031A734 | `SP_ERR_NULL_ERRORFUNC` | Known | Error/assertion message |
| 0x0031A950 | `SP_ERR_INVALID_CHANNEL_ID` | Known | Error/assertion message |
| 0x0031A96C | `SP_ERR_INVALID_COMMAND` | Known | Error/assertion message |
| 0x0031AB10 | `SP_ERR_HW_RESET_FAILED` | Known | Error/assertion message |
| 0x0031E1C3 | `( <- Error:Unsupported` | Known | Error/assertion message |
| 0x0031E1DC | ` <- Error:Internal Error` | Known | Error/assertion message |
| 0x0031E207 | `(Invalid Key Value arguments` | Known | Error/assertion message |
| 0x0031E7D4 | `***Error r_reg Unknown Ide Reg(%X)` | Known | Error/assertion message |
| 0x0031E8B0 | `***Error w_reg Unknown Ide Reg(%X)` | Known | Error/assertion message |
| 0x0031F0A5 | `command:  head No. error` | Known | Error/assertion message |
| 0x0031F0C5 | `command:  DRVH reg. compare error` | Known | Error/assertion message |
| 0x0031F0ED | `command:  CYL_high reg. compare error` | Known | Error/assertion message |
| 0x0031F119 | `command:  CYL_low reg. compare error` | Known | Error/assertion message |
| 0x0031F145 | `command:  SECT_num reg. compare error` | Known | Error/assertion message |
| 0x0031F171 | `command:  SECT_cnt reg. compare error` | Known | Error/assertion message |
| 0x0031FCF9 | `xfer_command: command failed.  sts=%x` | Known | Error/assertion message |
| 0x0031FF8D | `FW transfer error status_reg=%x` | Known | Error/assertion message |
| 0x0031FFDD | `Check FW error status_reg=%x` | Known | Error/assertion message |
| 0x00320179 | `FW Write error status_reg=%x` | Known | Error/assertion message |
| 0x00320199 | `DoToshiba: stand-by immediate command error status_reg=...` | Known | Error/assertion message |
| 0x00320211 | `Write Copyright error status_reg=%x` | Known | Error/assertion message |
| 0x00320261 | `Restart error status_reg=%x` | Known | Error/assertion message |
| 0x0032034D | `doToshiba:  sts error.  sts = %x` | Known | Error/assertion message |
| 0x00322889 | `get_fw_rev:  identify command failed.  cannot get FW Re...` | Known | Error/assertion message |
| 0x003228D5 | `check_fw:  buffer compare error.  i = %x, Wtbuf[i] = %x...` | Known | Error/assertion message |
| 0x00325D70 | `Power status is invalid.` | Known | Error/assertion message |
| 0x0032660F | `@Cannot set: no valid time given on command line.` | Known | Error/assertion message |
| 0x003271C0 | `Error opening device.` | Known | Error/assertion message |
| 0x00327354 | `Device %s error reading sectors.` | Known | Error/assertion message |
| 0x003289F4 | `***Error: [0x%08X]==0x%08X` | Known | Error/assertion message |
| 0x003F7321 | `Vccw Range Error` | Known | Error/assertion message |
| 0x003F7335 | `Device Protect Error` | Known | Error/assertion message |
| 0x003F734D | `Command Sequence Error` | Known | Error/assertion message |
| 0x003F7365 | `Clear Block Lock-Bits Error` | Known | Error/assertion message |
| 0x003F7445 | `      Vccw Range Error` | Known | Error/assertion message |
| 0x003F747D | `      Command Sequence Error` | Known | Error/assertion message |
| 0x003F749D | `      Block Erase Error` | Known | Error/assertion message |
| 0x003F7621 | `      Device Protect Error` | Known | Error/assertion message |
| 0x003F763D | `      Byte Write Error` | Known | Error/assertion message |
| 0x003F76F5 | `Set Block Lock-Bits Error` | Known | Error/assertion message |
| 0x003FB6BE | `XMODEM/XMODEM-1K transfer failed.` | Known | Error/assertion message |

---

## 8. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00049352 | `  # Name             Avail/Total Worst Count  Bytes Wai...` | Known | Filesystem path |
| 0x00049398 | `%3d %-16s %5d/%5ld %5d %5d %6ld` | Known | Filesystem path |
| 0x00049A5A | `  # Name            Current/Depth Worst Count Waiters` | Known | Filesystem path |
| 0x00049A94 | `%3d %-16s  %5d/%5d %5d %5d` | Known | Filesystem path |
| 0x0004A198 | ` %4p/%4p ticks` | Known | Filesystem path |
| 0x0004A5C0 | `C - Clock / Timers` | Known | Filesystem path |
| 0x0004A5EC | `Z - Zero Partition/Queue/Resource Statistics` | Known | Filesystem path |
| 0x0004AFE4 | `/ - Time slice` | Known | Filesystem path |
| 0x0009E790 | `             ROM Timing (R/W) : %1X/%1X` | Known | Filesystem path |
| 0x0009E7D8 | `             RAM Timing (R/W) : %1X/%1X` | Known | Filesystem path |
| 0x000D3DBC | `SASUC -D5S00A/` | Known | Filesystem path |
| 0x000D7A08 | `S/P-DIF` | Known | Filesystem path |
| 0x000DAFC8 | `  <VolumeIndex>%s</VolumeIndex>` | Known | Filesystem path |
| 0x000DAFEC | `  <Filename>%s</Filename>` | Known | Filesystem path |
| 0x000DB014 | ` </MusicItem>` | Known | Filesystem path |
| 0x000DB024 | `</GetMusicItemResponse>` | Known | Filesystem path |
| 0x000DB040 | `</GetMusicItemExResponse>` | Known | Filesystem path |
| 0x000DB1E8 | `</SetMusicItemResponse>` | Known | Filesystem path |
| 0x00145C15 | `url;type=work:apple.com/support/ipod` | Known | Filesystem path |
| 0x00148D28 | `\/:*?"<>\|` | Known | Filesystem path |
| 0x00149214 | `I/O PROCESS TERMINATED` | Known | Filesystem path |
| 0x0014AADB | `RMA/PMA IS FULL.BSR NCITS` | Known | Filesystem path |
| 0x0014B847 | `</getmusicitem>` | Known | Filesystem path |
| 0x0014B868 | `</getmusicitemex>` | Known | Filesystem path |
| 0x0014B889 | `</setmusicitem>` | Known | Filesystem path |
| 0x0014B8A6 | `</collection>` | Known | Filesystem path |
| 0x0014B8C0 | `</musicitem>` | Known | Filesystem path |
| 0x0014B8DB | `</volumeindex>` | Known | Filesystem path |
| 0x0014B8F5 | `</filename>` | Known | Filesystem path |
| 0x0014B90D | `</songindex>` | Known | Filesystem path |
| 0x0014B922 | `</title>` | Known | Filesystem path |
| 0x0014B936 | `</filesize>` | Known | Filesystem path |
| 0x0014B94F | `</timelength>` | Known | Filesystem path |
| 0x0014B967 | `</bitrate>` | Known | Filesystem path |
| 0x0014B984 | `</variablebitrate>` | Known | Filesystem path |
| 0x0014B9A6 | `</samplingrate>` | Known | Filesystem path |
| 0x0014B9C3 | `</samplesize>` | Known | Filesystem path |
| 0x0014B9E4 | `</numberofchannels>` | Known | Filesystem path |
| 0x0014BA1F | `</artist>` | Known | Filesystem path |
| 0x0014BA31 | `</album>` | Known | Filesystem path |
| 0x0014BA41 | `</year>` | Known | Filesystem path |
| 0x0014BA51 | `</genre>` | Known | Filesystem path |
| 0x0014BA68 | `</tracknumber>` | Known | Filesystem path |
| 0x0014BA86 | `</trackcredits>` | Known | Filesystem path |
| 0x0014BAA3 | `</tracknotes>` | Known | Filesystem path |
| 0x0014BABC | `</comments>` | Known | Filesystem path |
| 0x0014D85C | `You need to read the OpenSSL FAQ, http://www.openssl.or...` | Known | Filesystem path |
| 0x0014EF31 | `S/MIME Capabilities` | Known | Filesystem path |
| 0x001698C3 | `W/}lE>q` | Known | Filesystem path |
| 0x00198505 | `H."0*Bx/` | Known | Filesystem path |
| 0x0019FE62 | `U/~RERT` | Known | Filesystem path |
| 0x001A4BCA | `TUOPT/\|` | Known | Filesystem path |
| 0x001ABE5F | `HuGZp/$j` | Known | Filesystem path |
| 0x001B236F | `(bJ)b"(b2""z/` | Known | Filesystem path |
| 0x001B4ADF | `JUAPDD(/` | Known | Filesystem path |
| 0x001BADEE | `/B\|$BD'` | Known | Filesystem path |
| 0x001BBD6B | `$Bd$BT/` | Known | Filesystem path |
| 0x001C2043 | `/" +J\|!` | Known | Filesystem path |
| 0x001C8EC2 | `Fb""")/` | Known | Filesystem path |
| 0x001CA079 | `/RyO(UIH` | Known | Filesystem path |
| 0x001CB129 | `~$Bh'"~$Bz/` | Known | Filesystem path |
| 0x001CE993 | `$B +BZ/` | Known | Filesystem path |
| 0x001D63B1 | `0c(HBP/` | Known | Filesystem path |
| 0x001DA947 | `$B~("\|/` | Known | Filesystem path |
| 0x001EEA49 | `T/DDDDD` | Known | Filesystem path |
| 0x001EECB3 | `"~UeB /` | Known | Filesystem path |
| 0x001F1765 | `$B((B /` | Known | Filesystem path |
| 0x001F9848 | ` "\|$B~/` | Known | Filesystem path |
| 0x001FC9C4 | `@$B\|$"(/` | Known | Filesystem path |
| 0x001FDAC4 | `)"8/B""` | Known | Filesystem path |
| 0x001FE20C | `r4c6 bN/` | Known | Filesystem path |
| 0x0020360D | `RDT%B(/` | Known | Filesystem path |
| 0x00204799 | `RBHUE\|/` | Known | Filesystem path |
| 0x0020BE19 | `]B""B</` | Known | Filesystem path |
| 0x0020F696 | `,B\|RED/` | Known | Filesystem path |
| 0x00215119 | `$BT). /` | Known | Filesystem path |
| 0x00216319 | `#"TUB(/` | Known | Filesystem path |
| 0x0022FB1A | `\|($D/DB)$D%DD ` | Known | Filesystem path |
| 0x0023368D | `/" %BD"` | Known | Filesystem path |
| 0x00242698 | `ODD""(/` | Known | Filesystem path |
| 0x002438D3 | `B"$R%"B$" /` | Known | Filesystem path |
| 0x00244240 | `bG\|jG\|/` | Known | Filesystem path |
| 0x00245E2A | `$E$$BR/` | Known | Filesystem path |
| 0x00245ECB | `dRB~RA$/` | Known | Filesystem path |
| 0x0024681C | `TT&T%B(/` | Known | Filesystem path |
| 0x002651DD | `B$rD$"D/RP` | Known | Filesystem path |
| 0x002737E2 | `BDU!BJ ""/` | Known | Filesystem path |
| 0x00274746 | `Z-bD("(%B>/` | Known | Filesystem path |
| 0x00286984 | ` !"#$%&'()*+,-./0123456789:;<=>?@abcdefghijklmnopqrstuv...` | Known | Filesystem path |
| 0x00286A84 | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x00292A7B | `rende VoiceAge Corporation i USA og/eller andre lande o...` | Known | Filesystem path |
| 0x0029594F | ` ist entweder eine eingetragene Marke oder eine Marke d...` | Known | Filesystem path |
| 0x00295E80 | `USA/Ost (DST)` | Known | Filesystem path |
| 0x00295E90 | `USA/Ost` | Known | Filesystem path |
| 0x00295E98 | `USA/Zentral (DST)` | Known | Filesystem path |
| 0x00295EAC | `USA/Zentral` | Known | Filesystem path |
| 0x00295EB8 | `USA/Rockies (DST)` | Known | Filesystem path |
| 0x00295ECC | `USA/Rockies` | Known | Filesystem path |
| 0x00295ED8 | `USA/Pazifik (DST)` | Known | Filesystem path |
| 0x00295EEC | `USA/Pazifik` | Known | Filesystem path |
| 0x00295EF8 | `USA/Alaska (DST)` | Known | Filesystem path |
| 0x00295F0C | `USA/Alaska` | Known | Filesystem path |
| 0x00295F18 | `USA/Hawaii (DST)` | Known | Filesystem path |
| 0x00295F2C | `USA/Hawaii` | Known | Filesystem path |
| 0x002982F3 | ` es una marca registrada o una marca comercial de Voice...` | Known | Filesystem path |
| 0x0029A38C | `%d / %d` | Known | Filesystem path |
| 0x0029AD91 | `ity tavaramerkki Yhdysvalloissa ja/tai muissa maissa, j...` | Known | Filesystem path |
| 0x002A28C8 | `%b/%-d %-I:%M %2p` | Known | Filesystem path |
| 0x002A28DC | `%-m/%-d` | Known | Filesystem path |
| 0x002A2904 | `%y/%-m/%-d` | Known | Filesystem path |
| 0x002A2910 | `%Y/%b/%-d` | Known | Filesystem path |
| 0x002A5A18 | `%Y/%B/%d %A  %I:%M:%S %p` | Known | Filesystem path |
| 0x002A5A34 | `%Y/%B/%d` | Known | Filesystem path |
| 0x002A5A4C | `%-m-/%-d %-I:%M %p` | Known | Filesystem path |
| 0x002A5A80 | `%Y/-%-m/-%-d` | Known | Filesystem path |
| 0x002A986B | ` is hetzij een gedeponeerd handelsmerk hetzij een hande...` | Known | Filesystem path |
| 0x002AC167 | ` er enten et registrert varemerke eller et varemerke fo...` | Known | Filesystem path |
| 0x002AEB8B | `brica da VoiceAge Corporation nos Estados Unidos e/ou e...` | Known | Filesystem path |
| 0x002AEBE9 | `a da VoiceAge Corporation. O codificador/decodificador ...` | Known | Filesystem path |
| 0x002B08DC | `%-d/%-m` | Known | Filesystem path |
| 0x002B158C | `r VoiceAge Corporation i USA och/eller andra l` | Known | Filesystem path |
| 0x002B1BA0 | `ll in datum/tid` | Known | Filesystem path |
| 0x002B31AC | `%-m/%-d %-I:%M %p` | Known | Filesystem path |
| 0x002B31E4 | `%Y/%-m/%-d` | Known | Filesystem path |
| 0x002C6598 | `%-m/%-d/%y` | Known | Filesystem path |
| 0x002C72F7 | ` is either registered trademark or trademark of VoiceAg...` | Known | Filesystem path |
| 0x002D793F | `%#%3%+%;%K% %/%(%7%?%` | Known | Filesystem path |
| 0x002D7ABE | `02*2+2,2-2.2/2@272B2C292:212>24222;2623252<2=2?282` | Known | Filesystem path |
| 0x002D7FE9 | `\7_J_/`P`m`` | Known | Filesystem path |
| 0x002D81C5 | `OHSIT>T/Z` | Known | Filesystem path |
| 0x002D823B | `i_l*mim/n` | Known | Filesystem path |
| 0x002D872D | `N,p]u/f` | Known | Filesystem path |
| 0x002D9439 | `S#S/S1S3S8S@SFSES` | Known | Filesystem path |
| 0x002D9539 | `q4V6V2V8VkVdV/VlVjV` | Known | Filesystem path |
| 0x002D97E1 | `_)_-_8_A_H_L_N_/_Q_V_W_Y_a_m_s_w_` | Known | Filesystem path |
| 0x002D9EA7 | `s4s/s)s%s>sNsOs` | Known | Filesystem path |
| 0x002DB383 | `% %!%"%#%$%%%&%'%(%)%*%+%,%-%.%/%0%1%2%3%4%5%6%7%8%9%:%...` | Known | Filesystem path |
| 0x002DBB63 | `qWlIl/Ymg*` | Known | Filesystem path |
| 0x002DD3D1 | `X)W,W*W3W9W.W/W\W;WBWiW` | Known | Filesystem path |
| 0x002DDA9F | `n/o6oKoto*o` | Known | Filesystem path |
| 0x002DE037 | `q/q1qsq\qhqEqrqJqxqzq` | Known | Filesystem path |
| 0x002DEDAB | `b6bKbNb/e` | Known | Filesystem path |
| 0x002DEF61 | `V3W0W(W-W,W/W)W` | Known | Filesystem path |
| 0x002DF035 | `NMOOOGOWO^O4O[OUO0OPOQO=O:O8OCOTO<OFOcO\O`O/ONO6OYO]OHO...` | Known | Filesystem path |
| 0x002DF693 | `e%f-f f'f/f` | Known | Filesystem path |
| 0x002DF995 | `bNc>c/cUcBcFcOcIc:cPc=c*c+c(cMcLcHeIe` | Known | Filesystem path |
| 0x002DFC47 | `fFUjUfUDU^UaUCUJU1UVUOUUU/UdU8U.U\U,UcU3UAUWU` | Known | Filesystem path |
| 0x002DFEAD | `\|F}C}q}.}9}<}@}0}3}D}/}B}2}1}=` | Known | Filesystem path |
| 0x002E0035 | `W/X*X4X$X0X1X!X` | Known | Filesystem path |
| 0x002E014B | `k.l/l,l/n8nTn!n2ngnJn n%n#n` | Known | Filesystem path |
| 0x002E01AD | `r6s%s4s)s:t*t3t"t%t5t6t4t/t` | Known | Filesystem path |
| 0x002E0A2B | `S.V;V9V2V?V4V)VSVNVWVtV6V/V0V` | Known | Filesystem path |
| 0x002E1071 | `\|.~>~F~7~2~C~+~=~1~E~A~4~9~H~5~?~/~D` | Known | Filesystem path |
| 0x002E2061 | `P P'P5P/P1P` | Known | Filesystem path |
| 0x002E21B3 | `h5h+h-h/hNhDh4h` | Known | Filesystem path |
| 0x002E21C7 | `h&h(h.hMh:h%h h,k/k-k1k4kmk` | Known | Filesystem path |
| 0x002E22D3 | `w"w'w#x,x"x5x/x(x.x+x!x)x3x*x1xTy[yOy\ySyRyQy` | Known | Filesystem path |
| 0x002E275D | `v9w/w-w1w2w4w3w=w%w;w5wHxRxIxMxJxLx&xExPxdygyiyjycykyay` | Known | Filesystem path |
| 0x002E27B9 | `{1{+{-{/{2{8{` | Known | Filesystem path |
| 0x002E2C4D | `t/uoulu` | Known | Filesystem path |
| 0x002E2FB3 | ``IaJa+aEa6a2a.aFa/aOa)a@a bh` | Known | Filesystem path |
| 0x002E2FF9 | `d&d0d(dAd5d/d` | Known | Filesystem path |
| 0x002E3061 | `hxi4iii@ioiDiviXiAitiLi;iKi7i\iOiQi2iRi/i{i<iFkEkCkBkHk...` | Known | Filesystem path |
| 0x002E36CB | `y+zJz0z/z(z&z` | Known | Filesystem path |
| 0x002E3FAF | `u/v-v1v=v3v<v5v2v0v` | Known | Filesystem path |
| 0x002E42A5 | `X-[%[2[#[,['[&[/[.[{[` | Known | Filesystem path |
| 0x002E4BD3 | `j<p5p/p7p4p1pBp8p?p:p9p@p;p3pAp` | Known | Filesystem path |
| 0x002E4DD5 | `w-y1y/yT\|S\|` | Known | Filesystem path |
| 0x002E785B | `h>kLp/t` | Known | Filesystem path |
| 0x002E7EE5 | `o;v/}7~` | Known | Filesystem path |
| 0x002E8CA9 | `e1f/h\q6z` | Known | Filesystem path |
| 0x002E92F5 | `UuX/c"dIfKfmh` | Known | Filesystem path |
| 0x002F3352 | `/ 0 1 2 3 4 5 6 7 8 9 : ; < = > ? @ A B C D E F G H I J...` | Known | Filesystem path |
| 0x002F3533 | `! !!!"!#!$!%!&!'!(!)!*!+!,!-!.!/!0!1!2!3!4!5!6!7!8!9!:!...` | Known | Filesystem path |
| 0x002FC964 | `/1f;{1Q` | Known | Filesystem path |
| 0x00325844 | `Syntax: %s [FourCharKeyName] [val0] [val1 val2 val3] [/...` | Known | Filesystem path |
| 0x00325F37 | `Turn on/off backlight` | Known | Filesystem path |
| 0x00325F98 | `/u[012]` | Known | Filesystem path |
| 0x00325FA0 | `Set audio to off/standby/on` | Known | Filesystem path |
| 0x00326071 | `/w warm reboot instead of cold` | Known | Filesystem path |
| 0x003261FC | `/v[0-9]` | Known | Filesystem path |
| 0x00326644 | `%04d/%02d/%02d %02d %02d:%02d:%02d` | Known | Filesystem path |
| 0x003267CC | `Syntax: %s [/i /o /d] [srcStreamName] ` | Known | Filesystem path |
| 0x00326CC8 | `Syntax : %s [+\|-\|HexAddress] [HexSize] [/b\|/w\|/d]` | Known | Filesystem path |
| 0x00326E70 | `Syntax : %s HexAddress HexVal [/b\|/w\|/d]` | Known | Filesystem path |
| 0x00327BE8 | `%x bytes read/written.` | Known | Filesystem path |
| 0x00327C1F | ` Read/write files on the host.` | Known | Filesystem path |
| 0x003283D8 | `Syntax : %s $String\|/p` | Known | Filesystem path |
| 0x00328597 | `/? option gives command specific help (i.e., dir /?)` | Known | Filesystem path |
| 0x00329090 | `</MM_NODE>` | Known | Filesystem path |
| 0x003296FC | `</EXPORTS>` | Known | Filesystem path |
| 0x0032975C | `</IMPORTS>` | Known | Filesystem path |
| 0x003297E0 | `</HANDLE` | Known | Filesystem path |
| 0x00329840 | `</APIMGR>` | Known | Filesystem path |
| 0x00329A91 | `<BOARD RTC ADJUST         %c%d.%02d SEC/DAY>` | Known | Filesystem path |
| 0x00329BC9 | `</ADDRESS RANGES>` | Known | Filesystem path |
| 0x00329BDC | `</SYSINFO>` | Known | Filesystem path |
| 0x00333F3C | `Syntax: %s [/s] [year/month/dom [GMToffset] hour:minute...` | Known | Filesystem path |
| 0x0033434C | `{{~~  /-----\   {{~~ /       \  {{~~\|         \| {{~~\...` | Known | Filesystem path |
| 0x003F4DAA | `Begin XMODEM/XMODEM-1K upload of IRAM image file..........` | Known | Filesystem path |
| 0x003FB6E1 | `Restart XMODEM/XMODEM-1K upload of image file...` | Known | Filesystem path |
| 0x004157D8 | `LBA/CNT` | Known | Filesystem path |

---

## 9. Binary Structure (for Ghidra/IDA)

| Parameter | Value |
|-----------|-------|
| **Architecture** | ARM (32-bit, little-endian) |
| **Processor** | ARM7TDMI (ARMv4T), dual-core |
| **Base Address** | 0x00000000 |
| **Entry Point** | 0x00000000 |
| **Endianness** | Little-Endian |
| **File Size** | 5,066,752 bytes |

