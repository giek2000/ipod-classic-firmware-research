# iPod Classic 6G Initial - RetailOS 1.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.1 |
| **IPSW** | iPod_24.1.1.ipsw |
| **Device** | iPod Classic 6G Initial (2008, 80/160GB, Click Wheel, Cover Flow, CE-ATA HDD) |
| **UpdaterFamilyID** | 24 |
| **Binary Size** | 9,865,856 bytes (9.41 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 9,863,808 bytes |
| **Total Strings (>=4)** | 66,508 |
| **Function Prologues** | 21,057 (ARM: 15,950, Thumb: 5,107) |
| **DRAM References** | 85,324 |
| **Peripheral Refs** | 5,616 |
| **Build** | N25FirmwareWin-359 |
| **SoC** | S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Codename** | N25C |
| **DFU PID** | 0x1223 |
| **SHA-256** | `1171a5149d37bd599343f26adb0d29a5c2b3f6d6d9b688de68ca0435e9f1bf88` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00092D84 | `TSilverCntlr` | Known | Controller |
| 0x00092D9C | `TCExtrasMenu` | Known | Controller |
| 0x00092DB4 | `TCGameScreen` | Known | Controller |
| 0x00092DCC | `TCGamesMenu` | Known | Controller |
| 0x00092DE0 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x00092E08 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00092E30 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x00092E5C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x00092E80 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x00092EA8 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x00092ED0 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x00092EF8 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x00092F20 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x00092F48 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x00092F78 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x00092FA4 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x00092FD4 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x00092FFC | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x00093024 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x00093050 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x00093078 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x000930A0 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x000930D0 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x00093100 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x00093208 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x00093238 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x00093260 | `TCRentalNotification` | Known | Controller |
| 0x00093280 | `TCRentalInfo` | Known | Controller |
| 0x00093298 | `TCRentalConfirmDelete` | Known | Controller |
| 0x000932B8 | `TCRentalDispatcher` | Known | Controller |
| 0x000932D4 | `TSilverGlobalCntlr` | Known | Controller |
| 0x000932F0 | `TSilverTrainerCntlr` | Known | Controller |
| 0x000E8768 | `TCSlideshowLCD` | Known | Controller |
| 0x000E8780 | `TCSlideshowTVOut` | Known | Controller |
| 0x000E879C | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x000E87BC | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x0010B698 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x0010B6C4 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x0010B6F0 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x0010B718 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x0010B744 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x0010B76C | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x0010B798 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x00112950 | `TCRemoteUI` | Known | Controller |
| 0x00112964 | `TCUnsupported` | Known | Controller |
| 0x00118D08 | `TCSpeakers` | Known | Controller |
| 0x00118D1C | `TCEQSetting` | Known | Controller |
| 0x00141290 | `TCSportTimer` | Known | Controller |
| 0x001412A8 | `TCSportTimerMenu` | Known | Controller |
| 0x001412C4 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x001412E8 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x00142668 | `TCVoiceMemos` | Known | Controller |
| 0x00142680 | `TCVoiceMemosMenu` | Known | Controller |
| 0x0014269C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x001426BC | `TCVoiceMemosPlayback` | Known | Controller |
| 0x001426DC | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x00153164 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x0015318C | `TCSettings_MainMenu` | Known | Controller |
| 0x001531A8 | `TCSettings_MusicMenu` | Known | Controller |
| 0x001531C8 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x001531E8 | `TCSettings_Brightness` | Known | Controller |
| 0x00153208 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0015322C | `TCSettings_EQ` | Known | Controller |
| 0x00153244 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x0015326C | `TCSettings_RadioRegions` | Known | Controller |
| 0x0015328C | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x001532B0 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x001532D4 | `TCDateTimeScreen` | Known | Controller |
| 0x001532F0 | `TCTimeZoneScreen` | Known | Controller |
| 0x0015330C | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x00153334 | `TCFirstBoot` | Known | Controller |
| 0x00168810 | `TCDemoMode` | Known | Controller |
| 0x0018E2A0 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x0018E2C0 | `TCAddressViewerDetails` | Known | Controller |
| 0x0018E2E0 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x0018E304 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x001B9B98 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x001B9BBC | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x001C135C | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x0024925C | `TC_LockDialog` | Known | Controller |
| 0x00249274 | `TC_LockScreen` | Known | Controller |
| 0x0024928C | `TC_LockediPod` | Known | Controller |
| 0x002492A4 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x002492C8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0024EDA4 | `TCClock` | Known | Controller |
| 0x0024EDB4 | `TCClockCityMenu` | Known | Controller |
| 0x0024EDCC | `TCClockRegionMenu` | Known | Controller |
| 0x0024EDE8 | `TCAlarmMenu` | Known | Controller |
| 0x0024EDFC | `TCSleepTimerMenu` | Known | Controller |
| 0x0024EE18 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x0024EE38 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x0024EE60 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x0024EE84 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0024EEA8 | `TCAlarmDatePicker` | Known | Controller |
| 0x0024EEC4 | `TCAlarmTriggered` | Known | Controller |
| 0x00255D10 | `TCNotesDispatcher` | Known | Controller |
| 0x00255D2C | `TCNotesLoading` | Known | Controller |
| 0x00255D44 | `TCNotesList` | Known | Controller |
| 0x00255D58 | `TCNotesContents` | Known | Controller |
| 0x00377028 | `TCAlarmTriggered` | Known | Controller |
| 0x0037703C | `TSilverCntlr` | Known | Controller |
| 0x0037705C | `TCClock` | Known | Controller |
| 0x00377064 | `TCClockRegionMenu` | Known | Controller |
| 0x00377078 | `TCClockCityMenu` | Known | Controller |
| 0x00377088 | `TCAlarmMenu` | Known | Controller |
| 0x00377094 | `TCSleepTimerMenu` | Known | Controller |
| 0x003770A8 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003770C0 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003770E0 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003770FC | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00377118 | `TCAlarmDatePicker` | Known | Controller |
| 0x00377150 | `TSilverCntlr` | Known | Controller |
| 0x00377170 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x00377300 | `TSilverCntlr` | Known | Controller |
| 0x00377320 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x00377340 | `TCSettings_Brightness` | Known | Controller |
| 0x00377358 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x00377374 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x00377394 | `TCSettings_RadioRegions` | Known | Controller |
| 0x003773AC | `TCSettings_EQ` | Known | Controller |
| 0x003773BC | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x003773D8 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x003773F8 | `TCFirstBoot` | Known | Controller |
| 0x00377404 | `TCSettings_MainMenu` | Known | Controller |
| 0x00377418 | `TCSettings_MusicMenu` | Known | Controller |
| 0x00377430 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x00377448 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x00377464 | `TCDateTimeScreen` | Known | Controller |
| 0x00377478 | `TCTimeZoneScreen` | Known | Controller |
| 0x0037E460 | `TSilverCntlr` | Known | Controller |
| 0x0037E480 | `TCClock` | Known | Controller |
| 0x0037E488 | `TCClockRegionMenu` | Known | Controller |
| 0x0037E49C | `TCClockCityMenu` | Known | Controller |
| 0x0037E4AC | `TCAlarmMenu` | Known | Controller |
| 0x0037E4B8 | `TCSleepTimerMenu` | Known | Controller |
| 0x0037E4CC | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x0037E544 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x0037E564 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x0037E580 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0037E5B4 | `TCAlarmDatePicker` | Known | Controller |
| 0x0037E5C8 | `TCAlarmTriggered` | Known | Controller |
| 0x00380044 | `TSilverCntlr` | Known | Controller |
| 0x00380064 | `TC_LockDialog` | Known | Controller |
| 0x00380074 | `TC_LockScreen` | Known | Controller |
| 0x00380084 | `TC_LockediPod` | Known | Controller |
| 0x00380094 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003800B0 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003800C8 | `TSilverCntlr` | Known | Controller |
| 0x00380230 | `TSilverCntlr` | Known | Controller |
| 0x0038024C | `TSilverCntlr` | Known | Controller |
| 0x003802B0 | `TSilverCntlr` | Known | Controller |
| 0x003802D0 | `TCNotesDispatcher` | Known | Controller |
| 0x003802E4 | `TCNotesLoading` | Known | Controller |
| 0x003802F4 | `TCNotesBase` | Known | Controller |
| 0x00380300 | `TCNotesList` | Known | Controller |
| 0x0038030C | `TCNotesContents` | Known | Controller |
| 0x0038031C | `TSilverCntlr` | Known | Controller |
| 0x0038033C | `TCRemoteUI` | Known | Controller |
| 0x00380348 | `TCUnsupported` | Known | Controller |
| 0x00380358 | `TSilverCntlr` | Known | Controller |
| 0x003803BC | `TSilverCntlr` | Known | Controller |
| 0x003803DC | `TCSportTimer` | Known | Controller |
| 0x003803EC | `TCSportTimerMenu` | Known | Controller |
| 0x00380400 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x0038041C | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x0038044C | `TSilverCntlr` | Known | Controller |
| 0x00380574 | `TSilverCntlr` | Known | Controller |
| 0x00380594 | `TCDemoMode` | Known | Controller |
| 0x003805A0 | `TCClock` | Known | Controller |
| 0x003805A8 | `TCClockRegionMenu` | Known | Controller |
| 0x003805BC | `TCClockCityMenu` | Known | Controller |
| 0x003805CC | `TCAlarmMenu` | Known | Controller |
| 0x003805D8 | `TCSleepTimerMenu` | Known | Controller |
| 0x003805EC | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x00380604 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x00380624 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x00380640 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0038065C | `TCAlarmDatePicker` | Known | Controller |
| 0x00380670 | `TCAlarmTriggered` | Known | Controller |
| 0x00380690 | `TSilverCntlr` | Known | Controller |
| 0x003806AC | `TSilverCntlr` | Known | Controller |
| 0x003806BC | `TSilverCntlr` | Known | Controller |
| 0x003806DC | `TCVoiceMemos` | Known | Controller |
| 0x003806EC | `TCVoiceMemosMenu` | Known | Controller |
| 0x00380700 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x00380718 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x00380730 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x00380750 | `TSilverCntlr` | Known | Controller |
| 0x003807B0 | `TSilverCntlr` | Known | Controller |
| 0x0038081C | `TSilverCntlr` | Known | Controller |
| 0x0038180C | `TSilverCntlr` | Known | Controller |
| 0x00381918 | `TSilverCntlr` | Known | Controller |
| 0x0038A138 | `TSilverCntlr` | Known | Controller |
| 0x0038A158 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x0038A170 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x0038A18C | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x0038A1AC | `TCAddressViewerDetails` | Known | Controller |
| 0x0038A1C4 | `TSilverCntlr` | Known | Controller |
| 0x0038A1E4 | `TSilverCalendarCntlr_Base` | Known | Controller |
| 0x0038A200 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x0038A224 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x0038A248 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x0038A268 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x0038A28C | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x0038A2AC | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x0038A2D0 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x0038A4A8 | `TSilverCntlr` | Known | Controller |
| 0x0038A4C8 | `TC_LockDialog` | Known | Controller |
| 0x0038A4D8 | `TC_LockScreen` | Known | Controller |
| 0x0038A4E8 | `TC_LockediPod` | Known | Controller |
| 0x0038A4F8 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x0038A51C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0038A5D0 | `TSilverCntlr` | Known | Controller |
| 0x0038A6F0 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x0038A70C | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x0038A72C | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0038A74C | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0038A774 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x0038A798 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x0038A7C0 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x0038A7E0 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x0038A800 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x0038A820 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x0038A840 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0038A868 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x0038A890 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x0038A8B0 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x0038A8D0 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x0038A8F4 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x0038A914 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x0038A938 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x0038A960 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x0038A98C | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x0038A9AC | `TCRentalNotification` | Known | Controller |
| 0x0038A9C4 | `TCRentalInfo` | Known | Controller |
| 0x0038A9D4 | `TCRentalConfirmDelete` | Known | Controller |
| 0x0038A9EC | `TCRentalDispatcher` | Known | Controller |
| 0x0038B2DC | `TSilverCntlr` | Known | Controller |
| 0x0038B3A0 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x0038B3BC | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x0038B3DC | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0038B3FC | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0038B424 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x0038B448 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x0038B470 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x0038B490 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x0038B4B0 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x0038B4D0 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x0038B4F0 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0038B518 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x0038B568 | `TCSlideshowTVOut` | Known | Controller |
| 0x0038B57C | `TCSlideshowLCD` | Known | Controller |
| 0x0038B58C | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x0038B5A4 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x0038B5C4 | `TSilverCntlr` | Known | Controller |
| 0x0038B5F0 | `TSilverCntlr` | Known | Controller |
| 0x0038B610 | `TCUnsupported` | Known | Controller |
| 0x0038B630 | `TSilverCntlr` | Known | Controller |
| 0x0038B670 | `TSilverCntlr` | Known | Controller |
| 0x0038B690 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x0038B6AC | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x0038B6C4 | `TSilverCntlr` | Known | Controller |
| 0x0038B6E4 | `TCSpeakers` | Known | Controller |
| 0x0038B6F0 | `TCEQSetting` | Known | Controller |
| 0x0038B710 | `TSilverCntlr` | Known | Controller |
| 0x0038B778 | `TSilverCntlr` | Known | Controller |
| 0x0038B798 | `TCExtrasMenu` | Known | Controller |
| 0x0038B7A8 | `TCGamesMenu` | Known | Controller |
| 0x0038B7B4 | `TCGameScreen` | Known | Controller |
| 0x0038B7C4 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x0038B7E4 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x0038B804 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x0038B824 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x0038B848 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x0038B864 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x0038B884 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0038B8A4 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0038B8CC | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x0038B8F0 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x0038B918 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x0038B938 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x0038B958 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x0038B978 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x0038B998 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0038B9C0 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x0038B9E8 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x0038BA08 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x0038BA28 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x0038BA4C | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x0038BA6C | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x0038BA90 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x0038BAB8 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x0038BAE4 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x0038BB04 | `TCRentalNotification` | Known | Controller |
| 0x0038BB1C | `TCRentalInfo` | Known | Controller |
| 0x0038BB2C | `TCRentalConfirmDelete` | Known | Controller |
| 0x0038BB44 | `TCRentalDispatcher` | Known | Controller |
| 0x0038BB58 | `TSilverGlobalCntlr` | Known | Controller |
| 0x0038BB6C | `TSilverTrainerCntlr` | Known | Controller |
| 0x004101B8 | `TCCCCCCCCCCCCCCCCCCCCCCT` | Known | Controller |
| 0x006A2BFA | `TCNotesDispatcher"` | Known | Controller |
| 0x006A2CB9 | `TCLockChosenDispatcher"` | Known | Controller |
| 0x006A2D7C | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x006ACDE1 | `TCNotesDispatcher"` | Known | Controller |
| 0x006ACF43 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x006C1C7C | `TCExtrasMenuTCAddressViewerMainMenu` | Known | Controller |
| 0x006C1CA0 | `TCAddressViewerDetails` | Known | Controller |
| 0x006C1CB8 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x006C1CD4 | `TCAlarmMenu` | Known | Controller |
| 0x006C1CE0 | `TCSleepTimerMenuTCAlarmPropertiesMenu` | Known | Controller |
| 0x006C1D08 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x006C1D28 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x006C1D44 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x006C1D60 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x006C1D7C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x006C1D98 | `TCAlarmDatePicker` | Known | Controller |
| 0x006C1DAC | `TCAlarmDatePicker` | Known | Controller |
| 0x006C1DC0 | `TCAlarmTriggeredTSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x006C1DEC | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x006C1E10 | `TSilverCalendarCntlr_MonthViewerTSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x006C1E50 | `TSilverCalendarCntlr_EventViewerTSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x006C1E90 | `TSilverCalendarCntlr_EventViewerTSilverCntlrTCClockRegionMenu` | Known | Controller |
| 0x006C1ED0 | `TCClockCityMenu` | Known | Controller |
| 0x006C1EE0 | `TCClockCityMenu` | Known | Controller |
| 0x006C1EF0 | `TCClockCityMenu` | Known | Controller |
| 0x006C1F00 | `TCClockCityMenu` | Known | Controller |
| 0x006C1F10 | `TCClockCityMenu` | Known | Controller |
| 0x006C1F20 | `TCClockCityMenu` | Known | Controller |
| 0x006C1F30 | `TCClockCityMenu` | Known | Controller |
| 0x006C1F40 | `TCClockCityMenu` | Known | Controller |
| 0x006C1F50 | `TCClock` | Known | Controller |
| 0x006C1F68 | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDe` | Known | Controller |
| 0x006C1FC0 | `TCGamesMenu` | Known | Controller |
| 0x006C1FCC | `TCGameScreenTC_LockediPod` | Known | Controller |
| 0x006C1FE8 | `TC_LockDialog` | Known | Controller |
| 0x006C1FF8 | `TC_LockScreen` | Known | Controller |
| 0x006C2008 | `TC_VolumeLimitLockScreenTSilverCntlrTSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x006C204C | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x006C206C | `TSilverMediaListCntlr_AudiobooksTSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x006C20B4 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x006C20D0 | `TSilverMediaListCntlr_AlbumsTSilverMediaListCntlr_Artists` | Known | Controller |
| 0x006C210C | `TSilverMediaListCntlr_GenresTSilverMediaListCntlr_Composers` | Known | Controller |
| 0x006C2148 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x006C2168 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x006C2190 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x006C21B0 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x006C21D0 | `TSilverMediaListCntlr_TVEpisodesTSilverMediaListCntlr_MoviesTSilverMediaListCntl` | Known | Controller |
| 0x006C222C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x006C2254 | `TSilverMediaListCntlr_VideoPlaylistsTSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x006C2298 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x006C22C4 | `TSilverCntlrTSilverMainMediaListCntlr_VideosTSilverSettingsVideoCntlr` | Known | Controller |
| 0x006C230C | `TCFirstBoot` | Known | Controller |
| 0x006C23B4 | `TCRentalInfoTCRentalConfirmDelete` | Known | Controller |
| 0x006C23D8 | `TSilverCntlrTCRentalNotificationTCRentalNotificationTCRentalNotificationTCNotesL` | Known | Controller |
| 0x006C2430 | `TCNotesList` | Known | Controller |
| 0x006C243C | `TCNotesList` | Known | Controller |
| 0x006C2448 | `TCNotesContents` | Known | Controller |
| 0x006C2458 | `TCNotesContents` | Known | Controller |
| 0x006C2468 | `TCNotesContents` | Known | Controller |
| 0x006C2478 | `TCNotesContents` | Known | Controller |
| 0x006C2534 | `TCSlideshowLCD` | Known | Controller |
| 0x006C2544 | `TCSlideshowTVOutTCSlideshow_TVOutAskTCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x006C2594 | `TCRemoteUI` | Known | Controller |
| 0x006C25A0 | `TCUnsupported` | Known | Controller |
| 0x006C25B0 | `TSilverSettingsMenuListCntlrTSilverSettingsMenuListCntlrTSilverSettingsMenuListC` | Known | Controller |
| 0x006C2618 | `TCSettings_MusicMenuTCSettings_VolumeLimit` | Known | Controller |
| 0x006C2644 | `TCSettings_Brightness` | Known | Controller |
| 0x006C265C | `TCSettings_BacklightTimer` | Known | Controller |
| 0x006C2678 | `TCSettings_AudiobookSettingsTCSettings_RadioRegions` | Known | Controller |
| 0x006C26AC | `TCSettings_EQ` | Known | Controller |
| 0x006C26BC | `TSilverSettingsMenuListCntlrTSilverCntlrTCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x006C2704 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x006C2720 | `TCSettings_MainMenu` | Known | Controller |
| 0x006C2734 | `TCSettings_MusicMenuTCSportTimerTCSportTimerMenuTCSportTimerSessionScreen` | Known | Controller |
| 0x006C2780 | `TSilverCntlrTUnitTestSuiteCntlr` | Known | Controller |
| 0x006C2800 | `TCVoiceMemosTCVoiceMemosMainMenuTCVoiceMemosMainMenuTCVoiceMemosMainMenuTSearchC` | Known | Controller |
| 0x006C2860 | `TCEQSetting` | Known | Controller |
| 0x006C290E | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x006C3C11 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x006C981A | `TCLockChosenDispatcher` | Known | Controller |
| 0x006C9878 | `TCNotesDispatcher` | Known | Controller |
| 0x006CB456 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006CB4B4 | `TCNotesDispatcher` | Known | Controller |
| 0x006CD092 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006CD0F0 | `TCNotesDispatcher` | Known | Controller |
| 0x006CECCE | `TCLockChosenDispatcher` | Known | Controller |
| 0x006CED2C | `TCNotesDispatcher` | Known | Controller |
| 0x006D090A | `TCLockChosenDispatcher` | Known | Controller |
| 0x006D0968 | `TCNotesDispatcher` | Known | Controller |
| 0x006D2546 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006D25A4 | `TCNotesDispatcher` | Known | Controller |
| 0x006D4182 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006D41E0 | `TCNotesDispatcher` | Known | Controller |
| 0x006D5DBE | `TCLockChosenDispatcher` | Known | Controller |
| 0x006D5E1C | `TCNotesDispatcher` | Known | Controller |
| 0x006D79FA | `TCLockChosenDispatcher` | Known | Controller |
| 0x006D7A58 | `TCNotesDispatcher` | Known | Controller |
| 0x006D9636 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006D9694 | `TCNotesDispatcher` | Known | Controller |
| 0x006DB272 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006DB2D0 | `TCNotesDispatcher` | Known | Controller |
| 0x006DCEAE | `TCLockChosenDispatcher` | Known | Controller |
| 0x006DCF0C | `TCNotesDispatcher` | Known | Controller |
| 0x006DEAEA | `TCLockChosenDispatcher` | Known | Controller |
| 0x006DEB48 | `TCNotesDispatcher` | Known | Controller |
| 0x006E0726 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006E0784 | `TCNotesDispatcher` | Known | Controller |
| 0x006E2362 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006E23C0 | `TCNotesDispatcher` | Known | Controller |
| 0x006E3F9E | `TCLockChosenDispatcher` | Known | Controller |
| 0x006E3FFC | `TCNotesDispatcher` | Known | Controller |
| 0x006E5BDA | `TCLockChosenDispatcher` | Known | Controller |
| 0x006E5C38 | `TCNotesDispatcher` | Known | Controller |
| 0x006E7816 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006E7874 | `TCNotesDispatcher` | Known | Controller |
| 0x006E9452 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006E94B0 | `TCNotesDispatcher` | Known | Controller |
| 0x006EB08E | `TCLockChosenDispatcher` | Known | Controller |
| 0x006EB0EC | `TCNotesDispatcher` | Known | Controller |
| 0x006ECCCA | `TCLockChosenDispatcher` | Known | Controller |
| 0x006ECD28 | `TCNotesDispatcher` | Known | Controller |
| 0x006EE906 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006EE964 | `TCNotesDispatcher` | Known | Controller |
| 0x006F0542 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006F05A0 | `TCNotesDispatcher` | Known | Controller |
| 0x006F217E | `TCLockChosenDispatcher` | Known | Controller |
| 0x006F21DC | `TCNotesDispatcher` | Known | Controller |
| 0x006F3DBA | `TCLockChosenDispatcher` | Known | Controller |
| 0x006F3E18 | `TCNotesDispatcher` | Known | Controller |
| 0x006F59F6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006F5A54 | `TCNotesDispatcher` | Known | Controller |
| 0x006F7632 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006F7690 | `TCNotesDispatcher` | Known | Controller |
| 0x006F926E | `TCLockChosenDispatcher` | Known | Controller |
| 0x006F92CC | `TCNotesDispatcher` | Known | Controller |
| 0x006FAEAA | `TCLockChosenDispatcher` | Known | Controller |
| 0x006FAF08 | `TCNotesDispatcher` | Known | Controller |
| 0x006FCAE6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006FCB44 | `TCNotesDispatcher` | Known | Controller |
| 0x006FE722 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006FE780 | `TCNotesDispatcher` | Known | Controller |
| 0x0070035E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007003BC | `TCNotesDispatcher` | Known | Controller |
| 0x00701F9A | `TCLockChosenDispatcher` | Known | Controller |
| 0x00701FF8 | `TCNotesDispatcher` | Known | Controller |
| 0x00703BD6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00703C34 | `TCNotesDispatcher` | Known | Controller |
| 0x00705812 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00705870 | `TCNotesDispatcher` | Known | Controller |
| 0x0070744E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007074AC | `TCNotesDispatcher` | Known | Controller |
| 0x0070908A | `TCLockChosenDispatcher` | Known | Controller |
| 0x007090E8 | `TCNotesDispatcher` | Known | Controller |
| 0x00714CC0 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x00714F82 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x007157B8 | `TCRentalDispatcher` | Known | Controller |
| 0x00716070 | `TCRentalDispatcher` | Known | Controller |
| 0x00716928 | `TCRentalDispatcher` | Known | Controller |
| 0x007171E0 | `TCRentalDispatcher` | Known | Controller |
| 0x00717A98 | `TCRentalDispatcher` | Known | Controller |
| 0x00718350 | `TCRentalDispatcher` | Known | Controller |
| 0x00718C08 | `TCRentalDispatcher` | Known | Controller |
| 0x007194C0 | `TCRentalDispatcher` | Known | Controller |
| 0x0084D850 | `TCMockupModeNavScreen` | Known | Controller |
| 0x0084D868 | `TSilverCntlr` | Known | Controller |
| 0x0084D888 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x0084D8D8 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x0084D8F8 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x0084D918 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x0084D93C | `TCExtrasMenu` | Known | Controller |
| 0x0084DA4C | `TSilverCntlr` | Known | Controller |
| 0x0084DA6C | `TCSlideshowTVOut` | Known | Controller |
| 0x0084DA80 | `TCSlideshowLCD` | Known | Controller |
| 0x0084DA90 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x0084DAA8 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x0084DAE4 | `TSilverCntlr` | Known | Controller |
| 0x0084DB60 | `TCSlideshowTVOut` | Known | Controller |
| 0x0084DB74 | `TCSlideshowLCD` | Known | Controller |
| 0x0084DB84 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x0084DB9C | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x0084DBBC | `TSilverCntlr` | Known | Controller |
| 0x0084DC04 | `TSilverCntlr` | Known | Controller |
| 0x0084DC24 | `TCGamesMenu` | Known | Controller |
| 0x0084DC30 | `TCGameScreen` | Known | Controller |
| 0x00901B8B | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |

---

## 2. Settings (Toggle/Show/TC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00121E54 | `ShowSetting_EQ` | Known | User setting |
| 0x001C2B10 | `ToggleSetting_Repeat` | Known | User setting |
| 0x001C2B2C | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001C2B44 | `ToggleSetting_TVOut` | Known | User setting |
| 0x001C2B58 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x001EA054 | `ShowSetting_Backlight` | Known | User setting |
| 0x001FB804 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001FB820 | `ToggleSetting_Repeat` | Known | User setting |
| 0x001FB838 | `ToggleSetting_SortBy` | Known | User setting |
| 0x001FB850 | `ToggleSetting_ClassicUI` | Known | User setting |
| 0x001FB868 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x001FB884 | `ToggleSetting_Clicker` | Known | User setting |
| 0x001FB89C | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x001FB8BC | `ToggleSetting_24HourClock` | Known | User setting |
| 0x001FB8D8 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x001FB8F4 | `ShowSetting_Shuffle` | Known | User setting |
| 0x001FBAA0 | `ShowSetting_Repeat` | Known | User setting |
| 0x001FBAB4 | `ShowSetting_About` | Known | User setting |
| 0x001FBAC8 | `ShowSetting_MainMenu` | Known | User setting |
| 0x001FBAE0 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x001FBAF8 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x001FBB10 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x001FBB2C | `ShowSetting_Brightness` | Known | User setting |
| 0x001FBB44 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x001FBB5C | `ShowSetting_RadioRegions` | Known | User setting |
| 0x001FBB78 | `ShowSetting_EQ` | Known | User setting |
| 0x001FBB88 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x001FBD24 | `ShowSetting_Clicker` | Known | User setting |
| 0x001FBD38 | `ShowSetting_DateAndTime` | Known | User setting |
| 0x001FBD50 | `ShowSetting_SortBy` | Known | User setting |
| 0x001FBD64 | `ShowSetting_ClassicUI` | Known | User setting |
| 0x001FBD7C | `ShowSetting_Language` | Known | User setting |
| 0x001FBD94 | `ShowSetting_Legal` | Known | User setting |
| 0x001FBDA8 | `ShowSetting_ResetAll` | Known | User setting |
| 0x006ABC69 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x006ABD19 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x006AE3AE | `ShowSetting_About` | Known | User setting |
| 0x006AE450 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x006AE494 | `ShowSetting_Shuffle` | Known | User setting |
| 0x006AE50B | `ToggleSetting_Repeat` | Known | User setting |
| 0x006AE54E | `ShowSetting_Repeat` | Known | User setting |
| 0x006AE658 | `ShowSetting_MainMenu` | Known | User setting |
| 0x006AE768 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x006AE830 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x006AE8FA | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x006AEA12 | `ShowSetting_Brightness` | Known | User setting |
| 0x006AEB48 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x006AEC59 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x006AED5A | `ShowSetting_EQ` | Known | User setting |
| 0x006AEDC7 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x006AEE0E | `ShowSetting_SoundCheck` | Known | User setting |
| 0x006AEE8B | `ToggleSetting_Clicker` | Known | User setting |
| 0x006AEECF | `ShowSetting_Clicker` | Known | User setting |
| 0x006AF036 | `ToggleSetting_SortBy` | Known | User setting |
| 0x006AF079 | `ShowSetting_SortBy` | Known | User setting |
| 0x006AF17A | `ShowSetting_Language` | Known | User setting |
| 0x006AF28A | `ShowSetting_Legal` | Known | User setting |
| 0x006AF3BB | `ShowSetting_ResetAll` | Known | User setting |
| 0x006AF52C | `ShowSetting_Backlight` | Known | User setting |
| 0x006AF5DC | `ShowSetting_Backlight` | Known | User setting |
| 0x006AF68C | `ShowSetting_Backlight` | Known | User setting |
| 0x006AF73D | `ShowSetting_Backlight` | Known | User setting |
| 0x006AF7EE | `ShowSetting_Backlight` | Known | User setting |
| 0x006AF89F | `ShowSetting_Backlight` | Known | User setting |
| 0x006AF953 | `ShowSetting_Backlight` | Known | User setting |
| 0x006AFA02 | `ShowSetting_EQ` | Known | User setting |
| 0x006AFA77 | `ShowSetting_Language` | Known | User setting |
| 0x0072A5D4 | `ToggleSetting_Repeat` | Known | User setting |
| 0x0072A60E | `ToggleSetting_Shuffle` | Known | User setting |
| 0x0072A6D0 | `ToggleSetting_TVOut` | Known | User setting |
| 0x0072A709 | `ToggleSetting_TVSignal` | Known | User setting |

---

## 3. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0013D108 | `MockupMode/MockupMode.xml` | Hidden | Developer Tool |
| 0x0013D608 | `MockupMode/` | Hidden | Developer Tool |
| 0x00232D10 | `Channel UnitTests` | Hidden | Developer Tool |
| 0x00282315 | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x00282358 | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x0028236D | `RTXCbug> ` | Hidden | Developer Tool |
| 0x00282D49 | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x00293070 | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x0032B435 | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x0032B4FD | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x0037C421 | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x006C27A0 | `TUnitTestSuiteTestsCntlrTSilverCntlrTCVoiceMemosMainMenuTCVoiceMemosMenuTCVoiceM` | Hidden | Developer Tool |
| 0x00750524 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0078BFDC | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0079E4A8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x007B5A48 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x007C7B58 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x007D16B0 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x007DAF24 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x007EFD98 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x007F9858 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0081FA48 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0083DCD8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00846F30 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008F4309 | `Debug_ListItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x008F4321 | `Debug_MenuItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x008F4A0A | `DemoMode_Background_Image` | Hidden | Demo/Retail Mode |
| 0x008F551F | `DemoMode_Icon_Image` | Hidden | Demo/Retail Mode |
| 0x008F7099 | `Debug_Menu_Title` | Hidden | Debug/Diagnostic |
| 0x008F70BE | `UnitTest_Menu_Title` | Hidden | Developer Tool |
| 0x008FF850 | `UnitTestModel` | Hidden | Developer Tool |
| 0x0090021B | `UnitTest_ListItem` | Hidden | Developer Tool |
| 0x009012A3 | `DemoMode_View_Main` | Hidden | Demo/Retail Mode |
| 0x00901478 | `DemoMode_Icon` | Hidden | Demo/Retail Mode |
| 0x009021F3 | `UnitTestApp` | Hidden | Developer Tool |
| 0x0090278A | `Debug_ListItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009027A5 | `Debug_MenuItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x00902EA3 | `DemoMode_Text_Color` | Hidden | Demo/Retail Mode |
| 0x00903267 | `Debug_ListItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x0090327E | `Debug_MenuItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x00907143 | `Debug_ListItem_UnitTest` | Hidden | Developer Tool |
| 0x0090715B | `Debug_MenuItem_UnitTest` | Hidden | Developer Tool |
| 0x0090B34D | `Debug_ListItem_Memory` | Hidden | Debug/Diagnostic |
| 0x0090B363 | `Debug_MenuItem_Memory` | Hidden | Debug/Diagnostic |

---

## 4. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000678F | `"MeCCADecode` | Known | Audio system |
| 0x001337BC | `AudioCodecs` | Known | Audio system |
| 0x00175C30 | `MeCCA_PCM_Output.wav` | Known | Audio system |
| 0x0018D4DC | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x0019696C | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x00196B74 | `MeCCAVideoDecode` | Known | Audio system |
| 0x00859328 | `MeCCA_StreamCache` | Known | Audio system |

---

## 5. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E4564 | `HandleWheel` | Known | Event handler |
| 0x000E4570 | `HandlePlayPause` | Known | Event handler |
| 0x000E4580 | `HandleSelectDown` | Known | Event handler |
| 0x000E4594 | `HandleNext` | Known | Event handler |
| 0x000E45A0 | `HandlePrevious` | Known | Event handler |
| 0x000E45B0 | `HandleNextPushAndHold` | Known | Event handler |
| 0x000E45C8 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000E4860 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000E4880 | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x000F0234 | `HandleSelect` | Known | Event handler |
| 0x000F0248 | `HandleHilite` | Known | Event handler |
| 0x000F05E0 | `HandleEQSettingSelected` | Known | Event handler |
| 0x000F0A10 | `HandleSelect` | Known | Event handler |
| 0x000F0A24 | `HandleGameHilited` | Known | Event handler |
| 0x000F0CD4 | `HandleNotesSelected` | Known | Event handler |
| 0x000F0CEC | `HandleNotesPop` | Known | Event handler |
| 0x000F0CFC | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x000FEC80 | `HandleVolumeWheel` | Known | Event handler |
| 0x000FEC94 | `HandleVolumeChange` | Known | Event handler |
| 0x000FECA8 | `HandleTimerDone` | Known | Event handler |
| 0x000FECB8 | `HandleFrequencyChange` | Known | Event handler |
| 0x000FED30 | `HandleTuning` | Known | Event handler |
| 0x000FED40 | `HandleTuningSelect` | Known | Event handler |
| 0x00109600 | `HandleLock` | Known | Event handler |
| 0x00109610 | `HandleAddressBook` | Known | Event handler |
| 0x00109CF8 | `HandleSelect` | Known | Event handler |
| 0x0010A230 | `HandleExit` | Known | Event handler |
| 0x0010A240 | `HandleLap` | Known | Event handler |
| 0x0010A24C | `HandleResume` | Known | Event handler |
| 0x0010A25C | `HandleStartStop` | Known | Event handler |
| 0x0010A4E4 | `HandleWheel` | Known | Event handler |
| 0x0010A4F4 | `HandlePlayPause` | Known | Event handler |
| 0x0010A504 | `HandleSelectDown` | Known | Event handler |
| 0x0010A518 | `HandleHilite` | Known | Event handler |
| 0x00114100 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x00122088 | `HandleExitUnsupported` | Known | Event handler |
| 0x00138E74 | `HandleNotesPop` | Known | Event handler |
| 0x00138E88 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00139D6C | `HandleSelect` | Known | Event handler |
| 0x00139D80 | `HandleWheel` | Known | Event handler |
| 0x00139D8C | `HandleImageNext` | Known | Event handler |
| 0x00139D9C | `HandleImagePrev` | Known | Event handler |
| 0x00139DAC | `HandleImageLast` | Known | Event handler |
| 0x00139DBC | `HandleImageFirst` | Known | Event handler |
| 0x00139DD0 | `HandlePlayPause` | Known | Event handler |
| 0x00139DE0 | `HandlePlay` | Known | Event handler |
| 0x00139DEC | `HandlePause` | Known | Event handler |
| 0x0014DCA8 | `HandleSelectCity` | Known | Event handler |
| 0x0014DCC0 | `HandleHighlightCity` | Known | Event handler |
| 0x0014EBE8 | `HandleWantPopFlow` | Known | Event handler |
| 0x0014EC00 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x0014EC1C | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0014EC38 | `HandleFlowNext` | Known | Event handler |
| 0x0014EC48 | `HandleFlowPrev` | Known | Event handler |
| 0x0014EC58 | `HandleFlowWheel` | Known | Event handler |
| 0x0014EC68 | `HandleAlbumSelected` | Known | Event handler |
| 0x0014EC7C | `HandlePlayPause` | Known | Event handler |
| 0x0014EC8C | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00177ACC | `HandleLeaveAlarm` | Known | Event handler |
| 0x00177EBC | `HandleSelect` | Known | Event handler |
| 0x00178D7C | `HandleSelect` | Known | Event handler |
| 0x00178D90 | `HandleWheel` | Known | Event handler |
| 0x00178D9C | `HandleImageNext` | Known | Event handler |
| 0x00178DAC | `HandleImagePrev` | Known | Event handler |
| 0x00178DBC | `HandleImageLast` | Known | Event handler |
| 0x00178DCC | `HandleImageFirst` | Known | Event handler |
| 0x00178DE0 | `HandlePlayPause` | Known | Event handler |
| 0x00178DF0 | `HandlePlay` | Known | Event handler |
| 0x00178DFC | `HandlePause` | Known | Event handler |
| 0x0017929C | `HandleNew` | Known | Event handler |
| 0x001792AC | `HandleClear` | Known | Event handler |
| 0x001792B8 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x001792D4 | `HandleSelectIndexedSession` | Known | Event handler |
| 0x001795E4 | `HandleWheel` | Known | Event handler |
| 0x001795F4 | `HandleArrowUp` | Known | Event handler |
| 0x00179604 | `HandleArrowDown` | Known | Event handler |
| 0x0017B828 | `HandleHiliteAlbum` | Known | Event handler |
| 0x0017B840 | `HandleBrowseAlbum` | Known | Event handler |
| 0x0017B854 | `HandlePlayPause` | Known | Event handler |
| 0x00190AFC | `HandleSelect` | Known | Event handler |
| 0x00190C8C | `HandleSelectRegion` | Known | Event handler |
| 0x001A618C | `HandleImageWheel` | Known | Event handler |
| 0x001A61A4 | `HandlePlayPause` | Known | Event handler |
| 0x001A61B4 | `HandleBrowseLarge` | Known | Event handler |
| 0x001A61C8 | `HandleBrowseSmall` | Known | Event handler |
| 0x001A61DC | `HandleBrowseSlideshow` | Known | Event handler |
| 0x001A61F4 | `HandleImageNext` | Known | Event handler |
| 0x001A6204 | `HandleImagePrev` | Known | Event handler |
| 0x001A6214 | `HandleHilite` | Known | Event handler |
| 0x001A6224 | `HandleImageLast` | Known | Event handler |
| 0x001A6234 | `HandleImageFirst` | Known | Event handler |
| 0x001A6248 | `HandleScreenNext` | Known | Event handler |
| 0x001A625C | `HandleScreenPrev` | Known | Event handler |
| 0x001A8B40 | `HandlePlayPause` | Known | Event handler |
| 0x001A8B54 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x001A8B70 | `HandleNext` | Known | Event handler |
| 0x001A8B7C | `HandleNextPressAndHold` | Known | Event handler |
| 0x001A8B94 | `HandlePrevious` | Known | Event handler |
| 0x001A8BA4 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x001A8BC0 | `HandleRemotePlayPause` | Known | Event handler |
| 0x001A8BD8 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x001A8BFC | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x001A8C14 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x001A8C2C | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x001A8DFC | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x001A8E14 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x001A8E2C | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x001A8E48 | `HandleRemoteStop` | Known | Event handler |
| 0x001A8E5C | `HandleRemotePlay` | Known | Event handler |
| 0x001A8E70 | `HandleRemotePause` | Known | Event handler |
| 0x001A8E84 | `HandleRemoteMute` | Known | Event handler |
| 0x001A8E98 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x001A8EB0 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x001A8EC8 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x001A8EE4 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x001A9108 | `HandleRemoteShuffle` | Known | Event handler |
| 0x001A911C | `HandleRemoteRepeat` | Known | Event handler |
| 0x001A9130 | `HandleRemoteOn` | Known | Event handler |
| 0x001A9140 | `HandleRemoteOff` | Known | Event handler |
| 0x001A9150 | `HandleRemoteBacklight` | Known | Event handler |
| 0x001A9168 | `HandleRemoteFFDown` | Known | Event handler |
| 0x001A917C | `HandleRemoteFFUp` | Known | Event handler |
| 0x001A9190 | `HandleRemoteRewDown` | Known | Event handler |
| 0x001A91A4 | `HandleRemoteRewUp` | Known | Event handler |
| 0x001A91B8 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x001A91D0 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x001A91E4 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x001A91FC | `HandleRemoteSelectUp` | Known | Event handler |
| 0x001A93CC | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x001A93E4 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x001A93FC | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x001A9418 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x001A9430 | `HandleRemoteEvent` | Known | Event handler |
| 0x001A9444 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x001A9460 | `HandleAudioPlayPause` | Known | Event handler |
| 0x001A9478 | `HandleAudioNext` | Known | Event handler |
| 0x001A9488 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x001A94A4 | `HandleAudioPrevious` | Known | Event handler |
| 0x001A94B8 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x001A96B8 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x001A96D0 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x001A96E8 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x001A9700 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x001A9714 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x001A972C | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x001A9744 | `HandleAudioStop` | Known | Event handler |
| 0x001A9754 | `HandleAudioPlay` | Known | Event handler |
| 0x001A9764 | `HandleAudioPause` | Known | Event handler |
| 0x001A9778 | `HandleAudioMute` | Known | Event handler |
| 0x001A9788 | `HandleAudioNextChapter` | Known | Event handler |
| 0x001A97A0 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x001A99C0 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x001A99D8 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x001A99F0 | `HandleAudioShuffle` | Known | Event handler |
| 0x001A9A04 | `HandleAudioRepeat` | Known | Event handler |
| 0x001A9A18 | `HandleAudioFFDown` | Known | Event handler |
| 0x001A9A2C | `HandleAudioFFUp` | Known | Event handler |
| 0x001A9A3C | `HandleAudioRewDown` | Known | Event handler |
| 0x001A9A50 | `HandleAudioRewUp` | Known | Event handler |
| 0x001A9A64 | `HandleVideoPlayPause` | Known | Event handler |
| 0x001A9A7C | `HandleVideoNext` | Known | Event handler |
| 0x001A9A8C | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x001A9AA8 | `HandleVideoPrevious` | Known | Event handler |
| 0x001A9ABC | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x001A9CC4 | `HandleVideoStop` | Known | Event handler |
| 0x001A9CD4 | `HandleVideoPlay` | Known | Event handler |
| 0x001A9CE4 | `HandleVideoPause` | Known | Event handler |
| 0x001A9CF8 | `HandleVideoFFDown` | Known | Event handler |
| 0x001A9D0C | `HandleVideoFFUp` | Known | Event handler |
| 0x001A9D1C | `HandleVideoRewDown` | Known | Event handler |
| 0x001A9D30 | `HandleVideoRewUp` | Known | Event handler |
| 0x001A9D44 | `HandleVideoNextChapter` | Known | Event handler |
| 0x001A9D5C | `HandleVideoPrevChapter` | Known | Event handler |
| 0x001A9D74 | `HandleVideoNextFrame` | Known | Event handler |
| 0x001A9D8C | `HandleVideoPrevFrame` | Known | Event handler |
| 0x001A9DA4 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x001B67C0 | `HandleMainMenu` | Known | Event handler |
| 0x001BAD04 | `HandlePowerSongSelected` | Known | Event handler |
| 0x001BAD20 | `HandlePowerSongChosen` | Known | Event handler |
| 0x001BAD38 | `HandleSelectTimedWorkout` | Known | Event handler |
| 0x001C1274 | `HandleSelect` | Known | Event handler |
| 0x001C151C | `HandleMusicMenu` | Known | Event handler |
| 0x001C17DC | `HandleSelect` | Known | Event handler |
| 0x001C1B60 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x001C1B78 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x001C1B98 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x001C1BBC | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x001C1BD8 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x001C2074 | `HandleWheel` | Known | Event handler |
| 0x001C2084 | `HandlePlayPause` | Known | Event handler |
| 0x001C2094 | `HandleSelectDown` | Known | Event handler |
| 0x001C20A8 | `HandleNext` | Known | Event handler |
| 0x001C20B4 | `HandlePrevious` | Known | Event handler |
| 0x001C20C4 | `HandleNextPushAndHold` | Known | Event handler |
| 0x001C20DC | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x001CDAE0 | `HandleFrequencyChosen` | Known | Event handler |
| 0x001CDAF8 | `HandleDateChosen` | Known | Event handler |
| 0x001CDB0C | `HandleTimeChosen` | Known | Event handler |
| 0x001CDB20 | `HandleSoundChosen` | Known | Event handler |
| 0x001CDB34 | `HandleLabelChosen` | Known | Event handler |
| 0x001CDB48 | `HandleDeleteChosen` | Known | Event handler |
| 0x001CEC28 | `HandleSelect` | Known | Event handler |
| 0x001D3544 | `HandlePrev` | Known | Event handler |
| 0x001D3554 | `HandleNext` | Known | Event handler |
| 0x001D3560 | `HandlePlayPause` | Known | Event handler |
| 0x001DAAA4 | `HandleNextContact` | Known | Event handler |
| 0x001DAABC | `HandlePreviousContact` | Known | Event handler |
| 0x001E25B8 | `HandleItemSelected` | Known | Event handler |
| 0x001E27B0 | `HandleRadioRegion` | Known | Event handler |
| 0x001E2998 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x001E6998 | `HandlePlayPause` | Known | Event handler |
| 0x001EA330 | `HandleDelete` | Known | Event handler |
| 0x001EA344 | `HandleSelectLozinch` | Known | Event handler |
| 0x001EA5EC | `HandleSelect` | Known | Event handler |
| 0x001EA8B8 | `HandleTVOutChanged` | Known | Event handler |
| 0x001EA8D0 | `HandleTVSignalChanged` | Known | Event handler |
| 0x001EA8E8 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x001EA908 | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x001EA928 | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x001EA94C | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x001EA96C | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x001ED4C8 | `HandleSelectKey` | Known | Event handler |
| 0x001ED670 | `HandleSelect` | Known | Event handler |
| 0x001EE3EC | `HandlePlayPause` | Known | Event handler |
| 0x001EE400 | `HandleWheel` | Known | Event handler |
| 0x001EE40C | `HandleWheelRating` | Known | Event handler |
| 0x001EE420 | `HandleWheelScrub` | Known | Event handler |
| 0x001EE434 | `HandleWheelVolume` | Known | Event handler |
| 0x001EE4F4 | `HandleMenuKey` | Known | Event handler |
| 0x001EE560 | `HandleMenuLongpress` | Known | Event handler |
| 0x001EE574 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x001EF17C | `HandleSelect` | Known | Event handler |
| 0x001EFA74 | `HandleLeaveAlarm` | Known | Event handler |
| 0x001F0964 | `HandleSelect` | Known | Event handler |
| 0x001F0978 | `HandleHilite` | Known | Event handler |
| 0x001F0988 | `HandlePlayPause` | Known | Event handler |
| 0x001F0998 | `HandleAddToOTG` | Known | Event handler |
| 0x001F09A8 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x001F36D8 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x001F3EE8 | `HandleSelect` | Known | Event handler |
| 0x001F3EFC | `HandleWheel` | Known | Event handler |
| 0x001F3F08 | `HandleWheelProgress` | Known | Event handler |
| 0x001F3F1C | `HandleSelectProgress` | Known | Event handler |
| 0x001F3F34 | `HandleSelectVolume` | Known | Event handler |
| 0x001F3F48 | `HandleSelectScrub` | Known | Event handler |
| 0x001F3F5C | `HandleSelectRating` | Known | Event handler |
| 0x001F3F70 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x001F3F88 | `HandleSelectChapterArt` | Known | Event handler |
| 0x001F3FA0 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x001F3FBC | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x001F3FD8 | `HandleWheelBrightness` | Known | Event handler |
| 0x001F4120 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x001F5CB0 | `HandleSelect` | Known | Event handler |
| 0x001F5CC0 | `HandleSelectRating` | Known | Event handler |
| 0x001F5CD4 | `HandleSelectProgress` | Known | Event handler |
| 0x001F5CEC | `HandleWheelProgress` | Known | Event handler |
| 0x001F5D00 | `HandleSelectScrub` | Known | Event handler |
| 0x001F5D14 | `HandleWheelBrightness` | Known | Event handler |
| 0x001F5D2C | `HandleNextTrackGotoScrub` | Known | Event handler |
| 0x001F5D48 | `HandlePrevTrackGotoScrub` | Known | Event handler |
| 0x001F5D64 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x001FBDE0 | `HandleLanguage` | Known | Event handler |
| 0x001FBDF0 | `HandleResetAllSettings` | Known | Event handler |
| 0x001FBE08 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x001FC774 | `HandleSelect` | Known | Event handler |
| 0x001FC9A4 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x001FFF40 | `HandleSelect` | Known | Event handler |
| 0x002000DC | `HandleSelect` | Known | Event handler |
| 0x0020037C | `HandleNextDay` | Known | Event handler |
| 0x00200390 | `HandlePreviousDay` | Known | Event handler |
| 0x00200B94 | `HandleMusicHilited` | Known | Event handler |
| 0x00200BAC | `HandleVideosHilited` | Known | Event handler |
| 0x00200BC0 | `HandlePodcastsHilited` | Known | Event handler |
| 0x00200BD8 | `HandleGenericHilited` | Known | Event handler |
| 0x00200BF0 | `HandlePhotosHilited` | Known | Event handler |
| 0x00200C04 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x00200C1C | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x00200C38 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00200C50 | `HandleArtistsHilited` | Known | Event handler |
| 0x00200C68 | `HandleGenresHilited` | Known | Event handler |
| 0x00200C7C | `HandleAlbumsHilited` | Known | Event handler |
| 0x00200C90 | `HandleCompilationsHilited` | Known | Event handler |
| 0x00200E64 | `HandleComposersHilited` | Known | Event handler |
| 0x00200E7C | `HandleSongsHilited` | Known | Event handler |
| 0x00200E90 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x00200EA8 | `HandleTVShowsHilited` | Known | Event handler |
| 0x00200EC0 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00200EDC | `HandleMoviesHilited` | Known | Event handler |
| 0x00200EF0 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00200F0C | `HandleRentalsHilited` | Known | Event handler |
| 0x00200F24 | `HandleMusicSelected` | Known | Event handler |
| 0x00200F38 | `HandleVideosSelected` | Known | Event handler |
| 0x00200F50 | `HandlePodcastsSelected` | Known | Event handler |
| 0x00201120 | `HandlePhotosSelected` | Known | Event handler |
| 0x00201138 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00201150 | `HandleSongsSelected` | Known | Event handler |
| 0x00201164 | `HandleAlbumsSelected` | Known | Event handler |
| 0x0020117C | `HandleCompilationsSelected` | Known | Event handler |
| 0x00201198 | `HandleArtistsSelected` | Known | Event handler |
| 0x002011B0 | `HandleGenresSelected` | Known | Event handler |
| 0x002011C8 | `HandleComposersSelected` | Known | Event handler |
| 0x002011E0 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x002011FC | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00201218 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x002013E0 | `HandleNowPlaying` | Known | Event handler |
| 0x002013F4 | `HandleTVShowsSelected` | Known | Event handler |
| 0x0020140C | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00201428 | `HandleMoviesSelected` | Known | Event handler |
| 0x00201440 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00201460 | `HandleRentalsSelected` | Known | Event handler |
| 0x00201478 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00201490 | `HandleLock` | Known | Event handler |
| 0x0020149C | `HandleBacklightSelected` | Known | Event handler |
| 0x002014B4 | `HandleSleepSelected` | Known | Event handler |
| 0x002014C8 | `HandleNikePlusSelected` | Known | Event handler |
| 0x00203CB4 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x002042B8 | `HandleWheel` | Known | Event handler |
| 0x00205788 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x002059E0 | `HandleNextDay` | Known | Event handler |
| 0x002059F4 | `HandlePreviousDay` | Known | Event handler |
| 0x00205C3C | `HandleSelect` | Known | Event handler |
| 0x00205ED8 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00208804 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x00208820 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x00209788 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00209E68 | `HandleSelect` | Known | Event handler |
| 0x0020A534 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x0023F7E4 | `HandleDeleteClock` | Known | Event handler |
| 0x0023F7FC | `HandleSelectClock` | Known | Event handler |
| 0x0023F810 | `HandleHilited` | Known | Event handler |
| 0x0023F820 | `HandleWheel` | Known | Event handler |
| 0x0023F82C | `HandleSelectLozinch` | Known | Event handler |
| 0x003AB426 | `HandleAudioFFDown` | Known | Event handler |
| 0x003AB44F | `HandleAudioFFUp` | Known | Event handler |
| 0x003AB47A | `HandleAudioMute` | Known | Event handler |
| 0x003AB4AD | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x003AB4E2 | `HandleAudioNext` | Known | Event handler |
| 0x003AB512 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x003AB549 | `HandleAudioNextChapter` | Known | Event handler |
| 0x003AB583 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x003AB5B7 | `HandleAudioPause` | Known | Event handler |
| 0x003AB5E3 | `HandleAudioPlay` | Known | Event handler |
| 0x003AB611 | `HandleAudioPlayPause` | Known | Event handler |
| 0x003AB649 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x003AB682 | `HandleAudioPrevious` | Known | Event handler |
| 0x003AB6B6 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x003AB6ED | `HandleAudioPrevChapter` | Known | Event handler |
| 0x003AB727 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x003AB75C | `HandleAudioRepeat` | Known | Event handler |
| 0x003AB788 | `HandleAudioRewDown` | Known | Event handler |
| 0x003AB7B3 | `HandleAudioRewUp` | Known | Event handler |
| 0x003AB7E2 | `HandleAudioShuffle` | Known | Event handler |
| 0x003AB810 | `HandleAudioStop` | Known | Event handler |
| 0x003AB841 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x003AB876 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x003AB8AD | `HandleAudioVolumeUp` | Known | Event handler |
| 0x003AB8DE | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x003AB997 | `HandleNextPressAndHold` | Known | Event handler |
| 0x003AB9C8 | `HandleNext` | Known | Event handler |
| 0x003ABA00 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x003ABA3B | `HandlePlayPause` | Known | Event handler |
| 0x003ABA6F | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x003ABAA4 | `HandlePrevious` | Known | Event handler |
| 0x003ABB31 | `HandleRemoteBacklight` | Known | Event handler |
| 0x003ABB69 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x003ABBA3 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x003ABBDC | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x003ABC11 | `HandleRemoteEvent` | Known | Event handler |
| 0x003ABC3D | `HandleRemoteFFDown` | Known | Event handler |
| 0x003ABC68 | `HandleRemoteFFUp` | Known | Event handler |
| 0x003ABC95 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x003ABCC4 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x003ABCF3 | `HandleRemoteMute` | Known | Event handler |
| 0x003ABD25 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x003ABD5E | `HandleRemoteNextChapter` | Known | Event handler |
| 0x003ABD9A | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x003ABDDA | `HandleRemoteOff` | Known | Event handler |
| 0x003ABE03 | `HandleRemoteOff` | Known | Event handler |
| 0x003ABE2D | `HandleRemoteOn` | Known | Event handler |
| 0x003ABE59 | `HandleRemotePause` | Known | Event handler |
| 0x003ABE87 | `HandleRemotePlay` | Known | Event handler |
| 0x003ABEC5 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x003ABF06 | `HandleRemotePlayPause` | Known | Event handler |
| 0x003ABF3D | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x003ABF76 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x003ABFB2 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x003ABFE9 | `HandleRemoteRepeat` | Known | Event handler |
| 0x003AC017 | `HandleRemoteRewDown` | Known | Event handler |
| 0x003AC044 | `HandleRemoteRewUp` | Known | Event handler |
| 0x003AC074 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x003AC0A7 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x003AC0DB | `HandleRemoteShuffle` | Known | Event handler |
| 0x003AC10B | `HandleRemoteStop` | Known | Event handler |
| 0x003AC13B | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x003AC170 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x003AC1A8 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x003AC1DF | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x003AC218 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x003AC24B | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x003AC280 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x003AC2B3 | `HandleVideoFFDown` | Known | Event handler |
| 0x003AC2DC | `HandleVideoFFUp` | Known | Event handler |
| 0x003AC30F | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x003AC344 | `HandleVideoNext` | Known | Event handler |
| 0x003AC376 | `HandleVideoNextChapter` | Known | Event handler |
| 0x003AC3AD | `HandleVideoNextFrame` | Known | Event handler |
| 0x003AC3DE | `HandleVideoPause` | Known | Event handler |
| 0x003AC40A | `HandleVideoPlay` | Known | Event handler |
| 0x003AC438 | `HandleVideoPlayPause` | Known | Event handler |
| 0x003AC470 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x003AC4A9 | `HandleVideoPrevious` | Known | Event handler |
| 0x003AC4DF | `HandleVideoPrevChapter` | Known | Event handler |
| 0x003AC516 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x003AC545 | `HandleVideoRewDown` | Known | Event handler |
| 0x003AC570 | `HandleVideoRewUp` | Known | Event handler |
| 0x003AC59C | `HandleVideoStop` | Known | Event handler |
| 0x006A297E | `HandleAddressBook` | Known | Event handler |
| 0x006A2F12 | `HandleSelect` | Known | Event handler |
| 0x006A2F4D | `HandleHilite` | Known | Event handler |
| 0x006A2FCE | `HandleSelectRegion` | Known | Event handler |
| 0x006A306E | `HandleSelectRegion` | Known | Event handler |
| 0x006A310A | `HandleSelectRegion` | Known | Event handler |
| 0x006A31AE | `HandleSelectRegion` | Known | Event handler |
| 0x006A3254 | `HandleSelectRegion` | Known | Event handler |
| 0x006A32F4 | `HandleSelectRegion` | Known | Event handler |
| 0x006A33A0 | `HandleSelectRegion` | Known | Event handler |
| 0x006A3442 | `HandleSelectRegion` | Known | Event handler |
| 0x006A34F2 | `HandleSelectCity` | Known | Event handler |
| 0x006A355E | `HandleHighlightCity` | Known | Event handler |
| 0x006A3597 | `HandleSelectCity` | Known | Event handler |
| 0x006A3603 | `HandleHighlightCity` | Known | Event handler |
| 0x006A363C | `HandleSelectCity` | Known | Event handler |
| 0x006A36A8 | `HandleHighlightCity` | Known | Event handler |
| 0x006A36E1 | `HandleSelectCity` | Known | Event handler |
| 0x006A374D | `HandleHighlightCity` | Known | Event handler |
| 0x006A3786 | `HandleSelectCity` | Known | Event handler |
| 0x006A37F2 | `HandleHighlightCity` | Known | Event handler |
| 0x006A382B | `HandleSelectCity` | Known | Event handler |
| 0x006A3897 | `HandleHighlightCity` | Known | Event handler |
| 0x006A38D0 | `HandleSelectCity` | Known | Event handler |
| 0x006A393C | `HandleHighlightCity` | Known | Event handler |
| 0x006A3975 | `HandleSelectCity` | Known | Event handler |
| 0x006A39E1 | `HandleHighlightCity` | Known | Event handler |
| 0x006A3A1A | `HandleSelectCity` | Known | Event handler |
| 0x006A3A86 | `HandleHighlightCity` | Known | Event handler |
| 0x006A3ABF | `HandleSelectCity` | Known | Event handler |
| 0x006A3B2B | `HandleHighlightCity` | Known | Event handler |
| 0x006A3B64 | `HandleSelectCity` | Known | Event handler |
| 0x006A3BD0 | `HandleHighlightCity` | Known | Event handler |
| 0x006A3C09 | `HandleSelectCity` | Known | Event handler |
| 0x006A3C75 | `HandleHighlightCity` | Known | Event handler |
| 0x006A3CAE | `HandleSelectCity` | Known | Event handler |
| 0x006A3D1A | `HandleHighlightCity` | Known | Event handler |
| 0x006A3D53 | `HandleSelectCity` | Known | Event handler |
| 0x006A3DBF | `HandleHighlightCity` | Known | Event handler |
| 0x006A3DF8 | `HandleSelectCity` | Known | Event handler |
| 0x006A3E64 | `HandleHighlightCity` | Known | Event handler |
| 0x006A3E9D | `HandleSelectCity` | Known | Event handler |
| 0x006A3F09 | `HandleHighlightCity` | Known | Event handler |
| 0x006A3F42 | `HandleSelectCity` | Known | Event handler |
| 0x006A3FAE | `HandleHighlightCity` | Known | Event handler |
| 0x006A3FE7 | `HandleSelectCity` | Known | Event handler |
| 0x006A4053 | `HandleHighlightCity` | Known | Event handler |
| 0x006A408C | `HandleSelectCity` | Known | Event handler |
| 0x006A40F8 | `HandleHighlightCity` | Known | Event handler |
| 0x006A4131 | `HandleSelectCity` | Known | Event handler |
| 0x006A419D | `HandleHighlightCity` | Known | Event handler |
| 0x006A41D6 | `HandleSelectCity` | Known | Event handler |
| 0x006A4242 | `HandleHighlightCity` | Known | Event handler |
| 0x006A427B | `HandleSelectCity` | Known | Event handler |
| 0x006A42E7 | `HandleHighlightCity` | Known | Event handler |
| 0x006A4320 | `HandleSelectCity` | Known | Event handler |
| 0x006A438C | `HandleHighlightCity` | Known | Event handler |
| 0x006A43C5 | `HandleSelectCity` | Known | Event handler |
| 0x006A4431 | `HandleHighlightCity` | Known | Event handler |
| 0x006A446A | `HandleSelectCity` | Known | Event handler |
| 0x006A44D6 | `HandleHighlightCity` | Known | Event handler |
| 0x006A450F | `HandleSelectCity` | Known | Event handler |
| 0x006A457B | `HandleHighlightCity` | Known | Event handler |
| 0x006A45B4 | `HandleSelectCity` | Known | Event handler |
| 0x006A4620 | `HandleHighlightCity` | Known | Event handler |
| 0x006A4659 | `HandleSelectCity` | Known | Event handler |
| 0x006A46C5 | `HandleHighlightCity` | Known | Event handler |
| 0x006A46FE | `HandleSelectCity` | Known | Event handler |
| 0x006A476A | `HandleHighlightCity` | Known | Event handler |
| 0x006A47A3 | `HandleSelectCity` | Known | Event handler |
| 0x006A480F | `HandleHighlightCity` | Known | Event handler |
| 0x006A4848 | `HandleSelectCity` | Known | Event handler |
| 0x006A48B4 | `HandleHighlightCity` | Known | Event handler |
| 0x006A48F2 | `HandleSelectCity` | Known | Event handler |
| 0x006A495E | `HandleHighlightCity` | Known | Event handler |
| 0x006A4997 | `HandleSelectCity` | Known | Event handler |
| 0x006A4A03 | `HandleHighlightCity` | Known | Event handler |
| 0x006A4A3C | `HandleSelectCity` | Known | Event handler |
| 0x006A4AA8 | `HandleHighlightCity` | Known | Event handler |
| 0x006A4AE1 | `HandleSelectCity` | Known | Event handler |
| 0x006A4B4D | `HandleHighlightCity` | Known | Event handler |
| 0x006A4B86 | `HandleSelectCity` | Known | Event handler |
| 0x006A4BF2 | `HandleHighlightCity` | Known | Event handler |
| 0x006A4C2B | `HandleSelectCity` | Known | Event handler |
| 0x006A4C97 | `HandleHighlightCity` | Known | Event handler |
| 0x006A4CD0 | `HandleSelectCity` | Known | Event handler |
| 0x006A4D3C | `HandleHighlightCity` | Known | Event handler |
| 0x006A4D75 | `HandleSelectCity` | Known | Event handler |
| 0x006A4DE1 | `HandleHighlightCity` | Known | Event handler |
| 0x006A4E1A | `HandleSelectCity` | Known | Event handler |
| 0x006A4E86 | `HandleHighlightCity` | Known | Event handler |
| 0x006A4EBF | `HandleSelectCity` | Known | Event handler |
| 0x006A4F2B | `HandleHighlightCity` | Known | Event handler |
| 0x006A4F64 | `HandleSelectCity` | Known | Event handler |
| 0x006A4FD0 | `HandleHighlightCity` | Known | Event handler |
| 0x006A5009 | `HandleSelectCity` | Known | Event handler |
| 0x006A5075 | `HandleHighlightCity` | Known | Event handler |
| 0x006A50AE | `HandleSelectCity` | Known | Event handler |
| 0x006A511A | `HandleHighlightCity` | Known | Event handler |
| 0x006A5153 | `HandleSelectCity` | Known | Event handler |
| 0x006A51BF | `HandleHighlightCity` | Known | Event handler |
| 0x006A51F8 | `HandleSelectCity` | Known | Event handler |
| 0x006A5264 | `HandleHighlightCity` | Known | Event handler |
| 0x006A529D | `HandleSelectCity` | Known | Event handler |
| 0x006A5309 | `HandleHighlightCity` | Known | Event handler |
| 0x006A5342 | `HandleSelectCity` | Known | Event handler |
| 0x006A53AE | `HandleHighlightCity` | Known | Event handler |
| 0x006A53E7 | `HandleSelectCity` | Known | Event handler |
| 0x006A5453 | `HandleHighlightCity` | Known | Event handler |
| 0x006A548C | `HandleSelectCity` | Known | Event handler |
| 0x006A54F8 | `HandleHighlightCity` | Known | Event handler |
| 0x006A5531 | `HandleSelectCity` | Known | Event handler |
| 0x006A559D | `HandleHighlightCity` | Known | Event handler |
| 0x006A55D6 | `HandleSelectCity` | Known | Event handler |
| 0x006A5642 | `HandleHighlightCity` | Known | Event handler |
| 0x006A567B | `HandleSelectCity` | Known | Event handler |
| 0x006A56E7 | `HandleHighlightCity` | Known | Event handler |
| 0x006A5720 | `HandleSelectCity` | Known | Event handler |
| 0x006A578C | `HandleHighlightCity` | Known | Event handler |
| 0x006A57C5 | `HandleSelectCity` | Known | Event handler |
| 0x006A5831 | `HandleHighlightCity` | Known | Event handler |
| 0x006A586A | `HandleSelectCity` | Known | Event handler |
| 0x006A58D6 | `HandleHighlightCity` | Known | Event handler |
| 0x006A590F | `HandleSelectCity` | Known | Event handler |
| 0x006A597B | `HandleHighlightCity` | Known | Event handler |
| 0x006A59B4 | `HandleSelectCity` | Known | Event handler |
| 0x006A5A20 | `HandleHighlightCity` | Known | Event handler |
| 0x006A5A59 | `HandleSelectCity` | Known | Event handler |
| 0x006A5AC5 | `HandleHighlightCity` | Known | Event handler |
| 0x006A5AFE | `HandleSelectCity` | Known | Event handler |
| 0x006A5B6A | `HandleHighlightCity` | Known | Event handler |
| 0x006A5BA3 | `HandleSelectCity` | Known | Event handler |
| 0x006A5C0F | `HandleHighlightCity` | Known | Event handler |
| 0x006A5C48 | `HandleSelectCity` | Known | Event handler |
| 0x006A5CB4 | `HandleHighlightCity` | Known | Event handler |
| 0x006A5CED | `HandleSelectCity` | Known | Event handler |
| 0x006A5D59 | `HandleHighlightCity` | Known | Event handler |
| 0x006A5D92 | `HandleSelectCity` | Known | Event handler |
| 0x006A5DFE | `HandleHighlightCity` | Known | Event handler |
| 0x006A5E37 | `HandleSelectCity` | Known | Event handler |
| 0x006A5EA3 | `HandleHighlightCity` | Known | Event handler |
| 0x006A5EDC | `HandleSelectCity` | Known | Event handler |
| 0x006A5F48 | `HandleHighlightCity` | Known | Event handler |
| 0x006A5F81 | `HandleSelectCity` | Known | Event handler |
| 0x006A5FED | `HandleHighlightCity` | Known | Event handler |
| 0x006A6026 | `HandleSelectCity` | Known | Event handler |
| 0x006A6092 | `HandleHighlightCity` | Known | Event handler |
| 0x006A60CB | `HandleSelectCity` | Known | Event handler |
| 0x006A6137 | `HandleHighlightCity` | Known | Event handler |
| 0x006A6170 | `HandleSelectCity` | Known | Event handler |
| 0x006A61DC | `HandleHighlightCity` | Known | Event handler |
| 0x006A6215 | `HandleSelectCity` | Known | Event handler |
| 0x006A6281 | `HandleHighlightCity` | Known | Event handler |
| 0x006A62BA | `HandleSelectCity` | Known | Event handler |
| 0x006A6326 | `HandleHighlightCity` | Known | Event handler |
| 0x006A635F | `HandleSelectCity` | Known | Event handler |
| 0x006A63CB | `HandleHighlightCity` | Known | Event handler |
| 0x006A6404 | `HandleSelectCity` | Known | Event handler |
| 0x006A6470 | `HandleHighlightCity` | Known | Event handler |
| 0x006A64A9 | `HandleSelectCity` | Known | Event handler |
| 0x006A6515 | `HandleHighlightCity` | Known | Event handler |
| 0x006A654E | `HandleSelectCity` | Known | Event handler |
| 0x006A65BA | `HandleHighlightCity` | Known | Event handler |
| 0x006A65F3 | `HandleSelectCity` | Known | Event handler |
| 0x006A665F | `HandleHighlightCity` | Known | Event handler |
| 0x006A6698 | `HandleSelectCity` | Known | Event handler |
| 0x006A6704 | `HandleHighlightCity` | Known | Event handler |
| 0x006A673D | `HandleSelectCity` | Known | Event handler |
| 0x006A67A9 | `HandleHighlightCity` | Known | Event handler |
| 0x006A67E2 | `HandleSelectCity` | Known | Event handler |
| 0x006A684E | `HandleHighlightCity` | Known | Event handler |
| 0x006A6887 | `HandleSelectCity` | Known | Event handler |
| 0x006A68F3 | `HandleHighlightCity` | Known | Event handler |
| 0x006A692C | `HandleSelectCity` | Known | Event handler |
| 0x006A6998 | `HandleHighlightCity` | Known | Event handler |
| 0x006A69D1 | `HandleSelectCity` | Known | Event handler |
| 0x006A6A3D | `HandleHighlightCity` | Known | Event handler |
| 0x006A6A76 | `HandleSelectCity` | Known | Event handler |
| 0x006A6AE2 | `HandleHighlightCity` | Known | Event handler |
| 0x006A6B1B | `HandleSelectCity` | Known | Event handler |
| 0x006A6B87 | `HandleHighlightCity` | Known | Event handler |
| 0x006A6BC0 | `HandleSelectCity` | Known | Event handler |
| 0x006A6C2C | `HandleHighlightCity` | Known | Event handler |
| 0x006A6C65 | `HandleSelectCity` | Known | Event handler |
| 0x006A6CD1 | `HandleHighlightCity` | Known | Event handler |
| 0x006A6D0A | `HandleSelectCity` | Known | Event handler |
| 0x006A6D76 | `HandleHighlightCity` | Known | Event handler |
| 0x006A6DB6 | `HandleSelectCity` | Known | Event handler |
| 0x006A6E22 | `HandleHighlightCity` | Known | Event handler |
| 0x006A6E5B | `HandleSelectCity` | Known | Event handler |
| 0x006A6EC7 | `HandleHighlightCity` | Known | Event handler |
| 0x006A6F00 | `HandleSelectCity` | Known | Event handler |
| 0x006A6F6C | `HandleHighlightCity` | Known | Event handler |
| 0x006A6FAA | `HandleSelectCity` | Known | Event handler |
| 0x006A7016 | `HandleHighlightCity` | Known | Event handler |
| 0x006A704F | `HandleSelectCity` | Known | Event handler |
| 0x006A70BB | `HandleHighlightCity` | Known | Event handler |
| 0x006A70F4 | `HandleSelectCity` | Known | Event handler |
| 0x006A7160 | `HandleHighlightCity` | Known | Event handler |
| 0x006A7199 | `HandleSelectCity` | Known | Event handler |
| 0x006A7205 | `HandleHighlightCity` | Known | Event handler |
| 0x006A723E | `HandleSelectCity` | Known | Event handler |
| 0x006A72AA | `HandleHighlightCity` | Known | Event handler |
| 0x006A72E3 | `HandleSelectCity` | Known | Event handler |
| 0x006A734F | `HandleHighlightCity` | Known | Event handler |
| 0x006A7388 | `HandleSelectCity` | Known | Event handler |
| 0x006A73F4 | `HandleHighlightCity` | Known | Event handler |
| 0x006A742D | `HandleSelectCity` | Known | Event handler |
| 0x006A7499 | `HandleHighlightCity` | Known | Event handler |
| 0x006A74D6 | `HandleSelectCity` | Known | Event handler |
| 0x006A7542 | `HandleHighlightCity` | Known | Event handler |
| 0x006A757B | `HandleSelectCity` | Known | Event handler |
| 0x006A75E7 | `HandleHighlightCity` | Known | Event handler |
| 0x006A7620 | `HandleSelectCity` | Known | Event handler |
| 0x006A768C | `HandleHighlightCity` | Known | Event handler |
| 0x006A76C5 | `HandleSelectCity` | Known | Event handler |
| 0x006A7731 | `HandleHighlightCity` | Known | Event handler |
| 0x006A776A | `HandleSelectCity` | Known | Event handler |
| 0x006A77D6 | `HandleHighlightCity` | Known | Event handler |
| 0x006A780F | `HandleSelectCity` | Known | Event handler |
| 0x006A787B | `HandleHighlightCity` | Known | Event handler |
| 0x006A78B4 | `HandleSelectCity` | Known | Event handler |
| 0x006A7920 | `HandleHighlightCity` | Known | Event handler |
| 0x006A7959 | `HandleSelectCity` | Known | Event handler |
| 0x006A79C5 | `HandleHighlightCity` | Known | Event handler |
| 0x006A79FE | `HandleSelectCity` | Known | Event handler |
| 0x006A7A6A | `HandleHighlightCity` | Known | Event handler |
| 0x006A7AA3 | `HandleSelectCity` | Known | Event handler |
| 0x006A7B0F | `HandleHighlightCity` | Known | Event handler |
| 0x006A7B48 | `HandleSelectCity` | Known | Event handler |
| 0x006A7BB4 | `HandleHighlightCity` | Known | Event handler |
| 0x006A7BED | `HandleSelectCity` | Known | Event handler |
| 0x006A7C59 | `HandleHighlightCity` | Known | Event handler |
| 0x006A7C92 | `HandleSelectCity` | Known | Event handler |
| 0x006A7CFE | `HandleHighlightCity` | Known | Event handler |
| 0x006A7D37 | `HandleSelectCity` | Known | Event handler |
| 0x006A7DA3 | `HandleHighlightCity` | Known | Event handler |
| 0x006A7DDC | `HandleSelectCity` | Known | Event handler |
| 0x006A7E48 | `HandleHighlightCity` | Known | Event handler |
| 0x006A7E81 | `HandleSelectCity` | Known | Event handler |
| 0x006A7EED | `HandleHighlightCity` | Known | Event handler |
| 0x006A7F26 | `HandleSelectCity` | Known | Event handler |
| 0x006A7F92 | `HandleHighlightCity` | Known | Event handler |
| 0x006A7FCB | `HandleSelectCity` | Known | Event handler |
| 0x006A8037 | `HandleHighlightCity` | Known | Event handler |
| 0x006A8070 | `HandleSelectCity` | Known | Event handler |
| 0x006A80DC | `HandleHighlightCity` | Known | Event handler |
| 0x006A8115 | `HandleSelectCity` | Known | Event handler |
| 0x006A8181 | `HandleHighlightCity` | Known | Event handler |
| 0x006A81BA | `HandleSelectCity` | Known | Event handler |
| 0x006A8226 | `HandleHighlightCity` | Known | Event handler |
| 0x006A825F | `HandleSelectCity` | Known | Event handler |
| 0x006A82CB | `HandleHighlightCity` | Known | Event handler |
| 0x006A8304 | `HandleSelectCity` | Known | Event handler |
| 0x006A8370 | `HandleHighlightCity` | Known | Event handler |
| 0x006A83A9 | `HandleSelectCity` | Known | Event handler |
| 0x006A8415 | `HandleHighlightCity` | Known | Event handler |
| 0x006A844E | `HandleSelectCity` | Known | Event handler |
| 0x006A84BA | `HandleHighlightCity` | Known | Event handler |
| 0x006A84F3 | `HandleSelectCity` | Known | Event handler |
| 0x006A855F | `HandleHighlightCity` | Known | Event handler |
| 0x006A8598 | `HandleSelectCity` | Known | Event handler |
| 0x006A8604 | `HandleHighlightCity` | Known | Event handler |
| 0x006A863D | `HandleSelectCity` | Known | Event handler |
| 0x006A86A9 | `HandleHighlightCity` | Known | Event handler |
| 0x006A86E2 | `HandleSelectCity` | Known | Event handler |
| 0x006A874E | `HandleHighlightCity` | Known | Event handler |
| 0x006A8787 | `HandleSelectCity` | Known | Event handler |
| 0x006A87F3 | `HandleHighlightCity` | Known | Event handler |
| 0x006A882C | `HandleSelectCity` | Known | Event handler |
| 0x006A8898 | `HandleHighlightCity` | Known | Event handler |
| 0x006A88D1 | `HandleSelectCity` | Known | Event handler |
| 0x006A893D | `HandleHighlightCity` | Known | Event handler |
| 0x006A8976 | `HandleSelectCity` | Known | Event handler |
| 0x006A89E2 | `HandleHighlightCity` | Known | Event handler |
| 0x006A8A1B | `HandleSelectCity` | Known | Event handler |
| 0x006A8A87 | `HandleHighlightCity` | Known | Event handler |
| 0x006A8AC6 | `HandleSelectCity` | Known | Event handler |
| 0x006A8B32 | `HandleHighlightCity` | Known | Event handler |
| 0x006A8B6B | `HandleSelectCity` | Known | Event handler |
| 0x006A8BD7 | `HandleHighlightCity` | Known | Event handler |
| 0x006A8C10 | `HandleSelectCity` | Known | Event handler |
| 0x006A8C7C | `HandleHighlightCity` | Known | Event handler |
| 0x006A8CB5 | `HandleSelectCity` | Known | Event handler |
| 0x006A8D21 | `HandleHighlightCity` | Known | Event handler |
| 0x006A8D5A | `HandleSelectCity` | Known | Event handler |
| 0x006A8DC6 | `HandleHighlightCity` | Known | Event handler |
| 0x006A8DFF | `HandleSelectCity` | Known | Event handler |
| 0x006A8E6B | `HandleHighlightCity` | Known | Event handler |
| 0x006A8EA4 | `HandleSelectCity` | Known | Event handler |
| 0x006A8F10 | `HandleHighlightCity` | Known | Event handler |
| 0x006A8F49 | `HandleSelectCity` | Known | Event handler |
| 0x006A8FB5 | `HandleHighlightCity` | Known | Event handler |
| 0x006A8FEE | `HandleSelectCity` | Known | Event handler |
| 0x006A905A | `HandleHighlightCity` | Known | Event handler |
| 0x006A9093 | `HandleSelectCity` | Known | Event handler |
| 0x006A90FF | `HandleHighlightCity` | Known | Event handler |
| 0x006A9138 | `HandleSelectCity` | Known | Event handler |
| 0x006A91A4 | `HandleHighlightCity` | Known | Event handler |
| 0x006A91DD | `HandleSelectCity` | Known | Event handler |
| 0x006A9249 | `HandleHighlightCity` | Known | Event handler |
| 0x006A9282 | `HandleSelectCity` | Known | Event handler |
| 0x006A92EE | `HandleHighlightCity` | Known | Event handler |
| 0x006A9327 | `HandleSelectCity` | Known | Event handler |
| 0x006A9393 | `HandleHighlightCity` | Known | Event handler |
| 0x006A93CC | `HandleSelectCity` | Known | Event handler |
| 0x006A9438 | `HandleHighlightCity` | Known | Event handler |
| 0x006A9471 | `HandleSelectCity` | Known | Event handler |
| 0x006A94DD | `HandleHighlightCity` | Known | Event handler |
| 0x006A9516 | `HandleSelectCity` | Known | Event handler |
| 0x006A9582 | `HandleHighlightCity` | Known | Event handler |
| 0x006A95BB | `HandleSelectCity` | Known | Event handler |
| 0x006A9627 | `HandleHighlightCity` | Known | Event handler |
| 0x006A9660 | `HandleSelectCity` | Known | Event handler |
| 0x006A96CC | `HandleHighlightCity` | Known | Event handler |
| 0x006A9705 | `HandleSelectCity` | Known | Event handler |
| 0x006A9771 | `HandleHighlightCity` | Known | Event handler |
| 0x006A97AA | `HandleSelectCity` | Known | Event handler |
| 0x006A9816 | `HandleHighlightCity` | Known | Event handler |
| 0x006A984F | `HandleSelectCity` | Known | Event handler |
| 0x006A98BB | `HandleHighlightCity` | Known | Event handler |
| 0x006A98F4 | `HandleSelectCity` | Known | Event handler |
| 0x006A9960 | `HandleHighlightCity` | Known | Event handler |
| 0x006A9999 | `HandleSelectCity` | Known | Event handler |
| 0x006A9A05 | `HandleHighlightCity` | Known | Event handler |
| 0x006A9A3E | `HandleSelectCity` | Known | Event handler |
| 0x006A9AAA | `HandleHighlightCity` | Known | Event handler |
| 0x006A9AE3 | `HandleSelectCity` | Known | Event handler |
| 0x006A9B4F | `HandleHighlightCity` | Known | Event handler |
| 0x006A9B88 | `HandleSelectCity` | Known | Event handler |
| 0x006A9BF4 | `HandleHighlightCity` | Known | Event handler |
| 0x006A9C2D | `HandleSelectCity` | Known | Event handler |
| 0x006A9C99 | `HandleHighlightCity` | Known | Event handler |
| 0x006A9CD2 | `HandleSelectCity` | Known | Event handler |
| 0x006A9D3E | `HandleHighlightCity` | Known | Event handler |
| 0x006A9D77 | `HandleSelectCity` | Known | Event handler |
| 0x006A9DE3 | `HandleHighlightCity` | Known | Event handler |
| 0x006A9E1C | `HandleSelectCity` | Known | Event handler |
| 0x006A9E88 | `HandleHighlightCity` | Known | Event handler |
| 0x006A9EC1 | `HandleSelectCity` | Known | Event handler |
| 0x006A9F2D | `HandleHighlightCity` | Known | Event handler |
| 0x006A9F66 | `HandleSelectCity` | Known | Event handler |
| 0x006A9FD2 | `HandleHighlightCity` | Known | Event handler |
| 0x006AA00B | `HandleSelectCity` | Known | Event handler |
| 0x006AA077 | `HandleHighlightCity` | Known | Event handler |
| 0x006AA0B0 | `HandleSelectCity` | Known | Event handler |
| 0x006AA11C | `HandleHighlightCity` | Known | Event handler |
| 0x006AA155 | `HandleSelectCity` | Known | Event handler |
| 0x006AA1C1 | `HandleHighlightCity` | Known | Event handler |
| 0x006AA1FA | `HandleSelectCity` | Known | Event handler |
| 0x006AA266 | `HandleHighlightCity` | Known | Event handler |
| 0x006AA29F | `HandleSelectCity` | Known | Event handler |
| 0x006AA30B | `HandleHighlightCity` | Known | Event handler |
| 0x006AA344 | `HandleSelectCity` | Known | Event handler |
| 0x006AA3B0 | `HandleHighlightCity` | Known | Event handler |
| 0x006AA3E9 | `HandleSelectCity` | Known | Event handler |
| 0x006AA455 | `HandleHighlightCity` | Known | Event handler |
| 0x006AA48E | `HandleSelectCity` | Known | Event handler |
| 0x006AA4FA | `HandleHighlightCity` | Known | Event handler |
| 0x006AA533 | `HandleSelectCity` | Known | Event handler |
| 0x006AA59F | `HandleHighlightCity` | Known | Event handler |
| 0x006AA5D8 | `HandleSelectCity` | Known | Event handler |
| 0x006AA644 | `HandleHighlightCity` | Known | Event handler |
| 0x006AA67D | `HandleSelectCity` | Known | Event handler |
| 0x006AA6E9 | `HandleHighlightCity` | Known | Event handler |
| 0x006AA722 | `HandleSelectCity` | Known | Event handler |
| 0x006AA78E | `HandleHighlightCity` | Known | Event handler |
| 0x006AA7C7 | `HandleSelectCity` | Known | Event handler |
| 0x006AA833 | `HandleHighlightCity` | Known | Event handler |
| 0x006AA86C | `HandleSelectCity` | Known | Event handler |
| 0x006AA8D8 | `HandleHighlightCity` | Known | Event handler |
| 0x006AA911 | `HandleSelectCity` | Known | Event handler |
| 0x006AA97D | `HandleHighlightCity` | Known | Event handler |
| 0x006AA9B6 | `HandleSelectCity` | Known | Event handler |
| 0x006AAA22 | `HandleHighlightCity` | Known | Event handler |
| 0x006AAA5B | `HandleSelectCity` | Known | Event handler |
| 0x006AAAC7 | `HandleHighlightCity` | Known | Event handler |
| 0x006AAB06 | `HandleSelectCity` | Known | Event handler |
| 0x006AAB72 | `HandleHighlightCity` | Known | Event handler |
| 0x006AABAB | `HandleSelectCity` | Known | Event handler |
| 0x006AAC17 | `HandleHighlightCity` | Known | Event handler |
| 0x006AAC50 | `HandleSelectCity` | Known | Event handler |
| 0x006AACBC | `HandleHighlightCity` | Known | Event handler |
| 0x006AACF5 | `HandleSelectCity` | Known | Event handler |
| 0x006AAD61 | `HandleHighlightCity` | Known | Event handler |
| 0x006AAD9A | `HandleSelectCity` | Known | Event handler |
| 0x006AAE06 | `HandleHighlightCity` | Known | Event handler |
| 0x006AAE46 | `HandleSelectCity` | Known | Event handler |
| 0x006AAEB2 | `HandleHighlightCity` | Known | Event handler |
| 0x006AAEEB | `HandleSelectCity` | Known | Event handler |
| 0x006AAF57 | `HandleHighlightCity` | Known | Event handler |
| 0x006AAF90 | `HandleSelectCity` | Known | Event handler |
| 0x006AAFFC | `HandleHighlightCity` | Known | Event handler |
| 0x006AB035 | `HandleSelectCity` | Known | Event handler |
| 0x006AB0A1 | `HandleHighlightCity` | Known | Event handler |
| 0x006AB0DA | `HandleSelectCity` | Known | Event handler |
| 0x006AB146 | `HandleHighlightCity` | Known | Event handler |
| 0x006AB17F | `HandleSelectCity` | Known | Event handler |
| 0x006AB1EB | `HandleHighlightCity` | Known | Event handler |
| 0x006AB224 | `HandleSelectCity` | Known | Event handler |
| 0x006AB290 | `HandleHighlightCity` | Known | Event handler |
| 0x006AB2C9 | `HandleSelectCity` | Known | Event handler |
| 0x006AB335 | `HandleHighlightCity` | Known | Event handler |
| 0x006AB36E | `HandleSelectCity` | Known | Event handler |
| 0x006AB3DA | `HandleHighlightCity` | Known | Event handler |
| 0x006AB413 | `HandleSelectCity` | Known | Event handler |
| 0x006AB47F | `HandleHighlightCity` | Known | Event handler |
| 0x006AB4B8 | `HandleSelectCity` | Known | Event handler |
| 0x006AB524 | `HandleHighlightCity` | Known | Event handler |
| 0x006AB55D | `HandleSelectCity` | Known | Event handler |
| 0x006AB5C9 | `HandleHighlightCity` | Known | Event handler |
| 0x006AB602 | `HandleSelectCity` | Known | Event handler |
| 0x006AB66E | `HandleHighlightCity` | Known | Event handler |
| 0x006AB6A7 | `HandleSelectCity` | Known | Event handler |
| 0x006AB713 | `HandleHighlightCity` | Known | Event handler |
| 0x006AB74C | `HandleSelectCity` | Known | Event handler |
| 0x006AB7B8 | `HandleHighlightCity` | Known | Event handler |
| 0x006AB7F1 | `HandleSelectCity` | Known | Event handler |
| 0x006AB85D | `HandleHighlightCity` | Known | Event handler |
| 0x006AB896 | `HandleSelectCity` | Known | Event handler |
| 0x006AB902 | `HandleHighlightCity` | Known | Event handler |
| 0x006ABDFA | `HandleMusicSelected` | Known | Event handler |
| 0x006ABE3C | `HandleMusicHilited` | Known | Event handler |
| 0x006ABE74 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x006ABEBA | `HandleMusicHilited` | Known | Event handler |
| 0x006ABEF2 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x006ABF38 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x006ABF74 | `HandleArtistsSelected` | Known | Event handler |
| 0x006ABFB8 | `HandleArtistsHilited` | Known | Event handler |
| 0x006ABFF2 | `HandleAlbumsSelected` | Known | Event handler |
| 0x006AC035 | `HandleAlbumsHilited` | Known | Event handler |
| 0x006AC06E | `HandleCompilationsSelected` | Known | Event handler |
| 0x006AC0B7 | `HandleCompilationsHilited` | Known | Event handler |
| 0x006AC0F6 | `HandleSongsSelected` | Known | Event handler |
| 0x006AC138 | `HandleSongsHilited` | Known | Event handler |
| 0x006AC170 | `HandleGenresSelected` | Known | Event handler |
| 0x006AC1B3 | `HandleGenresHilited` | Known | Event handler |
| 0x006AC1EC | `HandleComposersSelected` | Known | Event handler |
| 0x006AC232 | `HandleComposersHilited` | Known | Event handler |
| 0x006AC26E | `HandleAudiobooksSelected` | Known | Event handler |
| 0x006AC2B5 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x006AC374 | `HandleMusicHilited` | Known | Event handler |
| 0x006AC3AC | `HandleVideosSelected` | Known | Event handler |
| 0x006AC3EF | `HandleVideosHilited` | Known | Event handler |
| 0x006AC428 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x006AC473 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x006AC4B4 | `HandleMoviesSelected` | Known | Event handler |
| 0x006AC4F7 | `HandleMoviesHilited` | Known | Event handler |
| 0x006AC530 | `HandleTVShowsSelected` | Known | Event handler |
| 0x006AC574 | `HandleTVShowsHilited` | Known | Event handler |
| 0x006AC5AE | `HandleMusicVideosSelected` | Known | Event handler |
| 0x006AC5F6 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x006AC634 | `HandleRentalsSelected` | Known | Event handler |
| 0x006AC678 | `HandleRentalsHilited` | Known | Event handler |
| 0x006AC6B2 | `HandlePhotosSelected` | Known | Event handler |
| 0x006AC6F5 | `HandlePhotosHilited` | Known | Event handler |
| 0x006AC72E | `HandlePhotosSelected` | Known | Event handler |
| 0x006AC771 | `HandlePhotosHilited` | Known | Event handler |
| 0x006AC7AA | `HandlePodcastsSelected` | Known | Event handler |
| 0x006AC7EF | `HandlePodcastsHilited` | Known | Event handler |
| 0x006AC8A2 | `HandleGenericHilited` | Known | Event handler |
| 0x006AC99B | `HandleGenericHilited` | Known | Event handler |
| 0x006ACE80 | `HandleLock` | Known | Event handler |
| 0x006ACFF1 | `HandleNikePlusSelected` | Known | Event handler |
| 0x006AD036 | `HandleGenericHilited` | Known | Event handler |
| 0x006AD13C | `HandleGenericHilited` | Known | Event handler |
| 0x006AD23B | `HandleGenericHilited` | Known | Event handler |
| 0x006AD328 | `HandleGenericHilited` | Known | Event handler |
| 0x006AD425 | `HandleGenericHilited` | Known | Event handler |
| 0x006AD49F | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x006AD4E8 | `HandleGenericHilited` | Known | Event handler |
| 0x006AD561 | `HandleBacklightSelected` | Known | Event handler |
| 0x006AD5A7 | `HandleGenericHilited` | Known | Event handler |
| 0x006AD622 | `HandleSleepSelected` | Known | Event handler |
| 0x006AD664 | `HandleGenericHilited` | Known | Event handler |
| 0x006AD6DB | `HandleNowPlaying` | Known | Event handler |
| 0x006AD753 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x006AD796 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x006AD7DC | `HandleMusicHilited` | Known | Event handler |
| 0x006AD814 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x006AD85A | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x006AD898 | `HandleArtistsSelected` | Known | Event handler |
| 0x006AD8DC | `HandleArtistsHilited` | Known | Event handler |
| 0x006AD916 | `HandleAlbumsSelected` | Known | Event handler |
| 0x006AD959 | `HandleAlbumsHilited` | Known | Event handler |
| 0x006AD992 | `HandleCompilationsSelected` | Known | Event handler |
| 0x006AD9DB | `HandleCompilationsHilited` | Known | Event handler |
| 0x006ADA1A | `HandleSongsSelected` | Known | Event handler |
| 0x006ADA5C | `HandleSongsHilited` | Known | Event handler |
| 0x006ADB07 | `HandleGenericHilited` | Known | Event handler |
| 0x006ADB7F | `HandleGenresSelected` | Known | Event handler |
| 0x006ADBC2 | `HandleGenresHilited` | Known | Event handler |
| 0x006ADBFB | `HandleComposersSelected` | Known | Event handler |
| 0x006ADC41 | `HandleComposersHilited` | Known | Event handler |
| 0x006ADC7D | `HandleAudiobooksSelected` | Known | Event handler |
| 0x006ADCC4 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x006ADD83 | `HandleMusicHilited` | Known | Event handler |
| 0x006ADDF9 | `HandlePlayPause` | Known | Event handler |
| 0x006ADE2E | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x006ADF18 | `HandleSelect` | Known | Event handler |
| 0x006ADF5E | `HandleMoviesSelected` | Known | Event handler |
| 0x006ADFA1 | `HandleMoviesHilited` | Known | Event handler |
| 0x006ADFDA | `HandleRentalsSelected` | Known | Event handler |
| 0x006AE01E | `HandleRentalsHilited` | Known | Event handler |
| 0x006AE058 | `HandleTVShowsSelected` | Known | Event handler |
| 0x006AE09C | `HandleTVShowsHilited` | Known | Event handler |
| 0x006AE0D6 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x006AE11E | `HandleMusicVideosHilited` | Known | Event handler |
| 0x006AE15C | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x006AE1A7 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x006AE26D | `HandleVideosHilited` | Known | Event handler |
| 0x006AE8AF | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x006AF436 | `HandleMainMenu` | Known | Event handler |
| 0x006AF46E | `HandleMusicMenu` | Known | Event handler |
| 0x006AF996 | `HandleRadioRegion` | Known | Event handler |
| 0x006AFA3A | `HandleLanguage` | Known | Event handler |
| 0x006AFB40 | `HandleNew` | Known | Event handler |
| 0x006AFBBB | `HandleClear` | Known | Event handler |
| 0x006AFBEC | `HandleSelectCurrentSession` | Known | Event handler |
| 0x006AFCA8 | `HandleSelectIndexedSession` | Known | Event handler |
| 0x006AFE11 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x006AFE64 | `HandleSelect` | Known | Event handler |
| 0x006AFF8E | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x006AFFC8 | `HandleEQSettingSelected` | Known | Event handler |
| 0x006B0000 | `HandleEQSettingSelected` | Known | Event handler |
| 0x006C2AE4 | `HandleItemSelected` | Known | Event handler |
| 0x006C2C2F | `HandleNextContact` | Known | Event handler |
| 0x006C2C5B | `HandlePreviousContact` | Known | Event handler |
| 0x006C2C91 | `HandleSelectKey` | Known | Event handler |
| 0x006C32A2 | `HandleSelect` | Known | Event handler |
| 0x006C35C9 | `HandleDateChosen` | Known | Event handler |
| 0x006C35FF | `HandleTimeChosen` | Known | Event handler |
| 0x006C3635 | `HandleFrequencyChosen` | Known | Event handler |
| 0x006C3670 | `HandleSoundChosen` | Known | Event handler |
| 0x006C36A7 | `HandleLabelChosen` | Known | Event handler |
| 0x006C36DE | `HandleDeleteChosen` | Known | Event handler |
| 0x006C371A | `HandleSelect` | Known | Event handler |
| 0x006C3752 | `HandleSelect` | Known | Event handler |
| 0x006C3A93 | `HandleLeaveAlarm` | Known | Event handler |
| 0x006C3AC0 | `HandleLeaveAlarm` | Known | Event handler |
| 0x006C3AEF | `HandleLeaveAlarm` | Known | Event handler |
| 0x006C3B1C | `HandleLeaveAlarm` | Known | Event handler |
| 0x006C3C56 | `HandleSelect` | Known | Event handler |
| 0x006C3C84 | `HandleSelect` | Known | Event handler |
| 0x006C3DE3 | `HandleNextDay` | Known | Event handler |
| 0x006C3E0B | `HandlePreviousDay` | Known | Event handler |
| 0x006C3FBA | `HandleSelect` | Known | Event handler |
| 0x006C3FE7 | `HandleNextDay` | Known | Event handler |
| 0x006C400F | `HandlePreviousDay` | Known | Event handler |
| 0x006C41B7 | `HandleNextDay` | Known | Event handler |
| 0x006C41DF | `HandlePreviousDay` | Known | Event handler |
| 0x006C42A0 | `HandleSelect` | Known | Event handler |
| 0x006C42CB | `HandleNextDay` | Known | Event handler |
| 0x006C42F3 | `HandlePreviousDay` | Known | Event handler |
| 0x006C446A | `HandleSelectLozinch` | Known | Event handler |
| 0x006C45E2 | `HandleSelectLozinch` | Known | Event handler |
| 0x006C4701 | `HandleFlowNext` | Known | Event handler |
| 0x006C472F | `HandlePlayPause` | Known | Event handler |
| 0x006C477E | `HandleFlowPrev` | Known | Event handler |
| 0x006C47A9 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x006C489D | `HandleAlbumSelected` | Known | Event handler |
| 0x006C4A38 | `HandleFlowNext` | Known | Event handler |
| 0x006C4A86 | `HandleFlowNext` | Known | Event handler |
| 0x006C4AB4 | `HandlePlayPause` | Known | Event handler |
| 0x006C4B03 | `HandleFlowPrev` | Known | Event handler |
| 0x006C4B2F | `HandleFlowPrev` | Known | Event handler |
| 0x006C4B4F | `HandleFlowWheel` | Known | Event handler |
| 0x006C4EDF | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x006C530A | `HandleArrowDown` | Known | Event handler |
| 0x006C5374 | `HandleArrowUp` | Known | Event handler |
| 0x006C5393 | `HandleWheel` | Known | Event handler |
| 0x006C541C | `HandleSelect` | Known | Event handler |
| 0x006C5499 | `HandleGameHilited` | Known | Event handler |
| 0x006C88FF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006CA53B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006CC177 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006CDDB3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006CF9EF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006D162B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006D3267 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006D4EA3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006D6ADF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006D871B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006DA357 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006DBF93 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006DDBCF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006DF80B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006E1447 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006E3083 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006E4CBF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006E68FB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006E8537 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006EA173 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006EBDAF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006ED9EB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006EF627 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006F1263 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006F2E9F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006F4ADB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006F6717 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006F8353 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006F9F8F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006FBBCB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006FD807 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006FF443 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0070107F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00702CBB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007048F7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00706533 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0070816F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00709D90 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0070A918 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0070B4A0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0070C028 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0070CBB0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0070D738 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0070E2C0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0070EE48 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0070F9D0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00710558 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007110E0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00711C68 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007127F0 | `HandlePlayPause` | Known | Event handler |
| 0x00712826 | `HandleAddToOTG` | Known | Event handler |
| 0x007129C3 | `HandlePlayPause` | Known | Event handler |
| 0x007129EA | `HandleSelect` | Known | Event handler |
| 0x00712A17 | `HandleHilite` | Known | Event handler |
| 0x00712A48 | `HandlePlayPause` | Known | Event handler |
| 0x00712ADB | `HandlePlayPause` | Known | Event handler |
| 0x00712B02 | `HandleSelect` | Known | Event handler |
| 0x00712B68 | `HandleHilite` | Known | Event handler |
| 0x00712B9A | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x00712BE4 | `HandlePlayPause` | Known | Event handler |
| 0x00712C1A | `HandleAddToOTG` | Known | Event handler |
| 0x00712CAC | `HandlePlayPause` | Known | Event handler |
| 0x00712CD3 | `HandleSelect` | Known | Event handler |
| 0x00712D3C | `HandlePlayPause` | Known | Event handler |
| 0x00712D72 | `HandleAddToOTG` | Known | Event handler |
| 0x00712E04 | `HandlePlayPause` | Known | Event handler |
| 0x00712E2B | `HandleSelect` | Known | Event handler |
| 0x00712E94 | `HandlePlayPause` | Known | Event handler |
| 0x00712F1A | `HandleSelect` | Known | Event handler |
| 0x00712F7F | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00712FC0 | `HandlePlayPause` | Known | Event handler |
| 0x00712FF6 | `HandleAddToOTG` | Known | Event handler |
| 0x00713228 | `HandlePlayPause` | Known | Event handler |
| 0x0071324F | `HandleSelect` | Known | Event handler |
| 0x0071327C | `HandleHilite` | Known | Event handler |
| 0x007132AC | `HandlePlayPause` | Known | Event handler |
| 0x007132E2 | `HandleAddToOTG` | Known | Event handler |
| 0x00713514 | `HandlePlayPause` | Known | Event handler |
| 0x0071353B | `HandleSelect` | Known | Event handler |
| 0x00713568 | `HandleHilite` | Known | Event handler |
| 0x00713598 | `HandlePlayPause` | Known | Event handler |
| 0x007135CE | `HandleAddToOTG` | Known | Event handler |
| 0x007138B9 | `HandlePlayPause` | Known | Event handler |
| 0x007138E0 | `HandleSelect` | Known | Event handler |
| 0x00713910 | `HandlePlayPause` | Known | Event handler |
| 0x00713946 | `HandleAddToOTG` | Known | Event handler |
| 0x007139D8 | `HandlePlayPause` | Known | Event handler |
| 0x007139FF | `HandleSelect` | Known | Event handler |
| 0x00713A90 | `HandlePlayPause` | Known | Event handler |
| 0x00713AC6 | `HandleAddToOTG` | Known | Event handler |
| 0x00713C7F | `HandlePlayPause` | Known | Event handler |
| 0x00713CA6 | `HandleSelect` | Known | Event handler |
| 0x00713CD8 | `HandlePlayPause` | Known | Event handler |
| 0x00713D0E | `HandleAddToOTG` | Known | Event handler |
| 0x00713D93 | `HandleSelect` | Known | Event handler |
| 0x00713E2C | `HandleHilite` | Known | Event handler |
| 0x00713E58 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00713E9C | `HandlePlayPause` | Known | Event handler |
| 0x00713ED2 | `HandleAddToOTG` | Known | Event handler |
| 0x00713F57 | `HandleSelect` | Known | Event handler |
| 0x00713FBC | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00714000 | `HandlePlayPause` | Known | Event handler |
| 0x007141A4 | `HandleSelect` | Known | Event handler |
| 0x007141D1 | `HandleHilite` | Known | Event handler |
| 0x007141FD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00714240 | `HandlePlayPause` | Known | Event handler |
| 0x007142C6 | `HandleSelect` | Known | Event handler |
| 0x00714354 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00714398 | `HandlePlayPause` | Known | Event handler |
| 0x0071441E | `HandleSelect` | Known | Event handler |
| 0x00714483 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007144C4 | `HandlePlayPause` | Known | Event handler |
| 0x0071454A | `HandleSelect` | Known | Event handler |
| 0x007145B0 | `HandleHilite` | Known | Event handler |
| 0x007145DC | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00714620 | `HandlePlayPause` | Known | Event handler |
| 0x00714656 | `HandleAddToOTG` | Known | Event handler |
| 0x00714819 | `HandlePlayPause` | Known | Event handler |
| 0x00714840 | `HandleSelect` | Known | Event handler |
| 0x00714870 | `HandlePlayPause` | Known | Event handler |
| 0x007148A6 | `HandleAddToOTG` | Known | Event handler |
| 0x00714AC7 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x00714BE0 | `HandlePlayPause` | Known | Event handler |
| 0x00714D0D | `HandleSelect` | Known | Event handler |
| 0x00714D39 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00714D7C | `HandlePlayPause` | Known | Event handler |
| 0x00714E02 | `HandleSelect` | Known | Event handler |
| 0x00714E2F | `HandleHilite` | Known | Event handler |
| 0x00714E5B | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00714E9C | `HandlePlayPause` | Known | Event handler |
| 0x00714FCF | `HandleSelect` | Known | Event handler |
| 0x00714FFB | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0071590D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007161C5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00716A7D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00717335 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00717BED | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007184A5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00718D5D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00719615 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0071965E | `HandleTVOutChanged` | Known | Event handler |
| 0x00719696 | `HandleTVSignalChanged` | Known | Event handler |
| 0x007196D1 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x00719722 | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x00719767 | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x007197B0 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x007197F2 | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x00719835 | `HandleSelect` | Known | Event handler |
| 0x00719865 | `HandleSelect` | Known | Event handler |
| 0x0071989D | `HandleMenuLongpress` | Known | Event handler |
| 0x007198CB | `HandleMenuKey` | Known | Event handler |
| 0x00719951 | `HandlePlayPause` | Known | Event handler |
| 0x007199D1 | `HandleSelect` | Known | Event handler |
| 0x0071A2DE | `HandlePlayPause` | Known | Event handler |
| 0x0071A353 | `HandleWheelProgress` | Known | Event handler |
| 0x0071A391 | `HandleMenuLongpress` | Known | Event handler |
| 0x0071A3BF | `HandleMenuKey` | Known | Event handler |
| 0x0071A445 | `HandlePlayPause` | Known | Event handler |
| 0x0071A4C5 | `HandleSelectProgress` | Known | Event handler |
| 0x0071ADDA | `HandlePlayPause` | Known | Event handler |
| 0x0071AE4F | `HandleWheelProgress` | Known | Event handler |
| 0x0071AE8D | `HandleMenuLongpress` | Known | Event handler |
| 0x0071AEBB | `HandleMenuKey` | Known | Event handler |
| 0x0071AF41 | `HandlePlayPause` | Known | Event handler |
| 0x0071AFC1 | `HandleSelectVolume` | Known | Event handler |
| 0x0071B8D4 | `HandlePlayPause` | Known | Event handler |
| 0x0071B949 | `HandleWheelVolume` | Known | Event handler |
| 0x0071B985 | `HandleMenuLongpress` | Known | Event handler |
| 0x0071B9B3 | `HandleMenuKey` | Known | Event handler |
| 0x0071BA39 | `HandlePlayPause` | Known | Event handler |
| 0x0071BAB9 | `HandleSelectRating` | Known | Event handler |
| 0x0071C3CC | `HandlePlayPause` | Known | Event handler |
| 0x0071C441 | `HandleWheelRating` | Known | Event handler |
| 0x0071C47D | `HandleMenuLongpress` | Known | Event handler |
| 0x0071C4AB | `HandleMenuKey` | Known | Event handler |
| 0x0071C523 | `HandlePlayPause` | Known | Event handler |
| 0x0071C59A | `HandleSelectScrub` | Known | Event handler |
| 0x0071CE9E | `HandlePlayPause` | Known | Event handler |
| 0x0071CF0A | `HandleWheelScrub` | Known | Event handler |
| 0x0071CF45 | `HandleMenuLongpress` | Known | Event handler |
| 0x0071CF73 | `HandleMenuKey` | Known | Event handler |
| 0x0071CFD0 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0071D008 | `HandlePlayPause` | Known | Event handler |
| 0x0071D062 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0071D097 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x0071D9B1 | `HandlePlayPause` | Known | Event handler |
| 0x0071DA26 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0071DA69 | `HandleMenuLongpress` | Known | Event handler |
| 0x0071DA97 | `HandleMenuKey` | Known | Event handler |
| 0x0071DB1D | `HandlePlayPause` | Known | Event handler |
| 0x0071DB9D | `HandleSelectExtraInfo` | Known | Event handler |
| 0x0071E4B3 | `HandlePlayPause` | Known | Event handler |
| 0x0071E551 | `HandleMenuLongpress` | Known | Event handler |
| 0x0071E57F | `HandleMenuKey` | Known | Event handler |
| 0x0071E605 | `HandlePlayPause` | Known | Event handler |
| 0x0071E685 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x0071EF9B | `HandlePlayPause` | Known | Event handler |
| 0x0071F039 | `HandleMenuLongpress` | Known | Event handler |
| 0x0071F067 | `HandleMenuKey` | Known | Event handler |
| 0x0071F0ED | `HandlePlayPause` | Known | Event handler |
| 0x0071F16D | `HandleSelectExtraInfo` | Known | Event handler |
| 0x0071FA83 | `HandlePlayPause` | Known | Event handler |
| 0x0071FB21 | `HandleMenuLongpress` | Known | Event handler |
| 0x0071FB4F | `HandleMenuKey` | Known | Event handler |
| 0x0071FBD5 | `HandlePlayPause` | Known | Event handler |
| 0x0071FC55 | `HandleSelectChapterArt` | Known | Event handler |
| 0x0072056C | `HandlePlayPause` | Known | Event handler |
| 0x007205E1 | `HandleWheelVolume` | Known | Event handler |
| 0x0072061D | `HandleMenuLongpress` | Known | Event handler |
| 0x0072064B | `HandleMenuKey` | Known | Event handler |
| 0x007206DA | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x00720771 | `HandleSelect` | Known | Event handler |
| 0x00721087 | `HandlePlayPause` | Known | Event handler |
| 0x00721105 | `HandleWheel` | Known | Event handler |
| 0x00721139 | `HandleMenuLongpress` | Known | Event handler |
| 0x00721167 | `HandleMenuKey` | Known | Event handler |
| 0x007211F6 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x0072128D | `HandleSelect` | Known | Event handler |
| 0x00721BA3 | `HandlePlayPause` | Known | Event handler |
| 0x00721C21 | `HandleWheel` | Known | Event handler |
| 0x00721C55 | `HandleMenuLongpress` | Known | Event handler |
| 0x00721C83 | `HandleMenuKey` | Known | Event handler |
| 0x00721D09 | `HandlePlayPause` | Known | Event handler |
| 0x00721D89 | `HandleSelect` | Known | Event handler |
| 0x00722696 | `HandlePlayPause` | Known | Event handler |
| 0x0072270B | `HandleWheel` | Known | Event handler |
| 0x00722741 | `HandleMenuLongpress` | Known | Event handler |
| 0x0072276F | `HandleMenuKey` | Known | Event handler |
| 0x007227F5 | `HandlePlayPause` | Known | Event handler |
| 0x00722875 | `HandleSelectProgress` | Known | Event handler |
| 0x0072318A | `HandlePlayPause` | Known | Event handler |
| 0x007231FF | `HandleWheelProgress` | Known | Event handler |
| 0x0072323D | `HandleMenuLongpress` | Known | Event handler |
| 0x0072326B | `HandleMenuKey` | Known | Event handler |
| 0x007232E3 | `HandlePlayPause` | Known | Event handler |
| 0x0072335A | `HandleSelectScrub` | Known | Event handler |
| 0x00723C5E | `HandlePlayPause` | Known | Event handler |
| 0x00723CCA | `HandleWheelScrub` | Known | Event handler |
| 0x00723D05 | `HandleMenuLongpress` | Known | Event handler |
| 0x00723D33 | `HandleMenuKey` | Known | Event handler |
| 0x00723DB9 | `HandlePlayPause` | Known | Event handler |
| 0x00724745 | `HandlePlayPause` | Known | Event handler |
| 0x007247BA | `HandleWheelVolume` | Known | Event handler |
| 0x007247F5 | `HandleMenuLongpress` | Known | Event handler |
| 0x00724823 | `HandleMenuKey` | Known | Event handler |
| 0x007248A9 | `HandlePlayPause` | Known | Event handler |
| 0x00725235 | `HandlePlayPause` | Known | Event handler |
| 0x007252AA | `HandleWheelBrightness` | Known | Event handler |
| 0x007253C1 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x00725D14 | `HandleWheel` | Known | Event handler |
| 0x00725D49 | `HandleMenuLongpress` | Known | Event handler |
| 0x00725D77 | `HandleMenuKey` | Known | Event handler |
| 0x00725DFD | `HandlePlayPause` | Known | Event handler |
| 0x00725E7D | `HandleSelect` | Known | Event handler |
| 0x0072631F | `HandlePlayPause` | Known | Event handler |
| 0x007263AD | `HandleMenuLongpress` | Known | Event handler |
| 0x007263DB | `HandleMenuKey` | Known | Event handler |
| 0x00726461 | `HandlePlayPause` | Known | Event handler |
| 0x007264E1 | `HandleSelectProgress` | Known | Event handler |
| 0x0072698B | `HandlePlayPause` | Known | Event handler |
| 0x00726A00 | `HandleWheelProgress` | Known | Event handler |
| 0x00726A3D | `HandleMenuLongpress` | Known | Event handler |
| 0x00726A6B | `HandleMenuKey` | Known | Event handler |
| 0x00726AF1 | `HandlePlayPause` | Known | Event handler |
| 0x00726B71 | `HandleSelectProgress` | Known | Event handler |
| 0x0072701B | `HandlePlayPause` | Known | Event handler |
| 0x00727090 | `HandleWheelProgress` | Known | Event handler |
| 0x007270CD | `HandleMenuLongpress` | Known | Event handler |
| 0x007270FB | `HandleMenuKey` | Known | Event handler |
| 0x00727181 | `HandlePlayPause` | Known | Event handler |
| 0x00727201 | `HandleSelectProgress` | Known | Event handler |
| 0x00727637 | `HandlePlayPause` | Known | Event handler |
| 0x007276AC | `HandleWheelProgress` | Known | Event handler |
| 0x007276E9 | `HandleMenuLongpress` | Known | Event handler |
| 0x00727717 | `HandleMenuKey` | Known | Event handler |
| 0x00727784 | `HandlePlayPause` | Known | Event handler |
| 0x007277F0 | `HandleSelectScrub` | Known | Event handler |
| 0x00727C0A | `HandlePlayPause` | Known | Event handler |
| 0x00727C6B | `HandleWheelScrub` | Known | Event handler |
| 0x00727CA5 | `HandleMenuLongpress` | Known | Event handler |
| 0x00727CD3 | `HandleMenuKey` | Known | Event handler |
| 0x00727D59 | `HandlePlayPause` | Known | Event handler |
| 0x00727DD9 | `HandleSelectVolume` | Known | Event handler |
| 0x0072820D | `HandlePlayPause` | Known | Event handler |
| 0x00728282 | `HandleWheelVolume` | Known | Event handler |
| 0x00728395 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x00728834 | `HandleSelect` | Known | Event handler |
| 0x00728861 | `HandleSelect` | Known | Event handler |
| 0x00728891 | `HandleSelect` | Known | Event handler |
| 0x007288C1 | `HandleSelect` | Known | Event handler |
| 0x007288F1 | `HandleSelect` | Known | Event handler |
| 0x00728921 | `HandleSelect` | Known | Event handler |
| 0x00728951 | `HandleSelect` | Known | Event handler |
| 0x00728981 | `HandleSelect` | Known | Event handler |
| 0x007289B1 | `HandleSelect` | Known | Event handler |
| 0x00728A21 | `HandleSelect` | Known | Event handler |
| 0x00728A51 | `HandleSelect` | Known | Event handler |
| 0x00728AC9 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00728AFC | `HandleNotesPop` | Known | Event handler |
| 0x00728B79 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00728BAC | `HandleNotesPop` | Known | Event handler |
| 0x00729068 | `HandleNotesSelected` | Known | Event handler |
| 0x007290A5 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007290D8 | `HandleNotesPop` | Known | Event handler |
| 0x00729594 | `HandleNotesSelected` | Known | Event handler |
| 0x007295D1 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00729604 | `HandleNotesPop` | Known | Event handler |
| 0x0072962F | `HandleNotesSelected` | Known | Event handler |
| 0x00729B01 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00729B34 | `HandleNotesPop` | Known | Event handler |
| 0x00729B5F | `HandleNotesSelected` | Known | Event handler |
| 0x0072A031 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0072A064 | `HandleNotesPop` | Known | Event handler |
| 0x0072A0E1 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0072A114 | `HandleNotesPop` | Known | Event handler |
| 0x0072A191 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0072A1C4 | `HandleNotesPop` | Known | Event handler |
| 0x0072A23C | `HandlePlayPause` | Known | Event handler |
| 0x0072A265 | `HandlePlayPause` | Known | Event handler |
| 0x0072A293 | `HandlePlayPause` | Known | Event handler |
| 0x0072A2C8 | `HandleBrowseAlbum` | Known | Event handler |
| 0x0072A348 | `HandleHiliteAlbum` | Known | Event handler |
| 0x0072A3F1 | `HandleBrowseAlbum` | Known | Event handler |
| 0x0072A478 | `HandleHiliteAlbum` | Known | Event handler |
| 0x0072A73C | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x0072A798 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x0072A94F | `HandleSelect` | Known | Event handler |
| 0x0072AAD3 | `HandleSelect` | Known | Event handler |
| 0x0072AB0D | `HandleImageLast` | Known | Event handler |
| 0x0072AB37 | `HandleImageNext` | Known | Event handler |
| 0x0072AB66 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x0072ABA0 | `HandleImageFirst` | Known | Event handler |
| 0x0072ABCB | `HandleImagePrev` | Known | Event handler |
| 0x0072ABF7 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x0072AC26 | `HandleImageNext` | Known | Event handler |
| 0x0072AC4F | `HandleBrowseSlideshow` | Known | Event handler |
| 0x0072AC83 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x0072ACB2 | `HandleImagePrev` | Known | Event handler |
| 0x0072ACD3 | `HandleImageWheel` | Known | Event handler |
| 0x0072AD71 | `HandleImageNext` | Known | Event handler |
| 0x0072ADA0 | `HandlePlayPause` | Known | Event handler |
| 0x0072ADEF | `HandleImagePrev` | Known | Event handler |
| 0x0072AE1B | `HandleSelect` | Known | Event handler |
| 0x0072B0EB | `HandleImageNext` | Known | Event handler |
| 0x0072B115 | `HandlePause` | Known | Event handler |
| 0x0072B13A | `HandlePlay` | Known | Event handler |
| 0x0072B163 | `HandlePlayPause` | Known | Event handler |
| 0x0072B18C | `HandleImagePrev` | Known | Event handler |
| 0x0072B1E5 | `HandleWheel` | Known | Event handler |
| 0x0072B27D | `HandleImageNext` | Known | Event handler |
| 0x0072B2AC | `HandlePlayPause` | Known | Event handler |
| 0x0072B2FB | `HandleImagePrev` | Known | Event handler |
| 0x0072B327 | `HandleSelect` | Known | Event handler |
| 0x0072B5F7 | `HandleImageNext` | Known | Event handler |
| 0x0072B621 | `HandlePause` | Known | Event handler |
| 0x0072B646 | `HandlePlay` | Known | Event handler |
| 0x0072B66F | `HandlePlayPause` | Known | Event handler |
| 0x0072B698 | `HandleImagePrev` | Known | Event handler |
| 0x0072B6F1 | `HandleWheel` | Known | Event handler |
| 0x0072B789 | `HandleImageNext` | Known | Event handler |
| 0x0072B7B8 | `HandlePlayPause` | Known | Event handler |
| 0x0072B807 | `HandleImagePrev` | Known | Event handler |
| 0x0072B833 | `HandleSelect` | Known | Event handler |
| 0x0072BB03 | `HandleImageNext` | Known | Event handler |
| 0x0072BB2D | `HandlePause` | Known | Event handler |
| 0x0072BB52 | `HandlePlay` | Known | Event handler |
| 0x0072BB7B | `HandlePlayPause` | Known | Event handler |
| 0x0072BBA4 | `HandleImagePrev` | Known | Event handler |
| 0x0072BBFD | `HandleWheel` | Known | Event handler |
| 0x0072BC95 | `HandleImageNext` | Known | Event handler |
| 0x0072BCC4 | `HandlePlayPause` | Known | Event handler |
| 0x0072BD13 | `HandleImagePrev` | Known | Event handler |
| 0x0072BD3F | `HandleSelect` | Known | Event handler |
| 0x0072C00F | `HandleImageNext` | Known | Event handler |
| 0x0072C039 | `HandlePause` | Known | Event handler |
| 0x0072C05E | `HandlePlay` | Known | Event handler |
| 0x0072C087 | `HandlePlayPause` | Known | Event handler |
| 0x0072C0B0 | `HandleImagePrev` | Known | Event handler |
| 0x0072C109 | `HandleWheel` | Known | Event handler |
| 0x0072C1A1 | `HandleImageNext` | Known | Event handler |
| 0x0072C1D0 | `HandlePlayPause` | Known | Event handler |
| 0x0072C21F | `HandleImagePrev` | Known | Event handler |
| 0x0072C24B | `HandleSelect` | Known | Event handler |
| 0x0072C51B | `HandleImageNext` | Known | Event handler |
| 0x0072C545 | `HandlePause` | Known | Event handler |
| 0x0072C56A | `HandlePlay` | Known | Event handler |
| 0x0072C593 | `HandlePlayPause` | Known | Event handler |
| 0x0072C5BC | `HandleImagePrev` | Known | Event handler |
| 0x0072C615 | `HandleWheel` | Known | Event handler |
| 0x0072C6AD | `HandleImageNext` | Known | Event handler |
| 0x0072C6DC | `HandlePlayPause` | Known | Event handler |
| 0x0072C72B | `HandleImagePrev` | Known | Event handler |
| 0x0072C757 | `HandleSelect` | Known | Event handler |
| 0x0072CA27 | `HandleImageNext` | Known | Event handler |
| 0x0072CA51 | `HandlePause` | Known | Event handler |
| 0x0072CA76 | `HandlePlay` | Known | Event handler |
| 0x0072CA9F | `HandlePlayPause` | Known | Event handler |
| 0x0072CAC8 | `HandleImagePrev` | Known | Event handler |
| 0x0072CB21 | `HandleWheel` | Known | Event handler |
| 0x0072CBB9 | `HandleImageNext` | Known | Event handler |
| 0x0072CBE8 | `HandlePlayPause` | Known | Event handler |
| 0x0072CC37 | `HandleImagePrev` | Known | Event handler |
| 0x0072CC63 | `HandleSelect` | Known | Event handler |
| 0x0072CEAE | `HandleImageNext` | Known | Event handler |
| 0x0072CED8 | `HandlePause` | Known | Event handler |
| 0x0072CEFD | `HandlePlay` | Known | Event handler |
| 0x0072CF26 | `HandlePlayPause` | Known | Event handler |
| 0x0072CF4F | `HandleImagePrev` | Known | Event handler |
| 0x0072CFB8 | `HandleWheel` | Known | Event handler |
| 0x0072D051 | `HandleImageNext` | Known | Event handler |
| 0x0072D080 | `HandlePlayPause` | Known | Event handler |
| 0x0072D0CF | `HandleImagePrev` | Known | Event handler |
| 0x0072D0FB | `HandleSelect` | Known | Event handler |
| 0x0072D346 | `HandleImageNext` | Known | Event handler |
| 0x0072D370 | `HandlePause` | Known | Event handler |
| 0x0072D395 | `HandlePlay` | Known | Event handler |
| 0x0072D3BE | `HandlePlayPause` | Known | Event handler |
| 0x0072D3E7 | `HandleImagePrev` | Known | Event handler |
| 0x0072D450 | `HandleWheel` | Known | Event handler |
| 0x0072D4E9 | `HandleImageNext` | Known | Event handler |
| 0x0072D518 | `HandlePlayPause` | Known | Event handler |
| 0x0072D567 | `HandleImagePrev` | Known | Event handler |
| 0x0072D593 | `HandleSelect` | Known | Event handler |
| 0x0072D7DE | `HandleImageNext` | Known | Event handler |
| 0x0072D808 | `HandlePause` | Known | Event handler |
| 0x0072D82D | `HandlePlay` | Known | Event handler |
| 0x0072D856 | `HandlePlayPause` | Known | Event handler |
| 0x0072D87F | `HandleImagePrev` | Known | Event handler |
| 0x0072D8E8 | `HandleWheel` | Known | Event handler |
| 0x0072D981 | `HandleImageNext` | Known | Event handler |
| 0x0072D9B0 | `HandlePlayPause` | Known | Event handler |
| 0x0072D9FF | `HandleImagePrev` | Known | Event handler |
| 0x0072DA2B | `HandleSelect` | Known | Event handler |
| 0x0072DC76 | `HandleImageNext` | Known | Event handler |
| 0x0072DCA0 | `HandlePause` | Known | Event handler |
| 0x0072DCC5 | `HandlePlay` | Known | Event handler |
| 0x0072DCEE | `HandlePlayPause` | Known | Event handler |
| 0x0072DD17 | `HandleImagePrev` | Known | Event handler |
| 0x0072DD80 | `HandleWheel` | Known | Event handler |
| 0x0072DE19 | `HandleImageNext` | Known | Event handler |
| 0x0072DE48 | `HandlePlayPause` | Known | Event handler |
| 0x0072DE97 | `HandleImagePrev` | Known | Event handler |
| 0x0072DEC3 | `HandleSelect` | Known | Event handler |
| 0x0072E10E | `HandleImageNext` | Known | Event handler |
| 0x0072E138 | `HandlePause` | Known | Event handler |
| 0x0072E15D | `HandlePlay` | Known | Event handler |
| 0x0072E186 | `HandlePlayPause` | Known | Event handler |
| 0x0072E1AF | `HandleImagePrev` | Known | Event handler |
| 0x0072E218 | `HandleWheel` | Known | Event handler |
| 0x0072E245 | `HandleSelect` | Known | Event handler |
| 0x0072E275 | `HandleSelect` | Known | Event handler |
| 0x0072E398 | `HandleTuning` | Known | Event handler |
| 0x0072E554 | `HandleVolumeChange` | Known | Event handler |
| 0x0072E6A0 | `HandleVolumeWheel` | Known | Event handler |
| 0x0072E7FB | `HandleTuningSelect` | Known | Event handler |
| 0x0072EADA | `HandleFrequencyChange` | Known | Event handler |
| 0x0072EC37 | `HandleTuningSelect` | Known | Event handler |
| 0x0072EF16 | `HandleFrequencyChange` | Known | Event handler |
| 0x0072F040 | `HandleTimerDone` | Known | Event handler |
| 0x0072F235 | `HandleVolumeChange` | Known | Event handler |
| 0x0072F34C | `HandleVolumeWheel` | Known | Event handler |
| 0x0072F92F | `HandleExitUnsupported` | Known | Event handler |
| 0x0072F961 | `HandleExitUnsupported` | Known | Event handler |
| 0x00734995 | `HandleSelectKey` | Known | Event handler |
| 0x007349CA | `HandleWheel` | Known | Event handler |
| 0x00734B18 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x00734B6B | `HandleSelectKey` | Known | Event handler |
| 0x00734B93 | `HandleSelectKey` | Known | Event handler |
| 0x00734BC3 | `HandleExit` | Known | Event handler |
| 0x00734BED | `HandleStartStop` | Known | Event handler |
| 0x00734C53 | `HandleStartStop` | Known | Event handler |
| 0x00734D6B | `HandleExit` | Known | Event handler |
| 0x00734D95 | `HandleStartStop` | Known | Event handler |
| 0x00734DC1 | `HandleLap` | Known | Event handler |
| 0x00734EC5 | `HandleSelectLozinch` | Known | Event handler |
| 0x007350E2 | `HandleSelect` | Known | Event handler |
| 0x0073516E | `HandleSelect` | Known | Event handler |
| 0x007351FC | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x007354E6 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x007355C7 | `HandlePlayPause` | Known | Event handler |
| 0x00735655 | `HandlePlayPause` | Known | Event handler |
| 0x007356E5 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x0073571D | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x00735759 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x0073579C | `HandlePlayPause` | Known | Event handler |
| 0x007357D2 | `HandleAddToOTG` | Known | Event handler |
| 0x00735A27 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x00735C83 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x00752116 | `HandleSelectClock` | Known | Event handler |
| 0x0075214F | `HandleHilited` | Known | Event handler |
| 0x00752181 | `HandleWheel` | Known | Event handler |
| 0x007521C8 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x0075224D | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00752451 | `HandleImageLast` | Known | Event handler |
| 0x0075247B | `HandleScreenNext` | Known | Event handler |
| 0x007524AB | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007524E5 | `HandleImageFirst` | Known | Event handler |
| 0x00752510 | `HandleScreenPrev` | Known | Event handler |
| 0x0075253D | `HandleBrowseLarge` | Known | Event handler |
| 0x007525BD | `HandleImageNext` | Known | Event handler |
| 0x007525E6 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x0075261A | `HandleBrowseSlideshow` | Known | Event handler |
| 0x00752649 | `HandleImagePrev` | Known | Event handler |
| 0x00752677 | `HandleBrowseSmall` | Known | Event handler |

---

## 6. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F0FF0 | `GotoNowPlaying` | Known | Navigation |
| 0x000F1068 | `GotoMainMenu` | Known | Navigation |
| 0x00109540 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x00109558 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x001096D0 | `GotoScreen_AddressBook` | Known | Navigation |
| 0x00114AE4 | `GotoNowPlaying` | Known | Navigation |
| 0x00114AF8 | `GotoAlbums` | Known | Navigation |
| 0x00114B04 | `GotoSongs` | Known | Navigation |
| 0x00122440 | `GotoScreen_EnterPassKey` | Known | Navigation |
| 0x00122458 | `GotoScreen_LockediPod` | Known | Navigation |
| 0x00122E5C | `GotoScreen_MainMenu` | Known | Navigation |
| 0x001390B8 | `GotoMainMenu` | Known | Navigation |
| 0x001B68A4 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001C1600 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001C1E50 | `GotoPlayDeleteMenu` | Known | Navigation |
| 0x001C1ED4 | `GotoNowPlaying` | Known | Navigation |
| 0x001DAF80 | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation |
| 0x001E6654 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x001E674C | `GotoScreen_SettingsMenu` | Known | Navigation |
| 0x001EDF70 | `GotoDefaultLayout` | Known | Navigation |
| 0x001EDFF4 | `GotoVolumeLayout` | Known | Navigation |
| 0x001EE12C | `GotoProgressLayout` | Known | Navigation |
| 0x001EE448 | `GotoDefault` | Known | Navigation |
| 0x001EE77C | `GotoProgressLayout` | Known | Navigation |
| 0x001EE93C | `GotoRentalWarningLayout` | Known | Navigation |
| 0x001EE9C0 | `GotoProgressLayout` | Known | Navigation |
| 0x001EECD0 | `GotoProgressLayout` | Known | Navigation |
| 0x001F085C | `GotoNowPlaying` | Known | Navigation |
| 0x001F1128 | `GotoNowPlaying` | Known | Navigation |
| 0x001F37C8 | `GotoScreen_Language` | Known | Navigation |
| 0x001F3B28 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x001F3B44 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x001F3B5C | `GotoDefaultLayout` | Known | Navigation |
| 0x001F3B70 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x001F3C08 | `GotoVolumeLayout` | Known | Navigation |
| 0x001F3C1C | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x001F3CBC | `GotoProgressLayout` | Known | Navigation |
| 0x001F3CD0 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x001F4198 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x001F4450 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x001F45EC | `GotoProgressLayout` | Known | Navigation |
| 0x001F4600 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x001F46C4 | `GotoBrightnessVideoLayout` | Known | Navigation |
| 0x001F46E0 | `GotoRatingLayout` | Known | Navigation |
| 0x001F4984 | `GotoChapterArtLayout` | Known | Navigation |
| 0x001F499C | `GotoShuffleLayout` | Known | Navigation |
| 0x001F4CF4 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x001F4D08 | `GotoExtraInfoLoadingLayout` | Known | Navigation |
| 0x001F4DD8 | `GotoVolumeLayout` | Known | Navigation |
| 0x001F4DF0 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x001F4E7C | `GotoVolumeLayout` | Known | Navigation |
| 0x001F4E90 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x001F50A0 | `GotoScrubLayout` | Known | Navigation |
| 0x001F50B0 | `GotoScrubVideoLayout` | Known | Navigation |
| 0x001F5140 | `GotoProgressLayout` | Known | Navigation |
| 0x001F5154 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x001F52F4 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x001F5310 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x001F5328 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x001F5344 | `GotoDefaultLayout` | Known | Navigation |
| 0x001F5588 | `GotoChapterArtLayout` | Known | Navigation |
| 0x001F5680 | `GotoProgressLayout` | Known | Navigation |
| 0x001F570C | `GotoProgressLayout` | Known | Navigation |
| 0x001F5720 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x001F57FC | `GotoExtraInfoLoadFailedLayout` | Known | Navigation |
| 0x001F581C | `GotoExtraInfoLayout` | Known | Navigation |
| 0x001F5AB4 | `GotoStatusBarLayout` | Known | Navigation |
| 0x001F5AC8 | `GotoDefaultLayout` | Known | Navigation |
| 0x001F5CA0 | `GotoDefault` | Known | Navigation |
| 0x001F5DD4 | `GotoProgressLayout` | Known | Navigation |
| 0x001F5F94 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x001F60E4 | `GotoBrightnessLayout` | Known | Navigation |
| 0x001F6168 | `GotoBrightnessLayout` | Known | Navigation |
| 0x001F61E8 | `GotoVolumeLayout` | Known | Navigation |
| 0x001F6234 | `GotoScrubLayout` | Known | Navigation |
| 0x001F62FC | `GotoStatusBarLayout` | Known | Navigation |
| 0x001F6310 | `GotoDefaultLayout` | Known | Navigation |
| 0x001F63E8 | `GotoScrubLayout` | Known | Navigation |
| 0x001F6438 | `GotoScrubLayout` | Known | Navigation |
| 0x001FBDC0 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x001FBF50 | `GotoFourCard_About` | Known | Navigation |
| 0x001FBF64 | `GotoThreeCard_About` | Known | Navigation |
| 0x001FC050 | `GotoScreen_BacklightTimer` | Known | Navigation |
| 0x001FC0E0 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x001FC0F8 | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation |
| 0x00200714 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x0020072C | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x00202C80 | `GotoNowPlaying` | Known | Navigation |
| 0x00203390 | `GotoNowPlaying` | Known | Navigation |
| 0x00203A10 | `GotoFirstBoot` | Known | Navigation |
| 0x00203A20 | `GotoNotesApp` | Known | Navigation |
| 0x00203A34 | `GotoLockApp` | Known | Navigation |
| 0x00208D14 | `GotoNowPlaying` | Known | Navigation |
| 0x0038A6A4 | `GotoProgressLayout` | Known | Navigation |
| 0x006AE7E3 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x00723E39 | `GotoDefault` | Known | Navigation |
| 0x00724929 | `GotoDefault` | Known | Navigation |
| 0x0080F208 | `Gotowe?` | Known | Navigation |

---

## 7. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0014F6B8 | `CoverFlow_Screen` | Known | Screen layout |
| 0x00179B64 | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x00179B84 | `VoiceMemos_Menu_Screen_Recording` | Known | Screen layout |
| 0x00179BA8 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x006A27EE | `Clock_Screen` | Known | Screen layout |
| 0x006A27FE | `Clock_Screen_Default"` | Known | Screen layout |
| 0x006A2863 | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x006A28C1 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x006A28D9 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x006A2946 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x006A29E4 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x006A2A43 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x006A2A59 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x006A2AC4 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x006A2B1E | `Games_Menu_Screen` | Known | Screen layout |
| 0x006A2B33 | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x006A2B9D | `Extras_Screen_Games` | Known | Screen layout |
| 0x006A2C5C | `Extras_Screen_Notes` | Known | Screen layout |
| 0x006A2D20 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x006A2DE9 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x006A2E46 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x006A2E5F | `Debug_MainMenu_Screen_Default"` | Known | Screen layout |
| 0x006A2ECD | `Extras_Screen_Debug` | Known | Screen layout |
| 0x006A3004 | `Clock_Africa_City_Screen ` | Known | Screen layout |
| 0x006A3020 | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x006A30A4 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x006A30BE | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x006A3140 | `Clock_Atlantic_City_Screen"` | Known | Screen layout |
| 0x006A315E | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x006A31E4 | `Clock_Australia_City_Screen#` | Known | Screen layout |
| 0x006A3203 | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x006A328A | `Clock_Europe_City_Screen ` | Known | Screen layout |
| 0x006A32A6 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x006A332A | `Clock_NorthAmerica_City_Screen&` | Known | Screen layout |
| 0x006A334C | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x006A33D6 | `Clock_Pacific_City_Screen!` | Known | Screen layout |
| 0x006A33F3 | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x006A3478 | `Clock_SouthAmerica_City_Screen&` | Known | Screen layout |
| 0x006A349A | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x006A3527 | `Clock_Screen"` | Known | Screen layout |
| 0x006A35CC | `Clock_Screen"` | Known | Screen layout |
| 0x006A3671 | `Clock_Screen"` | Known | Screen layout |
| 0x006A3716 | `Clock_Screen"` | Known | Screen layout |
| 0x006A37BB | `Clock_Screen"` | Known | Screen layout |
| 0x006A3860 | `Clock_Screen"` | Known | Screen layout |
| 0x006A3905 | `Clock_Screen"` | Known | Screen layout |
| 0x006A39AA | `Clock_Screen"` | Known | Screen layout |
| 0x006A3A4F | `Clock_Screen"` | Known | Screen layout |
| 0x006A3AF4 | `Clock_Screen"` | Known | Screen layout |
| 0x006A3B99 | `Clock_Screen"` | Known | Screen layout |
| 0x006A3C3E | `Clock_Screen"` | Known | Screen layout |
| 0x006A3CE3 | `Clock_Screen"` | Known | Screen layout |
| 0x006A3D88 | `Clock_Screen"` | Known | Screen layout |
| 0x006A3E2D | `Clock_Screen"` | Known | Screen layout |
| 0x006A3ED2 | `Clock_Screen"` | Known | Screen layout |
| 0x006A3F77 | `Clock_Screen"` | Known | Screen layout |
| 0x006A401C | `Clock_Screen"` | Known | Screen layout |
| 0x006A40C1 | `Clock_Screen"` | Known | Screen layout |
| 0x006A4166 | `Clock_Screen"` | Known | Screen layout |
| 0x006A420B | `Clock_Screen"` | Known | Screen layout |
| 0x006A42B0 | `Clock_Screen"` | Known | Screen layout |
| 0x006A4355 | `Clock_Screen"` | Known | Screen layout |
| 0x006A43FA | `Clock_Screen"` | Known | Screen layout |
| 0x006A449F | `Clock_Screen"` | Known | Screen layout |
| 0x006A4544 | `Clock_Screen"` | Known | Screen layout |
| 0x006A45E9 | `Clock_Screen"` | Known | Screen layout |
| 0x006A468E | `Clock_Screen"` | Known | Screen layout |
| 0x006A4733 | `Clock_Screen"` | Known | Screen layout |
| 0x006A47D8 | `Clock_Screen"` | Known | Screen layout |
| 0x006A487D | `Clock_Screen"` | Known | Screen layout |
| 0x006A4927 | `Clock_Screen"` | Known | Screen layout |
| 0x006A49CC | `Clock_Screen"` | Known | Screen layout |
| 0x006A4A71 | `Clock_Screen"` | Known | Screen layout |
| 0x006A4B16 | `Clock_Screen"` | Known | Screen layout |
| 0x006A4BBB | `Clock_Screen"` | Known | Screen layout |
| 0x006A4C60 | `Clock_Screen"` | Known | Screen layout |
| 0x006A4D05 | `Clock_Screen"` | Known | Screen layout |
| 0x006A4DAA | `Clock_Screen"` | Known | Screen layout |
| 0x006A4E4F | `Clock_Screen"` | Known | Screen layout |
| 0x006A4EF4 | `Clock_Screen"` | Known | Screen layout |
| 0x006A4F99 | `Clock_Screen"` | Known | Screen layout |
| 0x006A503E | `Clock_Screen"` | Known | Screen layout |
| 0x006A50E3 | `Clock_Screen"` | Known | Screen layout |
| 0x006A5188 | `Clock_Screen"` | Known | Screen layout |
| 0x006A522D | `Clock_Screen"` | Known | Screen layout |
| 0x006A52D2 | `Clock_Screen"` | Known | Screen layout |
| 0x006A5377 | `Clock_Screen"` | Known | Screen layout |
| 0x006A541C | `Clock_Screen"` | Known | Screen layout |
| 0x006A54C1 | `Clock_Screen"` | Known | Screen layout |
| 0x006A5566 | `Clock_Screen"` | Known | Screen layout |
| 0x006A560B | `Clock_Screen"` | Known | Screen layout |
| 0x006A56B0 | `Clock_Screen"` | Known | Screen layout |
| 0x006A5755 | `Clock_Screen"` | Known | Screen layout |
| 0x006A57FA | `Clock_Screen"` | Known | Screen layout |
| 0x006A589F | `Clock_Screen"` | Known | Screen layout |
| 0x006A5944 | `Clock_Screen"` | Known | Screen layout |
| 0x006A59E9 | `Clock_Screen"` | Known | Screen layout |
| 0x006A5A8E | `Clock_Screen"` | Known | Screen layout |
| 0x006A5B33 | `Clock_Screen"` | Known | Screen layout |
| 0x006A5BD8 | `Clock_Screen"` | Known | Screen layout |
| 0x006A5C7D | `Clock_Screen"` | Known | Screen layout |
| 0x006A5D22 | `Clock_Screen"` | Known | Screen layout |
| 0x006A5DC7 | `Clock_Screen"` | Known | Screen layout |
| 0x006A5E6C | `Clock_Screen"` | Known | Screen layout |
| 0x006A5F11 | `Clock_Screen"` | Known | Screen layout |
| 0x006A5FB6 | `Clock_Screen"` | Known | Screen layout |
| 0x006A605B | `Clock_Screen"` | Known | Screen layout |
| 0x006A6100 | `Clock_Screen"` | Known | Screen layout |
| 0x006A61A5 | `Clock_Screen"` | Known | Screen layout |
| 0x006A624A | `Clock_Screen"` | Known | Screen layout |
| 0x006A62EF | `Clock_Screen"` | Known | Screen layout |
| 0x006A6394 | `Clock_Screen"` | Known | Screen layout |
| 0x006A6439 | `Clock_Screen"` | Known | Screen layout |
| 0x006A64DE | `Clock_Screen"` | Known | Screen layout |
| 0x006A6583 | `Clock_Screen"` | Known | Screen layout |
| 0x006A6628 | `Clock_Screen"` | Known | Screen layout |
| 0x006A66CD | `Clock_Screen"` | Known | Screen layout |
| 0x006A6772 | `Clock_Screen"` | Known | Screen layout |
| 0x006A6817 | `Clock_Screen"` | Known | Screen layout |
| 0x006A68BC | `Clock_Screen"` | Known | Screen layout |
| 0x006A6961 | `Clock_Screen"` | Known | Screen layout |
| 0x006A6A06 | `Clock_Screen"` | Known | Screen layout |
| 0x006A6AAB | `Clock_Screen"` | Known | Screen layout |
| 0x006A6B50 | `Clock_Screen"` | Known | Screen layout |
| 0x006A6BF5 | `Clock_Screen"` | Known | Screen layout |
| 0x006A6C9A | `Clock_Screen"` | Known | Screen layout |
| 0x006A6D3F | `Clock_Screen"` | Known | Screen layout |
| 0x006A6DEB | `Clock_Screen"` | Known | Screen layout |
| 0x006A6E90 | `Clock_Screen"` | Known | Screen layout |
| 0x006A6F35 | `Clock_Screen"` | Known | Screen layout |
| 0x006A6FDF | `Clock_Screen"` | Known | Screen layout |
| 0x006A7084 | `Clock_Screen"` | Known | Screen layout |
| 0x006A7129 | `Clock_Screen"` | Known | Screen layout |
| 0x006A71CE | `Clock_Screen"` | Known | Screen layout |
| 0x006A7273 | `Clock_Screen"` | Known | Screen layout |
| 0x006A7318 | `Clock_Screen"` | Known | Screen layout |
| 0x006A73BD | `Clock_Screen"` | Known | Screen layout |
| 0x006A7462 | `Clock_Screen"` | Known | Screen layout |
| 0x006A750B | `Clock_Screen"` | Known | Screen layout |
| 0x006A75B0 | `Clock_Screen"` | Known | Screen layout |
| 0x006A7655 | `Clock_Screen"` | Known | Screen layout |
| 0x006A76FA | `Clock_Screen"` | Known | Screen layout |
| 0x006A779F | `Clock_Screen"` | Known | Screen layout |
| 0x006A7844 | `Clock_Screen"` | Known | Screen layout |
| 0x006A78E9 | `Clock_Screen"` | Known | Screen layout |
| 0x006A798E | `Clock_Screen"` | Known | Screen layout |
| 0x006A7A33 | `Clock_Screen"` | Known | Screen layout |
| 0x006A7AD8 | `Clock_Screen"` | Known | Screen layout |
| 0x006A7B7D | `Clock_Screen"` | Known | Screen layout |
| 0x006A7C22 | `Clock_Screen"` | Known | Screen layout |
| 0x006A7CC7 | `Clock_Screen"` | Known | Screen layout |
| 0x006A7D6C | `Clock_Screen"` | Known | Screen layout |
| 0x006A7E11 | `Clock_Screen"` | Known | Screen layout |
| 0x006A7EB6 | `Clock_Screen"` | Known | Screen layout |
| 0x006A7F5B | `Clock_Screen"` | Known | Screen layout |
| 0x006A8000 | `Clock_Screen"` | Known | Screen layout |
| 0x006A80A5 | `Clock_Screen"` | Known | Screen layout |
| 0x006A814A | `Clock_Screen"` | Known | Screen layout |
| 0x006A81EF | `Clock_Screen"` | Known | Screen layout |
| 0x006A8294 | `Clock_Screen"` | Known | Screen layout |
| 0x006A8339 | `Clock_Screen"` | Known | Screen layout |
| 0x006A83DE | `Clock_Screen"` | Known | Screen layout |
| 0x006A8483 | `Clock_Screen"` | Known | Screen layout |
| 0x006A8528 | `Clock_Screen"` | Known | Screen layout |
| 0x006A85CD | `Clock_Screen"` | Known | Screen layout |
| 0x006A8672 | `Clock_Screen"` | Known | Screen layout |
| 0x006A8717 | `Clock_Screen"` | Known | Screen layout |
| 0x006A87BC | `Clock_Screen"` | Known | Screen layout |
| 0x006A8861 | `Clock_Screen"` | Known | Screen layout |
| 0x006A8906 | `Clock_Screen"` | Known | Screen layout |
| 0x006A89AB | `Clock_Screen"` | Known | Screen layout |
| 0x006A8A50 | `Clock_Screen"` | Known | Screen layout |
| 0x006A8AFB | `Clock_Screen"` | Known | Screen layout |
| 0x006A8BA0 | `Clock_Screen"` | Known | Screen layout |
| 0x006A8C45 | `Clock_Screen"` | Known | Screen layout |
| 0x006A8CEA | `Clock_Screen"` | Known | Screen layout |
| 0x006A8D8F | `Clock_Screen"` | Known | Screen layout |
| 0x006A8E34 | `Clock_Screen"` | Known | Screen layout |
| 0x006A8ED9 | `Clock_Screen"` | Known | Screen layout |
| 0x006A8F7E | `Clock_Screen"` | Known | Screen layout |
| 0x006A9023 | `Clock_Screen"` | Known | Screen layout |
| 0x006A90C8 | `Clock_Screen"` | Known | Screen layout |
| 0x006A916D | `Clock_Screen"` | Known | Screen layout |
| 0x006A9212 | `Clock_Screen"` | Known | Screen layout |
| 0x006A92B7 | `Clock_Screen"` | Known | Screen layout |
| 0x006A935C | `Clock_Screen"` | Known | Screen layout |
| 0x006A9401 | `Clock_Screen"` | Known | Screen layout |
| 0x006A94A6 | `Clock_Screen"` | Known | Screen layout |
| 0x006A954B | `Clock_Screen"` | Known | Screen layout |
| 0x006A95F0 | `Clock_Screen"` | Known | Screen layout |
| 0x006A9695 | `Clock_Screen"` | Known | Screen layout |
| 0x006A973A | `Clock_Screen"` | Known | Screen layout |
| 0x006A97DF | `Clock_Screen"` | Known | Screen layout |
| 0x006A9884 | `Clock_Screen"` | Known | Screen layout |
| 0x006A9929 | `Clock_Screen"` | Known | Screen layout |
| 0x006A99CE | `Clock_Screen"` | Known | Screen layout |
| 0x006A9A73 | `Clock_Screen"` | Known | Screen layout |
| 0x006A9B18 | `Clock_Screen"` | Known | Screen layout |
| 0x006A9BBD | `Clock_Screen"` | Known | Screen layout |
| 0x006A9C62 | `Clock_Screen"` | Known | Screen layout |
| 0x006A9D07 | `Clock_Screen"` | Known | Screen layout |
| 0x006A9DAC | `Clock_Screen"` | Known | Screen layout |
| 0x006A9E51 | `Clock_Screen"` | Known | Screen layout |
| 0x006A9EF6 | `Clock_Screen"` | Known | Screen layout |
| 0x006A9F9B | `Clock_Screen"` | Known | Screen layout |
| 0x006AA040 | `Clock_Screen"` | Known | Screen layout |
| 0x006AA0E5 | `Clock_Screen"` | Known | Screen layout |
| 0x006AA18A | `Clock_Screen"` | Known | Screen layout |
| 0x006AA22F | `Clock_Screen"` | Known | Screen layout |
| 0x006AA2D4 | `Clock_Screen"` | Known | Screen layout |
| 0x006AA379 | `Clock_Screen"` | Known | Screen layout |
| 0x006AA41E | `Clock_Screen"` | Known | Screen layout |
| 0x006AA4C3 | `Clock_Screen"` | Known | Screen layout |
| 0x006AA568 | `Clock_Screen"` | Known | Screen layout |
| 0x006AA60D | `Clock_Screen"` | Known | Screen layout |
| 0x006AA6B2 | `Clock_Screen"` | Known | Screen layout |
| 0x006AA757 | `Clock_Screen"` | Known | Screen layout |
| 0x006AA7FC | `Clock_Screen"` | Known | Screen layout |
| 0x006AA8A1 | `Clock_Screen"` | Known | Screen layout |
| 0x006AA946 | `Clock_Screen"` | Known | Screen layout |
| 0x006AA9EB | `Clock_Screen"` | Known | Screen layout |
| 0x006AAA90 | `Clock_Screen"` | Known | Screen layout |
| 0x006AAB3B | `Clock_Screen"` | Known | Screen layout |
| 0x006AABE0 | `Clock_Screen"` | Known | Screen layout |
| 0x006AAC85 | `Clock_Screen"` | Known | Screen layout |
| 0x006AAD2A | `Clock_Screen"` | Known | Screen layout |
| 0x006AADCF | `Clock_Screen"` | Known | Screen layout |
| 0x006AAE7B | `Clock_Screen"` | Known | Screen layout |
| 0x006AAF20 | `Clock_Screen"` | Known | Screen layout |
| 0x006AAFC5 | `Clock_Screen"` | Known | Screen layout |
| 0x006AB06A | `Clock_Screen"` | Known | Screen layout |
| 0x006AB10F | `Clock_Screen"` | Known | Screen layout |
| 0x006AB1B4 | `Clock_Screen"` | Known | Screen layout |
| 0x006AB259 | `Clock_Screen"` | Known | Screen layout |
| 0x006AB2FE | `Clock_Screen"` | Known | Screen layout |
| 0x006AB3A3 | `Clock_Screen"` | Known | Screen layout |
| 0x006AB448 | `Clock_Screen"` | Known | Screen layout |
| 0x006AB4ED | `Clock_Screen"` | Known | Screen layout |
| 0x006AB592 | `Clock_Screen"` | Known | Screen layout |
| 0x006AB637 | `Clock_Screen"` | Known | Screen layout |
| 0x006AB6DC | `Clock_Screen"` | Known | Screen layout |
| 0x006AB781 | `Clock_Screen"` | Known | Screen layout |
| 0x006AB826 | `Clock_Screen"` | Known | Screen layout |
| 0x006AB8CB | `Clock_Screen"` | Known | Screen layout |
| 0x006AB96E | `Settings_DateTime_SetDate_Screen(` | Known | Screen layout |
| 0x006AB992 | `Settings_DateTime_SetDate_Screen_Default"` | Known | Screen layout |
| 0x006ABA0B | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x006ABA71 | `Settings_DateTime_SetTime_Screen(` | Known | Screen layout |
| 0x006ABA95 | `Settings_DateTime_SetTime_Screen_Default"` | Known | Screen layout |
| 0x006ABB0E | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x006ABB79 | `Settings_DateTime_SetTimeZone_Screen,` | Known | Screen layout |
| 0x006ABBA1 | `Settings_DateTime_SetTimeZone_Screen_Default"` | Known | Screen layout |
| 0x006ABC1E | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x006ABCD7 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x006ABD87 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x006AC316 | `Search_Main_Screen` | Known | Screen layout |
| 0x006AC32C | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x006AC84E | `Extras_Screen` | Known | Screen layout |
| 0x006AC85F | `Extras_Screen_WorldClock"` | Known | Screen layout |
| 0x006AC8DC | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x006AC93E | `Clock_Screen` | Known | Screen layout |
| 0x006AC94E | `Clock_Screen_Default` | Known | Screen layout |
| 0x006AC9D5 | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x006ACA3B | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x006ACA51 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x006ACABC | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x006ACB1E | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x006ACB36 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x006ACBA3 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x006ACC07 | `AddressViewer_Main_Screen!` | Known | Screen layout |
| 0x006ACC24 | `AddressViewer_Main_Screen_Default"` | Known | Screen layout |
| 0x006ACC96 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x006ACCFD | `Games_Menu_Screen` | Known | Screen layout |
| 0x006ACD12 | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x006ACD7C | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x006ACE43 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x006ACEDF | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x006ACFB0 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x006AD070 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x006AD0D4 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x006AD0F3 | `VoiceMemos_Menu_Screen_Default"` | Known | Screen layout |
| 0x006AD176 | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x006AD1DC | `Speakers_Main_Screen` | Known | Screen layout |
| 0x006AD1F4 | `Speakers_Main_Screen_Default"` | Known | Screen layout |
| 0x006AD275 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x006AD2D9 | `Radio_Screen` | Known | Screen layout |
| 0x006AD2E9 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x006AD362 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x006AD3C3 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x006AD45F | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x006AD522 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x006AD5E1 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x006AD69E | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x006ADAB8 | `Radio_Screen` | Known | Screen layout |
| 0x006ADAC8 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x006ADB41 | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x006ADD25 | `Search_Main_Screen` | Known | Screen layout |
| 0x006ADD3B | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x006ADE68 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x006ADECB | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x006AE20C | `Video_Settings_Screen` | Known | Screen layout |
| 0x006AE225 | `Video_Settings_Screen_Default"` | Known | Screen layout |
| 0x006AE322 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x006AE5E7 | `SettingsMenus_MainMenu_Screen%` | Known | Screen layout |
| 0x006AE6F5 | `SettingsMenus_MusicMenu_Screen&` | Known | Screen layout |
| 0x006AE99E | `SettingsMenus_Brightness_Screen&` | Known | Screen layout |
| 0x006AEAB3 | `SettingsMenus_AudiobookSettings_Screen.` | Known | Screen layout |
| 0x006AEBE9 | `SettingsMenus_RadioRegions_Screen ` | Known | Screen layout |
| 0x006AECFE | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x006AEF6A | `Settings_DateTime_Screen ` | Known | Screen layout |
| 0x006AEF86 | `Settings_DateTime_Screen_Default"` | Known | Screen layout |
| 0x006AF112 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x006AF217 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x006AF230 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x006AF321 | `SettingsMenus_ResetAllSettings_Screen4` | Known | Screen layout |
| 0x006AFAF2 | `Stopwatch_Screen` | Known | Screen layout |
| 0x006AFB06 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x006AFB6D | `Stopwatch_Screen` | Known | Screen layout |
| 0x006AFB81 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x006AFC2A | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x006AFC4D | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x006AFCE6 | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x006AFD09 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x006AFEBC | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x006AFF2A | `Speakers_ToneControl_Screen#` | Known | Screen layout |
| 0x006AFF49 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x006C2985 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C2A08 | `LockediPod_Screen` | Known | Screen layout |
| 0x006C2A90 | `Lock_Screen` | Known | Screen layout |
| 0x006C2A9F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C2B1A | `AddressViewer_ContactDetails_Screen+` | Known | Screen layout |
| 0x006C2B41 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x006C2BBC | `Extras_Screen` | Known | Screen layout |
| 0x006C2C07 | `Extras_Screen` | Known | Screen layout |
| 0x006C2CEE | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x006C2D4C | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x006C2D69 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x006C2DD7 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x006C2DF0 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x006C2E67 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x006C2E84 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x006C2EEF | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x006C2F0C | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x006C2F73 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x006C2FDA | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x006C3038 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x006C3055 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x006C30C3 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x006C30DC | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x006C3153 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x006C3170 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x006C31DB | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x006C31F8 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x006C325F | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x006C32FF | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x006C3388 | `Alarms_Set_Alarm_Frequency_Screen)` | Known | Screen layout |
| 0x006C33AD | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x006C341E | `Alarms_Set_Alarm_Label_Screen%` | Known | Screen layout |
| 0x006C343F | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x006C34AC | `Alarms_Set_Alarm_Sound_Screen%` | Known | Screen layout |
| 0x006C34CD | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x006C3539 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x006C37B4 | `Alarms_Set_Alarm_Playlist_Screen(` | Known | Screen layout |
| 0x006C37D8 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x006C3848 | `Alarms_Set_Alarm_Tones_Screen%` | Known | Screen layout |
| 0x006C3869 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x006C3B7C | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x006C3B97 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x006C3CE8 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x006C3CFF | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x006C3D80 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x006C3D97 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x006C3E6D | `Calendar_Event_Screen` | Known | Screen layout |
| 0x006C3E86 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x006C3F0B | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x006C3F7C | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x006C4071 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x006C408A | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x006C410F | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x006C4180 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x006C4240 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x006C4254 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x006C4383 | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x006C43E6 | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x006C443D | `Clock_Screen_Default` | Known | Screen layout |
| 0x006C44CE | `Clock_Region_Screen` | Known | Screen layout |
| 0x006C44E5 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x006C455E | `Clock_Screen_Default` | Known | Screen layout |
| 0x006C45B5 | `Clock_Screen_Default` | Known | Screen layout |
| 0x006C4646 | `Clock_Region_Screen` | Known | Screen layout |
| 0x006C465D | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x006C47E8 | `CoverFlow_Screen_DefaultNoAlbumInfo&` | Known | Screen layout |
| 0x006C48D6 | `CoverFlow_Screen_DefaultNoAlbumInfo#` | Known | Screen layout |
| 0x006C494B | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x006C4C41 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x006C4DF1 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x006C4F1F | `CoverFlow_Screen_Default&` | Known | Screen layout |
| 0x006C4FF5 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x006C518A | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x006C53EF | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x006C544C | `Game_Screen` | Known | Screen layout |
| 0x006C545B | `Game_Screen_Default` | Known | Screen layout |
| 0x006C54FD | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x006C555F | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x006C55C2 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x006C5625 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x006C5681 | `Game_Running_Screen` | Known | Screen layout |
| 0x006C56E1 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x006C5743 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x006C57A6 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x006C5809 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x006C5865 | `Game_Running_Screen` | Known | Screen layout |
| 0x006C58C5 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x006C5927 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x006C598A | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x006C59ED | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x006C5A49 | `Game_Running_Screen` | Known | Screen layout |
| 0x006C5AA9 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x006C5B0B | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x006C5B6E | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x006C5BD1 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x006C5C2D | `Game_Running_Screen` | Known | Screen layout |
| 0x006C5C8D | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x006C5CEF | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x006C5D52 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x006C5DB5 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x006C5E11 | `Game_Running_Screen` | Known | Screen layout |
| 0x006C6057 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x006C60B9 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x006C611C | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x006C617F | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x006C61DB | `Game_Running_Screen` | Known | Screen layout |
| 0x006C6292 | `Extras_Screen` | Known | Screen layout |
| 0x006C62A3 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x006C6301 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x006C649E | `Extras_Screen` | Known | Screen layout |
| 0x006C64AF | `Extras_Screen_Lock` | Known | Screen layout |
| 0x006C650D | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x006C66AA | `Extras_Screen` | Known | Screen layout |
| 0x006C66BB | `Extras_Screen_Lock` | Known | Screen layout |
| 0x006C6719 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x006C68B6 | `Extras_Screen` | Known | Screen layout |
| 0x006C68C7 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x006C6925 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x006C6AC7 | `Lock_Screen` | Known | Screen layout |
| 0x006C6AD6 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x006C6B38 | `Extras_Screen` | Known | Screen layout |
| 0x006C6B49 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x006C6BA8 | `LockediPod_Screen` | Known | Screen layout |
| 0x006C6C22 | `MainMenus_Main_Screen)` | Known | Screen layout |
| 0x006C6DF3 | `Lock_Screen` | Known | Screen layout |
| 0x006C6E02 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x006C6E64 | `Extras_Screen` | Known | Screen layout |
| 0x006C6E75 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x006C6ED4 | `LockediPod_Screen` | Known | Screen layout |
| 0x006C6F4E | `MainMenus_Main_Screen%` | Known | Screen layout |
| 0x006C6FB5 | `LockediPod_Screen` | Known | Screen layout |
| 0x006C6FCA | `LockediPod_Screen_Default,` | Known | Screen layout |
| 0x006C7119 | `Lock_Screen` | Known | Screen layout |
| 0x006C7128 | `Lock_Screen_Default_Layout"` | Known | Screen layout |
| 0x006C7191 | `Lock_Screen` | Known | Screen layout |
| 0x006C71A0 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x006C7202 | `Extras_Screen` | Known | Screen layout |
| 0x006C7213 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x006C7272 | `LockediPod_Screen` | Known | Screen layout |
| 0x006C72EC | `MainMenus_Main_Screen,` | Known | Screen layout |
| 0x006C7448 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x006C74AE | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x006C7512 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C75A1 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x006C760E | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x006C767B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x006C76E8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C7750 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x006C77B6 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x006C781A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C78A9 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x006C7916 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x006C7983 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x006C79F0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C7A58 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x006C7ABE | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x006C7B22 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C7BB1 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x006C7C1E | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x006C7C8B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x006C7CF8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C7D60 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x006C7DC6 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x006C7E2A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C7EB9 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x006C7F26 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x006C7F93 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x006C8000 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C8068 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x006C80CE | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x006C8132 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C81C1 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x006C822E | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x006C829B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x006C8308 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C8361 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x006C83CA | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x006C8431 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x006C84CC | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x006C8535 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x006C859E | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x006C8605 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x006C86A0 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x006C8709 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x006C8772 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x006C87D9 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x006C8874 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x006C8960 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C897C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C89EA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C8A07 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C8A72 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006C8A92 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006C8B09 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C8B25 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C8B95 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006C8BB4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006C8C20 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006C8C34 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006C8CAD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006C8D21 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006C8D91 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006C8DF9 | `NoContent_Screen` | Known | Screen layout |
| 0x006C8E0D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006C8E71 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006C8ED8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C8EF2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006C8F60 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006C8FD2 | `NoContent_Screen` | Known | Screen layout |
| 0x006C8FE6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006C9050 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006C90B9 | `No_Photos_Screen` | Known | Screen layout |
| 0x006C90CD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006C9133 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C91A1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006C920E | `NoContent_Screen` | Known | Screen layout |
| 0x006C9222 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006C928A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006C92F4 | `NoContent_Screen` | Known | Screen layout |
| 0x006C9308 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006C936F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006C93D9 | `NoContent_Screen` | Known | Screen layout |
| 0x006C93ED | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006C945A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006C94CC | `NoContent_Screen` | Known | Screen layout |
| 0x006C94E0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C9548 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006C95B1 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006C95CC | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006C9632 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006C964E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006C972D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006C9746 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006C97A7 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006C97BB | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006C9929 | `Radio_Screen` | Known | Screen layout |
| 0x006C9939 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006C999A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C9A1D | `LockediPod_Screen` | Known | Screen layout |
| 0x006C9AA5 | `Lock_Screen` | Known | Screen layout |
| 0x006C9AB4 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C9B17 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006C9B79 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006C9B95 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006C9C07 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006C9C26 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006C9C8E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C9CA8 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006C9D10 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C9D2D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C9D99 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006C9E03 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006C9E1D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006C9E8D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006C9F00 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006C9F71 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006C9FE0 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006CA04C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006CA067 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006CA0DC | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006CA143 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006CA1A5 | `Photos_Screen` | Known | Screen layout |
| 0x006CA209 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006CA227 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006CA299 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006CA2B6 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006CA31C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006CA337 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006CA3A0 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006CA3BD | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006CA434 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006CA458 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006CA4C6 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006CA4E1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006CA59C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CA5B8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CA626 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CA643 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CA6AE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006CA6CE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006CA745 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CA761 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CA7D1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006CA7F0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006CA85C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006CA870 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006CA8E9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006CA95D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006CA9CD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006CAA35 | `NoContent_Screen` | Known | Screen layout |
| 0x006CAA49 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006CAAAD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006CAB14 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CAB2E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006CAB9C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006CAC0E | `NoContent_Screen` | Known | Screen layout |
| 0x006CAC22 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006CAC8C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006CACF5 | `No_Photos_Screen` | Known | Screen layout |
| 0x006CAD09 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006CAD6F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006CADDD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006CAE4A | `NoContent_Screen` | Known | Screen layout |
| 0x006CAE5E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006CAEC6 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006CAF30 | `NoContent_Screen` | Known | Screen layout |
| 0x006CAF44 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006CAFAB | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006CB015 | `NoContent_Screen` | Known | Screen layout |
| 0x006CB029 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006CB096 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006CB108 | `NoContent_Screen` | Known | Screen layout |
| 0x006CB11C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006CB184 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006CB1ED | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006CB208 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006CB26E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006CB28A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006CB369 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006CB382 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006CB3E3 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006CB3F7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006CB565 | `Radio_Screen` | Known | Screen layout |
| 0x006CB575 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006CB5D6 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006CB659 | `LockediPod_Screen` | Known | Screen layout |
| 0x006CB6E1 | `Lock_Screen` | Known | Screen layout |
| 0x006CB6F0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006CB753 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006CB7B5 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006CB7D1 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006CB843 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006CB862 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006CB8CA | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CB8E4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006CB94C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CB969 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CB9D5 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006CBA3F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006CBA59 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006CBAC9 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006CBB3C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006CBBAD | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006CBC1C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006CBC88 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006CBCA3 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006CBD18 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006CBD7F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006CBDE1 | `Photos_Screen` | Known | Screen layout |
| 0x006CBE45 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006CBE63 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006CBED5 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006CBEF2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006CBF58 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006CBF73 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006CBFDC | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006CBFF9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006CC070 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006CC094 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006CC102 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006CC11D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006CC1D8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CC1F4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CC262 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CC27F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CC2EA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006CC30A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006CC381 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CC39D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CC40D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006CC42C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006CC498 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006CC4AC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006CC525 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006CC599 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006CC609 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006CC671 | `NoContent_Screen` | Known | Screen layout |
| 0x006CC685 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006CC6E9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006CC750 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CC76A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006CC7D8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006CC84A | `NoContent_Screen` | Known | Screen layout |
| 0x006CC85E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006CC8C8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006CC931 | `No_Photos_Screen` | Known | Screen layout |
| 0x006CC945 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006CC9AB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006CCA19 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006CCA86 | `NoContent_Screen` | Known | Screen layout |
| 0x006CCA9A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006CCB02 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006CCB6C | `NoContent_Screen` | Known | Screen layout |
| 0x006CCB80 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006CCBE7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006CCC51 | `NoContent_Screen` | Known | Screen layout |
| 0x006CCC65 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006CCCD2 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006CCD44 | `NoContent_Screen` | Known | Screen layout |
| 0x006CCD58 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006CCDC0 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006CCE29 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006CCE44 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006CCEAA | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006CCEC6 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006CCFA5 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006CCFBE | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006CD01F | `FirstBoot_Screen` | Known | Screen layout |
| 0x006CD033 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006CD1A1 | `Radio_Screen` | Known | Screen layout |
| 0x006CD1B1 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006CD212 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006CD295 | `LockediPod_Screen` | Known | Screen layout |
| 0x006CD31D | `Lock_Screen` | Known | Screen layout |
| 0x006CD32C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006CD38F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006CD3F1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006CD40D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006CD47F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006CD49E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006CD506 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CD520 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006CD588 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CD5A5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CD611 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006CD67B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006CD695 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006CD705 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006CD778 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006CD7E9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006CD858 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006CD8C4 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006CD8DF | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006CD954 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006CD9BB | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006CDA1D | `Photos_Screen` | Known | Screen layout |
| 0x006CDA81 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006CDA9F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006CDB11 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006CDB2E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006CDB94 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006CDBAF | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006CDC18 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006CDC35 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006CDCAC | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006CDCD0 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006CDD3E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006CDD59 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006CDE14 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CDE30 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CDE9E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CDEBB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CDF26 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006CDF46 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006CDFBD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CDFD9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CE049 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006CE068 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006CE0D4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006CE0E8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006CE161 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006CE1D5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006CE245 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006CE2AD | `NoContent_Screen` | Known | Screen layout |
| 0x006CE2C1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006CE325 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006CE38C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CE3A6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006CE414 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006CE486 | `NoContent_Screen` | Known | Screen layout |
| 0x006CE49A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006CE504 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006CE56D | `No_Photos_Screen` | Known | Screen layout |
| 0x006CE581 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006CE5E7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006CE655 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006CE6C2 | `NoContent_Screen` | Known | Screen layout |
| 0x006CE6D6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006CE73E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006CE7A8 | `NoContent_Screen` | Known | Screen layout |
| 0x006CE7BC | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006CE823 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006CE88D | `NoContent_Screen` | Known | Screen layout |
| 0x006CE8A1 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006CE90E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006CE980 | `NoContent_Screen` | Known | Screen layout |
| 0x006CE994 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006CE9FC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006CEA65 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006CEA80 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006CEAE6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006CEB02 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006CEBE1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006CEBFA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006CEC5B | `FirstBoot_Screen` | Known | Screen layout |
| 0x006CEC6F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006CEDDD | `Radio_Screen` | Known | Screen layout |
| 0x006CEDED | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006CEE4E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006CEED1 | `LockediPod_Screen` | Known | Screen layout |
| 0x006CEF59 | `Lock_Screen` | Known | Screen layout |
| 0x006CEF68 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006CEFCB | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006CF02D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006CF049 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006CF0BB | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006CF0DA | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006CF142 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CF15C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006CF1C4 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CF1E1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CF24D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006CF2B7 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006CF2D1 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006CF341 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006CF3B4 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006CF425 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006CF494 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006CF500 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006CF51B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006CF590 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006CF5F7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006CF659 | `Photos_Screen` | Known | Screen layout |
| 0x006CF6BD | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006CF6DB | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006CF74D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006CF76A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006CF7D0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006CF7EB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006CF854 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006CF871 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006CF8E8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006CF90C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006CF97A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006CF995 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006CFA50 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CFA6C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CFADA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CFAF7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CFB62 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006CFB82 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006CFBF9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CFC15 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CFC85 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006CFCA4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006CFD10 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006CFD24 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006CFD9D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006CFE11 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006CFE81 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006CFEE9 | `NoContent_Screen` | Known | Screen layout |
| 0x006CFEFD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006CFF61 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006CFFC8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CFFE2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006D0050 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006D00C2 | `NoContent_Screen` | Known | Screen layout |
| 0x006D00D6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006D0140 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006D01A9 | `No_Photos_Screen` | Known | Screen layout |
| 0x006D01BD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006D0223 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D0291 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006D02FE | `NoContent_Screen` | Known | Screen layout |
| 0x006D0312 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006D037A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006D03E4 | `NoContent_Screen` | Known | Screen layout |
| 0x006D03F8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006D045F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006D04C9 | `NoContent_Screen` | Known | Screen layout |
| 0x006D04DD | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006D054A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006D05BC | `NoContent_Screen` | Known | Screen layout |
| 0x006D05D0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D0638 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006D06A1 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006D06BC | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006D0722 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006D073E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006D081D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006D0836 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006D0897 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006D08AB | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006D0A19 | `Radio_Screen` | Known | Screen layout |
| 0x006D0A29 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006D0A8A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006D0B0D | `LockediPod_Screen` | Known | Screen layout |
| 0x006D0B95 | `Lock_Screen` | Known | Screen layout |
| 0x006D0BA4 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006D0C07 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006D0C69 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006D0C85 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006D0CF7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006D0D16 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006D0D7E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D0D98 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006D0E00 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D0E1D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D0E89 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006D0EF3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006D0F0D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006D0F7D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006D0FF0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006D1061 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006D10D0 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006D113C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006D1157 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006D11CC | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006D1233 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006D1295 | `Photos_Screen` | Known | Screen layout |
| 0x006D12F9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006D1317 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006D1389 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006D13A6 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006D140C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006D1427 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006D1490 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006D14AD | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006D1524 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006D1548 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006D15B6 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006D15D1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006D168C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D16A8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D1716 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D1733 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D179E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006D17BE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006D1835 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D1851 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D18C1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006D18E0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006D194C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006D1960 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006D19D9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006D1A4D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006D1ABD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006D1B25 | `NoContent_Screen` | Known | Screen layout |
| 0x006D1B39 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006D1B9D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006D1C04 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D1C1E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006D1C8C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006D1CFE | `NoContent_Screen` | Known | Screen layout |
| 0x006D1D12 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006D1D7C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006D1DE5 | `No_Photos_Screen` | Known | Screen layout |
| 0x006D1DF9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006D1E5F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D1ECD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006D1F3A | `NoContent_Screen` | Known | Screen layout |
| 0x006D1F4E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006D1FB6 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006D2020 | `NoContent_Screen` | Known | Screen layout |
| 0x006D2034 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006D209B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006D2105 | `NoContent_Screen` | Known | Screen layout |
| 0x006D2119 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006D2186 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006D21F8 | `NoContent_Screen` | Known | Screen layout |
| 0x006D220C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D2274 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006D22DD | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006D22F8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006D235E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006D237A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006D2459 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006D2472 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006D24D3 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006D24E7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006D2655 | `Radio_Screen` | Known | Screen layout |
| 0x006D2665 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006D26C6 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006D2749 | `LockediPod_Screen` | Known | Screen layout |
| 0x006D27D1 | `Lock_Screen` | Known | Screen layout |
| 0x006D27E0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006D2843 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006D28A5 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006D28C1 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006D2933 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006D2952 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006D29BA | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D29D4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006D2A3C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D2A59 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D2AC5 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006D2B2F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006D2B49 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006D2BB9 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006D2C2C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006D2C9D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006D2D0C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006D2D78 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006D2D93 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006D2E08 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006D2E6F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006D2ED1 | `Photos_Screen` | Known | Screen layout |
| 0x006D2F35 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006D2F53 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006D2FC5 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006D2FE2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006D3048 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006D3063 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006D30CC | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006D30E9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006D3160 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006D3184 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006D31F2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006D320D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006D32C8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D32E4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D3352 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D336F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D33DA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006D33FA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006D3471 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D348D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D34FD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006D351C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006D3588 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006D359C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006D3615 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006D3689 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006D36F9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006D3761 | `NoContent_Screen` | Known | Screen layout |
| 0x006D3775 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006D37D9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006D3840 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D385A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006D38C8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006D393A | `NoContent_Screen` | Known | Screen layout |
| 0x006D394E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006D39B8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006D3A21 | `No_Photos_Screen` | Known | Screen layout |
| 0x006D3A35 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006D3A9B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D3B09 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006D3B76 | `NoContent_Screen` | Known | Screen layout |
| 0x006D3B8A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006D3BF2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006D3C5C | `NoContent_Screen` | Known | Screen layout |
| 0x006D3C70 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006D3CD7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006D3D41 | `NoContent_Screen` | Known | Screen layout |
| 0x006D3D55 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006D3DC2 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006D3E34 | `NoContent_Screen` | Known | Screen layout |
| 0x006D3E48 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D3EB0 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006D3F19 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006D3F34 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006D3F9A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006D3FB6 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006D4095 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006D40AE | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006D410F | `FirstBoot_Screen` | Known | Screen layout |
| 0x006D4123 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006D4291 | `Radio_Screen` | Known | Screen layout |
| 0x006D42A1 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006D4302 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006D4385 | `LockediPod_Screen` | Known | Screen layout |
| 0x006D440D | `Lock_Screen` | Known | Screen layout |
| 0x006D441C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006D447F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006D44E1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006D44FD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006D456F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006D458E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006D45F6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D4610 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006D4678 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D4695 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D4701 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006D476B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006D4785 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006D47F5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006D4868 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006D48D9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006D4948 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006D49B4 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006D49CF | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006D4A44 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006D4AAB | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006D4B0D | `Photos_Screen` | Known | Screen layout |
| 0x006D4B71 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006D4B8F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006D4C01 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006D4C1E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006D4C84 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006D4C9F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006D4D08 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006D4D25 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006D4D9C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006D4DC0 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006D4E2E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006D4E49 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006D4F04 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D4F20 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D4F8E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D4FAB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D5016 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006D5036 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006D50AD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D50C9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D5139 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006D5158 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006D51C4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006D51D8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006D5251 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006D52C5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006D5335 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006D539D | `NoContent_Screen` | Known | Screen layout |
| 0x006D53B1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006D5415 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006D547C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D5496 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006D5504 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006D5576 | `NoContent_Screen` | Known | Screen layout |
| 0x006D558A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006D55F4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006D565D | `No_Photos_Screen` | Known | Screen layout |
| 0x006D5671 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006D56D7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D5745 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006D57B2 | `NoContent_Screen` | Known | Screen layout |
| 0x006D57C6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006D582E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006D5898 | `NoContent_Screen` | Known | Screen layout |
| 0x006D58AC | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006D5913 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006D597D | `NoContent_Screen` | Known | Screen layout |
| 0x006D5991 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006D59FE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006D5A70 | `NoContent_Screen` | Known | Screen layout |
| 0x006D5A84 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D5AEC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006D5B55 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006D5B70 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006D5BD6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006D5BF2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006D5CD1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006D5CEA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006D5D4B | `FirstBoot_Screen` | Known | Screen layout |
| 0x006D5D5F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006D5ECD | `Radio_Screen` | Known | Screen layout |
| 0x006D5EDD | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006D5F3E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006D5FC1 | `LockediPod_Screen` | Known | Screen layout |
| 0x006D6049 | `Lock_Screen` | Known | Screen layout |
| 0x006D6058 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006D60BB | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006D611D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006D6139 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006D61AB | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006D61CA | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006D6232 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D624C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006D62B4 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D62D1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D633D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006D63A7 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006D63C1 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006D6431 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006D64A4 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006D6515 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006D6584 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006D65F0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006D660B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006D6680 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006D66E7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006D6749 | `Photos_Screen` | Known | Screen layout |
| 0x006D67AD | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006D67CB | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006D683D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006D685A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006D68C0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006D68DB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006D6944 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006D6961 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006D69D8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006D69FC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006D6A6A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006D6A85 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006D6B40 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D6B5C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D6BCA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D6BE7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D6C52 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006D6C72 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006D6CE9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D6D05 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D6D75 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006D6D94 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006D6E00 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006D6E14 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006D6E8D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006D6F01 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006D6F71 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006D6FD9 | `NoContent_Screen` | Known | Screen layout |
| 0x006D6FED | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006D7051 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006D70B8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D70D2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006D7140 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006D71B2 | `NoContent_Screen` | Known | Screen layout |
| 0x006D71C6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006D7230 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006D7299 | `No_Photos_Screen` | Known | Screen layout |
| 0x006D72AD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006D7313 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D7381 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006D73EE | `NoContent_Screen` | Known | Screen layout |
| 0x006D7402 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006D746A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006D74D4 | `NoContent_Screen` | Known | Screen layout |
| 0x006D74E8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006D754F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006D75B9 | `NoContent_Screen` | Known | Screen layout |
| 0x006D75CD | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006D763A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006D76AC | `NoContent_Screen` | Known | Screen layout |
| 0x006D76C0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D7728 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006D7791 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006D77AC | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006D7812 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006D782E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006D790D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006D7926 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006D7987 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006D799B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006D7B09 | `Radio_Screen` | Known | Screen layout |
| 0x006D7B19 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006D7B7A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006D7BFD | `LockediPod_Screen` | Known | Screen layout |
| 0x006D7C85 | `Lock_Screen` | Known | Screen layout |
| 0x006D7C94 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006D7CF7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006D7D59 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006D7D75 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006D7DE7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006D7E06 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006D7E6E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D7E88 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006D7EF0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D7F0D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D7F79 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006D7FE3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006D7FFD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006D806D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006D80E0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006D8151 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006D81C0 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006D822C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006D8247 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006D82BC | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006D8323 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006D8385 | `Photos_Screen` | Known | Screen layout |
| 0x006D83E9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006D8407 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006D8479 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006D8496 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006D84FC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006D8517 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006D8580 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006D859D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006D8614 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006D8638 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006D86A6 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006D86C1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006D877C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D8798 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D8806 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D8823 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D888E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006D88AE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006D8925 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D8941 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D89B1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006D89D0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006D8A3C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006D8A50 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006D8AC9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006D8B3D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006D8BAD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006D8C15 | `NoContent_Screen` | Known | Screen layout |
| 0x006D8C29 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006D8C8D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006D8CF4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D8D0E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006D8D7C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006D8DEE | `NoContent_Screen` | Known | Screen layout |
| 0x006D8E02 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006D8E6C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006D8ED5 | `No_Photos_Screen` | Known | Screen layout |
| 0x006D8EE9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006D8F4F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D8FBD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006D902A | `NoContent_Screen` | Known | Screen layout |
| 0x006D903E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006D90A6 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006D9110 | `NoContent_Screen` | Known | Screen layout |
| 0x006D9124 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006D918B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006D91F5 | `NoContent_Screen` | Known | Screen layout |
| 0x006D9209 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006D9276 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006D92E8 | `NoContent_Screen` | Known | Screen layout |
| 0x006D92FC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D9364 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006D93CD | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006D93E8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006D944E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006D946A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006D9549 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006D9562 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006D95C3 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006D95D7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006D9745 | `Radio_Screen` | Known | Screen layout |
| 0x006D9755 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006D97B6 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006D9839 | `LockediPod_Screen` | Known | Screen layout |
| 0x006D98C1 | `Lock_Screen` | Known | Screen layout |
| 0x006D98D0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006D9933 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006D9995 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006D99B1 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006D9A23 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006D9A42 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006D9AAA | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D9AC4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006D9B2C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D9B49 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D9BB5 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006D9C1F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006D9C39 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006D9CA9 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006D9D1C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006D9D8D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006D9DFC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006D9E68 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006D9E83 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006D9EF8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006D9F5F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006D9FC1 | `Photos_Screen` | Known | Screen layout |
| 0x006DA025 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006DA043 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006DA0B5 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006DA0D2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006DA138 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006DA153 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006DA1BC | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006DA1D9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006DA250 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006DA274 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006DA2E2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006DA2FD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006DA3B8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DA3D4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DA442 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DA45F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DA4CA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006DA4EA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006DA561 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DA57D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DA5ED | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006DA60C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006DA678 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006DA68C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006DA705 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006DA779 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006DA7E9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006DA851 | `NoContent_Screen` | Known | Screen layout |
| 0x006DA865 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006DA8C9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006DA930 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DA94A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006DA9B8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006DAA2A | `NoContent_Screen` | Known | Screen layout |
| 0x006DAA3E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006DAAA8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006DAB11 | `No_Photos_Screen` | Known | Screen layout |
| 0x006DAB25 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006DAB8B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DABF9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006DAC66 | `NoContent_Screen` | Known | Screen layout |
| 0x006DAC7A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006DACE2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006DAD4C | `NoContent_Screen` | Known | Screen layout |
| 0x006DAD60 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006DADC7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006DAE31 | `NoContent_Screen` | Known | Screen layout |
| 0x006DAE45 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006DAEB2 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006DAF24 | `NoContent_Screen` | Known | Screen layout |
| 0x006DAF38 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DAFA0 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006DB009 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006DB024 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006DB08A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006DB0A6 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006DB185 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006DB19E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006DB1FF | `FirstBoot_Screen` | Known | Screen layout |
| 0x006DB213 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006DB381 | `Radio_Screen` | Known | Screen layout |
| 0x006DB391 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006DB3F2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006DB475 | `LockediPod_Screen` | Known | Screen layout |
| 0x006DB4FD | `Lock_Screen` | Known | Screen layout |
| 0x006DB50C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006DB56F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006DB5D1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006DB5ED | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006DB65F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006DB67E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006DB6E6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DB700 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006DB768 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DB785 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DB7F1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006DB85B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006DB875 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006DB8E5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006DB958 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006DB9C9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006DBA38 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006DBAA4 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006DBABF | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006DBB34 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006DBB9B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006DBBFD | `Photos_Screen` | Known | Screen layout |
| 0x006DBC61 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006DBC7F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006DBCF1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006DBD0E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006DBD74 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006DBD8F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006DBDF8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006DBE15 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006DBE8C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006DBEB0 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006DBF1E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006DBF39 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006DBFF4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DC010 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DC07E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DC09B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DC106 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006DC126 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006DC19D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DC1B9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DC229 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006DC248 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006DC2B4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006DC2C8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006DC341 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006DC3B5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006DC425 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006DC48D | `NoContent_Screen` | Known | Screen layout |
| 0x006DC4A1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006DC505 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006DC56C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DC586 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006DC5F4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006DC666 | `NoContent_Screen` | Known | Screen layout |
| 0x006DC67A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006DC6E4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006DC74D | `No_Photos_Screen` | Known | Screen layout |
| 0x006DC761 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006DC7C7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DC835 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006DC8A2 | `NoContent_Screen` | Known | Screen layout |
| 0x006DC8B6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006DC91E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006DC988 | `NoContent_Screen` | Known | Screen layout |
| 0x006DC99C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006DCA03 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006DCA6D | `NoContent_Screen` | Known | Screen layout |
| 0x006DCA81 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006DCAEE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006DCB60 | `NoContent_Screen` | Known | Screen layout |
| 0x006DCB74 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DCBDC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006DCC45 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006DCC60 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006DCCC6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006DCCE2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006DCDC1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006DCDDA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006DCE3B | `FirstBoot_Screen` | Known | Screen layout |
| 0x006DCE4F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006DCFBD | `Radio_Screen` | Known | Screen layout |
| 0x006DCFCD | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006DD02E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006DD0B1 | `LockediPod_Screen` | Known | Screen layout |
| 0x006DD139 | `Lock_Screen` | Known | Screen layout |
| 0x006DD148 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006DD1AB | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006DD20D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006DD229 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006DD29B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006DD2BA | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006DD322 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DD33C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006DD3A4 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DD3C1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DD42D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006DD497 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006DD4B1 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006DD521 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006DD594 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006DD605 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006DD674 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006DD6E0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006DD6FB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006DD770 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006DD7D7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006DD839 | `Photos_Screen` | Known | Screen layout |
| 0x006DD89D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006DD8BB | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006DD92D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006DD94A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006DD9B0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006DD9CB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006DDA34 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006DDA51 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006DDAC8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006DDAEC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006DDB5A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006DDB75 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006DDC30 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DDC4C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DDCBA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DDCD7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DDD42 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006DDD62 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006DDDD9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DDDF5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DDE65 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006DDE84 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006DDEF0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006DDF04 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006DDF7D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006DDFF1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006DE061 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006DE0C9 | `NoContent_Screen` | Known | Screen layout |
| 0x006DE0DD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006DE141 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006DE1A8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DE1C2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006DE230 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006DE2A2 | `NoContent_Screen` | Known | Screen layout |
| 0x006DE2B6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006DE320 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006DE389 | `No_Photos_Screen` | Known | Screen layout |
| 0x006DE39D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006DE403 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DE471 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006DE4DE | `NoContent_Screen` | Known | Screen layout |
| 0x006DE4F2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006DE55A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006DE5C4 | `NoContent_Screen` | Known | Screen layout |
| 0x006DE5D8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006DE63F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006DE6A9 | `NoContent_Screen` | Known | Screen layout |
| 0x006DE6BD | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006DE72A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006DE79C | `NoContent_Screen` | Known | Screen layout |
| 0x006DE7B0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DE818 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006DE881 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006DE89C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006DE902 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006DE91E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006DE9FD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006DEA16 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006DEA77 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006DEA8B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006DEBF9 | `Radio_Screen` | Known | Screen layout |
| 0x006DEC09 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006DEC6A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006DECED | `LockediPod_Screen` | Known | Screen layout |
| 0x006DED75 | `Lock_Screen` | Known | Screen layout |
| 0x006DED84 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006DEDE7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006DEE49 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006DEE65 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006DEED7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006DEEF6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006DEF5E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DEF78 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006DEFE0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DEFFD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DF069 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006DF0D3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006DF0ED | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006DF15D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006DF1D0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006DF241 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006DF2B0 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006DF31C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006DF337 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006DF3AC | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006DF413 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006DF475 | `Photos_Screen` | Known | Screen layout |
| 0x006DF4D9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006DF4F7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006DF569 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006DF586 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006DF5EC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006DF607 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006DF670 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006DF68D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006DF704 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006DF728 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006DF796 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006DF7B1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006DF86C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DF888 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DF8F6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DF913 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DF97E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006DF99E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006DFA15 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DFA31 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DFAA1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006DFAC0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006DFB2C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006DFB40 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006DFBB9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006DFC2D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006DFC9D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006DFD05 | `NoContent_Screen` | Known | Screen layout |
| 0x006DFD19 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006DFD7D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006DFDE4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DFDFE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006DFE6C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006DFEDE | `NoContent_Screen` | Known | Screen layout |
| 0x006DFEF2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006DFF5C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006DFFC5 | `No_Photos_Screen` | Known | Screen layout |
| 0x006DFFD9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006E003F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E00AD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006E011A | `NoContent_Screen` | Known | Screen layout |
| 0x006E012E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006E0196 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006E0200 | `NoContent_Screen` | Known | Screen layout |
| 0x006E0214 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006E027B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006E02E5 | `NoContent_Screen` | Known | Screen layout |
| 0x006E02F9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006E0366 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006E03D8 | `NoContent_Screen` | Known | Screen layout |
| 0x006E03EC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E0454 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006E04BD | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006E04D8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006E053E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006E055A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006E0639 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006E0652 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006E06B3 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006E06C7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006E0835 | `Radio_Screen` | Known | Screen layout |
| 0x006E0845 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006E08A6 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006E0929 | `LockediPod_Screen` | Known | Screen layout |
| 0x006E09B1 | `Lock_Screen` | Known | Screen layout |
| 0x006E09C0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006E0A23 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006E0A85 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006E0AA1 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006E0B13 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E0B32 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E0B9A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E0BB4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006E0C1C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E0C39 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E0CA5 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E0D0F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006E0D29 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006E0D99 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006E0E0C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006E0E7D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006E0EEC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006E0F58 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006E0F73 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006E0FE8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006E104F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006E10B1 | `Photos_Screen` | Known | Screen layout |
| 0x006E1115 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006E1133 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006E11A5 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006E11C2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006E1228 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E1243 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E12AC | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006E12C9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006E1340 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006E1364 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006E13D2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006E13ED | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006E14A8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E14C4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E1532 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E154F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E15BA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006E15DA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006E1651 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E166D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E16DD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006E16FC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006E1768 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006E177C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006E17F5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006E1869 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006E18D9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006E1941 | `NoContent_Screen` | Known | Screen layout |
| 0x006E1955 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006E19B9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006E1A20 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E1A3A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006E1AA8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006E1B1A | `NoContent_Screen` | Known | Screen layout |
| 0x006E1B2E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006E1B98 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006E1C01 | `No_Photos_Screen` | Known | Screen layout |
| 0x006E1C15 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006E1C7B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E1CE9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006E1D56 | `NoContent_Screen` | Known | Screen layout |
| 0x006E1D6A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006E1DD2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006E1E3C | `NoContent_Screen` | Known | Screen layout |
| 0x006E1E50 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006E1EB7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006E1F21 | `NoContent_Screen` | Known | Screen layout |
| 0x006E1F35 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006E1FA2 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006E2014 | `NoContent_Screen` | Known | Screen layout |
| 0x006E2028 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E2090 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006E20F9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006E2114 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006E217A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006E2196 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006E2275 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006E228E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006E22EF | `FirstBoot_Screen` | Known | Screen layout |
| 0x006E2303 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006E2471 | `Radio_Screen` | Known | Screen layout |
| 0x006E2481 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006E24E2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006E2565 | `LockediPod_Screen` | Known | Screen layout |
| 0x006E25ED | `Lock_Screen` | Known | Screen layout |
| 0x006E25FC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006E265F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006E26C1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006E26DD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006E274F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E276E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E27D6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E27F0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006E2858 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E2875 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E28E1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E294B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006E2965 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006E29D5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006E2A48 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006E2AB9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006E2B28 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006E2B94 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006E2BAF | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006E2C24 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006E2C8B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006E2CED | `Photos_Screen` | Known | Screen layout |
| 0x006E2D51 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006E2D6F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006E2DE1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006E2DFE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006E2E64 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E2E7F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E2EE8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006E2F05 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006E2F7C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006E2FA0 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006E300E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006E3029 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006E30E4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E3100 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E316E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E318B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E31F6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006E3216 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006E328D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E32A9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E3319 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006E3338 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006E33A4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006E33B8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006E3431 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006E34A5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006E3515 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006E357D | `NoContent_Screen` | Known | Screen layout |
| 0x006E3591 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006E35F5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006E365C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E3676 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006E36E4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006E3756 | `NoContent_Screen` | Known | Screen layout |
| 0x006E376A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006E37D4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006E383D | `No_Photos_Screen` | Known | Screen layout |
| 0x006E3851 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006E38B7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E3925 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006E3992 | `NoContent_Screen` | Known | Screen layout |
| 0x006E39A6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006E3A0E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006E3A78 | `NoContent_Screen` | Known | Screen layout |
| 0x006E3A8C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006E3AF3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006E3B5D | `NoContent_Screen` | Known | Screen layout |
| 0x006E3B71 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006E3BDE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006E3C50 | `NoContent_Screen` | Known | Screen layout |
| 0x006E3C64 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E3CCC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006E3D35 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006E3D50 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006E3DB6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006E3DD2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006E3EB1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006E3ECA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006E3F2B | `FirstBoot_Screen` | Known | Screen layout |
| 0x006E3F3F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006E40AD | `Radio_Screen` | Known | Screen layout |
| 0x006E40BD | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006E411E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006E41A1 | `LockediPod_Screen` | Known | Screen layout |
| 0x006E4229 | `Lock_Screen` | Known | Screen layout |
| 0x006E4238 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006E429B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006E42FD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006E4319 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006E438B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E43AA | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E4412 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E442C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006E4494 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E44B1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E451D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E4587 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006E45A1 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006E4611 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006E4684 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006E46F5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006E4764 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006E47D0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006E47EB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006E4860 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006E48C7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006E4929 | `Photos_Screen` | Known | Screen layout |
| 0x006E498D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006E49AB | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006E4A1D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006E4A3A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006E4AA0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E4ABB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E4B24 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006E4B41 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006E4BB8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006E4BDC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006E4C4A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006E4C65 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006E4D20 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E4D3C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E4DAA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E4DC7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E4E32 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006E4E52 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006E4EC9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E4EE5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E4F55 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006E4F74 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006E4FE0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006E4FF4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006E506D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006E50E1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006E5151 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006E51B9 | `NoContent_Screen` | Known | Screen layout |
| 0x006E51CD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006E5231 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006E5298 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E52B2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006E5320 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006E5392 | `NoContent_Screen` | Known | Screen layout |
| 0x006E53A6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006E5410 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006E5479 | `No_Photos_Screen` | Known | Screen layout |
| 0x006E548D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006E54F3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E5561 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006E55CE | `NoContent_Screen` | Known | Screen layout |
| 0x006E55E2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006E564A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006E56B4 | `NoContent_Screen` | Known | Screen layout |
| 0x006E56C8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006E572F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006E5799 | `NoContent_Screen` | Known | Screen layout |
| 0x006E57AD | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006E581A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006E588C | `NoContent_Screen` | Known | Screen layout |
| 0x006E58A0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E5908 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006E5971 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006E598C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006E59F2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006E5A0E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006E5AED | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006E5B06 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006E5B67 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006E5B7B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006E5CE9 | `Radio_Screen` | Known | Screen layout |
| 0x006E5CF9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006E5D5A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006E5DDD | `LockediPod_Screen` | Known | Screen layout |
| 0x006E5E65 | `Lock_Screen` | Known | Screen layout |
| 0x006E5E74 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006E5ED7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006E5F39 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006E5F55 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006E5FC7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E5FE6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E604E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E6068 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006E60D0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E60ED | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E6159 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E61C3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006E61DD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006E624D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006E62C0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006E6331 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006E63A0 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006E640C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006E6427 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006E649C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006E6503 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006E6565 | `Photos_Screen` | Known | Screen layout |
| 0x006E65C9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006E65E7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006E6659 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006E6676 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006E66DC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E66F7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E6760 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006E677D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006E67F4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006E6818 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006E6886 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006E68A1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006E695C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E6978 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E69E6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E6A03 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E6A6E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006E6A8E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006E6B05 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E6B21 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E6B91 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006E6BB0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006E6C1C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006E6C30 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006E6CA9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006E6D1D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006E6D8D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006E6DF5 | `NoContent_Screen` | Known | Screen layout |
| 0x006E6E09 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006E6E6D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006E6ED4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E6EEE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006E6F5C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006E6FCE | `NoContent_Screen` | Known | Screen layout |
| 0x006E6FE2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006E704C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006E70B5 | `No_Photos_Screen` | Known | Screen layout |
| 0x006E70C9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006E712F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E719D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006E720A | `NoContent_Screen` | Known | Screen layout |
| 0x006E721E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006E7286 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006E72F0 | `NoContent_Screen` | Known | Screen layout |
| 0x006E7304 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006E736B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006E73D5 | `NoContent_Screen` | Known | Screen layout |
| 0x006E73E9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006E7456 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006E74C8 | `NoContent_Screen` | Known | Screen layout |
| 0x006E74DC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E7544 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006E75AD | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006E75C8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006E762E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006E764A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006E7729 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006E7742 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006E77A3 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006E77B7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006E7925 | `Radio_Screen` | Known | Screen layout |
| 0x006E7935 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006E7996 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006E7A19 | `LockediPod_Screen` | Known | Screen layout |
| 0x006E7AA1 | `Lock_Screen` | Known | Screen layout |
| 0x006E7AB0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006E7B13 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006E7B75 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006E7B91 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006E7C03 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E7C22 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E7C8A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E7CA4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006E7D0C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E7D29 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E7D95 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E7DFF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006E7E19 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006E7E89 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006E7EFC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006E7F6D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006E7FDC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006E8048 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006E8063 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006E80D8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006E813F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006E81A1 | `Photos_Screen` | Known | Screen layout |
| 0x006E8205 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006E8223 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006E8295 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006E82B2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006E8318 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E8333 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E839C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006E83B9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006E8430 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006E8454 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006E84C2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006E84DD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006E8598 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E85B4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E8622 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E863F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E86AA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006E86CA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006E8741 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E875D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E87CD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006E87EC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006E8858 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006E886C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006E88E5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006E8959 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006E89C9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006E8A31 | `NoContent_Screen` | Known | Screen layout |
| 0x006E8A45 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006E8AA9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006E8B10 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E8B2A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006E8B98 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006E8C0A | `NoContent_Screen` | Known | Screen layout |
| 0x006E8C1E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006E8C88 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006E8CF1 | `No_Photos_Screen` | Known | Screen layout |
| 0x006E8D05 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006E8D6B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E8DD9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006E8E46 | `NoContent_Screen` | Known | Screen layout |
| 0x006E8E5A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006E8EC2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006E8F2C | `NoContent_Screen` | Known | Screen layout |
| 0x006E8F40 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006E8FA7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006E9011 | `NoContent_Screen` | Known | Screen layout |
| 0x006E9025 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006E9092 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006E9104 | `NoContent_Screen` | Known | Screen layout |
| 0x006E9118 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E9180 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006E91E9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006E9204 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006E926A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006E9286 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006E9365 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006E937E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006E93DF | `FirstBoot_Screen` | Known | Screen layout |
| 0x006E93F3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006E9561 | `Radio_Screen` | Known | Screen layout |
| 0x006E9571 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006E95D2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006E9655 | `LockediPod_Screen` | Known | Screen layout |
| 0x006E96DD | `Lock_Screen` | Known | Screen layout |
| 0x006E96EC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006E974F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006E97B1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006E97CD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006E983F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E985E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E98C6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E98E0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006E9948 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E9965 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E99D1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E9A3B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006E9A55 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006E9AC5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006E9B38 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006E9BA9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006E9C18 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006E9C84 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006E9C9F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006E9D14 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006E9D7B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006E9DDD | `Photos_Screen` | Known | Screen layout |
| 0x006E9E41 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006E9E5F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006E9ED1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006E9EEE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006E9F54 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E9F6F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E9FD8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006E9FF5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006EA06C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006EA090 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006EA0FE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006EA119 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006EA1D4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EA1F0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EA25E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006EA27B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006EA2E6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006EA306 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006EA37D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EA399 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EA409 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006EA428 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006EA494 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006EA4A8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006EA521 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006EA595 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006EA605 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006EA66D | `NoContent_Screen` | Known | Screen layout |
| 0x006EA681 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006EA6E5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006EA74C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006EA766 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006EA7D4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006EA846 | `NoContent_Screen` | Known | Screen layout |
| 0x006EA85A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006EA8C4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006EA92D | `No_Photos_Screen` | Known | Screen layout |
| 0x006EA941 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006EA9A7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006EAA15 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006EAA82 | `NoContent_Screen` | Known | Screen layout |
| 0x006EAA96 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006EAAFE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006EAB68 | `NoContent_Screen` | Known | Screen layout |
| 0x006EAB7C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006EABE3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006EAC4D | `NoContent_Screen` | Known | Screen layout |
| 0x006EAC61 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006EACCE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006EAD40 | `NoContent_Screen` | Known | Screen layout |
| 0x006EAD54 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006EADBC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006EAE25 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006EAE40 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006EAEA6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006EAEC2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006EAFA1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006EAFBA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006EB01B | `FirstBoot_Screen` | Known | Screen layout |
| 0x006EB02F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006EB19D | `Radio_Screen` | Known | Screen layout |
| 0x006EB1AD | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006EB20E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006EB291 | `LockediPod_Screen` | Known | Screen layout |
| 0x006EB319 | `Lock_Screen` | Known | Screen layout |
| 0x006EB328 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006EB38B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006EB3ED | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006EB409 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006EB47B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006EB49A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006EB502 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006EB51C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006EB584 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006EB5A1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006EB60D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006EB677 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006EB691 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006EB701 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006EB774 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006EB7E5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006EB854 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006EB8C0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006EB8DB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006EB950 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006EB9B7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006EBA19 | `Photos_Screen` | Known | Screen layout |
| 0x006EBA7D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006EBA9B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006EBB0D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006EBB2A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006EBB90 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006EBBAB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006EBC14 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006EBC31 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006EBCA8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006EBCCC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006EBD3A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006EBD55 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006EBE10 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EBE2C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EBE9A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006EBEB7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006EBF22 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006EBF42 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006EBFB9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EBFD5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EC045 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006EC064 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006EC0D0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006EC0E4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006EC15D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006EC1D1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006EC241 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006EC2A9 | `NoContent_Screen` | Known | Screen layout |
| 0x006EC2BD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006EC321 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006EC388 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006EC3A2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006EC410 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006EC482 | `NoContent_Screen` | Known | Screen layout |
| 0x006EC496 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006EC500 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006EC569 | `No_Photos_Screen` | Known | Screen layout |
| 0x006EC57D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006EC5E3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006EC651 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006EC6BE | `NoContent_Screen` | Known | Screen layout |
| 0x006EC6D2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006EC73A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006EC7A4 | `NoContent_Screen` | Known | Screen layout |
| 0x006EC7B8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006EC81F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006EC889 | `NoContent_Screen` | Known | Screen layout |
| 0x006EC89D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006EC90A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006EC97C | `NoContent_Screen` | Known | Screen layout |
| 0x006EC990 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006EC9F8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006ECA61 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006ECA7C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006ECAE2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006ECAFE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006ECBDD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006ECBF6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006ECC57 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006ECC6B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006ECDD9 | `Radio_Screen` | Known | Screen layout |
| 0x006ECDE9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006ECE4A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006ECECD | `LockediPod_Screen` | Known | Screen layout |
| 0x006ECF55 | `Lock_Screen` | Known | Screen layout |
| 0x006ECF64 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006ECFC7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006ED029 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006ED045 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006ED0B7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006ED0D6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006ED13E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006ED158 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006ED1C0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006ED1DD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006ED249 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006ED2B3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006ED2CD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006ED33D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006ED3B0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006ED421 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006ED490 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006ED4FC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006ED517 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006ED58C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006ED5F3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006ED655 | `Photos_Screen` | Known | Screen layout |
| 0x006ED6B9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006ED6D7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006ED749 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006ED766 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006ED7CC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006ED7E7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006ED850 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006ED86D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006ED8E4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006ED908 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006ED976 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006ED991 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006EDA4C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EDA68 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EDAD6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006EDAF3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006EDB5E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006EDB7E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006EDBF5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EDC11 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EDC81 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006EDCA0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006EDD0C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006EDD20 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006EDD99 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006EDE0D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006EDE7D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006EDEE5 | `NoContent_Screen` | Known | Screen layout |
| 0x006EDEF9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006EDF5D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006EDFC4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006EDFDE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006EE04C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006EE0BE | `NoContent_Screen` | Known | Screen layout |
| 0x006EE0D2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006EE13C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006EE1A5 | `No_Photos_Screen` | Known | Screen layout |
| 0x006EE1B9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006EE21F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006EE28D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006EE2FA | `NoContent_Screen` | Known | Screen layout |
| 0x006EE30E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006EE376 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006EE3E0 | `NoContent_Screen` | Known | Screen layout |
| 0x006EE3F4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006EE45B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006EE4C5 | `NoContent_Screen` | Known | Screen layout |
| 0x006EE4D9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006EE546 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006EE5B8 | `NoContent_Screen` | Known | Screen layout |
| 0x006EE5CC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006EE634 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006EE69D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006EE6B8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006EE71E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006EE73A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006EE819 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006EE832 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006EE893 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006EE8A7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006EEA15 | `Radio_Screen` | Known | Screen layout |
| 0x006EEA25 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006EEA86 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006EEB09 | `LockediPod_Screen` | Known | Screen layout |
| 0x006EEB91 | `Lock_Screen` | Known | Screen layout |
| 0x006EEBA0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006EEC03 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006EEC65 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006EEC81 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006EECF3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006EED12 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006EED7A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006EED94 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006EEDFC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006EEE19 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006EEE85 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006EEEEF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006EEF09 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006EEF79 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006EEFEC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006EF05D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006EF0CC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006EF138 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006EF153 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006EF1C8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006EF22F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006EF291 | `Photos_Screen` | Known | Screen layout |
| 0x006EF2F5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006EF313 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006EF385 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006EF3A2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006EF408 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006EF423 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006EF48C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006EF4A9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006EF520 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006EF544 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006EF5B2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006EF5CD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006EF688 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EF6A4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EF712 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006EF72F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006EF79A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006EF7BA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006EF831 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EF84D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EF8BD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006EF8DC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006EF948 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006EF95C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006EF9D5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006EFA49 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006EFAB9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006EFB21 | `NoContent_Screen` | Known | Screen layout |
| 0x006EFB35 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006EFB99 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006EFC00 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006EFC1A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006EFC88 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006EFCFA | `NoContent_Screen` | Known | Screen layout |
| 0x006EFD0E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006EFD78 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006EFDE1 | `No_Photos_Screen` | Known | Screen layout |
| 0x006EFDF5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006EFE5B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006EFEC9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006EFF36 | `NoContent_Screen` | Known | Screen layout |
| 0x006EFF4A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006EFFB2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006F001C | `NoContent_Screen` | Known | Screen layout |
| 0x006F0030 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006F0097 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006F0101 | `NoContent_Screen` | Known | Screen layout |
| 0x006F0115 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F0182 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F01F4 | `NoContent_Screen` | Known | Screen layout |
| 0x006F0208 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F0270 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006F02D9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006F02F4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F035A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006F0376 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006F0455 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006F046E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006F04CF | `FirstBoot_Screen` | Known | Screen layout |
| 0x006F04E3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006F0651 | `Radio_Screen` | Known | Screen layout |
| 0x006F0661 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006F06C2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006F0745 | `LockediPod_Screen` | Known | Screen layout |
| 0x006F07CD | `Lock_Screen` | Known | Screen layout |
| 0x006F07DC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006F083F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006F08A1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F08BD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006F092F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006F094E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006F09B6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006F09D0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006F0A38 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F0A55 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F0AC1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006F0B2B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006F0B45 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006F0BB5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006F0C28 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006F0C99 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006F0D08 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F0D74 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006F0D8F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006F0E04 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006F0E6B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006F0ECD | `Photos_Screen` | Known | Screen layout |
| 0x006F0F31 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006F0F4F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006F0FC1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006F0FDE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006F1044 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006F105F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006F10C8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F10E5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F115C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F1180 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F11EE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006F1209 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F12C4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F12E0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F134E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F136B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F13D6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006F13F6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006F146D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F1489 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F14F9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006F1518 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006F1584 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006F1598 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006F1611 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006F1685 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006F16F5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006F175D | `NoContent_Screen` | Known | Screen layout |
| 0x006F1771 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006F17D5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006F183C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006F1856 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006F18C4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006F1936 | `NoContent_Screen` | Known | Screen layout |
| 0x006F194A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006F19B4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006F1A1D | `No_Photos_Screen` | Known | Screen layout |
| 0x006F1A31 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006F1A97 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F1B05 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006F1B72 | `NoContent_Screen` | Known | Screen layout |
| 0x006F1B86 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006F1BEE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006F1C58 | `NoContent_Screen` | Known | Screen layout |
| 0x006F1C6C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006F1CD3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006F1D3D | `NoContent_Screen` | Known | Screen layout |
| 0x006F1D51 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F1DBE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F1E30 | `NoContent_Screen` | Known | Screen layout |
| 0x006F1E44 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F1EAC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006F1F15 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006F1F30 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F1F96 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006F1FB2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006F2091 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006F20AA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006F210B | `FirstBoot_Screen` | Known | Screen layout |
| 0x006F211F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006F228D | `Radio_Screen` | Known | Screen layout |
| 0x006F229D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006F22FE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006F2381 | `LockediPod_Screen` | Known | Screen layout |
| 0x006F2409 | `Lock_Screen` | Known | Screen layout |
| 0x006F2418 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006F247B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006F24DD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F24F9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006F256B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006F258A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006F25F2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006F260C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006F2674 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F2691 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F26FD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006F2767 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006F2781 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006F27F1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006F2864 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006F28D5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006F2944 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F29B0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006F29CB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006F2A40 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006F2AA7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006F2B09 | `Photos_Screen` | Known | Screen layout |
| 0x006F2B6D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006F2B8B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006F2BFD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006F2C1A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006F2C80 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006F2C9B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006F2D04 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F2D21 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F2D98 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F2DBC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F2E2A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006F2E45 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F2F00 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F2F1C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F2F8A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F2FA7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F3012 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006F3032 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006F30A9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F30C5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F3135 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006F3154 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006F31C0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006F31D4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006F324D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006F32C1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006F3331 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006F3399 | `NoContent_Screen` | Known | Screen layout |
| 0x006F33AD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006F3411 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006F3478 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006F3492 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006F3500 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006F3572 | `NoContent_Screen` | Known | Screen layout |
| 0x006F3586 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006F35F0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006F3659 | `No_Photos_Screen` | Known | Screen layout |
| 0x006F366D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006F36D3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F3741 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006F37AE | `NoContent_Screen` | Known | Screen layout |
| 0x006F37C2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006F382A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006F3894 | `NoContent_Screen` | Known | Screen layout |
| 0x006F38A8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006F390F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006F3979 | `NoContent_Screen` | Known | Screen layout |
| 0x006F398D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F39FA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F3A6C | `NoContent_Screen` | Known | Screen layout |
| 0x006F3A80 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F3AE8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006F3B51 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006F3B6C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F3BD2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006F3BEE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006F3CCD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006F3CE6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006F3D47 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006F3D5B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006F3EC9 | `Radio_Screen` | Known | Screen layout |
| 0x006F3ED9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006F3F3A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006F3FBD | `LockediPod_Screen` | Known | Screen layout |
| 0x006F4045 | `Lock_Screen` | Known | Screen layout |
| 0x006F4054 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006F40B7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006F4119 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F4135 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006F41A7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006F41C6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006F422E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006F4248 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006F42B0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F42CD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F4339 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006F43A3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006F43BD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006F442D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006F44A0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006F4511 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006F4580 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F45EC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006F4607 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006F467C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006F46E3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006F4745 | `Photos_Screen` | Known | Screen layout |
| 0x006F47A9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006F47C7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006F4839 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006F4856 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006F48BC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006F48D7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006F4940 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F495D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F49D4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F49F8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F4A66 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006F4A81 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F4B3C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F4B58 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F4BC6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F4BE3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F4C4E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006F4C6E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006F4CE5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F4D01 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F4D71 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006F4D90 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006F4DFC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006F4E10 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006F4E89 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006F4EFD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006F4F6D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006F4FD5 | `NoContent_Screen` | Known | Screen layout |
| 0x006F4FE9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006F504D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006F50B4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006F50CE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006F513C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006F51AE | `NoContent_Screen` | Known | Screen layout |
| 0x006F51C2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006F522C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006F5295 | `No_Photos_Screen` | Known | Screen layout |
| 0x006F52A9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006F530F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F537D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006F53EA | `NoContent_Screen` | Known | Screen layout |
| 0x006F53FE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006F5466 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006F54D0 | `NoContent_Screen` | Known | Screen layout |
| 0x006F54E4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006F554B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006F55B5 | `NoContent_Screen` | Known | Screen layout |
| 0x006F55C9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F5636 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F56A8 | `NoContent_Screen` | Known | Screen layout |
| 0x006F56BC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F5724 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006F578D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006F57A8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F580E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006F582A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006F5909 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006F5922 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006F5983 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006F5997 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006F5B05 | `Radio_Screen` | Known | Screen layout |
| 0x006F5B15 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006F5B76 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006F5BF9 | `LockediPod_Screen` | Known | Screen layout |
| 0x006F5C81 | `Lock_Screen` | Known | Screen layout |
| 0x006F5C90 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006F5CF3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006F5D55 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F5D71 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006F5DE3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006F5E02 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006F5E6A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006F5E84 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006F5EEC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F5F09 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F5F75 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006F5FDF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006F5FF9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006F6069 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006F60DC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006F614D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006F61BC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F6228 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006F6243 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006F62B8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006F631F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006F6381 | `Photos_Screen` | Known | Screen layout |
| 0x006F63E5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006F6403 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006F6475 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006F6492 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006F64F8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006F6513 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006F657C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F6599 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F6610 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F6634 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F66A2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006F66BD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F6778 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F6794 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F6802 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F681F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F688A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006F68AA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006F6921 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F693D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F69AD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006F69CC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006F6A38 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006F6A4C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006F6AC5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006F6B39 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006F6BA9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006F6C11 | `NoContent_Screen` | Known | Screen layout |
| 0x006F6C25 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006F6C89 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006F6CF0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006F6D0A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006F6D78 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006F6DEA | `NoContent_Screen` | Known | Screen layout |
| 0x006F6DFE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006F6E68 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006F6ED1 | `No_Photos_Screen` | Known | Screen layout |
| 0x006F6EE5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006F6F4B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F6FB9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006F7026 | `NoContent_Screen` | Known | Screen layout |
| 0x006F703A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006F70A2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006F710C | `NoContent_Screen` | Known | Screen layout |
| 0x006F7120 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006F7187 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006F71F1 | `NoContent_Screen` | Known | Screen layout |
| 0x006F7205 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F7272 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F72E4 | `NoContent_Screen` | Known | Screen layout |
| 0x006F72F8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F7360 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006F73C9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006F73E4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F744A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006F7466 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006F7545 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006F755E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006F75BF | `FirstBoot_Screen` | Known | Screen layout |
| 0x006F75D3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006F7741 | `Radio_Screen` | Known | Screen layout |
| 0x006F7751 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006F77B2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006F7835 | `LockediPod_Screen` | Known | Screen layout |
| 0x006F78BD | `Lock_Screen` | Known | Screen layout |
| 0x006F78CC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006F792F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006F7991 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F79AD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006F7A1F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006F7A3E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006F7AA6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006F7AC0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006F7B28 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F7B45 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F7BB1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006F7C1B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006F7C35 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006F7CA5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006F7D18 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006F7D89 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006F7DF8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F7E64 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006F7E7F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006F7EF4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006F7F5B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006F7FBD | `Photos_Screen` | Known | Screen layout |
| 0x006F8021 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006F803F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006F80B1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006F80CE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006F8134 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006F814F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006F81B8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F81D5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F824C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F8270 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F82DE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006F82F9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F83B4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F83D0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F843E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F845B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F84C6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006F84E6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006F855D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F8579 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F85E9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006F8608 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006F8674 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006F8688 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006F8701 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006F8775 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006F87E5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006F884D | `NoContent_Screen` | Known | Screen layout |
| 0x006F8861 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006F88C5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006F892C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006F8946 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006F89B4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006F8A26 | `NoContent_Screen` | Known | Screen layout |
| 0x006F8A3A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006F8AA4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006F8B0D | `No_Photos_Screen` | Known | Screen layout |
| 0x006F8B21 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006F8B87 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F8BF5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006F8C62 | `NoContent_Screen` | Known | Screen layout |
| 0x006F8C76 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006F8CDE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006F8D48 | `NoContent_Screen` | Known | Screen layout |
| 0x006F8D5C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006F8DC3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006F8E2D | `NoContent_Screen` | Known | Screen layout |
| 0x006F8E41 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F8EAE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F8F20 | `NoContent_Screen` | Known | Screen layout |
| 0x006F8F34 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F8F9C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006F9005 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006F9020 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F9086 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006F90A2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006F9181 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006F919A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006F91FB | `FirstBoot_Screen` | Known | Screen layout |
| 0x006F920F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006F937D | `Radio_Screen` | Known | Screen layout |
| 0x006F938D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006F93EE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006F9471 | `LockediPod_Screen` | Known | Screen layout |
| 0x006F94F9 | `Lock_Screen` | Known | Screen layout |
| 0x006F9508 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006F956B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006F95CD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F95E9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006F965B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006F967A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006F96E2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006F96FC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006F9764 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F9781 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F97ED | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006F9857 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006F9871 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006F98E1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006F9954 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006F99C5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006F9A34 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F9AA0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006F9ABB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006F9B30 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006F9B97 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006F9BF9 | `Photos_Screen` | Known | Screen layout |
| 0x006F9C5D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006F9C7B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006F9CED | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006F9D0A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006F9D70 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006F9D8B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006F9DF4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F9E11 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F9E88 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F9EAC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F9F1A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006F9F35 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F9FF0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006FA00C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006FA07A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006FA097 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006FA102 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006FA122 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006FA199 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006FA1B5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006FA225 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006FA244 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006FA2B0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006FA2C4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006FA33D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006FA3B1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006FA421 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006FA489 | `NoContent_Screen` | Known | Screen layout |
| 0x006FA49D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006FA501 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006FA568 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006FA582 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006FA5F0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006FA662 | `NoContent_Screen` | Known | Screen layout |
| 0x006FA676 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006FA6E0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006FA749 | `No_Photos_Screen` | Known | Screen layout |
| 0x006FA75D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006FA7C3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006FA831 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006FA89E | `NoContent_Screen` | Known | Screen layout |
| 0x006FA8B2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006FA91A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006FA984 | `NoContent_Screen` | Known | Screen layout |
| 0x006FA998 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006FA9FF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006FAA69 | `NoContent_Screen` | Known | Screen layout |
| 0x006FAA7D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006FAAEA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006FAB5C | `NoContent_Screen` | Known | Screen layout |
| 0x006FAB70 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006FABD8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006FAC41 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006FAC5C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006FACC2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006FACDE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006FADBD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006FADD6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006FAE37 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006FAE4B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006FAFB9 | `Radio_Screen` | Known | Screen layout |
| 0x006FAFC9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006FB02A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006FB0AD | `LockediPod_Screen` | Known | Screen layout |
| 0x006FB135 | `Lock_Screen` | Known | Screen layout |
| 0x006FB144 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006FB1A7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006FB209 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006FB225 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006FB297 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006FB2B6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006FB31E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006FB338 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006FB3A0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006FB3BD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006FB429 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006FB493 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006FB4AD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006FB51D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006FB590 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006FB601 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006FB670 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006FB6DC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006FB6F7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006FB76C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006FB7D3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006FB835 | `Photos_Screen` | Known | Screen layout |
| 0x006FB899 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006FB8B7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006FB929 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006FB946 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006FB9AC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006FB9C7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006FBA30 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006FBA4D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006FBAC4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006FBAE8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006FBB56 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006FBB71 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006FBC2C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006FBC48 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006FBCB6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006FBCD3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006FBD3E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006FBD5E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006FBDD5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006FBDF1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006FBE61 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006FBE80 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006FBEEC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006FBF00 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006FBF79 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006FBFED | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006FC05D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006FC0C5 | `NoContent_Screen` | Known | Screen layout |
| 0x006FC0D9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006FC13D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006FC1A4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006FC1BE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006FC22C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006FC29E | `NoContent_Screen` | Known | Screen layout |
| 0x006FC2B2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006FC31C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006FC385 | `No_Photos_Screen` | Known | Screen layout |
| 0x006FC399 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006FC3FF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006FC46D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006FC4DA | `NoContent_Screen` | Known | Screen layout |
| 0x006FC4EE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006FC556 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006FC5C0 | `NoContent_Screen` | Known | Screen layout |
| 0x006FC5D4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006FC63B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006FC6A5 | `NoContent_Screen` | Known | Screen layout |
| 0x006FC6B9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006FC726 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006FC798 | `NoContent_Screen` | Known | Screen layout |
| 0x006FC7AC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006FC814 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006FC87D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006FC898 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006FC8FE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006FC91A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006FC9F9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006FCA12 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006FCA73 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006FCA87 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006FCBF5 | `Radio_Screen` | Known | Screen layout |
| 0x006FCC05 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006FCC66 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006FCCE9 | `LockediPod_Screen` | Known | Screen layout |
| 0x006FCD71 | `Lock_Screen` | Known | Screen layout |
| 0x006FCD80 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006FCDE3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006FCE45 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006FCE61 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006FCED3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006FCEF2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006FCF5A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006FCF74 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006FCFDC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006FCFF9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006FD065 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006FD0CF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006FD0E9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006FD159 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006FD1CC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006FD23D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006FD2AC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006FD318 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006FD333 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006FD3A8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006FD40F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006FD471 | `Photos_Screen` | Known | Screen layout |
| 0x006FD4D5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006FD4F3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006FD565 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006FD582 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006FD5E8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006FD603 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006FD66C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006FD689 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006FD700 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006FD724 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006FD792 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006FD7AD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006FD868 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006FD884 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006FD8F2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006FD90F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006FD97A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006FD99A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006FDA11 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006FDA2D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006FDA9D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006FDABC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006FDB28 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006FDB3C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006FDBB5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006FDC29 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006FDC99 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006FDD01 | `NoContent_Screen` | Known | Screen layout |
| 0x006FDD15 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006FDD79 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006FDDE0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006FDDFA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006FDE68 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006FDEDA | `NoContent_Screen` | Known | Screen layout |
| 0x006FDEEE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006FDF58 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006FDFC1 | `No_Photos_Screen` | Known | Screen layout |
| 0x006FDFD5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006FE03B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006FE0A9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006FE116 | `NoContent_Screen` | Known | Screen layout |
| 0x006FE12A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006FE192 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006FE1FC | `NoContent_Screen` | Known | Screen layout |
| 0x006FE210 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006FE277 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006FE2E1 | `NoContent_Screen` | Known | Screen layout |
| 0x006FE2F5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006FE362 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006FE3D4 | `NoContent_Screen` | Known | Screen layout |
| 0x006FE3E8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006FE450 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006FE4B9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006FE4D4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006FE53A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006FE556 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006FE635 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006FE64E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006FE6AF | `FirstBoot_Screen` | Known | Screen layout |
| 0x006FE6C3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006FE831 | `Radio_Screen` | Known | Screen layout |
| 0x006FE841 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006FE8A2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006FE925 | `LockediPod_Screen` | Known | Screen layout |
| 0x006FE9AD | `Lock_Screen` | Known | Screen layout |
| 0x006FE9BC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006FEA1F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006FEA81 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006FEA9D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006FEB0F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006FEB2E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006FEB96 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006FEBB0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006FEC18 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006FEC35 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006FECA1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006FED0B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006FED25 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006FED95 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006FEE08 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006FEE79 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006FEEE8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006FEF54 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006FEF6F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006FEFE4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006FF04B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006FF0AD | `Photos_Screen` | Known | Screen layout |
| 0x006FF111 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006FF12F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006FF1A1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x006FF1BE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x006FF224 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006FF23F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006FF2A8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006FF2C5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006FF33C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006FF360 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006FF3CE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006FF3E9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006FF4A4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006FF4C0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006FF52E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006FF54B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006FF5B6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006FF5D6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006FF64D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006FF669 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006FF6D9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006FF6F8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006FF764 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006FF778 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006FF7F1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006FF865 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006FF8D5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006FF93D | `NoContent_Screen` | Known | Screen layout |
| 0x006FF951 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006FF9B5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006FFA1C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006FFA36 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006FFAA4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006FFB16 | `NoContent_Screen` | Known | Screen layout |
| 0x006FFB2A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006FFB94 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006FFBFD | `No_Photos_Screen` | Known | Screen layout |
| 0x006FFC11 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006FFC77 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006FFCE5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006FFD52 | `NoContent_Screen` | Known | Screen layout |
| 0x006FFD66 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006FFDCE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006FFE38 | `NoContent_Screen` | Known | Screen layout |
| 0x006FFE4C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x006FFEB3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006FFF1D | `NoContent_Screen` | Known | Screen layout |
| 0x006FFF31 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006FFF9E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00700010 | `NoContent_Screen` | Known | Screen layout |
| 0x00700024 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0070008C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007000F5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00700110 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00700176 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00700192 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00700271 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0070028A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007002EB | `FirstBoot_Screen` | Known | Screen layout |
| 0x007002FF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0070046D | `Radio_Screen` | Known | Screen layout |
| 0x0070047D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007004DE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00700561 | `LockediPod_Screen` | Known | Screen layout |
| 0x007005E9 | `Lock_Screen` | Known | Screen layout |
| 0x007005F8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0070065B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007006BD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007006D9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0070074B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0070076A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007007D2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007007EC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00700854 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00700871 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007008DD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00700947 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00700961 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007009D1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00700A44 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00700AB5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00700B24 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00700B90 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00700BAB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00700C20 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00700C87 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00700CE9 | `Photos_Screen` | Known | Screen layout |
| 0x00700D4D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00700D6B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00700DDD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00700DFA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00700E60 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00700E7B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00700EE4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00700F01 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00700F78 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00700F9C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0070100A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00701025 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007010E0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007010FC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070116A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00701187 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007011F2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00701212 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00701289 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007012A5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00701315 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00701334 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007013A0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007013B4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0070142D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007014A1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00701511 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00701579 | `NoContent_Screen` | Known | Screen layout |
| 0x0070158D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007015F1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00701658 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00701672 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007016E0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00701752 | `NoContent_Screen` | Known | Screen layout |
| 0x00701766 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007017D0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00701839 | `No_Photos_Screen` | Known | Screen layout |
| 0x0070184D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007018B3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00701921 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0070198E | `NoContent_Screen` | Known | Screen layout |
| 0x007019A2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00701A0A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00701A74 | `NoContent_Screen` | Known | Screen layout |
| 0x00701A88 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00701AEF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00701B59 | `NoContent_Screen` | Known | Screen layout |
| 0x00701B6D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00701BDA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00701C4C | `NoContent_Screen` | Known | Screen layout |
| 0x00701C60 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00701CC8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00701D31 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00701D4C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00701DB2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00701DCE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00701EAD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00701EC6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00701F27 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00701F3B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007020A9 | `Radio_Screen` | Known | Screen layout |
| 0x007020B9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0070211A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0070219D | `LockediPod_Screen` | Known | Screen layout |
| 0x00702225 | `Lock_Screen` | Known | Screen layout |
| 0x00702234 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00702297 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007022F9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00702315 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00702387 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007023A6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0070240E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00702428 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00702490 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007024AD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00702519 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00702583 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0070259D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0070260D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00702680 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007026F1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00702760 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007027CC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007027E7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0070285C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007028C3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00702925 | `Photos_Screen` | Known | Screen layout |
| 0x00702989 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007029A7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00702A19 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00702A36 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00702A9C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00702AB7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00702B20 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00702B3D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00702BB4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00702BD8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00702C46 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00702C61 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00702D1C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00702D38 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00702DA6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00702DC3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00702E2E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00702E4E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00702EC5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00702EE1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00702F51 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00702F70 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00702FDC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00702FF0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00703069 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007030DD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0070314D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007031B5 | `NoContent_Screen` | Known | Screen layout |
| 0x007031C9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0070322D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00703294 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007032AE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0070331C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0070338E | `NoContent_Screen` | Known | Screen layout |
| 0x007033A2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0070340C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00703475 | `No_Photos_Screen` | Known | Screen layout |
| 0x00703489 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007034EF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0070355D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007035CA | `NoContent_Screen` | Known | Screen layout |
| 0x007035DE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00703646 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007036B0 | `NoContent_Screen` | Known | Screen layout |
| 0x007036C4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0070372B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00703795 | `NoContent_Screen` | Known | Screen layout |
| 0x007037A9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00703816 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00703888 | `NoContent_Screen` | Known | Screen layout |
| 0x0070389C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00703904 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0070396D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00703988 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007039EE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00703A0A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00703AE9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00703B02 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00703B63 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00703B77 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00703CE5 | `Radio_Screen` | Known | Screen layout |
| 0x00703CF5 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00703D56 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00703DD9 | `LockediPod_Screen` | Known | Screen layout |
| 0x00703E61 | `Lock_Screen` | Known | Screen layout |
| 0x00703E70 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00703ED3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00703F35 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00703F51 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00703FC3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00703FE2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0070404A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00704064 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007040CC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007040E9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00704155 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007041BF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007041D9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00704249 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007042BC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0070432D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0070439C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00704408 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00704423 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00704498 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007044FF | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00704561 | `Photos_Screen` | Known | Screen layout |
| 0x007045C5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007045E3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00704655 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00704672 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007046D8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007046F3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0070475C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00704779 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007047F0 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00704814 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00704882 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0070489D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00704958 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00704974 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007049E2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007049FF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00704A6A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00704A8A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00704B01 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00704B1D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00704B8D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00704BAC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00704C18 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00704C2C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00704CA5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00704D19 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00704D89 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00704DF1 | `NoContent_Screen` | Known | Screen layout |
| 0x00704E05 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00704E69 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00704ED0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00704EEA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00704F58 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00704FCA | `NoContent_Screen` | Known | Screen layout |
| 0x00704FDE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00705048 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007050B1 | `No_Photos_Screen` | Known | Screen layout |
| 0x007050C5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0070512B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00705199 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00705206 | `NoContent_Screen` | Known | Screen layout |
| 0x0070521A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00705282 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007052EC | `NoContent_Screen` | Known | Screen layout |
| 0x00705300 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00705367 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007053D1 | `NoContent_Screen` | Known | Screen layout |
| 0x007053E5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00705452 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007054C4 | `NoContent_Screen` | Known | Screen layout |
| 0x007054D8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00705540 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007055A9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007055C4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0070562A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00705646 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00705725 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0070573E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0070579F | `FirstBoot_Screen` | Known | Screen layout |
| 0x007057B3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00705921 | `Radio_Screen` | Known | Screen layout |
| 0x00705931 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00705992 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00705A15 | `LockediPod_Screen` | Known | Screen layout |
| 0x00705A9D | `Lock_Screen` | Known | Screen layout |
| 0x00705AAC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00705B0F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00705B71 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00705B8D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00705BFF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00705C1E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00705C86 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00705CA0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00705D08 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00705D25 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00705D91 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00705DFB | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00705E15 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00705E85 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00705EF8 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00705F69 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00705FD8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00706044 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0070605F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007060D4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0070613B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0070619D | `Photos_Screen` | Known | Screen layout |
| 0x00706201 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0070621F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00706291 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007062AE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00706314 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0070632F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00706398 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007063B5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0070642C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00706450 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007064BE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007064D9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00706594 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007065B0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070661E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0070663B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007066A6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007066C6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0070673D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00706759 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007067C9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007067E8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00706854 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00706868 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007068E1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00706955 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007069C5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00706A2D | `NoContent_Screen` | Known | Screen layout |
| 0x00706A41 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00706AA5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00706B0C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00706B26 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00706B94 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00706C06 | `NoContent_Screen` | Known | Screen layout |
| 0x00706C1A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00706C84 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00706CED | `No_Photos_Screen` | Known | Screen layout |
| 0x00706D01 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00706D67 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00706DD5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00706E42 | `NoContent_Screen` | Known | Screen layout |
| 0x00706E56 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00706EBE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00706F28 | `NoContent_Screen` | Known | Screen layout |
| 0x00706F3C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00706FA3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0070700D | `NoContent_Screen` | Known | Screen layout |
| 0x00707021 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0070708E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00707100 | `NoContent_Screen` | Known | Screen layout |
| 0x00707114 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0070717C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007071E5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00707200 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00707266 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00707282 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00707361 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0070737A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007073DB | `FirstBoot_Screen` | Known | Screen layout |
| 0x007073EF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0070755D | `Radio_Screen` | Known | Screen layout |
| 0x0070756D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007075CE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00707651 | `LockediPod_Screen` | Known | Screen layout |
| 0x007076D9 | `Lock_Screen` | Known | Screen layout |
| 0x007076E8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0070774B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007077AD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007077C9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0070783B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0070785A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007078C2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007078DC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00707944 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00707961 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007079CD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00707A37 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00707A51 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00707AC1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00707B34 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00707BA5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00707C14 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00707C80 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00707C9B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00707D10 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00707D77 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00707DD9 | `Photos_Screen` | Known | Screen layout |
| 0x00707E3D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00707E5B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00707ECD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00707EEA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00707F50 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00707F6B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00707FD4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00707FF1 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00708068 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0070808C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007080FA | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00708115 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007081D0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007081EC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070825A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00708277 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007082E2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00708302 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00708379 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00708395 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00708405 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00708424 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00708490 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007084A4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0070851D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00708591 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00708601 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00708669 | `NoContent_Screen` | Known | Screen layout |
| 0x0070867D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007086E1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00708748 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00708762 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007087D0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00708842 | `NoContent_Screen` | Known | Screen layout |
| 0x00708856 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007088C0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00708929 | `No_Photos_Screen` | Known | Screen layout |
| 0x0070893D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007089A3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00708A11 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00708A7E | `NoContent_Screen` | Known | Screen layout |
| 0x00708A92 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00708AFA | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00708B64 | `NoContent_Screen` | Known | Screen layout |
| 0x00708B78 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00708BDF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00708C49 | `NoContent_Screen` | Known | Screen layout |
| 0x00708C5D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00708CCA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00708D3C | `NoContent_Screen` | Known | Screen layout |
| 0x00708D50 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00708DB8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00708E21 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00708E3C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00708EA2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00708EBE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00708F9D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00708FB6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00709017 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0070902B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00709199 | `Radio_Screen` | Known | Screen layout |
| 0x007091A9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0070920A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0070928D | `LockediPod_Screen` | Known | Screen layout |
| 0x00709315 | `Lock_Screen` | Known | Screen layout |
| 0x00709324 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00709387 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007093E9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00709405 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00709477 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00709496 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007094FE | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00709518 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00709580 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0070959D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00709609 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00709673 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0070968D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007096FD | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00709770 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007097E1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00709850 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007098BC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007098D7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0070994C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007099B3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00709A15 | `Photos_Screen` | Known | Screen layout |
| 0x00709A79 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00709A97 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00709B09 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00709B26 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00709B8C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00709BA7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00709C10 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00709C2D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00709CA4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00709CC8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00709D36 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00709D51 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00709DF1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00709E0D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00709E7B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00709E98 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00709F03 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00709F23 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00709F9A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00709FB6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070A026 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0070A045 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0070A0B1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0070A0C5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0070A13A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0070A1A5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0070A214 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0070A285 | `NoContent_Screen` | Known | Screen layout |
| 0x0070A299 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0070A308 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0070A37B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0070A3E8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0070A451 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0070A4C1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0070A531 | `NoContent_Screen` | Known | Screen layout |
| 0x0070A545 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0070A5A8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0070A60B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0070A627 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0070A6E7 | `Radio_Screen` | Known | Screen layout |
| 0x0070A6F7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0070A758 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0070A7C6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0070A7E5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0070A853 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0070A8B8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0070A8D3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0070A979 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070A995 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070AA03 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0070AA20 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0070AA8B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0070AAAB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0070AB22 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070AB3E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070ABAE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0070ABCD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0070AC39 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0070AC4D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0070ACC2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0070AD2D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0070AD9C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0070AE0D | `NoContent_Screen` | Known | Screen layout |
| 0x0070AE21 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0070AE90 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0070AF03 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0070AF70 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0070AFD9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0070B049 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0070B0B9 | `NoContent_Screen` | Known | Screen layout |
| 0x0070B0CD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0070B130 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0070B193 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0070B1AF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0070B26F | `Radio_Screen` | Known | Screen layout |
| 0x0070B27F | `Radio_Screen_Default` | Known | Screen layout |
| 0x0070B2E0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0070B34E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0070B36D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0070B3DB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0070B440 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0070B45B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0070B501 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070B51D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070B58B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0070B5A8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0070B613 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0070B633 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0070B6AA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070B6C6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070B736 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0070B755 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0070B7C1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0070B7D5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0070B84A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0070B8B5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0070B924 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0070B995 | `NoContent_Screen` | Known | Screen layout |
| 0x0070B9A9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0070BA18 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0070BA8B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0070BAF8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0070BB61 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0070BBD1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0070BC41 | `NoContent_Screen` | Known | Screen layout |
| 0x0070BC55 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0070BCB8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0070BD1B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0070BD37 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0070BDF7 | `Radio_Screen` | Known | Screen layout |
| 0x0070BE07 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0070BE68 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0070BED6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0070BEF5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0070BF63 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0070BFC8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0070BFE3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0070C089 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070C0A5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070C113 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0070C130 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0070C19B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0070C1BB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0070C232 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070C24E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070C2BE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0070C2DD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0070C349 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0070C35D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0070C3D2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0070C43D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0070C4AC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0070C51D | `NoContent_Screen` | Known | Screen layout |
| 0x0070C531 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0070C5A0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0070C613 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0070C680 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0070C6E9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0070C759 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0070C7C9 | `NoContent_Screen` | Known | Screen layout |
| 0x0070C7DD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0070C840 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0070C8A3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0070C8BF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0070C97F | `Radio_Screen` | Known | Screen layout |
| 0x0070C98F | `Radio_Screen_Default` | Known | Screen layout |
| 0x0070C9F0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0070CA5E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0070CA7D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0070CAEB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0070CB50 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0070CB6B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0070CC11 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070CC2D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070CC9B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0070CCB8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0070CD23 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0070CD43 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0070CDBA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070CDD6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070CE46 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0070CE65 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0070CED1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0070CEE5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0070CF5A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0070CFC5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0070D034 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0070D0A5 | `NoContent_Screen` | Known | Screen layout |
| 0x0070D0B9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0070D128 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0070D19B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0070D208 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0070D271 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0070D2E1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0070D351 | `NoContent_Screen` | Known | Screen layout |
| 0x0070D365 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0070D3C8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0070D42B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0070D447 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0070D507 | `Radio_Screen` | Known | Screen layout |
| 0x0070D517 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0070D578 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0070D5E6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0070D605 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0070D673 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0070D6D8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0070D6F3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0070D799 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070D7B5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070D823 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0070D840 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0070D8AB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0070D8CB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0070D942 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070D95E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070D9CE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0070D9ED | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0070DA59 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0070DA6D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0070DAE2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0070DB4D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0070DBBC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0070DC2D | `NoContent_Screen` | Known | Screen layout |
| 0x0070DC41 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0070DCB0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0070DD23 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0070DD90 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0070DDF9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0070DE69 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0070DED9 | `NoContent_Screen` | Known | Screen layout |
| 0x0070DEED | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0070DF50 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0070DFB3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0070DFCF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0070E08F | `Radio_Screen` | Known | Screen layout |
| 0x0070E09F | `Radio_Screen_Default` | Known | Screen layout |
| 0x0070E100 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0070E16E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0070E18D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0070E1FB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0070E260 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0070E27B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0070E321 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070E33D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070E3AB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0070E3C8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0070E433 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0070E453 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0070E4CA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070E4E6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070E556 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0070E575 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0070E5E1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0070E5F5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0070E66A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0070E6D5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0070E744 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0070E7B5 | `NoContent_Screen` | Known | Screen layout |
| 0x0070E7C9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0070E838 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0070E8AB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0070E918 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0070E981 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0070E9F1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0070EA61 | `NoContent_Screen` | Known | Screen layout |
| 0x0070EA75 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0070EAD8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0070EB3B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0070EB57 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0070EC17 | `Radio_Screen` | Known | Screen layout |
| 0x0070EC27 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0070EC88 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0070ECF6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0070ED15 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0070ED83 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0070EDE8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0070EE03 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0070EEA9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070EEC5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070EF33 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0070EF50 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0070EFBB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0070EFDB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0070F052 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070F06E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070F0DE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0070F0FD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0070F169 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0070F17D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0070F1F2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0070F25D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0070F2CC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0070F33D | `NoContent_Screen` | Known | Screen layout |
| 0x0070F351 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0070F3C0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0070F433 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0070F4A0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0070F509 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0070F579 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0070F5E9 | `NoContent_Screen` | Known | Screen layout |
| 0x0070F5FD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0070F660 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0070F6C3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0070F6DF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0070F79F | `Radio_Screen` | Known | Screen layout |
| 0x0070F7AF | `Radio_Screen_Default` | Known | Screen layout |
| 0x0070F810 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0070F87E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0070F89D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0070F90B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0070F970 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0070F98B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0070FA31 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070FA4D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070FABB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0070FAD8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0070FB43 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0070FB63 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0070FBDA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070FBF6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0070FC66 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0070FC85 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0070FCF1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0070FD05 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0070FD7A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0070FDE5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0070FE54 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0070FEC5 | `NoContent_Screen` | Known | Screen layout |
| 0x0070FED9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0070FF48 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0070FFBB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00710028 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00710091 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00710101 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00710171 | `NoContent_Screen` | Known | Screen layout |
| 0x00710185 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007101E8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0071024B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00710267 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00710327 | `Radio_Screen` | Known | Screen layout |
| 0x00710337 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00710398 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00710406 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00710425 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00710493 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007104F8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00710513 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007105B9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007105D5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00710643 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00710660 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007106CB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007106EB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00710762 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0071077E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007107EE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0071080D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00710879 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0071088D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00710902 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0071096D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007109DC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00710A4D | `NoContent_Screen` | Known | Screen layout |
| 0x00710A61 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00710AD0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00710B43 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00710BB0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00710C19 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00710C89 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00710CF9 | `NoContent_Screen` | Known | Screen layout |
| 0x00710D0D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00710D70 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00710DD3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00710DEF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00710EAF | `Radio_Screen` | Known | Screen layout |
| 0x00710EBF | `Radio_Screen_Default` | Known | Screen layout |
| 0x00710F20 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00710F8E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00710FAD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0071101B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00711080 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0071109B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00711141 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0071115D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007111CB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007111E8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00711253 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00711273 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007112EA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00711306 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00711376 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00711395 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00711401 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00711415 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0071148A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007114F5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00711564 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007115D5 | `NoContent_Screen` | Known | Screen layout |
| 0x007115E9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00711658 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007116CB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00711738 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007117A1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00711811 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00711881 | `NoContent_Screen` | Known | Screen layout |
| 0x00711895 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007118F8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0071195B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00711977 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00711A37 | `Radio_Screen` | Known | Screen layout |
| 0x00711A47 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00711AA8 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00711B16 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00711B35 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00711BA3 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00711C08 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00711C23 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00711CC9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00711CE5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00711D53 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00711D70 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00711DDB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00711DFB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00711E72 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00711E8E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00711EFE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00711F1D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00711F89 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00711F9D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00712012 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0071207D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007120EC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0071215D | `NoContent_Screen` | Known | Screen layout |
| 0x00712171 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007121E0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00712253 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007122C0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00712329 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00712399 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00712409 | `NoContent_Screen` | Known | Screen layout |
| 0x0071241D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00712480 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007124E3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007124FF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007125BF | `Radio_Screen` | Known | Screen layout |
| 0x007125CF | `Radio_Screen_Default` | Known | Screen layout |
| 0x00712630 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0071269E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007126BD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0071272B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00712790 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007127AB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0071288C | `MediaLists_AudiobookChapters_Screen)` | Known | Screen layout |
| 0x007128B3 | `MediaLists_AudiobookChapters_Screen_Plane!` | Known | Screen layout |
| 0x0071304D | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x00713068 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x007130D3 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007130EE | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x00713161 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0071317C | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x00713339 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x00713354 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x007133BF | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007133DA | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x0071344D | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x00713468 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x00713630 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0071364C | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x007136C7 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007136E3 | `MediaLists_Albums_Screen_WithTrackCount%` | Known | Screen layout |
| 0x0071375C | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x00713777 | `MediaLists_Songs_Screen_Plain1` | Known | Screen layout |
| 0x007137F2 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0071380D | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x00713A2F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00713A4C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00713B2B | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00713B47 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x00713BC2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00713BDD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00713DC3 | `MediaLists_PodcastEpisodes_Screen'` | Known | Screen layout |
| 0x00713DE8 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007140BA | `MediaLists_TVSeasons_Screen#` | Known | Screen layout |
| 0x007140D9 | `MediaLists_TVSeasons_Screen_Default%` | Known | Screen layout |
| 0x0071414E | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x0071416E | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007142F6 | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x00714316 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0071470F | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x00714734 | `MediaLists_NestedPlaylists_Screen_Default,` | Known | Screen layout |
| 0x007147B6 | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x007147D5 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00714965 | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x0071498A | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x00714A02 | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x00714A21 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00714A85 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x00714B32 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x00714BA4 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x00714C9A | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x00714F5C | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x0071505C | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x007150C8 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00715132 | `NoContent_Screen` | Known | Screen layout |
| 0x00715146 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007151B0 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00715224 | `NoContent_Screen` | Known | Screen layout |
| 0x00715238 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007152A3 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0071530F | `NoContent_Screen` | Known | Screen layout |
| 0x00715323 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0071538A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007153F6 | `NoContent_Screen` | Known | Screen layout |
| 0x0071540A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00715477 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007154EB | `NoContent_Screen` | Known | Screen layout |
| 0x007154FF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00715567 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007155D4 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00715638 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00715654 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007156C0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007156DD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0071574A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00715811 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0071582E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007158A5 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007158C9 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00715980 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007159EA | `NoContent_Screen` | Known | Screen layout |
| 0x007159FE | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00715A68 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00715ADC | `NoContent_Screen` | Known | Screen layout |
| 0x00715AF0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00715B5B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00715BC7 | `NoContent_Screen` | Known | Screen layout |
| 0x00715BDB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00715C42 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00715CAE | `NoContent_Screen` | Known | Screen layout |
| 0x00715CC2 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00715D2F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00715DA3 | `NoContent_Screen` | Known | Screen layout |
| 0x00715DB7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00715E1F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00715E8C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00715EF0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00715F0C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00715F78 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00715F95 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00716002 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007160C9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007160E6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0071615D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00716181 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00716238 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007162A2 | `NoContent_Screen` | Known | Screen layout |
| 0x007162B6 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00716320 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00716394 | `NoContent_Screen` | Known | Screen layout |
| 0x007163A8 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00716413 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0071647F | `NoContent_Screen` | Known | Screen layout |
| 0x00716493 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007164FA | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00716566 | `NoContent_Screen` | Known | Screen layout |
| 0x0071657A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007165E7 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0071665B | `NoContent_Screen` | Known | Screen layout |
| 0x0071666F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007166D7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00716744 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007167A8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007167C4 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00716830 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0071684D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007168BA | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00716981 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0071699E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00716A15 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00716A39 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00716AF0 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00716B5A | `NoContent_Screen` | Known | Screen layout |
| 0x00716B6E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00716BD8 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00716C4C | `NoContent_Screen` | Known | Screen layout |
| 0x00716C60 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00716CCB | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00716D37 | `NoContent_Screen` | Known | Screen layout |
| 0x00716D4B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00716DB2 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00716E1E | `NoContent_Screen` | Known | Screen layout |
| 0x00716E32 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00716E9F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00716F13 | `NoContent_Screen` | Known | Screen layout |
| 0x00716F27 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00716F8F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00716FFC | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00717060 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0071707C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007170E8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00717105 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00717172 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00717239 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00717256 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007172CD | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007172F1 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007173A8 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00717412 | `NoContent_Screen` | Known | Screen layout |
| 0x00717426 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00717490 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00717504 | `NoContent_Screen` | Known | Screen layout |
| 0x00717518 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00717583 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007175EF | `NoContent_Screen` | Known | Screen layout |
| 0x00717603 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0071766A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007176D6 | `NoContent_Screen` | Known | Screen layout |
| 0x007176EA | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00717757 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007177CB | `NoContent_Screen` | Known | Screen layout |
| 0x007177DF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00717847 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007178B4 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00717918 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00717934 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007179A0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007179BD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00717A2A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00717AF1 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00717B0E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00717B85 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00717BA9 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00717C60 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00717CCA | `NoContent_Screen` | Known | Screen layout |
| 0x00717CDE | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00717D48 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00717DBC | `NoContent_Screen` | Known | Screen layout |
| 0x00717DD0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00717E3B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00717EA7 | `NoContent_Screen` | Known | Screen layout |
| 0x00717EBB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00717F22 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00717F8E | `NoContent_Screen` | Known | Screen layout |
| 0x00717FA2 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0071800F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00718083 | `NoContent_Screen` | Known | Screen layout |
| 0x00718097 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007180FF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0071816C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007181D0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007181EC | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00718258 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00718275 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007182E2 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007183A9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007183C6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0071843D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00718461 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00718518 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00718582 | `NoContent_Screen` | Known | Screen layout |
| 0x00718596 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00718600 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00718674 | `NoContent_Screen` | Known | Screen layout |
| 0x00718688 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007186F3 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0071875F | `NoContent_Screen` | Known | Screen layout |
| 0x00718773 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007187DA | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00718846 | `NoContent_Screen` | Known | Screen layout |
| 0x0071885A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007188C7 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0071893B | `NoContent_Screen` | Known | Screen layout |
| 0x0071894F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007189B7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00718A24 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00718A88 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00718AA4 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00718B10 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00718B2D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00718B9A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00718C61 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00718C7E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00718CF5 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00718D19 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00718DD0 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00718E3A | `NoContent_Screen` | Known | Screen layout |
| 0x00718E4E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00718EB8 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00718F2C | `NoContent_Screen` | Known | Screen layout |
| 0x00718F40 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00718FAB | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00719017 | `NoContent_Screen` | Known | Screen layout |
| 0x0071902B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00719092 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007190FE | `NoContent_Screen` | Known | Screen layout |
| 0x00719112 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0071917F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007191F3 | `NoContent_Screen` | Known | Screen layout |
| 0x00719207 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0071926F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007192DC | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00719340 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0071935C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007193C8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007193E5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00719452 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00719519 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00719536 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007195AD | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007195D1 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00719A34 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00719AA6 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00719B11 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00719B76 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00719BE0 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00719C4A | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00719CBA | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x00719D31 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x00719D9F | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x00719E0A | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x00719E74 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x00719EDB | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x00719F4A | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x00719FB8 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0071A01D | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0071A085 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0071A0F0 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0071A15B | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0071A1C2 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0071A530 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0071A5A2 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0071A60D | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0071A672 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0071A6DC | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0071A746 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0071A7B6 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0071A82D | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0071A89B | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0071A906 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0071A970 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0071A9D7 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0071AA46 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0071AAB4 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0071AB19 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0071AB81 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0071ABEC | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0071AC57 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0071ACBE | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0071B02A | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0071B09C | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0071B107 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0071B16C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0071B1D6 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0071B240 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0071B2B0 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0071B327 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0071B395 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0071B400 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0071B46A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0071B4D1 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0071B540 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0071B5AE | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0071B613 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0071B67B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0071B6E6 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0071B751 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0071B7B8 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0071BB22 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0071BB94 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0071BBFF | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0071BC64 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0071BCCE | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0071BD38 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0071BDA8 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0071BE1F | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0071BE8D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0071BEF8 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0071BF62 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0071BFC9 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0071C038 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0071C0A6 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0071C10B | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0071C173 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0071C1DE | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0071C249 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0071C2B0 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0071C602 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0071C674 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0071C6DF | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0071C744 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0071C7AE | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0071C818 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0071C888 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0071C8FF | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0071C96D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0071C9D8 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0071CA42 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0071CAA9 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0071CB18 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0071CB86 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0071CBEB | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0071CC53 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0071CCBE | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0071CD29 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0071CD90 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0071D107 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0071D179 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0071D1E4 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0071D249 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0071D2B3 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0071D31D | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0071D38D | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0071D404 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0071D472 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0071D4DD | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0071D547 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0071D5AE | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0071D61D | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0071D68B | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0071D6F0 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0071D758 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0071D7C3 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0071D82E | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0071D895 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0071DC09 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0071DC7B | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0071DCE6 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0071DD4B | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0071DDB5 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0071DE1F | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0071DE8F | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0071DF06 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0071DF74 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0071DFDF | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0071E049 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0071E0B0 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0071E11F | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0071E18D | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0071E1F2 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0071E25A | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0071E2C5 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0071E330 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0071E397 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0071E6F1 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0071E763 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0071E7CE | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0071E833 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0071E89D | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0071E907 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0071E977 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0071E9EE | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0071EA5C | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0071EAC7 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0071EB31 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0071EB98 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0071EC07 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0071EC75 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0071ECDA | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0071ED42 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0071EDAD | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0071EE18 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0071EE7F | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0071F1D9 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0071F24B | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0071F2B6 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0071F31B | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0071F385 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0071F3EF | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0071F45F | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0071F4D6 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0071F544 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0071F5AF | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0071F619 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0071F680 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0071F6EF | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0071F75D | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0071F7C2 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0071F82A | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0071F895 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0071F900 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0071F967 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0071FCC2 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0071FD34 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0071FD9F | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0071FE04 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0071FE6E | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0071FED8 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0071FF48 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0071FFBF | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0072002D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x00720098 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x00720102 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x00720169 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007201D8 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x00720246 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007202AB | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x00720313 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0072037E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007203E9 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x00720450 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007207D4 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00720846 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007208B1 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00720916 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00720980 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007209EA | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00720A5A | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x00720AD1 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x00720B3F | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x00720BAA | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x00720C14 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x00720C7B | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x00720CEA | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x00720D58 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x00720DBD | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x00720E25 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x00720E90 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x00720EFB | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x00720F62 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007212F0 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00721362 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007213CD | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00721432 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0072149C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00721506 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00721576 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007215ED | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0072165B | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007216C6 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x00721730 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x00721797 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x00721806 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x00721874 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007218D9 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x00721941 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007219AC | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x00721A17 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x00721A7E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x00721DEC | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00721E5E | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00721EC9 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00721F2E | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00721F98 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00722002 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00722072 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007220E9 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x00722157 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007221C2 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0072222C | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x00722293 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x00722302 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x00722370 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007223D5 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0072243D | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007224A8 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x00722513 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0072257A | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007228E0 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00722952 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007229BD | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00722A22 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00722A8C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00722AF6 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00722B66 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x00722BDD | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x00722C4B | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x00722CB6 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x00722D20 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x00722D87 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x00722DF6 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x00722E64 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x00722EC9 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x00722F31 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x00722F9C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x00723007 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0072306E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007233C2 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00723434 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0072349F | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00723504 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0072356E | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007235D8 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00723648 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007236BF | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0072372D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x00723798 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x00723802 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x00723869 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007238D8 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x00723946 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007239AB | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x00723A13 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x00723A7E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x00723AE9 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x00723B50 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x00723E9B | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00723F0D | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00723F78 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00723FDD | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00724047 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007240B1 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00724121 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x00724198 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x00724206 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x00724271 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007242DB | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x00724342 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007243B1 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0072441F | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x00724484 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007244EC | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x00724557 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007245C2 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x00724629 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0072498B | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007249FD | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00724A68 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00724ACD | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00724B37 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00724BA1 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00724C11 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x00724C88 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x00724CF6 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x00724D61 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x00724DCB | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x00724E32 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x00724EA1 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x00724F0F | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x00724F74 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x00724FDC | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x00725047 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007250B2 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x00725119 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x00725431 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007254A3 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0072550E | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00725573 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007255DD | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00725647 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007256B7 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0072572E | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0072579C | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x00725807 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x00725871 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007258D8 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x00725947 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007259B5 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x00725A1A | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x00725A82 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x00725AED | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x00725B58 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x00725BBF | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x00725ED6 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x00725F4D | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x00725FCA | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x0072603C | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007260AC | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x00726122 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x00726190 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007261FD | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x00726542 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007265B9 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x00726636 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007266A8 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x00726718 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x0072678E | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007267FC | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x00726869 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x00726BD2 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x00726C49 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x00726CC6 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x00726D38 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x00726DA8 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x00726E1E | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x00726E8C | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x00726EF9 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x00727262 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007272D9 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x00727354 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007273C4 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x0072743A | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007274A8 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x00727515 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x0072784E | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007278C5 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x00727940 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007279B0 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x00727A26 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x00727A94 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x00727B01 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x00727E38 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x00727EAF | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x00727F2A | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x00727F9A | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x00728010 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x0072807E | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007280EB | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007283FB | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x00728472 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007284ED | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x0072855D | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007285D3 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x00728641 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007286AE | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x00728CB2 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x00728CCF | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x00728D4A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x00728D63 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x00728DDB | `Notes_Contents_Screen` | Known | Screen layout |
| 0x00728DF4 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x00728E69 | `Notes_Image_Screen` | Known | Screen layout |
| 0x00728E7F | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x00728EF6 | `Notes_Image_Screen` | Known | Screen layout |
| 0x00728F0C | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x00728F83 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x00728FA0 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x00729018 | `Notes_List_Screen` | Known | Screen layout |
| 0x0072902D | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007291DE | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007291FB | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x00729276 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x0072928F | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x00729307 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x00729320 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x00729395 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007293AB | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x00729422 | `Notes_Image_Screen` | Known | Screen layout |
| 0x00729438 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007294AF | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007294CC | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x00729544 | `Notes_List_Screen` | Known | Screen layout |
| 0x00729559 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x0072973A | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x00729757 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007297D2 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007297EB | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x00729863 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x0072987C | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007298F1 | `Notes_Image_Screen` | Known | Screen layout |
| 0x00729907 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x0072997E | `Notes_Image_Screen` | Known | Screen layout |
| 0x00729994 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x00729A0B | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x00729A28 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x00729AA0 | `Notes_List_Screen` | Known | Screen layout |
| 0x00729AB5 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x00729C6A | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x00729C87 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x00729D02 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x00729D1B | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x00729D93 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x00729DAC | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x00729E21 | `Notes_Image_Screen` | Known | Screen layout |
| 0x00729E37 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x00729EAE | `Notes_Image_Screen` | Known | Screen layout |
| 0x00729EC4 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x00729F3B | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x00729F58 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x00729FD0 | `Notes_List_Screen` | Known | Screen layout |
| 0x00729FE5 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x0072A2FD | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x0072A3A3 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0072A426 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x0072A4DE | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x0072A560 | `PhotosSettingsSlideshowMusic_Screen+` | Known | Screen layout |
| 0x0072A587 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x0072A66D | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x0072A825 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0072A885 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0072A8E2 | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x0072A909 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x0072A9A9 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0072AA09 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0072AA66 | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x0072AA8D | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x0072AD28 | `Photos_Screen` | Known | Screen layout |
| 0x0072AE74 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0072AED8 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x0072AF39 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x0072AF96 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x0072AFF3 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x0072B061 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x0072B0BE | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x0072B234 | `Photos_Screen` | Known | Screen layout |
| 0x0072B380 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0072B3E4 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x0072B445 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x0072B4A2 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x0072B4FF | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x0072B56D | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x0072B5CA | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x0072B740 | `Photos_Screen` | Known | Screen layout |
| 0x0072B88C | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0072B8F0 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x0072B951 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x0072B9AE | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x0072BA0B | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x0072BA79 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x0072BAD6 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x0072BC4C | `Photos_Screen` | Known | Screen layout |
| 0x0072BD98 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0072BDFC | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x0072BE5D | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x0072BEBA | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x0072BF17 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x0072BF85 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x0072BFE2 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x0072C158 | `Photos_Screen` | Known | Screen layout |
| 0x0072C2A4 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0072C308 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x0072C369 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x0072C3C6 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x0072C423 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x0072C491 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x0072C4EE | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x0072C664 | `Photos_Screen` | Known | Screen layout |
| 0x0072C7B0 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0072C814 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x0072C875 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x0072C8D2 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x0072C92F | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x0072C99D | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x0072C9FA | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x0072CB70 | `Photos_Screen` | Known | Screen layout |
| 0x0072CCBC | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0072CD22 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x0072CD84 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x0072CDE6 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x0072CE7C | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x0072CF9D | `Photos_Screen` | Known | Screen layout |
| 0x0072D008 | `Photos_Screen` | Known | Screen layout |
| 0x0072D154 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0072D1BA | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x0072D21C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x0072D27E | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x0072D314 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x0072D435 | `Photos_Screen` | Known | Screen layout |
| 0x0072D4A0 | `Photos_Screen` | Known | Screen layout |
| 0x0072D5EC | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0072D652 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x0072D6B4 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x0072D716 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x0072D7AC | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x0072D8CD | `Photos_Screen` | Known | Screen layout |
| 0x0072D938 | `Photos_Screen` | Known | Screen layout |
| 0x0072DA84 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0072DAEA | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x0072DB4C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x0072DBAE | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x0072DC44 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x0072DD65 | `Photos_Screen` | Known | Screen layout |
| 0x0072DDD0 | `Photos_Screen` | Known | Screen layout |
| 0x0072DF1C | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0072DF82 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x0072DFE4 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x0072E046 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x0072E0DC | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x0072E1FD | `Photos_Screen` | Known | Screen layout |
| 0x0072E3F1 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x0072E453 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x0072E4C1 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x0072E527 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x0072E58C | `Radio_Screen_Volume` | Known | Screen layout |
| 0x0072E85A | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x0072E8BC | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x0072E92A | `Radio_Screen_Default#` | Known | Screen layout |
| 0x0072E990 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x0072EC96 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x0072ECF8 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x0072ED66 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x0072EDCC | `Radio_Screen_Volume` | Known | Screen layout |
| 0x0072F075 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0072F0D2 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x0072F134 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x0072F1A2 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x0072F208 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x0072F502 | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x0072F56C | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x0072F7DA | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x0072F844 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x0072FA01 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x0072FA64 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x0072FAC9 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x0072FB31 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x0072FB94 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0072FBFC | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x0072FC65 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0072FCCB | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x0072FD30 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0072FD9D | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x0072FE0D | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x0072FE83 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x0072FEF9 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x0072FF69 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0072FFDE | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x00730055 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007300C9 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x0073013B | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007301B5 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00730228 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x0073029A | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0073031E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00730348 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007303CF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0073045C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007304FB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00730515 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x0073058D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007305A7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00730611 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x0073062E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007306A6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007306D0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x00730757 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007307E4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00730883 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0073089D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00730915 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0073092F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00730999 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007309B6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00730A2E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00730A58 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x00730ADF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00730B6C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00730C0B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00730C25 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00730C9D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00730CB7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00730D21 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x00730D3E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00730DB6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00730DE0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x00730E67 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00730EF4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00730F93 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00730FAD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00731025 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0073103F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007310A9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007310C6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0073113E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00731168 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007311EF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0073127C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0073131B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00731335 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007313AD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007313C7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00731431 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x0073144E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007314C6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007314F0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x00731577 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00731604 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007316A3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007316BD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00731735 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0073174F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007317B9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007317D6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0073184E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00731878 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007318FF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0073198C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00731A2B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00731A45 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00731ABD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00731AD7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00731B41 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x00731B5E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00731BD6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00731C00 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x00731C87 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00731D14 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00731DB3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00731DCD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00731E45 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00731E5F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00731EC9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x00731EE6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00731F5E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00731F88 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x0073200F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0073209C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0073213B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00732155 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007321CD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007321E7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00732251 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x0073226E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007322E6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00732310 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x00732397 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00732424 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007324C3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007324DD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00732555 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0073256F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007325D9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007325F6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0073266E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00732698 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x0073271F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007327AC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0073284B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00732865 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007328DD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007328F7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00732961 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x0073297E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007329F6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00732A20 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x00732AA7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00732B34 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00732BD3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00732BED | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00732C65 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00732C7F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00732CE9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x00732D06 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00732D7E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00732DA8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x00732E2F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00732EBC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00732F5B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00732F75 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00732FED | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00733007 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00733071 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x0073308E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00733106 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00733130 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007331B7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00733244 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007332E3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007332FD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00733375 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0073338F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007333F9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x00733416 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0073348E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007334B8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x0073353F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007335CC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0073366B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00733685 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007336FD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00733717 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00733781 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x0073379E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00733816 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00733840 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007338C7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00733954 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007339F3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00733A0D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00733A85 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00733A9F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00733B09 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x00733B26 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00733B9E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00733BC8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x00733C4F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00733CDC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00733D7B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00733D95 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00733E0D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00733E27 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00733E91 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x00733EAE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00733F26 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x00733F50 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x00733FD7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00734064 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00734103 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0073411D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00734195 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007341AF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00734219 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x00734236 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007342AE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007342D8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x0073435F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007343EC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0073448B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007344A5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x0073451D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00734537 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007345A1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007345BE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00734645 | `SettingsMenus_ResetMainMenu_Screen9` | Known | Screen layout |
| 0x00734715 | `SettingsMenus_ResetMusicMenu_Screen:` | Known | Screen layout |
| 0x007347C9 | `SettingsMenus_Main_Screen,` | Known | Screen layout |
| 0x0073483B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00734855 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007348CD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007348E7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00734C22 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x00734C88 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x00734CE5 | `Extras_Screen` | Known | Screen layout |
| 0x00734D39 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x00734E17 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x00734E85 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x00734F23 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x00734F3C | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x00734FA4 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x00735016 | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x0073502F | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x00735092 | `DemoMode_Screen` | Known | Screen layout |
| 0x007350A5 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x00735112 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x0073512B | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x0073519E | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x007351B9 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x007352C9 | `VoiceMemos_Context_PlayDelete_Screen)` | Known | Screen layout |
| 0x007352F1 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x00735368 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x00735434 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007354A3 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x00735591 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007355FA | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x0073561C | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x00735688 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007356AA | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x00735826 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00735842 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x00735909 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x00735924 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x00735987 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007359EA | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x00735A81 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00735A9D | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x00735B64 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x00735B7F | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x00735BE2 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x00735C45 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x00735CDD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00735CF9 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x00735DC0 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x00735DDB | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x00735E3E | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x00735EA1 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x00735F1E | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x00735F89 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x00735FF5 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x00736067 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007360D4 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x0073613F | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x007361AB | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x00736213 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x0073627F | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007362F3 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x00736361 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x007363DA | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x00752208 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x0075228D | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x00752572 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x008F1E69 | `VoiceMemos_Screen_InsufficientDiskSpace2` | Known | Screen layout |
| 0x008F36ED | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x008F3705 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x008F3723 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x008F382F | `NowPlaying_Screen_Music` | Known | Screen layout |
| 0x008F385B | `MainMenu_Music_Screen_NoMusic` | Known | Screen layout |
| 0x008F3879 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x008F3897 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x008F3998 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x008F3A4C | `DeleteRental_Confirmation_Screen_Unwatched` | Known | Screen layout |
| 0x008F3AA2 | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x008F3AEE | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x008F3BF0 | `NowPlaying_Screen_ExtraInfoLoadFailed` | Known | Screen layout |
| 0x008F3C4B | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x008F3C64 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x008F3C82 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x008F3CB1 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x008F3CE9 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x008F4120 | `Video_Settings_TV_Screen_Standard` | Known | Screen layout |
| 0x008F4152 | `Search_Main_Screen_WithKeyboard` | Known | Screen layout |
| 0x008F4172 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x008F41B7 | `VoiceMemos_Screen_InsufficientDiskSpace` | Known | Screen layout |
| 0x008F427B | `Video_Settings_TV_Screen_Wide` | Known | Screen layout |
| 0x008F42C3 | `CoverFlow_Screen_Backside` | Known | Screen layout |
| 0x008F6C40 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x008F6E45 | `NowPlaying_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x008F6E6A | `Slideshow_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x008F6F3A | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x008F6F54 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x008F6FE7 | `RentalDeleted_Screen_Title` | Known | Screen layout |
| 0x008F7002 | `SingleRentalExpiring_Screen_Title` | Known | Screen layout |
| 0x008F7024 | `MultipleRentalsExpiring_Screen_Title` | Known | Screen layout |
| 0x008F7049 | `DeleteRental_Screen_Title` | Known | Screen layout |
| 0x008F70EC | `VoiceMemos_DeleteAll_ScreenTitle` | Known | Screen layout |
| 0x008F7189 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x008F71CC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x008F73BD | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x008F74A6 | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x008F74BF | `Radio_Screen_Volume` | Known | Screen layout |
| 0x008F74D3 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x008F74F0 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x008F750F | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x008F75DB | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x008F7731 | `CoverFlow_Screen_BacksideForCapture` | Known | Screen layout |
| 0x008F86A6 | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x008F86C1 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x008F89B8 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x008F89EC | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x008F8A29 | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x008F8B3B | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x008F8C8B | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x008F8CC3 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x008F8CE9 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x008FE8EF | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x008FE91A | `TVOutSlideshow_Screen_Playing` | Known | Screen layout |
| 0x008FE938 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x008FE972 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x008FEA0F | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x008FEA7A | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x008FEAFA | `Extras_Screen_Debug` | Known | Screen layout |
| 0x008FEC04 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x008FEC24 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x008FF16F | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x008FF18A | `Extras_Screen_Lock` | Known | Screen layout |
| 0x008FF19D | `MainMenu_List_ScreenLock` | Known | Screen layout |
| 0x008FF1B6 | `ExtrasMenu_ScreenLock` | Known | Screen layout |
| 0x008FF229 | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x008FF24A | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x008FF31D | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x008FF33F | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x008FF446 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x008FF486 | `VoiceMemos_Screen_DeletAllAsk` | Known | Screen layout |
| 0x008FF4A4 | `NowPlaying_Screen_TVOutAsk` | Known | Screen layout |
| 0x008FF600 | `NowPlaying_Screen_Initial` | Known | Screen layout |
| 0x008FF61A | `NowPlaying_Screen_Video_TVOut_Initial` | Known | Screen layout |
| 0x0090036E | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x009003EF | `RemoteUI_Screen` | Known | Screen layout |
| 0x009003FF | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x00900417 | `Video_Settings_TV_Screen` | Known | Screen layout |
| 0x00900430 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00900447 | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x0090046B | `AddressViewer_PartialLoad_Screen` | Known | Screen layout |
| 0x0090048C | `Alarms_Alarm_Clock_Triggered_Screen` | Known | Screen layout |
| 0x009004B0 | `Alarms_Alarm_Triggered_Screen` | Known | Screen layout |
| 0x009004CE | `Unsupported_Screen` | Known | Screen layout |
| 0x009004E1 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x009004FF | `LockediPod_Screen` | Known | Screen layout |
| 0x00900511 | `DiskMode_Screen` | Known | Screen layout |
| 0x00900521 | `DemoMode_Screen` | Known | Screen layout |
| 0x00900531 | `Notes_Image_Screen` | Known | Screen layout |
| 0x00900544 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x00900562 | `NowPlaying_Idle_Screen` | Known | Screen layout |
| 0x00900579 | `Game_Screen` | Known | Screen layout |
| 0x00900585 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x009005A2 | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x009005BB | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x009005DC | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x00900601 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x00900614 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x00900631 | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x00900652 | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x00900677 | `Notes_Loading_Screen` | Known | Screen layout |
| 0x0090068C | `SettingsMenus_AdjustScrolling_Screen` | Known | Screen layout |
| 0x009006B1 | `Game_Running_Screen` | Known | Screen layout |
| 0x009006C5 | `Stopwatch_Screen` | Known | Screen layout |
| 0x009006D6 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x009006ED | `Clock_Screen` | Known | Screen layout |
| 0x009006FA | `TVOutSlideshowAsk_Screen` | Known | Screen layout |
| 0x00900713 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x00900729 | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x00900747 | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x00900763 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x00900774 | `DemoMode_Main_Screen` | Known | Screen layout |
| 0x00900789 | `Search_Main_Screen` | Known | Screen layout |
| 0x0090079C | `AddressViewer_Main_Screen` | Known | Screen layout |
| 0x009007B6 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x009007CB | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x009007E1 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x009007FB | `Clock_Region_Screen` | Known | Screen layout |
| 0x0090080F | `RentalDeleted_Notification_Screen` | Known | Screen layout |
| 0x00900831 | `SingleRentalExpiring_Notification_Screen` | Known | Screen layout |
| 0x0090085A | `MultipleRentalsExpiring_Notification_Screen` | Known | Screen layout |
| 0x00900886 | `RentalError_Notification_Screen` | Known | Screen layout |
| 0x009008A6 | `DeleteRental_Confirmation_Screen` | Known | Screen layout |
| 0x009008C7 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x009008DF | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x009008FD | `Video_Settings_Fit_To_Screen` | Known | Screen layout |
| 0x0090091A | `RentalInfo_Screen` | Known | Screen layout |
| 0x0090092C | `Radio_Screen` | Known | Screen layout |
| 0x00900939 | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x00900953 | `ChargingMode_LowPower_Screen` | Known | Screen layout |
| 0x00900970 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0090098A | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x009009A4 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x009009BE | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x009009D7 | `Extras_Screen` | Known | Screen layout |
| 0x009009E5 | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x00900A02 | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x00900A24 | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x00900A3D | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x00900A5B | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x00900A74 | `Video_Settings_Screen` | Known | Screen layout |
| 0x00900A8A | `SettingsMenus_AudiobookSettings_Screen` | Known | Screen layout |
| 0x00900AB1 | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x00900AD7 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x00900AED | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x00900B05 | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x00900B28 | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x00900B45 | `MediaLists_Rentals_Screen` | Known | Screen layout |
| 0x00900B5F | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x00900B83 | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x00900B9C | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x00900BBE | `ToDo_Instructions_Screen` | Known | Screen layout |
| 0x00900BD7 | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x00900BF3 | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x00900C0D | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x00900C2E | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x00900C4A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00900C62 | `VoiceMemos_Screen` | Known | Screen layout |
| 0x00900C74 | `No_Photos_Screen` | Known | Screen layout |
| 0x00900C85 | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x00900C9F | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x00900CBB | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x00900CDF | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x00900CFF | `ChargingMode_Progress_Screen` | Known | Screen layout |
| 0x00900D1C | `Notes_Contents_Screen` | Known | Screen layout |
| 0x00900D32 | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x00900D4D | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x00900D69 | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x00900D8B | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x00900DAC | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x00900DC6 | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x00900DE0 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00900DFF | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x00900E20 | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x00900E38 | `NoContent_Screen` | Known | Screen layout |
| 0x00900E49 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00900E5F | `FirstBoot_Screen` | Known | Screen layout |
| 0x00900E70 | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x00900E86 | `Notes_List_Screen` | Known | Screen layout |
| 0x00900E98 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x00900EAE | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x00900ECF | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x00900EE9 | `TODOS_Menu_Screen` | Known | Screen layout |
| 0x00900EFB | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x00900F11 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x00900F2D | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x00900F42 | `Games_Menu_Screen` | Known | Screen layout |
| 0x00900F54 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x00900F67 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x00900F86 | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x00900FA5 | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x00900FC9 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x00900FDF | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x00900FFD | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x00901020 | `TVOutSlideshow_Screen` | Known | Screen layout |
| 0x00901036 | `CoverFlow_Screen` | Known | Screen layout |
| 0x00901047 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0090105B | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x0090107D | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x00901095 | `Stopwatch_SessionSummary_Screen` | Known | Screen layout |
| 0x009010B5 | `SettingsMenu_AboutWithAccessory_Screen` | Known | Screen layout |
| 0x009010DC | `Clock_NorthAmerica_City_Screen` | Known | Screen layout |
| 0x009010FB | `Clock_SouthAmerica_City_Screen` | Known | Screen layout |
| 0x0090111A | `Clock_Africa_City_Screen` | Known | Screen layout |
| 0x00901133 | `Clock_Australia_City_Screen` | Known | Screen layout |
| 0x0090114F | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x00901166 | `Clock_Pacific_City_Screen` | Known | Screen layout |
| 0x00901180 | `Clock_Atlantic_City_Screen` | Known | Screen layout |
| 0x0090119B | `Clock_Europe_City_Screen` | Known | Screen layout |
| 0x0090127B | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x009012CC | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x009012EF | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00901317 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x00901647 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0090174A | `RentalInfo_Screen_ExpiringSoon` | Known | Screen layout |
| 0x009017A0 | `RentalInfo_Screen_NoAlbumArt_ExpiringSoon` | Known | Screen layout |
| 0x00901B6F | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x00901BC5 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x00901D16 | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x00901D33 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x00902107 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x00902229 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0090224B | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x009022B8 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x009022D7 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x009028FE | `ChargingMode_ScreenLayout_LowPower` | Known | Screen layout |
| 0x0090324E | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x00903393 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x0090346F | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0090348D | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x009034AD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x009035B8 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x009035D4 | `Extras_Screen_Games` | Known | Screen layout |
| 0x009036DA | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x009036F9 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x00903715 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x009037E0 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x009038BB | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00903A89 | `MainMenu_Music_Screen_NoAudiobooks` | Known | Screen layout |
| 0x00903AAC | `MainMenus_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x00903ACF | `NoContent_Screen_NoAudiobooks` | Known | Screen layout |
| 0x00903B09 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00903B28 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00903B49 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00903BF8 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x00903C15 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x00903C94 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00903D78 | `MainMenu_Music_Screen_NoCompilations` | Known | Screen layout |
| 0x00903D9D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00903F24 | `MainMenu_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00903F47 | `MainMenu_Videos_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00903F6C | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00903F8B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00903FAA | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00903FCB | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x00904009 | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x0090402A | `Extras_Screen_VoiceMemos` | Known | Screen layout |
| 0x00904095 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x009040C7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x009040E6 | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x00904193 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x009041FF | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x009042F8 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x00904314 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x00904397 | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x009043B2 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x009043D3 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x00904482 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x009044B6 | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x009044D7 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0090457A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0090459B | `MainMenu_Videos_Screen_NoPlaylists` | Known | Screen layout |
| 0x009045BE | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0090460D | `MainMenu_Music_Screen_NoArtists` | Known | Screen layout |
| 0x009046B4 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x009046D3 | `Extras_Screen_NikePlus` | Known | Screen layout |
| 0x00904823 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00904842 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00904863 | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x00904CCE | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x00904D81 | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x00904DFB | `Notes_Contents_Screen_Alt` | Known | Screen layout |
| 0x00904E15 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x00904EC1 | `VoiceMemos_Screen_InsufficientDiskSpace2_Default` | Known | Screen layout |
| 0x00904F73 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x00905018 | `VoiceMemos_Screen_InsufficientDiskSpace_Default` | Known | Screen layout |
| 0x00905048 | `NowPlaying_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x00905075 | `Slideshow_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x00905C54 | `Radio_Screen_Tuning_Default` | Known | Screen layout |
| 0x00905CB5 | `VoiceMemos_Screen_DeletAllAsk_Default` | Known | Screen layout |
| 0x00905CDB | `NowPlaying_Screen_TVOutAsk_Default` | Known | Screen layout |
| 0x00905CFE | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00905D1C | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x00905D48 | `AddressViewer_PartialLoad_Screen_Default` | Known | Screen layout |
| 0x00905D71 | `Alarms_Alarm_Clock_Triggered_Screen_Default` | Known | Screen layout |
| 0x00905D9D | `Alarms_Alarm_Triggered_Screen_Default` | Known | Screen layout |
| 0x00905DC3 | `Unsupported_Screen_Default` | Known | Screen layout |
| 0x00905DDE | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x00905E04 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x00905E1C | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x00905E37 | `Game_Screen_Default` | Known | Screen layout |
| 0x00905E4B | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x00905E71 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x00905E92 | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x00905EBB | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x00905EE5 | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x00905F12 | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x00905F3B | `Notes_Loading_Screen_Default` | Known | Screen layout |
| 0x00905F58 | `Clock_Screen_Default` | Known | Screen layout |
| 0x00905F6D | `TVOutSlideshowAsk_Screen_Default` | Known | Screen layout |
| 0x00905F8E | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x00905FAC | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x00905FD2 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x00905FF6 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x0090600F | `AddressViewer_Main_Screen_Default` | Known | Screen layout |
| 0x00906031 | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x0090604E | `MainMenus_Main_Screen_Default` | Known | Screen layout |
| 0x0090606C | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00906089 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x009060A5 | `RentalDeleted_Notification_Screen_Default` | Known | Screen layout |
| 0x009060CF | `SingleRentalExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x00906100 | `MultipleRentalsExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x00906134 | `RentalError_Notification_Screen_Default` | Known | Screen layout |
| 0x0090615C | `DeleteRental_Confirmation_Screen_Default` | Known | Screen layout |
| 0x00906185 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x009061B1 | `RentalInfo_Screen_Default` | Known | Screen layout |
| 0x009061CB | `Radio_Screen_Default` | Known | Screen layout |
| 0x009061E0 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x00906202 | `Extras_Screen_Default` | Known | Screen layout |
| 0x00906218 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x0090623E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0090625F | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x0090627D | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0090629F | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x009062CB | `ToDo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009062EC | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00906310 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x00906332 | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x00906356 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00906375 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0090638E | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x009063B0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x009063D4 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x009063F2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00906416 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x00906440 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00906469 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0090648B | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x009064AB | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x009064C9 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x009064E2 | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x00906500 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x0090651A | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x00906538 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x00906561 | `TODOS_Menu_Screen_Default` | Known | Screen layout |
| 0x0090657B | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x00906599 | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x009065B6 | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x009065D0 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x009065EB | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x0090660A | `Debug_MainMenu_Screen_Default` | Known | Screen layout |
| 0x00906628 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x00906646 | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x0090665F | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x0090667B | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x009066A5 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x009066C5 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x009066ED | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x00906714 | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0090673B | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x0090675C | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x00906780 | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x0090679F | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x009067C1 | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x009067E4 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x00906805 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00906893 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x009068C3 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x009068E5 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x00906956 | `RentalInfo_Screen_NoAlbumArt_Default` | Known | Screen layout |
| 0x0090697B | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x00906F56 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00906F82 | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00906FC7 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x00906FEF | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00907010 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00907031 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00907057 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00907074 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00907096 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x009070BA | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x009070DE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0090729A | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0090730A | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0090735B | `NowPlaying_Screen_Video_TVOut` | Known | Screen layout |
| 0x009074CD | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x009074F4 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x00907A2D | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x00907BEA | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x00907DDC | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x009080A8 | `Lock_Screen_Default_Layout` | Known | Screen layout |
| 0x0090813E | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x00908165 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x00908381 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0090845B | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x009084C2 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009084EC | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0090AC67 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0090ACB3 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x0090AD91 | `MainMenu_List_ScreenLock_x` | Known | Screen layout |
| 0x0090B05F | `RentalInfo_Screen_ExpiresToday` | Known | Screen layout |
| 0x0090B0B5 | `RentalInfo_Screen_NoAlbumArt_ExpiresToday` | Known | Screen layout |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00009077 | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | Known | RTOS |
| 0x0028232C | `  K - RTXC` | Known | RTOS |
| 0x00283314 | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | Known | RTOS |
| 0x008F0AF4 | `Returning from RTXCBug` | Known | RTOS |

---

## 9. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000CEF64 | `HostOSTask` | Known | RTOS task thread |
| 0x00126A00 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x0012BEB0 | `USBDeviceTask` | Known | RTOS task thread |
| 0x0013613C | `DiskReaderTask` | Known | RTOS task thread |
| 0x00145C34 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x00145C48 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x00193F9C | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x001CD1C4 | `MeCCAIOTask` | Known | RTOS task thread |
| 0x001FC400 | `StreamCacheMassStorageManagerTimeOutTask` | Known | RTOS task thread |
| 0x001FC57C | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x00275434 | `FirewireTask` | Known | RTOS task thread |
| 0x00275448 | `TouchwheelTask` | Known | RTOS task thread |
| 0x0027545C | `AudioOutStateTask` | Known | RTOS task thread |
| 0x00275488 | `DiskMgrTask` | Known | RTOS task thread |
| 0x00275498 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x002754AC | `TopPlugTask` | Known | RTOS task thread |
| 0x002754BC | `HPhoneDetTask` | Known | RTOS task thread |
| 0x00275534 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x0027555C | `AlarmTask` | Known | RTOS task thread |
| 0x0027557B | `"USBAudioTask` | Known | RTOS task thread |
| 0x002829CC | `Undefined Task` | Known | RTOS task thread |
| 0x0037DC04 | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x00381840 | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x00389EF4 | `MeCCARecordingTask` | Known | RTOS task thread |
| 0x0084C6CC | `WaveFileDebugTask` | Known | RTOS task thread |

---

## 10. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00232B1C | `Channel Reserved` | Known | Logging channel |
| 0x00232B30 | `Channel AppBoot` | Known | Logging channel |
| 0x00232B40 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x00232B5C | `Channel PrefsWriting` | Known | Logging channel |
| 0x00232B74 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x00232B94 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x00232BAC | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x00232BC8 | `Channel TestLogging` | Known | Logging channel |
| 0x00232BDC | `Channel AppFileLoading` | Known | Logging channel |
| 0x00232BF4 | `Channel VCardReading` | Known | Logging channel |
| 0x00232C0C | `Channel LongSongScanning` | Known | Logging channel |
| 0x00232C80 | `Channel VoiceRecording` | Known | Logging channel |
| 0x00232C98 | `Channel PhotoImporting` | Known | Logging channel |
| 0x00232CB0 | `Channel Notes` | Known | Logging channel |
| 0x00232CC0 | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x00232CDC | `Channel DiskMode` | Known | Logging channel |
| 0x00232CF0 | `Channel Firewire` | Known | Logging channel |
| 0x00232D04 | `Channel USB` | Known | Logging channel |
| 0x00232D24 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x00232D3C | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0007FBD0 | `gamedata_RW` | Known | Game system |
| 0x0007FBEC | `gamedata_ShareRW` | Known | Game system |
| 0x0007FC00 | `games_RO` | Known | Game system |
| 0x008F0B4E | `iPod_Control/games_RO/` | Known | Game system |
| 0x008F0B65 | `Resources/Games/games_RO/` | Known | Game system |
| 0x008FC1C2 | `GamesMenu_StatusBar_String` | Known | Game system |
| 0x008FC8FC | `AboutScreen_Games_String` | Known | Game system |
| 0x009035E8 | `MainMenu_List_Games` | Known | Game system |
| 0x009035FC | `ExtrasMenu_Games` | Known | Game system |
| 0x0090AE00 | `MainMenu_List_Games_x` | Known | Game system |

---

## 12. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0008FF5C | `adrmmp4a` | Known | DRM system |
| 0x00133790 | `AppleDRMVersion` | Known | DRM system |
| 0x00133830 | `AppleDRM` | Known | DRM system |
| 0x00134A14 | `AppleVideoDRM` | Known | DRM system |
| 0x00137EAC | `tx3gdrmsp608aavdmp4aesdsT` | Known | DRM system |
| 0x001DA674 | `drmttx3g` | Known | DRM system |
| 0x008F0F2D | `DRMLevel` | Known | DRM system |

---

## 13. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000304C8 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000304E0 | `iPod_Control\iTunes\firsttime` | Known | iTunes database |
| 0x00050DA8 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x00050DD0 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000574B8 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0007BE28 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x0007FB60 | `iPod_Control/iTunes/iTunesDB.p7b` | Known | iTunes database |
| 0x0009BF78 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x0009C160 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000A4530 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000A59D4 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000A5AD4 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0011F234 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x00376F94 | `iTunesDB` | Known | iTunes database |
| 0x00376FA0 | `iPod_Control\iTunes\` | Known | iTunes database |

---

## 14. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0005E058 | `cI: could not read CE-ATA task file` | Known | Hardware |
| 0x0005E080 | `cI: CE-ATA signature missing (%x,%x)` | Known | Hardware |
| 0x0005E0D8 | `cI: CE-ATA interrupt enable failed` | Known | Hardware |
| 0x0011EA98 | `NAND FLASH DRIVE` | Known | Hardware |
| 0x00133D00 | `FireWireGUID` | Known | FireWire |
| 0x00133D10 | `FireWireVersion` | Known | FireWire |
| 0x001343F4 | `FireWire` | Known | FireWire |
| 0x003262FC | `CE-ATA init failed` | Known | Hardware |
| 0x00326714 | `ISDIE: CE-ATA interrupt enable failed` | Known | Hardware |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x006AEC0E | `Settings_RadioRegionsMenu_Layout"` | Known | FM Radio |
| 0x006AEC97 | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x007516C0 | `Radio Regions` | Known | FM Radio |
| 0x0079F9D4 | `Radio-Regionen` | Known | FM Radio |
| 0x008F9399 | `Settings_RadioRegions_Australia_String` | Known | FM Radio |
| 0x008F93C0 | `Settings_RadioRegions_Asia_String` | Known | FM Radio |
| 0x008FA5A2 | `Settings_RadioRegions_Europe_String` | Known | FM Radio |
| 0x008FBB12 | `Settings_RadioRegions_Japan_String` | Known | FM Radio |
| 0x008FC719 | `Settings_RadioRegions_Americas_String` | Known | FM Radio |
| 0x008FCDFB | `SettingsMenu_RadioRegions_String` | Known | FM Radio |
| 0x0090022D | `RadioRegionsMenuItem` | Known | FM Radio |
| 0x00903D01 | `SettingsMenu_ListItem_RadioRegions` | Known | FM Radio |
| 0x00907CB6 | `SettingsInfo_Template_RadioRegions_Layout` | Known | FM Radio |
| 0x00907CE0 | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x00908342 | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 16. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x007DCDA4 | `Fotocamera` | Known | Camera |
| 0x007DD308 | `Collega fotocamera o inserisci scheda` | Known | Camera |
| 0x007DD380 | `Fotocamera non supportata` | Known | Camera |
| 0x007FB698 | `Camera` | Known | Camera |
| 0x007FBC18 | `Sluit camera of kaart aan` | Known | Camera |
| 0x007FBC84 | `Camera niet ondersteund` | Known | Camera |
| 0x008F93E2 | `Photos_DevName_Camera_String` | Known | Camera |

---

## 17. Pedometer

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0090B168 | `NikePlus_Remote_Step_Away` | Known | Pedometer |
| 0x0090B182 | `NikePlus_Step_Away` | Known | Pedometer |

---

## 18. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000304B4 | `iPod_Control` | Filesystem Path |  |
| 0x00030520 | `iPod_Control\Device` | Filesystem Path |  |
| 0x0003E6E8 | `iPod_Control\Device` | Filesystem Path |  |
| 0x00040764 | `iPod_Control` | Filesystem Path |  |
| 0x00040DCC | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x00050D88 | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path |  |
| 0x000538EC | `iPod_Control\Music\` | Filesystem Path |  |
| 0x00057338 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x00089940 | `iPod_Control` | Filesystem Path |  |
| 0x00089950 | `Resources/Games` | Filesystem Path |  |
| 0x00089960 | `iPod_Control/%s%s%s` | Filesystem Path |  |
| 0x000EF654 | `iPod_Control/Device/alarms` | Filesystem Path |  |
| 0x000FF750 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00100C54 | `iPod_Control/Device` | Filesystem Path |  |
| 0x00100C68 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00119FFC | `iPod_Control/Device/Users` | Filesystem Path |  |
| 0x00146E34 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x00147090 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x00152C40 | `Resources/UI/active.bin` | Filesystem Path |  |
| 0x00152C58 | `Resources/UI/` | Filesystem Path |  |
| 0x001748AC | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x0017556C | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path |  |
| 0x00175594 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001975E4 | `iPod_Control/Device/PlayCounts` | Filesystem Path |  |
| 0x001AD3F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AD4A4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AD620 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AD7B8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AD860 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001ADA10 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001ADAB4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001ADB58 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001ADBFC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001ADCA0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001ADD50 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001ADDF4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001ADE98 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001ADF48 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001ADFF8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AE0A8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AE214 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AE2C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AE374 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AE418 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AE4C8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AE5BC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AE660 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AE714 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AE7D0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AE880 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AE9A4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AEA60 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AEB10 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AECCC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AED90 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AEE40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AEEFC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF038 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF104 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF1C0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF264 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF308 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF3C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF480 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF548 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF5EC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF6B4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF77C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF82C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF8F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AF9BC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AFA6C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AFB1C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AFBE0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AFC90 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AFD40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AFDF0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AFEC4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AFF98 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B0098 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B0178 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B0280 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B036C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00377012 | `iPod_Control/Device` | Filesystem Path |  |
| 0x0037D4A8 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x0037FFF8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003803A6 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003819AC | `Resources/Fonts` | Filesystem Path |  |
| 0x00389EC0 | `iPod_Control:Device:voices:` | Filesystem Path |  |
| 0x008F0A29 | `Resources/Games/` | Filesystem Path |  |
| 0x008F0E0F | `iPod_Control/Device` | Filesystem Path |  |
| 0x008F0E23 | `iPod_Control/Device/accessories.reg` | Filesystem Path |  |
| 0x008F0EA4 | `iPod_Control/Tones` | Filesystem Path |  |

---

## 19. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0084EDF8 | `c:\bwa\N25FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftu` | Build Path |  |
| 0x0084EE50 | `c:\bwa\N25FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fts` | Build Path |  |
| 0x0084EEA8 | `c:\bwa\N25FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fto` | Build Path |  |
| 0x00859478 | `c:\bwa\N25FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\` | Build Path |  |
| 0x00859FF4 | `c:\bwa\N25FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfd` | Build Path |  |
| 0x0085B1F0 | `c:\bwa\N25FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrg` | Build Path |  |
| 0x0085B248 | `c:\bwa\N25FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrc` | Build Path |  |
| 0x0085B2A0 | `c:\bwa\N25FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfro` | Build Path |  |
| 0x0085B5E4 | `c:\bwa\N25FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1` | Build Path |  |
| 0x0086A98C | `c:\bwa\N25FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttc` | Build Path |  |
| 0x0086AC08 | `c:\bwa\N25FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype` | Build Path |  |
| 0x0086B174 | `c:\bwa\N25FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1` | Build Path |  |

---

## 20. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00087648 | `Acoustic` | EQ Preset |  |
| 0x00087654 | `Bass Booster` | EQ Preset |  |
| 0x00087674 | `Classical` | EQ Preset |  |
| 0x00087680 | `Dance` | EQ Preset |  |
| 0x00087690 | `Electronic` | EQ Preset |  |
| 0x000876A4 | `Hip Hop` | EQ Preset |  |
| 0x000876AC | `Jazz` | EQ Preset |  |
| 0x000876B4 | `Latin` | EQ Preset |  |
| 0x000876BC | `Loudness` | EQ Preset |  |
| 0x000876C8 | `Lounge` | EQ Preset |  |
| 0x000876D0 | `Piano` | EQ Preset |  |
| 0x000876E4 | `Rock` | EQ Preset |  |
| 0x000876EC | `Small Speakers` | EQ Preset |  |
| 0x000876FC | `Spoken Word` | EQ Preset |  |
| 0x00087708 | `Treble Booster` | EQ Preset |  |
| 0x00087754 | `Vocal Booster` | EQ Preset |  |
| 0x007519B0 | `Acoustic` | EQ Preset |  |
| 0x007519BC | `Bass Booster` | EQ Preset |  |
| 0x007519DC | `Classical` | EQ Preset |  |
| 0x007519E8 | `Dance` | EQ Preset |  |
| 0x007519F8 | `Electronic` | EQ Preset |  |
| 0x00751A0C | `Hip Hop` | EQ Preset |  |
| 0x00751A14 | `Jazz` | EQ Preset |  |
| 0x00751A1C | `Latin` | EQ Preset |  |
| 0x00751A24 | `Loudness` | EQ Preset |  |
| 0x00751A30 | `Lounge` | EQ Preset |  |
| 0x00751A38 | `Piano` | EQ Preset |  |
| 0x00751A48 | `Rock` | EQ Preset |  |
| 0x00751A50 | `Small Speakers` | EQ Preset |  |
| 0x00751A60 | `Spoken Word` | EQ Preset |  |
| 0x00751A6C | `Treble Booster` | EQ Preset |  |
| 0x00751A8C | `Vocal Booster` | EQ Preset |  |
| 0x0078D75C | `Acoustic` | EQ Preset |  |
| 0x0078D768 | `Bass Booster` | EQ Preset |  |
| 0x0078D788 | `Classical` | EQ Preset |  |
| 0x0078D794 | `Dance` | EQ Preset |  |
| 0x0078D7A4 | `Electronic` | EQ Preset |  |
| 0x0078D7B8 | `Hip Hop` | EQ Preset |  |
| 0x0078D7C0 | `Jazz` | EQ Preset |  |
| 0x0078D7C8 | `Latin` | EQ Preset |  |
| 0x0078D7D0 | `Loudness` | EQ Preset |  |
| 0x0078D7DC | `Lounge` | EQ Preset |  |
| 0x0078D7E4 | `Piano` | EQ Preset |  |
| 0x0078D7F4 | `Rock` | EQ Preset |  |
| 0x0078D7FC | `Small Speakers` | EQ Preset |  |
| 0x0078D80C | `Spoken Word` | EQ Preset |  |
| 0x0078D818 | `Treble Booster` | EQ Preset |  |
| 0x0078D838 | `Vocal Booster` | EQ Preset |  |
| 0x00796808 | `Acoustic` | EQ Preset |  |
| 0x00796814 | `Bass Booster` | EQ Preset |  |
| 0x00796834 | `Classical` | EQ Preset |  |
| 0x00796840 | `Dance` | EQ Preset |  |
| 0x00796850 | `Electronic` | EQ Preset |  |
| 0x00796864 | `Hip Hop` | EQ Preset |  |
| 0x0079686C | `Jazz` | EQ Preset |  |
| 0x00796874 | `Latin` | EQ Preset |  |
| 0x0079687C | `Loudness` | EQ Preset |  |
| 0x00796888 | `Lounge` | EQ Preset |  |
| 0x00796890 | `Piano` | EQ Preset |  |
| 0x007968A0 | `Rock` | EQ Preset |  |
| 0x007968A8 | `Small Speakers` | EQ Preset |  |
| 0x007968B8 | `Spoken Word` | EQ Preset |  |
| 0x007968C4 | `Treble Booster` | EQ Preset |  |
| 0x007968E4 | `Vocal Booster` | EQ Preset |  |
| 0x0079FD7C | `Acoustic` | EQ Preset |  |
| 0x0079FDAC | `Dance` | EQ Preset |  |
| 0x0079FDBC | `Electronic` | EQ Preset |  |
| 0x0079FDD8 | `Jazz` | EQ Preset |  |
| 0x0079FDE0 | `Latin` | EQ Preset |  |
| 0x0079FDE8 | `Loudness` | EQ Preset |  |
| 0x0079FDFC | `Piano` | EQ Preset |  |
| 0x0079FE0C | `Rock` | EQ Preset |  |
| 0x007B736C | `Dance` | EQ Preset |  |
| 0x007B7394 | `Hip Hop` | EQ Preset |  |
| 0x007B739C | `Jazz` | EQ Preset |  |
| 0x007B73AC | `Loudness` | EQ Preset |  |
| 0x007B73B8 | `Lounge` | EQ Preset |  |
| 0x007B73C0 | `Piano` | EQ Preset |  |
| 0x007B73D0 | `Rock` | EQ Preset |  |
| 0x007C0480 | `Jazz` | EQ Preset |  |
| 0x007C0488 | `Latin` | EQ Preset |  |
| 0x007C049C | `Lounge` | EQ Preset |  |
| 0x007C04A4 | `Piano` | EQ Preset |  |
| 0x007C04B4 | `Rock` | EQ Preset |  |
| 0x007C94BC | `Hip Hop` | EQ Preset |  |
| 0x007C94C4 | `Jazz` | EQ Preset |  |
| 0x007C94E0 | `Lounge` | EQ Preset |  |
| 0x007C94E8 | `Piano` | EQ Preset |  |
| 0x007C9500 | `Rock` | EQ Preset |  |
| 0x007D30F4 | `Latin` | EQ Preset |  |
| 0x007D3120 | `Rock` | EQ Preset |  |
| 0x007DC690 | `Dance` | EQ Preset |  |
| 0x007DC6B4 | `Hip Hop` | EQ Preset |  |
| 0x007DC6BC | `Jazz` | EQ Preset |  |
| 0x007DC6CC | `Loudness` | EQ Preset |  |
| 0x007DC6D8 | `Lounge` | EQ Preset |  |
| 0x007DC6E0 | `Piano` | EQ Preset |  |
| 0x007DC6F0 | `Rock` | EQ Preset |  |
| 0x007E6FF4 | `Acoustic` | EQ Preset |  |
| 0x007E7000 | `Bass Booster` | EQ Preset |  |
| 0x007E7020 | `Classical` | EQ Preset |  |
| 0x007E702C | `Dance` | EQ Preset |  |
| 0x007E703C | `Electronic` | EQ Preset |  |
| 0x007E7050 | `Hip Hop` | EQ Preset |  |
| 0x007E7058 | `Jazz` | EQ Preset |  |
| 0x007E7060 | `Latin` | EQ Preset |  |
| 0x007E7068 | `Loudness` | EQ Preset |  |
| 0x007E7074 | `Lounge` | EQ Preset |  |
| 0x007E707C | `Piano` | EQ Preset |  |
| 0x007E708C | `Rock` | EQ Preset |  |
| 0x007E7094 | `Small Speakers` | EQ Preset |  |
| 0x007E70A4 | `Spoken Word` | EQ Preset |  |
| 0x007E70B0 | `Treble Booster` | EQ Preset |  |
| 0x007E70D0 | `Vocal Booster` | EQ Preset |  |
| 0x007F17B8 | `Acoustic` | EQ Preset |  |
| 0x007F17C4 | `Bass Booster` | EQ Preset |  |
| 0x007F17E4 | `Classical` | EQ Preset |  |
| 0x007F17F0 | `Dance` | EQ Preset |  |
| 0x007F1800 | `Electronic` | EQ Preset |  |
| 0x007F1814 | `Hip Hop` | EQ Preset |  |
| 0x007F181C | `Jazz` | EQ Preset |  |
| 0x007F1824 | `Latin` | EQ Preset |  |
| 0x007F182C | `Loudness` | EQ Preset |  |
| 0x007F1838 | `Lounge` | EQ Preset |  |
| 0x007F1840 | `Piano` | EQ Preset |  |
| 0x007F1850 | `Rock` | EQ Preset |  |
| 0x007F1858 | `Small Speakers` | EQ Preset |  |
| 0x007F1868 | `Spoken Word` | EQ Preset |  |
| 0x007F1874 | `Treble Booster` | EQ Preset |  |
| 0x007F1894 | `Vocal Booster` | EQ Preset |  |
| 0x007FAF7C | `Dance` | EQ Preset |  |
| 0x007FAFB0 | `Jazz` | EQ Preset |  |
| 0x007FAFB8 | `Latin` | EQ Preset |  |
| 0x007FAFC0 | `Loudness` | EQ Preset |  |
| 0x007FAFCC | `Lounge` | EQ Preset |  |
| 0x007FAFD4 | `Piano` | EQ Preset |  |
| 0x007FAFE4 | `Rock` | EQ Preset |  |
| 0x00804000 | `Dance` | EQ Preset |  |
| 0x0080402C | `Jazz` | EQ Preset |  |
| 0x0080403C | `Loudness` | EQ Preset |  |
| 0x00804048 | `Lounge` | EQ Preset |  |
| 0x00804050 | `Piano` | EQ Preset |  |
| 0x00804060 | `Rock` | EQ Preset |  |
| 0x0080D330 | `Hip Hop` | EQ Preset |  |
| 0x0080D338 | `Jazz` | EQ Preset |  |
| 0x0080D35C | `Lounge` | EQ Preset |  |
| 0x0080D374 | `Rock` | EQ Preset |  |
| 0x00816A30 | `Hip Hop` | EQ Preset |  |
| 0x00816A38 | `Jazz` | EQ Preset |  |
| 0x00816A54 | `Lounge` | EQ Preset |  |
| 0x00816A5C | `Piano` | EQ Preset |  |
| 0x00816A6C | `Rock` | EQ Preset |  |
| 0x0082CBC4 | `Acoustic` | EQ Preset |  |
| 0x0082CBD0 | `Bass Booster` | EQ Preset |  |
| 0x0082CBF0 | `Classical` | EQ Preset |  |
| 0x0082CBFC | `Dance` | EQ Preset |  |
| 0x0082CC0C | `Electronic` | EQ Preset |  |
| 0x0082CC20 | `Hip Hop` | EQ Preset |  |
| 0x0082CC28 | `Jazz` | EQ Preset |  |
| 0x0082CC30 | `Latin` | EQ Preset |  |
| 0x0082CC38 | `Loudness` | EQ Preset |  |
| 0x0082CC44 | `Lounge` | EQ Preset |  |
| 0x0082CC4C | `Piano` | EQ Preset |  |
| 0x0082CC5C | `Rock` | EQ Preset |  |
| 0x0082CC64 | `Small Speakers` | EQ Preset |  |
| 0x0082CC74 | `Spoken Word` | EQ Preset |  |
| 0x0082CC80 | `Treble Booster` | EQ Preset |  |
| 0x0082CCA0 | `Vocal Booster` | EQ Preset |  |
| 0x00835ED8 | `Hip Hop` | EQ Preset |  |
| 0x00835EE4 | `Latin` | EQ Preset |  |
| 0x00835EEC | `Loudness` | EQ Preset |  |
| 0x00835EF8 | `Lounge` | EQ Preset |  |
| 0x00835F10 | `Rock` | EQ Preset |  |
| 0x0083F2D4 | `Acoustic` | EQ Preset |  |
| 0x0083F2E0 | `Bass Booster` | EQ Preset |  |
| 0x0083F300 | `Classical` | EQ Preset |  |
| 0x0083F30C | `Dance` | EQ Preset |  |
| 0x0083F31C | `Electronic` | EQ Preset |  |
| 0x0083F330 | `Hip Hop` | EQ Preset |  |
| 0x0083F338 | `Jazz` | EQ Preset |  |
| 0x0083F340 | `Latin` | EQ Preset |  |
| 0x0083F348 | `Loudness` | EQ Preset |  |
| 0x0083F354 | `Lounge` | EQ Preset |  |
| 0x0083F35C | `Piano` | EQ Preset |  |
| 0x0083F36C | `Rock` | EQ Preset |  |
| 0x0083F374 | `Small Speakers` | EQ Preset |  |
| 0x0083F384 | `Spoken Word` | EQ Preset |  |
| 0x0083F390 | `Treble Booster` | EQ Preset |  |
| 0x0083F3B0 | `Vocal Booster` | EQ Preset |  |
| 0x00848590 | `Acoustic` | EQ Preset |  |
| 0x0084859C | `Bass Booster` | EQ Preset |  |
| 0x008485BC | `Classical` | EQ Preset |  |
| 0x008485C8 | `Dance` | EQ Preset |  |
| 0x008485D8 | `Electronic` | EQ Preset |  |
| 0x008485EC | `Hip Hop` | EQ Preset |  |
| 0x008485F4 | `Jazz` | EQ Preset |  |
| 0x008485FC | `Latin` | EQ Preset |  |
| 0x00848604 | `Loudness` | EQ Preset |  |
| 0x00848610 | `Lounge` | EQ Preset |  |
| 0x00848618 | `Piano` | EQ Preset |  |
| 0x00848628 | `Rock` | EQ Preset |  |
| 0x00848630 | `Small Speakers` | EQ Preset |  |
| 0x00848640 | `Spoken Word` | EQ Preset |  |
| 0x0084864C | `Treble Booster` | EQ Preset |  |
| 0x0084866C | `Vocal Booster` | EQ Preset |  |

---
