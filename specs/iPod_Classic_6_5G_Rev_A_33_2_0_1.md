# iPod Classic 6.5G (Rev A, 120GB) - RetailOS 2.0.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 2.0.1 |
| **IPSW** | iPod_33.2.0.1.ipsw |
| **Device** | iPod Classic 6.5G (Rev A, 120GB) (2008, 120GB, Click Wheel, Cover Flow, Genius, CE-ATA HDD) |
| **UpdaterFamilyID** | 33 |
| **Binary Size** | 10,514,000 bytes (10.03 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,511,952 bytes |
| **Total Strings (>=4)** | 71,744 |
| **Function Prologues** | 22,810 (ARM: 17,413, Thumb: 5,397) |
| **DRAM References** | 106,510 |
| **Peripheral Refs** | 7,195 |
| **Build** | N25BFirmwareWin-93 |
| **SoC** | S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Codename** | N25C |
| **DFU PID** | 0x1223 |
| **SHA-256** | `600abfe55ba5504a0330639597a88cd405e465fa3014ad1487e75e95d340ce7a` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000954CC | `TSilverCntlr` | Known | Controller |
| 0x000954E4 | `TCExtrasMenu` | Known | Controller |
| 0x000954FC | `TCGameScreen` | Known | Controller |
| 0x00095514 | `TCGamesMenu` | Known | Controller |
| 0x00095528 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x00095550 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00095578 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x000955A4 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x000955C8 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x000955F0 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x00095618 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x00095640 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x00095668 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x00095690 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x000956C0 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x000956EC | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x0009571C | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x00095744 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x0009576C | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x00095798 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x000957C0 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x000957E8 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x00095818 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x00095848 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x000959A4 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x000959D4 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x000959FC | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x00095A24 | `TCRentalNotification` | Known | Controller |
| 0x00095A44 | `TCRentalInfo` | Known | Controller |
| 0x00095A5C | `TCRentalConfirmDelete` | Known | Controller |
| 0x00095A7C | `TCRentalDispatcher` | Known | Controller |
| 0x00095AD4 | `TSilverGlobalCntlr` | Known | Controller |
| 0x00095AF0 | `TSilverTrainerCntlr` | Known | Controller |
| 0x000ECEA0 | `TCSlideshowLCD` | Known | Controller |
| 0x000ECEB8 | `TCSlideshowTVOut` | Known | Controller |
| 0x000ECED4 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x000ECEF4 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x001109EC | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x00110A18 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x00110A44 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x00110A6C | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x00110A98 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x00110AC0 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x00110AEC | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x0011867C | `TCRemoteUI` | Known | Controller |
| 0x00118690 | `TCUnsupported` | Known | Controller |
| 0x0011EB1C | `TCSpeakers` | Known | Controller |
| 0x0011EB30 | `TCEQSetting` | Known | Controller |
| 0x00147994 | `TCSportTimer` | Known | Controller |
| 0x001479AC | `TCSportTimerMenu` | Known | Controller |
| 0x001479C8 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x001479EC | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x00148D9C | `TCVoiceMemos` | Known | Controller |
| 0x00148DB4 | `TCVoiceMemosMenu` | Known | Controller |
| 0x00148DD0 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x00148DF0 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x00148E10 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x00148E30 | `TCVoiceMemosAlert` | Known | Controller |
| 0x0015AAA8 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x0015AAD0 | `TCSettings_MainMenu` | Known | Controller |
| 0x0015AAEC | `TCSettings_MusicMenu` | Known | Controller |
| 0x0015AB0C | `TCSettings_VolumeLimit` | Known | Controller |
| 0x0015AB2C | `TCSettings_Brightness` | Known | Controller |
| 0x0015AB4C | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0015AB70 | `TCSettings_EQ` | Known | Controller |
| 0x0015AB88 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x0015ABB0 | `TCSettings_RadioRegions` | Known | Controller |
| 0x0015ABD0 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0015ABF4 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x0015AC18 | `TCDateTimeScreen` | Known | Controller |
| 0x0015AC34 | `TCTimeZoneScreen` | Known | Controller |
| 0x0015AC50 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0015AC78 | `TCFirstBoot` | Known | Controller |
| 0x00171190 | `TCDemoMode` | Known | Controller |
| 0x00199A10 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x00199A30 | `TCAddressViewerDetails` | Known | Controller |
| 0x00199A50 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x00199A74 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x001C62B0 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x001C62D4 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x001CDAFC | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x002630B4 | `TC_LockDialog` | Known | Controller |
| 0x002630CC | `TC_LockScreen` | Known | Controller |
| 0x002630E4 | `TC_LockediPod` | Known | Controller |
| 0x002630FC | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x00263120 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00268CA4 | `TCClock` | Known | Controller |
| 0x00268CB4 | `TCClockCityMenu` | Known | Controller |
| 0x00268CCC | `TCClockRegionMenu` | Known | Controller |
| 0x00268CE8 | `TCAlarmMenu` | Known | Controller |
| 0x00268CFC | `TCSleepTimerMenu` | Known | Controller |
| 0x00268D18 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x00268D38 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x00268D60 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x00268D84 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00268DA8 | `TCAlarmDatePicker` | Known | Controller |
| 0x00268DC4 | `TCAlarmTriggered` | Known | Controller |
| 0x0026FCE8 | `TCNotesDispatcher` | Known | Controller |
| 0x0026FD04 | `TCNotesLoading` | Known | Controller |
| 0x0026FD1C | `TCNotesList` | Known | Controller |
| 0x0026FD30 | `TCNotesContents` | Known | Controller |
| 0x003DE1C8 | `TCAlarmTriggered` | Known | Controller |
| 0x003DE1DC | `TSilverCntlr` | Known | Controller |
| 0x003DE1FC | `TCClock` | Known | Controller |
| 0x003DE204 | `TCClockRegionMenu` | Known | Controller |
| 0x003DE218 | `TCClockCityMenu` | Known | Controller |
| 0x003DE228 | `TCAlarmMenu` | Known | Controller |
| 0x003DE234 | `TCSleepTimerMenu` | Known | Controller |
| 0x003DE248 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003DE260 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003DE280 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003DE29C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003DE2B8 | `TCAlarmDatePicker` | Known | Controller |
| 0x003DE2F0 | `TSilverCntlr` | Known | Controller |
| 0x003DE310 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003DE4A0 | `TSilverCntlr` | Known | Controller |
| 0x003DE4C0 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x003DE4E0 | `TCSettings_Brightness` | Known | Controller |
| 0x003DE4F8 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x003DE514 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x003DE534 | `TCSettings_RadioRegions` | Known | Controller |
| 0x003DE54C | `TCSettings_EQ` | Known | Controller |
| 0x003DE55C | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x003DE578 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x003DE598 | `TCFirstBoot` | Known | Controller |
| 0x003DE5A4 | `TCSettings_MainMenu` | Known | Controller |
| 0x003DE5B8 | `TCSettings_MusicMenu` | Known | Controller |
| 0x003DE5D0 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003DE5E8 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x003DE604 | `TCDateTimeScreen` | Known | Controller |
| 0x003DE618 | `TCTimeZoneScreen` | Known | Controller |
| 0x003E570C | `TSilverCntlr` | Known | Controller |
| 0x003E572C | `TCClock` | Known | Controller |
| 0x003E5734 | `TCClockRegionMenu` | Known | Controller |
| 0x003E5748 | `TCClockCityMenu` | Known | Controller |
| 0x003E5758 | `TCAlarmMenu` | Known | Controller |
| 0x003E5764 | `TCSleepTimerMenu` | Known | Controller |
| 0x003E5778 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003E57F0 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003E5810 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003E582C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003E5860 | `TCAlarmDatePicker` | Known | Controller |
| 0x003E5874 | `TCAlarmTriggered` | Known | Controller |
| 0x003E6958 | `TSilverCntlr` | Known | Controller |
| 0x003E6978 | `TC_LockDialog` | Known | Controller |
| 0x003E6988 | `TC_LockScreen` | Known | Controller |
| 0x003E6998 | `TC_LockediPod` | Known | Controller |
| 0x003E69A8 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003E69C4 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003E69DC | `TSilverCntlr` | Known | Controller |
| 0x003E6BE8 | `TSilverCntlr` | Known | Controller |
| 0x003E6C04 | `TSilverCntlr` | Known | Controller |
| 0x003E6C68 | `TSilverCntlr` | Known | Controller |
| 0x003E6C88 | `TCNotesDispatcher` | Known | Controller |
| 0x003E6C9C | `TCNotesLoading` | Known | Controller |
| 0x003E6CAC | `TCNotesBase` | Known | Controller |
| 0x003E6CB8 | `TCNotesList` | Known | Controller |
| 0x003E6CC4 | `TCNotesContents` | Known | Controller |
| 0x003E6CD4 | `TSilverCntlr` | Known | Controller |
| 0x003E6CF4 | `TCRemoteUI` | Known | Controller |
| 0x003E6D00 | `TCUnsupported` | Known | Controller |
| 0x003E6D10 | `TSilverCntlr` | Known | Controller |
| 0x003E6D74 | `TSilverCntlr` | Known | Controller |
| 0x003E6D94 | `TCSportTimer` | Known | Controller |
| 0x003E6DA4 | `TCSportTimerMenu` | Known | Controller |
| 0x003E6DB8 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x003E6DD4 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x003E6E04 | `TSilverCntlr` | Known | Controller |
| 0x003E6F2C | `TSilverCntlr` | Known | Controller |
| 0x003E6F4C | `TCDemoMode` | Known | Controller |
| 0x003E6F58 | `TCClock` | Known | Controller |
| 0x003E6F60 | `TCClockRegionMenu` | Known | Controller |
| 0x003E6F74 | `TCClockCityMenu` | Known | Controller |
| 0x003E6F84 | `TCAlarmMenu` | Known | Controller |
| 0x003E6F90 | `TCSleepTimerMenu` | Known | Controller |
| 0x003E6FA4 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003E6FBC | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003E6FDC | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003E6FF8 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003E7014 | `TCAlarmDatePicker` | Known | Controller |
| 0x003E7028 | `TCAlarmTriggered` | Known | Controller |
| 0x003E7048 | `TSilverCntlr` | Known | Controller |
| 0x003E7064 | `TSilverCntlr` | Known | Controller |
| 0x003E7074 | `TSilverCntlr` | Known | Controller |
| 0x003E7094 | `TCVoiceMemos` | Known | Controller |
| 0x003E70A4 | `TCVoiceMemosMenu` | Known | Controller |
| 0x003E70B8 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x003E70D0 | `TCVoiceMemosAlert` | Known | Controller |
| 0x003E70E4 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x003E70FC | `TCVoiceMemosPlayback` | Known | Controller |
| 0x003E711C | `TSilverCntlr` | Known | Controller |
| 0x003E717C | `TSilverCntlr` | Known | Controller |
| 0x003E71E8 | `TSilverCntlr` | Known | Controller |
| 0x003E8510 | `TSilverCntlr` | Known | Controller |
| 0x003E861C | `TSilverCntlr` | Known | Controller |
| 0x003F0E94 | `TSilverCntlr` | Known | Controller |
| 0x003F0EB4 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x003F0ECC | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x003F0EE8 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x003F0F08 | `TCAddressViewerDetails` | Known | Controller |
| 0x003F0F20 | `TSilverCntlr` | Known | Controller |
| 0x003F0F40 | `TSilverCalendarCntlr_Base` | Known | Controller |
| 0x003F0F5C | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x003F0F80 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x003F0FA4 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x003F0FC4 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x003F0FE8 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x003F1008 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x003F102C | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x003F1204 | `TSilverCntlr` | Known | Controller |
| 0x003F1224 | `TC_LockDialog` | Known | Controller |
| 0x003F1234 | `TC_LockScreen` | Known | Controller |
| 0x003F1244 | `TC_LockediPod` | Known | Controller |
| 0x003F1254 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003F1278 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003F1398 | `TSilverCntlr` | Known | Controller |
| 0x003F14CC | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003F14E8 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003F1508 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003F1528 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003F1550 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003F1574 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003F159C | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003F15BC | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003F15DC | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003F15FC | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003F161C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003F1644 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003F166C | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x003F168C | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003F16AC | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003F16CC | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003F16F0 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003F1710 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003F1734 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003F175C | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003F1788 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003F17A8 | `TCRentalNotification` | Known | Controller |
| 0x003F17C0 | `TCRentalInfo` | Known | Controller |
| 0x003F17D0 | `TCRentalConfirmDelete` | Known | Controller |
| 0x003F17E8 | `TCRentalDispatcher` | Known | Controller |
| 0x003F20D8 | `TSilverCntlr` | Known | Controller |
| 0x003F219C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003F21B8 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003F21D8 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003F21F8 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003F2220 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003F2244 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003F226C | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003F228C | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003F22AC | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003F22CC | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003F22EC | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003F2314 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003F233C | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x003F2384 | `TCSlideshowTVOut` | Known | Controller |
| 0x003F2398 | `TCSlideshowLCD` | Known | Controller |
| 0x003F23A8 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x003F23C0 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x003F23E0 | `TSilverCntlr` | Known | Controller |
| 0x003F240C | `TSilverCntlr` | Known | Controller |
| 0x003F242C | `TCUnsupported` | Known | Controller |
| 0x003F244C | `TSilverCntlr` | Known | Controller |
| 0x003F248C | `TSilverCntlr` | Known | Controller |
| 0x003F24AC | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x003F24C8 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x003F24E0 | `TSilverCntlr` | Known | Controller |
| 0x003F2500 | `TCSpeakers` | Known | Controller |
| 0x003F250C | `TCEQSetting` | Known | Controller |
| 0x003F252C | `TSilverCntlr` | Known | Controller |
| 0x003F2594 | `TSilverCntlr` | Known | Controller |
| 0x003F25B4 | `TCExtrasMenu` | Known | Controller |
| 0x003F25C4 | `TCGamesMenu` | Known | Controller |
| 0x003F25D0 | `TCGameScreen` | Known | Controller |
| 0x003F25E0 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x003F2600 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x003F2620 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x003F2640 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x003F2664 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003F2680 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003F26A0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003F26C0 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003F26E8 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003F270C | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003F2734 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003F2754 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003F2774 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003F2794 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003F27B4 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003F27DC | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003F2804 | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x003F2824 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003F2844 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003F2864 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003F2888 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003F28A8 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003F28CC | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003F28F4 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003F2920 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003F2940 | `TCRentalNotification` | Known | Controller |
| 0x003F2958 | `TCRentalInfo` | Known | Controller |
| 0x003F2968 | `TCRentalConfirmDelete` | Known | Controller |
| 0x003F2980 | `TCRentalDispatcher` | Known | Controller |
| 0x003F2994 | `TSilverGlobalCntlr` | Known | Controller |
| 0x003F29A8 | `TSilverTrainerCntlr` | Known | Controller |
| 0x00477DD4 | `TCCCCCCCCCCCCCCCCCCCCCCT` | Known | Controller |
| 0x0071A17E | `TCNotesDispatcher"` | Known | Controller |
| 0x0071A23D | `TCLockChosenDispatcher"` | Known | Controller |
| 0x0071A300 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x007243E5 | `TCNotesDispatcher"` | Known | Controller |
| 0x00724547 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x00739858 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x00739870 | `TCAddressViewerDetails` | Known | Controller |
| 0x00739888 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x007398A4 | `TCAlarmMenu` | Known | Controller |
| 0x007398B0 | `TCSleepTimerMenuTCAlarmPropertiesMenu` | Known | Controller |
| 0x007398D8 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x007398F8 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x00739914 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00739930 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0073994C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00739968 | `TCAlarmDatePicker` | Known | Controller |
| 0x0073997C | `TCAlarmDatePicker` | Known | Controller |
| 0x00739990 | `TCAlarmTriggeredTSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x007399BC | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x007399E0 | `TSilverCalendarCntlr_MonthViewerTSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x00739A20 | `TSilverCalendarCntlr_EventViewerTSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x00739A60 | `TSilverCalendarCntlr_EventViewerTSilverCntlrTCClockRegionMenu` | Known | Controller |
| 0x00739AA0 | `TCClockCityMenu` | Known | Controller |
| 0x00739AB0 | `TCClockCityMenu` | Known | Controller |
| 0x00739AC0 | `TCClockCityMenu` | Known | Controller |
| 0x00739AD0 | `TCClockCityMenu` | Known | Controller |
| 0x00739AE0 | `TCClockCityMenu` | Known | Controller |
| 0x00739AF0 | `TCClockCityMenu` | Known | Controller |
| 0x00739B00 | `TCClockCityMenu` | Known | Controller |
| 0x00739B10 | `TCClockCityMenu` | Known | Controller |
| 0x00739B20 | `TCClock` | Known | Controller |
| 0x00739B38 | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDe` | Known | Controller |
| 0x00739B90 | `TCGamesMenu` | Known | Controller |
| 0x00739B9C | `TCGameScreenTC_LockediPod` | Known | Controller |
| 0x00739BB8 | `TC_LockDialog` | Known | Controller |
| 0x00739BC8 | `TC_LockScreen` | Known | Controller |
| 0x00739BD8 | `TC_VolumeLimitLockScreenTSilverCntlrTSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x00739C1C | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00739C3C | `TSilverMediaListCntlr_AudiobooksTSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x00739C84 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x00739CA0 | `TSilverMediaListCntlr_AlbumsTSilverMediaListCntlr_Artists` | Known | Controller |
| 0x00739CDC | `TSilverMediaListCntlr_GenresTSilverMediaListCntlr_Composers` | Known | Controller |
| 0x00739D18 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x00739D38 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x00739D60 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x00739D80 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x00739DA0 | `TSilverMediaListCntlr_TVEpisodesTSilverMediaListCntlr_MoviesTSilverMediaListCntl` | Known | Controller |
| 0x00739DFC | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x00739E24 | `TSilverMediaListCntlr_VideoPlaylistsTSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x00739E68 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x00739E94 | `TSilverCntlrTSilverMainMediaListCntlr_VideosTSilverSettingsVideoCntlr` | Known | Controller |
| 0x00739EDC | `TCFirstBoot` | Known | Controller |
| 0x00739EE8 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00739F08 | `TSilverMediaListCntlr_GeniusTSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x00739FE8 | `TCRentalInfoTCRentalConfirmDelete` | Known | Controller |
| 0x0073A00C | `TSilverCntlrTCRentalNotificationTCRentalNotificationTCRentalNotificationTCNotesL` | Known | Controller |
| 0x0073A064 | `TCNotesList` | Known | Controller |
| 0x0073A070 | `TCNotesList` | Known | Controller |
| 0x0073A07C | `TCNotesContents` | Known | Controller |
| 0x0073A08C | `TCNotesContents` | Known | Controller |
| 0x0073A09C | `TCNotesContents` | Known | Controller |
| 0x0073A0AC | `TCNotesContents` | Known | Controller |
| 0x0073A168 | `TCSlideshowLCD` | Known | Controller |
| 0x0073A178 | `TCSlideshowTVOutTCSlideshow_TVOutAskTCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x0073A1C8 | `TCRemoteUI` | Known | Controller |
| 0x0073A1D4 | `TCUnsupported` | Known | Controller |
| 0x0073A1E4 | `TSilverSettingsMenuListCntlrTSilverSettingsMenuListCntlrTSilverSettingsMenuListC` | Known | Controller |
| 0x0073A24C | `TCSettings_MusicMenuTCSettings_VolumeLimit` | Known | Controller |
| 0x0073A278 | `TCSettings_Brightness` | Known | Controller |
| 0x0073A290 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0073A2AC | `TCSettings_AudiobookSettingsTCSettings_RadioRegions` | Known | Controller |
| 0x0073A2E0 | `TCSettings_EQ` | Known | Controller |
| 0x0073A2F0 | `TSilverSettingsMenuListCntlrTSilverCntlrTCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0073A338 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0073A354 | `TCSettings_MainMenu` | Known | Controller |
| 0x0073A368 | `TCSettings_MusicMenuTCSportTimerTCSportTimerMenuTCSportTimerSessionScreen` | Known | Controller |
| 0x0073A3B4 | `TSilverCntlrTUnitTestSuiteCntlr` | Known | Controller |
| 0x0073A434 | `TCVoiceMemosTCVoiceMemosAlert` | Known | Controller |
| 0x0073A454 | `TCVoiceMemosAlert` | Known | Controller |
| 0x0073A468 | `TCVoiceMemosAlert` | Known | Controller |
| 0x0073A494 | `TCEQSetting` | Known | Controller |
| 0x0073A602 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x0073BB61 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x007418C8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00741926 | `TCNotesDispatcher` | Known | Controller |
| 0x00743664 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007436C2 | `TCNotesDispatcher` | Known | Controller |
| 0x00745400 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074545E | `TCNotesDispatcher` | Known | Controller |
| 0x0074719C | `TCLockChosenDispatcher` | Known | Controller |
| 0x007471FA | `TCNotesDispatcher` | Known | Controller |
| 0x00748F38 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00748F96 | `TCNotesDispatcher` | Known | Controller |
| 0x0074ACD4 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074AD32 | `TCNotesDispatcher` | Known | Controller |
| 0x0074CA70 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074CACE | `TCNotesDispatcher` | Known | Controller |
| 0x0074E80C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074E86A | `TCNotesDispatcher` | Known | Controller |
| 0x007505A8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00750606 | `TCNotesDispatcher` | Known | Controller |
| 0x00752344 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007523A2 | `TCNotesDispatcher` | Known | Controller |
| 0x007540E0 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075413E | `TCNotesDispatcher` | Known | Controller |
| 0x00755E7C | `TCLockChosenDispatcher` | Known | Controller |
| 0x00755EDA | `TCNotesDispatcher` | Known | Controller |
| 0x00757C18 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00757C76 | `TCNotesDispatcher` | Known | Controller |
| 0x007599B4 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00759A12 | `TCNotesDispatcher` | Known | Controller |
| 0x0075B750 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075B7AE | `TCNotesDispatcher` | Known | Controller |
| 0x0075D4EC | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075D54A | `TCNotesDispatcher` | Known | Controller |
| 0x0075F288 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075F2E6 | `TCNotesDispatcher` | Known | Controller |
| 0x00761024 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00761082 | `TCNotesDispatcher` | Known | Controller |
| 0x00762DC0 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00762E1E | `TCNotesDispatcher` | Known | Controller |
| 0x00764B5C | `TCLockChosenDispatcher` | Known | Controller |
| 0x00764BBA | `TCNotesDispatcher` | Known | Controller |
| 0x007668F8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00766956 | `TCNotesDispatcher` | Known | Controller |
| 0x00768694 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007686F2 | `TCNotesDispatcher` | Known | Controller |
| 0x0076A430 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076A48E | `TCNotesDispatcher` | Known | Controller |
| 0x0076C1CC | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076C22A | `TCNotesDispatcher` | Known | Controller |
| 0x0076DF68 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076DFC6 | `TCNotesDispatcher` | Known | Controller |
| 0x0076FD04 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076FD62 | `TCNotesDispatcher` | Known | Controller |
| 0x00771AA0 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00771AFE | `TCNotesDispatcher` | Known | Controller |
| 0x0077383C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077389A | `TCNotesDispatcher` | Known | Controller |
| 0x007755D8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00775636 | `TCNotesDispatcher` | Known | Controller |
| 0x00777374 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007773D2 | `TCNotesDispatcher` | Known | Controller |
| 0x00779110 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077916E | `TCNotesDispatcher` | Known | Controller |
| 0x0077AEAC | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077AF0A | `TCNotesDispatcher` | Known | Controller |
| 0x0077CC48 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077CCA6 | `TCNotesDispatcher` | Known | Controller |
| 0x0077E9E4 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077EA42 | `TCNotesDispatcher` | Known | Controller |
| 0x00780780 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007807DE | `TCNotesDispatcher` | Known | Controller |
| 0x0078251C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078257A | `TCNotesDispatcher` | Known | Controller |
| 0x007842B8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00784316 | `TCNotesDispatcher` | Known | Controller |
| 0x00791F6C | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x0079222E | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x00792A64 | `TCRentalDispatcher` | Known | Controller |
| 0x0079331C | `TCRentalDispatcher` | Known | Controller |
| 0x00793BD4 | `TCRentalDispatcher` | Known | Controller |
| 0x0079448C | `TCRentalDispatcher` | Known | Controller |
| 0x00794D44 | `TCRentalDispatcher` | Known | Controller |
| 0x007955FC | `TCRentalDispatcher` | Known | Controller |
| 0x00795EB4 | `TCRentalDispatcher` | Known | Controller |
| 0x0079676C | `TCRentalDispatcher` | Known | Controller |
| 0x008DD9D4 | `TCMockupModeNavScreen` | Known | Controller |
| 0x008DD9EC | `TSilverCntlr` | Known | Controller |
| 0x008DDA0C | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x008DDA5C | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x008DDA7C | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x008DDA9C | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x008DDAC0 | `TCExtrasMenu` | Known | Controller |
| 0x008DDBD0 | `TSilverCntlr` | Known | Controller |
| 0x008DDBF0 | `TCSlideshowTVOut` | Known | Controller |
| 0x008DDC04 | `TCSlideshowLCD` | Known | Controller |
| 0x008DDC14 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x008DDC2C | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x008DDC4C | `TSilverGlobalCntlr` | Known | Controller |
| 0x008DDC7C | `TSilverCntlr` | Known | Controller |
| 0x008DDCF8 | `TCSlideshowTVOut` | Known | Controller |
| 0x008DDD0C | `TCSlideshowLCD` | Known | Controller |
| 0x008DDD1C | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x008DDD34 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x008DDD54 | `TSilverCntlr` | Known | Controller |
| 0x008DDD9C | `TSilverCntlr` | Known | Controller |
| 0x008DDDBC | `TCGamesMenu` | Known | Controller |
| 0x008DDDC8 | `TCGameScreen` | Known | Controller |
| 0x0099B92C | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |

---

## 2. Settings (Toggle/Show/TC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001281A0 | `ShowSetting_EQ` | Known | User setting |
| 0x001CFAAC | `ToggleSetting_Repeat` | Known | User setting |
| 0x001CFAC8 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001CFAE0 | `ToggleSetting_TVOut` | Known | User setting |
| 0x001CFAF4 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x001F882C | `ShowSetting_Backlight` | Known | User setting |
| 0x0020D7F0 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x0020D80C | `ToggleSetting_Repeat` | Known | User setting |
| 0x0020D824 | `ToggleSetting_SortBy` | Known | User setting |
| 0x0020D83C | `ToggleSetting_ClassicUI` | Known | User setting |
| 0x0020D854 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x0020D870 | `ToggleSetting_Clicker` | Known | User setting |
| 0x0020D888 | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x0020D8A8 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x0020D8C4 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x0020D8E0 | `ShowSetting_Shuffle` | Known | User setting |
| 0x0020DA8C | `ShowSetting_Repeat` | Known | User setting |
| 0x0020DAA0 | `ShowSetting_About` | Known | User setting |
| 0x0020DAB4 | `ShowSetting_MainMenu` | Known | User setting |
| 0x0020DACC | `ShowSetting_MusicMenu` | Known | User setting |
| 0x0020DAE4 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x0020DAFC | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x0020DB18 | `ShowSetting_Brightness` | Known | User setting |
| 0x0020DB30 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x0020DB48 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x0020DB64 | `ShowSetting_EQ` | Known | User setting |
| 0x0020DB74 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x0020DD10 | `ShowSetting_Clicker` | Known | User setting |
| 0x0020DD24 | `ShowSetting_DateAndTime` | Known | User setting |
| 0x0020DD3C | `ShowSetting_SortBy` | Known | User setting |
| 0x0020DD50 | `ShowSetting_ClassicUI` | Known | User setting |
| 0x0020DD68 | `ShowSetting_Language` | Known | User setting |
| 0x0020DD80 | `ShowSetting_Legal` | Known | User setting |
| 0x0020DD94 | `ShowSetting_ResetAll` | Known | User setting |
| 0x007231F5 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x007232A5 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x007259F2 | `ShowSetting_About` | Known | User setting |
| 0x00725A94 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x00725AD8 | `ShowSetting_Shuffle` | Known | User setting |
| 0x00725B4F | `ToggleSetting_Repeat` | Known | User setting |
| 0x00725B92 | `ShowSetting_Repeat` | Known | User setting |
| 0x00725C9C | `ShowSetting_MainMenu` | Known | User setting |
| 0x00725DAC | `ShowSetting_MusicMenu` | Known | User setting |
| 0x00725E74 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x00725F3E | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x00726056 | `ShowSetting_Brightness` | Known | User setting |
| 0x0072618C | `ShowSetting_Audiobooks` | Known | User setting |
| 0x0072629D | `ShowSetting_RadioRegions` | Known | User setting |
| 0x0072639E | `ShowSetting_EQ` | Known | User setting |
| 0x0072640B | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x00726452 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x007264CF | `ToggleSetting_Clicker` | Known | User setting |
| 0x00726513 | `ShowSetting_Clicker` | Known | User setting |
| 0x0072667A | `ToggleSetting_SortBy` | Known | User setting |
| 0x007266BD | `ShowSetting_SortBy` | Known | User setting |
| 0x007267BE | `ShowSetting_Language` | Known | User setting |
| 0x007268CE | `ShowSetting_Legal` | Known | User setting |
| 0x007269FF | `ShowSetting_ResetAll` | Known | User setting |
| 0x00726B70 | `ShowSetting_Backlight` | Known | User setting |
| 0x00726C20 | `ShowSetting_Backlight` | Known | User setting |
| 0x00726CD0 | `ShowSetting_Backlight` | Known | User setting |
| 0x00726D81 | `ShowSetting_Backlight` | Known | User setting |
| 0x00726E32 | `ShowSetting_Backlight` | Known | User setting |
| 0x00726EE3 | `ShowSetting_Backlight` | Known | User setting |
| 0x00726F97 | `ShowSetting_Backlight` | Known | User setting |
| 0x00727046 | `ShowSetting_EQ` | Known | User setting |
| 0x007270BB | `ShowSetting_Language` | Known | User setting |
| 0x007AEDBC | `ToggleSetting_Repeat` | Known | User setting |
| 0x007AEDF6 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x007AEEB8 | `ToggleSetting_TVOut` | Known | User setting |
| 0x007AEEF1 | `ToggleSetting_TVSignal` | Known | User setting |

---

## 3. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0014380C | `MockupMode/MockupMode.xml` | Hidden | Developer Tool |
| 0x00143D0C | `MockupMode/` | Hidden | Developer Tool |
| 0x00248F40 | `Channel UnitTests` | Hidden | Developer Tool |
| 0x002A26A1 | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x002A26E4 | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x002A26F9 | `RTXCbug> ` | Hidden | Developer Tool |
| 0x002A30D5 | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x002BCA70 | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x003850CD | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x00385195 | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x003E360D | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x0073A3D4 | `TUnitTestSuiteTestsCntlrTSilverCntlrTCVoiceMemosMainMenuTCVoiceMemosMenuTCVoiceM` | Hidden | Developer Tool |
| 0x007D5EE0 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00812DDC | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00825AC4 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0083D9A8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008502F8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0085A2C4 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00863F60 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00879734 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00883648 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008AA800 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008C95F8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008D2C78 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0095F03D | `10TCDemoMode` | Hidden | Demo/Retail Mode |
| 0x0095F9B4 | `21TCMockupModeNavScreen` | Hidden | Developer Tool |
| 0x0095FE74 | `27TSilverCntlrTransitionAddonI10TCDemoModeE` | Hidden | Demo/Retail Mode |
| 0x0098DB1B | `Debug_ListItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x0098DB33 | `Debug_MenuItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x0098E238 | `DemoMode_Background_Image` | Hidden | Demo/Retail Mode |
| 0x0098EE26 | `DemoMode_Icon_Image` | Hidden | Demo/Retail Mode |
| 0x009909E8 | `Debug_Menu_Title` | Hidden | Debug/Diagnostic |
| 0x00990A0D | `UnitTest_Menu_Title` | Hidden | Developer Tool |
| 0x009994E4 | `UnitTestModel` | Hidden | Developer Tool |
| 0x00999EC3 | `UnitTest_ListItem` | Hidden | Developer Tool |
| 0x0099AFD9 | `DemoMode_View_Main` | Hidden | Demo/Retail Mode |
| 0x0099B1AE | `DemoMode_Icon` | Hidden | Demo/Retail Mode |
| 0x0099BF94 | `UnitTestApp` | Hidden | Developer Tool |
| 0x0099C546 | `Debug_ListItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x0099C561 | `Debug_MenuItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x0099CC77 | `DemoMode_Text_Color` | Hidden | Demo/Retail Mode |
| 0x0099D08C | `Debug_ListItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x0099D0A3 | `Debug_MenuItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009A1122 | `Debug_ListItem_UnitTest` | Hidden | Developer Tool |
| 0x009A113A | `Debug_MenuItem_UnitTest` | Hidden | Developer Tool |
| 0x009A5534 | `Debug_ListItem_Memory` | Hidden | Debug/Diagnostic |
| 0x009A554A | `Debug_MenuItem_Memory` | Hidden | Debug/Diagnostic |

---

## 4. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000067BB | `"MeCCADecode` | Known | Audio system |
| 0x00139850 | `AudioCodecs` | Known | Audio system |
| 0x00151194 | `MeCCA_RecordingBuffer` | Known | Audio system |
| 0x0017F960 | `MeCCA_PCM_Output.wav` | Known | Audio system |
| 0x00198C4C | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x001A2568 | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x001A2770 | `MeCCAVideoDecode` | Known | Audio system |
| 0x008E9B40 | `MeCCA_StreamCache` | Known | Audio system |

---

## 5. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E86C4 | `HandleWheel` | Known | Event handler |
| 0x000E86D0 | `HandlePlayPause` | Known | Event handler |
| 0x000E86E0 | `HandleSelectDown` | Known | Event handler |
| 0x000E86F4 | `HandleNext` | Known | Event handler |
| 0x000E8700 | `HandlePrevious` | Known | Event handler |
| 0x000E8710 | `HandleNextPushAndHold` | Known | Event handler |
| 0x000E8728 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000E89C0 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000E89E0 | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x000F4D38 | `HandleSelect` | Known | Event handler |
| 0x000F4D4C | `HandleHilite` | Known | Event handler |
| 0x000F50E4 | `HandleEQSettingSelected` | Known | Event handler |
| 0x000F5514 | `HandleSelect` | Known | Event handler |
| 0x000F5528 | `HandleGameHilited` | Known | Event handler |
| 0x000F57D8 | `HandleNotesSelected` | Known | Event handler |
| 0x000F57F0 | `HandleNotesPop` | Known | Event handler |
| 0x000F5800 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00103BA4 | `HandleVolumeWheel` | Known | Event handler |
| 0x00103BB8 | `HandleVolumeChange` | Known | Event handler |
| 0x00103BCC | `HandleTimerDone` | Known | Event handler |
| 0x00103BDC | `HandleFrequencyChange` | Known | Event handler |
| 0x00103C54 | `HandleTuning` | Known | Event handler |
| 0x00103C64 | `HandleTuningSelect` | Known | Event handler |
| 0x0010E7A8 | `HandleLock` | Known | Event handler |
| 0x0010E7B8 | `HandleAddressBook` | Known | Event handler |
| 0x0010EEA0 | `HandleSelect` | Known | Event handler |
| 0x0010F3D8 | `HandleExit` | Known | Event handler |
| 0x0010F3E8 | `HandleLap` | Known | Event handler |
| 0x0010F3F4 | `HandleResume` | Known | Event handler |
| 0x0010F404 | `HandleStartStop` | Known | Event handler |
| 0x0010F6B8 | `HandleWheel` | Known | Event handler |
| 0x0010F6C8 | `HandlePlayPause` | Known | Event handler |
| 0x0010F6D8 | `HandleSelectDown` | Known | Event handler |
| 0x0010F6EC | `HandleHilite` | Known | Event handler |
| 0x0010F710 | `HandleFinishRecording` | Known | Event handler |
| 0x00119E98 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x001283D4 | `HandleExitUnsupported` | Known | Event handler |
| 0x0013F08C | `HandleNotesPop` | Known | Event handler |
| 0x0013F0A0 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0013FFAC | `HandleSelect` | Known | Event handler |
| 0x0013FFC0 | `HandleWheel` | Known | Event handler |
| 0x0013FFCC | `HandleImageNext` | Known | Event handler |
| 0x0013FFDC | `HandleImagePrev` | Known | Event handler |
| 0x0013FFEC | `HandleImageLast` | Known | Event handler |
| 0x0013FFFC | `HandleImageFirst` | Known | Event handler |
| 0x00140010 | `HandlePlayPause` | Known | Event handler |
| 0x00140020 | `HandlePlay` | Known | Event handler |
| 0x0014002C | `HandlePause` | Known | Event handler |
| 0x00140038 | `HandleMikeyCenter` | Known | Event handler |
| 0x00154FBC | `HandleSelectCity` | Known | Event handler |
| 0x00154FD4 | `HandleHighlightCity` | Known | Event handler |
| 0x001560C0 | `HandleWantPopFlow` | Known | Event handler |
| 0x001560D8 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x001560F4 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x00156110 | `HandleFlowNext` | Known | Event handler |
| 0x00156120 | `HandleFlowPrev` | Known | Event handler |
| 0x00156130 | `HandleFlowWheel` | Known | Event handler |
| 0x00156140 | `HandleAlbumSelected` | Known | Event handler |
| 0x00156154 | `HandlePlayPause` | Known | Event handler |
| 0x00156164 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x001817FC | `HandleLeaveAlarm` | Known | Event handler |
| 0x00181BEC | `HandleSelect` | Known | Event handler |
| 0x00182AD4 | `HandleSelect` | Known | Event handler |
| 0x00182AE8 | `HandleWheel` | Known | Event handler |
| 0x00182AF4 | `HandleImageNext` | Known | Event handler |
| 0x00182B04 | `HandleImagePrev` | Known | Event handler |
| 0x00182B14 | `HandleImageLast` | Known | Event handler |
| 0x00182B24 | `HandleImageFirst` | Known | Event handler |
| 0x00182B38 | `HandlePlayPause` | Known | Event handler |
| 0x00182B48 | `HandlePlay` | Known | Event handler |
| 0x00182B54 | `HandlePause` | Known | Event handler |
| 0x00182B60 | `HandleMikeyCenter` | Known | Event handler |
| 0x00183008 | `HandleNew` | Known | Event handler |
| 0x00183018 | `HandleClear` | Known | Event handler |
| 0x00183024 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x00183040 | `HandleSelectIndexedSession` | Known | Event handler |
| 0x00183350 | `HandleWheel` | Known | Event handler |
| 0x00183360 | `HandleArrowUp` | Known | Event handler |
| 0x00183370 | `HandleArrowDown` | Known | Event handler |
| 0x00185894 | `HandleHiliteAlbum` | Known | Event handler |
| 0x001858AC | `HandleBrowseAlbum` | Known | Event handler |
| 0x001858C0 | `HandlePlayPause` | Known | Event handler |
| 0x0019C26C | `HandleSelect` | Known | Event handler |
| 0x0019C3FC | `HandleSelectRegion` | Known | Event handler |
| 0x0019C774 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x0019C790 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x0019C7AC | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x001B1E44 | `HandleImageWheel` | Known | Event handler |
| 0x001B1E5C | `HandlePlayPause` | Known | Event handler |
| 0x001B1E6C | `HandleBrowseLarge` | Known | Event handler |
| 0x001B1E80 | `HandleBrowseSmall` | Known | Event handler |
| 0x001B1E94 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x001B1EAC | `HandleImageNext` | Known | Event handler |
| 0x001B1EBC | `HandleImagePrev` | Known | Event handler |
| 0x001B1ECC | `HandleHilite` | Known | Event handler |
| 0x001B1EDC | `HandleImageLast` | Known | Event handler |
| 0x001B1EEC | `HandleImageFirst` | Known | Event handler |
| 0x001B1F00 | `HandleScreenNext` | Known | Event handler |
| 0x001B1F14 | `HandleScreenPrev` | Known | Event handler |
| 0x001B47DC | `HandlePlayPause` | Known | Event handler |
| 0x001B47F0 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x001B480C | `HandleNext` | Known | Event handler |
| 0x001B4818 | `HandleNextPressAndHold` | Known | Event handler |
| 0x001B4830 | `HandlePrevious` | Known | Event handler |
| 0x001B4840 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x001B485C | `HandleRemotePlayPause` | Known | Event handler |
| 0x001B4874 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x001B4898 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x001B48B0 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x001B48C8 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x001B4A6C | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x001B4A84 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x001B4A9C | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x001B4AB8 | `HandleRemoteStop` | Known | Event handler |
| 0x001B4ACC | `HandleRemotePlay` | Known | Event handler |
| 0x001B4AE0 | `HandleRemotePause` | Known | Event handler |
| 0x001B4AF4 | `HandleRemoteMute` | Known | Event handler |
| 0x001B4B08 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x001B4B20 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x001B4B38 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x001B4B54 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x001B4D5C | `HandleRemoteShuffle` | Known | Event handler |
| 0x001B4D70 | `HandleRemoteRepeat` | Known | Event handler |
| 0x001B4D84 | `HandleRemoteOn` | Known | Event handler |
| 0x001B4D98 | `HandleRemoteOff` | Known | Event handler |
| 0x001B4DA8 | `HandleRemoteBacklight` | Known | Event handler |
| 0x001B4DC0 | `HandleRemoteFFDown` | Known | Event handler |
| 0x001B4DD4 | `HandleRemoteFFUp` | Known | Event handler |
| 0x001B4DE8 | `HandleRemoteRewDown` | Known | Event handler |
| 0x001B4DFC | `HandleRemoteRewUp` | Known | Event handler |
| 0x001B4E10 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x001B4E28 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x001B4E3C | `HandleRemoteSelectDown` | Known | Event handler |
| 0x001B4E54 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x001B5004 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x001B501C | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x001B5034 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x001B5050 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x001B5068 | `HandleRemoteEvent` | Known | Event handler |
| 0x001B507C | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x001B5098 | `HandleAudioPlayPause` | Known | Event handler |
| 0x001B50B0 | `HandleAudioNext` | Known | Event handler |
| 0x001B50C0 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x001B50DC | `HandleAudioPrevious` | Known | Event handler |
| 0x001B50F0 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x001B5280 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x001B5298 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x001B52B0 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x001B52C8 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x001B52DC | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x001B52F4 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x001B530C | `HandleAudioStop` | Known | Event handler |
| 0x001B531C | `HandleAudioPlay` | Known | Event handler |
| 0x001B532C | `HandleAudioPause` | Known | Event handler |
| 0x001B5340 | `HandleAudioMute` | Known | Event handler |
| 0x001B5350 | `HandleAudioNextChapter` | Known | Event handler |
| 0x001B5368 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x001B5554 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x001B556C | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x001B5584 | `HandleAudioShuffle` | Known | Event handler |
| 0x001B5598 | `HandleAudioRepeat` | Known | Event handler |
| 0x001B55AC | `HandleAudioFFDown` | Known | Event handler |
| 0x001B55C0 | `HandleAudioFFUp` | Known | Event handler |
| 0x001B55D0 | `HandleAudioRewDown` | Known | Event handler |
| 0x001B55E4 | `HandleAudioRewUp` | Known | Event handler |
| 0x001B55F8 | `HandleVideoPlayPause` | Known | Event handler |
| 0x001B5610 | `HandleVideoNext` | Known | Event handler |
| 0x001B5620 | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x001B563C | `HandleVideoPrevious` | Known | Event handler |
| 0x001B5650 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x001B5814 | `HandleVideoStop` | Known | Event handler |
| 0x001B5824 | `HandleVideoPlay` | Known | Event handler |
| 0x001B5834 | `HandleVideoPause` | Known | Event handler |
| 0x001B5848 | `HandleVideoFFDown` | Known | Event handler |
| 0x001B585C | `HandleVideoFFUp` | Known | Event handler |
| 0x001B586C | `HandleVideoRewDown` | Known | Event handler |
| 0x001B5880 | `HandleVideoRewUp` | Known | Event handler |
| 0x001B5894 | `HandleVideoNextChapter` | Known | Event handler |
| 0x001B58AC | `HandleVideoPrevChapter` | Known | Event handler |
| 0x001B58C4 | `HandleVideoNextFrame` | Known | Event handler |
| 0x001B58DC | `HandleVideoPrevFrame` | Known | Event handler |
| 0x001B58F4 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x001B5910 | `HandleMikeyCenter` | Known | Event handler |
| 0x001B5A70 | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x001B5A90 | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x001B5AB0 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x001B5AD4 | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x001B5AF4 | `HandleMikeyAllUp` | Known | Event handler |
| 0x001B5B08 | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x001B5B1C | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x001B5B34 | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x001B5B4C | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x001C2914 | `HandleMainMenu` | Known | Event handler |
| 0x001C4A60 | `HandleLoadingCancelled` | Known | Event handler |
| 0x001C7494 | `HandlePowerSongSelected` | Known | Event handler |
| 0x001C74B0 | `HandlePowerSongChosen` | Known | Event handler |
| 0x001C74C8 | `HandleSelectTimedWorkout` | Known | Event handler |
| 0x001CDA14 | `HandleSelect` | Known | Event handler |
| 0x001CDCBC | `HandleMusicMenu` | Known | Event handler |
| 0x001CDF7C | `HandleSelect` | Known | Event handler |
| 0x001CE280 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x001CE2A0 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x001CE75C | `HandleWheel` | Known | Event handler |
| 0x001CE76C | `HandlePlayPause` | Known | Event handler |
| 0x001CE77C | `HandleSelectDown` | Known | Event handler |
| 0x001CE790 | `HandleNext` | Known | Event handler |
| 0x001CE79C | `HandlePrevious` | Known | Event handler |
| 0x001CE7AC | `HandleNextPushAndHold` | Known | Event handler |
| 0x001CE7C4 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x001CEEB8 | `HandleMenuSelection` | Known | Event handler |
| 0x001CEECC | `HandleViewAlbum` | Known | Event handler |
| 0x001CEEDC | `HandleViewArtist` | Known | Event handler |
| 0x001CEEF0 | `HandleViewCompilation` | Known | Event handler |
| 0x001CEF08 | `HandleStartGenius` | Known | Event handler |
| 0x001DBB18 | `HandleFrequencyChosen` | Known | Event handler |
| 0x001DBB30 | `HandleDateChosen` | Known | Event handler |
| 0x001DBB44 | `HandleTimeChosen` | Known | Event handler |
| 0x001DBB58 | `HandleSoundChosen` | Known | Event handler |
| 0x001DBB6C | `HandleLabelChosen` | Known | Event handler |
| 0x001DBB80 | `HandleDeleteChosen` | Known | Event handler |
| 0x001DCC60 | `HandleSelect` | Known | Event handler |
| 0x001E157C | `HandlePrev` | Known | Event handler |
| 0x001E158C | `HandleNext` | Known | Event handler |
| 0x001E1598 | `HandlePlayPause` | Known | Event handler |
| 0x001E8F54 | `HandleNextContact` | Known | Event handler |
| 0x001E8F6C | `HandlePreviousContact` | Known | Event handler |
| 0x001F0B1C | `HandleItemSelected` | Known | Event handler |
| 0x001F0D14 | `HandleRadioRegion` | Known | Event handler |
| 0x001F0EFC | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x001F5170 | `HandlePlayPause` | Known | Event handler |
| 0x001F8B08 | `HandleDelete` | Known | Event handler |
| 0x001F8B1C | `HandleSelectLozinch` | Known | Event handler |
| 0x001F8DC4 | `HandleSelect` | Known | Event handler |
| 0x001F9090 | `HandleTVOutChanged` | Known | Event handler |
| 0x001F90A8 | `HandleTVSignalChanged` | Known | Event handler |
| 0x001F90C0 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x001F90E0 | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x001F9100 | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x001F9124 | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x001F9144 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x001FBFA4 | `HandleSelectKey` | Known | Event handler |
| 0x001FC14C | `HandleSelect` | Known | Event handler |
| 0x001FCEC8 | `HandlePlayPause` | Known | Event handler |
| 0x001FCEDC | `HandleWheel` | Known | Event handler |
| 0x001FCEE8 | `HandleWheelRating` | Known | Event handler |
| 0x001FCEFC | `HandleWheelScrub` | Known | Event handler |
| 0x001FCF10 | `HandleWheelVolume` | Known | Event handler |
| 0x001FCFD0 | `HandleMenuKey` | Known | Event handler |
| 0x001FD03C | `HandleMenuLongpress` | Known | Event handler |
| 0x001FD050 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x001FDC58 | `HandleSelect` | Known | Event handler |
| 0x001FE550 | `HandleLeaveAlarm` | Known | Event handler |
| 0x001FF468 | `HandleSelect` | Known | Event handler |
| 0x001FF47C | `HandleHilite` | Known | Event handler |
| 0x001FF48C | `HandlePlayPause` | Known | Event handler |
| 0x001FF49C | `HandleAddToOTG` | Known | Event handler |
| 0x001FF4AC | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x001FF4CC | `HandleShowContextualMenu` | Known | Event handler |
| 0x00202554 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x00202D64 | `HandleSelect` | Known | Event handler |
| 0x00202D78 | `HandleWheel` | Known | Event handler |
| 0x00202D84 | `HandleWheelProgress` | Known | Event handler |
| 0x00202D98 | `HandleSelectProgress` | Known | Event handler |
| 0x00202DB0 | `HandleSelectVolume` | Known | Event handler |
| 0x00202DC4 | `HandleSelectScrub` | Known | Event handler |
| 0x00202DD8 | `HandleSelectGenius` | Known | Event handler |
| 0x00202DEC | `HandleSelectRating` | Known | Event handler |
| 0x00202E00 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x00202E18 | `HandleSelectChapterArt` | Known | Event handler |
| 0x00202E30 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x00202E4C | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x00203048 | `HandleWheelGenius` | Known | Event handler |
| 0x0020305C | `HandleWheelBrightness` | Known | Event handler |
| 0x002030C8 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x002030E8 | `HandlePushContextualMenu` | Known | Event handler |
| 0x00203104 | `HandleAddToOTG` | Known | Event handler |
| 0x00203114 | `HandleViewArtist` | Known | Event handler |
| 0x00203128 | `HandleViewAlbum` | Known | Event handler |
| 0x00203138 | `HandleViewCompilation` | Known | Event handler |
| 0x00203230 | `HandleStartGenius` | Known | Event handler |
| 0x00203244 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0020325C | `HandleAudiobookFaster` | Known | Event handler |
| 0x00203274 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0020328C | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00204D28 | `HandleStartGenius` | Known | Event handler |
| 0x0020507C | `HandleAudiobookSlower` | Known | Event handler |
| 0x00205094 | `HandleAudiobookNormal` | Known | Event handler |
| 0x002050AC | `HandleAudiobookFaster` | Known | Event handler |
| 0x002050C4 | `HandleStartGenius` | Known | Event handler |
| 0x002050D8 | `HandleAddToOTG` | Known | Event handler |
| 0x002050E8 | `HandleViewCompilation` | Known | Event handler |
| 0x00205100 | `HandleViewAlbum` | Known | Event handler |
| 0x00205110 | `HandleViewArtist` | Known | Event handler |
| 0x00205124 | `HandleCancel` | Known | Event handler |
| 0x00205BC0 | `HandleSelect` | Known | Event handler |
| 0x00205BD0 | `HandleSelectRating` | Known | Event handler |
| 0x00205BE4 | `HandleSelectProgress` | Known | Event handler |
| 0x00205BFC | `HandleWheelProgress` | Known | Event handler |
| 0x00205C10 | `HandleSelectScrub` | Known | Event handler |
| 0x00205C24 | `HandleWheelBrightness` | Known | Event handler |
| 0x00205C3C | `HandleNextTrackGotoScrub` | Known | Event handler |
| 0x00205C58 | `HandlePrevTrackGotoScrub` | Known | Event handler |
| 0x00205C74 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x0020885C | `HandleStartGenius` | Known | Event handler |
| 0x00208874 | `HandleViewArtist` | Known | Event handler |
| 0x00208888 | `HandleViewAlbum` | Known | Event handler |
| 0x00208898 | `HandleViewCompilation` | Known | Event handler |
| 0x002088B0 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00209234 | `HandleStartGenius` | Known | Event handler |
| 0x00209248 | `HandleAddToOTG` | Known | Event handler |
| 0x00209258 | `HandleViewCompilation` | Known | Event handler |
| 0x00209270 | `HandleViewAlbum` | Known | Event handler |
| 0x00209280 | `HandleViewArtist` | Known | Event handler |
| 0x00209294 | `HandleCancel` | Known | Event handler |
| 0x0020BC24 | `HandleAddToOTG` | Known | Event handler |
| 0x0020BC34 | `HandleCancel` | Known | Event handler |
| 0x0020BE28 | `HandleStartGenius` | Known | Event handler |
| 0x0020BE40 | `HandleViewAlbum` | Known | Event handler |
| 0x0020BE50 | `HandleViewArtist` | Known | Event handler |
| 0x0020BE64 | `HandleViewCompilation` | Known | Event handler |
| 0x0020BE7C | `HandleSaveGeniusPlaylist` | Known | Event handler |
| 0x0020BE98 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x0020BEB0 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0020CE18 | `HandleStartGenius` | Known | Event handler |
| 0x0020CE2C | `HandleAddToOTG` | Known | Event handler |
| 0x0020CE3C | `HandleViewCompilation` | Known | Event handler |
| 0x0020CE54 | `HandleViewAlbum` | Known | Event handler |
| 0x0020CE64 | `HandleViewArtist` | Known | Event handler |
| 0x0020CE78 | `HandleCancel` | Known | Event handler |
| 0x0020D324 | `HandleAddToOTG` | Known | Event handler |
| 0x0020D334 | `HandleCancel` | Known | Event handler |
| 0x0020DDCC | `HandleLanguage` | Known | Event handler |
| 0x0020DDDC | `HandleResetAllSettings` | Known | Event handler |
| 0x0020DDF4 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x0020E760 | `HandleSelect` | Known | Event handler |
| 0x0020E990 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x0020F870 | `HandleAddToOTG` | Known | Event handler |
| 0x0020F880 | `HandleCancel` | Known | Event handler |
| 0x00212368 | `HandleSelect` | Known | Event handler |
| 0x00212504 | `HandleSelect` | Known | Event handler |
| 0x002127A4 | `HandleNextDay` | Known | Event handler |
| 0x002127B8 | `HandlePreviousDay` | Known | Event handler |
| 0x00212FBC | `HandleMusicHilited` | Known | Event handler |
| 0x00212FD4 | `HandleVideosHilited` | Known | Event handler |
| 0x00212FE8 | `HandlePodcastsHilited` | Known | Event handler |
| 0x00213000 | `HandleGenericHilited` | Known | Event handler |
| 0x00213018 | `HandlePhotosHilited` | Known | Event handler |
| 0x0021302C | `HandleNowPlayingHilited` | Known | Event handler |
| 0x00213044 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x00213060 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00213078 | `HandleArtistsHilited` | Known | Event handler |
| 0x00213090 | `HandleGenresHilited` | Known | Event handler |
| 0x002130A4 | `HandleAlbumsHilited` | Known | Event handler |
| 0x002130B8 | `HandleCompilationsHilited` | Known | Event handler |
| 0x00213288 | `HandleComposersHilited` | Known | Event handler |
| 0x002132A0 | `HandleSongsHilited` | Known | Event handler |
| 0x002132B4 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x002132CC | `HandleGeniusHilited` | Known | Event handler |
| 0x002132E0 | `HandleTVShowsHilited` | Known | Event handler |
| 0x002132F8 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00213314 | `HandleMoviesHilited` | Known | Event handler |
| 0x00213328 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00213344 | `HandleRentalsHilited` | Known | Event handler |
| 0x0021335C | `HandleMusicSelected` | Known | Event handler |
| 0x00213370 | `HandleVideosSelected` | Known | Event handler |
| 0x00213540 | `HandlePodcastsSelected` | Known | Event handler |
| 0x00213558 | `HandlePhotosSelected` | Known | Event handler |
| 0x00213570 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00213588 | `HandleSongsSelected` | Known | Event handler |
| 0x0021359C | `HandleAlbumsSelected` | Known | Event handler |
| 0x002135B4 | `HandleCompilationsSelected` | Known | Event handler |
| 0x002135D0 | `HandleArtistsSelected` | Known | Event handler |
| 0x002135E8 | `HandleGenresSelected` | Known | Event handler |
| 0x00213600 | `HandleComposersSelected` | Known | Event handler |
| 0x00213618 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x00213634 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00213808 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00213820 | `HandleNowPlaying` | Known | Event handler |
| 0x00213834 | `HandleGotoGenius` | Known | Event handler |
| 0x00213848 | `HandleTVShowsSelected` | Known | Event handler |
| 0x00213860 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x0021387C | `HandleMoviesSelected` | Known | Event handler |
| 0x00213894 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x002138B4 | `HandleRentalsSelected` | Known | Event handler |
| 0x002138CC | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x002138E4 | `HandleLock` | Known | Event handler |
| 0x002138F0 | `HandleBacklightSelected` | Known | Event handler |
| 0x00213950 | `HandleSleepSelected` | Known | Event handler |
| 0x00213964 | `HandleNikePlusSelected` | Known | Event handler |
| 0x00216388 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00216924 | `HandleAddToOTG` | Known | Event handler |
| 0x00216934 | `HandleCancel` | Known | Event handler |
| 0x00216B04 | `HandleWheel` | Known | Event handler |
| 0x00217980 | `HandleAddToOTG` | Known | Event handler |
| 0x00217990 | `HandleCancel` | Known | Event handler |
| 0x00218150 | `HandleAddToOTG` | Known | Event handler |
| 0x00218160 | `HandleCancel` | Known | Event handler |
| 0x00218B10 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x00218D68 | `HandleNextDay` | Known | Event handler |
| 0x00218D7C | `HandlePreviousDay` | Known | Event handler |
| 0x00218FC4 | `HandleSelect` | Known | Event handler |
| 0x00219260 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00219730 | `HandleAddToOTG` | Known | Event handler |
| 0x00219740 | `HandleCancel` | Known | Event handler |
| 0x0021CE08 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x0021CE24 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x0021CE3C | `HandleStartGenius` | Known | Event handler |
| 0x0021CE50 | `HandleViewArtist` | Known | Event handler |
| 0x0021CE64 | `HandleViewAlbum` | Known | Event handler |
| 0x0021CE74 | `HandleViewCompilation` | Known | Event handler |
| 0x0021CE8C | `HandleShowContextualMenu` | Known | Event handler |
| 0x0021CEA8 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x0021CEC0 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0021E1E8 | `HandleStartGenius` | Known | Event handler |
| 0x0021E1FC | `HandleAddToOTG` | Known | Event handler |
| 0x0021E20C | `HandleViewCompilation` | Known | Event handler |
| 0x0021E224 | `HandleViewAlbum` | Known | Event handler |
| 0x0021E234 | `HandleViewArtist` | Known | Event handler |
| 0x0021E248 | `HandleCancel` | Known | Event handler |
| 0x0021E9BC | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0021EC20 | `HandleAddToOTG` | Known | Event handler |
| 0x0021EC30 | `HandleCancel` | Known | Event handler |
| 0x0021F124 | `HandleSelect` | Known | Event handler |
| 0x0021F7F0 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x002587FC | `HandleDeleteClock` | Known | Event handler |
| 0x00258814 | `HandleSelectClock` | Known | Event handler |
| 0x00258828 | `HandleHilited` | Known | Event handler |
| 0x00258838 | `HandleWheel` | Known | Event handler |
| 0x00258844 | `HandleSelectLozinch` | Known | Event handler |
| 0x00412E2A | `HandleAudioFFDown` | Known | Event handler |
| 0x00412E53 | `HandleAudioFFUp` | Known | Event handler |
| 0x00412E7E | `HandleAudioMute` | Known | Event handler |
| 0x00412EB1 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x00412EE6 | `HandleAudioNext` | Known | Event handler |
| 0x00412F16 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x00412F4D | `HandleAudioNextChapter` | Known | Event handler |
| 0x00412F87 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x00412FBB | `HandleAudioPause` | Known | Event handler |
| 0x00412FE7 | `HandleAudioPlay` | Known | Event handler |
| 0x00413015 | `HandleAudioPlayPause` | Known | Event handler |
| 0x0041304D | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x00413086 | `HandleAudioPrevious` | Known | Event handler |
| 0x004130BA | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x004130F1 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x0041312B | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x00413160 | `HandleAudioRepeat` | Known | Event handler |
| 0x0041318C | `HandleAudioRewDown` | Known | Event handler |
| 0x004131B7 | `HandleAudioRewUp` | Known | Event handler |
| 0x004131E6 | `HandleAudioShuffle` | Known | Event handler |
| 0x00413214 | `HandleAudioStop` | Known | Event handler |
| 0x00413245 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x0041327A | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x004132B1 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x004132E2 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x0041339B | `HandleNextPressAndHold` | Known | Event handler |
| 0x004133CC | `HandleNext` | Known | Event handler |
| 0x00413404 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x0041343F | `HandlePlayPause` | Known | Event handler |
| 0x00413473 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x004134A8 | `HandlePrevious` | Known | Event handler |
| 0x0041353A | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x00413582 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x004135CB | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x0041360D | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x00413645 | `HandleMikeyCenter` | Known | Event handler |
| 0x00413678 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x004136AE | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x004136E6 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x00413718 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x0041374E | `HandleRemoteBacklight` | Known | Event handler |
| 0x00413786 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x004137C0 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x004137F9 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x0041382E | `HandleRemoteEvent` | Known | Event handler |
| 0x0041385A | `HandleRemoteFFDown` | Known | Event handler |
| 0x00413885 | `HandleRemoteFFUp` | Known | Event handler |
| 0x004138B2 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x004138E1 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x00413910 | `HandleRemoteMute` | Known | Event handler |
| 0x00413942 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x0041397B | `HandleRemoteNextChapter` | Known | Event handler |
| 0x004139B7 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x004139F7 | `HandleRemoteOff` | Known | Event handler |
| 0x00413A20 | `HandleRemoteOff` | Known | Event handler |
| 0x00413A4A | `HandleRemoteOn` | Known | Event handler |
| 0x00413A76 | `HandleRemotePause` | Known | Event handler |
| 0x00413AA4 | `HandleRemotePlay` | Known | Event handler |
| 0x00413AE2 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x00413B23 | `HandleRemotePlayPause` | Known | Event handler |
| 0x00413B5A | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x00413B93 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x00413BCF | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x00413C06 | `HandleRemoteRepeat` | Known | Event handler |
| 0x00413C34 | `HandleRemoteRewDown` | Known | Event handler |
| 0x00413C61 | `HandleRemoteRewUp` | Known | Event handler |
| 0x00413C91 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x00413CC4 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x00413CF8 | `HandleRemoteShuffle` | Known | Event handler |
| 0x00413D28 | `HandleRemoteStop` | Known | Event handler |
| 0x00413D58 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x00413D8D | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x00413DC5 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x00413DFC | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x00413E35 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x00413E68 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x00413E9D | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x00413ED0 | `HandleVideoFFDown` | Known | Event handler |
| 0x00413EF9 | `HandleVideoFFUp` | Known | Event handler |
| 0x00413F2C | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x00413F61 | `HandleVideoNext` | Known | Event handler |
| 0x00413F93 | `HandleVideoNextChapter` | Known | Event handler |
| 0x00413FCA | `HandleVideoNextFrame` | Known | Event handler |
| 0x00413FFB | `HandleVideoPause` | Known | Event handler |
| 0x00414027 | `HandleVideoPlay` | Known | Event handler |
| 0x00414055 | `HandleVideoPlayPause` | Known | Event handler |
| 0x0041408D | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x004140C6 | `HandleVideoPrevious` | Known | Event handler |
| 0x004140FC | `HandleVideoPrevChapter` | Known | Event handler |
| 0x00414133 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x00414162 | `HandleVideoRewDown` | Known | Event handler |
| 0x0041418D | `HandleVideoRewUp` | Known | Event handler |
| 0x004141B9 | `HandleVideoStop` | Known | Event handler |
| 0x00719F02 | `HandleAddressBook` | Known | Event handler |
| 0x0071A49E | `HandleSelect` | Known | Event handler |
| 0x0071A4D9 | `HandleHilite` | Known | Event handler |
| 0x0071A55A | `HandleSelectRegion` | Known | Event handler |
| 0x0071A5FA | `HandleSelectRegion` | Known | Event handler |
| 0x0071A696 | `HandleSelectRegion` | Known | Event handler |
| 0x0071A73A | `HandleSelectRegion` | Known | Event handler |
| 0x0071A7E0 | `HandleSelectRegion` | Known | Event handler |
| 0x0071A880 | `HandleSelectRegion` | Known | Event handler |
| 0x0071A92C | `HandleSelectRegion` | Known | Event handler |
| 0x0071A9CE | `HandleSelectRegion` | Known | Event handler |
| 0x0071AA7E | `HandleSelectCity` | Known | Event handler |
| 0x0071AAEA | `HandleHighlightCity` | Known | Event handler |
| 0x0071AB23 | `HandleSelectCity` | Known | Event handler |
| 0x0071AB8F | `HandleHighlightCity` | Known | Event handler |
| 0x0071ABC8 | `HandleSelectCity` | Known | Event handler |
| 0x0071AC34 | `HandleHighlightCity` | Known | Event handler |
| 0x0071AC6D | `HandleSelectCity` | Known | Event handler |
| 0x0071ACD9 | `HandleHighlightCity` | Known | Event handler |
| 0x0071AD12 | `HandleSelectCity` | Known | Event handler |
| 0x0071AD7E | `HandleHighlightCity` | Known | Event handler |
| 0x0071ADB7 | `HandleSelectCity` | Known | Event handler |
| 0x0071AE23 | `HandleHighlightCity` | Known | Event handler |
| 0x0071AE5C | `HandleSelectCity` | Known | Event handler |
| 0x0071AEC8 | `HandleHighlightCity` | Known | Event handler |
| 0x0071AF01 | `HandleSelectCity` | Known | Event handler |
| 0x0071AF6D | `HandleHighlightCity` | Known | Event handler |
| 0x0071AFA6 | `HandleSelectCity` | Known | Event handler |
| 0x0071B012 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B04B | `HandleSelectCity` | Known | Event handler |
| 0x0071B0B7 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B0F0 | `HandleSelectCity` | Known | Event handler |
| 0x0071B15C | `HandleHighlightCity` | Known | Event handler |
| 0x0071B195 | `HandleSelectCity` | Known | Event handler |
| 0x0071B201 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B23A | `HandleSelectCity` | Known | Event handler |
| 0x0071B2A6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B2DF | `HandleSelectCity` | Known | Event handler |
| 0x0071B34B | `HandleHighlightCity` | Known | Event handler |
| 0x0071B384 | `HandleSelectCity` | Known | Event handler |
| 0x0071B3F0 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B429 | `HandleSelectCity` | Known | Event handler |
| 0x0071B495 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B4CE | `HandleSelectCity` | Known | Event handler |
| 0x0071B53A | `HandleHighlightCity` | Known | Event handler |
| 0x0071B573 | `HandleSelectCity` | Known | Event handler |
| 0x0071B5DF | `HandleHighlightCity` | Known | Event handler |
| 0x0071B618 | `HandleSelectCity` | Known | Event handler |
| 0x0071B684 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B6BD | `HandleSelectCity` | Known | Event handler |
| 0x0071B729 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B762 | `HandleSelectCity` | Known | Event handler |
| 0x0071B7CE | `HandleHighlightCity` | Known | Event handler |
| 0x0071B807 | `HandleSelectCity` | Known | Event handler |
| 0x0071B873 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B8AC | `HandleSelectCity` | Known | Event handler |
| 0x0071B918 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B951 | `HandleSelectCity` | Known | Event handler |
| 0x0071B9BD | `HandleHighlightCity` | Known | Event handler |
| 0x0071B9F6 | `HandleSelectCity` | Known | Event handler |
| 0x0071BA62 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BA9B | `HandleSelectCity` | Known | Event handler |
| 0x0071BB07 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BB40 | `HandleSelectCity` | Known | Event handler |
| 0x0071BBAC | `HandleHighlightCity` | Known | Event handler |
| 0x0071BBE5 | `HandleSelectCity` | Known | Event handler |
| 0x0071BC51 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BC8A | `HandleSelectCity` | Known | Event handler |
| 0x0071BCF6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BD2F | `HandleSelectCity` | Known | Event handler |
| 0x0071BD9B | `HandleHighlightCity` | Known | Event handler |
| 0x0071BDD4 | `HandleSelectCity` | Known | Event handler |
| 0x0071BE40 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BE7E | `HandleSelectCity` | Known | Event handler |
| 0x0071BEEA | `HandleHighlightCity` | Known | Event handler |
| 0x0071BF23 | `HandleSelectCity` | Known | Event handler |
| 0x0071BF8F | `HandleHighlightCity` | Known | Event handler |
| 0x0071BFC8 | `HandleSelectCity` | Known | Event handler |
| 0x0071C034 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C06D | `HandleSelectCity` | Known | Event handler |
| 0x0071C0D9 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C112 | `HandleSelectCity` | Known | Event handler |
| 0x0071C17E | `HandleHighlightCity` | Known | Event handler |
| 0x0071C1B7 | `HandleSelectCity` | Known | Event handler |
| 0x0071C223 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C25C | `HandleSelectCity` | Known | Event handler |
| 0x0071C2C8 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C301 | `HandleSelectCity` | Known | Event handler |
| 0x0071C36D | `HandleHighlightCity` | Known | Event handler |
| 0x0071C3A6 | `HandleSelectCity` | Known | Event handler |
| 0x0071C412 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C44B | `HandleSelectCity` | Known | Event handler |
| 0x0071C4B7 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C4F0 | `HandleSelectCity` | Known | Event handler |
| 0x0071C55C | `HandleHighlightCity` | Known | Event handler |
| 0x0071C595 | `HandleSelectCity` | Known | Event handler |
| 0x0071C601 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C63A | `HandleSelectCity` | Known | Event handler |
| 0x0071C6A6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C6DF | `HandleSelectCity` | Known | Event handler |
| 0x0071C74B | `HandleHighlightCity` | Known | Event handler |
| 0x0071C784 | `HandleSelectCity` | Known | Event handler |
| 0x0071C7F0 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C829 | `HandleSelectCity` | Known | Event handler |
| 0x0071C895 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C8CE | `HandleSelectCity` | Known | Event handler |
| 0x0071C93A | `HandleHighlightCity` | Known | Event handler |
| 0x0071C973 | `HandleSelectCity` | Known | Event handler |
| 0x0071C9DF | `HandleHighlightCity` | Known | Event handler |
| 0x0071CA18 | `HandleSelectCity` | Known | Event handler |
| 0x0071CA84 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CABD | `HandleSelectCity` | Known | Event handler |
| 0x0071CB29 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CB62 | `HandleSelectCity` | Known | Event handler |
| 0x0071CBCE | `HandleHighlightCity` | Known | Event handler |
| 0x0071CC07 | `HandleSelectCity` | Known | Event handler |
| 0x0071CC73 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CCAC | `HandleSelectCity` | Known | Event handler |
| 0x0071CD18 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CD51 | `HandleSelectCity` | Known | Event handler |
| 0x0071CDBD | `HandleHighlightCity` | Known | Event handler |
| 0x0071CDF6 | `HandleSelectCity` | Known | Event handler |
| 0x0071CE62 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CE9B | `HandleSelectCity` | Known | Event handler |
| 0x0071CF07 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CF40 | `HandleSelectCity` | Known | Event handler |
| 0x0071CFAC | `HandleHighlightCity` | Known | Event handler |
| 0x0071CFE5 | `HandleSelectCity` | Known | Event handler |
| 0x0071D051 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D08A | `HandleSelectCity` | Known | Event handler |
| 0x0071D0F6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D12F | `HandleSelectCity` | Known | Event handler |
| 0x0071D19B | `HandleHighlightCity` | Known | Event handler |
| 0x0071D1D4 | `HandleSelectCity` | Known | Event handler |
| 0x0071D240 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D279 | `HandleSelectCity` | Known | Event handler |
| 0x0071D2E5 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D31E | `HandleSelectCity` | Known | Event handler |
| 0x0071D38A | `HandleHighlightCity` | Known | Event handler |
| 0x0071D3C3 | `HandleSelectCity` | Known | Event handler |
| 0x0071D42F | `HandleHighlightCity` | Known | Event handler |
| 0x0071D468 | `HandleSelectCity` | Known | Event handler |
| 0x0071D4D4 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D50D | `HandleSelectCity` | Known | Event handler |
| 0x0071D579 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D5B2 | `HandleSelectCity` | Known | Event handler |
| 0x0071D61E | `HandleHighlightCity` | Known | Event handler |
| 0x0071D657 | `HandleSelectCity` | Known | Event handler |
| 0x0071D6C3 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D6FC | `HandleSelectCity` | Known | Event handler |
| 0x0071D768 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D7A1 | `HandleSelectCity` | Known | Event handler |
| 0x0071D80D | `HandleHighlightCity` | Known | Event handler |
| 0x0071D846 | `HandleSelectCity` | Known | Event handler |
| 0x0071D8B2 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D8EB | `HandleSelectCity` | Known | Event handler |
| 0x0071D957 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D990 | `HandleSelectCity` | Known | Event handler |
| 0x0071D9FC | `HandleHighlightCity` | Known | Event handler |
| 0x0071DA35 | `HandleSelectCity` | Known | Event handler |
| 0x0071DAA1 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DADA | `HandleSelectCity` | Known | Event handler |
| 0x0071DB46 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DB7F | `HandleSelectCity` | Known | Event handler |
| 0x0071DBEB | `HandleHighlightCity` | Known | Event handler |
| 0x0071DC24 | `HandleSelectCity` | Known | Event handler |
| 0x0071DC90 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DCC9 | `HandleSelectCity` | Known | Event handler |
| 0x0071DD35 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DD6E | `HandleSelectCity` | Known | Event handler |
| 0x0071DDDA | `HandleHighlightCity` | Known | Event handler |
| 0x0071DE13 | `HandleSelectCity` | Known | Event handler |
| 0x0071DE7F | `HandleHighlightCity` | Known | Event handler |
| 0x0071DEB8 | `HandleSelectCity` | Known | Event handler |
| 0x0071DF24 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DF5D | `HandleSelectCity` | Known | Event handler |
| 0x0071DFC9 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E002 | `HandleSelectCity` | Known | Event handler |
| 0x0071E06E | `HandleHighlightCity` | Known | Event handler |
| 0x0071E0A7 | `HandleSelectCity` | Known | Event handler |
| 0x0071E113 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E14C | `HandleSelectCity` | Known | Event handler |
| 0x0071E1B8 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E1F1 | `HandleSelectCity` | Known | Event handler |
| 0x0071E25D | `HandleHighlightCity` | Known | Event handler |
| 0x0071E296 | `HandleSelectCity` | Known | Event handler |
| 0x0071E302 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E342 | `HandleSelectCity` | Known | Event handler |
| 0x0071E3AE | `HandleHighlightCity` | Known | Event handler |
| 0x0071E3E7 | `HandleSelectCity` | Known | Event handler |
| 0x0071E453 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E48C | `HandleSelectCity` | Known | Event handler |
| 0x0071E4F8 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E536 | `HandleSelectCity` | Known | Event handler |
| 0x0071E5A2 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E5DB | `HandleSelectCity` | Known | Event handler |
| 0x0071E647 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E680 | `HandleSelectCity` | Known | Event handler |
| 0x0071E6EC | `HandleHighlightCity` | Known | Event handler |
| 0x0071E725 | `HandleSelectCity` | Known | Event handler |
| 0x0071E791 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E7CA | `HandleSelectCity` | Known | Event handler |
| 0x0071E836 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E86F | `HandleSelectCity` | Known | Event handler |
| 0x0071E8DB | `HandleHighlightCity` | Known | Event handler |
| 0x0071E914 | `HandleSelectCity` | Known | Event handler |
| 0x0071E980 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E9B9 | `HandleSelectCity` | Known | Event handler |
| 0x0071EA25 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EA62 | `HandleSelectCity` | Known | Event handler |
| 0x0071EACE | `HandleHighlightCity` | Known | Event handler |
| 0x0071EB07 | `HandleSelectCity` | Known | Event handler |
| 0x0071EB73 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EBAC | `HandleSelectCity` | Known | Event handler |
| 0x0071EC18 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EC51 | `HandleSelectCity` | Known | Event handler |
| 0x0071ECBD | `HandleHighlightCity` | Known | Event handler |
| 0x0071ECF6 | `HandleSelectCity` | Known | Event handler |
| 0x0071ED62 | `HandleHighlightCity` | Known | Event handler |
| 0x0071ED9B | `HandleSelectCity` | Known | Event handler |
| 0x0071EE07 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EE40 | `HandleSelectCity` | Known | Event handler |
| 0x0071EEAC | `HandleHighlightCity` | Known | Event handler |
| 0x0071EEE5 | `HandleSelectCity` | Known | Event handler |
| 0x0071EF51 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EF8A | `HandleSelectCity` | Known | Event handler |
| 0x0071EFF6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F02F | `HandleSelectCity` | Known | Event handler |
| 0x0071F09B | `HandleHighlightCity` | Known | Event handler |
| 0x0071F0D4 | `HandleSelectCity` | Known | Event handler |
| 0x0071F140 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F179 | `HandleSelectCity` | Known | Event handler |
| 0x0071F1E5 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F21E | `HandleSelectCity` | Known | Event handler |
| 0x0071F28A | `HandleHighlightCity` | Known | Event handler |
| 0x0071F2C3 | `HandleSelectCity` | Known | Event handler |
| 0x0071F32F | `HandleHighlightCity` | Known | Event handler |
| 0x0071F368 | `HandleSelectCity` | Known | Event handler |
| 0x0071F3D4 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F40D | `HandleSelectCity` | Known | Event handler |
| 0x0071F479 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F4B2 | `HandleSelectCity` | Known | Event handler |
| 0x0071F51E | `HandleHighlightCity` | Known | Event handler |
| 0x0071F557 | `HandleSelectCity` | Known | Event handler |
| 0x0071F5C3 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F5FC | `HandleSelectCity` | Known | Event handler |
| 0x0071F668 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F6A1 | `HandleSelectCity` | Known | Event handler |
| 0x0071F70D | `HandleHighlightCity` | Known | Event handler |
| 0x0071F746 | `HandleSelectCity` | Known | Event handler |
| 0x0071F7B2 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F7EB | `HandleSelectCity` | Known | Event handler |
| 0x0071F857 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F890 | `HandleSelectCity` | Known | Event handler |
| 0x0071F8FC | `HandleHighlightCity` | Known | Event handler |
| 0x0071F935 | `HandleSelectCity` | Known | Event handler |
| 0x0071F9A1 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F9DA | `HandleSelectCity` | Known | Event handler |
| 0x0071FA46 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FA7F | `HandleSelectCity` | Known | Event handler |
| 0x0071FAEB | `HandleHighlightCity` | Known | Event handler |
| 0x0071FB24 | `HandleSelectCity` | Known | Event handler |
| 0x0071FB90 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FBC9 | `HandleSelectCity` | Known | Event handler |
| 0x0071FC35 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FC6E | `HandleSelectCity` | Known | Event handler |
| 0x0071FCDA | `HandleHighlightCity` | Known | Event handler |
| 0x0071FD13 | `HandleSelectCity` | Known | Event handler |
| 0x0071FD7F | `HandleHighlightCity` | Known | Event handler |
| 0x0071FDB8 | `HandleSelectCity` | Known | Event handler |
| 0x0071FE24 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FE5D | `HandleSelectCity` | Known | Event handler |
| 0x0071FEC9 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FF02 | `HandleSelectCity` | Known | Event handler |
| 0x0071FF6E | `HandleHighlightCity` | Known | Event handler |
| 0x0071FFA7 | `HandleSelectCity` | Known | Event handler |
| 0x00720013 | `HandleHighlightCity` | Known | Event handler |
| 0x00720052 | `HandleSelectCity` | Known | Event handler |
| 0x007200BE | `HandleHighlightCity` | Known | Event handler |
| 0x007200F7 | `HandleSelectCity` | Known | Event handler |
| 0x00720163 | `HandleHighlightCity` | Known | Event handler |
| 0x0072019C | `HandleSelectCity` | Known | Event handler |
| 0x00720208 | `HandleHighlightCity` | Known | Event handler |
| 0x00720241 | `HandleSelectCity` | Known | Event handler |
| 0x007202AD | `HandleHighlightCity` | Known | Event handler |
| 0x007202E6 | `HandleSelectCity` | Known | Event handler |
| 0x00720352 | `HandleHighlightCity` | Known | Event handler |
| 0x0072038B | `HandleSelectCity` | Known | Event handler |
| 0x007203F7 | `HandleHighlightCity` | Known | Event handler |
| 0x00720430 | `HandleSelectCity` | Known | Event handler |
| 0x0072049C | `HandleHighlightCity` | Known | Event handler |
| 0x007204D5 | `HandleSelectCity` | Known | Event handler |
| 0x00720541 | `HandleHighlightCity` | Known | Event handler |
| 0x0072057A | `HandleSelectCity` | Known | Event handler |
| 0x007205E6 | `HandleHighlightCity` | Known | Event handler |
| 0x0072061F | `HandleSelectCity` | Known | Event handler |
| 0x0072068B | `HandleHighlightCity` | Known | Event handler |
| 0x007206C4 | `HandleSelectCity` | Known | Event handler |
| 0x00720730 | `HandleHighlightCity` | Known | Event handler |
| 0x00720769 | `HandleSelectCity` | Known | Event handler |
| 0x007207D5 | `HandleHighlightCity` | Known | Event handler |
| 0x0072080E | `HandleSelectCity` | Known | Event handler |
| 0x0072087A | `HandleHighlightCity` | Known | Event handler |
| 0x007208B3 | `HandleSelectCity` | Known | Event handler |
| 0x0072091F | `HandleHighlightCity` | Known | Event handler |
| 0x00720958 | `HandleSelectCity` | Known | Event handler |
| 0x007209C4 | `HandleHighlightCity` | Known | Event handler |
| 0x007209FD | `HandleSelectCity` | Known | Event handler |
| 0x00720A69 | `HandleHighlightCity` | Known | Event handler |
| 0x00720AA2 | `HandleSelectCity` | Known | Event handler |
| 0x00720B0E | `HandleHighlightCity` | Known | Event handler |
| 0x00720B47 | `HandleSelectCity` | Known | Event handler |
| 0x00720BB3 | `HandleHighlightCity` | Known | Event handler |
| 0x00720BEC | `HandleSelectCity` | Known | Event handler |
| 0x00720C58 | `HandleHighlightCity` | Known | Event handler |
| 0x00720C91 | `HandleSelectCity` | Known | Event handler |
| 0x00720CFD | `HandleHighlightCity` | Known | Event handler |
| 0x00720D36 | `HandleSelectCity` | Known | Event handler |
| 0x00720DA2 | `HandleHighlightCity` | Known | Event handler |
| 0x00720DDB | `HandleSelectCity` | Known | Event handler |
| 0x00720E47 | `HandleHighlightCity` | Known | Event handler |
| 0x00720E80 | `HandleSelectCity` | Known | Event handler |
| 0x00720EEC | `HandleHighlightCity` | Known | Event handler |
| 0x00720F25 | `HandleSelectCity` | Known | Event handler |
| 0x00720F91 | `HandleHighlightCity` | Known | Event handler |
| 0x00720FCA | `HandleSelectCity` | Known | Event handler |
| 0x00721036 | `HandleHighlightCity` | Known | Event handler |
| 0x0072106F | `HandleSelectCity` | Known | Event handler |
| 0x007210DB | `HandleHighlightCity` | Known | Event handler |
| 0x00721114 | `HandleSelectCity` | Known | Event handler |
| 0x00721180 | `HandleHighlightCity` | Known | Event handler |
| 0x007211B9 | `HandleSelectCity` | Known | Event handler |
| 0x00721225 | `HandleHighlightCity` | Known | Event handler |
| 0x0072125E | `HandleSelectCity` | Known | Event handler |
| 0x007212CA | `HandleHighlightCity` | Known | Event handler |
| 0x00721303 | `HandleSelectCity` | Known | Event handler |
| 0x0072136F | `HandleHighlightCity` | Known | Event handler |
| 0x007213A8 | `HandleSelectCity` | Known | Event handler |
| 0x00721414 | `HandleHighlightCity` | Known | Event handler |
| 0x0072144D | `HandleSelectCity` | Known | Event handler |
| 0x007214B9 | `HandleHighlightCity` | Known | Event handler |
| 0x007214F2 | `HandleSelectCity` | Known | Event handler |
| 0x0072155E | `HandleHighlightCity` | Known | Event handler |
| 0x00721597 | `HandleSelectCity` | Known | Event handler |
| 0x00721603 | `HandleHighlightCity` | Known | Event handler |
| 0x0072163C | `HandleSelectCity` | Known | Event handler |
| 0x007216A8 | `HandleHighlightCity` | Known | Event handler |
| 0x007216E1 | `HandleSelectCity` | Known | Event handler |
| 0x0072174D | `HandleHighlightCity` | Known | Event handler |
| 0x00721786 | `HandleSelectCity` | Known | Event handler |
| 0x007217F2 | `HandleHighlightCity` | Known | Event handler |
| 0x0072182B | `HandleSelectCity` | Known | Event handler |
| 0x00721897 | `HandleHighlightCity` | Known | Event handler |
| 0x007218D0 | `HandleSelectCity` | Known | Event handler |
| 0x0072193C | `HandleHighlightCity` | Known | Event handler |
| 0x00721975 | `HandleSelectCity` | Known | Event handler |
| 0x007219E1 | `HandleHighlightCity` | Known | Event handler |
| 0x00721A1A | `HandleSelectCity` | Known | Event handler |
| 0x00721A86 | `HandleHighlightCity` | Known | Event handler |
| 0x00721ABF | `HandleSelectCity` | Known | Event handler |
| 0x00721B2B | `HandleHighlightCity` | Known | Event handler |
| 0x00721B64 | `HandleSelectCity` | Known | Event handler |
| 0x00721BD0 | `HandleHighlightCity` | Known | Event handler |
| 0x00721C09 | `HandleSelectCity` | Known | Event handler |
| 0x00721C75 | `HandleHighlightCity` | Known | Event handler |
| 0x00721CAE | `HandleSelectCity` | Known | Event handler |
| 0x00721D1A | `HandleHighlightCity` | Known | Event handler |
| 0x00721D53 | `HandleSelectCity` | Known | Event handler |
| 0x00721DBF | `HandleHighlightCity` | Known | Event handler |
| 0x00721DF8 | `HandleSelectCity` | Known | Event handler |
| 0x00721E64 | `HandleHighlightCity` | Known | Event handler |
| 0x00721E9D | `HandleSelectCity` | Known | Event handler |
| 0x00721F09 | `HandleHighlightCity` | Known | Event handler |
| 0x00721F42 | `HandleSelectCity` | Known | Event handler |
| 0x00721FAE | `HandleHighlightCity` | Known | Event handler |
| 0x00721FE7 | `HandleSelectCity` | Known | Event handler |
| 0x00722053 | `HandleHighlightCity` | Known | Event handler |
| 0x00722092 | `HandleSelectCity` | Known | Event handler |
| 0x007220FE | `HandleHighlightCity` | Known | Event handler |
| 0x00722137 | `HandleSelectCity` | Known | Event handler |
| 0x007221A3 | `HandleHighlightCity` | Known | Event handler |
| 0x007221DC | `HandleSelectCity` | Known | Event handler |
| 0x00722248 | `HandleHighlightCity` | Known | Event handler |
| 0x00722281 | `HandleSelectCity` | Known | Event handler |
| 0x007222ED | `HandleHighlightCity` | Known | Event handler |
| 0x00722326 | `HandleSelectCity` | Known | Event handler |
| 0x00722392 | `HandleHighlightCity` | Known | Event handler |
| 0x007223D2 | `HandleSelectCity` | Known | Event handler |
| 0x0072243E | `HandleHighlightCity` | Known | Event handler |
| 0x00722477 | `HandleSelectCity` | Known | Event handler |
| 0x007224E3 | `HandleHighlightCity` | Known | Event handler |
| 0x0072251C | `HandleSelectCity` | Known | Event handler |
| 0x00722588 | `HandleHighlightCity` | Known | Event handler |
| 0x007225C1 | `HandleSelectCity` | Known | Event handler |
| 0x0072262D | `HandleHighlightCity` | Known | Event handler |
| 0x00722666 | `HandleSelectCity` | Known | Event handler |
| 0x007226D2 | `HandleHighlightCity` | Known | Event handler |
| 0x0072270B | `HandleSelectCity` | Known | Event handler |
| 0x00722777 | `HandleHighlightCity` | Known | Event handler |
| 0x007227B0 | `HandleSelectCity` | Known | Event handler |
| 0x0072281C | `HandleHighlightCity` | Known | Event handler |
| 0x00722855 | `HandleSelectCity` | Known | Event handler |
| 0x007228C1 | `HandleHighlightCity` | Known | Event handler |
| 0x007228FA | `HandleSelectCity` | Known | Event handler |
| 0x00722966 | `HandleHighlightCity` | Known | Event handler |
| 0x0072299F | `HandleSelectCity` | Known | Event handler |
| 0x00722A0B | `HandleHighlightCity` | Known | Event handler |
| 0x00722A44 | `HandleSelectCity` | Known | Event handler |
| 0x00722AB0 | `HandleHighlightCity` | Known | Event handler |
| 0x00722AE9 | `HandleSelectCity` | Known | Event handler |
| 0x00722B55 | `HandleHighlightCity` | Known | Event handler |
| 0x00722B8E | `HandleSelectCity` | Known | Event handler |
| 0x00722BFA | `HandleHighlightCity` | Known | Event handler |
| 0x00722C33 | `HandleSelectCity` | Known | Event handler |
| 0x00722C9F | `HandleHighlightCity` | Known | Event handler |
| 0x00722CD8 | `HandleSelectCity` | Known | Event handler |
| 0x00722D44 | `HandleHighlightCity` | Known | Event handler |
| 0x00722D7D | `HandleSelectCity` | Known | Event handler |
| 0x00722DE9 | `HandleHighlightCity` | Known | Event handler |
| 0x00722E22 | `HandleSelectCity` | Known | Event handler |
| 0x00722E8E | `HandleHighlightCity` | Known | Event handler |
| 0x00723386 | `HandleMusicSelected` | Known | Event handler |
| 0x007233C8 | `HandleMusicHilited` | Known | Event handler |
| 0x00723400 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00723446 | `HandleMusicHilited` | Known | Event handler |
| 0x0072347E | `HandleGotoGenius` | Known | Event handler |
| 0x007234BD | `HandleGeniusHilited` | Known | Event handler |
| 0x007234F6 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x0072353C | `HandlePlaylistsHilited` | Known | Event handler |
| 0x00723578 | `HandleArtistsSelected` | Known | Event handler |
| 0x007235BC | `HandleArtistsHilited` | Known | Event handler |
| 0x007235F6 | `HandleAlbumsSelected` | Known | Event handler |
| 0x00723639 | `HandleAlbumsHilited` | Known | Event handler |
| 0x00723672 | `HandleCompilationsSelected` | Known | Event handler |
| 0x007236BB | `HandleCompilationsHilited` | Known | Event handler |
| 0x007236FA | `HandleSongsSelected` | Known | Event handler |
| 0x0072373C | `HandleSongsHilited` | Known | Event handler |
| 0x00723774 | `HandleGenresSelected` | Known | Event handler |
| 0x007237B7 | `HandleGenresHilited` | Known | Event handler |
| 0x007237F0 | `HandleComposersSelected` | Known | Event handler |
| 0x00723836 | `HandleComposersHilited` | Known | Event handler |
| 0x00723872 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x007238B9 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00723978 | `HandleMusicHilited` | Known | Event handler |
| 0x007239B0 | `HandleVideosSelected` | Known | Event handler |
| 0x007239F3 | `HandleVideosHilited` | Known | Event handler |
| 0x00723A2C | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00723A77 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00723AB8 | `HandleMoviesSelected` | Known | Event handler |
| 0x00723AFB | `HandleMoviesHilited` | Known | Event handler |
| 0x00723B34 | `HandleTVShowsSelected` | Known | Event handler |
| 0x00723B78 | `HandleTVShowsHilited` | Known | Event handler |
| 0x00723BB2 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00723BFA | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00723C38 | `HandleRentalsSelected` | Known | Event handler |
| 0x00723C7C | `HandleRentalsHilited` | Known | Event handler |
| 0x00723CB6 | `HandlePhotosSelected` | Known | Event handler |
| 0x00723CF9 | `HandlePhotosHilited` | Known | Event handler |
| 0x00723D32 | `HandlePhotosSelected` | Known | Event handler |
| 0x00723D75 | `HandlePhotosHilited` | Known | Event handler |
| 0x00723DAE | `HandlePodcastsSelected` | Known | Event handler |
| 0x00723DF3 | `HandlePodcastsHilited` | Known | Event handler |
| 0x00723EA6 | `HandleGenericHilited` | Known | Event handler |
| 0x00723F9F | `HandleGenericHilited` | Known | Event handler |
| 0x00724484 | `HandleLock` | Known | Event handler |
| 0x007245F5 | `HandleNikePlusSelected` | Known | Event handler |
| 0x0072463A | `HandleGenericHilited` | Known | Event handler |
| 0x00724740 | `HandleGenericHilited` | Known | Event handler |
| 0x0072483F | `HandleGenericHilited` | Known | Event handler |
| 0x0072492C | `HandleGenericHilited` | Known | Event handler |
| 0x00724A29 | `HandleGenericHilited` | Known | Event handler |
| 0x00724AA3 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00724AEC | `HandleGenericHilited` | Known | Event handler |
| 0x00724B65 | `HandleBacklightSelected` | Known | Event handler |
| 0x00724BAB | `HandleGenericHilited` | Known | Event handler |
| 0x00724C26 | `HandleSleepSelected` | Known | Event handler |
| 0x00724C68 | `HandleGenericHilited` | Known | Event handler |
| 0x00724CDF | `HandleNowPlaying` | Known | Event handler |
| 0x00724D57 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x00724D9A | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00724DE0 | `HandleMusicHilited` | Known | Event handler |
| 0x00724E18 | `HandleGotoGenius` | Known | Event handler |
| 0x00724E4E | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00724E94 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x00724ED2 | `HandleArtistsSelected` | Known | Event handler |
| 0x00724F16 | `HandleArtistsHilited` | Known | Event handler |
| 0x00724F50 | `HandleAlbumsSelected` | Known | Event handler |
| 0x00724F93 | `HandleAlbumsHilited` | Known | Event handler |
| 0x00724FCC | `HandleCompilationsSelected` | Known | Event handler |
| 0x00725015 | `HandleCompilationsHilited` | Known | Event handler |
| 0x00725054 | `HandleSongsSelected` | Known | Event handler |
| 0x00725096 | `HandleSongsHilited` | Known | Event handler |
| 0x00725141 | `HandleGenericHilited` | Known | Event handler |
| 0x007251B9 | `HandleGenresSelected` | Known | Event handler |
| 0x007251FC | `HandleGenresHilited` | Known | Event handler |
| 0x00725235 | `HandleComposersSelected` | Known | Event handler |
| 0x0072527B | `HandleComposersHilited` | Known | Event handler |
| 0x007252B7 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x007252FE | `HandleAudiobooksHilited` | Known | Event handler |
| 0x007253BD | `HandleMusicHilited` | Known | Event handler |
| 0x00725431 | `HandlePlayPause` | Known | Event handler |
| 0x00725466 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x00725550 | `HandleSelect` | Known | Event handler |
| 0x00725596 | `HandleMoviesSelected` | Known | Event handler |
| 0x007255D9 | `HandleMoviesHilited` | Known | Event handler |
| 0x00725612 | `HandleRentalsSelected` | Known | Event handler |
| 0x00725656 | `HandleRentalsHilited` | Known | Event handler |
| 0x00725690 | `HandleTVShowsSelected` | Known | Event handler |
| 0x007256D4 | `HandleTVShowsHilited` | Known | Event handler |
| 0x0072570E | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00725756 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00725794 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x007257DF | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x007258A5 | `HandleVideosHilited` | Known | Event handler |
| 0x00725EF3 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x00726A7A | `HandleMainMenu` | Known | Event handler |
| 0x00726AB2 | `HandleMusicMenu` | Known | Event handler |
| 0x00726FDA | `HandleRadioRegion` | Known | Event handler |
| 0x0072707E | `HandleLanguage` | Known | Event handler |
| 0x00727184 | `HandleNew` | Known | Event handler |
| 0x007271FF | `HandleClear` | Known | Event handler |
| 0x00727230 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x007272EC | `HandleSelectIndexedSession` | Known | Event handler |
| 0x00727455 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x007274A8 | `HandleSelect` | Known | Event handler |
| 0x007275D2 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x0072760C | `HandleEQSettingSelected` | Known | Event handler |
| 0x00727644 | `HandleEQSettingSelected` | Known | Event handler |
| 0x0073A586 | `HandleMenuSelection` | Known | Event handler |
| 0x0073A8CB | `HandleLoadingCancelled` | Known | Event handler |
| 0x0073A967 | `HandleLoadingCancelled` | Known | Event handler |
| 0x0073AA34 | `HandleItemSelected` | Known | Event handler |
| 0x0073AB7F | `HandleNextContact` | Known | Event handler |
| 0x0073ABAB | `HandlePreviousContact` | Known | Event handler |
| 0x0073ABE1 | `HandleSelectKey` | Known | Event handler |
| 0x0073B1F2 | `HandleSelect` | Known | Event handler |
| 0x0073B519 | `HandleDateChosen` | Known | Event handler |
| 0x0073B54F | `HandleTimeChosen` | Known | Event handler |
| 0x0073B585 | `HandleFrequencyChosen` | Known | Event handler |
| 0x0073B5C0 | `HandleSoundChosen` | Known | Event handler |
| 0x0073B5F7 | `HandleLabelChosen` | Known | Event handler |
| 0x0073B62E | `HandleDeleteChosen` | Known | Event handler |
| 0x0073B66A | `HandleSelect` | Known | Event handler |
| 0x0073B6A2 | `HandleSelect` | Known | Event handler |
| 0x0073B9E3 | `HandleLeaveAlarm` | Known | Event handler |
| 0x0073BA10 | `HandleLeaveAlarm` | Known | Event handler |
| 0x0073BA3F | `HandleLeaveAlarm` | Known | Event handler |
| 0x0073BA6C | `HandleLeaveAlarm` | Known | Event handler |
| 0x0073BBA6 | `HandleSelect` | Known | Event handler |
| 0x0073BBD4 | `HandleSelect` | Known | Event handler |
| 0x0073BD33 | `HandleNextDay` | Known | Event handler |
| 0x0073BD5B | `HandlePreviousDay` | Known | Event handler |
| 0x0073BF0A | `HandleSelect` | Known | Event handler |
| 0x0073BF37 | `HandleNextDay` | Known | Event handler |
| 0x0073BF5F | `HandlePreviousDay` | Known | Event handler |
| 0x0073C107 | `HandleNextDay` | Known | Event handler |
| 0x0073C12F | `HandlePreviousDay` | Known | Event handler |
| 0x0073C1F0 | `HandleSelect` | Known | Event handler |
| 0x0073C21B | `HandleNextDay` | Known | Event handler |
| 0x0073C243 | `HandlePreviousDay` | Known | Event handler |
| 0x0073C3BA | `HandleSelectLozinch` | Known | Event handler |
| 0x0073C532 | `HandleSelectLozinch` | Known | Event handler |
| 0x0073C651 | `HandleFlowNext` | Known | Event handler |
| 0x0073C67F | `HandlePlayPause` | Known | Event handler |
| 0x0073C6CE | `HandleFlowPrev` | Known | Event handler |
| 0x0073C6F9 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x0073C7ED | `HandleAlbumSelected` | Known | Event handler |
| 0x0073C988 | `HandleFlowNext` | Known | Event handler |
| 0x0073C9D6 | `HandleFlowNext` | Known | Event handler |
| 0x0073CA04 | `HandlePlayPause` | Known | Event handler |
| 0x0073CA53 | `HandleFlowPrev` | Known | Event handler |
| 0x0073CA7F | `HandleFlowPrev` | Known | Event handler |
| 0x0073CA9F | `HandleFlowWheel` | Known | Event handler |
| 0x0073CE2F | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0073D25A | `HandleArrowDown` | Known | Event handler |
| 0x0073D2C4 | `HandleArrowUp` | Known | Event handler |
| 0x0073D2E3 | `HandleWheel` | Known | Event handler |
| 0x0073D36C | `HandleSelect` | Known | Event handler |
| 0x0073D3E9 | `HandleGameHilited` | Known | Event handler |
| 0x0074084F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007425EB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00744387 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00746123 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00747EBF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00749C5B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074B9F7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074D793 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074F52F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007512CB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00753067 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00754E03 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00756B9F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075893B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075A6D7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075C473 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075E20F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075FFAB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00761D47 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00763AE3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076587F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076761B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007693B7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076B153 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076CEEF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076EC8B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00770A27 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007727C3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077455F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007762FB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00778097 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00779E33 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077BBCF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077D96B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077F707 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007814A3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078323F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00784FC0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00785C3C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007868B8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00787534 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007881B0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00788E2C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00789AA8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078A724 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078B3A0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078C01C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078CC98 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078D914 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078E590 | `HandlePlayPause` | Known | Event handler |
| 0x0078E5C6 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078E608 | `HandleAddToOTG` | Known | Event handler |
| 0x0078E7A5 | `HandlePlayPause` | Known | Event handler |
| 0x0078E7CC | `HandleSelect` | Known | Event handler |
| 0x0078E7F9 | `HandleHilite` | Known | Event handler |
| 0x0078E82C | `HandlePlayPause` | Known | Event handler |
| 0x0078E8BF | `HandlePlayPause` | Known | Event handler |
| 0x0078E8E6 | `HandleSelect` | Known | Event handler |
| 0x0078E94C | `HandleHilite` | Known | Event handler |
| 0x0078E97E | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x0078E9C8 | `HandlePlayPause` | Known | Event handler |
| 0x0078E9FE | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078EA45 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0078EA88 | `HandleAddToOTG` | Known | Event handler |
| 0x0078EAEB | `HandleStartGenius` | Known | Event handler |
| 0x0078EB27 | `HandleViewAlbum` | Known | Event handler |
| 0x0078EB62 | `HandleViewArtist` | Known | Event handler |
| 0x0078EBA3 | `HandleViewCompilation` | Known | Event handler |
| 0x0078ED43 | `HandlePlayPause` | Known | Event handler |
| 0x0078ED6A | `HandleSelect` | Known | Event handler |
| 0x0078EDD4 | `HandlePlayPause` | Known | Event handler |
| 0x0078EE0A | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078EE51 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0078EE94 | `HandleAddToOTG` | Known | Event handler |
| 0x0078EEF7 | `HandleStartGenius` | Known | Event handler |
| 0x0078EF33 | `HandleViewAlbum` | Known | Event handler |
| 0x0078EF6E | `HandleViewArtist` | Known | Event handler |
| 0x0078EFAF | `HandleViewCompilation` | Known | Event handler |
| 0x0078F14F | `HandlePlayPause` | Known | Event handler |
| 0x0078F176 | `HandleSelect` | Known | Event handler |
| 0x0078F1E0 | `HandlePlayPause` | Known | Event handler |
| 0x0078F21E | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0078F261 | `HandleAddToOTG` | Known | Event handler |
| 0x0078F2C4 | `HandleStartGenius` | Known | Event handler |
| 0x0078F300 | `HandleViewAlbum` | Known | Event handler |
| 0x0078F33B | `HandleViewArtist` | Known | Event handler |
| 0x0078F37C | `HandleViewCompilation` | Known | Event handler |
| 0x0078F50F | `HandleSelect` | Known | Event handler |
| 0x0078F574 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0078F5B8 | `HandlePlayPause` | Known | Event handler |
| 0x0078F5EE | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078F630 | `HandleAddToOTG` | Known | Event handler |
| 0x0078F88A | `HandlePlayPause` | Known | Event handler |
| 0x0078F8B1 | `HandleSelect` | Known | Event handler |
| 0x0078F8DE | `HandleHilite` | Known | Event handler |
| 0x0078F910 | `HandlePlayPause` | Known | Event handler |
| 0x0078F946 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078F988 | `HandleAddToOTG` | Known | Event handler |
| 0x0078FBE2 | `HandlePlayPause` | Known | Event handler |
| 0x0078FC09 | `HandleSelect` | Known | Event handler |
| 0x0078FC36 | `HandleHilite` | Known | Event handler |
| 0x0078FC68 | `HandlePlayPause` | Known | Event handler |
| 0x0078FC9E | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078FCE0 | `HandleAddToOTG` | Known | Event handler |
| 0x0078FFF3 | `HandlePlayPause` | Known | Event handler |
| 0x0079001A | `HandleSelect` | Known | Event handler |
| 0x0079004C | `HandlePlayPause` | Known | Event handler |
| 0x00790082 | `HandleShowContextualMenu` | Known | Event handler |
| 0x007900C4 | `HandleAddToOTG` | Known | Event handler |
| 0x0079017E | `HandlePlayPause` | Known | Event handler |
| 0x007901A5 | `HandleSelect` | Known | Event handler |
| 0x00790234 | `HandlePlayPause` | Known | Event handler |
| 0x0079026A | `HandleShowContextualMenu` | Known | Event handler |
| 0x007902AC | `HandleAddToOTG` | Known | Event handler |
| 0x0079048D | `HandlePlayPause` | Known | Event handler |
| 0x007904B4 | `HandleSelect` | Known | Event handler |
| 0x007904E4 | `HandlePlayPause` | Known | Event handler |
| 0x0079051A | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079055C | `HandleAddToOTG` | Known | Event handler |
| 0x00790609 | `HandleSelect` | Known | Event handler |
| 0x007906A2 | `HandleHilite` | Known | Event handler |
| 0x007906CE | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00790710 | `HandlePlayPause` | Known | Event handler |
| 0x00790746 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00790788 | `HandleAddToOTG` | Known | Event handler |
| 0x00790835 | `HandleSelect` | Known | Event handler |
| 0x0079089A | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007908DC | `HandlePlayPause` | Known | Event handler |
| 0x00790A80 | `HandleSelect` | Known | Event handler |
| 0x00790AAD | `HandleHilite` | Known | Event handler |
| 0x00790AD9 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00790B1C | `HandlePlayPause` | Known | Event handler |
| 0x00790BA2 | `HandleSelect` | Known | Event handler |
| 0x00790C30 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00790C74 | `HandlePlayPause` | Known | Event handler |
| 0x00790CFA | `HandleSelect` | Known | Event handler |
| 0x00790D5F | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00790DA0 | `HandlePlayPause` | Known | Event handler |
| 0x00790E26 | `HandleSelect` | Known | Event handler |
| 0x00790E8C | `HandleHilite` | Known | Event handler |
| 0x00790EB8 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00790EFC | `HandlePlayPause` | Known | Event handler |
| 0x00790F32 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00790F74 | `HandleAddToOTG` | Known | Event handler |
| 0x007911F9 | `HandlePlayPause` | Known | Event handler |
| 0x00791220 | `HandleSelect` | Known | Event handler |
| 0x00791250 | `HandlePlayPause` | Known | Event handler |
| 0x00791286 | `HandleShowContextualMenu` | Known | Event handler |
| 0x007912CD | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00791310 | `HandleAddToOTG` | Known | Event handler |
| 0x00791373 | `HandleStartGenius` | Known | Event handler |
| 0x007913AF | `HandleViewAlbum` | Known | Event handler |
| 0x007913EA | `HandleViewArtist` | Known | Event handler |
| 0x0079142B | `HandleViewCompilation` | Known | Event handler |
| 0x00791913 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x00791958 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079199B | `HandleAddToOTG` | Known | Event handler |
| 0x007919FE | `HandleStartGenius` | Known | Event handler |
| 0x00791A3A | `HandleViewAlbum` | Known | Event handler |
| 0x00791A75 | `HandleViewArtist` | Known | Event handler |
| 0x00791AB6 | `HandleViewCompilation` | Known | Event handler |
| 0x00791E8C | `HandlePlayPause` | Known | Event handler |
| 0x00791FB9 | `HandleSelect` | Known | Event handler |
| 0x00791FE5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00792028 | `HandlePlayPause` | Known | Event handler |
| 0x007920AE | `HandleSelect` | Known | Event handler |
| 0x007920DB | `HandleHilite` | Known | Event handler |
| 0x00792107 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00792148 | `HandlePlayPause` | Known | Event handler |
| 0x0079227B | `HandleSelect` | Known | Event handler |
| 0x007922A7 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00792BB9 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00793471 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00793D29 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007945E1 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00794E99 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00795751 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00796009 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007968C1 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079690A | `HandleTVOutChanged` | Known | Event handler |
| 0x00796942 | `HandleTVSignalChanged` | Known | Event handler |
| 0x0079697D | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x007969CE | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x00796A13 | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x00796A5C | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x00796A9E | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x00796AE8 | `HandlePlayPause` | Known | Event handler |
| 0x00796B1E | `HandleShowContextualMenu` | Known | Event handler |
| 0x00796B65 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00796BA8 | `HandleAddToOTG` | Known | Event handler |
| 0x00796C0B | `HandleStartGenius` | Known | Event handler |
| 0x00796C47 | `HandleViewAlbum` | Known | Event handler |
| 0x00796C82 | `HandleViewArtist` | Known | Event handler |
| 0x00796CC3 | `HandleViewCompilation` | Known | Event handler |
| 0x00796EFF | `HandlePlayPause` | Known | Event handler |
| 0x00796F26 | `HandleSelect` | Known | Event handler |
| 0x00796F58 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x00796F93 | `HandleSaveGeniusPlaylist` | Known | Event handler |
| 0x00797034 | `HandlePlayPause` | Known | Event handler |
| 0x0079706A | `HandleShowContextualMenu` | Known | Event handler |
| 0x007970B1 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007970F4 | `HandleAddToOTG` | Known | Event handler |
| 0x00797157 | `HandleStartGenius` | Known | Event handler |
| 0x00797193 | `HandleViewAlbum` | Known | Event handler |
| 0x007971CE | `HandleViewArtist` | Known | Event handler |
| 0x0079720F | `HandleViewCompilation` | Known | Event handler |
| 0x0079727D | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x007976A5 | `HandlePlayPause` | Known | Event handler |
| 0x007976CC | `HandleSelect` | Known | Event handler |
| 0x007976FE | `HandleRefreshPlaylist` | Known | Event handler |
| 0x00797735 | `HandleSelect` | Known | Event handler |
| 0x00797765 | `HandleSelect` | Known | Event handler |
| 0x0079779D | `HandleMenuLongpress` | Known | Event handler |
| 0x007977CB | `HandleMenuKey` | Known | Event handler |
| 0x00797851 | `HandlePlayPause` | Known | Event handler |
| 0x007978DB | `HandlePushContextualMenu` | Known | Event handler |
| 0x00797910 | `HandleSelect` | Known | Event handler |
| 0x0079794B | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079798E | `HandleAddToOTG` | Known | Event handler |
| 0x007979CD | `HandleAudiobookFaster` | Known | Event handler |
| 0x00797A13 | `HandleAudiobookNormal` | Known | Event handler |
| 0x00797A59 | `HandleAudiobookSlower` | Known | Event handler |
| 0x00797AC3 | `HandleStartGenius` | Known | Event handler |
| 0x00797AFF | `HandleViewAlbum` | Known | Event handler |
| 0x00797B3A | `HandleViewArtist` | Known | Event handler |
| 0x00797B7B | `HandleViewCompilation` | Known | Event handler |
| 0x007985BD | `HandleStartGenius` | Known | Event handler |
| 0x007986D0 | `HandlePlayPause` | Known | Event handler |
| 0x00798745 | `HandleWheelProgress` | Known | Event handler |
| 0x00798781 | `HandleMenuLongpress` | Known | Event handler |
| 0x007987AF | `HandleMenuKey` | Known | Event handler |
| 0x00798835 | `HandlePlayPause` | Known | Event handler |
| 0x007988BF | `HandlePushContextualMenu` | Known | Event handler |
| 0x007988F4 | `HandleSelectProgress` | Known | Event handler |
| 0x00798937 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079897A | `HandleAddToOTG` | Known | Event handler |
| 0x007989B9 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007989FF | `HandleAudiobookNormal` | Known | Event handler |
| 0x00798A45 | `HandleAudiobookSlower` | Known | Event handler |
| 0x00798AAF | `HandleStartGenius` | Known | Event handler |
| 0x00798AEB | `HandleViewAlbum` | Known | Event handler |
| 0x00798B26 | `HandleViewArtist` | Known | Event handler |
| 0x00798B67 | `HandleViewCompilation` | Known | Event handler |
| 0x007995A9 | `HandleStartGenius` | Known | Event handler |
| 0x007996BC | `HandlePlayPause` | Known | Event handler |
| 0x00799731 | `HandleWheelProgress` | Known | Event handler |
| 0x0079976D | `HandleMenuLongpress` | Known | Event handler |
| 0x0079979B | `HandleMenuKey` | Known | Event handler |
| 0x00799821 | `HandlePlayPause` | Known | Event handler |
| 0x007998AB | `HandlePushContextualMenu` | Known | Event handler |
| 0x007998E0 | `HandleSelectVolume` | Known | Event handler |
| 0x00799921 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00799964 | `HandleAddToOTG` | Known | Event handler |
| 0x007999A3 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007999E9 | `HandleAudiobookNormal` | Known | Event handler |
| 0x00799A2F | `HandleAudiobookSlower` | Known | Event handler |
| 0x00799A99 | `HandleStartGenius` | Known | Event handler |
| 0x00799AD5 | `HandleViewAlbum` | Known | Event handler |
| 0x00799B10 | `HandleViewArtist` | Known | Event handler |
| 0x00799B51 | `HandleViewCompilation` | Known | Event handler |
| 0x0079A593 | `HandleStartGenius` | Known | Event handler |
| 0x0079A6A6 | `HandlePlayPause` | Known | Event handler |
| 0x0079A71B | `HandleWheelVolume` | Known | Event handler |
| 0x0079A755 | `HandleMenuLongpress` | Known | Event handler |
| 0x0079A783 | `HandleMenuKey` | Known | Event handler |
| 0x0079A809 | `HandlePlayPause` | Known | Event handler |
| 0x0079A893 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079A8C8 | `HandleSelectRating` | Known | Event handler |
| 0x0079A909 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079A94C | `HandleAddToOTG` | Known | Event handler |
| 0x0079A98B | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079A9D1 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079AA17 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079AA81 | `HandleStartGenius` | Known | Event handler |
| 0x0079AABD | `HandleViewAlbum` | Known | Event handler |
| 0x0079AAF8 | `HandleViewArtist` | Known | Event handler |
| 0x0079AB39 | `HandleViewCompilation` | Known | Event handler |
| 0x0079B57B | `HandleStartGenius` | Known | Event handler |
| 0x0079B68E | `HandlePlayPause` | Known | Event handler |
| 0x0079B703 | `HandleWheelRating` | Known | Event handler |
| 0x0079B73D | `HandleMenuLongpress` | Known | Event handler |
| 0x0079B76B | `HandleMenuKey` | Known | Event handler |
| 0x0079B7E3 | `HandlePlayPause` | Known | Event handler |
| 0x0079B864 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079B899 | `HandleSelectScrub` | Known | Event handler |
| 0x0079B8D9 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079B91C | `HandleAddToOTG` | Known | Event handler |
| 0x0079B95B | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079B9A1 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079B9E7 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079BA51 | `HandleStartGenius` | Known | Event handler |
| 0x0079BA8D | `HandleViewAlbum` | Known | Event handler |
| 0x0079BAC8 | `HandleViewArtist` | Known | Event handler |
| 0x0079BB09 | `HandleViewCompilation` | Known | Event handler |
| 0x0079C54B | `HandleStartGenius` | Known | Event handler |
| 0x0079C650 | `HandlePlayPause` | Known | Event handler |
| 0x0079C6BC | `HandleWheelScrub` | Known | Event handler |
| 0x0079C6F5 | `HandleMenuLongpress` | Known | Event handler |
| 0x0079C723 | `HandleMenuKey` | Known | Event handler |
| 0x0079C7A9 | `HandlePlayPause` | Known | Event handler |
| 0x0079C833 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079C868 | `HandleSelectGenius` | Known | Event handler |
| 0x0079C8A9 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079C8EC | `HandleAddToOTG` | Known | Event handler |
| 0x0079C92B | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079C971 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079C9B7 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079CA21 | `HandleStartGenius` | Known | Event handler |
| 0x0079CA5D | `HandleViewAlbum` | Known | Event handler |
| 0x0079CA98 | `HandleViewArtist` | Known | Event handler |
| 0x0079CAD9 | `HandleViewCompilation` | Known | Event handler |
| 0x0079D51B | `HandleStartGenius` | Known | Event handler |
| 0x0079D62E | `HandlePlayPause` | Known | Event handler |
| 0x0079D6A3 | `HandleWheelGenius` | Known | Event handler |
| 0x0079D6DD | `HandleMenuLongpress` | Known | Event handler |
| 0x0079D70B | `HandleMenuKey` | Known | Event handler |
| 0x0079D768 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0079D7A0 | `HandlePlayPause` | Known | Event handler |
| 0x0079D7FA | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0079D839 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079D86E | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x0079D8B6 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079D8F9 | `HandleAddToOTG` | Known | Event handler |
| 0x0079D938 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079D97E | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079D9C4 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079DA2E | `HandleStartGenius` | Known | Event handler |
| 0x0079DA6A | `HandleViewAlbum` | Known | Event handler |
| 0x0079DAA5 | `HandleViewArtist` | Known | Event handler |
| 0x0079DAE6 | `HandleViewCompilation` | Known | Event handler |
| 0x0079E528 | `HandleStartGenius` | Known | Event handler |
| 0x0079E63B | `HandlePlayPause` | Known | Event handler |
| 0x0079E6B0 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0079E6F1 | `HandleMenuLongpress` | Known | Event handler |
| 0x0079E71F | `HandleMenuKey` | Known | Event handler |
| 0x0079E7A5 | `HandlePlayPause` | Known | Event handler |
| 0x0079E82F | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079E864 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x0079E8A8 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079E8EB | `HandleAddToOTG` | Known | Event handler |
| 0x0079E92A | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079E970 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079E9B6 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079EA20 | `HandleStartGenius` | Known | Event handler |
| 0x0079EA5C | `HandleViewAlbum` | Known | Event handler |
| 0x0079EA97 | `HandleViewArtist` | Known | Event handler |
| 0x0079EAD8 | `HandleViewCompilation` | Known | Event handler |
| 0x0079F51A | `HandleStartGenius` | Known | Event handler |
| 0x0079F62D | `HandlePlayPause` | Known | Event handler |
| 0x0079F6CD | `HandleMenuLongpress` | Known | Event handler |
| 0x0079F6FB | `HandleMenuKey` | Known | Event handler |
| 0x0079F781 | `HandlePlayPause` | Known | Event handler |
| 0x0079F80B | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079F840 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x0079F884 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079F8C7 | `HandleAddToOTG` | Known | Event handler |
| 0x0079F906 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079F94C | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079F992 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079F9FC | `HandleStartGenius` | Known | Event handler |
| 0x0079FA38 | `HandleViewAlbum` | Known | Event handler |
| 0x0079FA73 | `HandleViewArtist` | Known | Event handler |
| 0x0079FAB4 | `HandleViewCompilation` | Known | Event handler |
| 0x007A04F6 | `HandleStartGenius` | Known | Event handler |
| 0x007A0609 | `HandlePlayPause` | Known | Event handler |
| 0x007A06A9 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A06D7 | `HandleMenuKey` | Known | Event handler |
| 0x007A075D | `HandlePlayPause` | Known | Event handler |
| 0x007A07E7 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A081C | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007A0860 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A08A3 | `HandleAddToOTG` | Known | Event handler |
| 0x007A08E2 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A0928 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A096E | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A09D8 | `HandleStartGenius` | Known | Event handler |
| 0x007A0A14 | `HandleViewAlbum` | Known | Event handler |
| 0x007A0A4F | `HandleViewArtist` | Known | Event handler |
| 0x007A0A90 | `HandleViewCompilation` | Known | Event handler |
| 0x007A14D2 | `HandleStartGenius` | Known | Event handler |
| 0x007A15E5 | `HandlePlayPause` | Known | Event handler |
| 0x007A1685 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A16B3 | `HandleMenuKey` | Known | Event handler |
| 0x007A1739 | `HandlePlayPause` | Known | Event handler |
| 0x007A17C3 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A17F8 | `HandleSelectChapterArt` | Known | Event handler |
| 0x007A183D | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A1880 | `HandleAddToOTG` | Known | Event handler |
| 0x007A18BF | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A1905 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A194B | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A19B5 | `HandleStartGenius` | Known | Event handler |
| 0x007A19F1 | `HandleViewAlbum` | Known | Event handler |
| 0x007A1A2C | `HandleViewArtist` | Known | Event handler |
| 0x007A1A6D | `HandleViewCompilation` | Known | Event handler |
| 0x007A24AF | `HandleStartGenius` | Known | Event handler |
| 0x007A25C2 | `HandlePlayPause` | Known | Event handler |
| 0x007A2637 | `HandleWheelVolume` | Known | Event handler |
| 0x007A2671 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A269F | `HandleMenuKey` | Known | Event handler |
| 0x007A272E | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007A27CF | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A2804 | `HandleSelect` | Known | Event handler |
| 0x007A283F | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A2882 | `HandleAddToOTG` | Known | Event handler |
| 0x007A28C1 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A2907 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A294D | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A29B7 | `HandleStartGenius` | Known | Event handler |
| 0x007A29F3 | `HandleViewAlbum` | Known | Event handler |
| 0x007A2A2E | `HandleViewArtist` | Known | Event handler |
| 0x007A2A6F | `HandleViewCompilation` | Known | Event handler |
| 0x007A34B1 | `HandleStartGenius` | Known | Event handler |
| 0x007A35CD | `HandlePlayPause` | Known | Event handler |
| 0x007A364B | `HandleWheel` | Known | Event handler |
| 0x007A3681 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A36AF | `HandleMenuKey` | Known | Event handler |
| 0x007A373E | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007A37DF | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A3814 | `HandleSelect` | Known | Event handler |
| 0x007A384F | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A3892 | `HandleAddToOTG` | Known | Event handler |
| 0x007A38D1 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A3917 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A395D | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A39C7 | `HandleStartGenius` | Known | Event handler |
| 0x007A3A03 | `HandleViewAlbum` | Known | Event handler |
| 0x007A3A3E | `HandleViewArtist` | Known | Event handler |
| 0x007A3A7F | `HandleViewCompilation` | Known | Event handler |
| 0x007A44C1 | `HandleStartGenius` | Known | Event handler |
| 0x007A45DD | `HandlePlayPause` | Known | Event handler |
| 0x007A465B | `HandleWheel` | Known | Event handler |
| 0x007A4691 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A46BF | `HandleMenuKey` | Known | Event handler |
| 0x007A4745 | `HandlePlayPause` | Known | Event handler |
| 0x007A47CF | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A4804 | `HandleSelect` | Known | Event handler |
| 0x007A483F | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A4882 | `HandleAddToOTG` | Known | Event handler |
| 0x007A48C1 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A4907 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A494D | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A49B7 | `HandleStartGenius` | Known | Event handler |
| 0x007A49F3 | `HandleViewAlbum` | Known | Event handler |
| 0x007A4A2E | `HandleViewArtist` | Known | Event handler |
| 0x007A4A6F | `HandleViewCompilation` | Known | Event handler |
| 0x007A54B1 | `HandleStartGenius` | Known | Event handler |
| 0x007A55C4 | `HandlePlayPause` | Known | Event handler |
| 0x007A5639 | `HandleWheel` | Known | Event handler |
| 0x007A566D | `HandleMenuLongpress` | Known | Event handler |
| 0x007A569B | `HandleMenuKey` | Known | Event handler |
| 0x007A5721 | `HandlePlayPause` | Known | Event handler |
| 0x007A57AB | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A57E0 | `HandleSelectProgress` | Known | Event handler |
| 0x007A5823 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A5866 | `HandleAddToOTG` | Known | Event handler |
| 0x007A58A5 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A58EB | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A5931 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A599B | `HandleStartGenius` | Known | Event handler |
| 0x007A59D7 | `HandleViewAlbum` | Known | Event handler |
| 0x007A5A12 | `HandleViewArtist` | Known | Event handler |
| 0x007A5A53 | `HandleViewCompilation` | Known | Event handler |
| 0x007A6495 | `HandleStartGenius` | Known | Event handler |
| 0x007A65A8 | `HandlePlayPause` | Known | Event handler |
| 0x007A661D | `HandleWheelProgress` | Known | Event handler |
| 0x007A6659 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A6687 | `HandleMenuKey` | Known | Event handler |
| 0x007A66FF | `HandlePlayPause` | Known | Event handler |
| 0x007A6780 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A67B5 | `HandleSelectScrub` | Known | Event handler |
| 0x007A67F5 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A6838 | `HandleAddToOTG` | Known | Event handler |
| 0x007A6877 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A68BD | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A6903 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A696D | `HandleStartGenius` | Known | Event handler |
| 0x007A69A9 | `HandleViewAlbum` | Known | Event handler |
| 0x007A69E4 | `HandleViewArtist` | Known | Event handler |
| 0x007A6A25 | `HandleViewCompilation` | Known | Event handler |
| 0x007A7467 | `HandleStartGenius` | Known | Event handler |
| 0x007A756C | `HandlePlayPause` | Known | Event handler |
| 0x007A75D8 | `HandleWheelScrub` | Known | Event handler |
| 0x007A7611 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A763F | `HandleMenuKey` | Known | Event handler |
| 0x007A76C5 | `HandlePlayPause` | Known | Event handler |
| 0x007A774F | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A77BE | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A7801 | `HandleAddToOTG` | Known | Event handler |
| 0x007A7840 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A7886 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A78CC | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A7936 | `HandleStartGenius` | Known | Event handler |
| 0x007A7972 | `HandleViewAlbum` | Known | Event handler |
| 0x007A79AD | `HandleViewArtist` | Known | Event handler |
| 0x007A79EE | `HandleViewCompilation` | Known | Event handler |
| 0x007A8430 | `HandleStartGenius` | Known | Event handler |
| 0x007A8543 | `HandlePlayPause` | Known | Event handler |
| 0x007A85B8 | `HandleWheelVolume` | Known | Event handler |
| 0x007A85F5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A8623 | `HandleMenuKey` | Known | Event handler |
| 0x007A86A9 | `HandlePlayPause` | Known | Event handler |
| 0x007A8733 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A87A2 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A87E5 | `HandleAddToOTG` | Known | Event handler |
| 0x007A8824 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A886A | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A88B0 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A891A | `HandleStartGenius` | Known | Event handler |
| 0x007A8956 | `HandleViewAlbum` | Known | Event handler |
| 0x007A8991 | `HandleViewArtist` | Known | Event handler |
| 0x007A89D2 | `HandleViewCompilation` | Known | Event handler |
| 0x007A9414 | `HandleStartGenius` | Known | Event handler |
| 0x007A9527 | `HandlePlayPause` | Known | Event handler |
| 0x007A959C | `HandleWheelBrightness` | Known | Event handler |
| 0x007A96BF | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A96F4 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007A973C | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A977F | `HandleAddToOTG` | Known | Event handler |
| 0x007A97BE | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A9804 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A984A | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A98B4 | `HandleStartGenius` | Known | Event handler |
| 0x007A98F0 | `HandleViewAlbum` | Known | Event handler |
| 0x007A992B | `HandleViewArtist` | Known | Event handler |
| 0x007A996C | `HandleViewCompilation` | Known | Event handler |
| 0x007AA3AE | `HandleStartGenius` | Known | Event handler |
| 0x007AA4FA | `HandleWheel` | Known | Event handler |
| 0x007AA531 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AA55F | `HandleMenuKey` | Known | Event handler |
| 0x007AA5E5 | `HandlePlayPause` | Known | Event handler |
| 0x007AA665 | `HandleSelect` | Known | Event handler |
| 0x007AAB07 | `HandlePlayPause` | Known | Event handler |
| 0x007AAB95 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AABC3 | `HandleMenuKey` | Known | Event handler |
| 0x007AAC49 | `HandlePlayPause` | Known | Event handler |
| 0x007AACC9 | `HandleSelectProgress` | Known | Event handler |
| 0x007AB173 | `HandlePlayPause` | Known | Event handler |
| 0x007AB1E8 | `HandleWheelProgress` | Known | Event handler |
| 0x007AB225 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AB253 | `HandleMenuKey` | Known | Event handler |
| 0x007AB2D9 | `HandlePlayPause` | Known | Event handler |
| 0x007AB359 | `HandleSelectProgress` | Known | Event handler |
| 0x007AB803 | `HandlePlayPause` | Known | Event handler |
| 0x007AB878 | `HandleWheelProgress` | Known | Event handler |
| 0x007AB8B5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AB8E3 | `HandleMenuKey` | Known | Event handler |
| 0x007AB969 | `HandlePlayPause` | Known | Event handler |
| 0x007AB9E9 | `HandleSelectProgress` | Known | Event handler |
| 0x007ABE1F | `HandlePlayPause` | Known | Event handler |
| 0x007ABE94 | `HandleWheelProgress` | Known | Event handler |
| 0x007ABED1 | `HandleMenuLongpress` | Known | Event handler |
| 0x007ABEFF | `HandleMenuKey` | Known | Event handler |
| 0x007ABF6C | `HandlePlayPause` | Known | Event handler |
| 0x007ABFD8 | `HandleSelectScrub` | Known | Event handler |
| 0x007AC3F2 | `HandlePlayPause` | Known | Event handler |
| 0x007AC453 | `HandleWheelScrub` | Known | Event handler |
| 0x007AC48D | `HandleMenuLongpress` | Known | Event handler |
| 0x007AC4BB | `HandleMenuKey` | Known | Event handler |
| 0x007AC541 | `HandlePlayPause` | Known | Event handler |
| 0x007AC5C1 | `HandleSelectVolume` | Known | Event handler |
| 0x007AC9F5 | `HandlePlayPause` | Known | Event handler |
| 0x007ACA6A | `HandleWheelVolume` | Known | Event handler |
| 0x007ACB7D | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007AD01C | `HandleSelect` | Known | Event handler |
| 0x007AD049 | `HandleSelect` | Known | Event handler |
| 0x007AD079 | `HandleSelect` | Known | Event handler |
| 0x007AD0A9 | `HandleSelect` | Known | Event handler |
| 0x007AD0D9 | `HandleSelect` | Known | Event handler |
| 0x007AD109 | `HandleSelect` | Known | Event handler |
| 0x007AD139 | `HandleSelect` | Known | Event handler |
| 0x007AD169 | `HandleSelect` | Known | Event handler |
| 0x007AD199 | `HandleSelect` | Known | Event handler |
| 0x007AD209 | `HandleSelect` | Known | Event handler |
| 0x007AD239 | `HandleSelect` | Known | Event handler |
| 0x007AD2B1 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AD2E4 | `HandleNotesPop` | Known | Event handler |
| 0x007AD361 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AD394 | `HandleNotesPop` | Known | Event handler |
| 0x007AD850 | `HandleNotesSelected` | Known | Event handler |
| 0x007AD88D | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AD8C0 | `HandleNotesPop` | Known | Event handler |
| 0x007ADD7C | `HandleNotesSelected` | Known | Event handler |
| 0x007ADDB9 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007ADDEC | `HandleNotesPop` | Known | Event handler |
| 0x007ADE17 | `HandleNotesSelected` | Known | Event handler |
| 0x007AE2E9 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AE31C | `HandleNotesPop` | Known | Event handler |
| 0x007AE347 | `HandleNotesSelected` | Known | Event handler |
| 0x007AE819 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AE84C | `HandleNotesPop` | Known | Event handler |
| 0x007AE8C9 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AE8FC | `HandleNotesPop` | Known | Event handler |
| 0x007AE979 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AE9AC | `HandleNotesPop` | Known | Event handler |
| 0x007AEA24 | `HandlePlayPause` | Known | Event handler |
| 0x007AEA4D | `HandlePlayPause` | Known | Event handler |
| 0x007AEA7B | `HandlePlayPause` | Known | Event handler |
| 0x007AEAB0 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007AEB30 | `HandleHiliteAlbum` | Known | Event handler |
| 0x007AEBD9 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007AEC60 | `HandleHiliteAlbum` | Known | Event handler |
| 0x007AEF24 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x007AEF80 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x007AF137 | `HandleSelect` | Known | Event handler |
| 0x007AF2BB | `HandleSelect` | Known | Event handler |
| 0x007AF2F5 | `HandleImageLast` | Known | Event handler |
| 0x007AF31F | `HandleImageNext` | Known | Event handler |
| 0x007AF34E | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AF388 | `HandleImageFirst` | Known | Event handler |
| 0x007AF3B3 | `HandleImagePrev` | Known | Event handler |
| 0x007AF3DF | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AF40E | `HandleImageNext` | Known | Event handler |
| 0x007AF437 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AF46B | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AF49A | `HandleImagePrev` | Known | Event handler |
| 0x007AF4BB | `HandleImageWheel` | Known | Event handler |
| 0x007AF559 | `HandleImageNext` | Known | Event handler |
| 0x007AF588 | `HandlePlayPause` | Known | Event handler |
| 0x007AF5D7 | `HandleImagePrev` | Known | Event handler |
| 0x007AF603 | `HandleSelect` | Known | Event handler |
| 0x007AF8D3 | `HandleImageNext` | Known | Event handler |
| 0x007AF8FD | `HandlePause` | Known | Event handler |
| 0x007AF922 | `HandlePlay` | Known | Event handler |
| 0x007AF94B | `HandlePlayPause` | Known | Event handler |
| 0x007AF974 | `HandleImagePrev` | Known | Event handler |
| 0x007AF9D7 | `HandleMikeyCenter` | Known | Event handler |
| 0x007AF9FA | `HandleWheel` | Known | Event handler |
| 0x007AFA95 | `HandleImageNext` | Known | Event handler |
| 0x007AFAC4 | `HandlePlayPause` | Known | Event handler |
| 0x007AFB13 | `HandleImagePrev` | Known | Event handler |
| 0x007AFB3F | `HandleSelect` | Known | Event handler |
| 0x007AFE0F | `HandleImageNext` | Known | Event handler |
| 0x007AFE39 | `HandlePause` | Known | Event handler |
| 0x007AFE5E | `HandlePlay` | Known | Event handler |
| 0x007AFE87 | `HandlePlayPause` | Known | Event handler |
| 0x007AFEB0 | `HandleImagePrev` | Known | Event handler |
| 0x007AFF13 | `HandleMikeyCenter` | Known | Event handler |
| 0x007AFF36 | `HandleWheel` | Known | Event handler |
| 0x007AFFD1 | `HandleImageNext` | Known | Event handler |
| 0x007B0000 | `HandlePlayPause` | Known | Event handler |
| 0x007B004F | `HandleImagePrev` | Known | Event handler |
| 0x007B007B | `HandleSelect` | Known | Event handler |
| 0x007B034B | `HandleImageNext` | Known | Event handler |
| 0x007B0375 | `HandlePause` | Known | Event handler |
| 0x007B039A | `HandlePlay` | Known | Event handler |
| 0x007B03C3 | `HandlePlayPause` | Known | Event handler |
| 0x007B03EC | `HandleImagePrev` | Known | Event handler |
| 0x007B044F | `HandleMikeyCenter` | Known | Event handler |
| 0x007B0472 | `HandleWheel` | Known | Event handler |
| 0x007B050D | `HandleImageNext` | Known | Event handler |
| 0x007B053C | `HandlePlayPause` | Known | Event handler |
| 0x007B058B | `HandleImagePrev` | Known | Event handler |
| 0x007B05B7 | `HandleSelect` | Known | Event handler |
| 0x007B0887 | `HandleImageNext` | Known | Event handler |
| 0x007B08B1 | `HandlePause` | Known | Event handler |
| 0x007B08D6 | `HandlePlay` | Known | Event handler |
| 0x007B08FF | `HandlePlayPause` | Known | Event handler |
| 0x007B0928 | `HandleImagePrev` | Known | Event handler |
| 0x007B098B | `HandleMikeyCenter` | Known | Event handler |
| 0x007B09AE | `HandleWheel` | Known | Event handler |
| 0x007B0A49 | `HandleImageNext` | Known | Event handler |
| 0x007B0A78 | `HandlePlayPause` | Known | Event handler |
| 0x007B0AC7 | `HandleImagePrev` | Known | Event handler |
| 0x007B0AF3 | `HandleSelect` | Known | Event handler |
| 0x007B0DC3 | `HandleImageNext` | Known | Event handler |
| 0x007B0DED | `HandlePause` | Known | Event handler |
| 0x007B0E12 | `HandlePlay` | Known | Event handler |
| 0x007B0E3B | `HandlePlayPause` | Known | Event handler |
| 0x007B0E64 | `HandleImagePrev` | Known | Event handler |
| 0x007B0EC7 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B0EEA | `HandleWheel` | Known | Event handler |
| 0x007B0F85 | `HandleImageNext` | Known | Event handler |
| 0x007B0FB4 | `HandlePlayPause` | Known | Event handler |
| 0x007B1003 | `HandleImagePrev` | Known | Event handler |
| 0x007B102F | `HandleSelect` | Known | Event handler |
| 0x007B12FF | `HandleImageNext` | Known | Event handler |
| 0x007B1329 | `HandlePause` | Known | Event handler |
| 0x007B134E | `HandlePlay` | Known | Event handler |
| 0x007B1377 | `HandlePlayPause` | Known | Event handler |
| 0x007B13A0 | `HandleImagePrev` | Known | Event handler |
| 0x007B1403 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B1426 | `HandleWheel` | Known | Event handler |
| 0x007B14C1 | `HandleImageNext` | Known | Event handler |
| 0x007B14F0 | `HandlePlayPause` | Known | Event handler |
| 0x007B153F | `HandleImagePrev` | Known | Event handler |
| 0x007B156B | `HandleSelect` | Known | Event handler |
| 0x007B17B6 | `HandleImageNext` | Known | Event handler |
| 0x007B17E0 | `HandlePause` | Known | Event handler |
| 0x007B1805 | `HandlePlay` | Known | Event handler |
| 0x007B182E | `HandlePlayPause` | Known | Event handler |
| 0x007B1857 | `HandleImagePrev` | Known | Event handler |
| 0x007B18CA | `HandleMikeyCenter` | Known | Event handler |
| 0x007B18ED | `HandleWheel` | Known | Event handler |
| 0x007B1985 | `HandleImageNext` | Known | Event handler |
| 0x007B19B4 | `HandlePlayPause` | Known | Event handler |
| 0x007B1A03 | `HandleImagePrev` | Known | Event handler |
| 0x007B1A2F | `HandleSelect` | Known | Event handler |
| 0x007B1C7A | `HandleImageNext` | Known | Event handler |
| 0x007B1CA4 | `HandlePause` | Known | Event handler |
| 0x007B1CC9 | `HandlePlay` | Known | Event handler |
| 0x007B1CF2 | `HandlePlayPause` | Known | Event handler |
| 0x007B1D1B | `HandleImagePrev` | Known | Event handler |
| 0x007B1D8E | `HandleMikeyCenter` | Known | Event handler |
| 0x007B1DB1 | `HandleWheel` | Known | Event handler |
| 0x007B1E49 | `HandleImageNext` | Known | Event handler |
| 0x007B1E78 | `HandlePlayPause` | Known | Event handler |
| 0x007B1EC7 | `HandleImagePrev` | Known | Event handler |
| 0x007B1EF3 | `HandleSelect` | Known | Event handler |
| 0x007B213E | `HandleImageNext` | Known | Event handler |
| 0x007B2168 | `HandlePause` | Known | Event handler |
| 0x007B218D | `HandlePlay` | Known | Event handler |
| 0x007B21B6 | `HandlePlayPause` | Known | Event handler |
| 0x007B21DF | `HandleImagePrev` | Known | Event handler |
| 0x007B2252 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B2275 | `HandleWheel` | Known | Event handler |
| 0x007B230D | `HandleImageNext` | Known | Event handler |
| 0x007B233C | `HandlePlayPause` | Known | Event handler |
| 0x007B238B | `HandleImagePrev` | Known | Event handler |
| 0x007B23B7 | `HandleSelect` | Known | Event handler |
| 0x007B2602 | `HandleImageNext` | Known | Event handler |
| 0x007B262C | `HandlePause` | Known | Event handler |
| 0x007B2651 | `HandlePlay` | Known | Event handler |
| 0x007B267A | `HandlePlayPause` | Known | Event handler |
| 0x007B26A3 | `HandleImagePrev` | Known | Event handler |
| 0x007B2716 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B2739 | `HandleWheel` | Known | Event handler |
| 0x007B27D1 | `HandleImageNext` | Known | Event handler |
| 0x007B2800 | `HandlePlayPause` | Known | Event handler |
| 0x007B284F | `HandleImagePrev` | Known | Event handler |
| 0x007B287B | `HandleSelect` | Known | Event handler |
| 0x007B2AC6 | `HandleImageNext` | Known | Event handler |
| 0x007B2AF0 | `HandlePause` | Known | Event handler |
| 0x007B2B15 | `HandlePlay` | Known | Event handler |
| 0x007B2B3E | `HandlePlayPause` | Known | Event handler |
| 0x007B2B67 | `HandleImagePrev` | Known | Event handler |
| 0x007B2BDA | `HandleMikeyCenter` | Known | Event handler |
| 0x007B2BFD | `HandleWheel` | Known | Event handler |
| 0x007B2C29 | `HandleSelect` | Known | Event handler |
| 0x007B2C59 | `HandleSelect` | Known | Event handler |
| 0x007B2D7C | `HandleTuning` | Known | Event handler |
| 0x007B2F3C | `HandleVolumeChange` | Known | Event handler |
| 0x007B2FA3 | `HandleVolumeChange` | Known | Event handler |
| 0x007B3008 | `HandleVolumeChange` | Known | Event handler |
| 0x007B3154 | `HandleVolumeWheel` | Known | Event handler |
| 0x007B32AF | `HandleTuningSelect` | Known | Event handler |
| 0x007B3475 | `HandleVolumeChange` | Known | Event handler |
| 0x007B34DC | `HandleVolumeChange` | Known | Event handler |
| 0x007B3541 | `HandleVolumeChange` | Known | Event handler |
| 0x007B368D | `HandleFrequencyChange` | Known | Event handler |
| 0x007B37EB | `HandleTuningSelect` | Known | Event handler |
| 0x007B39B1 | `HandleVolumeChange` | Known | Event handler |
| 0x007B3A18 | `HandleVolumeChange` | Known | Event handler |
| 0x007B3A7D | `HandleVolumeChange` | Known | Event handler |
| 0x007B3BC9 | `HandleFrequencyChange` | Known | Event handler |
| 0x007B3CF4 | `HandleTimerDone` | Known | Event handler |
| 0x007B3EED | `HandleVolumeChange` | Known | Event handler |
| 0x007B3F1F | `HandleVolumeChange` | Known | Event handler |
| 0x007B3F4F | `HandleVolumeChange` | Known | Event handler |
| 0x007B4066 | `HandleVolumeWheel` | Known | Event handler |
| 0x007B48B7 | `HandleExitUnsupported` | Known | Event handler |
| 0x007B48E9 | `HandleExitUnsupported` | Known | Event handler |
| 0x007B991D | `HandleSelectKey` | Known | Event handler |
| 0x007B9952 | `HandleWheel` | Known | Event handler |
| 0x007B9AA0 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x007B9AF3 | `HandleSelectKey` | Known | Event handler |
| 0x007B9B1B | `HandleSelectKey` | Known | Event handler |
| 0x007B9B4B | `HandleExit` | Known | Event handler |
| 0x007B9B75 | `HandleStartStop` | Known | Event handler |
| 0x007B9BDB | `HandleStartStop` | Known | Event handler |
| 0x007B9CF3 | `HandleExit` | Known | Event handler |
| 0x007B9D1D | `HandleStartStop` | Known | Event handler |
| 0x007B9D49 | `HandleLap` | Known | Event handler |
| 0x007B9E4D | `HandleSelectLozinch` | Known | Event handler |
| 0x007BA06A | `HandleSelect` | Known | Event handler |
| 0x007BA0F6 | `HandleSelect` | Known | Event handler |
| 0x007BA184 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x007BA482 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x007BA56D | `HandleFinishRecording` | Known | Event handler |
| 0x007BA5BE | `HandlePlayPause` | Known | Event handler |
| 0x007BA64C | `HandlePlayPause` | Known | Event handler |
| 0x007BA6DD | `HandleDeleteAllSelect` | Known | Event handler |
| 0x007BA715 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x007BA751 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x007BA794 | `HandlePlayPause` | Known | Event handler |
| 0x007BA7CA | `HandleAddToOTG` | Known | Event handler |
| 0x007BAA1F | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007BAC7B | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007D7B3E | `HandleSelectClock` | Known | Event handler |
| 0x007D7B77 | `HandleHilited` | Known | Event handler |
| 0x007D7BA9 | `HandleWheel` | Known | Event handler |
| 0x007D7BF0 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x007D7C75 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x007D7E81 | `HandleImageLast` | Known | Event handler |
| 0x007D7EAB | `HandleScreenNext` | Known | Event handler |
| 0x007D7EDB | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D7F15 | `HandleImageFirst` | Known | Event handler |
| 0x007D7F40 | `HandleScreenPrev` | Known | Event handler |
| 0x007D7F6D | `HandleBrowseLarge` | Known | Event handler |
| 0x007D7FED | `HandleImageNext` | Known | Event handler |
| 0x007D8016 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D804A | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D8079 | `HandleImagePrev` | Known | Event handler |
| 0x007D80A7 | `HandleBrowseSmall` | Known | Event handler |

---

## 6. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F5AF4 | `GotoNowPlaying` | Known | Navigation |
| 0x000F5B6C | `GotoMainMenu` | Known | Navigation |
| 0x0010E6E8 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x0010E700 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x0010E878 | `GotoScreen_AddressBook` | Known | Navigation |
| 0x0011A87C | `GotoNowPlaying` | Known | Navigation |
| 0x0011A890 | `GotoAlbums` | Known | Navigation |
| 0x0011A89C | `GotoSongs` | Known | Navigation |
| 0x0012878C | `GotoScreen_EnterPassKey` | Known | Navigation |
| 0x001287A4 | `GotoScreen_LockediPod` | Known | Navigation |
| 0x001291A8 | `GotoScreen_MainMenu` | Known | Navigation |
| 0x0013F2D0 | `GotoMainMenu` | Known | Navigation |
| 0x001C29F8 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001C4DC8 | `GotoErrorLayout` | Known | Navigation |
| 0x001CDDA0 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001CE464 | `GotoPlayDeleteMenu` | Known | Navigation |
| 0x001CE4E8 | `GotoNowPlaying` | Known | Navigation |
| 0x001E9430 | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation |
| 0x001F4E2C | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x001F4F24 | `GotoScreen_SettingsMenu` | Known | Navigation |
| 0x001FCA4C | `GotoDefaultLayout` | Known | Navigation |
| 0x001FCAD0 | `GotoVolumeLayout` | Known | Navigation |
| 0x001FCC08 | `GotoProgressLayout` | Known | Navigation |
| 0x001FCF24 | `GotoDefault` | Known | Navigation |
| 0x001FD258 | `GotoProgressLayout` | Known | Navigation |
| 0x001FD418 | `GotoRentalWarningLayout` | Known | Navigation |
| 0x001FD49C | `GotoProgressLayout` | Known | Navigation |
| 0x001FD7AC | `GotoProgressLayout` | Known | Navigation |
| 0x001FF338 | `GotoNowPlaying` | Known | Navigation |
| 0x001FFC48 | `GotoNowPlaying` | Known | Navigation |
| 0x001FFF4C | `GotoNowPlaying` | Known | Navigation |
| 0x00202644 | `GotoScreen_Language` | Known | Navigation |
| 0x002029A4 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x002029C0 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x002029D8 | `GotoDefaultLayout` | Known | Navigation |
| 0x002029EC | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00202A84 | `GotoVolumeLayout` | Known | Navigation |
| 0x00202A98 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x00202B38 | `GotoProgressLayout` | Known | Navigation |
| 0x00202B4C | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00203300 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00203768 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x002039D4 | `GotoProgressLayout` | Known | Navigation |
| 0x002039E8 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00203B80 | `GotoBrightnessVideoLayout` | Known | Navigation |
| 0x00203BA4 | `GotoGeniusLayout` | Known | Navigation |
| 0x00203BB8 | `GotoRatingLayout` | Known | Navigation |
| 0x00203D2C | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x00203D48 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x00203D60 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x00204060 | `GotoChapterArtLayout` | Known | Navigation |
| 0x00204078 | `GotoShuffleLayout` | Known | Navigation |
| 0x00204408 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x0020441C | `GotoExtraInfoLoadingLayout` | Known | Navigation |
| 0x002044EC | `GotoVolumeLayout` | Known | Navigation |
| 0x00204504 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x00204590 | `GotoVolumeLayout` | Known | Navigation |
| 0x002045A4 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x002047B4 | `GotoScrubLayout` | Known | Navigation |
| 0x002047C4 | `GotoScrubVideoLayout` | Known | Navigation |
| 0x00204854 | `GotoProgressLayout` | Known | Navigation |
| 0x00204868 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00204AC0 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x00204ADC | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00204AF4 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00204B10 | `GotoDefaultLayout` | Known | Navigation |
| 0x00204D3C | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x00204D58 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x002052F4 | `GotoChapterArtLayout` | Known | Navigation |
| 0x002053EC | `GotoProgressLayout` | Known | Navigation |
| 0x00205478 | `GotoProgressLayout` | Known | Navigation |
| 0x0020548C | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00205568 | `GotoExtraInfoLoadFailedLayout` | Known | Navigation |
| 0x00205588 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x002059C4 | `GotoStatusBarLayout` | Known | Navigation |
| 0x002059D8 | `GotoDefaultLayout` | Known | Navigation |
| 0x00205BB0 | `GotoDefault` | Known | Navigation |
| 0x00205CE4 | `GotoProgressLayout` | Known | Navigation |
| 0x00205EA4 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x00205FF4 | `GotoBrightnessLayout` | Known | Navigation |
| 0x00206078 | `GotoBrightnessLayout` | Known | Navigation |
| 0x002060F8 | `GotoVolumeLayout` | Known | Navigation |
| 0x00206144 | `GotoScrubLayout` | Known | Navigation |
| 0x0020620C | `GotoStatusBarLayout` | Known | Navigation |
| 0x00206220 | `GotoDefaultLayout` | Known | Navigation |
| 0x002062F8 | `GotoScrubLayout` | Known | Navigation |
| 0x00206348 | `GotoScrubLayout` | Known | Navigation |
| 0x00208E1C | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020C290 | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x0020C2AC | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020C2C4 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x0020C474 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020C968 | `GotoNowPlaying` | Known | Navigation |
| 0x0020CC50 | `GotoNowPlaying` | Known | Navigation |
| 0x0020DDAC | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x0020DF3C | `GotoFourCard_About` | Known | Navigation |
| 0x0020DF50 | `GotoThreeCard_About` | Known | Navigation |
| 0x0020E03C | `GotoScreen_BacklightTimer` | Known | Navigation |
| 0x0020E0CC | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x0020E0E4 | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation |
| 0x00212B3C | `GotoScreen_LockDialog` | Known | Navigation |
| 0x00212B54 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x00213AE4 | `GotoGeniusIntro` | Known | Navigation |
| 0x00213AF8 | `GotoGenius` | Known | Navigation |
| 0x002151A0 | `GotoNowPlaying` | Known | Navigation |
| 0x002158B0 | `GotoNowPlaying` | Known | Navigation |
| 0x00216094 | `GotoFirstBoot` | Known | Navigation |
| 0x002160A4 | `GotoNotesApp` | Known | Navigation |
| 0x002160B8 | `GotoLockApp` | Known | Navigation |
| 0x002173F0 | `GotoGenius` | Known | Navigation |
| 0x0021D284 | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x0021D2A0 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0021D2B8 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x0021D468 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0021DC44 | `GotoNowPlaying` | Known | Navigation |
| 0x003F146C | `GotoRatingLayout` | Known | Navigation |
| 0x003F1480 | `GotoProgressLayout` | Known | Navigation |
| 0x00725E27 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x007A7784 | `GotoDefault` | Known | Navigation |
| 0x007A8768 | `GotoDefault` | Known | Navigation |
| 0x008997AC | `Gotowe?` | Known | Navigation |

---

## 7. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00156B90 | `CoverFlow_Screen` | Known | Screen layout |
| 0x00719D72 | `Clock_Screen` | Known | Screen layout |
| 0x00719D82 | `Clock_Screen_Default"` | Known | Screen layout |
| 0x00719DE7 | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x00719E45 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x00719E5D | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x00719ECA | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x00719F68 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x00719FC7 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x00719FDD | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x0071A048 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x0071A0A2 | `Games_Menu_Screen` | Known | Screen layout |
| 0x0071A0B7 | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x0071A121 | `Extras_Screen_Games` | Known | Screen layout |
| 0x0071A1E0 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x0071A2A4 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0071A36D | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x0071A3CA | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x0071A3E3 | `Debug_MainMenu_Screen_Default"` | Known | Screen layout |
| 0x0071A451 | `Extras_Screen_Debug` | Known | Screen layout |
| 0x0071A590 | `Clock_Africa_City_Screen ` | Known | Screen layout |
| 0x0071A5AC | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x0071A630 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x0071A64A | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x0071A6CC | `Clock_Atlantic_City_Screen"` | Known | Screen layout |
| 0x0071A6EA | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x0071A770 | `Clock_Australia_City_Screen#` | Known | Screen layout |
| 0x0071A78F | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x0071A816 | `Clock_Europe_City_Screen ` | Known | Screen layout |
| 0x0071A832 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x0071A8B6 | `Clock_NorthAmerica_City_Screen&` | Known | Screen layout |
| 0x0071A8D8 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0071A962 | `Clock_Pacific_City_Screen!` | Known | Screen layout |
| 0x0071A97F | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x0071AA04 | `Clock_SouthAmerica_City_Screen&` | Known | Screen layout |
| 0x0071AA26 | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0071AAB3 | `Clock_Screen"` | Known | Screen layout |
| 0x0071AB58 | `Clock_Screen"` | Known | Screen layout |
| 0x0071ABFD | `Clock_Screen"` | Known | Screen layout |
| 0x0071ACA2 | `Clock_Screen"` | Known | Screen layout |
| 0x0071AD47 | `Clock_Screen"` | Known | Screen layout |
| 0x0071ADEC | `Clock_Screen"` | Known | Screen layout |
| 0x0071AE91 | `Clock_Screen"` | Known | Screen layout |
| 0x0071AF36 | `Clock_Screen"` | Known | Screen layout |
| 0x0071AFDB | `Clock_Screen"` | Known | Screen layout |
| 0x0071B080 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B125 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B1CA | `Clock_Screen"` | Known | Screen layout |
| 0x0071B26F | `Clock_Screen"` | Known | Screen layout |
| 0x0071B314 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B3B9 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B45E | `Clock_Screen"` | Known | Screen layout |
| 0x0071B503 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B5A8 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B64D | `Clock_Screen"` | Known | Screen layout |
| 0x0071B6F2 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B797 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B83C | `Clock_Screen"` | Known | Screen layout |
| 0x0071B8E1 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B986 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BA2B | `Clock_Screen"` | Known | Screen layout |
| 0x0071BAD0 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BB75 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BC1A | `Clock_Screen"` | Known | Screen layout |
| 0x0071BCBF | `Clock_Screen"` | Known | Screen layout |
| 0x0071BD64 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BE09 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BEB3 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BF58 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BFFD | `Clock_Screen"` | Known | Screen layout |
| 0x0071C0A2 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C147 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C1EC | `Clock_Screen"` | Known | Screen layout |
| 0x0071C291 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C336 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C3DB | `Clock_Screen"` | Known | Screen layout |
| 0x0071C480 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C525 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C5CA | `Clock_Screen"` | Known | Screen layout |
| 0x0071C66F | `Clock_Screen"` | Known | Screen layout |
| 0x0071C714 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C7B9 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C85E | `Clock_Screen"` | Known | Screen layout |
| 0x0071C903 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C9A8 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CA4D | `Clock_Screen"` | Known | Screen layout |
| 0x0071CAF2 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CB97 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CC3C | `Clock_Screen"` | Known | Screen layout |
| 0x0071CCE1 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CD86 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CE2B | `Clock_Screen"` | Known | Screen layout |
| 0x0071CED0 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CF75 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D01A | `Clock_Screen"` | Known | Screen layout |
| 0x0071D0BF | `Clock_Screen"` | Known | Screen layout |
| 0x0071D164 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D209 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D2AE | `Clock_Screen"` | Known | Screen layout |
| 0x0071D353 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D3F8 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D49D | `Clock_Screen"` | Known | Screen layout |
| 0x0071D542 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D5E7 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D68C | `Clock_Screen"` | Known | Screen layout |
| 0x0071D731 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D7D6 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D87B | `Clock_Screen"` | Known | Screen layout |
| 0x0071D920 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D9C5 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DA6A | `Clock_Screen"` | Known | Screen layout |
| 0x0071DB0F | `Clock_Screen"` | Known | Screen layout |
| 0x0071DBB4 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DC59 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DCFE | `Clock_Screen"` | Known | Screen layout |
| 0x0071DDA3 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DE48 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DEED | `Clock_Screen"` | Known | Screen layout |
| 0x0071DF92 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E037 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E0DC | `Clock_Screen"` | Known | Screen layout |
| 0x0071E181 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E226 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E2CB | `Clock_Screen"` | Known | Screen layout |
| 0x0071E377 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E41C | `Clock_Screen"` | Known | Screen layout |
| 0x0071E4C1 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E56B | `Clock_Screen"` | Known | Screen layout |
| 0x0071E610 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E6B5 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E75A | `Clock_Screen"` | Known | Screen layout |
| 0x0071E7FF | `Clock_Screen"` | Known | Screen layout |
| 0x0071E8A4 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E949 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E9EE | `Clock_Screen"` | Known | Screen layout |
| 0x0071EA97 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EB3C | `Clock_Screen"` | Known | Screen layout |
| 0x0071EBE1 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EC86 | `Clock_Screen"` | Known | Screen layout |
| 0x0071ED2B | `Clock_Screen"` | Known | Screen layout |
| 0x0071EDD0 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EE75 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EF1A | `Clock_Screen"` | Known | Screen layout |
| 0x0071EFBF | `Clock_Screen"` | Known | Screen layout |
| 0x0071F064 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F109 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F1AE | `Clock_Screen"` | Known | Screen layout |
| 0x0071F253 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F2F8 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F39D | `Clock_Screen"` | Known | Screen layout |
| 0x0071F442 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F4E7 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F58C | `Clock_Screen"` | Known | Screen layout |
| 0x0071F631 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F6D6 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F77B | `Clock_Screen"` | Known | Screen layout |
| 0x0071F820 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F8C5 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F96A | `Clock_Screen"` | Known | Screen layout |
| 0x0071FA0F | `Clock_Screen"` | Known | Screen layout |
| 0x0071FAB4 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FB59 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FBFE | `Clock_Screen"` | Known | Screen layout |
| 0x0071FCA3 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FD48 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FDED | `Clock_Screen"` | Known | Screen layout |
| 0x0071FE92 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FF37 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FFDC | `Clock_Screen"` | Known | Screen layout |
| 0x00720087 | `Clock_Screen"` | Known | Screen layout |
| 0x0072012C | `Clock_Screen"` | Known | Screen layout |
| 0x007201D1 | `Clock_Screen"` | Known | Screen layout |
| 0x00720276 | `Clock_Screen"` | Known | Screen layout |
| 0x0072031B | `Clock_Screen"` | Known | Screen layout |
| 0x007203C0 | `Clock_Screen"` | Known | Screen layout |
| 0x00720465 | `Clock_Screen"` | Known | Screen layout |
| 0x0072050A | `Clock_Screen"` | Known | Screen layout |
| 0x007205AF | `Clock_Screen"` | Known | Screen layout |
| 0x00720654 | `Clock_Screen"` | Known | Screen layout |
| 0x007206F9 | `Clock_Screen"` | Known | Screen layout |
| 0x0072079E | `Clock_Screen"` | Known | Screen layout |
| 0x00720843 | `Clock_Screen"` | Known | Screen layout |
| 0x007208E8 | `Clock_Screen"` | Known | Screen layout |
| 0x0072098D | `Clock_Screen"` | Known | Screen layout |
| 0x00720A32 | `Clock_Screen"` | Known | Screen layout |
| 0x00720AD7 | `Clock_Screen"` | Known | Screen layout |
| 0x00720B7C | `Clock_Screen"` | Known | Screen layout |
| 0x00720C21 | `Clock_Screen"` | Known | Screen layout |
| 0x00720CC6 | `Clock_Screen"` | Known | Screen layout |
| 0x00720D6B | `Clock_Screen"` | Known | Screen layout |
| 0x00720E10 | `Clock_Screen"` | Known | Screen layout |
| 0x00720EB5 | `Clock_Screen"` | Known | Screen layout |
| 0x00720F5A | `Clock_Screen"` | Known | Screen layout |
| 0x00720FFF | `Clock_Screen"` | Known | Screen layout |
| 0x007210A4 | `Clock_Screen"` | Known | Screen layout |
| 0x00721149 | `Clock_Screen"` | Known | Screen layout |
| 0x007211EE | `Clock_Screen"` | Known | Screen layout |
| 0x00721293 | `Clock_Screen"` | Known | Screen layout |
| 0x00721338 | `Clock_Screen"` | Known | Screen layout |
| 0x007213DD | `Clock_Screen"` | Known | Screen layout |
| 0x00721482 | `Clock_Screen"` | Known | Screen layout |
| 0x00721527 | `Clock_Screen"` | Known | Screen layout |
| 0x007215CC | `Clock_Screen"` | Known | Screen layout |
| 0x00721671 | `Clock_Screen"` | Known | Screen layout |
| 0x00721716 | `Clock_Screen"` | Known | Screen layout |
| 0x007217BB | `Clock_Screen"` | Known | Screen layout |
| 0x00721860 | `Clock_Screen"` | Known | Screen layout |
| 0x00721905 | `Clock_Screen"` | Known | Screen layout |
| 0x007219AA | `Clock_Screen"` | Known | Screen layout |
| 0x00721A4F | `Clock_Screen"` | Known | Screen layout |
| 0x00721AF4 | `Clock_Screen"` | Known | Screen layout |
| 0x00721B99 | `Clock_Screen"` | Known | Screen layout |
| 0x00721C3E | `Clock_Screen"` | Known | Screen layout |
| 0x00721CE3 | `Clock_Screen"` | Known | Screen layout |
| 0x00721D88 | `Clock_Screen"` | Known | Screen layout |
| 0x00721E2D | `Clock_Screen"` | Known | Screen layout |
| 0x00721ED2 | `Clock_Screen"` | Known | Screen layout |
| 0x00721F77 | `Clock_Screen"` | Known | Screen layout |
| 0x0072201C | `Clock_Screen"` | Known | Screen layout |
| 0x007220C7 | `Clock_Screen"` | Known | Screen layout |
| 0x0072216C | `Clock_Screen"` | Known | Screen layout |
| 0x00722211 | `Clock_Screen"` | Known | Screen layout |
| 0x007222B6 | `Clock_Screen"` | Known | Screen layout |
| 0x0072235B | `Clock_Screen"` | Known | Screen layout |
| 0x00722407 | `Clock_Screen"` | Known | Screen layout |
| 0x007224AC | `Clock_Screen"` | Known | Screen layout |
| 0x00722551 | `Clock_Screen"` | Known | Screen layout |
| 0x007225F6 | `Clock_Screen"` | Known | Screen layout |
| 0x0072269B | `Clock_Screen"` | Known | Screen layout |
| 0x00722740 | `Clock_Screen"` | Known | Screen layout |
| 0x007227E5 | `Clock_Screen"` | Known | Screen layout |
| 0x0072288A | `Clock_Screen"` | Known | Screen layout |
| 0x0072292F | `Clock_Screen"` | Known | Screen layout |
| 0x007229D4 | `Clock_Screen"` | Known | Screen layout |
| 0x00722A79 | `Clock_Screen"` | Known | Screen layout |
| 0x00722B1E | `Clock_Screen"` | Known | Screen layout |
| 0x00722BC3 | `Clock_Screen"` | Known | Screen layout |
| 0x00722C68 | `Clock_Screen"` | Known | Screen layout |
| 0x00722D0D | `Clock_Screen"` | Known | Screen layout |
| 0x00722DB2 | `Clock_Screen"` | Known | Screen layout |
| 0x00722E57 | `Clock_Screen"` | Known | Screen layout |
| 0x00722EFA | `Settings_DateTime_SetDate_Screen(` | Known | Screen layout |
| 0x00722F1E | `Settings_DateTime_SetDate_Screen_Default"` | Known | Screen layout |
| 0x00722F97 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x00722FFD | `Settings_DateTime_SetTime_Screen(` | Known | Screen layout |
| 0x00723021 | `Settings_DateTime_SetTime_Screen_Default"` | Known | Screen layout |
| 0x0072309A | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x00723105 | `Settings_DateTime_SetTimeZone_Screen,` | Known | Screen layout |
| 0x0072312D | `Settings_DateTime_SetTimeZone_Screen_Default"` | Known | Screen layout |
| 0x007231AA | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x00723263 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x00723313 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0072391A | `Search_Main_Screen` | Known | Screen layout |
| 0x00723930 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x00723E52 | `Extras_Screen` | Known | Screen layout |
| 0x00723E63 | `Extras_Screen_WorldClock"` | Known | Screen layout |
| 0x00723EE0 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x00723F42 | `Clock_Screen` | Known | Screen layout |
| 0x00723F52 | `Clock_Screen_Default` | Known | Screen layout |
| 0x00723FD9 | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x0072403F | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x00724055 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x007240C0 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x00724122 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0072413A | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x007241A7 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x0072420B | `AddressViewer_Main_Screen!` | Known | Screen layout |
| 0x00724228 | `AddressViewer_Main_Screen_Default"` | Known | Screen layout |
| 0x0072429A | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x00724301 | `Games_Menu_Screen` | Known | Screen layout |
| 0x00724316 | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x00724380 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x00724447 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x007244E3 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x007245B4 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x00724674 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x007246D8 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007246F7 | `VoiceMemos_Menu_Screen_Default"` | Known | Screen layout |
| 0x0072477A | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x007247E0 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x007247F8 | `Speakers_Main_Screen_Default"` | Known | Screen layout |
| 0x00724879 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x007248DD | `Radio_Screen` | Known | Screen layout |
| 0x007248ED | `Radio_Screen_Default"` | Known | Screen layout |
| 0x00724966 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x007249C7 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00724A63 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x00724B26 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x00724BE5 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x00724CA2 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x007250F2 | `Radio_Screen` | Known | Screen layout |
| 0x00725102 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x0072517B | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x0072535F | `Search_Main_Screen` | Known | Screen layout |
| 0x00725375 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x007254A0 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x00725503 | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x00725844 | `Video_Settings_Screen` | Known | Screen layout |
| 0x0072585D | `Video_Settings_Screen_Default"` | Known | Screen layout |
| 0x00725966 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x00725C2B | `SettingsMenus_MainMenu_Screen%` | Known | Screen layout |
| 0x00725D39 | `SettingsMenus_MusicMenu_Screen&` | Known | Screen layout |
| 0x00725FE2 | `SettingsMenus_Brightness_Screen&` | Known | Screen layout |
| 0x007260F7 | `SettingsMenus_AudiobookSettings_Screen.` | Known | Screen layout |
| 0x0072622D | `SettingsMenus_RadioRegions_Screen ` | Known | Screen layout |
| 0x00726342 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x007265AE | `Settings_DateTime_Screen ` | Known | Screen layout |
| 0x007265CA | `Settings_DateTime_Screen_Default"` | Known | Screen layout |
| 0x00726756 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x0072685B | `Settings_Legal_Screen` | Known | Screen layout |
| 0x00726874 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x00726965 | `SettingsMenus_ResetAllSettings_Screen4` | Known | Screen layout |
| 0x00727136 | `Stopwatch_Screen` | Known | Screen layout |
| 0x0072714A | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007271B1 | `Stopwatch_Screen` | Known | Screen layout |
| 0x007271C5 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x0072726E | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x00727291 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x0072732A | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x0072734D | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x00727500 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x0072756E | `Speakers_ToneControl_Screen#` | Known | Screen layout |
| 0x0072758D | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x0073A679 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073A6FC | `LockediPod_Screen` | Known | Screen layout |
| 0x0073A784 | `Lock_Screen` | Known | Screen layout |
| 0x0073A793 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073A92E | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x0073AA00 | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x0073AA6A | `AddressViewer_ContactDetails_Screen+` | Known | Screen layout |
| 0x0073AA91 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x0073AB0C | `Extras_Screen` | Known | Screen layout |
| 0x0073AB57 | `Extras_Screen` | Known | Screen layout |
| 0x0073AC3E | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x0073AC9C | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0073ACB9 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x0073AD27 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0073AD40 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0073ADB7 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0073ADD4 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x0073AE3F | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x0073AE5C | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0073AEC3 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x0073AF2A | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x0073AF88 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0073AFA5 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x0073B013 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0073B02C | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0073B0A3 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0073B0C0 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x0073B12B | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x0073B148 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0073B1AF | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x0073B24F | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x0073B2D8 | `Alarms_Set_Alarm_Frequency_Screen)` | Known | Screen layout |
| 0x0073B2FD | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x0073B36E | `Alarms_Set_Alarm_Label_Screen%` | Known | Screen layout |
| 0x0073B38F | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x0073B3FC | `Alarms_Set_Alarm_Sound_Screen%` | Known | Screen layout |
| 0x0073B41D | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x0073B489 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x0073B704 | `Alarms_Set_Alarm_Playlist_Screen(` | Known | Screen layout |
| 0x0073B728 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x0073B798 | `Alarms_Set_Alarm_Tones_Screen%` | Known | Screen layout |
| 0x0073B7B9 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x0073BACC | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x0073BAE7 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x0073BC38 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0073BC4F | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x0073BCD0 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0073BCE7 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0073BDBD | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0073BDD6 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0073BE5B | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x0073BECC | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0073BFC1 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0073BFDA | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0073C05F | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x0073C0D0 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0073C190 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x0073C1A4 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x0073C2D3 | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x0073C336 | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x0073C38D | `Clock_Screen_Default` | Known | Screen layout |
| 0x0073C41E | `Clock_Region_Screen` | Known | Screen layout |
| 0x0073C435 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x0073C4AE | `Clock_Screen_Default` | Known | Screen layout |
| 0x0073C505 | `Clock_Screen_Default` | Known | Screen layout |
| 0x0073C596 | `Clock_Region_Screen` | Known | Screen layout |
| 0x0073C5AD | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x0073C738 | `CoverFlow_Screen_DefaultNoAlbumInfo&` | Known | Screen layout |
| 0x0073C826 | `CoverFlow_Screen_DefaultNoAlbumInfo#` | Known | Screen layout |
| 0x0073C89B | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073CB91 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073CD41 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073CE6F | `CoverFlow_Screen_Default&` | Known | Screen layout |
| 0x0073CF45 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073D0DA | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073D33F | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x0073D39C | `Game_Screen` | Known | Screen layout |
| 0x0073D3AB | `Game_Screen_Default` | Known | Screen layout |
| 0x0073D44D | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073D4AF | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073D512 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073D575 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073D5D1 | `Game_Running_Screen` | Known | Screen layout |
| 0x0073D631 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073D693 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073D6F6 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073D759 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073D7B5 | `Game_Running_Screen` | Known | Screen layout |
| 0x0073D815 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073D877 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073D8DA | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073D93D | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073D999 | `Game_Running_Screen` | Known | Screen layout |
| 0x0073D9F9 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073DA5B | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073DABE | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073DB21 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073DB7D | `Game_Running_Screen` | Known | Screen layout |
| 0x0073DBDD | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073DC3F | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073DCA2 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073DD05 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073DD61 | `Game_Running_Screen` | Known | Screen layout |
| 0x0073DFA7 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073E009 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073E06C | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073E0CF | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073E12B | `Game_Running_Screen` | Known | Screen layout |
| 0x0073E1E2 | `Extras_Screen` | Known | Screen layout |
| 0x0073E1F3 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073E251 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0073E3EE | `Extras_Screen` | Known | Screen layout |
| 0x0073E3FF | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073E45D | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0073E5FA | `Extras_Screen` | Known | Screen layout |
| 0x0073E60B | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073E669 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0073E806 | `Extras_Screen` | Known | Screen layout |
| 0x0073E817 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073E875 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0073EA17 | `Lock_Screen` | Known | Screen layout |
| 0x0073EA26 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0073EA88 | `Extras_Screen` | Known | Screen layout |
| 0x0073EA99 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0073EAF8 | `LockediPod_Screen` | Known | Screen layout |
| 0x0073EB72 | `MainMenus_Main_Screen)` | Known | Screen layout |
| 0x0073ED43 | `Lock_Screen` | Known | Screen layout |
| 0x0073ED52 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0073EDB4 | `Extras_Screen` | Known | Screen layout |
| 0x0073EDC5 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0073EE24 | `LockediPod_Screen` | Known | Screen layout |
| 0x0073EE9E | `MainMenus_Main_Screen%` | Known | Screen layout |
| 0x0073EF05 | `LockediPod_Screen` | Known | Screen layout |
| 0x0073EF1A | `LockediPod_Screen_Default,` | Known | Screen layout |
| 0x0073F069 | `Lock_Screen` | Known | Screen layout |
| 0x0073F078 | `Lock_Screen_Default_Layout"` | Known | Screen layout |
| 0x0073F0E1 | `Lock_Screen` | Known | Screen layout |
| 0x0073F0F0 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0073F152 | `Extras_Screen` | Known | Screen layout |
| 0x0073F163 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0073F1C2 | `LockediPod_Screen` | Known | Screen layout |
| 0x0073F23C | `MainMenus_Main_Screen,` | Known | Screen layout |
| 0x0073F398 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0073F3FE | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0073F462 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073F4F1 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0073F55E | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0073F5CB | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0073F638 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073F6A0 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0073F706 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0073F76A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073F7F9 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0073F866 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0073F8D3 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0073F940 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073F9A8 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0073FA0E | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0073FA72 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073FB01 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0073FB6E | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0073FBDB | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0073FC48 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073FCB0 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0073FD16 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0073FD7A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073FE09 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0073FE76 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0073FEE3 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0073FF50 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073FFB8 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0074001E | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x00740082 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00740111 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0074017E | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x007401EB | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x00740258 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007402B1 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x0074031A | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00740381 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0074041C | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00740485 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x007404EE | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00740555 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007405F0 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00740659 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x007406C2 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00740729 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007407C4 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x007408B0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007408CC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074093A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00740957 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007409C2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007409E2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00740A59 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00740A75 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00740AE5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00740B04 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00740B70 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00740B84 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00740BFD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00740C71 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00740CE1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00740D48 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00740DB0 | `NoContent_Screen` | Known | Screen layout |
| 0x00740DC4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00740E28 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00740E8F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00740EA9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00740F17 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00740F89 | `NoContent_Screen` | Known | Screen layout |
| 0x00740F9D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00741007 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00741070 | `No_Photos_Screen` | Known | Screen layout |
| 0x00741084 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007410EA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00741158 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007411C5 | `NoContent_Screen` | Known | Screen layout |
| 0x007411D9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00741241 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007412AB | `NoContent_Screen` | Known | Screen layout |
| 0x007412BF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00741326 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00741390 | `NoContent_Screen` | Known | Screen layout |
| 0x007413A4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00741411 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00741483 | `NoContent_Screen` | Known | Screen layout |
| 0x00741497 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007414FF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00741568 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00741583 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007415E9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00741605 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007416E4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007416FD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074175E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00741772 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007417CC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007417E8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074184F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00741866 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007419D7 | `Radio_Screen` | Known | Screen layout |
| 0x007419E7 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00741A48 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00741ACB | `LockediPod_Screen` | Known | Screen layout |
| 0x00741B53 | `Lock_Screen` | Known | Screen layout |
| 0x00741B62 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00741BC5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00741C27 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00741C43 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00741CB5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00741CD4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00741D3C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00741D56 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00741DBE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00741DDB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00741E47 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00741EB1 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00741ECB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00741F3B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00741FAE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074201F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074208E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007420FA | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00742115 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074218A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007421F1 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00742253 | `Photos_Screen` | Known | Screen layout |
| 0x007422B7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007422D5 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00742347 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00742364 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007423CA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007423E5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074244E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074246B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007424E2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00742506 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00742574 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074258F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074264C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00742668 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007426D6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007426F3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074275E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074277E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007427F5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00742811 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00742881 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007428A0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074290C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00742920 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00742999 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00742A0D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00742A7D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00742AE4 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00742B4C | `NoContent_Screen` | Known | Screen layout |
| 0x00742B60 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00742BC4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00742C2B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00742C45 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00742CB3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00742D25 | `NoContent_Screen` | Known | Screen layout |
| 0x00742D39 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00742DA3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00742E0C | `No_Photos_Screen` | Known | Screen layout |
| 0x00742E20 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00742E86 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00742EF4 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00742F61 | `NoContent_Screen` | Known | Screen layout |
| 0x00742F75 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00742FDD | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00743047 | `NoContent_Screen` | Known | Screen layout |
| 0x0074305B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007430C2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074312C | `NoContent_Screen` | Known | Screen layout |
| 0x00743140 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007431AD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074321F | `NoContent_Screen` | Known | Screen layout |
| 0x00743233 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074329B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00743304 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074331F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00743385 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007433A1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00743480 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00743499 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007434FA | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074350E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00743568 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00743584 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007435EB | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00743602 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00743773 | `Radio_Screen` | Known | Screen layout |
| 0x00743783 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007437E4 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00743867 | `LockediPod_Screen` | Known | Screen layout |
| 0x007438EF | `Lock_Screen` | Known | Screen layout |
| 0x007438FE | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00743961 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007439C3 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007439DF | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00743A51 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00743A70 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00743AD8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00743AF2 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00743B5A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00743B77 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00743BE3 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00743C4D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00743C67 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00743CD7 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00743D4A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00743DBB | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00743E2A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00743E96 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00743EB1 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00743F26 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00743F8D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00743FEF | `Photos_Screen` | Known | Screen layout |
| 0x00744053 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00744071 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007440E3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00744100 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00744166 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00744181 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007441EA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00744207 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074427E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007442A2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00744310 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074432B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007443E8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00744404 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00744472 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074448F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007444FA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074451A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00744591 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007445AD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074461D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074463C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007446A8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007446BC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00744735 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007447A9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00744819 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00744880 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007448E8 | `NoContent_Screen` | Known | Screen layout |
| 0x007448FC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00744960 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007449C7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007449E1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00744A4F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00744AC1 | `NoContent_Screen` | Known | Screen layout |
| 0x00744AD5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00744B3F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00744BA8 | `No_Photos_Screen` | Known | Screen layout |
| 0x00744BBC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00744C22 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00744C90 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00744CFD | `NoContent_Screen` | Known | Screen layout |
| 0x00744D11 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00744D79 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00744DE3 | `NoContent_Screen` | Known | Screen layout |
| 0x00744DF7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00744E5E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00744EC8 | `NoContent_Screen` | Known | Screen layout |
| 0x00744EDC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00744F49 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00744FBB | `NoContent_Screen` | Known | Screen layout |
| 0x00744FCF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00745037 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007450A0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007450BB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00745121 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074513D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074521C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00745235 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00745296 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007452AA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00745304 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00745320 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00745387 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074539E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074550F | `Radio_Screen` | Known | Screen layout |
| 0x0074551F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00745580 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00745603 | `LockediPod_Screen` | Known | Screen layout |
| 0x0074568B | `Lock_Screen` | Known | Screen layout |
| 0x0074569A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007456FD | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074575F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074577B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007457ED | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074580C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00745874 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074588E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007458F6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00745913 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074597F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007459E9 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00745A03 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00745A73 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00745AE6 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00745B57 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00745BC6 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00745C32 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00745C4D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00745CC2 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00745D29 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00745D8B | `Photos_Screen` | Known | Screen layout |
| 0x00745DEF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00745E0D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00745E7F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00745E9C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00745F02 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00745F1D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00745F86 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00745FA3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074601A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074603E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007460AC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007460C7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00746184 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007461A0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074620E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074622B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00746296 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007462B6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074632D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00746349 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007463B9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007463D8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00746444 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00746458 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007464D1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00746545 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007465B5 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074661C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00746684 | `NoContent_Screen` | Known | Screen layout |
| 0x00746698 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007466FC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00746763 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074677D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007467EB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074685D | `NoContent_Screen` | Known | Screen layout |
| 0x00746871 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007468DB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00746944 | `No_Photos_Screen` | Known | Screen layout |
| 0x00746958 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007469BE | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00746A2C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00746A99 | `NoContent_Screen` | Known | Screen layout |
| 0x00746AAD | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00746B15 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00746B7F | `NoContent_Screen` | Known | Screen layout |
| 0x00746B93 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00746BFA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00746C64 | `NoContent_Screen` | Known | Screen layout |
| 0x00746C78 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00746CE5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00746D57 | `NoContent_Screen` | Known | Screen layout |
| 0x00746D6B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00746DD3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00746E3C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00746E57 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00746EBD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00746ED9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00746FB8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00746FD1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00747032 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00747046 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007470A0 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007470BC | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00747123 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074713A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007472AB | `Radio_Screen` | Known | Screen layout |
| 0x007472BB | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074731C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074739F | `LockediPod_Screen` | Known | Screen layout |
| 0x00747427 | `Lock_Screen` | Known | Screen layout |
| 0x00747436 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00747499 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007474FB | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00747517 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00747589 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007475A8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00747610 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074762A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00747692 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007476AF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074771B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00747785 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074779F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074780F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00747882 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007478F3 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00747962 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007479CE | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007479E9 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00747A5E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00747AC5 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00747B27 | `Photos_Screen` | Known | Screen layout |
| 0x00747B8B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00747BA9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00747C1B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00747C38 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00747C9E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00747CB9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00747D22 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00747D3F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00747DB6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00747DDA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00747E48 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00747E63 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00747F20 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00747F3C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00747FAA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00747FC7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00748032 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00748052 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007480C9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007480E5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00748155 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00748174 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007481E0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007481F4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074826D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007482E1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00748351 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007483B8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00748420 | `NoContent_Screen` | Known | Screen layout |
| 0x00748434 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00748498 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007484FF | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00748519 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00748587 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007485F9 | `NoContent_Screen` | Known | Screen layout |
| 0x0074860D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00748677 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007486E0 | `No_Photos_Screen` | Known | Screen layout |
| 0x007486F4 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074875A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007487C8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00748835 | `NoContent_Screen` | Known | Screen layout |
| 0x00748849 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007488B1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074891B | `NoContent_Screen` | Known | Screen layout |
| 0x0074892F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00748996 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00748A00 | `NoContent_Screen` | Known | Screen layout |
| 0x00748A14 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00748A81 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00748AF3 | `NoContent_Screen` | Known | Screen layout |
| 0x00748B07 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00748B6F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00748BD8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00748BF3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00748C59 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00748C75 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00748D54 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00748D6D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00748DCE | `FirstBoot_Screen` | Known | Screen layout |
| 0x00748DE2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00748E3C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00748E58 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00748EBF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00748ED6 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00749047 | `Radio_Screen` | Known | Screen layout |
| 0x00749057 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007490B8 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074913B | `LockediPod_Screen` | Known | Screen layout |
| 0x007491C3 | `Lock_Screen` | Known | Screen layout |
| 0x007491D2 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00749235 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00749297 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007492B3 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00749325 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00749344 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007493AC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007493C6 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074942E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074944B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007494B7 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00749521 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074953B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007495AB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074961E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074968F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007496FE | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074976A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00749785 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007497FA | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00749861 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007498C3 | `Photos_Screen` | Known | Screen layout |
| 0x00749927 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00749945 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007499B7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007499D4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00749A3A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00749A55 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00749ABE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00749ADB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00749B52 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00749B76 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00749BE4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00749BFF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00749CBC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00749CD8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00749D46 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00749D63 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00749DCE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00749DEE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00749E65 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00749E81 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00749EF1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00749F10 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00749F7C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00749F90 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074A009 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074A07D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074A0ED | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074A154 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074A1BC | `NoContent_Screen` | Known | Screen layout |
| 0x0074A1D0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074A234 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074A29B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074A2B5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074A323 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074A395 | `NoContent_Screen` | Known | Screen layout |
| 0x0074A3A9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074A413 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074A47C | `No_Photos_Screen` | Known | Screen layout |
| 0x0074A490 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074A4F6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074A564 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074A5D1 | `NoContent_Screen` | Known | Screen layout |
| 0x0074A5E5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074A64D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074A6B7 | `NoContent_Screen` | Known | Screen layout |
| 0x0074A6CB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074A732 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074A79C | `NoContent_Screen` | Known | Screen layout |
| 0x0074A7B0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074A81D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074A88F | `NoContent_Screen` | Known | Screen layout |
| 0x0074A8A3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074A90B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074A974 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074A98F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074A9F5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074AA11 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074AAF0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074AB09 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074AB6A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074AB7E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074ABD8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074ABF4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074AC5B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074AC72 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074ADE3 | `Radio_Screen` | Known | Screen layout |
| 0x0074ADF3 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074AE54 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074AED7 | `LockediPod_Screen` | Known | Screen layout |
| 0x0074AF5F | `Lock_Screen` | Known | Screen layout |
| 0x0074AF6E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074AFD1 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074B033 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074B04F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0074B0C1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074B0E0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074B148 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074B162 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074B1CA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074B1E7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074B253 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074B2BD | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074B2D7 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074B347 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074B3BA | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074B42B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074B49A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074B506 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074B521 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074B596 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074B5FD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074B65F | `Photos_Screen` | Known | Screen layout |
| 0x0074B6C3 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074B6E1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074B753 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074B770 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074B7D6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074B7F1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074B85A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074B877 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074B8EE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074B912 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074B980 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074B99B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074BA58 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074BA74 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074BAE2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074BAFF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074BB6A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074BB8A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074BC01 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074BC1D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074BC8D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074BCAC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074BD18 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074BD2C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074BDA5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074BE19 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074BE89 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074BEF0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074BF58 | `NoContent_Screen` | Known | Screen layout |
| 0x0074BF6C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074BFD0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074C037 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074C051 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074C0BF | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074C131 | `NoContent_Screen` | Known | Screen layout |
| 0x0074C145 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074C1AF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074C218 | `No_Photos_Screen` | Known | Screen layout |
| 0x0074C22C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074C292 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074C300 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074C36D | `NoContent_Screen` | Known | Screen layout |
| 0x0074C381 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074C3E9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074C453 | `NoContent_Screen` | Known | Screen layout |
| 0x0074C467 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074C4CE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074C538 | `NoContent_Screen` | Known | Screen layout |
| 0x0074C54C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074C5B9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074C62B | `NoContent_Screen` | Known | Screen layout |
| 0x0074C63F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074C6A7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074C710 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074C72B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074C791 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074C7AD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074C88C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074C8A5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074C906 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074C91A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074C974 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074C990 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074C9F7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074CA0E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074CB7F | `Radio_Screen` | Known | Screen layout |
| 0x0074CB8F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074CBF0 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074CC73 | `LockediPod_Screen` | Known | Screen layout |
| 0x0074CCFB | `Lock_Screen` | Known | Screen layout |
| 0x0074CD0A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074CD6D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074CDCF | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074CDEB | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0074CE5D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074CE7C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074CEE4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074CEFE | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074CF66 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074CF83 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074CFEF | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074D059 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074D073 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074D0E3 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074D156 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074D1C7 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074D236 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074D2A2 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074D2BD | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074D332 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074D399 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074D3FB | `Photos_Screen` | Known | Screen layout |
| 0x0074D45F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074D47D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074D4EF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074D50C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074D572 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074D58D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074D5F6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074D613 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074D68A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074D6AE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074D71C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074D737 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074D7F4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074D810 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074D87E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074D89B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074D906 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074D926 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074D99D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074D9B9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074DA29 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074DA48 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074DAB4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074DAC8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074DB41 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074DBB5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074DC25 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074DC8C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074DCF4 | `NoContent_Screen` | Known | Screen layout |
| 0x0074DD08 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074DD6C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074DDD3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074DDED | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074DE5B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074DECD | `NoContent_Screen` | Known | Screen layout |
| 0x0074DEE1 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074DF4B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074DFB4 | `No_Photos_Screen` | Known | Screen layout |
| 0x0074DFC8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074E02E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074E09C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074E109 | `NoContent_Screen` | Known | Screen layout |
| 0x0074E11D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074E185 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074E1EF | `NoContent_Screen` | Known | Screen layout |
| 0x0074E203 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074E26A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074E2D4 | `NoContent_Screen` | Known | Screen layout |
| 0x0074E2E8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074E355 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074E3C7 | `NoContent_Screen` | Known | Screen layout |
| 0x0074E3DB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074E443 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074E4AC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074E4C7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074E52D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074E549 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074E628 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074E641 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074E6A2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074E6B6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074E710 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074E72C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074E793 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074E7AA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074E91B | `Radio_Screen` | Known | Screen layout |
| 0x0074E92B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074E98C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074EA0F | `LockediPod_Screen` | Known | Screen layout |
| 0x0074EA97 | `Lock_Screen` | Known | Screen layout |
| 0x0074EAA6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074EB09 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074EB6B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074EB87 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0074EBF9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074EC18 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074EC80 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074EC9A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074ED02 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074ED1F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074ED8B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074EDF5 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074EE0F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074EE7F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074EEF2 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074EF63 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074EFD2 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074F03E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074F059 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074F0CE | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074F135 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074F197 | `Photos_Screen` | Known | Screen layout |
| 0x0074F1FB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074F219 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074F28B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074F2A8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074F30E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074F329 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074F392 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074F3AF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074F426 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074F44A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074F4B8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074F4D3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074F590 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074F5AC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074F61A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074F637 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074F6A2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074F6C2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074F739 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074F755 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074F7C5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074F7E4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074F850 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074F864 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074F8DD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074F951 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074F9C1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074FA28 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074FA90 | `NoContent_Screen` | Known | Screen layout |
| 0x0074FAA4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074FB08 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074FB6F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074FB89 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074FBF7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074FC69 | `NoContent_Screen` | Known | Screen layout |
| 0x0074FC7D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074FCE7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074FD50 | `No_Photos_Screen` | Known | Screen layout |
| 0x0074FD64 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074FDCA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074FE38 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074FEA5 | `NoContent_Screen` | Known | Screen layout |
| 0x0074FEB9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074FF21 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074FF8B | `NoContent_Screen` | Known | Screen layout |
| 0x0074FF9F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00750006 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00750070 | `NoContent_Screen` | Known | Screen layout |
| 0x00750084 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007500F1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00750163 | `NoContent_Screen` | Known | Screen layout |
| 0x00750177 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007501DF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00750248 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00750263 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007502C9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007502E5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007503C4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007503DD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075043E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00750452 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007504AC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007504C8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075052F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00750546 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007506B7 | `Radio_Screen` | Known | Screen layout |
| 0x007506C7 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00750728 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007507AB | `LockediPod_Screen` | Known | Screen layout |
| 0x00750833 | `Lock_Screen` | Known | Screen layout |
| 0x00750842 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007508A5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00750907 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00750923 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00750995 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007509B4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00750A1C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00750A36 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00750A9E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00750ABB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00750B27 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00750B91 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00750BAB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00750C1B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00750C8E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00750CFF | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00750D6E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00750DDA | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00750DF5 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00750E6A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00750ED1 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00750F33 | `Photos_Screen` | Known | Screen layout |
| 0x00750F97 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00750FB5 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00751027 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00751044 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007510AA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007510C5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075112E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075114B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007511C2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007511E6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00751254 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075126F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075132C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00751348 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007513B6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007513D3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075143E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075145E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007514D5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007514F1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00751561 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00751580 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007515EC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00751600 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00751679 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007516ED | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075175D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007517C4 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075182C | `NoContent_Screen` | Known | Screen layout |
| 0x00751840 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007518A4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075190B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00751925 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00751993 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00751A05 | `NoContent_Screen` | Known | Screen layout |
| 0x00751A19 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00751A83 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00751AEC | `No_Photos_Screen` | Known | Screen layout |
| 0x00751B00 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00751B66 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00751BD4 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00751C41 | `NoContent_Screen` | Known | Screen layout |
| 0x00751C55 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00751CBD | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00751D27 | `NoContent_Screen` | Known | Screen layout |
| 0x00751D3B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00751DA2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00751E0C | `NoContent_Screen` | Known | Screen layout |
| 0x00751E20 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00751E8D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00751EFF | `NoContent_Screen` | Known | Screen layout |
| 0x00751F13 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00751F7B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00751FE4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00751FFF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00752065 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00752081 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00752160 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00752179 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007521DA | `FirstBoot_Screen` | Known | Screen layout |
| 0x007521EE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00752248 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00752264 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007522CB | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007522E2 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00752453 | `Radio_Screen` | Known | Screen layout |
| 0x00752463 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007524C4 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00752547 | `LockediPod_Screen` | Known | Screen layout |
| 0x007525CF | `Lock_Screen` | Known | Screen layout |
| 0x007525DE | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00752641 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007526A3 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007526BF | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00752731 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00752750 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007527B8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007527D2 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075283A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00752857 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007528C3 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075292D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00752947 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007529B7 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00752A2A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00752A9B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00752B0A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00752B76 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00752B91 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00752C06 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00752C6D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00752CCF | `Photos_Screen` | Known | Screen layout |
| 0x00752D33 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00752D51 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00752DC3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00752DE0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00752E46 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00752E61 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00752ECA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00752EE7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00752F5E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00752F82 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00752FF0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075300B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007530C8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007530E4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00753152 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075316F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007531DA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007531FA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00753271 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075328D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007532FD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075331C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00753388 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075339C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00753415 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00753489 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007534F9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00753560 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007535C8 | `NoContent_Screen` | Known | Screen layout |
| 0x007535DC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00753640 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007536A7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007536C1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075372F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007537A1 | `NoContent_Screen` | Known | Screen layout |
| 0x007537B5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075381F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00753888 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075389C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00753902 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00753970 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007539DD | `NoContent_Screen` | Known | Screen layout |
| 0x007539F1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00753A59 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00753AC3 | `NoContent_Screen` | Known | Screen layout |
| 0x00753AD7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00753B3E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00753BA8 | `NoContent_Screen` | Known | Screen layout |
| 0x00753BBC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00753C29 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00753C9B | `NoContent_Screen` | Known | Screen layout |
| 0x00753CAF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00753D17 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00753D80 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00753D9B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00753E01 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00753E1D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00753EFC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00753F15 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00753F76 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00753F8A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00753FE4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00754000 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00754067 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075407E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007541EF | `Radio_Screen` | Known | Screen layout |
| 0x007541FF | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00754260 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007542E3 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075436B | `Lock_Screen` | Known | Screen layout |
| 0x0075437A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007543DD | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075443F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075445B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007544CD | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007544EC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00754554 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075456E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007545D6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007545F3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075465F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007546C9 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007546E3 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00754753 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007547C6 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00754837 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007548A6 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00754912 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075492D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007549A2 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00754A09 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00754A6B | `Photos_Screen` | Known | Screen layout |
| 0x00754ACF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00754AED | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00754B5F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00754B7C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00754BE2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00754BFD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00754C66 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00754C83 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00754CFA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00754D1E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00754D8C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00754DA7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00754E64 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00754E80 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00754EEE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00754F0B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00754F76 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00754F96 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075500D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00755029 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00755099 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007550B8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00755124 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00755138 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007551B1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00755225 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00755295 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007552FC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00755364 | `NoContent_Screen` | Known | Screen layout |
| 0x00755378 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007553DC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00755443 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075545D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007554CB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075553D | `NoContent_Screen` | Known | Screen layout |
| 0x00755551 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007555BB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00755624 | `No_Photos_Screen` | Known | Screen layout |
| 0x00755638 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075569E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075570C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00755779 | `NoContent_Screen` | Known | Screen layout |
| 0x0075578D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007557F5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075585F | `NoContent_Screen` | Known | Screen layout |
| 0x00755873 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007558DA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00755944 | `NoContent_Screen` | Known | Screen layout |
| 0x00755958 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007559C5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00755A37 | `NoContent_Screen` | Known | Screen layout |
| 0x00755A4B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00755AB3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00755B1C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00755B37 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00755B9D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00755BB9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00755C98 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00755CB1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00755D12 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00755D26 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00755D80 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00755D9C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00755E03 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00755E1A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00755F8B | `Radio_Screen` | Known | Screen layout |
| 0x00755F9B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00755FFC | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075607F | `LockediPod_Screen` | Known | Screen layout |
| 0x00756107 | `Lock_Screen` | Known | Screen layout |
| 0x00756116 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00756179 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007561DB | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007561F7 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00756269 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00756288 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007562F0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075630A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00756372 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075638F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007563FB | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00756465 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075647F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007564EF | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00756562 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007565D3 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00756642 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007566AE | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007566C9 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075673E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007567A5 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00756807 | `Photos_Screen` | Known | Screen layout |
| 0x0075686B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00756889 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007568FB | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00756918 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075697E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00756999 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00756A02 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00756A1F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00756A96 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00756ABA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00756B28 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00756B43 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00756C00 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00756C1C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00756C8A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00756CA7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00756D12 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00756D32 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00756DA9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00756DC5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00756E35 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00756E54 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00756EC0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00756ED4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00756F4D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00756FC1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00757031 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00757098 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00757100 | `NoContent_Screen` | Known | Screen layout |
| 0x00757114 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00757178 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007571DF | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007571F9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00757267 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007572D9 | `NoContent_Screen` | Known | Screen layout |
| 0x007572ED | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00757357 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007573C0 | `No_Photos_Screen` | Known | Screen layout |
| 0x007573D4 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075743A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007574A8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00757515 | `NoContent_Screen` | Known | Screen layout |
| 0x00757529 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00757591 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007575FB | `NoContent_Screen` | Known | Screen layout |
| 0x0075760F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00757676 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007576E0 | `NoContent_Screen` | Known | Screen layout |
| 0x007576F4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00757761 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007577D3 | `NoContent_Screen` | Known | Screen layout |
| 0x007577E7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075784F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007578B8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007578D3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00757939 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00757955 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00757A34 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00757A4D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00757AAE | `FirstBoot_Screen` | Known | Screen layout |
| 0x00757AC2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00757B1C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00757B38 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00757B9F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00757BB6 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00757D27 | `Radio_Screen` | Known | Screen layout |
| 0x00757D37 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00757D98 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00757E1B | `LockediPod_Screen` | Known | Screen layout |
| 0x00757EA3 | `Lock_Screen` | Known | Screen layout |
| 0x00757EB2 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00757F15 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00757F77 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00757F93 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00758005 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00758024 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075808C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007580A6 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075810E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075812B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00758197 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00758201 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075821B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075828B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007582FE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075836F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007583DE | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075844A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00758465 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007584DA | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00758541 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007585A3 | `Photos_Screen` | Known | Screen layout |
| 0x00758607 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00758625 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00758697 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007586B4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075871A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00758735 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075879E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007587BB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00758832 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00758856 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007588C4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007588DF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075899C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007589B8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00758A26 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00758A43 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00758AAE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00758ACE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00758B45 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00758B61 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00758BD1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00758BF0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00758C5C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00758C70 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00758CE9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00758D5D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00758DCD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00758E34 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00758E9C | `NoContent_Screen` | Known | Screen layout |
| 0x00758EB0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00758F14 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00758F7B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00758F95 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00759003 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00759075 | `NoContent_Screen` | Known | Screen layout |
| 0x00759089 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007590F3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075915C | `No_Photos_Screen` | Known | Screen layout |
| 0x00759170 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007591D6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00759244 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007592B1 | `NoContent_Screen` | Known | Screen layout |
| 0x007592C5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075932D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00759397 | `NoContent_Screen` | Known | Screen layout |
| 0x007593AB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00759412 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075947C | `NoContent_Screen` | Known | Screen layout |
| 0x00759490 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007594FD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075956F | `NoContent_Screen` | Known | Screen layout |
| 0x00759583 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007595EB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00759654 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075966F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007596D5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007596F1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007597D0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007597E9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075984A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075985E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007598B8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007598D4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075993B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00759952 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00759AC3 | `Radio_Screen` | Known | Screen layout |
| 0x00759AD3 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00759B34 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00759BB7 | `LockediPod_Screen` | Known | Screen layout |
| 0x00759C3F | `Lock_Screen` | Known | Screen layout |
| 0x00759C4E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00759CB1 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00759D13 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00759D2F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00759DA1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00759DC0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00759E28 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00759E42 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00759EAA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00759EC7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00759F33 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00759F9D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00759FB7 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075A027 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075A09A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075A10B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075A17A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075A1E6 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075A201 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075A276 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075A2DD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075A33F | `Photos_Screen` | Known | Screen layout |
| 0x0075A3A3 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075A3C1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075A433 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075A450 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075A4B6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075A4D1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075A53A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075A557 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075A5CE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075A5F2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075A660 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075A67B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075A738 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075A754 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075A7C2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075A7DF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075A84A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075A86A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075A8E1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075A8FD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075A96D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075A98C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075A9F8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075AA0C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075AA85 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075AAF9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075AB69 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075ABD0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075AC38 | `NoContent_Screen` | Known | Screen layout |
| 0x0075AC4C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075ACB0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075AD17 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075AD31 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075AD9F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075AE11 | `NoContent_Screen` | Known | Screen layout |
| 0x0075AE25 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075AE8F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075AEF8 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075AF0C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075AF72 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075AFE0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075B04D | `NoContent_Screen` | Known | Screen layout |
| 0x0075B061 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075B0C9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075B133 | `NoContent_Screen` | Known | Screen layout |
| 0x0075B147 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075B1AE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075B218 | `NoContent_Screen` | Known | Screen layout |
| 0x0075B22C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075B299 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075B30B | `NoContent_Screen` | Known | Screen layout |
| 0x0075B31F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075B387 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075B3F0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075B40B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075B471 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075B48D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075B56C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075B585 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075B5E6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075B5FA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075B654 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075B670 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075B6D7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075B6EE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075B85F | `Radio_Screen` | Known | Screen layout |
| 0x0075B86F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075B8D0 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075B953 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075B9DB | `Lock_Screen` | Known | Screen layout |
| 0x0075B9EA | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075BA4D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075BAAF | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075BACB | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075BB3D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075BB5C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075BBC4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075BBDE | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075BC46 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075BC63 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075BCCF | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075BD39 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075BD53 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075BDC3 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075BE36 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075BEA7 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075BF16 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075BF82 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075BF9D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075C012 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075C079 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075C0DB | `Photos_Screen` | Known | Screen layout |
| 0x0075C13F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075C15D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075C1CF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075C1EC | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075C252 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075C26D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075C2D6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075C2F3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075C36A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075C38E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075C3FC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075C417 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075C4D4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075C4F0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075C55E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075C57B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075C5E6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075C606 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075C67D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075C699 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075C709 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075C728 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075C794 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075C7A8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075C821 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075C895 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075C905 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075C96C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075C9D4 | `NoContent_Screen` | Known | Screen layout |
| 0x0075C9E8 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075CA4C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075CAB3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075CACD | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075CB3B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075CBAD | `NoContent_Screen` | Known | Screen layout |
| 0x0075CBC1 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075CC2B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075CC94 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075CCA8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075CD0E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075CD7C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075CDE9 | `NoContent_Screen` | Known | Screen layout |
| 0x0075CDFD | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075CE65 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075CECF | `NoContent_Screen` | Known | Screen layout |
| 0x0075CEE3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075CF4A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075CFB4 | `NoContent_Screen` | Known | Screen layout |
| 0x0075CFC8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075D035 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075D0A7 | `NoContent_Screen` | Known | Screen layout |
| 0x0075D0BB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075D123 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075D18C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075D1A7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075D20D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075D229 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075D308 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075D321 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075D382 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075D396 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075D3F0 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075D40C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075D473 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075D48A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075D5FB | `Radio_Screen` | Known | Screen layout |
| 0x0075D60B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075D66C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075D6EF | `LockediPod_Screen` | Known | Screen layout |
| 0x0075D777 | `Lock_Screen` | Known | Screen layout |
| 0x0075D786 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075D7E9 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075D84B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075D867 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075D8D9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075D8F8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075D960 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075D97A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075D9E2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075D9FF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075DA6B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075DAD5 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075DAEF | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075DB5F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075DBD2 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075DC43 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075DCB2 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075DD1E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075DD39 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075DDAE | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075DE15 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075DE77 | `Photos_Screen` | Known | Screen layout |
| 0x0075DEDB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075DEF9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075DF6B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075DF88 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075DFEE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075E009 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075E072 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075E08F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075E106 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075E12A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075E198 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075E1B3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075E270 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075E28C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075E2FA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075E317 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075E382 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075E3A2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075E419 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075E435 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075E4A5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075E4C4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075E530 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075E544 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075E5BD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075E631 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075E6A1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075E708 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075E770 | `NoContent_Screen` | Known | Screen layout |
| 0x0075E784 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075E7E8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075E84F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075E869 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075E8D7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075E949 | `NoContent_Screen` | Known | Screen layout |
| 0x0075E95D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075E9C7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075EA30 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075EA44 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075EAAA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075EB18 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075EB85 | `NoContent_Screen` | Known | Screen layout |
| 0x0075EB99 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075EC01 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075EC6B | `NoContent_Screen` | Known | Screen layout |
| 0x0075EC7F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075ECE6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075ED50 | `NoContent_Screen` | Known | Screen layout |
| 0x0075ED64 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075EDD1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075EE43 | `NoContent_Screen` | Known | Screen layout |
| 0x0075EE57 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075EEBF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075EF28 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075EF43 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075EFA9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075EFC5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075F0A4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075F0BD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075F11E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075F132 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075F18C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075F1A8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075F20F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075F226 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075F397 | `Radio_Screen` | Known | Screen layout |
| 0x0075F3A7 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075F408 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075F48B | `LockediPod_Screen` | Known | Screen layout |
| 0x0075F513 | `Lock_Screen` | Known | Screen layout |
| 0x0075F522 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075F585 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075F5E7 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075F603 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075F675 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075F694 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075F6FC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075F716 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075F77E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075F79B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075F807 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075F871 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075F88B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075F8FB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075F96E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075F9DF | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075FA4E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075FABA | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075FAD5 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075FB4A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075FBB1 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075FC13 | `Photos_Screen` | Known | Screen layout |
| 0x0075FC77 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075FC95 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075FD07 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075FD24 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075FD8A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075FDA5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075FE0E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075FE2B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075FEA2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075FEC6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075FF34 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075FF4F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076000C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00760028 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00760096 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007600B3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076011E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076013E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007601B5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007601D1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00760241 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00760260 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007602CC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007602E0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00760359 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007603CD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076043D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007604A4 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076050C | `NoContent_Screen` | Known | Screen layout |
| 0x00760520 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00760584 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007605EB | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00760605 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00760673 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007606E5 | `NoContent_Screen` | Known | Screen layout |
| 0x007606F9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00760763 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007607CC | `No_Photos_Screen` | Known | Screen layout |
| 0x007607E0 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00760846 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007608B4 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00760921 | `NoContent_Screen` | Known | Screen layout |
| 0x00760935 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076099D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00760A07 | `NoContent_Screen` | Known | Screen layout |
| 0x00760A1B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00760A82 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00760AEC | `NoContent_Screen` | Known | Screen layout |
| 0x00760B00 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00760B6D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00760BDF | `NoContent_Screen` | Known | Screen layout |
| 0x00760BF3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00760C5B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00760CC4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00760CDF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00760D45 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00760D61 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00760E40 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00760E59 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00760EBA | `FirstBoot_Screen` | Known | Screen layout |
| 0x00760ECE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00760F28 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00760F44 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00760FAB | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00760FC2 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00761133 | `Radio_Screen` | Known | Screen layout |
| 0x00761143 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007611A4 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00761227 | `LockediPod_Screen` | Known | Screen layout |
| 0x007612AF | `Lock_Screen` | Known | Screen layout |
| 0x007612BE | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00761321 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00761383 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076139F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00761411 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00761430 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00761498 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007614B2 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076151A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00761537 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007615A3 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076160D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00761627 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00761697 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076170A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076177B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007617EA | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00761856 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00761871 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007618E6 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076194D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007619AF | `Photos_Screen` | Known | Screen layout |
| 0x00761A13 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00761A31 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00761AA3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00761AC0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00761B26 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00761B41 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00761BAA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00761BC7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00761C3E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00761C62 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00761CD0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00761CEB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00761DA8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00761DC4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00761E32 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00761E4F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00761EBA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00761EDA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00761F51 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00761F6D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00761FDD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00761FFC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00762068 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076207C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007620F5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00762169 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007621D9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00762240 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007622A8 | `NoContent_Screen` | Known | Screen layout |
| 0x007622BC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00762320 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00762387 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007623A1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076240F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00762481 | `NoContent_Screen` | Known | Screen layout |
| 0x00762495 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007624FF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00762568 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076257C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007625E2 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00762650 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007626BD | `NoContent_Screen` | Known | Screen layout |
| 0x007626D1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00762739 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007627A3 | `NoContent_Screen` | Known | Screen layout |
| 0x007627B7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076281E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00762888 | `NoContent_Screen` | Known | Screen layout |
| 0x0076289C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00762909 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076297B | `NoContent_Screen` | Known | Screen layout |
| 0x0076298F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007629F7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00762A60 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00762A7B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00762AE1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00762AFD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00762BDC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00762BF5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00762C56 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00762C6A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00762CC4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00762CE0 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00762D47 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00762D5E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00762ECF | `Radio_Screen` | Known | Screen layout |
| 0x00762EDF | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00762F40 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00762FC3 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076304B | `Lock_Screen` | Known | Screen layout |
| 0x0076305A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007630BD | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076311F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076313B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007631AD | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007631CC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00763234 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076324E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007632B6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007632D3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076333F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007633A9 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007633C3 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00763433 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007634A6 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00763517 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00763586 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007635F2 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076360D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00763682 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007636E9 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076374B | `Photos_Screen` | Known | Screen layout |
| 0x007637AF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007637CD | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076383F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076385C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007638C2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007638DD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00763946 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00763963 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007639DA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007639FE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00763A6C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00763A87 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00763B44 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00763B60 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00763BCE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00763BEB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00763C56 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00763C76 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00763CED | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00763D09 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00763D79 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00763D98 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00763E04 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00763E18 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00763E91 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00763F05 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00763F75 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00763FDC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00764044 | `NoContent_Screen` | Known | Screen layout |
| 0x00764058 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007640BC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00764123 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076413D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007641AB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076421D | `NoContent_Screen` | Known | Screen layout |
| 0x00764231 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076429B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00764304 | `No_Photos_Screen` | Known | Screen layout |
| 0x00764318 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076437E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007643EC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00764459 | `NoContent_Screen` | Known | Screen layout |
| 0x0076446D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007644D5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076453F | `NoContent_Screen` | Known | Screen layout |
| 0x00764553 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007645BA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00764624 | `NoContent_Screen` | Known | Screen layout |
| 0x00764638 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007646A5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00764717 | `NoContent_Screen` | Known | Screen layout |
| 0x0076472B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00764793 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007647FC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00764817 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076487D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00764899 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00764978 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00764991 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007649F2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00764A06 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00764A60 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00764A7C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00764AE3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00764AFA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00764C6B | `Radio_Screen` | Known | Screen layout |
| 0x00764C7B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00764CDC | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00764D5F | `LockediPod_Screen` | Known | Screen layout |
| 0x00764DE7 | `Lock_Screen` | Known | Screen layout |
| 0x00764DF6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00764E59 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00764EBB | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00764ED7 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00764F49 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00764F68 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00764FD0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00764FEA | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00765052 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076506F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007650DB | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00765145 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076515F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007651CF | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00765242 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007652B3 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00765322 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076538E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007653A9 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076541E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00765485 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007654E7 | `Photos_Screen` | Known | Screen layout |
| 0x0076554B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00765569 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007655DB | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007655F8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076565E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00765679 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007656E2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007656FF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00765776 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076579A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00765808 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00765823 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007658E0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007658FC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076596A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00765987 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007659F2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00765A12 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00765A89 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00765AA5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00765B15 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00765B34 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00765BA0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00765BB4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00765C2D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00765CA1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00765D11 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00765D78 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00765DE0 | `NoContent_Screen` | Known | Screen layout |
| 0x00765DF4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00765E58 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00765EBF | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00765ED9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00765F47 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00765FB9 | `NoContent_Screen` | Known | Screen layout |
| 0x00765FCD | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00766037 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007660A0 | `No_Photos_Screen` | Known | Screen layout |
| 0x007660B4 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076611A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00766188 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007661F5 | `NoContent_Screen` | Known | Screen layout |
| 0x00766209 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00766271 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007662DB | `NoContent_Screen` | Known | Screen layout |
| 0x007662EF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00766356 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007663C0 | `NoContent_Screen` | Known | Screen layout |
| 0x007663D4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00766441 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007664B3 | `NoContent_Screen` | Known | Screen layout |
| 0x007664C7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076652F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00766598 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007665B3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00766619 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00766635 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00766714 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076672D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076678E | `FirstBoot_Screen` | Known | Screen layout |
| 0x007667A2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007667FC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00766818 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076687F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00766896 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00766A07 | `Radio_Screen` | Known | Screen layout |
| 0x00766A17 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00766A78 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00766AFB | `LockediPod_Screen` | Known | Screen layout |
| 0x00766B83 | `Lock_Screen` | Known | Screen layout |
| 0x00766B92 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00766BF5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00766C57 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00766C73 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00766CE5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00766D04 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00766D6C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00766D86 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00766DEE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00766E0B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00766E77 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00766EE1 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00766EFB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00766F6B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00766FDE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076704F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007670BE | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076712A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00767145 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007671BA | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00767221 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00767283 | `Photos_Screen` | Known | Screen layout |
| 0x007672E7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00767305 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00767377 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00767394 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007673FA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00767415 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076747E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076749B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00767512 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00767536 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007675A4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007675BF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076767C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00767698 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00767706 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00767723 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076778E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007677AE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00767825 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00767841 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007678B1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007678D0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076793C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00767950 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007679C9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00767A3D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00767AAD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00767B14 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00767B7C | `NoContent_Screen` | Known | Screen layout |
| 0x00767B90 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00767BF4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00767C5B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00767C75 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00767CE3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00767D55 | `NoContent_Screen` | Known | Screen layout |
| 0x00767D69 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00767DD3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00767E3C | `No_Photos_Screen` | Known | Screen layout |
| 0x00767E50 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00767EB6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00767F24 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00767F91 | `NoContent_Screen` | Known | Screen layout |
| 0x00767FA5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076800D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00768077 | `NoContent_Screen` | Known | Screen layout |
| 0x0076808B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007680F2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076815C | `NoContent_Screen` | Known | Screen layout |
| 0x00768170 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007681DD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076824F | `NoContent_Screen` | Known | Screen layout |
| 0x00768263 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007682CB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00768334 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076834F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007683B5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007683D1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007684B0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007684C9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076852A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076853E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00768598 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007685B4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076861B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00768632 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007687A3 | `Radio_Screen` | Known | Screen layout |
| 0x007687B3 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00768814 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00768897 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076891F | `Lock_Screen` | Known | Screen layout |
| 0x0076892E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00768991 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007689F3 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00768A0F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00768A81 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00768AA0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00768B08 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00768B22 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00768B8A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00768BA7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00768C13 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00768C7D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00768C97 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00768D07 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00768D7A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00768DEB | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00768E5A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00768EC6 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00768EE1 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00768F56 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00768FBD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076901F | `Photos_Screen` | Known | Screen layout |
| 0x00769083 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007690A1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00769113 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00769130 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00769196 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007691B1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076921A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00769237 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007692AE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007692D2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00769340 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076935B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00769418 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00769434 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007694A2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007694BF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076952A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076954A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007695C1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007695DD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076964D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076966C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007696D8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007696EC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00769765 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007697D9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00769849 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007698B0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00769918 | `NoContent_Screen` | Known | Screen layout |
| 0x0076992C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00769990 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007699F7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00769A11 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00769A7F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00769AF1 | `NoContent_Screen` | Known | Screen layout |
| 0x00769B05 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00769B6F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00769BD8 | `No_Photos_Screen` | Known | Screen layout |
| 0x00769BEC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00769C52 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00769CC0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00769D2D | `NoContent_Screen` | Known | Screen layout |
| 0x00769D41 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00769DA9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00769E13 | `NoContent_Screen` | Known | Screen layout |
| 0x00769E27 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00769E8E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00769EF8 | `NoContent_Screen` | Known | Screen layout |
| 0x00769F0C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00769F79 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00769FEB | `NoContent_Screen` | Known | Screen layout |
| 0x00769FFF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076A067 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076A0D0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076A0EB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076A151 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076A16D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076A24C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076A265 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076A2C6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076A2DA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076A334 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076A350 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076A3B7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076A3CE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076A53F | `Radio_Screen` | Known | Screen layout |
| 0x0076A54F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076A5B0 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076A633 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076A6BB | `Lock_Screen` | Known | Screen layout |
| 0x0076A6CA | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076A72D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076A78F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076A7AB | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076A81D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076A83C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076A8A4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076A8BE | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076A926 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076A943 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076A9AF | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076AA19 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076AA33 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076AAA3 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076AB16 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076AB87 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076ABF6 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076AC62 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076AC7D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076ACF2 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076AD59 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076ADBB | `Photos_Screen` | Known | Screen layout |
| 0x0076AE1F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076AE3D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076AEAF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076AECC | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076AF32 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076AF4D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076AFB6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076AFD3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076B04A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076B06E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076B0DC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076B0F7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076B1B4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076B1D0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076B23E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076B25B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076B2C6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076B2E6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076B35D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076B379 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076B3E9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076B408 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076B474 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076B488 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076B501 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076B575 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076B5E5 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076B64C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076B6B4 | `NoContent_Screen` | Known | Screen layout |
| 0x0076B6C8 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076B72C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076B793 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076B7AD | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076B81B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076B88D | `NoContent_Screen` | Known | Screen layout |
| 0x0076B8A1 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076B90B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076B974 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076B988 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076B9EE | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076BA5C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076BAC9 | `NoContent_Screen` | Known | Screen layout |
| 0x0076BADD | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076BB45 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076BBAF | `NoContent_Screen` | Known | Screen layout |
| 0x0076BBC3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076BC2A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076BC94 | `NoContent_Screen` | Known | Screen layout |
| 0x0076BCA8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076BD15 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076BD87 | `NoContent_Screen` | Known | Screen layout |
| 0x0076BD9B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076BE03 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076BE6C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076BE87 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076BEED | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076BF09 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076BFE8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076C001 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076C062 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076C076 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076C0D0 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076C0EC | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076C153 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076C16A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076C2DB | `Radio_Screen` | Known | Screen layout |
| 0x0076C2EB | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076C34C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076C3CF | `LockediPod_Screen` | Known | Screen layout |
| 0x0076C457 | `Lock_Screen` | Known | Screen layout |
| 0x0076C466 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076C4C9 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076C52B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076C547 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076C5B9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076C5D8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076C640 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076C65A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076C6C2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076C6DF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076C74B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076C7B5 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076C7CF | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076C83F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076C8B2 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076C923 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076C992 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076C9FE | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076CA19 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076CA8E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076CAF5 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076CB57 | `Photos_Screen` | Known | Screen layout |
| 0x0076CBBB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076CBD9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076CC4B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076CC68 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076CCCE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076CCE9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076CD52 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076CD6F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076CDE6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076CE0A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076CE78 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076CE93 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076CF50 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076CF6C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076CFDA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076CFF7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076D062 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076D082 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076D0F9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076D115 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076D185 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076D1A4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076D210 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076D224 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076D29D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076D311 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076D381 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076D3E8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076D450 | `NoContent_Screen` | Known | Screen layout |
| 0x0076D464 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076D4C8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076D52F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076D549 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076D5B7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076D629 | `NoContent_Screen` | Known | Screen layout |
| 0x0076D63D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076D6A7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076D710 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076D724 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076D78A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076D7F8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076D865 | `NoContent_Screen` | Known | Screen layout |
| 0x0076D879 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076D8E1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076D94B | `NoContent_Screen` | Known | Screen layout |
| 0x0076D95F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076D9C6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076DA30 | `NoContent_Screen` | Known | Screen layout |
| 0x0076DA44 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076DAB1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076DB23 | `NoContent_Screen` | Known | Screen layout |
| 0x0076DB37 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076DB9F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076DC08 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076DC23 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076DC89 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076DCA5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076DD84 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076DD9D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076DDFE | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076DE12 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076DE6C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076DE88 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076DEEF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076DF06 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076E077 | `Radio_Screen` | Known | Screen layout |
| 0x0076E087 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076E0E8 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076E16B | `LockediPod_Screen` | Known | Screen layout |
| 0x0076E1F3 | `Lock_Screen` | Known | Screen layout |
| 0x0076E202 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076E265 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076E2C7 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076E2E3 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076E355 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076E374 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076E3DC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076E3F6 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076E45E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076E47B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076E4E7 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076E551 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076E56B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076E5DB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076E64E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076E6BF | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076E72E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076E79A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076E7B5 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076E82A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076E891 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076E8F3 | `Photos_Screen` | Known | Screen layout |
| 0x0076E957 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076E975 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076E9E7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076EA04 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076EA6A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076EA85 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076EAEE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076EB0B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076EB82 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076EBA6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076EC14 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076EC2F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076ECEC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076ED08 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076ED76 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076ED93 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076EDFE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076EE1E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076EE95 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076EEB1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076EF21 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076EF40 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076EFAC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076EFC0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076F039 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076F0AD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076F11D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076F184 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076F1EC | `NoContent_Screen` | Known | Screen layout |
| 0x0076F200 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076F264 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076F2CB | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076F2E5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076F353 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076F3C5 | `NoContent_Screen` | Known | Screen layout |
| 0x0076F3D9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076F443 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076F4AC | `No_Photos_Screen` | Known | Screen layout |
| 0x0076F4C0 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076F526 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076F594 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076F601 | `NoContent_Screen` | Known | Screen layout |
| 0x0076F615 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076F67D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076F6E7 | `NoContent_Screen` | Known | Screen layout |
| 0x0076F6FB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076F762 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076F7CC | `NoContent_Screen` | Known | Screen layout |
| 0x0076F7E0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076F84D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076F8BF | `NoContent_Screen` | Known | Screen layout |
| 0x0076F8D3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076F93B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076F9A4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076F9BF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076FA25 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076FA41 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076FB20 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076FB39 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076FB9A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076FBAE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076FC08 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076FC24 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076FC8B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076FCA2 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076FE13 | `Radio_Screen` | Known | Screen layout |
| 0x0076FE23 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076FE84 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076FF07 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076FF8F | `Lock_Screen` | Known | Screen layout |
| 0x0076FF9E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00770001 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00770063 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077007F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007700F1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00770110 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00770178 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00770192 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007701FA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00770217 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00770283 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007702ED | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00770307 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00770377 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007703EA | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077045B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007704CA | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00770536 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00770551 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007705C6 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077062D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077068F | `Photos_Screen` | Known | Screen layout |
| 0x007706F3 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00770711 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00770783 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007707A0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00770806 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00770821 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077088A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007708A7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077091E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00770942 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007709B0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007709CB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00770A88 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00770AA4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00770B12 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00770B2F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00770B9A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00770BBA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00770C31 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00770C4D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00770CBD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00770CDC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00770D48 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00770D5C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00770DD5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00770E49 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00770EB9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00770F20 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00770F88 | `NoContent_Screen` | Known | Screen layout |
| 0x00770F9C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00771000 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00771067 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00771081 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007710EF | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00771161 | `NoContent_Screen` | Known | Screen layout |
| 0x00771175 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007711DF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00771248 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077125C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007712C2 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00771330 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077139D | `NoContent_Screen` | Known | Screen layout |
| 0x007713B1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00771419 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00771483 | `NoContent_Screen` | Known | Screen layout |
| 0x00771497 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007714FE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00771568 | `NoContent_Screen` | Known | Screen layout |
| 0x0077157C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007715E9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077165B | `NoContent_Screen` | Known | Screen layout |
| 0x0077166F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007716D7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00771740 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077175B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007717C1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007717DD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007718BC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007718D5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00771936 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077194A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007719A4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007719C0 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00771A27 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00771A3E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00771BAF | `Radio_Screen` | Known | Screen layout |
| 0x00771BBF | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00771C20 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00771CA3 | `LockediPod_Screen` | Known | Screen layout |
| 0x00771D2B | `Lock_Screen` | Known | Screen layout |
| 0x00771D3A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00771D9D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00771DFF | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00771E1B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00771E8D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00771EAC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00771F14 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00771F2E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00771F96 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00771FB3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077201F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00772089 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007720A3 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00772113 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00772186 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007721F7 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00772266 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007722D2 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007722ED | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00772362 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007723C9 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077242B | `Photos_Screen` | Known | Screen layout |
| 0x0077248F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007724AD | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077251F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077253C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007725A2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007725BD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00772626 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00772643 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007726BA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007726DE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077274C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00772767 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00772824 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00772840 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007728AE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007728CB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00772936 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00772956 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007729CD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007729E9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00772A59 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00772A78 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00772AE4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00772AF8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00772B71 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00772BE5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00772C55 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00772CBC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00772D24 | `NoContent_Screen` | Known | Screen layout |
| 0x00772D38 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00772D9C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00772E03 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00772E1D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00772E8B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00772EFD | `NoContent_Screen` | Known | Screen layout |
| 0x00772F11 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00772F7B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00772FE4 | `No_Photos_Screen` | Known | Screen layout |
| 0x00772FF8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077305E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007730CC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00773139 | `NoContent_Screen` | Known | Screen layout |
| 0x0077314D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007731B5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077321F | `NoContent_Screen` | Known | Screen layout |
| 0x00773233 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077329A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00773304 | `NoContent_Screen` | Known | Screen layout |
| 0x00773318 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00773385 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007733F7 | `NoContent_Screen` | Known | Screen layout |
| 0x0077340B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00773473 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007734DC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007734F7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077355D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00773579 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00773658 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00773671 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007736D2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007736E6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00773740 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077375C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007737C3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007737DA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077394B | `Radio_Screen` | Known | Screen layout |
| 0x0077395B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007739BC | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00773A3F | `LockediPod_Screen` | Known | Screen layout |
| 0x00773AC7 | `Lock_Screen` | Known | Screen layout |
| 0x00773AD6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00773B39 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00773B9B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00773BB7 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00773C29 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00773C48 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00773CB0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00773CCA | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00773D32 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00773D4F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00773DBB | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00773E25 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00773E3F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00773EAF | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00773F22 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00773F93 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00774002 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077406E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00774089 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007740FE | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00774165 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007741C7 | `Photos_Screen` | Known | Screen layout |
| 0x0077422B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00774249 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007742BB | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007742D8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077433E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00774359 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007743C2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007743DF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00774456 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077447A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007744E8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00774503 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007745C0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007745DC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077464A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00774667 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007746D2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007746F2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00774769 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00774785 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007747F5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00774814 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00774880 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00774894 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077490D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00774981 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007749F1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00774A58 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00774AC0 | `NoContent_Screen` | Known | Screen layout |
| 0x00774AD4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00774B38 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00774B9F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00774BB9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00774C27 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00774C99 | `NoContent_Screen` | Known | Screen layout |
| 0x00774CAD | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00774D17 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00774D80 | `No_Photos_Screen` | Known | Screen layout |
| 0x00774D94 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00774DFA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00774E68 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00774ED5 | `NoContent_Screen` | Known | Screen layout |
| 0x00774EE9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00774F51 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00774FBB | `NoContent_Screen` | Known | Screen layout |
| 0x00774FCF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00775036 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007750A0 | `NoContent_Screen` | Known | Screen layout |
| 0x007750B4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00775121 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00775193 | `NoContent_Screen` | Known | Screen layout |
| 0x007751A7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077520F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00775278 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00775293 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007752F9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00775315 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007753F4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077540D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077546E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00775482 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007754DC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007754F8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077555F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00775576 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007756E7 | `Radio_Screen` | Known | Screen layout |
| 0x007756F7 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00775758 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007757DB | `LockediPod_Screen` | Known | Screen layout |
| 0x00775863 | `Lock_Screen` | Known | Screen layout |
| 0x00775872 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007758D5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00775937 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00775953 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007759C5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007759E4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00775A4C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00775A66 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00775ACE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00775AEB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00775B57 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00775BC1 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00775BDB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00775C4B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00775CBE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00775D2F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00775D9E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00775E0A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00775E25 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00775E9A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00775F01 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00775F63 | `Photos_Screen` | Known | Screen layout |
| 0x00775FC7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00775FE5 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00776057 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00776074 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007760DA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007760F5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077615E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077617B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007761F2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00776216 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00776284 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077629F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077635C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00776378 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007763E6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00776403 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077646E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077648E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00776505 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00776521 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00776591 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007765B0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077661C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00776630 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007766A9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077671D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077678D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007767F4 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077685C | `NoContent_Screen` | Known | Screen layout |
| 0x00776870 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007768D4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077693B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00776955 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007769C3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00776A35 | `NoContent_Screen` | Known | Screen layout |
| 0x00776A49 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00776AB3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00776B1C | `No_Photos_Screen` | Known | Screen layout |
| 0x00776B30 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00776B96 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00776C04 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00776C71 | `NoContent_Screen` | Known | Screen layout |
| 0x00776C85 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00776CED | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00776D57 | `NoContent_Screen` | Known | Screen layout |
| 0x00776D6B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00776DD2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00776E3C | `NoContent_Screen` | Known | Screen layout |
| 0x00776E50 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00776EBD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00776F2F | `NoContent_Screen` | Known | Screen layout |
| 0x00776F43 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00776FAB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00777014 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077702F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00777095 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007770B1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00777190 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007771A9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077720A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077721E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00777278 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00777294 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007772FB | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00777312 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00777483 | `Radio_Screen` | Known | Screen layout |
| 0x00777493 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007774F4 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00777577 | `LockediPod_Screen` | Known | Screen layout |
| 0x007775FF | `Lock_Screen` | Known | Screen layout |
| 0x0077760E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00777671 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007776D3 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007776EF | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00777761 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00777780 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007777E8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00777802 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077786A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00777887 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007778F3 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077795D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00777977 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007779E7 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00777A5A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00777ACB | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00777B3A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00777BA6 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00777BC1 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00777C36 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00777C9D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00777CFF | `Photos_Screen` | Known | Screen layout |
| 0x00777D63 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00777D81 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00777DF3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00777E10 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00777E76 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00777E91 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00777EFA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00777F17 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00777F8E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00777FB2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00778020 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077803B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007780F8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00778114 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00778182 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077819F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077820A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077822A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007782A1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007782BD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077832D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077834C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007783B8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007783CC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00778445 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007784B9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00778529 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00778590 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007785F8 | `NoContent_Screen` | Known | Screen layout |
| 0x0077860C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00778670 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007786D7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007786F1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077875F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007787D1 | `NoContent_Screen` | Known | Screen layout |
| 0x007787E5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077884F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007788B8 | `No_Photos_Screen` | Known | Screen layout |
| 0x007788CC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00778932 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007789A0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00778A0D | `NoContent_Screen` | Known | Screen layout |
| 0x00778A21 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00778A89 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00778AF3 | `NoContent_Screen` | Known | Screen layout |
| 0x00778B07 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00778B6E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00778BD8 | `NoContent_Screen` | Known | Screen layout |
| 0x00778BEC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00778C59 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00778CCB | `NoContent_Screen` | Known | Screen layout |
| 0x00778CDF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00778D47 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00778DB0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00778DCB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00778E31 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00778E4D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00778F2C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00778F45 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00778FA6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00778FBA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00779014 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00779030 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00779097 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007790AE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077921F | `Radio_Screen` | Known | Screen layout |
| 0x0077922F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00779290 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00779313 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077939B | `Lock_Screen` | Known | Screen layout |
| 0x007793AA | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077940D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077946F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077948B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007794FD | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077951C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00779584 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077959E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00779606 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00779623 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077968F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007796F9 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00779713 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00779783 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007797F6 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00779867 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007798D6 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00779942 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077995D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007799D2 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00779A39 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00779A9B | `Photos_Screen` | Known | Screen layout |
| 0x00779AFF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00779B1D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00779B8F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00779BAC | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00779C12 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00779C2D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00779C96 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00779CB3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00779D2A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00779D4E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00779DBC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00779DD7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00779E94 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00779EB0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00779F1E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00779F3B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00779FA6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00779FC6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077A03D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077A059 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077A0C9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077A0E8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077A154 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077A168 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077A1E1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077A255 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077A2C5 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077A32C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077A394 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A3A8 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077A40C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077A473 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077A48D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077A4FB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077A56D | `NoContent_Screen` | Known | Screen layout |
| 0x0077A581 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077A5EB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077A654 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077A668 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077A6CE | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077A73C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077A7A9 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A7BD | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077A825 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077A88F | `NoContent_Screen` | Known | Screen layout |
| 0x0077A8A3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077A90A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077A974 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A988 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077A9F5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077AA67 | `NoContent_Screen` | Known | Screen layout |
| 0x0077AA7B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077AAE3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077AB4C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077AB67 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077ABCD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077ABE9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077ACC8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077ACE1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077AD42 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077AD56 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077ADB0 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077ADCC | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077AE33 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077AE4A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077AFBB | `Radio_Screen` | Known | Screen layout |
| 0x0077AFCB | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077B02C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077B0AF | `LockediPod_Screen` | Known | Screen layout |
| 0x0077B137 | `Lock_Screen` | Known | Screen layout |
| 0x0077B146 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077B1A9 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077B20B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077B227 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077B299 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077B2B8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077B320 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077B33A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077B3A2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077B3BF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077B42B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077B495 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077B4AF | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077B51F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077B592 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077B603 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077B672 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077B6DE | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077B6F9 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077B76E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077B7D5 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077B837 | `Photos_Screen` | Known | Screen layout |
| 0x0077B89B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077B8B9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077B92B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077B948 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077B9AE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077B9C9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077BA32 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077BA4F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077BAC6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077BAEA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077BB58 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077BB73 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077BC30 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077BC4C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077BCBA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077BCD7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077BD42 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077BD62 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077BDD9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077BDF5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077BE65 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077BE84 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077BEF0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077BF04 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077BF7D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077BFF1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077C061 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077C0C8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077C130 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C144 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077C1A8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077C20F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077C229 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077C297 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077C309 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C31D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077C387 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077C3F0 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077C404 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077C46A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077C4D8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077C545 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C559 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077C5C1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077C62B | `NoContent_Screen` | Known | Screen layout |
| 0x0077C63F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077C6A6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077C710 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C724 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077C791 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077C803 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C817 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077C87F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077C8E8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077C903 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077C969 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077C985 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077CA64 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077CA7D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077CADE | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077CAF2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077CB4C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077CB68 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077CBCF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077CBE6 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077CD57 | `Radio_Screen` | Known | Screen layout |
| 0x0077CD67 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077CDC8 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077CE4B | `LockediPod_Screen` | Known | Screen layout |
| 0x0077CED3 | `Lock_Screen` | Known | Screen layout |
| 0x0077CEE2 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077CF45 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077CFA7 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077CFC3 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077D035 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077D054 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077D0BC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077D0D6 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077D13E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077D15B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077D1C7 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077D231 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077D24B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077D2BB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077D32E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077D39F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077D40E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077D47A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077D495 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077D50A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077D571 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077D5D3 | `Photos_Screen` | Known | Screen layout |
| 0x0077D637 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077D655 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077D6C7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077D6E4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077D74A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077D765 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077D7CE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077D7EB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077D862 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077D886 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077D8F4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077D90F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077D9CC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077D9E8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077DA56 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077DA73 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077DADE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077DAFE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077DB75 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077DB91 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077DC01 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077DC20 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077DC8C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077DCA0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077DD19 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077DD8D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077DDFD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077DE64 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077DECC | `NoContent_Screen` | Known | Screen layout |
| 0x0077DEE0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077DF44 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077DFAB | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077DFC5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077E033 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077E0A5 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E0B9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077E123 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077E18C | `No_Photos_Screen` | Known | Screen layout |
| 0x0077E1A0 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077E206 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077E274 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077E2E1 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E2F5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077E35D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077E3C7 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E3DB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077E442 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077E4AC | `NoContent_Screen` | Known | Screen layout |
| 0x0077E4C0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077E52D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077E59F | `NoContent_Screen` | Known | Screen layout |
| 0x0077E5B3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077E61B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077E684 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077E69F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077E705 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077E721 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077E800 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077E819 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077E87A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077E88E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077E8E8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077E904 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077E96B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077E982 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077EAF3 | `Radio_Screen` | Known | Screen layout |
| 0x0077EB03 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077EB64 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077EBE7 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077EC6F | `Lock_Screen` | Known | Screen layout |
| 0x0077EC7E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077ECE1 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077ED43 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077ED5F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077EDD1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077EDF0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077EE58 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077EE72 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077EEDA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077EEF7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077EF63 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077EFCD | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077EFE7 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077F057 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077F0CA | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077F13B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077F1AA | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077F216 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077F231 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077F2A6 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077F30D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077F36F | `Photos_Screen` | Known | Screen layout |
| 0x0077F3D3 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077F3F1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077F463 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077F480 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077F4E6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077F501 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077F56A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077F587 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077F5FE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077F622 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077F690 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077F6AB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077F768 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077F784 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077F7F2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077F80F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077F87A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077F89A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077F911 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077F92D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077F99D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077F9BC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077FA28 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077FA3C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077FAB5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077FB29 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077FB99 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077FC00 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077FC68 | `NoContent_Screen` | Known | Screen layout |
| 0x0077FC7C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077FCE0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077FD47 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077FD61 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077FDCF | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077FE41 | `NoContent_Screen` | Known | Screen layout |
| 0x0077FE55 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077FEBF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077FF28 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077FF3C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077FFA2 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00780010 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078007D | `NoContent_Screen` | Known | Screen layout |
| 0x00780091 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007800F9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00780163 | `NoContent_Screen` | Known | Screen layout |
| 0x00780177 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007801DE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00780248 | `NoContent_Screen` | Known | Screen layout |
| 0x0078025C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007802C9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078033B | `NoContent_Screen` | Known | Screen layout |
| 0x0078034F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007803B7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00780420 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078043B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007804A1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007804BD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078059C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007805B5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00780616 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078062A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00780684 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007806A0 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00780707 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078071E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078088F | `Radio_Screen` | Known | Screen layout |
| 0x0078089F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00780900 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00780983 | `LockediPod_Screen` | Known | Screen layout |
| 0x00780A0B | `Lock_Screen` | Known | Screen layout |
| 0x00780A1A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00780A7D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00780ADF | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00780AFB | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00780B6D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00780B8C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00780BF4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00780C0E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00780C76 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00780C93 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00780CFF | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00780D69 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00780D83 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00780DF3 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00780E66 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00780ED7 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00780F46 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00780FB2 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00780FCD | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00781042 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007810A9 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078110B | `Photos_Screen` | Known | Screen layout |
| 0x0078116F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078118D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007811FF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078121C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00781282 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078129D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00781306 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00781323 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078139A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007813BE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078142C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00781447 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00781504 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00781520 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078158E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007815AB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00781616 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00781636 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007816AD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007816C9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00781739 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00781758 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007817C4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007817D8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00781851 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007818C5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00781935 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0078199C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00781A04 | `NoContent_Screen` | Known | Screen layout |
| 0x00781A18 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00781A7C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00781AE3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00781AFD | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00781B6B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00781BDD | `NoContent_Screen` | Known | Screen layout |
| 0x00781BF1 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00781C5B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00781CC4 | `No_Photos_Screen` | Known | Screen layout |
| 0x00781CD8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00781D3E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00781DAC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00781E19 | `NoContent_Screen` | Known | Screen layout |
| 0x00781E2D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00781E95 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00781EFF | `NoContent_Screen` | Known | Screen layout |
| 0x00781F13 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00781F7A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00781FE4 | `NoContent_Screen` | Known | Screen layout |
| 0x00781FF8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00782065 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007820D7 | `NoContent_Screen` | Known | Screen layout |
| 0x007820EB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00782153 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007821BC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007821D7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078223D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00782259 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00782338 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00782351 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007823B2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007823C6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00782420 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078243C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007824A3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007824BA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078262B | `Radio_Screen` | Known | Screen layout |
| 0x0078263B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078269C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078271F | `LockediPod_Screen` | Known | Screen layout |
| 0x007827A7 | `Lock_Screen` | Known | Screen layout |
| 0x007827B6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00782819 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078287B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00782897 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00782909 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00782928 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00782990 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007829AA | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00782A12 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00782A2F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00782A9B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00782B05 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00782B1F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00782B8F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00782C02 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00782C73 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00782CE2 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00782D4E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00782D69 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00782DDE | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00782E45 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00782EA7 | `Photos_Screen` | Known | Screen layout |
| 0x00782F0B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00782F29 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00782F9B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00782FB8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078301E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00783039 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007830A2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007830BF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00783136 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078315A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007831C8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007831E3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007832A0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007832BC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078332A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00783347 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007833B2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007833D2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00783449 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00783465 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007834D5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007834F4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00783560 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00783574 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007835ED | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00783661 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007836D1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00783738 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007837A0 | `NoContent_Screen` | Known | Screen layout |
| 0x007837B4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00783818 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078387F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00783899 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00783907 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00783979 | `NoContent_Screen` | Known | Screen layout |
| 0x0078398D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007839F7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00783A60 | `No_Photos_Screen` | Known | Screen layout |
| 0x00783A74 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00783ADA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00783B48 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00783BB5 | `NoContent_Screen` | Known | Screen layout |
| 0x00783BC9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00783C31 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00783C9B | `NoContent_Screen` | Known | Screen layout |
| 0x00783CAF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00783D16 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00783D80 | `NoContent_Screen` | Known | Screen layout |
| 0x00783D94 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00783E01 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00783E73 | `NoContent_Screen` | Known | Screen layout |
| 0x00783E87 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00783EEF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00783F58 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00783F73 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00783FD9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00783FF5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007840D4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007840ED | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078414E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00784162 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007841BC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007841D8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078423F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00784256 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007843C7 | `Radio_Screen` | Known | Screen layout |
| 0x007843D7 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00784438 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007844BB | `LockediPod_Screen` | Known | Screen layout |
| 0x00784543 | `Lock_Screen` | Known | Screen layout |
| 0x00784552 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007845B5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00784617 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00784633 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007846A5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007846C4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078472C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00784746 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007847AE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007847CB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00784837 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007848A1 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007848BB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078492B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078499E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00784A0F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00784A7E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00784AEA | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00784B05 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00784B7A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00784BE1 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00784C43 | `Photos_Screen` | Known | Screen layout |
| 0x00784CA7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00784CC5 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00784D37 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00784D54 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00784DBA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00784DD5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00784E3E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00784E5B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00784ED2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00784EF6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00784F64 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00784F7F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00785021 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078503D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007850AB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007850C8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00785133 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00785153 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007851CA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007851E6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00785256 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00785275 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007852E1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007852F5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078536A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007853D5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00785444 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007854B5 | `NoContent_Screen` | Known | Screen layout |
| 0x007854C9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00785538 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007855AB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00785618 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00785681 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007856F1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00785761 | `NoContent_Screen` | Known | Screen layout |
| 0x00785775 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007857D8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078583B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00785857 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007858B9 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007858D5 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078593C | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00785953 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00785A0E | `Radio_Screen` | Known | Screen layout |
| 0x00785A1E | `Radio_Screen_Default` | Known | Screen layout |
| 0x00785A7F | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00785AED | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00785B0C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00785B7A | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00785BDF | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00785BFA | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00785C9D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00785CB9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00785D27 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00785D44 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00785DAF | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00785DCF | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00785E46 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00785E62 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00785ED2 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00785EF1 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00785F5D | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00785F71 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00785FE6 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00786051 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007860C0 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00786131 | `NoContent_Screen` | Known | Screen layout |
| 0x00786145 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007861B4 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00786227 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00786294 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007862FD | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078636D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007863DD | `NoContent_Screen` | Known | Screen layout |
| 0x007863F1 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00786454 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007864B7 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007864D3 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00786535 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00786551 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007865B8 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007865CF | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078668A | `Radio_Screen` | Known | Screen layout |
| 0x0078669A | `Radio_Screen_Default` | Known | Screen layout |
| 0x007866FB | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00786769 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00786788 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007867F6 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078685B | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00786876 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00786919 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00786935 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007869A3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007869C0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00786A2B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00786A4B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00786AC2 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00786ADE | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00786B4E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00786B6D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00786BD9 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00786BED | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00786C62 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00786CCD | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00786D3C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00786DAD | `NoContent_Screen` | Known | Screen layout |
| 0x00786DC1 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00786E30 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00786EA3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00786F10 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00786F79 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00786FE9 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00787059 | `NoContent_Screen` | Known | Screen layout |
| 0x0078706D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007870D0 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00787133 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078714F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007871B1 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007871CD | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00787234 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078724B | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00787306 | `Radio_Screen` | Known | Screen layout |
| 0x00787316 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00787377 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007873E5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00787404 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00787472 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007874D7 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007874F2 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00787595 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007875B1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078761F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078763C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007876A7 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007876C7 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078773E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078775A | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007877CA | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007877E9 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00787855 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00787869 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007878DE | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00787949 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007879B8 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00787A29 | `NoContent_Screen` | Known | Screen layout |
| 0x00787A3D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00787AAC | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00787B1F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00787B8C | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00787BF5 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00787C65 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00787CD5 | `NoContent_Screen` | Known | Screen layout |
| 0x00787CE9 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00787D4C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00787DAF | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00787DCB | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00787E2D | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00787E49 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00787EB0 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00787EC7 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00787F82 | `Radio_Screen` | Known | Screen layout |
| 0x00787F92 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00787FF3 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00788061 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00788080 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007880EE | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00788153 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078816E | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00788211 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078822D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078829B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007882B8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00788323 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00788343 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007883BA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007883D6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00788446 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00788465 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007884D1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007884E5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078855A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007885C5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00788634 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007886A5 | `NoContent_Screen` | Known | Screen layout |
| 0x007886B9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00788728 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078879B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00788808 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00788871 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007888E1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00788951 | `NoContent_Screen` | Known | Screen layout |
| 0x00788965 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007889C8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00788A2B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00788A47 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00788AA9 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00788AC5 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00788B2C | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00788B43 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00788BFE | `Radio_Screen` | Known | Screen layout |
| 0x00788C0E | `Radio_Screen_Default` | Known | Screen layout |
| 0x00788C6F | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00788CDD | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00788CFC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00788D6A | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00788DCF | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00788DEA | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00788E8D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00788EA9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00788F17 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00788F34 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00788F9F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00788FBF | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00789036 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00789052 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007890C2 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007890E1 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078914D | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00789161 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007891D6 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00789241 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007892B0 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00789321 | `NoContent_Screen` | Known | Screen layout |
| 0x00789335 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007893A4 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00789417 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00789484 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007894ED | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078955D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007895CD | `NoContent_Screen` | Known | Screen layout |
| 0x007895E1 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00789644 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007896A7 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007896C3 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00789725 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00789741 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007897A8 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007897BF | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078987A | `Radio_Screen` | Known | Screen layout |
| 0x0078988A | `Radio_Screen_Default` | Known | Screen layout |
| 0x007898EB | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00789959 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00789978 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007899E6 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00789A4B | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00789A66 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00789B09 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00789B25 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00789B93 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00789BB0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00789C1B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00789C3B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00789CB2 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00789CCE | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00789D3E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00789D5D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00789DC9 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00789DDD | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00789E52 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00789EBD | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00789F2C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00789F9D | `NoContent_Screen` | Known | Screen layout |
| 0x00789FB1 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078A020 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078A093 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078A100 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078A169 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078A1D9 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078A249 | `NoContent_Screen` | Known | Screen layout |
| 0x0078A25D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078A2C0 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078A323 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078A33F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078A3A1 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078A3BD | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078A424 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078A43B | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078A4F6 | `Radio_Screen` | Known | Screen layout |
| 0x0078A506 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078A567 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078A5D5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078A5F4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078A662 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078A6C7 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078A6E2 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078A785 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078A7A1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078A80F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078A82C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078A897 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078A8B7 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078A92E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078A94A | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078A9BA | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078A9D9 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078AA45 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078AA59 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078AACE | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078AB39 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078ABA8 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078AC19 | `NoContent_Screen` | Known | Screen layout |
| 0x0078AC2D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078AC9C | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078AD0F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078AD7C | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078ADE5 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078AE55 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078AEC5 | `NoContent_Screen` | Known | Screen layout |
| 0x0078AED9 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078AF3C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078AF9F | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078AFBB | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078B01D | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078B039 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078B0A0 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078B0B7 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078B172 | `Radio_Screen` | Known | Screen layout |
| 0x0078B182 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078B1E3 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078B251 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078B270 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078B2DE | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078B343 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078B35E | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078B401 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078B41D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078B48B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078B4A8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078B513 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078B533 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078B5AA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078B5C6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078B636 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078B655 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078B6C1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078B6D5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078B74A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078B7B5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078B824 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078B895 | `NoContent_Screen` | Known | Screen layout |
| 0x0078B8A9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078B918 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078B98B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078B9F8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078BA61 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078BAD1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078BB41 | `NoContent_Screen` | Known | Screen layout |
| 0x0078BB55 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078BBB8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078BC1B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078BC37 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078BC99 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078BCB5 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078BD1C | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078BD33 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078BDEE | `Radio_Screen` | Known | Screen layout |
| 0x0078BDFE | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078BE5F | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078BECD | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078BEEC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078BF5A | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078BFBF | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078BFDA | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078C07D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078C099 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078C107 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078C124 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078C18F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078C1AF | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078C226 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078C242 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078C2B2 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078C2D1 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078C33D | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078C351 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078C3C6 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078C431 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078C4A0 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078C511 | `NoContent_Screen` | Known | Screen layout |
| 0x0078C525 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078C594 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078C607 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078C674 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078C6DD | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078C74D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078C7BD | `NoContent_Screen` | Known | Screen layout |
| 0x0078C7D1 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078C834 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078C897 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078C8B3 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078C915 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078C931 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078C998 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078C9AF | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078CA6A | `Radio_Screen` | Known | Screen layout |
| 0x0078CA7A | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078CADB | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078CB49 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078CB68 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078CBD6 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078CC3B | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078CC56 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078CCF9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078CD15 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078CD83 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078CDA0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078CE0B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078CE2B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078CEA2 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078CEBE | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078CF2E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078CF4D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078CFB9 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078CFCD | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078D042 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078D0AD | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078D11C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078D18D | `NoContent_Screen` | Known | Screen layout |
| 0x0078D1A1 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078D210 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078D283 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078D2F0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078D359 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078D3C9 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078D439 | `NoContent_Screen` | Known | Screen layout |
| 0x0078D44D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078D4B0 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078D513 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078D52F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078D591 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078D5AD | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078D614 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078D62B | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078D6E6 | `Radio_Screen` | Known | Screen layout |
| 0x0078D6F6 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078D757 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078D7C5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078D7E4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078D852 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078D8B7 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078D8D2 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078D975 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078D991 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078D9FF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078DA1C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078DA87 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078DAA7 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078DB1E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078DB3A | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078DBAA | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078DBC9 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078DC35 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078DC49 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078DCBE | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078DD29 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078DD98 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078DE09 | `NoContent_Screen` | Known | Screen layout |
| 0x0078DE1D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078DE8C | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078DEFF | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078DF6C | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078DFD5 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078E045 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078E0B5 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E0C9 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078E12C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078E18F | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078E1AB | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078E20D | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078E229 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078E290 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078E2A7 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078E362 | `Radio_Screen` | Known | Screen layout |
| 0x0078E372 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078E3D3 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078E441 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078E460 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078E4CE | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078E533 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078E54E | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078E66E | `MediaLists_AudiobookChapters_Screen)` | Known | Screen layout |
| 0x0078E695 | `MediaLists_AudiobookChapters_Screen_Plane!` | Known | Screen layout |
| 0x0078EC02 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078EC1E | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0078EC8D | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0078ECA6 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0078F00E | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078F02A | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0078F099 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0078F0B2 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0078F3DB | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078F3F7 | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0078F466 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0078F47F | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0078F6AF | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0078F6CA | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x0078F735 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078F750 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x0078F7C3 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0078F7DE | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0078FA07 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0078FA22 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x0078FA8D | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078FAA8 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x0078FB1B | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0078FB36 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0078FD6A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078FD86 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x0078FE01 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078FE1D | `MediaLists_Albums_Screen_WithTrackCount%` | Known | Screen layout |
| 0x0078FE96 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0078FEB1 | `MediaLists_Songs_Screen_Plain1` | Known | Screen layout |
| 0x0078FF2C | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0078FF47 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x007901D5 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007901F2 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00790339 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00790355 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x007903D0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007903EB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00790639 | `MediaLists_PodcastEpisodes_Screen'` | Known | Screen layout |
| 0x0079065E | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00790996 | `MediaLists_TVSeasons_Screen#` | Known | Screen layout |
| 0x007909B5 | `MediaLists_TVSeasons_Screen_Default%` | Known | Screen layout |
| 0x00790A2A | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x00790A4A | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00790BD2 | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x00790BF2 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00790FFC | `MediaLists_GeniusPlaylist_Screen(` | Known | Screen layout |
| 0x00791020 | `MediaLists_GeniusPlaylist_Screen_Default` | Known | Screen layout |
| 0x007910EF | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x00791114 | `MediaLists_NestedPlaylists_Screen_Default,` | Known | Screen layout |
| 0x00791196 | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x007911B5 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00791492 | `MediaLists_GeniusPlaylist_Screen(` | Known | Screen layout |
| 0x007914B6 | `MediaLists_GeniusPlaylist_Screen_Default#` | Known | Screen layout |
| 0x0079152E | `Genius_Error_Screen` | Known | Screen layout |
| 0x00791545 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007915BD | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007915D4 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x00791642 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0079165E | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x007916CD | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007916E6 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007917B0 | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x007917D5 | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x0079184D | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x0079186C | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x007918D1 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x00791B1D | `MediaLists_GeniusPlaylist_Screen(` | Known | Screen layout |
| 0x00791B41 | `MediaLists_GeniusPlaylist_Screen_Default"` | Known | Screen layout |
| 0x00791BBA | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x00791C2C | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00791C97 | `Genius_Error_Screen` | Known | Screen layout |
| 0x00791CAE | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x00791D26 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00791D3D | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x00791DAB | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00791DC7 | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x00791E36 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x00791E4F | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x00791F46 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x00792208 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x00792308 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x00792374 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007923DE | `NoContent_Screen` | Known | Screen layout |
| 0x007923F2 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079245C | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007924D0 | `NoContent_Screen` | Known | Screen layout |
| 0x007924E4 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079254F | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007925BB | `NoContent_Screen` | Known | Screen layout |
| 0x007925CF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00792636 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007926A2 | `NoContent_Screen` | Known | Screen layout |
| 0x007926B6 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00792723 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00792797 | `NoContent_Screen` | Known | Screen layout |
| 0x007927AB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00792813 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00792880 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007928E4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00792900 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0079296C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00792989 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007929F6 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00792ABD | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00792ADA | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00792B51 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00792B75 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00792C2C | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00792C96 | `NoContent_Screen` | Known | Screen layout |
| 0x00792CAA | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00792D14 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00792D88 | `NoContent_Screen` | Known | Screen layout |
| 0x00792D9C | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00792E07 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00792E73 | `NoContent_Screen` | Known | Screen layout |
| 0x00792E87 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00792EEE | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00792F5A | `NoContent_Screen` | Known | Screen layout |
| 0x00792F6E | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00792FDB | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079304F | `NoContent_Screen` | Known | Screen layout |
| 0x00793063 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007930CB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00793138 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079319C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007931B8 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00793224 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00793241 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007932AE | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00793375 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00793392 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00793409 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079342D | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007934E4 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079354E | `NoContent_Screen` | Known | Screen layout |
| 0x00793562 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007935CC | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00793640 | `NoContent_Screen` | Known | Screen layout |
| 0x00793654 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007936BF | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0079372B | `NoContent_Screen` | Known | Screen layout |
| 0x0079373F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007937A6 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00793812 | `NoContent_Screen` | Known | Screen layout |
| 0x00793826 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00793893 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00793907 | `NoContent_Screen` | Known | Screen layout |
| 0x0079391B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00793983 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007939F0 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00793A54 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00793A70 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00793ADC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00793AF9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00793B66 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00793C2D | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00793C4A | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00793CC1 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00793CE5 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00793D9C | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00793E06 | `NoContent_Screen` | Known | Screen layout |
| 0x00793E1A | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00793E84 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00793EF8 | `NoContent_Screen` | Known | Screen layout |
| 0x00793F0C | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00793F77 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00793FE3 | `NoContent_Screen` | Known | Screen layout |
| 0x00793FF7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079405E | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007940CA | `NoContent_Screen` | Known | Screen layout |
| 0x007940DE | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079414B | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007941BF | `NoContent_Screen` | Known | Screen layout |
| 0x007941D3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079423B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007942A8 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079430C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00794328 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00794394 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007943B1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079441E | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007944E5 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00794502 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00794579 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079459D | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00794654 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007946BE | `NoContent_Screen` | Known | Screen layout |
| 0x007946D2 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079473C | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007947B0 | `NoContent_Screen` | Known | Screen layout |
| 0x007947C4 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079482F | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0079489B | `NoContent_Screen` | Known | Screen layout |
| 0x007948AF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00794916 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00794982 | `NoContent_Screen` | Known | Screen layout |
| 0x00794996 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00794A03 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00794A77 | `NoContent_Screen` | Known | Screen layout |
| 0x00794A8B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00794AF3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00794B60 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00794BC4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00794BE0 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00794C4C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00794C69 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00794CD6 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00794D9D | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00794DBA | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00794E31 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00794E55 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00794F0C | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00794F76 | `NoContent_Screen` | Known | Screen layout |
| 0x00794F8A | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00794FF4 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00795068 | `NoContent_Screen` | Known | Screen layout |
| 0x0079507C | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007950E7 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00795153 | `NoContent_Screen` | Known | Screen layout |
| 0x00795167 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007951CE | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079523A | `NoContent_Screen` | Known | Screen layout |
| 0x0079524E | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007952BB | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079532F | `NoContent_Screen` | Known | Screen layout |
| 0x00795343 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007953AB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00795418 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079547C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00795498 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00795504 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00795521 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079558E | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00795655 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00795672 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007956E9 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079570D | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007957C4 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079582E | `NoContent_Screen` | Known | Screen layout |
| 0x00795842 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007958AC | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00795920 | `NoContent_Screen` | Known | Screen layout |
| 0x00795934 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079599F | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00795A0B | `NoContent_Screen` | Known | Screen layout |
| 0x00795A1F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00795A86 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00795AF2 | `NoContent_Screen` | Known | Screen layout |
| 0x00795B06 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00795B73 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00795BE7 | `NoContent_Screen` | Known | Screen layout |
| 0x00795BFB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00795C63 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00795CD0 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00795D34 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00795D50 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00795DBC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00795DD9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00795E46 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00795F0D | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00795F2A | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00795FA1 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00795FC5 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079607C | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007960E6 | `NoContent_Screen` | Known | Screen layout |
| 0x007960FA | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00796164 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007961D8 | `NoContent_Screen` | Known | Screen layout |
| 0x007961EC | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00796257 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007962C3 | `NoContent_Screen` | Known | Screen layout |
| 0x007962D7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079633E | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007963AA | `NoContent_Screen` | Known | Screen layout |
| 0x007963BE | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079642B | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079649F | `NoContent_Screen` | Known | Screen layout |
| 0x007964B3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079651B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00796588 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007965EC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00796608 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00796674 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00796691 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007966FE | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007967C5 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007967E2 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00796859 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079687D | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00796D2C | `Genius_Error_Screen` | Known | Screen layout |
| 0x00796D43 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x00796DBB | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00796DD2 | `Genius_Error_Screen_NoGeniusInfoForTrack"` | Known | Screen layout |
| 0x00796E49 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x00796E62 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x00796FCF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00796FEE | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007972B8 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00797323 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079733A | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007973B2 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007973C9 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x00797437 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00797453 | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x007974C2 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007974DB | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007975A5 | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x007975CA | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x00797642 | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x00797661 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00797BE7 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00797C59 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00797CC4 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00797D29 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00797D93 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00797DFD | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00797E6D | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x00797EE4 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x00797F56 | `Genius_Error_Screen` | Known | Screen layout |
| 0x00797F6D | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x00797FE5 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00797FFC | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079806E | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007980D5 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007980EE | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x00798157 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007981C2 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079822C | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x00798293 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x00798302 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x00798370 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007983D5 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079843D | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007984A8 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x00798513 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079857A | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x00798BD3 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00798C45 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00798CB0 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00798D15 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00798D7F | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00798DE9 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00798E59 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x00798ED0 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x00798F42 | `Genius_Error_Screen` | Known | Screen layout |
| 0x00798F59 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x00798FD1 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00798FE8 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079905A | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007990C1 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007990DA | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x00799143 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007991AE | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x00799218 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079927F | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007992EE | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079935C | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007993C1 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x00799429 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x00799494 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007994FF | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x00799566 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x00799BBD | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00799C2F | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00799C9A | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00799CFF | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00799D69 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00799DD3 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00799E43 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x00799EBA | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x00799F2C | `Genius_Error_Screen` | Known | Screen layout |
| 0x00799F43 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x00799FBB | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00799FD2 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079A044 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079A0AB | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079A0C4 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079A12D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079A198 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079A202 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079A269 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079A2D8 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079A346 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079A3AB | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079A413 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079A47E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079A4E9 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079A550 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079ABA5 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079AC17 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079AC82 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079ACE7 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079AD51 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079ADBB | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079AE2B | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079AEA2 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079AF14 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079AF2B | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079AFA3 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079AFBA | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079B02C | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079B093 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079B0AC | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079B115 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079B180 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079B1EA | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079B251 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079B2C0 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079B32E | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079B393 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079B3FB | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079B466 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079B4D1 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079B538 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079BB75 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079BBE7 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079BC52 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079BCB7 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079BD21 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079BD8B | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079BDFB | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079BE72 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079BEE4 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079BEFB | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079BF73 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079BF8A | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079BFFC | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079C063 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079C07C | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079C0E5 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079C150 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079C1BA | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079C221 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079C290 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079C2FE | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079C363 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079C3CB | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079C436 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079C4A1 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079C508 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079CB45 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079CBB7 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079CC22 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079CC87 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079CCF1 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079CD5B | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079CDCB | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079CE42 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079CEB4 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079CECB | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079CF43 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079CF5A | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079CFCC | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079D033 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079D04C | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079D0B5 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079D120 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079D18A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079D1F1 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079D260 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079D2CE | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079D333 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079D39B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079D406 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079D471 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079D4D8 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079DB52 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079DBC4 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079DC2F | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079DC94 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079DCFE | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079DD68 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079DDD8 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079DE4F | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079DEC1 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079DED8 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079DF50 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079DF67 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079DFD9 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079E040 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079E059 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079E0C2 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079E12D | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079E197 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079E1FE | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079E26D | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079E2DB | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079E340 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079E3A8 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079E413 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079E47E | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079E4E5 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079EB44 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079EBB6 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079EC21 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079EC86 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079ECF0 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079ED5A | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079EDCA | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079EE41 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079EEB3 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079EECA | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079EF42 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079EF59 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079EFCB | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079F032 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079F04B | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079F0B4 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079F11F | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079F189 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079F1F0 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079F25F | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079F2CD | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079F332 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079F39A | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079F405 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079F470 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079F4D7 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079FB20 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079FB92 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079FBFD | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079FC62 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079FCCC | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079FD36 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079FDA6 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079FE1D | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079FE8F | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079FEA6 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079FF1E | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079FF35 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079FFA7 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A000E | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A0027 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A0090 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A00FB | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A0165 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A01CC | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A023B | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A02A9 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A030E | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A0376 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A03E1 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A044C | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A04B3 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A0AFC | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A0B6E | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A0BD9 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A0C3E | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A0CA8 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A0D12 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A0D82 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A0DF9 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A0E6B | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A0E82 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A0EFA | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A0F11 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A0F83 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A0FEA | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A1003 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A106C | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A10D7 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A1141 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A11A8 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A1217 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A1285 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A12EA | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A1352 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A13BD | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A1428 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A148F | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A1AD9 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A1B4B | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A1BB6 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A1C1B | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A1C85 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A1CEF | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A1D5F | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A1DD6 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A1E48 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A1E5F | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A1ED7 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A1EEE | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A1F60 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A1FC7 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A1FE0 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A2049 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A20B4 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A211E | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A2185 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A21F4 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A2262 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A22C7 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A232F | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A239A | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A2405 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A246C | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A2ADB | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A2B4D | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A2BB8 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A2C1D | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A2C87 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A2CF1 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A2D61 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A2DD8 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A2E4A | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A2E61 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A2ED9 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A2EF0 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A2F62 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A2FC9 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A2FE2 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A304B | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A30B6 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A3120 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A3187 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A31F6 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A3264 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A32C9 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A3331 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A339C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A3407 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A346E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A3AEB | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A3B5D | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A3BC8 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A3C2D | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A3C97 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A3D01 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A3D71 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A3DE8 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A3E5A | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A3E71 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A3EE9 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A3F00 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A3F72 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A3FD9 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A3FF2 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A405B | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A40C6 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A4130 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A4197 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A4206 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A4274 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A42D9 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A4341 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A43AC | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A4417 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A447E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A4ADB | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A4B4D | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A4BB8 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A4C1D | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A4C87 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A4CF1 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A4D61 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A4DD8 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A4E4A | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A4E61 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A4ED9 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A4EF0 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A4F62 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A4FC9 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A4FE2 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A504B | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A50B6 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A5120 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A5187 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A51F6 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A5264 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A52C9 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A5331 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A539C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A5407 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A546E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A5ABF | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A5B31 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A5B9C | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A5C01 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A5C6B | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A5CD5 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A5D45 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A5DBC | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A5E2E | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A5E45 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A5EBD | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A5ED4 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A5F46 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A5FAD | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A5FC6 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A602F | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A609A | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A6104 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A616B | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A61DA | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A6248 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A62AD | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A6315 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A6380 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A63EB | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A6452 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A6A91 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A6B03 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A6B6E | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A6BD3 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A6C3D | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A6CA7 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A6D17 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A6D8E | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A6E00 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A6E17 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A6E8F | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A6EA6 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A6F18 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A6F7F | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A6F98 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A7001 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A706C | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A70D6 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A713D | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A71AC | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A721A | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A727F | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A72E7 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A7352 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A73BD | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A7424 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A7A5A | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A7ACC | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A7B37 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A7B9C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A7C06 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A7C70 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A7CE0 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A7D57 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A7DC9 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A7DE0 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A7E58 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A7E6F | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A7EE1 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A7F48 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A7F61 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A7FCA | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A8035 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A809F | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A8106 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A8175 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A81E3 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A8248 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A82B0 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A831B | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A8386 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A83ED | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A8A3E | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A8AB0 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A8B1B | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A8B80 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A8BEA | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A8C54 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A8CC4 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A8D3B | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A8DAD | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A8DC4 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A8E3C | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A8E53 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A8EC5 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A8F2C | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A8F45 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A8FAE | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A9019 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A9083 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A90EA | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A9159 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A91C7 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A922C | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A9294 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A92FF | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A936A | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A93D1 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A99D8 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A9A4A | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A9AB5 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A9B1A | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A9B84 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A9BEE | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A9C5E | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A9CD5 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A9D47 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A9D5E | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A9DD6 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A9DED | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A9E5F | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A9EC6 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A9EDF | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A9F48 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A9FB3 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007AA01D | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007AA084 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007AA0F3 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007AA161 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007AA1C6 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007AA22E | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007AA299 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007AA304 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007AA36B | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007AA6BE | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AA735 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AA7B2 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AA824 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AA894 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AA90A | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AA978 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AA9E5 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AAD2A | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AADA1 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AAE1E | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AAE90 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AAF00 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AAF76 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AAFE4 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AB051 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AB3BA | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AB431 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AB4AE | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AB520 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AB590 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AB606 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AB674 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AB6E1 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007ABA4A | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007ABAC1 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007ABB3C | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007ABBAC | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007ABC22 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007ABC90 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007ABCFD | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AC036 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AC0AD | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AC128 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AC198 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AC20E | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AC27C | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AC2E9 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AC620 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AC697 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AC712 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AC782 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AC7F8 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AC866 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AC8D3 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007ACBE3 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007ACC5A | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007ACCD5 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007ACD45 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007ACDBB | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007ACE29 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007ACE96 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AD49A | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007AD4B7 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007AD532 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007AD54B | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007AD5C3 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007AD5DC | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007AD651 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AD667 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007AD6DE | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AD6F4 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007AD76B | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007AD788 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007AD800 | `Notes_List_Screen` | Known | Screen layout |
| 0x007AD815 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007AD9C6 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007AD9E3 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007ADA5E | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007ADA77 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007ADAEF | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007ADB08 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007ADB7D | `Notes_Image_Screen` | Known | Screen layout |
| 0x007ADB93 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007ADC0A | `Notes_Image_Screen` | Known | Screen layout |
| 0x007ADC20 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007ADC97 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007ADCB4 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007ADD2C | `Notes_List_Screen` | Known | Screen layout |
| 0x007ADD41 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007ADF22 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007ADF3F | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007ADFBA | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007ADFD3 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007AE04B | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007AE064 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007AE0D9 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AE0EF | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007AE166 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AE17C | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007AE1F3 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007AE210 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007AE288 | `Notes_List_Screen` | Known | Screen layout |
| 0x007AE29D | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007AE452 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007AE46F | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007AE4EA | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007AE503 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007AE57B | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007AE594 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007AE609 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AE61F | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007AE696 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AE6AC | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007AE723 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007AE740 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007AE7B8 | `Notes_List_Screen` | Known | Screen layout |
| 0x007AE7CD | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007AEAE5 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007AEB8B | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AEC0E | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007AECC6 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x007AED48 | `PhotosSettingsSlideshowMusic_Screen+` | Known | Screen layout |
| 0x007AED6F | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x007AEE55 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x007AF00D | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AF06D | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AF0CA | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007AF0F1 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007AF191 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AF1F1 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AF24E | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007AF275 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007AF510 | `Photos_Screen` | Known | Screen layout |
| 0x007AF65C | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AF6C0 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007AF721 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007AF77E | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007AF7DB | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007AF849 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007AF8A6 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007AFA4C | `Photos_Screen` | Known | Screen layout |
| 0x007AFB98 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AFBFC | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007AFC5D | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007AFCBA | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007AFD17 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007AFD85 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007AFDE2 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007AFF88 | `Photos_Screen` | Known | Screen layout |
| 0x007B00D4 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B0138 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007B0199 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007B01F6 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007B0253 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B02C1 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B031E | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B04C4 | `Photos_Screen` | Known | Screen layout |
| 0x007B0610 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B0674 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007B06D5 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007B0732 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007B078F | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B07FD | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B085A | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B0A00 | `Photos_Screen` | Known | Screen layout |
| 0x007B0B4C | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B0BB0 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007B0C11 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007B0C6E | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007B0CCB | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B0D39 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B0D96 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B0F3C | `Photos_Screen` | Known | Screen layout |
| 0x007B1088 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B10EC | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007B114D | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007B11AA | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007B1207 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B1275 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B12D2 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B1478 | `Photos_Screen` | Known | Screen layout |
| 0x007B15C4 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B162A | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B168C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B16EE | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B1784 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B18A5 | `Photos_Screen` | Known | Screen layout |
| 0x007B193C | `Photos_Screen` | Known | Screen layout |
| 0x007B1A88 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B1AEE | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B1B50 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B1BB2 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B1C48 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B1D69 | `Photos_Screen` | Known | Screen layout |
| 0x007B1E00 | `Photos_Screen` | Known | Screen layout |
| 0x007B1F4C | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B1FB2 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B2014 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B2076 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B210C | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B222D | `Photos_Screen` | Known | Screen layout |
| 0x007B22C4 | `Photos_Screen` | Known | Screen layout |
| 0x007B2410 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B2476 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B24D8 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B253A | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B25D0 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B26F1 | `Photos_Screen` | Known | Screen layout |
| 0x007B2788 | `Photos_Screen` | Known | Screen layout |
| 0x007B28D4 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B293A | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B299C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B29FE | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B2A94 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B2BB5 | `Photos_Screen` | Known | Screen layout |
| 0x007B2DD5 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B2E37 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B2EA5 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B2F0B | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B2F74 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B2FDB | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3040 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B330E | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B3370 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B33DE | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B3444 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B34AD | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3514 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3579 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B384A | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B38AC | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B391A | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B3980 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B39E9 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3A50 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3AB5 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3D29 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007B3D86 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B3DE8 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B3E56 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B3EBC | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B421A | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007B4284 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007B462A | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007B4694 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007B4989 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007B49EC | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B4A51 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4AB9 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B4B1C | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B4B84 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B4BED | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B4C53 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4CB8 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B4D25 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4D95 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007B4E0B | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4E81 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B4EF1 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B4F66 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B4FDD | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007B5051 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007B50C3 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007B513D | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B51B0 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B5222 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B52A6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B52D0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B5357 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B53E4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B5483 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B549D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B5515 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B552F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B5599 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B55B6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B562E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B5658 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B56DF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B576C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B580B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5825 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B589D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B58B7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B5921 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B593E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B59B6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B59E0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B5A67 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B5AF4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B5B93 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5BAD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B5C25 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5C3F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B5CA9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B5CC6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B5D3E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B5D68 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B5DEF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B5E7C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B5F1B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5F35 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B5FAD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5FC7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B6031 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B604E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B60C6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B60F0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B6177 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B6204 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B62A3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B62BD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B6335 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B634F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B63B9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B63D6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B644E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B6478 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B64FF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B658C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B662B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6645 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B66BD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B66D7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B6741 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B675E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B67D6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B6800 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B6887 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B6914 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B69B3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B69CD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B6A45 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6A5F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B6AC9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B6AE6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B6B5E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B6B88 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B6C0F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B6C9C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B6D3B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6D55 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B6DCD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6DE7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B6E51 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B6E6E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B6EE6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B6F10 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B6F97 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B7024 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B70C3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B70DD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B7155 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B716F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B71D9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B71F6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B726E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B7298 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B731F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B73AC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B744B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7465 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B74DD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B74F7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B7561 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B757E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B75F6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B7620 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B76A7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B7734 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B77D3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B77ED | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B7865 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B787F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B78E9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B7906 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B797E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B79A8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B7A2F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B7ABC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B7B5B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7B75 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B7BED | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7C07 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B7C71 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B7C8E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B7D06 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B7D30 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B7DB7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B7E44 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B7EE3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7EFD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B7F75 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7F8F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B7FF9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B8016 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B808E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B80B8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B813F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B81CC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B826B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8285 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B82FD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8317 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B8381 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B839E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B8416 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B8440 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B84C7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B8554 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B85F3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B860D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B8685 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B869F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B8709 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B8726 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B879E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B87C8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B884F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B88DC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B897B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8995 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B8A0D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8A27 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B8A91 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B8AAE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B8B26 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B8B50 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B8BD7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B8C64 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B8D03 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8D1D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B8D95 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8DAF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B8E19 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B8E36 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B8EAE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B8ED8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B8F5F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B8FEC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B908B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B90A5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B911D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B9137 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B91A1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B91BE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B9236 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B9260 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B92E7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B9374 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B9413 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B942D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B94A5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B94BF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B9529 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B9546 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B95CD | `SettingsMenus_ResetMainMenu_Screen9` | Known | Screen layout |
| 0x007B969D | `SettingsMenus_ResetMusicMenu_Screen:` | Known | Screen layout |
| 0x007B9751 | `SettingsMenus_Main_Screen,` | Known | Screen layout |
| 0x007B97C3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B97DD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B9855 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B986F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B9BAA | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007B9C10 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007B9C6D | `Extras_Screen` | Known | Screen layout |
| 0x007B9CC1 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x007B9D9F | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x007B9E0D | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007B9EAB | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x007B9EC4 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x007B9F2C | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007B9F9E | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x007B9FB7 | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x007BA01A | `DemoMode_Screen` | Known | Screen layout |
| 0x007BA02D | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x007BA09A | `Debug_TestList_Screen` | Known | Screen layout |
| 0x007BA0B3 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x007BA126 | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x007BA141 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x007BA251 | `VoiceMemos_Context_PlayDelete_Screen)` | Known | Screen layout |
| 0x007BA279 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x007BA3D2 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007BA441 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007BA52D | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007BA5F1 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007BA613 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007BA67F | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007BA6A1 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007BA81E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BA83A | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007BA901 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007BA91C | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007BA97F | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007BA9E2 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007BAA79 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BAA95 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007BAB5C | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007BAB77 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007BABDA | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007BAC3D | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007BACD5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BACF1 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007BADB8 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007BADD3 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007BAE36 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007BAE99 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007BAF16 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007BAF81 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007BAFED | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007BB05F | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007BB0CC | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007BB137 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x007BB1A3 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007BB20B | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007BB277 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007BB2EB | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007BB359 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x007BB3D2 | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x007D7C30 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x007D7CB5 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x007D7FA2 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x0098B67B | `VoiceMemos_Screen_InsufficientDiskSpace2` | Known | Screen layout |
| 0x0098CEFF | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0098CF17 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0098CF35 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x0098D041 | `NowPlaying_Screen_Music` | Known | Screen layout |
| 0x0098D06D | `MainMenu_Music_Screen_NoMusic` | Known | Screen layout |
| 0x0098D08B | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0098D0A9 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0098D1AA | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x0098D25E | `DeleteRental_Confirmation_Screen_Unwatched` | Known | Screen layout |
| 0x0098D2B4 | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x0098D300 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0098D402 | `NowPlaying_Screen_ExtraInfoLoadFailed` | Known | Screen layout |
| 0x0098D45D | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x0098D476 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x0098D494 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x0098D4C3 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x0098D4FB | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x0098D932 | `Video_Settings_TV_Screen_Standard` | Known | Screen layout |
| 0x0098D964 | `Search_Main_Screen_WithKeyboard` | Known | Screen layout |
| 0x0098D984 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x0098D9C9 | `VoiceMemos_Screen_InsufficientDiskSpace` | Known | Screen layout |
| 0x0098DA8D | `Video_Settings_TV_Screen_Wide` | Known | Screen layout |
| 0x0098DAD5 | `CoverFlow_Screen_Backside` | Known | Screen layout |
| 0x0099058F | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x00990794 | `NowPlaying_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x009907B9 | `Slideshow_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x00990889 | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x009908A3 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x00990936 | `RentalDeleted_Screen_Title` | Known | Screen layout |
| 0x00990951 | `SingleRentalExpiring_Screen_Title` | Known | Screen layout |
| 0x00990973 | `MultipleRentalsExpiring_Screen_Title` | Known | Screen layout |
| 0x00990998 | `DeleteRental_Screen_Title` | Known | Screen layout |
| 0x00990A3B | `VoiceMemos_DeleteAll_ScreenTitle` | Known | Screen layout |
| 0x00990AD8 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00990B1B | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00990D0C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00990DF5 | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x00990E0E | `Radio_Screen_Volume` | Known | Screen layout |
| 0x00990E22 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x00990E3F | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x00990E5E | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x00990F2A | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x00991080 | `CoverFlow_Screen_BacksideForCapture` | Known | Screen layout |
| 0x0099207F | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x0099209A | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x00992391 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x009923C5 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x00992402 | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x00992514 | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x00992664 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x0099269C | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x009926C2 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x00998542 | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x0099856D | `TVOutSlideshow_Screen_Playing` | Known | Screen layout |
| 0x0099858B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x009985C5 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x00998662 | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x009986CD | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x0099874D | `Extras_Screen_Debug` | Known | Screen layout |
| 0x00998857 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x00998877 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x00998DA5 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x00998E03 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x00998E1E | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00998E31 | `MainMenu_List_ScreenLock` | Known | Screen layout |
| 0x00998E4A | `ExtrasMenu_ScreenLock` | Known | Screen layout |
| 0x00998EBD | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x00998EDE | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x00998FB1 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x00998FD3 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x009990DA | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0099911A | `VoiceMemos_Screen_DeletAllAsk` | Known | Screen layout |
| 0x00999138 | `NowPlaying_Screen_TVOutAsk` | Known | Screen layout |
| 0x00999294 | `NowPlaying_Screen_Initial` | Known | Screen layout |
| 0x009992AE | `NowPlaying_Screen_Video_TVOut_Initial` | Known | Screen layout |
| 0x0099A016 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x0099A097 | `RemoteUI_Screen` | Known | Screen layout |
| 0x0099A0A7 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x0099A0BF | `Video_Settings_TV_Screen` | Known | Screen layout |
| 0x0099A0D8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0099A0EF | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x0099A113 | `AddressViewer_PartialLoad_Screen` | Known | Screen layout |
| 0x0099A134 | `Alarms_Alarm_Clock_Triggered_Screen` | Known | Screen layout |
| 0x0099A158 | `Alarms_Alarm_Triggered_Screen` | Known | Screen layout |
| 0x0099A176 | `Unsupported_Screen` | Known | Screen layout |
| 0x0099A189 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x0099A1A7 | `LockediPod_Screen` | Known | Screen layout |
| 0x0099A1B9 | `DiskMode_Screen` | Known | Screen layout |
| 0x0099A1C9 | `DemoMode_Screen` | Known | Screen layout |
| 0x0099A1D9 | `Notes_Image_Screen` | Known | Screen layout |
| 0x0099A1EC | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x0099A20A | `NowPlaying_Idle_Screen` | Known | Screen layout |
| 0x0099A221 | `Game_Screen` | Known | Screen layout |
| 0x0099A22D | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x0099A24A | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x0099A263 | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x0099A284 | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x0099A2A9 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x0099A2BC | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x0099A2D9 | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x0099A2FA | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x0099A31F | `Notes_Loading_Screen` | Known | Screen layout |
| 0x0099A334 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0099A34A | `SettingsMenus_AdjustScrolling_Screen` | Known | Screen layout |
| 0x0099A36F | `Game_Running_Screen` | Known | Screen layout |
| 0x0099A383 | `Stopwatch_Screen` | Known | Screen layout |
| 0x0099A394 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0099A3AB | `Clock_Screen` | Known | Screen layout |
| 0x0099A3B8 | `TVOutSlideshowAsk_Screen` | Known | Screen layout |
| 0x0099A3D1 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x0099A3E7 | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x0099A405 | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x0099A421 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x0099A432 | `DemoMode_Main_Screen` | Known | Screen layout |
| 0x0099A447 | `Search_Main_Screen` | Known | Screen layout |
| 0x0099A45A | `AddressViewer_Main_Screen` | Known | Screen layout |
| 0x0099A474 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x0099A489 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x0099A49F | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0099A4B9 | `Clock_Region_Screen` | Known | Screen layout |
| 0x0099A4CD | `RentalDeleted_Notification_Screen` | Known | Screen layout |
| 0x0099A4EF | `SingleRentalExpiring_Notification_Screen` | Known | Screen layout |
| 0x0099A518 | `MultipleRentalsExpiring_Notification_Screen` | Known | Screen layout |
| 0x0099A544 | `RentalError_Notification_Screen` | Known | Screen layout |
| 0x0099A564 | `DeleteRental_Confirmation_Screen` | Known | Screen layout |
| 0x0099A585 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0099A59D | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x0099A5BB | `Video_Settings_Fit_To_Screen` | Known | Screen layout |
| 0x0099A5D8 | `RentalInfo_Screen` | Known | Screen layout |
| 0x0099A5EA | `Radio_Screen` | Known | Screen layout |
| 0x0099A5F7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0099A60B | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x0099A625 | `ChargingMode_LowPower_Screen` | Known | Screen layout |
| 0x0099A642 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0099A65C | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0099A676 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0099A690 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0099A6A4 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0099A6BD | `Extras_Screen` | Known | Screen layout |
| 0x0099A6CB | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x0099A6E8 | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x0099A70A | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x0099A723 | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x0099A741 | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x0099A75A | `Video_Settings_Screen` | Known | Screen layout |
| 0x0099A770 | `SettingsMenus_AudiobookSettings_Screen` | Known | Screen layout |
| 0x0099A797 | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x0099A7BD | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0099A7D3 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0099A7EB | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x0099A80E | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x0099A82B | `MediaLists_Rentals_Screen` | Known | Screen layout |
| 0x0099A845 | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x0099A869 | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x0099A882 | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x0099A8A4 | `ToDo_Instructions_Screen` | Known | Screen layout |
| 0x0099A8BD | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x0099A8D9 | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x0099A8F3 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x0099A914 | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x0099A930 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0099A948 | `VoiceMemos_Screen` | Known | Screen layout |
| 0x0099A95A | `No_Photos_Screen` | Known | Screen layout |
| 0x0099A96B | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x0099A985 | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x0099A9A1 | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x0099A9C5 | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x0099A9E5 | `ChargingMode_Progress_Screen` | Known | Screen layout |
| 0x0099AA02 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x0099AA18 | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x0099AA33 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x0099AA4F | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x0099AA71 | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x0099AA92 | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x0099AAAC | `MediaLists_Genius_Screen` | Known | Screen layout |
| 0x0099AAC5 | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x0099AADF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0099AAFE | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x0099AB1F | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x0099AB37 | `NoContent_Screen` | Known | Screen layout |
| 0x0099AB48 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0099AB5E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0099AB6F | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x0099AB85 | `Notes_List_Screen` | Known | Screen layout |
| 0x0099AB97 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x0099ABAD | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x0099ABCE | `MediaLists_GeniusPlaylist_Screen` | Known | Screen layout |
| 0x0099ABEF | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x0099AC09 | `TODOS_Menu_Screen` | Known | Screen layout |
| 0x0099AC1B | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x0099AC31 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0099AC4D | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0099AC62 | `Games_Menu_Screen` | Known | Screen layout |
| 0x0099AC74 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0099AC87 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x0099ACA6 | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x0099ACC5 | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x0099ACE9 | `ContextualMenu_Screen` | Known | Screen layout |
| 0x0099ACFF | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x0099AD15 | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x0099AD33 | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x0099AD56 | `TVOutSlideshow_Screen` | Known | Screen layout |
| 0x0099AD6C | `CoverFlow_Screen` | Known | Screen layout |
| 0x0099AD7D | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0099AD91 | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x0099ADB3 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x0099ADCB | `Stopwatch_SessionSummary_Screen` | Known | Screen layout |
| 0x0099ADEB | `SettingsMenu_AboutWithAccessory_Screen` | Known | Screen layout |
| 0x0099AE12 | `Clock_NorthAmerica_City_Screen` | Known | Screen layout |
| 0x0099AE31 | `Clock_SouthAmerica_City_Screen` | Known | Screen layout |
| 0x0099AE50 | `Clock_Africa_City_Screen` | Known | Screen layout |
| 0x0099AE69 | `Clock_Australia_City_Screen` | Known | Screen layout |
| 0x0099AE85 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x0099AE9C | `Clock_Pacific_City_Screen` | Known | Screen layout |
| 0x0099AEB6 | `Clock_Atlantic_City_Screen` | Known | Screen layout |
| 0x0099AED1 | `Clock_Europe_City_Screen` | Known | Screen layout |
| 0x0099AFB1 | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x0099B002 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0099B025 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0099B04D | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x0099B3E8 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0099B4EB | `RentalInfo_Screen_ExpiringSoon` | Known | Screen layout |
| 0x0099B541 | `RentalInfo_Screen_NoAlbumArt_ExpiringSoon` | Known | Screen layout |
| 0x0099B910 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x0099B966 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x0099BAB7 | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x0099BAD4 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x0099BEA8 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x0099BFCA | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0099BFEC | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0099C059 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x0099C078 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x0099C6BA | `ChargingMode_ScreenLayout_LowPower` | Known | Screen layout |
| 0x0099D057 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x0099D070 | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x0099D1B8 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x0099D294 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0099D2B2 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0099D2D2 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0099D3DD | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x0099D3F9 | `Extras_Screen_Games` | Known | Screen layout |
| 0x0099D4FF | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0099D51E | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x0099D53A | `Extras_Screen_Notes` | Known | Screen layout |
| 0x0099D605 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x0099D6E0 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0099D8AE | `MainMenu_Music_Screen_NoAudiobooks` | Known | Screen layout |
| 0x0099D8D1 | `MainMenus_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x0099D8F4 | `NoContent_Screen_NoAudiobooks` | Known | Screen layout |
| 0x0099D92E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0099D94D | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0099D96E | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0099DA1D | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x0099DA3A | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x0099DAB9 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0099DB9D | `MainMenu_Music_Screen_NoCompilations` | Known | Screen layout |
| 0x0099DBC2 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0099DD49 | `MainMenu_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0099DD6C | `MainMenu_Videos_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0099DD91 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0099DDB0 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0099DDCF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0099DDF0 | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x0099DE2E | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x0099DE4F | `Extras_Screen_VoiceMemos` | Known | Screen layout |
| 0x0099DEBA | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0099DEEC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0099DF0B | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x0099DFB8 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x0099E024 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0099E11D | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x0099E139 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x0099E1BC | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x0099E1D7 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0099E1F8 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x0099E2A7 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0099E2DB | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x0099E2FC | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0099E39F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0099E3C0 | `MainMenu_Videos_Screen_NoPlaylists` | Known | Screen layout |
| 0x0099E3E3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0099E432 | `MainMenu_Music_Screen_NoArtists` | Known | Screen layout |
| 0x0099E4D9 | `NowPlaying_Screen_Genius` | Known | Screen layout |
| 0x0099E522 | `Genius_Error_Screen_NoGenius` | Known | Screen layout |
| 0x0099E53F | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x0099E55E | `Extras_Screen_NikePlus` | Known | Screen layout |
| 0x0099E6AE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0099E6CD | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0099E6EE | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x0099EB59 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x0099EC0C | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x0099EC86 | `Notes_Contents_Screen_Alt` | Known | Screen layout |
| 0x0099ECA0 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x0099ED4C | `VoiceMemos_Screen_InsufficientDiskSpace2_Default` | Known | Screen layout |
| 0x0099EDFE | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x0099EEA3 | `VoiceMemos_Screen_InsufficientDiskSpace_Default` | Known | Screen layout |
| 0x0099EED3 | `NowPlaying_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x0099EF00 | `Slideshow_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x0099FB91 | `Radio_Screen_Tuning_Default` | Known | Screen layout |
| 0x0099FBF2 | `VoiceMemos_Screen_DeletAllAsk_Default` | Known | Screen layout |
| 0x0099FC18 | `NowPlaying_Screen_TVOutAsk_Default` | Known | Screen layout |
| 0x0099FC3B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0099FC59 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x0099FC85 | `AddressViewer_PartialLoad_Screen_Default` | Known | Screen layout |
| 0x0099FCAE | `Alarms_Alarm_Clock_Triggered_Screen_Default` | Known | Screen layout |
| 0x0099FCDA | `Alarms_Alarm_Triggered_Screen_Default` | Known | Screen layout |
| 0x0099FD00 | `Unsupported_Screen_Default` | Known | Screen layout |
| 0x0099FD1B | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x0099FD41 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x0099FD59 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x0099FD74 | `Game_Screen_Default` | Known | Screen layout |
| 0x0099FD88 | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x0099FDAE | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0099FDCF | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x0099FDF8 | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x0099FE22 | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x0099FE4F | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x0099FE78 | `Notes_Loading_Screen_Default` | Known | Screen layout |
| 0x0099FE95 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0099FEB3 | `Clock_Screen_Default` | Known | Screen layout |
| 0x0099FEC8 | `TVOutSlideshowAsk_Screen_Default` | Known | Screen layout |
| 0x0099FEE9 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x0099FF07 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x0099FF2D | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x0099FF51 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x0099FF6A | `AddressViewer_Main_Screen_Default` | Known | Screen layout |
| 0x0099FF8C | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x0099FFA9 | `MainMenus_Main_Screen_Default` | Known | Screen layout |
| 0x0099FFC7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0099FFE4 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x009A0000 | `RentalDeleted_Notification_Screen_Default` | Known | Screen layout |
| 0x009A002A | `SingleRentalExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x009A005B | `MultipleRentalsExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x009A008F | `RentalError_Notification_Screen_Default` | Known | Screen layout |
| 0x009A00B7 | `DeleteRental_Confirmation_Screen_Default` | Known | Screen layout |
| 0x009A00E0 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x009A010C | `RentalInfo_Screen_Default` | Known | Screen layout |
| 0x009A0126 | `Radio_Screen_Default` | Known | Screen layout |
| 0x009A013B | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x009A0157 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x009A0179 | `Extras_Screen_Default` | Known | Screen layout |
| 0x009A018F | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x009A01B5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x009A01D6 | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x009A01F4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x009A0216 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x009A0242 | `ToDo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009A0263 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009A0287 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x009A02A9 | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x009A02CD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x009A02EC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x009A0305 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x009A0327 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x009A034B | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x009A0369 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x009A038D | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x009A03B7 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x009A03E0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x009A0402 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x009A0423 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x009A0443 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x009A0461 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x009A047A | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x009A0498 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x009A04B2 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x009A04D0 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x009A04F9 | `MediaLists_GeniusPlaylist_Screen_Default` | Known | Screen layout |
| 0x009A0522 | `TODOS_Menu_Screen_Default` | Known | Screen layout |
| 0x009A053C | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x009A055A | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x009A0577 | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x009A0591 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x009A05AC | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x009A05CB | `ContextualMenu_Screen_Default` | Known | Screen layout |
| 0x009A05E9 | `Debug_MainMenu_Screen_Default` | Known | Screen layout |
| 0x009A0607 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x009A0625 | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x009A063E | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x009A065A | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x009A0684 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x009A06A4 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x009A06CC | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009A06F3 | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009A071A | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x009A073B | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x009A075F | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x009A077E | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x009A07A0 | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x009A07C3 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x009A07E4 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x009A0872 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x009A08A2 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x009A08C4 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x009A0935 | `RentalInfo_Screen_NoAlbumArt_Default` | Known | Screen layout |
| 0x009A095A | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x009A0F35 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009A0F61 | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009A0FA6 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x009A0FCE | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x009A0FEF | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x009A1010 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x009A1036 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x009A1053 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x009A1075 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x009A1099 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x009A10BD | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x009A128D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x009A1368 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x009A13B9 | `NowPlaying_Screen_Video_TVOut` | Known | Screen layout |
| 0x009A152B | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x009A1552 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x009A1A8B | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x009A1C48 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x009A1E3A | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x009A2106 | `Lock_Screen_Default_Layout` | Known | Screen layout |
| 0x009A219C | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x009A21C3 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x009A23DF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x009A24B9 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x009A2520 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009A254A | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009A4E4E | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x009A4E9A | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x009A4F78 | `MainMenu_List_ScreenLock_x` | Known | Screen layout |
| 0x009A5246 | `RentalInfo_Screen_ExpiresToday` | Known | Screen layout |
| 0x009A529C | `RentalInfo_Screen_NoAlbumArt_ExpiresToday` | Known | Screen layout |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000908B | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | Known | RTOS |
| 0x002A26B8 | `  K - RTXC` | Known | RTOS |
| 0x002A36C0 | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | Known | RTOS |
| 0x0098A26C | `Returning from RTXCBug` | Known | RTOS |

---

## 9. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000D2278 | `HostOSTask` | Known | RTOS task thread |
| 0x0012CD4C | `MP3ExampleTask` | Known | RTOS task thread |
| 0x00132214 | `USBDeviceTask` | Known | RTOS task thread |
| 0x0013C354 | `DiskReaderTask` | Known | RTOS task thread |
| 0x0014C4C4 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0014C4D8 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0019FB98 | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x001DB17C | `MeCCAIOTask` | Known | RTOS task thread |
| 0x0020E3EC | `StreamCacheMassStorageManagerTimeOutTask` | Known | RTOS task thread |
| 0x0020E568 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x00290858 | `FirewireTask` | Known | RTOS task thread |
| 0x0029086C | `TouchwheelTask` | Known | RTOS task thread |
| 0x00290880 | `AudioOutStateTask` | Known | RTOS task thread |
| 0x002908AC | `DiskMgrTask` | Known | RTOS task thread |
| 0x002908BC | `HoldSwitchTask` | Known | RTOS task thread |
| 0x002908D0 | `MikeyTask` | Known | RTOS task thread |
| 0x002908E0 | `TopPlugTask` | Known | RTOS task thread |
| 0x002908F0 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x00290968 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x00290990 | `AlarmTask` | Known | RTOS task thread |
| 0x002909AF | `"USBAudioTask` | Known | RTOS task thread |
| 0x002A2D58 | `Undefined Task` | Known | RTOS task thread |
| 0x003E4E78 | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x003E8544 | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x003F0C50 | `MeCCARecordingTask` | Known | RTOS task thread |
| 0x008DC850 | `WaveFileDebugTask` | Known | RTOS task thread |

---

## 10. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00248D4C | `Channel Reserved` | Known | Logging channel |
| 0x00248D60 | `Channel AppBoot` | Known | Logging channel |
| 0x00248D70 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x00248D8C | `Channel PrefsWriting` | Known | Logging channel |
| 0x00248DA4 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x00248DC4 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x00248DDC | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x00248DF8 | `Channel TestLogging` | Known | Logging channel |
| 0x00248E0C | `Channel AppFileLoading` | Known | Logging channel |
| 0x00248E24 | `Channel VCardReading` | Known | Logging channel |
| 0x00248E3C | `Channel LongSongScanning` | Known | Logging channel |
| 0x00248EB0 | `Channel VoiceRecording` | Known | Logging channel |
| 0x00248EC8 | `Channel PhotoImporting` | Known | Logging channel |
| 0x00248EE0 | `Channel Notes` | Known | Logging channel |
| 0x00248EF0 | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x00248F0C | `Channel DiskMode` | Known | Logging channel |
| 0x00248F20 | `Channel Firewire` | Known | Logging channel |
| 0x00248F34 | `Channel USB` | Known | Logging channel |
| 0x00248F54 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x00248F6C | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00081D20 | `gamedata_RW` | Known | Game system |
| 0x00081D3C | `gamedata_ShareRW` | Known | Game system |
| 0x00081D50 | `games_RO` | Known | Game system |
| 0x0095F0DF | `11TCGamesMenu` | Known | Game system |
| 0x0095F1B3 | `12TCGameScreen` | Known | Game system |
| 0x0095FF7F | `27TSilverCntlrTransitionAddonI11TCGamesMenuE` | Known | Game system |
| 0x00960034 | `27TSilverCntlrTransitionAddonI12TCGameScreenE` | Known | Game system |
| 0x0098A2C6 | `iPod_Control/games_RO/` | Known | Game system |
| 0x0098A2DD | `Resources/Games/games_RO/` | Known | Game system |
| 0x00995D54 | `GamesMenu_StatusBar_String` | Known | Game system |
| 0x009964DC | `AboutScreen_Games_String` | Known | Game system |
| 0x0099D40D | `MainMenu_List_Games` | Known | Game system |
| 0x0099D421 | `ExtrasMenu_Games` | Known | Game system |
| 0x009A4FE7 | `MainMenu_List_Games_x` | Known | Game system |

---

## 12. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00092210 | `adrmmp4a` | Known | DRM system |
| 0x00139824 | `AppleDRMVersion` | Known | DRM system |
| 0x001398C4 | `AppleDRM` | Known | DRM system |
| 0x0013AB34 | `AppleVideoDRM` | Known | DRM system |
| 0x0013E0C4 | `tx3gdrmsp608aavdmp4aesds@` | Known | DRM system |
| 0x001E8B24 | `drmttx3g` | Known | DRM system |
| 0x0098A74F | `DRMLevel` | Known | DRM system |

---

## 13. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00031350 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00031368 | `iPod_Control\iTunes\firsttime` | Known | iTunes database |
| 0x0005273C | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x00052764 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00058A64 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0007DD38 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x00081CB0 | `iPod_Control/iTunes/iTunesDB.p7b` | Known | iTunes database |
| 0x00095144 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0009E774 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x0009E95C | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x0009F288 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000A7430 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000A8928 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000A8A28 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0012540C | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x0021B420 | `%s/sqlite_` | Known | SQLite database |
| 0x00281970 | `iPod_Control/iTunes/primary.db` | Known | iTunes database |
| 0x00282538 | `iPod_Control/iTunes/Extras.itdb` | Known | iTunes database |
| 0x002A61AC | `sqlite3BtreeInitPage() returns error code %d` | Known | SQLite database |
| 0x002A9488 | `sqlite_master` | Known | SQLite database |
| 0x002A9498 | `sqlite_temp_master` | Known | SQLite database |
| 0x002C03D0 | `sqlite_stat1` | Known | SQLite database |
| 0x002C03E0 | `CREATE TABLE %Q.sqlite_stat1(tbl,idx,stat)` | Known | SQLite database |
| 0x002C040C | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x002CADB4 | `sqlite_subquery_%p_` | Known | SQLite database |
| 0x00360034 | `sqlite_master` | Known | SQLite database |
| 0x00360044 | `sqlite_temp_master` | Known | SQLite database |
| 0x00360368 | `sqlite_` | Known | SQLite database |
| 0x003603A8 | `sqlite_master` | Known | SQLite database |
| 0x003603B8 | `sqlite_temp_master` | Known | SQLite database |
| 0x003603D0 | `sqlite_sequence` | Known | SQLite database |
| 0x003603E0 | `UPDATE "%w".sqlite_sequence set name = %Q WHERE name = %Q` | Known | SQLite database |
| 0x003604C4 | `sqlite_stat1` | Known | SQLite database |
| 0x003604D4 | `SELECT idx, stat FROM %Q.sqlite_stat1` | Known | SQLite database |
| 0x003611B0 | `sqlite_` | Known | SQLite database |
| 0x003613AC | `sqlite_master` | Known | SQLite database |
| 0x003613BC | `sqlite_temp_master` | Known | SQLite database |
| 0x003640D8 | `sqlite_` | Known | SQLite database |
| 0x003653C4 | `sqlite_autoindex_` | Known | SQLite database |
| 0x003653D8 | `sqlite_master` | Known | SQLite database |
| 0x003653E8 | `sqlite_temp_master` | Known | SQLite database |
| 0x00366840 | `sqlite_master` | Known | SQLite database |
| 0x00366850 | `sqlite_temp_master` | Known | SQLite database |
| 0x00366884 | `sqlite_stat1` | Known | SQLite database |
| 0x00366894 | `DELETE FROM %Q.sqlite_stat1 WHERE idx=%Q` | Known | SQLite database |
| 0x00366B7C | `sqlite_master` | Known | SQLite database |
| 0x00366B8C | `sqlite_temp_master` | Known | SQLite database |
| 0x00366C00 | `DELETE FROM %s.sqlite_sequence WHERE name=%Q` | Known | SQLite database |
| 0x00366C68 | `sqlite_stat1` | Known | SQLite database |
| 0x00366C78 | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x00366FF0 | `sqlite_master` | Known | SQLite database |
| 0x00367000 | `sqlite_temp_master` | Known | SQLite database |
| 0x00367418 | `sqlite_master` | Known | SQLite database |
| 0x00367428 | `sqlite_temp_master` | Known | SQLite database |
| 0x00367440 | `CREATE TABLE %Q.sqlite_sequence(name,seq)` | Known | SQLite database |
| 0x0036A6C8 | `sqlite_master` | Known | SQLite database |
| 0x0036A6D8 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036CAC0 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036CAD8 | `sqlite_master` | Known | SQLite database |
| 0x0036E2B4 | `sqlite3_extension_init` | Known | SQLite database |
| 0x0036EAA8 | `sqlite_master` | Known | SQLite database |
| 0x0036EAB8 | `sqlite_temp_master` | Known | SQLite database |
| 0x00372E98 | `sqlite_attach` | Known | SQLite database |
| 0x00372EAC | `sqlite_detach` | Known | SQLite database |
| 0x00375BE0 | `sqlite_master` | Known | SQLite database |
| 0x00375BF0 | `sqlite_temp_master` | Known | SQLite database |
| 0x00375C40 | `sqlite_sequence` | Known | SQLite database |
| 0x0037B4CC | `sqlite_master` | Known | SQLite database |
| 0x0037B4DC | `sqlite_temp_master` | Known | SQLite database |
| 0x0037E870 | `sqlite_master` | Known | SQLite database |
| 0x0037E880 | `sqlite_temp_master` | Known | SQLite database |
| 0x0038CA1C | `sqlite_attach` | Known | SQLite database |
| 0x0038CA2C | `sqlite_detach` | Known | SQLite database |
| 0x003DE134 | `iTunesDB` | Known | iTunes database |
| 0x003DE140 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x008D889F | `SQLite format 3` | Known | SQLite database |
| 0x008DAF4C | `CREATE TABLE sqlite_master(` | Known | SQLite database |
| 0x008DAFB4 | `CREATE TEMP TABLE sqlite_temp_master(` | Known | SQLite database |
| 0x008DB67C | `illegal return value (%d) from the authorization function - should be SQLITE_OK,` | Known | SQLite database |
| 0x008DB734 | `SELECT 'CREATE TABLE vacuum_db.' || substr(sql,14)   FROM sqlite_master WHERE ty` | Known | SQLite database |
| 0x008DB7BC | `SELECT 'CREATE INDEX vacuum_db.' || substr(sql,14)  FROM sqlite_master WHERE sql` | Known | SQLite database |
| 0x008DB824 | `SELECT 'CREATE UNIQUE INDEX vacuum_db.' || substr(sql,21)   FROM sqlite_master W` | Known | SQLite database |
| 0x008DB89C | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x008DB94C | `SELECT 'DELETE FROM vacuum_db.' || quote(name) || ';' FROM vacuum_db.sqlite_mast` | Known | SQLite database |
| 0x008DB9C0 | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x008DBA58 | `INSERT INTO vacuum_db.sqlite_master   SELECT type, name, tbl_name, rootpage, sql` | Known | SQLite database |
| 0x008DBC18 | `UPDATE %Q.%s SET sql = CASE WHEN type = 'trigger' THEN sqlite_rename_trigger(sql` | Known | SQLite database |
| 0x008DBD8C | `UPDATE sqlite_temp_master SET sql = sqlite_rename_trigger(sql, %Q), tbl_name = %` | Known | SQLite database |
| 0x008DBFC8 | `sqlite3_get_table() called with two or more incompatible queries` | Known | SQLite database |
| 0x009A5A76 | `sqlite_rename_table` | Known | SQLite database |
| 0x009A5BF9 | `sqlite_version` | Known | SQLite database |
| 0x009A5C93 | `sqlite_rename_trigger` | Known | SQLite database |
| 0x009A5FB7 | `SQLite_iPod_VFS` | Known | SQLite database |

---

## 14. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0005EE5C | `cI: could not read CE-ATA task file` | Known | Hardware |
| 0x0005EE84 | `cI: CE-ATA signature missing (%x,%x)` | Known | Hardware |
| 0x0005EEDC | `cI: CE-ATA interrupt enable failed` | Known | Hardware |
| 0x00124C70 | `NAND FLASH DRIVE` | Known | Hardware |
| 0x00139D6C | `FireWireGUID` | Known | FireWire |
| 0x00139D7C | `FireWireVersion` | Known | FireWire |
| 0x0013A458 | `FireWire` | Known | FireWire |
| 0x0035B348 | `CE-ATA init failed` | Known | Hardware |
| 0x0035B808 | `ISDIE: CE-ATA interrupt enable failed` | Known | Hardware |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00726252 | `Settings_RadioRegionsMenu_Layout"` | Known | FM Radio |
| 0x007262DB | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x007D70D8 | `Radio Regions` | Known | FM Radio |
| 0x00827064 | `Radio-Regionen` | Known | FM Radio |
| 0x0095FB7C | `23TCSettings_RadioRegions` | Known | FM Radio |
| 0x00960A8F | `27TSilverCntlrTransitionAddonI23TCSettings_RadioRegionsE` | Known | FM Radio |
| 0x00992DB5 | `Settings_RadioRegions_Australia_String` | Known | FM Radio |
| 0x00992DDC | `Settings_RadioRegions_Asia_String` | Known | FM Radio |
| 0x00994041 | `Settings_RadioRegions_Europe_String` | Known | FM Radio |
| 0x00995650 | `Settings_RadioRegions_Japan_String` | Known | FM Radio |
| 0x009962F9 | `Settings_RadioRegions_Americas_String` | Known | FM Radio |
| 0x009969DB | `SettingsMenu_RadioRegions_String` | Known | FM Radio |
| 0x00999ED5 | `RadioRegionsMenuItem` | Known | FM Radio |
| 0x0099DB26 | `SettingsMenu_ListItem_RadioRegions` | Known | FM Radio |
| 0x009A1D14 | `SettingsInfo_Template_RadioRegions_Layout` | Known | FM Radio |
| 0x009A1D3E | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x009A23A0 | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 16. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00865E4C | `Fotocamera` | Known | Camera |
| 0x008663B0 | `Collega fotocamera o inserisci scheda` | Known | Camera |
| 0x00866428 | `Fotocamera non supportata` | Known | Camera |
| 0x008854F4 | `Camera` | Known | Camera |
| 0x00885A74 | `Sluit camera of kaart aan` | Known | Camera |
| 0x00885AE0 | `Camera niet ondersteund` | Known | Camera |
| 0x00992DFE | `Photos_DevName_Camera_String` | Known | Camera |

---

## 17. Pedometer

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x009A534F | `NikePlus_Remote_Step_Away` | Known | Pedometer |
| 0x009A5369 | `NikePlus_Step_Away` | Known | Pedometer |
| 0x009A5C34 | `AggStep` | Known | Pedometer |

---

## 18. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0003133C | `iPod_Control` | Filesystem Path |  |
| 0x000313A8 | `iPod_Control\Device` | Filesystem Path |  |
| 0x0003FC04 | `iPod_Control\Device` | Filesystem Path |  |
| 0x00041C90 | `iPod_Control` | Filesystem Path |  |
| 0x000422FC | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x0005271C | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path |  |
| 0x000552C4 | `iPod_Control\Music\` | Filesystem Path |  |
| 0x000588E4 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x0008BBF8 | `iPod_Control` | Filesystem Path |  |
| 0x0008BC08 | `Resources/Games` | Filesystem Path |  |
| 0x0008BC18 | `iPod_Control/%s%s%s` | Filesystem Path |  |
| 0x000E9058 | `iPod_Control\Device\dst` | Filesystem Path |  |
| 0x000F4158 | `iPod_Control/Device/alarms` | Filesystem Path |  |
| 0x00104674 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00105BB4 | `iPod_Control/Device` | Filesystem Path |  |
| 0x00105BC8 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x0011FE10 | `iPod_Control/Device/Users` | Filesystem Path |  |
| 0x0014D968 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x0014DBC4 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x0015A548 | `Resources/UI/active.bin` | Filesystem Path |  |
| 0x0015A560 | `Resources/UI/` | Filesystem Path |  |
| 0x0017DF94 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x0017EEC0 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path |  |
| 0x0017EEE8 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001A31E0 | `iPod_Control/Device/PlayCounts` | Filesystem Path |  |
| 0x001B91B0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9260 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B93DC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9574 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B961C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B97CC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9870 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9914 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B99B8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9A5C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9B0C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9BB0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9C54 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9D04 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9DB4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9E64 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9FD0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA080 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA130 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA1D4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA284 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA378 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA41C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA4D0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA58C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA63C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA760 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA81C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA8CC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAA88 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAB4C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BABFC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BACB8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BADF4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAEC0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAF7C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB020 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB0C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB180 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB23C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB304 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB3A8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB470 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB538 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB5E8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB6B0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB778 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB828 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB8D8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB99C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBA4C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBAFC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBBAC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBC80 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBD54 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBE54 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBF34 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC03C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC128 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003DE1B2 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003E4718 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x003E690C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003E6D5E | `iPod_Control/Device` | Filesystem Path |  |
| 0x003E86B0 | `Resources/Fonts` | Filesystem Path |  |
| 0x003F0C1C | `iPod_Control:Device:voices:` | Filesystem Path |  |
| 0x0098A1A1 | `Resources/Games/` | Filesystem Path |  |
| 0x0098A5BF | `iPod_Control/Device` | Filesystem Path |  |
| 0x0098A5D3 | `iPod_Control/Device/accessories.reg` | Filesystem Path |  |
| 0x0098A6C6 | `iPod_Control/Tones` | Filesystem Path |  |

---

## 19. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x008DEF90 | `c:\bwa\N25BFirmwareWin-93\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftu` | Build Path |  |
| 0x008DEFE8 | `c:\bwa\N25BFirmwareWin-93\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fts` | Build Path |  |
| 0x008DF040 | `c:\bwa\N25BFirmwareWin-93\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fto` | Build Path |  |
| 0x008E9C90 | `c:\bwa\N25BFirmwareWin-93\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\` | Build Path |  |
| 0x008EA80C | `c:\bwa\N25BFirmwareWin-93\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfd` | Build Path |  |
| 0x008EBA08 | `c:\bwa\N25BFirmwareWin-93\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrg` | Build Path |  |
| 0x008EBA60 | `c:\bwa\N25BFirmwareWin-93\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrc` | Build Path |  |
| 0x008EBAB8 | `c:\bwa\N25BFirmwareWin-93\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfro` | Build Path |  |
| 0x008EBDFC | `c:\bwa\N25BFirmwareWin-93\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1` | Build Path |  |
| 0x008FB1A4 | `c:\bwa\N25BFirmwareWin-93\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttc` | Build Path |  |
| 0x008FB420 | `c:\bwa\N25BFirmwareWin-93\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype` | Build Path |  |
| 0x008FB98C | `c:\bwa\N25BFirmwareWin-93\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1` | Build Path |  |

---

## 20. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00089840 | `Acoustic` | EQ Preset |  |
| 0x0008984C | `Bass Booster` | EQ Preset |  |
| 0x0008986C | `Classical` | EQ Preset |  |
| 0x00089878 | `Dance` | EQ Preset |  |
| 0x00089888 | `Electronic` | EQ Preset |  |
| 0x0008989C | `Hip Hop` | EQ Preset |  |
| 0x000898A4 | `Jazz` | EQ Preset |  |
| 0x000898AC | `Latin` | EQ Preset |  |
| 0x000898B4 | `Loudness` | EQ Preset |  |
| 0x000898C0 | `Lounge` | EQ Preset |  |
| 0x000898C8 | `Piano` | EQ Preset |  |
| 0x000898DC | `Rock` | EQ Preset |  |
| 0x000898E4 | `Small Speakers` | EQ Preset |  |
| 0x000898F4 | `Spoken Word` | EQ Preset |  |
| 0x00089900 | `Treble Booster` | EQ Preset |  |
| 0x0008994C | `Vocal Booster` | EQ Preset |  |
| 0x007D73C8 | `Acoustic` | EQ Preset |  |
| 0x007D73D4 | `Bass Booster` | EQ Preset |  |
| 0x007D73F4 | `Classical` | EQ Preset |  |
| 0x007D7400 | `Dance` | EQ Preset |  |
| 0x007D7410 | `Electronic` | EQ Preset |  |
| 0x007D7424 | `Hip Hop` | EQ Preset |  |
| 0x007D742C | `Jazz` | EQ Preset |  |
| 0x007D7434 | `Latin` | EQ Preset |  |
| 0x007D743C | `Loudness` | EQ Preset |  |
| 0x007D7448 | `Lounge` | EQ Preset |  |
| 0x007D7450 | `Piano` | EQ Preset |  |
| 0x007D7460 | `Rock` | EQ Preset |  |
| 0x007D7468 | `Small Speakers` | EQ Preset |  |
| 0x007D7478 | `Spoken Word` | EQ Preset |  |
| 0x007D7484 | `Treble Booster` | EQ Preset |  |
| 0x007D74A4 | `Vocal Booster` | EQ Preset |  |
| 0x008145CC | `Acoustic` | EQ Preset |  |
| 0x008145D8 | `Bass Booster` | EQ Preset |  |
| 0x008145F8 | `Classical` | EQ Preset |  |
| 0x00814604 | `Dance` | EQ Preset |  |
| 0x00814614 | `Electronic` | EQ Preset |  |
| 0x00814628 | `Hip Hop` | EQ Preset |  |
| 0x00814630 | `Jazz` | EQ Preset |  |
| 0x00814638 | `Latin` | EQ Preset |  |
| 0x00814640 | `Loudness` | EQ Preset |  |
| 0x0081464C | `Lounge` | EQ Preset |  |
| 0x00814654 | `Piano` | EQ Preset |  |
| 0x00814664 | `Rock` | EQ Preset |  |
| 0x0081466C | `Small Speakers` | EQ Preset |  |
| 0x0081467C | `Spoken Word` | EQ Preset |  |
| 0x00814688 | `Treble Booster` | EQ Preset |  |
| 0x008146A8 | `Vocal Booster` | EQ Preset |  |
| 0x0081DA40 | `Acoustic` | EQ Preset |  |
| 0x0081DA4C | `Bass Booster` | EQ Preset |  |
| 0x0081DA6C | `Classical` | EQ Preset |  |
| 0x0081DA78 | `Dance` | EQ Preset |  |
| 0x0081DA88 | `Electronic` | EQ Preset |  |
| 0x0081DA9C | `Hip Hop` | EQ Preset |  |
| 0x0081DAA4 | `Jazz` | EQ Preset |  |
| 0x0081DAAC | `Latin` | EQ Preset |  |
| 0x0081DAB4 | `Loudness` | EQ Preset |  |
| 0x0081DAC0 | `Lounge` | EQ Preset |  |
| 0x0081DAC8 | `Piano` | EQ Preset |  |
| 0x0081DAD8 | `Rock` | EQ Preset |  |
| 0x0081DAE0 | `Small Speakers` | EQ Preset |  |
| 0x0081DAF0 | `Spoken Word` | EQ Preset |  |
| 0x0081DAFC | `Treble Booster` | EQ Preset |  |
| 0x0081DB1C | `Vocal Booster` | EQ Preset |  |
| 0x0082740C | `Acoustic` | EQ Preset |  |
| 0x0082743C | `Dance` | EQ Preset |  |
| 0x0082744C | `Electronic` | EQ Preset |  |
| 0x00827468 | `Jazz` | EQ Preset |  |
| 0x00827470 | `Latin` | EQ Preset |  |
| 0x00827478 | `Loudness` | EQ Preset |  |
| 0x0082748C | `Piano` | EQ Preset |  |
| 0x0082749C | `Rock` | EQ Preset |  |
| 0x0083F348 | `Dance` | EQ Preset |  |
| 0x0083F370 | `Hip Hop` | EQ Preset |  |
| 0x0083F378 | `Jazz` | EQ Preset |  |
| 0x0083F388 | `Loudness` | EQ Preset |  |
| 0x0083F394 | `Lounge` | EQ Preset |  |
| 0x0083F39C | `Piano` | EQ Preset |  |
| 0x0083F3AC | `Rock` | EQ Preset |  |
| 0x0084886C | `Jazz` | EQ Preset |  |
| 0x00848874 | `Latin` | EQ Preset |  |
| 0x00848888 | `Lounge` | EQ Preset |  |
| 0x00848890 | `Piano` | EQ Preset |  |
| 0x008488A0 | `Rock` | EQ Preset |  |
| 0x00851CC0 | `Hip Hop` | EQ Preset |  |
| 0x00851CC8 | `Jazz` | EQ Preset |  |
| 0x00851CE4 | `Lounge` | EQ Preset |  |
| 0x00851CEC | `Piano` | EQ Preset |  |
| 0x00851D04 | `Rock` | EQ Preset |  |
| 0x0085BD78 | `Latin` | EQ Preset |  |
| 0x0085BDA4 | `Rock` | EQ Preset |  |
| 0x00865738 | `Dance` | EQ Preset |  |
| 0x0086575C | `Hip Hop` | EQ Preset |  |
| 0x00865764 | `Jazz` | EQ Preset |  |
| 0x00865774 | `Loudness` | EQ Preset |  |
| 0x00865780 | `Lounge` | EQ Preset |  |
| 0x00865788 | `Piano` | EQ Preset |  |
| 0x00865798 | `Rock` | EQ Preset |  |
| 0x0087056C | `Acoustic` | EQ Preset |  |
| 0x00870578 | `Bass Booster` | EQ Preset |  |
| 0x00870598 | `Classical` | EQ Preset |  |
| 0x008705A4 | `Dance` | EQ Preset |  |
| 0x008705B4 | `Electronic` | EQ Preset |  |
| 0x008705C8 | `Hip Hop` | EQ Preset |  |
| 0x008705D0 | `Jazz` | EQ Preset |  |
| 0x008705D8 | `Latin` | EQ Preset |  |
| 0x008705E0 | `Loudness` | EQ Preset |  |
| 0x008705EC | `Lounge` | EQ Preset |  |
| 0x008705F4 | `Piano` | EQ Preset |  |
| 0x00870604 | `Rock` | EQ Preset |  |
| 0x0087060C | `Small Speakers` | EQ Preset |  |
| 0x0087061C | `Spoken Word` | EQ Preset |  |
| 0x00870628 | `Treble Booster` | EQ Preset |  |
| 0x00870648 | `Vocal Booster` | EQ Preset |  |
| 0x0087B1D8 | `Acoustic` | EQ Preset |  |
| 0x0087B1E4 | `Bass Booster` | EQ Preset |  |
| 0x0087B204 | `Classical` | EQ Preset |  |
| 0x0087B210 | `Dance` | EQ Preset |  |
| 0x0087B220 | `Electronic` | EQ Preset |  |
| 0x0087B234 | `Hip Hop` | EQ Preset |  |
| 0x0087B23C | `Jazz` | EQ Preset |  |
| 0x0087B244 | `Latin` | EQ Preset |  |
| 0x0087B24C | `Loudness` | EQ Preset |  |
| 0x0087B258 | `Lounge` | EQ Preset |  |
| 0x0087B260 | `Piano` | EQ Preset |  |
| 0x0087B270 | `Rock` | EQ Preset |  |
| 0x0087B278 | `Small Speakers` | EQ Preset |  |
| 0x0087B288 | `Spoken Word` | EQ Preset |  |
| 0x0087B294 | `Treble Booster` | EQ Preset |  |
| 0x0087B2B4 | `Vocal Booster` | EQ Preset |  |
| 0x00884DD8 | `Dance` | EQ Preset |  |
| 0x00884E0C | `Jazz` | EQ Preset |  |
| 0x00884E14 | `Latin` | EQ Preset |  |
| 0x00884E1C | `Loudness` | EQ Preset |  |
| 0x00884E28 | `Lounge` | EQ Preset |  |
| 0x00884E30 | `Piano` | EQ Preset |  |
| 0x00884E40 | `Rock` | EQ Preset |  |
| 0x0088E1F4 | `Dance` | EQ Preset |  |
| 0x0088E220 | `Jazz` | EQ Preset |  |
| 0x0088E230 | `Loudness` | EQ Preset |  |
| 0x0088E23C | `Lounge` | EQ Preset |  |
| 0x0088E244 | `Piano` | EQ Preset |  |
| 0x0088E254 | `Rock` | EQ Preset |  |
| 0x008978BC | `Hip Hop` | EQ Preset |  |
| 0x008978C4 | `Jazz` | EQ Preset |  |
| 0x008978E8 | `Lounge` | EQ Preset |  |
| 0x00897900 | `Rock` | EQ Preset |  |
| 0x008A138C | `Hip Hop` | EQ Preset |  |
| 0x008A1394 | `Jazz` | EQ Preset |  |
| 0x008A13B0 | `Lounge` | EQ Preset |  |
| 0x008A13B8 | `Piano` | EQ Preset |  |
| 0x008A13C8 | `Rock` | EQ Preset |  |
| 0x008B7E14 | `Acoustic` | EQ Preset |  |
| 0x008B7E20 | `Bass Booster` | EQ Preset |  |
| 0x008B7E40 | `Classical` | EQ Preset |  |
| 0x008B7E4C | `Dance` | EQ Preset |  |
| 0x008B7E5C | `Electronic` | EQ Preset |  |
| 0x008B7E70 | `Hip Hop` | EQ Preset |  |
| 0x008B7E78 | `Jazz` | EQ Preset |  |
| 0x008B7E80 | `Latin` | EQ Preset |  |
| 0x008B7E88 | `Loudness` | EQ Preset |  |
| 0x008B7E94 | `Lounge` | EQ Preset |  |
| 0x008B7E9C | `Piano` | EQ Preset |  |
| 0x008B7EAC | `Rock` | EQ Preset |  |
| 0x008B7EB4 | `Small Speakers` | EQ Preset |  |
| 0x008B7EC4 | `Spoken Word` | EQ Preset |  |
| 0x008B7ED0 | `Treble Booster` | EQ Preset |  |
| 0x008B7EF0 | `Vocal Booster` | EQ Preset |  |
| 0x008C1494 | `Hip Hop` | EQ Preset |  |
| 0x008C14A0 | `Latin` | EQ Preset |  |
| 0x008C14D8 | `Rock` | EQ Preset |  |
| 0x008CAC54 | `Acoustic` | EQ Preset |  |
| 0x008CAC60 | `Bass Booster` | EQ Preset |  |
| 0x008CAC80 | `Classical` | EQ Preset |  |
| 0x008CAC8C | `Dance` | EQ Preset |  |
| 0x008CAC9C | `Electronic` | EQ Preset |  |
| 0x008CACB0 | `Hip Hop` | EQ Preset |  |
| 0x008CACB8 | `Jazz` | EQ Preset |  |
| 0x008CACC0 | `Latin` | EQ Preset |  |
| 0x008CACC8 | `Loudness` | EQ Preset |  |
| 0x008CACD4 | `Lounge` | EQ Preset |  |
| 0x008CACDC | `Piano` | EQ Preset |  |
| 0x008CACEC | `Rock` | EQ Preset |  |
| 0x008CACF4 | `Small Speakers` | EQ Preset |  |
| 0x008CAD04 | `Spoken Word` | EQ Preset |  |
| 0x008CAD10 | `Treble Booster` | EQ Preset |  |
| 0x008CAD30 | `Vocal Booster` | EQ Preset |  |
| 0x008D4344 | `Acoustic` | EQ Preset |  |
| 0x008D4350 | `Bass Booster` | EQ Preset |  |
| 0x008D4370 | `Classical` | EQ Preset |  |
| 0x008D437C | `Dance` | EQ Preset |  |
| 0x008D438C | `Electronic` | EQ Preset |  |
| 0x008D43A0 | `Hip Hop` | EQ Preset |  |
| 0x008D43A8 | `Jazz` | EQ Preset |  |
| 0x008D43B0 | `Latin` | EQ Preset |  |
| 0x008D43B8 | `Loudness` | EQ Preset |  |
| 0x008D43C4 | `Lounge` | EQ Preset |  |
| 0x008D43CC | `Piano` | EQ Preset |  |
| 0x008D43DC | `Rock` | EQ Preset |  |
| 0x008D43E4 | `Small Speakers` | EQ Preset |  |
| 0x008D43F4 | `Spoken Word` | EQ Preset |  |
| 0x008D4400 | `Treble Booster` | EQ Preset |  |
| 0x008D4420 | `Vocal Booster` | EQ Preset |  |

---
