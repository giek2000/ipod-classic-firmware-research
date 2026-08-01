# iPod 4th Generation (Color/U2) - RetailOS 1.2.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.2.1 |
| **IPSW** | iPod_11.1.2.1.ipsw |
| **Device** | iPod 4th Generation (Color/U2) (2004, Click Wheel, Color Display) |
| **UpdaterFamilyID** | 11 |
| **Binary Size** | 6,514,176 bytes (6.21 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 6,514,176 bytes |
| **Total Strings (>=6)** | 12,965 |
| **Function Prologues** | 14,868 (ARM: 9,777, Thumb: 5,091) |
| **SoC** | PortalPlayer PP5020 |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Encrypted** | No |
| **SHA-256** | `55845b4694263be104e8bfded72f11d1b1d5b9cbeec64f9ffaced80b0bcdc2f5` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001C9E78 | `MP3ExampleTask` | Hidden | Hidden Test |

---

## 2. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
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
| 0x001D0500 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x001D07FF | `5RunTestsTask` | Known | RTOS task thread |
| 0x002953B4 | `USBDeviceTask` | Known | RTOS task thread |
| 0x00295574 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x002A1504 | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x002A151C | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x002C1A28 | `FWInterruptHandlerTask` | Known | RTOS task thread |
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

## 3. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001D3A7A | `.net codec i dette produkt bruges i henhold til en licensaftale fra VoiceAge Cor` | Known | Audio system |
| 0x001D82A9 | `.net Codec in diesem Produkt wird mit der Lizenz der VoiceAge Corporation verwen` | Known | Audio system |
| 0x001E57CC | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x001E9BCE | ` utilizzato su licenza da VoiceAge Corporation. Il codec ACELP` | Known | Audio system |
| 0x001F3162 | `.net codec` | Known | Audio system |
| 0x001F7B2C | `.net-codec in dit product wordt gebruikt in licentie van VoiceAge Corporation. G` | Known | Audio system |
| 0x00272779 | `.net codec in this product is used under license from VoiceAge Corporation. Port` | Known | Audio system |
| 0x004FD314 | `AudioCodecs` | Known | Audio system |

---

## 4. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001D3928 | `Audible-softwaren i dette produkt bruges i henhold til en licensaftale fra Audib` | Known | Audible audiobook format |
| 0x001D3988 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x001D8160 | `Die Audible Software in diesem Produkt wird in Lizenz von Audible verwendet. Cop` | Known | Audible audiobook format |
| 0x001D81B9 | ` 2002 by Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x001DC540 | `El software Audible incluido en este producto se usa bajo licencia de Audible. C` | Known | Audible audiobook format |
| 0x001DC59B | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x001E084E | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright ` | Known | Audible audiobook format |
| 0x001E0888 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x001E56BC | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x001E5706 | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x001E571B | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x001E9B04 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x001E9B2D | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x001E9B5C | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x001EE58D | ` Audible ` | Known | Audible audiobook format |
| 0x001EE5AE | `Audible ` | Known | Audible audiobook format |
| 0x001EE607 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x001F3013 | ` Audible ` | Known | Audible audiobook format |
| 0x001F302E | ` Audible` | Known | Audible audiobook format |
| 0x001F3072 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x001F79E4 | `De Audible-software in dit product wordt gebruikt in licentie van Audible. Copyr` | Known | Audible audiobook format |
| 0x001F7A3B | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x001FBC04 | `Audible-programvaren i dette produktet brukes under lisens fra Audible. Copyrigh` | Known | Audible audiobook format |
| 0x001FBC58 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0020025C | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x0020028B | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x002002A2 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x00204647 | ` Audible ` | Known | Audible audiobook format |
| 0x00204659 | ` Audible ` | Known | Audible audiobook format |
| 0x0020467D | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002089DC | `Audible ` | Known | Audible audiobook format |
| 0x002089F0 | ` Audible ` | Known | Audible audiobook format |
| 0x00208A1A | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x00272640 | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x00272695 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x004FD23C | `Audible` | Known | Audible audiobook format |

---

## 5. Audio/Codec - Apple Lossless

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x004FD2B0 | `AppleLossless` | Known | Apple Lossless codec |

---

## 6. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001D3B24 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunhofer IIS og THOMSON multi` | Known | MP3 codec |
| 0x001D836F | `r MPEG Layer-3 wurde lizensiert von Fraunhofer IIS und THOMSON Multimedia.` | Known | MP3 codec |
| 0x001DC739 | `n de audio MPEG Layer-3 utilizada bajo licencia de Fraunhofer IIS y THOMSON mult` | Known | MP3 codec |
| 0x001E09F4 | `MPEG Layer-3 -` | Known | MP3 codec |
| 0x001E0A06 | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | MP3 codec |
| 0x001E58D8 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | MP3 codec |
| 0x001E9CA4 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da Fraunhofer IIS e THOMSON` | Known | MP3 codec |
| 0x001EE7BC | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x001EE808 | `Fraunhofer IIS ` | Known | MP3 codec |
| 0x001F31FC | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x001F3223 | ` Fraunhofer IIS` | Known | MP3 codec |
| 0x001F7BC8 | `Technologie voor codering van MPEG Layer-3-audio in licentie van Fraunhofer IIS ` | Known | MP3 codec |
| 0x001FBDD4 | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fraunhofer IIS og THOMSON m` | Known | MP3 codec |
| 0x0020043C | `Ljudkodningstekniken MPEG Layer-3 ` | Known | MP3 codec |
| 0x00200472 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | MP3 codec |
| 0x002047D4 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x002047F6 | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x00208B78 | `MPEG Layer-3 ` | Known | MP3 codec |
| 0x00208B9D | ` Fraunhofer IIS ` | Known | MP3 codec |
| 0x0027280C | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |

---

## 7. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x004FD20C | `AppleDRMVersion` | Known | DRM system |
| 0x004FD244 | `AppleDRM` | Known | DRM system |

---

## 8. Photo System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001CA61C | `Photos\Photo Database` | Known | Photo system |
| 0x001CA678 | `Photo Database` | Known | Photo system |
| 0x001D0484 | `Photo Database` | Known | Photo system |
| 0x001D04D8 | `Photo Database` | Known | Photo system |
| 0x001D04E8 | `Photo Import Database` | Known | Photo system |

---

## 9. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001CA579 | `#!#iTunesDB` | Known | iTunes database |
| 0x001D36E8 | `r fotografier til computeren, og synkroniser via iTunes for at vise dem p` | Known | iTunes database |
| 0x001D439E | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x001D442D | `r importerede fotografier til computeren, og synkroniser via iTunes for at vise ` | Known | iTunes database |
| 0x001D7EDC | `Importierte Fotos werden nicht auf dem TV angezeigt. Senden Sie sie erst an den ` | Known | iTunes database |
| 0x001D8CA9 | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x001D8D4A | `ber iTunes zur Anzeige auf dem iPod.` | Known | iTunes database |
| 0x001DC300 | `celas con iTunes para verlas en la TV.` | Known | iTunes database |
| 0x001DD064 | `Esta foto es demasiado grande para mostrarla en el iPod. Transfiera las fotos im` | Known | iTunes database |
| 0x001DD108 | `Este formato de foto no puede visualizarse en el iPod. Transfiera las fotos impo` | Known | iTunes database |
| 0x001E05E8 | ` kuvat tietokoneelle ja synkronoi ne iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x001E12AA | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x001E1331 | ` tuodut kuvat tietokoneelle ja synkronoi iTunesin kautta katsellaksesi niit` | Known | iTunes database |
| 0x001E5429 | `rez-les sur ordinateur et synchronisez-les avec iTunes.` | Known | iTunes database |
| 0x001E62AE | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x001E636B | `iTunes pour les afficher sur l` | Known | iTunes database |
| 0x001E9850 | `Le foto importate non possono visualizzarsi in TV. Trasferisci le foto al comput` | Known | iTunes database |
| 0x001EA53D | ` troppo grande per essere visualizzato in iPod. Trasferisci le foto al tuo compu` | Known | iTunes database |
| 0x001EA5EB | ` essere visualizzato in iPod. Trasferisci le foto al tuo computer e sincronizzal` | Known | iTunes database |
| 0x001EE2A4 | `iTunes ` | Known | iTunes database |
| 0x001EF347 | `iTunes ` | Known | iTunes database |
| 0x001EF40D | `iTunes ` | Known | iTunes database |
| 0x001F2DB8 | ` iTunes` | Known | iTunes database |
| 0x001F3B8A | ` iTunes` | Known | iTunes database |
| 0x001F3C2F | ` iTunes` | Known | iTunes database |
| 0x001F7799 | `mporteerde foto's op tv niet mogelijk. Kopieer foto's naar de computer en synchr` | Known | iTunes database |
| 0x001F84F3 | `mporteerde foto's naar de computer en synchroniseer ze met iTunes voor weergave ` | Known | iTunes database |
| 0x001F858A | `mporteerde foto's naar de computer en synchroniseer ze met iTunes voor weergave ` | Known | iTunes database |
| 0x001FB9C4 | `r bildene til datamaskinen, og synkroniser dem via iTunes for ` | Known | iTunes database |
| 0x001FC62C | `r importerte bilder til datamaskinen, og synkroniser via iTunes for ` | Known | iTunes database |
| 0x001FC6BC | `r importerte bilder til datamaskinen, og synkroniser via iTunes for ` | Known | iTunes database |
| 0x00200015 | `r bilder till din dator och synkronisera via iTunes f` | Known | iTunes database |
| 0x00200D26 | `r importerade bilder till din dator och synkronisera via iTunes f` | Known | iTunes database |
| 0x00200DBB | `r importerade bilder till din dator och synkronisera via iTunes f` | Known | iTunes database |
| 0x0020441D | ` iTunes ` | Known | iTunes database |
| 0x0020509E | ` iTunes ` | Known | iTunes database |
| 0x00205133 | ` iTunes ` | Known | iTunes database |
| 0x002087C9 | ` iTunes ` | Known | iTunes database |
| 0x00209492 | ` iTunes ` | Known | iTunes database |
| 0x00209528 | ` iTunes ` | Known | iTunes database |
| 0x002723C0 | `Imported photos cannot be viewed on TV. Transfer photos to your computer and syn` | Known | iTunes database |
| 0x00273094 | `This photo is too large to display on iPod. Transfer imported photos to your com` | Known | iTunes database |
| 0x00273124 | `This photo format cannot be viewed on iPod. Transfer imported photos to your com` | Known | iTunes database |
| 0x002A124C | `iTunes Image DB` | Known | iTunes database |
| 0x002C1F74 | `iTunes Image DB` | Known | iTunes database |

---

## 10. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001C9A3C | `FirewireInitiator` | Known | FireWire interface |
| 0x001C9A50 | `FirewireHandler` | Known | FireWire interface |
| 0x001C9DB0 | `FirewireGuid` | Known | FireWire interface |
| 0x001D4944 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x001D9246 | `ber FireWire verbunden` | Known | FireWire interface |
| 0x001DD6B0 | `FireWire conectado` | Known | FireWire interface |
| 0x001E1874 | `FireWire liitetty` | Known | FireWire interface |
| 0x001E68BC | `FireWire Connect` | Known | FireWire interface |
| 0x001EAB5C | `FireWire Connesso` | Known | FireWire interface |
| 0x001EFAA4 | `FireWire ` | Known | FireWire interface |
| 0x001F419C | `FireWire ` | Known | FireWire interface |
| 0x001F8A88 | `FireWire aangesloten` | Known | FireWire interface |
| 0x001FCBA0 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x001FF916 | `rst in din iPod som FireWire-h` | Known | FireWire interface |
| 0x002012B0 | `FireWire anslutet` | Known | FireWire interface |
| 0x002055D8 | `FireWire ` | Known | FireWire interface |
| 0x002099DC | `FireWire ` | Known | FireWire interface |
| 0x00273638 | `FireWire Connected` | Known | FireWire interface |
| 0x004FD320 | `FireWireVersion` | Known | FireWire interface |
| 0x004FD460 | `FireWire` | Known | FireWire interface |
| 0x00505898 | `FireWire` | Known | FireWire interface |
| 0x00505BCB | `<key>FireWireGUID</key>` | Known | FireWire interface |

---

## 11. USB

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00505C19 | `USBCompositeDevice1.6` | Known | USB interface |
| 0x00505C71 | `USBCompositeDevice1.6` | Known | USB interface |

---

## 12. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001C9CF0 | `diskModeImageRev` | Known | Hardware interface |
| 0x001CA148 | `I2C write Error` | Known | Hardware interface |
| 0x001CA15C | `I2C read Error %02x` | Known | Hardware interface |
| 0x00271A00 | `Disk Mode` | Known | Hardware interface |

---

## 13. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001C9A60 | `PCFPowerMgr` | Known | Power management |
| 0x001C9B20 | `USBPowerSense` | Known | Power management |
| 0x001CA088 | `PowerManager` | Known | Power management |
| 0x002715B4 | `Charging` | Known | Power management |
| 0x00273684 | `Low Battery` | Known | Power management |
| 0x004FD47C | `PowerInformation` | Known | Power management |

---

## 14. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001D3250 | `Alarmer` | Known | UI element |
| 0x001D4888 | `Alarmer` | Known | UI element |
| 0x001D8133 | `hlen" beendet Alarm` | Known | UI element |
| 0x001DB090 | `Calendario` | Known | UI element |
| 0x001DB09C | `Calendarios` | Known | UI element |
| 0x001DBD9C | `Calendario` | Known | UI element |
| 0x001DBDA8 | `Calendarios` | Known | UI element |
| 0x001DBDC4 | `Alarmas` | Known | UI element |
| 0x001DC524 | `Alarma` | Known | UI element |
| 0x001DCBC0 | `Alarma` | Known | UI element |
| 0x001DCC30 | `Alarma` | Known | UI element |
| 0x001DCD5A | `Calendario` | Known | UI element |
| 0x001DD278 | `Alarma` | Known | UI element |
| 0x001DD4E0 | `Calendarios` | Known | UI element |
| 0x001DD5B0 | `Alarma` | Known | UI element |
| 0x001DD5F8 | `Alarmas` | Known | UI element |
| 0x001E4EF8 | `Alarmes` | Known | UI element |
| 0x001E5698 | `Alarme` | Known | UI element |
| 0x001E5D1C | `Alarme` | Known | UI element |
| 0x001E5D88 | `Alarme` | Known | UI element |
| 0x001E643C | `Alarme` | Known | UI element |
| 0x001E6784 | `Alarme` | Known | UI element |
| 0x001E67E8 | `Alarmes` | Known | UI element |
| 0x001E8720 | `Calendario` | Known | UI element |
| 0x001E872C | `Calendari` | Known | UI element |
| 0x001E9388 | `Calendario` | Known | UI element |
| 0x001E9394 | `Calendari` | Known | UI element |
| 0x001EA22F | `Calendario` | Known | UI element |
| 0x001EA984 | `Calendari` | Known | UI element |
| 0x001FB528 | `Alarmer` | Known | UI element |
| 0x001FC7B8 | `Alarmtidspunkt` | Known | UI element |
| 0x001FCAE0 | `Alarmer` | Known | UI element |
| 0x00200EA0 | `Alarmtid` | Known | UI element |
| 0x00270FC0 | `Calendar` | Known | UI element |
| 0x00270FCC | `Calendars` | Known | UI element |
| 0x00271F80 | `Calendar` | Known | UI element |
| 0x00271F8C | `Calendars` | Known | UI element |
| 0x00271FA4 | `Alarms` | Known | UI element |
| 0x00272C44 | `Alarm Clock` | Known | UI element |
| 0x00272D82 | `Calendar` | Known | UI element |
| 0x00273230 | `Alarm Time` | Known | UI element |
| 0x0027323C | `Alarm Clock` | Known | UI element |
| 0x00273460 | `Calendars` | Known | UI element |
| 0x0027352C | `Alarm Clock` | Known | UI element |
| 0x0027358C | `Alarms` | Known | UI element |
| 0x002A0E83 | `Calendars` | Known | UI element |
| 0x002A0E8D | `Calendars/` | Known | UI element |

---

## 15. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001D268C | `Podcasts` | Known | Menu item |
| 0x001D411D | `Podcasts` | Known | Menu item |
| 0x001D4290 | `Podcasts` | Known | Menu item |
| 0x001D48B4 | `Podcasts` | Known | Menu item |
| 0x001D6CEC | `Extras` | Known | Menu item |
| 0x001D6D44 | `Podcasts` | Known | Menu item |
| 0x001D8954 | `Extras` | Known | Menu item |
| 0x001D89A6 | `Podcasts` | Known | Menu item |
| 0x001D8B2C | `Podcasts` | Known | Menu item |
| 0x001D91B0 | `Podcasts` | Known | Menu item |
| 0x001D92AC | `Extras` | Known | Menu item |
| 0x001DB0BC | `Extras` | Known | Menu item |
| 0x001DB110 | `Podcasts` | Known | Menu item |
| 0x001DCD84 | `Extras` | Known | Menu item |
| 0x001DCDE2 | `Podcasts` | Known | Menu item |
| 0x001DCF90 | `Podcasts` | Known | Menu item |
| 0x001DD620 | `Podcasts` | Known | Menu item |
| 0x001DD714 | `Extras` | Known | Menu item |
| 0x001E40D0 | `Albums` | Known | Menu item |
| 0x001E40E8 | `Genres` | Known | Menu item |
| 0x001E4104 | `Photos` | Known | Menu item |
| 0x001E417C | `Extras` | Known | Menu item |
| 0x001E41D4 | `Podcasts` | Known | Menu item |
| 0x001E421C | `Albums` | Known | Menu item |
| 0x001E5AB8 | `Photos` | Known | Menu item |
| 0x001E5EF8 | `Extras` | Known | Menu item |
| 0x001E5F18 | `Photos` | Known | Menu item |
| 0x001E5F3A | `Genres` | Known | Menu item |
| 0x001E5F5E | `Podcasts` | Known | Menu item |
| 0x001E5F92 | `Albums` | Known | Menu item |
| 0x001E610C | `Genres` | Known | Menu item |
| 0x001E6114 | `Podcasts` | Known | Menu item |
| 0x001E612C | `Albums` | Known | Menu item |
| 0x001E6394 | `Photos` | Known | Menu item |
| 0x001E67FC | `Genres` | Known | Menu item |
| 0x001E6810 | `Podcasts` | Known | Menu item |
| 0x001E682C | `Albums` | Known | Menu item |
| 0x001E6944 | `Extras` | Known | Menu item |
| 0x001F64F4 | `Albums` | Known | Menu item |
| 0x001F6508 | `Genres` | Known | Menu item |
| 0x001F65FC | `Podcasts` | Known | Menu item |
| 0x001F6640 | `Albums` | Known | Menu item |
| 0x001F8206 | `Genres` | Known | Menu item |
| 0x001F8222 | `Podcasts` | Known | Menu item |
| 0x001F824A | `Albums` | Known | Menu item |
| 0x001F83B4 | `Genres` | Known | Menu item |
| 0x001F83BC | `Podcasts` | Known | Menu item |
| 0x001F83D0 | `Albums` | Known | Menu item |
| 0x001F89E8 | `Genres` | Known | Menu item |
| 0x001F89FC | `Podcasts` | Known | Menu item |
| 0x001F8A14 | `Albums` | Known | Menu item |
| 0x00270F30 | `Now Playing` | Known | Menu item |
| 0x00270F3C | `Artists` | Known | Menu item |
| 0x00270F54 | `Albums` | Known | Menu item |
| 0x00270F6C | `Genres` | Known | Menu item |
| 0x00270F74 | `Composers` | Known | Menu item |
| 0x00270F88 | `Photos` | Known | Menu item |
| 0x00270FF4 | `Extras` | Known | Menu item |
| 0x00270FFC | `Playlists` | Known | Menu item |
| 0x00271008 | `Audiobooks` | Known | Menu item |
| 0x00271038 | `Shuffle Songs` | Known | Menu item |
| 0x00271048 | `Podcasts` | Known | Menu item |
| 0x00271088 | `Albums` | Known | Menu item |
| 0x002719CC | `Now Playing` | Known | Menu item |
| 0x00272A00 | `Photos` | Known | Menu item |
| 0x00272D40 | `Shuffle Songs` | Known | Menu item |
| 0x00272DA8 | `Extras` | Known | Menu item |
| 0x00272DC4 | `Photos` | Known | Menu item |
| 0x00272DD2 | `Composers` | Known | Menu item |
| 0x00272DE2 | `Genres` | Known | Menu item |
| 0x00272DF2 | `Audiobooks` | Known | Menu item |
| 0x00272E06 | `Podcasts` | Known | Menu item |
| 0x00272E36 | `Albums` | Known | Menu item |
| 0x00272E46 | `Artists` | Known | Menu item |
| 0x00272E56 | `Playlists` | Known | Menu item |
| 0x00272F54 | `Main Menu` | Known | Menu item |
| 0x00272FAC | `Audiobooks` | Known | Menu item |
| 0x00272FB8 | `Composers` | Known | Menu item |
| 0x00272FC4 | `Genres` | Known | Menu item |
| 0x00272FCC | `Podcasts` | Known | Menu item |
| 0x00272FE0 | `Albums` | Known | Menu item |
| 0x00272FE8 | `Artists` | Known | Menu item |
| 0x00272FF0 | `Playlists` | Known | Menu item |
| 0x00272FFC | `Settings` | Known | Menu item |
| 0x002731B4 | `Photos` | Known | Menu item |
| 0x00273550 | `Audiobooks` | Known | Menu item |
| 0x00273594 | `Settings` | Known | Menu item |
| 0x002735A0 | `Genres` | Known | Menu item |
| 0x002735A8 | `Artists` | Known | Menu item |
| 0x002735B0 | `Podcasts` | Known | Menu item |
| 0x002735CC | `Albums` | Known | Menu item |
| 0x002735D4 | `Composers` | Known | Menu item |
| 0x002735F0 | `Audiobooks` | Known | Menu item |
| 0x0027362C | `Playlists` | Known | Menu item |
| 0x00273698 | `Extras` | Known | Menu item |
| 0x002736E8 | `Main Menu` | Known | Menu item |

---

## 16. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001C9D54 | `iPod_Control\Device` | Filesystem Path | |
| 0x001C9D68 | `iPod_Control` | Filesystem Path | |
| 0x001C9D78 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x001CA518 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001CA588 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001CA5B0 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x001CA5E0 | `iPod_Control` | Filesystem Path | |
| 0x001CA604 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001CA634 | `System_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x001CA658 | `iPod_Control\Device` | Filesystem Path | |
| 0x001CA6A0 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x001CA6EC | `iPod_Control\Music\` | Filesystem Path | |
| 0x00295524 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x002A0EB1 | `iPod_Control/iTunes/` | Filesystem Path | |

---

## 17. EQ Presets

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
| 0x001D710C | `Acoustic` | EQ Preset | |
| 0x001D714C | `Electronic` | EQ Preset | |
| 0x001D7160 | `Hip Hop` | EQ Preset | |
| 0x001D7178 | `Loudness` | EQ Preset | |
| 0x001DB514 | `Hip Hop` | EQ Preset | |
| 0x001DB524 | `Latina` | EQ Preset | |
| 0x001DB52C | `Loudness` | EQ Preset | |
| 0x001DB538 | `Lounge` | EQ Preset | |
| 0x001DF9F0 | `Lounge` | EQ Preset | |
| 0x001E45DC | `Hip Hop` | EQ Preset | |
| 0x001E45EC | `Latino` | EQ Preset | |
| 0x001E45FC | `Lounge` | EQ Preset | |
| 0x001E8B94 | `Hip Hop` | EQ Preset | |
| 0x001E8BA4 | `Latina` | EQ Preset | |
| 0x001E8BAC | `Loudness` | EQ Preset | |
| 0x001E8BB8 | `Lounge` | EQ Preset | |
| 0x001ED0D0 | `Acoustic` | EQ Preset | |
| 0x001ED0DC | `Bass Booster` | EQ Preset | |
| 0x001ED0FC | `Classical` | EQ Preset | |
| 0x001ED118 | `Electronic` | EQ Preset | |
| 0x001ED12C | `Hip Hop` | EQ Preset | |
| 0x001ED144 | `Loudness` | EQ Preset | |
| 0x001ED150 | `Lounge` | EQ Preset | |
| 0x001ED174 | `Small Speakers` | EQ Preset | |
| 0x001ED184 | `Spoken Word` | EQ Preset | |
| 0x001ED190 | `Treble Booster` | EQ Preset | |
| 0x001ED1B0 | `Vocal Booster` | EQ Preset | |
| 0x001F1EA4 | `Acoustic` | EQ Preset | |
| 0x001F1EB0 | `Bass Booster` | EQ Preset | |
| 0x001F1ED0 | `Classical` | EQ Preset | |
| 0x001F1EEC | `Electronic` | EQ Preset | |
| 0x001F1F00 | `Hip Hop` | EQ Preset | |
| 0x001F1F18 | `Loudness` | EQ Preset | |
| 0x001F1F24 | `Lounge` | EQ Preset | |
| 0x001F1F48 | `Small Speakers` | EQ Preset | |
| 0x001F1F58 | `Spoken Word` | EQ Preset | |
| 0x001F1F64 | `Treble Booster` | EQ Preset | |
| 0x001F1F84 | `Vocal Booster` | EQ Preset | |
| 0x001F69EC | `Loudness` | EQ Preset | |
| 0x001F69F8 | `Lounge` | EQ Preset | |
| 0x001FAD28 | `Latino` | EQ Preset | |
| 0x001FAD30 | `Loudness` | EQ Preset | |
| 0x001FAD3C | `Lounge` | EQ Preset | |
| 0x001FF304 | `Acoustic` | EQ Preset | |
| 0x001FF310 | `Bass Booster` | EQ Preset | |
| 0x001FF330 | `Classical` | EQ Preset | |
| 0x001FF34C | `Electronic` | EQ Preset | |
| 0x001FF360 | `Hip Hop` | EQ Preset | |
| 0x001FF378 | `Loudness` | EQ Preset | |
| 0x001FF384 | `Lounge` | EQ Preset | |
| 0x001FF3A8 | `Small Speakers` | EQ Preset | |
| 0x001FF3B8 | `Spoken Word` | EQ Preset | |
| 0x001FF3C4 | `Treble Booster` | EQ Preset | |
| 0x001FF3E4 | `Vocal Booster` | EQ Preset | |
| 0x002036B8 | `Acoustic` | EQ Preset | |
| 0x002036C4 | `Bass Booster` | EQ Preset | |
| 0x002036E4 | `Classical` | EQ Preset | |
| 0x00203700 | `Electronic` | EQ Preset | |
| 0x00203714 | `Hip Hop` | EQ Preset | |
| 0x0020372C | `Loudness` | EQ Preset | |
| 0x00203738 | `Lounge` | EQ Preset | |
| 0x0020375C | `Small Speakers` | EQ Preset | |
| 0x0020376C | `Spoken Word` | EQ Preset | |
| 0x00203778 | `Treble Booster` | EQ Preset | |
| 0x00203798 | `Vocal Booster` | EQ Preset | |
| 0x002079DC | `Acoustic` | EQ Preset | |
| 0x002079E8 | `Bass Booster` | EQ Preset | |
| 0x00207A08 | `Classical` | EQ Preset | |
| 0x00207A24 | `Electronic` | EQ Preset | |
| 0x00207A38 | `Hip Hop` | EQ Preset | |
| 0x00207A50 | `Loudness` | EQ Preset | |
| 0x00207A5C | `Lounge` | EQ Preset | |
| 0x00207A80 | `Small Speakers` | EQ Preset | |
| 0x00207A90 | `Spoken Word` | EQ Preset | |
| 0x00207A9C | `Treble Booster` | EQ Preset | |
| 0x00207ABC | `Vocal Booster` | EQ Preset | |
| 0x00271430 | `Acoustic` | EQ Preset | |
| 0x0027143C | `Bass Booster` | EQ Preset | |
| 0x0027145C | `Classical` | EQ Preset | |
| 0x00271478 | `Electronic` | EQ Preset | |
| 0x0027148C | `Hip Hop` | EQ Preset | |
| 0x002714A4 | `Loudness` | EQ Preset | |
| 0x002714B0 | `Lounge` | EQ Preset | |
| 0x002714D4 | `Small Speakers` | EQ Preset | |
| 0x002714E4 | `Spoken Word` | EQ Preset | |
| 0x002714F0 | `Treble Booster` | EQ Preset | |
| 0x00271510 | `Vocal Booster` | EQ Preset | |

---

## 18. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001DC434 | `Error durante la importaci` | Diagnostic | |
| 0x001EAA18 | `Errore` | Diagnostic | |
| 0x001EAC14 | `Errore` | Diagnostic | |
| 0x002A2270 | `%s Error in file %s.` | Diagnostic | |
| 0x004FE0F0 | `Root hub Error Calling Add Device` | Diagnostic | |
| 0x004FE144 | `Root Hub Driver Internal Error unused case in hub handler` | Diagnostic | |

---

## 19. Assertions

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00005280 | `*** assertion failed: %s, file %s, line %d` | Assertion | |

---
