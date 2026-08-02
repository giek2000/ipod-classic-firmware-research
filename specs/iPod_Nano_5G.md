# iPod Nano 5th Generation - RetailOS 1.0.2 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.0.2 |
| **IPSW** | iPod_1.0.2_34A20020.ipsw |
| **Device** | iPod Nano 5th Generation (2009, 8/16GB, Click Wheel, Camera, Pedometer, FM Radio) |
| **UpdaterFamilyID** | 34 |
| **Binary Size** | 7,286,720 bytes (6.95 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 7,284,672 bytes |
| **Total Strings (>=4)** | 92,332 |
| **Function Prologues** | 32,358 (ARM: 1,649, Thumb: 30,709) |
| **DRAM References** | 36,681 |
| **Peripheral Refs** | 8,595 |
| **Build** | N33FirmwareWin-261 |
| **SoC** | S5L8730 |
| **Architecture** | ARM Cortex-A8 (ARMv7) |
| **Codename** | N33 |
| **DFU PID** | 0x1231 |
| **SHA-256** | `3269d9eda2e7e7c406d7bfd6895bd83f27603f03a2c9f52c6cf415924e41af81` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0004EE24 | `TCSportTimer` | Known | Controller |
| 0x0004EE3C | `TCSportTimerMenu` | Known | Controller |
| 0x0004EE58 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x0004EE7C | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x0004FAF4 | `TCFirewireUnsupported` | Known | Controller |
| 0x0004FC20 | `TCNotesDispatcher` | Known | Controller |
| 0x0004FC3C | `TCNotesLoading` | Known | Controller |
| 0x0004FC54 | `TCNotesList` | Known | Controller |
| 0x0004FC68 | `TCNotesContents` | Known | Controller |
| 0x000522B0 | `TCDemoMode` | Known | Controller |
| 0x000567B4 | `TCCamera` | Known | Controller |
| 0x000567C8 | `TCCameraInitial` | Known | Controller |
| 0x000567E0 | `TCCameraLocalMediaList` | Known | Controller |
| 0x00056800 | `TCCameraAllVideosList` | Known | Controller |
| 0x00056820 | `TCCameraDeleteAllDialog` | Known | Controller |
| 0x00056840 | `TCCameraDeleteDialog` | Known | Controller |
| 0x000570B8 | `TCPhotosDeleteAllDialog` | Known | Controller |
| 0x000570D8 | `TCPhotosDeleteDialog` | Known | Controller |
| 0x0005778C | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x000577B8 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x000577E4 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x0005780C | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x00057838 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x00057860 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x0005788C | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x00057DC4 | `TCRemoteUI` | Known | Controller |
| 0x00057DD8 | `TCUnsupported` | Known | Controller |
| 0x000581D8 | `TCSpeakers` | Known | Controller |
| 0x000581EC | `TCEQSetting` | Known | Controller |
| 0x00059770 | `TCVoiceMemos` | Known | Controller |
| 0x00059788 | `TCVoiceMemosIdle` | Known | Controller |
| 0x000597A4 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x000597C4 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x000597E4 | `TCVoiceMemosLabelSelectMenu` | Known | Controller |
| 0x00059808 | `TCVoiceMemosLoading` | Known | Controller |
| 0x00059824 | `TCVoiceMemosTimedStatus` | Known | Controller |
| 0x00059DBC | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x00059DE4 | `TCSettings_MainMenu` | Known | Controller |
| 0x00059E00 | `TCSettings_MusicMenu` | Known | Controller |
| 0x00059E20 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x00059E40 | `TCSettings_VolumeLimit_Dialogue` | Known | Controller |
| 0x00059E68 | `TCSettings_Brightness` | Known | Controller |
| 0x00059E88 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x00059EAC | `TCSettings_EQ` | Known | Controller |
| 0x00059EC4 | `TCSettings_RadioRegions` | Known | Controller |
| 0x00059EE4 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x00059F08 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x00059F2C | `TCDateTimeScreen` | Known | Controller |
| 0x00059F48 | `TCTimeZoneScreen` | Known | Controller |
| 0x00059F64 | `TCAddressViewerLoadingScreenCntlr` | Known | Controller |
| 0x00059F90 | `TCAddressViewerNoContactsCntlr` | Known | Controller |
| 0x00059FDC | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0005A004 | `TCAboutCntlr` | Known | Controller |
| 0x0005A01C | `TCSettings_Language` | Known | Controller |
| 0x0005AFE8 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x0005B008 | `TCAddressViewerDetails` | Known | Controller |
| 0x0005B028 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x0005B04C | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x0005B074 | `TCAddressViewerContactGroups` | Known | Controller |
| 0x0005C5E4 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x0005C608 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x0005E9D8 | `TC_LockDialog` | Known | Controller |
| 0x0005E9F0 | `TC_LockScreen` | Known | Controller |
| 0x0005EA08 | `TC_LockediPod` | Known | Controller |
| 0x0005EA20 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x0005EA44 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0005EA64 | `TCResetCombinationChosenDispatcher` | Known | Controller |
| 0x0005EA90 | `TCLockAppMenu` | Known | Controller |
| 0x0005EEF0 | `TCClock` | Known | Controller |
| 0x0005EF00 | `TCClockCityMenu` | Known | Controller |
| 0x0005EF18 | `TCClockRegionMenu` | Known | Controller |
| 0x0005EF34 | `TCAlarmMenu` | Known | Controller |
| 0x0005EF48 | `TCSleepTimerMenu` | Known | Controller |
| 0x0005EF64 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x0005EF84 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x0005EFAC | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x0005EFD0 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0005EFF4 | `TCAlarmDatePicker` | Known | Controller |
| 0x0005F010 | `TCAlarmTriggered` | Known | Controller |
| 0x0007262E | `TC3Fd` | Known | Controller |
| 0x00077D6C | `TSilverCntlr` | Known | Controller |
| 0x00077D84 | `TCExtrasMenu` | Known | Controller |
| 0x00077D9C | `TCGameScreen` | Known | Controller |
| 0x00077DB4 | `TCGameControls` | Known | Controller |
| 0x00077DCC | `TCGamesMenu` | Known | Controller |
| 0x00077DE0 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x00077E08 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00077E30 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x00077E5C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x00077E80 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x00077EA8 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x00077ED0 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x00077EF8 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x00077F20 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x00077F48 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x00077F78 | `TSilverMediaListCntlr_iTunesU` | Known | Controller |
| 0x00077FA0 | `TSilverMediaListCntlr_iTunesUEpisodes` | Known | Controller |
| 0x00077FD0 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x00077FFC | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x0007802C | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x00078054 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x0007807C | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x000780A8 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x000780D0 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x000780F8 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x00078128 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x00078158 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x00078284 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x000782B4 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x000782DC | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x00078304 | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x00078330 | `TCRentalNotification` | Known | Controller |
| 0x00078350 | `TCRentalInfo` | Known | Controller |
| 0x00078368 | `TCRentalConfirmDelete` | Known | Controller |
| 0x00078388 | `TCRentalDispatcher` | Known | Controller |
| 0x000783A4 | `TSilverOverlayCntlr` | Known | Controller |
| 0x0007843C | `TSilverGlobalCntlr` | Known | Controller |
| 0x00078458 | `TCGlobalCoverFlowEntry` | Known | Controller |
| 0x000EAE80 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x0045F634 | `TCClock` | Known | Controller |
| 0x0045F63C | `TCAlarmTriggered` | Known | Controller |
| 0x0045F660 | `TSilverCntlr` | Known | Controller |
| 0x0045F670 | `TCClockRegionMenu` | Known | Controller |
| 0x0045F684 | `TCClockCityMenu` | Known | Controller |
| 0x0045F694 | `TCAlarmMenu` | Known | Controller |
| 0x0045F6A0 | `TCSleepTimerMenu` | Known | Controller |
| 0x0045F6B4 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x0045F6CC | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x0045F6EC | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x0045F708 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0045F724 | `TCAlarmDatePicker` | Known | Controller |
| 0x0045F84C | `TSilverCntlr` | Known | Controller |
| 0x0045F85C | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x0045F87C | `TCSettings_Brightness` | Known | Controller |
| 0x0045F894 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0045F8B0 | `TCSettings_RadioRegions` | Known | Controller |
| 0x0045F8C8 | `TCSettings_EQ` | Known | Controller |
| 0x0045F8D8 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0045F8F4 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0045F914 | `TCAboutCntlr` | Known | Controller |
| 0x0045F924 | `TCSettings_Language` | Known | Controller |
| 0x0045F938 | `TCSettings_MainMenu` | Known | Controller |
| 0x0045F94C | `TCSettings_MusicMenu` | Known | Controller |
| 0x0045F964 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x0045F97C | `TCSettings_VolumeLimit_Dialogue` | Known | Controller |
| 0x0045F99C | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x0045F9B8 | `TCDateTimeScreen` | Known | Controller |
| 0x0045F9CC | `TCTimeZoneScreen` | Known | Controller |
| 0x0045F9E0 | `TCAddressViewerLoadingScreenCntlr` | Known | Controller |
| 0x0045FA04 | `TCAddressViewerNoContactsCntlr` | Known | Controller |
| 0x00462600 | `TSilverCntlr` | Known | Controller |
| 0x00462610 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x00462634 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x00462658 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x00462678 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x0046269C | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x004626BC | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x004626E0 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x00462880 | `TSilverCntlr` | Known | Controller |
| 0x00462890 | `TCCameraInitial` | Known | Controller |
| 0x004628A0 | `TCCamera` | Known | Controller |
| 0x004628AC | `TCCameraLocalMediaList` | Known | Controller |
| 0x004628C4 | `TCCameraAllVideosList` | Known | Controller |
| 0x004628DC | `TCCameraDeleteAllDialog` | Known | Controller |
| 0x004628F4 | `TCCameraDeleteDialog` | Known | Controller |
| 0x00463714 | `TSilverCntlr` | Known | Controller |
| 0x00463724 | `TC_LockDialog` | Known | Controller |
| 0x00463734 | `TC_LockScreen` | Known | Controller |
| 0x00463744 | `TC_LockediPod` | Known | Controller |
| 0x00463754 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x00463770 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00463788 | `TCResetCombinationChosenDispatcher` | Known | Controller |
| 0x004637AC | `TCLockAppMenu` | Known | Controller |
| 0x004637CC | `TSilverCntlr` | Known | Controller |
| 0x004637DC | `TCFirewireUnsupported` | Known | Controller |
| 0x00463A4C | `TCRemoteUI` | Known | Controller |
| 0x00463A58 | `TCUnsupported` | Known | Controller |
| 0x00463AE8 | `TSilverCntlr` | Known | Controller |
| 0x004640F4 | `TSilverCntlr` | Known | Controller |
| 0x00464320 | `TCDemoMode` | Known | Controller |
| 0x00464424 | `TSilverCntlr` | Known | Controller |
| 0x00464434 | `TCVoiceMemosIdle` | Known | Controller |
| 0x00464448 | `TCVoiceMemos` | Known | Controller |
| 0x00464458 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x00464470 | `TCVoiceMemosLabelSelectMenu` | Known | Controller |
| 0x0046448C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x004644A4 | `TCVoiceMemosLoading` | Known | Controller |
| 0x004644B8 | `TCVoiceMemosTimedStatus` | Known | Controller |
| 0x004660C0 | `TSilverCntlr` | Known | Controller |
| 0x0047AEFC | `TSilverCntlr` | Known | Controller |
| 0x0047AF0C | `TCAddressViewerMainMenu` | Known | Controller |
| 0x0047AF24 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x0047AF40 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x0047AF60 | `TCAddressViewerDetails` | Known | Controller |
| 0x0047AF78 | `TCAddressViewerContactGroups` | Known | Controller |
| 0x0047B4D8 | `TSilverCalendarCntlr_Base` | Known | Controller |
| 0x0047B7AC | `TCCameraInitial` | Known | Controller |
| 0x0047B7BC | `TCCamera` | Known | Controller |
| 0x0047B7C8 | `TCCameraMediaList_Base` | Known | Controller |
| 0x0047B7F0 | `TCCameraLocalMediaList` | Known | Controller |
| 0x0047B808 | `TCCameraAllVideosList` | Known | Controller |
| 0x0047B820 | `TCCameraDeleteAllDialog` | Known | Controller |
| 0x0047B838 | `TCCameraDeleteDialog` | Known | Controller |
| 0x0047B9BC | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x0047B9EC | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x0047BA0C | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x0047BA2C | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x0047BA60 | `TSilverCntlr` | Known | Controller |
| 0x0047BA70 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x0047BA8C | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x0047BAAC | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0047BACC | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0047BAF4 | `TSilverMediaListCntlr_iTunesU` | Known | Controller |
| 0x0047BB14 | `TSilverMediaListCntlr_iTunesUEpisodes` | Known | Controller |
| 0x0047BB3C | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x0047BB60 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x0047BB88 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x0047BBA8 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x0047BBC8 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x0047BBE8 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x0047BC08 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0047BC30 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x0047BC58 | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x0047BC78 | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x0047BCAC | `TSilverCntlr` | Known | Controller |
| 0x0047BE28 | `TCNotesDispatcher` | Known | Controller |
| 0x0047BE3C | `TCNotesLoading` | Known | Controller |
| 0x0047BE4C | `TCNotesList` | Known | Controller |
| 0x0047BE58 | `TCNotesContents` | Known | Controller |
| 0x0047BE68 | `TSilverCntlr` | Known | Controller |
| 0x0047BF70 | `TCPhotosDeleteAllDialog` | Known | Controller |
| 0x0047BF88 | `TCPhotosDeleteDialog` | Known | Controller |
| 0x0047BFB0 | `TSilverCntlr` | Known | Controller |
| 0x0047C084 | `TSilverCntlr` | Known | Controller |
| 0x0047C184 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x0047C1A0 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x0047C1B8 | `TCSpeakers` | Known | Controller |
| 0x0047C1C4 | `TCEQSetting` | Known | Controller |
| 0x0047C1FC | `TSilverCntlr` | Known | Controller |
| 0x0047C20C | `TCSportTimer` | Known | Controller |
| 0x0047C21C | `TCSportTimerMenu` | Known | Controller |
| 0x0047C230 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x0047C24C | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x0047C2B8 | `TSilverCntlr` | Known | Controller |
| 0x0047C2C8 | `TCVoiceMemosIdle` | Known | Controller |
| 0x0047C2DC | `TCVoiceMemos` | Known | Controller |
| 0x0047C2EC | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x0047C304 | `TCVoiceMemosLabelSelectMenu` | Known | Controller |
| 0x0047C320 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x0047C338 | `TCVoiceMemosLoading` | Known | Controller |
| 0x0047C34C | `TCVoiceMemosTimedStatus` | Known | Controller |
| 0x0047C3A4 | `TSilverCntlr` | Known | Controller |
| 0x0047C3C4 | `TCExtrasMenu` | Known | Controller |
| 0x0047C3D4 | `TCGamesMenu` | Known | Controller |
| 0x0047C3E0 | `TCGameControls` | Known | Controller |
| 0x0047C3F0 | `TCGameScreen` | Known | Controller |
| 0x0047C400 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x0047C41C | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x0047C43C | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0047C45C | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0047C484 | `TSilverMediaListCntlr_iTunesU` | Known | Controller |
| 0x0047C4A4 | `TSilverMediaListCntlr_iTunesUEpisodes` | Known | Controller |
| 0x0047C4CC | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x0047C4F0 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x0047C518 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x0047C538 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x0047C558 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x0047C578 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x0047C598 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0047C5C0 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x0047C5E8 | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x0047C608 | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x0047C62C | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x0047C64C | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x0047C66C | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x0047C690 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x0047C6B0 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x0047C6D0 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x0047C6F4 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x0047C714 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x0047C73C | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x0047C768 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x0047C788 | `TCRentalNotification` | Known | Controller |
| 0x0047C7A0 | `TCRentalInfo` | Known | Controller |
| 0x0047C7B0 | `TCRentalConfirmDelete` | Known | Controller |
| 0x0047C7C8 | `TCRentalDispatcher` | Known | Controller |
| 0x0047C7DC | `TSilverGlobalCntlr` | Known | Controller |
| 0x0047C7F0 | `TCGlobalCoverFlowEntry` | Known | Controller |
| 0x0047C808 | `TSilverOverlayCntlr` | Known | Controller |
| 0x0048AE04 | `TCNotesDispatcher` | Known | Controller |
| 0x0048AE18 | `TCNotesLoading` | Known | Controller |
| 0x0048AE28 | `TCNotesBase` | Known | Controller |
| 0x0048AE44 | `TCNotesList` | Known | Controller |
| 0x0048AE50 | `TCNotesContents` | Known | Controller |
| 0x0056D751 | `TCCameraInitial_InitialLayoutIsAppNotInitialized` | Known | Controller |
| 0x0056D816 | `TCCameraInitial_InitialLayoutIsActive` | Known | Controller |
| 0x0056D882 | `TCCameraInitial_InitialLayoutIsDiskFull` | Known | Controller |
| 0x0056D984 | `TCCameraMediaList_Base_DoDeleteAll` | Known | Controller |
| 0x0056D9A7 | `TCCameraMediaList_Base_DoDeleteItem` | Known | Controller |
| 0x005710EB | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x005AE3CC | `TSilverGlobalCntlr` | Known | Controller |
| 0x005AE3E0 | `TSilverCntlr` | Known | Controller |
| 0x005AE3F0 | `TSilverCntlr` | Known | Controller |
| 0x005AE400 | `TSilverCntlr` | Known | Controller |
| 0x005AE410 | `TSilverCntlr` | Known | Controller |
| 0x005AE420 | `TCGlobalCoverFlowEntry` | Known | Controller |
| 0x005AE438 | `TSilverCntlr` | Known | Controller |
| 0x005AE448 | `TCGlobalCoverFlowEntry` | Known | Controller |
| 0x005AE460 | `TSilverCntlr` | Known | Controller |
| 0x005AE470 | `TSilverCntlr` | Known | Controller |
| 0x005AE490 | `TSilverCntlr` | Known | Controller |
| 0x005AE4A0 | `TSilverCntlr` | Known | Controller |
| 0x005AE4B0 | `TSilverCntlr` | Known | Controller |
| 0x005AE4C0 | `TSilverCntlr` | Known | Controller |
| 0x005AE4D0 | `TSilverCntlr` | Known | Controller |
| 0x005AE4E0 | `TCGlobalCoverFlowEntry` | Known | Controller |
| 0x005AE4F8 | `TSilverCntlr` | Known | Controller |
| 0x005BD14C | `TCFirewireUnsupported` | Known | Controller |
| 0x005BD17C | `TCExtrasMenu` | Known | Controller |
| 0x005BD1B4 | `TCAddressViewerNoContactsCntlr` | Known | Controller |
| 0x005BD1D4 | `TCAddressViewerContactGroups` | Known | Controller |
| 0x005BD1F4 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x005BD20C | `TCAddressViewerDetails` | Known | Controller |
| 0x005BD224 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x005BD240 | `TCAddressViewerLoadingScreenCntlr` | Known | Controller |
| 0x005BD264 | `TCAlarmMenu` | Known | Controller |
| 0x005BD270 | `TCSleepTimerMenu` | Known | Controller |
| 0x005BD284 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x005BD29C | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x005BD2BC | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x005BD2D8 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x005BD2F4 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x005BD310 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x005BD32C | `TCAlarmDatePicker` | Known | Controller |
| 0x005BD340 | `TCAlarmDatePicker` | Known | Controller |
| 0x005BD354 | `TCAlarmTriggered` | Known | Controller |
| 0x005BD368 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x005BD384 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x005BD3A8 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x005BD3CC | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x005BD3EC | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x005BD410 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x005BD430 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x005BD454 | `TSilverCntlr` | Known | Controller |
| 0x005BD480 | `TCCamera` | Known | Controller |
| 0x005BD48C | `TCCameraLocalMediaList` | Known | Controller |
| 0x005BD4A4 | `TCCameraAllVideosList` | Known | Controller |
| 0x005BD4D4 | `TCCameraDeleteAllDialog` | Known | Controller |
| 0x005BD4EC | `TCCameraDeleteDialog` | Known | Controller |
| 0x005BD504 | `TCCameraDeleteDialog` | Known | Controller |
| 0x005BD51C | `TCCameraDeleteDialog` | Known | Controller |
| 0x005BD534 | `TCClockRegionMenu` | Known | Controller |
| 0x005BD548 | `TCClockCityMenu` | Known | Controller |
| 0x005BD558 | `TCClock` | Known | Controller |
| 0x005BD5A0 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x005BD5C0 | `TCDateTimeScreen` | Known | Controller |
| 0x005BD5D4 | `TCDateTimeScreen` | Known | Controller |
| 0x005BD5E8 | `TCTimeZoneScreen` | Known | Controller |
| 0x005BD5FC | `TCDemoMode` | Known | Controller |
| 0x005BD620 | `TCGamesMenu` | Known | Controller |
| 0x005BD62C | `TCGameControls` | Known | Controller |
| 0x005BD63C | `TCGameScreen` | Known | Controller |
| 0x005BD698 | `TCLockAppMenu` | Known | Controller |
| 0x005BD6A8 | `TC_LockediPod` | Known | Controller |
| 0x005BD6B8 | `TC_LockScreen` | Known | Controller |
| 0x005BD6C8 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x005BD6E4 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x005BD704 | `TSilverCntlr` | Known | Controller |
| 0x005BD714 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x005BD734 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x005BD758 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x005BD780 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x005BD79C | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x005BD7BC | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x005BD7DC | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x005BD7FC | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x005BD81C | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x005BD840 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x005BD860 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x005BD880 | `TSilverCntlr` | Known | Controller |
| 0x005BD890 | `TSilverCntlr` | Known | Controller |
| 0x005BD8A0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x005BD8C0 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x005BD8E8 | `TSilverMediaListCntlr_iTunesU` | Known | Controller |
| 0x005BD908 | `TSilverMediaListCntlr_iTunesUEpisodes` | Known | Controller |
| 0x005BD930 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x005BD950 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x005BD978 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x005BD9B8 | `TSilverCntlr` | Known | Controller |
| 0x005BD9C8 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x005BD9EC | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x005BDA0C | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x005BDA2C | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x005BDA4C | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x005BDA6C | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x005BDA90 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x005BDAB0 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x005BDACC | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x005BDAF4 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x005BDB20 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x005BDC9C | `TCRentalInfo` | Known | Controller |
| 0x005BDCAC | `TCRentalConfirmDelete` | Known | Controller |
| 0x005BDCC4 | `TSilverCntlr` | Known | Controller |
| 0x005BDCD4 | `TCRentalNotification` | Known | Controller |
| 0x005BDCEC | `TCRentalNotification` | Known | Controller |
| 0x005BDD04 | `TCRentalNotification` | Known | Controller |
| 0x005BDD1C | `TCNotesLoading` | Known | Controller |
| 0x005BDD2C | `TCNotesList` | Known | Controller |
| 0x005BDD38 | `TCNotesContents` | Known | Controller |
| 0x005BDD48 | `TCNotesContents` | Known | Controller |
| 0x005BDD58 | `TCNotesContents` | Known | Controller |
| 0x005BDDB8 | `TSilverCntlr` | Known | Controller |
| 0x005BDF14 | `TCPhotosDeleteAllDialog` | Known | Controller |
| 0x005BDF2C | `TCPhotosDeleteDialog` | Known | Controller |
| 0x005BDF44 | `TCPhotosDeleteAllDialog` | Known | Controller |
| 0x005BDF5C | `TCPhotosDeleteDialog` | Known | Controller |
| 0x005BE0A0 | `TCRemoteUI` | Known | Controller |
| 0x005BE0AC | `TCUnsupported` | Known | Controller |
| 0x005BE0D8 | `TCAboutCntlr` | Known | Controller |
| 0x005BE0E8 | `TCAboutCntlr` | Known | Controller |
| 0x005BE0F8 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x005BE118 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x005BE138 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x005BE158 | `TCSettings_MainMenu` | Known | Controller |
| 0x005BE16C | `TCSettings_MusicMenu` | Known | Controller |
| 0x005BE184 | `TCShakeAdjust_Intensity` | Known | Controller |
| 0x005BE19C | `TCShakeAdjust_Duration` | Known | Controller |
| 0x005BE1B4 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x005BE1CC | `TCSettings_VolumeLimit_Dialogue` | Known | Controller |
| 0x005BE1EC | `TCSettings_Brightness` | Known | Controller |
| 0x005BE204 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x005BE220 | `TCSettings_RadioRegions` | Known | Controller |
| 0x005BE238 | `TCSettings_EQ` | Known | Controller |
| 0x005BE248 | `TCSettings_Language` | Known | Controller |
| 0x005BE25C | `TSilverCntlr` | Known | Controller |
| 0x005BE26C | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x005BE28C | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x005BE2A8 | `TCSettings_MainMenu` | Known | Controller |
| 0x005BE2BC | `TCSettings_MusicMenu` | Known | Controller |
| 0x005BE2D4 | `TCSportTimer` | Known | Controller |
| 0x005BE2E4 | `TCSportTimerMenu` | Known | Controller |
| 0x005BE2F8 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x005BE628 | `TSilverCntlr` | Known | Controller |
| 0x005BE96C | `TSilverCntlr` | Known | Controller |
| 0x005BEAB8 | `TSilverCntlr` | Known | Controller |
| 0x005BEAC8 | `TSilverCntlr` | Known | Controller |
| 0x005BEB24 | `TSilverCntlr` | Known | Controller |
| 0x005BEB64 | `TSilverCntlr` | Known | Controller |
| 0x005BEB74 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BEB8C | `TCVoiceMemosIdle` | Known | Controller |
| 0x005BEBA0 | `TCVoiceMemos` | Known | Controller |
| 0x005BEBC8 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BEBE0 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x005BEBF8 | `TCVoiceMemosLabelSelectMenu` | Known | Controller |
| 0x005BEC14 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BEC2C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BEC44 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BEC5C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BEC74 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BEC8C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BECA4 | `TCVoiceMemosLoading` | Known | Controller |
| 0x005BECB8 | `TCVoiceMemosTimedStatus` | Known | Controller |
| 0x005BECE0 | `TCSpeakers` | Known | Controller |
| 0x005BECEC | `TCEQSetting` | Known | Controller |
| 0x00634B3C | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x00634E18 | `TCNotesDispatcher` | Known | Controller |
| 0x00634EF0 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x00635788 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x00635C6C | `TCCameraInitial` | Known | Controller |
| 0x00636E70 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00636EA4 | `TCResetCombinationChosenDispatcher` | Known | Controller |
| 0x0063A9BC | `TCRentalDispatcher` | Known | Controller |
| 0x0063AB18 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |

---

## 2. Settings (Toggle/Show/TC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000B008C | `ToggleSetting_PreviewPanel` | Known | User setting |
| 0x000B15E4 | `ToggleSetting_Repeat` | Known | User setting |
| 0x000B1600 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x000B1618 | `ToggleSetting_TVOut` | Known | User setting |
| 0x000B162C | `ToggleSetting_TVSignal` | Known | User setting |
| 0x000B8168 | `ToggleSetting_Audiobook` | Known | User setting |
| 0x000B8184 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x000B819C | `ToggleSetting_Repeat` | Known | User setting |
| 0x000B81B4 | `ToggleSetting_SortBy` | Known | User setting |
| 0x000B81CC | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x000B81E8 | `ToggleSetting_Clicker` | Known | User setting |
| 0x000B8200 | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x000B8220 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x000B823C | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x000B8258 | `ToggleSetting_EnergySaver` | Known | User setting |
| 0x000B8274 | `ToggleSetting_Crossfade` | Known | User setting |
| 0x000B828C | `ToggleSetting_FontSize` | Known | User setting |
| 0x000B82A4 | `ToggleSetting_Shake` | Known | User setting |
| 0x000B82B8 | `ToggleSetting_VoiceFeedback` | Known | User setting |
| 0x000B82D4 | `ToggleSetting_Rotate` | Known | User setting |
| 0x000B82EC | `ToggleSetting_MonoAudio` | Known | User setting |
| 0x000B8304 | `ShowSetting_About` | Known | User setting |
| 0x000B8318 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x000B8330 | `ShowSetting_DateAndTime` | Known | User setting |
| 0x000B8348 | `ShowSetting_Legal` | Known | User setting |
| 0x006357C8 | `ToggleSetting_Alarm` | Known | User setting |
| 0x006369B0 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x006369E8 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x0063BEA4 | `ToggleSetting_Repeat` | Known | User setting |
| 0x0063BED8 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x0063BF6C | `ToggleSetting_TVOut` | Known | User setting |
| 0x0063BF9C | `ToggleSetting_TVSignal` | Known | User setting |
| 0x0063DBF0 | `ShowSetting_About` | Known | User setting |
| 0x0063DF20 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x0063DF58 | `ToggleSetting_MonoAudio` | Known | User setting |
| 0x0063DFC8 | `ToggleSetting_Crossfade` | Known | User setting |
| 0x0063DFFC | `ToggleSetting_Audiobook` | Known | User setting |
| 0x0063E030 | `ToggleSetting_Shake` | Known | User setting |
| 0x0063E060 | `ToggleSetting_EnergySaver` | Known | User setting |
| 0x0063E100 | `ToggleSetting_Clicker` | Known | User setting |
| 0x0063E134 | `ToggleSetting_Rotate` | Known | User setting |
| 0x0063E168 | `ToggleSetting_VoiceFeedback` | Known | User setting |
| 0x0063E240 | `ToggleSetting_FontSize` | Known | User setting |
| 0x0063E33C | `ToggleSetting_SortBy` | Known | User setting |
| 0x0063E370 | `ToggleSetting_PreviewPanel` | Known | User setting |

---

## 3. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0011E0AC | `Channel UnitTests` | Hidden | Developer Tool |
| 0x0051C0FB | `10TCDemoMode` | Hidden | Demo/Retail Mode |
| 0x0051C498 | `12TUnitTestApp` | Hidden | Developer Tool |
| 0x0051E54F | `27TSilverCntlrTransitionAddonI10TCDemoModeE` | Hidden | Demo/Retail Mode |
| 0x005BEB34 | `TUnitTestSuiteCntlr` | Hidden | Developer Tool |
| 0x005BEB48 | `TUnitTestSuiteTestsCntlr` | Hidden | Developer Tool |
| 0x0062A86C | `DemoMode` | Hidden | Demo/Retail Mode |

---

## 4. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00035FC0 | `AudioCodecs` | Known | Audio system |
| 0x0006C904 | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x00164708 | `MeCCA_VdRecBufferMgr` | Known | Audio system |
| 0x0018D118 | `MeCCA_GlobalBMHeap` | Known | Audio system |
| 0x002D2214 | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x002DDF0F | `"MeCCADecode` | Known | Audio system |
| 0x002DDFBC | `MeCCAVideoDecode` | Known | Audio system |
| 0x0049663C | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Toolbox/MeCCA/MediaEngine/Video/Codec` | Known | Audio system |

---

## 5. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00079CF4 | `HandleNotesSelected` | Known | Event handler |
| 0x00079D0C | `HandleNotesPop` | Known | Event handler |
| 0x00079D1C | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0007ACCC | `HandleExit` | Known | Event handler |
| 0x0007ACDC | `HandleLap` | Known | Event handler |
| 0x0007ACE8 | `HandleResume` | Known | Event handler |
| 0x0007ACF8 | `HandleStartStop` | Known | Event handler |
| 0x0007C110 | `HandleNotesPop` | Known | Event handler |
| 0x0007C124 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0007CA64 | `HandleChosen` | Known | Event handler |
| 0x0007EFE8 | `HandleDelete` | Known | Event handler |
| 0x0007EFFC | `HandleSelectLozinch` | Known | Event handler |
| 0x0007F560 | `HandleSelect` | Known | Event handler |
| 0x0007FC60 | `HandleSelect` | Known | Event handler |
| 0x0007FC70 | `HandleSelectRating` | Known | Event handler |
| 0x0007FC84 | `HandleSelectProgress` | Known | Event handler |
| 0x0007FC9C | `HandleWheelProgress` | Known | Event handler |
| 0x0007FCB0 | `HandleSelectScrub` | Known | Event handler |
| 0x0007FCC4 | `HandleWheelBrightness` | Known | Event handler |
| 0x0007FCDC | `HandleNextTrackGotoScrub` | Known | Event handler |
| 0x0007FCF8 | `HandlePrevTrackGotoScrub` | Known | Event handler |
| 0x0007FD14 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x0007FD34 | `HandleMikeyCenter` | Known | Event handler |
| 0x000809F4 | `HandleSelect` | Known | Event handler |
| 0x00080A08 | `HandleWheel` | Known | Event handler |
| 0x00080A14 | `HandleWheelProgress` | Known | Event handler |
| 0x00080A28 | `HandleSelectProgress` | Known | Event handler |
| 0x00080A40 | `HandleSelectVolume` | Known | Event handler |
| 0x00080A54 | `HandleSelectScrub` | Known | Event handler |
| 0x00080A68 | `HandleSelectGenius` | Known | Event handler |
| 0x00080A7C | `HandleSelectRating` | Known | Event handler |
| 0x00080A90 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x00080AA8 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x00080AC4 | `HandleWheelGenius` | Known | Event handler |
| 0x00080AD8 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x00080AF4 | `HandleWheelBrightness` | Known | Event handler |
| 0x00080B0C | `HandleAddToOTG` | Known | Event handler |
| 0x00080B1C | `HandleViewArtist` | Known | Event handler |
| 0x00080B30 | `HandleViewAlbum` | Known | Event handler |
| 0x00080B40 | `HandleViewCompilation` | Known | Event handler |
| 0x00080B58 | `HandleStartGenius` | Known | Event handler |
| 0x00080B6C | `HandleAudiobookSlower` | Known | Event handler |
| 0x00080B84 | `HandleAudiobookFaster` | Known | Event handler |
| 0x00080B9C | `HandleAudiobookNormal` | Known | Event handler |
| 0x00080BB4 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00080BD0 | `HandlePushContextualMenu` | Known | Event handler |
| 0x00080C40 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x00080C60 | `HandleOrientationChange` | Known | Event handler |
| 0x00080C78 | `HandleRotationChange` | Known | Event handler |
| 0x00080C90 | `HandlePlayPauseTV` | Known | Event handler |
| 0x00080CA4 | `HandleSwapToVideoScreen` | Known | Event handler |
| 0x00080CBC | `HandleSwapToMusicScreen` | Known | Event handler |
| 0x00080CD4 | `HandleMikeyCenter` | Known | Event handler |
| 0x00080CE8 | `HandleRemoteMenu` | Known | Event handler |
| 0x0008115C | `HandleAudiobookSlower` | Known | Event handler |
| 0x00081174 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0008118C | `HandleAudiobookFaster` | Known | Event handler |
| 0x000811A4 | `HandleAddToOTG` | Known | Event handler |
| 0x000811B4 | `HandleStartGenius` | Known | Event handler |
| 0x000811CC | `HandleViewCompilation` | Known | Event handler |
| 0x000811E4 | `HandleViewAlbum` | Known | Event handler |
| 0x000811F4 | `HandleViewArtist` | Known | Event handler |
| 0x00081208 | `HandleCancel` | Known | Event handler |
| 0x00081264 | `HandleSelect` | Known | Event handler |
| 0x00081420 | `HandleSelect` | Known | Event handler |
| 0x0008CC94 | `HandleMenuSelection` | Known | Event handler |
| 0x000A39E0 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000A3A00 | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x000A4C40 | `HandleSelect` | Known | Event handler |
| 0x000A4C54 | `HandleHilite` | Known | Event handler |
| 0x000A4EB4 | `HandleEQSettingSelected` | Known | Event handler |
| 0x000A4ED0 | `HandleEQSettingPreview` | Known | Event handler |
| 0x000A5628 | `HandleTunerContextMenu` | Known | Event handler |
| 0x000A56B8 | `HandleVolumeChange` | Known | Event handler |
| 0x000A56CC | `HandleVolumeWheel` | Known | Event handler |
| 0x000A56E0 | `HandleTunerWheel` | Known | Event handler |
| 0x000A56F4 | `HandleBufferWheel` | Known | Event handler |
| 0x000A5708 | `HandleBandWheel` | Known | Event handler |
| 0x000A5718 | `HandlePreviousPress` | Known | Event handler |
| 0x000A572C | `HandleNextPress` | Known | Event handler |
| 0x000A573C | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x000A5758 | `HandleNextPressAndHold` | Known | Event handler |
| 0x000A5770 | `HandlePreviousTuning` | Known | Event handler |
| 0x000A5788 | `HandleNextTuning` | Known | Event handler |
| 0x000A579C | `HandlePlayPause` | Known | Event handler |
| 0x000A57AC | `HandleMikeyCenter` | Known | Event handler |
| 0x000A57C0 | `HandleMikeyPrevious` | Known | Event handler |
| 0x000A57D4 | `HandleMikeyNext` | Known | Event handler |
| 0x000A57E4 | `HandleMikeyVolume` | Known | Event handler |
| 0x000A62A0 | `HandlePushToCapacity` | Known | Event handler |
| 0x000A62BC | `HandlePopToCapacity` | Known | Event handler |
| 0x000A62D0 | `HandlePushToCount` | Known | Event handler |
| 0x000A62E4 | `HandlePopToCount` | Known | Event handler |
| 0x000A62F8 | `HandlePushToBasic` | Known | Event handler |
| 0x000A630C | `HandlePopToBasic` | Known | Event handler |
| 0x000A6320 | `HandlePushToAccessoryCapacity` | Known | Event handler |
| 0x000A6340 | `HandlePopToAccessoryCapacity` | Known | Event handler |
| 0x000A6360 | `HandlePushToAccessoryCount` | Known | Event handler |
| 0x000A637C | `HandlePopToAccessoryCount` | Known | Event handler |
| 0x000A6398 | `HandlePushToAccessoryBasic` | Known | Event handler |
| 0x000A63B4 | `HandlePopToAccessoryBasic` | Known | Event handler |
| 0x000A63D0 | `HandlePushToAccessoryAccessory` | Known | Event handler |
| 0x000A63F0 | `HandlePopToAccessoryAccessory` | Known | Event handler |
| 0x000A6F54 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x000AA278 | `HandleSelectCity` | Known | Event handler |
| 0x000AA5DC | `HandleOrientationChange` | Known | Event handler |
| 0x000AA5F8 | `HandleVolumePopup` | Known | Event handler |
| 0x000AADA4 | `HandleLeaveAlarm` | Known | Event handler |
| 0x000AAF40 | `HandleSelect` | Known | Event handler |
| 0x000AB1F0 | `HandleHiliteAlbum` | Known | Event handler |
| 0x000AB208 | `HandleBrowseAlbum` | Known | Event handler |
| 0x000AB21C | `HandlePlayPause` | Known | Event handler |
| 0x000AB22C | `HandlePushEvents` | Known | Event handler |
| 0x000AB240 | `HandlePopEvents` | Known | Event handler |
| 0x000AB250 | `HandlePushFaces` | Known | Event handler |
| 0x000AB260 | `HandlePopFaces` | Known | Event handler |
| 0x000AB270 | `HandlePushPlaces` | Known | Event handler |
| 0x000AB284 | `HandlePopPlaces` | Known | Event handler |
| 0x000ACD2C | `HandleSelect` | Known | Event handler |
| 0x000ACEC8 | `HandleSelectRegion_Africa` | Known | Event handler |
| 0x000ACEE8 | `HandleSelectRegion_Asia` | Known | Event handler |
| 0x000ACF00 | `HandleSelectRegion_Atlantic` | Known | Event handler |
| 0x000ACF1C | `HandleSelectRegion_Australia` | Known | Event handler |
| 0x000ACF3C | `HandleSelectRegion_Europe` | Known | Event handler |
| 0x000ACF58 | `HandleSelectRegion_NorthAmerica` | Known | Event handler |
| 0x000ACF78 | `HandleSelectRegion_Pacific` | Known | Event handler |
| 0x000ACF94 | `HandleSelectRegion_SouthAmerica` | Known | Event handler |
| 0x000AFE6C | `HandleLanguage` | Known | Event handler |
| 0x000AFE80 | `HandleLanguagePop` | Known | Event handler |
| 0x000B0078 | `HandleMainMenu` | Known | Event handler |
| 0x000B0590 | `HandlePlayRadio` | Known | Event handler |
| 0x000B05A4 | `HandleStopRadio` | Known | Event handler |
| 0x000B05B4 | `HandleAutoTune` | Known | Event handler |
| 0x000B05C4 | `HandleTogglePlayPause` | Known | Event handler |
| 0x000B05DC | `HandleToggleBufferSetting` | Known | Event handler |
| 0x000B05F8 | `HandleScanLogging` | Known | Event handler |
| 0x000B1010 | `HandleSelect` | Known | Event handler |
| 0x000B1158 | `HandleSelect` | Known | Event handler |
| 0x000B12AC | `HandleMusicMenu` | Known | Event handler |
| 0x000B2584 | `HandleSelectPreset` | Known | Event handler |
| 0x000B259C | `HandleRemovePreset` | Known | Event handler |
| 0x000B25B0 | `HandleTogglePlayPause` | Known | Event handler |
| 0x000B27D8 | `HandlePrev` | Known | Event handler |
| 0x000B27E8 | `HandleNext` | Known | Event handler |
| 0x000B27F4 | `HandlePlayPause` | Known | Event handler |
| 0x000B2A30 | `HandleNextContact` | Known | Event handler |
| 0x000B2A48 | `HandlePreviousContact` | Known | Event handler |
| 0x000B2C8C | `HandleSelectPressAndHold` | Known | Event handler |
| 0x000B2CAC | `HandleDeleteItem` | Known | Event handler |
| 0x000B2CC0 | `HandleDeleteAllItems` | Known | Event handler |
| 0x000B3658 | `HandleItemSelected` | Known | Event handler |
| 0x000B3768 | `HandleSelect` | Known | Event handler |
| 0x000B38AC | `HandleSelect` | Known | Event handler |
| 0x000B39E4 | `HandleRadioRegion` | Known | Event handler |
| 0x000B4B04 | `HandlePlayPause` | Known | Event handler |
| 0x000B5788 | `HandleCenterButtonSelected` | Known | Event handler |
| 0x000B57A8 | `HandleWheel` | Known | Event handler |
| 0x000B5E1C | `HandleTVOutChanged` | Known | Event handler |
| 0x000B5E34 | `HandleTVSignalChanged` | Known | Event handler |
| 0x000B5E4C | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x000B5E6C | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x000B5E8C | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x000B5EB0 | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x000B5ED0 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x000B6564 | `HandleSelect` | Known | Event handler |
| 0x000B6EAC | `HandleLeaveAlarm` | Known | Event handler |
| 0x000B7DD8 | `HandleItemSelected` | Known | Event handler |
| 0x000B837C | `HandleResetAllSettings` | Known | Event handler |
| 0x000B8394 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x000B8B90 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x000B8F94 | `HandleSelect` | Known | Event handler |
| 0x000BA15C | `HandleStartGenius` | Known | Event handler |
| 0x000BA540 | `HandleNextDay` | Known | Event handler |
| 0x000BA554 | `HandlePreviousDay` | Known | Event handler |
| 0x000BA81C | `HandleWheel` | Known | Event handler |
| 0x000BA9E0 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x000BAAE4 | `HandleNextDay` | Known | Event handler |
| 0x000BAAF8 | `HandlePreviousDay` | Known | Event handler |
| 0x000BABF4 | `HandleSelect` | Known | Event handler |
| 0x000BB92C | `HandleDeleteClock` | Known | Event handler |
| 0x000BB944 | `HandleSelectClock` | Known | Event handler |
| 0x000BB958 | `HandleHilited` | Known | Event handler |
| 0x000BB968 | `HandleWheel` | Known | Event handler |
| 0x000BB974 | `HandleSelectLozinch` | Known | Event handler |
| 0x000BB988 | `HandleAddClock` | Known | Event handler |
| 0x000BB998 | `HandleEditClock` | Known | Event handler |
| 0x000C1E30 | `HandleWheel` | Known | Event handler |
| 0x000C1E3C | `HandlePlayPause` | Known | Event handler |
| 0x000C1E4C | `HandleSelectDown` | Known | Event handler |
| 0x000C1E60 | `HandleNext` | Known | Event handler |
| 0x000C1E6C | `HandlePrevious` | Known | Event handler |
| 0x000C1E7C | `HandleNextPushAndHold` | Known | Event handler |
| 0x000C1E94 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000C5BF0 | `HandlePortEvents` | Known | Event handler |
| 0x000C63A8 | `HandleWantPopFlow` | Known | Event handler |
| 0x000C63C0 | `HandleSwapToNowPlayingFromOrientation` | Known | Event handler |
| 0x000C63E8 | `HandleSwapToNowPlaying` | Known | Event handler |
| 0x000C6400 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x000C641C | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x000C6438 | `HandleFlowNext` | Known | Event handler |
| 0x000C6448 | `HandleFlowPrev` | Known | Event handler |
| 0x000C6458 | `HandleFlowWheel` | Known | Event handler |
| 0x000C6468 | `HandleAlbumSelected` | Known | Event handler |
| 0x000C647C | `HandlePlayPause` | Known | Event handler |
| 0x000C648C | `HandleBacksideSongSelected` | Known | Event handler |
| 0x000C64A8 | `HandleScreenRotation` | Known | Event handler |
| 0x000C64C0 | `HandleNext` | Known | Event handler |
| 0x000C64CC | `HandleNextPressAndHold` | Known | Event handler |
| 0x000C64E4 | `HandlePrevious` | Known | Event handler |
| 0x000C64F4 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x000E1768 | `HandleSelect` | Known | Event handler |
| 0x000E177C | `HandleGameHilited` | Known | Event handler |
| 0x000E22A0 | `HandleLock` | Known | Event handler |
| 0x000E22B0 | `HandleGotoAddressBookScreen` | Known | Event handler |
| 0x000E22CC | `HandleGotoCalendarLoadingScreen` | Known | Event handler |
| 0x000E22EC | `HandleNikePlusSelected` | Known | Event handler |
| 0x000E2304 | `HandleVoiceMemosSelected` | Known | Event handler |
| 0x000E2320 | `HandleRadioSelected` | Known | Event handler |
| 0x000E2334 | `HandleRadioPlayPause` | Known | Event handler |
| 0x000E2634 | `HandleOrientationPortrait` | Known | Event handler |
| 0x000E2654 | `HandleOrientationLandscape` | Known | Event handler |
| 0x000E2670 | `HandleScreenRotation` | Known | Event handler |
| 0x000E2688 | `HandleGestureShake` | Known | Event handler |
| 0x000E269C | `HandleGestureSteer` | Known | Event handler |
| 0x000E2A3C | `HandlePlayPause` | Known | Event handler |
| 0x000E2A50 | `HandleAddChapterMark` | Known | Event handler |
| 0x000E3CEC | `HandleExitUnsupported` | Known | Event handler |
| 0x000E48B8 | `HandleScreenRotation` | Known | Event handler |
| 0x000E4F24 | `HandleNext` | Known | Event handler |
| 0x000E4F34 | `HandlePrev` | Known | Event handler |
| 0x000E4F40 | `HandleNextPressAndHold` | Known | Event handler |
| 0x000E4F58 | `HandlePrevPressAndHold` | Known | Event handler |
| 0x000E4F70 | `HandleWheel` | Known | Event handler |
| 0x000E4F7C | `HandleOrientationAlt` | Known | Event handler |
| 0x000E4F94 | `HandleOrientationDefault` | Known | Event handler |
| 0x000E4FB0 | `HandleScreenRotation` | Known | Event handler |
| 0x000E4FC8 | `HandlePlayPause` | Known | Event handler |
| 0x000E4FD8 | `HandlePlay` | Known | Event handler |
| 0x000E4FE4 | `HandlePause` | Known | Event handler |
| 0x000E4FF0 | `HandleMikeyPlayPause` | Known | Event handler |
| 0x000E5008 | `HandleSelect` | Known | Event handler |
| 0x000E5018 | `HandleMenuUp` | Known | Event handler |
| 0x000E7200 | `HandleWheel` | Known | Event handler |
| 0x000E7210 | `HandleArrowUp` | Known | Event handler |
| 0x000E7220 | `HandleArrowDown` | Known | Event handler |
| 0x000EAB24 | `HandleSteer` | Known | Event handler |
| 0x000EADD4 | `HandleSelect` | Known | Event handler |
| 0x000EB058 | `HandleShowRecordings` | Known | Event handler |
| 0x000EB07C | `HandleDeleteAllSelect` | Known | Event handler |
| 0x000EB094 | `HandleDeleteSelect` | Known | Event handler |
| 0x000EB0A8 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x000EB0C8 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x000EB0EC | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x000EB108 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x000EB128 | `HandleMicrophoneRequired` | Known | Event handler |
| 0x000EB144 | `HandleMicrophoneDisconnected` | Known | Event handler |
| 0x000EB56C | `HandleFrequencyChosen` | Known | Event handler |
| 0x000EB584 | `HandleDateChosen` | Known | Event handler |
| 0x000EB598 | `HandleTimeChosen` | Known | Event handler |
| 0x000EB5AC | `HandleSoundChosen` | Known | Event handler |
| 0x000EB5C0 | `HandleLabelChosen` | Known | Event handler |
| 0x000EB5D4 | `HandleDeleteChosen` | Known | Event handler |
| 0x000EB6C0 | `HandleSelect` | Known | Event handler |
| 0x000EBB74 | `HandleOrientationAlt` | Known | Event handler |
| 0x000EBF34 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x000EC314 | `HandleSelectPressAndHold` | Known | Event handler |
| 0x000EC334 | `HandleDeleteItem` | Known | Event handler |
| 0x000EC348 | `HandleDeleteAllItems` | Known | Event handler |
| 0x000ED448 | `HandleSelectLabel` | Known | Event handler |
| 0x000EE314 | `HandleStartGenius` | Known | Event handler |
| 0x000EE32C | `HandleViewArtist` | Known | Event handler |
| 0x000EE340 | `HandleViewAlbum` | Known | Event handler |
| 0x000EE350 | `HandleViewCompilation` | Known | Event handler |
| 0x000EE368 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x000EEAC8 | `HandleStartGenius` | Known | Event handler |
| 0x000EEADC | `HandleAddToOTG` | Known | Event handler |
| 0x000EEAF0 | `HandleViewCompilation` | Known | Event handler |
| 0x000EEB08 | `HandleViewAlbum` | Known | Event handler |
| 0x000EEB18 | `HandleViewArtist` | Known | Event handler |
| 0x000EEB2C | `HandleCancel` | Known | Event handler |
| 0x000EF354 | `HandleAddToOTG` | Known | Event handler |
| 0x000EF364 | `HandleCancel` | Known | Event handler |
| 0x000EF4F0 | `HandleStartGenius` | Known | Event handler |
| 0x000EF508 | `HandleViewAlbum` | Known | Event handler |
| 0x000EF518 | `HandleViewArtist` | Known | Event handler |
| 0x000EF52C | `HandleViewCompilation` | Known | Event handler |
| 0x000EF544 | `HandleSaveGeniusPlaylist` | Known | Event handler |
| 0x000EF560 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x000EF578 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x000F0094 | `HandleStartGenius` | Known | Event handler |
| 0x000F00A8 | `HandleAddToOTG` | Known | Event handler |
| 0x000F00BC | `HandleViewCompilation` | Known | Event handler |
| 0x000F00D4 | `HandleViewAlbum` | Known | Event handler |
| 0x000F00E4 | `HandleViewArtist` | Known | Event handler |
| 0x000F00F8 | `HandleCancel` | Known | Event handler |
| 0x000F0380 | `HandleAddToOTG` | Known | Event handler |
| 0x000F0390 | `HandleCancel` | Known | Event handler |
| 0x000F0C5C | `HandleAddToOTG` | Known | Event handler |
| 0x000F0C6C | `HandleCancel` | Known | Event handler |
| 0x000F1318 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x000F1694 | `HandleAddToOTG` | Known | Event handler |
| 0x000F16A4 | `HandleCancel` | Known | Event handler |
| 0x000F1908 | `HandleConfirmation` | Known | Event handler |
| 0x000F1DF0 | `HandleMusicHilited` | Known | Event handler |
| 0x000F1E08 | `HandleVideosHilited` | Known | Event handler |
| 0x000F1E1C | `HandlePodcastsHilited` | Known | Event handler |
| 0x000F1E34 | `HandleiTunesUHilited` | Known | Event handler |
| 0x000F1E4C | `HandleGenericHilited` | Known | Event handler |
| 0x000F1E64 | `HandlePhotosHilited` | Known | Event handler |
| 0x000F1E78 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x000F1E90 | `HandleExtrasHilited` | Known | Event handler |
| 0x000F1EA4 | `HandleSettingsHilited` | Known | Event handler |
| 0x000F1EBC | `HandleCameraHilited` | Known | Event handler |
| 0x000F1ED0 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x000F1EEC | `HandleAudiobooksHilited` | Known | Event handler |
| 0x000F1F04 | `HandleArtistsHilited` | Known | Event handler |
| 0x000F1F1C | `HandleGenresHilited` | Known | Event handler |
| 0x000F1F30 | `HandleAlbumsHilited` | Known | Event handler |
| 0x000F1F44 | `HandleCompilationsHilited` | Known | Event handler |
| 0x000F1F60 | `HandleComposersHilited` | Known | Event handler |
| 0x000F1F78 | `HandleSongsHilited` | Known | Event handler |
| 0x000F1F8C | `HandlePlaylistsHilited` | Known | Event handler |
| 0x000F1FA4 | `HandleGeniusMixesHilited` | Known | Event handler |
| 0x000F1FC0 | `HandleCoverflowHilited` | Known | Event handler |
| 0x000F1FD8 | `HandleTVShowsHilited` | Known | Event handler |
| 0x000F1FF0 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x000F200C | `HandleMoviesHilited` | Known | Event handler |
| 0x000F2020 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x000F203C | `HandleRentalsHilited` | Known | Event handler |
| 0x000F2054 | `HandleRadioHilited` | Known | Event handler |
| 0x000F2068 | `HandleVoiceMemosHilited` | Known | Event handler |
| 0x000F2080 | `HandleVoiceMemosSelected` | Known | Event handler |
| 0x000F209C | `HandleMusicSelected` | Known | Event handler |
| 0x000F20B0 | `HandleVideosSelected` | Known | Event handler |
| 0x000F20C8 | `HandlePodcastsSelected` | Known | Event handler |
| 0x000F20E0 | `HandleiTunesUSelected` | Known | Event handler |
| 0x000F20F8 | `HandlePhotosSelected` | Known | Event handler |
| 0x000F2110 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x000F2128 | `HandleSongsSelected` | Known | Event handler |
| 0x000F213C | `HandleAlbumsSelected` | Known | Event handler |
| 0x000F2154 | `HandleCompilationsSelected` | Known | Event handler |
| 0x000F2170 | `HandleArtistsSelected` | Known | Event handler |
| 0x000F2400 | `HandleGenresSelected` | Known | Event handler |
| 0x000F2418 | `HandleComposersSelected` | Known | Event handler |
| 0x000F2430 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x000F244C | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x000F2468 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x000F2480 | `HandleNowPlaying` | Known | Event handler |
| 0x000F2494 | `HandleCameraSelected` | Known | Event handler |
| 0x000F24AC | `HandleGeniusMixesSelected` | Known | Event handler |
| 0x000F24C8 | `HandleTVShowsSelected` | Known | Event handler |
| 0x000F24E0 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x000F24FC | `HandleMoviesSelected` | Known | Event handler |
| 0x000F2514 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x000F2534 | `HandleRentalsSelected` | Known | Event handler |
| 0x000F254C | `HandleCameraVideosSelected` | Known | Event handler |
| 0x000F2568 | `HandleAddressBookSelected` | Known | Event handler |
| 0x000F2584 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x000F259C | `HandleSleepSelected` | Known | Event handler |
| 0x000F25B0 | `HandleNikePlusSelected` | Known | Event handler |
| 0x000F25C8 | `HandleNikePlusHilited` | Known | Event handler |
| 0x000F25E0 | `HandleRadioSelected` | Known | Event handler |
| 0x000F25F4 | `HandleRadioPreviewPlayPause` | Known | Event handler |
| 0x000F2610 | `HandlePedometerSelected` | Known | Event handler |
| 0x000F2628 | `HandlePedometerHilited` | Known | Event handler |
| 0x000F2640 | `HandlePortraitToLandscape` | Known | Event handler |
| 0x000F2BC0 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x000F2E70 | `HandleAddToOTG` | Known | Event handler |
| 0x000F2E80 | `HandleCancel` | Known | Event handler |
| 0x000F38C8 | `HandleAddToOTG` | Known | Event handler |
| 0x000F38D8 | `HandleCancel` | Known | Event handler |
| 0x000F3F74 | `HandleAddToOTG` | Known | Event handler |
| 0x000F3F84 | `HandleCancel` | Known | Event handler |
| 0x000F4174 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x000F4668 | `HandleGeniusMixPlaylistReady` | Known | Event handler |
| 0x000F468C | `HandleSelectMix` | Known | Event handler |
| 0x000F469C | `HandlePrev` | Known | Event handler |
| 0x000F46A8 | `HandleNext` | Known | Event handler |
| 0x000F46B4 | `HandlePlayPause` | Known | Event handler |
| 0x000F46C4 | `HandleWheel` | Known | Event handler |
| 0x000F4FBC | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x000F4FD8 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x000F4FF0 | `HandleStartGenius` | Known | Event handler |
| 0x000F5004 | `HandleViewArtist` | Known | Event handler |
| 0x000F5018 | `HandleViewAlbum` | Known | Event handler |
| 0x000F5028 | `HandleViewCompilation` | Known | Event handler |
| 0x000F5040 | `HandleShowContextualMenu` | Known | Event handler |
| 0x000F505C | `HandleRefreshPlaylist` | Known | Event handler |
| 0x000F5074 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x000F5D04 | `HandleStartGenius` | Known | Event handler |
| 0x000F5D18 | `HandleAddToOTG` | Known | Event handler |
| 0x000F5D2C | `HandleViewCompilation` | Known | Event handler |
| 0x000F5D44 | `HandleViewAlbum` | Known | Event handler |
| 0x000F5D54 | `HandleViewArtist` | Known | Event handler |
| 0x000F5D68 | `HandleCancel` | Known | Event handler |
| 0x000F5DBC | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x000F5F1C | `HandleAddToOTG` | Known | Event handler |
| 0x000F5F2C | `HandleCancel` | Known | Event handler |
| 0x000F6EFC | `HandleSelect` | Known | Event handler |
| 0x000F6F0C | `HandleCameraRemoteSelect` | Known | Event handler |
| 0x000F6F28 | `HandleMikeyCenter` | Known | Event handler |
| 0x000F6F3C | `HandleMikeyCenterReleased` | Known | Event handler |
| 0x000F6F74 | `HandleSelectPressAndHold` | Known | Event handler |
| 0x000F76DC | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x000F76FC | `HandleMikeyCenterReleased` | Known | Event handler |
| 0x000F7718 | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x000F7738 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x000F775C | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x000F777C | `HandleMikeyAllUp` | Known | Event handler |
| 0x000F7790 | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x000F77A4 | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x000F77BC | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x000F77D4 | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x000F77EC | `HandleSelect` | Known | Event handler |
| 0x000F77FC | `HandleCameraRemoteSelect` | Known | Event handler |
| 0x000F7818 | `HandleMikeyCenter` | Known | Event handler |
| 0x000F7974 | `HandleSelect` | Known | Event handler |
| 0x000F7988 | `HandleNext` | Known | Event handler |
| 0x000F7994 | `HandlePrev` | Known | Event handler |
| 0x000F79A0 | `HandleWheel` | Known | Event handler |
| 0x000F7A60 | `HandleAppInitialized` | Known | Event handler |
| 0x000FC7A4 | `HandleSelect` | Known | Event handler |
| 0x00102E10 | `HandleRotationChange` | Known | Event handler |
| 0x00102E2C | `HandleSteerGesture` | Known | Event handler |
| 0x00102E40 | `HandlePlayPause` | Known | Event handler |
| 0x00102E50 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x00102E6C | `HandleTouchAndHoldPlayPause` | Known | Event handler |
| 0x00102E88 | `HandleNext` | Known | Event handler |
| 0x00102E94 | `HandleNextPressAndHold` | Known | Event handler |
| 0x00102EAC | `HandlePrevious` | Known | Event handler |
| 0x00102EBC | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x00102ED8 | `HandleRemotePlayPause` | Known | Event handler |
| 0x00102EF0 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x00102F14 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x00102F2C | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x00102F44 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x00102F5C | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x00102F74 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x00102F8C | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x00102FA8 | `HandleRemoteStop` | Known | Event handler |
| 0x00102FBC | `HandleRemotePlay` | Known | Event handler |
| 0x00102FD0 | `HandleRemotePause` | Known | Event handler |
| 0x00102FE4 | `HandleRemoteMute` | Known | Event handler |
| 0x00102FF8 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x00103010 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x00103028 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x00103044 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x00103060 | `HandleRemoteShuffle` | Known | Event handler |
| 0x00103074 | `HandleRemoteRepeat` | Known | Event handler |
| 0x00103088 | `HandleRemoteOn` | Known | Event handler |
| 0x00103098 | `HandleRemoteOff` | Known | Event handler |
| 0x001030A8 | `HandleRemoteBacklight` | Known | Event handler |
| 0x001030C0 | `HandleRemoteFFDown` | Known | Event handler |
| 0x001030D4 | `HandleRemoteFFUp` | Known | Event handler |
| 0x001030E8 | `HandleRemoteRewDown` | Known | Event handler |
| 0x001030FC | `HandleRemoteRewUp` | Known | Event handler |
| 0x00103110 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x00103128 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x0010313C | `HandleRemoteSelectDown` | Known | Event handler |
| 0x00103154 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x0010316C | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x00103184 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x0010319C | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x00103578 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x00103590 | `HandleRemoteEvent` | Known | Event handler |
| 0x001035A4 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x001035C0 | `HandleAudioPlayPause` | Known | Event handler |
| 0x001035D8 | `HandleAudioNext` | Known | Event handler |
| 0x001035E8 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x00103604 | `HandleAudioPrevious` | Known | Event handler |
| 0x00103618 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x00103638 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x00103650 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x00103668 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x00103680 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x00103694 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x001036AC | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x001036C4 | `HandleAudioStop` | Known | Event handler |
| 0x001036D4 | `HandleAudioPlay` | Known | Event handler |
| 0x001036E4 | `HandleAudioPause` | Known | Event handler |
| 0x001036F8 | `HandleAudioMute` | Known | Event handler |
| 0x00103708 | `HandleAudioNextChapter` | Known | Event handler |
| 0x00103720 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x00103738 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x00103750 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x00103768 | `HandleAudioShuffle` | Known | Event handler |
| 0x0010377C | `HandleAudioRepeat` | Known | Event handler |
| 0x00103790 | `HandleAudioFFDown` | Known | Event handler |
| 0x001037A4 | `HandleAudioFFUp` | Known | Event handler |
| 0x001037B4 | `HandleAudioRewDown` | Known | Event handler |
| 0x001037C8 | `HandleAudioRewUp` | Known | Event handler |
| 0x001037DC | `HandleVideoPlayPause` | Known | Event handler |
| 0x001037F4 | `HandleVideoNext` | Known | Event handler |
| 0x00103804 | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x00103820 | `HandleVideoPrevious` | Known | Event handler |
| 0x00103834 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x00103854 | `HandleVideoStop` | Known | Event handler |
| 0x00103864 | `HandleVideoPlay` | Known | Event handler |
| 0x00103874 | `HandleVideoPause` | Known | Event handler |
| 0x00103888 | `HandleVideoFFDown` | Known | Event handler |
| 0x0010389C | `HandleVideoFFUp` | Known | Event handler |
| 0x001038AC | `HandleVideoRewDown` | Known | Event handler |
| 0x001038C0 | `HandleVideoRewUp` | Known | Event handler |
| 0x001038D4 | `HandleVideoNextChapter` | Known | Event handler |
| 0x001038EC | `HandleVideoPrevChapter` | Known | Event handler |
| 0x00103BB0 | `HandleVideoNextFrame` | Known | Event handler |
| 0x00103BC8 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x00103BE0 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x00103BFC | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00103C18 | `HandleShakeShuffleSongsSelected` | Known | Event handler |
| 0x00103C38 | `HandleGlobalVolume` | Known | Event handler |
| 0x00103C4C | `HandleSimulateOrientationChange` | Known | Event handler |
| 0x00103C6C | `HandleMikeyCenter` | Known | Event handler |
| 0x00103C80 | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x00103CA0 | `HandleMikeyCenterReleased` | Known | Event handler |
| 0x00103CBC | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x00103CDC | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x00103D00 | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x00103D20 | `HandleMikeyCenterTripleClickAndHold` | Known | Event handler |
| 0x00103D44 | `HandleMikeyAllUp` | Known | Event handler |
| 0x00103D58 | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x00103D6C | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x00103D84 | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x00103D9C | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x00103DB4 | `HandleVoiceOverPlaylistSelected` | Known | Event handler |
| 0x00103DD4 | `HandleVoiceOverPodcastSelected` | Known | Event handler |
| 0x00103DF4 | `HandleVoiceOverAudiobookSelected` | Known | Event handler |
| 0x00103E18 | `HandleRadioTagUp` | Known | Event handler |
| 0x00103E2C | `HandleVoiceCommand` | Known | Event handler |
| 0x00103E40 | `HandleVoiceArtist` | Known | Event handler |
| 0x00103E54 | `HandleVoiceAlbum` | Known | Event handler |
| 0x00103E68 | `HandleNECIRMenuUp` | Known | Event handler |
| 0x00103E7C | `HandleNECIRPlayPauseUp` | Known | Event handler |
| 0x00103E94 | `HandleNECIRNextUp` | Known | Event handler |
| 0x00103EA8 | `HandleNECIRPrevUp` | Known | Event handler |
| 0x00103EBC | `HandleNECIRVolumeDownUp` | Known | Event handler |
| 0x00103ED4 | `HandleNECIRVolumeUpUp` | Known | Event handler |
| 0x00103EEC | `HandleCameraRemote` | Known | Event handler |
| 0x00107FB4 | `HandleLoadingCancelled` | Known | Event handler |
| 0x001B228C | `HandleWheelVolume` | Known | Event handler |
| 0x001B22A4 | `HandleMenuKey` | Known | Event handler |
| 0x001B22B4 | `HandlePauseKey` | Known | Event handler |
| 0x001B22C4 | `HandlePrevKey` | Known | Event handler |
| 0x001B22D4 | `HandleNextKey` | Known | Event handler |
| 0x001B2B48 | `HandleSelect` | Known | Event handler |
| 0x001B2FE4 | `HandleChooseLink` | Known | Event handler |
| 0x001B2FFC | `HandleUnlink` | Known | Event handler |
| 0x001B3124 | `HandleSelectedDayWorkout` | Known | Event handler |
| 0x001B3144 | `HandleMenuUp` | Known | Event handler |
| 0x001B337C | `HandleSelect` | Known | Event handler |
| 0x001B3390 | `HandleMenu` | Known | Event handler |
| 0x001B339C | `HandleLinkCancelOption` | Known | Event handler |
| 0x001B33B4 | `HandleLinkNewRemote` | Known | Event handler |
| 0x001B33C8 | `HandleLinkNewHeartMonitor` | Known | Event handler |
| 0x001B33E4 | `HandleCancelRemoteLinking` | Known | Event handler |
| 0x001B3664 | `HandleChoosePlaylistsPlay` | Known | Event handler |
| 0x001B3684 | `HandleChooseGeniusMixesPlay` | Known | Event handler |
| 0x001B36A0 | `HandleChoosePodcastsPlay` | Known | Event handler |
| 0x001B36BC | `HandleChooseAudiobooksPlay` | Known | Event handler |
| 0x001B36D8 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x001B36F4 | `HandleNoneSelected` | Known | Event handler |
| 0x001B3708 | `HandleNowPlayingSelected` | Known | Event handler |
| 0x001B3724 | `HandleMenuUp` | Known | Event handler |
| 0x001B3734 | `HandleMenuKeyNop` | Known | Event handler |
| 0x001B41F8 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x001B4570 | `HandleChoosePowerPlay` | Known | Event handler |
| 0x001B458C | `HandleChooseUnit` | Known | Event handler |
| 0x001B45A0 | `HandleChooseVoicePrompts` | Known | Event handler |
| 0x001B45BC | `HandleChoosePedometer` | Known | Event handler |
| 0x001B5170 | `HandlePopupDismissed` | Known | Event handler |
| 0x001B5234 | `HandleSelect` | Known | Event handler |
| 0x001B538C | `HandleListChoose` | Known | Event handler |
| 0x001B5904 | `HandleBasicSelected` | Known | Event handler |
| 0x001B591C | `HandleTimedWorkoutSelected` | Known | Event handler |
| 0x001B5938 | `HandleDistanceWorkoutSelected` | Known | Event handler |
| 0x001B5958 | `HandleCaloriesWorkoutSelected` | Known | Event handler |
| 0x001B606C | `HandleWalkCalibrationSelection` | Known | Event handler |
| 0x001B6090 | `HandleRunCalibrationSelection` | Known | Event handler |
| 0x001B6240 | `HandleNewWorkout` | Known | Event handler |
| 0x001B6258 | `HandleNewBasicWorkout` | Known | Event handler |
| 0x001B6270 | `HandleNewQuickstartWorkout` | Known | Event handler |
| 0x001B628C | `HandleResumeWorkout` | Known | Event handler |
| 0x001B64A4 | `HandleChooseTimedWorkoutFromList` | Known | Event handler |
| 0x001B64CC | `HandleChooseDistanceWorkoutFromList` | Known | Event handler |
| 0x001B64F0 | `HandleChooseCalorieWorkoutFromList` | Known | Event handler |
| 0x001B65D0 | `HandleSelect` | Known | Event handler |
| 0x001B65E4 | `HandleGeniusMixPlaylistReady` | Known | Event handler |
| 0x001B6848 | `HandleVerticalSelected` | Known | Event handler |
| 0x001B6864 | `HandleRightSelected` | Known | Event handler |
| 0x001B6878 | `HandleLeftSelected` | Known | Event handler |
| 0x001B7054 | `HandleBegin` | Known | Event handler |
| 0x001B7618 | `HandleItemSelected` | Known | Event handler |
| 0x001B7760 | `HandleSelect` | Known | Event handler |
| 0x001B7774 | `HandlePopBackToSongsScreen` | Known | Event handler |
| 0x001B7A44 | `HandleUnlinkRemote` | Known | Event handler |
| 0x001B7B54 | `HandleChooseRemoteSetting` | Known | Event handler |
| 0x001B7B74 | `HandleChooseRemoteLink` | Known | Event handler |
| 0x001B7B8C | `HandleChooseRemoteUnlink` | Known | Event handler |
| 0x001B7DC4 | `HandleDeleteWorkout` | Known | Event handler |
| 0x001B7DDC | `HandleDeleteAllWorkouts` | Known | Event handler |
| 0x001B7DF4 | `HandleOrientationChange` | Known | Event handler |
| 0x001B7E0C | `HandleSelectNextWorkout` | Known | Event handler |
| 0x001B7E24 | `HandleSelectPrevWorkout` | Known | Event handler |
| 0x001B82BC | `HandleDeleteAllWorkouts` | Known | Event handler |
| 0x001B82D8 | `HandleClearBests` | Known | Event handler |
| 0x001B82EC | `HandleClearTotals` | Known | Event handler |
| 0x001B8300 | `HandleResetWalkingCalibration` | Known | Event handler |
| 0x001B8320 | `HandleResetRuningCalibration` | Known | Event handler |
| 0x001B845C | `HandlePopSelf` | Known | Event handler |
| 0x001B8470 | `HandlePressAndHold` | Known | Event handler |
| 0x001B8908 | `HandleHerculesKey` | Known | Event handler |
| 0x001B8A1C | `HandleUnlinkHeartMonitor` | Known | Event handler |
| 0x001B8B20 | `HandleChooseHeartMonitorLink` | Known | Event handler |
| 0x001B8B44 | `HandleChooseHeartMonitorUnlink` | Known | Event handler |
| 0x001BC0AC | `HandleAddToOTG` | Known | Event handler |
| 0x001BC0BC | `HandleCancel` | Known | Event handler |
| 0x001BC3F8 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x001BC498 | `HandleAddToOTG` | Known | Event handler |
| 0x001BC4A8 | `HandleCancel` | Known | Event handler |
| 0x001BC4F0 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x001C8F88 | `HandleOrientationChange` | Known | Event handler |
| 0x001C8FA4 | `HandleNext` | Known | Event handler |
| 0x001C8FB0 | `HandlePrev` | Known | Event handler |
| 0x001C8FBC | `HandleWheel` | Known | Event handler |
| 0x001C8FC8 | `HandleSelect` | Known | Event handler |
| 0x001C91BC | `HandleSelectedNikeMainMenuItem` | Known | Event handler |
| 0x001C91E0 | `HandleGeniusMixPlaylistReady` | Known | Event handler |
| 0x001C9538 | `HandleSelect` | Known | Event handler |
| 0x001C954C | `HandleMenuKey` | Known | Event handler |
| 0x001C955C | `HandlePauseWorkout` | Known | Event handler |
| 0x001C9570 | `HandleEndWorkout` | Known | Event handler |
| 0x001C9584 | `HandleResumeWorkout` | Known | Event handler |
| 0x001C9598 | `HandleChooseMusic` | Known | Event handler |
| 0x001C9C14 | `HandleWeightWheel` | Known | Event handler |
| 0x001C9C2C | `HandleWeightSelect` | Known | Event handler |
| 0x001C9C40 | `HandleWeightSelectAltTrans` | Known | Event handler |
| 0x001C9C5C | `HandleDistanceWheel` | Known | Event handler |
| 0x001C9C70 | `HandleDistanceSelect` | Known | Event handler |
| 0x001C9C88 | `HandleTimeWheel` | Known | Event handler |
| 0x001C9C98 | `HandleTimeSelect` | Known | Event handler |
| 0x001C9CAC | `HandleCaloriesWheel` | Known | Event handler |
| 0x001C9CC0 | `HandleCaloriesSelect` | Known | Event handler |
| 0x001C9CD8 | `HandleStepGoalWheel` | Known | Event handler |
| 0x001C9CEC | `HandleStepGoalSelect` | Known | Event handler |
| 0x001C9D04 | `HandleWeightSelectPedometer` | Known | Event handler |
| 0x001C9F08 | `HandleNikeNestedPlaylistSelect` | Known | Event handler |
| 0x001C9F9C | `HandleDistanceWheel` | Known | Event handler |
| 0x001C9FB4 | `HandleDistanceSelect` | Known | Event handler |
| 0x001D5480 | `HandlePauseKey` | Known | Event handler |
| 0x001D5494 | `HandlePauseHold` | Known | Event handler |
| 0x001D54A4 | `HandleMenuKey` | Known | Event handler |
| 0x001D54B4 | `HandleMenuKeyNop` | Known | Event handler |
| 0x001D54C8 | `HandleWheel` | Known | Event handler |
| 0x001D5518 | `HandleSelectKeyDown` | Known | Event handler |
| 0x001D552C | `HandleOnDemandPrompt` | Known | Event handler |
| 0x001D5544 | `HandlePowerPlay` | Known | Event handler |
| 0x001D5554 | `HandlePauseWorkout` | Known | Event handler |
| 0x001D5568 | `HandleEndWorkout` | Known | Event handler |
| 0x001D557C | `HandleResumeWorkout` | Known | Event handler |
| 0x001D5590 | `HandleChooseMusic` | Known | Event handler |
| 0x001D55A4 | `HandleMikeyPressExtended` | Known | Event handler |
| 0x001D55C0 | `Handle3BitModeFinished` | Known | Event handler |
| 0x002127B4 | `HandlePlayPause` | Known | Event handler |
| 0x002127C8 | `HandleWheel` | Known | Event handler |
| 0x002127D4 | `HandleMTWheel` | Known | Event handler |
| 0x002127E4 | `HandleWheelRating` | Known | Event handler |
| 0x002127F8 | `HandleWheelScrub` | Known | Event handler |
| 0x0021280C | `HandleWheelVolume` | Known | Event handler |
| 0x002128CC | `HandleMenuKey` | Known | Event handler |
| 0x002128DC | `HandleMenuLongpress` | Known | Event handler |
| 0x002128F0 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x0021290C | `HandleSwapToCoverflow` | Known | Event handler |
| 0x00212924 | `HandleDefaultOrientation` | Known | Event handler |
| 0x0021613C | `HandleSelect` | Known | Event handler |
| 0x002401F0 | `HandleMikeyCenter` | Known | Event handler |
| 0x00240204 | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x00240224 | `HandleMikeyCenterReleased` | Known | Event handler |
| 0x00240240 | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x00240260 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x00240284 | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x002402A4 | `HandleMikeyAllUp` | Known | Event handler |
| 0x002402B8 | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x002402CC | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x002402E4 | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x002402FC | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x00240314 | `HandleNext` | Known | Event handler |
| 0x00240320 | `HandlePrev` | Known | Event handler |
| 0x0024032C | `HandleNextAndHold` | Known | Event handler |
| 0x00240340 | `HandlePrevAndHold` | Known | Event handler |
| 0x00240354 | `HandleScreenRotation` | Known | Event handler |
| 0x0024036C | `HandleAudioPlayPause` | Known | Event handler |
| 0x00240384 | `HandleCameraRemoteSelect` | Known | Event handler |
| 0x002949FC | `HandleSelect` | Known | Event handler |
| 0x00294A10 | `HandleHilite` | Known | Event handler |
| 0x00294A20 | `HandlePlayPause` | Known | Event handler |
| 0x00294A30 | `HandleAddToOTG` | Known | Event handler |
| 0x00294A40 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00294A60 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00294A7C | `HandleStartQuickNav` | Known | Event handler |
| 0x004641B4 | `HandlePauseKey` | Known | Event handler |
| 0x004641C4 | `HandleMenuKey` | Known | Event handler |
| 0x0056D722 | `HandleMikeyCenterReleased` | Known | Event handler |
| 0x0056D73C | `HandleAppInitialized` | Known | Event handler |
| 0x0056D782 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x0056D7A6 | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x0056D7C4 | `HandleSelectPressAndHold` | Known | Event handler |
| 0x0056D7DD | `HandleNextAndHold` | Known | Event handler |
| 0x0056D7EF | `HandlePrevAndHold` | Known | Event handler |
| 0x0056D801 | `HandleAudioPlayPause` | Known | Event handler |
| 0x0056D83C | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x0056D859 | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x0056D876 | `HandleWheel` | Known | Event handler |
| 0x0056D8AA | `HandleScreenRotation` | Known | Event handler |
| 0x0056D8BF | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x0056D8D5 | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x0056D8E9 | `HandleMikeyAllUp` | Known | Event handler |
| 0x0056D8FA | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x0056D912 | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x0056D928 | `HandleMikeyCenter` | Known | Event handler |
| 0x0056D93A | `HandleSelect` | Known | Event handler |
| 0x0056D947 | `HandleCameraRemoteSelect` | Known | Event handler |
| 0x0056D960 | `HandleNext` | Known | Event handler |
| 0x0056D979 | `HandlePrev` | Known | Event handler |
| 0x00633488 | `HandleShakeShuffleSongsSelected` | Known | Event handler |
| 0x006334B8 | `HandleAudioFFDown` | Known | Event handler |
| 0x006334D8 | `HandleAudioFFUp` | Known | Event handler |
| 0x006334F8 | `HandleAudioMute` | Known | Event handler |
| 0x00633520 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x0063354C | `HandleAudioNext` | Known | Event handler |
| 0x00633574 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x006335A4 | `HandleAudioNextChapter` | Known | Event handler |
| 0x006335D4 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x00633600 | `HandleAudioPause` | Known | Event handler |
| 0x00633624 | `HandleAudioPlay` | Known | Event handler |
| 0x00633648 | `HandleAudioPlayPause` | Known | Event handler |
| 0x00633678 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x006336A8 | `HandleAudioPrevious` | Known | Event handler |
| 0x006336D4 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x00633704 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x00633734 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x00633760 | `HandleAudioRepeat` | Known | Event handler |
| 0x00633784 | `HandleAudioRewDown` | Known | Event handler |
| 0x006337A8 | `HandleAudioRewUp` | Known | Event handler |
| 0x006337D0 | `HandleAudioShuffle` | Known | Event handler |
| 0x006337F4 | `HandleAudioStop` | Known | Event handler |
| 0x0063381C | `HandleAudioVolumeDown` | Known | Event handler |
| 0x00633848 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x00633874 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x0063389C | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x00633918 | `HandleNextPressAndHold` | Known | Event handler |
| 0x00633930 | `HandleNext` | Known | Event handler |
| 0x00633960 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x0063399C | `HandlePlayPause` | Known | Event handler |
| 0x006339AC | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x006339C8 | `HandlePrevious` | Known | Event handler |
| 0x006339F0 | `HandleCameraRemote` | Known | Event handler |
| 0x00633A44 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x00633A6C | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x00633ABC | `HandleRemoteMenuDown` | Known | Event handler |
| 0x00633AE0 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x00633B30 | `HandleRemotePlayPause` | Known | Event handler |
| 0x00633B54 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x00633B78 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x00633BF0 | `HandleMikeyAllUp` | Known | Event handler |
| 0x00633C1C | `HandleMikeyCenterReleased` | Known | Event handler |
| 0x00633C54 | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x00633C94 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x00633CD4 | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x00633D10 | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x00633D50 | `HandleMikeyCenterTripleClickAndHold` | Known | Event handler |
| 0x00633D84 | `HandleMikeyCenter` | Known | Event handler |
| 0x00633DB0 | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x00633DDC | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x00633E08 | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x00633E30 | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x00633EE4 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x00633F10 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x00633F40 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x00633F6C | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x00633F9C | `HandleRadioTagUp` | Known | Event handler |
| 0x00633FC4 | `HandleRemoteBacklight` | Known | Event handler |
| 0x00633FF4 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x00634050 | `HandleRemoteEvent` | Known | Event handler |
| 0x00634074 | `HandleRemoteFFDown` | Known | Event handler |
| 0x00634098 | `HandleRemoteFFUp` | Known | Event handler |
| 0x006340E4 | `HandleRemoteMute` | Known | Event handler |
| 0x00634110 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x00634140 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x00634174 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x006341AC | `HandleRemoteOff` | Known | Event handler |
| 0x006341DC | `HandleRemoteOn` | Known | Event handler |
| 0x00634200 | `HandleRemotePause` | Known | Event handler |
| 0x00634228 | `HandleRemotePlay` | Known | Event handler |
| 0x00634260 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x0063429C | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x006342CC | `HandleRemotePrevChapter` | Known | Event handler |
| 0x00634300 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x00634330 | `HandleRemoteRepeat` | Known | Event handler |
| 0x00634354 | `HandleRemoteRewDown` | Known | Event handler |
| 0x00634378 | `HandleRemoteRewUp` | Known | Event handler |
| 0x006343A0 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x006343CC | `HandleRemoteSelectUp` | Known | Event handler |
| 0x006343F8 | `HandleRemoteShuffle` | Known | Event handler |
| 0x00634420 | `HandleRemoteStop` | Known | Event handler |
| 0x006344C4 | `HandleRotationChange` | Known | Event handler |
| 0x006344F0 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x0063451C | `HandleVideoFFDown` | Known | Event handler |
| 0x0063453C | `HandleVideoFFUp` | Known | Event handler |
| 0x00634564 | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x00634590 | `HandleVideoNext` | Known | Event handler |
| 0x006345B8 | `HandleVideoNextChapter` | Known | Event handler |
| 0x006345E8 | `HandleVideoNextFrame` | Known | Event handler |
| 0x00634614 | `HandleVideoPause` | Known | Event handler |
| 0x00634638 | `HandleVideoPlay` | Known | Event handler |
| 0x0063465C | `HandleVideoPlayPause` | Known | Event handler |
| 0x0063468C | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x006346BC | `HandleVideoPrevious` | Known | Event handler |
| 0x006346E8 | `HandleVideoPrevChapter` | Known | Event handler |
| 0x00634718 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x00634740 | `HandleVideoRewDown` | Known | Event handler |
| 0x00634764 | `HandleVideoRewUp` | Known | Event handler |
| 0x00634788 | `HandleVideoStop` | Known | Event handler |
| 0x006347BC | `HandleVoiceOverAudiobookSelected` | Known | Event handler |
| 0x00634800 | `HandleVoiceOverPlaylistSelected` | Known | Event handler |
| 0x00634840 | `HandleVoiceOverPodcastSelected` | Known | Event handler |
| 0x00634890 | `HandleMenuSelection` | Known | Event handler |
| 0x006349B8 | `HandleLoadingCancelled` | Known | Event handler |
| 0x00634A28 | `HandleDialogDismiss` | Known | Event handler |
| 0x00634CF8 | `HandleGotoCalendarLoadingScreen` | Known | Event handler |
| 0x00634D78 | `HandleGotoAddressBookScreen` | Known | Event handler |
| 0x00634E48 | `HandleRadioSelected` | Known | Event handler |
| 0x00634E7C | `HandleRadioPlayPause` | Known | Event handler |
| 0x00634F2C | `HandleVoiceMemosSelected` | Known | Event handler |
| 0x00634FB4 | `HandleSteer` | Known | Event handler |
| 0x00634FD0 | `HandleItemSelected` | Known | Event handler |
| 0x00635080 | `HandleNextContact` | Known | Event handler |
| 0x00635094 | `HandlePreviousContact` | Known | Event handler |
| 0x006350AC | `HandleSelect` | Known | Event handler |
| 0x006350CC | `HandleHilite` | Known | Event handler |
| 0x00635490 | `HandleDateChosen` | Known | Event handler |
| 0x006354C0 | `HandleTimeChosen` | Known | Event handler |
| 0x006354F0 | `HandleFrequencyChosen` | Known | Event handler |
| 0x00635524 | `HandleSoundChosen` | Known | Event handler |
| 0x00635554 | `HandleLabelChosen` | Known | Event handler |
| 0x00635584 | `HandleDeleteChosen` | Known | Event handler |
| 0x00635704 | `HandleLeaveAlarm` | Known | Event handler |
| 0x0063588C | `HandleNextDay` | Known | Event handler |
| 0x0063589C | `HandlePreviousDay` | Known | Event handler |
| 0x00635938 | `HandleNextAndHold` | Known | Event handler |
| 0x0063594C | `HandlePrevAndHold` | Known | Event handler |
| 0x00635960 | `HandlePrev` | Known | Event handler |
| 0x00635988 | `HandleSelectPressAndHold` | Known | Event handler |
| 0x006359A4 | `HandleCameraRemoteSelect` | Known | Event handler |
| 0x006359DC | `HandleAppInitialized` | Known | Event handler |
| 0x00635B10 | `HandleScreenRotation` | Known | Event handler |
| 0x00635B84 | `HandleWheel` | Known | Event handler |
| 0x00635BB4 | `HandleDeleteAllItems` | Known | Event handler |
| 0x00635BEC | `HandleDeleteItem` | Known | Event handler |
| 0x00636014 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006360DC | `HandleSelectClock` | Known | Event handler |
| 0x0063610C | `HandleHilited` | Known | Event handler |
| 0x00636154 | `HandleSelectRegion_Africa` | Known | Event handler |
| 0x006361C8 | `HandleSelectRegion_Asia` | Known | Event handler |
| 0x006361FC | `HandleSelectRegion_Atlantic` | Known | Event handler |
| 0x00636234 | `HandleSelectRegion_Australia` | Known | Event handler |
| 0x00636270 | `HandleSelectRegion_Europe` | Known | Event handler |
| 0x006362A8 | `HandleSelectRegion_NorthAmerica` | Known | Event handler |
| 0x006362E4 | `HandleSelectRegion_Pacific` | Known | Event handler |
| 0x0063631C | `HandleSelectRegion_SouthAmerica` | Known | Event handler |
| 0x0063633C | `HandleSelectCity` | Known | Event handler |
| 0x00636370 | `HandleAddClock` | Known | Event handler |
| 0x006363A4 | `HandleDeleteClock` | Known | Event handler |
| 0x006363D8 | `HandleEditClock` | Known | Event handler |
| 0x006365CC | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00636624 | `HandleSwapToNowPlaying` | Known | Event handler |
| 0x0063663C | `HandleFlowNext` | Known | Event handler |
| 0x0063664C | `HandleFlowPrev` | Known | Event handler |
| 0x0063665C | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x006366A0 | `HandleAlbumSelected` | Known | Event handler |
| 0x006367D8 | `HandleSwapToNowPlayingFromOrientation` | Known | Event handler |
| 0x00636800 | `HandleFlowWheel` | Known | Event handler |
| 0x0063682C | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x00636A1C | `HandleArrowDown` | Known | Event handler |
| 0x00636A40 | `HandleArrowUp` | Known | Event handler |
| 0x00636ABC | `HandleGameHilited` | Known | Event handler |
| 0x00636CF0 | `HandleGestureShake` | Known | Event handler |
| 0x00636D4C | `HandleOrientationLandscape` | Known | Event handler |
| 0x00636D68 | `HandleOrientationPortrait` | Known | Event handler |
| 0x00636D84 | `HandleGestureSteer` | Known | Event handler |
| 0x00636DD0 | `HandleOrientionAlt` | Known | Event handler |
| 0x00636E2C | `HandlePauseKey` | Known | Event handler |
| 0x0063727C | `HandleOrientationAlt` | Known | Event handler |
| 0x006372B0 | `HandleMusicSelected` | Known | Event handler |
| 0x006372E0 | `HandleMusicHilited` | Known | Event handler |
| 0x00637310 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00637344 | `HandleCoverflowHilited` | Known | Event handler |
| 0x00637378 | `HandleGeniusMixesSelected` | Known | Event handler |
| 0x006373B0 | `HandleGeniusMixesHilited` | Known | Event handler |
| 0x006373E8 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x0063741C | `HandlePlaylistsHilited` | Known | Event handler |
| 0x00637450 | `HandleArtistsSelected` | Known | Event handler |
| 0x00637484 | `HandleArtistsHilited` | Known | Event handler |
| 0x006374B8 | `HandleAlbumsSelected` | Known | Event handler |
| 0x006374EC | `HandleAlbumsHilited` | Known | Event handler |
| 0x0063751C | `HandleCompilationsSelected` | Known | Event handler |
| 0x00637554 | `HandleCompilationsHilited` | Known | Event handler |
| 0x0063758C | `HandleSongsSelected` | Known | Event handler |
| 0x006375BC | `HandleSongsHilited` | Known | Event handler |
| 0x006375EC | `HandleGenresSelected` | Known | Event handler |
| 0x00637620 | `HandleGenresHilited` | Known | Event handler |
| 0x00637650 | `HandleComposersSelected` | Known | Event handler |
| 0x00637684 | `HandleComposersHilited` | Known | Event handler |
| 0x006376B8 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x006376F0 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x006377B8 | `HandleVideosSelected` | Known | Event handler |
| 0x006377EC | `HandleVideosHilited` | Known | Event handler |
| 0x0063781C | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00637858 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00637890 | `HandleMoviesSelected` | Known | Event handler |
| 0x006378C4 | `HandleMoviesHilited` | Known | Event handler |
| 0x006378F4 | `HandleTVShowsSelected` | Known | Event handler |
| 0x00637928 | `HandleTVShowsHilited` | Known | Event handler |
| 0x0063795C | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00637994 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x006379CC | `HandleRentalsSelected` | Known | Event handler |
| 0x00637A00 | `HandleRentalsHilited` | Known | Event handler |
| 0x00637A34 | `HandlePhotosSelected` | Known | Event handler |
| 0x00637A68 | `HandlePhotosHilited` | Known | Event handler |
| 0x00637A98 | `HandlePodcastsSelected` | Known | Event handler |
| 0x00637ACC | `HandlePodcastsHilited` | Known | Event handler |
| 0x00637B00 | `HandleiTunesUSelected` | Known | Event handler |
| 0x00637B34 | `HandleiTunesUHilited` | Known | Event handler |
| 0x00637B84 | `HandleRadioHilited` | Known | Event handler |
| 0x00637BD0 | `HandleCameraSelected` | Known | Event handler |
| 0x00637C04 | `HandleCameraHilited` | Known | Event handler |
| 0x00637C68 | `HandleExtrasHilited` | Known | Event handler |
| 0x00637DC4 | `HandleAddressBookSelected` | Known | Event handler |
| 0x00637F74 | `HandleNikePlusSelected` | Known | Event handler |
| 0x00637FA8 | `HandleNikePlusHilited` | Known | Event handler |
| 0x00638070 | `HandleVoiceMemosHilited` | Known | Event handler |
| 0x0063811C | `HandleGenericHilited` | Known | Event handler |
| 0x006381AC | `HandleSettingsHilited` | Known | Event handler |
| 0x00638200 | `HandleSleepSelected` | Known | Event handler |
| 0x00638268 | `HandlePedometerSelected` | Known | Event handler |
| 0x0063829C | `HandlePedometerHilited` | Known | Event handler |
| 0x006382D0 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00638344 | `HandleNowPlaying` | Known | Event handler |
| 0x00638374 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x0063838C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00639A9C | `HandlePortraitToLandscape` | Known | Event handler |
| 0x00639AB8 | `HandleRadioPreviewPlayPause` | Known | Event handler |
| 0x00639C74 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00639CB0 | `HandleAddToOTG` | Known | Event handler |
| 0x00639D80 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00639DDC | `HandleStartGenius` | Known | Event handler |
| 0x00639E10 | `HandleViewAlbum` | Known | Event handler |
| 0x00639E40 | `HandleViewArtist` | Known | Event handler |
| 0x00639E7C | `HandleViewCompilation` | Known | Event handler |
| 0x0063A1EC | `HandleRefreshPlaylist` | Known | Event handler |
| 0x0063A23C | `HandleSaveGeniusPlaylist` | Known | Event handler |
| 0x0063A258 | `HandleSelectMix` | Known | Event handler |
| 0x0063A290 | `HandleGeniusMixPlaylistReady` | Known | Event handler |
| 0x0063A5C0 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x0063A750 | `HandleCameraVideosSelected` | Known | Event handler |
| 0x0063AB70 | `HandleTVOutChanged` | Known | Event handler |
| 0x0063ABA0 | `HandleTVSignalChanged` | Known | Event handler |
| 0x0063ABD4 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x0063AC1C | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x0063AC58 | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x0063AC98 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x0063ACD4 | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x0063ACF4 | `HandleMenuLongpress` | Known | Event handler |
| 0x0063AD08 | `HandleMenuKey` | Known | Event handler |
| 0x0063AD50 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0063AD94 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0063ADD4 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0063AE14 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0063B340 | `HandleSwapToMusicScreen` | Known | Event handler |
| 0x0063B378 | `HandleSwapToVideoScreen` | Known | Event handler |
| 0x0063B398 | `HandleMTWheel` | Known | Event handler |
| 0x0063B3A8 | `HandleSwapToCoverflow` | Known | Event handler |
| 0x0063B3C0 | `HandleDefaultOrientation` | Known | Event handler |
| 0x0063B3DC | `HandleSelectProgress` | Known | Event handler |
| 0x0063B3F4 | `HandleWheelProgress` | Known | Event handler |
| 0x0063B408 | `HandleSelectVolume` | Known | Event handler |
| 0x0063B41C | `HandleWheelVolume` | Known | Event handler |
| 0x0063B430 | `HandleSelectGenius` | Known | Event handler |
| 0x0063B444 | `HandleWheelGenius` | Known | Event handler |
| 0x0063B458 | `HandleSelectRating` | Known | Event handler |
| 0x0063B46C | `HandleWheelRating` | Known | Event handler |
| 0x0063B498 | `HandleSelectScrub` | Known | Event handler |
| 0x0063B4AC | `HandleWheelScrub` | Known | Event handler |
| 0x0063B4C0 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0063B4DC | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x0063B4F8 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x0063B52C | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x0063B564 | `HandleOrientationChange` | Known | Event handler |
| 0x0063B588 | `HandleWheelBrightness` | Known | Event handler |
| 0x0063B5A0 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x0063B5BC | `HandlePlayPauseTV` | Known | Event handler |
| 0x0063B6E4 | `HandleCenterButtonSelected` | Known | Event handler |
| 0x0063B8B4 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0063B8D0 | `HandleNotesPop` | Known | Event handler |
| 0x0063BA9C | `HandleNotesSelected` | Known | Event handler |
| 0x0063BBE0 | `HandlePushEvents` | Known | Event handler |
| 0x0063BC44 | `HandlePushFaces` | Known | Event handler |
| 0x0063BCA0 | `HandlePushPlaces` | Known | Event handler |
| 0x0063BD50 | `HandleBrowseAlbum` | Known | Event handler |
| 0x0063BD64 | `HandleHiliteAlbum` | Known | Event handler |
| 0x0063BD78 | `HandlePopEvents` | Known | Event handler |
| 0x0063BD88 | `HandlePopFaces` | Known | Event handler |
| 0x0063BD98 | `HandlePopPlaces` | Known | Event handler |
| 0x0063BFB4 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x0063BFD8 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x0063C068 | `HandleImageLast` | Known | Event handler |
| 0x0063C078 | `HandleScreenNext` | Known | Event handler |
| 0x0063C08C | `HandleImageFirst` | Known | Event handler |
| 0x0063C0A0 | `HandleScreenPrev` | Known | Event handler |
| 0x0063C0B4 | `HandleBrowseLarge` | Known | Event handler |
| 0x0063C1B0 | `HandleImageNext` | Known | Event handler |
| 0x0063C1D0 | `HandleImagePrev` | Known | Event handler |
| 0x0063C1E0 | `HandleMenuUp` | Known | Event handler |
| 0x0063C1F0 | `HandlePrevPressAndHold` | Known | Event handler |
| 0x0063C700 | `HandlePause` | Known | Event handler |
| 0x0063C70C | `HandlePlay` | Known | Event handler |
| 0x0063C728 | `HandleMikeyPlayPause` | Known | Event handler |
| 0x0063C740 | `HandleOrientationDefault` | Known | Event handler |
| 0x0063CC68 | `HandleReadyForLargeBrowse` | Known | Event handler |
| 0x0063CD58 | `HandleNextTuning` | Known | Event handler |
| 0x0063CD78 | `HandlePreviousTuning` | Known | Event handler |
| 0x0063CD90 | `HandleTunerContextMenu` | Known | Event handler |
| 0x0063D0CC | `HandleMikeyNext` | Known | Event handler |
| 0x0063D0DC | `HandleMikeyPrevious` | Known | Event handler |
| 0x0063D0F0 | `HandleVolumeWheel` | Known | Event handler |
| 0x0063D18C | `HandleBufferWheel` | Known | Event handler |
| 0x0063D1B4 | `HandleTunerWheel` | Known | Event handler |
| 0x0063D1F8 | `HandleBandWheel` | Known | Event handler |
| 0x0063D208 | `HandleNextPress` | Known | Event handler |
| 0x0063D218 | `HandlePreviousPress` | Known | Event handler |
| 0x0063D328 | `HandleTogglePlayPause` | Known | Event handler |
| 0x0063D394 | `HandlePlayRadio` | Known | Event handler |
| 0x0063D3C0 | `HandleStopRadio` | Known | Event handler |
| 0x0063D3EC | `HandleAutoTune` | Known | Event handler |
| 0x0063D580 | `HandleToggleBufferSetting` | Known | Event handler |
| 0x0063D5B8 | `HandleScanLogging` | Known | Event handler |
| 0x0063D638 | `HandleRemovePreset` | Known | Event handler |
| 0x0063D664 | `HandleSelectPreset` | Known | Event handler |
| 0x0063D698 | `HandleConfirmation` | Known | Event handler |
| 0x0063D80C | `HandleExitUnsupported` | Known | Event handler |
| 0x0063D8F8 | `HandlePushToCount` | Known | Event handler |
| 0x0063D90C | `HandlePopToBasic` | Known | Event handler |
| 0x0063D920 | `HandlePushToBasic` | Known | Event handler |
| 0x0063D934 | `HandlePopToCapacity` | Known | Event handler |
| 0x0063D948 | `HandlePushToCapacity` | Known | Event handler |
| 0x0063D960 | `HandlePopToCount` | Known | Event handler |
| 0x0063D974 | `HandlePushToAccessoryCount` | Known | Event handler |
| 0x0063D990 | `HandlePopToAccessoryAccessory` | Known | Event handler |
| 0x0063D9B0 | `HandlePushToAccessoryBasic` | Known | Event handler |
| 0x0063D9CC | `HandlePopToAccessoryCapacity` | Known | Event handler |
| 0x0063D9EC | `HandlePushToAccessoryAccessory` | Known | Event handler |
| 0x0063DA0C | `HandlePopToAccessoryCount` | Known | Event handler |
| 0x0063DA28 | `HandlePushToAccessoryCapacity` | Known | Event handler |
| 0x0063DA48 | `HandlePopToAccessoryBasic` | Known | Event handler |
| 0x0063DDE8 | `HandleResetAllSettings` | Known | Event handler |
| 0x0063E1A0 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x0063E3A8 | `HandleMainMenu` | Known | Event handler |
| 0x0063E840 | `HandleMusicMenu` | Known | Event handler |
| 0x0063EBF4 | `HandleRadioRegion` | Known | Event handler |
| 0x0063EC14 | `HandleLanguagePop` | Known | Event handler |
| 0x0063EC5C | `HandleLanguage` | Known | Event handler |
| 0x0063EFC0 | `HandleSelectKey` | Known | Event handler |
| 0x0063EFD0 | `HandleExit` | Known | Event handler |
| 0x0063EFDC | `HandleStartStop` | Known | Event handler |
| 0x0063F020 | `HandleLap` | Known | Event handler |
| 0x0063F02C | `HandleChosen` | Known | Event handler |
| 0x0063F0A8 | `HandleDelete` | Known | Event handler |
| 0x0063F5B8 | `HandleSelectedNikeMainMenuItem` | Known | Event handler |
| 0x0063FBB4 | `HandleBasicSelected` | Known | Event handler |
| 0x0063FBE4 | `HandleTimedWorkoutSelected` | Known | Event handler |
| 0x0063FC1C | `HandleDistanceWorkoutSelected` | Known | Event handler |
| 0x0063FC58 | `HandleCaloriesWorkoutSelected` | Known | Event handler |
| 0x0063FC9C | `HandleChooseTimedWorkoutFromList` | Known | Event handler |
| 0x0063FCC0 | `HandleChooseDistanceWorkoutFromList` | Known | Event handler |
| 0x0063FCE4 | `HandleChooseCalorieWorkoutFromList` | Known | Event handler |
| 0x0063FD08 | `HandleBegin` | Known | Event handler |
| 0x0063FF98 | `HandleLinkNewRemote` | Known | Event handler |
| 0x00640028 | `HandleLinkNewHeartMonitor` | Known | Event handler |
| 0x006400B4 | `HandleCancelRemoteLinking` | Known | Event handler |
| 0x006400F4 | `HandleChooseMusic` | Known | Event handler |
| 0x00640128 | `HandleEndWorkout` | Known | Event handler |
| 0x00640160 | `HandlePauseWorkout` | Known | Event handler |
| 0x00640198 | `HandleResumeWorkout` | Known | Event handler |
| 0x00640548 | `HandleNewWorkout` | Known | Event handler |
| 0x006405A8 | `HandleNewBasicWorkout` | Known | Event handler |
| 0x006405C0 | `HandleNewQuickstartWorkout` | Known | Event handler |
| 0x00640650 | `HandleChoosePedometer` | Known | Event handler |
| 0x006406E8 | `HandleChoosePowerPlay` | Known | Event handler |
| 0x0064071C | `HandleChooseVoicePrompts` | Known | Event handler |
| 0x00640754 | `HandleChooseUnit` | Known | Event handler |
| 0x00640A0C | `HandleChooseRemoteSetting` | Known | Event handler |
| 0x00640A44 | `HandleChooseRemoteLink` | Known | Event handler |
| 0x00640A78 | `HandleChooseRemoteUnlink` | Known | Event handler |
| 0x00640BDC | `HandleChooseHeartMonitorLink` | Known | Event handler |
| 0x00640C18 | `HandleChooseHeartMonitorUnlink` | Known | Event handler |
| 0x00640D44 | `HandleListChoose` | Known | Event handler |
| 0x00640D58 | `HandlePopBackToSongsScreen` | Known | Event handler |
| 0x00640D90 | `HandleVerticalSelected` | Known | Event handler |
| 0x00640DC4 | `HandleRightSelected` | Known | Event handler |
| 0x00640DF4 | `HandleLeftSelected` | Known | Event handler |
| 0x00640E98 | `HandleChooseLink` | Known | Event handler |
| 0x00640F30 | `HandleMenuKeyNop` | Known | Event handler |
| 0x00640F68 | `HandlePauseHold` | Known | Event handler |
| 0x00640F8C | `HandleSelectKeyDown` | Known | Event handler |
| 0x00640FA0 | `HandlePowerPlay` | Known | Event handler |
| 0x00640FB0 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x00641574 | `Handle3BitModeFinished` | Known | Event handler |
| 0x0064158C | `HandleMikeyPressExtended` | Known | Event handler |
| 0x00641A98 | `HandleNowPlayingSelected` | Known | Event handler |
| 0x00641AD0 | `HandleChooseGeniusMixesPlay` | Known | Event handler |
| 0x00641B08 | `HandleChoosePlaylistsPlay` | Known | Event handler |
| 0x00641B40 | `HandleChoosePodcastsPlay` | Known | Event handler |
| 0x00641B78 | `HandleChooseAudiobooksPlay` | Known | Event handler |
| 0x00641BCC | `HandleNoneSelected` | Known | Event handler |
| 0x00641C4C | `HandleNikeNestedPlaylistSelect` | Known | Event handler |
| 0x00641FD4 | `HandleSelectedDayWorkout` | Known | Event handler |
| 0x00642040 | `HandleClearBests` | Known | Event handler |
| 0x006420A4 | `HandleClearTotals` | Known | Event handler |
| 0x006420B8 | `HandleHerculesKey` | Known | Event handler |
| 0x006420CC | `HandlePopSelf` | Known | Event handler |
| 0x006420DC | `HandlePressAndHold` | Known | Event handler |
| 0x006421C4 | `HandleSelectNextWorkout` | Known | Event handler |
| 0x006421DC | `HandleSelectPrevWorkout` | Known | Event handler |
| 0x0064226C | `HandleDeleteAllWorkouts` | Known | Event handler |
| 0x006422A8 | `HandleDeleteWorkout` | Known | Event handler |
| 0x006424B0 | `HandleNextKey` | Known | Event handler |
| 0x006424C0 | `HandlePrevKey` | Known | Event handler |
| 0x006426AC | `HandleWeightSelect` | Known | Event handler |
| 0x006426C0 | `HandleWeightWheel` | Known | Event handler |
| 0x006426D4 | `HandleWeightSelectAltTrans` | Known | Event handler |
| 0x006426F0 | `HandleWeightSelectPedometer` | Known | Event handler |
| 0x0064270C | `HandleDistanceSelect` | Known | Event handler |
| 0x00642724 | `HandleDistanceWheel` | Known | Event handler |
| 0x00642738 | `HandleTimeSelect` | Known | Event handler |
| 0x0064274C | `HandleTimeWheel` | Known | Event handler |
| 0x0064275C | `HandleCaloriesSelect` | Known | Event handler |
| 0x00642774 | `HandleCaloriesWheel` | Known | Event handler |
| 0x00642788 | `HandleStepGoalSelect` | Known | Event handler |
| 0x006427A0 | `HandleStepGoalWheel` | Known | Event handler |
| 0x0064284C | `HandleWalkCalibrationSelection` | Known | Event handler |
| 0x00642888 | `HandleRunCalibrationSelection` | Known | Event handler |
| 0x00642974 | `HandleResetRuningCalibration` | Known | Event handler |
| 0x00642994 | `HandleResetWalkingCalibration` | Known | Event handler |
| 0x00642C90 | `HandleUnlinkRemote` | Known | Event handler |
| 0x00642CE0 | `HandleUnlinkHeartMonitor` | Known | Event handler |
| 0x00643138 | `HandleShowRecordings` | Known | Event handler |
| 0x00643150 | `HandleAddChapterMark` | Known | Event handler |
| 0x006433EC | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x006434DC | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x00643534 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x006435E8 | `HandleSelectLabel` | Known | Event handler |
| 0x006435FC | `HandleDeleteAllSelect` | Known | Event handler |
| 0x00643614 | `HandleDeleteSelect` | Known | Event handler |
| 0x00643628 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x00643644 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x00643664 | `HandleMicrophoneRequired` | Known | Event handler |
| 0x00643680 | `HandleMicrophoneDisconnected` | Known | Event handler |
| 0x00643724 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x006437B4 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x006437D0 | `HandleEQSettingSelected` | Known | Event handler |
| 0x006437E8 | `HandleEQSettingPreview` | Known | Event handler |

---

## 6. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00079EE0 | `GotoNowPlaying` | Known | Navigation |
| 0x0007F3F8 | `GotoDefaultLayout` | Known | Navigation |
| 0x0007F464 | `GotoVolumeLayout` | Known | Navigation |
| 0x0007FAEC | `GotoStatusBarLayout` | Known | Navigation |
| 0x0007FB00 | `GotoDefaultLayout` | Known | Navigation |
| 0x0007FC50 | `GotoDefault` | Known | Navigation |
| 0x0007FD84 | `GotoProgressLayout` | Known | Navigation |
| 0x0007FED4 | `GotoBrightnessLayout` | Known | Navigation |
| 0x0007FF38 | `GotoBrightnessLayout` | Known | Navigation |
| 0x0007FF90 | `GotoVolumeLayout` | Known | Navigation |
| 0x0007FFC8 | `GotoScrubLayout` | Known | Navigation |
| 0x00080060 | `GotoStatusBarLayout` | Known | Navigation |
| 0x00080074 | `GotoDefaultLayout` | Known | Navigation |
| 0x000800D0 | `GotoScrubLayout` | Known | Navigation |
| 0x0008010C | `GotoScrubLayout` | Known | Navigation |
| 0x00080550 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x0008056C | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00080588 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x000805A4 | `GotoDefaultLayout` | Known | Navigation |
| 0x0008061C | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x00080634 | `GotoVolumeLayout` | Known | Navigation |
| 0x00080D40 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00080E8C | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x00080EA8 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00080EC4 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00080EE0 | `GotoDefaultLayout` | Known | Navigation |
| 0x000890D4 | `GotoMainMenu` | Known | Navigation |
| 0x0008E0C0 | `GotoDialogueScreen` | Known | Navigation |
| 0x000A7440 | `GotoNowPlaying` | Known | Navigation |
| 0x000A7454 | `GotoAlbums` | Known | Navigation |
| 0x000A7460 | `GotoSongs` | Known | Navigation |
| 0x000B3038 | `GotoNowPlaying` | Known | Navigation |
| 0x000B3270 | `GotoNowPlaying` | Known | Navigation |
| 0x000B4774 | `GotoScreen_PlaybackSettingsMenu` | Known | Navigation |
| 0x000B5844 | `GotoMediumPedometerLayout` | Known | Navigation |
| 0x000B69D8 | `GotoProgressLayout` | Known | Navigation |
| 0x000B6A84 | `GotoProgressLayout` | Known | Navigation |
| 0x000B6B70 | `GotoProgressLayout` | Known | Navigation |
| 0x000B6E20 | `GotoProgressLayout` | Known | Navigation |
| 0x000B7388 | `GotoMainCalendarPage` | Known | Navigation |
| 0x000B835C | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x000B8464 | `GotoFourCard_About` | Known | Navigation |
| 0x000B8478 | `GotoThreeCard_About` | Known | Navigation |
| 0x000B86E0 | `GotoScreen_ResetAllSettings` | Known | Navigation |
| 0x000B8960 | `GotoScreen_BacklightTimer` | Known | Navigation |
| 0x000B89D0 | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation |
| 0x000B89F4 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x000B908C | `GotoProgressVideoLayout` | Known | Navigation |
| 0x000B90A4 | `GotoProgressLayout` | Known | Navigation |
| 0x000B9344 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x000B935C | `GotoProgressLayout` | Known | Navigation |
| 0x000B942C | `GotoProgressLayout` | Known | Navigation |
| 0x000B954C | `GotoBrightnessVideoLayout` | Known | Navigation |
| 0x000B9568 | `GotoGeniusLayout` | Known | Navigation |
| 0x000B957C | `GotoRatingLayout` | Known | Navigation |
| 0x000B96E8 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x000B979C | `GotoRatingLayout` | Known | Navigation |
| 0x000B987C | `GotoShuffleLayout` | Known | Navigation |
| 0x000B98B4 | `GotoDefaultLayout` | Known | Navigation |
| 0x000B9B64 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x000B9B7C | `GotoVolumeLayout` | Known | Navigation |
| 0x000B9C04 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x000B9C1C | `GotoVolumeLayout` | Known | Navigation |
| 0x000B9D18 | `GotoScrubVideoLayout` | Known | Navigation |
| 0x000B9D30 | `GotoScrubLayout` | Known | Navigation |
| 0x000B9DA4 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x000B9DBC | `GotoProgressLayout` | Known | Navigation |
| 0x000BA2E8 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x000BA300 | `GotoProgressLayout` | Known | Navigation |
| 0x000BA390 | `GotoExtraInfoLoadFailedLayout` | Known | Navigation |
| 0x000BA3B0 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x000BAF70 | `GotoScreen_AddressViewerLoaded` | Known | Navigation |
| 0x000C4764 | `GotoScreen_LockediPod` | Known | Navigation |
| 0x000C477C | `GotoScreen_EnterPassKey` | Known | Navigation |
| 0x000C6510 | `GotoCoverFlowScreenBackside` | Known | Navigation |
| 0x000CB30C | `GotoScreenMainMenu` | Known | Navigation |
| 0x000CC8B0 | `GotoScreen_Language` | Known | Navigation |
| 0x000CD414 | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation |
| 0x000E2190 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x000E21B4 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x000E2480 | `GotoScreen_AddressViewerLoading` | Known | Navigation |
| 0x000E24A0 | `GotoScreen_AddressViewerLoaded` | Known | Navigation |
| 0x000E24FC | `GotoScreen_CalendarViewerLoading` | Known | Navigation |
| 0x000E2520 | `GotoScreen_CalendarView` | Known | Navigation |
| 0x000E3F78 | `GotoScreen_LockediPod` | Known | Navigation |
| 0x000E416C | `GotoScreen_MainMenu` | Known | Navigation |
| 0x000ECC40 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x000EE790 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x000EF8C0 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x000EF9D0 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x000EFB34 | `GotoPlaylistScreen` | Known | Navigation |
| 0x000EFD4C | `GotoNowPlaying` | Known | Navigation |
| 0x000EFF54 | `GotoNowPlaying` | Known | Navigation |
| 0x000F2A80 | `GotoFirstBoot` | Known | Navigation |
| 0x000F2A90 | `GotoNotesApp` | Known | Navigation |
| 0x000F2AA4 | `GotoLockApp` | Known | Navigation |
| 0x000F35BC | `GotoPlaylists` | Known | Navigation |
| 0x000F3AD8 | `GotoGenius` | Known | Navigation |
| 0x000F3D00 | `GotoGenius` | Known | Navigation |
| 0x000F3D0C | `GotoGeniusIntro` | Known | Navigation |
| 0x000F485C | `GotoNowPlaying` | Known | Navigation |
| 0x000F488C | `GotoGeniusMixLoadingScreen` | Known | Navigation |
| 0x000F53AC | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x000F54DC | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x000FC5B0 | `GotoGameRunningLayout` | Known | Navigation |
| 0x000FC73C | `GotoGameRunningLayout` | Known | Navigation |
| 0x0010C688 | `GotoNowPlaying` | Known | Navigation |
| 0x0010C6E8 | `GotoScreen_AddressViewerLoading` | Known | Navigation |
| 0x0010C708 | `GotoScreen_AddressViewerLoaded` | Known | Navigation |
| 0x0010C728 | `GotoScreen_AddressViewerNoContacts` | Known | Navigation |
| 0x0010C880 | `GotoGeniusMixesIntro` | Known | Navigation |
| 0x0010CBC4 | `GotoNowPlaying` | Known | Navigation |
| 0x0012CB64 | `GotoErrorLayout` | Known | Navigation |
| 0x001B3C74 | `GotoGeniusMixesIntro` | Known | Navigation |
| 0x001B6114 | `GotoCalibrateRcvMissing` | Known | Navigation |
| 0x001B6190 | `GotoCalibrateRcvMissing` | Known | Navigation |
| 0x001B676C | `GotoGeniusMixLoadingScreen` | Known | Navigation |
| 0x001B76CC | `GotoSettings` | Known | Navigation |
| 0x001B76DC | `GotoCustomStepGoal` | Known | Navigation |
| 0x001BBC60 | `GotoRecording` | Known | Navigation |
| 0x002109CC | `GotoExtraInfoLayout` | Known | Navigation |
| 0x002109E0 | `GotoExtraInfoLoadingLayout` | Known | Navigation |
| 0x00212820 | `GotoDefault` | Known | Navigation |
| 0x00216CAC | `GotoNowPlaying` | Known | Navigation |
| 0x0023AE24 | `GotoGeniusMixLoadingScreen` | Known | Navigation |
| 0x00241370 | `GotoDefaultLayout` | Known | Navigation |
| 0x0024C60C | `GotoEnteringNowPlaying` | Known | Navigation |
| 0x002522C0 | `GotoRentalWarningLayout` | Known | Navigation |
| 0x0026436C | `GotoNowPlaying` | Known | Navigation |
| 0x0028E198 | `GotoNowPlaying` | Known | Navigation |
| 0x0029732C | `GotoNowPlaying` | Known | Navigation |
| 0x00636728 | `GotoCoverFlowScreenBackside` | Known | Navigation |
| 0x0063B57C | `GotoDefault` | Known | Navigation |
| 0x0063DF8C | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |

---

## 7. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0056DBE0 | `VoiceMemos_Screen_InsufficientDiskSpace2` | Known | Screen layout |
| 0x0056DC09 | `NikePlus_EndWorkout_Screen_Contextual_Default_L` | Known | Screen layout |
| 0x0056DC39 | `MainMenus_Main_Screen_NoCamera` | Known | Screen layout |
| 0x0056DC58 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0056DC70 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0056DC8E | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x0056DCB2 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x0056DCD3 | `NowPlaying_Screen_Music` | Known | Screen layout |
| 0x0056DCEB | `NoContent_Screen_Music` | Known | Screen layout |
| 0x0056DD02 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0056DD20 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0056DD39 | `NikePlus_NowRunning_Screen_BasicNoMusic` | Known | Screen layout |
| 0x0056DD61 | `NikePlus_NowRunning_Screen_DistanceNoMusic` | Known | Screen layout |
| 0x0056DD8C | `NikePlus_NowRunning_Screen_TimeNoMusic` | Known | Screen layout |
| 0x0056DDB3 | `NikePlus_NowRunning_Screen_Basic_LandscapeNoMusic` | Known | Screen layout |
| 0x0056DDE5 | `NikePlus_NowRunning_Screen_Distance_LandscapeNoMusic` | Known | Screen layout |
| 0x0056DE1A | `NikePlus_NowRunning_Screen_Time_LandscapeNoMusic` | Known | Screen layout |
| 0x0056DE4B | `NikePlus_NowRunning_Screen_Calibrate_LandscapeNoMusic` | Known | Screen layout |
| 0x0056DE81 | `NikePlus_NowRunning_Screen_Calories_LandscapeNoMusic` | Known | Screen layout |
| 0x0056DEB6 | `NikePlus_NowRunning_Screen_CalibrateNoMusic` | Known | Screen layout |
| 0x0056DEE2 | `NoContent_Screen_MainNoMusic` | Known | Screen layout |
| 0x0056DEFF | `NikePlus_NowRunning_Screen_CaloriesNoMusic` | Known | Screen layout |
| 0x0056DF2A | `NikePlus_Custom_Screen_Weight_ToWorkoutMusic` | Known | Screen layout |
| 0x0056DF57 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x0056DF82 | `DeleteRental_Confirmation_Screen_Unwatched` | Known | Screen layout |
| 0x0056DFAD | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x0056DFCB | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0056E0DD | `NowPlaying_Screen_ExtraInfoLoadFailed` | Known | Screen layout |
| 0x0056E103 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x0056E11C | `VoiceMemos_Screen_MicrophoneRequired` | Known | Screen layout |
| 0x0056E141 | `PhotosGL_Camera_Screen_Paused` | Known | Screen layout |
| 0x0056E15F | `PhotosGL_Camera_Alt_Screen_Paused` | Known | Screen layout |
| 0x0056E181 | `PhotosGL_TvOut_Screen_Paused` | Known | Screen layout |
| 0x0056E19E | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x0056E1BC | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x0056E1DC | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x0056E1FF | `VoiceMemos_Screen_MicrophoneDisconnected` | Known | Screen layout |
| 0x0056E228 | `VoiceMemos_Status_Screen_Inserted` | Known | Screen layout |
| 0x0056E24A | `VoiceMemos_Screen_Recording_ChapterInserted` | Known | Screen layout |
| 0x0056E276 | `Camera_Screen_Uninitialized` | Known | Screen layout |
| 0x0056E292 | `Search_Main_Screen_WithKeyboard` | Known | Screen layout |
| 0x0056E2B2 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x0056E2D0 | `VoiceMemos_Screen_InsufficientDiskSpace` | Known | Screen layout |
| 0x0056E2F8 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x0056E31C | `NikePlus_Custom_Screen_Distance` | Known | Screen layout |
| 0x0056E33C | `NikePlus_Custom_Screen_Simple_CalibrationDistance` | Known | Screen layout |
| 0x0056E36E | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x0056E399 | `CoverFlow_Screen_Backside` | Known | Screen layout |
| 0x0056E3B3 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x0056E3D0 | `NowPlaying_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x0056E3F5 | `PhotosGL_Camera_Screen_TvOut_ConnectCable` | Known | Screen layout |
| 0x0056E42F | `VoiceMemos_Screen_Idle` | Known | Screen layout |
| 0x0056E446 | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x0056E460 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x0056E496 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0056E4C2 | `NikePlus_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0056E4EC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0056E514 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x0056E534 | `NikePlus_Custom_Screen_Time` | Known | Screen layout |
| 0x0056E550 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0056E572 | `PhotosGL_Camera_Screen_Volume` | Known | Screen layout |
| 0x0056E590 | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x0056E5A9 | `PhotosGL_Camera_Alt_Screen_Volume` | Known | Screen layout |
| 0x0056E5CB | `PhotosGL_TvOut_Screen_Volume` | Known | Screen layout |
| 0x0056E5E8 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0056E607 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x0056E62C | `NikePlus_StartWorkout_Screen_Resume` | Known | Screen layout |
| 0x0056E650 | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x0056E67A | `NikePlus_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x0056E6A2 | `NikePlus_NowRunning_Screen_Basic_Landscape` | Known | Screen layout |
| 0x0056E6CD | `NikePlus_NowRunning_Screen_Distance_Landscape` | Known | Screen layout |
| 0x0056E6FB | `NikePlus_NowRunning_Screen_Time_Landscape` | Known | Screen layout |
| 0x0056E725 | `NikePlus_NowRunning_Screen_Calibrate_Landscape` | Known | Screen layout |
| 0x0056E754 | `NikePlus_EndWorkout_Screen_Contextual_Landscape` | Known | Screen layout |
| 0x0056E784 | `NikePlus_EndWorkout_Screen_Calibration_Contextual_Landscape` | Known | Screen layout |
| 0x0056E7C0 | `Alarms_Alarm_Clock_Triggered_Screen_Landscape` | Known | Screen layout |
| 0x0056E7EE | `Alarms_Alarm_Triggered_Screen_Landscape` | Known | Screen layout |
| 0x0056E816 | `Nike_Volume_Screen_Landscape` | Known | Screen layout |
| 0x0056E833 | `Pedometer_Volume_Screen_Landscape` | Known | Screen layout |
| 0x0056E855 | `NikePlus_NowRunning_Screen_Landscape` | Known | Screen layout |
| 0x0056E87A | `NikePlus_NowRunning_Screen_Calories_Landscape` | Known | Screen layout |
| 0x0056E98D | `RemoteUI_Screen_DisplayImage_With_Unsupported_Firewire` | Known | Screen layout |
| 0x0056E9C4 | `RemoteUI_Screen_Main_With_Unsupported_Firewire` | Known | Screen layout |
| 0x0056E9F3 | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x0056EA18 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x0056EA42 | `Camera_Screen_Active` | Known | Screen layout |
| 0x0056EA57 | `Camera_Screen_ForcedOff` | Known | Screen layout |
| 0x0056EA6F | `NikePlus_CalibrationComplete_Screen_Pacing` | Known | Screen layout |
| 0x0056EA9A | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x0056EAB8 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0056EADB | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x0056EAF7 | `RemoteUI_Hercules_ScreenLayout_Recording` | Known | Screen layout |
| 0x0056EB20 | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x0056EBA0 | `NikePlus_Calibrate_ResetToDefault_Screen_Walking` | Known | Screen layout |
| 0x0056EBD1 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x0056EBEA | `NikePlus_Calibrate_ResetToDefault_Screen_Running` | Known | Screen layout |
| 0x0056EC1B | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0056EC41 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x0056EC67 | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x0056EC80 | `CoverFlow_Screen_Exiting` | Known | Screen layout |
| 0x0056ECBA | `VoiceMemos_Screen_Saving` | Known | Screen layout |
| 0x0056ECD3 | `PhotosGL_Camera_Screen_Playing` | Known | Screen layout |
| 0x0056ECF2 | `PhotosGL_Camera_Alt_Screen_Playing` | Known | Screen layout |
| 0x0056ED15 | `PhotosGL_TvOut_Screen_Playing` | Known | Screen layout |
| 0x0056ED33 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0056ED54 | `CoverFlow_Screen_EnteringNowPlaying` | Known | Screen layout |
| 0x0056ED78 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x0056ED9C | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x0056EDB6 | `MainMenu_Main_Screen_NoContentSearch` | Known | Screen layout |
| 0x0056EDDB | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x0056EDFB | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0056EE24 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x0056EE3F | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x0056EE7A | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x0056EE9C | `MediaLists_MusicVideos_Songs_Screen_WithAlbumAndArtwork` | Known | Screen layout |
| 0x0056EED4 | `MediaLists_Songs_Screen_WithAlbumAndArtwork` | Known | Screen layout |
| 0x0056EF00 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0056EF2D | `PhotosGL_Camera_Screen_TvOut_Ask` | Known | Screen layout |
| 0x0056EF4E | `VoiceMemos_Screen_DeleteAsk` | Known | Screen layout |
| 0x0056EF6A | `VoiceMemos_Screen_DeleteAllAsk` | Known | Screen layout |
| 0x0056EF89 | `NowPlaying_Screen_TVOutAsk` | Known | Screen layout |
| 0x0056EFA4 | `NowPlaying_Screen_Initial` | Known | Screen layout |
| 0x0056EFBE | `NikePlus_EndWorkout_Screen_Contextual` | Known | Screen layout |
| 0x0056EFE4 | `NikePlus_History_Screen_Contextual` | Known | Screen layout |
| 0x0056F03C | `Camera_Screen_DiskFull` | Known | Screen layout |
| 0x0056F053 | `MediaLists_Songs_Screen_WithAlbum` | Known | Screen layout |
| 0x0056F075 | `NikePlus_History_WorkoutSummary_Screen_Gym` | Known | Screen layout |
| 0x0056F0A0 | `RemoteUI_Screen` | Known | Screen layout |
| 0x0056F0B0 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x0056F0C8 | `MediaLists_iTunesU_Screen` | Known | Screen layout |
| 0x0056F0E2 | `MediaLists_Camera_Local_Media_Screen` | Known | Screen layout |
| 0x0056F107 | `Radio_InformationalOverlay_NoAntenna_Screen` | Known | Screen layout |
| 0x0056F133 | `PhotosGL_Camera_Screen` | Known | Screen layout |
| 0x0056F14A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0056F161 | `NikePlus_Workout_Music_Screen` | Known | Screen layout |
| 0x0056F17F | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x0056F1A3 | `AddressViewer_PartialLoad_Screen` | Known | Screen layout |
| 0x0056F1C4 | `NikePlus_IsLinked_Screen` | Known | Screen layout |
| 0x0056F1DD | `NikePlus_ActivityStopped_Screen` | Known | Screen layout |
| 0x0056F1FD | `Alarms_Alarm_Clock_Triggered_Screen` | Known | Screen layout |
| 0x0056F221 | `Alarms_Alarm_Triggered_Screen` | Known | Screen layout |
| 0x0056F23F | `Pedometer_Trainer_Paused_Screen` | Known | Screen layout |
| 0x0056F25F | `Radio_InformationalOverlay_AccessoryConnected_Screen` | Known | Screen layout |
| 0x0056F294 | `Firewire_Charging_Unsupported_Screen` | Known | Screen layout |
| 0x0056F2B9 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x0056F2D7 | `LockediPod_Screen` | Known | Screen layout |
| 0x0056F2E9 | `DiskMode_Screen` | Known | Screen layout |
| 0x0056F2F9 | `DemoMode_Screen` | Known | Screen layout |
| 0x0056F309 | `Notes_Image_Screen` | Known | Screen layout |
| 0x0056F31C | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x0056F33A | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0056F350 | `NowPlaying_Idle_Screen` | Known | Screen layout |
| 0x0056F367 | `Game_Screen` | Known | Screen layout |
| 0x0056F373 | `NikePlus_Deleteme_Screen` | Known | Screen layout |
| 0x0056F38C | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x0056F3A9 | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x0056F3C2 | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x0056F3E3 | `Nike_Volume_Screen` | Known | Screen layout |
| 0x0056F3F6 | `Pedometer_Volume_Screen` | Known | Screen layout |
| 0x0056F40E | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x0056F433 | `NikePlus_NowRunning_Idle_Landscape_Screen` | Known | Screen layout |
| 0x0056F45D | `Pedometer_Main_Landscape_Screen` | Known | Screen layout |
| 0x0056F47D | `Pedometer_Ambient_Landscape_Screen` | Known | Screen layout |
| 0x0056F4A0 | `NikePlus_Daily_landscape_Screen` | Known | Screen layout |
| 0x0056F4C0 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x0056F4DD | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x0056F4FE | `PhotosRotate_Screen` | Known | Screen layout |
| 0x0056F512 | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x0056F537 | `ContextualMenu_ThreeItem_White_Screen` | Known | Screen layout |
| 0x0056F55D | `ContextualMenu_FiveItem_White_Screen` | Known | Screen layout |
| 0x0056F582 | `ContextualMenu_TwoItem_White_Screen` | Known | Screen layout |
| 0x0056F5A6 | `ContextualMenu_FourItem_White_Screen` | Known | Screen layout |
| 0x0056F5CB | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x0056F5E2 | `Calendar_Loading_Screen` | Known | Screen layout |
| 0x0056F5FA | `AddressViewer_Loading_Screen` | Known | Screen layout |
| 0x0056F617 | `Notes_Loading_Screen` | Known | Screen layout |
| 0x0056F62C | `GeniusMixes_Loading_Screen` | Known | Screen layout |
| 0x0056F647 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0056F65D | `NikePlus_SensorSearching_Screen` | Known | Screen layout |
| 0x0056F67D | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x0056F69C | `NikePlus_HeartMonitor_Linking_Screen` | Known | Screen layout |
| 0x0056F6C1 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x0056F6D9 | `NikePlus_Remote_Unlinking_Screen` | Known | Screen layout |
| 0x0056F6FA | `NikePlus_HeartMonitor_Unlinking_Screen` | Known | Screen layout |
| 0x0056F721 | `SettingsMenus_AdjustScrolling_Screen` | Known | Screen layout |
| 0x0056F746 | `Game_Running_Screen` | Known | Screen layout |
| 0x0056F75A | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x0056F775 | `Radio_NowPlaying_Screen` | Known | Screen layout |
| 0x0056F78D | `NikePlus_SimpleCalibration_Walk_Dialog_Screen` | Known | Screen layout |
| 0x0056F7BB | `NikePlus_DeleteAllWorkouts_Confirmation_Dialog_Screen` | Known | Screen layout |
| 0x0056F7F1 | `NikePlus_SimpleCalibration_Dialog_Screen` | Known | Screen layout |
| 0x0056F81A | `NikePlus_SimpleCalibration_Run_Dialog_Screen` | Known | Screen layout |
| 0x0056F847 | `Stopwatch_Screen` | Known | Screen layout |
| 0x0056F858 | `NikePlus_History_WorkoutGraph_Screen` | Known | Screen layout |
| 0x0056F87D | `SettingsMenus_Playback_Screen` | Known | Screen layout |
| 0x0056F89B | `ContextualMenu_ThreeItem_Black_Screen` | Known | Screen layout |
| 0x0056F8C1 | `ContextualMenu_FiveItem_Black_Screen` | Known | Screen layout |
| 0x0056F8E6 | `ContextualMenu_TwoItem_Black_Screen` | Known | Screen layout |
| 0x0056F90A | `ContextualMenu_FourItem_Black_Screen` | Known | Screen layout |
| 0x0056F92F | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0056F946 | `Clock_Screen` | Known | Screen layout |
| 0x0056F953 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x0056F969 | `NikePlus_HeartMonitor_LinkingInitial_Screen` | Known | Screen layout |
| 0x0056F995 | `Pedometer_Step_Goal_Screen` | Known | Screen layout |
| 0x0056F9B0 | `NikePlus_Custom_StepGoal_Screen` | Known | Screen layout |
| 0x0056F9D0 | `SettingsMenus_General_Screen` | Known | Screen layout |
| 0x0056F9ED | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x0056FA0B | `Radio_InformationalOverlay_BufferFull_Screen` | Known | Screen layout |
| 0x0056FA38 | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x0056FA54 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x0056FA65 | `PhotosZoom_Screen` | Known | Screen layout |
| 0x0056FA77 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x0056FA8E | `MediaLists_OTG_ClearConfirm_Screen` | Known | Screen layout |
| 0x0056FAB1 | `Search_Main_Screen` | Known | Screen layout |
| 0x0056FAC4 | `Location_Main_Screen` | Known | Screen layout |
| 0x0056FAD9 | `Pedometer_Main_Screen` | Known | Screen layout |
| 0x0056FAEF | `AddressViewer_Main_Screen` | Known | Screen layout |
| 0x0056FB09 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x0056FB1E | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x0056FB34 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0056FB4E | `Clock_Region_Screen` | Known | Screen layout |
| 0x0056FB62 | `RentalDeleted_Notification_Screen` | Known | Screen layout |
| 0x0056FB84 | `SingleRentalExpiring_Notification_Screen` | Known | Screen layout |
| 0x0056FBAD | `MultipleRentalsExpiring_Notification_Screen` | Known | Screen layout |
| 0x0056FBD9 | `RentalError_Notification_Screen` | Known | Screen layout |
| 0x0056FBF9 | `DeleteRental_Confirmation_Screen` | Known | Screen layout |
| 0x0056FC1A | `Stopwatch_DeleteConfirmation_Screen` | Known | Screen layout |
| 0x0056FC3E | `NikePlus_SimpleCalibration_Screen` | Known | Screen layout |
| 0x0056FC60 | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x0056FC8E | `NikePlus_StartCalibration_Screen` | Known | Screen layout |
| 0x0056FCAF | `SettingsMenus_ShakeAdjust_Duration_Screen` | Known | Screen layout |
| 0x0056FCD9 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x0056FCF7 | `Radio_ConfirmationOverlay_ChangeStation_Screen` | Known | Screen layout |
| 0x0056FD26 | `Hercules_Connection_Screen` | Known | Screen layout |
| 0x0056FD41 | `RentalInfo_Screen` | Known | Screen layout |
| 0x0056FD53 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0056FD67 | `NikePlus_Calendar_Screen` | Known | Screen layout |
| 0x0056FD80 | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x0056FD9A | `ChargingMode_LowPower_Screen` | Known | Screen layout |
| 0x0056FDB7 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0056FDD1 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0056FDEB | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0056FE05 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0056FE19 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0056FE32 | `NikePlus_CalibrationCompleteError_Screen` | Known | Screen layout |
| 0x0056FE5B | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x0056FE72 | `NikePlus_HeartMonitor_Screen` | Known | Screen layout |
| 0x0056FE8F | `Extras_Screen` | Known | Screen layout |
| 0x0056FE9D | `PhotoBrowseThumbs_Screen` | Known | Screen layout |
| 0x0056FEB6 | `Photos_Faces_Screen` | Known | Screen layout |
| 0x0056FECA | `Photos_Places_Screen` | Known | Screen layout |
| 0x0056FEDF | `MediaLists_iTunesUEpisodes_Screen` | Known | Screen layout |
| 0x0056FF01 | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x0056FF1E | `Nike_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x0056FF3A | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x0056FF5C | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x0056FF75 | `RemoteUI_Hercules_Screen` | Known | Screen layout |
| 0x0056FF8E | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x0056FFAC | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x0056FFC5 | `MediaLists_GeniusMixes_Screen` | Known | Screen layout |
| 0x0056FFE3 | `NikePlus_GeniusMixes_Screen` | Known | Screen layout |
| 0x0056FFFF | `Video_Settings_Screen` | Known | Screen layout |
| 0x00570015 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x0057002E | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x00570054 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0057006A | `MediaLists_MusicVideos_Songs_Screen` | Known | Screen layout |
| 0x0057008E | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x005700A6 | `NikePlus_Songs_Screen` | Known | Screen layout |
| 0x005700BC | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x005700DF | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x005700FC | `NikePlus_Audiobooks_Screen` | Known | Screen layout |
| 0x00570117 | `MediaLists_Rentals_Screen` | Known | Screen layout |
| 0x00570131 | `NikePlus_History_Totals_Screen` | Known | Screen layout |
| 0x00570150 | `NikePlus_History_ClearTotals_Screen` | Known | Screen layout |
| 0x00570174 | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x00570198 | `Game_Controls_Screen` | Known | Screen layout |
| 0x005701AD | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x005701C6 | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x005701E8 | `ToDo_Instructions_Screen` | Known | Screen layout |
| 0x00570201 | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x0057021D | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x00570237 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x00570258 | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x00570274 | `MediaLists_Camera_All_Videos_Screen` | Known | Screen layout |
| 0x00570298 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x005702B0 | `VoiceMemos_Screen` | Known | Screen layout |
| 0x005702C2 | `No_Photos_Screen` | Known | Screen layout |
| 0x005702D3 | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x005702ED | `AddressViewer_ContactGroups_Screen` | Known | Screen layout |
| 0x00570310 | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x0057032C | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x00570350 | `NikePlus_AudiobookChapters_Screen` | Known | Screen layout |
| 0x00570372 | `NikePlus_CalibrationCompleteSuccess_Screen` | Known | Screen layout |
| 0x0057039D | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x005703BD | `ChargingMode_Progress_Screen` | Known | Screen layout |
| 0x005703DA | `Notes_Contents_Screen` | Known | Screen layout |
| 0x005703F0 | `Photos_Events_Screen` | Known | Screen layout |
| 0x00570405 | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x00570420 | `NikePlus_Podcasts_Screen` | Known | Screen layout |
| 0x00570439 | `NikePlus_History_ClearBests_Screen` | Known | Screen layout |
| 0x0057045C | `Location_Tests_Screen` | Known | Screen layout |
| 0x00570472 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x0057048E | `NikePlus_Playlists_Screen` | Known | Screen layout |
| 0x005704A8 | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x005704CA | `NikePlus_NestedPlaylists_Screen` | Known | Screen layout |
| 0x005704EA | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x0057050B | `MediaLists_MusicVideos_Artists_Screen` | Known | Screen layout |
| 0x00570531 | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x0057054B | `NikePlus_History_Day_Workouts_Screen` | Known | Screen layout |
| 0x00570570 | `NikePlus_History_BestWorkouts_Screen` | Known | Screen layout |
| 0x00570595 | `MediaLists_Genius_Screen` | Known | Screen layout |
| 0x005705AE | `VoiceMemos_Status_Screen` | Known | Screen layout |
| 0x005705C7 | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x005705E1 | `VoiceMemos_Label_Select_Screen` | Known | Screen layout |
| 0x00570600 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0057061F | `NikePlus_NowRunning_Idle_Portrait_Screen` | Known | Screen layout |
| 0x00570648 | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x00570669 | `PhotosRotateAlt_Screen` | Known | Screen layout |
| 0x00570680 | `PhotosZoomAlt_Screen` | Known | Screen layout |
| 0x00570695 | `NikePlus_Calibrate_ResetToDefault_Screen` | Known | Screen layout |
| 0x005706BE | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x005706D6 | `VoiceMemos_No_Content_Screen` | Known | Screen layout |
| 0x005706F3 | `AddressViewer_Intro_Content_Screen` | Known | Screen layout |
| 0x00570716 | `NoContent_Screen` | Known | Screen layout |
| 0x00570727 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0057073D | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x00570753 | `NikePlus_EquipmentAlert_Screen` | Known | Screen layout |
| 0x00570772 | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x00570788 | `Notes_List_Screen` | Known | Screen layout |
| 0x0057079A | `Radio_TagList_Screen` | Known | Screen layout |
| 0x005707AF | `Radio_PresetList_Screen` | Known | Screen layout |
| 0x005707C7 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x005707DD | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x005707FE | `NikePlus_PowerPlaylist_Screen` | Known | Screen layout |
| 0x0057081C | `MediaLists_GeniusPlaylist_Screen` | Known | Screen layout |
| 0x0057083D | `PhotosGL_TvOut_Screen` | Known | Screen layout |
| 0x00570853 | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x0057086D | `NikePlus_Dynamic_Workout_Screen` | Known | Screen layout |
| 0x0057088D | `NikePlus_New_Workout_Screen` | Known | Screen layout |
| 0x005708A9 | `NikePlus_EndPausedWorkout_Screen` | Known | Screen layout |
| 0x005708CA | `NikePlus_StartWorkout_Screen` | Known | Screen layout |
| 0x005708E7 | `TODOS_Menu_Screen` | Known | Screen layout |
| 0x005708F9 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x0057090F | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0057092B | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x00570940 | `Games_Menu_Screen` | Known | Screen layout |
| 0x00570952 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x00570965 | `VoiceMemos_RecordingList_Menu_Screen` | Known | Screen layout |
| 0x0057098A | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x005709A9 | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x005709C8 | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x005709EC | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x00570A02 | `Radio_MainMenu_Screen` | Known | Screen layout |
| 0x00570A18 | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x00570A36 | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x00570A59 | `Radio_TunerTagContextMenu_Screen` | Known | Screen layout |
| 0x00570A7A | `Radio_TunerContextMenu_Screen` | Known | Screen layout |
| 0x00570A98 | `Radio_PresetListContextMenu_Screen` | Known | Screen layout |
| 0x00570ABB | `CoverFlow_Screen` | Known | Screen layout |
| 0x00570ACC | `Calendar_Day_Screen` | Known | Screen layout |
| 0x00570AE0 | `Volume_Overlay_Screen` | Known | Screen layout |
| 0x00570AF6 | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x00570B18 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x00570B30 | `Stopwatch_SessionSummary_Screen` | Known | Screen layout |
| 0x00570B50 | `NikePlus_End_WorkoutSummary_Screen` | Known | Screen layout |
| 0x00570B73 | `NikePlus_EndHercules_WorkoutSummary_Screen` | Known | Screen layout |
| 0x00570B9E | `NikePlus_History_WorkoutSummary_Screen` | Known | Screen layout |
| 0x00570BC5 | `SettingsMenu_AboutWithAccessory_Screen` | Known | Screen layout |
| 0x00570BEC | `Location_NMEA_History_Screen` | Known | Screen layout |
| 0x00570C09 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x00570C21 | `Radio_TrackHistory_Screen` | Known | Screen layout |
| 0x00570C3B | `Clock_City_Screen` | Known | Screen layout |
| 0x00570C4D | `SettingsMenus_ShakeAdjust_Intensity_Screen` | Known | Screen layout |
| 0x00570CF7 | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x00570D0C | `MediaLists_iTunesUEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00570D34 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00570D57 | `Nike_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00570D79 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00570DA1 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x00570DBF | `NikePlus_History_WorkoutSummary_Screen_Pedometer_Session` | Known | Screen layout |
| 0x00570DF8 | `NikePlus_Custom_Screen_Weight_ToPedometerSession` | Known | Screen layout |
| 0x00570F58 | `RentalInfo_Screen_ExpiringSoon` | Known | Screen layout |
| 0x00570F77 | `PhotosGL_TvOut_NTSC_Screen_Paused_Video` | Known | Screen layout |
| 0x00570F9F | `PhotosGL_TvOut_PAL_Screen_Paused_Video` | Known | Screen layout |
| 0x00570FC6 | `PhotosGL_TvOut_NTSC_Screen_Volume_Video` | Known | Screen layout |
| 0x00570FEE | `PhotosGL_TvOut_PAL_Screen_Volume_Video` | Known | Screen layout |
| 0x00571015 | `PhotosGL_TvOut_NTSC_Screen_Playing_Video` | Known | Screen layout |
| 0x0057103E | `PhotosGL_TvOut_PAL_Screen_Playing_Video` | Known | Screen layout |
| 0x00571066 | `NowPlaying_Screen_Video` | Known | Screen layout |
| 0x0057107E | `PhotosGL_TvOut_NTSC_Screen_Default_Video` | Known | Screen layout |
| 0x005710A7 | `PhotosGL_TvOut_PAL_Screen_Default_Video` | Known | Screen layout |
| 0x005710CF | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x00571125 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x00571149 | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x00571166 | `Unsupported_Screen_Radio` | Known | Screen layout |
| 0x0057117F | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x0057119B | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x005711B7 | `MainMenus_Main_Screen_Filmstrip` | Known | Screen layout |
| 0x005711D7 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x00571261 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x00571287 | `ChargingMode_ScreenLayout_LowPower` | Known | Screen layout |
| 0x005712AA | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x005712C3 | `GeniusMixes_Loading_Screen_Error` | Known | Screen layout |
| 0x005712E4 | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x00571300 | `NikePlus_ActivityStopped_Screen_Contextual_FoundSensor` | Known | Screen layout |
| 0x00571337 | `NikePlus_ActivityStopped_Screen_Contextual_NoSensor` | Known | Screen layout |
| 0x0057136B | `PhotosGL_Camera_Screen_Thumbs` | Known | Screen layout |
| 0x00571389 | `PhotosGL_Camera_Alt_Screen_Thumbs` | Known | Screen layout |
| 0x005713AB | `NikePlus_NowRunning_Screen_Calories` | Known | Screen layout |
| 0x005713CF | `NikePlus_Custom_Screen_Calories` | Known | Screen layout |
| 0x005713EF | `MainMenus_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0057140E | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00571428 | `NoContent_Screen_MainNoMovies` | Known | Screen layout |
| 0x00571446 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x00571462 | `MainMenu_Main_Screen_NoGenres` | Known | Screen layout |
| 0x00571480 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x005714BC | `MainMenu_Main_Screen_GeniusMixes` | Known | Screen layout |
| 0x005714DD | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x005714FC | `NoContent_Screen_Audiobooks` | Known | Screen layout |
| 0x00571518 | `MainMenu_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x0057153A | `NoContent_Screen_NoAudiobooks` | Known | Screen layout |
| 0x00571558 | `MainMenus_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00571578 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00571593 | `NoContent_Screen_MainNoRentals` | Known | Screen layout |
| 0x005715B2 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x005715CF | `MainMenu_Main_Screen_NoAlbums` | Known | Screen layout |
| 0x005715ED | `MainMenu_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00571611 | `Radio_TagList_Screen_Instructions` | Known | Screen layout |
| 0x00571633 | `Radio_PresetList_Screen_Instructions` | Known | Screen layout |
| 0x00571658 | `Radio_TrackHistory_Screen_Instructions` | Known | Screen layout |
| 0x0057167F | `MainMenus_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x005716A3 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x005716C2 | `NoContent_Screen_MainNoMusicVideos` | Known | Screen layout |
| 0x005716E5 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00571704 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00571725 | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x0057173F | `NoContent_Screen_MainNoVideos` | Known | Screen layout |
| 0x0057175D | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x0057177E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0057179B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x005717BA | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x005717D4 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x005717F3 | `MainMenu_Main_Screen_NoComposers` | Known | Screen layout |
| 0x0057183A | `PhotosGL_Camera_Screen_Brightness` | Known | Screen layout |
| 0x0057185C | `PhotosGL_Camera_Alt_Screen_Brightness` | Known | Screen layout |
| 0x00571882 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x005718A5 | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x005718C0 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x005718E1 | `Camera_Screen_Effects` | Known | Screen layout |
| 0x005718F7 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x00571914 | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x00571935 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00571951 | `NikePlus_NoContent_Screen_Playlists` | Known | Screen layout |
| 0x00571975 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00571996 | `MainMenus_Main_Screen_NoVideoPlaylists` | Known | Screen layout |
| 0x005719BD | `MainMenu_Main_Screen_NoArtists` | Known | Screen layout |
| 0x005719DC | `NowPlaying_Screen_Genius` | Known | Screen layout |
| 0x005719F5 | `Genius_Error_Screen_NoGenius` | Known | Screen layout |
| 0x00571A12 | `MainMenu_Main_Screen_NikePlus` | Known | Screen layout |
| 0x00571A30 | `MainMenus_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00571A50 | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x00571A6B | `NoContent_Screen_MainNoTVShows` | Known | Screen layout |
| 0x00571A8A | `NikePlus_Custom_Screen_Weight` | Known | Screen layout |
| 0x00571B9D | `Firewire_Charging_Unsupported_Screen_Alt` | Known | Screen layout |
| 0x00571BC6 | `NowPlaying_Idle_Screen_Alt` | Known | Screen layout |
| 0x00571BE1 | `Volume_Overlay_Screen_Alt` | Known | Screen layout |
| 0x00571C98 | `VoiceMemos_Screen_InsufficientDiskSpace2_Default` | Known | Screen layout |
| 0x00571CC9 | `VoiceMemos_Screen_MicrophoneRequired_Default` | Known | Screen layout |
| 0x00571CF6 | `VoiceMemos_Screen_MicrophoneDisconnected_Default` | Known | Screen layout |
| 0x00571D27 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x00571D5A | `VoiceMemos_Screen_InsufficientDiskSpace_Default` | Known | Screen layout |
| 0x00571D8A | `NowPlaying_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x00571DB7 | `VoiceMemos_Screen_Idle_Default` | Known | Screen layout |
| 0x00571DF4 | `NikePlus_EndWorkout_Screen_Calibration_Contextual_Landscape_Default` | Known | Screen layout |
| 0x00571E38 | `Nike_Volume_Screen_Landscape_Default` | Known | Screen layout |
| 0x00571E5D | `Pedometer_Volume_Screen_Landscape_Default` | Known | Screen layout |
| 0x00571EDC | `VoiceMemos_Screen_DeleteAsk_Default` | Known | Screen layout |
| 0x00571F00 | `VoiceMemos_Screen_DeleteAllAsk_Default` | Known | Screen layout |
| 0x00571F27 | `NowPlaying_Screen_TVOutAsk_Default` | Known | Screen layout |
| 0x00571F4A | `NikePlus_EndWorkout_Screen_Contextual_Default` | Known | Screen layout |
| 0x00571F78 | `NikePlus_History_Screen_Contextual_Default` | Known | Screen layout |
| 0x00571FCE | `MediaLists_Camera_Local_Media_Screen_Default` | Known | Screen layout |
| 0x00571FFB | `PhotosGL_Camera_Screen_Default` | Known | Screen layout |
| 0x0057201A | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x00572040 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0057205E | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x0057208A | `AddressViewer_PartialLoad_Screen_Default` | Known | Screen layout |
| 0x005720B3 | `NikePlus_IsLinked_Screen_Default` | Known | Screen layout |
| 0x005720D4 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x005720FC | `Alarms_Alarm_Clock_Triggered_Screen_Default` | Known | Screen layout |
| 0x00572128 | `Alarms_Alarm_Triggered_Screen_Default` | Known | Screen layout |
| 0x0057214E | `Firewire_Charging_Unsupported_Screen_Default` | Known | Screen layout |
| 0x0057217B | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x005721A1 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x005721B9 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x005721D4 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x005721F1 | `Game_Screen_Default` | Known | Screen layout |
| 0x00572205 | `NikePlus_Deleteme_Screen_Default` | Known | Screen layout |
| 0x00572226 | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x0057224C | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0057226D | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x00572296 | `Nike_Volume_Screen_Default` | Known | Screen layout |
| 0x005722B1 | `Pedometer_Volume_Screen_Default` | Known | Screen layout |
| 0x005722D1 | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x005722FB | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x00572328 | `NikePlus_Daily_landscape_Screen_Default` | Known | Screen layout |
| 0x00572350 | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x00572379 | `PhotosRotate_Screen_Default` | Known | Screen layout |
| 0x00572395 | `ContextualMenu_ThreeItem_White_Screen_Default` | Known | Screen layout |
| 0x005723C3 | `ContextualMenu_FiveItem_White_Screen_Default` | Known | Screen layout |
| 0x005723F0 | `ContextualMenu_TwoItem_White_Screen_Default` | Known | Screen layout |
| 0x0057241C | `ContextualMenu_FourItem_White_Screen_Default` | Known | Screen layout |
| 0x00572449 | `Calendar_Loading_Screen_Default` | Known | Screen layout |
| 0x00572469 | `AddressViewer_Loading_Screen_Default` | Known | Screen layout |
| 0x0057248E | `Notes_Loading_Screen_Default` | Known | Screen layout |
| 0x005724AB | `GeniusMixes_Loading_Screen_Default` | Known | Screen layout |
| 0x005724CE | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x005724EC | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x00572514 | `NikePlus_Remote_Unlinking_Screen_Default` | Known | Screen layout |
| 0x0057253D | `NikePlus_HeartMonitor_Unlinking_Screen_Default` | Known | Screen layout |
| 0x0057256C | `NikePlus_History_WorkoutGraph_Screen_Default` | Known | Screen layout |
| 0x00572599 | `ContextualMenu_ThreeItem_Black_Screen_Default` | Known | Screen layout |
| 0x005725C7 | `ContextualMenu_FiveItem_Black_Screen_Default` | Known | Screen layout |
| 0x005725F4 | `ContextualMenu_TwoItem_Black_Screen_Default` | Known | Screen layout |
| 0x00572620 | `ContextualMenu_FourItem_Black_Screen_Default` | Known | Screen layout |
| 0x0057264D | `Clock_Screen_Default` | Known | Screen layout |
| 0x00572662 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x00572680 | `NikePlus_HeartMonitor_LinkingInitial_Screen_Default` | Known | Screen layout |
| 0x005726B4 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x005726DA | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x005726FE | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x00572717 | `PhotosZoom_Screen_Default` | Known | Screen layout |
| 0x00572731 | `Location_Main_Screen_Default` | Known | Screen layout |
| 0x0057274E | `Pedometer_Main_Screen_Default` | Known | Screen layout |
| 0x0057276C | `AddressViewer_Main_Screen_Default` | Known | Screen layout |
| 0x0057278E | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x005727AB | `MainMenus_Main_Screen_Default` | Known | Screen layout |
| 0x005727C9 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x005727E6 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x00572802 | `RentalDeleted_Notification_Screen_Default` | Known | Screen layout |
| 0x0057282C | `SingleRentalExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x0057285D | `MultipleRentalsExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x00572891 | `RentalError_Notification_Screen_Default` | Known | Screen layout |
| 0x005728B9 | `DeleteRental_Confirmation_Screen_Default` | Known | Screen layout |
| 0x005728E2 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x0057290E | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x00572937 | `RentalInfo_Screen_Default` | Known | Screen layout |
| 0x00572951 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00572972 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0057298E | `NikePlus_Calendar_Screen_Default` | Known | Screen layout |
| 0x005729AF | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x005729D1 | `NikePlus_CalibrationCompleteError_Screen_Default` | Known | Screen layout |
| 0x00572A02 | `Extras_Screen_Default` | Known | Screen layout |
| 0x00572A18 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x00572A3E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00572A5F | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00572A85 | `NikePlus_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00572AA9 | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x00572AC7 | `NikePlus_Settings_Screen_Default` | Known | Screen layout |
| 0x00572AE8 | `NikePlus_Songs_Screen_Default` | Known | Screen layout |
| 0x00572B06 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00572B28 | `NikePlus_History_Totals_Screen_Default` | Known | Screen layout |
| 0x00572B4F | `NikePlus_History_ClearTotals_Screen_Default` | Known | Screen layout |
| 0x00572B7B | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x00572BA7 | `ToDo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00572BC8 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00572BEC | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x00572C0E | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x00572C32 | `MediaLists_Camera_All_Videos_Screen_Default` | Known | Screen layout |
| 0x00572C5E | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00572C7D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00572C96 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x00572CB8 | `AddressViewer_ContactGroups_Screen_Default` | Known | Screen layout |
| 0x00572CE3 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00572D07 | `NikePlus_CalibrationCompleteSuccess_Screen_Default` | Known | Screen layout |
| 0x00572D3A | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x00572D58 | `NikePlus_History_ClearBests_Screen_Default` | Known | Screen layout |
| 0x00572D83 | `Location_Tests_Screen_Default` | Known | Screen layout |
| 0x00572DA1 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00572DC5 | `NikePlus_Playlists_Screen_Default` | Known | Screen layout |
| 0x00572DE7 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x00572E11 | `NikePlus_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x00572E39 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00572E62 | `MediaLists_MusicVideos_Artists_Screen_Default` | Known | Screen layout |
| 0x00572E90 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00572EB2 | `NikePlus_History_Day_Workouts_Screen_Default` | Known | Screen layout |
| 0x00572EDF | `NikePlus_History_BestWorkouts_Screen_Default` | Known | Screen layout |
| 0x00572F0C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00572F2D | `VoiceMemos_Label_Select_Screen_Default` | Known | Screen layout |
| 0x00572F54 | `PhotosGL_Camera_Alt_Screen_Default` | Known | Screen layout |
| 0x00572F77 | `PhotosRotateAlt_Screen_Default` | Known | Screen layout |
| 0x00572F96 | `PhotosZoomAlt_Screen_Default` | Known | Screen layout |
| 0x00572FB3 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x00572FD3 | `Pedometer_Ambient_Screen_Default` | Known | Screen layout |
| 0x00572FF4 | `AddressViewer_Intro_Content_Screen_Default` | Known | Screen layout |
| 0x0057301F | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0057303D | `Nike_Dummy_BarTest_Screen_Default` | Known | Screen layout |
| 0x0057305F | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x0057307D | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x00573097 | `Radio_TagList_Screen_Default` | Known | Screen layout |
| 0x005730B4 | `Radio_PresetList_Screen_Default` | Known | Screen layout |
| 0x005730D4 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x005730F2 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x0057311B | `NikePlus_PowerPlaylist_Screen_Default` | Known | Screen layout |
| 0x00573141 | `MediaLists_GeniusPlaylist_Screen_Default` | Known | Screen layout |
| 0x0057316A | `PhotosGL_TvOut_Screen_Default` | Known | Screen layout |
| 0x00573188 | `NikePlus_New_Workout_Screen_Default` | Known | Screen layout |
| 0x005731AC | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x005731D1 | `TODOS_Menu_Screen_Default` | Known | Screen layout |
| 0x005731EB | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x00573209 | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x00573226 | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x00573240 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x0057325B | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x0057327A | `VoiceMemos_RecordingList_Menu_Screen_Default` | Known | Screen layout |
| 0x005732A7 | `Debug_MainMenu_Screen_Default` | Known | Screen layout |
| 0x005732C5 | `Radio_TunerTagContextMenu_Screen_Default` | Known | Screen layout |
| 0x005732EE | `Radio_TunerContextMenu_Screen_Default` | Known | Screen layout |
| 0x00573314 | `Radio_PresetListContextMenu_Screen_Default` | Known | Screen layout |
| 0x0057333F | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x00573358 | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x00573374 | `Volume_Overlay_Screen_Default` | Known | Screen layout |
| 0x00573392 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x005733BC | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x005733DC | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x00573404 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x0057342F | `NikePlus_EndHercules_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x00573462 | `NikePlus_History_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x00573491 | `Location_NMEA_History_Screen_Default` | Known | Screen layout |
| 0x005734B6 | `NikePlus_History_Screen_Default` | Known | Screen layout |
| 0x005734D6 | `Radio_TrackHistory_Screen_Default` | Known | Screen layout |
| 0x005734F8 | `Clock_City_Screen_Default` | Known | Screen layout |
| 0x00573736 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00573756 | `NikePlus_ActivityStopped_Screen_Contextual_FoundSensor_Default` | Known | Screen layout |
| 0x00573795 | `NikePlus_ActivityStopped_Screen_Contextual_NoSensor_Default` | Known | Screen layout |
| 0x005737D1 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x00573801 | `NikePlus_NoContent_Screen_Playlists_Default` | Known | Screen layout |
| 0x0057382D | `Firewire_Charging_Unsupported_Screen_Alt_Default` | Known | Screen layout |
| 0x0057385E | `Volume_Overlay_Screen_Alt_Default` | Known | Screen layout |
| 0x005738E9 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x00573945 | `MediaLists_OTG_ClearConfirm_ScreenLayout_Default` | Known | Screen layout |
| 0x00573976 | `RemoteUI_Hercules_ScreenLayout_Default` | Known | Screen layout |
| 0x0057399D | `Radio_MainMenu_ScreenLayout_Default` | Known | Screen layout |
| 0x005739C1 | `Pedometer_Screen_Ambient` | Known | Screen layout |
| 0x005739DA | `NikePlus_History_WorkoutSummary_Screen_Pedometer_Ambient` | Known | Screen layout |
| 0x00573A13 | `MediaLists_iTunesU_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00573A3E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00573A6A | `NikePlus_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00573A94 | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00573ABF | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x00573AE7 | `MainMenus_Main_Screen_NoiTunesUArt` | Known | Screen layout |
| 0x00573B0A | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00573B2B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00573B4C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00573B72 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00573B94 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00573BB8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00573BDC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00573C13 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00573C3A | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00573C5D | `NowPlaying_Screen_Video_TVOut` | Known | Screen layout |
| 0x00573CDA | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x00573D07 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x00573D37 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x00573DAC | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x00573DD3 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x0057400A | `NikePlus_EndPausedWorkout_Screen_BasicSave_Layout` | Known | Screen layout |
| 0x0057403C | `NikePlus_EndPausedWorkout_Screen_DistanceSave_Layout` | Known | Screen layout |
| 0x00574071 | `NikePlus_EndPausedWorkout_Screen_TimeSave_Layout` | Known | Screen layout |
| 0x005740A2 | `NikePlus_EndPausedWorkout_Screen_CaloriesSave_Layout` | Known | Screen layout |
| 0x005740D7 | `NikePlus_EndPausedWorkout_Screen_QuickstartSave_Layout` | Known | Screen layout |
| 0x0057410E | `MainMenu_Main_Screen_Pedometer_InActive_Layout` | Known | Screen layout |
| 0x0057421D | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x00574238 | `Pedometer_Main_Landscape_Screen_Medium_Layout` | Known | Screen layout |
| 0x00574266 | `Pedometer_Trainer_Paused_Screen_Layout` | Known | Screen layout |
| 0x0057428D | `Pedometer_Main_Landscape_Screen_Layout` | Known | Screen layout |
| 0x005742B4 | `Pedometer_Ambient_Landscape_Screen_Layout` | Known | Screen layout |
| 0x005742DE | `Hercules_Complete_Screen_Layout` | Known | Screen layout |
| 0x005742FE | `Pedometer_Ambient_Landscape_Medium_Screen_Layout` | Known | Screen layout |
| 0x0057432F | `Hercules_Connection_Screen_Layout` | Known | Screen layout |
| 0x005743B6 | `MainMenu_Main_Screen_Pedometer_Layout` | Known | Screen layout |
| 0x005743DC | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0057454D | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x005746DB | `Lock_Screen_Default_Layout` | Known | Screen layout |
| 0x005746F6 | `Pedometer_Step_Goal_Screen_Default_Layout` | Known | Screen layout |
| 0x0057476D | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x00574794 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x0057486D | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00574884 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x005748BC | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x005748E6 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00574961 | `Unsupported_Upgrade_ScreenLayout` | Known | Screen layout |
| 0x00574982 | `Pedometer_Main_Screen_Medium_ScreenLayout` | Known | Screen layout |
| 0x005749AC | `Pedometer_Ambient_Screen_Medium_ScreenLayout` | Known | Screen layout |
| 0x00574CF1 | `AddressViewer_Intro_Content_Screen_MainMenu` | Known | Screen layout |
| 0x00574D6B | `CoverFlow_Screen_QuickNav` | Known | Screen layout |
| 0x00574D85 | `MainMenus_Main_Screen_NoPreview` | Known | Screen layout |
| 0x00574DC0 | `PhotosGL_Camera_Screen_BatteryLow` | Known | Screen layout |
| 0x00574DE2 | `PhotosGL_Camera_Alt_Screen_BatteryLow` | Known | Screen layout |
| 0x00574E08 | `PhotosGL_TvOut_Screen_BatteryLow` | Known | Screen layout |
| 0x00574E29 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x00574E4C | `NowPlaying_Screen_Initial_From_CoverFlow` | Known | Screen layout |
| 0x00574E75 | `NowPlaying_Screen_Exit_To_CoverFlow` | Known | Screen layout |
| 0x00574E99 | `MediaLists_GeniusMixes_Screen_SingleMix` | Known | Screen layout |
| 0x00574EC1 | `RentalInfo_Screen_ExpiresToday` | Known | Screen layout |
| 0x00634A04 | `GeniusMixes_Loading_Screen_Error` | Known | Screen layout |
| 0x00634A3C | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x00634AD8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00634AF0 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00634B88 | `AddressViewer_Loading_Screen` | Known | Screen layout |
| 0x00634BA8 | `AddressViewer_Loading_Screen_Default` | Known | Screen layout |
| 0x00634BF4 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x00634C0C | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x00634C58 | `Calendar_Loading_Screen` | Known | Screen layout |
| 0x00634C70 | `Calendar_Loading_Screen_Default` | Known | Screen layout |
| 0x00634CAC | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x00634CC0 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x00634D34 | `Clock_Screen` | Known | Screen layout |
| 0x00634D44 | `Clock_Screen_Default` | Known | Screen layout |
| 0x00634DCC | `Games_Menu_Screen` | Known | Screen layout |
| 0x00634DE0 | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x00634F64 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x00634F7C | `Debug_MainMenu_Screen_Default` | Known | Screen layout |
| 0x00635014 | `Extras_Screen` | Known | Screen layout |
| 0x00635024 | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x00635048 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x006350F4 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x0063512C | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x00635148 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x00635190 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x006351A8 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x00635204 | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x00635220 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x00635278 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x006352D8 | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x006352FC | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x00635344 | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x00635364 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x006353A8 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x006353C8 | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x0063540C | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x006355B8 | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x006355DC | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x00635624 | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x00635644 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x00635734 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x0063574C | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x00635814 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x00635828 | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x0063586C | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x006358CC | `ToDo_Item_Screen` | Known | Screen layout |
| 0x006358E0 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x00635A54 | `Camera_Screen_Active` | Known | Screen layout |
| 0x00635A84 | `Camera_Screen_DiskFull` | Known | Screen layout |
| 0x00635ABC | `Camera_Screen_Effects` | Known | Screen layout |
| 0x00635AF4 | `Camera_Screen_Uninitialized` | Known | Screen layout |
| 0x00635C9C | `PhotosGL_Camera_Screen` | Known | Screen layout |
| 0x00635CB4 | `PhotosGL_Camera_Screen_Thumbs` | Known | Screen layout |
| 0x00636170 | `Clock_City_Screen` | Known | Screen layout |
| 0x00636184 | `Clock_City_Screen_Default` | Known | Screen layout |
| 0x006364BC | `Clock_Region_Screen` | Known | Screen layout |
| 0x006364D0 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x006366D0 | `CoverFlow_Screen_QuickNav` | Known | Screen layout |
| 0x00636768 | `CoverFlow_Screen_EnteringNowPlaying` | Known | Screen layout |
| 0x00636810 | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x00636864 | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x00636888 | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x006368D0 | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x006368F4 | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x0063693C | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x00636964 | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x00636A6C | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x00636A9C | `Game_Screen` | Known | Screen layout |
| 0x00636AA8 | `Game_Screen_Default` | Known | Screen layout |
| 0x00636B7C | `Game_Controls_Screen` | Known | Screen layout |
| 0x00636BFC | `Game_Running_Screen` | Known | Screen layout |
| 0x00636C2C | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00636C64 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00636C9C | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00636CD4 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00636D98 | `Location_Tests_Screen` | Known | Screen layout |
| 0x00636DB0 | `Location_Tests_Screen_Default` | Known | Screen layout |
| 0x00636DE4 | `Location_NMEA_History_Screen` | Known | Screen layout |
| 0x00636E04 | `Location_NMEA_History_Screen_Default` | Known | Screen layout |
| 0x00636E3C | `Location_Main_Screen` | Known | Screen layout |
| 0x00636F14 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0063703C | `LockediPod_Screen` | Known | Screen layout |
| 0x0063708C | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x006370D0 | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x00637114 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x00637130 | `Lock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00637174 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0063719C | `SettingsMenus_Playback_Screen` | Known | Screen layout |
| 0x0063720C | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x00637254 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00637724 | `Search_Main_Screen` | Known | Screen layout |
| 0x00637738 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x00637774 | `MainMenu_Main_Screen_NoContentSearch` | Known | Screen layout |
| 0x00637B98 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x00637C34 | `Extras_Screen_Default` | Known | Screen layout |
| 0x00637C7C | `MainMenus_Main_Screen_Filmstrip` | Known | Screen layout |
| 0x00637CD4 | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x00637D30 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x00637D88 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x00637DFC | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x00637E90 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x00637EE4 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x00637F38 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x00637FDC | `Location_Main_Screen_Default` | Known | Screen layout |
| 0x00638018 | `MainMenus_Main_Screen_NoPreview` | Known | Screen layout |
| 0x00638088 | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x006380C8 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x006380E0 | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x00638134 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x006381C4 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x00638230 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x00638308 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x006383C0 | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x006383DC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00638420 | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x0063843C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00638480 | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x006384A0 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00638504 | `MediaLists_Camera_Local_Media_Screen` | Known | Screen layout |
| 0x0063852C | `MediaLists_Camera_Local_Media_Screen_Default` | Known | Screen layout |
| 0x0063859C | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x006385B8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006385FC | `CoverFlow_Screen` | Known | Screen layout |
| 0x00638630 | `MainMenu_Main_Screen_NoAlbums` | Known | Screen layout |
| 0x00638670 | `MainMenu_Main_Screen_NoArtists` | Known | Screen layout |
| 0x006386B4 | `MainMenu_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x006386FC | `NoContent_Screen_Audiobooks` | Known | Screen layout |
| 0x00638718 | `NoContent_Screen_NoAudiobooks` | Known | Screen layout |
| 0x00638758 | `MainMenus_Main_Screen_NoCamera` | Known | Screen layout |
| 0x0063879C | `MainMenu_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006387E4 | `MainMenu_Main_Screen_NoComposers` | Known | Screen layout |
| 0x0063882C | `NoContent_Screen_Music` | Known | Screen layout |
| 0x00638844 | `NoContent_Screen_MainNoMusic` | Known | Screen layout |
| 0x00638888 | `MainMenu_Main_Screen_GeniusMixes` | Known | Screen layout |
| 0x006388CC | `MainMenu_Main_Screen_NoGenres` | Known | Screen layout |
| 0x0063890C | `MainMenus_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0063894C | `NoContent_Screen` | Known | Screen layout |
| 0x00638960 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0063899C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006389DC | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00638A1C | `MainMenus_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00638A64 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00638AA4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00638AE4 | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x00638B20 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00638B88 | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x00638BD0 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00638C0C | `MainMenus_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00638C4C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00638C88 | `MainMenus_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00638CC8 | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x00638D0C | `MainMenus_Main_Screen_NoVideoPlaylists` | Known | Screen layout |
| 0x00638D54 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00638D94 | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x00638DD0 | `MediaLists_GeniusMixes_Screen` | Known | Screen layout |
| 0x00638DF0 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00638E34 | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x00638E50 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00638E9C | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x00638EB4 | `NikePlus_Custom_Screen_Weight_ToPedometerSession` | Known | Screen layout |
| 0x00638F10 | `Pedometer_Ambient_Landscape_Screen` | Known | Screen layout |
| 0x00638F34 | `Pedometer_Ambient_Landscape_Screen_Layout` | Known | Screen layout |
| 0x00638F88 | `Pedometer_Screen_Ambient` | Known | Screen layout |
| 0x00638FA4 | `Pedometer_Ambient_Screen_Default` | Known | Screen layout |
| 0x00638FF4 | `Pedometer_Main_Landscape_Screen` | Known | Screen layout |
| 0x00639014 | `Pedometer_Main_Landscape_Screen_Layout` | Known | Screen layout |
| 0x00639064 | `Pedometer_Main_Screen` | Known | Screen layout |
| 0x0063907C | `Pedometer_Main_Screen_Default` | Known | Screen layout |
| 0x006390B8 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x00639118 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0063912C | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x006391B0 | `AddressViewer_Intro_Content_Screen` | Known | Screen layout |
| 0x006391D4 | `AddressViewer_Intro_Content_Screen_MainMenu` | Known | Screen layout |
| 0x00639220 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0063925C | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x00639278 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x006392DC | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x006392F8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00639338 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00639350 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006393AC | `MainMenu_Main_Screen_NikePlus` | Known | Screen layout |
| 0x006393EC | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00639430 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00639478 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006394C0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00639504 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00639548 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0063958C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006395A4 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x006395E8 | `MainMenus_Main_Screen_NoiTunesUArt` | Known | Screen layout |
| 0x00639630 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00639670 | `MainMenu_Main_Screen_Pedometer_Layout` | Known | Screen layout |
| 0x006396BC | `MainMenu_Main_Screen_Pedometer_InActive_Layout` | Known | Screen layout |
| 0x00639724 | `Photos_Screen` | Known | Screen layout |
| 0x00639778 | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x00639794 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00639808 | `MediaLists_Rentals_Screen` | Known | Screen layout |
| 0x00639824 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0063986C | `MediaLists_GeniusMixes_Screen_SingleMix` | Known | Screen layout |
| 0x006398B0 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x006398C8 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00639908 | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x00639924 | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00639974 | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x00639998 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006399FC | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00639A54 | `MediaLists_iTunesU_Screen` | Known | Screen layout |
| 0x00639A70 | `MediaLists_iTunesU_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00639AFC | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00639CE8 | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x00639D0C | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x00639EB0 | `MediaLists_Genius_Screen` | Known | Screen layout |
| 0x00639ECC | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00639F14 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x00639F2C | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x00639F68 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x00639FAC | `MediaLists_Songs_Screen_WithAlbum` | Known | Screen layout |
| 0x00639FF8 | `MediaLists_Songs_Screen_WithAlbumAndArtwork` | Known | Screen layout |
| 0x0063A024 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0063A0A0 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x0063A344 | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x0063A368 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0063A390 | `MediaLists_iTunesUEpisodes_Screen` | Known | Screen layout |
| 0x0063A3B4 | `MediaLists_iTunesUEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0063A400 | `MediaLists_GeniusPlaylist_Screen` | Known | Screen layout |
| 0x0063A424 | `MediaLists_GeniusPlaylist_Screen_Default` | Known | Screen layout |
| 0x0063A484 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0063A4BC | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x0063A4E0 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x0063A53C | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x0063A558 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x0063A5FC | `MediaLists_OTG_ClearConfirm_Screen` | Known | Screen layout |
| 0x0063A620 | `MediaLists_OTG_ClearConfirm_ScreenLayout_Default` | Known | Screen layout |
| 0x0063A8A0 | `Video_Settings_Screen` | Known | Screen layout |
| 0x0063A8B8 | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x0063A914 | `MediaLists_Camera_All_Videos_Screen` | Known | Screen layout |
| 0x0063A938 | `MediaLists_Camera_All_Videos_Screen_Default` | Known | Screen layout |
| 0x0063A964 | `MediaLists_MusicVideos_Artists_Screen` | Known | Screen layout |
| 0x0063A98C | `MediaLists_MusicVideos_Artists_Screen_Default` | Known | Screen layout |
| 0x0063A9EC | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x0063AA08 | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x0063AA54 | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x0063AA74 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0063AA98 | `MediaLists_MusicVideos_Songs_Screen` | Known | Screen layout |
| 0x0063AABC | `MediaLists_MusicVideos_Songs_Screen_WithAlbumAndArtwork` | Known | Screen layout |
| 0x0063AAF4 | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x0063AE54 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x0063AEB0 | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x0063AEF0 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0063AF30 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x0063AF78 | `NowPlaying_Screen_ExtraInfoLoadFailed` | Known | Screen layout |
| 0x0063AFC8 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0063B008 | `NowPlaying_Screen_Genius` | Known | Screen layout |
| 0x0063B068 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0063B0A8 | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x0063B0E8 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x0063B12C | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0063B170 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0063B1A8 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0063B1E8 | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x0063B228 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0063B268 | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x0063B2A8 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0063B5D0 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x0063B620 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x0063B670 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x0063B694 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x0063B6BC | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x0063B720 | `Pedometer_Volume_Screen` | Known | Screen layout |
| 0x0063B738 | `Pedometer_Volume_Screen_Default` | Known | Screen layout |
| 0x0063B7A4 | `Pedometer_Main_Screen_Medium_ScreenLayout` | Known | Screen layout |
| 0x0063B7D0 | `Pedometer_Volume_Screen_Landscape` | Known | Screen layout |
| 0x0063B7F4 | `Pedometer_Volume_Screen_Landscape_Default` | Known | Screen layout |
| 0x0063B820 | `Pedometer_Main_Landscape_Screen_Medium_Layout` | Known | Screen layout |
| 0x0063B850 | `Pedometer_Ambient_Screen_Medium_ScreenLayout` | Known | Screen layout |
| 0x0063B880 | `Pedometer_Ambient_Landscape_Medium_Screen_Layout` | Known | Screen layout |
| 0x0063B91C | `Notes_Contents_Screen` | Known | Screen layout |
| 0x0063B934 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x0063B974 | `Notes_Image_Screen` | Known | Screen layout |
| 0x0063B988 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x0063B9C8 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x0063BA0C | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x0063BA28 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x0063BA6C | `Notes_List_Screen` | Known | Screen layout |
| 0x0063BA80 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x0063BB08 | `PhotosGL_Camera_Alt_Screen_Thumbs` | Known | Screen layout |
| 0x0063BB50 | `PhotosGL_Screen` | Known | Screen layout |
| 0x0063BB60 | `PhotosGL_Screen_Thumbs` | Known | Screen layout |
| 0x0063BBA0 | `PhotosGL_Alt_Screen_Thumbs` | Known | Screen layout |
| 0x0063BBF4 | `Photos_Events_Screen` | Known | Screen layout |
| 0x0063BC54 | `Photos_Faces_Screen` | Known | Screen layout |
| 0x0063BCB4 | `Photos_Places_Screen` | Known | Screen layout |
| 0x0063BD00 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0063BDE0 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x0063BE38 | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x0063BE5C | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x0063BF0C | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x0063C034 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x0063C0E8 | `PhotosZoomAlt_Screen` | Known | Screen layout |
| 0x0063C100 | `PhotosZoomAlt_Screen_Default` | Known | Screen layout |
| 0x0063C138 | `PhotosGL_Screen_Default` | Known | Screen layout |
| 0x0063C170 | `PhotosZoom_Screen` | Known | Screen layout |
| 0x0063C184 | `PhotosZoom_Screen_Default` | Known | Screen layout |
| 0x0063C21C | `PhotosGL_Camera_Screen_TvOut_Ask` | Known | Screen layout |
| 0x0063C25C | `PhotosGL_Camera_Screen_Brightness` | Known | Screen layout |
| 0x0063C2A0 | `PhotosGL_Camera_Alt_Screen_Brightness` | Known | Screen layout |
| 0x0063C2E8 | `PhotosGL_Camera_Screen_TvOut_ConnectCable` | Known | Screen layout |
| 0x0063C32C | `PhotosGL_Camera_Screen_Default` | Known | Screen layout |
| 0x0063C368 | `PhotosGL_Camera_Alt_Screen_Default` | Known | Screen layout |
| 0x0063C3A4 | `PhotosGL_Camera_Screen_Paused` | Known | Screen layout |
| 0x0063C3E0 | `PhotosGL_Camera_Alt_Screen_Paused` | Known | Screen layout |
| 0x0063C41C | `PhotosGL_Camera_Screen_Playing` | Known | Screen layout |
| 0x0063C458 | `PhotosGL_Camera_Alt_Screen_Playing` | Known | Screen layout |
| 0x0063C4EC | `PhotosGL_Camera_Screen_Volume` | Known | Screen layout |
| 0x0063C528 | `PhotosGL_Camera_Alt_Screen_Volume` | Known | Screen layout |
| 0x0063C960 | `PhotosGL_TvOut_Screen_Default` | Known | Screen layout |
| 0x0063C9A4 | `PhotosGL_TvOut_NTSC_Screen_Default_Video` | Known | Screen layout |
| 0x0063C9F4 | `PhotosGL_TvOut_PAL_Screen_Default_Video` | Known | Screen layout |
| 0x0063CA1C | `PhotosGL_TvOut_Screen_Paused` | Known | Screen layout |
| 0x0063CA60 | `PhotosGL_TvOut_NTSC_Screen_Paused_Video` | Known | Screen layout |
| 0x0063CAAC | `PhotosGL_TvOut_PAL_Screen_Paused_Video` | Known | Screen layout |
| 0x0063CAD4 | `PhotosGL_TvOut_Screen_Playing` | Known | Screen layout |
| 0x0063CB18 | `PhotosGL_TvOut_NTSC_Screen_Playing_Video` | Known | Screen layout |
| 0x0063CB68 | `PhotosGL_TvOut_PAL_Screen_Playing_Video` | Known | Screen layout |
| 0x0063CB90 | `PhotosGL_TvOut_Screen_Volume` | Known | Screen layout |
| 0x0063CBD4 | `PhotosGL_TvOut_NTSC_Screen_Volume_Video` | Known | Screen layout |
| 0x0063CC20 | `PhotosGL_TvOut_PAL_Screen_Volume_Video` | Known | Screen layout |
| 0x0063CCA0 | `SlideshowAlt_Screen` | Known | Screen layout |
| 0x0063CCD0 | `Slideshow_Screen` | Known | Screen layout |
| 0x0063CD00 | `SlideshowAlt_Screen_Default` | Known | Screen layout |
| 0x0063CD34 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x0063CE38 | `Radio_TunerContextMenu_Screen` | Known | Screen layout |
| 0x0063CE58 | `Radio_TunerContextMenu_Screen_Default` | Known | Screen layout |
| 0x0063CEA4 | `Radio_TunerTagContextMenu_Screen` | Known | Screen layout |
| 0x0063CEC8 | `Radio_TunerTagContextMenu_Screen_Default` | Known | Screen layout |
| 0x0063D360 | `Radio_NowPlaying_Screen` | Known | Screen layout |
| 0x0063D418 | `Radio_PresetList_Screen` | Known | Screen layout |
| 0x0063D430 | `Radio_PresetList_Screen_Default` | Known | Screen layout |
| 0x0063D46C | `Radio_TagList_Screen` | Known | Screen layout |
| 0x0063D484 | `Radio_TagList_Screen_Default` | Known | Screen layout |
| 0x0063D4C0 | `Radio_TrackHistory_Screen` | Known | Screen layout |
| 0x0063D4DC | `Radio_TrackHistory_Screen_Default` | Known | Screen layout |
| 0x0063D51C | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x0063D5CC | `Radio_PresetListContextMenu_Screen` | Known | Screen layout |
| 0x0063D5F0 | `Radio_PresetListContextMenu_Screen_Default` | Known | Screen layout |
| 0x0063D6DC | `RemoteUI_Screen_Main_With_Unsupported_Firewire` | Known | Screen layout |
| 0x0063D730 | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x0063D770 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x0063D7D4 | `RemoteUI_Screen_DisplayImage_With_Unsupported_Firewire` | Known | Screen layout |
| 0x0063D85C | `RemoteUI_Hercules_ScreenLayout_Recording` | Known | Screen layout |
| 0x0063D8A0 | `NikePlus_History_WorkoutSummary_Screen` | Known | Screen layout |
| 0x0063D8C8 | `NikePlus_History_WorkoutSummary_Screen_Hercules` | Known | Screen layout |
| 0x0063DA84 | `SettingsMenu_AboutWithAccessory_Screen` | Known | Screen layout |
| 0x0063DAAC | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0063DB0C | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x0063DB8C | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x0063DBA8 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0063DC9C | `SettingsMenus_General_Screen` | Known | Screen layout |
| 0x0063DD00 | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x0063DD1C | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0063DD94 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x0063DDAC | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x0063DE30 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0063DE48 | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x0063DE98 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0063DED4 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x0063E0A4 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0063E1DC | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x0063E274 | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x0063E2D8 | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x0063E7C4 | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x0063E9A0 | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x0063EFEC | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x0063F03C | `Stopwatch_DeleteConfirmation_Screen` | Known | Screen layout |
| 0x0063F060 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x0063F10C | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x0063F124 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x0063F164 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x0063F1B0 | `NikePlus_EquipmentAlert_Screen` | Known | Screen layout |
| 0x0063F244 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x0063F2C8 | `NikePlus_EndPausedWorkout_Screen` | Known | Screen layout |
| 0x0063F2EC | `NikePlus_EndPausedWorkout_Screen_QuickstartSave_Layout` | Known | Screen layout |
| 0x0063F378 | `NikePlus_New_Workout_Screen` | Known | Screen layout |
| 0x0063F394 | `NikePlus_New_Workout_Screen_Default` | Known | Screen layout |
| 0x0063F3D8 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x0063F3F4 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x0063F444 | `NikePlus_NowRunning_Screen_Landscape` | Known | Screen layout |
| 0x0063F46C | `NikePlus_NowRunning_Screen_Basic_Landscape` | Known | Screen layout |
| 0x0063F4B8 | `NikePlus_SensorSearching_Screen` | Known | Screen layout |
| 0x0063F4D8 | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x0063F520 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x0063F538 | `NikePlus_History_Screen_Default` | Known | Screen layout |
| 0x0063F578 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x0063F594 | `NikePlus_Settings_Screen_Default` | Known | Screen layout |
| 0x0063F844 | `NikePlus_Workout_Music_Screen` | Known | Screen layout |
| 0x0063F864 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x0063F88C | `NikePlus_Custom_Screen_Weight_ToWorkoutMusic` | Known | Screen layout |
| 0x0063F8F0 | `NikePlus_Dynamic_Workout_Screen` | Known | Screen layout |
| 0x0063F910 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x0063F974 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x0063F9D8 | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x0063FA34 | `NikePlus_EndPausedWorkout_Screen_BasicSave_Layout` | Known | Screen layout |
| 0x0063FA98 | `NikePlus_EndPausedWorkout_Screen_CaloriesSave_Layout` | Known | Screen layout |
| 0x0063FB00 | `NikePlus_EndPausedWorkout_Screen_DistanceSave_Layout` | Known | Screen layout |
| 0x0063FB64 | `NikePlus_EndPausedWorkout_Screen_TimeSave_Layout` | Known | Screen layout |
| 0x0063FD7C | `NikePlus_StartWorkout_Screen` | Known | Screen layout |
| 0x0063FD9C | `NikePlus_StartWorkout_Screen_Resume` | Known | Screen layout |
| 0x0063FDEC | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x0063FE74 | `NikePlus_StartCalibration_Screen` | Known | Screen layout |
| 0x0063FE98 | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x0063FEC4 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x0063FEDC | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x0063FF58 | `NikePlus_EquipmentAlert_Screen_Default` | Known | Screen layout |
| 0x0063FF80 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x0063FFCC | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x00640008 | `NikePlus_HeartMonitor_Screen` | Known | Screen layout |
| 0x0064006C | `NikePlus_HeartMonitor_Linking_Screen` | Known | Screen layout |
| 0x006401D4 | `NikePlus_ActivityStopped_Screen_Contextual_FoundSensor` | Known | Screen layout |
| 0x0064020C | `NikePlus_ActivityStopped_Screen_Contextual_FoundSensor_Default` | Known | Screen layout |
| 0x0064028C | `NikePlus_ActivityStopped_Screen_Contextual_NoSensor` | Known | Screen layout |
| 0x006402C0 | `NikePlus_ActivityStopped_Screen_Contextual_NoSensor_Default` | Known | Screen layout |
| 0x00640320 | `NikePlus_End_WorkoutSummary_Screen` | Known | Screen layout |
| 0x00640344 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x0064038C | `NikePlus_HeartMonitor_LinkingInitial_Screen` | Known | Screen layout |
| 0x006403B8 | `NikePlus_HeartMonitor_LinkingInitial_Screen_Default` | Known | Screen layout |
| 0x006403EC | `NikePlus_Remote_Unlinking_Screen` | Known | Screen layout |
| 0x00640410 | `NikePlus_Remote_Unlinking_Screen_Default` | Known | Screen layout |
| 0x0064043C | `NikePlus_HeartMonitor_Unlinking_Screen` | Known | Screen layout |
| 0x00640464 | `NikePlus_HeartMonitor_Unlinking_Screen_Default` | Known | Screen layout |
| 0x00640494 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x006405FC | `NikePlus_Songs_Screen` | Known | Screen layout |
| 0x00640614 | `NikePlus_Songs_Screen_Default` | Known | Screen layout |
| 0x00640684 | `Pedometer_Step_Goal_Screen` | Known | Screen layout |
| 0x006406A0 | `Pedometer_Step_Goal_Screen_Default_Layout` | Known | Screen layout |
| 0x00640784 | `NikePlus_Custom_Screen_Weight` | Known | Screen layout |
| 0x00640C5C | `NikePlus_NoContent_Screen_Playlists` | Known | Screen layout |
| 0x00640C80 | `NikePlus_NoContent_Screen_Playlists_Default` | Known | Screen layout |
| 0x00640CD0 | `NikePlus_PowerPlaylist_Screen` | Known | Screen layout |
| 0x00640CF0 | `NikePlus_PowerPlaylist_Screen_Default` | Known | Screen layout |
| 0x00640EC8 | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x00641114 | `NikePlus_NowRunning_Idle_Portrait_Screen` | Known | Screen layout |
| 0x0064130C | `NikePlus_CalibrationCompleteError_Screen` | Known | Screen layout |
| 0x00641338 | `NikePlus_CalibrationCompleteError_Screen_Default` | Known | Screen layout |
| 0x00641394 | `NikePlus_CalibrationComplete_Screen_Pacing` | Known | Screen layout |
| 0x006413E8 | `NikePlus_CalibrationCompleteSuccess_Screen` | Known | Screen layout |
| 0x00641414 | `NikePlus_CalibrationCompleteSuccess_Screen_Default` | Known | Screen layout |
| 0x00641468 | `NikePlus_EndWorkout_Screen_Contextual` | Known | Screen layout |
| 0x00641490 | `NikePlus_EndWorkout_Screen_Contextual_Default` | Known | Screen layout |
| 0x006414E0 | `NikePlus_ActivityStopped_Screen` | Known | Screen layout |
| 0x00641500 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x00641528 | `Nike_Volume_Screen` | Known | Screen layout |
| 0x0064153C | `Nike_Volume_Screen_Default` | Known | Screen layout |
| 0x006415D0 | `NikePlus_NowRunning_Idle_Landscape_Screen` | Known | Screen layout |
| 0x00641778 | `NikePlus_EndWorkout_Screen_Calibration_Contextual_Landscape` | Known | Screen layout |
| 0x006417B4 | `NikePlus_EndWorkout_Screen_Calibration_Contextual_Landscape_Default` | Known | Screen layout |
| 0x00641820 | `NikePlus_EndWorkout_Screen_Contextual_Landscape` | Known | Screen layout |
| 0x00641850 | `NikePlus_EndWorkout_Screen_Contextual_Default_L` | Known | Screen layout |
| 0x00641880 | `Nike_Volume_Screen_Landscape` | Known | Screen layout |
| 0x006418A0 | `Nike_Volume_Screen_Landscape_Default` | Known | Screen layout |
| 0x006418E8 | `NikePlus_Audiobooks_Screen` | Known | Screen layout |
| 0x00641904 | `NikePlus_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00641954 | `NikePlus_GeniusMixes_Screen` | Known | Screen layout |
| 0x00641970 | `NikePlus_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x006419F4 | `NikePlus_Playlists_Screen` | Known | Screen layout |
| 0x00641A10 | `NikePlus_Playlists_Screen_Default` | Known | Screen layout |
| 0x00641A34 | `NikePlus_Podcasts_Screen` | Known | Screen layout |
| 0x00641A50 | `NikePlus_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00641C04 | `NikePlus_NestedPlaylists_Screen` | Known | Screen layout |
| 0x00641C24 | `NikePlus_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x00641C6C | `Nike_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x00641C88 | `Nike_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00641CAC | `NikePlus_AudiobookChapters_Screen` | Known | Screen layout |
| 0x00641CD0 | `NikePlus_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x00641D18 | `NikePlus_Calendar_Screen` | Known | Screen layout |
| 0x00641D34 | `NikePlus_Calendar_Screen_Default` | Known | Screen layout |
| 0x00641D90 | `NikePlus_History_Totals_Screen` | Known | Screen layout |
| 0x00641DB0 | `NikePlus_History_Totals_Screen_Default` | Known | Screen layout |
| 0x00641DF4 | `NikePlus_History_BestWorkouts_Screen` | Known | Screen layout |
| 0x00641E1C | `NikePlus_History_BestWorkouts_Screen_Default` | Known | Screen layout |
| 0x00641E74 | `NikePlus_History_WorkoutSummary_Screen_Gym` | Known | Screen layout |
| 0x00641ED4 | `NikePlus_History_WorkoutSummary_Screen_Pedometer_Session` | Known | Screen layout |
| 0x00641F44 | `NikePlus_History_WorkoutSummary_Screen_Pedometer_Ambient` | Known | Screen layout |
| 0x00641FA4 | `NikePlus_History_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x00641FF0 | `NikePlus_History_ClearBests_Screen` | Known | Screen layout |
| 0x00642014 | `NikePlus_History_ClearBests_Screen_Default` | Known | Screen layout |
| 0x00642054 | `NikePlus_History_ClearTotals_Screen` | Known | Screen layout |
| 0x00642078 | `NikePlus_History_ClearTotals_Screen_Default` | Known | Screen layout |
| 0x00642114 | `NikePlus_SimpleCalibration_Dialog_Screen` | Known | Screen layout |
| 0x00642140 | `NikePlus_SimpleCalibration_Run_Dialog_Screen` | Known | Screen layout |
| 0x00642194 | `NikePlus_SimpleCalibration_Walk_Dialog_Screen` | Known | Screen layout |
| 0x006421F4 | `NikePlus_History_Screen_Contextual` | Known | Screen layout |
| 0x00642218 | `NikePlus_History_Screen_Contextual_Default` | Known | Screen layout |
| 0x00642304 | `NikePlus_DeleteAllWorkouts_Confirmation_Dialog_Screen` | Known | Screen layout |
| 0x006423B0 | `NikePlus_History_Day_Workouts_Screen` | Known | Screen layout |
| 0x00642400 | `NikePlus_Daily_landscape_Screen` | Known | Screen layout |
| 0x00642420 | `NikePlus_Daily_landscape_Screen_Default` | Known | Screen layout |
| 0x006428C4 | `NikePlus_Calibrate_ResetToDefault_Screen` | Known | Screen layout |
| 0x006428F0 | `NikePlus_Calibrate_ResetToDefault_Screen_Walking` | Known | Screen layout |
| 0x00642940 | `NikePlus_Calibrate_ResetToDefault_Screen_Running` | Known | Screen layout |
| 0x006429B4 | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x006429DC | `NikePlus_NowRunning_Screen_Calibrate_Landscape` | Known | Screen layout |
| 0x00642A0C | `NikePlus_SimpleCalibration_Screen` | Known | Screen layout |
| 0x00642A30 | `NikePlus_Custom_Screen_Simple_CalibrationDistance` | Known | Screen layout |
| 0x00642AB8 | `NikePlus_IsLinked_Screen` | Known | Screen layout |
| 0x00642AD4 | `NikePlus_IsLinked_Screen_Default` | Known | Screen layout |
| 0x00642E6C | `NikePlus_Custom_StepGoal_Screen` | Known | Screen layout |
| 0x00642EC0 | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x00642ED8 | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x00642F14 | `DemoMode_Screen` | Known | Screen layout |
| 0x00642F24 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x00642F58 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x00642F70 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x00642FAC | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x00642FC4 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x00643058 | `VoiceMemos_RecordingList_Menu_Screen` | Known | Screen layout |
| 0x00643080 | `VoiceMemos_RecordingList_Menu_Screen_Default` | Known | Screen layout |
| 0x006430D4 | `VoiceMemos_No_Content_Screen` | Known | Screen layout |
| 0x0064320C | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0064324C | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x00643288 | `VoiceMemos_Status_Screen` | Known | Screen layout |
| 0x006432A4 | `VoiceMemos_Status_Screen_Inserted` | Known | Screen layout |
| 0x006432E8 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x00643308 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x00643348 | `VoiceMemos_Screen_Saving` | Known | Screen layout |
| 0x00643428 | `VoiceMemos_Screen_DeleteAllAsk` | Known | Screen layout |
| 0x00643448 | `VoiceMemos_Screen_DeleteAllAsk_Default` | Known | Screen layout |
| 0x00643488 | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x006434B0 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x0064356C | `VoiceMemos_Label_Select_Screen` | Known | Screen layout |
| 0x0064358C | `VoiceMemos_Label_Select_Screen_Default` | Known | Screen layout |
| 0x00643700 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x00643758 | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x00643774 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x00643820 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x00643860 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x006438A4 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x006438E8 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x00643924 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x0064396C | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00000628 | `Amici 1.0.2 34A20020 RTXC SCM Administrator@w02 2009/11/02 20:12:09 CL162258` | Known | RTOS |
| 0x00493014 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Known | RTOS |
| 0x0052297C | `N3ISL13TRFTuner_RTXCE` | Known | RTOS |
| 0x00522DB7 | `N3ISL20TLocationDevice_RTXCE` | Known | RTOS |

---

## 9. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00005780 | `BootTask` | Known | RTOS task thread |
| 0x00035044 | `USBDeviceTask` | Known | RTOS task thread |
| 0x00042834 | `FirewireTask` | Known | RTOS task thread |
| 0x00042848 | `TouchwheelTask` | Known | RTOS task thread |
| 0x00042870 | `DiskMgrTask` | Known | RTOS task thread |
| 0x00042880 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x00042894 | `MikeyTask` | Known | RTOS task thread |
| 0x000428A4 | `RadioTask` | Known | RTOS task thread |
| 0x0004291C | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x00042944 | `AlarmTask` | Known | RTOS task thread |
| 0x00042963 | `"USBAudioTask` | Known | RTOS task thread |
| 0x00042978 | `ChargeMgmtTask` | Known | RTOS task thread |
| 0x00049014 | `Terminator Task` | Known | RTOS task thread |
| 0x0004CC5C | `MainAppTask` | Known | RTOS task thread |
| 0x000C6C68 | `TLogPedDiskWritingTask` | Known | RTOS task thread |
| 0x000E9974 | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x00149E84 | `DiskReaderTask` | Known | RTOS task thread |
| 0x0015B074 | `MeCCABufferedRDSUpdateTask` | Known | RTOS task thread |
| 0x00184E30 | `TPodMediaPlayerFileUpdate Task` | Known | RTOS task thread |
| 0x001A4E00 | `TTrainerApp_LocaleChangedLoadingTask` | Known | RTOS task thread |
| 0x001B2184 | `GeniusMixesTask` | Known | RTOS task thread |
| 0x0029A1FC | `MeCCARecordingTask` | Known | RTOS task thread |
| 0x002D2924 | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x002D2E44 | `MeCCAInputTask` | Known | RTOS task thread |
| 0x002D2E58 | `MeCCAOutputTask` | Known | RTOS task thread |
| 0x003395AC | `InputBufferLoadTask` | Known | RTOS task thread |
| 0x00342E64 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x00351B50 | `StreamCacheMassStorageManagerTimeOutTask` | Known | RTOS task thread |
| 0x0036A9F0 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0036AA04 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0038BF84 | `Task` | Known | RTOS task thread |
| 0x0045FEC0 | `HostOSTask` | Known | RTOS task thread |

---

## 10. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0011DF10 | `Channel Reserved` | Known | Logging channel |
| 0x0011DF24 | `Channel AppBoot` | Known | Logging channel |
| 0x0011DF34 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x0011DF50 | `Channel PrefsWriting` | Known | Logging channel |
| 0x0011DF68 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x0011DF88 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x0011DFA0 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x0011DFBC | `Channel TestLogging` | Known | Logging channel |
| 0x0011DFD0 | `Channel AppFileLoading` | Known | Logging channel |
| 0x0011DFE8 | `Channel VCardReading` | Known | Logging channel |
| 0x0011E000 | `Channel LongSongScanning` | Known | Logging channel |
| 0x0011E01C | `Channel VoiceRecording` | Known | Logging channel |
| 0x0011E034 | `Channel PhotoImporting` | Known | Logging channel |
| 0x0011E04C | `Channel Notes` | Known | Logging channel |
| 0x0011E05C | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x0011E078 | `Channel DiskMode` | Known | Logging channel |
| 0x0011E08C | `Channel Firewire` | Known | Logging channel |
| 0x0011E0A0 | `Channel USB` | Known | Logging channel |
| 0x0011E0C0 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x0011E0D8 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001F05F0 | `gamedata_RW` | Known | Game system |
| 0x001F060C | `gamedata_ShareRW` | Known | Game system |
| 0x001F0620 | `games_RO` | Known | Game system |
| 0x0051C213 | `11TCGamesMenu` | Known | Game system |
| 0x0051C321 | `12TCGameScreen` | Known | Game system |
| 0x0051C679 | `14TCGameControls` | Known | Game system |
| 0x0051E62D | `27TSilverCntlrTransitionAddonI11TCGamesMenuE` | Known | Game system |
| 0x0051E710 | `27TSilverCntlrTransitionAddonI12TCGameScreenE` | Known | Game system |
| 0x0051E93E | `27TSilverCntlrTransitionAddonI14TCGameControlsE` | Known | Game system |
| 0x0056AB13 | `iPod_Control/games_RO/` | Known | Game system |
| 0x0056AB2A | `Resources/Games/games_RO/` | Known | Game system |

---

## 12. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00035688 | `AppleDRMVersion` | Known | DRM system |
| 0x00036038 | `AppleDRM` | Known | DRM system |
| 0x00036408 | `AppleVideoDRM` | Known | DRM system |
| 0x00036440 | `AppleDRM` | Known | DRM system |
| 0x00058AAC | `FairPlayDeviceType` | Known | DRM system |
| 0x001B1354 | `adrmmp4a` | Known | DRM system |
| 0x001B4084 | `drmttx3gp` | Known | DRM system |
| 0x0056C0C6 | `DRMLevel` | Known | DRM system |
| 0x0069DD44 | `$Apple FairPlay Certificate Authority0` | Known | DRM system |
| 0x0069E0C9 | `&Apple FairPlay Certification Authority0` | Known | DRM system |

---

## 13. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00058A74 | `SQLiteDB` | Known | SQLite database |
| 0x000E00B8 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000E047C | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x00186BC2 | `pGiPod_Control/iTunes/iTunes Library.itlp/Dynamic.itdb` | Known | iTunes database |
| 0x00198CE6 | `pGiPod_Control/iTunes/iTunes Library.itlp/Library.itdb` | Known | iTunes database |
| 0x001CCF1E | `pGiPod_Control/iTunes/iTunes Library.itlp/Extras.itdb` | Known | iTunes database |
| 0x001CCF56 | `pGiPod_Control/iTunes/iTunes Library.itlp/Genius.itdb` | Known | iTunes database |
| 0x001CD052 | `pGiPod_Control/iTunes/iTunes Library.itlp/Locations.itdb` | Known | iTunes database |
| 0x001CD73C | `sqlite3_extension_init` | Known | SQLite database |
| 0x001CDA1C | `sqlite_attach` | Known | SQLite database |
| 0x001CDA30 | `sqlite_detach` | Known | SQLite database |
| 0x001D7268 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x001D7280 | `iPod_Control\iTunes\firsttime` | Known | iTunes database |
| 0x001DB138 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x001DB15C | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x001EE0AC | `sqlite_temp_master` | Known | SQLite database |
| 0x001EE0C0 | `sqlite_master` | Known | SQLite database |
| 0x001F1380 | `sqlite_stat1` | Known | SQLite database |
| 0x001F1390 | `CREATE TABLE %Q.sqlite_stat1(tbl,idx,stat)` | Known | SQLite database |
| 0x001F13BC | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x00204888 | `sqlite_temp_master` | Known | SQLite database |
| 0x0020489C | `sqlite_master` | Known | SQLite database |
| 0x00204CC4 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x002051A0 | `sqlite_temp_master` | Known | SQLite database |
| 0x002051B4 | `sqlite_master` | Known | SQLite database |
| 0x00211D04 | `sqlite_temp_master` | Known | SQLite database |
| 0x00211D18 | `sqlite_master` | Known | SQLite database |
| 0x00219BA0 | `sqlite3BtreeInitPage() returns error code %d` | Known | SQLite database |
| 0x0021E2B8 | `sqlite_autoindex_` | Known | SQLite database |
| 0x0021E2CC | `sqlite_temp_master` | Known | SQLite database |
| 0x0021E2E0 | `sqlite_master` | Known | SQLite database |
| 0x00220E80 | `sqlite_temp_master` | Known | SQLite database |
| 0x00220E94 | `sqlite_master` | Known | SQLite database |
| 0x00220EA8 | `CREATE TABLE %Q.sqlite_sequence(name,seq)` | Known | SQLite database |
| 0x00224400 | `sqlite_subquery_%p_` | Known | SQLite database |
| 0x00225BC4 | `sqlite_temp_master` | Known | SQLite database |
| 0x00225BD8 | `sqlite_master` | Known | SQLite database |
| 0x00225C24 | `sqlite_sequence` | Known | SQLite database |
| 0x0022E170 | `sqlite_stat1` | Known | SQLite database |
| 0x0022E180 | `SELECT idx, stat FROM %Q.sqlite_stat1` | Known | SQLite database |
| 0x0022EFF4 | `sqlite_` | Known | SQLite database |
| 0x00241F14 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00296018 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0029EB98 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0029EBB0 | `iTunesDB` | Known | iTunes database |
| 0x002B3080 | `sqlite_temp_master` | Known | SQLite database |
| 0x002B3094 | `sqlite_master` | Known | SQLite database |
| 0x002B7210 | `%s/sqlite_` | Known | SQLite database |
| 0x002BE094 | `sqlite_attach` | Known | SQLite database |
| 0x002BE0A4 | `sqlite_detach` | Known | SQLite database |
| 0x002C45AC | `sqlite_temp_master` | Known | SQLite database |
| 0x002C45C0 | `sqlite_master` | Known | SQLite database |
| 0x002C47C4 | `sqlite_` | Known | SQLite database |
| 0x002C4804 | `sqlite_temp_master` | Known | SQLite database |
| 0x002C4818 | `sqlite_master` | Known | SQLite database |
| 0x002C482C | `sqlite_sequence` | Known | SQLite database |
| 0x002C483C | `UPDATE "%w".sqlite_sequence set name = %Q WHERE name = %Q` | Known | SQLite database |
| 0x002C4C6C | `sqlite_` | Known | SQLite database |
| 0x002C4D04 | `sqlite_temp_master` | Known | SQLite database |
| 0x002C4D18 | `sqlite_master` | Known | SQLite database |
| 0x002C5270 | `sqlite_temp_master` | Known | SQLite database |
| 0x002C5284 | `sqlite_master` | Known | SQLite database |
| 0x002C52B4 | `sqlite_stat1` | Known | SQLite database |
| 0x002C52C4 | `DELETE FROM %Q.sqlite_stat1 WHERE idx=%Q` | Known | SQLite database |
| 0x002C5568 | `sqlite_temp_master` | Known | SQLite database |
| 0x002C557C | `sqlite_master` | Known | SQLite database |
| 0x002C55EC | `DELETE FROM %s.sqlite_sequence WHERE name=%Q` | Known | SQLite database |
| 0x002C5654 | `sqlite_stat1` | Known | SQLite database |
| 0x002C5664 | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x002C5B7C | `sqlite_temp_master` | Known | SQLite database |
| 0x002C5B90 | `sqlite_master` | Known | SQLite database |
| 0x002C76A4 | `sqlite_temp_master` | Known | SQLite database |
| 0x002C76B8 | `sqlite_master` | Known | SQLite database |
| 0x00475730 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00475770 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0047612F | `SQLite format 3` | Known | SQLite database |
| 0x004787DC | `CREATE TABLE sqlite_master(` | Known | SQLite database |
| 0x00478844 | `CREATE TEMP TABLE sqlite_temp_master(` | Known | SQLite database |
| 0x00478F0C | `illegal return value (%d) from the authorization function - should be SQLITE_OK,` | Known | SQLite database |
| 0x00478FC4 | `SELECT 'CREATE TABLE vacuum_db.' || substr(sql,14)   FROM sqlite_master WHERE ty` | Known | SQLite database |
| 0x0047904C | `SELECT 'CREATE INDEX vacuum_db.' || substr(sql,14)  FROM sqlite_master WHERE sql` | Known | SQLite database |
| 0x004790B4 | `SELECT 'CREATE UNIQUE INDEX vacuum_db.' || substr(sql,21)   FROM sqlite_master W` | Known | SQLite database |
| 0x0047912C | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x004791DC | `SELECT 'DELETE FROM vacuum_db.' || quote(name) || ';' FROM vacuum_db.sqlite_mast` | Known | SQLite database |
| 0x00479250 | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x004792E8 | `INSERT INTO vacuum_db.sqlite_master   SELECT type, name, tbl_name, rootpage, sql` | Known | SQLite database |
| 0x004794A8 | `UPDATE %Q.%s SET sql = CASE WHEN type = 'trigger' THEN sqlite_rename_trigger(sql` | Known | SQLite database |
| 0x0047961C | `UPDATE sqlite_temp_master SET sql = sqlite_rename_trigger(sql, %Q), tbl_name = %` | Known | SQLite database |
| 0x00479830 | `sqlite3_get_table() called with two or more incompatible queries` | Known | SQLite database |
| 0x0056C4D8 | `sqlite_rename_table` | Known | SQLite database |
| 0x0056C657 | `sqlite_version` | Known | SQLite database |
| 0x0056C6F1 | `sqlite_rename_trigger` | Known | SQLite database |
| 0x0057538E | `SQLite_iPod_VFS` | Known | SQLite database |
| 0x0057BBFD | `CREATE TABLE _SqliteDatabaseProperties (key TEXT, value TEXT, UNIQUE(key));` | Known | SQLite database |
| 0x0062E5CC | `Richard Hipp (SQLite) SQLite Copyright` | Known | SQLite database |
| 0x0062E5F4 | `All of the deliverable code in SQLite has been dedicated to the public domain by` | Known | SQLite database |
| 0x0062E820 | `The previous paragraph applies to the deliverable code in SQLite - those parts o` | Known | SQLite database |
| 0x0062E9F0 | `All of the deliverable code in SQLite has been written from scratch. No code has` | Known | SQLite database |
| 0x0062EB5C | `Obtaining An Explicit License To Use SQLite` | Known | SQLite database |
| 0x0062EB88 | `Even though SQLite is in the public domain and does not require a license, some ` | Known | SQLite database |
| 0x0062EC80 | `-You are using SQLite in a jurisdiction that does not recognize the right of an ` | Known | SQLite database |
| 0x0062ED04 | `-You want to hold a tangible legal document as evidence that you have the legal ` | Known | SQLite database |
| 0x0062EDC0 | `If you feel like you really have to purchase a license for SQLite, Hwaci, the co` | Known | SQLite database |
| 0x0062EE80 | `In order to keep SQLite completely free and unencumbered by copyright, all new c` | Known | SQLite database |
| 0x0062F13C | `We are not able to accept patches or changes to SQLite that are not accompanied ` | Known | SQLite database |

---

## 14. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0004D858 | `[FTL:MSG] Apple NAND Driver (AND) RW` | Known | Hardware |
| 0x0004D8D0 | `[FTL:MSG] No NAND attached` | Known | Hardware |
| 0x000589D0 | `FireWireGUID` | Known | FireWire |
| 0x000589E0 | `FireWireVersion` | Known | FireWire |
| 0x00058C4C | `FireWire` | Known | FireWire |
| 0x00060824 | `[FIL:INF] could not find NAND config in the new NAND tables` | Known | Hardware |
| 0x000635FC | `NANDDRIVERSIGN` | Known | Hardware |
| 0x00063844 | `NANDDRIVERSIGN` | Known | Hardware |
| 0x001DBF68 | `NANDDRIVERSIGN` | Known | Hardware |
| 0x0036A958 | `NAND FLASH DRIVE` | Known | Hardware |
| 0x00492FD8 | `[NAND] Panic! %s:%d` | Known | Hardware |
| 0x00492FF0 | `[NAND] Failed WMR_ASSERT(%s) %s:%d` | Known | Hardware |
| 0x00493084 | `[NAND] %s:%d IOCtl on buffer of size %d with %d bytes of src data!` | Known | Hardware |
| 0x0049328C | `[WMR:ERR] NAND format invalid (mismatch, corrupt, read error or blank NAND devic` | Known | Hardware |
| 0x0049337C | `AND: NAND initialisation failed due to format mismatch or uninitialised NAND.` | Known | Hardware |
| 0x004C2160 | `[FTL:WRN] Recovering NAND Data Structures - this will take some time!` | Known | Hardware |
| 0x004C31A0 | `(bReadEdoClocks * dwMaxNSPerClock) < (_GetReadValidNanosecs() + SOC_RISE_TIME_NS` | Known | Hardware |
| 0x0051D69A | `21TCFirewireUnsupported` | Known | FireWire |
| 0x0051F2BF | `27TSilverCntlrTransitionAddonI21TCFirewireUnsupportedE` | Known | FireWire |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0051DB3C | `23TCSettings_RadioRegions` | Known | FM Radio |
| 0x0051DED4 | `24TSilverRadioTunerBarView` | Known | FM Radio |
| 0x0051F674 | `27TSilverCntlrTransitionAddonI23TCSettings_RadioRegionsE` | Known | FM Radio |
| 0x0056A955 | `General.RadioRegion` | Known | FM Radio |
| 0x00571B51 | `Radio_ConfirmationOverlay_ChangeStation_Layout_Portrait` | Known | FM Radio |
| 0x00574376 | `Radio_NowPlaying_TunerBar_Layout` | Known | FM Radio |
| 0x0057482E | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |
| 0x0062C7A4 | `Please use the built in FM tuner to listen to the Radio.` | Known | FM Radio |
| 0x0062CA18 | `Radio Regions` | Known | FM Radio |
| 0x0063D068 | `Radio_NowPlaying_TunerBar_Layout` | Known | FM Radio |
| 0x0063D540 | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 16. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0005707C | `TPhotosGLCntlrLcdCamera` | Known | Camera |
| 0x0008F184 | `CameraShutter` | Known | Camera |
| 0x000B2B20 | `EnterCameraScreen` | Known | Camera |
| 0x000B3010 | `EnterCameraScreen` | Known | Camera |
| 0x000C8C70 | `PushScreen_BrowseCameraPhotos` | Known | Camera |
| 0x000EB7AC | `Camera Videos` | Known | Camera |
| 0x000F6BF8 | `FinishedPopCamera` | Known | Camera |
| 0x000F6F58 | `PopCamera` | Known | Camera |
| 0x000F70E0 | `PopCamera` | Known | Camera |
| 0x000F7260 | `PopCamera` | Known | Camera |
| 0x000F75B4 | `PopCamera` | Known | Camera |
| 0x000F782C | `PopCamera` | Known | Camera |
| 0x000F79BC | `PopCamera` | Known | Camera |
| 0x000F7A44 | `PopCamera` | Known | Camera |
| 0x000FA7C4 | `CameraApp` | Known | Camera |
| 0x0010B4FC | `EmptyCameraHilited` | Known | Camera |
| 0x0010B510 | `CameraHilited` | Known | Camera |
| 0x0010B940 | `CameraSelected` | Known | Camera |
| 0x0010CA6C | `CameraVideosSelected` | Known | Camera |
| 0x0012E51C | `CameraDeviceManager` | Known | Camera |
| 0x00156B34 | `%d: GKCameraDriver - camera overflow occurred!` | Known | Camera |
| 0x0047BF44 | `TPhotosGLCntlrLcdCamera` | Known | Camera |
| 0x0051C122 | `10TCameraApp` | Known | Camera |
| 0x0051C170 | `10TPodCamera` | Known | Camera |
| 0x0051C657 | `14GKCameraDevice` | Known | Camera |
| 0x0051C668 | `14GKCameraDriver` | Known | Camera |
| 0x0051C870 | `15TCCameraInitial` | Known | Camera |
| 0x0051CE4B | `17TCameraMediaModel` | Known | Camera |
| 0x0051D43D | `20TCCameraDeleteDialog` | Known | Camera |
| 0x0051D682 | `21TCCameraAllVideosList` | Known | Camera |
| 0x0051D8C7 | `22TCCameraLocalMediaList` | Known | Camera |
| 0x0051D8E0 | `22TCCameraMediaList_Base` | Known | Camera |
| 0x0051DB08 | `23TCCameraDeleteAllDialog` | Known | Camera |
| 0x0051DBF2 | `23TPhotosGLCntlrLcdCamera` | Known | Camera |
| 0x0051DC5A | `23TRecentCameraMediaModel` | Known | Camera |
| 0x0051DE17 | `24TCameraMediaModel_Import` | Known | Camera |
| 0x0051DE9E | `24TSilverCameraShutterView` | Known | Camera |
| 0x0051DFFF | `25TCameraApplication_Import` | Known | Camera |
| 0x0051E234 | `26TCameraMediaDatabaseEntity` | Known | Camera |
| 0x0051E49B | `27TCameraMediaDatabaseContext` | Known | Camera |
| 0x0051E9CE | `27TSilverCntlrTransitionAddonI15TCCameraInitialE` | Known | Camera |
| 0x0051F035 | `27TSilverCntlrTransitionAddonI20TCCameraDeleteDialogE` | Known | Camera |
| 0x0051F288 | `27TSilverCntlrTransitionAddonI21TCCameraAllVideosListE` | Known | Camera |
| 0x0051F441 | `27TSilverCntlrTransitionAddonI22TCCameraLocalMediaListE` | Known | Camera |
| 0x0051F602 | `27TSilverCntlrTransitionAddonI23TCCameraDeleteAllDialogE` | Known | Camera |
| 0x0051F758 | `27TSilverCntlrTransitionAddonI23TPhotosGLCntlrLcdCameraE` | Known | Camera |
| 0x00521300 | `27TSilverCntlrTransitionAddonI8TCCameraE` | Known | Camera |
| 0x00521666 | `29TCameraMediaDatabaseInterface` | Known | Camera |
| 0x0052198C | `30TRecentCameraMediaModel_Import` | Known | Camera |
| 0x00522416 | `8TCCamera` | Known | Camera |
| 0x00522521 | `N10TCameraApp19TCameraStateMachineE` | Known | Camera |
| 0x005225B2 | `N17TCameraMediaModel15CameraItemPropsE` | Known | Camera |
| 0x005225D9 | `N17TCameraMediaModel17LaunchCameraPropsE` | Known | Camera |
| 0x00522602 | `N17TCameraMediaModel18SyncedContentPropsE` | Known | Camera |
| 0x0052262C | `N17TCameraMediaModel22UnSyncedPhotoListPropsE` | Known | Camera |
| 0x0052265A | `N17TCameraMediaModel22UnSyncedVideoFilePropsE` | Known | Camera |
| 0x00522688 | `N17TCameraMediaModel22UnSyncedVideoListPropsE` | Known | Camera |
| 0x005226B6 | `N17TCameraMediaModel23UnSyncedAllContentPropsE` | Known | Camera |
| 0x0052270D | `N24TSilverCameraShutterView25TShutterAnimationStateMsgE` | Known | Camera |
| 0x00522745 | `N27TCameraMediaDatabaseContext16ContextualEntityE` | Known | Camera |
| 0x0052278E | `N3ISL10IPodCameraE` | Known | Camera |
| 0x0056AC05 | `cameraremote` | Known | Camera |
| 0x0056D718 | `PopCamera` | Known | Camera |
| 0x00570E29 | `PhotosGL_Camera_Delete_All_Confirmation` | Known | Camera |
| 0x00570E51 | `MediaLists_Camera_Delete_All_Confirmation` | Known | Camera |
| 0x00570E7B | `PhotosGL_Camera_Delete_Item_Confirmation` | Known | Camera |
| 0x00570EA4 | `MediaLists_Camera_Delete_Video_Confirmation` | Known | Camera |
| 0x00570ED0 | `MediaLists_Camera_Delete_Video_Event_Confirmation` | Known | Camera |
| 0x00570F02 | `MediaLists_Camera_Delete_Photo_Event_Confirmation` | Known | Camera |
| 0x005711F9 | `PhotosGL_Camera_All_Media_Delete_Menu_Alt_NoStatusBar` | Known | Camera |
| 0x0057122F | `PhotosGL_Camera_All_Media_Delete_Menu_NoStatusBar` | Known | Camera |
| 0x00571BFB | `PhotosGL_Camera_Delete_All_Confirmation_Alt` | Known | Camera |
| 0x00571C27 | `PhotosGL_Camera_Delete_Item_Confirmation_Alt` | Known | Camera |
| 0x00571C54 | `PhotosGL_Camera_All_Media_Delete_Menu_Alt` | Known | Camera |
| 0x005735CF | `PhotosGL_Camera_Delete_All_Confirmation_Default` | Known | Camera |
| 0x005735FF | `MediaLists_Camera_Delete_All_Confirmation_Default` | Known | Camera |
| 0x00573631 | `PhotosGL_Camera_Delete_Item_Confirmation_Default` | Known | Camera |
| 0x00573662 | `MediaLists_Camera_Delete_Video_Confirmation_Default` | Known | Camera |
| 0x00573696 | `MediaLists_Camera_Delete_Video_Event_Confirmation_Default` | Known | Camera |
| 0x005736D0 | `MediaLists_Camera_Delete_Photo_Event_Confirmation_Default` | Known | Camera |
| 0x00573880 | `PhotosGL_Camera_Delete_All_Confirmation_Alt_Default` | Known | Camera |
| 0x005738B4 | `PhotosGL_Camera_Delete_Item_Confirmation_Alt_Default` | Known | Camera |
| 0x00574B5D | `PhotosGL_Camera_All_Media_Delete_Menu` | Known | Camera |
| 0x00574B83 | `MediaLists_Camera_All_Media_Delete_Menu` | Known | Camera |
| 0x00574BC0 | `PhotosGL_Camera_All_Media_Contextual_Menu` | Known | Camera |
| 0x00574BEA | `MediaLists_Camera_All_Media_Contextual_Menu` | Known | Camera |
| 0x00574C16 | `PhotosGL_Camera_All_Media_Alt_Contextual_Menu` | Known | Camera |
| 0x00574C44 | `MediaLists_Camera_All_Media_Delete_All_Menu` | Known | Camera |
| 0x005BDE68 | `TPhotosGLCntlrLcdCamera` | Known | Camera |
| 0x0062851C | `Video Camera` | Known | Camera |
| 0x00629BAC | `Camera Roll` | Known | Camera |
| 0x00629BB8 | `Camera Videos` | Known | Camera |
| 0x00629BC8 | `Video Camera` | Known | Camera |
| 0x00629C50 | `Delete all camera videos from your iPod?` | Known | Camera |
| 0x00629CF0 | `Camera Initializing` | Known | Camera |
| 0x0062AEA8 | `Video Camera` | Known | Camera |
| 0x0062AF2C | `Camera Videos` | Known | Camera |
| 0x0062BFD4 | `Camera Roll` | Known | Camera |
| 0x0062C208 | `Delete all camera videos from your iPod?` | Known | Camera |
| 0x0062C234 | `Delete this camera video from your iPod?` | Known | Camera |
| 0x006339D8 | `cameraremote.action.up` | Known | Camera |
| 0x00633A04 | `cameraremote.photo.up` | Known | Camera |
| 0x00633A1C | `cameraremote.video.up` | Known | Camera |
| 0x0063592C | `PopCamera` | Known | Camera |
| 0x006359F4 | `controller.FinishedPopCamera` | Known | Camera |
| 0x00635C4C | `controller.EnterCameraScreen` | Known | Camera |
| 0x00635D00 | `MediaLists_Camera_Delete_All_Confirmation` | Known | Camera |
| 0x00635D2C | `MediaLists_Camera_Delete_All_Confirmation_Default` | Known | Camera |
| 0x00635DA0 | `MediaLists_Camera_All_Media_Contextual_Menu` | Known | Camera |
| 0x00635DCC | `MediaLists_Camera_All_Media_Delete_All_Menu` | Known | Camera |
| 0x00635E14 | `MediaLists_Camera_All_Media_Delete_Menu` | Known | Camera |
| 0x00635E70 | `MediaLists_Camera_Delete_Photo_Event_Confirmation` | Known | Camera |
| 0x00635EA4 | `MediaLists_Camera_Delete_Photo_Event_Confirmation_Default` | Known | Camera |
| 0x00635F10 | `MediaLists_Camera_Delete_Video_Confirmation` | Known | Camera |
| 0x00635F3C | `MediaLists_Camera_Delete_Video_Confirmation_Default` | Known | Camera |
| 0x00635FA4 | `MediaLists_Camera_Delete_Video_Event_Confirmation` | Known | Camera |
| 0x00635FD8 | `MediaLists_Camera_Delete_Video_Event_Confirmation_Default` | Known | Camera |
| 0x006384CC | `controller.CameraHilited` | Known | Camera |
| 0x006384E8 | `controller.CameraSelected` | Known | Camera |
| 0x00638738 | `controller.EmptyCameraHilited` | Known | Camera |
| 0x0063A8F4 | `controller.CameraVideosSelected` | Known | Camera |
| 0x0063BAB0 | `controller.PushScreen_BrowseCameraPhotos` | Known | Camera |
| 0x0063BADC | `controller.PushScreen_BrowseCameraPhotosAlt` | Known | Camera |
| 0x0063C564 | `PhotosGL_Camera_All_Media_Contextual_Menu` | Known | Camera |
| 0x0063C590 | `PhotosGL_Camera_All_Media_Delete_Menu` | Known | Camera |
| 0x0063C5B8 | `PhotosGL_Camera_Delete_All_Confirmation` | Known | Camera |
| 0x0063C5E0 | `PhotosGL_Camera_Delete_All_Confirmation_Default` | Known | Camera |
| 0x0063C638 | `PhotosGL_Camera_Delete_Item_Confirmation` | Known | Camera |
| 0x0063C664 | `PhotosGL_Camera_Delete_Item_Confirmation_Default` | Known | Camera |
| 0x0063C6BC | `PhotosGL_Camera_All_Media_Delete_Menu_NoStatusBar` | Known | Camera |
| 0x0063C75C | `PhotosGL_Camera_All_Media_Alt_Contextual_Menu` | Known | Camera |
| 0x0063C78C | `PhotosGL_Camera_All_Media_Delete_Menu_Alt` | Known | Camera |
| 0x0063C7B8 | `PhotosGL_Camera_Delete_All_Confirmation_Alt` | Known | Camera |
| 0x0063C7E4 | `PhotosGL_Camera_Delete_All_Confirmation_Alt_Default` | Known | Camera |
| 0x0063C818 | `PhotosGL_Camera_Delete_Item_Confirmation_Alt` | Known | Camera |
| 0x0063C848 | `PhotosGL_Camera_Delete_Item_Confirmation_Alt_Default` | Known | Camera |
| 0x0063C880 | `PhotosGL_Camera_All_Media_Delete_Menu_Alt_NoStatusBar` | Known | Camera |

---

## 17. Pedometer

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00038F90 | `Pedometer` | Known | Pedometer |
| 0x00058460 | `TTrainer_Cntlr_Ambient_Pedometer` | Known | Pedometer |
| 0x000A5CE4 | `] Step: ` | Known | Pedometer |
| 0x0010C0B8 | `PedometerHilited` | Known | Pedometer |
| 0x0010C0CC | `PedometerInactiveHilited` | Known | Pedometer |
| 0x0011D208 | `/Pedometer/` | Known | Pedometer |
| 0x0018AE18 | `TTrainer_Cntlr_Pedometer_Goal` | Known | Pedometer |
| 0x0018AE40 | `TPedometer_Hourly_Cntlr` | Known | Pedometer |
| 0x001B2C08 | `GoToWorkoutPedometerAmbientSummaryScreen` | Known | Pedometer |
| 0x001B2C34 | `GoToWorkoutPedometerSummaryScreen` | Known | Pedometer |
| 0x001B322C | `GoToWorkoutPedometerAmbientSummaryScreen` | Known | Pedometer |
| 0x001B3258 | `GoToPedometerSessionWorkoutSummaryScreen` | Known | Pedometer |
| 0x002406C8 | `pedometer` | Known | Pedometer |
| 0x00252580 | `pedometer` | Known | Pedometer |
| 0x0029F77C | `TPedometerThread` | Known | Pedometer |
| 0x002CADA4 | `TPedometerHeartbeatThread` | Known | Pedometer |
| 0x002D76C0 | `GoToPedometerSession` | Known | Pedometer |
| 0x002D76D8 | `GoToPedometerDaily` | Known | Pedometer |
| 0x002D8DF4 | `/Pedometer/` | Known | Pedometer |
| 0x002DFDF0 | `pedometer` | Known | Pedometer |
| 0x002F87B4 | `pedometer` | Known | Pedometer |
| 0x0045F7C7 | `TotalSteps` | Known | Pedometer |
| 0x0046365C | `PedometerModel - No steps for ambient workout . Discarding and deleting session!` | Known | Pedometer |
| 0x004636B0 | `PedometerModel - No steps for session workout . Discarding and deleting session!` | Known | Pedometer |
| 0x00463E6C | `TPedometer_Hourly_Cntlr` | Known | Pedometer |
| 0x00463F30 | `TTrainer_Cntlr_Pedometer_Goal` | Known | Pedometer |
| 0x0047B87C | `TTrainer_Cntlr_Ambient_Pedometer` | Known | Pedometer |
| 0x0047B8A0 | `TTrainer_Cntlr_Pedometer` | Known | Pedometer |
| 0x0047B8D8 | `TTrainer_Cntlr_Ambient_Pedometer` | Known | Pedometer |
| 0x0047F880 | `Stepper` | Known | Pedometer |
| 0x0047F9F3 | `Steps` | Known | Pedometer |
| 0x0047FA87 | `pedometer` | Known | Pedometer |
| 0x0047FA91 | `ambient_pedometer` | Known | Pedometer |
| 0x0051C5F7 | `13TPedometerApp` | Known | Pedometer |
| 0x0051C97A | `15TPedometerModel` | Known | Pedometer |
| 0x0051D9A8 | `22TPedometerModel_Import` | Known | Pedometer |
| 0x0051DBD8 | `23TPedometer_Hourly_Cntlr` | Known | Pedometer |
| 0x0051DCC2 | `23TSilverStepBarGraphView` | Known | Pedometer |
| 0x0051DF25 | `24TTrainer_Cntlr_Pedometer` | Known | Pedometer |
| 0x0051F71F | `27TSilverCntlrTransitionAddonI23TPedometer_Hourly_CntlrE` | Known | Pedometer |
| 0x0052053F | `27TSilverCntlrTransitionAddonI29TTrainer_Cntlr_Pedometer_GoalE` | Known | Pedometer |
| 0x00520C50 | `27TSilverCntlrTransitionAddonI32TTrainer_Cntlr_Ambient_PedometerE` | Known | Pedometer |
| 0x005213FB | `27TTrainer_PedometerGoalModel` | Known | Pedometer |
| 0x00521806 | `29TTrainer_Cntlr_Pedometer_Goal` | Known | Pedometer |
| 0x00521E29 | `32TTrainer_Cntlr_Ambient_Pedometer` | Known | Pedometer |
| 0x00522071 | `34TTrainer_PedometerGoalModel_Import` | Known | Pedometer |
| 0x005229D6 | `N3ISL14TStepPedometerE` | Known | Pedometer |
| 0x00522B79 | `N3ISL17IPodStepPedometerE` | Known | Pedometer |
| 0x0056BD43 | `Trainer.PedometerStepGoal` | Known | Pedometer |
| 0x0056BD71 | `Trainer.Pedometer` | Known | Pedometer |
| 0x0056C692 | `AggStep` | Known | Pedometer |
| 0x005BDC54 | `TTrainer_Cntlr_Ambient_Pedometer` | Known | Pedometer |
| 0x005BDC78 | `TTrainer_Cntlr_Ambient_Pedometer` | Known | Pedometer |
| 0x005BEAEC | `TPedometer_Hourly_Cntlr` | Known | Pedometer |
| 0x005BEB04 | `TTrainer_Cntlr_Pedometer_Goal` | Known | Pedometer |
| 0x00628634 | `Pedometer` | Known | Pedometer |
| 0x0062A880 | `Pedometer` | Known | Pedometer |
| 0x0062A8A0 | `Please quit your Nike+ workout to begin using the Pedometer.` | Known | Pedometer |
| 0x0062B978 | `Pedometer` | Known | Pedometer |
| 0x0062BA00 | `Step Workout` | Known | Pedometer |
| 0x00631770 | `Stepper Workout` | Known | Pedometer |
| 0x0063226C | `Pedometer` | Known | Pedometer |
| 0x006325F0 | `Steps` | Known | Pedometer |
| 0x006329DC | `Step away from all other sensors` | Known | Pedometer |
| 0x00632B28 | `Step away from all other remotes` | Known | Pedometer |
| 0x00632DE4 | `Step away from all other monitors.` | Known | Pedometer |
| 0x00632E1C | `Daily Step View` | Known | Pedometer |
| 0x00632E38 | `Total Steps:` | Known | Pedometer |
| 0x00632E74 | `Steps` | Known | Pedometer |
| 0x00632E88 | `Step Goal` | Known | Pedometer |
| 0x00632EC4 | `Daily Step Goal` | Known | Pedometer |
| 0x00632EE4 | `Weekly Step Total` | Known | Pedometer |
| 0x00632EF8 | `Monthly Step Total` | Known | Pedometer |
| 0x00638EE8 | `controller.GoToPedometerDailyLandscape` | Known | Pedometer |
| 0x00638F60 | `controller.GoToPedometerDailyPortrait` | Known | Pedometer |
| 0x00638FC8 | `controller.GoToPedometerSessionLandscape` | Known | Pedometer |
| 0x0063903C | `controller.GoToPedometerSessionPortrait` | Known | Pedometer |
| 0x00639654 | `controller.PedometerHilited` | Known | Pedometer |
| 0x00639698 | `controller.PedometerInactiveHilited` | Known | Pedometer |
| 0x0063B77C | `controller.GotoMediumPedometerLayout` | Known | Pedometer |
| 0x00641EA0 | `controller.GoToPedometerSessionWorkoutSummaryScreen` | Known | Pedometer |
| 0x00641F10 | `controller.GoToWorkoutPedometerAmbientSummaryScreen` | Known | Pedometer |
| 0x006423D8 | `controller.Goto_Pedometer_Daily_Graph` | Known | Pedometer |
| 0x00642E1C | `controller.GoToWorkoutPedometerSummaryScreen` | Known | Pedometer |
| 0x00642E4C | `controller.GotoCustomStepGoal` | Known | Pedometer |

---

## 18. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00054C4C | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x00061134 | `Resources/Sounds/clicker.wav` | Filesystem Path |  |
| 0x0006C918 | `iPod_Control/Device/Radio/RadioBuffer` | Filesystem Path |  |
| 0x00078520 | `iPod_Control/Device/Accessories/Tags/` | Filesystem Path |  |
| 0x00084724 | `iPod_Control\Device` | Filesystem Path |  |
| 0x00086498 | `Resources/Sounds/camera.wav` | Filesystem Path |  |
| 0x00088520 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x000A5CA0 | `iPod_Control/Device/Radio/Tuner_Scan.log` | Filesystem Path |  |
| 0x000DF608 | `iPod_Control` | Filesystem Path |  |
| 0x000E0970 | `iPod_Control/Logs/crash000.bin` | Filesystem Path |  |
| 0x000E09A4 | `pytcgsmlrddamfducpafksthpsafpytegerfktsfglveiPod_Control/Logs` | Filesystem Path |  |
| 0x000E0F48 | `iPod_Control\Device\dst` | Filesystem Path |  |
| 0x00104040 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001041C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104234 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001042B0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010445C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001044CC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010453C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001045AC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010461C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010468C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001046FC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010476C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001047DC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104854 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001048C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010493C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001049B4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104A24 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104A94 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104B14 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104B84 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104BEC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104C64 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104CDC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104DC4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104E3C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104EB4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104F38 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104FB4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010502C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010509C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105114 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010518C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010522C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001052A4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105368 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001053E0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105458 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001054C8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105540 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001055C0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105638 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001056B0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105728 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001057A8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105828 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001058A8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105920 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105998 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105A18 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105B40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105BC0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105C40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105CB8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105D38 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105E08 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105E94 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105F14 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105F94 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106004 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106084 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106104 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106174 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001061F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106274 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001062F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106380 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106400 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106488 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106510 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106598 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106608 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106690 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106718 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106790 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106818 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106890 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106918 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001069A0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106A18 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106A90 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106B18 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106B90 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106C08 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106C88 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106D10 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106D90 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106E24 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106EAC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106F24 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00106FA4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010701C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00107094 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00107124 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001071BC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010724C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001072DC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00107354 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001073CC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00107468 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00107504 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010759C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00107648 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001076E0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00107784 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00107828 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001078A0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00107944 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001079DC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00107A54 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00107B04 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00107BA4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0011F0B8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0011F108 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00159EC0 | `Resources/Sounds/volumebeep.wav` | Filesystem Path |  |
| 0x0017B40C | `iPod_Control/Device/Radio/RadioBuffer` | Filesystem Path |  |
| 0x00181CCC | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x001893A0 | `iPod_Control\Device` | Filesystem Path |  |
| 0x0018A15C | `iPod_Control/Device/Radio` | Filesystem Path |  |
| 0x0018A178 | `iPod_Control/Device/Radio/TunerSettings` | Filesystem Path |  |
| 0x001A131C | `iPod_Control/Device/Radio/TunerSettings` | Filesystem Path |  |
| 0x001C2544 | `Resources/Games` | Filesystem Path |  |
| 0x001C2554 | `iPod_Control` | Filesystem Path |  |
| 0x001C2574 | `iPod_Control/%s%s%s` | Filesystem Path |  |
| 0x001D6004 | `Resources/UI/` | Filesystem Path |  |
| 0x001D6024 | `Resources/UI/SilverDB.%s.LE.bin` | Filesystem Path |  |
| 0x001D7254 | `iPod_Control` | Filesystem Path |  |
| 0x001DB118 | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path |  |
| 0x001DE42C | `iPod_Control/Device/Users` | Filesystem Path |  |
| 0x001DE594 | `iPod_Control/Device/PlayCounts` | Filesystem Path |  |
| 0x001E64D4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0021AA40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x002337A0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x002337F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00249280 | `iPod_Control\Music\` | Filesystem Path |  |
| 0x0025B3A0 | `Resources/TrainerTemplates` | Filesystem Path |  |
| 0x00260530 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00277838 | `/iPod_Control/Device/iPod_Contacts.db` | Filesystem Path |  |
| 0x0027FB98 | `Resources/Sounds/shake.wav` | Filesystem Path |  |
| 0x0028F874 | `iPod_Control/Device/alarms` | Filesystem Path |  |
| 0x002936BC | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x00297110 | `Resources/Sounds/marimba.wav` | Filesystem Path |  |
| 0x002BBC9C | `iPod_Control/Device/Radio/Tuner_Metadata.log` | Filesystem Path |  |
| 0x002BBD40 | `iPod_Control/Device/Radio/Tuner_Readings.log` | Filesystem Path |  |
| 0x002EA3F8 | `Resources/Fonts` | Filesystem Path |  |
| 0x0032F7B0 | `iPod_Control/Device/Accessories/Tags/` | Filesystem Path |  |
| 0x0034F08C | `iPod_Control:Device:voices:` | Filesystem Path |  |
| 0x0045F61E | `iPod_Control/Device` | Filesystem Path |  |
| 0x0046CF6C | `Resources/UI/SilverImagesDB.LE.bin` | Filesystem Path |  |
| 0x0047C1D6 | `iPod_Control/Device` | Filesystem Path |  |
| 0x0056A881 | `Resources/Games/` | Filesystem Path |  |
| 0x0056BE72 | `iPod_Control\Device\log` | Filesystem Path |  |
| 0x0056BE92 | `/iPod_Control/Speakable` | Filesystem Path |  |
| 0x0056BEAA | `/iPod_Control/Speakable/UISS.plist` | Filesystem Path |  |
| 0x0056BECD | `/iPod_Control/Speakable/CacheInfo.plist` | Filesystem Path |  |
| 0x0056BEF5 | `/iPod_Control/Speakable/ConfigInfo.plist` | Filesystem Path |  |
| 0x0056BF1E | `/iPod_Control/Speakable/UISS_combined.plist.gz` | Filesystem Path |  |
| 0x0056BF4D | `/Resources/Speakable/UISS_combined.plist.gz` | Filesystem Path |  |
| 0x0056BF79 | `/Resources/Speakable` | Filesystem Path |  |
| 0x0056BFC6 | `iPod_Control/Tones` | Filesystem Path |  |
| 0x0056BFE6 | `/iPod_Control/Device/` | Filesystem Path |  |
| 0x0056BFFC | `iPod_Control/Device` | Filesystem Path |  |
| 0x0056C010 | `iPod_Control/Device/accessories.reg` | Filesystem Path |  |

---

## 19. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0045F3C4 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/eAppHostLib/eAppHostL` | Build Path |  |
| 0x0045FFAC | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/eAppHostLib/eAppMotor` | Build Path |  |
| 0x0047FDC4 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x00493154 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x00493220 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x004935F0 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/LIBXML/xpath.c` | Build Path |  |
| 0x004AE2F4 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Silver/3rdParty/freetype/src/autofit/` | Build Path |  |
| 0x004AEE70 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Silver/3rdParty/freetype/src/bdf/bdfd` | Build Path |  |
| 0x004B0038 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Silver/3rdParty/freetype/src/base/ftu` | Build Path |  |
| 0x004B0090 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Silver/3rdParty/freetype/src/base/fts` | Build Path |  |
| 0x004B00E8 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Silver/3rdParty/freetype/src/base/fto` | Build Path |  |
| 0x004B0208 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Silver/3rdParty/freetype/src/pfr/pfrg` | Build Path |  |
| 0x004B0260 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Silver/3rdParty/freetype/src/pfr/pfrc` | Build Path |  |
| 0x004B02B8 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Silver/3rdParty/freetype/src/pfr/pfro` | Build Path |  |
| 0x004B05FC | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Silver/3rdParty/freetype/src/psaux/t1` | Build Path |  |
| 0x004BF9A4 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Silver/3rdParty/freetype/src/sfnt/ttc` | Build Path |  |
| 0x004BFC20 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Silver/3rdParty/freetype/src/truetype` | Build Path |  |
| 0x004C018C | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Silver/3rdParty/freetype/src/type1/t1` | Build Path |  |
| 0x004C183C | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x004C1B74 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x004C2788 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x004C2E80 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x004D6684 | `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/Image3/Image3.c` | Build Path |  |

---

## 20. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0056D44A | `Electronic` | EQ Preset |  |
| 0x0056D455 | `Acoustic` | EQ Preset |  |
| 0x0056D472 | `Dance` | EQ Preset |  |
| 0x0056D478 | `Lounge` | EQ Preset |  |
| 0x0056D47F | `Rock` | EQ Preset |  |
| 0x0056D484 | `Classical` | EQ Preset |  |
| 0x0056D48E | `Latin` | EQ Preset |  |
| 0x0056D494 | `Piano` | EQ Preset |  |
| 0x0056D4F5 | `Loudness` | EQ Preset |  |
| 0x0056D503 | `Jazz` | EQ Preset |  |
| 0x0062CCB0 | `Acoustic` | EQ Preset |  |
| 0x0062CCBC | `Bass Booster` | EQ Preset |  |
| 0x0062CCDC | `Classical` | EQ Preset |  |
| 0x0062CCE8 | `Dance` | EQ Preset |  |
| 0x0062CCF8 | `Electronic` | EQ Preset |  |
| 0x0062CD0C | `Hip Hop` | EQ Preset |  |
| 0x0062CD14 | `Jazz` | EQ Preset |  |
| 0x0062CD1C | `Latin` | EQ Preset |  |
| 0x0062CD24 | `Loudness` | EQ Preset |  |
| 0x0062CD30 | `Lounge` | EQ Preset |  |
| 0x0062CD38 | `Piano` | EQ Preset |  |
| 0x0062CD48 | `Rock` | EQ Preset |  |
| 0x0062CD50 | `Small Speakers` | EQ Preset |  |
| 0x0062CD60 | `Spoken Word` | EQ Preset |  |
| 0x0062CD6C | `Treble Booster` | EQ Preset |  |
| 0x0062CD8C | `Vocal Booster` | EQ Preset |  |

---
