# iPod Classic 7G Rev B - RetailOS 2.0.2 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 2.0.2 |
| **IPSW** | iPod_35.2.0.2.ipsw |
| **Device** | iPod Classic 7G Rev B (2009, 160GB, Click Wheel, Cover Flow, Genius, CE-ATA HDD) |
| **UpdaterFamilyID** | 35 |
| **Binary Size** | 10,514,672 bytes (10.03 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,512,624 bytes |
| **Total Strings (>=4)** | 71,747 |
| **Function Prologues** | 22,792 (ARM: 17,413, Thumb: 5,379) |
| **DRAM References** | 106,510 |
| **Peripheral Refs** | 7,200 |
| **Build** | N25CFirmwareWin-10 |
| **SoC** | S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Codename** | N25C |
| **DFU PID** | 0x1223 |
| **SHA-256** | `59ec5fdee0afb620844ad6d9e832d39c7f76d46b56a60c3001c019e4551731a5` |

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
| 0x001479C0 | `TCSportTimer` | Known | Controller |
| 0x001479D8 | `TCSportTimerMenu` | Known | Controller |
| 0x001479F4 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x00147A18 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x00148DC8 | `TCVoiceMemos` | Known | Controller |
| 0x00148DE0 | `TCVoiceMemosMenu` | Known | Controller |
| 0x00148DFC | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x00148E1C | `TCVoiceMemosPlayback` | Known | Controller |
| 0x00148E3C | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x00148E5C | `TCVoiceMemosAlert` | Known | Controller |
| 0x0015AAD4 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x0015AAFC | `TCSettings_MainMenu` | Known | Controller |
| 0x0015AB18 | `TCSettings_MusicMenu` | Known | Controller |
| 0x0015AB38 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x0015AB58 | `TCSettings_Brightness` | Known | Controller |
| 0x0015AB78 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0015AB9C | `TCSettings_EQ` | Known | Controller |
| 0x0015ABB4 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x0015ABDC | `TCSettings_RadioRegions` | Known | Controller |
| 0x0015ABFC | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0015AC20 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x0015AC44 | `TCDateTimeScreen` | Known | Controller |
| 0x0015AC60 | `TCTimeZoneScreen` | Known | Controller |
| 0x0015AC7C | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0015ACA4 | `TCFirstBoot` | Known | Controller |
| 0x001711BC | `TCDemoMode` | Known | Controller |
| 0x00199A3C | `TCAddressViewerMainMenu` | Known | Controller |
| 0x00199A5C | `TCAddressViewerDetails` | Known | Controller |
| 0x00199A7C | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x00199AA0 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x001C658C | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x001C65B0 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x001CDDD8 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x00263390 | `TC_LockDialog` | Known | Controller |
| 0x002633A8 | `TC_LockScreen` | Known | Controller |
| 0x002633C0 | `TC_LockediPod` | Known | Controller |
| 0x002633D8 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x002633FC | `TCLockChosenDispatcher` | Known | Controller |
| 0x00268F80 | `TCClock` | Known | Controller |
| 0x00268F90 | `TCClockCityMenu` | Known | Controller |
| 0x00268FA8 | `TCClockRegionMenu` | Known | Controller |
| 0x00268FC4 | `TCAlarmMenu` | Known | Controller |
| 0x00268FD8 | `TCSleepTimerMenu` | Known | Controller |
| 0x00268FF4 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x00269014 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x0026903C | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x00269060 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00269084 | `TCAlarmDatePicker` | Known | Controller |
| 0x002690A0 | `TCAlarmTriggered` | Known | Controller |
| 0x0026FFC4 | `TCNotesDispatcher` | Known | Controller |
| 0x0026FFE0 | `TCNotesLoading` | Known | Controller |
| 0x0026FFF8 | `TCNotesList` | Known | Controller |
| 0x0027000C | `TCNotesContents` | Known | Controller |
| 0x003DE478 | `TCAlarmTriggered` | Known | Controller |
| 0x003DE48C | `TSilverCntlr` | Known | Controller |
| 0x003DE4AC | `TCClock` | Known | Controller |
| 0x003DE4B4 | `TCClockRegionMenu` | Known | Controller |
| 0x003DE4C8 | `TCClockCityMenu` | Known | Controller |
| 0x003DE4D8 | `TCAlarmMenu` | Known | Controller |
| 0x003DE4E4 | `TCSleepTimerMenu` | Known | Controller |
| 0x003DE4F8 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003DE510 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003DE530 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003DE54C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003DE568 | `TCAlarmDatePicker` | Known | Controller |
| 0x003DE5A0 | `TSilverCntlr` | Known | Controller |
| 0x003DE5C0 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003DE750 | `TSilverCntlr` | Known | Controller |
| 0x003DE770 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x003DE790 | `TCSettings_Brightness` | Known | Controller |
| 0x003DE7A8 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x003DE7C4 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x003DE7E4 | `TCSettings_RadioRegions` | Known | Controller |
| 0x003DE7FC | `TCSettings_EQ` | Known | Controller |
| 0x003DE80C | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x003DE828 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x003DE848 | `TCFirstBoot` | Known | Controller |
| 0x003DE854 | `TCSettings_MainMenu` | Known | Controller |
| 0x003DE868 | `TCSettings_MusicMenu` | Known | Controller |
| 0x003DE880 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003DE898 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x003DE8B4 | `TCDateTimeScreen` | Known | Controller |
| 0x003DE8C8 | `TCTimeZoneScreen` | Known | Controller |
| 0x003E59BC | `TSilverCntlr` | Known | Controller |
| 0x003E59DC | `TCClock` | Known | Controller |
| 0x003E59E4 | `TCClockRegionMenu` | Known | Controller |
| 0x003E59F8 | `TCClockCityMenu` | Known | Controller |
| 0x003E5A08 | `TCAlarmMenu` | Known | Controller |
| 0x003E5A14 | `TCSleepTimerMenu` | Known | Controller |
| 0x003E5A28 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003E5AA0 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003E5AC0 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003E5ADC | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003E5B10 | `TCAlarmDatePicker` | Known | Controller |
| 0x003E5B24 | `TCAlarmTriggered` | Known | Controller |
| 0x003E6C08 | `TSilverCntlr` | Known | Controller |
| 0x003E6C28 | `TC_LockDialog` | Known | Controller |
| 0x003E6C38 | `TC_LockScreen` | Known | Controller |
| 0x003E6C48 | `TC_LockediPod` | Known | Controller |
| 0x003E6C58 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003E6C74 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003E6C8C | `TSilverCntlr` | Known | Controller |
| 0x003E6E98 | `TSilverCntlr` | Known | Controller |
| 0x003E6EB4 | `TSilverCntlr` | Known | Controller |
| 0x003E6F18 | `TSilverCntlr` | Known | Controller |
| 0x003E6F38 | `TCNotesDispatcher` | Known | Controller |
| 0x003E6F4C | `TCNotesLoading` | Known | Controller |
| 0x003E6F5C | `TCNotesBase` | Known | Controller |
| 0x003E6F68 | `TCNotesList` | Known | Controller |
| 0x003E6F74 | `TCNotesContents` | Known | Controller |
| 0x003E6F84 | `TSilverCntlr` | Known | Controller |
| 0x003E6FA4 | `TCRemoteUI` | Known | Controller |
| 0x003E6FB0 | `TCUnsupported` | Known | Controller |
| 0x003E6FC0 | `TSilverCntlr` | Known | Controller |
| 0x003E7024 | `TSilverCntlr` | Known | Controller |
| 0x003E7044 | `TCSportTimer` | Known | Controller |
| 0x003E7054 | `TCSportTimerMenu` | Known | Controller |
| 0x003E7068 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x003E7084 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x003E70B4 | `TSilverCntlr` | Known | Controller |
| 0x003E71DC | `TSilverCntlr` | Known | Controller |
| 0x003E71FC | `TCDemoMode` | Known | Controller |
| 0x003E7208 | `TCClock` | Known | Controller |
| 0x003E7210 | `TCClockRegionMenu` | Known | Controller |
| 0x003E7224 | `TCClockCityMenu` | Known | Controller |
| 0x003E7234 | `TCAlarmMenu` | Known | Controller |
| 0x003E7240 | `TCSleepTimerMenu` | Known | Controller |
| 0x003E7254 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003E726C | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003E728C | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003E72A8 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003E72C4 | `TCAlarmDatePicker` | Known | Controller |
| 0x003E72D8 | `TCAlarmTriggered` | Known | Controller |
| 0x003E72F8 | `TSilverCntlr` | Known | Controller |
| 0x003E7314 | `TSilverCntlr` | Known | Controller |
| 0x003E7324 | `TSilverCntlr` | Known | Controller |
| 0x003E7344 | `TCVoiceMemos` | Known | Controller |
| 0x003E7354 | `TCVoiceMemosMenu` | Known | Controller |
| 0x003E7368 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x003E7380 | `TCVoiceMemosAlert` | Known | Controller |
| 0x003E7394 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x003E73AC | `TCVoiceMemosPlayback` | Known | Controller |
| 0x003E73CC | `TSilverCntlr` | Known | Controller |
| 0x003E742C | `TSilverCntlr` | Known | Controller |
| 0x003E7498 | `TSilverCntlr` | Known | Controller |
| 0x003E87C0 | `TSilverCntlr` | Known | Controller |
| 0x003E88CC | `TSilverCntlr` | Known | Controller |
| 0x003F1144 | `TSilverCntlr` | Known | Controller |
| 0x003F1164 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x003F117C | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x003F1198 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x003F11B8 | `TCAddressViewerDetails` | Known | Controller |
| 0x003F11D0 | `TSilverCntlr` | Known | Controller |
| 0x003F11F0 | `TSilverCalendarCntlr_Base` | Known | Controller |
| 0x003F120C | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x003F1230 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x003F1254 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x003F1274 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x003F1298 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x003F12B8 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x003F12DC | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x003F14B4 | `TSilverCntlr` | Known | Controller |
| 0x003F14D4 | `TC_LockDialog` | Known | Controller |
| 0x003F14E4 | `TC_LockScreen` | Known | Controller |
| 0x003F14F4 | `TC_LockediPod` | Known | Controller |
| 0x003F1504 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003F1528 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003F1648 | `TSilverCntlr` | Known | Controller |
| 0x003F177C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003F1798 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003F17B8 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003F17D8 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003F1800 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003F1824 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003F184C | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003F186C | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003F188C | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003F18AC | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003F18CC | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003F18F4 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003F191C | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x003F193C | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003F195C | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003F197C | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003F19A0 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003F19C0 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003F19E4 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003F1A0C | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003F1A38 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003F1A58 | `TCRentalNotification` | Known | Controller |
| 0x003F1A70 | `TCRentalInfo` | Known | Controller |
| 0x003F1A80 | `TCRentalConfirmDelete` | Known | Controller |
| 0x003F1A98 | `TCRentalDispatcher` | Known | Controller |
| 0x003F2388 | `TSilverCntlr` | Known | Controller |
| 0x003F244C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003F2468 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003F2488 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003F24A8 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003F24D0 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003F24F4 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003F251C | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003F253C | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003F255C | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003F257C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003F259C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003F25C4 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003F25EC | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x003F2634 | `TCSlideshowTVOut` | Known | Controller |
| 0x003F2648 | `TCSlideshowLCD` | Known | Controller |
| 0x003F2658 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x003F2670 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x003F2690 | `TSilverCntlr` | Known | Controller |
| 0x003F26BC | `TSilverCntlr` | Known | Controller |
| 0x003F26DC | `TCUnsupported` | Known | Controller |
| 0x003F26FC | `TSilverCntlr` | Known | Controller |
| 0x003F273C | `TSilverCntlr` | Known | Controller |
| 0x003F275C | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x003F2778 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x003F2790 | `TSilverCntlr` | Known | Controller |
| 0x003F27B0 | `TCSpeakers` | Known | Controller |
| 0x003F27BC | `TCEQSetting` | Known | Controller |
| 0x003F27DC | `TSilverCntlr` | Known | Controller |
| 0x003F2844 | `TSilverCntlr` | Known | Controller |
| 0x003F2864 | `TCExtrasMenu` | Known | Controller |
| 0x003F2874 | `TCGamesMenu` | Known | Controller |
| 0x003F2880 | `TCGameScreen` | Known | Controller |
| 0x003F2890 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x003F28B0 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x003F28D0 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x003F28F0 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x003F2914 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003F2930 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003F2950 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003F2970 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003F2998 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003F29BC | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003F29E4 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003F2A04 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003F2A24 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003F2A44 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003F2A64 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003F2A8C | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003F2AB4 | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x003F2AD4 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003F2AF4 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003F2B14 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003F2B38 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003F2B58 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003F2B7C | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003F2BA4 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003F2BD0 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003F2BF0 | `TCRentalNotification` | Known | Controller |
| 0x003F2C08 | `TCRentalInfo` | Known | Controller |
| 0x003F2C18 | `TCRentalConfirmDelete` | Known | Controller |
| 0x003F2C30 | `TCRentalDispatcher` | Known | Controller |
| 0x003F2C44 | `TSilverGlobalCntlr` | Known | Controller |
| 0x003F2C58 | `TSilverTrainerCntlr` | Known | Controller |
| 0x00478084 | `TCCCCCCCCCCCCCCCCCCCCCCT` | Known | Controller |
| 0x0071A42E | `TCNotesDispatcher"` | Known | Controller |
| 0x0071A4ED | `TCLockChosenDispatcher"` | Known | Controller |
| 0x0071A5B0 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x00724695 | `TCNotesDispatcher"` | Known | Controller |
| 0x007247F7 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x00739B08 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x00739B20 | `TCAddressViewerDetails` | Known | Controller |
| 0x00739B38 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x00739B54 | `TCAlarmMenu` | Known | Controller |
| 0x00739B60 | `TCSleepTimerMenuTCAlarmPropertiesMenu` | Known | Controller |
| 0x00739B88 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x00739BA8 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x00739BC4 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00739BE0 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00739BFC | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00739C18 | `TCAlarmDatePicker` | Known | Controller |
| 0x00739C2C | `TCAlarmDatePicker` | Known | Controller |
| 0x00739C40 | `TCAlarmTriggeredTSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x00739C6C | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x00739C90 | `TSilverCalendarCntlr_MonthViewerTSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x00739CD0 | `TSilverCalendarCntlr_EventViewerTSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x00739D10 | `TSilverCalendarCntlr_EventViewerTSilverCntlrTCClockRegionMenu` | Known | Controller |
| 0x00739D50 | `TCClockCityMenu` | Known | Controller |
| 0x00739D60 | `TCClockCityMenu` | Known | Controller |
| 0x00739D70 | `TCClockCityMenu` | Known | Controller |
| 0x00739D80 | `TCClockCityMenu` | Known | Controller |
| 0x00739D90 | `TCClockCityMenu` | Known | Controller |
| 0x00739DA0 | `TCClockCityMenu` | Known | Controller |
| 0x00739DB0 | `TCClockCityMenu` | Known | Controller |
| 0x00739DC0 | `TCClockCityMenu` | Known | Controller |
| 0x00739DD0 | `TCClock` | Known | Controller |
| 0x00739DE8 | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDe` | Known | Controller |
| 0x00739E40 | `TCGamesMenu` | Known | Controller |
| 0x00739E4C | `TCGameScreenTC_LockediPod` | Known | Controller |
| 0x00739E68 | `TC_LockDialog` | Known | Controller |
| 0x00739E78 | `TC_LockScreen` | Known | Controller |
| 0x00739E88 | `TC_VolumeLimitLockScreenTSilverCntlrTSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x00739ECC | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00739EEC | `TSilverMediaListCntlr_AudiobooksTSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x00739F34 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x00739F50 | `TSilverMediaListCntlr_AlbumsTSilverMediaListCntlr_Artists` | Known | Controller |
| 0x00739F8C | `TSilverMediaListCntlr_GenresTSilverMediaListCntlr_Composers` | Known | Controller |
| 0x00739FC8 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x00739FE8 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0073A010 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x0073A030 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x0073A050 | `TSilverMediaListCntlr_TVEpisodesTSilverMediaListCntlr_MoviesTSilverMediaListCntl` | Known | Controller |
| 0x0073A0AC | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0073A0D4 | `TSilverMediaListCntlr_VideoPlaylistsTSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x0073A118 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x0073A144 | `TSilverCntlrTSilverMainMediaListCntlr_VideosTSilverSettingsVideoCntlr` | Known | Controller |
| 0x0073A18C | `TCFirstBoot` | Known | Controller |
| 0x0073A198 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x0073A1B8 | `TSilverMediaListCntlr_GeniusTSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0073A298 | `TCRentalInfoTCRentalConfirmDelete` | Known | Controller |
| 0x0073A2BC | `TSilverCntlrTCRentalNotificationTCRentalNotificationTCRentalNotificationTCNotesL` | Known | Controller |
| 0x0073A314 | `TCNotesList` | Known | Controller |
| 0x0073A320 | `TCNotesList` | Known | Controller |
| 0x0073A32C | `TCNotesContents` | Known | Controller |
| 0x0073A33C | `TCNotesContents` | Known | Controller |
| 0x0073A34C | `TCNotesContents` | Known | Controller |
| 0x0073A35C | `TCNotesContents` | Known | Controller |
| 0x0073A418 | `TCSlideshowLCD` | Known | Controller |
| 0x0073A428 | `TCSlideshowTVOutTCSlideshow_TVOutAskTCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x0073A478 | `TCRemoteUI` | Known | Controller |
| 0x0073A484 | `TCUnsupported` | Known | Controller |
| 0x0073A494 | `TSilverSettingsMenuListCntlrTSilverSettingsMenuListCntlrTSilverSettingsMenuListC` | Known | Controller |
| 0x0073A4FC | `TCSettings_MusicMenuTCSettings_VolumeLimit` | Known | Controller |
| 0x0073A528 | `TCSettings_Brightness` | Known | Controller |
| 0x0073A540 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0073A55C | `TCSettings_AudiobookSettingsTCSettings_RadioRegions` | Known | Controller |
| 0x0073A590 | `TCSettings_EQ` | Known | Controller |
| 0x0073A5A0 | `TSilverSettingsMenuListCntlrTSilverCntlrTCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0073A5E8 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0073A604 | `TCSettings_MainMenu` | Known | Controller |
| 0x0073A618 | `TCSettings_MusicMenuTCSportTimerTCSportTimerMenuTCSportTimerSessionScreen` | Known | Controller |
| 0x0073A664 | `TSilverCntlrTUnitTestSuiteCntlr` | Known | Controller |
| 0x0073A6E4 | `TCVoiceMemosTCVoiceMemosAlert` | Known | Controller |
| 0x0073A704 | `TCVoiceMemosAlert` | Known | Controller |
| 0x0073A718 | `TCVoiceMemosAlert` | Known | Controller |
| 0x0073A744 | `TCEQSetting` | Known | Controller |
| 0x0073A8B2 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x0073BE11 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x00741B78 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00741BD6 | `TCNotesDispatcher` | Known | Controller |
| 0x00743914 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00743972 | `TCNotesDispatcher` | Known | Controller |
| 0x007456B0 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074570E | `TCNotesDispatcher` | Known | Controller |
| 0x0074744C | `TCLockChosenDispatcher` | Known | Controller |
| 0x007474AA | `TCNotesDispatcher` | Known | Controller |
| 0x007491E8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00749246 | `TCNotesDispatcher` | Known | Controller |
| 0x0074AF84 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074AFE2 | `TCNotesDispatcher` | Known | Controller |
| 0x0074CD20 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074CD7E | `TCNotesDispatcher` | Known | Controller |
| 0x0074EABC | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074EB1A | `TCNotesDispatcher` | Known | Controller |
| 0x00750858 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007508B6 | `TCNotesDispatcher` | Known | Controller |
| 0x007525F4 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00752652 | `TCNotesDispatcher` | Known | Controller |
| 0x00754390 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007543EE | `TCNotesDispatcher` | Known | Controller |
| 0x0075612C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075618A | `TCNotesDispatcher` | Known | Controller |
| 0x00757EC8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00757F26 | `TCNotesDispatcher` | Known | Controller |
| 0x00759C64 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00759CC2 | `TCNotesDispatcher` | Known | Controller |
| 0x0075BA00 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075BA5E | `TCNotesDispatcher` | Known | Controller |
| 0x0075D79C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075D7FA | `TCNotesDispatcher` | Known | Controller |
| 0x0075F538 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075F596 | `TCNotesDispatcher` | Known | Controller |
| 0x007612D4 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00761332 | `TCNotesDispatcher` | Known | Controller |
| 0x00763070 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007630CE | `TCNotesDispatcher` | Known | Controller |
| 0x00764E0C | `TCLockChosenDispatcher` | Known | Controller |
| 0x00764E6A | `TCNotesDispatcher` | Known | Controller |
| 0x00766BA8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00766C06 | `TCNotesDispatcher` | Known | Controller |
| 0x00768944 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007689A2 | `TCNotesDispatcher` | Known | Controller |
| 0x0076A6E0 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076A73E | `TCNotesDispatcher` | Known | Controller |
| 0x0076C47C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076C4DA | `TCNotesDispatcher` | Known | Controller |
| 0x0076E218 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076E276 | `TCNotesDispatcher` | Known | Controller |
| 0x0076FFB4 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00770012 | `TCNotesDispatcher` | Known | Controller |
| 0x00771D50 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00771DAE | `TCNotesDispatcher` | Known | Controller |
| 0x00773AEC | `TCLockChosenDispatcher` | Known | Controller |
| 0x00773B4A | `TCNotesDispatcher` | Known | Controller |
| 0x00775888 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007758E6 | `TCNotesDispatcher` | Known | Controller |
| 0x00777624 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00777682 | `TCNotesDispatcher` | Known | Controller |
| 0x007793C0 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077941E | `TCNotesDispatcher` | Known | Controller |
| 0x0077B15C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077B1BA | `TCNotesDispatcher` | Known | Controller |
| 0x0077CEF8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077CF56 | `TCNotesDispatcher` | Known | Controller |
| 0x0077EC94 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077ECF2 | `TCNotesDispatcher` | Known | Controller |
| 0x00780A30 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00780A8E | `TCNotesDispatcher` | Known | Controller |
| 0x007827CC | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078282A | `TCNotesDispatcher` | Known | Controller |
| 0x00784568 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007845C6 | `TCNotesDispatcher` | Known | Controller |
| 0x0079221C | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x007924DE | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x00792D14 | `TCRentalDispatcher` | Known | Controller |
| 0x007935CC | `TCRentalDispatcher` | Known | Controller |
| 0x00793E84 | `TCRentalDispatcher` | Known | Controller |
| 0x0079473C | `TCRentalDispatcher` | Known | Controller |
| 0x00794FF4 | `TCRentalDispatcher` | Known | Controller |
| 0x007958AC | `TCRentalDispatcher` | Known | Controller |
| 0x00796164 | `TCRentalDispatcher` | Known | Controller |
| 0x00796A1C | `TCRentalDispatcher` | Known | Controller |
| 0x008DDC84 | `TCMockupModeNavScreen` | Known | Controller |
| 0x008DDC9C | `TSilverCntlr` | Known | Controller |
| 0x008DDCBC | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x008DDD0C | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x008DDD2C | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x008DDD4C | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x008DDD70 | `TCExtrasMenu` | Known | Controller |
| 0x008DDE80 | `TSilverCntlr` | Known | Controller |
| 0x008DDEA0 | `TCSlideshowTVOut` | Known | Controller |
| 0x008DDEB4 | `TCSlideshowLCD` | Known | Controller |
| 0x008DDEC4 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x008DDEDC | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x008DDEFC | `TSilverGlobalCntlr` | Known | Controller |
| 0x008DDF2C | `TSilverCntlr` | Known | Controller |
| 0x008DDFA8 | `TCSlideshowTVOut` | Known | Controller |
| 0x008DDFBC | `TCSlideshowLCD` | Known | Controller |
| 0x008DDFCC | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x008DDFE4 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x008DE004 | `TSilverCntlr` | Known | Controller |
| 0x008DE04C | `TSilverCntlr` | Known | Controller |
| 0x008DE06C | `TCGamesMenu` | Known | Controller |
| 0x008DE078 | `TCGameScreen` | Known | Controller |
| 0x0099BBDC | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |

---

## 2. Settings (Toggle/Show/TC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001281A0 | `ShowSetting_EQ` | Known | User setting |
| 0x001CFD88 | `ToggleSetting_Repeat` | Known | User setting |
| 0x001CFDA4 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001CFDBC | `ToggleSetting_TVOut` | Known | User setting |
| 0x001CFDD0 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x001F8B08 | `ShowSetting_Backlight` | Known | User setting |
| 0x0020DACC | `ToggleSetting_Shuffle` | Known | User setting |
| 0x0020DAE8 | `ToggleSetting_Repeat` | Known | User setting |
| 0x0020DB00 | `ToggleSetting_SortBy` | Known | User setting |
| 0x0020DB18 | `ToggleSetting_ClassicUI` | Known | User setting |
| 0x0020DB30 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x0020DB4C | `ToggleSetting_Clicker` | Known | User setting |
| 0x0020DB64 | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x0020DB84 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x0020DBA0 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x0020DBBC | `ShowSetting_Shuffle` | Known | User setting |
| 0x0020DD68 | `ShowSetting_Repeat` | Known | User setting |
| 0x0020DD7C | `ShowSetting_About` | Known | User setting |
| 0x0020DD90 | `ShowSetting_MainMenu` | Known | User setting |
| 0x0020DDA8 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x0020DDC0 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x0020DDD8 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x0020DDF4 | `ShowSetting_Brightness` | Known | User setting |
| 0x0020DE0C | `ShowSetting_Audiobooks` | Known | User setting |
| 0x0020DE24 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x0020DE40 | `ShowSetting_EQ` | Known | User setting |
| 0x0020DE50 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x0020DFEC | `ShowSetting_Clicker` | Known | User setting |
| 0x0020E000 | `ShowSetting_DateAndTime` | Known | User setting |
| 0x0020E018 | `ShowSetting_SortBy` | Known | User setting |
| 0x0020E02C | `ShowSetting_ClassicUI` | Known | User setting |
| 0x0020E044 | `ShowSetting_Language` | Known | User setting |
| 0x0020E05C | `ShowSetting_Legal` | Known | User setting |
| 0x0020E070 | `ShowSetting_ResetAll` | Known | User setting |
| 0x007234A5 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x00723555 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x00725CA2 | `ShowSetting_About` | Known | User setting |
| 0x00725D44 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x00725D88 | `ShowSetting_Shuffle` | Known | User setting |
| 0x00725DFF | `ToggleSetting_Repeat` | Known | User setting |
| 0x00725E42 | `ShowSetting_Repeat` | Known | User setting |
| 0x00725F4C | `ShowSetting_MainMenu` | Known | User setting |
| 0x0072605C | `ShowSetting_MusicMenu` | Known | User setting |
| 0x00726124 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x007261EE | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x00726306 | `ShowSetting_Brightness` | Known | User setting |
| 0x0072643C | `ShowSetting_Audiobooks` | Known | User setting |
| 0x0072654D | `ShowSetting_RadioRegions` | Known | User setting |
| 0x0072664E | `ShowSetting_EQ` | Known | User setting |
| 0x007266BB | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x00726702 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x0072677F | `ToggleSetting_Clicker` | Known | User setting |
| 0x007267C3 | `ShowSetting_Clicker` | Known | User setting |
| 0x0072692A | `ToggleSetting_SortBy` | Known | User setting |
| 0x0072696D | `ShowSetting_SortBy` | Known | User setting |
| 0x00726A6E | `ShowSetting_Language` | Known | User setting |
| 0x00726B7E | `ShowSetting_Legal` | Known | User setting |
| 0x00726CAF | `ShowSetting_ResetAll` | Known | User setting |
| 0x00726E20 | `ShowSetting_Backlight` | Known | User setting |
| 0x00726ED0 | `ShowSetting_Backlight` | Known | User setting |
| 0x00726F80 | `ShowSetting_Backlight` | Known | User setting |
| 0x00727031 | `ShowSetting_Backlight` | Known | User setting |
| 0x007270E2 | `ShowSetting_Backlight` | Known | User setting |
| 0x00727193 | `ShowSetting_Backlight` | Known | User setting |
| 0x00727247 | `ShowSetting_Backlight` | Known | User setting |
| 0x007272F6 | `ShowSetting_EQ` | Known | User setting |
| 0x0072736B | `ShowSetting_Language` | Known | User setting |
| 0x007AF06C | `ToggleSetting_Repeat` | Known | User setting |
| 0x007AF0A6 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x007AF168 | `ToggleSetting_TVOut` | Known | User setting |
| 0x007AF1A1 | `ToggleSetting_TVSignal` | Known | User setting |

---

## 3. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00143838 | `MockupMode/MockupMode.xml` | Hidden | Developer Tool |
| 0x00143D38 | `MockupMode/` | Hidden | Developer Tool |
| 0x0024921C | `Channel UnitTests` | Hidden | Developer Tool |
| 0x002A2979 | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x002A29BC | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x002A29D1 | `RTXCbug> ` | Hidden | Developer Tool |
| 0x002A33AD | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x002BCD48 | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x0038537D | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x00385445 | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x003E38BD | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x0073A684 | `TUnitTestSuiteTestsCntlrTSilverCntlrTCVoiceMemosMainMenuTCVoiceMemosMenuTCVoiceM` | Hidden | Developer Tool |
| 0x007D6190 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0081308C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00825D74 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0083DC58 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008505A8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0085A574 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00864210 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008799E4 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008838F8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008AAAB0 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008C98A8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008D2F28 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0095F2ED | `10TCDemoMode` | Hidden | Demo/Retail Mode |
| 0x0095FC64 | `21TCMockupModeNavScreen` | Hidden | Developer Tool |
| 0x00960124 | `27TSilverCntlrTransitionAddonI10TCDemoModeE` | Hidden | Demo/Retail Mode |
| 0x0098DDCB | `Debug_ListItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x0098DDE3 | `Debug_MenuItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x0098E4E8 | `DemoMode_Background_Image` | Hidden | Demo/Retail Mode |
| 0x0098F0D6 | `DemoMode_Icon_Image` | Hidden | Demo/Retail Mode |
| 0x00990C98 | `Debug_Menu_Title` | Hidden | Debug/Diagnostic |
| 0x00990CBD | `UnitTest_Menu_Title` | Hidden | Developer Tool |
| 0x00999794 | `UnitTestModel` | Hidden | Developer Tool |
| 0x0099A173 | `UnitTest_ListItem` | Hidden | Developer Tool |
| 0x0099B289 | `DemoMode_View_Main` | Hidden | Demo/Retail Mode |
| 0x0099B45E | `DemoMode_Icon` | Hidden | Demo/Retail Mode |
| 0x0099C244 | `UnitTestApp` | Hidden | Developer Tool |
| 0x0099C7F6 | `Debug_ListItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x0099C811 | `Debug_MenuItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x0099CF27 | `DemoMode_Text_Color` | Hidden | Demo/Retail Mode |
| 0x0099D33C | `Debug_ListItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x0099D353 | `Debug_MenuItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009A13D2 | `Debug_ListItem_UnitTest` | Hidden | Developer Tool |
| 0x009A13EA | `Debug_MenuItem_UnitTest` | Hidden | Developer Tool |
| 0x009A57E4 | `Debug_ListItem_Memory` | Hidden | Debug/Diagnostic |
| 0x009A57FA | `Debug_MenuItem_Memory` | Hidden | Debug/Diagnostic |

---

## 4. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000067BB | `"MeCCADecode` | Known | Audio system |
| 0x0013985C | `AudioCodecs` | Known | Audio system |
| 0x001511C0 | `MeCCA_RecordingBuffer` | Known | Audio system |
| 0x0017F98C | `MeCCA_PCM_Output.wav` | Known | Audio system |
| 0x00198C78 | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x001A2594 | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x001A279C | `MeCCAVideoDecode` | Known | Audio system |
| 0x008E9DF0 | `MeCCA_StreamCache` | Known | Audio system |

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
| 0x0013F0B8 | `HandleNotesPop` | Known | Event handler |
| 0x0013F0CC | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0013FFD8 | `HandleSelect` | Known | Event handler |
| 0x0013FFEC | `HandleWheel` | Known | Event handler |
| 0x0013FFF8 | `HandleImageNext` | Known | Event handler |
| 0x00140008 | `HandleImagePrev` | Known | Event handler |
| 0x00140018 | `HandleImageLast` | Known | Event handler |
| 0x00140028 | `HandleImageFirst` | Known | Event handler |
| 0x0014003C | `HandlePlayPause` | Known | Event handler |
| 0x0014004C | `HandlePlay` | Known | Event handler |
| 0x00140058 | `HandlePause` | Known | Event handler |
| 0x00140064 | `HandleMikeyCenter` | Known | Event handler |
| 0x00154FE8 | `HandleSelectCity` | Known | Event handler |
| 0x00155000 | `HandleHighlightCity` | Known | Event handler |
| 0x001560EC | `HandleWantPopFlow` | Known | Event handler |
| 0x00156104 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x00156120 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0015613C | `HandleFlowNext` | Known | Event handler |
| 0x0015614C | `HandleFlowPrev` | Known | Event handler |
| 0x0015615C | `HandleFlowWheel` | Known | Event handler |
| 0x0015616C | `HandleAlbumSelected` | Known | Event handler |
| 0x00156180 | `HandlePlayPause` | Known | Event handler |
| 0x00156190 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00181828 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00181C18 | `HandleSelect` | Known | Event handler |
| 0x00182B00 | `HandleSelect` | Known | Event handler |
| 0x00182B14 | `HandleWheel` | Known | Event handler |
| 0x00182B20 | `HandleImageNext` | Known | Event handler |
| 0x00182B30 | `HandleImagePrev` | Known | Event handler |
| 0x00182B40 | `HandleImageLast` | Known | Event handler |
| 0x00182B50 | `HandleImageFirst` | Known | Event handler |
| 0x00182B64 | `HandlePlayPause` | Known | Event handler |
| 0x00182B74 | `HandlePlay` | Known | Event handler |
| 0x00182B80 | `HandlePause` | Known | Event handler |
| 0x00182B8C | `HandleMikeyCenter` | Known | Event handler |
| 0x00183034 | `HandleNew` | Known | Event handler |
| 0x00183044 | `HandleClear` | Known | Event handler |
| 0x00183050 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x0018306C | `HandleSelectIndexedSession` | Known | Event handler |
| 0x0018337C | `HandleWheel` | Known | Event handler |
| 0x0018338C | `HandleArrowUp` | Known | Event handler |
| 0x0018339C | `HandleArrowDown` | Known | Event handler |
| 0x001858C0 | `HandleHiliteAlbum` | Known | Event handler |
| 0x001858D8 | `HandleBrowseAlbum` | Known | Event handler |
| 0x001858EC | `HandlePlayPause` | Known | Event handler |
| 0x0019C298 | `HandleSelect` | Known | Event handler |
| 0x0019C428 | `HandleSelectRegion` | Known | Event handler |
| 0x0019C7A0 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x0019C7BC | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x0019C7D8 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x001B2120 | `HandleImageWheel` | Known | Event handler |
| 0x001B2138 | `HandlePlayPause` | Known | Event handler |
| 0x001B2148 | `HandleBrowseLarge` | Known | Event handler |
| 0x001B215C | `HandleBrowseSmall` | Known | Event handler |
| 0x001B2170 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x001B2188 | `HandleImageNext` | Known | Event handler |
| 0x001B2198 | `HandleImagePrev` | Known | Event handler |
| 0x001B21A8 | `HandleHilite` | Known | Event handler |
| 0x001B21B8 | `HandleImageLast` | Known | Event handler |
| 0x001B21C8 | `HandleImageFirst` | Known | Event handler |
| 0x001B21DC | `HandleScreenNext` | Known | Event handler |
| 0x001B21F0 | `HandleScreenPrev` | Known | Event handler |
| 0x001B4AB8 | `HandlePlayPause` | Known | Event handler |
| 0x001B4ACC | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x001B4AE8 | `HandleNext` | Known | Event handler |
| 0x001B4AF4 | `HandleNextPressAndHold` | Known | Event handler |
| 0x001B4B0C | `HandlePrevious` | Known | Event handler |
| 0x001B4B1C | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x001B4B38 | `HandleRemotePlayPause` | Known | Event handler |
| 0x001B4B50 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x001B4B74 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x001B4B8C | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x001B4BA4 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x001B4D48 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x001B4D60 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x001B4D78 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x001B4D94 | `HandleRemoteStop` | Known | Event handler |
| 0x001B4DA8 | `HandleRemotePlay` | Known | Event handler |
| 0x001B4DBC | `HandleRemotePause` | Known | Event handler |
| 0x001B4DD0 | `HandleRemoteMute` | Known | Event handler |
| 0x001B4DE4 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x001B4DFC | `HandleRemotePrevChapter` | Known | Event handler |
| 0x001B4E14 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x001B4E30 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x001B5038 | `HandleRemoteShuffle` | Known | Event handler |
| 0x001B504C | `HandleRemoteRepeat` | Known | Event handler |
| 0x001B5060 | `HandleRemoteOn` | Known | Event handler |
| 0x001B5074 | `HandleRemoteOff` | Known | Event handler |
| 0x001B5084 | `HandleRemoteBacklight` | Known | Event handler |
| 0x001B509C | `HandleRemoteFFDown` | Known | Event handler |
| 0x001B50B0 | `HandleRemoteFFUp` | Known | Event handler |
| 0x001B50C4 | `HandleRemoteRewDown` | Known | Event handler |
| 0x001B50D8 | `HandleRemoteRewUp` | Known | Event handler |
| 0x001B50EC | `HandleRemoteMenuDown` | Known | Event handler |
| 0x001B5104 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x001B5118 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x001B5130 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x001B52E0 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x001B52F8 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x001B5310 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x001B532C | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x001B5344 | `HandleRemoteEvent` | Known | Event handler |
| 0x001B5358 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x001B5374 | `HandleAudioPlayPause` | Known | Event handler |
| 0x001B538C | `HandleAudioNext` | Known | Event handler |
| 0x001B539C | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x001B53B8 | `HandleAudioPrevious` | Known | Event handler |
| 0x001B53CC | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x001B555C | `HandleAudioNextAlbum` | Known | Event handler |
| 0x001B5574 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x001B558C | `HandleAudioVolumeDown` | Known | Event handler |
| 0x001B55A4 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x001B55B8 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x001B55D0 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x001B55E8 | `HandleAudioStop` | Known | Event handler |
| 0x001B55F8 | `HandleAudioPlay` | Known | Event handler |
| 0x001B5608 | `HandleAudioPause` | Known | Event handler |
| 0x001B561C | `HandleAudioMute` | Known | Event handler |
| 0x001B562C | `HandleAudioNextChapter` | Known | Event handler |
| 0x001B5644 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x001B5830 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x001B5848 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x001B5860 | `HandleAudioShuffle` | Known | Event handler |
| 0x001B5874 | `HandleAudioRepeat` | Known | Event handler |
| 0x001B5888 | `HandleAudioFFDown` | Known | Event handler |
| 0x001B589C | `HandleAudioFFUp` | Known | Event handler |
| 0x001B58AC | `HandleAudioRewDown` | Known | Event handler |
| 0x001B58C0 | `HandleAudioRewUp` | Known | Event handler |
| 0x001B58D4 | `HandleVideoPlayPause` | Known | Event handler |
| 0x001B58EC | `HandleVideoNext` | Known | Event handler |
| 0x001B58FC | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x001B5918 | `HandleVideoPrevious` | Known | Event handler |
| 0x001B592C | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x001B5AF0 | `HandleVideoStop` | Known | Event handler |
| 0x001B5B00 | `HandleVideoPlay` | Known | Event handler |
| 0x001B5B10 | `HandleVideoPause` | Known | Event handler |
| 0x001B5B24 | `HandleVideoFFDown` | Known | Event handler |
| 0x001B5B38 | `HandleVideoFFUp` | Known | Event handler |
| 0x001B5B48 | `HandleVideoRewDown` | Known | Event handler |
| 0x001B5B5C | `HandleVideoRewUp` | Known | Event handler |
| 0x001B5B70 | `HandleVideoNextChapter` | Known | Event handler |
| 0x001B5B88 | `HandleVideoPrevChapter` | Known | Event handler |
| 0x001B5BA0 | `HandleVideoNextFrame` | Known | Event handler |
| 0x001B5BB8 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x001B5BD0 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x001B5BEC | `HandleMikeyCenter` | Known | Event handler |
| 0x001B5D4C | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x001B5D6C | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x001B5D8C | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x001B5DB0 | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x001B5DD0 | `HandleMikeyAllUp` | Known | Event handler |
| 0x001B5DE4 | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x001B5DF8 | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x001B5E10 | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x001B5E28 | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x001C2BF0 | `HandleMainMenu` | Known | Event handler |
| 0x001C4D3C | `HandleLoadingCancelled` | Known | Event handler |
| 0x001C7770 | `HandlePowerSongSelected` | Known | Event handler |
| 0x001C778C | `HandlePowerSongChosen` | Known | Event handler |
| 0x001C77A4 | `HandleSelectTimedWorkout` | Known | Event handler |
| 0x001CDCF0 | `HandleSelect` | Known | Event handler |
| 0x001CDF98 | `HandleMusicMenu` | Known | Event handler |
| 0x001CE258 | `HandleSelect` | Known | Event handler |
| 0x001CE55C | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x001CE57C | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x001CEA38 | `HandleWheel` | Known | Event handler |
| 0x001CEA48 | `HandlePlayPause` | Known | Event handler |
| 0x001CEA58 | `HandleSelectDown` | Known | Event handler |
| 0x001CEA6C | `HandleNext` | Known | Event handler |
| 0x001CEA78 | `HandlePrevious` | Known | Event handler |
| 0x001CEA88 | `HandleNextPushAndHold` | Known | Event handler |
| 0x001CEAA0 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x001CF194 | `HandleMenuSelection` | Known | Event handler |
| 0x001CF1A8 | `HandleViewAlbum` | Known | Event handler |
| 0x001CF1B8 | `HandleViewArtist` | Known | Event handler |
| 0x001CF1CC | `HandleViewCompilation` | Known | Event handler |
| 0x001CF1E4 | `HandleStartGenius` | Known | Event handler |
| 0x001DBDF4 | `HandleFrequencyChosen` | Known | Event handler |
| 0x001DBE0C | `HandleDateChosen` | Known | Event handler |
| 0x001DBE20 | `HandleTimeChosen` | Known | Event handler |
| 0x001DBE34 | `HandleSoundChosen` | Known | Event handler |
| 0x001DBE48 | `HandleLabelChosen` | Known | Event handler |
| 0x001DBE5C | `HandleDeleteChosen` | Known | Event handler |
| 0x001DCF3C | `HandleSelect` | Known | Event handler |
| 0x001E1858 | `HandlePrev` | Known | Event handler |
| 0x001E1868 | `HandleNext` | Known | Event handler |
| 0x001E1874 | `HandlePlayPause` | Known | Event handler |
| 0x001E9230 | `HandleNextContact` | Known | Event handler |
| 0x001E9248 | `HandlePreviousContact` | Known | Event handler |
| 0x001F0DF8 | `HandleItemSelected` | Known | Event handler |
| 0x001F0FF0 | `HandleRadioRegion` | Known | Event handler |
| 0x001F11D8 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x001F544C | `HandlePlayPause` | Known | Event handler |
| 0x001F8DE4 | `HandleDelete` | Known | Event handler |
| 0x001F8DF8 | `HandleSelectLozinch` | Known | Event handler |
| 0x001F90A0 | `HandleSelect` | Known | Event handler |
| 0x001F936C | `HandleTVOutChanged` | Known | Event handler |
| 0x001F9384 | `HandleTVSignalChanged` | Known | Event handler |
| 0x001F939C | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x001F93BC | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x001F93DC | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x001F9400 | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x001F9420 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x001FC280 | `HandleSelectKey` | Known | Event handler |
| 0x001FC428 | `HandleSelect` | Known | Event handler |
| 0x001FD1A4 | `HandlePlayPause` | Known | Event handler |
| 0x001FD1B8 | `HandleWheel` | Known | Event handler |
| 0x001FD1C4 | `HandleWheelRating` | Known | Event handler |
| 0x001FD1D8 | `HandleWheelScrub` | Known | Event handler |
| 0x001FD1EC | `HandleWheelVolume` | Known | Event handler |
| 0x001FD2AC | `HandleMenuKey` | Known | Event handler |
| 0x001FD318 | `HandleMenuLongpress` | Known | Event handler |
| 0x001FD32C | `HandleRentalWarningChoice` | Known | Event handler |
| 0x001FDF34 | `HandleSelect` | Known | Event handler |
| 0x001FE82C | `HandleLeaveAlarm` | Known | Event handler |
| 0x001FF744 | `HandleSelect` | Known | Event handler |
| 0x001FF758 | `HandleHilite` | Known | Event handler |
| 0x001FF768 | `HandlePlayPause` | Known | Event handler |
| 0x001FF778 | `HandleAddToOTG` | Known | Event handler |
| 0x001FF788 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x001FF7A8 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00202830 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x00203040 | `HandleSelect` | Known | Event handler |
| 0x00203054 | `HandleWheel` | Known | Event handler |
| 0x00203060 | `HandleWheelProgress` | Known | Event handler |
| 0x00203074 | `HandleSelectProgress` | Known | Event handler |
| 0x0020308C | `HandleSelectVolume` | Known | Event handler |
| 0x002030A0 | `HandleSelectScrub` | Known | Event handler |
| 0x002030B4 | `HandleSelectGenius` | Known | Event handler |
| 0x002030C8 | `HandleSelectRating` | Known | Event handler |
| 0x002030DC | `HandleSelectExtraInfo` | Known | Event handler |
| 0x002030F4 | `HandleSelectChapterArt` | Known | Event handler |
| 0x0020310C | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x00203128 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x00203324 | `HandleWheelGenius` | Known | Event handler |
| 0x00203338 | `HandleWheelBrightness` | Known | Event handler |
| 0x002033A4 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x002033C4 | `HandlePushContextualMenu` | Known | Event handler |
| 0x002033E0 | `HandleAddToOTG` | Known | Event handler |
| 0x002033F0 | `HandleViewArtist` | Known | Event handler |
| 0x00203404 | `HandleViewAlbum` | Known | Event handler |
| 0x00203414 | `HandleViewCompilation` | Known | Event handler |
| 0x0020350C | `HandleStartGenius` | Known | Event handler |
| 0x00203520 | `HandleAudiobookSlower` | Known | Event handler |
| 0x00203538 | `HandleAudiobookFaster` | Known | Event handler |
| 0x00203550 | `HandleAudiobookNormal` | Known | Event handler |
| 0x00203568 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00205004 | `HandleStartGenius` | Known | Event handler |
| 0x00205358 | `HandleAudiobookSlower` | Known | Event handler |
| 0x00205370 | `HandleAudiobookNormal` | Known | Event handler |
| 0x00205388 | `HandleAudiobookFaster` | Known | Event handler |
| 0x002053A0 | `HandleStartGenius` | Known | Event handler |
| 0x002053B4 | `HandleAddToOTG` | Known | Event handler |
| 0x002053C4 | `HandleViewCompilation` | Known | Event handler |
| 0x002053DC | `HandleViewAlbum` | Known | Event handler |
| 0x002053EC | `HandleViewArtist` | Known | Event handler |
| 0x00205400 | `HandleCancel` | Known | Event handler |
| 0x00205E9C | `HandleSelect` | Known | Event handler |
| 0x00205EAC | `HandleSelectRating` | Known | Event handler |
| 0x00205EC0 | `HandleSelectProgress` | Known | Event handler |
| 0x00205ED8 | `HandleWheelProgress` | Known | Event handler |
| 0x00205EEC | `HandleSelectScrub` | Known | Event handler |
| 0x00205F00 | `HandleWheelBrightness` | Known | Event handler |
| 0x00205F18 | `HandleNextTrackGotoScrub` | Known | Event handler |
| 0x00205F34 | `HandlePrevTrackGotoScrub` | Known | Event handler |
| 0x00205F50 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x00208B38 | `HandleStartGenius` | Known | Event handler |
| 0x00208B50 | `HandleViewArtist` | Known | Event handler |
| 0x00208B64 | `HandleViewAlbum` | Known | Event handler |
| 0x00208B74 | `HandleViewCompilation` | Known | Event handler |
| 0x00208B8C | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00209510 | `HandleStartGenius` | Known | Event handler |
| 0x00209524 | `HandleAddToOTG` | Known | Event handler |
| 0x00209534 | `HandleViewCompilation` | Known | Event handler |
| 0x0020954C | `HandleViewAlbum` | Known | Event handler |
| 0x0020955C | `HandleViewArtist` | Known | Event handler |
| 0x00209570 | `HandleCancel` | Known | Event handler |
| 0x0020BF00 | `HandleAddToOTG` | Known | Event handler |
| 0x0020BF10 | `HandleCancel` | Known | Event handler |
| 0x0020C104 | `HandleStartGenius` | Known | Event handler |
| 0x0020C11C | `HandleViewAlbum` | Known | Event handler |
| 0x0020C12C | `HandleViewArtist` | Known | Event handler |
| 0x0020C140 | `HandleViewCompilation` | Known | Event handler |
| 0x0020C158 | `HandleSaveGeniusPlaylist` | Known | Event handler |
| 0x0020C174 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x0020C18C | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0020D0F4 | `HandleStartGenius` | Known | Event handler |
| 0x0020D108 | `HandleAddToOTG` | Known | Event handler |
| 0x0020D118 | `HandleViewCompilation` | Known | Event handler |
| 0x0020D130 | `HandleViewAlbum` | Known | Event handler |
| 0x0020D140 | `HandleViewArtist` | Known | Event handler |
| 0x0020D154 | `HandleCancel` | Known | Event handler |
| 0x0020D600 | `HandleAddToOTG` | Known | Event handler |
| 0x0020D610 | `HandleCancel` | Known | Event handler |
| 0x0020E0A8 | `HandleLanguage` | Known | Event handler |
| 0x0020E0B8 | `HandleResetAllSettings` | Known | Event handler |
| 0x0020E0D0 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x0020EA3C | `HandleSelect` | Known | Event handler |
| 0x0020EC6C | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x0020FB4C | `HandleAddToOTG` | Known | Event handler |
| 0x0020FB5C | `HandleCancel` | Known | Event handler |
| 0x00212644 | `HandleSelect` | Known | Event handler |
| 0x002127E0 | `HandleSelect` | Known | Event handler |
| 0x00212A80 | `HandleNextDay` | Known | Event handler |
| 0x00212A94 | `HandlePreviousDay` | Known | Event handler |
| 0x00213298 | `HandleMusicHilited` | Known | Event handler |
| 0x002132B0 | `HandleVideosHilited` | Known | Event handler |
| 0x002132C4 | `HandlePodcastsHilited` | Known | Event handler |
| 0x002132DC | `HandleGenericHilited` | Known | Event handler |
| 0x002132F4 | `HandlePhotosHilited` | Known | Event handler |
| 0x00213308 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x00213320 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x0021333C | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00213354 | `HandleArtistsHilited` | Known | Event handler |
| 0x0021336C | `HandleGenresHilited` | Known | Event handler |
| 0x00213380 | `HandleAlbumsHilited` | Known | Event handler |
| 0x00213394 | `HandleCompilationsHilited` | Known | Event handler |
| 0x00213564 | `HandleComposersHilited` | Known | Event handler |
| 0x0021357C | `HandleSongsHilited` | Known | Event handler |
| 0x00213590 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x002135A8 | `HandleGeniusHilited` | Known | Event handler |
| 0x002135BC | `HandleTVShowsHilited` | Known | Event handler |
| 0x002135D4 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x002135F0 | `HandleMoviesHilited` | Known | Event handler |
| 0x00213604 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00213620 | `HandleRentalsHilited` | Known | Event handler |
| 0x00213638 | `HandleMusicSelected` | Known | Event handler |
| 0x0021364C | `HandleVideosSelected` | Known | Event handler |
| 0x0021381C | `HandlePodcastsSelected` | Known | Event handler |
| 0x00213834 | `HandlePhotosSelected` | Known | Event handler |
| 0x0021384C | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00213864 | `HandleSongsSelected` | Known | Event handler |
| 0x00213878 | `HandleAlbumsSelected` | Known | Event handler |
| 0x00213890 | `HandleCompilationsSelected` | Known | Event handler |
| 0x002138AC | `HandleArtistsSelected` | Known | Event handler |
| 0x002138C4 | `HandleGenresSelected` | Known | Event handler |
| 0x002138DC | `HandleComposersSelected` | Known | Event handler |
| 0x002138F4 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x00213910 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00213AE4 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00213AFC | `HandleNowPlaying` | Known | Event handler |
| 0x00213B10 | `HandleGotoGenius` | Known | Event handler |
| 0x00213B24 | `HandleTVShowsSelected` | Known | Event handler |
| 0x00213B3C | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00213B58 | `HandleMoviesSelected` | Known | Event handler |
| 0x00213B70 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00213B90 | `HandleRentalsSelected` | Known | Event handler |
| 0x00213BA8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00213BC0 | `HandleLock` | Known | Event handler |
| 0x00213BCC | `HandleBacklightSelected` | Known | Event handler |
| 0x00213C2C | `HandleSleepSelected` | Known | Event handler |
| 0x00213C40 | `HandleNikePlusSelected` | Known | Event handler |
| 0x00216664 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00216C00 | `HandleAddToOTG` | Known | Event handler |
| 0x00216C10 | `HandleCancel` | Known | Event handler |
| 0x00216DE0 | `HandleWheel` | Known | Event handler |
| 0x00217C5C | `HandleAddToOTG` | Known | Event handler |
| 0x00217C6C | `HandleCancel` | Known | Event handler |
| 0x0021842C | `HandleAddToOTG` | Known | Event handler |
| 0x0021843C | `HandleCancel` | Known | Event handler |
| 0x00218DEC | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x00219044 | `HandleNextDay` | Known | Event handler |
| 0x00219058 | `HandlePreviousDay` | Known | Event handler |
| 0x002192A0 | `HandleSelect` | Known | Event handler |
| 0x0021953C | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00219A0C | `HandleAddToOTG` | Known | Event handler |
| 0x00219A1C | `HandleCancel` | Known | Event handler |
| 0x0021D0E4 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x0021D100 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x0021D118 | `HandleStartGenius` | Known | Event handler |
| 0x0021D12C | `HandleViewArtist` | Known | Event handler |
| 0x0021D140 | `HandleViewAlbum` | Known | Event handler |
| 0x0021D150 | `HandleViewCompilation` | Known | Event handler |
| 0x0021D168 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0021D184 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x0021D19C | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0021E4C4 | `HandleStartGenius` | Known | Event handler |
| 0x0021E4D8 | `HandleAddToOTG` | Known | Event handler |
| 0x0021E4E8 | `HandleViewCompilation` | Known | Event handler |
| 0x0021E500 | `HandleViewAlbum` | Known | Event handler |
| 0x0021E510 | `HandleViewArtist` | Known | Event handler |
| 0x0021E524 | `HandleCancel` | Known | Event handler |
| 0x0021EC98 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0021EEFC | `HandleAddToOTG` | Known | Event handler |
| 0x0021EF0C | `HandleCancel` | Known | Event handler |
| 0x0021F400 | `HandleSelect` | Known | Event handler |
| 0x0021FACC | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x00258AD8 | `HandleDeleteClock` | Known | Event handler |
| 0x00258AF0 | `HandleSelectClock` | Known | Event handler |
| 0x00258B04 | `HandleHilited` | Known | Event handler |
| 0x00258B14 | `HandleWheel` | Known | Event handler |
| 0x00258B20 | `HandleSelectLozinch` | Known | Event handler |
| 0x004130DA | `HandleAudioFFDown` | Known | Event handler |
| 0x00413103 | `HandleAudioFFUp` | Known | Event handler |
| 0x0041312E | `HandleAudioMute` | Known | Event handler |
| 0x00413161 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x00413196 | `HandleAudioNext` | Known | Event handler |
| 0x004131C6 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x004131FD | `HandleAudioNextChapter` | Known | Event handler |
| 0x00413237 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x0041326B | `HandleAudioPause` | Known | Event handler |
| 0x00413297 | `HandleAudioPlay` | Known | Event handler |
| 0x004132C5 | `HandleAudioPlayPause` | Known | Event handler |
| 0x004132FD | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x00413336 | `HandleAudioPrevious` | Known | Event handler |
| 0x0041336A | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x004133A1 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x004133DB | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x00413410 | `HandleAudioRepeat` | Known | Event handler |
| 0x0041343C | `HandleAudioRewDown` | Known | Event handler |
| 0x00413467 | `HandleAudioRewUp` | Known | Event handler |
| 0x00413496 | `HandleAudioShuffle` | Known | Event handler |
| 0x004134C4 | `HandleAudioStop` | Known | Event handler |
| 0x004134F5 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x0041352A | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x00413561 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x00413592 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x0041364B | `HandleNextPressAndHold` | Known | Event handler |
| 0x0041367C | `HandleNext` | Known | Event handler |
| 0x004136B4 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x004136EF | `HandlePlayPause` | Known | Event handler |
| 0x00413723 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x00413758 | `HandlePrevious` | Known | Event handler |
| 0x004137EA | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x00413832 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x0041387B | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x004138BD | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x004138F5 | `HandleMikeyCenter` | Known | Event handler |
| 0x00413928 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x0041395E | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x00413996 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x004139C8 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x004139FE | `HandleRemoteBacklight` | Known | Event handler |
| 0x00413A36 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x00413A70 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x00413AA9 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x00413ADE | `HandleRemoteEvent` | Known | Event handler |
| 0x00413B0A | `HandleRemoteFFDown` | Known | Event handler |
| 0x00413B35 | `HandleRemoteFFUp` | Known | Event handler |
| 0x00413B62 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x00413B91 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x00413BC0 | `HandleRemoteMute` | Known | Event handler |
| 0x00413BF2 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x00413C2B | `HandleRemoteNextChapter` | Known | Event handler |
| 0x00413C67 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x00413CA7 | `HandleRemoteOff` | Known | Event handler |
| 0x00413CD0 | `HandleRemoteOff` | Known | Event handler |
| 0x00413CFA | `HandleRemoteOn` | Known | Event handler |
| 0x00413D26 | `HandleRemotePause` | Known | Event handler |
| 0x00413D54 | `HandleRemotePlay` | Known | Event handler |
| 0x00413D92 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x00413DD3 | `HandleRemotePlayPause` | Known | Event handler |
| 0x00413E0A | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x00413E43 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x00413E7F | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x00413EB6 | `HandleRemoteRepeat` | Known | Event handler |
| 0x00413EE4 | `HandleRemoteRewDown` | Known | Event handler |
| 0x00413F11 | `HandleRemoteRewUp` | Known | Event handler |
| 0x00413F41 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x00413F74 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x00413FA8 | `HandleRemoteShuffle` | Known | Event handler |
| 0x00413FD8 | `HandleRemoteStop` | Known | Event handler |
| 0x00414008 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x0041403D | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x00414075 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x004140AC | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x004140E5 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x00414118 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x0041414D | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x00414180 | `HandleVideoFFDown` | Known | Event handler |
| 0x004141A9 | `HandleVideoFFUp` | Known | Event handler |
| 0x004141DC | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x00414211 | `HandleVideoNext` | Known | Event handler |
| 0x00414243 | `HandleVideoNextChapter` | Known | Event handler |
| 0x0041427A | `HandleVideoNextFrame` | Known | Event handler |
| 0x004142AB | `HandleVideoPause` | Known | Event handler |
| 0x004142D7 | `HandleVideoPlay` | Known | Event handler |
| 0x00414305 | `HandleVideoPlayPause` | Known | Event handler |
| 0x0041433D | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x00414376 | `HandleVideoPrevious` | Known | Event handler |
| 0x004143AC | `HandleVideoPrevChapter` | Known | Event handler |
| 0x004143E3 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x00414412 | `HandleVideoRewDown` | Known | Event handler |
| 0x0041443D | `HandleVideoRewUp` | Known | Event handler |
| 0x00414469 | `HandleVideoStop` | Known | Event handler |
| 0x0071A1B2 | `HandleAddressBook` | Known | Event handler |
| 0x0071A74E | `HandleSelect` | Known | Event handler |
| 0x0071A789 | `HandleHilite` | Known | Event handler |
| 0x0071A80A | `HandleSelectRegion` | Known | Event handler |
| 0x0071A8AA | `HandleSelectRegion` | Known | Event handler |
| 0x0071A946 | `HandleSelectRegion` | Known | Event handler |
| 0x0071A9EA | `HandleSelectRegion` | Known | Event handler |
| 0x0071AA90 | `HandleSelectRegion` | Known | Event handler |
| 0x0071AB30 | `HandleSelectRegion` | Known | Event handler |
| 0x0071ABDC | `HandleSelectRegion` | Known | Event handler |
| 0x0071AC7E | `HandleSelectRegion` | Known | Event handler |
| 0x0071AD2E | `HandleSelectCity` | Known | Event handler |
| 0x0071AD9A | `HandleHighlightCity` | Known | Event handler |
| 0x0071ADD3 | `HandleSelectCity` | Known | Event handler |
| 0x0071AE3F | `HandleHighlightCity` | Known | Event handler |
| 0x0071AE78 | `HandleSelectCity` | Known | Event handler |
| 0x0071AEE4 | `HandleHighlightCity` | Known | Event handler |
| 0x0071AF1D | `HandleSelectCity` | Known | Event handler |
| 0x0071AF89 | `HandleHighlightCity` | Known | Event handler |
| 0x0071AFC2 | `HandleSelectCity` | Known | Event handler |
| 0x0071B02E | `HandleHighlightCity` | Known | Event handler |
| 0x0071B067 | `HandleSelectCity` | Known | Event handler |
| 0x0071B0D3 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B10C | `HandleSelectCity` | Known | Event handler |
| 0x0071B178 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B1B1 | `HandleSelectCity` | Known | Event handler |
| 0x0071B21D | `HandleHighlightCity` | Known | Event handler |
| 0x0071B256 | `HandleSelectCity` | Known | Event handler |
| 0x0071B2C2 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B2FB | `HandleSelectCity` | Known | Event handler |
| 0x0071B367 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B3A0 | `HandleSelectCity` | Known | Event handler |
| 0x0071B40C | `HandleHighlightCity` | Known | Event handler |
| 0x0071B445 | `HandleSelectCity` | Known | Event handler |
| 0x0071B4B1 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B4EA | `HandleSelectCity` | Known | Event handler |
| 0x0071B556 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B58F | `HandleSelectCity` | Known | Event handler |
| 0x0071B5FB | `HandleHighlightCity` | Known | Event handler |
| 0x0071B634 | `HandleSelectCity` | Known | Event handler |
| 0x0071B6A0 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B6D9 | `HandleSelectCity` | Known | Event handler |
| 0x0071B745 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B77E | `HandleSelectCity` | Known | Event handler |
| 0x0071B7EA | `HandleHighlightCity` | Known | Event handler |
| 0x0071B823 | `HandleSelectCity` | Known | Event handler |
| 0x0071B88F | `HandleHighlightCity` | Known | Event handler |
| 0x0071B8C8 | `HandleSelectCity` | Known | Event handler |
| 0x0071B934 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B96D | `HandleSelectCity` | Known | Event handler |
| 0x0071B9D9 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BA12 | `HandleSelectCity` | Known | Event handler |
| 0x0071BA7E | `HandleHighlightCity` | Known | Event handler |
| 0x0071BAB7 | `HandleSelectCity` | Known | Event handler |
| 0x0071BB23 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BB5C | `HandleSelectCity` | Known | Event handler |
| 0x0071BBC8 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BC01 | `HandleSelectCity` | Known | Event handler |
| 0x0071BC6D | `HandleHighlightCity` | Known | Event handler |
| 0x0071BCA6 | `HandleSelectCity` | Known | Event handler |
| 0x0071BD12 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BD4B | `HandleSelectCity` | Known | Event handler |
| 0x0071BDB7 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BDF0 | `HandleSelectCity` | Known | Event handler |
| 0x0071BE5C | `HandleHighlightCity` | Known | Event handler |
| 0x0071BE95 | `HandleSelectCity` | Known | Event handler |
| 0x0071BF01 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BF3A | `HandleSelectCity` | Known | Event handler |
| 0x0071BFA6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BFDF | `HandleSelectCity` | Known | Event handler |
| 0x0071C04B | `HandleHighlightCity` | Known | Event handler |
| 0x0071C084 | `HandleSelectCity` | Known | Event handler |
| 0x0071C0F0 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C12E | `HandleSelectCity` | Known | Event handler |
| 0x0071C19A | `HandleHighlightCity` | Known | Event handler |
| 0x0071C1D3 | `HandleSelectCity` | Known | Event handler |
| 0x0071C23F | `HandleHighlightCity` | Known | Event handler |
| 0x0071C278 | `HandleSelectCity` | Known | Event handler |
| 0x0071C2E4 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C31D | `HandleSelectCity` | Known | Event handler |
| 0x0071C389 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C3C2 | `HandleSelectCity` | Known | Event handler |
| 0x0071C42E | `HandleHighlightCity` | Known | Event handler |
| 0x0071C467 | `HandleSelectCity` | Known | Event handler |
| 0x0071C4D3 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C50C | `HandleSelectCity` | Known | Event handler |
| 0x0071C578 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C5B1 | `HandleSelectCity` | Known | Event handler |
| 0x0071C61D | `HandleHighlightCity` | Known | Event handler |
| 0x0071C656 | `HandleSelectCity` | Known | Event handler |
| 0x0071C6C2 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C6FB | `HandleSelectCity` | Known | Event handler |
| 0x0071C767 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C7A0 | `HandleSelectCity` | Known | Event handler |
| 0x0071C80C | `HandleHighlightCity` | Known | Event handler |
| 0x0071C845 | `HandleSelectCity` | Known | Event handler |
| 0x0071C8B1 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C8EA | `HandleSelectCity` | Known | Event handler |
| 0x0071C956 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C98F | `HandleSelectCity` | Known | Event handler |
| 0x0071C9FB | `HandleHighlightCity` | Known | Event handler |
| 0x0071CA34 | `HandleSelectCity` | Known | Event handler |
| 0x0071CAA0 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CAD9 | `HandleSelectCity` | Known | Event handler |
| 0x0071CB45 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CB7E | `HandleSelectCity` | Known | Event handler |
| 0x0071CBEA | `HandleHighlightCity` | Known | Event handler |
| 0x0071CC23 | `HandleSelectCity` | Known | Event handler |
| 0x0071CC8F | `HandleHighlightCity` | Known | Event handler |
| 0x0071CCC8 | `HandleSelectCity` | Known | Event handler |
| 0x0071CD34 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CD6D | `HandleSelectCity` | Known | Event handler |
| 0x0071CDD9 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CE12 | `HandleSelectCity` | Known | Event handler |
| 0x0071CE7E | `HandleHighlightCity` | Known | Event handler |
| 0x0071CEB7 | `HandleSelectCity` | Known | Event handler |
| 0x0071CF23 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CF5C | `HandleSelectCity` | Known | Event handler |
| 0x0071CFC8 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D001 | `HandleSelectCity` | Known | Event handler |
| 0x0071D06D | `HandleHighlightCity` | Known | Event handler |
| 0x0071D0A6 | `HandleSelectCity` | Known | Event handler |
| 0x0071D112 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D14B | `HandleSelectCity` | Known | Event handler |
| 0x0071D1B7 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D1F0 | `HandleSelectCity` | Known | Event handler |
| 0x0071D25C | `HandleHighlightCity` | Known | Event handler |
| 0x0071D295 | `HandleSelectCity` | Known | Event handler |
| 0x0071D301 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D33A | `HandleSelectCity` | Known | Event handler |
| 0x0071D3A6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D3DF | `HandleSelectCity` | Known | Event handler |
| 0x0071D44B | `HandleHighlightCity` | Known | Event handler |
| 0x0071D484 | `HandleSelectCity` | Known | Event handler |
| 0x0071D4F0 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D529 | `HandleSelectCity` | Known | Event handler |
| 0x0071D595 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D5CE | `HandleSelectCity` | Known | Event handler |
| 0x0071D63A | `HandleHighlightCity` | Known | Event handler |
| 0x0071D673 | `HandleSelectCity` | Known | Event handler |
| 0x0071D6DF | `HandleHighlightCity` | Known | Event handler |
| 0x0071D718 | `HandleSelectCity` | Known | Event handler |
| 0x0071D784 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D7BD | `HandleSelectCity` | Known | Event handler |
| 0x0071D829 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D862 | `HandleSelectCity` | Known | Event handler |
| 0x0071D8CE | `HandleHighlightCity` | Known | Event handler |
| 0x0071D907 | `HandleSelectCity` | Known | Event handler |
| 0x0071D973 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D9AC | `HandleSelectCity` | Known | Event handler |
| 0x0071DA18 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DA51 | `HandleSelectCity` | Known | Event handler |
| 0x0071DABD | `HandleHighlightCity` | Known | Event handler |
| 0x0071DAF6 | `HandleSelectCity` | Known | Event handler |
| 0x0071DB62 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DB9B | `HandleSelectCity` | Known | Event handler |
| 0x0071DC07 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DC40 | `HandleSelectCity` | Known | Event handler |
| 0x0071DCAC | `HandleHighlightCity` | Known | Event handler |
| 0x0071DCE5 | `HandleSelectCity` | Known | Event handler |
| 0x0071DD51 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DD8A | `HandleSelectCity` | Known | Event handler |
| 0x0071DDF6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DE2F | `HandleSelectCity` | Known | Event handler |
| 0x0071DE9B | `HandleHighlightCity` | Known | Event handler |
| 0x0071DED4 | `HandleSelectCity` | Known | Event handler |
| 0x0071DF40 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DF79 | `HandleSelectCity` | Known | Event handler |
| 0x0071DFE5 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E01E | `HandleSelectCity` | Known | Event handler |
| 0x0071E08A | `HandleHighlightCity` | Known | Event handler |
| 0x0071E0C3 | `HandleSelectCity` | Known | Event handler |
| 0x0071E12F | `HandleHighlightCity` | Known | Event handler |
| 0x0071E168 | `HandleSelectCity` | Known | Event handler |
| 0x0071E1D4 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E20D | `HandleSelectCity` | Known | Event handler |
| 0x0071E279 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E2B2 | `HandleSelectCity` | Known | Event handler |
| 0x0071E31E | `HandleHighlightCity` | Known | Event handler |
| 0x0071E357 | `HandleSelectCity` | Known | Event handler |
| 0x0071E3C3 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E3FC | `HandleSelectCity` | Known | Event handler |
| 0x0071E468 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E4A1 | `HandleSelectCity` | Known | Event handler |
| 0x0071E50D | `HandleHighlightCity` | Known | Event handler |
| 0x0071E546 | `HandleSelectCity` | Known | Event handler |
| 0x0071E5B2 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E5F2 | `HandleSelectCity` | Known | Event handler |
| 0x0071E65E | `HandleHighlightCity` | Known | Event handler |
| 0x0071E697 | `HandleSelectCity` | Known | Event handler |
| 0x0071E703 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E73C | `HandleSelectCity` | Known | Event handler |
| 0x0071E7A8 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E7E6 | `HandleSelectCity` | Known | Event handler |
| 0x0071E852 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E88B | `HandleSelectCity` | Known | Event handler |
| 0x0071E8F7 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E930 | `HandleSelectCity` | Known | Event handler |
| 0x0071E99C | `HandleHighlightCity` | Known | Event handler |
| 0x0071E9D5 | `HandleSelectCity` | Known | Event handler |
| 0x0071EA41 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EA7A | `HandleSelectCity` | Known | Event handler |
| 0x0071EAE6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EB1F | `HandleSelectCity` | Known | Event handler |
| 0x0071EB8B | `HandleHighlightCity` | Known | Event handler |
| 0x0071EBC4 | `HandleSelectCity` | Known | Event handler |
| 0x0071EC30 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EC69 | `HandleSelectCity` | Known | Event handler |
| 0x0071ECD5 | `HandleHighlightCity` | Known | Event handler |
| 0x0071ED12 | `HandleSelectCity` | Known | Event handler |
| 0x0071ED7E | `HandleHighlightCity` | Known | Event handler |
| 0x0071EDB7 | `HandleSelectCity` | Known | Event handler |
| 0x0071EE23 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EE5C | `HandleSelectCity` | Known | Event handler |
| 0x0071EEC8 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EF01 | `HandleSelectCity` | Known | Event handler |
| 0x0071EF6D | `HandleHighlightCity` | Known | Event handler |
| 0x0071EFA6 | `HandleSelectCity` | Known | Event handler |
| 0x0071F012 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F04B | `HandleSelectCity` | Known | Event handler |
| 0x0071F0B7 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F0F0 | `HandleSelectCity` | Known | Event handler |
| 0x0071F15C | `HandleHighlightCity` | Known | Event handler |
| 0x0071F195 | `HandleSelectCity` | Known | Event handler |
| 0x0071F201 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F23A | `HandleSelectCity` | Known | Event handler |
| 0x0071F2A6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F2DF | `HandleSelectCity` | Known | Event handler |
| 0x0071F34B | `HandleHighlightCity` | Known | Event handler |
| 0x0071F384 | `HandleSelectCity` | Known | Event handler |
| 0x0071F3F0 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F429 | `HandleSelectCity` | Known | Event handler |
| 0x0071F495 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F4CE | `HandleSelectCity` | Known | Event handler |
| 0x0071F53A | `HandleHighlightCity` | Known | Event handler |
| 0x0071F573 | `HandleSelectCity` | Known | Event handler |
| 0x0071F5DF | `HandleHighlightCity` | Known | Event handler |
| 0x0071F618 | `HandleSelectCity` | Known | Event handler |
| 0x0071F684 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F6BD | `HandleSelectCity` | Known | Event handler |
| 0x0071F729 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F762 | `HandleSelectCity` | Known | Event handler |
| 0x0071F7CE | `HandleHighlightCity` | Known | Event handler |
| 0x0071F807 | `HandleSelectCity` | Known | Event handler |
| 0x0071F873 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F8AC | `HandleSelectCity` | Known | Event handler |
| 0x0071F918 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F951 | `HandleSelectCity` | Known | Event handler |
| 0x0071F9BD | `HandleHighlightCity` | Known | Event handler |
| 0x0071F9F6 | `HandleSelectCity` | Known | Event handler |
| 0x0071FA62 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FA9B | `HandleSelectCity` | Known | Event handler |
| 0x0071FB07 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FB40 | `HandleSelectCity` | Known | Event handler |
| 0x0071FBAC | `HandleHighlightCity` | Known | Event handler |
| 0x0071FBE5 | `HandleSelectCity` | Known | Event handler |
| 0x0071FC51 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FC8A | `HandleSelectCity` | Known | Event handler |
| 0x0071FCF6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FD2F | `HandleSelectCity` | Known | Event handler |
| 0x0071FD9B | `HandleHighlightCity` | Known | Event handler |
| 0x0071FDD4 | `HandleSelectCity` | Known | Event handler |
| 0x0071FE40 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FE79 | `HandleSelectCity` | Known | Event handler |
| 0x0071FEE5 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FF1E | `HandleSelectCity` | Known | Event handler |
| 0x0071FF8A | `HandleHighlightCity` | Known | Event handler |
| 0x0071FFC3 | `HandleSelectCity` | Known | Event handler |
| 0x0072002F | `HandleHighlightCity` | Known | Event handler |
| 0x00720068 | `HandleSelectCity` | Known | Event handler |
| 0x007200D4 | `HandleHighlightCity` | Known | Event handler |
| 0x0072010D | `HandleSelectCity` | Known | Event handler |
| 0x00720179 | `HandleHighlightCity` | Known | Event handler |
| 0x007201B2 | `HandleSelectCity` | Known | Event handler |
| 0x0072021E | `HandleHighlightCity` | Known | Event handler |
| 0x00720257 | `HandleSelectCity` | Known | Event handler |
| 0x007202C3 | `HandleHighlightCity` | Known | Event handler |
| 0x00720302 | `HandleSelectCity` | Known | Event handler |
| 0x0072036E | `HandleHighlightCity` | Known | Event handler |
| 0x007203A7 | `HandleSelectCity` | Known | Event handler |
| 0x00720413 | `HandleHighlightCity` | Known | Event handler |
| 0x0072044C | `HandleSelectCity` | Known | Event handler |
| 0x007204B8 | `HandleHighlightCity` | Known | Event handler |
| 0x007204F1 | `HandleSelectCity` | Known | Event handler |
| 0x0072055D | `HandleHighlightCity` | Known | Event handler |
| 0x00720596 | `HandleSelectCity` | Known | Event handler |
| 0x00720602 | `HandleHighlightCity` | Known | Event handler |
| 0x0072063B | `HandleSelectCity` | Known | Event handler |
| 0x007206A7 | `HandleHighlightCity` | Known | Event handler |
| 0x007206E0 | `HandleSelectCity` | Known | Event handler |
| 0x0072074C | `HandleHighlightCity` | Known | Event handler |
| 0x00720785 | `HandleSelectCity` | Known | Event handler |
| 0x007207F1 | `HandleHighlightCity` | Known | Event handler |
| 0x0072082A | `HandleSelectCity` | Known | Event handler |
| 0x00720896 | `HandleHighlightCity` | Known | Event handler |
| 0x007208CF | `HandleSelectCity` | Known | Event handler |
| 0x0072093B | `HandleHighlightCity` | Known | Event handler |
| 0x00720974 | `HandleSelectCity` | Known | Event handler |
| 0x007209E0 | `HandleHighlightCity` | Known | Event handler |
| 0x00720A19 | `HandleSelectCity` | Known | Event handler |
| 0x00720A85 | `HandleHighlightCity` | Known | Event handler |
| 0x00720ABE | `HandleSelectCity` | Known | Event handler |
| 0x00720B2A | `HandleHighlightCity` | Known | Event handler |
| 0x00720B63 | `HandleSelectCity` | Known | Event handler |
| 0x00720BCF | `HandleHighlightCity` | Known | Event handler |
| 0x00720C08 | `HandleSelectCity` | Known | Event handler |
| 0x00720C74 | `HandleHighlightCity` | Known | Event handler |
| 0x00720CAD | `HandleSelectCity` | Known | Event handler |
| 0x00720D19 | `HandleHighlightCity` | Known | Event handler |
| 0x00720D52 | `HandleSelectCity` | Known | Event handler |
| 0x00720DBE | `HandleHighlightCity` | Known | Event handler |
| 0x00720DF7 | `HandleSelectCity` | Known | Event handler |
| 0x00720E63 | `HandleHighlightCity` | Known | Event handler |
| 0x00720E9C | `HandleSelectCity` | Known | Event handler |
| 0x00720F08 | `HandleHighlightCity` | Known | Event handler |
| 0x00720F41 | `HandleSelectCity` | Known | Event handler |
| 0x00720FAD | `HandleHighlightCity` | Known | Event handler |
| 0x00720FE6 | `HandleSelectCity` | Known | Event handler |
| 0x00721052 | `HandleHighlightCity` | Known | Event handler |
| 0x0072108B | `HandleSelectCity` | Known | Event handler |
| 0x007210F7 | `HandleHighlightCity` | Known | Event handler |
| 0x00721130 | `HandleSelectCity` | Known | Event handler |
| 0x0072119C | `HandleHighlightCity` | Known | Event handler |
| 0x007211D5 | `HandleSelectCity` | Known | Event handler |
| 0x00721241 | `HandleHighlightCity` | Known | Event handler |
| 0x0072127A | `HandleSelectCity` | Known | Event handler |
| 0x007212E6 | `HandleHighlightCity` | Known | Event handler |
| 0x0072131F | `HandleSelectCity` | Known | Event handler |
| 0x0072138B | `HandleHighlightCity` | Known | Event handler |
| 0x007213C4 | `HandleSelectCity` | Known | Event handler |
| 0x00721430 | `HandleHighlightCity` | Known | Event handler |
| 0x00721469 | `HandleSelectCity` | Known | Event handler |
| 0x007214D5 | `HandleHighlightCity` | Known | Event handler |
| 0x0072150E | `HandleSelectCity` | Known | Event handler |
| 0x0072157A | `HandleHighlightCity` | Known | Event handler |
| 0x007215B3 | `HandleSelectCity` | Known | Event handler |
| 0x0072161F | `HandleHighlightCity` | Known | Event handler |
| 0x00721658 | `HandleSelectCity` | Known | Event handler |
| 0x007216C4 | `HandleHighlightCity` | Known | Event handler |
| 0x007216FD | `HandleSelectCity` | Known | Event handler |
| 0x00721769 | `HandleHighlightCity` | Known | Event handler |
| 0x007217A2 | `HandleSelectCity` | Known | Event handler |
| 0x0072180E | `HandleHighlightCity` | Known | Event handler |
| 0x00721847 | `HandleSelectCity` | Known | Event handler |
| 0x007218B3 | `HandleHighlightCity` | Known | Event handler |
| 0x007218EC | `HandleSelectCity` | Known | Event handler |
| 0x00721958 | `HandleHighlightCity` | Known | Event handler |
| 0x00721991 | `HandleSelectCity` | Known | Event handler |
| 0x007219FD | `HandleHighlightCity` | Known | Event handler |
| 0x00721A36 | `HandleSelectCity` | Known | Event handler |
| 0x00721AA2 | `HandleHighlightCity` | Known | Event handler |
| 0x00721ADB | `HandleSelectCity` | Known | Event handler |
| 0x00721B47 | `HandleHighlightCity` | Known | Event handler |
| 0x00721B80 | `HandleSelectCity` | Known | Event handler |
| 0x00721BEC | `HandleHighlightCity` | Known | Event handler |
| 0x00721C25 | `HandleSelectCity` | Known | Event handler |
| 0x00721C91 | `HandleHighlightCity` | Known | Event handler |
| 0x00721CCA | `HandleSelectCity` | Known | Event handler |
| 0x00721D36 | `HandleHighlightCity` | Known | Event handler |
| 0x00721D6F | `HandleSelectCity` | Known | Event handler |
| 0x00721DDB | `HandleHighlightCity` | Known | Event handler |
| 0x00721E14 | `HandleSelectCity` | Known | Event handler |
| 0x00721E80 | `HandleHighlightCity` | Known | Event handler |
| 0x00721EB9 | `HandleSelectCity` | Known | Event handler |
| 0x00721F25 | `HandleHighlightCity` | Known | Event handler |
| 0x00721F5E | `HandleSelectCity` | Known | Event handler |
| 0x00721FCA | `HandleHighlightCity` | Known | Event handler |
| 0x00722003 | `HandleSelectCity` | Known | Event handler |
| 0x0072206F | `HandleHighlightCity` | Known | Event handler |
| 0x007220A8 | `HandleSelectCity` | Known | Event handler |
| 0x00722114 | `HandleHighlightCity` | Known | Event handler |
| 0x0072214D | `HandleSelectCity` | Known | Event handler |
| 0x007221B9 | `HandleHighlightCity` | Known | Event handler |
| 0x007221F2 | `HandleSelectCity` | Known | Event handler |
| 0x0072225E | `HandleHighlightCity` | Known | Event handler |
| 0x00722297 | `HandleSelectCity` | Known | Event handler |
| 0x00722303 | `HandleHighlightCity` | Known | Event handler |
| 0x00722342 | `HandleSelectCity` | Known | Event handler |
| 0x007223AE | `HandleHighlightCity` | Known | Event handler |
| 0x007223E7 | `HandleSelectCity` | Known | Event handler |
| 0x00722453 | `HandleHighlightCity` | Known | Event handler |
| 0x0072248C | `HandleSelectCity` | Known | Event handler |
| 0x007224F8 | `HandleHighlightCity` | Known | Event handler |
| 0x00722531 | `HandleSelectCity` | Known | Event handler |
| 0x0072259D | `HandleHighlightCity` | Known | Event handler |
| 0x007225D6 | `HandleSelectCity` | Known | Event handler |
| 0x00722642 | `HandleHighlightCity` | Known | Event handler |
| 0x00722682 | `HandleSelectCity` | Known | Event handler |
| 0x007226EE | `HandleHighlightCity` | Known | Event handler |
| 0x00722727 | `HandleSelectCity` | Known | Event handler |
| 0x00722793 | `HandleHighlightCity` | Known | Event handler |
| 0x007227CC | `HandleSelectCity` | Known | Event handler |
| 0x00722838 | `HandleHighlightCity` | Known | Event handler |
| 0x00722871 | `HandleSelectCity` | Known | Event handler |
| 0x007228DD | `HandleHighlightCity` | Known | Event handler |
| 0x00722916 | `HandleSelectCity` | Known | Event handler |
| 0x00722982 | `HandleHighlightCity` | Known | Event handler |
| 0x007229BB | `HandleSelectCity` | Known | Event handler |
| 0x00722A27 | `HandleHighlightCity` | Known | Event handler |
| 0x00722A60 | `HandleSelectCity` | Known | Event handler |
| 0x00722ACC | `HandleHighlightCity` | Known | Event handler |
| 0x00722B05 | `HandleSelectCity` | Known | Event handler |
| 0x00722B71 | `HandleHighlightCity` | Known | Event handler |
| 0x00722BAA | `HandleSelectCity` | Known | Event handler |
| 0x00722C16 | `HandleHighlightCity` | Known | Event handler |
| 0x00722C4F | `HandleSelectCity` | Known | Event handler |
| 0x00722CBB | `HandleHighlightCity` | Known | Event handler |
| 0x00722CF4 | `HandleSelectCity` | Known | Event handler |
| 0x00722D60 | `HandleHighlightCity` | Known | Event handler |
| 0x00722D99 | `HandleSelectCity` | Known | Event handler |
| 0x00722E05 | `HandleHighlightCity` | Known | Event handler |
| 0x00722E3E | `HandleSelectCity` | Known | Event handler |
| 0x00722EAA | `HandleHighlightCity` | Known | Event handler |
| 0x00722EE3 | `HandleSelectCity` | Known | Event handler |
| 0x00722F4F | `HandleHighlightCity` | Known | Event handler |
| 0x00722F88 | `HandleSelectCity` | Known | Event handler |
| 0x00722FF4 | `HandleHighlightCity` | Known | Event handler |
| 0x0072302D | `HandleSelectCity` | Known | Event handler |
| 0x00723099 | `HandleHighlightCity` | Known | Event handler |
| 0x007230D2 | `HandleSelectCity` | Known | Event handler |
| 0x0072313E | `HandleHighlightCity` | Known | Event handler |
| 0x00723636 | `HandleMusicSelected` | Known | Event handler |
| 0x00723678 | `HandleMusicHilited` | Known | Event handler |
| 0x007236B0 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x007236F6 | `HandleMusicHilited` | Known | Event handler |
| 0x0072372E | `HandleGotoGenius` | Known | Event handler |
| 0x0072376D | `HandleGeniusHilited` | Known | Event handler |
| 0x007237A6 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x007237EC | `HandlePlaylistsHilited` | Known | Event handler |
| 0x00723828 | `HandleArtistsSelected` | Known | Event handler |
| 0x0072386C | `HandleArtistsHilited` | Known | Event handler |
| 0x007238A6 | `HandleAlbumsSelected` | Known | Event handler |
| 0x007238E9 | `HandleAlbumsHilited` | Known | Event handler |
| 0x00723922 | `HandleCompilationsSelected` | Known | Event handler |
| 0x0072396B | `HandleCompilationsHilited` | Known | Event handler |
| 0x007239AA | `HandleSongsSelected` | Known | Event handler |
| 0x007239EC | `HandleSongsHilited` | Known | Event handler |
| 0x00723A24 | `HandleGenresSelected` | Known | Event handler |
| 0x00723A67 | `HandleGenresHilited` | Known | Event handler |
| 0x00723AA0 | `HandleComposersSelected` | Known | Event handler |
| 0x00723AE6 | `HandleComposersHilited` | Known | Event handler |
| 0x00723B22 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x00723B69 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00723C28 | `HandleMusicHilited` | Known | Event handler |
| 0x00723C60 | `HandleVideosSelected` | Known | Event handler |
| 0x00723CA3 | `HandleVideosHilited` | Known | Event handler |
| 0x00723CDC | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00723D27 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00723D68 | `HandleMoviesSelected` | Known | Event handler |
| 0x00723DAB | `HandleMoviesHilited` | Known | Event handler |
| 0x00723DE4 | `HandleTVShowsSelected` | Known | Event handler |
| 0x00723E28 | `HandleTVShowsHilited` | Known | Event handler |
| 0x00723E62 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00723EAA | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00723EE8 | `HandleRentalsSelected` | Known | Event handler |
| 0x00723F2C | `HandleRentalsHilited` | Known | Event handler |
| 0x00723F66 | `HandlePhotosSelected` | Known | Event handler |
| 0x00723FA9 | `HandlePhotosHilited` | Known | Event handler |
| 0x00723FE2 | `HandlePhotosSelected` | Known | Event handler |
| 0x00724025 | `HandlePhotosHilited` | Known | Event handler |
| 0x0072405E | `HandlePodcastsSelected` | Known | Event handler |
| 0x007240A3 | `HandlePodcastsHilited` | Known | Event handler |
| 0x00724156 | `HandleGenericHilited` | Known | Event handler |
| 0x0072424F | `HandleGenericHilited` | Known | Event handler |
| 0x00724734 | `HandleLock` | Known | Event handler |
| 0x007248A5 | `HandleNikePlusSelected` | Known | Event handler |
| 0x007248EA | `HandleGenericHilited` | Known | Event handler |
| 0x007249F0 | `HandleGenericHilited` | Known | Event handler |
| 0x00724AEF | `HandleGenericHilited` | Known | Event handler |
| 0x00724BDC | `HandleGenericHilited` | Known | Event handler |
| 0x00724CD9 | `HandleGenericHilited` | Known | Event handler |
| 0x00724D53 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00724D9C | `HandleGenericHilited` | Known | Event handler |
| 0x00724E15 | `HandleBacklightSelected` | Known | Event handler |
| 0x00724E5B | `HandleGenericHilited` | Known | Event handler |
| 0x00724ED6 | `HandleSleepSelected` | Known | Event handler |
| 0x00724F18 | `HandleGenericHilited` | Known | Event handler |
| 0x00724F8F | `HandleNowPlaying` | Known | Event handler |
| 0x00725007 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x0072504A | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00725090 | `HandleMusicHilited` | Known | Event handler |
| 0x007250C8 | `HandleGotoGenius` | Known | Event handler |
| 0x007250FE | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00725144 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x00725182 | `HandleArtistsSelected` | Known | Event handler |
| 0x007251C6 | `HandleArtistsHilited` | Known | Event handler |
| 0x00725200 | `HandleAlbumsSelected` | Known | Event handler |
| 0x00725243 | `HandleAlbumsHilited` | Known | Event handler |
| 0x0072527C | `HandleCompilationsSelected` | Known | Event handler |
| 0x007252C5 | `HandleCompilationsHilited` | Known | Event handler |
| 0x00725304 | `HandleSongsSelected` | Known | Event handler |
| 0x00725346 | `HandleSongsHilited` | Known | Event handler |
| 0x007253F1 | `HandleGenericHilited` | Known | Event handler |
| 0x00725469 | `HandleGenresSelected` | Known | Event handler |
| 0x007254AC | `HandleGenresHilited` | Known | Event handler |
| 0x007254E5 | `HandleComposersSelected` | Known | Event handler |
| 0x0072552B | `HandleComposersHilited` | Known | Event handler |
| 0x00725567 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x007255AE | `HandleAudiobooksHilited` | Known | Event handler |
| 0x0072566D | `HandleMusicHilited` | Known | Event handler |
| 0x007256E1 | `HandlePlayPause` | Known | Event handler |
| 0x00725716 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x00725800 | `HandleSelect` | Known | Event handler |
| 0x00725846 | `HandleMoviesSelected` | Known | Event handler |
| 0x00725889 | `HandleMoviesHilited` | Known | Event handler |
| 0x007258C2 | `HandleRentalsSelected` | Known | Event handler |
| 0x00725906 | `HandleRentalsHilited` | Known | Event handler |
| 0x00725940 | `HandleTVShowsSelected` | Known | Event handler |
| 0x00725984 | `HandleTVShowsHilited` | Known | Event handler |
| 0x007259BE | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00725A06 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00725A44 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00725A8F | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00725B55 | `HandleVideosHilited` | Known | Event handler |
| 0x007261A3 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x00726D2A | `HandleMainMenu` | Known | Event handler |
| 0x00726D62 | `HandleMusicMenu` | Known | Event handler |
| 0x0072728A | `HandleRadioRegion` | Known | Event handler |
| 0x0072732E | `HandleLanguage` | Known | Event handler |
| 0x00727434 | `HandleNew` | Known | Event handler |
| 0x007274AF | `HandleClear` | Known | Event handler |
| 0x007274E0 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x0072759C | `HandleSelectIndexedSession` | Known | Event handler |
| 0x00727705 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x00727758 | `HandleSelect` | Known | Event handler |
| 0x00727882 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x007278BC | `HandleEQSettingSelected` | Known | Event handler |
| 0x007278F4 | `HandleEQSettingSelected` | Known | Event handler |
| 0x0073A836 | `HandleMenuSelection` | Known | Event handler |
| 0x0073AB7B | `HandleLoadingCancelled` | Known | Event handler |
| 0x0073AC17 | `HandleLoadingCancelled` | Known | Event handler |
| 0x0073ACE4 | `HandleItemSelected` | Known | Event handler |
| 0x0073AE2F | `HandleNextContact` | Known | Event handler |
| 0x0073AE5B | `HandlePreviousContact` | Known | Event handler |
| 0x0073AE91 | `HandleSelectKey` | Known | Event handler |
| 0x0073B4A2 | `HandleSelect` | Known | Event handler |
| 0x0073B7C9 | `HandleDateChosen` | Known | Event handler |
| 0x0073B7FF | `HandleTimeChosen` | Known | Event handler |
| 0x0073B835 | `HandleFrequencyChosen` | Known | Event handler |
| 0x0073B870 | `HandleSoundChosen` | Known | Event handler |
| 0x0073B8A7 | `HandleLabelChosen` | Known | Event handler |
| 0x0073B8DE | `HandleDeleteChosen` | Known | Event handler |
| 0x0073B91A | `HandleSelect` | Known | Event handler |
| 0x0073B952 | `HandleSelect` | Known | Event handler |
| 0x0073BC93 | `HandleLeaveAlarm` | Known | Event handler |
| 0x0073BCC0 | `HandleLeaveAlarm` | Known | Event handler |
| 0x0073BCEF | `HandleLeaveAlarm` | Known | Event handler |
| 0x0073BD1C | `HandleLeaveAlarm` | Known | Event handler |
| 0x0073BE56 | `HandleSelect` | Known | Event handler |
| 0x0073BE84 | `HandleSelect` | Known | Event handler |
| 0x0073BFE3 | `HandleNextDay` | Known | Event handler |
| 0x0073C00B | `HandlePreviousDay` | Known | Event handler |
| 0x0073C1BA | `HandleSelect` | Known | Event handler |
| 0x0073C1E7 | `HandleNextDay` | Known | Event handler |
| 0x0073C20F | `HandlePreviousDay` | Known | Event handler |
| 0x0073C3B7 | `HandleNextDay` | Known | Event handler |
| 0x0073C3DF | `HandlePreviousDay` | Known | Event handler |
| 0x0073C4A0 | `HandleSelect` | Known | Event handler |
| 0x0073C4CB | `HandleNextDay` | Known | Event handler |
| 0x0073C4F3 | `HandlePreviousDay` | Known | Event handler |
| 0x0073C66A | `HandleSelectLozinch` | Known | Event handler |
| 0x0073C7E2 | `HandleSelectLozinch` | Known | Event handler |
| 0x0073C901 | `HandleFlowNext` | Known | Event handler |
| 0x0073C92F | `HandlePlayPause` | Known | Event handler |
| 0x0073C97E | `HandleFlowPrev` | Known | Event handler |
| 0x0073C9A9 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x0073CA9D | `HandleAlbumSelected` | Known | Event handler |
| 0x0073CC38 | `HandleFlowNext` | Known | Event handler |
| 0x0073CC86 | `HandleFlowNext` | Known | Event handler |
| 0x0073CCB4 | `HandlePlayPause` | Known | Event handler |
| 0x0073CD03 | `HandleFlowPrev` | Known | Event handler |
| 0x0073CD2F | `HandleFlowPrev` | Known | Event handler |
| 0x0073CD4F | `HandleFlowWheel` | Known | Event handler |
| 0x0073D0DF | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0073D50A | `HandleArrowDown` | Known | Event handler |
| 0x0073D574 | `HandleArrowUp` | Known | Event handler |
| 0x0073D593 | `HandleWheel` | Known | Event handler |
| 0x0073D61C | `HandleSelect` | Known | Event handler |
| 0x0073D699 | `HandleGameHilited` | Known | Event handler |
| 0x00740AFF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074289B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00744637 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007463D3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074816F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00749F0B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074BCA7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074DA43 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074F7DF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075157B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00753317 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007550B3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00756E4F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00758BEB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075A987 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075C723 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075E4BF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076025B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00761FF7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00763D93 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00765B2F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007678CB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00769667 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076B403 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076D19F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076EF3B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00770CD7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00772A73 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077480F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007765AB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00778347 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077A0E3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077BE7F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077DC1B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077F9B7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00781753 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007834EF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00785270 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00785EEC | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00786B68 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007877E4 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00788460 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007890DC | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00789D58 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078A9D4 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078B650 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078C2CC | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078CF48 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078DBC4 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078E840 | `HandlePlayPause` | Known | Event handler |
| 0x0078E876 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078E8B8 | `HandleAddToOTG` | Known | Event handler |
| 0x0078EA55 | `HandlePlayPause` | Known | Event handler |
| 0x0078EA7C | `HandleSelect` | Known | Event handler |
| 0x0078EAA9 | `HandleHilite` | Known | Event handler |
| 0x0078EADC | `HandlePlayPause` | Known | Event handler |
| 0x0078EB6F | `HandlePlayPause` | Known | Event handler |
| 0x0078EB96 | `HandleSelect` | Known | Event handler |
| 0x0078EBFC | `HandleHilite` | Known | Event handler |
| 0x0078EC2E | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x0078EC78 | `HandlePlayPause` | Known | Event handler |
| 0x0078ECAE | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078ECF5 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0078ED38 | `HandleAddToOTG` | Known | Event handler |
| 0x0078ED9B | `HandleStartGenius` | Known | Event handler |
| 0x0078EDD7 | `HandleViewAlbum` | Known | Event handler |
| 0x0078EE12 | `HandleViewArtist` | Known | Event handler |
| 0x0078EE53 | `HandleViewCompilation` | Known | Event handler |
| 0x0078EFF3 | `HandlePlayPause` | Known | Event handler |
| 0x0078F01A | `HandleSelect` | Known | Event handler |
| 0x0078F084 | `HandlePlayPause` | Known | Event handler |
| 0x0078F0BA | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078F101 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0078F144 | `HandleAddToOTG` | Known | Event handler |
| 0x0078F1A7 | `HandleStartGenius` | Known | Event handler |
| 0x0078F1E3 | `HandleViewAlbum` | Known | Event handler |
| 0x0078F21E | `HandleViewArtist` | Known | Event handler |
| 0x0078F25F | `HandleViewCompilation` | Known | Event handler |
| 0x0078F3FF | `HandlePlayPause` | Known | Event handler |
| 0x0078F426 | `HandleSelect` | Known | Event handler |
| 0x0078F490 | `HandlePlayPause` | Known | Event handler |
| 0x0078F4CE | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0078F511 | `HandleAddToOTG` | Known | Event handler |
| 0x0078F574 | `HandleStartGenius` | Known | Event handler |
| 0x0078F5B0 | `HandleViewAlbum` | Known | Event handler |
| 0x0078F5EB | `HandleViewArtist` | Known | Event handler |
| 0x0078F62C | `HandleViewCompilation` | Known | Event handler |
| 0x0078F7BF | `HandleSelect` | Known | Event handler |
| 0x0078F824 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0078F868 | `HandlePlayPause` | Known | Event handler |
| 0x0078F89E | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078F8E0 | `HandleAddToOTG` | Known | Event handler |
| 0x0078FB3A | `HandlePlayPause` | Known | Event handler |
| 0x0078FB61 | `HandleSelect` | Known | Event handler |
| 0x0078FB8E | `HandleHilite` | Known | Event handler |
| 0x0078FBC0 | `HandlePlayPause` | Known | Event handler |
| 0x0078FBF6 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078FC38 | `HandleAddToOTG` | Known | Event handler |
| 0x0078FE92 | `HandlePlayPause` | Known | Event handler |
| 0x0078FEB9 | `HandleSelect` | Known | Event handler |
| 0x0078FEE6 | `HandleHilite` | Known | Event handler |
| 0x0078FF18 | `HandlePlayPause` | Known | Event handler |
| 0x0078FF4E | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078FF90 | `HandleAddToOTG` | Known | Event handler |
| 0x007902A3 | `HandlePlayPause` | Known | Event handler |
| 0x007902CA | `HandleSelect` | Known | Event handler |
| 0x007902FC | `HandlePlayPause` | Known | Event handler |
| 0x00790332 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00790374 | `HandleAddToOTG` | Known | Event handler |
| 0x0079042E | `HandlePlayPause` | Known | Event handler |
| 0x00790455 | `HandleSelect` | Known | Event handler |
| 0x007904E4 | `HandlePlayPause` | Known | Event handler |
| 0x0079051A | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079055C | `HandleAddToOTG` | Known | Event handler |
| 0x0079073D | `HandlePlayPause` | Known | Event handler |
| 0x00790764 | `HandleSelect` | Known | Event handler |
| 0x00790794 | `HandlePlayPause` | Known | Event handler |
| 0x007907CA | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079080C | `HandleAddToOTG` | Known | Event handler |
| 0x007908B9 | `HandleSelect` | Known | Event handler |
| 0x00790952 | `HandleHilite` | Known | Event handler |
| 0x0079097E | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007909C0 | `HandlePlayPause` | Known | Event handler |
| 0x007909F6 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00790A38 | `HandleAddToOTG` | Known | Event handler |
| 0x00790AE5 | `HandleSelect` | Known | Event handler |
| 0x00790B4A | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00790B8C | `HandlePlayPause` | Known | Event handler |
| 0x00790D30 | `HandleSelect` | Known | Event handler |
| 0x00790D5D | `HandleHilite` | Known | Event handler |
| 0x00790D89 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00790DCC | `HandlePlayPause` | Known | Event handler |
| 0x00790E52 | `HandleSelect` | Known | Event handler |
| 0x00790EE0 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00790F24 | `HandlePlayPause` | Known | Event handler |
| 0x00790FAA | `HandleSelect` | Known | Event handler |
| 0x0079100F | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00791050 | `HandlePlayPause` | Known | Event handler |
| 0x007910D6 | `HandleSelect` | Known | Event handler |
| 0x0079113C | `HandleHilite` | Known | Event handler |
| 0x00791168 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007911AC | `HandlePlayPause` | Known | Event handler |
| 0x007911E2 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00791224 | `HandleAddToOTG` | Known | Event handler |
| 0x007914A9 | `HandlePlayPause` | Known | Event handler |
| 0x007914D0 | `HandleSelect` | Known | Event handler |
| 0x00791500 | `HandlePlayPause` | Known | Event handler |
| 0x00791536 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079157D | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007915C0 | `HandleAddToOTG` | Known | Event handler |
| 0x00791623 | `HandleStartGenius` | Known | Event handler |
| 0x0079165F | `HandleViewAlbum` | Known | Event handler |
| 0x0079169A | `HandleViewArtist` | Known | Event handler |
| 0x007916DB | `HandleViewCompilation` | Known | Event handler |
| 0x00791BC3 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x00791C08 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00791C4B | `HandleAddToOTG` | Known | Event handler |
| 0x00791CAE | `HandleStartGenius` | Known | Event handler |
| 0x00791CEA | `HandleViewAlbum` | Known | Event handler |
| 0x00791D25 | `HandleViewArtist` | Known | Event handler |
| 0x00791D66 | `HandleViewCompilation` | Known | Event handler |
| 0x0079213C | `HandlePlayPause` | Known | Event handler |
| 0x00792269 | `HandleSelect` | Known | Event handler |
| 0x00792295 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007922D8 | `HandlePlayPause` | Known | Event handler |
| 0x0079235E | `HandleSelect` | Known | Event handler |
| 0x0079238B | `HandleHilite` | Known | Event handler |
| 0x007923B7 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007923F8 | `HandlePlayPause` | Known | Event handler |
| 0x0079252B | `HandleSelect` | Known | Event handler |
| 0x00792557 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00792E69 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00793721 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00793FD9 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00794891 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00795149 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00795A01 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007962B9 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00796B71 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00796BBA | `HandleTVOutChanged` | Known | Event handler |
| 0x00796BF2 | `HandleTVSignalChanged` | Known | Event handler |
| 0x00796C2D | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x00796C7E | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x00796CC3 | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x00796D0C | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x00796D4E | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x00796D98 | `HandlePlayPause` | Known | Event handler |
| 0x00796DCE | `HandleShowContextualMenu` | Known | Event handler |
| 0x00796E15 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00796E58 | `HandleAddToOTG` | Known | Event handler |
| 0x00796EBB | `HandleStartGenius` | Known | Event handler |
| 0x00796EF7 | `HandleViewAlbum` | Known | Event handler |
| 0x00796F32 | `HandleViewArtist` | Known | Event handler |
| 0x00796F73 | `HandleViewCompilation` | Known | Event handler |
| 0x007971AF | `HandlePlayPause` | Known | Event handler |
| 0x007971D6 | `HandleSelect` | Known | Event handler |
| 0x00797208 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x00797243 | `HandleSaveGeniusPlaylist` | Known | Event handler |
| 0x007972E4 | `HandlePlayPause` | Known | Event handler |
| 0x0079731A | `HandleShowContextualMenu` | Known | Event handler |
| 0x00797361 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007973A4 | `HandleAddToOTG` | Known | Event handler |
| 0x00797407 | `HandleStartGenius` | Known | Event handler |
| 0x00797443 | `HandleViewAlbum` | Known | Event handler |
| 0x0079747E | `HandleViewArtist` | Known | Event handler |
| 0x007974BF | `HandleViewCompilation` | Known | Event handler |
| 0x0079752D | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x00797955 | `HandlePlayPause` | Known | Event handler |
| 0x0079797C | `HandleSelect` | Known | Event handler |
| 0x007979AE | `HandleRefreshPlaylist` | Known | Event handler |
| 0x007979E5 | `HandleSelect` | Known | Event handler |
| 0x00797A15 | `HandleSelect` | Known | Event handler |
| 0x00797A4D | `HandleMenuLongpress` | Known | Event handler |
| 0x00797A7B | `HandleMenuKey` | Known | Event handler |
| 0x00797B01 | `HandlePlayPause` | Known | Event handler |
| 0x00797B8B | `HandlePushContextualMenu` | Known | Event handler |
| 0x00797BC0 | `HandleSelect` | Known | Event handler |
| 0x00797BFB | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00797C3E | `HandleAddToOTG` | Known | Event handler |
| 0x00797C7D | `HandleAudiobookFaster` | Known | Event handler |
| 0x00797CC3 | `HandleAudiobookNormal` | Known | Event handler |
| 0x00797D09 | `HandleAudiobookSlower` | Known | Event handler |
| 0x00797D73 | `HandleStartGenius` | Known | Event handler |
| 0x00797DAF | `HandleViewAlbum` | Known | Event handler |
| 0x00797DEA | `HandleViewArtist` | Known | Event handler |
| 0x00797E2B | `HandleViewCompilation` | Known | Event handler |
| 0x0079886D | `HandleStartGenius` | Known | Event handler |
| 0x00798980 | `HandlePlayPause` | Known | Event handler |
| 0x007989F5 | `HandleWheelProgress` | Known | Event handler |
| 0x00798A31 | `HandleMenuLongpress` | Known | Event handler |
| 0x00798A5F | `HandleMenuKey` | Known | Event handler |
| 0x00798AE5 | `HandlePlayPause` | Known | Event handler |
| 0x00798B6F | `HandlePushContextualMenu` | Known | Event handler |
| 0x00798BA4 | `HandleSelectProgress` | Known | Event handler |
| 0x00798BE7 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00798C2A | `HandleAddToOTG` | Known | Event handler |
| 0x00798C69 | `HandleAudiobookFaster` | Known | Event handler |
| 0x00798CAF | `HandleAudiobookNormal` | Known | Event handler |
| 0x00798CF5 | `HandleAudiobookSlower` | Known | Event handler |
| 0x00798D5F | `HandleStartGenius` | Known | Event handler |
| 0x00798D9B | `HandleViewAlbum` | Known | Event handler |
| 0x00798DD6 | `HandleViewArtist` | Known | Event handler |
| 0x00798E17 | `HandleViewCompilation` | Known | Event handler |
| 0x00799859 | `HandleStartGenius` | Known | Event handler |
| 0x0079996C | `HandlePlayPause` | Known | Event handler |
| 0x007999E1 | `HandleWheelProgress` | Known | Event handler |
| 0x00799A1D | `HandleMenuLongpress` | Known | Event handler |
| 0x00799A4B | `HandleMenuKey` | Known | Event handler |
| 0x00799AD1 | `HandlePlayPause` | Known | Event handler |
| 0x00799B5B | `HandlePushContextualMenu` | Known | Event handler |
| 0x00799B90 | `HandleSelectVolume` | Known | Event handler |
| 0x00799BD1 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00799C14 | `HandleAddToOTG` | Known | Event handler |
| 0x00799C53 | `HandleAudiobookFaster` | Known | Event handler |
| 0x00799C99 | `HandleAudiobookNormal` | Known | Event handler |
| 0x00799CDF | `HandleAudiobookSlower` | Known | Event handler |
| 0x00799D49 | `HandleStartGenius` | Known | Event handler |
| 0x00799D85 | `HandleViewAlbum` | Known | Event handler |
| 0x00799DC0 | `HandleViewArtist` | Known | Event handler |
| 0x00799E01 | `HandleViewCompilation` | Known | Event handler |
| 0x0079A843 | `HandleStartGenius` | Known | Event handler |
| 0x0079A956 | `HandlePlayPause` | Known | Event handler |
| 0x0079A9CB | `HandleWheelVolume` | Known | Event handler |
| 0x0079AA05 | `HandleMenuLongpress` | Known | Event handler |
| 0x0079AA33 | `HandleMenuKey` | Known | Event handler |
| 0x0079AAB9 | `HandlePlayPause` | Known | Event handler |
| 0x0079AB43 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079AB78 | `HandleSelectRating` | Known | Event handler |
| 0x0079ABB9 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079ABFC | `HandleAddToOTG` | Known | Event handler |
| 0x0079AC3B | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079AC81 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079ACC7 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079AD31 | `HandleStartGenius` | Known | Event handler |
| 0x0079AD6D | `HandleViewAlbum` | Known | Event handler |
| 0x0079ADA8 | `HandleViewArtist` | Known | Event handler |
| 0x0079ADE9 | `HandleViewCompilation` | Known | Event handler |
| 0x0079B82B | `HandleStartGenius` | Known | Event handler |
| 0x0079B93E | `HandlePlayPause` | Known | Event handler |
| 0x0079B9B3 | `HandleWheelRating` | Known | Event handler |
| 0x0079B9ED | `HandleMenuLongpress` | Known | Event handler |
| 0x0079BA1B | `HandleMenuKey` | Known | Event handler |
| 0x0079BA93 | `HandlePlayPause` | Known | Event handler |
| 0x0079BB14 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079BB49 | `HandleSelectScrub` | Known | Event handler |
| 0x0079BB89 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079BBCC | `HandleAddToOTG` | Known | Event handler |
| 0x0079BC0B | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079BC51 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079BC97 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079BD01 | `HandleStartGenius` | Known | Event handler |
| 0x0079BD3D | `HandleViewAlbum` | Known | Event handler |
| 0x0079BD78 | `HandleViewArtist` | Known | Event handler |
| 0x0079BDB9 | `HandleViewCompilation` | Known | Event handler |
| 0x0079C7FB | `HandleStartGenius` | Known | Event handler |
| 0x0079C900 | `HandlePlayPause` | Known | Event handler |
| 0x0079C96C | `HandleWheelScrub` | Known | Event handler |
| 0x0079C9A5 | `HandleMenuLongpress` | Known | Event handler |
| 0x0079C9D3 | `HandleMenuKey` | Known | Event handler |
| 0x0079CA59 | `HandlePlayPause` | Known | Event handler |
| 0x0079CAE3 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079CB18 | `HandleSelectGenius` | Known | Event handler |
| 0x0079CB59 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079CB9C | `HandleAddToOTG` | Known | Event handler |
| 0x0079CBDB | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079CC21 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079CC67 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079CCD1 | `HandleStartGenius` | Known | Event handler |
| 0x0079CD0D | `HandleViewAlbum` | Known | Event handler |
| 0x0079CD48 | `HandleViewArtist` | Known | Event handler |
| 0x0079CD89 | `HandleViewCompilation` | Known | Event handler |
| 0x0079D7CB | `HandleStartGenius` | Known | Event handler |
| 0x0079D8DE | `HandlePlayPause` | Known | Event handler |
| 0x0079D953 | `HandleWheelGenius` | Known | Event handler |
| 0x0079D98D | `HandleMenuLongpress` | Known | Event handler |
| 0x0079D9BB | `HandleMenuKey` | Known | Event handler |
| 0x0079DA18 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0079DA50 | `HandlePlayPause` | Known | Event handler |
| 0x0079DAAA | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0079DAE9 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079DB1E | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x0079DB66 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079DBA9 | `HandleAddToOTG` | Known | Event handler |
| 0x0079DBE8 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079DC2E | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079DC74 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079DCDE | `HandleStartGenius` | Known | Event handler |
| 0x0079DD1A | `HandleViewAlbum` | Known | Event handler |
| 0x0079DD55 | `HandleViewArtist` | Known | Event handler |
| 0x0079DD96 | `HandleViewCompilation` | Known | Event handler |
| 0x0079E7D8 | `HandleStartGenius` | Known | Event handler |
| 0x0079E8EB | `HandlePlayPause` | Known | Event handler |
| 0x0079E960 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0079E9A1 | `HandleMenuLongpress` | Known | Event handler |
| 0x0079E9CF | `HandleMenuKey` | Known | Event handler |
| 0x0079EA55 | `HandlePlayPause` | Known | Event handler |
| 0x0079EADF | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079EB14 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x0079EB58 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079EB9B | `HandleAddToOTG` | Known | Event handler |
| 0x0079EBDA | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079EC20 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079EC66 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079ECD0 | `HandleStartGenius` | Known | Event handler |
| 0x0079ED0C | `HandleViewAlbum` | Known | Event handler |
| 0x0079ED47 | `HandleViewArtist` | Known | Event handler |
| 0x0079ED88 | `HandleViewCompilation` | Known | Event handler |
| 0x0079F7CA | `HandleStartGenius` | Known | Event handler |
| 0x0079F8DD | `HandlePlayPause` | Known | Event handler |
| 0x0079F97D | `HandleMenuLongpress` | Known | Event handler |
| 0x0079F9AB | `HandleMenuKey` | Known | Event handler |
| 0x0079FA31 | `HandlePlayPause` | Known | Event handler |
| 0x0079FABB | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079FAF0 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x0079FB34 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079FB77 | `HandleAddToOTG` | Known | Event handler |
| 0x0079FBB6 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079FBFC | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079FC42 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079FCAC | `HandleStartGenius` | Known | Event handler |
| 0x0079FCE8 | `HandleViewAlbum` | Known | Event handler |
| 0x0079FD23 | `HandleViewArtist` | Known | Event handler |
| 0x0079FD64 | `HandleViewCompilation` | Known | Event handler |
| 0x007A07A6 | `HandleStartGenius` | Known | Event handler |
| 0x007A08B9 | `HandlePlayPause` | Known | Event handler |
| 0x007A0959 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A0987 | `HandleMenuKey` | Known | Event handler |
| 0x007A0A0D | `HandlePlayPause` | Known | Event handler |
| 0x007A0A97 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A0ACC | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007A0B10 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A0B53 | `HandleAddToOTG` | Known | Event handler |
| 0x007A0B92 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A0BD8 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A0C1E | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A0C88 | `HandleStartGenius` | Known | Event handler |
| 0x007A0CC4 | `HandleViewAlbum` | Known | Event handler |
| 0x007A0CFF | `HandleViewArtist` | Known | Event handler |
| 0x007A0D40 | `HandleViewCompilation` | Known | Event handler |
| 0x007A1782 | `HandleStartGenius` | Known | Event handler |
| 0x007A1895 | `HandlePlayPause` | Known | Event handler |
| 0x007A1935 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A1963 | `HandleMenuKey` | Known | Event handler |
| 0x007A19E9 | `HandlePlayPause` | Known | Event handler |
| 0x007A1A73 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A1AA8 | `HandleSelectChapterArt` | Known | Event handler |
| 0x007A1AED | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A1B30 | `HandleAddToOTG` | Known | Event handler |
| 0x007A1B6F | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A1BB5 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A1BFB | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A1C65 | `HandleStartGenius` | Known | Event handler |
| 0x007A1CA1 | `HandleViewAlbum` | Known | Event handler |
| 0x007A1CDC | `HandleViewArtist` | Known | Event handler |
| 0x007A1D1D | `HandleViewCompilation` | Known | Event handler |
| 0x007A275F | `HandleStartGenius` | Known | Event handler |
| 0x007A2872 | `HandlePlayPause` | Known | Event handler |
| 0x007A28E7 | `HandleWheelVolume` | Known | Event handler |
| 0x007A2921 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A294F | `HandleMenuKey` | Known | Event handler |
| 0x007A29DE | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007A2A7F | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A2AB4 | `HandleSelect` | Known | Event handler |
| 0x007A2AEF | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A2B32 | `HandleAddToOTG` | Known | Event handler |
| 0x007A2B71 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A2BB7 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A2BFD | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A2C67 | `HandleStartGenius` | Known | Event handler |
| 0x007A2CA3 | `HandleViewAlbum` | Known | Event handler |
| 0x007A2CDE | `HandleViewArtist` | Known | Event handler |
| 0x007A2D1F | `HandleViewCompilation` | Known | Event handler |
| 0x007A3761 | `HandleStartGenius` | Known | Event handler |
| 0x007A387D | `HandlePlayPause` | Known | Event handler |
| 0x007A38FB | `HandleWheel` | Known | Event handler |
| 0x007A3931 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A395F | `HandleMenuKey` | Known | Event handler |
| 0x007A39EE | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007A3A8F | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A3AC4 | `HandleSelect` | Known | Event handler |
| 0x007A3AFF | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A3B42 | `HandleAddToOTG` | Known | Event handler |
| 0x007A3B81 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A3BC7 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A3C0D | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A3C77 | `HandleStartGenius` | Known | Event handler |
| 0x007A3CB3 | `HandleViewAlbum` | Known | Event handler |
| 0x007A3CEE | `HandleViewArtist` | Known | Event handler |
| 0x007A3D2F | `HandleViewCompilation` | Known | Event handler |
| 0x007A4771 | `HandleStartGenius` | Known | Event handler |
| 0x007A488D | `HandlePlayPause` | Known | Event handler |
| 0x007A490B | `HandleWheel` | Known | Event handler |
| 0x007A4941 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A496F | `HandleMenuKey` | Known | Event handler |
| 0x007A49F5 | `HandlePlayPause` | Known | Event handler |
| 0x007A4A7F | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A4AB4 | `HandleSelect` | Known | Event handler |
| 0x007A4AEF | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A4B32 | `HandleAddToOTG` | Known | Event handler |
| 0x007A4B71 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A4BB7 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A4BFD | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A4C67 | `HandleStartGenius` | Known | Event handler |
| 0x007A4CA3 | `HandleViewAlbum` | Known | Event handler |
| 0x007A4CDE | `HandleViewArtist` | Known | Event handler |
| 0x007A4D1F | `HandleViewCompilation` | Known | Event handler |
| 0x007A5761 | `HandleStartGenius` | Known | Event handler |
| 0x007A5874 | `HandlePlayPause` | Known | Event handler |
| 0x007A58E9 | `HandleWheel` | Known | Event handler |
| 0x007A591D | `HandleMenuLongpress` | Known | Event handler |
| 0x007A594B | `HandleMenuKey` | Known | Event handler |
| 0x007A59D1 | `HandlePlayPause` | Known | Event handler |
| 0x007A5A5B | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A5A90 | `HandleSelectProgress` | Known | Event handler |
| 0x007A5AD3 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A5B16 | `HandleAddToOTG` | Known | Event handler |
| 0x007A5B55 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A5B9B | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A5BE1 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A5C4B | `HandleStartGenius` | Known | Event handler |
| 0x007A5C87 | `HandleViewAlbum` | Known | Event handler |
| 0x007A5CC2 | `HandleViewArtist` | Known | Event handler |
| 0x007A5D03 | `HandleViewCompilation` | Known | Event handler |
| 0x007A6745 | `HandleStartGenius` | Known | Event handler |
| 0x007A6858 | `HandlePlayPause` | Known | Event handler |
| 0x007A68CD | `HandleWheelProgress` | Known | Event handler |
| 0x007A6909 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A6937 | `HandleMenuKey` | Known | Event handler |
| 0x007A69AF | `HandlePlayPause` | Known | Event handler |
| 0x007A6A30 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A6A65 | `HandleSelectScrub` | Known | Event handler |
| 0x007A6AA5 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A6AE8 | `HandleAddToOTG` | Known | Event handler |
| 0x007A6B27 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A6B6D | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A6BB3 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A6C1D | `HandleStartGenius` | Known | Event handler |
| 0x007A6C59 | `HandleViewAlbum` | Known | Event handler |
| 0x007A6C94 | `HandleViewArtist` | Known | Event handler |
| 0x007A6CD5 | `HandleViewCompilation` | Known | Event handler |
| 0x007A7717 | `HandleStartGenius` | Known | Event handler |
| 0x007A781C | `HandlePlayPause` | Known | Event handler |
| 0x007A7888 | `HandleWheelScrub` | Known | Event handler |
| 0x007A78C1 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A78EF | `HandleMenuKey` | Known | Event handler |
| 0x007A7975 | `HandlePlayPause` | Known | Event handler |
| 0x007A79FF | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A7A6E | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A7AB1 | `HandleAddToOTG` | Known | Event handler |
| 0x007A7AF0 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A7B36 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A7B7C | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A7BE6 | `HandleStartGenius` | Known | Event handler |
| 0x007A7C22 | `HandleViewAlbum` | Known | Event handler |
| 0x007A7C5D | `HandleViewArtist` | Known | Event handler |
| 0x007A7C9E | `HandleViewCompilation` | Known | Event handler |
| 0x007A86E0 | `HandleStartGenius` | Known | Event handler |
| 0x007A87F3 | `HandlePlayPause` | Known | Event handler |
| 0x007A8868 | `HandleWheelVolume` | Known | Event handler |
| 0x007A88A5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A88D3 | `HandleMenuKey` | Known | Event handler |
| 0x007A8959 | `HandlePlayPause` | Known | Event handler |
| 0x007A89E3 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A8A52 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A8A95 | `HandleAddToOTG` | Known | Event handler |
| 0x007A8AD4 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A8B1A | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A8B60 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A8BCA | `HandleStartGenius` | Known | Event handler |
| 0x007A8C06 | `HandleViewAlbum` | Known | Event handler |
| 0x007A8C41 | `HandleViewArtist` | Known | Event handler |
| 0x007A8C82 | `HandleViewCompilation` | Known | Event handler |
| 0x007A96C4 | `HandleStartGenius` | Known | Event handler |
| 0x007A97D7 | `HandlePlayPause` | Known | Event handler |
| 0x007A984C | `HandleWheelBrightness` | Known | Event handler |
| 0x007A996F | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A99A4 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007A99EC | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A9A2F | `HandleAddToOTG` | Known | Event handler |
| 0x007A9A6E | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A9AB4 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A9AFA | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A9B64 | `HandleStartGenius` | Known | Event handler |
| 0x007A9BA0 | `HandleViewAlbum` | Known | Event handler |
| 0x007A9BDB | `HandleViewArtist` | Known | Event handler |
| 0x007A9C1C | `HandleViewCompilation` | Known | Event handler |
| 0x007AA65E | `HandleStartGenius` | Known | Event handler |
| 0x007AA7AA | `HandleWheel` | Known | Event handler |
| 0x007AA7E1 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AA80F | `HandleMenuKey` | Known | Event handler |
| 0x007AA895 | `HandlePlayPause` | Known | Event handler |
| 0x007AA915 | `HandleSelect` | Known | Event handler |
| 0x007AADB7 | `HandlePlayPause` | Known | Event handler |
| 0x007AAE45 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AAE73 | `HandleMenuKey` | Known | Event handler |
| 0x007AAEF9 | `HandlePlayPause` | Known | Event handler |
| 0x007AAF79 | `HandleSelectProgress` | Known | Event handler |
| 0x007AB423 | `HandlePlayPause` | Known | Event handler |
| 0x007AB498 | `HandleWheelProgress` | Known | Event handler |
| 0x007AB4D5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AB503 | `HandleMenuKey` | Known | Event handler |
| 0x007AB589 | `HandlePlayPause` | Known | Event handler |
| 0x007AB609 | `HandleSelectProgress` | Known | Event handler |
| 0x007ABAB3 | `HandlePlayPause` | Known | Event handler |
| 0x007ABB28 | `HandleWheelProgress` | Known | Event handler |
| 0x007ABB65 | `HandleMenuLongpress` | Known | Event handler |
| 0x007ABB93 | `HandleMenuKey` | Known | Event handler |
| 0x007ABC19 | `HandlePlayPause` | Known | Event handler |
| 0x007ABC99 | `HandleSelectProgress` | Known | Event handler |
| 0x007AC0CF | `HandlePlayPause` | Known | Event handler |
| 0x007AC144 | `HandleWheelProgress` | Known | Event handler |
| 0x007AC181 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AC1AF | `HandleMenuKey` | Known | Event handler |
| 0x007AC21C | `HandlePlayPause` | Known | Event handler |
| 0x007AC288 | `HandleSelectScrub` | Known | Event handler |
| 0x007AC6A2 | `HandlePlayPause` | Known | Event handler |
| 0x007AC703 | `HandleWheelScrub` | Known | Event handler |
| 0x007AC73D | `HandleMenuLongpress` | Known | Event handler |
| 0x007AC76B | `HandleMenuKey` | Known | Event handler |
| 0x007AC7F1 | `HandlePlayPause` | Known | Event handler |
| 0x007AC871 | `HandleSelectVolume` | Known | Event handler |
| 0x007ACCA5 | `HandlePlayPause` | Known | Event handler |
| 0x007ACD1A | `HandleWheelVolume` | Known | Event handler |
| 0x007ACE2D | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007AD2CC | `HandleSelect` | Known | Event handler |
| 0x007AD2F9 | `HandleSelect` | Known | Event handler |
| 0x007AD329 | `HandleSelect` | Known | Event handler |
| 0x007AD359 | `HandleSelect` | Known | Event handler |
| 0x007AD389 | `HandleSelect` | Known | Event handler |
| 0x007AD3B9 | `HandleSelect` | Known | Event handler |
| 0x007AD3E9 | `HandleSelect` | Known | Event handler |
| 0x007AD419 | `HandleSelect` | Known | Event handler |
| 0x007AD449 | `HandleSelect` | Known | Event handler |
| 0x007AD4B9 | `HandleSelect` | Known | Event handler |
| 0x007AD4E9 | `HandleSelect` | Known | Event handler |
| 0x007AD561 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AD594 | `HandleNotesPop` | Known | Event handler |
| 0x007AD611 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AD644 | `HandleNotesPop` | Known | Event handler |
| 0x007ADB00 | `HandleNotesSelected` | Known | Event handler |
| 0x007ADB3D | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007ADB70 | `HandleNotesPop` | Known | Event handler |
| 0x007AE02C | `HandleNotesSelected` | Known | Event handler |
| 0x007AE069 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AE09C | `HandleNotesPop` | Known | Event handler |
| 0x007AE0C7 | `HandleNotesSelected` | Known | Event handler |
| 0x007AE599 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AE5CC | `HandleNotesPop` | Known | Event handler |
| 0x007AE5F7 | `HandleNotesSelected` | Known | Event handler |
| 0x007AEAC9 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AEAFC | `HandleNotesPop` | Known | Event handler |
| 0x007AEB79 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AEBAC | `HandleNotesPop` | Known | Event handler |
| 0x007AEC29 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AEC5C | `HandleNotesPop` | Known | Event handler |
| 0x007AECD4 | `HandlePlayPause` | Known | Event handler |
| 0x007AECFD | `HandlePlayPause` | Known | Event handler |
| 0x007AED2B | `HandlePlayPause` | Known | Event handler |
| 0x007AED60 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007AEDE0 | `HandleHiliteAlbum` | Known | Event handler |
| 0x007AEE89 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007AEF10 | `HandleHiliteAlbum` | Known | Event handler |
| 0x007AF1D4 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x007AF230 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x007AF3E7 | `HandleSelect` | Known | Event handler |
| 0x007AF56B | `HandleSelect` | Known | Event handler |
| 0x007AF5A5 | `HandleImageLast` | Known | Event handler |
| 0x007AF5CF | `HandleImageNext` | Known | Event handler |
| 0x007AF5FE | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AF638 | `HandleImageFirst` | Known | Event handler |
| 0x007AF663 | `HandleImagePrev` | Known | Event handler |
| 0x007AF68F | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AF6BE | `HandleImageNext` | Known | Event handler |
| 0x007AF6E7 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AF71B | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AF74A | `HandleImagePrev` | Known | Event handler |
| 0x007AF76B | `HandleImageWheel` | Known | Event handler |
| 0x007AF809 | `HandleImageNext` | Known | Event handler |
| 0x007AF838 | `HandlePlayPause` | Known | Event handler |
| 0x007AF887 | `HandleImagePrev` | Known | Event handler |
| 0x007AF8B3 | `HandleSelect` | Known | Event handler |
| 0x007AFB83 | `HandleImageNext` | Known | Event handler |
| 0x007AFBAD | `HandlePause` | Known | Event handler |
| 0x007AFBD2 | `HandlePlay` | Known | Event handler |
| 0x007AFBFB | `HandlePlayPause` | Known | Event handler |
| 0x007AFC24 | `HandleImagePrev` | Known | Event handler |
| 0x007AFC87 | `HandleMikeyCenter` | Known | Event handler |
| 0x007AFCAA | `HandleWheel` | Known | Event handler |
| 0x007AFD45 | `HandleImageNext` | Known | Event handler |
| 0x007AFD74 | `HandlePlayPause` | Known | Event handler |
| 0x007AFDC3 | `HandleImagePrev` | Known | Event handler |
| 0x007AFDEF | `HandleSelect` | Known | Event handler |
| 0x007B00BF | `HandleImageNext` | Known | Event handler |
| 0x007B00E9 | `HandlePause` | Known | Event handler |
| 0x007B010E | `HandlePlay` | Known | Event handler |
| 0x007B0137 | `HandlePlayPause` | Known | Event handler |
| 0x007B0160 | `HandleImagePrev` | Known | Event handler |
| 0x007B01C3 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B01E6 | `HandleWheel` | Known | Event handler |
| 0x007B0281 | `HandleImageNext` | Known | Event handler |
| 0x007B02B0 | `HandlePlayPause` | Known | Event handler |
| 0x007B02FF | `HandleImagePrev` | Known | Event handler |
| 0x007B032B | `HandleSelect` | Known | Event handler |
| 0x007B05FB | `HandleImageNext` | Known | Event handler |
| 0x007B0625 | `HandlePause` | Known | Event handler |
| 0x007B064A | `HandlePlay` | Known | Event handler |
| 0x007B0673 | `HandlePlayPause` | Known | Event handler |
| 0x007B069C | `HandleImagePrev` | Known | Event handler |
| 0x007B06FF | `HandleMikeyCenter` | Known | Event handler |
| 0x007B0722 | `HandleWheel` | Known | Event handler |
| 0x007B07BD | `HandleImageNext` | Known | Event handler |
| 0x007B07EC | `HandlePlayPause` | Known | Event handler |
| 0x007B083B | `HandleImagePrev` | Known | Event handler |
| 0x007B0867 | `HandleSelect` | Known | Event handler |
| 0x007B0B37 | `HandleImageNext` | Known | Event handler |
| 0x007B0B61 | `HandlePause` | Known | Event handler |
| 0x007B0B86 | `HandlePlay` | Known | Event handler |
| 0x007B0BAF | `HandlePlayPause` | Known | Event handler |
| 0x007B0BD8 | `HandleImagePrev` | Known | Event handler |
| 0x007B0C3B | `HandleMikeyCenter` | Known | Event handler |
| 0x007B0C5E | `HandleWheel` | Known | Event handler |
| 0x007B0CF9 | `HandleImageNext` | Known | Event handler |
| 0x007B0D28 | `HandlePlayPause` | Known | Event handler |
| 0x007B0D77 | `HandleImagePrev` | Known | Event handler |
| 0x007B0DA3 | `HandleSelect` | Known | Event handler |
| 0x007B1073 | `HandleImageNext` | Known | Event handler |
| 0x007B109D | `HandlePause` | Known | Event handler |
| 0x007B10C2 | `HandlePlay` | Known | Event handler |
| 0x007B10EB | `HandlePlayPause` | Known | Event handler |
| 0x007B1114 | `HandleImagePrev` | Known | Event handler |
| 0x007B1177 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B119A | `HandleWheel` | Known | Event handler |
| 0x007B1235 | `HandleImageNext` | Known | Event handler |
| 0x007B1264 | `HandlePlayPause` | Known | Event handler |
| 0x007B12B3 | `HandleImagePrev` | Known | Event handler |
| 0x007B12DF | `HandleSelect` | Known | Event handler |
| 0x007B15AF | `HandleImageNext` | Known | Event handler |
| 0x007B15D9 | `HandlePause` | Known | Event handler |
| 0x007B15FE | `HandlePlay` | Known | Event handler |
| 0x007B1627 | `HandlePlayPause` | Known | Event handler |
| 0x007B1650 | `HandleImagePrev` | Known | Event handler |
| 0x007B16B3 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B16D6 | `HandleWheel` | Known | Event handler |
| 0x007B1771 | `HandleImageNext` | Known | Event handler |
| 0x007B17A0 | `HandlePlayPause` | Known | Event handler |
| 0x007B17EF | `HandleImagePrev` | Known | Event handler |
| 0x007B181B | `HandleSelect` | Known | Event handler |
| 0x007B1A66 | `HandleImageNext` | Known | Event handler |
| 0x007B1A90 | `HandlePause` | Known | Event handler |
| 0x007B1AB5 | `HandlePlay` | Known | Event handler |
| 0x007B1ADE | `HandlePlayPause` | Known | Event handler |
| 0x007B1B07 | `HandleImagePrev` | Known | Event handler |
| 0x007B1B7A | `HandleMikeyCenter` | Known | Event handler |
| 0x007B1B9D | `HandleWheel` | Known | Event handler |
| 0x007B1C35 | `HandleImageNext` | Known | Event handler |
| 0x007B1C64 | `HandlePlayPause` | Known | Event handler |
| 0x007B1CB3 | `HandleImagePrev` | Known | Event handler |
| 0x007B1CDF | `HandleSelect` | Known | Event handler |
| 0x007B1F2A | `HandleImageNext` | Known | Event handler |
| 0x007B1F54 | `HandlePause` | Known | Event handler |
| 0x007B1F79 | `HandlePlay` | Known | Event handler |
| 0x007B1FA2 | `HandlePlayPause` | Known | Event handler |
| 0x007B1FCB | `HandleImagePrev` | Known | Event handler |
| 0x007B203E | `HandleMikeyCenter` | Known | Event handler |
| 0x007B2061 | `HandleWheel` | Known | Event handler |
| 0x007B20F9 | `HandleImageNext` | Known | Event handler |
| 0x007B2128 | `HandlePlayPause` | Known | Event handler |
| 0x007B2177 | `HandleImagePrev` | Known | Event handler |
| 0x007B21A3 | `HandleSelect` | Known | Event handler |
| 0x007B23EE | `HandleImageNext` | Known | Event handler |
| 0x007B2418 | `HandlePause` | Known | Event handler |
| 0x007B243D | `HandlePlay` | Known | Event handler |
| 0x007B2466 | `HandlePlayPause` | Known | Event handler |
| 0x007B248F | `HandleImagePrev` | Known | Event handler |
| 0x007B2502 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B2525 | `HandleWheel` | Known | Event handler |
| 0x007B25BD | `HandleImageNext` | Known | Event handler |
| 0x007B25EC | `HandlePlayPause` | Known | Event handler |
| 0x007B263B | `HandleImagePrev` | Known | Event handler |
| 0x007B2667 | `HandleSelect` | Known | Event handler |
| 0x007B28B2 | `HandleImageNext` | Known | Event handler |
| 0x007B28DC | `HandlePause` | Known | Event handler |
| 0x007B2901 | `HandlePlay` | Known | Event handler |
| 0x007B292A | `HandlePlayPause` | Known | Event handler |
| 0x007B2953 | `HandleImagePrev` | Known | Event handler |
| 0x007B29C6 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B29E9 | `HandleWheel` | Known | Event handler |
| 0x007B2A81 | `HandleImageNext` | Known | Event handler |
| 0x007B2AB0 | `HandlePlayPause` | Known | Event handler |
| 0x007B2AFF | `HandleImagePrev` | Known | Event handler |
| 0x007B2B2B | `HandleSelect` | Known | Event handler |
| 0x007B2D76 | `HandleImageNext` | Known | Event handler |
| 0x007B2DA0 | `HandlePause` | Known | Event handler |
| 0x007B2DC5 | `HandlePlay` | Known | Event handler |
| 0x007B2DEE | `HandlePlayPause` | Known | Event handler |
| 0x007B2E17 | `HandleImagePrev` | Known | Event handler |
| 0x007B2E8A | `HandleMikeyCenter` | Known | Event handler |
| 0x007B2EAD | `HandleWheel` | Known | Event handler |
| 0x007B2ED9 | `HandleSelect` | Known | Event handler |
| 0x007B2F09 | `HandleSelect` | Known | Event handler |
| 0x007B302C | `HandleTuning` | Known | Event handler |
| 0x007B31EC | `HandleVolumeChange` | Known | Event handler |
| 0x007B3253 | `HandleVolumeChange` | Known | Event handler |
| 0x007B32B8 | `HandleVolumeChange` | Known | Event handler |
| 0x007B3404 | `HandleVolumeWheel` | Known | Event handler |
| 0x007B355F | `HandleTuningSelect` | Known | Event handler |
| 0x007B3725 | `HandleVolumeChange` | Known | Event handler |
| 0x007B378C | `HandleVolumeChange` | Known | Event handler |
| 0x007B37F1 | `HandleVolumeChange` | Known | Event handler |
| 0x007B393D | `HandleFrequencyChange` | Known | Event handler |
| 0x007B3A9B | `HandleTuningSelect` | Known | Event handler |
| 0x007B3C61 | `HandleVolumeChange` | Known | Event handler |
| 0x007B3CC8 | `HandleVolumeChange` | Known | Event handler |
| 0x007B3D2D | `HandleVolumeChange` | Known | Event handler |
| 0x007B3E79 | `HandleFrequencyChange` | Known | Event handler |
| 0x007B3FA4 | `HandleTimerDone` | Known | Event handler |
| 0x007B419D | `HandleVolumeChange` | Known | Event handler |
| 0x007B41CF | `HandleVolumeChange` | Known | Event handler |
| 0x007B41FF | `HandleVolumeChange` | Known | Event handler |
| 0x007B4316 | `HandleVolumeWheel` | Known | Event handler |
| 0x007B4B67 | `HandleExitUnsupported` | Known | Event handler |
| 0x007B4B99 | `HandleExitUnsupported` | Known | Event handler |
| 0x007B9BCD | `HandleSelectKey` | Known | Event handler |
| 0x007B9C02 | `HandleWheel` | Known | Event handler |
| 0x007B9D50 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x007B9DA3 | `HandleSelectKey` | Known | Event handler |
| 0x007B9DCB | `HandleSelectKey` | Known | Event handler |
| 0x007B9DFB | `HandleExit` | Known | Event handler |
| 0x007B9E25 | `HandleStartStop` | Known | Event handler |
| 0x007B9E8B | `HandleStartStop` | Known | Event handler |
| 0x007B9FA3 | `HandleExit` | Known | Event handler |
| 0x007B9FCD | `HandleStartStop` | Known | Event handler |
| 0x007B9FF9 | `HandleLap` | Known | Event handler |
| 0x007BA0FD | `HandleSelectLozinch` | Known | Event handler |
| 0x007BA31A | `HandleSelect` | Known | Event handler |
| 0x007BA3A6 | `HandleSelect` | Known | Event handler |
| 0x007BA434 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x007BA732 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x007BA81D | `HandleFinishRecording` | Known | Event handler |
| 0x007BA86E | `HandlePlayPause` | Known | Event handler |
| 0x007BA8FC | `HandlePlayPause` | Known | Event handler |
| 0x007BA98D | `HandleDeleteAllSelect` | Known | Event handler |
| 0x007BA9C5 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x007BAA01 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x007BAA44 | `HandlePlayPause` | Known | Event handler |
| 0x007BAA7A | `HandleAddToOTG` | Known | Event handler |
| 0x007BACCF | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007BAF2B | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007D7DEE | `HandleSelectClock` | Known | Event handler |
| 0x007D7E27 | `HandleHilited` | Known | Event handler |
| 0x007D7E59 | `HandleWheel` | Known | Event handler |
| 0x007D7EA0 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x007D7F25 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x007D8131 | `HandleImageLast` | Known | Event handler |
| 0x007D815B | `HandleScreenNext` | Known | Event handler |
| 0x007D818B | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D81C5 | `HandleImageFirst` | Known | Event handler |
| 0x007D81F0 | `HandleScreenPrev` | Known | Event handler |
| 0x007D821D | `HandleBrowseLarge` | Known | Event handler |
| 0x007D829D | `HandleImageNext` | Known | Event handler |
| 0x007D82C6 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D82FA | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D8329 | `HandleImagePrev` | Known | Event handler |
| 0x007D8357 | `HandleBrowseSmall` | Known | Event handler |

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
| 0x0013F2FC | `GotoMainMenu` | Known | Navigation |
| 0x001C2CD4 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001C50A4 | `GotoErrorLayout` | Known | Navigation |
| 0x001CE07C | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001CE740 | `GotoPlayDeleteMenu` | Known | Navigation |
| 0x001CE7C4 | `GotoNowPlaying` | Known | Navigation |
| 0x001E970C | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation |
| 0x001F5108 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x001F5200 | `GotoScreen_SettingsMenu` | Known | Navigation |
| 0x001FCD28 | `GotoDefaultLayout` | Known | Navigation |
| 0x001FCDAC | `GotoVolumeLayout` | Known | Navigation |
| 0x001FCEE4 | `GotoProgressLayout` | Known | Navigation |
| 0x001FD200 | `GotoDefault` | Known | Navigation |
| 0x001FD534 | `GotoProgressLayout` | Known | Navigation |
| 0x001FD6F4 | `GotoRentalWarningLayout` | Known | Navigation |
| 0x001FD778 | `GotoProgressLayout` | Known | Navigation |
| 0x001FDA88 | `GotoProgressLayout` | Known | Navigation |
| 0x001FF614 | `GotoNowPlaying` | Known | Navigation |
| 0x001FFF24 | `GotoNowPlaying` | Known | Navigation |
| 0x00200228 | `GotoNowPlaying` | Known | Navigation |
| 0x00202920 | `GotoScreen_Language` | Known | Navigation |
| 0x00202C80 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x00202C9C | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00202CB4 | `GotoDefaultLayout` | Known | Navigation |
| 0x00202CC8 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00202D60 | `GotoVolumeLayout` | Known | Navigation |
| 0x00202D74 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x00202E14 | `GotoProgressLayout` | Known | Navigation |
| 0x00202E28 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x002035DC | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00203A44 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x00203CB0 | `GotoProgressLayout` | Known | Navigation |
| 0x00203CC4 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00203E5C | `GotoBrightnessVideoLayout` | Known | Navigation |
| 0x00203E80 | `GotoGeniusLayout` | Known | Navigation |
| 0x00203E94 | `GotoRatingLayout` | Known | Navigation |
| 0x00204008 | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x00204024 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020403C | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x0020433C | `GotoChapterArtLayout` | Known | Navigation |
| 0x00204354 | `GotoShuffleLayout` | Known | Navigation |
| 0x002046E4 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x002046F8 | `GotoExtraInfoLoadingLayout` | Known | Navigation |
| 0x002047C8 | `GotoVolumeLayout` | Known | Navigation |
| 0x002047E0 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x0020486C | `GotoVolumeLayout` | Known | Navigation |
| 0x00204880 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x00204A90 | `GotoScrubLayout` | Known | Navigation |
| 0x00204AA0 | `GotoScrubVideoLayout` | Known | Navigation |
| 0x00204B30 | `GotoProgressLayout` | Known | Navigation |
| 0x00204B44 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00204D9C | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x00204DB8 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00204DD0 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00204DEC | `GotoDefaultLayout` | Known | Navigation |
| 0x00205018 | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x00205034 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x002055D0 | `GotoChapterArtLayout` | Known | Navigation |
| 0x002056C8 | `GotoProgressLayout` | Known | Navigation |
| 0x00205754 | `GotoProgressLayout` | Known | Navigation |
| 0x00205768 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00205844 | `GotoExtraInfoLoadFailedLayout` | Known | Navigation |
| 0x00205864 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x00205CA0 | `GotoStatusBarLayout` | Known | Navigation |
| 0x00205CB4 | `GotoDefaultLayout` | Known | Navigation |
| 0x00205E8C | `GotoDefault` | Known | Navigation |
| 0x00205FC0 | `GotoProgressLayout` | Known | Navigation |
| 0x00206180 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x002062D0 | `GotoBrightnessLayout` | Known | Navigation |
| 0x00206354 | `GotoBrightnessLayout` | Known | Navigation |
| 0x002063D4 | `GotoVolumeLayout` | Known | Navigation |
| 0x00206420 | `GotoScrubLayout` | Known | Navigation |
| 0x002064E8 | `GotoStatusBarLayout` | Known | Navigation |
| 0x002064FC | `GotoDefaultLayout` | Known | Navigation |
| 0x002065D4 | `GotoScrubLayout` | Known | Navigation |
| 0x00206624 | `GotoScrubLayout` | Known | Navigation |
| 0x002090F8 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020C56C | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x0020C588 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020C5A0 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x0020C750 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020CC44 | `GotoNowPlaying` | Known | Navigation |
| 0x0020CF2C | `GotoNowPlaying` | Known | Navigation |
| 0x0020E088 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x0020E218 | `GotoFourCard_About` | Known | Navigation |
| 0x0020E22C | `GotoThreeCard_About` | Known | Navigation |
| 0x0020E318 | `GotoScreen_BacklightTimer` | Known | Navigation |
| 0x0020E3A8 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x0020E3C0 | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation |
| 0x00212E18 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x00212E30 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x00213DC0 | `GotoGeniusIntro` | Known | Navigation |
| 0x00213DD4 | `GotoGenius` | Known | Navigation |
| 0x0021547C | `GotoNowPlaying` | Known | Navigation |
| 0x00215B8C | `GotoNowPlaying` | Known | Navigation |
| 0x00216370 | `GotoFirstBoot` | Known | Navigation |
| 0x00216380 | `GotoNotesApp` | Known | Navigation |
| 0x00216394 | `GotoLockApp` | Known | Navigation |
| 0x002176CC | `GotoGenius` | Known | Navigation |
| 0x0021D560 | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x0021D57C | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0021D594 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x0021D744 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0021DF20 | `GotoNowPlaying` | Known | Navigation |
| 0x003F171C | `GotoRatingLayout` | Known | Navigation |
| 0x003F1730 | `GotoProgressLayout` | Known | Navigation |
| 0x007260D7 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x007A7A34 | `GotoDefault` | Known | Navigation |
| 0x007A8A18 | `GotoDefault` | Known | Navigation |
| 0x00899A5C | `Gotowe?` | Known | Navigation |

---

## 7. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00156BBC | `CoverFlow_Screen` | Known | Screen layout |
| 0x0071A022 | `Clock_Screen` | Known | Screen layout |
| 0x0071A032 | `Clock_Screen_Default"` | Known | Screen layout |
| 0x0071A097 | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x0071A0F5 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0071A10D | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x0071A17A | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x0071A218 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x0071A277 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0071A28D | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x0071A2F8 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x0071A352 | `Games_Menu_Screen` | Known | Screen layout |
| 0x0071A367 | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x0071A3D1 | `Extras_Screen_Games` | Known | Screen layout |
| 0x0071A490 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x0071A554 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0071A61D | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x0071A67A | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x0071A693 | `Debug_MainMenu_Screen_Default"` | Known | Screen layout |
| 0x0071A701 | `Extras_Screen_Debug` | Known | Screen layout |
| 0x0071A840 | `Clock_Africa_City_Screen ` | Known | Screen layout |
| 0x0071A85C | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x0071A8E0 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x0071A8FA | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x0071A97C | `Clock_Atlantic_City_Screen"` | Known | Screen layout |
| 0x0071A99A | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x0071AA20 | `Clock_Australia_City_Screen#` | Known | Screen layout |
| 0x0071AA3F | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x0071AAC6 | `Clock_Europe_City_Screen ` | Known | Screen layout |
| 0x0071AAE2 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x0071AB66 | `Clock_NorthAmerica_City_Screen&` | Known | Screen layout |
| 0x0071AB88 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0071AC12 | `Clock_Pacific_City_Screen!` | Known | Screen layout |
| 0x0071AC2F | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x0071ACB4 | `Clock_SouthAmerica_City_Screen&` | Known | Screen layout |
| 0x0071ACD6 | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0071AD63 | `Clock_Screen"` | Known | Screen layout |
| 0x0071AE08 | `Clock_Screen"` | Known | Screen layout |
| 0x0071AEAD | `Clock_Screen"` | Known | Screen layout |
| 0x0071AF52 | `Clock_Screen"` | Known | Screen layout |
| 0x0071AFF7 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B09C | `Clock_Screen"` | Known | Screen layout |
| 0x0071B141 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B1E6 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B28B | `Clock_Screen"` | Known | Screen layout |
| 0x0071B330 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B3D5 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B47A | `Clock_Screen"` | Known | Screen layout |
| 0x0071B51F | `Clock_Screen"` | Known | Screen layout |
| 0x0071B5C4 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B669 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B70E | `Clock_Screen"` | Known | Screen layout |
| 0x0071B7B3 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B858 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B8FD | `Clock_Screen"` | Known | Screen layout |
| 0x0071B9A2 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BA47 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BAEC | `Clock_Screen"` | Known | Screen layout |
| 0x0071BB91 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BC36 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BCDB | `Clock_Screen"` | Known | Screen layout |
| 0x0071BD80 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BE25 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BECA | `Clock_Screen"` | Known | Screen layout |
| 0x0071BF6F | `Clock_Screen"` | Known | Screen layout |
| 0x0071C014 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C0B9 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C163 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C208 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C2AD | `Clock_Screen"` | Known | Screen layout |
| 0x0071C352 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C3F7 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C49C | `Clock_Screen"` | Known | Screen layout |
| 0x0071C541 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C5E6 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C68B | `Clock_Screen"` | Known | Screen layout |
| 0x0071C730 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C7D5 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C87A | `Clock_Screen"` | Known | Screen layout |
| 0x0071C91F | `Clock_Screen"` | Known | Screen layout |
| 0x0071C9C4 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CA69 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CB0E | `Clock_Screen"` | Known | Screen layout |
| 0x0071CBB3 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CC58 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CCFD | `Clock_Screen"` | Known | Screen layout |
| 0x0071CDA2 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CE47 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CEEC | `Clock_Screen"` | Known | Screen layout |
| 0x0071CF91 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D036 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D0DB | `Clock_Screen"` | Known | Screen layout |
| 0x0071D180 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D225 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D2CA | `Clock_Screen"` | Known | Screen layout |
| 0x0071D36F | `Clock_Screen"` | Known | Screen layout |
| 0x0071D414 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D4B9 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D55E | `Clock_Screen"` | Known | Screen layout |
| 0x0071D603 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D6A8 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D74D | `Clock_Screen"` | Known | Screen layout |
| 0x0071D7F2 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D897 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D93C | `Clock_Screen"` | Known | Screen layout |
| 0x0071D9E1 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DA86 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DB2B | `Clock_Screen"` | Known | Screen layout |
| 0x0071DBD0 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DC75 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DD1A | `Clock_Screen"` | Known | Screen layout |
| 0x0071DDBF | `Clock_Screen"` | Known | Screen layout |
| 0x0071DE64 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DF09 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DFAE | `Clock_Screen"` | Known | Screen layout |
| 0x0071E053 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E0F8 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E19D | `Clock_Screen"` | Known | Screen layout |
| 0x0071E242 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E2E7 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E38C | `Clock_Screen"` | Known | Screen layout |
| 0x0071E431 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E4D6 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E57B | `Clock_Screen"` | Known | Screen layout |
| 0x0071E627 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E6CC | `Clock_Screen"` | Known | Screen layout |
| 0x0071E771 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E81B | `Clock_Screen"` | Known | Screen layout |
| 0x0071E8C0 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E965 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EA0A | `Clock_Screen"` | Known | Screen layout |
| 0x0071EAAF | `Clock_Screen"` | Known | Screen layout |
| 0x0071EB54 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EBF9 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EC9E | `Clock_Screen"` | Known | Screen layout |
| 0x0071ED47 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EDEC | `Clock_Screen"` | Known | Screen layout |
| 0x0071EE91 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EF36 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EFDB | `Clock_Screen"` | Known | Screen layout |
| 0x0071F080 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F125 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F1CA | `Clock_Screen"` | Known | Screen layout |
| 0x0071F26F | `Clock_Screen"` | Known | Screen layout |
| 0x0071F314 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F3B9 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F45E | `Clock_Screen"` | Known | Screen layout |
| 0x0071F503 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F5A8 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F64D | `Clock_Screen"` | Known | Screen layout |
| 0x0071F6F2 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F797 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F83C | `Clock_Screen"` | Known | Screen layout |
| 0x0071F8E1 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F986 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FA2B | `Clock_Screen"` | Known | Screen layout |
| 0x0071FAD0 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FB75 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FC1A | `Clock_Screen"` | Known | Screen layout |
| 0x0071FCBF | `Clock_Screen"` | Known | Screen layout |
| 0x0071FD64 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FE09 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FEAE | `Clock_Screen"` | Known | Screen layout |
| 0x0071FF53 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FFF8 | `Clock_Screen"` | Known | Screen layout |
| 0x0072009D | `Clock_Screen"` | Known | Screen layout |
| 0x00720142 | `Clock_Screen"` | Known | Screen layout |
| 0x007201E7 | `Clock_Screen"` | Known | Screen layout |
| 0x0072028C | `Clock_Screen"` | Known | Screen layout |
| 0x00720337 | `Clock_Screen"` | Known | Screen layout |
| 0x007203DC | `Clock_Screen"` | Known | Screen layout |
| 0x00720481 | `Clock_Screen"` | Known | Screen layout |
| 0x00720526 | `Clock_Screen"` | Known | Screen layout |
| 0x007205CB | `Clock_Screen"` | Known | Screen layout |
| 0x00720670 | `Clock_Screen"` | Known | Screen layout |
| 0x00720715 | `Clock_Screen"` | Known | Screen layout |
| 0x007207BA | `Clock_Screen"` | Known | Screen layout |
| 0x0072085F | `Clock_Screen"` | Known | Screen layout |
| 0x00720904 | `Clock_Screen"` | Known | Screen layout |
| 0x007209A9 | `Clock_Screen"` | Known | Screen layout |
| 0x00720A4E | `Clock_Screen"` | Known | Screen layout |
| 0x00720AF3 | `Clock_Screen"` | Known | Screen layout |
| 0x00720B98 | `Clock_Screen"` | Known | Screen layout |
| 0x00720C3D | `Clock_Screen"` | Known | Screen layout |
| 0x00720CE2 | `Clock_Screen"` | Known | Screen layout |
| 0x00720D87 | `Clock_Screen"` | Known | Screen layout |
| 0x00720E2C | `Clock_Screen"` | Known | Screen layout |
| 0x00720ED1 | `Clock_Screen"` | Known | Screen layout |
| 0x00720F76 | `Clock_Screen"` | Known | Screen layout |
| 0x0072101B | `Clock_Screen"` | Known | Screen layout |
| 0x007210C0 | `Clock_Screen"` | Known | Screen layout |
| 0x00721165 | `Clock_Screen"` | Known | Screen layout |
| 0x0072120A | `Clock_Screen"` | Known | Screen layout |
| 0x007212AF | `Clock_Screen"` | Known | Screen layout |
| 0x00721354 | `Clock_Screen"` | Known | Screen layout |
| 0x007213F9 | `Clock_Screen"` | Known | Screen layout |
| 0x0072149E | `Clock_Screen"` | Known | Screen layout |
| 0x00721543 | `Clock_Screen"` | Known | Screen layout |
| 0x007215E8 | `Clock_Screen"` | Known | Screen layout |
| 0x0072168D | `Clock_Screen"` | Known | Screen layout |
| 0x00721732 | `Clock_Screen"` | Known | Screen layout |
| 0x007217D7 | `Clock_Screen"` | Known | Screen layout |
| 0x0072187C | `Clock_Screen"` | Known | Screen layout |
| 0x00721921 | `Clock_Screen"` | Known | Screen layout |
| 0x007219C6 | `Clock_Screen"` | Known | Screen layout |
| 0x00721A6B | `Clock_Screen"` | Known | Screen layout |
| 0x00721B10 | `Clock_Screen"` | Known | Screen layout |
| 0x00721BB5 | `Clock_Screen"` | Known | Screen layout |
| 0x00721C5A | `Clock_Screen"` | Known | Screen layout |
| 0x00721CFF | `Clock_Screen"` | Known | Screen layout |
| 0x00721DA4 | `Clock_Screen"` | Known | Screen layout |
| 0x00721E49 | `Clock_Screen"` | Known | Screen layout |
| 0x00721EEE | `Clock_Screen"` | Known | Screen layout |
| 0x00721F93 | `Clock_Screen"` | Known | Screen layout |
| 0x00722038 | `Clock_Screen"` | Known | Screen layout |
| 0x007220DD | `Clock_Screen"` | Known | Screen layout |
| 0x00722182 | `Clock_Screen"` | Known | Screen layout |
| 0x00722227 | `Clock_Screen"` | Known | Screen layout |
| 0x007222CC | `Clock_Screen"` | Known | Screen layout |
| 0x00722377 | `Clock_Screen"` | Known | Screen layout |
| 0x0072241C | `Clock_Screen"` | Known | Screen layout |
| 0x007224C1 | `Clock_Screen"` | Known | Screen layout |
| 0x00722566 | `Clock_Screen"` | Known | Screen layout |
| 0x0072260B | `Clock_Screen"` | Known | Screen layout |
| 0x007226B7 | `Clock_Screen"` | Known | Screen layout |
| 0x0072275C | `Clock_Screen"` | Known | Screen layout |
| 0x00722801 | `Clock_Screen"` | Known | Screen layout |
| 0x007228A6 | `Clock_Screen"` | Known | Screen layout |
| 0x0072294B | `Clock_Screen"` | Known | Screen layout |
| 0x007229F0 | `Clock_Screen"` | Known | Screen layout |
| 0x00722A95 | `Clock_Screen"` | Known | Screen layout |
| 0x00722B3A | `Clock_Screen"` | Known | Screen layout |
| 0x00722BDF | `Clock_Screen"` | Known | Screen layout |
| 0x00722C84 | `Clock_Screen"` | Known | Screen layout |
| 0x00722D29 | `Clock_Screen"` | Known | Screen layout |
| 0x00722DCE | `Clock_Screen"` | Known | Screen layout |
| 0x00722E73 | `Clock_Screen"` | Known | Screen layout |
| 0x00722F18 | `Clock_Screen"` | Known | Screen layout |
| 0x00722FBD | `Clock_Screen"` | Known | Screen layout |
| 0x00723062 | `Clock_Screen"` | Known | Screen layout |
| 0x00723107 | `Clock_Screen"` | Known | Screen layout |
| 0x007231AA | `Settings_DateTime_SetDate_Screen(` | Known | Screen layout |
| 0x007231CE | `Settings_DateTime_SetDate_Screen_Default"` | Known | Screen layout |
| 0x00723247 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x007232AD | `Settings_DateTime_SetTime_Screen(` | Known | Screen layout |
| 0x007232D1 | `Settings_DateTime_SetTime_Screen_Default"` | Known | Screen layout |
| 0x0072334A | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x007233B5 | `Settings_DateTime_SetTimeZone_Screen,` | Known | Screen layout |
| 0x007233DD | `Settings_DateTime_SetTimeZone_Screen_Default"` | Known | Screen layout |
| 0x0072345A | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x00723513 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x007235C3 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x00723BCA | `Search_Main_Screen` | Known | Screen layout |
| 0x00723BE0 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x00724102 | `Extras_Screen` | Known | Screen layout |
| 0x00724113 | `Extras_Screen_WorldClock"` | Known | Screen layout |
| 0x00724190 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x007241F2 | `Clock_Screen` | Known | Screen layout |
| 0x00724202 | `Clock_Screen_Default` | Known | Screen layout |
| 0x00724289 | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x007242EF | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x00724305 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x00724370 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x007243D2 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x007243EA | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x00724457 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x007244BB | `AddressViewer_Main_Screen!` | Known | Screen layout |
| 0x007244D8 | `AddressViewer_Main_Screen_Default"` | Known | Screen layout |
| 0x0072454A | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x007245B1 | `Games_Menu_Screen` | Known | Screen layout |
| 0x007245C6 | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x00724630 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x007246F7 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x00724793 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x00724864 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x00724924 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x00724988 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007249A7 | `VoiceMemos_Menu_Screen_Default"` | Known | Screen layout |
| 0x00724A2A | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x00724A90 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x00724AA8 | `Speakers_Main_Screen_Default"` | Known | Screen layout |
| 0x00724B29 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x00724B8D | `Radio_Screen` | Known | Screen layout |
| 0x00724B9D | `Radio_Screen_Default"` | Known | Screen layout |
| 0x00724C16 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x00724C77 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00724D13 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x00724DD6 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x00724E95 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x00724F52 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x007253A2 | `Radio_Screen` | Known | Screen layout |
| 0x007253B2 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x0072542B | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x0072560F | `Search_Main_Screen` | Known | Screen layout |
| 0x00725625 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x00725750 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x007257B3 | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x00725AF4 | `Video_Settings_Screen` | Known | Screen layout |
| 0x00725B0D | `Video_Settings_Screen_Default"` | Known | Screen layout |
| 0x00725C16 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x00725EDB | `SettingsMenus_MainMenu_Screen%` | Known | Screen layout |
| 0x00725FE9 | `SettingsMenus_MusicMenu_Screen&` | Known | Screen layout |
| 0x00726292 | `SettingsMenus_Brightness_Screen&` | Known | Screen layout |
| 0x007263A7 | `SettingsMenus_AudiobookSettings_Screen.` | Known | Screen layout |
| 0x007264DD | `SettingsMenus_RadioRegions_Screen ` | Known | Screen layout |
| 0x007265F2 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x0072685E | `Settings_DateTime_Screen ` | Known | Screen layout |
| 0x0072687A | `Settings_DateTime_Screen_Default"` | Known | Screen layout |
| 0x00726A06 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x00726B0B | `Settings_Legal_Screen` | Known | Screen layout |
| 0x00726B24 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x00726C15 | `SettingsMenus_ResetAllSettings_Screen4` | Known | Screen layout |
| 0x007273E6 | `Stopwatch_Screen` | Known | Screen layout |
| 0x007273FA | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x00727461 | `Stopwatch_Screen` | Known | Screen layout |
| 0x00727475 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x0072751E | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x00727541 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007275DA | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x007275FD | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007277B0 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x0072781E | `Speakers_ToneControl_Screen#` | Known | Screen layout |
| 0x0072783D | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x0073A929 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073A9AC | `LockediPod_Screen` | Known | Screen layout |
| 0x0073AA34 | `Lock_Screen` | Known | Screen layout |
| 0x0073AA43 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073ABDE | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x0073ACB0 | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x0073AD1A | `AddressViewer_ContactDetails_Screen+` | Known | Screen layout |
| 0x0073AD41 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x0073ADBC | `Extras_Screen` | Known | Screen layout |
| 0x0073AE07 | `Extras_Screen` | Known | Screen layout |
| 0x0073AEEE | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x0073AF4C | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0073AF69 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x0073AFD7 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0073AFF0 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0073B067 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0073B084 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x0073B0EF | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x0073B10C | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0073B173 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x0073B1DA | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x0073B238 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0073B255 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x0073B2C3 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0073B2DC | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0073B353 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0073B370 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x0073B3DB | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x0073B3F8 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0073B45F | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x0073B4FF | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x0073B588 | `Alarms_Set_Alarm_Frequency_Screen)` | Known | Screen layout |
| 0x0073B5AD | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x0073B61E | `Alarms_Set_Alarm_Label_Screen%` | Known | Screen layout |
| 0x0073B63F | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x0073B6AC | `Alarms_Set_Alarm_Sound_Screen%` | Known | Screen layout |
| 0x0073B6CD | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x0073B739 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x0073B9B4 | `Alarms_Set_Alarm_Playlist_Screen(` | Known | Screen layout |
| 0x0073B9D8 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x0073BA48 | `Alarms_Set_Alarm_Tones_Screen%` | Known | Screen layout |
| 0x0073BA69 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x0073BD7C | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x0073BD97 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x0073BEE8 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0073BEFF | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x0073BF80 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0073BF97 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0073C06D | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0073C086 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0073C10B | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x0073C17C | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0073C271 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0073C28A | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0073C30F | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x0073C380 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0073C440 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x0073C454 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x0073C583 | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x0073C5E6 | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x0073C63D | `Clock_Screen_Default` | Known | Screen layout |
| 0x0073C6CE | `Clock_Region_Screen` | Known | Screen layout |
| 0x0073C6E5 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x0073C75E | `Clock_Screen_Default` | Known | Screen layout |
| 0x0073C7B5 | `Clock_Screen_Default` | Known | Screen layout |
| 0x0073C846 | `Clock_Region_Screen` | Known | Screen layout |
| 0x0073C85D | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x0073C9E8 | `CoverFlow_Screen_DefaultNoAlbumInfo&` | Known | Screen layout |
| 0x0073CAD6 | `CoverFlow_Screen_DefaultNoAlbumInfo#` | Known | Screen layout |
| 0x0073CB4B | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073CE41 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073CFF1 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073D11F | `CoverFlow_Screen_Default&` | Known | Screen layout |
| 0x0073D1F5 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073D38A | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073D5EF | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x0073D64C | `Game_Screen` | Known | Screen layout |
| 0x0073D65B | `Game_Screen_Default` | Known | Screen layout |
| 0x0073D6FD | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073D75F | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073D7C2 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073D825 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073D881 | `Game_Running_Screen` | Known | Screen layout |
| 0x0073D8E1 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073D943 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073D9A6 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073DA09 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073DA65 | `Game_Running_Screen` | Known | Screen layout |
| 0x0073DAC5 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073DB27 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073DB8A | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073DBED | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073DC49 | `Game_Running_Screen` | Known | Screen layout |
| 0x0073DCA9 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073DD0B | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073DD6E | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073DDD1 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073DE2D | `Game_Running_Screen` | Known | Screen layout |
| 0x0073DE8D | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073DEEF | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073DF52 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073DFB5 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073E011 | `Game_Running_Screen` | Known | Screen layout |
| 0x0073E257 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073E2B9 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073E31C | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073E37F | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073E3DB | `Game_Running_Screen` | Known | Screen layout |
| 0x0073E492 | `Extras_Screen` | Known | Screen layout |
| 0x0073E4A3 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073E501 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0073E69E | `Extras_Screen` | Known | Screen layout |
| 0x0073E6AF | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073E70D | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0073E8AA | `Extras_Screen` | Known | Screen layout |
| 0x0073E8BB | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073E919 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0073EAB6 | `Extras_Screen` | Known | Screen layout |
| 0x0073EAC7 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073EB25 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0073ECC7 | `Lock_Screen` | Known | Screen layout |
| 0x0073ECD6 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0073ED38 | `Extras_Screen` | Known | Screen layout |
| 0x0073ED49 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0073EDA8 | `LockediPod_Screen` | Known | Screen layout |
| 0x0073EE22 | `MainMenus_Main_Screen)` | Known | Screen layout |
| 0x0073EFF3 | `Lock_Screen` | Known | Screen layout |
| 0x0073F002 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0073F064 | `Extras_Screen` | Known | Screen layout |
| 0x0073F075 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0073F0D4 | `LockediPod_Screen` | Known | Screen layout |
| 0x0073F14E | `MainMenus_Main_Screen%` | Known | Screen layout |
| 0x0073F1B5 | `LockediPod_Screen` | Known | Screen layout |
| 0x0073F1CA | `LockediPod_Screen_Default,` | Known | Screen layout |
| 0x0073F319 | `Lock_Screen` | Known | Screen layout |
| 0x0073F328 | `Lock_Screen_Default_Layout"` | Known | Screen layout |
| 0x0073F391 | `Lock_Screen` | Known | Screen layout |
| 0x0073F3A0 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0073F402 | `Extras_Screen` | Known | Screen layout |
| 0x0073F413 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0073F472 | `LockediPod_Screen` | Known | Screen layout |
| 0x0073F4EC | `MainMenus_Main_Screen,` | Known | Screen layout |
| 0x0073F648 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0073F6AE | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0073F712 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073F7A1 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0073F80E | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0073F87B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0073F8E8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073F950 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0073F9B6 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0073FA1A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073FAA9 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0073FB16 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0073FB83 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0073FBF0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073FC58 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0073FCBE | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0073FD22 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073FDB1 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0073FE1E | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0073FE8B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0073FEF8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073FF60 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0073FFC6 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0074002A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007400B9 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x00740126 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x00740193 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x00740200 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00740268 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x007402CE | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x00740332 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007403C1 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0074042E | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0074049B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x00740508 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00740561 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x007405CA | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00740631 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007406CC | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00740735 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x0074079E | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00740805 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007408A0 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00740909 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x00740972 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x007409D9 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00740A74 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00740B60 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00740B7C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00740BEA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00740C07 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00740C72 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00740C92 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00740D09 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00740D25 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00740D95 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00740DB4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00740E20 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00740E34 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00740EAD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00740F21 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00740F91 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00740FF8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00741060 | `NoContent_Screen` | Known | Screen layout |
| 0x00741074 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007410D8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074113F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00741159 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007411C7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00741239 | `NoContent_Screen` | Known | Screen layout |
| 0x0074124D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007412B7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00741320 | `No_Photos_Screen` | Known | Screen layout |
| 0x00741334 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074139A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00741408 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00741475 | `NoContent_Screen` | Known | Screen layout |
| 0x00741489 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007414F1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074155B | `NoContent_Screen` | Known | Screen layout |
| 0x0074156F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007415D6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00741640 | `NoContent_Screen` | Known | Screen layout |
| 0x00741654 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007416C1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00741733 | `NoContent_Screen` | Known | Screen layout |
| 0x00741747 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007417AF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00741818 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00741833 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00741899 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007418B5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00741994 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007419AD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00741A0E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00741A22 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00741A7C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00741A98 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00741AFF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00741B16 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00741C87 | `Radio_Screen` | Known | Screen layout |
| 0x00741C97 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00741CF8 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00741D7B | `LockediPod_Screen` | Known | Screen layout |
| 0x00741E03 | `Lock_Screen` | Known | Screen layout |
| 0x00741E12 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00741E75 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00741ED7 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00741EF3 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00741F65 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00741F84 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00741FEC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00742006 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074206E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074208B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007420F7 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00742161 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074217B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007421EB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074225E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007422CF | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074233E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007423AA | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007423C5 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074243A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007424A1 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00742503 | `Photos_Screen` | Known | Screen layout |
| 0x00742567 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00742585 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007425F7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00742614 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074267A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00742695 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007426FE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074271B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00742792 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007427B6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00742824 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074283F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007428FC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00742918 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00742986 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007429A3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00742A0E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00742A2E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00742AA5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00742AC1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00742B31 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00742B50 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00742BBC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00742BD0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00742C49 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00742CBD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00742D2D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00742D94 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00742DFC | `NoContent_Screen` | Known | Screen layout |
| 0x00742E10 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00742E74 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00742EDB | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00742EF5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00742F63 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00742FD5 | `NoContent_Screen` | Known | Screen layout |
| 0x00742FE9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00743053 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007430BC | `No_Photos_Screen` | Known | Screen layout |
| 0x007430D0 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00743136 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007431A4 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00743211 | `NoContent_Screen` | Known | Screen layout |
| 0x00743225 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074328D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007432F7 | `NoContent_Screen` | Known | Screen layout |
| 0x0074330B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00743372 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007433DC | `NoContent_Screen` | Known | Screen layout |
| 0x007433F0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074345D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007434CF | `NoContent_Screen` | Known | Screen layout |
| 0x007434E3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074354B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007435B4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007435CF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00743635 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00743651 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00743730 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00743749 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007437AA | `FirstBoot_Screen` | Known | Screen layout |
| 0x007437BE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00743818 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00743834 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074389B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007438B2 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00743A23 | `Radio_Screen` | Known | Screen layout |
| 0x00743A33 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00743A94 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00743B17 | `LockediPod_Screen` | Known | Screen layout |
| 0x00743B9F | `Lock_Screen` | Known | Screen layout |
| 0x00743BAE | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00743C11 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00743C73 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00743C8F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00743D01 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00743D20 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00743D88 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00743DA2 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00743E0A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00743E27 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00743E93 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00743EFD | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00743F17 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00743F87 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00743FFA | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074406B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007440DA | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00744146 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00744161 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007441D6 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074423D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074429F | `Photos_Screen` | Known | Screen layout |
| 0x00744303 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00744321 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00744393 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007443B0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00744416 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00744431 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074449A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007444B7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074452E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00744552 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007445C0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007445DB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00744698 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007446B4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00744722 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074473F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007447AA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007447CA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00744841 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074485D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007448CD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007448EC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00744958 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074496C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007449E5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00744A59 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00744AC9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00744B30 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00744B98 | `NoContent_Screen` | Known | Screen layout |
| 0x00744BAC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00744C10 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00744C77 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00744C91 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00744CFF | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00744D71 | `NoContent_Screen` | Known | Screen layout |
| 0x00744D85 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00744DEF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00744E58 | `No_Photos_Screen` | Known | Screen layout |
| 0x00744E6C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00744ED2 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00744F40 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00744FAD | `NoContent_Screen` | Known | Screen layout |
| 0x00744FC1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00745029 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00745093 | `NoContent_Screen` | Known | Screen layout |
| 0x007450A7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074510E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00745178 | `NoContent_Screen` | Known | Screen layout |
| 0x0074518C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007451F9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074526B | `NoContent_Screen` | Known | Screen layout |
| 0x0074527F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007452E7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00745350 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074536B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007453D1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007453ED | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007454CC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007454E5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00745546 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074555A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007455B4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007455D0 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00745637 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074564E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007457BF | `Radio_Screen` | Known | Screen layout |
| 0x007457CF | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00745830 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007458B3 | `LockediPod_Screen` | Known | Screen layout |
| 0x0074593B | `Lock_Screen` | Known | Screen layout |
| 0x0074594A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007459AD | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00745A0F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00745A2B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00745A9D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00745ABC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00745B24 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00745B3E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00745BA6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00745BC3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00745C2F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00745C99 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00745CB3 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00745D23 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00745D96 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00745E07 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00745E76 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00745EE2 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00745EFD | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00745F72 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00745FD9 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074603B | `Photos_Screen` | Known | Screen layout |
| 0x0074609F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007460BD | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074612F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074614C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007461B2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007461CD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00746236 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00746253 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007462CA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007462EE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074635C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00746377 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00746434 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00746450 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007464BE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007464DB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00746546 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00746566 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007465DD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007465F9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00746669 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00746688 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007466F4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00746708 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00746781 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007467F5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00746865 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007468CC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00746934 | `NoContent_Screen` | Known | Screen layout |
| 0x00746948 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007469AC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00746A13 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00746A2D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00746A9B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00746B0D | `NoContent_Screen` | Known | Screen layout |
| 0x00746B21 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00746B8B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00746BF4 | `No_Photos_Screen` | Known | Screen layout |
| 0x00746C08 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00746C6E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00746CDC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00746D49 | `NoContent_Screen` | Known | Screen layout |
| 0x00746D5D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00746DC5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00746E2F | `NoContent_Screen` | Known | Screen layout |
| 0x00746E43 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00746EAA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00746F14 | `NoContent_Screen` | Known | Screen layout |
| 0x00746F28 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00746F95 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00747007 | `NoContent_Screen` | Known | Screen layout |
| 0x0074701B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00747083 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007470EC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00747107 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074716D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00747189 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00747268 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00747281 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007472E2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007472F6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00747350 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074736C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007473D3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007473EA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074755B | `Radio_Screen` | Known | Screen layout |
| 0x0074756B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007475CC | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074764F | `LockediPod_Screen` | Known | Screen layout |
| 0x007476D7 | `Lock_Screen` | Known | Screen layout |
| 0x007476E6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00747749 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007477AB | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007477C7 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00747839 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00747858 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007478C0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007478DA | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00747942 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074795F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007479CB | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00747A35 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00747A4F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00747ABF | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00747B32 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00747BA3 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00747C12 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00747C7E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00747C99 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00747D0E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00747D75 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00747DD7 | `Photos_Screen` | Known | Screen layout |
| 0x00747E3B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00747E59 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00747ECB | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00747EE8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00747F4E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00747F69 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00747FD2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00747FEF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00748066 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074808A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007480F8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00748113 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007481D0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007481EC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074825A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00748277 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007482E2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00748302 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00748379 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00748395 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00748405 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00748424 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00748490 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007484A4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074851D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00748591 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00748601 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00748668 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007486D0 | `NoContent_Screen` | Known | Screen layout |
| 0x007486E4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00748748 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007487AF | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007487C9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00748837 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007488A9 | `NoContent_Screen` | Known | Screen layout |
| 0x007488BD | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00748927 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00748990 | `No_Photos_Screen` | Known | Screen layout |
| 0x007489A4 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00748A0A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00748A78 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00748AE5 | `NoContent_Screen` | Known | Screen layout |
| 0x00748AF9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00748B61 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00748BCB | `NoContent_Screen` | Known | Screen layout |
| 0x00748BDF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00748C46 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00748CB0 | `NoContent_Screen` | Known | Screen layout |
| 0x00748CC4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00748D31 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00748DA3 | `NoContent_Screen` | Known | Screen layout |
| 0x00748DB7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00748E1F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00748E88 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00748EA3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00748F09 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00748F25 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00749004 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074901D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074907E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00749092 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007490EC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00749108 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074916F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00749186 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007492F7 | `Radio_Screen` | Known | Screen layout |
| 0x00749307 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00749368 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007493EB | `LockediPod_Screen` | Known | Screen layout |
| 0x00749473 | `Lock_Screen` | Known | Screen layout |
| 0x00749482 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007494E5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00749547 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00749563 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007495D5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007495F4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074965C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00749676 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007496DE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007496FB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00749767 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007497D1 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007497EB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074985B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007498CE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074993F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007499AE | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00749A1A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00749A35 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00749AAA | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00749B11 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00749B73 | `Photos_Screen` | Known | Screen layout |
| 0x00749BD7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00749BF5 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00749C67 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00749C84 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00749CEA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00749D05 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00749D6E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00749D8B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00749E02 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00749E26 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00749E94 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00749EAF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00749F6C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00749F88 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00749FF6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074A013 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074A07E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074A09E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074A115 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074A131 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074A1A1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074A1C0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074A22C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074A240 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074A2B9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074A32D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074A39D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074A404 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074A46C | `NoContent_Screen` | Known | Screen layout |
| 0x0074A480 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074A4E4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074A54B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074A565 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074A5D3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074A645 | `NoContent_Screen` | Known | Screen layout |
| 0x0074A659 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074A6C3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074A72C | `No_Photos_Screen` | Known | Screen layout |
| 0x0074A740 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074A7A6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074A814 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074A881 | `NoContent_Screen` | Known | Screen layout |
| 0x0074A895 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074A8FD | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074A967 | `NoContent_Screen` | Known | Screen layout |
| 0x0074A97B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074A9E2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074AA4C | `NoContent_Screen` | Known | Screen layout |
| 0x0074AA60 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074AACD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074AB3F | `NoContent_Screen` | Known | Screen layout |
| 0x0074AB53 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074ABBB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074AC24 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074AC3F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074ACA5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074ACC1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074ADA0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074ADB9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074AE1A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074AE2E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074AE88 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074AEA4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074AF0B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074AF22 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074B093 | `Radio_Screen` | Known | Screen layout |
| 0x0074B0A3 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074B104 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074B187 | `LockediPod_Screen` | Known | Screen layout |
| 0x0074B20F | `Lock_Screen` | Known | Screen layout |
| 0x0074B21E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074B281 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074B2E3 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074B2FF | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0074B371 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074B390 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074B3F8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074B412 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074B47A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074B497 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074B503 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074B56D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074B587 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074B5F7 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074B66A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074B6DB | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074B74A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074B7B6 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074B7D1 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074B846 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074B8AD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074B90F | `Photos_Screen` | Known | Screen layout |
| 0x0074B973 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074B991 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074BA03 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074BA20 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074BA86 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074BAA1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074BB0A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074BB27 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074BB9E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074BBC2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074BC30 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074BC4B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074BD08 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074BD24 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074BD92 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074BDAF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074BE1A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074BE3A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074BEB1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074BECD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074BF3D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074BF5C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074BFC8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074BFDC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074C055 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074C0C9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074C139 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074C1A0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074C208 | `NoContent_Screen` | Known | Screen layout |
| 0x0074C21C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074C280 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074C2E7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074C301 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074C36F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074C3E1 | `NoContent_Screen` | Known | Screen layout |
| 0x0074C3F5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074C45F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074C4C8 | `No_Photos_Screen` | Known | Screen layout |
| 0x0074C4DC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074C542 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074C5B0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074C61D | `NoContent_Screen` | Known | Screen layout |
| 0x0074C631 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074C699 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074C703 | `NoContent_Screen` | Known | Screen layout |
| 0x0074C717 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074C77E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074C7E8 | `NoContent_Screen` | Known | Screen layout |
| 0x0074C7FC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074C869 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074C8DB | `NoContent_Screen` | Known | Screen layout |
| 0x0074C8EF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074C957 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074C9C0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074C9DB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074CA41 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074CA5D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074CB3C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074CB55 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074CBB6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074CBCA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074CC24 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074CC40 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074CCA7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074CCBE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074CE2F | `Radio_Screen` | Known | Screen layout |
| 0x0074CE3F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074CEA0 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074CF23 | `LockediPod_Screen` | Known | Screen layout |
| 0x0074CFAB | `Lock_Screen` | Known | Screen layout |
| 0x0074CFBA | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074D01D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074D07F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074D09B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0074D10D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074D12C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074D194 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074D1AE | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074D216 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074D233 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074D29F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074D309 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074D323 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074D393 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074D406 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074D477 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074D4E6 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074D552 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074D56D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074D5E2 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074D649 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074D6AB | `Photos_Screen` | Known | Screen layout |
| 0x0074D70F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074D72D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074D79F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074D7BC | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074D822 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074D83D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074D8A6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074D8C3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074D93A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074D95E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074D9CC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074D9E7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074DAA4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074DAC0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074DB2E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074DB4B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074DBB6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074DBD6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074DC4D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074DC69 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074DCD9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074DCF8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074DD64 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074DD78 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074DDF1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074DE65 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074DED5 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074DF3C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074DFA4 | `NoContent_Screen` | Known | Screen layout |
| 0x0074DFB8 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074E01C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074E083 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074E09D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074E10B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074E17D | `NoContent_Screen` | Known | Screen layout |
| 0x0074E191 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074E1FB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074E264 | `No_Photos_Screen` | Known | Screen layout |
| 0x0074E278 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074E2DE | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074E34C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074E3B9 | `NoContent_Screen` | Known | Screen layout |
| 0x0074E3CD | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074E435 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074E49F | `NoContent_Screen` | Known | Screen layout |
| 0x0074E4B3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074E51A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074E584 | `NoContent_Screen` | Known | Screen layout |
| 0x0074E598 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074E605 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074E677 | `NoContent_Screen` | Known | Screen layout |
| 0x0074E68B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074E6F3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074E75C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074E777 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074E7DD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074E7F9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074E8D8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074E8F1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074E952 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074E966 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074E9C0 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074E9DC | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074EA43 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074EA5A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074EBCB | `Radio_Screen` | Known | Screen layout |
| 0x0074EBDB | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074EC3C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074ECBF | `LockediPod_Screen` | Known | Screen layout |
| 0x0074ED47 | `Lock_Screen` | Known | Screen layout |
| 0x0074ED56 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074EDB9 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074EE1B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074EE37 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0074EEA9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074EEC8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074EF30 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074EF4A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074EFB2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074EFCF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074F03B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074F0A5 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074F0BF | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074F12F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074F1A2 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074F213 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074F282 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074F2EE | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074F309 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074F37E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074F3E5 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074F447 | `Photos_Screen` | Known | Screen layout |
| 0x0074F4AB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074F4C9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074F53B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074F558 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074F5BE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074F5D9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074F642 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074F65F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074F6D6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074F6FA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074F768 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074F783 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074F840 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074F85C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074F8CA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074F8E7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074F952 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074F972 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074F9E9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074FA05 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074FA75 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074FA94 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074FB00 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074FB14 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074FB8D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074FC01 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074FC71 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074FCD8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074FD40 | `NoContent_Screen` | Known | Screen layout |
| 0x0074FD54 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074FDB8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074FE1F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074FE39 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074FEA7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074FF19 | `NoContent_Screen` | Known | Screen layout |
| 0x0074FF2D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074FF97 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00750000 | `No_Photos_Screen` | Known | Screen layout |
| 0x00750014 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075007A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007500E8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00750155 | `NoContent_Screen` | Known | Screen layout |
| 0x00750169 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007501D1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075023B | `NoContent_Screen` | Known | Screen layout |
| 0x0075024F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007502B6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00750320 | `NoContent_Screen` | Known | Screen layout |
| 0x00750334 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007503A1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00750413 | `NoContent_Screen` | Known | Screen layout |
| 0x00750427 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075048F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007504F8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00750513 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00750579 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00750595 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00750674 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075068D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007506EE | `FirstBoot_Screen` | Known | Screen layout |
| 0x00750702 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075075C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00750778 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007507DF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007507F6 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00750967 | `Radio_Screen` | Known | Screen layout |
| 0x00750977 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007509D8 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00750A5B | `LockediPod_Screen` | Known | Screen layout |
| 0x00750AE3 | `Lock_Screen` | Known | Screen layout |
| 0x00750AF2 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00750B55 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00750BB7 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00750BD3 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00750C45 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00750C64 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00750CCC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00750CE6 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00750D4E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00750D6B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00750DD7 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00750E41 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00750E5B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00750ECB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00750F3E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00750FAF | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075101E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075108A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007510A5 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075111A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00751181 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007511E3 | `Photos_Screen` | Known | Screen layout |
| 0x00751247 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00751265 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007512D7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007512F4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075135A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00751375 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007513DE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007513FB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00751472 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00751496 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00751504 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075151F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007515DC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007515F8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00751666 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00751683 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007516EE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075170E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00751785 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007517A1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00751811 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00751830 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075189C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007518B0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00751929 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075199D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00751A0D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00751A74 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00751ADC | `NoContent_Screen` | Known | Screen layout |
| 0x00751AF0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00751B54 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00751BBB | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00751BD5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00751C43 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00751CB5 | `NoContent_Screen` | Known | Screen layout |
| 0x00751CC9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00751D33 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00751D9C | `No_Photos_Screen` | Known | Screen layout |
| 0x00751DB0 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00751E16 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00751E84 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00751EF1 | `NoContent_Screen` | Known | Screen layout |
| 0x00751F05 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00751F6D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00751FD7 | `NoContent_Screen` | Known | Screen layout |
| 0x00751FEB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00752052 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007520BC | `NoContent_Screen` | Known | Screen layout |
| 0x007520D0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075213D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007521AF | `NoContent_Screen` | Known | Screen layout |
| 0x007521C3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075222B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00752294 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007522AF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00752315 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00752331 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00752410 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00752429 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075248A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075249E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007524F8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00752514 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075257B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00752592 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00752703 | `Radio_Screen` | Known | Screen layout |
| 0x00752713 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00752774 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007527F7 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075287F | `Lock_Screen` | Known | Screen layout |
| 0x0075288E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007528F1 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00752953 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075296F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007529E1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00752A00 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00752A68 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00752A82 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00752AEA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00752B07 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00752B73 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00752BDD | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00752BF7 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00752C67 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00752CDA | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00752D4B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00752DBA | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00752E26 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00752E41 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00752EB6 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00752F1D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00752F7F | `Photos_Screen` | Known | Screen layout |
| 0x00752FE3 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00753001 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00753073 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00753090 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007530F6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00753111 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075317A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00753197 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075320E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00753232 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007532A0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007532BB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00753378 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00753394 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00753402 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075341F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075348A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007534AA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00753521 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075353D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007535AD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007535CC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00753638 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075364C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007536C5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00753739 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007537A9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00753810 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00753878 | `NoContent_Screen` | Known | Screen layout |
| 0x0075388C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007538F0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00753957 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00753971 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007539DF | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00753A51 | `NoContent_Screen` | Known | Screen layout |
| 0x00753A65 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00753ACF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00753B38 | `No_Photos_Screen` | Known | Screen layout |
| 0x00753B4C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00753BB2 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00753C20 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00753C8D | `NoContent_Screen` | Known | Screen layout |
| 0x00753CA1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00753D09 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00753D73 | `NoContent_Screen` | Known | Screen layout |
| 0x00753D87 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00753DEE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00753E58 | `NoContent_Screen` | Known | Screen layout |
| 0x00753E6C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00753ED9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00753F4B | `NoContent_Screen` | Known | Screen layout |
| 0x00753F5F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00753FC7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00754030 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075404B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007540B1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007540CD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007541AC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007541C5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00754226 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075423A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00754294 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007542B0 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00754317 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075432E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075449F | `Radio_Screen` | Known | Screen layout |
| 0x007544AF | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00754510 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00754593 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075461B | `Lock_Screen` | Known | Screen layout |
| 0x0075462A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075468D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007546EF | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075470B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075477D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075479C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00754804 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075481E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00754886 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007548A3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075490F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00754979 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00754993 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00754A03 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00754A76 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00754AE7 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00754B56 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00754BC2 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00754BDD | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00754C52 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00754CB9 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00754D1B | `Photos_Screen` | Known | Screen layout |
| 0x00754D7F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00754D9D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00754E0F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00754E2C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00754E92 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00754EAD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00754F16 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00754F33 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00754FAA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00754FCE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075503C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00755057 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00755114 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00755130 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075519E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007551BB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00755226 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00755246 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007552BD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007552D9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00755349 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00755368 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007553D4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007553E8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00755461 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007554D5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00755545 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007555AC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00755614 | `NoContent_Screen` | Known | Screen layout |
| 0x00755628 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075568C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007556F3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075570D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075577B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007557ED | `NoContent_Screen` | Known | Screen layout |
| 0x00755801 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075586B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007558D4 | `No_Photos_Screen` | Known | Screen layout |
| 0x007558E8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075594E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007559BC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00755A29 | `NoContent_Screen` | Known | Screen layout |
| 0x00755A3D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00755AA5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00755B0F | `NoContent_Screen` | Known | Screen layout |
| 0x00755B23 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00755B8A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00755BF4 | `NoContent_Screen` | Known | Screen layout |
| 0x00755C08 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00755C75 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00755CE7 | `NoContent_Screen` | Known | Screen layout |
| 0x00755CFB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00755D63 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00755DCC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00755DE7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00755E4D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00755E69 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00755F48 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00755F61 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00755FC2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00755FD6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00756030 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075604C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007560B3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007560CA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075623B | `Radio_Screen` | Known | Screen layout |
| 0x0075624B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007562AC | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075632F | `LockediPod_Screen` | Known | Screen layout |
| 0x007563B7 | `Lock_Screen` | Known | Screen layout |
| 0x007563C6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00756429 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075648B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007564A7 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00756519 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00756538 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007565A0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007565BA | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00756622 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075663F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007566AB | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00756715 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075672F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075679F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00756812 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00756883 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007568F2 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075695E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00756979 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007569EE | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00756A55 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00756AB7 | `Photos_Screen` | Known | Screen layout |
| 0x00756B1B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00756B39 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00756BAB | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00756BC8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00756C2E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00756C49 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00756CB2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00756CCF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00756D46 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00756D6A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00756DD8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00756DF3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00756EB0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00756ECC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00756F3A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00756F57 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00756FC2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00756FE2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00757059 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00757075 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007570E5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00757104 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00757170 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00757184 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007571FD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00757271 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007572E1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00757348 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007573B0 | `NoContent_Screen` | Known | Screen layout |
| 0x007573C4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00757428 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075748F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007574A9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00757517 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00757589 | `NoContent_Screen` | Known | Screen layout |
| 0x0075759D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00757607 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00757670 | `No_Photos_Screen` | Known | Screen layout |
| 0x00757684 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007576EA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00757758 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007577C5 | `NoContent_Screen` | Known | Screen layout |
| 0x007577D9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00757841 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007578AB | `NoContent_Screen` | Known | Screen layout |
| 0x007578BF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00757926 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00757990 | `NoContent_Screen` | Known | Screen layout |
| 0x007579A4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00757A11 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00757A83 | `NoContent_Screen` | Known | Screen layout |
| 0x00757A97 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00757AFF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00757B68 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00757B83 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00757BE9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00757C05 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00757CE4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00757CFD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00757D5E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00757D72 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00757DCC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00757DE8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00757E4F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00757E66 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00757FD7 | `Radio_Screen` | Known | Screen layout |
| 0x00757FE7 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00758048 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007580CB | `LockediPod_Screen` | Known | Screen layout |
| 0x00758153 | `Lock_Screen` | Known | Screen layout |
| 0x00758162 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007581C5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00758227 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00758243 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007582B5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007582D4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075833C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00758356 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007583BE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007583DB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00758447 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007584B1 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007584CB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075853B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007585AE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075861F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075868E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007586FA | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00758715 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075878A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007587F1 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00758853 | `Photos_Screen` | Known | Screen layout |
| 0x007588B7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007588D5 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00758947 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00758964 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007589CA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007589E5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00758A4E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00758A6B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00758AE2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00758B06 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00758B74 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00758B8F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00758C4C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00758C68 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00758CD6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00758CF3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00758D5E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00758D7E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00758DF5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00758E11 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00758E81 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00758EA0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00758F0C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00758F20 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00758F99 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075900D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075907D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007590E4 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075914C | `NoContent_Screen` | Known | Screen layout |
| 0x00759160 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007591C4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075922B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00759245 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007592B3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00759325 | `NoContent_Screen` | Known | Screen layout |
| 0x00759339 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007593A3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075940C | `No_Photos_Screen` | Known | Screen layout |
| 0x00759420 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00759486 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007594F4 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00759561 | `NoContent_Screen` | Known | Screen layout |
| 0x00759575 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007595DD | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00759647 | `NoContent_Screen` | Known | Screen layout |
| 0x0075965B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007596C2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075972C | `NoContent_Screen` | Known | Screen layout |
| 0x00759740 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007597AD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075981F | `NoContent_Screen` | Known | Screen layout |
| 0x00759833 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075989B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00759904 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075991F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00759985 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007599A1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00759A80 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00759A99 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00759AFA | `FirstBoot_Screen` | Known | Screen layout |
| 0x00759B0E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00759B68 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00759B84 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00759BEB | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00759C02 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00759D73 | `Radio_Screen` | Known | Screen layout |
| 0x00759D83 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00759DE4 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00759E67 | `LockediPod_Screen` | Known | Screen layout |
| 0x00759EEF | `Lock_Screen` | Known | Screen layout |
| 0x00759EFE | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00759F61 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00759FC3 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00759FDF | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075A051 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075A070 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075A0D8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075A0F2 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075A15A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075A177 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075A1E3 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075A24D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075A267 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075A2D7 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075A34A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075A3BB | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075A42A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075A496 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075A4B1 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075A526 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075A58D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075A5EF | `Photos_Screen` | Known | Screen layout |
| 0x0075A653 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075A671 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075A6E3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075A700 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075A766 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075A781 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075A7EA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075A807 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075A87E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075A8A2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075A910 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075A92B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075A9E8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075AA04 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075AA72 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075AA8F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075AAFA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075AB1A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075AB91 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075ABAD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075AC1D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075AC3C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075ACA8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075ACBC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075AD35 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075ADA9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075AE19 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075AE80 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075AEE8 | `NoContent_Screen` | Known | Screen layout |
| 0x0075AEFC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075AF60 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075AFC7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075AFE1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075B04F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075B0C1 | `NoContent_Screen` | Known | Screen layout |
| 0x0075B0D5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075B13F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075B1A8 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075B1BC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075B222 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075B290 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075B2FD | `NoContent_Screen` | Known | Screen layout |
| 0x0075B311 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075B379 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075B3E3 | `NoContent_Screen` | Known | Screen layout |
| 0x0075B3F7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075B45E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075B4C8 | `NoContent_Screen` | Known | Screen layout |
| 0x0075B4DC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075B549 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075B5BB | `NoContent_Screen` | Known | Screen layout |
| 0x0075B5CF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075B637 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075B6A0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075B6BB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075B721 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075B73D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075B81C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075B835 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075B896 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075B8AA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075B904 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075B920 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075B987 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075B99E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075BB0F | `Radio_Screen` | Known | Screen layout |
| 0x0075BB1F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075BB80 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075BC03 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075BC8B | `Lock_Screen` | Known | Screen layout |
| 0x0075BC9A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075BCFD | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075BD5F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075BD7B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075BDED | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075BE0C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075BE74 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075BE8E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075BEF6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075BF13 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075BF7F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075BFE9 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075C003 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075C073 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075C0E6 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075C157 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075C1C6 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075C232 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075C24D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075C2C2 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075C329 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075C38B | `Photos_Screen` | Known | Screen layout |
| 0x0075C3EF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075C40D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075C47F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075C49C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075C502 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075C51D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075C586 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075C5A3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075C61A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075C63E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075C6AC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075C6C7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075C784 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075C7A0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075C80E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075C82B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075C896 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075C8B6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075C92D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075C949 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075C9B9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075C9D8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075CA44 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075CA58 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075CAD1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075CB45 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075CBB5 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075CC1C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075CC84 | `NoContent_Screen` | Known | Screen layout |
| 0x0075CC98 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075CCFC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075CD63 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075CD7D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075CDEB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075CE5D | `NoContent_Screen` | Known | Screen layout |
| 0x0075CE71 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075CEDB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075CF44 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075CF58 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075CFBE | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075D02C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075D099 | `NoContent_Screen` | Known | Screen layout |
| 0x0075D0AD | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075D115 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075D17F | `NoContent_Screen` | Known | Screen layout |
| 0x0075D193 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075D1FA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075D264 | `NoContent_Screen` | Known | Screen layout |
| 0x0075D278 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075D2E5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075D357 | `NoContent_Screen` | Known | Screen layout |
| 0x0075D36B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075D3D3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075D43C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075D457 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075D4BD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075D4D9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075D5B8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075D5D1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075D632 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075D646 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075D6A0 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075D6BC | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075D723 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075D73A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075D8AB | `Radio_Screen` | Known | Screen layout |
| 0x0075D8BB | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075D91C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075D99F | `LockediPod_Screen` | Known | Screen layout |
| 0x0075DA27 | `Lock_Screen` | Known | Screen layout |
| 0x0075DA36 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075DA99 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075DAFB | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075DB17 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075DB89 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075DBA8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075DC10 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075DC2A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075DC92 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075DCAF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075DD1B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075DD85 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075DD9F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075DE0F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075DE82 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075DEF3 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075DF62 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075DFCE | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075DFE9 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075E05E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075E0C5 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075E127 | `Photos_Screen` | Known | Screen layout |
| 0x0075E18B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075E1A9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075E21B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075E238 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075E29E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075E2B9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075E322 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075E33F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075E3B6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075E3DA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075E448 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075E463 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075E520 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075E53C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075E5AA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075E5C7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075E632 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075E652 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075E6C9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075E6E5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075E755 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075E774 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075E7E0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075E7F4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075E86D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075E8E1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075E951 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075E9B8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075EA20 | `NoContent_Screen` | Known | Screen layout |
| 0x0075EA34 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075EA98 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075EAFF | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075EB19 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075EB87 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075EBF9 | `NoContent_Screen` | Known | Screen layout |
| 0x0075EC0D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075EC77 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075ECE0 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075ECF4 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075ED5A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075EDC8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075EE35 | `NoContent_Screen` | Known | Screen layout |
| 0x0075EE49 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075EEB1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075EF1B | `NoContent_Screen` | Known | Screen layout |
| 0x0075EF2F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075EF96 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075F000 | `NoContent_Screen` | Known | Screen layout |
| 0x0075F014 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075F081 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075F0F3 | `NoContent_Screen` | Known | Screen layout |
| 0x0075F107 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075F16F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075F1D8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075F1F3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075F259 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075F275 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075F354 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075F36D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075F3CE | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075F3E2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075F43C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075F458 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075F4BF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075F4D6 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075F647 | `Radio_Screen` | Known | Screen layout |
| 0x0075F657 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075F6B8 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075F73B | `LockediPod_Screen` | Known | Screen layout |
| 0x0075F7C3 | `Lock_Screen` | Known | Screen layout |
| 0x0075F7D2 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075F835 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075F897 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075F8B3 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075F925 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075F944 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075F9AC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075F9C6 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075FA2E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075FA4B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075FAB7 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075FB21 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075FB3B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075FBAB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075FC1E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075FC8F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075FCFE | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075FD6A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075FD85 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075FDFA | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075FE61 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075FEC3 | `Photos_Screen` | Known | Screen layout |
| 0x0075FF27 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075FF45 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075FFB7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075FFD4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076003A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00760055 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007600BE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007600DB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00760152 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00760176 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007601E4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007601FF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007602BC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007602D8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00760346 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00760363 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007603CE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007603EE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00760465 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00760481 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007604F1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00760510 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076057C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00760590 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00760609 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076067D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007606ED | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00760754 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007607BC | `NoContent_Screen` | Known | Screen layout |
| 0x007607D0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00760834 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076089B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007608B5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00760923 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00760995 | `NoContent_Screen` | Known | Screen layout |
| 0x007609A9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00760A13 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00760A7C | `No_Photos_Screen` | Known | Screen layout |
| 0x00760A90 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00760AF6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00760B64 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00760BD1 | `NoContent_Screen` | Known | Screen layout |
| 0x00760BE5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00760C4D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00760CB7 | `NoContent_Screen` | Known | Screen layout |
| 0x00760CCB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00760D32 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00760D9C | `NoContent_Screen` | Known | Screen layout |
| 0x00760DB0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00760E1D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00760E8F | `NoContent_Screen` | Known | Screen layout |
| 0x00760EA3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00760F0B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00760F74 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00760F8F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00760FF5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00761011 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007610F0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00761109 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076116A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076117E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007611D8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007611F4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076125B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00761272 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007613E3 | `Radio_Screen` | Known | Screen layout |
| 0x007613F3 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00761454 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007614D7 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076155F | `Lock_Screen` | Known | Screen layout |
| 0x0076156E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007615D1 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00761633 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076164F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007616C1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007616E0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00761748 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00761762 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007617CA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007617E7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00761853 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007618BD | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007618D7 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00761947 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007619BA | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00761A2B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00761A9A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00761B06 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00761B21 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00761B96 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00761BFD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00761C5F | `Photos_Screen` | Known | Screen layout |
| 0x00761CC3 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00761CE1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00761D53 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00761D70 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00761DD6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00761DF1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00761E5A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00761E77 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00761EEE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00761F12 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00761F80 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00761F9B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00762058 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00762074 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007620E2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007620FF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076216A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076218A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00762201 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076221D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076228D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007622AC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00762318 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076232C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007623A5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00762419 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00762489 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007624F0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00762558 | `NoContent_Screen` | Known | Screen layout |
| 0x0076256C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007625D0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00762637 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00762651 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007626BF | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00762731 | `NoContent_Screen` | Known | Screen layout |
| 0x00762745 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007627AF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00762818 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076282C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00762892 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00762900 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076296D | `NoContent_Screen` | Known | Screen layout |
| 0x00762981 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007629E9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00762A53 | `NoContent_Screen` | Known | Screen layout |
| 0x00762A67 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00762ACE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00762B38 | `NoContent_Screen` | Known | Screen layout |
| 0x00762B4C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00762BB9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00762C2B | `NoContent_Screen` | Known | Screen layout |
| 0x00762C3F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00762CA7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00762D10 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00762D2B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00762D91 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00762DAD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00762E8C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00762EA5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00762F06 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00762F1A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00762F74 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00762F90 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00762FF7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076300E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076317F | `Radio_Screen` | Known | Screen layout |
| 0x0076318F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007631F0 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00763273 | `LockediPod_Screen` | Known | Screen layout |
| 0x007632FB | `Lock_Screen` | Known | Screen layout |
| 0x0076330A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076336D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007633CF | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007633EB | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076345D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076347C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007634E4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007634FE | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00763566 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00763583 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007635EF | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00763659 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00763673 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007636E3 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00763756 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007637C7 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00763836 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007638A2 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007638BD | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00763932 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00763999 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007639FB | `Photos_Screen` | Known | Screen layout |
| 0x00763A5F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00763A7D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00763AEF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00763B0C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00763B72 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00763B8D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00763BF6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00763C13 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00763C8A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00763CAE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00763D1C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00763D37 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00763DF4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00763E10 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00763E7E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00763E9B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00763F06 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00763F26 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00763F9D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00763FB9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00764029 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00764048 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007640B4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007640C8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00764141 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007641B5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00764225 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076428C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007642F4 | `NoContent_Screen` | Known | Screen layout |
| 0x00764308 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076436C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007643D3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007643ED | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076445B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007644CD | `NoContent_Screen` | Known | Screen layout |
| 0x007644E1 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076454B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007645B4 | `No_Photos_Screen` | Known | Screen layout |
| 0x007645C8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076462E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076469C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00764709 | `NoContent_Screen` | Known | Screen layout |
| 0x0076471D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00764785 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007647EF | `NoContent_Screen` | Known | Screen layout |
| 0x00764803 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076486A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007648D4 | `NoContent_Screen` | Known | Screen layout |
| 0x007648E8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00764955 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007649C7 | `NoContent_Screen` | Known | Screen layout |
| 0x007649DB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00764A43 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00764AAC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00764AC7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00764B2D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00764B49 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00764C28 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00764C41 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00764CA2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00764CB6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00764D10 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00764D2C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00764D93 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00764DAA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00764F1B | `Radio_Screen` | Known | Screen layout |
| 0x00764F2B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00764F8C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076500F | `LockediPod_Screen` | Known | Screen layout |
| 0x00765097 | `Lock_Screen` | Known | Screen layout |
| 0x007650A6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00765109 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076516B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00765187 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007651F9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00765218 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00765280 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076529A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00765302 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076531F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076538B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007653F5 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076540F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076547F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007654F2 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00765563 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007655D2 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076563E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00765659 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007656CE | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00765735 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00765797 | `Photos_Screen` | Known | Screen layout |
| 0x007657FB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00765819 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076588B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007658A8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076590E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00765929 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00765992 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007659AF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00765A26 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00765A4A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00765AB8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00765AD3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00765B90 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00765BAC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00765C1A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00765C37 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00765CA2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00765CC2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00765D39 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00765D55 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00765DC5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00765DE4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00765E50 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00765E64 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00765EDD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00765F51 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00765FC1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00766028 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00766090 | `NoContent_Screen` | Known | Screen layout |
| 0x007660A4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00766108 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076616F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00766189 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007661F7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00766269 | `NoContent_Screen` | Known | Screen layout |
| 0x0076627D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007662E7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00766350 | `No_Photos_Screen` | Known | Screen layout |
| 0x00766364 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007663CA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00766438 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007664A5 | `NoContent_Screen` | Known | Screen layout |
| 0x007664B9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00766521 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076658B | `NoContent_Screen` | Known | Screen layout |
| 0x0076659F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00766606 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00766670 | `NoContent_Screen` | Known | Screen layout |
| 0x00766684 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007666F1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00766763 | `NoContent_Screen` | Known | Screen layout |
| 0x00766777 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007667DF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00766848 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00766863 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007668C9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007668E5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007669C4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007669DD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00766A3E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00766A52 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00766AAC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00766AC8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00766B2F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00766B46 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00766CB7 | `Radio_Screen` | Known | Screen layout |
| 0x00766CC7 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00766D28 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00766DAB | `LockediPod_Screen` | Known | Screen layout |
| 0x00766E33 | `Lock_Screen` | Known | Screen layout |
| 0x00766E42 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00766EA5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00766F07 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00766F23 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00766F95 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00766FB4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076701C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00767036 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076709E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007670BB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00767127 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00767191 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007671AB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076721B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076728E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007672FF | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076736E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007673DA | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007673F5 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076746A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007674D1 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00767533 | `Photos_Screen` | Known | Screen layout |
| 0x00767597 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007675B5 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00767627 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00767644 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007676AA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007676C5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076772E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076774B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007677C2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007677E6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00767854 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076786F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076792C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00767948 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007679B6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007679D3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00767A3E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00767A5E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00767AD5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00767AF1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00767B61 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00767B80 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00767BEC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00767C00 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00767C79 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00767CED | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00767D5D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00767DC4 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00767E2C | `NoContent_Screen` | Known | Screen layout |
| 0x00767E40 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00767EA4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00767F0B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00767F25 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00767F93 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00768005 | `NoContent_Screen` | Known | Screen layout |
| 0x00768019 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00768083 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007680EC | `No_Photos_Screen` | Known | Screen layout |
| 0x00768100 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00768166 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007681D4 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00768241 | `NoContent_Screen` | Known | Screen layout |
| 0x00768255 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007682BD | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00768327 | `NoContent_Screen` | Known | Screen layout |
| 0x0076833B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007683A2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076840C | `NoContent_Screen` | Known | Screen layout |
| 0x00768420 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076848D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007684FF | `NoContent_Screen` | Known | Screen layout |
| 0x00768513 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076857B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007685E4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007685FF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00768665 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00768681 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00768760 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00768779 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007687DA | `FirstBoot_Screen` | Known | Screen layout |
| 0x007687EE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00768848 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00768864 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007688CB | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007688E2 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00768A53 | `Radio_Screen` | Known | Screen layout |
| 0x00768A63 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00768AC4 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00768B47 | `LockediPod_Screen` | Known | Screen layout |
| 0x00768BCF | `Lock_Screen` | Known | Screen layout |
| 0x00768BDE | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00768C41 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00768CA3 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00768CBF | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00768D31 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00768D50 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00768DB8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00768DD2 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00768E3A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00768E57 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00768EC3 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00768F2D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00768F47 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00768FB7 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076902A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076909B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076910A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00769176 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00769191 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00769206 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076926D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007692CF | `Photos_Screen` | Known | Screen layout |
| 0x00769333 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00769351 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007693C3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007693E0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00769446 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00769461 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007694CA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007694E7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076955E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00769582 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007695F0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076960B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007696C8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007696E4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00769752 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076976F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007697DA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007697FA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00769871 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076988D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007698FD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076991C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00769988 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076999C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00769A15 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00769A89 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00769AF9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00769B60 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00769BC8 | `NoContent_Screen` | Known | Screen layout |
| 0x00769BDC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00769C40 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00769CA7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00769CC1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00769D2F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00769DA1 | `NoContent_Screen` | Known | Screen layout |
| 0x00769DB5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00769E1F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00769E88 | `No_Photos_Screen` | Known | Screen layout |
| 0x00769E9C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00769F02 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00769F70 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00769FDD | `NoContent_Screen` | Known | Screen layout |
| 0x00769FF1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076A059 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076A0C3 | `NoContent_Screen` | Known | Screen layout |
| 0x0076A0D7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076A13E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076A1A8 | `NoContent_Screen` | Known | Screen layout |
| 0x0076A1BC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076A229 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076A29B | `NoContent_Screen` | Known | Screen layout |
| 0x0076A2AF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076A317 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076A380 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076A39B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076A401 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076A41D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076A4FC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076A515 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076A576 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076A58A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076A5E4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076A600 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076A667 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076A67E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076A7EF | `Radio_Screen` | Known | Screen layout |
| 0x0076A7FF | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076A860 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076A8E3 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076A96B | `Lock_Screen` | Known | Screen layout |
| 0x0076A97A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076A9DD | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076AA3F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076AA5B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076AACD | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076AAEC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076AB54 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076AB6E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076ABD6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076ABF3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076AC5F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076ACC9 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076ACE3 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076AD53 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076ADC6 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076AE37 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076AEA6 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076AF12 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076AF2D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076AFA2 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076B009 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076B06B | `Photos_Screen` | Known | Screen layout |
| 0x0076B0CF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076B0ED | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076B15F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076B17C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076B1E2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076B1FD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076B266 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076B283 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076B2FA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076B31E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076B38C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076B3A7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076B464 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076B480 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076B4EE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076B50B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076B576 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076B596 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076B60D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076B629 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076B699 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076B6B8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076B724 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076B738 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076B7B1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076B825 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076B895 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076B8FC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076B964 | `NoContent_Screen` | Known | Screen layout |
| 0x0076B978 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076B9DC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076BA43 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076BA5D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076BACB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076BB3D | `NoContent_Screen` | Known | Screen layout |
| 0x0076BB51 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076BBBB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076BC24 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076BC38 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076BC9E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076BD0C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076BD79 | `NoContent_Screen` | Known | Screen layout |
| 0x0076BD8D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076BDF5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076BE5F | `NoContent_Screen` | Known | Screen layout |
| 0x0076BE73 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076BEDA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076BF44 | `NoContent_Screen` | Known | Screen layout |
| 0x0076BF58 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076BFC5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076C037 | `NoContent_Screen` | Known | Screen layout |
| 0x0076C04B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076C0B3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076C11C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076C137 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076C19D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076C1B9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076C298 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076C2B1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076C312 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076C326 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076C380 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076C39C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076C403 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076C41A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076C58B | `Radio_Screen` | Known | Screen layout |
| 0x0076C59B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076C5FC | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076C67F | `LockediPod_Screen` | Known | Screen layout |
| 0x0076C707 | `Lock_Screen` | Known | Screen layout |
| 0x0076C716 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076C779 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076C7DB | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076C7F7 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076C869 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076C888 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076C8F0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076C90A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076C972 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076C98F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076C9FB | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076CA65 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076CA7F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076CAEF | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076CB62 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076CBD3 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076CC42 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076CCAE | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076CCC9 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076CD3E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076CDA5 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076CE07 | `Photos_Screen` | Known | Screen layout |
| 0x0076CE6B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076CE89 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076CEFB | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076CF18 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076CF7E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076CF99 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076D002 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076D01F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076D096 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076D0BA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076D128 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076D143 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076D200 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076D21C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076D28A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076D2A7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076D312 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076D332 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076D3A9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076D3C5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076D435 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076D454 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076D4C0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076D4D4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076D54D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076D5C1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076D631 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076D698 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076D700 | `NoContent_Screen` | Known | Screen layout |
| 0x0076D714 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076D778 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076D7DF | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076D7F9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076D867 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076D8D9 | `NoContent_Screen` | Known | Screen layout |
| 0x0076D8ED | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076D957 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076D9C0 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076D9D4 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076DA3A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076DAA8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076DB15 | `NoContent_Screen` | Known | Screen layout |
| 0x0076DB29 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076DB91 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076DBFB | `NoContent_Screen` | Known | Screen layout |
| 0x0076DC0F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076DC76 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076DCE0 | `NoContent_Screen` | Known | Screen layout |
| 0x0076DCF4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076DD61 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076DDD3 | `NoContent_Screen` | Known | Screen layout |
| 0x0076DDE7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076DE4F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076DEB8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076DED3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076DF39 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076DF55 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076E034 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076E04D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076E0AE | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076E0C2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076E11C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076E138 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076E19F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076E1B6 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076E327 | `Radio_Screen` | Known | Screen layout |
| 0x0076E337 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076E398 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076E41B | `LockediPod_Screen` | Known | Screen layout |
| 0x0076E4A3 | `Lock_Screen` | Known | Screen layout |
| 0x0076E4B2 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076E515 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076E577 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076E593 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076E605 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076E624 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076E68C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076E6A6 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076E70E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076E72B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076E797 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076E801 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076E81B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076E88B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076E8FE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076E96F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076E9DE | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076EA4A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076EA65 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076EADA | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076EB41 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076EBA3 | `Photos_Screen` | Known | Screen layout |
| 0x0076EC07 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076EC25 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076EC97 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076ECB4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076ED1A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076ED35 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076ED9E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076EDBB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076EE32 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076EE56 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076EEC4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076EEDF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076EF9C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076EFB8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076F026 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076F043 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076F0AE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076F0CE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076F145 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076F161 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076F1D1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076F1F0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076F25C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076F270 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076F2E9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076F35D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076F3CD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076F434 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076F49C | `NoContent_Screen` | Known | Screen layout |
| 0x0076F4B0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076F514 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076F57B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076F595 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076F603 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076F675 | `NoContent_Screen` | Known | Screen layout |
| 0x0076F689 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076F6F3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076F75C | `No_Photos_Screen` | Known | Screen layout |
| 0x0076F770 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076F7D6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076F844 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076F8B1 | `NoContent_Screen` | Known | Screen layout |
| 0x0076F8C5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076F92D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076F997 | `NoContent_Screen` | Known | Screen layout |
| 0x0076F9AB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076FA12 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076FA7C | `NoContent_Screen` | Known | Screen layout |
| 0x0076FA90 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076FAFD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076FB6F | `NoContent_Screen` | Known | Screen layout |
| 0x0076FB83 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076FBEB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076FC54 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076FC6F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076FCD5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076FCF1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076FDD0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076FDE9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076FE4A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076FE5E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076FEB8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076FED4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076FF3B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076FF52 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007700C3 | `Radio_Screen` | Known | Screen layout |
| 0x007700D3 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00770134 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007701B7 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077023F | `Lock_Screen` | Known | Screen layout |
| 0x0077024E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007702B1 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00770313 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077032F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007703A1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007703C0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00770428 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00770442 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007704AA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007704C7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00770533 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077059D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007705B7 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00770627 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077069A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077070B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077077A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007707E6 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00770801 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00770876 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007708DD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077093F | `Photos_Screen` | Known | Screen layout |
| 0x007709A3 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007709C1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00770A33 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00770A50 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00770AB6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00770AD1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00770B3A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00770B57 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00770BCE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00770BF2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00770C60 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00770C7B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00770D38 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00770D54 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00770DC2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00770DDF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00770E4A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00770E6A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00770EE1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00770EFD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00770F6D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00770F8C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00770FF8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077100C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00771085 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007710F9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00771169 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007711D0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00771238 | `NoContent_Screen` | Known | Screen layout |
| 0x0077124C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007712B0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00771317 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00771331 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077139F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00771411 | `NoContent_Screen` | Known | Screen layout |
| 0x00771425 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077148F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007714F8 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077150C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00771572 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007715E0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077164D | `NoContent_Screen` | Known | Screen layout |
| 0x00771661 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007716C9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00771733 | `NoContent_Screen` | Known | Screen layout |
| 0x00771747 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007717AE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00771818 | `NoContent_Screen` | Known | Screen layout |
| 0x0077182C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00771899 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077190B | `NoContent_Screen` | Known | Screen layout |
| 0x0077191F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00771987 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007719F0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00771A0B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00771A71 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00771A8D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00771B6C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00771B85 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00771BE6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00771BFA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00771C54 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00771C70 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00771CD7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00771CEE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00771E5F | `Radio_Screen` | Known | Screen layout |
| 0x00771E6F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00771ED0 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00771F53 | `LockediPod_Screen` | Known | Screen layout |
| 0x00771FDB | `Lock_Screen` | Known | Screen layout |
| 0x00771FEA | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077204D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007720AF | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007720CB | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077213D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077215C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007721C4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007721DE | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00772246 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00772263 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007722CF | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00772339 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00772353 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007723C3 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00772436 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007724A7 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00772516 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00772582 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077259D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00772612 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00772679 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007726DB | `Photos_Screen` | Known | Screen layout |
| 0x0077273F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077275D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007727CF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007727EC | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00772852 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077286D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007728D6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007728F3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077296A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077298E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007729FC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00772A17 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00772AD4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00772AF0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00772B5E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00772B7B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00772BE6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00772C06 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00772C7D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00772C99 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00772D09 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00772D28 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00772D94 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00772DA8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00772E21 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00772E95 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00772F05 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00772F6C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00772FD4 | `NoContent_Screen` | Known | Screen layout |
| 0x00772FE8 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077304C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007730B3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007730CD | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077313B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007731AD | `NoContent_Screen` | Known | Screen layout |
| 0x007731C1 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077322B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00773294 | `No_Photos_Screen` | Known | Screen layout |
| 0x007732A8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077330E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077337C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007733E9 | `NoContent_Screen` | Known | Screen layout |
| 0x007733FD | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00773465 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007734CF | `NoContent_Screen` | Known | Screen layout |
| 0x007734E3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077354A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007735B4 | `NoContent_Screen` | Known | Screen layout |
| 0x007735C8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00773635 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007736A7 | `NoContent_Screen` | Known | Screen layout |
| 0x007736BB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00773723 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077378C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007737A7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077380D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00773829 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00773908 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00773921 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00773982 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00773996 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007739F0 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00773A0C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00773A73 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00773A8A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00773BFB | `Radio_Screen` | Known | Screen layout |
| 0x00773C0B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00773C6C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00773CEF | `LockediPod_Screen` | Known | Screen layout |
| 0x00773D77 | `Lock_Screen` | Known | Screen layout |
| 0x00773D86 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00773DE9 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00773E4B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00773E67 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00773ED9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00773EF8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00773F60 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00773F7A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00773FE2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00773FFF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077406B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007740D5 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007740EF | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077415F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007741D2 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00774243 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007742B2 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077431E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00774339 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007743AE | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00774415 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00774477 | `Photos_Screen` | Known | Screen layout |
| 0x007744DB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007744F9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077456B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00774588 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007745EE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00774609 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00774672 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077468F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00774706 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077472A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00774798 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007747B3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00774870 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077488C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007748FA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00774917 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00774982 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007749A2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00774A19 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00774A35 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00774AA5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00774AC4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00774B30 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00774B44 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00774BBD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00774C31 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00774CA1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00774D08 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00774D70 | `NoContent_Screen` | Known | Screen layout |
| 0x00774D84 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00774DE8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00774E4F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00774E69 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00774ED7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00774F49 | `NoContent_Screen` | Known | Screen layout |
| 0x00774F5D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00774FC7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00775030 | `No_Photos_Screen` | Known | Screen layout |
| 0x00775044 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007750AA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00775118 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00775185 | `NoContent_Screen` | Known | Screen layout |
| 0x00775199 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00775201 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077526B | `NoContent_Screen` | Known | Screen layout |
| 0x0077527F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007752E6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00775350 | `NoContent_Screen` | Known | Screen layout |
| 0x00775364 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007753D1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00775443 | `NoContent_Screen` | Known | Screen layout |
| 0x00775457 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007754BF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00775528 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00775543 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007755A9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007755C5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007756A4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007756BD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077571E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00775732 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077578C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007757A8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077580F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00775826 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00775997 | `Radio_Screen` | Known | Screen layout |
| 0x007759A7 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00775A08 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00775A8B | `LockediPod_Screen` | Known | Screen layout |
| 0x00775B13 | `Lock_Screen` | Known | Screen layout |
| 0x00775B22 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00775B85 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00775BE7 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00775C03 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00775C75 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00775C94 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00775CFC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00775D16 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00775D7E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00775D9B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00775E07 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00775E71 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00775E8B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00775EFB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00775F6E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00775FDF | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077604E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007760BA | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007760D5 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077614A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007761B1 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00776213 | `Photos_Screen` | Known | Screen layout |
| 0x00776277 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00776295 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00776307 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00776324 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077638A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007763A5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077640E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077642B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007764A2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007764C6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00776534 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077654F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077660C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00776628 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00776696 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007766B3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077671E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077673E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007767B5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007767D1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00776841 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00776860 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007768CC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007768E0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00776959 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007769CD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00776A3D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00776AA4 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00776B0C | `NoContent_Screen` | Known | Screen layout |
| 0x00776B20 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00776B84 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00776BEB | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00776C05 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00776C73 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00776CE5 | `NoContent_Screen` | Known | Screen layout |
| 0x00776CF9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00776D63 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00776DCC | `No_Photos_Screen` | Known | Screen layout |
| 0x00776DE0 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00776E46 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00776EB4 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00776F21 | `NoContent_Screen` | Known | Screen layout |
| 0x00776F35 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00776F9D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00777007 | `NoContent_Screen` | Known | Screen layout |
| 0x0077701B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00777082 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007770EC | `NoContent_Screen` | Known | Screen layout |
| 0x00777100 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077716D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007771DF | `NoContent_Screen` | Known | Screen layout |
| 0x007771F3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077725B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007772C4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007772DF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00777345 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00777361 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00777440 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00777459 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007774BA | `FirstBoot_Screen` | Known | Screen layout |
| 0x007774CE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00777528 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00777544 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007775AB | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007775C2 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00777733 | `Radio_Screen` | Known | Screen layout |
| 0x00777743 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007777A4 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00777827 | `LockediPod_Screen` | Known | Screen layout |
| 0x007778AF | `Lock_Screen` | Known | Screen layout |
| 0x007778BE | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00777921 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00777983 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077799F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00777A11 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00777A30 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00777A98 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00777AB2 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00777B1A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00777B37 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00777BA3 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00777C0D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00777C27 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00777C97 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00777D0A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00777D7B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00777DEA | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00777E56 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00777E71 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00777EE6 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00777F4D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00777FAF | `Photos_Screen` | Known | Screen layout |
| 0x00778013 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00778031 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007780A3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007780C0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00778126 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00778141 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007781AA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007781C7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077823E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00778262 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007782D0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007782EB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007783A8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007783C4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00778432 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077844F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007784BA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007784DA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00778551 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077856D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007785DD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007785FC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00778668 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077867C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007786F5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00778769 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007787D9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00778840 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007788A8 | `NoContent_Screen` | Known | Screen layout |
| 0x007788BC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00778920 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00778987 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007789A1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00778A0F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00778A81 | `NoContent_Screen` | Known | Screen layout |
| 0x00778A95 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00778AFF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00778B68 | `No_Photos_Screen` | Known | Screen layout |
| 0x00778B7C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00778BE2 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00778C50 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00778CBD | `NoContent_Screen` | Known | Screen layout |
| 0x00778CD1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00778D39 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00778DA3 | `NoContent_Screen` | Known | Screen layout |
| 0x00778DB7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00778E1E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00778E88 | `NoContent_Screen` | Known | Screen layout |
| 0x00778E9C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00778F09 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00778F7B | `NoContent_Screen` | Known | Screen layout |
| 0x00778F8F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00778FF7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00779060 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077907B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007790E1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007790FD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007791DC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007791F5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00779256 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077926A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007792C4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007792E0 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00779347 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077935E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007794CF | `Radio_Screen` | Known | Screen layout |
| 0x007794DF | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00779540 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007795C3 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077964B | `Lock_Screen` | Known | Screen layout |
| 0x0077965A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007796BD | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077971F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077973B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007797AD | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007797CC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00779834 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077984E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007798B6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007798D3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077993F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007799A9 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007799C3 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00779A33 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00779AA6 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00779B17 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00779B86 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00779BF2 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00779C0D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00779C82 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00779CE9 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00779D4B | `Photos_Screen` | Known | Screen layout |
| 0x00779DAF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00779DCD | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00779E3F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00779E5C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00779EC2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00779EDD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00779F46 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00779F63 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00779FDA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00779FFE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077A06C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077A087 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077A144 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077A160 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077A1CE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077A1EB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077A256 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077A276 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077A2ED | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077A309 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077A379 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077A398 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077A404 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077A418 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077A491 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077A505 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077A575 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077A5DC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077A644 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A658 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077A6BC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077A723 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077A73D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077A7AB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077A81D | `NoContent_Screen` | Known | Screen layout |
| 0x0077A831 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077A89B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077A904 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077A918 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077A97E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077A9EC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077AA59 | `NoContent_Screen` | Known | Screen layout |
| 0x0077AA6D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077AAD5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077AB3F | `NoContent_Screen` | Known | Screen layout |
| 0x0077AB53 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077ABBA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077AC24 | `NoContent_Screen` | Known | Screen layout |
| 0x0077AC38 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077ACA5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077AD17 | `NoContent_Screen` | Known | Screen layout |
| 0x0077AD2B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077AD93 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077ADFC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077AE17 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077AE7D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077AE99 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077AF78 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077AF91 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077AFF2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077B006 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077B060 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077B07C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077B0E3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077B0FA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077B26B | `Radio_Screen` | Known | Screen layout |
| 0x0077B27B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077B2DC | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077B35F | `LockediPod_Screen` | Known | Screen layout |
| 0x0077B3E7 | `Lock_Screen` | Known | Screen layout |
| 0x0077B3F6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077B459 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077B4BB | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077B4D7 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077B549 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077B568 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077B5D0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077B5EA | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077B652 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077B66F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077B6DB | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077B745 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077B75F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077B7CF | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077B842 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077B8B3 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077B922 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077B98E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077B9A9 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077BA1E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077BA85 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077BAE7 | `Photos_Screen` | Known | Screen layout |
| 0x0077BB4B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077BB69 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077BBDB | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077BBF8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077BC5E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077BC79 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077BCE2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077BCFF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077BD76 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077BD9A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077BE08 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077BE23 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077BEE0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077BEFC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077BF6A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077BF87 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077BFF2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077C012 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077C089 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077C0A5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077C115 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077C134 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077C1A0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077C1B4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077C22D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077C2A1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077C311 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077C378 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077C3E0 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C3F4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077C458 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077C4BF | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077C4D9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077C547 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077C5B9 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C5CD | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077C637 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077C6A0 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077C6B4 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077C71A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077C788 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077C7F5 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C809 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077C871 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077C8DB | `NoContent_Screen` | Known | Screen layout |
| 0x0077C8EF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077C956 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077C9C0 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C9D4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077CA41 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077CAB3 | `NoContent_Screen` | Known | Screen layout |
| 0x0077CAC7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077CB2F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077CB98 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077CBB3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077CC19 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077CC35 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077CD14 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077CD2D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077CD8E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077CDA2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077CDFC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077CE18 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077CE7F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077CE96 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077D007 | `Radio_Screen` | Known | Screen layout |
| 0x0077D017 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077D078 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077D0FB | `LockediPod_Screen` | Known | Screen layout |
| 0x0077D183 | `Lock_Screen` | Known | Screen layout |
| 0x0077D192 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077D1F5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077D257 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077D273 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077D2E5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077D304 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077D36C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077D386 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077D3EE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077D40B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077D477 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077D4E1 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077D4FB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077D56B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077D5DE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077D64F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077D6BE | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077D72A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077D745 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077D7BA | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077D821 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077D883 | `Photos_Screen` | Known | Screen layout |
| 0x0077D8E7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077D905 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077D977 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077D994 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077D9FA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077DA15 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077DA7E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077DA9B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077DB12 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077DB36 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077DBA4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077DBBF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077DC7C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077DC98 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077DD06 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077DD23 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077DD8E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077DDAE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077DE25 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077DE41 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077DEB1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077DED0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077DF3C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077DF50 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077DFC9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077E03D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077E0AD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077E114 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077E17C | `NoContent_Screen` | Known | Screen layout |
| 0x0077E190 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077E1F4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077E25B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077E275 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077E2E3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077E355 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E369 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077E3D3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077E43C | `No_Photos_Screen` | Known | Screen layout |
| 0x0077E450 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077E4B6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077E524 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077E591 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E5A5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077E60D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077E677 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E68B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077E6F2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077E75C | `NoContent_Screen` | Known | Screen layout |
| 0x0077E770 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077E7DD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077E84F | `NoContent_Screen` | Known | Screen layout |
| 0x0077E863 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077E8CB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077E934 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077E94F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077E9B5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077E9D1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077EAB0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077EAC9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077EB2A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077EB3E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077EB98 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077EBB4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077EC1B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077EC32 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077EDA3 | `Radio_Screen` | Known | Screen layout |
| 0x0077EDB3 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077EE14 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077EE97 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077EF1F | `Lock_Screen` | Known | Screen layout |
| 0x0077EF2E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077EF91 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077EFF3 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077F00F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077F081 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077F0A0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077F108 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077F122 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077F18A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077F1A7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077F213 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077F27D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077F297 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077F307 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077F37A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077F3EB | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077F45A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077F4C6 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077F4E1 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077F556 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077F5BD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077F61F | `Photos_Screen` | Known | Screen layout |
| 0x0077F683 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077F6A1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077F713 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077F730 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077F796 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077F7B1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077F81A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077F837 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077F8AE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077F8D2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077F940 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077F95B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077FA18 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077FA34 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077FAA2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077FABF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077FB2A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077FB4A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077FBC1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077FBDD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077FC4D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077FC6C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077FCD8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077FCEC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077FD65 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077FDD9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077FE49 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077FEB0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077FF18 | `NoContent_Screen` | Known | Screen layout |
| 0x0077FF2C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077FF90 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077FFF7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00780011 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078007F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007800F1 | `NoContent_Screen` | Known | Screen layout |
| 0x00780105 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078016F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007801D8 | `No_Photos_Screen` | Known | Screen layout |
| 0x007801EC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00780252 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007802C0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078032D | `NoContent_Screen` | Known | Screen layout |
| 0x00780341 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007803A9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00780413 | `NoContent_Screen` | Known | Screen layout |
| 0x00780427 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078048E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007804F8 | `NoContent_Screen` | Known | Screen layout |
| 0x0078050C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00780579 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007805EB | `NoContent_Screen` | Known | Screen layout |
| 0x007805FF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00780667 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007806D0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007806EB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00780751 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078076D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078084C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00780865 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007808C6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007808DA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00780934 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00780950 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007809B7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007809CE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00780B3F | `Radio_Screen` | Known | Screen layout |
| 0x00780B4F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00780BB0 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00780C33 | `LockediPod_Screen` | Known | Screen layout |
| 0x00780CBB | `Lock_Screen` | Known | Screen layout |
| 0x00780CCA | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00780D2D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00780D8F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00780DAB | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00780E1D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00780E3C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00780EA4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00780EBE | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00780F26 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00780F43 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00780FAF | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00781019 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00781033 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007810A3 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00781116 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00781187 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007811F6 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00781262 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078127D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007812F2 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00781359 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007813BB | `Photos_Screen` | Known | Screen layout |
| 0x0078141F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078143D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007814AF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007814CC | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00781532 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078154D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007815B6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007815D3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078164A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078166E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007816DC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007816F7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007817B4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007817D0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078183E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078185B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007818C6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007818E6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078195D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00781979 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007819E9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00781A08 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00781A74 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00781A88 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00781B01 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00781B75 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00781BE5 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00781C4C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00781CB4 | `NoContent_Screen` | Known | Screen layout |
| 0x00781CC8 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00781D2C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00781D93 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00781DAD | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00781E1B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00781E8D | `NoContent_Screen` | Known | Screen layout |
| 0x00781EA1 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00781F0B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00781F74 | `No_Photos_Screen` | Known | Screen layout |
| 0x00781F88 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00781FEE | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078205C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007820C9 | `NoContent_Screen` | Known | Screen layout |
| 0x007820DD | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00782145 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007821AF | `NoContent_Screen` | Known | Screen layout |
| 0x007821C3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078222A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00782294 | `NoContent_Screen` | Known | Screen layout |
| 0x007822A8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00782315 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00782387 | `NoContent_Screen` | Known | Screen layout |
| 0x0078239B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00782403 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078246C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00782487 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007824ED | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00782509 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007825E8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00782601 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00782662 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00782676 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007826D0 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007826EC | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00782753 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078276A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007828DB | `Radio_Screen` | Known | Screen layout |
| 0x007828EB | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078294C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007829CF | `LockediPod_Screen` | Known | Screen layout |
| 0x00782A57 | `Lock_Screen` | Known | Screen layout |
| 0x00782A66 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00782AC9 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00782B2B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00782B47 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00782BB9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00782BD8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00782C40 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00782C5A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00782CC2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00782CDF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00782D4B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00782DB5 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00782DCF | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00782E3F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00782EB2 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00782F23 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00782F92 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00782FFE | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00783019 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078308E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007830F5 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00783157 | `Photos_Screen` | Known | Screen layout |
| 0x007831BB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007831D9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078324B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00783268 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007832CE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007832E9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00783352 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078336F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007833E6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078340A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00783478 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00783493 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00783550 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078356C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007835DA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007835F7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00783662 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00783682 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007836F9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00783715 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00783785 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007837A4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00783810 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00783824 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078389D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00783911 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00783981 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007839E8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00783A50 | `NoContent_Screen` | Known | Screen layout |
| 0x00783A64 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00783AC8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00783B2F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00783B49 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00783BB7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00783C29 | `NoContent_Screen` | Known | Screen layout |
| 0x00783C3D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00783CA7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00783D10 | `No_Photos_Screen` | Known | Screen layout |
| 0x00783D24 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00783D8A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00783DF8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00783E65 | `NoContent_Screen` | Known | Screen layout |
| 0x00783E79 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00783EE1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00783F4B | `NoContent_Screen` | Known | Screen layout |
| 0x00783F5F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00783FC6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00784030 | `NoContent_Screen` | Known | Screen layout |
| 0x00784044 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007840B1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00784123 | `NoContent_Screen` | Known | Screen layout |
| 0x00784137 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078419F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00784208 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00784223 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00784289 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007842A5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00784384 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078439D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007843FE | `FirstBoot_Screen` | Known | Screen layout |
| 0x00784412 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078446C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00784488 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007844EF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00784506 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00784677 | `Radio_Screen` | Known | Screen layout |
| 0x00784687 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007846E8 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078476B | `LockediPod_Screen` | Known | Screen layout |
| 0x007847F3 | `Lock_Screen` | Known | Screen layout |
| 0x00784802 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00784865 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007848C7 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007848E3 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00784955 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00784974 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007849DC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007849F6 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00784A5E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00784A7B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00784AE7 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00784B51 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00784B6B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00784BDB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00784C4E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00784CBF | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00784D2E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00784D9A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00784DB5 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00784E2A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00784E91 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00784EF3 | `Photos_Screen` | Known | Screen layout |
| 0x00784F57 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00784F75 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00784FE7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00785004 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078506A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00785085 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007850EE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078510B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00785182 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007851A6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00785214 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078522F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007852D1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007852ED | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078535B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00785378 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007853E3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00785403 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078547A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00785496 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00785506 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00785525 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00785591 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007855A5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078561A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00785685 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007856F4 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00785765 | `NoContent_Screen` | Known | Screen layout |
| 0x00785779 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007857E8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078585B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007858C8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00785931 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007859A1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00785A11 | `NoContent_Screen` | Known | Screen layout |
| 0x00785A25 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00785A88 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00785AEB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00785B07 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00785B69 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00785B85 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00785BEC | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00785C03 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00785CBE | `Radio_Screen` | Known | Screen layout |
| 0x00785CCE | `Radio_Screen_Default` | Known | Screen layout |
| 0x00785D2F | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00785D9D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00785DBC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00785E2A | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00785E8F | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00785EAA | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00785F4D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00785F69 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00785FD7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00785FF4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078605F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078607F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007860F6 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00786112 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00786182 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007861A1 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078620D | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00786221 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00786296 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00786301 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00786370 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007863E1 | `NoContent_Screen` | Known | Screen layout |
| 0x007863F5 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00786464 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007864D7 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00786544 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007865AD | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078661D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078668D | `NoContent_Screen` | Known | Screen layout |
| 0x007866A1 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00786704 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00786767 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00786783 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007867E5 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00786801 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00786868 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078687F | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078693A | `Radio_Screen` | Known | Screen layout |
| 0x0078694A | `Radio_Screen_Default` | Known | Screen layout |
| 0x007869AB | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00786A19 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00786A38 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00786AA6 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00786B0B | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00786B26 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00786BC9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00786BE5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00786C53 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00786C70 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00786CDB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00786CFB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00786D72 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00786D8E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00786DFE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00786E1D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00786E89 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00786E9D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00786F12 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00786F7D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00786FEC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078705D | `NoContent_Screen` | Known | Screen layout |
| 0x00787071 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007870E0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00787153 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007871C0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00787229 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00787299 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00787309 | `NoContent_Screen` | Known | Screen layout |
| 0x0078731D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00787380 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007873E3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007873FF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00787461 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078747D | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007874E4 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007874FB | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007875B6 | `Radio_Screen` | Known | Screen layout |
| 0x007875C6 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00787627 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00787695 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007876B4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00787722 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00787787 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007877A2 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00787845 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00787861 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007878CF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007878EC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00787957 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00787977 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007879EE | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00787A0A | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00787A7A | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00787A99 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00787B05 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00787B19 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00787B8E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00787BF9 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00787C68 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00787CD9 | `NoContent_Screen` | Known | Screen layout |
| 0x00787CED | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00787D5C | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00787DCF | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00787E3C | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00787EA5 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00787F15 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00787F85 | `NoContent_Screen` | Known | Screen layout |
| 0x00787F99 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00787FFC | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078805F | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078807B | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007880DD | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007880F9 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00788160 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00788177 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00788232 | `Radio_Screen` | Known | Screen layout |
| 0x00788242 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007882A3 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00788311 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00788330 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078839E | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00788403 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078841E | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007884C1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007884DD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078854B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00788568 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007885D3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007885F3 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078866A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00788686 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007886F6 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00788715 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00788781 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00788795 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078880A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00788875 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007888E4 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00788955 | `NoContent_Screen` | Known | Screen layout |
| 0x00788969 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007889D8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00788A4B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00788AB8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00788B21 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00788B91 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00788C01 | `NoContent_Screen` | Known | Screen layout |
| 0x00788C15 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00788C78 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00788CDB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00788CF7 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00788D59 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00788D75 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00788DDC | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00788DF3 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00788EAE | `Radio_Screen` | Known | Screen layout |
| 0x00788EBE | `Radio_Screen_Default` | Known | Screen layout |
| 0x00788F1F | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00788F8D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00788FAC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078901A | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078907F | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078909A | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078913D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00789159 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007891C7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007891E4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078924F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078926F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007892E6 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00789302 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00789372 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00789391 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007893FD | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00789411 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00789486 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007894F1 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00789560 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007895D1 | `NoContent_Screen` | Known | Screen layout |
| 0x007895E5 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00789654 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007896C7 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00789734 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078979D | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078980D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078987D | `NoContent_Screen` | Known | Screen layout |
| 0x00789891 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007898F4 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00789957 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00789973 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007899D5 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007899F1 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00789A58 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00789A6F | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00789B2A | `Radio_Screen` | Known | Screen layout |
| 0x00789B3A | `Radio_Screen_Default` | Known | Screen layout |
| 0x00789B9B | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00789C09 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00789C28 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00789C96 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00789CFB | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00789D16 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00789DB9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00789DD5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00789E43 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00789E60 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00789ECB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00789EEB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00789F62 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00789F7E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00789FEE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078A00D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078A079 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078A08D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078A102 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078A16D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078A1DC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078A24D | `NoContent_Screen` | Known | Screen layout |
| 0x0078A261 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078A2D0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078A343 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078A3B0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078A419 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078A489 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078A4F9 | `NoContent_Screen` | Known | Screen layout |
| 0x0078A50D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078A570 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078A5D3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078A5EF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078A651 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078A66D | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078A6D4 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078A6EB | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078A7A6 | `Radio_Screen` | Known | Screen layout |
| 0x0078A7B6 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078A817 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078A885 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078A8A4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078A912 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078A977 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078A992 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078AA35 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078AA51 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078AABF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078AADC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078AB47 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078AB67 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078ABDE | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078ABFA | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078AC6A | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078AC89 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078ACF5 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078AD09 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078AD7E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078ADE9 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078AE58 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078AEC9 | `NoContent_Screen` | Known | Screen layout |
| 0x0078AEDD | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078AF4C | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078AFBF | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078B02C | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078B095 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078B105 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078B175 | `NoContent_Screen` | Known | Screen layout |
| 0x0078B189 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078B1EC | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078B24F | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078B26B | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078B2CD | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078B2E9 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078B350 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078B367 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078B422 | `Radio_Screen` | Known | Screen layout |
| 0x0078B432 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078B493 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078B501 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078B520 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078B58E | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078B5F3 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078B60E | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078B6B1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078B6CD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078B73B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078B758 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078B7C3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078B7E3 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078B85A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078B876 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078B8E6 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078B905 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078B971 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078B985 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078B9FA | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078BA65 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078BAD4 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078BB45 | `NoContent_Screen` | Known | Screen layout |
| 0x0078BB59 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078BBC8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078BC3B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078BCA8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078BD11 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078BD81 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078BDF1 | `NoContent_Screen` | Known | Screen layout |
| 0x0078BE05 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078BE68 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078BECB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078BEE7 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078BF49 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078BF65 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078BFCC | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078BFE3 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078C09E | `Radio_Screen` | Known | Screen layout |
| 0x0078C0AE | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078C10F | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078C17D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078C19C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078C20A | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078C26F | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078C28A | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078C32D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078C349 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078C3B7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078C3D4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078C43F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078C45F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078C4D6 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078C4F2 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078C562 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078C581 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078C5ED | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078C601 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078C676 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078C6E1 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078C750 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078C7C1 | `NoContent_Screen` | Known | Screen layout |
| 0x0078C7D5 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078C844 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078C8B7 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078C924 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078C98D | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078C9FD | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078CA6D | `NoContent_Screen` | Known | Screen layout |
| 0x0078CA81 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078CAE4 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078CB47 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078CB63 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078CBC5 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078CBE1 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078CC48 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078CC5F | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078CD1A | `Radio_Screen` | Known | Screen layout |
| 0x0078CD2A | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078CD8B | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078CDF9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078CE18 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078CE86 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078CEEB | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078CF06 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078CFA9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078CFC5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078D033 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078D050 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078D0BB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078D0DB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078D152 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078D16E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078D1DE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078D1FD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078D269 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078D27D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078D2F2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078D35D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078D3CC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078D43D | `NoContent_Screen` | Known | Screen layout |
| 0x0078D451 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078D4C0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078D533 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078D5A0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078D609 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078D679 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078D6E9 | `NoContent_Screen` | Known | Screen layout |
| 0x0078D6FD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078D760 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078D7C3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078D7DF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078D841 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078D85D | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078D8C4 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078D8DB | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078D996 | `Radio_Screen` | Known | Screen layout |
| 0x0078D9A6 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078DA07 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078DA75 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078DA94 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078DB02 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078DB67 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078DB82 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078DC25 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078DC41 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078DCAF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078DCCC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078DD37 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078DD57 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078DDCE | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078DDEA | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078DE5A | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078DE79 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078DEE5 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078DEF9 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078DF6E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078DFD9 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078E048 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078E0B9 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E0CD | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078E13C | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078E1AF | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078E21C | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078E285 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078E2F5 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078E365 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E379 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078E3DC | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078E43F | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078E45B | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078E4BD | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078E4D9 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078E540 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078E557 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078E612 | `Radio_Screen` | Known | Screen layout |
| 0x0078E622 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078E683 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078E6F1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078E710 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078E77E | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078E7E3 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078E7FE | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078E91E | `MediaLists_AudiobookChapters_Screen)` | Known | Screen layout |
| 0x0078E945 | `MediaLists_AudiobookChapters_Screen_Plane!` | Known | Screen layout |
| 0x0078EEB2 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078EECE | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0078EF3D | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0078EF56 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0078F2BE | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078F2DA | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0078F349 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0078F362 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0078F68B | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078F6A7 | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0078F716 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0078F72F | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0078F95F | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0078F97A | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x0078F9E5 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078FA00 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x0078FA73 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0078FA8E | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0078FCB7 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0078FCD2 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x0078FD3D | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078FD58 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x0078FDCB | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0078FDE6 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0079001A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00790036 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x007900B1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007900CD | `MediaLists_Albums_Screen_WithTrackCount%` | Known | Screen layout |
| 0x00790146 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x00790161 | `MediaLists_Songs_Screen_Plain1` | Known | Screen layout |
| 0x007901DC | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x007901F7 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x00790485 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007904A2 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007905E9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00790605 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x00790680 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079069B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007908E9 | `MediaLists_PodcastEpisodes_Screen'` | Known | Screen layout |
| 0x0079090E | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00790C46 | `MediaLists_TVSeasons_Screen#` | Known | Screen layout |
| 0x00790C65 | `MediaLists_TVSeasons_Screen_Default%` | Known | Screen layout |
| 0x00790CDA | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x00790CFA | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00790E82 | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x00790EA2 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007912AC | `MediaLists_GeniusPlaylist_Screen(` | Known | Screen layout |
| 0x007912D0 | `MediaLists_GeniusPlaylist_Screen_Default` | Known | Screen layout |
| 0x0079139F | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x007913C4 | `MediaLists_NestedPlaylists_Screen_Default,` | Known | Screen layout |
| 0x00791446 | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x00791465 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00791742 | `MediaLists_GeniusPlaylist_Screen(` | Known | Screen layout |
| 0x00791766 | `MediaLists_GeniusPlaylist_Screen_Default#` | Known | Screen layout |
| 0x007917DE | `Genius_Error_Screen` | Known | Screen layout |
| 0x007917F5 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079186D | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00791884 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007918F2 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0079190E | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0079197D | `Genius_Loading_Screen` | Known | Screen layout |
| 0x00791996 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x00791A60 | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x00791A85 | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x00791AFD | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x00791B1C | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00791B81 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x00791DCD | `MediaLists_GeniusPlaylist_Screen(` | Known | Screen layout |
| 0x00791DF1 | `MediaLists_GeniusPlaylist_Screen_Default"` | Known | Screen layout |
| 0x00791E6A | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x00791EDC | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00791F47 | `Genius_Error_Screen` | Known | Screen layout |
| 0x00791F5E | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x00791FD6 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00791FED | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079205B | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00792077 | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x007920E6 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007920FF | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007921F6 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x007924B8 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x007925B8 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x00792624 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079268E | `NoContent_Screen` | Known | Screen layout |
| 0x007926A2 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079270C | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00792780 | `NoContent_Screen` | Known | Screen layout |
| 0x00792794 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007927FF | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0079286B | `NoContent_Screen` | Known | Screen layout |
| 0x0079287F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007928E6 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00792952 | `NoContent_Screen` | Known | Screen layout |
| 0x00792966 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007929D3 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00792A47 | `NoContent_Screen` | Known | Screen layout |
| 0x00792A5B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00792AC3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00792B30 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00792B94 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00792BB0 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00792C1C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00792C39 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00792CA6 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00792D6D | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00792D8A | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00792E01 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00792E25 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00792EDC | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00792F46 | `NoContent_Screen` | Known | Screen layout |
| 0x00792F5A | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00792FC4 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00793038 | `NoContent_Screen` | Known | Screen layout |
| 0x0079304C | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007930B7 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00793123 | `NoContent_Screen` | Known | Screen layout |
| 0x00793137 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079319E | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079320A | `NoContent_Screen` | Known | Screen layout |
| 0x0079321E | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079328B | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007932FF | `NoContent_Screen` | Known | Screen layout |
| 0x00793313 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079337B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007933E8 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079344C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00793468 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007934D4 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007934F1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079355E | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00793625 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00793642 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007936B9 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007936DD | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00793794 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007937FE | `NoContent_Screen` | Known | Screen layout |
| 0x00793812 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079387C | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007938F0 | `NoContent_Screen` | Known | Screen layout |
| 0x00793904 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079396F | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007939DB | `NoContent_Screen` | Known | Screen layout |
| 0x007939EF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00793A56 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00793AC2 | `NoContent_Screen` | Known | Screen layout |
| 0x00793AD6 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00793B43 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00793BB7 | `NoContent_Screen` | Known | Screen layout |
| 0x00793BCB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00793C33 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00793CA0 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00793D04 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00793D20 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00793D8C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00793DA9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00793E16 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00793EDD | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00793EFA | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00793F71 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00793F95 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079404C | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007940B6 | `NoContent_Screen` | Known | Screen layout |
| 0x007940CA | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00794134 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007941A8 | `NoContent_Screen` | Known | Screen layout |
| 0x007941BC | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00794227 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00794293 | `NoContent_Screen` | Known | Screen layout |
| 0x007942A7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079430E | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079437A | `NoContent_Screen` | Known | Screen layout |
| 0x0079438E | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007943FB | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079446F | `NoContent_Screen` | Known | Screen layout |
| 0x00794483 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007944EB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00794558 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007945BC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007945D8 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00794644 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00794661 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007946CE | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00794795 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007947B2 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00794829 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079484D | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00794904 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079496E | `NoContent_Screen` | Known | Screen layout |
| 0x00794982 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007949EC | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00794A60 | `NoContent_Screen` | Known | Screen layout |
| 0x00794A74 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00794ADF | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00794B4B | `NoContent_Screen` | Known | Screen layout |
| 0x00794B5F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00794BC6 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00794C32 | `NoContent_Screen` | Known | Screen layout |
| 0x00794C46 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00794CB3 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00794D27 | `NoContent_Screen` | Known | Screen layout |
| 0x00794D3B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00794DA3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00794E10 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00794E74 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00794E90 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00794EFC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00794F19 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00794F86 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079504D | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079506A | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007950E1 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00795105 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007951BC | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00795226 | `NoContent_Screen` | Known | Screen layout |
| 0x0079523A | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007952A4 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00795318 | `NoContent_Screen` | Known | Screen layout |
| 0x0079532C | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00795397 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00795403 | `NoContent_Screen` | Known | Screen layout |
| 0x00795417 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079547E | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007954EA | `NoContent_Screen` | Known | Screen layout |
| 0x007954FE | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079556B | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007955DF | `NoContent_Screen` | Known | Screen layout |
| 0x007955F3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079565B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007956C8 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079572C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00795748 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007957B4 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007957D1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079583E | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00795905 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00795922 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00795999 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007959BD | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00795A74 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00795ADE | `NoContent_Screen` | Known | Screen layout |
| 0x00795AF2 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00795B5C | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00795BD0 | `NoContent_Screen` | Known | Screen layout |
| 0x00795BE4 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00795C4F | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00795CBB | `NoContent_Screen` | Known | Screen layout |
| 0x00795CCF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00795D36 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00795DA2 | `NoContent_Screen` | Known | Screen layout |
| 0x00795DB6 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00795E23 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00795E97 | `NoContent_Screen` | Known | Screen layout |
| 0x00795EAB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00795F13 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00795F80 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00795FE4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00796000 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0079606C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00796089 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007960F6 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007961BD | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007961DA | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00796251 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00796275 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079632C | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00796396 | `NoContent_Screen` | Known | Screen layout |
| 0x007963AA | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00796414 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00796488 | `NoContent_Screen` | Known | Screen layout |
| 0x0079649C | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00796507 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00796573 | `NoContent_Screen` | Known | Screen layout |
| 0x00796587 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007965EE | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079665A | `NoContent_Screen` | Known | Screen layout |
| 0x0079666E | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007966DB | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079674F | `NoContent_Screen` | Known | Screen layout |
| 0x00796763 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007967CB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00796838 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079689C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007968B8 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00796924 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00796941 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007969AE | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00796A75 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00796A92 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00796B09 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00796B2D | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00796FDC | `Genius_Error_Screen` | Known | Screen layout |
| 0x00796FF3 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079706B | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00797082 | `Genius_Error_Screen_NoGeniusInfoForTrack"` | Known | Screen layout |
| 0x007970F9 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x00797112 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079727F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079729E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00797568 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007975D3 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007975EA | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x00797662 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00797679 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007976E7 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00797703 | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x00797772 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079778B | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x00797855 | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x0079787A | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x007978F2 | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x00797911 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00797E97 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00797F09 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00797F74 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00797FD9 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00798043 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007980AD | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079811D | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x00798194 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x00798206 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079821D | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x00798295 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007982AC | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079831E | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x00798385 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079839E | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x00798407 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x00798472 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007984DC | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x00798543 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007985B2 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x00798620 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x00798685 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007986ED | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x00798758 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007987C3 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079882A | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x00798E83 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00798EF5 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00798F60 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00798FC5 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079902F | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00799099 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00799109 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x00799180 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007991F2 | `Genius_Error_Screen` | Known | Screen layout |
| 0x00799209 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x00799281 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00799298 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079930A | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x00799371 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079938A | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007993F3 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079945E | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007994C8 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079952F | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079959E | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079960C | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x00799671 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007996D9 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x00799744 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007997AF | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x00799816 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x00799E6D | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00799EDF | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00799F4A | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00799FAF | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079A019 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079A083 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079A0F3 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079A16A | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079A1DC | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079A1F3 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079A26B | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079A282 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079A2F4 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079A35B | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079A374 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079A3DD | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079A448 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079A4B2 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079A519 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079A588 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079A5F6 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079A65B | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079A6C3 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079A72E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079A799 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079A800 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079AE55 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079AEC7 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079AF32 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079AF97 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079B001 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079B06B | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079B0DB | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079B152 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079B1C4 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079B1DB | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079B253 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079B26A | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079B2DC | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079B343 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079B35C | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079B3C5 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079B430 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079B49A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079B501 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079B570 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079B5DE | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079B643 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079B6AB | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079B716 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079B781 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079B7E8 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079BE25 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079BE97 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079BF02 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079BF67 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079BFD1 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079C03B | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079C0AB | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079C122 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079C194 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079C1AB | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079C223 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079C23A | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079C2AC | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079C313 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079C32C | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079C395 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079C400 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079C46A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079C4D1 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079C540 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079C5AE | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079C613 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079C67B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079C6E6 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079C751 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079C7B8 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079CDF5 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079CE67 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079CED2 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079CF37 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079CFA1 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079D00B | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079D07B | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079D0F2 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079D164 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079D17B | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079D1F3 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079D20A | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079D27C | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079D2E3 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079D2FC | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079D365 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079D3D0 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079D43A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079D4A1 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079D510 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079D57E | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079D5E3 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079D64B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079D6B6 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079D721 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079D788 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079DE02 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079DE74 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079DEDF | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079DF44 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079DFAE | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079E018 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079E088 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079E0FF | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079E171 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079E188 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079E200 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079E217 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079E289 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079E2F0 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079E309 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079E372 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079E3DD | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079E447 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079E4AE | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079E51D | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079E58B | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079E5F0 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079E658 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079E6C3 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079E72E | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079E795 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079EDF4 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079EE66 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079EED1 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079EF36 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079EFA0 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079F00A | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079F07A | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079F0F1 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079F163 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079F17A | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079F1F2 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079F209 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079F27B | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079F2E2 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079F2FB | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079F364 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079F3CF | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079F439 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079F4A0 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079F50F | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079F57D | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079F5E2 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079F64A | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079F6B5 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079F720 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079F787 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079FDD0 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079FE42 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079FEAD | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079FF12 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079FF7C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079FFE6 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A0056 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A00CD | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A013F | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A0156 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A01CE | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A01E5 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A0257 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A02BE | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A02D7 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A0340 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A03AB | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A0415 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A047C | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A04EB | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A0559 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A05BE | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A0626 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A0691 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A06FC | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A0763 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A0DAC | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A0E1E | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A0E89 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A0EEE | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A0F58 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A0FC2 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A1032 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A10A9 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A111B | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A1132 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A11AA | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A11C1 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A1233 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A129A | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A12B3 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A131C | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A1387 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A13F1 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A1458 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A14C7 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A1535 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A159A | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A1602 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A166D | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A16D8 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A173F | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A1D89 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A1DFB | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A1E66 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A1ECB | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A1F35 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A1F9F | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A200F | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A2086 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A20F8 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A210F | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A2187 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A219E | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A2210 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A2277 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A2290 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A22F9 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A2364 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A23CE | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A2435 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A24A4 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A2512 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A2577 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A25DF | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A264A | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A26B5 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A271C | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A2D8B | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A2DFD | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A2E68 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A2ECD | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A2F37 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A2FA1 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A3011 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A3088 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A30FA | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A3111 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A3189 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A31A0 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A3212 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A3279 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A3292 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A32FB | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A3366 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A33D0 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A3437 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A34A6 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A3514 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A3579 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A35E1 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A364C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A36B7 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A371E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A3D9B | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A3E0D | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A3E78 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A3EDD | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A3F47 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A3FB1 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A4021 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A4098 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A410A | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A4121 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A4199 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A41B0 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A4222 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A4289 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A42A2 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A430B | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A4376 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A43E0 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A4447 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A44B6 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A4524 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A4589 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A45F1 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A465C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A46C7 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A472E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A4D8B | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A4DFD | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A4E68 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A4ECD | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A4F37 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A4FA1 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A5011 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A5088 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A50FA | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A5111 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A5189 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A51A0 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A5212 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A5279 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A5292 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A52FB | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A5366 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A53D0 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A5437 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A54A6 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A5514 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A5579 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A55E1 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A564C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A56B7 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A571E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A5D6F | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A5DE1 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A5E4C | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A5EB1 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A5F1B | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A5F85 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A5FF5 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A606C | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A60DE | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A60F5 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A616D | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A6184 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A61F6 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A625D | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A6276 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A62DF | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A634A | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A63B4 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A641B | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A648A | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A64F8 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A655D | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A65C5 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A6630 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A669B | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A6702 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A6D41 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A6DB3 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A6E1E | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A6E83 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A6EED | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A6F57 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A6FC7 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A703E | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A70B0 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A70C7 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A713F | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A7156 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A71C8 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A722F | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A7248 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A72B1 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A731C | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A7386 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A73ED | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A745C | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A74CA | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A752F | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A7597 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A7602 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A766D | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A76D4 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A7D0A | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A7D7C | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A7DE7 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A7E4C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A7EB6 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A7F20 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A7F90 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A8007 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A8079 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A8090 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A8108 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A811F | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A8191 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A81F8 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A8211 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A827A | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A82E5 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A834F | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A83B6 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A8425 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A8493 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A84F8 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A8560 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A85CB | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A8636 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A869D | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A8CEE | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A8D60 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A8DCB | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A8E30 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A8E9A | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A8F04 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A8F74 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A8FEB | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A905D | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A9074 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A90EC | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A9103 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A9175 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A91DC | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A91F5 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A925E | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A92C9 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A9333 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A939A | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A9409 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A9477 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A94DC | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A9544 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A95AF | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A961A | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A9681 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A9C88 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A9CFA | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A9D65 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A9DCA | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A9E34 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A9E9E | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A9F0E | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A9F85 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A9FF7 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007AA00E | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007AA086 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007AA09D | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007AA10F | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007AA176 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007AA18F | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007AA1F8 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007AA263 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007AA2CD | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007AA334 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007AA3A3 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007AA411 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007AA476 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007AA4DE | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007AA549 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007AA5B4 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007AA61B | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007AA96E | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AA9E5 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AAA62 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AAAD4 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AAB44 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AABBA | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AAC28 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AAC95 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AAFDA | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AB051 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AB0CE | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AB140 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AB1B0 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AB226 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AB294 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AB301 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AB66A | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AB6E1 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AB75E | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AB7D0 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AB840 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AB8B6 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AB924 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AB991 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007ABCFA | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007ABD71 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007ABDEC | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007ABE5C | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007ABED2 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007ABF40 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007ABFAD | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AC2E6 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AC35D | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AC3D8 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AC448 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AC4BE | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AC52C | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AC599 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AC8D0 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AC947 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AC9C2 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007ACA32 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007ACAA8 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007ACB16 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007ACB83 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007ACE93 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007ACF0A | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007ACF85 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007ACFF5 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AD06B | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AD0D9 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AD146 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AD74A | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007AD767 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007AD7E2 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007AD7FB | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007AD873 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007AD88C | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007AD901 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AD917 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007AD98E | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AD9A4 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007ADA1B | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007ADA38 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007ADAB0 | `Notes_List_Screen` | Known | Screen layout |
| 0x007ADAC5 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007ADC76 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007ADC93 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007ADD0E | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007ADD27 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007ADD9F | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007ADDB8 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007ADE2D | `Notes_Image_Screen` | Known | Screen layout |
| 0x007ADE43 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007ADEBA | `Notes_Image_Screen` | Known | Screen layout |
| 0x007ADED0 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007ADF47 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007ADF64 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007ADFDC | `Notes_List_Screen` | Known | Screen layout |
| 0x007ADFF1 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007AE1D2 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007AE1EF | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007AE26A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007AE283 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007AE2FB | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007AE314 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007AE389 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AE39F | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007AE416 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AE42C | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007AE4A3 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007AE4C0 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007AE538 | `Notes_List_Screen` | Known | Screen layout |
| 0x007AE54D | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007AE702 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007AE71F | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007AE79A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007AE7B3 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007AE82B | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007AE844 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007AE8B9 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AE8CF | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007AE946 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AE95C | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007AE9D3 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007AE9F0 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007AEA68 | `Notes_List_Screen` | Known | Screen layout |
| 0x007AEA7D | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007AED95 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007AEE3B | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AEEBE | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007AEF76 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x007AEFF8 | `PhotosSettingsSlideshowMusic_Screen+` | Known | Screen layout |
| 0x007AF01F | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x007AF105 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x007AF2BD | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AF31D | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AF37A | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007AF3A1 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007AF441 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AF4A1 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AF4FE | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007AF525 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007AF7C0 | `Photos_Screen` | Known | Screen layout |
| 0x007AF90C | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AF970 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007AF9D1 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007AFA2E | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007AFA8B | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007AFAF9 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007AFB56 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007AFCFC | `Photos_Screen` | Known | Screen layout |
| 0x007AFE48 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AFEAC | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007AFF0D | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007AFF6A | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007AFFC7 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B0035 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B0092 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B0238 | `Photos_Screen` | Known | Screen layout |
| 0x007B0384 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B03E8 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007B0449 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007B04A6 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007B0503 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B0571 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B05CE | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B0774 | `Photos_Screen` | Known | Screen layout |
| 0x007B08C0 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B0924 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007B0985 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007B09E2 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007B0A3F | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B0AAD | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B0B0A | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B0CB0 | `Photos_Screen` | Known | Screen layout |
| 0x007B0DFC | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B0E60 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007B0EC1 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007B0F1E | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007B0F7B | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B0FE9 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B1046 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B11EC | `Photos_Screen` | Known | Screen layout |
| 0x007B1338 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B139C | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007B13FD | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007B145A | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007B14B7 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B1525 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B1582 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B1728 | `Photos_Screen` | Known | Screen layout |
| 0x007B1874 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B18DA | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B193C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B199E | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B1A34 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B1B55 | `Photos_Screen` | Known | Screen layout |
| 0x007B1BEC | `Photos_Screen` | Known | Screen layout |
| 0x007B1D38 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B1D9E | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B1E00 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B1E62 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B1EF8 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B2019 | `Photos_Screen` | Known | Screen layout |
| 0x007B20B0 | `Photos_Screen` | Known | Screen layout |
| 0x007B21FC | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B2262 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B22C4 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B2326 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B23BC | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B24DD | `Photos_Screen` | Known | Screen layout |
| 0x007B2574 | `Photos_Screen` | Known | Screen layout |
| 0x007B26C0 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B2726 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B2788 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B27EA | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B2880 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B29A1 | `Photos_Screen` | Known | Screen layout |
| 0x007B2A38 | `Photos_Screen` | Known | Screen layout |
| 0x007B2B84 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B2BEA | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B2C4C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B2CAE | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B2D44 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B2E65 | `Photos_Screen` | Known | Screen layout |
| 0x007B3085 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B30E7 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B3155 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B31BB | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3224 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B328B | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B32F0 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B35BE | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B3620 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B368E | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B36F4 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B375D | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B37C4 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3829 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3AFA | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B3B5C | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B3BCA | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B3C30 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3C99 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3D00 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3D65 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3FD9 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007B4036 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B4098 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B4106 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B416C | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B44CA | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007B4534 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007B48DA | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007B4944 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007B4C39 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4C9C | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B4D01 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4D69 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B4DCC | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B4E34 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B4E9D | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B4F03 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4F68 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B4FD5 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007B5045 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007B50BB | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007B5131 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B51A1 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B5216 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B528D | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007B5301 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007B5373 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007B53ED | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B5460 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B54D2 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B5556 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B5580 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B5607 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B5694 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B5733 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B574D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B57C5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B57DF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B5849 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B5866 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B58DE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B5908 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B598F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B5A1C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B5ABB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5AD5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B5B4D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5B67 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B5BD1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B5BEE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B5C66 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B5C90 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B5D17 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B5DA4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B5E43 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5E5D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B5ED5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5EEF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B5F59 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B5F76 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B5FEE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B6018 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B609F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B612C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B61CB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B61E5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B625D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6277 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B62E1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B62FE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B6376 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B63A0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B6427 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B64B4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B6553 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B656D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B65E5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B65FF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B6669 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B6686 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B66FE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B6728 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B67AF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B683C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B68DB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B68F5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B696D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6987 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B69F1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B6A0E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B6A86 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B6AB0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B6B37 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B6BC4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B6C63 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6C7D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B6CF5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6D0F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B6D79 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B6D96 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B6E0E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B6E38 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B6EBF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B6F4C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B6FEB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7005 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B707D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7097 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B7101 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B711E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B7196 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B71C0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B7247 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B72D4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B7373 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B738D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B7405 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B741F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B7489 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B74A6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B751E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B7548 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B75CF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B765C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B76FB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7715 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B778D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B77A7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B7811 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B782E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B78A6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B78D0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B7957 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B79E4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B7A83 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7A9D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B7B15 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7B2F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B7B99 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B7BB6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B7C2E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B7C58 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B7CDF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B7D6C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B7E0B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7E25 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B7E9D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7EB7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B7F21 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B7F3E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B7FB6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B7FE0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B8067 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B80F4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B8193 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B81AD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B8225 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B823F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B82A9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B82C6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B833E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B8368 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B83EF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B847C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B851B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8535 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B85AD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B85C7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B8631 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B864E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B86C6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B86F0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B8777 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B8804 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B88A3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B88BD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B8935 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B894F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B89B9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B89D6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B8A4E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B8A78 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B8AFF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B8B8C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B8C2B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8C45 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B8CBD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8CD7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B8D41 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B8D5E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B8DD6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B8E00 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B8E87 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B8F14 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B8FB3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8FCD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B9045 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B905F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B90C9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B90E6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B915E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B9188 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B920F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B929C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B933B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B9355 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B93CD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B93E7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B9451 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B946E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B94E6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B9510 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B9597 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B9624 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B96C3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B96DD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B9755 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B976F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B97D9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B97F6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B987D | `SettingsMenus_ResetMainMenu_Screen9` | Known | Screen layout |
| 0x007B994D | `SettingsMenus_ResetMusicMenu_Screen:` | Known | Screen layout |
| 0x007B9A01 | `SettingsMenus_Main_Screen,` | Known | Screen layout |
| 0x007B9A73 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B9A8D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B9B05 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B9B1F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B9E5A | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007B9EC0 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007B9F1D | `Extras_Screen` | Known | Screen layout |
| 0x007B9F71 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x007BA04F | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x007BA0BD | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007BA15B | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x007BA174 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x007BA1DC | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007BA24E | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x007BA267 | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x007BA2CA | `DemoMode_Screen` | Known | Screen layout |
| 0x007BA2DD | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x007BA34A | `Debug_TestList_Screen` | Known | Screen layout |
| 0x007BA363 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x007BA3D6 | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x007BA3F1 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x007BA501 | `VoiceMemos_Context_PlayDelete_Screen)` | Known | Screen layout |
| 0x007BA529 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x007BA682 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007BA6F1 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007BA7DD | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007BA8A1 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007BA8C3 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007BA92F | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007BA951 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007BAACE | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BAAEA | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007BABB1 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007BABCC | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007BAC2F | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007BAC92 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007BAD29 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BAD45 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007BAE0C | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007BAE27 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007BAE8A | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007BAEED | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007BAF85 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BAFA1 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007BB068 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007BB083 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007BB0E6 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007BB149 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007BB1C6 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007BB231 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007BB29D | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007BB30F | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007BB37C | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007BB3E7 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x007BB453 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007BB4BB | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007BB527 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007BB59B | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007BB609 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x007BB682 | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x007D7EE0 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x007D7F65 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x007D8252 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x0098B92B | `VoiceMemos_Screen_InsufficientDiskSpace2` | Known | Screen layout |
| 0x0098D1AF | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0098D1C7 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0098D1E5 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x0098D2F1 | `NowPlaying_Screen_Music` | Known | Screen layout |
| 0x0098D31D | `MainMenu_Music_Screen_NoMusic` | Known | Screen layout |
| 0x0098D33B | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0098D359 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0098D45A | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x0098D50E | `DeleteRental_Confirmation_Screen_Unwatched` | Known | Screen layout |
| 0x0098D564 | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x0098D5B0 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0098D6B2 | `NowPlaying_Screen_ExtraInfoLoadFailed` | Known | Screen layout |
| 0x0098D70D | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x0098D726 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x0098D744 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x0098D773 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x0098D7AB | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x0098DBE2 | `Video_Settings_TV_Screen_Standard` | Known | Screen layout |
| 0x0098DC14 | `Search_Main_Screen_WithKeyboard` | Known | Screen layout |
| 0x0098DC34 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x0098DC79 | `VoiceMemos_Screen_InsufficientDiskSpace` | Known | Screen layout |
| 0x0098DD3D | `Video_Settings_TV_Screen_Wide` | Known | Screen layout |
| 0x0098DD85 | `CoverFlow_Screen_Backside` | Known | Screen layout |
| 0x0099083F | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x00990A44 | `NowPlaying_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x00990A69 | `Slideshow_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x00990B39 | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x00990B53 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x00990BE6 | `RentalDeleted_Screen_Title` | Known | Screen layout |
| 0x00990C01 | `SingleRentalExpiring_Screen_Title` | Known | Screen layout |
| 0x00990C23 | `MultipleRentalsExpiring_Screen_Title` | Known | Screen layout |
| 0x00990C48 | `DeleteRental_Screen_Title` | Known | Screen layout |
| 0x00990CEB | `VoiceMemos_DeleteAll_ScreenTitle` | Known | Screen layout |
| 0x00990D88 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00990DCB | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00990FBC | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x009910A5 | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x009910BE | `Radio_Screen_Volume` | Known | Screen layout |
| 0x009910D2 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x009910EF | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0099110E | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x009911DA | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x00991330 | `CoverFlow_Screen_BacksideForCapture` | Known | Screen layout |
| 0x0099232F | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x0099234A | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x00992641 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x00992675 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x009926B2 | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x009927C4 | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x00992914 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x0099294C | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x00992972 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x009987F2 | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x0099881D | `TVOutSlideshow_Screen_Playing` | Known | Screen layout |
| 0x0099883B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00998875 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x00998912 | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x0099897D | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x009989FD | `Extras_Screen_Debug` | Known | Screen layout |
| 0x00998B07 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x00998B27 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x00999055 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x009990B3 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x009990CE | `Extras_Screen_Lock` | Known | Screen layout |
| 0x009990E1 | `MainMenu_List_ScreenLock` | Known | Screen layout |
| 0x009990FA | `ExtrasMenu_ScreenLock` | Known | Screen layout |
| 0x0099916D | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x0099918E | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x00999261 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x00999283 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x0099938A | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x009993CA | `VoiceMemos_Screen_DeletAllAsk` | Known | Screen layout |
| 0x009993E8 | `NowPlaying_Screen_TVOutAsk` | Known | Screen layout |
| 0x00999544 | `NowPlaying_Screen_Initial` | Known | Screen layout |
| 0x0099955E | `NowPlaying_Screen_Video_TVOut_Initial` | Known | Screen layout |
| 0x0099A2C6 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x0099A347 | `RemoteUI_Screen` | Known | Screen layout |
| 0x0099A357 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x0099A36F | `Video_Settings_TV_Screen` | Known | Screen layout |
| 0x0099A388 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0099A39F | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x0099A3C3 | `AddressViewer_PartialLoad_Screen` | Known | Screen layout |
| 0x0099A3E4 | `Alarms_Alarm_Clock_Triggered_Screen` | Known | Screen layout |
| 0x0099A408 | `Alarms_Alarm_Triggered_Screen` | Known | Screen layout |
| 0x0099A426 | `Unsupported_Screen` | Known | Screen layout |
| 0x0099A439 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x0099A457 | `LockediPod_Screen` | Known | Screen layout |
| 0x0099A469 | `DiskMode_Screen` | Known | Screen layout |
| 0x0099A479 | `DemoMode_Screen` | Known | Screen layout |
| 0x0099A489 | `Notes_Image_Screen` | Known | Screen layout |
| 0x0099A49C | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x0099A4BA | `NowPlaying_Idle_Screen` | Known | Screen layout |
| 0x0099A4D1 | `Game_Screen` | Known | Screen layout |
| 0x0099A4DD | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x0099A4FA | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x0099A513 | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x0099A534 | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x0099A559 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x0099A56C | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x0099A589 | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x0099A5AA | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x0099A5CF | `Notes_Loading_Screen` | Known | Screen layout |
| 0x0099A5E4 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0099A5FA | `SettingsMenus_AdjustScrolling_Screen` | Known | Screen layout |
| 0x0099A61F | `Game_Running_Screen` | Known | Screen layout |
| 0x0099A633 | `Stopwatch_Screen` | Known | Screen layout |
| 0x0099A644 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0099A65B | `Clock_Screen` | Known | Screen layout |
| 0x0099A668 | `TVOutSlideshowAsk_Screen` | Known | Screen layout |
| 0x0099A681 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x0099A697 | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x0099A6B5 | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x0099A6D1 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x0099A6E2 | `DemoMode_Main_Screen` | Known | Screen layout |
| 0x0099A6F7 | `Search_Main_Screen` | Known | Screen layout |
| 0x0099A70A | `AddressViewer_Main_Screen` | Known | Screen layout |
| 0x0099A724 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x0099A739 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x0099A74F | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0099A769 | `Clock_Region_Screen` | Known | Screen layout |
| 0x0099A77D | `RentalDeleted_Notification_Screen` | Known | Screen layout |
| 0x0099A79F | `SingleRentalExpiring_Notification_Screen` | Known | Screen layout |
| 0x0099A7C8 | `MultipleRentalsExpiring_Notification_Screen` | Known | Screen layout |
| 0x0099A7F4 | `RentalError_Notification_Screen` | Known | Screen layout |
| 0x0099A814 | `DeleteRental_Confirmation_Screen` | Known | Screen layout |
| 0x0099A835 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0099A84D | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x0099A86B | `Video_Settings_Fit_To_Screen` | Known | Screen layout |
| 0x0099A888 | `RentalInfo_Screen` | Known | Screen layout |
| 0x0099A89A | `Radio_Screen` | Known | Screen layout |
| 0x0099A8A7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0099A8BB | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x0099A8D5 | `ChargingMode_LowPower_Screen` | Known | Screen layout |
| 0x0099A8F2 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0099A90C | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0099A926 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0099A940 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0099A954 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0099A96D | `Extras_Screen` | Known | Screen layout |
| 0x0099A97B | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x0099A998 | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x0099A9BA | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x0099A9D3 | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x0099A9F1 | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x0099AA0A | `Video_Settings_Screen` | Known | Screen layout |
| 0x0099AA20 | `SettingsMenus_AudiobookSettings_Screen` | Known | Screen layout |
| 0x0099AA47 | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x0099AA6D | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0099AA83 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0099AA9B | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x0099AABE | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x0099AADB | `MediaLists_Rentals_Screen` | Known | Screen layout |
| 0x0099AAF5 | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x0099AB19 | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x0099AB32 | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x0099AB54 | `ToDo_Instructions_Screen` | Known | Screen layout |
| 0x0099AB6D | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x0099AB89 | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x0099ABA3 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x0099ABC4 | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x0099ABE0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0099ABF8 | `VoiceMemos_Screen` | Known | Screen layout |
| 0x0099AC0A | `No_Photos_Screen` | Known | Screen layout |
| 0x0099AC1B | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x0099AC35 | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x0099AC51 | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x0099AC75 | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x0099AC95 | `ChargingMode_Progress_Screen` | Known | Screen layout |
| 0x0099ACB2 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x0099ACC8 | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x0099ACE3 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x0099ACFF | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x0099AD21 | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x0099AD42 | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x0099AD5C | `MediaLists_Genius_Screen` | Known | Screen layout |
| 0x0099AD75 | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x0099AD8F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0099ADAE | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x0099ADCF | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x0099ADE7 | `NoContent_Screen` | Known | Screen layout |
| 0x0099ADF8 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0099AE0E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0099AE1F | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x0099AE35 | `Notes_List_Screen` | Known | Screen layout |
| 0x0099AE47 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x0099AE5D | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x0099AE7E | `MediaLists_GeniusPlaylist_Screen` | Known | Screen layout |
| 0x0099AE9F | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x0099AEB9 | `TODOS_Menu_Screen` | Known | Screen layout |
| 0x0099AECB | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x0099AEE1 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0099AEFD | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0099AF12 | `Games_Menu_Screen` | Known | Screen layout |
| 0x0099AF24 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0099AF37 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x0099AF56 | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x0099AF75 | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x0099AF99 | `ContextualMenu_Screen` | Known | Screen layout |
| 0x0099AFAF | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x0099AFC5 | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x0099AFE3 | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x0099B006 | `TVOutSlideshow_Screen` | Known | Screen layout |
| 0x0099B01C | `CoverFlow_Screen` | Known | Screen layout |
| 0x0099B02D | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0099B041 | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x0099B063 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x0099B07B | `Stopwatch_SessionSummary_Screen` | Known | Screen layout |
| 0x0099B09B | `SettingsMenu_AboutWithAccessory_Screen` | Known | Screen layout |
| 0x0099B0C2 | `Clock_NorthAmerica_City_Screen` | Known | Screen layout |
| 0x0099B0E1 | `Clock_SouthAmerica_City_Screen` | Known | Screen layout |
| 0x0099B100 | `Clock_Africa_City_Screen` | Known | Screen layout |
| 0x0099B119 | `Clock_Australia_City_Screen` | Known | Screen layout |
| 0x0099B135 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x0099B14C | `Clock_Pacific_City_Screen` | Known | Screen layout |
| 0x0099B166 | `Clock_Atlantic_City_Screen` | Known | Screen layout |
| 0x0099B181 | `Clock_Europe_City_Screen` | Known | Screen layout |
| 0x0099B261 | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x0099B2B2 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0099B2D5 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0099B2FD | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x0099B698 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0099B79B | `RentalInfo_Screen_ExpiringSoon` | Known | Screen layout |
| 0x0099B7F1 | `RentalInfo_Screen_NoAlbumArt_ExpiringSoon` | Known | Screen layout |
| 0x0099BBC0 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x0099BC16 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x0099BD67 | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x0099BD84 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x0099C158 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x0099C27A | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0099C29C | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0099C309 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x0099C328 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x0099C96A | `ChargingMode_ScreenLayout_LowPower` | Known | Screen layout |
| 0x0099D307 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x0099D320 | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x0099D468 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x0099D544 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0099D562 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0099D582 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0099D68D | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x0099D6A9 | `Extras_Screen_Games` | Known | Screen layout |
| 0x0099D7AF | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0099D7CE | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x0099D7EA | `Extras_Screen_Notes` | Known | Screen layout |
| 0x0099D8B5 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x0099D990 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0099DB5E | `MainMenu_Music_Screen_NoAudiobooks` | Known | Screen layout |
| 0x0099DB81 | `MainMenus_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x0099DBA4 | `NoContent_Screen_NoAudiobooks` | Known | Screen layout |
| 0x0099DBDE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0099DBFD | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0099DC1E | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0099DCCD | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x0099DCEA | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x0099DD69 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0099DE4D | `MainMenu_Music_Screen_NoCompilations` | Known | Screen layout |
| 0x0099DE72 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0099DFF9 | `MainMenu_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0099E01C | `MainMenu_Videos_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0099E041 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0099E060 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0099E07F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0099E0A0 | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x0099E0DE | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x0099E0FF | `Extras_Screen_VoiceMemos` | Known | Screen layout |
| 0x0099E16A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0099E19C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0099E1BB | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x0099E268 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x0099E2D4 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0099E3CD | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x0099E3E9 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x0099E46C | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x0099E487 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0099E4A8 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x0099E557 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0099E58B | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x0099E5AC | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0099E64F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0099E670 | `MainMenu_Videos_Screen_NoPlaylists` | Known | Screen layout |
| 0x0099E693 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0099E6E2 | `MainMenu_Music_Screen_NoArtists` | Known | Screen layout |
| 0x0099E789 | `NowPlaying_Screen_Genius` | Known | Screen layout |
| 0x0099E7D2 | `Genius_Error_Screen_NoGenius` | Known | Screen layout |
| 0x0099E7EF | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x0099E80E | `Extras_Screen_NikePlus` | Known | Screen layout |
| 0x0099E95E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0099E97D | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0099E99E | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x0099EE09 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x0099EEBC | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x0099EF36 | `Notes_Contents_Screen_Alt` | Known | Screen layout |
| 0x0099EF50 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x0099EFFC | `VoiceMemos_Screen_InsufficientDiskSpace2_Default` | Known | Screen layout |
| 0x0099F0AE | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x0099F153 | `VoiceMemos_Screen_InsufficientDiskSpace_Default` | Known | Screen layout |
| 0x0099F183 | `NowPlaying_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x0099F1B0 | `Slideshow_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x0099FE41 | `Radio_Screen_Tuning_Default` | Known | Screen layout |
| 0x0099FEA2 | `VoiceMemos_Screen_DeletAllAsk_Default` | Known | Screen layout |
| 0x0099FEC8 | `NowPlaying_Screen_TVOutAsk_Default` | Known | Screen layout |
| 0x0099FEEB | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0099FF09 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x0099FF35 | `AddressViewer_PartialLoad_Screen_Default` | Known | Screen layout |
| 0x0099FF5E | `Alarms_Alarm_Clock_Triggered_Screen_Default` | Known | Screen layout |
| 0x0099FF8A | `Alarms_Alarm_Triggered_Screen_Default` | Known | Screen layout |
| 0x0099FFB0 | `Unsupported_Screen_Default` | Known | Screen layout |
| 0x0099FFCB | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x0099FFF1 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x009A0009 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x009A0024 | `Game_Screen_Default` | Known | Screen layout |
| 0x009A0038 | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x009A005E | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x009A007F | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x009A00A8 | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x009A00D2 | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x009A00FF | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x009A0128 | `Notes_Loading_Screen_Default` | Known | Screen layout |
| 0x009A0145 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x009A0163 | `Clock_Screen_Default` | Known | Screen layout |
| 0x009A0178 | `TVOutSlideshowAsk_Screen_Default` | Known | Screen layout |
| 0x009A0199 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x009A01B7 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x009A01DD | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x009A0201 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x009A021A | `AddressViewer_Main_Screen_Default` | Known | Screen layout |
| 0x009A023C | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x009A0259 | `MainMenus_Main_Screen_Default` | Known | Screen layout |
| 0x009A0277 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x009A0294 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x009A02B0 | `RentalDeleted_Notification_Screen_Default` | Known | Screen layout |
| 0x009A02DA | `SingleRentalExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x009A030B | `MultipleRentalsExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x009A033F | `RentalError_Notification_Screen_Default` | Known | Screen layout |
| 0x009A0367 | `DeleteRental_Confirmation_Screen_Default` | Known | Screen layout |
| 0x009A0390 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x009A03BC | `RentalInfo_Screen_Default` | Known | Screen layout |
| 0x009A03D6 | `Radio_Screen_Default` | Known | Screen layout |
| 0x009A03EB | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x009A0407 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x009A0429 | `Extras_Screen_Default` | Known | Screen layout |
| 0x009A043F | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x009A0465 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x009A0486 | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x009A04A4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x009A04C6 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x009A04F2 | `ToDo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009A0513 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009A0537 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x009A0559 | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x009A057D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x009A059C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x009A05B5 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x009A05D7 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x009A05FB | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x009A0619 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x009A063D | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x009A0667 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x009A0690 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x009A06B2 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x009A06D3 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x009A06F3 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x009A0711 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x009A072A | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x009A0748 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x009A0762 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x009A0780 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x009A07A9 | `MediaLists_GeniusPlaylist_Screen_Default` | Known | Screen layout |
| 0x009A07D2 | `TODOS_Menu_Screen_Default` | Known | Screen layout |
| 0x009A07EC | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x009A080A | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x009A0827 | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x009A0841 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x009A085C | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x009A087B | `ContextualMenu_Screen_Default` | Known | Screen layout |
| 0x009A0899 | `Debug_MainMenu_Screen_Default` | Known | Screen layout |
| 0x009A08B7 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x009A08D5 | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x009A08EE | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x009A090A | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x009A0934 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x009A0954 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x009A097C | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009A09A3 | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009A09CA | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x009A09EB | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x009A0A0F | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x009A0A2E | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x009A0A50 | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x009A0A73 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x009A0A94 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x009A0B22 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x009A0B52 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x009A0B74 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x009A0BE5 | `RentalInfo_Screen_NoAlbumArt_Default` | Known | Screen layout |
| 0x009A0C0A | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x009A11E5 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009A1211 | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009A1256 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x009A127E | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x009A129F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x009A12C0 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x009A12E6 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x009A1303 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x009A1325 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x009A1349 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x009A136D | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x009A153D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x009A1618 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x009A1669 | `NowPlaying_Screen_Video_TVOut` | Known | Screen layout |
| 0x009A17DB | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x009A1802 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x009A1D3B | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x009A1EF8 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x009A20EA | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x009A23B6 | `Lock_Screen_Default_Layout` | Known | Screen layout |
| 0x009A244C | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x009A2473 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x009A268F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x009A2769 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x009A27D0 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009A27FA | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009A50FE | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x009A514A | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x009A5228 | `MainMenu_List_ScreenLock_x` | Known | Screen layout |
| 0x009A54F6 | `RentalInfo_Screen_ExpiresToday` | Known | Screen layout |
| 0x009A554C | `RentalInfo_Screen_NoAlbumArt_ExpiresToday` | Known | Screen layout |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000908B | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | Known | RTOS |
| 0x002A2990 | `  K - RTXC` | Known | RTOS |
| 0x002A3998 | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | Known | RTOS |
| 0x0098A51C | `Returning from RTXCBug` | Known | RTOS |

---

## 9. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000D2278 | `HostOSTask` | Known | RTOS task thread |
| 0x0012CD4C | `MP3ExampleTask` | Known | RTOS task thread |
| 0x00132214 | `USBDeviceTask` | Known | RTOS task thread |
| 0x0013C380 | `DiskReaderTask` | Known | RTOS task thread |
| 0x0014C4F0 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0014C504 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0019FBC4 | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x001DB458 | `MeCCAIOTask` | Known | RTOS task thread |
| 0x0020E6C8 | `StreamCacheMassStorageManagerTimeOutTask` | Known | RTOS task thread |
| 0x0020E844 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x00290B30 | `FirewireTask` | Known | RTOS task thread |
| 0x00290B44 | `TouchwheelTask` | Known | RTOS task thread |
| 0x00290B58 | `AudioOutStateTask` | Known | RTOS task thread |
| 0x00290B84 | `DiskMgrTask` | Known | RTOS task thread |
| 0x00290B94 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x00290BA8 | `MikeyTask` | Known | RTOS task thread |
| 0x00290BB8 | `TopPlugTask` | Known | RTOS task thread |
| 0x00290BC8 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x00290C40 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x00290C68 | `AlarmTask` | Known | RTOS task thread |
| 0x00290C87 | `"USBAudioTask` | Known | RTOS task thread |
| 0x002A3030 | `Undefined Task` | Known | RTOS task thread |
| 0x003E5128 | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x003E87F4 | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x003F0F00 | `MeCCARecordingTask` | Known | RTOS task thread |
| 0x008DCB00 | `WaveFileDebugTask` | Known | RTOS task thread |

---

## 10. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00249028 | `Channel Reserved` | Known | Logging channel |
| 0x0024903C | `Channel AppBoot` | Known | Logging channel |
| 0x0024904C | `Channel BufferedSongReading` | Known | Logging channel |
| 0x00249068 | `Channel PrefsWriting` | Known | Logging channel |
| 0x00249080 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x002490A0 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x002490B8 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x002490D4 | `Channel TestLogging` | Known | Logging channel |
| 0x002490E8 | `Channel AppFileLoading` | Known | Logging channel |
| 0x00249100 | `Channel VCardReading` | Known | Logging channel |
| 0x00249118 | `Channel LongSongScanning` | Known | Logging channel |
| 0x0024918C | `Channel VoiceRecording` | Known | Logging channel |
| 0x002491A4 | `Channel PhotoImporting` | Known | Logging channel |
| 0x002491BC | `Channel Notes` | Known | Logging channel |
| 0x002491CC | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x002491E8 | `Channel DiskMode` | Known | Logging channel |
| 0x002491FC | `Channel Firewire` | Known | Logging channel |
| 0x00249210 | `Channel USB` | Known | Logging channel |
| 0x00249230 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x00249248 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00081D20 | `gamedata_RW` | Known | Game system |
| 0x00081D3C | `gamedata_ShareRW` | Known | Game system |
| 0x00081D50 | `games_RO` | Known | Game system |
| 0x0095F38F | `11TCGamesMenu` | Known | Game system |
| 0x0095F463 | `12TCGameScreen` | Known | Game system |
| 0x0096022F | `27TSilverCntlrTransitionAddonI11TCGamesMenuE` | Known | Game system |
| 0x009602E4 | `27TSilverCntlrTransitionAddonI12TCGameScreenE` | Known | Game system |
| 0x0098A576 | `iPod_Control/games_RO/` | Known | Game system |
| 0x0098A58D | `Resources/Games/games_RO/` | Known | Game system |
| 0x00996004 | `GamesMenu_StatusBar_String` | Known | Game system |
| 0x0099678C | `AboutScreen_Games_String` | Known | Game system |
| 0x0099D6BD | `MainMenu_List_Games` | Known | Game system |
| 0x0099D6D1 | `ExtrasMenu_Games` | Known | Game system |
| 0x009A5297 | `MainMenu_List_Games_x` | Known | Game system |

---

## 12. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00092210 | `adrmmp4a` | Known | DRM system |
| 0x00139830 | `AppleDRMVersion` | Known | DRM system |
| 0x001398D0 | `AppleDRM` | Known | DRM system |
| 0x0013AB50 | `AppleVideoDRM` | Known | DRM system |
| 0x0013E0F0 | `tx3gdrmsp608aavdmp4aesds` | Known | DRM system |
| 0x001E8E00 | `drmttx3g` | Known | DRM system |
| 0x0098A9FF | `DRMLevel` | Known | DRM system |

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
| 0x0021B6FC | `%s/sqlite_` | Known | SQLite database |
| 0x00281C4C | `iPod_Control/iTunes/primary.db` | Known | iTunes database |
| 0x00282814 | `iPod_Control/iTunes/Extras.itdb` | Known | iTunes database |
| 0x002A6484 | `sqlite3BtreeInitPage() returns error code %d` | Known | SQLite database |
| 0x002A9760 | `sqlite_master` | Known | SQLite database |
| 0x002A9770 | `sqlite_temp_master` | Known | SQLite database |
| 0x002C06A8 | `sqlite_stat1` | Known | SQLite database |
| 0x002C06B8 | `CREATE TABLE %Q.sqlite_stat1(tbl,idx,stat)` | Known | SQLite database |
| 0x002C06E4 | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x002CB08C | `sqlite_subquery_%p_` | Known | SQLite database |
| 0x003602E4 | `sqlite_master` | Known | SQLite database |
| 0x003602F4 | `sqlite_temp_master` | Known | SQLite database |
| 0x00360618 | `sqlite_` | Known | SQLite database |
| 0x00360658 | `sqlite_master` | Known | SQLite database |
| 0x00360668 | `sqlite_temp_master` | Known | SQLite database |
| 0x00360680 | `sqlite_sequence` | Known | SQLite database |
| 0x00360690 | `UPDATE "%w".sqlite_sequence set name = %Q WHERE name = %Q` | Known | SQLite database |
| 0x00360774 | `sqlite_stat1` | Known | SQLite database |
| 0x00360784 | `SELECT idx, stat FROM %Q.sqlite_stat1` | Known | SQLite database |
| 0x00361460 | `sqlite_` | Known | SQLite database |
| 0x0036165C | `sqlite_master` | Known | SQLite database |
| 0x0036166C | `sqlite_temp_master` | Known | SQLite database |
| 0x00364388 | `sqlite_` | Known | SQLite database |
| 0x00365674 | `sqlite_autoindex_` | Known | SQLite database |
| 0x00365688 | `sqlite_master` | Known | SQLite database |
| 0x00365698 | `sqlite_temp_master` | Known | SQLite database |
| 0x00366AF0 | `sqlite_master` | Known | SQLite database |
| 0x00366B00 | `sqlite_temp_master` | Known | SQLite database |
| 0x00366B34 | `sqlite_stat1` | Known | SQLite database |
| 0x00366B44 | `DELETE FROM %Q.sqlite_stat1 WHERE idx=%Q` | Known | SQLite database |
| 0x00366E2C | `sqlite_master` | Known | SQLite database |
| 0x00366E3C | `sqlite_temp_master` | Known | SQLite database |
| 0x00366EB0 | `DELETE FROM %s.sqlite_sequence WHERE name=%Q` | Known | SQLite database |
| 0x00366F18 | `sqlite_stat1` | Known | SQLite database |
| 0x00366F28 | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x003672A0 | `sqlite_master` | Known | SQLite database |
| 0x003672B0 | `sqlite_temp_master` | Known | SQLite database |
| 0x003676C8 | `sqlite_master` | Known | SQLite database |
| 0x003676D8 | `sqlite_temp_master` | Known | SQLite database |
| 0x003676F0 | `CREATE TABLE %Q.sqlite_sequence(name,seq)` | Known | SQLite database |
| 0x0036A978 | `sqlite_master` | Known | SQLite database |
| 0x0036A988 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036CD70 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036CD88 | `sqlite_master` | Known | SQLite database |
| 0x0036E564 | `sqlite3_extension_init` | Known | SQLite database |
| 0x0036ED58 | `sqlite_master` | Known | SQLite database |
| 0x0036ED68 | `sqlite_temp_master` | Known | SQLite database |
| 0x00373148 | `sqlite_attach` | Known | SQLite database |
| 0x0037315C | `sqlite_detach` | Known | SQLite database |
| 0x00375E90 | `sqlite_master` | Known | SQLite database |
| 0x00375EA0 | `sqlite_temp_master` | Known | SQLite database |
| 0x00375EF0 | `sqlite_sequence` | Known | SQLite database |
| 0x0037B77C | `sqlite_master` | Known | SQLite database |
| 0x0037B78C | `sqlite_temp_master` | Known | SQLite database |
| 0x0037EB20 | `sqlite_master` | Known | SQLite database |
| 0x0037EB30 | `sqlite_temp_master` | Known | SQLite database |
| 0x0038CCCC | `sqlite_attach` | Known | SQLite database |
| 0x0038CCDC | `sqlite_detach` | Known | SQLite database |
| 0x003DE3E4 | `iTunesDB` | Known | iTunes database |
| 0x003DE3F0 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x008D8B4F | `SQLite format 3` | Known | SQLite database |
| 0x008DB1FC | `CREATE TABLE sqlite_master(` | Known | SQLite database |
| 0x008DB264 | `CREATE TEMP TABLE sqlite_temp_master(` | Known | SQLite database |
| 0x008DB92C | `illegal return value (%d) from the authorization function - should be SQLITE_OK,` | Known | SQLite database |
| 0x008DB9E4 | `SELECT 'CREATE TABLE vacuum_db.' || substr(sql,14)   FROM sqlite_master WHERE ty` | Known | SQLite database |
| 0x008DBA6C | `SELECT 'CREATE INDEX vacuum_db.' || substr(sql,14)  FROM sqlite_master WHERE sql` | Known | SQLite database |
| 0x008DBAD4 | `SELECT 'CREATE UNIQUE INDEX vacuum_db.' || substr(sql,21)   FROM sqlite_master W` | Known | SQLite database |
| 0x008DBB4C | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x008DBBFC | `SELECT 'DELETE FROM vacuum_db.' || quote(name) || ';' FROM vacuum_db.sqlite_mast` | Known | SQLite database |
| 0x008DBC70 | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x008DBD08 | `INSERT INTO vacuum_db.sqlite_master   SELECT type, name, tbl_name, rootpage, sql` | Known | SQLite database |
| 0x008DBEC8 | `UPDATE %Q.%s SET sql = CASE WHEN type = 'trigger' THEN sqlite_rename_trigger(sql` | Known | SQLite database |
| 0x008DC03C | `UPDATE sqlite_temp_master SET sql = sqlite_rename_trigger(sql, %Q), tbl_name = %` | Known | SQLite database |
| 0x008DC278 | `sqlite3_get_table() called with two or more incompatible queries` | Known | SQLite database |
| 0x009A5D26 | `sqlite_rename_table` | Known | SQLite database |
| 0x009A5EA9 | `sqlite_version` | Known | SQLite database |
| 0x009A5F43 | `sqlite_rename_trigger` | Known | SQLite database |
| 0x009A6267 | `SQLite_iPod_VFS` | Known | SQLite database |

---

## 14. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0005EE5C | `cI: could not read CE-ATA task file` | Known | Hardware |
| 0x0005EE84 | `cI: CE-ATA signature missing (%x,%x)` | Known | Hardware |
| 0x0005EEDC | `cI: CE-ATA interrupt enable failed` | Known | Hardware |
| 0x00124C70 | `NAND FLASH DRIVE` | Known | Hardware |
| 0x00139D78 | `FireWireGUID` | Known | FireWire |
| 0x00139D88 | `FireWireVersion` | Known | FireWire |
| 0x0013A464 | `FireWire` | Known | FireWire |
| 0x0035B620 | `CE-ATA init failed` | Known | Hardware |
| 0x0035BAE0 | `ISDIE: CE-ATA interrupt enable failed` | Known | Hardware |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00726502 | `Settings_RadioRegionsMenu_Layout"` | Known | FM Radio |
| 0x0072658B | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x007D7388 | `Radio Regions` | Known | FM Radio |
| 0x00827314 | `Radio-Regionen` | Known | FM Radio |
| 0x0095FE2C | `23TCSettings_RadioRegions` | Known | FM Radio |
| 0x00960D3F | `27TSilverCntlrTransitionAddonI23TCSettings_RadioRegionsE` | Known | FM Radio |
| 0x00993065 | `Settings_RadioRegions_Australia_String` | Known | FM Radio |
| 0x0099308C | `Settings_RadioRegions_Asia_String` | Known | FM Radio |
| 0x009942F1 | `Settings_RadioRegions_Europe_String` | Known | FM Radio |
| 0x00995900 | `Settings_RadioRegions_Japan_String` | Known | FM Radio |
| 0x009965A9 | `Settings_RadioRegions_Americas_String` | Known | FM Radio |
| 0x00996C8B | `SettingsMenu_RadioRegions_String` | Known | FM Radio |
| 0x0099A185 | `RadioRegionsMenuItem` | Known | FM Radio |
| 0x0099DDD6 | `SettingsMenu_ListItem_RadioRegions` | Known | FM Radio |
| 0x009A1FC4 | `SettingsInfo_Template_RadioRegions_Layout` | Known | FM Radio |
| 0x009A1FEE | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x009A2650 | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 16. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x008660FC | `Fotocamera` | Known | Camera |
| 0x00866660 | `Collega fotocamera o inserisci scheda` | Known | Camera |
| 0x008666D8 | `Fotocamera non supportata` | Known | Camera |
| 0x008857A4 | `Camera` | Known | Camera |
| 0x00885D24 | `Sluit camera of kaart aan` | Known | Camera |
| 0x00885D90 | `Camera niet ondersteund` | Known | Camera |
| 0x009930AE | `Photos_DevName_Camera_String` | Known | Camera |

---

## 17. Pedometer

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x009A55FF | `NikePlus_Remote_Step_Away` | Known | Pedometer |
| 0x009A5619 | `NikePlus_Step_Away` | Known | Pedometer |
| 0x009A5EE4 | `AggStep` | Known | Pedometer |

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
| 0x0014D994 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x0014DBF0 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x0015A574 | `Resources/UI/active.bin` | Filesystem Path |  |
| 0x0015A58C | `Resources/UI/` | Filesystem Path |  |
| 0x0017DFC0 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x0017EEEC | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path |  |
| 0x0017EF14 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001A320C | `iPod_Control/Device/PlayCounts` | Filesystem Path |  |
| 0x001B948C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B953C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B96B8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9850 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B98F8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9AA8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9B4C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9BF0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9C94 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9D38 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9DE8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9E8C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9F30 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9FE0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA090 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA140 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA2AC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA35C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA40C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA4B0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA560 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA654 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA6F8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA7AC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA868 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA918 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAA3C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAAF8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BABA8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAD64 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAE28 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAED8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAF94 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB0D0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB19C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB258 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB2FC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB3A0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB45C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB518 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB5E0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB684 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB74C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB814 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB8C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB98C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBA54 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBB04 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBBB4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBC78 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBD28 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBDD8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBE88 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBF5C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC030 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC130 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC210 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC318 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC404 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003DE462 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003E49C8 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x003E6BBC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003E700E | `iPod_Control/Device` | Filesystem Path |  |
| 0x003E8960 | `Resources/Fonts` | Filesystem Path |  |
| 0x003F0ECC | `iPod_Control:Device:voices:` | Filesystem Path |  |
| 0x0098A451 | `Resources/Games/` | Filesystem Path |  |
| 0x0098A86F | `iPod_Control/Device` | Filesystem Path |  |
| 0x0098A883 | `iPod_Control/Device/accessories.reg` | Filesystem Path |  |
| 0x0098A976 | `iPod_Control/Tones` | Filesystem Path |  |

---

## 19. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x008DF240 | `c:\BWA\N25CFirmwareWin-10\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftu` | Build Path |  |
| 0x008DF298 | `c:\BWA\N25CFirmwareWin-10\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fts` | Build Path |  |
| 0x008DF2F0 | `c:\BWA\N25CFirmwareWin-10\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fto` | Build Path |  |
| 0x008E9F40 | `c:\BWA\N25CFirmwareWin-10\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\` | Build Path |  |
| 0x008EAABC | `c:\BWA\N25CFirmwareWin-10\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfd` | Build Path |  |
| 0x008EBCB8 | `c:\BWA\N25CFirmwareWin-10\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrg` | Build Path |  |
| 0x008EBD10 | `c:\BWA\N25CFirmwareWin-10\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrc` | Build Path |  |
| 0x008EBD68 | `c:\BWA\N25CFirmwareWin-10\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfro` | Build Path |  |
| 0x008EC0AC | `c:\BWA\N25CFirmwareWin-10\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1` | Build Path |  |
| 0x008FB454 | `c:\BWA\N25CFirmwareWin-10\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttc` | Build Path |  |
| 0x008FB6D0 | `c:\BWA\N25CFirmwareWin-10\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype` | Build Path |  |
| 0x008FBC3C | `c:\BWA\N25CFirmwareWin-10\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1` | Build Path |  |

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
| 0x007D7678 | `Acoustic` | EQ Preset |  |
| 0x007D7684 | `Bass Booster` | EQ Preset |  |
| 0x007D76A4 | `Classical` | EQ Preset |  |
| 0x007D76B0 | `Dance` | EQ Preset |  |
| 0x007D76C0 | `Electronic` | EQ Preset |  |
| 0x007D76D4 | `Hip Hop` | EQ Preset |  |
| 0x007D76DC | `Jazz` | EQ Preset |  |
| 0x007D76E4 | `Latin` | EQ Preset |  |
| 0x007D76EC | `Loudness` | EQ Preset |  |
| 0x007D76F8 | `Lounge` | EQ Preset |  |
| 0x007D7700 | `Piano` | EQ Preset |  |
| 0x007D7710 | `Rock` | EQ Preset |  |
| 0x007D7718 | `Small Speakers` | EQ Preset |  |
| 0x007D7728 | `Spoken Word` | EQ Preset |  |
| 0x007D7734 | `Treble Booster` | EQ Preset |  |
| 0x007D7754 | `Vocal Booster` | EQ Preset |  |
| 0x0081487C | `Acoustic` | EQ Preset |  |
| 0x00814888 | `Bass Booster` | EQ Preset |  |
| 0x008148A8 | `Classical` | EQ Preset |  |
| 0x008148B4 | `Dance` | EQ Preset |  |
| 0x008148C4 | `Electronic` | EQ Preset |  |
| 0x008148D8 | `Hip Hop` | EQ Preset |  |
| 0x008148E0 | `Jazz` | EQ Preset |  |
| 0x008148E8 | `Latin` | EQ Preset |  |
| 0x008148F0 | `Loudness` | EQ Preset |  |
| 0x008148FC | `Lounge` | EQ Preset |  |
| 0x00814904 | `Piano` | EQ Preset |  |
| 0x00814914 | `Rock` | EQ Preset |  |
| 0x0081491C | `Small Speakers` | EQ Preset |  |
| 0x0081492C | `Spoken Word` | EQ Preset |  |
| 0x00814938 | `Treble Booster` | EQ Preset |  |
| 0x00814958 | `Vocal Booster` | EQ Preset |  |
| 0x0081DCF0 | `Acoustic` | EQ Preset |  |
| 0x0081DCFC | `Bass Booster` | EQ Preset |  |
| 0x0081DD1C | `Classical` | EQ Preset |  |
| 0x0081DD28 | `Dance` | EQ Preset |  |
| 0x0081DD38 | `Electronic` | EQ Preset |  |
| 0x0081DD4C | `Hip Hop` | EQ Preset |  |
| 0x0081DD54 | `Jazz` | EQ Preset |  |
| 0x0081DD5C | `Latin` | EQ Preset |  |
| 0x0081DD64 | `Loudness` | EQ Preset |  |
| 0x0081DD70 | `Lounge` | EQ Preset |  |
| 0x0081DD78 | `Piano` | EQ Preset |  |
| 0x0081DD88 | `Rock` | EQ Preset |  |
| 0x0081DD90 | `Small Speakers` | EQ Preset |  |
| 0x0081DDA0 | `Spoken Word` | EQ Preset |  |
| 0x0081DDAC | `Treble Booster` | EQ Preset |  |
| 0x0081DDCC | `Vocal Booster` | EQ Preset |  |
| 0x008276BC | `Acoustic` | EQ Preset |  |
| 0x008276EC | `Dance` | EQ Preset |  |
| 0x008276FC | `Electronic` | EQ Preset |  |
| 0x00827718 | `Jazz` | EQ Preset |  |
| 0x00827720 | `Latin` | EQ Preset |  |
| 0x00827728 | `Loudness` | EQ Preset |  |
| 0x0082773C | `Piano` | EQ Preset |  |
| 0x0082774C | `Rock` | EQ Preset |  |
| 0x0083F5F8 | `Dance` | EQ Preset |  |
| 0x0083F620 | `Hip Hop` | EQ Preset |  |
| 0x0083F628 | `Jazz` | EQ Preset |  |
| 0x0083F638 | `Loudness` | EQ Preset |  |
| 0x0083F644 | `Lounge` | EQ Preset |  |
| 0x0083F64C | `Piano` | EQ Preset |  |
| 0x0083F65C | `Rock` | EQ Preset |  |
| 0x00848B1C | `Jazz` | EQ Preset |  |
| 0x00848B24 | `Latin` | EQ Preset |  |
| 0x00848B38 | `Lounge` | EQ Preset |  |
| 0x00848B40 | `Piano` | EQ Preset |  |
| 0x00848B50 | `Rock` | EQ Preset |  |
| 0x00851F70 | `Hip Hop` | EQ Preset |  |
| 0x00851F78 | `Jazz` | EQ Preset |  |
| 0x00851F94 | `Lounge` | EQ Preset |  |
| 0x00851F9C | `Piano` | EQ Preset |  |
| 0x00851FB4 | `Rock` | EQ Preset |  |
| 0x0085C028 | `Latin` | EQ Preset |  |
| 0x0085C054 | `Rock` | EQ Preset |  |
| 0x008659E8 | `Dance` | EQ Preset |  |
| 0x00865A0C | `Hip Hop` | EQ Preset |  |
| 0x00865A14 | `Jazz` | EQ Preset |  |
| 0x00865A24 | `Loudness` | EQ Preset |  |
| 0x00865A30 | `Lounge` | EQ Preset |  |
| 0x00865A38 | `Piano` | EQ Preset |  |
| 0x00865A48 | `Rock` | EQ Preset |  |
| 0x0087081C | `Acoustic` | EQ Preset |  |
| 0x00870828 | `Bass Booster` | EQ Preset |  |
| 0x00870848 | `Classical` | EQ Preset |  |
| 0x00870854 | `Dance` | EQ Preset |  |
| 0x00870864 | `Electronic` | EQ Preset |  |
| 0x00870878 | `Hip Hop` | EQ Preset |  |
| 0x00870880 | `Jazz` | EQ Preset |  |
| 0x00870888 | `Latin` | EQ Preset |  |
| 0x00870890 | `Loudness` | EQ Preset |  |
| 0x0087089C | `Lounge` | EQ Preset |  |
| 0x008708A4 | `Piano` | EQ Preset |  |
| 0x008708B4 | `Rock` | EQ Preset |  |
| 0x008708BC | `Small Speakers` | EQ Preset |  |
| 0x008708CC | `Spoken Word` | EQ Preset |  |
| 0x008708D8 | `Treble Booster` | EQ Preset |  |
| 0x008708F8 | `Vocal Booster` | EQ Preset |  |
| 0x0087B488 | `Acoustic` | EQ Preset |  |
| 0x0087B494 | `Bass Booster` | EQ Preset |  |
| 0x0087B4B4 | `Classical` | EQ Preset |  |
| 0x0087B4C0 | `Dance` | EQ Preset |  |
| 0x0087B4D0 | `Electronic` | EQ Preset |  |
| 0x0087B4E4 | `Hip Hop` | EQ Preset |  |
| 0x0087B4EC | `Jazz` | EQ Preset |  |
| 0x0087B4F4 | `Latin` | EQ Preset |  |
| 0x0087B4FC | `Loudness` | EQ Preset |  |
| 0x0087B508 | `Lounge` | EQ Preset |  |
| 0x0087B510 | `Piano` | EQ Preset |  |
| 0x0087B520 | `Rock` | EQ Preset |  |
| 0x0087B528 | `Small Speakers` | EQ Preset |  |
| 0x0087B538 | `Spoken Word` | EQ Preset |  |
| 0x0087B544 | `Treble Booster` | EQ Preset |  |
| 0x0087B564 | `Vocal Booster` | EQ Preset |  |
| 0x00885088 | `Dance` | EQ Preset |  |
| 0x008850BC | `Jazz` | EQ Preset |  |
| 0x008850C4 | `Latin` | EQ Preset |  |
| 0x008850CC | `Loudness` | EQ Preset |  |
| 0x008850D8 | `Lounge` | EQ Preset |  |
| 0x008850E0 | `Piano` | EQ Preset |  |
| 0x008850F0 | `Rock` | EQ Preset |  |
| 0x0088E4A4 | `Dance` | EQ Preset |  |
| 0x0088E4D0 | `Jazz` | EQ Preset |  |
| 0x0088E4E0 | `Loudness` | EQ Preset |  |
| 0x0088E4EC | `Lounge` | EQ Preset |  |
| 0x0088E4F4 | `Piano` | EQ Preset |  |
| 0x0088E504 | `Rock` | EQ Preset |  |
| 0x00897B6C | `Hip Hop` | EQ Preset |  |
| 0x00897B74 | `Jazz` | EQ Preset |  |
| 0x00897B98 | `Lounge` | EQ Preset |  |
| 0x00897BB0 | `Rock` | EQ Preset |  |
| 0x008A163C | `Hip Hop` | EQ Preset |  |
| 0x008A1644 | `Jazz` | EQ Preset |  |
| 0x008A1660 | `Lounge` | EQ Preset |  |
| 0x008A1668 | `Piano` | EQ Preset |  |
| 0x008A1678 | `Rock` | EQ Preset |  |
| 0x008B80C4 | `Acoustic` | EQ Preset |  |
| 0x008B80D0 | `Bass Booster` | EQ Preset |  |
| 0x008B80F0 | `Classical` | EQ Preset |  |
| 0x008B80FC | `Dance` | EQ Preset |  |
| 0x008B810C | `Electronic` | EQ Preset |  |
| 0x008B8120 | `Hip Hop` | EQ Preset |  |
| 0x008B8128 | `Jazz` | EQ Preset |  |
| 0x008B8130 | `Latin` | EQ Preset |  |
| 0x008B8138 | `Loudness` | EQ Preset |  |
| 0x008B8144 | `Lounge` | EQ Preset |  |
| 0x008B814C | `Piano` | EQ Preset |  |
| 0x008B815C | `Rock` | EQ Preset |  |
| 0x008B8164 | `Small Speakers` | EQ Preset |  |
| 0x008B8174 | `Spoken Word` | EQ Preset |  |
| 0x008B8180 | `Treble Booster` | EQ Preset |  |
| 0x008B81A0 | `Vocal Booster` | EQ Preset |  |
| 0x008C1744 | `Hip Hop` | EQ Preset |  |
| 0x008C1750 | `Latin` | EQ Preset |  |
| 0x008C1788 | `Rock` | EQ Preset |  |
| 0x008CAF04 | `Acoustic` | EQ Preset |  |
| 0x008CAF10 | `Bass Booster` | EQ Preset |  |
| 0x008CAF30 | `Classical` | EQ Preset |  |
| 0x008CAF3C | `Dance` | EQ Preset |  |
| 0x008CAF4C | `Electronic` | EQ Preset |  |
| 0x008CAF60 | `Hip Hop` | EQ Preset |  |
| 0x008CAF68 | `Jazz` | EQ Preset |  |
| 0x008CAF70 | `Latin` | EQ Preset |  |
| 0x008CAF78 | `Loudness` | EQ Preset |  |
| 0x008CAF84 | `Lounge` | EQ Preset |  |
| 0x008CAF8C | `Piano` | EQ Preset |  |
| 0x008CAF9C | `Rock` | EQ Preset |  |
| 0x008CAFA4 | `Small Speakers` | EQ Preset |  |
| 0x008CAFB4 | `Spoken Word` | EQ Preset |  |
| 0x008CAFC0 | `Treble Booster` | EQ Preset |  |
| 0x008CAFE0 | `Vocal Booster` | EQ Preset |  |
| 0x008D45F4 | `Acoustic` | EQ Preset |  |
| 0x008D4600 | `Bass Booster` | EQ Preset |  |
| 0x008D4620 | `Classical` | EQ Preset |  |
| 0x008D462C | `Dance` | EQ Preset |  |
| 0x008D463C | `Electronic` | EQ Preset |  |
| 0x008D4650 | `Hip Hop` | EQ Preset |  |
| 0x008D4658 | `Jazz` | EQ Preset |  |
| 0x008D4660 | `Latin` | EQ Preset |  |
| 0x008D4668 | `Loudness` | EQ Preset |  |
| 0x008D4674 | `Lounge` | EQ Preset |  |
| 0x008D467C | `Piano` | EQ Preset |  |
| 0x008D468C | `Rock` | EQ Preset |  |
| 0x008D4694 | `Small Speakers` | EQ Preset |  |
| 0x008D46A4 | `Spoken Word` | EQ Preset |  |
| 0x008D46B0 | `Treble Booster` | EQ Preset |  |
| 0x008D46D0 | `Vocal Booster` | EQ Preset |  |

---
