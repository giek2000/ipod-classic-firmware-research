# iPod 1st Generation - RetailOS 1.1.5 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.1.5 |
| **IPSW** | iPod_1.1.5.ipsw |
| **Device** | iPod 1st Generation (2001, Scroll Wheel) |
| **UpdaterFamilyID** | 1 |
| **Binary Size** | 5,066,752 bytes (4.83 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 5,066,752 bytes |
| **Total Strings (>=6)** | 26,630 |
| **Function Prologues** | 7,696 (ARM: 6,885, Thumb: 811) |
| **SoC** | PortalPlayer PP5002 |
| **Architecture** | ARM7TDMI (ARMv4T), dual-core |
| **Encrypted** | No |
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
| 0x002EC47D | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x002F0085 | `T_RTXCBUG` | Hidden | Developer Tool |
| 0x00326A94 | `BTM Debug Zones %s (0x%08X)` | Hidden | Debug/Diagnostic |
| 0x0032A634 | `Retail mode` | Hidden | Demo/Retail Mode |
| 0x0032A644 | `Debug mode` | Hidden | Debug/Diagnostic |

---

## 2. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00009CE8 | `HostOSTask` | Known | RTOS task thread |
| 0x0002C88C | `RunTestsTask` | Known | RTOS task thread |
| 0x00048A8B | `** Stack Snapshot **` | Known | RTOS task thread |
| 0x00048B0E | `RTXC Kernel           %08p %5d %5d %5d` | Known | RTOS task thread |
| 0x00048D5B | `** Clock Snapshot **` | Known | RTOS task thread |
| 0x00049197 | `** Mailbox Snapshot **` | Known | RTOS task thread |
| 0x00049333 | `** Partition Snapshot **` | Known | RTOS task thread |
| 0x000494DB | `** Resource Snapshot **` | Known | RTOS task thread |
| 0x000497CF | `** Semaphore Snapshot **` | Known | RTOS task thread |
| 0x00049A3F | `** Queue Snapshot **` | Known | RTOS task thread |
| 0x00049E77 | `** Task Snapshot **` | Known | RTOS task thread |
| 0x0004A82C | `  K - RTXC` | Known | RTOS task thread |
| 0x0004AAC8 | `Undefined Task` | Known | RTOS task thread |
| 0x0004AB32 | `** Task Register Snapshot **` | Known | RTOS task thread |
| 0x0007C554 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x00145284 | `RTXC v3.2fpp for ARM and Thumb - ARM ADS 1.0 Jul-08-00 Key: 24104` | Known | RTOS task thread |

---

## 3. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0014B9F8 | `<codec>` | Known | Audio system |
| 0x0014BA00 | `</codec>` | Known | Audio system |
| 0x0014EBE3 | `msCodeCom` | Known | Audio system |
| 0x00292AFA | `.net codec i dette produkt bruges i henhold til en licensaftale fra VoiceAge Cor` | Known | Audio system |
| 0x00295A05 | `.net Codec in diesem Produkt wird mit der Lizenz der VoiceAge Corporation verwen` | Known | Audio system |
| 0x0029ADF1 | `.net codec t` | Known | Audio system |
| 0x0029E2F0 | ` sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x002A0D0A | ` utilizzato su licenza da VoiceAge Corporation. Il codec ACELP` | Known | Audio system |
| 0x002A68E2 | `.net codec` | Known | Audio system |
| 0x002A992F | `.net-codec in dit product wordt gebruikt in licentie van VoiceAge Corporation. G` | Known | Audio system |
| 0x002C73A5 | `.net codec in this product is used under license from VoiceAge Corporation. Port` | Known | Audio system |
| 0x002EC604 | `S_CSA_CODECMSG` | Known | Audio system |

---

## 4. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002929A8 | `Audible-softwaren i dette produkt bruges i henhold til en licensaftale fra Audib` | Known | Audible audiobook format |
| 0x00292A08 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x002958BC | `Dies Audible Software in diesem Produkt wird in Lizenz von Audible verwendet. Co` | Known | Audible audiobook format |
| 0x00295916 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x00298258 | `El software Audible incorporado en este producto se usa bajo licencia de Audible` | Known | Audible audiobook format |
| 0x002982B6 | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x0029ACC8 | `Audible-ohjelmistoa t` | Known | Audible audiobook format |
| 0x0029ACFA | `n Audiblen lisenssill` | Known | Audible audiobook format |
| 0x0029AD1F | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x0029E1F8 | `Le logiciel Audible de ce produit est utilis` | Known | Audible audiobook format |
| 0x0029E226 | ` sous licence Audible. Copyright ` | Known | Audible audiobook format |
| 0x0029E249 | ` 2002 par Audible, Inc. Tous drois r` | Known | Audible audiobook format |
| 0x002A0C40 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x002A0C69 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x002A0C98 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x002A38E1 | ` Audible ` | Known | Audible audiobook format |
| 0x002A3902 | `Audible ` | Known | Audible audiobook format |
| 0x002A394F | ` Copyright 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002A6793 | ` Audible ` | Known | Audible audiobook format |
| 0x002A67AE | ` Audible` | Known | Audible audiobook format |
| 0x002A67F1 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002A97DC | `De Audible-software in dit product wordt gebruikt in licentie van Audible. Copyr` | Known | Audible audiobook format |
| 0x002A9833 | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x002AC0D8 | `Audible-programvaren i dette produktet brukes p` | Known | Audible audiobook format |
| 0x002AC109 | ` lisens fra Audible. Copyright ` | Known | Audible audiobook format |
| 0x002AC12A | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x002AEAD8 | `O software Audible deste produto ` | Known | Audible audiobook format |
| 0x002AEB0D | `a da Audible. Copyright ` | Known | Audible audiobook format |
| 0x002AEB27 | ` 2002 por Audible, Inc. Todos os direitos reservados.` | Known | Audible audiobook format |
| 0x002B14C0 | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x002B14EF | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x002B1506 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x002B3DD7 | ` Audible ` | Known | Audible audiobook format |
| 0x002B3DE9 | ` Audible ` | Known | Audible audiobook format |
| 0x002B3E0D | ` 2002 by Audible, Inc.` | Known | Audible audiobook format |
| 0x002B6900 | `Audible ` | Known | Audible audiobook format |
| 0x002B6914 | ` Audible ` | Known | Audible audiobook format |
| 0x002B6932 | `Copyright (C) 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002C726C | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x002C72C1 | ` 2004 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |

---

## 5. Audio/Codec - AAC

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
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

## 6. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00292BA4 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunhofer IIS og THOMSON multi` | Known | MP3 codec |
| 0x00295ACB | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und THOMSON Multimedia.` | Known | MP3 codec |
| 0x00298455 | `n de audio MPEG Layer-3 utilizada bajo licencia de Fraunhofer IIS y THOMSON mult` | Known | MP3 codec |
| 0x0029AE88 | `MPEG Layer-3 -` | Known | MP3 codec |
| 0x0029AE9A | `nikoodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | MP3 codec |
| 0x0029E3C8 | `encodage audio MPEG Layer-3 sous licence de Fraunhofer IIS et THOMSON multimedia` | Known | MP3 codec |
| 0x002A0DE0 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da Fraunhofer IIS e THOMSON` | Known | MP3 codec |
| 0x002A3B20 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002A3B6C | `Fraunhofer IIS ` | Known | MP3 codec |
| 0x002A697C | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002A69A3 | ` Fraunhofer IIS` | Known | MP3 codec |
| 0x002A99CC | `Technologie voor codering van MPEG Layer-3-audio in licentie van Fraunhofer IIS ` | Known | MP3 codec |
| 0x002AC2B4 | `MPEG Layer-3 lydkodingsteknologi lisensiert fra Fraunhofer IIS og THOMSON multim` | Known | MP3 codec |
| 0x002AECDB | `udio MPEG Layer-3 ` | Known | MP3 codec |
| 0x002AECEF | ` licenciada da Fraunhofer IIS e THOMSON multimedia.` | Known | MP3 codec |
| 0x002B16A0 | `Ljudkodningstekniken MPEG Layer-3 ` | Known | MP3 codec |
| 0x002B16D6 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | MP3 codec |
| 0x002B3F70 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002B3F92 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x002B6AA0 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002B6AC5 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x002C7438 | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |

---

## 7. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000390B0 | `iTunesDB` | Known | iTunes database |
| 0x00295614 | ` von iTunes.` | Known | iTunes database |
| 0x00297FB5 | ` Ayuda iTunes.` | Known | iTunes database |
| 0x0029AA38 | `ytyy iPod-ohjeista iTunesin ohjevalikosta.` | Known | iTunes database |
| 0x0029DF1C | `es, consultez l'Aide iPod dans le menu Aide iTunes.` | Known | iTunes database |
| 0x002A0981 | ` dettagliate, fai clic su "Aiuto iPod" disponibile dal menu Aiuto di iTunes.` | Known | iTunes database |
| 0x002A3497 | `iTunes` | Known | iTunes database |
| 0x002A6497 | `, iTunes ` | Known | iTunes database |
| 0x002ABE45 | `res automatisk til iPod. Hvis du vil vite mer, kan du velge iPod Hjelp fra Hjelp` | Known | iTunes database |
| 0x002AE80D | ` no menu Ajuda do iTunes.` | Known | iTunes database |
| 0x002B124C | `lp i iTunes Hj` | Known | iTunes database |
| 0x002B3B66 | ` iTunes ` | Known | iTunes database |
| 0x002B664E | ` iTunes ` | Known | iTunes database |

---

## 8. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0003533C | `pu8FirewireGuid` | Known | FireWire interface |
| 0x0014B129 | `FirewirePower` | Known | FireWire interface |
| 0x0029263D | `r du har tilsluttet iPod som FireWire-disk, skal du dobbeltklikke p` | Known | FireWire interface |
| 0x0029280D | `r du har tilsluttet iPod som FireWire-disk, skal du dobbeltklikke p` | Known | FireWire interface |
| 0x002954A9 | `nnen die enthaltenen Kontakte im Standard-Format "vCards" exportieren. Nachdem S` | Known | FireWire interface |
| 0x002956B0 | `nnen die enthaltenen Kontakte im Standard-Format "vCards" exportieren. Nachdem S` | Known | FireWire interface |
| 0x00297EC1 | `gido FireWire y hacer doble clic en el icono del iPod en su escritorio. Arrastre` | Known | FireWire interface |
| 0x002980D4 | `gido FireWire y hacer doble clic en el icono del iPod en su escritorio. Arrastre` | Known | FireWire interface |
| 0x0029A93A | ` tietonsa vakiintuneessa "vCards"-muodossa. Kun iPodin FireWire-levytila on sall` | Known | FireWire interface |
| 0x0029AAFE | ` tietonsa vakiintuneessa "vCards"-muodossa. Kun iPodin FireWire-levytila on sall` | Known | FireWire interface |
| 0x0029DE19 | ` votre iPod comme disque dur FireWire, double-cliquez simplement sur l` | Known | FireWire interface |
| 0x0029E041 | ` votre iPod comme disque dur FireWire, double-cliquez simplement sur l` | Known | FireWire interface |
| 0x002A07CB | ` contenere, insieme alla musica, fino a un migliaio di contatti. Rubrica Indiriz` | Known | FireWire interface |
| 0x002A09DB | ` contenere, insieme alla musica, fino a un migliaio di contatti. Microsoft Outlo` | Known | FireWire interface |
| 0x002A330A | ` FireWire ` | Known | FireWire interface |
| 0x002A3638 | ` FireWire ` | Known | FireWire interface |
| 0x002A638C | ` FireWire ` | Known | FireWire interface |
| 0x002A65B4 | ` FireWire ` | Known | FireWire interface |
| 0x002A9314 | `De iPod biedt ruimte voor maar liefst duizend adressen die u samen met uw muziek` | Known | FireWire interface |
| 0x002A9544 | `De iPod biedt ruimte voor maar liefst duizend adressen die u samen met uw muziek` | Known | FireWire interface |
| 0x002ABD97 | `rst aktiverer du iPod som en FireWire-disk, deretter dobbeltklikker du p` | Known | FireWire interface |
| 0x002ABF5F | `rst aktiverer du iPod som en FireWire-disk, deretter dobbeltklikker du p` | Known | FireWire interface |
| 0x002AE708 | `gido FireWire, basta clicar duas vezes no ` | Known | FireWire interface |
| 0x002AE92F | `gido FireWire, basta clicar duas vezes no ` | Known | FireWire interface |
| 0x002B115A | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x002B1334 | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x002B3AA9 | ` FireWire ` | Known | FireWire interface |
| 0x002B3C6E | ` FireWire ` | Known | FireWire interface |
| 0x002B657E | ` FireWire ` | Known | FireWire interface |
| 0x002B677A | ` FireWire ` | Known | FireWire interface |
| 0x002C6E6B | `X's Address Book, Microsoft Entourage, and Palm Desktop can all export their con` | Known | FireWire interface |
| 0x002C6FFA | `Your iPod can store up to one thousand contacts right alongside your music. Micr` | Known | FireWire interface |
| 0x00325D34 | `Running on FireWire power.` | Known | FireWire interface |
| 0x00325F57 | `Turn on/off firewire` | Known | FireWire interface |

---

## 9. Hardware (GPIO)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002EC516 | `S_GPIO_IR_TMOUT` | Known | GPIO hardware |
| 0x002EC56B | `S_GPIO` | Known | GPIO hardware |
| 0x002EC57C | `S_GPIO_CD_A` | Known | GPIO hardware |
| 0x002EC58D | `S_GPIO_CD_B` | Known | GPIO hardware |
| 0x002EC59E | `S_GPIO_ACK` | Known | GPIO hardware |

---

## 10. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00035140 | `diskModeImageRev` | Known | Hardware interface |
| 0x0014956A | `SPINDLE SERVO FAILURE` | Known | Hardware interface |
| 0x002EC5AF | `S_LCD_TIMER` | Known | Hardware interface |
| 0x002EC5C0 | `S_I2C_SCAN` | Known | Hardware interface |

---

## 11. Storage (ATA/Disk)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x003271D8 | `!ATADisk is FDISK format.` | Known | ATA/disk interface |

---

## 12. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0014B11C | `BatteryLevel` | Known | Power management |
| 0x002C7D2C | `Low Battery` | Known | Power management |
| 0x002C7D40 | `Charging` | Known | Power management |
| 0x002C7D4C | `Charging` | Known | Power management |
| 0x00325D70 | `Power status is invalid.` | Known | Power management |
| 0x00325E1F | `@Battery reading is %d (%x) E%4cF` | Known | Power management |

---

## 13. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001489A3 | `\Calendars` | Known | UI element |
| 0x001489AE | `Calendars` | Known | UI element |
| 0x002931F8 | `Alarmer` | Known | UI element |
| 0x00295882 | `hlen", um den Alarm zu beenden.` | Known | UI element |
| 0x00297684 | `Calendario` | Known | UI element |
| 0x00297690 | `Calendarios` | Known | UI element |
| 0x002981D4 | `Calendario` | Known | UI element |
| 0x002981E0 | `Calendarios` | Known | UI element |
| 0x0029823C | `Alarma` | Known | UI element |
| 0x00298A6E | `Calendario` | Known | UI element |
| 0x00298B90 | `Alarmas` | Known | UI element |
| 0x00298CCC | `Calendarios` | Known | UI element |
| 0x0029E1D8 | `Alarme` | Known | UI element |
| 0x0029EAA8 | `Alarmes` | Known | UI element |
| 0x002A0040 | `Calendario` | Known | UI element |
| 0x002A004C | `Calendari` | Known | UI element |
| 0x002A0BC0 | `Calendario` | Known | UI element |
| 0x002A0BCC | `Calendari` | Known | UI element |
| 0x002A1306 | `Calendario` | Known | UI element |
| 0x002A152C | `Calendari` | Known | UI element |
| 0x002AC948 | `Alarmer` | Known | UI element |
| 0x002AEAB8 | `Alarme` | Known | UI element |
| 0x002AF43C | `Alarmes` | Known | UI element |
| 0x002C6698 | `Calendar` | Known | UI element |
| 0x002C66A4 | `Calendars` | Known | UI element |
| 0x002C71C4 | `Calendar` | Known | UI element |
| 0x002C71D0 | `Calendars` | Known | UI element |
| 0x002C799A | `Calendar` | Known | UI element |
| 0x002C7AB4 | `Alarms` | Known | UI element |
| 0x002C7BE4 | `Calendars` | Known | UI element |

---

## 14. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00294CAC | `Extras` | Known | Menu item |
| 0x00296030 | `Extras` | Known | Menu item |
| 0x002963C0 | `Extras` | Known | Menu item |
| 0x002976A8 | `Extras` | Known | Menu item |
| 0x00298A90 | `Extras` | Known | Menu item |
| 0x00298E20 | `Extras` | Known | Menu item |
| 0x0029D504 | `Albums` | Known | Menu item |
| 0x0029D50C | `Genres` | Known | Menu item |
| 0x0029D56C | `Extras` | Known | Menu item |
| 0x0029D5C0 | `Albums` | Known | Menu item |
| 0x0029D720 | `Albums` | Known | Menu item |
| 0x0029D744 | `Genres` | Known | Menu item |
| 0x0029E99C | `Extras` | Known | Menu item |
| 0x0029E9C6 | `Genres` | Known | Menu item |
| 0x0029E9DE | `Albums` | Known | Menu item |
| 0x0029EB94 | `Genres` | Known | Menu item |
| 0x0029EBA8 | `Albums` | Known | Menu item |
| 0x0029EC84 | `Genres` | Known | Menu item |
| 0x0029EC98 | `Albums` | Known | Menu item |
| 0x0029ED6C | `Extras` | Known | Menu item |
| 0x002A8B40 | `Albums` | Known | Menu item |
| 0x002A8B48 | `Genres` | Known | Menu item |
| 0x002A8BF0 | `Albums` | Known | Menu item |
| 0x002A9FAA | `Genres` | Known | Menu item |
| 0x002A9FC2 | `Albums` | Known | Menu item |
| 0x002AA15C | `Genres` | Known | Menu item |
| 0x002AA16C | `Albums` | Known | Menu item |
| 0x002AA250 | `Genres` | Known | Menu item |
| 0x002AA264 | `Albums` | Known | Menu item |
| 0x002ADF08 | `Extras` | Known | Menu item |
| 0x002ADF20 | `Audiobooks` | Known | Menu item |
| 0x002AF34C | `Extras` | Known | Menu item |
| 0x002AF356 | `Audiobooks` | Known | Menu item |
| 0x002AF508 | `Audiobooks` | Known | Menu item |
| 0x002AF64C | `Audiobooks` | Known | Menu item |
| 0x002AF6D4 | `Extras` | Known | Menu item |
| 0x002C663C | `Now Playing` | Known | Menu item |
| 0x002C6648 | `Artists` | Known | Menu item |
| 0x002C6660 | `Albums` | Known | Menu item |
| 0x002C6668 | `Genres` | Known | Menu item |
| 0x002C6670 | `Composers` | Known | Menu item |
| 0x002C66C4 | `Extras` | Known | Menu item |
| 0x002C66CC | `Playlists` | Known | Menu item |
| 0x002C66D8 | `Audiobooks` | Known | Menu item |
| 0x002C66E4 | `Shuffle Songs` | Known | Menu item |
| 0x002C6708 | `Albums` | Known | Menu item |
| 0x002C6D18 | `Now Playing` | Known | Menu item |
| 0x002C7980 | `Shuffle Songs` | Known | Menu item |
| 0x002C79B8 | `Extras` | Known | Menu item |
| 0x002C79C2 | `Audiobooks` | Known | Menu item |
| 0x002C79D2 | `Composers` | Known | Menu item |
| 0x002C79DE | `Genres` | Known | Menu item |
| 0x002C79F2 | `Albums` | Known | Menu item |
| 0x002C79FE | `Artists` | Known | Menu item |
| 0x002C7A0A | `Playlists` | Known | Menu item |
| 0x002C7B34 | `Main Menu` | Known | Menu item |
| 0x002C7B64 | `Audiobooks` | Known | Menu item |
| 0x002C7B70 | `Composers` | Known | Menu item |
| 0x002C7B7C | `Genres` | Known | Menu item |
| 0x002C7B8C | `Albums` | Known | Menu item |
| 0x002C7B94 | `Artists` | Known | Menu item |
| 0x002C7B9C | `Playlists` | Known | Menu item |
| 0x002C7BA8 | `Settings` | Known | Menu item |
| 0x002C7C7C | `Settings` | Known | Menu item |
| 0x002C7C88 | `Genres` | Known | Menu item |
| 0x002C7C90 | `Artists` | Known | Menu item |
| 0x002C7C98 | `Albums` | Known | Menu item |
| 0x002C7CA0 | `Composers` | Known | Menu item |
| 0x002C7CBC | `Audiobooks` | Known | Menu item |
| 0x002C7CC8 | `Playlists` | Known | Menu item |
| 0x002C7D58 | `Extras` | Known | Menu item |
| 0x002C7D68 | `Main Menu` | Known | Menu item |

---

## 15. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000351A4 | `iPod_Control\Device` | Filesystem Path | |
| 0x000351F4 | `iPod_Control` | Filesystem Path | |
| 0x000352FC | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x00038D64 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x00039098 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000390C4 | `System_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x0006D00C | `iPod_Control\Testing\TestLog.txt` | Filesystem Path | |
| 0x0006D030 | `\iPod_Control\Testing\TestLog.txt` | Filesystem Path | |
| 0x0006D364 | `iPod_Control\Testing` | Filesystem Path | |
| 0x0006D37C | `\iPod_Control\Testing` | Filesystem Path | |
| 0x0006D3A0 | `iPod_Control\Testing\Tests.Lock` | Filesystem Path | |
| 0x0006D3C0 | `\iPod_Control\Testing\Tests.Lock` | Filesystem Path | |
| 0x0007CAB8 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x0007CF18 | `iPod_Control\Device\Limit` | Filesystem Path | |

---

## 16. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
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
| 0x00294D2C | `Acoustic` | EQ Preset | |
| 0x00294D6C | `Electronic` | EQ Preset | |
| 0x00294D80 | `Hip Hop` | EQ Preset | |
| 0x00294D98 | `Loudness` | EQ Preset | |
| 0x002977A0 | `Hip Hop` | EQ Preset | |
| 0x002977B0 | `Latina` | EQ Preset | |
| 0x002977B8 | `Loudness` | EQ Preset | |
| 0x002977C4 | `Lounge` | EQ Preset | |
| 0x0029A29C | `Lounge` | EQ Preset | |
| 0x0029D668 | `Hip Hop` | EQ Preset | |
| 0x0029D678 | `Latino` | EQ Preset | |
| 0x002A015C | `Hip Hop` | EQ Preset | |
| 0x002A016C | `Latina` | EQ Preset | |
| 0x002A0174 | `Loudness` | EQ Preset | |
| 0x002A0180 | `Lounge` | EQ Preset | |
| 0x002A2AC8 | `Acoustic` | EQ Preset | |
| 0x002A2AD4 | `Bass Booster` | EQ Preset | |
| 0x002A2AF4 | `Classical` | EQ Preset | |
| 0x002A2B10 | `Electronic` | EQ Preset | |
| 0x002A2B24 | `Hip Hop` | EQ Preset | |
| 0x002A2B3C | `Loudness` | EQ Preset | |
| 0x002A2B48 | `Lounge` | EQ Preset | |
| 0x002A2B6C | `Small Speakers` | EQ Preset | |
| 0x002A2B7C | `Spoken Word` | EQ Preset | |
| 0x002A2B88 | `Treble Booster` | EQ Preset | |
| 0x002A2BA8 | `Vocal Booster` | EQ Preset | |
| 0x002A5BEC | `Acoustic` | EQ Preset | |
| 0x002A5BF8 | `Bass Booster` | EQ Preset | |
| 0x002A5C18 | `Classical` | EQ Preset | |
| 0x002A5C34 | `Electronic` | EQ Preset | |
| 0x002A5C48 | `Hip Hop` | EQ Preset | |
| 0x002A5C60 | `Loudness` | EQ Preset | |
| 0x002A5C6C | `Lounge` | EQ Preset | |
| 0x002A5C90 | `Small Speakers` | EQ Preset | |
| 0x002A5CA0 | `Spoken Word` | EQ Preset | |
| 0x002A5CAC | `Treble Booster` | EQ Preset | |
| 0x002A5CCC | `Vocal Booster` | EQ Preset | |
| 0x002A8C94 | `Loudness` | EQ Preset | |
| 0x002A8CA0 | `Lounge` | EQ Preset | |
| 0x002AB6C8 | `Hip Hop` | EQ Preset | |
| 0x002AB6D8 | `Latino` | EQ Preset | |
| 0x002AB6E0 | `Loudness` | EQ Preset | |
| 0x002AB6EC | `Lounge` | EQ Preset | |
| 0x002ADFFC | `Hip Hop` | EQ Preset | |
| 0x002AE00C | `Latino` | EQ Preset | |
| 0x002B0A40 | `Acoustic` | EQ Preset | |
| 0x002B0A4C | `Bass Booster` | EQ Preset | |
| 0x002B0A6C | `Classical` | EQ Preset | |
| 0x002B0A88 | `Electronic` | EQ Preset | |
| 0x002B0A9C | `Hip Hop` | EQ Preset | |
| 0x002B0AB4 | `Loudness` | EQ Preset | |
| 0x002B0AC0 | `Lounge` | EQ Preset | |
| 0x002B0AE4 | `Small Speakers` | EQ Preset | |
| 0x002B0AF4 | `Spoken Word` | EQ Preset | |
| 0x002B0B00 | `Treble Booster` | EQ Preset | |
| 0x002B0B20 | `Vocal Booster` | EQ Preset | |
| 0x002B334C | `Acoustic` | EQ Preset | |
| 0x002B3358 | `Bass Booster` | EQ Preset | |
| 0x002B3378 | `Classical` | EQ Preset | |
| 0x002B3394 | `Electronic` | EQ Preset | |
| 0x002B33A8 | `Hip Hop` | EQ Preset | |
| 0x002B33C0 | `Loudness` | EQ Preset | |
| 0x002B33CC | `Lounge` | EQ Preset | |
| 0x002B33F0 | `Small Speakers` | EQ Preset | |
| 0x002B3400 | `Spoken Word` | EQ Preset | |
| 0x002B340C | `Treble Booster` | EQ Preset | |
| 0x002B342C | `Vocal Booster` | EQ Preset | |
| 0x002B5DF8 | `Acoustic` | EQ Preset | |
| 0x002B5E04 | `Bass Booster` | EQ Preset | |
| 0x002B5E24 | `Classical` | EQ Preset | |
| 0x002B5E40 | `Electronic` | EQ Preset | |
| 0x002B5E54 | `Hip Hop` | EQ Preset | |
| 0x002B5E6C | `Loudness` | EQ Preset | |
| 0x002B5E78 | `Lounge` | EQ Preset | |
| 0x002B5E9C | `Small Speakers` | EQ Preset | |
| 0x002B5EAC | `Spoken Word` | EQ Preset | |
| 0x002B5EB8 | `Treble Booster` | EQ Preset | |
| 0x002B5ED8 | `Vocal Booster` | EQ Preset | |
| 0x002C67D0 | `Acoustic` | EQ Preset | |
| 0x002C67DC | `Bass Booster` | EQ Preset | |
| 0x002C67FC | `Classical` | EQ Preset | |
| 0x002C6818 | `Electronic` | EQ Preset | |
| 0x002C682C | `Hip Hop` | EQ Preset | |
| 0x002C6844 | `Loudness` | EQ Preset | |
| 0x002C6850 | `Lounge` | EQ Preset | |
| 0x002C6874 | `Small Speakers` | EQ Preset | |
| 0x002C6884 | `Spoken Word` | EQ Preset | |
| 0x002C6890 | `Treble Booster` | EQ Preset | |
| 0x002C68B0 | `Vocal Booster` | EQ Preset | |

---

## 17. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009FEC4 | `SetError: %s` | Diagnostic | |
| 0x000A01B8 | `Error reading write parameters page from drive.` | Diagnostic | |
| 0x000A01E8 | `Error sending write parameters page to drive.` | Diagnostic | |
| 0x000C6ABC | `SetTAOWriteMode: Error reading write parameters page.` | Diagnostic | |
| 0x000C6B18 | `SetTAOWriteMode: Error sending write parameters page.` | Diagnostic | |
| 0x00148E71 | `Error loading operating system. Setup cannot continue.` | Diagnostic | |
| 0x00149138 | `PSM_Test() Error` | Diagnostic | |
| 0x0014918D | `   Errors: ` | Diagnostic | |
| 0x00150DF4 | `setct-ErrorTBS` | Diagnostic | |
| 0x0031A250 | `SP_ERR_COMPATIBILITY` | Diagnostic | |
| 0x0031A268 | `SP_ERR_MAJOR_VERSION` | Diagnostic | |
| 0x0031A280 | `SP_ERR_COMP_VERSION` | Diagnostic | |
| 0x0031A294 | `SP_ERR_BAD_MODULE_ID` | Diagnostic | |
| 0x0031A2AC | `SP_ERR_BAD_UNIT_NUMBER` | Diagnostic | |
| 0x0031A2C4 | `SP_ERR_BAD_INSTANCE` | Diagnostic | |
| 0x0031A2D8 | `SP_ERR_BAD_HANDLE` | Diagnostic | |
| 0x0031A2EC | `SP_ERR_BAD_INDEX` | Diagnostic | |
| 0x0031A300 | `SP_ERR_BAD_PARAMETER` | Diagnostic | |
| 0x0031A318 | `SP_ERR_NO_INSTANCES` | Diagnostic | |
| 0x0031A32C | `SP_ERR_NO_COMPONENT` | Diagnostic | |
| 0x0031A3A8 | `SP_ERR_NO_RESOURCES` | Diagnostic | |
| 0x0031A3BC | `SP_ERR_INSTANCE_IN_USE` | Diagnostic | |
| 0x0031A3D4 | `SP_ERR_RESOURCE_OWNED` | Diagnostic | |
| 0x0031A3EC | `SP_ERR_RESOURCE_NOT_OWNED` | Diagnostic | |
| 0x0031A408 | `SP_ERR_INCONSISTENT_PARAMS` | Diagnostic | |
| 0x0031A424 | `SP_ERR_NOT_INITIALIZED` | Diagnostic | |
| 0x0031A43C | `SP_ERR_NOT_ENABLED` | Diagnostic | |
| 0x0031A450 | `SP_ERR_NOT_SUPPORTED` | Diagnostic | |
| 0x0031A468 | `SP_ERR_INIT_FAILED` | Diagnostic | |
| 0x0031A47C | `SP_ERR_BUSY` | Diagnostic | |

---

## 18. Assertions

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F60F8 | `%s(%d): OpenSSL internal error, assertion failed: %s` | Assertion | |
| 0x0011D9C4 | `*** assertion failed: %s, file %s, line %d` | Assertion | |
| 0x0031AAE0 | `SP_ERR_ASSERTION` | Assertion | |
| 0x0032C39C | `SP_ERR_ASSERTION` | Assertion | |

---
