# iPod Mini 1st Gen (Silver) - RetailOS 1.4.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.4.1 |
| **IPSW** | iPod_3.1.4.1.ipsw |
| **Device** | iPod Mini 1st Gen (Silver) (2004, 4GB Microdrive, Click Wheel, Anodized Aluminum) |
| **UpdaterFamilyID** | 3 |
| **Binary Size** | 4,506,624 bytes (4.30 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 4,506,624 bytes |
| **Total Strings (>=4)** | 32,487 |
| **Function Prologues** | 10,416 (ARM: 7,563, Thumb: 2,853) |
| **DRAM References** | 26,472 |
| **Peripheral Refs** | 7,200 |
| **Build** | Unknown |
| **SoC** | PortalPlayer PP5020 |
| **Architecture** | ARM7TDMI (ARMv4T) dual-core |
| **Codename** | P86 |
| **DFU PID** | N/A |
| **SHA-256** | `a69031d594a0b54649c0a6cc087241808463b9d94a9e45793cedb7f02abd357f` |

---

## 1. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x003141C4 | `AudioCodecs` | Known | Audio system |

---

## 2. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168324 | `Task` | Known | RTOS task thread |
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
| 0x0016DAAB | `5RunTestsTask` | Known | RTOS task thread |
| 0x001BD794 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x001CEDB8 | `FWInterruptHandlerTask` | Known | RTOS task thread |
| 0x00314A94 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x00314AA8 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x00314ABC | `SBP2CommandTask` | Known | RTOS task thread |

---

## 3. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x003140BC | `AppleDRMVersion` | Known | DRM system |
| 0x003140F4 | `AppleDRM` | Known | DRM system |

---

## 4. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168FC0 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00169021 | `#!#iTunesDB` | Known | iTunes database |
| 0x00169030 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00169058 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x001690A0 | `System_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x001BD744 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x001CD5DD | `iPod_Control\iTunes\` | Known | iTunes database |

---

## 5. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001684B8 | `FirewireHandler` | Known | FireWire |
| 0x00168890 | `FirewireGuid` | Known | FireWire |
| 0x001708DC | `FireWire tilsluttet` | Known | FireWire |
| 0x00172D5F | `ffnen Sie das Adressbuch, Microsoft Entourage oder Palm Desktop und exportieren ` | Known | FireWire |
| 0x00173D5A | `ber FireWire verbunden` | Known | FireWire |
| 0x00176CC8 | `FireWire conectado` | Known | FireWire |
| 0x00179B18 | `FireWire liitetty` | Known | FireWire |
| 0x0017C585 | `utiliser comme disque FireWire. Puis glissez les vCards dans le dossier Contacts` | Known | FireWire |
| 0x0017D5DC | `FireWire Connect` | Known | FireWire |
| 0x001803C0 | `FireWire Connesso` | Known | FireWire |
| 0x00183BC0 | `FireWire ` | Known | FireWire |
| 0x00186D4C | `FireWire ` | Known | FireWire |
| 0x00189064 | `Op de iPod kunt u adres- en agendagegevens opslaan. Als u met iSync werkt (Mac O` | Known | FireWire |
| 0x0018A1A4 | `FireWire aangesloten` | Known | FireWire |
| 0x0018CF90 | `Koblet til via FireWire` | Known | FireWire |
| 0x0018F53E | `rst in din iPod som FireWire-h` | Known | FireWire |
| 0x00190308 | `FireWire anslutet` | Known | FireWire |
| 0x00193268 | `FireWire ` | Known | FireWire |
| 0x0019627C | `FireWire ` | Known | FireWire |
| 0x001A52BC | `FireWire Connected` | Known | FireWire |
| 0x00314214 | `FireWire` | Known | FireWire |
| 0x00314284 | `FireWireVersion` | Known | FireWire |
| 0x0031D7AD | `FireWire` | Known | FireWire |
| 0x0031DB03 | `<key>FireWireGUID</key>` | Known | FireWire |

---

## 6. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00168834 | `iPod_Control\Device` | Filesystem Path |  |
| 0x00168848 | `iPod_Control` | Filesystem Path |  |
| 0x00168858 | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x00169088 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001690C4 | `iPod_Control\Device` | Filesystem Path |  |
| 0x001690E4 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x0016912C | `iPod_Control\Music\` | Filesystem Path |  |

---

## 7. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016F520 | `Acoustic` | EQ Preset |  |
| 0x0016F52C | `Bass Booster` | EQ Preset |  |
| 0x0016F54C | `Classical` | EQ Preset |  |
| 0x0016F558 | `Dance` | EQ Preset |  |
| 0x0016F568 | `Electronic` | EQ Preset |  |
| 0x0016F57C | `Hip Hop` | EQ Preset |  |
| 0x0016F584 | `Jazz` | EQ Preset |  |
| 0x0016F58C | `Latin` | EQ Preset |  |
| 0x0016F594 | `Loudness` | EQ Preset |  |
| 0x0016F5A0 | `Lounge` | EQ Preset |  |
| 0x0016F5A8 | `Piano` | EQ Preset |  |
| 0x0016F5BC | `Rock` | EQ Preset |  |
| 0x0016F5C4 | `Small Speakers` | EQ Preset |  |
| 0x0016F5D4 | `Spoken Word` | EQ Preset |  |
| 0x0016F5E0 | `Treble Booster` | EQ Preset |  |
| 0x0016F600 | `Vocal Booster` | EQ Preset |  |
| 0x001728CC | `Acoustic` | EQ Preset |  |
| 0x001728FC | `Dance` | EQ Preset |  |
| 0x0017290C | `Electronic` | EQ Preset |  |
| 0x00172920 | `Hip Hop` | EQ Preset |  |
| 0x00172928 | `Jazz` | EQ Preset |  |
| 0x00172930 | `Latin` | EQ Preset |  |
| 0x00172938 | `Loudness` | EQ Preset |  |
| 0x0017294C | `Piano` | EQ Preset |  |
| 0x00172960 | `Rock` | EQ Preset |  |
| 0x0017581C | `Dance` | EQ Preset |  |
| 0x00175844 | `Hip Hop` | EQ Preset |  |
| 0x0017584C | `Jazz` | EQ Preset |  |
| 0x0017585C | `Loudness` | EQ Preset |  |
| 0x00175868 | `Lounge` | EQ Preset |  |
| 0x00175870 | `Piano` | EQ Preset |  |
| 0x00175884 | `Rock` | EQ Preset |  |
| 0x00178804 | `Hip Hop` | EQ Preset |  |
| 0x0017880C | `Jazz` | EQ Preset |  |
| 0x00178814 | `Latin` | EQ Preset |  |
| 0x00178828 | `Lounge` | EQ Preset |  |
| 0x00178830 | `Piano` | EQ Preset |  |
| 0x00178844 | `Rock` | EQ Preset |  |
| 0x0017C098 | `Dance` | EQ Preset |  |
| 0x0017C0C0 | `Hip Hop` | EQ Preset |  |
| 0x0017C0C8 | `Jazz` | EQ Preset |  |
| 0x0017C0D0 | `Latin` | EQ Preset |  |
| 0x0017C0D8 | `Loudness` | EQ Preset |  |
| 0x0017C0EC | `Piano` | EQ Preset |  |
| 0x0017C104 | `Rock` | EQ Preset |  |
| 0x0017F0B8 | `Dance` | EQ Preset |  |
| 0x0017F0DC | `Hip Hop` | EQ Preset |  |
| 0x0017F0E4 | `Jazz` | EQ Preset |  |
| 0x0017F0F4 | `Loudness` | EQ Preset |  |
| 0x0017F100 | `Lounge` | EQ Preset |  |
| 0x0017F108 | `Piano` | EQ Preset |  |
| 0x0017F11C | `Rock` | EQ Preset |  |
| 0x00182130 | `Acoustic` | EQ Preset |  |
| 0x0018213C | `Bass Booster` | EQ Preset |  |
| 0x0018215C | `Classical` | EQ Preset |  |
| 0x00182168 | `Dance` | EQ Preset |  |
| 0x00182178 | `Electronic` | EQ Preset |  |
| 0x0018218C | `Hip Hop` | EQ Preset |  |
| 0x00182194 | `Jazz` | EQ Preset |  |
| 0x0018219C | `Latin` | EQ Preset |  |
| 0x001821A4 | `Loudness` | EQ Preset |  |
| 0x001821B0 | `Lounge` | EQ Preset |  |
| 0x001821B8 | `Piano` | EQ Preset |  |
| 0x001821CC | `Rock` | EQ Preset |  |
| 0x001821D4 | `Small Speakers` | EQ Preset |  |
| 0x001821E4 | `Spoken Word` | EQ Preset |  |
| 0x001821F0 | `Treble Booster` | EQ Preset |  |
| 0x00182210 | `Vocal Booster` | EQ Preset |  |
| 0x001857E0 | `Acoustic` | EQ Preset |  |
| 0x001857EC | `Bass Booster` | EQ Preset |  |
| 0x0018580C | `Classical` | EQ Preset |  |
| 0x00185818 | `Dance` | EQ Preset |  |
| 0x00185828 | `Electronic` | EQ Preset |  |
| 0x0018583C | `Hip Hop` | EQ Preset |  |
| 0x00185844 | `Jazz` | EQ Preset |  |
| 0x0018584C | `Latin` | EQ Preset |  |
| 0x00185854 | `Loudness` | EQ Preset |  |
| 0x00185860 | `Lounge` | EQ Preset |  |
| 0x00185868 | `Piano` | EQ Preset |  |
| 0x0018587C | `Rock` | EQ Preset |  |
| 0x00185884 | `Small Speakers` | EQ Preset |  |
| 0x00185894 | `Spoken Word` | EQ Preset |  |
| 0x001858A0 | `Treble Booster` | EQ Preset |  |
| 0x001858C0 | `Vocal Booster` | EQ Preset |  |
| 0x00188D74 | `Dance` | EQ Preset |  |
| 0x00188DA8 | `Jazz` | EQ Preset |  |
| 0x00188DB0 | `Latin` | EQ Preset |  |
| 0x00188DB8 | `Loudness` | EQ Preset |  |
| 0x00188DC4 | `Lounge` | EQ Preset |  |
| 0x00188DCC | `Piano` | EQ Preset |  |
| 0x00188DDC | `Rock` | EQ Preset |  |
| 0x0018BC24 | `Dance` | EQ Preset |  |
| 0x0018BC48 | `Hip Hop` | EQ Preset |  |
| 0x0018BC50 | `Jazz` | EQ Preset |  |
| 0x0018BC60 | `Loudness` | EQ Preset |  |
| 0x0018BC6C | `Lounge` | EQ Preset |  |
| 0x0018BC74 | `Piano` | EQ Preset |  |
| 0x0018BC88 | `Rock` | EQ Preset |  |
| 0x0018EF34 | `Acoustic` | EQ Preset |  |
| 0x0018EF40 | `Bass Booster` | EQ Preset |  |
| 0x0018EF60 | `Classical` | EQ Preset |  |
| 0x0018EF6C | `Dance` | EQ Preset |  |
| 0x0018EF7C | `Electronic` | EQ Preset |  |
| 0x0018EF90 | `Hip Hop` | EQ Preset |  |
| 0x0018EF98 | `Jazz` | EQ Preset |  |
| 0x0018EFA0 | `Latin` | EQ Preset |  |
| 0x0018EFA8 | `Loudness` | EQ Preset |  |
| 0x0018EFB4 | `Lounge` | EQ Preset |  |
| 0x0018EFBC | `Piano` | EQ Preset |  |
| 0x0018EFD0 | `Rock` | EQ Preset |  |
| 0x0018EFD8 | `Small Speakers` | EQ Preset |  |
| 0x0018EFE8 | `Spoken Word` | EQ Preset |  |
| 0x0018EFF4 | `Treble Booster` | EQ Preset |  |
| 0x0018F014 | `Vocal Booster` | EQ Preset |  |
| 0x00191F2C | `Acoustic` | EQ Preset |  |
| 0x00191F38 | `Bass Booster` | EQ Preset |  |
| 0x00191F58 | `Classical` | EQ Preset |  |
| 0x00191F64 | `Dance` | EQ Preset |  |
| 0x00191F74 | `Electronic` | EQ Preset |  |
| 0x00191F88 | `Hip Hop` | EQ Preset |  |
| 0x00191F90 | `Jazz` | EQ Preset |  |
| 0x00191F98 | `Latin` | EQ Preset |  |
| 0x00191FA0 | `Loudness` | EQ Preset |  |
| 0x00191FAC | `Lounge` | EQ Preset |  |
| 0x00191FB4 | `Piano` | EQ Preset |  |
| 0x00191FC8 | `Rock` | EQ Preset |  |
| 0x00191FD0 | `Small Speakers` | EQ Preset |  |
| 0x00191FE0 | `Spoken Word` | EQ Preset |  |
| 0x00191FEC | `Treble Booster` | EQ Preset |  |
| 0x0019200C | `Vocal Booster` | EQ Preset |  |
| 0x00194E9C | `Acoustic` | EQ Preset |  |
| 0x00194EA8 | `Bass Booster` | EQ Preset |  |
| 0x00194EC8 | `Classical` | EQ Preset |  |
| 0x00194ED4 | `Dance` | EQ Preset |  |
| 0x00194EE4 | `Electronic` | EQ Preset |  |
| 0x00194EF8 | `Hip Hop` | EQ Preset |  |
| 0x00194F00 | `Jazz` | EQ Preset |  |
| 0x00194F08 | `Latin` | EQ Preset |  |
| 0x00194F10 | `Loudness` | EQ Preset |  |
| 0x00194F1C | `Lounge` | EQ Preset |  |
| 0x00194F24 | `Piano` | EQ Preset |  |
| 0x00194F38 | `Rock` | EQ Preset |  |
| 0x00194F40 | `Small Speakers` | EQ Preset |  |
| 0x00194F50 | `Spoken Word` | EQ Preset |  |
| 0x00194F5C | `Treble Booster` | EQ Preset |  |
| 0x00194F7C | `Vocal Booster` | EQ Preset |  |
| 0x001A3C04 | `Acoustic` | EQ Preset |  |
| 0x001A3C10 | `Bass Booster` | EQ Preset |  |
| 0x001A3C30 | `Classical` | EQ Preset |  |
| 0x001A3C3C | `Dance` | EQ Preset |  |
| 0x001A3C4C | `Electronic` | EQ Preset |  |
| 0x001A3C60 | `Hip Hop` | EQ Preset |  |
| 0x001A3C68 | `Jazz` | EQ Preset |  |
| 0x001A3C70 | `Latin` | EQ Preset |  |
| 0x001A3C78 | `Loudness` | EQ Preset |  |
| 0x001A3C84 | `Lounge` | EQ Preset |  |
| 0x001A3C8C | `Piano` | EQ Preset |  |
| 0x001A3CA0 | `Rock` | EQ Preset |  |
| 0x001A3CA8 | `Small Speakers` | EQ Preset |  |
| 0x001A3CB8 | `Spoken Word` | EQ Preset |  |
| 0x001A3CC4 | `Treble Booster` | EQ Preset |  |
| 0x001A3CE4 | `Vocal Booster` | EQ Preset |  |

---
