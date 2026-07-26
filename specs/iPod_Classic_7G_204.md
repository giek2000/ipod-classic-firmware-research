# iPod Classic 7th Generation - RetailOS 2.0.4 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 2.0.4 |
| **Device** | iPod Classic 7th Generation (Late 2009, 160GB) |
| **Binary Size** | 10,599,920 bytes (10.11 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,597,872 bytes |
| **Total Strings** | 55,243 |
| **Function Prologues** | 17,721 |
| **DRAM References** | 62,655 |
| **Peripheral Refs** | 9,477 |

---

## 1. Known User-Facing Features

### 1.1 Main Menu Items (English UI Strings)

| Offset | String | Classification |
|--------|--------|----------------|
| 0x007FF654 | `Videos` | Known - Top-level menu |
| 0x007FF65C | `Extras` | Known - Top-level menu |
| 0x007FF664 | `Speakers` | Known - Top-level menu |
| 0x007FF678 | `Settings` | Known - Top-level menu |
| 0x007FF684 | `Shuffle Songs` | Known - Top-level menu |
| 0x007FF694 | `Now Playing` | Known - Top-level menu |
| 0x007FF6BC | `Cover Flow` | Known - Top-level menu |
| 0x007FF6C8 | `Playlists` | Known - Music submenu |
| 0x007FF6D4 | `Artists` | Known - Music submenu |
| 0x007FF6DC | `Albums` | Known - Music submenu |
| 0x007FF6E4 | `Compilations` | Known - Music submenu |
| 0x007FF6FC | `Genres` | Known - Music submenu |
| 0x007FF704 | `Composers` | Known - Music submenu |
| 0x007FF710 | `Audiobooks` | Known - Music submenu |
| 0x007FF71C | `Search` | Known - Music submenu |
| 0x007FF724 | `Backlight` | Known - Settings |
| 0x007FF738 | `Nike+iPod` | Known - Extras submenu |
| 0x007FF744 | `Rentals` | Known - Videos submenu |
| 0x007FF74C | `Genius` | Known - Music submenu |
| 0x007FF754 | `Genius Mixes` | Known - Music submenu |
| 0x007FF764 | `Alarms` | Known - Extras submenu |
| 0x007FF76C | `Contacts` | Known - Extras submenu |
| 0x007FF778 | `Calendars` | Known - Extras submenu |
| 0x007FF794 | `Screen Lock` | Known - Extras submenu |
| 0x007FF7A0 | `Stopwatch` | Known - Extras submenu |
| 0x007FF7AC | `Voice Memos` | Known - Extras submenu |
| 0x007FF7B8 | `Clocks` | Known - Extras submenu |
| 0x001AAF98 | `Movies` | Known - Videos submenu |
| 0x001AAFA0 | `Music Videos` | Known - Videos submenu |
| 0x001AAFB0 | `TV Shows` | Known - Videos submenu |
| 0x001AAFBC | `Video Podcasts` | Known - Videos submenu |
| 0x001AAFCC | `Rentals` | Known - Videos submenu |
| 0x007FF634 | `Podcasts` | Known - Music submenu |
| 0x007FF640 | `iTunes` | Known - Reference |
| 0x007FF64C | `Photos` | Known - Top-level menu |
| 0x00802CEC | `Disk Browser` | Known - Extras submenu |

### 1.2 Settings Menu

| Offset | String | Classification |
|--------|--------|----------------|
| 0x00802634 | `Volume Limit` | Known - Settings |
| 0x00802664 | `Audiobooks` | Known - Settings |
| 0x00802684 | `Sound Check` | Known - Settings |
| 0x00802714 | `Play Music Library in Sequence` | Known - Settings |
| 0x00801A94 | `Alternate Audio` | Known - Video Settings |
| 0x00800288 | `Sleep Timer` | Known - Clock Settings |
| 0x0080284C | `Set or Lock Volume Limit` | Known - Settings prompt |
| 0x00802868 | `Press Play to Lock Volume Limit` | Known - Settings prompt |

### 1.3 Settings Controllers (TC/ShowSetting/ToggleSetting)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001725F8 | `TCSettings_MainMenu` | Known | Main settings controller |
| 0x00172614 | `TCSettings_MusicMenu` | Known | Music menu settings |
| 0x00172634 | `TCSettings_VolumeLimit` | Known | Volume limit settings |
| 0x00172654 | `TCSettings_Brightness` | Known | Display brightness |
| 0x00172674 | `TCSettings_BacklightTimer` | Known | Backlight auto-off timer |
| 0x00172698 | `TCSettings_EQ` | Known | Equalizer setting |
| 0x001726B0 | `TCSettings_AudiobookSettings` | Known | Audiobook speed control |
| 0x001726D8 | `TCSettings_RadioRegions` | Known | FM radio region selector |
| 0x001726F8 | `TCSettings_ResetAllSettings` | Known | Factory reset |
| 0x00172778 | `TCSettings_AdjustScrollingCntlr` | Known | Scroll speed adjustment |

### 1.4 ToggleSetting Functions

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x001E9D0C | `ToggleSetting_Repeat` | Cycle repeat mode (off/one/all) |
| 0x001E9D28 | `ToggleSetting_Shuffle` | Toggle shuffle mode |
| 0x001E9D40 | `ToggleSetting_TVOut` | Enable/disable TV output |
| 0x001E9D54 | `ToggleSetting_TVSignal` | Toggle PAL/NTSC signal |
| 0x00227F24 | `ToggleSetting_SortBy` | Change sort order |
| 0x00227F3C | `ToggleSetting_ClassicUI` | Toggle Classic UI mode |
| 0x00227F54 | `ToggleSetting_SoundCheck` | Toggle volume normalization |
| 0x00227F70 | `ToggleSetting_Clicker` | Toggle click sound |
| 0x00227F88 | `ToggleSetting_DaylightSavings` | Toggle DST |
| 0x00227FA8 | `ToggleSetting_24HourClock` | Toggle 24h time format |
| 0x00227FC4 | `ToggleSetting_TimeInTitle` | Show time in title bar |

### 1.5 ShowSetting Functions

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x0013F7E4 | `ShowSetting_EQ` | Display EQ selection screen |
| 0x00212D30 | `ShowSetting_Backlight` | Display backlight timer options |
| 0x00227FE0 | `ShowSetting_Shuffle` | Display shuffle state |
| 0x0022818C | `ShowSetting_Repeat` | Display repeat state |
| 0x002281A0 | `ShowSetting_About` | Show About screen |
| 0x002281B4 | `ShowSetting_MainMenu` | Show main menu customization |
| 0x002281CC | `ShowSetting_MusicMenu` | Show music menu customization |
| 0x002281E4 | `ShowSetting_VolumeLimit` | Show volume limit control |
| 0x002281FC | `ShowSetting_BacklightTimer` | Show backlight timer |
| 0x00228218 | `ShowSetting_Brightness` | Show brightness slider |
| 0x00228230 | `ShowSetting_Audiobooks` | Show audiobook speed |
| 0x00228248 | `ShowSetting_RadioRegions` | Show radio region picker |
| 0x00228264 | `ShowSetting_EQ` | Show EQ preset list |
| 0x00228274 | `ShowSetting_SoundCheck` | Show SoundCheck toggle |
| 0x00228410 | `ShowSetting_Clicker` | Show clicker toggle |
| 0x00228424 | `ShowSetting_DateAndTime` | Show date/time settings |
| 0x0022843C | `ShowSetting_SortBy` | Show sort preferences |
| 0x00228450 | `ShowSetting_ClassicUI` | Show Classic UI toggle |
| 0x00228468 | `ShowSetting_Language` | Show language selector |
| 0x00228480 | `ShowSetting_Legal` | Show legal/license info |
| 0x00228494 | `ShowSetting_ResetAll` | Show reset confirmation |

---

## 2. Hidden/Disabled Features

### 2.1 Demo Mode

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001892F4 | `TCDemoMode` | Hidden | Retail store demo controller - cycles through features automatically |

### 2.2 Debug / Test Menus

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001E04F4 | `TSilverCntlrTestAppCntlr` | Hidden | Test app controller in Silver framework |
| 0x001E0518 | `TSilverCntlrTestCntlr` | Hidden | Generic test controller |
| 0x007395C2 | `Debug_MainMenu_Screen` | Hidden | Debug main menu screen definition |
| 0x007395DB | `Debug_MainMenu_Screen_Default` | Hidden | Default layout for debug menu |
| 0x007FF6A0 | `Test Menu Item (has no ID)` | Hidden | Placeholder test menu entry |
| 0x001443A0 | `MP3ExampleTask` | Hidden | MP3 decode test/example task |
| 0x0090A214 | `WaveFileDebugTask` | Hidden | WAV file debug/test task |

### 2.3 RTXCbug Debugger

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002BFA1D | `** RTXCbug -` | Hidden | RTXCbug debugger banner |
| 0x002BFA60 | `X - Exit RTXCbug` | Hidden | Exit debug command |
| 0x002BFA75 | `RTXCbug>` | Hidden | Debug shell prompt |
| 0x002C0451 | `RTXCbug - RTXC Objects>` | Hidden | RTOS object inspector prompt |
| 0x002D9D24 | `X - Exit RTXCbug` | Hidden | Exit from RTXCbug (second ref) |
| 0x003947E1 | `Re-entering RTXCbug mode` | Hidden | Re-entry message |
| 0x00394719 | `$RTXCbug>` | Hidden | Alternative prompt (task manager) |

### 2.4 RTXCbug Commands

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x002D9C08 | `T - Tasks` | List all RTXC tasks |
| 0x002D9C14 | `M - Mailboxes` | List mailbox objects |
| 0x002D9C24 | `Q - Queues` | List queue objects |
| 0x002D9C30 | `R - Resources` | List resource objects |
| 0x002D9C40 | `S - Semaphores` | List semaphore objects |
| 0x002D9C50 | `C - Clock / Timers` | Show clock/timer state |
| 0x002D9C64 | `K - Stack Limits` | Show stack usage |
| 0x002D9C78 | `Z - Zero Partition/Queue/Resource Statistics` | Reset statistics |
| 0x002D9CA8 | `$ - Enter Task Manager Mode` | Enter task manager |
| 0x002D9CC8 | `# - Task Registers` | Show task registers |
| 0x002D9CDC | `G - Go to Multitasking Mode` | Resume OS |
| 0x002D9CFC | `H - Help` | Show help |
| 0x002D9D08 | `U - Return To Main Menu` | Return to main menu |
| 0x00394760 | `C - Change task priority` | Change priority (task mgr) |
| 0x003947C4 | `X - Exit Task Manager Mode` | Exit task manager |

### 2.5 Unit Test / Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002650C0 | `Channel Reserved` | Hidden | Reserved logging channel |
| 0x002650D4 | `Channel AppBoot` | Hidden | Application boot events |
| 0x002650E4 | `Channel BufferedSongReading` | Hidden | Song buffer logging |
| 0x00265100 | `Channel PrefsWriting` | Hidden | Preferences write events |
| 0x00265118 | `Channel GeneralUserExperience` | Hidden | UX event logging |
| 0x00265138 | `Channel PlayFromDisk` | Hidden | Disk playback logging |
| 0x00265150 | `Channel CacheSpinupDrive` | Hidden | Drive cache spinup |
| 0x0026516C | `Channel TestLogging` | Hidden | Test/debug logging |
| 0x00265180 | `Channel AppFileLoading` | Hidden | App file load events |
| 0x00265198 | `Channel VCardReading` | Hidden | VCard import logging |
| 0x002651B0 | `Channel LongSongScanning` | Hidden | Long song scan events |
| 0x00265224 | `Channel VoiceRecording` | Hidden | Voice memo recording |
| 0x0026523C | `Channel PhotoImporting` | Hidden | Photo import events |
| 0x00265254 | `Channel Notes` | Hidden | Notes app logging |
| 0x00265264 | `Channel PhotoFileManagement` | Hidden | Photo file management |
| 0x00265280 | `Channel DiskMode` | Hidden | Disk mode events |
| 0x00265294 | `Channel Firewire` | Hidden | FireWire events |
| 0x002652A8 | `Channel USB` | Hidden | USB events |
| 0x002652B4 | `Channel UnitTests` | Hidden | Unit test channel |
| 0x002652C8 | `Channel FreeSpaceCache` | Hidden | Free space tracking |
| 0x002652E0 | `Channel OnTheGoFileMgmt` | Hidden | OTG playlist management |
| 0x002652F8 | `Unknown Disk Channel` | Hidden | Unknown disk channel |

---

## 3. Audio System

### 3.1 MeCCA Framework (Media Codec/Container Architecture)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000067EF | `MeCCADecode` | Known | Main MeCCA decoder entry |
| 0x00168A34 | `MeCCA_RecordingBuffer` | Known | Recording input buffer |
| 0x00197BE0 | `MeCCA_PCM_Output.wav` | Hidden | PCM output debug file |
| 0x001B1678 | `MeCCA_MediaPlayer` | Known | Main media player interface |
| 0x001BC3C8 | `MeCCA_VideoBufferMgr` | Known | Video buffer manager |
| 0x001BC5D0 | `MeCCAVideoDecode` | Known | Video decoder component |
| 0x001F5470 | `MeCCAIOTask` | Known | MeCCA I/O task thread |
| 0x004007B4 | `MeCCARecordingTask` | Known | Recording task thread |
| 0x00917268 | `MeCCA_StreamCache` | Known | Stream caching layer |

### 3.2 Supported Audio Codecs/Formats

| Offset | Evidence | Format | Classification |
|--------|----------|--------|----------------|
| 0x000A6ACF | `alac: bit depth = %d, pb = 0x%X, mb = 0x%X, kb = 0x%X` | Apple Lossless (ALAC) | Known |
| 0x000A8D94 | `adrmmp4a` | AAC (DRM protected) | Known |
| 0x0015592C | `alacmp4v@KL` | ALAC + MP4 video | Known |
| 0x00161F8C | `RIFFWAVEfmt data` | WAV/RIFF PCM | Known |
| 0x00150F8C | `AppleLossless` | Apple Lossless capability flag | Known |
| 0x00150EDC | `AudioCodecs` | Codec capability query | Known |
| 0x00150FB8 | `Audible` | Audible audiobook format | Known |
| 0x007FC700 | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON` | MP3 | Known |
| 0x001C26A0 | `tkhdedtselstmdiamdhdminfstblstsdstcoco64stscstszsttsstssdrmidrms` | MP4/M4A container atoms | Known |
| 0x001C3FD0 | `elsttkhdmdhdstsdsttsstszstscstcomp4aalac` | MP4A + ALAC atoms | Known |
| 0x00155910 | `tx3gdrmsp608aavdmp4aesdsX{` | DRM + text + AAC atoms | Known |
| 0x0092B578 | `ERROR: unknownCodec loaded !!!` | Unknown codec error handler | Hidden |

**Confirmed Supported Formats:** MP3, AAC, Protected AAC (FairPlay), Apple Lossless (ALAC), WAV/AIFF, Audible (.aa/.aax), MP4 audio

### 3.3 Equalizer Presets

| Offset | Preset Name | Type |
|--------|-------------|------|
| 0x000A02B0 | `Bass Booster` | EQ Preset |
| 0x000A02C0 | `Bass Reducer` | EQ Preset |
| 0x000A0364 | `Treble Booster` | EQ Preset |
| 0x000A0374 | `Treble Reducer` | EQ Preset |
| 0x009BFF1A | `EQMenu_Deep_String` | EQ Preset |
| 0x009BFF70 | `EQMenu_HipHop_String` | EQ Preset |
| 0x009BFF85 | `EQMenu_Pop_String` | EQ Preset |
| 0x009BEF7A | `ContextualMenu_Audiobook_Normal_String` | Audiobook speed |
| 0x009BEE77 | `EQMenu_Rock_String` | EQ Preset |
| 0x009BEF12 | `EQMenu_Classical_String` | EQ Preset |
| 0x009BF9A1 | `EQMenu_Latin_String` | EQ Preset |
| 0x009BFE53 | `EQMenu_Piano_String` | EQ Preset |
| 0x009BDC6C | `EQMenu_Dance_String` | EQ Preset |
| 0x009BDE08 | `EQMenu_Lounge_String` | EQ Preset |
| 0x009BD210 | `EQMenu_Electronic_String` | EQ Preset |
| 0x009BD30A | `EQMenu_Acoustic_String` | EQ Preset |
| 0x009BD921 | `EQMenu_SpokenWord_String` | EQ Preset |
| 0x009C0F0D | `EQMenu_SmallSpeakers_String` | EQ Preset |
| 0x009C1007 | `EQMenu_Loudness_String` | EQ Preset |
| 0x009C1532 | `EQMenu_Flat_String` | EQ Preset |
| 0x009BE761 | `EQMenu_Off_String` | EQ Preset |
| 0x009C24F7 | `EQMenu_Jazz_String` | EQ Preset |
| 0x009C00CC | `EQMenu_TrebleReducer_String` | EQ Preset |
| 0x009C00E8 | `EQMenu_BassReducer_String` | EQ Preset |
| 0x009C02A4 | `EQMenu_TrebleBooster_String` | EQ Preset |
| 0x009C02C0 | `EQMenu_VocalBooster_String` | EQ Preset |
| 0x009C02DB | `EQMenu_BassBooster_String` | EQ Preset |
| 0x009B8087 | `Settings_EQ_RandB_Image` | EQ Preset (R&B) |

### 3.4 Audio Processing Chain

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x000D2FDC | `TrackData::EQCache` | EQ data cache per track |
| 0x0010C2A0 | `HandleEQSettingSelected` | EQ preset selection handler |
| 0x00135D74 | `TCEQSetting` | EQ setting controller class |
| 0x0013F7F8 | `HandleEQ` | Main EQ handler |
| 0x00150F10 | `MaximumSampleRate` | Maximum supported sample rate |
| 0x00150F60 | `VariableBitRate` | VBR support indicator |
| 0x00802A4C | `Use Original Volume Level` | SoundCheck off description |
| 0x00802A68 | `Normalize Volume Across All Songs` | SoundCheck on description |
| 0x00A1B59C | `SoundEffect` | Sound effect system |
| 0x00988D20 | `21SoundEffectDescriptor` | RTTI: sound effect descriptor class |

### 3.5 DRM System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00150EB0 | `AppleDRMVersion` | Known | DRM version capability |
| 0x00150F50 | `AppleDRM` | Known | DRM support flag |
| 0x00152510 | `AppleVideoDRM` | Known | Video DRM flag |
| 0x000A8D94 | `adrmmp4a` | Known | DRM-protected AAC atom |
| 0x001C26A0 | `...drmidrms` | Known | DRM atoms in MP4 container |
| 0x00202E8C | `drmttx3g` | Known | DRM + subtitle text track |
| 0x009B487F | `DRMLevel` | Known | DRM protection level |
| 0x009DA728 | `$Apple FairPlay Certificate Authority0` | Known | FairPlay CA cert |
| 0x009DAAAD | `&Apple FairPlay Certification Authority0` | Known | FairPlay cert authority |
| 0x00A0F655 | `&Apple FairPlay Certification Authority0` | Known | Second FairPlay CA ref |
| 0x00A0F6CB | `Apple FairPlay1402` | Known | FairPlay version string |
| 0x000985D8 | `iPod_Control/iTunes/iTunesDB.p7b` | Known | PKCS#7 signed iTunes DB |

---

## 4. UI Architecture - "Silver" Framework

### 4.1 Core Silver Controllers (TSilver*)

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x000AC06C | `TSilverCntlr` | Base controller class |
| 0x000AC0C8 | `TSilverMainMediaListCntlr_Main` | Main media list root |
| 0x000AC0F0 | `TSilverMainMediaListCntlr_Music` | Music top-level list |
| 0x000AC118 | `TSilverMainMediaListCntlr_Videos` | Videos top-level list |
| 0x000AC73C | `TSilverGlobalCntlr` | Global/system controller |
| 0x000AC758 | `TSilverTrainerCntlr` | Nike+ trainer controller |
| 0x001725D0 | `TSilverSettingsMenuListCntlr` | Settings menu list |
| 0x0017271C | `TSilverSettingsVideoCntlr` | Video settings controller |
| 0x001E04F4 | `TSilverCntlrTestAppCntlr` | Test app (hidden) |
| 0x001E0518 | `TSilverCntlrTestCntlr` | Test controller (hidden) |

### 4.2 Media List Controllers

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x000AC144 | `TSilverMediaListCntlr_Songs` | Song list |
| 0x000AC168 | `TSilverMediaListCntlr_Albums` | Album list |
| 0x000AC190 | `TSilverMediaListCntlr_Artists` | Artist list |
| 0x000AC1B8 | `TSilverMediaListCntlr_Genres` | Genre list |
| 0x000AC1E0 | `TSilverMediaListCntlr_Composers` | Composer list |
| 0x000AC208 | `TSilverMediaListCntlr_Podcasts` | Podcast list |
| 0x000AC230 | `TSilverMediaListCntlr_PodcastEpisodes` | Podcast episodes |
| 0x000AC260 | `TSilverMediaListCntlr_iTunesU` | iTunes U list |
| 0x000AC288 | `TSilverMediaListCntlr_iTunesUEpisodes` | iTunes U episodes |
| 0x000AC2B8 | `TSilverMediaListCntlr_Audiobooks` | Audiobook list |
| 0x000AC2E4 | `TSilverMediaListCntlr_AudiobookChapters` | Audiobook chapters |
| 0x000AC314 | `TSilverMediaListCntlr_TVShows` | TV show list |
| 0x000AC33C | `TSilverMediaListCntlr_TVSeasons` | TV season list |
| 0x000AC364 | `TSilverMediaListCntlr_TVEpisodes` | TV episode list |
| 0x000AC390 | `TSilverMediaListCntlr_Movies` | Movie list |
| 0x000AC3B8 | `TSilverMediaListCntlr_Playlists` | Playlist list |
| 0x000AC3E0 | `TSilverMediaListCntlr_NestedPlaylists` | Nested/folder playlists |
| 0x000AC410 | `TSilverMediaListCntlr_VideoPlaylists` | Video playlists |
| 0x000AC5AC | `TSilverMediaListCntlr_NestedVideoPlaylists` | Nested video playlists |
| 0x000AC5E0 | `TSilverMediaListCntlr_PlaylistChooser` | Playlist chooser |
| 0x000AC610 | `TSilverMediaListCntlr_Rentals` | Rental videos list |
| 0x000AC638 | `TSilverMediaListCntlr_Genius` | Genius playlist |
| 0x000AC660 | `TSilverMediaListCntlr_GeniusMixes` | Genius Mixes list |

### 4.3 Calendar/PIM Controllers

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x00127E04 | `TSilverCalendarCntlr_CalendarMenu` | Calendar menu |
| 0x00127E30 | `TSilverCalendarCntlr_MonthViewer` | Month view |
| 0x00127E5C | `TSilverCalendarCntlr_DayViewer` | Day view |
| 0x00127E84 | `TSilverCalendarCntlr_EventViewer` | Event detail |
| 0x00127EB0 | `TSilverCalendarCntlr_ToDoList` | To-do list |
| 0x00127ED8 | `TSilverCalendarCntlr_ToDoDispatcher` | To-do dispatcher |
| 0x00127F04 | `TSilverCalendarCntlr_Alarm` | Alarm controller |

### 4.4 Other TC Controllers

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x000AC084 | `TCExtrasMenu` | Extras menu controller |
| 0x000AC09C | `TCGameScreen` | Game display controller |
| 0x000AC0B4 | `TCGamesMenu` | Games menu controller |
| 0x000AC68C | `TCRentalNotification` | Rental expiry notification |
| 0x000AC6AC | `TCRentalInfo` | Rental info display |
| 0x000AC6C4 | `TCRentalConfirmDelete` | Rental delete confirm |
| 0x000AC6E4 | `TCRentalDispatcher` | Rental action dispatcher |
| 0x000AC700 | `TContextualMenuCntlr` | Context menu (hold center) |
| 0x000AC720 | `TGeniusLoadingCntlr` | Genius loading screen |
| 0x0010402C | `TCSlideshowLCD` | Photo slideshow (LCD) |
| 0x00104044 | `TCSlideshowTVOut` | Photo slideshow (TV) |
| 0x00104060 | `TCSlideshow_TVOutAsk` | TV out prompt |
| 0x00104080 | `TCSlideshow_TVOutCableConnect` | Cable connect prompt |
| 0x001040A8 | `TPhotosBrowseCntlr` | Photo browser |
| 0x001040C4 | `TPhotosBrowseTransitionCntlr` | Photo transition |
| 0x001040EC | `TPhotosMenuCntlr` | Photos menu |
| 0x00104108 | `TPhotosSettingsCntlr` | Photo settings |
| 0x001040EC | `TPhotosSettingsCntlr_Transitions` | Slideshow transitions |
| 0x00104154 | `TPhotosSettingsCntlr_Duration` | Slideshow duration |
| 0x0010417C | `TPhotosSettingsSlideshowPlaylistCntlr` | Slideshow music |
| 0x00104494 | `TSearchCntlr` | Search controller |
| 0x0011CD4C | `TRadioCntlr` | FM Radio controller |
| 0x0012F880 | `TCRemoteUI` | Remote control UI |
| 0x0012F894 | `TCUnsupported` | Unsupported feature screen |
| 0x00135D60 | `TCSpeakers` | Speaker output controller |
| 0x00144430 | `TChargingModeCntlr` | Charging display |
| 0x0014444C | `TChargingModeLowPowerCntlr` | Low power charging |
| 0x0015F23C | `TCSportTimer` | Sport timer |
| 0x0015F254 | `TCSportTimerMenu` | Sport timer menu |
| 0x0015F270 | `TCSportTimerSessionScreen` | Timer session |
| 0x0015F294 | `TCSportTimerChosenDispatcher` | Timer action dispatch |
| 0x00160644 | `TCVoiceMemos` | Voice memo main |
| 0x0016065C | `TCVoiceMemosMenu` | Voice memo menu |
| 0x00160678 | `TCVoiceMemosMainMenu` | Voice memo main menu |
| 0x00160698 | `TCVoiceMemosPlayback` | Voice memo playback |
| 0x001606B8 | `TCVoiceMemosContextMenu` | Voice memo context |
| 0x001606D8 | `TCVoiceMemosAlert` | Voice memo alerts |
| 0x001727A0 | `TCFirstBoot` | First boot/setup wizard |
| 0x00172740 | `TCDateTimeScreen` | Date/time setting |
| 0x0017275C | `TCTimeZoneScreen` | Time zone selection |
| 0x001B243C | `TCAddressViewerMainMenu` | Address book menu |
| 0x001B245C | `TCAddressViewerDetails` | Contact details |
| 0x001B247C | `TCAddressViewerPartialLoad` | Partial contact load |
| 0x001B24A0 | `TCAddressViewerMainDispatcher` | Address dispatch |
| 0x0027EDD0 | `TCoverFlowCntlr` | Cover Flow controller |
| 0x0027FC00 | `TC_LockDialog` | Lock dialog controller |
| 0x0027FC18 | `TC_LockScreen` | Lock screen controller |
| 0x0027FC30 | `TC_LockediPod` | Locked iPod screen |
| 0x0027FC48 | `TC_VolumeLimitLockScreen` | Volume lock screen |
| 0x0027FC6C | `TCLockChosenDispatcher` | Lock action dispatcher |
| 0x0028582C | `TCClock` | Clock display controller |
| 0x0028583C | `TCClockCityMenu` | Clock city selector |
| 0x00285854 | `TCClockRegionMenu` | Clock region selector |
| 0x00285870 | `TCAlarmMenu` | Alarm list |
| 0x00285884 | `TCSleepTimerMenu` | Sleep timer menu |
| 0x002858A0 | `TCAlarmPropertiesMenu` | Alarm properties |
| 0x002858C0 | `TCAlarmPropertiesFrequencyMenu` | Alarm frequency |
| 0x002858E8 | `TCAlarmPropertiesLabelMenu` | Alarm label |
| 0x0028590C | `TCAlarmPropertiesSoundMenu` | Alarm sound |
| 0x00285930 | `TCAlarmDatePicker` | Alarm date picker |
| 0x0028594C | `TCAlarmTriggered` | Alarm triggered screen |
| 0x0028C8F4 | `TCNotesDispatcher` | Notes dispatcher |
| 0x0028C910 | `TCNotesLoading` | Notes loading screen |

### 4.5 Key Screen Names

| Offset | Screen Name | Description |
|--------|-------------|-------------|
| 0x00743B5D | `Radio_Screen` | FM Radio screen |
| 0x009BAFE8 | `NowPlaying_Screen_Volume` | Now Playing volume overlay |
| 0x009BB001 | `Radio_Screen_Volume` | Radio volume overlay |
| 0x009BB015 | `TVOutSlideshow_Screen_Volume` | TV out slideshow volume |
| 0x009BB032 | `NowPlaying_Screen_Video_Volume` | Video volume overlay |
| 0x009BB051 | `NowPlaying_Screen_Video_TVOut_Volume` | TV out video volume |
| 0x009C4387 | `SettingsMenus_EQ_Screen` | EQ selection screen |
| 0x009C4483 | `Alarms_Set_Alarm_Sound_Screen` | Alarm sound picker |
| 0x009C468E | `VolumeLimitLock_Screen` | Volume lock code entry |
| 0x009C4AC3 | `SettingsMenus_AudiobookSettings_Screen` | Audiobook settings |
| 0x009C4B61 | `MediaLists_Audiobooks_Screen` | Audiobooks list screen |
| 0x009C4CF4 | `MediaLists_AudiobookChapters_Screen` | Chapters screen |
| 0x009C4E51 | `SettingsMenus_VolumeLimit_Screen` | Volume limit screen |
| 0x009C50E4 | `Alarms_Set_Alarm_Frequency_Screen` | Alarm repeat picker |
| 0x009C44B3 | `DiskMode_Screen` | Disk mode screen |
| 0x009C451B | `Game_Screen` | Game display screen |
| 0x009C4669 | `Game_Running_Screen` | Game running state |
| 0x009C4955 | `Game_Signing_Error_Screen` | Game signature error |
| 0x009C496F | `Game_Version_Error_Screen` | Game version error |
| 0x009C4989 | `Game_Unknown_Error_Screen` | Game unknown error |
| 0x009C49B7 | `Game_Memory_Error_Screen` | Game memory error |
| 0x009C4FB5 | `Games_Menu_Screen` | Games list menu |
| 0x0075C3C7 | `CoverFlow_Screen_Default` | Cover Flow screen |
| 0x0075BDF3 | `CoverFlow_Screen_Backside` | Album back info |

### 4.6 Key Layouts

| Offset | Layout Name | Description |
|--------|-------------|-------------|
| 0x009CBD68 | `SettingsInfo_Template_SoundCheck_Layout` | SoundCheck info |
| 0x009CBD90 | `SettingsMenus_SoundCheck_Layout` | SoundCheck menu |
| 0x009CBF8F | `Settings_VolumeLimitControl_Layout` | Volume limit slider |
| 0x009CC1C6 | `SettingsMenus_DialogNotice_Audiobooks_Layout` | Audiobook notice |
| 0x009CC21B | `SettingsMenus_Audiobooks_Layout` | Audiobook menu |
| 0x009CC361 | `VolumeLimitLock_Screen_Incorrect_Layout` | Wrong code layout |
| 0x009CC4CE | `SettingsMenus_VolumeLimit_Layout` | Volume limit layout |
| 0x009CC7A8 | `Settings_EQMenu_Layout` | EQ menu layout |
| 0x007E5E5E | `DiskMode_ScreenLayout_Disconnected` | Disk disconnected |
| 0x007E5EC9 | `DiskMode_ScreenLayout_Loading` | Disk loading |
| 0x007E5F35 | `DiskMode_ScreenLayout_Synchronizing` | Disk syncing |
| 0x007E5FA7 | `DiskMode_ScreenLayout_Connected` | Disk connected |

---

## 5. Game System

### 5.1 Game Loading Infrastructure

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00098648 | `gamedata_RW` | Known | Game read/write data partition |
| 0x00098654 | `gamestats_WO` | Known | Game statistics write-only |
| 0x00098664 | `gamedata_ShareRW` | Known | Shared game data (R/W) |
| 0x00098678 | `games_RO` | Known | Game binaries read-only |
| 0x00152408 | `GamesPlatformID` | Known | Platform identification for games |
| 0x00152418 | `GamesPlatformVersion` | Known | Platform version for games |

### 5.2 Game UI Screens

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x009C451B | `Game_Screen` | Known | Main game rendering screen |
| 0x009C4669 | `Game_Running_Screen` | Known | Active game screen |
| 0x009C4955 | `Game_Signing_Error_Screen` | Known | Code signing failure |
| 0x009C496F | `Game_Version_Error_Screen` | Known | Version mismatch error |
| 0x009C4989 | `Game_Unknown_Error_Screen` | Known | Unknown game error |
| 0x009C49B7 | `Game_Memory_Error_Screen` | Known | Insufficient memory |
| 0x009C4FB5 | `Games_Menu_Screen` | Known | Games browser menu |
| 0x009BC1AF | `Game_Daisy_Template` | Known | Game loading template ("daisy" spinner) |
| 0x009BB735 | `Game_Error_Message_Template` | Known | Error message template |

### 5.3 Game Error Strings

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x009C0459 | `Games_Launch_Error_String` | Failed to launch game |
| 0x009C0473 | `Games_Version_Error_String` | Game version incompatible |
| 0x009C048E | `Games_Unknown_Error_String` | Unknown game error |
| 0x009C04A9 | `Games_Error_String` | Generic game error |
| 0x009BF662 | `Games_Reinstall_String` | Reinstall prompt |
| 0x009BFD91 | `Games_MoreInfo_String` | More info prompt |
| 0x009BFBF2 | `Games_Resource_Action_String` | Resource action |
| 0x009BFC0F | `Games_Version_Action_String` | Version action |
| 0x009C1658 | `Games_Format_String` | Game format/platform info |
| 0x009BE871 | `Games_Loading_String` | "Loading" text |

### 5.4 Game Signing Verification

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0075C9C1 | `controller.ShowSigningError1` | Known | Trigger signing error screen |
| 0x0075CA07 | `Game_Signing_Error_Screen` | Known | Signing error layout |
| 0x000985D8 | `iPod_Control/iTunes/iTunesDB.p7b` | Known | PKCS#7 signed database |
| 0x00063A68 | `signature has problems, re-make with post SSLeay045` | Known | OpenSSL signature error |
| 0x002737F8 | `http://www.w3.org/2000/09/xmldsig#rsa-sha1` | Known | RSA-SHA1 XML signature |
| 0x00273874 | `http://www.w3.org/2000/09/xmldsig#sha1` | Known | SHA1 XML signature |
| 0x003EE91F | `SERIALVERIFIER` | Hidden | Serial number verifier |
| 0x003EE930 | `RESISTORVERIFIER` | Hidden | Hardware resistor verifier |

### 5.5 Game Play Tracking

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x009B7B68 | `Games_PlayCount_Once` | Played once |
| 0x009C6A14 | `Games_PlayCount_Never` | Never played |
| 0x009CFB54 | `Games_PlayCount_Many` | Played multiple times |
| 0x009BB566 | `Games_Background_Template` | Game background |
| 0x009BC57D | `Games_Background_Template_Loading` | Loading background |
| 0x009BC5CB | `Games_Background_Color_Loading` | Loading color |
| 0x009BC880 | `Games_Background_Template_Running` | Running background |
| 0x009BC8D6 | `Games_Background_Color_Running` | Running color |
| 0x009BA5C0 | `Games_Proxy_Image` | Game proxy/placeholder image |
| 0x009B8807 | `Games_Preview_Template_Image` | Game preview image |

---

## 6. Database System

### 6.1 SQLite Integration

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00906263 | `SQLite format 3` | Known | SQLite file header magic |
| 0x009D06D5 | `SQLite_iPod_VFS` | Known | Custom VFS for iPod storage |
| 0x0036FA1C | `sqlite_sequence` | Known | Auto-increment sequence table |
| 0x00376A8C | `CREATE TABLE %Q.sqlite_sequence(name,seq)` | Known | Sequence table creation |
| 0x002DCA04 | `sqlite_stat1` | Known | SQLite statistics table |
| 0x002DCA14 | `CREATE TABLE %Q.sqlite_stat1(tbl,idx,stat)` | Known | Stats table DDL |
| 0x00376BF4 | `file is encrypted or is not a database` | Known | Encryption check |
| 0x00376D94 | `disk I/O error` | Known | Disk I/O error string |
| 0x00376DC8 | `database or disk is full` | Known | Full disk error |
| 0x00376DA4 | `database disk image is malformed` | Known | Corruption error |
| 0x00237040 | `%s/sqlite_` | Known | SQLite path format |

### 6.2 iTunesDB System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00069108 | `iPod_Control\iTunes\iTunesDB` | Known | Main iTunes database path |
| 0x000985D8 | `iPod_Control/iTunes/iTunesDB.p7b` | Known | PKCS#7 signed DB |
| 0x0013C650 | `iPod_Control\iTunes\iTunesDB` | Known | Second iTunesDB ref |
| 0x000B55CC | `iPod_Control\iTunes\Play Counts` | Known | Play counts file |
| 0x000BF64C | `iPod_Control\iTunes\Play Counts` | Known | Play counts (2nd ref) |
| 0x000BC1E4 | `TDBEtadb` | Known | Extended track attribute DB |
| 0x003EDC9C | `iTunesDB` | Known | iTunesDB reference |
| 0x0011F4EC | `iTunes Image DB.itdb` | Known | Artwork database |
| 0x0014AF30 | `iTunes Image DB.itdb` | Known | Artwork DB (2nd ref) |
| 0x000690CC | `rtSPhotos\Photo Database` | Known | Photo database |

### 6.3 Database Metrics (Usage Log)

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x002663CD | `Music database size: %d KB` | Track DB size |
| 0x002663ED | `Music database num songs: %d` | Number of songs |
| 0x0026640D | `Photo database size: %d KB` | Photo DB size |
| 0x0026642D | `Photo database num photos: %d` | Number of photos |
| 0x0026644D | `Album art database size: %d KB` | Artwork DB size |
| 0x00266471 | `Album art num images: %d` | Number of artworks |
| 0x00267024 | `Music Database Size` | Event log field |
| 0x00267038 | `Photo Database Size` | Event log field |
| 0x0026704C | `Artwork Database Size` | Event log field |

### 6.4 Data Caches

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x0005EDD4 | `SearchCriteria::DataCache` | Search criteria cache |
| 0x000CBAA4 | `Playlists::NameCache` | Playlist name cache |
| 0x000CBABC | `Playlists::DescCache` | Playlist description cache |
| 0x000D2F38 | `TrackData::ArtistComposerCache` | Artist/composer cache |
| 0x000D2F58 | `TrackData::DisplayArtistCache` | Display artist cache |
| 0x000D2F78 | `TrackData::AlbumCache` | Album name cache |
| 0x000D2FAC | `TrackData::GenreCache` | Genre name cache |
| 0x000D3044 | `TrackData::PodcastURLCache` | Podcast URL cache |
| 0x000ABD08 | `playlistPersistentID` | Playlist persistent ID key |

---

## 7. Hardware Interfaces

### 7.1 Storage / Disk Interface

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00163D64 | `ATAWorkLoopTask` | Known | ATA command work loop |
| 0x00163D78 | `ATAWorkLoopIRQTask` | Known | ATA interrupt handler task |
| 0x002AD900 | `DiskMgrTask` | Known | Disk manager task |
| 0x00153B9C | `DiskReaderTask` | Known | Disk reader task |
| 0x0013BEB4 | `NAND FLASH DRIVE` | Known | Flash storage identifier |
| 0x0036A9A8 | `MMC init failed` | Known | MMC initialization error |
| 0x0036A9BC | `CE-ATA init failed` | Known | CE-ATA initialization error |
| 0x000756D0 | `cI: CE-ATA signature missing (%x,%x)` | Known | CE-ATA signature check |
| 0x000755EC | `cI: Set drive to MMC high speed failed` | Known | MMC speed negotiation |
| 0x00075614 | `cI: card failed HS_TIMING SWITCH` | Known | High-speed timing fail |
| 0x0007563C | `cI: SWITCH to %d-bit bus width failed` | Known | Bus width negotiation |
| 0x000C5220 | `CGLE: Retry failed in getting ATA STATUS registers` | Known | ATA status retry |
| 0x000CE7FC | `cIC12: ATA Status Error! Could not get error code.` | Known | ATA error handler |
| 0x000CE830 | `cIC12: ATA Status Error! Error code (0x%2x)` | Known | ATA error code |

### 7.2 USB Interface

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00149868 | `USBDeviceTask` | Known | USB device mode task |
| 0x002ADA03 | `USBAudioTask` | Known | USB audio class task |
| 0x0016126C | `USBCUSBS` | Known | USB composite/storage |
| 0x001617F8 | `USB MSC` | Known | USB Mass Storage Class |
| 0x002AD8EC | `USBPowerSense` | Known | USB power detection |
| 0x009B46E0 | `USBRoleManager` | Known | USB OTG role manager |
| 0x003EE8CA | `USB_GRANT` | Hidden | USB grant semaphore |
| 0x003EE8DB | `USB_RESP_INIT` | Hidden | USB responder init |
| 0x003EE8EC | `USB_RESPONDER` | Hidden | USB responder task |

### 7.3 Serial / UART

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00078A9C | `iPod Serial Number` | Known | Device serial number |
| 0x000AF67C | `serial` | Known | Serial port reference |
| 0x00151D48 | `SerialNumber` | Known | Serial number capability |
| 0x003EE91F | `SERIALVERIFIER` | Hidden | Serial verification task |

### 7.4 I2C / Hardware Control

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x003EE8B9 | `I2C_MASTER` | Hidden | I2C master bus |
| 0x003EE875 | `GPIO_REG_WRITE` | Hidden | GPIO register write |
| 0x003EE886 | `GPIO_INT_INIT` | Hidden | GPIO interrupt init |
| 0x003EE8FD | `DISKPWRMGRSEND` | Hidden | Disk power manager |
| 0x003EE90E | `PIEZOMGRSEND` | Hidden | Piezo buzzer manager |
| 0x003EE930 | `RESISTORVERIFIER` | Hidden | Accessory resistor check |

### 7.5 Input Devices

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002AD8C0 | `TouchwheelTask` | Known | Click wheel input task |
| 0x002AD910 | `HoldSwitchTask` | Known | Hold switch monitoring |
| 0x002AD924 | `MikeyTask` | Known | Apple Mikey chip (remote) |
| 0x002AD934 | `TopPlugTask` | Known | Top plug detection |
| 0x002AD944 | `HPhoneDetTask` | Known | Headphone detect task |
| 0x002AD9BC | `LowBattDebounceTask` | Known | Low battery debounce |

### 7.6 Power Management

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x002AD8AC | `FirewireTask` | FireWire (charging) task |
| 0x002AD9E4 | `AlarmTask` | Alarm wakeup task |
| 0x00267004 | `Begin Charging` | Charge start event |
| 0x00267014 | `Stop Charging` | Charge stop event |
| 0x007FFE74 | `Low Battery` | Low battery warning string |
| 0x007FFE80 | `Connect to Power` | Connect power prompt |

### 7.7 Disk Mode Screens

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x009BD04A | `DiskMode__String` | Disk mode label |
| 0x009BD861 | `DiskMode_Connected_String` | "Connected" |
| 0x009BE824 | `DiskMode_Syncing_String` | "Syncing" |
| 0x009BE83C | `DiskMode_Loading_String` | "Loading" |
| 0x009BEC92 | `DiskMode_Synchronizing_String` | "Synchronizing" |
| 0x009C1690 | `DiskMode_UseiTunesToEject_String` | Eject instruction |
| 0x009C16D2 | `DiskMode_OKToDisconnect_String` | Safe to disconnect |
| 0x009C1712 | `DiskMode_EjectingYouMayDisconnect_String` | Ejecting message |
| 0x009C190B | `DiskMode_PleaseWait_String` | Please wait |
| 0x009C1926 | `DiskMode_EjectingPleaseWait_String` | Ejecting please wait |
| 0x001596D0 | `SwitchToSynchronizing` | Transition to sync mode |

---

## 8. RTOS Architecture (RTXC)

### 8.1 RTXC Kernel Objects

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x002BF57D | `** Clock Snapshot **` | Clock state dump header |
| 0x002BF598 | `Clock rate is...Hz, Tick interval is...ms` | Clock config display |
| 0x002BF5E0 | `Maximum of...timers` | Timer pool info |
| 0x002BF64C | `Time  Cyclic  Task  Timer  Object` | Timer list columns |
| 0x002BF860 | `Timer` | Timer type label |
| 0x002BF878 | `Delay` | Delay timer type |
| 0x002BF890 | `Semaphore` | Semaphore type |
| 0x002BF8AC | `Mailbox` | Mailbox type |
| 0x002BF8F4 | `Resource` | Resource type |
| 0x002BFB99 | `** Mailbox Snapshot **` | Mailbox dump header |
| 0x002BFDDD | `** Queue Snapshot **` | Queue dump header |
| 0x002C0215 | `** Resource Snapshot **` | Resource dump header |
| 0x002C066D | `** Semaphore Snapshot **` | Semaphore dump header |
| 0x002C09A5 | `** Stack Snapshot **` | Stack dump header |
| 0x002C0CB9 | `** Task Snapshot **` | Task dump header |

### 8.2 Known RTXC Tasks

| Offset | Task Name | Description |
|--------|-----------|-------------|
| 0x000E8FBC | `HostOSTask` | Host OS main task |
| 0x00149868 | `USBDeviceTask` | USB device mode |
| 0x00153B9C | `DiskReaderTask` | Disk I/O reader |
| 0x00163D64 | `ATAWorkLoopTask` | ATA command loop |
| 0x00163D78 | `ATAWorkLoopIRQTask` | ATA IRQ handler |
| 0x0019CA7C | `GeniusMixesTask` | Genius playlist generation |
| 0x001B99F8 | `TMusicLoadingTask` | Music library load |
| 0x001F5470 | `MeCCAIOTask` | MeCCA codec I/O |
| 0x002AD8AC | `FirewireTask` | FireWire interface |
| 0x002AD8C0 | `TouchwheelTask` | Click wheel input |
| 0x002AD8D4 | `AudioOutStateTask` | Audio output state |
| 0x002AD900 | `DiskMgrTask` | Disk management |
| 0x002AD910 | `HoldSwitchTask` | Hold switch detect |
| 0x002AD924 | `MikeyTask` | Apple Mikey remote IC |
| 0x002AD934 | `TopPlugTask` | Top plug detect |
| 0x002AD944 | `HPhoneDetTask` | Headphone detection |
| 0x002AD9BC | `LowBattDebounceTask` | Battery debounce |
| 0x002AD9E4 | `AlarmTask` | Alarm management |
| 0x002ADA03 | `USBAudioTask` | USB Audio class |
| 0x004007B4 | `MeCCARecordingTask` | Voice recording |
| 0x003F49DC | `ArtworkLoadTask` | Album art loading |
| 0x003F80A8 | `TPodMediaPlayer Task` | Main media player |
| 0x00228AEC | `StreamCacheMassStorageManagerTimeOutTask` | Stream cache timeout |
| 0x00228C68 | `StreamCacheReadTask` | Stream cache reader |
| 0x001FC308 | `iMAAudioPromptMangThread` | iMA audio prompt |
| 0x00133458 | `SearchHelperThread` | Search helper |
| 0x0019CF6C | `KeyRepeatTimer` | Key repeat timing |

### 8.3 RTXC Task States (from RTXCbug output)

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x002C0D28 | `SUSPENDED` | Task suspended state |
| 0x002C0D34 | `INACTIVE` | Task inactive state |
| 0x002C0D40 | `Semaphore` | Waiting on semaphore |
| 0x002C0CD0 | `#  Name  Priority  CPUTime State` | Task snapshot columns |
| 0x002C09BC | `#  Task  TopOfStk  Size  Used  Spare` | Stack snapshot columns |
| 0x002C0A73 | `Worst case interrupt nesting =` | IRQ nesting metric |
| 0x002C0A3C | `RTXC Kernel...` | Kernel version display |

### 8.4 RTXC Mailbox/Queue/Semaphore Infrastructure

The firmware uses RTXC's full object model:
- **Tasks**: Independent threads of execution (17,721 function prologues detected)
- **Mailboxes**: Inter-task messaging (single message)
- **Queues**: FIFO message buffers (with depth tracking)
- **Semaphores**: Binary/counting synchronization
- **Resources**: Mutual exclusion locks
- **Timers**: One-shot and cyclic timers

---

## 9. Logging & Analytics System

### 9.1 Usage Log Framework

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00265C7C | `iPod Usage Stats` | Hidden | Usage statistics header |
| 0x00266F18 | `Flush Usage Log Data` | Hidden | Flush log to disk event |
| 0x0028BFE8 | `Notes/Activity.log` | Hidden | Activity log file path |
| 0x009B5001 | `LogActivity` | Hidden | Activity logging function |
| 0x002658A8 | `Max Events In Queue: %d` | Hidden | Event queue capacity |

### 9.2 Tracked Time Metrics

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x00265AB9 | `Total time the disk was running in the app: %d seconds` | Disk activity time |
| 0x00265AF5 | `As a percent of the total log time: %d%%` | Disk time percentage |
| 0x00265B21 | `As a percent of the total playback time: %d%%` | Disk vs playback |
| 0x00265B61 | `The disk was turned on %d %s` | Disk spinup count |
| 0x00265B91 | `Total time in DMIA: %d %s` | Time in disk mode |
| 0x00265BAD | `Time in DMIA as percent of total log time: %d%%` | DMIA percentage |
| 0x00265BE1 | `DMIA was entered %d %s` | DMIA entry count |
| 0x00265EDD | `Total time in light sleep: %d seconds` | Light sleep time |
| 0x00265F41 | `Light sleep was entered %d %s` | Light sleep count |
| 0x00265F61 | `Total time in deep sleep: %d seconds` | Deep sleep time |
| 0x00265F89 | `Deep sleep was entered %d %s` | Deep sleep count |
| 0x0026677D | `Total log length is %d seconds` | Total log duration |

### 9.3 Playback Statistics

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x002667A1 | `%d%% Playing Music` | Music play percentage |
| 0x002667B9 | `%d%% in DMIA` | DMIA percentage |
| 0x002667C9 | `%d%% Sleeping` | Sleep percentage |
| 0x002667D9 | `%d%% Idle` | Idle percentage |
| 0x002669E9 | `Total time playing: %d seconds` | Total play time |
| 0x00266A0D | `As percent of total log time: %d%%` | Play percentage |
| 0x00266A35 | `Average playback duration was %d seconds` | Avg session length |
| 0x00266A71 | `Playback was started %d %s` | Playback start count |
| 0x00266A91 | `Average navigation (Next/Prev) per playback duration: %d` | Skip rate |

### 9.4 Backlight Statistics

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x00266C45 | `Total time on: %d %s` | Backlight on time |
| 0x00266C5D | `As percent of total log time: %d%%` | Backlight percentage |
| 0x00266C95 | `It was turned on %d %s` | Backlight on count |

### 9.5 Logged Events

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x00265AA8 | `Disk Activity` | Disk activity event |
| 0x00266E2C | `Disk Spinup` | Disk spin up event |
| 0x00266E38 | `Disk Spindown` | Disk spin down event |
| 0x00266E48 | `Disk Obtain Access` | Disk access acquired |
| 0x00266E5C | `Disk Release Access` | Disk access released |
| 0x00266E70 | `Next Track` | Track advance (auto) |
| 0x00266E7C | `Next Track User Initiated` | Track advance (user) |
| 0x00266E98 | `Previous Track User Initiated` | Track back (user) |
| 0x00266EB8 | `Playback Begin` | Playback started |
| 0x00266EC8 | `Playback Resume` | Playback resumed |
| 0x00266ED8 | `Playback Pause` | Playback paused |
| 0x00266EE8 | `Playback Stop` | Playback stopped |
| 0x00266EF8 | `Backlight On` | Backlight on event |
| 0x00266F08 | `Backlight Off` | Backlight off event |
| 0x00266F98 | `Enter Disk Mode` | Disk mode entered |
| 0x00266FA8 | `Exit Disk Mode` | Disk mode exited |
| 0x00266FB8 | `Enter Light Sleep` | Light sleep entered |
| 0x00266FCC | `Exit Light Sleep` | Light sleep exited |
| 0x00266FE0 | `Enter Deep Sleep` | Deep sleep entered |
| 0x00266FF4 | `Exit Deep Sleep` | Deep sleep exited |
| 0x00267004 | `Begin Charging` | Charging started |
| 0x00267014 | `Stop Charging` | Charging stopped |

### 9.6 Memory/Resource Metrics

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x002663C4 | `Memory` | Memory section header |
| 0x0026648D | `Block Manager blocks: %d x %d KB = %d KB` | Block manager stats |
| 0x002664B9 | `Requested reserve heap size: %d KB` | Heap reservation |
| 0x002664E1 | `Free reserve heap size: %d KB` | Free heap |
| 0x002670D0 | `Free Reserve Heap Size` | Event log field |
| 0x002670E8 | `Number of Songs` | Song count field |
| 0x002670F8 | `Number of Photos` | Photo count field |
| 0x0026710C | `Number of Album Arts` | Artwork count field |
| 0x00267124 | `Heap Start Size` | Heap at start |
| 0x00267134 | `Heap End Size` | Heap at end |

---

## 10. Network/Sync Features

### 10.1 iTunes Sync

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00047D58 | `iPod_Control\iTunes\firsttime` | Known | First-time sync flag file |
| 0x00069108 | `iPod_Control\iTunes\iTunesDB` | Known | Main database path |
| 0x000B55CC | `iPod_Control\iTunes\Play Counts` | Known | Play count sync file |
| 0x00151964 | `MinITunesVersion` | Known | Minimum iTunes version check |
| 0x0006BD30 | `iPod_Control\Music\` | Known | Music file storage path |
| 0x00079418 | `rbsync` | Known | RB sync command/flag |
| 0x0006EEEC | `OTGPlaylistInfo` | Known | On-The-Go playlist info |
| 0x00113BD0 | `GeniusPlaylist_` | Known | Genius playlist prefix |
| 0x00113BE0 | `OTGPlaylistInfo_` | Known | OTG playlist prefix |

### 10.2 Nike+ iPod

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000AC758 | `TSilverTrainerCntlr` | Known | Nike+ trainer controller |
| 0x009BC23C | `NikePlus_Walk_To_Activate` | Known | Sensor activation prompt |
| 0x009BC256 | `NikePlus_Remote_Press_Button_To_Activate` | Known | Remote activation |
| 0x009BC7DD | `NikePlus_Remote_Searching` | Known | Searching for remote |
| 0x009BC7F7 | `NikePlus_Searching` | Known | Searching for sensor |
| 0x009BC832 | `NikePlus_Remote_Linking` | Known | Linking remote |
| 0x009BC84A | `NikePlus_Remote_Unlinking` | Known | Unlinking remote |
| 0x009B74B4 | `NikePlus_Remote_Now_Linked` | Known | Remote linked confirm |
| 0x009B74CF | `NikePlus_Now_Linked` | Known | Sensor linked confirm |
| 0x009BD7CA | `NikePlus_Alert_WeightEntryRequired_String` | Known | Weight required alert |
| 0x009BDA04 | `NikePlus_Alert_NotEnoughDiskSpace_String` | Known | Low space alert |
| 0x009BDD0A | `NikePlus_CustomDistance_String` | Known | Custom distance input |
| 0x009BDE88 | `NikePlus_Workout_Shuffle_String` | Known | Workout shuffle |
| 0x009BF6DB | `NikePlus_NowRunning_Distance_Km_String` | Known | Km distance display |
| 0x009C0719 | `NikePlus_NowRunning_Distance_Miles_String` | Known | Miles distance display |
| 0x009BDB18 | `NikePlus_ResetToDefault_Walking_Notice_String` | Known | Reset walking notice |
| 0x009BDB46 | `NikePlus_ResetToDefault_Running_Notice_String` | Known | Reset running notice |
| 0x009BDB74 | `NikePlus_History_ClearTotals_Notice_String` | Known | Clear totals notice |
| 0x009BDB9F | `NikePlus_History_DeleteAllWorkouts_Notice_String` | Known | Delete all workouts |
| 0x009BDBD0 | `NikePlus_EndPausedWorkout_Notice_String` | Known | End paused workout |
| 0x009BDBF8 | `NikePlus_History_DeleteActiveWorkout_Notice_String` | Known | Delete active workout |

### 10.3 Nike+ Workout Types

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x009BDE1D | `NikePlus_Male_String` | Gender: Male |
| 0x009BDE32 | `NikePlus_Female_String` | Gender: Female |
| 0x009BD794 | `NikePlus_Workout_Timed_String` | Timed workout |
| 0x009C1DF0 | `NikePlus_Distance_Workout_String` | Distance workout |
| 0x009C1E11 | `NikePlus_Time_Workout_String` | Time workout |
| 0x009C1E2E | `NikePlus_Calories_Workout_String` | Calorie workout |
| 0x009BED20 | `NikePlus_PowerSong_String` | Power song feature |
| 0x009BEDCF | `NikePlus_Spoken_Feedback_String` | Spoken feedback |
| 0x009BF6B5 | `NikePlus_CalibrationSuccessful_String` | Calibration success |
| 0x009C0F9A | `NikePlus_400Meters_String` | 400m calibration |
| 0x009C0FB4 | `NikePlus_Kilometers_String` | Kilometers unit |
| 0x009C076C | `NikePlus_Miles_String` | Miles unit |

### 10.4 Genius System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00152430 | `SupportsGenius` | Known | Genius capability flag |
| 0x00152440 | `GeniusConfigMinVersion` | Known | Min Genius config version |
| 0x00152458 | `GeniusMetadataMinVersion` | Known | Min metadata version |
| 0x00152474 | `GeniusSimilaritiesMinVersion` | Known | Min similarities version |
| 0x00152494 | `GeniusConfigMaxVersion` | Known | Max config version |
| 0x001524AC | `GeniusMetadataMaxVersion` | Known | Max metadata version |
| 0x001524C8 | `GeniusSimilaritiesMaxVersion` | Known | Max similarities version |
| 0x001524E8 | `SupportsGeniusMixes` | Known | Genius Mixes capability |
| 0x001DED04 | `RefreshingGenius` | Known | Genius refresh action |
| 0x001DED1C | `CreatingGeniusMix` | Known | Creating a Genius Mix |
| 0x001DF008 | `GeniusPlaylistReady` | Known | Genius playlist complete |
| 0x001DF01C | `GeniusMixPlaylistReady` | Known | Genius Mix complete |
| 0x0021E248 | `GotoGeniusError_NoGenius` | Known | No Genius data error |
| 0x0021E27C | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Track not in Genius DB |
| 0x007FFAB8 | `Connect to iTunes to activate Genius.` | Known | Genius activation msg |
| 0x007FFAE0 | `Genius is unavailable for the selected song.` | Known | Song not in Genius |
| 0x007FFB10 | `Creating Genius Playlist` | Known | Creating playlist msg |

---

## 11. Security System

### 11.1 FairPlay DRM

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x009DA728 | `$Apple FairPlay Certificate Authority0` | Known | Root CA certificate |
| 0x009DAAAD | `&Apple FairPlay Certification Authority0` | Known | Intermediate CA |
| 0x00A0F6CB | `Apple FairPlay1402` | Known | FairPlay version 1402 |
| 0x00150EB0 | `AppleDRMVersion` | Known | DRM version field |
| 0x00150F50 | `AppleDRM` | Known | DRM support flag |
| 0x001C26A0 | `drmidrms` (in atom chain) | Known | DRM MP4 atoms |
| 0x000E6EE4 | `EncryptedBlocks` | Known | Encrypted data blocks |

### 11.2 Game Code Signing (PKCS#7)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000985D8 | `iPod_Control/iTunes/iTunesDB.p7b` | Known | PKCS#7 binary signature |
| 0x002737F8 | `xmldsig#rsa-sha1` | Known | RSA-SHA1 algorithm ref |
| 0x00273874 | `xmldsig#sha1` | Known | SHA1 digest algorithm |
| 0x00273838 | `xmldsig#enveloped-signature` | Known | Enveloped signature type |
| 0x0075C9C1 | `controller.ShowSigningError1` | Known | Signing error trigger |
| 0x009C4955 | `Game_Signing_Error_Screen` | Known | Signing error display |
| 0x002D982C | `../../apps/openssl.cnf` | Hidden | OpenSSL config path |
| 0x0005FD74 | `ssl2-md5` | Known | SSL2-MD5 algorithm |
| 0x0005FD80 | `ssl3-md5` | Known | SSL3-MD5 algorithm |
| 0x0005FD94 | `ssl3-sha1` | Known | SSL3-SHA1 algorithm |
| 0x0005FDA0 | `RSA-SHA1` | Known | RSA-SHA1 algorithm |
| 0x0005FDAC | `RSA-SHA1-2` | Known | RSA-SHA1 variant |
| 0x00070120 | `NO X509_NAME` | Known | X509 name missing error |

### 11.3 X.509 / PKI Infrastructure

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x009D081D | `ASN1_SEQUENCE` | ASN.1 sequence type |
| 0x009D1D9A | `Netscape Certificate Sequence` | Certificate sequence |
| 0x009D1DB8 | `nsCertSequence` | NS cert sequence OID |
| 0x009D4B63 | `x509_req` | X509 request type |
| 0x009D4C23 | `OCSP request` | OCSP request type |
| 0x009D523E | `OCSP_REQINFO` | OCSP request info |
| 0x009D5295 | `OCSP_REQUEST` | OCSP request object |
| 0x009D52FE | `reqCert` | Request certificate |
| 0x009D5306 | `requestorName` | Requestor name |
| 0x009D1E89 | `X509v3 Key Usage` | Key usage extension |
| 0x009D1E9A | `X509v3 Extended Key Usage` | Extended key usage |

### 11.4 Lock / Security Screens

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00125B00 | `GotoScreen_LockDialog` | Known | Show lock PIN dialog |
| 0x00125B18 | `GotoScreen_SetCombinationFirstTime` | Known | First-time PIN setup |
| 0x00125BC0 | `HandleLock` | Known | Lock action handler |
| 0x0013FD5C | `SwitchLayout_AnimateLockLayout` | Known | Lock animation |
| 0x0013FDD0 | `GotoScreen_EnterPassKey` | Known | Enter passkey screen |
| 0x0013FDE8 | `GotoScreen_LockediPod` | Known | Locked iPod screen |
| 0x001400B8 | `SwitchLayout_EnterPasskey` | Known | Passkey entry layout |
| 0x001400D4 | `SwitchLayout_NewPasskey` | Known | New passkey layout |
| 0x00140224 | `PopScreen_LockDialog` | Known | Dismiss lock dialog |
| 0x0014023C | `PushScreen_LockDialog` | Known | Push lock dialog |
| 0x0014026C | `SwitchLayout_ConfirmPasskey` | Known | Confirm passkey |
| 0x001406C0 | `SwitchLayout_Locked` | Known | Locked state layout |
| 0x001407D8 | `SwitchLayout_Unlock` | Known | Unlock layout |
| 0x0027FC00 | `TC_LockDialog` | Known | Lock dialog controller |
| 0x0027FC18 | `TC_LockScreen` | Known | Lock screen controller |
| 0x0027FC30 | `TC_LockediPod` | Known | Locked state controller |
| 0x007FFC90 | `If you forget the combination, connect to iTunes to unlock your iPod.` | Known | Recovery message |

### 11.5 Serial/Hardware Verification

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x003EE91F | `SERIALVERIFIER` | Hidden | Serial number verifier task |
| 0x003EE930 | `RESISTORVERIFIER` | Hidden | Resistor ID verifier (dock auth) |
| 0x003F2FCC | `S_BTMREVERIFY` | Hidden | Bluetooth/bottom re-verify |
| 0x00078A9C | `iPod Serial Number` | Known | Serial number string |
| 0x001529BC | `RentalClockBias` | Known | Rental clock anti-tamper |

---

## 12. Error Handling

### 12.1 Assertion Patterns

The firmware contains extensive assertion checks. Pattern: `assertion failed on line %d of file %s`

| Offset | Context | Description |
|--------|---------|-------------|
| 0x0004F6E8 | Generic | Assertion at earliest offset |
| 0x0004F7D4 | Generic | Multiple assertion locations |
| 0x00086508 | Storage layer | Disk/storage assertion |
| 0x0009C360 | Driver layer | Hardware driver assertion |
| 0x000A3A04 | Font renderer | PFR font assertion |
| 0x000A983C | Audio/Codec | Audio codec assertion |
| 0x000B6388 | Database | Database layer assertion |
| 0x000C13E4 | Font/Type | CFF font assertion |
| 0x000CDD40 | Disk I/O | Disk assertion |
| 0x002D0268 | FreeType | FT2 rendering assertion |
| 0x002E5518 | Core | Core assertion |
| 0x00393408 | Late code | Late-stage assertion |
| 0x00395718 | Late code | Additional assertion |

### 12.2 Error Screens and Handlers

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x001DEFF0 | `GotoErrorLayout` | General error layout |
| 0x0021E248 | `GotoGeniusError_NoGenius` | Genius not available |
| 0x0021E27C | `GotoGeniusError_NoGeniusInfoForTrack` | Track not in Genius |
| 0x0021FAEC | `GotoExtraInfoLoadFailedLayout` | Info load failure |
| 0x0010CC98 | `SwitchToNotesImageError` | Notes image error |
| 0x0075D5C4 | `Game_Unknown_Error_Screen` | Game unknown error |
| 0x0075D627 | `Game_Version_Error_Screen` | Game version error |
| 0x009C49B7 | `Game_Memory_Error_Screen` | Game out of memory |
| 0x0075D683 | `Game_Running_Screen` | Game running state |
| 0x007FFF70 | `Insufficient Disk Space` | Low disk space |
| 0x007FFF88 | `Insufficient disk space to record.` | Recording space error |
| 0x007FFFAC | `There is not enough disk space to continue recording.` | Recording full |
| 0x00103318 | `No songs selected to play. (102)` | Empty playback error |
| 0x00119D88 | `Memory full. %d notes loaded, some notes not loaded. (30)` | Notes memory |
| 0x0011A3F0 | `Note too long, truncated. (7)` | Note truncation |

### 12.3 Hardware Error Messages

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x00075688 | `cI: Soft Reset sequence error` | Soft reset failure |
| 0x00075728 | `cI: CE-ATA interrupt enable failed` | ATA IRQ enable fail |
| 0x00077E9C | `mC: command failed` | MMC command failure |
| 0x000882DC | `mDS: write data CRC response error` | Write CRC error |
| 0x00088300 | `mDS: read data CRC error` | Read CRC error |
| 0x0008831C | `mDS: write data CRC error` | Write CRC error |
| 0x00088338 | `mDS: write data card CRC error` | Card CRC error |
| 0x00088358 | `mDS: read data end bit error(s)` | Read end bit error |
| 0x00099788 | `mRM: READ MULTIPLE REGISTER(0x%x, 0x%x) failed` | Register read fail |
| 0x000997EC | `mRM: data transfer error(s)` | Data transfer error |
| 0x000A13A0 | `mCS: response timeout error` | Response timeout |
| 0x000A13C0 | `mCS: response end bit error` | Response end bit |
| 0x000A13E0 | `mCS: response index error` | Response index |
| 0x000A13FC | `mCS: response CRC error` | Response CRC |
| 0x000A15D4 | `mWM: WRITE MULTIPLE REGISTER(0x%x, 0x%x) failed` | Register write fail |
| 0x000A163C | `mWM: data transfer error(s)` | Write transfer error |
| 0x000A9F90 | `ISR: Can't soft-reset device` | ISR reset failure |
| 0x000C5220 | `CGLE: Retry failed in getting ATA STATUS registers` | ATA status retry |
| 0x000C5254 | `CGLE: Retry failed in getting ATA ERR registers` | ATA error retry |
| 0x000D7BB0 | `cSTFR: FAST_IO Write failed` | Fast I/O write failure |
| 0x0036ADFC | `RMS: data transfer ATA status error` | RMS ATA error |
| 0x0036AE7C | `ISDIE: CE-ATA interrupt enable failed` | CE-ATA IRQ fail |

### 12.4 SQLite Error Messages

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x00376C60 | `not an error` | SQLITE_OK |
| 0x00376C70 | `SQL logic error or missing database` | SQLITE_ERROR |
| 0x00376CB0 | `callback requested query abort` | SQLITE_ABORT |
| 0x00376D20 | `database is locked` | SQLITE_BUSY |
| 0x00376D60 | `attempt to write a readonly database` | SQLITE_READONLY |
| 0x00376D94 | `disk I/O error` | SQLITE_IOERR |
| 0x00376DA4 | `database disk image is malformed` | SQLITE_CORRUPT |
| 0x00376DC8 | `database or disk is full` | SQLITE_FULL |
| 0x00376DE4 | `unable to open database file` | SQLITE_CANTOPEN |
| 0x00376E04 | `database schema has changed` | SQLITE_SCHEMA |
| 0x00376E74 | `constraint failed` | SQLITE_CONSTRAINT |
| 0x00376E9C | `library routine called out of sequence` | SQLITE_MISUSE |
| 0x00376EC4 | `large file support is disabled` | SQLITE_NOLFS |
| 0x00376EE4 | `unknown error` | Unknown |
| 0x0037FA44 | `near "%T": syntax error` | SQLITE_PARSE |

### 12.5 Exception Handling

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x00267D04 | `Exception: %s.` | C++ exception format |
| 0x00267D18 | `exception` | Exception base |
| 0x00267D24 | `unexpected exception` | Unexpected exception |
| 0x00267D3C | `bad_alloc: out of memory` | Memory allocation failure |
| 0x00267D58 | `unknown exception` | Unknown exception |
| 0x00986DEA | `Termination request` | std::terminate called |

---

## 13. Event Handler Reference (Handle* Functions)

### 13.1 Audio Playback Handlers

| Offset | Handler | Input Event |
|--------|---------|-------------|
| 0x001CF1D4 | `HandleAudioPlayPause` | `audio.playpause.up` |
| 0x001CF1EC | `HandleAudioNext` | `audio.next.up` |
| 0x001CF1FC | `HandleAudioNextPressAndHold` | `audio.next.pressandhold` |
| 0x001CF218 | `HandleAudioPrevious` | `audio.prev.up` |
| 0x001CF22C | `HandleAudioPreviousPressAndHold` | `audio.prev.pressandhold` |
| 0x001CF3BC | `HandleAudioNextAlbum` | `audio.nextalbum.down` |
| 0x001CF3D4 | `HandleAudioPrevAlbum` | `audio.prevalbum.down` |
| 0x001CF3EC | `HandleAudioVolumeDown` | `audio.volumedown.down` |
| 0x001CF404 | `HandleAudioVolumeUp` | `audio.volumeup.down` |
| 0x001CF418 | `HandleAudioVolumeDownUp` | `audio.volumedown.up` |
| 0x001CF430 | `HandleAudioVolumeUpUp` | `audio.volumeup.up` |
| 0x001CF448 | `HandleAudioStop` | `audio.stop.down` |
| 0x001CF458 | `HandleAudioPlay` | `audio.play.down` |
| 0x001CF468 | `HandleAudioPause` | `audio.pause.down` |
| 0x001CF47C | `HandleAudioMute` | `audio.mute.down` |
| 0x001CF48C | `HandleAudioNextChapter` | `audio.nextchapter.down` |
| 0x001CF4A4 | `HandleAudioPrevChapter` | `audio.prevchapter.down` |
| 0x001CF690 | `HandleAudioNextPlaylist` | `audio.nextplaylist.down` |
| 0x001CF6A8 | `HandleAudioPrevPlaylist` | `audio.prevplaylist.down` |
| 0x001CF6C0 | `HandleAudioShuffle` | `audio.shuffle.down` |
| 0x001CF6D4 | `HandleAudioRepeat` | `audio.repeat.down` |
| 0x001CF6E8 | `HandleAudioFFDown` | `audio.ff.down` |
| 0x001CF6FC | `HandleAudioFFUp` | `audio.ff.up` |
| 0x001CF70C | `HandleAudioRewDown` | `audio.rew.down` |
| 0x001CF720 | `HandleAudioRewUp` | `audio.rew.up` |

### 13.2 Video Playback Handlers

| Offset | Handler | Description |
|--------|---------|-------------|
| 0x001CF734 | `HandleVideoPlayPause` | Video play/pause |
| 0x001CF74C | `HandleVideoNext` | Next video |
| 0x001CF75C | `HandleVideoNextPressAndHold` | Fast forward |
| 0x001CF778 | `HandleVideoPrevious` | Previous video |
| 0x001CF78C | `HandleVideoPreviousPressAndHold` | Rewind |
| 0x001CF950 | `HandleVideoStop` | Stop video |
| 0x001CF960 | `HandleVideoPlay` | Play video |
| 0x001CF970 | `HandleVideoPause` | Pause video |
| 0x001CF984 | `HandleVideoFFDown` | FF press |
| 0x001CF998 | `HandleVideoFFUp` | FF release |
| 0x001CF9A8 | `HandleVideoRewDown` | Rew press |
| 0x001CF9BC | `HandleVideoRewUp` | Rew release |
| 0x001CF9D0 | `HandleVideoNextChapter` | Next chapter |
| 0x001CF9E8 | `HandleVideoPrevChapter` | Previous chapter |
| 0x001CFA00 | `HandleVideoNextFrame` | Frame advance |
| 0x001CFA18 | `HandleVideoPrevFrame` | Frame back |
| 0x001CFA30 | `HandleVideoCaptionAdvance` | Next subtitle |

### 13.3 Remote Control Handlers (Mikey/Dock)

| Offset | Handler | Input Event |
|--------|---------|-------------|
| 0x001CEA04 | `HandleRemoteVolumeUp` | `remote.volumeup.down` |
| 0x001CEBA8 | `HandleRemoteVolumeDown` | `remote.volumedown.down` |
| 0x001CEBC0 | `HandleRemoteVolumeUpUp` | `remote.volumeup.up` |
| 0x001CEBD8 | `HandleRemoteVolumeDownUp` | `remote.volumedown.up` |
| 0x001CFC44 | `HandleMikeyVolumeUp` | `mikey.volumeup.down` |
| 0x001CFC58 | `HandleMikeyVolumeUpUp` | `mikey.volumeup.up` |
| 0x001CFC70 | `HandleMikeyVolumeDown` | `mikey.volumedown.down` |
| 0x001CFC88 | `HandleMikeyVolumeDownUp` | `mikey.volumedown.up` |
| 0x001CEBF4 | `HandleRemoteStop` | Remote stop button |
| 0x001CE9D4 | `HandleRemoteNextAlbum` | Remote next album |
| 0x001CE9EC | `HandleRemotePrevAlbum` | Remote prev album |
| 0x001CEC74 | `HandleRemoteNextPlaylist` | Remote next playlist |
| 0x001CEC90 | `HandleRemotePrevPlaylist` | Remote prev playlist |
| 0x001CEE98 | `HandleRemoteShuffle` | Remote shuffle |
| 0x001CEEAC | `HandleRemoteRepeat` | Remote repeat |
| 0x001CEEE4 | `HandleRemoteBacklight` | Remote backlight |
| 0x001CF1B8 | `HandleRemoteBacklightOff` | Remote backlight off |
| 0x001CEF4C | `HandleRemoteMenuDown` | Remote menu press |
| 0x001CEF64 | `HandleRemoteMenuUp` | Remote menu release |

### 13.4 UI/Navigation Handlers

| Offset | Handler | Description |
|--------|---------|-------------|
| 0x001DCA54 | `HandleMainMenu` | Main menu selected |
| 0x001E7F1C | `HandleMusicMenu` | Music menu |
| 0x001E9118 | `HandleMenuSelection` | Menu item selected |
| 0x001E912C | `HandleViewAlbum` | View album action |
| 0x001E913C | `HandleViewArtist` | View artist action |
| 0x001E9168 | `HandleStartGenius` | Start Genius |
| 0x001E16D8 | `HandlePowerSongSelected` | Nike+ power song |
| 0x001E16F4 | `HandlePowerSongChosen` | Power song confirmed |
| 0x00217540 | `HandleMenuLongpress` | Long press handler |
| 0x002174D4 | `HandleMenuKey` | Menu button handler |
| 0x0016DA60 | `HandleAlbumSelected` | Album selected |
| 0x0016DA84 | `HandleBacksideSongSelected` | Song from album back |
| 0x0016D9F8 | `HandleFlipToAlbumBackside` | Flip to album back |
| 0x0016DA14 | `HandleFlipToAlbumFrontside` | Flip to album front |
| 0x001CBFD0 | `HandleBrowseSlideshow` | Browse slideshow |
| 0x0010C994 | `HandleNotesSelected` | Notes selected |
| 0x0010C9AC | `HandleNotesPop` | Pop notes screen |
| 0x0010C9BC | `HandleNotesPopToMainMenu` | Notes back to main |
| 0x00131A88 | `GotoNowPlaying` | Navigate to Now Playing |
| 0x00131A9C | `GotoAlbums` | Navigate to Albums |
| 0x00131AA8 | `GotoSongs` | Navigate to Songs |
| 0x001407EC | `GotoScreen_MainMenu` | Navigate to Main Menu |
| 0x0020F28C | `GotoScreen_SettingsMenu` | Navigate to Settings |
| 0x0021CB50 | `GotoScreen_Language` | Navigate to Language |

### 13.5 Volume / Audio Setting Handlers

| Offset | Handler | Description |
|--------|---------|-------------|
| 0x0011AFBC | `HandleVolumeWheel` | Wheel volume adjustment |
| 0x0011AFD0 | `HandleVolumeChange` | Volume change event |
| 0x0021D2BC | `HandleSelectVolume` | Volume selection |
| 0x00217414 | `HandleWheelVolume` | Wheel volume control |
| 0x0021D750 | `HandleAudiobookSlower` | Audiobook speed slower |
| 0x0021D768 | `HandleAudiobookFaster` | Audiobook speed faster |
| 0x0021D780 | `HandleAudiobookNormal` | Audiobook speed normal |
| 0x0010D328 | `EnterVolume` | Enter volume mode |

### 13.6 Radio Handlers

| Offset | Handler | Description |
|--------|---------|-------------|
| 0x0011AF0C | `TuneToNextPreset` | Next FM preset |
| 0x0011AF24 | `TuneToPreviousPreset` | Previous FM preset |
| 0x0011AF3C | `TuneToNextPresetWithTimer` | Next preset with delay |
| 0x0011AF58 | `TuneToPreviousPresetWithTimer` | Prev preset with delay |
| 0x0011AF9C | `TogglePreset` | Toggle preset mode |
| 0x0011AFAC | `TogglePlayPause` | Radio play/pause |
| 0x0011AFF4 | `HandleFrequencyChange` | Frequency change |
| 0x001F5E0C | `HandleFrequencyChosen` | Frequency confirmed |
| 0x001F5E4C | `HandleSoundChosen` | Alarm sound chosen |
| 0x001F5DFC | `ToggleAlarm` | Toggle alarm state |
| 0x0020B07C | `HandleRadioRegion` | Radio region change |
| 0x0011B378 | `SwitchLayout_RadioVolume` | Radio volume layout |

### 13.7 Video Settings Handlers

| Offset | Handler | Description |
|--------|---------|-------------|
| 0x002135C4 | `HandleCaptionSettingsChanged` | Caption change |
| 0x002135E4 | `HandleSubtitleSettingsChanged` | Subtitle change |
| 0x00213604 | `HandleAlternateAudioSettingsChanged` | Alt audio change |
| 0x00213628 | `HandleWideScreenSettingsChanged` | Widescreen toggle |
| 0x00213648 | `HandleVideoTVWideAspectRatioSettingsChanged` | TV aspect ratio |

### 13.8 Backlight / Display Handlers

| Offset | Handler | Description |
|--------|---------|-------------|
| 0x000FFAAC | `HandleCycleBacklightSetting` | Cycle backlight timer |
| 0x000FFACC | `HandleCycleLargeAlbumSetting` | Cycle album art size |
| 0x00212C40 | `SetBacklight_AlwaysOff` | Backlight always off |
| 0x00212C64 | `SetBacklight_2sec` | 2 second backlight |
| 0x00212C7C | `SetBacklight_5sec` | 5 second backlight |
| 0x00212C94 | `SetBacklight_10sec` | 10 second backlight |
| 0x00212CAC | `SetBacklight_15sec` | 15 second backlight |
| 0x00212CC4 | `SetBacklight_20sec` | 20 second backlight |
| 0x00212CDC | `SetBacklight_30sec` | 30 second backlight |
| 0x00212CF4 | `SetBacklight_AlwaysOn` | Backlight always on |
| 0x00212D10 | `SaveBacklight` | Save backlight setting |
| 0x00212D20 | `CancelBacklight` | Cancel backlight change |
| 0x00157954 | `EnterBrightness` | Enter brightness adjust |

---

## 14. Cover Flow System

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x0014370C | `Coverflow` | Cover Flow feature flag |
| 0x0027EDD0 | `TCoverFlowCntlr` | Cover Flow controller |
| 0x0016DF44 | `CoverflowPlayAlbum` | Play album from CFlow |
| 0x0016E118 | `FlipToAlbumBacksideEnded` | Back flip animation done |
| 0x0016E134 | `AlbumSelectedAnimationEnded` | Select animation done |
| 0x0016E150 | `SongSelectedAnimationEnded` | Song select animation |
| 0x0019E278 | `HandleHiliteAlbum` | Album highlight |
| 0x0019E290 | `HandleBrowseAlbum` | Album browse |
| 0x0022E35C | `HandleCoverFlowSelected` | Cover Flow entered |
| 0x0023015C | `CoverFlowSelected` | Cover Flow selection event |
| 0x0075BDF3 | `CoverFlow_Screen_Backside` | Album back screen |
| 0x0075C3C7 | `CoverFlow_Screen_Default` | Default CFlow screen |
| 0x0075BC90 | `CoverFlow_Screen_DefaultNoAlbumInfo` | No album info state |

---

## 15. Video Codec System

### 15.1 Video Container Atoms

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x001C3884 | `trakelsttkhdmdhdsttsstsdstcostszstscstssmp4vavc1` | Video track MP4V/AVC1 |
| 0x001524FC | `VideoCodecs` | Video codec capability |
| 0x00152510 | `AppleVideoDRM` | Video DRM flag |
| 0x0015254C | `MaximumBitRate` | Max bitrate supported |
| 0x00152580 | `MaximumAverageBitRate` | Max avg bitrate |
| 0x00152598 | `MaximumPeakBitRate` | Max peak bitrate |

### 15.2 Video UI

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x0021CEB0 | `GotoStatusBarVideoLayout` | Video status bar |
| 0x0021CECC | `GotoDefaultVideoLayout` | Default video layout |
| 0x0021CFA4 | `GotoVolumeVideoLayout` | Video volume overlay |
| 0x0021D058 | `GotoProgressVideoLayout` | Video progress bar |
| 0x0021DC74 | `GotoCaptionVideoLayout` | Caption/subtitle layout |
| 0x0021E09C | `GotoBrightnessVideoLayout` | Video brightness adjust |
| 0x0021ED00 | `GotoScrubVideoLayout` | Video scrub/seek |
| 0x002199B0 | `HandleRemotePlayPauseForVideo` | Remote video play/pause |

---

## 16. Filesystem Paths

| Offset | Path | Description |
|--------|------|-------------|
| 0x00047D40 | `iPod_Control\iTunes\` | iTunes control directory |
| 0x00047D58 | `iPod_Control\iTunes\firsttime` | First sync flag |
| 0x00069108 | `iPod_Control\iTunes\iTunesDB` | Main database |
| 0x000690CC | `rtSPhotos\Photo Database` | Photo database |
| 0x0006BD30 | `iPod_Control\Music\` | Music storage |
| 0x000800DC | `Contacts/` | Contacts directory |
| 0x000AB67C | `Calendars/` | Calendars directory |
| 0x000BDE88 | `Calendars` | Calendars root |
| 0x000C7560 | `Recordings/` | Voice memos storage |
| 0x000E9F34 | `Notes/` | Notes directory |
| 0x0010B314 | `iPod_Control/Device/alarms` | Alarm data |
| 0x0011BA8C | `iPod_Control/Device/radio` | Radio presets |
| 0x0028BFE8 | `Notes/Activity.log` | Activity log file |
| 0x00098648 | `gamedata_RW` | Game R/W data |
| 0x00098654 | `gamestats_WO` | Game stats write |
| 0x00098664 | `gamedata_ShareRW` | Shared game data |
| 0x00098678 | `games_RO` | Game binaries |

---

## 17. RTTI Class Hierarchy (C++ Mangled Names)

These represent the C++ class inheritance tree as found in RTTI data:

| Offset | Mangled Name | Demangled |
|--------|--------------|-----------|
| 0x0098845F | `11TCEQSetting` | TCEQSetting (11 chars) |
| 0x00988624 | `13TCSettings_EQ` | TCSettings_EQ |
| 0x0098880C | `15iFSVolumeClient` | iFSVolumeClient |
| 0x00988D20 | `21SoundEffectDescriptor` | SoundEffectDescriptor |
| 0x00988E9A | `22TCSettings_VolumeLimit` | TCSettings_VolumeLimit |
| 0x00988F9B | `24TC_VolumeLimitLockScreen` | TC_VolumeLimitLockScreen |
| 0x009890EB | `26TCAlarmPropertiesSoundMenu` | TCAlarmPropertiesSoundMenu |
| 0x009892C1 | `27TSilverCntlrTransitionAddonI11TCEQSettingE` | TSilverCntlrTransitionAddon<TCEQSetting> |
| 0x009894E4 | `27TSilverCntlrTransitionAddonI13TCSettings_EQE` | TSilverCntlrTransitionAddon<TCSettings_EQ> |
| 0x00989DBA | `27TSilverCntlrTransitionAddonI22TCSettings_VolumeLimitE` | TSilverCntlrTransitionAddon<TCSettings_VolumeLimit> |
| 0x0098A03A | `27TSilverCntlrTransitionAddonI26TCAlarmPropertiesSoundMenuE` | TSilverCntlrTransitionAddon<TCAlarmPropertiesSoundMenu> |
| 0x0098A25A | `27TSilverCntlrTransitionAddonI28TCSettings_AudiobookSettingsE` | TSilverCntlrTransitionAddon<TCSettings_AudiobookSettings> |
| 0x0098A642 | `27TSilverCntlrTransitionAddonI30TCAlarmPropertiesFrequencyMenuE` | TSilverCntlrTransitionAddon<TCAlarmPropertiesFrequencyMenu> |
| 0x0098A9CF | `27TSilverCntlrTransitionAddonI32TSilverMediaListCntlr_AudiobooksE` | TSilverCntlrTransitionAddon<TSilverMediaListCntlr_Audiobooks> |
| 0x0098AD10 | `27TSilverCntlrTransitionAddonI39TSilverMediaListCntlr_AudiobookChaptersE` | TSilverCntlrTransitionAddon<TSilverMediaListCntlr_AudiobookChapters> |
| 0x0098ADEB | `28TCSettings_AudiobookSettings` | TCSettings_AudiobookSettings |
| 0x0098AFE3 | `30TCAlarmPropertiesFrequencyMenu` | TCAlarmPropertiesFrequencyMenu |
| 0x0098B1DF | `32TSilverMediaListCntlr_Audiobooks` | TSilverMediaListCntlr_Audiobooks |
| 0x0098B41F | `39TSilverMediaListCntlr_AudiobookChapters` | TSilverMediaListCntlr_AudiobookChapters |

---

## 18. Supported Languages

The firmware contains full UI string tables in the following languages (detected from translated strings in 0x0083xxxx-0x008Fxxxx range):

1. English (primary, at 0x007FFxxx)
2. Czech (0x0083xxxx - "Audio knihy")
3. Danish (0x0084xxxx)
4. German (0x0085xxxx - "Alternatives Audio")
5. Spanish (0x0086xxxx - "Audiolibros")
6. Finnish (0x0087xxxx region)
7. French (0x0087xxxx - "Livres audio")
8. Hungarian (0x0088xxxx)
9. Italian (0x0088xxxx - "Audiolibri")
10. Japanese (0x0089xxxx region)
11. Korean (0x008Axxxx region)
12. Dutch (0x008Axxxx - "Audioboeken")
13. Norwegian (0x008Bxxxx - "Spiller nå")
14. Polish (0x008Cxxxx - "Książki audio")
15. Portuguese (0x008Cxxxx - "Audiolivros")
16. Russian (0x008Dxxxx)
17. Swedish (0x008Exxxx - "Bildspelsmusik")
18. Turkish (0x008Fxxxx)
19. Chinese Simplified (0x0090xxxx region)
20. Chinese Traditional (0x0090xxxx region)

---

## 19. Binary Structure Summary for Ghidra

### 19.1 Load Configuration

| Parameter | Value |
|-----------|-------|
| **Architecture** | ARM (32-bit, little-endian) |
| **Base Address** | 0x00000000 (raw binary offset 0x800) |
| **Entry Point** | Vector table at offset 0x800 |
| **Code Region** | 0x800 - ~0xA1FFFF |
| **String Tables** | Concentrated at 0x007Fxxxx - 0x009Fxxxx |
| **RTTI Data** | 0x0098xxxx - 0x009Exxxx |
| **UI Layout Data** | 0x0073xxxx - 0x007Exxxx |
| **Localization** | 0x0080xxxx - 0x0090xxxx |

### 19.2 Key Regions for Analysis

| Offset Range | Content | Priority |
|--------------|---------|----------|
| 0x000AC000 - 0x000AD000 | Controller class definitions | HIGH |
| 0x00125000 - 0x00145000 | Lock/security code | HIGH |
| 0x00150000 - 0x00160000 | Device capabilities | HIGH |
| 0x00189000 - 0x0018A000 | Demo mode | MEDIUM |
| 0x001CF000 - 0x001D0000 | Event handlers (audio/video) | HIGH |
| 0x001E0000 - 0x001F0000 | Settings/UI handlers | MEDIUM |
| 0x00220000 - 0x00235000 | Settings implementation | MEDIUM |
| 0x00265000 - 0x00268000 | Logging/analytics system | HIGH |
| 0x002AD000 - 0x002AE000 | RTOS task definitions | HIGH |
| 0x002BF000 - 0x002C1000 | RTXCbug debugger | HIGH |
| 0x0036A000 - 0x003A0000 | SQLite + disk drivers | MEDIUM |
| 0x003ED000 - 0x00410000 | Late-init controllers + RTOS objects | HIGH |
| 0x009B4000 - 0x009D6000 | UI resource names (screens/views/images) | MEDIUM |
| 0x009DA000 - 0x00A20000 | Crypto/DRM certificates | HIGH |

### 19.3 Function Identification Tips

- **TSilver* classes**: MVC controllers, look for vtable references
- **TC* classes**: Concrete screen controllers with transition addons
- **Handle* functions**: Event callbacks, registered in dispatch tables
- **Goto* functions**: Navigation/screen transition triggers
- **Toggle* functions**: Binary state flip + UI refresh
- **Show* functions**: Display current setting value
- **Switch* functions**: Layout/view transitions within a screen

---

## 20. Miscellaneous Features

### 20.1 Rental System (iTunes Movie Rentals)

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x000AC610 | `TSilverMediaListCntlr_Rentals` | Rental list |
| 0x000AC68C | `TCRentalNotification` | Expiry notification |
| 0x000AC6AC | `TCRentalInfo` | Rental info screen |
| 0x000AC6C4 | `TCRentalConfirmDelete` | Confirm rental delete |
| 0x000AC6E4 | `TCRentalDispatcher` | Rental action dispatch |
| 0x001529BC | `RentalClockBias` | Clock tamper protection |
| 0x00217554 | `HandleRentalWarningChoice` | Rental warning handler |
| 0x0021791C | `GotoRentalWarningLayout` | Warning layout |
| 0x007FFFF4 | `This rental has expired.` | Expired message |
| 0x008000A4 | `View Rentals` | View rentals button |

### 20.2 On-The-Go Playlists

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x0006EEEC | `OTGPlaylistInfo` | OTG playlist metadata |
| 0x007FFA74 | `Add to On-The-Go` | Add to OTG menu item |
| 0x007FFA88 | `Remove from On-The-Go` | Remove from OTG |
| 0x001E86C4 | `GotoPlayDeleteMenu` | Play/delete context menu |

### 20.3 Contextual Menu

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x000AC700 | `TContextualMenuCntlr` | Context menu controller |
| 0x007FFA40 | `Browse Artist` | Context: browse artist |
| 0x007FFA50 | `Browse Album` | Context: browse album |
| 0x007FFA60 | `Browse Compilation` | Context: browse compilation |
| 0x007FFA74 | `Add to On-The-Go` | Context: add to OTG |
| 0x007FFA88 | `Remove from On-The-Go` | Context: remove from OTG |
| 0x007FFAA0 | `Cancel` | Context: cancel |
| 0x007FFA30 | `Start Genius` | Context: start Genius |
| 0x0021D5F4 | `HandlePushContextualMenu` | Push context menu |
| 0x002199D0 | `HandleShowContextualMenu` | Show context menu |

### 20.4 Showcase / Demo Content

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x001B84A0 | `Showcase [audio]` | Demo audio content |
| 0x001B84B4 | `Showcase [video]` | Demo video content |
| 0x001B84C8 | `Showcase [podcast]` | Demo podcast content |
| 0x001B84DC | `Showcase [photo]` | Demo photo content |

### 20.5 Reset Confirmation Dialogs

| Offset | String | Description |
|--------|--------|-------------|
| 0x007FFCD8 | `Do you want to reset your iPod?` | Reset confirm |
| 0x007FFCF8 | `This will reset all your settings. Your synced content will not be modified.` | Reset explanation |
| 0x007FFD48 | `Do you want to reset the main menu?` | Menu reset |
| 0x007FFD6C | `This will restore your menu to its original state.` | Menu reset detail |
| 0x007FFDA0 | `Do you want to reset the music menu?` | Music menu reset |
| 0x007FFDC8 | `This will restore your music menu to its original state.` | Music reset detail |
| 0x001DCB38 | `GotoScreen_ConfirmCancelResetMenu` | Reset confirmation |

---

*Document generated from firmware string analysis of RetailOS 2.0.4 binary.*
*Total unique features cataloged: 500+ symbols across 12 functional categories.*
*Recommended Ghidra base address: 0x00000000 with binary offset 0x800 for ARM code.*
