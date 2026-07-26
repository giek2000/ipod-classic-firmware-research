# iPod 5th Generation (Video) - RetailOS 13.1.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 13.1.3 |
| **IPSW** | iPod_13.1.3.ipsw |
| **Device** | iPod Video 5th Generation (2005, 30/60GB) |
| **Binary Size** | 13,893,632 bytes (13.25 MB) |
| **ARM Code Start** | 0x0 |
| **ARM Code Size** | 13,893,632 bytes |
| **Total Strings (>=6)** | 30,182 |
| **Function Prologues** | 12,890 |
| **SoC** | PortalPlayer PP5021C / Broadcom BCM2722 |
| **Architecture** | ARM7TDMI (ARMv4T), dual-core + video DSP |
| **Encrypted** | No |
| **Decryption Method** | Extract from IPSW (unencrypted) |
| **SHA-256** | `182632b68c54103693ef8cfe8b248fe23f8b733973ee3b8a756bc1c3f9aa2c88` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00106C80 | `Root Hub Driver Internal Error unused case in hub handl...` | Hidden | Undocumented UI |
| 0x0016EF04 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x00215CF4 | `Channel Reserved` | Hidden | Logging/Telemetry |
| 0x00215D08 | `Channel AppBoot` | Hidden | Logging/Telemetry |
| 0x00215D18 | `Channel BufferedSongReading` | Hidden | Logging/Telemetry |
| 0x00215D34 | `Channel PrefsWriting` | Hidden | Logging/Telemetry |
| 0x00215D4C | `Channel GeneralUserExperience` | Hidden | Logging/Telemetry |
| 0x00215D6C | `Channel PlayFromDisk` | Hidden | Logging/Telemetry |
| 0x00215D84 | `Channel CacheSpinupDrive` | Hidden | Logging/Telemetry |
| 0x00215DA0 | `Channel TestLogging` | Hidden | Logging/Telemetry |
| 0x00215DB4 | `Channel AppFileLoading` | Hidden | Logging/Telemetry |
| 0x00215DCC | `Channel VCardReading` | Hidden | Logging/Telemetry |
| 0x00215DE4 | `Channel LongSongScanning` | Hidden | Logging/Telemetry |
| 0x00215E58 | `Channel VoiceRecording` | Hidden | Logging/Telemetry |
| 0x00215E70 | `Channel VoiceRecordingNewFileSegment` | Hidden | Logging/Telemetry |
| 0x00215E98 | `Channel PhotoBrowse` | Hidden | Logging/Telemetry |
| 0x00215EAC | `Channel PhotoImporting` | Hidden | Logging/Telemetry |
| 0x00215EC4 | `Channel Notes` | Hidden | Logging/Telemetry |
| 0x00215ED4 | `Channel PhotoFileManagement` | Hidden | Logging/Telemetry |
| 0x00215EF0 | `Channel DiskModeChannel` | Hidden | Logging/Telemetry |
| 0x00215F08 | `Channel FirewireChannel` | Hidden | Logging/Telemetry |
| 0x00215F20 | `Channel USBChannel` | Hidden | Logging/Telemetry |
| 0x00215F34 | `Channel UnitTests` | Hidden | Hidden Test |
| 0x00215F48 | `Channel FreeSpaceCache` | Hidden | Logging/Telemetry |
| 0x00215FC0 | `Channel OnTheGoFileMgmt` | Hidden | Logging/Telemetry |
| 0x00215FD8 | `Channel SlideShow` | Hidden | Logging/Telemetry |
| 0x00215FEC | `Channel ImageCache` | Hidden | Logging/Telemetry |
| 0x00216000 | `Channel AlbumArtReading` | Hidden | Logging/Telemetry |
| 0x00216018 | `Channel Video` | Hidden | Logging/Telemetry |
| 0x00216028 | `Channel DiskImage` | Hidden | Logging/Telemetry |
| 0x0021603C | `Channel ResourceAccess` | Hidden | Logging/Telemetry |
| 0x00216054 | `Channel VideoCoreBoot` | Hidden | Logging/Telemetry |
| 0x0021606C | `Channel DiskFormatConvert` | Hidden | Logging/Telemetry |
| 0x00216088 | `Channel StreamCacheAddFile` | Hidden | Logging/Telemetry |
| 0x002160A4 | `Channel FontFileAccess` | Hidden | Logging/Telemetry |
| 0x002160BC | `Channel ScreenLock` | Hidden | Logging/Telemetry |
| 0x00216128 | `Channel DiskReaderTask` | Hidden | Logging/Telemetry |
| 0x00216140 | `Channel ProfilerAccess` | Hidden | Logging/Telemetry |
| 0x00216158 | `Channel eAppAccess` | Hidden | Logging/Telemetry |
| 0x0021616C | `Channel eAppWriteBackCache` | Hidden | Logging/Telemetry |
| 0x00216188 | `Channel TrainerFileAccess` | Hidden | Logging/Telemetry |
| 0x002161A4 | `Channel IapStorage` | Hidden | Logging/Telemetry |
| 0x002161B8 | `Channel XMLParsing` | Hidden | Logging/Telemetry |
| 0x002161CC | `Channel AudioPrompt` | Hidden | Logging/Telemetry |
| 0x002161E0 | `Channel AudioPromptXML` | Hidden | Logging/Telemetry |
| 0x002161F8 | `Channel StreamCacheSeek` | Hidden | Logging/Telemetry |
| 0x00216210 | `Channel PredictiveCacheSpinup` | Hidden | Logging/Telemetry |
| 0x00216B3C | `iPod Usage Stats` | Hidden | Logging/Telemetry |
| 0x00217CB8 | `Flush Usage Log Data` | Hidden | Logging/Telemetry |

---

## 2. Discovered Features

### EQ Preset

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001FD91C | `Disp seq_num` | EQ Preset | |
| 0x001FD984 | `core_freq_khz` | EQ Preset | |
| 0x00216DF1 | `Total time in deep sleep: %d seconds` | EQ Preset | |
| 0x00216E19 | `Deep sleep was entered %d %s` | EQ Preset | |
| 0x00217D88 | `Enter Deep Sleep` | EQ Preset | |
| 0x00217D9C | `Exit Deep Sleep` | EQ Preset | |
| 0x002B2598 | `Acoustic` | EQ Preset | |
| 0x002B25A4 | `Bass Booster` | EQ Preset | |
| 0x002B25C4 | `Classical` | EQ Preset | |
| 0x002B25E0 | `Electronic` | EQ Preset | |
| 0x002B25F4 | `Hip Hop` | EQ Preset | |
| 0x002B260C | `Loudness` | EQ Preset | |
| 0x002B2618 | `Lounge` | EQ Preset | |
| 0x002B263C | `Small Speakers` | EQ Preset | |
| 0x002B264C | `Spoken Word` | EQ Preset | |
| 0x002B2658 | `Treble Booster` | EQ Preset | |
| 0x002B2678 | `Vocal Booster` | EQ Preset | |
| 0x002B8B9C | `Treble Boost` | EQ Preset | |
| 0x002B8BAC | `Bass Boost` | EQ Preset | |
| 0x002CECAC | `Latina` | EQ Preset | |
| 0x002DC148 | `rique latine` | EQ Preset | |
| 0x003045C8 | `Latino` | EQ Preset | |
| 0x0063021B | `~ BR&B$"` | EQ Preset | |
| 0x0067D03D | `LATIN-1` | EQ Preset | |
| 0x0067D045 | `LATIN1` | EQ Preset | |
| 0x0068066D | `Secure Electronic Transactions` | EQ Preset | |
| 0x00B26AC0 | `hostreq_notify` | EQ Preset | |
| 0x00B26ACF | `hostreq_read_iphoto_block` | EQ Preset | |
| 0x00B26AE9 | `hostreq_rendertext` | EQ Preset | |
| 0x00B850C1 | `Disp seq_num=%u` | EQ Preset | |

### Localization

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x004B53DC | `English` | Localization | |
| 0x004B5414 | `Italiano` | Localization | |
| 0x0067CDD8 | `x-mac-japanese` | Localization | |
| 0x0067D075 | `X-MAC-CHINESETRAD` | Localization | |
| 0x0067D087 | `X-MAC-JAPANESE` | Localization | |
| 0x0067D096 | `MACJAPANESE` | Localization | |
| 0x0067D0B5 | `X-MAC-KOREAN` | Localization | |
| 0x0067D0D7 | `X-MAC-CHINESESIMP` | Localization | |

### Filesystem Path

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00097CD0 | `iPod_Control` | Filesystem Path | |
| 0x00097CFC | `iPod_Control\Device` | Filesystem Path | |
| 0x000A6548 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x000B8BE0 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000BB708 | `iPod_Control\Music\` | Filesystem Path | |
| 0x000BFCF4 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x000E2E94 | `iPod_Control/%s%s%s` | Filesystem Path | |
| 0x000E2EA8 | `iPod_Control/%s/%s%s%s` | Filesystem Path | |
| 0x000EFCC8 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x001AB9AC | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001AC340 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path | |
| 0x001D7564 | `iPod_Control\Device\` | Filesystem Path | |
| 0x0024683C | `iPod_Control/Device` | Filesystem Path | |
| 0x00246850 | `iPod_Control/Device/radio` | Filesystem Path | |
| 0x0067C6BB | `iPod_Control/games_RO/` | Filesystem Path | |
| 0x0067C785 | `iPod_Control/Device/accessories` | Filesystem Path | |
| 0x0067CAB8 | `iPod_Control/iTunes/` | Filesystem Path | |

---

## 3. Known User-Facing Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0007DA90 | `Settings` | Known | User setting |
| 0x001436D4 | `Memory full. %d notes loaded, some notes not loaded. (3...` | Known | UI element |
| 0x0017AE8C | `CanFlashBacklight` | Known | UI element |
| 0x001AFD68 | `KeyRepeatTimer` | Known | UI element |
| 0x001C3600 | `Could not find settingsHandler for pid %d` | Known | User setting |
| 0x001DD070 | `Contextual menu up!` | Known | Menu item |
| 0x002179D0 | `Backlight` | Known | UI element |
| 0x00217C98 | `Backlight On` | Known | UI element |
| 0x00217CA8 | `Backlight Off` | Known | UI element |
| 0x002B7EB8 | `Alarmer` | Known | UI element |
| 0x002B803B | `kke tekstarkiverne til mappen Notes p` | Known | UI element |
| 0x002B864B | ` menuknappen for at annullere.` | Known | Menu item |
| 0x002BA93C | `Nulstil menu` | Known | Menu item |
| 0x002BABA4 | `Hovedmenu` | Known | Menu item |
| 0x002BB1C4 | `Menuer` | Known | Menu item |
| 0x002BF180 | `Extras` | Known | UI element |
| 0x002C0298 | `Contacts` | Known | UI element |
| 0x002C5210 | ` Notes ` | Known | UI element |
| 0x002C7B78 | `Shuffle %s` | Known | UI element |
| 0x002CDEAC | `Calendario` | Known | UI element |
| 0x002CDEB8 | `Calendarios` | Known | UI element |
| 0x002CDF00 | `Alarmas` | Known | UI element |
| 0x002CE089 | `n de usar el iPod como disco y arrastrar los archivos d...` | Known | UI element |
| 0x002CFA00 | `mo sincronizar contactos, calendarios y listas de tarea...` | Known | UI element |
| 0x002CFC74 | `Alarma` | Known | UI element |
| 0x002D0F2C | `Hora alarma` | Known | UI element |
| 0x002D15C8 | `Contraste` | Known | UI element |
| 0x002D46DA | ` sitten tekstitiedostot iPodin Notes-kansioon. Lis` | Known | UI element |
| 0x002D4CC5 | `n. Kumoa painamalla menu-painiketta.` | Known | Menu item |
| 0x002DB2A0 | `Alarmes` | Known | UI element |
| 0x002DB446 | `utilisation comme disque dur puis faites glisser ces fi...` | Known | UI element |
| 0x002DBB85 | `es de la liste. Cliquez le bouton central pour lancer l...` | Known | Menu item |
| 0x002DCE68 | `chargement des contacts.` | Known | UI element |
| 0x002DCF21 | `iPod pour obtenir des instructions sur la synchronisati...` | Known | UI element |
| 0x002DD1A4 | ` Contacts ` | Known | UI element |
| 0x002DD1CC | `Alarme` | Known | UI element |
| 0x002DD1D4 | `Chargement des contacts.` | Known | UI element |
| 0x002DD71C | `Chargement des notes.` | Known | UI element |
| 0x002DE2B7 | `initialiser le menu principal` | Known | Menu item |
| 0x002DE584 | `Menu principal` | Known | Menu item |
| 0x002DE616 | `alarme` | Known | UI element |
| 0x002E1FA2 | ` Notes mapp` | Known | UI element |
| 0x002E86F0 | `Calendari` | Known | UI element |
| 0x002E888C | `Per visualizzare documenti di testo qui, abilita iPod p...` | Known | UI element |
| 0x002EA1A8 | ` di sincronizzazione dei contatti, dei calendari e degl...` | Known | UI element |
| 0x002EB2FC | `Ora legale` | Known | UI element |
| 0x002EB344 | `Ripristina menu principale` | Known | Menu item |
| 0x002EBCA8 | `Contrasto` | Known | UI element |
| 0x002FDBD8 | `Met 'Zoek zenders' zoekt u naar alle beschikbare radioz...` | Known | Menu item |
| 0x002FE090 | `Shuffle nummers` | Known | UI element |
| 0x00300018 | `Shuffle foto's` | Known | UI element |
| 0x00300134 | `Herstel menu` | Known | Menu item |
| 0x00300384 | `Shuffle` | Known | UI element |
| 0x0030038C | `Hoofdmenu` | Known | Menu item |
| 0x00300A0C | `Menu's` | Known | Menu item |
| 0x00300A34 | `Contrast` | Known | UI element |
| 0x00306658 | `Alarmtidspunkt` | Known | UI element |
| 0x00309B94 | `Alarmy` | Known | UI element |
| 0x00309D43 | `gnij te pliki tekstowe do teczki Notes w iPodzie. Szcze...` | Known | UI element |
| 0x0030A398 | ` skanowanie i menu przycisk, by odwo` | Known | Menu item |
| 0x0030C890 | `Wyzeruj menu g` | Known | Menu item |
| 0x0030CB34 | `Menu g` | Known | Menu item |
| 0x0030CBB0 | `Czas alarmu` | Known | UI element |
| 0x003103A7 | `o como disco e, em seguida, arraste ficheiros de texto ...` | Known | UI element |
| 0x00310AAA | `o de menu para cancelar.` | Known | Menu item |
| 0x00312EC8 | `Repor menu pri.` | Known | Menu item |
| 0x0031F86E | `ndning av iPod som extern enhet och drar sedan textfile...` | Known | UI element |
| 0x00322428 | `Alarmtid` | Known | UI element |
| 0x0032591C | `Alarmlar` | Known | UI element |
| 0x00325AF1 | ` iPod'daki Notes klas` | Known | UI element |
| 0x003261A0 | `in Menu d` | Known | Menu item |
| 0x003288E8 | `Alarm Zaman` | Known | UI element |
| 0x0032C65E | ` Menu ` | Known | Menu item |
| 0x00332AEE | ` menu ` | Known | Menu item |
| 0x004B4778 | `Calendar` | Known | UI element |
| 0x004B4784 | `Calendars` | Known | UI element |
| 0x004B47C4 | `Alarms` | Known | UI element |
| 0x004B4BF0 | `Slideshow Settings` | Known | User setting |
| 0x004B4DFC | `Find Stations will scan through all available radio sta...` | Known | Menu item |
| 0x004B51C8 | `Now Playing` | Known | UI element |
| 0x004B52F4 | `Shuffle Songs` | Known | UI element |
| 0x004B53A8 | `Volume Limit` | Known | UI element |
| 0x004B5658 | `New Clock` | Known | UI element |
| 0x004B6564 | `contacts loading.` | Known | UI element |
| 0x004B6864 | `Contacts loading.` | Known | UI element |
| 0x004B6C9C | `Notes loading.` | Known | UI element |
| 0x004B7374 | `Delete This Clock` | Known | UI element |
| 0x004B73E8 | `Sleep Timer` | Known | UI element |
| 0x004B73F4 | `Alarm Clock` | Known | UI element |
| 0x004B7400 | `World Clock` | Known | UI element |
| 0x004B743C | `Video Settings` | Known | User setting |
| 0x004B7570 | `Shuffle Photos` | Known | UI element |
| 0x004B7580 | `Repeat` | Known | UI element |
| 0x004B76B8 | `Reset Main Menu` | Known | Menu item |
| 0x004B784C | `Reset All Settings` | Known | User setting |
| 0x004B7908 | `Backlight Timer` | Known | UI element |
| 0x004B7928 | `Main Menu` | Known | Menu item |
| 0x004B7990 | `Alarm Time` | Known | UI element |
| 0x004B79B0 | `Delete Clock` | Known | UI element |
| 0x004B7A38 | `Radio Settings` | Known | User setting |
| 0x004B7FB4 | `Reset All` | Known | UI element |
| 0x00503A84 | `TCalendarCntlr_Alarm` | Known | UI element |
| 0x00504668 | `To check song links, set the preference NotesOnly to tr...` | Known | UI element |
| 0x005046CC | `The NotesOnly pref can only be set globally in the Pref...` | Known | UI element |
| 0x0050471C | `Warning: Preferences file must be in the Notes folder, ...` | Known | UI element |
| 0x00669AD9 | `Illegal instruction` | Known | UI element |
| 0x00669B07 | `Illegal address` | Known | UI element |
| 0x0067CEFA | `dalarm` | Known | UI element |
| 0x0067CF01 | `valarm` | Known | UI element |
| 0x0067CF53 | `vcalendar` | Known | UI element |
| 0x0067D48C | `NotesOnly` | Known | UI element |
| 0x00AFD0C9 | `backlight` | Known | UI element |
| 0x00B266CE | `audio_playgetclock` | Known | UI element |
| 0x00B266E1 | `audio_playgetendclock` | Known | UI element |
| 0x00B26D89 | `TMT_Retrieve_Clock` | Known | UI element |
| 0x00B8517D | `framenum_checksum=%u` | Known | Menu item |
| 0x00B8E29A | `mp_clock_init` | Known | UI element |
| 0x00B8E2A8 | `mp_clock_destroy` | Known | UI element |
| 0x00B8E2B9 | `mp_clock_reset` | Known | UI element |
| 0x00B8E2C8 | `clock_stop` | Known | UI element |
| 0x00B8E2D3 | `mp_clock_start` | Known | UI element |
| 0x00B8E2E2 | `clock_fetch` | Known | UI element |
| 0x00B8E2EE | `mp_clock_fetch` | Known | UI element |
| 0x00B8E2FD | `mp_clock_fetch2` | Known | UI element |
| 0x00B8E30D | `mp_clock_audio` | Known | UI element |
| 0x00B8E31C | `mp_clock_stc` | Known | UI element |
| 0x00B8E329 | `mp_clock_stop` | Known | UI element |
| 0x00B8E337 | `mp_clock_set` | Known | UI element |

---

## 4. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00098774 | `TrackCacheReadTask` | Known | RTOS task thread |
| 0x000C9120 | `USB Secondary Interrupt Task` | Known | RTOS task thread |
| 0x000E294C | `ICAPTPCameraIOTask` | Known | RTOS task thread |
| 0x0011D1A8 | `USBStatusTask` | Known | RTOS task thread |
| 0x0011D1C4 | `USBTaskTimeTask` | Known | RTOS task thread |
| 0x0014CB44 | `RtcTaskClass` | Known | RTOS task thread |
| 0x0016902C | `VCUpdateTask` | Known | RTOS task thread |
| 0x0016EF04 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x00174E7C | `USBDeviceTask` | Known | RTOS task thread |
| 0x0017B5B4 | `DiskReaderTask` | Known | RTOS task thread |
| 0x00187E00 | `TimerTaskClass` | Known | RTOS task thread |
| 0x0018A9BC | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0018A9D0 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0019DDF8 | `PhotoCopyTask` | Known | RTOS task thread |
| 0x001BE158 | `LoadDataTasks` | Known | RTOS task thread |
| 0x0020FC64 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x00216128 | `Channel DiskReaderTask` | Known | RTOS task thread |
| 0x0026143C | `FirewireTask` | Known | RTOS task thread |
| 0x00261454 | `OptoTask` | Known | RTOS task thread |
| 0x00261464 | `SerialOptoTask` | Known | RTOS task thread |
| 0x00261478 | `BacklightTask` | Known | RTOS task thread |
| 0x0026148C | `CNATask` | Known | RTOS task thread |
| 0x002614AC | `DiskMgrTask` | Known | RTOS task thread |
| 0x002614BC | `HoldSwitchTask` | Known | RTOS task thread |
| 0x002614D0 | `TopPlugTask` | Known | RTOS task thread |
| 0x002614E0 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x002614F4 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x0026150C | `AccessoryDetectTask` | Known | RTOS task thread |
| 0x00261534 | `AlarmTask` | Known | RTOS task thread |
| 0x00261544 | `WatchdogTask` | Known | RTOS task thread |
| 0x002615BC | `USBAudioTask` | Known | RTOS task thread |
| 0x002ADA10 | `HostOSTask` | Known | RTOS task thread |
| 0x002AE5FC | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x002E3438 | `Taskent` | Known | RTOS task thread |
| 0x0050495C | `StreamCopierWriteTask` | Known | RTOS task thread |
| 0x00504974 | `StreamCopierReadTask` | Known | RTOS task thread |
| 0x00504A98 | `VideoDaisyTask` | Known | RTOS task thread |
| 0x00B26C68 | `TCC_Create_Task` | Known | RTOS task thread |
| 0x00B26C78 | `TCC_Current_Task_Pointer` | Known | RTOS task thread |
| 0x00B26CA1 | `TCC_Delete_Task` | Known | RTOS task thread |
| 0x00B26CC0 | `TCC_Reset_Task` | Known | RTOS task thread |
| 0x00B26CE2 | `TCC_Task_Sleep` | Known | RTOS task thread |
| 0x00B26CF1 | `TCC_Terminate_Task` | Known | RTOS task thread |
| 0x00B26D04 | `TCF_Task_Information` | Known | RTOS task thread |

---

## 5. Audio/Codec

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000D95E0 | `RIFFWAVEfmt data"V` | Known | PCM audio format |
| 0x0015DFE4 | `RIFFWAVEfmt data` | Known | PCM audio format |
| 0x00179E88 | `AudioCodecs` | Known | Audio system |
| 0x00179F64 | `Audible` | Known | Audible audiobook format |
| 0x0017AF28 | `VideoCodecs` | Known | Audio system |
| 0x0017B00C | `H.264LC` | Known | Audio system |
| 0x001FD7B8 | `max_decoded_buffer` | Known | Decoder component |
| 0x001FD7CC | `min_decoded_buffer` | Known | Decoder component |
| 0x001FD888 | `total_bytes_decoded` | Known | Decoder component |
| 0x001FD89C | `max_decoded_bytes` | Known | Decoder component |
| 0x001FD93C | `Disp decoded data` | Known | Decoder component |
| 0x001FDE58 | `lost frames - decoder (failed to decode):` | Known | Decoder component |
| 0x001FDE88 | `max decoded buffer:` | Known | Decoder component |
| 0x001FDEA0 | `min decoded buffer:` | Known | Decoder component |
| 0x001FDED0 | `total bytes decoded:` | Known | Decoder component |
| 0x001FDEEC | `max decoded bytes:` | Known | Decoder component |
| 0x002619CF | `@mp3dec_sync` | Known | MP3 codec |
| 0x00262217 | `@mp4_aacdec_sync` | Known | AAC codec |
| 0x002B3B1D | ` Audible v` | Known | Audible audiobook format |
| 0x002B3B6F | ` Audible. Copyright ` | Known | Audible audiobook format |
| 0x002B3B85 | ` 2002 Audible, Inc. V` | Known | Audible audiobook format |
| 0x002B3D62 | ` zvuku MPEG Layer-3 byla poskytnuta spole` | Known | Audio system |
| 0x002B3D8D | `nostmi Fraunhofer IIS a` | Known | Audio system |
| 0x002B9FA8 | `Audible-softwaren i dette produkt bruges i henhold til ...` | Known | Audible audiobook format |
| 0x002BA008 | ` 2002 Audible, Inc. Alle rettigheder forbeholdes.` | Known | Audible audiobook format |
| 0x002BA0FA | `.net codec i dette produkt bruges i henhold til en lice...` | Known | Audio system |
| 0x002BA1A4 | `MPEG Layer-3-lydkodningsteknologi licenseret fra Fraunh...` | Known | Audio system |
| 0x002C0A84 | `Die Audible Software in diesem Produkt wird in Lizenz d...` | Known | Audible audiobook format |
| 0x002C0ADD | ` 2002 Audible, Inc. Alle Rechte vorbehalten.` | Known | Audible audiobook format |
| 0x002C0BC9 | `.net Codec in diesem Produkt wird in Lizenz der VoiceAg...` | Known | Audio system |
| 0x002C0C87 | `r MPEG Layer-3 wurde lizenziert von Fraunhofer IIS und ...` | Known | Audio system |
| 0x002C8B63 | ` Audible ` | Known | Audible audiobook format |
| 0x002C8BC0 | ` Audible. ` | Known | Audible audiobook format |
| 0x002C8BF6 | ` 2002 by Audible, Inc. ` | Known | Audible audiobook format |
| 0x002C8D74 | `.net codec ` | Known | Audio system |
| 0x002C8EBB | ` MPEG Layer-3 ` | Known | Audio system |
| 0x002C8EF9 | ` Fraunhofer IIS ` | Known | Audio system |
| 0x002D01A8 | `El software Audible incluido en este producto se usa ba...` | Known | Audible audiobook format |
| 0x002D0203 | ` 2002 de Audible, Inc. Todos los derechos reservados.` | Known | Audible audiobook format |
| 0x002D03A1 | `n de audio MPEG Layer-3 utilizada bajo licencia de Frau...` | Known | Audio system |
| 0x002D6696 | `n Audiblelta lisensoitua Audible-ohjelmistoa. Copyright...` | Known | Audible audiobook format |
| 0x002D66D0 | ` 2002 Audible, Inc. Kaikki oikeudet pid` | Known | Audible audiobook format |
| 0x002D683C | `MPEG Layer-3 -` | Known | Audio system |
| 0x002D684E | `nen koodaustekniikka on lisensoitu Fraunhofer IIS:lt` | Known | Audio system |
| 0x002DD7D8 | `Le logiciel Audible contenu dans ce produit est utilis` | Known | Audible audiobook format |
| 0x002DD822 | `Audible. Copyright ` | Known | Audible audiobook format |
| 0x002DD837 | ` 2002 par Audible, Inc. Tous droits r` | Known | Audible audiobook format |
| 0x002DD8E8 | `e sous licence de VoiceAge Corporation. Le codec ACELP` | Known | Audio system |
| 0x002DD9BC | `La technologie de codage audio MPEG Layer 3 est utilis` | Known | Audio system |
| 0x002DD9F4 | `e sous licence de Fraunhofer IIS et THOMSON multim` | Known | Audio system |
| 0x002E4242 | ` Audible szoftver az Audible licence alatt van haszn` | Known | Audible audiobook format |
| 0x002E428C | ` 2002, Audible, Inc. Minden jog fenntartva.` | Known | Audible audiobook format |
| 0x002E4381 | `.net codec a VoiceAge Coporation c` | Known | Audio system |
| 0x002E4414 | `Az MPEG Layer-3 hangk` | Known | Audio system |
| 0x002E443C | `gia a Fraunhofer IIS ` | Known | Audio system |
| 0x002E479C | `l alacsony.` | Known | Apple Lossless codec |
| 0x002E9C3C | `La Mecca` | Known | Audio system |
| 0x002EA9C8 | `Il software Audible di questo prodotto ` | Known | Audible audiobook format |
| 0x002EA9F1 | ` utilizzato su licenza da Audible. Copyright ` | Known | Audible audiobook format |
| 0x002EAA20 | ` 2002 di Audible, Inc. Tutti i diritti riservati.` | Known | Audible audiobook format |
| 0x002EAA92 | ` utilizzato su licenza da VoiceAge Corporation. Il code...` | Known | Audio system |
| 0x002EAB68 | `Tecnologia di codifica audio MPEG Layer-3 su licenza da...` | Known | Audio system |
| 0x002F1B62 | `Audible ` | Known | Audible audiobook format |
| 0x002F1BBB | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x002F1D70 | `MPEG Layer-3 ` | Known | Audio system |
| 0x002F1DBC | `Fraunhofer IIS ` | Known | Audio system |
| 0x002F8C0E | ` Audible` | Known | Audible audiobook format |
| 0x002F8D42 | `.net codec` | Known | Audio system |
| 0x002F8E03 | ` Fraunhofer IIS` | Known | Audio system |
| 0x002FF72C | `De Audible-software in dit product wordt gebruikt in li...` | Known | Audible audiobook format |
| 0x002FF783 | ` 2002 Audible, Inc. Alle rechten voorbehouden.` | Known | Audible audiobook format |
| 0x002FF874 | `.net-codec in dit product wordt gebruikt in licentie va...` | Known | Audio system |
| 0x002FF910 | `Technologie voor codering van MPEG Layer-3-audio in lic...` | Known | Audio system |
| 0x00305A5C | `Audible-programvaren i dette produktet brukes under lis...` | Known | Audible audiobook format |
| 0x00305AB0 | ` 2002 by Audible, Inc. Alle rettigheter forbeholdes.` | Known | Audible audiobook format |
| 0x00305C2C | `Lydkodingsteknologien MPEG Layer-3 er lisensiert fra Fr...` | Known | Audio system |
| 0x0030BE64 | `Oprogramowanie Audible w tym produkcie jest wykorzystyw...` | Known | Audible audiobook format |
| 0x0030BED0 | ` 2002 Audible, Inc. Wszystkie prawa zastrze` | Known | Audible audiobook format |
| 0x0030C078 | `Technologia kodowania audio MPEG Layer-3 licencjonowana...` | Known | Audio system |
| 0x0031244C | `O software Audible ` | Known | Audible audiobook format |
| 0x00312482 | `a da Audible. Copyright ` | Known | Audible audiobook format |
| 0x0031249C | ` 2002 da Audible, Inc. Reservados todos os direitos.` | Known | Audible audiobook format |
| 0x0031255D | `a da VoiceAge Corporation. O codec ACELP` | Known | Audio system |
| 0x0031263E | `udio MPEG Layer-3 licenciada pela Fraunhofer IIS e THOM...` | Known | Audio system |
| 0x0031A734 | `MPEG Layer-3: ` | Known | Audio system |
| 0x003217AC | `Audible-programvaran anv` | Known | Audible audiobook format |
| 0x003217DB | `n Audible. Copyright ` | Known | Audible audiobook format |
| 0x003217F2 | ` 2002 Audible, Inc. Alla r` | Known | Audible audiobook format |
| 0x0032198C | `Ljudkodningstekniken MPEG Layer-3 ` | Known | Audio system |
| 0x003219C2 | `n Fraunhofer IIS och THOMSON multimedia.` | Known | Audio system |
| 0x00327BC0 | `ndeki Audible yaz` | Known | Audible audiobook format |
| 0x00327BD9 | ` Audible lisans` | Known | Audible audiobook format |
| 0x00327C0E | ` 2002, Audible, Inc. T` | Known | Audible audiobook format |
| 0x00327D09 | `.net codec'i VoiceAge Corporation lisans` | Known | Audio system |
| 0x00327D98 | `MPEG Layer-3 ses kodlama teknolojisi Fraunhofer IIS ve ...` | Known | Audio system |
| 0x004B6D7C | `The Audible software in this product is used under lice...` | Known | Audible audiobook format |
| 0x004B6EB5 | `.net codec in this product is used under license from V...` | Known | Audio system |
| 0x004B6F48 | `MPEG Layer-3 audio coding technology licensed from Frau...` | Known | Audio system |
| 0x004B7434 | `TV Out` | Known | Audio system |
| 0x0065F0D1 | ``0aLaCfDf` | Known | Apple Lossless codec |
| 0x00669E14 | `21SoundEffectDescriptor` | Known | Audio system |
| 0x0067D274 | `&Aacute` | Known | AAC codec |
| 0x0067D29C | `&aacute` | Known | AAC codec |
| 0x0067FA8E | `msCodeCom` | Known | Audio system |
| 0x006805E3 | `aaControls` | Known | AAC codec |
| 0x00B26A4F | `gencmd_decode_fourcc` | Known | Decoder component |
| 0x00B26A64 | `gencmd_decode_int` | Known | Decoder component |
| 0x00B37C60 | `AACDEC  VLL ` | Known | AAC codec |
| 0x00B413A8 | `MPEG4 AAC LC Decoder` | Known | AAC codec |
| 0x00B4380D | `AACDecoderGetMem` | Known | AAC codec |
| 0x00B4381E | `AACDecoderInit` | Known | AAC codec |
| 0x00B4382D | `AACDecoderGetConfig` | Known | AAC codec |
| 0x00B43841 | `AACDecoderSetConfig` | Known | AAC codec |
| 0x00B43855 | `AACHeaderDecode` | Known | AAC codec |
| 0x00B43865 | `AACDecode` | Known | AAC codec |
| 0x00B4386F | `AACDecoderInit_Ittiam` | Known | AAC codec |
| 0x00B43885 | `AACDecoderGetConfig_Ittiam` | Known | AAC codec |
| 0x00B438A0 | `AACDecoderSetConfig_Ittiam` | Known | AAC codec |
| 0x00B438BB | `AACHeaderDecode_Ittiam` | Known | AAC codec |
| 0x00B438D2 | `AACDecode_Ittiam` | Known | AAC codec |
| 0x00B438FA | `get_aac_dec_func_table` | Known | AAC codec |
| 0x00B43992 | `aac_initbits` | Known | AAC codec |
| 0x00B4399F | `aac_get_processed_bits` | Known | AAC codec |
| 0x00B439B6 | `aac_byte_align` | Known | AAC codec |
| 0x00B43BA6 | `is_decode` | Known | Decoder component |
| 0x00B43BB0 | `can_decode_objType` | Known | Decoder component |
| 0x00B43BD7 | `ms_decode` | Known | Decoder component |
| 0x00B43C1E | `pns_decode` | Known | Decoder component |
| 0x00B43C29 | `pulse_decode` | Known | Decoder component |
| 0x00B43C56 | `tns_decode_frame` | Known | Decoder component |
| 0x00B522E0 | `H.264 Video Decoder` | Known | Decoder component |
| 0x00B5B400 | `H264InitDecoder` | Known | Decoder component |
| 0x00B5B410 | `init_decoder` | Known | Decoder component |
| 0x00B5B41D | `H264DecodeFrame` | Known | Decoder component |
| 0x00B5B42D | `H264ReleaseDecoder` | Known | Decoder component |
| 0x00B5B440 | `shutdown_decoder` | Known | Decoder component |
| 0x00B5B5A8 | `decode_one_frame` | Known | Decoder component |
| 0x00B5B92D | `h264_refstripe_prepare_decode` | Known | Decoder component |
| 0x00B5B94B | `h264_refstripe_finished_decode` | Known | Decoder component |
| 0x00B5C0FF | `mvpairdecode_table` | Known | Decoder component |
| 0x00B5C112 | `mvpairdecodelen_table` | Known | Decoder component |
| 0x00B5CAC1 | `h264_writestripe_prepare_decode` | Known | Decoder component |
| 0x00B774A0 | `MPEG-4 video decoder` | Known | Decoder component |
| 0x00B7F86A | `UBVInitDecoder` | Known | Decoder component |
| 0x00B7F879 | `UBVDecodeFrame` | Known | Decoder component |
| 0x00B7F888 | `UBVReleaseDecoder` | Known | Decoder component |
| 0x00B7F993 | `macroblockdecode` | Known | Decoder component |
| 0x00B7FA10 | `vc_ClipIquantMPEG4` | Known | Audio system |
| 0x00B7FA23 | `vc_MPEG4InterIQuant` | Known | Audio system |
| 0x00B7FC64 | `vc_MPEG4getDC` | Known | Audio system |
| 0x00B7FC8F | `MPEG4CopyMBs` | Known | Audio system |
| 0x00B7FC9C | `DecodeFrame` | Known | Decoder component |
| 0x00B7FCA8 | `MPEG4DecodeFrame` | Known | Decoder component |
| 0x00B7FCB9 | `H263DecodeFrame` | Known | Decoder component |
| 0x00B7FCC9 | `MPEG4DecodeEnhancementFrame` | Known | Decoder component |
| 0x00B7FCE5 | `MPEG4InitParams` | Known | Audio system |
| 0x00B802F5 | `H263DecodeMotionVectors` | Known | Decoder component |
| 0x00B8030D | `H263DecodeIPic` | Known | Decoder component |
| 0x00B80346 | `H263DecodePPic` | Known | Decoder component |
| 0x00B80407 | `pmacroblockdecode` | Known | Decoder component |
| 0x00B80513 | `MPEG4PPictureSave6BlocksD` | Known | Audio system |
| 0x00B80591 | `MPEG4GetCBPY` | Known | Audio system |
| 0x00B8059E | `MPEG4MCBPCIntra` | Known | Audio system |
| 0x00B805AE | `MPEG4MCBPCInter` | Known | Audio system |
| 0x00B805BE | `MPEG4InitMVD` | Known | Audio system |
| 0x00B805CB | `MPEG4GetTMNMV` | Known | Audio system |
| 0x00B805D9 | `MPEG4AddPMV` | Known | Audio system |
| 0x00B805E5 | `MPEG4GetMotionVectorsF` | Known | Audio system |
| 0x00B805FC | `MPEG4GetMotionVectorData` | Known | Audio system |
| 0x00B80615 | `MPEG4FindPos4MVD` | Known | Audio system |
| 0x00B80626 | `MPEG4ClipMV` | Known | Audio system |
| 0x00B80632 | `MPEG4PredictIntraBlock` | Known | Audio system |
| 0x00B80649 | `MPEG4InitRowBlocksD` | Known | Audio system |
| 0x00B8065D | `MPEG4IntraAdvancedPredictionDecodeD` | Known | Decoder component |
| 0x00B80681 | `MPEG4DecodeIPic` | Known | Decoder component |
| 0x00B80691 | `MPEG4DecodeDataPartitionedIPic` | Known | Decoder component |
| 0x00B807AD | `MPEG4_INTRA_MP4BDEC_UBV_DCT3D0` | Known | Audio system |
| 0x00B807CC | `MPEG4_INTRA_MP4BDEC_UBV_DCT3D1` | Known | Audio system |
| 0x00B807EB | `MPEG4_INTRA_MP4BDEC_UBV_DCT3D2` | Known | Audio system |
| 0x00B8081E | `MPEG4GetInterBlock` | Known | Audio system |
| 0x00B80845 | `MPEG4GetIntraBlock` | Known | Audio system |
| 0x00B80858 | `MPEG4RvlcDecTCOEF` | Known | Audio system |
| 0x00B8086A | `MPEG4GetIntraBlockRVLC` | Known | Audio system |
| 0x00B80881 | `MPEG4GetInterBlockRVLC` | Known | Audio system |
| 0x00B80898 | `MPEG4ParseVOLHeader` | Known | Audio system |
| 0x00B808AC | `MPEG4FlushUserData` | Known | Audio system |
| 0x00B808BF | `MPEG4FindStartCode` | Known | Audio system |
| 0x00B808D2 | `MPEG4ParseVOPHeader` | Known | Audio system |
| 0x00B808E6 | `MPEG4CheckMotionMarker` | Known | Audio system |
| 0x00B808FD | `MPEG4DecodePPic` | Known | Decoder component |
| 0x00B8090D | `MPEG4DecodeInterTextureMacroblock` | Known | Decoder component |
| 0x00B8092F | `MPEG4DecodeDataPartitionedPPic` | Known | Decoder component |
| 0x00B8094E | `MPEG4ReadIMacroBlocks` | Known | Audio system |
| 0x00B80964 | `MPEG4ReadDQuant` | Known | Audio system |
| 0x00B80974 | `MPEG4ReadPMacroBlocks` | Known | Audio system |
| 0x00B8098A | `MPEG4DataPartitionReadIMacroBlocks` | Known | Audio system |
| 0x00B809AD | `MPEG4DataPartitionReadPMacroBlocks` | Known | Audio system |
| 0x00B809D0 | `MPEG4DataPartitionReadPMacroBlocks2` | Known | Audio system |
| 0x00B809F4 | `MPEG4ReadVideoPacketHeader` | Known | Audio system |
| 0x00B80A16 | `InitDecoder` | Known | Decoder component |
| 0x00B80A31 | `MPEG4InitMemory` | Known | Audio system |
| 0x00B80A41 | `InitEnhancementDecoder` | Known | Decoder component |
| 0x00B80A58 | `ReleaseDecoder` | Known | Decoder component |
| 0x00B80A67 | `H263ReleaseH263Decoder` | Known | Decoder component |
| 0x00B80A7E | `MPEG4ReleaseMPEG4Decoder` | Known | Decoder component |
| 0x00B80A97 | `ReleaseEnhancementDecoder` | Known | Decoder component |
| 0x00B80F9A | `MPEG4ShowBitsAlignedD` | Known | Audio system |
| 0x00B80FB0 | `MPEG4AlignInput` | Known | Audio system |
| 0x00B80FC0 | `MPEG4PeekNextStartCode` | Known | Audio system |
| 0x00B8105E | `Decoders` | Known | Decoder component |
| 0x00B84C10 | `forbid_decoder_panic=` | Known | Decoder component |
| 0x00B85165 | `max_decoded_bytes=%d` | Known | Decoder component |
| 0x00B85195 | `max_decoded_buffer=%u` | Known | Decoder component |
| 0x00B851AD | `min_decoded_buffer=%u` | Known | Decoder component |
| 0x00B851C5 | `total_bytes_decoded=%u` | Known | Decoder component |
| 0x00B851DD | `Disp decoded_data=%08x` | Known | Decoder component |
| 0x00B8B28F | `no audio decoder` | Known | Decoder component |
| 0x00B8B2A0 | `no video decoder` | Known | Decoder component |
| 0x00B8DF9D | `codec_string` | Known | Audio system |
| 0x00B8DFAA | `codec_name` | Known | Audio system |
| 0x00B9CBF1 | `decode_int` | Known | Decoder component |

---

## 6. Storage/Hardware

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00003A00 | `!ATAsoso` | Known | ATA/disk interface |
| 0x00003A28 | `!ATAcrsr` | Known | ATA/disk interface |
| 0x00003A50 | `!ATAdpua` | Known | ATA/disk interface |
| 0x00004334 | `!ATAebih` | Known | ATA/disk interface |
| 0x00005589 | `diskmode` | Known | Hardware interface |
| 0x00005592 | `diskscan` | Known | Hardware interface |
| 0x0007D718 | `Metadata` | Known | ATA/disk interface |
| 0x00097F74 | `Photo Database` | Known | ATA/disk interface |
| 0x000A968C | `atadmrts` | Known | ATA/disk interface |
| 0x000B3B4C | `atadmhbddbhmmhsd` | Known | ATA/disk interface |
| 0x000B3E28 | `atadmhfddfhmmhsd\|@-` | Known | ATA/disk interface |
| 0x000B8C04 | `Photos\Photo Database` | Known | ATA/disk interface |
| 0x000BFE8C | `nutiatad` | Known | ATA/disk interface |
| 0x000C0154 | `atadImage DB Temp` | Known | ATA/disk interface |
| 0x000CA0AC | `]ih[!ATA` | Known | ATA/disk interface |
| 0x000E9C8C | `atadmhpo0@-` | Known | ATA/disk interface |
| 0x000EDAE4 | `data abort` | Known | ATA/disk interface |
| 0x000EFCE8 | `atadmhdp` | Known | ATA/disk interface |
| 0x001424EC | `Bad meta data, name not found. (23)` | Known | ATA/disk interface |
| 0x00142514 | `Bad meta data, name termination quote not found. (24)` | Known | ATA/disk interface |
| 0x00142550 | `Bad meta data, content not found. (25)` | Known | ATA/disk interface |
| 0x00142578 | `Bad meta data, content termination quote not found. (26...` | Known | ATA/disk interface |
| 0x0014321C | `Bad meta data, encoding not found. (8)` | Known | ATA/disk interface |
| 0x00143244 | `Bad meta data, encoding termination quote not found. (9...` | Known | ATA/disk interface |
| 0x00143280 | `Bad meta data, unknown encoding. (17)` | Known | ATA/disk interface |
| 0x00143624 | `Bad meta data, name not a recognized preference. (27)` | Known | ATA/disk interface |
| 0x00143D30 | `Bad data. (32)` | Known | ATA/disk interface |
| 0x0017706D | `lyrdata` | Known | ATA/disk interface |
| 0x00179EE4 | `MaximumDataRate` | Known | ATA/disk interface |
| 0x0017A394 | `FireWireGUID` | Known | FireWire interface |
| 0x0017A3A4 | `FireWireVersion` | Known | FireWire interface |
| 0x0017A990 | `FireWire` | Known | FireWire interface |
| 0x0017AE38 | `ForcedDiskMode` | Known | Hardware interface |
| 0x0017AE58 | `CorruptDataPartition` | Known | ATA/disk interface |
| 0x00188624 | `USB MSC` | Known | USB interface |
| 0x0019B0BC | `23iUPhoto Database` | Known | ATA/disk interface |
| 0x0019D724 | `Photo Import Database` | Known | ATA/disk interface |
| 0x001CBB38 | `spiral` | Known | Hardware interface |
| 0x00215D6C | `Channel PlayFromDisk` | Known | Hardware interface |
| 0x00215D84 | `Channel CacheSpinupDrive` | Known | Hardware interface |
| 0x00215EF0 | `Channel DiskModeChannel` | Known | Hardware interface |
| 0x00215F08 | `Channel FirewireChannel` | Known | FireWire interface |
| 0x00216028 | `Channel DiskImage` | Known | Hardware interface |
| 0x0021606C | `Channel DiskFormatConvert` | Known | Hardware interface |
| 0x00216210 | `Channel PredictiveCacheSpinup` | Known | Hardware interface |
| 0x00216238 | `Unknown Disk Channel` | Known | Hardware interface |
| 0x00216968 | `Disk Activity` | Known | Hardware interface |
| 0x00216979 | `Total time the disk was running in the app: %d seconds` | Known | Hardware interface |
| 0x00216A21 | `The disk was turned on %d %s` | Known | Hardware interface |
| 0x002171D1 | `Music database size: %d KB` | Known | ATA/disk interface |
| 0x002171F1 | `Music database num songs: %d` | Known | ATA/disk interface |
| 0x00217211 | `Photo database size: %d KB` | Known | ATA/disk interface |
| 0x00217231 | `Photo database num photos: %d` | Known | ATA/disk interface |
| 0x00217251 | `Album art database size: %d KB` | Known | ATA/disk interface |
| 0x00217BCC | `Disk Spinup` | Known | Hardware interface |
| 0x00217BD8 | `Disk Spindown` | Known | Hardware interface |
| 0x00217BE8 | `Disk Obtain Access` | Known | Hardware interface |
| 0x00217BFC | `Disk Release Access` | Known | Hardware interface |
| 0x00217CB8 | `Flush Usage Log Data` | Known | ATA/disk interface |
| 0x00217D40 | `Enter Disk Mode` | Known | Hardware interface |
| 0x00217D50 | `Exit Disk Mode` | Known | Hardware interface |
| 0x00217DCC | `Music Database Size` | Known | ATA/disk interface |
| 0x00217DE0 | `Photo Database Size` | Known | ATA/disk interface |
| 0x00217DF4 | `Artwork Database Size` | Known | ATA/disk interface |
| 0x00261C10 | `[CDATA[` | Known | ATA/disk interface |
| 0x0026BDE4 | `MEMDISK` | Known | Hardware interface |
| 0x002AD24B | `glBufferData` | Known | ATA/disk interface |
| 0x002AD258 | `glBufferSubData` | Known | ATA/disk interface |
| 0x002ADA48 | `gamedata_RW` | Known | ATA/disk interface |
| 0x002ADA64 | `gamedata_ShareRW` | Known | ATA/disk interface |
| 0x002B1A60 | `e v diskov` | Known | Hardware interface |
| 0x002B2120 | `Data RDS nenalezena` | Known | ATA/disk interface |
| 0x002B2D78 | `Kalkata` | Known | ATA/disk interface |
| 0x002B311C | `im disku` | Known | Hardware interface |
| 0x002B3137 | `es FireWire nen` | Known | FireWire interface |
| 0x002B32DB | `dat a zobrazovat data importovan` | Known | ATA/disk interface |
| 0x002B4E59 | `e na disku nen` | Known | Hardware interface |
| 0x002B50B0 | `FireWire p` | Known | FireWire interface |
| 0x002B8011 | ` brug af iPod som ekstern disk til og tr` | Known | Hardware interface |
| 0x002B8358 | `Videospillelister` | Known | Hardware interface |
| 0x002B8680 | `Ingen RDS-data fundet` | Known | ATA/disk interface |
| 0x002B86A0 | ` Afspil for at lytte til radio` | Known | Hardware interface |
| 0x002B86C8 | ` Afspil for at slukke radioen` | Known | Hardware interface |
| 0x002B8820 | `Spiller nu` | Known | Hardware interface |
| 0x002B8904 | `Spillelister` | Known | Hardware interface |
| 0x002B8930 | `Genoptag spil` | Known | Hardware interface |
| 0x002B899C | `Ved afspilning` | Known | Hardware interface |
| 0x002B9260 | `Kolkata (Calcutta)` | Known | ATA/disk interface |
| 0x002B93A4 | `Ulaanbaatar` | Known | ATA/disk interface |
| 0x002B9494 | `Slet spilleliste` | Known | Hardware interface |
| 0x002B94A8 | `Arkiver spilleliste` | Known | Hardware interface |
| 0x002B9568 | `Ny spilleliste %lu` | Known | Hardware interface |
| 0x002B95F4 | `Harddisk` | Known | Hardware interface |
| 0x002B9600 | `FireWire-forbindelser underst` | Known | FireWire interface |
| 0x002B9664 | `re sange og data.` | Known | ATA/disk interface |
| 0x002B9684 | `Afspil %s` | Known | Hardware interface |
| 0x002B9B24 | `Slut iPod til iTunes, og installer spillet igen.` | Known | Hardware interface |
| 0x002B9B58 | `Spillet kan ikke spilles.` | Known | Hardware interface |
| 0x002B9BF4 | `Denne version af spillet underst` | Known | Hardware interface |
| 0x002B9D8F | ` afspilningsknappen p` | Known | Hardware interface |
| 0x002BA42F | `je den til spillelisten On-The-Go. Hold knappen nede, n` | Known | Hardware interface |
| 0x002BA468 | `r en spilleliste, kunstner eller et album er valgt for ...` | Known | Hardware interface |
| 0x002BA4A5 | `je alle sangene til spillelisten On-The-Go.` | Known | Hardware interface |
| 0x002BA880 | `Nyt spil` | Known | Hardware interface |
| 0x002BA894 | `Afspil` | Known | Hardware interface |
| 0x002BACAC | `Afspil video` | Known | Hardware interface |
| 0x002BAD28 | `Dette mediearkiv kan ikke vises eller afspilles p` | Known | Hardware interface |
| 0x002BB320 | `FireWire tilsluttet` | Known | FireWire interface |
| 0x002BF138 | `Spiele` | Known | Hardware interface |
| 0x002BF1B8 | `Weiterspielen` | Known | Hardware interface |
| 0x002BFAF8 | `Kolkata (Kalkutta)` | Known | ATA/disk interface |
| 0x002BFEA8 | `FireWire wird nicht unterst` | Known | FireWire interface |
| 0x002BFF80 | `Spitzname` | Known | Hardware interface |
| 0x002C0324 | `Beispiel` | Known | Hardware interface |
| 0x002C033C | `Beispielfirma GmbH` | Known | Hardware interface |
| 0x002C0350 | `Dieses Beispiel zeigt, welche Infos Sie bei einem Konta...` | Known | Hardware interface |
| 0x002C0590 | `Verbinden Sie Ihren iPod mit iTunes und installieren Si...` | Known | Hardware interface |
| 0x002C05DC | `Dieses Spiel kann nicht gespielt werden.` | Known | Hardware interface |
| 0x002C069C | `Diese Version des Spiels wird nicht mehr unterst` | Known | Hardware interface |
| 0x002C13D0 | `Neues Spiel` | Known | Hardware interface |
| 0x002C18A8 | `Die Mediendatei kann nicht auf dem iPod angezeigt oder ...` | Known | Hardware interface |
| 0x002C1F2E | `ber FireWire verbunden` | Known | FireWire interface |
| 0x002C7AC6 | ` FireWire. ` | Known | FireWire interface |
| 0x002CB066 | ` FireWire` | Known | FireWire interface |
| 0x002CF414 | `Kolkata (Calcuta)` | Known | ATA/disk interface |
| 0x002CF7B9 | `de canciones o archivos no son posibles con FireWire :` | Known | FireWire interface |
| 0x002D09F4 | `Espiral` | Known | Hardware interface |
| 0x002D16E8 | `FireWire conectado` | Known | FireWire interface |
| 0x002D4B90 | `Etsi kanavia -komento etsii kaikki saatavilla olevat ra...` | Known | ATA/disk interface |
| 0x002D4D04 | `RDS-dataa ei havaittu` | Known | ATA/disk interface |
| 0x002D5250 | `Diskanttivahv.` | Known | Hardware interface |
| 0x002D5260 | `Diskanttiheik.` | Known | Hardware interface |
| 0x002D528C | `Diskantinkorostus` | Known | Hardware interface |
| 0x002D5C44 | `Ladataan` | Known | ATA/disk interface |
| 0x002D5D18 | `FireWire-tiedonsiirtoa ei tueta. Siirt` | Known | FireWire interface |
| 0x002D5DB4 | `Ladataan...` | Known | ATA/disk interface |
| 0x002D5E98 | `yhteystietoa ladataan.` | Known | ATA/disk interface |
| 0x002D61AC | `Yhteystietoa ladataan.` | Known | ATA/disk interface |
| 0x002D626E | ` ei voi pelata.` | Known | ATA/disk interface |
| 0x002D65D8 | `Muistiinpanoja ladataan.` | Known | ATA/disk interface |
| 0x002D7788 | `nityksen jatkamiseen ei ole tarpeeksi vapaata levytilaa...` | Known | ATA/disk interface |
| 0x002D77C8 | `nityksen aloittamiseen ei ole tarpeeksi vapaata levytil...` | Known | ATA/disk interface |
| 0x002D7A18 | `FireWire liitetty` | Known | FireWire interface |
| 0x002DCCE5 | `s via FireWire : connectez l` | Known | FireWire interface |
| 0x002DEF08 | `FireWire Connect` | Known | FireWire interface |
| 0x002E3714 | `A FireWire kapcsolat nem t` | Known | FireWire interface |
| 0x002E5924 | `FireWire csatlakozik` | Known | FireWire interface |
| 0x002E8B70 | `Durata diapositiva` | Known | ATA/disk interface |
| 0x002E8ED8 | ` stata effettuata. Premi e mantieni premuto il pulsante...` | Known | ATA/disk interface |
| 0x002E9F8C | `Connessioni di dati via FireWire non sono supportate. P...` | Known | USB interface |
| 0x002EA090 | `auto privata` | Known | ATA/disk interface |
| 0x002EA60D | ` supportata.` | Known | ATA/disk interface |
| 0x002EAFF0 | `Data & ora` | Known | ATA/disk interface |
| 0x002EB0EC | `Spazzata dal centro` | Known | ATA/disk interface |
| 0x002EB100 | `Spazzata Verso il basso` | Known | ATA/disk interface |
| 0x002EB118 | `Spazzata di lato` | Known | ATA/disk interface |
| 0x002EB12C | `Spirale` | Known | Hardware interface |
| 0x002EB13C | `Spinta verso il basso` | Known | Hardware interface |
| 0x002EB154 | `Spinta di lato` | Known | Hardware interface |
| 0x002EB308 | `Imposta data & ora` | Known | ATA/disk interface |
| 0x002EBD18 | `Data & Ora` | Known | ATA/disk interface |
| 0x002EBDDC | `FireWire connesso` | Known | FireWire interface |
| 0x002F0DAC | `FireWire ` | Known | FireWire interface |
| 0x002FD95C | `Handmatig` | Known | Hardware interface |
| 0x002FEB20 | `Jekatarinenburg` | Known | ATA/disk interface |
| 0x002FED7A | `ren via FireWire, maar alleen via de meegeleverde USB-k...` | Known | USB interface |
| 0x00300B88 | `FireWire aangesloten` | Known | FireWire interface |
| 0x00303AAC | `Hvis du vil vise tekstfiler her, aktiverer du iPod for ...` | Known | Hardware interface |
| 0x00304154 | `Finner ikke RDS-data` | Known | ATA/disk interface |
| 0x003042E0 | `Spilles n` | Known | Hardware interface |
| 0x003043EC | `Fortsett spill` | Known | Hardware interface |
| 0x00304458 | `Under avspilling` | Known | Hardware interface |
| 0x0030461C | `Diskantforsterkning` | Known | Hardware interface |
| 0x00304630 | `Diskantreduksjon` | Known | Hardware interface |
| 0x00304F5C | `Slett spilleliste` | Known | Hardware interface |
| 0x003050BC | `Diskmodus` | Known | Hardware interface |
| 0x003050CF | `ring via FireWire st` | Known | FireWire interface |
| 0x00305123 | `re sanger eller data.` | Known | ATA/disk interface |
| 0x00305148 | `Spill %s` | Known | Hardware interface |
| 0x003055F0 | `Koble iPod til iTunes, og installer spillet p` | Known | Hardware interface |
| 0x00305628 | `Dette spillet kan ikke spilles.` | Known | Hardware interface |
| 0x003056C0 | `Denne versjonen av dette spillet st` | Known | Hardware interface |
| 0x003056F8 | `Hvis du glemmer kombinasjonen, kan du koble iPod til da...` | Known | ATA/disk interface |
| 0x003057D0 | `r bildene til datamaskinen, og synkroniser dem via iTun...` | Known | ATA/disk interface |
| 0x0030585E | ` avspillingsknappen p` | Known | Hardware interface |
| 0x00305E92 | ` legge den til i On-The-Go-spillelisten. Spillelister, ...` | Known | Hardware interface |
| 0x00305FB7 | ` denne iPod-enheten. Koble iPod til datamaskinen, og st...` | Known | ATA/disk interface |
| 0x003062E0 | `Nytt spill` | Known | Hardware interface |
| 0x0030677C | `Denne mediefilen kan ikke vises eller spilles p` | Known | Hardware interface |
| 0x003067C3 | ` datamaskinen ved hjelp av QuickTime.` | Known | ATA/disk interface |
| 0x00306824 | `r importerte bilder til datamaskinen, og synkroniser vi...` | Known | ATA/disk interface |
| 0x00306B08 | `Det er ikke nok ledig diskplass til ` | Known | Hardware interface |
| 0x00306D70 | `Koblet til via FireWire` | Known | FireWire interface |
| 0x00309F98 | `Strata` | Known | ATA/disk interface |
| 0x0030B47F | `czenie FireWire nie jest wspierane. By przes` | Known | FireWire interface |
| 0x0030C53C | `Data i czas` | Known | ATA/disk interface |
| 0x0030D3CF | `czony przez Firewire` | Known | FireWire interface |
| 0x00311700 | `Kolkata (Calcut` | Known | ATA/disk interface |
| 0x00311ACF | `es FireWire n` | Known | FireWire interface |
| 0x00312B08 | `Data & hora` | Known | ATA/disk interface |
| 0x00312E84 | `Definir data & hora` | Known | ATA/disk interface |
| 0x003139D0 | `FireWire ligado` | Known | FireWire interface |
| 0x00319395 | ` FireWire ` | Known | FireWire interface |
| 0x0031F970 | `inget kort inmatat` | Known | ATA/disk interface |
| 0x0031FEB0 | `Inga RDS-data kan hittas` | Known | ATA/disk interface |
| 0x00320E24 | `FireWire-` | Known | FireWire interface |
| 0x00320E74 | `ver musik eller data.` | Known | ATA/disk interface |
| 0x00321BD8 | `Stort bildmaterial` | Known | Hardware interface |
| 0x00322B40 | `FireWire anslutet` | Known | FireWire interface |
| 0x00325AAE | `in iPod'u disk kullan` | Known | Hardware interface |
| 0x00327170 | `Disk Modu` | Known | Hardware interface |
| 0x0032717C | `FireWire ba` | Known | FireWire interface |
| 0x00327A48 | `nda bir hata olu` | Known | ATA/disk interface |
| 0x0032891C | `Bilinmeyen Hata` | Known | ATA/disk interface |
| 0x00328E21 | ` disk alan` | Known | Hardware interface |
| 0x00329088 | `FireWire Ba` | Known | FireWire interface |
| 0x004B48E8 | `To view text files here, enable iPod for disk use, then...` | Known | Hardware interface |
| 0x004B4F74 | `No RDS Data Detected` | Known | ATA/disk interface |
| 0x004B63FC | `Disk Mode` | Known | Hardware interface |
| 0x004B6408 | `FireWire connections are not supported. To transfer son...` | Known | USB interface |
| 0x004B723C | `The battery level is too low.` | Known | Power management |
| 0x004B75E0 | `Disk Browser` | Known | Hardware interface |
| 0x004B7E14 | `There is not enough free disk space to continue recordi...` | Known | Hardware interface |
| 0x004B7E50 | `There is not enough free disk space to start recording.` | Known | Hardware interface |
| 0x004B80B0 | `FireWire Connected` | Known | FireWire interface |
| 0x004B80C4 | `No battery power remains. Please connect iPod to power.` | Known | Power management |
| 0x004B80FC | `Low Battery` | Known | Power management |
| 0x004FFE50 | `I2C write Error` | Known | Hardware interface |
| 0x004FFE64 | `I2C read Error %02x` | Known | Hardware interface |
| 0x00642A48 | `TROMResourceDB - unknown header version! (Try regenerat...` | Known | ATA/disk interface |
| 0x00653EF3 | `ataTaza[aea;ajaaaVa)b'b+b+dMd[d]dtdvdrdsd}dudfd` | Known | ATA/disk interface |
| 0x00658ADD | `aGa>a(a'aJa?a<a,a4a=aBaDasawaXaYaZakataoaeaqa_a]aSaua` | Known | ATA/disk interface |
| 0x00669DA9 | `15TCountedPointerI10SImageDataE` | Known | ATA/disk interface |
| 0x00669DDB | `15iMAXMLParseData` | Known | ATA/disk interface |
| 0x00669EE5 | `N4eApp17ManifestDataProxyE` | Known | ATA/disk interface |
| 0x0067C65F | `HoldSwitch` | Known | Hardware interface |
| 0x0067D1D9 | `Bad Data` | Known | ATA/disk interface |
| 0x0067D9CB | `ex_data` | Known | ATA/disk interface |
| 0x0067DAC4 | `RSA Data Security, Inc.` | Known | ATA/disk interface |
| 0x0067DF7E | `set-brand-IATA-ATA` | Known | ATA/disk interface |
| 0x0067E6D3 | `RSA Data Security, Inc. PKCS` | Known | ATA/disk interface |
| 0x0067E8DB | `setCext-Track2Data` | Known | ATA/disk interface |
| 0x0067E8EE | `id-cct-PKIData` | Known | ATA/disk interface |
| 0x0067E8FD | `setct-OIData` | Known | ATA/disk interface |
| 0x0067E90A | `setct-PIData` | Known | ATA/disk interface |
| 0x0067E917 | `setct-PANData` | Known | ATA/disk interface |
| 0x0067E925 | `qualityLabelledData` | Known | ATA/disk interface |
| 0x0067E939 | `pkcs7-signedData` | Known | ATA/disk interface |
| 0x0067E94A | `pkcs7-signedAndEnvelopedData` | Known | ATA/disk interface |
| 0x0067E967 | `pkcs7-envelopedData` | Known | ATA/disk interface |
| 0x0067E97B | `pkcs7-encryptedData` | Known | ATA/disk interface |
| 0x0067E98F | `id-smime-ct-DVCSResponseData` | Known | ATA/disk interface |
| 0x0067E9AC | `setCext-merchData` | Known | ATA/disk interface |
| 0x0067E9BE | `id-smime-ct-authData` | Known | ATA/disk interface |
| 0x0067E9D3 | `id-on-personalData` | Known | ATA/disk interface |
| 0x0067E9E6 | `setct-CapTokenData` | Known | ATA/disk interface |
| 0x0067E9F9 | `setct-BatchAdminReqData` | Known | ATA/disk interface |
| 0x0067EA11 | `setct-CertReqData` | Known | ATA/disk interface |
| 0x0067EA23 | `setct-PCertReqData` | Known | ATA/disk interface |
| 0x0067EA36 | `setct-PResData` | Known | ATA/disk interface |
| 0x0067EA45 | `setct-CredResData` | Known | ATA/disk interface |
| 0x0067EA57 | `setct-BatchAdminResData` | Known | ATA/disk interface |
| 0x0067EA6F | `setct-CapResData` | Known | ATA/disk interface |
| 0x0067EA80 | `setct-PInitResData` | Known | ATA/disk interface |
| 0x0067EA93 | `setct-CertResData` | Known | ATA/disk interface |
| 0x0067EAA5 | `setct-CredRevResData` | Known | ATA/disk interface |
| 0x0067EABA | `setct-AuthRevResData` | Known | ATA/disk interface |
| 0x0067EACF | `setct-CapRevResData` | Known | ATA/disk interface |
| 0x0067EAE3 | `pkcs7-digestData` | Known | ATA/disk interface |
| 0x0067EAF4 | `id-smime-ct-DVCSRequestData` | Known | ATA/disk interface |
| 0x0067EB10 | `pkcs7-data` | Known | ATA/disk interface |
| 0x0067ED6D | `setct-PIDataUnsigned` | Known | ATA/disk interface |
| 0x0067F34F | `Netscape Data Type` | Known | ATA/disk interface |
| 0x0067F375 | `nsDataType` | Known | ATA/disk interface |
| 0x0067FDBE | `id-cmc-dataReturn` | Known | ATA/disk interface |
| 0x00680F1A | `d.data` | Known | ATA/disk interface |
| 0x00680FC0 | `enc_data` | Known | ATA/disk interface |
| 0x0068164D | `Data Encipherment` | Known | ATA/disk interface |
| 0x00681670 | `dataEncipherment` | Known | ATA/disk interface |
| 0x00681950 | `OCSP_RESPDATA` | Known | ATA/disk interface |
| 0x00681979 | `OCSP_RESPID` | Known | Hardware interface |
| 0x00681BC5 | `tbsResponseData` | Known | ATA/disk interface |
| 0x0071C467 | `@!ATAp@-` | Known | ATA/disk interface |
| 0x00AED7E8 | `gldMallocSlow` | Known | Hardware interface |
| 0x00AEE792 | `0BgldMallocSlow` | Known | Hardware interface |
| 0x00AF1B18 | `Length is less than data described in texture sub data ...` | Known | ATA/disk interface |
| 0x00B12264 | `disk_notify` | Known | Hardware interface |
| 0x00B258FC | `TV_DMA_INIT` | Known | Hardware interface |
| 0x00B25908 | `TV_DMA_START` | Known | Hardware interface |
| 0x00B25915 | `TV_DMA_MIDDLE` | Known | Hardware interface |
| 0x00B25923 | `TV_DMA_END` | Known | Hardware interface |
| 0x00B2592E | `TV_DMA_BLOCK` | Known | Hardware interface |
| 0x00B2593B | `TV_DMA_STOP` | Known | Hardware interface |
| 0x00B26851 | `dma_get_transfer_queue` | Known | Hardware interface |
| 0x00B26868 | `dma_memcpy` | Known | Hardware interface |
| 0x00B26873 | `dma_memcpy2d_uncached` | Known | Hardware interface |
| 0x00B26889 | `dma_subchan_free` | Known | Hardware interface |
| 0x00B2689A | `dma_subchan_request` | Known | Hardware interface |
| 0x00B268AE | `dma_transfer_chain` | Known | Hardware interface |
| 0x00B268C1 | `dma_transfer_has_finished` | Known | Hardware interface |
| 0x00B268DB | `dma_transfer_queue_post` | Known | Hardware interface |
| 0x00B268F3 | `dma_transfer_queue_release` | Known | Hardware interface |
| 0x00B2690E | `dma_transfer_set_callback` | Known | Hardware interface |
| 0x00B26928 | `dma_transfer_setup_memcpy` | Known | Hardware interface |
| 0x00B26942 | `dma_transfer_setup_memcpy_uncached` | Known | Hardware interface |
| 0x00B26965 | `dma_transfer_setup_memcpy2d_uncached` | Known | Hardware interface |
| 0x00B2698A | `dma_transfer_wait` | Known | Hardware interface |
| 0x00B26F7F | `vc_image_set_image_data` | Known | ATA/disk interface |
| 0x00B26F97 | `vc_image_set_image_data_yuv` | Known | ATA/disk interface |
| 0x00B3F2C2 | `X(4ATA` | Known | ATA/disk interface |
| 0x00B43947 | `pulse_data` | Known | ATA/disk interface |
| 0x00B43952 | `data_stream_element` | Known | ATA/disk interface |
| 0x00B43985 | `section_data` | Known | ATA/disk interface |
| 0x00B43C36 | `scale_factor_data` | Known | ATA/disk interface |
| 0x00B43C48 | `spectral_data` | Known | ATA/disk interface |
| 0x00B43C7B | `tns_data` | Known | ATA/disk interface |
| 0x00B44D73 | `.rdata` | Known | ATA/disk interface |
| 0x00B44D7A | `.rsdata` | Known | ATA/disk interface |
| 0x00B44D88 | `.sdata` | Known | ATA/disk interface |
| 0x00B44DA8 | `.rela.rdata` | Known | ATA/disk interface |
| 0x00B44DB4 | `.rela.rsdata` | Known | ATA/disk interface |
| 0x00B44DD6 | `.rela.sdata` | Known | ATA/disk interface |
| 0x00B44E32 | `.rela.data` | Known | ATA/disk interface |
| 0x00B5B320 | `h264_setrefdata` | Known | ATA/disk interface |
| 0x00B5B35F | `h264_chromaplane_data` | Known | ATA/disk interface |
| 0x00B5B375 | `h264_lumaplane_data` | Known | ATA/disk interface |
| 0x00B5CB64 | `g_refdata` | Known | ATA/disk interface |
| 0x00B7F921 | `mpeg4_startdata` | Known | ATA/disk interface |
| 0x00B7F931 | `mpeg4_blockdata` | Known | ATA/disk interface |
| 0x00B7F941 | `mpeg4_lastblockdata` | Known | ATA/disk interface |
| 0x00B7FDF4 | `mpeg4dec_fetch_blocks_dma_subchan` | Known | Hardware interface |
| 0x00B7FE16 | `mpeg4dec_fetch_blocks_dma_chan` | Known | Hardware interface |
| 0x00B7FE35 | `mpeg4dec_fetch_blocks_dma_cba` | Known | Hardware interface |
| 0x00B7FE5D | `mpeg4_deststripedatay` | Known | ATA/disk interface |
| 0x00B7FE73 | `mpeg4_deststripedatau` | Known | ATA/disk interface |
| 0x00B7FE89 | `mpeg4_deststripedatav` | Known | ATA/disk interface |
| 0x00B7FEB7 | `mpeg4dec_dma_xfer_q` | Known | Hardware interface |
| 0x00B7FF31 | `mpeg4_stripedatay` | Known | ATA/disk interface |
| 0x00B7FF43 | `mpeg4_stripedatau` | Known | ATA/disk interface |
| 0x00B7FF55 | `mpeg4_stripedatav` | Known | ATA/disk interface |
| 0x00B8001A | `mpeg4_fetcheddata` | Known | ATA/disk interface |
| 0x00B80065 | `fastparse_preparedma` | Known | Hardware interface |
| 0x00B801E8 | `launchdma` | Known | Hardware interface |
| 0x00B801F2 | `launchdma2` | Known | Hardware interface |
| 0x00B80831 | `ubv_initintratables` | Known | ATA/disk interface |
| 0x00B80B73 | `waitfordma` | Known | Hardware interface |
| 0x00B80B83 | `waitfordma2` | Known | Hardware interface |
| 0x00B80C64 | `ubv_vc_intratable` | Known | ATA/disk interface |

---

## 7. Logging/Analytics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00215CF4 | `Channel Reserved` | Hidden | Logging channel |
| 0x00215D08 | `Channel AppBoot` | Hidden | Logging channel |
| 0x00215D18 | `Channel BufferedSongReading` | Hidden | Logging channel |
| 0x00215D34 | `Channel PrefsWriting` | Hidden | Logging channel |
| 0x00215D4C | `Channel GeneralUserExperience` | Hidden | Logging channel |
| 0x00215DA0 | `Channel TestLogging` | Hidden | Logging channel |
| 0x00215DB4 | `Channel AppFileLoading` | Hidden | Logging channel |
| 0x00215DCC | `Channel VCardReading` | Hidden | Logging channel |
| 0x00215DE4 | `Channel LongSongScanning` | Hidden | Logging channel |
| 0x00215E58 | `Channel VoiceRecording` | Hidden | Logging channel |
| 0x00215E70 | `Channel VoiceRecordingNewFileSegment` | Hidden | Logging channel |
| 0x00215E98 | `Channel PhotoBrowse` | Hidden | Logging channel |
| 0x00215EAC | `Channel PhotoImporting` | Hidden | Logging channel |
| 0x00215EC4 | `Channel Notes` | Hidden | Logging channel |
| 0x00215ED4 | `Channel PhotoFileManagement` | Hidden | Logging channel |
| 0x00215F20 | `Channel USBChannel` | Hidden | Logging channel |
| 0x00215F34 | `Channel UnitTests` | Hidden | Logging channel |
| 0x00215F48 | `Channel FreeSpaceCache` | Hidden | Logging channel |
| 0x00215FC0 | `Channel OnTheGoFileMgmt` | Hidden | Logging channel |
| 0x00215FD8 | `Channel SlideShow` | Hidden | Logging channel |
| 0x00215FEC | `Channel ImageCache` | Hidden | Logging channel |
| 0x00216000 | `Channel AlbumArtReading` | Hidden | Logging channel |
| 0x00216018 | `Channel Video` | Hidden | Logging channel |
| 0x0021603C | `Channel ResourceAccess` | Hidden | Logging channel |
| 0x00216054 | `Channel VideoCoreBoot` | Hidden | Logging channel |
| 0x00216088 | `Channel StreamCacheAddFile` | Hidden | Logging channel |
| 0x002160A4 | `Channel FontFileAccess` | Hidden | Logging channel |
| 0x002160BC | `Channel ScreenLock` | Hidden | Logging channel |
| 0x00216140 | `Channel ProfilerAccess` | Hidden | Logging channel |
| 0x00216158 | `Channel eAppAccess` | Hidden | Logging channel |
| 0x0021616C | `Channel eAppWriteBackCache` | Hidden | Logging channel |
| 0x00216188 | `Channel TrainerFileAccess` | Hidden | Logging channel |
| 0x002161A4 | `Channel IapStorage` | Hidden | Logging channel |
| 0x002161B8 | `Channel XMLParsing` | Hidden | Logging channel |
| 0x002161CC | `Channel AudioPrompt` | Hidden | Logging channel |
| 0x002161E0 | `Channel AudioPromptXML` | Hidden | Logging channel |
| 0x002161F8 | `Channel StreamCacheSeek` | Hidden | Logging channel |
| 0x00216B3C | `iPod Usage Stats` | Hidden | Usage telemetry |
| 0x00B15DBC | `pm_stop_logging` | Hidden | Internal logging |
| 0x00B15DCC | `pm_start_logging` | Hidden | Internal logging |
| 0x00B265A9 | `Pm Logging` | Hidden | Internal logging |

---

## 8. Error Messages

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000863DC | `Invalid Operation` | Known | Error/assertion message |
| 0x0009ED40 | `IP Address:<invalid>` | Known | Error/assertion message |
| 0x000EEB4C | `internal error: list index %ld out of range` | Known | Error/assertion message |
| 0x00106C80 | `Root Hub Driver Internal Error unused case in hub handl...` | Known | Error/assertion message |
| 0x00106CBC | `Root hub Error Calling Add Device` | Known | Error/assertion message |
| 0x00141F90 | `Too many errors, further errors discarded. (31)` | Known | Error/assertion message |
| 0x001422E0 | `%s Error in file %s.` | Known | Error/assertion message |
| 0x00142988 | `Cannot link to a .link file. (29)` | Known | Error/assertion message |
| 0x002D0034 | `Error durante la importaci` | Known | Error/assertion message |
| 0x002D0F64 | `Error desconocido` | Known | Error/assertion message |
| 0x002EA85C | `Errore durante l'importazione` | Known | Error/assertion message |
| 0x002EB634 | `Errore sconosciuto` | Known | Error/assertion message |
| 0x004B49D8 | `connection failed` | Known | Error/assertion message |
| 0x004B6918 | `This game cannot be played.` | Known | Error/assertion message |
| 0x004B6A9C | `Imported photos cannot be viewed on TV. Transfer photos...` | Known | Error/assertion message |
| 0x004B6BF0 | `An error occurred while importing` | Known | Error/assertion message |
| 0x004B7278 | `%s failed to launch because its resources cannot be fou...` | Known | Error/assertion message |
| 0x004B7AA4 | `This file cannot be viewed on iPod.` | Known | Error/assertion message |
| 0x004B7AC8 | `This media file cannot be viewed or played on iPod. Use...` | Known | Error/assertion message |
| 0x004B7BC0 | `This photo format cannot be viewed on iPod. Transfer im...` | Known | Error/assertion message |
| 0x004B7DDC | `Cannot record because there is no microphone attached.` | Known | Error/assertion message |
| 0x00681CA6 | `%s: range error: invalid range [%d, %d)` | Known | Error/assertion message |
| 0x00681CF9 | `%s: conversion failed` | Known | Error/assertion message |
| 0x00681D35 | `%s: failed to construct locale name` | Known | Error/assertion message |
| 0x00681D80 | `%s: invalid pointer %p` | Known | Error/assertion message |
| 0x00681DA7 | `%s: unspecified error` | Known | Error/assertion message |
| 0x00681DBD | `%s: runtime error` | Known | Error/assertion message |
| 0x00681DCF | `%s: underflow error` | Known | Error/assertion message |
| 0x00681DE3 | `%s: overflow error` | Known | Error/assertion message |
| 0x00681E93 | `%s: length error: %u > %u` | Known | Error/assertion message |
| 0x00AF14CE | `@>ShaderMachine: Invalid shader type found` | Known | Error/assertion message |
| 0x00B0B8DC | `error=%d error_msg="odd number of arguments"` | Known | Error/assertion message |
| 0x00B0F7D0 | `error=%d error_msg="missing argument"` | Known | Error/assertion message |
| 0x00B0F9CC | `error=%d error_msg="Invalid arguments"` | Known | Error/assertion message |
| 0x00B0F9F4 | `error=%d error_msg="Command not registered"` | Known | Error/assertion message |
| 0x00B11F12 | `0Berror=%d error_msg="bad display"` | Known | Error/assertion message |
| 0x00B12404 | `error=%d error_msg="bad argument"` | Known | Error/assertion message |
| 0x00B12518 | `error=%d error_msg="dlopen: %s"` | Known | Error/assertion message |
| 0x00B12538 | `error=%d error_msg="dl_local_sym: %s"` | Known | Error/assertion message |
| 0x00B12560 | `error=%d error_msg="app already loaded"` | Known | Error/assertion message |
| 0x00B271C8 | `:Cannot print floating point:` | Known | Error/assertion message |
| 0x00B43AA3 | `adts_error_check` | Known | Error/assertion message |
| 0x00B7F853 | `global_bitstream_error` | Known | Error/assertion message |
| 0x00B7F955 | `mpeg4_numinvalidabove` | Known | Error/assertion message |
| 0x00B8404C | `error=%d error_msg="ff/rew unavailable"` | Known | Error/assertion message |
| 0x00B84C28 | `error=%d error_msg="bad parameters"` | Known | Error/assertion message |
| 0x00B852F4 | `error=%d error_msg="bad parameter"` | Known | Error/assertion message |
| 0x00B855CC | `error=%d error_msg="suspended"` | Known | Error/assertion message |
| 0x00B856CC | `error=%d error_msg="not playing or recording"` | Known | Error/assertion message |
| 0x00B857C2 | `0Berror=%d error_msg="not available"` | Known | Error/assertion message |
| 0x00B8584C | `error=%d error_msg="not recording"` | Known | Error/assertion message |
| 0x00B85D5C | `error=%d error_msg="not suspended"` | Known | Error/assertion message |
| 0x00B85EEA | `0Berror=%d error_msg="not playing"` | Known | Error/assertion message |
| 0x00B85F10 | `error=%d error_msg="no video stream"` | Known | Error/assertion message |
| 0x00B85F38 | `error=%d error_msg="screen capture in progress"` | Known | Error/assertion message |
| 0x00B85FEC | `error=%d error_msg="recording"` | Known | Error/assertion message |
| 0x00B8600C | `error=%d error_msg="out of range"` | Known | Error/assertion message |
| 0x00B86030 | `error=%d error_msg="stream not active"` | Known | Error/assertion message |
| 0x00B86134 | `error=%d error_msg="busy"` | Known | Error/assertion message |
| 0x00B86276 | `0Berror=%d error_msg="bad transform"` | Known | Error/assertion message |
| 0x00B86352 | `0Berror=%d error_msg="step unavailable"` | Known | Error/assertion message |
| 0x00B8637C | `error=%d error_msg="step in progress"` | Known | Error/assertion message |
| 0x00B86408 | `error=%d error_msg="idle"` | Known | Error/assertion message |
| 0x00B86484 | `error=%d error_msg="already suspended"` | Known | Error/assertion message |
| 0x00B86504 | `error=%d error_msg="failed"` | Known | Error/assertion message |
| 0x00B96F1A | `0Berror=%d error_msg="bad parameters"` | Known | Error/assertion message |
| 0x00B96FB4 | `error=%d error_msg="bad transform"` | Known | Error/assertion message |
| 0x00B972BE | `0Berror=%d error_msg="busy"` | Known | Error/assertion message |
| 0x00B972FC | `error=%d error_msg="bad direction"` | Known | Error/assertion message |
| 0x00B97320 | `error=%d error_msg="not playing image"` | Known | Error/assertion message |

---

## 9. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0007CC8C | `;9?=3175+)/-#!'%[Y_]SQWUKIOMCAGE{y` | Known | Filesystem path |
| 0x0007CDD9 | `\|yz;8=>7412# %&/,)*` | Known | Filesystem path |
| 0x0007CE8C | `\|ungXQJC4=&/` | Known | Filesystem path |
| 0x0007D0E0 | `85"/di~sP]JG` | Known | Filesystem path |
| 0x0007D14C | `MCQ_u{ig=3!/` | Known | Filesystem path |
| 0x000E2E94 | `iPod_Control/%s%s%s` | Known | Filesystem path |
| 0x000E2EA8 | `iPod_Control/%s/%s%s%s` | Known | Filesystem path |
| 0x00142F8C | `Bad link, no matching </a> for anchor tag. (20)` | Known | Filesystem path |
| 0x0014C240 | `%s<key>%s</key>` | Known | Filesystem path |
| 0x0014C250 | `%s<integer>%d</integer>` | Known | Filesystem path |
| 0x0014C3B0 | `%s<string>%s</string>` | Known | Filesystem path |
| 0x0014C434 | `%s<%s/>` | Known | Filesystem path |
| 0x0014C474 | `%s</dict>` | Known | Filesystem path |
| 0x0014C4B4 | `%s</array>` | Known | Filesystem path |
| 0x0014C5B0 | `%s<real>%s</real>` | Known | Filesystem path |
| 0x00171AD8 | `paMB rtSDIrp/P` | Known | Filesystem path |
| 0x00192FFC | `Created: %d/%d/%4d %d:%02d:%02d %s` | Known | Filesystem path |
| 0x00193020 | `Last Accessed: %d/%d/%4d %2d:%02d:%02d %s` | Known | Filesystem path |
| 0x0019304C | `Modified: %d/%d/%4d %2d:%02d:%02d %s` | Known | Filesystem path |
| 0x001AB9AC | `/iPod_Control/Device/Accessories` | Known | Filesystem path |
| 0x001AC340 | `/iPod_Control/Device/Accessories/Tags` | Known | Filesystem path |
| 0x001AC38C | `%s/Tags/%lu.plist` | Known | Filesystem path |
| 0x001AC3A0 | `%s/Tags/%lu.p7` | Known | Filesystem path |
| 0x00217839 | `Average navigation (Next/Prev) per playback duration: %...` | Known | Filesystem path |
| 0x0024683C | `iPod_Control/Device` | Known | Filesystem path |
| 0x00246850 | `iPod_Control/Device/radio` | Known | Filesystem path |
| 0x00246D44 | `Resources/Fonts` | Known | Filesystem path |
| 0x0027A998 | `iPod S/N` | Known | Filesystem path |
| 0x002B17EC | `%-m/%-d` | Known | Filesystem path |
| 0x002B180C | `%-m/%-d/%y` | Known | Filesystem path |
| 0x002B1AE1 | `ce Features Guide nebo na adrese www.apple.com/support/...` | Known | Filesystem path |
| 0x002B1CF8 | `re: %d (%d/%d)` | Known | Filesystem path |
| 0x002B339F | `ky naleznete na adrese www.apple.com/support/ipod.` | Known | Filesystem path |
| 0x002B3438 | `apple.com/support/ipod` | Known | Filesystem path |
| 0x002B35BB | ` informace naleznete na adrese http://apple.com/support...` | Known | Filesystem path |
| 0x002B3C21 | `USA a/nebo dal` | Known | Filesystem path |
| 0x002B5124 | `www.apple.com/support` | Known | Filesystem path |
| 0x002B80AB | ` www.apple.com/dk/support/ipod.` | Known | Filesystem path |
| 0x002B826C | `Point: %d (%d/%d)` | Known | Filesystem path |
| 0x002B9896 | ` adressen www.apple.com/support/dk/ipod.` | Known | Filesystem path |
| 0x002B9948 | `apple.com/dk/support/ipod` | Known | Filesystem path |
| 0x002B9AD8 | ` http://www.apple.com/dk/support/ipod/` | Known | Filesystem path |
| 0x002BA07B | `rende VoiceAge Corporation i USA og/eller andre lande o...` | Known | Filesystem path |
| 0x002BB39C | `www.apple.com/dk/support` | Known | Filesystem path |
| 0x002BE80A | ` bewegen. Weitere Informationen finden Sie im iPod Hand...` | Known | Filesystem path |
| 0x002BEA38 | `Punkte: %d (%d/%d)` | Known | Filesystem path |
| 0x002BF43C | `Vorn./Nachn.` | Known | Filesystem path |
| 0x002BF44C | `Nachn./Vorn.` | Known | Filesystem path |
| 0x002C02A3 | ` auf Ihrem iPod. Weitere Anleitungen finden Sie im iPod...` | Known | Filesystem path |
| 0x002C03AC | `apple.com/de/support/ipod` | Known | Filesystem path |
| 0x002C0518 | `Weitere Informationen finden Sie unter: http://apple.co...` | Known | Filesystem path |
| 0x002C084C | `ber die Start/Pause-Taste von jedem ausgew` | Known | Filesystem path |
| 0x002C0B13 | ` ist entweder eine eingetragene Marke oder eine Marke d...` | Known | Filesystem path |
| 0x002C1F9C | `www.apple.com/de/support` | Known | Filesystem path |
| 0x002C52AA | ` www.apple.com/support/ipod.` | Known | Filesystem path |
| 0x002C55AC | `: %d (%d/%d)` | Known | Filesystem path |
| 0x002C8265 | ` http://apple.com/support/ipod/` | Known | Filesystem path |
| 0x002CE12D | `jase a www.apple.com/es/support/ipod.` | Known | Filesystem path |
| 0x002CE32C | `Result.: %d (%d/%d)` | Known | Filesystem path |
| 0x002CFA5A | `nea de dicho manual en www.apple.com/es/support/ipod.` | Known | Filesystem path |
| 0x002CFB14 | `apple.com/es/support/ipod` | Known | Filesystem path |
| 0x002CFCA5 | `n, visite http://apple.com/es/support/ipod/` | Known | Filesystem path |
| 0x002D0243 | ` es una marca registrada o una marca comercial de Voice...` | Known | Filesystem path |
| 0x002D0898 | `Fecha/hora` | Known | Filesystem path |
| 0x002D1754 | `www.apple.com/es/support` | Known | Filesystem path |
| 0x002D470E | `tietoja annetaan iPodin ominaisuusoppaassa tai osoittee...` | Known | Filesystem path |
| 0x002D47E4 | `%s / %s` | Known | Filesystem path |
| 0x002D4808 | `%d / %d` | Known | Filesystem path |
| 0x002D4810 | `%d / %d valokuvaa tuotu` | Known | Filesystem path |
| 0x002D490C | `Tulos: %d (%d/%d)` | Known | Filesystem path |
| 0x002D5F77 | `ytyy verkosta osoitteesta www.apple.com/fi/support/ipod...` | Known | Filesystem path |
| 0x002D6040 | `apple.com/fi/support/ipod` | Known | Filesystem path |
| 0x002D61D4 | `ytyy osoitteesta http://www.apple.com/fi/support/ipod/` | Known | Filesystem path |
| 0x002D670B | ` on VoiceAge Corporationin Yhdysvalloissa ja/tai muissa...` | Known | Filesystem path |
| 0x002D7A80 | `www.apple.com/fi/support` | Known | Filesystem path |
| 0x002DB4FA | ` www.apple.com/fr/support/ipod.` | Known | Filesystem path |
| 0x002DCFB6 | ` l'adresse www.apple.com/fr/support/ipod.` | Known | Filesystem path |
| 0x002DD058 | `apple.com/fr/support/ipod` | Known | Filesystem path |
| 0x002DD1F0 | `Pour en savoir plus, veuillez visiter le site http://ap...` | Known | Filesystem path |
| 0x002DD8BA | `tats-Unis et/ou dans d` | Known | Filesystem path |
| 0x002DE273 | `gler date/heure` | Known | Filesystem path |
| 0x002DEF98 | `www.apple.com/fr/support` | Known | Filesystem path |
| 0x002E200B | `togassa meg a www.apple.com/support/ipod weboldalt.` | Known | Filesystem path |
| 0x002E20F0 | `%d / %d f` | Known | Filesystem path |
| 0x002E2214 | `m: %d (%d/%d)` | Known | Filesystem path |
| 0x002E39FF | `i a www.apple.com/support/ipod c` | Known | Filesystem path |
| 0x002E3C68 | `togasson el a http://apple.com/support/ipod/ c` | Known | Filesystem path |
| 0x002E42F3 | `s/vagy m` | Known | Filesystem path |
| 0x002E8944 | ` di iPod" o vai al sito web www.apple.com/it/support/ip...` | Known | Filesystem path |
| 0x002E8B54 | `Punti: %d (%d/%d)` | Known | Filesystem path |
| 0x002EA1F2 | `, consulta la Guida alle caratteristiche di iPod. Sono ...` | Known | Filesystem path |
| 0x002EA308 | `apple.com/it/support/ipod` | Known | Filesystem path |
| 0x002EA47C | `Per ulteriori informazioni, consulta il sito http://app...` | Known | Filesystem path |
| 0x002EA55C | `Per ulteriori informazioni, consulta http://apple.com/i...` | Known | Filesystem path |
| 0x002EBE4C | `www.apple.com/it/support` | Known | Filesystem path |
| 0x002EEB54 | `%b/%-d %-I:%M %p` | Known | Filesystem path |
| 0x002EEB8C | `%y/%-m/%-d` | Known | Filesystem path |
| 0x002EEB98 | `%Y/%b/%-d` | Known | Filesystem path |
| 0x002EEBA4 | `%y/%b/%-d` | Known | Filesystem path |
| 0x002EEE9C | ` www.apple.com/jp/support/ipod ` | Known | Filesystem path |
| 0x002EF0F0 | `%d (%d/%d)` | Known | Filesystem path |
| 0x002F1122 | `www.apple.com/jp/support/ipod ` | Known | Filesystem path |
| 0x002F1214 | `apple.com/jp/support/ipod` | Known | Filesystem path |
| 0x002F13F0 | `http://www.apple.com/jp/support/ipod/ ` | Known | Filesystem path |
| 0x002F3580 | `www.apple.com/jp/support` | Known | Filesystem path |
| 0x002F62E4 | `%Y/%b/%d %A  %I:%M:%S %p` | Known | Filesystem path |
| 0x002F6300 | `%Y/%b/%d` | Known | Filesystem path |
| 0x002F6318 | `%-m/%-d %-I:%M %p` | Known | Filesystem path |
| 0x002F634C | `%Y/%-m/%-d` | Known | Filesystem path |
| 0x002F6639 | ` www.apple.co.kr/support/ipod` | Known | Filesystem path |
| 0x002F6714 | `%d / %d ` | Known | Filesystem path |
| 0x002F8494 | `apple.co.kr/support/ipod` | Known | Filesystem path |
| 0x002F8639 | ` http://www.apple.co.kr/support/ipod/` | Known | Filesystem path |
| 0x002FA240 | `www.apple.co.kr/support` | Known | Filesystem path |
| 0x002FD694 | `Om hier tekstbestanden te bekijken, stelt u de iPod in ...` | Known | Filesystem path |
| 0x002FD924 | `Score: %d (%d/%d)` | Known | Filesystem path |
| 0x002FEF4F | `Raadpleeg de iPod-overzichtshandleiding voor informatie...` | Known | Filesystem path |
| 0x002FF0B4 | `apple.com/nl/support/ipod` | Known | Filesystem path |
| 0x002FF244 | `Meer informatie vindt u op http://apple.com/nl/support/...` | Known | Filesystem path |
| 0x002FF7BB | ` is een gedeponeerd handelsmerk of een handelsmerk van ...` | Known | Filesystem path |
| 0x003000F8 | `Stel datum/tijd in` | Known | Filesystem path |
| 0x00300BFC | `www.apple.com/nl/support` | Known | Filesystem path |
| 0x00303B75 | `r til www.apple.com/no/support/ipod.` | Known | Filesystem path |
| 0x00303D28 | `Poeng: %d (%d/%d)` | Known | Filesystem path |
| 0x00305386 | ` www.apple.com/no/support/ipod.` | Known | Filesystem path |
| 0x00305418 | `apple.com/no/support/ipod` | Known | Filesystem path |
| 0x003055A3 | ` http://www.apple.com/no/support/ipod/` | Known | Filesystem path |
| 0x00305AEF | ` er enten et registrert varemerke eller et varemerke fo...` | Known | Filesystem path |
| 0x00306DE0 | `www.apple.com/no/support` | Known | Filesystem path |
| 0x00309AD4 | `%-d/%-m/%y` | Known | Filesystem path |
| 0x00309DB1 | `ytkownika iPoda lub na stronie www.apple.com/support/ip...` | Known | Filesystem path |
| 0x00309FA0 | `Punkty: %d (%d/%d)` | Known | Filesystem path |
| 0x0030B6E3 | `ugi iPoda. Wersja elektroniczna instrukcji na stronie w...` | Known | Filesystem path |
| 0x0030B95B | ` pod adresem http://apple.com/support/ipod/` | Known | Filesystem path |
| 0x0030BF1A | `onym znakiem towarowym lub znakiem towarowym firmy Voic...` | Known | Filesystem path |
| 0x00310144 | `%-d/%-m` | Known | Filesystem path |
| 0x0031041F | ` para www.apple.com/support/ipod para obter mais inform...` | Known | Filesystem path |
| 0x0031062A | `o: %d (%d/%d)` | Known | Filesystem path |
| 0x003107FE | `d. p/ desbloq.` | Known | Filesystem path |
| 0x00310B08 | `Prima Repr. p/ desligar r` | Known | Filesystem path |
| 0x00310B40 | `o central p/ guardar a esta` | Known | Filesystem path |
| 0x00310BD8 | `Prima Anterior/Seguinte para mudar de Esta` | Known | Filesystem path |
| 0x00311D3D | `es online deste guia podem ser encontradas em www.apple...` | Known | Filesystem path |
| 0x00311F81 | `es, consulte http://apple.com/support/ipod/` | Known | Filesystem path |
| 0x00312028 | `es adicionais, consulte http://apple.com/support/ipod/` | Known | Filesystem path |
| 0x003123E4 | `o p/ impedir mais altera` | Known | Filesystem path |
| 0x00312432 | `o central p/ continuar.` | Known | Filesystem path |
| 0x003124DE | ` uma marca comercial ou marca registada da VoiceAge Cor...` | Known | Filesystem path |
| 0x00316C2B | ` www.apple.com/support/ipod ` | Known | Filesystem path |
| 0x00319892 | `: www.apple.com/ru/support/ipod.` | Known | Filesystem path |
| 0x00319C03 | `: http://apple.com/support/ipod/` | Known | Filesystem path |
| 0x0031F8FE | ` www.apple.com/se/support/ipod.` | Known | Filesystem path |
| 0x0031FAB8 | `ng: %d (%d/%d)` | Known | Filesystem path |
| 0x003210AD | ` adressen www.apple.com/se/support/ipod.` | Known | Filesystem path |
| 0x00321150 | `apple.com/support/se/ipod` | Known | Filesystem path |
| 0x003212E1 | ` http://www.apple.com/se/support/ipod/` | Known | Filesystem path |
| 0x00321878 | `r VoiceAge Corporation i USA och/eller andra l` | Known | Filesystem path |
| 0x00322124 | `ll in datum/tid` | Known | Filesystem path |
| 0x00322BA8 | `www.apple.com/se/support` | Known | Filesystem path |
| 0x00325840 | `%d/%m %-H:%M` | Known | Filesystem path |
| 0x00325B56 | `n ya da www.apple.com/support/ipod adresine gidin.` | Known | Filesystem path |
| 0x00325C3C | `%d / %d foto` | Known | Filesystem path |
| 0x00325D50 | `Puan: %d (%d/%d)` | Known | Filesystem path |
| 0x00327404 | `mleri www.apple.com/support/ipod adresinde bulunabilir.` | Known | Filesystem path |
| 0x00327664 | `in http://apple.com/support/ipod/ adresini ziyaret edin...` | Known | Filesystem path |
| 0x0032773E | `tfen http://apple.com/support/ipod/ adresini ziyaret ed...` | Known | Filesystem path |
| 0x00327971 | `alma/oynatma d` | Known | Filesystem path |
| 0x00327C5F | `n ABD ve/veya di` | Known | Filesystem path |
| 0x0032C0E4 | ` www.apple.com.cn/support/ipod` | Known | Filesystem path |
| 0x0032C1AD | ` %d/%d ` | Known | Filesystem path |
| 0x0032C2B1 | ` %d (%d/%d)` | Known | Filesystem path |
| 0x0032DABF | ` www.apple.com.cn/support/ipod ` | Known | Filesystem path |
| 0x0032DB5C | `apple.com.cn/support/ipod` | Known | Filesystem path |
| 0x0032DCFE | ` http://www.apple.com.cn/support/ipod/` | Known | Filesystem path |
| 0x0032F504 | `www.apple.com.cn/support` | Known | Filesystem path |
| 0x0033255C | ` www.apple.com.tw/support/ipod` | Known | Filesystem path |
| 0x00332621 | ` %d / %d ` | Known | Filesystem path |
| 0x00333F3A | `www.apple.com.tw/support/ipod` | Known | Filesystem path |
| 0x00334189 | `http://www.apple.com.tw/support/ipod/` | Known | Filesystem path |
| 0x00335A38 | `http://www.apple.com.tw/support` | Known | Filesystem path |
| 0x003EADAA | `pcefefefefefefefefefefefefefefefefefefefefef/[6` | Known | Filesystem path |
| 0x003EAE0D | `=/[2\|pc` | Known | Filesystem path |
| 0x003EBF4A | `2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2...` | Known | Filesystem path |
| 0x003EC240 | `2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2...` | Known | Filesystem path |
| 0x003EC536 | `2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|/...` | Known | Filesystem path |
| 0x003EC82C | `2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003ECB22 | `2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003ECE18 | `2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003ED0FE | `pcefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|...` | Known | Filesystem path |
| 0x003ED3F2 | `pcefefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2...` | Known | Filesystem path |
| 0x003ED6E6 | `pcefefefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\...` | Known | Filesystem path |
| 0x003ED9DA | `pcefefefefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|...` | Known | Filesystem path |
| 0x003EDCCE | `pcefefefefefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|2\|2\|/...` | Known | Filesystem path |
| 0x003EDFC2 | `pcefefefefefefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EE2B6 | `pcefefefefefefefefefefefefef2\|2\|2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EE5AA | `pcefefefefefefefefefefefefefef2\|2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EE89E | `pcefefefefefefefefefefefefefefef2\|2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EEB92 | `pcefefefefefefefefefefefefefefefef2\|2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EEE86 | `pcefefefefefefefefefefefefefefefefef2\|2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EF17A | `pcefefefefefefefefefefefefefefefefefef2\|2\|2\|/[` | Known | Filesystem path |
| 0x003EF46E | `pcefefefefefefefefefefefefefefefefefefef2\|2\|/[` | Known | Filesystem path |
| 0x003EF762 | `pcefefefefefefefefefefefefefefefefefefefef2\|/[6` | Known | Filesystem path |
| 0x00404A25 | `d/(1010101010101010101010101010101010101010101010101010...` | Known | Filesystem path |
| 0x004A0DC5 | `cpcocococOc/[ocpcOc` | Known | Filesystem path |
| 0x004A1555 | `cO[/[ococ` | Known | Filesystem path |
| 0x004A1F6F | `cococococococOc/[ocococococ` | Known | Filesystem path |
| 0x004A1FE3 | `kO[N[oc/[O[pc` | Known | Filesystem path |
| 0x004A21CF | `R/[ocO[N[ocpcOcocN[` | Known | Filesystem path |
| 0x004A21E9 | `kococococococOcOc/[.[N[/[O[/[N[Oc` | Known | Filesystem path |
| 0x004A2427 | `kocOc/[O[` | Known | Filesystem path |
| 0x004A2435 | `cOcO[N[ocOc/[oc` | Known | Filesystem path |
| 0x004A2445 | `c/[.[O[/[` | Known | Filesystem path |
| 0x004A244F | `S/[OcN[N[OcOcococ` | Known | Filesystem path |
| 0x004A2467 | `cocococ.[N[oc.[/[O[/[N[N[Ococ/[/[Oc/[OcO[.[N[ocococ/[oc...` | Known | Filesystem path |
| 0x004A2593 | `cO[/[OcpcPc` | Known | Filesystem path |
| 0x004A26B5 | `kocOcO[O[oc/[O[Oc.[` | Known | Filesystem path |
| 0x004A2937 | `kOcoc/[OcOcO[.[.[.[/[` | Known | Filesystem path |
| 0x004A29B3 | `SO[O[/[ocOc.[.[` | Known | Filesystem path |
| 0x004A29C5 | `[.[/[.[ococN[.[.[.[N[Oc` | Known | Filesystem path |
| 0x004A2A01 | `[/[.[N[O[N[` | Known | Filesystem path |
| 0x004A2A0D | `[N[O[.[OcO[.[N[ococococN[O[OcOcoc/[ococpcocOcococN[.[oc...` | Known | Filesystem path |
| 0x004A2A53 | `cococococ/[/[/[/[O[OcocOcpcpcpcpcpcoc/[/[pc` | Known | Filesystem path |
| 0x004A2AF3 | `koc/[/[ocOc` | Known | Filesystem path |
| 0x004A2BA9 | `cocOc/[ocococ` | Known | Filesystem path |
| 0x004A2D87 | `[.[/[/[` | Known | Filesystem path |
| 0x004A2DB1 | `[.[/[.[.[/[/[.[` | Known | Filesystem path |
| 0x004A2DF5 | `[Oc/[N[Oc.[/[` | Known | Filesystem path |
| 0x004A2E35 | `kpcO[/[/[.[N[` | Known | Filesystem path |
| 0x004A3335 | `[/[Oc.[` | Known | Filesystem path |
| 0x004A3827 | `k/[OcO[OcpcocO[.[.[` | Known | Filesystem path |
| 0x004A3AC4 | `/[O[O[Oc` | Known | Filesystem path |
| 0x004A3AF3 | `B/[O[/[O[O[O[.[.[.[` | Known | Filesystem path |
| 0x004A3B09 | `JkBOcO[O[/[O[Oc.[` | Known | Filesystem path |
| 0x004A3BC9 | `J*:O[O[/[` | Known | Filesystem path |
| 0x004A3C11 | `JJB/[.[O[N[` | Known | Filesystem path |
| 0x004A3C73 | `:O[/[O[.[` | Known | Filesystem path |
| 0x004A3CB9 | `J*:O[/[O[O[.[` | Known | Filesystem path |
| 0x004A3D33 | `1/[.[.[` | Known | Filesystem path |
| 0x004A8E85 | `R/[O[Ocpc` | Known | Filesystem path |
| 0x004A8E9D | `J/[OcOcpc` | Known | Filesystem path |
| 0x004A8ECD | `R/[O[OcOcpcpc` | Known | Filesystem path |
| 0x004A8F19 | `S.S/[pc` | Known | Filesystem path |
| 0x004A8F79 | `SOcO[/[pcpc` | Known | Filesystem path |
| 0x004A8F93 | `[/[O[/[.SO[` | Known | Filesystem path |
| 0x004A8FAF | `S/[O[/[pcpc` | Known | Filesystem path |
| 0x004A8FD7 | `R/[O[/[` | Known | Filesystem path |
| 0x004A9029 | `cOc/[/[/[` | Known | Filesystem path |
| 0x004A9039 | `R/[/[O[pcpc` | Known | Filesystem path |
| 0x004A9053 | `S/[/[.S/[` | Known | Filesystem path |
| 0x004A906D | `[/[/[O[OcpcO[` | Known | Filesystem path |
| 0x004A9083 | `S/[/[PcP[/[O[pc` | Known | Filesystem path |
| 0x004A90B3 | `[/[/[P[pcpcOcO[` | Known | Filesystem path |
| 0x004A90D1 | `S/[pcpc` | Known | Filesystem path |
| 0x004A90E7 | `R/[O[pcpcpc` | Known | Filesystem path |
| 0x004A90FF | `Spcpc/[/[` | Known | Filesystem path |
| 0x004A9131 | `R/[/[pc` | Known | Filesystem path |
| 0x004A9147 | `R/[pc/[/[P[` | Known | Filesystem path |
| 0x004A915D | `[/[/[O[` | Known | Filesystem path |
| 0x004A9193 | `S/[O[/[` | Known | Filesystem path |
| 0x004A91A7 | `S/[/[/[P[P[` | Known | Filesystem path |
| 0x004A91BF | `R/[/[P[P[/[` | Known | Filesystem path |
| 0x004A91DB | `S/[pcO[` | Known | Filesystem path |
| 0x004A9205 | `S.SPcP[/[O[` | Known | Filesystem path |
| 0x004A921F | `S/[pcP[/[` | Known | Filesystem path |
| 0x004A92B1 | `[O[/[P[/[` | Known | Filesystem path |
| 0x004B65D5 | `Refer to the iPod Features Guide for instructions on ho...` | Known | Filesystem path |
| 0x004B6878 | `For more information, please visit http://apple.com/sup...` | Known | Filesystem path |
| 0x004B6934 | `For additional information, please visit http://apple.c...` | Known | Filesystem path |
| 0x004B6E07 | ` is either registered trademark or trademark of VoiceAg...` | Known | Filesystem path |
| 0x004F8591 | ` !"#$%&'()*+,-./0123456789:;<=>?@abcdefghijklmnopqrstuv...` | Known | Filesystem path |
| 0x004F8691 | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x004FEDD0 | `$X/wTNw` | Known | Filesystem path |
| 0x004FFEC0 | `{{~~  /-----\   {{~~ /       \  {{~~\|         \| {{~~\...` | Known | Filesystem path |
| 0x00500007 | `<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1....` | Known | Filesystem path |
| 0x0050008E | `</plist>` | Known | Filesystem path |
| 0x005002AA | `_/:>v?J7` | Known | Filesystem path |
| 0x0050498C | `!"#$%&'.,()+-=_/:;<>?@[]abcdefghijklmnopqrstuvwxyzABCDE...` | Known | Filesystem path |
| 0x0051C4B7 | `W/}lE>q` | Known | Filesystem path |
| 0x0054B0F9 | `H."0*Bx/` | Known | Filesystem path |
| 0x00552A56 | `U/~RERT` | Known | Filesystem path |
| 0x005577BE | `TUOPT/\|` | Known | Filesystem path |
| 0x0055EA53 | `HuGZp/$j` | Known | Filesystem path |
| 0x00564F63 | `(bJ)b"(b2""z/` | Known | Filesystem path |
| 0x005676D3 | `JUAPDD(/` | Known | Filesystem path |
| 0x0056D9E2 | `/B\|$BD'` | Known | Filesystem path |
| 0x0056E95F | `$Bd$BT/` | Known | Filesystem path |
| 0x00574C37 | `/" +J\|!` | Known | Filesystem path |
| 0x0057BAB6 | `Fb""")/` | Known | Filesystem path |
| 0x0057CC6D | `/RyO(UIH` | Known | Filesystem path |
| 0x0057DD1D | `~$Bh'"~$Bz/` | Known | Filesystem path |
| 0x00581587 | `$B +BZ/` | Known | Filesystem path |
| 0x00588FA5 | `0c(HBP/` | Known | Filesystem path |
| 0x0058D53B | `$B~("\|/` | Known | Filesystem path |
| 0x005A552D | `T/DDDDD` | Known | Filesystem path |
| 0x005A5797 | `"~UeB /` | Known | Filesystem path |
| 0x005A8249 | `$B((B /` | Known | Filesystem path |
| 0x005B032C | ` "\|$B~/` | Known | Filesystem path |
| 0x005B34A8 | `@$B\|$"(/` | Known | Filesystem path |
| 0x005B45A8 | `)"8/B""` | Known | Filesystem path |
| 0x005B4CF0 | `r4c6 bN/` | Known | Filesystem path |
| 0x005BA0F1 | `RDT%B(/` | Known | Filesystem path |
| 0x005BB27D | `RBHUE\|/` | Known | Filesystem path |
| 0x005C28FD | `]B""B</` | Known | Filesystem path |
| 0x005C617A | `,B\|RED/` | Known | Filesystem path |
| 0x005CBBFD | `$BT). /` | Known | Filesystem path |
| 0x005CCDFD | `#"TUB(/` | Known | Filesystem path |
| 0x005FE78D | `x$DDC/T` | Known | Filesystem path |
| 0x006163AD | `/" %BD"` | Known | Filesystem path |
| 0x0061CCD0 | `ODD""(/` | Known | Filesystem path |
| 0x0061DF0B | `B"$R%"B$" /` | Known | Filesystem path |
| 0x0061E878 | `bG\|jG\|/` | Known | Filesystem path |
| 0x00620462 | `$E$$BR/` | Known | Filesystem path |
| 0x00620503 | `dRB~RA$/` | Known | Filesystem path |
| 0x00620E54 | `TT&T%B(/` | Known | Filesystem path |
| 0x0062FB48 | `)'>$B8/` | Known | Filesystem path |
| 0x00632156 | `$B\|%EV/` | Known | Filesystem path |
| 0x006396AA | `BDU!BJ ""/` | Known | Filesystem path |
| 0x0063A60E | `Z-bD("(%B>/` | Known | Filesystem path |
| 0x00642D5C | `!"#$%&'()*+,-./` | Known | Filesystem path |
| 0x00643516 | `on543k'78%/e/"#`34 '=3?49-?:))60` | Known | Filesystem path |
| 0x006436B1 | `VcYcmo8jics' EfFf~z/` | Known | Filesystem path |
| 0x00643869 | `J=/&5 1Y` | Known | Filesystem path |
| 0x006453C8 | ` ,;=+[]*?<>\|":/\` | Known | Filesystem path |
| 0x006453D9 | `\/:*?"<>\|` | Known | Filesystem path |
| 0x0064B83D | `% %!%"%#%$%%%&%'%(%)%*%+%,%-%.%/%0%1%2%3%4%5%6%7%8%9%:%...` | Known | Filesystem path |
| 0x0064C01D | `qWlIl/Ymg*` | Known | Filesystem path |
| 0x0064D88B | `X)W,W*W3W9W.W/W\W;WBWiW` | Known | Filesystem path |
| 0x0064DF59 | `n/o6oKoto*o` | Known | Filesystem path |
| 0x0064E4F1 | `q/q1qsq\qhqEqrqJqxqzq` | Known | Filesystem path |
| 0x0064F779 | `b6bKbNb/e` | Known | Filesystem path |
| 0x0064F92F | `V3W0W(W-W,W/W)W` | Known | Filesystem path |
| 0x0064FA03 | `NMOOOGOWO^O4O[OUO0OPOQO=O:O8OCOTO<OFOcO\O`O/ONO6OYO]OHO...` | Known | Filesystem path |
| 0x00650061 | `e%f-f f'f/f` | Known | Filesystem path |
| 0x00650363 | `bNc>c/cUcBcFcOcIc:cPc=c*c+c(cMcLcHeIe` | Known | Filesystem path |
| 0x00650615 | `fFUjUfUDU^UaUCUJU1UVUOUUU/UdU8U.U\U,UcU3UAUWU` | Known | Filesystem path |
| 0x0065087B | `\|F}C}q}.}9}<}@}0}3}D}/}B}2}1}=` | Known | Filesystem path |
| 0x00650A03 | `W/X*X4X$X0X1X!X` | Known | Filesystem path |
| 0x00650B19 | `k.l/l,l/n8nTn!n2ngnJn n%n#n` | Known | Filesystem path |
| 0x00650B7B | `r6s%s4s)s:t*t3t"t%t5t6t4t/t` | Known | Filesystem path |
| 0x006513F9 | `S.V;V9V2V?V4V)VSVNVWVtV6V/V0V` | Known | Filesystem path |
| 0x00651A3F | `\|.~>~F~7~2~C~+~=~1~E~A~4~9~H~5~?~/~D` | Known | Filesystem path |
| 0x00652A2F | `P P'P5P/P1P` | Known | Filesystem path |
| 0x00652B81 | `h5h+h-h/hNhDh4h` | Known | Filesystem path |
| 0x00652B95 | `h&h(h.hMh:h%h h,k/k-k1k4kmk` | Known | Filesystem path |
| 0x00652CA1 | `w"w'w#x,x"x5x/x(x.x+x!x)x3x*x1xTy[yOy\ySyRyQy` | Known | Filesystem path |
| 0x0065312B | `v9w/w-w1w2w4w3w=w%w;w5wHxRxIxMxJxLx&xExPxdygyiyjycykyay` | Known | Filesystem path |
| 0x00653187 | `{1{+{-{/{2{8{` | Known | Filesystem path |
| 0x0065361B | `t/uoulu` | Known | Filesystem path |
| 0x00653981 | ``IaJa+aEa6a2a.aFa/aOa)a@a bh` | Known | Filesystem path |
| 0x006539C7 | `d&d0d(dAd5d/d` | Known | Filesystem path |
| 0x00653A2F | `hxi4iii@ioiDiviXiAitiLi;iKi7i\iOiQi2iRi/i{i<iFkEkCkBkHk...` | Known | Filesystem path |
| 0x00654099 | `y+zJz0z/z(z&z` | Known | Filesystem path |
| 0x0065497D | `u/v-v1v=v3v<v5v2v0v` | Known | Filesystem path |
| 0x00654C73 | `X-[%[2[#[,['[&[/[.[{[` | Known | Filesystem path |
| 0x006555A1 | `j<p5p/p7p4p1pBp8p?p:p9p@p;p3pAp` | Known | Filesystem path |
| 0x006557A3 | `w-y1y/yT\|S\|` | Known | Filesystem path |
| 0x00656B6F | `%#%3%+%;%K% %/%(%7%?%` | Known | Filesystem path |
| 0x00656CEE | `02*2+2,2-2.2/2@272B2C292:212>24222;2623252<2=2?282` | Known | Filesystem path |
| 0x00657219 | `\7_J_/`P`m`` | Known | Filesystem path |
| 0x006573F5 | `OHSIT>T/Z` | Known | Filesystem path |
| 0x0065746B | `i_l*mim/n` | Known | Filesystem path |
| 0x0065795D | `N,p]u/f` | Known | Filesystem path |
| 0x00658669 | `S#S/S1S3S8S@SFSES` | Known | Filesystem path |
| 0x00658769 | `q4V6V2V8VkVdV/VlVjV` | Known | Filesystem path |
| 0x00658A11 | `_)_-_8_A_H_L_N_/_Q_V_W_Y_a_m_s_w_` | Known | Filesystem path |
| 0x006590D7 | `s4s/s)s%s>sNsOs` | Known | Filesystem path |
| 0x0065CEDB | `h>kLp/t` | Known | Filesystem path |
| 0x0065D565 | `o;v/}7~` | Known | Filesystem path |
| 0x0065E329 | `e1f/h\q6z` | Known | Filesystem path |
| 0x0065E975 | `UuX/c"dIfKfmh` | Known | Filesystem path |
| 0x00663422 | `  !"##$%&&'())*+,-../01234556789:;<=>?@ABCDEFGHIJKMNOPQ...` | Known | Filesystem path |
| 0x00663648 | ` !""#$%&''()*+,-./0123456789:;<>?@ABDEFGIJKMNOQRTUVXY[\...` | Known | Filesystem path |
| 0x00663C90 | `/B'2N6REMQLEVJ\|aViu\J]lLm` | Known | Filesystem path |
| 0x0067C6BB | `iPod_Control/games_RO/` | Known | Filesystem path |
| 0x0067C785 | `iPod_Control/Device/accessories` | Known | Filesystem path |
| 0x0067CAB8 | `iPod_Control/iTunes/` | Known | Filesystem path |
| 0x0067CAD4 | `Recordings/` | Known | Filesystem path |
| 0x0067CAE0 | `Calendars/` | Known | Filesystem path |
| 0x0067CAEB | `Contacts/` | Known | Filesystem path |
| 0x0067CBF6 | `/Resources/VideoCore` | Known | Filesystem path |
| 0x0067D137 | `file://` | Known | Filesystem path |
| 0x0067D145 | `</ROT13>` | Known | Filesystem path |
| 0x0067D15E | `</TITLE>` | Known | Filesystem path |
| 0x0067D1AD | `</BODY>` | Known | Filesystem path |
| 0x006803DD | `S/MIME Capabilities` | Known | Filesystem path |
| 0x0068119A | `S/MIME email` | Known | Filesystem path |
| 0x00681203 | `S/MIME signing` | Known | Filesystem path |
| 0x00681230 | `S/MIME encryption` | Known | Filesystem path |
| 0x00681564 | `S/MIME CA` | Known | Filesystem path |
| 0x006A7CD9 | `WV>P7C/;` | Known | Filesystem path |
| 0x006BC510 | `/,)&/,)&/,)&/,)&` | Known | Filesystem path |
| 0x006BD54E | `/ 0 1 2 3 4 5 6 7 8 9 : ; < = > ? @ A B C D E F G H I J...` | Known | Filesystem path |
| 0x006BD72F | `! !!!"!#!$!%!&!'!(!)!*!+!,!-!.!/!0!1!2!3!4!5!6!7!8!9!:!...` | Known | Filesystem path |
| 0x006C072F | `*Zj"/8'6V` | Known | Filesystem path |
| 0x0072B310 | `/1f;{1Q` | Known | Filesystem path |
| 0x0075AF62 | `8/868@8\8a8` | Known | Filesystem path |
| 0x0075B2BA | `S!S$S*S/S3S5S;SOSXSZS\SiSuS{S` | Known | Filesystem path |
| 0x0075B310 | `T!T$T/T1T6T9T@TCTFTJTNTQTUTXT\T_TbTdTfThTlTxT}T` | Known | Filesystem path |
| 0x0075B54A | `Y Y%Y'Y/Y2Y4Y:Y>YBYDYIYKYQY[Y]YcYeYnYvYyY}Y` | Known | Filesystem path |
| 0x0075B5BE | `Z Z#Z%Z'Z)Z+Z-Z/Z6Z<ZAZGZJZUZZZ`ZgZjZmZwZzZ` | Known | Filesystem path |
| 0x0075B7EA | `_#_)_-_/_1_>_A_H_J_N_Y_]_g_m_s_w_y_` | Known | Filesystem path |
| 0x0075B84A | ``"`+`/`1`3`5`<`C`G`M`P`R`U`Z`]`e`m`p`s`x`}`` | Known | Filesystem path |
| 0x0075B97C | `c!c%c(c+c/c2c6c?cFcIcPcScUcWcYc\c_cccecicncrcwc}c` | Known | Filesystem path |
| 0x0075B9FE | `d*d-d/d6d:dDdHdJdNdRdTdXd`dgdidmdpd{d}d` | Known | Filesystem path |
| 0x0075BB62 | `g(g/g1g8g:g=gCgIgQgWgYgfghgjgwg\|g` | Known | Filesystem path |
| 0x0075BD1C | `k$k'k,k/k2k5k;k?kCkGkJkLkNkPkTkVkYk\kgkjklkokuk` | Known | Filesystem path |
| 0x0075BE00 | `m+m/m6m?mAmHmKmOmTm\m^mamfmjmlmpmtmym\|m` | Known | Filesystem path |
| 0x0075BE68 | `n'n)n/n2n4n6n:n>nEnOnQnTnXn\n_nancngninknonsnvn{n` | Known | Filesystem path |
| 0x0075C0E2 | `s"s%s,s/s1s4s7s;s?sAsEsPsRsXs`sdshsrsusxs\|s~s` | Known | Filesystem path |
| 0x0075C288 | `w w"w&w)w-w/w:w<w>wAwGwJwQwXw\whwlwrwzw` | Known | Filesystem path |
| 0x0075C46A | `{ {({/{1{3{6{={A{V{[{]{`{g{i{u{w{z{` | Known | Filesystem path |
| 0x0075C5BC | `~#~(~/~7~;~?~A~H~K~M~R~V~Z~^~b~g~k~p~s~u~y~` | Known | Filesystem path |
| 0x0075E7E4 | `X X#X&X*X-X/X4X9X=X@XIXOXQXTXWX^XaXdXgXiXkXmXoXuXyX\|X` | Known | Filesystem path |
| 0x0075E8D8 | `Z#Z%Z'Z)Z+Z-Z/Z1Z<Z@ZFZIZUZZZ`ZbZjZlZtZzZ~Z` | Known | Filesystem path |
| 0x0075EB04 | `_%_-_/_1_4_@_E_J_L_P_[_a_i_p_w_y_{_` | Known | Filesystem path |
| 0x0075EB62 | `` `$`/`1`3`5`:`A`F`I`P`R`T`Y`]`_`g`o`s`u`z`` | Known | Filesystem path |
| 0x0075EC96 | `c#c'c*c/c2c5c9cAcIcKcScUcWcYc\c^cacecgckcqctczc` | Known | Filesystem path |
| 0x0075ED16 | `d d,d/d4d:d=dFdJdNdQdTdXdZdgdidmdodsd}d` | Known | Filesystem path |
| 0x0075EEE8 | `h!h(h/h7h;h@hHhLhPhWhYh[h_hbhehkhmhrhthyh\|h~h` | Known | Filesystem path |
| 0x0075F036 | `k'k,k/k2k5k7k=kCkFkIkLkNkPkSkVkXk[k_kiklkokrkwk` | Known | Filesystem path |
| 0x0075F210 | `o o"o%o)o/o5o8o<o>oAoEoGoKoMoQoToWo^ofohomotoxozo\|o` | Known | Filesystem path |
| 0x0075F298 | `p#p&p,p/p2p7p9p<p>pCpGpIpNpTpXp]p`pcpipkpupxp\|p` | Known | Filesystem path |
| 0x0075F4DE | `u(u/u7u:uDuYu`ubuduiuou}u` | Known | Filesystem path |
| 0x0075F5A2 | `w"w$w(w,w/w4w<w>w@wEwJwMwXwZw^wjwrwyw\|w` | Known | Filesystem path |
| 0x0075F7EC | `\|#\|&\|*\|/\|3\|6\|=\|E\|J\|L\|O\|V\|[\|c\|g\|i\|l\|r\...` | Known | Filesystem path |
| 0x007643A8 | `t\|tutjthtet^t]t\t:t8t6t5t0t/t.t,t+t)t(t't%t!t t` | Known | Filesystem path |
| 0x0076442C | `rxrirgrfrDrBr?r8r5r1r/r.r-r&r%r r` | Known | Filesystem path |
| 0x007FD0C5 | `UUUUUUa/` | Known | Filesystem path |
| 0x007FD294 | `DDDDDDQ/` | Known | Filesystem path |
| 0x007FD311 | `""""""0/` | Known | Filesystem path |
| 0x0080E0CD | `""f"" /` | Known | Filesystem path |
| 0x00812EAF | `33333330/` | Known | Filesystem path |
| 0x0081396C | `UUUUUUQ/` | Known | Filesystem path |
| 0x008180CA | `fffffffa/` | Known | Filesystem path |
| 0x0082AFEA | `fffffffp/` | Known | Filesystem path |
| 0x008372A3 | `""""""1/` | Known | Filesystem path |
| 0x0084B71B | `/`~UDDE0` | Known | Filesystem path |
| 0x0084E985 | `/Wwx'wv` | Known | Filesystem path |
| 0x0084FF7D | `/PWwww`` | Known | Filesystem path |
| 0x00854D54 | `b""""0/` | Known | Filesystem path |
| 0x00857E60 | `o`o@oP/` | Known | Filesystem path |
| 0x00857E7B | `ofo@oP/` | Known | Filesystem path |
| 0x00870336 | `UUUUUUUQ/` | Known | Filesystem path |
| 0x0087D614 | `"""2""" /` | Known | Filesystem path |
| 0x00883038 | `fb6ffff@/` | Known | Filesystem path |
| 0x0088530F | `/uwwQAy` | Known | Filesystem path |
| 0x00890ED6 | `Gwwww /` | Known | Filesystem path |
| 0x008B31A7 | `33333331/` | Known | Filesystem path |
| 0x008B464D | `vfggffgq/` | Known | Filesystem path |
| 0x008BFB11 | `fffffffq/` | Known | Filesystem path |
| 0x008CD279 | `VffffP/` | Known | Filesystem path |
| 0x008CF701 | `S3W33 /` | Known | Filesystem path |
| 0x008E9675 | `#T"E21/` | Known | Filesystem path |
| 0x008E9A8A | `/`eUUS/` | Known | Filesystem path |
| 0x008EF1AD | `/SDDDDQ` | Known | Filesystem path |
| 0x008FB1F7 | `3333!/{` | Known | Filesystem path |
| 0x008FB29D | `DDDDDB/^` | Known | Filesystem path |
| 0x008FB903 | `/p_OpOJ` | Known | Filesystem path |
| 0x008FD5CE | `/`o_@33L` | Known | Filesystem path |
| 0x0090870C | `""#wwww /` | Known | Filesystem path |
| 0x009090A8 | `DD&wwww`/` | Known | Filesystem path |
| 0x009091ED | `DD8wwwwt/` | Known | Filesystem path |
| 0x0090B70B | `336ffffb/` | Known | Filesystem path |
| 0x00912C7B | `"""""""1/` | Known | Filesystem path |
| 0x009156C3 | `DDDDD@/` | Known | Filesystem path |
| 0x009179FD | `33333A/` | Known | Filesystem path |
| 0x009188FE | `fc4DDB/` | Known | Filesystem path |
| 0x00919758 | `fffff`/` | Known | Filesystem path |
| 0x0091AAE1 | `wwwwwrL/` | Known | Filesystem path |
| 0x0092915B | `wwwwwwwq/` | Known | Filesystem path |
| 0x0093C32C | `C4343 /` | Known | Filesystem path |
| 0x009634E9 | `/rDDDD ` | Known | Filesystem path |
| 0x009763B9 | `UUU5UUUa/` | Known | Filesystem path |
| 0x0097F4E1 | `OqWwwwwP/` | Known | Filesystem path |
| 0x009915D1 | `6ffffc/` | Known | Filesystem path |
| 0x00991F6B | `DDDDDDDA/` | Known | Filesystem path |
| 0x0099286E | `UUU$fff0/` | Known | Filesystem path |
| 0x00992B9F | `3334www@/` | Known | Filesystem path |
| 0x00997227 | `UUPDDD/` | Known | Filesystem path |
| 0x0099AF26 | `_o@/?_i` | Known | Filesystem path |
| 0x0099E8D6 | `7wwwwws0/` | Known | Filesystem path |
| 0x009BCCA1 | `TDDDD /` | Known | Filesystem path |
| 0x009DD2BA | `fffffa/` | Known | Filesystem path |
| 0x009EEC68 | `www7wwwq/` | Known | Filesystem path |
| 0x00A8C63E | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x00AADE1D | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x00AB3E0E | ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUV...` | Known | Filesystem path |
| 0x00AC9473 | `po`?pOp/` | Known | Filesystem path |
| 0x00ACC41A | `/PEBLC2` | Known | Filesystem path |
| 0x00ACE206 | ` " & / 0 : D ` p y ~ ` | Known | Filesystem path |
| 0x00AD6B52 | `/"/0/>/L/Z/h/v/` | Known | Filesystem path |
| 0x00AEC693 | `;0K0B0B0B0B0B0B0Bb/` | Known | Filesystem path |
| 0x00AEC903 | `n0P0De/` | Known | Filesystem path |
| 0x00AEC98B | `gWi7PQcd/` | Known | Filesystem path |
| 0x00AEC9F7 | `gVi6PScd/` | Known | Filesystem path |
| 0x00AECC15 | `gUi5PSch/` | Known | Filesystem path |
| 0x00AECC81 | `gTi4PQch/` | Known | Filesystem path |
| 0x00AECF8F | `gYi9PScd/` | Known | Filesystem path |
| 0x00AECFFB | `gXi8PQcd/` | Known | Filesystem path |
| 0x00AED387 | `10D3P`/` | Known | Filesystem path |
| 0x00AF1730 | `/1An->`/4` | Known | Filesystem path |
| 0x00AF4FFA | `d/wwwwt` | Known | Filesystem path |
| 0x00AF92F0 | `0B0B0B0B0B0B0B0Bg>`/` | Known | Filesystem path |
| 0x00AFA0FA | `0B0B0B0B0B0B0B0B0B0B0B0B0B0B0B0B0B0B0B0B`/` | Known | Filesystem path |
| 0x00AFE17C | `!#%')+-/13579;=?` | Known | Filesystem path |
| 0x00AFE1AC | ` !"#$%&'()*+,-./0123456789:;<=>?` | Known | Filesystem path |
| 0x00B06B12 | `b/ANYDa` | Known | Filesystem path |
| 0x00B06D44 | ``/TNVEa` | Known | Filesystem path |
| 0x00B06F34 | `0D0R0X0R0X0R0X0R0X0R0X0R0Xb/0` | Known | Filesystem path |
| 0x00B07180 | `e/UEUQa` | Known | Filesystem path |
| 0x00B076E8 | `d/AMESa` | Known | Filesystem path |
| 0x00B078FE | ``/RSIHa` | Known | Filesystem path |
| 0x00B083F2 | `a/ffff1y` | Known | Filesystem path |
| 0x00B0843A | `a/wwww1y` | Known | Filesystem path |
| 0x00B08B16 | `d/EMITa` | Known | Filesystem path |
| 0x00B1310D | `o$>P@QA&/ ` | Known | Filesystem path |
| 0x00B13360 | `P@QA / ` | Known | Filesystem path |
| 0x00B1362F | `C"fk>!/^` | Known | Filesystem path |
| 0x00B1A5A8 | `0D0D0D`/` | Known | Filesystem path |
| 0x00B1A8E2 | `0J0D"/d` | Known | Filesystem path |
| 0x00B227E8 | `>?^@_A./` | Known | Filesystem path |
| 0x00B23433 | `;0K0D /`9` | Known | Filesystem path |
| 0x00B25EFB | `#'+/26:?CGMQVZ_dhmrv{` | Known | Filesystem path |
| 0x00B25FD9 | `"&*/37<@EINRW\aejotx}` | Known | Filesystem path |
| 0x00B2925C | `!#%')+-/13579;=?X` | Known | Filesystem path |
| 0x00B2F9A0 | `/mfs/vlls/` | Known | Filesystem path |
| 0x00B38963 | `q/qKqfq` | Known | Filesystem path |
| 0x00B389C5 | `v/vHvavzv` | Known | Filesystem path |
| 0x00B3A2CB | `R/RVR\|R` | Known | Filesystem path |
| 0x00B3A4AF | `p/pGp_pwp` | Known | Filesystem path |
| 0x00B3B3C7 | `pwp_pGp/p` | Known | Filesystem path |
| 0x00B3B5AD | `R\|RVR/R` | Known | Filesystem path |
| 0x00B3F0D3 | `n0Dd/\|%` | Known | Filesystem path |
| 0x00B3F0ED | `L n0Da/` | Known | Filesystem path |
| 0x00B3F0FB | `D0n0Db/` | Known | Filesystem path |
| 0x00B42590 | `))/113//+++(` | Known | Filesystem path |
| 0x00B49655 | `."1PN#/` | Known | Filesystem path |
| 0x00B73171 | `f??R@!/` | Known | Filesystem path |
| 0x00B735AE | `'.5<=6/7>?` | Known | Filesystem path |
| 0x00B735EA | `"#()01*+$%&',-./2389:;4567<=>?` | Known | Filesystem path |
| 0x00B73640 | `&.6>'/7?` | Known | Filesystem path |
| 0x00B848E0 | `/mfs/temp.mid` | Known | Filesystem path |
| 0x00B88D4C | `d/IDIMt(f` | Known | Filesystem path |
| 0x00B89DEC | `e/9AMW$.` | Known | Filesystem path |
| 0x00B94A59 | `Fh`10"/` | Known | Filesystem path |
| 0x00B94D8D | `gh`10 /` | Known | Filesystem path |
| 0x00C39C20 | `]/i>!sa-` | Known | Filesystem path |
| 0x00C45A23 | `VN8e9k/` | Known | Filesystem path |
| 0x00C482DB | `SsTKw/-` | Known | Filesystem path |
| 0x00C494A2 | `Pvt4/_&` | Known | Filesystem path |
| 0x00C5099D | `y2B</;@Szl` | Known | Filesystem path |
| 0x00C5C8F5 | `Z_/Cta]Q` | Known | Filesystem path |
| 0x00C61CF7 | `/1Q5w@;` | Known | Filesystem path |
| 0x00C72694 | `+0$/{#%` | Known | Filesystem path |
| 0x00C904E9 | `0&X/E!F` | Known | Filesystem path |
| 0x00C9250F | `%=/?hW}.` | Known | Filesystem path |
| 0x00C959CD | `\'MiG/u{Z` | Known | Filesystem path |
| 0x00C9DF13 | `/df*l=M3F` | Known | Filesystem path |
| 0x00CB178E | `J/(5kW$` | Known | Filesystem path |
| 0x00CBA028 | `f\FHXQ/` | Known | Filesystem path |
| 0x00CC2F51 | `/%\|m1l~` | Known | Filesystem path |
| 0x00CC7872 | `(q/Gc]m` | Known | Filesystem path |
| 0x00CD38B9 | `/${'",L>` | Known | Filesystem path |
| 0x00CDC56C | `fMiBJZ/` | Known | Filesystem path |
| 0x00CDD595 | `Jc/($]%` | Known | Filesystem path |
| 0x00CEC62A | `9$K9l?/` | Known | Filesystem path |
| 0x00CED6FE | `Td\|/>B,Xok` | Known | Filesystem path |
| 0x00CEDAEB | `e/L&mt)` | Known | Filesystem path |
| 0x00CEEAC7 | `~z"rmi/` | Known | Filesystem path |
| 0x00CEEEEC | `f/nhb`t` | Known | Filesystem path |
| 0x00CF3995 | `h/;^$m-` | Known | Filesystem path |
| 0x00CF4532 | `!v/lRtK` | Known | Filesystem path |
| 0x00CF625B | `/WrX36K*` | Known | Filesystem path |
| 0x00CF8AEB | `3i2/!Md` | Known | Filesystem path |
| 0x00CFB3BC | `64S/_]z_K` | Known | Filesystem path |
| 0x00CFE497 | `>*C8JG/` | Known | Filesystem path |
| 0x00CFFCB0 | `+R/p0=Ez` | Known | Filesystem path |
| 0x00D061F1 | `48/1iE(J` | Known | Filesystem path |
| 0x00D09E46 | `ad/kTQk?` | Known | Filesystem path |
| 0x00D15710 | `6Z<2IbH/` | Known | Filesystem path |
| 0x00D17E0A | `{5Q/{0P}` | Known | Filesystem path |
| 0x00D21370 | `mUBbE/T4` | Known | Filesystem path |
| 0x00D3653E | `m{\Q:%/n` | Known | Filesystem path |
| 0x00D3B6B9 | `9X/zd$9u3` | Known | Filesystem path |
| 0x00D3CE0F | `h/`i2)u` | Known | Filesystem path |
| 0x00D3DB39 | `9/,oR+9f\|H` | Known | Filesystem path |

---

## 10. Nike+/Fitness

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00216188 | `Channel TrainerFileAccess` | Known | Nike+ integration |

---

## 11. Video Playback

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0017AF28 | `VideoCodecs` | Known | Video playback |
| 0x0017B00C | `H.264LC` | Known | Video playback |
| 0x004B7434 | `TV Out` | Known | Video playback |
| 0x00B522E0 | `H.264 Video Decoder` | Known | Video playback |
| 0x00B774A0 | `MPEG-4 video decoder` | Known | Video playback |

---

## 12. Binary Structure (for Ghidra/IDA)

| Parameter | Value |
|-----------|-------|
| **Architecture** | ARM (32-bit, little-endian) |
| **Processor** | ARM7TDMI (ARMv4T), dual-core + video DSP |
| **Base Address** | 0x00000000 |
| **Entry Point** | 0x00000000 |
| **Endianness** | Little-Endian |
| **File Size** | 13,893,632 bytes |

