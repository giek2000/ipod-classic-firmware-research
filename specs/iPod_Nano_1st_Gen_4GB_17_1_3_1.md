# iPod Nano 1st Generation (4GB) - RetailOS 1.3.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.3.1 |
| **IPSW** | iPod_17.1.3.1.ipsw |
| **Device** | iPod Nano 1st Generation (4GB) (2005, Click Wheel, Thin Form Factor) |
| **UpdaterFamilyID** | 17 |
| **Binary Size** | 22,905,856 bytes (21.84 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 22,905,856 bytes |
| **Total Strings (>=6)** | 54,684 |
| **Function Prologues** | 19,901 (ARM: 11,306, Thumb: 8,595) |
| **SoC** | PortalPlayer PP5021C |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Encrypted** | No |
| **SHA-256** | `462c975ef81b697e248e48c8471049ad6fffd6a651908449a491f88d1962db8c` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000CCB50 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x0016D0EC | `Channel UnitTests` | Hidden | Developer Tool |

---

## 2. Controllers (TSilver/TC Classes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x008B5B03 | `TCK 6l` | Known | UI controller |
| 0x009DC645 | `TCXPU?P E` | Known | UI controller |

---

## 3. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0001F7D4 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x000C7BCC | `FX_RenderTask` | Known | RTOS task thread |
| 0x000D4D08 | `USBDeviceTask` | Known | RTOS task thread |
| 0x000DB284 | `DiskReaderTask` | Known | RTOS task thread |
| 0x000DF3EC | `LcdUpdateTask` | Known | RTOS task thread |
| 0x000E6AAC | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x000E6AC0 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x001B3F04 | `FirewireTask` | Known | RTOS task thread |
| 0x001B3F1C | `OptoTask` | Known | RTOS task thread |
| 0x001B3F2C | `SerialOptoTask` | Known | RTOS task thread |
| 0x001B3F40 | `BacklightTask` | Known | RTOS task thread |
| 0x001B3F54 | `CNATask` | Known | RTOS task thread |
| 0x001B3F74 | `DiskMgrTask` | Known | RTOS task thread |
| 0x001B3F84 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x001B3F98 | `TopPlugTask` | Known | RTOS task thread |
| 0x001B3FA8 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x001B3FBC | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x001B3FD4 | `AccessoryDetectTask` | Known | RTOS task thread |
| 0x001B3FFC | `AlarmTask` | Known | RTOS task thread |
| 0x001B400C | `WatchdogTask` | Known | RTOS task thread |
| 0x001B4084 | `USBAudioTask` | Known | RTOS task thread |
| 0x001EF9E0 | `HostOSTask` | Known | RTOS task thread |
| 0x0035BCD4 | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x004A8184 | `FX_DisplayTask` | Known | RTOS task thread |

---

## 4. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016CEF8 | `Channel Reserved` | Known | Logging channel |
| 0x0016CF0C | `Channel AppBoot` | Known | Logging channel |
| 0x0016CF1C | `Channel BufferedSongReading` | Known | Logging channel |
| 0x0016CF38 | `Channel PrefsWriting` | Known | Logging channel |
| 0x0016CF50 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x0016CF70 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x0016CF88 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x0016CFA4 | `Channel TestLogging` | Known | Logging channel |
| 0x0016CFB8 | `Channel AppFileLoading` | Known | Logging channel |
| 0x0016CFD0 | `Channel VCardReading` | Known | Logging channel |
| 0x0016CFE8 | `Channel LongSongScanning` | Known | Logging channel |
| 0x0016D05C | `Channel VoiceRecording` | Known | Logging channel |
| 0x0016D074 | `Channel PhotoImporting` | Known | Logging channel |
| 0x0016D08C | `Channel Notes` | Known | Logging channel |
| 0x0016D09C | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x0016D0B8 | `Channel DiskMode` | Known | Logging channel |
| 0x0016D0CC | `Channel Firewire` | Known | Logging channel |
| 0x0016D0E0 | `Channel USB` | Known | Logging channel |
| 0x0016D100 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x0016D118 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 5. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000DA354 | `AudioCodecs` | Known | Audio system |
| 0x002006C6 | `.net codec i dette produkt bruges i henhold til en licensaftale fra VoiceAge Cor` | Known | Audio system |
| 0x00208DD5 | `.net Codec in diesem Produkt wird mit der Lizenz der VoiceAge Corporation verwen` | Known | Audio system |
| 0x002137D0 | `.net codec ` | Known | Audio system |
| 0x0022CD84 | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x0023535D | `.net codec a VoiceAge Coporation c` | Known | Audio system |
| 0x0023D2B2 | ` utilizzato su licenza da VoiceAge Corporation. Il codec ACELP` | Known | Audio system |
| 0x0024F4EA | `.net codec` | Known | Audio system |
| 0x00257A04 | `.net-codec in dit product wordt gebruikt in licentie van VoiceAge Corporation. G` | Known | Audio system |
| 0x0026F485 | `a da VoiceAge Corporation. O codec ACELP` | Known | Audio system |
| 0x0028A785 | `.net codec'i VoiceAge Corporation lisans` | Known | Audio system |
| 0x00323489 | `.net codec in this product is used under license from VoiceAge Corporation. Port` | Known | Audio system |
| 0x00519203 | `msCodeCom` | Known | Audio system |

---

## 6. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000DA434 | `Audible` | Known | Audible audiobook format |
| 0x001F8771 | ` Audible v` | Known | Audible audiobook format |
| 0x001F87C3 | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x001F87D9 | ` 2002 Audible, Inc. V` | Known | Audible audiobook format |
| 0x00200574 | `Audible-softwaren i dette produkt bruges i henhold til en licensaftale fra Audib` | Known | Audible audiobook format |
| 0x002005D4 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x00208C8C | `Die Audible Software in diesem Produkt wird in Lizenz von Audible verwendet. Cop` | Known | Audible audiobook format |
| 0x00208CE5 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x002135BF | ` Audible ` | Known | Audible audiobook format |
| 0x0021361C | ` Audible. ` | Known | Audible audiobook format |
| 0x00213652 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x0021C388 | `El software Audible incluido en este producto se usa bajo licencia de Audible. C` | Known | Audible audiobook format |
| 0x0021C3E3 | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x002240C2 | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright ` | Known | Audible audiobook format |
| 0x002240FC | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x0022CC74 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x0022CCBE | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x0022CCD3 | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x0023521E | ` Audible szoftver az Audible licence alatt van haszn` | Known | Audible audiobook format |
| 0x00235268 | ` 2002, Audible, Inc. Minden jog fenntartva.` | Known | Audible audiobook format |
| 0x0023D1E8 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x0023D211 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0023D240 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x002467D1 | ` Audible ` | Known | Audible audiobook format |
| 0x002467F2 | `Audible ` | Known | Audible audiobook format |
| 0x0024684B | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x0024F39B | ` Audible ` | Known | Audible audiobook format |
| 0x0024F3B6 | ` Audible` | Known | Audible audiobook format |
| 0x0024F3FA | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002578BC | `De Audible-software in dit product wordt gebruikt in licentie van Audible. Copyr` | Known | Audible audiobook format |
| 0x00257913 | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x0025F6A4 | `Audible-programvaren i dette produktet brukes under lisens fra Audible. Copyrigh` | Known | Audible audiobook format |
| 0x0025F6F8 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x002674C4 | `Oprogramowanie Audible w tym produkcie jest wykorzystywane na podstawie licencji` | Known | Audible audiobook format |
| 0x00267530 | ` 2002 Audible, Inc. Wszystkie prawa zastrze` | Known | Audible audiobook format |
| 0x0026F374 | `O software Audible ` | Known | Audible audiobook format |
| 0x0026F3AA | `a da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0026F3C4 | ` 2002 da Audible, Inc. Reservados todos os direitos.` | Known | Audible audiobook format |
| 0x00279DDD | ` Audible ` | Known | Audible audiobook format |
| 0x00279E2F | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x00279E45 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002826A0 | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x002826CF | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x002826E6 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x0028A63C | `ndeki Audible yaz` | Known | Audible audiobook format |
| 0x0028A655 | ` Audible lisans` | Known | Audible audiobook format |
| 0x0028A68A | ` 2002, Audible, Inc. T` | Known | Audible audiobook format |
| 0x0029265B | ` Audible ` | Known | Audible audiobook format |
| 0x0029266D | ` Audible ` | Known | Audible audiobook format |
| 0x00292691 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x0029A4F0 | `Audible ` | Known | Audible audiobook format |
| 0x0029A504 | ` Audible ` | Known | Audible audiobook format |
| 0x0029A52E | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x00323350 | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x003233A5 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |

---

## 7. Audio/Codec - Apple Lossless

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000DA408 | `AppleLossless` | Known | Apple Lossless codec |

---

## 8. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001F89B6 | ` zvuku MPEG Layer-3 byla poskytnuta spole` | Known | MP3 codec |
| 0x001F89E1 | `nostmi Fraunhofer IIS a` | Known | MP3 codec |
| 0x00200770 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunhofer IIS og THOMSON multi` | Known | MP3 codec |
| 0x00208E9B | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und THOMSON Multimedia.` | Known | MP3 codec |
| 0x00213917 | ` MPEG Layer-3 ` | Known | MP3 codec |
| 0x00213955 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x0021C581 | `n de audio MPEG Layer-3 utilizada bajo licencia de Fraunhofer IIS y THOMSON mult` | Known | MP3 codec |
| 0x00224268 | `MPEG Layer-3 -` | Known | MP3 codec |
| 0x0022427A | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | MP3 codec |
| 0x0022CE90 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | MP3 codec |
| 0x002353F0 | `Az MPEG Layer-3 hangk` | Known | MP3 codec |
| 0x00235418 | `gia a Fraunhofer IIS ` | Known | MP3 codec |
| 0x0023D388 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da Fraunhofer IIS e THOMSON` | Known | MP3 codec |
| 0x00246A00 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00246A4C | `Fraunhofer IIS ` | Known | MP3 codec |
| 0x0024F584 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x0024F5AB | ` Fraunhofer IIS` | Known | MP3 codec |
| 0x00257AA0 | `Technologie voor codering van MPEG Layer-3-audio in licentie van Fraunhofer IIS ` | Known | MP3 codec |
| 0x0025F874 | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fraunhofer IIS og THOMSON m` | Known | MP3 codec |
| 0x002676D8 | `Technologia kodowania audio MPEG Layer-3 licencjonowana od Fraunhofer IIS oraz T` | Known | MP3 codec |
| 0x0026F566 | `udio MPEG Layer-3 licenciada pela Fraunhofer IIS e THOMSON multimedia.` | Known | MP3 codec |
| 0x0027A064 | `MPEG Layer-3: ` | Known | MP3 codec |
| 0x0027A0BD | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x00282880 | `Ljudkodningstekniken MPEG Layer-3 ` | Known | MP3 codec |
| 0x002828B6 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | MP3 codec |
| 0x0028A814 | `MPEG Layer-3 ses kodlama teknolojisi Fraunhofer IIS ve THOMSON multimedia'dan li` | Known | MP3 codec |
| 0x002927E8 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x0029280A | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x0029A68C | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x0029A6B1 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x0032351C | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |

---

## 9. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000DA32C | `AppleDRMVersion` | Known | DRM system |
| 0x000DA3CC | `AppleDRM` | Known | DRM system |
| 0x002070B8 | `Trainingsinfos` | Known | DRM system |
| 0x01209B4F | `9ODRM7N9D)6` | Known | DRM system |

---

## 10. Photo System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0003A5B4 | `Photos\Photo Database` | Known | Photo system |
| 0x0016ECDC | `Photo Database Size` | Known | Photo system |

---

## 11. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0003A5A8 | `iTunesDB` | Known | iTunes database |
| 0x0011C1A8 | `iTunes Image DB` | Known | iTunes database |
| 0x00164B10 | `iTunes Image DB` | Known | iTunes database |
| 0x001F02B8 | `iTunesDB` | Known | iTunes database |
| 0x001F6765 | ` z iTunes nebo z vCards.` | Known | iTunes database |
| 0x001F679A | `it kontakty z iTunes, vyberte iPod ze seznamu zdroj` | Known | iTunes database |
| 0x001F6A01 | ` z iTunes nebo z vCards.` | Known | iTunes database |
| 0x001F6A36 | `it kontakty z iTunes, vyberte iPod ze seznamu zdroj` | Known | iTunes database |
| 0x001F9587 | `es iTunes.` | Known | iTunes database |
| 0x001F962D | `es iTunes.` | Known | iTunes database |
| 0x001FE614 | `Din iPod kan opbevare og vise adresser importeret fra iTunes eller vCards. Du ov` | Known | iTunes database |
| 0x001FE669 | `rer adresser automatisk fra iTunes ved at v` | Known | iTunes database |
| 0x001FE6A6 | ` kildeoversigten i iTunes. V` | Known | iTunes database |
| 0x001FE80E | ` iPod. Der findes flere oplysninger i iPod - Oversigt over funktioner, iTunes-hj` | Known | iTunes database |
| 0x001FE88C | `Din iPod kan opbevare og vise adresser importeret fra iTunes eller vCards. Du ov` | Known | iTunes database |
| 0x001FE8E1 | `rer adresser automatisk fra iTunes ved at v` | Known | iTunes database |
| 0x001FE91E | ` kildeoversigten i iTunes. V` | Known | iTunes database |
| 0x001FEA9B | ` iPod. Der findes flere oplysninger i iPod - Oversigt over funktioner, iTunes-hj` | Known | iTunes database |
| 0x002011B2 | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x00201241 | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x002068FB | `nnen Kontakte (mit iTunes importiert oder vCards) auf Ihrem iPod sichern und anz` | Known | iTunes database |
| 0x00206BDF | `nnen Kontakte (mit iTunes importiert oder vCards) auf Ihrem iPod sichern und anz` | Known | iTunes database |
| 0x00209971 | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x00209A12 | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x00210034 | ` iTunes ` | Known | iTunes database |
| 0x0021008C | ` iTunes, ` | Known | iTunes database |
| 0x0021034F | ` iTunes Help ` | Known | iTunes database |
| 0x00210424 | ` iTunes ` | Known | iTunes database |
| 0x0021047C | ` iTunes, ` | Known | iTunes database |
| 0x0021073F | ` iTunes Help ` | Known | iTunes database |
| 0x00214D9A | ` iTunes ` | Known | iTunes database |
| 0x00214EC3 | ` iTunes ` | Known | iTunes database |
| 0x0021A2D8 | `El iPod puede almacenar y mostrar contactos importados desde iTunes o tarjetas v` | Known | iTunes database |
| 0x0021A363 | `ticamente desde iTunes, seleccione el iPod en la lista Fuente. A continuaci` | Known | iTunes database |
| 0x0021A5EC | `El iPod puede almacenar y mostrar contactos importados desde iTunes o tarjetas v` | Known | iTunes database |
| 0x0021A677 | `ticamente desde iTunes, seleccione el iPod en la lista Fuente. A continuaci` | Known | iTunes database |
| 0x0021D078 | `Esta foto es demasiado grande para mostrarla en el iPod. Transfiera las fotos im` | Known | iTunes database |
| 0x0021D11C | `Este formato de foto no puede visualizarse en el iPod. Transfiera las fotos impo` | Known | iTunes database |
| 0x00222298 | ` yhteystietoja, jotka on tuotu iTunesista tai vCardeina. Jos haluat tallentaa yh` | Known | iTunes database |
| 0x00222510 | ` yhteystietoja, jotka on tuotu iTunesista tai vCardeina. Jos haluat tallentaa yh` | Known | iTunes database |
| 0x00224CE6 | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x00224D6D | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x0022AA85 | `iTunes ou de vCards. Pour stocker des contacts provenant d` | Known | iTunes database |
| 0x0022AAC2 | `iTunes automatiquement, s` | Known | iTunes database |
| 0x0022AD37 | `Aide iTunes ou rendez-vous ` | Known | iTunes database |
| 0x0022ADC1 | `iTunes ou de vCards. Pour stocker des contacts provenant d` | Known | iTunes database |
| 0x0022ADFE | `iTunes automatiquement, s` | Known | iTunes database |
| 0x0022B077 | `Aide iTunes ou rendez-vous ` | Known | iTunes database |
| 0x0022DA6E | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x0022DB2B | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x00232D90 | `Az iPod, az iTunes programb` | Known | iTunes database |
| 0x00232E11 | `. Az iTunes programb` | Known | iTunes database |
| 0x00233094 | `t, az iTunes seg` | Known | iTunes database |
| 0x002330F0 | `Az iPod, az iTunes programb` | Known | iTunes database |
| 0x00233171 | `. Az iTunes programb` | Known | iTunes database |
| 0x002333F4 | `t, az iTunes seg` | Known | iTunes database |
| 0x00234F6B | `phez, hogy az iTunes feloldja a z` | Known | iTunes database |
| 0x00236057 | `lja azokat az iTunes programmal.` | Known | iTunes database |
| 0x0023612E | `lja azokat az iTunes programmal.` | Known | iTunes database |
| 0x0023B1F8 | ` memorizzare e visualizzare contatti importati da iTunes o vCards. Per memorizza` | Known | iTunes database |
| 0x0023B4B4 | ` memorizzare e visualizzare contatti importati da iTunes o vCards. Per memorizza` | Known | iTunes database |
| 0x0023DDF5 | ` troppo grande per essere visualizzato in iPod. Trasferisci le foto al tuo compu` | Known | iTunes database |
| 0x0023DEA3 | ` essere visualizzato in iPod. Trasferisci le foto al tuo computer e sincronizzal` | Known | iTunes database |
| 0x00243B4B | `iTunes ` | Known | iTunes database |
| 0x00243BB0 | `iTunes ` | Known | iTunes database |
| 0x00243ECD | `iTunes ` | Known | iTunes database |
| 0x00243F47 | `iTunes ` | Known | iTunes database |
| 0x00243FAC | `iTunes ` | Known | iTunes database |
| 0x002442C7 | `iTunes ` | Known | iTunes database |
| 0x002477C7 | `iTunes ` | Known | iTunes database |
| 0x0024788D | `iTunes ` | Known | iTunes database |
| 0x0024D05F | ` iTunes ` | Known | iTunes database |
| 0x0024D0BA | `. iTunes` | Known | iTunes database |
| 0x0024D2FF | `, iTunes ` | Known | iTunes database |
| 0x0024D363 | ` iTunes ` | Known | iTunes database |
| 0x0024D3BE | `. iTunes` | Known | iTunes database |
| 0x0024D601 | `, iTunes ` | Known | iTunes database |
| 0x0025010E | ` iTunes` | Known | iTunes database |
| 0x002501B3 | ` iTunes` | Known | iTunes database |
| 0x0025857F | `mporteerde foto's naar de computer en synchroniseer ze met iTunes voor weergave ` | Known | iTunes database |
| 0x00258616 | `mporteerde foto's naar de computer en synchroniseer ze met iTunes voor weergave ` | Known | iTunes database |
| 0x0025D6AC | `iPod kan oppbevare og vise kontaktinformasjon som importeres fra iTunes eller vC` | Known | iTunes database |
| 0x0025D719 | `re kontaktinformasjon automatisk med iTunes, markerer du iPod i kildetreet. Klik` | Known | iTunes database |
| 0x0025D8F3 | `r du opp i iPod Funksjonsoversikt eller iTunes Hjelp eller g` | Known | iTunes database |
| 0x0025D958 | `iPod kan oppbevare og vise kontaktinformasjon som importeres fra iTunes eller vC` | Known | iTunes database |
| 0x0025D9C5 | `re kontaktinformasjon automatisk med iTunes, markerer du iPod i kildetreet. Klik` | Known | iTunes database |
| 0x0025DB9D | `r du opp i iPod Funksjonsoversikt eller iTunes Hjelp eller g` | Known | iTunes database |
| 0x00260280 | `r importerte bilder til datamaskinen, og synkroniser via iTunes for ` | Known | iTunes database |
| 0x00260310 | `r importerte bilder til datamaskinen, og synkroniser via iTunes for ` | Known | iTunes database |
| 0x002654BA | ` kontakty zaimportowane z iTunes lub vCards.` | Known | iTunes database |
| 0x002654F6 | ` kontakty z iTunes automatycznie, zaznacz iPod na li` | Known | iTunes database |
| 0x0026577D | ` kontakty importowane z iTunes lub vCards. By przechowywa` | Known | iTunes database |
| 0x002657B8 | ` kontakty z iTunes automatycznie, zaznacz na li` | Known | iTunes database |
| 0x00267214 | `cz iPod do komputera, a iTunes go odblokuje.` | Known | iTunes database |
| 0x0026821C | `cia do komputera i zsynchronizuj je poprzez iTunes, by wy` | Known | iTunes database |
| 0x002682C9 | `cia do komputera i zsynchronizuj je poprzez iTunes, by wy` | Known | iTunes database |
| 0x0026D470 | `O iPod pode armazenar e apresentar contactos importados do iTunes ou vCards. Par` | Known | iTunes database |
| 0x0026D6F0 | `O iPod pode armazenar e apresentar contactos importados do iTunes ou vCards. Par` | Known | iTunes database |
| 0x0026F115 | `o, ligue o iPod ao computador e o iTunes desbloqueia-o.` | Known | iTunes database |
| 0x00270112 | `s do iTunes para as visualizar no iPod.` | Known | iTunes database |
| | *...and 43 more* | | |

---

## 12. Nike+ iPod

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0013AAD8 | `CIapLingoSports::SetSportsMode` | Known | Nike+ fitness |
| 0x0013AC1C | `CIapLingoSports::SetFilterList` | Known | Nike+ fitness |
| 0x0013AC9C | `CIapLingoSports::SetRxdWindowTiming` | Known | Nike+ fitness |
| 0x0013B10C | `GetSportsMode` | Known | Nike+ fitness |
| 0x0013B7A8 | `CIapLingoSports::UnregisterHandler` | Known | Nike+ fitness |
| 0x0013BB34 | `CIapLingoSports::RegisterHandler` | Known | Nike+ fitness |
| 0x0013BB80 | `CIapLingoSports::GetSportsCaps` | Known | Nike+ fitness |
| 0x001F5538 | `Nike+iPod` | Known | Nike+ fitness |
| 0x001F83A1 | `te technologii Nike+iPod p` | Known | Nike+ fitness |
| 0x001F9146 | `Nike+iPod` | Known | Nike+ fitness |
| 0x001FD4B8 | `Nike+iPod` | Known | Nike+ fitness |
| 0x002001F2 | `lpe Nike+iPod med at tilpasse sig din tr` | Known | Nike+ fitness |
| 0x00200E46 | `Nike+iPod` | Known | Nike+ fitness |
| 0x00205738 | `Nike+iPod` | Known | Nike+ fitness |
| 0x002088B0 | `Laufen oder gehen Sie eine bestimmte Strecke mit normalem Tempo. So kann Nike+iP` | Known | Nike+ fitness |
| 0x00209586 | `Nike+iPod` | Known | Nike+ fitness |
| 0x0020E16C | `Nike+iPod` | Known | Nike+ fitness |
| 0x00212F6D | ` Nike+iPod ` | Known | Nike+ fitness |
| 0x002145E6 | `Nike+iPod` | Known | Nike+ fitness |
| 0x002190F8 | `Nike+iPod` | Known | Nike+ fitness |
| 0x0021BFBC | `Correr o caminar una determinada distancia a un ritmo natural, permite a Nike+iP` | Known | Nike+ fitness |
| 0x0021CCF6 | `Nike+iPod` | Known | Nike+ fitness |
| 0x00221094 | `Nike+iPod` | Known | Nike+ fitness |
| 0x00223D02 | ` tunnetun matkan voit auttaa Nike+iPodia tottumaan liikkumistapaasi ja tuottamaa` | Known | Nike+ fitness |
| 0x0022493A | `Nike+iPod` | Known | Nike+ fitness |
| 0x002297C4 | `Nike+iPod` | Known | Nike+ fitness |
| 0x0022C884 | ` un rythme naturel, vous aidez Nike+iPod ` | Known | Nike+ fitness |
| 0x0022D606 | `Nike+iPod` | Known | Nike+ fitness |
| 0x00231B58 | `Nike+iPod` | Known | Nike+ fitness |
| 0x00234E16 | `thet abban, hogy a Nike+iPod alkalmazkodjon az ` | Known | Nike+ fitness |
| 0x00235B9C | ` Nike+iPod` | Known | Nike+ fitness |
| 0x0023A09C | `Nike+iPod` | Known | Nike+ fitness |
| 0x0023CDFC | `Se percorri camminando o correndo una distanza nota con una frequenza normale, p` | Known | Nike+ fitness |
| 0x0023DA72 | `Nike+iPod` | Known | Nike+ fitness |
| 0x00242058 | `Nike+iPod` | Known | Nike+ fitness |
| 0x00246314 | `Nike+iPod ` | Known | Nike+ fitness |
| 0x00247292 | `Nike+iPod` | Known | Nike+ fitness |
| 0x0024BA84 | `Nike+iPod` | Known | Nike+ fitness |
| 0x0024EFC8 | ` Nike+iPod` | Known | Nike+ fitness |
| 0x0024FCE2 | `Nike+iPod` | Known | Nike+ fitness |
| 0x00254674 | `Nike+iPod` | Known | Nike+ fitness |
| 0x00257518 | `Door een vaste afstand te rennen of te lopen in een natuurlijk tempo, kan Nike+i` | Known | Nike+ fitness |
| 0x002581EA | `Nike+iPod` | Known | Nike+ fitness |
| 0x0025C528 | `Nike+iPod` | Known | Nike+ fitness |
| 0x0025F31E | ` en kjent distanse med normal hastighet blir det enklere for Nike+iPod ` | Known | Nike+ fitness |
| 0x0025FF16 | `Nike+iPod` | Known | Nike+ fitness |
| 0x00264248 | `Nike+iPod` | Known | Nike+ fitness |
| 0x00267103 | ` zestaw Nike+iPod do twojego stylu treningu i dostarczy` | Known | Nike+ fitness |
| 0x00267E16 | `Nike+iPod` | Known | Nike+ fitness |
| 0x0026C2B4 | `Nike+iPod` | Known | Nike+ fitness |
| 0x0026EFEB | ` ao Nike+iPod adaptar-se ao seu estilo de exerc` | Known | Nike+ fitness |
| 0x0026FCEF | `Nike+iPod` | Known | Nike+ fitness |
| 0x00274814 | `Nike+iPod` | Known | Nike+ fitness |
| 0x002797A6 | ` Nike+iPod ` | Known | Nike+ fitness |
| 0x0027AC93 | `Nike+iPod` | Known | Nike+ fitness |
| 0x0027F658 | `Nike+iPod` | Known | Nike+ fitness |
| 0x0028232D | `lper du Nike+iPod att anpassa sig till ditt tr` | Known | Nike+ fitness |
| 0x00282F7E | `Nike+iPod` | Known | Nike+ fitness |
| 0x00287308 | `Nike+iPod` | Known | Nike+ fitness |
| 0x0028A263 | `yerek Nike+iPod'un antrenman stilinize uyum sa` | Known | Nike+ fitness |
| 0x0028AF6E | `Nike+iPod` | Known | Nike+ fitness |
| 0x0028F2EC | `Nike+iPod` | Known | Nike+ fitness |
| 0x00292307 | ` Nike+iPod ` | Known | Nike+ fitness |
| 0x00292E5A | `Nike+iPod` | Known | Nike+ fitness |
| 0x002971A4 | `Nike+iPod` | Known | Nike+ fitness |
| 0x0029A1A9 | ` Nike+iPod ` | Known | Nike+ fitness |
| 0x0029AD3E | `Nike+iPod` | Known | Nike+ fitness |
| 0x00320218 | `Nike+iPod` | Known | Nike+ fitness |
| 0x00322FA8 | `By running or walking a known distance at a natural pace you can help Nike+iPod ` | Known | Nike+ fitness |
| 0x00323BBE | `Nike+iPod` | Known | Nike+ fitness |
| 0x0037A9A4 | `TTrainerCntlr` | Known | Nike+ fitness |
| 0x0037A9B4 | `TTrainerHistoryCntlr` | Known | Nike+ fitness |
| 0x0037A9CC | `TTrainerEndWorkoutCntlr` | Known | Nike+ fitness |
| 0x0037A9E4 | `TTrainerNowRunningCntlr` | Known | Nike+ fitness |
| 0x0037A9FC | `TTrainerPauseCntlr` | Known | Nike+ fitness |
| 0x0037AA10 | `TTrainerCalibrateMenuCntlr` | Known | Nike+ fitness |

---

## 13. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0020528C | `Radio-Region` | Known | FM Radio |
| 0x0031FD24 | `Radio Region` | Known | FM Radio |
| 0x0032397C | `Radio Region` | Known | FM Radio |
| 0x00323E88 | `Radio Settings` | Known | FM Radio |

---

## 14. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000DA894 | `FireWireGUID` | Known | FireWire interface |
| 0x000DA8A4 | `FireWireVersion` | Known | FireWire interface |
| 0x000DAE0C | `FireWire` | Known | FireWire interface |
| 0x001F65B3 | ` FireWire nen` | Known | FireWire interface |
| 0x001F98A8 | `FireWire p` | Known | FireWire interface |
| 0x001FE49C | `FireWire-forbindelser underst` | Known | FireWire interface |
| 0x0020150C | `FireWire tilsluttet` | Known | FireWire interface |
| 0x00206754 | `FireWire wird nicht unterst` | Known | FireWire interface |
| 0x00209CB6 | `ber FireWire verbunden` | Known | FireWire interface |
| 0x0020FD8E | ` FireWire. ` | Known | FireWire interface |
| 0x0021534E | ` FireWire` | Known | FireWire interface |
| 0x0021A173 | `n FireWire. Para hacerlo, utilice el cable USB suministrado.` | Known | FireWire interface |
| 0x0021D42C | `FireWire conectado` | Known | FireWire interface |
| 0x002220E4 | `FireWire-tiedonsiirtoa ei tueta. Siirt` | Known | FireWire interface |
| 0x00225040 | `FireWire liitetty` | Known | FireWire interface |
| 0x0022A868 | `Les transferts de morceaux via FireWire ne sont pas pris en charge` | Known | FireWire interface |
| 0x0022DDF8 | `FireWire Connect` | Known | FireWire interface |
| 0x00232C14 | `A FireWire kapcsolat nem t` | Known | FireWire interface |
| 0x002363E8 | `FireWire csatlakozik` | Known | FireWire interface |
| 0x0023B074 | `Le connessioni FireWire non sono supportate. Per trasferire brani, collega il ca` | Known | FireWire interface |
| 0x0023E188 | `FireWire Connesso` | Known | FireWire interface |
| 0x002438F4 | `FireWire ` | Known | FireWire interface |
| 0x00247BD4 | `FireWire ` | Known | FireWire interface |
| 0x0024CEA4 | `FireWire ` | Known | FireWire interface |
| 0x00250454 | `FireWire ` | Known | FireWire interface |
| 0x002556A2 | `ren via FireWire, maar alleen via de meegeleverde USB-kabel.` | Known | FireWire interface |
| 0x00258908 | `FireWire aangesloten` | Known | FireWire interface |
| 0x0025D530 | `Tilkobling via FireWire st` | Known | FireWire interface |
| 0x002605D4 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x0026531B | `czenie FireWire nie jest wspierane. By przes` | Known | FireWire interface |
| 0x002685EF | `czony przez Firewire` | Known | FireWire interface |
| 0x0026D30F | `es FireWire n` | Known | FireWire interface |
| 0x00270498 | `FireWire ligado` | Known | FireWire interface |
| 0x00276485 | ` FireWire ` | Known | FireWire interface |
| 0x0027B95B | ` FireWire` | Known | FireWire interface |
| 0x00280614 | `FireWire-` | Known | FireWire interface |
| 0x00283640 | `FireWire anslutet` | Known | FireWire interface |
| 0x002882AC | `FireWire ba` | Known | FireWire interface |
| 0x0028B6FC | `FireWire Ba` | Known | FireWire interface |
| 0x002906D1 | ` FireWire ` | Known | FireWire interface |
| 0x00293520 | `FireWire ` | Known | FireWire interface |
| 0x00298521 | ` FireWire ` | Known | FireWire interface |
| 0x0029B43C | `FireWire ` | Known | FireWire interface |
| 0x00321644 | `FireWire connections are not supported. To transfer songs, connect the USB cable` | Known | FireWire interface |
| 0x003242CC | `FireWire Connected` | Known | FireWire interface |

---

## 15. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000673A8 | `LCD Module could not be determined.` | Known | Hardware interface |
| 0x000DAED0 | `ForcedDiskMode` | Known | Hardware interface |
| 0x0016EC3C | `Enter Disk Mode` | Known | Hardware interface |
| 0x0016EC4C | `Exit Disk Mode` | Known | Hardware interface |
| 0x001EFD54 | `I2C write Error` | Known | Hardware interface |
| 0x001EFD68 | `I2C read Error %02x` | Known | Hardware interface |
| 0x00321638 | `Disk Mode` | Known | Hardware interface |
| 0x0051B044 | `OCSP_RESPID` | Known | Hardware interface |
| 0x00D15C71 | `T/RzN7I2C` | Known | Hardware interface |

---

## 16. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000B0AB0 | `PowerManager` | Known | Power management |
| 0x000DADE8 | `PowerInformation` | Known | Power management |
| 0x0016ECA8 | `Begin Charging` | Known | Power management |
| 0x0016ECB8 | `Stop Charging` | Known | Power management |
| 0x001B3F60 | `USBPowerSense` | Known | Power management |
| 0x001B4020 | `PCFPowerMgr` | Known | Power management |
| 0x001F6F1C | `PowerSong` | Known | Power management |
| 0x001F7208 | `PowerSong` | Known | Power management |
| 0x001FEDCC | `PowerSong` | Known | Power management |
| 0x001FF06C | `PowerSong` | Known | Power management |
| 0x002071A4 | `PowerSong` | Known | Power management |
| 0x00207474 | `PowerSong` | Known | Power management |
| 0x00210C0C | `PowerSong` | Known | Power management |
| 0x002110FC | `PowerSong` | Known | Power management |
| 0x0021ABBC | `PowerSong` | Known | Power management |
| 0x0021AE98 | `PowerSong` | Known | Power management |
| 0x00222A3C | `PowerSong` | Known | Power management |
| 0x00222CF8 | `PowerSong` | Known | Power management |
| 0x0022B37C | `PowerSong` | Known | Power management |
| 0x0022B674 | `PowerSong` | Known | Power management |
| 0x00233730 | `PowerSong` | Known | Power management |
| 0x00233A50 | `PowerSong` | Known | Power management |
| 0x0023BA30 | `PowerSong` | Known | Power management |
| 0x0023BD10 | `PowerSong` | Known | Power management |
| 0x00244680 | `PowerSong` | Known | Power management |
| 0x002449F4 | `PowerSong` | Known | Power management |
| 0x0024D920 | `PowerSong` | Known | Power management |
| 0x0024DC00 | `PowerSong` | Known | Power management |
| 0x00256018 | `PowerSong` | Known | Power management |
| 0x002562D4 | `PowerSong` | Known | Power management |
| 0x0025DEA0 | `PowerSong` | Known | Power management |
| 0x0025E138 | `PowerSong` | Known | Power management |
| 0x00265CC0 | `PowerSong` | Known | Power management |
| 0x00265F68 | `PowerSong` | Known | Power management |
| 0x0026DC50 | `PowerSong` | Known | Power management |
| 0x0026DF50 | `PowerSong` | Known | Power management |
| 0x00277474 | `PowerSong` | Known | Power management |
| 0x00277964 | `PowerSong` | Known | Power management |
| 0x00280EF8 | `PowerSong` | Known | Power management |
| 0x002811A0 | `PowerSong` | Known | Power management |
| 0x00288D44 | `PowerSong` | Known | Power management |
| 0x00289020 | `PowerSong` | Known | Power management |
| 0x00290FAC | `PowerSong` | Known | Power management |
| 0x0029126C | `PowerSong` | Known | Power management |
| 0x00298E44 | `PowerSong` | Known | Power management |
| 0x002990F4 | `PowerSong` | Known | Power management |
| 0x003211EC | `Charging` | Known | Power management |
| 0x00321E24 | `Sensor Battery` | Known | Power management |
| 0x00321E7C | `PowerSong` | Known | Power management |
| 0x003220F0 | `PowerSong` | Known | Power management |
| 0x00322DD8 | `Low Battery` | Known | Power management |
| 0x00324318 | `Low Battery` | Known | Power management |

---

## 17. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000B2E24 | `Calendar alarm!` | Known | UI element |
| 0x000B2E38 | `Calendar Not ready` | Known | UI element |
| 0x001FCD8C | `Alarmer` | Known | UI element |
| 0x00201450 | `Alarmer` | Known | UI element |
| 0x002188A8 | `Calendario` | Known | UI element |
| 0x002188B4 | `Calendarios` | Known | UI element |
| 0x002188C0 | `Calendarios` | Known | UI element |
| 0x002188FC | `Alarmas` | Known | UI element |
| 0x00219104 | `Calendario` | Known | UI element |
| 0x00219110 | `Calendarios` | Known | UI element |
| 0x0021BDBC | `Alarma` | Known | UI element |
| 0x0021C9AC | `Alarma` | Known | UI element |
| 0x0021CA1C | `Alarma` | Known | UI element |
| 0x0021CCA6 | `Calendario` | Known | UI element |
| 0x0021CF48 | `Alarma` | Known | UI element |
| 0x0021D324 | `Alarma` | Known | UI element |
| 0x0021D374 | `Alarmas` | Known | UI element |
| 0x00228FA4 | `Alarmes` | Known | UI element |
| 0x0022C600 | `Alarme` | Known | UI element |
| 0x0022D2A0 | `Alarme` | Known | UI element |
| 0x0022D308 | `Alarme` | Known | UI element |
| 0x0022D878 | `Alarme` | Known | UI element |
| 0x0022DCA8 | `Alarme` | Known | UI element |
| 0x0022DD20 | `Alarmes` | Known | UI element |
| 0x0023987C | `Calendario` | Known | UI element |
| 0x00239888 | `Calendari` | Known | UI element |
| 0x00239894 | `Calendari` | Known | UI element |
| 0x0023A0A8 | `Calendario` | Known | UI element |
| 0x0023A0B4 | `Calendari` | Known | UI element |
| 0x0023DA22 | `Calendario` | Known | UI element |
| 0x0025BDD0 | `Alarmer` | Known | UI element |
| 0x00260118 | `Alarmtidspunkt` | Known | UI element |
| 0x00260510 | `Alarmer` | Known | UI element |
| 0x00263AC0 | `Alarmy` | Known | UI element |
| 0x00263DA8 | `Gotowe` | Known | UI element |
| 0x00265E00 | `Gotowe?` | Known | UI element |
| 0x00268518 | `Alarmy` | Known | UI element |
| 0x0026BAF8 | `Alarmes` | Known | UI element |
| 0x0026EDE0 | `Alarme` | Known | UI element |
| 0x0026F974 | `Alarme` | Known | UI element |
| 0x002703D8 | `Alarmes` | Known | UI element |
| 0x0028319C | `Alarmtid` | Known | UI element |
| 0x00286B0C | `Alarmlar` | Known | UI element |
| 0x0028B1AC | `Alarm Zaman` | Known | UI element |
| 0x0028B640 | `Alarmlar` | Known | UI element |
| 0x0031F9D0 | `Calendar` | Known | UI element |
| 0x0031F9DC | `Calendars` | Known | UI element |
| 0x0031F9E8 | `Calendars` | Known | UI element |
| 0x0031FA1C | `Alarms` | Known | UI element |
| 0x00320224 | `Calendar` | Known | UI element |
| 0x00320230 | `Calendars` | Known | UI element |
| 0x0032393C | `Alarm Clock` | Known | UI element |
| 0x00323B7A | `Calendar` | Known | UI element |
| 0x00323E00 | `Alarm Time` | Known | UI element |
| 0x00323E0C | `Alarm Clock` | Known | UI element |
| 0x003241B4 | `Alarm Clock` | Known | UI element |
| 0x00324220 | `Alarms` | Known | UI element |
| 0x00324420 | `GotoBackToIdleCommand` | Known | UI element |
| 0x0035E89C | `TCalendarCntlr_Alarm` | Known | UI element |
| 0x005166E8 | `Calendars/` | Known | UI element |
| 0x00516703 | `Calendars` | Known | UI element |

---

## 18. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001F5494 | `Podcasts` | Known | Menu item |
| 0x001F55CC | `Podcasts` | Known | Menu item |
| 0x001F91BE | `Podcasts` | Known | Menu item |
| 0x001F9820 | `Podcasts` | Known | Menu item |
| 0x001FD424 | `Podcasts` | Known | Menu item |
| 0x001FD554 | `Podcasts` | Known | Menu item |
| 0x00200EB9 | `Podcasts` | Known | Menu item |
| 0x0020147C | `Podcasts` | Known | Menu item |
| 0x00205690 | `Podcasts` | Known | Menu item |
| 0x0020577C | `Extras` | Known | Menu item |
| 0x002057DC | `Podcasts` | Known | Menu item |
| 0x00209598 | `Extras` | Known | Menu item |
| 0x002095EB | `Podcasts` | Known | Menu item |
| 0x00209C20 | `Podcasts` | Known | Menu item |
| 0x00209D1C | `Extras` | Known | Menu item |
| 0x0020E038 | `Podcasts` | Known | Menu item |
| 0x0020E284 | `Podcasts` | Known | Menu item |
| 0x002146A9 | `Podcasts` | Known | Menu item |
| 0x00215228 | `Podcasts` | Known | Menu item |
| 0x00219044 | `Podcasts` | Known | Menu item |
| 0x00219130 | `Extras` | Known | Menu item |
| 0x00219194 | `Podcasts` | Known | Menu item |
| 0x0021CD0C | `Extras` | Known | Menu item |
| 0x0021CD62 | `Podcasts` | Known | Menu item |
| 0x0021D39C | `Podcasts` | Known | Menu item |
| 0x0021D490 | `Extras` | Known | Menu item |
| 0x00229710 | `Podcasts` | Known | Menu item |
| 0x00229758 | `Albums` | Known | Menu item |
| 0x00229770 | `Genres` | Known | Menu item |
| 0x002297B4 | `Photos` | Known | Menu item |
| 0x0022980C | `Extras` | Known | Menu item |
| 0x0022986C | `Podcasts` | Known | Menu item |
| 0x00229924 | `Albums` | Known | Menu item |
| 0x0022D054 | `Photos` | Known | Menu item |
| 0x0022D620 | `Extras` | Known | Menu item |
| 0x0022D628 | `Photos` | Known | Menu item |
| 0x0022D656 | `Genres` | Known | Menu item |
| 0x0022D67A | `Podcasts` | Known | Menu item |
| 0x0022D6AE | `Albums` | Known | Menu item |
| 0x0022D81C | `Genres` | Known | Menu item |
| 0x0022D830 | `Albums` | Known | Menu item |
| 0x0022DB54 | `Photos` | Known | Menu item |
| 0x0022DD34 | `Genres` | Known | Menu item |
| 0x0022DD48 | `Podcasts` | Known | Menu item |
| 0x0022DD64 | `Albums` | Known | Menu item |
| 0x0022DE80 | `Extras` | Known | Menu item |
| 0x00231AA4 | `Podcasts` | Known | Menu item |
| 0x00231BF4 | `Podcasts` | Known | Menu item |
| 0x00235C0E | `Podcasts` | Known | Menu item |
| 0x00236348 | `Podcasts` | Known | Menu item |
| 0x002545D8 | `Podcasts` | Known | Menu item |
| 0x0025461C | `Albums` | Known | Menu item |
| 0x00254630 | `Genres` | Known | Menu item |
| 0x00254718 | `Podcasts` | Known | Menu item |
| 0x002547E0 | `Albums` | Known | Menu item |
| 0x0025822B | `Genres` | Known | Menu item |
| 0x0025824B | `Podcasts` | Known | Menu item |
| 0x00258273 | `Albums` | Known | Menu item |
| 0x002583B4 | `Genres` | Known | Menu item |
| 0x002583C4 | `Albums` | Known | Menu item |
| 0x0025885C | `Genres` | Known | Menu item |
| 0x00258870 | `Podcasts` | Known | Menu item |
| 0x00258888 | `Albums` | Known | Menu item |
| 0x00267E8E | `Podcasts` | Known | Menu item |
| 0x0026C204 | `Podcasts` | Known | Menu item |
| 0x0026C300 | `Extras` | Known | Menu item |
| 0x0026C35C | `Podcasts` | Known | Menu item |
| 0x0026FD0C | `Extras` | Known | Menu item |
| 0x0026FD62 | `Podcasts` | Known | Menu item |
| 0x00270410 | `Podcasts` | Known | Menu item |
| 0x002704E4 | `Extras` | Known | Menu item |
| 0x002746F4 | `Podcasts` | Known | Menu item |
| 0x00274918 | `Podcasts` | Known | Menu item |
| 0x0027AD41 | `Podcasts` | Known | Menu item |
| 0x0027B840 | `Podcasts` | Known | Menu item |
| 0x002835B8 | `Podcasts` | Known | Menu item |
| 0x00287250 | `Podcasts` | Known | Menu item |
| 0x0028739C | `Podcasts` | Known | Menu item |
| 0x0028AFEA | `Podcasts` | Known | Menu item |
| 0x0028B66C | `Podcasts` | Known | Menu item |
| 0x0029B3B0 | `Podcasts` | Known | Menu item |
| 0x00320088 | `Podcasts` | Known | Menu item |
| 0x00320198 | `Now Playing` | Known | Menu item |
| 0x003201A4 | `Artists` | Known | Menu item |
| 0x003201BC | `Albums` | Known | Menu item |
| 0x003201D4 | `Genres` | Known | Menu item |
| 0x003201DC | `Composers` | Known | Menu item |
| 0x00320208 | `Photos` | Known | Menu item |
| 0x00320258 | `Extras` | Known | Menu item |
| 0x00320260 | `Playlists` | Known | Menu item |
| 0x0032026C | `Audiobooks` | Known | Menu item |
| 0x003202A4 | `Shuffle Songs` | Known | Menu item |
| 0x003202B4 | `Podcasts` | Known | Menu item |
| 0x00320364 | `Albums` | Known | Menu item |
| 0x00321604 | `Now Playing` | Known | Menu item |
| 0x003216B4 | `Audiobooks` | Known | Menu item |
| 0x00321D8C | `Settings` | Known | Menu item |
| 0x00322058 | `Settings` | Known | Menu item |
| 0x003236F4 | `Photos` | Known | Menu item |
| 0x00323B4C | `Shuffle Songs` | Known | Menu item |
| | *...and 29 more* | | |

---

## 19. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0001EEDC | `iPod_Control` | Filesystem Path | |
| 0x0001EF08 | `iPod_Control\Device` | Filesystem Path | |
| 0x00028CF4 | `iPod_Control\Device` | Filesystem Path | |
| 0x0002A358 | `iPod_Control` | Filesystem Path | |
| 0x0002A948 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x0003A590 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x0003A5D0 | `System_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x0003C884 | `iPod_Control\Music\` | Filesystem Path | |
| 0x0003F3DC | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x0003F55C | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000623B0 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000692E0 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x0006A824 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x0006A920 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x0010268C | `iPod_Control/Accessories` | Filesystem Path | |
| 0x001033B4 | `iPod_Control/Accessories` | Filesystem Path | |
| 0x00129DCC | `iPod_Control\Device\` | Filesystem Path | |
| 0x00129F7C | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012A034 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012A184 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012A274 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012A354 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012A424 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012A5B0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012A660 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012A754 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012A7F8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012A8AC | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012A968 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012AA9C | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012AC10 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012ACC8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012AE04 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012AED0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012AF9C | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012B064 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012B12C | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012B1F4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012B2B4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012B378 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012B428 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012B508 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012B5F4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012B6B4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012B764 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012B864 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012B944 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012BA4C | `iPod_Control\Device\` | Filesystem Path | |
| 0x0012BB38 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0019E6E4 | `iPod_Control/Device` | Filesystem Path | |

---

## 20. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001F5A18 | `Acoustic` | EQ Preset | |
| 0x001F5A24 | `Bass Booster` | EQ Preset | |
| 0x001F5A44 | `Classical` | EQ Preset | |
| 0x001F5A60 | `Electronic` | EQ Preset | |
| 0x001F5A74 | `Hip Hop` | EQ Preset | |
| 0x001F5A8C | `Loudness` | EQ Preset | |
| 0x001F5A98 | `Lounge` | EQ Preset | |
| 0x001F5ABC | `Small Speakers` | EQ Preset | |
| 0x001F5ACC | `Spoken Word` | EQ Preset | |
| 0x001F5AD8 | `Treble Booster` | EQ Preset | |
| 0x001F5AF8 | `Vocal Booster` | EQ Preset | |
| 0x001FD944 | `Acoustic` | EQ Preset | |
| 0x001FD950 | `Bass Booster` | EQ Preset | |
| 0x001FD970 | `Classical` | EQ Preset | |
| 0x001FD98C | `Electronic` | EQ Preset | |
| 0x001FD9A0 | `Hip Hop` | EQ Preset | |
| 0x001FD9B8 | `Loudness` | EQ Preset | |
| 0x001FD9C4 | `Lounge` | EQ Preset | |
| 0x001FD9E8 | `Small Speakers` | EQ Preset | |
| 0x001FD9F8 | `Spoken Word` | EQ Preset | |
| 0x001FDA04 | `Treble Booster` | EQ Preset | |
| 0x001FDA24 | `Vocal Booster` | EQ Preset | |
| 0x00205BF4 | `Acoustic` | EQ Preset | |
| 0x00205C34 | `Electronic` | EQ Preset | |
| 0x00205C48 | `Hip Hop` | EQ Preset | |
| 0x00205C60 | `Loudness` | EQ Preset | |
| 0x0020E988 | `Hip Hop` | EQ Preset | |
| 0x0020E9A0 | `Loudness` | EQ Preset | |
| 0x0020E9AC | `Lounge` | EQ Preset | |
| 0x0021962C | `Hip Hop` | EQ Preset | |
| 0x0021963C | `Latina` | EQ Preset | |
| 0x00219644 | `Loudness` | EQ Preset | |
| 0x00219650 | `Lounge` | EQ Preset | |
| 0x002215E0 | `Lounge` | EQ Preset | |
| 0x00229D00 | `Hip Hop` | EQ Preset | |
| 0x00229D10 | `Latino` | EQ Preset | |
| 0x00229D24 | `Lounge` | EQ Preset | |
| 0x0023A560 | `Hip Hop` | EQ Preset | |
| 0x0023A570 | `Latina` | EQ Preset | |
| 0x0023A578 | `Loudness` | EQ Preset | |
| 0x0023A584 | `Lounge` | EQ Preset | |
| 0x0024277C | `Acoustic` | EQ Preset | |
| 0x00242788 | `Bass Booster` | EQ Preset | |
| 0x002427A8 | `Classical` | EQ Preset | |
| 0x002427C4 | `Electronic` | EQ Preset | |
| 0x002427D8 | `Hip Hop` | EQ Preset | |
| 0x002427F0 | `Loudness` | EQ Preset | |
| 0x002427FC | `Lounge` | EQ Preset | |
| 0x00242820 | `Small Speakers` | EQ Preset | |
| 0x00242830 | `Spoken Word` | EQ Preset | |
| 0x0024283C | `Treble Booster` | EQ Preset | |
| 0x0024285C | `Vocal Booster` | EQ Preset | |
| 0x0024BFC0 | `Acoustic` | EQ Preset | |
| 0x0024BFCC | `Bass Booster` | EQ Preset | |
| 0x0024BFEC | `Classical` | EQ Preset | |
| 0x0024C008 | `Electronic` | EQ Preset | |
| 0x0024C01C | `Hip Hop` | EQ Preset | |
| 0x0024C034 | `Loudness` | EQ Preset | |
| 0x0024C040 | `Lounge` | EQ Preset | |
| 0x0024C064 | `Small Speakers` | EQ Preset | |
| 0x0024C074 | `Spoken Word` | EQ Preset | |
| 0x0024C080 | `Treble Booster` | EQ Preset | |
| 0x0024C0A0 | `Vocal Booster` | EQ Preset | |
| 0x00254B8C | `Loudness` | EQ Preset | |
| 0x00254B98 | `Lounge` | EQ Preset | |
| 0x0025CA30 | `Latino` | EQ Preset | |
| 0x0025CA38 | `Loudness` | EQ Preset | |
| 0x0025CA44 | `Lounge` | EQ Preset | |
| 0x002647CC | `Hip Hop` | EQ Preset | |
| 0x00264800 | `Lounge` | EQ Preset | |
| 0x0026C7D8 | `Hip Hop` | EQ Preset | |
| 0x0026C7E8 | `Latina` | EQ Preset | |
| 0x0026C7F0 | `Loudness` | EQ Preset | |
| 0x0026C7FC | `Lounge` | EQ Preset | |
| 0x0027FACC | `Acoustic` | EQ Preset | |
| 0x0027FAD8 | `Bass Booster` | EQ Preset | |
| 0x0027FAF8 | `Classical` | EQ Preset | |
| 0x0027FB14 | `Electronic` | EQ Preset | |
| 0x0027FB28 | `Hip Hop` | EQ Preset | |
| 0x0027FB40 | `Loudness` | EQ Preset | |
| 0x0027FB4C | `Lounge` | EQ Preset | |
| 0x0027FB70 | `Small Speakers` | EQ Preset | |
| 0x0027FB80 | `Spoken Word` | EQ Preset | |
| 0x0027FB8C | `Treble Booster` | EQ Preset | |
| 0x0027FBAC | `Vocal Booster` | EQ Preset | |
| 0x002877D8 | `Hip Hop` | EQ Preset | |
| 0x002877EC | `Loudness` | EQ Preset | |
| 0x002877F8 | `Lounge` | EQ Preset | |
| 0x0028F920 | `Acoustic` | EQ Preset | |
| 0x0028F92C | `Bass Booster` | EQ Preset | |
| 0x0028F94C | `Classical` | EQ Preset | |
| 0x0028F968 | `Electronic` | EQ Preset | |
| 0x0028F97C | `Hip Hop` | EQ Preset | |
| 0x0028F994 | `Loudness` | EQ Preset | |
| 0x0028F9A0 | `Lounge` | EQ Preset | |
| 0x0028F9C4 | `Small Speakers` | EQ Preset | |
| 0x0028F9D4 | `Spoken Word` | EQ Preset | |
| 0x0028F9E0 | `Treble Booster` | EQ Preset | |
| 0x0028FA00 | `Vocal Booster` | EQ Preset | |
| 0x00297790 | `Acoustic` | EQ Preset | |
| 0x0029779C | `Bass Booster` | EQ Preset | |
| 0x002977BC | `Classical` | EQ Preset | |
| 0x002977D8 | `Electronic` | EQ Preset | |
| 0x002977EC | `Hip Hop` | EQ Preset | |
| 0x00297804 | `Loudness` | EQ Preset | |
| 0x00297810 | `Lounge` | EQ Preset | |
| 0x00297834 | `Small Speakers` | EQ Preset | |
| 0x00297844 | `Spoken Word` | EQ Preset | |
| 0x00297850 | `Treble Booster` | EQ Preset | |
| 0x00297870 | `Vocal Booster` | EQ Preset | |
| 0x00320764 | `Acoustic` | EQ Preset | |
| 0x00320770 | `Bass Booster` | EQ Preset | |
| 0x00320790 | `Classical` | EQ Preset | |
| 0x003207AC | `Electronic` | EQ Preset | |
| 0x003207C0 | `Hip Hop` | EQ Preset | |
| 0x003207D8 | `Loudness` | EQ Preset | |
| 0x003207E4 | `Lounge` | EQ Preset | |
| 0x00320808 | `Small Speakers` | EQ Preset | |
| 0x00320818 | `Spoken Word` | EQ Preset | |
| 0x00320824 | `Treble Booster` | EQ Preset | |
| 0x00320844 | `Vocal Booster` | EQ Preset | |

---

## 21. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00074840 | `Error-SDriver` | Diagnostic | |
| 0x00074850 | `Error-AClient` | Diagnostic | |
| 0x000A2E7C | `%s Error in file %s.` | Diagnostic | |
| 0x00161538 | `Error...no cases match!` | Diagnostic | |
| 0x0023BB8C | `Errore` | Diagnostic | |
| 0x0023E00C | `Errore` | Diagnostic | |
| 0x0023E2A4 | `Errore` | Diagnostic | |
| 0x00517DA6 | `setct-ErrorTBS` | Diagnostic | |
| 0x0051E41F | `Error!` | Diagnostic | |

---
