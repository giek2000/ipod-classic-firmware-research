# iPod Nano 3rd Generation - RetailOS 1.1.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.1.3 |
| **IPSW** | iPod_26.1.1.3.ipsw |
| **Device** | iPod Nano 3rd Generation (2009, 4/8GB, Click Wheel, Cover Flow, Video) |
| **UpdaterFamilyID** | 26 |
| **Binary Size** | 10,792,304 bytes (10.29 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,790,256 bytes |
| **Total Strings (>=4)** | 72,489 |
| **Function Prologues** | 22,820 (ARM: 17,473, Thumb: 5,347) |
| **DRAM References** | 105,535 |
| **Peripheral Refs** | 7,435 |
| **Build** | N46FirmwareWin-465 |
| **SoC** | S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Codename** | N46 |
| **DFU PID** | 0x1229 |
| **SHA-256** | `41f3782d9ae5ab8437e30e9f4b26789e959fea2dae2eea7b60fc14272c669877` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000A0B48 | `TSilverCntlr` | Known | Controller |
| 0x000A0B60 | `TCExtrasMenu` | Known | Controller |
| 0x000A0B78 | `TCGameScreen` | Known | Controller |
| 0x000A0B90 | `TCGamesMenu` | Known | Controller |
| 0x000A0BA4 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x000A0BCC | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x000A0BF4 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x000A0C20 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x000A0C44 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x000A0C6C | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x000A0C94 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x000A0CBC | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x000A0CE4 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x000A0D0C | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x000A0D3C | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x000A0D68 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x000A0D98 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x000A0DC0 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x000A0DE8 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x000A0E14 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x000A0E3C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x000A0E64 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x000A0E94 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x000A0EC4 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x000A0FCC | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x000A0FFC | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x000A1024 | `TCRentalNotification` | Known | Controller |
| 0x000A1044 | `TCRentalInfo` | Known | Controller |
| 0x000A105C | `TCRentalConfirmDelete` | Known | Controller |
| 0x000A107C | `TCRentalDispatcher` | Known | Controller |
| 0x000A1098 | `TSilverGlobalCntlr` | Known | Controller |
| 0x000A10B4 | `TSilverTrainerCntlr` | Known | Controller |
| 0x000FA23C | `TCSlideshowLCD` | Known | Controller |
| 0x000FA254 | `TCSlideshowTVOut` | Known | Controller |
| 0x000FA270 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x000FA290 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00123790 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x001237BC | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x001237E8 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x00123810 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x0012383C | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x00123864 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x00123890 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x0012AA98 | `TCRemoteUI` | Known | Controller |
| 0x0012AAAC | `TCUnsupported` | Known | Controller |
| 0x00130E58 | `TCSpeakers` | Known | Controller |
| 0x00130E6C | `TCEQSetting` | Known | Controller |
| 0x0015C990 | `TCSportTimer` | Known | Controller |
| 0x0015C9A8 | `TCSportTimerMenu` | Known | Controller |
| 0x0015C9C4 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x0015C9E8 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x0015DD68 | `TCVoiceMemos` | Known | Controller |
| 0x0015DD80 | `TCVoiceMemosMenu` | Known | Controller |
| 0x0015DD9C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x0015DDBC | `TCVoiceMemosPlayback` | Known | Controller |
| 0x0015DDDC | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x0016F4EC | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x0016F514 | `TCSettings_MainMenu` | Known | Controller |
| 0x0016F530 | `TCSettings_MusicMenu` | Known | Controller |
| 0x0016F550 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x0016F570 | `TCSettings_Brightness` | Known | Controller |
| 0x0016F590 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0016F5B4 | `TCSettings_EQ` | Known | Controller |
| 0x0016F5CC | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x0016F5F4 | `TCSettings_RadioRegions` | Known | Controller |
| 0x0016F614 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0016F638 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x0016F65C | `TCDateTimeScreen` | Known | Controller |
| 0x0016F678 | `TCTimeZoneScreen` | Known | Controller |
| 0x0016F694 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0016F6BC | `TCFirstBoot` | Known | Controller |
| 0x00184F68 | `TCDemoMode` | Known | Controller |
| 0x001B1E0C | `TCAddressViewerMainMenu` | Known | Controller |
| 0x001B1E2C | `TCAddressViewerDetails` | Known | Controller |
| 0x001B1E4C | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x001B1E70 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x001E22E0 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x001E2304 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x001EA2D8 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x00287D0C | `TC_LockDialog` | Known | Controller |
| 0x00287D24 | `TC_LockScreen` | Known | Controller |
| 0x00287D3C | `TC_LockediPod` | Known | Controller |
| 0x00287D54 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x00287D78 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0028DC24 | `TCClock` | Known | Controller |
| 0x0028DC34 | `TCClockCityMenu` | Known | Controller |
| 0x0028DC4C | `TCClockRegionMenu` | Known | Controller |
| 0x0028DC68 | `TCAlarmMenu` | Known | Controller |
| 0x0028DC7C | `TCSleepTimerMenu` | Known | Controller |
| 0x0028DC98 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x0028DCB8 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x0028DCE0 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x0028DD04 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0028DD28 | `TCAlarmDatePicker` | Known | Controller |
| 0x0028DD44 | `TCAlarmTriggered` | Known | Controller |
| 0x00294C64 | `TCNotesDispatcher` | Known | Controller |
| 0x00294C80 | `TCNotesLoading` | Known | Controller |
| 0x00294C98 | `TCNotesList` | Known | Controller |
| 0x00294CAC | `TCNotesContents` | Known | Controller |
| 0x003C7308 | `TCAlarmTriggered` | Known | Controller |
| 0x003C731C | `TSilverCntlr` | Known | Controller |
| 0x003C733C | `TCClock` | Known | Controller |
| 0x003C7344 | `TCClockRegionMenu` | Known | Controller |
| 0x003C7358 | `TCClockCityMenu` | Known | Controller |
| 0x003C7368 | `TCAlarmMenu` | Known | Controller |
| 0x003C7374 | `TCSleepTimerMenu` | Known | Controller |
| 0x003C7388 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003C73A0 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003C73C0 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003C73DC | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003C73F8 | `TCAlarmDatePicker` | Known | Controller |
| 0x003C7430 | `TSilverCntlr` | Known | Controller |
| 0x003C7450 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003C75E0 | `TSilverCntlr` | Known | Controller |
| 0x003C7600 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x003C7620 | `TCSettings_Brightness` | Known | Controller |
| 0x003C7638 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x003C7654 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x003C7674 | `TCSettings_RadioRegions` | Known | Controller |
| 0x003C768C | `TCSettings_EQ` | Known | Controller |
| 0x003C769C | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x003C76B8 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x003C76D8 | `TCFirstBoot` | Known | Controller |
| 0x003C76E4 | `TCSettings_MainMenu` | Known | Controller |
| 0x003C76F8 | `TCSettings_MusicMenu` | Known | Controller |
| 0x003C7710 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003C7728 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x003C7744 | `TCDateTimeScreen` | Known | Controller |
| 0x003C7758 | `TCTimeZoneScreen` | Known | Controller |
| 0x003CE75C | `TSilverCntlr` | Known | Controller |
| 0x003CE77C | `TCClock` | Known | Controller |
| 0x003CE784 | `TCClockRegionMenu` | Known | Controller |
| 0x003CE798 | `TCClockCityMenu` | Known | Controller |
| 0x003CE7A8 | `TCAlarmMenu` | Known | Controller |
| 0x003CE7B4 | `TCSleepTimerMenu` | Known | Controller |
| 0x003CE7C8 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003CE840 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003CE860 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003CE87C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003CE8B0 | `TCAlarmDatePicker` | Known | Controller |
| 0x003CE8C4 | `TCAlarmTriggered` | Known | Controller |
| 0x003D0344 | `TSilverCntlr` | Known | Controller |
| 0x003D0364 | `TC_LockDialog` | Known | Controller |
| 0x003D0374 | `TC_LockScreen` | Known | Controller |
| 0x003D0384 | `TC_LockediPod` | Known | Controller |
| 0x003D0394 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003D03B0 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003D03C8 | `TSilverCntlr` | Known | Controller |
| 0x003D0530 | `TSilverCntlr` | Known | Controller |
| 0x003D054C | `TSilverCntlr` | Known | Controller |
| 0x003D05B0 | `TSilverCntlr` | Known | Controller |
| 0x003D05D0 | `TCNotesDispatcher` | Known | Controller |
| 0x003D05E4 | `TCNotesLoading` | Known | Controller |
| 0x003D05F4 | `TCNotesBase` | Known | Controller |
| 0x003D0600 | `TCNotesList` | Known | Controller |
| 0x003D060C | `TCNotesContents` | Known | Controller |
| 0x003D061C | `TSilverCntlr` | Known | Controller |
| 0x003D063C | `TCRemoteUI` | Known | Controller |
| 0x003D0648 | `TCUnsupported` | Known | Controller |
| 0x003D0658 | `TSilverCntlr` | Known | Controller |
| 0x003D06BC | `TSilverCntlr` | Known | Controller |
| 0x003D06DC | `TCSportTimer` | Known | Controller |
| 0x003D06EC | `TCSportTimerMenu` | Known | Controller |
| 0x003D0700 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x003D071C | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x003D0874 | `TSilverCntlr` | Known | Controller |
| 0x003D0894 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003D08B0 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003D08D0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003D08F0 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003D0918 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003D093C | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003D0964 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003D0984 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003D09A4 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003D09C4 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003D09E4 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003D0A0C | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003D0D68 | `TSilverCntlr` | Known | Controller |
| 0x003D0E90 | `TSilverCntlr` | Known | Controller |
| 0x003D0EB0 | `TCDemoMode` | Known | Controller |
| 0x003D0EBC | `TCClock` | Known | Controller |
| 0x003D0EC4 | `TCClockRegionMenu` | Known | Controller |
| 0x003D0ED8 | `TCClockCityMenu` | Known | Controller |
| 0x003D0EE8 | `TCAlarmMenu` | Known | Controller |
| 0x003D0EF4 | `TCSleepTimerMenu` | Known | Controller |
| 0x003D0F08 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003D0F20 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003D0F40 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003D0F5C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003D0F78 | `TCAlarmDatePicker` | Known | Controller |
| 0x003D0F8C | `TCAlarmTriggered` | Known | Controller |
| 0x003D0FAC | `TSilverCntlr` | Known | Controller |
| 0x003D0FC8 | `TSilverCntlr` | Known | Controller |
| 0x003D0FD8 | `TSilverCntlr` | Known | Controller |
| 0x003D0FF8 | `TCVoiceMemos` | Known | Controller |
| 0x003D1008 | `TCVoiceMemosMenu` | Known | Controller |
| 0x003D101C | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x003D1034 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x003D104C | `TCVoiceMemosPlayback` | Known | Controller |
| 0x003D106C | `TSilverCntlr` | Known | Controller |
| 0x003D10CC | `TSilverCntlr` | Known | Controller |
| 0x003D1138 | `TSilverCntlr` | Known | Controller |
| 0x003D23D4 | `TSilverCntlr` | Known | Controller |
| 0x003D24E0 | `TSilverCntlr` | Known | Controller |
| 0x003DAD64 | `TSilverCntlr` | Known | Controller |
| 0x003DAD84 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x003DAD9C | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x003DADB8 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x003DADD8 | `TCAddressViewerDetails` | Known | Controller |
| 0x003DADF0 | `TSilverCntlr` | Known | Controller |
| 0x003DAE10 | `TSilverCalendarCntlr_Base` | Known | Controller |
| 0x003DAE2C | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x003DAE50 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x003DAE74 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x003DAE94 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x003DAEB8 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x003DAED8 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x003DAEFC | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x003DB0D4 | `TSilverCntlr` | Known | Controller |
| 0x003DB0F4 | `TC_LockDialog` | Known | Controller |
| 0x003DB104 | `TC_LockScreen` | Known | Controller |
| 0x003DB114 | `TC_LockediPod` | Known | Controller |
| 0x003DB124 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003DB148 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003DB1FC | `TSilverCntlr` | Known | Controller |
| 0x003DB21C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003DB238 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003DB258 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003DB278 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003DB2A0 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003DB2C4 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003DB2EC | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003DB30C | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003DB32C | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003DB34C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003DB36C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003DB394 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003DB3BC | `TSilverCntlr` | Known | Controller |
| 0x003DB4DC | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003DB4F8 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003DB518 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003DB538 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003DB560 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003DB584 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003DB5AC | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003DB5CC | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003DB5EC | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003DB60C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003DB62C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003DB654 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003DB67C | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003DB69C | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003DB6BC | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003DB6E0 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003DB700 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003DB724 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003DB74C | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003DB778 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003DB798 | `TCRentalNotification` | Known | Controller |
| 0x003DB7B0 | `TCRentalInfo` | Known | Controller |
| 0x003DB7C0 | `TCRentalConfirmDelete` | Known | Controller |
| 0x003DB7D8 | `TCRentalDispatcher` | Known | Controller |
| 0x003DC0C8 | `TSilverCntlr` | Known | Controller |
| 0x003DC18C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003DC1A8 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003DC1C8 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003DC1E8 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003DC210 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003DC234 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003DC25C | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003DC27C | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003DC29C | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003DC2BC | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003DC2DC | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003DC304 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003DC354 | `TCSlideshowTVOut` | Known | Controller |
| 0x003DC368 | `TCSlideshowLCD` | Known | Controller |
| 0x003DC378 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x003DC390 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x003DC3B0 | `TSilverCntlr` | Known | Controller |
| 0x003DC3DC | `TSilverCntlr` | Known | Controller |
| 0x003DC3FC | `TCUnsupported` | Known | Controller |
| 0x003DC41C | `TSilverCntlr` | Known | Controller |
| 0x003DC45C | `TSilverCntlr` | Known | Controller |
| 0x003DC47C | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x003DC498 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x003DC4B0 | `TSilverCntlr` | Known | Controller |
| 0x003DC4D0 | `TCSpeakers` | Known | Controller |
| 0x003DC4DC | `TCEQSetting` | Known | Controller |
| 0x003DC584 | `TSilverCntlr` | Known | Controller |
| 0x003DC594 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003DC5B0 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003DC5D0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003DC5F0 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003DC618 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003DC63C | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003DC664 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003DC684 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003DC6A4 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003DC6C4 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003DC6E4 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003DC70C | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003DCCB4 | `TSilverCntlr` | Known | Controller |
| 0x003DCCD8 | `TSilverCntlr` | Known | Controller |
| 0x003DCD44 | `TSilverCntlr` | Known | Controller |
| 0x003DCD64 | `TCExtrasMenu` | Known | Controller |
| 0x003DCD74 | `TCGamesMenu` | Known | Controller |
| 0x003DCD80 | `TCGameScreen` | Known | Controller |
| 0x003DCD90 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x003DCDB0 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x003DCDD0 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x003DCDF0 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x003DCE14 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003DCE30 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003DCE50 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003DCE70 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003DCE98 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003DCEBC | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003DCEE4 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003DCF04 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003DCF24 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003DCF44 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003DCF64 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003DCF8C | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003DCFB4 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003DCFD4 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003DCFF4 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003DD018 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003DD038 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003DD05C | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003DD084 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003DD0B0 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003DD0D0 | `TCRentalNotification` | Known | Controller |
| 0x003DD0E8 | `TCRentalInfo` | Known | Controller |
| 0x003DD0F8 | `TCRentalConfirmDelete` | Known | Controller |
| 0x003DD110 | `TCRentalDispatcher` | Known | Controller |
| 0x003DD124 | `TSilverGlobalCntlr` | Known | Controller |
| 0x003DD138 | `TSilverTrainerCntlr` | Known | Controller |
| 0x00467CA0 | `TCCCCCCCCCCCCCCCCCCCCCCT` | Known | Controller |
| 0x007541C6 | `TCNotesDispatcher"` | Known | Controller |
| 0x00754285 | `TCLockChosenDispatcher"` | Known | Controller |
| 0x00754348 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x0075E3AD | `TCNotesDispatcher"` | Known | Controller |
| 0x0075E50F | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x00775B70 | `TCExtrasMenuTCAddressViewerMainMenu` | Known | Controller |
| 0x00775B94 | `TCAddressViewerDetails` | Known | Controller |
| 0x00775BAC | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x00775BC8 | `TCAlarmMenu` | Known | Controller |
| 0x00775BD4 | `TCSleepTimerMenuTCAlarmPropertiesMenu` | Known | Controller |
| 0x00775BFC | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x00775C1C | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x00775C38 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00775C54 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00775C70 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00775C8C | `TCAlarmDatePicker` | Known | Controller |
| 0x00775CA0 | `TCAlarmDatePicker` | Known | Controller |
| 0x00775CB4 | `TCAlarmTriggeredTSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x00775CE0 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x00775D04 | `TSilverCalendarCntlr_MonthViewerTSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x00775D44 | `TSilverCalendarCntlr_EventViewerTSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x00775D84 | `TSilverCalendarCntlr_EventViewerTSilverCntlrTCClockRegionMenu` | Known | Controller |
| 0x00775DC4 | `TCClockCityMenu` | Known | Controller |
| 0x00775DD4 | `TCClockCityMenu` | Known | Controller |
| 0x00775DE4 | `TCClockCityMenu` | Known | Controller |
| 0x00775DF4 | `TCClockCityMenu` | Known | Controller |
| 0x00775E04 | `TCClockCityMenu` | Known | Controller |
| 0x00775E14 | `TCClockCityMenu` | Known | Controller |
| 0x00775E24 | `TCClockCityMenu` | Known | Controller |
| 0x00775E34 | `TCClockCityMenu` | Known | Controller |
| 0x00775E44 | `TCClock` | Known | Controller |
| 0x00775E5C | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDe` | Known | Controller |
| 0x00775EB4 | `TCGamesMenu` | Known | Controller |
| 0x00775EC0 | `TCGameScreenTC_LockediPod` | Known | Controller |
| 0x00775EDC | `TC_LockDialog` | Known | Controller |
| 0x00775EEC | `TC_LockScreen` | Known | Controller |
| 0x00775EFC | `TC_VolumeLimitLockScreenTSilverCntlrTSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x00775F40 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00775F60 | `TSilverMediaListCntlr_AudiobooksTSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x00775FA8 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x00775FC4 | `TSilverMediaListCntlr_AlbumsTSilverMediaListCntlr_Artists` | Known | Controller |
| 0x00776000 | `TSilverMediaListCntlr_GenresTSilverMediaListCntlr_Composers` | Known | Controller |
| 0x0077603C | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0077605C | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x00776084 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x007760A4 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x007760C4 | `TSilverMediaListCntlr_TVEpisodesTSilverMediaListCntlr_MoviesTSilverMediaListCntl` | Known | Controller |
| 0x00776120 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x00776148 | `TSilverMediaListCntlr_VideoPlaylistsTSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x0077618C | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x007761B8 | `TSilverCntlrTSilverMainMediaListCntlr_VideosTSilverSettingsVideoCntlr` | Known | Controller |
| 0x00776200 | `TCFirstBoot` | Known | Controller |
| 0x007762A8 | `TCRentalInfoTCRentalConfirmDelete` | Known | Controller |
| 0x007762CC | `TSilverCntlrTCRentalNotificationTCRentalNotificationTCRentalNotificationTCNotesL` | Known | Controller |
| 0x00776324 | `TCNotesList` | Known | Controller |
| 0x00776330 | `TCNotesList` | Known | Controller |
| 0x0077633C | `TCNotesContents` | Known | Controller |
| 0x0077634C | `TCNotesContents` | Known | Controller |
| 0x0077635C | `TCNotesContents` | Known | Controller |
| 0x0077636C | `TCNotesContents` | Known | Controller |
| 0x00776428 | `TCSlideshowLCD` | Known | Controller |
| 0x00776438 | `TCSlideshowTVOutTCSlideshow_TVOutAskTCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00776488 | `TCRemoteUI` | Known | Controller |
| 0x00776494 | `TCUnsupported` | Known | Controller |
| 0x007764A4 | `TSilverSettingsMenuListCntlrTSilverSettingsMenuListCntlrTSilverSettingsMenuListC` | Known | Controller |
| 0x0077650C | `TCSettings_MusicMenuTCSettings_VolumeLimit` | Known | Controller |
| 0x00776538 | `TCSettings_Brightness` | Known | Controller |
| 0x00776550 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0077656C | `TCSettings_AudiobookSettingsTCSettings_RadioRegions` | Known | Controller |
| 0x007765A0 | `TCSettings_EQ` | Known | Controller |
| 0x007765B0 | `TSilverSettingsMenuListCntlrTSilverCntlrTCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x007765F8 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x00776614 | `TCSettings_MainMenu` | Known | Controller |
| 0x00776628 | `TCSettings_MusicMenuTCSportTimerTCSportTimerMenuTCSportTimerSessionScreen` | Known | Controller |
| 0x0077678C | `TSilverCntlrTTrainerEndSessionCntlr` | Known | Controller |
| 0x00776804 | `TSilverCntlrTTrainerCalibrateWalkMenuCntlr` | Known | Controller |
| 0x00776A98 | `TCVoiceMemosTCVoiceMemosMainMenuTCVoiceMemosMainMenuTCVoiceMemosMainMenuTSearchC` | Known | Controller |
| 0x00776AF8 | `TCEQSetting` | Known | Controller |
| 0x00776BA6 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x00777EA9 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x0077DAB2 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077DB10 | `TCNotesDispatcher` | Known | Controller |
| 0x0077F6EE | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077F74C | `TCNotesDispatcher` | Known | Controller |
| 0x0078132A | `TCLockChosenDispatcher` | Known | Controller |
| 0x00781388 | `TCNotesDispatcher` | Known | Controller |
| 0x00782F66 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00782FC4 | `TCNotesDispatcher` | Known | Controller |
| 0x00784BA2 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00784C00 | `TCNotesDispatcher` | Known | Controller |
| 0x007867DE | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078683C | `TCNotesDispatcher` | Known | Controller |
| 0x0078841A | `TCLockChosenDispatcher` | Known | Controller |
| 0x00788478 | `TCNotesDispatcher` | Known | Controller |
| 0x0078A056 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078A0B4 | `TCNotesDispatcher` | Known | Controller |
| 0x0078BC92 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078BCF0 | `TCNotesDispatcher` | Known | Controller |
| 0x0078D8CE | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078D92C | `TCNotesDispatcher` | Known | Controller |
| 0x0078F50A | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078F568 | `TCNotesDispatcher` | Known | Controller |
| 0x00791146 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007911A4 | `TCNotesDispatcher` | Known | Controller |
| 0x00792D82 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00792DE0 | `TCNotesDispatcher` | Known | Controller |
| 0x007949BE | `TCLockChosenDispatcher` | Known | Controller |
| 0x00794A1C | `TCNotesDispatcher` | Known | Controller |
| 0x007965FA | `TCLockChosenDispatcher` | Known | Controller |
| 0x00796658 | `TCNotesDispatcher` | Known | Controller |
| 0x00798236 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00798294 | `TCNotesDispatcher` | Known | Controller |
| 0x00799E72 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00799ED0 | `TCNotesDispatcher` | Known | Controller |
| 0x0079BAAE | `TCLockChosenDispatcher` | Known | Controller |
| 0x0079BB0C | `TCNotesDispatcher` | Known | Controller |
| 0x0079D6EA | `TCLockChosenDispatcher` | Known | Controller |
| 0x0079D748 | `TCNotesDispatcher` | Known | Controller |
| 0x0079F326 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0079F384 | `TCNotesDispatcher` | Known | Controller |
| 0x007A0F62 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A0FC0 | `TCNotesDispatcher` | Known | Controller |
| 0x007A2B9E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A2BFC | `TCNotesDispatcher` | Known | Controller |
| 0x007A47DA | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A4838 | `TCNotesDispatcher` | Known | Controller |
| 0x007A6416 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A6474 | `TCNotesDispatcher` | Known | Controller |
| 0x007A8052 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A80B0 | `TCNotesDispatcher` | Known | Controller |
| 0x007A9C8E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A9CEC | `TCNotesDispatcher` | Known | Controller |
| 0x007AB8CA | `TCLockChosenDispatcher` | Known | Controller |
| 0x007AB928 | `TCNotesDispatcher` | Known | Controller |
| 0x007AD506 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007AD564 | `TCNotesDispatcher` | Known | Controller |
| 0x007AF142 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007AF1A0 | `TCNotesDispatcher` | Known | Controller |
| 0x007B0D7E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B0DDC | `TCNotesDispatcher` | Known | Controller |
| 0x007B29BA | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B2A18 | `TCNotesDispatcher` | Known | Controller |
| 0x007B45F6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B4654 | `TCNotesDispatcher` | Known | Controller |
| 0x007B6232 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B6290 | `TCNotesDispatcher` | Known | Controller |
| 0x007B7E6E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B7ECC | `TCNotesDispatcher` | Known | Controller |
| 0x007B9AAA | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B9B08 | `TCNotesDispatcher` | Known | Controller |
| 0x007BB6E6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007BB744 | `TCNotesDispatcher` | Known | Controller |
| 0x007BD322 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007BD380 | `TCNotesDispatcher` | Known | Controller |
| 0x007C8F58 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x007C921A | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x007C9A50 | `TCRentalDispatcher` | Known | Controller |
| 0x007CA308 | `TCRentalDispatcher` | Known | Controller |
| 0x007CABC0 | `TCRentalDispatcher` | Known | Controller |
| 0x007CB478 | `TCRentalDispatcher` | Known | Controller |
| 0x007CBD30 | `TCRentalDispatcher` | Known | Controller |
| 0x007CC5E8 | `TCRentalDispatcher` | Known | Controller |
| 0x007CCEA0 | `TCRentalDispatcher` | Known | Controller |
| 0x007CD758 | `TCRentalDispatcher` | Known | Controller |
| 0x00920974 | `TCMockupModeNavScreen` | Known | Controller |
| 0x0092098C | `TSilverCntlr` | Known | Controller |
| 0x009209AC | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x009209E4 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x00920A04 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00920A24 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x00920A48 | `TCExtrasMenu` | Known | Controller |
| 0x00920B58 | `TSilverCntlr` | Known | Controller |
| 0x00920B78 | `TCSlideshowTVOut` | Known | Controller |
| 0x00920B8C | `TCSlideshowLCD` | Known | Controller |
| 0x00920B9C | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x00920BB4 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00920BF0 | `TSilverCntlr` | Known | Controller |
| 0x00920C6C | `TCSlideshowTVOut` | Known | Controller |
| 0x00920C80 | `TCSlideshowLCD` | Known | Controller |
| 0x00920C90 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x00920CA8 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00920CC8 | `TSilverCntlr` | Known | Controller |
| 0x00920F80 | `TSilverCntlr` | Known | Controller |
| 0x00920FA0 | `TCGamesMenu` | Known | Controller |
| 0x00920FAC | `TCGameScreen` | Known | Controller |
| 0x009DF2CF | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |

---

## 2. Settings (Toggle/Show/TC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0013A158 | `ShowSetting_EQ` | Known | User setting |
| 0x001EBA8C | `ToggleSetting_Repeat` | Known | User setting |
| 0x001EBAA8 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001EBAC0 | `ToggleSetting_TVOut` | Known | User setting |
| 0x001EBAD4 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x0021B8D8 | `ShowSetting_Backlight` | Known | User setting |
| 0x00230680 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x0023069C | `ToggleSetting_Repeat` | Known | User setting |
| 0x002306B4 | `ToggleSetting_SortBy` | Known | User setting |
| 0x002306CC | `ToggleSetting_ClassicUI` | Known | User setting |
| 0x002306E4 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x00230700 | `ToggleSetting_Clicker` | Known | User setting |
| 0x00230718 | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x00230738 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x00230754 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x00230770 | `ShowSetting_Shuffle` | Known | User setting |
| 0x0023091C | `ShowSetting_Repeat` | Known | User setting |
| 0x00230930 | `ShowSetting_About` | Known | User setting |
| 0x00230944 | `ShowSetting_MainMenu` | Known | User setting |
| 0x0023095C | `ShowSetting_MusicMenu` | Known | User setting |
| 0x00230974 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x0023098C | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x002309A8 | `ShowSetting_Brightness` | Known | User setting |
| 0x002309C0 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x002309D8 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x002309F4 | `ShowSetting_EQ` | Known | User setting |
| 0x00230A04 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x00230BA0 | `ShowSetting_Clicker` | Known | User setting |
| 0x00230BB4 | `ShowSetting_DateAndTime` | Known | User setting |
| 0x00230BCC | `ShowSetting_SortBy` | Known | User setting |
| 0x00230BE0 | `ShowSetting_ClassicUI` | Known | User setting |
| 0x00230BF8 | `ShowSetting_Language` | Known | User setting |
| 0x00230C10 | `ShowSetting_Legal` | Known | User setting |
| 0x00230C24 | `ShowSetting_ResetAll` | Known | User setting |
| 0x0075D235 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x0075D2E5 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x0075F97A | `ShowSetting_About` | Known | User setting |
| 0x0075FA1C | `ToggleSetting_Shuffle` | Known | User setting |
| 0x0075FA60 | `ShowSetting_Shuffle` | Known | User setting |
| 0x0075FAD7 | `ToggleSetting_Repeat` | Known | User setting |
| 0x0075FB1A | `ShowSetting_Repeat` | Known | User setting |
| 0x0075FC24 | `ShowSetting_MainMenu` | Known | User setting |
| 0x0075FD34 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x0075FDFC | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x0075FEC6 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x0075FFDE | `ShowSetting_Brightness` | Known | User setting |
| 0x00760114 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x00760225 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x00760326 | `ShowSetting_EQ` | Known | User setting |
| 0x00760393 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x007603DA | `ShowSetting_SoundCheck` | Known | User setting |
| 0x00760457 | `ToggleSetting_Clicker` | Known | User setting |
| 0x0076049B | `ShowSetting_Clicker` | Known | User setting |
| 0x00760602 | `ToggleSetting_SortBy` | Known | User setting |
| 0x00760645 | `ShowSetting_SortBy` | Known | User setting |
| 0x00760746 | `ShowSetting_Language` | Known | User setting |
| 0x00760856 | `ShowSetting_Legal` | Known | User setting |
| 0x00760987 | `ShowSetting_ResetAll` | Known | User setting |
| 0x00760AF8 | `ShowSetting_Backlight` | Known | User setting |
| 0x00760BA8 | `ShowSetting_Backlight` | Known | User setting |
| 0x00760C58 | `ShowSetting_Backlight` | Known | User setting |
| 0x00760D09 | `ShowSetting_Backlight` | Known | User setting |
| 0x00760DBA | `ShowSetting_Backlight` | Known | User setting |
| 0x00760E6B | `ShowSetting_Backlight` | Known | User setting |
| 0x00760F1F | `ShowSetting_Backlight` | Known | User setting |
| 0x00760FCE | `ShowSetting_EQ` | Known | User setting |
| 0x00761043 | `ShowSetting_Language` | Known | User setting |
| 0x007DE86C | `ToggleSetting_Repeat` | Known | User setting |
| 0x007DE8A6 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x007DE968 | `ToggleSetting_TVOut` | Known | User setting |
| 0x007DE9A1 | `ToggleSetting_TVSignal` | Known | User setting |

---

## 3. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00158808 | `MockupMode/MockupMode.xml` | Hidden | Developer Tool |
| 0x00158D08 | `MockupMode/` | Hidden | Developer Tool |
| 0x0026CC5C | `Channel UnitTests` | Hidden | Developer Tool |
| 0x002C3141 | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x002C3184 | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x002C3199 | `RTXCbug> ` | Hidden | Developer Tool |
| 0x002C3B75 | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x002D5370 | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x0036F9DD | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x0036FAA5 | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x003CC709 | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x00776A04 | `TTrainerLoadingCntlrTSilverCntlrTUnitTestSuiteCntlr` | Hidden | Developer Tool |
| 0x00776A38 | `TUnitTestSuiteTestsCntlrTSilverCntlrTCVoiceMemosMainMenuTCVoiceMemosMenuTCVoiceM` | Hidden | Developer Tool |
| 0x00815518 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0085D9E8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0086FFF4 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00887754 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008999C0 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008A35CC | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008ACEFC | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008C1F0C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008CBA98 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008F1F0C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0091038C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00919694 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x009CFBCE | `Debug_ListItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x009CFBE6 | `Debug_MenuItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x009D07FB | `DemoMode_Background_Image` | Hidden | Demo/Retail Mode |
| 0x009D180E | `DemoMode_Icon_Image` | Hidden | Demo/Retail Mode |
| 0x009D3688 | `Debug_Menu_Title` | Hidden | Debug/Diagnostic |
| 0x009D36AD | `UnitTest_Menu_Title` | Hidden | Developer Tool |
| 0x009DC79C | `UnitTestModel` | Hidden | Developer Tool |
| 0x009DD3D4 | `UnitTest_ListItem` | Hidden | Developer Tool |
| 0x009DE966 | `DemoMode_View_Main` | Hidden | Demo/Retail Mode |
| 0x009DEB4E | `DemoMode_Icon` | Hidden | Demo/Retail Mode |
| 0x009DF994 | `UnitTestApp` | Hidden | Developer Tool |
| 0x009DFFF6 | `Debug_ListItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009E0011 | `Debug_MenuItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009E081C | `DemoMode_Text_Color` | Hidden | Demo/Retail Mode |
| 0x009E0C25 | `Debug_ListItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009E0C3C | `Debug_MenuItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009E56E1 | `Debug_ListItem_UnitTest` | Hidden | Developer Tool |
| 0x009E56F9 | `Debug_MenuItem_UnitTest` | Hidden | Developer Tool |
| 0x009EA627 | `Debug_ListItem_Memory` | Hidden | Debug/Diagnostic |
| 0x009EA63D | `Debug_MenuItem_Memory` | Hidden | Debug/Diagnostic |

---

## 4. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000677B | `"MeCCADecode` | Known | Audio system |
| 0x0014EDF0 | `AudioCodecs` | Known | Audio system |
| 0x00193C98 | `MeCCA_PCM_Output.wav` | Known | Audio system |
| 0x001B1048 | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x001BB954 | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x001BBB5C | `MeCCAVideoDecode` | Known | Audio system |
| 0x0092E308 | `MeCCA_StreamCache` | Known | Audio system |

---

## 5. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F5FA8 | `HandleWheel` | Known | Event handler |
| 0x000F5FB4 | `HandlePlayPause` | Known | Event handler |
| 0x000F5FC4 | `HandleSelectDown` | Known | Event handler |
| 0x000F5FD8 | `HandleNext` | Known | Event handler |
| 0x000F5FE4 | `HandlePrevious` | Known | Event handler |
| 0x000F5FF4 | `HandleNextPushAndHold` | Known | Event handler |
| 0x000F600C | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000F62A4 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000F62C4 | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x00102310 | `HandleSelect` | Known | Event handler |
| 0x00102324 | `HandleHilite` | Known | Event handler |
| 0x001026BC | `HandleEQSettingSelected` | Known | Event handler |
| 0x00102AEC | `HandleSelect` | Known | Event handler |
| 0x00102B00 | `HandleGameHilited` | Known | Event handler |
| 0x00102DB0 | `HandleNotesSelected` | Known | Event handler |
| 0x00102DC8 | `HandleNotesPop` | Known | Event handler |
| 0x00102DD8 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00110E68 | `HandleVolumeWheel` | Known | Event handler |
| 0x00110E7C | `HandleVolumeChange` | Known | Event handler |
| 0x00110E90 | `HandleTimerDone` | Known | Event handler |
| 0x00110EA0 | `HandleFrequencyChange` | Known | Event handler |
| 0x00110F18 | `HandleTuning` | Known | Event handler |
| 0x00110F28 | `HandleTuningSelect` | Known | Event handler |
| 0x001216F8 | `HandleLock` | Known | Event handler |
| 0x00121708 | `HandleAddressBook` | Known | Event handler |
| 0x00121DF0 | `HandleSelect` | Known | Event handler |
| 0x00122328 | `HandleExit` | Known | Event handler |
| 0x00122338 | `HandleLap` | Known | Event handler |
| 0x00122344 | `HandleResume` | Known | Event handler |
| 0x00122354 | `HandleStartStop` | Known | Event handler |
| 0x001225DC | `HandleWheel` | Known | Event handler |
| 0x001225EC | `HandlePlayPause` | Known | Event handler |
| 0x001225FC | `HandleSelectDown` | Known | Event handler |
| 0x00122610 | `HandleHilite` | Known | Event handler |
| 0x0012C250 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x0013A38C | `HandleExitUnsupported` | Known | Event handler |
| 0x0014599C | `HandleBasicSelected` | Known | Event handler |
| 0x001459B4 | `HandleTimedWorkoutSelected` | Known | Event handler |
| 0x001459D0 | `HandleDistanceWorkoutSelected` | Known | Event handler |
| 0x001459F0 | `HandleCaloriesWorkoutSelected` | Known | Event handler |
| 0x00145A10 | `HandleSelectWorkout` | Known | Event handler |
| 0x00154574 | `HandleNotesPop` | Known | Event handler |
| 0x00154588 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0015546C | `HandleSelect` | Known | Event handler |
| 0x00155480 | `HandleWheel` | Known | Event handler |
| 0x0015548C | `HandleImageNext` | Known | Event handler |
| 0x0015549C | `HandleImagePrev` | Known | Event handler |
| 0x001554AC | `HandleImageLast` | Known | Event handler |
| 0x001554BC | `HandleImageFirst` | Known | Event handler |
| 0x001554D0 | `HandlePlayPause` | Known | Event handler |
| 0x001554E0 | `HandlePlay` | Known | Event handler |
| 0x001554EC | `HandlePause` | Known | Event handler |
| 0x00169FAC | `HandleSelectCity` | Known | Event handler |
| 0x00169FC4 | `HandleHighlightCity` | Known | Event handler |
| 0x0016AEEC | `HandleWantPopFlow` | Known | Event handler |
| 0x0016AF04 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x0016AF20 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0016AF3C | `HandleFlowNext` | Known | Event handler |
| 0x0016AF4C | `HandleFlowPrev` | Known | Event handler |
| 0x0016AF5C | `HandleFlowWheel` | Known | Event handler |
| 0x0016AF6C | `HandleAlbumSelected` | Known | Event handler |
| 0x0016AF80 | `HandlePlayPause` | Known | Event handler |
| 0x0016AF90 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00195B34 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00195F24 | `HandleSelect` | Known | Event handler |
| 0x00196DE4 | `HandleSelect` | Known | Event handler |
| 0x00196DF8 | `HandleWheel` | Known | Event handler |
| 0x00196E04 | `HandleImageNext` | Known | Event handler |
| 0x00196E14 | `HandleImagePrev` | Known | Event handler |
| 0x00196E24 | `HandleImageLast` | Known | Event handler |
| 0x00196E34 | `HandleImageFirst` | Known | Event handler |
| 0x00196E48 | `HandlePlayPause` | Known | Event handler |
| 0x00196E58 | `HandlePlay` | Known | Event handler |
| 0x00196E64 | `HandlePause` | Known | Event handler |
| 0x00197304 | `HandleNew` | Known | Event handler |
| 0x00197314 | `HandleClear` | Known | Event handler |
| 0x00197320 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x0019733C | `HandleSelectIndexedSession` | Known | Event handler |
| 0x0019764C | `HandleWheel` | Known | Event handler |
| 0x0019765C | `HandleArrowUp` | Known | Event handler |
| 0x0019766C | `HandleArrowDown` | Known | Event handler |
| 0x00199890 | `HandleHiliteAlbum` | Known | Event handler |
| 0x001998A8 | `HandleBrowseAlbum` | Known | Event handler |
| 0x001998BC | `HandlePlayPause` | Known | Event handler |
| 0x001B56DC | `HandleSelect` | Known | Event handler |
| 0x001B586C | `HandleSelectRegion` | Known | Event handler |
| 0x001BA2F8 | `HandleChooseLink` | Known | Event handler |
| 0x001BA310 | `HandleChooseCalibrate` | Known | Event handler |
| 0x001BA328 | `HandleUnlink` | Known | Event handler |
| 0x001CB2D0 | `HandleImageWheel` | Known | Event handler |
| 0x001CB2E8 | `HandlePlayPause` | Known | Event handler |
| 0x001CB2F8 | `HandleBrowseLarge` | Known | Event handler |
| 0x001CB30C | `HandleBrowseSmall` | Known | Event handler |
| 0x001CB320 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x001CB338 | `HandleImageNext` | Known | Event handler |
| 0x001CB348 | `HandleImagePrev` | Known | Event handler |
| 0x001CB358 | `HandleHilite` | Known | Event handler |
| 0x001CB368 | `HandleImageLast` | Known | Event handler |
| 0x001CB378 | `HandleImageFirst` | Known | Event handler |
| 0x001CB38C | `HandleScreenNext` | Known | Event handler |
| 0x001CB3A0 | `HandleScreenPrev` | Known | Event handler |
| 0x001CDC84 | `HandlePlayPause` | Known | Event handler |
| 0x001CDC98 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x001CDCB4 | `HandleNext` | Known | Event handler |
| 0x001CDCC0 | `HandleNextPressAndHold` | Known | Event handler |
| 0x001CDCD8 | `HandlePrevious` | Known | Event handler |
| 0x001CDCE8 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x001CDD04 | `HandleRemotePlayPause` | Known | Event handler |
| 0x001CDD1C | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x001CDD40 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x001CDD58 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x001CDD70 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x001CDF40 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x001CDF58 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x001CDF70 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x001CDF8C | `HandleRemoteStop` | Known | Event handler |
| 0x001CDFA0 | `HandleRemotePlay` | Known | Event handler |
| 0x001CDFB4 | `HandleRemotePause` | Known | Event handler |
| 0x001CDFC8 | `HandleRemoteMute` | Known | Event handler |
| 0x001CDFDC | `HandleRemoteNextChapter` | Known | Event handler |
| 0x001CDFF4 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x001CE00C | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x001CE028 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x001CE24C | `HandleRemoteShuffle` | Known | Event handler |
| 0x001CE260 | `HandleRemoteRepeat` | Known | Event handler |
| 0x001CE274 | `HandleRemoteOn` | Known | Event handler |
| 0x001CE284 | `HandleRemoteOff` | Known | Event handler |
| 0x001CE294 | `HandleRemoteBacklight` | Known | Event handler |
| 0x001CE2AC | `HandleRemoteFFDown` | Known | Event handler |
| 0x001CE2C0 | `HandleRemoteFFUp` | Known | Event handler |
| 0x001CE2D4 | `HandleRemoteRewDown` | Known | Event handler |
| 0x001CE2E8 | `HandleRemoteRewUp` | Known | Event handler |
| 0x001CE2FC | `HandleRemoteMenuDown` | Known | Event handler |
| 0x001CE314 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x001CE328 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x001CE340 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x001CE510 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x001CE528 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x001CE540 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x001CE55C | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x001CE574 | `HandleRemoteEvent` | Known | Event handler |
| 0x001CE588 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x001CE5A4 | `HandleAudioPlayPause` | Known | Event handler |
| 0x001CE5BC | `HandleAudioNext` | Known | Event handler |
| 0x001CE5CC | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x001CE5E8 | `HandleAudioPrevious` | Known | Event handler |
| 0x001CE5FC | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x001CE7FC | `HandleAudioNextAlbum` | Known | Event handler |
| 0x001CE814 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x001CE82C | `HandleAudioVolumeDown` | Known | Event handler |
| 0x001CE844 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x001CE858 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x001CE870 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x001CE888 | `HandleAudioStop` | Known | Event handler |
| 0x001CE898 | `HandleAudioPlay` | Known | Event handler |
| 0x001CE8A8 | `HandleAudioPause` | Known | Event handler |
| 0x001CE8BC | `HandleAudioMute` | Known | Event handler |
| 0x001CE8CC | `HandleAudioNextChapter` | Known | Event handler |
| 0x001CE8E4 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x001CEB04 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x001CEB1C | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x001CEB34 | `HandleAudioShuffle` | Known | Event handler |
| 0x001CEB48 | `HandleAudioRepeat` | Known | Event handler |
| 0x001CEB5C | `HandleAudioFFDown` | Known | Event handler |
| 0x001CEB70 | `HandleAudioFFUp` | Known | Event handler |
| 0x001CEB80 | `HandleAudioRewDown` | Known | Event handler |
| 0x001CEB94 | `HandleAudioRewUp` | Known | Event handler |
| 0x001CEBA8 | `HandleVideoPlayPause` | Known | Event handler |
| 0x001CEBC0 | `HandleVideoNext` | Known | Event handler |
| 0x001CEBD0 | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x001CEBEC | `HandleVideoPrevious` | Known | Event handler |
| 0x001CEC00 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x001CEE08 | `HandleVideoStop` | Known | Event handler |
| 0x001CEE18 | `HandleVideoPlay` | Known | Event handler |
| 0x001CEE28 | `HandleVideoPause` | Known | Event handler |
| 0x001CEE3C | `HandleVideoFFDown` | Known | Event handler |
| 0x001CEE50 | `HandleVideoFFUp` | Known | Event handler |
| 0x001CEE60 | `HandleVideoRewDown` | Known | Event handler |
| 0x001CEE74 | `HandleVideoRewUp` | Known | Event handler |
| 0x001CEE88 | `HandleVideoNextChapter` | Known | Event handler |
| 0x001CEEA0 | `HandleVideoPrevChapter` | Known | Event handler |
| 0x001CEEB8 | `HandleVideoNextFrame` | Known | Event handler |
| 0x001CEED0 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x001CEEE8 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x001D2480 | `HandleSelect` | Known | Event handler |
| 0x001D2494 | `HandleMenu` | Known | Event handler |
| 0x001D24A0 | `HandleLinkCancelOption` | Known | Event handler |
| 0x001D24B8 | `HandleLinkNewRemote` | Known | Event handler |
| 0x001D24CC | `HandleCancelRemoteLinking` | Known | Event handler |
| 0x001D282C | `HandleChoosePlaylistsPlay` | Known | Event handler |
| 0x001D284C | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x001D2868 | `HandleNoneSelected` | Known | Event handler |
| 0x001D287C | `HandleNowPlayingSelected` | Known | Event handler |
| 0x001D2898 | `HandleMenuKeyNop` | Known | Event handler |
| 0x001D28AC | `HandlePlaylistSelected` | Known | Event handler |
| 0x001D3078 | `HandlePauseWorkout` | Known | Event handler |
| 0x001D3090 | `HandleEndWorkout` | Known | Event handler |
| 0x001D30A4 | `HandleResumeWorkout` | Known | Event handler |
| 0x001D30B8 | `HandleChooseMusic` | Known | Event handler |
| 0x001D30CC | `HandleMenuKeyNop` | Known | Event handler |
| 0x001DEF08 | `HandleMainMenu` | Known | Event handler |
| 0x001E344C | `HandlePowerSongSelected` | Known | Event handler |
| 0x001E3468 | `HandlePowerSongChosen` | Known | Event handler |
| 0x001E3480 | `HandleSelectTimedWorkout` | Known | Event handler |
| 0x001E3D04 | `HandleChooseRemoteSetting` | Known | Event handler |
| 0x001E3D24 | `HandleChooseRemoteLink` | Known | Event handler |
| 0x001E3D3C | `HandleChooseRemoteUnlink` | Known | Event handler |
| 0x001E4078 | `HandleSelectResume` | Known | Event handler |
| 0x001E4090 | `HandleEndWorkout` | Known | Event handler |
| 0x001EA1F0 | `HandleSelect` | Known | Event handler |
| 0x001EA498 | `HandleMusicMenu` | Known | Event handler |
| 0x001EA758 | `HandleSelect` | Known | Event handler |
| 0x001EAADC | `HandleDeleteAllSelect` | Known | Event handler |
| 0x001EAAF4 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x001EAB14 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x001EAB38 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x001EAB54 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x001EAFF0 | `HandleWheel` | Known | Event handler |
| 0x001EB000 | `HandlePlayPause` | Known | Event handler |
| 0x001EB010 | `HandleSelectDown` | Known | Event handler |
| 0x001EB024 | `HandleNext` | Known | Event handler |
| 0x001EB030 | `HandlePrevious` | Known | Event handler |
| 0x001EB040 | `HandleNextPushAndHold` | Known | Event handler |
| 0x001EB058 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x001F086C | `HandleChooseLast` | Known | Event handler |
| 0x001F0884 | `HandleChooseRecent` | Known | Event handler |
| 0x001F0898 | `HandleChooseWorkout` | Known | Event handler |
| 0x001F08AC | `HandleChooseBest` | Known | Event handler |
| 0x001F08C0 | `HandleChooseRecentWorkout` | Known | Event handler |
| 0x001F2F48 | `HandleSelect` | Known | Event handler |
| 0x001F2F5C | `HandleMenu` | Known | Event handler |
| 0x001FAF64 | `HandleFrequencyChosen` | Known | Event handler |
| 0x001FAF7C | `HandleDateChosen` | Known | Event handler |
| 0x001FAF90 | `HandleTimeChosen` | Known | Event handler |
| 0x001FAFA4 | `HandleSoundChosen` | Known | Event handler |
| 0x001FAFB8 | `HandleLabelChosen` | Known | Event handler |
| 0x001FAFCC | `HandleDeleteChosen` | Known | Event handler |
| 0x001FC0AC | `HandleSelect` | Known | Event handler |
| 0x002009C8 | `HandlePrev` | Known | Event handler |
| 0x002009D8 | `HandleNext` | Known | Event handler |
| 0x002009E4 | `HandlePlayPause` | Known | Event handler |
| 0x002011C0 | `HandleChoosePowerPlay` | Known | Event handler |
| 0x002011DC | `HandleChooseUnit` | Known | Event handler |
| 0x002011F0 | `HandleChooseVoicePrompts` | Known | Event handler |
| 0x00209B68 | `HandleNextContact` | Known | Event handler |
| 0x00209B80 | `HandlePreviousContact` | Known | Event handler |
| 0x0020CF64 | `HandleSelect` | Known | Event handler |
| 0x0020D240 | `HandleListChoose` | Known | Event handler |
| 0x00211C74 | `HandleItemSelected` | Known | Event handler |
| 0x00211E6C | `HandleRadioRegion` | Known | Event handler |
| 0x00212054 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x002128A0 | `HandleSelect` | Known | Event handler |
| 0x00212CE8 | `HandlePauseKey` | Known | Event handler |
| 0x00212CFC | `HandlePauseHold` | Known | Event handler |
| 0x00212D0C | `HandlePauseKeyNop` | Known | Event handler |
| 0x00212D20 | `HandleMenuKey` | Known | Event handler |
| 0x00212D30 | `HandleMenuKeyNop` | Known | Event handler |
| 0x00212D44 | `HandleWheel` | Known | Event handler |
| 0x00212D94 | `HandleSelectKeyDown` | Known | Event handler |
| 0x00212DA8 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x00212DC0 | `HandlePowerPlay` | Known | Event handler |
| 0x00217B90 | `HandlePlayPause` | Known | Event handler |
| 0x00218DFC | `HandleSelect` | Known | Event handler |
| 0x0021908C | `HandleChooseCustomTimedWorkout` | Known | Event handler |
| 0x002190B0 | `HandleChooseTimedWorkoutFromList` | Known | Event handler |
| 0x002190D4 | `HandleChooseCustomDistanceWorkout` | Known | Event handler |
| 0x002190F8 | `HandleChooseDistanceWorkoutFromList` | Known | Event handler |
| 0x0021911C | `HandleChooseCustomCalorieWorkout` | Known | Event handler |
| 0x00219140 | `HandleChooseCalorieWorkoutFromList` | Known | Event handler |
| 0x0021BBB4 | `HandleDelete` | Known | Event handler |
| 0x0021BBC8 | `HandleSelectLozinch` | Known | Event handler |
| 0x0021BE70 | `HandleSelect` | Known | Event handler |
| 0x0021C13C | `HandleTVOutChanged` | Known | Event handler |
| 0x0021C154 | `HandleTVSignalChanged` | Known | Event handler |
| 0x0021C16C | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x0021C18C | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x0021C1AC | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x0021C1D0 | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x0021C1F0 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x0021CADC | `HandleBegin` | Known | Event handler |
| 0x0021FD60 | `HandleSelectKey` | Known | Event handler |
| 0x0021FF08 | `HandleSelect` | Known | Event handler |
| 0x00220C84 | `HandlePlayPause` | Known | Event handler |
| 0x00220C98 | `HandleWheel` | Known | Event handler |
| 0x00220CA4 | `HandleWheelRating` | Known | Event handler |
| 0x00220CB8 | `HandleWheelScrub` | Known | Event handler |
| 0x00220CCC | `HandleWheelVolume` | Known | Event handler |
| 0x00220D8C | `HandleMenuKey` | Known | Event handler |
| 0x00220DF8 | `HandleMenuLongpress` | Known | Event handler |
| 0x00220E0C | `HandleRentalWarningChoice` | Known | Event handler |
| 0x00221A14 | `HandleSelect` | Known | Event handler |
| 0x002222E4 | `HandleLeaveAlarm` | Known | Event handler |
| 0x002231D4 | `HandleSelect` | Known | Event handler |
| 0x002231E8 | `HandleHilite` | Known | Event handler |
| 0x002231F8 | `HandlePlayPause` | Known | Event handler |
| 0x00223208 | `HandleAddToOTG` | Known | Event handler |
| 0x00223218 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00224254 | `HandleWeightWheel` | Known | Event handler |
| 0x0022426C | `HandleWeightSelect` | Known | Event handler |
| 0x00224280 | `HandleDistanceWheel` | Known | Event handler |
| 0x00224294 | `HandleDistanceSelect` | Known | Event handler |
| 0x002242AC | `HandleTimeWheel` | Known | Event handler |
| 0x002242BC | `HandleTimeSelect` | Known | Event handler |
| 0x002242D0 | `HandleCaloriesWheel` | Known | Event handler |
| 0x002242E4 | `HandleCaloriesSelect` | Known | Event handler |
| 0x002248B0 | `HandleSelect` | Known | Event handler |
| 0x002248C4 | `HandleSelectFromLastWorkout` | Known | Event handler |
| 0x00227178 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x00227988 | `HandleSelect` | Known | Event handler |
| 0x0022799C | `HandleWheel` | Known | Event handler |
| 0x002279A8 | `HandleWheelProgress` | Known | Event handler |
| 0x002279BC | `HandleSelectProgress` | Known | Event handler |
| 0x002279D4 | `HandleSelectVolume` | Known | Event handler |
| 0x002279E8 | `HandleSelectScrub` | Known | Event handler |
| 0x002279FC | `HandleSelectRating` | Known | Event handler |
| 0x00227A10 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x00227A28 | `HandleSelectChapterArt` | Known | Event handler |
| 0x00227A40 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x00227A5C | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x00227A78 | `HandleWheelBrightness` | Known | Event handler |
| 0x00227BC0 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x002299D4 | `HandleSelect` | Known | Event handler |
| 0x002299E4 | `HandleSelectRating` | Known | Event handler |
| 0x002299F8 | `HandleSelectProgress` | Known | Event handler |
| 0x00229A10 | `HandleWheelProgress` | Known | Event handler |
| 0x00229A24 | `HandleSelectScrub` | Known | Event handler |
| 0x00229A38 | `HandleWheelBrightness` | Known | Event handler |
| 0x00229A50 | `HandleNextTrackGotoScrub` | Known | Event handler |
| 0x00229A6C | `HandlePrevTrackGotoScrub` | Known | Event handler |
| 0x00229A88 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x0022CEF0 | `HandleSelectWalking` | Known | Event handler |
| 0x0022CF08 | `HandleSelectRunning` | Known | Event handler |
| 0x00230C5C | `HandleLanguage` | Known | Event handler |
| 0x00230C6C | `HandleResetAllSettings` | Known | Event handler |
| 0x00230C84 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x0023104C | `HandleUnlinkRemote` | Known | Event handler |
| 0x00231B3C | `HandleSelect` | Known | Event handler |
| 0x00231D6C | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x002334E4 | `Handle400MetersRun` | Known | Event handler |
| 0x002334FC | `HandleCustomRun` | Known | Event handler |
| 0x0023350C | `HandleResetToDefault` | Known | Event handler |
| 0x0023396C | `HandleSelect_Basic` | Known | Event handler |
| 0x00233984 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x00235A44 | `HandleSelect` | Known | Event handler |
| 0x00235BE0 | `HandleSelect` | Known | Event handler |
| 0x00235E80 | `HandleNextDay` | Known | Event handler |
| 0x00235E94 | `HandlePreviousDay` | Known | Event handler |
| 0x00236698 | `HandleMusicHilited` | Known | Event handler |
| 0x002366B0 | `HandleVideosHilited` | Known | Event handler |
| 0x002366C4 | `HandlePodcastsHilited` | Known | Event handler |
| 0x002366DC | `HandleGenericHilited` | Known | Event handler |
| 0x002366F4 | `HandlePhotosHilited` | Known | Event handler |
| 0x00236708 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x00236720 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x0023673C | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00236754 | `HandleArtistsHilited` | Known | Event handler |
| 0x0023676C | `HandleGenresHilited` | Known | Event handler |
| 0x00236780 | `HandleAlbumsHilited` | Known | Event handler |
| 0x00236794 | `HandleCompilationsHilited` | Known | Event handler |
| 0x00236968 | `HandleComposersHilited` | Known | Event handler |
| 0x00236980 | `HandleSongsHilited` | Known | Event handler |
| 0x00236994 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x002369AC | `HandleTVShowsHilited` | Known | Event handler |
| 0x002369C4 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x002369E0 | `HandleMoviesHilited` | Known | Event handler |
| 0x002369F4 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00236A10 | `HandleRentalsHilited` | Known | Event handler |
| 0x00236A28 | `HandleMusicSelected` | Known | Event handler |
| 0x00236A3C | `HandleVideosSelected` | Known | Event handler |
| 0x00236A54 | `HandlePodcastsSelected` | Known | Event handler |
| 0x00236C24 | `HandlePhotosSelected` | Known | Event handler |
| 0x00236C3C | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00236C54 | `HandleSongsSelected` | Known | Event handler |
| 0x00236C68 | `HandleAlbumsSelected` | Known | Event handler |
| 0x00236C80 | `HandleCompilationsSelected` | Known | Event handler |
| 0x00236C9C | `HandleArtistsSelected` | Known | Event handler |
| 0x00236CB4 | `HandleGenresSelected` | Known | Event handler |
| 0x00236CCC | `HandleComposersSelected` | Known | Event handler |
| 0x00236CE4 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x00236D00 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00236D1C | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00236EF4 | `HandleNowPlaying` | Known | Event handler |
| 0x00236F08 | `HandleTVShowsSelected` | Known | Event handler |
| 0x00236F20 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00236F3C | `HandleMoviesSelected` | Known | Event handler |
| 0x00236F54 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00236F74 | `HandleRentalsSelected` | Known | Event handler |
| 0x00236F8C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00236FA4 | `HandleLock` | Known | Event handler |
| 0x00236FB0 | `HandleBacklightSelected` | Known | Event handler |
| 0x00236FC8 | `HandleSleepSelected` | Known | Event handler |
| 0x00236FDC | `HandleNikePlusSelected` | Known | Event handler |
| 0x002398A0 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00239EB4 | `Handle400MetersWalk` | Known | Event handler |
| 0x00239ECC | `HandleCustomWalk` | Known | Event handler |
| 0x00239EE0 | `HandleResetToDefault` | Known | Event handler |
| 0x0023A1CC | `HandleSelect` | Known | Event handler |
| 0x0023A47C | `HandleWheel` | Known | Event handler |
| 0x0023BCB0 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x0023BF08 | `HandleNextDay` | Known | Event handler |
| 0x0023BF1C | `HandlePreviousDay` | Known | Event handler |
| 0x0023C164 | `HandleSelect` | Known | Event handler |
| 0x0023C400 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0023F0B8 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x0023F0D4 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x0024003C | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0024071C | `HandleSelect` | Known | Event handler |
| 0x00240DE8 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x0027E294 | `HandleDeleteClock` | Known | Event handler |
| 0x0027E2AC | `HandleSelectClock` | Known | Event handler |
| 0x0027E2C0 | `HandleHilited` | Known | Event handler |
| 0x0027E2D0 | `HandleWheel` | Known | Event handler |
| 0x0027E2DC | `HandleSelectLozinch` | Known | Event handler |
| 0x00402F0E | `HandleAudioFFDown` | Known | Event handler |
| 0x00402F37 | `HandleAudioFFUp` | Known | Event handler |
| 0x00402F62 | `HandleAudioMute` | Known | Event handler |
| 0x00402F95 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x00402FCA | `HandleAudioNext` | Known | Event handler |
| 0x00402FFA | `HandleAudioNextAlbum` | Known | Event handler |
| 0x00403031 | `HandleAudioNextChapter` | Known | Event handler |
| 0x0040306B | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x0040309F | `HandleAudioPause` | Known | Event handler |
| 0x004030CB | `HandleAudioPlay` | Known | Event handler |
| 0x004030F9 | `HandleAudioPlayPause` | Known | Event handler |
| 0x00403131 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x0040316A | `HandleAudioPrevious` | Known | Event handler |
| 0x0040319E | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x004031D5 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x0040320F | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x00403244 | `HandleAudioRepeat` | Known | Event handler |
| 0x00403270 | `HandleAudioRewDown` | Known | Event handler |
| 0x0040329B | `HandleAudioRewUp` | Known | Event handler |
| 0x004032CA | `HandleAudioShuffle` | Known | Event handler |
| 0x004032F8 | `HandleAudioStop` | Known | Event handler |
| 0x00403329 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x0040335E | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x00403395 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x004033C6 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x0040347F | `HandleNextPressAndHold` | Known | Event handler |
| 0x004034B0 | `HandleNext` | Known | Event handler |
| 0x004034E8 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x00403523 | `HandlePlayPause` | Known | Event handler |
| 0x00403557 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x0040358C | `HandlePrevious` | Known | Event handler |
| 0x00403619 | `HandleRemoteBacklight` | Known | Event handler |
| 0x00403651 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x0040368B | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x004036C4 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x004036F9 | `HandleRemoteEvent` | Known | Event handler |
| 0x00403725 | `HandleRemoteFFDown` | Known | Event handler |
| 0x00403750 | `HandleRemoteFFUp` | Known | Event handler |
| 0x0040377D | `HandleRemoteMenuDown` | Known | Event handler |
| 0x004037AC | `HandleRemoteMenuUp` | Known | Event handler |
| 0x004037DB | `HandleRemoteMute` | Known | Event handler |
| 0x0040380D | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x00403846 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x00403882 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x004038C2 | `HandleRemoteOff` | Known | Event handler |
| 0x004038EB | `HandleRemoteOff` | Known | Event handler |
| 0x00403915 | `HandleRemoteOn` | Known | Event handler |
| 0x00403941 | `HandleRemotePause` | Known | Event handler |
| 0x0040396F | `HandleRemotePlay` | Known | Event handler |
| 0x004039AD | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x004039EE | `HandleRemotePlayPause` | Known | Event handler |
| 0x00403A25 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x00403A5E | `HandleRemotePrevChapter` | Known | Event handler |
| 0x00403A9A | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x00403AD1 | `HandleRemoteRepeat` | Known | Event handler |
| 0x00403AFF | `HandleRemoteRewDown` | Known | Event handler |
| 0x00403B2C | `HandleRemoteRewUp` | Known | Event handler |
| 0x00403B5C | `HandleRemoteSelectDown` | Known | Event handler |
| 0x00403B8F | `HandleRemoteSelectUp` | Known | Event handler |
| 0x00403BC3 | `HandleRemoteShuffle` | Known | Event handler |
| 0x00403BF3 | `HandleRemoteStop` | Known | Event handler |
| 0x00403C23 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x00403C58 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x00403C90 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x00403CC7 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x00403D00 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x00403D33 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x00403D68 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x00403D9B | `HandleVideoFFDown` | Known | Event handler |
| 0x00403DC4 | `HandleVideoFFUp` | Known | Event handler |
| 0x00403DF7 | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x00403E2C | `HandleVideoNext` | Known | Event handler |
| 0x00403E5E | `HandleVideoNextChapter` | Known | Event handler |
| 0x00403E95 | `HandleVideoNextFrame` | Known | Event handler |
| 0x00403EC6 | `HandleVideoPause` | Known | Event handler |
| 0x00403EF2 | `HandleVideoPlay` | Known | Event handler |
| 0x00403F20 | `HandleVideoPlayPause` | Known | Event handler |
| 0x00403F58 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x00403F91 | `HandleVideoPrevious` | Known | Event handler |
| 0x00403FC7 | `HandleVideoPrevChapter` | Known | Event handler |
| 0x00403FFE | `HandleVideoPrevFrame` | Known | Event handler |
| 0x0040402D | `HandleVideoRewDown` | Known | Event handler |
| 0x00404058 | `HandleVideoRewUp` | Known | Event handler |
| 0x00404084 | `HandleVideoStop` | Known | Event handler |
| 0x00753F4A | `HandleAddressBook` | Known | Event handler |
| 0x007544DE | `HandleSelect` | Known | Event handler |
| 0x00754519 | `HandleHilite` | Known | Event handler |
| 0x0075459A | `HandleSelectRegion` | Known | Event handler |
| 0x0075463A | `HandleSelectRegion` | Known | Event handler |
| 0x007546D6 | `HandleSelectRegion` | Known | Event handler |
| 0x0075477A | `HandleSelectRegion` | Known | Event handler |
| 0x00754820 | `HandleSelectRegion` | Known | Event handler |
| 0x007548C0 | `HandleSelectRegion` | Known | Event handler |
| 0x0075496C | `HandleSelectRegion` | Known | Event handler |
| 0x00754A0E | `HandleSelectRegion` | Known | Event handler |
| 0x00754ABE | `HandleSelectCity` | Known | Event handler |
| 0x00754B2A | `HandleHighlightCity` | Known | Event handler |
| 0x00754B63 | `HandleSelectCity` | Known | Event handler |
| 0x00754BCF | `HandleHighlightCity` | Known | Event handler |
| 0x00754C08 | `HandleSelectCity` | Known | Event handler |
| 0x00754C74 | `HandleHighlightCity` | Known | Event handler |
| 0x00754CAD | `HandleSelectCity` | Known | Event handler |
| 0x00754D19 | `HandleHighlightCity` | Known | Event handler |
| 0x00754D52 | `HandleSelectCity` | Known | Event handler |
| 0x00754DBE | `HandleHighlightCity` | Known | Event handler |
| 0x00754DF7 | `HandleSelectCity` | Known | Event handler |
| 0x00754E63 | `HandleHighlightCity` | Known | Event handler |
| 0x00754E9C | `HandleSelectCity` | Known | Event handler |
| 0x00754F08 | `HandleHighlightCity` | Known | Event handler |
| 0x00754F41 | `HandleSelectCity` | Known | Event handler |
| 0x00754FAD | `HandleHighlightCity` | Known | Event handler |
| 0x00754FE6 | `HandleSelectCity` | Known | Event handler |
| 0x00755052 | `HandleHighlightCity` | Known | Event handler |
| 0x0075508B | `HandleSelectCity` | Known | Event handler |
| 0x007550F7 | `HandleHighlightCity` | Known | Event handler |
| 0x00755130 | `HandleSelectCity` | Known | Event handler |
| 0x0075519C | `HandleHighlightCity` | Known | Event handler |
| 0x007551D5 | `HandleSelectCity` | Known | Event handler |
| 0x00755241 | `HandleHighlightCity` | Known | Event handler |
| 0x0075527A | `HandleSelectCity` | Known | Event handler |
| 0x007552E6 | `HandleHighlightCity` | Known | Event handler |
| 0x0075531F | `HandleSelectCity` | Known | Event handler |
| 0x0075538B | `HandleHighlightCity` | Known | Event handler |
| 0x007553C4 | `HandleSelectCity` | Known | Event handler |
| 0x00755430 | `HandleHighlightCity` | Known | Event handler |
| 0x00755469 | `HandleSelectCity` | Known | Event handler |
| 0x007554D5 | `HandleHighlightCity` | Known | Event handler |
| 0x0075550E | `HandleSelectCity` | Known | Event handler |
| 0x0075557A | `HandleHighlightCity` | Known | Event handler |
| 0x007555B3 | `HandleSelectCity` | Known | Event handler |
| 0x0075561F | `HandleHighlightCity` | Known | Event handler |
| 0x00755658 | `HandleSelectCity` | Known | Event handler |
| 0x007556C4 | `HandleHighlightCity` | Known | Event handler |
| 0x007556FD | `HandleSelectCity` | Known | Event handler |
| 0x00755769 | `HandleHighlightCity` | Known | Event handler |
| 0x007557A2 | `HandleSelectCity` | Known | Event handler |
| 0x0075580E | `HandleHighlightCity` | Known | Event handler |
| 0x00755847 | `HandleSelectCity` | Known | Event handler |
| 0x007558B3 | `HandleHighlightCity` | Known | Event handler |
| 0x007558EC | `HandleSelectCity` | Known | Event handler |
| 0x00755958 | `HandleHighlightCity` | Known | Event handler |
| 0x00755991 | `HandleSelectCity` | Known | Event handler |
| 0x007559FD | `HandleHighlightCity` | Known | Event handler |
| 0x00755A36 | `HandleSelectCity` | Known | Event handler |
| 0x00755AA2 | `HandleHighlightCity` | Known | Event handler |
| 0x00755ADB | `HandleSelectCity` | Known | Event handler |
| 0x00755B47 | `HandleHighlightCity` | Known | Event handler |
| 0x00755B80 | `HandleSelectCity` | Known | Event handler |
| 0x00755BEC | `HandleHighlightCity` | Known | Event handler |
| 0x00755C25 | `HandleSelectCity` | Known | Event handler |
| 0x00755C91 | `HandleHighlightCity` | Known | Event handler |
| 0x00755CCA | `HandleSelectCity` | Known | Event handler |
| 0x00755D36 | `HandleHighlightCity` | Known | Event handler |
| 0x00755D6F | `HandleSelectCity` | Known | Event handler |
| 0x00755DDB | `HandleHighlightCity` | Known | Event handler |
| 0x00755E14 | `HandleSelectCity` | Known | Event handler |
| 0x00755E80 | `HandleHighlightCity` | Known | Event handler |
| 0x00755EBE | `HandleSelectCity` | Known | Event handler |
| 0x00755F2A | `HandleHighlightCity` | Known | Event handler |
| 0x00755F63 | `HandleSelectCity` | Known | Event handler |
| 0x00755FCF | `HandleHighlightCity` | Known | Event handler |
| 0x00756008 | `HandleSelectCity` | Known | Event handler |
| 0x00756074 | `HandleHighlightCity` | Known | Event handler |
| 0x007560AD | `HandleSelectCity` | Known | Event handler |
| 0x00756119 | `HandleHighlightCity` | Known | Event handler |
| 0x00756152 | `HandleSelectCity` | Known | Event handler |
| 0x007561BE | `HandleHighlightCity` | Known | Event handler |
| 0x007561F7 | `HandleSelectCity` | Known | Event handler |
| 0x00756263 | `HandleHighlightCity` | Known | Event handler |
| 0x0075629C | `HandleSelectCity` | Known | Event handler |
| 0x00756308 | `HandleHighlightCity` | Known | Event handler |
| 0x00756341 | `HandleSelectCity` | Known | Event handler |
| 0x007563AD | `HandleHighlightCity` | Known | Event handler |
| 0x007563E6 | `HandleSelectCity` | Known | Event handler |
| 0x00756452 | `HandleHighlightCity` | Known | Event handler |
| 0x0075648B | `HandleSelectCity` | Known | Event handler |
| 0x007564F7 | `HandleHighlightCity` | Known | Event handler |
| 0x00756530 | `HandleSelectCity` | Known | Event handler |
| 0x0075659C | `HandleHighlightCity` | Known | Event handler |
| 0x007565D5 | `HandleSelectCity` | Known | Event handler |
| 0x00756641 | `HandleHighlightCity` | Known | Event handler |
| 0x0075667A | `HandleSelectCity` | Known | Event handler |
| 0x007566E6 | `HandleHighlightCity` | Known | Event handler |
| 0x0075671F | `HandleSelectCity` | Known | Event handler |
| 0x0075678B | `HandleHighlightCity` | Known | Event handler |
| 0x007567C4 | `HandleSelectCity` | Known | Event handler |
| 0x00756830 | `HandleHighlightCity` | Known | Event handler |
| 0x00756869 | `HandleSelectCity` | Known | Event handler |
| 0x007568D5 | `HandleHighlightCity` | Known | Event handler |
| 0x0075690E | `HandleSelectCity` | Known | Event handler |
| 0x0075697A | `HandleHighlightCity` | Known | Event handler |
| 0x007569B3 | `HandleSelectCity` | Known | Event handler |
| 0x00756A1F | `HandleHighlightCity` | Known | Event handler |
| 0x00756A58 | `HandleSelectCity` | Known | Event handler |
| 0x00756AC4 | `HandleHighlightCity` | Known | Event handler |
| 0x00756AFD | `HandleSelectCity` | Known | Event handler |
| 0x00756B69 | `HandleHighlightCity` | Known | Event handler |
| 0x00756BA2 | `HandleSelectCity` | Known | Event handler |
| 0x00756C0E | `HandleHighlightCity` | Known | Event handler |
| 0x00756C47 | `HandleSelectCity` | Known | Event handler |
| 0x00756CB3 | `HandleHighlightCity` | Known | Event handler |
| 0x00756CEC | `HandleSelectCity` | Known | Event handler |
| 0x00756D58 | `HandleHighlightCity` | Known | Event handler |
| 0x00756D91 | `HandleSelectCity` | Known | Event handler |
| 0x00756DFD | `HandleHighlightCity` | Known | Event handler |
| 0x00756E36 | `HandleSelectCity` | Known | Event handler |
| 0x00756EA2 | `HandleHighlightCity` | Known | Event handler |
| 0x00756EDB | `HandleSelectCity` | Known | Event handler |
| 0x00756F47 | `HandleHighlightCity` | Known | Event handler |
| 0x00756F80 | `HandleSelectCity` | Known | Event handler |
| 0x00756FEC | `HandleHighlightCity` | Known | Event handler |
| 0x00757025 | `HandleSelectCity` | Known | Event handler |
| 0x00757091 | `HandleHighlightCity` | Known | Event handler |
| 0x007570CA | `HandleSelectCity` | Known | Event handler |
| 0x00757136 | `HandleHighlightCity` | Known | Event handler |
| 0x0075716F | `HandleSelectCity` | Known | Event handler |
| 0x007571DB | `HandleHighlightCity` | Known | Event handler |
| 0x00757214 | `HandleSelectCity` | Known | Event handler |
| 0x00757280 | `HandleHighlightCity` | Known | Event handler |
| 0x007572B9 | `HandleSelectCity` | Known | Event handler |
| 0x00757325 | `HandleHighlightCity` | Known | Event handler |
| 0x0075735E | `HandleSelectCity` | Known | Event handler |
| 0x007573CA | `HandleHighlightCity` | Known | Event handler |
| 0x00757403 | `HandleSelectCity` | Known | Event handler |
| 0x0075746F | `HandleHighlightCity` | Known | Event handler |
| 0x007574A8 | `HandleSelectCity` | Known | Event handler |
| 0x00757514 | `HandleHighlightCity` | Known | Event handler |
| 0x0075754D | `HandleSelectCity` | Known | Event handler |
| 0x007575B9 | `HandleHighlightCity` | Known | Event handler |
| 0x007575F2 | `HandleSelectCity` | Known | Event handler |
| 0x0075765E | `HandleHighlightCity` | Known | Event handler |
| 0x00757697 | `HandleSelectCity` | Known | Event handler |
| 0x00757703 | `HandleHighlightCity` | Known | Event handler |
| 0x0075773C | `HandleSelectCity` | Known | Event handler |
| 0x007577A8 | `HandleHighlightCity` | Known | Event handler |
| 0x007577E1 | `HandleSelectCity` | Known | Event handler |
| 0x0075784D | `HandleHighlightCity` | Known | Event handler |
| 0x00757886 | `HandleSelectCity` | Known | Event handler |
| 0x007578F2 | `HandleHighlightCity` | Known | Event handler |
| 0x0075792B | `HandleSelectCity` | Known | Event handler |
| 0x00757997 | `HandleHighlightCity` | Known | Event handler |
| 0x007579D0 | `HandleSelectCity` | Known | Event handler |
| 0x00757A3C | `HandleHighlightCity` | Known | Event handler |
| 0x00757A75 | `HandleSelectCity` | Known | Event handler |
| 0x00757AE1 | `HandleHighlightCity` | Known | Event handler |
| 0x00757B1A | `HandleSelectCity` | Known | Event handler |
| 0x00757B86 | `HandleHighlightCity` | Known | Event handler |
| 0x00757BBF | `HandleSelectCity` | Known | Event handler |
| 0x00757C2B | `HandleHighlightCity` | Known | Event handler |
| 0x00757C64 | `HandleSelectCity` | Known | Event handler |
| 0x00757CD0 | `HandleHighlightCity` | Known | Event handler |
| 0x00757D09 | `HandleSelectCity` | Known | Event handler |
| 0x00757D75 | `HandleHighlightCity` | Known | Event handler |
| 0x00757DAE | `HandleSelectCity` | Known | Event handler |
| 0x00757E1A | `HandleHighlightCity` | Known | Event handler |
| 0x00757E53 | `HandleSelectCity` | Known | Event handler |
| 0x00757EBF | `HandleHighlightCity` | Known | Event handler |
| 0x00757EF8 | `HandleSelectCity` | Known | Event handler |
| 0x00757F64 | `HandleHighlightCity` | Known | Event handler |
| 0x00757F9D | `HandleSelectCity` | Known | Event handler |
| 0x00758009 | `HandleHighlightCity` | Known | Event handler |
| 0x00758042 | `HandleSelectCity` | Known | Event handler |
| 0x007580AE | `HandleHighlightCity` | Known | Event handler |
| 0x007580E7 | `HandleSelectCity` | Known | Event handler |
| 0x00758153 | `HandleHighlightCity` | Known | Event handler |
| 0x0075818C | `HandleSelectCity` | Known | Event handler |
| 0x007581F8 | `HandleHighlightCity` | Known | Event handler |
| 0x00758231 | `HandleSelectCity` | Known | Event handler |
| 0x0075829D | `HandleHighlightCity` | Known | Event handler |
| 0x007582D6 | `HandleSelectCity` | Known | Event handler |
| 0x00758342 | `HandleHighlightCity` | Known | Event handler |
| 0x00758382 | `HandleSelectCity` | Known | Event handler |
| 0x007583EE | `HandleHighlightCity` | Known | Event handler |
| 0x00758427 | `HandleSelectCity` | Known | Event handler |
| 0x00758493 | `HandleHighlightCity` | Known | Event handler |
| 0x007584CC | `HandleSelectCity` | Known | Event handler |
| 0x00758538 | `HandleHighlightCity` | Known | Event handler |
| 0x00758576 | `HandleSelectCity` | Known | Event handler |
| 0x007585E2 | `HandleHighlightCity` | Known | Event handler |
| 0x0075861B | `HandleSelectCity` | Known | Event handler |
| 0x00758687 | `HandleHighlightCity` | Known | Event handler |
| 0x007586C0 | `HandleSelectCity` | Known | Event handler |
| 0x0075872C | `HandleHighlightCity` | Known | Event handler |
| 0x00758765 | `HandleSelectCity` | Known | Event handler |
| 0x007587D1 | `HandleHighlightCity` | Known | Event handler |
| 0x0075880A | `HandleSelectCity` | Known | Event handler |
| 0x00758876 | `HandleHighlightCity` | Known | Event handler |
| 0x007588AF | `HandleSelectCity` | Known | Event handler |
| 0x0075891B | `HandleHighlightCity` | Known | Event handler |
| 0x00758954 | `HandleSelectCity` | Known | Event handler |
| 0x007589C0 | `HandleHighlightCity` | Known | Event handler |
| 0x007589F9 | `HandleSelectCity` | Known | Event handler |
| 0x00758A65 | `HandleHighlightCity` | Known | Event handler |
| 0x00758AA2 | `HandleSelectCity` | Known | Event handler |
| 0x00758B0E | `HandleHighlightCity` | Known | Event handler |
| 0x00758B47 | `HandleSelectCity` | Known | Event handler |
| 0x00758BB3 | `HandleHighlightCity` | Known | Event handler |
| 0x00758BEC | `HandleSelectCity` | Known | Event handler |
| 0x00758C58 | `HandleHighlightCity` | Known | Event handler |
| 0x00758C91 | `HandleSelectCity` | Known | Event handler |
| 0x00758CFD | `HandleHighlightCity` | Known | Event handler |
| 0x00758D36 | `HandleSelectCity` | Known | Event handler |
| 0x00758DA2 | `HandleHighlightCity` | Known | Event handler |
| 0x00758DDB | `HandleSelectCity` | Known | Event handler |
| 0x00758E47 | `HandleHighlightCity` | Known | Event handler |
| 0x00758E80 | `HandleSelectCity` | Known | Event handler |
| 0x00758EEC | `HandleHighlightCity` | Known | Event handler |
| 0x00758F25 | `HandleSelectCity` | Known | Event handler |
| 0x00758F91 | `HandleHighlightCity` | Known | Event handler |
| 0x00758FCA | `HandleSelectCity` | Known | Event handler |
| 0x00759036 | `HandleHighlightCity` | Known | Event handler |
| 0x0075906F | `HandleSelectCity` | Known | Event handler |
| 0x007590DB | `HandleHighlightCity` | Known | Event handler |
| 0x00759114 | `HandleSelectCity` | Known | Event handler |
| 0x00759180 | `HandleHighlightCity` | Known | Event handler |
| 0x007591B9 | `HandleSelectCity` | Known | Event handler |
| 0x00759225 | `HandleHighlightCity` | Known | Event handler |
| 0x0075925E | `HandleSelectCity` | Known | Event handler |
| 0x007592CA | `HandleHighlightCity` | Known | Event handler |
| 0x00759303 | `HandleSelectCity` | Known | Event handler |
| 0x0075936F | `HandleHighlightCity` | Known | Event handler |
| 0x007593A8 | `HandleSelectCity` | Known | Event handler |
| 0x00759414 | `HandleHighlightCity` | Known | Event handler |
| 0x0075944D | `HandleSelectCity` | Known | Event handler |
| 0x007594B9 | `HandleHighlightCity` | Known | Event handler |
| 0x007594F2 | `HandleSelectCity` | Known | Event handler |
| 0x0075955E | `HandleHighlightCity` | Known | Event handler |
| 0x00759597 | `HandleSelectCity` | Known | Event handler |
| 0x00759603 | `HandleHighlightCity` | Known | Event handler |
| 0x0075963C | `HandleSelectCity` | Known | Event handler |
| 0x007596A8 | `HandleHighlightCity` | Known | Event handler |
| 0x007596E1 | `HandleSelectCity` | Known | Event handler |
| 0x0075974D | `HandleHighlightCity` | Known | Event handler |
| 0x00759786 | `HandleSelectCity` | Known | Event handler |
| 0x007597F2 | `HandleHighlightCity` | Known | Event handler |
| 0x0075982B | `HandleSelectCity` | Known | Event handler |
| 0x00759897 | `HandleHighlightCity` | Known | Event handler |
| 0x007598D0 | `HandleSelectCity` | Known | Event handler |
| 0x0075993C | `HandleHighlightCity` | Known | Event handler |
| 0x00759975 | `HandleSelectCity` | Known | Event handler |
| 0x007599E1 | `HandleHighlightCity` | Known | Event handler |
| 0x00759A1A | `HandleSelectCity` | Known | Event handler |
| 0x00759A86 | `HandleHighlightCity` | Known | Event handler |
| 0x00759ABF | `HandleSelectCity` | Known | Event handler |
| 0x00759B2B | `HandleHighlightCity` | Known | Event handler |
| 0x00759B64 | `HandleSelectCity` | Known | Event handler |
| 0x00759BD0 | `HandleHighlightCity` | Known | Event handler |
| 0x00759C09 | `HandleSelectCity` | Known | Event handler |
| 0x00759C75 | `HandleHighlightCity` | Known | Event handler |
| 0x00759CAE | `HandleSelectCity` | Known | Event handler |
| 0x00759D1A | `HandleHighlightCity` | Known | Event handler |
| 0x00759D53 | `HandleSelectCity` | Known | Event handler |
| 0x00759DBF | `HandleHighlightCity` | Known | Event handler |
| 0x00759DF8 | `HandleSelectCity` | Known | Event handler |
| 0x00759E64 | `HandleHighlightCity` | Known | Event handler |
| 0x00759E9D | `HandleSelectCity` | Known | Event handler |
| 0x00759F09 | `HandleHighlightCity` | Known | Event handler |
| 0x00759F42 | `HandleSelectCity` | Known | Event handler |
| 0x00759FAE | `HandleHighlightCity` | Known | Event handler |
| 0x00759FE7 | `HandleSelectCity` | Known | Event handler |
| 0x0075A053 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A092 | `HandleSelectCity` | Known | Event handler |
| 0x0075A0FE | `HandleHighlightCity` | Known | Event handler |
| 0x0075A137 | `HandleSelectCity` | Known | Event handler |
| 0x0075A1A3 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A1DC | `HandleSelectCity` | Known | Event handler |
| 0x0075A248 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A281 | `HandleSelectCity` | Known | Event handler |
| 0x0075A2ED | `HandleHighlightCity` | Known | Event handler |
| 0x0075A326 | `HandleSelectCity` | Known | Event handler |
| 0x0075A392 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A3CB | `HandleSelectCity` | Known | Event handler |
| 0x0075A437 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A470 | `HandleSelectCity` | Known | Event handler |
| 0x0075A4DC | `HandleHighlightCity` | Known | Event handler |
| 0x0075A515 | `HandleSelectCity` | Known | Event handler |
| 0x0075A581 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A5BA | `HandleSelectCity` | Known | Event handler |
| 0x0075A626 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A65F | `HandleSelectCity` | Known | Event handler |
| 0x0075A6CB | `HandleHighlightCity` | Known | Event handler |
| 0x0075A704 | `HandleSelectCity` | Known | Event handler |
| 0x0075A770 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A7A9 | `HandleSelectCity` | Known | Event handler |
| 0x0075A815 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A84E | `HandleSelectCity` | Known | Event handler |
| 0x0075A8BA | `HandleHighlightCity` | Known | Event handler |
| 0x0075A8F3 | `HandleSelectCity` | Known | Event handler |
| 0x0075A95F | `HandleHighlightCity` | Known | Event handler |
| 0x0075A998 | `HandleSelectCity` | Known | Event handler |
| 0x0075AA04 | `HandleHighlightCity` | Known | Event handler |
| 0x0075AA3D | `HandleSelectCity` | Known | Event handler |
| 0x0075AAA9 | `HandleHighlightCity` | Known | Event handler |
| 0x0075AAE2 | `HandleSelectCity` | Known | Event handler |
| 0x0075AB4E | `HandleHighlightCity` | Known | Event handler |
| 0x0075AB87 | `HandleSelectCity` | Known | Event handler |
| 0x0075ABF3 | `HandleHighlightCity` | Known | Event handler |
| 0x0075AC2C | `HandleSelectCity` | Known | Event handler |
| 0x0075AC98 | `HandleHighlightCity` | Known | Event handler |
| 0x0075ACD1 | `HandleSelectCity` | Known | Event handler |
| 0x0075AD3D | `HandleHighlightCity` | Known | Event handler |
| 0x0075AD76 | `HandleSelectCity` | Known | Event handler |
| 0x0075ADE2 | `HandleHighlightCity` | Known | Event handler |
| 0x0075AE1B | `HandleSelectCity` | Known | Event handler |
| 0x0075AE87 | `HandleHighlightCity` | Known | Event handler |
| 0x0075AEC0 | `HandleSelectCity` | Known | Event handler |
| 0x0075AF2C | `HandleHighlightCity` | Known | Event handler |
| 0x0075AF65 | `HandleSelectCity` | Known | Event handler |
| 0x0075AFD1 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B00A | `HandleSelectCity` | Known | Event handler |
| 0x0075B076 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B0AF | `HandleSelectCity` | Known | Event handler |
| 0x0075B11B | `HandleHighlightCity` | Known | Event handler |
| 0x0075B154 | `HandleSelectCity` | Known | Event handler |
| 0x0075B1C0 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B1F9 | `HandleSelectCity` | Known | Event handler |
| 0x0075B265 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B29E | `HandleSelectCity` | Known | Event handler |
| 0x0075B30A | `HandleHighlightCity` | Known | Event handler |
| 0x0075B343 | `HandleSelectCity` | Known | Event handler |
| 0x0075B3AF | `HandleHighlightCity` | Known | Event handler |
| 0x0075B3E8 | `HandleSelectCity` | Known | Event handler |
| 0x0075B454 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B48D | `HandleSelectCity` | Known | Event handler |
| 0x0075B4F9 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B532 | `HandleSelectCity` | Known | Event handler |
| 0x0075B59E | `HandleHighlightCity` | Known | Event handler |
| 0x0075B5D7 | `HandleSelectCity` | Known | Event handler |
| 0x0075B643 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B67C | `HandleSelectCity` | Known | Event handler |
| 0x0075B6E8 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B721 | `HandleSelectCity` | Known | Event handler |
| 0x0075B78D | `HandleHighlightCity` | Known | Event handler |
| 0x0075B7C6 | `HandleSelectCity` | Known | Event handler |
| 0x0075B832 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B86B | `HandleSelectCity` | Known | Event handler |
| 0x0075B8D7 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B910 | `HandleSelectCity` | Known | Event handler |
| 0x0075B97C | `HandleHighlightCity` | Known | Event handler |
| 0x0075B9B5 | `HandleSelectCity` | Known | Event handler |
| 0x0075BA21 | `HandleHighlightCity` | Known | Event handler |
| 0x0075BA5A | `HandleSelectCity` | Known | Event handler |
| 0x0075BAC6 | `HandleHighlightCity` | Known | Event handler |
| 0x0075BAFF | `HandleSelectCity` | Known | Event handler |
| 0x0075BB6B | `HandleHighlightCity` | Known | Event handler |
| 0x0075BBA4 | `HandleSelectCity` | Known | Event handler |
| 0x0075BC10 | `HandleHighlightCity` | Known | Event handler |
| 0x0075BC49 | `HandleSelectCity` | Known | Event handler |
| 0x0075BCB5 | `HandleHighlightCity` | Known | Event handler |
| 0x0075BCEE | `HandleSelectCity` | Known | Event handler |
| 0x0075BD5A | `HandleHighlightCity` | Known | Event handler |
| 0x0075BD93 | `HandleSelectCity` | Known | Event handler |
| 0x0075BDFF | `HandleHighlightCity` | Known | Event handler |
| 0x0075BE38 | `HandleSelectCity` | Known | Event handler |
| 0x0075BEA4 | `HandleHighlightCity` | Known | Event handler |
| 0x0075BEDD | `HandleSelectCity` | Known | Event handler |
| 0x0075BF49 | `HandleHighlightCity` | Known | Event handler |
| 0x0075BF82 | `HandleSelectCity` | Known | Event handler |
| 0x0075BFEE | `HandleHighlightCity` | Known | Event handler |
| 0x0075C027 | `HandleSelectCity` | Known | Event handler |
| 0x0075C093 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C0D2 | `HandleSelectCity` | Known | Event handler |
| 0x0075C13E | `HandleHighlightCity` | Known | Event handler |
| 0x0075C177 | `HandleSelectCity` | Known | Event handler |
| 0x0075C1E3 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C21C | `HandleSelectCity` | Known | Event handler |
| 0x0075C288 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C2C1 | `HandleSelectCity` | Known | Event handler |
| 0x0075C32D | `HandleHighlightCity` | Known | Event handler |
| 0x0075C366 | `HandleSelectCity` | Known | Event handler |
| 0x0075C3D2 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C412 | `HandleSelectCity` | Known | Event handler |
| 0x0075C47E | `HandleHighlightCity` | Known | Event handler |
| 0x0075C4B7 | `HandleSelectCity` | Known | Event handler |
| 0x0075C523 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C55C | `HandleSelectCity` | Known | Event handler |
| 0x0075C5C8 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C601 | `HandleSelectCity` | Known | Event handler |
| 0x0075C66D | `HandleHighlightCity` | Known | Event handler |
| 0x0075C6A6 | `HandleSelectCity` | Known | Event handler |
| 0x0075C712 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C74B | `HandleSelectCity` | Known | Event handler |
| 0x0075C7B7 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C7F0 | `HandleSelectCity` | Known | Event handler |
| 0x0075C85C | `HandleHighlightCity` | Known | Event handler |
| 0x0075C895 | `HandleSelectCity` | Known | Event handler |
| 0x0075C901 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C93A | `HandleSelectCity` | Known | Event handler |
| 0x0075C9A6 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C9DF | `HandleSelectCity` | Known | Event handler |
| 0x0075CA4B | `HandleHighlightCity` | Known | Event handler |
| 0x0075CA84 | `HandleSelectCity` | Known | Event handler |
| 0x0075CAF0 | `HandleHighlightCity` | Known | Event handler |
| 0x0075CB29 | `HandleSelectCity` | Known | Event handler |
| 0x0075CB95 | `HandleHighlightCity` | Known | Event handler |
| 0x0075CBCE | `HandleSelectCity` | Known | Event handler |
| 0x0075CC3A | `HandleHighlightCity` | Known | Event handler |
| 0x0075CC73 | `HandleSelectCity` | Known | Event handler |
| 0x0075CCDF | `HandleHighlightCity` | Known | Event handler |
| 0x0075CD18 | `HandleSelectCity` | Known | Event handler |
| 0x0075CD84 | `HandleHighlightCity` | Known | Event handler |
| 0x0075CDBD | `HandleSelectCity` | Known | Event handler |
| 0x0075CE29 | `HandleHighlightCity` | Known | Event handler |
| 0x0075CE62 | `HandleSelectCity` | Known | Event handler |
| 0x0075CECE | `HandleHighlightCity` | Known | Event handler |
| 0x0075D3C6 | `HandleMusicSelected` | Known | Event handler |
| 0x0075D408 | `HandleMusicHilited` | Known | Event handler |
| 0x0075D440 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x0075D486 | `HandleMusicHilited` | Known | Event handler |
| 0x0075D4BE | `HandlePlaylistsSelected` | Known | Event handler |
| 0x0075D504 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x0075D540 | `HandleArtistsSelected` | Known | Event handler |
| 0x0075D584 | `HandleArtistsHilited` | Known | Event handler |
| 0x0075D5BE | `HandleAlbumsSelected` | Known | Event handler |
| 0x0075D601 | `HandleAlbumsHilited` | Known | Event handler |
| 0x0075D63A | `HandleCompilationsSelected` | Known | Event handler |
| 0x0075D683 | `HandleCompilationsHilited` | Known | Event handler |
| 0x0075D6C2 | `HandleSongsSelected` | Known | Event handler |
| 0x0075D704 | `HandleSongsHilited` | Known | Event handler |
| 0x0075D73C | `HandleGenresSelected` | Known | Event handler |
| 0x0075D77F | `HandleGenresHilited` | Known | Event handler |
| 0x0075D7B8 | `HandleComposersSelected` | Known | Event handler |
| 0x0075D7FE | `HandleComposersHilited` | Known | Event handler |
| 0x0075D83A | `HandleAudiobooksSelected` | Known | Event handler |
| 0x0075D881 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x0075D940 | `HandleMusicHilited` | Known | Event handler |
| 0x0075D978 | `HandleVideosSelected` | Known | Event handler |
| 0x0075D9BB | `HandleVideosHilited` | Known | Event handler |
| 0x0075D9F4 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x0075DA3F | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x0075DA80 | `HandleMoviesSelected` | Known | Event handler |
| 0x0075DAC3 | `HandleMoviesHilited` | Known | Event handler |
| 0x0075DAFC | `HandleTVShowsSelected` | Known | Event handler |
| 0x0075DB40 | `HandleTVShowsHilited` | Known | Event handler |
| 0x0075DB7A | `HandleMusicVideosSelected` | Known | Event handler |
| 0x0075DBC2 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x0075DC00 | `HandleRentalsSelected` | Known | Event handler |
| 0x0075DC44 | `HandleRentalsHilited` | Known | Event handler |
| 0x0075DC7E | `HandlePhotosSelected` | Known | Event handler |
| 0x0075DCC1 | `HandlePhotosHilited` | Known | Event handler |
| 0x0075DCFA | `HandlePhotosSelected` | Known | Event handler |
| 0x0075DD3D | `HandlePhotosHilited` | Known | Event handler |
| 0x0075DD76 | `HandlePodcastsSelected` | Known | Event handler |
| 0x0075DDBB | `HandlePodcastsHilited` | Known | Event handler |
| 0x0075DE6E | `HandleGenericHilited` | Known | Event handler |
| 0x0075DF67 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E44C | `HandleLock` | Known | Event handler |
| 0x0075E5BD | `HandleNikePlusSelected` | Known | Event handler |
| 0x0075E602 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E708 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E807 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E8F4 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E9F1 | `HandleGenericHilited` | Known | Event handler |
| 0x0075EA6B | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x0075EAB4 | `HandleGenericHilited` | Known | Event handler |
| 0x0075EB2D | `HandleBacklightSelected` | Known | Event handler |
| 0x0075EB73 | `HandleGenericHilited` | Known | Event handler |
| 0x0075EBEE | `HandleSleepSelected` | Known | Event handler |
| 0x0075EC30 | `HandleGenericHilited` | Known | Event handler |
| 0x0075ECA7 | `HandleNowPlaying` | Known | Event handler |
| 0x0075ED1F | `HandleNowPlayingHilited` | Known | Event handler |
| 0x0075ED62 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x0075EDA8 | `HandleMusicHilited` | Known | Event handler |
| 0x0075EDE0 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x0075EE26 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x0075EE64 | `HandleArtistsSelected` | Known | Event handler |
| 0x0075EEA8 | `HandleArtistsHilited` | Known | Event handler |
| 0x0075EEE2 | `HandleAlbumsSelected` | Known | Event handler |
| 0x0075EF25 | `HandleAlbumsHilited` | Known | Event handler |
| 0x0075EF5E | `HandleCompilationsSelected` | Known | Event handler |
| 0x0075EFA7 | `HandleCompilationsHilited` | Known | Event handler |
| 0x0075EFE6 | `HandleSongsSelected` | Known | Event handler |
| 0x0075F028 | `HandleSongsHilited` | Known | Event handler |
| 0x0075F0D3 | `HandleGenericHilited` | Known | Event handler |
| 0x0075F14B | `HandleGenresSelected` | Known | Event handler |
| 0x0075F18E | `HandleGenresHilited` | Known | Event handler |
| 0x0075F1C7 | `HandleComposersSelected` | Known | Event handler |
| 0x0075F20D | `HandleComposersHilited` | Known | Event handler |
| 0x0075F249 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x0075F290 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x0075F34F | `HandleMusicHilited` | Known | Event handler |
| 0x0075F3C5 | `HandlePlayPause` | Known | Event handler |
| 0x0075F3FA | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x0075F4E4 | `HandleSelect` | Known | Event handler |
| 0x0075F52A | `HandleMoviesSelected` | Known | Event handler |
| 0x0075F56D | `HandleMoviesHilited` | Known | Event handler |
| 0x0075F5A6 | `HandleRentalsSelected` | Known | Event handler |
| 0x0075F5EA | `HandleRentalsHilited` | Known | Event handler |
| 0x0075F624 | `HandleTVShowsSelected` | Known | Event handler |
| 0x0075F668 | `HandleTVShowsHilited` | Known | Event handler |
| 0x0075F6A2 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x0075F6EA | `HandleMusicVideosHilited` | Known | Event handler |
| 0x0075F728 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x0075F773 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x0075F839 | `HandleVideosHilited` | Known | Event handler |
| 0x0075FE7B | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x00760A02 | `HandleMainMenu` | Known | Event handler |
| 0x00760A3A | `HandleMusicMenu` | Known | Event handler |
| 0x00760F62 | `HandleRadioRegion` | Known | Event handler |
| 0x00761006 | `HandleLanguage` | Known | Event handler |
| 0x0076110C | `HandleNew` | Known | Event handler |
| 0x00761187 | `HandleClear` | Known | Event handler |
| 0x007611B8 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x00761274 | `HandleSelectIndexedSession` | Known | Event handler |
| 0x0076143B | `HandleBasicSelected` | Known | Event handler |
| 0x007614E1 | `HandleTimedWorkoutSelected` | Known | Event handler |
| 0x0076158E | `HandleDistanceWorkoutSelected` | Known | Event handler |
| 0x0076163E | `HandleCaloriesWorkoutSelected` | Known | Event handler |
| 0x00761A29 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x00761A7C | `HandleSelect` | Known | Event handler |
| 0x00761BA6 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x00761BE0 | `HandleEQSettingSelected` | Known | Event handler |
| 0x00761C18 | `HandleEQSettingSelected` | Known | Event handler |
| 0x00776D7C | `HandleItemSelected` | Known | Event handler |
| 0x00776EC7 | `HandleNextContact` | Known | Event handler |
| 0x00776EF3 | `HandlePreviousContact` | Known | Event handler |
| 0x00776F29 | `HandleSelectKey` | Known | Event handler |
| 0x0077753A | `HandleSelect` | Known | Event handler |
| 0x00777861 | `HandleDateChosen` | Known | Event handler |
| 0x00777897 | `HandleTimeChosen` | Known | Event handler |
| 0x007778CD | `HandleFrequencyChosen` | Known | Event handler |
| 0x00777908 | `HandleSoundChosen` | Known | Event handler |
| 0x0077793F | `HandleLabelChosen` | Known | Event handler |
| 0x00777976 | `HandleDeleteChosen` | Known | Event handler |
| 0x007779B2 | `HandleSelect` | Known | Event handler |
| 0x007779EA | `HandleSelect` | Known | Event handler |
| 0x00777D2B | `HandleLeaveAlarm` | Known | Event handler |
| 0x00777D58 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00777D87 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00777DB4 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00777EEE | `HandleSelect` | Known | Event handler |
| 0x00777F1C | `HandleSelect` | Known | Event handler |
| 0x0077807B | `HandleNextDay` | Known | Event handler |
| 0x007780A3 | `HandlePreviousDay` | Known | Event handler |
| 0x00778252 | `HandleSelect` | Known | Event handler |
| 0x0077827F | `HandleNextDay` | Known | Event handler |
| 0x007782A7 | `HandlePreviousDay` | Known | Event handler |
| 0x0077844F | `HandleNextDay` | Known | Event handler |
| 0x00778477 | `HandlePreviousDay` | Known | Event handler |
| 0x00778538 | `HandleSelect` | Known | Event handler |
| 0x00778563 | `HandleNextDay` | Known | Event handler |
| 0x0077858B | `HandlePreviousDay` | Known | Event handler |
| 0x00778702 | `HandleSelectLozinch` | Known | Event handler |
| 0x0077887A | `HandleSelectLozinch` | Known | Event handler |
| 0x00778999 | `HandleFlowNext` | Known | Event handler |
| 0x007789C7 | `HandlePlayPause` | Known | Event handler |
| 0x00778A16 | `HandleFlowPrev` | Known | Event handler |
| 0x00778A41 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x00778B35 | `HandleAlbumSelected` | Known | Event handler |
| 0x00778CD0 | `HandleFlowNext` | Known | Event handler |
| 0x00778D1E | `HandleFlowNext` | Known | Event handler |
| 0x00778D4C | `HandlePlayPause` | Known | Event handler |
| 0x00778D9B | `HandleFlowPrev` | Known | Event handler |
| 0x00778DC7 | `HandleFlowPrev` | Known | Event handler |
| 0x00778DE7 | `HandleFlowWheel` | Known | Event handler |
| 0x00779177 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x007795A2 | `HandleArrowDown` | Known | Event handler |
| 0x0077960C | `HandleArrowUp` | Known | Event handler |
| 0x0077962B | `HandleWheel` | Known | Event handler |
| 0x007796B4 | `HandleSelect` | Known | Event handler |
| 0x00779731 | `HandleGameHilited` | Known | Event handler |
| 0x0077CB97 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077E7D3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078040F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078204B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00783C87 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007858C3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007874FF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078913B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078AD77 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078C9B3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078E5EF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079022B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00791E67 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00793AA3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007956DF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079731B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00798F57 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079AB93 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079C7CF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079E40B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A0047 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A1C83 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A38BF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A54FB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A7137 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A8D73 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007AA9AF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007AC5EB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007AE227 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007AFE63 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B1A9F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B36DB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B5317 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B6F53 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B8B8F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BA7CB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BC407 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BE028 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BEBB0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BF738 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C02C0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C0E48 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C19D0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C2558 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C30E0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C3C68 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C47F0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C5378 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C5F00 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C6A88 | `HandlePlayPause` | Known | Event handler |
| 0x007C6ABE | `HandleAddToOTG` | Known | Event handler |
| 0x007C6C5B | `HandlePlayPause` | Known | Event handler |
| 0x007C6C82 | `HandleSelect` | Known | Event handler |
| 0x007C6CAF | `HandleHilite` | Known | Event handler |
| 0x007C6CE0 | `HandlePlayPause` | Known | Event handler |
| 0x007C6D73 | `HandlePlayPause` | Known | Event handler |
| 0x007C6D9A | `HandleSelect` | Known | Event handler |
| 0x007C6E00 | `HandleHilite` | Known | Event handler |
| 0x007C6E32 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x007C6E7C | `HandlePlayPause` | Known | Event handler |
| 0x007C6EB2 | `HandleAddToOTG` | Known | Event handler |
| 0x007C6F44 | `HandlePlayPause` | Known | Event handler |
| 0x007C6F6B | `HandleSelect` | Known | Event handler |
| 0x007C6FD4 | `HandlePlayPause` | Known | Event handler |
| 0x007C700A | `HandleAddToOTG` | Known | Event handler |
| 0x007C709C | `HandlePlayPause` | Known | Event handler |
| 0x007C70C3 | `HandleSelect` | Known | Event handler |
| 0x007C712C | `HandlePlayPause` | Known | Event handler |
| 0x007C71B2 | `HandleSelect` | Known | Event handler |
| 0x007C7217 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C7258 | `HandlePlayPause` | Known | Event handler |
| 0x007C728E | `HandleAddToOTG` | Known | Event handler |
| 0x007C74C0 | `HandlePlayPause` | Known | Event handler |
| 0x007C74E7 | `HandleSelect` | Known | Event handler |
| 0x007C7514 | `HandleHilite` | Known | Event handler |
| 0x007C7544 | `HandlePlayPause` | Known | Event handler |
| 0x007C757A | `HandleAddToOTG` | Known | Event handler |
| 0x007C77AC | `HandlePlayPause` | Known | Event handler |
| 0x007C77D3 | `HandleSelect` | Known | Event handler |
| 0x007C7800 | `HandleHilite` | Known | Event handler |
| 0x007C7830 | `HandlePlayPause` | Known | Event handler |
| 0x007C7866 | `HandleAddToOTG` | Known | Event handler |
| 0x007C7B51 | `HandlePlayPause` | Known | Event handler |
| 0x007C7B78 | `HandleSelect` | Known | Event handler |
| 0x007C7BA8 | `HandlePlayPause` | Known | Event handler |
| 0x007C7BDE | `HandleAddToOTG` | Known | Event handler |
| 0x007C7C70 | `HandlePlayPause` | Known | Event handler |
| 0x007C7C97 | `HandleSelect` | Known | Event handler |
| 0x007C7D28 | `HandlePlayPause` | Known | Event handler |
| 0x007C7D5E | `HandleAddToOTG` | Known | Event handler |
| 0x007C7F17 | `HandlePlayPause` | Known | Event handler |
| 0x007C7F3E | `HandleSelect` | Known | Event handler |
| 0x007C7F70 | `HandlePlayPause` | Known | Event handler |
| 0x007C7FA6 | `HandleAddToOTG` | Known | Event handler |
| 0x007C802B | `HandleSelect` | Known | Event handler |
| 0x007C80C4 | `HandleHilite` | Known | Event handler |
| 0x007C80F0 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C8134 | `HandlePlayPause` | Known | Event handler |
| 0x007C816A | `HandleAddToOTG` | Known | Event handler |
| 0x007C81EF | `HandleSelect` | Known | Event handler |
| 0x007C8254 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C8298 | `HandlePlayPause` | Known | Event handler |
| 0x007C843C | `HandleSelect` | Known | Event handler |
| 0x007C8469 | `HandleHilite` | Known | Event handler |
| 0x007C8495 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C84D8 | `HandlePlayPause` | Known | Event handler |
| 0x007C855E | `HandleSelect` | Known | Event handler |
| 0x007C85EC | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C8630 | `HandlePlayPause` | Known | Event handler |
| 0x007C86B6 | `HandleSelect` | Known | Event handler |
| 0x007C871B | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C875C | `HandlePlayPause` | Known | Event handler |
| 0x007C87E2 | `HandleSelect` | Known | Event handler |
| 0x007C8848 | `HandleHilite` | Known | Event handler |
| 0x007C8874 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C88B8 | `HandlePlayPause` | Known | Event handler |
| 0x007C88EE | `HandleAddToOTG` | Known | Event handler |
| 0x007C8AB1 | `HandlePlayPause` | Known | Event handler |
| 0x007C8AD8 | `HandleSelect` | Known | Event handler |
| 0x007C8B08 | `HandlePlayPause` | Known | Event handler |
| 0x007C8B3E | `HandleAddToOTG` | Known | Event handler |
| 0x007C8D5F | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x007C8E78 | `HandlePlayPause` | Known | Event handler |
| 0x007C8FA5 | `HandleSelect` | Known | Event handler |
| 0x007C8FD1 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C9014 | `HandlePlayPause` | Known | Event handler |
| 0x007C909A | `HandleSelect` | Known | Event handler |
| 0x007C90C7 | `HandleHilite` | Known | Event handler |
| 0x007C90F3 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C9134 | `HandlePlayPause` | Known | Event handler |
| 0x007C9267 | `HandleSelect` | Known | Event handler |
| 0x007C9293 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C9BA5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CA45D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CAD15 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CB5CD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CBE85 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CC73D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CCFF5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CD8AD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CD8F6 | `HandleTVOutChanged` | Known | Event handler |
| 0x007CD92E | `HandleTVSignalChanged` | Known | Event handler |
| 0x007CD969 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x007CD9BA | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x007CD9FF | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x007CDA48 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x007CDA8A | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x007CDACD | `HandleSelect` | Known | Event handler |
| 0x007CDAFD | `HandleSelect` | Known | Event handler |
| 0x007CDB35 | `HandleMenuLongpress` | Known | Event handler |
| 0x007CDB63 | `HandleMenuKey` | Known | Event handler |
| 0x007CDBE9 | `HandlePlayPause` | Known | Event handler |
| 0x007CDC69 | `HandleSelect` | Known | Event handler |
| 0x007CE576 | `HandlePlayPause` | Known | Event handler |
| 0x007CE5EB | `HandleWheelProgress` | Known | Event handler |
| 0x007CE629 | `HandleMenuLongpress` | Known | Event handler |
| 0x007CE657 | `HandleMenuKey` | Known | Event handler |
| 0x007CE6DD | `HandlePlayPause` | Known | Event handler |
| 0x007CE75D | `HandleSelectProgress` | Known | Event handler |
| 0x007CF072 | `HandlePlayPause` | Known | Event handler |
| 0x007CF0E7 | `HandleWheelProgress` | Known | Event handler |
| 0x007CF125 | `HandleMenuLongpress` | Known | Event handler |
| 0x007CF153 | `HandleMenuKey` | Known | Event handler |
| 0x007CF1D9 | `HandlePlayPause` | Known | Event handler |
| 0x007CF259 | `HandleSelectVolume` | Known | Event handler |
| 0x007CFB6C | `HandlePlayPause` | Known | Event handler |
| 0x007CFBE1 | `HandleWheelVolume` | Known | Event handler |
| 0x007CFC1D | `HandleMenuLongpress` | Known | Event handler |
| 0x007CFC4B | `HandleMenuKey` | Known | Event handler |
| 0x007CFCD1 | `HandlePlayPause` | Known | Event handler |
| 0x007CFD51 | `HandleSelectRating` | Known | Event handler |
| 0x007D0664 | `HandlePlayPause` | Known | Event handler |
| 0x007D06D9 | `HandleWheelRating` | Known | Event handler |
| 0x007D0715 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D0743 | `HandleMenuKey` | Known | Event handler |
| 0x007D07BB | `HandlePlayPause` | Known | Event handler |
| 0x007D0832 | `HandleSelectScrub` | Known | Event handler |
| 0x007D1136 | `HandlePlayPause` | Known | Event handler |
| 0x007D11A2 | `HandleWheelScrub` | Known | Event handler |
| 0x007D11DD | `HandleMenuLongpress` | Known | Event handler |
| 0x007D120B | `HandleMenuKey` | Known | Event handler |
| 0x007D1268 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007D12A0 | `HandlePlayPause` | Known | Event handler |
| 0x007D12FA | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007D132F | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x007D1C49 | `HandlePlayPause` | Known | Event handler |
| 0x007D1CBE | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007D1D01 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D1D2F | `HandleMenuKey` | Known | Event handler |
| 0x007D1DB5 | `HandlePlayPause` | Known | Event handler |
| 0x007D1E35 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007D274B | `HandlePlayPause` | Known | Event handler |
| 0x007D27E9 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D2817 | `HandleMenuKey` | Known | Event handler |
| 0x007D289D | `HandlePlayPause` | Known | Event handler |
| 0x007D291D | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007D3233 | `HandlePlayPause` | Known | Event handler |
| 0x007D32D1 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D32FF | `HandleMenuKey` | Known | Event handler |
| 0x007D3385 | `HandlePlayPause` | Known | Event handler |
| 0x007D3405 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007D3D1B | `HandlePlayPause` | Known | Event handler |
| 0x007D3DB9 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D3DE7 | `HandleMenuKey` | Known | Event handler |
| 0x007D3E6D | `HandlePlayPause` | Known | Event handler |
| 0x007D3EED | `HandleSelectChapterArt` | Known | Event handler |
| 0x007D4804 | `HandlePlayPause` | Known | Event handler |
| 0x007D4879 | `HandleWheelVolume` | Known | Event handler |
| 0x007D48B5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D48E3 | `HandleMenuKey` | Known | Event handler |
| 0x007D4972 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007D4A09 | `HandleSelect` | Known | Event handler |
| 0x007D531F | `HandlePlayPause` | Known | Event handler |
| 0x007D539D | `HandleWheel` | Known | Event handler |
| 0x007D53D1 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D53FF | `HandleMenuKey` | Known | Event handler |
| 0x007D548E | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007D5525 | `HandleSelect` | Known | Event handler |
| 0x007D5E3B | `HandlePlayPause` | Known | Event handler |
| 0x007D5EB9 | `HandleWheel` | Known | Event handler |
| 0x007D5EED | `HandleMenuLongpress` | Known | Event handler |
| 0x007D5F1B | `HandleMenuKey` | Known | Event handler |
| 0x007D5FA1 | `HandlePlayPause` | Known | Event handler |
| 0x007D6021 | `HandleSelect` | Known | Event handler |
| 0x007D692E | `HandlePlayPause` | Known | Event handler |
| 0x007D69A3 | `HandleWheel` | Known | Event handler |
| 0x007D69D9 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D6A07 | `HandleMenuKey` | Known | Event handler |
| 0x007D6A8D | `HandlePlayPause` | Known | Event handler |
| 0x007D6B0D | `HandleSelectProgress` | Known | Event handler |
| 0x007D7422 | `HandlePlayPause` | Known | Event handler |
| 0x007D7497 | `HandleWheelProgress` | Known | Event handler |
| 0x007D74D5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D7503 | `HandleMenuKey` | Known | Event handler |
| 0x007D757B | `HandlePlayPause` | Known | Event handler |
| 0x007D75F2 | `HandleSelectScrub` | Known | Event handler |
| 0x007D7EF6 | `HandlePlayPause` | Known | Event handler |
| 0x007D7F62 | `HandleWheelScrub` | Known | Event handler |
| 0x007D7F9D | `HandleMenuLongpress` | Known | Event handler |
| 0x007D7FCB | `HandleMenuKey` | Known | Event handler |
| 0x007D8051 | `HandlePlayPause` | Known | Event handler |
| 0x007D89DD | `HandlePlayPause` | Known | Event handler |
| 0x007D8A52 | `HandleWheelVolume` | Known | Event handler |
| 0x007D8A8D | `HandleMenuLongpress` | Known | Event handler |
| 0x007D8ABB | `HandleMenuKey` | Known | Event handler |
| 0x007D8B41 | `HandlePlayPause` | Known | Event handler |
| 0x007D94CD | `HandlePlayPause` | Known | Event handler |
| 0x007D9542 | `HandleWheelBrightness` | Known | Event handler |
| 0x007D9659 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007D9FAC | `HandleWheel` | Known | Event handler |
| 0x007D9FE1 | `HandleMenuLongpress` | Known | Event handler |
| 0x007DA00F | `HandleMenuKey` | Known | Event handler |
| 0x007DA095 | `HandlePlayPause` | Known | Event handler |
| 0x007DA115 | `HandleSelect` | Known | Event handler |
| 0x007DA5B7 | `HandlePlayPause` | Known | Event handler |
| 0x007DA645 | `HandleMenuLongpress` | Known | Event handler |
| 0x007DA673 | `HandleMenuKey` | Known | Event handler |
| 0x007DA6F9 | `HandlePlayPause` | Known | Event handler |
| 0x007DA779 | `HandleSelectProgress` | Known | Event handler |
| 0x007DAC23 | `HandlePlayPause` | Known | Event handler |
| 0x007DAC98 | `HandleWheelProgress` | Known | Event handler |
| 0x007DACD5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007DAD03 | `HandleMenuKey` | Known | Event handler |
| 0x007DAD89 | `HandlePlayPause` | Known | Event handler |
| 0x007DAE09 | `HandleSelectProgress` | Known | Event handler |
| 0x007DB2B3 | `HandlePlayPause` | Known | Event handler |
| 0x007DB328 | `HandleWheelProgress` | Known | Event handler |
| 0x007DB365 | `HandleMenuLongpress` | Known | Event handler |
| 0x007DB393 | `HandleMenuKey` | Known | Event handler |
| 0x007DB419 | `HandlePlayPause` | Known | Event handler |
| 0x007DB499 | `HandleSelectProgress` | Known | Event handler |
| 0x007DB8CF | `HandlePlayPause` | Known | Event handler |
| 0x007DB944 | `HandleWheelProgress` | Known | Event handler |
| 0x007DB981 | `HandleMenuLongpress` | Known | Event handler |
| 0x007DB9AF | `HandleMenuKey` | Known | Event handler |
| 0x007DBA1C | `HandlePlayPause` | Known | Event handler |
| 0x007DBA88 | `HandleSelectScrub` | Known | Event handler |
| 0x007DBEA2 | `HandlePlayPause` | Known | Event handler |
| 0x007DBF03 | `HandleWheelScrub` | Known | Event handler |
| 0x007DBF3D | `HandleMenuLongpress` | Known | Event handler |
| 0x007DBF6B | `HandleMenuKey` | Known | Event handler |
| 0x007DBFF1 | `HandlePlayPause` | Known | Event handler |
| 0x007DC071 | `HandleSelectVolume` | Known | Event handler |
| 0x007DC4A5 | `HandlePlayPause` | Known | Event handler |
| 0x007DC51A | `HandleWheelVolume` | Known | Event handler |
| 0x007DC62D | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007DCACC | `HandleSelect` | Known | Event handler |
| 0x007DCAF9 | `HandleSelect` | Known | Event handler |
| 0x007DCB29 | `HandleSelect` | Known | Event handler |
| 0x007DCB59 | `HandleSelect` | Known | Event handler |
| 0x007DCB89 | `HandleSelect` | Known | Event handler |
| 0x007DCBB9 | `HandleSelect` | Known | Event handler |
| 0x007DCBE9 | `HandleSelect` | Known | Event handler |
| 0x007DCC19 | `HandleSelect` | Known | Event handler |
| 0x007DCC49 | `HandleSelect` | Known | Event handler |
| 0x007DCCB9 | `HandleSelect` | Known | Event handler |
| 0x007DCCE9 | `HandleSelect` | Known | Event handler |
| 0x007DCD61 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DCD94 | `HandleNotesPop` | Known | Event handler |
| 0x007DCE11 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DCE44 | `HandleNotesPop` | Known | Event handler |
| 0x007DD300 | `HandleNotesSelected` | Known | Event handler |
| 0x007DD33D | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DD370 | `HandleNotesPop` | Known | Event handler |
| 0x007DD82C | `HandleNotesSelected` | Known | Event handler |
| 0x007DD869 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DD89C | `HandleNotesPop` | Known | Event handler |
| 0x007DD8C7 | `HandleNotesSelected` | Known | Event handler |
| 0x007DDD99 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DDDCC | `HandleNotesPop` | Known | Event handler |
| 0x007DDDF7 | `HandleNotesSelected` | Known | Event handler |
| 0x007DE2C9 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DE2FC | `HandleNotesPop` | Known | Event handler |
| 0x007DE379 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DE3AC | `HandleNotesPop` | Known | Event handler |
| 0x007DE429 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DE45C | `HandleNotesPop` | Known | Event handler |
| 0x007DE4D4 | `HandlePlayPause` | Known | Event handler |
| 0x007DE4FD | `HandlePlayPause` | Known | Event handler |
| 0x007DE52B | `HandlePlayPause` | Known | Event handler |
| 0x007DE560 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007DE5E0 | `HandleHiliteAlbum` | Known | Event handler |
| 0x007DE689 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007DE710 | `HandleHiliteAlbum` | Known | Event handler |
| 0x007DE9D4 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x007DEA30 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x007DEBE7 | `HandleSelect` | Known | Event handler |
| 0x007DED6B | `HandleSelect` | Known | Event handler |
| 0x007DEDA5 | `HandleImageLast` | Known | Event handler |
| 0x007DEDCF | `HandleImageNext` | Known | Event handler |
| 0x007DEDFE | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007DEE38 | `HandleImageFirst` | Known | Event handler |
| 0x007DEE63 | `HandleImagePrev` | Known | Event handler |
| 0x007DEE8F | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007DEEBE | `HandleImageNext` | Known | Event handler |
| 0x007DEEE7 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007DEF1B | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007DEF4A | `HandleImagePrev` | Known | Event handler |
| 0x007DEF6B | `HandleImageWheel` | Known | Event handler |
| 0x007DF009 | `HandleImageNext` | Known | Event handler |
| 0x007DF038 | `HandlePlayPause` | Known | Event handler |
| 0x007DF087 | `HandleImagePrev` | Known | Event handler |
| 0x007DF0B3 | `HandleSelect` | Known | Event handler |
| 0x007DF383 | `HandleImageNext` | Known | Event handler |
| 0x007DF3AD | `HandlePause` | Known | Event handler |
| 0x007DF3D2 | `HandlePlay` | Known | Event handler |
| 0x007DF3FB | `HandlePlayPause` | Known | Event handler |
| 0x007DF424 | `HandleImagePrev` | Known | Event handler |
| 0x007DF47D | `HandleWheel` | Known | Event handler |
| 0x007DF515 | `HandleImageNext` | Known | Event handler |
| 0x007DF544 | `HandlePlayPause` | Known | Event handler |
| 0x007DF593 | `HandleImagePrev` | Known | Event handler |
| 0x007DF5BF | `HandleSelect` | Known | Event handler |
| 0x007DF88F | `HandleImageNext` | Known | Event handler |
| 0x007DF8B9 | `HandlePause` | Known | Event handler |
| 0x007DF8DE | `HandlePlay` | Known | Event handler |
| 0x007DF907 | `HandlePlayPause` | Known | Event handler |
| 0x007DF930 | `HandleImagePrev` | Known | Event handler |
| 0x007DF989 | `HandleWheel` | Known | Event handler |
| 0x007DFA21 | `HandleImageNext` | Known | Event handler |
| 0x007DFA50 | `HandlePlayPause` | Known | Event handler |
| 0x007DFA9F | `HandleImagePrev` | Known | Event handler |
| 0x007DFACB | `HandleSelect` | Known | Event handler |
| 0x007DFD9B | `HandleImageNext` | Known | Event handler |
| 0x007DFDC5 | `HandlePause` | Known | Event handler |
| 0x007DFDEA | `HandlePlay` | Known | Event handler |
| 0x007DFE13 | `HandlePlayPause` | Known | Event handler |
| 0x007DFE3C | `HandleImagePrev` | Known | Event handler |
| 0x007DFE95 | `HandleWheel` | Known | Event handler |
| 0x007DFF2D | `HandleImageNext` | Known | Event handler |
| 0x007DFF5C | `HandlePlayPause` | Known | Event handler |
| 0x007DFFAB | `HandleImagePrev` | Known | Event handler |
| 0x007DFFD7 | `HandleSelect` | Known | Event handler |
| 0x007E02A7 | `HandleImageNext` | Known | Event handler |
| 0x007E02D1 | `HandlePause` | Known | Event handler |
| 0x007E02F6 | `HandlePlay` | Known | Event handler |
| 0x007E031F | `HandlePlayPause` | Known | Event handler |
| 0x007E0348 | `HandleImagePrev` | Known | Event handler |
| 0x007E03A1 | `HandleWheel` | Known | Event handler |
| 0x007E0439 | `HandleImageNext` | Known | Event handler |
| 0x007E0468 | `HandlePlayPause` | Known | Event handler |
| 0x007E04B7 | `HandleImagePrev` | Known | Event handler |
| 0x007E04E3 | `HandleSelect` | Known | Event handler |
| 0x007E07B3 | `HandleImageNext` | Known | Event handler |
| 0x007E07DD | `HandlePause` | Known | Event handler |
| 0x007E0802 | `HandlePlay` | Known | Event handler |
| 0x007E082B | `HandlePlayPause` | Known | Event handler |
| 0x007E0854 | `HandleImagePrev` | Known | Event handler |
| 0x007E08AD | `HandleWheel` | Known | Event handler |
| 0x007E0945 | `HandleImageNext` | Known | Event handler |
| 0x007E0974 | `HandlePlayPause` | Known | Event handler |
| 0x007E09C3 | `HandleImagePrev` | Known | Event handler |
| 0x007E09EF | `HandleSelect` | Known | Event handler |
| 0x007E0CBF | `HandleImageNext` | Known | Event handler |
| 0x007E0CE9 | `HandlePause` | Known | Event handler |
| 0x007E0D0E | `HandlePlay` | Known | Event handler |
| 0x007E0D37 | `HandlePlayPause` | Known | Event handler |
| 0x007E0D60 | `HandleImagePrev` | Known | Event handler |
| 0x007E0DB9 | `HandleWheel` | Known | Event handler |
| 0x007E0E51 | `HandleImageNext` | Known | Event handler |
| 0x007E0E80 | `HandlePlayPause` | Known | Event handler |
| 0x007E0ECF | `HandleImagePrev` | Known | Event handler |
| 0x007E0EFB | `HandleSelect` | Known | Event handler |
| 0x007E1146 | `HandleImageNext` | Known | Event handler |
| 0x007E1170 | `HandlePause` | Known | Event handler |
| 0x007E1195 | `HandlePlay` | Known | Event handler |
| 0x007E11BE | `HandlePlayPause` | Known | Event handler |
| 0x007E11E7 | `HandleImagePrev` | Known | Event handler |
| 0x007E1250 | `HandleWheel` | Known | Event handler |
| 0x007E12E9 | `HandleImageNext` | Known | Event handler |
| 0x007E1318 | `HandlePlayPause` | Known | Event handler |
| 0x007E1367 | `HandleImagePrev` | Known | Event handler |
| 0x007E1393 | `HandleSelect` | Known | Event handler |
| 0x007E15DE | `HandleImageNext` | Known | Event handler |
| 0x007E1608 | `HandlePause` | Known | Event handler |
| 0x007E162D | `HandlePlay` | Known | Event handler |
| 0x007E1656 | `HandlePlayPause` | Known | Event handler |
| 0x007E167F | `HandleImagePrev` | Known | Event handler |
| 0x007E16E8 | `HandleWheel` | Known | Event handler |
| 0x007E1781 | `HandleImageNext` | Known | Event handler |
| 0x007E17B0 | `HandlePlayPause` | Known | Event handler |
| 0x007E17FF | `HandleImagePrev` | Known | Event handler |
| 0x007E182B | `HandleSelect` | Known | Event handler |
| 0x007E1A76 | `HandleImageNext` | Known | Event handler |
| 0x007E1AA0 | `HandlePause` | Known | Event handler |
| 0x007E1AC5 | `HandlePlay` | Known | Event handler |
| 0x007E1AEE | `HandlePlayPause` | Known | Event handler |
| 0x007E1B17 | `HandleImagePrev` | Known | Event handler |
| 0x007E1B80 | `HandleWheel` | Known | Event handler |
| 0x007E1C19 | `HandleImageNext` | Known | Event handler |
| 0x007E1C48 | `HandlePlayPause` | Known | Event handler |
| 0x007E1C97 | `HandleImagePrev` | Known | Event handler |
| 0x007E1CC3 | `HandleSelect` | Known | Event handler |
| 0x007E1F0E | `HandleImageNext` | Known | Event handler |
| 0x007E1F38 | `HandlePause` | Known | Event handler |
| 0x007E1F5D | `HandlePlay` | Known | Event handler |
| 0x007E1F86 | `HandlePlayPause` | Known | Event handler |
| 0x007E1FAF | `HandleImagePrev` | Known | Event handler |
| 0x007E2018 | `HandleWheel` | Known | Event handler |
| 0x007E20B1 | `HandleImageNext` | Known | Event handler |
| 0x007E20E0 | `HandlePlayPause` | Known | Event handler |
| 0x007E212F | `HandleImagePrev` | Known | Event handler |
| 0x007E215B | `HandleSelect` | Known | Event handler |
| 0x007E23A6 | `HandleImageNext` | Known | Event handler |
| 0x007E23D0 | `HandlePause` | Known | Event handler |
| 0x007E23F5 | `HandlePlay` | Known | Event handler |
| 0x007E241E | `HandlePlayPause` | Known | Event handler |
| 0x007E2447 | `HandleImagePrev` | Known | Event handler |
| 0x007E24B0 | `HandleWheel` | Known | Event handler |
| 0x007E24DD | `HandleSelect` | Known | Event handler |
| 0x007E250D | `HandleSelect` | Known | Event handler |
| 0x007E2630 | `HandleTuning` | Known | Event handler |
| 0x007E27EC | `HandleVolumeChange` | Known | Event handler |
| 0x007E2938 | `HandleVolumeWheel` | Known | Event handler |
| 0x007E2A93 | `HandleTuningSelect` | Known | Event handler |
| 0x007E2D72 | `HandleFrequencyChange` | Known | Event handler |
| 0x007E2ECF | `HandleTuningSelect` | Known | Event handler |
| 0x007E31AE | `HandleFrequencyChange` | Known | Event handler |
| 0x007E32D8 | `HandleTimerDone` | Known | Event handler |
| 0x007E34CD | `HandleVolumeChange` | Known | Event handler |
| 0x007E35E4 | `HandleVolumeWheel` | Known | Event handler |
| 0x007E3BC7 | `HandleExitUnsupported` | Known | Event handler |
| 0x007E3BF9 | `HandleExitUnsupported` | Known | Event handler |
| 0x007E8C2D | `HandleSelectKey` | Known | Event handler |
| 0x007E8C62 | `HandleWheel` | Known | Event handler |
| 0x007E8DB0 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x007E8E03 | `HandleSelectKey` | Known | Event handler |
| 0x007E8E2B | `HandleSelectKey` | Known | Event handler |
| 0x007E8E5B | `HandleExit` | Known | Event handler |
| 0x007E8E85 | `HandleStartStop` | Known | Event handler |
| 0x007E8EEB | `HandleStartStop` | Known | Event handler |
| 0x007E9003 | `HandleExit` | Known | Event handler |
| 0x007E902D | `HandleStartStop` | Known | Event handler |
| 0x007E9059 | `HandleLap` | Known | Event handler |
| 0x007E915D | `HandleSelectLozinch` | Known | Event handler |
| 0x007E9C04 | `HandleSelect` | Known | Event handler |
| 0x007EA70F | `HandleChoosePowerPlay` | Known | Event handler |
| 0x007EA74A | `HandleChooseVoicePrompts` | Known | Event handler |
| 0x007EA788 | `HandleChooseUnit` | Known | Event handler |
| 0x007EA91C | `HandleListChoose` | Known | Event handler |
| 0x007EAB7B | `HandleSelect` | Known | Event handler |
| 0x007EAD9B | `HandleSelect` | Known | Event handler |
| 0x007EADD1 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EB002 | `HandleNowPlayingSelected` | Known | Event handler |
| 0x007EB040 | `HandleChoosePlaylistsPlay` | Known | Event handler |
| 0x007EB07F | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x007EB0BF | `HandleNoneSelected` | Known | Event handler |
| 0x007EB0F5 | `HandleBegin` | Known | Event handler |
| 0x007EB3E2 | `HandleBegin` | Known | Event handler |
| 0x007EB411 | `HandleBegin` | Known | Event handler |
| 0x007EB4CD | `HandleBegin` | Known | Event handler |
| 0x007EB4F9 | `HandleBegin` | Known | Event handler |
| 0x007EB5B5 | `HandleBegin` | Known | Event handler |
| 0x007EB5E1 | `HandleBegin` | Known | Event handler |
| 0x007EB69D | `HandleBegin` | Known | Event handler |
| 0x007EB6D1 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EB6FC | `HandleMenuKey` | Known | Event handler |
| 0x007EB793 | `HandlePauseHold` | Known | Event handler |
| 0x007EB7C2 | `HandlePauseKey` | Known | Event handler |
| 0x007EB84C | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EB886 | `HandlePowerPlay` | Known | Event handler |
| 0x007EB8B2 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EBD1F | `HandlePauseHold` | Known | Event handler |
| 0x007EBD4E | `HandlePauseKey` | Known | Event handler |
| 0x007EBD79 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EBDB7 | `HandlePowerPlay` | Known | Event handler |
| 0x007EBDE6 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EBE0C | `HandleWheel` | Known | Event handler |
| 0x007EBE41 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EBE6C | `HandleMenuKey` | Known | Event handler |
| 0x007EBF03 | `HandlePauseHold` | Known | Event handler |
| 0x007EBF32 | `HandlePauseKey` | Known | Event handler |
| 0x007EBFBC | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EBFEC | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EC44C | `HandlePauseHold` | Known | Event handler |
| 0x007EC47B | `HandlePauseKey` | Known | Event handler |
| 0x007EC4A6 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EC4DA | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EC500 | `HandleWheel` | Known | Event handler |
| 0x007EC535 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EC560 | `HandleMenuKey` | Known | Event handler |
| 0x007EC5F7 | `HandlePauseHold` | Known | Event handler |
| 0x007EC626 | `HandlePauseKey` | Known | Event handler |
| 0x007EC6B0 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EC6EA | `HandlePowerPlay` | Known | Event handler |
| 0x007EC716 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ECB82 | `HandlePauseHold` | Known | Event handler |
| 0x007ECBB1 | `HandlePauseKey` | Known | Event handler |
| 0x007ECBDC | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ECC1A | `HandlePowerPlay` | Known | Event handler |
| 0x007ECC49 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ECC6F | `HandleWheel` | Known | Event handler |
| 0x007ECCA5 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007ECCD0 | `HandleMenuKey` | Known | Event handler |
| 0x007ECD67 | `HandlePauseHold` | Known | Event handler |
| 0x007ECD96 | `HandlePauseKey` | Known | Event handler |
| 0x007ECE20 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007ECE50 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ED2AF | `HandlePauseHold` | Known | Event handler |
| 0x007ED2DE | `HandlePauseKey` | Known | Event handler |
| 0x007ED309 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ED33D | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ED363 | `HandleWheel` | Known | Event handler |
| 0x007ED399 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007ED3C4 | `HandleMenuKey` | Known | Event handler |
| 0x007ED45B | `HandlePauseHold` | Known | Event handler |
| 0x007ED48A | `HandlePauseKey` | Known | Event handler |
| 0x007ED514 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007ED54E | `HandlePowerPlay` | Known | Event handler |
| 0x007ED57A | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ED9EA | `HandlePauseHold` | Known | Event handler |
| 0x007EDA19 | `HandlePauseKey` | Known | Event handler |
| 0x007EDA44 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EDA82 | `HandlePowerPlay` | Known | Event handler |
| 0x007EDAB1 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EDAD7 | `HandleWheel` | Known | Event handler |
| 0x007EDB0D | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EDB38 | `HandleMenuKey` | Known | Event handler |
| 0x007EDBCF | `HandlePauseHold` | Known | Event handler |
| 0x007EDBFE | `HandlePauseKey` | Known | Event handler |
| 0x007EDC88 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EDCB8 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EE11B | `HandlePauseHold` | Known | Event handler |
| 0x007EE14A | `HandlePauseKey` | Known | Event handler |
| 0x007EE175 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EE1A9 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EE1CF | `HandleWheel` | Known | Event handler |
| 0x007EE205 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EE230 | `HandleMenuKey` | Known | Event handler |
| 0x007EE2C7 | `HandlePauseHold` | Known | Event handler |
| 0x007EE2F6 | `HandlePauseKey` | Known | Event handler |
| 0x007EE380 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EE3BA | `HandlePowerPlay` | Known | Event handler |
| 0x007EE3E6 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EE856 | `HandlePauseHold` | Known | Event handler |
| 0x007EE885 | `HandlePauseKey` | Known | Event handler |
| 0x007EE8B0 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EE8EE | `HandlePowerPlay` | Known | Event handler |
| 0x007EE91D | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EE943 | `HandleWheel` | Known | Event handler |
| 0x007EE979 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EE9A4 | `HandleMenuKey` | Known | Event handler |
| 0x007EEA3B | `HandlePauseHold` | Known | Event handler |
| 0x007EEA6A | `HandlePauseKey` | Known | Event handler |
| 0x007EEAF4 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EEB24 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EEF87 | `HandlePauseHold` | Known | Event handler |
| 0x007EEFB6 | `HandlePauseKey` | Known | Event handler |
| 0x007EEFE1 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EF015 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EF03B | `HandleWheel` | Known | Event handler |
| 0x007EF071 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EF09C | `HandleMenuKey` | Known | Event handler |
| 0x007EF133 | `HandlePauseHold` | Known | Event handler |
| 0x007EF162 | `HandlePauseKey` | Known | Event handler |
| 0x007EF1EC | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EF21C | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EF615 | `HandlePauseHold` | Known | Event handler |
| 0x007EF644 | `HandlePauseKey` | Known | Event handler |
| 0x007EF66F | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EF6A3 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EF6C9 | `HandleWheel` | Known | Event handler |
| 0x007EF6FD | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EF728 | `HandleResumeWorkout` | Known | Event handler |
| 0x007EF803 | `HandleResumeWorkout` | Known | Event handler |
| 0x007EF877 | `HandlePauseWorkout` | Known | Event handler |
| 0x007EF8E5 | `HandleChooseMusic` | Known | Event handler |
| 0x007EF982 | `HandleEndWorkout` | Known | Event handler |
| 0x007EFA2D | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EFCD4 | `HandleEndWorkout` | Known | Event handler |
| 0x007F0163 | `HandleSelectResume` | Known | Event handler |
| 0x007F019B | `HandleEndWorkout` | Known | Event handler |
| 0x007F0246 | `HandleChooseCustomTimedWorkout` | Known | Event handler |
| 0x007F02DF | `HandleChooseTimedWorkoutFromList` | Known | Event handler |
| 0x007F0392 | `HandleChooseCustomDistanceWorkout` | Known | Event handler |
| 0x007F0432 | `HandleChooseDistanceWorkoutFromList` | Known | Event handler |
| 0x007F0618 | `HandleChooseCustomCalorieWorkout` | Known | Event handler |
| 0x007F06B7 | `HandleChooseCalorieWorkoutFromList` | Known | Event handler |
| 0x007F0A26 | `HandleChooseLink` | Known | Event handler |
| 0x007F0A5C | `HandleChooseCalibrate` | Known | Event handler |
| 0x007F0DB5 | `HandleChooseRemoteSetting` | Known | Event handler |
| 0x007F0DF4 | `HandleChooseRemoteLink` | Known | Event handler |
| 0x007F0E30 | `HandleChooseRemoteUnlink` | Known | Event handler |
| 0x007F11DE | `Handle400MetersWalk` | Known | Event handler |
| 0x007F1217 | `HandleCustomWalk` | Known | Event handler |
| 0x007F12ED | `HandleSelectWalking` | Known | Event handler |
| 0x007F1411 | `HandleSelectRunning` | Known | Event handler |
| 0x007F175E | `Handle400MetersRun` | Known | Event handler |
| 0x007F1796 | `HandleCustomRun` | Known | Event handler |
| 0x007F1A69 | `HandleSelect` | Known | Event handler |
| 0x007F1A99 | `HandleSelect` | Known | Event handler |
| 0x007F1C0F | `HandleLinkNewRemote` | Known | Event handler |
| 0x007F1D7D | `HandleSelect` | Known | Event handler |
| 0x007F1DAB | `HandleCancelRemoteLinking` | Known | Event handler |
| 0x007F1E18 | `HandleSelect` | Known | Event handler |
| 0x007F2309 | `HandleUnlinkRemote` | Known | Event handler |
| 0x007F256D | `HandleWeightSelect` | Known | Event handler |
| 0x007F25CA | `HandleWeightWheel` | Known | Event handler |
| 0x007F25FD | `HandleWeightSelect` | Known | Event handler |
| 0x007F2687 | `HandleWeightWheel` | Known | Event handler |
| 0x007F26B9 | `HandleDistanceSelect` | Known | Event handler |
| 0x007F2745 | `HandleDistanceWheel` | Known | Event handler |
| 0x007F2779 | `HandleDistanceSelect` | Known | Event handler |
| 0x007F2805 | `HandleDistanceWheel` | Known | Event handler |
| 0x007F2839 | `HandleTimeSelect` | Known | Event handler |
| 0x007F28C1 | `HandleTimeWheel` | Known | Event handler |
| 0x007F28F1 | `HandleCaloriesSelect` | Known | Event handler |
| 0x007F2A49 | `HandleCaloriesWheel` | Known | Event handler |
| 0x007F2DB5 | `HandleChooseLast` | Known | Event handler |
| 0x007F2DEB | `HandleChooseRecent` | Known | Event handler |
| 0x007F2E23 | `HandleChooseBest` | Known | Event handler |
| 0x007F3139 | `HandleSelect` | Known | Event handler |
| 0x007F3321 | `HandleSelectFromLastWorkout` | Known | Event handler |
| 0x007F3519 | `HandleSelect` | Known | Event handler |
| 0x007F37D2 | `HandleChooseRecentWorkout` | Known | Event handler |
| 0x007F38A5 | `HandleSelect` | Known | Event handler |
| 0x007F3939 | `HandleSelect_Basic` | Known | Event handler |
| 0x007F3C1D | `HandleSelect_Dynamic` | Known | Event handler |
| 0x007F3F11 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x007F4201 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x007F45F2 | `HandleSelect` | Known | Event handler |
| 0x007F467E | `HandleSelect` | Known | Event handler |
| 0x007F470C | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x007F49F6 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x007F4AD7 | `HandlePlayPause` | Known | Event handler |
| 0x007F4B65 | `HandlePlayPause` | Known | Event handler |
| 0x007F4BF5 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x007F4C2D | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x007F4C69 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x007F4CAC | `HandlePlayPause` | Known | Event handler |
| 0x007F4CE2 | `HandleAddToOTG` | Known | Event handler |
| 0x007F4F37 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007F5193 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x00818392 | `HandleSelectClock` | Known | Event handler |
| 0x008183CB | `HandleHilited` | Known | Event handler |
| 0x008183FD | `HandleWheel` | Known | Event handler |
| 0x00818444 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x008184C9 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x008186D5 | `HandleImageLast` | Known | Event handler |
| 0x008186FF | `HandleScreenNext` | Known | Event handler |
| 0x0081872F | `HandleBrowseSlideshow` | Known | Event handler |
| 0x00818769 | `HandleImageFirst` | Known | Event handler |
| 0x00818794 | `HandleScreenPrev` | Known | Event handler |
| 0x008187C1 | `HandleBrowseLarge` | Known | Event handler |
| 0x00818841 | `HandleImageNext` | Known | Event handler |
| 0x0081886A | `HandleBrowseSlideshow` | Known | Event handler |
| 0x0081889E | `HandleBrowseSlideshow` | Known | Event handler |
| 0x008188CD | `HandleImagePrev` | Known | Event handler |
| 0x008188FB | `HandleBrowseSmall` | Known | Event handler |

---

## 6. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001030CC | `GotoNowPlaying` | Known | Navigation |
| 0x00103144 | `GotoMainMenu` | Known | Navigation |
| 0x00121638 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x00121650 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x001217C8 | `GotoScreen_AddressBook` | Known | Navigation |
| 0x0012CC34 | `GotoNowPlaying` | Known | Navigation |
| 0x0012CC48 | `GotoAlbums` | Known | Navigation |
| 0x0012CC54 | `GotoSongs` | Known | Navigation |
| 0x0013A744 | `GotoScreen_EnterPassKey` | Known | Navigation |
| 0x0013A75C | `GotoScreen_LockediPod` | Known | Navigation |
| 0x0013B160 | `GotoScreen_MainMenu` | Known | Navigation |
| 0x001547B8 | `GotoMainMenu` | Known | Navigation |
| 0x001DEFEC | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001EA57C | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001EADCC | `GotoPlayDeleteMenu` | Known | Navigation |
| 0x001EAE50 | `GotoNowPlaying` | Known | Navigation |
| 0x0020A044 | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation |
| 0x0021784C | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x00217944 | `GotoScreen_SettingsMenu` | Known | Navigation |
| 0x00220808 | `GotoDefaultLayout` | Known | Navigation |
| 0x0022088C | `GotoVolumeLayout` | Known | Navigation |
| 0x002209C4 | `GotoProgressLayout` | Known | Navigation |
| 0x00220CE0 | `GotoDefault` | Known | Navigation |
| 0x00221014 | `GotoProgressLayout` | Known | Navigation |
| 0x002211D4 | `GotoRentalWarningLayout` | Known | Navigation |
| 0x00221258 | `GotoProgressLayout` | Known | Navigation |
| 0x00221568 | `GotoProgressLayout` | Known | Navigation |
| 0x002230CC | `GotoNowPlaying` | Known | Navigation |
| 0x00223998 | `GotoNowPlaying` | Known | Navigation |
| 0x00227268 | `GotoScreen_Language` | Known | Navigation |
| 0x002275C8 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x002275E4 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x002275FC | `GotoDefaultLayout` | Known | Navigation |
| 0x00227610 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x002276A8 | `GotoVolumeLayout` | Known | Navigation |
| 0x002276BC | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x0022775C | `GotoProgressLayout` | Known | Navigation |
| 0x00227770 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00227C38 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00227F74 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x0022812C | `GotoProgressLayout` | Known | Navigation |
| 0x00228140 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00228204 | `GotoBrightnessVideoLayout` | Known | Navigation |
| 0x00228220 | `GotoRatingLayout` | Known | Navigation |
| 0x002284C4 | `GotoChapterArtLayout` | Known | Navigation |
| 0x002284DC | `GotoShuffleLayout` | Known | Navigation |
| 0x0022886C | `GotoExtraInfoLayout` | Known | Navigation |
| 0x00228880 | `GotoExtraInfoLoadingLayout` | Known | Navigation |
| 0x00228950 | `GotoVolumeLayout` | Known | Navigation |
| 0x00228968 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x002289F4 | `GotoVolumeLayout` | Known | Navigation |
| 0x00228A08 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x00228C18 | `GotoScrubLayout` | Known | Navigation |
| 0x00228C28 | `GotoScrubVideoLayout` | Known | Navigation |
| 0x00228CB8 | `GotoProgressLayout` | Known | Navigation |
| 0x00228CCC | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00228E6C | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x00228E88 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00228EA0 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00228EBC | `GotoDefaultLayout` | Known | Navigation |
| 0x00229108 | `GotoChapterArtLayout` | Known | Navigation |
| 0x00229200 | `GotoProgressLayout` | Known | Navigation |
| 0x0022928C | `GotoProgressLayout` | Known | Navigation |
| 0x002292A0 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x0022937C | `GotoExtraInfoLoadFailedLayout` | Known | Navigation |
| 0x0022939C | `GotoExtraInfoLayout` | Known | Navigation |
| 0x002297D8 | `GotoStatusBarLayout` | Known | Navigation |
| 0x002297EC | `GotoDefaultLayout` | Known | Navigation |
| 0x002299C4 | `GotoDefault` | Known | Navigation |
| 0x00229AF8 | `GotoProgressLayout` | Known | Navigation |
| 0x00229CB8 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x00229E08 | `GotoBrightnessLayout` | Known | Navigation |
| 0x00229E8C | `GotoBrightnessLayout` | Known | Navigation |
| 0x00229F0C | `GotoVolumeLayout` | Known | Navigation |
| 0x00229F58 | `GotoScrubLayout` | Known | Navigation |
| 0x0022A020 | `GotoStatusBarLayout` | Known | Navigation |
| 0x0022A034 | `GotoDefaultLayout` | Known | Navigation |
| 0x0022A10C | `GotoScrubLayout` | Known | Navigation |
| 0x0022A15C | `GotoScrubLayout` | Known | Navigation |
| 0x00230C3C | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x00230DCC | `GotoFourCard_About` | Known | Navigation |
| 0x00230DE0 | `GotoThreeCard_About` | Known | Navigation |
| 0x00230ECC | `GotoScreen_BacklightTimer` | Known | Navigation |
| 0x00230F5C | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x00230F74 | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation |
| 0x00236218 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x00236230 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x00238824 | `GotoNowPlaying` | Known | Navigation |
| 0x00238F7C | `GotoNowPlaying` | Known | Navigation |
| 0x002395FC | `GotoFirstBoot` | Known | Navigation |
| 0x0023960C | `GotoNotesApp` | Known | Navigation |
| 0x00239620 | `GotoLockApp` | Known | Navigation |
| 0x0023F5C8 | `GotoNowPlaying` | Known | Navigation |
| 0x003DB490 | `GotoProgressLayout` | Known | Navigation |
| 0x0075FDAF | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x007D80D1 | `GotoDefault` | Known | Navigation |
| 0x007D8BC1 | `GotoDefault` | Known | Navigation |
| 0x008E09A4 | `Gotowe?` | Known | Navigation |

---

## 7. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016B9A8 | `CoverFlow_Screen` | Known | Screen layout |
| 0x00197BCC | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x00197BEC | `VoiceMemos_Menu_Screen_Recording` | Known | Screen layout |
| 0x00197C10 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x00753DBA | `Clock_Screen` | Known | Screen layout |
| 0x00753DCA | `Clock_Screen_Default"` | Known | Screen layout |
| 0x00753E2F | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x00753E8D | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x00753EA5 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x00753F12 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x00753FB0 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x0075400F | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x00754025 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x00754090 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x007540EA | `Games_Menu_Screen` | Known | Screen layout |
| 0x007540FF | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x00754169 | `Extras_Screen_Games` | Known | Screen layout |
| 0x00754228 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x007542EC | `Extras_Screen_Lock` | Known | Screen layout |
| 0x007543B5 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x00754412 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x0075442B | `Debug_MainMenu_Screen_Default"` | Known | Screen layout |
| 0x00754499 | `Extras_Screen_Debug` | Known | Screen layout |
| 0x007545D0 | `Clock_Africa_City_Screen ` | Known | Screen layout |
| 0x007545EC | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x00754670 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x0075468A | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x0075470C | `Clock_Atlantic_City_Screen"` | Known | Screen layout |
| 0x0075472A | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x007547B0 | `Clock_Australia_City_Screen#` | Known | Screen layout |
| 0x007547CF | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x00754856 | `Clock_Europe_City_Screen ` | Known | Screen layout |
| 0x00754872 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x007548F6 | `Clock_NorthAmerica_City_Screen&` | Known | Screen layout |
| 0x00754918 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x007549A2 | `Clock_Pacific_City_Screen!` | Known | Screen layout |
| 0x007549BF | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x00754A44 | `Clock_SouthAmerica_City_Screen&` | Known | Screen layout |
| 0x00754A66 | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x00754AF3 | `Clock_Screen"` | Known | Screen layout |
| 0x00754B98 | `Clock_Screen"` | Known | Screen layout |
| 0x00754C3D | `Clock_Screen"` | Known | Screen layout |
| 0x00754CE2 | `Clock_Screen"` | Known | Screen layout |
| 0x00754D87 | `Clock_Screen"` | Known | Screen layout |
| 0x00754E2C | `Clock_Screen"` | Known | Screen layout |
| 0x00754ED1 | `Clock_Screen"` | Known | Screen layout |
| 0x00754F76 | `Clock_Screen"` | Known | Screen layout |
| 0x0075501B | `Clock_Screen"` | Known | Screen layout |
| 0x007550C0 | `Clock_Screen"` | Known | Screen layout |
| 0x00755165 | `Clock_Screen"` | Known | Screen layout |
| 0x0075520A | `Clock_Screen"` | Known | Screen layout |
| 0x007552AF | `Clock_Screen"` | Known | Screen layout |
| 0x00755354 | `Clock_Screen"` | Known | Screen layout |
| 0x007553F9 | `Clock_Screen"` | Known | Screen layout |
| 0x0075549E | `Clock_Screen"` | Known | Screen layout |
| 0x00755543 | `Clock_Screen"` | Known | Screen layout |
| 0x007555E8 | `Clock_Screen"` | Known | Screen layout |
| 0x0075568D | `Clock_Screen"` | Known | Screen layout |
| 0x00755732 | `Clock_Screen"` | Known | Screen layout |
| 0x007557D7 | `Clock_Screen"` | Known | Screen layout |
| 0x0075587C | `Clock_Screen"` | Known | Screen layout |
| 0x00755921 | `Clock_Screen"` | Known | Screen layout |
| 0x007559C6 | `Clock_Screen"` | Known | Screen layout |
| 0x00755A6B | `Clock_Screen"` | Known | Screen layout |
| 0x00755B10 | `Clock_Screen"` | Known | Screen layout |
| 0x00755BB5 | `Clock_Screen"` | Known | Screen layout |
| 0x00755C5A | `Clock_Screen"` | Known | Screen layout |
| 0x00755CFF | `Clock_Screen"` | Known | Screen layout |
| 0x00755DA4 | `Clock_Screen"` | Known | Screen layout |
| 0x00755E49 | `Clock_Screen"` | Known | Screen layout |
| 0x00755EF3 | `Clock_Screen"` | Known | Screen layout |
| 0x00755F98 | `Clock_Screen"` | Known | Screen layout |
| 0x0075603D | `Clock_Screen"` | Known | Screen layout |
| 0x007560E2 | `Clock_Screen"` | Known | Screen layout |
| 0x00756187 | `Clock_Screen"` | Known | Screen layout |
| 0x0075622C | `Clock_Screen"` | Known | Screen layout |
| 0x007562D1 | `Clock_Screen"` | Known | Screen layout |
| 0x00756376 | `Clock_Screen"` | Known | Screen layout |
| 0x0075641B | `Clock_Screen"` | Known | Screen layout |
| 0x007564C0 | `Clock_Screen"` | Known | Screen layout |
| 0x00756565 | `Clock_Screen"` | Known | Screen layout |
| 0x0075660A | `Clock_Screen"` | Known | Screen layout |
| 0x007566AF | `Clock_Screen"` | Known | Screen layout |
| 0x00756754 | `Clock_Screen"` | Known | Screen layout |
| 0x007567F9 | `Clock_Screen"` | Known | Screen layout |
| 0x0075689E | `Clock_Screen"` | Known | Screen layout |
| 0x00756943 | `Clock_Screen"` | Known | Screen layout |
| 0x007569E8 | `Clock_Screen"` | Known | Screen layout |
| 0x00756A8D | `Clock_Screen"` | Known | Screen layout |
| 0x00756B32 | `Clock_Screen"` | Known | Screen layout |
| 0x00756BD7 | `Clock_Screen"` | Known | Screen layout |
| 0x00756C7C | `Clock_Screen"` | Known | Screen layout |
| 0x00756D21 | `Clock_Screen"` | Known | Screen layout |
| 0x00756DC6 | `Clock_Screen"` | Known | Screen layout |
| 0x00756E6B | `Clock_Screen"` | Known | Screen layout |
| 0x00756F10 | `Clock_Screen"` | Known | Screen layout |
| 0x00756FB5 | `Clock_Screen"` | Known | Screen layout |
| 0x0075705A | `Clock_Screen"` | Known | Screen layout |
| 0x007570FF | `Clock_Screen"` | Known | Screen layout |
| 0x007571A4 | `Clock_Screen"` | Known | Screen layout |
| 0x00757249 | `Clock_Screen"` | Known | Screen layout |
| 0x007572EE | `Clock_Screen"` | Known | Screen layout |
| 0x00757393 | `Clock_Screen"` | Known | Screen layout |
| 0x00757438 | `Clock_Screen"` | Known | Screen layout |
| 0x007574DD | `Clock_Screen"` | Known | Screen layout |
| 0x00757582 | `Clock_Screen"` | Known | Screen layout |
| 0x00757627 | `Clock_Screen"` | Known | Screen layout |
| 0x007576CC | `Clock_Screen"` | Known | Screen layout |
| 0x00757771 | `Clock_Screen"` | Known | Screen layout |
| 0x00757816 | `Clock_Screen"` | Known | Screen layout |
| 0x007578BB | `Clock_Screen"` | Known | Screen layout |
| 0x00757960 | `Clock_Screen"` | Known | Screen layout |
| 0x00757A05 | `Clock_Screen"` | Known | Screen layout |
| 0x00757AAA | `Clock_Screen"` | Known | Screen layout |
| 0x00757B4F | `Clock_Screen"` | Known | Screen layout |
| 0x00757BF4 | `Clock_Screen"` | Known | Screen layout |
| 0x00757C99 | `Clock_Screen"` | Known | Screen layout |
| 0x00757D3E | `Clock_Screen"` | Known | Screen layout |
| 0x00757DE3 | `Clock_Screen"` | Known | Screen layout |
| 0x00757E88 | `Clock_Screen"` | Known | Screen layout |
| 0x00757F2D | `Clock_Screen"` | Known | Screen layout |
| 0x00757FD2 | `Clock_Screen"` | Known | Screen layout |
| 0x00758077 | `Clock_Screen"` | Known | Screen layout |
| 0x0075811C | `Clock_Screen"` | Known | Screen layout |
| 0x007581C1 | `Clock_Screen"` | Known | Screen layout |
| 0x00758266 | `Clock_Screen"` | Known | Screen layout |
| 0x0075830B | `Clock_Screen"` | Known | Screen layout |
| 0x007583B7 | `Clock_Screen"` | Known | Screen layout |
| 0x0075845C | `Clock_Screen"` | Known | Screen layout |
| 0x00758501 | `Clock_Screen"` | Known | Screen layout |
| 0x007585AB | `Clock_Screen"` | Known | Screen layout |
| 0x00758650 | `Clock_Screen"` | Known | Screen layout |
| 0x007586F5 | `Clock_Screen"` | Known | Screen layout |
| 0x0075879A | `Clock_Screen"` | Known | Screen layout |
| 0x0075883F | `Clock_Screen"` | Known | Screen layout |
| 0x007588E4 | `Clock_Screen"` | Known | Screen layout |
| 0x00758989 | `Clock_Screen"` | Known | Screen layout |
| 0x00758A2E | `Clock_Screen"` | Known | Screen layout |
| 0x00758AD7 | `Clock_Screen"` | Known | Screen layout |
| 0x00758B7C | `Clock_Screen"` | Known | Screen layout |
| 0x00758C21 | `Clock_Screen"` | Known | Screen layout |
| 0x00758CC6 | `Clock_Screen"` | Known | Screen layout |
| 0x00758D6B | `Clock_Screen"` | Known | Screen layout |
| 0x00758E10 | `Clock_Screen"` | Known | Screen layout |
| 0x00758EB5 | `Clock_Screen"` | Known | Screen layout |
| 0x00758F5A | `Clock_Screen"` | Known | Screen layout |
| 0x00758FFF | `Clock_Screen"` | Known | Screen layout |
| 0x007590A4 | `Clock_Screen"` | Known | Screen layout |
| 0x00759149 | `Clock_Screen"` | Known | Screen layout |
| 0x007591EE | `Clock_Screen"` | Known | Screen layout |
| 0x00759293 | `Clock_Screen"` | Known | Screen layout |
| 0x00759338 | `Clock_Screen"` | Known | Screen layout |
| 0x007593DD | `Clock_Screen"` | Known | Screen layout |
| 0x00759482 | `Clock_Screen"` | Known | Screen layout |
| 0x00759527 | `Clock_Screen"` | Known | Screen layout |
| 0x007595CC | `Clock_Screen"` | Known | Screen layout |
| 0x00759671 | `Clock_Screen"` | Known | Screen layout |
| 0x00759716 | `Clock_Screen"` | Known | Screen layout |
| 0x007597BB | `Clock_Screen"` | Known | Screen layout |
| 0x00759860 | `Clock_Screen"` | Known | Screen layout |
| 0x00759905 | `Clock_Screen"` | Known | Screen layout |
| 0x007599AA | `Clock_Screen"` | Known | Screen layout |
| 0x00759A4F | `Clock_Screen"` | Known | Screen layout |
| 0x00759AF4 | `Clock_Screen"` | Known | Screen layout |
| 0x00759B99 | `Clock_Screen"` | Known | Screen layout |
| 0x00759C3E | `Clock_Screen"` | Known | Screen layout |
| 0x00759CE3 | `Clock_Screen"` | Known | Screen layout |
| 0x00759D88 | `Clock_Screen"` | Known | Screen layout |
| 0x00759E2D | `Clock_Screen"` | Known | Screen layout |
| 0x00759ED2 | `Clock_Screen"` | Known | Screen layout |
| 0x00759F77 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A01C | `Clock_Screen"` | Known | Screen layout |
| 0x0075A0C7 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A16C | `Clock_Screen"` | Known | Screen layout |
| 0x0075A211 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A2B6 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A35B | `Clock_Screen"` | Known | Screen layout |
| 0x0075A400 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A4A5 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A54A | `Clock_Screen"` | Known | Screen layout |
| 0x0075A5EF | `Clock_Screen"` | Known | Screen layout |
| 0x0075A694 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A739 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A7DE | `Clock_Screen"` | Known | Screen layout |
| 0x0075A883 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A928 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A9CD | `Clock_Screen"` | Known | Screen layout |
| 0x0075AA72 | `Clock_Screen"` | Known | Screen layout |
| 0x0075AB17 | `Clock_Screen"` | Known | Screen layout |
| 0x0075ABBC | `Clock_Screen"` | Known | Screen layout |
| 0x0075AC61 | `Clock_Screen"` | Known | Screen layout |
| 0x0075AD06 | `Clock_Screen"` | Known | Screen layout |
| 0x0075ADAB | `Clock_Screen"` | Known | Screen layout |
| 0x0075AE50 | `Clock_Screen"` | Known | Screen layout |
| 0x0075AEF5 | `Clock_Screen"` | Known | Screen layout |
| 0x0075AF9A | `Clock_Screen"` | Known | Screen layout |
| 0x0075B03F | `Clock_Screen"` | Known | Screen layout |
| 0x0075B0E4 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B189 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B22E | `Clock_Screen"` | Known | Screen layout |
| 0x0075B2D3 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B378 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B41D | `Clock_Screen"` | Known | Screen layout |
| 0x0075B4C2 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B567 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B60C | `Clock_Screen"` | Known | Screen layout |
| 0x0075B6B1 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B756 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B7FB | `Clock_Screen"` | Known | Screen layout |
| 0x0075B8A0 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B945 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B9EA | `Clock_Screen"` | Known | Screen layout |
| 0x0075BA8F | `Clock_Screen"` | Known | Screen layout |
| 0x0075BB34 | `Clock_Screen"` | Known | Screen layout |
| 0x0075BBD9 | `Clock_Screen"` | Known | Screen layout |
| 0x0075BC7E | `Clock_Screen"` | Known | Screen layout |
| 0x0075BD23 | `Clock_Screen"` | Known | Screen layout |
| 0x0075BDC8 | `Clock_Screen"` | Known | Screen layout |
| 0x0075BE6D | `Clock_Screen"` | Known | Screen layout |
| 0x0075BF12 | `Clock_Screen"` | Known | Screen layout |
| 0x0075BFB7 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C05C | `Clock_Screen"` | Known | Screen layout |
| 0x0075C107 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C1AC | `Clock_Screen"` | Known | Screen layout |
| 0x0075C251 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C2F6 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C39B | `Clock_Screen"` | Known | Screen layout |
| 0x0075C447 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C4EC | `Clock_Screen"` | Known | Screen layout |
| 0x0075C591 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C636 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C6DB | `Clock_Screen"` | Known | Screen layout |
| 0x0075C780 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C825 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C8CA | `Clock_Screen"` | Known | Screen layout |
| 0x0075C96F | `Clock_Screen"` | Known | Screen layout |
| 0x0075CA14 | `Clock_Screen"` | Known | Screen layout |
| 0x0075CAB9 | `Clock_Screen"` | Known | Screen layout |
| 0x0075CB5E | `Clock_Screen"` | Known | Screen layout |
| 0x0075CC03 | `Clock_Screen"` | Known | Screen layout |
| 0x0075CCA8 | `Clock_Screen"` | Known | Screen layout |
| 0x0075CD4D | `Clock_Screen"` | Known | Screen layout |
| 0x0075CDF2 | `Clock_Screen"` | Known | Screen layout |
| 0x0075CE97 | `Clock_Screen"` | Known | Screen layout |
| 0x0075CF3A | `Settings_DateTime_SetDate_Screen(` | Known | Screen layout |
| 0x0075CF5E | `Settings_DateTime_SetDate_Screen_Default"` | Known | Screen layout |
| 0x0075CFD7 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0075D03D | `Settings_DateTime_SetTime_Screen(` | Known | Screen layout |
| 0x0075D061 | `Settings_DateTime_SetTime_Screen_Default"` | Known | Screen layout |
| 0x0075D0DA | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x0075D145 | `Settings_DateTime_SetTimeZone_Screen,` | Known | Screen layout |
| 0x0075D16D | `Settings_DateTime_SetTimeZone_Screen_Default"` | Known | Screen layout |
| 0x0075D1EA | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x0075D2A3 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0075D353 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0075D8E2 | `Search_Main_Screen` | Known | Screen layout |
| 0x0075D8F8 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x0075DE1A | `Extras_Screen` | Known | Screen layout |
| 0x0075DE2B | `Extras_Screen_WorldClock"` | Known | Screen layout |
| 0x0075DEA8 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x0075DF0A | `Clock_Screen` | Known | Screen layout |
| 0x0075DF1A | `Clock_Screen_Default` | Known | Screen layout |
| 0x0075DFA1 | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x0075E007 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0075E01D | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x0075E088 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x0075E0EA | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0075E102 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x0075E16F | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x0075E1D3 | `AddressViewer_Main_Screen!` | Known | Screen layout |
| 0x0075E1F0 | `AddressViewer_Main_Screen_Default"` | Known | Screen layout |
| 0x0075E262 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x0075E2C9 | `Games_Menu_Screen` | Known | Screen layout |
| 0x0075E2DE | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x0075E348 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x0075E40F | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x0075E4AB | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x0075E57C | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x0075E63C | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x0075E6A0 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0075E6BF | `VoiceMemos_Menu_Screen_Default"` | Known | Screen layout |
| 0x0075E742 | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x0075E7A8 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x0075E7C0 | `Speakers_Main_Screen_Default"` | Known | Screen layout |
| 0x0075E841 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x0075E8A5 | `Radio_Screen` | Known | Screen layout |
| 0x0075E8B5 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x0075E92E | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x0075E98F | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0075EA2B | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x0075EAEE | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x0075EBAD | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x0075EC6A | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x0075F084 | `Radio_Screen` | Known | Screen layout |
| 0x0075F094 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x0075F10D | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x0075F2F1 | `Search_Main_Screen` | Known | Screen layout |
| 0x0075F307 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x0075F434 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x0075F497 | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x0075F7D8 | `Video_Settings_Screen` | Known | Screen layout |
| 0x0075F7F1 | `Video_Settings_Screen_Default"` | Known | Screen layout |
| 0x0075F8EE | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0075FBB3 | `SettingsMenus_MainMenu_Screen%` | Known | Screen layout |
| 0x0075FCC1 | `SettingsMenus_MusicMenu_Screen&` | Known | Screen layout |
| 0x0075FF6A | `SettingsMenus_Brightness_Screen&` | Known | Screen layout |
| 0x0076007F | `SettingsMenus_AudiobookSettings_Screen.` | Known | Screen layout |
| 0x007601B5 | `SettingsMenus_RadioRegions_Screen ` | Known | Screen layout |
| 0x007602CA | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x00760536 | `Settings_DateTime_Screen ` | Known | Screen layout |
| 0x00760552 | `Settings_DateTime_Screen_Default"` | Known | Screen layout |
| 0x007606DE | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x007607E3 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x007607FC | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x007608ED | `SettingsMenus_ResetAllSettings_Screen4` | Known | Screen layout |
| 0x007610BE | `Stopwatch_Screen` | Known | Screen layout |
| 0x007610D2 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x00761139 | `Stopwatch_Screen` | Known | Screen layout |
| 0x0076114D | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x007611F6 | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x00761219 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007612B2 | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x007612D5 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x00761366 | `NikePlus_ResumeWorkout_Screen%` | Known | Screen layout |
| 0x00761387 | `NikePlus_ResumeWorkout_Screen_Default"` | Known | Screen layout |
| 0x007613FD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007614A3 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00761550 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00761600 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007616B0 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00761712 | `NikePlus_Settings_Screen ` | Known | Screen layout |
| 0x0076172E | `NikePlus_Settings_Screen_Default` | Known | Screen layout |
| 0x007617B1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00761813 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x0076182E | `NikePlus_History_Screen_Default` | Known | Screen layout |
| 0x007618B0 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00761AD4 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x00761B42 | `Speakers_ToneControl_Screen#` | Known | Screen layout |
| 0x00761B61 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x00776C1D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00776CA0 | `LockediPod_Screen` | Known | Screen layout |
| 0x00776D28 | `Lock_Screen` | Known | Screen layout |
| 0x00776D37 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00776DB2 | `AddressViewer_ContactDetails_Screen+` | Known | Screen layout |
| 0x00776DD9 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x00776E54 | `Extras_Screen` | Known | Screen layout |
| 0x00776E9F | `Extras_Screen` | Known | Screen layout |
| 0x00776F86 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x00776FE4 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x00777001 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x0077706F | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00777088 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x007770FF | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0077711C | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x00777187 | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x007771A4 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0077720B | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x00777272 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x007772D0 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x007772ED | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x0077735B | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00777374 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x007773EB | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x00777408 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x00777473 | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x00777490 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x007774F7 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x00777597 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x00777620 | `Alarms_Set_Alarm_Frequency_Screen)` | Known | Screen layout |
| 0x00777645 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x007776B6 | `Alarms_Set_Alarm_Label_Screen%` | Known | Screen layout |
| 0x007776D7 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x00777744 | `Alarms_Set_Alarm_Sound_Screen%` | Known | Screen layout |
| 0x00777765 | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x007777D1 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x00777A4C | `Alarms_Set_Alarm_Playlist_Screen(` | Known | Screen layout |
| 0x00777A70 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x00777AE0 | `Alarms_Set_Alarm_Tones_Screen%` | Known | Screen layout |
| 0x00777B01 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x00777E14 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x00777E2F | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x00777F80 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x00777F97 | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x00778018 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0077802F | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x00778105 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0077811E | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x007781A3 | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x00778214 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x00778309 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00778322 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x007783A7 | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x00778418 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x007784D8 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x007784EC | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x0077861B | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x0077867E | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x007786D5 | `Clock_Screen_Default` | Known | Screen layout |
| 0x00778766 | `Clock_Region_Screen` | Known | Screen layout |
| 0x0077877D | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x007787F6 | `Clock_Screen_Default` | Known | Screen layout |
| 0x0077884D | `Clock_Screen_Default` | Known | Screen layout |
| 0x007788DE | `Clock_Region_Screen` | Known | Screen layout |
| 0x007788F5 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x00778A80 | `CoverFlow_Screen_DefaultNoAlbumInfo&` | Known | Screen layout |
| 0x00778B6E | `CoverFlow_Screen_DefaultNoAlbumInfo#` | Known | Screen layout |
| 0x00778BE3 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00778ED9 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00779089 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x007791B7 | `CoverFlow_Screen_Default&` | Known | Screen layout |
| 0x0077928D | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00779422 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00779687 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x007796E4 | `Game_Screen` | Known | Screen layout |
| 0x007796F3 | `Game_Screen_Default` | Known | Screen layout |
| 0x00779795 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x007797F7 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0077985A | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x007798BD | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00779919 | `Game_Running_Screen` | Known | Screen layout |
| 0x00779979 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x007799DB | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00779A3E | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00779AA1 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00779AFD | `Game_Running_Screen` | Known | Screen layout |
| 0x00779B5D | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00779BBF | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00779C22 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00779C85 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00779CE1 | `Game_Running_Screen` | Known | Screen layout |
| 0x00779D41 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00779DA3 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00779E06 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00779E69 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00779EC5 | `Game_Running_Screen` | Known | Screen layout |
| 0x00779F25 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00779F87 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00779FEA | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0077A04D | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0077A0A9 | `Game_Running_Screen` | Known | Screen layout |
| 0x0077A2EF | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0077A351 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0077A3B4 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0077A417 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0077A473 | `Game_Running_Screen` | Known | Screen layout |
| 0x0077A52A | `Extras_Screen` | Known | Screen layout |
| 0x0077A53B | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0077A599 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0077A736 | `Extras_Screen` | Known | Screen layout |
| 0x0077A747 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0077A7A5 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0077A942 | `Extras_Screen` | Known | Screen layout |
| 0x0077A953 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0077A9B1 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0077AB4E | `Extras_Screen` | Known | Screen layout |
| 0x0077AB5F | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0077ABBD | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0077AD5F | `Lock_Screen` | Known | Screen layout |
| 0x0077AD6E | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0077ADD0 | `Extras_Screen` | Known | Screen layout |
| 0x0077ADE1 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0077AE40 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077AEBA | `MainMenus_Main_Screen)` | Known | Screen layout |
| 0x0077B08B | `Lock_Screen` | Known | Screen layout |
| 0x0077B09A | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0077B0FC | `Extras_Screen` | Known | Screen layout |
| 0x0077B10D | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0077B16C | `LockediPod_Screen` | Known | Screen layout |
| 0x0077B1E6 | `MainMenus_Main_Screen%` | Known | Screen layout |
| 0x0077B24D | `LockediPod_Screen` | Known | Screen layout |
| 0x0077B262 | `LockediPod_Screen_Default,` | Known | Screen layout |
| 0x0077B3B1 | `Lock_Screen` | Known | Screen layout |
| 0x0077B3C0 | `Lock_Screen_Default_Layout"` | Known | Screen layout |
| 0x0077B429 | `Lock_Screen` | Known | Screen layout |
| 0x0077B438 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0077B49A | `Extras_Screen` | Known | Screen layout |
| 0x0077B4AB | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0077B50A | `LockediPod_Screen` | Known | Screen layout |
| 0x0077B584 | `MainMenus_Main_Screen,` | Known | Screen layout |
| 0x0077B6E0 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0077B746 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0077B7AA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077B839 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0077B8A6 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0077B913 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0077B980 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077B9E8 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0077BA4E | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0077BAB2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077BB41 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0077BBAE | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0077BC1B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0077BC88 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077BCF0 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0077BD56 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0077BDBA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077BE49 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0077BEB6 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0077BF23 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0077BF90 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077BFF8 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0077C05E | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0077C0C2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077C151 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0077C1BE | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0077C22B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0077C298 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077C300 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0077C366 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0077C3CA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077C459 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0077C4C6 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0077C533 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0077C5A0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077C5F9 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x0077C662 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x0077C6C9 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0077C764 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x0077C7CD | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x0077C836 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x0077C89D | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0077C938 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x0077C9A1 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x0077CA0A | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x0077CA71 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0077CB0C | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x0077CBF8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077CC14 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077CC82 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077CC9F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077CD0A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077CD2A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077CDA1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077CDBD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077CE2D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077CE4C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077CEB8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077CECC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077CF45 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077CFB9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077D029 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077D091 | `NoContent_Screen` | Known | Screen layout |
| 0x0077D0A5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077D109 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077D170 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077D18A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077D1F8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077D26A | `NoContent_Screen` | Known | Screen layout |
| 0x0077D27E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077D2E8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077D351 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077D365 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077D3CB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077D439 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077D4A6 | `NoContent_Screen` | Known | Screen layout |
| 0x0077D4BA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077D522 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077D58C | `NoContent_Screen` | Known | Screen layout |
| 0x0077D5A0 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077D607 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077D671 | `NoContent_Screen` | Known | Screen layout |
| 0x0077D685 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077D6F2 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077D764 | `NoContent_Screen` | Known | Screen layout |
| 0x0077D778 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077D7E0 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077D849 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077D864 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077D8CA | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077D8E6 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077D9C5 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077D9DE | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077DA3F | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077DA53 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077DBC1 | `Radio_Screen` | Known | Screen layout |
| 0x0077DBD1 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077DC32 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077DCB5 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077DD3D | `Lock_Screen` | Known | Screen layout |
| 0x0077DD4C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077DDAF | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077DE11 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077DE2D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077DE9F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077DEBE | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077DF26 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077DF40 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077DFA8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077DFC5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077E031 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077E09B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077E0B5 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077E125 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077E198 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077E209 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077E278 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077E2E4 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077E2FF | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077E374 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077E3DB | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077E43D | `Photos_Screen` | Known | Screen layout |
| 0x0077E4A1 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077E4BF | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077E531 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077E54E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077E5B4 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077E5CF | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077E638 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077E655 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077E6CC | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077E6F0 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077E75E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077E779 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077E834 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077E850 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077E8BE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077E8DB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077E946 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077E966 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077E9DD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077E9F9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077EA69 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077EA88 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077EAF4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077EB08 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077EB81 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077EBF5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077EC65 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077ECCD | `NoContent_Screen` | Known | Screen layout |
| 0x0077ECE1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077ED45 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077EDAC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077EDC6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077EE34 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077EEA6 | `NoContent_Screen` | Known | Screen layout |
| 0x0077EEBA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077EF24 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077EF8D | `No_Photos_Screen` | Known | Screen layout |
| 0x0077EFA1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077F007 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077F075 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077F0E2 | `NoContent_Screen` | Known | Screen layout |
| 0x0077F0F6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077F15E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077F1C8 | `NoContent_Screen` | Known | Screen layout |
| 0x0077F1DC | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077F243 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077F2AD | `NoContent_Screen` | Known | Screen layout |
| 0x0077F2C1 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077F32E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077F3A0 | `NoContent_Screen` | Known | Screen layout |
| 0x0077F3B4 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077F41C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077F485 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077F4A0 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077F506 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077F522 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077F601 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077F61A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077F67B | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077F68F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077F7FD | `Radio_Screen` | Known | Screen layout |
| 0x0077F80D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077F86E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077F8F1 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077F979 | `Lock_Screen` | Known | Screen layout |
| 0x0077F988 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077F9EB | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077FA4D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077FA69 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077FADB | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077FAFA | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077FB62 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077FB7C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077FBE4 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077FC01 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077FC6D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077FCD7 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077FCF1 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077FD61 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077FDD4 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077FE45 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077FEB4 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077FF20 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077FF3B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077FFB0 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00780017 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00780079 | `Photos_Screen` | Known | Screen layout |
| 0x007800DD | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007800FB | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078016D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078018A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007801F0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078020B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00780274 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00780291 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00780308 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078032C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078039A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007803B5 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00780470 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078048C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007804FA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00780517 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00780582 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007805A2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00780619 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00780635 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007806A5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007806C4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00780730 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00780744 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007807BD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00780831 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007808A1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00780909 | `NoContent_Screen` | Known | Screen layout |
| 0x0078091D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00780981 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007809E8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00780A02 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00780A70 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00780AE2 | `NoContent_Screen` | Known | Screen layout |
| 0x00780AF6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00780B60 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00780BC9 | `No_Photos_Screen` | Known | Screen layout |
| 0x00780BDD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00780C43 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00780CB1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00780D1E | `NoContent_Screen` | Known | Screen layout |
| 0x00780D32 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00780D9A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00780E04 | `NoContent_Screen` | Known | Screen layout |
| 0x00780E18 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00780E7F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00780EE9 | `NoContent_Screen` | Known | Screen layout |
| 0x00780EFD | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00780F6A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00780FDC | `NoContent_Screen` | Known | Screen layout |
| 0x00780FF0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00781058 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007810C1 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007810DC | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00781142 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078115E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078123D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00781256 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007812B7 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007812CB | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00781439 | `Radio_Screen` | Known | Screen layout |
| 0x00781449 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007814AA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078152D | `LockediPod_Screen` | Known | Screen layout |
| 0x007815B5 | `Lock_Screen` | Known | Screen layout |
| 0x007815C4 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00781627 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00781689 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007816A5 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00781717 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00781736 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078179E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007817B8 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00781820 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078183D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007818A9 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00781913 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078192D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078199D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00781A10 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00781A81 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00781AF0 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00781B5C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00781B77 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00781BEC | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00781C53 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00781CB5 | `Photos_Screen` | Known | Screen layout |
| 0x00781D19 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00781D37 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00781DA9 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00781DC6 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00781E2C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00781E47 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00781EB0 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00781ECD | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00781F44 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00781F68 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00781FD6 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00781FF1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007820AC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007820C8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00782136 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00782153 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007821BE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007821DE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00782255 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00782271 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007822E1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00782300 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078236C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00782380 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007823F9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078246D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007824DD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00782545 | `NoContent_Screen` | Known | Screen layout |
| 0x00782559 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007825BD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00782624 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078263E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007826AC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078271E | `NoContent_Screen` | Known | Screen layout |
| 0x00782732 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078279C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00782805 | `No_Photos_Screen` | Known | Screen layout |
| 0x00782819 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078287F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007828ED | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078295A | `NoContent_Screen` | Known | Screen layout |
| 0x0078296E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007829D6 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00782A40 | `NoContent_Screen` | Known | Screen layout |
| 0x00782A54 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00782ABB | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00782B25 | `NoContent_Screen` | Known | Screen layout |
| 0x00782B39 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00782BA6 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00782C18 | `NoContent_Screen` | Known | Screen layout |
| 0x00782C2C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00782C94 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00782CFD | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00782D18 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00782D7E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00782D9A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00782E79 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00782E92 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00782EF3 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00782F07 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00783075 | `Radio_Screen` | Known | Screen layout |
| 0x00783085 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007830E6 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00783169 | `LockediPod_Screen` | Known | Screen layout |
| 0x007831F1 | `Lock_Screen` | Known | Screen layout |
| 0x00783200 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00783263 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007832C5 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007832E1 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00783353 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00783372 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007833DA | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007833F4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078345C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00783479 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007834E5 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078354F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00783569 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007835D9 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078364C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007836BD | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078372C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00783798 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007837B3 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00783828 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078388F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007838F1 | `Photos_Screen` | Known | Screen layout |
| 0x00783955 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00783973 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007839E5 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00783A02 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00783A68 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00783A83 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00783AEC | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00783B09 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00783B80 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00783BA4 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00783C12 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00783C2D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00783CE8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00783D04 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00783D72 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00783D8F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00783DFA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00783E1A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00783E91 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00783EAD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00783F1D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00783F3C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00783FA8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00783FBC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00784035 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007840A9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00784119 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00784181 | `NoContent_Screen` | Known | Screen layout |
| 0x00784195 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007841F9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00784260 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078427A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007842E8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078435A | `NoContent_Screen` | Known | Screen layout |
| 0x0078436E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007843D8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00784441 | `No_Photos_Screen` | Known | Screen layout |
| 0x00784455 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007844BB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00784529 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00784596 | `NoContent_Screen` | Known | Screen layout |
| 0x007845AA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00784612 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078467C | `NoContent_Screen` | Known | Screen layout |
| 0x00784690 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007846F7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00784761 | `NoContent_Screen` | Known | Screen layout |
| 0x00784775 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007847E2 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00784854 | `NoContent_Screen` | Known | Screen layout |
| 0x00784868 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007848D0 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00784939 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00784954 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007849BA | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007849D6 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00784AB5 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00784ACE | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00784B2F | `FirstBoot_Screen` | Known | Screen layout |
| 0x00784B43 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00784CB1 | `Radio_Screen` | Known | Screen layout |
| 0x00784CC1 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00784D22 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00784DA5 | `LockediPod_Screen` | Known | Screen layout |
| 0x00784E2D | `Lock_Screen` | Known | Screen layout |
| 0x00784E3C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00784E9F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00784F01 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00784F1D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00784F8F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00784FAE | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00785016 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00785030 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00785098 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007850B5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00785121 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078518B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007851A5 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00785215 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00785288 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007852F9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00785368 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007853D4 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007853EF | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00785464 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007854CB | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078552D | `Photos_Screen` | Known | Screen layout |
| 0x00785591 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007855AF | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00785621 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078563E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007856A4 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007856BF | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00785728 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00785745 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007857BC | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007857E0 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078584E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00785869 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00785924 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00785940 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007859AE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007859CB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00785A36 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00785A56 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00785ACD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00785AE9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00785B59 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00785B78 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00785BE4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00785BF8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00785C71 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00785CE5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00785D55 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00785DBD | `NoContent_Screen` | Known | Screen layout |
| 0x00785DD1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00785E35 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00785E9C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00785EB6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00785F24 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00785F96 | `NoContent_Screen` | Known | Screen layout |
| 0x00785FAA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00786014 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078607D | `No_Photos_Screen` | Known | Screen layout |
| 0x00786091 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007860F7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00786165 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007861D2 | `NoContent_Screen` | Known | Screen layout |
| 0x007861E6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078624E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007862B8 | `NoContent_Screen` | Known | Screen layout |
| 0x007862CC | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00786333 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078639D | `NoContent_Screen` | Known | Screen layout |
| 0x007863B1 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078641E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00786490 | `NoContent_Screen` | Known | Screen layout |
| 0x007864A4 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078650C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00786575 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00786590 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007865F6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00786612 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007866F1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078670A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078676B | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078677F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007868ED | `Radio_Screen` | Known | Screen layout |
| 0x007868FD | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078695E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007869E1 | `LockediPod_Screen` | Known | Screen layout |
| 0x00786A69 | `Lock_Screen` | Known | Screen layout |
| 0x00786A78 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00786ADB | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00786B3D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00786B59 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00786BCB | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00786BEA | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00786C52 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00786C6C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00786CD4 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00786CF1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00786D5D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00786DC7 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00786DE1 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00786E51 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00786EC4 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00786F35 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00786FA4 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00787010 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078702B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007870A0 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00787107 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00787169 | `Photos_Screen` | Known | Screen layout |
| 0x007871CD | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007871EB | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078725D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078727A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007872E0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007872FB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00787364 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00787381 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007873F8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078741C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078748A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007874A5 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00787560 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078757C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007875EA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00787607 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00787672 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00787692 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00787709 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00787725 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00787795 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007877B4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00787820 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00787834 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007878AD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00787921 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00787991 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007879F9 | `NoContent_Screen` | Known | Screen layout |
| 0x00787A0D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00787A71 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00787AD8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00787AF2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00787B60 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00787BD2 | `NoContent_Screen` | Known | Screen layout |
| 0x00787BE6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00787C50 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00787CB9 | `No_Photos_Screen` | Known | Screen layout |
| 0x00787CCD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00787D33 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00787DA1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00787E0E | `NoContent_Screen` | Known | Screen layout |
| 0x00787E22 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00787E8A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00787EF4 | `NoContent_Screen` | Known | Screen layout |
| 0x00787F08 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00787F6F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00787FD9 | `NoContent_Screen` | Known | Screen layout |
| 0x00787FED | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078805A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007880CC | `NoContent_Screen` | Known | Screen layout |
| 0x007880E0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00788148 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007881B1 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007881CC | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00788232 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078824E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078832D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00788346 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007883A7 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007883BB | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00788529 | `Radio_Screen` | Known | Screen layout |
| 0x00788539 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078859A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078861D | `LockediPod_Screen` | Known | Screen layout |
| 0x007886A5 | `Lock_Screen` | Known | Screen layout |
| 0x007886B4 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00788717 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00788779 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00788795 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00788807 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00788826 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078888E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007888A8 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00788910 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078892D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00788999 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00788A03 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00788A1D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00788A8D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00788B00 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00788B71 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00788BE0 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00788C4C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00788C67 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00788CDC | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00788D43 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00788DA5 | `Photos_Screen` | Known | Screen layout |
| 0x00788E09 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00788E27 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00788E99 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00788EB6 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00788F1C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00788F37 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00788FA0 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00788FBD | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00789034 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00789058 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007890C6 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007890E1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078919C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007891B8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00789226 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00789243 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007892AE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007892CE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00789345 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00789361 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007893D1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007893F0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078945C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00789470 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007894E9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078955D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007895CD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00789635 | `NoContent_Screen` | Known | Screen layout |
| 0x00789649 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007896AD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00789714 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078972E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078979C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078980E | `NoContent_Screen` | Known | Screen layout |
| 0x00789822 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078988C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007898F5 | `No_Photos_Screen` | Known | Screen layout |
| 0x00789909 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078996F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007899DD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00789A4A | `NoContent_Screen` | Known | Screen layout |
| 0x00789A5E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00789AC6 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00789B30 | `NoContent_Screen` | Known | Screen layout |
| 0x00789B44 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00789BAB | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00789C15 | `NoContent_Screen` | Known | Screen layout |
| 0x00789C29 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00789C96 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00789D08 | `NoContent_Screen` | Known | Screen layout |
| 0x00789D1C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00789D84 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00789DED | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00789E08 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00789E6E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00789E8A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00789F69 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00789F82 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00789FE3 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00789FF7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078A165 | `Radio_Screen` | Known | Screen layout |
| 0x0078A175 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078A1D6 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078A259 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078A2E1 | `Lock_Screen` | Known | Screen layout |
| 0x0078A2F0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078A353 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078A3B5 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078A3D1 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078A443 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078A462 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078A4CA | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078A4E4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078A54C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078A569 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078A5D5 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078A63F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078A659 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078A6C9 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078A73C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078A7AD | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078A81C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078A888 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078A8A3 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078A918 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078A97F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078A9E1 | `Photos_Screen` | Known | Screen layout |
| 0x0078AA45 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078AA63 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078AAD5 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078AAF2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078AB58 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078AB73 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078ABDC | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078ABF9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078AC70 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078AC94 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078AD02 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078AD1D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078ADD8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078ADF4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078AE62 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078AE7F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078AEEA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078AF0A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078AF81 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078AF9D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078B00D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078B02C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078B098 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078B0AC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078B125 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078B199 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078B209 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078B271 | `NoContent_Screen` | Known | Screen layout |
| 0x0078B285 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078B2E9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078B350 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078B36A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078B3D8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078B44A | `NoContent_Screen` | Known | Screen layout |
| 0x0078B45E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078B4C8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078B531 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078B545 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078B5AB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078B619 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078B686 | `NoContent_Screen` | Known | Screen layout |
| 0x0078B69A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078B702 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078B76C | `NoContent_Screen` | Known | Screen layout |
| 0x0078B780 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078B7E7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078B851 | `NoContent_Screen` | Known | Screen layout |
| 0x0078B865 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078B8D2 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078B944 | `NoContent_Screen` | Known | Screen layout |
| 0x0078B958 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078B9C0 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078BA29 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078BA44 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078BAAA | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078BAC6 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078BBA5 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078BBBE | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078BC1F | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078BC33 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078BDA1 | `Radio_Screen` | Known | Screen layout |
| 0x0078BDB1 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078BE12 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078BE95 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078BF1D | `Lock_Screen` | Known | Screen layout |
| 0x0078BF2C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078BF8F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078BFF1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078C00D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078C07F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078C09E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078C106 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078C120 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078C188 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078C1A5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078C211 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078C27B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078C295 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078C305 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078C378 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078C3E9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078C458 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078C4C4 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078C4DF | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078C554 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078C5BB | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078C61D | `Photos_Screen` | Known | Screen layout |
| 0x0078C681 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078C69F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078C711 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078C72E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078C794 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078C7AF | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078C818 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078C835 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078C8AC | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078C8D0 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078C93E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078C959 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078CA14 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078CA30 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078CA9E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078CABB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078CB26 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078CB46 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078CBBD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078CBD9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078CC49 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078CC68 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078CCD4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078CCE8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078CD61 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078CDD5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078CE45 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078CEAD | `NoContent_Screen` | Known | Screen layout |
| 0x0078CEC1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078CF25 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078CF8C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078CFA6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078D014 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078D086 | `NoContent_Screen` | Known | Screen layout |
| 0x0078D09A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078D104 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078D16D | `No_Photos_Screen` | Known | Screen layout |
| 0x0078D181 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078D1E7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078D255 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078D2C2 | `NoContent_Screen` | Known | Screen layout |
| 0x0078D2D6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078D33E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078D3A8 | `NoContent_Screen` | Known | Screen layout |
| 0x0078D3BC | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078D423 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078D48D | `NoContent_Screen` | Known | Screen layout |
| 0x0078D4A1 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078D50E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078D580 | `NoContent_Screen` | Known | Screen layout |
| 0x0078D594 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078D5FC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078D665 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078D680 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078D6E6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078D702 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078D7E1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078D7FA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078D85B | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078D86F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078D9DD | `Radio_Screen` | Known | Screen layout |
| 0x0078D9ED | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078DA4E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078DAD1 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078DB59 | `Lock_Screen` | Known | Screen layout |
| 0x0078DB68 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078DBCB | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078DC2D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078DC49 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078DCBB | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078DCDA | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078DD42 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078DD5C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078DDC4 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078DDE1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078DE4D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078DEB7 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078DED1 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078DF41 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078DFB4 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078E025 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078E094 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078E100 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078E11B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078E190 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078E1F7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078E259 | `Photos_Screen` | Known | Screen layout |
| 0x0078E2BD | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078E2DB | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078E34D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078E36A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078E3D0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078E3EB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078E454 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078E471 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078E4E8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078E50C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078E57A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078E595 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078E650 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078E66C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078E6DA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078E6F7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078E762 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078E782 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078E7F9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078E815 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078E885 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078E8A4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078E910 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078E924 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078E99D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078EA11 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078EA81 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078EAE9 | `NoContent_Screen` | Known | Screen layout |
| 0x0078EAFD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078EB61 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078EBC8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078EBE2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078EC50 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078ECC2 | `NoContent_Screen` | Known | Screen layout |
| 0x0078ECD6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078ED40 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078EDA9 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078EDBD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078EE23 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078EE91 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078EEFE | `NoContent_Screen` | Known | Screen layout |
| 0x0078EF12 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078EF7A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078EFE4 | `NoContent_Screen` | Known | Screen layout |
| 0x0078EFF8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078F05F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078F0C9 | `NoContent_Screen` | Known | Screen layout |
| 0x0078F0DD | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078F14A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078F1BC | `NoContent_Screen` | Known | Screen layout |
| 0x0078F1D0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078F238 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078F2A1 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078F2BC | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078F322 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078F33E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078F41D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078F436 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078F497 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078F4AB | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078F619 | `Radio_Screen` | Known | Screen layout |
| 0x0078F629 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078F68A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078F70D | `LockediPod_Screen` | Known | Screen layout |
| 0x0078F795 | `Lock_Screen` | Known | Screen layout |
| 0x0078F7A4 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078F807 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078F869 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078F885 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078F8F7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078F916 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078F97E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078F998 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078FA00 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078FA1D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078FA89 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078FAF3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078FB0D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078FB7D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078FBF0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078FC61 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078FCD0 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078FD3C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078FD57 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078FDCC | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078FE33 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078FE95 | `Photos_Screen` | Known | Screen layout |
| 0x0078FEF9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078FF17 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078FF89 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078FFA6 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079000C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00790027 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00790090 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007900AD | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00790124 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00790148 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007901B6 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007901D1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079028C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007902A8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00790316 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00790333 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079039E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007903BE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00790435 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00790451 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007904C1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007904E0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079054C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00790560 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007905D9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079064D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007906BD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00790725 | `NoContent_Screen` | Known | Screen layout |
| 0x00790739 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079079D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00790804 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079081E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079088C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007908FE | `NoContent_Screen` | Known | Screen layout |
| 0x00790912 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079097C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007909E5 | `No_Photos_Screen` | Known | Screen layout |
| 0x007909F9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00790A5F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00790ACD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00790B3A | `NoContent_Screen` | Known | Screen layout |
| 0x00790B4E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00790BB6 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00790C20 | `NoContent_Screen` | Known | Screen layout |
| 0x00790C34 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00790C9B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00790D05 | `NoContent_Screen` | Known | Screen layout |
| 0x00790D19 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00790D86 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00790DF8 | `NoContent_Screen` | Known | Screen layout |
| 0x00790E0C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00790E74 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00790EDD | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00790EF8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00790F5E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00790F7A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00791059 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00791072 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007910D3 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007910E7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00791255 | `Radio_Screen` | Known | Screen layout |
| 0x00791265 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007912C6 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00791349 | `LockediPod_Screen` | Known | Screen layout |
| 0x007913D1 | `Lock_Screen` | Known | Screen layout |
| 0x007913E0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00791443 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007914A5 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007914C1 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00791533 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00791552 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007915BA | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007915D4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079163C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00791659 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007916C5 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079172F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00791749 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007917B9 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079182C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079189D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079190C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00791978 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00791993 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00791A08 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00791A6F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00791AD1 | `Photos_Screen` | Known | Screen layout |
| 0x00791B35 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00791B53 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00791BC5 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00791BE2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00791C48 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00791C63 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00791CCC | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00791CE9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00791D60 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00791D84 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00791DF2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00791E0D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00791EC8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00791EE4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00791F52 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00791F6F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00791FDA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00791FFA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00792071 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079208D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007920FD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079211C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00792188 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079219C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00792215 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00792289 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007922F9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00792361 | `NoContent_Screen` | Known | Screen layout |
| 0x00792375 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007923D9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00792440 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079245A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007924C8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079253A | `NoContent_Screen` | Known | Screen layout |
| 0x0079254E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007925B8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00792621 | `No_Photos_Screen` | Known | Screen layout |
| 0x00792635 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079269B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00792709 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00792776 | `NoContent_Screen` | Known | Screen layout |
| 0x0079278A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007927F2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0079285C | `NoContent_Screen` | Known | Screen layout |
| 0x00792870 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007928D7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00792941 | `NoContent_Screen` | Known | Screen layout |
| 0x00792955 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007929C2 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00792A34 | `NoContent_Screen` | Known | Screen layout |
| 0x00792A48 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00792AB0 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00792B19 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00792B34 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00792B9A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00792BB6 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00792C95 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00792CAE | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00792D0F | `FirstBoot_Screen` | Known | Screen layout |
| 0x00792D23 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00792E91 | `Radio_Screen` | Known | Screen layout |
| 0x00792EA1 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00792F02 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00792F85 | `LockediPod_Screen` | Known | Screen layout |
| 0x0079300D | `Lock_Screen` | Known | Screen layout |
| 0x0079301C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079307F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007930E1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007930FD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0079316F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079318E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007931F6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00793210 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00793278 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00793295 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00793301 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079336B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00793385 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007933F5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00793468 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007934D9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00793548 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007935B4 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007935CF | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00793644 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007936AB | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079370D | `Photos_Screen` | Known | Screen layout |
| 0x00793771 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079378F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00793801 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079381E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00793884 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079389F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00793908 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00793925 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079399C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007939C0 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00793A2E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00793A49 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00793B04 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00793B20 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00793B8E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00793BAB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00793C16 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00793C36 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00793CAD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00793CC9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00793D39 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00793D58 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00793DC4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00793DD8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00793E51 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00793EC5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00793F35 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00793F9D | `NoContent_Screen` | Known | Screen layout |
| 0x00793FB1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00794015 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0079407C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00794096 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00794104 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00794176 | `NoContent_Screen` | Known | Screen layout |
| 0x0079418A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007941F4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0079425D | `No_Photos_Screen` | Known | Screen layout |
| 0x00794271 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007942D7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00794345 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007943B2 | `NoContent_Screen` | Known | Screen layout |
| 0x007943C6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079442E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00794498 | `NoContent_Screen` | Known | Screen layout |
| 0x007944AC | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00794513 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079457D | `NoContent_Screen` | Known | Screen layout |
| 0x00794591 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007945FE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00794670 | `NoContent_Screen` | Known | Screen layout |
| 0x00794684 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007946EC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00794755 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00794770 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007947D6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007947F2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007948D1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007948EA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0079494B | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079495F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00794ACD | `Radio_Screen` | Known | Screen layout |
| 0x00794ADD | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00794B3E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00794BC1 | `LockediPod_Screen` | Known | Screen layout |
| 0x00794C49 | `Lock_Screen` | Known | Screen layout |
| 0x00794C58 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00794CBB | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00794D1D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00794D39 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00794DAB | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00794DCA | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00794E32 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00794E4C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00794EB4 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00794ED1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00794F3D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00794FA7 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00794FC1 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00795031 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007950A4 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00795115 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00795184 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007951F0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079520B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00795280 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007952E7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00795349 | `Photos_Screen` | Known | Screen layout |
| 0x007953AD | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007953CB | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0079543D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079545A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007954C0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007954DB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00795544 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00795561 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007955D8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007955FC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079566A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00795685 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00795740 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079575C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007957CA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007957E7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00795852 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00795872 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007958E9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00795905 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00795975 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00795994 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00795A00 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00795A14 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00795A8D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00795B01 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00795B71 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00795BD9 | `NoContent_Screen` | Known | Screen layout |
| 0x00795BED | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00795C51 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00795CB8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00795CD2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00795D40 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00795DB2 | `NoContent_Screen` | Known | Screen layout |
| 0x00795DC6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00795E30 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00795E99 | `No_Photos_Screen` | Known | Screen layout |
| 0x00795EAD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00795F13 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00795F81 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00795FEE | `NoContent_Screen` | Known | Screen layout |
| 0x00796002 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079606A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007960D4 | `NoContent_Screen` | Known | Screen layout |
| 0x007960E8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079614F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007961B9 | `NoContent_Screen` | Known | Screen layout |
| 0x007961CD | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079623A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007962AC | `NoContent_Screen` | Known | Screen layout |
| 0x007962C0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00796328 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00796391 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007963AC | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00796412 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079642E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079650D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00796526 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00796587 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079659B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00796709 | `Radio_Screen` | Known | Screen layout |
| 0x00796719 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079677A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007967FD | `LockediPod_Screen` | Known | Screen layout |
| 0x00796885 | `Lock_Screen` | Known | Screen layout |
| 0x00796894 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007968F7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00796959 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00796975 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007969E7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00796A06 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00796A6E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00796A88 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00796AF0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00796B0D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00796B79 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00796BE3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00796BFD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00796C6D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00796CE0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00796D51 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00796DC0 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00796E2C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00796E47 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00796EBC | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00796F23 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00796F85 | `Photos_Screen` | Known | Screen layout |
| 0x00796FE9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00797007 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00797079 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00797096 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007970FC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00797117 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00797180 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079719D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00797214 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00797238 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007972A6 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007972C1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079737C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00797398 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00797406 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00797423 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079748E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007974AE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00797525 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00797541 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007975B1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007975D0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079763C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00797650 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007976C9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079773D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007977AD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00797815 | `NoContent_Screen` | Known | Screen layout |
| 0x00797829 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079788D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007978F4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079790E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079797C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007979EE | `NoContent_Screen` | Known | Screen layout |
| 0x00797A02 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00797A6C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00797AD5 | `No_Photos_Screen` | Known | Screen layout |
| 0x00797AE9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00797B4F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00797BBD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00797C2A | `NoContent_Screen` | Known | Screen layout |
| 0x00797C3E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00797CA6 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00797D10 | `NoContent_Screen` | Known | Screen layout |
| 0x00797D24 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00797D8B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00797DF5 | `NoContent_Screen` | Known | Screen layout |
| 0x00797E09 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00797E76 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00797EE8 | `NoContent_Screen` | Known | Screen layout |
| 0x00797EFC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00797F64 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00797FCD | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00797FE8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079804E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079806A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00798149 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00798162 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007981C3 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007981D7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00798345 | `Radio_Screen` | Known | Screen layout |
| 0x00798355 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007983B6 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00798439 | `LockediPod_Screen` | Known | Screen layout |
| 0x007984C1 | `Lock_Screen` | Known | Screen layout |
| 0x007984D0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00798533 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00798595 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007985B1 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00798623 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00798642 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007986AA | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007986C4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079872C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00798749 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007987B5 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079881F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00798839 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007988A9 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079891C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079898D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007989FC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00798A68 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00798A83 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00798AF8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00798B5F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00798BC1 | `Photos_Screen` | Known | Screen layout |
| 0x00798C25 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00798C43 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00798CB5 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00798CD2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00798D38 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00798D53 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00798DBC | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00798DD9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00798E50 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00798E74 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00798EE2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00798EFD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00798FB8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00798FD4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00799042 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079905F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007990CA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007990EA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00799161 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079917D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007991ED | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079920C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00799278 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079928C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00799305 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00799379 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007993E9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00799451 | `NoContent_Screen` | Known | Screen layout |
| 0x00799465 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007994C9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00799530 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079954A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007995B8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079962A | `NoContent_Screen` | Known | Screen layout |
| 0x0079963E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007996A8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00799711 | `No_Photos_Screen` | Known | Screen layout |
| 0x00799725 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079978B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007997F9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00799866 | `NoContent_Screen` | Known | Screen layout |
| 0x0079987A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007998E2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0079994C | `NoContent_Screen` | Known | Screen layout |
| 0x00799960 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007999C7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00799A31 | `NoContent_Screen` | Known | Screen layout |
| 0x00799A45 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00799AB2 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00799B24 | `NoContent_Screen` | Known | Screen layout |
| 0x00799B38 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00799BA0 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00799C09 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00799C24 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00799C8A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00799CA6 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00799D85 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00799D9E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00799DFF | `FirstBoot_Screen` | Known | Screen layout |
| 0x00799E13 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00799F81 | `Radio_Screen` | Known | Screen layout |
| 0x00799F91 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00799FF2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079A075 | `LockediPod_Screen` | Known | Screen layout |
| 0x0079A0FD | `Lock_Screen` | Known | Screen layout |
| 0x0079A10C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079A16F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0079A1D1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079A1ED | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0079A25F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079A27E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079A2E6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079A300 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079A368 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079A385 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079A3F1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079A45B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0079A475 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0079A4E5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079A558 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079A5C9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079A638 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079A6A4 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079A6BF | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0079A734 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079A79B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079A7FD | `Photos_Screen` | Known | Screen layout |
| 0x0079A861 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079A87F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0079A8F1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079A90E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079A974 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079A98F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079A9F8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079AA15 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079AA8C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079AAB0 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079AB1E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0079AB39 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079ABF4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079AC10 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079AC7E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079AC9B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079AD06 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079AD26 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079AD9D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079ADB9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079AE29 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079AE48 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079AEB4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079AEC8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0079AF41 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079AFB5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079B025 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079B08D | `NoContent_Screen` | Known | Screen layout |
| 0x0079B0A1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079B105 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0079B16C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079B186 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079B1F4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079B266 | `NoContent_Screen` | Known | Screen layout |
| 0x0079B27A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079B2E4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0079B34D | `No_Photos_Screen` | Known | Screen layout |
| 0x0079B361 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079B3C7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079B435 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0079B4A2 | `NoContent_Screen` | Known | Screen layout |
| 0x0079B4B6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079B51E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0079B588 | `NoContent_Screen` | Known | Screen layout |
| 0x0079B59C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079B603 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079B66D | `NoContent_Screen` | Known | Screen layout |
| 0x0079B681 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079B6EE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079B760 | `NoContent_Screen` | Known | Screen layout |
| 0x0079B774 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079B7DC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0079B845 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0079B860 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079B8C6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079B8E2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079B9C1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079B9DA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0079BA3B | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079BA4F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0079BBBD | `Radio_Screen` | Known | Screen layout |
| 0x0079BBCD | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079BC2E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079BCB1 | `LockediPod_Screen` | Known | Screen layout |
| 0x0079BD39 | `Lock_Screen` | Known | Screen layout |
| 0x0079BD48 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079BDAB | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0079BE0D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079BE29 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0079BE9B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079BEBA | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079BF22 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079BF3C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079BFA4 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079BFC1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079C02D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079C097 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0079C0B1 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0079C121 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079C194 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079C205 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079C274 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079C2E0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079C2FB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0079C370 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079C3D7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079C439 | `Photos_Screen` | Known | Screen layout |
| 0x0079C49D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079C4BB | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0079C52D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079C54A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079C5B0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079C5CB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079C634 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079C651 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079C6C8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079C6EC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079C75A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0079C775 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079C830 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079C84C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079C8BA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079C8D7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079C942 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079C962 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079C9D9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079C9F5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079CA65 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079CA84 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079CAF0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079CB04 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0079CB7D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079CBF1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079CC61 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079CCC9 | `NoContent_Screen` | Known | Screen layout |
| 0x0079CCDD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079CD41 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0079CDA8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079CDC2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079CE30 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079CEA2 | `NoContent_Screen` | Known | Screen layout |
| 0x0079CEB6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079CF20 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0079CF89 | `No_Photos_Screen` | Known | Screen layout |
| 0x0079CF9D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079D003 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079D071 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0079D0DE | `NoContent_Screen` | Known | Screen layout |
| 0x0079D0F2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079D15A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0079D1C4 | `NoContent_Screen` | Known | Screen layout |
| 0x0079D1D8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079D23F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079D2A9 | `NoContent_Screen` | Known | Screen layout |
| 0x0079D2BD | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079D32A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079D39C | `NoContent_Screen` | Known | Screen layout |
| 0x0079D3B0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079D418 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0079D481 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0079D49C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079D502 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079D51E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079D5FD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079D616 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0079D677 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079D68B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0079D7F9 | `Radio_Screen` | Known | Screen layout |
| 0x0079D809 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079D86A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079D8ED | `LockediPod_Screen` | Known | Screen layout |
| 0x0079D975 | `Lock_Screen` | Known | Screen layout |
| 0x0079D984 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079D9E7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0079DA49 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079DA65 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0079DAD7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079DAF6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079DB5E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079DB78 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079DBE0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079DBFD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079DC69 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079DCD3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0079DCED | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0079DD5D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079DDD0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079DE41 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079DEB0 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079DF1C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079DF37 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0079DFAC | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079E013 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079E075 | `Photos_Screen` | Known | Screen layout |
| 0x0079E0D9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079E0F7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0079E169 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079E186 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079E1EC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079E207 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079E270 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079E28D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079E304 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079E328 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079E396 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0079E3B1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079E46C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079E488 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079E4F6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079E513 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079E57E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079E59E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079E615 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079E631 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079E6A1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079E6C0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079E72C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079E740 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0079E7B9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079E82D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079E89D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079E905 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E919 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079E97D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0079E9E4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079E9FE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079EA6C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079EADE | `NoContent_Screen` | Known | Screen layout |
| 0x0079EAF2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079EB5C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0079EBC5 | `No_Photos_Screen` | Known | Screen layout |
| 0x0079EBD9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079EC3F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079ECAD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0079ED1A | `NoContent_Screen` | Known | Screen layout |
| 0x0079ED2E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079ED96 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0079EE00 | `NoContent_Screen` | Known | Screen layout |
| 0x0079EE14 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079EE7B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079EEE5 | `NoContent_Screen` | Known | Screen layout |
| 0x0079EEF9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079EF66 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079EFD8 | `NoContent_Screen` | Known | Screen layout |
| 0x0079EFEC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079F054 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0079F0BD | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0079F0D8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079F13E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079F15A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079F239 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079F252 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0079F2B3 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079F2C7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0079F435 | `Radio_Screen` | Known | Screen layout |
| 0x0079F445 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079F4A6 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079F529 | `LockediPod_Screen` | Known | Screen layout |
| 0x0079F5B1 | `Lock_Screen` | Known | Screen layout |
| 0x0079F5C0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079F623 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0079F685 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079F6A1 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0079F713 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079F732 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079F79A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079F7B4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079F81C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079F839 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079F8A5 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079F90F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0079F929 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0079F999 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079FA0C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079FA7D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079FAEC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079FB58 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079FB73 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0079FBE8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079FC4F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079FCB1 | `Photos_Screen` | Known | Screen layout |
| 0x0079FD15 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079FD33 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0079FDA5 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079FDC2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079FE28 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079FE43 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079FEAC | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079FEC9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079FF40 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079FF64 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079FFD2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0079FFED | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A00A8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A00C4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A0132 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A014F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A01BA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A01DA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A0251 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A026D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A02DD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A02FC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A0368 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A037C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A03F5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A0469 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A04D9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A0541 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0555 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A05B9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A0620 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A063A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A06A8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A071A | `NoContent_Screen` | Known | Screen layout |
| 0x007A072E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A0798 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A0801 | `No_Photos_Screen` | Known | Screen layout |
| 0x007A0815 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A087B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A08E9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A0956 | `NoContent_Screen` | Known | Screen layout |
| 0x007A096A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A09D2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A0A3C | `NoContent_Screen` | Known | Screen layout |
| 0x007A0A50 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A0AB7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A0B21 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0B35 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A0BA2 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A0C14 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0C28 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A0C90 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A0CF9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A0D14 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A0D7A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A0D96 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A0E75 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A0E8E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A0EEF | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A0F03 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A1071 | `Radio_Screen` | Known | Screen layout |
| 0x007A1081 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A10E2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A1165 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A11ED | `Lock_Screen` | Known | Screen layout |
| 0x007A11FC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A125F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A12C1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A12DD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A134F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A136E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A13D6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A13F0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A1458 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A1475 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A14E1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A154B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A1565 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A15D5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A1648 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A16B9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A1728 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A1794 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A17AF | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A1824 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A188B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A18ED | `Photos_Screen` | Known | Screen layout |
| 0x007A1951 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A196F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A19E1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A19FE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A1A64 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A1A7F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A1AE8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A1B05 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A1B7C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A1BA0 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A1C0E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A1C29 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A1CE4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A1D00 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A1D6E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A1D8B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A1DF6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A1E16 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A1E8D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A1EA9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A1F19 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A1F38 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A1FA4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A1FB8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A2031 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A20A5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A2115 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A217D | `NoContent_Screen` | Known | Screen layout |
| 0x007A2191 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A21F5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A225C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A2276 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A22E4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A2356 | `NoContent_Screen` | Known | Screen layout |
| 0x007A236A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A23D4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A243D | `No_Photos_Screen` | Known | Screen layout |
| 0x007A2451 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A24B7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A2525 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A2592 | `NoContent_Screen` | Known | Screen layout |
| 0x007A25A6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A260E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A2678 | `NoContent_Screen` | Known | Screen layout |
| 0x007A268C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A26F3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A275D | `NoContent_Screen` | Known | Screen layout |
| 0x007A2771 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A27DE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A2850 | `NoContent_Screen` | Known | Screen layout |
| 0x007A2864 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A28CC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A2935 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A2950 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A29B6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A29D2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A2AB1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A2ACA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A2B2B | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A2B3F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A2CAD | `Radio_Screen` | Known | Screen layout |
| 0x007A2CBD | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A2D1E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A2DA1 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A2E29 | `Lock_Screen` | Known | Screen layout |
| 0x007A2E38 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A2E9B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A2EFD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A2F19 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A2F8B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A2FAA | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A3012 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A302C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A3094 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A30B1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A311D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A3187 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A31A1 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A3211 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A3284 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A32F5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A3364 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A33D0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A33EB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A3460 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A34C7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A3529 | `Photos_Screen` | Known | Screen layout |
| 0x007A358D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A35AB | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A361D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A363A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A36A0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A36BB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A3724 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A3741 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A37B8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A37DC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A384A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A3865 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A3920 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A393C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A39AA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A39C7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A3A32 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A3A52 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A3AC9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A3AE5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A3B55 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A3B74 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A3BE0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A3BF4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A3C6D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A3CE1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A3D51 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A3DB9 | `NoContent_Screen` | Known | Screen layout |
| 0x007A3DCD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A3E31 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A3E98 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A3EB2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A3F20 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A3F92 | `NoContent_Screen` | Known | Screen layout |
| 0x007A3FA6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A4010 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A4079 | `No_Photos_Screen` | Known | Screen layout |
| 0x007A408D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A40F3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A4161 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A41CE | `NoContent_Screen` | Known | Screen layout |
| 0x007A41E2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A424A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A42B4 | `NoContent_Screen` | Known | Screen layout |
| 0x007A42C8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A432F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A4399 | `NoContent_Screen` | Known | Screen layout |
| 0x007A43AD | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A441A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A448C | `NoContent_Screen` | Known | Screen layout |
| 0x007A44A0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A4508 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A4571 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A458C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A45F2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A460E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A46ED | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A4706 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A4767 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A477B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A48E9 | `Radio_Screen` | Known | Screen layout |
| 0x007A48F9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A495A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A49DD | `LockediPod_Screen` | Known | Screen layout |
| 0x007A4A65 | `Lock_Screen` | Known | Screen layout |
| 0x007A4A74 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A4AD7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A4B39 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A4B55 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A4BC7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A4BE6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A4C4E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A4C68 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A4CD0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A4CED | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A4D59 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A4DC3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A4DDD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A4E4D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A4EC0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A4F31 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A4FA0 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A500C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A5027 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A509C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A5103 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A5165 | `Photos_Screen` | Known | Screen layout |
| 0x007A51C9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A51E7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A5259 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A5276 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A52DC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A52F7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A5360 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A537D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A53F4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A5418 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A5486 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A54A1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A555C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A5578 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A55E6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A5603 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A566E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A568E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A5705 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A5721 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A5791 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A57B0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A581C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A5830 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A58A9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A591D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A598D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A59F5 | `NoContent_Screen` | Known | Screen layout |
| 0x007A5A09 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A5A6D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A5AD4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A5AEE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A5B5C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A5BCE | `NoContent_Screen` | Known | Screen layout |
| 0x007A5BE2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A5C4C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A5CB5 | `No_Photos_Screen` | Known | Screen layout |
| 0x007A5CC9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A5D2F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A5D9D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A5E0A | `NoContent_Screen` | Known | Screen layout |
| 0x007A5E1E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A5E86 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A5EF0 | `NoContent_Screen` | Known | Screen layout |
| 0x007A5F04 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A5F6B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A5FD5 | `NoContent_Screen` | Known | Screen layout |
| 0x007A5FE9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A6056 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A60C8 | `NoContent_Screen` | Known | Screen layout |
| 0x007A60DC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A6144 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A61AD | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A61C8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A622E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A624A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A6329 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A6342 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A63A3 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A63B7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A6525 | `Radio_Screen` | Known | Screen layout |
| 0x007A6535 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A6596 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A6619 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A66A1 | `Lock_Screen` | Known | Screen layout |
| 0x007A66B0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A6713 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A6775 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A6791 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A6803 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A6822 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A688A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A68A4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A690C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A6929 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A6995 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A69FF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A6A19 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A6A89 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A6AFC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A6B6D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A6BDC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A6C48 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A6C63 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A6CD8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A6D3F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A6DA1 | `Photos_Screen` | Known | Screen layout |
| 0x007A6E05 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A6E23 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A6E95 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A6EB2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A6F18 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A6F33 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A6F9C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A6FB9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A7030 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A7054 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A70C2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A70DD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A7198 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A71B4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A7222 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A723F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A72AA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A72CA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A7341 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A735D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A73CD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A73EC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A7458 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A746C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A74E5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A7559 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A75C9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A7631 | `NoContent_Screen` | Known | Screen layout |
| 0x007A7645 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A76A9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A7710 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A772A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A7798 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A780A | `NoContent_Screen` | Known | Screen layout |
| 0x007A781E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A7888 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A78F1 | `No_Photos_Screen` | Known | Screen layout |
| 0x007A7905 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A796B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A79D9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A7A46 | `NoContent_Screen` | Known | Screen layout |
| 0x007A7A5A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A7AC2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A7B2C | `NoContent_Screen` | Known | Screen layout |
| 0x007A7B40 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A7BA7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A7C11 | `NoContent_Screen` | Known | Screen layout |
| 0x007A7C25 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A7C92 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A7D04 | `NoContent_Screen` | Known | Screen layout |
| 0x007A7D18 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A7D80 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A7DE9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A7E04 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A7E6A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A7E86 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A7F65 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A7F7E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A7FDF | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A7FF3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A8161 | `Radio_Screen` | Known | Screen layout |
| 0x007A8171 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A81D2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A8255 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A82DD | `Lock_Screen` | Known | Screen layout |
| 0x007A82EC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A834F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A83B1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A83CD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A843F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A845E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A84C6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A84E0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A8548 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A8565 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A85D1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A863B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A8655 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A86C5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A8738 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A87A9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A8818 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A8884 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A889F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A8914 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A897B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A89DD | `Photos_Screen` | Known | Screen layout |
| 0x007A8A41 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A8A5F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A8AD1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A8AEE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A8B54 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A8B6F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A8BD8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A8BF5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A8C6C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A8C90 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A8CFE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A8D19 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A8DD4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A8DF0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A8E5E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A8E7B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A8EE6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A8F06 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A8F7D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A8F99 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A9009 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A9028 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A9094 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A90A8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A9121 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A9195 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A9205 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A926D | `NoContent_Screen` | Known | Screen layout |
| 0x007A9281 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A92E5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A934C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A9366 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A93D4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A9446 | `NoContent_Screen` | Known | Screen layout |
| 0x007A945A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A94C4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A952D | `No_Photos_Screen` | Known | Screen layout |
| 0x007A9541 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A95A7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A9615 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A9682 | `NoContent_Screen` | Known | Screen layout |
| 0x007A9696 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A96FE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A9768 | `NoContent_Screen` | Known | Screen layout |
| 0x007A977C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A97E3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A984D | `NoContent_Screen` | Known | Screen layout |
| 0x007A9861 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A98CE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A9940 | `NoContent_Screen` | Known | Screen layout |
| 0x007A9954 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A99BC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A9A25 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A9A40 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A9AA6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A9AC2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A9BA1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A9BBA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A9C1B | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A9C2F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A9D9D | `Radio_Screen` | Known | Screen layout |
| 0x007A9DAD | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A9E0E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A9E91 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A9F19 | `Lock_Screen` | Known | Screen layout |
| 0x007A9F28 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A9F8B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A9FED | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007AA009 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007AA07B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007AA09A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007AA102 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AA11C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007AA184 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AA1A1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AA20D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007AA277 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007AA291 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007AA301 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007AA374 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007AA3E5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007AA454 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007AA4C0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007AA4DB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007AA550 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007AA5B7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007AA619 | `Photos_Screen` | Known | Screen layout |
| 0x007AA67D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007AA69B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007AA70D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007AA72A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007AA790 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007AA7AB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007AA814 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007AA831 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007AA8A8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007AA8CC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007AA93A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007AA955 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007AAA10 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AAA2C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AAA9A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AAAB7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AAB22 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007AAB42 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007AABB9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AABD5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AAC45 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007AAC64 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007AACD0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007AACE4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007AAD5D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007AADD1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007AAE41 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007AAEA9 | `NoContent_Screen` | Known | Screen layout |
| 0x007AAEBD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007AAF21 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007AAF88 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AAFA2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007AB010 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007AB082 | `NoContent_Screen` | Known | Screen layout |
| 0x007AB096 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007AB100 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007AB169 | `No_Photos_Screen` | Known | Screen layout |
| 0x007AB17D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007AB1E3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AB251 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007AB2BE | `NoContent_Screen` | Known | Screen layout |
| 0x007AB2D2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007AB33A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007AB3A4 | `NoContent_Screen` | Known | Screen layout |
| 0x007AB3B8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007AB41F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007AB489 | `NoContent_Screen` | Known | Screen layout |
| 0x007AB49D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007AB50A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007AB57C | `NoContent_Screen` | Known | Screen layout |
| 0x007AB590 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AB5F8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007AB661 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007AB67C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007AB6E2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007AB6FE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007AB7DD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007AB7F6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007AB857 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007AB86B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007AB9D9 | `Radio_Screen` | Known | Screen layout |
| 0x007AB9E9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007ABA4A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007ABACD | `LockediPod_Screen` | Known | Screen layout |
| 0x007ABB55 | `Lock_Screen` | Known | Screen layout |
| 0x007ABB64 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007ABBC7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007ABC29 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007ABC45 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007ABCB7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007ABCD6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007ABD3E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007ABD58 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007ABDC0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007ABDDD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007ABE49 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007ABEB3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007ABECD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007ABF3D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007ABFB0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007AC021 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007AC090 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007AC0FC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007AC117 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007AC18C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007AC1F3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007AC255 | `Photos_Screen` | Known | Screen layout |
| 0x007AC2B9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007AC2D7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007AC349 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007AC366 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007AC3CC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007AC3E7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007AC450 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007AC46D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007AC4E4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007AC508 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007AC576 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007AC591 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007AC64C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AC668 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AC6D6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AC6F3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AC75E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007AC77E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007AC7F5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AC811 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AC881 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007AC8A0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007AC90C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007AC920 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007AC999 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007ACA0D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007ACA7D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007ACAE5 | `NoContent_Screen` | Known | Screen layout |
| 0x007ACAF9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007ACB5D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007ACBC4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007ACBDE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007ACC4C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007ACCBE | `NoContent_Screen` | Known | Screen layout |
| 0x007ACCD2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007ACD3C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007ACDA5 | `No_Photos_Screen` | Known | Screen layout |
| 0x007ACDB9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007ACE1F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007ACE8D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007ACEFA | `NoContent_Screen` | Known | Screen layout |
| 0x007ACF0E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007ACF76 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007ACFE0 | `NoContent_Screen` | Known | Screen layout |
| 0x007ACFF4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007AD05B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007AD0C5 | `NoContent_Screen` | Known | Screen layout |
| 0x007AD0D9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007AD146 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007AD1B8 | `NoContent_Screen` | Known | Screen layout |
| 0x007AD1CC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AD234 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007AD29D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007AD2B8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007AD31E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007AD33A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007AD419 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007AD432 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007AD493 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007AD4A7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007AD615 | `Radio_Screen` | Known | Screen layout |
| 0x007AD625 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007AD686 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007AD709 | `LockediPod_Screen` | Known | Screen layout |
| 0x007AD791 | `Lock_Screen` | Known | Screen layout |
| 0x007AD7A0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007AD803 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007AD865 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007AD881 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007AD8F3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007AD912 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007AD97A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AD994 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007AD9FC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007ADA19 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007ADA85 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007ADAEF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007ADB09 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007ADB79 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007ADBEC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007ADC5D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007ADCCC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007ADD38 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007ADD53 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007ADDC8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007ADE2F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007ADE91 | `Photos_Screen` | Known | Screen layout |
| 0x007ADEF5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007ADF13 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007ADF85 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007ADFA2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007AE008 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007AE023 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007AE08C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007AE0A9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007AE120 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007AE144 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007AE1B2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007AE1CD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007AE288 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AE2A4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AE312 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AE32F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AE39A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007AE3BA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007AE431 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AE44D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AE4BD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007AE4DC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007AE548 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007AE55C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007AE5D5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007AE649 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007AE6B9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007AE721 | `NoContent_Screen` | Known | Screen layout |
| 0x007AE735 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007AE799 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007AE800 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AE81A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007AE888 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007AE8FA | `NoContent_Screen` | Known | Screen layout |
| 0x007AE90E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007AE978 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007AE9E1 | `No_Photos_Screen` | Known | Screen layout |
| 0x007AE9F5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007AEA5B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AEAC9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007AEB36 | `NoContent_Screen` | Known | Screen layout |
| 0x007AEB4A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007AEBB2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007AEC1C | `NoContent_Screen` | Known | Screen layout |
| 0x007AEC30 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007AEC97 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007AED01 | `NoContent_Screen` | Known | Screen layout |
| 0x007AED15 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007AED82 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007AEDF4 | `NoContent_Screen` | Known | Screen layout |
| 0x007AEE08 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AEE70 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007AEED9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007AEEF4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007AEF5A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007AEF76 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007AF055 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007AF06E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007AF0CF | `FirstBoot_Screen` | Known | Screen layout |
| 0x007AF0E3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007AF251 | `Radio_Screen` | Known | Screen layout |
| 0x007AF261 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007AF2C2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007AF345 | `LockediPod_Screen` | Known | Screen layout |
| 0x007AF3CD | `Lock_Screen` | Known | Screen layout |
| 0x007AF3DC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007AF43F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007AF4A1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007AF4BD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007AF52F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007AF54E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007AF5B6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AF5D0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007AF638 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AF655 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AF6C1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007AF72B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007AF745 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007AF7B5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007AF828 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007AF899 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007AF908 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007AF974 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007AF98F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007AFA04 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007AFA6B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007AFACD | `Photos_Screen` | Known | Screen layout |
| 0x007AFB31 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007AFB4F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007AFBC1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007AFBDE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007AFC44 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007AFC5F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007AFCC8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007AFCE5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007AFD5C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007AFD80 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007AFDEE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007AFE09 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007AFEC4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AFEE0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AFF4E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AFF6B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AFFD6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007AFFF6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B006D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B0089 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B00F9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B0118 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B0184 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B0198 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B0211 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B0285 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B02F5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B035D | `NoContent_Screen` | Known | Screen layout |
| 0x007B0371 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B03D5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B043C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B0456 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B04C4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B0536 | `NoContent_Screen` | Known | Screen layout |
| 0x007B054A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B05B4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B061D | `No_Photos_Screen` | Known | Screen layout |
| 0x007B0631 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B0697 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B0705 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B0772 | `NoContent_Screen` | Known | Screen layout |
| 0x007B0786 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B07EE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B0858 | `NoContent_Screen` | Known | Screen layout |
| 0x007B086C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B08D3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B093D | `NoContent_Screen` | Known | Screen layout |
| 0x007B0951 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B09BE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B0A30 | `NoContent_Screen` | Known | Screen layout |
| 0x007B0A44 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B0AAC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B0B15 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B0B30 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B0B96 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B0BB2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B0C91 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B0CAA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B0D0B | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B0D1F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B0E8D | `Radio_Screen` | Known | Screen layout |
| 0x007B0E9D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B0EFE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B0F81 | `LockediPod_Screen` | Known | Screen layout |
| 0x007B1009 | `Lock_Screen` | Known | Screen layout |
| 0x007B1018 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B107B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B10DD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B10F9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B116B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B118A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B11F2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B120C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B1274 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B1291 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B12FD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B1367 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B1381 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B13F1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B1464 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B14D5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B1544 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B15B0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B15CB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B1640 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B16A7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B1709 | `Photos_Screen` | Known | Screen layout |
| 0x007B176D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B178B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B17FD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B181A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B1880 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B189B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B1904 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B1921 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B1998 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B19BC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B1A2A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B1A45 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B1B00 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B1B1C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B1B8A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B1BA7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B1C12 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B1C32 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B1CA9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B1CC5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B1D35 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B1D54 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B1DC0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B1DD4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B1E4D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B1EC1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B1F31 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B1F99 | `NoContent_Screen` | Known | Screen layout |
| 0x007B1FAD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B2011 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B2078 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B2092 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B2100 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B2172 | `NoContent_Screen` | Known | Screen layout |
| 0x007B2186 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B21F0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B2259 | `No_Photos_Screen` | Known | Screen layout |
| 0x007B226D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B22D3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B2341 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B23AE | `NoContent_Screen` | Known | Screen layout |
| 0x007B23C2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B242A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B2494 | `NoContent_Screen` | Known | Screen layout |
| 0x007B24A8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B250F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B2579 | `NoContent_Screen` | Known | Screen layout |
| 0x007B258D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B25FA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B266C | `NoContent_Screen` | Known | Screen layout |
| 0x007B2680 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B26E8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B2751 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B276C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B27D2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B27EE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B28CD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B28E6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B2947 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B295B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B2AC9 | `Radio_Screen` | Known | Screen layout |
| 0x007B2AD9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B2B3A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B2BBD | `LockediPod_Screen` | Known | Screen layout |
| 0x007B2C45 | `Lock_Screen` | Known | Screen layout |
| 0x007B2C54 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B2CB7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B2D19 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B2D35 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B2DA7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B2DC6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B2E2E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B2E48 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B2EB0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B2ECD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B2F39 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B2FA3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B2FBD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B302D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B30A0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B3111 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B3180 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B31EC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B3207 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B327C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B32E3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B3345 | `Photos_Screen` | Known | Screen layout |
| 0x007B33A9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B33C7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B3439 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B3456 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B34BC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B34D7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B3540 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B355D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B35D4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B35F8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B3666 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B3681 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B373C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B3758 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B37C6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B37E3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B384E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B386E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B38E5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B3901 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B3971 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B3990 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B39FC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B3A10 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B3A89 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B3AFD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B3B6D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B3BD5 | `NoContent_Screen` | Known | Screen layout |
| 0x007B3BE9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B3C4D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B3CB4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B3CCE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B3D3C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B3DAE | `NoContent_Screen` | Known | Screen layout |
| 0x007B3DC2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B3E2C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B3E95 | `No_Photos_Screen` | Known | Screen layout |
| 0x007B3EA9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B3F0F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B3F7D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B3FEA | `NoContent_Screen` | Known | Screen layout |
| 0x007B3FFE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B4066 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B40D0 | `NoContent_Screen` | Known | Screen layout |
| 0x007B40E4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B414B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B41B5 | `NoContent_Screen` | Known | Screen layout |
| 0x007B41C9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B4236 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B42A8 | `NoContent_Screen` | Known | Screen layout |
| 0x007B42BC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B4324 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B438D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B43A8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B440E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B442A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B4509 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B4522 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B4583 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B4597 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B4705 | `Radio_Screen` | Known | Screen layout |
| 0x007B4715 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B4776 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B47F9 | `LockediPod_Screen` | Known | Screen layout |
| 0x007B4881 | `Lock_Screen` | Known | Screen layout |
| 0x007B4890 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B48F3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B4955 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B4971 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B49E3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B4A02 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B4A6A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B4A84 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B4AEC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B4B09 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B4B75 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B4BDF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B4BF9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B4C69 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B4CDC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B4D4D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B4DBC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B4E28 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B4E43 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B4EB8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B4F1F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B4F81 | `Photos_Screen` | Known | Screen layout |
| 0x007B4FE5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B5003 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B5075 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B5092 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B50F8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B5113 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B517C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B5199 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B5210 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B5234 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B52A2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B52BD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B5378 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B5394 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B5402 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B541F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B548A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B54AA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B5521 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B553D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B55AD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B55CC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B5638 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B564C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B56C5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B5739 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B57A9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B5811 | `NoContent_Screen` | Known | Screen layout |
| 0x007B5825 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B5889 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B58F0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B590A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B5978 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B59EA | `NoContent_Screen` | Known | Screen layout |
| 0x007B59FE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B5A68 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B5AD1 | `No_Photos_Screen` | Known | Screen layout |
| 0x007B5AE5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B5B4B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B5BB9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B5C26 | `NoContent_Screen` | Known | Screen layout |
| 0x007B5C3A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B5CA2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B5D0C | `NoContent_Screen` | Known | Screen layout |
| 0x007B5D20 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B5D87 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B5DF1 | `NoContent_Screen` | Known | Screen layout |
| 0x007B5E05 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B5E72 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B5EE4 | `NoContent_Screen` | Known | Screen layout |
| 0x007B5EF8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B5F60 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B5FC9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B5FE4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B604A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B6066 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B6145 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B615E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B61BF | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B61D3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B6341 | `Radio_Screen` | Known | Screen layout |
| 0x007B6351 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B63B2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B6435 | `LockediPod_Screen` | Known | Screen layout |
| 0x007B64BD | `Lock_Screen` | Known | Screen layout |
| 0x007B64CC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B652F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B6591 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B65AD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B661F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B663E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B66A6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B66C0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B6728 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B6745 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B67B1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B681B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B6835 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B68A5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B6918 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B6989 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B69F8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B6A64 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B6A7F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B6AF4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B6B5B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B6BBD | `Photos_Screen` | Known | Screen layout |
| 0x007B6C21 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B6C3F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B6CB1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B6CCE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B6D34 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B6D4F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B6DB8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B6DD5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B6E4C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B6E70 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B6EDE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B6EF9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B6FB4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B6FD0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B703E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B705B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B70C6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B70E6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B715D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B7179 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B71E9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B7208 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B7274 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B7288 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B7301 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B7375 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B73E5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B744D | `NoContent_Screen` | Known | Screen layout |
| 0x007B7461 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B74C5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B752C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B7546 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B75B4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B7626 | `NoContent_Screen` | Known | Screen layout |
| 0x007B763A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B76A4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B770D | `No_Photos_Screen` | Known | Screen layout |
| 0x007B7721 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B7787 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B77F5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B7862 | `NoContent_Screen` | Known | Screen layout |
| 0x007B7876 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B78DE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B7948 | `NoContent_Screen` | Known | Screen layout |
| 0x007B795C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B79C3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B7A2D | `NoContent_Screen` | Known | Screen layout |
| 0x007B7A41 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B7AAE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B7B20 | `NoContent_Screen` | Known | Screen layout |
| 0x007B7B34 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B7B9C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B7C05 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B7C20 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B7C86 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B7CA2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B7D81 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B7D9A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B7DFB | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B7E0F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B7F7D | `Radio_Screen` | Known | Screen layout |
| 0x007B7F8D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B7FEE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B8071 | `LockediPod_Screen` | Known | Screen layout |
| 0x007B80F9 | `Lock_Screen` | Known | Screen layout |
| 0x007B8108 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B816B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B81CD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B81E9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B825B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B827A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B82E2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B82FC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B8364 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B8381 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B83ED | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B8457 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B8471 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B84E1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B8554 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B85C5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B8634 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B86A0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B86BB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B8730 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B8797 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B87F9 | `Photos_Screen` | Known | Screen layout |
| 0x007B885D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B887B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B88ED | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B890A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B8970 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B898B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B89F4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B8A11 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B8A88 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B8AAC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B8B1A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B8B35 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B8BF0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B8C0C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B8C7A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B8C97 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B8D02 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B8D22 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B8D99 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B8DB5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B8E25 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B8E44 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B8EB0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B8EC4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B8F3D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B8FB1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B9021 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B9089 | `NoContent_Screen` | Known | Screen layout |
| 0x007B909D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B9101 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B9168 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B9182 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B91F0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B9262 | `NoContent_Screen` | Known | Screen layout |
| 0x007B9276 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B92E0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B9349 | `No_Photos_Screen` | Known | Screen layout |
| 0x007B935D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B93C3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B9431 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B949E | `NoContent_Screen` | Known | Screen layout |
| 0x007B94B2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B951A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B9584 | `NoContent_Screen` | Known | Screen layout |
| 0x007B9598 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B95FF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B9669 | `NoContent_Screen` | Known | Screen layout |
| 0x007B967D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B96EA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B975C | `NoContent_Screen` | Known | Screen layout |
| 0x007B9770 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B97D8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B9841 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B985C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B98C2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B98DE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B99BD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B99D6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B9A37 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B9A4B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B9BB9 | `Radio_Screen` | Known | Screen layout |
| 0x007B9BC9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B9C2A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B9CAD | `LockediPod_Screen` | Known | Screen layout |
| 0x007B9D35 | `Lock_Screen` | Known | Screen layout |
| 0x007B9D44 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B9DA7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B9E09 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B9E25 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B9E97 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B9EB6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B9F1E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B9F38 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B9FA0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B9FBD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BA029 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BA093 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007BA0AD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007BA11D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007BA190 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007BA201 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007BA270 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007BA2DC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007BA2F7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007BA36C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007BA3D3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007BA435 | `Photos_Screen` | Known | Screen layout |
| 0x007BA499 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007BA4B7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007BA529 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007BA546 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007BA5AC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BA5C7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BA630 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007BA64D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007BA6C4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007BA6E8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007BA756 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007BA771 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007BA82C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BA848 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BA8B6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BA8D3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BA93E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BA95E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BA9D5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BA9F1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BAA61 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BAA80 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BAAEC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BAB00 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007BAB79 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BABED | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007BAC5D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007BACC5 | `NoContent_Screen` | Known | Screen layout |
| 0x007BACD9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007BAD3D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007BADA4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007BADBE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007BAE2C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007BAE9E | `NoContent_Screen` | Known | Screen layout |
| 0x007BAEB2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007BAF1C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007BAF85 | `No_Photos_Screen` | Known | Screen layout |
| 0x007BAF99 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007BAFFF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007BB06D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007BB0DA | `NoContent_Screen` | Known | Screen layout |
| 0x007BB0EE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007BB156 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007BB1C0 | `NoContent_Screen` | Known | Screen layout |
| 0x007BB1D4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007BB23B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007BB2A5 | `NoContent_Screen` | Known | Screen layout |
| 0x007BB2B9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007BB326 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007BB398 | `NoContent_Screen` | Known | Screen layout |
| 0x007BB3AC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007BB414 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007BB47D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007BB498 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007BB4FE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BB51A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BB5F9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007BB612 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007BB673 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007BB687 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007BB7F5 | `Radio_Screen` | Known | Screen layout |
| 0x007BB805 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007BB866 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007BB8E9 | `LockediPod_Screen` | Known | Screen layout |
| 0x007BB971 | `Lock_Screen` | Known | Screen layout |
| 0x007BB980 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007BB9E3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007BBA45 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007BBA61 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007BBAD3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BBAF2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BBB5A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007BBB74 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007BBBDC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BBBF9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BBC65 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BBCCF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007BBCE9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007BBD59 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007BBDCC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007BBE3D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007BBEAC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007BBF18 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007BBF33 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007BBFA8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007BC00F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007BC071 | `Photos_Screen` | Known | Screen layout |
| 0x007BC0D5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007BC0F3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007BC165 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007BC182 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007BC1E8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BC203 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BC26C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007BC289 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007BC300 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007BC324 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007BC392 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007BC3AD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007BC468 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BC484 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BC4F2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BC50F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BC57A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BC59A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BC611 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BC62D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BC69D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BC6BC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BC728 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BC73C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007BC7B5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BC829 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007BC899 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007BC901 | `NoContent_Screen` | Known | Screen layout |
| 0x007BC915 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007BC979 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007BC9E0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007BC9FA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007BCA68 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007BCADA | `NoContent_Screen` | Known | Screen layout |
| 0x007BCAEE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007BCB58 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007BCBC1 | `No_Photos_Screen` | Known | Screen layout |
| 0x007BCBD5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007BCC3B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007BCCA9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007BCD16 | `NoContent_Screen` | Known | Screen layout |
| 0x007BCD2A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007BCD92 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007BCDFC | `NoContent_Screen` | Known | Screen layout |
| 0x007BCE10 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007BCE77 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007BCEE1 | `NoContent_Screen` | Known | Screen layout |
| 0x007BCEF5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007BCF62 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007BCFD4 | `NoContent_Screen` | Known | Screen layout |
| 0x007BCFE8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007BD050 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007BD0B9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007BD0D4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007BD13A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BD156 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BD235 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007BD24E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007BD2AF | `FirstBoot_Screen` | Known | Screen layout |
| 0x007BD2C3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007BD431 | `Radio_Screen` | Known | Screen layout |
| 0x007BD441 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007BD4A2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007BD525 | `LockediPod_Screen` | Known | Screen layout |
| 0x007BD5AD | `Lock_Screen` | Known | Screen layout |
| 0x007BD5BC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007BD61F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007BD681 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007BD69D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007BD70F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BD72E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BD796 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007BD7B0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007BD818 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BD835 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BD8A1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BD90B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007BD925 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007BD995 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007BDA08 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007BDA79 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007BDAE8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007BDB54 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007BDB6F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007BDBE4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007BDC4B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007BDCAD | `Photos_Screen` | Known | Screen layout |
| 0x007BDD11 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007BDD2F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007BDDA1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007BDDBE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007BDE24 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BDE3F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BDEA8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007BDEC5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007BDF3C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007BDF60 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007BDFCE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007BDFE9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007BE089 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BE0A5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BE113 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BE130 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BE19B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BE1BB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BE232 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BE24E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BE2BE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BE2DD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BE349 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BE35D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007BE3D2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007BE43D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007BE4AC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007BE51D | `NoContent_Screen` | Known | Screen layout |
| 0x007BE531 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BE5A0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007BE613 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007BE680 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007BE6E9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007BE759 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007BE7C9 | `NoContent_Screen` | Known | Screen layout |
| 0x007BE7DD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007BE840 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007BE8A3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BE8BF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BE97F | `Radio_Screen` | Known | Screen layout |
| 0x007BE98F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007BE9F0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007BEA5E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BEA7D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BEAEB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BEB50 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BEB6B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BEC11 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BEC2D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BEC9B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BECB8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BED23 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BED43 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BEDBA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BEDD6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BEE46 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BEE65 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BEED1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BEEE5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007BEF5A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007BEFC5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007BF034 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007BF0A5 | `NoContent_Screen` | Known | Screen layout |
| 0x007BF0B9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BF128 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007BF19B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007BF208 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007BF271 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007BF2E1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007BF351 | `NoContent_Screen` | Known | Screen layout |
| 0x007BF365 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007BF3C8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007BF42B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BF447 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BF507 | `Radio_Screen` | Known | Screen layout |
| 0x007BF517 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007BF578 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007BF5E6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BF605 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BF673 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BF6D8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BF6F3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BF799 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BF7B5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BF823 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BF840 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BF8AB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BF8CB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BF942 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BF95E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BF9CE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BF9ED | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BFA59 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BFA6D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007BFAE2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007BFB4D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007BFBBC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007BFC2D | `NoContent_Screen` | Known | Screen layout |
| 0x007BFC41 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BFCB0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007BFD23 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007BFD90 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007BFDF9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007BFE69 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007BFED9 | `NoContent_Screen` | Known | Screen layout |
| 0x007BFEED | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007BFF50 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007BFFB3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BFFCF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C008F | `Radio_Screen` | Known | Screen layout |
| 0x007C009F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C0100 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C016E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C018D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C01FB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C0260 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C027B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C0321 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C033D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C03AB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C03C8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C0433 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C0453 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C04CA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C04E6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C0556 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C0575 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C05E1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C05F5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C066A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C06D5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C0744 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C07B5 | `NoContent_Screen` | Known | Screen layout |
| 0x007C07C9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C0838 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C08AB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C0918 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C0981 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C09F1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C0A61 | `NoContent_Screen` | Known | Screen layout |
| 0x007C0A75 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C0AD8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C0B3B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C0B57 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C0C17 | `Radio_Screen` | Known | Screen layout |
| 0x007C0C27 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C0C88 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C0CF6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C0D15 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C0D83 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C0DE8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C0E03 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C0EA9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C0EC5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C0F33 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C0F50 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C0FBB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C0FDB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C1052 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C106E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C10DE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C10FD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C1169 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C117D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C11F2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C125D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C12CC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C133D | `NoContent_Screen` | Known | Screen layout |
| 0x007C1351 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C13C0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C1433 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C14A0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C1509 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C1579 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C15E9 | `NoContent_Screen` | Known | Screen layout |
| 0x007C15FD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C1660 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C16C3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C16DF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C179F | `Radio_Screen` | Known | Screen layout |
| 0x007C17AF | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C1810 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C187E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C189D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C190B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C1970 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C198B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C1A31 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C1A4D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C1ABB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C1AD8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C1B43 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C1B63 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C1BDA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C1BF6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C1C66 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C1C85 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C1CF1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C1D05 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C1D7A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C1DE5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C1E54 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C1EC5 | `NoContent_Screen` | Known | Screen layout |
| 0x007C1ED9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C1F48 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C1FBB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C2028 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C2091 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C2101 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C2171 | `NoContent_Screen` | Known | Screen layout |
| 0x007C2185 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C21E8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C224B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C2267 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C2327 | `Radio_Screen` | Known | Screen layout |
| 0x007C2337 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C2398 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C2406 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C2425 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C2493 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C24F8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C2513 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C25B9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C25D5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C2643 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C2660 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C26CB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C26EB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C2762 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C277E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C27EE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C280D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C2879 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C288D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C2902 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C296D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C29DC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C2A4D | `NoContent_Screen` | Known | Screen layout |
| 0x007C2A61 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C2AD0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C2B43 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C2BB0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C2C19 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C2C89 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C2CF9 | `NoContent_Screen` | Known | Screen layout |
| 0x007C2D0D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C2D70 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C2DD3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C2DEF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C2EAF | `Radio_Screen` | Known | Screen layout |
| 0x007C2EBF | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C2F20 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C2F8E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C2FAD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C301B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C3080 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C309B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C3141 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C315D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C31CB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C31E8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C3253 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C3273 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C32EA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C3306 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C3376 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C3395 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C3401 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C3415 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C348A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C34F5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C3564 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C35D5 | `NoContent_Screen` | Known | Screen layout |
| 0x007C35E9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C3658 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C36CB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C3738 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C37A1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C3811 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C3881 | `NoContent_Screen` | Known | Screen layout |
| 0x007C3895 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C38F8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C395B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C3977 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C3A37 | `Radio_Screen` | Known | Screen layout |
| 0x007C3A47 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C3AA8 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C3B16 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C3B35 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C3BA3 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C3C08 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C3C23 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C3CC9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C3CE5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C3D53 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C3D70 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C3DDB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C3DFB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C3E72 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C3E8E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C3EFE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C3F1D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C3F89 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C3F9D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C4012 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C407D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C40EC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C415D | `NoContent_Screen` | Known | Screen layout |
| 0x007C4171 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C41E0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C4253 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C42C0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C4329 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C4399 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C4409 | `NoContent_Screen` | Known | Screen layout |
| 0x007C441D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C4480 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C44E3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C44FF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C45BF | `Radio_Screen` | Known | Screen layout |
| 0x007C45CF | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C4630 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C469E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C46BD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C472B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C4790 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C47AB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C4851 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C486D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C48DB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C48F8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C4963 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C4983 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C49FA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C4A16 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C4A86 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C4AA5 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C4B11 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C4B25 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C4B9A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C4C05 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C4C74 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C4CE5 | `NoContent_Screen` | Known | Screen layout |
| 0x007C4CF9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C4D68 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C4DDB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C4E48 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C4EB1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C4F21 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C4F91 | `NoContent_Screen` | Known | Screen layout |
| 0x007C4FA5 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C5008 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C506B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C5087 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C5147 | `Radio_Screen` | Known | Screen layout |
| 0x007C5157 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C51B8 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C5226 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C5245 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C52B3 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C5318 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C5333 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C53D9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C53F5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C5463 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C5480 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C54EB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C550B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C5582 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C559E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C560E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C562D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C5699 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C56AD | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C5722 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C578D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C57FC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C586D | `NoContent_Screen` | Known | Screen layout |
| 0x007C5881 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C58F0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C5963 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C59D0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C5A39 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C5AA9 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C5B19 | `NoContent_Screen` | Known | Screen layout |
| 0x007C5B2D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C5B90 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C5BF3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C5C0F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C5CCF | `Radio_Screen` | Known | Screen layout |
| 0x007C5CDF | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C5D40 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C5DAE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C5DCD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C5E3B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C5EA0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C5EBB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C5F61 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C5F7D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C5FEB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C6008 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C6073 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C6093 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C610A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C6126 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C6196 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C61B5 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C6221 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C6235 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C62AA | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C6315 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C6384 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C63F5 | `NoContent_Screen` | Known | Screen layout |
| 0x007C6409 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C6478 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C64EB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C6558 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C65C1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C6631 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C66A1 | `NoContent_Screen` | Known | Screen layout |
| 0x007C66B5 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C6718 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C677B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C6797 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C6857 | `Radio_Screen` | Known | Screen layout |
| 0x007C6867 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C68C8 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C6936 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C6955 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C69C3 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C6A28 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C6A43 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C6B24 | `MediaLists_AudiobookChapters_Screen)` | Known | Screen layout |
| 0x007C6B4B | `MediaLists_AudiobookChapters_Screen_Plane!` | Known | Screen layout |
| 0x007C72E5 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007C7300 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x007C736B | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C7386 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x007C73F9 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x007C7414 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x007C75D1 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007C75EC | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x007C7657 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C7672 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x007C76E5 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x007C7700 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x007C78C8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C78E4 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x007C795F | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C797B | `MediaLists_Albums_Screen_WithTrackCount%` | Known | Screen layout |
| 0x007C79F4 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007C7A0F | `MediaLists_Songs_Screen_Plain1` | Known | Screen layout |
| 0x007C7A8A | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x007C7AA5 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x007C7CC7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C7CE4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C7DC3 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C7DDF | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x007C7E5A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C7E75 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C805B | `MediaLists_PodcastEpisodes_Screen'` | Known | Screen layout |
| 0x007C8080 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007C8352 | `MediaLists_TVSeasons_Screen#` | Known | Screen layout |
| 0x007C8371 | `MediaLists_TVSeasons_Screen_Default%` | Known | Screen layout |
| 0x007C83E6 | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x007C8406 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007C858E | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x007C85AE | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007C89A7 | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x007C89CC | `MediaLists_NestedPlaylists_Screen_Default,` | Known | Screen layout |
| 0x007C8A4E | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x007C8A6D | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x007C8BFD | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x007C8C22 | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x007C8C9A | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x007C8CB9 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x007C8D1D | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C8DCA | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C8E3C | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x007C8F32 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x007C91F4 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x007C92F4 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x007C9360 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007C93CA | `NoContent_Screen` | Known | Screen layout |
| 0x007C93DE | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007C9448 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007C94BC | `NoContent_Screen` | Known | Screen layout |
| 0x007C94D0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007C953B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007C95A7 | `NoContent_Screen` | Known | Screen layout |
| 0x007C95BB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007C9622 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007C968E | `NoContent_Screen` | Known | Screen layout |
| 0x007C96A2 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007C970F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007C9783 | `NoContent_Screen` | Known | Screen layout |
| 0x007C9797 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007C97FF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007C986C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007C98D0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007C98EC | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007C9958 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C9975 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C99E2 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007C9AA9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007C9AC6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007C9B3D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007C9B61 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C9C18 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007C9C82 | `NoContent_Screen` | Known | Screen layout |
| 0x007C9C96 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007C9D00 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007C9D74 | `NoContent_Screen` | Known | Screen layout |
| 0x007C9D88 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007C9DF3 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007C9E5F | `NoContent_Screen` | Known | Screen layout |
| 0x007C9E73 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007C9EDA | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007C9F46 | `NoContent_Screen` | Known | Screen layout |
| 0x007C9F5A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007C9FC7 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CA03B | `NoContent_Screen` | Known | Screen layout |
| 0x007CA04F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CA0B7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CA124 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CA188 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CA1A4 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CA210 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CA22D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CA29A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CA361 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CA37E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CA3F5 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CA419 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CA4D0 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007CA53A | `NoContent_Screen` | Known | Screen layout |
| 0x007CA54E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007CA5B8 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007CA62C | `NoContent_Screen` | Known | Screen layout |
| 0x007CA640 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007CA6AB | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007CA717 | `NoContent_Screen` | Known | Screen layout |
| 0x007CA72B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007CA792 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007CA7FE | `NoContent_Screen` | Known | Screen layout |
| 0x007CA812 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007CA87F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CA8F3 | `NoContent_Screen` | Known | Screen layout |
| 0x007CA907 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CA96F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CA9DC | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CAA40 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CAA5C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CAAC8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CAAE5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CAB52 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CAC19 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CAC36 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CACAD | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CACD1 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CAD88 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007CADF2 | `NoContent_Screen` | Known | Screen layout |
| 0x007CAE06 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007CAE70 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007CAEE4 | `NoContent_Screen` | Known | Screen layout |
| 0x007CAEF8 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007CAF63 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007CAFCF | `NoContent_Screen` | Known | Screen layout |
| 0x007CAFE3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007CB04A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007CB0B6 | `NoContent_Screen` | Known | Screen layout |
| 0x007CB0CA | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007CB137 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CB1AB | `NoContent_Screen` | Known | Screen layout |
| 0x007CB1BF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CB227 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CB294 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CB2F8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CB314 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CB380 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CB39D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CB40A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CB4D1 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CB4EE | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CB565 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CB589 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CB640 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007CB6AA | `NoContent_Screen` | Known | Screen layout |
| 0x007CB6BE | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007CB728 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007CB79C | `NoContent_Screen` | Known | Screen layout |
| 0x007CB7B0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007CB81B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007CB887 | `NoContent_Screen` | Known | Screen layout |
| 0x007CB89B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007CB902 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007CB96E | `NoContent_Screen` | Known | Screen layout |
| 0x007CB982 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007CB9EF | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CBA63 | `NoContent_Screen` | Known | Screen layout |
| 0x007CBA77 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CBADF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CBB4C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CBBB0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CBBCC | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CBC38 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CBC55 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CBCC2 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CBD89 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CBDA6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CBE1D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CBE41 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CBEF8 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007CBF62 | `NoContent_Screen` | Known | Screen layout |
| 0x007CBF76 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007CBFE0 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007CC054 | `NoContent_Screen` | Known | Screen layout |
| 0x007CC068 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007CC0D3 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007CC13F | `NoContent_Screen` | Known | Screen layout |
| 0x007CC153 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007CC1BA | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007CC226 | `NoContent_Screen` | Known | Screen layout |
| 0x007CC23A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007CC2A7 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CC31B | `NoContent_Screen` | Known | Screen layout |
| 0x007CC32F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CC397 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CC404 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CC468 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CC484 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CC4F0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CC50D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CC57A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CC641 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CC65E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CC6D5 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CC6F9 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CC7B0 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007CC81A | `NoContent_Screen` | Known | Screen layout |
| 0x007CC82E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007CC898 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007CC90C | `NoContent_Screen` | Known | Screen layout |
| 0x007CC920 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007CC98B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007CC9F7 | `NoContent_Screen` | Known | Screen layout |
| 0x007CCA0B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007CCA72 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007CCADE | `NoContent_Screen` | Known | Screen layout |
| 0x007CCAF2 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007CCB5F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CCBD3 | `NoContent_Screen` | Known | Screen layout |
| 0x007CCBE7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CCC4F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CCCBC | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CCD20 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CCD3C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CCDA8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CCDC5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CCE32 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CCEF9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CCF16 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CCF8D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CCFB1 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CD068 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007CD0D2 | `NoContent_Screen` | Known | Screen layout |
| 0x007CD0E6 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007CD150 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007CD1C4 | `NoContent_Screen` | Known | Screen layout |
| 0x007CD1D8 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007CD243 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007CD2AF | `NoContent_Screen` | Known | Screen layout |
| 0x007CD2C3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007CD32A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007CD396 | `NoContent_Screen` | Known | Screen layout |
| 0x007CD3AA | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007CD417 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CD48B | `NoContent_Screen` | Known | Screen layout |
| 0x007CD49F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CD507 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CD574 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CD5D8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CD5F4 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CD660 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CD67D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CD6EA | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CD7B1 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CD7CE | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CD845 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CD869 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CDCCC | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CDD3E | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CDDA9 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CDE0E | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CDE78 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CDEE2 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CDF52 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CDFC9 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CE037 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CE0A2 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CE10C | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CE173 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CE1E2 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CE250 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CE2B5 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CE31D | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CE388 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CE3F3 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CE45A | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CE7C8 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CE83A | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CE8A5 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CE90A | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CE974 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CE9DE | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CEA4E | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CEAC5 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CEB33 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CEB9E | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CEC08 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CEC6F | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CECDE | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CED4C | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CEDB1 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CEE19 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CEE84 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CEEEF | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CEF56 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CF2C2 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CF334 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CF39F | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CF404 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CF46E | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CF4D8 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CF548 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CF5BF | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CF62D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CF698 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CF702 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CF769 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CF7D8 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CF846 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CF8AB | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CF913 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CF97E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CF9E9 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CFA50 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CFDBA | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CFE2C | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CFE97 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CFEFC | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CFF66 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CFFD0 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D0040 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D00B7 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D0125 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D0190 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D01FA | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D0261 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D02D0 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D033E | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D03A3 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D040B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D0476 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D04E1 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D0548 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D089A | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D090C | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D0977 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D09DC | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D0A46 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D0AB0 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D0B20 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D0B97 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D0C05 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D0C70 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D0CDA | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D0D41 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D0DB0 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D0E1E | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D0E83 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D0EEB | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D0F56 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D0FC1 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D1028 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D139F | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D1411 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D147C | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D14E1 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D154B | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D15B5 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D1625 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D169C | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D170A | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D1775 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D17DF | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D1846 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D18B5 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D1923 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D1988 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D19F0 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D1A5B | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D1AC6 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D1B2D | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D1EA1 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D1F13 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D1F7E | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D1FE3 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D204D | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D20B7 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D2127 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D219E | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D220C | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D2277 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D22E1 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D2348 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D23B7 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D2425 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D248A | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D24F2 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D255D | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D25C8 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D262F | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D2989 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D29FB | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D2A66 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D2ACB | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D2B35 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D2B9F | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D2C0F | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D2C86 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D2CF4 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D2D5F | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D2DC9 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D2E30 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D2E9F | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D2F0D | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D2F72 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D2FDA | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D3045 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D30B0 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D3117 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D3471 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D34E3 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D354E | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D35B3 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D361D | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D3687 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D36F7 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D376E | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D37DC | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D3847 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D38B1 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D3918 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D3987 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D39F5 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D3A5A | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D3AC2 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D3B2D | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D3B98 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D3BFF | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D3F5A | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D3FCC | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D4037 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D409C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D4106 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D4170 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D41E0 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D4257 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D42C5 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D4330 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D439A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D4401 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D4470 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D44DE | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D4543 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D45AB | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D4616 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D4681 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D46E8 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D4A6C | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D4ADE | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D4B49 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D4BAE | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D4C18 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D4C82 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D4CF2 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D4D69 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D4DD7 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D4E42 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D4EAC | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D4F13 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D4F82 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D4FF0 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D5055 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D50BD | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D5128 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D5193 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D51FA | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D5588 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D55FA | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D5665 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D56CA | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D5734 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D579E | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D580E | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D5885 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D58F3 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D595E | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D59C8 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D5A2F | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D5A9E | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D5B0C | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D5B71 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D5BD9 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D5C44 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D5CAF | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D5D16 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D6084 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D60F6 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D6161 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D61C6 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D6230 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D629A | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D630A | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D6381 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D63EF | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D645A | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D64C4 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D652B | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D659A | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D6608 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D666D | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D66D5 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D6740 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D67AB | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D6812 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D6B78 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D6BEA | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D6C55 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D6CBA | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D6D24 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D6D8E | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D6DFE | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D6E75 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D6EE3 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D6F4E | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D6FB8 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D701F | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D708E | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D70FC | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D7161 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D71C9 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D7234 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D729F | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D7306 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D765A | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D76CC | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D7737 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D779C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D7806 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D7870 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D78E0 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D7957 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D79C5 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D7A30 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D7A9A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D7B01 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D7B70 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D7BDE | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D7C43 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D7CAB | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D7D16 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D7D81 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D7DE8 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D8133 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D81A5 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D8210 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D8275 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D82DF | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D8349 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D83B9 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D8430 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D849E | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D8509 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D8573 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D85DA | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D8649 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D86B7 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D871C | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D8784 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D87EF | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D885A | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D88C1 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D8C23 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D8C95 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D8D00 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D8D65 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D8DCF | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D8E39 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D8EA9 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D8F20 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D8F8E | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D8FF9 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D9063 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D90CA | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D9139 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D91A7 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D920C | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D9274 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D92DF | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D934A | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D93B1 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D96C9 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D973B | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D97A6 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D980B | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D9875 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D98DF | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D994F | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D99C6 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D9A34 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D9A9F | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D9B09 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D9B70 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D9BDF | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D9C4D | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D9CB2 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D9D1A | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D9D85 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D9DF0 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D9E57 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007DA16E | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DA1E5 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DA262 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DA2D4 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DA344 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DA3BA | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DA428 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DA495 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DA7DA | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DA851 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DA8CE | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DA940 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DA9B0 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DAA26 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DAA94 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DAB01 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DAE6A | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DAEE1 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DAF5E | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DAFD0 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DB040 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DB0B6 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DB124 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DB191 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DB4FA | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DB571 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DB5EC | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DB65C | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DB6D2 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DB740 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DB7AD | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DBAE6 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DBB5D | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DBBD8 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DBC48 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DBCBE | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DBD2C | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DBD99 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DC0D0 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DC147 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DC1C2 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DC232 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DC2A8 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DC316 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DC383 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DC693 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DC70A | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DC785 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DC7F5 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DC86B | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DC8D9 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DC946 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DCF4A | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007DCF67 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007DCFE2 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007DCFFB | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007DD073 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007DD08C | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007DD101 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DD117 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007DD18E | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DD1A4 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007DD21B | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007DD238 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007DD2B0 | `Notes_List_Screen` | Known | Screen layout |
| 0x007DD2C5 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007DD476 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007DD493 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007DD50E | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007DD527 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007DD59F | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007DD5B8 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007DD62D | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DD643 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007DD6BA | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DD6D0 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007DD747 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007DD764 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007DD7DC | `Notes_List_Screen` | Known | Screen layout |
| 0x007DD7F1 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007DD9D2 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007DD9EF | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007DDA6A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007DDA83 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007DDAFB | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007DDB14 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007DDB89 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DDB9F | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007DDC16 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DDC2C | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007DDCA3 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007DDCC0 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007DDD38 | `Notes_List_Screen` | Known | Screen layout |
| 0x007DDD4D | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007DDF02 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007DDF1F | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007DDF9A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007DDFB3 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007DE02B | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007DE044 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007DE0B9 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DE0CF | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007DE146 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DE15C | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007DE1D3 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007DE1F0 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007DE268 | `Notes_List_Screen` | Known | Screen layout |
| 0x007DE27D | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007DE595 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007DE63B | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007DE6BE | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007DE776 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x007DE7F8 | `PhotosSettingsSlideshowMusic_Screen+` | Known | Screen layout |
| 0x007DE81F | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x007DE905 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x007DEABD | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007DEB1D | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007DEB7A | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007DEBA1 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007DEC41 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007DECA1 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007DECFE | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007DED25 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007DEFC0 | `Photos_Screen` | Known | Screen layout |
| 0x007DF10C | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007DF170 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007DF1D1 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007DF22E | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007DF28B | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007DF2F9 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007DF356 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007DF4CC | `Photos_Screen` | Known | Screen layout |
| 0x007DF618 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007DF67C | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007DF6DD | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007DF73A | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007DF797 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007DF805 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007DF862 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007DF9D8 | `Photos_Screen` | Known | Screen layout |
| 0x007DFB24 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007DFB88 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007DFBE9 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007DFC46 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007DFCA3 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007DFD11 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007DFD6E | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007DFEE4 | `Photos_Screen` | Known | Screen layout |
| 0x007E0030 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E0094 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007E00F5 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007E0152 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007E01AF | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E021D | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007E027A | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007E03F0 | `Photos_Screen` | Known | Screen layout |
| 0x007E053C | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E05A0 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007E0601 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007E065E | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007E06BB | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E0729 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007E0786 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007E08FC | `Photos_Screen` | Known | Screen layout |
| 0x007E0A48 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E0AAC | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007E0B0D | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007E0B6A | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007E0BC7 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E0C35 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007E0C92 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007E0E08 | `Photos_Screen` | Known | Screen layout |
| 0x007E0F54 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E0FBA | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007E101C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007E107E | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E1114 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007E1235 | `Photos_Screen` | Known | Screen layout |
| 0x007E12A0 | `Photos_Screen` | Known | Screen layout |
| 0x007E13EC | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E1452 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007E14B4 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007E1516 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E15AC | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007E16CD | `Photos_Screen` | Known | Screen layout |
| 0x007E1738 | `Photos_Screen` | Known | Screen layout |
| 0x007E1884 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E18EA | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007E194C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007E19AE | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E1A44 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007E1B65 | `Photos_Screen` | Known | Screen layout |
| 0x007E1BD0 | `Photos_Screen` | Known | Screen layout |
| 0x007E1D1C | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E1D82 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007E1DE4 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007E1E46 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E1EDC | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007E1FFD | `Photos_Screen` | Known | Screen layout |
| 0x007E2068 | `Photos_Screen` | Known | Screen layout |
| 0x007E21B4 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E221A | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007E227C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007E22DE | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E2374 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007E2495 | `Photos_Screen` | Known | Screen layout |
| 0x007E2689 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007E26EB | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007E2759 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007E27BF | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007E2824 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007E2AF2 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007E2B54 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007E2BC2 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007E2C28 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007E2F2E | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007E2F90 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007E2FFE | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007E3064 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007E330D | `Radio_Screen_Default` | Known | Screen layout |
| 0x007E336A | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007E33CC | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007E343A | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007E34A0 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007E379A | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007E3804 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007E3A72 | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007E3ADC | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007E3C99 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007E3CFC | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007E3D61 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007E3DC9 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007E3E2C | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E3E94 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007E3EFD | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E3F63 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007E3FC8 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E4035 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007E40A5 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007E411B | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007E4191 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007E4201 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E4276 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007E42ED | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007E4361 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007E43D3 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007E444D | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E44C0 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007E4532 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E45B6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E45E0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E4667 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E46F4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E4793 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E47AD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E4825 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E483F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E48A9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E48C6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E493E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E4968 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E49EF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E4A7C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E4B1B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E4B35 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E4BAD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E4BC7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E4C31 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E4C4E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E4CC6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E4CF0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E4D77 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E4E04 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E4EA3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E4EBD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E4F35 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E4F4F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E4FB9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E4FD6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E504E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E5078 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E50FF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E518C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E522B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E5245 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E52BD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E52D7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E5341 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E535E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E53D6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E5400 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E5487 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E5514 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E55B3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E55CD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E5645 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E565F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E56C9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E56E6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E575E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E5788 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E580F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E589C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E593B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E5955 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E59CD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E59E7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E5A51 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E5A6E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E5AE6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E5B10 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E5B97 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E5C24 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E5CC3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E5CDD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E5D55 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E5D6F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E5DD9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E5DF6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E5E6E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E5E98 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E5F1F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E5FAC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E604B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E6065 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E60DD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E60F7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E6161 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E617E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E61F6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E6220 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E62A7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E6334 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E63D3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E63ED | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E6465 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E647F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E64E9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E6506 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E657E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E65A8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E662F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E66BC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E675B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E6775 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E67ED | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E6807 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E6871 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E688E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E6906 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E6930 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E69B7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E6A44 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E6AE3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E6AFD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E6B75 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E6B8F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E6BF9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E6C16 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E6C8E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E6CB8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E6D3F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E6DCC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E6E6B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E6E85 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E6EFD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E6F17 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E6F81 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E6F9E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E7016 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E7040 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E70C7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E7154 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E71F3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E720D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E7285 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E729F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E7309 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E7326 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E739E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E73C8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E744F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E74DC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E757B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E7595 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E760D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E7627 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E7691 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E76AE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E7726 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E7750 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E77D7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E7864 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E7903 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E791D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E7995 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E79AF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E7A19 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E7A36 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E7AAE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E7AD8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E7B5F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E7BEC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E7C8B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E7CA5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E7D1D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E7D37 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E7DA1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E7DBE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E7E36 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E7E60 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E7EE7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E7F74 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E8013 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E802D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E80A5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E80BF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E8129 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E8146 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E81BE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E81E8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E826F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E82FC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E839B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E83B5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E842D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E8447 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E84B1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E84CE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E8546 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E8570 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E85F7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E8684 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E8723 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E873D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E87B5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E87CF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E8839 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E8856 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E88DD | `SettingsMenus_ResetMainMenu_Screen9` | Known | Screen layout |
| 0x007E89AD | `SettingsMenus_ResetMusicMenu_Screen:` | Known | Screen layout |
| 0x007E8A61 | `SettingsMenus_Main_Screen,` | Known | Screen layout |
| 0x007E8AD3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E8AED | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E8B65 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E8B7F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E8EBA | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007E8F20 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007E8F7D | `Extras_Screen` | Known | Screen layout |
| 0x007E8FD1 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x007E90AF | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x007E911D | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007E91BB | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x007E91D4 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x007E923C | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007E92AF | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E9331 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007E9392 | `NikePlus_Alert_Screen$` | Known | Screen layout |
| 0x007E9412 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E948B | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E9505 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E958A | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E95AB | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E961A | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x007E96A2 | `NikePlus_Remote_Unlinking_Screen(` | Known | Screen layout |
| 0x007E96C6 | `NikePlus_Remote_Unlinking_Screen_Default` | Known | Screen layout |
| 0x007E973A | `NikePlus_Alert_Screen(` | Known | Screen layout |
| 0x007E97C6 | `NikePlus_Dynamic_Workout_Screen,` | Known | Screen layout |
| 0x007E97E9 | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x007E9862 | `NikePlus_Dynamic_Workout_Screen/` | Known | Screen layout |
| 0x007E9885 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x007E98FE | `NikePlus_Dynamic_Workout_Screen.` | Known | Screen layout |
| 0x007E9921 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x007E999A | `NikePlus_Alert_Screen(` | Known | Screen layout |
| 0x007E9A1D | `NikePlus_Custom_Screen,` | Known | Screen layout |
| 0x007E9A37 | `NikePlus_Custom_Screen_Weight_ToWorkoutMusic` | Known | Screen layout |
| 0x007E9AB1 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E9B33 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007E9BAB | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007E9BC9 | `NikePlus_EndWorkout_Screen_Default` | Known | Screen layout |
| 0x007E9C61 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007E9CDD | `NikePlus_EquipmentAlert_Screen<` | Known | Screen layout |
| 0x007E9DAA | `NikePlus_EquipmentAlert_Screen<` | Known | Screen layout |
| 0x007E9E74 | `NikePlus_EquipmentAlert_Screen8` | Known | Screen layout |
| 0x007E9F41 | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007EA00E | `NikePlus_EquipmentAlert_Screen1` | Known | Screen layout |
| 0x007EA0C0 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007EA0E1 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007EA178 | `NikePlus_Dynamic_Workout_Screen.` | Known | Screen layout |
| 0x007EA19B | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x007EA23B | `NikePlus_Dynamic_Workout_Screen/` | Known | Screen layout |
| 0x007EA25E | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x007EA2FC | `NikePlus_Dynamic_Workout_Screen,` | Known | Screen layout |
| 0x007EA31F | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x007EA3B5 | `NikePlus_EndPausedWorkout_Screen1` | Known | Screen layout |
| 0x007EA3D9 | `NikePlus_EndPausedWorkout_Screen_BasicSave_Layout` | Known | Screen layout |
| 0x007EA477 | `NikePlus_EndPausedWorkout_Screen4` | Known | Screen layout |
| 0x007EA49B | `NikePlus_EndPausedWorkout_Screen_CaloriesSave_Layout` | Known | Screen layout |
| 0x007EA53C | `NikePlus_EndPausedWorkout_Screen4` | Known | Screen layout |
| 0x007EA560 | `NikePlus_EndPausedWorkout_Screen_DistanceSave_Layout` | Known | Screen layout |
| 0x007EA5FE | `NikePlus_EndPausedWorkout_Screen0` | Known | Screen layout |
| 0x007EA622 | `NikePlus_EndPausedWorkout_Screen_TimeSave_Layout` | Known | Screen layout |
| 0x007EA6B7 | `NikePlus_Songs_Screen` | Known | Screen layout |
| 0x007EA6D0 | `NikePlus_Songs_Screen_Default` | Known | Screen layout |
| 0x007EA7E2 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007EA7FC | `NikePlus_Custom_Screen_Weight` | Known | Screen layout |
| 0x007EA85F | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007EA8D3 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007EA951 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x007EA9BB | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x007EAA1B | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007EAA98 | `NikePlus_SensorSearching_Screen'` | Known | Screen layout |
| 0x007EAABB | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x007EAB2A | `NikePlus_Playlists_Screen ` | Known | Screen layout |
| 0x007EAB47 | `NikePlus_Playlists_Screen_Nested` | Known | Screen layout |
| 0x007EABDB | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x007EAC3B | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007EACB8 | `NikePlus_SensorSearching_Screen'` | Known | Screen layout |
| 0x007EACDB | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x007EAD4A | `NikePlus_Playlists_Screen ` | Known | Screen layout |
| 0x007EAD67 | `NikePlus_Playlists_Screen_Nested` | Known | Screen layout |
| 0x007EAE30 | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x007EAE90 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007EAF0D | `NikePlus_Playlists_Screen!` | Known | Screen layout |
| 0x007EAF2A | `NikePlus_Playlists_Screen_Default` | Known | Screen layout |
| 0x007EAF96 | `NikePlus_SensorSearching_Screen'` | Known | Screen layout |
| 0x007EAFB9 | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x007EB155 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007EB173 | `NikePlus_NowRunning_Screen_Basic'` | Known | Screen layout |
| 0x007EB1E7 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EB205 | `NikePlus_NowRunning_Screen_Calories'` | Known | Screen layout |
| 0x007EB27C | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EB29A | `NikePlus_NowRunning_Screen_Distance#` | Known | Screen layout |
| 0x007EB30D | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EB32B | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007EB393 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007EB46F | `NikePlus_NowRunning_Screen$` | Known | Screen layout |
| 0x007EB48D | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x007EB557 | `NikePlus_NowRunning_Screen$` | Known | Screen layout |
| 0x007EB575 | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x007EB63F | `NikePlus_NowRunning_Screen$` | Known | Screen layout |
| 0x007EB65D | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x007EB922 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EB94E | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EB9D0 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EB9FE | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EBA80 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EBAA2 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EBB14 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EBB37 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EBBA7 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EBBC5 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EBC35 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EBC5B | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EBCCF | `NikePlus_NowRunning_Screen_Basic_PowerSong` | Known | Screen layout |
| 0x007EC05C | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EC088 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EC10A | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EC138 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EC1BA | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EC1DC | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EC24E | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EC271 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EC2E1 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EC2FF | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EC36F | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EC395 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EC406 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007EC786 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EC7B2 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EC834 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EC862 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EC8E4 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EC906 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EC978 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EC99B | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007ECA0B | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007ECA29 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007ECA99 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007ECABF | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007ECB33 | `NikePlus_NowRunning_Screen_Time_PowerSong` | Known | Screen layout |
| 0x007ECEC0 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007ECEEC | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007ECF6E | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007ECF9C | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007ED01E | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007ED040 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007ED0B2 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007ED0D5 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007ED145 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007ED163 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007ED1D3 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007ED1F9 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007ED26A | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007ED5EA | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007ED616 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007ED698 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007ED6C6 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007ED748 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007ED76A | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007ED7DC | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007ED7FF | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007ED86F | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007ED88D | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007ED8FD | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007ED923 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007ED997 | `NikePlus_NowRunning_Screen_Distance_PowerSong` | Known | Screen layout |
| 0x007EDD28 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EDD54 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EDDD6 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EDE04 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EDE86 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EDEA8 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EDF1A | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EDF3D | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EDFAD | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EDFCB | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EE03B | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EE061 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EE0D2 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007EE456 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EE482 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EE504 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EE532 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EE5B4 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EE5D6 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EE648 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EE66B | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EE6DB | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EE6F9 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EE769 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EE78F | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EE803 | `NikePlus_NowRunning_Screen_Calories_PowerSong` | Known | Screen layout |
| 0x007EEB94 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EEBC0 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EEC42 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EEC70 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EECF2 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EED14 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EED86 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EEDA9 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EEE19 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EEE37 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EEEA7 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EEECD | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EEF3E | `NikePlus_NowRunning_Screen_Calories` | Known | Screen layout |
| 0x007EF28C | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EF2B8 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EF33A | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EF368 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EF3EA | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EF40C | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EF47E | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EF4A1 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EF511 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EF52F | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EF59F | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EF5C5 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EF760 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EF7C7 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EF83B | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EF8AE | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007EF91A | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007EF93B | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007EF9B6 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EF9DC | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EFA99 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EFAC5 | `NikePlus_CalibrationCompleteError_Screen_Default'` | Known | Screen layout |
| 0x007EFB49 | `NikePlus_CalibrationCompleteError_Screen*` | Known | Screen layout |
| 0x007EFB75 | `NikePlus_CalibrationComplete_Screen_Pacing%` | Known | Screen layout |
| 0x007EFBF1 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EFC1F | `NikePlus_CalibrationCompleteSuccess_Screen_Default` | Known | Screen layout |
| 0x007EFC98 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EFD28 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007EFD7B | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x007EFDE8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007EFE3B | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x007EFEA8 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007EFEFC | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007EFF62 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007EFF80 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007EFFEC | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F000A | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007F007A | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F0098 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007F0104 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007F0122 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007F01CF | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007F01F5 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007F0288 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007F02A2 | `NikePlus_Custom_Screen_Time` | Known | Screen layout |
| 0x007F0323 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F0344 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F03D7 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007F03F1 | `NikePlus_Custom_Screen_Distance` | Known | Screen layout |
| 0x007F0479 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F049A | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F0517 | `NikePlus_EquipmentAlert_Screen)` | Known | Screen layout |
| 0x007F05B0 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F05D1 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F065C | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007F0676 | `NikePlus_Custom_Screen_Calories` | Known | Screen layout |
| 0x007F0729 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F07B0 | `NikePlus_Alert_Screen(` | Known | Screen layout |
| 0x007F0859 | `NikePlus_EquipmentAlert_ScreenK` | Known | Screen layout |
| 0x007F091C | `NikePlus_EquipmentAlert_Screen8` | Known | Screen layout |
| 0x007F09D0 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007F0AC0 | `NikePlus_EquipmentAlert_Screen>` | Known | Screen layout |
| 0x007F0B7E | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007F0C3D | `NikePlus_Alert_Screen$` | Known | Screen layout |
| 0x007F0CBE | `NikePlus_Remote_Unlinking_Screen(` | Known | Screen layout |
| 0x007F0CE2 | `NikePlus_Remote_Unlinking_Screen_Default!` | Known | Screen layout |
| 0x007F0D58 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x007F0E96 | `NikePlus_Calibration_CalibrateWalk_Screen1` | Known | Screen layout |
| 0x007F0F3A | `NikePlus_Calibration_CalibrateRun_Screen0` | Known | Screen layout |
| 0x007F0FFD | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007F10BE | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F10DF | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F1166 | `NikePlus_Custom_Screen*` | Known | Screen layout |
| 0x007F1180 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x007F1271 | `NikePlus_Calibrate_ResetToDefault_Screen0` | Known | Screen layout |
| 0x007F129D | `NikePlus_Calibrate_ResetToDefault_Screen_Walking` | Known | Screen layout |
| 0x007F1353 | `NikePlus_Calibration_CalibrateRun_Screen"` | Known | Screen layout |
| 0x007F13CB | `NikePlus_Calibration_CalibrateWalk_Screen` | Known | Screen layout |
| 0x007F1477 | `NikePlus_Calibration_CalibrateRun_Screen"` | Known | Screen layout |
| 0x007F14EF | `NikePlus_Calibration_CalibrateWalk_Screen` | Known | Screen layout |
| 0x007F157D | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007F163E | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F165F | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F16E6 | `NikePlus_Custom_Screen*` | Known | Screen layout |
| 0x007F1700 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x007F17EF | `NikePlus_Calibrate_ResetToDefault_Screen0` | Known | Screen layout |
| 0x007F181B | `NikePlus_Calibrate_ResetToDefault_Screen_Running` | Known | Screen layout |
| 0x007F188D | `NikePlus_StartWorkout_Screen$` | Known | Screen layout |
| 0x007F18AD | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x007F1914 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007F1967 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007F19DC | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007F1A36 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007F1AED | `NikePlus_Custom_Screen!` | Known | Screen layout |
| 0x007F1B69 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007F1BE0 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007F1C72 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x007F1CF0 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007F1D4A | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007F1DE9 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007F1E6A | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007F1EC1 | `NikePlus_Calibration_ChooseCalibration_Screen5` | Known | Screen layout |
| 0x007F1F6E | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F1FED | `NikePlus_StartCalibration_Screen(` | Known | Screen layout |
| 0x007F2011 | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x007F2077 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F20F3 | `NikePlus_StartWorkout_Screen$` | Known | Screen layout |
| 0x007F2113 | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x007F217E | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F21F7 | `NikePlus_Workout_Music_Screen` | Known | Screen layout |
| 0x007F225E | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007F22B9 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F2365 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F23F7 | `NikePlus_StartWorkout_Screen$` | Known | Screen layout |
| 0x007F2417 | `NikePlus_StartWorkout_Screen_Default#` | Known | Screen layout |
| 0x007F248B | `NikePlus_StartCalibration_Screen(` | Known | Screen layout |
| 0x007F24AF | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x007F251C | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F25A4 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x007F2633 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F2654 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F26F1 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F2712 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F27B1 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F27D2 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F286D | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F288E | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F295C | `NikePlus_EquipmentAlert_Screen)` | Known | Screen layout |
| 0x007F29F5 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F2A16 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F2AB2 | `NikePlus_History_BestWorkouts_Screen,` | Known | Screen layout |
| 0x007F2ADA | `NikePlus_History_BestWorkouts_Screen_Default#` | Known | Screen layout |
| 0x007F2B56 | `NikePlus_History_RecentWorkouts_Screen.` | Known | Screen layout |
| 0x007F2B80 | `NikePlus_History_RecentWorkouts_Screen_Default'` | Known | Screen layout |
| 0x007F2C02 | `NikePlus_History_WorkoutSummary_Screen+` | Known | Screen layout |
| 0x007F2C2C | `NikePlus_History_WorkoutSummary_Screen_Last1` | Known | Screen layout |
| 0x007F2CB5 | `NikePlus_NoData_Screen%` | Known | Screen layout |
| 0x007F2CCF | `NikePlus_NoData_Screen_NoBestWorkouts2` | Known | Screen layout |
| 0x007F2D53 | `NikePlus_NoData_Screen&` | Known | Screen layout |
| 0x007F2D6D | `NikePlus_NoData_Screen_NoSavedWorkouts` | Known | Screen layout |
| 0x007F2E7D | `NikePlus_History_Totals_Screen&` | Known | Screen layout |
| 0x007F2E9F | `NikePlus_History_Totals_Screen_Default` | Known | Screen layout |
| 0x007F2F11 | `NikePlus_History_DeleteActiveWorkout_Screen2` | Known | Screen layout |
| 0x007F2F40 | `NikePlus_History_DeleteActiveWorkout_Screen_Cancel` | Known | Screen layout |
| 0x007F2FB5 | `NikePlus_History_DeleteActiveWorkout_Screen7` | Known | Screen layout |
| 0x007F2FE4 | `NikePlus_History_DeleteActiveWorkout_Screen_LastWorkout` | Known | Screen layout |
| 0x007F305C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007F30AF | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007F3105 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007F31C0 | `NikePlus_History_RecentWorkouts_ScreenB` | Known | Screen layout |
| 0x007F3256 | `NikePlus_History_WorkoutSummary_ScreenA` | Known | Screen layout |
| 0x007F32EB | `NikePlus_History_Screen` | Known | Screen layout |
| 0x007F33B7 | `NikePlus_History_RecentWorkouts_ScreenB` | Known | Screen layout |
| 0x007F344D | `NikePlus_History_WorkoutSummary_ScreenA` | Known | Screen layout |
| 0x007F34E2 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x007F359F | `NikePlus_History_ScreenG` | Known | Screen layout |
| 0x007F362B | `NikePlus_History_RecentWorkouts_Screen` | Known | Screen layout |
| 0x007F36AB | `NikePlus_History_DeleteAllWorkouts_Screen0` | Known | Screen layout |
| 0x007F36D8 | `NikePlus_History_DeleteAllWorkouts_Screen_Cancel#` | Known | Screen layout |
| 0x007F3758 | `NikePlus_History_WorkoutSummary_Screen.` | Known | Screen layout |
| 0x007F3782 | `NikePlus_History_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007F3835 | `NikePlus_History_ClearTotals_Screen+` | Known | Screen layout |
| 0x007F385C | `NikePlus_History_ClearTotals_Screen_Default` | Known | Screen layout |
| 0x007F38FE | `NikePlus_History_Totals_Screen` | Known | Screen layout |
| 0x007F3991 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F39B2 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F3A21 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007F3A3F | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007F3AAB | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F3AC9 | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007F3B39 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F3B57 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007F3BC3 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007F3BE1 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007F3C77 | `NikePlus_Dynamic_Workout_Screen/` | Known | Screen layout |
| 0x007F3C9A | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x007F3D13 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007F3D31 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007F3D9D | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F3DBB | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007F3E2B | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F3E49 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007F3EB5 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007F3ED3 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007F3F6B | `NikePlus_Dynamic_Workout_Screen,` | Known | Screen layout |
| 0x007F3F8E | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x007F4004 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007F4022 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007F408E | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F40AC | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007F411C | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F413A | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007F41A6 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007F41C4 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007F425B | `NikePlus_Dynamic_Workout_Screen.` | Known | Screen layout |
| 0x007F427E | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x007F42F6 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007F4314 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007F4380 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F439E | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007F440E | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F442C | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007F4498 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007F44B6 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007F4526 | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x007F453F | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x007F45A2 | `DemoMode_Screen` | Known | Screen layout |
| 0x007F45B5 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x007F4622 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x007F463B | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x007F46AE | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x007F46C9 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x007F47D9 | `VoiceMemos_Context_PlayDelete_Screen)` | Known | Screen layout |
| 0x007F4801 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x007F4878 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007F4944 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007F49B3 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007F4AA1 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007F4B0A | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007F4B2C | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007F4B98 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007F4BBA | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007F4D36 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007F4D52 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007F4E19 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007F4E34 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007F4E97 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007F4EFA | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007F4F91 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007F4FAD | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007F5074 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007F508F | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007F50F2 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007F5155 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007F51ED | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007F5209 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007F52D0 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007F52EB | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007F534E | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007F53B1 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007F542E | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007F5499 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007F5505 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007F5577 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007F55E4 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007F564F | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x007F56BB | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007F5723 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007F578F | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007F5803 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007F5871 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x007F58EA | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x00818484 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x00818509 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x008187F6 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x009CD356 | `VoiceMemos_Screen_InsufficientDiskSpace2` | Known | Screen layout |
| 0x009CECF6 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x009CED0E | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x009CED2C | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x009CEDF3 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x009CEE71 | `NowPlaying_Screen_Music` | Known | Screen layout |
| 0x009CEEB2 | `MainMenu_Music_Screen_NoMusic` | Known | Screen layout |
| 0x009CEED0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x009CEEEE | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x009CEF07 | `NikePlus_Custom_Screen_Weight_ToWorkoutMusic` | Known | Screen layout |
| 0x009CF01C | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x009CF0D0 | `DeleteRental_Confirmation_Screen_Unwatched` | Known | Screen layout |
| 0x009CF126 | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x009CF172 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x009CF3D5 | `NowPlaying_Screen_ExtraInfoLoadFailed` | Known | Screen layout |
| 0x009CF430 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x009CF449 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x009CF467 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x009CF496 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x009CF4CE | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x009CF55E | `NikePlus_Playlists_Screen_Nested` | Known | Screen layout |
| 0x009CF93E | `Video_Settings_TV_Screen_Standard` | Known | Screen layout |
| 0x009CF970 | `Search_Main_Screen_WithKeyboard` | Known | Screen layout |
| 0x009CF990 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x009CF9D5 | `VoiceMemos_Screen_InsufficientDiskSpace` | Known | Screen layout |
| 0x009CFA3A | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x009CFA5E | `NikePlus_Custom_Screen_Distance` | Known | Screen layout |
| 0x009CFAB9 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x009CFB40 | `Video_Settings_TV_Screen_Wide` | Known | Screen layout |
| 0x009CFB88 | `CoverFlow_Screen_Backside` | Known | Screen layout |
| 0x009D31CA | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x009D33CF | `NowPlaying_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x009D33F4 | `Slideshow_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x009D34C4 | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x009D34DE | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x009D35D6 | `RentalDeleted_Screen_Title` | Known | Screen layout |
| 0x009D35F1 | `SingleRentalExpiring_Screen_Title` | Known | Screen layout |
| 0x009D3613 | `MultipleRentalsExpiring_Screen_Title` | Known | Screen layout |
| 0x009D3638 | `DeleteRental_Screen_Title` | Known | Screen layout |
| 0x009D36DB | `VoiceMemos_DeleteAll_ScreenTitle` | Known | Screen layout |
| 0x009D3778 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x009D37BB | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x009D38AD | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x009D38CD | `NikePlus_Custom_Screen_Time` | Known | Screen layout |
| 0x009D3A18 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x009D3B01 | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x009D3B1A | `Radio_Screen_Volume` | Known | Screen layout |
| 0x009D3B2E | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x009D3B4B | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x009D3B6A | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x009D3C75 | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x009D3DE1 | `CoverFlow_Screen_BacksideForCapture` | Known | Screen layout |
| 0x009D4FB6 | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x009D50AE | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x009D50C9 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x009D536C | `NikePlus_CalibrationComplete_Screen_Pacing` | Known | Screen layout |
| 0x009D5404 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x009D5438 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x009D5475 | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x009D5587 | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x009D56B5 | `NikePlus_Calibrate_ResetToDefault_Screen_Walking` | Known | Screen layout |
| 0x009D57E7 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x009D5800 | `NikePlus_Calibrate_ResetToDefault_Screen_Running` | Known | Screen layout |
| 0x009D5850 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x009D5876 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x009DB509 | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x009DB578 | `TVOutSlideshow_Screen_Playing` | Known | Screen layout |
| 0x009DB596 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x009DB5EC | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x009DB656 | `NikePlus_NowRunning_Screen_Basic_PowerSong` | Known | Screen layout |
| 0x009DB681 | `NikePlus_NowRunning_Screen_Distance_PowerSong` | Known | Screen layout |
| 0x009DB6AF | `NikePlus_NowRunning_Screen_Time_PowerSong` | Known | Screen layout |
| 0x009DB6FC | `NikePlus_NowRunning_Screen_Calories_PowerSong` | Known | Screen layout |
| 0x009DB779 | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x009DB7E4 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x009DB864 | `Extras_Screen_Debug` | Known | Screen layout |
| 0x009DB96E | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x009DB98E | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x009DBF00 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x009DBF1B | `Extras_Screen_Lock` | Known | Screen layout |
| 0x009DBF2E | `MainMenu_List_ScreenLock` | Known | Screen layout |
| 0x009DBF47 | `ExtrasMenu_ScreenLock` | Known | Screen layout |
| 0x009DBFCA | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x009DBFEB | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x009DC096 | `NikePlus_StartCalibration_Screen_Walk` | Known | Screen layout |
| 0x009DC11E | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x009DC140 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x009DC247 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x009DC287 | `VoiceMemos_Screen_DeletAllAsk` | Known | Screen layout |
| 0x009DC2A5 | `NowPlaying_Screen_TVOutAsk` | Known | Screen layout |
| 0x009DC401 | `NowPlaying_Screen_Initial` | Known | Screen layout |
| 0x009DC41B | `NowPlaying_Screen_Video_TVOut_Initial` | Known | Screen layout |
| 0x009DC6ED | `NikePlus_History_DeleteAllWorkouts_Screen_Cancel` | Known | Screen layout |
| 0x009DC71E | `NikePlus_History_DeleteActiveWorkout_Screen_Cancel` | Known | Screen layout |
| 0x009DD540 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x009DD5C1 | `RemoteUI_Screen` | Known | Screen layout |
| 0x009DD5D1 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x009DD5E9 | `Video_Settings_TV_Screen` | Known | Screen layout |
| 0x009DD602 | `NikePlus_NoData_Screen` | Known | Screen layout |
| 0x009DD619 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x009DD630 | `NikePlus_Workout_Music_Screen` | Known | Screen layout |
| 0x009DD64E | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x009DD672 | `AddressViewer_PartialLoad_Screen` | Known | Screen layout |
| 0x009DD693 | `NikePlus_ActivityStopped_Screen` | Known | Screen layout |
| 0x009DD6B3 | `Alarms_Alarm_Clock_Triggered_Screen` | Known | Screen layout |
| 0x009DD6D7 | `Alarms_Alarm_Triggered_Screen` | Known | Screen layout |
| 0x009DD6F5 | `Unsupported_Screen` | Known | Screen layout |
| 0x009DD708 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x009DD726 | `LockediPod_Screen` | Known | Screen layout |
| 0x009DD738 | `DiskMode_Screen` | Known | Screen layout |
| 0x009DD748 | `DemoMode_Screen` | Known | Screen layout |
| 0x009DD758 | `Notes_Image_Screen` | Known | Screen layout |
| 0x009DD76B | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x009DD789 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x009DD79F | `NowPlaying_Idle_Screen` | Known | Screen layout |
| 0x009DD7B6 | `Game_Screen` | Known | Screen layout |
| 0x009DD7C2 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x009DD7DF | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x009DD7F8 | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x009DD819 | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x009DD83E | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x009DD851 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x009DD86E | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x009DD88F | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x009DD8B4 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x009DD8CB | `Notes_Loading_Screen` | Known | Screen layout |
| 0x009DD8E0 | `NikePlus_SensorSearching_Screen` | Known | Screen layout |
| 0x009DD900 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x009DD91F | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x009DD937 | `NikePlus_Remote_Unlinking_Screen` | Known | Screen layout |
| 0x009DD958 | `SettingsMenus_AdjustScrolling_Screen` | Known | Screen layout |
| 0x009DD97D | `Game_Running_Screen` | Known | Screen layout |
| 0x009DD991 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x009DD9AC | `Stopwatch_Screen` | Known | Screen layout |
| 0x009DD9BD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x009DD9D4 | `Clock_Screen` | Known | Screen layout |
| 0x009DD9E1 | `NikePlus_Calibration_CalibrateWalk_Screen` | Known | Screen layout |
| 0x009DDA0B | `TVOutSlideshowAsk_Screen` | Known | Screen layout |
| 0x009DDA24 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x009DDA3A | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x009DDA58 | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x009DDA74 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x009DDA85 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x009DDA9C | `DemoMode_Main_Screen` | Known | Screen layout |
| 0x009DDAB1 | `Search_Main_Screen` | Known | Screen layout |
| 0x009DDAC4 | `AddressViewer_Main_Screen` | Known | Screen layout |
| 0x009DDADE | `Speakers_Main_Screen` | Known | Screen layout |
| 0x009DDAF3 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x009DDB09 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x009DDB23 | `Clock_Region_Screen` | Known | Screen layout |
| 0x009DDB37 | `RentalDeleted_Notification_Screen` | Known | Screen layout |
| 0x009DDB59 | `SingleRentalExpiring_Notification_Screen` | Known | Screen layout |
| 0x009DDB82 | `MultipleRentalsExpiring_Notification_Screen` | Known | Screen layout |
| 0x009DDBAE | `RentalError_Notification_Screen` | Known | Screen layout |
| 0x009DDBCE | `DeleteRental_Confirmation_Screen` | Known | Screen layout |
| 0x009DDBEF | `LockConfirmation_Screen` | Known | Screen layout |
| 0x009DDC07 | `NikePlus_EndCalibration_Screen` | Known | Screen layout |
| 0x009DDC26 | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x009DDC54 | `NikePlus_StartCalibration_Screen` | Known | Screen layout |
| 0x009DDC75 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x009DDC93 | `NikePlus_Calibration_CalibrateRun_Screen` | Known | Screen layout |
| 0x009DDCBC | `Video_Settings_Fit_To_Screen` | Known | Screen layout |
| 0x009DDCD9 | `RentalInfo_Screen` | Known | Screen layout |
| 0x009DDCEB | `Radio_Screen` | Known | Screen layout |
| 0x009DDCF8 | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x009DDD12 | `ChargingMode_LowPower_Screen` | Known | Screen layout |
| 0x009DDD2F | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x009DDD49 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x009DDD63 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x009DDD7D | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x009DDD96 | `NikePlus_CalibrationCompleteError_Screen` | Known | Screen layout |
| 0x009DDDBF | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x009DDDD6 | `Extras_Screen` | Known | Screen layout |
| 0x009DDDE4 | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x009DDE01 | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x009DDE23 | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x009DDE3C | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x009DDE5A | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x009DDE73 | `Video_Settings_Screen` | Known | Screen layout |
| 0x009DDE89 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x009DDEA2 | `SettingsMenus_AudiobookSettings_Screen` | Known | Screen layout |
| 0x009DDEC9 | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x009DDEEF | `PhotosSettings_Screen` | Known | Screen layout |
| 0x009DDF05 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x009DDF1D | `NikePlus_Songs_Screen` | Known | Screen layout |
| 0x009DDF33 | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x009DDF56 | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x009DDF73 | `MediaLists_Rentals_Screen` | Known | Screen layout |
| 0x009DDF8D | `NikePlus_History_Totals_Screen` | Known | Screen layout |
| 0x009DDFAC | `NikePlus_History_ClearTotals_Screen` | Known | Screen layout |
| 0x009DDFD0 | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x009DDFF4 | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x009DE00D | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x009DE02F | `ToDo_Instructions_Screen` | Known | Screen layout |
| 0x009DE048 | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x009DE064 | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x009DE07E | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x009DE09F | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x009DE0BB | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x009DE0D3 | `VoiceMemos_Screen` | Known | Screen layout |
| 0x009DE0E5 | `No_Photos_Screen` | Known | Screen layout |
| 0x009DE0F6 | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x009DE110 | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x009DE12C | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x009DE150 | `NikePlus_CalibrationCompleteSuccess_Screen` | Known | Screen layout |
| 0x009DE17B | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x009DE19B | `ChargingMode_Progress_Screen` | Known | Screen layout |
| 0x009DE1B8 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x009DE1CE | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x009DE1E9 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x009DE205 | `NikePlus_Playlists_Screen` | Known | Screen layout |
| 0x009DE21F | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x009DE241 | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x009DE262 | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x009DE27C | `NikePlus_History_DeleteAllWorkouts_Screen` | Known | Screen layout |
| 0x009DE2A6 | `NikePlus_History_RecentWorkouts_Screen` | Known | Screen layout |
| 0x009DE2CD | `NikePlus_History_BestWorkouts_Screen` | Known | Screen layout |
| 0x009DE2F2 | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x009DE30C | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x009DE32B | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x009DE34C | `NikePlus_Calibrate_ResetToDefault_Screen` | Known | Screen layout |
| 0x009DE375 | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x009DE38D | `NoContent_Screen` | Known | Screen layout |
| 0x009DE39E | `Calendar_Event_Screen` | Known | Screen layout |
| 0x009DE3B4 | `FirstBoot_Screen` | Known | Screen layout |
| 0x009DE3C5 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x009DE3DB | `NikePlus_EquipmentAlert_Screen` | Known | Screen layout |
| 0x009DE3FA | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x009DE410 | `Notes_List_Screen` | Known | Screen layout |
| 0x009DE422 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x009DE438 | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x009DE459 | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x009DE473 | `NikePlus_Dynamic_Workout_Screen` | Known | Screen layout |
| 0x009DE493 | `NikePlus_EndPausedWorkout_Screen` | Known | Screen layout |
| 0x009DE4B4 | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x009DE4CF | `NikePlus_ResumeWorkout_Screen` | Known | Screen layout |
| 0x009DE4ED | `NikePlus_History_DeleteActiveWorkout_Screen` | Known | Screen layout |
| 0x009DE519 | `NikePlus_StartWorkout_Screen` | Known | Screen layout |
| 0x009DE536 | `TODOS_Menu_Screen` | Known | Screen layout |
| 0x009DE548 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x009DE55E | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x009DE57A | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x009DE58F | `Games_Menu_Screen` | Known | Screen layout |
| 0x009DE5A1 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x009DE5B4 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x009DE5D3 | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x009DE5F2 | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x009DE616 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x009DE62C | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x009DE64A | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x009DE66D | `TVOutSlideshow_Screen` | Known | Screen layout |
| 0x009DE683 | `CoverFlow_Screen` | Known | Screen layout |
| 0x009DE694 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x009DE6A8 | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x009DE6CA | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x009DE6E2 | `Stopwatch_SessionSummary_Screen` | Known | Screen layout |
| 0x009DE702 | `NikePlus_End_WorkoutSummary_Screen` | Known | Screen layout |
| 0x009DE725 | `NikePlus_History_WorkoutSummary_Screen` | Known | Screen layout |
| 0x009DE74C | `SettingsMenu_AboutWithAccessory_Screen` | Known | Screen layout |
| 0x009DE773 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x009DE78B | `Clock_NorthAmerica_City_Screen` | Known | Screen layout |
| 0x009DE7AA | `Clock_SouthAmerica_City_Screen` | Known | Screen layout |
| 0x009DE7C9 | `Clock_Africa_City_Screen` | Known | Screen layout |
| 0x009DE7E2 | `Clock_Australia_City_Screen` | Known | Screen layout |
| 0x009DE7FE | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x009DE815 | `Clock_Pacific_City_Screen` | Known | Screen layout |
| 0x009DE82F | `Clock_Atlantic_City_Screen` | Known | Screen layout |
| 0x009DE84A | `Clock_Europe_City_Screen` | Known | Screen layout |
| 0x009DE93E | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x009DE98F | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x009DE9B2 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x009DE9DA | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x009DED66 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x009DEE69 | `RentalInfo_Screen_ExpiringSoon` | Known | Screen layout |
| 0x009DEEBF | `RentalInfo_Screen_NoAlbumArt_ExpiringSoon` | Known | Screen layout |
| 0x009DEFC9 | `NikePlus_StartCalibration_Screen_Run` | Known | Screen layout |
| 0x009DF2B3 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x009DF309 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x009DF45A | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x009DF477 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x009DF8A8 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x009DF9CA | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x009DF9EC | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x009DFB24 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x009DFB43 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x009E0233 | `ChargingMode_ScreenLayout_LowPower` | Known | Screen layout |
| 0x009E0BF3 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x009E0D51 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x009E0E08 | `NikePlus_NowRunning_Screen_Calories` | Known | Screen layout |
| 0x009E0E2C | `NikePlus_Custom_Screen_Calories` | Known | Screen layout |
| 0x009E0EC5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x009E0EE3 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x009E0F03 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x009E100E | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x009E102A | `Extras_Screen_Games` | Known | Screen layout |
| 0x009E1130 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x009E114F | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x009E116B | `Extras_Screen_Notes` | Known | Screen layout |
| 0x009E1256 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x009E1372 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x009E1540 | `MainMenu_Music_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009E1563 | `MainMenus_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009E1586 | `NoContent_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009E15C0 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x009E15DF | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x009E1600 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x009E16F7 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x009E1714 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x009E1793 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x009E1877 | `MainMenu_Music_Screen_NoCompilations` | Known | Screen layout |
| 0x009E189C | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x009E1A45 | `MainMenu_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x009E1A68 | `MainMenu_Videos_Screen_NoMusicVideos` | Known | Screen layout |
| 0x009E1A8D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x009E1AAC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x009E1ACB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x009E1AEC | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x009E1B2A | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x009E1B4B | `Extras_Screen_VoiceMemos` | Known | Screen layout |
| 0x009E1BB6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x009E1BE8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x009E1C07 | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x009E1CB4 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x009E1D20 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x009E1E19 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x009E1E35 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x009E1EB8 | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x009E1ED3 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x009E1EF4 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x009E1FA3 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x009E1FD7 | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x009E1FF8 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x009E20B6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x009E20D7 | `MainMenu_Videos_Screen_NoPlaylists` | Known | Screen layout |
| 0x009E20FA | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x009E2149 | `MainMenu_Music_Screen_NoArtists` | Known | Screen layout |
| 0x009E21B9 | `NikePlus_NoData_Screen_NoSavedWorkouts` | Known | Screen layout |
| 0x009E226B | `NikePlus_NoData_Screen_NoBestWorkouts` | Known | Screen layout |
| 0x009E2318 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x009E2337 | `Extras_Screen_NikePlus` | Known | Screen layout |
| 0x009E2487 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x009E24A6 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x009E24C7 | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x009E299E | `NikePlus_Custom_Screen_Weight` | Known | Screen layout |
| 0x009E2A13 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x009E2AC6 | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x009E2B40 | `Notes_Contents_Screen_Alt` | Known | Screen layout |
| 0x009E2B5A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x009E2C06 | `VoiceMemos_Screen_InsufficientDiskSpace2_Default` | Known | Screen layout |
| 0x009E2CB8 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x009E2D5D | `VoiceMemos_Screen_InsufficientDiskSpace_Default` | Known | Screen layout |
| 0x009E2D8D | `NowPlaying_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x009E2DBA | `Slideshow_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x009E3C94 | `Radio_Screen_Tuning_Default` | Known | Screen layout |
| 0x009E3D20 | `VoiceMemos_Screen_DeletAllAsk_Default` | Known | Screen layout |
| 0x009E3D46 | `NowPlaying_Screen_TVOutAsk_Default` | Known | Screen layout |
| 0x009E3D7D | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x009E3DA3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x009E3DC1 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x009E3DED | `AddressViewer_PartialLoad_Screen_Default` | Known | Screen layout |
| 0x009E3E16 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x009E3E3E | `Alarms_Alarm_Clock_Triggered_Screen_Default` | Known | Screen layout |
| 0x009E3E6A | `Alarms_Alarm_Triggered_Screen_Default` | Known | Screen layout |
| 0x009E3E90 | `Unsupported_Screen_Default` | Known | Screen layout |
| 0x009E3EAB | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x009E3ED1 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x009E3EE9 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x009E3F04 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x009E3F21 | `Game_Screen_Default` | Known | Screen layout |
| 0x009E3F35 | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x009E3F5B | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x009E3F7C | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x009E3FA5 | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x009E3FCF | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x009E3FFC | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x009E4025 | `Notes_Loading_Screen_Default` | Known | Screen layout |
| 0x009E4042 | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x009E406A | `NikePlus_Remote_Unlinking_Screen_Default` | Known | Screen layout |
| 0x009E4093 | `Clock_Screen_Default` | Known | Screen layout |
| 0x009E40A8 | `TVOutSlideshowAsk_Screen_Default` | Known | Screen layout |
| 0x009E40C9 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x009E40E7 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x009E410D | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x009E4131 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x009E414A | `AddressViewer_Main_Screen_Default` | Known | Screen layout |
| 0x009E416C | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x009E4189 | `MainMenus_Main_Screen_Default` | Known | Screen layout |
| 0x009E41A7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x009E41C4 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x009E41E0 | `RentalDeleted_Notification_Screen_Default` | Known | Screen layout |
| 0x009E420A | `SingleRentalExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x009E423B | `MultipleRentalsExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x009E426F | `RentalError_Notification_Screen_Default` | Known | Screen layout |
| 0x009E4297 | `DeleteRental_Confirmation_Screen_Default` | Known | Screen layout |
| 0x009E42C0 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x009E42EC | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x009E4313 | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x009E433C | `RentalInfo_Screen_Default` | Known | Screen layout |
| 0x009E4356 | `Radio_Screen_Default` | Known | Screen layout |
| 0x009E436B | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x009E438D | `NikePlus_CalibrationCompleteError_Screen_Default` | Known | Screen layout |
| 0x009E43BE | `Extras_Screen_Default` | Known | Screen layout |
| 0x009E43D4 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x009E43FA | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x009E441B | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x009E4439 | `NikePlus_Settings_Screen_Default` | Known | Screen layout |
| 0x009E445A | `NikePlus_Songs_Screen_Default` | Known | Screen layout |
| 0x009E4478 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x009E449A | `NikePlus_History_Totals_Screen_Default` | Known | Screen layout |
| 0x009E44C1 | `NikePlus_History_ClearTotals_Screen_Default` | Known | Screen layout |
| 0x009E44ED | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x009E4519 | `ToDo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009E453A | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009E455E | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x009E4580 | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x009E45A4 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x009E45C3 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x009E45DC | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x009E45FE | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x009E4622 | `NikePlus_CalibrationCompleteSuccess_Screen_Default` | Known | Screen layout |
| 0x009E4655 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x009E4673 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x009E4697 | `NikePlus_Playlists_Screen_Default` | Known | Screen layout |
| 0x009E46B9 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x009E46E3 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x009E470C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x009E472E | `NikePlus_History_RecentWorkouts_Screen_Default` | Known | Screen layout |
| 0x009E475D | `NikePlus_History_BestWorkouts_Screen_Default` | Known | Screen layout |
| 0x009E478A | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x009E47AA | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x009E47C8 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x009E47E1 | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x009E47FF | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x009E4819 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x009E4837 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x009E4860 | `NikePlus_EndWorkout_Screen_Default` | Known | Screen layout |
| 0x009E4883 | `NikePlus_ResumeWorkout_Screen_Default` | Known | Screen layout |
| 0x009E48A9 | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x009E48CE | `TODOS_Menu_Screen_Default` | Known | Screen layout |
| 0x009E48E8 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x009E4906 | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x009E4923 | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x009E493D | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x009E4958 | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x009E4977 | `Debug_MainMenu_Screen_Default` | Known | Screen layout |
| 0x009E4995 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x009E49B3 | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x009E49CC | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x009E49E8 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x009E4A12 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x009E4A32 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x009E4A5A | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x009E4A85 | `NikePlus_History_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x009E4AB4 | `NikePlus_History_Screen_Default` | Known | Screen layout |
| 0x009E4AD4 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009E4AFB | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009E4B22 | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x009E4B43 | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x009E4B67 | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x009E4B86 | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x009E4BA8 | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x009E4BCB | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x009E4C08 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x009E4C96 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x009E4CC6 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x009E4CE8 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x009E4D59 | `RentalInfo_Screen_NoAlbumArt_Default` | Known | Screen layout |
| 0x009E4D7E | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x009E54AA | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009E54D6 | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009E551B | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x009E5543 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x009E5564 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x009E5585 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x009E55AB | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x009E55C8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x009E55EA | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x009E560E | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x009E5632 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x009E5697 | `NikePlus_History_WorkoutSummary_Screen_Last` | Known | Screen layout |
| 0x009E5838 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x009E58A8 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x009E58F9 | `NowPlaying_Screen_Video_TVOut` | Known | Screen layout |
| 0x009E5A0C | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x009E5A69 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x009E5AB8 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x009E5B7F | `NikePlus_History_DeleteActiveWorkout_Screen_LastWorkout` | Known | Screen layout |
| 0x009E5CC8 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x009E5CEF | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x009E6147 | `NikePlus_EndPausedWorkout_Screen_BasicSave_Layout` | Known | Screen layout |
| 0x009E6179 | `NikePlus_EndPausedWorkout_Screen_DistanceSave_Layout` | Known | Screen layout |
| 0x009E61AE | `NikePlus_EndPausedWorkout_Screen_TimeSave_Layout` | Known | Screen layout |
| 0x009E61DF | `NikePlus_EndPausedWorkout_Screen_CaloriesSave_Layout` | Known | Screen layout |
| 0x009E6497 | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x009E6694 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x009E6938 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x009E6C55 | `Lock_Screen_Default_Layout` | Known | Screen layout |
| 0x009E6CEB | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x009E6D12 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x009E6F2E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x009E7008 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x009E706F | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009E7099 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009E9F41 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x009E9F8D | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x009EA06B | `MainMenu_List_ScreenLock_x` | Known | Screen layout |
| 0x009EA339 | `RentalInfo_Screen_ExpiresToday` | Known | Screen layout |
| 0x009EA38F | `RentalInfo_Screen_NoAlbumArt_ExpiresToday` | Known | Screen layout |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000904B | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | Known | RTOS |
| 0x002C3158 | `  K - RTXC` | Known | RTOS |
| 0x002C4140 | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | Known | RTOS |
| 0x009CBD10 | `Returning from RTXCBug` | Known | RTOS |

---

## 9. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000DDDE0 | `HostOSTask` | Known | RTOS task thread |
| 0x0013ED04 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x001472D0 | `USBDeviceTask` | Known | RTOS task thread |
| 0x001517D8 | `DiskReaderTask` | Known | RTOS task thread |
| 0x00161394 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x001613A8 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x001B8C14 | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x001FA648 | `MeCCAIOTask` | Known | RTOS task thread |
| 0x002317C8 | `StreamCacheMassStorageManagerTimeOutTask` | Known | RTOS task thread |
| 0x00231944 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x002B6220 | `FirewireTask` | Known | RTOS task thread |
| 0x002B6234 | `TouchwheelTask` | Known | RTOS task thread |
| 0x002B6248 | `AudioOutStateTask` | Known | RTOS task thread |
| 0x002B6274 | `DiskMgrTask` | Known | RTOS task thread |
| 0x002B6284 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x002B6298 | `TopPlugTask` | Known | RTOS task thread |
| 0x002B62A8 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x002B6320 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x002B6348 | `AlarmTask` | Known | RTOS task thread |
| 0x002B6367 | `"USBAudioTask` | Known | RTOS task thread |
| 0x002C37F8 | `Undefined Task` | Known | RTOS task thread |
| 0x003CDF00 | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x003D2408 | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x003DAB50 | `MeCCARecordingTask` | Known | RTOS task thread |
| 0x0091F7F0 | `WaveFileDebugTask` | Known | RTOS task thread |

---

## 10. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0026CA68 | `Channel Reserved` | Known | Logging channel |
| 0x0026CA7C | `Channel AppBoot` | Known | Logging channel |
| 0x0026CA8C | `Channel BufferedSongReading` | Known | Logging channel |
| 0x0026CAA8 | `Channel PrefsWriting` | Known | Logging channel |
| 0x0026CAC0 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x0026CAE0 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x0026CAF8 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x0026CB14 | `Channel TestLogging` | Known | Logging channel |
| 0x0026CB28 | `Channel AppFileLoading` | Known | Logging channel |
| 0x0026CB40 | `Channel VCardReading` | Known | Logging channel |
| 0x0026CB58 | `Channel LongSongScanning` | Known | Logging channel |
| 0x0026CBCC | `Channel VoiceRecording` | Known | Logging channel |
| 0x0026CBE4 | `Channel PhotoImporting` | Known | Logging channel |
| 0x0026CBFC | `Channel Notes` | Known | Logging channel |
| 0x0026CC0C | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x0026CC28 | `Channel DiskMode` | Known | Logging channel |
| 0x0026CC3C | `Channel Firewire` | Known | Logging channel |
| 0x0026CC50 | `Channel USB` | Known | Logging channel |
| 0x0026CC70 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x0026CC88 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0008CFFC | `gamedata_RW` | Known | Game system |
| 0x0008D018 | `gamedata_ShareRW` | Known | Game system |
| 0x0008D02C | `games_RO` | Known | Game system |
| 0x009CBD6A | `iPod_Control/games_RO/` | Known | Game system |
| 0x009CBD81 | `Resources/Games/games_RO/` | Known | Game system |
| 0x009D8DDC | `GamesMenu_StatusBar_String` | Known | Game system |
| 0x009D9516 | `AboutScreen_Games_String` | Known | Game system |
| 0x009E103E | `MainMenu_List_Games` | Known | Game system |
| 0x009E1052 | `ExtrasMenu_Games` | Known | Game system |
| 0x009EA0DA | `MainMenu_List_Games_x` | Known | Game system |

---

## 12. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009E180 | `adrmmp4a` | Known | DRM system |
| 0x0014EDC4 | `AppleDRMVersion` | Known | DRM system |
| 0x0014EE64 | `AppleDRM` | Known | DRM system |
| 0x00150034 | `AppleVideoDRM` | Known | DRM system |
| 0x00153548 | `tx3gdrmsp608aavdmp4aesdsd` | Known | DRM system |
| 0x00209738 | `drmttx3g` | Known | DRM system |
| 0x009CC14F | `DRMLevel` | Known | DRM system |

---

## 13. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00036294 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000362AC | `iPod_Control\iTunes\firsttime` | Known | iTunes database |
| 0x00058858 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x00058880 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00061550 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0008905C | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x0008CF8C | `iPod_Control/iTunes/iTunesDB.p7b` | Known | iTunes database |
| 0x000AA2E4 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x000AA4CC | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000B2F04 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000B43A8 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000B44A8 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00137490 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x003C7274 | `iTunesDB` | Known | iTunes database |
| 0x003C7280 | `iPod_Control\iTunes\` | Known | iTunes database |

---

## 14. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0005FA04 | `NANDDRIVERSIGN` | Known | Hardware |
| 0x00060540 | `NANDDRIVERSIGN` | Known | Hardware |
| 0x00060EF8 | `[FTL:MSG] Apple NAND Driver (AND) 0x%08x` | Known | Hardware |
| 0x00061010 | `[FTL:MSG] Valid Signature not found! Re-initializing NAND!` | Known | Hardware |
| 0x00136CF4 | `NAND FLASH DRIVE` | Known | Hardware |
| 0x0014F30C | `FireWireGUID` | Known | FireWire |
| 0x0014F31C | `FireWireVersion` | Known | FireWire |
| 0x0014F9F8 | `FireWire` | Known | FireWire |
| 0x002D4098 | `[FIL:ERR] No recognized NAND found (0x%X, 0x%X) (line:%d)!` | Known | Hardware |
| 0x009263EC | `[FTL:WRN] Recovering NAND Data Structures - this will take some time!` | Known | Hardware |
| 0x00927930 | `[FIL:WRN]  FNAND_GetStruct 0x%X is not identified is FIL data struct identifier!` | Known | Hardware |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x007601DA | `Settings_RadioRegionsMenu_Layout"` | Known | FM Radio |
| 0x00760263 | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x008166B4 | `Radio Regions` | Known | FM Radio |
| 0x00871520 | `Radio-Regionen` | Known | FM Radio |
| 0x009D5F39 | `Settings_RadioRegions_Australia_String` | Known | FM Radio |
| 0x009D5F60 | `Settings_RadioRegions_Asia_String` | Known | FM Radio |
| 0x009D718E | `Settings_RadioRegions_Europe_String` | Known | FM Radio |
| 0x009D86FE | `Settings_RadioRegions_Japan_String` | Known | FM Radio |
| 0x009D9333 | `Settings_RadioRegions_Americas_String` | Known | FM Radio |
| 0x009D9A15 | `SettingsMenu_RadioRegions_String` | Known | FM Radio |
| 0x009DD3FF | `RadioRegionsMenuItem` | Known | FM Radio |
| 0x009E1800 | `SettingsMenu_ListItem_RadioRegions` | Known | FM Radio |
| 0x009E679D | `SettingsInfo_Template_RadioRegions_Layout` | Known | FM Radio |
| 0x009E67C7 | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x009E6EEF | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 16. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x008B02B0 | `Fotocamera` | Known | Camera |
| 0x008B0440 | `Collega fotocamera o inserisci scheda` | Known | Camera |
| 0x008B04B8 | `Fotocamera non supportata` | Known | Camera |
| 0x008CEE48 | `Camera` | Known | Camera |
| 0x008CEFD4 | `Sluit camera of kaart aan` | Known | Camera |
| 0x008CF040 | `Camera niet ondersteund` | Known | Camera |
| 0x009D5F82 | `Photos_DevName_Camera_String` | Known | Camera |

---

## 17. Pedometer

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00817660 | `Step away from all other sensors.` | Known | Pedometer |
| 0x00817844 | `Step away from all other remotes.` | Known | Pedometer |
| 0x009EA442 | `NikePlus_Remote_Step_Away` | Known | Pedometer |
| 0x009EA45C | `NikePlus_Step_Away` | Known | Pedometer |

---

## 18. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00036280 | `iPod_Control` | Filesystem Path |  |
| 0x000362EC | `iPod_Control\Device` | Filesystem Path |  |
| 0x00045D00 | `iPod_Control\Device` | Filesystem Path |  |
| 0x00047D8C | `iPod_Control` | Filesystem Path |  |
| 0x000483F4 | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x00058838 | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path |  |
| 0x0005B39C | `iPod_Control\Music\` | Filesystem Path |  |
| 0x000613D0 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x00097238 | `iPod_Control` | Filesystem Path |  |
| 0x00097248 | `Resources/Games` | Filesystem Path |  |
| 0x00097258 | `iPod_Control/%s%s%s` | Filesystem Path |  |
| 0x001015FC | `iPod_Control/Device/alarms` | Filesystem Path |  |
| 0x00111938 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00112E3C | `iPod_Control/Device` | Filesystem Path |  |
| 0x00112E50 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00132164 | `iPod_Control/Device/Users` | Filesystem Path |  |
| 0x00162F68 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x001631C4 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x0016EF48 | `Resources/UI/active.bin` | Filesystem Path |  |
| 0x0016EF60 | `Resources/UI/` | Filesystem Path |  |
| 0x001916E4 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x00192DDC | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path |  |
| 0x00192E04 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001BC5CC | `iPod_Control/Device/PlayCounts` | Filesystem Path |  |
| 0x001D3978 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3A28 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3BA4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3D3C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3DE4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3F94 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4038 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D40DC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4180 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4224 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D42D4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4378 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D441C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D44CC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D457C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D462C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4798 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4848 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D48F8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D499C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4A4C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4B40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4BE4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4C98 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4D54 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4E04 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4F28 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4FE4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5094 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5250 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5314 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D53C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5480 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D55BC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5688 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5744 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D57E8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D588C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5948 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5A04 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5ACC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5B70 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5C38 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5D00 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5DB0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5E78 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5F40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5FF0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D60A0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D6164 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D6214 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D62C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D6374 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D6448 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D651C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D661C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D66FC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D6804 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D68F0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003C72F2 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003CD7A0 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x003D02F8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003D06A6 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003D0764 | `iPod_Control/Device/Trainer/Workouts/Empeds/` | Filesystem Path |  |
| 0x003D2574 | `Resources/Fonts` | Filesystem Path |  |
| 0x003DAB1C | `iPod_Control:Device:voices:` | Filesystem Path |  |
| 0x003DCADE | `Resources/TrainerTemplates` | Filesystem Path |  |
| 0x003DCAF9 | `iPod_Control/Device/Trainer/TrainerTemplates` | Filesystem Path |  |
| 0x003DD14C | `iPod_Control/Device/Trainer/Workouts/Empeds/` | Filesystem Path |  |
| 0x003DD1D3 | `/iPod_Control/Device/Trainer/Workouts/Empeds` | Filesystem Path |  |
| 0x009CBC45 | `Resources/Games/` | Filesystem Path |  |
| 0x009CC031 | `iPod_Control/Device` | Filesystem Path |  |
| 0x009CC045 | `iPod_Control/Device/accessories.reg` | Filesystem Path |  |
| 0x009CC0C6 | `iPod_Control/Tones` | Filesystem Path |  |

---

## 19. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0091D42C | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x00922174 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftu` | Build Path |  |
| 0x009221CC | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fts` | Build Path |  |
| 0x00922224 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fto` | Build Path |  |
| 0x00925D84 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x00925DF8 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x00926A14 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x00927148 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x00927594 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x0092E458 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\` | Build Path |  |
| 0x0092EFD4 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfd` | Build Path |  |
| 0x009301D0 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrg` | Build Path |  |
| 0x00930228 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrc` | Build Path |  |
| 0x00930280 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfro` | Build Path |  |
| 0x009305C4 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1` | Build Path |  |
| 0x0093F96C | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttc` | Build Path |  |
| 0x0093FBE8 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype` | Build Path |  |
| 0x00940154 | `c:\bwa\N46FirmwareWin-465\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1` | Build Path |  |

---

## 20. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00095204 | `Acoustic` | EQ Preset |  |
| 0x00095210 | `Bass Booster` | EQ Preset |  |
| 0x00095230 | `Classical` | EQ Preset |  |
| 0x0009523C | `Dance` | EQ Preset |  |
| 0x0009524C | `Electronic` | EQ Preset |  |
| 0x00095260 | `Hip Hop` | EQ Preset |  |
| 0x00095268 | `Jazz` | EQ Preset |  |
| 0x00095270 | `Latin` | EQ Preset |  |
| 0x00095278 | `Loudness` | EQ Preset |  |
| 0x00095284 | `Lounge` | EQ Preset |  |
| 0x0009528C | `Piano` | EQ Preset |  |
| 0x000952A0 | `Rock` | EQ Preset |  |
| 0x000952A8 | `Small Speakers` | EQ Preset |  |
| 0x000952B8 | `Spoken Word` | EQ Preset |  |
| 0x000952C4 | `Treble Booster` | EQ Preset |  |
| 0x00095310 | `Vocal Booster` | EQ Preset |  |
| 0x008169A4 | `Acoustic` | EQ Preset |  |
| 0x008169B0 | `Bass Booster` | EQ Preset |  |
| 0x008169D0 | `Classical` | EQ Preset |  |
| 0x008169DC | `Dance` | EQ Preset |  |
| 0x008169EC | `Electronic` | EQ Preset |  |
| 0x00816A00 | `Hip Hop` | EQ Preset |  |
| 0x00816A08 | `Jazz` | EQ Preset |  |
| 0x00816A10 | `Latin` | EQ Preset |  |
| 0x00816A18 | `Loudness` | EQ Preset |  |
| 0x00816A24 | `Lounge` | EQ Preset |  |
| 0x00816A2C | `Piano` | EQ Preset |  |
| 0x00816A3C | `Rock` | EQ Preset |  |
| 0x00816A44 | `Small Speakers` | EQ Preset |  |
| 0x00816A54 | `Spoken Word` | EQ Preset |  |
| 0x00816A60 | `Treble Booster` | EQ Preset |  |
| 0x00816A80 | `Vocal Booster` | EQ Preset |  |
| 0x0085F168 | `Acoustic` | EQ Preset |  |
| 0x0085F174 | `Bass Booster` | EQ Preset |  |
| 0x0085F194 | `Classical` | EQ Preset |  |
| 0x0085F1A0 | `Dance` | EQ Preset |  |
| 0x0085F1B0 | `Electronic` | EQ Preset |  |
| 0x0085F1C4 | `Hip Hop` | EQ Preset |  |
| 0x0085F1CC | `Jazz` | EQ Preset |  |
| 0x0085F1D4 | `Latin` | EQ Preset |  |
| 0x0085F1DC | `Loudness` | EQ Preset |  |
| 0x0085F1E8 | `Lounge` | EQ Preset |  |
| 0x0085F1F0 | `Piano` | EQ Preset |  |
| 0x0085F200 | `Rock` | EQ Preset |  |
| 0x0085F208 | `Small Speakers` | EQ Preset |  |
| 0x0085F218 | `Spoken Word` | EQ Preset |  |
| 0x0085F224 | `Treble Booster` | EQ Preset |  |
| 0x0085F244 | `Vocal Booster` | EQ Preset |  |
| 0x008682A8 | `Acoustic` | EQ Preset |  |
| 0x008682B4 | `Bass Booster` | EQ Preset |  |
| 0x008682D4 | `Classical` | EQ Preset |  |
| 0x008682E0 | `Dance` | EQ Preset |  |
| 0x008682F0 | `Electronic` | EQ Preset |  |
| 0x00868304 | `Hip Hop` | EQ Preset |  |
| 0x0086830C | `Jazz` | EQ Preset |  |
| 0x00868314 | `Latin` | EQ Preset |  |
| 0x0086831C | `Loudness` | EQ Preset |  |
| 0x00868328 | `Lounge` | EQ Preset |  |
| 0x00868330 | `Piano` | EQ Preset |  |
| 0x00868340 | `Rock` | EQ Preset |  |
| 0x00868348 | `Small Speakers` | EQ Preset |  |
| 0x00868358 | `Spoken Word` | EQ Preset |  |
| 0x00868364 | `Treble Booster` | EQ Preset |  |
| 0x00868384 | `Vocal Booster` | EQ Preset |  |
| 0x008718C8 | `Acoustic` | EQ Preset |  |
| 0x008718F8 | `Dance` | EQ Preset |  |
| 0x00871908 | `Electronic` | EQ Preset |  |
| 0x00871924 | `Jazz` | EQ Preset |  |
| 0x0087192C | `Latin` | EQ Preset |  |
| 0x00871934 | `Loudness` | EQ Preset |  |
| 0x00871948 | `Piano` | EQ Preset |  |
| 0x00871958 | `Rock` | EQ Preset |  |
| 0x00889078 | `Dance` | EQ Preset |  |
| 0x008890A0 | `Hip Hop` | EQ Preset |  |
| 0x008890A8 | `Jazz` | EQ Preset |  |
| 0x008890B8 | `Loudness` | EQ Preset |  |
| 0x008890C4 | `Lounge` | EQ Preset |  |
| 0x008890CC | `Piano` | EQ Preset |  |
| 0x008890DC | `Rock` | EQ Preset |  |
| 0x0089224C | `Jazz` | EQ Preset |  |
| 0x00892254 | `Latin` | EQ Preset |  |
| 0x00892268 | `Lounge` | EQ Preset |  |
| 0x00892270 | `Piano` | EQ Preset |  |
| 0x00892280 | `Rock` | EQ Preset |  |
| 0x0089B324 | `Hip Hop` | EQ Preset |  |
| 0x0089B32C | `Jazz` | EQ Preset |  |
| 0x0089B348 | `Lounge` | EQ Preset |  |
| 0x0089B350 | `Piano` | EQ Preset |  |
| 0x0089B368 | `Rock` | EQ Preset |  |
| 0x008A5010 | `Latin` | EQ Preset |  |
| 0x008A503C | `Rock` | EQ Preset |  |
| 0x008AE668 | `Dance` | EQ Preset |  |
| 0x008AE68C | `Hip Hop` | EQ Preset |  |
| 0x008AE694 | `Jazz` | EQ Preset |  |
| 0x008AE6A4 | `Loudness` | EQ Preset |  |
| 0x008AE6B0 | `Lounge` | EQ Preset |  |
| 0x008AE6B8 | `Piano` | EQ Preset |  |
| 0x008AE6C8 | `Rock` | EQ Preset |  |
| 0x008B9078 | `Acoustic` | EQ Preset |  |
| 0x008B9084 | `Bass Booster` | EQ Preset |  |
| 0x008B90A4 | `Classical` | EQ Preset |  |
| 0x008B90B0 | `Dance` | EQ Preset |  |
| 0x008B90C0 | `Electronic` | EQ Preset |  |
| 0x008B90D4 | `Hip Hop` | EQ Preset |  |
| 0x008B90DC | `Jazz` | EQ Preset |  |
| 0x008B90E4 | `Latin` | EQ Preset |  |
| 0x008B90EC | `Loudness` | EQ Preset |  |
| 0x008B90F8 | `Lounge` | EQ Preset |  |
| 0x008B9100 | `Piano` | EQ Preset |  |
| 0x008B9110 | `Rock` | EQ Preset |  |
| 0x008B9118 | `Small Speakers` | EQ Preset |  |
| 0x008B9128 | `Spoken Word` | EQ Preset |  |
| 0x008B9134 | `Treble Booster` | EQ Preset |  |
| 0x008B9154 | `Vocal Booster` | EQ Preset |  |
| 0x008C392C | `Acoustic` | EQ Preset |  |
| 0x008C3938 | `Bass Booster` | EQ Preset |  |
| 0x008C3958 | `Classical` | EQ Preset |  |
| 0x008C3964 | `Dance` | EQ Preset |  |
| 0x008C3974 | `Electronic` | EQ Preset |  |
| 0x008C3988 | `Hip Hop` | EQ Preset |  |
| 0x008C3990 | `Jazz` | EQ Preset |  |
| 0x008C3998 | `Latin` | EQ Preset |  |
| 0x008C39A0 | `Loudness` | EQ Preset |  |
| 0x008C39AC | `Lounge` | EQ Preset |  |
| 0x008C39B4 | `Piano` | EQ Preset |  |
| 0x008C39C4 | `Rock` | EQ Preset |  |
| 0x008C39CC | `Small Speakers` | EQ Preset |  |
| 0x008C39DC | `Spoken Word` | EQ Preset |  |
| 0x008C39E8 | `Treble Booster` | EQ Preset |  |
| 0x008C3A08 | `Vocal Booster` | EQ Preset |  |
| 0x008CD1BC | `Dance` | EQ Preset |  |
| 0x008CD1F0 | `Jazz` | EQ Preset |  |
| 0x008CD1F8 | `Latin` | EQ Preset |  |
| 0x008CD200 | `Loudness` | EQ Preset |  |
| 0x008CD20C | `Lounge` | EQ Preset |  |
| 0x008CD214 | `Piano` | EQ Preset |  |
| 0x008CD224 | `Rock` | EQ Preset |  |
| 0x008D62D8 | `Dance` | EQ Preset |  |
| 0x008D6304 | `Jazz` | EQ Preset |  |
| 0x008D6314 | `Loudness` | EQ Preset |  |
| 0x008D6320 | `Lounge` | EQ Preset |  |
| 0x008D6328 | `Piano` | EQ Preset |  |
| 0x008D6338 | `Rock` | EQ Preset |  |
| 0x008DF694 | `Hip Hop` | EQ Preset |  |
| 0x008DF69C | `Jazz` | EQ Preset |  |
| 0x008DF6C0 | `Lounge` | EQ Preset |  |
| 0x008DF6D8 | `Rock` | EQ Preset |  |
| 0x008E8E30 | `Hip Hop` | EQ Preset |  |
| 0x008E8E38 | `Jazz` | EQ Preset |  |
| 0x008E8E54 | `Lounge` | EQ Preset |  |
| 0x008E8E5C | `Piano` | EQ Preset |  |
| 0x008E8E6C | `Rock` | EQ Preset |  |
| 0x008FF148 | `Acoustic` | EQ Preset |  |
| 0x008FF154 | `Bass Booster` | EQ Preset |  |
| 0x008FF174 | `Classical` | EQ Preset |  |
| 0x008FF180 | `Dance` | EQ Preset |  |
| 0x008FF190 | `Electronic` | EQ Preset |  |
| 0x008FF1A4 | `Hip Hop` | EQ Preset |  |
| 0x008FF1AC | `Jazz` | EQ Preset |  |
| 0x008FF1B4 | `Latin` | EQ Preset |  |
| 0x008FF1BC | `Loudness` | EQ Preset |  |
| 0x008FF1C8 | `Lounge` | EQ Preset |  |
| 0x008FF1D0 | `Piano` | EQ Preset |  |
| 0x008FF1E0 | `Rock` | EQ Preset |  |
| 0x008FF1E8 | `Small Speakers` | EQ Preset |  |
| 0x008FF1F8 | `Spoken Word` | EQ Preset |  |
| 0x008FF204 | `Treble Booster` | EQ Preset |  |
| 0x008FF224 | `Vocal Booster` | EQ Preset |  |
| 0x009084F4 | `Hip Hop` | EQ Preset |  |
| 0x00908500 | `Latin` | EQ Preset |  |
| 0x00908508 | `Loudness` | EQ Preset |  |
| 0x00908514 | `Lounge` | EQ Preset |  |
| 0x0090852C | `Rock` | EQ Preset |  |
| 0x00911988 | `Acoustic` | EQ Preset |  |
| 0x00911994 | `Bass Booster` | EQ Preset |  |
| 0x009119B4 | `Classical` | EQ Preset |  |
| 0x009119C0 | `Dance` | EQ Preset |  |
| 0x009119D0 | `Electronic` | EQ Preset |  |
| 0x009119E4 | `Hip Hop` | EQ Preset |  |
| 0x009119EC | `Jazz` | EQ Preset |  |
| 0x009119F4 | `Latin` | EQ Preset |  |
| 0x009119FC | `Loudness` | EQ Preset |  |
| 0x00911A08 | `Lounge` | EQ Preset |  |
| 0x00911A10 | `Piano` | EQ Preset |  |
| 0x00911A20 | `Rock` | EQ Preset |  |
| 0x00911A28 | `Small Speakers` | EQ Preset |  |
| 0x00911A38 | `Spoken Word` | EQ Preset |  |
| 0x00911A44 | `Treble Booster` | EQ Preset |  |
| 0x00911A64 | `Vocal Booster` | EQ Preset |  |
| 0x0091ACF4 | `Acoustic` | EQ Preset |  |
| 0x0091AD00 | `Bass Booster` | EQ Preset |  |
| 0x0091AD20 | `Classical` | EQ Preset |  |
| 0x0091AD2C | `Dance` | EQ Preset |  |
| 0x0091AD3C | `Electronic` | EQ Preset |  |
| 0x0091AD50 | `Hip Hop` | EQ Preset |  |
| 0x0091AD58 | `Jazz` | EQ Preset |  |
| 0x0091AD60 | `Latin` | EQ Preset |  |
| 0x0091AD68 | `Loudness` | EQ Preset |  |
| 0x0091AD74 | `Lounge` | EQ Preset |  |
| 0x0091AD7C | `Piano` | EQ Preset |  |
| 0x0091AD8C | `Rock` | EQ Preset |  |
| 0x0091AD94 | `Small Speakers` | EQ Preset |  |
| 0x0091ADA4 | `Spoken Word` | EQ Preset |  |
| 0x0091ADB0 | `Treble Booster` | EQ Preset |  |
| 0x0091ADD0 | `Vocal Booster` | EQ Preset |  |

---
