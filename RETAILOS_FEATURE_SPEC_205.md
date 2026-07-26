# iPod Classic 7th Generation - RetailOS 2.0.5 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 2.0.5 |
| **IPSW** | iPod_38.2.0.5.ipsw |
| **Device** | iPod Classic (Rev C, Late 2012, 160GB) |
| **UpdaterFamilyID** | 38 |
| **Binary Size** | 10,634,528 bytes (10.14 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,632,480 bytes |
| **Total Strings** | 72,926 |
| **Function Prologues** | 23,164 (ARM: 17,762, Thumb: 5,402) |
| **DRAM References** | 2,578 |
| **Peripheral Refs** | 9,481 |
| **Build** | N25CFirmwareWin-247 |

---

## 1. Main Menu Items

| Offset | String | Classification |
|--------|--------|----------------|
| 0x0005EEBC | `SearchCriteria::DataCache` | Known - UI |
| 0x00069230 | ` rtSPhotos\Photo Database` | Known - UI |
| 0x000AC1F0 | `TCExtrasMenu` | Known - UI |
| 0x000AC2D4 | `TSilverMediaListCntlr_Albums` | Known - UI |
| 0x000AC2FC | `TSilverMediaListCntlr_Artists` | Known - UI |
| 0x000AC324 | `TSilverMediaListCntlr_Genres` | Known - UI |
| 0x000AC7A4 | `TSilverMediaListCntlr_Genius` | Known - UI |
| 0x000AC88C | `TGeniusLoadingCntlr` | Known - UI |
| 0x000B608C | `GeniusPlaylist` | Known - UI |
| 0x000CBC08 | `Playlists::NameCache` | Known - UI |
| 0x000CBC20 | `Playlists::DescCache` | Known - UI |
| 0x000FFC44 | `HandleCycleLargeAlbumSetting` | Known - UI |
| 0x00104220 | `TPhotosBrowseCntlr` | Known - UI |
| 0x0010423C | `TPhotosBrowseTransitionCntlr` | Known - UI |
| 0x00104264 | `TPhotosMenuCntlr` | Known - UI |
| 0x00104280 | `TPhotosSettingsCntlr` | Known - UI |
| 0x001042CC | `TPhotosSettingsCntlr_Duration` | Known - UI |
| 0x0010460C | `TSearchCntlr` | Known - UI |
| 0x0010C418 | `HandleEQSettingSelected` | Known - UI |
| 0x00113E38 | `GeniusPlaylist_` | Known - UI |
| 0x001269E0 | `ExitToExtras` | Known - UI |
| 0x00131378 | `EnableSearchListHilite` | Known - UI |
| 0x00131D54 | `GotoAlbums` | Known - UI |
| 0x00133710 | `SearchHelperThread` | Known - UI |
| 0x00151F44 | `PodcastsSupported` | Known - UI |
| 0x00152A50 | `SupportsGenius` | Known - UI |
| 0x00152A60 | `GeniusConfigMinVersion` | Known - UI |
| 0x00152A78 | `GeniusMetadataMinVersion` | Known - UI |
| 0x00152A94 | `GeniusSimilaritiesMinVersion` | Known - UI |
| 0x00152AB4 | `GeniusConfigMaxVersion` | Known - UI |
| 0x00152ACC | `GeniusMetadataMaxVersion` | Known - UI |
| 0x00152AE8 | `GeniusSimilaritiesMaxVersion` | Known - UI |
| 0x00152B08 | `SupportsGeniusMixes` | Known - UI |
| 0x0016E2E4 | `HandleAlbumSelected` | Known - UI |
| 0x0016E9B8 | `AlbumSelectedAnimationEnded` | Known - UI |
| 0x00172E70 | `TSilverSettingsMenuListCntlr` | Known - UI |
| 0x00172E98 | `TCSettings_MainMenu` | Known - UI |
| 0x00172EB4 | `TCSettings_MusicMenu` | Known - UI |
| 0x00172ED4 | `TCSettings_VolumeLimit` | Known - UI |
| 0x00172EF4 | `TCSettings_Brightness` | Known - UI |

---

## 2. Controllers (TSilver / TC Classes)

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x000AC1D8 | `TSilverCntlr` | Controller |
| 0x000AC1F0 | `TCExtrasMenu` | Controller |
| 0x000AC208 | `TCGameScreen` | Controller |
| 0x000AC234 | `TSilverMainMediaListCntlr_Main` | Controller |
| 0x000AC25C | `TSilverMainMediaListCntlr_Music` | Controller |
| 0x000AC284 | `TSilverMainMediaListCntlr_Videos` | Controller |
| 0x000AC2B0 | `TSilverMediaListCntlr_Songs` | Controller |
| 0x000AC2D4 | `TSilverMediaListCntlr_Albums` | Controller |
| 0x000AC2FC | `TSilverMediaListCntlr_Artists` | Controller |
| 0x000AC324 | `TSilverMediaListCntlr_Genres` | Controller |
| 0x000AC34C | `TSilverMediaListCntlr_Composers` | Controller |
| 0x000AC374 | `TSilverMediaListCntlr_Podcasts` | Controller |
| 0x000AC39C | `TSilverMediaListCntlr_PodcastEpisodes` | Controller |
| 0x000AC3CC | `TSilverMediaListCntlr_iTunesU` | Controller |
| 0x000AC3F4 | `TSilverMediaListCntlr_iTunesUEpisodes` | Controller |
| 0x000AC424 | `TSilverMediaListCntlr_Audiobooks` | Controller |
| 0x000AC450 | `TSilverMediaListCntlr_AudiobookChapters` | Controller |
| 0x000AC480 | `TSilverMediaListCntlr_TVShows` | Controller |
| 0x000AC4A8 | `TSilverMediaListCntlr_TVSeasons` | Controller |
| 0x000AC4D0 | `TSilverMediaListCntlr_TVEpisodes` | Controller |
| 0x000AC4FC | `TSilverMediaListCntlr_Movies` | Controller |
| 0x000AC524 | `TSilverMediaListCntlr_Playlists` | Controller |
| 0x000AC54C | `TSilverMediaListCntlr_NestedPlaylists` | Controller |
| 0x000AC57C | `TSilverMediaListCntlr_VideoPlaylists` | Controller |
| 0x000AC718 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Controller |
| 0x000AC74C | `TSilverMediaListCntlr_PlaylistChooser` | Controller |
| 0x000AC77C | `TSilverMediaListCntlr_Rentals` | Controller |
| 0x000AC7A4 | `TSilverMediaListCntlr_Genius` | Controller |
| 0x000AC7CC | `TSilverMediaListCntlr_GeniusMixes` | Controller |
| 0x000AC86C | `TContextualMenuCntlr` | Controller |
| 0x000AC8A8 | `TSilverGlobalCntlr` | Controller |
| 0x000AC8C4 | `TSilverTrainerCntlr` | Controller |
| 0x00104220 | `TPhotosBrowseCntlr` | Controller |
| 0x0010423C | `TPhotosBrowseTransitionCntlr` | Controller |
| 0x00104264 | `TPhotosMenuCntlr` | Controller |
| 0x00104280 | `TPhotosSettingsCntlr` | Controller |
| 0x001042A0 | `TPhotosSettingsCntlr_Transitions` | Controller |
| 0x001042CC | `TPhotosSettingsCntlr_Duration` | Controller |
| 0x001042F4 | `TPhotosSettingsSlideshowPlaylistCntlr` | Controller |
| 0x0010460C | `TSearchCntlr` | Controller |
| 0x001280B8 | `TSilverCalendarCntlr_CalendarMenu` | Controller |
| 0x001280E4 | `TSilverCalendarCntlr_MonthViewer` | Controller |
| 0x00128110 | `TSilverCalendarCntlr_DayViewer` | Controller |
| 0x00128138 | `TSilverCalendarCntlr_EventViewer` | Controller |
| 0x00128164 | `TSilverCalendarCntlr_ToDoList` | Controller |
| 0x0012818C | `TSilverCalendarCntlr_ToDoDispatcher` | Controller |
| 0x001281B8 | `TSilverCalendarCntlr_Alarm` | Controller |
| 0x0012FB38 | `TCRemoteUI` | Controller |
| 0x00136350 | `TCSpeakers` | Controller |
| 0x00144A34 | `TChargingModeCntlr` | Controller |
| 0x00144A50 | `TChargingModeLowPowerCntlr` | Controller |
| 0x0015F85C | `TCSportTimer` | Controller |
| 0x0015F874 | `TCSportTimerMenu` | Controller |
| 0x0015F890 | `TCSportTimerSessionScreen` | Controller |
| 0x0015F8B4 | `TCSportTimerChosenDispatcher` | Controller |
| 0x00160C64 | `TCVoiceMemos` | Controller |
| 0x00160C7C | `TCVoiceMemosMenu` | Controller |
| 0x00160C98 | `TCVoiceMemosMainMenu` | Controller |
| 0x00160CB8 | `TCVoiceMemosPlayback` | Controller |
| 0x00160CD8 | `TCVoiceMemosContextMenu` | Controller |
| 0x00160CF8 | `TCVoiceMemosAlert` | Controller |
| 0x00172E70 | `TSilverSettingsMenuListCntlr` | Controller |
| 0x00172E98 | `TCSettings_MainMenu` | Controller |
| 0x00172EB4 | `TCSettings_MusicMenu` | Controller |
| 0x00172ED4 | `TCSettings_VolumeLimit` | Controller |
| 0x00172EF4 | `TCSettings_Brightness` | Controller |
| 0x00172F14 | `TCSettings_BacklightTimer` | Controller |
| 0x00172F38 | `TCSettings_EQ` | Controller |
| 0x00172F50 | `TCSettings_AudiobookSettings` | Controller |
| 0x00172F78 | `TCSettings_RadioRegions` | Controller |
| 0x00172F98 | `TCSettings_ResetAllSettings` | Controller |
| 0x00172FBC | `TCSettings_EULimitConfirmation` | Controller |
| 0x00172FE4 | `TSilverSettingsVideoCntlr` | Controller |
| 0x00173040 | `TCSettings_AdjustScrollingCntlr` | Controller |
| 0x00173068 | `TCFirstBoot` | Controller |
| 0x00189C74 | `TCDemoMode` | Controller |
| 0x001B2E80 | `TCAddressViewerMainMenu` | Controller |
| 0x001B2EA0 | `TCAddressViewerDetails` | Controller |
| 0x001B2EC0 | `TCAddressViewerPartialLoad` | Controller |
| 0x001B2EE4 | `TCAddressViewerMainDispatcher` | Controller |

---

## 3. Settings (Toggle/Show/TC)

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x0013FDE8 | `ShowSetting_EQ` | Setting |
| 0x00172E98 | `TCSettings_MainMenu` | Setting |
| 0x00172EB4 | `TCSettings_MusicMenu` | Setting |
| 0x00172ED4 | `TCSettings_VolumeLimit` | Setting |
| 0x00172EF4 | `TCSettings_Brightness` | Setting |
| 0x00172F14 | `TCSettings_BacklightTimer` | Setting |
| 0x00172F38 | `TCSettings_EQ` | Setting |
| 0x00172F50 | `TCSettings_AudiobookSettings` | Setting |
| 0x00172F78 | `TCSettings_RadioRegions` | Setting |
| 0x00172F98 | `TCSettings_ResetAllSettings` | Setting |
| 0x00172FBC | `TCSettings_EULimitConfirmation` | Setting |
| 0x00173040 | `TCSettings_AdjustScrollingCntlr` | Setting |
| 0x001EABD4 | `ToggleSetting_Repeat` | Setting |
| 0x001EABF0 | `ToggleSetting_Shuffle` | Setting |
| 0x001EAC08 | `ToggleSetting_TVOut` | Setting |
| 0x001EAC1C | `ToggleSetting_TVSignal` | Setting |
| 0x002140EC | `ShowSetting_Backlight` | Setting |
| 0x00229448 | `ToggleSetting_SortBy` | Setting |
| 0x00229460 | `ToggleSetting_ClassicUI` | Setting |
| 0x00229478 | `ToggleSetting_SoundCheck` | Setting |
| 0x00229494 | `ToggleSetting_Clicker` | Setting |
| 0x002294AC | `ToggleSetting_DaylightSavings` | Setting |
| 0x002294CC | `ToggleSetting_24HourClock` | Setting |
| 0x002294E8 | `ToggleSetting_TimeInTitle` | Setting |
| 0x00229504 | `ShowSetting_Shuffle` | Setting |
| 0x002296B0 | `ShowSetting_Repeat` | Setting |
| 0x002296C4 | `ShowSetting_About` | Setting |
| 0x002296D8 | `ShowSetting_MainMenu` | Setting |
| 0x002296F0 | `ShowSetting_MusicMenu` | Setting |
| 0x00229708 | `ShowSetting_VolumeLimit` | Setting |
| 0x00229720 | `ShowSetting_BacklightTimer` | Setting |
| 0x0022973C | `ShowSetting_Brightness` | Setting |
| 0x00229754 | `ShowSetting_Audiobooks` | Setting |
| 0x0022976C | `ShowSetting_RadioRegions` | Setting |
| 0x00229798 | `ShowSetting_SoundCheck` | Setting |
| 0x00229918 | `ShowSetting_Clicker` | Setting |
| 0x0022992C | `ShowSetting_DateAndTime` | Setting |
| 0x00229944 | `ShowSetting_SortBy` | Setting |
| 0x00229958 | `ShowSetting_ClassicUI` | Setting |
| 0x00229970 | `ShowSetting_Language` | Setting |
| 0x00229988 | `ShowSetting_Legal` | Setting |
| 0x0022999C | `ShowSetting_ResetAll` | Setting |
| 0x002299FC | `ToggleSetting_RecommendedVolumeLimit` | Setting |
| 0x0075B800 | `TSilverSettingsMenuListCntlrTSilverSettingsMenuListCntlrTSilverSettingsMenuListCntlrTCSettings_MainMenu` | Setting |
| 0x0075B868 | `TCSettings_MusicMenuTCSettings_VolumeLimit` | Setting |
| 0x0075B8C8 | `TCSettings_AudiobookSettingsTCSettings_RadioRegions` | Setting |
| 0x0075B8FC | `TSilverSettingsMenuListCntlrTCSettings_EULimitConfirmation` | Setting |
| 0x0075B948 | `TSilverSettingsMenuListCntlrTSilverCntlrTCSettings_AdjustScrollingCntlr` | Setting |
| 0x0075B9C0 | `TCSettings_MusicMenuTCSportTimerTCSportTimerMenuTCSportTimerSessionScreen` | Setting |
| 0x009902EC | `13TCSettings_EQ` | Setting |

---

## 4. NEW IN 2.0.5: EU Volume Limit System

This feature was added in 2.0.5 to comply with EU regulation 2006/95/EC on maximum headphone volume.

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x00172FBC | `TCSettings_EULimitConfirmation` | EU Volume Limit |
| 0x002299FC | `ToggleSetting_RecommendedVolumeLimit` | EU Volume Limit |
| 0x00229E54 | `GotoScreen_VolumeLimitEU` | EU Volume Limit |
| 0x00229F18 | `GotoScreen_EUVolumeLimitConfirmation` | EU Volume Limit |
| 0x0075B8FC | `TSilverSettingsMenuListCntlrTCSettings_EULimitConfirmation` | EU Volume Limit |
| 0x00761961 | `SettingsMenus_VolumeLimitEU_Screen"` | EU Volume Limit |
| 0x00761A3E | `SettingsMenus_VolumeLimitEU_Screen!` | EU Volume Limit |
| 0x007E26F2 | `controller.GotoScreen_VolumeLimitEU1` | EU Volume Limit |
| 0x007E273E | `SettingsMenus_VolumeLimitEU_Screen*` | EU Volume Limit |
| 0x007E2764 | `SettingsMenus_VolumeLimitEU_Screen_Default,` | EU Volume Limit |
| 0x007E764D | `SettingsMenus_VolumeLimitEU_Screen,` | EU Volume Limit |
| 0x007E7840 | `controller.GotoScreen_EUVolumeLimitConfirmation1` | EU Volume Limit |
| 0x007E7898 | `SettingsMenus_EUVolume_Confirmation_Screen2` | EU Volume Limit |
| 0x007E78C6 | `SettingsMenus_EUVolume_Confirmation_Screen_Default!` | EU Volume Limit |
| 0x008071FC | `This will limit the maximum headphone volume to the European Union recommended level.` | EU Volume Limit |
| 0x008AEEDF | `(European Union) ` | EU Volume Limit |
| 0x0099234A | `27TSilverCntlrTransitionAddonI30TCSettings_EULimitConfirmationE` | EU Volume Limit |
| 0x00992D0C | `30TCSettings_EULimitConfirmation` | EU Volume Limit |
| 0x009BE80D | `SettingsMenu_ListItem_VolumeLimitEU` | EU Volume Limit |
| 0x009C2B33 | `SettingsMenu_ListItem_VolumeLimitEU_Toggle` | EU Volume Limit |
| 0x009C9CA9 | `SettingsMenu_EUVolumeLimit_String` | EU Volume Limit |
| 0x009CC744 | `SettingsMenus_VolumeLimitEU_Screen` | EU Volume Limit |
| 0x009CCC26 | `SettingsMenus_EUVolume_Confirmation_Screen` | EU Volume Limit |
| 0x009D2492 | `SettingsMenus_VolumeLimitEU_Screen_Default` | EU Volume Limit |
| 0x009D2939 | `SettingsMenus_EUVolume_Confirmation_Screen_Default` | EU Volume Limit |
| 0x009D44FE | `SettingsMenus_DialogNotice_EULimitConfirmation_Layout` | EU Volume Limit |
| 0x008071FC | `This will limit the maximum headphone volume to the European Union rec...` | EU Localized |
| 0x00872001 | `ximo de los auriculares al nivel recomendado por la Uni` | EU Localized |
| 0x00872294 | `Auriculares` | EU Localized |
| 0x008722EC | `Emitir sonido por los auriculares` | EU Localized |
| 0x00872310 | `Emitir sonido por los auriculares y el altavoz` | EU Localized |
| 0x00884EC7 | `Union europ` | EU Localized |
| 0x00898E6C | ` limitato al livello consigliato dall'Unione europea.` | EU Localized |
| 0x00899560 | `Limita il volume massimo delle cuffie al livello consigliato dall'Unio...` | EU Localized |
| 0x008B8CFC | `Hiermee beperkt u het maximale koptelefoonvolume tot het door de Europ...` | EU Localized |
| 0x008B942C | `Beperk het maximale koptelefoonvolume tot het door de Europese Unie aa...` | EU Localized |
| 0x008C2390 | `Dette begrenser maksimalvolumet for hodetelefonene til niv` | EU Localized |
| 0x008C2A28 | `Begrens maksimalvolumet for hodetelefonene til niv` | EU Localized |

---

## 5. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001449A4 | `MP3ExampleTask` | Hidden | |
| 0x00189C74 | `TCDemoMode` | Hidden | |
| 0x001E13BC | `TSilverCntlrTestAppCntlr` | Hidden | |
| 0x001E13E0 | `TSilverCntlrTestCntlr` | Hidden | |
| 0x002669E4 | `Channel Reserved` | Hidden | |
| 0x002669F8 | `Channel AppBoot` | Hidden | |
| 0x00266A08 | `Channel BufferedSongReading` | Hidden | |
| 0x00266A24 | `Channel PrefsWriting` | Hidden | |
| 0x00266A3C | `Channel GeneralUserExperience` | Hidden | |
| 0x00266A5C | `Channel PlayFromDisk` | Hidden | |
| 0x00266A74 | `Channel CacheSpinupDrive` | Hidden | |
| 0x00266A90 | `Channel TestLogging` | Hidden | |
| 0x00266AA4 | `Channel AppFileLoading` | Hidden | |
| 0x00266ABC | `Channel VCardReading` | Hidden | |
| 0x00266AD4 | `Channel LongSongScanning` | Hidden | |
| 0x00266B48 | `Channel VoiceRecording` | Hidden | |
| 0x00266B60 | `Channel PhotoImporting` | Hidden | |
| 0x00266B78 | `Channel Notes` | Hidden | |
| 0x00266B88 | `Channel PhotoFileManagement` | Hidden | |
| 0x00266BA4 | `Channel DiskMode` | Hidden | |
| 0x00266BB8 | `Channel Firewire` | Hidden | |
| 0x00266BCC | `Channel USB` | Hidden | |
| 0x00266BD8 | `Channel UnitTests` | Hidden | |
| 0x00266BEC | `Channel FreeSpaceCache` | Hidden | |
| 0x00266C04 | `Channel OnTheGoFileMgmt` | Hidden | |
| 0x002C13DD | `** RTXCbug - ` | Hidden | |
| 0x002C1420 | `  X - Exit RTXCbug` | Hidden | |
| 0x002C1435 | `RTXCbug> ` | Hidden | |
| 0x002C1E11 | `RTXCbug - RTXC Objects> ` | Hidden | |
| 0x002DB6F0 | `X - Exit RTXCbug` | Hidden | |
| 0x003960E5 | `$RTXCbug> ` | Hidden | |
| 0x003961AD | `Re-entering RTXCbug mode` | Hidden | |
| 0x003F4E0D | `S_RTXCBUG` | Hidden | |
| 0x0073B49E | `Debug_MainMenu_Screen` | Hidden | |
| 0x0073B4B7 | `Debug_MainMenu_Screen_Default"` | Hidden | |
| 0x0075B0C8 | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDemoMode` | Hidden | |
| 0x0075BA0C | `TSilverCntlrTUnitTestSuiteCntlr` | Hidden | |
| 0x0075BA2C | `TUnitTestSuiteTestsCntlrTSilverCntlrTCVoiceMemosMainMenuTCVoiceMemosMenuTCVoiceMemosContextMenu` | Hidden | |
| 0x007E81B6 | `Debug_UnitTest_Screen` | Hidden | |
| 0x007E81CF | `Debug_UnitTest_Screen_Default` | Hidden | |

---

## 6. Audio System (MeCCA Framework)

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x0000680F | `"MeCCADecode` | Audio |
| 0x0004F2FC | `FT_Request_Size: bitmap strike %lu matched` | Audio |
| 0x000757E4 | `cI: Soft Reset sequence error` | Audio |
| 0x000A6C43 | `"alac: bit depth = %d, pb = 0x%X, mb = 0x%X, kb = 0x%X ` | Audio |
| 0x000D3140 | `TrackData::EQCache` | Audio |
| 0x0010C418 | `HandleEQSettingSelected` | Audio |
| 0x0011B2A8 | `HandleFrequencyChange` | Audio |
| 0x00136364 | `TCEQSetting` | Audio |
| 0x0013FDE8 | `ShowSetting_EQ` | Audio |
| 0x0013FDFC | `HandleEQ` | Audio |
| 0x001514FC | `AudioCodecs` | Audio |
| 0x001515AC | `AppleLossless` | Audio |
| 0x001515D8 | `Audible` | Audio |
| 0x001523D4 | `RBRequestVersion` | Audio |
| 0x00155F4C | `alacmp4v@KL` | Audio |
| 0x001692B8 | `MeCCA_RecordingBuffer` | Audio |
| 0x00172F38 | `TCSettings_EQ` | Audio |
| 0x00198560 | `MeCCA_PCM_Output.wav` | Audio |
| 0x001B20BC | `MeCCA_MediaPlayer` | Audio |
| 0x001BCE0C | `MeCCA_VideoBufferMgr` | Audio |
| 0x001BD014 | `MeCCAVideoDecode` | Audio |
| 0x001C4A14 | `elsttkhdmdhdstsdsttsstszstscstcomp4aalac` | Audio |
| 0x001E938C | `ToglleQuality` | Audio |
| 0x001F6358 | `MeCCAIOTask` | Audio |
| 0x001F6CF4 | `HandleFrequencyChosen` | Audio |
| 0x00229478 | `ToggleSetting_SoundCheck` | Audio |
| 0x00229798 | `ShowSetting_SoundCheck` | Audio |
| 0x00267DDD | `Requested reserve heap size: %d KB` | Audio |
| 0x002689A0 | `Requested Reserve Heap Size` | Audio |
| 0x00287250 | `TCAlarmPropertiesFrequencyMenu` | Audio |
| 0x002C8980 | `collseq(%.20s)` | Audio |
| 0x003713E8 | `sqlite_sequence` | Audio |
| 0x003713F8 | `UPDATE "%w".sqlite_sequence set name = %Q WHERE name = %Q` | Audio |
| 0x00377C18 | `DELETE FROM %s.sqlite_sequence WHERE name=%Q` | Audio |
| 0x00378458 | `CREATE TABLE %Q.sqlite_sequence(name,seq)` | Audio |
| 0x0037867C | `callback requested query abort` | Audio |
| 0x00378868 | `library routine called out of sequence` | Audio |
| 0x00379E50 | `no such collation sequence: %s` | Audio |
| 0x0037EC94 | `unable to use function %s in the requested context` | Audio |
| 0x0037F3F8 | `no such collation sequence: %.*s` | Audio |

---

## 7. Event Handlers

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x000FF928 | `HandleWheel` | Handler |
| 0x0010CB0C | `HandleNotesSelected` | Handler |
| 0x0010CB24 | `HandleNotesPop` | Handler |
| 0x0010CB34 | `HandleNotesPopToMainMenu` | Handler |
| 0x0010CE28 | `GotoNowPlaying` | Handler |
| 0x0011B270 | `HandleVolumeWheel` | Handler |
| 0x0011B284 | `HandleVolumeChange` | Handler |
| 0x0011B2A8 | `HandleFrequencyChange` | Handler |
| 0x00125DB4 | `GotoScreen_LockDialog` | Handler |
| 0x00125DCC | `GotoScreen_SetCombinationFirstTime` | Handler |
| 0x00125E74 | `HandleLock` | Handler |
| 0x00125F44 | `GotoScreen_AddressBook` | Handler |
| 0x0013135C | `HandleMenuOnKeyboard` | Handler |
| 0x00131D54 | `GotoAlbums` | Handler |
| 0x00131D60 | `GotoSongs` | Handler |
| 0x001403D4 | `GotoScreen_EnterPassKey` | Handler |
| 0x001403EC | `GotoScreen_LockediPod` | Handler |
| 0x00140DF0 | `GotoScreen_MainMenu` | Handler |
| 0x00157ECC | `HandleMikeyCenter` | Handler |
| 0x0019EC14 | `HandleBrowseAlbum` | Handler |
| 0x001CC9EC | `HandleBrowseLarge` | Handler |
| 0x001CCA00 | `HandleBrowseSmall` | Handler |
| 0x001CCA14 | `HandleBrowseSlideshow` | Handler |
| 0x001CF3DC | `HandleRemotePlayPause` | Handler |
| 0x001CF3F4 | `HandleRemotePlayPausePressAndHold` | Handler |
| 0x001CF418 | `HandleRemoteNextAlbum` | Handler |
| 0x001CF430 | `HandleRemotePrevAlbum` | Handler |
| 0x001CF448 | `HandleRemoteVolumeUp` | Handler |
| 0x001CF5EC | `HandleRemoteVolumeDown` | Handler |
| 0x001CF604 | `HandleRemoteVolumeUpUp` | Handler |
| 0x001CF61C | `HandleRemoteVolumeDownUp` | Handler |
| 0x001CF638 | `HandleRemoteStop` | Handler |
| 0x001CF64C | `HandleRemotePlay` | Handler |
| 0x001CF660 | `HandleRemotePause` | Handler |
| 0x001CF674 | `HandleRemoteMute` | Handler |
| 0x001CF688 | `HandleRemoteNextChapter` | Handler |
| 0x001CF6A0 | `HandleRemotePrevChapter` | Handler |
| 0x001CF6B8 | `HandleRemoteNextPlaylist` | Handler |
| 0x001CF6D4 | `HandleRemotePrevPlaylist` | Handler |
| 0x001CF8DC | `HandleRemoteShuffle` | Handler |
| 0x001CF8F0 | `HandleRemoteRepeat` | Handler |
| 0x001CF904 | `HandleRemoteOn` | Handler |
| 0x001CF918 | `HandleRemoteOff` | Handler |
| 0x001CF928 | `HandleRemoteBacklight` | Handler |
| 0x001CF940 | `HandleRemoteFFDown` | Handler |
| 0x001CF954 | `HandleRemoteFFUp` | Handler |
| 0x001CF968 | `HandleRemoteRewDown` | Handler |
| 0x001CF97C | `HandleRemoteRewUp` | Handler |
| 0x001CF990 | `HandleRemoteMenuDown` | Handler |
| 0x001CF9A8 | `HandleRemoteMenuUp` | Handler |
| 0x001CF9BC | `HandleRemoteSelectDown` | Handler |
| 0x001CF9D4 | `HandleRemoteSelectUp` | Handler |
| 0x001CFB84 | `HandleRemoteUpArrowDown` | Handler |
| 0x001CFB9C | `HandleRemoteUpArrowUp` | Handler |
| 0x001CFBB4 | `HandleRemoteDownArrowDown` | Handler |
| 0x001CFBD0 | `HandleRemoteDownArrowUp` | Handler |
| 0x001CFBE8 | `HandleRemoteEvent` | Handler |
| 0x001CFBFC | `HandleRemoteBacklightOff` | Handler |
| 0x001CFC18 | `HandleAudioPlayPause` | Handler |
| 0x001CFC30 | `HandleAudioNext` | Handler |
| 0x001CFC40 | `HandleAudioNextPressAndHold` | Handler |
| 0x001CFC5C | `HandleAudioPrevious` | Handler |
| 0x001CFC70 | `HandleAudioPreviousPressAndHold` | Handler |
| 0x001CFE00 | `HandleAudioNextAlbum` | Handler |
| 0x001CFE18 | `HandleAudioPrevAlbum` | Handler |
| 0x001CFE30 | `HandleAudioVolumeDown` | Handler |
| 0x001CFE48 | `HandleAudioVolumeUp` | Handler |
| 0x001CFE5C | `HandleAudioVolumeDownUp` | Handler |
| 0x001CFE74 | `HandleAudioVolumeUpUp` | Handler |
| 0x001CFE8C | `HandleAudioStop` | Handler |
| 0x001CFE9C | `HandleAudioPlay` | Handler |
| 0x001CFEAC | `HandleAudioPause` | Handler |
| 0x001CFEC0 | `HandleAudioMute` | Handler |
| 0x001CFED0 | `HandleAudioNextChapter` | Handler |
| 0x001CFEE8 | `HandleAudioPrevChapter` | Handler |
| 0x001D00D4 | `HandleAudioNextPlaylist` | Handler |
| 0x001D00EC | `HandleAudioPrevPlaylist` | Handler |
| 0x001D0104 | `HandleAudioShuffle` | Handler |
| 0x001D0118 | `HandleAudioRepeat` | Handler |
| 0x001D012C | `HandleAudioFFDown` | Handler |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x000091AB | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | RTXC |
| 0x002C0F3D | `** Clock Snapshot **` | RTXC |
| 0x002C13DD | `** RTXCbug - ` | RTXC |
| 0x002C13F4 | `  K - RTXC` | RTXC |
| 0x002C1420 | `  X - Exit RTXCbug` | RTXC |
| 0x002C1435 | `RTXCbug> ` | RTXC |
| 0x002C1559 | `** Mailbox Snapshot **` | RTXC |
| 0x002C179D | `** Queue Snapshot **` | RTXC |
| 0x002C1E11 | `RTXCbug - RTXC Objects> ` | RTXC |
| 0x002C202D | `** Semaphore Snapshot **` | RTXC |
| 0x002C2365 | `** Stack Snapshot **` | RTXC |
| 0x002C23FC | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | RTXC |
| 0x002C2679 | `** Task Snapshot **` | RTXC |
| 0x002DB6F0 | `X - Exit RTXCbug` | RTXC |
| 0x003960E5 | `$RTXCbug> ` | RTXC |
| 0x003961AD | `Re-entering RTXCbug mode` | RTXC |
| 0x003F4E0D | `S_RTXCBUG` | RTXC |
| 0x009BC360 | `Returning from RTXCBug` | RTXC |

---

## 9. Logging Channels

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x002669E4 | `Channel Reserved` | Logging |
| 0x002669F8 | `Channel AppBoot` | Logging |
| 0x00266A08 | `Channel BufferedSongReading` | Logging |
| 0x00266A24 | `Channel PrefsWriting` | Logging |
| 0x00266A3C | `Channel GeneralUserExperience` | Logging |
| 0x00266A5C | `Channel PlayFromDisk` | Logging |
| 0x00266A74 | `Channel CacheSpinupDrive` | Logging |
| 0x00266A90 | `Channel TestLogging` | Logging |
| 0x00266AA4 | `Channel AppFileLoading` | Logging |
| 0x00266ABC | `Channel VCardReading` | Logging |
| 0x00266AD4 | `Channel LongSongScanning` | Logging |
| 0x00266B48 | `Channel VoiceRecording` | Logging |
| 0x00266B60 | `Channel PhotoImporting` | Logging |
| 0x00266B78 | `Channel Notes` | Logging |
| 0x00266B88 | `Channel PhotoFileManagement` | Logging |
| 0x00266BA4 | `Channel DiskMode` | Logging |
| 0x00266BB8 | `Channel Firewire` | Logging |
| 0x00266BCC | `Channel USB` | Logging |
| 0x00266BD8 | `Channel UnitTests` | Logging |
| 0x00266BEC | `Channel FreeSpaceCache` | Logging |
| 0x00266C04 | `Channel OnTheGoFileMgmt` | Logging |

---

## 10. Game System

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x000987A8 | `gamedata_RW` | Game |
| 0x000987B4 | `gamestats_WO` | Game |
| 0x000987C4 | `gamedata_ShareRW` | Game |
| 0x000987D8 | `games_RO` | Game |
| 0x000A27E0 | `Resources/Games` | Game |
| 0x000AC208 | `TCGameScreen` | Game |
| 0x000AC220 | `TCGamesMenu` | Game |
| 0x0010C85C | `HandleGameHilited` | Game |
| 0x001261DC | `StartGame` | Game |
| 0x00151044 | `BuiltInGames` | Game |
| 0x00152A28 | `GamesPlatformID` | Game |
| 0x00152A38 | `GamesPlatformVersion` | Game |
| 0x0073B176 | `Games_Menu_Screen` | Game |
| 0x0073B18B | `Games_Menu_Screen_Default"` | Game |
| 0x0073B1F5 | `Extras_Screen_Games` | Game |
| 0x007454DC | `MainMenus_Main_Screen_Games` | Game |
| 0x0075B12C | `TCGameScreenTC_LockediPod` | Game |
| 0x0075E9F4 | `Game_Screen` | Game |
| 0x0075EA03 | `Game_Screen_Default` | Game |
| 0x0075EAA5 | `Game_Memory_Error_Screen` | Game |
| 0x0075EB07 | `Game_Signing_Error_Screen` | Game |
| 0x0075EB6A | `Game_Unknown_Error_Screen` | Game |
| 0x0075EBCD | `Game_Version_Error_Screen` | Game |
| 0x0075EBEA | `controller.StartGame1` | Game |
| 0x0075EC29 | `Game_Running_Screen` | Game |
| 0x008040A0 | `Games` | Game |
| 0x00805CE4 | `No Games` | Game |
| 0x00805CF0 | `1 Game` | Game |
| 0x00805CF8 | `%d Games` | Game |
| 0x00805D48 | `This version of the game is no longer supported.` | Game |
| 0x00805D7C | `This game cannot be launched.` | Game |
| 0x00805E04 | `Connect your iPod to iTunes and reinstall the game.` | Game |
| 0x00990143 | `11TCGamesMenu` | Game |
| 0x00990217 | `12TCGameScreen` | Game |
| 0x00990FE3 | `27TSilverCntlrTransitionAddonI11TCGamesMenuE` | Game |
| 0x00991098 | `27TSilverCntlrTransitionAddonI12TCGameScreenE` | Game |
| 0x009BC295 | `Resources/Games/` | Game |
| 0x009BC3BA | `iPod_Control/games_RO/` | Game |
| 0x009BC3D1 | `Resources/Games/games_RO/` | Game |
| 0x009BFBB7 | `Games_PlayCount_Once` | Game |

---

## 11. Database System

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x00069230 | ` rtSPhotos\Photo Database` | Database |
| 0x0006926C | `iPod_Control\iTunes\iTunesDB` | Database |
| 0x00098738 | `iPod_Control/iTunes/iTunesDB.p7b` | Database |
| 0x000B5738 | `iPod_Control\iTunes\Play Counts` | Database |
| 0x0011F7A0 | `iTunes Image DB.itdb` | Database |
| 0x002255C0 | `iTunes Image DB` | Database |
| 0x00238964 | `%s/sqlite_` | Database |
| 0x00267D31 | `Photo database size: %d KB` | Database |
| 0x00267D51 | `Photo database num photos: %d` | Database |
| 0x0026895C | `Photo Database Size` | Database |
| 0x002C4EE8 | `sqlite3BtreeInitPage() returns error code %d` | Database |
| 0x002C8200 | `sqlite_master` | Database |
| 0x002C8210 | `sqlite_temp_master` | Database |
| 0x002DE3D0 | `sqlite_stat1` | Database |
| 0x002DE3E0 | `CREATE TABLE %Q.sqlite_stat1(tbl,idx,stat)` | Database |
| 0x002DE40C | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Database |
| 0x002E9148 | `sqlite_subquery_%p_` | Database |
| 0x00371380 | `sqlite_` | Database |
| 0x003713E8 | `sqlite_sequence` | Database |
| 0x003713F8 | `UPDATE "%w".sqlite_sequence set name = %Q WHERE name = %Q` | Database |
| 0x003714EC | `SELECT idx, stat FROM %Q.sqlite_stat1` | Database |
| 0x003763DC | `sqlite_autoindex_` | Database |
| 0x003778AC | `DELETE FROM %Q.sqlite_stat1 WHERE idx=%Q` | Database |
| 0x00377C18 | `DELETE FROM %s.sqlite_sequence WHERE name=%Q` | Database |
| 0x00378458 | `CREATE TABLE %Q.sqlite_sequence(name,seq)` | Database |
| 0x0037F2CC | `sqlite3_extension_init` | Database |
| 0x00383EB0 | `sqlite_attach` | Database |
| 0x00383EC4 | `sqlite_detach` | Database |
| 0x003EF904 | `iTunesDB` | Database |
| 0x0080319C | `Richard Hipp (SQLite)` | Database |

---

## 12. Hardware Interfaces

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x00075748 | `cI: Set drive to MMC high speed failed` | Hardware |
| 0x00075804 | `cI: could not read CE-ATA task file` | Hardware |
| 0x0007582C | `cI: CE-ATA signature missing (%x,%x)` | Hardware |
| 0x00075884 | `cI: CE-ATA interrupt enable failed` | Hardware |
| 0x000EEE50 | `mI: card not in MMC TRAN state as expected` | Hardware |
| 0x0013C4B8 | `NAND FLASH DRIVE` | Hardware |
| 0x00149E6C | `USBDeviceTask` | Hardware |
| 0x001645E8 | `ATAWorkLoopTask` | Hardware |
| 0x001645FC | `ATAWorkLoopIRQTask` | Hardware |
| 0x002AF250 | `FirewireTask` | Hardware |
| 0x002AF264 | `TouchwheelTask` | Hardware |
| 0x002AF2A4 | `DiskMgrTask` | Hardware |
| 0x002AF2B4 | `HoldSwitchTask` | Hardware |
| 0x002AF2C8 | `MikeyTask` | Hardware |
| 0x0036C374 | `MMC init failed` | Hardware |
| 0x0036C388 | `CE-ATA init failed` | Hardware |
| 0x0036C848 | `ISDIE: CE-ATA interrupt enable failed` | Hardware |
| 0x003F03F9 | `M_DISKMGR` | Hardware |
| 0x003F050D | `GPIO_REG_WRITE` | Hardware |
| 0x003F051E | `GPIO_INT_INIT` | Hardware |
| 0x003F0551 | `I2C_MASTER` | Hardware |
| 0x003F4AEE | `S_DISKMGRQ` | Hardware |
| 0x003F4E62 | `S_I2C_DONE` | Hardware |
| 0x009BC27F | `HoldSwitch` | Hardware |

---

## 13. Security System

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x0005FE88 | `RSA-SHA1` | Security |
| 0x0005FE94 | `RSA-SHA1-2` | Security |
| 0x0007027C | `NO X509_NAME` | Security |
| 0x001514D0 | `AppleDRMVersion` | Security |
| 0x00151570 | `AppleDRM` | Security |
| 0x00275188 | `http://www.w3.org/2000/09/xmldsig#rsa-sha1` | Security |
| 0x003F05B7 | `SERIALVERIFIER` | Security |
| 0x003F05C8 | `RESISTORVERIFIER` | Security |
| 0x0090CBDF | `X509Data` | Security |
| 0x0090CBE8 | `X509Certificate` | Security |
| 0x009BC8AA | `DRMLevel` | Security |
| 0x009D8F42 | `X509` | Security |
| 0x009D8F47 | `X509_CINF` | Security |
| 0x009D90D1 | `pkcs1` | Security |
| 0x009D9243 | `pkcs3` | Security |
| 0x009D92F7 | `pkcs5` | Security |
| 0x009D935C | `pkcs7` | Security |
| 0x009D941B | `pkcs9` | Security |
| 0x009D9BDB | `RSA Data Security, Inc. PKCS` | Security |
| 0x009D9E45 | `pkcs7-signedData` | Security |
| 0x009D9E56 | `pkcs7-signedAndEnvelopedData` | Security |
| 0x009D9E73 | `pkcs7-envelopedData` | Security |
| 0x009D9E87 | `pkcs7-encryptedData` | Security |
| 0x009D9FEF | `pkcs7-digestData` | Security |
| 0x009DA01C | `pkcs7-data` | Security |
| 0x009DA339 | `X509v3 Private Key Usage Period` | Security |
| 0x009DA515 | `X509v3 CRL Reason Code` | Security |
| 0x009DA585 | `X509v3 Key Usage` | Security |
| 0x009DA596 | `X509v3 Extended Key Usage` | Security |
| 0x009DA60F | `X509v3 No Revocation Available` | Security |

---

## 14. Usage Analytics / Logging

| Offset | Symbol | Description |
|--------|--------|-------------|
| 0x00267419 | `As a percent of the total log time: %d%%` | Analytics |
| 0x002674D1 | `Time in DMIA as percent of total log time: %d%%` | Analytics |
| 0x002675A0 | `iPod Usage Stats` | Analytics |
| 0x002680A1 | `Total log length is %d seconds` | Analytics |
| 0x00268331 | `As percent of total log time: %d%%` | Analytics |
| 0x00268750 | `Disk Spinup` | Analytics |
| 0x002687DC | `Playback Begin` | Analytics |
| 0x002687FC | `Playback Pause` | Analytics |
| 0x0026881C | `Backlight On` | Analytics |
| 0x0026883C | `Flush Usage Log Data` | Analytics |
| 0x002688BC | `Enter Disk Mode` | Analytics |
| 0x002688DC | `Enter Light Sleep` | Analytics |
| 0x00268904 | `Enter Deep Sleep` | Analytics |

---

## 15. Filesystem Paths

| Offset | Path | Description |
|--------|------|-------------|
| 0x00047DEC | `iPod_Control` | Path |
| 0x00047E00 | `iPod_Control\iTunes\` | Path |
| 0x00047E18 | `iPod_Control\iTunes\firsttime` | Path |
| 0x00047E58 | `iPod_Control\Device` | Path |
| 0x000592A8 | `iPod_Control\Device\SysInfo` | Path |
| 0x0006924C | `iPod_Control\Artwork\ArtworkDB` | Path |
| 0x0006926C | `iPod_Control\iTunes\iTunesDB` | Path |
| 0x0006BE94 | `iPod_Control\Music\` | Path |
| 0x0006EED4 | `iPod_Control\Device\Preferences` | Path |
| 0x00080238 | `Contacts/` | Path |
| 0x000946E0 | `iPod_Control/iTunes/` | Path |
| 0x00098738 | `iPod_Control/iTunes/iTunesDB.p7b` | Path |
| 0x000987A8 | `gamedata_RW` | Path |
| 0x000987B4 | `gamestats_WO` | Path |
| 0x000987C4 | `gamedata_ShareRW` | Path |
| 0x000A27F0 | `iPod_Control/%s%s%s` | Path |
| 0x000AB7E8 | `Calendars/` | Path |
| 0x000B5738 | `iPod_Control\iTunes\Play Counts` | Path |
| 0x000C76C4 | `Recordings/` | Path |
| 0x000EA0AC | `Notes/` | Path |
| 0x00100354 | `iPod_Control\Device\dst` | Path |
| 0x0010B48C | `iPod_Control/Device/alarms` | Path |
| 0x0011A978 | `NOTES/` | Path |
| 0x0011BD40 | `iPod_Control/Device/radio` | Path |
| 0x0011D280 | `iPod_Control/Device` | Path |
| 0x00137644 | `iPod_Control/Device/Users` | Path |
| 0x00165A8C | `/iPod_Control/Device/1da` | Path |
| 0x00196B94 | `/iPod_Control/Device/Accessories` | Path |
| 0x00197AC0 | `/iPod_Control/Device/Accessories/Tags` | Path |
| 0x001BDA84 | `iPod_Control/Device/PlayCounts` | Path |

---

## 16. Screen Names

| Offset | Screen | Description |
|--------|--------|-------------|
| 0x0016ED34 | `CoverFlow_Screen` | Screen |
| 0x002299B4 | `GotoVolumeLimit_or_Lock_Screen` | Screen |
| 0x002299D4 | `GotoVolumeLimit_or_Lock_or_EU_Screen` | Screen |
| 0x0073AE46 | `Clock_Screen` | Screen |
| 0x0073AE56 | `Clock_Screen_Default"` | Screen |
| 0x0073AEBB | `Extras_Screen_WorldClock` | Screen |
| 0x0073AF19 | `Calendar_Menu_Screen` | Screen |
| 0x0073AF31 | `Calendar_Menu_Screen_Default"` | Screen |
| 0x0073AF9E | `Extras_Screen_Calendar` | Screen |
| 0x0073B03C | `Extras_Screen_AddressBook` | Screen |
| 0x0073B09B | `Alarms_Menu_Screen` | Screen |
| 0x0073B0B1 | `Alarms_Menu_Screen_Default"` | Screen |
| 0x0073B11C | `Extras_Screen_Alarms` | Screen |
| 0x0073B176 | `Games_Menu_Screen` | Screen |
| 0x0073B18B | `Games_Menu_Screen_Default"` | Screen |
| 0x0073B1F5 | `Extras_Screen_Games` | Screen |
| 0x0073B2B4 | `Extras_Screen_Notes` | Screen |
| 0x0073B378 | `Extras_Screen_Lock` | Screen |
| 0x0073B441 | `Extras_Screen_Stopwatch` | Screen |
| 0x0073B49E | `Debug_MainMenu_Screen` | Screen |
| 0x0073B4B7 | `Debug_MainMenu_Screen_Default"` | Screen |
| 0x0073B525 | `Extras_Screen_Debug` | Screen |
| 0x0073B664 | `Clock_Africa_City_Screen ` | Screen |
| 0x0073B680 | `Clock_Africa_City_Screen_Default` | Screen |
| 0x0073B704 | `Clock_Asia_City_Screen` | Screen |
| 0x0073B71E | `Clock_Asia_City_Screen_Default` | Screen |
| 0x0073B7A0 | `Clock_Atlantic_City_Screen"` | Screen |
| 0x0073B7BE | `Clock_Atlantic_City_Screen_Default` | Screen |
| 0x0073B844 | `Clock_Australia_City_Screen#` | Screen |
| 0x0073B863 | `Clock_Australia_City_Screen_Default` | Screen |
| 0x0073B8EA | `Clock_Europe_City_Screen ` | Screen |
| 0x0073B906 | `Clock_Europe_City_Screen_Default` | Screen |
| 0x0073B98A | `Clock_NorthAmerica_City_Screen&` | Screen |
| 0x0073B9AC | `Clock_NorthAmerica_City_Screen_Default` | Screen |
| 0x0073BA36 | `Clock_Pacific_City_Screen!` | Screen |
| 0x0073BA53 | `Clock_Pacific_City_Screen_Default` | Screen |
| 0x0073BAD8 | `Clock_SouthAmerica_City_Screen&` | Screen |
| 0x0073BAFA | `Clock_SouthAmerica_City_Screen_Default` | Screen |
| 0x0073BB87 | `Clock_Screen"` | Screen |
| 0x00743FCE | `Settings_DateTime_SetDate_Screen(` | Screen |
| 0x00743FF2 | `Settings_DateTime_SetDate_Screen_Default"` | Screen |
| 0x0074406B | `Settings_DateTime_Screen_Default` | Screen |
| 0x007440D1 | `Settings_DateTime_SetTime_Screen(` | Screen |
| 0x007440F5 | `Settings_DateTime_SetTime_Screen_Default"` | Screen |
| 0x0074416E | `Settings_DateTime_Time_Screen_Default` | Screen |
| 0x007441D9 | `Settings_DateTime_SetTimeZone_Screen,` | Screen |
| 0x00744201 | `Settings_DateTime_SetTimeZone_Screen_Default"` | Screen |
| 0x0074427E | `Settings_DateTime_TimeZone_Screen_Default` | Screen |
| 0x007449F8 | `Search_Main_Screen` | Screen |
| 0x00744A0E | `Search_Main_Screen_NoKeyboard"` | Screen |
| 0x00744FAE | `Extras_Screen` | Screen |
| 0x00744FBF | `Extras_Screen_WorldClock"` | Screen |
| 0x0074503C | `MainMenus_Main_Screen_Extras` | Screen |
| 0x007450AE | `Clock_Screen_Default` | Screen |
| 0x00745135 | `MainMenus_Main_Screen_WorldClock` | Screen |
| 0x0074521C | `MainMenus_Main_Screen_Alarms` | Screen |
| 0x00745303 | `MainMenus_Main_Screen_Calendar` | Screen |
| 0x00745367 | `AddressViewer_Main_Screen!` | Screen |
| 0x00745384 | `AddressViewer_Main_Screen_Default"` | Screen |
| 0x007453F6 | `MainMenus_Main_Screen_AddressBook` | Screen |

---

## 17. Binary Structure (for Ghidra)

| Parameter | Value |
|-----------|-------|
| **Architecture** | ARM (32-bit, little-endian) |
| **Processor** | ARM926EJ-S (ARMv5TEJ) |
| **Base Address** | 0x00000000 |
| **Entry Point** | Vector table at offset 0x800 |
| **Code Region** | 0x800 - ~0xA2FFFF |
| **String Tables** | 0x007Fxxxx - 0x009Fxxxx |
| **RTTI Data** | 0x0098xxxx - 0x009Exxxx |
| **UI Layout Data** | 0x0073xxxx - 0x007Exxxx |
| **Localization** | 0x0080xxxx - 0x0090xxxx |

---

## 18. Changes from 2.0.4

| Change | Details |
|--------|---------|
| EU Volume Limit | New feature: full EU-mandated volume limiter with confirmation dialogs |
| Build number | 75 → 247 |
| FreeType2 | Updated (new source paths) |
| zlib | Updated (new copyright/license text) |
| SQLite | Updated (new license text) |
| Functions | +131 new functions (23,164 total) |
| Binary size | +34,608 bytes |
| New controller | `TCSettings_EULimitConfirmation` |
| New toggle | `ToggleSetting_RecommendedVolumeLimit` |
| New screens | `SettingsMenus_VolumeLimitEU_Screen`, `SettingsMenus_EUVolume_Confirmation_Screen` |
