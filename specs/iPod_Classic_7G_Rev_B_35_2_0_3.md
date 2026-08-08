# iPod Classic 7G Rev B - RetailOS 2.0.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 2.0.3 |
| **IPSW** | iPod_35.2.0.3.ipsw |
| **Device** | iPod Classic 7G Rev B (2009, 160GB, Click Wheel, Cover Flow, Genius, CE-ATA HDD) |
| **UpdaterFamilyID** | 35 |
| **Binary Size** | 10,573,952 bytes (10.08 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,571,904 bytes |
| **Total Strings (>=4)** | 72,600 |
| **Function Prologues** | 22,905 (ARM: 17,538, Thumb: 5,367) |
| **DRAM References** | 106,490 |
| **Peripheral Refs** | 7,225 |
| **Build** | N25CFirmwareWin-33 |
| **SoC** | S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Codename** | N25C |
| **DFU PID** | 0x1223 |
| **SHA-256** | `f78b304df02cf2f9f1aa0eed33db0de5f9f36a778ba21f13b270dab44733086a` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00095788 | `TSilverCntlr` | Known | Controller |
| 0x000957A0 | `TCExtrasMenu` | Known | Controller |
| 0x000957B8 | `TCGameScreen` | Known | Controller |
| 0x000957D0 | `TCGamesMenu` | Known | Controller |
| 0x000957E4 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x0009580C | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00095834 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x00095860 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x00095884 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x000958AC | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x000958D4 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x000958FC | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x00095924 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0009594C | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0009597C | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x000959A8 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x000959D8 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x00095A00 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x00095A28 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x00095A54 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x00095A7C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x00095AA4 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x00095AD4 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x00095B04 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x00095C7C | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x00095CAC | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x00095CD4 | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x00095CFC | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x00095D28 | `TCRentalNotification` | Known | Controller |
| 0x00095D48 | `TCRentalInfo` | Known | Controller |
| 0x00095D60 | `TCRentalConfirmDelete` | Known | Controller |
| 0x00095D80 | `TCRentalDispatcher` | Known | Controller |
| 0x00095DD8 | `TSilverGlobalCntlr` | Known | Controller |
| 0x00095DF4 | `TSilverTrainerCntlr` | Known | Controller |
| 0x000ED27C | `TCSlideshowLCD` | Known | Controller |
| 0x000ED294 | `TCSlideshowTVOut` | Known | Controller |
| 0x000ED2B0 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x000ED2D0 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00110DF0 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x00110E1C | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x00110E48 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x00110E70 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x00110E9C | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x00110EC4 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x00110EF0 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x0011888C | `TCRemoteUI` | Known | Controller |
| 0x001188A0 | `TCUnsupported` | Known | Controller |
| 0x0011ED2C | `TCSpeakers` | Known | Controller |
| 0x0011ED40 | `TCEQSetting` | Known | Controller |
| 0x00147E30 | `TCSportTimer` | Known | Controller |
| 0x00147E48 | `TCSportTimerMenu` | Known | Controller |
| 0x00147E64 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x00147E88 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x00149238 | `TCVoiceMemos` | Known | Controller |
| 0x00149250 | `TCVoiceMemosMenu` | Known | Controller |
| 0x0014926C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x0014928C | `TCVoiceMemosPlayback` | Known | Controller |
| 0x001492AC | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x001492CC | `TCVoiceMemosAlert` | Known | Controller |
| 0x0015B1C0 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x0015B1E8 | `TCSettings_MainMenu` | Known | Controller |
| 0x0015B204 | `TCSettings_MusicMenu` | Known | Controller |
| 0x0015B224 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x0015B244 | `TCSettings_Brightness` | Known | Controller |
| 0x0015B264 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0015B288 | `TCSettings_EQ` | Known | Controller |
| 0x0015B2A0 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x0015B2C8 | `TCSettings_RadioRegions` | Known | Controller |
| 0x0015B2E8 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0015B30C | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x0015B330 | `TCDateTimeScreen` | Known | Controller |
| 0x0015B34C | `TCTimeZoneScreen` | Known | Controller |
| 0x0015B368 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0015B390 | `TCFirstBoot` | Known | Controller |
| 0x00171C0C | `TCDemoMode` | Known | Controller |
| 0x0019AC24 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x0019AC44 | `TCAddressViewerDetails` | Known | Controller |
| 0x0019AC64 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x0019AC88 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x001C8C84 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x001C8CA8 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x001D04EC | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x002673A0 | `TC_LockDialog` | Known | Controller |
| 0x002673B8 | `TC_LockScreen` | Known | Controller |
| 0x002673D0 | `TC_LockediPod` | Known | Controller |
| 0x002673E8 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x0026740C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0026CFCC | `TCClock` | Known | Controller |
| 0x0026CFDC | `TCClockCityMenu` | Known | Controller |
| 0x0026CFF4 | `TCClockRegionMenu` | Known | Controller |
| 0x0026D010 | `TCAlarmMenu` | Known | Controller |
| 0x0026D024 | `TCSleepTimerMenu` | Known | Controller |
| 0x0026D040 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x0026D060 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x0026D088 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x0026D0AC | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0026D0D0 | `TCAlarmDatePicker` | Known | Controller |
| 0x0026D0EC | `TCAlarmTriggered` | Known | Controller |
| 0x00274010 | `TCNotesDispatcher` | Known | Controller |
| 0x0027402C | `TCNotesLoading` | Known | Controller |
| 0x00274044 | `TCNotesList` | Known | Controller |
| 0x00274058 | `TCNotesContents` | Known | Controller |
| 0x003E2890 | `TCAlarmTriggered` | Known | Controller |
| 0x003E28A4 | `TSilverCntlr` | Known | Controller |
| 0x003E28C4 | `TCClock` | Known | Controller |
| 0x003E28CC | `TCClockRegionMenu` | Known | Controller |
| 0x003E28E0 | `TCClockCityMenu` | Known | Controller |
| 0x003E28F0 | `TCAlarmMenu` | Known | Controller |
| 0x003E28FC | `TCSleepTimerMenu` | Known | Controller |
| 0x003E2910 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003E2928 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003E2948 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003E2964 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003E2980 | `TCAlarmDatePicker` | Known | Controller |
| 0x003E29B8 | `TSilverCntlr` | Known | Controller |
| 0x003E29D8 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003E2B68 | `TSilverCntlr` | Known | Controller |
| 0x003E2B88 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x003E2BA8 | `TCSettings_Brightness` | Known | Controller |
| 0x003E2BC0 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x003E2BDC | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x003E2BFC | `TCSettings_RadioRegions` | Known | Controller |
| 0x003E2C14 | `TCSettings_EQ` | Known | Controller |
| 0x003E2C24 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x003E2C40 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x003E2C60 | `TCFirstBoot` | Known | Controller |
| 0x003E2C6C | `TCSettings_MainMenu` | Known | Controller |
| 0x003E2C80 | `TCSettings_MusicMenu` | Known | Controller |
| 0x003E2C98 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003E2CB0 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x003E2CCC | `TCDateTimeScreen` | Known | Controller |
| 0x003E2CE0 | `TCTimeZoneScreen` | Known | Controller |
| 0x003E9DD4 | `TSilverCntlr` | Known | Controller |
| 0x003E9DF4 | `TCClock` | Known | Controller |
| 0x003E9DFC | `TCClockRegionMenu` | Known | Controller |
| 0x003E9E10 | `TCClockCityMenu` | Known | Controller |
| 0x003E9E20 | `TCAlarmMenu` | Known | Controller |
| 0x003E9E2C | `TCSleepTimerMenu` | Known | Controller |
| 0x003E9E40 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003E9EB8 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003E9ED8 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003E9EF4 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003E9F28 | `TCAlarmDatePicker` | Known | Controller |
| 0x003E9F3C | `TCAlarmTriggered` | Known | Controller |
| 0x003EB020 | `TSilverCntlr` | Known | Controller |
| 0x003EB040 | `TC_LockDialog` | Known | Controller |
| 0x003EB050 | `TC_LockScreen` | Known | Controller |
| 0x003EB060 | `TC_LockediPod` | Known | Controller |
| 0x003EB070 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003EB08C | `TCLockChosenDispatcher` | Known | Controller |
| 0x003EB0A4 | `TSilverCntlr` | Known | Controller |
| 0x003EB2B0 | `TSilverCntlr` | Known | Controller |
| 0x003EB2CC | `TSilverCntlr` | Known | Controller |
| 0x003EB330 | `TSilverCntlr` | Known | Controller |
| 0x003EB350 | `TCNotesDispatcher` | Known | Controller |
| 0x003EB364 | `TCNotesLoading` | Known | Controller |
| 0x003EB374 | `TCNotesBase` | Known | Controller |
| 0x003EB380 | `TCNotesList` | Known | Controller |
| 0x003EB38C | `TCNotesContents` | Known | Controller |
| 0x003EB39C | `TSilverCntlr` | Known | Controller |
| 0x003EB3BC | `TCRemoteUI` | Known | Controller |
| 0x003EB3C8 | `TCUnsupported` | Known | Controller |
| 0x003EB3D8 | `TSilverCntlr` | Known | Controller |
| 0x003EB43C | `TSilverCntlr` | Known | Controller |
| 0x003EB45C | `TCSportTimer` | Known | Controller |
| 0x003EB46C | `TCSportTimerMenu` | Known | Controller |
| 0x003EB480 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x003EB49C | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x003EB4CC | `TSilverCntlr` | Known | Controller |
| 0x003EB5F4 | `TSilverCntlr` | Known | Controller |
| 0x003EB614 | `TCDemoMode` | Known | Controller |
| 0x003EB620 | `TCClock` | Known | Controller |
| 0x003EB628 | `TCClockRegionMenu` | Known | Controller |
| 0x003EB63C | `TCClockCityMenu` | Known | Controller |
| 0x003EB64C | `TCAlarmMenu` | Known | Controller |
| 0x003EB658 | `TCSleepTimerMenu` | Known | Controller |
| 0x003EB66C | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003EB684 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003EB6A4 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003EB6C0 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003EB6DC | `TCAlarmDatePicker` | Known | Controller |
| 0x003EB6F0 | `TCAlarmTriggered` | Known | Controller |
| 0x003EB710 | `TSilverCntlr` | Known | Controller |
| 0x003EB72C | `TSilverCntlr` | Known | Controller |
| 0x003EB73C | `TSilverCntlr` | Known | Controller |
| 0x003EB75C | `TCVoiceMemos` | Known | Controller |
| 0x003EB76C | `TCVoiceMemosMenu` | Known | Controller |
| 0x003EB780 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x003EB798 | `TCVoiceMemosAlert` | Known | Controller |
| 0x003EB7AC | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x003EB7C4 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x003EB7E4 | `TSilverCntlr` | Known | Controller |
| 0x003EB844 | `TSilverCntlr` | Known | Controller |
| 0x003EB8B0 | `TSilverCntlr` | Known | Controller |
| 0x003ECBD8 | `TSilverCntlr` | Known | Controller |
| 0x003ECCE4 | `TSilverCntlr` | Known | Controller |
| 0x003F555C | `TSilverCntlr` | Known | Controller |
| 0x003F557C | `TCAddressViewerMainMenu` | Known | Controller |
| 0x003F5594 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x003F55B0 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x003F55D0 | `TCAddressViewerDetails` | Known | Controller |
| 0x003F55E8 | `TSilverCntlr` | Known | Controller |
| 0x003F5608 | `TSilverCalendarCntlr_Base` | Known | Controller |
| 0x003F5624 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x003F5648 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x003F566C | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x003F568C | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x003F56B0 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x003F56D0 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x003F56F4 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x003F58CC | `TSilverCntlr` | Known | Controller |
| 0x003F58EC | `TC_LockDialog` | Known | Controller |
| 0x003F58FC | `TC_LockScreen` | Known | Controller |
| 0x003F590C | `TC_LockediPod` | Known | Controller |
| 0x003F591C | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003F5940 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003F5A60 | `TSilverCntlr` | Known | Controller |
| 0x003F5B94 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003F5BB0 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003F5BD0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003F5BF0 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003F5C18 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003F5C3C | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003F5C64 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003F5C84 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003F5CA4 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003F5CC4 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003F5CE4 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003F5D0C | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003F5D34 | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x003F5D54 | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x003F5D78 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003F5D98 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003F5DB8 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003F5DDC | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003F5DFC | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003F5E20 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003F5E48 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003F5E74 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003F5E94 | `TCRentalNotification` | Known | Controller |
| 0x003F5EAC | `TCRentalInfo` | Known | Controller |
| 0x003F5EBC | `TCRentalConfirmDelete` | Known | Controller |
| 0x003F5ED4 | `TCRentalDispatcher` | Known | Controller |
| 0x003F67C4 | `TSilverCntlr` | Known | Controller |
| 0x003F6888 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003F68A4 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003F68C4 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003F68E4 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003F690C | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003F6930 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003F6958 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003F6978 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003F6998 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003F69B8 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003F69D8 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003F6A00 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003F6A28 | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x003F6A48 | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x003F6A94 | `TCSlideshowTVOut` | Known | Controller |
| 0x003F6AA8 | `TCSlideshowLCD` | Known | Controller |
| 0x003F6AB8 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x003F6AD0 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x003F6AF0 | `TSilverCntlr` | Known | Controller |
| 0x003F6B1C | `TSilverCntlr` | Known | Controller |
| 0x003F6B3C | `TCUnsupported` | Known | Controller |
| 0x003F6B5C | `TSilverCntlr` | Known | Controller |
| 0x003F6B9C | `TSilverCntlr` | Known | Controller |
| 0x003F6BBC | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x003F6BD8 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x003F6BF0 | `TSilverCntlr` | Known | Controller |
| 0x003F6C10 | `TCSpeakers` | Known | Controller |
| 0x003F6C1C | `TCEQSetting` | Known | Controller |
| 0x003F6C3C | `TSilverCntlr` | Known | Controller |
| 0x003F6CA4 | `TSilverCntlr` | Known | Controller |
| 0x003F6CC4 | `TCExtrasMenu` | Known | Controller |
| 0x003F6CD4 | `TCGamesMenu` | Known | Controller |
| 0x003F6CE0 | `TCGameScreen` | Known | Controller |
| 0x003F6CF0 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x003F6D10 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x003F6D30 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x003F6D50 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x003F6D74 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003F6D90 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003F6DB0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003F6DD0 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003F6DF8 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003F6E1C | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003F6E44 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003F6E64 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003F6E84 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003F6EA4 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003F6EC4 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003F6EEC | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003F6F14 | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x003F6F34 | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x003F6F58 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003F6F78 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003F6F98 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003F6FBC | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003F6FDC | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003F7000 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003F7028 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003F7054 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003F7074 | `TCRentalNotification` | Known | Controller |
| 0x003F708C | `TCRentalInfo` | Known | Controller |
| 0x003F709C | `TCRentalConfirmDelete` | Known | Controller |
| 0x003F70B4 | `TCRentalDispatcher` | Known | Controller |
| 0x003F70C8 | `TSilverGlobalCntlr` | Known | Controller |
| 0x003F70DC | `TSilverTrainerCntlr` | Known | Controller |
| 0x0047C884 | `TCCCCCCCCCCCCCCCCCCCCCCT` | Known | Controller |
| 0x00720A6A | `TCNotesDispatcher"` | Known | Controller |
| 0x00720B29 | `TCLockChosenDispatcher"` | Known | Controller |
| 0x00720BEC | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x0072ACDB | `TCNotesDispatcher"` | Known | Controller |
| 0x0072AE3D | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x0074019C | `TCAddressViewerMainMenu` | Known | Controller |
| 0x007401B4 | `TCAddressViewerDetails` | Known | Controller |
| 0x007401CC | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x007401E8 | `TCAlarmMenu` | Known | Controller |
| 0x007401F4 | `TCSleepTimerMenuTCAlarmPropertiesMenu` | Known | Controller |
| 0x0074021C | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x0074023C | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x00740258 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00740274 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00740290 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x007402AC | `TCAlarmDatePicker` | Known | Controller |
| 0x007402C0 | `TCAlarmDatePicker` | Known | Controller |
| 0x007402D4 | `TCAlarmTriggeredTSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x00740300 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x00740324 | `TSilverCalendarCntlr_MonthViewerTSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x00740364 | `TSilverCalendarCntlr_EventViewerTSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x007403A4 | `TSilverCalendarCntlr_EventViewerTSilverCntlrTCClockRegionMenu` | Known | Controller |
| 0x007403E4 | `TCClockCityMenu` | Known | Controller |
| 0x007403F4 | `TCClockCityMenu` | Known | Controller |
| 0x00740404 | `TCClockCityMenu` | Known | Controller |
| 0x00740414 | `TCClockCityMenu` | Known | Controller |
| 0x00740424 | `TCClockCityMenu` | Known | Controller |
| 0x00740434 | `TCClockCityMenu` | Known | Controller |
| 0x00740444 | `TCClockCityMenu` | Known | Controller |
| 0x00740454 | `TCClockCityMenu` | Known | Controller |
| 0x00740464 | `TCClock` | Known | Controller |
| 0x0074047C | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDe` | Known | Controller |
| 0x007404D4 | `TCGamesMenu` | Known | Controller |
| 0x007404E0 | `TCGameScreenTC_LockediPod` | Known | Controller |
| 0x007404FC | `TC_LockDialog` | Known | Controller |
| 0x0074050C | `TC_LockScreen` | Known | Controller |
| 0x0074051C | `TC_VolumeLimitLockScreenTSilverCntlrTSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x00740560 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00740580 | `TSilverMediaListCntlr_AudiobooksTSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x007405C8 | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x007405EC | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x00740608 | `TSilverMediaListCntlr_AlbumsTSilverMediaListCntlr_Artists` | Known | Controller |
| 0x00740644 | `TSilverMediaListCntlr_GenresTSilverMediaListCntlr_Composers` | Known | Controller |
| 0x00740680 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x007406A0 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x007406C8 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x007406E8 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x00740708 | `TSilverMediaListCntlr_TVEpisodesTSilverMediaListCntlr_MoviesTSilverMediaListCntl` | Known | Controller |
| 0x00740764 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0074078C | `TSilverMediaListCntlr_VideoPlaylistsTSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x007407D0 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x007407FC | `TSilverCntlrTSilverMainMediaListCntlr_VideosTSilverSettingsVideoCntlr` | Known | Controller |
| 0x00740844 | `TCFirstBoot` | Known | Controller |
| 0x00740850 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00740870 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00740890 | `TSilverMediaListCntlr_GeniusTSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x00740970 | `TCRentalInfoTCRentalConfirmDelete` | Known | Controller |
| 0x00740994 | `TSilverCntlrTCRentalNotificationTCRentalNotificationTCRentalNotificationTCNotesL` | Known | Controller |
| 0x007409EC | `TCNotesList` | Known | Controller |
| 0x007409F8 | `TCNotesList` | Known | Controller |
| 0x00740A04 | `TCNotesContents` | Known | Controller |
| 0x00740A14 | `TCNotesContents` | Known | Controller |
| 0x00740A24 | `TCNotesContents` | Known | Controller |
| 0x00740A34 | `TCNotesContents` | Known | Controller |
| 0x00740AF0 | `TCSlideshowLCD` | Known | Controller |
| 0x00740B00 | `TCSlideshowTVOutTCSlideshow_TVOutAskTCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00740B50 | `TCRemoteUI` | Known | Controller |
| 0x00740B5C | `TCUnsupported` | Known | Controller |
| 0x00740B6C | `TSilverSettingsMenuListCntlrTSilverSettingsMenuListCntlrTSilverSettingsMenuListC` | Known | Controller |
| 0x00740BD4 | `TCSettings_MusicMenuTCSettings_VolumeLimit` | Known | Controller |
| 0x00740C00 | `TCSettings_Brightness` | Known | Controller |
| 0x00740C18 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x00740C34 | `TCSettings_AudiobookSettingsTCSettings_RadioRegions` | Known | Controller |
| 0x00740C68 | `TCSettings_EQ` | Known | Controller |
| 0x00740C78 | `TSilverSettingsMenuListCntlrTSilverCntlrTCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x00740CC0 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x00740CDC | `TCSettings_MainMenu` | Known | Controller |
| 0x00740CF0 | `TCSettings_MusicMenuTCSportTimerTCSportTimerMenuTCSportTimerSessionScreen` | Known | Controller |
| 0x00740D3C | `TSilverCntlrTUnitTestSuiteCntlr` | Known | Controller |
| 0x00740DBC | `TCVoiceMemosTCVoiceMemosAlert` | Known | Controller |
| 0x00740DDC | `TCVoiceMemosAlert` | Known | Controller |
| 0x00740DF0 | `TCVoiceMemosAlert` | Known | Controller |
| 0x00740E1C | `TCEQSetting` | Known | Controller |
| 0x00740F8A | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x007424E9 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x00748365 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007483C3 | `TCNotesDispatcher` | Known | Controller |
| 0x0074A2A1 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074A2FF | `TCNotesDispatcher` | Known | Controller |
| 0x0074C1DD | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074C23B | `TCNotesDispatcher` | Known | Controller |
| 0x0074E119 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074E177 | `TCNotesDispatcher` | Known | Controller |
| 0x00750055 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007500B3 | `TCNotesDispatcher` | Known | Controller |
| 0x00751F91 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00751FEF | `TCNotesDispatcher` | Known | Controller |
| 0x00753ECD | `TCLockChosenDispatcher` | Known | Controller |
| 0x00753F2B | `TCNotesDispatcher` | Known | Controller |
| 0x00755E09 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00755E67 | `TCNotesDispatcher` | Known | Controller |
| 0x00757D45 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00757DA3 | `TCNotesDispatcher` | Known | Controller |
| 0x00759C81 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00759CDF | `TCNotesDispatcher` | Known | Controller |
| 0x0075BBBD | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075BC1B | `TCNotesDispatcher` | Known | Controller |
| 0x0075DAF9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075DB57 | `TCNotesDispatcher` | Known | Controller |
| 0x0075FA35 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075FA93 | `TCNotesDispatcher` | Known | Controller |
| 0x00761971 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007619CF | `TCNotesDispatcher` | Known | Controller |
| 0x007638AD | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076390B | `TCNotesDispatcher` | Known | Controller |
| 0x007657E9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00765847 | `TCNotesDispatcher` | Known | Controller |
| 0x00767725 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00767783 | `TCNotesDispatcher` | Known | Controller |
| 0x00769661 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007696BF | `TCNotesDispatcher` | Known | Controller |
| 0x0076B59D | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076B5FB | `TCNotesDispatcher` | Known | Controller |
| 0x0076D4D9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076D537 | `TCNotesDispatcher` | Known | Controller |
| 0x0076F415 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076F473 | `TCNotesDispatcher` | Known | Controller |
| 0x00771351 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007713AF | `TCNotesDispatcher` | Known | Controller |
| 0x0077328D | `TCLockChosenDispatcher` | Known | Controller |
| 0x007732EB | `TCNotesDispatcher` | Known | Controller |
| 0x007751C9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00775227 | `TCNotesDispatcher` | Known | Controller |
| 0x00777105 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00777163 | `TCNotesDispatcher` | Known | Controller |
| 0x00779041 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077909F | `TCNotesDispatcher` | Known | Controller |
| 0x0077AF7D | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077AFDB | `TCNotesDispatcher` | Known | Controller |
| 0x0077CEB9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077CF17 | `TCNotesDispatcher` | Known | Controller |
| 0x0077EDF5 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077EE53 | `TCNotesDispatcher` | Known | Controller |
| 0x00780D31 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00780D8F | `TCNotesDispatcher` | Known | Controller |
| 0x00782C6D | `TCLockChosenDispatcher` | Known | Controller |
| 0x00782CCB | `TCNotesDispatcher` | Known | Controller |
| 0x00784BA9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00784C07 | `TCNotesDispatcher` | Known | Controller |
| 0x00786AE5 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00786B43 | `TCNotesDispatcher` | Known | Controller |
| 0x00788A21 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00788A7F | `TCNotesDispatcher` | Known | Controller |
| 0x0078A95D | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078A9BB | `TCNotesDispatcher` | Known | Controller |
| 0x0078C899 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078C8F7 | `TCNotesDispatcher` | Known | Controller |
| 0x0078E7D5 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078E833 | `TCNotesDispatcher` | Known | Controller |
| 0x0079DB88 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x0079DE4A | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x0079E680 | `TCRentalDispatcher` | Known | Controller |
| 0x0079EF38 | `TCRentalDispatcher` | Known | Controller |
| 0x0079F7F0 | `TCRentalDispatcher` | Known | Controller |
| 0x007A00A8 | `TCRentalDispatcher` | Known | Controller |
| 0x007A0960 | `TCRentalDispatcher` | Known | Controller |
| 0x007A1218 | `TCRentalDispatcher` | Known | Controller |
| 0x007A1AD0 | `TCRentalDispatcher` | Known | Controller |
| 0x007A2388 | `TCRentalDispatcher` | Known | Controller |
| 0x008EB604 | `TCMockupModeNavScreen` | Known | Controller |
| 0x008EB61C | `TSilverCntlr` | Known | Controller |
| 0x008EB63C | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x008EB68C | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x008EB6AC | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x008EB6CC | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x008EB6F0 | `TCExtrasMenu` | Known | Controller |
| 0x008EB800 | `TSilverCntlr` | Known | Controller |
| 0x008EB820 | `TCSlideshowTVOut` | Known | Controller |
| 0x008EB834 | `TCSlideshowLCD` | Known | Controller |
| 0x008EB844 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x008EB85C | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x008EB87C | `TSilverGlobalCntlr` | Known | Controller |
| 0x008EB8AC | `TSilverCntlr` | Known | Controller |
| 0x008EB928 | `TCSlideshowTVOut` | Known | Controller |
| 0x008EB93C | `TCSlideshowLCD` | Known | Controller |
| 0x008EB94C | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x008EB964 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x008EB984 | `TSilverCntlr` | Known | Controller |
| 0x008EB9CC | `TSilverCntlr` | Known | Controller |
| 0x008EB9EC | `TCGamesMenu` | Known | Controller |
| 0x008EB9F8 | `TCGameScreen` | Known | Controller |
| 0x009AA0BB | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |

---

## 2. Settings (Toggle/Show/TC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00128458 | `ShowSetting_EQ` | Known | User setting |
| 0x001D249C | `ToggleSetting_Repeat` | Known | User setting |
| 0x001D24B8 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001D24D0 | `ToggleSetting_TVOut` | Known | User setting |
| 0x001D24E4 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x001FB214 | `ShowSetting_Backlight` | Known | User setting |
| 0x002102C8 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x002102E4 | `ToggleSetting_Repeat` | Known | User setting |
| 0x002102FC | `ToggleSetting_SortBy` | Known | User setting |
| 0x00210314 | `ToggleSetting_ClassicUI` | Known | User setting |
| 0x0021032C | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x00210348 | `ToggleSetting_Clicker` | Known | User setting |
| 0x00210360 | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x00210380 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x0021039C | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x002103B8 | `ShowSetting_Shuffle` | Known | User setting |
| 0x00210564 | `ShowSetting_Repeat` | Known | User setting |
| 0x00210578 | `ShowSetting_About` | Known | User setting |
| 0x0021058C | `ShowSetting_MainMenu` | Known | User setting |
| 0x002105A4 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x002105BC | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x002105D4 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x002105F0 | `ShowSetting_Brightness` | Known | User setting |
| 0x00210608 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x00210620 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x0021063C | `ShowSetting_EQ` | Known | User setting |
| 0x0021064C | `ShowSetting_SoundCheck` | Known | User setting |
| 0x002107E8 | `ShowSetting_Clicker` | Known | User setting |
| 0x002107FC | `ShowSetting_DateAndTime` | Known | User setting |
| 0x00210814 | `ShowSetting_SortBy` | Known | User setting |
| 0x00210828 | `ShowSetting_ClassicUI` | Known | User setting |
| 0x00210840 | `ShowSetting_Language` | Known | User setting |
| 0x00210858 | `ShowSetting_Legal` | Known | User setting |
| 0x0021086C | `ShowSetting_ResetAll` | Known | User setting |
| 0x00729AE1 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x00729B91 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x0072C2F2 | `ShowSetting_About` | Known | User setting |
| 0x0072C394 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x0072C3D8 | `ShowSetting_Shuffle` | Known | User setting |
| 0x0072C44F | `ToggleSetting_Repeat` | Known | User setting |
| 0x0072C492 | `ShowSetting_Repeat` | Known | User setting |
| 0x0072C59C | `ShowSetting_MainMenu` | Known | User setting |
| 0x0072C6AC | `ShowSetting_MusicMenu` | Known | User setting |
| 0x0072C774 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x0072C83E | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x0072C956 | `ShowSetting_Brightness` | Known | User setting |
| 0x0072CA8C | `ShowSetting_Audiobooks` | Known | User setting |
| 0x0072CB9D | `ShowSetting_RadioRegions` | Known | User setting |
| 0x0072CC9E | `ShowSetting_EQ` | Known | User setting |
| 0x0072CD0B | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x0072CD52 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x0072CDCF | `ToggleSetting_Clicker` | Known | User setting |
| 0x0072CE13 | `ShowSetting_Clicker` | Known | User setting |
| 0x0072CF7A | `ToggleSetting_SortBy` | Known | User setting |
| 0x0072CFBD | `ShowSetting_SortBy` | Known | User setting |
| 0x0072D0BE | `ShowSetting_Language` | Known | User setting |
| 0x0072D1CE | `ShowSetting_Legal` | Known | User setting |
| 0x0072D2FF | `ShowSetting_ResetAll` | Known | User setting |
| 0x0072D470 | `ShowSetting_Backlight` | Known | User setting |
| 0x0072D520 | `ShowSetting_Backlight` | Known | User setting |
| 0x0072D5D0 | `ShowSetting_Backlight` | Known | User setting |
| 0x0072D681 | `ShowSetting_Backlight` | Known | User setting |
| 0x0072D732 | `ShowSetting_Backlight` | Known | User setting |
| 0x0072D7E3 | `ShowSetting_Backlight` | Known | User setting |
| 0x0072D897 | `ShowSetting_Backlight` | Known | User setting |
| 0x0072D946 | `ShowSetting_EQ` | Known | User setting |
| 0x0072D9BB | `ShowSetting_Language` | Known | User setting |
| 0x007BA9B4 | `ToggleSetting_Repeat` | Known | User setting |
| 0x007BA9EE | `ToggleSetting_Shuffle` | Known | User setting |
| 0x007BAAB0 | `ToggleSetting_TVOut` | Known | User setting |
| 0x007BAAE9 | `ToggleSetting_TVSignal` | Known | User setting |

---

## 3. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00143CA8 | `MockupMode/MockupMode.xml` | Hidden | Developer Tool |
| 0x001441A8 | `MockupMode/` | Hidden | Developer Tool |
| 0x0024CA6C | `Channel UnitTests` | Hidden | Developer Tool |
| 0x002A6DAD | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x002A6DF0 | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x002A6E05 | `RTXCbug> ` | Hidden | Developer Tool |
| 0x002A77E1 | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x002C1310 | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x00389945 | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x00389A0D | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x003E7CD5 | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x00740D5C | `TUnitTestSuiteTestsCntlrTSilverCntlrTCVoiceMemosMainMenuTCVoiceMemosMenuTCVoiceM` | Hidden | Developer Tool |
| 0x007E1CBC | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0081F58C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00832430 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0084A5D8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0085D10C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008671DC | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00870F8C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00886954 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00890984 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008B7EE8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008D7084 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008E07F8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0096D001 | `10TCDemoMode` | Hidden | Demo/Retail Mode |
| 0x0096D978 | `21TCMockupModeNavScreen` | Hidden | Developer Tool |
| 0x0096DE38 | `27TSilverCntlrTransitionAddonI10TCDemoModeE` | Hidden | Demo/Retail Mode |
| 0x0099C197 | `Debug_ListItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x0099C1AF | `Debug_MenuItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x0099C8B4 | `DemoMode_Background_Image` | Hidden | Demo/Retail Mode |
| 0x0099D4F1 | `DemoMode_Icon_Image` | Hidden | Demo/Retail Mode |
| 0x0099F0B3 | `Debug_Menu_Title` | Hidden | Debug/Diagnostic |
| 0x0099F0D8 | `UnitTest_Menu_Title` | Hidden | Developer Tool |
| 0x009A7C21 | `UnitTestModel` | Hidden | Developer Tool |
| 0x009A8600 | `UnitTest_ListItem` | Hidden | Developer Tool |
| 0x009A974D | `DemoMode_View_Main` | Hidden | Demo/Retail Mode |
| 0x009A9922 | `DemoMode_Icon` | Hidden | Demo/Retail Mode |
| 0x009AA723 | `UnitTestApp` | Hidden | Developer Tool |
| 0x009AACD5 | `Debug_ListItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009AACF0 | `Debug_MenuItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009AB44C | `DemoMode_Text_Color` | Hidden | Demo/Retail Mode |
| 0x009AB861 | `Debug_ListItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009AB878 | `Debug_MenuItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009AF9BE | `Debug_ListItem_UnitTest` | Hidden | Developer Tool |
| 0x009AF9D6 | `Debug_MenuItem_UnitTest` | Hidden | Developer Tool |
| 0x009B3F50 | `Debug_ListItem_Memory` | Hidden | Debug/Diagnostic |
| 0x009B3F66 | `Debug_MenuItem_Memory` | Hidden | Debug/Diagnostic |

---

## 4. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000067BB | `"MeCCADecode` | Known | Audio system |
| 0x00139B3C | `AudioCodecs` | Known | Audio system |
| 0x00151630 | `MeCCA_RecordingBuffer` | Known | Audio system |
| 0x001803DC | `MeCCA_PCM_Output.wav` | Known | Audio system |
| 0x00199E60 | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x001A4BB0 | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x001A4DB8 | `MeCCAVideoDecode` | Known | Audio system |
| 0x008F7A40 | `MeCCA_StreamCache` | Known | Audio system |

---

## 5. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E8A04 | `HandleWheel` | Known | Event handler |
| 0x000E8A10 | `HandlePlayPause` | Known | Event handler |
| 0x000E8A20 | `HandleSelectDown` | Known | Event handler |
| 0x000E8A34 | `HandleNext` | Known | Event handler |
| 0x000E8A40 | `HandlePrevious` | Known | Event handler |
| 0x000E8A50 | `HandleNextPushAndHold` | Known | Event handler |
| 0x000E8A68 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000E8D00 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000E8D20 | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x000F513C | `HandleSelect` | Known | Event handler |
| 0x000F5150 | `HandleHilite` | Known | Event handler |
| 0x000F54E8 | `HandleEQSettingSelected` | Known | Event handler |
| 0x000F5918 | `HandleSelect` | Known | Event handler |
| 0x000F592C | `HandleGameHilited` | Known | Event handler |
| 0x000F5BDC | `HandleNotesSelected` | Known | Event handler |
| 0x000F5BF4 | `HandleNotesPop` | Known | Event handler |
| 0x000F5C04 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00103FA8 | `HandleVolumeWheel` | Known | Event handler |
| 0x00103FBC | `HandleVolumeChange` | Known | Event handler |
| 0x00103FD0 | `HandleTimerDone` | Known | Event handler |
| 0x00103FE0 | `HandleFrequencyChange` | Known | Event handler |
| 0x00104058 | `HandleTuning` | Known | Event handler |
| 0x00104068 | `HandleTuningSelect` | Known | Event handler |
| 0x0010EBAC | `HandleLock` | Known | Event handler |
| 0x0010EBBC | `HandleAddressBook` | Known | Event handler |
| 0x0010F2A4 | `HandleSelect` | Known | Event handler |
| 0x0010F7DC | `HandleExit` | Known | Event handler |
| 0x0010F7EC | `HandleLap` | Known | Event handler |
| 0x0010F7F8 | `HandleResume` | Known | Event handler |
| 0x0010F808 | `HandleStartStop` | Known | Event handler |
| 0x0010FABC | `HandleWheel` | Known | Event handler |
| 0x0010FACC | `HandlePlayPause` | Known | Event handler |
| 0x0010FADC | `HandleSelectDown` | Known | Event handler |
| 0x0010FAF0 | `HandleHilite` | Known | Event handler |
| 0x0010FB14 | `HandleFinishRecording` | Known | Event handler |
| 0x0011A0A8 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x0012868C | `HandleExitUnsupported` | Known | Event handler |
| 0x0013F4F4 | `HandleNotesPop` | Known | Event handler |
| 0x0013F508 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00140414 | `HandleSelect` | Known | Event handler |
| 0x00140428 | `HandleWheel` | Known | Event handler |
| 0x00140434 | `HandleImageNext` | Known | Event handler |
| 0x00140444 | `HandleImagePrev` | Known | Event handler |
| 0x00140454 | `HandleImageLast` | Known | Event handler |
| 0x00140464 | `HandleImageFirst` | Known | Event handler |
| 0x00140478 | `HandlePlayPause` | Known | Event handler |
| 0x00140488 | `HandlePlay` | Known | Event handler |
| 0x00140494 | `HandlePause` | Known | Event handler |
| 0x001404A0 | `HandleMikeyCenter` | Known | Event handler |
| 0x001554D8 | `HandleSelectCity` | Known | Event handler |
| 0x001554F0 | `HandleHighlightCity` | Known | Event handler |
| 0x001565DC | `HandleWantPopFlow` | Known | Event handler |
| 0x001565F4 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x00156610 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0015662C | `HandleFlowNext` | Known | Event handler |
| 0x0015663C | `HandleFlowPrev` | Known | Event handler |
| 0x0015664C | `HandleFlowWheel` | Known | Event handler |
| 0x0015665C | `HandleAlbumSelected` | Known | Event handler |
| 0x00156670 | `HandlePlayPause` | Known | Event handler |
| 0x00156680 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00182278 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00182668 | `HandleSelect` | Known | Event handler |
| 0x00183550 | `HandleSelect` | Known | Event handler |
| 0x00183564 | `HandleWheel` | Known | Event handler |
| 0x00183570 | `HandleImageNext` | Known | Event handler |
| 0x00183580 | `HandleImagePrev` | Known | Event handler |
| 0x00183590 | `HandleImageLast` | Known | Event handler |
| 0x001835A0 | `HandleImageFirst` | Known | Event handler |
| 0x001835B4 | `HandlePlayPause` | Known | Event handler |
| 0x001835C4 | `HandlePlay` | Known | Event handler |
| 0x001835D0 | `HandlePause` | Known | Event handler |
| 0x001835DC | `HandleMikeyCenter` | Known | Event handler |
| 0x00183A84 | `HandleNew` | Known | Event handler |
| 0x00183A94 | `HandleClear` | Known | Event handler |
| 0x00183AA0 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x00183ABC | `HandleSelectIndexedSession` | Known | Event handler |
| 0x00183DCC | `HandleWheel` | Known | Event handler |
| 0x00183DDC | `HandleArrowUp` | Known | Event handler |
| 0x00183DEC | `HandleArrowDown` | Known | Event handler |
| 0x00186A78 | `HandleHiliteAlbum` | Known | Event handler |
| 0x00186A90 | `HandleBrowseAlbum` | Known | Event handler |
| 0x00186AA4 | `HandlePlayPause` | Known | Event handler |
| 0x0019D480 | `HandleSelect` | Known | Event handler |
| 0x0019D610 | `HandleSelectRegion` | Known | Event handler |
| 0x0019D988 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x0019D9A4 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x0019D9C0 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x001B473C | `HandleImageWheel` | Known | Event handler |
| 0x001B4754 | `HandlePlayPause` | Known | Event handler |
| 0x001B4764 | `HandleBrowseLarge` | Known | Event handler |
| 0x001B4778 | `HandleBrowseSmall` | Known | Event handler |
| 0x001B478C | `HandleBrowseSlideshow` | Known | Event handler |
| 0x001B47A4 | `HandleImageNext` | Known | Event handler |
| 0x001B47B4 | `HandleImagePrev` | Known | Event handler |
| 0x001B47C4 | `HandleHilite` | Known | Event handler |
| 0x001B47D4 | `HandleImageLast` | Known | Event handler |
| 0x001B47E4 | `HandleImageFirst` | Known | Event handler |
| 0x001B47F8 | `HandleScreenNext` | Known | Event handler |
| 0x001B480C | `HandleScreenPrev` | Known | Event handler |
| 0x001B70D4 | `HandlePlayPause` | Known | Event handler |
| 0x001B70E8 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x001B7104 | `HandleNext` | Known | Event handler |
| 0x001B7110 | `HandleNextPressAndHold` | Known | Event handler |
| 0x001B7128 | `HandlePrevious` | Known | Event handler |
| 0x001B7138 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x001B7154 | `HandleRemotePlayPause` | Known | Event handler |
| 0x001B716C | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x001B7190 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x001B71A8 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x001B71C0 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x001B7364 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x001B737C | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x001B7394 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x001B73B0 | `HandleRemoteStop` | Known | Event handler |
| 0x001B73C4 | `HandleRemotePlay` | Known | Event handler |
| 0x001B73D8 | `HandleRemotePause` | Known | Event handler |
| 0x001B73EC | `HandleRemoteMute` | Known | Event handler |
| 0x001B7400 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x001B7418 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x001B7430 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x001B744C | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x001B7654 | `HandleRemoteShuffle` | Known | Event handler |
| 0x001B7668 | `HandleRemoteRepeat` | Known | Event handler |
| 0x001B767C | `HandleRemoteOn` | Known | Event handler |
| 0x001B7690 | `HandleRemoteOff` | Known | Event handler |
| 0x001B76A0 | `HandleRemoteBacklight` | Known | Event handler |
| 0x001B76B8 | `HandleRemoteFFDown` | Known | Event handler |
| 0x001B76CC | `HandleRemoteFFUp` | Known | Event handler |
| 0x001B76E0 | `HandleRemoteRewDown` | Known | Event handler |
| 0x001B76F4 | `HandleRemoteRewUp` | Known | Event handler |
| 0x001B7708 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x001B7720 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x001B7734 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x001B774C | `HandleRemoteSelectUp` | Known | Event handler |
| 0x001B78FC | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x001B7914 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x001B792C | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x001B7948 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x001B7960 | `HandleRemoteEvent` | Known | Event handler |
| 0x001B7974 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x001B7990 | `HandleAudioPlayPause` | Known | Event handler |
| 0x001B79A8 | `HandleAudioNext` | Known | Event handler |
| 0x001B79B8 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x001B79D4 | `HandleAudioPrevious` | Known | Event handler |
| 0x001B79E8 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x001B7B78 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x001B7B90 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x001B7BA8 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x001B7BC0 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x001B7BD4 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x001B7BEC | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x001B7C04 | `HandleAudioStop` | Known | Event handler |
| 0x001B7C14 | `HandleAudioPlay` | Known | Event handler |
| 0x001B7C24 | `HandleAudioPause` | Known | Event handler |
| 0x001B7C38 | `HandleAudioMute` | Known | Event handler |
| 0x001B7C48 | `HandleAudioNextChapter` | Known | Event handler |
| 0x001B7C60 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x001B7E4C | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x001B7E64 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x001B7E7C | `HandleAudioShuffle` | Known | Event handler |
| 0x001B7E90 | `HandleAudioRepeat` | Known | Event handler |
| 0x001B7EA4 | `HandleAudioFFDown` | Known | Event handler |
| 0x001B7EB8 | `HandleAudioFFUp` | Known | Event handler |
| 0x001B7EC8 | `HandleAudioRewDown` | Known | Event handler |
| 0x001B7EDC | `HandleAudioRewUp` | Known | Event handler |
| 0x001B7EF0 | `HandleVideoPlayPause` | Known | Event handler |
| 0x001B7F08 | `HandleVideoNext` | Known | Event handler |
| 0x001B7F18 | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x001B7F34 | `HandleVideoPrevious` | Known | Event handler |
| 0x001B7F48 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x001B810C | `HandleVideoStop` | Known | Event handler |
| 0x001B811C | `HandleVideoPlay` | Known | Event handler |
| 0x001B812C | `HandleVideoPause` | Known | Event handler |
| 0x001B8140 | `HandleVideoFFDown` | Known | Event handler |
| 0x001B8154 | `HandleVideoFFUp` | Known | Event handler |
| 0x001B8164 | `HandleVideoRewDown` | Known | Event handler |
| 0x001B8178 | `HandleVideoRewUp` | Known | Event handler |
| 0x001B818C | `HandleVideoNextChapter` | Known | Event handler |
| 0x001B81A4 | `HandleVideoPrevChapter` | Known | Event handler |
| 0x001B81BC | `HandleVideoNextFrame` | Known | Event handler |
| 0x001B81D4 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x001B81EC | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x001B8208 | `HandleMikeyCenter` | Known | Event handler |
| 0x001B8368 | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x001B8388 | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x001B83A8 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x001B83CC | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x001B83EC | `HandleMikeyAllUp` | Known | Event handler |
| 0x001B8400 | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x001B8414 | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x001B842C | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x001B8444 | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x001C520C | `HandleMainMenu` | Known | Event handler |
| 0x001C7358 | `HandleLoadingCancelled` | Known | Event handler |
| 0x001C9E68 | `HandlePowerSongSelected` | Known | Event handler |
| 0x001C9E84 | `HandlePowerSongChosen` | Known | Event handler |
| 0x001C9E9C | `HandleSelectTimedWorkout` | Known | Event handler |
| 0x001D0404 | `HandleSelect` | Known | Event handler |
| 0x001D06AC | `HandleMusicMenu` | Known | Event handler |
| 0x001D096C | `HandleSelect` | Known | Event handler |
| 0x001D0C70 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x001D0C90 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x001D114C | `HandleWheel` | Known | Event handler |
| 0x001D115C | `HandlePlayPause` | Known | Event handler |
| 0x001D116C | `HandleSelectDown` | Known | Event handler |
| 0x001D1180 | `HandleNext` | Known | Event handler |
| 0x001D118C | `HandlePrevious` | Known | Event handler |
| 0x001D119C | `HandleNextPushAndHold` | Known | Event handler |
| 0x001D11B4 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x001D18A8 | `HandleMenuSelection` | Known | Event handler |
| 0x001D18BC | `HandleViewAlbum` | Known | Event handler |
| 0x001D18CC | `HandleViewArtist` | Known | Event handler |
| 0x001D18E0 | `HandleViewCompilation` | Known | Event handler |
| 0x001D18F8 | `HandleStartGenius` | Known | Event handler |
| 0x001DE4F4 | `HandleFrequencyChosen` | Known | Event handler |
| 0x001DE50C | `HandleDateChosen` | Known | Event handler |
| 0x001DE520 | `HandleTimeChosen` | Known | Event handler |
| 0x001DE534 | `HandleSoundChosen` | Known | Event handler |
| 0x001DE548 | `HandleLabelChosen` | Known | Event handler |
| 0x001DE55C | `HandleDeleteChosen` | Known | Event handler |
| 0x001DF63C | `HandleSelect` | Known | Event handler |
| 0x001E3F64 | `HandlePrev` | Known | Event handler |
| 0x001E3F74 | `HandleNext` | Known | Event handler |
| 0x001E3F80 | `HandlePlayPause` | Known | Event handler |
| 0x001EB93C | `HandleNextContact` | Known | Event handler |
| 0x001EB954 | `HandlePreviousContact` | Known | Event handler |
| 0x001F3504 | `HandleItemSelected` | Known | Event handler |
| 0x001F36FC | `HandleRadioRegion` | Known | Event handler |
| 0x001F38E4 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x001F7B58 | `HandlePlayPause` | Known | Event handler |
| 0x001FB4F0 | `HandleDelete` | Known | Event handler |
| 0x001FB504 | `HandleSelectLozinch` | Known | Event handler |
| 0x001FB7AC | `HandleSelect` | Known | Event handler |
| 0x001FBA78 | `HandleTVOutChanged` | Known | Event handler |
| 0x001FBA90 | `HandleTVSignalChanged` | Known | Event handler |
| 0x001FBAA8 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x001FBAC8 | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x001FBAE8 | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x001FBB0C | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x001FBB2C | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x001FE98C | `HandleSelectKey` | Known | Event handler |
| 0x001FEB34 | `HandleSelect` | Known | Event handler |
| 0x001FF8B0 | `HandlePlayPause` | Known | Event handler |
| 0x001FF8C4 | `HandleWheel` | Known | Event handler |
| 0x001FF8D0 | `HandleWheelRating` | Known | Event handler |
| 0x001FF8E4 | `HandleWheelScrub` | Known | Event handler |
| 0x001FF8F8 | `HandleWheelVolume` | Known | Event handler |
| 0x001FF9B8 | `HandleMenuKey` | Known | Event handler |
| 0x001FFA24 | `HandleMenuLongpress` | Known | Event handler |
| 0x001FFA38 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x00200640 | `HandleSelect` | Known | Event handler |
| 0x00200F38 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00201E50 | `HandleSelect` | Known | Event handler |
| 0x00201E64 | `HandleHilite` | Known | Event handler |
| 0x00201E74 | `HandlePlayPause` | Known | Event handler |
| 0x00201E84 | `HandleAddToOTG` | Known | Event handler |
| 0x00201E94 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00201EB4 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00204F44 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x00205754 | `HandleSelect` | Known | Event handler |
| 0x00205768 | `HandleWheel` | Known | Event handler |
| 0x00205774 | `HandleWheelProgress` | Known | Event handler |
| 0x00205788 | `HandleSelectProgress` | Known | Event handler |
| 0x002057A0 | `HandleSelectVolume` | Known | Event handler |
| 0x002057B4 | `HandleSelectScrub` | Known | Event handler |
| 0x002057C8 | `HandleSelectGenius` | Known | Event handler |
| 0x002057DC | `HandleSelectRating` | Known | Event handler |
| 0x002057F0 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x00205808 | `HandleSelectChapterArt` | Known | Event handler |
| 0x00205820 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x0020583C | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x00205A38 | `HandleWheelGenius` | Known | Event handler |
| 0x00205A4C | `HandleWheelBrightness` | Known | Event handler |
| 0x00205AB8 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x00205AD8 | `HandlePushContextualMenu` | Known | Event handler |
| 0x00205AF4 | `HandleAddToOTG` | Known | Event handler |
| 0x00205B04 | `HandleViewArtist` | Known | Event handler |
| 0x00205B18 | `HandleViewAlbum` | Known | Event handler |
| 0x00205B28 | `HandleViewCompilation` | Known | Event handler |
| 0x00205C20 | `HandleStartGenius` | Known | Event handler |
| 0x00205C34 | `HandleAudiobookSlower` | Known | Event handler |
| 0x00205C4C | `HandleAudiobookFaster` | Known | Event handler |
| 0x00205C64 | `HandleAudiobookNormal` | Known | Event handler |
| 0x00205C7C | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00207730 | `HandleStartGenius` | Known | Event handler |
| 0x00207A84 | `HandleAudiobookSlower` | Known | Event handler |
| 0x00207A9C | `HandleAudiobookNormal` | Known | Event handler |
| 0x00207AB4 | `HandleAudiobookFaster` | Known | Event handler |
| 0x00207ACC | `HandleStartGenius` | Known | Event handler |
| 0x00207AE0 | `HandleAddToOTG` | Known | Event handler |
| 0x00207AF0 | `HandleViewCompilation` | Known | Event handler |
| 0x00207B08 | `HandleViewAlbum` | Known | Event handler |
| 0x00207B18 | `HandleViewArtist` | Known | Event handler |
| 0x00207B2C | `HandleCancel` | Known | Event handler |
| 0x002085D8 | `HandleSelect` | Known | Event handler |
| 0x002085E8 | `HandleSelectRating` | Known | Event handler |
| 0x002085FC | `HandleSelectProgress` | Known | Event handler |
| 0x00208614 | `HandleWheelProgress` | Known | Event handler |
| 0x00208628 | `HandleSelectScrub` | Known | Event handler |
| 0x0020863C | `HandleWheelBrightness` | Known | Event handler |
| 0x00208654 | `HandleNextTrackGotoScrub` | Known | Event handler |
| 0x00208670 | `HandlePrevTrackGotoScrub` | Known | Event handler |
| 0x0020868C | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x0020B2AC | `HandleStartGenius` | Known | Event handler |
| 0x0020B2C4 | `HandleViewArtist` | Known | Event handler |
| 0x0020B2D8 | `HandleViewAlbum` | Known | Event handler |
| 0x0020B2E8 | `HandleViewCompilation` | Known | Event handler |
| 0x0020B300 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0020BCA8 | `HandleStartGenius` | Known | Event handler |
| 0x0020BCBC | `HandleAddToOTG` | Known | Event handler |
| 0x0020BCCC | `HandleViewCompilation` | Known | Event handler |
| 0x0020BCE4 | `HandleViewAlbum` | Known | Event handler |
| 0x0020BCF4 | `HandleViewArtist` | Known | Event handler |
| 0x0020BD08 | `HandleCancel` | Known | Event handler |
| 0x0020E698 | `HandleAddToOTG` | Known | Event handler |
| 0x0020E6A8 | `HandleCancel` | Known | Event handler |
| 0x0020E89C | `HandleStartGenius` | Known | Event handler |
| 0x0020E8B4 | `HandleViewAlbum` | Known | Event handler |
| 0x0020E8C4 | `HandleViewArtist` | Known | Event handler |
| 0x0020E8D8 | `HandleViewCompilation` | Known | Event handler |
| 0x0020E8F0 | `HandleSaveGeniusPlaylist` | Known | Event handler |
| 0x0020E90C | `HandleRefreshPlaylist` | Known | Event handler |
| 0x0020E924 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0020F8F0 | `HandleStartGenius` | Known | Event handler |
| 0x0020F904 | `HandleAddToOTG` | Known | Event handler |
| 0x0020F914 | `HandleViewCompilation` | Known | Event handler |
| 0x0020F92C | `HandleViewAlbum` | Known | Event handler |
| 0x0020F93C | `HandleViewArtist` | Known | Event handler |
| 0x0020F950 | `HandleCancel` | Known | Event handler |
| 0x0020FDFC | `HandleAddToOTG` | Known | Event handler |
| 0x0020FE0C | `HandleCancel` | Known | Event handler |
| 0x002108A4 | `HandleLanguage` | Known | Event handler |
| 0x002108B4 | `HandleResetAllSettings` | Known | Event handler |
| 0x002108CC | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x00211238 | `HandleSelect` | Known | Event handler |
| 0x00211468 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x00212348 | `HandleAddToOTG` | Known | Event handler |
| 0x00212358 | `HandleCancel` | Known | Event handler |
| 0x00214E40 | `HandleSelect` | Known | Event handler |
| 0x00214FDC | `HandleSelect` | Known | Event handler |
| 0x0021527C | `HandleNextDay` | Known | Event handler |
| 0x00215290 | `HandlePreviousDay` | Known | Event handler |
| 0x00215A90 | `HandleMusicHilited` | Known | Event handler |
| 0x00215AA8 | `HandleVideosHilited` | Known | Event handler |
| 0x00215ABC | `HandlePodcastsHilited` | Known | Event handler |
| 0x00215AD4 | `HandleGenericHilited` | Known | Event handler |
| 0x00215AEC | `HandlePhotosHilited` | Known | Event handler |
| 0x00215B00 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x00215B18 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x00215B34 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00215B4C | `HandleArtistsHilited` | Known | Event handler |
| 0x00215B64 | `HandleGenresHilited` | Known | Event handler |
| 0x00215B78 | `HandleAlbumsHilited` | Known | Event handler |
| 0x00215B8C | `HandleCompilationsHilited` | Known | Event handler |
| 0x00215D60 | `HandleComposersHilited` | Known | Event handler |
| 0x00215D78 | `HandleSongsHilited` | Known | Event handler |
| 0x00215D8C | `HandlePlaylistsHilited` | Known | Event handler |
| 0x00215DA4 | `HandleGeniusHilited` | Known | Event handler |
| 0x00215DB8 | `HandleGeniusMixesHilited` | Known | Event handler |
| 0x00215DD4 | `HandleTVShowsHilited` | Known | Event handler |
| 0x00215DEC | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00215E08 | `HandleMoviesHilited` | Known | Event handler |
| 0x00215E1C | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00215E38 | `HandleRentalsHilited` | Known | Event handler |
| 0x00215E50 | `HandleMusicSelected` | Known | Event handler |
| 0x0021601C | `HandleVideosSelected` | Known | Event handler |
| 0x00216034 | `HandlePodcastsSelected` | Known | Event handler |
| 0x0021604C | `HandlePhotosSelected` | Known | Event handler |
| 0x00216064 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x0021607C | `HandleSongsSelected` | Known | Event handler |
| 0x00216090 | `HandleAlbumsSelected` | Known | Event handler |
| 0x002160A8 | `HandleCompilationsSelected` | Known | Event handler |
| 0x002160C4 | `HandleArtistsSelected` | Known | Event handler |
| 0x002160DC | `HandleGenresSelected` | Known | Event handler |
| 0x002160F4 | `HandleComposersSelected` | Known | Event handler |
| 0x0021610C | `HandleAudiobooksSelected` | Known | Event handler |
| 0x002162E0 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x002162FC | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00216314 | `HandleNowPlaying` | Known | Event handler |
| 0x00216328 | `HandleGotoGeniusMixes` | Known | Event handler |
| 0x00216340 | `HandleTVShowsSelected` | Known | Event handler |
| 0x00216358 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00216374 | `HandleMoviesSelected` | Known | Event handler |
| 0x0021638C | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x002163AC | `HandleRentalsSelected` | Known | Event handler |
| 0x002163C4 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x002163DC | `HandleLock` | Known | Event handler |
| 0x00216458 | `HandleBacklightSelected` | Known | Event handler |
| 0x00216470 | `HandleSleepSelected` | Known | Event handler |
| 0x00216484 | `HandleNikePlusSelected` | Known | Event handler |
| 0x00218FE8 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00219584 | `HandleAddToOTG` | Known | Event handler |
| 0x00219594 | `HandleCancel` | Known | Event handler |
| 0x00219764 | `HandleWheel` | Known | Event handler |
| 0x0021A5B0 | `HandleAddToOTG` | Known | Event handler |
| 0x0021A5C0 | `HandleCancel` | Known | Event handler |
| 0x0021B074 | `HandleAddToOTG` | Known | Event handler |
| 0x0021B084 | `HandleCancel` | Known | Event handler |
| 0x0021BA34 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x0021BC8C | `HandleNextDay` | Known | Event handler |
| 0x0021BCA0 | `HandlePreviousDay` | Known | Event handler |
| 0x0021BEE8 | `HandleSelect` | Known | Event handler |
| 0x0021C184 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0021C654 | `HandleAddToOTG` | Known | Event handler |
| 0x0021C664 | `HandleCancel` | Known | Event handler |
| 0x0021D580 | `HandleGeniusMixPlaylistReady` | Known | Event handler |
| 0x0021D5A4 | `HandleSelectMix` | Known | Event handler |
| 0x0021D5B4 | `HandlePlayPause` | Known | Event handler |
| 0x0021D5C4 | `HandlePrev` | Known | Event handler |
| 0x0021D5D0 | `HandleNext` | Known | Event handler |
| 0x0021D5DC | `HandleWheel` | Known | Event handler |
| 0x00220548 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x00220564 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x0022057C | `HandleStartGenius` | Known | Event handler |
| 0x00220590 | `HandleViewArtist` | Known | Event handler |
| 0x002205A4 | `HandleViewAlbum` | Known | Event handler |
| 0x002205B4 | `HandleViewCompilation` | Known | Event handler |
| 0x002205CC | `HandleShowContextualMenu` | Known | Event handler |
| 0x002205E8 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x00220600 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00221970 | `HandleStartGenius` | Known | Event handler |
| 0x00221984 | `HandleAddToOTG` | Known | Event handler |
| 0x00221994 | `HandleViewCompilation` | Known | Event handler |
| 0x002219AC | `HandleViewAlbum` | Known | Event handler |
| 0x002219BC | `HandleViewArtist` | Known | Event handler |
| 0x002219D0 | `HandleCancel` | Known | Event handler |
| 0x00222144 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x002223A8 | `HandleAddToOTG` | Known | Event handler |
| 0x002223B8 | `HandleCancel` | Known | Event handler |
| 0x002228AC | `HandleSelect` | Known | Event handler |
| 0x00222F78 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x0025CAE8 | `HandleDeleteClock` | Known | Event handler |
| 0x0025CB00 | `HandleSelectClock` | Known | Event handler |
| 0x0025CB14 | `HandleHilited` | Known | Event handler |
| 0x0025CB24 | `HandleWheel` | Known | Event handler |
| 0x0025CB30 | `HandleSelectLozinch` | Known | Event handler |
| 0x004178DA | `HandleAudioFFDown` | Known | Event handler |
| 0x00417903 | `HandleAudioFFUp` | Known | Event handler |
| 0x0041792E | `HandleAudioMute` | Known | Event handler |
| 0x00417961 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x00417996 | `HandleAudioNext` | Known | Event handler |
| 0x004179C6 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x004179FD | `HandleAudioNextChapter` | Known | Event handler |
| 0x00417A37 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x00417A6B | `HandleAudioPause` | Known | Event handler |
| 0x00417A97 | `HandleAudioPlay` | Known | Event handler |
| 0x00417AC5 | `HandleAudioPlayPause` | Known | Event handler |
| 0x00417AFD | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x00417B36 | `HandleAudioPrevious` | Known | Event handler |
| 0x00417B6A | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x00417BA1 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x00417BDB | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x00417C10 | `HandleAudioRepeat` | Known | Event handler |
| 0x00417C3C | `HandleAudioRewDown` | Known | Event handler |
| 0x00417C67 | `HandleAudioRewUp` | Known | Event handler |
| 0x00417C96 | `HandleAudioShuffle` | Known | Event handler |
| 0x00417CC4 | `HandleAudioStop` | Known | Event handler |
| 0x00417CF5 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x00417D2A | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x00417D61 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x00417D92 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x00417E4B | `HandleNextPressAndHold` | Known | Event handler |
| 0x00417E7C | `HandleNext` | Known | Event handler |
| 0x00417EB4 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x00417EEF | `HandlePlayPause` | Known | Event handler |
| 0x00417F23 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x00417F58 | `HandlePrevious` | Known | Event handler |
| 0x00417FEA | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x00418032 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x0041807B | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x004180BD | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x004180F5 | `HandleMikeyCenter` | Known | Event handler |
| 0x00418128 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x0041815E | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x00418196 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x004181C8 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x004181FE | `HandleRemoteBacklight` | Known | Event handler |
| 0x00418236 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x00418270 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x004182A9 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x004182DE | `HandleRemoteEvent` | Known | Event handler |
| 0x0041830A | `HandleRemoteFFDown` | Known | Event handler |
| 0x00418335 | `HandleRemoteFFUp` | Known | Event handler |
| 0x00418362 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x00418391 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x004183C0 | `HandleRemoteMute` | Known | Event handler |
| 0x004183F2 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x0041842B | `HandleRemoteNextChapter` | Known | Event handler |
| 0x00418467 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x004184A7 | `HandleRemoteOff` | Known | Event handler |
| 0x004184D0 | `HandleRemoteOff` | Known | Event handler |
| 0x004184FA | `HandleRemoteOn` | Known | Event handler |
| 0x00418526 | `HandleRemotePause` | Known | Event handler |
| 0x00418554 | `HandleRemotePlay` | Known | Event handler |
| 0x00418592 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x004185D3 | `HandleRemotePlayPause` | Known | Event handler |
| 0x0041860A | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x00418643 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x0041867F | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x004186B6 | `HandleRemoteRepeat` | Known | Event handler |
| 0x004186E4 | `HandleRemoteRewDown` | Known | Event handler |
| 0x00418711 | `HandleRemoteRewUp` | Known | Event handler |
| 0x00418741 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x00418774 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x004187A8 | `HandleRemoteShuffle` | Known | Event handler |
| 0x004187D8 | `HandleRemoteStop` | Known | Event handler |
| 0x00418808 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x0041883D | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x00418875 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x004188AC | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x004188E5 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x00418918 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x0041894D | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x00418980 | `HandleVideoFFDown` | Known | Event handler |
| 0x004189A9 | `HandleVideoFFUp` | Known | Event handler |
| 0x004189DC | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x00418A11 | `HandleVideoNext` | Known | Event handler |
| 0x00418A43 | `HandleVideoNextChapter` | Known | Event handler |
| 0x00418A7A | `HandleVideoNextFrame` | Known | Event handler |
| 0x00418AAB | `HandleVideoPause` | Known | Event handler |
| 0x00418AD7 | `HandleVideoPlay` | Known | Event handler |
| 0x00418B05 | `HandleVideoPlayPause` | Known | Event handler |
| 0x00418B3D | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x00418B76 | `HandleVideoPrevious` | Known | Event handler |
| 0x00418BAC | `HandleVideoPrevChapter` | Known | Event handler |
| 0x00418BE3 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x00418C12 | `HandleVideoRewDown` | Known | Event handler |
| 0x00418C3D | `HandleVideoRewUp` | Known | Event handler |
| 0x00418C69 | `HandleVideoStop` | Known | Event handler |
| 0x007207EE | `HandleAddressBook` | Known | Event handler |
| 0x00720D8A | `HandleSelect` | Known | Event handler |
| 0x00720DC5 | `HandleHilite` | Known | Event handler |
| 0x00720E46 | `HandleSelectRegion` | Known | Event handler |
| 0x00720EE6 | `HandleSelectRegion` | Known | Event handler |
| 0x00720F82 | `HandleSelectRegion` | Known | Event handler |
| 0x00721026 | `HandleSelectRegion` | Known | Event handler |
| 0x007210CC | `HandleSelectRegion` | Known | Event handler |
| 0x0072116C | `HandleSelectRegion` | Known | Event handler |
| 0x00721218 | `HandleSelectRegion` | Known | Event handler |
| 0x007212BA | `HandleSelectRegion` | Known | Event handler |
| 0x0072136A | `HandleSelectCity` | Known | Event handler |
| 0x007213D6 | `HandleHighlightCity` | Known | Event handler |
| 0x0072140F | `HandleSelectCity` | Known | Event handler |
| 0x0072147B | `HandleHighlightCity` | Known | Event handler |
| 0x007214B4 | `HandleSelectCity` | Known | Event handler |
| 0x00721520 | `HandleHighlightCity` | Known | Event handler |
| 0x00721559 | `HandleSelectCity` | Known | Event handler |
| 0x007215C5 | `HandleHighlightCity` | Known | Event handler |
| 0x007215FE | `HandleSelectCity` | Known | Event handler |
| 0x0072166A | `HandleHighlightCity` | Known | Event handler |
| 0x007216A3 | `HandleSelectCity` | Known | Event handler |
| 0x0072170F | `HandleHighlightCity` | Known | Event handler |
| 0x00721748 | `HandleSelectCity` | Known | Event handler |
| 0x007217B4 | `HandleHighlightCity` | Known | Event handler |
| 0x007217ED | `HandleSelectCity` | Known | Event handler |
| 0x00721859 | `HandleHighlightCity` | Known | Event handler |
| 0x00721892 | `HandleSelectCity` | Known | Event handler |
| 0x007218FE | `HandleHighlightCity` | Known | Event handler |
| 0x00721937 | `HandleSelectCity` | Known | Event handler |
| 0x007219A3 | `HandleHighlightCity` | Known | Event handler |
| 0x007219DC | `HandleSelectCity` | Known | Event handler |
| 0x00721A48 | `HandleHighlightCity` | Known | Event handler |
| 0x00721A81 | `HandleSelectCity` | Known | Event handler |
| 0x00721AED | `HandleHighlightCity` | Known | Event handler |
| 0x00721B26 | `HandleSelectCity` | Known | Event handler |
| 0x00721B92 | `HandleHighlightCity` | Known | Event handler |
| 0x00721BCB | `HandleSelectCity` | Known | Event handler |
| 0x00721C37 | `HandleHighlightCity` | Known | Event handler |
| 0x00721C70 | `HandleSelectCity` | Known | Event handler |
| 0x00721CDC | `HandleHighlightCity` | Known | Event handler |
| 0x00721D15 | `HandleSelectCity` | Known | Event handler |
| 0x00721D81 | `HandleHighlightCity` | Known | Event handler |
| 0x00721DBA | `HandleSelectCity` | Known | Event handler |
| 0x00721E26 | `HandleHighlightCity` | Known | Event handler |
| 0x00721E5F | `HandleSelectCity` | Known | Event handler |
| 0x00721ECB | `HandleHighlightCity` | Known | Event handler |
| 0x00721F04 | `HandleSelectCity` | Known | Event handler |
| 0x00721F70 | `HandleHighlightCity` | Known | Event handler |
| 0x00721FA9 | `HandleSelectCity` | Known | Event handler |
| 0x00722015 | `HandleHighlightCity` | Known | Event handler |
| 0x0072204E | `HandleSelectCity` | Known | Event handler |
| 0x007220BA | `HandleHighlightCity` | Known | Event handler |
| 0x007220F3 | `HandleSelectCity` | Known | Event handler |
| 0x0072215F | `HandleHighlightCity` | Known | Event handler |
| 0x00722198 | `HandleSelectCity` | Known | Event handler |
| 0x00722204 | `HandleHighlightCity` | Known | Event handler |
| 0x0072223D | `HandleSelectCity` | Known | Event handler |
| 0x007222A9 | `HandleHighlightCity` | Known | Event handler |
| 0x007222E2 | `HandleSelectCity` | Known | Event handler |
| 0x0072234E | `HandleHighlightCity` | Known | Event handler |
| 0x00722387 | `HandleSelectCity` | Known | Event handler |
| 0x007223F3 | `HandleHighlightCity` | Known | Event handler |
| 0x0072242C | `HandleSelectCity` | Known | Event handler |
| 0x00722498 | `HandleHighlightCity` | Known | Event handler |
| 0x007224D1 | `HandleSelectCity` | Known | Event handler |
| 0x0072253D | `HandleHighlightCity` | Known | Event handler |
| 0x00722576 | `HandleSelectCity` | Known | Event handler |
| 0x007225E2 | `HandleHighlightCity` | Known | Event handler |
| 0x0072261B | `HandleSelectCity` | Known | Event handler |
| 0x00722687 | `HandleHighlightCity` | Known | Event handler |
| 0x007226C0 | `HandleSelectCity` | Known | Event handler |
| 0x0072272C | `HandleHighlightCity` | Known | Event handler |
| 0x0072276A | `HandleSelectCity` | Known | Event handler |
| 0x007227D6 | `HandleHighlightCity` | Known | Event handler |
| 0x0072280F | `HandleSelectCity` | Known | Event handler |
| 0x0072287B | `HandleHighlightCity` | Known | Event handler |
| 0x007228B4 | `HandleSelectCity` | Known | Event handler |
| 0x00722920 | `HandleHighlightCity` | Known | Event handler |
| 0x00722959 | `HandleSelectCity` | Known | Event handler |
| 0x007229C5 | `HandleHighlightCity` | Known | Event handler |
| 0x007229FE | `HandleSelectCity` | Known | Event handler |
| 0x00722A6A | `HandleHighlightCity` | Known | Event handler |
| 0x00722AA3 | `HandleSelectCity` | Known | Event handler |
| 0x00722B0F | `HandleHighlightCity` | Known | Event handler |
| 0x00722B48 | `HandleSelectCity` | Known | Event handler |
| 0x00722BB4 | `HandleHighlightCity` | Known | Event handler |
| 0x00722BED | `HandleSelectCity` | Known | Event handler |
| 0x00722C59 | `HandleHighlightCity` | Known | Event handler |
| 0x00722C92 | `HandleSelectCity` | Known | Event handler |
| 0x00722CFE | `HandleHighlightCity` | Known | Event handler |
| 0x00722D37 | `HandleSelectCity` | Known | Event handler |
| 0x00722DA3 | `HandleHighlightCity` | Known | Event handler |
| 0x00722DDC | `HandleSelectCity` | Known | Event handler |
| 0x00722E48 | `HandleHighlightCity` | Known | Event handler |
| 0x00722E81 | `HandleSelectCity` | Known | Event handler |
| 0x00722EED | `HandleHighlightCity` | Known | Event handler |
| 0x00722F26 | `HandleSelectCity` | Known | Event handler |
| 0x00722F92 | `HandleHighlightCity` | Known | Event handler |
| 0x00722FCB | `HandleSelectCity` | Known | Event handler |
| 0x00723037 | `HandleHighlightCity` | Known | Event handler |
| 0x00723070 | `HandleSelectCity` | Known | Event handler |
| 0x007230DC | `HandleHighlightCity` | Known | Event handler |
| 0x00723115 | `HandleSelectCity` | Known | Event handler |
| 0x00723181 | `HandleHighlightCity` | Known | Event handler |
| 0x007231BA | `HandleSelectCity` | Known | Event handler |
| 0x00723226 | `HandleHighlightCity` | Known | Event handler |
| 0x0072325F | `HandleSelectCity` | Known | Event handler |
| 0x007232CB | `HandleHighlightCity` | Known | Event handler |
| 0x00723304 | `HandleSelectCity` | Known | Event handler |
| 0x00723370 | `HandleHighlightCity` | Known | Event handler |
| 0x007233A9 | `HandleSelectCity` | Known | Event handler |
| 0x00723415 | `HandleHighlightCity` | Known | Event handler |
| 0x0072344E | `HandleSelectCity` | Known | Event handler |
| 0x007234BA | `HandleHighlightCity` | Known | Event handler |
| 0x007234F3 | `HandleSelectCity` | Known | Event handler |
| 0x0072355F | `HandleHighlightCity` | Known | Event handler |
| 0x00723598 | `HandleSelectCity` | Known | Event handler |
| 0x00723604 | `HandleHighlightCity` | Known | Event handler |
| 0x0072363D | `HandleSelectCity` | Known | Event handler |
| 0x007236A9 | `HandleHighlightCity` | Known | Event handler |
| 0x007236E2 | `HandleSelectCity` | Known | Event handler |
| 0x0072374E | `HandleHighlightCity` | Known | Event handler |
| 0x00723787 | `HandleSelectCity` | Known | Event handler |
| 0x007237F3 | `HandleHighlightCity` | Known | Event handler |
| 0x0072382C | `HandleSelectCity` | Known | Event handler |
| 0x00723898 | `HandleHighlightCity` | Known | Event handler |
| 0x007238D1 | `HandleSelectCity` | Known | Event handler |
| 0x0072393D | `HandleHighlightCity` | Known | Event handler |
| 0x00723976 | `HandleSelectCity` | Known | Event handler |
| 0x007239E2 | `HandleHighlightCity` | Known | Event handler |
| 0x00723A1B | `HandleSelectCity` | Known | Event handler |
| 0x00723A87 | `HandleHighlightCity` | Known | Event handler |
| 0x00723AC0 | `HandleSelectCity` | Known | Event handler |
| 0x00723B2C | `HandleHighlightCity` | Known | Event handler |
| 0x00723B65 | `HandleSelectCity` | Known | Event handler |
| 0x00723BD1 | `HandleHighlightCity` | Known | Event handler |
| 0x00723C0A | `HandleSelectCity` | Known | Event handler |
| 0x00723C76 | `HandleHighlightCity` | Known | Event handler |
| 0x00723CAF | `HandleSelectCity` | Known | Event handler |
| 0x00723D1B | `HandleHighlightCity` | Known | Event handler |
| 0x00723D54 | `HandleSelectCity` | Known | Event handler |
| 0x00723DC0 | `HandleHighlightCity` | Known | Event handler |
| 0x00723DF9 | `HandleSelectCity` | Known | Event handler |
| 0x00723E65 | `HandleHighlightCity` | Known | Event handler |
| 0x00723E9E | `HandleSelectCity` | Known | Event handler |
| 0x00723F0A | `HandleHighlightCity` | Known | Event handler |
| 0x00723F43 | `HandleSelectCity` | Known | Event handler |
| 0x00723FAF | `HandleHighlightCity` | Known | Event handler |
| 0x00723FE8 | `HandleSelectCity` | Known | Event handler |
| 0x00724054 | `HandleHighlightCity` | Known | Event handler |
| 0x0072408D | `HandleSelectCity` | Known | Event handler |
| 0x007240F9 | `HandleHighlightCity` | Known | Event handler |
| 0x00724132 | `HandleSelectCity` | Known | Event handler |
| 0x0072419E | `HandleHighlightCity` | Known | Event handler |
| 0x007241D7 | `HandleSelectCity` | Known | Event handler |
| 0x00724243 | `HandleHighlightCity` | Known | Event handler |
| 0x0072427C | `HandleSelectCity` | Known | Event handler |
| 0x007242E8 | `HandleHighlightCity` | Known | Event handler |
| 0x00724321 | `HandleSelectCity` | Known | Event handler |
| 0x0072438D | `HandleHighlightCity` | Known | Event handler |
| 0x007243C6 | `HandleSelectCity` | Known | Event handler |
| 0x00724432 | `HandleHighlightCity` | Known | Event handler |
| 0x0072446B | `HandleSelectCity` | Known | Event handler |
| 0x007244D7 | `HandleHighlightCity` | Known | Event handler |
| 0x00724510 | `HandleSelectCity` | Known | Event handler |
| 0x0072457C | `HandleHighlightCity` | Known | Event handler |
| 0x007245B5 | `HandleSelectCity` | Known | Event handler |
| 0x00724621 | `HandleHighlightCity` | Known | Event handler |
| 0x0072465A | `HandleSelectCity` | Known | Event handler |
| 0x007246C6 | `HandleHighlightCity` | Known | Event handler |
| 0x007246FF | `HandleSelectCity` | Known | Event handler |
| 0x0072476B | `HandleHighlightCity` | Known | Event handler |
| 0x007247A4 | `HandleSelectCity` | Known | Event handler |
| 0x00724810 | `HandleHighlightCity` | Known | Event handler |
| 0x00724849 | `HandleSelectCity` | Known | Event handler |
| 0x007248B5 | `HandleHighlightCity` | Known | Event handler |
| 0x007248EE | `HandleSelectCity` | Known | Event handler |
| 0x0072495A | `HandleHighlightCity` | Known | Event handler |
| 0x00724993 | `HandleSelectCity` | Known | Event handler |
| 0x007249FF | `HandleHighlightCity` | Known | Event handler |
| 0x00724A38 | `HandleSelectCity` | Known | Event handler |
| 0x00724AA4 | `HandleHighlightCity` | Known | Event handler |
| 0x00724ADD | `HandleSelectCity` | Known | Event handler |
| 0x00724B49 | `HandleHighlightCity` | Known | Event handler |
| 0x00724B82 | `HandleSelectCity` | Known | Event handler |
| 0x00724BEE | `HandleHighlightCity` | Known | Event handler |
| 0x00724C2E | `HandleSelectCity` | Known | Event handler |
| 0x00724C9A | `HandleHighlightCity` | Known | Event handler |
| 0x00724CD3 | `HandleSelectCity` | Known | Event handler |
| 0x00724D3F | `HandleHighlightCity` | Known | Event handler |
| 0x00724D78 | `HandleSelectCity` | Known | Event handler |
| 0x00724DE4 | `HandleHighlightCity` | Known | Event handler |
| 0x00724E22 | `HandleSelectCity` | Known | Event handler |
| 0x00724E8E | `HandleHighlightCity` | Known | Event handler |
| 0x00724EC7 | `HandleSelectCity` | Known | Event handler |
| 0x00724F33 | `HandleHighlightCity` | Known | Event handler |
| 0x00724F6C | `HandleSelectCity` | Known | Event handler |
| 0x00724FD8 | `HandleHighlightCity` | Known | Event handler |
| 0x00725011 | `HandleSelectCity` | Known | Event handler |
| 0x0072507D | `HandleHighlightCity` | Known | Event handler |
| 0x007250B6 | `HandleSelectCity` | Known | Event handler |
| 0x00725122 | `HandleHighlightCity` | Known | Event handler |
| 0x0072515B | `HandleSelectCity` | Known | Event handler |
| 0x007251C7 | `HandleHighlightCity` | Known | Event handler |
| 0x00725200 | `HandleSelectCity` | Known | Event handler |
| 0x0072526C | `HandleHighlightCity` | Known | Event handler |
| 0x007252A5 | `HandleSelectCity` | Known | Event handler |
| 0x00725311 | `HandleHighlightCity` | Known | Event handler |
| 0x0072534E | `HandleSelectCity` | Known | Event handler |
| 0x007253BA | `HandleHighlightCity` | Known | Event handler |
| 0x007253F3 | `HandleSelectCity` | Known | Event handler |
| 0x0072545F | `HandleHighlightCity` | Known | Event handler |
| 0x00725498 | `HandleSelectCity` | Known | Event handler |
| 0x00725504 | `HandleHighlightCity` | Known | Event handler |
| 0x0072553D | `HandleSelectCity` | Known | Event handler |
| 0x007255A9 | `HandleHighlightCity` | Known | Event handler |
| 0x007255E2 | `HandleSelectCity` | Known | Event handler |
| 0x0072564E | `HandleHighlightCity` | Known | Event handler |
| 0x00725687 | `HandleSelectCity` | Known | Event handler |
| 0x007256F3 | `HandleHighlightCity` | Known | Event handler |
| 0x0072572C | `HandleSelectCity` | Known | Event handler |
| 0x00725798 | `HandleHighlightCity` | Known | Event handler |
| 0x007257D1 | `HandleSelectCity` | Known | Event handler |
| 0x0072583D | `HandleHighlightCity` | Known | Event handler |
| 0x00725876 | `HandleSelectCity` | Known | Event handler |
| 0x007258E2 | `HandleHighlightCity` | Known | Event handler |
| 0x0072591B | `HandleSelectCity` | Known | Event handler |
| 0x00725987 | `HandleHighlightCity` | Known | Event handler |
| 0x007259C0 | `HandleSelectCity` | Known | Event handler |
| 0x00725A2C | `HandleHighlightCity` | Known | Event handler |
| 0x00725A65 | `HandleSelectCity` | Known | Event handler |
| 0x00725AD1 | `HandleHighlightCity` | Known | Event handler |
| 0x00725B0A | `HandleSelectCity` | Known | Event handler |
| 0x00725B76 | `HandleHighlightCity` | Known | Event handler |
| 0x00725BAF | `HandleSelectCity` | Known | Event handler |
| 0x00725C1B | `HandleHighlightCity` | Known | Event handler |
| 0x00725C54 | `HandleSelectCity` | Known | Event handler |
| 0x00725CC0 | `HandleHighlightCity` | Known | Event handler |
| 0x00725CF9 | `HandleSelectCity` | Known | Event handler |
| 0x00725D65 | `HandleHighlightCity` | Known | Event handler |
| 0x00725D9E | `HandleSelectCity` | Known | Event handler |
| 0x00725E0A | `HandleHighlightCity` | Known | Event handler |
| 0x00725E43 | `HandleSelectCity` | Known | Event handler |
| 0x00725EAF | `HandleHighlightCity` | Known | Event handler |
| 0x00725EE8 | `HandleSelectCity` | Known | Event handler |
| 0x00725F54 | `HandleHighlightCity` | Known | Event handler |
| 0x00725F8D | `HandleSelectCity` | Known | Event handler |
| 0x00725FF9 | `HandleHighlightCity` | Known | Event handler |
| 0x00726032 | `HandleSelectCity` | Known | Event handler |
| 0x0072609E | `HandleHighlightCity` | Known | Event handler |
| 0x007260D7 | `HandleSelectCity` | Known | Event handler |
| 0x00726143 | `HandleHighlightCity` | Known | Event handler |
| 0x0072617C | `HandleSelectCity` | Known | Event handler |
| 0x007261E8 | `HandleHighlightCity` | Known | Event handler |
| 0x00726221 | `HandleSelectCity` | Known | Event handler |
| 0x0072628D | `HandleHighlightCity` | Known | Event handler |
| 0x007262C6 | `HandleSelectCity` | Known | Event handler |
| 0x00726332 | `HandleHighlightCity` | Known | Event handler |
| 0x0072636B | `HandleSelectCity` | Known | Event handler |
| 0x007263D7 | `HandleHighlightCity` | Known | Event handler |
| 0x00726410 | `HandleSelectCity` | Known | Event handler |
| 0x0072647C | `HandleHighlightCity` | Known | Event handler |
| 0x007264B5 | `HandleSelectCity` | Known | Event handler |
| 0x00726521 | `HandleHighlightCity` | Known | Event handler |
| 0x0072655A | `HandleSelectCity` | Known | Event handler |
| 0x007265C6 | `HandleHighlightCity` | Known | Event handler |
| 0x007265FF | `HandleSelectCity` | Known | Event handler |
| 0x0072666B | `HandleHighlightCity` | Known | Event handler |
| 0x007266A4 | `HandleSelectCity` | Known | Event handler |
| 0x00726710 | `HandleHighlightCity` | Known | Event handler |
| 0x00726749 | `HandleSelectCity` | Known | Event handler |
| 0x007267B5 | `HandleHighlightCity` | Known | Event handler |
| 0x007267EE | `HandleSelectCity` | Known | Event handler |
| 0x0072685A | `HandleHighlightCity` | Known | Event handler |
| 0x00726893 | `HandleSelectCity` | Known | Event handler |
| 0x007268FF | `HandleHighlightCity` | Known | Event handler |
| 0x0072693E | `HandleSelectCity` | Known | Event handler |
| 0x007269AA | `HandleHighlightCity` | Known | Event handler |
| 0x007269E3 | `HandleSelectCity` | Known | Event handler |
| 0x00726A4F | `HandleHighlightCity` | Known | Event handler |
| 0x00726A88 | `HandleSelectCity` | Known | Event handler |
| 0x00726AF4 | `HandleHighlightCity` | Known | Event handler |
| 0x00726B2D | `HandleSelectCity` | Known | Event handler |
| 0x00726B99 | `HandleHighlightCity` | Known | Event handler |
| 0x00726BD2 | `HandleSelectCity` | Known | Event handler |
| 0x00726C3E | `HandleHighlightCity` | Known | Event handler |
| 0x00726C77 | `HandleSelectCity` | Known | Event handler |
| 0x00726CE3 | `HandleHighlightCity` | Known | Event handler |
| 0x00726D1C | `HandleSelectCity` | Known | Event handler |
| 0x00726D88 | `HandleHighlightCity` | Known | Event handler |
| 0x00726DC1 | `HandleSelectCity` | Known | Event handler |
| 0x00726E2D | `HandleHighlightCity` | Known | Event handler |
| 0x00726E66 | `HandleSelectCity` | Known | Event handler |
| 0x00726ED2 | `HandleHighlightCity` | Known | Event handler |
| 0x00726F0B | `HandleSelectCity` | Known | Event handler |
| 0x00726F77 | `HandleHighlightCity` | Known | Event handler |
| 0x00726FB0 | `HandleSelectCity` | Known | Event handler |
| 0x0072701C | `HandleHighlightCity` | Known | Event handler |
| 0x00727055 | `HandleSelectCity` | Known | Event handler |
| 0x007270C1 | `HandleHighlightCity` | Known | Event handler |
| 0x007270FA | `HandleSelectCity` | Known | Event handler |
| 0x00727166 | `HandleHighlightCity` | Known | Event handler |
| 0x0072719F | `HandleSelectCity` | Known | Event handler |
| 0x0072720B | `HandleHighlightCity` | Known | Event handler |
| 0x00727244 | `HandleSelectCity` | Known | Event handler |
| 0x007272B0 | `HandleHighlightCity` | Known | Event handler |
| 0x007272E9 | `HandleSelectCity` | Known | Event handler |
| 0x00727355 | `HandleHighlightCity` | Known | Event handler |
| 0x0072738E | `HandleSelectCity` | Known | Event handler |
| 0x007273FA | `HandleHighlightCity` | Known | Event handler |
| 0x00727433 | `HandleSelectCity` | Known | Event handler |
| 0x0072749F | `HandleHighlightCity` | Known | Event handler |
| 0x007274D8 | `HandleSelectCity` | Known | Event handler |
| 0x00727544 | `HandleHighlightCity` | Known | Event handler |
| 0x0072757D | `HandleSelectCity` | Known | Event handler |
| 0x007275E9 | `HandleHighlightCity` | Known | Event handler |
| 0x00727622 | `HandleSelectCity` | Known | Event handler |
| 0x0072768E | `HandleHighlightCity` | Known | Event handler |
| 0x007276C7 | `HandleSelectCity` | Known | Event handler |
| 0x00727733 | `HandleHighlightCity` | Known | Event handler |
| 0x0072776C | `HandleSelectCity` | Known | Event handler |
| 0x007277D8 | `HandleHighlightCity` | Known | Event handler |
| 0x00727811 | `HandleSelectCity` | Known | Event handler |
| 0x0072787D | `HandleHighlightCity` | Known | Event handler |
| 0x007278B6 | `HandleSelectCity` | Known | Event handler |
| 0x00727922 | `HandleHighlightCity` | Known | Event handler |
| 0x0072795B | `HandleSelectCity` | Known | Event handler |
| 0x007279C7 | `HandleHighlightCity` | Known | Event handler |
| 0x00727A00 | `HandleSelectCity` | Known | Event handler |
| 0x00727A6C | `HandleHighlightCity` | Known | Event handler |
| 0x00727AA5 | `HandleSelectCity` | Known | Event handler |
| 0x00727B11 | `HandleHighlightCity` | Known | Event handler |
| 0x00727B4A | `HandleSelectCity` | Known | Event handler |
| 0x00727BB6 | `HandleHighlightCity` | Known | Event handler |
| 0x00727BEF | `HandleSelectCity` | Known | Event handler |
| 0x00727C5B | `HandleHighlightCity` | Known | Event handler |
| 0x00727C94 | `HandleSelectCity` | Known | Event handler |
| 0x00727D00 | `HandleHighlightCity` | Known | Event handler |
| 0x00727D39 | `HandleSelectCity` | Known | Event handler |
| 0x00727DA5 | `HandleHighlightCity` | Known | Event handler |
| 0x00727DDE | `HandleSelectCity` | Known | Event handler |
| 0x00727E4A | `HandleHighlightCity` | Known | Event handler |
| 0x00727E83 | `HandleSelectCity` | Known | Event handler |
| 0x00727EEF | `HandleHighlightCity` | Known | Event handler |
| 0x00727F28 | `HandleSelectCity` | Known | Event handler |
| 0x00727F94 | `HandleHighlightCity` | Known | Event handler |
| 0x00727FCD | `HandleSelectCity` | Known | Event handler |
| 0x00728039 | `HandleHighlightCity` | Known | Event handler |
| 0x00728072 | `HandleSelectCity` | Known | Event handler |
| 0x007280DE | `HandleHighlightCity` | Known | Event handler |
| 0x00728117 | `HandleSelectCity` | Known | Event handler |
| 0x00728183 | `HandleHighlightCity` | Known | Event handler |
| 0x007281BC | `HandleSelectCity` | Known | Event handler |
| 0x00728228 | `HandleHighlightCity` | Known | Event handler |
| 0x00728261 | `HandleSelectCity` | Known | Event handler |
| 0x007282CD | `HandleHighlightCity` | Known | Event handler |
| 0x00728306 | `HandleSelectCity` | Known | Event handler |
| 0x00728372 | `HandleHighlightCity` | Known | Event handler |
| 0x007283AB | `HandleSelectCity` | Known | Event handler |
| 0x00728417 | `HandleHighlightCity` | Known | Event handler |
| 0x00728450 | `HandleSelectCity` | Known | Event handler |
| 0x007284BC | `HandleHighlightCity` | Known | Event handler |
| 0x007284F5 | `HandleSelectCity` | Known | Event handler |
| 0x00728561 | `HandleHighlightCity` | Known | Event handler |
| 0x0072859A | `HandleSelectCity` | Known | Event handler |
| 0x00728606 | `HandleHighlightCity` | Known | Event handler |
| 0x0072863F | `HandleSelectCity` | Known | Event handler |
| 0x007286AB | `HandleHighlightCity` | Known | Event handler |
| 0x007286E4 | `HandleSelectCity` | Known | Event handler |
| 0x00728750 | `HandleHighlightCity` | Known | Event handler |
| 0x00728789 | `HandleSelectCity` | Known | Event handler |
| 0x007287F5 | `HandleHighlightCity` | Known | Event handler |
| 0x0072882E | `HandleSelectCity` | Known | Event handler |
| 0x0072889A | `HandleHighlightCity` | Known | Event handler |
| 0x007288D3 | `HandleSelectCity` | Known | Event handler |
| 0x0072893F | `HandleHighlightCity` | Known | Event handler |
| 0x0072897E | `HandleSelectCity` | Known | Event handler |
| 0x007289EA | `HandleHighlightCity` | Known | Event handler |
| 0x00728A23 | `HandleSelectCity` | Known | Event handler |
| 0x00728A8F | `HandleHighlightCity` | Known | Event handler |
| 0x00728AC8 | `HandleSelectCity` | Known | Event handler |
| 0x00728B34 | `HandleHighlightCity` | Known | Event handler |
| 0x00728B6D | `HandleSelectCity` | Known | Event handler |
| 0x00728BD9 | `HandleHighlightCity` | Known | Event handler |
| 0x00728C12 | `HandleSelectCity` | Known | Event handler |
| 0x00728C7E | `HandleHighlightCity` | Known | Event handler |
| 0x00728CBE | `HandleSelectCity` | Known | Event handler |
| 0x00728D2A | `HandleHighlightCity` | Known | Event handler |
| 0x00728D63 | `HandleSelectCity` | Known | Event handler |
| 0x00728DCF | `HandleHighlightCity` | Known | Event handler |
| 0x00728E08 | `HandleSelectCity` | Known | Event handler |
| 0x00728E74 | `HandleHighlightCity` | Known | Event handler |
| 0x00728EAD | `HandleSelectCity` | Known | Event handler |
| 0x00728F19 | `HandleHighlightCity` | Known | Event handler |
| 0x00728F52 | `HandleSelectCity` | Known | Event handler |
| 0x00728FBE | `HandleHighlightCity` | Known | Event handler |
| 0x00728FF7 | `HandleSelectCity` | Known | Event handler |
| 0x00729063 | `HandleHighlightCity` | Known | Event handler |
| 0x0072909C | `HandleSelectCity` | Known | Event handler |
| 0x00729108 | `HandleHighlightCity` | Known | Event handler |
| 0x00729141 | `HandleSelectCity` | Known | Event handler |
| 0x007291AD | `HandleHighlightCity` | Known | Event handler |
| 0x007291E6 | `HandleSelectCity` | Known | Event handler |
| 0x00729252 | `HandleHighlightCity` | Known | Event handler |
| 0x0072928B | `HandleSelectCity` | Known | Event handler |
| 0x007292F7 | `HandleHighlightCity` | Known | Event handler |
| 0x00729330 | `HandleSelectCity` | Known | Event handler |
| 0x0072939C | `HandleHighlightCity` | Known | Event handler |
| 0x007293D5 | `HandleSelectCity` | Known | Event handler |
| 0x00729441 | `HandleHighlightCity` | Known | Event handler |
| 0x0072947A | `HandleSelectCity` | Known | Event handler |
| 0x007294E6 | `HandleHighlightCity` | Known | Event handler |
| 0x0072951F | `HandleSelectCity` | Known | Event handler |
| 0x0072958B | `HandleHighlightCity` | Known | Event handler |
| 0x007295C4 | `HandleSelectCity` | Known | Event handler |
| 0x00729630 | `HandleHighlightCity` | Known | Event handler |
| 0x00729669 | `HandleSelectCity` | Known | Event handler |
| 0x007296D5 | `HandleHighlightCity` | Known | Event handler |
| 0x0072970E | `HandleSelectCity` | Known | Event handler |
| 0x0072977A | `HandleHighlightCity` | Known | Event handler |
| 0x00729C72 | `HandleMusicSelected` | Known | Event handler |
| 0x00729CB4 | `HandleMusicHilited` | Known | Event handler |
| 0x00729CEC | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00729D32 | `HandleMusicHilited` | Known | Event handler |
| 0x00729D6A | `HandleGotoGeniusMixes` | Known | Event handler |
| 0x00729DAE | `HandleGeniusMixesHilited` | Known | Event handler |
| 0x00729DEC | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00729E32 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x00729E6E | `HandleArtistsSelected` | Known | Event handler |
| 0x00729EB2 | `HandleArtistsHilited` | Known | Event handler |
| 0x00729EEC | `HandleAlbumsSelected` | Known | Event handler |
| 0x00729F2F | `HandleAlbumsHilited` | Known | Event handler |
| 0x00729F68 | `HandleCompilationsSelected` | Known | Event handler |
| 0x00729FB1 | `HandleCompilationsHilited` | Known | Event handler |
| 0x00729FF0 | `HandleSongsSelected` | Known | Event handler |
| 0x0072A032 | `HandleSongsHilited` | Known | Event handler |
| 0x0072A06A | `HandleGenresSelected` | Known | Event handler |
| 0x0072A0AD | `HandleGenresHilited` | Known | Event handler |
| 0x0072A0E6 | `HandleComposersSelected` | Known | Event handler |
| 0x0072A12C | `HandleComposersHilited` | Known | Event handler |
| 0x0072A168 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x0072A1AF | `HandleAudiobooksHilited` | Known | Event handler |
| 0x0072A26E | `HandleMusicHilited` | Known | Event handler |
| 0x0072A2A6 | `HandleVideosSelected` | Known | Event handler |
| 0x0072A2E9 | `HandleVideosHilited` | Known | Event handler |
| 0x0072A322 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x0072A36D | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x0072A3AE | `HandleMoviesSelected` | Known | Event handler |
| 0x0072A3F1 | `HandleMoviesHilited` | Known | Event handler |
| 0x0072A42A | `HandleTVShowsSelected` | Known | Event handler |
| 0x0072A46E | `HandleTVShowsHilited` | Known | Event handler |
| 0x0072A4A8 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x0072A4F0 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x0072A52E | `HandleRentalsSelected` | Known | Event handler |
| 0x0072A572 | `HandleRentalsHilited` | Known | Event handler |
| 0x0072A5AC | `HandlePhotosSelected` | Known | Event handler |
| 0x0072A5EF | `HandlePhotosHilited` | Known | Event handler |
| 0x0072A628 | `HandlePhotosSelected` | Known | Event handler |
| 0x0072A66B | `HandlePhotosHilited` | Known | Event handler |
| 0x0072A6A4 | `HandlePodcastsSelected` | Known | Event handler |
| 0x0072A6E9 | `HandlePodcastsHilited` | Known | Event handler |
| 0x0072A79C | `HandleGenericHilited` | Known | Event handler |
| 0x0072A895 | `HandleGenericHilited` | Known | Event handler |
| 0x0072AD7A | `HandleLock` | Known | Event handler |
| 0x0072AEEB | `HandleNikePlusSelected` | Known | Event handler |
| 0x0072AF30 | `HandleGenericHilited` | Known | Event handler |
| 0x0072B036 | `HandleGenericHilited` | Known | Event handler |
| 0x0072B135 | `HandleGenericHilited` | Known | Event handler |
| 0x0072B222 | `HandleGenericHilited` | Known | Event handler |
| 0x0072B31F | `HandleGenericHilited` | Known | Event handler |
| 0x0072B399 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x0072B3E2 | `HandleGenericHilited` | Known | Event handler |
| 0x0072B45B | `HandleBacklightSelected` | Known | Event handler |
| 0x0072B4A1 | `HandleGenericHilited` | Known | Event handler |
| 0x0072B51C | `HandleSleepSelected` | Known | Event handler |
| 0x0072B55E | `HandleGenericHilited` | Known | Event handler |
| 0x0072B5D5 | `HandleNowPlaying` | Known | Event handler |
| 0x0072B64D | `HandleNowPlayingHilited` | Known | Event handler |
| 0x0072B68E | `HandleCoverFlowSelected` | Known | Event handler |
| 0x0072B6D4 | `HandleMusicHilited` | Known | Event handler |
| 0x0072B70C | `HandleGotoGeniusMixes` | Known | Event handler |
| 0x0072B747 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x0072B78D | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x0072B7CB | `HandleArtistsSelected` | Known | Event handler |
| 0x0072B80F | `HandleArtistsHilited` | Known | Event handler |
| 0x0072B849 | `HandleAlbumsSelected` | Known | Event handler |
| 0x0072B88C | `HandleAlbumsHilited` | Known | Event handler |
| 0x0072B8C5 | `HandleCompilationsSelected` | Known | Event handler |
| 0x0072B90E | `HandleCompilationsHilited` | Known | Event handler |
| 0x0072B94D | `HandleSongsSelected` | Known | Event handler |
| 0x0072B98F | `HandleSongsHilited` | Known | Event handler |
| 0x0072BA3A | `HandleGenericHilited` | Known | Event handler |
| 0x0072BAB2 | `HandleGenresSelected` | Known | Event handler |
| 0x0072BAF5 | `HandleGenresHilited` | Known | Event handler |
| 0x0072BB2E | `HandleComposersSelected` | Known | Event handler |
| 0x0072BB74 | `HandleComposersHilited` | Known | Event handler |
| 0x0072BBB0 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x0072BBF7 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x0072BCB6 | `HandleMusicHilited` | Known | Event handler |
| 0x0072BD2D | `HandlePlayPause` | Known | Event handler |
| 0x0072BD62 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x0072BE4C | `HandleSelect` | Known | Event handler |
| 0x0072BE92 | `HandleMoviesSelected` | Known | Event handler |
| 0x0072BED5 | `HandleMoviesHilited` | Known | Event handler |
| 0x0072BF0E | `HandleRentalsSelected` | Known | Event handler |
| 0x0072BF52 | `HandleRentalsHilited` | Known | Event handler |
| 0x0072BF8C | `HandleTVShowsSelected` | Known | Event handler |
| 0x0072BFD0 | `HandleTVShowsHilited` | Known | Event handler |
| 0x0072C00A | `HandleMusicVideosSelected` | Known | Event handler |
| 0x0072C052 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x0072C090 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x0072C0DB | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x0072C1A1 | `HandleVideosHilited` | Known | Event handler |
| 0x0072C7F3 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x0072D37A | `HandleMainMenu` | Known | Event handler |
| 0x0072D3B2 | `HandleMusicMenu` | Known | Event handler |
| 0x0072D8DA | `HandleRadioRegion` | Known | Event handler |
| 0x0072D97E | `HandleLanguage` | Known | Event handler |
| 0x0072DA84 | `HandleNew` | Known | Event handler |
| 0x0072DAFF | `HandleClear` | Known | Event handler |
| 0x0072DB30 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x0072DBEC | `HandleSelectIndexedSession` | Known | Event handler |
| 0x0072DD55 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x0072DDA8 | `HandleSelect` | Known | Event handler |
| 0x0072DED2 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x0072DF0C | `HandleEQSettingSelected` | Known | Event handler |
| 0x0072DF44 | `HandleEQSettingSelected` | Known | Event handler |
| 0x00740F0E | `HandleMenuSelection` | Known | Event handler |
| 0x00741253 | `HandleLoadingCancelled` | Known | Event handler |
| 0x007412EF | `HandleLoadingCancelled` | Known | Event handler |
| 0x007413BC | `HandleItemSelected` | Known | Event handler |
| 0x00741507 | `HandleNextContact` | Known | Event handler |
| 0x00741533 | `HandlePreviousContact` | Known | Event handler |
| 0x00741569 | `HandleSelectKey` | Known | Event handler |
| 0x00741B7A | `HandleSelect` | Known | Event handler |
| 0x00741EA1 | `HandleDateChosen` | Known | Event handler |
| 0x00741ED7 | `HandleTimeChosen` | Known | Event handler |
| 0x00741F0D | `HandleFrequencyChosen` | Known | Event handler |
| 0x00741F48 | `HandleSoundChosen` | Known | Event handler |
| 0x00741F7F | `HandleLabelChosen` | Known | Event handler |
| 0x00741FB6 | `HandleDeleteChosen` | Known | Event handler |
| 0x00741FF2 | `HandleSelect` | Known | Event handler |
| 0x0074202A | `HandleSelect` | Known | Event handler |
| 0x0074236B | `HandleLeaveAlarm` | Known | Event handler |
| 0x00742398 | `HandleLeaveAlarm` | Known | Event handler |
| 0x007423C7 | `HandleLeaveAlarm` | Known | Event handler |
| 0x007423F4 | `HandleLeaveAlarm` | Known | Event handler |
| 0x0074252E | `HandleSelect` | Known | Event handler |
| 0x0074255C | `HandleSelect` | Known | Event handler |
| 0x007426BB | `HandleNextDay` | Known | Event handler |
| 0x007426E3 | `HandlePreviousDay` | Known | Event handler |
| 0x00742892 | `HandleSelect` | Known | Event handler |
| 0x007428BF | `HandleNextDay` | Known | Event handler |
| 0x007428E7 | `HandlePreviousDay` | Known | Event handler |
| 0x00742A8F | `HandleNextDay` | Known | Event handler |
| 0x00742AB7 | `HandlePreviousDay` | Known | Event handler |
| 0x00742B78 | `HandleSelect` | Known | Event handler |
| 0x00742BA3 | `HandleNextDay` | Known | Event handler |
| 0x00742BCB | `HandlePreviousDay` | Known | Event handler |
| 0x00742D42 | `HandleSelectLozinch` | Known | Event handler |
| 0x00742EBA | `HandleSelectLozinch` | Known | Event handler |
| 0x00742FD9 | `HandleFlowNext` | Known | Event handler |
| 0x00743007 | `HandlePlayPause` | Known | Event handler |
| 0x00743056 | `HandleFlowPrev` | Known | Event handler |
| 0x00743081 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x00743175 | `HandleAlbumSelected` | Known | Event handler |
| 0x00743310 | `HandleFlowNext` | Known | Event handler |
| 0x0074335E | `HandleFlowNext` | Known | Event handler |
| 0x0074338C | `HandlePlayPause` | Known | Event handler |
| 0x007433DB | `HandleFlowPrev` | Known | Event handler |
| 0x00743407 | `HandleFlowPrev` | Known | Event handler |
| 0x00743427 | `HandleFlowWheel` | Known | Event handler |
| 0x007437B7 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x00743BE2 | `HandleArrowDown` | Known | Event handler |
| 0x00743C4C | `HandleArrowUp` | Known | Event handler |
| 0x00743C6B | `HandleWheel` | Known | Event handler |
| 0x00743CF4 | `HandleSelect` | Known | Event handler |
| 0x00743D71 | `HandleGameHilited` | Known | Event handler |
| 0x007471D7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00749113 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074B04F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074CF8B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074EEC7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00750E03 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00752D3F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00754C7B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00756BB7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00758AF3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075AA2F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075C96B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075E8A7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007607E3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076271F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076465B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00766597 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007684D3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076A40F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076C34B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076E287 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007701C3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007720FF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077403B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00775F77 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00777EB3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00779DEF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077BD2B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077DC67 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077FBA3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00781ADF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00783A1B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00785957 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00787893 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007897CF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078B70B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078D647 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078F568 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00790384 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007911A0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00791FBC | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00792DD8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00793BF4 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00794A10 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079582C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00796648 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00797464 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00798280 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079909C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00799EB8 | `HandlePlayPause` | Known | Event handler |
| 0x00799EEE | `HandleShowContextualMenu` | Known | Event handler |
| 0x00799F30 | `HandleAddToOTG` | Known | Event handler |
| 0x0079A0CD | `HandlePlayPause` | Known | Event handler |
| 0x0079A0F4 | `HandleSelect` | Known | Event handler |
| 0x0079A121 | `HandleHilite` | Known | Event handler |
| 0x0079A154 | `HandlePlayPause` | Known | Event handler |
| 0x0079A1E7 | `HandlePlayPause` | Known | Event handler |
| 0x0079A20E | `HandleSelect` | Known | Event handler |
| 0x0079A274 | `HandleHilite` | Known | Event handler |
| 0x0079A2A6 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x0079A2EB | `HandleNext` | Known | Event handler |
| 0x0079A315 | `HandlePlayPause` | Known | Event handler |
| 0x0079A33F | `HandlePrev` | Known | Event handler |
| 0x0079A366 | `HandleSelectMix` | Known | Event handler |
| 0x0079A3A7 | `HandleGeniusMixPlaylistReady` | Known | Event handler |
| 0x0079A4BC | `HandleWheel` | Known | Event handler |
| 0x0079A4EC | `HandlePlayPause` | Known | Event handler |
| 0x0079A522 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079A569 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079A5AC | `HandleAddToOTG` | Known | Event handler |
| 0x0079A60F | `HandleStartGenius` | Known | Event handler |
| 0x0079A64B | `HandleViewAlbum` | Known | Event handler |
| 0x0079A686 | `HandleViewArtist` | Known | Event handler |
| 0x0079A6C7 | `HandleViewCompilation` | Known | Event handler |
| 0x0079A867 | `HandlePlayPause` | Known | Event handler |
| 0x0079A88E | `HandleSelect` | Known | Event handler |
| 0x0079A8F8 | `HandlePlayPause` | Known | Event handler |
| 0x0079A92E | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079A975 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079A9B8 | `HandleAddToOTG` | Known | Event handler |
| 0x0079AA1B | `HandleStartGenius` | Known | Event handler |
| 0x0079AA57 | `HandleViewAlbum` | Known | Event handler |
| 0x0079AA92 | `HandleViewArtist` | Known | Event handler |
| 0x0079AAD3 | `HandleViewCompilation` | Known | Event handler |
| 0x0079AC73 | `HandlePlayPause` | Known | Event handler |
| 0x0079AC9A | `HandleSelect` | Known | Event handler |
| 0x0079AD04 | `HandlePlayPause` | Known | Event handler |
| 0x0079AD42 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079AD85 | `HandleAddToOTG` | Known | Event handler |
| 0x0079ADE8 | `HandleStartGenius` | Known | Event handler |
| 0x0079AE24 | `HandleViewAlbum` | Known | Event handler |
| 0x0079AE5F | `HandleViewArtist` | Known | Event handler |
| 0x0079AEA0 | `HandleViewCompilation` | Known | Event handler |
| 0x0079B033 | `HandleSelect` | Known | Event handler |
| 0x0079B098 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079B0DC | `HandlePlayPause` | Known | Event handler |
| 0x0079B112 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079B154 | `HandleAddToOTG` | Known | Event handler |
| 0x0079B3AE | `HandlePlayPause` | Known | Event handler |
| 0x0079B3D5 | `HandleSelect` | Known | Event handler |
| 0x0079B402 | `HandleHilite` | Known | Event handler |
| 0x0079B434 | `HandlePlayPause` | Known | Event handler |
| 0x0079B46A | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079B4AC | `HandleAddToOTG` | Known | Event handler |
| 0x0079B706 | `HandlePlayPause` | Known | Event handler |
| 0x0079B72D | `HandleSelect` | Known | Event handler |
| 0x0079B75A | `HandleHilite` | Known | Event handler |
| 0x0079B78C | `HandlePlayPause` | Known | Event handler |
| 0x0079B7C2 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079B804 | `HandleAddToOTG` | Known | Event handler |
| 0x0079BB17 | `HandlePlayPause` | Known | Event handler |
| 0x0079BB3E | `HandleSelect` | Known | Event handler |
| 0x0079BB70 | `HandlePlayPause` | Known | Event handler |
| 0x0079BBA6 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079BBE8 | `HandleAddToOTG` | Known | Event handler |
| 0x0079BCA2 | `HandlePlayPause` | Known | Event handler |
| 0x0079BCC9 | `HandleSelect` | Known | Event handler |
| 0x0079BD58 | `HandlePlayPause` | Known | Event handler |
| 0x0079BD8E | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079BDD0 | `HandleAddToOTG` | Known | Event handler |
| 0x0079BFB1 | `HandlePlayPause` | Known | Event handler |
| 0x0079BFD8 | `HandleSelect` | Known | Event handler |
| 0x0079C008 | `HandlePlayPause` | Known | Event handler |
| 0x0079C03E | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079C080 | `HandleAddToOTG` | Known | Event handler |
| 0x0079C12D | `HandleSelect` | Known | Event handler |
| 0x0079C1C6 | `HandleHilite` | Known | Event handler |
| 0x0079C1F2 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079C234 | `HandlePlayPause` | Known | Event handler |
| 0x0079C26A | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079C2AC | `HandleAddToOTG` | Known | Event handler |
| 0x0079C359 | `HandleSelect` | Known | Event handler |
| 0x0079C3BE | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079C400 | `HandlePlayPause` | Known | Event handler |
| 0x0079C5A4 | `HandleSelect` | Known | Event handler |
| 0x0079C5D1 | `HandleHilite` | Known | Event handler |
| 0x0079C5FD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079C640 | `HandlePlayPause` | Known | Event handler |
| 0x0079C6C6 | `HandleSelect` | Known | Event handler |
| 0x0079C754 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079C798 | `HandlePlayPause` | Known | Event handler |
| 0x0079C81E | `HandleSelect` | Known | Event handler |
| 0x0079C883 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079C8C4 | `HandlePlayPause` | Known | Event handler |
| 0x0079C94A | `HandleSelect` | Known | Event handler |
| 0x0079C9B0 | `HandleHilite` | Known | Event handler |
| 0x0079C9DC | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079CA20 | `HandlePlayPause` | Known | Event handler |
| 0x0079CA56 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079CA98 | `HandleAddToOTG` | Known | Event handler |
| 0x0079CE14 | `HandlePlayPause` | Known | Event handler |
| 0x0079CE3B | `HandleSelect` | Known | Event handler |
| 0x0079CE6C | `HandlePlayPause` | Known | Event handler |
| 0x0079CEA2 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079CEE9 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079CF2C | `HandleAddToOTG` | Known | Event handler |
| 0x0079CF8F | `HandleStartGenius` | Known | Event handler |
| 0x0079CFCB | `HandleViewAlbum` | Known | Event handler |
| 0x0079D006 | `HandleViewArtist` | Known | Event handler |
| 0x0079D047 | `HandleViewCompilation` | Known | Event handler |
| 0x0079D52F | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x0079D574 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079D5B7 | `HandleAddToOTG` | Known | Event handler |
| 0x0079D61A | `HandleStartGenius` | Known | Event handler |
| 0x0079D656 | `HandleViewAlbum` | Known | Event handler |
| 0x0079D691 | `HandleViewArtist` | Known | Event handler |
| 0x0079D6D2 | `HandleViewCompilation` | Known | Event handler |
| 0x0079DAA8 | `HandlePlayPause` | Known | Event handler |
| 0x0079DBD5 | `HandleSelect` | Known | Event handler |
| 0x0079DC01 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079DC44 | `HandlePlayPause` | Known | Event handler |
| 0x0079DCCA | `HandleSelect` | Known | Event handler |
| 0x0079DCF7 | `HandleHilite` | Known | Event handler |
| 0x0079DD23 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079DD64 | `HandlePlayPause` | Known | Event handler |
| 0x0079DE97 | `HandleSelect` | Known | Event handler |
| 0x0079DEC3 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079E7D5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079F08D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079F945 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007A01FD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007A0AB5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007A136D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007A1C25 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007A24DD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007A2526 | `HandleTVOutChanged` | Known | Event handler |
| 0x007A255E | `HandleTVSignalChanged` | Known | Event handler |
| 0x007A2599 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x007A25EA | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x007A262F | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x007A2678 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x007A26BA | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x007A2708 | `HandlePlayPause` | Known | Event handler |
| 0x007A273E | `HandleShowContextualMenu` | Known | Event handler |
| 0x007A2785 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A27C8 | `HandleAddToOTG` | Known | Event handler |
| 0x007A282B | `HandleStartGenius` | Known | Event handler |
| 0x007A2867 | `HandleViewAlbum` | Known | Event handler |
| 0x007A28A2 | `HandleViewArtist` | Known | Event handler |
| 0x007A28E3 | `HandleViewCompilation` | Known | Event handler |
| 0x007A2B1F | `HandlePlayPause` | Known | Event handler |
| 0x007A2B46 | `HandleSelect` | Known | Event handler |
| 0x007A2B78 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x007A2BB3 | `HandleSaveGeniusPlaylist` | Known | Event handler |
| 0x007A2C2C | `HandlePlayPause` | Known | Event handler |
| 0x007A2C62 | `HandleShowContextualMenu` | Known | Event handler |
| 0x007A2CA9 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A2CEC | `HandleAddToOTG` | Known | Event handler |
| 0x007A2D4F | `HandleStartGenius` | Known | Event handler |
| 0x007A2D8B | `HandleViewAlbum` | Known | Event handler |
| 0x007A2DC6 | `HandleViewArtist` | Known | Event handler |
| 0x007A2E07 | `HandleViewCompilation` | Known | Event handler |
| 0x007A2E75 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x007A329D | `HandlePlayPause` | Known | Event handler |
| 0x007A32C4 | `HandleSelect` | Known | Event handler |
| 0x007A32F6 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x007A332D | `HandleSelect` | Known | Event handler |
| 0x007A335D | `HandleSelect` | Known | Event handler |
| 0x007A3395 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A33C3 | `HandleMenuKey` | Known | Event handler |
| 0x007A3449 | `HandlePlayPause` | Known | Event handler |
| 0x007A34D3 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A3508 | `HandleSelect` | Known | Event handler |
| 0x007A3543 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A3586 | `HandleAddToOTG` | Known | Event handler |
| 0x007A35C5 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A360B | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A3651 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A36BB | `HandleStartGenius` | Known | Event handler |
| 0x007A36F7 | `HandleViewAlbum` | Known | Event handler |
| 0x007A3732 | `HandleViewArtist` | Known | Event handler |
| 0x007A3773 | `HandleViewCompilation` | Known | Event handler |
| 0x007A41B5 | `HandleStartGenius` | Known | Event handler |
| 0x007A42C8 | `HandlePlayPause` | Known | Event handler |
| 0x007A433D | `HandleWheelProgress` | Known | Event handler |
| 0x007A4379 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A43A7 | `HandleMenuKey` | Known | Event handler |
| 0x007A442D | `HandlePlayPause` | Known | Event handler |
| 0x007A44B7 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A44EC | `HandleSelectProgress` | Known | Event handler |
| 0x007A452F | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A4572 | `HandleAddToOTG` | Known | Event handler |
| 0x007A45B1 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A45F7 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A463D | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A46A7 | `HandleStartGenius` | Known | Event handler |
| 0x007A46E3 | `HandleViewAlbum` | Known | Event handler |
| 0x007A471E | `HandleViewArtist` | Known | Event handler |
| 0x007A475F | `HandleViewCompilation` | Known | Event handler |
| 0x007A51A1 | `HandleStartGenius` | Known | Event handler |
| 0x007A52B4 | `HandlePlayPause` | Known | Event handler |
| 0x007A5329 | `HandleWheelProgress` | Known | Event handler |
| 0x007A5365 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A5393 | `HandleMenuKey` | Known | Event handler |
| 0x007A5419 | `HandlePlayPause` | Known | Event handler |
| 0x007A54A3 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A54D8 | `HandleSelectVolume` | Known | Event handler |
| 0x007A5519 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A555C | `HandleAddToOTG` | Known | Event handler |
| 0x007A559B | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A55E1 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A5627 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A5691 | `HandleStartGenius` | Known | Event handler |
| 0x007A56CD | `HandleViewAlbum` | Known | Event handler |
| 0x007A5708 | `HandleViewArtist` | Known | Event handler |
| 0x007A5749 | `HandleViewCompilation` | Known | Event handler |
| 0x007A618B | `HandleStartGenius` | Known | Event handler |
| 0x007A629E | `HandlePlayPause` | Known | Event handler |
| 0x007A6313 | `HandleWheelVolume` | Known | Event handler |
| 0x007A634D | `HandleMenuLongpress` | Known | Event handler |
| 0x007A637B | `HandleMenuKey` | Known | Event handler |
| 0x007A6401 | `HandlePlayPause` | Known | Event handler |
| 0x007A648B | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A64C0 | `HandleSelectRating` | Known | Event handler |
| 0x007A6501 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A6544 | `HandleAddToOTG` | Known | Event handler |
| 0x007A6583 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A65C9 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A660F | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A6679 | `HandleStartGenius` | Known | Event handler |
| 0x007A66B5 | `HandleViewAlbum` | Known | Event handler |
| 0x007A66F0 | `HandleViewArtist` | Known | Event handler |
| 0x007A6731 | `HandleViewCompilation` | Known | Event handler |
| 0x007A7173 | `HandleStartGenius` | Known | Event handler |
| 0x007A7286 | `HandlePlayPause` | Known | Event handler |
| 0x007A72FB | `HandleWheelRating` | Known | Event handler |
| 0x007A7335 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A7363 | `HandleMenuKey` | Known | Event handler |
| 0x007A73DB | `HandlePlayPause` | Known | Event handler |
| 0x007A745C | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A7491 | `HandleSelectScrub` | Known | Event handler |
| 0x007A74D1 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A7514 | `HandleAddToOTG` | Known | Event handler |
| 0x007A7553 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A7599 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A75DF | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A7649 | `HandleStartGenius` | Known | Event handler |
| 0x007A7685 | `HandleViewAlbum` | Known | Event handler |
| 0x007A76C0 | `HandleViewArtist` | Known | Event handler |
| 0x007A7701 | `HandleViewCompilation` | Known | Event handler |
| 0x007A8143 | `HandleStartGenius` | Known | Event handler |
| 0x007A8248 | `HandlePlayPause` | Known | Event handler |
| 0x007A82B4 | `HandleWheelScrub` | Known | Event handler |
| 0x007A82ED | `HandleMenuLongpress` | Known | Event handler |
| 0x007A831B | `HandleMenuKey` | Known | Event handler |
| 0x007A83A1 | `HandlePlayPause` | Known | Event handler |
| 0x007A842B | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A8460 | `HandleSelectGenius` | Known | Event handler |
| 0x007A84A1 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A84E4 | `HandleAddToOTG` | Known | Event handler |
| 0x007A8523 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A8569 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A85AF | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A8619 | `HandleStartGenius` | Known | Event handler |
| 0x007A8655 | `HandleViewAlbum` | Known | Event handler |
| 0x007A8690 | `HandleViewArtist` | Known | Event handler |
| 0x007A86D1 | `HandleViewCompilation` | Known | Event handler |
| 0x007A9113 | `HandleStartGenius` | Known | Event handler |
| 0x007A9226 | `HandlePlayPause` | Known | Event handler |
| 0x007A929B | `HandleWheelGenius` | Known | Event handler |
| 0x007A92D5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A9303 | `HandleMenuKey` | Known | Event handler |
| 0x007A9360 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007A9398 | `HandlePlayPause` | Known | Event handler |
| 0x007A93F2 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007A9431 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A9466 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x007A94AE | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A94F1 | `HandleAddToOTG` | Known | Event handler |
| 0x007A9530 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A9576 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A95BC | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A9626 | `HandleStartGenius` | Known | Event handler |
| 0x007A9662 | `HandleViewAlbum` | Known | Event handler |
| 0x007A969D | `HandleViewArtist` | Known | Event handler |
| 0x007A96DE | `HandleViewCompilation` | Known | Event handler |
| 0x007AA120 | `HandleStartGenius` | Known | Event handler |
| 0x007AA233 | `HandlePlayPause` | Known | Event handler |
| 0x007AA2A8 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007AA2E9 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AA317 | `HandleMenuKey` | Known | Event handler |
| 0x007AA39D | `HandlePlayPause` | Known | Event handler |
| 0x007AA427 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007AA45C | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007AA4A0 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007AA4E3 | `HandleAddToOTG` | Known | Event handler |
| 0x007AA522 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007AA568 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007AA5AE | `HandleAudiobookSlower` | Known | Event handler |
| 0x007AA618 | `HandleStartGenius` | Known | Event handler |
| 0x007AA654 | `HandleViewAlbum` | Known | Event handler |
| 0x007AA68F | `HandleViewArtist` | Known | Event handler |
| 0x007AA6D0 | `HandleViewCompilation` | Known | Event handler |
| 0x007AB112 | `HandleStartGenius` | Known | Event handler |
| 0x007AB225 | `HandlePlayPause` | Known | Event handler |
| 0x007AB2C5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AB2F3 | `HandleMenuKey` | Known | Event handler |
| 0x007AB379 | `HandlePlayPause` | Known | Event handler |
| 0x007AB403 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007AB438 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007AB47C | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007AB4BF | `HandleAddToOTG` | Known | Event handler |
| 0x007AB4FE | `HandleAudiobookFaster` | Known | Event handler |
| 0x007AB544 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007AB58A | `HandleAudiobookSlower` | Known | Event handler |
| 0x007AB5F4 | `HandleStartGenius` | Known | Event handler |
| 0x007AB630 | `HandleViewAlbum` | Known | Event handler |
| 0x007AB66B | `HandleViewArtist` | Known | Event handler |
| 0x007AB6AC | `HandleViewCompilation` | Known | Event handler |
| 0x007AC0EE | `HandleStartGenius` | Known | Event handler |
| 0x007AC201 | `HandlePlayPause` | Known | Event handler |
| 0x007AC2A1 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AC2CF | `HandleMenuKey` | Known | Event handler |
| 0x007AC355 | `HandlePlayPause` | Known | Event handler |
| 0x007AC3DF | `HandlePushContextualMenu` | Known | Event handler |
| 0x007AC414 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007AC458 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007AC49B | `HandleAddToOTG` | Known | Event handler |
| 0x007AC4DA | `HandleAudiobookFaster` | Known | Event handler |
| 0x007AC520 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007AC566 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007AC5D0 | `HandleStartGenius` | Known | Event handler |
| 0x007AC60C | `HandleViewAlbum` | Known | Event handler |
| 0x007AC647 | `HandleViewArtist` | Known | Event handler |
| 0x007AC688 | `HandleViewCompilation` | Known | Event handler |
| 0x007AD0CA | `HandleStartGenius` | Known | Event handler |
| 0x007AD1DD | `HandlePlayPause` | Known | Event handler |
| 0x007AD27D | `HandleMenuLongpress` | Known | Event handler |
| 0x007AD2AB | `HandleMenuKey` | Known | Event handler |
| 0x007AD331 | `HandlePlayPause` | Known | Event handler |
| 0x007AD3BB | `HandlePushContextualMenu` | Known | Event handler |
| 0x007AD3F0 | `HandleSelectChapterArt` | Known | Event handler |
| 0x007AD435 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007AD478 | `HandleAddToOTG` | Known | Event handler |
| 0x007AD4B7 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007AD4FD | `HandleAudiobookNormal` | Known | Event handler |
| 0x007AD543 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007AD5AD | `HandleStartGenius` | Known | Event handler |
| 0x007AD5E9 | `HandleViewAlbum` | Known | Event handler |
| 0x007AD624 | `HandleViewArtist` | Known | Event handler |
| 0x007AD665 | `HandleViewCompilation` | Known | Event handler |
| 0x007AE0A7 | `HandleStartGenius` | Known | Event handler |
| 0x007AE1BA | `HandlePlayPause` | Known | Event handler |
| 0x007AE22F | `HandleWheelVolume` | Known | Event handler |
| 0x007AE269 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AE297 | `HandleMenuKey` | Known | Event handler |
| 0x007AE326 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007AE3C7 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007AE3FC | `HandleSelect` | Known | Event handler |
| 0x007AE437 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007AE47A | `HandleAddToOTG` | Known | Event handler |
| 0x007AE4B9 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007AE4FF | `HandleAudiobookNormal` | Known | Event handler |
| 0x007AE545 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007AE5AF | `HandleStartGenius` | Known | Event handler |
| 0x007AE5EB | `HandleViewAlbum` | Known | Event handler |
| 0x007AE626 | `HandleViewArtist` | Known | Event handler |
| 0x007AE667 | `HandleViewCompilation` | Known | Event handler |
| 0x007AF0A9 | `HandleStartGenius` | Known | Event handler |
| 0x007AF1C5 | `HandlePlayPause` | Known | Event handler |
| 0x007AF243 | `HandleWheel` | Known | Event handler |
| 0x007AF279 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AF2A7 | `HandleMenuKey` | Known | Event handler |
| 0x007AF336 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007AF3D7 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007AF40C | `HandleSelect` | Known | Event handler |
| 0x007AF447 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007AF48A | `HandleAddToOTG` | Known | Event handler |
| 0x007AF4C9 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007AF50F | `HandleAudiobookNormal` | Known | Event handler |
| 0x007AF555 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007AF5BF | `HandleStartGenius` | Known | Event handler |
| 0x007AF5FB | `HandleViewAlbum` | Known | Event handler |
| 0x007AF636 | `HandleViewArtist` | Known | Event handler |
| 0x007AF677 | `HandleViewCompilation` | Known | Event handler |
| 0x007B00B9 | `HandleStartGenius` | Known | Event handler |
| 0x007B01D5 | `HandlePlayPause` | Known | Event handler |
| 0x007B0253 | `HandleWheel` | Known | Event handler |
| 0x007B0289 | `HandleMenuLongpress` | Known | Event handler |
| 0x007B02B7 | `HandleMenuKey` | Known | Event handler |
| 0x007B033D | `HandlePlayPause` | Known | Event handler |
| 0x007B03C7 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007B03FC | `HandleSelect` | Known | Event handler |
| 0x007B0437 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007B047A | `HandleAddToOTG` | Known | Event handler |
| 0x007B04B9 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007B04FF | `HandleAudiobookNormal` | Known | Event handler |
| 0x007B0545 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007B05AF | `HandleStartGenius` | Known | Event handler |
| 0x007B05EB | `HandleViewAlbum` | Known | Event handler |
| 0x007B0626 | `HandleViewArtist` | Known | Event handler |
| 0x007B0667 | `HandleViewCompilation` | Known | Event handler |
| 0x007B10A9 | `HandleStartGenius` | Known | Event handler |
| 0x007B11BC | `HandlePlayPause` | Known | Event handler |
| 0x007B1231 | `HandleWheel` | Known | Event handler |
| 0x007B1265 | `HandleMenuLongpress` | Known | Event handler |
| 0x007B1293 | `HandleMenuKey` | Known | Event handler |
| 0x007B1319 | `HandlePlayPause` | Known | Event handler |
| 0x007B13A3 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007B13D8 | `HandleSelectProgress` | Known | Event handler |
| 0x007B141B | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007B145E | `HandleAddToOTG` | Known | Event handler |
| 0x007B149D | `HandleAudiobookFaster` | Known | Event handler |
| 0x007B14E3 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007B1529 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007B1593 | `HandleStartGenius` | Known | Event handler |
| 0x007B15CF | `HandleViewAlbum` | Known | Event handler |
| 0x007B160A | `HandleViewArtist` | Known | Event handler |
| 0x007B164B | `HandleViewCompilation` | Known | Event handler |
| 0x007B208D | `HandleStartGenius` | Known | Event handler |
| 0x007B21A0 | `HandlePlayPause` | Known | Event handler |
| 0x007B2215 | `HandleWheelProgress` | Known | Event handler |
| 0x007B2251 | `HandleMenuLongpress` | Known | Event handler |
| 0x007B227F | `HandleMenuKey` | Known | Event handler |
| 0x007B22F7 | `HandlePlayPause` | Known | Event handler |
| 0x007B2378 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007B23AD | `HandleSelectScrub` | Known | Event handler |
| 0x007B23ED | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007B2430 | `HandleAddToOTG` | Known | Event handler |
| 0x007B246F | `HandleAudiobookFaster` | Known | Event handler |
| 0x007B24B5 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007B24FB | `HandleAudiobookSlower` | Known | Event handler |
| 0x007B2565 | `HandleStartGenius` | Known | Event handler |
| 0x007B25A1 | `HandleViewAlbum` | Known | Event handler |
| 0x007B25DC | `HandleViewArtist` | Known | Event handler |
| 0x007B261D | `HandleViewCompilation` | Known | Event handler |
| 0x007B305F | `HandleStartGenius` | Known | Event handler |
| 0x007B3164 | `HandlePlayPause` | Known | Event handler |
| 0x007B31D0 | `HandleWheelScrub` | Known | Event handler |
| 0x007B3209 | `HandleMenuLongpress` | Known | Event handler |
| 0x007B3237 | `HandleMenuKey` | Known | Event handler |
| 0x007B32BD | `HandlePlayPause` | Known | Event handler |
| 0x007B3347 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007B33B6 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007B33F9 | `HandleAddToOTG` | Known | Event handler |
| 0x007B3438 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007B347E | `HandleAudiobookNormal` | Known | Event handler |
| 0x007B34C4 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007B352E | `HandleStartGenius` | Known | Event handler |
| 0x007B356A | `HandleViewAlbum` | Known | Event handler |
| 0x007B35A5 | `HandleViewArtist` | Known | Event handler |
| 0x007B35E6 | `HandleViewCompilation` | Known | Event handler |
| 0x007B4028 | `HandleStartGenius` | Known | Event handler |
| 0x007B413B | `HandlePlayPause` | Known | Event handler |
| 0x007B41B0 | `HandleWheelVolume` | Known | Event handler |
| 0x007B41ED | `HandleMenuLongpress` | Known | Event handler |
| 0x007B421B | `HandleMenuKey` | Known | Event handler |
| 0x007B42A1 | `HandlePlayPause` | Known | Event handler |
| 0x007B432B | `HandlePushContextualMenu` | Known | Event handler |
| 0x007B439A | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007B43DD | `HandleAddToOTG` | Known | Event handler |
| 0x007B441C | `HandleAudiobookFaster` | Known | Event handler |
| 0x007B4462 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007B44A8 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007B4512 | `HandleStartGenius` | Known | Event handler |
| 0x007B454E | `HandleViewAlbum` | Known | Event handler |
| 0x007B4589 | `HandleViewArtist` | Known | Event handler |
| 0x007B45CA | `HandleViewCompilation` | Known | Event handler |
| 0x007B500C | `HandleStartGenius` | Known | Event handler |
| 0x007B511F | `HandlePlayPause` | Known | Event handler |
| 0x007B5194 | `HandleWheelBrightness` | Known | Event handler |
| 0x007B52B7 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007B52EC | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007B5334 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007B5377 | `HandleAddToOTG` | Known | Event handler |
| 0x007B53B6 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007B53FC | `HandleAudiobookNormal` | Known | Event handler |
| 0x007B5442 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007B54AC | `HandleStartGenius` | Known | Event handler |
| 0x007B54E8 | `HandleViewAlbum` | Known | Event handler |
| 0x007B5523 | `HandleViewArtist` | Known | Event handler |
| 0x007B5564 | `HandleViewCompilation` | Known | Event handler |
| 0x007B5FA6 | `HandleStartGenius` | Known | Event handler |
| 0x007B60F2 | `HandleWheel` | Known | Event handler |
| 0x007B6129 | `HandleMenuLongpress` | Known | Event handler |
| 0x007B6157 | `HandleMenuKey` | Known | Event handler |
| 0x007B61DD | `HandlePlayPause` | Known | Event handler |
| 0x007B625D | `HandleSelect` | Known | Event handler |
| 0x007B66FF | `HandlePlayPause` | Known | Event handler |
| 0x007B678D | `HandleMenuLongpress` | Known | Event handler |
| 0x007B67BB | `HandleMenuKey` | Known | Event handler |
| 0x007B6841 | `HandlePlayPause` | Known | Event handler |
| 0x007B68C1 | `HandleSelectProgress` | Known | Event handler |
| 0x007B6D6B | `HandlePlayPause` | Known | Event handler |
| 0x007B6DE0 | `HandleWheelProgress` | Known | Event handler |
| 0x007B6E1D | `HandleMenuLongpress` | Known | Event handler |
| 0x007B6E4B | `HandleMenuKey` | Known | Event handler |
| 0x007B6ED1 | `HandlePlayPause` | Known | Event handler |
| 0x007B6F51 | `HandleSelectProgress` | Known | Event handler |
| 0x007B73FB | `HandlePlayPause` | Known | Event handler |
| 0x007B7470 | `HandleWheelProgress` | Known | Event handler |
| 0x007B74AD | `HandleMenuLongpress` | Known | Event handler |
| 0x007B74DB | `HandleMenuKey` | Known | Event handler |
| 0x007B7561 | `HandlePlayPause` | Known | Event handler |
| 0x007B75E1 | `HandleSelectProgress` | Known | Event handler |
| 0x007B7A17 | `HandlePlayPause` | Known | Event handler |
| 0x007B7A8C | `HandleWheelProgress` | Known | Event handler |
| 0x007B7AC9 | `HandleMenuLongpress` | Known | Event handler |
| 0x007B7AF7 | `HandleMenuKey` | Known | Event handler |
| 0x007B7B64 | `HandlePlayPause` | Known | Event handler |
| 0x007B7BD0 | `HandleSelectScrub` | Known | Event handler |
| 0x007B7FEA | `HandlePlayPause` | Known | Event handler |
| 0x007B804B | `HandleWheelScrub` | Known | Event handler |
| 0x007B8085 | `HandleMenuLongpress` | Known | Event handler |
| 0x007B80B3 | `HandleMenuKey` | Known | Event handler |
| 0x007B8139 | `HandlePlayPause` | Known | Event handler |
| 0x007B81B9 | `HandleSelectVolume` | Known | Event handler |
| 0x007B85ED | `HandlePlayPause` | Known | Event handler |
| 0x007B8662 | `HandleWheelVolume` | Known | Event handler |
| 0x007B8775 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007B8C14 | `HandleSelect` | Known | Event handler |
| 0x007B8C41 | `HandleSelect` | Known | Event handler |
| 0x007B8C71 | `HandleSelect` | Known | Event handler |
| 0x007B8CA1 | `HandleSelect` | Known | Event handler |
| 0x007B8CD1 | `HandleSelect` | Known | Event handler |
| 0x007B8D01 | `HandleSelect` | Known | Event handler |
| 0x007B8D31 | `HandleSelect` | Known | Event handler |
| 0x007B8D61 | `HandleSelect` | Known | Event handler |
| 0x007B8D91 | `HandleSelect` | Known | Event handler |
| 0x007B8E01 | `HandleSelect` | Known | Event handler |
| 0x007B8E31 | `HandleSelect` | Known | Event handler |
| 0x007B8EA9 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007B8EDC | `HandleNotesPop` | Known | Event handler |
| 0x007B8F59 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007B8F8C | `HandleNotesPop` | Known | Event handler |
| 0x007B9448 | `HandleNotesSelected` | Known | Event handler |
| 0x007B9485 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007B94B8 | `HandleNotesPop` | Known | Event handler |
| 0x007B9974 | `HandleNotesSelected` | Known | Event handler |
| 0x007B99B1 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007B99E4 | `HandleNotesPop` | Known | Event handler |
| 0x007B9A0F | `HandleNotesSelected` | Known | Event handler |
| 0x007B9EE1 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007B9F14 | `HandleNotesPop` | Known | Event handler |
| 0x007B9F3F | `HandleNotesSelected` | Known | Event handler |
| 0x007BA411 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007BA444 | `HandleNotesPop` | Known | Event handler |
| 0x007BA4C1 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007BA4F4 | `HandleNotesPop` | Known | Event handler |
| 0x007BA571 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007BA5A4 | `HandleNotesPop` | Known | Event handler |
| 0x007BA61C | `HandlePlayPause` | Known | Event handler |
| 0x007BA645 | `HandlePlayPause` | Known | Event handler |
| 0x007BA673 | `HandlePlayPause` | Known | Event handler |
| 0x007BA6A8 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007BA728 | `HandleHiliteAlbum` | Known | Event handler |
| 0x007BA7D1 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007BA858 | `HandleHiliteAlbum` | Known | Event handler |
| 0x007BAB1C | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x007BAB78 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x007BAD2F | `HandleSelect` | Known | Event handler |
| 0x007BAEB3 | `HandleSelect` | Known | Event handler |
| 0x007BAEED | `HandleImageLast` | Known | Event handler |
| 0x007BAF17 | `HandleImageNext` | Known | Event handler |
| 0x007BAF46 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007BAF80 | `HandleImageFirst` | Known | Event handler |
| 0x007BAFAB | `HandleImagePrev` | Known | Event handler |
| 0x007BAFD7 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007BB006 | `HandleImageNext` | Known | Event handler |
| 0x007BB02F | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007BB063 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007BB092 | `HandleImagePrev` | Known | Event handler |
| 0x007BB0B3 | `HandleImageWheel` | Known | Event handler |
| 0x007BB151 | `HandleImageNext` | Known | Event handler |
| 0x007BB180 | `HandlePlayPause` | Known | Event handler |
| 0x007BB1CF | `HandleImagePrev` | Known | Event handler |
| 0x007BB1FB | `HandleSelect` | Known | Event handler |
| 0x007BB4CB | `HandleImageNext` | Known | Event handler |
| 0x007BB4F5 | `HandlePause` | Known | Event handler |
| 0x007BB51A | `HandlePlay` | Known | Event handler |
| 0x007BB543 | `HandlePlayPause` | Known | Event handler |
| 0x007BB56C | `HandleImagePrev` | Known | Event handler |
| 0x007BB5CF | `HandleMikeyCenter` | Known | Event handler |
| 0x007BB5F2 | `HandleWheel` | Known | Event handler |
| 0x007BB68D | `HandleImageNext` | Known | Event handler |
| 0x007BB6BC | `HandlePlayPause` | Known | Event handler |
| 0x007BB70B | `HandleImagePrev` | Known | Event handler |
| 0x007BB737 | `HandleSelect` | Known | Event handler |
| 0x007BBA07 | `HandleImageNext` | Known | Event handler |
| 0x007BBA31 | `HandlePause` | Known | Event handler |
| 0x007BBA56 | `HandlePlay` | Known | Event handler |
| 0x007BBA7F | `HandlePlayPause` | Known | Event handler |
| 0x007BBAA8 | `HandleImagePrev` | Known | Event handler |
| 0x007BBB0B | `HandleMikeyCenter` | Known | Event handler |
| 0x007BBB2E | `HandleWheel` | Known | Event handler |
| 0x007BBBC9 | `HandleImageNext` | Known | Event handler |
| 0x007BBBF8 | `HandlePlayPause` | Known | Event handler |
| 0x007BBC47 | `HandleImagePrev` | Known | Event handler |
| 0x007BBC73 | `HandleSelect` | Known | Event handler |
| 0x007BBF43 | `HandleImageNext` | Known | Event handler |
| 0x007BBF6D | `HandlePause` | Known | Event handler |
| 0x007BBF92 | `HandlePlay` | Known | Event handler |
| 0x007BBFBB | `HandlePlayPause` | Known | Event handler |
| 0x007BBFE4 | `HandleImagePrev` | Known | Event handler |
| 0x007BC047 | `HandleMikeyCenter` | Known | Event handler |
| 0x007BC06A | `HandleWheel` | Known | Event handler |
| 0x007BC105 | `HandleImageNext` | Known | Event handler |
| 0x007BC134 | `HandlePlayPause` | Known | Event handler |
| 0x007BC183 | `HandleImagePrev` | Known | Event handler |
| 0x007BC1AF | `HandleSelect` | Known | Event handler |
| 0x007BC47F | `HandleImageNext` | Known | Event handler |
| 0x007BC4A9 | `HandlePause` | Known | Event handler |
| 0x007BC4CE | `HandlePlay` | Known | Event handler |
| 0x007BC4F7 | `HandlePlayPause` | Known | Event handler |
| 0x007BC520 | `HandleImagePrev` | Known | Event handler |
| 0x007BC583 | `HandleMikeyCenter` | Known | Event handler |
| 0x007BC5A6 | `HandleWheel` | Known | Event handler |
| 0x007BC641 | `HandleImageNext` | Known | Event handler |
| 0x007BC670 | `HandlePlayPause` | Known | Event handler |
| 0x007BC6BF | `HandleImagePrev` | Known | Event handler |
| 0x007BC6EB | `HandleSelect` | Known | Event handler |
| 0x007BC9BB | `HandleImageNext` | Known | Event handler |
| 0x007BC9E5 | `HandlePause` | Known | Event handler |
| 0x007BCA0A | `HandlePlay` | Known | Event handler |
| 0x007BCA33 | `HandlePlayPause` | Known | Event handler |
| 0x007BCA5C | `HandleImagePrev` | Known | Event handler |
| 0x007BCABF | `HandleMikeyCenter` | Known | Event handler |
| 0x007BCAE2 | `HandleWheel` | Known | Event handler |
| 0x007BCB7D | `HandleImageNext` | Known | Event handler |
| 0x007BCBAC | `HandlePlayPause` | Known | Event handler |
| 0x007BCBFB | `HandleImagePrev` | Known | Event handler |
| 0x007BCC27 | `HandleSelect` | Known | Event handler |
| 0x007BCEF7 | `HandleImageNext` | Known | Event handler |
| 0x007BCF21 | `HandlePause` | Known | Event handler |
| 0x007BCF46 | `HandlePlay` | Known | Event handler |
| 0x007BCF6F | `HandlePlayPause` | Known | Event handler |
| 0x007BCF98 | `HandleImagePrev` | Known | Event handler |
| 0x007BCFFB | `HandleMikeyCenter` | Known | Event handler |
| 0x007BD01E | `HandleWheel` | Known | Event handler |
| 0x007BD0B9 | `HandleImageNext` | Known | Event handler |
| 0x007BD0E8 | `HandlePlayPause` | Known | Event handler |
| 0x007BD137 | `HandleImagePrev` | Known | Event handler |
| 0x007BD163 | `HandleSelect` | Known | Event handler |
| 0x007BD3AE | `HandleImageNext` | Known | Event handler |
| 0x007BD3D8 | `HandlePause` | Known | Event handler |
| 0x007BD3FD | `HandlePlay` | Known | Event handler |
| 0x007BD426 | `HandlePlayPause` | Known | Event handler |
| 0x007BD44F | `HandleImagePrev` | Known | Event handler |
| 0x007BD4C2 | `HandleMikeyCenter` | Known | Event handler |
| 0x007BD4E5 | `HandleWheel` | Known | Event handler |
| 0x007BD57D | `HandleImageNext` | Known | Event handler |
| 0x007BD5AC | `HandlePlayPause` | Known | Event handler |
| 0x007BD5FB | `HandleImagePrev` | Known | Event handler |
| 0x007BD627 | `HandleSelect` | Known | Event handler |
| 0x007BD872 | `HandleImageNext` | Known | Event handler |
| 0x007BD89C | `HandlePause` | Known | Event handler |
| 0x007BD8C1 | `HandlePlay` | Known | Event handler |
| 0x007BD8EA | `HandlePlayPause` | Known | Event handler |
| 0x007BD913 | `HandleImagePrev` | Known | Event handler |
| 0x007BD986 | `HandleMikeyCenter` | Known | Event handler |
| 0x007BD9A9 | `HandleWheel` | Known | Event handler |
| 0x007BDA41 | `HandleImageNext` | Known | Event handler |
| 0x007BDA70 | `HandlePlayPause` | Known | Event handler |
| 0x007BDABF | `HandleImagePrev` | Known | Event handler |
| 0x007BDAEB | `HandleSelect` | Known | Event handler |
| 0x007BDD36 | `HandleImageNext` | Known | Event handler |
| 0x007BDD60 | `HandlePause` | Known | Event handler |
| 0x007BDD85 | `HandlePlay` | Known | Event handler |
| 0x007BDDAE | `HandlePlayPause` | Known | Event handler |
| 0x007BDDD7 | `HandleImagePrev` | Known | Event handler |
| 0x007BDE4A | `HandleMikeyCenter` | Known | Event handler |
| 0x007BDE6D | `HandleWheel` | Known | Event handler |
| 0x007BDF05 | `HandleImageNext` | Known | Event handler |
| 0x007BDF34 | `HandlePlayPause` | Known | Event handler |
| 0x007BDF83 | `HandleImagePrev` | Known | Event handler |
| 0x007BDFAF | `HandleSelect` | Known | Event handler |
| 0x007BE1FA | `HandleImageNext` | Known | Event handler |
| 0x007BE224 | `HandlePause` | Known | Event handler |
| 0x007BE249 | `HandlePlay` | Known | Event handler |
| 0x007BE272 | `HandlePlayPause` | Known | Event handler |
| 0x007BE29B | `HandleImagePrev` | Known | Event handler |
| 0x007BE30E | `HandleMikeyCenter` | Known | Event handler |
| 0x007BE331 | `HandleWheel` | Known | Event handler |
| 0x007BE3C9 | `HandleImageNext` | Known | Event handler |
| 0x007BE3F8 | `HandlePlayPause` | Known | Event handler |
| 0x007BE447 | `HandleImagePrev` | Known | Event handler |
| 0x007BE473 | `HandleSelect` | Known | Event handler |
| 0x007BE6BE | `HandleImageNext` | Known | Event handler |
| 0x007BE6E8 | `HandlePause` | Known | Event handler |
| 0x007BE70D | `HandlePlay` | Known | Event handler |
| 0x007BE736 | `HandlePlayPause` | Known | Event handler |
| 0x007BE75F | `HandleImagePrev` | Known | Event handler |
| 0x007BE7D2 | `HandleMikeyCenter` | Known | Event handler |
| 0x007BE7F5 | `HandleWheel` | Known | Event handler |
| 0x007BE821 | `HandleSelect` | Known | Event handler |
| 0x007BE851 | `HandleSelect` | Known | Event handler |
| 0x007BE974 | `HandleTuning` | Known | Event handler |
| 0x007BEB34 | `HandleVolumeChange` | Known | Event handler |
| 0x007BEB9B | `HandleVolumeChange` | Known | Event handler |
| 0x007BEC00 | `HandleVolumeChange` | Known | Event handler |
| 0x007BED4C | `HandleVolumeWheel` | Known | Event handler |
| 0x007BEEA7 | `HandleTuningSelect` | Known | Event handler |
| 0x007BF06D | `HandleVolumeChange` | Known | Event handler |
| 0x007BF0D4 | `HandleVolumeChange` | Known | Event handler |
| 0x007BF139 | `HandleVolumeChange` | Known | Event handler |
| 0x007BF285 | `HandleFrequencyChange` | Known | Event handler |
| 0x007BF3E3 | `HandleTuningSelect` | Known | Event handler |
| 0x007BF5A9 | `HandleVolumeChange` | Known | Event handler |
| 0x007BF610 | `HandleVolumeChange` | Known | Event handler |
| 0x007BF675 | `HandleVolumeChange` | Known | Event handler |
| 0x007BF7C1 | `HandleFrequencyChange` | Known | Event handler |
| 0x007BF8EC | `HandleTimerDone` | Known | Event handler |
| 0x007BFAE5 | `HandleVolumeChange` | Known | Event handler |
| 0x007BFB17 | `HandleVolumeChange` | Known | Event handler |
| 0x007BFB47 | `HandleVolumeChange` | Known | Event handler |
| 0x007BFC5E | `HandleVolumeWheel` | Known | Event handler |
| 0x007C04AF | `HandleExitUnsupported` | Known | Event handler |
| 0x007C04E1 | `HandleExitUnsupported` | Known | Event handler |
| 0x007C5515 | `HandleSelectKey` | Known | Event handler |
| 0x007C554A | `HandleWheel` | Known | Event handler |
| 0x007C5698 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x007C56EB | `HandleSelectKey` | Known | Event handler |
| 0x007C5713 | `HandleSelectKey` | Known | Event handler |
| 0x007C5743 | `HandleExit` | Known | Event handler |
| 0x007C576D | `HandleStartStop` | Known | Event handler |
| 0x007C57D3 | `HandleStartStop` | Known | Event handler |
| 0x007C58EB | `HandleExit` | Known | Event handler |
| 0x007C5915 | `HandleStartStop` | Known | Event handler |
| 0x007C5941 | `HandleLap` | Known | Event handler |
| 0x007C5A45 | `HandleSelectLozinch` | Known | Event handler |
| 0x007C5C62 | `HandleSelect` | Known | Event handler |
| 0x007C5CEE | `HandleSelect` | Known | Event handler |
| 0x007C5D7C | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x007C607A | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x007C6165 | `HandleFinishRecording` | Known | Event handler |
| 0x007C61B6 | `HandlePlayPause` | Known | Event handler |
| 0x007C6244 | `HandlePlayPause` | Known | Event handler |
| 0x007C62D5 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x007C630D | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x007C6349 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x007C638C | `HandlePlayPause` | Known | Event handler |
| 0x007C63C2 | `HandleAddToOTG` | Known | Event handler |
| 0x007C6617 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007C6873 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007E39C6 | `HandleSelectClock` | Known | Event handler |
| 0x007E39FF | `HandleHilited` | Known | Event handler |
| 0x007E3A31 | `HandleWheel` | Known | Event handler |
| 0x007E3A78 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x007E3AFD | `HandleBacksideSongSelected` | Known | Event handler |
| 0x007E3D11 | `HandleImageLast` | Known | Event handler |
| 0x007E3D3B | `HandleScreenNext` | Known | Event handler |
| 0x007E3D6B | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007E3DA5 | `HandleImageFirst` | Known | Event handler |
| 0x007E3DD0 | `HandleScreenPrev` | Known | Event handler |
| 0x007E3DFD | `HandleBrowseLarge` | Known | Event handler |
| 0x007E3E7D | `HandleImageNext` | Known | Event handler |
| 0x007E3EA6 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007E3EDA | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007E3F09 | `HandleImagePrev` | Known | Event handler |
| 0x007E3F37 | `HandleBrowseSmall` | Known | Event handler |

---

## 6. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F5EF8 | `GotoNowPlaying` | Known | Navigation |
| 0x000F5F70 | `GotoMainMenu` | Known | Navigation |
| 0x0010EAEC | `GotoScreen_LockDialog` | Known | Navigation |
| 0x0010EB04 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x0010EC7C | `GotoScreen_AddressBook` | Known | Navigation |
| 0x0011AA8C | `GotoNowPlaying` | Known | Navigation |
| 0x0011AAA0 | `GotoAlbums` | Known | Navigation |
| 0x0011AAAC | `GotoSongs` | Known | Navigation |
| 0x00128A44 | `GotoScreen_EnterPassKey` | Known | Navigation |
| 0x00128A5C | `GotoScreen_LockediPod` | Known | Navigation |
| 0x00129460 | `GotoScreen_MainMenu` | Known | Navigation |
| 0x0013F738 | `GotoMainMenu` | Known | Navigation |
| 0x001C52F0 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001C7780 | `GotoErrorLayout` | Known | Navigation |
| 0x001D0790 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001D0E54 | `GotoPlayDeleteMenu` | Known | Navigation |
| 0x001D0ED8 | `GotoNowPlaying` | Known | Navigation |
| 0x001EBE18 | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation |
| 0x001F7814 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x001F790C | `GotoScreen_SettingsMenu` | Known | Navigation |
| 0x001FF434 | `GotoDefaultLayout` | Known | Navigation |
| 0x001FF4B8 | `GotoVolumeLayout` | Known | Navigation |
| 0x001FF5F0 | `GotoProgressLayout` | Known | Navigation |
| 0x001FF90C | `GotoDefault` | Known | Navigation |
| 0x001FFC40 | `GotoProgressLayout` | Known | Navigation |
| 0x001FFE00 | `GotoRentalWarningLayout` | Known | Navigation |
| 0x001FFE84 | `GotoProgressLayout` | Known | Navigation |
| 0x00200194 | `GotoProgressLayout` | Known | Navigation |
| 0x00201D20 | `GotoNowPlaying` | Known | Navigation |
| 0x00202630 | `GotoNowPlaying` | Known | Navigation |
| 0x0020293C | `GotoNowPlaying` | Known | Navigation |
| 0x00205034 | `GotoScreen_Language` | Known | Navigation |
| 0x00205394 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x002053B0 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x002053C8 | `GotoDefaultLayout` | Known | Navigation |
| 0x002053DC | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00205474 | `GotoVolumeLayout` | Known | Navigation |
| 0x00205488 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x00205528 | `GotoProgressLayout` | Known | Navigation |
| 0x0020553C | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00205CF0 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00206158 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x002063C4 | `GotoProgressLayout` | Known | Navigation |
| 0x002063D8 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00206570 | `GotoBrightnessVideoLayout` | Known | Navigation |
| 0x00206594 | `GotoGeniusLayout` | Known | Navigation |
| 0x002065A8 | `GotoRatingLayout` | Known | Navigation |
| 0x0020671C | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x00206738 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x00206750 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x00206A60 | `GotoChapterArtLayout` | Known | Navigation |
| 0x00206A78 | `GotoShuffleLayout` | Known | Navigation |
| 0x00206E08 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x00206E1C | `GotoExtraInfoLoadingLayout` | Known | Navigation |
| 0x00206EEC | `GotoVolumeLayout` | Known | Navigation |
| 0x00206F04 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x00206F90 | `GotoVolumeLayout` | Known | Navigation |
| 0x00206FA4 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x002071B4 | `GotoScrubLayout` | Known | Navigation |
| 0x002071C4 | `GotoScrubVideoLayout` | Known | Navigation |
| 0x00207254 | `GotoProgressLayout` | Known | Navigation |
| 0x00207268 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x002074C0 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x002074DC | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x002074F4 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00207510 | `GotoDefaultLayout` | Known | Navigation |
| 0x00207744 | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x00207760 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x00207D0C | `GotoChapterArtLayout` | Known | Navigation |
| 0x00207E04 | `GotoProgressLayout` | Known | Navigation |
| 0x00207E90 | `GotoProgressLayout` | Known | Navigation |
| 0x00207EA4 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00207F80 | `GotoExtraInfoLoadFailedLayout` | Known | Navigation |
| 0x00207FA0 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x002083DC | `GotoStatusBarLayout` | Known | Navigation |
| 0x002083F0 | `GotoDefaultLayout` | Known | Navigation |
| 0x002085C8 | `GotoDefault` | Known | Navigation |
| 0x002086FC | `GotoProgressLayout` | Known | Navigation |
| 0x002088BC | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x00208A0C | `GotoBrightnessLayout` | Known | Navigation |
| 0x00208A90 | `GotoBrightnessLayout` | Known | Navigation |
| 0x00208B10 | `GotoVolumeLayout` | Known | Navigation |
| 0x00208B5C | `GotoScrubLayout` | Known | Navigation |
| 0x00208C24 | `GotoStatusBarLayout` | Known | Navigation |
| 0x00208C38 | `GotoDefaultLayout` | Known | Navigation |
| 0x00208D10 | `GotoScrubLayout` | Known | Navigation |
| 0x00208D60 | `GotoScrubLayout` | Known | Navigation |
| 0x0020B880 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020ED18 | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x0020ED34 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020ED4C | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x0020EF18 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020F42C | `GotoNowPlaying` | Known | Navigation |
| 0x0020F728 | `GotoNowPlaying` | Known | Navigation |
| 0x00210884 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x00210A14 | `GotoFourCard_About` | Known | Navigation |
| 0x00210A28 | `GotoThreeCard_About` | Known | Navigation |
| 0x00210B14 | `GotoScreen_BacklightTimer` | Known | Navigation |
| 0x00210BA4 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x00210BBC | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation |
| 0x00215614 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x0021562C | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x002175CC | `GotoGeniusMixesIntro` | Known | Navigation |
| 0x002175E8 | `GotoGeniusMixes` | Known | Navigation |
| 0x00217E00 | `GotoNowPlaying` | Known | Navigation |
| 0x00218510 | `GotoNowPlaying` | Known | Navigation |
| 0x00218CF4 | `GotoFirstBoot` | Known | Navigation |
| 0x00218D04 | `GotoNotesApp` | Known | Navigation |
| 0x00218D18 | `GotoLockApp` | Known | Navigation |
| 0x0021A854 | `GotoGenius` | Known | Navigation |
| 0x0021ACD4 | `GotoGeniusIntro` | Known | Navigation |
| 0x0021ACE8 | `GotoGenius` | Known | Navigation |
| 0x0021D850 | `GotoNowPlaying` | Known | Navigation |
| 0x0021D87C | `GotoGeniusMixLoadingScreen` | Known | Navigation |
| 0x002209D8 | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x002209F4 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x00220A0C | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x00220BCC | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x002213CC | `GotoNowPlaying` | Known | Navigation |
| 0x003F5B34 | `GotoRatingLayout` | Known | Navigation |
| 0x003F5B48 | `GotoProgressLayout` | Known | Navigation |
| 0x0072C727 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x007B337C | `GotoDefault` | Known | Navigation |
| 0x007B4360 | `GotoDefault` | Known | Navigation |
| 0x008A6D68 | `Gotowe?` | Known | Navigation |

---

## 7. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001570AC | `CoverFlow_Screen` | Known | Screen layout |
| 0x0072065E | `Clock_Screen` | Known | Screen layout |
| 0x0072066E | `Clock_Screen_Default"` | Known | Screen layout |
| 0x007206D3 | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x00720731 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x00720749 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x007207B6 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x00720854 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x007208B3 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x007208C9 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x00720934 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x0072098E | `Games_Menu_Screen` | Known | Screen layout |
| 0x007209A3 | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x00720A0D | `Extras_Screen_Games` | Known | Screen layout |
| 0x00720ACC | `Extras_Screen_Notes` | Known | Screen layout |
| 0x00720B90 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00720C59 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x00720CB6 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x00720CCF | `Debug_MainMenu_Screen_Default"` | Known | Screen layout |
| 0x00720D3D | `Extras_Screen_Debug` | Known | Screen layout |
| 0x00720E7C | `Clock_Africa_City_Screen ` | Known | Screen layout |
| 0x00720E98 | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x00720F1C | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x00720F36 | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x00720FB8 | `Clock_Atlantic_City_Screen"` | Known | Screen layout |
| 0x00720FD6 | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x0072105C | `Clock_Australia_City_Screen#` | Known | Screen layout |
| 0x0072107B | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x00721102 | `Clock_Europe_City_Screen ` | Known | Screen layout |
| 0x0072111E | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x007211A2 | `Clock_NorthAmerica_City_Screen&` | Known | Screen layout |
| 0x007211C4 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0072124E | `Clock_Pacific_City_Screen!` | Known | Screen layout |
| 0x0072126B | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x007212F0 | `Clock_SouthAmerica_City_Screen&` | Known | Screen layout |
| 0x00721312 | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0072139F | `Clock_Screen"` | Known | Screen layout |
| 0x00721444 | `Clock_Screen"` | Known | Screen layout |
| 0x007214E9 | `Clock_Screen"` | Known | Screen layout |
| 0x0072158E | `Clock_Screen"` | Known | Screen layout |
| 0x00721633 | `Clock_Screen"` | Known | Screen layout |
| 0x007216D8 | `Clock_Screen"` | Known | Screen layout |
| 0x0072177D | `Clock_Screen"` | Known | Screen layout |
| 0x00721822 | `Clock_Screen"` | Known | Screen layout |
| 0x007218C7 | `Clock_Screen"` | Known | Screen layout |
| 0x0072196C | `Clock_Screen"` | Known | Screen layout |
| 0x00721A11 | `Clock_Screen"` | Known | Screen layout |
| 0x00721AB6 | `Clock_Screen"` | Known | Screen layout |
| 0x00721B5B | `Clock_Screen"` | Known | Screen layout |
| 0x00721C00 | `Clock_Screen"` | Known | Screen layout |
| 0x00721CA5 | `Clock_Screen"` | Known | Screen layout |
| 0x00721D4A | `Clock_Screen"` | Known | Screen layout |
| 0x00721DEF | `Clock_Screen"` | Known | Screen layout |
| 0x00721E94 | `Clock_Screen"` | Known | Screen layout |
| 0x00721F39 | `Clock_Screen"` | Known | Screen layout |
| 0x00721FDE | `Clock_Screen"` | Known | Screen layout |
| 0x00722083 | `Clock_Screen"` | Known | Screen layout |
| 0x00722128 | `Clock_Screen"` | Known | Screen layout |
| 0x007221CD | `Clock_Screen"` | Known | Screen layout |
| 0x00722272 | `Clock_Screen"` | Known | Screen layout |
| 0x00722317 | `Clock_Screen"` | Known | Screen layout |
| 0x007223BC | `Clock_Screen"` | Known | Screen layout |
| 0x00722461 | `Clock_Screen"` | Known | Screen layout |
| 0x00722506 | `Clock_Screen"` | Known | Screen layout |
| 0x007225AB | `Clock_Screen"` | Known | Screen layout |
| 0x00722650 | `Clock_Screen"` | Known | Screen layout |
| 0x007226F5 | `Clock_Screen"` | Known | Screen layout |
| 0x0072279F | `Clock_Screen"` | Known | Screen layout |
| 0x00722844 | `Clock_Screen"` | Known | Screen layout |
| 0x007228E9 | `Clock_Screen"` | Known | Screen layout |
| 0x0072298E | `Clock_Screen"` | Known | Screen layout |
| 0x00722A33 | `Clock_Screen"` | Known | Screen layout |
| 0x00722AD8 | `Clock_Screen"` | Known | Screen layout |
| 0x00722B7D | `Clock_Screen"` | Known | Screen layout |
| 0x00722C22 | `Clock_Screen"` | Known | Screen layout |
| 0x00722CC7 | `Clock_Screen"` | Known | Screen layout |
| 0x00722D6C | `Clock_Screen"` | Known | Screen layout |
| 0x00722E11 | `Clock_Screen"` | Known | Screen layout |
| 0x00722EB6 | `Clock_Screen"` | Known | Screen layout |
| 0x00722F5B | `Clock_Screen"` | Known | Screen layout |
| 0x00723000 | `Clock_Screen"` | Known | Screen layout |
| 0x007230A5 | `Clock_Screen"` | Known | Screen layout |
| 0x0072314A | `Clock_Screen"` | Known | Screen layout |
| 0x007231EF | `Clock_Screen"` | Known | Screen layout |
| 0x00723294 | `Clock_Screen"` | Known | Screen layout |
| 0x00723339 | `Clock_Screen"` | Known | Screen layout |
| 0x007233DE | `Clock_Screen"` | Known | Screen layout |
| 0x00723483 | `Clock_Screen"` | Known | Screen layout |
| 0x00723528 | `Clock_Screen"` | Known | Screen layout |
| 0x007235CD | `Clock_Screen"` | Known | Screen layout |
| 0x00723672 | `Clock_Screen"` | Known | Screen layout |
| 0x00723717 | `Clock_Screen"` | Known | Screen layout |
| 0x007237BC | `Clock_Screen"` | Known | Screen layout |
| 0x00723861 | `Clock_Screen"` | Known | Screen layout |
| 0x00723906 | `Clock_Screen"` | Known | Screen layout |
| 0x007239AB | `Clock_Screen"` | Known | Screen layout |
| 0x00723A50 | `Clock_Screen"` | Known | Screen layout |
| 0x00723AF5 | `Clock_Screen"` | Known | Screen layout |
| 0x00723B9A | `Clock_Screen"` | Known | Screen layout |
| 0x00723C3F | `Clock_Screen"` | Known | Screen layout |
| 0x00723CE4 | `Clock_Screen"` | Known | Screen layout |
| 0x00723D89 | `Clock_Screen"` | Known | Screen layout |
| 0x00723E2E | `Clock_Screen"` | Known | Screen layout |
| 0x00723ED3 | `Clock_Screen"` | Known | Screen layout |
| 0x00723F78 | `Clock_Screen"` | Known | Screen layout |
| 0x0072401D | `Clock_Screen"` | Known | Screen layout |
| 0x007240C2 | `Clock_Screen"` | Known | Screen layout |
| 0x00724167 | `Clock_Screen"` | Known | Screen layout |
| 0x0072420C | `Clock_Screen"` | Known | Screen layout |
| 0x007242B1 | `Clock_Screen"` | Known | Screen layout |
| 0x00724356 | `Clock_Screen"` | Known | Screen layout |
| 0x007243FB | `Clock_Screen"` | Known | Screen layout |
| 0x007244A0 | `Clock_Screen"` | Known | Screen layout |
| 0x00724545 | `Clock_Screen"` | Known | Screen layout |
| 0x007245EA | `Clock_Screen"` | Known | Screen layout |
| 0x0072468F | `Clock_Screen"` | Known | Screen layout |
| 0x00724734 | `Clock_Screen"` | Known | Screen layout |
| 0x007247D9 | `Clock_Screen"` | Known | Screen layout |
| 0x0072487E | `Clock_Screen"` | Known | Screen layout |
| 0x00724923 | `Clock_Screen"` | Known | Screen layout |
| 0x007249C8 | `Clock_Screen"` | Known | Screen layout |
| 0x00724A6D | `Clock_Screen"` | Known | Screen layout |
| 0x00724B12 | `Clock_Screen"` | Known | Screen layout |
| 0x00724BB7 | `Clock_Screen"` | Known | Screen layout |
| 0x00724C63 | `Clock_Screen"` | Known | Screen layout |
| 0x00724D08 | `Clock_Screen"` | Known | Screen layout |
| 0x00724DAD | `Clock_Screen"` | Known | Screen layout |
| 0x00724E57 | `Clock_Screen"` | Known | Screen layout |
| 0x00724EFC | `Clock_Screen"` | Known | Screen layout |
| 0x00724FA1 | `Clock_Screen"` | Known | Screen layout |
| 0x00725046 | `Clock_Screen"` | Known | Screen layout |
| 0x007250EB | `Clock_Screen"` | Known | Screen layout |
| 0x00725190 | `Clock_Screen"` | Known | Screen layout |
| 0x00725235 | `Clock_Screen"` | Known | Screen layout |
| 0x007252DA | `Clock_Screen"` | Known | Screen layout |
| 0x00725383 | `Clock_Screen"` | Known | Screen layout |
| 0x00725428 | `Clock_Screen"` | Known | Screen layout |
| 0x007254CD | `Clock_Screen"` | Known | Screen layout |
| 0x00725572 | `Clock_Screen"` | Known | Screen layout |
| 0x00725617 | `Clock_Screen"` | Known | Screen layout |
| 0x007256BC | `Clock_Screen"` | Known | Screen layout |
| 0x00725761 | `Clock_Screen"` | Known | Screen layout |
| 0x00725806 | `Clock_Screen"` | Known | Screen layout |
| 0x007258AB | `Clock_Screen"` | Known | Screen layout |
| 0x00725950 | `Clock_Screen"` | Known | Screen layout |
| 0x007259F5 | `Clock_Screen"` | Known | Screen layout |
| 0x00725A9A | `Clock_Screen"` | Known | Screen layout |
| 0x00725B3F | `Clock_Screen"` | Known | Screen layout |
| 0x00725BE4 | `Clock_Screen"` | Known | Screen layout |
| 0x00725C89 | `Clock_Screen"` | Known | Screen layout |
| 0x00725D2E | `Clock_Screen"` | Known | Screen layout |
| 0x00725DD3 | `Clock_Screen"` | Known | Screen layout |
| 0x00725E78 | `Clock_Screen"` | Known | Screen layout |
| 0x00725F1D | `Clock_Screen"` | Known | Screen layout |
| 0x00725FC2 | `Clock_Screen"` | Known | Screen layout |
| 0x00726067 | `Clock_Screen"` | Known | Screen layout |
| 0x0072610C | `Clock_Screen"` | Known | Screen layout |
| 0x007261B1 | `Clock_Screen"` | Known | Screen layout |
| 0x00726256 | `Clock_Screen"` | Known | Screen layout |
| 0x007262FB | `Clock_Screen"` | Known | Screen layout |
| 0x007263A0 | `Clock_Screen"` | Known | Screen layout |
| 0x00726445 | `Clock_Screen"` | Known | Screen layout |
| 0x007264EA | `Clock_Screen"` | Known | Screen layout |
| 0x0072658F | `Clock_Screen"` | Known | Screen layout |
| 0x00726634 | `Clock_Screen"` | Known | Screen layout |
| 0x007266D9 | `Clock_Screen"` | Known | Screen layout |
| 0x0072677E | `Clock_Screen"` | Known | Screen layout |
| 0x00726823 | `Clock_Screen"` | Known | Screen layout |
| 0x007268C8 | `Clock_Screen"` | Known | Screen layout |
| 0x00726973 | `Clock_Screen"` | Known | Screen layout |
| 0x00726A18 | `Clock_Screen"` | Known | Screen layout |
| 0x00726ABD | `Clock_Screen"` | Known | Screen layout |
| 0x00726B62 | `Clock_Screen"` | Known | Screen layout |
| 0x00726C07 | `Clock_Screen"` | Known | Screen layout |
| 0x00726CAC | `Clock_Screen"` | Known | Screen layout |
| 0x00726D51 | `Clock_Screen"` | Known | Screen layout |
| 0x00726DF6 | `Clock_Screen"` | Known | Screen layout |
| 0x00726E9B | `Clock_Screen"` | Known | Screen layout |
| 0x00726F40 | `Clock_Screen"` | Known | Screen layout |
| 0x00726FE5 | `Clock_Screen"` | Known | Screen layout |
| 0x0072708A | `Clock_Screen"` | Known | Screen layout |
| 0x0072712F | `Clock_Screen"` | Known | Screen layout |
| 0x007271D4 | `Clock_Screen"` | Known | Screen layout |
| 0x00727279 | `Clock_Screen"` | Known | Screen layout |
| 0x0072731E | `Clock_Screen"` | Known | Screen layout |
| 0x007273C3 | `Clock_Screen"` | Known | Screen layout |
| 0x00727468 | `Clock_Screen"` | Known | Screen layout |
| 0x0072750D | `Clock_Screen"` | Known | Screen layout |
| 0x007275B2 | `Clock_Screen"` | Known | Screen layout |
| 0x00727657 | `Clock_Screen"` | Known | Screen layout |
| 0x007276FC | `Clock_Screen"` | Known | Screen layout |
| 0x007277A1 | `Clock_Screen"` | Known | Screen layout |
| 0x00727846 | `Clock_Screen"` | Known | Screen layout |
| 0x007278EB | `Clock_Screen"` | Known | Screen layout |
| 0x00727990 | `Clock_Screen"` | Known | Screen layout |
| 0x00727A35 | `Clock_Screen"` | Known | Screen layout |
| 0x00727ADA | `Clock_Screen"` | Known | Screen layout |
| 0x00727B7F | `Clock_Screen"` | Known | Screen layout |
| 0x00727C24 | `Clock_Screen"` | Known | Screen layout |
| 0x00727CC9 | `Clock_Screen"` | Known | Screen layout |
| 0x00727D6E | `Clock_Screen"` | Known | Screen layout |
| 0x00727E13 | `Clock_Screen"` | Known | Screen layout |
| 0x00727EB8 | `Clock_Screen"` | Known | Screen layout |
| 0x00727F5D | `Clock_Screen"` | Known | Screen layout |
| 0x00728002 | `Clock_Screen"` | Known | Screen layout |
| 0x007280A7 | `Clock_Screen"` | Known | Screen layout |
| 0x0072814C | `Clock_Screen"` | Known | Screen layout |
| 0x007281F1 | `Clock_Screen"` | Known | Screen layout |
| 0x00728296 | `Clock_Screen"` | Known | Screen layout |
| 0x0072833B | `Clock_Screen"` | Known | Screen layout |
| 0x007283E0 | `Clock_Screen"` | Known | Screen layout |
| 0x00728485 | `Clock_Screen"` | Known | Screen layout |
| 0x0072852A | `Clock_Screen"` | Known | Screen layout |
| 0x007285CF | `Clock_Screen"` | Known | Screen layout |
| 0x00728674 | `Clock_Screen"` | Known | Screen layout |
| 0x00728719 | `Clock_Screen"` | Known | Screen layout |
| 0x007287BE | `Clock_Screen"` | Known | Screen layout |
| 0x00728863 | `Clock_Screen"` | Known | Screen layout |
| 0x00728908 | `Clock_Screen"` | Known | Screen layout |
| 0x007289B3 | `Clock_Screen"` | Known | Screen layout |
| 0x00728A58 | `Clock_Screen"` | Known | Screen layout |
| 0x00728AFD | `Clock_Screen"` | Known | Screen layout |
| 0x00728BA2 | `Clock_Screen"` | Known | Screen layout |
| 0x00728C47 | `Clock_Screen"` | Known | Screen layout |
| 0x00728CF3 | `Clock_Screen"` | Known | Screen layout |
| 0x00728D98 | `Clock_Screen"` | Known | Screen layout |
| 0x00728E3D | `Clock_Screen"` | Known | Screen layout |
| 0x00728EE2 | `Clock_Screen"` | Known | Screen layout |
| 0x00728F87 | `Clock_Screen"` | Known | Screen layout |
| 0x0072902C | `Clock_Screen"` | Known | Screen layout |
| 0x007290D1 | `Clock_Screen"` | Known | Screen layout |
| 0x00729176 | `Clock_Screen"` | Known | Screen layout |
| 0x0072921B | `Clock_Screen"` | Known | Screen layout |
| 0x007292C0 | `Clock_Screen"` | Known | Screen layout |
| 0x00729365 | `Clock_Screen"` | Known | Screen layout |
| 0x0072940A | `Clock_Screen"` | Known | Screen layout |
| 0x007294AF | `Clock_Screen"` | Known | Screen layout |
| 0x00729554 | `Clock_Screen"` | Known | Screen layout |
| 0x007295F9 | `Clock_Screen"` | Known | Screen layout |
| 0x0072969E | `Clock_Screen"` | Known | Screen layout |
| 0x00729743 | `Clock_Screen"` | Known | Screen layout |
| 0x007297E6 | `Settings_DateTime_SetDate_Screen(` | Known | Screen layout |
| 0x0072980A | `Settings_DateTime_SetDate_Screen_Default"` | Known | Screen layout |
| 0x00729883 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x007298E9 | `Settings_DateTime_SetTime_Screen(` | Known | Screen layout |
| 0x0072990D | `Settings_DateTime_SetTime_Screen_Default"` | Known | Screen layout |
| 0x00729986 | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x007299F1 | `Settings_DateTime_SetTimeZone_Screen,` | Known | Screen layout |
| 0x00729A19 | `Settings_DateTime_SetTimeZone_Screen_Default"` | Known | Screen layout |
| 0x00729A96 | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x00729B4F | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x00729BFF | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0072A210 | `Search_Main_Screen` | Known | Screen layout |
| 0x0072A226 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x0072A748 | `Extras_Screen` | Known | Screen layout |
| 0x0072A759 | `Extras_Screen_WorldClock"` | Known | Screen layout |
| 0x0072A7D6 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x0072A838 | `Clock_Screen` | Known | Screen layout |
| 0x0072A848 | `Clock_Screen_Default` | Known | Screen layout |
| 0x0072A8CF | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x0072A935 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0072A94B | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x0072A9B6 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x0072AA18 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0072AA30 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x0072AA9D | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x0072AB01 | `AddressViewer_Main_Screen!` | Known | Screen layout |
| 0x0072AB1E | `AddressViewer_Main_Screen_Default"` | Known | Screen layout |
| 0x0072AB90 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x0072ABF7 | `Games_Menu_Screen` | Known | Screen layout |
| 0x0072AC0C | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x0072AC76 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x0072AD3D | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x0072ADD9 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x0072AEAA | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x0072AF6A | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x0072AFCE | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0072AFED | `VoiceMemos_Menu_Screen_Default"` | Known | Screen layout |
| 0x0072B070 | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x0072B0D6 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x0072B0EE | `Speakers_Main_Screen_Default"` | Known | Screen layout |
| 0x0072B16F | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x0072B1D3 | `Radio_Screen` | Known | Screen layout |
| 0x0072B1E3 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x0072B25C | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x0072B2BD | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0072B359 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x0072B41C | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x0072B4DB | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x0072B598 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x0072B9EB | `Radio_Screen` | Known | Screen layout |
| 0x0072B9FB | `Radio_Screen_Default"` | Known | Screen layout |
| 0x0072BA74 | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x0072BC58 | `Search_Main_Screen` | Known | Screen layout |
| 0x0072BC6E | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x0072BD9C | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x0072BDFF | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x0072C140 | `Video_Settings_Screen` | Known | Screen layout |
| 0x0072C159 | `Video_Settings_Screen_Default"` | Known | Screen layout |
| 0x0072C266 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0072C52B | `SettingsMenus_MainMenu_Screen%` | Known | Screen layout |
| 0x0072C639 | `SettingsMenus_MusicMenu_Screen&` | Known | Screen layout |
| 0x0072C8E2 | `SettingsMenus_Brightness_Screen&` | Known | Screen layout |
| 0x0072C9F7 | `SettingsMenus_AudiobookSettings_Screen.` | Known | Screen layout |
| 0x0072CB2D | `SettingsMenus_RadioRegions_Screen ` | Known | Screen layout |
| 0x0072CC42 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x0072CEAE | `Settings_DateTime_Screen ` | Known | Screen layout |
| 0x0072CECA | `Settings_DateTime_Screen_Default"` | Known | Screen layout |
| 0x0072D056 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x0072D15B | `Settings_Legal_Screen` | Known | Screen layout |
| 0x0072D174 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x0072D265 | `SettingsMenus_ResetAllSettings_Screen4` | Known | Screen layout |
| 0x0072DA36 | `Stopwatch_Screen` | Known | Screen layout |
| 0x0072DA4A | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x0072DAB1 | `Stopwatch_Screen` | Known | Screen layout |
| 0x0072DAC5 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x0072DB6E | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x0072DB91 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x0072DC2A | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x0072DC4D | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x0072DE00 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x0072DE6E | `Speakers_ToneControl_Screen#` | Known | Screen layout |
| 0x0072DE8D | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x00741001 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00741084 | `LockediPod_Screen` | Known | Screen layout |
| 0x0074110C | `Lock_Screen` | Known | Screen layout |
| 0x0074111B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007412B6 | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x00741388 | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x007413F2 | `AddressViewer_ContactDetails_Screen+` | Known | Screen layout |
| 0x00741419 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x00741494 | `Extras_Screen` | Known | Screen layout |
| 0x007414DF | `Extras_Screen` | Known | Screen layout |
| 0x007415C6 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x00741624 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x00741641 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x007416AF | `Calendar_Event_Screen` | Known | Screen layout |
| 0x007416C8 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0074173F | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0074175C | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x007417C7 | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x007417E4 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0074184B | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x007418B2 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x00741910 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0074192D | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x0074199B | `Calendar_Event_Screen` | Known | Screen layout |
| 0x007419B4 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x00741A2B | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x00741A48 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x00741AB3 | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x00741AD0 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x00741B37 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x00741BD7 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x00741C60 | `Alarms_Set_Alarm_Frequency_Screen)` | Known | Screen layout |
| 0x00741C85 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x00741CF6 | `Alarms_Set_Alarm_Label_Screen%` | Known | Screen layout |
| 0x00741D17 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x00741D84 | `Alarms_Set_Alarm_Sound_Screen%` | Known | Screen layout |
| 0x00741DA5 | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x00741E11 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x0074208C | `Alarms_Set_Alarm_Playlist_Screen(` | Known | Screen layout |
| 0x007420B0 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x00742120 | `Alarms_Set_Alarm_Tones_Screen%` | Known | Screen layout |
| 0x00742141 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x00742454 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x0074246F | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x007425C0 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x007425D7 | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x00742658 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0074266F | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x00742745 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0074275E | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x007427E3 | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x00742854 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x00742949 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00742962 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x007429E7 | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x00742A58 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x00742B18 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x00742B2C | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x00742C5B | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x00742CBE | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x00742D15 | `Clock_Screen_Default` | Known | Screen layout |
| 0x00742DA6 | `Clock_Region_Screen` | Known | Screen layout |
| 0x00742DBD | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x00742E36 | `Clock_Screen_Default` | Known | Screen layout |
| 0x00742E8D | `Clock_Screen_Default` | Known | Screen layout |
| 0x00742F1E | `Clock_Region_Screen` | Known | Screen layout |
| 0x00742F35 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x007430C0 | `CoverFlow_Screen_DefaultNoAlbumInfo&` | Known | Screen layout |
| 0x007431AE | `CoverFlow_Screen_DefaultNoAlbumInfo#` | Known | Screen layout |
| 0x00743223 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00743519 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x007436C9 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x007437F7 | `CoverFlow_Screen_Default&` | Known | Screen layout |
| 0x007438CD | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00743A62 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00743CC7 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x00743D24 | `Game_Screen` | Known | Screen layout |
| 0x00743D33 | `Game_Screen_Default` | Known | Screen layout |
| 0x00743DD5 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00743E37 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00743E9A | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00743EFD | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00743F59 | `Game_Running_Screen` | Known | Screen layout |
| 0x00743FB9 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0074401B | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0074407E | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x007440E1 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0074413D | `Game_Running_Screen` | Known | Screen layout |
| 0x0074419D | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x007441FF | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00744262 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x007442C5 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00744321 | `Game_Running_Screen` | Known | Screen layout |
| 0x00744381 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x007443E3 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00744446 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x007444A9 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00744505 | `Game_Running_Screen` | Known | Screen layout |
| 0x00744565 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x007445C7 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0074462A | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0074468D | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x007446E9 | `Game_Running_Screen` | Known | Screen layout |
| 0x0074492F | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00744991 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x007449F4 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00744A57 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00744AB3 | `Game_Running_Screen` | Known | Screen layout |
| 0x00744B6A | `Extras_Screen` | Known | Screen layout |
| 0x00744B7B | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00744BD9 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x00744D76 | `Extras_Screen` | Known | Screen layout |
| 0x00744D87 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00744DE5 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x00744F82 | `Extras_Screen` | Known | Screen layout |
| 0x00744F93 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00744FF1 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0074518E | `Extras_Screen` | Known | Screen layout |
| 0x0074519F | `Extras_Screen_Lock` | Known | Screen layout |
| 0x007451FD | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0074539F | `Lock_Screen` | Known | Screen layout |
| 0x007453AE | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x00745410 | `Extras_Screen` | Known | Screen layout |
| 0x00745421 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x00745480 | `LockediPod_Screen` | Known | Screen layout |
| 0x007454FA | `MainMenus_Main_Screen)` | Known | Screen layout |
| 0x007456CB | `Lock_Screen` | Known | Screen layout |
| 0x007456DA | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0074573C | `Extras_Screen` | Known | Screen layout |
| 0x0074574D | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x007457AC | `LockediPod_Screen` | Known | Screen layout |
| 0x00745826 | `MainMenus_Main_Screen%` | Known | Screen layout |
| 0x0074588D | `LockediPod_Screen` | Known | Screen layout |
| 0x007458A2 | `LockediPod_Screen_Default,` | Known | Screen layout |
| 0x007459F1 | `Lock_Screen` | Known | Screen layout |
| 0x00745A00 | `Lock_Screen_Default_Layout"` | Known | Screen layout |
| 0x00745A69 | `Lock_Screen` | Known | Screen layout |
| 0x00745A78 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x00745ADA | `Extras_Screen` | Known | Screen layout |
| 0x00745AEB | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x00745B4A | `LockediPod_Screen` | Known | Screen layout |
| 0x00745BC4 | `MainMenus_Main_Screen,` | Known | Screen layout |
| 0x00745D20 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00745D86 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x00745DEA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00745E79 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x00745EE6 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x00745F53 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x00745FC0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00746028 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0074608E | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x007460F2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00746181 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x007461EE | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0074625B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x007462C8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00746330 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00746396 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x007463FA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00746489 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x007464F6 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x00746563 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x007465D0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00746638 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0074669E | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x00746702 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00746791 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x007467FE | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0074686B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x007468D8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00746940 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x007469A6 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x00746A0A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00746A99 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x00746B06 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x00746B73 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x00746BE0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00746C39 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x00746CA2 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00746D09 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00746DA4 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00746E0D | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x00746E76 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00746EDD | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00746F78 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00746FE1 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x0074704A | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x007470B1 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0074714C | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00747238 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00747254 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007472C2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007472DF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074734A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074736A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007473E1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007473FD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074746D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074748C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007474F8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074750C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00747585 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007475F9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00747669 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007476D0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00747738 | `NoContent_Screen` | Known | Screen layout |
| 0x0074774C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007477B0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00747817 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00747831 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074789F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00747911 | `NoContent_Screen` | Known | Screen layout |
| 0x00747925 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074798F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007479F8 | `No_Photos_Screen` | Known | Screen layout |
| 0x00747A0C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00747A72 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00747AE0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00747B4D | `NoContent_Screen` | Known | Screen layout |
| 0x00747B61 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00747BC9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00747C33 | `NoContent_Screen` | Known | Screen layout |
| 0x00747C47 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00747CAE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00747D18 | `NoContent_Screen` | Known | Screen layout |
| 0x00747D2C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00747D99 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00747E0B | `NoContent_Screen` | Known | Screen layout |
| 0x00747E1F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00747E87 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00747EF0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00747F0B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00747F71 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00747F8D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074806C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00748085 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007480E6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007480FA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00748154 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00748170 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007481D7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007481EE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00748250 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00748271 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x007482E2 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x007482FE | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00748474 | `Radio_Screen` | Known | Screen layout |
| 0x00748484 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007484E5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00748568 | `LockediPod_Screen` | Known | Screen layout |
| 0x007485F0 | `Lock_Screen` | Known | Screen layout |
| 0x007485FF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00748662 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007486C4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007486E0 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00748752 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00748771 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007487D9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007487F3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074885B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00748878 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007488E4 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074894E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00748968 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007489D8 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00748A4B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00748ABC | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00748B2B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00748B97 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00748BB2 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00748C27 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00748C8E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00748CF0 | `Photos_Screen` | Known | Screen layout |
| 0x00748D55 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00748D74 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00748DDF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00748DFD | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00748E6F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00748E8C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00748EF2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00748F0D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00748F76 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00748F93 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074900A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074902E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074909C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007490B7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00749174 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00749190 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007491FE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074921B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00749286 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007492A6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074931D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00749339 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007493A9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007493C8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00749434 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00749448 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007494C1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00749535 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007495A5 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074960C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00749674 | `NoContent_Screen` | Known | Screen layout |
| 0x00749688 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007496EC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00749753 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074976D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007497DB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074984D | `NoContent_Screen` | Known | Screen layout |
| 0x00749861 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007498CB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00749934 | `No_Photos_Screen` | Known | Screen layout |
| 0x00749948 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007499AE | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00749A1C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00749A89 | `NoContent_Screen` | Known | Screen layout |
| 0x00749A9D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00749B05 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00749B6F | `NoContent_Screen` | Known | Screen layout |
| 0x00749B83 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00749BEA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00749C54 | `NoContent_Screen` | Known | Screen layout |
| 0x00749C68 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00749CD5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00749D47 | `NoContent_Screen` | Known | Screen layout |
| 0x00749D5B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00749DC3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00749E2C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00749E47 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00749EAD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00749EC9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00749FA8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00749FC1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074A022 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074A036 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074A090 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074A0AC | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074A113 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074A12A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074A18C | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0074A1AD | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0074A21E | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0074A23A | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0074A3B0 | `Radio_Screen` | Known | Screen layout |
| 0x0074A3C0 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074A421 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074A4A4 | `LockediPod_Screen` | Known | Screen layout |
| 0x0074A52C | `Lock_Screen` | Known | Screen layout |
| 0x0074A53B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074A59E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074A600 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074A61C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0074A68E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074A6AD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074A715 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074A72F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074A797 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074A7B4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074A820 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074A88A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074A8A4 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074A914 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074A987 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074A9F8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074AA67 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074AAD3 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074AAEE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074AB63 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074ABCA | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074AC2C | `Photos_Screen` | Known | Screen layout |
| 0x0074AC91 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074ACB0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074AD1B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074AD39 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074ADAB | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074ADC8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074AE2E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074AE49 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074AEB2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074AECF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074AF46 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074AF6A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074AFD8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074AFF3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074B0B0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074B0CC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074B13A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074B157 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074B1C2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074B1E2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074B259 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074B275 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074B2E5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074B304 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074B370 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074B384 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074B3FD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074B471 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074B4E1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074B548 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074B5B0 | `NoContent_Screen` | Known | Screen layout |
| 0x0074B5C4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074B628 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074B68F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074B6A9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074B717 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074B789 | `NoContent_Screen` | Known | Screen layout |
| 0x0074B79D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074B807 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074B870 | `No_Photos_Screen` | Known | Screen layout |
| 0x0074B884 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074B8EA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074B958 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074B9C5 | `NoContent_Screen` | Known | Screen layout |
| 0x0074B9D9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074BA41 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074BAAB | `NoContent_Screen` | Known | Screen layout |
| 0x0074BABF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074BB26 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074BB90 | `NoContent_Screen` | Known | Screen layout |
| 0x0074BBA4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074BC11 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074BC83 | `NoContent_Screen` | Known | Screen layout |
| 0x0074BC97 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074BCFF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074BD68 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074BD83 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074BDE9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074BE05 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074BEE4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074BEFD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074BF5E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074BF72 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074BFCC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074BFE8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074C04F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074C066 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074C0C8 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0074C0E9 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0074C15A | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0074C176 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0074C2EC | `Radio_Screen` | Known | Screen layout |
| 0x0074C2FC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074C35D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074C3E0 | `LockediPod_Screen` | Known | Screen layout |
| 0x0074C468 | `Lock_Screen` | Known | Screen layout |
| 0x0074C477 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074C4DA | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074C53C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074C558 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0074C5CA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074C5E9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074C651 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074C66B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074C6D3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074C6F0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074C75C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074C7C6 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074C7E0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074C850 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074C8C3 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074C934 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074C9A3 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074CA0F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074CA2A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074CA9F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074CB06 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074CB68 | `Photos_Screen` | Known | Screen layout |
| 0x0074CBCD | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074CBEC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074CC57 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074CC75 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074CCE7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074CD04 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074CD6A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074CD85 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074CDEE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074CE0B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074CE82 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074CEA6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074CF14 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074CF2F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074CFEC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074D008 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074D076 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074D093 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074D0FE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074D11E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074D195 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074D1B1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074D221 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074D240 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074D2AC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074D2C0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074D339 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074D3AD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074D41D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074D484 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074D4EC | `NoContent_Screen` | Known | Screen layout |
| 0x0074D500 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074D564 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074D5CB | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074D5E5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074D653 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074D6C5 | `NoContent_Screen` | Known | Screen layout |
| 0x0074D6D9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074D743 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074D7AC | `No_Photos_Screen` | Known | Screen layout |
| 0x0074D7C0 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074D826 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074D894 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074D901 | `NoContent_Screen` | Known | Screen layout |
| 0x0074D915 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074D97D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074D9E7 | `NoContent_Screen` | Known | Screen layout |
| 0x0074D9FB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074DA62 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074DACC | `NoContent_Screen` | Known | Screen layout |
| 0x0074DAE0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074DB4D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074DBBF | `NoContent_Screen` | Known | Screen layout |
| 0x0074DBD3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074DC3B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074DCA4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074DCBF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074DD25 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074DD41 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074DE20 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074DE39 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074DE9A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074DEAE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074DF08 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074DF24 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074DF8B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074DFA2 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074E004 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0074E025 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0074E096 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0074E0B2 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0074E228 | `Radio_Screen` | Known | Screen layout |
| 0x0074E238 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074E299 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074E31C | `LockediPod_Screen` | Known | Screen layout |
| 0x0074E3A4 | `Lock_Screen` | Known | Screen layout |
| 0x0074E3B3 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074E416 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074E478 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074E494 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0074E506 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074E525 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074E58D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074E5A7 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074E60F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074E62C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074E698 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074E702 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074E71C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074E78C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074E7FF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074E870 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074E8DF | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074E94B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074E966 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074E9DB | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074EA42 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074EAA4 | `Photos_Screen` | Known | Screen layout |
| 0x0074EB09 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074EB28 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074EB93 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074EBB1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074EC23 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074EC40 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074ECA6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074ECC1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074ED2A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074ED47 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074EDBE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074EDE2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074EE50 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074EE6B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074EF28 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074EF44 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074EFB2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074EFCF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074F03A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074F05A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074F0D1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074F0ED | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074F15D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074F17C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074F1E8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074F1FC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074F275 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074F2E9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074F359 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074F3C0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074F428 | `NoContent_Screen` | Known | Screen layout |
| 0x0074F43C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074F4A0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074F507 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074F521 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074F58F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074F601 | `NoContent_Screen` | Known | Screen layout |
| 0x0074F615 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074F67F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074F6E8 | `No_Photos_Screen` | Known | Screen layout |
| 0x0074F6FC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074F762 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074F7D0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074F83D | `NoContent_Screen` | Known | Screen layout |
| 0x0074F851 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074F8B9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074F923 | `NoContent_Screen` | Known | Screen layout |
| 0x0074F937 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074F99E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074FA08 | `NoContent_Screen` | Known | Screen layout |
| 0x0074FA1C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074FA89 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074FAFB | `NoContent_Screen` | Known | Screen layout |
| 0x0074FB0F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074FB77 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074FBE0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074FBFB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074FC61 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074FC7D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074FD5C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074FD75 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074FDD6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074FDEA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074FE44 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074FE60 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074FEC7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074FEDE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074FF40 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0074FF61 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0074FFD2 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0074FFEE | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00750164 | `Radio_Screen` | Known | Screen layout |
| 0x00750174 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007501D5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00750258 | `LockediPod_Screen` | Known | Screen layout |
| 0x007502E0 | `Lock_Screen` | Known | Screen layout |
| 0x007502EF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00750352 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007503B4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007503D0 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00750442 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00750461 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007504C9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007504E3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075054B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00750568 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007505D4 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075063E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00750658 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007506C8 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075073B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007507AC | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075081B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00750887 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007508A2 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00750917 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075097E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007509E0 | `Photos_Screen` | Known | Screen layout |
| 0x00750A45 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00750A64 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00750ACF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00750AED | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00750B5F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00750B7C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00750BE2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00750BFD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00750C66 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00750C83 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00750CFA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00750D1E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00750D8C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00750DA7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00750E64 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00750E80 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00750EEE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00750F0B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00750F76 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00750F96 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075100D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00751029 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00751099 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007510B8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00751124 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00751138 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007511B1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00751225 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00751295 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007512FC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00751364 | `NoContent_Screen` | Known | Screen layout |
| 0x00751378 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007513DC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00751443 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075145D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007514CB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075153D | `NoContent_Screen` | Known | Screen layout |
| 0x00751551 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007515BB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00751624 | `No_Photos_Screen` | Known | Screen layout |
| 0x00751638 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075169E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075170C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00751779 | `NoContent_Screen` | Known | Screen layout |
| 0x0075178D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007517F5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075185F | `NoContent_Screen` | Known | Screen layout |
| 0x00751873 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007518DA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00751944 | `NoContent_Screen` | Known | Screen layout |
| 0x00751958 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007519C5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00751A37 | `NoContent_Screen` | Known | Screen layout |
| 0x00751A4B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00751AB3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00751B1C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00751B37 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00751B9D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00751BB9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00751C98 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00751CB1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00751D12 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00751D26 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00751D80 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00751D9C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00751E03 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00751E1A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00751E7C | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00751E9D | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00751F0E | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00751F2A | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x007520A0 | `Radio_Screen` | Known | Screen layout |
| 0x007520B0 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00752111 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00752194 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075221C | `Lock_Screen` | Known | Screen layout |
| 0x0075222B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075228E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007522F0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075230C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075237E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075239D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00752405 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075241F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00752487 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007524A4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00752510 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075257A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00752594 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00752604 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00752677 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007526E8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00752757 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007527C3 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007527DE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00752853 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007528BA | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075291C | `Photos_Screen` | Known | Screen layout |
| 0x00752981 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007529A0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00752A0B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00752A29 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00752A9B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00752AB8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00752B1E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00752B39 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00752BA2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00752BBF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00752C36 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00752C5A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00752CC8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00752CE3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00752DA0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00752DBC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00752E2A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00752E47 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00752EB2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00752ED2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00752F49 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00752F65 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00752FD5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00752FF4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00753060 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00753074 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007530ED | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00753161 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007531D1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00753238 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007532A0 | `NoContent_Screen` | Known | Screen layout |
| 0x007532B4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00753318 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075337F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00753399 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00753407 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00753479 | `NoContent_Screen` | Known | Screen layout |
| 0x0075348D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007534F7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00753560 | `No_Photos_Screen` | Known | Screen layout |
| 0x00753574 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007535DA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00753648 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007536B5 | `NoContent_Screen` | Known | Screen layout |
| 0x007536C9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00753731 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075379B | `NoContent_Screen` | Known | Screen layout |
| 0x007537AF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00753816 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00753880 | `NoContent_Screen` | Known | Screen layout |
| 0x00753894 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00753901 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00753973 | `NoContent_Screen` | Known | Screen layout |
| 0x00753987 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007539EF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00753A58 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00753A73 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00753AD9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00753AF5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00753BD4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00753BED | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00753C4E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00753C62 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00753CBC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00753CD8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00753D3F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00753D56 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00753DB8 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00753DD9 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00753E4A | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00753E66 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00753FDC | `Radio_Screen` | Known | Screen layout |
| 0x00753FEC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075404D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007540D0 | `LockediPod_Screen` | Known | Screen layout |
| 0x00754158 | `Lock_Screen` | Known | Screen layout |
| 0x00754167 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007541CA | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075422C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00754248 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007542BA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007542D9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00754341 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075435B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007543C3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007543E0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075444C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007544B6 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007544D0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00754540 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007545B3 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00754624 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00754693 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007546FF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075471A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075478F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007547F6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00754858 | `Photos_Screen` | Known | Screen layout |
| 0x007548BD | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007548DC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00754947 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00754965 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007549D7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007549F4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00754A5A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00754A75 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00754ADE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00754AFB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00754B72 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00754B96 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00754C04 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00754C1F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00754CDC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00754CF8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00754D66 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00754D83 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00754DEE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00754E0E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00754E85 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00754EA1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00754F11 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00754F30 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00754F9C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00754FB0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00755029 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075509D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075510D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00755174 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007551DC | `NoContent_Screen` | Known | Screen layout |
| 0x007551F0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00755254 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007552BB | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007552D5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00755343 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007553B5 | `NoContent_Screen` | Known | Screen layout |
| 0x007553C9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00755433 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075549C | `No_Photos_Screen` | Known | Screen layout |
| 0x007554B0 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00755516 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00755584 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007555F1 | `NoContent_Screen` | Known | Screen layout |
| 0x00755605 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075566D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007556D7 | `NoContent_Screen` | Known | Screen layout |
| 0x007556EB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00755752 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007557BC | `NoContent_Screen` | Known | Screen layout |
| 0x007557D0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075583D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007558AF | `NoContent_Screen` | Known | Screen layout |
| 0x007558C3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075592B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00755994 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007559AF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00755A15 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00755A31 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00755B10 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00755B29 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00755B8A | `FirstBoot_Screen` | Known | Screen layout |
| 0x00755B9E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00755BF8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00755C14 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00755C7B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00755C92 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00755CF4 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00755D15 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00755D86 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00755DA2 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00755F18 | `Radio_Screen` | Known | Screen layout |
| 0x00755F28 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00755F89 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075600C | `LockediPod_Screen` | Known | Screen layout |
| 0x00756094 | `Lock_Screen` | Known | Screen layout |
| 0x007560A3 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00756106 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00756168 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00756184 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007561F6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00756215 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075627D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00756297 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007562FF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075631C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00756388 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007563F2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075640C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075647C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007564EF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00756560 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007565CF | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075663B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00756656 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007566CB | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00756732 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00756794 | `Photos_Screen` | Known | Screen layout |
| 0x007567F9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00756818 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00756883 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007568A1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00756913 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00756930 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00756996 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007569B1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00756A1A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00756A37 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00756AAE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00756AD2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00756B40 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00756B5B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00756C18 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00756C34 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00756CA2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00756CBF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00756D2A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00756D4A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00756DC1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00756DDD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00756E4D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00756E6C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00756ED8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00756EEC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00756F65 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00756FD9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00757049 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007570B0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00757118 | `NoContent_Screen` | Known | Screen layout |
| 0x0075712C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00757190 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007571F7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00757211 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075727F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007572F1 | `NoContent_Screen` | Known | Screen layout |
| 0x00757305 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075736F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007573D8 | `No_Photos_Screen` | Known | Screen layout |
| 0x007573EC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00757452 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007574C0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075752D | `NoContent_Screen` | Known | Screen layout |
| 0x00757541 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007575A9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00757613 | `NoContent_Screen` | Known | Screen layout |
| 0x00757627 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075768E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007576F8 | `NoContent_Screen` | Known | Screen layout |
| 0x0075770C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00757779 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007577EB | `NoContent_Screen` | Known | Screen layout |
| 0x007577FF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00757867 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007578D0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007578EB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00757951 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075796D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00757A4C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00757A65 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00757AC6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00757ADA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00757B34 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00757B50 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00757BB7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00757BCE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00757C30 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00757C51 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00757CC2 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00757CDE | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00757E54 | `Radio_Screen` | Known | Screen layout |
| 0x00757E64 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00757EC5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00757F48 | `LockediPod_Screen` | Known | Screen layout |
| 0x00757FD0 | `Lock_Screen` | Known | Screen layout |
| 0x00757FDF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00758042 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007580A4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007580C0 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00758132 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00758151 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007581B9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007581D3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075823B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00758258 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007582C4 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075832E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00758348 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007583B8 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075842B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075849C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075850B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00758577 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00758592 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00758607 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075866E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007586D0 | `Photos_Screen` | Known | Screen layout |
| 0x00758735 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00758754 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007587BF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007587DD | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075884F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075886C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007588D2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007588ED | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00758956 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00758973 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007589EA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00758A0E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00758A7C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00758A97 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00758B54 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00758B70 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00758BDE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00758BFB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00758C66 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00758C86 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00758CFD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00758D19 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00758D89 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00758DA8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00758E14 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00758E28 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00758EA1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00758F15 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00758F85 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00758FEC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00759054 | `NoContent_Screen` | Known | Screen layout |
| 0x00759068 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007590CC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00759133 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075914D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007591BB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075922D | `NoContent_Screen` | Known | Screen layout |
| 0x00759241 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007592AB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00759314 | `No_Photos_Screen` | Known | Screen layout |
| 0x00759328 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075938E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007593FC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00759469 | `NoContent_Screen` | Known | Screen layout |
| 0x0075947D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007594E5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075954F | `NoContent_Screen` | Known | Screen layout |
| 0x00759563 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007595CA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00759634 | `NoContent_Screen` | Known | Screen layout |
| 0x00759648 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007596B5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00759727 | `NoContent_Screen` | Known | Screen layout |
| 0x0075973B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007597A3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075980C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00759827 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075988D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007598A9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00759988 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007599A1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00759A02 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00759A16 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00759A70 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00759A8C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00759AF3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00759B0A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00759B6C | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00759B8D | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00759BFE | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00759C1A | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00759D90 | `Radio_Screen` | Known | Screen layout |
| 0x00759DA0 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00759E01 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00759E84 | `LockediPod_Screen` | Known | Screen layout |
| 0x00759F0C | `Lock_Screen` | Known | Screen layout |
| 0x00759F1B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00759F7E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00759FE0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00759FFC | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075A06E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075A08D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075A0F5 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075A10F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075A177 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075A194 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075A200 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075A26A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075A284 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075A2F4 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075A367 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075A3D8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075A447 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075A4B3 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075A4CE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075A543 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075A5AA | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075A60C | `Photos_Screen` | Known | Screen layout |
| 0x0075A671 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075A690 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075A6FB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075A719 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075A78B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075A7A8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075A80E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075A829 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075A892 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075A8AF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075A926 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075A94A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075A9B8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075A9D3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075AA90 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075AAAC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075AB1A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075AB37 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075ABA2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075ABC2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075AC39 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075AC55 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075ACC5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075ACE4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075AD50 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075AD64 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075ADDD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075AE51 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075AEC1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075AF28 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075AF90 | `NoContent_Screen` | Known | Screen layout |
| 0x0075AFA4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075B008 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075B06F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075B089 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075B0F7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075B169 | `NoContent_Screen` | Known | Screen layout |
| 0x0075B17D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075B1E7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075B250 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075B264 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075B2CA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075B338 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075B3A5 | `NoContent_Screen` | Known | Screen layout |
| 0x0075B3B9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075B421 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075B48B | `NoContent_Screen` | Known | Screen layout |
| 0x0075B49F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075B506 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075B570 | `NoContent_Screen` | Known | Screen layout |
| 0x0075B584 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075B5F1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075B663 | `NoContent_Screen` | Known | Screen layout |
| 0x0075B677 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075B6DF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075B748 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075B763 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075B7C9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075B7E5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075B8C4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075B8DD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075B93E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075B952 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075B9AC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075B9C8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075BA2F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075BA46 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075BAA8 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0075BAC9 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0075BB3A | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0075BB56 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0075BCCC | `Radio_Screen` | Known | Screen layout |
| 0x0075BCDC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075BD3D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075BDC0 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075BE48 | `Lock_Screen` | Known | Screen layout |
| 0x0075BE57 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075BEBA | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075BF1C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075BF38 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075BFAA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075BFC9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075C031 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075C04B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075C0B3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075C0D0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075C13C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075C1A6 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075C1C0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075C230 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075C2A3 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075C314 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075C383 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075C3EF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075C40A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075C47F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075C4E6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075C548 | `Photos_Screen` | Known | Screen layout |
| 0x0075C5AD | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075C5CC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075C637 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075C655 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075C6C7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075C6E4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075C74A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075C765 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075C7CE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075C7EB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075C862 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075C886 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075C8F4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075C90F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075C9CC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075C9E8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075CA56 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075CA73 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075CADE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075CAFE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075CB75 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075CB91 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075CC01 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075CC20 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075CC8C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075CCA0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075CD19 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075CD8D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075CDFD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075CE64 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075CECC | `NoContent_Screen` | Known | Screen layout |
| 0x0075CEE0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075CF44 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075CFAB | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075CFC5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075D033 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075D0A5 | `NoContent_Screen` | Known | Screen layout |
| 0x0075D0B9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075D123 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075D18C | `No_Photos_Screen` | Known | Screen layout |
| 0x0075D1A0 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075D206 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075D274 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075D2E1 | `NoContent_Screen` | Known | Screen layout |
| 0x0075D2F5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075D35D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075D3C7 | `NoContent_Screen` | Known | Screen layout |
| 0x0075D3DB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075D442 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075D4AC | `NoContent_Screen` | Known | Screen layout |
| 0x0075D4C0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075D52D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075D59F | `NoContent_Screen` | Known | Screen layout |
| 0x0075D5B3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075D61B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075D684 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075D69F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075D705 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075D721 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075D800 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075D819 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075D87A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075D88E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075D8E8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075D904 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075D96B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075D982 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075D9E4 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0075DA05 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0075DA76 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0075DA92 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0075DC08 | `Radio_Screen` | Known | Screen layout |
| 0x0075DC18 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075DC79 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075DCFC | `LockediPod_Screen` | Known | Screen layout |
| 0x0075DD84 | `Lock_Screen` | Known | Screen layout |
| 0x0075DD93 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075DDF6 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075DE58 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075DE74 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075DEE6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075DF05 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075DF6D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075DF87 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075DFEF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075E00C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075E078 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075E0E2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075E0FC | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075E16C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075E1DF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075E250 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075E2BF | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075E32B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075E346 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075E3BB | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075E422 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075E484 | `Photos_Screen` | Known | Screen layout |
| 0x0075E4E9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075E508 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075E573 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075E591 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075E603 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075E620 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075E686 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075E6A1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075E70A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075E727 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075E79E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075E7C2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075E830 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075E84B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075E908 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075E924 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075E992 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075E9AF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075EA1A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075EA3A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075EAB1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075EACD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075EB3D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075EB5C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075EBC8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075EBDC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075EC55 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075ECC9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075ED39 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075EDA0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075EE08 | `NoContent_Screen` | Known | Screen layout |
| 0x0075EE1C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075EE80 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075EEE7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075EF01 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075EF6F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075EFE1 | `NoContent_Screen` | Known | Screen layout |
| 0x0075EFF5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075F05F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075F0C8 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075F0DC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075F142 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075F1B0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075F21D | `NoContent_Screen` | Known | Screen layout |
| 0x0075F231 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075F299 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075F303 | `NoContent_Screen` | Known | Screen layout |
| 0x0075F317 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075F37E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075F3E8 | `NoContent_Screen` | Known | Screen layout |
| 0x0075F3FC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075F469 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075F4DB | `NoContent_Screen` | Known | Screen layout |
| 0x0075F4EF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075F557 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075F5C0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075F5DB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075F641 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075F65D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075F73C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075F755 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075F7B6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075F7CA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075F824 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075F840 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075F8A7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075F8BE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075F920 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0075F941 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0075F9B2 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0075F9CE | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0075FB44 | `Radio_Screen` | Known | Screen layout |
| 0x0075FB54 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075FBB5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075FC38 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075FCC0 | `Lock_Screen` | Known | Screen layout |
| 0x0075FCCF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075FD32 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075FD94 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075FDB0 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075FE22 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075FE41 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075FEA9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075FEC3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075FF2B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075FF48 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075FFB4 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076001E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00760038 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007600A8 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076011B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076018C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007601FB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00760267 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00760282 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007602F7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076035E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007603C0 | `Photos_Screen` | Known | Screen layout |
| 0x00760425 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00760444 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007604AF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007604CD | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076053F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076055C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007605C2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007605DD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00760646 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00760663 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007606DA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007606FE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076076C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00760787 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00760844 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00760860 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007608CE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007608EB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00760956 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00760976 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007609ED | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00760A09 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00760A79 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00760A98 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00760B04 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00760B18 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00760B91 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00760C05 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00760C75 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00760CDC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00760D44 | `NoContent_Screen` | Known | Screen layout |
| 0x00760D58 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00760DBC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00760E23 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00760E3D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00760EAB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00760F1D | `NoContent_Screen` | Known | Screen layout |
| 0x00760F31 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00760F9B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00761004 | `No_Photos_Screen` | Known | Screen layout |
| 0x00761018 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076107E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007610EC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00761159 | `NoContent_Screen` | Known | Screen layout |
| 0x0076116D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007611D5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076123F | `NoContent_Screen` | Known | Screen layout |
| 0x00761253 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007612BA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00761324 | `NoContent_Screen` | Known | Screen layout |
| 0x00761338 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007613A5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00761417 | `NoContent_Screen` | Known | Screen layout |
| 0x0076142B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00761493 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007614FC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00761517 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076157D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00761599 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00761678 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00761691 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007616F2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00761706 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00761760 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076177C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007617E3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007617FA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076185C | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0076187D | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x007618EE | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0076190A | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00761A80 | `Radio_Screen` | Known | Screen layout |
| 0x00761A90 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00761AF1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00761B74 | `LockediPod_Screen` | Known | Screen layout |
| 0x00761BFC | `Lock_Screen` | Known | Screen layout |
| 0x00761C0B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00761C6E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00761CD0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00761CEC | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00761D5E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00761D7D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00761DE5 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00761DFF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00761E67 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00761E84 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00761EF0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00761F5A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00761F74 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00761FE4 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00762057 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007620C8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00762137 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007621A3 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007621BE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00762233 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076229A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007622FC | `Photos_Screen` | Known | Screen layout |
| 0x00762361 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00762380 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007623EB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00762409 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076247B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00762498 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007624FE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00762519 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00762582 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076259F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00762616 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076263A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007626A8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007626C3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00762780 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076279C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076280A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00762827 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00762892 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007628B2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00762929 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00762945 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007629B5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007629D4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00762A40 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00762A54 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00762ACD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00762B41 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00762BB1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00762C18 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00762C80 | `NoContent_Screen` | Known | Screen layout |
| 0x00762C94 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00762CF8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00762D5F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00762D79 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00762DE7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00762E59 | `NoContent_Screen` | Known | Screen layout |
| 0x00762E6D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00762ED7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00762F40 | `No_Photos_Screen` | Known | Screen layout |
| 0x00762F54 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00762FBA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00763028 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00763095 | `NoContent_Screen` | Known | Screen layout |
| 0x007630A9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00763111 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076317B | `NoContent_Screen` | Known | Screen layout |
| 0x0076318F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007631F6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00763260 | `NoContent_Screen` | Known | Screen layout |
| 0x00763274 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007632E1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00763353 | `NoContent_Screen` | Known | Screen layout |
| 0x00763367 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007633CF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00763438 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00763453 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007634B9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007634D5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007635B4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007635CD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076362E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00763642 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076369C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007636B8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076371F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00763736 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00763798 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x007637B9 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0076382A | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00763846 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x007639BC | `Radio_Screen` | Known | Screen layout |
| 0x007639CC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00763A2D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00763AB0 | `LockediPod_Screen` | Known | Screen layout |
| 0x00763B38 | `Lock_Screen` | Known | Screen layout |
| 0x00763B47 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00763BAA | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00763C0C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00763C28 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00763C9A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00763CB9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00763D21 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00763D3B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00763DA3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00763DC0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00763E2C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00763E96 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00763EB0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00763F20 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00763F93 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00764004 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00764073 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007640DF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007640FA | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076416F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007641D6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00764238 | `Photos_Screen` | Known | Screen layout |
| 0x0076429D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007642BC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00764327 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00764345 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007643B7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007643D4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076443A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00764455 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007644BE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007644DB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00764552 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00764576 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007645E4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007645FF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007646BC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007646D8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00764746 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00764763 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007647CE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007647EE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00764865 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00764881 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007648F1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00764910 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076497C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00764990 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00764A09 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00764A7D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00764AED | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00764B54 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00764BBC | `NoContent_Screen` | Known | Screen layout |
| 0x00764BD0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00764C34 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00764C9B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00764CB5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00764D23 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00764D95 | `NoContent_Screen` | Known | Screen layout |
| 0x00764DA9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00764E13 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00764E7C | `No_Photos_Screen` | Known | Screen layout |
| 0x00764E90 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00764EF6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00764F64 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00764FD1 | `NoContent_Screen` | Known | Screen layout |
| 0x00764FE5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076504D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007650B7 | `NoContent_Screen` | Known | Screen layout |
| 0x007650CB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00765132 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076519C | `NoContent_Screen` | Known | Screen layout |
| 0x007651B0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076521D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076528F | `NoContent_Screen` | Known | Screen layout |
| 0x007652A3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076530B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00765374 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076538F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007653F5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00765411 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007654F0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00765509 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076556A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076557E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007655D8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007655F4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076565B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00765672 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007656D4 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x007656F5 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00765766 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00765782 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x007658F8 | `Radio_Screen` | Known | Screen layout |
| 0x00765908 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00765969 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007659EC | `LockediPod_Screen` | Known | Screen layout |
| 0x00765A74 | `Lock_Screen` | Known | Screen layout |
| 0x00765A83 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00765AE6 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00765B48 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00765B64 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00765BD6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00765BF5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00765C5D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00765C77 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00765CDF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00765CFC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00765D68 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00765DD2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00765DEC | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00765E5C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00765ECF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00765F40 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00765FAF | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076601B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00766036 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007660AB | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00766112 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00766174 | `Photos_Screen` | Known | Screen layout |
| 0x007661D9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007661F8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00766263 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00766281 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007662F3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00766310 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00766376 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00766391 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007663FA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00766417 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076648E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007664B2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00766520 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076653B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007665F8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00766614 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00766682 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076669F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076670A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076672A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007667A1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007667BD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076682D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076684C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007668B8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007668CC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00766945 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007669B9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00766A29 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00766A90 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00766AF8 | `NoContent_Screen` | Known | Screen layout |
| 0x00766B0C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00766B70 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00766BD7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00766BF1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00766C5F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00766CD1 | `NoContent_Screen` | Known | Screen layout |
| 0x00766CE5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00766D4F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00766DB8 | `No_Photos_Screen` | Known | Screen layout |
| 0x00766DCC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00766E32 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00766EA0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00766F0D | `NoContent_Screen` | Known | Screen layout |
| 0x00766F21 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00766F89 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00766FF3 | `NoContent_Screen` | Known | Screen layout |
| 0x00767007 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076706E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007670D8 | `NoContent_Screen` | Known | Screen layout |
| 0x007670EC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00767159 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007671CB | `NoContent_Screen` | Known | Screen layout |
| 0x007671DF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00767247 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007672B0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007672CB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00767331 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076734D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076742C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00767445 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007674A6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007674BA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00767514 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00767530 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00767597 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007675AE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00767610 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00767631 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x007676A2 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x007676BE | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00767834 | `Radio_Screen` | Known | Screen layout |
| 0x00767844 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007678A5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00767928 | `LockediPod_Screen` | Known | Screen layout |
| 0x007679B0 | `Lock_Screen` | Known | Screen layout |
| 0x007679BF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00767A22 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00767A84 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00767AA0 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00767B12 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00767B31 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00767B99 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00767BB3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00767C1B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00767C38 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00767CA4 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00767D0E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00767D28 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00767D98 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00767E0B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00767E7C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00767EEB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00767F57 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00767F72 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00767FE7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076804E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007680B0 | `Photos_Screen` | Known | Screen layout |
| 0x00768115 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00768134 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076819F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007681BD | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076822F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076824C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007682B2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007682CD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00768336 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00768353 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007683CA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007683EE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076845C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00768477 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00768534 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00768550 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007685BE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007685DB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00768646 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00768666 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007686DD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007686F9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00768769 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00768788 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007687F4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00768808 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00768881 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007688F5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00768965 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007689CC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00768A34 | `NoContent_Screen` | Known | Screen layout |
| 0x00768A48 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00768AAC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00768B13 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00768B2D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00768B9B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00768C0D | `NoContent_Screen` | Known | Screen layout |
| 0x00768C21 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00768C8B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00768CF4 | `No_Photos_Screen` | Known | Screen layout |
| 0x00768D08 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00768D6E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00768DDC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00768E49 | `NoContent_Screen` | Known | Screen layout |
| 0x00768E5D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00768EC5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00768F2F | `NoContent_Screen` | Known | Screen layout |
| 0x00768F43 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00768FAA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00769014 | `NoContent_Screen` | Known | Screen layout |
| 0x00769028 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00769095 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00769107 | `NoContent_Screen` | Known | Screen layout |
| 0x0076911B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00769183 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007691EC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00769207 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076926D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00769289 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00769368 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00769381 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007693E2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007693F6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00769450 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076946C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007694D3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007694EA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076954C | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0076956D | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x007695DE | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x007695FA | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00769770 | `Radio_Screen` | Known | Screen layout |
| 0x00769780 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007697E1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00769864 | `LockediPod_Screen` | Known | Screen layout |
| 0x007698EC | `Lock_Screen` | Known | Screen layout |
| 0x007698FB | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076995E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007699C0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007699DC | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00769A4E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00769A6D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00769AD5 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00769AEF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00769B57 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00769B74 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00769BE0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00769C4A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00769C64 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00769CD4 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00769D47 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00769DB8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00769E27 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00769E93 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00769EAE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00769F23 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00769F8A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00769FEC | `Photos_Screen` | Known | Screen layout |
| 0x0076A051 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076A070 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076A0DB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076A0F9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076A16B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076A188 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076A1EE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076A209 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076A272 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076A28F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076A306 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076A32A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076A398 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076A3B3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076A470 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076A48C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076A4FA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076A517 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076A582 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076A5A2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076A619 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076A635 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076A6A5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076A6C4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076A730 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076A744 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076A7BD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076A831 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076A8A1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076A908 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076A970 | `NoContent_Screen` | Known | Screen layout |
| 0x0076A984 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076A9E8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076AA4F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076AA69 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076AAD7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076AB49 | `NoContent_Screen` | Known | Screen layout |
| 0x0076AB5D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076ABC7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076AC30 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076AC44 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076ACAA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076AD18 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076AD85 | `NoContent_Screen` | Known | Screen layout |
| 0x0076AD99 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076AE01 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076AE6B | `NoContent_Screen` | Known | Screen layout |
| 0x0076AE7F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076AEE6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076AF50 | `NoContent_Screen` | Known | Screen layout |
| 0x0076AF64 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076AFD1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076B043 | `NoContent_Screen` | Known | Screen layout |
| 0x0076B057 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076B0BF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076B128 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076B143 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076B1A9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076B1C5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076B2A4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076B2BD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076B31E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076B332 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076B38C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076B3A8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076B40F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076B426 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076B488 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0076B4A9 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0076B51A | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0076B536 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0076B6AC | `Radio_Screen` | Known | Screen layout |
| 0x0076B6BC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076B71D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076B7A0 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076B828 | `Lock_Screen` | Known | Screen layout |
| 0x0076B837 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076B89A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076B8FC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076B918 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076B98A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076B9A9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076BA11 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076BA2B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076BA93 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076BAB0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076BB1C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076BB86 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076BBA0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076BC10 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076BC83 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076BCF4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076BD63 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076BDCF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076BDEA | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076BE5F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076BEC6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076BF28 | `Photos_Screen` | Known | Screen layout |
| 0x0076BF8D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076BFAC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076C017 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076C035 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076C0A7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076C0C4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076C12A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076C145 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076C1AE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076C1CB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076C242 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076C266 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076C2D4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076C2EF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076C3AC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076C3C8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076C436 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076C453 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076C4BE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076C4DE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076C555 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076C571 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076C5E1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076C600 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076C66C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076C680 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076C6F9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076C76D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076C7DD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076C844 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076C8AC | `NoContent_Screen` | Known | Screen layout |
| 0x0076C8C0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076C924 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076C98B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076C9A5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076CA13 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076CA85 | `NoContent_Screen` | Known | Screen layout |
| 0x0076CA99 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076CB03 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076CB6C | `No_Photos_Screen` | Known | Screen layout |
| 0x0076CB80 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076CBE6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076CC54 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076CCC1 | `NoContent_Screen` | Known | Screen layout |
| 0x0076CCD5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076CD3D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076CDA7 | `NoContent_Screen` | Known | Screen layout |
| 0x0076CDBB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076CE22 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076CE8C | `NoContent_Screen` | Known | Screen layout |
| 0x0076CEA0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076CF0D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076CF7F | `NoContent_Screen` | Known | Screen layout |
| 0x0076CF93 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076CFFB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076D064 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076D07F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076D0E5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076D101 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076D1E0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076D1F9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076D25A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076D26E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076D2C8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076D2E4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076D34B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076D362 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076D3C4 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0076D3E5 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0076D456 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0076D472 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0076D5E8 | `Radio_Screen` | Known | Screen layout |
| 0x0076D5F8 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076D659 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076D6DC | `LockediPod_Screen` | Known | Screen layout |
| 0x0076D764 | `Lock_Screen` | Known | Screen layout |
| 0x0076D773 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076D7D6 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076D838 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076D854 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076D8C6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076D8E5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076D94D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076D967 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076D9CF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076D9EC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076DA58 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076DAC2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076DADC | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076DB4C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076DBBF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076DC30 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076DC9F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076DD0B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076DD26 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076DD9B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076DE02 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076DE64 | `Photos_Screen` | Known | Screen layout |
| 0x0076DEC9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076DEE8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076DF53 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076DF71 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076DFE3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076E000 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076E066 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076E081 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076E0EA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076E107 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076E17E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076E1A2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076E210 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076E22B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076E2E8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076E304 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076E372 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076E38F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076E3FA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076E41A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076E491 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076E4AD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076E51D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076E53C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076E5A8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076E5BC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076E635 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076E6A9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076E719 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076E780 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076E7E8 | `NoContent_Screen` | Known | Screen layout |
| 0x0076E7FC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076E860 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076E8C7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076E8E1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076E94F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076E9C1 | `NoContent_Screen` | Known | Screen layout |
| 0x0076E9D5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076EA3F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076EAA8 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076EABC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076EB22 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076EB90 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076EBFD | `NoContent_Screen` | Known | Screen layout |
| 0x0076EC11 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076EC79 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076ECE3 | `NoContent_Screen` | Known | Screen layout |
| 0x0076ECF7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076ED5E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076EDC8 | `NoContent_Screen` | Known | Screen layout |
| 0x0076EDDC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076EE49 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076EEBB | `NoContent_Screen` | Known | Screen layout |
| 0x0076EECF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076EF37 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076EFA0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076EFBB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076F021 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076F03D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076F11C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076F135 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076F196 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076F1AA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076F204 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076F220 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076F287 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076F29E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076F300 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0076F321 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0076F392 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0076F3AE | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0076F524 | `Radio_Screen` | Known | Screen layout |
| 0x0076F534 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076F595 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076F618 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076F6A0 | `Lock_Screen` | Known | Screen layout |
| 0x0076F6AF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076F712 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076F774 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076F790 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076F802 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076F821 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076F889 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076F8A3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076F90B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076F928 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076F994 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076F9FE | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076FA18 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076FA88 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076FAFB | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076FB6C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076FBDB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076FC47 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076FC62 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076FCD7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076FD3E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076FDA0 | `Photos_Screen` | Known | Screen layout |
| 0x0076FE05 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076FE24 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076FE8F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076FEAD | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076FF1F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076FF3C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076FFA2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076FFBD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00770026 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00770043 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007700BA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007700DE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077014C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00770167 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00770224 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00770240 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007702AE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007702CB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00770336 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00770356 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007703CD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007703E9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00770459 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00770478 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007704E4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007704F8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00770571 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007705E5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00770655 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007706BC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00770724 | `NoContent_Screen` | Known | Screen layout |
| 0x00770738 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077079C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00770803 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077081D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077088B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007708FD | `NoContent_Screen` | Known | Screen layout |
| 0x00770911 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077097B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007709E4 | `No_Photos_Screen` | Known | Screen layout |
| 0x007709F8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00770A5E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00770ACC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00770B39 | `NoContent_Screen` | Known | Screen layout |
| 0x00770B4D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00770BB5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00770C1F | `NoContent_Screen` | Known | Screen layout |
| 0x00770C33 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00770C9A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00770D04 | `NoContent_Screen` | Known | Screen layout |
| 0x00770D18 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00770D85 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00770DF7 | `NoContent_Screen` | Known | Screen layout |
| 0x00770E0B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00770E73 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00770EDC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00770EF7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00770F5D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00770F79 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00771058 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00771071 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007710D2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007710E6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00771140 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077115C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007711C3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007711DA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077123C | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0077125D | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x007712CE | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x007712EA | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00771460 | `Radio_Screen` | Known | Screen layout |
| 0x00771470 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007714D1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00771554 | `LockediPod_Screen` | Known | Screen layout |
| 0x007715DC | `Lock_Screen` | Known | Screen layout |
| 0x007715EB | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077164E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007716B0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007716CC | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077173E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077175D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007717C5 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007717DF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00771847 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00771864 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007718D0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077193A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00771954 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007719C4 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00771A37 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00771AA8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00771B17 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00771B83 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00771B9E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00771C13 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00771C7A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00771CDC | `Photos_Screen` | Known | Screen layout |
| 0x00771D41 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00771D60 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00771DCB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00771DE9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00771E5B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00771E78 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00771EDE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00771EF9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00771F62 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00771F7F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00771FF6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077201A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00772088 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007720A3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00772160 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077217C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007721EA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00772207 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00772272 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00772292 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00772309 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00772325 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00772395 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007723B4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00772420 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00772434 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007724AD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00772521 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00772591 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007725F8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00772660 | `NoContent_Screen` | Known | Screen layout |
| 0x00772674 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007726D8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077273F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00772759 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007727C7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00772839 | `NoContent_Screen` | Known | Screen layout |
| 0x0077284D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007728B7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00772920 | `No_Photos_Screen` | Known | Screen layout |
| 0x00772934 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077299A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00772A08 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00772A75 | `NoContent_Screen` | Known | Screen layout |
| 0x00772A89 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00772AF1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00772B5B | `NoContent_Screen` | Known | Screen layout |
| 0x00772B6F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00772BD6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00772C40 | `NoContent_Screen` | Known | Screen layout |
| 0x00772C54 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00772CC1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00772D33 | `NoContent_Screen` | Known | Screen layout |
| 0x00772D47 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00772DAF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00772E18 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00772E33 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00772E99 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00772EB5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00772F94 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00772FAD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077300E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00773022 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077307C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00773098 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007730FF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00773116 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00773178 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00773199 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0077320A | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00773226 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0077339C | `Radio_Screen` | Known | Screen layout |
| 0x007733AC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077340D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00773490 | `LockediPod_Screen` | Known | Screen layout |
| 0x00773518 | `Lock_Screen` | Known | Screen layout |
| 0x00773527 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077358A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007735EC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00773608 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077367A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00773699 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00773701 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077371B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00773783 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007737A0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077380C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00773876 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00773890 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00773900 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00773973 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007739E4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00773A53 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00773ABF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00773ADA | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00773B4F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00773BB6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00773C18 | `Photos_Screen` | Known | Screen layout |
| 0x00773C7D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00773C9C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00773D07 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00773D25 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00773D97 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00773DB4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00773E1A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00773E35 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00773E9E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00773EBB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00773F32 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00773F56 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00773FC4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00773FDF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077409C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007740B8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00774126 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00774143 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007741AE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007741CE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00774245 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00774261 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007742D1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007742F0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077435C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00774370 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007743E9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077445D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007744CD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00774534 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077459C | `NoContent_Screen` | Known | Screen layout |
| 0x007745B0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00774614 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077467B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00774695 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00774703 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00774775 | `NoContent_Screen` | Known | Screen layout |
| 0x00774789 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007747F3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077485C | `No_Photos_Screen` | Known | Screen layout |
| 0x00774870 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007748D6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00774944 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007749B1 | `NoContent_Screen` | Known | Screen layout |
| 0x007749C5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00774A2D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00774A97 | `NoContent_Screen` | Known | Screen layout |
| 0x00774AAB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00774B12 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00774B7C | `NoContent_Screen` | Known | Screen layout |
| 0x00774B90 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00774BFD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00774C6F | `NoContent_Screen` | Known | Screen layout |
| 0x00774C83 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00774CEB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00774D54 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00774D6F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00774DD5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00774DF1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00774ED0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00774EE9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00774F4A | `FirstBoot_Screen` | Known | Screen layout |
| 0x00774F5E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00774FB8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00774FD4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077503B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00775052 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007750B4 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x007750D5 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00775146 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00775162 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x007752D8 | `Radio_Screen` | Known | Screen layout |
| 0x007752E8 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00775349 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007753CC | `LockediPod_Screen` | Known | Screen layout |
| 0x00775454 | `Lock_Screen` | Known | Screen layout |
| 0x00775463 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007754C6 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00775528 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00775544 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007755B6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007755D5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077563D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00775657 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007756BF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007756DC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00775748 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007757B2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007757CC | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077583C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007758AF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00775920 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077598F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007759FB | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00775A16 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00775A8B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00775AF2 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00775B54 | `Photos_Screen` | Known | Screen layout |
| 0x00775BB9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00775BD8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00775C43 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00775C61 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00775CD3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00775CF0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00775D56 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00775D71 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00775DDA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00775DF7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00775E6E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00775E92 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00775F00 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00775F1B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00775FD8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00775FF4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00776062 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077607F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007760EA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077610A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00776181 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077619D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077620D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077622C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00776298 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007762AC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00776325 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00776399 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00776409 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00776470 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007764D8 | `NoContent_Screen` | Known | Screen layout |
| 0x007764EC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00776550 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007765B7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007765D1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077663F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007766B1 | `NoContent_Screen` | Known | Screen layout |
| 0x007766C5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077672F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00776798 | `No_Photos_Screen` | Known | Screen layout |
| 0x007767AC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00776812 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00776880 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007768ED | `NoContent_Screen` | Known | Screen layout |
| 0x00776901 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00776969 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007769D3 | `NoContent_Screen` | Known | Screen layout |
| 0x007769E7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00776A4E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00776AB8 | `NoContent_Screen` | Known | Screen layout |
| 0x00776ACC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00776B39 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00776BAB | `NoContent_Screen` | Known | Screen layout |
| 0x00776BBF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00776C27 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00776C90 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00776CAB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00776D11 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00776D2D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00776E0C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00776E25 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00776E86 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00776E9A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00776EF4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00776F10 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00776F77 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00776F8E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00776FF0 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00777011 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00777082 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0077709E | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00777214 | `Radio_Screen` | Known | Screen layout |
| 0x00777224 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00777285 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00777308 | `LockediPod_Screen` | Known | Screen layout |
| 0x00777390 | `Lock_Screen` | Known | Screen layout |
| 0x0077739F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00777402 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00777464 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00777480 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007774F2 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00777511 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00777579 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00777593 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007775FB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00777618 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00777684 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007776EE | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00777708 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00777778 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007777EB | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077785C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007778CB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00777937 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00777952 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007779C7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00777A2E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00777A90 | `Photos_Screen` | Known | Screen layout |
| 0x00777AF5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00777B14 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00777B7F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00777B9D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00777C0F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00777C2C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00777C92 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00777CAD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00777D16 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00777D33 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00777DAA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00777DCE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00777E3C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00777E57 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00777F14 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00777F30 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00777F9E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00777FBB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00778026 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00778046 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007780BD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007780D9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00778149 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00778168 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007781D4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007781E8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00778261 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007782D5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00778345 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007783AC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00778414 | `NoContent_Screen` | Known | Screen layout |
| 0x00778428 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077848C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007784F3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077850D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077857B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007785ED | `NoContent_Screen` | Known | Screen layout |
| 0x00778601 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077866B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007786D4 | `No_Photos_Screen` | Known | Screen layout |
| 0x007786E8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077874E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007787BC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00778829 | `NoContent_Screen` | Known | Screen layout |
| 0x0077883D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007788A5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077890F | `NoContent_Screen` | Known | Screen layout |
| 0x00778923 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077898A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007789F4 | `NoContent_Screen` | Known | Screen layout |
| 0x00778A08 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00778A75 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00778AE7 | `NoContent_Screen` | Known | Screen layout |
| 0x00778AFB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00778B63 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00778BCC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00778BE7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00778C4D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00778C69 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00778D48 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00778D61 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00778DC2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00778DD6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00778E30 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00778E4C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00778EB3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00778ECA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00778F2C | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00778F4D | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00778FBE | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00778FDA | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00779150 | `Radio_Screen` | Known | Screen layout |
| 0x00779160 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007791C1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00779244 | `LockediPod_Screen` | Known | Screen layout |
| 0x007792CC | `Lock_Screen` | Known | Screen layout |
| 0x007792DB | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077933E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007793A0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007793BC | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077942E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077944D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007794B5 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007794CF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00779537 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00779554 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007795C0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077962A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00779644 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007796B4 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00779727 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00779798 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00779807 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00779873 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077988E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00779903 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077996A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007799CC | `Photos_Screen` | Known | Screen layout |
| 0x00779A31 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00779A50 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00779ABB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00779AD9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00779B4B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00779B68 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00779BCE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00779BE9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00779C52 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00779C6F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00779CE6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00779D0A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00779D78 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00779D93 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00779E50 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00779E6C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00779EDA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00779EF7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00779F62 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00779F82 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00779FF9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077A015 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077A085 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077A0A4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077A110 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077A124 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077A19D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077A211 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077A281 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077A2E8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077A350 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A364 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077A3C8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077A42F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077A449 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077A4B7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077A529 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A53D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077A5A7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077A610 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077A624 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077A68A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077A6F8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077A765 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A779 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077A7E1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077A84B | `NoContent_Screen` | Known | Screen layout |
| 0x0077A85F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077A8C6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077A930 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A944 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077A9B1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077AA23 | `NoContent_Screen` | Known | Screen layout |
| 0x0077AA37 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077AA9F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077AB08 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077AB23 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077AB89 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077ABA5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077AC84 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077AC9D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077ACFE | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077AD12 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077AD6C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077AD88 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077ADEF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077AE06 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077AE68 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0077AE89 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0077AEFA | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0077AF16 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0077B08C | `Radio_Screen` | Known | Screen layout |
| 0x0077B09C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077B0FD | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077B180 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077B208 | `Lock_Screen` | Known | Screen layout |
| 0x0077B217 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077B27A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077B2DC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077B2F8 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077B36A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077B389 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077B3F1 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077B40B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077B473 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077B490 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077B4FC | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077B566 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077B580 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077B5F0 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077B663 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077B6D4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077B743 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077B7AF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077B7CA | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077B83F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077B8A6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077B908 | `Photos_Screen` | Known | Screen layout |
| 0x0077B96D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077B98C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077B9F7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077BA15 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077BA87 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077BAA4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077BB0A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077BB25 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077BB8E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077BBAB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077BC22 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077BC46 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077BCB4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077BCCF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077BD8C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077BDA8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077BE16 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077BE33 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077BE9E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077BEBE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077BF35 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077BF51 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077BFC1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077BFE0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077C04C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077C060 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077C0D9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077C14D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077C1BD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077C224 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077C28C | `NoContent_Screen` | Known | Screen layout |
| 0x0077C2A0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077C304 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077C36B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077C385 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077C3F3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077C465 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C479 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077C4E3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077C54C | `No_Photos_Screen` | Known | Screen layout |
| 0x0077C560 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077C5C6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077C634 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077C6A1 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C6B5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077C71D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077C787 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C79B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077C802 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077C86C | `NoContent_Screen` | Known | Screen layout |
| 0x0077C880 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077C8ED | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077C95F | `NoContent_Screen` | Known | Screen layout |
| 0x0077C973 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077C9DB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077CA44 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077CA5F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077CAC5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077CAE1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077CBC0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077CBD9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077CC3A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077CC4E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077CCA8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077CCC4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077CD2B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077CD42 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077CDA4 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0077CDC5 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0077CE36 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0077CE52 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0077CFC8 | `Radio_Screen` | Known | Screen layout |
| 0x0077CFD8 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077D039 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077D0BC | `LockediPod_Screen` | Known | Screen layout |
| 0x0077D144 | `Lock_Screen` | Known | Screen layout |
| 0x0077D153 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077D1B6 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077D218 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077D234 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077D2A6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077D2C5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077D32D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077D347 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077D3AF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077D3CC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077D438 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077D4A2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077D4BC | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077D52C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077D59F | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077D610 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077D67F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077D6EB | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077D706 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077D77B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077D7E2 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077D844 | `Photos_Screen` | Known | Screen layout |
| 0x0077D8A9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077D8C8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077D933 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077D951 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077D9C3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077D9E0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077DA46 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077DA61 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077DACA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077DAE7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077DB5E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077DB82 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077DBF0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077DC0B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077DCC8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077DCE4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077DD52 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077DD6F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077DDDA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077DDFA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077DE71 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077DE8D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077DEFD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077DF1C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077DF88 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077DF9C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077E015 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077E089 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077E0F9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077E160 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077E1C8 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E1DC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077E240 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077E2A7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077E2C1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077E32F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077E3A1 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E3B5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077E41F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077E488 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077E49C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077E502 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077E570 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077E5DD | `NoContent_Screen` | Known | Screen layout |
| 0x0077E5F1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077E659 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077E6C3 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E6D7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077E73E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077E7A8 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E7BC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077E829 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077E89B | `NoContent_Screen` | Known | Screen layout |
| 0x0077E8AF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077E917 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077E980 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077E99B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077EA01 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077EA1D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077EAFC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077EB15 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077EB76 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077EB8A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077EBE4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077EC00 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077EC67 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077EC7E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077ECE0 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0077ED01 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0077ED72 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0077ED8E | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0077EF04 | `Radio_Screen` | Known | Screen layout |
| 0x0077EF14 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077EF75 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077EFF8 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077F080 | `Lock_Screen` | Known | Screen layout |
| 0x0077F08F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077F0F2 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077F154 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077F170 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077F1E2 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077F201 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077F269 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077F283 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077F2EB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077F308 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077F374 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077F3DE | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077F3F8 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077F468 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077F4DB | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077F54C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077F5BB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077F627 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077F642 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077F6B7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077F71E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077F780 | `Photos_Screen` | Known | Screen layout |
| 0x0077F7E5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077F804 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077F86F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077F88D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077F8FF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077F91C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077F982 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077F99D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077FA06 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077FA23 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077FA9A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077FABE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077FB2C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077FB47 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077FC04 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077FC20 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077FC8E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077FCAB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077FD16 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077FD36 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077FDAD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077FDC9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077FE39 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077FE58 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077FEC4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077FED8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077FF51 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077FFC5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00780035 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0078009C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00780104 | `NoContent_Screen` | Known | Screen layout |
| 0x00780118 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078017C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007801E3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007801FD | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078026B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007802DD | `NoContent_Screen` | Known | Screen layout |
| 0x007802F1 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078035B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007803C4 | `No_Photos_Screen` | Known | Screen layout |
| 0x007803D8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078043E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007804AC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00780519 | `NoContent_Screen` | Known | Screen layout |
| 0x0078052D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00780595 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007805FF | `NoContent_Screen` | Known | Screen layout |
| 0x00780613 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078067A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007806E4 | `NoContent_Screen` | Known | Screen layout |
| 0x007806F8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00780765 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007807D7 | `NoContent_Screen` | Known | Screen layout |
| 0x007807EB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00780853 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007808BC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007808D7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078093D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00780959 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00780A38 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00780A51 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00780AB2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00780AC6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00780B20 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00780B3C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00780BA3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00780BBA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00780C1C | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00780C3D | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00780CAE | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00780CCA | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00780E40 | `Radio_Screen` | Known | Screen layout |
| 0x00780E50 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00780EB1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00780F34 | `LockediPod_Screen` | Known | Screen layout |
| 0x00780FBC | `Lock_Screen` | Known | Screen layout |
| 0x00780FCB | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078102E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00781090 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007810AC | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078111E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078113D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007811A5 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007811BF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00781227 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00781244 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007812B0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078131A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00781334 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007813A4 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00781417 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00781488 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007814F7 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00781563 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078157E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007815F3 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078165A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007816BC | `Photos_Screen` | Known | Screen layout |
| 0x00781721 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00781740 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007817AB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007817C9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078183B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00781858 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007818BE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007818D9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00781942 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078195F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007819D6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007819FA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00781A68 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00781A83 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00781B40 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00781B5C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00781BCA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00781BE7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00781C52 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00781C72 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00781CE9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00781D05 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00781D75 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00781D94 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00781E00 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00781E14 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00781E8D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00781F01 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00781F71 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00781FD8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00782040 | `NoContent_Screen` | Known | Screen layout |
| 0x00782054 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007820B8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078211F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00782139 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007821A7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00782219 | `NoContent_Screen` | Known | Screen layout |
| 0x0078222D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00782297 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00782300 | `No_Photos_Screen` | Known | Screen layout |
| 0x00782314 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078237A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007823E8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00782455 | `NoContent_Screen` | Known | Screen layout |
| 0x00782469 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007824D1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078253B | `NoContent_Screen` | Known | Screen layout |
| 0x0078254F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007825B6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00782620 | `NoContent_Screen` | Known | Screen layout |
| 0x00782634 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007826A1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00782713 | `NoContent_Screen` | Known | Screen layout |
| 0x00782727 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078278F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007827F8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00782813 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00782879 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00782895 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00782974 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078298D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007829EE | `FirstBoot_Screen` | Known | Screen layout |
| 0x00782A02 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00782A5C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00782A78 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00782ADF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00782AF6 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00782B58 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00782B79 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00782BEA | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00782C06 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00782D7C | `Radio_Screen` | Known | Screen layout |
| 0x00782D8C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00782DED | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00782E70 | `LockediPod_Screen` | Known | Screen layout |
| 0x00782EF8 | `Lock_Screen` | Known | Screen layout |
| 0x00782F07 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00782F6A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00782FCC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00782FE8 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078305A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00783079 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007830E1 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007830FB | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00783163 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00783180 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007831EC | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00783256 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00783270 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007832E0 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00783353 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007833C4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00783433 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078349F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007834BA | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078352F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00783596 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007835F8 | `Photos_Screen` | Known | Screen layout |
| 0x0078365D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078367C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007836E7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00783705 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00783777 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00783794 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007837FA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00783815 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078387E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078389B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00783912 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00783936 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007839A4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007839BF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00783A7C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00783A98 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00783B06 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00783B23 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00783B8E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00783BAE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00783C25 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00783C41 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00783CB1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00783CD0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00783D3C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00783D50 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00783DC9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00783E3D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00783EAD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00783F14 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00783F7C | `NoContent_Screen` | Known | Screen layout |
| 0x00783F90 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00783FF4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078405B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00784075 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007840E3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00784155 | `NoContent_Screen` | Known | Screen layout |
| 0x00784169 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007841D3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078423C | `No_Photos_Screen` | Known | Screen layout |
| 0x00784250 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007842B6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00784324 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00784391 | `NoContent_Screen` | Known | Screen layout |
| 0x007843A5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078440D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00784477 | `NoContent_Screen` | Known | Screen layout |
| 0x0078448B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007844F2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078455C | `NoContent_Screen` | Known | Screen layout |
| 0x00784570 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007845DD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078464F | `NoContent_Screen` | Known | Screen layout |
| 0x00784663 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007846CB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00784734 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078474F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007847B5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007847D1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007848B0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007848C9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078492A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078493E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00784998 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007849B4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00784A1B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00784A32 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00784A94 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00784AB5 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00784B26 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00784B42 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00784CB8 | `Radio_Screen` | Known | Screen layout |
| 0x00784CC8 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00784D29 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00784DAC | `LockediPod_Screen` | Known | Screen layout |
| 0x00784E34 | `Lock_Screen` | Known | Screen layout |
| 0x00784E43 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00784EA6 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00784F08 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00784F24 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00784F96 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00784FB5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078501D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00785037 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078509F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007850BC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00785128 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00785192 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007851AC | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078521C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078528F | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00785300 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078536F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007853DB | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007853F6 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078546B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007854D2 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00785534 | `Photos_Screen` | Known | Screen layout |
| 0x00785599 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007855B8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00785623 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00785641 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007856B3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007856D0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00785736 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00785751 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007857BA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007857D7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078584E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00785872 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007858E0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007858FB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007859B8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007859D4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00785A42 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00785A5F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00785ACA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00785AEA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00785B61 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00785B7D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00785BED | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00785C0C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00785C78 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00785C8C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00785D05 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00785D79 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00785DE9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00785E50 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00785EB8 | `NoContent_Screen` | Known | Screen layout |
| 0x00785ECC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00785F30 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00785F97 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00785FB1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078601F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00786091 | `NoContent_Screen` | Known | Screen layout |
| 0x007860A5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078610F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00786178 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078618C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007861F2 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00786260 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007862CD | `NoContent_Screen` | Known | Screen layout |
| 0x007862E1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00786349 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007863B3 | `NoContent_Screen` | Known | Screen layout |
| 0x007863C7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078642E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00786498 | `NoContent_Screen` | Known | Screen layout |
| 0x007864AC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00786519 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078658B | `NoContent_Screen` | Known | Screen layout |
| 0x0078659F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00786607 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00786670 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078668B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007866F1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078670D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007867EC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00786805 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00786866 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078687A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007868D4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007868F0 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00786957 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078696E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007869D0 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x007869F1 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00786A62 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00786A7E | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00786BF4 | `Radio_Screen` | Known | Screen layout |
| 0x00786C04 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00786C65 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00786CE8 | `LockediPod_Screen` | Known | Screen layout |
| 0x00786D70 | `Lock_Screen` | Known | Screen layout |
| 0x00786D7F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00786DE2 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00786E44 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00786E60 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00786ED2 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00786EF1 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00786F59 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00786F73 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00786FDB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00786FF8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00787064 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007870CE | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007870E8 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00787158 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007871CB | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078723C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007872AB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00787317 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00787332 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007873A7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078740E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00787470 | `Photos_Screen` | Known | Screen layout |
| 0x007874D5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007874F4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078755F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078757D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007875EF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078760C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00787672 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078768D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007876F6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00787713 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078778A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007877AE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078781C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00787837 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007878F4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00787910 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078797E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078799B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00787A06 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00787A26 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00787A9D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00787AB9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00787B29 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00787B48 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00787BB4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00787BC8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00787C41 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00787CB5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00787D25 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00787D8C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00787DF4 | `NoContent_Screen` | Known | Screen layout |
| 0x00787E08 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00787E6C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00787ED3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00787EED | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00787F5B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00787FCD | `NoContent_Screen` | Known | Screen layout |
| 0x00787FE1 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078804B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007880B4 | `No_Photos_Screen` | Known | Screen layout |
| 0x007880C8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078812E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078819C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00788209 | `NoContent_Screen` | Known | Screen layout |
| 0x0078821D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00788285 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007882EF | `NoContent_Screen` | Known | Screen layout |
| 0x00788303 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078836A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007883D4 | `NoContent_Screen` | Known | Screen layout |
| 0x007883E8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00788455 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007884C7 | `NoContent_Screen` | Known | Screen layout |
| 0x007884DB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00788543 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007885AC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007885C7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078862D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00788649 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00788728 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00788741 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007887A2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007887B6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00788810 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078882C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00788893 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007888AA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078890C | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0078892D | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0078899E | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x007889BA | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00788B30 | `Radio_Screen` | Known | Screen layout |
| 0x00788B40 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00788BA1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00788C24 | `LockediPod_Screen` | Known | Screen layout |
| 0x00788CAC | `Lock_Screen` | Known | Screen layout |
| 0x00788CBB | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00788D1E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00788D80 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00788D9C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00788E0E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00788E2D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00788E95 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00788EAF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00788F17 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00788F34 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00788FA0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078900A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00789024 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00789094 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00789107 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00789178 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007891E7 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00789253 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078926E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007892E3 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078934A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007893AC | `Photos_Screen` | Known | Screen layout |
| 0x00789411 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00789430 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078949B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007894B9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078952B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00789548 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007895AE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007895C9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00789632 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078964F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007896C6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007896EA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00789758 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00789773 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00789830 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078984C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007898BA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007898D7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00789942 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00789962 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007899D9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007899F5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00789A65 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00789A84 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00789AF0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00789B04 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00789B7D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00789BF1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00789C61 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00789CC8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00789D30 | `NoContent_Screen` | Known | Screen layout |
| 0x00789D44 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00789DA8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00789E0F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00789E29 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00789E97 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00789F09 | `NoContent_Screen` | Known | Screen layout |
| 0x00789F1D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00789F87 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00789FF0 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078A004 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078A06A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078A0D8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078A145 | `NoContent_Screen` | Known | Screen layout |
| 0x0078A159 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078A1C1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078A22B | `NoContent_Screen` | Known | Screen layout |
| 0x0078A23F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078A2A6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078A310 | `NoContent_Screen` | Known | Screen layout |
| 0x0078A324 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078A391 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078A403 | `NoContent_Screen` | Known | Screen layout |
| 0x0078A417 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078A47F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078A4E8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078A503 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078A569 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078A585 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078A664 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078A67D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078A6DE | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078A6F2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078A74C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078A768 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078A7CF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078A7E6 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078A848 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0078A869 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0078A8DA | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0078A8F6 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0078AA6C | `Radio_Screen` | Known | Screen layout |
| 0x0078AA7C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078AADD | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078AB60 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078ABE8 | `Lock_Screen` | Known | Screen layout |
| 0x0078ABF7 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078AC5A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078ACBC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078ACD8 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078AD4A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078AD69 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078ADD1 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078ADEB | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078AE53 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078AE70 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078AEDC | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078AF46 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078AF60 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078AFD0 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078B043 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078B0B4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078B123 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078B18F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078B1AA | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078B21F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078B286 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078B2E8 | `Photos_Screen` | Known | Screen layout |
| 0x0078B34D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078B36C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078B3D7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078B3F5 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078B467 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078B484 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078B4EA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078B505 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078B56E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078B58B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078B602 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078B626 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078B694 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078B6AF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078B76C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078B788 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078B7F6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078B813 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078B87E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078B89E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078B915 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078B931 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078B9A1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078B9C0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078BA2C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078BA40 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078BAB9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078BB2D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078BB9D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0078BC04 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078BC6C | `NoContent_Screen` | Known | Screen layout |
| 0x0078BC80 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078BCE4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078BD4B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078BD65 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078BDD3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078BE45 | `NoContent_Screen` | Known | Screen layout |
| 0x0078BE59 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078BEC3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078BF2C | `No_Photos_Screen` | Known | Screen layout |
| 0x0078BF40 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078BFA6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078C014 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078C081 | `NoContent_Screen` | Known | Screen layout |
| 0x0078C095 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078C0FD | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078C167 | `NoContent_Screen` | Known | Screen layout |
| 0x0078C17B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078C1E2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078C24C | `NoContent_Screen` | Known | Screen layout |
| 0x0078C260 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078C2CD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078C33F | `NoContent_Screen` | Known | Screen layout |
| 0x0078C353 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078C3BB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078C424 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078C43F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078C4A5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078C4C1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078C5A0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078C5B9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078C61A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078C62E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078C688 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078C6A4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078C70B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078C722 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078C784 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0078C7A5 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0078C816 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0078C832 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0078C9A8 | `Radio_Screen` | Known | Screen layout |
| 0x0078C9B8 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078CA19 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078CA9C | `LockediPod_Screen` | Known | Screen layout |
| 0x0078CB24 | `Lock_Screen` | Known | Screen layout |
| 0x0078CB33 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078CB96 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078CBF8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078CC14 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078CC86 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078CCA5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078CD0D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078CD27 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078CD8F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078CDAC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078CE18 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078CE82 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078CE9C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078CF0C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078CF7F | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078CFF0 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078D05F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078D0CB | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078D0E6 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078D15B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078D1C2 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078D224 | `Photos_Screen` | Known | Screen layout |
| 0x0078D289 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078D2A8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078D313 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078D331 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078D3A3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078D3C0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078D426 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078D441 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078D4AA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078D4C7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078D53E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078D562 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078D5D0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078D5EB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078D6A8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078D6C4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078D732 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078D74F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078D7BA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078D7DA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078D851 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078D86D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078D8DD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078D8FC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078D968 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078D97C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078D9F5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078DA69 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078DAD9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0078DB40 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078DBA8 | `NoContent_Screen` | Known | Screen layout |
| 0x0078DBBC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078DC20 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078DC87 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078DCA1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078DD0F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078DD81 | `NoContent_Screen` | Known | Screen layout |
| 0x0078DD95 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078DDFF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078DE68 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078DE7C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078DEE2 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078DF50 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078DFBD | `NoContent_Screen` | Known | Screen layout |
| 0x0078DFD1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078E039 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078E0A3 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E0B7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078E11E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078E188 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E19C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078E209 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078E27B | `NoContent_Screen` | Known | Screen layout |
| 0x0078E28F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078E2F7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078E360 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078E37B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078E3E1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078E3FD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078E4DC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078E4F5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078E556 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078E56A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078E5C4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078E5E0 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078E647 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078E65E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078E6C0 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0078E6E1 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0078E752 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0078E76E | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0078E8E4 | `Radio_Screen` | Known | Screen layout |
| 0x0078E8F4 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078E955 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078E9D8 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078EA60 | `Lock_Screen` | Known | Screen layout |
| 0x0078EA6F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078EAD2 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078EB34 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078EB50 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078EBC2 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078EBE1 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078EC49 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078EC63 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078ECCB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078ECE8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078ED54 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078EDBE | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078EDD8 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078EE48 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078EEBB | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078EF2C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078EF9B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078F007 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078F022 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078F097 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078F0FE | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078F160 | `Photos_Screen` | Known | Screen layout |
| 0x0078F1C5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078F1E4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078F24F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078F26D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078F2DF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078F2FC | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078F362 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078F37D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078F3E6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078F403 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078F47A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078F49E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078F50C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078F527 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078F5C9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078F5E5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078F653 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078F670 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078F6DB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078F6FB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078F772 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078F78E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078F7FE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078F81D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078F889 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078F89D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078F912 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078F97D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078F9EC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078FA5D | `NoContent_Screen` | Known | Screen layout |
| 0x0078FA71 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078FAE0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078FB53 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078FBC0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078FC29 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078FC99 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078FD09 | `NoContent_Screen` | Known | Screen layout |
| 0x0078FD1D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078FD80 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078FDE3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078FDFF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078FE61 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078FE7D | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078FEE4 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078FEFB | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078FF5D | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0078FF7E | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0078FFEF | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0079000B | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x007900CB | `Radio_Screen` | Known | Screen layout |
| 0x007900DB | `Radio_Screen_Default` | Known | Screen layout |
| 0x0079013C | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007901AA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007901C9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00790237 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007902A0 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007902BF | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00790327 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00790342 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007903E5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00790401 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079046F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079048C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007904F7 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00790517 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079058E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007905AA | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079061A | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00790639 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007906A5 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007906B9 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0079072E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00790799 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00790808 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00790879 | `NoContent_Screen` | Known | Screen layout |
| 0x0079088D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007908FC | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0079096F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007909DC | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00790A45 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00790AB5 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00790B25 | `NoContent_Screen` | Known | Screen layout |
| 0x00790B39 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00790B9C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00790BFF | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00790C1B | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00790C7D | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00790C99 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00790D00 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00790D17 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00790D79 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00790D9A | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00790E0B | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00790E27 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00790EE7 | `Radio_Screen` | Known | Screen layout |
| 0x00790EF7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00790F58 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00790FC6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00790FE5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00791053 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007910BC | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007910DB | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00791143 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079115E | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00791201 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079121D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079128B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007912A8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00791313 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00791333 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007913AA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007913C6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00791436 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00791455 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007914C1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007914D5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0079154A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007915B5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00791624 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00791695 | `NoContent_Screen` | Known | Screen layout |
| 0x007916A9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00791718 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0079178B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007917F8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00791861 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007918D1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00791941 | `NoContent_Screen` | Known | Screen layout |
| 0x00791955 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007919B8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00791A1B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00791A37 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00791A99 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00791AB5 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00791B1C | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00791B33 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00791B95 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00791BB6 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00791C27 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00791C43 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00791D03 | `Radio_Screen` | Known | Screen layout |
| 0x00791D13 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00791D74 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00791DE2 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00791E01 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00791E6F | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00791ED8 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00791EF7 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00791F5F | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00791F7A | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079201D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00792039 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007920A7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007920C4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079212F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079214F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007921C6 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007921E2 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00792252 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00792271 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007922DD | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007922F1 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00792366 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007923D1 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00792440 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007924B1 | `NoContent_Screen` | Known | Screen layout |
| 0x007924C5 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00792534 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007925A7 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00792614 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0079267D | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007926ED | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0079275D | `NoContent_Screen` | Known | Screen layout |
| 0x00792771 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007927D4 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00792837 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00792853 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007928B5 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007928D1 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00792938 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0079294F | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007929B1 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x007929D2 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00792A43 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00792A5F | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00792B1F | `Radio_Screen` | Known | Screen layout |
| 0x00792B2F | `Radio_Screen_Default` | Known | Screen layout |
| 0x00792B90 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00792BFE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00792C1D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00792C8B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00792CF4 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00792D13 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00792D7B | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00792D96 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00792E39 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00792E55 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00792EC3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00792EE0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00792F4B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00792F6B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00792FE2 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00792FFE | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079306E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079308D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007930F9 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079310D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00793182 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007931ED | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0079325C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007932CD | `NoContent_Screen` | Known | Screen layout |
| 0x007932E1 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00793350 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007933C3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00793430 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00793499 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00793509 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00793579 | `NoContent_Screen` | Known | Screen layout |
| 0x0079358D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007935F0 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00793653 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079366F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007936D1 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007936ED | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00793754 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0079376B | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007937CD | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x007937EE | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0079385F | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x0079387B | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0079393B | `Radio_Screen` | Known | Screen layout |
| 0x0079394B | `Radio_Screen_Default` | Known | Screen layout |
| 0x007939AC | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00793A1A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00793A39 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00793AA7 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00793B10 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00793B2F | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00793B97 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00793BB2 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00793C55 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00793C71 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00793CDF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00793CFC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00793D67 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00793D87 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00793DFE | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00793E1A | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00793E8A | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00793EA9 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00793F15 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00793F29 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00793F9E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00794009 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00794078 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007940E9 | `NoContent_Screen` | Known | Screen layout |
| 0x007940FD | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079416C | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007941DF | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0079424C | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007942B5 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00794325 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00794395 | `NoContent_Screen` | Known | Screen layout |
| 0x007943A9 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0079440C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0079446F | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079448B | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007944ED | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00794509 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00794570 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00794587 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007945E9 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0079460A | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x0079467B | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00794697 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00794757 | `Radio_Screen` | Known | Screen layout |
| 0x00794767 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007947C8 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00794836 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00794855 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007948C3 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079492C | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079494B | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007949B3 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007949CE | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00794A71 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00794A8D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00794AFB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00794B18 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00794B83 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00794BA3 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00794C1A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00794C36 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00794CA6 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00794CC5 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00794D31 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00794D45 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00794DBA | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00794E25 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00794E94 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00794F05 | `NoContent_Screen` | Known | Screen layout |
| 0x00794F19 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00794F88 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00794FFB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00795068 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007950D1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00795141 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007951B1 | `NoContent_Screen` | Known | Screen layout |
| 0x007951C5 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00795228 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0079528B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007952A7 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00795309 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00795325 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0079538C | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007953A3 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00795405 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00795426 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00795497 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x007954B3 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00795573 | `Radio_Screen` | Known | Screen layout |
| 0x00795583 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007955E4 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00795652 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00795671 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007956DF | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00795748 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00795767 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007957CF | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007957EA | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079588D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007958A9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00795917 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00795934 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079599F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007959BF | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00795A36 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00795A52 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00795AC2 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00795AE1 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00795B4D | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00795B61 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00795BD6 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00795C41 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00795CB0 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00795D21 | `NoContent_Screen` | Known | Screen layout |
| 0x00795D35 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00795DA4 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00795E17 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00795E84 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00795EED | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00795F5D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00795FCD | `NoContent_Screen` | Known | Screen layout |
| 0x00795FE1 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00796044 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007960A7 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007960C3 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00796125 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00796141 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007961A8 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007961BF | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00796221 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00796242 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x007962B3 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x007962CF | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x0079638F | `Radio_Screen` | Known | Screen layout |
| 0x0079639F | `Radio_Screen_Default` | Known | Screen layout |
| 0x00796400 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0079646E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079648D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007964FB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00796564 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00796583 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007965EB | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00796606 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007966A9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007966C5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00796733 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00796750 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007967BB | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007967DB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00796852 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079686E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007968DE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007968FD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00796969 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079697D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007969F2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00796A5D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00796ACC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00796B3D | `NoContent_Screen` | Known | Screen layout |
| 0x00796B51 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00796BC0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00796C33 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00796CA0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00796D09 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00796D79 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00796DE9 | `NoContent_Screen` | Known | Screen layout |
| 0x00796DFD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00796E60 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00796EC3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00796EDF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00796F41 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00796F5D | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00796FC4 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00796FDB | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0079703D | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x0079705E | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x007970CF | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x007970EB | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x007971AB | `Radio_Screen` | Known | Screen layout |
| 0x007971BB | `Radio_Screen_Default` | Known | Screen layout |
| 0x0079721C | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0079728A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007972A9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00797317 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00797380 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079739F | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00797407 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00797422 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007974C5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007974E1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079754F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079756C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007975D7 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007975F7 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079766E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079768A | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007976FA | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00797719 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00797785 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00797799 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0079780E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00797879 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007978E8 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00797959 | `NoContent_Screen` | Known | Screen layout |
| 0x0079796D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007979DC | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00797A4F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00797ABC | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00797B25 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00797B95 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00797C05 | `NoContent_Screen` | Known | Screen layout |
| 0x00797C19 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00797C7C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00797CDF | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00797CFB | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00797D5D | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00797D79 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00797DE0 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00797DF7 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00797E59 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00797E7A | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00797EEB | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00797F07 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00797FC7 | `Radio_Screen` | Known | Screen layout |
| 0x00797FD7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00798038 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007980A6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007980C5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00798133 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079819C | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007981BB | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00798223 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079823E | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007982E1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007982FD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079836B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00798388 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007983F3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00798413 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079848A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007984A6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00798516 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00798535 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007985A1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007985B5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0079862A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00798695 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00798704 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00798775 | `NoContent_Screen` | Known | Screen layout |
| 0x00798789 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007987F8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0079886B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007988D8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00798941 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007989B1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00798A21 | `NoContent_Screen` | Known | Screen layout |
| 0x00798A35 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00798A98 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00798AFB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00798B17 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00798B79 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00798B95 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00798BFC | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00798C13 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00798C75 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00798C96 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00798D07 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00798D23 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00798DE3 | `Radio_Screen` | Known | Screen layout |
| 0x00798DF3 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00798E54 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00798EC2 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00798EE1 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00798F4F | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00798FB8 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00798FD7 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079903F | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079905A | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007990FD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00799119 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00799187 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007991A4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079920F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079922F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007992A6 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007992C2 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00799332 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00799351 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007993BD | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007993D1 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00799446 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007994B1 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00799520 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00799591 | `NoContent_Screen` | Known | Screen layout |
| 0x007995A5 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00799614 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00799687 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007996F4 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0079975D | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007997CD | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0079983D | `NoContent_Screen` | Known | Screen layout |
| 0x00799851 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007998B4 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00799917 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00799933 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00799995 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007999B1 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00799A18 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00799A2F | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00799A91 | `MediaLists_GeniusMixes_Screen%` | Known | Screen layout |
| 0x00799AB2 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00799B23 | `GeniusMixes_Intro_Screen ` | Known | Screen layout |
| 0x00799B3F | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00799BFF | `Radio_Screen` | Known | Screen layout |
| 0x00799C0F | `Radio_Screen_Default` | Known | Screen layout |
| 0x00799C70 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00799CDE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00799CFD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00799D6B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00799DD4 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00799DF3 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00799E5B | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00799E76 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00799F96 | `MediaLists_AudiobookChapters_Screen)` | Known | Screen layout |
| 0x00799FBD | `MediaLists_AudiobookChapters_Screen_Plane!` | Known | Screen layout |
| 0x0079A419 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079A432 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079A726 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0079A742 | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0079A7B1 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079A7CA | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079AB32 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0079AB4E | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0079ABBD | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079ABD6 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079AEFF | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0079AF1B | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0079AF8A | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079AFA3 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079B1D3 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0079B1EE | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x0079B259 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079B274 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x0079B2E7 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0079B302 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0079B52B | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0079B546 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x0079B5B1 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079B5CC | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x0079B63F | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0079B65A | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0079B88E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079B8AA | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x0079B925 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079B941 | `MediaLists_Albums_Screen_WithTrackCount%` | Known | Screen layout |
| 0x0079B9BA | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0079B9D5 | `MediaLists_Songs_Screen_Plain1` | Known | Screen layout |
| 0x0079BA50 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0079BA6B | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0079BCF9 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079BD16 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079BE5D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079BE79 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x0079BEF4 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079BF0F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079C15D | `MediaLists_PodcastEpisodes_Screen'` | Known | Screen layout |
| 0x0079C182 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0079C4BA | `MediaLists_TVSeasons_Screen#` | Known | Screen layout |
| 0x0079C4D9 | `MediaLists_TVSeasons_Screen_Default%` | Known | Screen layout |
| 0x0079C54E | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x0079C56E | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0079C6F6 | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x0079C716 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0079CB20 | `MediaLists_GeniusPlaylist_Screen(` | Known | Screen layout |
| 0x0079CB44 | `MediaLists_GeniusPlaylist_Screen_Default` | Known | Screen layout |
| 0x0079CBAE | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0079CBCA | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0079CC31 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0079CC48 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0079CD0A | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x0079CD2F | `MediaLists_NestedPlaylists_Screen_Default,` | Known | Screen layout |
| 0x0079CDB1 | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x0079CDD0 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x0079D0AE | `MediaLists_GeniusPlaylist_Screen(` | Known | Screen layout |
| 0x0079D0D2 | `MediaLists_GeniusPlaylist_Screen_Default#` | Known | Screen layout |
| 0x0079D14A | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079D161 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079D1D9 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079D1F0 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079D25E | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0079D27A | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0079D2E9 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079D302 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079D3CC | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x0079D3F1 | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x0079D469 | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x0079D488 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x0079D4ED | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079D739 | `MediaLists_GeniusPlaylist_Screen(` | Known | Screen layout |
| 0x0079D75D | `MediaLists_GeniusPlaylist_Screen_Default"` | Known | Screen layout |
| 0x0079D7D6 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079D848 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079D8B3 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079D8CA | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079D942 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079D959 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079D9C7 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0079D9E3 | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0079DA52 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079DA6B | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079DB62 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x0079DE24 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x0079DF24 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x0079DF90 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079DFFA | `NoContent_Screen` | Known | Screen layout |
| 0x0079E00E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079E078 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079E0EC | `NoContent_Screen` | Known | Screen layout |
| 0x0079E100 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079E16B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0079E1D7 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E1EB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079E252 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079E2BE | `NoContent_Screen` | Known | Screen layout |
| 0x0079E2D2 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079E33F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079E3B3 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E3C7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079E42F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079E49C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079E500 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079E51C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0079E588 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079E5A5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079E612 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079E6D9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079E6F6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079E76D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079E791 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079E848 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079E8B2 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E8C6 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079E930 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079E9A4 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E9B8 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079EA23 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0079EA8F | `NoContent_Screen` | Known | Screen layout |
| 0x0079EAA3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079EB0A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079EB76 | `NoContent_Screen` | Known | Screen layout |
| 0x0079EB8A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079EBF7 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079EC6B | `NoContent_Screen` | Known | Screen layout |
| 0x0079EC7F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079ECE7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079ED54 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079EDB8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079EDD4 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0079EE40 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079EE5D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079EECA | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079EF91 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079EFAE | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079F025 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079F049 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079F100 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079F16A | `NoContent_Screen` | Known | Screen layout |
| 0x0079F17E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079F1E8 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079F25C | `NoContent_Screen` | Known | Screen layout |
| 0x0079F270 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079F2DB | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0079F347 | `NoContent_Screen` | Known | Screen layout |
| 0x0079F35B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079F3C2 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079F42E | `NoContent_Screen` | Known | Screen layout |
| 0x0079F442 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079F4AF | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079F523 | `NoContent_Screen` | Known | Screen layout |
| 0x0079F537 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079F59F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079F60C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079F670 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079F68C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0079F6F8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079F715 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079F782 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079F849 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079F866 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079F8DD | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079F901 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079F9B8 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079FA22 | `NoContent_Screen` | Known | Screen layout |
| 0x0079FA36 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079FAA0 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079FB14 | `NoContent_Screen` | Known | Screen layout |
| 0x0079FB28 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079FB93 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0079FBFF | `NoContent_Screen` | Known | Screen layout |
| 0x0079FC13 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079FC7A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079FCE6 | `NoContent_Screen` | Known | Screen layout |
| 0x0079FCFA | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079FD67 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079FDDB | `NoContent_Screen` | Known | Screen layout |
| 0x0079FDEF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079FE57 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079FEC4 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079FF28 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079FF44 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0079FFB0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079FFCD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A003A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A0101 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A011E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A0195 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A01B9 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A0270 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007A02DA | `NoContent_Screen` | Known | Screen layout |
| 0x007A02EE | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007A0358 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A03CC | `NoContent_Screen` | Known | Screen layout |
| 0x007A03E0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A044B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007A04B7 | `NoContent_Screen` | Known | Screen layout |
| 0x007A04CB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A0532 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007A059E | `NoContent_Screen` | Known | Screen layout |
| 0x007A05B2 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A061F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A0693 | `NoContent_Screen` | Known | Screen layout |
| 0x007A06A7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A070F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A077C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A07E0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A07FC | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007A0868 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A0885 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A08F2 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A09B9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A09D6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A0A4D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A0A71 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A0B28 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007A0B92 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0BA6 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007A0C10 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A0C84 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0C98 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A0D03 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007A0D6F | `NoContent_Screen` | Known | Screen layout |
| 0x007A0D83 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A0DEA | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007A0E56 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0E6A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A0ED7 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A0F4B | `NoContent_Screen` | Known | Screen layout |
| 0x007A0F5F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A0FC7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A1034 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A1098 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A10B4 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007A1120 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A113D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A11AA | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A1271 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A128E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A1305 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A1329 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A13E0 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007A144A | `NoContent_Screen` | Known | Screen layout |
| 0x007A145E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007A14C8 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A153C | `NoContent_Screen` | Known | Screen layout |
| 0x007A1550 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A15BB | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007A1627 | `NoContent_Screen` | Known | Screen layout |
| 0x007A163B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A16A2 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007A170E | `NoContent_Screen` | Known | Screen layout |
| 0x007A1722 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A178F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A1803 | `NoContent_Screen` | Known | Screen layout |
| 0x007A1817 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A187F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A18EC | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A1950 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A196C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007A19D8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A19F5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A1A62 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A1B29 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A1B46 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A1BBD | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A1BE1 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A1C98 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007A1D02 | `NoContent_Screen` | Known | Screen layout |
| 0x007A1D16 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007A1D80 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A1DF4 | `NoContent_Screen` | Known | Screen layout |
| 0x007A1E08 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A1E73 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007A1EDF | `NoContent_Screen` | Known | Screen layout |
| 0x007A1EF3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A1F5A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007A1FC6 | `NoContent_Screen` | Known | Screen layout |
| 0x007A1FDA | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A2047 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A20BB | `NoContent_Screen` | Known | Screen layout |
| 0x007A20CF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A2137 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A21A4 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A2208 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A2224 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007A2290 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A22AD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A231A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A23E1 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A23FE | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A2475 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A2499 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A294C | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A2963 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A29DB | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A29F2 | `Genius_Error_Screen_NoGeniusInfoForTrack"` | Known | Screen layout |
| 0x007A2A69 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A2A82 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A2BF0 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x007A2EB0 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A2F1B | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A2F32 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A2FAA | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A2FC1 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A302F | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007A304B | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x007A30BA | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A30D3 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A319D | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x007A31C2 | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x007A323A | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x007A3259 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x007A37DF | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A3851 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A38BC | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A3921 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A398B | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A39F5 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A3A65 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A3ADC | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A3B4E | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A3B65 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A3BDD | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A3BF4 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A3C66 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A3CCD | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A3CE6 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A3D4F | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A3DBA | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A3E24 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A3E8B | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A3EFA | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A3F68 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A3FCD | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A4035 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A40A0 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A410B | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A4172 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A47CB | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A483D | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A48A8 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A490D | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A4977 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A49E1 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A4A51 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A4AC8 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A4B3A | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A4B51 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A4BC9 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A4BE0 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A4C52 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A4CB9 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A4CD2 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A4D3B | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A4DA6 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A4E10 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A4E77 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A4EE6 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A4F54 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A4FB9 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A5021 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A508C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A50F7 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A515E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A57B5 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A5827 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A5892 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A58F7 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A5961 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A59CB | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A5A3B | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A5AB2 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A5B24 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A5B3B | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A5BB3 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A5BCA | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A5C3C | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A5CA3 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A5CBC | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A5D25 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A5D90 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A5DFA | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A5E61 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A5ED0 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A5F3E | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A5FA3 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A600B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A6076 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A60E1 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A6148 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A679D | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A680F | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A687A | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A68DF | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A6949 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A69B3 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A6A23 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A6A9A | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A6B0C | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A6B23 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A6B9B | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A6BB2 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A6C24 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A6C8B | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A6CA4 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A6D0D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A6D78 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A6DE2 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A6E49 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A6EB8 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A6F26 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A6F8B | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A6FF3 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A705E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A70C9 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A7130 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A776D | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A77DF | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A784A | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A78AF | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A7919 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A7983 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A79F3 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A7A6A | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A7ADC | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A7AF3 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A7B6B | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A7B82 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A7BF4 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A7C5B | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A7C74 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A7CDD | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A7D48 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A7DB2 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A7E19 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A7E88 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A7EF6 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A7F5B | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A7FC3 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A802E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A8099 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A8100 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A873D | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A87AF | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A881A | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A887F | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A88E9 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A8953 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A89C3 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A8A3A | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A8AAC | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A8AC3 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A8B3B | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A8B52 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A8BC4 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A8C2B | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A8C44 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A8CAD | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A8D18 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A8D82 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A8DE9 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A8E58 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A8EC6 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A8F2B | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A8F93 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A8FFE | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A9069 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A90D0 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A974A | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A97BC | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A9827 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A988C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A98F6 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A9960 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A99D0 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A9A47 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A9AB9 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A9AD0 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A9B48 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A9B5F | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A9BD1 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A9C38 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A9C51 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A9CBA | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A9D25 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A9D8F | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A9DF6 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A9E65 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A9ED3 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A9F38 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A9FA0 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007AA00B | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007AA076 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007AA0DD | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007AA73C | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007AA7AE | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007AA819 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007AA87E | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007AA8E8 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007AA952 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007AA9C2 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007AAA39 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007AAAAB | `Genius_Error_Screen` | Known | Screen layout |
| 0x007AAAC2 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007AAB3A | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007AAB51 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007AABC3 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007AAC2A | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007AAC43 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007AACAC | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007AAD17 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007AAD81 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007AADE8 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007AAE57 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007AAEC5 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007AAF2A | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007AAF92 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007AAFFD | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007AB068 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007AB0CF | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007AB718 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007AB78A | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007AB7F5 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007AB85A | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007AB8C4 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007AB92E | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007AB99E | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007ABA15 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007ABA87 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007ABA9E | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007ABB16 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007ABB2D | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007ABB9F | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007ABC06 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007ABC1F | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007ABC88 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007ABCF3 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007ABD5D | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007ABDC4 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007ABE33 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007ABEA1 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007ABF06 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007ABF6E | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007ABFD9 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007AC044 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007AC0AB | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007AC6F4 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007AC766 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007AC7D1 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007AC836 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007AC8A0 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007AC90A | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007AC97A | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007AC9F1 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007ACA63 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007ACA7A | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007ACAF2 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007ACB09 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007ACB7B | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007ACBE2 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007ACBFB | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007ACC64 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007ACCCF | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007ACD39 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007ACDA0 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007ACE0F | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007ACE7D | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007ACEE2 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007ACF4A | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007ACFB5 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007AD020 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007AD087 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007AD6D1 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007AD743 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007AD7AE | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007AD813 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007AD87D | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007AD8E7 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007AD957 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007AD9CE | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007ADA40 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007ADA57 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007ADACF | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007ADAE6 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007ADB58 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007ADBBF | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007ADBD8 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007ADC41 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007ADCAC | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007ADD16 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007ADD7D | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007ADDEC | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007ADE5A | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007ADEBF | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007ADF27 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007ADF92 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007ADFFD | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007AE064 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007AE6D3 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007AE745 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007AE7B0 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007AE815 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007AE87F | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007AE8E9 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007AE959 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007AE9D0 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007AEA42 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007AEA59 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007AEAD1 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007AEAE8 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007AEB5A | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007AEBC1 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007AEBDA | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007AEC43 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007AECAE | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007AED18 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007AED7F | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007AEDEE | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007AEE5C | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007AEEC1 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007AEF29 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007AEF94 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007AEFFF | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007AF066 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007AF6E3 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007AF755 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007AF7C0 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007AF825 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007AF88F | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007AF8F9 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007AF969 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007AF9E0 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007AFA52 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007AFA69 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007AFAE1 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007AFAF8 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007AFB6A | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007AFBD1 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007AFBEA | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007AFC53 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007AFCBE | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007AFD28 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007AFD8F | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007AFDFE | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007AFE6C | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007AFED1 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007AFF39 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007AFFA4 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007B000F | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007B0076 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007B06D3 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007B0745 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007B07B0 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007B0815 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007B087F | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007B08E9 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007B0959 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007B09D0 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007B0A42 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007B0A59 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007B0AD1 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007B0AE8 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007B0B5A | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007B0BC1 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007B0BDA | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007B0C43 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007B0CAE | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007B0D18 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007B0D7F | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007B0DEE | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007B0E5C | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007B0EC1 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007B0F29 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007B0F94 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007B0FFF | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007B1066 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007B16B7 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007B1729 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007B1794 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007B17F9 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007B1863 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007B18CD | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007B193D | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007B19B4 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007B1A26 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007B1A3D | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007B1AB5 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007B1ACC | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007B1B3E | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007B1BA5 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007B1BBE | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007B1C27 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007B1C92 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007B1CFC | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007B1D63 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007B1DD2 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007B1E40 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007B1EA5 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007B1F0D | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007B1F78 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007B1FE3 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007B204A | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007B2689 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007B26FB | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007B2766 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007B27CB | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007B2835 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007B289F | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007B290F | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007B2986 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007B29F8 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007B2A0F | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007B2A87 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007B2A9E | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007B2B10 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007B2B77 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007B2B90 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007B2BF9 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007B2C64 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007B2CCE | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007B2D35 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007B2DA4 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007B2E12 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007B2E77 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007B2EDF | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007B2F4A | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007B2FB5 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007B301C | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007B3652 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007B36C4 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007B372F | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007B3794 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007B37FE | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007B3868 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007B38D8 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007B394F | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007B39C1 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007B39D8 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007B3A50 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007B3A67 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007B3AD9 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007B3B40 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007B3B59 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007B3BC2 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007B3C2D | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007B3C97 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007B3CFE | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007B3D6D | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007B3DDB | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007B3E40 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007B3EA8 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007B3F13 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007B3F7E | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007B3FE5 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007B4636 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007B46A8 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007B4713 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007B4778 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007B47E2 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007B484C | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007B48BC | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007B4933 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007B49A5 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007B49BC | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007B4A34 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007B4A4B | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007B4ABD | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007B4B24 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007B4B3D | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007B4BA6 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007B4C11 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007B4C7B | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007B4CE2 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007B4D51 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007B4DBF | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007B4E24 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007B4E8C | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007B4EF7 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007B4F62 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007B4FC9 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007B55D0 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007B5642 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007B56AD | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007B5712 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007B577C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007B57E6 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007B5856 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007B58CD | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007B593F | `Genius_Error_Screen` | Known | Screen layout |
| 0x007B5956 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007B59CE | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007B59E5 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007B5A57 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007B5ABE | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007B5AD7 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007B5B40 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007B5BAB | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007B5C15 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007B5C7C | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007B5CEB | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007B5D59 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007B5DBE | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007B5E26 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007B5E91 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007B5EFC | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007B5F63 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007B62B6 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007B632D | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007B63AA | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007B641C | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007B648C | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007B6502 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007B6570 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007B65DD | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007B6922 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007B6999 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007B6A16 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007B6A88 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007B6AF8 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007B6B6E | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007B6BDC | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007B6C49 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007B6FB2 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007B7029 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007B70A6 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007B7118 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007B7188 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007B71FE | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007B726C | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007B72D9 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007B7642 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007B76B9 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007B7734 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007B77A4 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007B781A | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007B7888 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007B78F5 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007B7C2E | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007B7CA5 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007B7D20 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007B7D90 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007B7E06 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007B7E74 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007B7EE1 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007B8218 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007B828F | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007B830A | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007B837A | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007B83F0 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007B845E | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007B84CB | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007B87DB | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007B8852 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007B88CD | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007B893D | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007B89B3 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007B8A21 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007B8A8E | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007B9092 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007B90AF | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007B912A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007B9143 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007B91BB | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007B91D4 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007B9249 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007B925F | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007B92D6 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007B92EC | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007B9363 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007B9380 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007B93F8 | `Notes_List_Screen` | Known | Screen layout |
| 0x007B940D | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007B95BE | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007B95DB | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007B9656 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007B966F | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007B96E7 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007B9700 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007B9775 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007B978B | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007B9802 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007B9818 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007B988F | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007B98AC | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007B9924 | `Notes_List_Screen` | Known | Screen layout |
| 0x007B9939 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007B9B1A | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007B9B37 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007B9BB2 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007B9BCB | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007B9C43 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007B9C5C | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007B9CD1 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007B9CE7 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007B9D5E | `Notes_Image_Screen` | Known | Screen layout |
| 0x007B9D74 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007B9DEB | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007B9E08 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007B9E80 | `Notes_List_Screen` | Known | Screen layout |
| 0x007B9E95 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007BA04A | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007BA067 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007BA0E2 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007BA0FB | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007BA173 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007BA18C | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007BA201 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007BA217 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007BA28E | `Notes_Image_Screen` | Known | Screen layout |
| 0x007BA2A4 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007BA31B | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007BA338 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007BA3B0 | `Notes_List_Screen` | Known | Screen layout |
| 0x007BA3C5 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007BA6DD | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007BA783 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007BA806 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007BA8BE | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x007BA940 | `PhotosSettingsSlideshowMusic_Screen+` | Known | Screen layout |
| 0x007BA967 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x007BAA4D | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x007BAC05 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007BAC65 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007BACC2 | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007BACE9 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007BAD89 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007BADE9 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007BAE46 | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007BAE6D | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007BB108 | `Photos_Screen` | Known | Screen layout |
| 0x007BB254 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007BB2B8 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007BB319 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007BB376 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007BB3D3 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007BB441 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007BB49E | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007BB644 | `Photos_Screen` | Known | Screen layout |
| 0x007BB790 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007BB7F4 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007BB855 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007BB8B2 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007BB90F | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007BB97D | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007BB9DA | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007BBB80 | `Photos_Screen` | Known | Screen layout |
| 0x007BBCCC | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007BBD30 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007BBD91 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007BBDEE | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007BBE4B | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007BBEB9 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007BBF16 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007BC0BC | `Photos_Screen` | Known | Screen layout |
| 0x007BC208 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007BC26C | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007BC2CD | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007BC32A | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007BC387 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007BC3F5 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007BC452 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007BC5F8 | `Photos_Screen` | Known | Screen layout |
| 0x007BC744 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007BC7A8 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007BC809 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007BC866 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007BC8C3 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007BC931 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007BC98E | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007BCB34 | `Photos_Screen` | Known | Screen layout |
| 0x007BCC80 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007BCCE4 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007BCD45 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007BCDA2 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007BCDFF | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007BCE6D | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007BCECA | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007BD070 | `Photos_Screen` | Known | Screen layout |
| 0x007BD1BC | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007BD222 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007BD284 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007BD2E6 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007BD37C | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007BD49D | `Photos_Screen` | Known | Screen layout |
| 0x007BD534 | `Photos_Screen` | Known | Screen layout |
| 0x007BD680 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007BD6E6 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007BD748 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007BD7AA | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007BD840 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007BD961 | `Photos_Screen` | Known | Screen layout |
| 0x007BD9F8 | `Photos_Screen` | Known | Screen layout |
| 0x007BDB44 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007BDBAA | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007BDC0C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007BDC6E | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007BDD04 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007BDE25 | `Photos_Screen` | Known | Screen layout |
| 0x007BDEBC | `Photos_Screen` | Known | Screen layout |
| 0x007BE008 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007BE06E | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007BE0D0 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007BE132 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007BE1C8 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007BE2E9 | `Photos_Screen` | Known | Screen layout |
| 0x007BE380 | `Photos_Screen` | Known | Screen layout |
| 0x007BE4CC | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007BE532 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007BE594 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007BE5F6 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007BE68C | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007BE7AD | `Photos_Screen` | Known | Screen layout |
| 0x007BE9CD | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007BEA2F | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007BEA9D | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007BEB03 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BEB6C | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BEBD3 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BEC38 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BEF06 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007BEF68 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007BEFD6 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007BF03C | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BF0A5 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BF10C | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BF171 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BF442 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007BF4A4 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007BF512 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007BF578 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BF5E1 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BF648 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BF6AD | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BF921 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007BF97E | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007BF9E0 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007BFA4E | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007BFAB4 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007BFE12 | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007BFE7C | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007C0222 | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007C028C | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007C0581 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007C05E4 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007C0649 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007C06B1 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007C0714 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C077C | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007C07E5 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C084B | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007C08B0 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C091D | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007C098D | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007C0A03 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007C0A79 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007C0AE9 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C0B5E | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007C0BD5 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007C0C49 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007C0CBB | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007C0D35 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C0DA8 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007C0E1A | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C0E9E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C0EC8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C0F4F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C0FDC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C107B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C1095 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C110D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C1127 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C1191 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C11AE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C1226 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C1250 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C12D7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C1364 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C1403 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C141D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C1495 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C14AF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C1519 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C1536 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C15AE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C15D8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C165F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C16EC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C178B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C17A5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C181D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C1837 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C18A1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C18BE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C1936 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C1960 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C19E7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C1A74 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C1B13 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C1B2D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C1BA5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C1BBF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C1C29 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C1C46 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C1CBE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C1CE8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C1D6F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C1DFC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C1E9B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C1EB5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C1F2D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C1F47 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C1FB1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C1FCE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C2046 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C2070 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C20F7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C2184 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C2223 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C223D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C22B5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C22CF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C2339 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C2356 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C23CE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C23F8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C247F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C250C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C25AB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C25C5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C263D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C2657 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C26C1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C26DE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C2756 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C2780 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C2807 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C2894 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C2933 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C294D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C29C5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C29DF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C2A49 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C2A66 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C2ADE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C2B08 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C2B8F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C2C1C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C2CBB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C2CD5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C2D4D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C2D67 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C2DD1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C2DEE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C2E66 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C2E90 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C2F17 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C2FA4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C3043 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C305D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C30D5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C30EF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C3159 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C3176 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C31EE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C3218 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C329F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C332C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C33CB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C33E5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C345D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C3477 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C34E1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C34FE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C3576 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C35A0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C3627 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C36B4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C3753 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C376D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C37E5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C37FF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C3869 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C3886 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C38FE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C3928 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C39AF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C3A3C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C3ADB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C3AF5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C3B6D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C3B87 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C3BF1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C3C0E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C3C86 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C3CB0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C3D37 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C3DC4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C3E63 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C3E7D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C3EF5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C3F0F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C3F79 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C3F96 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C400E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C4038 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C40BF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C414C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C41EB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C4205 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C427D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C4297 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C4301 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C431E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C4396 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C43C0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C4447 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C44D4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C4573 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C458D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C4605 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C461F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C4689 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C46A6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C471E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C4748 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C47CF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C485C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C48FB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C4915 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C498D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C49A7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C4A11 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C4A2E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C4AA6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C4AD0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C4B57 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C4BE4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C4C83 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C4C9D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C4D15 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C4D2F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C4D99 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C4DB6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C4E2E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007C4E58 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007C4EDF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007C4F6C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007C500B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C5025 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C509D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C50B7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C5121 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007C513E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007C51C5 | `SettingsMenus_ResetMainMenu_Screen9` | Known | Screen layout |
| 0x007C5295 | `SettingsMenus_ResetMusicMenu_Screen:` | Known | Screen layout |
| 0x007C5349 | `SettingsMenus_Main_Screen,` | Known | Screen layout |
| 0x007C53BB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C53D5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007C544D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007C5467 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007C57A2 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007C5808 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007C5865 | `Extras_Screen` | Known | Screen layout |
| 0x007C58B9 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x007C5997 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x007C5A05 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007C5AA3 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x007C5ABC | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x007C5B24 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007C5B96 | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x007C5BAF | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x007C5C12 | `DemoMode_Screen` | Known | Screen layout |
| 0x007C5C25 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x007C5C92 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x007C5CAB | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x007C5D1E | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x007C5D39 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x007C5E49 | `VoiceMemos_Context_PlayDelete_Screen)` | Known | Screen layout |
| 0x007C5E71 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x007C5FCA | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007C6039 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007C6125 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007C61E9 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007C620B | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007C6277 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007C6299 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007C6416 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C6432 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007C64F9 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007C6514 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007C6577 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007C65DA | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007C6671 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C668D | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007C6754 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007C676F | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007C67D2 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007C6835 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007C68CD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C68E9 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007C69B0 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007C69CB | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007C6A2E | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007C6A91 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007C6B0E | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007C6B79 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007C6BE5 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007C6C57 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007C6CC4 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007C6D2F | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x007C6D9B | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007C6E03 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007C6E6F | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007C6EE3 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007C6F51 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x007C6FCA | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x007E3AB8 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x007E3B3D | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x007E3E32 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x00999CF7 | `VoiceMemos_Screen_InsufficientDiskSpace2` | Known | Screen layout |
| 0x0099B57B | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0099B593 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0099B5B1 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x0099B6BD | `NowPlaying_Screen_Music` | Known | Screen layout |
| 0x0099B6E9 | `MainMenu_Music_Screen_NoMusic` | Known | Screen layout |
| 0x0099B707 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0099B725 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0099B826 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x0099B8DA | `DeleteRental_Confirmation_Screen_Unwatched` | Known | Screen layout |
| 0x0099B930 | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x0099B97C | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0099BA7E | `NowPlaying_Screen_ExtraInfoLoadFailed` | Known | Screen layout |
| 0x0099BAD9 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x0099BAF2 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x0099BB10 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x0099BB3F | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x0099BB77 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x0099BFAE | `Video_Settings_TV_Screen_Standard` | Known | Screen layout |
| 0x0099BFE0 | `Search_Main_Screen_WithKeyboard` | Known | Screen layout |
| 0x0099C000 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x0099C045 | `VoiceMemos_Screen_InsufficientDiskSpace` | Known | Screen layout |
| 0x0099C109 | `Video_Settings_TV_Screen_Wide` | Known | Screen layout |
| 0x0099C151 | `CoverFlow_Screen_Backside` | Known | Screen layout |
| 0x0099EC5A | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x0099EE5F | `NowPlaying_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x0099EE84 | `Slideshow_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x0099EF54 | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x0099EF6E | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x0099F001 | `RentalDeleted_Screen_Title` | Known | Screen layout |
| 0x0099F01C | `SingleRentalExpiring_Screen_Title` | Known | Screen layout |
| 0x0099F03E | `MultipleRentalsExpiring_Screen_Title` | Known | Screen layout |
| 0x0099F063 | `DeleteRental_Screen_Title` | Known | Screen layout |
| 0x0099F106 | `VoiceMemos_DeleteAll_ScreenTitle` | Known | Screen layout |
| 0x0099F1A3 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0099F1E6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0099F3D7 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0099F4C0 | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x0099F4D9 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x0099F4ED | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x0099F50A | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0099F529 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x0099F5F5 | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x0099F74B | `CoverFlow_Screen_BacksideForCapture` | Known | Screen layout |
| 0x009A0782 | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x009A079D | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x009A0A94 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x009A0AC8 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x009A0B05 | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x009A0C17 | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x009A0D67 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x009A0D9F | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x009A0DC5 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x009A6C7F | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x009A6CAA | `TVOutSlideshow_Screen_Playing` | Known | Screen layout |
| 0x009A6CC8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x009A6D02 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x009A6D9F | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x009A6E0A | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x009A6E8A | `Extras_Screen_Debug` | Known | Screen layout |
| 0x009A6F94 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x009A6FB4 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x009A74E2 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x009A7540 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x009A755B | `Extras_Screen_Lock` | Known | Screen layout |
| 0x009A756E | `MainMenu_List_ScreenLock` | Known | Screen layout |
| 0x009A7587 | `ExtrasMenu_ScreenLock` | Known | Screen layout |
| 0x009A75FA | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x009A761B | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x009A76EE | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x009A7710 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x009A7817 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x009A7857 | `VoiceMemos_Screen_DeletAllAsk` | Known | Screen layout |
| 0x009A7875 | `NowPlaying_Screen_TVOutAsk` | Known | Screen layout |
| 0x009A79D1 | `NowPlaying_Screen_Initial` | Known | Screen layout |
| 0x009A79EB | `NowPlaying_Screen_Video_TVOut_Initial` | Known | Screen layout |
| 0x009A8753 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x009A87D4 | `RemoteUI_Screen` | Known | Screen layout |
| 0x009A87E4 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x009A87FC | `Video_Settings_TV_Screen` | Known | Screen layout |
| 0x009A8815 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x009A882C | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x009A8850 | `AddressViewer_PartialLoad_Screen` | Known | Screen layout |
| 0x009A8871 | `Alarms_Alarm_Clock_Triggered_Screen` | Known | Screen layout |
| 0x009A8895 | `Alarms_Alarm_Triggered_Screen` | Known | Screen layout |
| 0x009A88B3 | `Unsupported_Screen` | Known | Screen layout |
| 0x009A88C6 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x009A88E4 | `LockediPod_Screen` | Known | Screen layout |
| 0x009A88F6 | `DiskMode_Screen` | Known | Screen layout |
| 0x009A8906 | `DemoMode_Screen` | Known | Screen layout |
| 0x009A8916 | `Notes_Image_Screen` | Known | Screen layout |
| 0x009A8929 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x009A8947 | `NowPlaying_Idle_Screen` | Known | Screen layout |
| 0x009A895E | `Game_Screen` | Known | Screen layout |
| 0x009A896A | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x009A8987 | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x009A89A0 | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x009A89C1 | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x009A89E6 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x009A89F9 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x009A8A16 | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x009A8A37 | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x009A8A5C | `Notes_Loading_Screen` | Known | Screen layout |
| 0x009A8A71 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x009A8A87 | `SettingsMenus_AdjustScrolling_Screen` | Known | Screen layout |
| 0x009A8AAC | `Game_Running_Screen` | Known | Screen layout |
| 0x009A8AC0 | `Stopwatch_Screen` | Known | Screen layout |
| 0x009A8AD1 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x009A8AE8 | `Clock_Screen` | Known | Screen layout |
| 0x009A8AF5 | `TVOutSlideshowAsk_Screen` | Known | Screen layout |
| 0x009A8B0E | `Settings_Legal_Screen` | Known | Screen layout |
| 0x009A8B24 | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x009A8B42 | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x009A8B5E | `ToDo_Item_Screen` | Known | Screen layout |
| 0x009A8B6F | `DemoMode_Main_Screen` | Known | Screen layout |
| 0x009A8B84 | `Search_Main_Screen` | Known | Screen layout |
| 0x009A8B97 | `AddressViewer_Main_Screen` | Known | Screen layout |
| 0x009A8BB1 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x009A8BC6 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x009A8BDC | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x009A8BF6 | `Clock_Region_Screen` | Known | Screen layout |
| 0x009A8C0A | `RentalDeleted_Notification_Screen` | Known | Screen layout |
| 0x009A8C2C | `SingleRentalExpiring_Notification_Screen` | Known | Screen layout |
| 0x009A8C55 | `MultipleRentalsExpiring_Notification_Screen` | Known | Screen layout |
| 0x009A8C81 | `RentalError_Notification_Screen` | Known | Screen layout |
| 0x009A8CA1 | `DeleteRental_Confirmation_Screen` | Known | Screen layout |
| 0x009A8CC2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x009A8CDA | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x009A8CF8 | `Video_Settings_Fit_To_Screen` | Known | Screen layout |
| 0x009A8D15 | `RentalInfo_Screen` | Known | Screen layout |
| 0x009A8D27 | `Radio_Screen` | Known | Screen layout |
| 0x009A8D34 | `GeniusMixes_Intro_Screen` | Known | Screen layout |
| 0x009A8D4D | `Genius_Intro_Screen` | Known | Screen layout |
| 0x009A8D61 | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x009A8D7B | `ChargingMode_LowPower_Screen` | Known | Screen layout |
| 0x009A8D98 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x009A8DB2 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x009A8DCC | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x009A8DE6 | `Genius_Error_Screen` | Known | Screen layout |
| 0x009A8DFA | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x009A8E13 | `Extras_Screen` | Known | Screen layout |
| 0x009A8E21 | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x009A8E3E | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x009A8E60 | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x009A8E79 | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x009A8E97 | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x009A8EB0 | `MediaLists_GeniusMixes_Screen` | Known | Screen layout |
| 0x009A8ECE | `Video_Settings_Screen` | Known | Screen layout |
| 0x009A8EE4 | `SettingsMenus_AudiobookSettings_Screen` | Known | Screen layout |
| 0x009A8F0B | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x009A8F31 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x009A8F47 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x009A8F5F | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x009A8F82 | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x009A8F9F | `MediaLists_Rentals_Screen` | Known | Screen layout |
| 0x009A8FB9 | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x009A8FDD | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x009A8FF6 | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x009A9018 | `ToDo_Instructions_Screen` | Known | Screen layout |
| 0x009A9031 | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x009A904D | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x009A9067 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x009A9088 | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x009A90A4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x009A90BC | `VoiceMemos_Screen` | Known | Screen layout |
| 0x009A90CE | `No_Photos_Screen` | Known | Screen layout |
| 0x009A90DF | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x009A90F9 | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x009A9115 | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x009A9139 | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x009A9159 | `ChargingMode_Progress_Screen` | Known | Screen layout |
| 0x009A9176 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x009A918C | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x009A91A7 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x009A91C3 | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x009A91E5 | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x009A9206 | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x009A9220 | `MediaLists_Genius_Screen` | Known | Screen layout |
| 0x009A9239 | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x009A9253 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x009A9272 | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x009A9293 | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x009A92AB | `NoContent_Screen` | Known | Screen layout |
| 0x009A92BC | `Calendar_Event_Screen` | Known | Screen layout |
| 0x009A92D2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x009A92E3 | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x009A92F9 | `Notes_List_Screen` | Known | Screen layout |
| 0x009A930B | `Debug_TestList_Screen` | Known | Screen layout |
| 0x009A9321 | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x009A9342 | `MediaLists_GeniusPlaylist_Screen` | Known | Screen layout |
| 0x009A9363 | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x009A937D | `TODOS_Menu_Screen` | Known | Screen layout |
| 0x009A938F | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x009A93A5 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x009A93C1 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x009A93D6 | `Games_Menu_Screen` | Known | Screen layout |
| 0x009A93E8 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x009A93FB | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x009A941A | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x009A9439 | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x009A945D | `ContextualMenu_Screen` | Known | Screen layout |
| 0x009A9473 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x009A9489 | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x009A94A7 | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x009A94CA | `TVOutSlideshow_Screen` | Known | Screen layout |
| 0x009A94E0 | `CoverFlow_Screen` | Known | Screen layout |
| 0x009A94F1 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x009A9505 | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x009A9527 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x009A953F | `Stopwatch_SessionSummary_Screen` | Known | Screen layout |
| 0x009A955F | `SettingsMenu_AboutWithAccessory_Screen` | Known | Screen layout |
| 0x009A9586 | `Clock_NorthAmerica_City_Screen` | Known | Screen layout |
| 0x009A95A5 | `Clock_SouthAmerica_City_Screen` | Known | Screen layout |
| 0x009A95C4 | `Clock_Africa_City_Screen` | Known | Screen layout |
| 0x009A95DD | `Clock_Australia_City_Screen` | Known | Screen layout |
| 0x009A95F9 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x009A9610 | `Clock_Pacific_City_Screen` | Known | Screen layout |
| 0x009A962A | `Clock_Atlantic_City_Screen` | Known | Screen layout |
| 0x009A9645 | `Clock_Europe_City_Screen` | Known | Screen layout |
| 0x009A9725 | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x009A9776 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x009A9799 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x009A97C1 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x009A9B77 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x009A9C7A | `RentalInfo_Screen_ExpiringSoon` | Known | Screen layout |
| 0x009A9CD0 | `RentalInfo_Screen_NoAlbumArt_ExpiringSoon` | Known | Screen layout |
| 0x009AA09F | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x009AA0F5 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x009AA246 | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x009AA263 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x009AA637 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x009AA759 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x009AA77B | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x009AA7E8 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x009AA807 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x009AAE8F | `ChargingMode_ScreenLayout_LowPower` | Known | Screen layout |
| 0x009AB82C | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x009AB845 | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x009AB98D | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x009ABA69 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x009ABA87 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x009ABAA7 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x009ABBB2 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x009ABBCE | `Extras_Screen_Games` | Known | Screen layout |
| 0x009ABCD4 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x009ABCF3 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x009ABD0F | `Extras_Screen_Notes` | Known | Screen layout |
| 0x009ABE14 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x009ABEEF | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x009AC0BD | `MainMenu_Music_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009AC0E0 | `MainMenus_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009AC103 | `NoContent_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009AC13D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x009AC15C | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x009AC17D | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x009AC22C | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x009AC249 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x009AC2C8 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x009AC3AC | `MainMenu_Music_Screen_NoCompilations` | Known | Screen layout |
| 0x009AC3D1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x009AC558 | `MainMenu_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x009AC57B | `MainMenu_Videos_Screen_NoMusicVideos` | Known | Screen layout |
| 0x009AC5A0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x009AC5BF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x009AC5DE | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x009AC5FF | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x009AC63D | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x009AC65E | `Extras_Screen_VoiceMemos` | Known | Screen layout |
| 0x009AC6C9 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x009AC6FB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x009AC71A | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x009AC7C7 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x009AC833 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x009AC92C | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x009AC948 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x009AC9CB | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x009AC9E6 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x009ACA07 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x009ACAB6 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x009ACAEA | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x009ACB0B | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x009ACBAE | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x009ACBCF | `MainMenu_Videos_Screen_NoPlaylists` | Known | Screen layout |
| 0x009ACBF2 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x009ACC41 | `MainMenu_Music_Screen_NoArtists` | Known | Screen layout |
| 0x009ACCE8 | `NowPlaying_Screen_Genius` | Known | Screen layout |
| 0x009ACD01 | `Genius_Error_Screen_NoGenius` | Known | Screen layout |
| 0x009ACD1E | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x009ACD3D | `Extras_Screen_NikePlus` | Known | Screen layout |
| 0x009ACE8D | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x009ACEAC | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x009ACECD | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x009AD338 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x009AD3EB | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x009AD465 | `Notes_Contents_Screen_Alt` | Known | Screen layout |
| 0x009AD47F | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x009AD52B | `VoiceMemos_Screen_InsufficientDiskSpace2_Default` | Known | Screen layout |
| 0x009AD5DD | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x009AD682 | `VoiceMemos_Screen_InsufficientDiskSpace_Default` | Known | Screen layout |
| 0x009AD6B2 | `NowPlaying_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x009AD6DF | `Slideshow_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x009AE3B8 | `Radio_Screen_Tuning_Default` | Known | Screen layout |
| 0x009AE419 | `VoiceMemos_Screen_DeletAllAsk_Default` | Known | Screen layout |
| 0x009AE43F | `NowPlaying_Screen_TVOutAsk_Default` | Known | Screen layout |
| 0x009AE462 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x009AE480 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x009AE4AC | `AddressViewer_PartialLoad_Screen_Default` | Known | Screen layout |
| 0x009AE4D5 | `Alarms_Alarm_Clock_Triggered_Screen_Default` | Known | Screen layout |
| 0x009AE501 | `Alarms_Alarm_Triggered_Screen_Default` | Known | Screen layout |
| 0x009AE527 | `Unsupported_Screen_Default` | Known | Screen layout |
| 0x009AE542 | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x009AE568 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x009AE580 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x009AE59B | `Game_Screen_Default` | Known | Screen layout |
| 0x009AE5AF | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x009AE5D5 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x009AE5F6 | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x009AE61F | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x009AE649 | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x009AE676 | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x009AE69F | `Notes_Loading_Screen_Default` | Known | Screen layout |
| 0x009AE6BC | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x009AE6DA | `Clock_Screen_Default` | Known | Screen layout |
| 0x009AE6EF | `TVOutSlideshowAsk_Screen_Default` | Known | Screen layout |
| 0x009AE710 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x009AE72E | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x009AE754 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x009AE778 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x009AE791 | `AddressViewer_Main_Screen_Default` | Known | Screen layout |
| 0x009AE7B3 | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x009AE7D0 | `MainMenus_Main_Screen_Default` | Known | Screen layout |
| 0x009AE7EE | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x009AE80B | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x009AE827 | `RentalDeleted_Notification_Screen_Default` | Known | Screen layout |
| 0x009AE851 | `SingleRentalExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x009AE882 | `MultipleRentalsExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x009AE8B6 | `RentalError_Notification_Screen_Default` | Known | Screen layout |
| 0x009AE8DE | `DeleteRental_Confirmation_Screen_Default` | Known | Screen layout |
| 0x009AE907 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x009AE933 | `RentalInfo_Screen_Default` | Known | Screen layout |
| 0x009AE94D | `Radio_Screen_Default` | Known | Screen layout |
| 0x009AE962 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x009AE983 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x009AE99F | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x009AE9C1 | `Extras_Screen_Default` | Known | Screen layout |
| 0x009AE9D7 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x009AE9FD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x009AEA1E | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x009AEA44 | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x009AEA62 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x009AEA84 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x009AEAB0 | `ToDo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009AEAD1 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009AEAF5 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x009AEB17 | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x009AEB3B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x009AEB5A | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x009AEB73 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x009AEB95 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x009AEBB9 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x009AEBD7 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x009AEBFB | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x009AEC25 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x009AEC4E | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x009AEC70 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x009AEC91 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x009AECB1 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x009AECCF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x009AECE8 | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x009AED06 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x009AED20 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x009AED3E | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x009AED67 | `MediaLists_GeniusPlaylist_Screen_Default` | Known | Screen layout |
| 0x009AED90 | `TODOS_Menu_Screen_Default` | Known | Screen layout |
| 0x009AEDAA | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x009AEDC8 | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x009AEDE5 | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x009AEDFF | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x009AEE1A | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x009AEE39 | `ContextualMenu_Screen_Default` | Known | Screen layout |
| 0x009AEE57 | `Debug_MainMenu_Screen_Default` | Known | Screen layout |
| 0x009AEE75 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x009AEE93 | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x009AEEAC | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x009AEEC8 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x009AEEF2 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x009AEF12 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x009AEF3A | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009AEF61 | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009AEF88 | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x009AEFA9 | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x009AEFCD | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x009AEFEC | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x009AF00E | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x009AF031 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x009AF052 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x009AF0E0 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x009AF110 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x009AF132 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x009AF1A3 | `RentalInfo_Screen_NoAlbumArt_Default` | Known | Screen layout |
| 0x009AF1C8 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x009AF7D1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009AF7FD | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009AF842 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x009AF86A | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x009AF88B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x009AF8AC | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x009AF8D2 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x009AF8EF | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x009AF911 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x009AF935 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x009AF959 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x009AFB29 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x009AFC04 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x009AFC55 | `NowPlaying_Screen_Video_TVOut` | Known | Screen layout |
| 0x009AFDC7 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x009AFDEE | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x009B0327 | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x009B04E4 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x009B06D6 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x009B09A2 | `Lock_Screen_Default_Layout` | Known | Screen layout |
| 0x009B0A38 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x009B0A5F | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x009B0C7B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x009B0D55 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x009B0DBC | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009B0DE6 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009B383B | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x009B3887 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x009B3965 | `MainMenu_List_ScreenLock_x` | Known | Screen layout |
| 0x009B3C4B | `RentalInfo_Screen_ExpiresToday` | Known | Screen layout |
| 0x009B3CA1 | `RentalInfo_Screen_NoAlbumArt_ExpiresToday` | Known | Screen layout |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000908B | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | Known | RTOS |
| 0x002A6DC4 | `  K - RTXC` | Known | RTOS |
| 0x002A7DCC | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | Known | RTOS |
| 0x009988E8 | `Returning from RTXCBug` | Known | RTOS |

---

## 9. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000D2530 | `HostOSTask` | Known | RTOS task thread |
| 0x0012D004 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x001324CC | `USBDeviceTask` | Known | RTOS task thread |
| 0x0013C7BC | `DiskReaderTask` | Known | RTOS task thread |
| 0x0014C960 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0014C974 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0018527C | `GeniusMixesTask` | Known | RTOS task thread |
| 0x001A21E0 | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x001DDB58 | `MeCCAIOTask` | Known | RTOS task thread |
| 0x00210EC4 | `StreamCacheMassStorageManagerTimeOutTask` | Known | RTOS task thread |
| 0x00211040 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x00294F44 | `FirewireTask` | Known | RTOS task thread |
| 0x00294F58 | `TouchwheelTask` | Known | RTOS task thread |
| 0x00294F6C | `AudioOutStateTask` | Known | RTOS task thread |
| 0x00294F98 | `DiskMgrTask` | Known | RTOS task thread |
| 0x00294FA8 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x00294FBC | `MikeyTask` | Known | RTOS task thread |
| 0x00294FCC | `TopPlugTask` | Known | RTOS task thread |
| 0x00294FDC | `HPhoneDetTask` | Known | RTOS task thread |
| 0x00295054 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x0029507C | `AlarmTask` | Known | RTOS task thread |
| 0x0029509B | `"USBAudioTask` | Known | RTOS task thread |
| 0x002A7464 | `Undefined Task` | Known | RTOS task thread |
| 0x003E9540 | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x003ECC0C | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x003F5318 | `MeCCARecordingTask` | Known | RTOS task thread |
| 0x008EA480 | `WaveFileDebugTask` | Known | RTOS task thread |

---

## 10. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0024C878 | `Channel Reserved` | Known | Logging channel |
| 0x0024C88C | `Channel AppBoot` | Known | Logging channel |
| 0x0024C89C | `Channel BufferedSongReading` | Known | Logging channel |
| 0x0024C8B8 | `Channel PrefsWriting` | Known | Logging channel |
| 0x0024C8D0 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x0024C8F0 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x0024C908 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x0024C924 | `Channel TestLogging` | Known | Logging channel |
| 0x0024C938 | `Channel AppFileLoading` | Known | Logging channel |
| 0x0024C950 | `Channel VCardReading` | Known | Logging channel |
| 0x0024C968 | `Channel LongSongScanning` | Known | Logging channel |
| 0x0024C9DC | `Channel VoiceRecording` | Known | Logging channel |
| 0x0024C9F4 | `Channel PhotoImporting` | Known | Logging channel |
| 0x0024CA0C | `Channel Notes` | Known | Logging channel |
| 0x0024CA1C | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x0024CA38 | `Channel DiskMode` | Known | Logging channel |
| 0x0024CA4C | `Channel Firewire` | Known | Logging channel |
| 0x0024CA60 | `Channel USB` | Known | Logging channel |
| 0x0024CA80 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x0024CA98 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00081FDC | `gamedata_RW` | Known | Game system |
| 0x00081FF8 | `gamedata_ShareRW` | Known | Game system |
| 0x0008200C | `games_RO` | Known | Game system |
| 0x0096D0A3 | `11TCGamesMenu` | Known | Game system |
| 0x0096D177 | `12TCGameScreen` | Known | Game system |
| 0x0096DF43 | `27TSilverCntlrTransitionAddonI11TCGamesMenuE` | Known | Game system |
| 0x0096DFF8 | `27TSilverCntlrTransitionAddonI12TCGameScreenE` | Known | Game system |
| 0x00998942 | `iPod_Control/games_RO/` | Known | Game system |
| 0x00998959 | `Resources/Games/games_RO/` | Known | Game system |
| 0x009A4457 | `GamesMenu_StatusBar_String` | Known | Game system |
| 0x009A4BDF | `AboutScreen_Games_String` | Known | Game system |
| 0x009ABBE2 | `MainMenu_List_Games` | Known | Game system |
| 0x009ABBF6 | `ExtrasMenu_Games` | Known | Game system |
| 0x009B39D4 | `MainMenu_List_Games_x` | Known | Game system |

---

## 12. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000924CC | `adrmmp4a` | Known | DRM system |
| 0x00139B10 | `AppleDRMVersion` | Known | DRM system |
| 0x00139BB0 | `AppleDRM` | Known | DRM system |
| 0x0013B168 | `AppleVideoDRM` | Known | DRM system |
| 0x0013E52C | `tx3gdrmsp608aavdmp4aesds@` | Known | DRM system |
| 0x001EB50C | `drmttx3g` | Known | DRM system |
| 0x00998DCB | `DRMLevel` | Known | DRM system |

---

## 13. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00031350 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00031368 | `iPod_Control\iTunes\firsttime` | Known | iTunes database |
| 0x00052910 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x00052938 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00058CD0 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0007DFA4 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x00081F6C | `iPod_Control/iTunes/iTunesDB.p7b` | Known | iTunes database |
| 0x00095400 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0009EA78 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x0009EC60 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x0009F58C | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000A7734 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000A8C44 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000A8D44 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0012561C | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x0021EB60 | `%s/sqlite_` | Known | SQLite database |
| 0x0028607C | `iPod_Control/iTunes/primary.db` | Known | iTunes database |
| 0x00286C44 | `iPod_Control/iTunes/Extras.itdb` | Known | iTunes database |
| 0x002AA8B8 | `sqlite3BtreeInitPage() returns error code %d` | Known | SQLite database |
| 0x002ADB94 | `sqlite_master` | Known | SQLite database |
| 0x002ADBA4 | `sqlite_temp_master` | Known | SQLite database |
| 0x002C4C70 | `sqlite_stat1` | Known | SQLite database |
| 0x002C4C80 | `CREATE TABLE %Q.sqlite_stat1(tbl,idx,stat)` | Known | SQLite database |
| 0x002C4CAC | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x002CF654 | `sqlite_subquery_%p_` | Known | SQLite database |
| 0x003648AC | `sqlite_master` | Known | SQLite database |
| 0x003648BC | `sqlite_temp_master` | Known | SQLite database |
| 0x00364BE0 | `sqlite_` | Known | SQLite database |
| 0x00364C20 | `sqlite_master` | Known | SQLite database |
| 0x00364C30 | `sqlite_temp_master` | Known | SQLite database |
| 0x00364C48 | `sqlite_sequence` | Known | SQLite database |
| 0x00364C58 | `UPDATE "%w".sqlite_sequence set name = %Q WHERE name = %Q` | Known | SQLite database |
| 0x00364D3C | `sqlite_stat1` | Known | SQLite database |
| 0x00364D4C | `SELECT idx, stat FROM %Q.sqlite_stat1` | Known | SQLite database |
| 0x00365A28 | `sqlite_` | Known | SQLite database |
| 0x00365C24 | `sqlite_master` | Known | SQLite database |
| 0x00365C34 | `sqlite_temp_master` | Known | SQLite database |
| 0x00368950 | `sqlite_` | Known | SQLite database |
| 0x00369C3C | `sqlite_autoindex_` | Known | SQLite database |
| 0x00369C50 | `sqlite_master` | Known | SQLite database |
| 0x00369C60 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036B0B8 | `sqlite_master` | Known | SQLite database |
| 0x0036B0C8 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036B0FC | `sqlite_stat1` | Known | SQLite database |
| 0x0036B10C | `DELETE FROM %Q.sqlite_stat1 WHERE idx=%Q` | Known | SQLite database |
| 0x0036B3F4 | `sqlite_master` | Known | SQLite database |
| 0x0036B404 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036B478 | `DELETE FROM %s.sqlite_sequence WHERE name=%Q` | Known | SQLite database |
| 0x0036B4E0 | `sqlite_stat1` | Known | SQLite database |
| 0x0036B4F0 | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x0036B868 | `sqlite_master` | Known | SQLite database |
| 0x0036B878 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036BC90 | `sqlite_master` | Known | SQLite database |
| 0x0036BCA0 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036BCB8 | `CREATE TABLE %Q.sqlite_sequence(name,seq)` | Known | SQLite database |
| 0x0036EF40 | `sqlite_master` | Known | SQLite database |
| 0x0036EF50 | `sqlite_temp_master` | Known | SQLite database |
| 0x00371338 | `sqlite_temp_master` | Known | SQLite database |
| 0x00371350 | `sqlite_master` | Known | SQLite database |
| 0x00372B2C | `sqlite3_extension_init` | Known | SQLite database |
| 0x00373320 | `sqlite_master` | Known | SQLite database |
| 0x00373330 | `sqlite_temp_master` | Known | SQLite database |
| 0x00377710 | `sqlite_attach` | Known | SQLite database |
| 0x00377724 | `sqlite_detach` | Known | SQLite database |
| 0x0037A458 | `sqlite_master` | Known | SQLite database |
| 0x0037A468 | `sqlite_temp_master` | Known | SQLite database |
| 0x0037A4B8 | `sqlite_sequence` | Known | SQLite database |
| 0x0037FD44 | `sqlite_master` | Known | SQLite database |
| 0x0037FD54 | `sqlite_temp_master` | Known | SQLite database |
| 0x003830E8 | `sqlite_master` | Known | SQLite database |
| 0x003830F8 | `sqlite_temp_master` | Known | SQLite database |
| 0x00391294 | `sqlite_attach` | Known | SQLite database |
| 0x003912A4 | `sqlite_detach` | Known | SQLite database |
| 0x003E27FC | `iTunesDB` | Known | iTunes database |
| 0x003E2808 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x008E64CF | `SQLite format 3` | Known | SQLite database |
| 0x008E8B7C | `CREATE TABLE sqlite_master(` | Known | SQLite database |
| 0x008E8BE4 | `CREATE TEMP TABLE sqlite_temp_master(` | Known | SQLite database |
| 0x008E92AC | `illegal return value (%d) from the authorization function - should be SQLITE_OK,` | Known | SQLite database |
| 0x008E9364 | `SELECT 'CREATE TABLE vacuum_db.' || substr(sql,14)   FROM sqlite_master WHERE ty` | Known | SQLite database |
| 0x008E93EC | `SELECT 'CREATE INDEX vacuum_db.' || substr(sql,14)  FROM sqlite_master WHERE sql` | Known | SQLite database |
| 0x008E9454 | `SELECT 'CREATE UNIQUE INDEX vacuum_db.' || substr(sql,21)   FROM sqlite_master W` | Known | SQLite database |
| 0x008E94CC | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x008E957C | `SELECT 'DELETE FROM vacuum_db.' || quote(name) || ';' FROM vacuum_db.sqlite_mast` | Known | SQLite database |
| 0x008E95F0 | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x008E9688 | `INSERT INTO vacuum_db.sqlite_master   SELECT type, name, tbl_name, rootpage, sql` | Known | SQLite database |
| 0x008E9848 | `UPDATE %Q.%s SET sql = CASE WHEN type = 'trigger' THEN sqlite_rename_trigger(sql` | Known | SQLite database |
| 0x008E99BC | `UPDATE sqlite_temp_master SET sql = sqlite_rename_trigger(sql, %Q), tbl_name = %` | Known | SQLite database |
| 0x008E9BF8 | `sqlite3_get_table() called with two or more incompatible queries` | Known | SQLite database |
| 0x009B4492 | `sqlite_rename_table` | Known | SQLite database |
| 0x009B4615 | `sqlite_version` | Known | SQLite database |
| 0x009B46AF | `sqlite_rename_trigger` | Known | SQLite database |
| 0x009B49D3 | `SQLite_iPod_VFS` | Known | SQLite database |

---

## 14. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0005F0C8 | `cI: could not read CE-ATA task file` | Known | Hardware |
| 0x0005F0F0 | `cI: CE-ATA signature missing (%x,%x)` | Known | Hardware |
| 0x0005F148 | `cI: CE-ATA interrupt enable failed` | Known | Hardware |
| 0x00124E80 | `NAND FLASH DRIVE` | Known | Hardware |
| 0x0013A078 | `FireWireGUID` | Known | FireWire |
| 0x0013A088 | `FireWireVersion` | Known | FireWire |
| 0x0013AA5C | `FireWire` | Known | FireWire |
| 0x0035FBE8 | `CE-ATA init failed` | Known | Hardware |
| 0x003600A8 | `ISDIE: CE-ATA interrupt enable failed` | Known | Hardware |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0072CB52 | `Settings_RadioRegionsMenu_Layout"` | Known | FM Radio |
| 0x0072CBDB | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x007E2F60 | `Radio Regions` | Known | FM Radio |
| 0x00833A74 | `Radio-Regionen` | Known | FM Radio |
| 0x0096DB40 | `23TCSettings_RadioRegions` | Known | FM Radio |
| 0x0096EA53 | `27TSilverCntlrTransitionAddonI23TCSettings_RadioRegionsE` | Known | FM Radio |
| 0x009A14B8 | `Settings_RadioRegions_Australia_String` | Known | FM Radio |
| 0x009A14DF | `Settings_RadioRegions_Asia_String` | Known | FM Radio |
| 0x009A2744 | `Settings_RadioRegions_Europe_String` | Known | FM Radio |
| 0x009A3D53 | `Settings_RadioRegions_Japan_String` | Known | FM Radio |
| 0x009A49FC | `Settings_RadioRegions_Americas_String` | Known | FM Radio |
| 0x009A5118 | `SettingsMenu_RadioRegions_String` | Known | FM Radio |
| 0x009A8612 | `RadioRegionsMenuItem` | Known | FM Radio |
| 0x009AC335 | `SettingsMenu_ListItem_RadioRegions` | Known | FM Radio |
| 0x009B05B0 | `SettingsInfo_Template_RadioRegions_Layout` | Known | FM Radio |
| 0x009B05DA | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x009B0C3C | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 16. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00872F18 | `Fotocamera` | Known | Camera |
| 0x0087347C | `Collega fotocamera o inserisci scheda` | Known | Camera |
| 0x008734F4 | `Fotocamera non supportata` | Known | Camera |
| 0x008928EC | `Camera` | Known | Camera |
| 0x00892E6C | `Sluit camera of kaart aan` | Known | Camera |
| 0x00892ED8 | `Camera niet ondersteund` | Known | Camera |
| 0x009A1501 | `Photos_DevName_Camera_String` | Known | Camera |

---

## 17. Pedometer

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x009B3D6B | `NikePlus_Remote_Step_Away` | Known | Pedometer |
| 0x009B3D85 | `NikePlus_Step_Away` | Known | Pedometer |
| 0x009B4650 | `AggStep` | Known | Pedometer |

---

## 18. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0003133C | `iPod_Control` | Filesystem Path |  |
| 0x000313A8 | `iPod_Control\Device` | Filesystem Path |  |
| 0x0003FDDC | `iPod_Control\Device` | Filesystem Path |  |
| 0x00041E68 | `iPod_Control` | Filesystem Path |  |
| 0x000424D4 | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x000528F0 | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path |  |
| 0x00055530 | `iPod_Control\Music\` | Filesystem Path |  |
| 0x00058B50 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x0008BEB4 | `iPod_Control` | Filesystem Path |  |
| 0x0008BEC4 | `Resources/Games` | Filesystem Path |  |
| 0x0008BED4 | `iPod_Control/%s%s%s` | Filesystem Path |  |
| 0x000E9430 | `iPod_Control\Device\dst` | Filesystem Path |  |
| 0x000F455C | `iPod_Control/Device/alarms` | Filesystem Path |  |
| 0x00104A78 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00105FB8 | `iPod_Control/Device` | Filesystem Path |  |
| 0x00105FCC | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00120020 | `iPod_Control/Device/Users` | Filesystem Path |  |
| 0x0014DE04 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x0014E060 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x0015AC60 | `Resources/UI/active.bin` | Filesystem Path |  |
| 0x0015AC78 | `Resources/UI/` | Filesystem Path |  |
| 0x0017EA10 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x0017F93C | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path |  |
| 0x0017F964 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001A5828 | `iPod_Control/Device/PlayCounts` | Filesystem Path |  |
| 0x001BBAA8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBB58 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBCD4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBE6C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BBF14 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC0C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC168 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC20C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC2B0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC354 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC404 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC4A8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC54C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC5FC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC6AC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC75C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC8C8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BC978 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BCA28 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BCACC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BCB7C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BCC70 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BCD14 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BCDC8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BCE84 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BCF34 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BD058 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BD114 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BD1C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BD380 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BD444 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BD4F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BD5B0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BD6EC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BD7B8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BD874 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BD918 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BD9BC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BDA78 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BDB34 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BDBFC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BDCA0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BDD68 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BDE30 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BDEE0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BDFA8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BE070 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BE120 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BE1D0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BE294 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BE344 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BE3F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BE4A4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BE578 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BE64C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BE74C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BE82C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BE934 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BEA20 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003E287A | `iPod_Control/Device` | Filesystem Path |  |
| 0x003E8DE0 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x003EAFD4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003EB426 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003ECD78 | `Resources/Fonts` | Filesystem Path |  |
| 0x003F52E4 | `iPod_Control:Device:voices:` | Filesystem Path |  |
| 0x0099881D | `Resources/Games/` | Filesystem Path |  |
| 0x00998C3B | `iPod_Control/Device` | Filesystem Path |  |
| 0x00998C4F | `iPod_Control/Device/accessories.reg` | Filesystem Path |  |
| 0x00998D42 | `iPod_Control/Tones` | Filesystem Path |  |

---

## 19. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x008ECBC0 | `c:\BWA\N25CFirmwareWin-33\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftu` | Build Path |  |
| 0x008ECC18 | `c:\BWA\N25CFirmwareWin-33\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fts` | Build Path |  |
| 0x008ECC70 | `c:\BWA\N25CFirmwareWin-33\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fto` | Build Path |  |
| 0x008F7B90 | `c:\BWA\N25CFirmwareWin-33\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\` | Build Path |  |
| 0x008F870C | `c:\BWA\N25CFirmwareWin-33\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfd` | Build Path |  |
| 0x008F9908 | `c:\BWA\N25CFirmwareWin-33\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrg` | Build Path |  |
| 0x008F9960 | `c:\BWA\N25CFirmwareWin-33\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrc` | Build Path |  |
| 0x008F99B8 | `c:\BWA\N25CFirmwareWin-33\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfro` | Build Path |  |
| 0x008F9CFC | `c:\BWA\N25CFirmwareWin-33\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1` | Build Path |  |
| 0x009090A4 | `c:\BWA\N25CFirmwareWin-33\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttc` | Build Path |  |
| 0x00909320 | `c:\BWA\N25CFirmwareWin-33\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype` | Build Path |  |
| 0x0090988C | `c:\BWA\N25CFirmwareWin-33\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1` | Build Path |  |

---

## 20. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00089AFC | `Acoustic` | EQ Preset |  |
| 0x00089B08 | `Bass Booster` | EQ Preset |  |
| 0x00089B28 | `Classical` | EQ Preset |  |
| 0x00089B34 | `Dance` | EQ Preset |  |
| 0x00089B44 | `Electronic` | EQ Preset |  |
| 0x00089B58 | `Hip Hop` | EQ Preset |  |
| 0x00089B60 | `Jazz` | EQ Preset |  |
| 0x00089B68 | `Latin` | EQ Preset |  |
| 0x00089B70 | `Loudness` | EQ Preset |  |
| 0x00089B7C | `Lounge` | EQ Preset |  |
| 0x00089B84 | `Piano` | EQ Preset |  |
| 0x00089B98 | `Rock` | EQ Preset |  |
| 0x00089BA0 | `Small Speakers` | EQ Preset |  |
| 0x00089BB0 | `Spoken Word` | EQ Preset |  |
| 0x00089BBC | `Treble Booster` | EQ Preset |  |
| 0x00089C08 | `Vocal Booster` | EQ Preset |  |
| 0x007E3250 | `Acoustic` | EQ Preset |  |
| 0x007E325C | `Bass Booster` | EQ Preset |  |
| 0x007E327C | `Classical` | EQ Preset |  |
| 0x007E3288 | `Dance` | EQ Preset |  |
| 0x007E3298 | `Electronic` | EQ Preset |  |
| 0x007E32AC | `Hip Hop` | EQ Preset |  |
| 0x007E32B4 | `Jazz` | EQ Preset |  |
| 0x007E32BC | `Latin` | EQ Preset |  |
| 0x007E32C4 | `Loudness` | EQ Preset |  |
| 0x007E32D0 | `Lounge` | EQ Preset |  |
| 0x007E32D8 | `Piano` | EQ Preset |  |
| 0x007E32E8 | `Rock` | EQ Preset |  |
| 0x007E32F0 | `Small Speakers` | EQ Preset |  |
| 0x007E3300 | `Spoken Word` | EQ Preset |  |
| 0x007E330C | `Treble Booster` | EQ Preset |  |
| 0x007E332C | `Vocal Booster` | EQ Preset |  |
| 0x00820E30 | `Acoustic` | EQ Preset |  |
| 0x00820E3C | `Bass Booster` | EQ Preset |  |
| 0x00820E5C | `Classical` | EQ Preset |  |
| 0x00820E68 | `Dance` | EQ Preset |  |
| 0x00820E78 | `Electronic` | EQ Preset |  |
| 0x00820E8C | `Hip Hop` | EQ Preset |  |
| 0x00820E94 | `Jazz` | EQ Preset |  |
| 0x00820E9C | `Latin` | EQ Preset |  |
| 0x00820EA4 | `Loudness` | EQ Preset |  |
| 0x00820EB0 | `Lounge` | EQ Preset |  |
| 0x00820EB8 | `Piano` | EQ Preset |  |
| 0x00820EC8 | `Rock` | EQ Preset |  |
| 0x00820ED0 | `Small Speakers` | EQ Preset |  |
| 0x00820EE0 | `Spoken Word` | EQ Preset |  |
| 0x00820EEC | `Treble Booster` | EQ Preset |  |
| 0x00820F0C | `Vocal Booster` | EQ Preset |  |
| 0x0082A37C | `Acoustic` | EQ Preset |  |
| 0x0082A388 | `Bass Booster` | EQ Preset |  |
| 0x0082A3A8 | `Classical` | EQ Preset |  |
| 0x0082A3B4 | `Dance` | EQ Preset |  |
| 0x0082A3C4 | `Electronic` | EQ Preset |  |
| 0x0082A3D8 | `Hip Hop` | EQ Preset |  |
| 0x0082A3E0 | `Jazz` | EQ Preset |  |
| 0x0082A3E8 | `Latin` | EQ Preset |  |
| 0x0082A3F0 | `Loudness` | EQ Preset |  |
| 0x0082A3FC | `Lounge` | EQ Preset |  |
| 0x0082A404 | `Piano` | EQ Preset |  |
| 0x0082A414 | `Rock` | EQ Preset |  |
| 0x0082A41C | `Small Speakers` | EQ Preset |  |
| 0x0082A42C | `Spoken Word` | EQ Preset |  |
| 0x0082A438 | `Treble Booster` | EQ Preset |  |
| 0x0082A458 | `Vocal Booster` | EQ Preset |  |
| 0x00833E1C | `Acoustic` | EQ Preset |  |
| 0x00833E4C | `Dance` | EQ Preset |  |
| 0x00833E5C | `Electronic` | EQ Preset |  |
| 0x00833E78 | `Jazz` | EQ Preset |  |
| 0x00833E80 | `Latin` | EQ Preset |  |
| 0x00833E88 | `Loudness` | EQ Preset |  |
| 0x00833E9C | `Piano` | EQ Preset |  |
| 0x00833EAC | `Rock` | EQ Preset |  |
| 0x0084C030 | `Dance` | EQ Preset |  |
| 0x0084C058 | `Hip Hop` | EQ Preset |  |
| 0x0084C060 | `Jazz` | EQ Preset |  |
| 0x0084C070 | `Loudness` | EQ Preset |  |
| 0x0084C07C | `Lounge` | EQ Preset |  |
| 0x0084C084 | `Piano` | EQ Preset |  |
| 0x0084C094 | `Rock` | EQ Preset |  |
| 0x00855650 | `Jazz` | EQ Preset |  |
| 0x00855658 | `Latin` | EQ Preset |  |
| 0x0085566C | `Lounge` | EQ Preset |  |
| 0x00855674 | `Piano` | EQ Preset |  |
| 0x00855684 | `Rock` | EQ Preset |  |
| 0x0085EBA4 | `Hip Hop` | EQ Preset |  |
| 0x0085EBAC | `Jazz` | EQ Preset |  |
| 0x0085EBC8 | `Lounge` | EQ Preset |  |
| 0x0085EBD0 | `Piano` | EQ Preset |  |
| 0x0085EBE8 | `Rock` | EQ Preset |  |
| 0x00868D74 | `Latin` | EQ Preset |  |
| 0x00868DA0 | `Rock` | EQ Preset |  |
| 0x00872804 | `Dance` | EQ Preset |  |
| 0x00872828 | `Hip Hop` | EQ Preset |  |
| 0x00872830 | `Jazz` | EQ Preset |  |
| 0x00872840 | `Loudness` | EQ Preset |  |
| 0x0087284C | `Lounge` | EQ Preset |  |
| 0x00872854 | `Piano` | EQ Preset |  |
| 0x00872864 | `Rock` | EQ Preset |  |
| 0x0087D758 | `Acoustic` | EQ Preset |  |
| 0x0087D764 | `Bass Booster` | EQ Preset |  |
| 0x0087D784 | `Classical` | EQ Preset |  |
| 0x0087D790 | `Dance` | EQ Preset |  |
| 0x0087D7A0 | `Electronic` | EQ Preset |  |
| 0x0087D7B4 | `Hip Hop` | EQ Preset |  |
| 0x0087D7BC | `Jazz` | EQ Preset |  |
| 0x0087D7C4 | `Latin` | EQ Preset |  |
| 0x0087D7CC | `Loudness` | EQ Preset |  |
| 0x0087D7D8 | `Lounge` | EQ Preset |  |
| 0x0087D7E0 | `Piano` | EQ Preset |  |
| 0x0087D7F0 | `Rock` | EQ Preset |  |
| 0x0087D7F8 | `Small Speakers` | EQ Preset |  |
| 0x0087D808 | `Spoken Word` | EQ Preset |  |
| 0x0087D814 | `Treble Booster` | EQ Preset |  |
| 0x0087D834 | `Vocal Booster` | EQ Preset |  |
| 0x008884E0 | `Acoustic` | EQ Preset |  |
| 0x008884EC | `Bass Booster` | EQ Preset |  |
| 0x0088850C | `Classical` | EQ Preset |  |
| 0x00888518 | `Dance` | EQ Preset |  |
| 0x00888528 | `Electronic` | EQ Preset |  |
| 0x0088853C | `Hip Hop` | EQ Preset |  |
| 0x00888544 | `Jazz` | EQ Preset |  |
| 0x0088854C | `Latin` | EQ Preset |  |
| 0x00888554 | `Loudness` | EQ Preset |  |
| 0x00888560 | `Lounge` | EQ Preset |  |
| 0x00888568 | `Piano` | EQ Preset |  |
| 0x00888578 | `Rock` | EQ Preset |  |
| 0x00888580 | `Small Speakers` | EQ Preset |  |
| 0x00888590 | `Spoken Word` | EQ Preset |  |
| 0x0088859C | `Treble Booster` | EQ Preset |  |
| 0x008885BC | `Vocal Booster` | EQ Preset |  |
| 0x008921D0 | `Dance` | EQ Preset |  |
| 0x00892204 | `Jazz` | EQ Preset |  |
| 0x0089220C | `Latin` | EQ Preset |  |
| 0x00892214 | `Loudness` | EQ Preset |  |
| 0x00892220 | `Lounge` | EQ Preset |  |
| 0x00892228 | `Piano` | EQ Preset |  |
| 0x00892238 | `Rock` | EQ Preset |  |
| 0x0089B6C8 | `Dance` | EQ Preset |  |
| 0x0089B6F4 | `Jazz` | EQ Preset |  |
| 0x0089B704 | `Loudness` | EQ Preset |  |
| 0x0089B710 | `Lounge` | EQ Preset |  |
| 0x0089B718 | `Piano` | EQ Preset |  |
| 0x0089B728 | `Rock` | EQ Preset |  |
| 0x008A4E78 | `Hip Hop` | EQ Preset |  |
| 0x008A4E80 | `Jazz` | EQ Preset |  |
| 0x008A4EA4 | `Lounge` | EQ Preset |  |
| 0x008A4EBC | `Rock` | EQ Preset |  |
| 0x008AEA3C | `Hip Hop` | EQ Preset |  |
| 0x008AEA44 | `Jazz` | EQ Preset |  |
| 0x008AEA60 | `Lounge` | EQ Preset |  |
| 0x008AEA68 | `Piano` | EQ Preset |  |
| 0x008AEA78 | `Rock` | EQ Preset |  |
| 0x008C575C | `Acoustic` | EQ Preset |  |
| 0x008C5768 | `Bass Booster` | EQ Preset |  |
| 0x008C5788 | `Classical` | EQ Preset |  |
| 0x008C5794 | `Dance` | EQ Preset |  |
| 0x008C57A4 | `Electronic` | EQ Preset |  |
| 0x008C57B8 | `Hip Hop` | EQ Preset |  |
| 0x008C57C0 | `Jazz` | EQ Preset |  |
| 0x008C57C8 | `Latin` | EQ Preset |  |
| 0x008C57D0 | `Loudness` | EQ Preset |  |
| 0x008C57DC | `Lounge` | EQ Preset |  |
| 0x008C57E4 | `Piano` | EQ Preset |  |
| 0x008C57F4 | `Rock` | EQ Preset |  |
| 0x008C57FC | `Small Speakers` | EQ Preset |  |
| 0x008C580C | `Spoken Word` | EQ Preset |  |
| 0x008C5818 | `Treble Booster` | EQ Preset |  |
| 0x008C5838 | `Vocal Booster` | EQ Preset |  |
| 0x008CEEE8 | `Hip Hop` | EQ Preset |  |
| 0x008CEEF4 | `Latin` | EQ Preset |  |
| 0x008CEF2C | `Rock` | EQ Preset |  |
| 0x008D87A0 | `Acoustic` | EQ Preset |  |
| 0x008D87AC | `Bass Booster` | EQ Preset |  |
| 0x008D87CC | `Classical` | EQ Preset |  |
| 0x008D87D8 | `Dance` | EQ Preset |  |
| 0x008D87E8 | `Electronic` | EQ Preset |  |
| 0x008D87FC | `Hip Hop` | EQ Preset |  |
| 0x008D8804 | `Jazz` | EQ Preset |  |
| 0x008D880C | `Latin` | EQ Preset |  |
| 0x008D8814 | `Loudness` | EQ Preset |  |
| 0x008D8820 | `Lounge` | EQ Preset |  |
| 0x008D8828 | `Piano` | EQ Preset |  |
| 0x008D8838 | `Rock` | EQ Preset |  |
| 0x008D8840 | `Small Speakers` | EQ Preset |  |
| 0x008D8850 | `Spoken Word` | EQ Preset |  |
| 0x008D885C | `Treble Booster` | EQ Preset |  |
| 0x008D887C | `Vocal Booster` | EQ Preset |  |
| 0x008E1F74 | `Acoustic` | EQ Preset |  |
| 0x008E1F80 | `Bass Booster` | EQ Preset |  |
| 0x008E1FA0 | `Classical` | EQ Preset |  |
| 0x008E1FAC | `Dance` | EQ Preset |  |
| 0x008E1FBC | `Electronic` | EQ Preset |  |
| 0x008E1FD0 | `Hip Hop` | EQ Preset |  |
| 0x008E1FD8 | `Jazz` | EQ Preset |  |
| 0x008E1FE0 | `Latin` | EQ Preset |  |
| 0x008E1FE8 | `Loudness` | EQ Preset |  |
| 0x008E1FF4 | `Lounge` | EQ Preset |  |
| 0x008E1FFC | `Piano` | EQ Preset |  |
| 0x008E200C | `Rock` | EQ Preset |  |
| 0x008E2014 | `Small Speakers` | EQ Preset |  |
| 0x008E2024 | `Spoken Word` | EQ Preset |  |
| 0x008E2030 | `Treble Booster` | EQ Preset |  |
| 0x008E2050 | `Vocal Booster` | EQ Preset |  |

---
