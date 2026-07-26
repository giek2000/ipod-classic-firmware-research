# iPod 5.5G (Enhanced/Search) - RetailOS 25.1.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 25.1.3 |
| **IPSW** | iPod_25.1.3.ipsw |
| **Device** | iPod 5.5G Enhanced (2006, 30/80GB, Brighter Screen, Search) |
| **Binary Size** | 13,903,872 bytes (13.26 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 13,903,872 bytes |
| **Total Strings (>=6)** | 30,181 |
| **Function Prologues** | 12,890 |
| **SoC** | PortalPlayer PP5022C / Broadcom BCM2722 |
| **Architecture** | ARM7TDMI (ARMv4T), dual-core + video DSP |
| **Encrypted** | No |
| **Decryption Method** | Extract from IPSW (unencrypted) |
| **SHA-256** | `7830d1345aa2313db154e06ae93f2b5961e1cb04e8edeaae27dadc303e0d9fb3` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00107680 | `Root Hub Driver Internal Error unused case in hub handl...` | Hidden | Undocumented UI |
| 0x0016F904 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x002166F4 | `Channel Reserved` | Hidden | Logging/Telemetry |
| 0x00216708 | `Channel AppBoot` | Hidden | Logging/Telemetry |
| 0x00216718 | `Channel BufferedSongReading` | Hidden | Logging/Telemetry |
| 0x00216734 | `Channel PrefsWriting` | Hidden | Logging/Telemetry |
| 0x0021674C | `Channel GeneralUserExperience` | Hidden | Logging/Telemetry |
| 0x0021676C | `Channel PlayFromDisk` | Hidden | Logging/Telemetry |
| 0x00216784 | `Channel CacheSpinupDrive` | Hidden | Logging/Telemetry |
| 0x002167A0 | `Channel TestLogging` | Hidden | Logging/Telemetry |
| 0x002167B4 | `Channel AppFileLoading` | Hidden | Logging/Telemetry |
| 0x002167CC | `Channel VCardReading` | Hidden | Logging/Telemetry |
| 0x002167E4 | `Channel LongSongScanning` | Hidden | Logging/Telemetry |
| 0x00216858 | `Channel VoiceRecording` | Hidden | Logging/Telemetry |
| 0x00216870 | `Channel VoiceRecordingNewFileSegment` | Hidden | Logging/Telemetry |
| 0x00216898 | `Channel PhotoBrowse` | Hidden | Logging/Telemetry |
| 0x002168AC | `Channel PhotoImporting` | Hidden | Logging/Telemetry |
| 0x002168C4 | `Channel Notes` | Hidden | Logging/Telemetry |
| 0x002168D4 | `Channel PhotoFileManagement` | Hidden | Logging/Telemetry |
| 0x002168F0 | `Channel DiskModeChannel` | Hidden | Logging/Telemetry |
| 0x00216908 | `Channel FirewireChannel` | Hidden | Logging/Telemetry |
| 0x00216920 | `Channel USBChannel` | Hidden | Logging/Telemetry |
| 0x00216934 | `Channel UnitTests` | Hidden | Hidden Test |
| 0x00216948 | `Channel FreeSpaceCache` | Hidden | Logging/Telemetry |
| 0x002169C0 | `Channel OnTheGoFileMgmt` | Hidden | Logging/Telemetry |
| 0x002169D8 | `Channel SlideShow` | Hidden | Logging/Telemetry |
| 0x002169EC | `Channel ImageCache` | Hidden | Logging/Telemetry |
| 0x00216A00 | `Channel AlbumArtReading` | Hidden | Logging/Telemetry |
| 0x00216A18 | `Channel Video` | Hidden | Logging/Telemetry |
| 0x00216A28 | `Channel DiskImage` | Hidden | Logging/Telemetry |
| 0x00216A3C | `Channel ResourceAccess` | Hidden | Logging/Telemetry |
| 0x00216A54 | `Channel VideoCoreBoot` | Hidden | Logging/Telemetry |
| 0x00216A6C | `Channel DiskFormatConvert` | Hidden | Logging/Telemetry |
| 0x00216A88 | `Channel StreamCacheAddFile` | Hidden | Logging/Telemetry |
| 0x00216AA4 | `Channel FontFileAccess` | Hidden | Logging/Telemetry |
| 0x00216ABC | `Channel ScreenLock` | Hidden | Logging/Telemetry |
| 0x00216B28 | `Channel DiskReaderTask` | Hidden | Logging/Telemetry |
| 0x00216B40 | `Channel ProfilerAccess` | Hidden | Logging/Telemetry |
| 0x00216B58 | `Channel eAppAccess` | Hidden | Logging/Telemetry |
| 0x00216B6C | `Channel eAppWriteBackCache` | Hidden | Logging/Telemetry |
| 0x00216B88 | `Channel TrainerFileAccess` | Hidden | Logging/Telemetry |
| 0x00216BA4 | `Channel IapStorage` | Hidden | Logging/Telemetry |
| 0x00216BB8 | `Channel XMLParsing` | Hidden | Logging/Telemetry |
| 0x00216BCC | `Channel AudioPrompt` | Hidden | Logging/Telemetry |
| 0x00216BE0 | `Channel AudioPromptXML` | Hidden | Logging/Telemetry |
| 0x00216BF8 | `Channel StreamCacheSeek` | Hidden | Logging/Telemetry |
| 0x00216C10 | `Channel PredictiveCacheSpinup` | Hidden | Logging/Telemetry |
| 0x0021753C | `iPod Usage Stats` | Hidden | Logging/Telemetry |
| 0x002186B8 | `Flush Usage Log Data` | Hidden | Logging/Telemetry |

---

## 2. Discovered Features

### EQ Preset

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001FE31C | `Disp seq_num` | EQ Preset | |
| 0x001FE384 | `core_freq_khz` | EQ Preset | |
| 0x002177F1 | `Total time in deep sleep: %d seconds` | EQ Preset | |
| 0x00217819 | `Deep sleep was entered %d %s` | EQ Preset | |
| 0x00218788 | `Enter Deep Sleep` | EQ Preset | |
| 0x0021879C | `Exit Deep Sleep` | EQ Preset | |
| 0x002B2F98 | `Acoustic` | EQ Preset | |
| 0x002B2FA4 | `Bass Booster` | EQ Preset | |
| 0x002B2FC4 | `Classical` | EQ Preset | |
| 0x002B2FE0 | `Electronic` | EQ Preset | |
| 0x002B2FF4 | `Hip Hop` | EQ Preset | |
| 0x002B300C | `Loudness` | EQ Preset | |
| 0x002B3018 | `Lounge` | EQ Preset | |
| 0x002B303C | `Small Speakers` | EQ Preset | |
| 0x002B304C | `Spoken Word` | EQ Preset | |
| 0x002B3058 | `Treble Booster` | EQ Preset | |
| 0x002B3078 | `Vocal Booster` | EQ Preset | |
| 0x002B959C | `Treble Boost` | EQ Preset | |
| 0x002B95AC | `Bass Boost` | EQ Preset | |
| 0x002CF6AC | `Latina` | EQ Preset | |
| 0x002DCB48 | `rique latine` | EQ Preset | |
| 0x00304FC8 | `Latino` | EQ Preset | |
| 0x00630C1B | `~ BR&B$"` | EQ Preset | |
| 0x0067DA3D | `LATIN-1` | EQ Preset | |
| 0x0067DA45 | `LATIN1` | EQ Preset | |
| 0x0068106D | `Secure Electronic Transactions` | EQ Preset | |
| 0x00B286C0 | `hostreq_notify` | EQ Preset | |
| 0x00B286CF | `hostreq_read_iphoto_block` | EQ Preset | |
| 0x00B286E9 | `hostreq_rendertext` | EQ Preset | |
| 0x00B86CC1 | `Disp seq_num=%u` | EQ Preset | |

### Localization

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x004B5DDC | `English` | Localization | |
| 0x004B5E14 | `Italiano` | Localization | |
| 0x0067D7D8 | `x-mac-japanese` | Localization | |
| 0x0067DA75 | `X-MAC-CHINESETRAD` | Localization | |
| 0x0067DA87 | `X-MAC-JAPANESE` | Localization | |
| 0x0067DA96 | `MACJAPANESE` | Localization | |
| 0x0067DAB5 | `X-MAC-KOREAN` | Localization | |
| 0x0067DAD7 | `X-MAC-CHINESESIMP` | Localization | |

### Filesystem Path

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000986D0 | `iPod_Control` | Filesystem Path | |
| 0x000986FC | `iPod_Control\Device` | Filesystem Path | |
| 0x000A6F48 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x000B95E0 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000BC108 | `iPod_Control\Music\` | Filesystem Path | |
| 0x000C06F4 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x000E3894 | `iPod_Control/%s%s%s` | Filesystem Path | |
| 0x000E38A8 | `iPod_Control/%s/%s%s%s` | Filesystem Path | |
| 0x000F06C8 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x001AC3AC | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001ACD40 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path | |
| 0x001D7F64 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0024723C | `iPod_Control/Device` | Filesystem Path | |
| 0x00247250 | `iPod_Control/Device/radio` | Filesystem Path | |
| 0x0067D0BB | `iPod_Control/games_RO/` | Filesystem Path | |
| 0x0067D185 | `iPod_Control/Device/accessories` | Filesystem Path | |
| 0x0067D4B8 | `iPod_Control/iTunes/` | Filesystem Path | |

---

## 3. Known User-Facing Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0007E490 | `Settings` | Known | User setting |
| 0x001440D4 | `Memory full. %d notes loaded, some notes not loaded. (3...` | Known | UI element |
| 0x0017B88C | `CanFlashBacklight` | Known | UI element |
| 0x001B0768 | `KeyRepeatTimer` | Known | UI element |
| 0x001C4000 | `Could not find settingsHandler for pid %d` | Known | User setting |
| 0x001DDA70 | `Contextual menu up!` | Known | Menu item |
| 0x002183D0 | `Backlight` | Known | UI element |
| 0x00218698 | `Backlight On` | Known | UI element |
| 0x002186A8 | `Backlight Off` | Known | UI element |
| 0x002B88B8 | `Alarmer` | Known | UI element |
| 0x002B8A3B | `kke tekstarkiverne til mappen Notes p` | Known | UI element |
| 0x002B904B | ` menuknappen for at annullere.` | Known | Menu item |
| 0x002BB33C | `Nulstil menu` | Known | Menu item |
| 0x002BB5A4 | `Hovedmenu` | Known | Menu item |
| 0x002BBBC4 | `Menuer` | Known | Menu item |
| 0x002BFB80 | `Extras` | Known | UI element |
| 0x002C0C98 | `Contacts` | Known | UI element |
| 0x002C5C10 | ` Notes ` | Known | UI element |
| 0x002C8578 | `Shuffle %s` | Known | UI element |
| 0x002CE8AC | `Calendario` | Known | UI element |
| 0x002CE8B8 | `Calendarios` | Known | UI element |
| 0x002CE900 | `Alarmas` | Known | UI element |
| 0x002CEA89 | `n de usar el iPod como disco y arrastrar los archivos d...` | Known | UI element |
| 0x002D0400 | `mo sincronizar contactos, calendarios y listas de tarea...` | Known | UI element |
| 0x002D0674 | `Alarma` | Known | UI element |
| 0x002D192C | `Hora alarma` | Known | UI element |
| 0x002D1FC8 | `Contraste` | Known | UI element |
| 0x002D50DA | ` sitten tekstitiedostot iPodin Notes-kansioon. Lis` | Known | UI element |
| 0x002D56C5 | `n. Kumoa painamalla menu-painiketta.` | Known | Menu item |
| 0x002DBCA0 | `Alarmes` | Known | UI element |
| 0x002DBE46 | `utilisation comme disque dur puis faites glisser ces fi...` | Known | UI element |
| 0x002DC585 | `es de la liste. Cliquez le bouton central pour lancer l...` | Known | Menu item |
| 0x002DD868 | `chargement des contacts.` | Known | UI element |
| 0x002DD921 | `iPod pour obtenir des instructions sur la synchronisati...` | Known | UI element |
| 0x002DDBA4 | ` Contacts ` | Known | UI element |
| 0x002DDBCC | `Alarme` | Known | UI element |
| 0x002DDBD4 | `Chargement des contacts.` | Known | UI element |
| 0x002DE11C | `Chargement des notes.` | Known | UI element |
| 0x002DECB7 | `initialiser le menu principal` | Known | Menu item |
| 0x002DEF84 | `Menu principal` | Known | Menu item |
| 0x002DF016 | `alarme` | Known | UI element |
| 0x002E29A2 | ` Notes mapp` | Known | UI element |
| 0x002E90F0 | `Calendari` | Known | UI element |
| 0x002E928C | `Per visualizzare documenti di testo qui, abilita iPod p...` | Known | UI element |
| 0x002EABA8 | ` di sincronizzazione dei contatti, dei calendari e degl...` | Known | UI element |
| 0x002EBCFC | `Ora legale` | Known | UI element |
| 0x002EBD44 | `Ripristina menu principale` | Known | Menu item |
| 0x002EC6A8 | `Contrasto` | Known | UI element |
| 0x002FE5D8 | `Met 'Zoek zenders' zoekt u naar alle beschikbare radioz...` | Known | Menu item |
| 0x002FEA90 | `Shuffle nummers` | Known | UI element |
| 0x00300A18 | `Shuffle foto's` | Known | UI element |
| 0x00300B34 | `Herstel menu` | Known | Menu item |
| 0x00300D84 | `Shuffle` | Known | UI element |
| 0x00300D8C | `Hoofdmenu` | Known | Menu item |
| 0x0030140C | `Menu's` | Known | Menu item |
| 0x00301434 | `Contrast` | Known | UI element |
| 0x00307058 | `Alarmtidspunkt` | Known | UI element |
| 0x0030A594 | `Alarmy` | Known | UI element |
| 0x0030A743 | `gnij te pliki tekstowe do teczki Notes w iPodzie. Szcze...` | Known | UI element |
| 0x0030AD98 | ` skanowanie i menu przycisk, by odwo` | Known | Menu item |
| 0x0030D290 | `Wyzeruj menu g` | Known | Menu item |
| 0x0030D534 | `Menu g` | Known | Menu item |
| 0x0030D5B0 | `Czas alarmu` | Known | UI element |
| 0x00310DA7 | `o como disco e, em seguida, arraste ficheiros de texto ...` | Known | UI element |
| 0x003114AA | `o de menu para cancelar.` | Known | Menu item |
| 0x003138C8 | `Repor menu pri.` | Known | Menu item |
| 0x0032026E | `ndning av iPod som extern enhet och drar sedan textfile...` | Known | UI element |
| 0x00322E28 | `Alarmtid` | Known | UI element |
| 0x0032631C | `Alarmlar` | Known | UI element |
| 0x003264F1 | ` iPod'daki Notes klas` | Known | UI element |
| 0x00326BA0 | `in Menu d` | Known | Menu item |
| 0x003292E8 | `Alarm Zaman` | Known | UI element |
| 0x0032D05E | ` Menu ` | Known | Menu item |
| 0x003334EE | ` menu ` | Known | Menu item |
| 0x004B5178 | `Calendar` | Known | UI element |
| 0x004B5184 | `Calendars` | Known | UI element |
| 0x004B51C4 | `Alarms` | Known | UI element |
| 0x004B55F0 | `Slideshow Settings` | Known | User setting |
| 0x004B57FC | `Find Stations will scan through all available radio sta...` | Known | Menu item |
| 0x004B5BC8 | `Now Playing` | Known | UI element |
| 0x004B5CF4 | `Shuffle Songs` | Known | UI element |
| 0x004B5DA8 | `Volume Limit` | Known | UI element |
| 0x004B6058 | `New Clock` | Known | UI element |
| 0x004B6F64 | `contacts loading.` | Known | UI element |
| 0x004B7264 | `Contacts loading.` | Known | UI element |
| 0x004B769C | `Notes loading.` | Known | UI element |
| 0x004B7D74 | `Delete This Clock` | Known | UI element |
| 0x004B7DE8 | `Sleep Timer` | Known | UI element |
| 0x004B7DF4 | `Alarm Clock` | Known | UI element |
| 0x004B7E00 | `World Clock` | Known | UI element |
| 0x004B7E3C | `Video Settings` | Known | User setting |
| 0x004B7F70 | `Shuffle Photos` | Known | UI element |
| 0x004B7F80 | `Repeat` | Known | UI element |
| 0x004B80B8 | `Reset Main Menu` | Known | Menu item |
| 0x004B824C | `Reset All Settings` | Known | User setting |
| 0x004B8308 | `Backlight Timer` | Known | UI element |
| 0x004B8328 | `Main Menu` | Known | Menu item |
| 0x004B8390 | `Alarm Time` | Known | UI element |
| 0x004B83B0 | `Delete Clock` | Known | UI element |
| 0x004B8438 | `Radio Settings` | Known | User setting |
| 0x004B89B4 | `Reset All` | Known | UI element |
| 0x00504484 | `TCalendarCntlr_Alarm` | Known | UI element |
| 0x00505068 | `To check song links, set the preference NotesOnly to tr...` | Known | UI element |
| 0x005050CC | `The NotesOnly pref can only be set globally in the Pref...` | Known | UI element |
| 0x0050511C | `Warning: Preferences file must be in the Notes folder, ...` | Known | UI element |
| 0x0066A4D9 | `Illegal instruction` | Known | UI element |
| 0x0066A507 | `Illegal address` | Known | UI element |
| 0x0067D8FA | `dalarm` | Known | UI element |
| 0x0067D901 | `valarm` | Known | UI element |
| 0x0067D953 | `vcalendar` | Known | UI element |
| 0x0067DE8C | `NotesOnly` | Known | UI element |
| 0x00AFECC9 | `backlight` | Known | UI element |
| 0x00B282CE | `audio_playgetclock` | Known | UI element |
| 0x00B282E1 | `audio_playgetendclock` | Known | UI element |
| 0x00B28989 | `TMT_Retrieve_Clock` | Known | UI element |
| 0x00B86D7D | `framenum_checksum=%u` | Known | Menu item |
| 0x00B8FE9A | `mp_clock_init` | Known | UI element |
| 0x00B8FEA8 | `mp_clock_destroy` | Known | UI element |
| 0x00B8FEB9 | `mp_clock_reset` | Known | UI element |
| 0x00B8FEC8 | `clock_stop` | Known | UI element |
| 0x00B8FED3 | `mp_clock_start` | Known | UI element |
| 0x00B8FEE2 | `clock_fetch` | Known | UI element |
| 0x00B8FEEE | `mp_clock_fetch` | Known | UI element |
| 0x00B8FEFD | `mp_clock_fetch2` | Known | UI element |
| 0x00B8FF0D | `mp_clock_audio` | Known | UI element |
| 0x00B8FF1C | `mp_clock_stc` | Known | UI element |
| 0x00B8FF29 | `mp_clock_stop` | Known | UI element |
| 0x00B8FF37 | `mp_clock_set` | Known | UI element |

---

## 4. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00099174 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x000C9B20 | `USB Secondary Interrupt Task` | Known | RTOS task thread |
| 0x000E334C | `ICAPTPCameraIOTask` | Known | RTOS task thread |
| 0x0011DBA8 | `USBStatusTask` | Known | RTOS task thread |
| 0x0011DBC4 | `USBTaskTimeTask` | Known | RTOS task thread |
| 0x0014D544 | `RtcTaskClass` | Known | RTOS task thread |
| 0x00169A2C | `VCUpdateTask` | Known | RTOS task thread |
| 0x0016F904 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x0017587C | `USBDeviceTask` | Known | RTOS task thread |
| 0x0017BFB4 | `DiskReaderTask` | Known | RTOS task thread |
| 0x00188800 | `TimerTaskClass` | Known | RTOS task thread |
| 0x0018B3BC | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0018B3D0 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0019E7F8 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x001BEB58 | `LoadDataTasks` | Known | RTOS task thread |
| 0x00210664 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x00216B28 | `Channel DiskReaderTask` | Known | RTOS task thread |
| 0x00261E3C | `FirewireTask` | Known | RTOS task thread |
| 0x00261E54 | `OptoTask` | Known | RTOS task thread |
| 0x00261E64 | `SerialOptoTask` | Known | RTOS task thread |
| 0x00261E78 | `BacklightTask` | Known | RTOS task thread |
| 0x00261E8C | `CNATask` | Known | RTOS task thread |
| 0x00261EAC | `DiskMgrTask` | Known | RTOS task thread |
| 0x00261EBC | `HoldSwitchTask` | Known | RTOS task thread |
| 0x00261ED0 | `TopPlugTask` | Known | RTOS task thread |
| 0x00261EE0 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x00261EF4 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x00261F0C | `AccessoryDetectTask` | Known | RTOS task thread |
| 0x00261F34 | `AlarmTask` | Known | RTOS task thread |
| 0x00261F44 | `WatchdogTask` | Known | RTOS task thread |
| 0x00261FBC | `USBAudioTask` | Known | RTOS task thread |
| 0x002AE410 | `HostOSTask` | Known | RTOS task thread |
| 0x002AEFFC | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x002E3E38 | `Taskent` | Known | RTOS task thread |
| 0x0050535C | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x00505374 | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x00505498 | `VideoDaisyTask` | Known | RTOS task thread |
| 0x00B28868 | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B28878 | `TCC_Current_Task_Pointer` | Known | RTOS task thread |
| 0x00B288A1 | `TCC_Delete_Task` | Known | RTOS task thread |
| 0x00B288C0 | `TCC_Reset_Task` | Known | RTOS task thread |
| 0x00B288E2 | `TCC_Task_Sleep` | Known | RTOS task thread |
| 0x00B288F1 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B28904 | `TCF_Task_Information` | Known | RTOS task thread |

---

## 5. Audio/Codec

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000D9FE0 | `RIFFWAVEfmt data"V` | Known | PCM audio format |
| 0x0015E9E4 | `RIFFWAVEfmt data` | Known | PCM audio format |
| 0x0017A888 | `AudioCodecs` | Known | Audio system |
| 0x0017A964 | `Audible` | Known | Audible audiobook format |
| 0x0017B928 | `VideoCodecs` | Known | Audio system |
| 0x0017BA0C | `H.264LC` | Known | Audio system |
| 0x001FE1B8 | `max_decoded_buffer` | Known | Decoder component |
| 0x001FE1CC | `min_decoded_buffer` | Known | Decoder component |
| 0x001FE288 | `total_bytes_decoded` | Known | Decoder component |
| 0x001FE29C | `max_decoded_bytes` | Known | Decoder component |
| 0x001FE33C | `Disp decoded data` | Known | Decoder component |
| 0x001FE858 | `lost frames - decoder (failed to decode):` | Known | Decoder component |
| 0x001FE888 | `max decoded buffer:` | Known | Decoder component |
| 0x001FE8A0 | `min decoded buffer:` | Known | Decoder component |
| 0x001FE8D0 | `total bytes decoded:` | Known | Decoder component |
| 0x001FE8EC | `max decoded bytes:` | Known | Decoder component |
| 0x002623CF | `@mp3dec_sync` | Known | MP3 codec |
| 0x00262C17 | `@mp4_aacdec_sync` | Known | AAC codec |
| 0x002B451D | ` Audible v` | Known | Audible audiobook format |
| 0x002B456F | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x002B4585 | ` 2002 Audible, Inc. V` | Known | Audible audiobook format |
| 0x002B4762 | ` zvuku MPEG Layer-3 byla poskytnuta spole` | Known | Audio system |
| 0x002B478D | `nostmi Fraunhofer IIS a` | Known | Audio system |
| 0x002BA9A8 | `Audible-softwaren i dette produkt bruges i henhold til ...` | Known | Audible audiobook format |
| 0x002BAA08 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x002BAAFA | `.net codec i dette produkt bruges i henhold til en lice...` | Known | Audio system |
| 0x002BABA4 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunh...` | Known | Audio system |
| 0x002C1484 | `Die Audible Software in diesem Produkt wird in Lizenz d...` | Known | Audible audiobook format |
| 0x002C14DD | ` 2002 Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x002C15C9 | `.net Codec in diesem Produkt wird in Lizenz der VoiceAg...` | Known | Audio system |
| 0x002C1687 | `r MPEG Layer-3 wurde lizenziert von Fraunhofer IIS und ...` | Known | Audio system |
| 0x002C9563 | ` Audible ` | Known | Audible audiobook format |
| 0x002C95C0 | ` Audible. ` | Known | Audible audiobook format |
| 0x002C95F6 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002C9774 | `.net codec ` | Known | Audio system |
| 0x002C98BB | ` MPEG Layer-3 ` | Known | Audio system |
| 0x002C98F9 | ` Fraunhofer IIS ` | Known | Audio system |
| 0x002D0BA8 | `El software Audible incluido en este producto se usa ba...` | Known | Audible audiobook format |
| 0x002D0C03 | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x002D0DA1 | `n de audio MPEG Layer-3 utilizada bajo licencia de Frau...` | Known | Audio system |
| 0x002D7096 | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright...` | Known | Audible audiobook format |
| 0x002D70D0 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x002D723C | `MPEG Layer-3 -` | Known | Audio system |
| 0x002D724E | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | Audio system |
| 0x002DE1D8 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x002DE222 | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x002DE237 | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x002DE2E8 | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x002DE3BC | `La technologie de codage audio MPEG Layer 3 est utilis` | Known | Audio system |
| 0x002DE3F4 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | Audio system |
| 0x002E4C42 | ` Audible szoftver az Audible licence alatt van haszn` | Known | Audible audiobook format |
| 0x002E4C8C | ` 2002, Audible, Inc. Minden jog fenntartva.` | Known | Audible audiobook format |
| 0x002E4D81 | `.net codec a VoiceAge Coporation c` | Known | Audio system |
| 0x002E4E14 | `Az MPEG Layer-3 hangk` | Known | Audio system |
| 0x002E4E3C | `gia a Fraunhofer IIS ` | Known | Audio system |
| 0x002E519C | `l alacsony.` | Known | Apple Lossless codec |
| 0x002EA63C | `La Mecca` | Known | Audio system |
| 0x002EB3C8 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x002EB3F1 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x002EB420 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x002EB492 | ` utilizzato su licenza da VoiceAge Corporation. Il code...` | Known | Audio system |
| 0x002EB568 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da...` | Known | Audio system |
| 0x002F2562 | `Audible ` | Known | Audible audiobook format |
| 0x002F25BB | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002F2770 | `MPEG Layer-3 ` | Known | Audio system |
| 0x002F27BC | `Fraunhofer IIS ` | Known | Audio system |
| 0x002F960E | ` Audible` | Known | Audible audiobook format |
| 0x002F9742 | `.net codec` | Known | Audio system |
| 0x002F9803 | ` Fraunhofer IIS` | Known | Audio system |
| 0x0030012C | `De Audible-software in dit product wordt gebruikt in li...` | Known | Audible audiobook format |
| 0x00300183 | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x00300274 | `.net-codec in dit product wordt gebruikt in licentie va...` | Known | Audio system |
| 0x00300310 | `Technologie voor codering van MPEG Layer-3-audio in lic...` | Known | Audio system |
| 0x0030645C | `Audible-programvaren i dette produktet brukes under lis...` | Known | Audible audiobook format |
| 0x003064B0 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x0030662C | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fr...` | Known | Audio system |
| 0x0030C864 | `Oprogramowanie Audible w tym produkcie jest wykorzystyw...` | Known | Audible audiobook format |
| 0x0030C8D0 | ` 2002 Audible, Inc. Wszystkie prawa zastrze` | Known | Audible audiobook format |
| 0x0030CA78 | `Technologia kodowania audio MPEG Layer-3 licencjonowana...` | Known | Audio system |
| 0x00312E4C | `O software Audible ` | Known | Audible audiobook format |
| 0x00312E82 | `a da Audible. Copyright ` | Known | Audible audiobook format |
| 0x00312E9C | ` 2002 da Audible, Inc. Reservados todos os direitos.` | Known | Audible audiobook format |
| 0x00312F5D | `a da VoiceAge Corporation. O codec ACELP` | Known | Audio system |
| 0x0031303E | `udio MPEG Layer-3 licenciada pela Fraunhofer IIS e THOM...` | Known | Audio system |
| 0x0031B134 | `MPEG Layer-3: ` | Known | Audio system |
| 0x003221AC | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x003221DB | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x003221F2 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x0032238C | `Ljudkodningstekniken MPEG Layer-3 ` | Known | Audio system |
| 0x003223C2 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | Audio system |
| 0x003285C0 | `ndeki Audible yaz` | Known | Audible audiobook format |
| 0x003285D9 | ` Audible lisans` | Known | Audible audiobook format |
| 0x0032860E | ` 2002, Audible, Inc. T` | Known | Audible audiobook format |
| 0x00328709 | `.net codec'i VoiceAge Corporation lisans` | Known | Audio system |
| 0x00328798 | `MPEG Layer-3 ses kodlama teknolojisi Fraunhofer IIS ve ...` | Known | Audio system |
| 0x004B777C | `The Audible software in this product is used under lice...` | Known | Audible audiobook format |
| 0x004B78B5 | `.net codec in this product is used under license from V...` | Known | Audio system |
| 0x004B7948 | `MPEG Layer-3 audio coding technology licensed from Frau...` | Known | Audio system |
| 0x004B7E34 | `TV Out` | Known | Audio system |
| 0x0065FAD1 | ``0aLaCfDf` | Known | Apple Lossless codec |
| 0x0066A814 | `21SoundEffectDescriptor` | Known | Audio system |
| 0x0067DC74 | `&Aacute` | Known | AAC codec |
| 0x0067DC9C | `&aacute` | Known | AAC codec |
| 0x0068048E | `msCodeCom` | Known | Audio system |
| 0x00680FE3 | `aaControls` | Known | AAC codec |
| 0x00B2864F | `gencmd_decode_fourcc` | Known | Decoder component |
| 0x00B28664 | `gencmd_decode_int` | Known | Decoder component |
| 0x00B39860 | `AACDEC  VLL ` | Known | AAC codec |
| 0x00B42FA8 | `MPEG4 AAC LC Decoder` | Known | AAC codec |
| 0x00B4540D | `AACDecoderGetMem` | Known | AAC codec |
| 0x00B4541E | `AACDecoderInit` | Known | AAC codec |
| 0x00B4542D | `AACDecoderGetConfig` | Known | AAC codec |
| 0x00B45441 | `AACDecoderSetConfig` | Known | AAC codec |
| 0x00B45455 | `AACHeaderDecode` | Known | AAC codec |
| 0x00B45465 | `AACDecode` | Known | AAC codec |
| 0x00B4546F | `AACDecoderInit_Ittiam` | Known | AAC codec |
| 0x00B45485 | `AACDecoderGetConfig_Ittiam` | Known | AAC codec |
| 0x00B454A0 | `AACDecoderSetConfig_Ittiam` | Known | AAC codec |
| 0x00B454BB | `AACHeaderDecode_Ittiam` | Known | AAC codec |
| 0x00B454D2 | `AACDecode_Ittiam` | Known | AAC codec |
| 0x00B454FA | `get_aac_dec_func_table` | Known | AAC codec |
| 0x00B45592 | `aac_initbits` | Known | AAC codec |
| 0x00B4559F | `aac_get_processed_bits` | Known | AAC codec |
| 0x00B455B6 | `aac_byte_align` | Known | AAC codec |
| 0x00B457A6 | `is_decode` | Known | Decoder component |
| 0x00B457B0 | `can_decode_objType` | Known | Decoder component |
| 0x00B457D7 | `ms_decode` | Known | Decoder component |
| 0x00B4581E | `pns_decode` | Known | Decoder component |
| 0x00B45829 | `pulse_decode` | Known | Decoder component |
| 0x00B45856 | `tns_decode_frame` | Known | Decoder component |
| 0x00B53EE0 | `H.264 Video Decoder` | Known | Decoder component |
| 0x00B5D000 | `H264InitDecoder` | Known | Decoder component |
| 0x00B5D010 | `init_decoder` | Known | Decoder component |
| 0x00B5D01D | `H264DecodeFrame` | Known | Decoder component |
| 0x00B5D02D | `H264ReleaseDecoder` | Known | Decoder component |
| 0x00B5D040 | `shutdown_decoder` | Known | Decoder component |
| 0x00B5D1A8 | `decode_one_frame` | Known | Decoder component |
| 0x00B5D52D | `h264_refstripe_prepare_decode` | Known | Decoder component |
| 0x00B5D54B | `h264_refstripe_finished_decode` | Known | Decoder component |
| 0x00B5DCFF | `mvpairdecode_table` | Known | Decoder component |
| 0x00B5DD12 | `mvpairdecodelen_table` | Known | Decoder component |
| 0x00B5E6C1 | `h264_writestripe_prepare_decode` | Known | Decoder component |
| 0x00B790A0 | `MPEG-4 video decoder` | Known | Decoder component |
| 0x00B8146A | `UBVInitDecoder` | Known | Decoder component |
| 0x00B81479 | `UBVDecodeFrame` | Known | Decoder component |
| 0x00B81488 | `UBVReleaseDecoder` | Known | Decoder component |
| 0x00B81593 | `macroblockdecode` | Known | Decoder component |
| 0x00B81610 | `vc_ClipIquantMPEG4` | Known | Audio system |
| 0x00B81623 | `vc_MPEG4InterIQuant` | Known | Audio system |
| 0x00B81864 | `vc_MPEG4getDC` | Known | Audio system |
| 0x00B8188F | `MPEG4CopyMBs` | Known | Audio system |
| 0x00B8189C | `DecodeFrame` | Known | Decoder component |
| 0x00B818A8 | `MPEG4DecodeFrame` | Known | Decoder component |
| 0x00B818B9 | `H263DecodeFrame` | Known | Decoder component |
| 0x00B818C9 | `MPEG4DecodeEnhancementFrame` | Known | Decoder component |
| 0x00B818E5 | `MPEG4InitParams` | Known | Audio system |
| 0x00B81EF5 | `H263DecodeMotionVectors` | Known | Decoder component |
| 0x00B81F0D | `H263DecodeIPic` | Known | Decoder component |
| 0x00B81F46 | `H263DecodePPic` | Known | Decoder component |
| 0x00B82007 | `pmacroblockdecode` | Known | Decoder component |
| 0x00B82113 | `MPEG4PPictureSave6BlocksD` | Known | Audio system |
| 0x00B82191 | `MPEG4GetCBPY` | Known | Audio system |
| 0x00B8219E | `MPEG4MCBPCIntra` | Known | Audio system |
| 0x00B821AE | `MPEG4MCBPCInter` | Known | Audio system |
| 0x00B821BE | `MPEG4InitMVD` | Known | Audio system |
| 0x00B821CB | `MPEG4GetTMNMV` | Known | Audio system |
| 0x00B821D9 | `MPEG4AddPMV` | Known | Audio system |
| 0x00B821E5 | `MPEG4GetMotionVectorsF` | Known | Audio system |
| 0x00B821FC | `MPEG4GetMotionVectorData` | Known | Audio system |
| 0x00B82215 | `MPEG4FindPos4MVD` | Known | Audio system |
| 0x00B82226 | `MPEG4ClipMV` | Known | Audio system |
| 0x00B82232 | `MPEG4PredictIntraBlock` | Known | Audio system |
| 0x00B82249 | `MPEG4InitRowBlocksD` | Known | Audio system |
| 0x00B8225D | `MPEG4IntraAdvancedPredictionDecodeD` | Known | Decoder component |
| 0x00B82281 | `MPEG4DecodeIPic` | Known | Decoder component |
| 0x00B82291 | `MPEG4DecodeDataPartitionedIPic` | Known | Decoder component |
| 0x00B823AD | `MPEG4_INTRA_MP4BDEC_UBV_DCT3D0` | Known | Audio system |
| 0x00B823CC | `MPEG4_INTRA_MP4BDEC_UBV_DCT3D1` | Known | Audio system |
| 0x00B823EB | `MPEG4_INTRA_MP4BDEC_UBV_DCT3D2` | Known | Audio system |
| 0x00B8241E | `MPEG4GetInterBlock` | Known | Audio system |
| 0x00B82445 | `MPEG4GetIntraBlock` | Known | Audio system |
| 0x00B82458 | `MPEG4RvlcDecTCOEF` | Known | Audio system |
| 0x00B8246A | `MPEG4GetIntraBlockRVLC` | Known | Audio system |
| 0x00B82481 | `MPEG4GetInterBlockRVLC` | Known | Audio system |
| 0x00B82498 | `MPEG4ParseVOLHeader` | Known | Audio system |
| 0x00B824AC | `MPEG4FlushUserData` | Known | Audio system |
| 0x00B824BF | `MPEG4FindStartCode` | Known | Audio system |
| 0x00B824D2 | `MPEG4ParseVOPHeader` | Known | Audio system |
| 0x00B824E6 | `MPEG4CheckMotionMarker` | Known | Audio system |
| 0x00B824FD | `MPEG4DecodePPic` | Known | Decoder component |
| 0x00B8250D | `MPEG4DecodeInterTextureMacroblock` | Known | Decoder component |
| 0x00B8252F | `MPEG4DecodeDataPartitionedPPic` | Known | Decoder component |
| 0x00B8254E | `MPEG4ReadIMacroBlocks` | Known | Audio system |
| 0x00B82564 | `MPEG4ReadDQuant` | Known | Audio system |
| 0x00B82574 | `MPEG4ReadPMacroBlocks` | Known | Audio system |
| 0x00B8258A | `MPEG4DataPartitionReadIMacroBlocks` | Known | Audio system |
| 0x00B825AD | `MPEG4DataPartitionReadPMacroBlocks` | Known | Audio system |
| 0x00B825D0 | `MPEG4DataPartitionReadPMacroBlocks2` | Known | Audio system |
| 0x00B825F4 | `MPEG4ReadVideoPacketHeader` | Known | Audio system |
| 0x00B82616 | `InitDecoder` | Known | Decoder component |
| 0x00B82631 | `MPEG4InitMemory` | Known | Audio system |
| 0x00B82641 | `InitEnhancementDecoder` | Known | Decoder component |
| 0x00B82658 | `ReleaseDecoder` | Known | Decoder component |
| 0x00B82667 | `H263ReleaseH263Decoder` | Known | Decoder component |
| 0x00B8267E | `MPEG4ReleaseMPEG4Decoder` | Known | Decoder component |
| 0x00B82697 | `ReleaseEnhancementDecoder` | Known | Decoder component |
| 0x00B82B9A | `MPEG4ShowBitsAlignedD` | Known | Audio system |
| 0x00B82BB0 | `MPEG4AlignInput` | Known | Audio system |
| 0x00B82BC0 | `MPEG4PeekNextStartCode` | Known | Audio system |
| 0x00B82C5E | `Decoders` | Known | Decoder component |
| 0x00B86810 | `forbid_decoder_panic=` | Known | Decoder component |
| 0x00B86D65 | `max_decoded_bytes=%d` | Known | Decoder component |
| 0x00B86D95 | `max_decoded_buffer=%u` | Known | Decoder component |
| 0x00B86DAD | `min_decoded_buffer=%u` | Known | Decoder component |
| 0x00B86DC5 | `total_bytes_decoded=%u` | Known | Decoder component |
| 0x00B86DDD | `Disp decoded_data=%08x` | Known | Decoder component |
| 0x00B8CE8F | `no audio decoder` | Known | Decoder component |
| 0x00B8CEA0 | `no video decoder` | Known | Decoder component |
| 0x00B8FB9D | `codec_string` | Known | Audio system |
| 0x00B8FBAA | `codec_name` | Known | Audio system |
| 0x00B9E7F1 | `decode_int` | Known | Decoder component |

---

## 6. Storage/Hardware

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00003A00 | `!ATAsoso` | Known | ATA/disk interface |
| 0x00003A28 | `!ATAcrsr` | Known | ATA/disk interface |
| 0x00003A50 | `!ATAdpua` | Known | ATA/disk interface |
| 0x00004D34 | `!ATAebih` | Known | ATA/disk interface |
| 0x00005F89 | `diskmode` | Known | Hardware interface |
| 0x00005F92 | `diskscan` | Known | Hardware interface |
| 0x0007E118 | `Metadata` | Known | ATA/disk interface |
| 0x00098974 | `Photo Database` | Known | ATA/disk interface |
| 0x000AA08C | `atadmrts` | Known | ATA/disk interface |
| 0x000B454C | `atadmhbddbhmmhsd` | Known | ATA/disk interface |
| 0x000B4828 | `atadmhfddfhmmhsd\|@-` | Known | ATA/disk interface |
| 0x000B9604 | `Photos\Photo Database` | Known | ATA/disk interface |
| 0x000C088C | `nutiatad` | Known | ATA/disk interface |
| 0x000C0B54 | `atadImage DB Temp` | Known | ATA/disk interface |
| 0x000CAAAC | `]ih[!ATA` | Known | ATA/disk interface |
| 0x000EA68C | `atadmhpo0@-` | Known | ATA/disk interface |
| 0x000EE4E4 | `data abort` | Known | ATA/disk interface |
| 0x000F06E8 | `atadmhdp` | Known | ATA/disk interface |
| 0x00142EEC | `Bad meta data, name not found. (23)` | Known | ATA/disk interface |
| 0x00142F14 | `Bad meta data, name termination quote not found. (24)` | Known | ATA/disk interface |
| 0x00142F50 | `Bad meta data, content not found. (25)` | Known | ATA/disk interface |
| 0x00142F78 | `Bad meta data, content termination quote not found. (26...` | Known | ATA/disk interface |
| 0x00143C1C | `Bad meta data, encoding not found. (8)` | Known | ATA/disk interface |
| 0x00143C44 | `Bad meta data, encoding termination quote not found. (9...` | Known | ATA/disk interface |
| 0x00143C80 | `Bad meta data, unknown encoding. (17)` | Known | ATA/disk interface |
| 0x00144024 | `Bad meta data, name not a recognized preference. (27)` | Known | ATA/disk interface |
| 0x00144730 | `Bad data. (32)` | Known | ATA/disk interface |
| 0x00177A6D | `lyrdata` | Known | ATA/disk interface |
| 0x0017A8E4 | `MaximumDataRate` | Known | ATA/disk interface |
| 0x0017AD94 | `FireWireGUID` | Known | FireWire interface |
| 0x0017ADA4 | `FireWireVersion` | Known | FireWire interface |
| 0x0017B390 | `FireWire` | Known | FireWire interface |
| 0x0017B838 | `ForcedDiskMode` | Known | Hardware interface |
| 0x0017B858 | `CorruptDataPartition` | Known | ATA/disk interface |
| 0x00189024 | `USB MSC` | Known | USB interface |
| 0x0019BABC | `23iUPhoto Database` | Known | ATA/disk interface |
| 0x0019E124 | `Photo Import Database` | Known | ATA/disk interface |
| 0x001CC538 | `spiral` | Known | Hardware interface |
| 0x0021676C | `Channel PlayFromDisk` | Known | Hardware interface |
| 0x00216784 | `Channel CacheSpinupDrive` | Known | Hardware interface |
| 0x002168F0 | `Channel DiskModeChannel` | Known | Hardware interface |
| 0x00216908 | `Channel FirewireChannel` | Known | FireWire interface |
| 0x00216A28 | `Channel DiskImage` | Known | Hardware interface |
| 0x00216A6C | `Channel DiskFormatConvert` | Known | Hardware interface |
| 0x00216C10 | `Channel PredictiveCacheSpinup` | Known | Hardware interface |
| 0x00216C38 | `Unknown Disk Channel` | Known | Hardware interface |
| 0x00217368 | `Disk Activity` | Known | Hardware interface |
| 0x00217379 | `Total time the disk was running in the app: %d seconds` | Known | Hardware interface |
| 0x00217421 | `The disk was turned on %d %s` | Known | Hardware interface |
| 0x00217BD1 | `Music database size: %d KB` | Known | ATA/disk interface |
| 0x00217BF1 | `Music database num songs: %d` | Known | ATA/disk interface |
| 0x00217C11 | `Photo database size: %d KB` | Known | ATA/disk interface |
| 0x00217C31 | `Photo database num photos: %d` | Known | ATA/disk interface |
| 0x00217C51 | `Album art database size: %d KB` | Known | ATA/disk interface |
| 0x002185CC | `Disk Spinup` | Known | Hardware interface |
| 0x002185D8 | `Disk Spindown` | Known | Hardware interface |
| 0x002185E8 | `Disk Obtain Access` | Known | Hardware interface |
| 0x002185FC | `Disk Release Access` | Known | Hardware interface |
| 0x002186B8 | `Flush Usage Log Data` | Known | ATA/disk interface |
| 0x00218740 | `Enter Disk Mode` | Known | Hardware interface |
| 0x00218750 | `Exit Disk Mode` | Known | Hardware interface |
| 0x002187CC | `Music Database Size` | Known | ATA/disk interface |
| 0x002187E0 | `Photo Database Size` | Known | ATA/disk interface |
| 0x002187F4 | `Artwork Database Size` | Known | ATA/disk interface |
| 0x00262610 | `[CDATA[` | Known | ATA/disk interface |
| 0x0026C7E4 | `MEMDISK` | Known | Hardware interface |
| 0x002ADC4B | `glBufferData` | Known | ATA/disk interface |
| 0x002ADC58 | `glBufferSubData` | Known | ATA/disk interface |
| 0x002AE448 | `gamedata_RW` | Known | ATA/disk interface |
| 0x002AE464 | `gamedata_ShareRW` | Known | ATA/disk interface |
| 0x002B2460 | `e v diskov` | Known | Hardware interface |
| 0x002B2B20 | `Data RDS nenalezena` | Known | ATA/disk interface |
| 0x002B3778 | `Kalkata` | Known | ATA/disk interface |
| 0x002B3B1C | `im disku` | Known | Hardware interface |
| 0x002B3B37 | `es FireWire nen` | Known | FireWire interface |
| 0x002B3CDB | `dat a zobrazovat data importovan` | Known | ATA/disk interface |
| 0x002B5859 | `e na disku nen` | Known | Hardware interface |
| 0x002B5AB0 | `FireWire p` | Known | FireWire interface |
| 0x002B8A11 | ` brug af iPod som ekstern disk til og tr` | Known | Hardware interface |
| 0x002B8D58 | `Videospillelister` | Known | Hardware interface |
| 0x002B9080 | `Ingen RDS-data fundet` | Known | ATA/disk interface |
| 0x002B90A0 | ` Afspil for at lytte til radio` | Known | Hardware interface |
| 0x002B90C8 | ` Afspil for at slukke radioen` | Known | Hardware interface |
| 0x002B9220 | `Spiller nu` | Known | Hardware interface |
| 0x002B9304 | `Spillelister` | Known | Hardware interface |
| 0x002B9330 | `Genoptag spil` | Known | Hardware interface |
| 0x002B939C | `Ved afspilning` | Known | Hardware interface |
| 0x002B9C60 | `Kolkata (Calcutta)` | Known | ATA/disk interface |
| 0x002B9DA4 | `Ulaanbaatar` | Known | ATA/disk interface |
| 0x002B9E94 | `Slet spilleliste` | Known | Hardware interface |
| 0x002B9EA8 | `Arkiver spilleliste` | Known | Hardware interface |
| 0x002B9F68 | `Ny spilleliste %lu` | Known | Hardware interface |
| 0x002B9FF4 | `Harddisk` | Known | Hardware interface |
| 0x002BA000 | `FireWire-forbindelser underst` | Known | FireWire interface |
| 0x002BA064 | `re sange og data.` | Known | ATA/disk interface |
| 0x002BA084 | `Afspil %s` | Known | Hardware interface |
| 0x002BA524 | `Slut iPod til iTunes, og installer spillet igen.` | Known | Hardware interface |
| 0x002BA558 | `Spillet kan ikke spilles.` | Known | Hardware interface |
| 0x002BA5F4 | `Denne version af spillet underst` | Known | Hardware interface |
| 0x002BA78F | ` afspilningsknappen p` | Known | Hardware interface |
| 0x002BAE2F | `je den til spillelisten On-The-Go. Hold knappen nede, n` | Known | Hardware interface |
| 0x002BAE68 | `r en spilleliste, kunstner eller et album er valgt for ...` | Known | Hardware interface |
| 0x002BAEA5 | `je alle sangene til spillelisten On-The-Go.` | Known | Hardware interface |
| 0x002BB280 | `Nyt spil` | Known | Hardware interface |
| 0x002BB294 | `Afspil` | Known | Hardware interface |
| 0x002BB6AC | `Afspil video` | Known | Hardware interface |
| 0x002BB728 | `Dette mediearkiv kan ikke vises eller afspilles p` | Known | Hardware interface |
| 0x002BBD20 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x002BFB38 | `Spiele` | Known | Hardware interface |
| 0x002BFBB8 | `Weiterspielen` | Known | Hardware interface |
| 0x002C04F8 | `Kolkata (Kalkutta)` | Known | ATA/disk interface |
| 0x002C08A8 | `FireWire wird nicht unterst` | Known | FireWire interface |
| 0x002C0980 | `Spitzname` | Known | Hardware interface |
| 0x002C0D24 | `Beispiel` | Known | Hardware interface |
| 0x002C0D3C | `Beispielfirma GmbH` | Known | Hardware interface |
| 0x002C0D50 | `Dieses Beispiel zeigt, welche Infos Sie bei einem Konta...` | Known | Hardware interface |
| 0x002C0F90 | `Verbinden Sie Ihren iPod mit iTunes und installieren Si...` | Known | Hardware interface |
| 0x002C0FDC | `Dieses Spiel kann nicht gespielt werden.` | Known | Hardware interface |
| 0x002C109C | `Diese Version des Spiels wird nicht mehr unterst` | Known | Hardware interface |
| 0x002C1DD0 | `Neues Spiel` | Known | Hardware interface |
| 0x002C22A8 | `Die Mediendatei kann nicht auf dem iPod angezeigt oder ...` | Known | Hardware interface |
| 0x002C292E | `ber FireWire verbunden` | Known | FireWire interface |
| 0x002C84C6 | ` FireWire. ` | Known | FireWire interface |
| 0x002CBA66 | ` FireWire` | Known | FireWire interface |
| 0x002CFE14 | `Kolkata (Calcuta)` | Known | ATA/disk interface |
| 0x002D01B9 | `de canciones o archivos no son posibles con FireWire :` | Known | FireWire interface |
| 0x002D13F4 | `Espiral` | Known | Hardware interface |
| 0x002D20E8 | `FireWire conectado` | Known | FireWire interface |
| 0x002D5590 | `Etsi kanavia -komento etsii kaikki saatavilla olevat ra...` | Known | ATA/disk interface |
| 0x002D5704 | `RDS-dataa ei havaittu` | Known | ATA/disk interface |
| 0x002D5C50 | `Diskanttivahv.` | Known | Hardware interface |
| 0x002D5C60 | `Diskanttiheik.` | Known | Hardware interface |
| 0x002D5C8C | `Diskantinkorostus` | Known | Hardware interface |
| 0x002D6644 | `Ladataan` | Known | ATA/disk interface |
| 0x002D6718 | `FireWire-tiedonsiirtoa ei tueta. Siirt` | Known | FireWire interface |
| 0x002D67B4 | `Ladataan...` | Known | ATA/disk interface |
| 0x002D6898 | `yhteystietoa ladataan.` | Known | ATA/disk interface |
| 0x002D6BAC | `Yhteystietoa ladataan.` | Known | ATA/disk interface |
| 0x002D6C6E | ` ei voi pelata.` | Known | ATA/disk interface |
| 0x002D6FD8 | `Muistiinpanoja ladataan.` | Known | ATA/disk interface |
| 0x002D8188 | `nityksen jatkamiseen ei ole tarpeeksi vapaata levytilaa...` | Known | ATA/disk interface |
| 0x002D81C8 | `nityksen aloittamiseen ei ole tarpeeksi vapaata levytil...` | Known | ATA/disk interface |
| 0x002D8418 | `FireWire liitetty` | Known | FireWire interface |
| 0x002DD6E5 | `s via FireWire : connectez l` | Known | FireWire interface |
| 0x002DF908 | `FireWire Connect` | Known | FireWire interface |
| 0x002E4114 | `A FireWire kapcsolat nem t` | Known | FireWire interface |
| 0x002E6324 | `FireWire csatlakozik` | Known | FireWire interface |
| 0x002E9570 | `Durata diapositiva` | Known | ATA/disk interface |
| 0x002E98D8 | ` stata effettuata. Premi e mantieni premuto il pulsante...` | Known | ATA/disk interface |
| 0x002EA98C | `Connessioni di dati via FireWire non sono supportate. P...` | Known | USB interface |
| 0x002EAA90 | `auto privata` | Known | ATA/disk interface |
| 0x002EB00D | ` supportata.` | Known | ATA/disk interface |
| 0x002EB9F0 | `Data & ora` | Known | ATA/disk interface |
| 0x002EBAEC | `Spazzata dal centro` | Known | ATA/disk interface |
| 0x002EBB00 | `Spazzata Verso il basso` | Known | ATA/disk interface |
| 0x002EBB18 | `Spazzata di lato` | Known | ATA/disk interface |
| 0x002EBB2C | `Spirale` | Known | Hardware interface |
| 0x002EBB3C | `Spinta verso il basso` | Known | Hardware interface |
| 0x002EBB54 | `Spinta di lato` | Known | Hardware interface |
| 0x002EBD08 | `Imposta data & ora` | Known | ATA/disk interface |
| 0x002EC718 | `Data & Ora` | Known | ATA/disk interface |
| 0x002EC7DC | `FireWire connesso` | Known | FireWire interface |
| 0x002F17AC | `FireWire ` | Known | FireWire interface |
| 0x002FE35C | `Handmatig` | Known | Hardware interface |
| 0x002FF520 | `Jekatarinenburg` | Known | ATA/disk interface |
| 0x002FF77A | `ren via FireWire, maar alleen via de meegeleverde USB-k...` | Known | USB interface |
| 0x00301588 | `FireWire aangesloten` | Known | FireWire interface |
| 0x003044AC | `Hvis du vil vise tekstfiler her, aktiverer du iPod for ...` | Known | Hardware interface |
| 0x00304B54 | `Finner ikke RDS-data` | Known | ATA/disk interface |
| 0x00304CE0 | `Spilles n` | Known | Hardware interface |
| 0x00304DEC | `Fortsett spill` | Known | Hardware interface |
| 0x00304E58 | `Under avspilling` | Known | Hardware interface |
| 0x0030501C | `Diskantforsterkning` | Known | Hardware interface |
| 0x00305030 | `Diskantreduksjon` | Known | Hardware interface |
| 0x0030595C | `Slett spilleliste` | Known | Hardware interface |
| 0x00305ABC | `Diskmodus` | Known | Hardware interface |
| 0x00305ACF | `ring via FireWire st` | Known | FireWire interface |
| 0x00305B23 | `re sanger eller data.` | Known | ATA/disk interface |
| 0x00305B48 | `Spill %s` | Known | Hardware interface |
| 0x00305FF0 | `Koble iPod til iTunes, og installer spillet p` | Known | Hardware interface |
| 0x00306028 | `Dette spillet kan ikke spilles.` | Known | Hardware interface |
| 0x003060C0 | `Denne versjonen av dette spillet st` | Known | Hardware interface |
| 0x003060F8 | `Hvis du glemmer kombinasjonen, kan du koble iPod til da...` | Known | ATA/disk interface |
| 0x003061D0 | `r bildene til datamaskinen, og synkroniser dem via iTun...` | Known | ATA/disk interface |
| 0x0030625E | ` avspillingsknappen p` | Known | Hardware interface |
| 0x00306892 | ` legge den til i On-The-Go-spillelisten. Spillelister, ...` | Known | Hardware interface |
| 0x003069B7 | ` denne iPod-enheten. Koble iPod til datamaskinen, og st...` | Known | ATA/disk interface |
| 0x00306CE0 | `Nytt spill` | Known | Hardware interface |
| 0x0030717C | `Denne mediefilen kan ikke vises eller spilles p` | Known | Hardware interface |
| 0x003071C3 | ` datamaskinen ved hjelp av QuickTime.` | Known | ATA/disk interface |
| 0x00307224 | `r importerte bilder til datamaskinen, og synkroniser vi...` | Known | ATA/disk interface |
| 0x00307508 | `Det er ikke nok ledig diskplass til ` | Known | Hardware interface |
| 0x00307770 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x0030A998 | `Strata` | Known | ATA/disk interface |
| 0x0030BE7F | `czenie FireWire nie jest wspierane. By przes` | Known | FireWire interface |
| 0x0030CF3C | `Data i czas` | Known | ATA/disk interface |
| 0x0030DDCF | `czony przez Firewire` | Known | FireWire interface |
| 0x00312100 | `Kolkata (Calcut` | Known | ATA/disk interface |
| 0x003124CF | `es FireWire n` | Known | FireWire interface |
| 0x00313508 | `Data & hora` | Known | ATA/disk interface |
| 0x00313884 | `Definir data & hora` | Known | ATA/disk interface |
| 0x003143D0 | `FireWire ligado` | Known | FireWire interface |
| 0x00319D95 | ` FireWire ` | Known | FireWire interface |
| 0x00320370 | `inget kort inmatat` | Known | ATA/disk interface |
| 0x003208B0 | `Inga RDS-data kan hittas` | Known | ATA/disk interface |
| 0x00321824 | `FireWire-` | Known | FireWire interface |
| 0x00321874 | `ver musik eller data.` | Known | ATA/disk interface |
| 0x003225D8 | `Stort bildmaterial` | Known | Hardware interface |
| 0x00323540 | `FireWire anslutet` | Known | FireWire interface |
| 0x003264AE | `in iPod'u disk kullan` | Known | Hardware interface |
| 0x00327B70 | `Disk Modu` | Known | Hardware interface |
| 0x00327B7C | `FireWire ba` | Known | FireWire interface |
| 0x00328448 | `nda bir hata olu` | Known | ATA/disk interface |
| 0x0032931C | `Bilinmeyen Hata` | Known | ATA/disk interface |
| 0x00329821 | ` disk alan` | Known | Hardware interface |
| 0x00329A88 | `FireWire Ba` | Known | FireWire interface |
| 0x004B52E8 | `To view text files here, enable iPod for disk use, then...` | Known | Hardware interface |
| 0x004B5974 | `No RDS Data Detected` | Known | ATA/disk interface |
| 0x004B6DFC | `Disk Mode` | Known | Hardware interface |
| 0x004B6E08 | `FireWire connections are not supported. To transfer son...` | Known | USB interface |
| 0x004B7C3C | `The battery level is too low.` | Known | Power management |
| 0x004B7FE0 | `Disk Browser` | Known | Hardware interface |
| 0x004B8814 | `There is not enough free disk space to continue recordi...` | Known | Hardware interface |
| 0x004B8850 | `There is not enough free disk space to start recording.` | Known | Hardware interface |
| 0x004B8AB0 | `FireWire Connected` | Known | FireWire interface |
| 0x004B8AC4 | `No battery power remains. Please connect iPod to power.` | Known | Power management |
| 0x004B8AFC | `Low Battery` | Known | Power management |
| 0x00500850 | `I2C write Error` | Known | Hardware interface |
| 0x00500864 | `I2C read Error %02x` | Known | Hardware interface |
| 0x00643448 | `TROMResourceDB - unknown header version! (Try regenerat...` | Known | ATA/disk interface |
| 0x006548F3 | `ataTaza[aea;ajaaaVa)b'b+b+dMd[d]dtdvdrdsd}dudfd` | Known | ATA/disk interface |
| 0x006594DD | `aGa>a(a'aJa?a<a,a4a=aBaDasawaXaYaZakataoaeaqa_a]aSaua` | Known | ATA/disk interface |
| 0x0066A7A9 | `15TCountedPointerI10SImageDataE` | Known | ATA/disk interface |
| 0x0066A7DB | `15iMAXMLParseData` | Known | ATA/disk interface |
| 0x0066A8E5 | `N4eApp17ManifestDataProxyE` | Known | ATA/disk interface |
| 0x0067D05F | `HoldSwitch` | Known | Hardware interface |
| 0x0067DBD9 | `Bad Data` | Known | ATA/disk interface |
| 0x0067E3CB | `ex_data` | Known | ATA/disk interface |
| 0x0067E4C4 | `RSA Data Security, Inc.` | Known | ATA/disk interface |
| 0x0067E97E | `set-brand-IATA-ATA` | Known | ATA/disk interface |
| 0x0067F0D3 | `RSA Data Security, Inc. PKCS` | Known | ATA/disk interface |
| 0x0067F2DB | `setCext-Track2Data` | Known | ATA/disk interface |
| 0x0067F2EE | `id-cct-PKIData` | Known | ATA/disk interface |
| 0x0067F2FD | `setct-OIData` | Known | ATA/disk interface |
| 0x0067F30A | `setct-PIData` | Known | ATA/disk interface |
| 0x0067F317 | `setct-PANData` | Known | ATA/disk interface |
| 0x0067F325 | `qualityLabelledData` | Known | ATA/disk interface |
| 0x0067F339 | `pkcs7-signedData` | Known | ATA/disk interface |
| 0x0067F34A | `pkcs7-signedAndEnvelopedData` | Known | ATA/disk interface |
| 0x0067F367 | `pkcs7-envelopedData` | Known | ATA/disk interface |
| 0x0067F37B | `pkcs7-encryptedData` | Known | ATA/disk interface |
| 0x0067F38F | `id-smime-ct-DVCSResponseData` | Known | ATA/disk interface |
| 0x0067F3AC | `setCext-merchData` | Known | ATA/disk interface |
| 0x0067F3BE | `id-smime-ct-authData` | Known | ATA/disk interface |
| 0x0067F3D3 | `id-on-personalData` | Known | ATA/disk interface |
| 0x0067F3E6 | `setct-CapTokenData` | Known | ATA/disk interface |
| 0x0067F3F9 | `setct-BatchAdminReqData` | Known | ATA/disk interface |
| 0x0067F411 | `setct-CertReqData` | Known | ATA/disk interface |
| 0x0067F423 | `setct-PCertReqData` | Known | ATA/disk interface |
| 0x0067F436 | `setct-PResData` | Known | ATA/disk interface |
| 0x0067F445 | `setct-CredResData` | Known | ATA/disk interface |
| 0x0067F457 | `setct-BatchAdminResData` | Known | ATA/disk interface |
| 0x0067F46F | `setct-CapResData` | Known | ATA/disk interface |
| 0x0067F480 | `setct-PInitResData` | Known | ATA/disk interface |
| 0x0067F493 | `setct-CertResData` | Known | ATA/disk interface |
| 0x0067F4A5 | `setct-CredRevResData` | Known | ATA/disk interface |
| 0x0067F4BA | `setct-AuthRevResData` | Known | ATA/disk interface |
| 0x0067F4CF | `setct-CapRevResData` | Known | ATA/disk interface |
| 0x0067F4E3 | `pkcs7-digestData` | Known | ATA/disk interface |
| 0x0067F4F4 | `id-smime-ct-DVCSRequestData` | Known | ATA/disk interface |
| 0x0067F510 | `pkcs7-data` | Known | ATA/disk interface |
| 0x0067F76D | `setct-PIDataUnsigned` | Known | ATA/disk interface |
| 0x0067FD4F | `Netscape Data Type` | Known | ATA/disk interface |
| 0x0067FD75 | `nsDataType` | Known | ATA/disk interface |
| 0x006807BE | `id-cmc-dataReturn` | Known | ATA/disk interface |
| 0x0068191A | `d.data` | Known | ATA/disk interface |
| 0x006819C0 | `enc_data` | Known | ATA/disk interface |
| 0x0068204D | `Data Encipherment` | Known | ATA/disk interface |
| 0x00682070 | `dataEncipherment` | Known | ATA/disk interface |
| 0x00682350 | `OCSP_RESPDATA` | Known | ATA/disk interface |
| 0x00682379 | `OCSP_RESPID` | Known | Hardware interface |
| 0x006825C5 | `tbsResponseData` | Known | ATA/disk interface |
| 0x0071CE67 | `@!ATAp@-` | Known | ATA/disk interface |
| 0x00AEF3E8 | `gldMallocSlow` | Known | Hardware interface |
| 0x00AF0392 | `0BgldMallocSlow` | Known | Hardware interface |
| 0x00AF3718 | `Length is less than data described in texture sub data ...` | Known | ATA/disk interface |
| 0x00B13E64 | `disk_notify` | Known | Hardware interface |
| 0x00B274FC | `TV_DMA_INIT` | Known | Hardware interface |
| 0x00B27508 | `TV_DMA_START` | Known | Hardware interface |
| 0x00B27515 | `TV_DMA_MIDDLE` | Known | Hardware interface |
| 0x00B27523 | `TV_DMA_END` | Known | Hardware interface |
| 0x00B2752E | `TV_DMA_BLOCK` | Known | Hardware interface |
| 0x00B2753B | `TV_DMA_STOP` | Known | Hardware interface |
| 0x00B28451 | `dma_get_transfer_queue` | Known | Hardware interface |
| 0x00B28468 | `dma_memcpy` | Known | Hardware interface |
| 0x00B28473 | `dma_memcpy2d_uncached` | Known | Hardware interface |
| 0x00B28489 | `dma_subchan_free` | Known | Hardware interface |
| 0x00B2849A | `dma_subchan_request` | Known | Hardware interface |
| 0x00B284AE | `dma_transfer_chain` | Known | Hardware interface |
| 0x00B284C1 | `dma_transfer_has_finished` | Known | Hardware interface |
| 0x00B284DB | `dma_transfer_queue_post` | Known | Hardware interface |
| 0x00B284F3 | `dma_transfer_queue_release` | Known | Hardware interface |
| 0x00B2850E | `dma_transfer_set_callback` | Known | Hardware interface |
| 0x00B28528 | `dma_transfer_setup_memcpy` | Known | Hardware interface |
| 0x00B28542 | `dma_transfer_setup_memcpy_uncached` | Known | Hardware interface |
| 0x00B28565 | `dma_transfer_setup_memcpy2d_uncached` | Known | Hardware interface |
| 0x00B2858A | `dma_transfer_wait` | Known | Hardware interface |
| 0x00B28B7F | `vc_image_set_image_data` | Known | ATA/disk interface |
| 0x00B28B97 | `vc_image_set_image_data_yuv` | Known | ATA/disk interface |
| 0x00B40EC2 | `X(4ATA` | Known | ATA/disk interface |
| 0x00B45547 | `pulse_data` | Known | ATA/disk interface |
| 0x00B45552 | `data_stream_element` | Known | ATA/disk interface |
| 0x00B45585 | `section_data` | Known | ATA/disk interface |
| 0x00B45836 | `scale_factor_data` | Known | ATA/disk interface |
| 0x00B45848 | `spectral_data` | Known | ATA/disk interface |
| 0x00B4587B | `tns_data` | Known | ATA/disk interface |
| 0x00B46973 | `.rdata` | Known | ATA/disk interface |
| 0x00B4697A | `.rsdata` | Known | ATA/disk interface |
| 0x00B46988 | `.sdata` | Known | ATA/disk interface |
| 0x00B469A8 | `.rela.rdata` | Known | ATA/disk interface |
| 0x00B469B4 | `.rela.rsdata` | Known | ATA/disk interface |
| 0x00B469D6 | `.rela.sdata` | Known | ATA/disk interface |
| 0x00B46A32 | `.rela.data` | Known | ATA/disk interface |
| 0x00B5CF20 | `h264_setrefdata` | Known | ATA/disk interface |
| 0x00B5CF5F | `h264_chromaplane_data` | Known | ATA/disk interface |
| 0x00B5CF75 | `h264_lumaplane_data` | Known | ATA/disk interface |
| 0x00B5E764 | `g_refdata` | Known | ATA/disk interface |
| 0x00B81521 | `mpeg4_startdata` | Known | ATA/disk interface |
| 0x00B81531 | `mpeg4_blockdata` | Known | ATA/disk interface |
| 0x00B81541 | `mpeg4_lastblockdata` | Known | ATA/disk interface |
| 0x00B819F4 | `mpeg4dec_fetch_blocks_dma_subchan` | Known | Hardware interface |
| 0x00B81A16 | `mpeg4dec_fetch_blocks_dma_chan` | Known | Hardware interface |
| 0x00B81A35 | `mpeg4dec_fetch_blocks_dma_cba` | Known | Hardware interface |
| 0x00B81A5D | `mpeg4_deststripedatay` | Known | ATA/disk interface |
| 0x00B81A73 | `mpeg4_deststripedatau` | Known | ATA/disk interface |
| 0x00B81A89 | `mpeg4_deststripedatav` | Known | ATA/disk interface |
| 0x00B81AB7 | `mpeg4dec_dma_xfer_q` | Known | Hardware interface |
| 0x00B81B31 | `mpeg4_stripedatay` | Known | ATA/disk interface |
| 0x00B81B43 | `mpeg4_stripedatau` | Known | ATA/disk interface |
| 0x00B81B55 | `mpeg4_stripedatav` | Known | ATA/disk interface |
| 0x00B81C1A | `mpeg4_fetcheddata` | Known | ATA/disk interface |
| 0x00B81C65 | `fastparse_preparedma` | Known | Hardware interface |
| 0x00B81DE8 | `launchdma` | Known | Hardware interface |
| 0x00B81DF2 | `launchdma2` | Known | Hardware interface |
| 0x00B82431 | `ubv_initintratables` | Known | ATA/disk interface |
| 0x00B82773 | `waitfordma` | Known | Hardware interface |
| 0x00B82783 | `waitfordma2` | Known | Hardware interface |
| 0x00B82864 | `ubv_vc_intratable` | Known | ATA/disk interface |
| 0x00C49BDF | `o}yzsPI` | Known | Hardware interface |

---

## 7. Logging/Analytics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002166F4 | `Channel Reserved` | Hidden | Logging channel |
| 0x00216708 | `Channel AppBoot` | Hidden | Logging channel |
| 0x00216718 | `Channel BufferedSongReading` | Hidden | Logging channel |
| 0x00216734 | `Channel PrefsWriting` | Hidden | Logging channel |
| 0x0021674C | `Channel GeneralUserExperience` | Hidden | Logging channel |
| 0x002167A0 | `Channel TestLogging` | Hidden | Logging channel |
| 0x002167B4 | `Channel AppFileLoading` | Hidden | Logging channel |
| 0x002167CC | `Channel VCardReading` | Hidden | Logging channel |
| 0x002167E4 | `Channel LongSongScanning` | Hidden | Logging channel |
| 0x00216858 | `Channel VoiceRecording` | Hidden | Logging channel |
| 0x00216870 | `Channel VoiceRecordingNewFileSegment` | Hidden | Logging channel |
| 0x00216898 | `Channel PhotoBrowse` | Hidden | Logging channel |
| 0x002168AC | `Channel PhotoImporting` | Hidden | Logging channel |
| 0x002168C4 | `Channel Notes` | Hidden | Logging channel |
| 0x002168D4 | `Channel PhotoFileManagement` | Hidden | Logging channel |
| 0x00216920 | `Channel USBChannel` | Hidden | Logging channel |
| 0x00216934 | `Channel UnitTests` | Hidden | Logging channel |
| 0x00216948 | `Channel FreeSpaceCache` | Hidden | Logging channel |
| 0x002169C0 | `Channel OnTheGoFileMgmt` | Hidden | Logging channel |
| 0x002169D8 | `Channel SlideShow` | Hidden | Logging channel |
| 0x002169EC | `Channel ImageCache` | Hidden | Logging channel |
| 0x00216A00 | `Channel AlbumArtReading` | Hidden | Logging channel |
| 0x00216A18 | `Channel Video` | Hidden | Logging channel |
| 0x00216A3C | `Channel ResourceAccess` | Hidden | Logging channel |
| 0x00216A54 | `Channel VideoCoreBoot` | Hidden | Logging channel |
| 0x00216A88 | `Channel StreamCacheAddFile` | Hidden | Logging channel |
| 0x00216AA4 | `Channel FontFileAccess` | Hidden | Logging channel |
| 0x00216ABC | `Channel ScreenLock` | Hidden | Logging channel |
| 0x00216B40 | `Channel ProfilerAccess` | Hidden | Logging channel |
| 0x00216B58 | `Channel eAppAccess` | Hidden | Logging channel |
| 0x00216B6C | `Channel eAppWriteBackCache` | Hidden | Logging channel |
| 0x00216B88 | `Channel TrainerFileAccess` | Hidden | Logging channel |
| 0x00216BA4 | `Channel IapStorage` | Hidden | Logging channel |
| 0x00216BB8 | `Channel XMLParsing` | Hidden | Logging channel |
| 0x00216BCC | `Channel AudioPrompt` | Hidden | Logging channel |
| 0x00216BE0 | `Channel AudioPromptXML` | Hidden | Logging channel |
| 0x00216BF8 | `Channel StreamCacheSeek` | Hidden | Logging channel |
| 0x0021753C | `iPod Usage Stats` | Hidden | Usage telemetry |
| 0x00B179BC | `pm_stop_logging` | Hidden | Internal logging |
| 0x00B179CC | `pm_start_logging` | Hidden | Internal logging |
| 0x00B281A9 | `Pm Logging` | Hidden | Internal logging |

---

## 8. Error Messages

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00086DDC | `Invalid Operation` | Known | Error/assertion message |
| 0x0009F740 | `IP Address:<invalid>` | Known | Error/assertion message |
| 0x000EF54C | `internal error: list index %ld out of range` | Known | Error/assertion message |
| 0x00107680 | `Root Hub Driver Internal Error unused case in hub handl...` | Known | Error/assertion message |
| 0x001076BC | `Root hub Error Calling Add Device` | Known | Error/assertion message |
| 0x00142990 | `Too many errors, further errors discarded. (31)` | Known | Error/assertion message |
| 0x00142CE0 | `%s Error in file %s.` | Known | Error/assertion message |
| 0x00143388 | `Cannot link to a .link file. (29)` | Known | Error/assertion message |
| 0x002D0A34 | `Error durante la importaci` | Known | Error/assertion message |
| 0x002D1964 | `Error desconocido` | Known | Error/assertion message |
| 0x002EB25C | `Errore durante l'importazione` | Known | Error/assertion message |
| 0x002EC034 | `Errore sconosciuto` | Known | Error/assertion message |
| 0x004B53D8 | `connection failed` | Known | Error/assertion message |
| 0x004B7318 | `This game cannot be played.` | Known | Error/assertion message |
| 0x004B749C | `Imported photos cannot be viewed on TV. Transfer photos...` | Known | Error/assertion message |
| 0x004B75F0 | `An error occurred while importing` | Known | Error/assertion message |
| 0x004B7C78 | `%s failed to launch because its resources cannot be fou...` | Known | Error/assertion message |
| 0x004B84A4 | `This file cannot be viewed on iPod.` | Known | Error/assertion message |
| 0x004B84C8 | `This media file cannot be viewed or played on iPod. Use...` | Known | Error/assertion message |
| 0x004B85C0 | `This photo format cannot be viewed on iPod. Transfer im...` | Known | Error/assertion message |
| 0x004B87DC | `Cannot record because there is no microphone attached.` | Known | Error/assertion message |
| 0x006826A6 | `%s: range error: invalid range [%d, %d)` | Known | Error/assertion message |
| 0x006826F9 | `%s: conversion failed` | Known | Error/assertion message |
| 0x00682735 | `%s: failed to construct locale name` | Known | Error/assertion message |
| 0x00682780 | `%s: invalid pointer %p` | Known | Error/assertion message |
| 0x006827A7 | `%s: unspecified error` | Known | Error/assertion message |
| 0x006827BD | `%s: runtime error` | Known | Error/assertion message |
| 0x006827CF | `%s: underflow error` | Known | Error/assertion message |
| 0x006827E3 | `%s: overflow error` | Known | Error/assertion message |
| 0x00682893 | `%s: length error: %u > %u` | Known | Error/assertion message |
| 0x00AF30CE | `@>ShaderMachine: Invalid shader type found` | Known | Error/assertion message |
| 0x00B0D4DC | `error=%d error_msg="odd number of arguments"` | Known | Error/assertion message |
| 0x00B113D0 | `error=%d error_msg="missing argument"` | Known | Error/assertion message |
| 0x00B115CC | `error=%d error_msg="Invalid arguments"` | Known | Error/assertion message |
| 0x00B115F4 | `error=%d error_msg="Command not registered"` | Known | Error/assertion message |
| 0x00B13B12 | `0Berror=%d error_msg="bad display"` | Known | Error/assertion message |
| 0x00B14004 | `error=%d error_msg="bad argument"` | Known | Error/assertion message |
| 0x00B14118 | `error=%d error_msg="dlopen: %s"` | Known | Error/assertion message |
| 0x00B14138 | `error=%d error_msg="dl_local_sym: %s"` | Known | Error/assertion message |
| 0x00B14160 | `error=%d error_msg="app already loaded"` | Known | Error/assertion message |
| 0x00B28DC8 | `:Cannot print floating point:` | Known | Error/assertion message |
| 0x00B456A3 | `adts_error_check` | Known | Error/assertion message |
| 0x00B81453 | `global_bitstream_error` | Known | Error/assertion message |
| 0x00B81555 | `mpeg4_numinvalidabove` | Known | Error/assertion message |
| 0x00B85C4C | `error=%d error_msg="ff/rew unavailable"` | Known | Error/assertion message |
| 0x00B86828 | `error=%d error_msg="bad parameters"` | Known | Error/assertion message |
| 0x00B86EF4 | `error=%d error_msg="bad parameter"` | Known | Error/assertion message |
| 0x00B871CC | `error=%d error_msg="suspended"` | Known | Error/assertion message |
| 0x00B872CC | `error=%d error_msg="not playing or recording"` | Known | Error/assertion message |
| 0x00B873C2 | `0Berror=%d error_msg="not available"` | Known | Error/assertion message |
| 0x00B8744C | `error=%d error_msg="not recording"` | Known | Error/assertion message |
| 0x00B8795C | `error=%d error_msg="not suspended"` | Known | Error/assertion message |
| 0x00B87AEA | `0Berror=%d error_msg="not playing"` | Known | Error/assertion message |
| 0x00B87B10 | `error=%d error_msg="no video stream"` | Known | Error/assertion message |
| 0x00B87B38 | `error=%d error_msg="screen capture in progress"` | Known | Error/assertion message |
| 0x00B87BEC | `error=%d error_msg="recording"` | Known | Error/assertion message |
| 0x00B87C0C | `error=%d error_msg="out of range"` | Known | Error/assertion message |
| 0x00B87C30 | `error=%d error_msg="stream not active"` | Known | Error/assertion message |
| 0x00B87D34 | `error=%d error_msg="busy"` | Known | Error/assertion message |
| 0x00B87E76 | `0Berror=%d error_msg="bad transform"` | Known | Error/assertion message |
| 0x00B87F52 | `0Berror=%d error_msg="step unavailable"` | Known | Error/assertion message |
| 0x00B87F7C | `error=%d error_msg="step in progress"` | Known | Error/assertion message |
| 0x00B88008 | `error=%d error_msg="idle"` | Known | Error/assertion message |
| 0x00B88084 | `error=%d error_msg="already suspended"` | Known | Error/assertion message |
| 0x00B88104 | `error=%d error_msg="failed"` | Known | Error/assertion message |
| 0x00B98B1A | `0Berror=%d error_msg="bad parameters"` | Known | Error/assertion message |
| 0x00B98BB4 | `error=%d error_msg="bad transform"` | Known | Error/assertion message |
| 0x00B98EBE | `0Berror=%d error_msg="busy"` | Known | Error/assertion message |
| 0x00B98EFC | `error=%d error_msg="bad direction"` | Known | Error/assertion message |
| 0x00B98F20 | `error=%d error_msg="not playing image"` | Known | Error/assertion message |

---

## 9. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0007D68C | `;9?=3175+)/-#!'%[Y_]SQWUKIOMCAGE{y` | Known | Filesystem path |
| 0x0007D7D9 | `\|yz;8=>7412# %&/,)*` | Known | Filesystem path |
| 0x0007D88C | `\|ungXQJC4=&/` | Known | Filesystem path |
| 0x0007DAE0 | `85"/di~sP]JG` | Known | Filesystem path |
| 0x0007DB4C | `MCQ_u{ig=3!/` | Known | Filesystem path |
| 0x000E3894 | `iPod_Control/%s%s%s` | Known | Filesystem path |
| 0x000E38A8 | `iPod_Control/%s/%s%s%s` | Known | Filesystem path |
| 0x0014398C | `Bad link, no matching </a> for anchor tag. (20)` | Known | Filesystem path |
| 0x0014CC40 | `%s<key>%s</key>` | Known | Filesystem path |
| 0x0014CC50 | `%s<integer>%d</integer>` | Known | Filesystem path |
| 0x0014CDB0 | `%s<string>%s</string>` | Known | Filesystem path |
| 0x0014CE34 | `%s<%s/>` | Known | Filesystem path |
| 0x0014CE74 | `%s</dict>` | Known | Filesystem path |
| 0x0014CEB4 | `%s</array>` | Known | Filesystem path |
| 0x0014CFB0 | `%s<real>%s</real>` | Known | Filesystem path |
| 0x001724D8 | `paMB rtSDIrp/P` | Known | Filesystem path |
| 0x001939FC | `Created: %d/%d/%4d %d:%02d:%02d %s` | Known | Filesystem path |
| 0x00193A20 | `Last Accessed: %d/%d/%4d %2d:%02d:%02d %s` | Known | Filesystem path |
| 0x00193A4C | `Modified: %d/%d/%4d %2d:%02d:%02d %s` | Known | Filesystem path |
| 0x001AC3AC | `/iPod_Control/Device/Accessories` | Known | Filesystem path |
| 0x001ACD40 | `/iPod_Control/Device/Accessories/Tags` | Known | Filesystem path |
| 0x001ACD8C | `%s/Tags/%lu.plist` | Known | Filesystem path |
| 0x001ACDA0 | `%s/Tags/%lu.p7` | Known | Filesystem path |
| 0x00218239 | `Average navigation (Next/Prev) per playback duration: %...` | Known | Filesystem path |
| 0x0024723C | `iPod_Control/Device` | Known | Filesystem path |
| 0x00247250 | `iPod_Control/Device/radio` | Known | Filesystem path |
| 0x00247744 | `Resources/Fonts` | Known | Filesystem path |
| 0x0027B398 | `iPod S/N` | Known | Filesystem path |
| 0x002B21EC | `%-m/%-d` | Known | Filesystem path |
| 0x002B220C | `%-m/%-d/%y` | Known | Filesystem path |
| 0x002B24E1 | `ce Features Guide nebo na adrese www.apple.com/support/...` | Known | Filesystem path |
| 0x002B26F8 | `re: %d (%d/%d)` | Known | Filesystem path |
| 0x002B3D9F | `ky naleznete na adrese www.apple.com/support/ipod.` | Known | Filesystem path |
| 0x002B3E38 | `apple.com/support/ipod` | Known | Filesystem path |
| 0x002B3FBB | ` informace naleznete na adrese http://apple.com/support...` | Known | Filesystem path |
| 0x002B4621 | `USA a/nebo dal` | Known | Filesystem path |
| 0x002B5B24 | `www.apple.com/support` | Known | Filesystem path |
| 0x002B8AAB | ` www.apple.com/dk/support/ipod.` | Known | Filesystem path |
| 0x002B8C6C | `Point: %d (%d/%d)` | Known | Filesystem path |
| 0x002BA296 | ` adressen www.apple.com/support/dk/ipod.` | Known | Filesystem path |
| 0x002BA348 | `apple.com/dk/support/ipod` | Known | Filesystem path |
| 0x002BA4D8 | ` http://www.apple.com/dk/support/ipod/` | Known | Filesystem path |
| 0x002BAA7B | `rende VoiceAge Corporation i USA og/eller andre lande o...` | Known | Filesystem path |
| 0x002BBD9C | `www.apple.com/dk/support` | Known | Filesystem path |
| 0x002BF20A | ` bewegen. Weitere Informationen finden Sie im iPod Hand...` | Known | Filesystem path |
| 0x002BF438 | `Punkte: %d (%d/%d)` | Known | Filesystem path |
| 0x002BFE3C | `Vorn./Nachn.` | Known | Filesystem path |
| 0x002BFE4C | `Nachn./Vorn.` | Known | Filesystem path |
| 0x002C0CA3 | ` auf Ihrem iPod. Weitere Anleitungen finden Sie im iPod...` | Known | Filesystem path |
| 0x002C0DAC | `apple.com/de/support/ipod` | Known | Filesystem path |
| 0x002C0F18 | `Weitere Informationen finden Sie unter: http://apple.co...` | Known | Filesystem path |
| 0x002C124C | `ber die Start/Pause-Taste von jedem ausgew` | Known | Filesystem path |
| 0x002C1513 | ` ist entweder eine eingetragene Marke oder eine Marke d...` | Known | Filesystem path |
| 0x002C299C | `www.apple.com/de/support` | Known | Filesystem path |
| 0x002C5CAA | ` www.apple.com/support/ipod.` | Known | Filesystem path |
| 0x002C5FAC | `: %d (%d/%d)` | Known | Filesystem path |
| 0x002C8C65 | ` http://apple.com/support/ipod/` | Known | Filesystem path |
| 0x002CEB2D | `jase a www.apple.com/es/support/ipod.` | Known | Filesystem path |
| 0x002CED2C | `Result.: %d (%d/%d)` | Known | Filesystem path |
| 0x002D045A | `nea de dicho manual en www.apple.com/es/support/ipod.` | Known | Filesystem path |
| 0x002D0514 | `apple.com/es/support/ipod` | Known | Filesystem path |
| 0x002D06A5 | `n, visite http://apple.com/es/support/ipod/` | Known | Filesystem path |
| 0x002D0C43 | ` es una marca registrada o una marca comercial de Voice...` | Known | Filesystem path |
| 0x002D1298 | `Fecha/hora` | Known | Filesystem path |
| 0x002D2154 | `www.apple.com/es/support` | Known | Filesystem path |
| 0x002D510E | `tietoja annetaan iPodin ominaisuusoppaassa tai osoittee...` | Known | Filesystem path |
| 0x002D51E4 | `%s / %s` | Known | Filesystem path |
| 0x002D5208 | `%d / %d` | Known | Filesystem path |
| 0x002D5210 | `%d / %d valokuvaa tuotu` | Known | Filesystem path |
| 0x002D530C | `Tulos: %d (%d/%d)` | Known | Filesystem path |
| 0x002D6977 | `ytyy verkosta osoitteesta www.apple.com/fi/support/ipod...` | Known | Filesystem path |
| 0x002D6A40 | `apple.com/fi/support/ipod` | Known | Filesystem path |
| 0x002D6BD4 | `ytyy osoitteesta http://www.apple.com/fi/support/ipod/` | Known | Filesystem path |
| 0x002D710B | ` on VoiceAge Corporationin Yhdysvalloissa ja/tai muissa...` | Known | Filesystem path |
| 0x002D8480 | `www.apple.com/fi/support` | Known | Filesystem path |
| 0x002DBEFA | ` www.apple.com/fr/support/ipod.` | Known | Filesystem path |
| 0x002DD9B6 | ` l'adresse www.apple.com/fr/support/ipod.` | Known | Filesystem path |
| 0x002DDA58 | `apple.com/fr/support/ipod` | Known | Filesystem path |
| 0x002DDBF0 | `Pour en savoir plus, veuillez visiter le site http://ap...` | Known | Filesystem path |
| 0x002DE2BA | `tats-Unis et/ou dans d` | Known | Filesystem path |
| 0x002DEC73 | `gler date/heure` | Known | Filesystem path |
| 0x002DF998 | `www.apple.com/fr/support` | Known | Filesystem path |
| 0x002E2A0B | `togassa meg a www.apple.com/support/ipod weboldalt.` | Known | Filesystem path |
| 0x002E2AF0 | `%d / %d f` | Known | Filesystem path |
| 0x002E2C14 | `m: %d (%d/%d)` | Known | Filesystem path |
| 0x002E43FF | `i a www.apple.com/support/ipod c` | Known | Filesystem path |
| 0x002E4668 | `togasson el a http://apple.com/support/ipod/ c` | Known | Filesystem path |
| 0x002E4CF3 | `s/vagy m` | Known | Filesystem path |
| 0x002E9344 | ` di iPod" o vai al sito web www.apple.com/it/support/ip...` | Known | Filesystem path |
| 0x002E9554 | `Punti: %d (%d/%d)` | Known | Filesystem path |
| 0x002EABF2 | `, consulta la Guida alle caratteristiche di iPod. Sono ...` | Known | Filesystem path |
| 0x002EAD08 | `apple.com/it/support/ipod` | Known | Filesystem path |
| 0x002EAE7C | `Per ulteriori informazioni, consulta il sito http://app...` | Known | Filesystem path |
| 0x002EAF5C | `Per ulteriori informazioni, consulta http://apple.com/i...` | Known | Filesystem path |
| 0x002EC84C | `www.apple.com/it/support` | Known | Filesystem path |
| 0x002EF554 | `%b/%-d %-I:%M %p` | Known | Filesystem path |
| 0x002EF58C | `%y/%-m/%-d` | Known | Filesystem path |
| 0x002EF598 | `%Y/%b/%-d` | Known | Filesystem path |
| 0x002EF5A4 | `%y/%b/%-d` | Known | Filesystem path |
| 0x002EF89C | ` www.apple.com/jp/support/ipod ` | Known | Filesystem path |
| 0x002EFAF0 | `%d (%d/%d)` | Known | Filesystem path |
| 0x002F1B22 | `www.apple.com/jp/support/ipod ` | Known | Filesystem path |
| 0x002F1C14 | `apple.com/jp/support/ipod` | Known | Filesystem path |
| 0x002F1DF0 | `http://www.apple.com/jp/support/ipod/ ` | Known | Filesystem path |
| 0x002F3F80 | `www.apple.com/jp/support` | Known | Filesystem path |
| 0x002F6CE4 | `%Y/%b/%d %A  %I:%M:%S %p` | Known | Filesystem path |
| 0x002F6D00 | `%Y/%b/%d` | Known | Filesystem path |
| 0x002F6D18 | `%-m/%-d %-I:%M %p` | Known | Filesystem path |
| 0x002F6D4C | `%Y/%-m/%-d` | Known | Filesystem path |
| 0x002F7039 | ` www.apple.co.kr/support/ipod` | Known | Filesystem path |
| 0x002F7114 | `%d / %d ` | Known | Filesystem path |
| 0x002F8E94 | `apple.co.kr/support/ipod` | Known | Filesystem path |
| 0x002F9039 | ` http://www.apple.co.kr/support/ipod/` | Known | Filesystem path |
| 0x002FAC40 | `www.apple.co.kr/support` | Known | Filesystem path |
| 0x002FE094 | `Om hier tekstbestanden te bekijken, stelt u de iPod in ...` | Known | Filesystem path |
| 0x002FE324 | `Score: %d (%d/%d)` | Known | Filesystem path |
| 0x002FF94F | `Raadpleeg de iPod-overzichtshandleiding voor informatie...` | Known | Filesystem path |
| 0x002FFAB4 | `apple.com/nl/support/ipod` | Known | Filesystem path |
| 0x002FFC44 | `Meer informatie vindt u op http://apple.com/nl/support/...` | Known | Filesystem path |
| 0x003001BB | ` is een gedeponeerd handelsmerk of een handelsmerk van ...` | Known | Filesystem path |
| 0x00300AF8 | `Stel datum/tijd in` | Known | Filesystem path |
| 0x003015FC | `www.apple.com/nl/support` | Known | Filesystem path |
| 0x00304575 | `r til www.apple.com/no/support/ipod.` | Known | Filesystem path |
| 0x00304728 | `Poeng: %d (%d/%d)` | Known | Filesystem path |
| 0x00305D86 | ` www.apple.com/no/support/ipod.` | Known | Filesystem path |
| 0x00305E18 | `apple.com/no/support/ipod` | Known | Filesystem path |
| 0x00305FA3 | ` http://www.apple.com/no/support/ipod/` | Known | Filesystem path |
| 0x003064EF | ` er enten et registrert varemerke eller et varemerke fo...` | Known | Filesystem path |
| 0x003077E0 | `www.apple.com/no/support` | Known | Filesystem path |
| 0x0030A4D4 | `%-d/%-m/%y` | Known | Filesystem path |
| 0x0030A7B1 | `ytkownika iPoda lub na stronie www.apple.com/support/ip...` | Known | Filesystem path |
| 0x0030A9A0 | `Punkty: %d (%d/%d)` | Known | Filesystem path |
| 0x0030C0E3 | `ugi iPoda. Wersja elektroniczna instrukcji na stronie w...` | Known | Filesystem path |
| 0x0030C35B | ` pod adresem http://apple.com/support/ipod/` | Known | Filesystem path |
| 0x0030C91A | `onym znakiem towarowym lub znakiem towarowym firmy Voic...` | Known | Filesystem path |
| 0x00310B44 | `%-d/%-m` | Known | Filesystem path |
| 0x00310E1F | ` para www.apple.com/support/ipod para obter mais inform...` | Known | Filesystem path |
| 0x0031102A | `o: %d (%d/%d)` | Known | Filesystem path |
| 0x003111FE | `d. p/ desbloq.` | Known | Filesystem path |
| 0x00311508 | `Prima Repr. p/ desligar r` | Known | Filesystem path |
| 0x00311540 | `o central p/ guardar a esta` | Known | Filesystem path |
| 0x003115D8 | `Prima Anterior/Seguinte para mudar de Esta` | Known | Filesystem path |
| 0x0031273D | `es online deste guia podem ser encontradas em www.apple...` | Known | Filesystem path |
| 0x00312981 | `es, consulte http://apple.com/support/ipod/` | Known | Filesystem path |
| 0x00312A28 | `es adicionais, consulte http://apple.com/support/ipod/` | Known | Filesystem path |
| 0x00312DE4 | `o p/ impedir mais altera` | Known | Filesystem path |
| 0x00312E32 | `o central p/ continuar.` | Known | Filesystem path |
| 0x00312EDE | ` uma marca comercial ou marca registada da VoiceAge Cor...` | Known | Filesystem path |
| 0x0031762B | ` www.apple.com/support/ipod ` | Known | Filesystem path |
| 0x0031A292 | `: www.apple.com/ru/support/ipod.` | Known | Filesystem path |
| 0x0031A603 | `: http://apple.com/support/ipod/` | Known | Filesystem path |
| 0x003202FE | ` www.apple.com/se/support/ipod.` | Known | Filesystem path |
| 0x003204B8 | `ng: %d (%d/%d)` | Known | Filesystem path |
| 0x00321AAD | ` adressen www.apple.com/se/support/ipod.` | Known | Filesystem path |
| 0x00321B50 | `apple.com/support/se/ipod` | Known | Filesystem path |
| 0x00321CE1 | ` http://www.apple.com/se/support/ipod/` | Known | Filesystem path |
| 0x00322278 | `r VoiceAge Corporation i USA och/eller andra l` | Known | Filesystem path |
| 0x00322B24 | `ll in datum/tid` | Known | Filesystem path |
| 0x003235A8 | `www.apple.com/se/support` | Known | Filesystem path |
| 0x00326240 | `%d/%m %-H:%M` | Known | Filesystem path |
| 0x00326556 | `n ya da www.apple.com/support/ipod adresine gidin.` | Known | Filesystem path |
| 0x0032663C | `%d / %d foto` | Known | Filesystem path |
| 0x00326750 | `Puan: %d (%d/%d)` | Known | Filesystem path |
| 0x00327E04 | `mleri www.apple.com/support/ipod adresinde bulunabilir.` | Known | Filesystem path |
| 0x00328064 | `in http://apple.com/support/ipod/ adresini ziyaret edin...` | Known | Filesystem path |
| 0x0032813E | `tfen http://apple.com/support/ipod/ adresini ziyaret ed...` | Known | Filesystem path |
| 0x00328371 | `alma/oynatma d` | Known | Filesystem path |
| 0x0032865F | `n ABD ve/veya di` | Known | Filesystem path |
| 0x0032CAE4 | ` www.apple.com.cn/support/ipod` | Known | Filesystem path |
| 0x0032CBAD | ` %d/%d ` | Known | Filesystem path |
| 0x0032CCB1 | ` %d (%d/%d)` | Known | Filesystem path |
| 0x0032E4BF | ` www.apple.com.cn/support/ipod ` | Known | Filesystem path |
| 0x0032E55C | `apple.com.cn/support/ipod` | Known | Filesystem path |
| 0x0032E6FE | ` http://www.apple.com.cn/support/ipod/` | Known | Filesystem path |
| 0x0032FF04 | `www.apple.com.cn/support` | Known | Filesystem path |
| 0x00332F5C | ` www.apple.com.tw/support/ipod` | Known | Filesystem path |
| 0x00333021 | ` %d / %d ` | Known | Filesystem path |
| 0x0033493A | `www.apple.com.tw/support/ipod` | Known | Filesystem path |
| 0x00334B89 | `http://www.apple.com.tw/support/ipod/` | Known | Filesystem path |
| 0x00336438 | `http://www.apple.com.tw/support` | Known | Filesystem path |
| 0x003EB7AA | `pcefefefefefefefefefefefefefefefefefefefefef/[6` | Known | Filesystem path |
| 0x003EB80D | `=/[2\|pc` | Known | Filesystem path |
| 0x003EC94A | `2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2...` | Known | Filesystem path |
| 0x003ECC40 | `2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2...` | Known | Filesystem path |
| 0x003ECF36 | `2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|/...` | Known | Filesystem path |
| 0x003ED22C | `2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003ED522 | `2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003ED818 | `2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EDAFE | `pcefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|...` | Known | Filesystem path |
| 0x003EDDF2 | `pcefefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2...` | Known | Filesystem path |
| 0x003EE0E6 | `pcefefefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\...` | Known | Filesystem path |
| 0x003EE3DA | `pcefefefefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|...` | Known | Filesystem path |
| 0x003EE6CE | `pcefefefefefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|/...` | Known | Filesystem path |
| 0x003EE9C2 | `pcefefefefefefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EECB6 | `pcefefefefefefefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EEFAA | `pcefefefefefefefefefefefefefef2\|2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EF29E | `pcefefefefefefefefefefefefefefef2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EF592 | `pcefefefefefefefefefefefefefefefef2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EF886 | `pcefefefefefefefefefefefefefefefefef2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EFB7A | `pcefefefefefefefefefefefefefefefefefef2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EFE6E | `pcefefefefefefefefefefefefefefefefefefef2\|2\|/[` | Known | Filesystem path |
| 0x003F0162 | `pcefefefefefefefefefefefefefefefefefefefef2\|/[6` | Known | Filesystem path |
| 0x00405425 | `d/(1010101010101010101010101010101010101010101010101010...` | Known | Filesystem path |
| 0x004A17C5 | `cpcocococOc/[ocpcOc` | Known | Filesystem path |
| 0x004A1F55 | `cO[/[ococ` | Known | Filesystem path |
| 0x004A296F | `cococococococOc/[ocococococ` | Known | Filesystem path |
| 0x004A29E3 | `kO[N[oc/[O[pc` | Known | Filesystem path |
| 0x004A2BCF | `R/[ocO[N[ocpcOcocN[` | Known | Filesystem path |
| 0x004A2BE9 | `kococococococOcOc/[.[N[/[O[/[N[Oc` | Known | Filesystem path |
| 0x004A2E27 | `kocOc/[O[` | Known | Filesystem path |
| 0x004A2E35 | `cOcO[N[ocOc/[oc` | Known | Filesystem path |
| 0x004A2E45 | `c/[.[O[/[` | Known | Filesystem path |
| 0x004A2E4F | `S/[OcN[N[OcOcococ` | Known | Filesystem path |
| 0x004A2E67 | `cocococ.[N[oc.[/[O[/[N[N[Ococ/[/[Oc/[OcO[.[N[ocococ/[oc...` | Known | Filesystem path |
| 0x004A2F93 | `cO[/[OcpcPc` | Known | Filesystem path |
| 0x004A30B5 | `kocOcO[O[oc/[O[Oc.[` | Known | Filesystem path |
| 0x004A3337 | `kOcoc/[OcOcO[.[.[.[/[` | Known | Filesystem path |
| 0x004A33B3 | `SO[O[/[ocOc.[.[` | Known | Filesystem path |
| 0x004A33C5 | `[.[/[.[ococN[.[.[.[N[Oc` | Known | Filesystem path |
| 0x004A3401 | `[/[.[N[O[N[` | Known | Filesystem path |
| 0x004A340D | `[N[O[.[OcO[.[N[ococococN[O[OcOcoc/[ococpcocOcococN[.[oc...` | Known | Filesystem path |
| 0x004A3453 | `cococococ/[/[/[/[O[OcocOcpcpcpcpcpcoc/[/[pc` | Known | Filesystem path |
| 0x004A34F3 | `koc/[/[ocOc` | Known | Filesystem path |
| 0x004A35A9 | `cocOc/[ocococ` | Known | Filesystem path |
| 0x004A3787 | `[.[/[/[` | Known | Filesystem path |
| 0x004A37B1 | `[.[/[.[.[/[/[.[` | Known | Filesystem path |
| 0x004A37F5 | `[Oc/[N[Oc.[/[` | Known | Filesystem path |
| 0x004A3835 | `kpcO[/[/[.[N[` | Known | Filesystem path |
| 0x004A3D35 | `[/[Oc.[` | Known | Filesystem path |
| 0x004A4227 | `k/[OcO[OcpcocO[.[.[` | Known | Filesystem path |
| 0x004A44C4 | `/[O[O[Oc` | Known | Filesystem path |
| 0x004A44F3 | `B/[O[/[O[O[O[.[.[.[` | Known | Filesystem path |
| 0x004A4509 | `JkBOcO[O[/[O[Oc.[` | Known | Filesystem path |
| 0x004A45C9 | `J*:O[O[/[` | Known | Filesystem path |
| 0x004A4611 | `JJB/[.[O[N[` | Known | Filesystem path |
| 0x004A4673 | `:O[/[O[.[` | Known | Filesystem path |
| 0x004A46B9 | `J*:O[/[O[O[.[` | Known | Filesystem path |
| 0x004A4733 | `1/[.[.[` | Known | Filesystem path |
| 0x004A9885 | `R/[O[Ocpc` | Known | Filesystem path |
| 0x004A989D | `J/[OcOcpc` | Known | Filesystem path |
| 0x004A98CD | `R/[O[OcOcpcpc` | Known | Filesystem path |
| 0x004A9919 | `S.S/[pc` | Known | Filesystem path |
| 0x004A9979 | `SOcO[/[pcpc` | Known | Filesystem path |
| 0x004A9993 | `[/[O[/[.SO[` | Known | Filesystem path |
| 0x004A99AF | `S/[O[/[pcpc` | Known | Filesystem path |
| 0x004A99D7 | `R/[O[/[` | Known | Filesystem path |
| 0x004A9A29 | `cOc/[/[/[` | Known | Filesystem path |
| 0x004A9A39 | `R/[/[O[pcpc` | Known | Filesystem path |
| 0x004A9A53 | `S/[/[.S/[` | Known | Filesystem path |
| 0x004A9A6D | `[/[/[O[OcpcO[` | Known | Filesystem path |
| 0x004A9A83 | `S/[/[PcP[/[O[pc` | Known | Filesystem path |
| 0x004A9AB3 | `[/[/[P[pcpcOcO[` | Known | Filesystem path |
| 0x004A9AD1 | `S/[pcpc` | Known | Filesystem path |
| 0x004A9AE7 | `R/[O[pcpcpc` | Known | Filesystem path |
| 0x004A9AFF | `Spcpc/[/[` | Known | Filesystem path |
| 0x004A9B31 | `R/[/[pc` | Known | Filesystem path |
| 0x004A9B47 | `R/[pc/[/[P[` | Known | Filesystem path |
| 0x004A9B5D | `[/[/[O[` | Known | Filesystem path |
| 0x004A9B93 | `S/[O[/[` | Known | Filesystem path |
| 0x004A9BA7 | `S/[/[/[P[P[` | Known | Filesystem path |
| 0x004A9BBF | `R/[/[P[P[/[` | Known | Filesystem path |
| 0x004A9BDB | `S/[pcO[` | Known | Filesystem path |
| 0x004A9C05 | `S.SPcP[/[O[` | Known | Filesystem path |
| 0x004A9C1F | `S/[pcP[/[` | Known | Filesystem path |
| 0x004A9CB1 | `[O[/[P[/[` | Known | Filesystem path |
| 0x004B6FD5 | `Refer to the iPod Features Guide for instructions on ho...` | Known | Filesystem path |
| 0x004B7278 | `For more information, please visit http://apple.com/sup...` | Known | Filesystem path |
| 0x004B7334 | `For additional information, please visit http://apple.c...` | Known | Filesystem path |
| 0x004B7807 | ` is either registered trademark or trademark of VoiceAg...` | Known | Filesystem path |
| 0x004F8F91 | ` !"#$%&'()*+,-./0123456789:;<=>?@abcdefghijklmnopqrstuv...` | Known | Filesystem path |
| 0x004F9091 | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x004FF7D0 | `$X/wTNw` | Known | Filesystem path |
| 0x005008C0 | `{{~~  /-----\   {{~~ /       \  {{~~\|         \| {{~~\...` | Known | Filesystem path |
| 0x00500A07 | `<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1....` | Known | Filesystem path |
| 0x00500A8E | `</plist>` | Known | Filesystem path |
| 0x00500CAA | `_/:>v?J7` | Known | Filesystem path |
| 0x0050538C | `!"#$%&'.,()+-=_/:;<>?@[]abcdefghijklmnopqrstuvwxyzABCDE...` | Known | Filesystem path |
| 0x0051CEB7 | `W/}lE>q` | Known | Filesystem path |
| 0x0054BAF9 | `H."0*Bx/` | Known | Filesystem path |
| 0x00553456 | `U/~RERT` | Known | Filesystem path |
| 0x005581BE | `TUOPT/\|` | Known | Filesystem path |
| 0x0055F453 | `HuGZp/$j` | Known | Filesystem path |
| 0x00565963 | `(bJ)b"(b2""z/` | Known | Filesystem path |
| 0x005680D3 | `JUAPDD(/` | Known | Filesystem path |
| 0x0056E3E2 | `/B\|$BD'` | Known | Filesystem path |
| 0x0056F35F | `$Bd$BT/` | Known | Filesystem path |
| 0x00575637 | `/" +J\|!` | Known | Filesystem path |
| 0x0057C4B6 | `Fb""")/` | Known | Filesystem path |
| 0x0057D66D | `/RyO(UIH` | Known | Filesystem path |
| 0x0057E71D | `~$Bh'"~$Bz/` | Known | Filesystem path |
| 0x00581F87 | `$B +BZ/` | Known | Filesystem path |
| 0x005899A5 | `0c(HBP/` | Known | Filesystem path |
| 0x0058DF3B | `$B~("\|/` | Known | Filesystem path |
| 0x005A5F2D | `T/DDDDD` | Known | Filesystem path |
| 0x005A6197 | `"~UeB /` | Known | Filesystem path |
| 0x005A8C49 | `$B((B /` | Known | Filesystem path |
| 0x005B0D2C | ` "\|$B~/` | Known | Filesystem path |
| 0x005B3EA8 | `@$B\|$"(/` | Known | Filesystem path |
| 0x005B4FA8 | `)"8/B""` | Known | Filesystem path |
| 0x005B56F0 | `r4c6 bN/` | Known | Filesystem path |
| 0x005BAAF1 | `RDT%B(/` | Known | Filesystem path |
| 0x005BBC7D | `RBHUE\|/` | Known | Filesystem path |
| 0x005C32FD | `]B""B</` | Known | Filesystem path |
| 0x005C6B7A | `,B\|RED/` | Known | Filesystem path |
| 0x005CC5FD | `$BT). /` | Known | Filesystem path |
| 0x005CD7FD | `#"TUB(/` | Known | Filesystem path |
| 0x005FF18D | `x$DDC/T` | Known | Filesystem path |
| 0x00616DAD | `/" %BD"` | Known | Filesystem path |
| 0x0061D6D0 | `ODD""(/` | Known | Filesystem path |
| 0x0061E90B | `B"$R%"B$" /` | Known | Filesystem path |
| 0x0061F278 | `bG\|jG\|/` | Known | Filesystem path |
| 0x00620E62 | `$E$$BR/` | Known | Filesystem path |
| 0x00620F03 | `dRB~RA$/` | Known | Filesystem path |
| 0x00621854 | `TT&T%B(/` | Known | Filesystem path |
| 0x00630548 | `)'>$B8/` | Known | Filesystem path |
| 0x00632B56 | `$B\|%EV/` | Known | Filesystem path |
| 0x0063A0AA | `BDU!BJ ""/` | Known | Filesystem path |
| 0x0063B00E | `Z-bD("(%B>/` | Known | Filesystem path |
| 0x0064375C | `!"#$%&'()*+,-./` | Known | Filesystem path |
| 0x00643F16 | `on543k'78%/e/"#`34 '=3?49-?:))60` | Known | Filesystem path |
| 0x006440B1 | `VcYcmo8jics' EfFf~z/` | Known | Filesystem path |
| 0x00644269 | `J=/&5 1Y` | Known | Filesystem path |
| 0x00645DC8 | ` ,;=+[]*?<>\|":/\` | Known | Filesystem path |
| 0x00645DD9 | `\/:*?"<>\|` | Known | Filesystem path |
| 0x0064C23D | `% %!%"%#%$%%%&%'%(%)%*%+%,%-%.%/%0%1%2%3%4%5%6%7%8%9%:%...` | Known | Filesystem path |
| 0x0064CA1D | `qWlIl/Ymg*` | Known | Filesystem path |
| 0x0064E28B | `X)W,W*W3W9W.W/W\W;WBWiW` | Known | Filesystem path |
| 0x0064E959 | `n/o6oKoto*o` | Known | Filesystem path |
| 0x0064EEF1 | `q/q1qsq\qhqEqrqJqxqzq` | Known | Filesystem path |
| 0x00650179 | `b6bKbNb/e` | Known | Filesystem path |
| 0x0065032F | `V3W0W(W-W,W/W)W` | Known | Filesystem path |
| 0x00650403 | `NMOOOGOWO^O4O[OUO0OPOQO=O:O8OCOTO<OFOcO\O`O/ONO6OYO]OHO...` | Known | Filesystem path |
| 0x00650A61 | `e%f-f f'f/f` | Known | Filesystem path |
| 0x00650D63 | `bNc>c/cUcBcFcOcIc:cPc=c*c+c(cMcLcHeIe` | Known | Filesystem path |
| 0x00651015 | `fFUjUfUDU^UaUCUJU1UVUOUUU/UdU8U.U\U,UcU3UAUWU` | Known | Filesystem path |
| 0x0065127B | `\|F}C}q}.}9}<}@}0}3}D}/}B}2}1}=` | Known | Filesystem path |
| 0x00651403 | `W/X*X4X$X0X1X!X` | Known | Filesystem path |
| 0x00651519 | `k.l/l,l/n8nTn!n2ngnJn n%n#n` | Known | Filesystem path |
| 0x0065157B | `r6s%s4s)s:t*t3t"t%t5t6t4t/t` | Known | Filesystem path |
| 0x00651DF9 | `S.V;V9V2V?V4V)VSVNVWVtV6V/V0V` | Known | Filesystem path |
| 0x0065243F | `\|.~>~F~7~2~C~+~=~1~E~A~4~9~H~5~?~/~D` | Known | Filesystem path |
| 0x0065342F | `P P'P5P/P1P` | Known | Filesystem path |
| 0x00653581 | `h5h+h-h/hNhDh4h` | Known | Filesystem path |
| 0x00653595 | `h&h(h.hMh:h%h h,k/k-k1k4kmk` | Known | Filesystem path |
| 0x006536A1 | `w"w'w#x,x"x5x/x(x.x+x!x)x3x*x1xTy[yOy\ySyRyQy` | Known | Filesystem path |
| 0x00653B2B | `v9w/w-w1w2w4w3w=w%w;w5wHxRxIxMxJxLx&xExPxdygyiyjycykyay` | Known | Filesystem path |
| 0x00653B87 | `{1{+{-{/{2{8{` | Known | Filesystem path |
| 0x0065401B | `t/uoulu` | Known | Filesystem path |
| 0x00654381 | ``IaJa+aEa6a2a.aFa/aOa)a@a bh` | Known | Filesystem path |
| 0x006543C7 | `d&d0d(dAd5d/d` | Known | Filesystem path |
| 0x0065442F | `hxi4iii@ioiDiviXiAitiLi;iKi7i\iOiQi2iRi/i{i<iFkEkCkBkHk...` | Known | Filesystem path |
| 0x00654A99 | `y+zJz0z/z(z&z` | Known | Filesystem path |
| 0x0065537D | `u/v-v1v=v3v<v5v2v0v` | Known | Filesystem path |
| 0x00655673 | `X-[%[2[#[,['[&[/[.[{[` | Known | Filesystem path |
| 0x00655FA1 | `j<p5p/p7p4p1pBp8p?p:p9p@p;p3pAp` | Known | Filesystem path |
| 0x006561A3 | `w-y1y/yT\|S\|` | Known | Filesystem path |
| 0x0065756F | `%#%3%+%;%K% %/%(%7%?%` | Known | Filesystem path |
| 0x006576EE | `02*2+2,2-2.2/2@272B2C292:212>24222;2623252<2=2?282` | Known | Filesystem path |
| 0x00657C19 | `\7_J_/`P`m`` | Known | Filesystem path |
| 0x00657DF5 | `OHSIT>T/Z` | Known | Filesystem path |
| 0x00657E6B | `i_l*mim/n` | Known | Filesystem path |
| 0x0065835D | `N,p]u/f` | Known | Filesystem path |
| 0x00659069 | `S#S/S1S3S8S@SFSES` | Known | Filesystem path |
| 0x00659169 | `q4V6V2V8VkVdV/VlVjV` | Known | Filesystem path |
| 0x00659411 | `_)_-_8_A_H_L_N_/_Q_V_W_Y_a_m_s_w_` | Known | Filesystem path |
| 0x00659AD7 | `s4s/s)s%s>sNsOs` | Known | Filesystem path |
| 0x0065D8DB | `h>kLp/t` | Known | Filesystem path |
| 0x0065DF65 | `o;v/}7~` | Known | Filesystem path |
| 0x0065ED29 | `e1f/h\q6z` | Known | Filesystem path |
| 0x0065F375 | `UuX/c"dIfKfmh` | Known | Filesystem path |
| 0x00663E22 | `  !"##$%&&'())*+,-../01234556789:;<=>?@ABCDEFGHIJKMNOPQ...` | Known | Filesystem path |
| 0x00664048 | ` !""#$%&''()*+,-./0123456789:;<>?@ABDEFGIJKMNOQRTUVXY[\...` | Known | Filesystem path |
| 0x00664690 | `/B'2N6REMQLEVJ\|aViu\J]lLm` | Known | Filesystem path |
| 0x0067D0BB | `iPod_Control/games_RO/` | Known | Filesystem path |
| 0x0067D185 | `iPod_Control/Device/accessories` | Known | Filesystem path |
| 0x0067D4B8 | `iPod_Control/iTunes/` | Known | Filesystem path |
| 0x0067D4D4 | `Recordings/` | Known | Filesystem path |
| 0x0067D4E0 | `Calendars/` | Known | Filesystem path |
| 0x0067D4EB | `Contacts/` | Known | Filesystem path |
| 0x0067D5F6 | `/Resources/VideoCore` | Known | Filesystem path |
| 0x0067DB37 | `file://` | Known | Filesystem path |
| 0x0067DB45 | `</ROT13>` | Known | Filesystem path |
| 0x0067DB5E | `</TITLE>` | Known | Filesystem path |
| 0x0067DBAD | `</BODY>` | Known | Filesystem path |
| 0x00680DDD | `S/MIME Capabilities` | Known | Filesystem path |
| 0x00681B9A | `S/MIME email` | Known | Filesystem path |
| 0x00681C03 | `S/MIME signing` | Known | Filesystem path |
| 0x00681C30 | `S/MIME encryption` | Known | Filesystem path |
| 0x00681F64 | `S/MIME CA` | Known | Filesystem path |
| 0x006A86D9 | `WV>P7C/;` | Known | Filesystem path |
| 0x006BCF10 | `/,)&/,)&/,)&/,)&` | Known | Filesystem path |
| 0x006BDF4E | `/ 0 1 2 3 4 5 6 7 8 9 : ; < = > ? @ A B C D E F G H I J...` | Known | Filesystem path |
| 0x006BE12F | `! !!!"!#!$!%!&!'!(!)!*!+!,!-!.!/!0!1!2!3!4!5!6!7!8!9!:!...` | Known | Filesystem path |
| 0x006C112F | `*Zj"/8'6V` | Known | Filesystem path |
| 0x0072BD10 | `/1f;{1Q` | Known | Filesystem path |
| 0x0075CB62 | `8/868@8\8a8` | Known | Filesystem path |
| 0x0075CEBA | `S!S$S*S/S3S5S;SOSXSZS\SiSuS{S` | Known | Filesystem path |
| 0x0075CF10 | `T!T$T/T1T6T9T@TCTFTJTNTQTUTXT\T_TbTdTfThTlTxT}T` | Known | Filesystem path |
| 0x0075D14A | `Y Y%Y'Y/Y2Y4Y:Y>YBYDYIYKYQY[Y]YcYeYnYvYyY}Y` | Known | Filesystem path |
| 0x0075D1BE | `Z Z#Z%Z'Z)Z+Z-Z/Z6Z<ZAZGZJZUZZZ`ZgZjZmZwZzZ` | Known | Filesystem path |
| 0x0075D3EA | `_#_)_-_/_1_>_A_H_J_N_Y_]_g_m_s_w_y_` | Known | Filesystem path |
| 0x0075D44A | ``"`+`/`1`3`5`<`C`G`M`P`R`U`Z`]`e`m`p`s`x`}`` | Known | Filesystem path |
| 0x0075D57C | `c!c%c(c+c/c2c6c?cFcIcPcScUcWcYc\c_cccecicncrcwc}c` | Known | Filesystem path |
| 0x0075D5FE | `d*d-d/d6d:dDdHdJdNdRdTdXd`dgdidmdpd{d}d` | Known | Filesystem path |
| 0x0075D762 | `g(g/g1g8g:g=gCgIgQgWgYgfghgjgwg\|g` | Known | Filesystem path |
| 0x0075D91C | `k$k'k,k/k2k5k;k?kCkGkJkLkNkPkTkVkYk\kgkjklkokuk` | Known | Filesystem path |
| 0x0075DA00 | `m+m/m6m?mAmHmKmOmTm\m^mamfmjmlmpmtmym\|m` | Known | Filesystem path |
| 0x0075DA68 | `n'n)n/n2n4n6n:n>nEnOnQnTnXn\n_nancngninknonsnvn{n` | Known | Filesystem path |
| 0x0075DCE2 | `s"s%s,s/s1s4s7s;s?sAsEsPsRsXs`sdshsrsusxs\|s~s` | Known | Filesystem path |
| 0x0075DE88 | `w w"w&w)w-w/w:w<w>wAwGwJwQwXw\whwlwrwzw` | Known | Filesystem path |
| 0x0075E06A | `{ {({/{1{3{6{={A{V{[{]{`{g{i{u{w{z{` | Known | Filesystem path |
| 0x0075E1BC | `~#~(~/~7~;~?~A~H~K~M~R~V~Z~^~b~g~k~p~s~u~y~` | Known | Filesystem path |
| 0x007603E4 | `X X#X&X*X-X/X4X9X=X@XIXOXQXTXWX^XaXdXgXiXkXmXoXuXyX\|X` | Known | Filesystem path |
| 0x007604D8 | `Z#Z%Z'Z)Z+Z-Z/Z1Z<Z@ZFZIZUZZZ`ZbZjZlZtZzZ~Z` | Known | Filesystem path |
| 0x00760704 | `_%_-_/_1_4_@_E_J_L_P_[_a_i_p_w_y_{_` | Known | Filesystem path |
| 0x00760762 | `` `$`/`1`3`5`:`A`F`I`P`R`T`Y`]`_`g`o`s`u`z`` | Known | Filesystem path |
| 0x00760896 | `c#c'c*c/c2c5c9cAcIcKcScUcWcYc\c^cacecgckcqctczc` | Known | Filesystem path |
| 0x00760916 | `d d,d/d4d:d=dFdJdNdQdTdXdZdgdidmdodsd}d` | Known | Filesystem path |
| 0x00760AE8 | `h!h(h/h7h;h@hHhLhPhWhYh[h_hbhehkhmhrhthyh\|h~h` | Known | Filesystem path |
| 0x00760C36 | `k'k,k/k2k5k7k=kCkFkIkLkNkPkSkVkXk[k_kiklkokrkwk` | Known | Filesystem path |
| 0x00760E10 | `o o"o%o)o/o5o8o<o>oAoEoGoKoMoQoToWo^ofohomotoxozo\|o` | Known | Filesystem path |
| 0x00760E98 | `p#p&p,p/p2p7p9p<p>pCpGpIpNpTpXp]p`pcpipkpupxp\|p` | Known | Filesystem path |
| 0x007610DE | `u(u/u7u:uDuYu`ubuduiuou}u` | Known | Filesystem path |
| 0x007611A2 | `w"w$w(w,w/w4w<w>w@wEwJwMwXwZw^wjwrwyw\|w` | Known | Filesystem path |
| 0x007613EC | `\|#\|&\|*\|/\|3\|6\|=\|E\|J\|L\|O\|V\|[\|c\|g\|i\|l\|r\...` | Known | Filesystem path |
| 0x00765FA8 | `t\|tutjthtet^t]t\t:t8t6t5t0t/t.t,t+t)t(t't%t!t t` | Known | Filesystem path |
| 0x0076602C | `rxrirgrfrDrBr?r8r5r1r/r.r-r&r%r r` | Known | Filesystem path |
| 0x007FECC5 | `UUUUUUa/` | Known | Filesystem path |
| 0x007FEE94 | `DDDDDDQ/` | Known | Filesystem path |
| 0x007FEF11 | `""""""0/` | Known | Filesystem path |
| 0x0080FCCD | `""f"" /` | Known | Filesystem path |
| 0x00814AAF | `33333330/` | Known | Filesystem path |
| 0x0081556C | `UUUUUUQ/` | Known | Filesystem path |
| 0x00819CCA | `fffffffa/` | Known | Filesystem path |
| 0x0082CBEA | `fffffffp/` | Known | Filesystem path |
| 0x00838EA3 | `""""""1/` | Known | Filesystem path |
| 0x0084D31B | `/`~UDDE0` | Known | Filesystem path |
| 0x00850585 | `/Wwx'wv` | Known | Filesystem path |
| 0x00851B7D | `/PWwww`` | Known | Filesystem path |
| 0x00856954 | `b""""0/` | Known | Filesystem path |
| 0x00859A60 | `o`o@oP/` | Known | Filesystem path |
| 0x00859A7B | `ofo@oP/` | Known | Filesystem path |
| 0x00871F36 | `UUUUUUUQ/` | Known | Filesystem path |
| 0x0087F214 | `"""2""" /` | Known | Filesystem path |
| 0x00884C38 | `fb6ffff@/` | Known | Filesystem path |
| 0x00886F0F | `/uwwQAy` | Known | Filesystem path |
| 0x00892AD6 | `Gwwww /` | Known | Filesystem path |
| 0x008B4DA7 | `33333331/` | Known | Filesystem path |
| 0x008B624D | `vfggffgq/` | Known | Filesystem path |
| 0x008C1711 | `fffffffq/` | Known | Filesystem path |
| 0x008CEE79 | `VffffP/` | Known | Filesystem path |
| 0x008D1301 | `S3W33 /` | Known | Filesystem path |
| 0x008EB275 | `#T"E21/` | Known | Filesystem path |
| 0x008EB68A | `/`eUUS/` | Known | Filesystem path |
| 0x008F0DAD | `/SDDDDQ` | Known | Filesystem path |
| 0x008FCDF7 | `3333!/{` | Known | Filesystem path |
| 0x008FCE9D | `DDDDDB/^` | Known | Filesystem path |
| 0x008FD503 | `/p_OpOJ` | Known | Filesystem path |
| 0x008FF1CE | `/`o_@33L` | Known | Filesystem path |
| 0x0090A30C | `""#wwww /` | Known | Filesystem path |
| 0x0090ACA8 | `DD&wwww`/` | Known | Filesystem path |
| 0x0090ADED | `DD8wwwwt/` | Known | Filesystem path |
| 0x0090D30B | `336ffffb/` | Known | Filesystem path |
| 0x0091487B | `"""""""1/` | Known | Filesystem path |
| 0x009172C3 | `DDDDD@/` | Known | Filesystem path |
| 0x009195FD | `33333A/` | Known | Filesystem path |
| 0x0091A4FE | `fc4DDB/` | Known | Filesystem path |
| 0x0091B358 | `fffff`/` | Known | Filesystem path |
| 0x0091C6E1 | `wwwwwrL/` | Known | Filesystem path |
| 0x0092AD5B | `wwwwwwwq/` | Known | Filesystem path |
| 0x0093DF2C | `C4343 /` | Known | Filesystem path |
| 0x009650E9 | `/rDDDD ` | Known | Filesystem path |
| 0x00977FB9 | `UUU5UUUa/` | Known | Filesystem path |
| 0x009810E1 | `OqWwwwwP/` | Known | Filesystem path |
| 0x009931D1 | `6ffffc/` | Known | Filesystem path |
| 0x00993B6B | `DDDDDDDA/` | Known | Filesystem path |
| 0x0099446E | `UUU$fff0/` | Known | Filesystem path |
| 0x0099479F | `3334www@/` | Known | Filesystem path |
| 0x00998E27 | `UUPDDD/` | Known | Filesystem path |
| 0x0099CB26 | `_o@/?_i` | Known | Filesystem path |
| 0x009A04D6 | `7wwwwws0/` | Known | Filesystem path |
| 0x009BE8A1 | `TDDDD /` | Known | Filesystem path |
| 0x009DEEBA | `fffffa/` | Known | Filesystem path |
| 0x009F0868 | `www7wwwq/` | Known | Filesystem path |
| 0x00A8E23E | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x00AAFA1D | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x00AB5A0E | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x00ACB073 | `po`?pOp/` | Known | Filesystem path |
| 0x00ACE01A | `/PEBLC2` | Known | Filesystem path |
| 0x00ACFE06 | ` " & / 0 : D ` p y ~ ` | Known | Filesystem path |
| 0x00AD8752 | `/"/0/>/L/Z/h/v/` | Known | Filesystem path |
| 0x00AEE293 | `;0K0B0B0B0B0B0B0Bb/` | Known | Filesystem path |
| 0x00AEE503 | `n0P0De/` | Known | Filesystem path |
| 0x00AEE58B | `gWi7PQcd/` | Known | Filesystem path |
| 0x00AEE5F7 | `gVi6PScd/` | Known | Filesystem path |
| 0x00AEE815 | `gUi5PSch/` | Known | Filesystem path |
| 0x00AEE881 | `gTi4PQch/` | Known | Filesystem path |
| 0x00AEEB8F | `gYi9PScd/` | Known | Filesystem path |
| 0x00AEEBFB | `gXi8PQcd/` | Known | Filesystem path |
| 0x00AEEF87 | `10D3P`/` | Known | Filesystem path |
| 0x00AF3330 | `/1An->`/4` | Known | Filesystem path |
| 0x00AF6BFA | `d/wwwwt` | Known | Filesystem path |
| 0x00AFAEF0 | `0B0B0B0B0B0B0B0Bg>`/` | Known | Filesystem path |
| 0x00AFBCFA | `0B0B0B0B0B0B0B0B0B0B0B0B0B0B0B0B0B0B0B0B`/` | Known | Filesystem path |
| 0x00AFFD7C | `!#%')+-/13579;=?` | Known | Filesystem path |
| 0x00AFFDAC | ` !"#$%&'()*+,-./0123456789:;<=>?` | Known | Filesystem path |
| 0x00B08712 | `b/ANYDa` | Known | Filesystem path |
| 0x00B08944 | ``/TNVEa` | Known | Filesystem path |
| 0x00B08B34 | `0D0R0X0R0X0R0X0R0X0R0X0R0Xb/0` | Known | Filesystem path |
| 0x00B08D80 | `e/UEUQa` | Known | Filesystem path |
| 0x00B092E8 | `d/AMESa` | Known | Filesystem path |
| 0x00B094FE | ``/RSIHa` | Known | Filesystem path |
| 0x00B09FF2 | `a/ffff1y` | Known | Filesystem path |
| 0x00B0A03A | `a/wwww1y` | Known | Filesystem path |
| 0x00B0A716 | `d/EMITa` | Known | Filesystem path |
| 0x00B14D0D | `o$>P@QA&/ ` | Known | Filesystem path |
| 0x00B14F60 | `P@QA / ` | Known | Filesystem path |
| 0x00B1522F | `C"fk>!/^` | Known | Filesystem path |
| 0x00B1C1A8 | `0D0D0D`/` | Known | Filesystem path |
| 0x00B1C4E2 | `0J0D"/d` | Known | Filesystem path |
| 0x00B243E8 | `>?^@_A./` | Known | Filesystem path |
| 0x00B25033 | `;0K0D /`9` | Known | Filesystem path |
| 0x00B27AFB | `#'+/26:?CGMQVZ_dhmrv{` | Known | Filesystem path |
| 0x00B27BD9 | `"&*/37<@EINRW\aejotx}` | Known | Filesystem path |
| 0x00B2AE5C | `!#%')+-/13579;=?X` | Known | Filesystem path |
| 0x00B315A0 | `/mfs/vlls/` | Known | Filesystem path |
| 0x00B3A563 | `q/qKqfq` | Known | Filesystem path |
| 0x00B3A5C5 | `v/vHvavzv` | Known | Filesystem path |
| 0x00B3BECB | `R/RVR\|R` | Known | Filesystem path |
| 0x00B3C0AF | `p/pGp_pwp` | Known | Filesystem path |
| 0x00B3CFC7 | `pwp_pGp/p` | Known | Filesystem path |
| 0x00B3D1AD | `R\|RVR/R` | Known | Filesystem path |
| 0x00B40CD3 | `n0Dd/\|%` | Known | Filesystem path |
| 0x00B40CED | `L n0Da/` | Known | Filesystem path |
| 0x00B40CFB | `D0n0Db/` | Known | Filesystem path |
| 0x00B44190 | `))/113//+++(` | Known | Filesystem path |
| 0x00B4B255 | `."1PN#/` | Known | Filesystem path |
| 0x00B74D71 | `f??R@!/` | Known | Filesystem path |
| 0x00B751AE | `'.5<=6/7>?` | Known | Filesystem path |
| 0x00B751EA | `"#()01*+$%&',-./2389:;4567<=>?` | Known | Filesystem path |
| 0x00B75240 | `&.6>'/7?` | Known | Filesystem path |
| 0x00B864E0 | `/mfs/temp.mid` | Known | Filesystem path |
| 0x00B8A94C | `d/IDIMt(f` | Known | Filesystem path |
| 0x00B8B9EC | `e/9AMW$.` | Known | Filesystem path |
| 0x00B96659 | `Fh`10"/` | Known | Filesystem path |
| 0x00B9698D | `gh`10 /` | Known | Filesystem path |
| 0x00C3CCDE | `*/?n{'d8` | Known | Filesystem path |
| 0x00C4114D | `MS#%;O/` | Known | Filesystem path |
| 0x00C46CA0 | `O/R/e0T` | Known | Filesystem path |
| 0x00C4918A | `?VEU/vvh` | Known | Filesystem path |
| 0x00C491E7 | `8[/sxHoD` | Known | Filesystem path |
| 0x00C4B8F9 | `m/;X\|ZZ` | Known | Filesystem path |
| 0x00C4FA12 | `,-Mew/L:-IA` | Known | Filesystem path |
| 0x00C53206 | `^/[Zh\|~` | Known | Filesystem path |
| 0x00C5B391 | `"d/cX!c` | Known | Filesystem path |
| 0x00C5E900 | `/VN+>p4u` | Known | Filesystem path |
| 0x00C60DCC | `zi/IlN@` | Known | Filesystem path |
| 0x00C61129 | `$)[P/lT` | Known | Filesystem path |
| 0x00C613D4 | `Ij56zsM/` | Known | Filesystem path |
| 0x00C63413 | `=-cY/A^g` | Known | Filesystem path |
| 0x00C70824 | `Ky<V/jG` | Known | Filesystem path |
| 0x00C79AB6 | `x27.Yl/` | Known | Filesystem path |
| 0x00C7C62F | `~,w/#U1` | Known | Filesystem path |
| 0x00C7EF58 | `O\|/SF93` | Known | Filesystem path |
| 0x00C89501 | `+w%/o1o\B` | Known | Filesystem path |
| 0x00C91DBB | `<a/lv22` | Known | Filesystem path |
| 0x00C99EB2 | `/12xt@t-` | Known | Filesystem path |
| 0x00C9A1B1 | `uAb4/Ee]` | Known | Filesystem path |
| 0x00C9A771 | `/-o;i9K` | Known | Filesystem path |
| 0x00C9C1AC | `;Mi'_B/` | Known | Filesystem path |
| 0x00CA21B9 | `/a7NY7rD` | Known | Filesystem path |
| 0x00CAF899 | `\|/Wyw^pLa` | Known | Filesystem path |
| 0x00CB350B | `/ram!9YFr` | Known | Filesystem path |
| 0x00CBBD12 | `9\Q/1N)` | Known | Filesystem path |
| 0x00CBC0A1 | `/.48?-a` | Known | Filesystem path |
| 0x00CBE6A2 | `FH/**R1*` | Known | Filesystem path |
| 0x00CC4B65 | `PLpe/g~e` | Known | Filesystem path |
| 0x00CC7AA0 | `/5>oG9$` | Known | Filesystem path |
| 0x00CC91DB | `Ln=8*j/E1` | Known | Filesystem path |
| 0x00CCEE67 | `/v%PH*[T` | Known | Filesystem path |
| 0x00CDA930 | `m72Y"?,sVSv/u` | Known | Filesystem path |
| 0x00CDD5C6 | `kU!R/h:` | Known | Filesystem path |
| 0x00CDF09D | `Q/{sob"3xB` | Known | Filesystem path |
| 0x00CE4C27 | `4T&Mf/g9` | Known | Filesystem path |
| 0x00CE596D | `>ne"(]</` | Known | Filesystem path |
| 0x00CE60EB | `/9OeL"H_` | Known | Filesystem path |
| 0x00CECF8B | `#~CBT/(o` | Known | Filesystem path |
| 0x00CFB8B0 | `t3>HBA/#` | Known | Filesystem path |
| 0x00CFC97A | `+l]kO&j/` | Known | Filesystem path |
| 0x00D00A54 | `/""T~pT` | Known | Filesystem path |
| 0x00D0E879 | `T~#-yt/O` | Known | Filesystem path |
| 0x00D171C2 | `&X/6qN_` | Known | Filesystem path |
| 0x00D1A02E | `Xu<)a/U` | Known | Filesystem path |
| 0x00D1CC1E | `)yi9s\|/` | Known | Filesystem path |
| 0x00D271E5 | `/8Mt4(P` | Known | Filesystem path |
| 0x00D3109D | `.wt:iu@/` | Known | Filesystem path |
| 0x00D31A66 | `/t~S\MK` | Known | Filesystem path |
| 0x00D31B8F | `k{JR7i/` | Known | Filesystem path |
| 0x00D32B70 | `/N1 sJ-W` | Known | Filesystem path |
| 0x00D36DF9 | `Br/ay7:,` | Known | Filesystem path |
| 0x00D3B855 | `!RRZ~&nI/4/` | Known | Filesystem path |

---

## 10. Nike+/Fitness

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00216B88 | `Channel TrainerFileAccess` | Known | Nike+ integration |

---

## 11. Video Playback

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017B928 | `VideoCodecs` | Known | Video playback |
| 0x0017BA0C | `H.264LC` | Known | Video playback |
| 0x004B7E34 | `TV Out` | Known | Video playback |
| 0x00B53EE0 | `H.264 Video Decoder` | Known | Video playback |
| 0x00B790A0 | `MPEG-4 video decoder` | Known | Video playback |

---

## 12. Binary Structure (for Ghidra/IDA)

| Parameter | Value |
|-----------|-------|
| **Architecture** | ARM (32-bit, little-endian) |
| **Processor** | ARM7TDMI (ARMv4T), dual-core + video DSP |
| **Base Address** | 0x00000000 |
| **Entry Point** | 0x00000000 |
| **Endianness** | Little-Endian |
| **File Size** | 13,903,872 bytes |

