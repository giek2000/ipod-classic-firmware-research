# iPod Classic 6G Initial - RetailOS 1.0.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.0.3 |
| **IPSW** | iPod_24.1.0.3.ipsw |
| **Device** | iPod Classic 6G Initial (2007, 80/160GB, Click Wheel, Cover Flow, CE-ATA HDD) |
| **UpdaterFamilyID** | 24 |
| **Binary Size** | 9,493,328 bytes (9.05 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 9,491,280 bytes |
| **Total Strings (>=4)** | 61,888 |
| **Function Prologues** | 20,457 (ARM: 15,735, Thumb: 4,722) |
| **DRAM References** | 84,960 |
| **Peripheral Refs** | 5,664 |
| **Build** | N25FirmwareWin-313 |
| **SoC** | S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Codename** | N25C |
| **DFU PID** | 0x1223 |
| **SHA-256** | `3f4980e2053dc10c3504885aceb16a02876d586acbe4cc306e6d639faa98b008` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00093980 | `TSilverCntlr` | Known | Controller |
| 0x00093998 | `TCExtrasMenu` | Known | Controller |
| 0x000939B0 | `TCGameScreen` | Known | Controller |
| 0x000939C8 | `TCGamesMenu` | Known | Controller |
| 0x000939DC | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x00093A04 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00093A2C | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x00093A58 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x00093A7C | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x00093AA4 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x00093ACC | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x00093AF4 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x00093B1C | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x00093B44 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x00093B74 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x00093BA0 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x00093BD0 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x00093BF8 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x00093C20 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x00093C4C | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x00093C74 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x00093C9C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x00093CCC | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x00093CFC | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x00093D78 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x00093DA8 | `TSilverGlobalCntlr` | Known | Controller |
| 0x00093DC4 | `TSilverTrainerCntlr` | Known | Controller |
| 0x000E8F60 | `TCSlideshowLCD` | Known | Controller |
| 0x000E8F78 | `TCSlideshowTVOut` | Known | Controller |
| 0x000E8F94 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x000E8FB4 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x0010B0D4 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x0010B100 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x0010B12C | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x0010B154 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x0010B180 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x0010B1A8 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x0010B1D4 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x001122E8 | `TCRemoteUI` | Known | Controller |
| 0x001122FC | `TCUnsupported` | Known | Controller |
| 0x001177AC | `TCSpeakers` | Known | Controller |
| 0x001177C0 | `TCEQSetting` | Known | Controller |
| 0x0013F4C8 | `TCSportTimer` | Known | Controller |
| 0x0013F4E0 | `TCSportTimerMenu` | Known | Controller |
| 0x0013F4FC | `TCSportTimerSessionScreen` | Known | Controller |
| 0x0013F520 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x00140890 | `TCVoiceMemos` | Known | Controller |
| 0x001408A8 | `TCVoiceMemosMenu` | Known | Controller |
| 0x001408C4 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x001408E4 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x00140904 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x00151218 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x00151240 | `TCSettings_MainMenu` | Known | Controller |
| 0x0015125C | `TCSettings_MusicMenu` | Known | Controller |
| 0x0015127C | `TCSettings_VolumeLimit` | Known | Controller |
| 0x0015129C | `TCSettings_Brightness` | Known | Controller |
| 0x001512BC | `TCSettings_BacklightTimer` | Known | Controller |
| 0x001512E0 | `TCSettings_EQ` | Known | Controller |
| 0x001512F8 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x00151320 | `TCSettings_RadioRegions` | Known | Controller |
| 0x00151340 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x00151364 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x00151388 | `TCDateTimeScreen` | Known | Controller |
| 0x001513A4 | `TCTimeZoneScreen` | Known | Controller |
| 0x001513C0 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x001513E8 | `TCFirstBoot` | Known | Controller |
| 0x0016567C | `TCDemoMode` | Known | Controller |
| 0x00189FAC | `TCAddressViewerMainMenu` | Known | Controller |
| 0x00189FCC | `TCAddressViewerDetails` | Known | Controller |
| 0x00189FEC | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x0018A010 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x001B4228 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x001B424C | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x00240FA4 | `TC_LockDialog` | Known | Controller |
| 0x00240FBC | `TC_LockScreen` | Known | Controller |
| 0x00240FD4 | `TC_LockediPod` | Known | Controller |
| 0x00240FEC | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x00241010 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00246804 | `TCClock` | Known | Controller |
| 0x00246814 | `TCClockCityMenu` | Known | Controller |
| 0x0024682C | `TCClockRegionMenu` | Known | Controller |
| 0x00246848 | `TCAlarmMenu` | Known | Controller |
| 0x0024685C | `TCSleepTimerMenu` | Known | Controller |
| 0x00246878 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x00246898 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x002468C0 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x002468E4 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00246908 | `TCAlarmDatePicker` | Known | Controller |
| 0x00246924 | `TCAlarmTriggered` | Known | Controller |
| 0x0024D43C | `TCNotesDispatcher` | Known | Controller |
| 0x0024D458 | `TCNotesLoading` | Known | Controller |
| 0x0024D470 | `TCNotesList` | Known | Controller |
| 0x0024D484 | `TCNotesContents` | Known | Controller |
| 0x00362188 | `TCAlarmTriggered` | Known | Controller |
| 0x0036219C | `TSilverCntlr` | Known | Controller |
| 0x003621BC | `TCClock` | Known | Controller |
| 0x003621C4 | `TCClockRegionMenu` | Known | Controller |
| 0x003621D8 | `TCClockCityMenu` | Known | Controller |
| 0x003621E8 | `TCAlarmMenu` | Known | Controller |
| 0x003621F4 | `TCSleepTimerMenu` | Known | Controller |
| 0x00362208 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x00362220 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x00362240 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x0036225C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00362278 | `TCAlarmDatePicker` | Known | Controller |
| 0x003622B0 | `TSilverCntlr` | Known | Controller |
| 0x003622D0 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x00362460 | `TSilverCntlr` | Known | Controller |
| 0x00362480 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x003624A0 | `TCSettings_Brightness` | Known | Controller |
| 0x003624B8 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x003624D4 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x003624F4 | `TCSettings_RadioRegions` | Known | Controller |
| 0x0036250C | `TCSettings_EQ` | Known | Controller |
| 0x0036251C | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x00362538 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x00362558 | `TCFirstBoot` | Known | Controller |
| 0x00362564 | `TCSettings_MainMenu` | Known | Controller |
| 0x00362578 | `TCSettings_MusicMenu` | Known | Controller |
| 0x00362590 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003625A8 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x003625C4 | `TCDateTimeScreen` | Known | Controller |
| 0x003625D8 | `TCTimeZoneScreen` | Known | Controller |
| 0x003695C4 | `TSilverCntlr` | Known | Controller |
| 0x003695E4 | `TCClock` | Known | Controller |
| 0x003695EC | `TCClockRegionMenu` | Known | Controller |
| 0x00369600 | `TCClockCityMenu` | Known | Controller |
| 0x00369610 | `TCAlarmMenu` | Known | Controller |
| 0x0036961C | `TCSleepTimerMenu` | Known | Controller |
| 0x00369630 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003696A8 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003696C8 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003696E4 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00369718 | `TCAlarmDatePicker` | Known | Controller |
| 0x0036972C | `TCAlarmTriggered` | Known | Controller |
| 0x0036B1A8 | `TSilverCntlr` | Known | Controller |
| 0x0036B1C8 | `TC_LockDialog` | Known | Controller |
| 0x0036B1D8 | `TC_LockScreen` | Known | Controller |
| 0x0036B1E8 | `TC_LockediPod` | Known | Controller |
| 0x0036B1F8 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x0036B214 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0036B22C | `TSilverCntlr` | Known | Controller |
| 0x0036B388 | `TSilverCntlr` | Known | Controller |
| 0x0036B3A8 | `TCRemoteUI` | Known | Controller |
| 0x0036B3B4 | `TCUnsupported` | Known | Controller |
| 0x0036B3C4 | `TSilverCntlr` | Known | Controller |
| 0x0036B428 | `TSilverCntlr` | Known | Controller |
| 0x0036B448 | `TCSportTimer` | Known | Controller |
| 0x0036B458 | `TCSportTimerMenu` | Known | Controller |
| 0x0036B46C | `TCSportTimerSessionScreen` | Known | Controller |
| 0x0036B488 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x0036B4B8 | `TSilverCntlr` | Known | Controller |
| 0x0036B5E0 | `TSilverCntlr` | Known | Controller |
| 0x0036B600 | `TCDemoMode` | Known | Controller |
| 0x0036B618 | `TSilverCntlr` | Known | Controller |
| 0x0036B634 | `TSilverCntlr` | Known | Controller |
| 0x0036B644 | `TSilverCntlr` | Known | Controller |
| 0x0036B664 | `TCVoiceMemos` | Known | Controller |
| 0x0036B674 | `TCVoiceMemosMenu` | Known | Controller |
| 0x0036B688 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x0036B6A0 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x0036B6B8 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x0036B6D8 | `TSilverCntlr` | Known | Controller |
| 0x0036B738 | `TSilverCntlr` | Known | Controller |
| 0x0036B7A4 | `TSilverCntlr` | Known | Controller |
| 0x0036C5DC | `TSilverCntlr` | Known | Controller |
| 0x0036C6E8 | `TSilverCntlr` | Known | Controller |
| 0x00374C50 | `TSilverCntlr` | Known | Controller |
| 0x00374C70 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x00374C88 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x00374CA4 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x00374CC4 | `TCAddressViewerDetails` | Known | Controller |
| 0x00374CDC | `TSilverCntlr` | Known | Controller |
| 0x00374CFC | `TSilverCalendarCntlr_Base` | Known | Controller |
| 0x00374D18 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x00374D3C | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x00374D60 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x00374D80 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x00374DA4 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x00374DC4 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x00374DE8 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x00374FC0 | `TSilverCntlr` | Known | Controller |
| 0x00374FE0 | `TC_LockDialog` | Known | Controller |
| 0x00374FF0 | `TC_LockScreen` | Known | Controller |
| 0x00375000 | `TC_LockediPod` | Known | Controller |
| 0x00375010 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x00375034 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003750F4 | `TSilverCntlr` | Known | Controller |
| 0x00375104 | `TSilverCntlr` | Known | Controller |
| 0x00375278 | `TSilverCntlr` | Known | Controller |
| 0x00375298 | `TCNotesDispatcher` | Known | Controller |
| 0x003752AC | `TCNotesLoading` | Known | Controller |
| 0x003752BC | `TCNotesBase` | Known | Controller |
| 0x003752C8 | `TCNotesList` | Known | Controller |
| 0x003752D4 | `TCNotesContents` | Known | Controller |
| 0x003752E4 | `TSilverCntlr` | Known | Controller |
| 0x003753A8 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003753C4 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003753E4 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x00375404 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0037542C | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x00375450 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x00375478 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x00375498 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003754B8 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003754D8 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003754F8 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x00375520 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x00375570 | `TCSlideshowTVOut` | Known | Controller |
| 0x00375584 | `TCSlideshowLCD` | Known | Controller |
| 0x00375594 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x003755AC | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x003755CC | `TSilverCntlr` | Known | Controller |
| 0x003755F8 | `TSilverCntlr` | Known | Controller |
| 0x00375618 | `TCUnsupported` | Known | Controller |
| 0x00375638 | `TSilverCntlr` | Known | Controller |
| 0x00375678 | `TSilverCntlr` | Known | Controller |
| 0x00375698 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x003756B4 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x003756CC | `TSilverCntlr` | Known | Controller |
| 0x003756EC | `TCSpeakers` | Known | Controller |
| 0x003756F8 | `TCEQSetting` | Known | Controller |
| 0x00375718 | `TSilverCntlr` | Known | Controller |
| 0x00375780 | `TSilverCntlr` | Known | Controller |
| 0x003757A0 | `TCExtrasMenu` | Known | Controller |
| 0x003757B0 | `TCGamesMenu` | Known | Controller |
| 0x003757BC | `TCGameScreen` | Known | Controller |
| 0x003757CC | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x003757EC | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x0037580C | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x0037582C | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x00375850 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x0037586C | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x0037588C | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003758AC | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003758D4 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003758F8 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x00375920 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x00375940 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x00375960 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x00375980 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003759A0 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003759C8 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003759F0 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x00375A10 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x00375A30 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x00375A54 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x00375A74 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x00375A98 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x00375AC0 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x00375AEC | `TSilverGlobalCntlr` | Known | Controller |
| 0x00375B00 | `TSilverTrainerCntlr` | Known | Controller |
| 0x003F8404 | `TCCCCCCCCCCCCCCCCCCCCCCT` | Known | Controller |
| 0x006877BE | `TCNotesDispatcher"` | Known | Controller |
| 0x0068787D | `TCLockChosenDispatcher"` | Known | Controller |
| 0x00687940 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x00691847 | `TCNotesDispatcher"` | Known | Controller |
| 0x006919A9 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x006A5E14 | `TCExtrasMenuTCAddressViewerMainMenu` | Known | Controller |
| 0x006A5E38 | `TCAddressViewerDetails` | Known | Controller |
| 0x006A5E50 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x006A5E6C | `TCAlarmMenu` | Known | Controller |
| 0x006A5E78 | `TCSleepTimerMenuTCAlarmPropertiesMenu` | Known | Controller |
| 0x006A5EA0 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x006A5EC0 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x006A5EDC | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x006A5EF8 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x006A5F14 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x006A5F30 | `TCAlarmDatePicker` | Known | Controller |
| 0x006A5F44 | `TCAlarmDatePicker` | Known | Controller |
| 0x006A5F58 | `TCAlarmTriggeredTSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x006A5F84 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x006A5FA8 | `TSilverCalendarCntlr_MonthViewerTSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x006A5FE8 | `TSilverCalendarCntlr_EventViewerTSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x006A6028 | `TSilverCalendarCntlr_EventViewerTSilverCntlrTCClockRegionMenu` | Known | Controller |
| 0x006A6068 | `TCClockCityMenu` | Known | Controller |
| 0x006A6078 | `TCClockCityMenu` | Known | Controller |
| 0x006A6088 | `TCClockCityMenu` | Known | Controller |
| 0x006A6098 | `TCClockCityMenu` | Known | Controller |
| 0x006A60A8 | `TCClockCityMenu` | Known | Controller |
| 0x006A60B8 | `TCClockCityMenu` | Known | Controller |
| 0x006A60C8 | `TCClockCityMenu` | Known | Controller |
| 0x006A60D8 | `TCClockCityMenu` | Known | Controller |
| 0x006A60E8 | `TCClock` | Known | Controller |
| 0x006A6100 | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDe` | Known | Controller |
| 0x006A6158 | `TCGamesMenu` | Known | Controller |
| 0x006A6164 | `TCGameScreenTC_LockediPod` | Known | Controller |
| 0x006A6180 | `TC_LockDialog` | Known | Controller |
| 0x006A6190 | `TC_LockScreen` | Known | Controller |
| 0x006A61A0 | `TC_VolumeLimitLockScreenTSilverCntlrTSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x006A61E4 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x006A6204 | `TSilverMediaListCntlr_AudiobooksTSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x006A624C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x006A6268 | `TSilverMediaListCntlr_AlbumsTSilverMediaListCntlr_Artists` | Known | Controller |
| 0x006A62A4 | `TSilverMediaListCntlr_GenresTSilverMediaListCntlr_Composers` | Known | Controller |
| 0x006A62E0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x006A6300 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x006A6328 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x006A6348 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x006A6368 | `TSilverMediaListCntlr_TVEpisodesTSilverMediaListCntlr_MoviesTSilverMediaListCntl` | Known | Controller |
| 0x006A63C4 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x006A63EC | `TSilverMediaListCntlr_VideoPlaylistsTSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x006A643C | `TSilverCntlrTSilverMainMediaListCntlr_VideosTSilverSettingsVideoCntlr` | Known | Controller |
| 0x006A6484 | `TCFirstBoot` | Known | Controller |
| 0x006A652C | `TCNotesLoading` | Known | Controller |
| 0x006A653C | `TCNotesList` | Known | Controller |
| 0x006A6548 | `TCNotesList` | Known | Controller |
| 0x006A6554 | `TCNotesContents` | Known | Controller |
| 0x006A6564 | `TCNotesContents` | Known | Controller |
| 0x006A6574 | `TCNotesContents` | Known | Controller |
| 0x006A6584 | `TCNotesContents` | Known | Controller |
| 0x006A6640 | `TCSlideshowLCD` | Known | Controller |
| 0x006A6650 | `TCSlideshowTVOutTCSlideshow_TVOutAskTCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x006A66A0 | `TCRemoteUI` | Known | Controller |
| 0x006A66AC | `TCUnsupported` | Known | Controller |
| 0x006A66BC | `TSilverSettingsMenuListCntlrTSilverSettingsMenuListCntlrTCSettings_MainMenu` | Known | Controller |
| 0x006A6708 | `TCSettings_MusicMenuTCSettings_VolumeLimit` | Known | Controller |
| 0x006A6734 | `TCSettings_Brightness` | Known | Controller |
| 0x006A674C | `TCSettings_BacklightTimer` | Known | Controller |
| 0x006A6768 | `TCSettings_AudiobookSettingsTCSettings_RadioRegions` | Known | Controller |
| 0x006A679C | `TCSettings_EQ` | Known | Controller |
| 0x006A67AC | `TSilverSettingsMenuListCntlrTSilverCntlrTCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x006A67F4 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x006A6810 | `TCSettings_MainMenu` | Known | Controller |
| 0x006A6824 | `TCSettings_MusicMenuTCSportTimerTCSportTimerMenuTCSportTimerSessionScreen` | Known | Controller |
| 0x006A6870 | `TCVoiceMemosMainMenuTCVoiceMemosMenuTCVoiceMemosContextMenu` | Known | Controller |
| 0x006A68AC | `TCVoiceMemosTCVoiceMemosMainMenuTCVoiceMemosMainMenuTCVoiceMemosMainMenuTSearchC` | Known | Controller |
| 0x006A690C | `TCEQSetting` | Known | Controller |
| 0x006A69BA | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x006A7CB9 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x006AD7D9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006AD837 | `TCNotesDispatcher` | Known | Controller |
| 0x006AF2AD | `TCLockChosenDispatcher` | Known | Controller |
| 0x006AF30B | `TCNotesDispatcher` | Known | Controller |
| 0x006B0D81 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006B0DDF | `TCNotesDispatcher` | Known | Controller |
| 0x006B2855 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006B28B3 | `TCNotesDispatcher` | Known | Controller |
| 0x006B4329 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006B4387 | `TCNotesDispatcher` | Known | Controller |
| 0x006B5DFD | `TCLockChosenDispatcher` | Known | Controller |
| 0x006B5E5B | `TCNotesDispatcher` | Known | Controller |
| 0x006B78D1 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006B792F | `TCNotesDispatcher` | Known | Controller |
| 0x006B93A5 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006B9403 | `TCNotesDispatcher` | Known | Controller |
| 0x006BAE79 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006BAED7 | `TCNotesDispatcher` | Known | Controller |
| 0x006BC94D | `TCLockChosenDispatcher` | Known | Controller |
| 0x006BC9AB | `TCNotesDispatcher` | Known | Controller |
| 0x006BE421 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006BE47F | `TCNotesDispatcher` | Known | Controller |
| 0x006BFEF5 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006BFF53 | `TCNotesDispatcher` | Known | Controller |
| 0x006C19C9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006C1A27 | `TCNotesDispatcher` | Known | Controller |
| 0x006C349D | `TCLockChosenDispatcher` | Known | Controller |
| 0x006C34FB | `TCNotesDispatcher` | Known | Controller |
| 0x006C4F71 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006C4FCF | `TCNotesDispatcher` | Known | Controller |
| 0x006C6A45 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006C6AA3 | `TCNotesDispatcher` | Known | Controller |
| 0x006C8519 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006C8577 | `TCNotesDispatcher` | Known | Controller |
| 0x006C9FED | `TCLockChosenDispatcher` | Known | Controller |
| 0x006CA04B | `TCNotesDispatcher` | Known | Controller |
| 0x006CBAC1 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006CBB1F | `TCNotesDispatcher` | Known | Controller |
| 0x006CD595 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006CD5F3 | `TCNotesDispatcher` | Known | Controller |
| 0x006CF069 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006CF0C7 | `TCNotesDispatcher` | Known | Controller |
| 0x006D0B3D | `TCLockChosenDispatcher` | Known | Controller |
| 0x006D0B9B | `TCNotesDispatcher` | Known | Controller |
| 0x006D2611 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006D266F | `TCNotesDispatcher` | Known | Controller |
| 0x006D40E5 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006D4143 | `TCNotesDispatcher` | Known | Controller |
| 0x006D5BB9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006D5C17 | `TCNotesDispatcher` | Known | Controller |
| 0x006D768D | `TCLockChosenDispatcher` | Known | Controller |
| 0x006D76EB | `TCNotesDispatcher` | Known | Controller |
| 0x006D9161 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006D91BF | `TCNotesDispatcher` | Known | Controller |
| 0x006DAC35 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006DAC93 | `TCNotesDispatcher` | Known | Controller |
| 0x006DC709 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006DC767 | `TCNotesDispatcher` | Known | Controller |
| 0x006DE1DD | `TCLockChosenDispatcher` | Known | Controller |
| 0x006DE23B | `TCNotesDispatcher` | Known | Controller |
| 0x006DFCB1 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006DFD0F | `TCNotesDispatcher` | Known | Controller |
| 0x006E1785 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006E17E3 | `TCNotesDispatcher` | Known | Controller |
| 0x006E3259 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006E32B7 | `TCNotesDispatcher` | Known | Controller |
| 0x006E4D2D | `TCLockChosenDispatcher` | Known | Controller |
| 0x006E4D8B | `TCNotesDispatcher` | Known | Controller |
| 0x006E6801 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006E685F | `TCNotesDispatcher` | Known | Controller |
| 0x006E82D5 | `TCLockChosenDispatcher` | Known | Controller |
| 0x006E8333 | `TCNotesDispatcher` | Known | Controller |
| 0x006F3E88 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x006F402A | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x0081128C | `TCMockupModeNavScreen` | Known | Controller |
| 0x008112A4 | `TSilverCntlr` | Known | Controller |
| 0x008112C4 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x00811314 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x00811334 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00811354 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x00811378 | `TCExtrasMenu` | Known | Controller |
| 0x00811D64 | `TSilverCntlr` | Known | Controller |
| 0x00811D84 | `TCSlideshowTVOut` | Known | Controller |
| 0x00811D98 | `TCSlideshowLCD` | Known | Controller |
| 0x00811DA8 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x00811DC0 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00811DFC | `TSilverCntlr` | Known | Controller |
| 0x00811E78 | `TCSlideshowTVOut` | Known | Controller |
| 0x00811E8C | `TCSlideshowLCD` | Known | Controller |
| 0x00811E9C | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x00811EB4 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00811ED4 | `TSilverCntlr` | Known | Controller |
| 0x00811F1C | `TSilverCntlr` | Known | Controller |
| 0x00811F3C | `TCGamesMenu` | Known | Controller |
| 0x00811F48 | `TCGameScreen` | Known | Controller |
| 0x008C3B21 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x008F6919 | `TCL$]` | Known | Controller |

---

## 2. Settings (Toggle/Show/TC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00120800 | `ShowSetting_EQ` | Known | User setting |
| 0x001BCD78 | `ToggleSetting_Repeat` | Known | User setting |
| 0x001BCD94 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001BCDAC | `ToggleSetting_TVOut` | Known | User setting |
| 0x001BCDC0 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x001E3714 | `ShowSetting_Backlight` | Known | User setting |
| 0x001F46C4 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001F46E0 | `ToggleSetting_Repeat` | Known | User setting |
| 0x001F46F8 | `ToggleSetting_SortBy` | Known | User setting |
| 0x001F4710 | `ToggleSetting_ClassicUI` | Known | User setting |
| 0x001F4728 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x001F4744 | `ToggleSetting_Clicker` | Known | User setting |
| 0x001F475C | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x001F477C | `ToggleSetting_24HourClock` | Known | User setting |
| 0x001F4798 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x001F47B4 | `ShowSetting_Shuffle` | Known | User setting |
| 0x001F494C | `ShowSetting_Repeat` | Known | User setting |
| 0x001F4960 | `ShowSetting_About` | Known | User setting |
| 0x001F4974 | `ShowSetting_MainMenu` | Known | User setting |
| 0x001F498C | `ShowSetting_MusicMenu` | Known | User setting |
| 0x001F49A4 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x001F49BC | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x001F49D8 | `ShowSetting_Brightness` | Known | User setting |
| 0x001F49F0 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x001F4A08 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x001F4A24 | `ShowSetting_EQ` | Known | User setting |
| 0x001F4A34 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x001F4BD0 | `ShowSetting_Clicker` | Known | User setting |
| 0x001F4BE4 | `ShowSetting_DateAndTime` | Known | User setting |
| 0x001F4BFC | `ShowSetting_SortBy` | Known | User setting |
| 0x001F4C10 | `ShowSetting_ClassicUI` | Known | User setting |
| 0x001F4C28 | `ShowSetting_Language` | Known | User setting |
| 0x001F4C40 | `ShowSetting_Legal` | Known | User setting |
| 0x001F4C54 | `ShowSetting_ResetAll` | Known | User setting |
| 0x0069074D | `ToggleSetting_24HourClock` | Known | User setting |
| 0x006907FD | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x00692DF2 | `ShowSetting_About` | Known | User setting |
| 0x00692EFA | `ToggleSetting_Shuffle` | Known | User setting |
| 0x00692F3E | `ShowSetting_Shuffle` | Known | User setting |
| 0x00692FB5 | `ToggleSetting_Repeat` | Known | User setting |
| 0x00692FF8 | `ShowSetting_Repeat` | Known | User setting |
| 0x00693102 | `ShowSetting_MainMenu` | Known | User setting |
| 0x00693212 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x006932DA | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x006933A4 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x006934BC | `ShowSetting_Brightness` | Known | User setting |
| 0x006935F2 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x00693703 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x00693804 | `ShowSetting_EQ` | Known | User setting |
| 0x00693871 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x006938B8 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x00693935 | `ToggleSetting_Clicker` | Known | User setting |
| 0x00693979 | `ShowSetting_Clicker` | Known | User setting |
| 0x00693AE0 | `ToggleSetting_SortBy` | Known | User setting |
| 0x00693B23 | `ShowSetting_SortBy` | Known | User setting |
| 0x00693C24 | `ShowSetting_Language` | Known | User setting |
| 0x00693D34 | `ShowSetting_Legal` | Known | User setting |
| 0x00693E65 | `ShowSetting_ResetAll` | Known | User setting |
| 0x00693FD8 | `ShowSetting_Backlight` | Known | User setting |
| 0x00694088 | `ShowSetting_Backlight` | Known | User setting |
| 0x00694138 | `ShowSetting_Backlight` | Known | User setting |
| 0x006941E9 | `ShowSetting_Backlight` | Known | User setting |
| 0x0069429A | `ShowSetting_Backlight` | Known | User setting |
| 0x0069434B | `ShowSetting_Backlight` | Known | User setting |
| 0x006943FF | `ShowSetting_Backlight` | Known | User setting |
| 0x006944AE | `ShowSetting_EQ` | Known | User setting |
| 0x00694523 | `ShowSetting_Language` | Known | User setting |
| 0x00702F14 | `ToggleSetting_Repeat` | Known | User setting |
| 0x00702F4E | `ToggleSetting_Shuffle` | Known | User setting |
| 0x00703010 | `ToggleSetting_TVOut` | Known | User setting |
| 0x00703049 | `ToggleSetting_TVSignal` | Known | User setting |

---

## 3. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0013B340 | `MockupMode/MockupMode.xml` | Hidden | Developer Tool |
| 0x0013B840 | `MockupMode/` | Hidden | Developer Tool |
| 0x0022ACFC | `Channel UnitTests` | Hidden | Developer Tool |
| 0x00278945 | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x00278988 | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x0027899D | `RTXCbug> ` | Hidden | Developer Tool |
| 0x00279379 | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x0028977C | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x00317931 | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x003179F9 | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x00367581 | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x00725178 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0075C5F0 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0076D85C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00783594 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x007942A4 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0079D448 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x007A62C8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x007B9AF4 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x007C2B48 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x007E660C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00802664 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0080AFF8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008B6E08 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008B74A5 | `DemoMode_Background_Image` | Hidden | Demo/Retail Mode |
| 0x008B7FBA | `DemoMode_Icon_Image` | Hidden | Demo/Retail Mode |
| 0x008B9ACC | `UnitTest_Menu_Title` | Hidden | Developer Tool |
| 0x008C1AA3 | `UnitTestModel` | Hidden | Developer Tool |
| 0x008C3344 | `DemoMode_View_Main` | Hidden | Demo/Retail Mode |
| 0x008C3519 | `DemoMode_Icon` | Hidden | Demo/Retail Mode |
| 0x008C4D99 | `DemoMode_Text_Color` | Hidden | Demo/Retail Mode |

---

## 4. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000672F | `"MeCCADecode` | Known | Audio system |
| 0x00131D38 | `AudioCodecs` | Known | Audio system |
| 0x00172640 | `MeCCA_PCM_Output.wav` | Known | Audio system |
| 0x001891D8 | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x00192668 | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x00192870 | `MeCCAVideoDecode` | Known | Audio system |
| 0x0081D640 | `MeCCA_StreamCache` | Known | Audio system |

---

## 5. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E4D14 | `HandleRemoteUIRemotePlayPause` | Known | Event handler |
| 0x000E4D38 | `HandleWheel` | Known | Event handler |
| 0x000E4D44 | `HandlePlayPause` | Known | Event handler |
| 0x000E4D54 | `HandleSelectDown` | Known | Event handler |
| 0x000E4D68 | `HandleNext` | Known | Event handler |
| 0x000E4D74 | `HandlePrevious` | Known | Event handler |
| 0x000E4D84 | `HandleNextPushAndHold` | Known | Event handler |
| 0x000E4D9C | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000E5088 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000E50A8 | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x000F09E4 | `HandleSelect` | Known | Event handler |
| 0x000F09F8 | `HandleHilite` | Known | Event handler |
| 0x000F0D1C | `HandleEQSettingSelected` | Known | Event handler |
| 0x000F114C | `HandleSelect` | Known | Event handler |
| 0x000F1160 | `HandleGameHilited` | Known | Event handler |
| 0x000F1410 | `HandleNotesSelected` | Known | Event handler |
| 0x000F1428 | `HandleNotesPop` | Known | Event handler |
| 0x000F1438 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x000FF2A4 | `HandleVolumeWheel` | Known | Event handler |
| 0x000FF2B8 | `HandleVolumeChange` | Known | Event handler |
| 0x000FF2CC | `HandleTimerDone` | Known | Event handler |
| 0x000FF2DC | `HandleFrequencyChange` | Known | Event handler |
| 0x000FF324 | `HandleTuning` | Known | Event handler |
| 0x00109478 | `HandleLock` | Known | Event handler |
| 0x00109488 | `HandleAddressBook` | Known | Event handler |
| 0x00109C88 | `HandleExit` | Known | Event handler |
| 0x00109C98 | `HandleLap` | Known | Event handler |
| 0x00109CA4 | `HandleResume` | Known | Event handler |
| 0x00109CB4 | `HandleStartStop` | Known | Event handler |
| 0x00109F3C | `HandleWheel` | Known | Event handler |
| 0x00109F4C | `HandlePlayPause` | Known | Event handler |
| 0x00109F5C | `HandleSelectDown` | Known | Event handler |
| 0x00109F70 | `HandleHilite` | Known | Event handler |
| 0x00112BA4 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x00120A34 | `HandleExitUnsupported` | Known | Event handler |
| 0x00137184 | `HandleNotesPop` | Known | Event handler |
| 0x00137198 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0013807C | `HandleSelect` | Known | Event handler |
| 0x00138090 | `HandleWheel` | Known | Event handler |
| 0x0013809C | `HandleImageNext` | Known | Event handler |
| 0x001380AC | `HandleImagePrev` | Known | Event handler |
| 0x001380BC | `HandleImageLast` | Known | Event handler |
| 0x001380CC | `HandleImageFirst` | Known | Event handler |
| 0x001380E0 | `HandlePlayPause` | Known | Event handler |
| 0x001380F0 | `HandlePlay` | Known | Event handler |
| 0x001380FC | `HandlePause` | Known | Event handler |
| 0x0014BEE8 | `HandleSelectCity` | Known | Event handler |
| 0x0014BF00 | `HandleHighlightCity` | Known | Event handler |
| 0x0014CE28 | `HandleWantPopFlow` | Known | Event handler |
| 0x0014CE40 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x0014CE5C | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0014CE78 | `HandleFlowNext` | Known | Event handler |
| 0x0014CE88 | `HandleFlowPrev` | Known | Event handler |
| 0x0014CE98 | `HandleFlowWheel` | Known | Event handler |
| 0x0014CEA8 | `HandleAlbumSelected` | Known | Event handler |
| 0x0014CEBC | `HandlePlayPause` | Known | Event handler |
| 0x0014CECC | `HandleBacksideSongSelected` | Known | Event handler |
| 0x001744DC | `HandleLeaveAlarm` | Known | Event handler |
| 0x001748C4 | `HandleSelect` | Known | Event handler |
| 0x00175784 | `HandleSelect` | Known | Event handler |
| 0x00175798 | `HandleWheel` | Known | Event handler |
| 0x001757A4 | `HandleImageNext` | Known | Event handler |
| 0x001757B4 | `HandleImagePrev` | Known | Event handler |
| 0x001757C4 | `HandleImageLast` | Known | Event handler |
| 0x001757D4 | `HandleImageFirst` | Known | Event handler |
| 0x001757E8 | `HandlePlayPause` | Known | Event handler |
| 0x001757F8 | `HandlePlay` | Known | Event handler |
| 0x00175804 | `HandlePause` | Known | Event handler |
| 0x00175CA4 | `HandleNew` | Known | Event handler |
| 0x00175CB4 | `HandleClear` | Known | Event handler |
| 0x00175CC0 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x00175CDC | `HandleSelectIndexedSession` | Known | Event handler |
| 0x00175FEC | `HandleWheel` | Known | Event handler |
| 0x00175FFC | `HandleArrowUp` | Known | Event handler |
| 0x0017600C | `HandleArrowDown` | Known | Event handler |
| 0x00178230 | `HandleHiliteAlbum` | Known | Event handler |
| 0x00178248 | `HandleBrowseAlbum` | Known | Event handler |
| 0x0017825C | `HandlePlayPause` | Known | Event handler |
| 0x0018C808 | `HandleSelect` | Known | Event handler |
| 0x0018C998 | `HandleSelectRegion` | Known | Event handler |
| 0x001A0E34 | `HandleImageWheel` | Known | Event handler |
| 0x001A0E4C | `HandlePlayPause` | Known | Event handler |
| 0x001A0E5C | `HandleBrowseLarge` | Known | Event handler |
| 0x001A0E70 | `HandleBrowseSmall` | Known | Event handler |
| 0x001A0E84 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x001A0E9C | `HandleImageNext` | Known | Event handler |
| 0x001A0EAC | `HandleImagePrev` | Known | Event handler |
| 0x001A0EBC | `HandleHilite` | Known | Event handler |
| 0x001A0ECC | `HandleImageLast` | Known | Event handler |
| 0x001A0EDC | `HandleImageFirst` | Known | Event handler |
| 0x001A0EF0 | `HandleScreenNext` | Known | Event handler |
| 0x001A0F04 | `HandleScreenPrev` | Known | Event handler |
| 0x001A37A0 | `HandlePlayPause` | Known | Event handler |
| 0x001A37B4 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x001A37D0 | `HandleNext` | Known | Event handler |
| 0x001A37DC | `HandleNextPressAndHold` | Known | Event handler |
| 0x001A37F4 | `HandlePrevious` | Known | Event handler |
| 0x001A3804 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x001A3820 | `HandleRemotePlayPause` | Known | Event handler |
| 0x001A3838 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x001A385C | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x001A3874 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x001A388C | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x001A3A5C | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x001A3A74 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x001A3A8C | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x001A3AA8 | `HandleRemoteStop` | Known | Event handler |
| 0x001A3ABC | `HandleRemotePlay` | Known | Event handler |
| 0x001A3AD0 | `HandleRemotePause` | Known | Event handler |
| 0x001A3AE4 | `HandleRemoteMute` | Known | Event handler |
| 0x001A3AF8 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x001A3B10 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x001A3B28 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x001A3B44 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x001A3D68 | `HandleRemoteShuffle` | Known | Event handler |
| 0x001A3D7C | `HandleRemoteRepeat` | Known | Event handler |
| 0x001A3D90 | `HandleRemoteOn` | Known | Event handler |
| 0x001A3DA0 | `HandleRemoteOff` | Known | Event handler |
| 0x001A3DB0 | `HandleRemoteBacklight` | Known | Event handler |
| 0x001A3DC8 | `HandleRemoteFFDown` | Known | Event handler |
| 0x001A3DDC | `HandleRemoteFFUp` | Known | Event handler |
| 0x001A3DF0 | `HandleRemoteRewDown` | Known | Event handler |
| 0x001A3E04 | `HandleRemoteRewUp` | Known | Event handler |
| 0x001A3E18 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x001A3E30 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x001A3E44 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x001A3E5C | `HandleRemoteSelectUp` | Known | Event handler |
| 0x001A402C | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x001A4044 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x001A405C | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x001A4078 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x001A4090 | `HandleRemoteEvent` | Known | Event handler |
| 0x001A40A4 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x001A40C0 | `HandleAudioPlayPause` | Known | Event handler |
| 0x001A40D8 | `HandleAudioNext` | Known | Event handler |
| 0x001A40E8 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x001A4104 | `HandleAudioPrevious` | Known | Event handler |
| 0x001A4118 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x001A4318 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x001A4330 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x001A4348 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x001A4360 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x001A4374 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x001A438C | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x001A43A4 | `HandleAudioStop` | Known | Event handler |
| 0x001A43B4 | `HandleAudioPlay` | Known | Event handler |
| 0x001A43C4 | `HandleAudioPause` | Known | Event handler |
| 0x001A43D8 | `HandleAudioMute` | Known | Event handler |
| 0x001A43E8 | `HandleAudioNextChapter` | Known | Event handler |
| 0x001A4400 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x001A4620 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x001A4638 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x001A4650 | `HandleAudioShuffle` | Known | Event handler |
| 0x001A4664 | `HandleAudioRepeat` | Known | Event handler |
| 0x001A4678 | `HandleAudioFFDown` | Known | Event handler |
| 0x001A468C | `HandleAudioFFUp` | Known | Event handler |
| 0x001A469C | `HandleAudioRewDown` | Known | Event handler |
| 0x001A46B0 | `HandleAudioRewUp` | Known | Event handler |
| 0x001A46C4 | `HandleVideoPlayPause` | Known | Event handler |
| 0x001A46DC | `HandleVideoNext` | Known | Event handler |
| 0x001A46EC | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x001A4708 | `HandleVideoPrevious` | Known | Event handler |
| 0x001A471C | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x001A4924 | `HandleVideoStop` | Known | Event handler |
| 0x001A4934 | `HandleVideoPlay` | Known | Event handler |
| 0x001A4944 | `HandleVideoPause` | Known | Event handler |
| 0x001A4958 | `HandleVideoFFDown` | Known | Event handler |
| 0x001A496C | `HandleVideoFFUp` | Known | Event handler |
| 0x001A497C | `HandleVideoRewDown` | Known | Event handler |
| 0x001A4990 | `HandleVideoRewUp` | Known | Event handler |
| 0x001A49A4 | `HandleVideoNextChapter` | Known | Event handler |
| 0x001A49BC | `HandleVideoPrevChapter` | Known | Event handler |
| 0x001A49D4 | `HandleVideoNextFrame` | Known | Event handler |
| 0x001A49EC | `HandleVideoPrevFrame` | Known | Event handler |
| 0x001A4A04 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x001B0E50 | `HandleMainMenu` | Known | Event handler |
| 0x001B5204 | `HandlePowerSongSelected` | Known | Event handler |
| 0x001B5220 | `HandlePowerSongChosen` | Known | Event handler |
| 0x001B5238 | `HandleSelectTimedWorkout` | Known | Event handler |
| 0x001BB784 | `HandleMusicMenu` | Known | Event handler |
| 0x001BBA44 | `HandleSelect` | Known | Event handler |
| 0x001BBDC8 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x001BBDE0 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x001BBE00 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x001BBE24 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x001BBE40 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x001BC2DC | `HandleWheel` | Known | Event handler |
| 0x001BC2EC | `HandlePlayPause` | Known | Event handler |
| 0x001BC2FC | `HandleSelectDown` | Known | Event handler |
| 0x001BC310 | `HandleNext` | Known | Event handler |
| 0x001BC31C | `HandlePrevious` | Known | Event handler |
| 0x001BC32C | `HandleNextPushAndHold` | Known | Event handler |
| 0x001BC344 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x001C7A7C | `HandleFrequencyChosen` | Known | Event handler |
| 0x001C7A94 | `HandleDateChosen` | Known | Event handler |
| 0x001C7AA8 | `HandleTimeChosen` | Known | Event handler |
| 0x001C7ABC | `HandleSoundChosen` | Known | Event handler |
| 0x001C7AD0 | `HandleLabelChosen` | Known | Event handler |
| 0x001C7AE4 | `HandleDeleteChosen` | Known | Event handler |
| 0x001CCFC8 | `HandlePrev` | Known | Event handler |
| 0x001CCFD8 | `HandleNext` | Known | Event handler |
| 0x001CCFE4 | `HandlePlayPause` | Known | Event handler |
| 0x001D41B8 | `HandleNextContact` | Known | Event handler |
| 0x001D41D0 | `HandlePreviousContact` | Known | Event handler |
| 0x001DBCD8 | `HandleItemSelected` | Known | Event handler |
| 0x001DBED0 | `HandleRadioRegion` | Known | Event handler |
| 0x001DC0B8 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x001E0058 | `HandlePlayPause` | Known | Event handler |
| 0x001E39F0 | `HandleDelete` | Known | Event handler |
| 0x001E3A04 | `HandleSelectLozinch` | Known | Event handler |
| 0x001E3CAC | `HandleSelect` | Known | Event handler |
| 0x001E3F28 | `HandleTVOutChanged` | Known | Event handler |
| 0x001E3F40 | `HandleTVSignalChanged` | Known | Event handler |
| 0x001E3F58 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x001E3F78 | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x001E3F98 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x001E6AF4 | `HandleSelectKey` | Known | Event handler |
| 0x001E6C9C | `HandleSelect` | Known | Event handler |
| 0x001E7A1C | `HandlePlayPause` | Known | Event handler |
| 0x001E7A30 | `HandleWheel` | Known | Event handler |
| 0x001E7A3C | `HandleWheelRating` | Known | Event handler |
| 0x001E7A50 | `HandleWheelScrub` | Known | Event handler |
| 0x001E7A64 | `HandleWheelVolume` | Known | Event handler |
| 0x001E85EC | `HandleSelect` | Known | Event handler |
| 0x001E8DA8 | `HandleLeaveAlarm` | Known | Event handler |
| 0x001E9BDC | `HandleSelect` | Known | Event handler |
| 0x001E9BF0 | `HandleHilite` | Known | Event handler |
| 0x001E9C00 | `HandlePlayPause` | Known | Event handler |
| 0x001E9C10 | `HandleAddToOTG` | Known | Event handler |
| 0x001E9C20 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x001EC928 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x001ED134 | `HandleSelect` | Known | Event handler |
| 0x001ED148 | `HandleWheel` | Known | Event handler |
| 0x001ED154 | `HandleWheelProgress` | Known | Event handler |
| 0x001ED168 | `HandleSelectProgress` | Known | Event handler |
| 0x001ED180 | `HandleSelectVolume` | Known | Event handler |
| 0x001ED194 | `HandleSelectScrub` | Known | Event handler |
| 0x001ED1A8 | `HandleSelectRating` | Known | Event handler |
| 0x001ED1BC | `HandleSelectExtraInfo` | Known | Event handler |
| 0x001ED1D4 | `HandleSelectChapterArt` | Known | Event handler |
| 0x001ED1EC | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x001ED208 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x001ED224 | `HandleWheelBrightness` | Known | Event handler |
| 0x001ED36C | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x001EECA0 | `HandleSelect` | Known | Event handler |
| 0x001EECB0 | `HandleSelectRating` | Known | Event handler |
| 0x001EECC4 | `HandleSelectProgress` | Known | Event handler |
| 0x001EECDC | `HandleWheelProgress` | Known | Event handler |
| 0x001EECF0 | `HandleSelectScrub` | Known | Event handler |
| 0x001EED04 | `HandleWheelBrightness` | Known | Event handler |
| 0x001EED1C | `HandleNextTrackGotoScrub` | Known | Event handler |
| 0x001EED38 | `HandlePrevTrackGotoScrub` | Known | Event handler |
| 0x001EED54 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x001F4C8C | `HandleLanguage` | Known | Event handler |
| 0x001F4C9C | `HandleResetAllSettings` | Known | Event handler |
| 0x001F4CB4 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x001F5598 | `HandleSelect` | Known | Event handler |
| 0x001F57C8 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x001F8698 | `HandleSelect` | Known | Event handler |
| 0x001F8834 | `HandleSelect` | Known | Event handler |
| 0x001F8AD4 | `HandleNextDay` | Known | Event handler |
| 0x001F8AE8 | `HandlePreviousDay` | Known | Event handler |
| 0x001F92EC | `HandleMusicHilited` | Known | Event handler |
| 0x001F9304 | `HandleVideosHilited` | Known | Event handler |
| 0x001F9318 | `HandlePodcastsHilited` | Known | Event handler |
| 0x001F9330 | `HandleGenericHilited` | Known | Event handler |
| 0x001F9348 | `HandlePhotosHilited` | Known | Event handler |
| 0x001F935C | `HandleNowPlayingHilited` | Known | Event handler |
| 0x001F9374 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x001F9390 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x001F93A8 | `HandleArtistsHilited` | Known | Event handler |
| 0x001F93C0 | `HandleGenresHilited` | Known | Event handler |
| 0x001F93D4 | `HandleAlbumsHilited` | Known | Event handler |
| 0x001F93E8 | `HandleCompilationsHilited` | Known | Event handler |
| 0x001F95BC | `HandleComposersHilited` | Known | Event handler |
| 0x001F95D4 | `HandleSongsHilited` | Known | Event handler |
| 0x001F95E8 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x001F9600 | `HandleTVShowsHilited` | Known | Event handler |
| 0x001F9618 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x001F9634 | `HandleMoviesHilited` | Known | Event handler |
| 0x001F9648 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x001F9664 | `HandleMusicSelected` | Known | Event handler |
| 0x001F9678 | `HandleVideosSelected` | Known | Event handler |
| 0x001F9690 | `HandlePodcastsSelected` | Known | Event handler |
| 0x001F96A8 | `HandlePhotosSelected` | Known | Event handler |
| 0x001F9878 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x001F9890 | `HandleSongsSelected` | Known | Event handler |
| 0x001F98A4 | `HandleAlbumsSelected` | Known | Event handler |
| 0x001F98BC | `HandleCompilationsSelected` | Known | Event handler |
| 0x001F98D8 | `HandleArtistsSelected` | Known | Event handler |
| 0x001F98F0 | `HandleGenresSelected` | Known | Event handler |
| 0x001F9908 | `HandleComposersSelected` | Known | Event handler |
| 0x001F9920 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x001F993C | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x001F9958 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x001F9970 | `HandleNowPlaying` | Known | Event handler |
| 0x001F9B0C | `HandleTVShowsSelected` | Known | Event handler |
| 0x001F9B24 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x001F9B40 | `HandleMoviesSelected` | Known | Event handler |
| 0x001F9B58 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x001F9B78 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x001F9B90 | `HandleRadioPlayPause` | Known | Event handler |
| 0x001F9BA8 | `HandleLock` | Known | Event handler |
| 0x001F9BB4 | `HandleBacklightSelected` | Known | Event handler |
| 0x001F9BCC | `HandleSleepSelected` | Known | Event handler |
| 0x001F9BE0 | `HandleNikePlusSelected` | Known | Event handler |
| 0x001FC268 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x001FC86C | `HandleWheel` | Known | Event handler |
| 0x001FDC6C | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x001FDEC4 | `HandleNextDay` | Known | Event handler |
| 0x001FDED8 | `HandlePreviousDay` | Known | Event handler |
| 0x001FE120 | `HandleSelect` | Known | Event handler |
| 0x001FE3BC | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00200C94 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x00200CB0 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x00201C18 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x002022F8 | `HandleSelect` | Known | Event handler |
| 0x002029C4 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x00237530 | `HandleDeleteClock` | Known | Event handler |
| 0x00237548 | `HandleSelectClock` | Known | Event handler |
| 0x0023755C | `HandleHilited` | Known | Event handler |
| 0x0023756C | `HandleWheel` | Known | Event handler |
| 0x00237578 | `HandleSelectLozinch` | Known | Event handler |
| 0x00393672 | `HandleAudioFFDown` | Known | Event handler |
| 0x0039369B | `HandleAudioFFUp` | Known | Event handler |
| 0x003936C6 | `HandleAudioMute` | Known | Event handler |
| 0x003936F9 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x0039372E | `HandleAudioNext` | Known | Event handler |
| 0x0039375E | `HandleAudioNextAlbum` | Known | Event handler |
| 0x00393795 | `HandleAudioNextChapter` | Known | Event handler |
| 0x003937CF | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x00393803 | `HandleAudioPause` | Known | Event handler |
| 0x0039382F | `HandleAudioPlay` | Known | Event handler |
| 0x0039385D | `HandleAudioPlayPause` | Known | Event handler |
| 0x00393895 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x003938CE | `HandleAudioPrevious` | Known | Event handler |
| 0x00393902 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x00393939 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x00393973 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x003939A8 | `HandleAudioRepeat` | Known | Event handler |
| 0x003939D4 | `HandleAudioRewDown` | Known | Event handler |
| 0x003939FF | `HandleAudioRewUp` | Known | Event handler |
| 0x00393A2E | `HandleAudioShuffle` | Known | Event handler |
| 0x00393A5C | `HandleAudioStop` | Known | Event handler |
| 0x00393A8D | `HandleAudioVolumeDown` | Known | Event handler |
| 0x00393AC2 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x00393AF9 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x00393B2A | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x00393BE3 | `HandleNextPressAndHold` | Known | Event handler |
| 0x00393C14 | `HandleNext` | Known | Event handler |
| 0x00393C4C | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x00393C87 | `HandlePlayPause` | Known | Event handler |
| 0x00393CBB | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x00393CF0 | `HandlePrevious` | Known | Event handler |
| 0x00393D7D | `HandleRemoteBacklight` | Known | Event handler |
| 0x00393DB5 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x00393DEF | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x00393E28 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x00393E5D | `HandleRemoteEvent` | Known | Event handler |
| 0x00393E89 | `HandleRemoteFFDown` | Known | Event handler |
| 0x00393EB4 | `HandleRemoteFFUp` | Known | Event handler |
| 0x00393EE1 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x00393F10 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x00393F3F | `HandleRemoteMute` | Known | Event handler |
| 0x00393F71 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x00393FAA | `HandleRemoteNextChapter` | Known | Event handler |
| 0x00393FE6 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x00394026 | `HandleRemoteOff` | Known | Event handler |
| 0x0039404F | `HandleRemoteOff` | Known | Event handler |
| 0x00394079 | `HandleRemoteOn` | Known | Event handler |
| 0x003940A5 | `HandleRemotePause` | Known | Event handler |
| 0x003940D3 | `HandleRemotePlay` | Known | Event handler |
| 0x00394111 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x00394152 | `HandleRemotePlayPause` | Known | Event handler |
| 0x00394189 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x003941C2 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x003941FE | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x00394235 | `HandleRemoteRepeat` | Known | Event handler |
| 0x00394263 | `HandleRemoteRewDown` | Known | Event handler |
| 0x00394290 | `HandleRemoteRewUp` | Known | Event handler |
| 0x003942C0 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x003942F3 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x00394327 | `HandleRemoteShuffle` | Known | Event handler |
| 0x00394357 | `HandleRemoteStop` | Known | Event handler |
| 0x00394387 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x003943BC | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x003943F4 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x0039442B | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x00394464 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x00394497 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x003944CC | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x003944FF | `HandleVideoFFDown` | Known | Event handler |
| 0x00394528 | `HandleVideoFFUp` | Known | Event handler |
| 0x0039455B | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x00394590 | `HandleVideoNext` | Known | Event handler |
| 0x003945C2 | `HandleVideoNextChapter` | Known | Event handler |
| 0x003945F9 | `HandleVideoNextFrame` | Known | Event handler |
| 0x0039462A | `HandleVideoPause` | Known | Event handler |
| 0x00394656 | `HandleVideoPlay` | Known | Event handler |
| 0x00394684 | `HandleVideoPlayPause` | Known | Event handler |
| 0x003946BC | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x003946F5 | `HandleVideoPrevious` | Known | Event handler |
| 0x0039472B | `HandleVideoPrevChapter` | Known | Event handler |
| 0x00394762 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x00394791 | `HandleVideoRewDown` | Known | Event handler |
| 0x003947BC | `HandleVideoRewUp` | Known | Event handler |
| 0x003947E8 | `HandleVideoStop` | Known | Event handler |
| 0x00687542 | `HandleAddressBook` | Known | Event handler |
| 0x006879F6 | `HandleSelect` | Known | Event handler |
| 0x00687A31 | `HandleHilite` | Known | Event handler |
| 0x00687AB2 | `HandleSelectRegion` | Known | Event handler |
| 0x00687B52 | `HandleSelectRegion` | Known | Event handler |
| 0x00687BEE | `HandleSelectRegion` | Known | Event handler |
| 0x00687C92 | `HandleSelectRegion` | Known | Event handler |
| 0x00687D38 | `HandleSelectRegion` | Known | Event handler |
| 0x00687DD8 | `HandleSelectRegion` | Known | Event handler |
| 0x00687E84 | `HandleSelectRegion` | Known | Event handler |
| 0x00687F26 | `HandleSelectRegion` | Known | Event handler |
| 0x00687FD6 | `HandleSelectCity` | Known | Event handler |
| 0x00688042 | `HandleHighlightCity` | Known | Event handler |
| 0x0068807B | `HandleSelectCity` | Known | Event handler |
| 0x006880E7 | `HandleHighlightCity` | Known | Event handler |
| 0x00688120 | `HandleSelectCity` | Known | Event handler |
| 0x0068818C | `HandleHighlightCity` | Known | Event handler |
| 0x006881C5 | `HandleSelectCity` | Known | Event handler |
| 0x00688231 | `HandleHighlightCity` | Known | Event handler |
| 0x0068826A | `HandleSelectCity` | Known | Event handler |
| 0x006882D6 | `HandleHighlightCity` | Known | Event handler |
| 0x0068830F | `HandleSelectCity` | Known | Event handler |
| 0x0068837B | `HandleHighlightCity` | Known | Event handler |
| 0x006883B4 | `HandleSelectCity` | Known | Event handler |
| 0x00688420 | `HandleHighlightCity` | Known | Event handler |
| 0x00688459 | `HandleSelectCity` | Known | Event handler |
| 0x006884C5 | `HandleHighlightCity` | Known | Event handler |
| 0x006884FE | `HandleSelectCity` | Known | Event handler |
| 0x0068856A | `HandleHighlightCity` | Known | Event handler |
| 0x006885A3 | `HandleSelectCity` | Known | Event handler |
| 0x0068860F | `HandleHighlightCity` | Known | Event handler |
| 0x00688648 | `HandleSelectCity` | Known | Event handler |
| 0x006886B4 | `HandleHighlightCity` | Known | Event handler |
| 0x006886ED | `HandleSelectCity` | Known | Event handler |
| 0x00688759 | `HandleHighlightCity` | Known | Event handler |
| 0x00688792 | `HandleSelectCity` | Known | Event handler |
| 0x006887FE | `HandleHighlightCity` | Known | Event handler |
| 0x00688837 | `HandleSelectCity` | Known | Event handler |
| 0x006888A3 | `HandleHighlightCity` | Known | Event handler |
| 0x006888DC | `HandleSelectCity` | Known | Event handler |
| 0x00688948 | `HandleHighlightCity` | Known | Event handler |
| 0x00688981 | `HandleSelectCity` | Known | Event handler |
| 0x006889ED | `HandleHighlightCity` | Known | Event handler |
| 0x00688A26 | `HandleSelectCity` | Known | Event handler |
| 0x00688A92 | `HandleHighlightCity` | Known | Event handler |
| 0x00688ACB | `HandleSelectCity` | Known | Event handler |
| 0x00688B37 | `HandleHighlightCity` | Known | Event handler |
| 0x00688B70 | `HandleSelectCity` | Known | Event handler |
| 0x00688BDC | `HandleHighlightCity` | Known | Event handler |
| 0x00688C15 | `HandleSelectCity` | Known | Event handler |
| 0x00688C81 | `HandleHighlightCity` | Known | Event handler |
| 0x00688CBA | `HandleSelectCity` | Known | Event handler |
| 0x00688D26 | `HandleHighlightCity` | Known | Event handler |
| 0x00688D5F | `HandleSelectCity` | Known | Event handler |
| 0x00688DCB | `HandleHighlightCity` | Known | Event handler |
| 0x00688E04 | `HandleSelectCity` | Known | Event handler |
| 0x00688E70 | `HandleHighlightCity` | Known | Event handler |
| 0x00688EA9 | `HandleSelectCity` | Known | Event handler |
| 0x00688F15 | `HandleHighlightCity` | Known | Event handler |
| 0x00688F4E | `HandleSelectCity` | Known | Event handler |
| 0x00688FBA | `HandleHighlightCity` | Known | Event handler |
| 0x00688FF3 | `HandleSelectCity` | Known | Event handler |
| 0x0068905F | `HandleHighlightCity` | Known | Event handler |
| 0x00689098 | `HandleSelectCity` | Known | Event handler |
| 0x00689104 | `HandleHighlightCity` | Known | Event handler |
| 0x0068913D | `HandleSelectCity` | Known | Event handler |
| 0x006891A9 | `HandleHighlightCity` | Known | Event handler |
| 0x006891E2 | `HandleSelectCity` | Known | Event handler |
| 0x0068924E | `HandleHighlightCity` | Known | Event handler |
| 0x00689287 | `HandleSelectCity` | Known | Event handler |
| 0x006892F3 | `HandleHighlightCity` | Known | Event handler |
| 0x0068932C | `HandleSelectCity` | Known | Event handler |
| 0x00689398 | `HandleHighlightCity` | Known | Event handler |
| 0x006893D6 | `HandleSelectCity` | Known | Event handler |
| 0x00689442 | `HandleHighlightCity` | Known | Event handler |
| 0x0068947B | `HandleSelectCity` | Known | Event handler |
| 0x006894E7 | `HandleHighlightCity` | Known | Event handler |
| 0x00689520 | `HandleSelectCity` | Known | Event handler |
| 0x0068958C | `HandleHighlightCity` | Known | Event handler |
| 0x006895C5 | `HandleSelectCity` | Known | Event handler |
| 0x00689631 | `HandleHighlightCity` | Known | Event handler |
| 0x0068966A | `HandleSelectCity` | Known | Event handler |
| 0x006896D6 | `HandleHighlightCity` | Known | Event handler |
| 0x0068970F | `HandleSelectCity` | Known | Event handler |
| 0x0068977B | `HandleHighlightCity` | Known | Event handler |
| 0x006897B4 | `HandleSelectCity` | Known | Event handler |
| 0x00689820 | `HandleHighlightCity` | Known | Event handler |
| 0x00689859 | `HandleSelectCity` | Known | Event handler |
| 0x006898C5 | `HandleHighlightCity` | Known | Event handler |
| 0x006898FE | `HandleSelectCity` | Known | Event handler |
| 0x0068996A | `HandleHighlightCity` | Known | Event handler |
| 0x006899A3 | `HandleSelectCity` | Known | Event handler |
| 0x00689A0F | `HandleHighlightCity` | Known | Event handler |
| 0x00689A48 | `HandleSelectCity` | Known | Event handler |
| 0x00689AB4 | `HandleHighlightCity` | Known | Event handler |
| 0x00689AED | `HandleSelectCity` | Known | Event handler |
| 0x00689B59 | `HandleHighlightCity` | Known | Event handler |
| 0x00689B92 | `HandleSelectCity` | Known | Event handler |
| 0x00689BFE | `HandleHighlightCity` | Known | Event handler |
| 0x00689C37 | `HandleSelectCity` | Known | Event handler |
| 0x00689CA3 | `HandleHighlightCity` | Known | Event handler |
| 0x00689CDC | `HandleSelectCity` | Known | Event handler |
| 0x00689D48 | `HandleHighlightCity` | Known | Event handler |
| 0x00689D81 | `HandleSelectCity` | Known | Event handler |
| 0x00689DED | `HandleHighlightCity` | Known | Event handler |
| 0x00689E26 | `HandleSelectCity` | Known | Event handler |
| 0x00689E92 | `HandleHighlightCity` | Known | Event handler |
| 0x00689ECB | `HandleSelectCity` | Known | Event handler |
| 0x00689F37 | `HandleHighlightCity` | Known | Event handler |
| 0x00689F70 | `HandleSelectCity` | Known | Event handler |
| 0x00689FDC | `HandleHighlightCity` | Known | Event handler |
| 0x0068A015 | `HandleSelectCity` | Known | Event handler |
| 0x0068A081 | `HandleHighlightCity` | Known | Event handler |
| 0x0068A0BA | `HandleSelectCity` | Known | Event handler |
| 0x0068A126 | `HandleHighlightCity` | Known | Event handler |
| 0x0068A15F | `HandleSelectCity` | Known | Event handler |
| 0x0068A1CB | `HandleHighlightCity` | Known | Event handler |
| 0x0068A204 | `HandleSelectCity` | Known | Event handler |
| 0x0068A270 | `HandleHighlightCity` | Known | Event handler |
| 0x0068A2A9 | `HandleSelectCity` | Known | Event handler |
| 0x0068A315 | `HandleHighlightCity` | Known | Event handler |
| 0x0068A34E | `HandleSelectCity` | Known | Event handler |
| 0x0068A3BA | `HandleHighlightCity` | Known | Event handler |
| 0x0068A3F3 | `HandleSelectCity` | Known | Event handler |
| 0x0068A45F | `HandleHighlightCity` | Known | Event handler |
| 0x0068A498 | `HandleSelectCity` | Known | Event handler |
| 0x0068A504 | `HandleHighlightCity` | Known | Event handler |
| 0x0068A53D | `HandleSelectCity` | Known | Event handler |
| 0x0068A5A9 | `HandleHighlightCity` | Known | Event handler |
| 0x0068A5E2 | `HandleSelectCity` | Known | Event handler |
| 0x0068A64E | `HandleHighlightCity` | Known | Event handler |
| 0x0068A687 | `HandleSelectCity` | Known | Event handler |
| 0x0068A6F3 | `HandleHighlightCity` | Known | Event handler |
| 0x0068A72C | `HandleSelectCity` | Known | Event handler |
| 0x0068A798 | `HandleHighlightCity` | Known | Event handler |
| 0x0068A7D1 | `HandleSelectCity` | Known | Event handler |
| 0x0068A83D | `HandleHighlightCity` | Known | Event handler |
| 0x0068A876 | `HandleSelectCity` | Known | Event handler |
| 0x0068A8E2 | `HandleHighlightCity` | Known | Event handler |
| 0x0068A91B | `HandleSelectCity` | Known | Event handler |
| 0x0068A987 | `HandleHighlightCity` | Known | Event handler |
| 0x0068A9C0 | `HandleSelectCity` | Known | Event handler |
| 0x0068AA2C | `HandleHighlightCity` | Known | Event handler |
| 0x0068AA65 | `HandleSelectCity` | Known | Event handler |
| 0x0068AAD1 | `HandleHighlightCity` | Known | Event handler |
| 0x0068AB0A | `HandleSelectCity` | Known | Event handler |
| 0x0068AB76 | `HandleHighlightCity` | Known | Event handler |
| 0x0068ABAF | `HandleSelectCity` | Known | Event handler |
| 0x0068AC1B | `HandleHighlightCity` | Known | Event handler |
| 0x0068AC54 | `HandleSelectCity` | Known | Event handler |
| 0x0068ACC0 | `HandleHighlightCity` | Known | Event handler |
| 0x0068ACF9 | `HandleSelectCity` | Known | Event handler |
| 0x0068AD65 | `HandleHighlightCity` | Known | Event handler |
| 0x0068AD9E | `HandleSelectCity` | Known | Event handler |
| 0x0068AE0A | `HandleHighlightCity` | Known | Event handler |
| 0x0068AE43 | `HandleSelectCity` | Known | Event handler |
| 0x0068AEAF | `HandleHighlightCity` | Known | Event handler |
| 0x0068AEE8 | `HandleSelectCity` | Known | Event handler |
| 0x0068AF54 | `HandleHighlightCity` | Known | Event handler |
| 0x0068AF8D | `HandleSelectCity` | Known | Event handler |
| 0x0068AFF9 | `HandleHighlightCity` | Known | Event handler |
| 0x0068B032 | `HandleSelectCity` | Known | Event handler |
| 0x0068B09E | `HandleHighlightCity` | Known | Event handler |
| 0x0068B0D7 | `HandleSelectCity` | Known | Event handler |
| 0x0068B143 | `HandleHighlightCity` | Known | Event handler |
| 0x0068B17C | `HandleSelectCity` | Known | Event handler |
| 0x0068B1E8 | `HandleHighlightCity` | Known | Event handler |
| 0x0068B221 | `HandleSelectCity` | Known | Event handler |
| 0x0068B28D | `HandleHighlightCity` | Known | Event handler |
| 0x0068B2C6 | `HandleSelectCity` | Known | Event handler |
| 0x0068B332 | `HandleHighlightCity` | Known | Event handler |
| 0x0068B36B | `HandleSelectCity` | Known | Event handler |
| 0x0068B3D7 | `HandleHighlightCity` | Known | Event handler |
| 0x0068B410 | `HandleSelectCity` | Known | Event handler |
| 0x0068B47C | `HandleHighlightCity` | Known | Event handler |
| 0x0068B4B5 | `HandleSelectCity` | Known | Event handler |
| 0x0068B521 | `HandleHighlightCity` | Known | Event handler |
| 0x0068B55A | `HandleSelectCity` | Known | Event handler |
| 0x0068B5C6 | `HandleHighlightCity` | Known | Event handler |
| 0x0068B5FF | `HandleSelectCity` | Known | Event handler |
| 0x0068B66B | `HandleHighlightCity` | Known | Event handler |
| 0x0068B6A4 | `HandleSelectCity` | Known | Event handler |
| 0x0068B710 | `HandleHighlightCity` | Known | Event handler |
| 0x0068B749 | `HandleSelectCity` | Known | Event handler |
| 0x0068B7B5 | `HandleHighlightCity` | Known | Event handler |
| 0x0068B7EE | `HandleSelectCity` | Known | Event handler |
| 0x0068B85A | `HandleHighlightCity` | Known | Event handler |
| 0x0068B89A | `HandleSelectCity` | Known | Event handler |
| 0x0068B906 | `HandleHighlightCity` | Known | Event handler |
| 0x0068B93F | `HandleSelectCity` | Known | Event handler |
| 0x0068B9AB | `HandleHighlightCity` | Known | Event handler |
| 0x0068B9E4 | `HandleSelectCity` | Known | Event handler |
| 0x0068BA50 | `HandleHighlightCity` | Known | Event handler |
| 0x0068BA8E | `HandleSelectCity` | Known | Event handler |
| 0x0068BAFA | `HandleHighlightCity` | Known | Event handler |
| 0x0068BB33 | `HandleSelectCity` | Known | Event handler |
| 0x0068BB9F | `HandleHighlightCity` | Known | Event handler |
| 0x0068BBD8 | `HandleSelectCity` | Known | Event handler |
| 0x0068BC44 | `HandleHighlightCity` | Known | Event handler |
| 0x0068BC7D | `HandleSelectCity` | Known | Event handler |
| 0x0068BCE9 | `HandleHighlightCity` | Known | Event handler |
| 0x0068BD22 | `HandleSelectCity` | Known | Event handler |
| 0x0068BD8E | `HandleHighlightCity` | Known | Event handler |
| 0x0068BDC7 | `HandleSelectCity` | Known | Event handler |
| 0x0068BE33 | `HandleHighlightCity` | Known | Event handler |
| 0x0068BE6C | `HandleSelectCity` | Known | Event handler |
| 0x0068BED8 | `HandleHighlightCity` | Known | Event handler |
| 0x0068BF11 | `HandleSelectCity` | Known | Event handler |
| 0x0068BF7D | `HandleHighlightCity` | Known | Event handler |
| 0x0068BFBA | `HandleSelectCity` | Known | Event handler |
| 0x0068C026 | `HandleHighlightCity` | Known | Event handler |
| 0x0068C05F | `HandleSelectCity` | Known | Event handler |
| 0x0068C0CB | `HandleHighlightCity` | Known | Event handler |
| 0x0068C104 | `HandleSelectCity` | Known | Event handler |
| 0x0068C170 | `HandleHighlightCity` | Known | Event handler |
| 0x0068C1A9 | `HandleSelectCity` | Known | Event handler |
| 0x0068C215 | `HandleHighlightCity` | Known | Event handler |
| 0x0068C24E | `HandleSelectCity` | Known | Event handler |
| 0x0068C2BA | `HandleHighlightCity` | Known | Event handler |
| 0x0068C2F3 | `HandleSelectCity` | Known | Event handler |
| 0x0068C35F | `HandleHighlightCity` | Known | Event handler |
| 0x0068C398 | `HandleSelectCity` | Known | Event handler |
| 0x0068C404 | `HandleHighlightCity` | Known | Event handler |
| 0x0068C43D | `HandleSelectCity` | Known | Event handler |
| 0x0068C4A9 | `HandleHighlightCity` | Known | Event handler |
| 0x0068C4E2 | `HandleSelectCity` | Known | Event handler |
| 0x0068C54E | `HandleHighlightCity` | Known | Event handler |
| 0x0068C587 | `HandleSelectCity` | Known | Event handler |
| 0x0068C5F3 | `HandleHighlightCity` | Known | Event handler |
| 0x0068C62C | `HandleSelectCity` | Known | Event handler |
| 0x0068C698 | `HandleHighlightCity` | Known | Event handler |
| 0x0068C6D1 | `HandleSelectCity` | Known | Event handler |
| 0x0068C73D | `HandleHighlightCity` | Known | Event handler |
| 0x0068C776 | `HandleSelectCity` | Known | Event handler |
| 0x0068C7E2 | `HandleHighlightCity` | Known | Event handler |
| 0x0068C81B | `HandleSelectCity` | Known | Event handler |
| 0x0068C887 | `HandleHighlightCity` | Known | Event handler |
| 0x0068C8C0 | `HandleSelectCity` | Known | Event handler |
| 0x0068C92C | `HandleHighlightCity` | Known | Event handler |
| 0x0068C965 | `HandleSelectCity` | Known | Event handler |
| 0x0068C9D1 | `HandleHighlightCity` | Known | Event handler |
| 0x0068CA0A | `HandleSelectCity` | Known | Event handler |
| 0x0068CA76 | `HandleHighlightCity` | Known | Event handler |
| 0x0068CAAF | `HandleSelectCity` | Known | Event handler |
| 0x0068CB1B | `HandleHighlightCity` | Known | Event handler |
| 0x0068CB54 | `HandleSelectCity` | Known | Event handler |
| 0x0068CBC0 | `HandleHighlightCity` | Known | Event handler |
| 0x0068CBF9 | `HandleSelectCity` | Known | Event handler |
| 0x0068CC65 | `HandleHighlightCity` | Known | Event handler |
| 0x0068CC9E | `HandleSelectCity` | Known | Event handler |
| 0x0068CD0A | `HandleHighlightCity` | Known | Event handler |
| 0x0068CD43 | `HandleSelectCity` | Known | Event handler |
| 0x0068CDAF | `HandleHighlightCity` | Known | Event handler |
| 0x0068CDE8 | `HandleSelectCity` | Known | Event handler |
| 0x0068CE54 | `HandleHighlightCity` | Known | Event handler |
| 0x0068CE8D | `HandleSelectCity` | Known | Event handler |
| 0x0068CEF9 | `HandleHighlightCity` | Known | Event handler |
| 0x0068CF32 | `HandleSelectCity` | Known | Event handler |
| 0x0068CF9E | `HandleHighlightCity` | Known | Event handler |
| 0x0068CFD7 | `HandleSelectCity` | Known | Event handler |
| 0x0068D043 | `HandleHighlightCity` | Known | Event handler |
| 0x0068D07C | `HandleSelectCity` | Known | Event handler |
| 0x0068D0E8 | `HandleHighlightCity` | Known | Event handler |
| 0x0068D121 | `HandleSelectCity` | Known | Event handler |
| 0x0068D18D | `HandleHighlightCity` | Known | Event handler |
| 0x0068D1C6 | `HandleSelectCity` | Known | Event handler |
| 0x0068D232 | `HandleHighlightCity` | Known | Event handler |
| 0x0068D26B | `HandleSelectCity` | Known | Event handler |
| 0x0068D2D7 | `HandleHighlightCity` | Known | Event handler |
| 0x0068D310 | `HandleSelectCity` | Known | Event handler |
| 0x0068D37C | `HandleHighlightCity` | Known | Event handler |
| 0x0068D3B5 | `HandleSelectCity` | Known | Event handler |
| 0x0068D421 | `HandleHighlightCity` | Known | Event handler |
| 0x0068D45A | `HandleSelectCity` | Known | Event handler |
| 0x0068D4C6 | `HandleHighlightCity` | Known | Event handler |
| 0x0068D4FF | `HandleSelectCity` | Known | Event handler |
| 0x0068D56B | `HandleHighlightCity` | Known | Event handler |
| 0x0068D5AA | `HandleSelectCity` | Known | Event handler |
| 0x0068D616 | `HandleHighlightCity` | Known | Event handler |
| 0x0068D64F | `HandleSelectCity` | Known | Event handler |
| 0x0068D6BB | `HandleHighlightCity` | Known | Event handler |
| 0x0068D6F4 | `HandleSelectCity` | Known | Event handler |
| 0x0068D760 | `HandleHighlightCity` | Known | Event handler |
| 0x0068D799 | `HandleSelectCity` | Known | Event handler |
| 0x0068D805 | `HandleHighlightCity` | Known | Event handler |
| 0x0068D83E | `HandleSelectCity` | Known | Event handler |
| 0x0068D8AA | `HandleHighlightCity` | Known | Event handler |
| 0x0068D8E3 | `HandleSelectCity` | Known | Event handler |
| 0x0068D94F | `HandleHighlightCity` | Known | Event handler |
| 0x0068D988 | `HandleSelectCity` | Known | Event handler |
| 0x0068D9F4 | `HandleHighlightCity` | Known | Event handler |
| 0x0068DA2D | `HandleSelectCity` | Known | Event handler |
| 0x0068DA99 | `HandleHighlightCity` | Known | Event handler |
| 0x0068DAD2 | `HandleSelectCity` | Known | Event handler |
| 0x0068DB3E | `HandleHighlightCity` | Known | Event handler |
| 0x0068DB77 | `HandleSelectCity` | Known | Event handler |
| 0x0068DBE3 | `HandleHighlightCity` | Known | Event handler |
| 0x0068DC1C | `HandleSelectCity` | Known | Event handler |
| 0x0068DC88 | `HandleHighlightCity` | Known | Event handler |
| 0x0068DCC1 | `HandleSelectCity` | Known | Event handler |
| 0x0068DD2D | `HandleHighlightCity` | Known | Event handler |
| 0x0068DD66 | `HandleSelectCity` | Known | Event handler |
| 0x0068DDD2 | `HandleHighlightCity` | Known | Event handler |
| 0x0068DE0B | `HandleSelectCity` | Known | Event handler |
| 0x0068DE77 | `HandleHighlightCity` | Known | Event handler |
| 0x0068DEB0 | `HandleSelectCity` | Known | Event handler |
| 0x0068DF1C | `HandleHighlightCity` | Known | Event handler |
| 0x0068DF55 | `HandleSelectCity` | Known | Event handler |
| 0x0068DFC1 | `HandleHighlightCity` | Known | Event handler |
| 0x0068DFFA | `HandleSelectCity` | Known | Event handler |
| 0x0068E066 | `HandleHighlightCity` | Known | Event handler |
| 0x0068E09F | `HandleSelectCity` | Known | Event handler |
| 0x0068E10B | `HandleHighlightCity` | Known | Event handler |
| 0x0068E144 | `HandleSelectCity` | Known | Event handler |
| 0x0068E1B0 | `HandleHighlightCity` | Known | Event handler |
| 0x0068E1E9 | `HandleSelectCity` | Known | Event handler |
| 0x0068E255 | `HandleHighlightCity` | Known | Event handler |
| 0x0068E28E | `HandleSelectCity` | Known | Event handler |
| 0x0068E2FA | `HandleHighlightCity` | Known | Event handler |
| 0x0068E333 | `HandleSelectCity` | Known | Event handler |
| 0x0068E39F | `HandleHighlightCity` | Known | Event handler |
| 0x0068E3D8 | `HandleSelectCity` | Known | Event handler |
| 0x0068E444 | `HandleHighlightCity` | Known | Event handler |
| 0x0068E47D | `HandleSelectCity` | Known | Event handler |
| 0x0068E4E9 | `HandleHighlightCity` | Known | Event handler |
| 0x0068E522 | `HandleSelectCity` | Known | Event handler |
| 0x0068E58E | `HandleHighlightCity` | Known | Event handler |
| 0x0068E5C7 | `HandleSelectCity` | Known | Event handler |
| 0x0068E633 | `HandleHighlightCity` | Known | Event handler |
| 0x0068E66C | `HandleSelectCity` | Known | Event handler |
| 0x0068E6D8 | `HandleHighlightCity` | Known | Event handler |
| 0x0068E711 | `HandleSelectCity` | Known | Event handler |
| 0x0068E77D | `HandleHighlightCity` | Known | Event handler |
| 0x0068E7B6 | `HandleSelectCity` | Known | Event handler |
| 0x0068E822 | `HandleHighlightCity` | Known | Event handler |
| 0x0068E85B | `HandleSelectCity` | Known | Event handler |
| 0x0068E8C7 | `HandleHighlightCity` | Known | Event handler |
| 0x0068E900 | `HandleSelectCity` | Known | Event handler |
| 0x0068E96C | `HandleHighlightCity` | Known | Event handler |
| 0x0068E9A5 | `HandleSelectCity` | Known | Event handler |
| 0x0068EA11 | `HandleHighlightCity` | Known | Event handler |
| 0x0068EA4A | `HandleSelectCity` | Known | Event handler |
| 0x0068EAB6 | `HandleHighlightCity` | Known | Event handler |
| 0x0068EAEF | `HandleSelectCity` | Known | Event handler |
| 0x0068EB5B | `HandleHighlightCity` | Known | Event handler |
| 0x0068EB94 | `HandleSelectCity` | Known | Event handler |
| 0x0068EC00 | `HandleHighlightCity` | Known | Event handler |
| 0x0068EC39 | `HandleSelectCity` | Known | Event handler |
| 0x0068ECA5 | `HandleHighlightCity` | Known | Event handler |
| 0x0068ECDE | `HandleSelectCity` | Known | Event handler |
| 0x0068ED4A | `HandleHighlightCity` | Known | Event handler |
| 0x0068ED83 | `HandleSelectCity` | Known | Event handler |
| 0x0068EDEF | `HandleHighlightCity` | Known | Event handler |
| 0x0068EE28 | `HandleSelectCity` | Known | Event handler |
| 0x0068EE94 | `HandleHighlightCity` | Known | Event handler |
| 0x0068EECD | `HandleSelectCity` | Known | Event handler |
| 0x0068EF39 | `HandleHighlightCity` | Known | Event handler |
| 0x0068EF72 | `HandleSelectCity` | Known | Event handler |
| 0x0068EFDE | `HandleHighlightCity` | Known | Event handler |
| 0x0068F017 | `HandleSelectCity` | Known | Event handler |
| 0x0068F083 | `HandleHighlightCity` | Known | Event handler |
| 0x0068F0BC | `HandleSelectCity` | Known | Event handler |
| 0x0068F128 | `HandleHighlightCity` | Known | Event handler |
| 0x0068F161 | `HandleSelectCity` | Known | Event handler |
| 0x0068F1CD | `HandleHighlightCity` | Known | Event handler |
| 0x0068F206 | `HandleSelectCity` | Known | Event handler |
| 0x0068F272 | `HandleHighlightCity` | Known | Event handler |
| 0x0068F2AB | `HandleSelectCity` | Known | Event handler |
| 0x0068F317 | `HandleHighlightCity` | Known | Event handler |
| 0x0068F350 | `HandleSelectCity` | Known | Event handler |
| 0x0068F3BC | `HandleHighlightCity` | Known | Event handler |
| 0x0068F3F5 | `HandleSelectCity` | Known | Event handler |
| 0x0068F461 | `HandleHighlightCity` | Known | Event handler |
| 0x0068F49A | `HandleSelectCity` | Known | Event handler |
| 0x0068F506 | `HandleHighlightCity` | Known | Event handler |
| 0x0068F53F | `HandleSelectCity` | Known | Event handler |
| 0x0068F5AB | `HandleHighlightCity` | Known | Event handler |
| 0x0068F5EA | `HandleSelectCity` | Known | Event handler |
| 0x0068F656 | `HandleHighlightCity` | Known | Event handler |
| 0x0068F68F | `HandleSelectCity` | Known | Event handler |
| 0x0068F6FB | `HandleHighlightCity` | Known | Event handler |
| 0x0068F734 | `HandleSelectCity` | Known | Event handler |
| 0x0068F7A0 | `HandleHighlightCity` | Known | Event handler |
| 0x0068F7D9 | `HandleSelectCity` | Known | Event handler |
| 0x0068F845 | `HandleHighlightCity` | Known | Event handler |
| 0x0068F87E | `HandleSelectCity` | Known | Event handler |
| 0x0068F8EA | `HandleHighlightCity` | Known | Event handler |
| 0x0068F92A | `HandleSelectCity` | Known | Event handler |
| 0x0068F996 | `HandleHighlightCity` | Known | Event handler |
| 0x0068F9CF | `HandleSelectCity` | Known | Event handler |
| 0x0068FA3B | `HandleHighlightCity` | Known | Event handler |
| 0x0068FA74 | `HandleSelectCity` | Known | Event handler |
| 0x0068FAE0 | `HandleHighlightCity` | Known | Event handler |
| 0x0068FB19 | `HandleSelectCity` | Known | Event handler |
| 0x0068FB85 | `HandleHighlightCity` | Known | Event handler |
| 0x0068FBBE | `HandleSelectCity` | Known | Event handler |
| 0x0068FC2A | `HandleHighlightCity` | Known | Event handler |
| 0x0068FC63 | `HandleSelectCity` | Known | Event handler |
| 0x0068FCCF | `HandleHighlightCity` | Known | Event handler |
| 0x0068FD08 | `HandleSelectCity` | Known | Event handler |
| 0x0068FD74 | `HandleHighlightCity` | Known | Event handler |
| 0x0068FDAD | `HandleSelectCity` | Known | Event handler |
| 0x0068FE19 | `HandleHighlightCity` | Known | Event handler |
| 0x0068FE52 | `HandleSelectCity` | Known | Event handler |
| 0x0068FEBE | `HandleHighlightCity` | Known | Event handler |
| 0x0068FEF7 | `HandleSelectCity` | Known | Event handler |
| 0x0068FF63 | `HandleHighlightCity` | Known | Event handler |
| 0x0068FF9C | `HandleSelectCity` | Known | Event handler |
| 0x00690008 | `HandleHighlightCity` | Known | Event handler |
| 0x00690041 | `HandleSelectCity` | Known | Event handler |
| 0x006900AD | `HandleHighlightCity` | Known | Event handler |
| 0x006900E6 | `HandleSelectCity` | Known | Event handler |
| 0x00690152 | `HandleHighlightCity` | Known | Event handler |
| 0x0069018B | `HandleSelectCity` | Known | Event handler |
| 0x006901F7 | `HandleHighlightCity` | Known | Event handler |
| 0x00690230 | `HandleSelectCity` | Known | Event handler |
| 0x0069029C | `HandleHighlightCity` | Known | Event handler |
| 0x006902D5 | `HandleSelectCity` | Known | Event handler |
| 0x00690341 | `HandleHighlightCity` | Known | Event handler |
| 0x0069037A | `HandleSelectCity` | Known | Event handler |
| 0x006903E6 | `HandleHighlightCity` | Known | Event handler |
| 0x006908DE | `HandleMusicSelected` | Known | Event handler |
| 0x00690920 | `HandleMusicHilited` | Known | Event handler |
| 0x00690958 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x0069099E | `HandleMusicHilited` | Known | Event handler |
| 0x006909D6 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00690A1C | `HandlePlaylistsHilited` | Known | Event handler |
| 0x00690A58 | `HandleArtistsSelected` | Known | Event handler |
| 0x00690A9C | `HandleArtistsHilited` | Known | Event handler |
| 0x00690AD6 | `HandleAlbumsSelected` | Known | Event handler |
| 0x00690B19 | `HandleAlbumsHilited` | Known | Event handler |
| 0x00690B52 | `HandleCompilationsSelected` | Known | Event handler |
| 0x00690B9B | `HandleCompilationsHilited` | Known | Event handler |
| 0x00690BDA | `HandleSongsSelected` | Known | Event handler |
| 0x00690C1C | `HandleSongsHilited` | Known | Event handler |
| 0x00690C54 | `HandleGenresSelected` | Known | Event handler |
| 0x00690C97 | `HandleGenresHilited` | Known | Event handler |
| 0x00690CD0 | `HandleComposersSelected` | Known | Event handler |
| 0x00690D16 | `HandleComposersHilited` | Known | Event handler |
| 0x00690D52 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x00690D99 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00690E58 | `HandleMusicHilited` | Known | Event handler |
| 0x00690E90 | `HandleVideosSelected` | Known | Event handler |
| 0x00690ED3 | `HandleVideosHilited` | Known | Event handler |
| 0x00690F0C | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00690F57 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00690F98 | `HandleMoviesSelected` | Known | Event handler |
| 0x00690FDB | `HandleMoviesHilited` | Known | Event handler |
| 0x00691014 | `HandleTVShowsSelected` | Known | Event handler |
| 0x00691058 | `HandleTVShowsHilited` | Known | Event handler |
| 0x00691092 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x006910DA | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00691118 | `HandlePhotosSelected` | Known | Event handler |
| 0x0069115B | `HandlePhotosHilited` | Known | Event handler |
| 0x00691194 | `HandlePhotosSelected` | Known | Event handler |
| 0x006911D7 | `HandlePhotosHilited` | Known | Event handler |
| 0x00691210 | `HandlePodcastsSelected` | Known | Event handler |
| 0x00691255 | `HandlePodcastsHilited` | Known | Event handler |
| 0x00691308 | `HandleGenericHilited` | Known | Event handler |
| 0x00691401 | `HandleGenericHilited` | Known | Event handler |
| 0x006918E6 | `HandleLock` | Known | Event handler |
| 0x00691A57 | `HandleNikePlusSelected` | Known | Event handler |
| 0x00691A9C | `HandleGenericHilited` | Known | Event handler |
| 0x00691BA2 | `HandleGenericHilited` | Known | Event handler |
| 0x00691CA1 | `HandleGenericHilited` | Known | Event handler |
| 0x00691D8E | `HandleGenericHilited` | Known | Event handler |
| 0x00691E08 | `HandleRadioPlayPause` | Known | Event handler |
| 0x00691EC8 | `HandleGenericHilited` | Known | Event handler |
| 0x00691F42 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00691F8B | `HandleGenericHilited` | Known | Event handler |
| 0x00692004 | `HandleBacklightSelected` | Known | Event handler |
| 0x0069204A | `HandleGenericHilited` | Known | Event handler |
| 0x006920C5 | `HandleSleepSelected` | Known | Event handler |
| 0x00692107 | `HandleGenericHilited` | Known | Event handler |
| 0x0069217E | `HandleNowPlaying` | Known | Event handler |
| 0x006921F6 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x0069223A | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00692280 | `HandleMusicHilited` | Known | Event handler |
| 0x006922B8 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x006922FE | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x0069233C | `HandleArtistsSelected` | Known | Event handler |
| 0x00692380 | `HandleArtistsHilited` | Known | Event handler |
| 0x006923BA | `HandleAlbumsSelected` | Known | Event handler |
| 0x006923FD | `HandleAlbumsHilited` | Known | Event handler |
| 0x00692436 | `HandleCompilationsSelected` | Known | Event handler |
| 0x0069247F | `HandleCompilationsHilited` | Known | Event handler |
| 0x006924BE | `HandleSongsSelected` | Known | Event handler |
| 0x00692500 | `HandleSongsHilited` | Known | Event handler |
| 0x006925AB | `HandleGenericHilited` | Known | Event handler |
| 0x00692626 | `HandleRadioPlayPause` | Known | Event handler |
| 0x00692660 | `HandleGenresSelected` | Known | Event handler |
| 0x006926A3 | `HandleGenresHilited` | Known | Event handler |
| 0x006926DC | `HandleComposersSelected` | Known | Event handler |
| 0x00692722 | `HandleComposersHilited` | Known | Event handler |
| 0x0069275E | `HandleAudiobooksSelected` | Known | Event handler |
| 0x006927A5 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00692864 | `HandleMusicHilited` | Known | Event handler |
| 0x006928D9 | `HandlePlayPause` | Known | Event handler |
| 0x0069290E | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x006929F8 | `HandleSelect` | Known | Event handler |
| 0x00692A3A | `HandleMoviesSelected` | Known | Event handler |
| 0x00692A7D | `HandleMoviesHilited` | Known | Event handler |
| 0x00692AB6 | `HandleTVShowsSelected` | Known | Event handler |
| 0x00692AFA | `HandleTVShowsHilited` | Known | Event handler |
| 0x00692B34 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00692B7C | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00692BBA | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00692C05 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00692CCB | `HandleVideosHilited` | Known | Event handler |
| 0x00693359 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x00693EE2 | `HandleMainMenu` | Known | Event handler |
| 0x00693F1A | `HandleMusicMenu` | Known | Event handler |
| 0x00694442 | `HandleRadioRegion` | Known | Event handler |
| 0x006944E6 | `HandleLanguage` | Known | Event handler |
| 0x006945EC | `HandleNew` | Known | Event handler |
| 0x00694667 | `HandleClear` | Known | Event handler |
| 0x00694698 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x00694754 | `HandleSelectIndexedSession` | Known | Event handler |
| 0x006948AD | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x00694900 | `HandleSelect` | Known | Event handler |
| 0x00694A2A | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x00694A64 | `HandleEQSettingSelected` | Known | Event handler |
| 0x00694A9C | `HandleEQSettingSelected` | Known | Event handler |
| 0x006A6B8C | `HandleItemSelected` | Known | Event handler |
| 0x006A6CD7 | `HandleNextContact` | Known | Event handler |
| 0x006A6D03 | `HandlePreviousContact` | Known | Event handler |
| 0x006A6D39 | `HandleSelectKey` | Known | Event handler |
| 0x006A734A | `HandleSelect` | Known | Event handler |
| 0x006A7671 | `HandleDateChosen` | Known | Event handler |
| 0x006A76A7 | `HandleTimeChosen` | Known | Event handler |
| 0x006A76DD | `HandleFrequencyChosen` | Known | Event handler |
| 0x006A7718 | `HandleSoundChosen` | Known | Event handler |
| 0x006A774F | `HandleLabelChosen` | Known | Event handler |
| 0x006A7786 | `HandleDeleteChosen` | Known | Event handler |
| 0x006A77C2 | `HandleSelect` | Known | Event handler |
| 0x006A77FA | `HandleSelect` | Known | Event handler |
| 0x006A7B3B | `HandleLeaveAlarm` | Known | Event handler |
| 0x006A7B68 | `HandleLeaveAlarm` | Known | Event handler |
| 0x006A7B97 | `HandleLeaveAlarm` | Known | Event handler |
| 0x006A7BC4 | `HandleLeaveAlarm` | Known | Event handler |
| 0x006A7CFE | `HandleSelect` | Known | Event handler |
| 0x006A7D2C | `HandleSelect` | Known | Event handler |
| 0x006A7E8B | `HandleNextDay` | Known | Event handler |
| 0x006A7EB3 | `HandlePreviousDay` | Known | Event handler |
| 0x006A8062 | `HandleSelect` | Known | Event handler |
| 0x006A808F | `HandleNextDay` | Known | Event handler |
| 0x006A80B7 | `HandlePreviousDay` | Known | Event handler |
| 0x006A825F | `HandleNextDay` | Known | Event handler |
| 0x006A8287 | `HandlePreviousDay` | Known | Event handler |
| 0x006A8348 | `HandleSelect` | Known | Event handler |
| 0x006A8373 | `HandleNextDay` | Known | Event handler |
| 0x006A839B | `HandlePreviousDay` | Known | Event handler |
| 0x006A8512 | `HandleSelectLozinch` | Known | Event handler |
| 0x006A868A | `HandleSelectLozinch` | Known | Event handler |
| 0x006A87A9 | `HandleFlowNext` | Known | Event handler |
| 0x006A87D7 | `HandlePlayPause` | Known | Event handler |
| 0x006A8826 | `HandleFlowPrev` | Known | Event handler |
| 0x006A8851 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x006A8945 | `HandleAlbumSelected` | Known | Event handler |
| 0x006A8AE0 | `HandleFlowNext` | Known | Event handler |
| 0x006A8B2E | `HandleFlowNext` | Known | Event handler |
| 0x006A8B5C | `HandlePlayPause` | Known | Event handler |
| 0x006A8BAB | `HandleFlowPrev` | Known | Event handler |
| 0x006A8BD7 | `HandleFlowPrev` | Known | Event handler |
| 0x006A8BF7 | `HandleFlowWheel` | Known | Event handler |
| 0x006A8F87 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x006A93B2 | `HandleArrowDown` | Known | Event handler |
| 0x006A941C | `HandleArrowUp` | Known | Event handler |
| 0x006A943B | `HandleWheel` | Known | Event handler |
| 0x006A94C4 | `HandleSelect` | Known | Event handler |
| 0x006A9541 | `HandleGameHilited` | Known | Event handler |
| 0x006AC9A3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006AE477 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006AFF4B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006B1A1F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006B34F3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006B4FC7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006B6A9B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006B856F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006BA043 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006BBB17 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006BD5EB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006BF0BF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006C0B93 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006C2667 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006C413B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006C5C0F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006C76E3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006C91B7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006CAC8B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006CC75F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006CE233 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006CFD07 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006D17DB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006D32AF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006D4D83 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006D6857 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006D832B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006D9DFF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006DB8D3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006DD3A7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006DEE7B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006E094F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006E2423 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006E3EF7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006E59CB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006E749F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006E8F58 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006E9AE0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006EA668 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006EB1F0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006EBD78 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006EC900 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006ED488 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006EE010 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006EEB98 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006EF720 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006F02A8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006F0E30 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x006F19B8 | `HandlePlayPause` | Known | Event handler |
| 0x006F19EE | `HandleAddToOTG` | Known | Event handler |
| 0x006F1B8B | `HandlePlayPause` | Known | Event handler |
| 0x006F1BB2 | `HandleSelect` | Known | Event handler |
| 0x006F1BDF | `HandleHilite` | Known | Event handler |
| 0x006F1C10 | `HandlePlayPause` | Known | Event handler |
| 0x006F1CA3 | `HandlePlayPause` | Known | Event handler |
| 0x006F1CCA | `HandleSelect` | Known | Event handler |
| 0x006F1D30 | `HandleHilite` | Known | Event handler |
| 0x006F1D62 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x006F1DAC | `HandlePlayPause` | Known | Event handler |
| 0x006F1DE2 | `HandleAddToOTG` | Known | Event handler |
| 0x006F1E74 | `HandlePlayPause` | Known | Event handler |
| 0x006F1E9B | `HandleSelect` | Known | Event handler |
| 0x006F1F04 | `HandlePlayPause` | Known | Event handler |
| 0x006F1F3A | `HandleAddToOTG` | Known | Event handler |
| 0x006F1FCC | `HandlePlayPause` | Known | Event handler |
| 0x006F1FF3 | `HandleSelect` | Known | Event handler |
| 0x006F205C | `HandlePlayPause` | Known | Event handler |
| 0x006F20E2 | `HandleSelect` | Known | Event handler |
| 0x006F2147 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F2188 | `HandlePlayPause` | Known | Event handler |
| 0x006F21BE | `HandleAddToOTG` | Known | Event handler |
| 0x006F23F0 | `HandlePlayPause` | Known | Event handler |
| 0x006F2417 | `HandleSelect` | Known | Event handler |
| 0x006F2444 | `HandleHilite` | Known | Event handler |
| 0x006F2474 | `HandlePlayPause` | Known | Event handler |
| 0x006F24AA | `HandleAddToOTG` | Known | Event handler |
| 0x006F26DC | `HandlePlayPause` | Known | Event handler |
| 0x006F2703 | `HandleSelect` | Known | Event handler |
| 0x006F2730 | `HandleHilite` | Known | Event handler |
| 0x006F2760 | `HandlePlayPause` | Known | Event handler |
| 0x006F2796 | `HandleAddToOTG` | Known | Event handler |
| 0x006F2A81 | `HandlePlayPause` | Known | Event handler |
| 0x006F2AA8 | `HandleSelect` | Known | Event handler |
| 0x006F2AD8 | `HandlePlayPause` | Known | Event handler |
| 0x006F2B0E | `HandleAddToOTG` | Known | Event handler |
| 0x006F2BA0 | `HandlePlayPause` | Known | Event handler |
| 0x006F2BC7 | `HandleSelect` | Known | Event handler |
| 0x006F2C58 | `HandlePlayPause` | Known | Event handler |
| 0x006F2C8E | `HandleAddToOTG` | Known | Event handler |
| 0x006F2E47 | `HandlePlayPause` | Known | Event handler |
| 0x006F2E6E | `HandleSelect` | Known | Event handler |
| 0x006F2EA0 | `HandlePlayPause` | Known | Event handler |
| 0x006F2ED6 | `HandleAddToOTG` | Known | Event handler |
| 0x006F2F5B | `HandleSelect` | Known | Event handler |
| 0x006F2FF4 | `HandleHilite` | Known | Event handler |
| 0x006F3020 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F3064 | `HandlePlayPause` | Known | Event handler |
| 0x006F309A | `HandleAddToOTG` | Known | Event handler |
| 0x006F311F | `HandleSelect` | Known | Event handler |
| 0x006F3184 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F31C8 | `HandlePlayPause` | Known | Event handler |
| 0x006F336C | `HandleSelect` | Known | Event handler |
| 0x006F3399 | `HandleHilite` | Known | Event handler |
| 0x006F33C5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F3408 | `HandlePlayPause` | Known | Event handler |
| 0x006F348E | `HandleSelect` | Known | Event handler |
| 0x006F351C | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F3560 | `HandlePlayPause` | Known | Event handler |
| 0x006F35E6 | `HandleSelect` | Known | Event handler |
| 0x006F364B | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F368C | `HandlePlayPause` | Known | Event handler |
| 0x006F3712 | `HandleSelect` | Known | Event handler |
| 0x006F3778 | `HandleHilite` | Known | Event handler |
| 0x006F37A4 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F37E8 | `HandlePlayPause` | Known | Event handler |
| 0x006F381E | `HandleAddToOTG` | Known | Event handler |
| 0x006F39E1 | `HandlePlayPause` | Known | Event handler |
| 0x006F3A08 | `HandleSelect` | Known | Event handler |
| 0x006F3A38 | `HandlePlayPause` | Known | Event handler |
| 0x006F3A6E | `HandleAddToOTG` | Known | Event handler |
| 0x006F3C8F | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x006F3DA8 | `HandlePlayPause` | Known | Event handler |
| 0x006F3ED5 | `HandleSelect` | Known | Event handler |
| 0x006F3F01 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F3F44 | `HandlePlayPause` | Known | Event handler |
| 0x006F4077 | `HandleSelect` | Known | Event handler |
| 0x006F40A3 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F4871 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F4FE5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F5759 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F5ECD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F6641 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F6DB5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F7529 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x006F7572 | `HandleTVOutChanged` | Known | Event handler |
| 0x006F75AA | `HandleTVSignalChanged` | Known | Event handler |
| 0x006F75E5 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x006F7636 | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x006F767B | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x006F76BD | `HandleSelect` | Known | Event handler |
| 0x006F76ED | `HandleSelect` | Known | Event handler |
| 0x006F7779 | `HandlePlayPause` | Known | Event handler |
| 0x006F77F9 | `HandleSelect` | Known | Event handler |
| 0x006F7F5C | `HandlePlayPause` | Known | Event handler |
| 0x006F7FD1 | `HandleWheelProgress` | Known | Event handler |
| 0x006F8061 | `HandlePlayPause` | Known | Event handler |
| 0x006F80E1 | `HandleSelectProgress` | Known | Event handler |
| 0x006F884C | `HandlePlayPause` | Known | Event handler |
| 0x006F88C1 | `HandleWheelProgress` | Known | Event handler |
| 0x006F8951 | `HandlePlayPause` | Known | Event handler |
| 0x006F89D1 | `HandleSelectVolume` | Known | Event handler |
| 0x006F913A | `HandlePlayPause` | Known | Event handler |
| 0x006F91AF | `HandleWheelVolume` | Known | Event handler |
| 0x006F923D | `HandlePlayPause` | Known | Event handler |
| 0x006F92BD | `HandleSelectRating` | Known | Event handler |
| 0x006F9A26 | `HandlePlayPause` | Known | Event handler |
| 0x006F9A9B | `HandleWheelRating` | Known | Event handler |
| 0x006F9B1B | `HandlePlayPause` | Known | Event handler |
| 0x006F9B92 | `HandleSelectScrub` | Known | Event handler |
| 0x006FA2EC | `HandlePlayPause` | Known | Event handler |
| 0x006FA358 | `HandleWheelScrub` | Known | Event handler |
| 0x006FA3BC | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x006FA3F4 | `HandlePlayPause` | Known | Event handler |
| 0x006FA44E | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x006FA483 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x006FABF3 | `HandlePlayPause` | Known | Event handler |
| 0x006FAC68 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x006FACFD | `HandlePlayPause` | Known | Event handler |
| 0x006FAD7D | `HandleSelectExtraInfo` | Known | Event handler |
| 0x006FB4E9 | `HandlePlayPause` | Known | Event handler |
| 0x006FB5DD | `HandlePlayPause` | Known | Event handler |
| 0x006FB65D | `HandleSelectChapterArt` | Known | Event handler |
| 0x006FBDCA | `HandlePlayPause` | Known | Event handler |
| 0x006FBE3F | `HandleWheelVolume` | Known | Event handler |
| 0x006FBED6 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x006FBF6D | `HandleSelect` | Known | Event handler |
| 0x006FC6D9 | `HandlePlayPause` | Known | Event handler |
| 0x006FC757 | `HandleWheel` | Known | Event handler |
| 0x006FC7EA | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x006FC881 | `HandleSelect` | Known | Event handler |
| 0x006FCFED | `HandlePlayPause` | Known | Event handler |
| 0x006FD06B | `HandleWheel` | Known | Event handler |
| 0x006FD0F5 | `HandlePlayPause` | Known | Event handler |
| 0x006FD175 | `HandleSelect` | Known | Event handler |
| 0x006FD8D8 | `HandlePlayPause` | Known | Event handler |
| 0x006FD94D | `HandleWheel` | Known | Event handler |
| 0x006FD9D5 | `HandlePlayPause` | Known | Event handler |
| 0x006FDA55 | `HandleSelectProgress` | Known | Event handler |
| 0x006FE1C0 | `HandlePlayPause` | Known | Event handler |
| 0x006FE235 | `HandleWheelProgress` | Known | Event handler |
| 0x006FE2B7 | `HandlePlayPause` | Known | Event handler |
| 0x006FE32E | `HandleSelectScrub` | Known | Event handler |
| 0x006FEA88 | `HandlePlayPause` | Known | Event handler |
| 0x006FEAF4 | `HandleWheelScrub` | Known | Event handler |
| 0x006FEB81 | `HandlePlayPause` | Known | Event handler |
| 0x006FF363 | `HandlePlayPause` | Known | Event handler |
| 0x006FF3D8 | `HandleWheelVolume` | Known | Event handler |
| 0x006FF469 | `HandlePlayPause` | Known | Event handler |
| 0x006FFC4B | `HandlePlayPause` | Known | Event handler |
| 0x006FFCC0 | `HandleWheelBrightness` | Known | Event handler |
| 0x006FFD55 | `HandlePlayPause` | Known | Event handler |
| 0x006FFDD5 | `HandleSelect` | Known | Event handler |
| 0x0070013A | `HandlePlayPause` | Known | Event handler |
| 0x0070021D | `HandlePlayPause` | Known | Event handler |
| 0x0070029D | `HandleSelectProgress` | Known | Event handler |
| 0x0070060A | `HandlePlayPause` | Known | Event handler |
| 0x0070067F | `HandleWheelProgress` | Known | Event handler |
| 0x00700711 | `HandlePlayPause` | Known | Event handler |
| 0x00700791 | `HandleSelectProgress` | Known | Event handler |
| 0x00700A8A | `HandlePlayPause` | Known | Event handler |
| 0x00700AFF | `HandleWheelProgress` | Known | Event handler |
| 0x00700B78 | `HandlePlayPause` | Known | Event handler |
| 0x00700BE4 | `HandleSelectScrub` | Known | Event handler |
| 0x00700EC1 | `HandlePlayPause` | Known | Event handler |
| 0x00700F22 | `HandleWheelScrub` | Known | Event handler |
| 0x00700FB1 | `HandlePlayPause` | Known | Event handler |
| 0x00701031 | `HandleSelectVolume` | Known | Event handler |
| 0x00701328 | `HandlePlayPause` | Known | Event handler |
| 0x0070139D | `HandleWheelVolume` | Known | Event handler |
| 0x007013D1 | `HandleSelect` | Known | Event handler |
| 0x00701409 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0070143C | `HandleNotesPop` | Known | Event handler |
| 0x007014B9 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007014EC | `HandleNotesPop` | Known | Event handler |
| 0x007019A8 | `HandleNotesSelected` | Known | Event handler |
| 0x007019E5 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00701A18 | `HandleNotesPop` | Known | Event handler |
| 0x00701ED4 | `HandleNotesSelected` | Known | Event handler |
| 0x00701F11 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00701F44 | `HandleNotesPop` | Known | Event handler |
| 0x00701F6F | `HandleNotesSelected` | Known | Event handler |
| 0x00702441 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00702474 | `HandleNotesPop` | Known | Event handler |
| 0x0070249F | `HandleNotesSelected` | Known | Event handler |
| 0x00702971 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007029A4 | `HandleNotesPop` | Known | Event handler |
| 0x00702A21 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00702A54 | `HandleNotesPop` | Known | Event handler |
| 0x00702AD1 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00702B04 | `HandleNotesPop` | Known | Event handler |
| 0x00702B7C | `HandlePlayPause` | Known | Event handler |
| 0x00702BA5 | `HandlePlayPause` | Known | Event handler |
| 0x00702BD3 | `HandlePlayPause` | Known | Event handler |
| 0x00702C08 | `HandleBrowseAlbum` | Known | Event handler |
| 0x00702C88 | `HandleHiliteAlbum` | Known | Event handler |
| 0x00702D31 | `HandleBrowseAlbum` | Known | Event handler |
| 0x00702DB8 | `HandleHiliteAlbum` | Known | Event handler |
| 0x0070307C | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x007030D8 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x0070328F | `HandleSelect` | Known | Event handler |
| 0x00703413 | `HandleSelect` | Known | Event handler |
| 0x0070344D | `HandleImageLast` | Known | Event handler |
| 0x00703477 | `HandleImageNext` | Known | Event handler |
| 0x007034A6 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007034E0 | `HandleImageFirst` | Known | Event handler |
| 0x0070350B | `HandleImagePrev` | Known | Event handler |
| 0x00703537 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x00703566 | `HandleImageNext` | Known | Event handler |
| 0x0070358F | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007035C3 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007035F2 | `HandleImagePrev` | Known | Event handler |
| 0x00703613 | `HandleImageWheel` | Known | Event handler |
| 0x007036B1 | `HandleImageNext` | Known | Event handler |
| 0x007036E0 | `HandlePlayPause` | Known | Event handler |
| 0x0070372F | `HandleImagePrev` | Known | Event handler |
| 0x0070375B | `HandleSelect` | Known | Event handler |
| 0x00703A2B | `HandleImageNext` | Known | Event handler |
| 0x00703A55 | `HandlePause` | Known | Event handler |
| 0x00703A7A | `HandlePlay` | Known | Event handler |
| 0x00703AA3 | `HandlePlayPause` | Known | Event handler |
| 0x00703ACC | `HandleImagePrev` | Known | Event handler |
| 0x00703B25 | `HandleWheel` | Known | Event handler |
| 0x00703BBD | `HandleImageNext` | Known | Event handler |
| 0x00703BEC | `HandlePlayPause` | Known | Event handler |
| 0x00703C3B | `HandleImagePrev` | Known | Event handler |
| 0x00703C67 | `HandleSelect` | Known | Event handler |
| 0x00703F37 | `HandleImageNext` | Known | Event handler |
| 0x00703F61 | `HandlePause` | Known | Event handler |
| 0x00703F86 | `HandlePlay` | Known | Event handler |
| 0x00703FAF | `HandlePlayPause` | Known | Event handler |
| 0x00703FD8 | `HandleImagePrev` | Known | Event handler |
| 0x00704031 | `HandleWheel` | Known | Event handler |
| 0x007040C9 | `HandleImageNext` | Known | Event handler |
| 0x007040F8 | `HandlePlayPause` | Known | Event handler |
| 0x00704147 | `HandleImagePrev` | Known | Event handler |
| 0x00704173 | `HandleSelect` | Known | Event handler |
| 0x00704443 | `HandleImageNext` | Known | Event handler |
| 0x0070446D | `HandlePause` | Known | Event handler |
| 0x00704492 | `HandlePlay` | Known | Event handler |
| 0x007044BB | `HandlePlayPause` | Known | Event handler |
| 0x007044E4 | `HandleImagePrev` | Known | Event handler |
| 0x0070453D | `HandleWheel` | Known | Event handler |
| 0x007045D5 | `HandleImageNext` | Known | Event handler |
| 0x00704604 | `HandlePlayPause` | Known | Event handler |
| 0x00704653 | `HandleImagePrev` | Known | Event handler |
| 0x0070467F | `HandleSelect` | Known | Event handler |
| 0x0070494F | `HandleImageNext` | Known | Event handler |
| 0x00704979 | `HandlePause` | Known | Event handler |
| 0x0070499E | `HandlePlay` | Known | Event handler |
| 0x007049C7 | `HandlePlayPause` | Known | Event handler |
| 0x007049F0 | `HandleImagePrev` | Known | Event handler |
| 0x00704A49 | `HandleWheel` | Known | Event handler |
| 0x00704AE1 | `HandleImageNext` | Known | Event handler |
| 0x00704B10 | `HandlePlayPause` | Known | Event handler |
| 0x00704B5F | `HandleImagePrev` | Known | Event handler |
| 0x00704B8B | `HandleSelect` | Known | Event handler |
| 0x00704E5B | `HandleImageNext` | Known | Event handler |
| 0x00704E85 | `HandlePause` | Known | Event handler |
| 0x00704EAA | `HandlePlay` | Known | Event handler |
| 0x00704ED3 | `HandlePlayPause` | Known | Event handler |
| 0x00704EFC | `HandleImagePrev` | Known | Event handler |
| 0x00704F55 | `HandleWheel` | Known | Event handler |
| 0x00704FED | `HandleImageNext` | Known | Event handler |
| 0x0070501C | `HandlePlayPause` | Known | Event handler |
| 0x0070506B | `HandleImagePrev` | Known | Event handler |
| 0x00705097 | `HandleSelect` | Known | Event handler |
| 0x00705367 | `HandleImageNext` | Known | Event handler |
| 0x00705391 | `HandlePause` | Known | Event handler |
| 0x007053B6 | `HandlePlay` | Known | Event handler |
| 0x007053DF | `HandlePlayPause` | Known | Event handler |
| 0x00705408 | `HandleImagePrev` | Known | Event handler |
| 0x00705461 | `HandleWheel` | Known | Event handler |
| 0x007054F9 | `HandleImageNext` | Known | Event handler |
| 0x00705528 | `HandlePlayPause` | Known | Event handler |
| 0x00705577 | `HandleImagePrev` | Known | Event handler |
| 0x007055A3 | `HandleSelect` | Known | Event handler |
| 0x007057EE | `HandleImageNext` | Known | Event handler |
| 0x00705818 | `HandlePause` | Known | Event handler |
| 0x0070583D | `HandlePlay` | Known | Event handler |
| 0x00705866 | `HandlePlayPause` | Known | Event handler |
| 0x0070588F | `HandleImagePrev` | Known | Event handler |
| 0x007058F8 | `HandleWheel` | Known | Event handler |
| 0x00705991 | `HandleImageNext` | Known | Event handler |
| 0x007059C0 | `HandlePlayPause` | Known | Event handler |
| 0x00705A0F | `HandleImagePrev` | Known | Event handler |
| 0x00705A3B | `HandleSelect` | Known | Event handler |
| 0x00705C86 | `HandleImageNext` | Known | Event handler |
| 0x00705CB0 | `HandlePause` | Known | Event handler |
| 0x00705CD5 | `HandlePlay` | Known | Event handler |
| 0x00705CFE | `HandlePlayPause` | Known | Event handler |
| 0x00705D27 | `HandleImagePrev` | Known | Event handler |
| 0x00705D90 | `HandleWheel` | Known | Event handler |
| 0x00705E29 | `HandleImageNext` | Known | Event handler |
| 0x00705E58 | `HandlePlayPause` | Known | Event handler |
| 0x00705EA7 | `HandleImagePrev` | Known | Event handler |
| 0x00705ED3 | `HandleSelect` | Known | Event handler |
| 0x0070611E | `HandleImageNext` | Known | Event handler |
| 0x00706148 | `HandlePause` | Known | Event handler |
| 0x0070616D | `HandlePlay` | Known | Event handler |
| 0x00706196 | `HandlePlayPause` | Known | Event handler |
| 0x007061BF | `HandleImagePrev` | Known | Event handler |
| 0x00706228 | `HandleWheel` | Known | Event handler |
| 0x007062C1 | `HandleImageNext` | Known | Event handler |
| 0x007062F0 | `HandlePlayPause` | Known | Event handler |
| 0x0070633F | `HandleImagePrev` | Known | Event handler |
| 0x0070636B | `HandleSelect` | Known | Event handler |
| 0x007065B6 | `HandleImageNext` | Known | Event handler |
| 0x007065E0 | `HandlePause` | Known | Event handler |
| 0x00706605 | `HandlePlay` | Known | Event handler |
| 0x0070662E | `HandlePlayPause` | Known | Event handler |
| 0x00706657 | `HandleImagePrev` | Known | Event handler |
| 0x007066C0 | `HandleWheel` | Known | Event handler |
| 0x00706759 | `HandleImageNext` | Known | Event handler |
| 0x00706788 | `HandlePlayPause` | Known | Event handler |
| 0x007067D7 | `HandleImagePrev` | Known | Event handler |
| 0x00706803 | `HandleSelect` | Known | Event handler |
| 0x00706A4E | `HandleImageNext` | Known | Event handler |
| 0x00706A78 | `HandlePause` | Known | Event handler |
| 0x00706A9D | `HandlePlay` | Known | Event handler |
| 0x00706AC6 | `HandlePlayPause` | Known | Event handler |
| 0x00706AEF | `HandleImagePrev` | Known | Event handler |
| 0x00706B58 | `HandleWheel` | Known | Event handler |
| 0x00706B85 | `HandleSelect` | Known | Event handler |
| 0x00706BB5 | `HandleSelect` | Known | Event handler |
| 0x00706CC5 | `HandleTuning` | Known | Event handler |
| 0x00706E81 | `HandleVolumeChange` | Known | Event handler |
| 0x00706FBA | `HandleVolumeWheel` | Known | Event handler |
| 0x00707124 | `HandleTimerDone` | Known | Event handler |
| 0x00707381 | `HandleFrequencyChange` | Known | Event handler |
| 0x007074F0 | `HandleTimerDone` | Known | Event handler |
| 0x0070774D | `HandleFrequencyChange` | Known | Event handler |
| 0x00707865 | `HandleTimerDone` | Known | Event handler |
| 0x00707994 | `HandleVolumeChange` | Known | Event handler |
| 0x00707A98 | `HandleVolumeWheel` | Known | Event handler |
| 0x00707CF0 | `HandleRemoteUIRemotePlayPause` | Known | Event handler |
| 0x00707F64 | `HandleRemoteUIRemotePlayPause` | Known | Event handler |
| 0x00707FB3 | `HandleExitUnsupported` | Known | Event handler |
| 0x00707FE5 | `HandleExitUnsupported` | Known | Event handler |
| 0x0070B335 | `HandleSelectKey` | Known | Event handler |
| 0x0070B36A | `HandleWheel` | Known | Event handler |
| 0x0070B4B8 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x0070B50B | `HandleSelectKey` | Known | Event handler |
| 0x0070B533 | `HandleSelectKey` | Known | Event handler |
| 0x0070B563 | `HandleExit` | Known | Event handler |
| 0x0070B58D | `HandleStartStop` | Known | Event handler |
| 0x0070B5F3 | `HandleStartStop` | Known | Event handler |
| 0x0070B70B | `HandleExit` | Known | Event handler |
| 0x0070B735 | `HandleStartStop` | Known | Event handler |
| 0x0070B761 | `HandleLap` | Known | Event handler |
| 0x0070B865 | `HandleSelectLozinch` | Known | Event handler |
| 0x0070B98C | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x0070BC76 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x0070BD57 | `HandlePlayPause` | Known | Event handler |
| 0x0070BDE5 | `HandlePlayPause` | Known | Event handler |
| 0x0070BE75 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x0070BEAD | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x0070BEE9 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x0070BF2C | `HandlePlayPause` | Known | Event handler |
| 0x0070BF62 | `HandleAddToOTG` | Known | Event handler |
| 0x0070C1B7 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x0070C413 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007267EE | `HandleSelectClock` | Known | Event handler |
| 0x00726827 | `HandleHilited` | Known | Event handler |
| 0x00726859 | `HandleWheel` | Known | Event handler |
| 0x007268A0 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00726925 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00726ADD | `HandleImageLast` | Known | Event handler |
| 0x00726B07 | `HandleScreenNext` | Known | Event handler |
| 0x00726B37 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x00726B71 | `HandleImageFirst` | Known | Event handler |
| 0x00726B9C | `HandleScreenPrev` | Known | Event handler |
| 0x00726BC9 | `HandleBrowseLarge` | Known | Event handler |
| 0x00726C49 | `HandleImageNext` | Known | Event handler |
| 0x00726C72 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x00726CA6 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x00726CD5 | `HandleImagePrev` | Known | Event handler |
| 0x00726D03 | `HandleBrowseSmall` | Known | Event handler |

---

## 6. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F1660 | `GotoNowPlaying` | Known | Navigation |
| 0x000F16D8 | `GotoMainMenu` | Known | Navigation |
| 0x001093B8 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x001093D0 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x00109548 | `GotoScreen_AddressBook` | Known | Navigation |
| 0x00113588 | `GotoNowPlaying` | Known | Navigation |
| 0x0011359C | `GotoAlbums` | Known | Navigation |
| 0x001135A8 | `GotoSongs` | Known | Navigation |
| 0x00120DEC | `GotoScreen_EnterPassKey` | Known | Navigation |
| 0x00120E04 | `GotoScreen_LockediPod` | Known | Navigation |
| 0x001217A8 | `GotoScreen_MainMenu` | Known | Navigation |
| 0x001373C8 | `GotoMainMenu` | Known | Navigation |
| 0x001B0F34 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001BB868 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001BC0B8 | `GotoPlayDeleteMenu` | Known | Navigation |
| 0x001BC13C | `GotoNowPlaying` | Known | Navigation |
| 0x001D4694 | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation |
| 0x001DFD54 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x001DFE0C | `GotoScreen_SettingsMenu` | Known | Navigation |
| 0x001E759C | `GotoDefaultLayout` | Known | Navigation |
| 0x001E7620 | `GotoVolumeLayout` | Known | Navigation |
| 0x001E7768 | `GotoProgressLayout` | Known | Navigation |
| 0x001E7A78 | `GotoDefault` | Known | Navigation |
| 0x001E7C8C | `GotoPausedIconLayout` | Known | Navigation |
| 0x001E7E00 | `GotoProgressLayout` | Known | Navigation |
| 0x001E7F54 | `GotoDefaultLayout` | Known | Navigation |
| 0x001E8014 | `GotoDefaultLayout` | Known | Navigation |
| 0x001E80E8 | `GotoProgressLayout` | Known | Navigation |
| 0x001E8274 | `GotoProgressLayout` | Known | Navigation |
| 0x001E9AD4 | `GotoNowPlaying` | Known | Navigation |
| 0x001EA38C | `GotoNowPlaying` | Known | Navigation |
| 0x001ECA18 | `GotoScreen_Language` | Known | Navigation |
| 0x001ECD90 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x001ECDAC | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x001ECDC4 | `GotoDefaultLayout` | Known | Navigation |
| 0x001ECE54 | `GotoVolumeLayout` | Known | Navigation |
| 0x001ECE68 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x001ECF08 | `GotoProgressLayout` | Known | Navigation |
| 0x001ECF1C | `GotoProgressVideoLayout` | Known | Navigation |
| 0x001ED3E4 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x001ED644 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x001ED7D8 | `GotoProgressLayout` | Known | Navigation |
| 0x001ED7EC | `GotoProgressVideoLayout` | Known | Navigation |
| 0x001ED8B0 | `GotoBrightnessVideoLayout` | Known | Navigation |
| 0x001ED8CC | `GotoRatingLayout` | Known | Navigation |
| 0x001EDB78 | `GotoChapterArtLayout` | Known | Navigation |
| 0x001EDB90 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x001EDBA4 | `GotoShuffleLayout` | Known | Navigation |
| 0x001EDF08 | `GotoVolumeLayout` | Known | Navigation |
| 0x001EDF20 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x001EDFAC | `GotoVolumeLayout` | Known | Navigation |
| 0x001EDFC0 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x001EE1D8 | `GotoScrubLayout` | Known | Navigation |
| 0x001EE1E8 | `GotoScrubVideoLayout` | Known | Navigation |
| 0x001EE278 | `GotoProgressLayout` | Known | Navigation |
| 0x001EE28C | `GotoProgressVideoLayout` | Known | Navigation |
| 0x001EE3F0 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x001EE40C | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x001EE424 | `GotoDefaultLayout` | Known | Navigation |
| 0x001EE508 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x001EE6A0 | `GotoChapterArtLayout` | Known | Navigation |
| 0x001EE6B8 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x001EE7A8 | `GotoProgressLayout` | Known | Navigation |
| 0x001EE834 | `GotoProgressLayout` | Known | Navigation |
| 0x001EE848 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x001EEAA4 | `GotoStatusBarLayout` | Known | Navigation |
| 0x001EEAB8 | `GotoDefaultLayout` | Known | Navigation |
| 0x001EEC90 | `GotoDefault` | Known | Navigation |
| 0x001EEDC4 | `GotoProgressLayout` | Known | Navigation |
| 0x001EEF84 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x001EF0D4 | `GotoBrightnessLayout` | Known | Navigation |
| 0x001EF158 | `GotoBrightnessLayout` | Known | Navigation |
| 0x001EF1D8 | `GotoVolumeLayout` | Known | Navigation |
| 0x001EF224 | `GotoScrubLayout` | Known | Navigation |
| 0x001EF2F8 | `GotoStatusBarLayout` | Known | Navigation |
| 0x001EF30C | `GotoDefaultLayout` | Known | Navigation |
| 0x001EF3E4 | `GotoScrubLayout` | Known | Navigation |
| 0x001EF434 | `GotoScrubLayout` | Known | Navigation |
| 0x001F4C6C | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x001F4E74 | `GotoScreen_BacklightTimer` | Known | Navigation |
| 0x001F4F04 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x001F4F1C | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation |
| 0x001F8E6C | `GotoScreen_LockDialog` | Known | Navigation |
| 0x001F8E84 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x001FA89C | `GotoRadio` | Known | Navigation |
| 0x001FB234 | `GotoNowPlaying` | Known | Navigation |
| 0x001FB944 | `GotoNowPlaying` | Known | Navigation |
| 0x001FBFC4 | `GotoFirstBoot` | Known | Navigation |
| 0x001FBFD4 | `GotoNotesApp` | Known | Navigation |
| 0x001FBFE8 | `GotoLockApp` | Known | Navigation |
| 0x002011A4 | `GotoNowPlaying` | Known | Navigation |
| 0x003751D8 | `GotoProgressLayout` | Known | Navigation |
| 0x0069328D | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x006FEC01 | `GotoDefault` | Known | Navigation |
| 0x006FF4E9 | `GotoDefault` | Known | Navigation |
| 0x007D6CA4 | `Gotowe?` | Known | Navigation |

---

## 7. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0014D8F8 | `CoverFlow_Screen` | Known | Screen layout |
| 0x0017656C | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x0017658C | `VoiceMemos_Menu_Screen_Recording` | Known | Screen layout |
| 0x001765B0 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x006873B2 | `Clock_Screen` | Known | Screen layout |
| 0x006873C2 | `Clock_Screen_Default"` | Known | Screen layout |
| 0x00687427 | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x00687485 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0068749D | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x0068750A | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x006875A8 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x00687607 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0068761D | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x00687688 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x006876E2 | `Games_Menu_Screen` | Known | Screen layout |
| 0x006876F7 | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x00687761 | `Extras_Screen_Games` | Known | Screen layout |
| 0x00687820 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x006878E4 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x006879AD | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x00687AE8 | `Clock_Africa_City_Screen ` | Known | Screen layout |
| 0x00687B04 | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x00687B88 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x00687BA2 | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x00687C24 | `Clock_Atlantic_City_Screen"` | Known | Screen layout |
| 0x00687C42 | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x00687CC8 | `Clock_Australia_City_Screen#` | Known | Screen layout |
| 0x00687CE7 | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x00687D6E | `Clock_Europe_City_Screen ` | Known | Screen layout |
| 0x00687D8A | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x00687E0E | `Clock_NorthAmerica_City_Screen&` | Known | Screen layout |
| 0x00687E30 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x00687EBA | `Clock_Pacific_City_Screen!` | Known | Screen layout |
| 0x00687ED7 | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x00687F5C | `Clock_SouthAmerica_City_Screen&` | Known | Screen layout |
| 0x00687F7E | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0068800B | `Clock_Screen"` | Known | Screen layout |
| 0x006880B0 | `Clock_Screen"` | Known | Screen layout |
| 0x00688155 | `Clock_Screen"` | Known | Screen layout |
| 0x006881FA | `Clock_Screen"` | Known | Screen layout |
| 0x0068829F | `Clock_Screen"` | Known | Screen layout |
| 0x00688344 | `Clock_Screen"` | Known | Screen layout |
| 0x006883E9 | `Clock_Screen"` | Known | Screen layout |
| 0x0068848E | `Clock_Screen"` | Known | Screen layout |
| 0x00688533 | `Clock_Screen"` | Known | Screen layout |
| 0x006885D8 | `Clock_Screen"` | Known | Screen layout |
| 0x0068867D | `Clock_Screen"` | Known | Screen layout |
| 0x00688722 | `Clock_Screen"` | Known | Screen layout |
| 0x006887C7 | `Clock_Screen"` | Known | Screen layout |
| 0x0068886C | `Clock_Screen"` | Known | Screen layout |
| 0x00688911 | `Clock_Screen"` | Known | Screen layout |
| 0x006889B6 | `Clock_Screen"` | Known | Screen layout |
| 0x00688A5B | `Clock_Screen"` | Known | Screen layout |
| 0x00688B00 | `Clock_Screen"` | Known | Screen layout |
| 0x00688BA5 | `Clock_Screen"` | Known | Screen layout |
| 0x00688C4A | `Clock_Screen"` | Known | Screen layout |
| 0x00688CEF | `Clock_Screen"` | Known | Screen layout |
| 0x00688D94 | `Clock_Screen"` | Known | Screen layout |
| 0x00688E39 | `Clock_Screen"` | Known | Screen layout |
| 0x00688EDE | `Clock_Screen"` | Known | Screen layout |
| 0x00688F83 | `Clock_Screen"` | Known | Screen layout |
| 0x00689028 | `Clock_Screen"` | Known | Screen layout |
| 0x006890CD | `Clock_Screen"` | Known | Screen layout |
| 0x00689172 | `Clock_Screen"` | Known | Screen layout |
| 0x00689217 | `Clock_Screen"` | Known | Screen layout |
| 0x006892BC | `Clock_Screen"` | Known | Screen layout |
| 0x00689361 | `Clock_Screen"` | Known | Screen layout |
| 0x0068940B | `Clock_Screen"` | Known | Screen layout |
| 0x006894B0 | `Clock_Screen"` | Known | Screen layout |
| 0x00689555 | `Clock_Screen"` | Known | Screen layout |
| 0x006895FA | `Clock_Screen"` | Known | Screen layout |
| 0x0068969F | `Clock_Screen"` | Known | Screen layout |
| 0x00689744 | `Clock_Screen"` | Known | Screen layout |
| 0x006897E9 | `Clock_Screen"` | Known | Screen layout |
| 0x0068988E | `Clock_Screen"` | Known | Screen layout |
| 0x00689933 | `Clock_Screen"` | Known | Screen layout |
| 0x006899D8 | `Clock_Screen"` | Known | Screen layout |
| 0x00689A7D | `Clock_Screen"` | Known | Screen layout |
| 0x00689B22 | `Clock_Screen"` | Known | Screen layout |
| 0x00689BC7 | `Clock_Screen"` | Known | Screen layout |
| 0x00689C6C | `Clock_Screen"` | Known | Screen layout |
| 0x00689D11 | `Clock_Screen"` | Known | Screen layout |
| 0x00689DB6 | `Clock_Screen"` | Known | Screen layout |
| 0x00689E5B | `Clock_Screen"` | Known | Screen layout |
| 0x00689F00 | `Clock_Screen"` | Known | Screen layout |
| 0x00689FA5 | `Clock_Screen"` | Known | Screen layout |
| 0x0068A04A | `Clock_Screen"` | Known | Screen layout |
| 0x0068A0EF | `Clock_Screen"` | Known | Screen layout |
| 0x0068A194 | `Clock_Screen"` | Known | Screen layout |
| 0x0068A239 | `Clock_Screen"` | Known | Screen layout |
| 0x0068A2DE | `Clock_Screen"` | Known | Screen layout |
| 0x0068A383 | `Clock_Screen"` | Known | Screen layout |
| 0x0068A428 | `Clock_Screen"` | Known | Screen layout |
| 0x0068A4CD | `Clock_Screen"` | Known | Screen layout |
| 0x0068A572 | `Clock_Screen"` | Known | Screen layout |
| 0x0068A617 | `Clock_Screen"` | Known | Screen layout |
| 0x0068A6BC | `Clock_Screen"` | Known | Screen layout |
| 0x0068A761 | `Clock_Screen"` | Known | Screen layout |
| 0x0068A806 | `Clock_Screen"` | Known | Screen layout |
| 0x0068A8AB | `Clock_Screen"` | Known | Screen layout |
| 0x0068A950 | `Clock_Screen"` | Known | Screen layout |
| 0x0068A9F5 | `Clock_Screen"` | Known | Screen layout |
| 0x0068AA9A | `Clock_Screen"` | Known | Screen layout |
| 0x0068AB3F | `Clock_Screen"` | Known | Screen layout |
| 0x0068ABE4 | `Clock_Screen"` | Known | Screen layout |
| 0x0068AC89 | `Clock_Screen"` | Known | Screen layout |
| 0x0068AD2E | `Clock_Screen"` | Known | Screen layout |
| 0x0068ADD3 | `Clock_Screen"` | Known | Screen layout |
| 0x0068AE78 | `Clock_Screen"` | Known | Screen layout |
| 0x0068AF1D | `Clock_Screen"` | Known | Screen layout |
| 0x0068AFC2 | `Clock_Screen"` | Known | Screen layout |
| 0x0068B067 | `Clock_Screen"` | Known | Screen layout |
| 0x0068B10C | `Clock_Screen"` | Known | Screen layout |
| 0x0068B1B1 | `Clock_Screen"` | Known | Screen layout |
| 0x0068B256 | `Clock_Screen"` | Known | Screen layout |
| 0x0068B2FB | `Clock_Screen"` | Known | Screen layout |
| 0x0068B3A0 | `Clock_Screen"` | Known | Screen layout |
| 0x0068B445 | `Clock_Screen"` | Known | Screen layout |
| 0x0068B4EA | `Clock_Screen"` | Known | Screen layout |
| 0x0068B58F | `Clock_Screen"` | Known | Screen layout |
| 0x0068B634 | `Clock_Screen"` | Known | Screen layout |
| 0x0068B6D9 | `Clock_Screen"` | Known | Screen layout |
| 0x0068B77E | `Clock_Screen"` | Known | Screen layout |
| 0x0068B823 | `Clock_Screen"` | Known | Screen layout |
| 0x0068B8CF | `Clock_Screen"` | Known | Screen layout |
| 0x0068B974 | `Clock_Screen"` | Known | Screen layout |
| 0x0068BA19 | `Clock_Screen"` | Known | Screen layout |
| 0x0068BAC3 | `Clock_Screen"` | Known | Screen layout |
| 0x0068BB68 | `Clock_Screen"` | Known | Screen layout |
| 0x0068BC0D | `Clock_Screen"` | Known | Screen layout |
| 0x0068BCB2 | `Clock_Screen"` | Known | Screen layout |
| 0x0068BD57 | `Clock_Screen"` | Known | Screen layout |
| 0x0068BDFC | `Clock_Screen"` | Known | Screen layout |
| 0x0068BEA1 | `Clock_Screen"` | Known | Screen layout |
| 0x0068BF46 | `Clock_Screen"` | Known | Screen layout |
| 0x0068BFEF | `Clock_Screen"` | Known | Screen layout |
| 0x0068C094 | `Clock_Screen"` | Known | Screen layout |
| 0x0068C139 | `Clock_Screen"` | Known | Screen layout |
| 0x0068C1DE | `Clock_Screen"` | Known | Screen layout |
| 0x0068C283 | `Clock_Screen"` | Known | Screen layout |
| 0x0068C328 | `Clock_Screen"` | Known | Screen layout |
| 0x0068C3CD | `Clock_Screen"` | Known | Screen layout |
| 0x0068C472 | `Clock_Screen"` | Known | Screen layout |
| 0x0068C517 | `Clock_Screen"` | Known | Screen layout |
| 0x0068C5BC | `Clock_Screen"` | Known | Screen layout |
| 0x0068C661 | `Clock_Screen"` | Known | Screen layout |
| 0x0068C706 | `Clock_Screen"` | Known | Screen layout |
| 0x0068C7AB | `Clock_Screen"` | Known | Screen layout |
| 0x0068C850 | `Clock_Screen"` | Known | Screen layout |
| 0x0068C8F5 | `Clock_Screen"` | Known | Screen layout |
| 0x0068C99A | `Clock_Screen"` | Known | Screen layout |
| 0x0068CA3F | `Clock_Screen"` | Known | Screen layout |
| 0x0068CAE4 | `Clock_Screen"` | Known | Screen layout |
| 0x0068CB89 | `Clock_Screen"` | Known | Screen layout |
| 0x0068CC2E | `Clock_Screen"` | Known | Screen layout |
| 0x0068CCD3 | `Clock_Screen"` | Known | Screen layout |
| 0x0068CD78 | `Clock_Screen"` | Known | Screen layout |
| 0x0068CE1D | `Clock_Screen"` | Known | Screen layout |
| 0x0068CEC2 | `Clock_Screen"` | Known | Screen layout |
| 0x0068CF67 | `Clock_Screen"` | Known | Screen layout |
| 0x0068D00C | `Clock_Screen"` | Known | Screen layout |
| 0x0068D0B1 | `Clock_Screen"` | Known | Screen layout |
| 0x0068D156 | `Clock_Screen"` | Known | Screen layout |
| 0x0068D1FB | `Clock_Screen"` | Known | Screen layout |
| 0x0068D2A0 | `Clock_Screen"` | Known | Screen layout |
| 0x0068D345 | `Clock_Screen"` | Known | Screen layout |
| 0x0068D3EA | `Clock_Screen"` | Known | Screen layout |
| 0x0068D48F | `Clock_Screen"` | Known | Screen layout |
| 0x0068D534 | `Clock_Screen"` | Known | Screen layout |
| 0x0068D5DF | `Clock_Screen"` | Known | Screen layout |
| 0x0068D684 | `Clock_Screen"` | Known | Screen layout |
| 0x0068D729 | `Clock_Screen"` | Known | Screen layout |
| 0x0068D7CE | `Clock_Screen"` | Known | Screen layout |
| 0x0068D873 | `Clock_Screen"` | Known | Screen layout |
| 0x0068D918 | `Clock_Screen"` | Known | Screen layout |
| 0x0068D9BD | `Clock_Screen"` | Known | Screen layout |
| 0x0068DA62 | `Clock_Screen"` | Known | Screen layout |
| 0x0068DB07 | `Clock_Screen"` | Known | Screen layout |
| 0x0068DBAC | `Clock_Screen"` | Known | Screen layout |
| 0x0068DC51 | `Clock_Screen"` | Known | Screen layout |
| 0x0068DCF6 | `Clock_Screen"` | Known | Screen layout |
| 0x0068DD9B | `Clock_Screen"` | Known | Screen layout |
| 0x0068DE40 | `Clock_Screen"` | Known | Screen layout |
| 0x0068DEE5 | `Clock_Screen"` | Known | Screen layout |
| 0x0068DF8A | `Clock_Screen"` | Known | Screen layout |
| 0x0068E02F | `Clock_Screen"` | Known | Screen layout |
| 0x0068E0D4 | `Clock_Screen"` | Known | Screen layout |
| 0x0068E179 | `Clock_Screen"` | Known | Screen layout |
| 0x0068E21E | `Clock_Screen"` | Known | Screen layout |
| 0x0068E2C3 | `Clock_Screen"` | Known | Screen layout |
| 0x0068E368 | `Clock_Screen"` | Known | Screen layout |
| 0x0068E40D | `Clock_Screen"` | Known | Screen layout |
| 0x0068E4B2 | `Clock_Screen"` | Known | Screen layout |
| 0x0068E557 | `Clock_Screen"` | Known | Screen layout |
| 0x0068E5FC | `Clock_Screen"` | Known | Screen layout |
| 0x0068E6A1 | `Clock_Screen"` | Known | Screen layout |
| 0x0068E746 | `Clock_Screen"` | Known | Screen layout |
| 0x0068E7EB | `Clock_Screen"` | Known | Screen layout |
| 0x0068E890 | `Clock_Screen"` | Known | Screen layout |
| 0x0068E935 | `Clock_Screen"` | Known | Screen layout |
| 0x0068E9DA | `Clock_Screen"` | Known | Screen layout |
| 0x0068EA7F | `Clock_Screen"` | Known | Screen layout |
| 0x0068EB24 | `Clock_Screen"` | Known | Screen layout |
| 0x0068EBC9 | `Clock_Screen"` | Known | Screen layout |
| 0x0068EC6E | `Clock_Screen"` | Known | Screen layout |
| 0x0068ED13 | `Clock_Screen"` | Known | Screen layout |
| 0x0068EDB8 | `Clock_Screen"` | Known | Screen layout |
| 0x0068EE5D | `Clock_Screen"` | Known | Screen layout |
| 0x0068EF02 | `Clock_Screen"` | Known | Screen layout |
| 0x0068EFA7 | `Clock_Screen"` | Known | Screen layout |
| 0x0068F04C | `Clock_Screen"` | Known | Screen layout |
| 0x0068F0F1 | `Clock_Screen"` | Known | Screen layout |
| 0x0068F196 | `Clock_Screen"` | Known | Screen layout |
| 0x0068F23B | `Clock_Screen"` | Known | Screen layout |
| 0x0068F2E0 | `Clock_Screen"` | Known | Screen layout |
| 0x0068F385 | `Clock_Screen"` | Known | Screen layout |
| 0x0068F42A | `Clock_Screen"` | Known | Screen layout |
| 0x0068F4CF | `Clock_Screen"` | Known | Screen layout |
| 0x0068F574 | `Clock_Screen"` | Known | Screen layout |
| 0x0068F61F | `Clock_Screen"` | Known | Screen layout |
| 0x0068F6C4 | `Clock_Screen"` | Known | Screen layout |
| 0x0068F769 | `Clock_Screen"` | Known | Screen layout |
| 0x0068F80E | `Clock_Screen"` | Known | Screen layout |
| 0x0068F8B3 | `Clock_Screen"` | Known | Screen layout |
| 0x0068F95F | `Clock_Screen"` | Known | Screen layout |
| 0x0068FA04 | `Clock_Screen"` | Known | Screen layout |
| 0x0068FAA9 | `Clock_Screen"` | Known | Screen layout |
| 0x0068FB4E | `Clock_Screen"` | Known | Screen layout |
| 0x0068FBF3 | `Clock_Screen"` | Known | Screen layout |
| 0x0068FC98 | `Clock_Screen"` | Known | Screen layout |
| 0x0068FD3D | `Clock_Screen"` | Known | Screen layout |
| 0x0068FDE2 | `Clock_Screen"` | Known | Screen layout |
| 0x0068FE87 | `Clock_Screen"` | Known | Screen layout |
| 0x0068FF2C | `Clock_Screen"` | Known | Screen layout |
| 0x0068FFD1 | `Clock_Screen"` | Known | Screen layout |
| 0x00690076 | `Clock_Screen"` | Known | Screen layout |
| 0x0069011B | `Clock_Screen"` | Known | Screen layout |
| 0x006901C0 | `Clock_Screen"` | Known | Screen layout |
| 0x00690265 | `Clock_Screen"` | Known | Screen layout |
| 0x0069030A | `Clock_Screen"` | Known | Screen layout |
| 0x006903AF | `Clock_Screen"` | Known | Screen layout |
| 0x00690452 | `Settings_DateTime_SetDate_Screen(` | Known | Screen layout |
| 0x00690476 | `Settings_DateTime_SetDate_Screen_Default"` | Known | Screen layout |
| 0x006904EF | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x00690555 | `Settings_DateTime_SetTime_Screen(` | Known | Screen layout |
| 0x00690579 | `Settings_DateTime_SetTime_Screen_Default"` | Known | Screen layout |
| 0x006905F2 | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x0069065D | `Settings_DateTime_SetTimeZone_Screen,` | Known | Screen layout |
| 0x00690685 | `Settings_DateTime_SetTimeZone_Screen_Default"` | Known | Screen layout |
| 0x00690702 | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x006907BB | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0069086B | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x00690DFA | `Search_Main_Screen` | Known | Screen layout |
| 0x00690E10 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x006912B4 | `Extras_Screen` | Known | Screen layout |
| 0x006912C5 | `Extras_Screen_WorldClock"` | Known | Screen layout |
| 0x00691342 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x006913A4 | `Clock_Screen` | Known | Screen layout |
| 0x006913B4 | `Clock_Screen_Default` | Known | Screen layout |
| 0x0069143B | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x006914A1 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x006914B7 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x00691522 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x00691584 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0069159C | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x00691609 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x0069166D | `AddressViewer_Main_Screen!` | Known | Screen layout |
| 0x0069168A | `AddressViewer_Main_Screen_Default"` | Known | Screen layout |
| 0x006916FC | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x00691763 | `Games_Menu_Screen` | Known | Screen layout |
| 0x00691778 | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x006917E2 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x006918A9 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x00691945 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x00691A16 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x00691AD6 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x00691B3A | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x00691B59 | `VoiceMemos_Menu_Screen_Default"` | Known | Screen layout |
| 0x00691BDC | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x00691C42 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x00691C5A | `Speakers_Main_Screen_Default"` | Known | Screen layout |
| 0x00691CDB | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x00691D3F | `Radio_Screen` | Known | Screen layout |
| 0x00691D4F | `Radio_Screen_Default"` | Known | Screen layout |
| 0x00691DC8 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x00691E66 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00691F02 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x00691FC5 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x00692084 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x00692141 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x0069255C | `Radio_Screen` | Known | Screen layout |
| 0x0069256C | `Radio_Screen_Default"` | Known | Screen layout |
| 0x006925E5 | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x00692806 | `Search_Main_Screen` | Known | Screen layout |
| 0x0069281C | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x00692948 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x006929AB | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x00692C6A | `Video_Settings_Screen` | Known | Screen layout |
| 0x00692C83 | `Video_Settings_Screen_Default"` | Known | Screen layout |
| 0x00692D6A | `PhotosSettings_Screen` | Known | Screen layout |
| 0x00692E27 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x00692E44 | `SettingsMenu_About_Screen_Capacity_Layout"` | Known | Screen layout |
| 0x00693091 | `SettingsMenus_MainMenu_Screen%` | Known | Screen layout |
| 0x0069319F | `SettingsMenus_MusicMenu_Screen&` | Known | Screen layout |
| 0x00693448 | `SettingsMenus_Brightness_Screen&` | Known | Screen layout |
| 0x0069355D | `SettingsMenus_AudiobookSettings_Screen.` | Known | Screen layout |
| 0x00693693 | `SettingsMenus_RadioRegions_Screen ` | Known | Screen layout |
| 0x006937A8 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x00693A14 | `Settings_DateTime_Screen ` | Known | Screen layout |
| 0x00693A30 | `Settings_DateTime_Screen_Default"` | Known | Screen layout |
| 0x00693BBC | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x00693CC1 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x00693CDA | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x00693DCB | `SettingsMenus_ResetAllSettings_Screen4` | Known | Screen layout |
| 0x0069459E | `Stopwatch_Screen` | Known | Screen layout |
| 0x006945B2 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x00694619 | `Stopwatch_Screen` | Known | Screen layout |
| 0x0069462D | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x006946D6 | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x006946F9 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x00694792 | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x006947B5 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x00694958 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x006949C6 | `Speakers_ToneControl_Screen#` | Known | Screen layout |
| 0x006949E5 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x006A6A31 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006A6AB4 | `LockediPod_Screen` | Known | Screen layout |
| 0x006A6B3C | `Lock_Screen` | Known | Screen layout |
| 0x006A6B4B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006A6BC2 | `AddressViewer_ContactDetails_Screen+` | Known | Screen layout |
| 0x006A6BE9 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x006A6C64 | `Extras_Screen` | Known | Screen layout |
| 0x006A6CAF | `Extras_Screen` | Known | Screen layout |
| 0x006A6D96 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x006A6DF4 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x006A6E11 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x006A6E7F | `Calendar_Event_Screen` | Known | Screen layout |
| 0x006A6E98 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x006A6F0F | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x006A6F2C | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x006A6F97 | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x006A6FB4 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x006A701B | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x006A7082 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x006A70E0 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x006A70FD | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x006A716B | `Calendar_Event_Screen` | Known | Screen layout |
| 0x006A7184 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x006A71FB | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x006A7218 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x006A7283 | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x006A72A0 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x006A7307 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x006A73A7 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x006A7430 | `Alarms_Set_Alarm_Frequency_Screen)` | Known | Screen layout |
| 0x006A7455 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x006A74C6 | `Alarms_Set_Alarm_Label_Screen%` | Known | Screen layout |
| 0x006A74E7 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x006A7554 | `Alarms_Set_Alarm_Sound_Screen%` | Known | Screen layout |
| 0x006A7575 | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x006A75E1 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x006A785C | `Alarms_Set_Alarm_Playlist_Screen(` | Known | Screen layout |
| 0x006A7880 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x006A78F0 | `Alarms_Set_Alarm_Tones_Screen%` | Known | Screen layout |
| 0x006A7911 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x006A7C24 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x006A7C3F | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x006A7D90 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x006A7DA7 | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x006A7E28 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x006A7E3F | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x006A7F15 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x006A7F2E | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x006A7FB3 | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x006A8024 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x006A8119 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x006A8132 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x006A81B7 | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x006A8228 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x006A82E8 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x006A82FC | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x006A842B | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x006A848E | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x006A84E5 | `Clock_Screen_Default` | Known | Screen layout |
| 0x006A8576 | `Clock_Region_Screen` | Known | Screen layout |
| 0x006A858D | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x006A8606 | `Clock_Screen_Default` | Known | Screen layout |
| 0x006A865D | `Clock_Screen_Default` | Known | Screen layout |
| 0x006A86EE | `Clock_Region_Screen` | Known | Screen layout |
| 0x006A8705 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x006A8890 | `CoverFlow_Screen_DefaultNoAlbumInfo&` | Known | Screen layout |
| 0x006A897E | `CoverFlow_Screen_DefaultNoAlbumInfo#` | Known | Screen layout |
| 0x006A89F3 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x006A8CE9 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x006A8E99 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x006A8FC7 | `CoverFlow_Screen_Default&` | Known | Screen layout |
| 0x006A909D | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x006A9232 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x006A9497 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x006A94F4 | `Game_Screen` | Known | Screen layout |
| 0x006A9503 | `Game_Screen_Default` | Known | Screen layout |
| 0x006A95A5 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x006A9607 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x006A966A | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x006A96CD | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x006A9729 | `Game_Running_Screen` | Known | Screen layout |
| 0x006A9789 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x006A97EB | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x006A984E | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x006A98B1 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x006A990D | `Game_Running_Screen` | Known | Screen layout |
| 0x006A996D | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x006A99CF | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x006A9A32 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x006A9A95 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x006A9AF1 | `Game_Running_Screen` | Known | Screen layout |
| 0x006A9B51 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x006A9BB3 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x006A9C16 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x006A9C79 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x006A9CD5 | `Game_Running_Screen` | Known | Screen layout |
| 0x006A9D35 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x006A9D97 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x006A9DFA | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x006A9E5D | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x006A9EB9 | `Game_Running_Screen` | Known | Screen layout |
| 0x006AA0FF | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x006AA161 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x006AA1C4 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x006AA227 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x006AA283 | `Game_Running_Screen` | Known | Screen layout |
| 0x006AA33A | `Extras_Screen` | Known | Screen layout |
| 0x006AA34B | `Extras_Screen_Lock` | Known | Screen layout |
| 0x006AA3A9 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x006AA546 | `Extras_Screen` | Known | Screen layout |
| 0x006AA557 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x006AA5B5 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x006AA752 | `Extras_Screen` | Known | Screen layout |
| 0x006AA763 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x006AA7C1 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x006AA95E | `Extras_Screen` | Known | Screen layout |
| 0x006AA96F | `Extras_Screen_Lock` | Known | Screen layout |
| 0x006AA9CD | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x006AAB6F | `Lock_Screen` | Known | Screen layout |
| 0x006AAB7E | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x006AABE0 | `Extras_Screen` | Known | Screen layout |
| 0x006AABF1 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x006AAC50 | `LockediPod_Screen` | Known | Screen layout |
| 0x006AACCA | `MainMenus_Main_Screen)` | Known | Screen layout |
| 0x006AAE9B | `Lock_Screen` | Known | Screen layout |
| 0x006AAEAA | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x006AAF0C | `Extras_Screen` | Known | Screen layout |
| 0x006AAF1D | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x006AAF7C | `LockediPod_Screen` | Known | Screen layout |
| 0x006AAFF6 | `MainMenus_Main_Screen%` | Known | Screen layout |
| 0x006AB05D | `LockediPod_Screen` | Known | Screen layout |
| 0x006AB072 | `LockediPod_Screen_Default,` | Known | Screen layout |
| 0x006AB1C1 | `Lock_Screen` | Known | Screen layout |
| 0x006AB1D0 | `Lock_Screen_Default_Layout"` | Known | Screen layout |
| 0x006AB239 | `Lock_Screen` | Known | Screen layout |
| 0x006AB248 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x006AB2AA | `Extras_Screen` | Known | Screen layout |
| 0x006AB2BB | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x006AB31A | `LockediPod_Screen` | Known | Screen layout |
| 0x006AB394 | `MainMenus_Main_Screen,` | Known | Screen layout |
| 0x006AB4F0 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x006AB556 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x006AB5BA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006AB649 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x006AB6B6 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x006AB723 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x006AB790 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006AB7F8 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x006AB85E | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x006AB8C2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006AB951 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x006AB9BE | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x006ABA2B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x006ABA98 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006ABB00 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x006ABB66 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x006ABBCA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006ABC59 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x006ABCC6 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x006ABD33 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x006ABDA0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006ABE08 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x006ABE6E | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x006ABED2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006ABF61 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x006ABFCE | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x006AC03B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x006AC0A8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006AC110 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x006AC176 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x006AC1DA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006AC269 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x006AC2D6 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x006AC343 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x006AC3B0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006AC409 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x006AC472 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x006AC4D9 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x006AC574 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x006AC5DD | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x006AC646 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x006AC6AD | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x006AC748 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x006AC7B1 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x006AC81A | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x006AC881 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x006AC91C | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x006ACA04 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006ACA20 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006ACA8E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006ACAAB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006ACB16 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006ACB36 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006ACBAD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006ACBC9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006ACC39 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006ACC58 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006ACCC4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006ACCD8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006ACD51 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006ACDC5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006ACE35 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006ACE9D | `NoContent_Screen` | Known | Screen layout |
| 0x006ACEB1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006ACF15 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006ACF7C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006ACF96 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006AD004 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006AD076 | `NoContent_Screen` | Known | Screen layout |
| 0x006AD08A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006AD0F4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006AD15D | `No_Photos_Screen` | Known | Screen layout |
| 0x006AD171 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006AD1D7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006AD245 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006AD2B2 | `NoContent_Screen` | Known | Screen layout |
| 0x006AD2C6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006AD32E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006AD398 | `NoContent_Screen` | Known | Screen layout |
| 0x006AD3AC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006AD419 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006AD48B | `NoContent_Screen` | Known | Screen layout |
| 0x006AD49F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006AD507 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006AD570 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006AD58B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006AD5F1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006AD60D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006AD6EC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006AD705 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006AD766 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006AD77A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006AD8E8 | `Radio_Screen` | Known | Screen layout |
| 0x006AD8F8 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006AD959 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006AD9DC | `LockediPod_Screen` | Known | Screen layout |
| 0x006ADA64 | `Lock_Screen` | Known | Screen layout |
| 0x006ADA73 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006ADAD6 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006ADB38 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006ADB54 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006ADBC6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006ADBE5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006ADC4D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006ADC67 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006ADCCF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006ADCEC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006ADD58 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006ADDC2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006ADDDC | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006ADE4C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006ADEBF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006ADF30 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006ADF9F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006AE00B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006AE026 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006AE09B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006AE102 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006AE164 | `Photos_Screen` | Known | Screen layout |
| 0x006AE1C8 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006AE1E6 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006AE256 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006AE271 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006AE2DA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006AE2F7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006AE36E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006AE392 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006AE400 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006AE41B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006AE4D8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006AE4F4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006AE562 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006AE57F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006AE5EA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006AE60A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006AE681 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006AE69D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006AE70D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006AE72C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006AE798 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006AE7AC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006AE825 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006AE899 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006AE909 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006AE971 | `NoContent_Screen` | Known | Screen layout |
| 0x006AE985 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006AE9E9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006AEA50 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006AEA6A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006AEAD8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006AEB4A | `NoContent_Screen` | Known | Screen layout |
| 0x006AEB5E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006AEBC8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006AEC31 | `No_Photos_Screen` | Known | Screen layout |
| 0x006AEC45 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006AECAB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006AED19 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006AED86 | `NoContent_Screen` | Known | Screen layout |
| 0x006AED9A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006AEE02 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006AEE6C | `NoContent_Screen` | Known | Screen layout |
| 0x006AEE80 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006AEEED | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006AEF5F | `NoContent_Screen` | Known | Screen layout |
| 0x006AEF73 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006AEFDB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006AF044 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006AF05F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006AF0C5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006AF0E1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006AF1C0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006AF1D9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006AF23A | `FirstBoot_Screen` | Known | Screen layout |
| 0x006AF24E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006AF3BC | `Radio_Screen` | Known | Screen layout |
| 0x006AF3CC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006AF42D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006AF4B0 | `LockediPod_Screen` | Known | Screen layout |
| 0x006AF538 | `Lock_Screen` | Known | Screen layout |
| 0x006AF547 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006AF5AA | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006AF60C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006AF628 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006AF69A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006AF6B9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006AF721 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006AF73B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006AF7A3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006AF7C0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006AF82C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006AF896 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006AF8B0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006AF920 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006AF993 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006AFA04 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006AFA73 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006AFADF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006AFAFA | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006AFB6F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006AFBD6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006AFC38 | `Photos_Screen` | Known | Screen layout |
| 0x006AFC9C | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006AFCBA | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006AFD2A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006AFD45 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006AFDAE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006AFDCB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006AFE42 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006AFE66 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006AFED4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006AFEEF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006AFFAC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006AFFC8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006B0036 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006B0053 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006B00BE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006B00DE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006B0155 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006B0171 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006B01E1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006B0200 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006B026C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006B0280 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006B02F9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006B036D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006B03DD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006B0445 | `NoContent_Screen` | Known | Screen layout |
| 0x006B0459 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006B04BD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006B0524 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006B053E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006B05AC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006B061E | `NoContent_Screen` | Known | Screen layout |
| 0x006B0632 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006B069C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006B0705 | `No_Photos_Screen` | Known | Screen layout |
| 0x006B0719 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006B077F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006B07ED | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006B085A | `NoContent_Screen` | Known | Screen layout |
| 0x006B086E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006B08D6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006B0940 | `NoContent_Screen` | Known | Screen layout |
| 0x006B0954 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006B09C1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006B0A33 | `NoContent_Screen` | Known | Screen layout |
| 0x006B0A47 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006B0AAF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006B0B18 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006B0B33 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006B0B99 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006B0BB5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006B0C94 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006B0CAD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006B0D0E | `FirstBoot_Screen` | Known | Screen layout |
| 0x006B0D22 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006B0E90 | `Radio_Screen` | Known | Screen layout |
| 0x006B0EA0 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006B0F01 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006B0F84 | `LockediPod_Screen` | Known | Screen layout |
| 0x006B100C | `Lock_Screen` | Known | Screen layout |
| 0x006B101B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006B107E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006B10E0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006B10FC | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006B116E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006B118D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006B11F5 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006B120F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006B1277 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006B1294 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006B1300 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006B136A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006B1384 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006B13F4 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006B1467 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006B14D8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006B1547 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006B15B3 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006B15CE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006B1643 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006B16AA | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006B170C | `Photos_Screen` | Known | Screen layout |
| 0x006B1770 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006B178E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006B17FE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006B1819 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006B1882 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006B189F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006B1916 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006B193A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006B19A8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006B19C3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006B1A80 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006B1A9C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006B1B0A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006B1B27 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006B1B92 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006B1BB2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006B1C29 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006B1C45 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006B1CB5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006B1CD4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006B1D40 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006B1D54 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006B1DCD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006B1E41 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006B1EB1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006B1F19 | `NoContent_Screen` | Known | Screen layout |
| 0x006B1F2D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006B1F91 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006B1FF8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006B2012 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006B2080 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006B20F2 | `NoContent_Screen` | Known | Screen layout |
| 0x006B2106 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006B2170 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006B21D9 | `No_Photos_Screen` | Known | Screen layout |
| 0x006B21ED | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006B2253 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006B22C1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006B232E | `NoContent_Screen` | Known | Screen layout |
| 0x006B2342 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006B23AA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006B2414 | `NoContent_Screen` | Known | Screen layout |
| 0x006B2428 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006B2495 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006B2507 | `NoContent_Screen` | Known | Screen layout |
| 0x006B251B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006B2583 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006B25EC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006B2607 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006B266D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006B2689 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006B2768 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006B2781 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006B27E2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006B27F6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006B2964 | `Radio_Screen` | Known | Screen layout |
| 0x006B2974 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006B29D5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006B2A58 | `LockediPod_Screen` | Known | Screen layout |
| 0x006B2AE0 | `Lock_Screen` | Known | Screen layout |
| 0x006B2AEF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006B2B52 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006B2BB4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006B2BD0 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006B2C42 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006B2C61 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006B2CC9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006B2CE3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006B2D4B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006B2D68 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006B2DD4 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006B2E3E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006B2E58 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006B2EC8 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006B2F3B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006B2FAC | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006B301B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006B3087 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006B30A2 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006B3117 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006B317E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006B31E0 | `Photos_Screen` | Known | Screen layout |
| 0x006B3244 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006B3262 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006B32D2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006B32ED | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006B3356 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006B3373 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006B33EA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006B340E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006B347C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006B3497 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006B3554 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006B3570 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006B35DE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006B35FB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006B3666 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006B3686 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006B36FD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006B3719 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006B3789 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006B37A8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006B3814 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006B3828 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006B38A1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006B3915 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006B3985 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006B39ED | `NoContent_Screen` | Known | Screen layout |
| 0x006B3A01 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006B3A65 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006B3ACC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006B3AE6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006B3B54 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006B3BC6 | `NoContent_Screen` | Known | Screen layout |
| 0x006B3BDA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006B3C44 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006B3CAD | `No_Photos_Screen` | Known | Screen layout |
| 0x006B3CC1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006B3D27 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006B3D95 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006B3E02 | `NoContent_Screen` | Known | Screen layout |
| 0x006B3E16 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006B3E7E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006B3EE8 | `NoContent_Screen` | Known | Screen layout |
| 0x006B3EFC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006B3F69 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006B3FDB | `NoContent_Screen` | Known | Screen layout |
| 0x006B3FEF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006B4057 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006B40C0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006B40DB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006B4141 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006B415D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006B423C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006B4255 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006B42B6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006B42CA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006B4438 | `Radio_Screen` | Known | Screen layout |
| 0x006B4448 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006B44A9 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006B452C | `LockediPod_Screen` | Known | Screen layout |
| 0x006B45B4 | `Lock_Screen` | Known | Screen layout |
| 0x006B45C3 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006B4626 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006B4688 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006B46A4 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006B4716 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006B4735 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006B479D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006B47B7 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006B481F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006B483C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006B48A8 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006B4912 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006B492C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006B499C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006B4A0F | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006B4A80 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006B4AEF | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006B4B5B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006B4B76 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006B4BEB | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006B4C52 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006B4CB4 | `Photos_Screen` | Known | Screen layout |
| 0x006B4D18 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006B4D36 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006B4DA6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006B4DC1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006B4E2A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006B4E47 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006B4EBE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006B4EE2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006B4F50 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006B4F6B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006B5028 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006B5044 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006B50B2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006B50CF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006B513A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006B515A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006B51D1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006B51ED | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006B525D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006B527C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006B52E8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006B52FC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006B5375 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006B53E9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006B5459 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006B54C1 | `NoContent_Screen` | Known | Screen layout |
| 0x006B54D5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006B5539 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006B55A0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006B55BA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006B5628 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006B569A | `NoContent_Screen` | Known | Screen layout |
| 0x006B56AE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006B5718 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006B5781 | `No_Photos_Screen` | Known | Screen layout |
| 0x006B5795 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006B57FB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006B5869 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006B58D6 | `NoContent_Screen` | Known | Screen layout |
| 0x006B58EA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006B5952 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006B59BC | `NoContent_Screen` | Known | Screen layout |
| 0x006B59D0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006B5A3D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006B5AAF | `NoContent_Screen` | Known | Screen layout |
| 0x006B5AC3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006B5B2B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006B5B94 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006B5BAF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006B5C15 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006B5C31 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006B5D10 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006B5D29 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006B5D8A | `FirstBoot_Screen` | Known | Screen layout |
| 0x006B5D9E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006B5F0C | `Radio_Screen` | Known | Screen layout |
| 0x006B5F1C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006B5F7D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006B6000 | `LockediPod_Screen` | Known | Screen layout |
| 0x006B6088 | `Lock_Screen` | Known | Screen layout |
| 0x006B6097 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006B60FA | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006B615C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006B6178 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006B61EA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006B6209 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006B6271 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006B628B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006B62F3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006B6310 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006B637C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006B63E6 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006B6400 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006B6470 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006B64E3 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006B6554 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006B65C3 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006B662F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006B664A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006B66BF | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006B6726 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006B6788 | `Photos_Screen` | Known | Screen layout |
| 0x006B67EC | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006B680A | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006B687A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006B6895 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006B68FE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006B691B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006B6992 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006B69B6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006B6A24 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006B6A3F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006B6AFC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006B6B18 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006B6B86 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006B6BA3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006B6C0E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006B6C2E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006B6CA5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006B6CC1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006B6D31 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006B6D50 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006B6DBC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006B6DD0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006B6E49 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006B6EBD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006B6F2D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006B6F95 | `NoContent_Screen` | Known | Screen layout |
| 0x006B6FA9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006B700D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006B7074 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006B708E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006B70FC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006B716E | `NoContent_Screen` | Known | Screen layout |
| 0x006B7182 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006B71EC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006B7255 | `No_Photos_Screen` | Known | Screen layout |
| 0x006B7269 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006B72CF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006B733D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006B73AA | `NoContent_Screen` | Known | Screen layout |
| 0x006B73BE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006B7426 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006B7490 | `NoContent_Screen` | Known | Screen layout |
| 0x006B74A4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006B7511 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006B7583 | `NoContent_Screen` | Known | Screen layout |
| 0x006B7597 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006B75FF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006B7668 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006B7683 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006B76E9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006B7705 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006B77E4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006B77FD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006B785E | `FirstBoot_Screen` | Known | Screen layout |
| 0x006B7872 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006B79E0 | `Radio_Screen` | Known | Screen layout |
| 0x006B79F0 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006B7A51 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006B7AD4 | `LockediPod_Screen` | Known | Screen layout |
| 0x006B7B5C | `Lock_Screen` | Known | Screen layout |
| 0x006B7B6B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006B7BCE | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006B7C30 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006B7C4C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006B7CBE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006B7CDD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006B7D45 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006B7D5F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006B7DC7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006B7DE4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006B7E50 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006B7EBA | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006B7ED4 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006B7F44 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006B7FB7 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006B8028 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006B8097 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006B8103 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006B811E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006B8193 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006B81FA | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006B825C | `Photos_Screen` | Known | Screen layout |
| 0x006B82C0 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006B82DE | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006B834E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006B8369 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006B83D2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006B83EF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006B8466 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006B848A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006B84F8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006B8513 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006B85D0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006B85EC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006B865A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006B8677 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006B86E2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006B8702 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006B8779 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006B8795 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006B8805 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006B8824 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006B8890 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006B88A4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006B891D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006B8991 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006B8A01 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006B8A69 | `NoContent_Screen` | Known | Screen layout |
| 0x006B8A7D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006B8AE1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006B8B48 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006B8B62 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006B8BD0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006B8C42 | `NoContent_Screen` | Known | Screen layout |
| 0x006B8C56 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006B8CC0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006B8D29 | `No_Photos_Screen` | Known | Screen layout |
| 0x006B8D3D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006B8DA3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006B8E11 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006B8E7E | `NoContent_Screen` | Known | Screen layout |
| 0x006B8E92 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006B8EFA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006B8F64 | `NoContent_Screen` | Known | Screen layout |
| 0x006B8F78 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006B8FE5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006B9057 | `NoContent_Screen` | Known | Screen layout |
| 0x006B906B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006B90D3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006B913C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006B9157 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006B91BD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006B91D9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006B92B8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006B92D1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006B9332 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006B9346 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006B94B4 | `Radio_Screen` | Known | Screen layout |
| 0x006B94C4 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006B9525 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006B95A8 | `LockediPod_Screen` | Known | Screen layout |
| 0x006B9630 | `Lock_Screen` | Known | Screen layout |
| 0x006B963F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006B96A2 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006B9704 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006B9720 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006B9792 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006B97B1 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006B9819 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006B9833 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006B989B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006B98B8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006B9924 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006B998E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006B99A8 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006B9A18 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006B9A8B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006B9AFC | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006B9B6B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006B9BD7 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006B9BF2 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006B9C67 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006B9CCE | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006B9D30 | `Photos_Screen` | Known | Screen layout |
| 0x006B9D94 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006B9DB2 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006B9E22 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006B9E3D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006B9EA6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006B9EC3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006B9F3A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006B9F5E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006B9FCC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006B9FE7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006BA0A4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006BA0C0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006BA12E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006BA14B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006BA1B6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006BA1D6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006BA24D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006BA269 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006BA2D9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006BA2F8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006BA364 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006BA378 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006BA3F1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006BA465 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006BA4D5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006BA53D | `NoContent_Screen` | Known | Screen layout |
| 0x006BA551 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006BA5B5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006BA61C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006BA636 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006BA6A4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006BA716 | `NoContent_Screen` | Known | Screen layout |
| 0x006BA72A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006BA794 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006BA7FD | `No_Photos_Screen` | Known | Screen layout |
| 0x006BA811 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006BA877 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006BA8E5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006BA952 | `NoContent_Screen` | Known | Screen layout |
| 0x006BA966 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006BA9CE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006BAA38 | `NoContent_Screen` | Known | Screen layout |
| 0x006BAA4C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006BAAB9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006BAB2B | `NoContent_Screen` | Known | Screen layout |
| 0x006BAB3F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006BABA7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006BAC10 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006BAC2B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006BAC91 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006BACAD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006BAD8C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006BADA5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006BAE06 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006BAE1A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006BAF88 | `Radio_Screen` | Known | Screen layout |
| 0x006BAF98 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006BAFF9 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006BB07C | `LockediPod_Screen` | Known | Screen layout |
| 0x006BB104 | `Lock_Screen` | Known | Screen layout |
| 0x006BB113 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006BB176 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006BB1D8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006BB1F4 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006BB266 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006BB285 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006BB2ED | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006BB307 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006BB36F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006BB38C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006BB3F8 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006BB462 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006BB47C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006BB4EC | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006BB55F | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006BB5D0 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006BB63F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006BB6AB | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006BB6C6 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006BB73B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006BB7A2 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006BB804 | `Photos_Screen` | Known | Screen layout |
| 0x006BB868 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006BB886 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006BB8F6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006BB911 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006BB97A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006BB997 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006BBA0E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006BBA32 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006BBAA0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006BBABB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006BBB78 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006BBB94 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006BBC02 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006BBC1F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006BBC8A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006BBCAA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006BBD21 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006BBD3D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006BBDAD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006BBDCC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006BBE38 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006BBE4C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006BBEC5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006BBF39 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006BBFA9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006BC011 | `NoContent_Screen` | Known | Screen layout |
| 0x006BC025 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006BC089 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006BC0F0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006BC10A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006BC178 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006BC1EA | `NoContent_Screen` | Known | Screen layout |
| 0x006BC1FE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006BC268 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006BC2D1 | `No_Photos_Screen` | Known | Screen layout |
| 0x006BC2E5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006BC34B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006BC3B9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006BC426 | `NoContent_Screen` | Known | Screen layout |
| 0x006BC43A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006BC4A2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006BC50C | `NoContent_Screen` | Known | Screen layout |
| 0x006BC520 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006BC58D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006BC5FF | `NoContent_Screen` | Known | Screen layout |
| 0x006BC613 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006BC67B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006BC6E4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006BC6FF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006BC765 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006BC781 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006BC860 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006BC879 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006BC8DA | `FirstBoot_Screen` | Known | Screen layout |
| 0x006BC8EE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006BCA5C | `Radio_Screen` | Known | Screen layout |
| 0x006BCA6C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006BCACD | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006BCB50 | `LockediPod_Screen` | Known | Screen layout |
| 0x006BCBD8 | `Lock_Screen` | Known | Screen layout |
| 0x006BCBE7 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006BCC4A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006BCCAC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006BCCC8 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006BCD3A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006BCD59 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006BCDC1 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006BCDDB | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006BCE43 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006BCE60 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006BCECC | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006BCF36 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006BCF50 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006BCFC0 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006BD033 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006BD0A4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006BD113 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006BD17F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006BD19A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006BD20F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006BD276 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006BD2D8 | `Photos_Screen` | Known | Screen layout |
| 0x006BD33C | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006BD35A | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006BD3CA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006BD3E5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006BD44E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006BD46B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006BD4E2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006BD506 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006BD574 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006BD58F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006BD64C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006BD668 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006BD6D6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006BD6F3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006BD75E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006BD77E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006BD7F5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006BD811 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006BD881 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006BD8A0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006BD90C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006BD920 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006BD999 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006BDA0D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006BDA7D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006BDAE5 | `NoContent_Screen` | Known | Screen layout |
| 0x006BDAF9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006BDB5D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006BDBC4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006BDBDE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006BDC4C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006BDCBE | `NoContent_Screen` | Known | Screen layout |
| 0x006BDCD2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006BDD3C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006BDDA5 | `No_Photos_Screen` | Known | Screen layout |
| 0x006BDDB9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006BDE1F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006BDE8D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006BDEFA | `NoContent_Screen` | Known | Screen layout |
| 0x006BDF0E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006BDF76 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006BDFE0 | `NoContent_Screen` | Known | Screen layout |
| 0x006BDFF4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006BE061 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006BE0D3 | `NoContent_Screen` | Known | Screen layout |
| 0x006BE0E7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006BE14F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006BE1B8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006BE1D3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006BE239 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006BE255 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006BE334 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006BE34D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006BE3AE | `FirstBoot_Screen` | Known | Screen layout |
| 0x006BE3C2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006BE530 | `Radio_Screen` | Known | Screen layout |
| 0x006BE540 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006BE5A1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006BE624 | `LockediPod_Screen` | Known | Screen layout |
| 0x006BE6AC | `Lock_Screen` | Known | Screen layout |
| 0x006BE6BB | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006BE71E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006BE780 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006BE79C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006BE80E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006BE82D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006BE895 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006BE8AF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006BE917 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006BE934 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006BE9A0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006BEA0A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006BEA24 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006BEA94 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006BEB07 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006BEB78 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006BEBE7 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006BEC53 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006BEC6E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006BECE3 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006BED4A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006BEDAC | `Photos_Screen` | Known | Screen layout |
| 0x006BEE10 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006BEE2E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006BEE9E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006BEEB9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006BEF22 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006BEF3F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006BEFB6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006BEFDA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006BF048 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006BF063 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006BF120 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006BF13C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006BF1AA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006BF1C7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006BF232 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006BF252 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006BF2C9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006BF2E5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006BF355 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006BF374 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006BF3E0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006BF3F4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006BF46D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006BF4E1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006BF551 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006BF5B9 | `NoContent_Screen` | Known | Screen layout |
| 0x006BF5CD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006BF631 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006BF698 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006BF6B2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006BF720 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006BF792 | `NoContent_Screen` | Known | Screen layout |
| 0x006BF7A6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006BF810 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006BF879 | `No_Photos_Screen` | Known | Screen layout |
| 0x006BF88D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006BF8F3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006BF961 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006BF9CE | `NoContent_Screen` | Known | Screen layout |
| 0x006BF9E2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006BFA4A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006BFAB4 | `NoContent_Screen` | Known | Screen layout |
| 0x006BFAC8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006BFB35 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006BFBA7 | `NoContent_Screen` | Known | Screen layout |
| 0x006BFBBB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006BFC23 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006BFC8C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006BFCA7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006BFD0D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006BFD29 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006BFE08 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006BFE21 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006BFE82 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006BFE96 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006C0004 | `Radio_Screen` | Known | Screen layout |
| 0x006C0014 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006C0075 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C00F8 | `LockediPod_Screen` | Known | Screen layout |
| 0x006C0180 | `Lock_Screen` | Known | Screen layout |
| 0x006C018F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C01F2 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006C0254 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006C0270 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006C02E2 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006C0301 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006C0369 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C0383 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006C03EB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C0408 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C0474 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006C04DE | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006C04F8 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006C0568 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006C05DB | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006C064C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006C06BB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006C0727 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006C0742 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006C07B7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006C081E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006C0880 | `Photos_Screen` | Known | Screen layout |
| 0x006C08E4 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006C0902 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006C0972 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006C098D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006C09F6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006C0A13 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006C0A8A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006C0AAE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006C0B1C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006C0B37 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006C0BF4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C0C10 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C0C7E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C0C9B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C0D06 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006C0D26 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006C0D9D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C0DB9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C0E29 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006C0E48 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006C0EB4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006C0EC8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006C0F41 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006C0FB5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006C1025 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006C108D | `NoContent_Screen` | Known | Screen layout |
| 0x006C10A1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006C1105 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006C116C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C1186 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006C11F4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006C1266 | `NoContent_Screen` | Known | Screen layout |
| 0x006C127A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006C12E4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006C134D | `No_Photos_Screen` | Known | Screen layout |
| 0x006C1361 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006C13C7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C1435 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006C14A2 | `NoContent_Screen` | Known | Screen layout |
| 0x006C14B6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006C151E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006C1588 | `NoContent_Screen` | Known | Screen layout |
| 0x006C159C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006C1609 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006C167B | `NoContent_Screen` | Known | Screen layout |
| 0x006C168F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C16F7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006C1760 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006C177B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006C17E1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006C17FD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006C18DC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006C18F5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006C1956 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006C196A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006C1AD8 | `Radio_Screen` | Known | Screen layout |
| 0x006C1AE8 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006C1B49 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C1BCC | `LockediPod_Screen` | Known | Screen layout |
| 0x006C1C54 | `Lock_Screen` | Known | Screen layout |
| 0x006C1C63 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C1CC6 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006C1D28 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006C1D44 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006C1DB6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006C1DD5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006C1E3D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C1E57 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006C1EBF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C1EDC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C1F48 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006C1FB2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006C1FCC | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006C203C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006C20AF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006C2120 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006C218F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006C21FB | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006C2216 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006C228B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006C22F2 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006C2354 | `Photos_Screen` | Known | Screen layout |
| 0x006C23B8 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006C23D6 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006C2446 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006C2461 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006C24CA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006C24E7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006C255E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006C2582 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006C25F0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006C260B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006C26C8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C26E4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C2752 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C276F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C27DA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006C27FA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006C2871 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C288D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C28FD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006C291C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006C2988 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006C299C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006C2A15 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006C2A89 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006C2AF9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006C2B61 | `NoContent_Screen` | Known | Screen layout |
| 0x006C2B75 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006C2BD9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006C2C40 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C2C5A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006C2CC8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006C2D3A | `NoContent_Screen` | Known | Screen layout |
| 0x006C2D4E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006C2DB8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006C2E21 | `No_Photos_Screen` | Known | Screen layout |
| 0x006C2E35 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006C2E9B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C2F09 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006C2F76 | `NoContent_Screen` | Known | Screen layout |
| 0x006C2F8A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006C2FF2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006C305C | `NoContent_Screen` | Known | Screen layout |
| 0x006C3070 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006C30DD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006C314F | `NoContent_Screen` | Known | Screen layout |
| 0x006C3163 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C31CB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006C3234 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006C324F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006C32B5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006C32D1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006C33B0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006C33C9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006C342A | `FirstBoot_Screen` | Known | Screen layout |
| 0x006C343E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006C35AC | `Radio_Screen` | Known | Screen layout |
| 0x006C35BC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006C361D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C36A0 | `LockediPod_Screen` | Known | Screen layout |
| 0x006C3728 | `Lock_Screen` | Known | Screen layout |
| 0x006C3737 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C379A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006C37FC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006C3818 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006C388A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006C38A9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006C3911 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C392B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006C3993 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C39B0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C3A1C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006C3A86 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006C3AA0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006C3B10 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006C3B83 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006C3BF4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006C3C63 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006C3CCF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006C3CEA | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006C3D5F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006C3DC6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006C3E28 | `Photos_Screen` | Known | Screen layout |
| 0x006C3E8C | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006C3EAA | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006C3F1A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006C3F35 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006C3F9E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006C3FBB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006C4032 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006C4056 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006C40C4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006C40DF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006C419C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C41B8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C4226 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C4243 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C42AE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006C42CE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006C4345 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C4361 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C43D1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006C43F0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006C445C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006C4470 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006C44E9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006C455D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006C45CD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006C4635 | `NoContent_Screen` | Known | Screen layout |
| 0x006C4649 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006C46AD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006C4714 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C472E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006C479C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006C480E | `NoContent_Screen` | Known | Screen layout |
| 0x006C4822 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006C488C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006C48F5 | `No_Photos_Screen` | Known | Screen layout |
| 0x006C4909 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006C496F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C49DD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006C4A4A | `NoContent_Screen` | Known | Screen layout |
| 0x006C4A5E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006C4AC6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006C4B30 | `NoContent_Screen` | Known | Screen layout |
| 0x006C4B44 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006C4BB1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006C4C23 | `NoContent_Screen` | Known | Screen layout |
| 0x006C4C37 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C4C9F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006C4D08 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006C4D23 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006C4D89 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006C4DA5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006C4E84 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006C4E9D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006C4EFE | `FirstBoot_Screen` | Known | Screen layout |
| 0x006C4F12 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006C5080 | `Radio_Screen` | Known | Screen layout |
| 0x006C5090 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006C50F1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C5174 | `LockediPod_Screen` | Known | Screen layout |
| 0x006C51FC | `Lock_Screen` | Known | Screen layout |
| 0x006C520B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C526E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006C52D0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006C52EC | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006C535E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006C537D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006C53E5 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C53FF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006C5467 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C5484 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C54F0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006C555A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006C5574 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006C55E4 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006C5657 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006C56C8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006C5737 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006C57A3 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006C57BE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006C5833 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006C589A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006C58FC | `Photos_Screen` | Known | Screen layout |
| 0x006C5960 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006C597E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006C59EE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006C5A09 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006C5A72 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006C5A8F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006C5B06 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006C5B2A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006C5B98 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006C5BB3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006C5C70 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C5C8C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C5CFA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C5D17 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C5D82 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006C5DA2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006C5E19 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C5E35 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C5EA5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006C5EC4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006C5F30 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006C5F44 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006C5FBD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006C6031 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006C60A1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006C6109 | `NoContent_Screen` | Known | Screen layout |
| 0x006C611D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006C6181 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006C61E8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C6202 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006C6270 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006C62E2 | `NoContent_Screen` | Known | Screen layout |
| 0x006C62F6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006C6360 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006C63C9 | `No_Photos_Screen` | Known | Screen layout |
| 0x006C63DD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006C6443 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C64B1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006C651E | `NoContent_Screen` | Known | Screen layout |
| 0x006C6532 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006C659A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006C6604 | `NoContent_Screen` | Known | Screen layout |
| 0x006C6618 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006C6685 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006C66F7 | `NoContent_Screen` | Known | Screen layout |
| 0x006C670B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C6773 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006C67DC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006C67F7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006C685D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006C6879 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006C6958 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006C6971 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006C69D2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006C69E6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006C6B54 | `Radio_Screen` | Known | Screen layout |
| 0x006C6B64 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006C6BC5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C6C48 | `LockediPod_Screen` | Known | Screen layout |
| 0x006C6CD0 | `Lock_Screen` | Known | Screen layout |
| 0x006C6CDF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C6D42 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006C6DA4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006C6DC0 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006C6E32 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006C6E51 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006C6EB9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C6ED3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006C6F3B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C6F58 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C6FC4 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006C702E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006C7048 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006C70B8 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006C712B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006C719C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006C720B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006C7277 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006C7292 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006C7307 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006C736E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006C73D0 | `Photos_Screen` | Known | Screen layout |
| 0x006C7434 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006C7452 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006C74C2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006C74DD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006C7546 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006C7563 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006C75DA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006C75FE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006C766C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006C7687 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006C7744 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C7760 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C77CE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C77EB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C7856 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006C7876 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006C78ED | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C7909 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C7979 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006C7998 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006C7A04 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006C7A18 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006C7A91 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006C7B05 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006C7B75 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006C7BDD | `NoContent_Screen` | Known | Screen layout |
| 0x006C7BF1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006C7C55 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006C7CBC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C7CD6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006C7D44 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006C7DB6 | `NoContent_Screen` | Known | Screen layout |
| 0x006C7DCA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006C7E34 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006C7E9D | `No_Photos_Screen` | Known | Screen layout |
| 0x006C7EB1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006C7F17 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C7F85 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006C7FF2 | `NoContent_Screen` | Known | Screen layout |
| 0x006C8006 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006C806E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006C80D8 | `NoContent_Screen` | Known | Screen layout |
| 0x006C80EC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006C8159 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006C81CB | `NoContent_Screen` | Known | Screen layout |
| 0x006C81DF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C8247 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006C82B0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006C82CB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006C8331 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006C834D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006C842C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006C8445 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006C84A6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006C84BA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006C8628 | `Radio_Screen` | Known | Screen layout |
| 0x006C8638 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006C8699 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006C871C | `LockediPod_Screen` | Known | Screen layout |
| 0x006C87A4 | `Lock_Screen` | Known | Screen layout |
| 0x006C87B3 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006C8816 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006C8878 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006C8894 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006C8906 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006C8925 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006C898D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C89A7 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006C8A0F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C8A2C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C8A98 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006C8B02 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006C8B1C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006C8B8C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006C8BFF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006C8C70 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006C8CDF | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006C8D4B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006C8D66 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006C8DDB | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006C8E42 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006C8EA4 | `Photos_Screen` | Known | Screen layout |
| 0x006C8F08 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006C8F26 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006C8F96 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006C8FB1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006C901A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006C9037 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006C90AE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006C90D2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006C9140 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006C915B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006C9218 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C9234 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C92A2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006C92BF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006C932A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006C934A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006C93C1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006C93DD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006C944D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006C946C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006C94D8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006C94EC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006C9565 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006C95D9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006C9649 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006C96B1 | `NoContent_Screen` | Known | Screen layout |
| 0x006C96C5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006C9729 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006C9790 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006C97AA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006C9818 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006C988A | `NoContent_Screen` | Known | Screen layout |
| 0x006C989E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006C9908 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006C9971 | `No_Photos_Screen` | Known | Screen layout |
| 0x006C9985 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006C99EB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C9A59 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006C9AC6 | `NoContent_Screen` | Known | Screen layout |
| 0x006C9ADA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006C9B42 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006C9BAC | `NoContent_Screen` | Known | Screen layout |
| 0x006C9BC0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006C9C2D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006C9C9F | `NoContent_Screen` | Known | Screen layout |
| 0x006C9CB3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006C9D1B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006C9D84 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006C9D9F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006C9E05 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006C9E21 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006C9F00 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006C9F19 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006C9F7A | `FirstBoot_Screen` | Known | Screen layout |
| 0x006C9F8E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006CA0FC | `Radio_Screen` | Known | Screen layout |
| 0x006CA10C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006CA16D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006CA1F0 | `LockediPod_Screen` | Known | Screen layout |
| 0x006CA278 | `Lock_Screen` | Known | Screen layout |
| 0x006CA287 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006CA2EA | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006CA34C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006CA368 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006CA3DA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006CA3F9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006CA461 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CA47B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006CA4E3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CA500 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CA56C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006CA5D6 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006CA5F0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006CA660 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006CA6D3 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006CA744 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006CA7B3 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006CA81F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006CA83A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006CA8AF | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006CA916 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006CA978 | `Photos_Screen` | Known | Screen layout |
| 0x006CA9DC | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006CA9FA | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006CAA6A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006CAA85 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006CAAEE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006CAB0B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006CAB82 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006CABA6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006CAC14 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006CAC2F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006CACEC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CAD08 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CAD76 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CAD93 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CADFE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006CAE1E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006CAE95 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CAEB1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CAF21 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006CAF40 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006CAFAC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006CAFC0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006CB039 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006CB0AD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006CB11D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006CB185 | `NoContent_Screen` | Known | Screen layout |
| 0x006CB199 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006CB1FD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006CB264 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CB27E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006CB2EC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006CB35E | `NoContent_Screen` | Known | Screen layout |
| 0x006CB372 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006CB3DC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006CB445 | `No_Photos_Screen` | Known | Screen layout |
| 0x006CB459 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006CB4BF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006CB52D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006CB59A | `NoContent_Screen` | Known | Screen layout |
| 0x006CB5AE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006CB616 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006CB680 | `NoContent_Screen` | Known | Screen layout |
| 0x006CB694 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006CB701 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006CB773 | `NoContent_Screen` | Known | Screen layout |
| 0x006CB787 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006CB7EF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006CB858 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006CB873 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006CB8D9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006CB8F5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006CB9D4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006CB9ED | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006CBA4E | `FirstBoot_Screen` | Known | Screen layout |
| 0x006CBA62 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006CBBD0 | `Radio_Screen` | Known | Screen layout |
| 0x006CBBE0 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006CBC41 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006CBCC4 | `LockediPod_Screen` | Known | Screen layout |
| 0x006CBD4C | `Lock_Screen` | Known | Screen layout |
| 0x006CBD5B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006CBDBE | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006CBE20 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006CBE3C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006CBEAE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006CBECD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006CBF35 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CBF4F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006CBFB7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CBFD4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CC040 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006CC0AA | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006CC0C4 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006CC134 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006CC1A7 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006CC218 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006CC287 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006CC2F3 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006CC30E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006CC383 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006CC3EA | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006CC44C | `Photos_Screen` | Known | Screen layout |
| 0x006CC4B0 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006CC4CE | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006CC53E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006CC559 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006CC5C2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006CC5DF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006CC656 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006CC67A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006CC6E8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006CC703 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006CC7C0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CC7DC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CC84A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CC867 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CC8D2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006CC8F2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006CC969 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CC985 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CC9F5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006CCA14 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006CCA80 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006CCA94 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006CCB0D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006CCB81 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006CCBF1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006CCC59 | `NoContent_Screen` | Known | Screen layout |
| 0x006CCC6D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006CCCD1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006CCD38 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CCD52 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006CCDC0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006CCE32 | `NoContent_Screen` | Known | Screen layout |
| 0x006CCE46 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006CCEB0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006CCF19 | `No_Photos_Screen` | Known | Screen layout |
| 0x006CCF2D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006CCF93 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006CD001 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006CD06E | `NoContent_Screen` | Known | Screen layout |
| 0x006CD082 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006CD0EA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006CD154 | `NoContent_Screen` | Known | Screen layout |
| 0x006CD168 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006CD1D5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006CD247 | `NoContent_Screen` | Known | Screen layout |
| 0x006CD25B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006CD2C3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006CD32C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006CD347 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006CD3AD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006CD3C9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006CD4A8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006CD4C1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006CD522 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006CD536 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006CD6A4 | `Radio_Screen` | Known | Screen layout |
| 0x006CD6B4 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006CD715 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006CD798 | `LockediPod_Screen` | Known | Screen layout |
| 0x006CD820 | `Lock_Screen` | Known | Screen layout |
| 0x006CD82F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006CD892 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006CD8F4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006CD910 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006CD982 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006CD9A1 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006CDA09 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CDA23 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006CDA8B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CDAA8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CDB14 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006CDB7E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006CDB98 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006CDC08 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006CDC7B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006CDCEC | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006CDD5B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006CDDC7 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006CDDE2 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006CDE57 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006CDEBE | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006CDF20 | `Photos_Screen` | Known | Screen layout |
| 0x006CDF84 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006CDFA2 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006CE012 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006CE02D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006CE096 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006CE0B3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006CE12A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006CE14E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006CE1BC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006CE1D7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006CE294 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CE2B0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CE31E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CE33B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CE3A6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006CE3C6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006CE43D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CE459 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CE4C9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006CE4E8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006CE554 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006CE568 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006CE5E1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006CE655 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006CE6C5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006CE72D | `NoContent_Screen` | Known | Screen layout |
| 0x006CE741 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006CE7A5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006CE80C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CE826 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006CE894 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006CE906 | `NoContent_Screen` | Known | Screen layout |
| 0x006CE91A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006CE984 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006CE9ED | `No_Photos_Screen` | Known | Screen layout |
| 0x006CEA01 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006CEA67 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006CEAD5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006CEB42 | `NoContent_Screen` | Known | Screen layout |
| 0x006CEB56 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006CEBBE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006CEC28 | `NoContent_Screen` | Known | Screen layout |
| 0x006CEC3C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006CECA9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006CED1B | `NoContent_Screen` | Known | Screen layout |
| 0x006CED2F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006CED97 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006CEE00 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006CEE1B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006CEE81 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006CEE9D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006CEF7C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006CEF95 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006CEFF6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006CF00A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006CF178 | `Radio_Screen` | Known | Screen layout |
| 0x006CF188 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006CF1E9 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006CF26C | `LockediPod_Screen` | Known | Screen layout |
| 0x006CF2F4 | `Lock_Screen` | Known | Screen layout |
| 0x006CF303 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006CF366 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006CF3C8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006CF3E4 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006CF456 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006CF475 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006CF4DD | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006CF4F7 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006CF55F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CF57C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CF5E8 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006CF652 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006CF66C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006CF6DC | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006CF74F | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006CF7C0 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006CF82F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006CF89B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006CF8B6 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006CF92B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006CF992 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006CF9F4 | `Photos_Screen` | Known | Screen layout |
| 0x006CFA58 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006CFA76 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006CFAE6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006CFB01 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006CFB6A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006CFB87 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006CFBFE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006CFC22 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006CFC90 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006CFCAB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006CFD68 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CFD84 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CFDF2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006CFE0F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006CFE7A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006CFE9A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006CFF11 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006CFF2D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006CFF9D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006CFFBC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006D0028 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006D003C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006D00B5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006D0129 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006D0199 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006D0201 | `NoContent_Screen` | Known | Screen layout |
| 0x006D0215 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006D0279 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006D02E0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D02FA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006D0368 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006D03DA | `NoContent_Screen` | Known | Screen layout |
| 0x006D03EE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006D0458 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006D04C1 | `No_Photos_Screen` | Known | Screen layout |
| 0x006D04D5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006D053B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D05A9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006D0616 | `NoContent_Screen` | Known | Screen layout |
| 0x006D062A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006D0692 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006D06FC | `NoContent_Screen` | Known | Screen layout |
| 0x006D0710 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006D077D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006D07EF | `NoContent_Screen` | Known | Screen layout |
| 0x006D0803 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D086B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006D08D4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006D08EF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006D0955 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006D0971 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006D0A50 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006D0A69 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006D0ACA | `FirstBoot_Screen` | Known | Screen layout |
| 0x006D0ADE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006D0C4C | `Radio_Screen` | Known | Screen layout |
| 0x006D0C5C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006D0CBD | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006D0D40 | `LockediPod_Screen` | Known | Screen layout |
| 0x006D0DC8 | `Lock_Screen` | Known | Screen layout |
| 0x006D0DD7 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006D0E3A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006D0E9C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006D0EB8 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006D0F2A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006D0F49 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006D0FB1 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D0FCB | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006D1033 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D1050 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D10BC | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006D1126 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006D1140 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006D11B0 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006D1223 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006D1294 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006D1303 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006D136F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006D138A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006D13FF | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006D1466 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006D14C8 | `Photos_Screen` | Known | Screen layout |
| 0x006D152C | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006D154A | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006D15BA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006D15D5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006D163E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006D165B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006D16D2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006D16F6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006D1764 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006D177F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006D183C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D1858 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D18C6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D18E3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D194E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006D196E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006D19E5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D1A01 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D1A71 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006D1A90 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006D1AFC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006D1B10 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006D1B89 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006D1BFD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006D1C6D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006D1CD5 | `NoContent_Screen` | Known | Screen layout |
| 0x006D1CE9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006D1D4D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006D1DB4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D1DCE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006D1E3C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006D1EAE | `NoContent_Screen` | Known | Screen layout |
| 0x006D1EC2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006D1F2C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006D1F95 | `No_Photos_Screen` | Known | Screen layout |
| 0x006D1FA9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006D200F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D207D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006D20EA | `NoContent_Screen` | Known | Screen layout |
| 0x006D20FE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006D2166 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006D21D0 | `NoContent_Screen` | Known | Screen layout |
| 0x006D21E4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006D2251 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006D22C3 | `NoContent_Screen` | Known | Screen layout |
| 0x006D22D7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D233F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006D23A8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006D23C3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006D2429 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006D2445 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006D2524 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006D253D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006D259E | `FirstBoot_Screen` | Known | Screen layout |
| 0x006D25B2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006D2720 | `Radio_Screen` | Known | Screen layout |
| 0x006D2730 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006D2791 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006D2814 | `LockediPod_Screen` | Known | Screen layout |
| 0x006D289C | `Lock_Screen` | Known | Screen layout |
| 0x006D28AB | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006D290E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006D2970 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006D298C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006D29FE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006D2A1D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006D2A85 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D2A9F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006D2B07 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D2B24 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D2B90 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006D2BFA | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006D2C14 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006D2C84 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006D2CF7 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006D2D68 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006D2DD7 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006D2E43 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006D2E5E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006D2ED3 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006D2F3A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006D2F9C | `Photos_Screen` | Known | Screen layout |
| 0x006D3000 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006D301E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006D308E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006D30A9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006D3112 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006D312F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006D31A6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006D31CA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006D3238 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006D3253 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006D3310 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D332C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D339A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D33B7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D3422 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006D3442 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006D34B9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D34D5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D3545 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006D3564 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006D35D0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006D35E4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006D365D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006D36D1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006D3741 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006D37A9 | `NoContent_Screen` | Known | Screen layout |
| 0x006D37BD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006D3821 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006D3888 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D38A2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006D3910 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006D3982 | `NoContent_Screen` | Known | Screen layout |
| 0x006D3996 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006D3A00 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006D3A69 | `No_Photos_Screen` | Known | Screen layout |
| 0x006D3A7D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006D3AE3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D3B51 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006D3BBE | `NoContent_Screen` | Known | Screen layout |
| 0x006D3BD2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006D3C3A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006D3CA4 | `NoContent_Screen` | Known | Screen layout |
| 0x006D3CB8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006D3D25 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006D3D97 | `NoContent_Screen` | Known | Screen layout |
| 0x006D3DAB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D3E13 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006D3E7C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006D3E97 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006D3EFD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006D3F19 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006D3FF8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006D4011 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006D4072 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006D4086 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006D41F4 | `Radio_Screen` | Known | Screen layout |
| 0x006D4204 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006D4265 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006D42E8 | `LockediPod_Screen` | Known | Screen layout |
| 0x006D4370 | `Lock_Screen` | Known | Screen layout |
| 0x006D437F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006D43E2 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006D4444 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006D4460 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006D44D2 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006D44F1 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006D4559 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D4573 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006D45DB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D45F8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D4664 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006D46CE | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006D46E8 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006D4758 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006D47CB | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006D483C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006D48AB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006D4917 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006D4932 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006D49A7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006D4A0E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006D4A70 | `Photos_Screen` | Known | Screen layout |
| 0x006D4AD4 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006D4AF2 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006D4B62 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006D4B7D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006D4BE6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006D4C03 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006D4C7A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006D4C9E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006D4D0C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006D4D27 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006D4DE4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D4E00 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D4E6E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D4E8B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D4EF6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006D4F16 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006D4F8D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D4FA9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D5019 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006D5038 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006D50A4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006D50B8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006D5131 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006D51A5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006D5215 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006D527D | `NoContent_Screen` | Known | Screen layout |
| 0x006D5291 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006D52F5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006D535C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D5376 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006D53E4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006D5456 | `NoContent_Screen` | Known | Screen layout |
| 0x006D546A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006D54D4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006D553D | `No_Photos_Screen` | Known | Screen layout |
| 0x006D5551 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006D55B7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D5625 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006D5692 | `NoContent_Screen` | Known | Screen layout |
| 0x006D56A6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006D570E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006D5778 | `NoContent_Screen` | Known | Screen layout |
| 0x006D578C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006D57F9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006D586B | `NoContent_Screen` | Known | Screen layout |
| 0x006D587F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D58E7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006D5950 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006D596B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006D59D1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006D59ED | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006D5ACC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006D5AE5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006D5B46 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006D5B5A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006D5CC8 | `Radio_Screen` | Known | Screen layout |
| 0x006D5CD8 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006D5D39 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006D5DBC | `LockediPod_Screen` | Known | Screen layout |
| 0x006D5E44 | `Lock_Screen` | Known | Screen layout |
| 0x006D5E53 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006D5EB6 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006D5F18 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006D5F34 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006D5FA6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006D5FC5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006D602D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D6047 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006D60AF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D60CC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D6138 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006D61A2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006D61BC | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006D622C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006D629F | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006D6310 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006D637F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006D63EB | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006D6406 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006D647B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006D64E2 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006D6544 | `Photos_Screen` | Known | Screen layout |
| 0x006D65A8 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006D65C6 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006D6636 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006D6651 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006D66BA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006D66D7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006D674E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006D6772 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006D67E0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006D67FB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006D68B8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D68D4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D6942 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D695F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D69CA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006D69EA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006D6A61 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D6A7D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D6AED | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006D6B0C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006D6B78 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006D6B8C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006D6C05 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006D6C79 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006D6CE9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006D6D51 | `NoContent_Screen` | Known | Screen layout |
| 0x006D6D65 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006D6DC9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006D6E30 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D6E4A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006D6EB8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006D6F2A | `NoContent_Screen` | Known | Screen layout |
| 0x006D6F3E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006D6FA8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006D7011 | `No_Photos_Screen` | Known | Screen layout |
| 0x006D7025 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006D708B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D70F9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006D7166 | `NoContent_Screen` | Known | Screen layout |
| 0x006D717A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006D71E2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006D724C | `NoContent_Screen` | Known | Screen layout |
| 0x006D7260 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006D72CD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006D733F | `NoContent_Screen` | Known | Screen layout |
| 0x006D7353 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D73BB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006D7424 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006D743F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006D74A5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006D74C1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006D75A0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006D75B9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006D761A | `FirstBoot_Screen` | Known | Screen layout |
| 0x006D762E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006D779C | `Radio_Screen` | Known | Screen layout |
| 0x006D77AC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006D780D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006D7890 | `LockediPod_Screen` | Known | Screen layout |
| 0x006D7918 | `Lock_Screen` | Known | Screen layout |
| 0x006D7927 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006D798A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006D79EC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006D7A08 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006D7A7A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006D7A99 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006D7B01 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D7B1B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006D7B83 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D7BA0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D7C0C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006D7C76 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006D7C90 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006D7D00 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006D7D73 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006D7DE4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006D7E53 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006D7EBF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006D7EDA | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006D7F4F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006D7FB6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006D8018 | `Photos_Screen` | Known | Screen layout |
| 0x006D807C | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006D809A | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006D810A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006D8125 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006D818E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006D81AB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006D8222 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006D8246 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006D82B4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006D82CF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006D838C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D83A8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D8416 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D8433 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D849E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006D84BE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006D8535 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D8551 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D85C1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006D85E0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006D864C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006D8660 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006D86D9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006D874D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006D87BD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006D8825 | `NoContent_Screen` | Known | Screen layout |
| 0x006D8839 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006D889D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006D8904 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D891E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006D898C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006D89FE | `NoContent_Screen` | Known | Screen layout |
| 0x006D8A12 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006D8A7C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006D8AE5 | `No_Photos_Screen` | Known | Screen layout |
| 0x006D8AF9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006D8B5F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D8BCD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006D8C3A | `NoContent_Screen` | Known | Screen layout |
| 0x006D8C4E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006D8CB6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006D8D20 | `NoContent_Screen` | Known | Screen layout |
| 0x006D8D34 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006D8DA1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006D8E13 | `NoContent_Screen` | Known | Screen layout |
| 0x006D8E27 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006D8E8F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006D8EF8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006D8F13 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006D8F79 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006D8F95 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006D9074 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006D908D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006D90EE | `FirstBoot_Screen` | Known | Screen layout |
| 0x006D9102 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006D9270 | `Radio_Screen` | Known | Screen layout |
| 0x006D9280 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006D92E1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006D9364 | `LockediPod_Screen` | Known | Screen layout |
| 0x006D93EC | `Lock_Screen` | Known | Screen layout |
| 0x006D93FB | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006D945E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006D94C0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006D94DC | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006D954E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006D956D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006D95D5 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006D95EF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006D9657 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D9674 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D96E0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006D974A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006D9764 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006D97D4 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006D9847 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006D98B8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006D9927 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006D9993 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006D99AE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006D9A23 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006D9A8A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006D9AEC | `Photos_Screen` | Known | Screen layout |
| 0x006D9B50 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006D9B6E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006D9BDE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006D9BF9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006D9C62 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006D9C7F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006D9CF6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006D9D1A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006D9D88 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006D9DA3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006D9E60 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006D9E7C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006D9EEA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006D9F07 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006D9F72 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006D9F92 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006DA009 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DA025 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DA095 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006DA0B4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006DA120 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006DA134 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006DA1AD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006DA221 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006DA291 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006DA2F9 | `NoContent_Screen` | Known | Screen layout |
| 0x006DA30D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006DA371 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006DA3D8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DA3F2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006DA460 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006DA4D2 | `NoContent_Screen` | Known | Screen layout |
| 0x006DA4E6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006DA550 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006DA5B9 | `No_Photos_Screen` | Known | Screen layout |
| 0x006DA5CD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006DA633 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DA6A1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006DA70E | `NoContent_Screen` | Known | Screen layout |
| 0x006DA722 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006DA78A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006DA7F4 | `NoContent_Screen` | Known | Screen layout |
| 0x006DA808 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006DA875 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006DA8E7 | `NoContent_Screen` | Known | Screen layout |
| 0x006DA8FB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DA963 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006DA9CC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006DA9E7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006DAA4D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006DAA69 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006DAB48 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006DAB61 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006DABC2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006DABD6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006DAD44 | `Radio_Screen` | Known | Screen layout |
| 0x006DAD54 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006DADB5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006DAE38 | `LockediPod_Screen` | Known | Screen layout |
| 0x006DAEC0 | `Lock_Screen` | Known | Screen layout |
| 0x006DAECF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006DAF32 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006DAF94 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006DAFB0 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006DB022 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006DB041 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006DB0A9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DB0C3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006DB12B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DB148 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DB1B4 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006DB21E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006DB238 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006DB2A8 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006DB31B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006DB38C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006DB3FB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006DB467 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006DB482 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006DB4F7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006DB55E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006DB5C0 | `Photos_Screen` | Known | Screen layout |
| 0x006DB624 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006DB642 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006DB6B2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006DB6CD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006DB736 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006DB753 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006DB7CA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006DB7EE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006DB85C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006DB877 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006DB934 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DB950 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DB9BE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DB9DB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DBA46 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006DBA66 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006DBADD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DBAF9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DBB69 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006DBB88 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006DBBF4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006DBC08 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006DBC81 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006DBCF5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006DBD65 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006DBDCD | `NoContent_Screen` | Known | Screen layout |
| 0x006DBDE1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006DBE45 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006DBEAC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DBEC6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006DBF34 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006DBFA6 | `NoContent_Screen` | Known | Screen layout |
| 0x006DBFBA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006DC024 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006DC08D | `No_Photos_Screen` | Known | Screen layout |
| 0x006DC0A1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006DC107 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DC175 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006DC1E2 | `NoContent_Screen` | Known | Screen layout |
| 0x006DC1F6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006DC25E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006DC2C8 | `NoContent_Screen` | Known | Screen layout |
| 0x006DC2DC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006DC349 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006DC3BB | `NoContent_Screen` | Known | Screen layout |
| 0x006DC3CF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DC437 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006DC4A0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006DC4BB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006DC521 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006DC53D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006DC61C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006DC635 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006DC696 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006DC6AA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006DC818 | `Radio_Screen` | Known | Screen layout |
| 0x006DC828 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006DC889 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006DC90C | `LockediPod_Screen` | Known | Screen layout |
| 0x006DC994 | `Lock_Screen` | Known | Screen layout |
| 0x006DC9A3 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006DCA06 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006DCA68 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006DCA84 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006DCAF6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006DCB15 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006DCB7D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DCB97 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006DCBFF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DCC1C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DCC88 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006DCCF2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006DCD0C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006DCD7C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006DCDEF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006DCE60 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006DCECF | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006DCF3B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006DCF56 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006DCFCB | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006DD032 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006DD094 | `Photos_Screen` | Known | Screen layout |
| 0x006DD0F8 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006DD116 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006DD186 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006DD1A1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006DD20A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006DD227 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006DD29E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006DD2C2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006DD330 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006DD34B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006DD408 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DD424 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DD492 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DD4AF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DD51A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006DD53A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006DD5B1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DD5CD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DD63D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006DD65C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006DD6C8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006DD6DC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006DD755 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006DD7C9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006DD839 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006DD8A1 | `NoContent_Screen` | Known | Screen layout |
| 0x006DD8B5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006DD919 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006DD980 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DD99A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006DDA08 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006DDA7A | `NoContent_Screen` | Known | Screen layout |
| 0x006DDA8E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006DDAF8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006DDB61 | `No_Photos_Screen` | Known | Screen layout |
| 0x006DDB75 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006DDBDB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DDC49 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006DDCB6 | `NoContent_Screen` | Known | Screen layout |
| 0x006DDCCA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006DDD32 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006DDD9C | `NoContent_Screen` | Known | Screen layout |
| 0x006DDDB0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006DDE1D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006DDE8F | `NoContent_Screen` | Known | Screen layout |
| 0x006DDEA3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DDF0B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006DDF74 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006DDF8F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006DDFF5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006DE011 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006DE0F0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006DE109 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006DE16A | `FirstBoot_Screen` | Known | Screen layout |
| 0x006DE17E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006DE2EC | `Radio_Screen` | Known | Screen layout |
| 0x006DE2FC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006DE35D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006DE3E0 | `LockediPod_Screen` | Known | Screen layout |
| 0x006DE468 | `Lock_Screen` | Known | Screen layout |
| 0x006DE477 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006DE4DA | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006DE53C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006DE558 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006DE5CA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006DE5E9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006DE651 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DE66B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006DE6D3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DE6F0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DE75C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006DE7C6 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006DE7E0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006DE850 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006DE8C3 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006DE934 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006DE9A3 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006DEA0F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006DEA2A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006DEA9F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006DEB06 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006DEB68 | `Photos_Screen` | Known | Screen layout |
| 0x006DEBCC | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006DEBEA | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006DEC5A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006DEC75 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006DECDE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006DECFB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006DED72 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006DED96 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006DEE04 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006DEE1F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006DEEDC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DEEF8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DEF66 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006DEF83 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006DEFEE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006DF00E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006DF085 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006DF0A1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006DF111 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006DF130 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006DF19C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006DF1B0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006DF229 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006DF29D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006DF30D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006DF375 | `NoContent_Screen` | Known | Screen layout |
| 0x006DF389 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006DF3ED | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006DF454 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006DF46E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006DF4DC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006DF54E | `NoContent_Screen` | Known | Screen layout |
| 0x006DF562 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006DF5CC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006DF635 | `No_Photos_Screen` | Known | Screen layout |
| 0x006DF649 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006DF6AF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DF71D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006DF78A | `NoContent_Screen` | Known | Screen layout |
| 0x006DF79E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006DF806 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006DF870 | `NoContent_Screen` | Known | Screen layout |
| 0x006DF884 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006DF8F1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006DF963 | `NoContent_Screen` | Known | Screen layout |
| 0x006DF977 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006DF9DF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006DFA48 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006DFA63 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006DFAC9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006DFAE5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006DFBC4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006DFBDD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006DFC3E | `FirstBoot_Screen` | Known | Screen layout |
| 0x006DFC52 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006DFDC0 | `Radio_Screen` | Known | Screen layout |
| 0x006DFDD0 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006DFE31 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006DFEB4 | `LockediPod_Screen` | Known | Screen layout |
| 0x006DFF3C | `Lock_Screen` | Known | Screen layout |
| 0x006DFF4B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006DFFAE | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006E0010 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006E002C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006E009E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E00BD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E0125 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E013F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006E01A7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E01C4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E0230 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E029A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006E02B4 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006E0324 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006E0397 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006E0408 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006E0477 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006E04E3 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006E04FE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006E0573 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006E05DA | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006E063C | `Photos_Screen` | Known | Screen layout |
| 0x006E06A0 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006E06BE | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006E072E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E0749 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E07B2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006E07CF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006E0846 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006E086A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006E08D8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006E08F3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006E09B0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E09CC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E0A3A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E0A57 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E0AC2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006E0AE2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006E0B59 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E0B75 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E0BE5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006E0C04 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006E0C70 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006E0C84 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006E0CFD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006E0D71 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006E0DE1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006E0E49 | `NoContent_Screen` | Known | Screen layout |
| 0x006E0E5D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006E0EC1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006E0F28 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E0F42 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006E0FB0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006E1022 | `NoContent_Screen` | Known | Screen layout |
| 0x006E1036 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006E10A0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006E1109 | `No_Photos_Screen` | Known | Screen layout |
| 0x006E111D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006E1183 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E11F1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006E125E | `NoContent_Screen` | Known | Screen layout |
| 0x006E1272 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006E12DA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006E1344 | `NoContent_Screen` | Known | Screen layout |
| 0x006E1358 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006E13C5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006E1437 | `NoContent_Screen` | Known | Screen layout |
| 0x006E144B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E14B3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006E151C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006E1537 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006E159D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006E15B9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006E1698 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006E16B1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006E1712 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006E1726 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006E1894 | `Radio_Screen` | Known | Screen layout |
| 0x006E18A4 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006E1905 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006E1988 | `LockediPod_Screen` | Known | Screen layout |
| 0x006E1A10 | `Lock_Screen` | Known | Screen layout |
| 0x006E1A1F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006E1A82 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006E1AE4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006E1B00 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006E1B72 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E1B91 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E1BF9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E1C13 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006E1C7B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E1C98 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E1D04 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E1D6E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006E1D88 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006E1DF8 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006E1E6B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006E1EDC | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006E1F4B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006E1FB7 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006E1FD2 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006E2047 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006E20AE | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006E2110 | `Photos_Screen` | Known | Screen layout |
| 0x006E2174 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006E2192 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006E2202 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E221D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E2286 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006E22A3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006E231A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006E233E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006E23AC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006E23C7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006E2484 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E24A0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E250E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E252B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E2596 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006E25B6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006E262D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E2649 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E26B9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006E26D8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006E2744 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006E2758 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006E27D1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006E2845 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006E28B5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006E291D | `NoContent_Screen` | Known | Screen layout |
| 0x006E2931 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006E2995 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006E29FC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E2A16 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006E2A84 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006E2AF6 | `NoContent_Screen` | Known | Screen layout |
| 0x006E2B0A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006E2B74 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006E2BDD | `No_Photos_Screen` | Known | Screen layout |
| 0x006E2BF1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006E2C57 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E2CC5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006E2D32 | `NoContent_Screen` | Known | Screen layout |
| 0x006E2D46 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006E2DAE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006E2E18 | `NoContent_Screen` | Known | Screen layout |
| 0x006E2E2C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006E2E99 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006E2F0B | `NoContent_Screen` | Known | Screen layout |
| 0x006E2F1F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E2F87 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006E2FF0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006E300B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006E3071 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006E308D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006E316C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006E3185 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006E31E6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006E31FA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006E3368 | `Radio_Screen` | Known | Screen layout |
| 0x006E3378 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006E33D9 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006E345C | `LockediPod_Screen` | Known | Screen layout |
| 0x006E34E4 | `Lock_Screen` | Known | Screen layout |
| 0x006E34F3 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006E3556 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006E35B8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006E35D4 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006E3646 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E3665 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E36CD | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E36E7 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006E374F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E376C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E37D8 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E3842 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006E385C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006E38CC | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006E393F | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006E39B0 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006E3A1F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006E3A8B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006E3AA6 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006E3B1B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006E3B82 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006E3BE4 | `Photos_Screen` | Known | Screen layout |
| 0x006E3C48 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006E3C66 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006E3CD6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E3CF1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E3D5A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006E3D77 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006E3DEE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006E3E12 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006E3E80 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006E3E9B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006E3F58 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E3F74 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E3FE2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E3FFF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E406A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006E408A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006E4101 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E411D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E418D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006E41AC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006E4218 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006E422C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006E42A5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006E4319 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006E4389 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006E43F1 | `NoContent_Screen` | Known | Screen layout |
| 0x006E4405 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006E4469 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006E44D0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E44EA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006E4558 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006E45CA | `NoContent_Screen` | Known | Screen layout |
| 0x006E45DE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006E4648 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006E46B1 | `No_Photos_Screen` | Known | Screen layout |
| 0x006E46C5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006E472B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E4799 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006E4806 | `NoContent_Screen` | Known | Screen layout |
| 0x006E481A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006E4882 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006E48EC | `NoContent_Screen` | Known | Screen layout |
| 0x006E4900 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006E496D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006E49DF | `NoContent_Screen` | Known | Screen layout |
| 0x006E49F3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E4A5B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006E4AC4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006E4ADF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006E4B45 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006E4B61 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006E4C40 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006E4C59 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006E4CBA | `FirstBoot_Screen` | Known | Screen layout |
| 0x006E4CCE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006E4E3C | `Radio_Screen` | Known | Screen layout |
| 0x006E4E4C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006E4EAD | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006E4F30 | `LockediPod_Screen` | Known | Screen layout |
| 0x006E4FB8 | `Lock_Screen` | Known | Screen layout |
| 0x006E4FC7 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006E502A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006E508C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006E50A8 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006E511A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E5139 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E51A1 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E51BB | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006E5223 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E5240 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E52AC | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E5316 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006E5330 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006E53A0 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006E5413 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006E5484 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006E54F3 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006E555F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006E557A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006E55EF | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006E5656 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006E56B8 | `Photos_Screen` | Known | Screen layout |
| 0x006E571C | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006E573A | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006E57AA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E57C5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E582E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006E584B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006E58C2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006E58E6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006E5954 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006E596F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006E5A2C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E5A48 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E5AB6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E5AD3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E5B3E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006E5B5E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006E5BD5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E5BF1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E5C61 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006E5C80 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006E5CEC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006E5D00 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006E5D79 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006E5DED | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006E5E5D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006E5EC5 | `NoContent_Screen` | Known | Screen layout |
| 0x006E5ED9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006E5F3D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006E5FA4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E5FBE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006E602C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006E609E | `NoContent_Screen` | Known | Screen layout |
| 0x006E60B2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006E611C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006E6185 | `No_Photos_Screen` | Known | Screen layout |
| 0x006E6199 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006E61FF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E626D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006E62DA | `NoContent_Screen` | Known | Screen layout |
| 0x006E62EE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006E6356 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006E63C0 | `NoContent_Screen` | Known | Screen layout |
| 0x006E63D4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006E6441 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006E64B3 | `NoContent_Screen` | Known | Screen layout |
| 0x006E64C7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E652F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006E6598 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006E65B3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006E6619 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006E6635 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006E6714 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006E672D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006E678E | `FirstBoot_Screen` | Known | Screen layout |
| 0x006E67A2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006E6910 | `Radio_Screen` | Known | Screen layout |
| 0x006E6920 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006E6981 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006E6A04 | `LockediPod_Screen` | Known | Screen layout |
| 0x006E6A8C | `Lock_Screen` | Known | Screen layout |
| 0x006E6A9B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006E6AFE | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006E6B60 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006E6B7C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006E6BEE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E6C0D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E6C75 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E6C8F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006E6CF7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E6D14 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E6D80 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E6DEA | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006E6E04 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006E6E74 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006E6EE7 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006E6F58 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006E6FC7 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006E7033 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006E704E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006E70C3 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006E712A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006E718C | `Photos_Screen` | Known | Screen layout |
| 0x006E71F0 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006E720E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006E727E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E7299 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E7302 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006E731F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006E7396 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006E73BA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006E7428 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006E7443 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006E7500 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E751C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E758A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E75A7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E7612 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006E7632 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006E76A9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E76C5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E7735 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006E7754 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006E77C0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006E77D4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x006E784D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006E78C1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x006E7931 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006E7999 | `NoContent_Screen` | Known | Screen layout |
| 0x006E79AD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x006E7A11 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x006E7A78 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E7A92 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x006E7B00 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006E7B72 | `NoContent_Screen` | Known | Screen layout |
| 0x006E7B86 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006E7BF0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x006E7C59 | `No_Photos_Screen` | Known | Screen layout |
| 0x006E7C6D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x006E7CD3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E7D41 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x006E7DAE | `NoContent_Screen` | Known | Screen layout |
| 0x006E7DC2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006E7E2A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x006E7E94 | `NoContent_Screen` | Known | Screen layout |
| 0x006E7EA8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006E7F15 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006E7F87 | `NoContent_Screen` | Known | Screen layout |
| 0x006E7F9B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006E8003 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x006E806C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x006E8087 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006E80ED | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006E8109 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006E81E8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006E8201 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x006E8262 | `FirstBoot_Screen` | Known | Screen layout |
| 0x006E8276 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x006E83E4 | `Radio_Screen` | Known | Screen layout |
| 0x006E83F4 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x006E8455 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x006E84D8 | `LockediPod_Screen` | Known | Screen layout |
| 0x006E8560 | `Lock_Screen` | Known | Screen layout |
| 0x006E856F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x006E85D2 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x006E8634 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006E8650 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x006E86C2 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E86E1 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E8749 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x006E8763 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x006E87CB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E87E8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E8854 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E88BE | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x006E88D8 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x006E8948 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x006E89BB | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006E8A2C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006E8A9B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x006E8B07 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x006E8B22 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x006E8B97 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x006E8BFE | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x006E8C60 | `Photos_Screen` | Known | Screen layout |
| 0x006E8CC4 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x006E8CE2 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006E8D52 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E8D6D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E8DD6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006E8DF3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006E8E6A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006E8E8E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006E8EFC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x006E8F17 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006E8FB9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E8FD5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E9043 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E9060 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E90CB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006E90EB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006E9162 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E917E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E91EE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006E920D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006E9279 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006E928D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x006E9302 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x006E936D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x006E93DC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x006E944D | `NoContent_Screen` | Known | Screen layout |
| 0x006E9461 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006E94D0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x006E9543 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x006E95B0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x006E9619 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x006E9689 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x006E96F9 | `NoContent_Screen` | Known | Screen layout |
| 0x006E970D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006E9770 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x006E97D3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006E97EF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006E98AF | `Radio_Screen` | Known | Screen layout |
| 0x006E98BF | `Radio_Screen_Default` | Known | Screen layout |
| 0x006E9920 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x006E998E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006E99AD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006E9A1B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x006E9A80 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006E9A9B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006E9B41 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E9B5D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E9BCB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006E9BE8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006E9C53 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006E9C73 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006E9CEA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006E9D06 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006E9D76 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006E9D95 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006E9E01 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006E9E15 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x006E9E8A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x006E9EF5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x006E9F64 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x006E9FD5 | `NoContent_Screen` | Known | Screen layout |
| 0x006E9FE9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006EA058 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x006EA0CB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x006EA138 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x006EA1A1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x006EA211 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x006EA281 | `NoContent_Screen` | Known | Screen layout |
| 0x006EA295 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006EA2F8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x006EA35B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006EA377 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006EA437 | `Radio_Screen` | Known | Screen layout |
| 0x006EA447 | `Radio_Screen_Default` | Known | Screen layout |
| 0x006EA4A8 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x006EA516 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006EA535 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006EA5A3 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x006EA608 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006EA623 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006EA6C9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EA6E5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EA753 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006EA770 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006EA7DB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006EA7FB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006EA872 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EA88E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EA8FE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006EA91D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006EA989 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006EA99D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x006EAA12 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x006EAA7D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x006EAAEC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x006EAB5D | `NoContent_Screen` | Known | Screen layout |
| 0x006EAB71 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006EABE0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x006EAC53 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x006EACC0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x006EAD29 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x006EAD99 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x006EAE09 | `NoContent_Screen` | Known | Screen layout |
| 0x006EAE1D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006EAE80 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x006EAEE3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006EAEFF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006EAFBF | `Radio_Screen` | Known | Screen layout |
| 0x006EAFCF | `Radio_Screen_Default` | Known | Screen layout |
| 0x006EB030 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x006EB09E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006EB0BD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006EB12B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x006EB190 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006EB1AB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006EB251 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EB26D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EB2DB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006EB2F8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006EB363 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006EB383 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006EB3FA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EB416 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EB486 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006EB4A5 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006EB511 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006EB525 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x006EB59A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x006EB605 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x006EB674 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x006EB6E5 | `NoContent_Screen` | Known | Screen layout |
| 0x006EB6F9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006EB768 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x006EB7DB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x006EB848 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x006EB8B1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x006EB921 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x006EB991 | `NoContent_Screen` | Known | Screen layout |
| 0x006EB9A5 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006EBA08 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x006EBA6B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006EBA87 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006EBB47 | `Radio_Screen` | Known | Screen layout |
| 0x006EBB57 | `Radio_Screen_Default` | Known | Screen layout |
| 0x006EBBB8 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x006EBC26 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006EBC45 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006EBCB3 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x006EBD18 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006EBD33 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006EBDD9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EBDF5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EBE63 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006EBE80 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006EBEEB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006EBF0B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006EBF82 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EBF9E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EC00E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006EC02D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006EC099 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006EC0AD | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x006EC122 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x006EC18D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x006EC1FC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x006EC26D | `NoContent_Screen` | Known | Screen layout |
| 0x006EC281 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006EC2F0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x006EC363 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x006EC3D0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x006EC439 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x006EC4A9 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x006EC519 | `NoContent_Screen` | Known | Screen layout |
| 0x006EC52D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006EC590 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x006EC5F3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006EC60F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006EC6CF | `Radio_Screen` | Known | Screen layout |
| 0x006EC6DF | `Radio_Screen_Default` | Known | Screen layout |
| 0x006EC740 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x006EC7AE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006EC7CD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006EC83B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x006EC8A0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006EC8BB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006EC961 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EC97D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EC9EB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006ECA08 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006ECA73 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006ECA93 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006ECB0A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006ECB26 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006ECB96 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006ECBB5 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006ECC21 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006ECC35 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x006ECCAA | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x006ECD15 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x006ECD84 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x006ECDF5 | `NoContent_Screen` | Known | Screen layout |
| 0x006ECE09 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006ECE78 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x006ECEEB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x006ECF58 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x006ECFC1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x006ED031 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x006ED0A1 | `NoContent_Screen` | Known | Screen layout |
| 0x006ED0B5 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006ED118 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x006ED17B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006ED197 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006ED257 | `Radio_Screen` | Known | Screen layout |
| 0x006ED267 | `Radio_Screen_Default` | Known | Screen layout |
| 0x006ED2C8 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x006ED336 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006ED355 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006ED3C3 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x006ED428 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006ED443 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006ED4E9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006ED505 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006ED573 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006ED590 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006ED5FB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006ED61B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006ED692 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006ED6AE | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006ED71E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006ED73D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006ED7A9 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006ED7BD | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x006ED832 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x006ED89D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x006ED90C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x006ED97D | `NoContent_Screen` | Known | Screen layout |
| 0x006ED991 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006EDA00 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x006EDA73 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x006EDAE0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x006EDB49 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x006EDBB9 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x006EDC29 | `NoContent_Screen` | Known | Screen layout |
| 0x006EDC3D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006EDCA0 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x006EDD03 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006EDD1F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006EDDDF | `Radio_Screen` | Known | Screen layout |
| 0x006EDDEF | `Radio_Screen_Default` | Known | Screen layout |
| 0x006EDE50 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x006EDEBE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006EDEDD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006EDF4B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x006EDFB0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006EDFCB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006EE071 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EE08D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EE0FB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006EE118 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006EE183 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006EE1A3 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006EE21A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EE236 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EE2A6 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006EE2C5 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006EE331 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006EE345 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x006EE3BA | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x006EE425 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x006EE494 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x006EE505 | `NoContent_Screen` | Known | Screen layout |
| 0x006EE519 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006EE588 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x006EE5FB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x006EE668 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x006EE6D1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x006EE741 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x006EE7B1 | `NoContent_Screen` | Known | Screen layout |
| 0x006EE7C5 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006EE828 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x006EE88B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006EE8A7 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006EE967 | `Radio_Screen` | Known | Screen layout |
| 0x006EE977 | `Radio_Screen_Default` | Known | Screen layout |
| 0x006EE9D8 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x006EEA46 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006EEA65 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006EEAD3 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x006EEB38 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006EEB53 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006EEBF9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EEC15 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EEC83 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006EECA0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006EED0B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006EED2B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006EEDA2 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EEDBE | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EEE2E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006EEE4D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006EEEB9 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006EEECD | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x006EEF42 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x006EEFAD | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x006EF01C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x006EF08D | `NoContent_Screen` | Known | Screen layout |
| 0x006EF0A1 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006EF110 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x006EF183 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x006EF1F0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x006EF259 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x006EF2C9 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x006EF339 | `NoContent_Screen` | Known | Screen layout |
| 0x006EF34D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006EF3B0 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x006EF413 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006EF42F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006EF4EF | `Radio_Screen` | Known | Screen layout |
| 0x006EF4FF | `Radio_Screen_Default` | Known | Screen layout |
| 0x006EF560 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x006EF5CE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006EF5ED | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006EF65B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x006EF6C0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006EF6DB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006EF781 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EF79D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EF80B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006EF828 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006EF893 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006EF8B3 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006EF92A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006EF946 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006EF9B6 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006EF9D5 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006EFA41 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006EFA55 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x006EFACA | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x006EFB35 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x006EFBA4 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x006EFC15 | `NoContent_Screen` | Known | Screen layout |
| 0x006EFC29 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006EFC98 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x006EFD0B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x006EFD78 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x006EFDE1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x006EFE51 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x006EFEC1 | `NoContent_Screen` | Known | Screen layout |
| 0x006EFED5 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006EFF38 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x006EFF9B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006EFFB7 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006F0077 | `Radio_Screen` | Known | Screen layout |
| 0x006F0087 | `Radio_Screen_Default` | Known | Screen layout |
| 0x006F00E8 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x006F0156 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006F0175 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006F01E3 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x006F0248 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006F0263 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006F0309 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F0325 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F0393 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F03B0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F041B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006F043B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006F04B2 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F04CE | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F053E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006F055D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006F05C9 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006F05DD | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x006F0652 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x006F06BD | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x006F072C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x006F079D | `NoContent_Screen` | Known | Screen layout |
| 0x006F07B1 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006F0820 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x006F0893 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x006F0900 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x006F0969 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x006F09D9 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x006F0A49 | `NoContent_Screen` | Known | Screen layout |
| 0x006F0A5D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006F0AC0 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x006F0B23 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006F0B3F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006F0BFF | `Radio_Screen` | Known | Screen layout |
| 0x006F0C0F | `Radio_Screen_Default` | Known | Screen layout |
| 0x006F0C70 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x006F0CDE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006F0CFD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006F0D6B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x006F0DD0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006F0DEB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006F0E91 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F0EAD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F0F1B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F0F38 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F0FA3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x006F0FC3 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006F103A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F1056 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x006F10C6 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x006F10E5 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x006F1151 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x006F1165 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x006F11DA | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x006F1245 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x006F12B4 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x006F1325 | `NoContent_Screen` | Known | Screen layout |
| 0x006F1339 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x006F13A8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x006F141B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x006F1488 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x006F14F1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x006F1561 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x006F15D1 | `NoContent_Screen` | Known | Screen layout |
| 0x006F15E5 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006F1648 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x006F16AB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x006F16C7 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x006F1787 | `Radio_Screen` | Known | Screen layout |
| 0x006F1797 | `Radio_Screen_Default` | Known | Screen layout |
| 0x006F17F8 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x006F1866 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x006F1885 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006F18F3 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x006F1958 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006F1973 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006F1A54 | `MediaLists_AudiobookChapters_Screen)` | Known | Screen layout |
| 0x006F1A7B | `MediaLists_AudiobookChapters_Screen_Plane!` | Known | Screen layout |
| 0x006F2215 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x006F2230 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x006F229B | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006F22B6 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x006F2329 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x006F2344 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x006F2501 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x006F251C | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x006F2587 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006F25A2 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x006F2615 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x006F2630 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x006F27F8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F2814 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x006F288F | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F28AB | `MediaLists_Albums_Screen_WithTrackCount%` | Known | Screen layout |
| 0x006F2924 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x006F293F | `MediaLists_Songs_Screen_Plain1` | Known | Screen layout |
| 0x006F29BA | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x006F29D5 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x006F2BF7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F2C14 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F2CF3 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x006F2D0F | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x006F2D8A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x006F2DA5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x006F2F8B | `MediaLists_PodcastEpisodes_Screen'` | Known | Screen layout |
| 0x006F2FB0 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x006F3282 | `MediaLists_TVSeasons_Screen#` | Known | Screen layout |
| 0x006F32A1 | `MediaLists_TVSeasons_Screen_Default%` | Known | Screen layout |
| 0x006F3316 | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x006F3336 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x006F34BE | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x006F34DE | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x006F38D7 | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x006F38FC | `MediaLists_NestedPlaylists_Screen_Default,` | Known | Screen layout |
| 0x006F397E | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x006F399D | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x006F3B2D | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x006F3B52 | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x006F3BCA | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x006F3BE9 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x006F3C4D | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F3CFA | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F3D6C | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x006F3E62 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x006F4004 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x006F4104 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x006F4170 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x006F41DA | `NoContent_Screen` | Known | Screen layout |
| 0x006F41EE | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x006F4258 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006F42CC | `NoContent_Screen` | Known | Screen layout |
| 0x006F42E0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006F434B | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x006F43B7 | `NoContent_Screen` | Known | Screen layout |
| 0x006F43CB | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F4438 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F44AC | `NoContent_Screen` | Known | Screen layout |
| 0x006F44C0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F4528 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F4595 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F45F9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F4615 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x006F4681 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F469E | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F470B | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F4775 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F4792 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F4809 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F482D | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F48E4 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x006F494E | `NoContent_Screen` | Known | Screen layout |
| 0x006F4962 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x006F49CC | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006F4A40 | `NoContent_Screen` | Known | Screen layout |
| 0x006F4A54 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006F4ABF | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x006F4B2B | `NoContent_Screen` | Known | Screen layout |
| 0x006F4B3F | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F4BAC | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F4C20 | `NoContent_Screen` | Known | Screen layout |
| 0x006F4C34 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F4C9C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F4D09 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F4D6D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F4D89 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x006F4DF5 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F4E12 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F4E7F | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F4EE9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F4F06 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F4F7D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F4FA1 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F5058 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x006F50C2 | `NoContent_Screen` | Known | Screen layout |
| 0x006F50D6 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x006F5140 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006F51B4 | `NoContent_Screen` | Known | Screen layout |
| 0x006F51C8 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006F5233 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x006F529F | `NoContent_Screen` | Known | Screen layout |
| 0x006F52B3 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F5320 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F5394 | `NoContent_Screen` | Known | Screen layout |
| 0x006F53A8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F5410 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F547D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F54E1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F54FD | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x006F5569 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F5586 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F55F3 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F565D | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F567A | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F56F1 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F5715 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F57CC | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x006F5836 | `NoContent_Screen` | Known | Screen layout |
| 0x006F584A | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x006F58B4 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006F5928 | `NoContent_Screen` | Known | Screen layout |
| 0x006F593C | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006F59A7 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x006F5A13 | `NoContent_Screen` | Known | Screen layout |
| 0x006F5A27 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F5A94 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F5B08 | `NoContent_Screen` | Known | Screen layout |
| 0x006F5B1C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F5B84 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F5BF1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F5C55 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F5C71 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x006F5CDD | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F5CFA | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F5D67 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F5DD1 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F5DEE | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F5E65 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F5E89 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F5F40 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x006F5FAA | `NoContent_Screen` | Known | Screen layout |
| 0x006F5FBE | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x006F6028 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006F609C | `NoContent_Screen` | Known | Screen layout |
| 0x006F60B0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006F611B | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x006F6187 | `NoContent_Screen` | Known | Screen layout |
| 0x006F619B | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F6208 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F627C | `NoContent_Screen` | Known | Screen layout |
| 0x006F6290 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F62F8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F6365 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F63C9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F63E5 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x006F6451 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F646E | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F64DB | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F6545 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F6562 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F65D9 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F65FD | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F66B4 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x006F671E | `NoContent_Screen` | Known | Screen layout |
| 0x006F6732 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x006F679C | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006F6810 | `NoContent_Screen` | Known | Screen layout |
| 0x006F6824 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006F688F | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x006F68FB | `NoContent_Screen` | Known | Screen layout |
| 0x006F690F | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F697C | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F69F0 | `NoContent_Screen` | Known | Screen layout |
| 0x006F6A04 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F6A6C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F6AD9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F6B3D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F6B59 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x006F6BC5 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F6BE2 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F6C4F | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F6CB9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F6CD6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F6D4D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F6D71 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F6E28 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x006F6E92 | `NoContent_Screen` | Known | Screen layout |
| 0x006F6EA6 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x006F6F10 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x006F6F84 | `NoContent_Screen` | Known | Screen layout |
| 0x006F6F98 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x006F7003 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x006F706F | `NoContent_Screen` | Known | Screen layout |
| 0x006F7083 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x006F70F0 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x006F7164 | `NoContent_Screen` | Known | Screen layout |
| 0x006F7178 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x006F71E0 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x006F724D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006F72B1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x006F72CD | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x006F7339 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x006F7356 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x006F73C3 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x006F742D | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x006F744A | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x006F74C1 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x006F74E5 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x006F785C | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006F78CE | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006F7939 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006F799E | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006F7A08 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006F7A72 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006F7AD9 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006F7B44 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006F7BAE | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006F7C15 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006F7C7C | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006F7CE1 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006F7D49 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006F7DB4 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006F7E1F | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006F7E86 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006F814C | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006F81BE | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006F8229 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006F828E | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006F82F8 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006F8362 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006F83C9 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006F8434 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006F849E | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006F8505 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006F856C | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006F85D1 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006F8639 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006F86A4 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006F870F | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006F8776 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006F8A3A | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006F8AAC | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006F8B17 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006F8B7C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006F8BE6 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006F8C50 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006F8CB7 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006F8D22 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006F8D8C | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006F8DF3 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006F8E5A | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006F8EBF | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006F8F27 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006F8F92 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006F8FFD | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006F9064 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006F9326 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006F9398 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006F9403 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006F9468 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006F94D2 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006F953C | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006F95A3 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006F960E | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006F9678 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006F96DF | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006F9746 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006F97AB | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006F9813 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006F987E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006F98E9 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006F9950 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006F9BFA | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006F9C6C | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006F9CD7 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006F9D3C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006F9DA6 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006F9E10 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006F9E77 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006F9EE2 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006F9F4C | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006F9FB3 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006FA01A | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006FA07F | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006FA0E7 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006FA152 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006FA1BD | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006FA224 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006FA4F3 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006FA565 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006FA5D0 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006FA635 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006FA69F | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006FA709 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006FA770 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006FA7DB | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006FA845 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006FA8AC | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006FA913 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006FA978 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006FA9E0 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006FAA4B | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006FAAB6 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006FAB1D | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006FADE9 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006FAE5B | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006FAEC6 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006FAF2B | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006FAF95 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006FAFFF | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006FB066 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006FB0D1 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006FB13B | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006FB1A2 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006FB209 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006FB26E | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006FB2D6 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006FB341 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006FB3AC | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006FB413 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006FB6CA | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006FB73C | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006FB7A7 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006FB80C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006FB876 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006FB8E0 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006FB947 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006FB9B2 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006FBA1C | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006FBA83 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006FBAEA | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006FBB4F | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006FBBB7 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006FBC22 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006FBC8D | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006FBCF4 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006FBFD0 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006FC042 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006FC0AD | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006FC112 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006FC17C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006FC1E6 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006FC24D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006FC2B8 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006FC322 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006FC389 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006FC3F0 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006FC455 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006FC4BD | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006FC528 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006FC593 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006FC5FA | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006FC8E4 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006FC956 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006FC9C1 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006FCA26 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006FCA90 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006FCAFA | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006FCB61 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006FCBCC | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006FCC36 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006FCC9D | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006FCD04 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006FCD69 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006FCDD1 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006FCE3C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006FCEA7 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006FCF0E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006FD1D8 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006FD24A | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006FD2B5 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006FD31A | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006FD384 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006FD3EE | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006FD455 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006FD4C0 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006FD52A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006FD591 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006FD5F8 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006FD65D | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006FD6C5 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006FD730 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006FD79B | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006FD802 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006FDAC0 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006FDB32 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006FDB9D | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006FDC02 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006FDC6C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006FDCD6 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006FDD3D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006FDDA8 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006FDE12 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006FDE79 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006FDEE0 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006FDF45 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006FDFAD | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006FE018 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006FE083 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006FE0EA | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006FE396 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006FE408 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006FE473 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006FE4D8 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006FE542 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006FE5AC | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006FE613 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006FE67E | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006FE6E8 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006FE74F | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006FE7B6 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006FE81B | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006FE883 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006FE8EE | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006FE959 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006FE9C0 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006FEC63 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006FECD5 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006FED40 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006FEDA5 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006FEE0F | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006FEE79 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006FEEE0 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006FEF4B | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006FEFB5 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006FF01C | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006FF083 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006FF0E8 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006FF150 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006FF1BB | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006FF226 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006FF28D | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006FF54B | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x006FF5BD | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x006FF628 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x006FF68D | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x006FF6F7 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x006FF761 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x006FF7C8 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x006FF833 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x006FF89D | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x006FF904 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x006FF96B | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x006FF9D0 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x006FFA38 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x006FFAA3 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x006FFB0E | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x006FFB75 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x006FFE2E | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x006FFEA1 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x006FFF13 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x006FFF83 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x006FFFF1 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x0070005E | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007002FE | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x00700371 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007003E3 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x00700453 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x007004C1 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x0070052E | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007007F2 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x00700863 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007008D3 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x00700941 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007009AE | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x00700C42 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x00700CB3 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x00700D23 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x00700D91 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x00700DFE | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x00701090 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x00701101 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x00701171 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x007011DF | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x0070124C | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007015F2 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x0070160F | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x0070168A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007016A3 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x0070171B | `Notes_Contents_Screen` | Known | Screen layout |
| 0x00701734 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007017A9 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007017BF | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x00701836 | `Notes_Image_Screen` | Known | Screen layout |
| 0x0070184C | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007018C3 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007018E0 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x00701958 | `Notes_List_Screen` | Known | Screen layout |
| 0x0070196D | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x00701B1E | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x00701B3B | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x00701BB6 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x00701BCF | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x00701C47 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x00701C60 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x00701CD5 | `Notes_Image_Screen` | Known | Screen layout |
| 0x00701CEB | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x00701D62 | `Notes_Image_Screen` | Known | Screen layout |
| 0x00701D78 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x00701DEF | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x00701E0C | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x00701E84 | `Notes_List_Screen` | Known | Screen layout |
| 0x00701E99 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x0070207A | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x00702097 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x00702112 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x0070212B | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007021A3 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007021BC | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x00702231 | `Notes_Image_Screen` | Known | Screen layout |
| 0x00702247 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007022BE | `Notes_Image_Screen` | Known | Screen layout |
| 0x007022D4 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x0070234B | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x00702368 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007023E0 | `Notes_List_Screen` | Known | Screen layout |
| 0x007023F5 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007025AA | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007025C7 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x00702642 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x0070265B | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007026D3 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007026EC | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x00702761 | `Notes_Image_Screen` | Known | Screen layout |
| 0x00702777 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007027EE | `Notes_Image_Screen` | Known | Screen layout |
| 0x00702804 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x0070287B | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x00702898 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x00702910 | `Notes_List_Screen` | Known | Screen layout |
| 0x00702925 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x00702C3D | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x00702CE3 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x00702D66 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x00702E1E | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x00702EA0 | `PhotosSettingsSlideshowMusic_Screen+` | Known | Screen layout |
| 0x00702EC7 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x00702FAD | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x00703165 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007031C5 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x00703222 | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x00703249 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007032E9 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x00703349 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007033A6 | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007033CD | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x00703668 | `Photos_Screen` | Known | Screen layout |
| 0x007037B4 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x00703818 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x00703879 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007038D6 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x00703933 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007039A1 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007039FE | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x00703B74 | `Photos_Screen` | Known | Screen layout |
| 0x00703CC0 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x00703D24 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x00703D85 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x00703DE2 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x00703E3F | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x00703EAD | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x00703F0A | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x00704080 | `Photos_Screen` | Known | Screen layout |
| 0x007041CC | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x00704230 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x00704291 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007042EE | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x0070434B | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007043B9 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x00704416 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x0070458C | `Photos_Screen` | Known | Screen layout |
| 0x007046D8 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0070473C | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x0070479D | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007047FA | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x00704857 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007048C5 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x00704922 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x00704A98 | `Photos_Screen` | Known | Screen layout |
| 0x00704BE4 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x00704C48 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x00704CA9 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x00704D06 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x00704D63 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x00704DD1 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x00704E2E | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x00704FA4 | `Photos_Screen` | Known | Screen layout |
| 0x007050F0 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x00705154 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007051B5 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x00705212 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x0070526F | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007052DD | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x0070533A | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007054B0 | `Photos_Screen` | Known | Screen layout |
| 0x007055FC | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x00705662 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007056C4 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x00705726 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007057BC | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007058DD | `Photos_Screen` | Known | Screen layout |
| 0x00705948 | `Photos_Screen` | Known | Screen layout |
| 0x00705A94 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x00705AFA | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x00705B5C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x00705BBE | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x00705C54 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x00705D75 | `Photos_Screen` | Known | Screen layout |
| 0x00705DE0 | `Photos_Screen` | Known | Screen layout |
| 0x00705F2C | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x00705F92 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x00705FF4 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x00706056 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007060EC | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x0070620D | `Photos_Screen` | Known | Screen layout |
| 0x00706278 | `Photos_Screen` | Known | Screen layout |
| 0x007063C4 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0070642A | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x0070648C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007064EE | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x00706584 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007066A5 | `Photos_Screen` | Known | Screen layout |
| 0x00706710 | `Photos_Screen` | Known | Screen layout |
| 0x0070685C | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007068C2 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x00706924 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x00706986 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x00706A1C | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x00706B3D | `Photos_Screen` | Known | Screen layout |
| 0x00706D1E | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x00706D80 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x00706DEE | `Radio_Screen_Default#` | Known | Screen layout |
| 0x00706E54 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x00706EB9 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x00707159 | `Radio_Screen_Default$` | Known | Screen layout |
| 0x007071C0 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x00707226 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x00707525 | `Radio_Screen_Default$` | Known | Screen layout |
| 0x0070758C | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007075F2 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x0070789A | `Radio_Screen_Default$` | Known | Screen layout |
| 0x00707901 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x00707967 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x00707C4E | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x00707CB8 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x00707EC2 | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x00707F2C | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x00708085 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007080E8 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x0070814D | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007081B5 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x00708218 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00708280 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007082E9 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0070834F | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007083B4 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00708439 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007084C6 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00708565 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070857F | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007085F7 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00708611 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00708685 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00708712 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007087B1 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007087CB | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00708843 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070885D | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007088D1 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0070895E | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007089FD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00708A17 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00708A8F | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00708AA9 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00708B1D | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00708BAA | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00708C49 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00708C63 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00708CDB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00708CF5 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00708D69 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00708DF6 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00708E95 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00708EAF | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00708F27 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00708F41 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00708FB5 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00709042 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007090E1 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007090FB | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00709173 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070918D | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00709201 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0070928E | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0070932D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00709347 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007093BF | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007093D9 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0070944D | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007094DA | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00709579 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00709593 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x0070960B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00709625 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00709699 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00709726 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007097C5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007097DF | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00709857 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00709871 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007098E5 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00709972 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00709A11 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00709A2B | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00709AA3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00709ABD | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00709B31 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00709BBE | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00709C5D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00709C77 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00709CEF | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00709D09 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00709D7D | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00709E0A | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00709EA9 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00709EC3 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x00709F3B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00709F55 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x00709FC9 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0070A056 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0070A0F5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070A10F | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x0070A187 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070A1A1 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0070A215 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0070A2A2 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0070A341 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070A35B | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x0070A3D3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070A3ED | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0070A461 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0070A4EE | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0070A58D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070A5A7 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x0070A61F | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070A639 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0070A6AD | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0070A73A | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0070A7D9 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070A7F3 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x0070A86B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070A885 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0070A8F9 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0070A986 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0070AA25 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070AA3F | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x0070AAB7 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070AAD1 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0070AB45 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0070ABD2 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0070AC71 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070AC8B | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x0070AD03 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070AD1D | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0070AD91 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0070AE1E | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0070AEBD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070AED7 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x0070AF4F | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070AF69 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0070AFE5 | `SettingsMenus_ResetMainMenu_Screen9` | Known | Screen layout |
| 0x0070B0B5 | `SettingsMenus_ResetMusicMenu_Screen:` | Known | Screen layout |
| 0x0070B169 | `SettingsMenus_Main_Screen,` | Known | Screen layout |
| 0x0070B1DB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070B1F5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x0070B26D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0070B287 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0070B5C2 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x0070B628 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x0070B685 | `Extras_Screen` | Known | Screen layout |
| 0x0070B6D9 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x0070B7B7 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x0070B825 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x0070B8C3 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x0070B8DC | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x0070B944 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x0070BA59 | `VoiceMemos_Context_PlayDelete_Screen)` | Known | Screen layout |
| 0x0070BA81 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x0070BAF8 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0070BBC4 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0070BC33 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0070BD21 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0070BD8A | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x0070BDAC | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x0070BE18 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x0070BE3A | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x0070BFB6 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070BFD2 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x0070C099 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0070C0B4 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x0070C117 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x0070C17A | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x0070C211 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070C22D | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x0070C2F4 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0070C30F | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x0070C372 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x0070C3D5 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x0070C46D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0070C489 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x0070C550 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0070C56B | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x0070C5CE | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x0070C631 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x0070C6AE | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x0070C719 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x0070C785 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x0070C7F7 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x0070C864 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x0070C8CF | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x0070C93B | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x0070C9A3 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x0070CA0F | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x0070CA83 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x0070CAF1 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x0070CB6A | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x007268E0 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x00726965 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x00726BFE | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x008B4A92 | `VoiceMemos_Screen_InsufficientDiskSpace2` | Known | Screen layout |
| 0x008B62C2 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x008B62DA | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x008B62F8 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x008B6404 | `NowPlaying_Screen_Music` | Known | Screen layout |
| 0x008B6430 | `MainMenu_Music_Screen_NoMusic` | Known | Screen layout |
| 0x008B644E | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x008B646C | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x008B656D | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x008B65F6 | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x008B6642 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x008B674A | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x008B6763 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x008B6781 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x008B67B0 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x008B67E8 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x008B6C1F | `Video_Settings_TV_Screen_Standard` | Known | Screen layout |
| 0x008B6C51 | `Search_Main_Screen_WithKeyboard` | Known | Screen layout |
| 0x008B6C71 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x008B6CB6 | `VoiceMemos_Screen_InsufficientDiskSpace` | Known | Screen layout |
| 0x008B6D7A | `Video_Settings_TV_Screen_Wide` | Known | Screen layout |
| 0x008B6DC2 | `CoverFlow_Screen_Backside` | Known | Screen layout |
| 0x008B96DB | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x008B98E0 | `NowPlaying_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x008B9905 | `Slideshow_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x008B99D5 | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x008B99EF | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x008B9AFA | `VoiceMemos_DeleteAll_ScreenTitle` | Known | Screen layout |
| 0x008B9B97 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x008B9BDA | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x008B9DCB | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x008B9EB4 | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x008B9ECD | `Radio_Screen_Volume` | Known | Screen layout |
| 0x008B9EE1 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x008B9EFE | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x008B9F1D | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x008B9FE9 | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x008BA13F | `CoverFlow_Screen_BacksideForCapture` | Known | Screen layout |
| 0x008BAFDD | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x008BAFF8 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x008BB2C3 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x008BB311 | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x008BB423 | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x008BB573 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x008C0B6A | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x008C0B95 | `TVOutSlideshow_Screen_Playing` | Known | Screen layout |
| 0x008C0BB3 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x008C0BED | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x008C0C8A | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x008C0CF5 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x008C0E57 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x008C0E77 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x008C13C2 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x008C13DD | `Extras_Screen_Lock` | Known | Screen layout |
| 0x008C13F0 | `MainMenu_List_ScreenLock` | Known | Screen layout |
| 0x008C1409 | `ExtrasMenu_ScreenLock` | Known | Screen layout |
| 0x008C147C | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x008C149D | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x008C1570 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x008C1592 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x008C1699 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x008C16D9 | `VoiceMemos_Screen_DeletAllAsk` | Known | Screen layout |
| 0x008C16F7 | `NowPlaying_Screen_TVOutAsk` | Known | Screen layout |
| 0x008C1853 | `NowPlaying_Screen_Initial` | Known | Screen layout |
| 0x008C186D | `NowPlaying_Screen_Video_TVOut_Initial` | Known | Screen layout |
| 0x008C2574 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x008C25F5 | `RemoteUI_Screen` | Known | Screen layout |
| 0x008C2605 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x008C261D | `Video_Settings_TV_Screen` | Known | Screen layout |
| 0x008C2636 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x008C264D | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x008C2671 | `AddressViewer_PartialLoad_Screen` | Known | Screen layout |
| 0x008C2692 | `Alarms_Alarm_Clock_Triggered_Screen` | Known | Screen layout |
| 0x008C26B6 | `Alarms_Alarm_Triggered_Screen` | Known | Screen layout |
| 0x008C26D4 | `Unsupported_Screen` | Known | Screen layout |
| 0x008C26E7 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x008C2705 | `LockediPod_Screen` | Known | Screen layout |
| 0x008C2717 | `DiskMode_Screen` | Known | Screen layout |
| 0x008C2727 | `DemoMode_Screen` | Known | Screen layout |
| 0x008C2737 | `Notes_Image_Screen` | Known | Screen layout |
| 0x008C274A | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x008C2768 | `NowPlaying_Idle_Screen` | Known | Screen layout |
| 0x008C277F | `Game_Screen` | Known | Screen layout |
| 0x008C278B | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x008C27A8 | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x008C27C1 | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x008C27E2 | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x008C2807 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x008C281A | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x008C2837 | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x008C2858 | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x008C287D | `Notes_Loading_Screen` | Known | Screen layout |
| 0x008C2892 | `SettingsMenus_AdjustScrolling_Screen` | Known | Screen layout |
| 0x008C28B7 | `Game_Running_Screen` | Known | Screen layout |
| 0x008C28CB | `Stopwatch_Screen` | Known | Screen layout |
| 0x008C28DC | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x008C28F3 | `Clock_Screen` | Known | Screen layout |
| 0x008C2900 | `TVOutSlideshowAsk_Screen` | Known | Screen layout |
| 0x008C2919 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x008C292F | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x008C294D | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x008C2969 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x008C297A | `DemoMode_Main_Screen` | Known | Screen layout |
| 0x008C298F | `Search_Main_Screen` | Known | Screen layout |
| 0x008C29A2 | `AddressViewer_Main_Screen` | Known | Screen layout |
| 0x008C29BC | `Speakers_Main_Screen` | Known | Screen layout |
| 0x008C29D1 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x008C29E7 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x008C2A01 | `Clock_Region_Screen` | Known | Screen layout |
| 0x008C2A15 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x008C2A2D | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x008C2A4B | `Video_Settings_Fit_To_Screen` | Known | Screen layout |
| 0x008C2A68 | `Radio_Screen` | Known | Screen layout |
| 0x008C2A75 | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x008C2A8F | `ChargingMode_LowPower_Screen` | Known | Screen layout |
| 0x008C2AAC | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x008C2AC6 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x008C2AE0 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x008C2AFA | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x008C2B13 | `Extras_Screen` | Known | Screen layout |
| 0x008C2B21 | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x008C2B3E | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x008C2B60 | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x008C2B79 | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x008C2B97 | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x008C2BB0 | `Video_Settings_Screen` | Known | Screen layout |
| 0x008C2BC6 | `SettingsMenus_AudiobookSettings_Screen` | Known | Screen layout |
| 0x008C2BED | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x008C2C13 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x008C2C29 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x008C2C41 | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x008C2C64 | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x008C2C81 | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x008C2CA5 | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x008C2CBE | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x008C2CE0 | `ToDo_Instructions_Screen` | Known | Screen layout |
| 0x008C2CF9 | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x008C2D15 | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x008C2D2F | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x008C2D50 | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x008C2D6C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x008C2D84 | `VoiceMemos_Screen` | Known | Screen layout |
| 0x008C2D96 | `No_Photos_Screen` | Known | Screen layout |
| 0x008C2DA7 | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x008C2DC1 | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x008C2DDD | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x008C2E01 | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x008C2E21 | `ChargingMode_Progress_Screen` | Known | Screen layout |
| 0x008C2E3E | `Notes_Contents_Screen` | Known | Screen layout |
| 0x008C2E54 | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x008C2E6F | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x008C2E8B | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x008C2EAD | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x008C2ECE | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x008C2EE8 | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x008C2F02 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x008C2F21 | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x008C2F42 | `NoContent_Screen` | Known | Screen layout |
| 0x008C2F53 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x008C2F69 | `FirstBoot_Screen` | Known | Screen layout |
| 0x008C2F7A | `Notes_List_Screen` | Known | Screen layout |
| 0x008C2F8C | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x008C2FAD | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x008C2FC7 | `TODOS_Menu_Screen` | Known | Screen layout |
| 0x008C2FD9 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x008C2FEF | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x008C300B | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x008C3020 | `Games_Menu_Screen` | Known | Screen layout |
| 0x008C3032 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x008C3045 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x008C3064 | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x008C3083 | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x008C30A7 | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x008C30C5 | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x008C30E8 | `TVOutSlideshow_Screen` | Known | Screen layout |
| 0x008C30FE | `CoverFlow_Screen` | Known | Screen layout |
| 0x008C310F | `Calendar_Day_Screen` | Known | Screen layout |
| 0x008C3123 | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x008C3145 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x008C315D | `Stopwatch_SessionSummary_Screen` | Known | Screen layout |
| 0x008C317D | `Clock_NorthAmerica_City_Screen` | Known | Screen layout |
| 0x008C319C | `Clock_SouthAmerica_City_Screen` | Known | Screen layout |
| 0x008C31BB | `Clock_Africa_City_Screen` | Known | Screen layout |
| 0x008C31D4 | `Clock_Australia_City_Screen` | Known | Screen layout |
| 0x008C31F0 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x008C3207 | `Clock_Pacific_City_Screen` | Known | Screen layout |
| 0x008C3221 | `Clock_Atlantic_City_Screen` | Known | Screen layout |
| 0x008C323C | `Clock_Europe_City_Screen` | Known | Screen layout |
| 0x008C331C | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x008C336D | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x008C3390 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x008C33B8 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x008C36E8 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x008C3B05 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x008C3B5B | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x008C3C8F | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x008C3CAC | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x008C403F | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x008C4155 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x008C4177 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x008C41E4 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x008C4203 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x008C47F4 | `ChargingMode_ScreenLayout_LowPower` | Known | Screen layout |
| 0x008C50E6 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x008C51FD | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x008C52D9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x008C52F7 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x008C5317 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x008C53EB | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x008C5407 | `Extras_Screen_Games` | Known | Screen layout |
| 0x008C550D | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x008C552C | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x008C5548 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x008C5613 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x008C56EE | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x008C58BC | `MainMenu_Music_Screen_NoAudiobooks` | Known | Screen layout |
| 0x008C58DF | `MainMenus_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x008C5902 | `NoContent_Screen_NoAudiobooks` | Known | Screen layout |
| 0x008C59B4 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x008C59D1 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x008C5A50 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x008C5B34 | `MainMenu_Music_Screen_NoCompilations` | Known | Screen layout |
| 0x008C5B59 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x008C5CA1 | `MainMenu_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x008C5CC4 | `MainMenu_Videos_Screen_NoMusicVideos` | Known | Screen layout |
| 0x008C5CE9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x008C5D08 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x008C5D27 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x008C5D48 | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x008C5D86 | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x008C5DA7 | `Extras_Screen_VoiceMemos` | Known | Screen layout |
| 0x008C5E12 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x008C5E44 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x008C5E63 | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x008C5F10 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x008C5F7C | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x008C6075 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x008C6091 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x008C6114 | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x008C612F | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x008C6150 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x008C61FF | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x008C6233 | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x008C6254 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x008C62F7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x008C6318 | `MainMenu_Videos_Screen_NoPlaylists` | Known | Screen layout |
| 0x008C633B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x008C638A | `MainMenu_Music_Screen_NoArtists` | Known | Screen layout |
| 0x008C6431 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x008C6450 | `Extras_Screen_NikePlus` | Known | Screen layout |
| 0x008C65A0 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x008C65BF | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x008C65E0 | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x008C69FC | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x008C6AAF | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x008C6B29 | `Notes_Contents_Screen_Alt` | Known | Screen layout |
| 0x008C6B43 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x008C6BC3 | `VoiceMemos_Screen_InsufficientDiskSpace2_Default` | Known | Screen layout |
| 0x008C6C75 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x008C6D1A | `VoiceMemos_Screen_InsufficientDiskSpace_Default` | Known | Screen layout |
| 0x008C6D4A | `NowPlaying_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x008C6D77 | `Slideshow_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x008C790A | `Radio_Screen_Tuning_Default` | Known | Screen layout |
| 0x008C796B | `VoiceMemos_Screen_DeletAllAsk_Default` | Known | Screen layout |
| 0x008C7991 | `NowPlaying_Screen_TVOutAsk_Default` | Known | Screen layout |
| 0x008C79B4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x008C79D2 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x008C79FE | `AddressViewer_PartialLoad_Screen_Default` | Known | Screen layout |
| 0x008C7A27 | `Alarms_Alarm_Clock_Triggered_Screen_Default` | Known | Screen layout |
| 0x008C7A53 | `Alarms_Alarm_Triggered_Screen_Default` | Known | Screen layout |
| 0x008C7A79 | `Unsupported_Screen_Default` | Known | Screen layout |
| 0x008C7A94 | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x008C7ABA | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x008C7AD2 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x008C7AED | `Game_Screen_Default` | Known | Screen layout |
| 0x008C7B01 | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x008C7B27 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x008C7B48 | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x008C7B71 | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x008C7B9B | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x008C7BC8 | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x008C7BF1 | `Notes_Loading_Screen_Default` | Known | Screen layout |
| 0x008C7C0E | `Clock_Screen_Default` | Known | Screen layout |
| 0x008C7C23 | `TVOutSlideshowAsk_Screen_Default` | Known | Screen layout |
| 0x008C7C44 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x008C7C62 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x008C7C88 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x008C7CAC | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x008C7CC5 | `AddressViewer_Main_Screen_Default` | Known | Screen layout |
| 0x008C7CE7 | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x008C7D04 | `MainMenus_Main_Screen_Default` | Known | Screen layout |
| 0x008C7D22 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x008C7D3F | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x008C7D5B | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x008C7D87 | `Radio_Screen_Default` | Known | Screen layout |
| 0x008C7D9C | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x008C7DBE | `Extras_Screen_Default` | Known | Screen layout |
| 0x008C7DD4 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x008C7DFA | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x008C7E1B | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x008C7E39 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x008C7E65 | `ToDo_Instructions_Screen_Default` | Known | Screen layout |
| 0x008C7E86 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x008C7EAA | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x008C7ECC | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x008C7EF0 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x008C7F0F | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x008C7F28 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x008C7F4A | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x008C7F6E | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x008C7F8C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x008C7FB0 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x008C7FDA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x008C8003 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x008C8025 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x008C8043 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x008C805C | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x008C8076 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x008C809F | `TODOS_Menu_Screen_Default` | Known | Screen layout |
| 0x008C80B9 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x008C80D7 | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x008C80F4 | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x008C810E | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x008C8129 | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x008C8148 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x008C8166 | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x008C817F | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x008C819B | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x008C81C5 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x008C81E5 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x008C820D | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x008C8234 | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x008C825B | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x008C827C | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x008C82A0 | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x008C82BF | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x008C82E1 | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x008C8304 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x008C8325 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x008C83B3 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x008C83D5 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x008C8414 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x008C89E0 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x008C8A0C | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x008C8A51 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x008C8A79 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x008C8A9A | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x008C8ABB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x008C8AE1 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x008C8AFE | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x008C8B20 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x008C8B44 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x008C8B68 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x008C8CE2 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x008C8D52 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x008C8DA3 | `NowPlaying_Screen_Video_TVOut` | Known | Screen layout |
| 0x008C8EE9 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x008C9415 | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x008C94F2 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x008C96E4 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x008C99B0 | `Lock_Screen_Default_Layout` | Known | Screen layout |
| 0x008C9A1A | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x008C9C29 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x008C9CD3 | `SettingsMenu_About_Screen_Accessory_Layout` | Known | Screen layout |
| 0x008C9D26 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x008CC13C | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x008CC188 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x008CC266 | `MainMenu_List_ScreenLock_x` | Known | Screen layout |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00009003 | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | Known | RTOS |
| 0x0027895C | `  K - RTXC` | Known | RTOS |
| 0x00279944 | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | Known | RTOS |
| 0x008B3C58 | `Returning from RTXCBug` | Known | RTOS |

---

## 9. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000CF850 | `HostOSTask` | Known | RTOS task thread |
| 0x0012533C | `MP3ExampleTask` | Known | RTOS task thread |
| 0x0012A468 | `USBDeviceTask` | Known | RTOS task thread |
| 0x00134564 | `DiskReaderTask` | Known | RTOS task thread |
| 0x00143E44 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x00143E58 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0018FC98 | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x001C7160 | `MeCCAIOTask` | Known | RTOS task thread |
| 0x001F5224 | `StreamCacheMassStorageManagerTimeOutTask` | Known | RTOS task thread |
| 0x001F53A0 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x0026C414 | `FirewireTask` | Known | RTOS task thread |
| 0x0026C428 | `TouchwheelTask` | Known | RTOS task thread |
| 0x0026C43C | `AudioOutStateTask` | Known | RTOS task thread |
| 0x0026C468 | `DiskMgrTask` | Known | RTOS task thread |
| 0x0026C478 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x0026C48C | `TopPlugTask` | Known | RTOS task thread |
| 0x0026C49C | `HPhoneDetTask` | Known | RTOS task thread |
| 0x0026C514 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x0026C53C | `AlarmTask` | Known | RTOS task thread |
| 0x0026C55B | `"USBAudioTask` | Known | RTOS task thread |
| 0x00278FFC | `Undefined Task` | Known | RTOS task thread |
| 0x00368D68 | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x0036C610 | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x00374A0C | `MeCCARecordingTask` | Known | RTOS task thread |
| 0x00810108 | `WaveFileDebugTask` | Known | RTOS task thread |

---

## 10. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0022AB08 | `Channel Reserved` | Known | Logging channel |
| 0x0022AB1C | `Channel AppBoot` | Known | Logging channel |
| 0x0022AB2C | `Channel BufferedSongReading` | Known | Logging channel |
| 0x0022AB48 | `Channel PrefsWriting` | Known | Logging channel |
| 0x0022AB60 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x0022AB80 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x0022AB98 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x0022ABB4 | `Channel TestLogging` | Known | Logging channel |
| 0x0022ABC8 | `Channel AppFileLoading` | Known | Logging channel |
| 0x0022ABE0 | `Channel VCardReading` | Known | Logging channel |
| 0x0022ABF8 | `Channel LongSongScanning` | Known | Logging channel |
| 0x0022AC6C | `Channel VoiceRecording` | Known | Logging channel |
| 0x0022AC84 | `Channel PhotoImporting` | Known | Logging channel |
| 0x0022AC9C | `Channel Notes` | Known | Logging channel |
| 0x0022ACAC | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x0022ACC8 | `Channel DiskMode` | Known | Logging channel |
| 0x0022ACDC | `Channel Firewire` | Known | Logging channel |
| 0x0022ACF0 | `Channel USB` | Known | Logging channel |
| 0x0022AD10 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x0022AD28 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00080848 | `gamedata_RW` | Known | Game system |
| 0x00080864 | `gamedata_ShareRW` | Known | Game system |
| 0x00080878 | `games_RO` | Known | Game system |
| 0x008B3CB2 | `iPod_Control/games_RO/` | Known | Game system |
| 0x008B3CC9 | `Resources/Games/games_RO/` | Known | Game system |
| 0x008BE655 | `GamesMenu_StatusBar_String` | Known | Game system |
| 0x008BECF0 | `AboutScreen_Games_String` | Known | Game system |
| 0x008C541B | `MainMenu_List_Games` | Known | Game system |
| 0x008C542F | `ExtrasMenu_Games` | Known | Game system |
| 0x008CC2D5 | `MainMenu_List_Games_x` | Known | Game system |

---

## 12. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00090B64 | `adrmmp4a` | Known | DRM system |
| 0x00131D0C | `AppleDRMVersion` | Known | DRM system |
| 0x00131DAC | `AppleDRM` | Known | DRM system |
| 0x00132EB0 | `AppleVideoDRM` | Known | DRM system |
| 0x00136288 | `drmsp608aavdmp4aesds` | Known | DRM system |
| 0x008B4091 | `DRMLevel` | Known | DRM system |

---

## 13. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0002FE90 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0002FEA8 | `iPod_Control\iTunes\firsttime` | Known | iTunes database |
| 0x00051790 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x000517B8 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00058144 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0007CC24 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x000807DC | `iPod_Control/iTunes/iTunesDB.p7b` | Known | iTunes database |
| 0x0009C984 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x0009CB58 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000A4F70 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000A63B4 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000A64B4 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0011DC38 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x003620F4 | `iTunesDB` | Known | iTunes database |
| 0x00362100 | `iPod_Control\iTunes\` | Known | iTunes database |

---

## 14. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0005EDCC | `cI: could not read CE-ATA task file` | Known | Hardware |
| 0x0005EDF4 | `cI: CE-ATA signature missing (%x,%x)` | Known | Hardware |
| 0x0005EE4C | `cI: CE-ATA interrupt enable failed` | Known | Hardware |
| 0x0011D49C | `NAND FLASH DRIVE` | Known | Hardware |
| 0x0013224C | `FireWireGUID` | Known | FireWire |
| 0x0013225C | `FireWireVersion` | Known | FireWire |
| 0x00132890 | `FireWire` | Known | FireWire |
| 0x0031285C | `CE-ATA init failed` | Known | Hardware |
| 0x00312C74 | `ISDIE: CE-ATA interrupt enable failed` | Known | Hardware |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x006936B8 | `Settings_RadioRegionsMenu_Layout"` | Known | FM Radio |
| 0x00693741 | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x00725DE8 | `Radio Regions` | Known | FM Radio |
| 0x0076E730 | `Radio-Regionen` | Known | FM Radio |
| 0x008BBC35 | `Settings_RadioRegions_Australia_String` | Known | FM Radio |
| 0x008BBC5C | `Settings_RadioRegions_Asia_String` | Known | FM Radio |
| 0x008BCD9D | `Settings_RadioRegions_Europe_String` | Known | FM Radio |
| 0x008BE09F | `Settings_RadioRegions_Japan_String` | Known | FM Radio |
| 0x008BEB0D | `Settings_RadioRegions_Americas_String` | Known | FM Radio |
| 0x008BF14F | `SettingsMenu_RadioRegions_String` | Known | FM Radio |
| 0x008C2433 | `RadioRegionsMenuItem` | Known | FM Radio |
| 0x008C5ABD | `SettingsMenu_ListItem_RadioRegions` | Known | FM Radio |
| 0x008C95BE | `SettingsInfo_Template_RadioRegions_Layout` | Known | FM Radio |
| 0x008C95E8 | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x008C9BEA | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 16. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x007A7B48 | `Fotocamera` | Known | Camera |
| 0x007A80BC | `Collega fotocamera o inserisci scheda` | Known | Camera |
| 0x007A8134 | `Fotocamera non supportata` | Known | Camera |
| 0x007C4400 | `Camera` | Known | Camera |
| 0x007C4990 | `Sluit camera of kaart aan` | Known | Camera |
| 0x007C49FC | `Camera niet ondersteund` | Known | Camera |
| 0x008BBC7E | `Photos_DevName_Camera_String` | Known | Camera |

---

## 17. Pedometer

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x008CC573 | `NikePlus_Remote_Step_Away` | Known | Pedometer |
| 0x008CC58D | `NikePlus_Step_Away` | Known | Pedometer |

---

## 18. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0002FE7C | `iPod_Control` | Filesystem Path |  |
| 0x0002FEE8 | `iPod_Control\Device` | Filesystem Path |  |
| 0x0003E3F0 | `iPod_Control\Device` | Filesystem Path |  |
| 0x000404E8 | `iPod_Control` | Filesystem Path |  |
| 0x00040B70 | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x00051770 | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path |  |
| 0x00054708 | `iPod_Control\Music\` | Filesystem Path |  |
| 0x00057FC4 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x0008A550 | `iPod_Control` | Filesystem Path |  |
| 0x0008A560 | `Resources/Games` | Filesystem Path |  |
| 0x0008A570 | `iPod_Control/%s%s%s` | Filesystem Path |  |
| 0x000EFE08 | `iPod_Control/Device/alarms` | Filesystem Path |  |
| 0x000FFBD0 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00100F30 | `iPod_Control/Device` | Filesystem Path |  |
| 0x00100F44 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x001189C4 | `iPod_Control/Device/Users` | Filesystem Path |  |
| 0x00145044 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x001452A0 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x00150DB8 | `Resources/UI/active.bin` | Filesystem Path |  |
| 0x00150DD0 | `Resources/UI/` | Filesystem Path |  |
| 0x001712BC | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x00171F7C | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path |  |
| 0x00171FA4 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001932E0 | `iPod_Control/Device/PlayCounts` | Filesystem Path |  |
| 0x001A8008 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A80B8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A8234 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A83CC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A8474 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A861C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A86C0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A8764 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A8808 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A88AC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A8950 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A89F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A8AA4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A8B54 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A8C04 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A8D70 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A8E20 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A8ED0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A8F74 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9024 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9118 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A91BC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9270 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A932C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A93DC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9500 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A95BC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A966C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9828 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A98EC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A999C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9A58 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9B94 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9C60 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9D1C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9DC0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9E64 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9F20 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001A9FDC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA0A4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA148 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA210 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA2C0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA388 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA450 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA500 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA5B0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA674 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA724 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA7D4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA884 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AA958 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AAA2C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AAB2C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AAC0C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AAD14 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001AAE00 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00237D08 | `Resources/Fonts` | Filesystem Path |  |
| 0x00250BD4 | `Resources/Fonts` | Filesystem Path |  |
| 0x00362172 | `iPod_Control/Device` | Filesystem Path |  |
| 0x00368608 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x0036B15C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0036B412 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003749D8 | `iPod_Control:Device:voices:` | Filesystem Path |  |
| 0x008B3B8D | `Resources/Games/` | Filesystem Path |  |
| 0x008B3F73 | `iPod_Control/Device` | Filesystem Path |  |
| 0x008B3F87 | `iPod_Control/Device/accessories.reg` | Filesystem Path |  |
| 0x008B4008 | `iPod_Control/Tones` | Filesystem Path |  |

---

## 19. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00813110 | `c:\bwa\N25FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftu` | Build Path |  |
| 0x00813168 | `c:\bwa\N25FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fts` | Build Path |  |
| 0x008131C0 | `c:\bwa\N25FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fto` | Build Path |  |
| 0x0081D790 | `c:\bwa\N25FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\` | Build Path |  |
| 0x0081E30C | `c:\bwa\N25FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfd` | Build Path |  |
| 0x0081F508 | `c:\bwa\N25FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrg` | Build Path |  |
| 0x0081F560 | `c:\bwa\N25FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrc` | Build Path |  |
| 0x0081F5B8 | `c:\bwa\N25FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfro` | Build Path |  |
| 0x0081F8FC | `c:\bwa\N25FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1` | Build Path |  |
| 0x0082ECA4 | `c:\bwa\N25FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttc` | Build Path |  |
| 0x0082EF20 | `c:\bwa\N25FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype` | Build Path |  |
| 0x0082F48C | `c:\bwa\N25FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1` | Build Path |  |

---

## 20. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00088258 | `Acoustic` | EQ Preset |  |
| 0x00088264 | `Bass Booster` | EQ Preset |  |
| 0x00088284 | `Classical` | EQ Preset |  |
| 0x00088290 | `Dance` | EQ Preset |  |
| 0x000882A0 | `Electronic` | EQ Preset |  |
| 0x000882B4 | `Hip Hop` | EQ Preset |  |
| 0x000882BC | `Jazz` | EQ Preset |  |
| 0x000882C4 | `Latin` | EQ Preset |  |
| 0x000882CC | `Loudness` | EQ Preset |  |
| 0x000882D8 | `Lounge` | EQ Preset |  |
| 0x000882E0 | `Piano` | EQ Preset |  |
| 0x000882F4 | `Rock` | EQ Preset |  |
| 0x000882FC | `Small Speakers` | EQ Preset |  |
| 0x0008830C | `Spoken Word` | EQ Preset |  |
| 0x00088318 | `Treble Booster` | EQ Preset |  |
| 0x00088364 | `Vocal Booster` | EQ Preset |  |
| 0x007260D8 | `Acoustic` | EQ Preset |  |
| 0x007260E4 | `Bass Booster` | EQ Preset |  |
| 0x00726104 | `Classical` | EQ Preset |  |
| 0x00726110 | `Dance` | EQ Preset |  |
| 0x00726120 | `Electronic` | EQ Preset |  |
| 0x00726134 | `Hip Hop` | EQ Preset |  |
| 0x0072613C | `Jazz` | EQ Preset |  |
| 0x00726144 | `Latin` | EQ Preset |  |
| 0x0072614C | `Loudness` | EQ Preset |  |
| 0x00726158 | `Lounge` | EQ Preset |  |
| 0x00726160 | `Piano` | EQ Preset |  |
| 0x00726170 | `Rock` | EQ Preset |  |
| 0x00726178 | `Small Speakers` | EQ Preset |  |
| 0x00726188 | `Spoken Word` | EQ Preset |  |
| 0x00726194 | `Treble Booster` | EQ Preset |  |
| 0x007261B4 | `Vocal Booster` | EQ Preset |  |
| 0x0075D7F8 | `Acoustic` | EQ Preset |  |
| 0x0075D804 | `Bass Booster` | EQ Preset |  |
| 0x0075D824 | `Classical` | EQ Preset |  |
| 0x0075D830 | `Dance` | EQ Preset |  |
| 0x0075D840 | `Electronic` | EQ Preset |  |
| 0x0075D854 | `Hip Hop` | EQ Preset |  |
| 0x0075D85C | `Jazz` | EQ Preset |  |
| 0x0075D864 | `Latin` | EQ Preset |  |
| 0x0075D86C | `Loudness` | EQ Preset |  |
| 0x0075D878 | `Lounge` | EQ Preset |  |
| 0x0075D880 | `Piano` | EQ Preset |  |
| 0x0075D890 | `Rock` | EQ Preset |  |
| 0x0075D898 | `Small Speakers` | EQ Preset |  |
| 0x0075D8A8 | `Spoken Word` | EQ Preset |  |
| 0x0075D8B4 | `Treble Booster` | EQ Preset |  |
| 0x0075D8D4 | `Vocal Booster` | EQ Preset |  |
| 0x00765FC8 | `Acoustic` | EQ Preset |  |
| 0x00765FD4 | `Bass Booster` | EQ Preset |  |
| 0x00765FF4 | `Classical` | EQ Preset |  |
| 0x00766000 | `Dance` | EQ Preset |  |
| 0x00766010 | `Electronic` | EQ Preset |  |
| 0x00766024 | `Hip Hop` | EQ Preset |  |
| 0x0076602C | `Jazz` | EQ Preset |  |
| 0x00766034 | `Latin` | EQ Preset |  |
| 0x0076603C | `Loudness` | EQ Preset |  |
| 0x00766048 | `Lounge` | EQ Preset |  |
| 0x00766050 | `Piano` | EQ Preset |  |
| 0x00766060 | `Rock` | EQ Preset |  |
| 0x00766068 | `Small Speakers` | EQ Preset |  |
| 0x00766078 | `Spoken Word` | EQ Preset |  |
| 0x00766084 | `Treble Booster` | EQ Preset |  |
| 0x007660A4 | `Vocal Booster` | EQ Preset |  |
| 0x0076EAD8 | `Acoustic` | EQ Preset |  |
| 0x0076EB08 | `Dance` | EQ Preset |  |
| 0x0076EB18 | `Electronic` | EQ Preset |  |
| 0x0076EB34 | `Jazz` | EQ Preset |  |
| 0x0076EB3C | `Latin` | EQ Preset |  |
| 0x0076EB44 | `Loudness` | EQ Preset |  |
| 0x0076EB58 | `Piano` | EQ Preset |  |
| 0x0076EB68 | `Rock` | EQ Preset |  |
| 0x00784844 | `Dance` | EQ Preset |  |
| 0x0078486C | `Hip Hop` | EQ Preset |  |
| 0x00784874 | `Jazz` | EQ Preset |  |
| 0x00784884 | `Loudness` | EQ Preset |  |
| 0x00784890 | `Lounge` | EQ Preset |  |
| 0x00784898 | `Piano` | EQ Preset |  |
| 0x007848A8 | `Rock` | EQ Preset |  |
| 0x0078CF98 | `Jazz` | EQ Preset |  |
| 0x0078CFA0 | `Latin` | EQ Preset |  |
| 0x0078CFB4 | `Lounge` | EQ Preset |  |
| 0x0078CFBC | `Piano` | EQ Preset |  |
| 0x0078CFCC | `Rock` | EQ Preset |  |
| 0x00795638 | `Hip Hop` | EQ Preset |  |
| 0x00795640 | `Jazz` | EQ Preset |  |
| 0x0079565C | `Lounge` | EQ Preset |  |
| 0x00795664 | `Piano` | EQ Preset |  |
| 0x0079567C | `Rock` | EQ Preset |  |
| 0x0079E874 | `Latin` | EQ Preset |  |
| 0x0079E8A0 | `Rock` | EQ Preset |  |
| 0x007A7444 | `Dance` | EQ Preset |  |
| 0x007A7468 | `Hip Hop` | EQ Preset |  |
| 0x007A7470 | `Jazz` | EQ Preset |  |
| 0x007A7480 | `Loudness` | EQ Preset |  |
| 0x007A748C | `Lounge` | EQ Preset |  |
| 0x007A7494 | `Piano` | EQ Preset |  |
| 0x007A74A4 | `Rock` | EQ Preset |  |
| 0x007B1158 | `Acoustic` | EQ Preset |  |
| 0x007B1164 | `Bass Booster` | EQ Preset |  |
| 0x007B1184 | `Classical` | EQ Preset |  |
| 0x007B1190 | `Dance` | EQ Preset |  |
| 0x007B11A0 | `Electronic` | EQ Preset |  |
| 0x007B11B4 | `Hip Hop` | EQ Preset |  |
| 0x007B11BC | `Jazz` | EQ Preset |  |
| 0x007B11C4 | `Latin` | EQ Preset |  |
| 0x007B11CC | `Loudness` | EQ Preset |  |
| 0x007B11D8 | `Lounge` | EQ Preset |  |
| 0x007B11E0 | `Piano` | EQ Preset |  |
| 0x007B11F0 | `Rock` | EQ Preset |  |
| 0x007B11F8 | `Small Speakers` | EQ Preset |  |
| 0x007B1208 | `Spoken Word` | EQ Preset |  |
| 0x007B1214 | `Treble Booster` | EQ Preset |  |
| 0x007B1234 | `Vocal Booster` | EQ Preset |  |
| 0x007BAE78 | `Acoustic` | EQ Preset |  |
| 0x007BAE84 | `Bass Booster` | EQ Preset |  |
| 0x007BAEA4 | `Classical` | EQ Preset |  |
| 0x007BAEB0 | `Dance` | EQ Preset |  |
| 0x007BAEC0 | `Electronic` | EQ Preset |  |
| 0x007BAED4 | `Hip Hop` | EQ Preset |  |
| 0x007BAEDC | `Jazz` | EQ Preset |  |
| 0x007BAEE4 | `Latin` | EQ Preset |  |
| 0x007BAEEC | `Loudness` | EQ Preset |  |
| 0x007BAEF8 | `Lounge` | EQ Preset |  |
| 0x007BAF00 | `Piano` | EQ Preset |  |
| 0x007BAF10 | `Rock` | EQ Preset |  |
| 0x007BAF18 | `Small Speakers` | EQ Preset |  |
| 0x007BAF28 | `Spoken Word` | EQ Preset |  |
| 0x007BAF34 | `Treble Booster` | EQ Preset |  |
| 0x007BAF54 | `Vocal Booster` | EQ Preset |  |
| 0x007C3CF4 | `Dance` | EQ Preset |  |
| 0x007C3D28 | `Jazz` | EQ Preset |  |
| 0x007C3D30 | `Latin` | EQ Preset |  |
| 0x007C3D38 | `Loudness` | EQ Preset |  |
| 0x007C3D44 | `Lounge` | EQ Preset |  |
| 0x007C3D4C | `Piano` | EQ Preset |  |
| 0x007C3D5C | `Rock` | EQ Preset |  |
| 0x007CC4C8 | `Dance` | EQ Preset |  |
| 0x007CC4F4 | `Jazz` | EQ Preset |  |
| 0x007CC504 | `Loudness` | EQ Preset |  |
| 0x007CC510 | `Lounge` | EQ Preset |  |
| 0x007CC518 | `Piano` | EQ Preset |  |
| 0x007CC528 | `Rock` | EQ Preset |  |
| 0x007D4DCC | `Hip Hop` | EQ Preset |  |
| 0x007D4DD4 | `Jazz` | EQ Preset |  |
| 0x007D4DF8 | `Lounge` | EQ Preset |  |
| 0x007D4E00 | `Piano` | EQ Preset |  |
| 0x007D4E10 | `Rock` | EQ Preset |  |
| 0x007DDAE0 | `Hip Hop` | EQ Preset |  |
| 0x007DDAE8 | `Jazz` | EQ Preset |  |
| 0x007DDB04 | `Lounge` | EQ Preset |  |
| 0x007DDB0C | `Piano` | EQ Preset |  |
| 0x007DDB1C | `Rock` | EQ Preset |  |
| 0x007F22B0 | `Acoustic` | EQ Preset |  |
| 0x007F22BC | `Bass Booster` | EQ Preset |  |
| 0x007F22DC | `Classical` | EQ Preset |  |
| 0x007F22E8 | `Dance` | EQ Preset |  |
| 0x007F22F8 | `Electronic` | EQ Preset |  |
| 0x007F230C | `Hip Hop` | EQ Preset |  |
| 0x007F2314 | `Jazz` | EQ Preset |  |
| 0x007F231C | `Latin` | EQ Preset |  |
| 0x007F2324 | `Loudness` | EQ Preset |  |
| 0x007F2330 | `Lounge` | EQ Preset |  |
| 0x007F2338 | `Piano` | EQ Preset |  |
| 0x007F2348 | `Rock` | EQ Preset |  |
| 0x007F2350 | `Small Speakers` | EQ Preset |  |
| 0x007F2360 | `Spoken Word` | EQ Preset |  |
| 0x007F236C | `Treble Booster` | EQ Preset |  |
| 0x007F238C | `Vocal Booster` | EQ Preset |  |
| 0x007FAC1C | `Hip Hop` | EQ Preset |  |
| 0x007FAC28 | `Latin` | EQ Preset |  |
| 0x007FAC30 | `Loudness` | EQ Preset |  |
| 0x007FAC3C | `Lounge` | EQ Preset |  |
| 0x007FAC54 | `Rock` | EQ Preset |  |
| 0x00803754 | `Acoustic` | EQ Preset |  |
| 0x00803760 | `Bass Booster` | EQ Preset |  |
| 0x00803780 | `Classical` | EQ Preset |  |
| 0x0080378C | `Dance` | EQ Preset |  |
| 0x0080379C | `Electronic` | EQ Preset |  |
| 0x008037B0 | `Hip Hop` | EQ Preset |  |
| 0x008037B8 | `Jazz` | EQ Preset |  |
| 0x008037C0 | `Latin` | EQ Preset |  |
| 0x008037C8 | `Loudness` | EQ Preset |  |
| 0x008037D4 | `Lounge` | EQ Preset |  |
| 0x008037DC | `Piano` | EQ Preset |  |
| 0x008037EC | `Rock` | EQ Preset |  |
| 0x008037F4 | `Small Speakers` | EQ Preset |  |
| 0x00803804 | `Spoken Word` | EQ Preset |  |
| 0x00803810 | `Treble Booster` | EQ Preset |  |
| 0x00803830 | `Vocal Booster` | EQ Preset |  |
| 0x0080C0FC | `Acoustic` | EQ Preset |  |
| 0x0080C108 | `Bass Booster` | EQ Preset |  |
| 0x0080C128 | `Classical` | EQ Preset |  |
| 0x0080C134 | `Dance` | EQ Preset |  |
| 0x0080C144 | `Electronic` | EQ Preset |  |
| 0x0080C158 | `Hip Hop` | EQ Preset |  |
| 0x0080C160 | `Jazz` | EQ Preset |  |
| 0x0080C168 | `Latin` | EQ Preset |  |
| 0x0080C170 | `Loudness` | EQ Preset |  |
| 0x0080C17C | `Lounge` | EQ Preset |  |
| 0x0080C184 | `Piano` | EQ Preset |  |
| 0x0080C194 | `Rock` | EQ Preset |  |
| 0x0080C19C | `Small Speakers` | EQ Preset |  |
| 0x0080C1AC | `Spoken Word` | EQ Preset |  |
| 0x0080C1B8 | `Treble Booster` | EQ Preset |  |
| 0x0080C1D8 | `Vocal Booster` | EQ Preset |  |

---
