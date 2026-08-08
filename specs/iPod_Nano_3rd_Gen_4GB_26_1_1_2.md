# iPod Nano 3rd Gen - RetailOS 1.1.2 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.1.2 |
| **IPSW** | iPod_26.1.1.2.ipsw |
| **Device** | iPod Nano 3rd Gen (2008, 4/8GB NAND, Click Wheel, Cover Flow, Video) |
| **UpdaterFamilyID** | 26 |
| **Binary Size** | 10,790,960 bytes (10.29 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,788,912 bytes |
| **Total Strings (>=4)** | 72,485 |
| **Function Prologues** | 22,709 (ARM: 17,472, Thumb: 5,237) |
| **DRAM References** | 105,630 |
| **Peripheral Refs** | 7,430 |
| **Build** | N46FirmwareWin-435 |
| **SoC** | S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Codename** | N46 |
| **DFU PID** | 0x1229 |
| **SHA-256** | `27877a2dc4cd70a38f6f3f2e4ba7a9866beeec2552bd17da22a275dc036b4b46` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000A0788 | `TSilverCntlr` | Known | Controller |
| 0x000A07A0 | `TCExtrasMenu` | Known | Controller |
| 0x000A07B8 | `TCGameScreen` | Known | Controller |
| 0x000A07D0 | `TCGamesMenu` | Known | Controller |
| 0x000A07E4 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x000A080C | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x000A0834 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x000A0860 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x000A0884 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x000A08AC | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x000A08D4 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x000A08FC | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x000A0924 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x000A094C | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x000A097C | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x000A09A8 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x000A09D8 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x000A0A00 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x000A0A28 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x000A0A54 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x000A0A7C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x000A0AA4 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x000A0AD4 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x000A0B04 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x000A0C0C | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x000A0C3C | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x000A0C64 | `TCRentalNotification` | Known | Controller |
| 0x000A0C84 | `TCRentalInfo` | Known | Controller |
| 0x000A0C9C | `TCRentalConfirmDelete` | Known | Controller |
| 0x000A0CBC | `TCRentalDispatcher` | Known | Controller |
| 0x000A0CD8 | `TSilverGlobalCntlr` | Known | Controller |
| 0x000A0CF4 | `TSilverTrainerCntlr` | Known | Controller |
| 0x000F9DEC | `TCSlideshowLCD` | Known | Controller |
| 0x000F9E04 | `TCSlideshowTVOut` | Known | Controller |
| 0x000F9E20 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x000F9E40 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00123328 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x00123354 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x00123380 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x001233A8 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x001233D4 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x001233FC | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x00123428 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x0012A630 | `TCRemoteUI` | Known | Controller |
| 0x0012A644 | `TCUnsupported` | Known | Controller |
| 0x001309F0 | `TCSpeakers` | Known | Controller |
| 0x00130A04 | `TCEQSetting` | Known | Controller |
| 0x0015C528 | `TCSportTimer` | Known | Controller |
| 0x0015C540 | `TCSportTimerMenu` | Known | Controller |
| 0x0015C55C | `TCSportTimerSessionScreen` | Known | Controller |
| 0x0015C580 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x0015D900 | `TCVoiceMemos` | Known | Controller |
| 0x0015D918 | `TCVoiceMemosMenu` | Known | Controller |
| 0x0015D934 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x0015D954 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x0015D974 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x0016F040 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x0016F068 | `TCSettings_MainMenu` | Known | Controller |
| 0x0016F084 | `TCSettings_MusicMenu` | Known | Controller |
| 0x0016F0A4 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x0016F0C4 | `TCSettings_Brightness` | Known | Controller |
| 0x0016F0E4 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0016F108 | `TCSettings_EQ` | Known | Controller |
| 0x0016F120 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x0016F148 | `TCSettings_RadioRegions` | Known | Controller |
| 0x0016F168 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0016F18C | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x0016F1B0 | `TCDateTimeScreen` | Known | Controller |
| 0x0016F1CC | `TCTimeZoneScreen` | Known | Controller |
| 0x0016F1E8 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0016F210 | `TCFirstBoot` | Known | Controller |
| 0x00184A3C | `TCDemoMode` | Known | Controller |
| 0x001B18E0 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x001B1900 | `TCAddressViewerDetails` | Known | Controller |
| 0x001B1920 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x001B1944 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x001E1DB4 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x001E1DD8 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x001E9DAC | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x002877C0 | `TC_LockDialog` | Known | Controller |
| 0x002877D8 | `TC_LockScreen` | Known | Controller |
| 0x002877F0 | `TC_LockediPod` | Known | Controller |
| 0x00287808 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x0028782C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0028D6D8 | `TCClock` | Known | Controller |
| 0x0028D6E8 | `TCClockCityMenu` | Known | Controller |
| 0x0028D700 | `TCClockRegionMenu` | Known | Controller |
| 0x0028D71C | `TCAlarmMenu` | Known | Controller |
| 0x0028D730 | `TCSleepTimerMenu` | Known | Controller |
| 0x0028D74C | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x0028D76C | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x0028D794 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x0028D7B8 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0028D7DC | `TCAlarmDatePicker` | Known | Controller |
| 0x0028D7F8 | `TCAlarmTriggered` | Known | Controller |
| 0x00294718 | `TCNotesDispatcher` | Known | Controller |
| 0x00294734 | `TCNotesLoading` | Known | Controller |
| 0x0029474C | `TCNotesList` | Known | Controller |
| 0x00294760 | `TCNotesContents` | Known | Controller |
| 0x003C6DA8 | `TCAlarmTriggered` | Known | Controller |
| 0x003C6DBC | `TSilverCntlr` | Known | Controller |
| 0x003C6DDC | `TCClock` | Known | Controller |
| 0x003C6DE4 | `TCClockRegionMenu` | Known | Controller |
| 0x003C6DF8 | `TCClockCityMenu` | Known | Controller |
| 0x003C6E08 | `TCAlarmMenu` | Known | Controller |
| 0x003C6E14 | `TCSleepTimerMenu` | Known | Controller |
| 0x003C6E28 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003C6E40 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003C6E60 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003C6E7C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003C6E98 | `TCAlarmDatePicker` | Known | Controller |
| 0x003C6ED0 | `TSilverCntlr` | Known | Controller |
| 0x003C6EF0 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003C7080 | `TSilverCntlr` | Known | Controller |
| 0x003C70A0 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x003C70C0 | `TCSettings_Brightness` | Known | Controller |
| 0x003C70D8 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x003C70F4 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x003C7114 | `TCSettings_RadioRegions` | Known | Controller |
| 0x003C712C | `TCSettings_EQ` | Known | Controller |
| 0x003C713C | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x003C7158 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x003C7178 | `TCFirstBoot` | Known | Controller |
| 0x003C7184 | `TCSettings_MainMenu` | Known | Controller |
| 0x003C7198 | `TCSettings_MusicMenu` | Known | Controller |
| 0x003C71B0 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003C71C8 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x003C71E4 | `TCDateTimeScreen` | Known | Controller |
| 0x003C71F8 | `TCTimeZoneScreen` | Known | Controller |
| 0x003CE1FC | `TSilverCntlr` | Known | Controller |
| 0x003CE21C | `TCClock` | Known | Controller |
| 0x003CE224 | `TCClockRegionMenu` | Known | Controller |
| 0x003CE238 | `TCClockCityMenu` | Known | Controller |
| 0x003CE248 | `TCAlarmMenu` | Known | Controller |
| 0x003CE254 | `TCSleepTimerMenu` | Known | Controller |
| 0x003CE268 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003CE2E0 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003CE300 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003CE31C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003CE350 | `TCAlarmDatePicker` | Known | Controller |
| 0x003CE364 | `TCAlarmTriggered` | Known | Controller |
| 0x003CFDE0 | `TSilverCntlr` | Known | Controller |
| 0x003CFE00 | `TC_LockDialog` | Known | Controller |
| 0x003CFE10 | `TC_LockScreen` | Known | Controller |
| 0x003CFE20 | `TC_LockediPod` | Known | Controller |
| 0x003CFE30 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003CFE4C | `TCLockChosenDispatcher` | Known | Controller |
| 0x003CFE64 | `TSilverCntlr` | Known | Controller |
| 0x003CFFCC | `TSilverCntlr` | Known | Controller |
| 0x003CFFE8 | `TSilverCntlr` | Known | Controller |
| 0x003D004C | `TSilverCntlr` | Known | Controller |
| 0x003D006C | `TCNotesDispatcher` | Known | Controller |
| 0x003D0080 | `TCNotesLoading` | Known | Controller |
| 0x003D0090 | `TCNotesBase` | Known | Controller |
| 0x003D009C | `TCNotesList` | Known | Controller |
| 0x003D00A8 | `TCNotesContents` | Known | Controller |
| 0x003D00B8 | `TSilverCntlr` | Known | Controller |
| 0x003D00D8 | `TCRemoteUI` | Known | Controller |
| 0x003D00E4 | `TCUnsupported` | Known | Controller |
| 0x003D00F4 | `TSilverCntlr` | Known | Controller |
| 0x003D0158 | `TSilverCntlr` | Known | Controller |
| 0x003D0178 | `TCSportTimer` | Known | Controller |
| 0x003D0188 | `TCSportTimerMenu` | Known | Controller |
| 0x003D019C | `TCSportTimerSessionScreen` | Known | Controller |
| 0x003D01B8 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x003D0310 | `TSilverCntlr` | Known | Controller |
| 0x003D0330 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003D034C | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003D036C | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003D038C | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003D03B4 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003D03D8 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003D0400 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003D0420 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003D0440 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003D0460 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003D0480 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003D04A8 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003D0804 | `TSilverCntlr` | Known | Controller |
| 0x003D092C | `TSilverCntlr` | Known | Controller |
| 0x003D094C | `TCDemoMode` | Known | Controller |
| 0x003D0958 | `TCClock` | Known | Controller |
| 0x003D0960 | `TCClockRegionMenu` | Known | Controller |
| 0x003D0974 | `TCClockCityMenu` | Known | Controller |
| 0x003D0984 | `TCAlarmMenu` | Known | Controller |
| 0x003D0990 | `TCSleepTimerMenu` | Known | Controller |
| 0x003D09A4 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003D09BC | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003D09DC | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003D09F8 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003D0A14 | `TCAlarmDatePicker` | Known | Controller |
| 0x003D0A28 | `TCAlarmTriggered` | Known | Controller |
| 0x003D0A48 | `TSilverCntlr` | Known | Controller |
| 0x003D0A64 | `TSilverCntlr` | Known | Controller |
| 0x003D0A74 | `TSilverCntlr` | Known | Controller |
| 0x003D0A94 | `TCVoiceMemos` | Known | Controller |
| 0x003D0AA4 | `TCVoiceMemosMenu` | Known | Controller |
| 0x003D0AB8 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x003D0AD0 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x003D0AE8 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x003D0B08 | `TSilverCntlr` | Known | Controller |
| 0x003D0B68 | `TSilverCntlr` | Known | Controller |
| 0x003D0BD4 | `TSilverCntlr` | Known | Controller |
| 0x003D1E70 | `TSilverCntlr` | Known | Controller |
| 0x003D1F7C | `TSilverCntlr` | Known | Controller |
| 0x003DA804 | `TSilverCntlr` | Known | Controller |
| 0x003DA824 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x003DA83C | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x003DA858 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x003DA878 | `TCAddressViewerDetails` | Known | Controller |
| 0x003DA890 | `TSilverCntlr` | Known | Controller |
| 0x003DA8B0 | `TSilverCalendarCntlr_Base` | Known | Controller |
| 0x003DA8CC | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x003DA8F0 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x003DA914 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x003DA934 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x003DA958 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x003DA978 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x003DA99C | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x003DAB74 | `TSilverCntlr` | Known | Controller |
| 0x003DAB94 | `TC_LockDialog` | Known | Controller |
| 0x003DABA4 | `TC_LockScreen` | Known | Controller |
| 0x003DABB4 | `TC_LockediPod` | Known | Controller |
| 0x003DABC4 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003DABE8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003DAC9C | `TSilverCntlr` | Known | Controller |
| 0x003DACBC | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003DACD8 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003DACF8 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003DAD18 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003DAD40 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003DAD64 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003DAD8C | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003DADAC | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003DADCC | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003DADEC | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003DAE0C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003DAE34 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003DAE5C | `TSilverCntlr` | Known | Controller |
| 0x003DAF7C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003DAF98 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003DAFB8 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003DAFD8 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003DB000 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003DB024 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003DB04C | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003DB06C | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003DB08C | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003DB0AC | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003DB0CC | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003DB0F4 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003DB11C | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003DB13C | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003DB15C | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003DB180 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003DB1A0 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003DB1C4 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003DB1EC | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003DB218 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003DB238 | `TCRentalNotification` | Known | Controller |
| 0x003DB250 | `TCRentalInfo` | Known | Controller |
| 0x003DB260 | `TCRentalConfirmDelete` | Known | Controller |
| 0x003DB278 | `TCRentalDispatcher` | Known | Controller |
| 0x003DBB68 | `TSilverCntlr` | Known | Controller |
| 0x003DBC2C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003DBC48 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003DBC68 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003DBC88 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003DBCB0 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003DBCD4 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003DBCFC | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003DBD1C | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003DBD3C | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003DBD5C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003DBD7C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003DBDA4 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003DBDF4 | `TCSlideshowTVOut` | Known | Controller |
| 0x003DBE08 | `TCSlideshowLCD` | Known | Controller |
| 0x003DBE18 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x003DBE30 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x003DBE50 | `TSilverCntlr` | Known | Controller |
| 0x003DBE7C | `TSilverCntlr` | Known | Controller |
| 0x003DBE9C | `TCUnsupported` | Known | Controller |
| 0x003DBEBC | `TSilverCntlr` | Known | Controller |
| 0x003DBEFC | `TSilverCntlr` | Known | Controller |
| 0x003DBF1C | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x003DBF38 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x003DBF50 | `TSilverCntlr` | Known | Controller |
| 0x003DBF70 | `TCSpeakers` | Known | Controller |
| 0x003DBF7C | `TCEQSetting` | Known | Controller |
| 0x003DC024 | `TSilverCntlr` | Known | Controller |
| 0x003DC034 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003DC050 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003DC070 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003DC090 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003DC0B8 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003DC0DC | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003DC104 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003DC124 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003DC144 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003DC164 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003DC184 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003DC1AC | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003DC754 | `TSilverCntlr` | Known | Controller |
| 0x003DC778 | `TSilverCntlr` | Known | Controller |
| 0x003DC7E4 | `TSilverCntlr` | Known | Controller |
| 0x003DC804 | `TCExtrasMenu` | Known | Controller |
| 0x003DC814 | `TCGamesMenu` | Known | Controller |
| 0x003DC820 | `TCGameScreen` | Known | Controller |
| 0x003DC830 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x003DC850 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x003DC870 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x003DC890 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x003DC8B4 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003DC8D0 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003DC8F0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003DC910 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003DC938 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003DC95C | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003DC984 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003DC9A4 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003DC9C4 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003DC9E4 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003DCA04 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003DCA2C | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003DCA54 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003DCA74 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003DCA94 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003DCAB8 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003DCAD8 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003DCAFC | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003DCB24 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003DCB50 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003DCB70 | `TCRentalNotification` | Known | Controller |
| 0x003DCB88 | `TCRentalInfo` | Known | Controller |
| 0x003DCB98 | `TCRentalConfirmDelete` | Known | Controller |
| 0x003DCBB0 | `TCRentalDispatcher` | Known | Controller |
| 0x003DCBC4 | `TSilverGlobalCntlr` | Known | Controller |
| 0x003DCBD8 | `TSilverTrainerCntlr` | Known | Controller |
| 0x00467740 | `TCCCCCCCCCCCCCCCCCCCCCCT` | Known | Controller |
| 0x00753C66 | `TCNotesDispatcher"` | Known | Controller |
| 0x00753D25 | `TCLockChosenDispatcher"` | Known | Controller |
| 0x00753DE8 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x0075DE4D | `TCNotesDispatcher"` | Known | Controller |
| 0x0075DFAF | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x00775610 | `TCExtrasMenuTCAddressViewerMainMenu` | Known | Controller |
| 0x00775634 | `TCAddressViewerDetails` | Known | Controller |
| 0x0077564C | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x00775668 | `TCAlarmMenu` | Known | Controller |
| 0x00775674 | `TCSleepTimerMenuTCAlarmPropertiesMenu` | Known | Controller |
| 0x0077569C | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x007756BC | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x007756D8 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x007756F4 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00775710 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0077572C | `TCAlarmDatePicker` | Known | Controller |
| 0x00775740 | `TCAlarmDatePicker` | Known | Controller |
| 0x00775754 | `TCAlarmTriggeredTSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x00775780 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x007757A4 | `TSilverCalendarCntlr_MonthViewerTSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x007757E4 | `TSilverCalendarCntlr_EventViewerTSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x00775824 | `TSilverCalendarCntlr_EventViewerTSilverCntlrTCClockRegionMenu` | Known | Controller |
| 0x00775864 | `TCClockCityMenu` | Known | Controller |
| 0x00775874 | `TCClockCityMenu` | Known | Controller |
| 0x00775884 | `TCClockCityMenu` | Known | Controller |
| 0x00775894 | `TCClockCityMenu` | Known | Controller |
| 0x007758A4 | `TCClockCityMenu` | Known | Controller |
| 0x007758B4 | `TCClockCityMenu` | Known | Controller |
| 0x007758C4 | `TCClockCityMenu` | Known | Controller |
| 0x007758D4 | `TCClockCityMenu` | Known | Controller |
| 0x007758E4 | `TCClock` | Known | Controller |
| 0x007758FC | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDe` | Known | Controller |
| 0x00775954 | `TCGamesMenu` | Known | Controller |
| 0x00775960 | `TCGameScreenTC_LockediPod` | Known | Controller |
| 0x0077597C | `TC_LockDialog` | Known | Controller |
| 0x0077598C | `TC_LockScreen` | Known | Controller |
| 0x0077599C | `TC_VolumeLimitLockScreenTSilverCntlrTSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x007759E0 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00775A00 | `TSilverMediaListCntlr_AudiobooksTSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x00775A48 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x00775A64 | `TSilverMediaListCntlr_AlbumsTSilverMediaListCntlr_Artists` | Known | Controller |
| 0x00775AA0 | `TSilverMediaListCntlr_GenresTSilverMediaListCntlr_Composers` | Known | Controller |
| 0x00775ADC | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x00775AFC | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x00775B24 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x00775B44 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x00775B64 | `TSilverMediaListCntlr_TVEpisodesTSilverMediaListCntlr_MoviesTSilverMediaListCntl` | Known | Controller |
| 0x00775BC0 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x00775BE8 | `TSilverMediaListCntlr_VideoPlaylistsTSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x00775C2C | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x00775C58 | `TSilverCntlrTSilverMainMediaListCntlr_VideosTSilverSettingsVideoCntlr` | Known | Controller |
| 0x00775CA0 | `TCFirstBoot` | Known | Controller |
| 0x00775D48 | `TCRentalInfoTCRentalConfirmDelete` | Known | Controller |
| 0x00775D6C | `TSilverCntlrTCRentalNotificationTCRentalNotificationTCRentalNotificationTCNotesL` | Known | Controller |
| 0x00775DC4 | `TCNotesList` | Known | Controller |
| 0x00775DD0 | `TCNotesList` | Known | Controller |
| 0x00775DDC | `TCNotesContents` | Known | Controller |
| 0x00775DEC | `TCNotesContents` | Known | Controller |
| 0x00775DFC | `TCNotesContents` | Known | Controller |
| 0x00775E0C | `TCNotesContents` | Known | Controller |
| 0x00775EC8 | `TCSlideshowLCD` | Known | Controller |
| 0x00775ED8 | `TCSlideshowTVOutTCSlideshow_TVOutAskTCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00775F28 | `TCRemoteUI` | Known | Controller |
| 0x00775F34 | `TCUnsupported` | Known | Controller |
| 0x00775F44 | `TSilverSettingsMenuListCntlrTSilverSettingsMenuListCntlrTSilverSettingsMenuListC` | Known | Controller |
| 0x00775FAC | `TCSettings_MusicMenuTCSettings_VolumeLimit` | Known | Controller |
| 0x00775FD8 | `TCSettings_Brightness` | Known | Controller |
| 0x00775FF0 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0077600C | `TCSettings_AudiobookSettingsTCSettings_RadioRegions` | Known | Controller |
| 0x00776040 | `TCSettings_EQ` | Known | Controller |
| 0x00776050 | `TSilverSettingsMenuListCntlrTSilverCntlrTCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x00776098 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x007760B4 | `TCSettings_MainMenu` | Known | Controller |
| 0x007760C8 | `TCSettings_MusicMenuTCSportTimerTCSportTimerMenuTCSportTimerSessionScreen` | Known | Controller |
| 0x0077622C | `TSilverCntlrTTrainerEndSessionCntlr` | Known | Controller |
| 0x007762A4 | `TSilverCntlrTTrainerCalibrateWalkMenuCntlr` | Known | Controller |
| 0x00776538 | `TCVoiceMemosTCVoiceMemosMainMenuTCVoiceMemosMainMenuTCVoiceMemosMainMenuTSearchC` | Known | Controller |
| 0x00776598 | `TCEQSetting` | Known | Controller |
| 0x00776646 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x00777949 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x0077D552 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077D5B0 | `TCNotesDispatcher` | Known | Controller |
| 0x0077F18E | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077F1EC | `TCNotesDispatcher` | Known | Controller |
| 0x00780DCA | `TCLockChosenDispatcher` | Known | Controller |
| 0x00780E28 | `TCNotesDispatcher` | Known | Controller |
| 0x00782A06 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00782A64 | `TCNotesDispatcher` | Known | Controller |
| 0x00784642 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007846A0 | `TCNotesDispatcher` | Known | Controller |
| 0x0078627E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007862DC | `TCNotesDispatcher` | Known | Controller |
| 0x00787EBA | `TCLockChosenDispatcher` | Known | Controller |
| 0x00787F18 | `TCNotesDispatcher` | Known | Controller |
| 0x00789AF6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00789B54 | `TCNotesDispatcher` | Known | Controller |
| 0x0078B732 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078B790 | `TCNotesDispatcher` | Known | Controller |
| 0x0078D36E | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078D3CC | `TCNotesDispatcher` | Known | Controller |
| 0x0078EFAA | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078F008 | `TCNotesDispatcher` | Known | Controller |
| 0x00790BE6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00790C44 | `TCNotesDispatcher` | Known | Controller |
| 0x00792822 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00792880 | `TCNotesDispatcher` | Known | Controller |
| 0x0079445E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007944BC | `TCNotesDispatcher` | Known | Controller |
| 0x0079609A | `TCLockChosenDispatcher` | Known | Controller |
| 0x007960F8 | `TCNotesDispatcher` | Known | Controller |
| 0x00797CD6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00797D34 | `TCNotesDispatcher` | Known | Controller |
| 0x00799912 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00799970 | `TCNotesDispatcher` | Known | Controller |
| 0x0079B54E | `TCLockChosenDispatcher` | Known | Controller |
| 0x0079B5AC | `TCNotesDispatcher` | Known | Controller |
| 0x0079D18A | `TCLockChosenDispatcher` | Known | Controller |
| 0x0079D1E8 | `TCNotesDispatcher` | Known | Controller |
| 0x0079EDC6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0079EE24 | `TCNotesDispatcher` | Known | Controller |
| 0x007A0A02 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A0A60 | `TCNotesDispatcher` | Known | Controller |
| 0x007A263E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A269C | `TCNotesDispatcher` | Known | Controller |
| 0x007A427A | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A42D8 | `TCNotesDispatcher` | Known | Controller |
| 0x007A5EB6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A5F14 | `TCNotesDispatcher` | Known | Controller |
| 0x007A7AF2 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A7B50 | `TCNotesDispatcher` | Known | Controller |
| 0x007A972E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A978C | `TCNotesDispatcher` | Known | Controller |
| 0x007AB36A | `TCLockChosenDispatcher` | Known | Controller |
| 0x007AB3C8 | `TCNotesDispatcher` | Known | Controller |
| 0x007ACFA6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007AD004 | `TCNotesDispatcher` | Known | Controller |
| 0x007AEBE2 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007AEC40 | `TCNotesDispatcher` | Known | Controller |
| 0x007B081E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B087C | `TCNotesDispatcher` | Known | Controller |
| 0x007B245A | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B24B8 | `TCNotesDispatcher` | Known | Controller |
| 0x007B4096 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B40F4 | `TCNotesDispatcher` | Known | Controller |
| 0x007B5CD2 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B5D30 | `TCNotesDispatcher` | Known | Controller |
| 0x007B790E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B796C | `TCNotesDispatcher` | Known | Controller |
| 0x007B954A | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B95A8 | `TCNotesDispatcher` | Known | Controller |
| 0x007BB186 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007BB1E4 | `TCNotesDispatcher` | Known | Controller |
| 0x007BCDC2 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007BCE20 | `TCNotesDispatcher` | Known | Controller |
| 0x007C89F8 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x007C8CBA | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x007C94F0 | `TCRentalDispatcher` | Known | Controller |
| 0x007C9DA8 | `TCRentalDispatcher` | Known | Controller |
| 0x007CA660 | `TCRentalDispatcher` | Known | Controller |
| 0x007CAF18 | `TCRentalDispatcher` | Known | Controller |
| 0x007CB7D0 | `TCRentalDispatcher` | Known | Controller |
| 0x007CC088 | `TCRentalDispatcher` | Known | Controller |
| 0x007CC940 | `TCRentalDispatcher` | Known | Controller |
| 0x007CD1F8 | `TCRentalDispatcher` | Known | Controller |
| 0x00920414 | `TCMockupModeNavScreen` | Known | Controller |
| 0x0092042C | `TSilverCntlr` | Known | Controller |
| 0x0092044C | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x00920484 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x009204A4 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x009204C4 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x009204E8 | `TCExtrasMenu` | Known | Controller |
| 0x009205F8 | `TSilverCntlr` | Known | Controller |
| 0x00920618 | `TCSlideshowTVOut` | Known | Controller |
| 0x0092062C | `TCSlideshowLCD` | Known | Controller |
| 0x0092063C | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x00920654 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00920690 | `TSilverCntlr` | Known | Controller |
| 0x0092070C | `TCSlideshowTVOut` | Known | Controller |
| 0x00920720 | `TCSlideshowLCD` | Known | Controller |
| 0x00920730 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x00920748 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00920768 | `TSilverCntlr` | Known | Controller |
| 0x00920A20 | `TSilverCntlr` | Known | Controller |
| 0x00920A40 | `TCGamesMenu` | Known | Controller |
| 0x00920A4C | `TCGameScreen` | Known | Controller |
| 0x009DED6F | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |

---

## 2. Settings (Toggle/Show/TC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00139CF0 | `ShowSetting_EQ` | Known | User setting |
| 0x001EB560 | `ToggleSetting_Repeat` | Known | User setting |
| 0x001EB57C | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001EB594 | `ToggleSetting_TVOut` | Known | User setting |
| 0x001EB5A8 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x0021B38C | `ShowSetting_Backlight` | Known | User setting |
| 0x00230134 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x00230150 | `ToggleSetting_Repeat` | Known | User setting |
| 0x00230168 | `ToggleSetting_SortBy` | Known | User setting |
| 0x00230180 | `ToggleSetting_ClassicUI` | Known | User setting |
| 0x00230198 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x002301B4 | `ToggleSetting_Clicker` | Known | User setting |
| 0x002301CC | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x002301EC | `ToggleSetting_24HourClock` | Known | User setting |
| 0x00230208 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x00230224 | `ShowSetting_Shuffle` | Known | User setting |
| 0x002303D0 | `ShowSetting_Repeat` | Known | User setting |
| 0x002303E4 | `ShowSetting_About` | Known | User setting |
| 0x002303F8 | `ShowSetting_MainMenu` | Known | User setting |
| 0x00230410 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x00230428 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x00230440 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x0023045C | `ShowSetting_Brightness` | Known | User setting |
| 0x00230474 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x0023048C | `ShowSetting_RadioRegions` | Known | User setting |
| 0x002304A8 | `ShowSetting_EQ` | Known | User setting |
| 0x002304B8 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x00230654 | `ShowSetting_Clicker` | Known | User setting |
| 0x00230668 | `ShowSetting_DateAndTime` | Known | User setting |
| 0x00230680 | `ShowSetting_SortBy` | Known | User setting |
| 0x00230694 | `ShowSetting_ClassicUI` | Known | User setting |
| 0x002306AC | `ShowSetting_Language` | Known | User setting |
| 0x002306C4 | `ShowSetting_Legal` | Known | User setting |
| 0x002306D8 | `ShowSetting_ResetAll` | Known | User setting |
| 0x0075CCD5 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x0075CD85 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x0075F41A | `ShowSetting_About` | Known | User setting |
| 0x0075F4BC | `ToggleSetting_Shuffle` | Known | User setting |
| 0x0075F500 | `ShowSetting_Shuffle` | Known | User setting |
| 0x0075F577 | `ToggleSetting_Repeat` | Known | User setting |
| 0x0075F5BA | `ShowSetting_Repeat` | Known | User setting |
| 0x0075F6C4 | `ShowSetting_MainMenu` | Known | User setting |
| 0x0075F7D4 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x0075F89C | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x0075F966 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x0075FA7E | `ShowSetting_Brightness` | Known | User setting |
| 0x0075FBB4 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x0075FCC5 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x0075FDC6 | `ShowSetting_EQ` | Known | User setting |
| 0x0075FE33 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x0075FE7A | `ShowSetting_SoundCheck` | Known | User setting |
| 0x0075FEF7 | `ToggleSetting_Clicker` | Known | User setting |
| 0x0075FF3B | `ShowSetting_Clicker` | Known | User setting |
| 0x007600A2 | `ToggleSetting_SortBy` | Known | User setting |
| 0x007600E5 | `ShowSetting_SortBy` | Known | User setting |
| 0x007601E6 | `ShowSetting_Language` | Known | User setting |
| 0x007602F6 | `ShowSetting_Legal` | Known | User setting |
| 0x00760427 | `ShowSetting_ResetAll` | Known | User setting |
| 0x00760598 | `ShowSetting_Backlight` | Known | User setting |
| 0x00760648 | `ShowSetting_Backlight` | Known | User setting |
| 0x007606F8 | `ShowSetting_Backlight` | Known | User setting |
| 0x007607A9 | `ShowSetting_Backlight` | Known | User setting |
| 0x0076085A | `ShowSetting_Backlight` | Known | User setting |
| 0x0076090B | `ShowSetting_Backlight` | Known | User setting |
| 0x007609BF | `ShowSetting_Backlight` | Known | User setting |
| 0x00760A6E | `ShowSetting_EQ` | Known | User setting |
| 0x00760AE3 | `ShowSetting_Language` | Known | User setting |
| 0x007DE30C | `ToggleSetting_Repeat` | Known | User setting |
| 0x007DE346 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x007DE408 | `ToggleSetting_TVOut` | Known | User setting |
| 0x007DE441 | `ToggleSetting_TVSignal` | Known | User setting |

---

## 3. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001583A0 | `MockupMode/MockupMode.xml` | Hidden | Developer Tool |
| 0x001588A0 | `MockupMode/` | Hidden | Developer Tool |
| 0x0026C710 | `Channel UnitTests` | Hidden | Developer Tool |
| 0x002C2BF5 | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x002C2C38 | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x002C2C4D | `RTXCbug> ` | Hidden | Developer Tool |
| 0x002C3629 | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x002D4E10 | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x0036F47D | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x0036F545 | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x003CC1A9 | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x007764A4 | `TTrainerLoadingCntlrTSilverCntlrTUnitTestSuiteCntlr` | Hidden | Developer Tool |
| 0x007764D8 | `TUnitTestSuiteTestsCntlrTSilverCntlrTCVoiceMemosMainMenuTCVoiceMemosMenuTCVoiceM` | Hidden | Developer Tool |
| 0x00814FB8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0085D488 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0086FA94 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008871F4 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00899460 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008A306C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008AC99C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008C19AC | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008CB538 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008F19AC | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0090FE2C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00919134 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x009CF66E | `Debug_ListItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x009CF686 | `Debug_MenuItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x009D029B | `DemoMode_Background_Image` | Hidden | Demo/Retail Mode |
| 0x009D12AE | `DemoMode_Icon_Image` | Hidden | Demo/Retail Mode |
| 0x009D3128 | `Debug_Menu_Title` | Hidden | Debug/Diagnostic |
| 0x009D314D | `UnitTest_Menu_Title` | Hidden | Developer Tool |
| 0x009DC23C | `UnitTestModel` | Hidden | Developer Tool |
| 0x009DCE74 | `UnitTest_ListItem` | Hidden | Developer Tool |
| 0x009DE406 | `DemoMode_View_Main` | Hidden | Demo/Retail Mode |
| 0x009DE5EE | `DemoMode_Icon` | Hidden | Demo/Retail Mode |
| 0x009DF434 | `UnitTestApp` | Hidden | Developer Tool |
| 0x009DFA96 | `Debug_ListItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009DFAB1 | `Debug_MenuItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009E02BC | `DemoMode_Text_Color` | Hidden | Demo/Retail Mode |
| 0x009E06C5 | `Debug_ListItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009E06DC | `Debug_MenuItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009E5181 | `Debug_ListItem_UnitTest` | Hidden | Developer Tool |
| 0x009E5199 | `Debug_MenuItem_UnitTest` | Hidden | Developer Tool |
| 0x009EA0C7 | `Debug_ListItem_Memory` | Hidden | Debug/Diagnostic |
| 0x009EA0DD | `Debug_MenuItem_Memory` | Hidden | Debug/Diagnostic |

---

## 4. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000677B | `"MeCCADecode` | Known | Audio system |
| 0x0014E988 | `AudioCodecs` | Known | Audio system |
| 0x0019376C | `MeCCA_PCM_Output.wav` | Known | Audio system |
| 0x001B0B1C | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x001BB428 | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x001BB630 | `MeCCAVideoDecode` | Known | Audio system |
| 0x0092DDA8 | `MeCCA_StreamCache` | Known | Audio system |

---

## 5. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F5BE8 | `HandleWheel` | Known | Event handler |
| 0x000F5BF4 | `HandlePlayPause` | Known | Event handler |
| 0x000F5C04 | `HandleSelectDown` | Known | Event handler |
| 0x000F5C18 | `HandleNext` | Known | Event handler |
| 0x000F5C24 | `HandlePrevious` | Known | Event handler |
| 0x000F5C34 | `HandleNextPushAndHold` | Known | Event handler |
| 0x000F5C4C | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000F5EE4 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000F5F04 | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x00101EC0 | `HandleSelect` | Known | Event handler |
| 0x00101ED4 | `HandleHilite` | Known | Event handler |
| 0x0010226C | `HandleEQSettingSelected` | Known | Event handler |
| 0x0010269C | `HandleSelect` | Known | Event handler |
| 0x001026B0 | `HandleGameHilited` | Known | Event handler |
| 0x00102960 | `HandleNotesSelected` | Known | Event handler |
| 0x00102978 | `HandleNotesPop` | Known | Event handler |
| 0x00102988 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00110A18 | `HandleVolumeWheel` | Known | Event handler |
| 0x00110A2C | `HandleVolumeChange` | Known | Event handler |
| 0x00110A40 | `HandleTimerDone` | Known | Event handler |
| 0x00110A50 | `HandleFrequencyChange` | Known | Event handler |
| 0x00110AC8 | `HandleTuning` | Known | Event handler |
| 0x00110AD8 | `HandleTuningSelect` | Known | Event handler |
| 0x00121290 | `HandleLock` | Known | Event handler |
| 0x001212A0 | `HandleAddressBook` | Known | Event handler |
| 0x00121988 | `HandleSelect` | Known | Event handler |
| 0x00121EC0 | `HandleExit` | Known | Event handler |
| 0x00121ED0 | `HandleLap` | Known | Event handler |
| 0x00121EDC | `HandleResume` | Known | Event handler |
| 0x00121EEC | `HandleStartStop` | Known | Event handler |
| 0x00122174 | `HandleWheel` | Known | Event handler |
| 0x00122184 | `HandlePlayPause` | Known | Event handler |
| 0x00122194 | `HandleSelectDown` | Known | Event handler |
| 0x001221A8 | `HandleHilite` | Known | Event handler |
| 0x0012BDE8 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x00139F24 | `HandleExitUnsupported` | Known | Event handler |
| 0x00145534 | `HandleBasicSelected` | Known | Event handler |
| 0x0014554C | `HandleTimedWorkoutSelected` | Known | Event handler |
| 0x00145568 | `HandleDistanceWorkoutSelected` | Known | Event handler |
| 0x00145588 | `HandleCaloriesWorkoutSelected` | Known | Event handler |
| 0x001455A8 | `HandleSelectWorkout` | Known | Event handler |
| 0x0015410C | `HandleNotesPop` | Known | Event handler |
| 0x00154120 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00155004 | `HandleSelect` | Known | Event handler |
| 0x00155018 | `HandleWheel` | Known | Event handler |
| 0x00155024 | `HandleImageNext` | Known | Event handler |
| 0x00155034 | `HandleImagePrev` | Known | Event handler |
| 0x00155044 | `HandleImageLast` | Known | Event handler |
| 0x00155054 | `HandleImageFirst` | Known | Event handler |
| 0x00155068 | `HandlePlayPause` | Known | Event handler |
| 0x00155078 | `HandlePlay` | Known | Event handler |
| 0x00155084 | `HandlePause` | Known | Event handler |
| 0x00169B44 | `HandleSelectCity` | Known | Event handler |
| 0x00169B5C | `HandleHighlightCity` | Known | Event handler |
| 0x0016AA84 | `HandleWantPopFlow` | Known | Event handler |
| 0x0016AA9C | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x0016AAB8 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0016AAD4 | `HandleFlowNext` | Known | Event handler |
| 0x0016AAE4 | `HandleFlowPrev` | Known | Event handler |
| 0x0016AAF4 | `HandleFlowWheel` | Known | Event handler |
| 0x0016AB04 | `HandleAlbumSelected` | Known | Event handler |
| 0x0016AB18 | `HandlePlayPause` | Known | Event handler |
| 0x0016AB28 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00195608 | `HandleLeaveAlarm` | Known | Event handler |
| 0x001959F8 | `HandleSelect` | Known | Event handler |
| 0x001968B8 | `HandleSelect` | Known | Event handler |
| 0x001968CC | `HandleWheel` | Known | Event handler |
| 0x001968D8 | `HandleImageNext` | Known | Event handler |
| 0x001968E8 | `HandleImagePrev` | Known | Event handler |
| 0x001968F8 | `HandleImageLast` | Known | Event handler |
| 0x00196908 | `HandleImageFirst` | Known | Event handler |
| 0x0019691C | `HandlePlayPause` | Known | Event handler |
| 0x0019692C | `HandlePlay` | Known | Event handler |
| 0x00196938 | `HandlePause` | Known | Event handler |
| 0x00196DD8 | `HandleNew` | Known | Event handler |
| 0x00196DE8 | `HandleClear` | Known | Event handler |
| 0x00196DF4 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x00196E10 | `HandleSelectIndexedSession` | Known | Event handler |
| 0x00197120 | `HandleWheel` | Known | Event handler |
| 0x00197130 | `HandleArrowUp` | Known | Event handler |
| 0x00197140 | `HandleArrowDown` | Known | Event handler |
| 0x00199364 | `HandleHiliteAlbum` | Known | Event handler |
| 0x0019937C | `HandleBrowseAlbum` | Known | Event handler |
| 0x00199390 | `HandlePlayPause` | Known | Event handler |
| 0x001B51B0 | `HandleSelect` | Known | Event handler |
| 0x001B5340 | `HandleSelectRegion` | Known | Event handler |
| 0x001B9DCC | `HandleChooseLink` | Known | Event handler |
| 0x001B9DE4 | `HandleChooseCalibrate` | Known | Event handler |
| 0x001B9DFC | `HandleUnlink` | Known | Event handler |
| 0x001CADA4 | `HandleImageWheel` | Known | Event handler |
| 0x001CADBC | `HandlePlayPause` | Known | Event handler |
| 0x001CADCC | `HandleBrowseLarge` | Known | Event handler |
| 0x001CADE0 | `HandleBrowseSmall` | Known | Event handler |
| 0x001CADF4 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x001CAE0C | `HandleImageNext` | Known | Event handler |
| 0x001CAE1C | `HandleImagePrev` | Known | Event handler |
| 0x001CAE2C | `HandleHilite` | Known | Event handler |
| 0x001CAE3C | `HandleImageLast` | Known | Event handler |
| 0x001CAE4C | `HandleImageFirst` | Known | Event handler |
| 0x001CAE60 | `HandleScreenNext` | Known | Event handler |
| 0x001CAE74 | `HandleScreenPrev` | Known | Event handler |
| 0x001CD758 | `HandlePlayPause` | Known | Event handler |
| 0x001CD76C | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x001CD788 | `HandleNext` | Known | Event handler |
| 0x001CD794 | `HandleNextPressAndHold` | Known | Event handler |
| 0x001CD7AC | `HandlePrevious` | Known | Event handler |
| 0x001CD7BC | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x001CD7D8 | `HandleRemotePlayPause` | Known | Event handler |
| 0x001CD7F0 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x001CD814 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x001CD82C | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x001CD844 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x001CDA14 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x001CDA2C | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x001CDA44 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x001CDA60 | `HandleRemoteStop` | Known | Event handler |
| 0x001CDA74 | `HandleRemotePlay` | Known | Event handler |
| 0x001CDA88 | `HandleRemotePause` | Known | Event handler |
| 0x001CDA9C | `HandleRemoteMute` | Known | Event handler |
| 0x001CDAB0 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x001CDAC8 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x001CDAE0 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x001CDAFC | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x001CDD20 | `HandleRemoteShuffle` | Known | Event handler |
| 0x001CDD34 | `HandleRemoteRepeat` | Known | Event handler |
| 0x001CDD48 | `HandleRemoteOn` | Known | Event handler |
| 0x001CDD58 | `HandleRemoteOff` | Known | Event handler |
| 0x001CDD68 | `HandleRemoteBacklight` | Known | Event handler |
| 0x001CDD80 | `HandleRemoteFFDown` | Known | Event handler |
| 0x001CDD94 | `HandleRemoteFFUp` | Known | Event handler |
| 0x001CDDA8 | `HandleRemoteRewDown` | Known | Event handler |
| 0x001CDDBC | `HandleRemoteRewUp` | Known | Event handler |
| 0x001CDDD0 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x001CDDE8 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x001CDDFC | `HandleRemoteSelectDown` | Known | Event handler |
| 0x001CDE14 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x001CDFE4 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x001CDFFC | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x001CE014 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x001CE030 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x001CE048 | `HandleRemoteEvent` | Known | Event handler |
| 0x001CE05C | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x001CE078 | `HandleAudioPlayPause` | Known | Event handler |
| 0x001CE090 | `HandleAudioNext` | Known | Event handler |
| 0x001CE0A0 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x001CE0BC | `HandleAudioPrevious` | Known | Event handler |
| 0x001CE0D0 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x001CE2D0 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x001CE2E8 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x001CE300 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x001CE318 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x001CE32C | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x001CE344 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x001CE35C | `HandleAudioStop` | Known | Event handler |
| 0x001CE36C | `HandleAudioPlay` | Known | Event handler |
| 0x001CE37C | `HandleAudioPause` | Known | Event handler |
| 0x001CE390 | `HandleAudioMute` | Known | Event handler |
| 0x001CE3A0 | `HandleAudioNextChapter` | Known | Event handler |
| 0x001CE3B8 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x001CE5D8 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x001CE5F0 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x001CE608 | `HandleAudioShuffle` | Known | Event handler |
| 0x001CE61C | `HandleAudioRepeat` | Known | Event handler |
| 0x001CE630 | `HandleAudioFFDown` | Known | Event handler |
| 0x001CE644 | `HandleAudioFFUp` | Known | Event handler |
| 0x001CE654 | `HandleAudioRewDown` | Known | Event handler |
| 0x001CE668 | `HandleAudioRewUp` | Known | Event handler |
| 0x001CE67C | `HandleVideoPlayPause` | Known | Event handler |
| 0x001CE694 | `HandleVideoNext` | Known | Event handler |
| 0x001CE6A4 | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x001CE6C0 | `HandleVideoPrevious` | Known | Event handler |
| 0x001CE6D4 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x001CE8DC | `HandleVideoStop` | Known | Event handler |
| 0x001CE8EC | `HandleVideoPlay` | Known | Event handler |
| 0x001CE8FC | `HandleVideoPause` | Known | Event handler |
| 0x001CE910 | `HandleVideoFFDown` | Known | Event handler |
| 0x001CE924 | `HandleVideoFFUp` | Known | Event handler |
| 0x001CE934 | `HandleVideoRewDown` | Known | Event handler |
| 0x001CE948 | `HandleVideoRewUp` | Known | Event handler |
| 0x001CE95C | `HandleVideoNextChapter` | Known | Event handler |
| 0x001CE974 | `HandleVideoPrevChapter` | Known | Event handler |
| 0x001CE98C | `HandleVideoNextFrame` | Known | Event handler |
| 0x001CE9A4 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x001CE9BC | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x001D1F54 | `HandleSelect` | Known | Event handler |
| 0x001D1F68 | `HandleMenu` | Known | Event handler |
| 0x001D1F74 | `HandleLinkCancelOption` | Known | Event handler |
| 0x001D1F8C | `HandleLinkNewRemote` | Known | Event handler |
| 0x001D1FA0 | `HandleCancelRemoteLinking` | Known | Event handler |
| 0x001D2300 | `HandleChoosePlaylistsPlay` | Known | Event handler |
| 0x001D2320 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x001D233C | `HandleNoneSelected` | Known | Event handler |
| 0x001D2350 | `HandleNowPlayingSelected` | Known | Event handler |
| 0x001D236C | `HandleMenuKeyNop` | Known | Event handler |
| 0x001D2380 | `HandlePlaylistSelected` | Known | Event handler |
| 0x001D2B4C | `HandlePauseWorkout` | Known | Event handler |
| 0x001D2B64 | `HandleEndWorkout` | Known | Event handler |
| 0x001D2B78 | `HandleResumeWorkout` | Known | Event handler |
| 0x001D2B8C | `HandleChooseMusic` | Known | Event handler |
| 0x001D2BA0 | `HandleMenuKeyNop` | Known | Event handler |
| 0x001DE9DC | `HandleMainMenu` | Known | Event handler |
| 0x001E2F20 | `HandlePowerSongSelected` | Known | Event handler |
| 0x001E2F3C | `HandlePowerSongChosen` | Known | Event handler |
| 0x001E2F54 | `HandleSelectTimedWorkout` | Known | Event handler |
| 0x001E37D8 | `HandleChooseRemoteSetting` | Known | Event handler |
| 0x001E37F8 | `HandleChooseRemoteLink` | Known | Event handler |
| 0x001E3810 | `HandleChooseRemoteUnlink` | Known | Event handler |
| 0x001E3B4C | `HandleSelectResume` | Known | Event handler |
| 0x001E3B64 | `HandleEndWorkout` | Known | Event handler |
| 0x001E9CC4 | `HandleSelect` | Known | Event handler |
| 0x001E9F6C | `HandleMusicMenu` | Known | Event handler |
| 0x001EA22C | `HandleSelect` | Known | Event handler |
| 0x001EA5B0 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x001EA5C8 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x001EA5E8 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x001EA60C | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x001EA628 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x001EAAC4 | `HandleWheel` | Known | Event handler |
| 0x001EAAD4 | `HandlePlayPause` | Known | Event handler |
| 0x001EAAE4 | `HandleSelectDown` | Known | Event handler |
| 0x001EAAF8 | `HandleNext` | Known | Event handler |
| 0x001EAB04 | `HandlePrevious` | Known | Event handler |
| 0x001EAB14 | `HandleNextPushAndHold` | Known | Event handler |
| 0x001EAB2C | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x001F0340 | `HandleChooseLast` | Known | Event handler |
| 0x001F0358 | `HandleChooseRecent` | Known | Event handler |
| 0x001F036C | `HandleChooseWorkout` | Known | Event handler |
| 0x001F0380 | `HandleChooseBest` | Known | Event handler |
| 0x001F0394 | `HandleChooseRecentWorkout` | Known | Event handler |
| 0x001F2A1C | `HandleSelect` | Known | Event handler |
| 0x001F2A30 | `HandleMenu` | Known | Event handler |
| 0x001FAA38 | `HandleFrequencyChosen` | Known | Event handler |
| 0x001FAA50 | `HandleDateChosen` | Known | Event handler |
| 0x001FAA64 | `HandleTimeChosen` | Known | Event handler |
| 0x001FAA78 | `HandleSoundChosen` | Known | Event handler |
| 0x001FAA8C | `HandleLabelChosen` | Known | Event handler |
| 0x001FAAA0 | `HandleDeleteChosen` | Known | Event handler |
| 0x001FBB80 | `HandleSelect` | Known | Event handler |
| 0x0020049C | `HandlePrev` | Known | Event handler |
| 0x002004AC | `HandleNext` | Known | Event handler |
| 0x002004B8 | `HandlePlayPause` | Known | Event handler |
| 0x00200C94 | `HandleChoosePowerPlay` | Known | Event handler |
| 0x00200CB0 | `HandleChooseUnit` | Known | Event handler |
| 0x00200CC4 | `HandleChooseVoicePrompts` | Known | Event handler |
| 0x0020963C | `HandleNextContact` | Known | Event handler |
| 0x00209654 | `HandlePreviousContact` | Known | Event handler |
| 0x0020CA18 | `HandleSelect` | Known | Event handler |
| 0x0020CCF4 | `HandleListChoose` | Known | Event handler |
| 0x00211728 | `HandleItemSelected` | Known | Event handler |
| 0x00211920 | `HandleRadioRegion` | Known | Event handler |
| 0x00211B08 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x00212354 | `HandleSelect` | Known | Event handler |
| 0x0021279C | `HandlePauseKey` | Known | Event handler |
| 0x002127B0 | `HandlePauseHold` | Known | Event handler |
| 0x002127C0 | `HandlePauseKeyNop` | Known | Event handler |
| 0x002127D4 | `HandleMenuKey` | Known | Event handler |
| 0x002127E4 | `HandleMenuKeyNop` | Known | Event handler |
| 0x002127F8 | `HandleWheel` | Known | Event handler |
| 0x00212848 | `HandleSelectKeyDown` | Known | Event handler |
| 0x0021285C | `HandleOnDemandPrompt` | Known | Event handler |
| 0x00212874 | `HandlePowerPlay` | Known | Event handler |
| 0x00217644 | `HandlePlayPause` | Known | Event handler |
| 0x002188B0 | `HandleSelect` | Known | Event handler |
| 0x00218B40 | `HandleChooseCustomTimedWorkout` | Known | Event handler |
| 0x00218B64 | `HandleChooseTimedWorkoutFromList` | Known | Event handler |
| 0x00218B88 | `HandleChooseCustomDistanceWorkout` | Known | Event handler |
| 0x00218BAC | `HandleChooseDistanceWorkoutFromList` | Known | Event handler |
| 0x00218BD0 | `HandleChooseCustomCalorieWorkout` | Known | Event handler |
| 0x00218BF4 | `HandleChooseCalorieWorkoutFromList` | Known | Event handler |
| 0x0021B668 | `HandleDelete` | Known | Event handler |
| 0x0021B67C | `HandleSelectLozinch` | Known | Event handler |
| 0x0021B924 | `HandleSelect` | Known | Event handler |
| 0x0021BBF0 | `HandleTVOutChanged` | Known | Event handler |
| 0x0021BC08 | `HandleTVSignalChanged` | Known | Event handler |
| 0x0021BC20 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x0021BC40 | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x0021BC60 | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x0021BC84 | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x0021BCA4 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x0021C590 | `HandleBegin` | Known | Event handler |
| 0x0021F814 | `HandleSelectKey` | Known | Event handler |
| 0x0021F9BC | `HandleSelect` | Known | Event handler |
| 0x00220738 | `HandlePlayPause` | Known | Event handler |
| 0x0022074C | `HandleWheel` | Known | Event handler |
| 0x00220758 | `HandleWheelRating` | Known | Event handler |
| 0x0022076C | `HandleWheelScrub` | Known | Event handler |
| 0x00220780 | `HandleWheelVolume` | Known | Event handler |
| 0x00220840 | `HandleMenuKey` | Known | Event handler |
| 0x002208AC | `HandleMenuLongpress` | Known | Event handler |
| 0x002208C0 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x002214C8 | `HandleSelect` | Known | Event handler |
| 0x00221D98 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00222C88 | `HandleSelect` | Known | Event handler |
| 0x00222C9C | `HandleHilite` | Known | Event handler |
| 0x00222CAC | `HandlePlayPause` | Known | Event handler |
| 0x00222CBC | `HandleAddToOTG` | Known | Event handler |
| 0x00222CCC | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00223D08 | `HandleWeightWheel` | Known | Event handler |
| 0x00223D20 | `HandleWeightSelect` | Known | Event handler |
| 0x00223D34 | `HandleDistanceWheel` | Known | Event handler |
| 0x00223D48 | `HandleDistanceSelect` | Known | Event handler |
| 0x00223D60 | `HandleTimeWheel` | Known | Event handler |
| 0x00223D70 | `HandleTimeSelect` | Known | Event handler |
| 0x00223D84 | `HandleCaloriesWheel` | Known | Event handler |
| 0x00223D98 | `HandleCaloriesSelect` | Known | Event handler |
| 0x00224364 | `HandleSelect` | Known | Event handler |
| 0x00224378 | `HandleSelectFromLastWorkout` | Known | Event handler |
| 0x00226C2C | `HandleLanguageAfterReset` | Known | Event handler |
| 0x0022743C | `HandleSelect` | Known | Event handler |
| 0x00227450 | `HandleWheel` | Known | Event handler |
| 0x0022745C | `HandleWheelProgress` | Known | Event handler |
| 0x00227470 | `HandleSelectProgress` | Known | Event handler |
| 0x00227488 | `HandleSelectVolume` | Known | Event handler |
| 0x0022749C | `HandleSelectScrub` | Known | Event handler |
| 0x002274B0 | `HandleSelectRating` | Known | Event handler |
| 0x002274C4 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x002274DC | `HandleSelectChapterArt` | Known | Event handler |
| 0x002274F4 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x00227510 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0022752C | `HandleWheelBrightness` | Known | Event handler |
| 0x00227674 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x00229488 | `HandleSelect` | Known | Event handler |
| 0x00229498 | `HandleSelectRating` | Known | Event handler |
| 0x002294AC | `HandleSelectProgress` | Known | Event handler |
| 0x002294C4 | `HandleWheelProgress` | Known | Event handler |
| 0x002294D8 | `HandleSelectScrub` | Known | Event handler |
| 0x002294EC | `HandleWheelBrightness` | Known | Event handler |
| 0x00229504 | `HandleNextTrackGotoScrub` | Known | Event handler |
| 0x00229520 | `HandlePrevTrackGotoScrub` | Known | Event handler |
| 0x0022953C | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x0022C9A4 | `HandleSelectWalking` | Known | Event handler |
| 0x0022C9BC | `HandleSelectRunning` | Known | Event handler |
| 0x00230710 | `HandleLanguage` | Known | Event handler |
| 0x00230720 | `HandleResetAllSettings` | Known | Event handler |
| 0x00230738 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x00230B00 | `HandleUnlinkRemote` | Known | Event handler |
| 0x002315F0 | `HandleSelect` | Known | Event handler |
| 0x00231820 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x00232F98 | `Handle400MetersRun` | Known | Event handler |
| 0x00232FB0 | `HandleCustomRun` | Known | Event handler |
| 0x00232FC0 | `HandleResetToDefault` | Known | Event handler |
| 0x00233420 | `HandleSelect_Basic` | Known | Event handler |
| 0x00233438 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x002354F8 | `HandleSelect` | Known | Event handler |
| 0x00235694 | `HandleSelect` | Known | Event handler |
| 0x00235934 | `HandleNextDay` | Known | Event handler |
| 0x00235948 | `HandlePreviousDay` | Known | Event handler |
| 0x0023614C | `HandleMusicHilited` | Known | Event handler |
| 0x00236164 | `HandleVideosHilited` | Known | Event handler |
| 0x00236178 | `HandlePodcastsHilited` | Known | Event handler |
| 0x00236190 | `HandleGenericHilited` | Known | Event handler |
| 0x002361A8 | `HandlePhotosHilited` | Known | Event handler |
| 0x002361BC | `HandleNowPlayingHilited` | Known | Event handler |
| 0x002361D4 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x002361F0 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00236208 | `HandleArtistsHilited` | Known | Event handler |
| 0x00236220 | `HandleGenresHilited` | Known | Event handler |
| 0x00236234 | `HandleAlbumsHilited` | Known | Event handler |
| 0x00236248 | `HandleCompilationsHilited` | Known | Event handler |
| 0x0023641C | `HandleComposersHilited` | Known | Event handler |
| 0x00236434 | `HandleSongsHilited` | Known | Event handler |
| 0x00236448 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x00236460 | `HandleTVShowsHilited` | Known | Event handler |
| 0x00236478 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00236494 | `HandleMoviesHilited` | Known | Event handler |
| 0x002364A8 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x002364C4 | `HandleRentalsHilited` | Known | Event handler |
| 0x002364DC | `HandleMusicSelected` | Known | Event handler |
| 0x002364F0 | `HandleVideosSelected` | Known | Event handler |
| 0x00236508 | `HandlePodcastsSelected` | Known | Event handler |
| 0x002366D8 | `HandlePhotosSelected` | Known | Event handler |
| 0x002366F0 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00236708 | `HandleSongsSelected` | Known | Event handler |
| 0x0023671C | `HandleAlbumsSelected` | Known | Event handler |
| 0x00236734 | `HandleCompilationsSelected` | Known | Event handler |
| 0x00236750 | `HandleArtistsSelected` | Known | Event handler |
| 0x00236768 | `HandleGenresSelected` | Known | Event handler |
| 0x00236780 | `HandleComposersSelected` | Known | Event handler |
| 0x00236798 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x002367B4 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x002367D0 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x002369A8 | `HandleNowPlaying` | Known | Event handler |
| 0x002369BC | `HandleTVShowsSelected` | Known | Event handler |
| 0x002369D4 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x002369F0 | `HandleMoviesSelected` | Known | Event handler |
| 0x00236A08 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00236A28 | `HandleRentalsSelected` | Known | Event handler |
| 0x00236A40 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00236A58 | `HandleLock` | Known | Event handler |
| 0x00236A64 | `HandleBacklightSelected` | Known | Event handler |
| 0x00236A7C | `HandleSleepSelected` | Known | Event handler |
| 0x00236A90 | `HandleNikePlusSelected` | Known | Event handler |
| 0x00239354 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00239968 | `Handle400MetersWalk` | Known | Event handler |
| 0x00239980 | `HandleCustomWalk` | Known | Event handler |
| 0x00239994 | `HandleResetToDefault` | Known | Event handler |
| 0x00239C80 | `HandleSelect` | Known | Event handler |
| 0x00239F30 | `HandleWheel` | Known | Event handler |
| 0x0023B764 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x0023B9BC | `HandleNextDay` | Known | Event handler |
| 0x0023B9D0 | `HandlePreviousDay` | Known | Event handler |
| 0x0023BC18 | `HandleSelect` | Known | Event handler |
| 0x0023BEB4 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0023EB6C | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x0023EB88 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x0023FAF0 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x002401D0 | `HandleSelect` | Known | Event handler |
| 0x0024089C | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x0027DD48 | `HandleDeleteClock` | Known | Event handler |
| 0x0027DD60 | `HandleSelectClock` | Known | Event handler |
| 0x0027DD74 | `HandleHilited` | Known | Event handler |
| 0x0027DD84 | `HandleWheel` | Known | Event handler |
| 0x0027DD90 | `HandleSelectLozinch` | Known | Event handler |
| 0x004029AE | `HandleAudioFFDown` | Known | Event handler |
| 0x004029D7 | `HandleAudioFFUp` | Known | Event handler |
| 0x00402A02 | `HandleAudioMute` | Known | Event handler |
| 0x00402A35 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x00402A6A | `HandleAudioNext` | Known | Event handler |
| 0x00402A9A | `HandleAudioNextAlbum` | Known | Event handler |
| 0x00402AD1 | `HandleAudioNextChapter` | Known | Event handler |
| 0x00402B0B | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x00402B3F | `HandleAudioPause` | Known | Event handler |
| 0x00402B6B | `HandleAudioPlay` | Known | Event handler |
| 0x00402B99 | `HandleAudioPlayPause` | Known | Event handler |
| 0x00402BD1 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x00402C0A | `HandleAudioPrevious` | Known | Event handler |
| 0x00402C3E | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x00402C75 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x00402CAF | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x00402CE4 | `HandleAudioRepeat` | Known | Event handler |
| 0x00402D10 | `HandleAudioRewDown` | Known | Event handler |
| 0x00402D3B | `HandleAudioRewUp` | Known | Event handler |
| 0x00402D6A | `HandleAudioShuffle` | Known | Event handler |
| 0x00402D98 | `HandleAudioStop` | Known | Event handler |
| 0x00402DC9 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x00402DFE | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x00402E35 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x00402E66 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x00402F1F | `HandleNextPressAndHold` | Known | Event handler |
| 0x00402F50 | `HandleNext` | Known | Event handler |
| 0x00402F88 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x00402FC3 | `HandlePlayPause` | Known | Event handler |
| 0x00402FF7 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x0040302C | `HandlePrevious` | Known | Event handler |
| 0x004030B9 | `HandleRemoteBacklight` | Known | Event handler |
| 0x004030F1 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x0040312B | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x00403164 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x00403199 | `HandleRemoteEvent` | Known | Event handler |
| 0x004031C5 | `HandleRemoteFFDown` | Known | Event handler |
| 0x004031F0 | `HandleRemoteFFUp` | Known | Event handler |
| 0x0040321D | `HandleRemoteMenuDown` | Known | Event handler |
| 0x0040324C | `HandleRemoteMenuUp` | Known | Event handler |
| 0x0040327B | `HandleRemoteMute` | Known | Event handler |
| 0x004032AD | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x004032E6 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x00403322 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x00403362 | `HandleRemoteOff` | Known | Event handler |
| 0x0040338B | `HandleRemoteOff` | Known | Event handler |
| 0x004033B5 | `HandleRemoteOn` | Known | Event handler |
| 0x004033E1 | `HandleRemotePause` | Known | Event handler |
| 0x0040340F | `HandleRemotePlay` | Known | Event handler |
| 0x0040344D | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x0040348E | `HandleRemotePlayPause` | Known | Event handler |
| 0x004034C5 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x004034FE | `HandleRemotePrevChapter` | Known | Event handler |
| 0x0040353A | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x00403571 | `HandleRemoteRepeat` | Known | Event handler |
| 0x0040359F | `HandleRemoteRewDown` | Known | Event handler |
| 0x004035CC | `HandleRemoteRewUp` | Known | Event handler |
| 0x004035FC | `HandleRemoteSelectDown` | Known | Event handler |
| 0x0040362F | `HandleRemoteSelectUp` | Known | Event handler |
| 0x00403663 | `HandleRemoteShuffle` | Known | Event handler |
| 0x00403693 | `HandleRemoteStop` | Known | Event handler |
| 0x004036C3 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x004036F8 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x00403730 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x00403767 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x004037A0 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x004037D3 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x00403808 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x0040383B | `HandleVideoFFDown` | Known | Event handler |
| 0x00403864 | `HandleVideoFFUp` | Known | Event handler |
| 0x00403897 | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x004038CC | `HandleVideoNext` | Known | Event handler |
| 0x004038FE | `HandleVideoNextChapter` | Known | Event handler |
| 0x00403935 | `HandleVideoNextFrame` | Known | Event handler |
| 0x00403966 | `HandleVideoPause` | Known | Event handler |
| 0x00403992 | `HandleVideoPlay` | Known | Event handler |
| 0x004039C0 | `HandleVideoPlayPause` | Known | Event handler |
| 0x004039F8 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x00403A31 | `HandleVideoPrevious` | Known | Event handler |
| 0x00403A67 | `HandleVideoPrevChapter` | Known | Event handler |
| 0x00403A9E | `HandleVideoPrevFrame` | Known | Event handler |
| 0x00403ACD | `HandleVideoRewDown` | Known | Event handler |
| 0x00403AF8 | `HandleVideoRewUp` | Known | Event handler |
| 0x00403B24 | `HandleVideoStop` | Known | Event handler |
| 0x007539EA | `HandleAddressBook` | Known | Event handler |
| 0x00753F7E | `HandleSelect` | Known | Event handler |
| 0x00753FB9 | `HandleHilite` | Known | Event handler |
| 0x0075403A | `HandleSelectRegion` | Known | Event handler |
| 0x007540DA | `HandleSelectRegion` | Known | Event handler |
| 0x00754176 | `HandleSelectRegion` | Known | Event handler |
| 0x0075421A | `HandleSelectRegion` | Known | Event handler |
| 0x007542C0 | `HandleSelectRegion` | Known | Event handler |
| 0x00754360 | `HandleSelectRegion` | Known | Event handler |
| 0x0075440C | `HandleSelectRegion` | Known | Event handler |
| 0x007544AE | `HandleSelectRegion` | Known | Event handler |
| 0x0075455E | `HandleSelectCity` | Known | Event handler |
| 0x007545CA | `HandleHighlightCity` | Known | Event handler |
| 0x00754603 | `HandleSelectCity` | Known | Event handler |
| 0x0075466F | `HandleHighlightCity` | Known | Event handler |
| 0x007546A8 | `HandleSelectCity` | Known | Event handler |
| 0x00754714 | `HandleHighlightCity` | Known | Event handler |
| 0x0075474D | `HandleSelectCity` | Known | Event handler |
| 0x007547B9 | `HandleHighlightCity` | Known | Event handler |
| 0x007547F2 | `HandleSelectCity` | Known | Event handler |
| 0x0075485E | `HandleHighlightCity` | Known | Event handler |
| 0x00754897 | `HandleSelectCity` | Known | Event handler |
| 0x00754903 | `HandleHighlightCity` | Known | Event handler |
| 0x0075493C | `HandleSelectCity` | Known | Event handler |
| 0x007549A8 | `HandleHighlightCity` | Known | Event handler |
| 0x007549E1 | `HandleSelectCity` | Known | Event handler |
| 0x00754A4D | `HandleHighlightCity` | Known | Event handler |
| 0x00754A86 | `HandleSelectCity` | Known | Event handler |
| 0x00754AF2 | `HandleHighlightCity` | Known | Event handler |
| 0x00754B2B | `HandleSelectCity` | Known | Event handler |
| 0x00754B97 | `HandleHighlightCity` | Known | Event handler |
| 0x00754BD0 | `HandleSelectCity` | Known | Event handler |
| 0x00754C3C | `HandleHighlightCity` | Known | Event handler |
| 0x00754C75 | `HandleSelectCity` | Known | Event handler |
| 0x00754CE1 | `HandleHighlightCity` | Known | Event handler |
| 0x00754D1A | `HandleSelectCity` | Known | Event handler |
| 0x00754D86 | `HandleHighlightCity` | Known | Event handler |
| 0x00754DBF | `HandleSelectCity` | Known | Event handler |
| 0x00754E2B | `HandleHighlightCity` | Known | Event handler |
| 0x00754E64 | `HandleSelectCity` | Known | Event handler |
| 0x00754ED0 | `HandleHighlightCity` | Known | Event handler |
| 0x00754F09 | `HandleSelectCity` | Known | Event handler |
| 0x00754F75 | `HandleHighlightCity` | Known | Event handler |
| 0x00754FAE | `HandleSelectCity` | Known | Event handler |
| 0x0075501A | `HandleHighlightCity` | Known | Event handler |
| 0x00755053 | `HandleSelectCity` | Known | Event handler |
| 0x007550BF | `HandleHighlightCity` | Known | Event handler |
| 0x007550F8 | `HandleSelectCity` | Known | Event handler |
| 0x00755164 | `HandleHighlightCity` | Known | Event handler |
| 0x0075519D | `HandleSelectCity` | Known | Event handler |
| 0x00755209 | `HandleHighlightCity` | Known | Event handler |
| 0x00755242 | `HandleSelectCity` | Known | Event handler |
| 0x007552AE | `HandleHighlightCity` | Known | Event handler |
| 0x007552E7 | `HandleSelectCity` | Known | Event handler |
| 0x00755353 | `HandleHighlightCity` | Known | Event handler |
| 0x0075538C | `HandleSelectCity` | Known | Event handler |
| 0x007553F8 | `HandleHighlightCity` | Known | Event handler |
| 0x00755431 | `HandleSelectCity` | Known | Event handler |
| 0x0075549D | `HandleHighlightCity` | Known | Event handler |
| 0x007554D6 | `HandleSelectCity` | Known | Event handler |
| 0x00755542 | `HandleHighlightCity` | Known | Event handler |
| 0x0075557B | `HandleSelectCity` | Known | Event handler |
| 0x007555E7 | `HandleHighlightCity` | Known | Event handler |
| 0x00755620 | `HandleSelectCity` | Known | Event handler |
| 0x0075568C | `HandleHighlightCity` | Known | Event handler |
| 0x007556C5 | `HandleSelectCity` | Known | Event handler |
| 0x00755731 | `HandleHighlightCity` | Known | Event handler |
| 0x0075576A | `HandleSelectCity` | Known | Event handler |
| 0x007557D6 | `HandleHighlightCity` | Known | Event handler |
| 0x0075580F | `HandleSelectCity` | Known | Event handler |
| 0x0075587B | `HandleHighlightCity` | Known | Event handler |
| 0x007558B4 | `HandleSelectCity` | Known | Event handler |
| 0x00755920 | `HandleHighlightCity` | Known | Event handler |
| 0x0075595E | `HandleSelectCity` | Known | Event handler |
| 0x007559CA | `HandleHighlightCity` | Known | Event handler |
| 0x00755A03 | `HandleSelectCity` | Known | Event handler |
| 0x00755A6F | `HandleHighlightCity` | Known | Event handler |
| 0x00755AA8 | `HandleSelectCity` | Known | Event handler |
| 0x00755B14 | `HandleHighlightCity` | Known | Event handler |
| 0x00755B4D | `HandleSelectCity` | Known | Event handler |
| 0x00755BB9 | `HandleHighlightCity` | Known | Event handler |
| 0x00755BF2 | `HandleSelectCity` | Known | Event handler |
| 0x00755C5E | `HandleHighlightCity` | Known | Event handler |
| 0x00755C97 | `HandleSelectCity` | Known | Event handler |
| 0x00755D03 | `HandleHighlightCity` | Known | Event handler |
| 0x00755D3C | `HandleSelectCity` | Known | Event handler |
| 0x00755DA8 | `HandleHighlightCity` | Known | Event handler |
| 0x00755DE1 | `HandleSelectCity` | Known | Event handler |
| 0x00755E4D | `HandleHighlightCity` | Known | Event handler |
| 0x00755E86 | `HandleSelectCity` | Known | Event handler |
| 0x00755EF2 | `HandleHighlightCity` | Known | Event handler |
| 0x00755F2B | `HandleSelectCity` | Known | Event handler |
| 0x00755F97 | `HandleHighlightCity` | Known | Event handler |
| 0x00755FD0 | `HandleSelectCity` | Known | Event handler |
| 0x0075603C | `HandleHighlightCity` | Known | Event handler |
| 0x00756075 | `HandleSelectCity` | Known | Event handler |
| 0x007560E1 | `HandleHighlightCity` | Known | Event handler |
| 0x0075611A | `HandleSelectCity` | Known | Event handler |
| 0x00756186 | `HandleHighlightCity` | Known | Event handler |
| 0x007561BF | `HandleSelectCity` | Known | Event handler |
| 0x0075622B | `HandleHighlightCity` | Known | Event handler |
| 0x00756264 | `HandleSelectCity` | Known | Event handler |
| 0x007562D0 | `HandleHighlightCity` | Known | Event handler |
| 0x00756309 | `HandleSelectCity` | Known | Event handler |
| 0x00756375 | `HandleHighlightCity` | Known | Event handler |
| 0x007563AE | `HandleSelectCity` | Known | Event handler |
| 0x0075641A | `HandleHighlightCity` | Known | Event handler |
| 0x00756453 | `HandleSelectCity` | Known | Event handler |
| 0x007564BF | `HandleHighlightCity` | Known | Event handler |
| 0x007564F8 | `HandleSelectCity` | Known | Event handler |
| 0x00756564 | `HandleHighlightCity` | Known | Event handler |
| 0x0075659D | `HandleSelectCity` | Known | Event handler |
| 0x00756609 | `HandleHighlightCity` | Known | Event handler |
| 0x00756642 | `HandleSelectCity` | Known | Event handler |
| 0x007566AE | `HandleHighlightCity` | Known | Event handler |
| 0x007566E7 | `HandleSelectCity` | Known | Event handler |
| 0x00756753 | `HandleHighlightCity` | Known | Event handler |
| 0x0075678C | `HandleSelectCity` | Known | Event handler |
| 0x007567F8 | `HandleHighlightCity` | Known | Event handler |
| 0x00756831 | `HandleSelectCity` | Known | Event handler |
| 0x0075689D | `HandleHighlightCity` | Known | Event handler |
| 0x007568D6 | `HandleSelectCity` | Known | Event handler |
| 0x00756942 | `HandleHighlightCity` | Known | Event handler |
| 0x0075697B | `HandleSelectCity` | Known | Event handler |
| 0x007569E7 | `HandleHighlightCity` | Known | Event handler |
| 0x00756A20 | `HandleSelectCity` | Known | Event handler |
| 0x00756A8C | `HandleHighlightCity` | Known | Event handler |
| 0x00756AC5 | `HandleSelectCity` | Known | Event handler |
| 0x00756B31 | `HandleHighlightCity` | Known | Event handler |
| 0x00756B6A | `HandleSelectCity` | Known | Event handler |
| 0x00756BD6 | `HandleHighlightCity` | Known | Event handler |
| 0x00756C0F | `HandleSelectCity` | Known | Event handler |
| 0x00756C7B | `HandleHighlightCity` | Known | Event handler |
| 0x00756CB4 | `HandleSelectCity` | Known | Event handler |
| 0x00756D20 | `HandleHighlightCity` | Known | Event handler |
| 0x00756D59 | `HandleSelectCity` | Known | Event handler |
| 0x00756DC5 | `HandleHighlightCity` | Known | Event handler |
| 0x00756DFE | `HandleSelectCity` | Known | Event handler |
| 0x00756E6A | `HandleHighlightCity` | Known | Event handler |
| 0x00756EA3 | `HandleSelectCity` | Known | Event handler |
| 0x00756F0F | `HandleHighlightCity` | Known | Event handler |
| 0x00756F48 | `HandleSelectCity` | Known | Event handler |
| 0x00756FB4 | `HandleHighlightCity` | Known | Event handler |
| 0x00756FED | `HandleSelectCity` | Known | Event handler |
| 0x00757059 | `HandleHighlightCity` | Known | Event handler |
| 0x00757092 | `HandleSelectCity` | Known | Event handler |
| 0x007570FE | `HandleHighlightCity` | Known | Event handler |
| 0x00757137 | `HandleSelectCity` | Known | Event handler |
| 0x007571A3 | `HandleHighlightCity` | Known | Event handler |
| 0x007571DC | `HandleSelectCity` | Known | Event handler |
| 0x00757248 | `HandleHighlightCity` | Known | Event handler |
| 0x00757281 | `HandleSelectCity` | Known | Event handler |
| 0x007572ED | `HandleHighlightCity` | Known | Event handler |
| 0x00757326 | `HandleSelectCity` | Known | Event handler |
| 0x00757392 | `HandleHighlightCity` | Known | Event handler |
| 0x007573CB | `HandleSelectCity` | Known | Event handler |
| 0x00757437 | `HandleHighlightCity` | Known | Event handler |
| 0x00757470 | `HandleSelectCity` | Known | Event handler |
| 0x007574DC | `HandleHighlightCity` | Known | Event handler |
| 0x00757515 | `HandleSelectCity` | Known | Event handler |
| 0x00757581 | `HandleHighlightCity` | Known | Event handler |
| 0x007575BA | `HandleSelectCity` | Known | Event handler |
| 0x00757626 | `HandleHighlightCity` | Known | Event handler |
| 0x0075765F | `HandleSelectCity` | Known | Event handler |
| 0x007576CB | `HandleHighlightCity` | Known | Event handler |
| 0x00757704 | `HandleSelectCity` | Known | Event handler |
| 0x00757770 | `HandleHighlightCity` | Known | Event handler |
| 0x007577A9 | `HandleSelectCity` | Known | Event handler |
| 0x00757815 | `HandleHighlightCity` | Known | Event handler |
| 0x0075784E | `HandleSelectCity` | Known | Event handler |
| 0x007578BA | `HandleHighlightCity` | Known | Event handler |
| 0x007578F3 | `HandleSelectCity` | Known | Event handler |
| 0x0075795F | `HandleHighlightCity` | Known | Event handler |
| 0x00757998 | `HandleSelectCity` | Known | Event handler |
| 0x00757A04 | `HandleHighlightCity` | Known | Event handler |
| 0x00757A3D | `HandleSelectCity` | Known | Event handler |
| 0x00757AA9 | `HandleHighlightCity` | Known | Event handler |
| 0x00757AE2 | `HandleSelectCity` | Known | Event handler |
| 0x00757B4E | `HandleHighlightCity` | Known | Event handler |
| 0x00757B87 | `HandleSelectCity` | Known | Event handler |
| 0x00757BF3 | `HandleHighlightCity` | Known | Event handler |
| 0x00757C2C | `HandleSelectCity` | Known | Event handler |
| 0x00757C98 | `HandleHighlightCity` | Known | Event handler |
| 0x00757CD1 | `HandleSelectCity` | Known | Event handler |
| 0x00757D3D | `HandleHighlightCity` | Known | Event handler |
| 0x00757D76 | `HandleSelectCity` | Known | Event handler |
| 0x00757DE2 | `HandleHighlightCity` | Known | Event handler |
| 0x00757E22 | `HandleSelectCity` | Known | Event handler |
| 0x00757E8E | `HandleHighlightCity` | Known | Event handler |
| 0x00757EC7 | `HandleSelectCity` | Known | Event handler |
| 0x00757F33 | `HandleHighlightCity` | Known | Event handler |
| 0x00757F6C | `HandleSelectCity` | Known | Event handler |
| 0x00757FD8 | `HandleHighlightCity` | Known | Event handler |
| 0x00758016 | `HandleSelectCity` | Known | Event handler |
| 0x00758082 | `HandleHighlightCity` | Known | Event handler |
| 0x007580BB | `HandleSelectCity` | Known | Event handler |
| 0x00758127 | `HandleHighlightCity` | Known | Event handler |
| 0x00758160 | `HandleSelectCity` | Known | Event handler |
| 0x007581CC | `HandleHighlightCity` | Known | Event handler |
| 0x00758205 | `HandleSelectCity` | Known | Event handler |
| 0x00758271 | `HandleHighlightCity` | Known | Event handler |
| 0x007582AA | `HandleSelectCity` | Known | Event handler |
| 0x00758316 | `HandleHighlightCity` | Known | Event handler |
| 0x0075834F | `HandleSelectCity` | Known | Event handler |
| 0x007583BB | `HandleHighlightCity` | Known | Event handler |
| 0x007583F4 | `HandleSelectCity` | Known | Event handler |
| 0x00758460 | `HandleHighlightCity` | Known | Event handler |
| 0x00758499 | `HandleSelectCity` | Known | Event handler |
| 0x00758505 | `HandleHighlightCity` | Known | Event handler |
| 0x00758542 | `HandleSelectCity` | Known | Event handler |
| 0x007585AE | `HandleHighlightCity` | Known | Event handler |
| 0x007585E7 | `HandleSelectCity` | Known | Event handler |
| 0x00758653 | `HandleHighlightCity` | Known | Event handler |
| 0x0075868C | `HandleSelectCity` | Known | Event handler |
| 0x007586F8 | `HandleHighlightCity` | Known | Event handler |
| 0x00758731 | `HandleSelectCity` | Known | Event handler |
| 0x0075879D | `HandleHighlightCity` | Known | Event handler |
| 0x007587D6 | `HandleSelectCity` | Known | Event handler |
| 0x00758842 | `HandleHighlightCity` | Known | Event handler |
| 0x0075887B | `HandleSelectCity` | Known | Event handler |
| 0x007588E7 | `HandleHighlightCity` | Known | Event handler |
| 0x00758920 | `HandleSelectCity` | Known | Event handler |
| 0x0075898C | `HandleHighlightCity` | Known | Event handler |
| 0x007589C5 | `HandleSelectCity` | Known | Event handler |
| 0x00758A31 | `HandleHighlightCity` | Known | Event handler |
| 0x00758A6A | `HandleSelectCity` | Known | Event handler |
| 0x00758AD6 | `HandleHighlightCity` | Known | Event handler |
| 0x00758B0F | `HandleSelectCity` | Known | Event handler |
| 0x00758B7B | `HandleHighlightCity` | Known | Event handler |
| 0x00758BB4 | `HandleSelectCity` | Known | Event handler |
| 0x00758C20 | `HandleHighlightCity` | Known | Event handler |
| 0x00758C59 | `HandleSelectCity` | Known | Event handler |
| 0x00758CC5 | `HandleHighlightCity` | Known | Event handler |
| 0x00758CFE | `HandleSelectCity` | Known | Event handler |
| 0x00758D6A | `HandleHighlightCity` | Known | Event handler |
| 0x00758DA3 | `HandleSelectCity` | Known | Event handler |
| 0x00758E0F | `HandleHighlightCity` | Known | Event handler |
| 0x00758E48 | `HandleSelectCity` | Known | Event handler |
| 0x00758EB4 | `HandleHighlightCity` | Known | Event handler |
| 0x00758EED | `HandleSelectCity` | Known | Event handler |
| 0x00758F59 | `HandleHighlightCity` | Known | Event handler |
| 0x00758F92 | `HandleSelectCity` | Known | Event handler |
| 0x00758FFE | `HandleHighlightCity` | Known | Event handler |
| 0x00759037 | `HandleSelectCity` | Known | Event handler |
| 0x007590A3 | `HandleHighlightCity` | Known | Event handler |
| 0x007590DC | `HandleSelectCity` | Known | Event handler |
| 0x00759148 | `HandleHighlightCity` | Known | Event handler |
| 0x00759181 | `HandleSelectCity` | Known | Event handler |
| 0x007591ED | `HandleHighlightCity` | Known | Event handler |
| 0x00759226 | `HandleSelectCity` | Known | Event handler |
| 0x00759292 | `HandleHighlightCity` | Known | Event handler |
| 0x007592CB | `HandleSelectCity` | Known | Event handler |
| 0x00759337 | `HandleHighlightCity` | Known | Event handler |
| 0x00759370 | `HandleSelectCity` | Known | Event handler |
| 0x007593DC | `HandleHighlightCity` | Known | Event handler |
| 0x00759415 | `HandleSelectCity` | Known | Event handler |
| 0x00759481 | `HandleHighlightCity` | Known | Event handler |
| 0x007594BA | `HandleSelectCity` | Known | Event handler |
| 0x00759526 | `HandleHighlightCity` | Known | Event handler |
| 0x0075955F | `HandleSelectCity` | Known | Event handler |
| 0x007595CB | `HandleHighlightCity` | Known | Event handler |
| 0x00759604 | `HandleSelectCity` | Known | Event handler |
| 0x00759670 | `HandleHighlightCity` | Known | Event handler |
| 0x007596A9 | `HandleSelectCity` | Known | Event handler |
| 0x00759715 | `HandleHighlightCity` | Known | Event handler |
| 0x0075974E | `HandleSelectCity` | Known | Event handler |
| 0x007597BA | `HandleHighlightCity` | Known | Event handler |
| 0x007597F3 | `HandleSelectCity` | Known | Event handler |
| 0x0075985F | `HandleHighlightCity` | Known | Event handler |
| 0x00759898 | `HandleSelectCity` | Known | Event handler |
| 0x00759904 | `HandleHighlightCity` | Known | Event handler |
| 0x0075993D | `HandleSelectCity` | Known | Event handler |
| 0x007599A9 | `HandleHighlightCity` | Known | Event handler |
| 0x007599E2 | `HandleSelectCity` | Known | Event handler |
| 0x00759A4E | `HandleHighlightCity` | Known | Event handler |
| 0x00759A87 | `HandleSelectCity` | Known | Event handler |
| 0x00759AF3 | `HandleHighlightCity` | Known | Event handler |
| 0x00759B32 | `HandleSelectCity` | Known | Event handler |
| 0x00759B9E | `HandleHighlightCity` | Known | Event handler |
| 0x00759BD7 | `HandleSelectCity` | Known | Event handler |
| 0x00759C43 | `HandleHighlightCity` | Known | Event handler |
| 0x00759C7C | `HandleSelectCity` | Known | Event handler |
| 0x00759CE8 | `HandleHighlightCity` | Known | Event handler |
| 0x00759D21 | `HandleSelectCity` | Known | Event handler |
| 0x00759D8D | `HandleHighlightCity` | Known | Event handler |
| 0x00759DC6 | `HandleSelectCity` | Known | Event handler |
| 0x00759E32 | `HandleHighlightCity` | Known | Event handler |
| 0x00759E6B | `HandleSelectCity` | Known | Event handler |
| 0x00759ED7 | `HandleHighlightCity` | Known | Event handler |
| 0x00759F10 | `HandleSelectCity` | Known | Event handler |
| 0x00759F7C | `HandleHighlightCity` | Known | Event handler |
| 0x00759FB5 | `HandleSelectCity` | Known | Event handler |
| 0x0075A021 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A05A | `HandleSelectCity` | Known | Event handler |
| 0x0075A0C6 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A0FF | `HandleSelectCity` | Known | Event handler |
| 0x0075A16B | `HandleHighlightCity` | Known | Event handler |
| 0x0075A1A4 | `HandleSelectCity` | Known | Event handler |
| 0x0075A210 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A249 | `HandleSelectCity` | Known | Event handler |
| 0x0075A2B5 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A2EE | `HandleSelectCity` | Known | Event handler |
| 0x0075A35A | `HandleHighlightCity` | Known | Event handler |
| 0x0075A393 | `HandleSelectCity` | Known | Event handler |
| 0x0075A3FF | `HandleHighlightCity` | Known | Event handler |
| 0x0075A438 | `HandleSelectCity` | Known | Event handler |
| 0x0075A4A4 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A4DD | `HandleSelectCity` | Known | Event handler |
| 0x0075A549 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A582 | `HandleSelectCity` | Known | Event handler |
| 0x0075A5EE | `HandleHighlightCity` | Known | Event handler |
| 0x0075A627 | `HandleSelectCity` | Known | Event handler |
| 0x0075A693 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A6CC | `HandleSelectCity` | Known | Event handler |
| 0x0075A738 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A771 | `HandleSelectCity` | Known | Event handler |
| 0x0075A7DD | `HandleHighlightCity` | Known | Event handler |
| 0x0075A816 | `HandleSelectCity` | Known | Event handler |
| 0x0075A882 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A8BB | `HandleSelectCity` | Known | Event handler |
| 0x0075A927 | `HandleHighlightCity` | Known | Event handler |
| 0x0075A960 | `HandleSelectCity` | Known | Event handler |
| 0x0075A9CC | `HandleHighlightCity` | Known | Event handler |
| 0x0075AA05 | `HandleSelectCity` | Known | Event handler |
| 0x0075AA71 | `HandleHighlightCity` | Known | Event handler |
| 0x0075AAAA | `HandleSelectCity` | Known | Event handler |
| 0x0075AB16 | `HandleHighlightCity` | Known | Event handler |
| 0x0075AB4F | `HandleSelectCity` | Known | Event handler |
| 0x0075ABBB | `HandleHighlightCity` | Known | Event handler |
| 0x0075ABF4 | `HandleSelectCity` | Known | Event handler |
| 0x0075AC60 | `HandleHighlightCity` | Known | Event handler |
| 0x0075AC99 | `HandleSelectCity` | Known | Event handler |
| 0x0075AD05 | `HandleHighlightCity` | Known | Event handler |
| 0x0075AD3E | `HandleSelectCity` | Known | Event handler |
| 0x0075ADAA | `HandleHighlightCity` | Known | Event handler |
| 0x0075ADE3 | `HandleSelectCity` | Known | Event handler |
| 0x0075AE4F | `HandleHighlightCity` | Known | Event handler |
| 0x0075AE88 | `HandleSelectCity` | Known | Event handler |
| 0x0075AEF4 | `HandleHighlightCity` | Known | Event handler |
| 0x0075AF2D | `HandleSelectCity` | Known | Event handler |
| 0x0075AF99 | `HandleHighlightCity` | Known | Event handler |
| 0x0075AFD2 | `HandleSelectCity` | Known | Event handler |
| 0x0075B03E | `HandleHighlightCity` | Known | Event handler |
| 0x0075B077 | `HandleSelectCity` | Known | Event handler |
| 0x0075B0E3 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B11C | `HandleSelectCity` | Known | Event handler |
| 0x0075B188 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B1C1 | `HandleSelectCity` | Known | Event handler |
| 0x0075B22D | `HandleHighlightCity` | Known | Event handler |
| 0x0075B266 | `HandleSelectCity` | Known | Event handler |
| 0x0075B2D2 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B30B | `HandleSelectCity` | Known | Event handler |
| 0x0075B377 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B3B0 | `HandleSelectCity` | Known | Event handler |
| 0x0075B41C | `HandleHighlightCity` | Known | Event handler |
| 0x0075B455 | `HandleSelectCity` | Known | Event handler |
| 0x0075B4C1 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B4FA | `HandleSelectCity` | Known | Event handler |
| 0x0075B566 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B59F | `HandleSelectCity` | Known | Event handler |
| 0x0075B60B | `HandleHighlightCity` | Known | Event handler |
| 0x0075B644 | `HandleSelectCity` | Known | Event handler |
| 0x0075B6B0 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B6E9 | `HandleSelectCity` | Known | Event handler |
| 0x0075B755 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B78E | `HandleSelectCity` | Known | Event handler |
| 0x0075B7FA | `HandleHighlightCity` | Known | Event handler |
| 0x0075B833 | `HandleSelectCity` | Known | Event handler |
| 0x0075B89F | `HandleHighlightCity` | Known | Event handler |
| 0x0075B8D8 | `HandleSelectCity` | Known | Event handler |
| 0x0075B944 | `HandleHighlightCity` | Known | Event handler |
| 0x0075B97D | `HandleSelectCity` | Known | Event handler |
| 0x0075B9E9 | `HandleHighlightCity` | Known | Event handler |
| 0x0075BA22 | `HandleSelectCity` | Known | Event handler |
| 0x0075BA8E | `HandleHighlightCity` | Known | Event handler |
| 0x0075BAC7 | `HandleSelectCity` | Known | Event handler |
| 0x0075BB33 | `HandleHighlightCity` | Known | Event handler |
| 0x0075BB72 | `HandleSelectCity` | Known | Event handler |
| 0x0075BBDE | `HandleHighlightCity` | Known | Event handler |
| 0x0075BC17 | `HandleSelectCity` | Known | Event handler |
| 0x0075BC83 | `HandleHighlightCity` | Known | Event handler |
| 0x0075BCBC | `HandleSelectCity` | Known | Event handler |
| 0x0075BD28 | `HandleHighlightCity` | Known | Event handler |
| 0x0075BD61 | `HandleSelectCity` | Known | Event handler |
| 0x0075BDCD | `HandleHighlightCity` | Known | Event handler |
| 0x0075BE06 | `HandleSelectCity` | Known | Event handler |
| 0x0075BE72 | `HandleHighlightCity` | Known | Event handler |
| 0x0075BEB2 | `HandleSelectCity` | Known | Event handler |
| 0x0075BF1E | `HandleHighlightCity` | Known | Event handler |
| 0x0075BF57 | `HandleSelectCity` | Known | Event handler |
| 0x0075BFC3 | `HandleHighlightCity` | Known | Event handler |
| 0x0075BFFC | `HandleSelectCity` | Known | Event handler |
| 0x0075C068 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C0A1 | `HandleSelectCity` | Known | Event handler |
| 0x0075C10D | `HandleHighlightCity` | Known | Event handler |
| 0x0075C146 | `HandleSelectCity` | Known | Event handler |
| 0x0075C1B2 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C1EB | `HandleSelectCity` | Known | Event handler |
| 0x0075C257 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C290 | `HandleSelectCity` | Known | Event handler |
| 0x0075C2FC | `HandleHighlightCity` | Known | Event handler |
| 0x0075C335 | `HandleSelectCity` | Known | Event handler |
| 0x0075C3A1 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C3DA | `HandleSelectCity` | Known | Event handler |
| 0x0075C446 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C47F | `HandleSelectCity` | Known | Event handler |
| 0x0075C4EB | `HandleHighlightCity` | Known | Event handler |
| 0x0075C524 | `HandleSelectCity` | Known | Event handler |
| 0x0075C590 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C5C9 | `HandleSelectCity` | Known | Event handler |
| 0x0075C635 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C66E | `HandleSelectCity` | Known | Event handler |
| 0x0075C6DA | `HandleHighlightCity` | Known | Event handler |
| 0x0075C713 | `HandleSelectCity` | Known | Event handler |
| 0x0075C77F | `HandleHighlightCity` | Known | Event handler |
| 0x0075C7B8 | `HandleSelectCity` | Known | Event handler |
| 0x0075C824 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C85D | `HandleSelectCity` | Known | Event handler |
| 0x0075C8C9 | `HandleHighlightCity` | Known | Event handler |
| 0x0075C902 | `HandleSelectCity` | Known | Event handler |
| 0x0075C96E | `HandleHighlightCity` | Known | Event handler |
| 0x0075CE66 | `HandleMusicSelected` | Known | Event handler |
| 0x0075CEA8 | `HandleMusicHilited` | Known | Event handler |
| 0x0075CEE0 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x0075CF26 | `HandleMusicHilited` | Known | Event handler |
| 0x0075CF5E | `HandlePlaylistsSelected` | Known | Event handler |
| 0x0075CFA4 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x0075CFE0 | `HandleArtistsSelected` | Known | Event handler |
| 0x0075D024 | `HandleArtistsHilited` | Known | Event handler |
| 0x0075D05E | `HandleAlbumsSelected` | Known | Event handler |
| 0x0075D0A1 | `HandleAlbumsHilited` | Known | Event handler |
| 0x0075D0DA | `HandleCompilationsSelected` | Known | Event handler |
| 0x0075D123 | `HandleCompilationsHilited` | Known | Event handler |
| 0x0075D162 | `HandleSongsSelected` | Known | Event handler |
| 0x0075D1A4 | `HandleSongsHilited` | Known | Event handler |
| 0x0075D1DC | `HandleGenresSelected` | Known | Event handler |
| 0x0075D21F | `HandleGenresHilited` | Known | Event handler |
| 0x0075D258 | `HandleComposersSelected` | Known | Event handler |
| 0x0075D29E | `HandleComposersHilited` | Known | Event handler |
| 0x0075D2DA | `HandleAudiobooksSelected` | Known | Event handler |
| 0x0075D321 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x0075D3E0 | `HandleMusicHilited` | Known | Event handler |
| 0x0075D418 | `HandleVideosSelected` | Known | Event handler |
| 0x0075D45B | `HandleVideosHilited` | Known | Event handler |
| 0x0075D494 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x0075D4DF | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x0075D520 | `HandleMoviesSelected` | Known | Event handler |
| 0x0075D563 | `HandleMoviesHilited` | Known | Event handler |
| 0x0075D59C | `HandleTVShowsSelected` | Known | Event handler |
| 0x0075D5E0 | `HandleTVShowsHilited` | Known | Event handler |
| 0x0075D61A | `HandleMusicVideosSelected` | Known | Event handler |
| 0x0075D662 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x0075D6A0 | `HandleRentalsSelected` | Known | Event handler |
| 0x0075D6E4 | `HandleRentalsHilited` | Known | Event handler |
| 0x0075D71E | `HandlePhotosSelected` | Known | Event handler |
| 0x0075D761 | `HandlePhotosHilited` | Known | Event handler |
| 0x0075D79A | `HandlePhotosSelected` | Known | Event handler |
| 0x0075D7DD | `HandlePhotosHilited` | Known | Event handler |
| 0x0075D816 | `HandlePodcastsSelected` | Known | Event handler |
| 0x0075D85B | `HandlePodcastsHilited` | Known | Event handler |
| 0x0075D90E | `HandleGenericHilited` | Known | Event handler |
| 0x0075DA07 | `HandleGenericHilited` | Known | Event handler |
| 0x0075DEEC | `HandleLock` | Known | Event handler |
| 0x0075E05D | `HandleNikePlusSelected` | Known | Event handler |
| 0x0075E0A2 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E1A8 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E2A7 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E394 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E491 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E50B | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x0075E554 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E5CD | `HandleBacklightSelected` | Known | Event handler |
| 0x0075E613 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E68E | `HandleSleepSelected` | Known | Event handler |
| 0x0075E6D0 | `HandleGenericHilited` | Known | Event handler |
| 0x0075E747 | `HandleNowPlaying` | Known | Event handler |
| 0x0075E7BF | `HandleNowPlayingHilited` | Known | Event handler |
| 0x0075E802 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x0075E848 | `HandleMusicHilited` | Known | Event handler |
| 0x0075E880 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x0075E8C6 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x0075E904 | `HandleArtistsSelected` | Known | Event handler |
| 0x0075E948 | `HandleArtistsHilited` | Known | Event handler |
| 0x0075E982 | `HandleAlbumsSelected` | Known | Event handler |
| 0x0075E9C5 | `HandleAlbumsHilited` | Known | Event handler |
| 0x0075E9FE | `HandleCompilationsSelected` | Known | Event handler |
| 0x0075EA47 | `HandleCompilationsHilited` | Known | Event handler |
| 0x0075EA86 | `HandleSongsSelected` | Known | Event handler |
| 0x0075EAC8 | `HandleSongsHilited` | Known | Event handler |
| 0x0075EB73 | `HandleGenericHilited` | Known | Event handler |
| 0x0075EBEB | `HandleGenresSelected` | Known | Event handler |
| 0x0075EC2E | `HandleGenresHilited` | Known | Event handler |
| 0x0075EC67 | `HandleComposersSelected` | Known | Event handler |
| 0x0075ECAD | `HandleComposersHilited` | Known | Event handler |
| 0x0075ECE9 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x0075ED30 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x0075EDEF | `HandleMusicHilited` | Known | Event handler |
| 0x0075EE65 | `HandlePlayPause` | Known | Event handler |
| 0x0075EE9A | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x0075EF84 | `HandleSelect` | Known | Event handler |
| 0x0075EFCA | `HandleMoviesSelected` | Known | Event handler |
| 0x0075F00D | `HandleMoviesHilited` | Known | Event handler |
| 0x0075F046 | `HandleRentalsSelected` | Known | Event handler |
| 0x0075F08A | `HandleRentalsHilited` | Known | Event handler |
| 0x0075F0C4 | `HandleTVShowsSelected` | Known | Event handler |
| 0x0075F108 | `HandleTVShowsHilited` | Known | Event handler |
| 0x0075F142 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x0075F18A | `HandleMusicVideosHilited` | Known | Event handler |
| 0x0075F1C8 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x0075F213 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x0075F2D9 | `HandleVideosHilited` | Known | Event handler |
| 0x0075F91B | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x007604A2 | `HandleMainMenu` | Known | Event handler |
| 0x007604DA | `HandleMusicMenu` | Known | Event handler |
| 0x00760A02 | `HandleRadioRegion` | Known | Event handler |
| 0x00760AA6 | `HandleLanguage` | Known | Event handler |
| 0x00760BAC | `HandleNew` | Known | Event handler |
| 0x00760C27 | `HandleClear` | Known | Event handler |
| 0x00760C58 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x00760D14 | `HandleSelectIndexedSession` | Known | Event handler |
| 0x00760EDB | `HandleBasicSelected` | Known | Event handler |
| 0x00760F81 | `HandleTimedWorkoutSelected` | Known | Event handler |
| 0x0076102E | `HandleDistanceWorkoutSelected` | Known | Event handler |
| 0x007610DE | `HandleCaloriesWorkoutSelected` | Known | Event handler |
| 0x007614C9 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x0076151C | `HandleSelect` | Known | Event handler |
| 0x00761646 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x00761680 | `HandleEQSettingSelected` | Known | Event handler |
| 0x007616B8 | `HandleEQSettingSelected` | Known | Event handler |
| 0x0077681C | `HandleItemSelected` | Known | Event handler |
| 0x00776967 | `HandleNextContact` | Known | Event handler |
| 0x00776993 | `HandlePreviousContact` | Known | Event handler |
| 0x007769C9 | `HandleSelectKey` | Known | Event handler |
| 0x00776FDA | `HandleSelect` | Known | Event handler |
| 0x00777301 | `HandleDateChosen` | Known | Event handler |
| 0x00777337 | `HandleTimeChosen` | Known | Event handler |
| 0x0077736D | `HandleFrequencyChosen` | Known | Event handler |
| 0x007773A8 | `HandleSoundChosen` | Known | Event handler |
| 0x007773DF | `HandleLabelChosen` | Known | Event handler |
| 0x00777416 | `HandleDeleteChosen` | Known | Event handler |
| 0x00777452 | `HandleSelect` | Known | Event handler |
| 0x0077748A | `HandleSelect` | Known | Event handler |
| 0x007777CB | `HandleLeaveAlarm` | Known | Event handler |
| 0x007777F8 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00777827 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00777854 | `HandleLeaveAlarm` | Known | Event handler |
| 0x0077798E | `HandleSelect` | Known | Event handler |
| 0x007779BC | `HandleSelect` | Known | Event handler |
| 0x00777B1B | `HandleNextDay` | Known | Event handler |
| 0x00777B43 | `HandlePreviousDay` | Known | Event handler |
| 0x00777CF2 | `HandleSelect` | Known | Event handler |
| 0x00777D1F | `HandleNextDay` | Known | Event handler |
| 0x00777D47 | `HandlePreviousDay` | Known | Event handler |
| 0x00777EEF | `HandleNextDay` | Known | Event handler |
| 0x00777F17 | `HandlePreviousDay` | Known | Event handler |
| 0x00777FD8 | `HandleSelect` | Known | Event handler |
| 0x00778003 | `HandleNextDay` | Known | Event handler |
| 0x0077802B | `HandlePreviousDay` | Known | Event handler |
| 0x007781A2 | `HandleSelectLozinch` | Known | Event handler |
| 0x0077831A | `HandleSelectLozinch` | Known | Event handler |
| 0x00778439 | `HandleFlowNext` | Known | Event handler |
| 0x00778467 | `HandlePlayPause` | Known | Event handler |
| 0x007784B6 | `HandleFlowPrev` | Known | Event handler |
| 0x007784E1 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x007785D5 | `HandleAlbumSelected` | Known | Event handler |
| 0x00778770 | `HandleFlowNext` | Known | Event handler |
| 0x007787BE | `HandleFlowNext` | Known | Event handler |
| 0x007787EC | `HandlePlayPause` | Known | Event handler |
| 0x0077883B | `HandleFlowPrev` | Known | Event handler |
| 0x00778867 | `HandleFlowPrev` | Known | Event handler |
| 0x00778887 | `HandleFlowWheel` | Known | Event handler |
| 0x00778C17 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x00779042 | `HandleArrowDown` | Known | Event handler |
| 0x007790AC | `HandleArrowUp` | Known | Event handler |
| 0x007790CB | `HandleWheel` | Known | Event handler |
| 0x00779154 | `HandleSelect` | Known | Event handler |
| 0x007791D1 | `HandleGameHilited` | Known | Event handler |
| 0x0077C637 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077E273 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077FEAF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00781AEB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00783727 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00785363 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00786F9F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00788BDB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078A817 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078C453 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078E08F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078FCCB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00791907 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00793543 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079517F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00796DBB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007989F7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079A633 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079C26F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079DEAB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079FAE7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A1723 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A335F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A4F9B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A6BD7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A8813 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007AA44F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007AC08B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007ADCC7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007AF903 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B153F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B317B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B4DB7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B69F3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B862F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BA26B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BBEA7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BDAC8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BE650 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BF1D8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BFD60 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C08E8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C1470 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C1FF8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C2B80 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C3708 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C4290 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C4E18 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C59A0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007C6528 | `HandlePlayPause` | Known | Event handler |
| 0x007C655E | `HandleAddToOTG` | Known | Event handler |
| 0x007C66FB | `HandlePlayPause` | Known | Event handler |
| 0x007C6722 | `HandleSelect` | Known | Event handler |
| 0x007C674F | `HandleHilite` | Known | Event handler |
| 0x007C6780 | `HandlePlayPause` | Known | Event handler |
| 0x007C6813 | `HandlePlayPause` | Known | Event handler |
| 0x007C683A | `HandleSelect` | Known | Event handler |
| 0x007C68A0 | `HandleHilite` | Known | Event handler |
| 0x007C68D2 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x007C691C | `HandlePlayPause` | Known | Event handler |
| 0x007C6952 | `HandleAddToOTG` | Known | Event handler |
| 0x007C69E4 | `HandlePlayPause` | Known | Event handler |
| 0x007C6A0B | `HandleSelect` | Known | Event handler |
| 0x007C6A74 | `HandlePlayPause` | Known | Event handler |
| 0x007C6AAA | `HandleAddToOTG` | Known | Event handler |
| 0x007C6B3C | `HandlePlayPause` | Known | Event handler |
| 0x007C6B63 | `HandleSelect` | Known | Event handler |
| 0x007C6BCC | `HandlePlayPause` | Known | Event handler |
| 0x007C6C52 | `HandleSelect` | Known | Event handler |
| 0x007C6CB7 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C6CF8 | `HandlePlayPause` | Known | Event handler |
| 0x007C6D2E | `HandleAddToOTG` | Known | Event handler |
| 0x007C6F60 | `HandlePlayPause` | Known | Event handler |
| 0x007C6F87 | `HandleSelect` | Known | Event handler |
| 0x007C6FB4 | `HandleHilite` | Known | Event handler |
| 0x007C6FE4 | `HandlePlayPause` | Known | Event handler |
| 0x007C701A | `HandleAddToOTG` | Known | Event handler |
| 0x007C724C | `HandlePlayPause` | Known | Event handler |
| 0x007C7273 | `HandleSelect` | Known | Event handler |
| 0x007C72A0 | `HandleHilite` | Known | Event handler |
| 0x007C72D0 | `HandlePlayPause` | Known | Event handler |
| 0x007C7306 | `HandleAddToOTG` | Known | Event handler |
| 0x007C75F1 | `HandlePlayPause` | Known | Event handler |
| 0x007C7618 | `HandleSelect` | Known | Event handler |
| 0x007C7648 | `HandlePlayPause` | Known | Event handler |
| 0x007C767E | `HandleAddToOTG` | Known | Event handler |
| 0x007C7710 | `HandlePlayPause` | Known | Event handler |
| 0x007C7737 | `HandleSelect` | Known | Event handler |
| 0x007C77C8 | `HandlePlayPause` | Known | Event handler |
| 0x007C77FE | `HandleAddToOTG` | Known | Event handler |
| 0x007C79B7 | `HandlePlayPause` | Known | Event handler |
| 0x007C79DE | `HandleSelect` | Known | Event handler |
| 0x007C7A10 | `HandlePlayPause` | Known | Event handler |
| 0x007C7A46 | `HandleAddToOTG` | Known | Event handler |
| 0x007C7ACB | `HandleSelect` | Known | Event handler |
| 0x007C7B64 | `HandleHilite` | Known | Event handler |
| 0x007C7B90 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C7BD4 | `HandlePlayPause` | Known | Event handler |
| 0x007C7C0A | `HandleAddToOTG` | Known | Event handler |
| 0x007C7C8F | `HandleSelect` | Known | Event handler |
| 0x007C7CF4 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C7D38 | `HandlePlayPause` | Known | Event handler |
| 0x007C7EDC | `HandleSelect` | Known | Event handler |
| 0x007C7F09 | `HandleHilite` | Known | Event handler |
| 0x007C7F35 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C7F78 | `HandlePlayPause` | Known | Event handler |
| 0x007C7FFE | `HandleSelect` | Known | Event handler |
| 0x007C808C | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C80D0 | `HandlePlayPause` | Known | Event handler |
| 0x007C8156 | `HandleSelect` | Known | Event handler |
| 0x007C81BB | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C81FC | `HandlePlayPause` | Known | Event handler |
| 0x007C8282 | `HandleSelect` | Known | Event handler |
| 0x007C82E8 | `HandleHilite` | Known | Event handler |
| 0x007C8314 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C8358 | `HandlePlayPause` | Known | Event handler |
| 0x007C838E | `HandleAddToOTG` | Known | Event handler |
| 0x007C8551 | `HandlePlayPause` | Known | Event handler |
| 0x007C8578 | `HandleSelect` | Known | Event handler |
| 0x007C85A8 | `HandlePlayPause` | Known | Event handler |
| 0x007C85DE | `HandleAddToOTG` | Known | Event handler |
| 0x007C87FF | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x007C8918 | `HandlePlayPause` | Known | Event handler |
| 0x007C8A45 | `HandleSelect` | Known | Event handler |
| 0x007C8A71 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C8AB4 | `HandlePlayPause` | Known | Event handler |
| 0x007C8B3A | `HandleSelect` | Known | Event handler |
| 0x007C8B67 | `HandleHilite` | Known | Event handler |
| 0x007C8B93 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C8BD4 | `HandlePlayPause` | Known | Event handler |
| 0x007C8D07 | `HandleSelect` | Known | Event handler |
| 0x007C8D33 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C9645 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C9EFD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CA7B5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CB06D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CB925 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CC1DD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CCA95 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CD34D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007CD396 | `HandleTVOutChanged` | Known | Event handler |
| 0x007CD3CE | `HandleTVSignalChanged` | Known | Event handler |
| 0x007CD409 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x007CD45A | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x007CD49F | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x007CD4E8 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x007CD52A | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x007CD56D | `HandleSelect` | Known | Event handler |
| 0x007CD59D | `HandleSelect` | Known | Event handler |
| 0x007CD5D5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007CD603 | `HandleMenuKey` | Known | Event handler |
| 0x007CD689 | `HandlePlayPause` | Known | Event handler |
| 0x007CD709 | `HandleSelect` | Known | Event handler |
| 0x007CE016 | `HandlePlayPause` | Known | Event handler |
| 0x007CE08B | `HandleWheelProgress` | Known | Event handler |
| 0x007CE0C9 | `HandleMenuLongpress` | Known | Event handler |
| 0x007CE0F7 | `HandleMenuKey` | Known | Event handler |
| 0x007CE17D | `HandlePlayPause` | Known | Event handler |
| 0x007CE1FD | `HandleSelectProgress` | Known | Event handler |
| 0x007CEB12 | `HandlePlayPause` | Known | Event handler |
| 0x007CEB87 | `HandleWheelProgress` | Known | Event handler |
| 0x007CEBC5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007CEBF3 | `HandleMenuKey` | Known | Event handler |
| 0x007CEC79 | `HandlePlayPause` | Known | Event handler |
| 0x007CECF9 | `HandleSelectVolume` | Known | Event handler |
| 0x007CF60C | `HandlePlayPause` | Known | Event handler |
| 0x007CF681 | `HandleWheelVolume` | Known | Event handler |
| 0x007CF6BD | `HandleMenuLongpress` | Known | Event handler |
| 0x007CF6EB | `HandleMenuKey` | Known | Event handler |
| 0x007CF771 | `HandlePlayPause` | Known | Event handler |
| 0x007CF7F1 | `HandleSelectRating` | Known | Event handler |
| 0x007D0104 | `HandlePlayPause` | Known | Event handler |
| 0x007D0179 | `HandleWheelRating` | Known | Event handler |
| 0x007D01B5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D01E3 | `HandleMenuKey` | Known | Event handler |
| 0x007D025B | `HandlePlayPause` | Known | Event handler |
| 0x007D02D2 | `HandleSelectScrub` | Known | Event handler |
| 0x007D0BD6 | `HandlePlayPause` | Known | Event handler |
| 0x007D0C42 | `HandleWheelScrub` | Known | Event handler |
| 0x007D0C7D | `HandleMenuLongpress` | Known | Event handler |
| 0x007D0CAB | `HandleMenuKey` | Known | Event handler |
| 0x007D0D08 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007D0D40 | `HandlePlayPause` | Known | Event handler |
| 0x007D0D9A | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007D0DCF | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x007D16E9 | `HandlePlayPause` | Known | Event handler |
| 0x007D175E | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007D17A1 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D17CF | `HandleMenuKey` | Known | Event handler |
| 0x007D1855 | `HandlePlayPause` | Known | Event handler |
| 0x007D18D5 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007D21EB | `HandlePlayPause` | Known | Event handler |
| 0x007D2289 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D22B7 | `HandleMenuKey` | Known | Event handler |
| 0x007D233D | `HandlePlayPause` | Known | Event handler |
| 0x007D23BD | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007D2CD3 | `HandlePlayPause` | Known | Event handler |
| 0x007D2D71 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D2D9F | `HandleMenuKey` | Known | Event handler |
| 0x007D2E25 | `HandlePlayPause` | Known | Event handler |
| 0x007D2EA5 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007D37BB | `HandlePlayPause` | Known | Event handler |
| 0x007D3859 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D3887 | `HandleMenuKey` | Known | Event handler |
| 0x007D390D | `HandlePlayPause` | Known | Event handler |
| 0x007D398D | `HandleSelectChapterArt` | Known | Event handler |
| 0x007D42A4 | `HandlePlayPause` | Known | Event handler |
| 0x007D4319 | `HandleWheelVolume` | Known | Event handler |
| 0x007D4355 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D4383 | `HandleMenuKey` | Known | Event handler |
| 0x007D4412 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007D44A9 | `HandleSelect` | Known | Event handler |
| 0x007D4DBF | `HandlePlayPause` | Known | Event handler |
| 0x007D4E3D | `HandleWheel` | Known | Event handler |
| 0x007D4E71 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D4E9F | `HandleMenuKey` | Known | Event handler |
| 0x007D4F2E | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007D4FC5 | `HandleSelect` | Known | Event handler |
| 0x007D58DB | `HandlePlayPause` | Known | Event handler |
| 0x007D5959 | `HandleWheel` | Known | Event handler |
| 0x007D598D | `HandleMenuLongpress` | Known | Event handler |
| 0x007D59BB | `HandleMenuKey` | Known | Event handler |
| 0x007D5A41 | `HandlePlayPause` | Known | Event handler |
| 0x007D5AC1 | `HandleSelect` | Known | Event handler |
| 0x007D63CE | `HandlePlayPause` | Known | Event handler |
| 0x007D6443 | `HandleWheel` | Known | Event handler |
| 0x007D6479 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D64A7 | `HandleMenuKey` | Known | Event handler |
| 0x007D652D | `HandlePlayPause` | Known | Event handler |
| 0x007D65AD | `HandleSelectProgress` | Known | Event handler |
| 0x007D6EC2 | `HandlePlayPause` | Known | Event handler |
| 0x007D6F37 | `HandleWheelProgress` | Known | Event handler |
| 0x007D6F75 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D6FA3 | `HandleMenuKey` | Known | Event handler |
| 0x007D701B | `HandlePlayPause` | Known | Event handler |
| 0x007D7092 | `HandleSelectScrub` | Known | Event handler |
| 0x007D7996 | `HandlePlayPause` | Known | Event handler |
| 0x007D7A02 | `HandleWheelScrub` | Known | Event handler |
| 0x007D7A3D | `HandleMenuLongpress` | Known | Event handler |
| 0x007D7A6B | `HandleMenuKey` | Known | Event handler |
| 0x007D7AF1 | `HandlePlayPause` | Known | Event handler |
| 0x007D847D | `HandlePlayPause` | Known | Event handler |
| 0x007D84F2 | `HandleWheelVolume` | Known | Event handler |
| 0x007D852D | `HandleMenuLongpress` | Known | Event handler |
| 0x007D855B | `HandleMenuKey` | Known | Event handler |
| 0x007D85E1 | `HandlePlayPause` | Known | Event handler |
| 0x007D8F6D | `HandlePlayPause` | Known | Event handler |
| 0x007D8FE2 | `HandleWheelBrightness` | Known | Event handler |
| 0x007D90F9 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007D9A4C | `HandleWheel` | Known | Event handler |
| 0x007D9A81 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D9AAF | `HandleMenuKey` | Known | Event handler |
| 0x007D9B35 | `HandlePlayPause` | Known | Event handler |
| 0x007D9BB5 | `HandleSelect` | Known | Event handler |
| 0x007DA057 | `HandlePlayPause` | Known | Event handler |
| 0x007DA0E5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007DA113 | `HandleMenuKey` | Known | Event handler |
| 0x007DA199 | `HandlePlayPause` | Known | Event handler |
| 0x007DA219 | `HandleSelectProgress` | Known | Event handler |
| 0x007DA6C3 | `HandlePlayPause` | Known | Event handler |
| 0x007DA738 | `HandleWheelProgress` | Known | Event handler |
| 0x007DA775 | `HandleMenuLongpress` | Known | Event handler |
| 0x007DA7A3 | `HandleMenuKey` | Known | Event handler |
| 0x007DA829 | `HandlePlayPause` | Known | Event handler |
| 0x007DA8A9 | `HandleSelectProgress` | Known | Event handler |
| 0x007DAD53 | `HandlePlayPause` | Known | Event handler |
| 0x007DADC8 | `HandleWheelProgress` | Known | Event handler |
| 0x007DAE05 | `HandleMenuLongpress` | Known | Event handler |
| 0x007DAE33 | `HandleMenuKey` | Known | Event handler |
| 0x007DAEB9 | `HandlePlayPause` | Known | Event handler |
| 0x007DAF39 | `HandleSelectProgress` | Known | Event handler |
| 0x007DB36F | `HandlePlayPause` | Known | Event handler |
| 0x007DB3E4 | `HandleWheelProgress` | Known | Event handler |
| 0x007DB421 | `HandleMenuLongpress` | Known | Event handler |
| 0x007DB44F | `HandleMenuKey` | Known | Event handler |
| 0x007DB4BC | `HandlePlayPause` | Known | Event handler |
| 0x007DB528 | `HandleSelectScrub` | Known | Event handler |
| 0x007DB942 | `HandlePlayPause` | Known | Event handler |
| 0x007DB9A3 | `HandleWheelScrub` | Known | Event handler |
| 0x007DB9DD | `HandleMenuLongpress` | Known | Event handler |
| 0x007DBA0B | `HandleMenuKey` | Known | Event handler |
| 0x007DBA91 | `HandlePlayPause` | Known | Event handler |
| 0x007DBB11 | `HandleSelectVolume` | Known | Event handler |
| 0x007DBF45 | `HandlePlayPause` | Known | Event handler |
| 0x007DBFBA | `HandleWheelVolume` | Known | Event handler |
| 0x007DC0CD | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007DC56C | `HandleSelect` | Known | Event handler |
| 0x007DC599 | `HandleSelect` | Known | Event handler |
| 0x007DC5C9 | `HandleSelect` | Known | Event handler |
| 0x007DC5F9 | `HandleSelect` | Known | Event handler |
| 0x007DC629 | `HandleSelect` | Known | Event handler |
| 0x007DC659 | `HandleSelect` | Known | Event handler |
| 0x007DC689 | `HandleSelect` | Known | Event handler |
| 0x007DC6B9 | `HandleSelect` | Known | Event handler |
| 0x007DC6E9 | `HandleSelect` | Known | Event handler |
| 0x007DC759 | `HandleSelect` | Known | Event handler |
| 0x007DC789 | `HandleSelect` | Known | Event handler |
| 0x007DC801 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DC834 | `HandleNotesPop` | Known | Event handler |
| 0x007DC8B1 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DC8E4 | `HandleNotesPop` | Known | Event handler |
| 0x007DCDA0 | `HandleNotesSelected` | Known | Event handler |
| 0x007DCDDD | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DCE10 | `HandleNotesPop` | Known | Event handler |
| 0x007DD2CC | `HandleNotesSelected` | Known | Event handler |
| 0x007DD309 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DD33C | `HandleNotesPop` | Known | Event handler |
| 0x007DD367 | `HandleNotesSelected` | Known | Event handler |
| 0x007DD839 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DD86C | `HandleNotesPop` | Known | Event handler |
| 0x007DD897 | `HandleNotesSelected` | Known | Event handler |
| 0x007DDD69 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DDD9C | `HandleNotesPop` | Known | Event handler |
| 0x007DDE19 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DDE4C | `HandleNotesPop` | Known | Event handler |
| 0x007DDEC9 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007DDEFC | `HandleNotesPop` | Known | Event handler |
| 0x007DDF74 | `HandlePlayPause` | Known | Event handler |
| 0x007DDF9D | `HandlePlayPause` | Known | Event handler |
| 0x007DDFCB | `HandlePlayPause` | Known | Event handler |
| 0x007DE000 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007DE080 | `HandleHiliteAlbum` | Known | Event handler |
| 0x007DE129 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007DE1B0 | `HandleHiliteAlbum` | Known | Event handler |
| 0x007DE474 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x007DE4D0 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x007DE687 | `HandleSelect` | Known | Event handler |
| 0x007DE80B | `HandleSelect` | Known | Event handler |
| 0x007DE845 | `HandleImageLast` | Known | Event handler |
| 0x007DE86F | `HandleImageNext` | Known | Event handler |
| 0x007DE89E | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007DE8D8 | `HandleImageFirst` | Known | Event handler |
| 0x007DE903 | `HandleImagePrev` | Known | Event handler |
| 0x007DE92F | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007DE95E | `HandleImageNext` | Known | Event handler |
| 0x007DE987 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007DE9BB | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007DE9EA | `HandleImagePrev` | Known | Event handler |
| 0x007DEA0B | `HandleImageWheel` | Known | Event handler |
| 0x007DEAA9 | `HandleImageNext` | Known | Event handler |
| 0x007DEAD8 | `HandlePlayPause` | Known | Event handler |
| 0x007DEB27 | `HandleImagePrev` | Known | Event handler |
| 0x007DEB53 | `HandleSelect` | Known | Event handler |
| 0x007DEE23 | `HandleImageNext` | Known | Event handler |
| 0x007DEE4D | `HandlePause` | Known | Event handler |
| 0x007DEE72 | `HandlePlay` | Known | Event handler |
| 0x007DEE9B | `HandlePlayPause` | Known | Event handler |
| 0x007DEEC4 | `HandleImagePrev` | Known | Event handler |
| 0x007DEF1D | `HandleWheel` | Known | Event handler |
| 0x007DEFB5 | `HandleImageNext` | Known | Event handler |
| 0x007DEFE4 | `HandlePlayPause` | Known | Event handler |
| 0x007DF033 | `HandleImagePrev` | Known | Event handler |
| 0x007DF05F | `HandleSelect` | Known | Event handler |
| 0x007DF32F | `HandleImageNext` | Known | Event handler |
| 0x007DF359 | `HandlePause` | Known | Event handler |
| 0x007DF37E | `HandlePlay` | Known | Event handler |
| 0x007DF3A7 | `HandlePlayPause` | Known | Event handler |
| 0x007DF3D0 | `HandleImagePrev` | Known | Event handler |
| 0x007DF429 | `HandleWheel` | Known | Event handler |
| 0x007DF4C1 | `HandleImageNext` | Known | Event handler |
| 0x007DF4F0 | `HandlePlayPause` | Known | Event handler |
| 0x007DF53F | `HandleImagePrev` | Known | Event handler |
| 0x007DF56B | `HandleSelect` | Known | Event handler |
| 0x007DF83B | `HandleImageNext` | Known | Event handler |
| 0x007DF865 | `HandlePause` | Known | Event handler |
| 0x007DF88A | `HandlePlay` | Known | Event handler |
| 0x007DF8B3 | `HandlePlayPause` | Known | Event handler |
| 0x007DF8DC | `HandleImagePrev` | Known | Event handler |
| 0x007DF935 | `HandleWheel` | Known | Event handler |
| 0x007DF9CD | `HandleImageNext` | Known | Event handler |
| 0x007DF9FC | `HandlePlayPause` | Known | Event handler |
| 0x007DFA4B | `HandleImagePrev` | Known | Event handler |
| 0x007DFA77 | `HandleSelect` | Known | Event handler |
| 0x007DFD47 | `HandleImageNext` | Known | Event handler |
| 0x007DFD71 | `HandlePause` | Known | Event handler |
| 0x007DFD96 | `HandlePlay` | Known | Event handler |
| 0x007DFDBF | `HandlePlayPause` | Known | Event handler |
| 0x007DFDE8 | `HandleImagePrev` | Known | Event handler |
| 0x007DFE41 | `HandleWheel` | Known | Event handler |
| 0x007DFED9 | `HandleImageNext` | Known | Event handler |
| 0x007DFF08 | `HandlePlayPause` | Known | Event handler |
| 0x007DFF57 | `HandleImagePrev` | Known | Event handler |
| 0x007DFF83 | `HandleSelect` | Known | Event handler |
| 0x007E0253 | `HandleImageNext` | Known | Event handler |
| 0x007E027D | `HandlePause` | Known | Event handler |
| 0x007E02A2 | `HandlePlay` | Known | Event handler |
| 0x007E02CB | `HandlePlayPause` | Known | Event handler |
| 0x007E02F4 | `HandleImagePrev` | Known | Event handler |
| 0x007E034D | `HandleWheel` | Known | Event handler |
| 0x007E03E5 | `HandleImageNext` | Known | Event handler |
| 0x007E0414 | `HandlePlayPause` | Known | Event handler |
| 0x007E0463 | `HandleImagePrev` | Known | Event handler |
| 0x007E048F | `HandleSelect` | Known | Event handler |
| 0x007E075F | `HandleImageNext` | Known | Event handler |
| 0x007E0789 | `HandlePause` | Known | Event handler |
| 0x007E07AE | `HandlePlay` | Known | Event handler |
| 0x007E07D7 | `HandlePlayPause` | Known | Event handler |
| 0x007E0800 | `HandleImagePrev` | Known | Event handler |
| 0x007E0859 | `HandleWheel` | Known | Event handler |
| 0x007E08F1 | `HandleImageNext` | Known | Event handler |
| 0x007E0920 | `HandlePlayPause` | Known | Event handler |
| 0x007E096F | `HandleImagePrev` | Known | Event handler |
| 0x007E099B | `HandleSelect` | Known | Event handler |
| 0x007E0BE6 | `HandleImageNext` | Known | Event handler |
| 0x007E0C10 | `HandlePause` | Known | Event handler |
| 0x007E0C35 | `HandlePlay` | Known | Event handler |
| 0x007E0C5E | `HandlePlayPause` | Known | Event handler |
| 0x007E0C87 | `HandleImagePrev` | Known | Event handler |
| 0x007E0CF0 | `HandleWheel` | Known | Event handler |
| 0x007E0D89 | `HandleImageNext` | Known | Event handler |
| 0x007E0DB8 | `HandlePlayPause` | Known | Event handler |
| 0x007E0E07 | `HandleImagePrev` | Known | Event handler |
| 0x007E0E33 | `HandleSelect` | Known | Event handler |
| 0x007E107E | `HandleImageNext` | Known | Event handler |
| 0x007E10A8 | `HandlePause` | Known | Event handler |
| 0x007E10CD | `HandlePlay` | Known | Event handler |
| 0x007E10F6 | `HandlePlayPause` | Known | Event handler |
| 0x007E111F | `HandleImagePrev` | Known | Event handler |
| 0x007E1188 | `HandleWheel` | Known | Event handler |
| 0x007E1221 | `HandleImageNext` | Known | Event handler |
| 0x007E1250 | `HandlePlayPause` | Known | Event handler |
| 0x007E129F | `HandleImagePrev` | Known | Event handler |
| 0x007E12CB | `HandleSelect` | Known | Event handler |
| 0x007E1516 | `HandleImageNext` | Known | Event handler |
| 0x007E1540 | `HandlePause` | Known | Event handler |
| 0x007E1565 | `HandlePlay` | Known | Event handler |
| 0x007E158E | `HandlePlayPause` | Known | Event handler |
| 0x007E15B7 | `HandleImagePrev` | Known | Event handler |
| 0x007E1620 | `HandleWheel` | Known | Event handler |
| 0x007E16B9 | `HandleImageNext` | Known | Event handler |
| 0x007E16E8 | `HandlePlayPause` | Known | Event handler |
| 0x007E1737 | `HandleImagePrev` | Known | Event handler |
| 0x007E1763 | `HandleSelect` | Known | Event handler |
| 0x007E19AE | `HandleImageNext` | Known | Event handler |
| 0x007E19D8 | `HandlePause` | Known | Event handler |
| 0x007E19FD | `HandlePlay` | Known | Event handler |
| 0x007E1A26 | `HandlePlayPause` | Known | Event handler |
| 0x007E1A4F | `HandleImagePrev` | Known | Event handler |
| 0x007E1AB8 | `HandleWheel` | Known | Event handler |
| 0x007E1B51 | `HandleImageNext` | Known | Event handler |
| 0x007E1B80 | `HandlePlayPause` | Known | Event handler |
| 0x007E1BCF | `HandleImagePrev` | Known | Event handler |
| 0x007E1BFB | `HandleSelect` | Known | Event handler |
| 0x007E1E46 | `HandleImageNext` | Known | Event handler |
| 0x007E1E70 | `HandlePause` | Known | Event handler |
| 0x007E1E95 | `HandlePlay` | Known | Event handler |
| 0x007E1EBE | `HandlePlayPause` | Known | Event handler |
| 0x007E1EE7 | `HandleImagePrev` | Known | Event handler |
| 0x007E1F50 | `HandleWheel` | Known | Event handler |
| 0x007E1F7D | `HandleSelect` | Known | Event handler |
| 0x007E1FAD | `HandleSelect` | Known | Event handler |
| 0x007E20D0 | `HandleTuning` | Known | Event handler |
| 0x007E228C | `HandleVolumeChange` | Known | Event handler |
| 0x007E23D8 | `HandleVolumeWheel` | Known | Event handler |
| 0x007E2533 | `HandleTuningSelect` | Known | Event handler |
| 0x007E2812 | `HandleFrequencyChange` | Known | Event handler |
| 0x007E296F | `HandleTuningSelect` | Known | Event handler |
| 0x007E2C4E | `HandleFrequencyChange` | Known | Event handler |
| 0x007E2D78 | `HandleTimerDone` | Known | Event handler |
| 0x007E2F6D | `HandleVolumeChange` | Known | Event handler |
| 0x007E3084 | `HandleVolumeWheel` | Known | Event handler |
| 0x007E3667 | `HandleExitUnsupported` | Known | Event handler |
| 0x007E3699 | `HandleExitUnsupported` | Known | Event handler |
| 0x007E86CD | `HandleSelectKey` | Known | Event handler |
| 0x007E8702 | `HandleWheel` | Known | Event handler |
| 0x007E8850 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x007E88A3 | `HandleSelectKey` | Known | Event handler |
| 0x007E88CB | `HandleSelectKey` | Known | Event handler |
| 0x007E88FB | `HandleExit` | Known | Event handler |
| 0x007E8925 | `HandleStartStop` | Known | Event handler |
| 0x007E898B | `HandleStartStop` | Known | Event handler |
| 0x007E8AA3 | `HandleExit` | Known | Event handler |
| 0x007E8ACD | `HandleStartStop` | Known | Event handler |
| 0x007E8AF9 | `HandleLap` | Known | Event handler |
| 0x007E8BFD | `HandleSelectLozinch` | Known | Event handler |
| 0x007E96A4 | `HandleSelect` | Known | Event handler |
| 0x007EA1AF | `HandleChoosePowerPlay` | Known | Event handler |
| 0x007EA1EA | `HandleChooseVoicePrompts` | Known | Event handler |
| 0x007EA228 | `HandleChooseUnit` | Known | Event handler |
| 0x007EA3BC | `HandleListChoose` | Known | Event handler |
| 0x007EA61B | `HandleSelect` | Known | Event handler |
| 0x007EA83B | `HandleSelect` | Known | Event handler |
| 0x007EA871 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EAAA2 | `HandleNowPlayingSelected` | Known | Event handler |
| 0x007EAAE0 | `HandleChoosePlaylistsPlay` | Known | Event handler |
| 0x007EAB1F | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x007EAB5F | `HandleNoneSelected` | Known | Event handler |
| 0x007EAB95 | `HandleBegin` | Known | Event handler |
| 0x007EAE82 | `HandleBegin` | Known | Event handler |
| 0x007EAEB1 | `HandleBegin` | Known | Event handler |
| 0x007EAF6D | `HandleBegin` | Known | Event handler |
| 0x007EAF99 | `HandleBegin` | Known | Event handler |
| 0x007EB055 | `HandleBegin` | Known | Event handler |
| 0x007EB081 | `HandleBegin` | Known | Event handler |
| 0x007EB13D | `HandleBegin` | Known | Event handler |
| 0x007EB171 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EB19C | `HandleMenuKey` | Known | Event handler |
| 0x007EB233 | `HandlePauseHold` | Known | Event handler |
| 0x007EB262 | `HandlePauseKey` | Known | Event handler |
| 0x007EB2EC | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EB326 | `HandlePowerPlay` | Known | Event handler |
| 0x007EB352 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EB7BF | `HandlePauseHold` | Known | Event handler |
| 0x007EB7EE | `HandlePauseKey` | Known | Event handler |
| 0x007EB819 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EB857 | `HandlePowerPlay` | Known | Event handler |
| 0x007EB886 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EB8AC | `HandleWheel` | Known | Event handler |
| 0x007EB8E1 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EB90C | `HandleMenuKey` | Known | Event handler |
| 0x007EB9A3 | `HandlePauseHold` | Known | Event handler |
| 0x007EB9D2 | `HandlePauseKey` | Known | Event handler |
| 0x007EBA5C | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EBA8C | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EBEEC | `HandlePauseHold` | Known | Event handler |
| 0x007EBF1B | `HandlePauseKey` | Known | Event handler |
| 0x007EBF46 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EBF7A | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EBFA0 | `HandleWheel` | Known | Event handler |
| 0x007EBFD5 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EC000 | `HandleMenuKey` | Known | Event handler |
| 0x007EC097 | `HandlePauseHold` | Known | Event handler |
| 0x007EC0C6 | `HandlePauseKey` | Known | Event handler |
| 0x007EC150 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EC18A | `HandlePowerPlay` | Known | Event handler |
| 0x007EC1B6 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EC622 | `HandlePauseHold` | Known | Event handler |
| 0x007EC651 | `HandlePauseKey` | Known | Event handler |
| 0x007EC67C | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EC6BA | `HandlePowerPlay` | Known | Event handler |
| 0x007EC6E9 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EC70F | `HandleWheel` | Known | Event handler |
| 0x007EC745 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EC770 | `HandleMenuKey` | Known | Event handler |
| 0x007EC807 | `HandlePauseHold` | Known | Event handler |
| 0x007EC836 | `HandlePauseKey` | Known | Event handler |
| 0x007EC8C0 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EC8F0 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ECD4F | `HandlePauseHold` | Known | Event handler |
| 0x007ECD7E | `HandlePauseKey` | Known | Event handler |
| 0x007ECDA9 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ECDDD | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ECE03 | `HandleWheel` | Known | Event handler |
| 0x007ECE39 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007ECE64 | `HandleMenuKey` | Known | Event handler |
| 0x007ECEFB | `HandlePauseHold` | Known | Event handler |
| 0x007ECF2A | `HandlePauseKey` | Known | Event handler |
| 0x007ECFB4 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007ECFEE | `HandlePowerPlay` | Known | Event handler |
| 0x007ED01A | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ED48A | `HandlePauseHold` | Known | Event handler |
| 0x007ED4B9 | `HandlePauseKey` | Known | Event handler |
| 0x007ED4E4 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ED522 | `HandlePowerPlay` | Known | Event handler |
| 0x007ED551 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007ED577 | `HandleWheel` | Known | Event handler |
| 0x007ED5AD | `HandleMenuKeyNop` | Known | Event handler |
| 0x007ED5D8 | `HandleMenuKey` | Known | Event handler |
| 0x007ED66F | `HandlePauseHold` | Known | Event handler |
| 0x007ED69E | `HandlePauseKey` | Known | Event handler |
| 0x007ED728 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007ED758 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EDBBB | `HandlePauseHold` | Known | Event handler |
| 0x007EDBEA | `HandlePauseKey` | Known | Event handler |
| 0x007EDC15 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EDC49 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EDC6F | `HandleWheel` | Known | Event handler |
| 0x007EDCA5 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EDCD0 | `HandleMenuKey` | Known | Event handler |
| 0x007EDD67 | `HandlePauseHold` | Known | Event handler |
| 0x007EDD96 | `HandlePauseKey` | Known | Event handler |
| 0x007EDE20 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EDE5A | `HandlePowerPlay` | Known | Event handler |
| 0x007EDE86 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EE2F6 | `HandlePauseHold` | Known | Event handler |
| 0x007EE325 | `HandlePauseKey` | Known | Event handler |
| 0x007EE350 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EE38E | `HandlePowerPlay` | Known | Event handler |
| 0x007EE3BD | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EE3E3 | `HandleWheel` | Known | Event handler |
| 0x007EE419 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EE444 | `HandleMenuKey` | Known | Event handler |
| 0x007EE4DB | `HandlePauseHold` | Known | Event handler |
| 0x007EE50A | `HandlePauseKey` | Known | Event handler |
| 0x007EE594 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EE5C4 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EEA27 | `HandlePauseHold` | Known | Event handler |
| 0x007EEA56 | `HandlePauseKey` | Known | Event handler |
| 0x007EEA81 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EEAB5 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EEADB | `HandleWheel` | Known | Event handler |
| 0x007EEB11 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EEB3C | `HandleMenuKey` | Known | Event handler |
| 0x007EEBD3 | `HandlePauseHold` | Known | Event handler |
| 0x007EEC02 | `HandlePauseKey` | Known | Event handler |
| 0x007EEC8C | `HandleSelectKeyDown` | Known | Event handler |
| 0x007EECBC | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EF0B5 | `HandlePauseHold` | Known | Event handler |
| 0x007EF0E4 | `HandlePauseKey` | Known | Event handler |
| 0x007EF10F | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EF143 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007EF169 | `HandleWheel` | Known | Event handler |
| 0x007EF19D | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EF1C8 | `HandleResumeWorkout` | Known | Event handler |
| 0x007EF2A3 | `HandleResumeWorkout` | Known | Event handler |
| 0x007EF317 | `HandlePauseWorkout` | Known | Event handler |
| 0x007EF385 | `HandleChooseMusic` | Known | Event handler |
| 0x007EF422 | `HandleEndWorkout` | Known | Event handler |
| 0x007EF4CD | `HandleMenuKeyNop` | Known | Event handler |
| 0x007EF774 | `HandleEndWorkout` | Known | Event handler |
| 0x007EFC03 | `HandleSelectResume` | Known | Event handler |
| 0x007EFC3B | `HandleEndWorkout` | Known | Event handler |
| 0x007EFCE6 | `HandleChooseCustomTimedWorkout` | Known | Event handler |
| 0x007EFD7F | `HandleChooseTimedWorkoutFromList` | Known | Event handler |
| 0x007EFE32 | `HandleChooseCustomDistanceWorkout` | Known | Event handler |
| 0x007EFED2 | `HandleChooseDistanceWorkoutFromList` | Known | Event handler |
| 0x007F00B8 | `HandleChooseCustomCalorieWorkout` | Known | Event handler |
| 0x007F0157 | `HandleChooseCalorieWorkoutFromList` | Known | Event handler |
| 0x007F04C6 | `HandleChooseLink` | Known | Event handler |
| 0x007F04FC | `HandleChooseCalibrate` | Known | Event handler |
| 0x007F0855 | `HandleChooseRemoteSetting` | Known | Event handler |
| 0x007F0894 | `HandleChooseRemoteLink` | Known | Event handler |
| 0x007F08D0 | `HandleChooseRemoteUnlink` | Known | Event handler |
| 0x007F0C7E | `Handle400MetersWalk` | Known | Event handler |
| 0x007F0CB7 | `HandleCustomWalk` | Known | Event handler |
| 0x007F0D8D | `HandleSelectWalking` | Known | Event handler |
| 0x007F0EB1 | `HandleSelectRunning` | Known | Event handler |
| 0x007F11FE | `Handle400MetersRun` | Known | Event handler |
| 0x007F1236 | `HandleCustomRun` | Known | Event handler |
| 0x007F1509 | `HandleSelect` | Known | Event handler |
| 0x007F1539 | `HandleSelect` | Known | Event handler |
| 0x007F16AF | `HandleLinkNewRemote` | Known | Event handler |
| 0x007F181D | `HandleSelect` | Known | Event handler |
| 0x007F184B | `HandleCancelRemoteLinking` | Known | Event handler |
| 0x007F18B8 | `HandleSelect` | Known | Event handler |
| 0x007F1DA9 | `HandleUnlinkRemote` | Known | Event handler |
| 0x007F200D | `HandleWeightSelect` | Known | Event handler |
| 0x007F206A | `HandleWeightWheel` | Known | Event handler |
| 0x007F209D | `HandleWeightSelect` | Known | Event handler |
| 0x007F2127 | `HandleWeightWheel` | Known | Event handler |
| 0x007F2159 | `HandleDistanceSelect` | Known | Event handler |
| 0x007F21E5 | `HandleDistanceWheel` | Known | Event handler |
| 0x007F2219 | `HandleDistanceSelect` | Known | Event handler |
| 0x007F22A5 | `HandleDistanceWheel` | Known | Event handler |
| 0x007F22D9 | `HandleTimeSelect` | Known | Event handler |
| 0x007F2361 | `HandleTimeWheel` | Known | Event handler |
| 0x007F2391 | `HandleCaloriesSelect` | Known | Event handler |
| 0x007F24E9 | `HandleCaloriesWheel` | Known | Event handler |
| 0x007F2855 | `HandleChooseLast` | Known | Event handler |
| 0x007F288B | `HandleChooseRecent` | Known | Event handler |
| 0x007F28C3 | `HandleChooseBest` | Known | Event handler |
| 0x007F2BD9 | `HandleSelect` | Known | Event handler |
| 0x007F2DC1 | `HandleSelectFromLastWorkout` | Known | Event handler |
| 0x007F2FB9 | `HandleSelect` | Known | Event handler |
| 0x007F3272 | `HandleChooseRecentWorkout` | Known | Event handler |
| 0x007F3345 | `HandleSelect` | Known | Event handler |
| 0x007F33D9 | `HandleSelect_Basic` | Known | Event handler |
| 0x007F36BD | `HandleSelect_Dynamic` | Known | Event handler |
| 0x007F39B1 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x007F3CA1 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x007F4092 | `HandleSelect` | Known | Event handler |
| 0x007F411E | `HandleSelect` | Known | Event handler |
| 0x007F41AC | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x007F4496 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x007F4577 | `HandlePlayPause` | Known | Event handler |
| 0x007F4605 | `HandlePlayPause` | Known | Event handler |
| 0x007F4695 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x007F46CD | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x007F4709 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x007F474C | `HandlePlayPause` | Known | Event handler |
| 0x007F4782 | `HandleAddToOTG` | Known | Event handler |
| 0x007F49D7 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007F4C33 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x00817E32 | `HandleSelectClock` | Known | Event handler |
| 0x00817E6B | `HandleHilited` | Known | Event handler |
| 0x00817E9D | `HandleWheel` | Known | Event handler |
| 0x00817EE4 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00817F69 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00818175 | `HandleImageLast` | Known | Event handler |
| 0x0081819F | `HandleScreenNext` | Known | Event handler |
| 0x008181CF | `HandleBrowseSlideshow` | Known | Event handler |
| 0x00818209 | `HandleImageFirst` | Known | Event handler |
| 0x00818234 | `HandleScreenPrev` | Known | Event handler |
| 0x00818261 | `HandleBrowseLarge` | Known | Event handler |
| 0x008182E1 | `HandleImageNext` | Known | Event handler |
| 0x0081830A | `HandleBrowseSlideshow` | Known | Event handler |
| 0x0081833E | `HandleBrowseSlideshow` | Known | Event handler |
| 0x0081836D | `HandleImagePrev` | Known | Event handler |
| 0x0081839B | `HandleBrowseSmall` | Known | Event handler |

---

## 6. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00102C7C | `GotoNowPlaying` | Known | Navigation |
| 0x00102CF4 | `GotoMainMenu` | Known | Navigation |
| 0x001211D0 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x001211E8 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x00121360 | `GotoScreen_AddressBook` | Known | Navigation |
| 0x0012C7CC | `GotoNowPlaying` | Known | Navigation |
| 0x0012C7E0 | `GotoAlbums` | Known | Navigation |
| 0x0012C7EC | `GotoSongs` | Known | Navigation |
| 0x0013A2DC | `GotoScreen_EnterPassKey` | Known | Navigation |
| 0x0013A2F4 | `GotoScreen_LockediPod` | Known | Navigation |
| 0x0013ACF8 | `GotoScreen_MainMenu` | Known | Navigation |
| 0x00154350 | `GotoMainMenu` | Known | Navigation |
| 0x001DEAC0 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001EA050 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001EA8A0 | `GotoPlayDeleteMenu` | Known | Navigation |
| 0x001EA924 | `GotoNowPlaying` | Known | Navigation |
| 0x00209B18 | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation |
| 0x00217300 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x002173F8 | `GotoScreen_SettingsMenu` | Known | Navigation |
| 0x002202BC | `GotoDefaultLayout` | Known | Navigation |
| 0x00220340 | `GotoVolumeLayout` | Known | Navigation |
| 0x00220478 | `GotoProgressLayout` | Known | Navigation |
| 0x00220794 | `GotoDefault` | Known | Navigation |
| 0x00220AC8 | `GotoProgressLayout` | Known | Navigation |
| 0x00220C88 | `GotoRentalWarningLayout` | Known | Navigation |
| 0x00220D0C | `GotoProgressLayout` | Known | Navigation |
| 0x0022101C | `GotoProgressLayout` | Known | Navigation |
| 0x00222B80 | `GotoNowPlaying` | Known | Navigation |
| 0x0022344C | `GotoNowPlaying` | Known | Navigation |
| 0x00226D1C | `GotoScreen_Language` | Known | Navigation |
| 0x0022707C | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x00227098 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x002270B0 | `GotoDefaultLayout` | Known | Navigation |
| 0x002270C4 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x0022715C | `GotoVolumeLayout` | Known | Navigation |
| 0x00227170 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x00227210 | `GotoProgressLayout` | Known | Navigation |
| 0x00227224 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x002276EC | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00227A28 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x00227BE0 | `GotoProgressLayout` | Known | Navigation |
| 0x00227BF4 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00227CB8 | `GotoBrightnessVideoLayout` | Known | Navigation |
| 0x00227CD4 | `GotoRatingLayout` | Known | Navigation |
| 0x00227F78 | `GotoChapterArtLayout` | Known | Navigation |
| 0x00227F90 | `GotoShuffleLayout` | Known | Navigation |
| 0x00228320 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x00228334 | `GotoExtraInfoLoadingLayout` | Known | Navigation |
| 0x00228404 | `GotoVolumeLayout` | Known | Navigation |
| 0x0022841C | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x002284A8 | `GotoVolumeLayout` | Known | Navigation |
| 0x002284BC | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x002286CC | `GotoScrubLayout` | Known | Navigation |
| 0x002286DC | `GotoScrubVideoLayout` | Known | Navigation |
| 0x0022876C | `GotoProgressLayout` | Known | Navigation |
| 0x00228780 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00228920 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x0022893C | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00228954 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00228970 | `GotoDefaultLayout` | Known | Navigation |
| 0x00228BBC | `GotoChapterArtLayout` | Known | Navigation |
| 0x00228CB4 | `GotoProgressLayout` | Known | Navigation |
| 0x00228D40 | `GotoProgressLayout` | Known | Navigation |
| 0x00228D54 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00228E30 | `GotoExtraInfoLoadFailedLayout` | Known | Navigation |
| 0x00228E50 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x0022928C | `GotoStatusBarLayout` | Known | Navigation |
| 0x002292A0 | `GotoDefaultLayout` | Known | Navigation |
| 0x00229478 | `GotoDefault` | Known | Navigation |
| 0x002295AC | `GotoProgressLayout` | Known | Navigation |
| 0x0022976C | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x002298BC | `GotoBrightnessLayout` | Known | Navigation |
| 0x00229940 | `GotoBrightnessLayout` | Known | Navigation |
| 0x002299C0 | `GotoVolumeLayout` | Known | Navigation |
| 0x00229A0C | `GotoScrubLayout` | Known | Navigation |
| 0x00229AD4 | `GotoStatusBarLayout` | Known | Navigation |
| 0x00229AE8 | `GotoDefaultLayout` | Known | Navigation |
| 0x00229BC0 | `GotoScrubLayout` | Known | Navigation |
| 0x00229C10 | `GotoScrubLayout` | Known | Navigation |
| 0x002306F0 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x00230880 | `GotoFourCard_About` | Known | Navigation |
| 0x00230894 | `GotoThreeCard_About` | Known | Navigation |
| 0x00230980 | `GotoScreen_BacklightTimer` | Known | Navigation |
| 0x00230A10 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x00230A28 | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation |
| 0x00235CCC | `GotoScreen_LockDialog` | Known | Navigation |
| 0x00235CE4 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x002382D8 | `GotoNowPlaying` | Known | Navigation |
| 0x00238A30 | `GotoNowPlaying` | Known | Navigation |
| 0x002390B0 | `GotoFirstBoot` | Known | Navigation |
| 0x002390C0 | `GotoNotesApp` | Known | Navigation |
| 0x002390D4 | `GotoLockApp` | Known | Navigation |
| 0x0023F07C | `GotoNowPlaying` | Known | Navigation |
| 0x003DAF30 | `GotoProgressLayout` | Known | Navigation |
| 0x0075F84F | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x007D7B71 | `GotoDefault` | Known | Navigation |
| 0x007D8661 | `GotoDefault` | Known | Navigation |
| 0x008E0444 | `Gotowe?` | Known | Navigation |

---

## 7. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016B540 | `CoverFlow_Screen` | Known | Screen layout |
| 0x001976A0 | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x001976C0 | `VoiceMemos_Menu_Screen_Recording` | Known | Screen layout |
| 0x001976E4 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x0075385A | `Clock_Screen` | Known | Screen layout |
| 0x0075386A | `Clock_Screen_Default"` | Known | Screen layout |
| 0x007538CF | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x0075392D | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x00753945 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x007539B2 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x00753A50 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x00753AAF | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x00753AC5 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x00753B30 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x00753B8A | `Games_Menu_Screen` | Known | Screen layout |
| 0x00753B9F | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x00753C09 | `Extras_Screen_Games` | Known | Screen layout |
| 0x00753CC8 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x00753D8C | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00753E55 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x00753EB2 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x00753ECB | `Debug_MainMenu_Screen_Default"` | Known | Screen layout |
| 0x00753F39 | `Extras_Screen_Debug` | Known | Screen layout |
| 0x00754070 | `Clock_Africa_City_Screen ` | Known | Screen layout |
| 0x0075408C | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x00754110 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x0075412A | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x007541AC | `Clock_Atlantic_City_Screen"` | Known | Screen layout |
| 0x007541CA | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x00754250 | `Clock_Australia_City_Screen#` | Known | Screen layout |
| 0x0075426F | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x007542F6 | `Clock_Europe_City_Screen ` | Known | Screen layout |
| 0x00754312 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x00754396 | `Clock_NorthAmerica_City_Screen&` | Known | Screen layout |
| 0x007543B8 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x00754442 | `Clock_Pacific_City_Screen!` | Known | Screen layout |
| 0x0075445F | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x007544E4 | `Clock_SouthAmerica_City_Screen&` | Known | Screen layout |
| 0x00754506 | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x00754593 | `Clock_Screen"` | Known | Screen layout |
| 0x00754638 | `Clock_Screen"` | Known | Screen layout |
| 0x007546DD | `Clock_Screen"` | Known | Screen layout |
| 0x00754782 | `Clock_Screen"` | Known | Screen layout |
| 0x00754827 | `Clock_Screen"` | Known | Screen layout |
| 0x007548CC | `Clock_Screen"` | Known | Screen layout |
| 0x00754971 | `Clock_Screen"` | Known | Screen layout |
| 0x00754A16 | `Clock_Screen"` | Known | Screen layout |
| 0x00754ABB | `Clock_Screen"` | Known | Screen layout |
| 0x00754B60 | `Clock_Screen"` | Known | Screen layout |
| 0x00754C05 | `Clock_Screen"` | Known | Screen layout |
| 0x00754CAA | `Clock_Screen"` | Known | Screen layout |
| 0x00754D4F | `Clock_Screen"` | Known | Screen layout |
| 0x00754DF4 | `Clock_Screen"` | Known | Screen layout |
| 0x00754E99 | `Clock_Screen"` | Known | Screen layout |
| 0x00754F3E | `Clock_Screen"` | Known | Screen layout |
| 0x00754FE3 | `Clock_Screen"` | Known | Screen layout |
| 0x00755088 | `Clock_Screen"` | Known | Screen layout |
| 0x0075512D | `Clock_Screen"` | Known | Screen layout |
| 0x007551D2 | `Clock_Screen"` | Known | Screen layout |
| 0x00755277 | `Clock_Screen"` | Known | Screen layout |
| 0x0075531C | `Clock_Screen"` | Known | Screen layout |
| 0x007553C1 | `Clock_Screen"` | Known | Screen layout |
| 0x00755466 | `Clock_Screen"` | Known | Screen layout |
| 0x0075550B | `Clock_Screen"` | Known | Screen layout |
| 0x007555B0 | `Clock_Screen"` | Known | Screen layout |
| 0x00755655 | `Clock_Screen"` | Known | Screen layout |
| 0x007556FA | `Clock_Screen"` | Known | Screen layout |
| 0x0075579F | `Clock_Screen"` | Known | Screen layout |
| 0x00755844 | `Clock_Screen"` | Known | Screen layout |
| 0x007558E9 | `Clock_Screen"` | Known | Screen layout |
| 0x00755993 | `Clock_Screen"` | Known | Screen layout |
| 0x00755A38 | `Clock_Screen"` | Known | Screen layout |
| 0x00755ADD | `Clock_Screen"` | Known | Screen layout |
| 0x00755B82 | `Clock_Screen"` | Known | Screen layout |
| 0x00755C27 | `Clock_Screen"` | Known | Screen layout |
| 0x00755CCC | `Clock_Screen"` | Known | Screen layout |
| 0x00755D71 | `Clock_Screen"` | Known | Screen layout |
| 0x00755E16 | `Clock_Screen"` | Known | Screen layout |
| 0x00755EBB | `Clock_Screen"` | Known | Screen layout |
| 0x00755F60 | `Clock_Screen"` | Known | Screen layout |
| 0x00756005 | `Clock_Screen"` | Known | Screen layout |
| 0x007560AA | `Clock_Screen"` | Known | Screen layout |
| 0x0075614F | `Clock_Screen"` | Known | Screen layout |
| 0x007561F4 | `Clock_Screen"` | Known | Screen layout |
| 0x00756299 | `Clock_Screen"` | Known | Screen layout |
| 0x0075633E | `Clock_Screen"` | Known | Screen layout |
| 0x007563E3 | `Clock_Screen"` | Known | Screen layout |
| 0x00756488 | `Clock_Screen"` | Known | Screen layout |
| 0x0075652D | `Clock_Screen"` | Known | Screen layout |
| 0x007565D2 | `Clock_Screen"` | Known | Screen layout |
| 0x00756677 | `Clock_Screen"` | Known | Screen layout |
| 0x0075671C | `Clock_Screen"` | Known | Screen layout |
| 0x007567C1 | `Clock_Screen"` | Known | Screen layout |
| 0x00756866 | `Clock_Screen"` | Known | Screen layout |
| 0x0075690B | `Clock_Screen"` | Known | Screen layout |
| 0x007569B0 | `Clock_Screen"` | Known | Screen layout |
| 0x00756A55 | `Clock_Screen"` | Known | Screen layout |
| 0x00756AFA | `Clock_Screen"` | Known | Screen layout |
| 0x00756B9F | `Clock_Screen"` | Known | Screen layout |
| 0x00756C44 | `Clock_Screen"` | Known | Screen layout |
| 0x00756CE9 | `Clock_Screen"` | Known | Screen layout |
| 0x00756D8E | `Clock_Screen"` | Known | Screen layout |
| 0x00756E33 | `Clock_Screen"` | Known | Screen layout |
| 0x00756ED8 | `Clock_Screen"` | Known | Screen layout |
| 0x00756F7D | `Clock_Screen"` | Known | Screen layout |
| 0x00757022 | `Clock_Screen"` | Known | Screen layout |
| 0x007570C7 | `Clock_Screen"` | Known | Screen layout |
| 0x0075716C | `Clock_Screen"` | Known | Screen layout |
| 0x00757211 | `Clock_Screen"` | Known | Screen layout |
| 0x007572B6 | `Clock_Screen"` | Known | Screen layout |
| 0x0075735B | `Clock_Screen"` | Known | Screen layout |
| 0x00757400 | `Clock_Screen"` | Known | Screen layout |
| 0x007574A5 | `Clock_Screen"` | Known | Screen layout |
| 0x0075754A | `Clock_Screen"` | Known | Screen layout |
| 0x007575EF | `Clock_Screen"` | Known | Screen layout |
| 0x00757694 | `Clock_Screen"` | Known | Screen layout |
| 0x00757739 | `Clock_Screen"` | Known | Screen layout |
| 0x007577DE | `Clock_Screen"` | Known | Screen layout |
| 0x00757883 | `Clock_Screen"` | Known | Screen layout |
| 0x00757928 | `Clock_Screen"` | Known | Screen layout |
| 0x007579CD | `Clock_Screen"` | Known | Screen layout |
| 0x00757A72 | `Clock_Screen"` | Known | Screen layout |
| 0x00757B17 | `Clock_Screen"` | Known | Screen layout |
| 0x00757BBC | `Clock_Screen"` | Known | Screen layout |
| 0x00757C61 | `Clock_Screen"` | Known | Screen layout |
| 0x00757D06 | `Clock_Screen"` | Known | Screen layout |
| 0x00757DAB | `Clock_Screen"` | Known | Screen layout |
| 0x00757E57 | `Clock_Screen"` | Known | Screen layout |
| 0x00757EFC | `Clock_Screen"` | Known | Screen layout |
| 0x00757FA1 | `Clock_Screen"` | Known | Screen layout |
| 0x0075804B | `Clock_Screen"` | Known | Screen layout |
| 0x007580F0 | `Clock_Screen"` | Known | Screen layout |
| 0x00758195 | `Clock_Screen"` | Known | Screen layout |
| 0x0075823A | `Clock_Screen"` | Known | Screen layout |
| 0x007582DF | `Clock_Screen"` | Known | Screen layout |
| 0x00758384 | `Clock_Screen"` | Known | Screen layout |
| 0x00758429 | `Clock_Screen"` | Known | Screen layout |
| 0x007584CE | `Clock_Screen"` | Known | Screen layout |
| 0x00758577 | `Clock_Screen"` | Known | Screen layout |
| 0x0075861C | `Clock_Screen"` | Known | Screen layout |
| 0x007586C1 | `Clock_Screen"` | Known | Screen layout |
| 0x00758766 | `Clock_Screen"` | Known | Screen layout |
| 0x0075880B | `Clock_Screen"` | Known | Screen layout |
| 0x007588B0 | `Clock_Screen"` | Known | Screen layout |
| 0x00758955 | `Clock_Screen"` | Known | Screen layout |
| 0x007589FA | `Clock_Screen"` | Known | Screen layout |
| 0x00758A9F | `Clock_Screen"` | Known | Screen layout |
| 0x00758B44 | `Clock_Screen"` | Known | Screen layout |
| 0x00758BE9 | `Clock_Screen"` | Known | Screen layout |
| 0x00758C8E | `Clock_Screen"` | Known | Screen layout |
| 0x00758D33 | `Clock_Screen"` | Known | Screen layout |
| 0x00758DD8 | `Clock_Screen"` | Known | Screen layout |
| 0x00758E7D | `Clock_Screen"` | Known | Screen layout |
| 0x00758F22 | `Clock_Screen"` | Known | Screen layout |
| 0x00758FC7 | `Clock_Screen"` | Known | Screen layout |
| 0x0075906C | `Clock_Screen"` | Known | Screen layout |
| 0x00759111 | `Clock_Screen"` | Known | Screen layout |
| 0x007591B6 | `Clock_Screen"` | Known | Screen layout |
| 0x0075925B | `Clock_Screen"` | Known | Screen layout |
| 0x00759300 | `Clock_Screen"` | Known | Screen layout |
| 0x007593A5 | `Clock_Screen"` | Known | Screen layout |
| 0x0075944A | `Clock_Screen"` | Known | Screen layout |
| 0x007594EF | `Clock_Screen"` | Known | Screen layout |
| 0x00759594 | `Clock_Screen"` | Known | Screen layout |
| 0x00759639 | `Clock_Screen"` | Known | Screen layout |
| 0x007596DE | `Clock_Screen"` | Known | Screen layout |
| 0x00759783 | `Clock_Screen"` | Known | Screen layout |
| 0x00759828 | `Clock_Screen"` | Known | Screen layout |
| 0x007598CD | `Clock_Screen"` | Known | Screen layout |
| 0x00759972 | `Clock_Screen"` | Known | Screen layout |
| 0x00759A17 | `Clock_Screen"` | Known | Screen layout |
| 0x00759ABC | `Clock_Screen"` | Known | Screen layout |
| 0x00759B67 | `Clock_Screen"` | Known | Screen layout |
| 0x00759C0C | `Clock_Screen"` | Known | Screen layout |
| 0x00759CB1 | `Clock_Screen"` | Known | Screen layout |
| 0x00759D56 | `Clock_Screen"` | Known | Screen layout |
| 0x00759DFB | `Clock_Screen"` | Known | Screen layout |
| 0x00759EA0 | `Clock_Screen"` | Known | Screen layout |
| 0x00759F45 | `Clock_Screen"` | Known | Screen layout |
| 0x00759FEA | `Clock_Screen"` | Known | Screen layout |
| 0x0075A08F | `Clock_Screen"` | Known | Screen layout |
| 0x0075A134 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A1D9 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A27E | `Clock_Screen"` | Known | Screen layout |
| 0x0075A323 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A3C8 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A46D | `Clock_Screen"` | Known | Screen layout |
| 0x0075A512 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A5B7 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A65C | `Clock_Screen"` | Known | Screen layout |
| 0x0075A701 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A7A6 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A84B | `Clock_Screen"` | Known | Screen layout |
| 0x0075A8F0 | `Clock_Screen"` | Known | Screen layout |
| 0x0075A995 | `Clock_Screen"` | Known | Screen layout |
| 0x0075AA3A | `Clock_Screen"` | Known | Screen layout |
| 0x0075AADF | `Clock_Screen"` | Known | Screen layout |
| 0x0075AB84 | `Clock_Screen"` | Known | Screen layout |
| 0x0075AC29 | `Clock_Screen"` | Known | Screen layout |
| 0x0075ACCE | `Clock_Screen"` | Known | Screen layout |
| 0x0075AD73 | `Clock_Screen"` | Known | Screen layout |
| 0x0075AE18 | `Clock_Screen"` | Known | Screen layout |
| 0x0075AEBD | `Clock_Screen"` | Known | Screen layout |
| 0x0075AF62 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B007 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B0AC | `Clock_Screen"` | Known | Screen layout |
| 0x0075B151 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B1F6 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B29B | `Clock_Screen"` | Known | Screen layout |
| 0x0075B340 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B3E5 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B48A | `Clock_Screen"` | Known | Screen layout |
| 0x0075B52F | `Clock_Screen"` | Known | Screen layout |
| 0x0075B5D4 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B679 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B71E | `Clock_Screen"` | Known | Screen layout |
| 0x0075B7C3 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B868 | `Clock_Screen"` | Known | Screen layout |
| 0x0075B90D | `Clock_Screen"` | Known | Screen layout |
| 0x0075B9B2 | `Clock_Screen"` | Known | Screen layout |
| 0x0075BA57 | `Clock_Screen"` | Known | Screen layout |
| 0x0075BAFC | `Clock_Screen"` | Known | Screen layout |
| 0x0075BBA7 | `Clock_Screen"` | Known | Screen layout |
| 0x0075BC4C | `Clock_Screen"` | Known | Screen layout |
| 0x0075BCF1 | `Clock_Screen"` | Known | Screen layout |
| 0x0075BD96 | `Clock_Screen"` | Known | Screen layout |
| 0x0075BE3B | `Clock_Screen"` | Known | Screen layout |
| 0x0075BEE7 | `Clock_Screen"` | Known | Screen layout |
| 0x0075BF8C | `Clock_Screen"` | Known | Screen layout |
| 0x0075C031 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C0D6 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C17B | `Clock_Screen"` | Known | Screen layout |
| 0x0075C220 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C2C5 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C36A | `Clock_Screen"` | Known | Screen layout |
| 0x0075C40F | `Clock_Screen"` | Known | Screen layout |
| 0x0075C4B4 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C559 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C5FE | `Clock_Screen"` | Known | Screen layout |
| 0x0075C6A3 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C748 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C7ED | `Clock_Screen"` | Known | Screen layout |
| 0x0075C892 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C937 | `Clock_Screen"` | Known | Screen layout |
| 0x0075C9DA | `Settings_DateTime_SetDate_Screen(` | Known | Screen layout |
| 0x0075C9FE | `Settings_DateTime_SetDate_Screen_Default"` | Known | Screen layout |
| 0x0075CA77 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0075CADD | `Settings_DateTime_SetTime_Screen(` | Known | Screen layout |
| 0x0075CB01 | `Settings_DateTime_SetTime_Screen_Default"` | Known | Screen layout |
| 0x0075CB7A | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x0075CBE5 | `Settings_DateTime_SetTimeZone_Screen,` | Known | Screen layout |
| 0x0075CC0D | `Settings_DateTime_SetTimeZone_Screen_Default"` | Known | Screen layout |
| 0x0075CC8A | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x0075CD43 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0075CDF3 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0075D382 | `Search_Main_Screen` | Known | Screen layout |
| 0x0075D398 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x0075D8BA | `Extras_Screen` | Known | Screen layout |
| 0x0075D8CB | `Extras_Screen_WorldClock"` | Known | Screen layout |
| 0x0075D948 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x0075D9AA | `Clock_Screen` | Known | Screen layout |
| 0x0075D9BA | `Clock_Screen_Default` | Known | Screen layout |
| 0x0075DA41 | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x0075DAA7 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0075DABD | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x0075DB28 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x0075DB8A | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0075DBA2 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x0075DC0F | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x0075DC73 | `AddressViewer_Main_Screen!` | Known | Screen layout |
| 0x0075DC90 | `AddressViewer_Main_Screen_Default"` | Known | Screen layout |
| 0x0075DD02 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x0075DD69 | `Games_Menu_Screen` | Known | Screen layout |
| 0x0075DD7E | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x0075DDE8 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x0075DEAF | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x0075DF4B | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x0075E01C | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x0075E0DC | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x0075E140 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0075E15F | `VoiceMemos_Menu_Screen_Default"` | Known | Screen layout |
| 0x0075E1E2 | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x0075E248 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x0075E260 | `Speakers_Main_Screen_Default"` | Known | Screen layout |
| 0x0075E2E1 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x0075E345 | `Radio_Screen` | Known | Screen layout |
| 0x0075E355 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x0075E3CE | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x0075E42F | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0075E4CB | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x0075E58E | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x0075E64D | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x0075E70A | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x0075EB24 | `Radio_Screen` | Known | Screen layout |
| 0x0075EB34 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x0075EBAD | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x0075ED91 | `Search_Main_Screen` | Known | Screen layout |
| 0x0075EDA7 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x0075EED4 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x0075EF37 | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x0075F278 | `Video_Settings_Screen` | Known | Screen layout |
| 0x0075F291 | `Video_Settings_Screen_Default"` | Known | Screen layout |
| 0x0075F38E | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0075F653 | `SettingsMenus_MainMenu_Screen%` | Known | Screen layout |
| 0x0075F761 | `SettingsMenus_MusicMenu_Screen&` | Known | Screen layout |
| 0x0075FA0A | `SettingsMenus_Brightness_Screen&` | Known | Screen layout |
| 0x0075FB1F | `SettingsMenus_AudiobookSettings_Screen.` | Known | Screen layout |
| 0x0075FC55 | `SettingsMenus_RadioRegions_Screen ` | Known | Screen layout |
| 0x0075FD6A | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x0075FFD6 | `Settings_DateTime_Screen ` | Known | Screen layout |
| 0x0075FFF2 | `Settings_DateTime_Screen_Default"` | Known | Screen layout |
| 0x0076017E | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x00760283 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x0076029C | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x0076038D | `SettingsMenus_ResetAllSettings_Screen4` | Known | Screen layout |
| 0x00760B5E | `Stopwatch_Screen` | Known | Screen layout |
| 0x00760B72 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x00760BD9 | `Stopwatch_Screen` | Known | Screen layout |
| 0x00760BED | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x00760C96 | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x00760CB9 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x00760D52 | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x00760D75 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x00760E06 | `NikePlus_ResumeWorkout_Screen%` | Known | Screen layout |
| 0x00760E27 | `NikePlus_ResumeWorkout_Screen_Default"` | Known | Screen layout |
| 0x00760E9D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00760F43 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00760FF0 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007610A0 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00761150 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007611B2 | `NikePlus_Settings_Screen ` | Known | Screen layout |
| 0x007611CE | `NikePlus_Settings_Screen_Default` | Known | Screen layout |
| 0x00761251 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007612B3 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x007612CE | `NikePlus_History_Screen_Default` | Known | Screen layout |
| 0x00761350 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00761574 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007615E2 | `Speakers_ToneControl_Screen#` | Known | Screen layout |
| 0x00761601 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x007766BD | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00776740 | `LockediPod_Screen` | Known | Screen layout |
| 0x007767C8 | `Lock_Screen` | Known | Screen layout |
| 0x007767D7 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00776852 | `AddressViewer_ContactDetails_Screen+` | Known | Screen layout |
| 0x00776879 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x007768F4 | `Extras_Screen` | Known | Screen layout |
| 0x0077693F | `Extras_Screen` | Known | Screen layout |
| 0x00776A26 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x00776A84 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x00776AA1 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x00776B0F | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00776B28 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x00776B9F | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x00776BBC | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x00776C27 | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x00776C44 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x00776CAB | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x00776D12 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x00776D70 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x00776D8D | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x00776DFB | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00776E14 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x00776E8B | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x00776EA8 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x00776F13 | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x00776F30 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x00776F97 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x00777037 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x007770C0 | `Alarms_Set_Alarm_Frequency_Screen)` | Known | Screen layout |
| 0x007770E5 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x00777156 | `Alarms_Set_Alarm_Label_Screen%` | Known | Screen layout |
| 0x00777177 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x007771E4 | `Alarms_Set_Alarm_Sound_Screen%` | Known | Screen layout |
| 0x00777205 | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x00777271 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x007774EC | `Alarms_Set_Alarm_Playlist_Screen(` | Known | Screen layout |
| 0x00777510 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x00777580 | `Alarms_Set_Alarm_Tones_Screen%` | Known | Screen layout |
| 0x007775A1 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x007778B4 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x007778CF | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x00777A20 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x00777A37 | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x00777AB8 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x00777ACF | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x00777BA5 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00777BBE | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x00777C43 | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x00777CB4 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x00777DA9 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00777DC2 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x00777E47 | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x00777EB8 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x00777F78 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x00777F8C | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x007780BB | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x0077811E | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x00778175 | `Clock_Screen_Default` | Known | Screen layout |
| 0x00778206 | `Clock_Region_Screen` | Known | Screen layout |
| 0x0077821D | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x00778296 | `Clock_Screen_Default` | Known | Screen layout |
| 0x007782ED | `Clock_Screen_Default` | Known | Screen layout |
| 0x0077837E | `Clock_Region_Screen` | Known | Screen layout |
| 0x00778395 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x00778520 | `CoverFlow_Screen_DefaultNoAlbumInfo&` | Known | Screen layout |
| 0x0077860E | `CoverFlow_Screen_DefaultNoAlbumInfo#` | Known | Screen layout |
| 0x00778683 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00778979 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00778B29 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00778C57 | `CoverFlow_Screen_Default&` | Known | Screen layout |
| 0x00778D2D | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00778EC2 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00779127 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x00779184 | `Game_Screen` | Known | Screen layout |
| 0x00779193 | `Game_Screen_Default` | Known | Screen layout |
| 0x00779235 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00779297 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x007792FA | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0077935D | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x007793B9 | `Game_Running_Screen` | Known | Screen layout |
| 0x00779419 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0077947B | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x007794DE | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00779541 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0077959D | `Game_Running_Screen` | Known | Screen layout |
| 0x007795FD | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0077965F | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x007796C2 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00779725 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00779781 | `Game_Running_Screen` | Known | Screen layout |
| 0x007797E1 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00779843 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x007798A6 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00779909 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00779965 | `Game_Running_Screen` | Known | Screen layout |
| 0x007799C5 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00779A27 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00779A8A | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00779AED | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00779B49 | `Game_Running_Screen` | Known | Screen layout |
| 0x00779D8F | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00779DF1 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00779E54 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00779EB7 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00779F13 | `Game_Running_Screen` | Known | Screen layout |
| 0x00779FCA | `Extras_Screen` | Known | Screen layout |
| 0x00779FDB | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0077A039 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0077A1D6 | `Extras_Screen` | Known | Screen layout |
| 0x0077A1E7 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0077A245 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0077A3E2 | `Extras_Screen` | Known | Screen layout |
| 0x0077A3F3 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0077A451 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0077A5EE | `Extras_Screen` | Known | Screen layout |
| 0x0077A5FF | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0077A65D | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0077A7FF | `Lock_Screen` | Known | Screen layout |
| 0x0077A80E | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0077A870 | `Extras_Screen` | Known | Screen layout |
| 0x0077A881 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0077A8E0 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077A95A | `MainMenus_Main_Screen)` | Known | Screen layout |
| 0x0077AB2B | `Lock_Screen` | Known | Screen layout |
| 0x0077AB3A | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0077AB9C | `Extras_Screen` | Known | Screen layout |
| 0x0077ABAD | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0077AC0C | `LockediPod_Screen` | Known | Screen layout |
| 0x0077AC86 | `MainMenus_Main_Screen%` | Known | Screen layout |
| 0x0077ACED | `LockediPod_Screen` | Known | Screen layout |
| 0x0077AD02 | `LockediPod_Screen_Default,` | Known | Screen layout |
| 0x0077AE51 | `Lock_Screen` | Known | Screen layout |
| 0x0077AE60 | `Lock_Screen_Default_Layout"` | Known | Screen layout |
| 0x0077AEC9 | `Lock_Screen` | Known | Screen layout |
| 0x0077AED8 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0077AF3A | `Extras_Screen` | Known | Screen layout |
| 0x0077AF4B | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0077AFAA | `LockediPod_Screen` | Known | Screen layout |
| 0x0077B024 | `MainMenus_Main_Screen,` | Known | Screen layout |
| 0x0077B180 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0077B1E6 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0077B24A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077B2D9 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0077B346 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0077B3B3 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0077B420 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077B488 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0077B4EE | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0077B552 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077B5E1 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0077B64E | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0077B6BB | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0077B728 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077B790 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0077B7F6 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0077B85A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077B8E9 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0077B956 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0077B9C3 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0077BA30 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077BA98 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0077BAFE | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0077BB62 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077BBF1 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0077BC5E | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0077BCCB | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0077BD38 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077BDA0 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0077BE06 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0077BE6A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077BEF9 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0077BF66 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0077BFD3 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0077C040 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077C099 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x0077C102 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x0077C169 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0077C204 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x0077C26D | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x0077C2D6 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x0077C33D | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0077C3D8 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x0077C441 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x0077C4AA | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x0077C511 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x0077C5AC | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x0077C698 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077C6B4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077C722 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077C73F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077C7AA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077C7CA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077C841 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077C85D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077C8CD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077C8EC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077C958 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077C96C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077C9E5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077CA59 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077CAC9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077CB31 | `NoContent_Screen` | Known | Screen layout |
| 0x0077CB45 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077CBA9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077CC10 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077CC2A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077CC98 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077CD0A | `NoContent_Screen` | Known | Screen layout |
| 0x0077CD1E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077CD88 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077CDF1 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077CE05 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077CE6B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077CED9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077CF46 | `NoContent_Screen` | Known | Screen layout |
| 0x0077CF5A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077CFC2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077D02C | `NoContent_Screen` | Known | Screen layout |
| 0x0077D040 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077D0A7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077D111 | `NoContent_Screen` | Known | Screen layout |
| 0x0077D125 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077D192 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077D204 | `NoContent_Screen` | Known | Screen layout |
| 0x0077D218 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077D280 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077D2E9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077D304 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077D36A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077D386 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077D465 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077D47E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077D4DF | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077D4F3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077D661 | `Radio_Screen` | Known | Screen layout |
| 0x0077D671 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077D6D2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077D755 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077D7DD | `Lock_Screen` | Known | Screen layout |
| 0x0077D7EC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077D84F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077D8B1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077D8CD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077D93F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077D95E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077D9C6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077D9E0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077DA48 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077DA65 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077DAD1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077DB3B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077DB55 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077DBC5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077DC38 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077DCA9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077DD18 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077DD84 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077DD9F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077DE14 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077DE7B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077DEDD | `Photos_Screen` | Known | Screen layout |
| 0x0077DF41 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077DF5F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077DFD1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077DFEE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077E054 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077E06F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077E0D8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077E0F5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077E16C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077E190 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077E1FE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077E219 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077E2D4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077E2F0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077E35E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077E37B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077E3E6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077E406 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077E47D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077E499 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077E509 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077E528 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077E594 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077E5A8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077E621 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077E695 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077E705 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077E76D | `NoContent_Screen` | Known | Screen layout |
| 0x0077E781 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077E7E5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077E84C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077E866 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077E8D4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077E946 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E95A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077E9C4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077EA2D | `No_Photos_Screen` | Known | Screen layout |
| 0x0077EA41 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077EAA7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077EB15 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077EB82 | `NoContent_Screen` | Known | Screen layout |
| 0x0077EB96 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077EBFE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077EC68 | `NoContent_Screen` | Known | Screen layout |
| 0x0077EC7C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077ECE3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077ED4D | `NoContent_Screen` | Known | Screen layout |
| 0x0077ED61 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077EDCE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077EE40 | `NoContent_Screen` | Known | Screen layout |
| 0x0077EE54 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077EEBC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077EF25 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077EF40 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077EFA6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077EFC2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077F0A1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077F0BA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077F11B | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077F12F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077F29D | `Radio_Screen` | Known | Screen layout |
| 0x0077F2AD | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077F30E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077F391 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077F419 | `Lock_Screen` | Known | Screen layout |
| 0x0077F428 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077F48B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077F4ED | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077F509 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077F57B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077F59A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077F602 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077F61C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077F684 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077F6A1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077F70D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077F777 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077F791 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077F801 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077F874 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077F8E5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077F954 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077F9C0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077F9DB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077FA50 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077FAB7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077FB19 | `Photos_Screen` | Known | Screen layout |
| 0x0077FB7D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077FB9B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077FC0D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077FC2A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077FC90 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077FCAB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077FD14 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077FD31 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077FDA8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077FDCC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077FE3A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077FE55 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077FF10 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077FF2C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077FF9A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077FFB7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00780022 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00780042 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007800B9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007800D5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00780145 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00780164 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007801D0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007801E4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078025D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007802D1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00780341 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007803A9 | `NoContent_Screen` | Known | Screen layout |
| 0x007803BD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00780421 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00780488 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007804A2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00780510 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00780582 | `NoContent_Screen` | Known | Screen layout |
| 0x00780596 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00780600 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00780669 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078067D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007806E3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00780751 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007807BE | `NoContent_Screen` | Known | Screen layout |
| 0x007807D2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078083A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007808A4 | `NoContent_Screen` | Known | Screen layout |
| 0x007808B8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078091F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00780989 | `NoContent_Screen` | Known | Screen layout |
| 0x0078099D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00780A0A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00780A7C | `NoContent_Screen` | Known | Screen layout |
| 0x00780A90 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00780AF8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00780B61 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00780B7C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00780BE2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00780BFE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00780CDD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00780CF6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00780D57 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00780D6B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00780ED9 | `Radio_Screen` | Known | Screen layout |
| 0x00780EE9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00780F4A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00780FCD | `LockediPod_Screen` | Known | Screen layout |
| 0x00781055 | `Lock_Screen` | Known | Screen layout |
| 0x00781064 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007810C7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00781129 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00781145 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007811B7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007811D6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078123E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00781258 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007812C0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007812DD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00781349 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007813B3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007813CD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078143D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007814B0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00781521 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00781590 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007815FC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00781617 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078168C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007816F3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00781755 | `Photos_Screen` | Known | Screen layout |
| 0x007817B9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007817D7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00781849 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00781866 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007818CC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007818E7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00781950 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078196D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007819E4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00781A08 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00781A76 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00781A91 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00781B4C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00781B68 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00781BD6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00781BF3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00781C5E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00781C7E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00781CF5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00781D11 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00781D81 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00781DA0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00781E0C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00781E20 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00781E99 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00781F0D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00781F7D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00781FE5 | `NoContent_Screen` | Known | Screen layout |
| 0x00781FF9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078205D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007820C4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007820DE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078214C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007821BE | `NoContent_Screen` | Known | Screen layout |
| 0x007821D2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078223C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007822A5 | `No_Photos_Screen` | Known | Screen layout |
| 0x007822B9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078231F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078238D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007823FA | `NoContent_Screen` | Known | Screen layout |
| 0x0078240E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00782476 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007824E0 | `NoContent_Screen` | Known | Screen layout |
| 0x007824F4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078255B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007825C5 | `NoContent_Screen` | Known | Screen layout |
| 0x007825D9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00782646 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007826B8 | `NoContent_Screen` | Known | Screen layout |
| 0x007826CC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00782734 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078279D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007827B8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078281E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078283A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00782919 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00782932 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00782993 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007829A7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00782B15 | `Radio_Screen` | Known | Screen layout |
| 0x00782B25 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00782B86 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00782C09 | `LockediPod_Screen` | Known | Screen layout |
| 0x00782C91 | `Lock_Screen` | Known | Screen layout |
| 0x00782CA0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00782D03 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00782D65 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00782D81 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00782DF3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00782E12 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00782E7A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00782E94 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00782EFC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00782F19 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00782F85 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00782FEF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00783009 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00783079 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007830EC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078315D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007831CC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00783238 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00783253 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007832C8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078332F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00783391 | `Photos_Screen` | Known | Screen layout |
| 0x007833F5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00783413 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00783485 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007834A2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00783508 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00783523 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078358C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007835A9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00783620 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00783644 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007836B2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007836CD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00783788 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007837A4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00783812 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078382F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078389A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007838BA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00783931 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078394D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007839BD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007839DC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00783A48 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00783A5C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00783AD5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00783B49 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00783BB9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00783C21 | `NoContent_Screen` | Known | Screen layout |
| 0x00783C35 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00783C99 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00783D00 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00783D1A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00783D88 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00783DFA | `NoContent_Screen` | Known | Screen layout |
| 0x00783E0E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00783E78 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00783EE1 | `No_Photos_Screen` | Known | Screen layout |
| 0x00783EF5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00783F5B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00783FC9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00784036 | `NoContent_Screen` | Known | Screen layout |
| 0x0078404A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007840B2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078411C | `NoContent_Screen` | Known | Screen layout |
| 0x00784130 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00784197 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00784201 | `NoContent_Screen` | Known | Screen layout |
| 0x00784215 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00784282 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007842F4 | `NoContent_Screen` | Known | Screen layout |
| 0x00784308 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00784370 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007843D9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007843F4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078445A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00784476 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00784555 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078456E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007845CF | `FirstBoot_Screen` | Known | Screen layout |
| 0x007845E3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00784751 | `Radio_Screen` | Known | Screen layout |
| 0x00784761 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007847C2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00784845 | `LockediPod_Screen` | Known | Screen layout |
| 0x007848CD | `Lock_Screen` | Known | Screen layout |
| 0x007848DC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078493F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007849A1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007849BD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00784A2F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00784A4E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00784AB6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00784AD0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00784B38 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00784B55 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00784BC1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00784C2B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00784C45 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00784CB5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00784D28 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00784D99 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00784E08 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00784E74 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00784E8F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00784F04 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00784F6B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00784FCD | `Photos_Screen` | Known | Screen layout |
| 0x00785031 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078504F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007850C1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007850DE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00785144 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078515F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007851C8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007851E5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078525C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00785280 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007852EE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00785309 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007853C4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007853E0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078544E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078546B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007854D6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007854F6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078556D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00785589 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007855F9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00785618 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00785684 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00785698 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00785711 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00785785 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007857F5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078585D | `NoContent_Screen` | Known | Screen layout |
| 0x00785871 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007858D5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078593C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00785956 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007859C4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00785A36 | `NoContent_Screen` | Known | Screen layout |
| 0x00785A4A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00785AB4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00785B1D | `No_Photos_Screen` | Known | Screen layout |
| 0x00785B31 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00785B97 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00785C05 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00785C72 | `NoContent_Screen` | Known | Screen layout |
| 0x00785C86 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00785CEE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00785D58 | `NoContent_Screen` | Known | Screen layout |
| 0x00785D6C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00785DD3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00785E3D | `NoContent_Screen` | Known | Screen layout |
| 0x00785E51 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00785EBE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00785F30 | `NoContent_Screen` | Known | Screen layout |
| 0x00785F44 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00785FAC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00786015 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00786030 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00786096 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007860B2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00786191 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007861AA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078620B | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078621F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078638D | `Radio_Screen` | Known | Screen layout |
| 0x0078639D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007863FE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00786481 | `LockediPod_Screen` | Known | Screen layout |
| 0x00786509 | `Lock_Screen` | Known | Screen layout |
| 0x00786518 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078657B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007865DD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007865F9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078666B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078668A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007866F2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078670C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00786774 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00786791 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007867FD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00786867 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00786881 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007868F1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00786964 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007869D5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00786A44 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00786AB0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00786ACB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00786B40 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00786BA7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00786C09 | `Photos_Screen` | Known | Screen layout |
| 0x00786C6D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00786C8B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00786CFD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00786D1A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00786D80 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00786D9B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00786E04 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00786E21 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00786E98 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00786EBC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00786F2A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00786F45 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00787000 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078701C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078708A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007870A7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00787112 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00787132 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007871A9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007871C5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00787235 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00787254 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007872C0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007872D4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078734D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007873C1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00787431 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00787499 | `NoContent_Screen` | Known | Screen layout |
| 0x007874AD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00787511 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00787578 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00787592 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00787600 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00787672 | `NoContent_Screen` | Known | Screen layout |
| 0x00787686 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007876F0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00787759 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078776D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007877D3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00787841 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007878AE | `NoContent_Screen` | Known | Screen layout |
| 0x007878C2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078792A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00787994 | `NoContent_Screen` | Known | Screen layout |
| 0x007879A8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00787A0F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00787A79 | `NoContent_Screen` | Known | Screen layout |
| 0x00787A8D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00787AFA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00787B6C | `NoContent_Screen` | Known | Screen layout |
| 0x00787B80 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00787BE8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00787C51 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00787C6C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00787CD2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00787CEE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00787DCD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00787DE6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00787E47 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00787E5B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00787FC9 | `Radio_Screen` | Known | Screen layout |
| 0x00787FD9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078803A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007880BD | `LockediPod_Screen` | Known | Screen layout |
| 0x00788145 | `Lock_Screen` | Known | Screen layout |
| 0x00788154 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007881B7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00788219 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00788235 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007882A7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007882C6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078832E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00788348 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007883B0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007883CD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00788439 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007884A3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007884BD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078852D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007885A0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00788611 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00788680 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007886EC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00788707 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078877C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007887E3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00788845 | `Photos_Screen` | Known | Screen layout |
| 0x007888A9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007888C7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00788939 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00788956 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007889BC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007889D7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00788A40 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00788A5D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00788AD4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00788AF8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00788B66 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00788B81 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00788C3C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00788C58 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00788CC6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00788CE3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00788D4E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00788D6E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00788DE5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00788E01 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00788E71 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00788E90 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00788EFC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00788F10 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00788F89 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00788FFD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078906D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007890D5 | `NoContent_Screen` | Known | Screen layout |
| 0x007890E9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078914D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007891B4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007891CE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078923C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007892AE | `NoContent_Screen` | Known | Screen layout |
| 0x007892C2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078932C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00789395 | `No_Photos_Screen` | Known | Screen layout |
| 0x007893A9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078940F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078947D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007894EA | `NoContent_Screen` | Known | Screen layout |
| 0x007894FE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00789566 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007895D0 | `NoContent_Screen` | Known | Screen layout |
| 0x007895E4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078964B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007896B5 | `NoContent_Screen` | Known | Screen layout |
| 0x007896C9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00789736 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007897A8 | `NoContent_Screen` | Known | Screen layout |
| 0x007897BC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00789824 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078988D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007898A8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078990E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078992A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00789A09 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00789A22 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00789A83 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00789A97 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00789C05 | `Radio_Screen` | Known | Screen layout |
| 0x00789C15 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00789C76 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00789CF9 | `LockediPod_Screen` | Known | Screen layout |
| 0x00789D81 | `Lock_Screen` | Known | Screen layout |
| 0x00789D90 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00789DF3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00789E55 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00789E71 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00789EE3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00789F02 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00789F6A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00789F84 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00789FEC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078A009 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078A075 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078A0DF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078A0F9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078A169 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078A1DC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078A24D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078A2BC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078A328 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078A343 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078A3B8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078A41F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078A481 | `Photos_Screen` | Known | Screen layout |
| 0x0078A4E5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078A503 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078A575 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078A592 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078A5F8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078A613 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078A67C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078A699 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078A710 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078A734 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078A7A2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078A7BD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078A878 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078A894 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078A902 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078A91F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078A98A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078A9AA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078AA21 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078AA3D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078AAAD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078AACC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078AB38 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078AB4C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078ABC5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078AC39 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078ACA9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078AD11 | `NoContent_Screen` | Known | Screen layout |
| 0x0078AD25 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078AD89 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078ADF0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078AE0A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078AE78 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078AEEA | `NoContent_Screen` | Known | Screen layout |
| 0x0078AEFE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078AF68 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078AFD1 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078AFE5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078B04B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078B0B9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078B126 | `NoContent_Screen` | Known | Screen layout |
| 0x0078B13A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078B1A2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078B20C | `NoContent_Screen` | Known | Screen layout |
| 0x0078B220 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078B287 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078B2F1 | `NoContent_Screen` | Known | Screen layout |
| 0x0078B305 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078B372 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078B3E4 | `NoContent_Screen` | Known | Screen layout |
| 0x0078B3F8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078B460 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078B4C9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078B4E4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078B54A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078B566 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078B645 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078B65E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078B6BF | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078B6D3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078B841 | `Radio_Screen` | Known | Screen layout |
| 0x0078B851 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078B8B2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078B935 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078B9BD | `Lock_Screen` | Known | Screen layout |
| 0x0078B9CC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078BA2F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078BA91 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078BAAD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078BB1F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078BB3E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078BBA6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078BBC0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078BC28 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078BC45 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078BCB1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078BD1B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078BD35 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078BDA5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078BE18 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078BE89 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078BEF8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078BF64 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078BF7F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078BFF4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078C05B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078C0BD | `Photos_Screen` | Known | Screen layout |
| 0x0078C121 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078C13F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078C1B1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078C1CE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078C234 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078C24F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078C2B8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078C2D5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078C34C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078C370 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078C3DE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078C3F9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078C4B4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078C4D0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078C53E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078C55B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078C5C6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078C5E6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078C65D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078C679 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078C6E9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078C708 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078C774 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078C788 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078C801 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078C875 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078C8E5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078C94D | `NoContent_Screen` | Known | Screen layout |
| 0x0078C961 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078C9C5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078CA2C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078CA46 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078CAB4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078CB26 | `NoContent_Screen` | Known | Screen layout |
| 0x0078CB3A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078CBA4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078CC0D | `No_Photos_Screen` | Known | Screen layout |
| 0x0078CC21 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078CC87 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078CCF5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078CD62 | `NoContent_Screen` | Known | Screen layout |
| 0x0078CD76 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078CDDE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078CE48 | `NoContent_Screen` | Known | Screen layout |
| 0x0078CE5C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078CEC3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078CF2D | `NoContent_Screen` | Known | Screen layout |
| 0x0078CF41 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078CFAE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078D020 | `NoContent_Screen` | Known | Screen layout |
| 0x0078D034 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078D09C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078D105 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078D120 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078D186 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078D1A2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078D281 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078D29A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078D2FB | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078D30F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078D47D | `Radio_Screen` | Known | Screen layout |
| 0x0078D48D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078D4EE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078D571 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078D5F9 | `Lock_Screen` | Known | Screen layout |
| 0x0078D608 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078D66B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078D6CD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078D6E9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078D75B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078D77A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078D7E2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078D7FC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078D864 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078D881 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078D8ED | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078D957 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078D971 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078D9E1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078DA54 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078DAC5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078DB34 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078DBA0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078DBBB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078DC30 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078DC97 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078DCF9 | `Photos_Screen` | Known | Screen layout |
| 0x0078DD5D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078DD7B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078DDED | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078DE0A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078DE70 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078DE8B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078DEF4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078DF11 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078DF88 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078DFAC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078E01A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078E035 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078E0F0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078E10C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078E17A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078E197 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078E202 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078E222 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078E299 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078E2B5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078E325 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078E344 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078E3B0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078E3C4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078E43D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078E4B1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078E521 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078E589 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E59D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078E601 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078E668 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078E682 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078E6F0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078E762 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E776 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078E7E0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078E849 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078E85D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078E8C3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078E931 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078E99E | `NoContent_Screen` | Known | Screen layout |
| 0x0078E9B2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078EA1A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078EA84 | `NoContent_Screen` | Known | Screen layout |
| 0x0078EA98 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078EAFF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078EB69 | `NoContent_Screen` | Known | Screen layout |
| 0x0078EB7D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078EBEA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078EC5C | `NoContent_Screen` | Known | Screen layout |
| 0x0078EC70 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078ECD8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078ED41 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078ED5C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078EDC2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078EDDE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078EEBD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078EED6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078EF37 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078EF4B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078F0B9 | `Radio_Screen` | Known | Screen layout |
| 0x0078F0C9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078F12A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078F1AD | `LockediPod_Screen` | Known | Screen layout |
| 0x0078F235 | `Lock_Screen` | Known | Screen layout |
| 0x0078F244 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078F2A7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078F309 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078F325 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078F397 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078F3B6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078F41E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078F438 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078F4A0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078F4BD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078F529 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078F593 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078F5AD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078F61D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078F690 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078F701 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078F770 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078F7DC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078F7F7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078F86C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078F8D3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078F935 | `Photos_Screen` | Known | Screen layout |
| 0x0078F999 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078F9B7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078FA29 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078FA46 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078FAAC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078FAC7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078FB30 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078FB4D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078FBC4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078FBE8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078FC56 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078FC71 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078FD2C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078FD48 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078FDB6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078FDD3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078FE3E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078FE5E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078FED5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078FEF1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078FF61 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078FF80 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078FFEC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00790000 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00790079 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007900ED | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079015D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007901C5 | `NoContent_Screen` | Known | Screen layout |
| 0x007901D9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079023D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007902A4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007902BE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079032C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079039E | `NoContent_Screen` | Known | Screen layout |
| 0x007903B2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079041C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00790485 | `No_Photos_Screen` | Known | Screen layout |
| 0x00790499 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007904FF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079056D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007905DA | `NoContent_Screen` | Known | Screen layout |
| 0x007905EE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00790656 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007906C0 | `NoContent_Screen` | Known | Screen layout |
| 0x007906D4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079073B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007907A5 | `NoContent_Screen` | Known | Screen layout |
| 0x007907B9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00790826 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00790898 | `NoContent_Screen` | Known | Screen layout |
| 0x007908AC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00790914 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0079097D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00790998 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007909FE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00790A1A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00790AF9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00790B12 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00790B73 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00790B87 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00790CF5 | `Radio_Screen` | Known | Screen layout |
| 0x00790D05 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00790D66 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00790DE9 | `LockediPod_Screen` | Known | Screen layout |
| 0x00790E71 | `Lock_Screen` | Known | Screen layout |
| 0x00790E80 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00790EE3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00790F45 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00790F61 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00790FD3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00790FF2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079105A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00791074 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007910DC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007910F9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00791165 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007911CF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007911E9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00791259 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007912CC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079133D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007913AC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00791418 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00791433 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007914A8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079150F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00791571 | `Photos_Screen` | Known | Screen layout |
| 0x007915D5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007915F3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00791665 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00791682 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007916E8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00791703 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079176C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00791789 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00791800 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00791824 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00791892 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007918AD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00791968 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00791984 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007919F2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00791A0F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00791A7A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00791A9A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00791B11 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00791B2D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00791B9D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00791BBC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00791C28 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00791C3C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00791CB5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00791D29 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00791D99 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00791E01 | `NoContent_Screen` | Known | Screen layout |
| 0x00791E15 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00791E79 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00791EE0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00791EFA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00791F68 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00791FDA | `NoContent_Screen` | Known | Screen layout |
| 0x00791FEE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00792058 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007920C1 | `No_Photos_Screen` | Known | Screen layout |
| 0x007920D5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079213B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007921A9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00792216 | `NoContent_Screen` | Known | Screen layout |
| 0x0079222A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00792292 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007922FC | `NoContent_Screen` | Known | Screen layout |
| 0x00792310 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00792377 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007923E1 | `NoContent_Screen` | Known | Screen layout |
| 0x007923F5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00792462 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007924D4 | `NoContent_Screen` | Known | Screen layout |
| 0x007924E8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00792550 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007925B9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007925D4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079263A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00792656 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00792735 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079274E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007927AF | `FirstBoot_Screen` | Known | Screen layout |
| 0x007927C3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00792931 | `Radio_Screen` | Known | Screen layout |
| 0x00792941 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007929A2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00792A25 | `LockediPod_Screen` | Known | Screen layout |
| 0x00792AAD | `Lock_Screen` | Known | Screen layout |
| 0x00792ABC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00792B1F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00792B81 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00792B9D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00792C0F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00792C2E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00792C96 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00792CB0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00792D18 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00792D35 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00792DA1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00792E0B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00792E25 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00792E95 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00792F08 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00792F79 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00792FE8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00793054 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079306F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007930E4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079314B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007931AD | `Photos_Screen` | Known | Screen layout |
| 0x00793211 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079322F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007932A1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007932BE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00793324 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079333F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007933A8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007933C5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079343C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00793460 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007934CE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007934E9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007935A4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007935C0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079362E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079364B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007936B6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007936D6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079374D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00793769 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007937D9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007937F8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00793864 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00793878 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007938F1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00793965 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007939D5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00793A3D | `NoContent_Screen` | Known | Screen layout |
| 0x00793A51 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00793AB5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00793B1C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00793B36 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00793BA4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00793C16 | `NoContent_Screen` | Known | Screen layout |
| 0x00793C2A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00793C94 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00793CFD | `No_Photos_Screen` | Known | Screen layout |
| 0x00793D11 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00793D77 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00793DE5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00793E52 | `NoContent_Screen` | Known | Screen layout |
| 0x00793E66 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00793ECE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00793F38 | `NoContent_Screen` | Known | Screen layout |
| 0x00793F4C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00793FB3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079401D | `NoContent_Screen` | Known | Screen layout |
| 0x00794031 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079409E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00794110 | `NoContent_Screen` | Known | Screen layout |
| 0x00794124 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079418C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007941F5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00794210 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00794276 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00794292 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00794371 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079438A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007943EB | `FirstBoot_Screen` | Known | Screen layout |
| 0x007943FF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0079456D | `Radio_Screen` | Known | Screen layout |
| 0x0079457D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007945DE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00794661 | `LockediPod_Screen` | Known | Screen layout |
| 0x007946E9 | `Lock_Screen` | Known | Screen layout |
| 0x007946F8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079475B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007947BD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007947D9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0079484B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079486A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007948D2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007948EC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00794954 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00794971 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007949DD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00794A47 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00794A61 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00794AD1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00794B44 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00794BB5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00794C24 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00794C90 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00794CAB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00794D20 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00794D87 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00794DE9 | `Photos_Screen` | Known | Screen layout |
| 0x00794E4D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00794E6B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00794EDD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00794EFA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00794F60 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00794F7B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00794FE4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00795001 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00795078 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079509C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079510A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00795125 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007951E0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007951FC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079526A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00795287 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007952F2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00795312 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00795389 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007953A5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00795415 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00795434 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007954A0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007954B4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0079552D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007955A1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00795611 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00795679 | `NoContent_Screen` | Known | Screen layout |
| 0x0079568D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007956F1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00795758 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00795772 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007957E0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00795852 | `NoContent_Screen` | Known | Screen layout |
| 0x00795866 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007958D0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00795939 | `No_Photos_Screen` | Known | Screen layout |
| 0x0079594D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007959B3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00795A21 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00795A8E | `NoContent_Screen` | Known | Screen layout |
| 0x00795AA2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00795B0A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00795B74 | `NoContent_Screen` | Known | Screen layout |
| 0x00795B88 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00795BEF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00795C59 | `NoContent_Screen` | Known | Screen layout |
| 0x00795C6D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00795CDA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00795D4C | `NoContent_Screen` | Known | Screen layout |
| 0x00795D60 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00795DC8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00795E31 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00795E4C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00795EB2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00795ECE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00795FAD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00795FC6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00796027 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079603B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007961A9 | `Radio_Screen` | Known | Screen layout |
| 0x007961B9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079621A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079629D | `LockediPod_Screen` | Known | Screen layout |
| 0x00796325 | `Lock_Screen` | Known | Screen layout |
| 0x00796334 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00796397 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007963F9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00796415 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00796487 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007964A6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079650E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00796528 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00796590 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007965AD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00796619 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00796683 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0079669D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0079670D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00796780 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007967F1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00796860 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007968CC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007968E7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0079695C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007969C3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00796A25 | `Photos_Screen` | Known | Screen layout |
| 0x00796A89 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00796AA7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00796B19 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00796B36 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00796B9C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00796BB7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00796C20 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00796C3D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00796CB4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00796CD8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00796D46 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00796D61 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00796E1C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00796E38 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00796EA6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00796EC3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00796F2E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00796F4E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00796FC5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00796FE1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00797051 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00797070 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007970DC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007970F0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00797169 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007971DD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079724D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007972B5 | `NoContent_Screen` | Known | Screen layout |
| 0x007972C9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079732D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00797394 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007973AE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079741C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079748E | `NoContent_Screen` | Known | Screen layout |
| 0x007974A2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079750C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00797575 | `No_Photos_Screen` | Known | Screen layout |
| 0x00797589 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007975EF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079765D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007976CA | `NoContent_Screen` | Known | Screen layout |
| 0x007976DE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00797746 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007977B0 | `NoContent_Screen` | Known | Screen layout |
| 0x007977C4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079782B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00797895 | `NoContent_Screen` | Known | Screen layout |
| 0x007978A9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00797916 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00797988 | `NoContent_Screen` | Known | Screen layout |
| 0x0079799C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00797A04 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00797A6D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00797A88 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00797AEE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00797B0A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00797BE9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00797C02 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00797C63 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00797C77 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00797DE5 | `Radio_Screen` | Known | Screen layout |
| 0x00797DF5 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00797E56 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00797ED9 | `LockediPod_Screen` | Known | Screen layout |
| 0x00797F61 | `Lock_Screen` | Known | Screen layout |
| 0x00797F70 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00797FD3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00798035 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00798051 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007980C3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007980E2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079814A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00798164 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007981CC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007981E9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00798255 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007982BF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007982D9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00798349 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007983BC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079842D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079849C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00798508 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00798523 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00798598 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007985FF | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00798661 | `Photos_Screen` | Known | Screen layout |
| 0x007986C5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007986E3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00798755 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00798772 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007987D8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007987F3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079885C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00798879 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007988F0 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00798914 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00798982 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0079899D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00798A58 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00798A74 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00798AE2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00798AFF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00798B6A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00798B8A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00798C01 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00798C1D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00798C8D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00798CAC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00798D18 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00798D2C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00798DA5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00798E19 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00798E89 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00798EF1 | `NoContent_Screen` | Known | Screen layout |
| 0x00798F05 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00798F69 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00798FD0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00798FEA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00799058 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007990CA | `NoContent_Screen` | Known | Screen layout |
| 0x007990DE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00799148 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007991B1 | `No_Photos_Screen` | Known | Screen layout |
| 0x007991C5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079922B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00799299 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00799306 | `NoContent_Screen` | Known | Screen layout |
| 0x0079931A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00799382 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007993EC | `NoContent_Screen` | Known | Screen layout |
| 0x00799400 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00799467 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007994D1 | `NoContent_Screen` | Known | Screen layout |
| 0x007994E5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00799552 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007995C4 | `NoContent_Screen` | Known | Screen layout |
| 0x007995D8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00799640 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007996A9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007996C4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079972A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00799746 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00799825 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079983E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0079989F | `FirstBoot_Screen` | Known | Screen layout |
| 0x007998B3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00799A21 | `Radio_Screen` | Known | Screen layout |
| 0x00799A31 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00799A92 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00799B15 | `LockediPod_Screen` | Known | Screen layout |
| 0x00799B9D | `Lock_Screen` | Known | Screen layout |
| 0x00799BAC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00799C0F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00799C71 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00799C8D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00799CFF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00799D1E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00799D86 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00799DA0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00799E08 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00799E25 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00799E91 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00799EFB | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00799F15 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00799F85 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00799FF8 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079A069 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079A0D8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079A144 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079A15F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0079A1D4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079A23B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079A29D | `Photos_Screen` | Known | Screen layout |
| 0x0079A301 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079A31F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0079A391 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079A3AE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079A414 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079A42F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079A498 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079A4B5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079A52C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079A550 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079A5BE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0079A5D9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079A694 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079A6B0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079A71E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079A73B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079A7A6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079A7C6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079A83D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079A859 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079A8C9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079A8E8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079A954 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079A968 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0079A9E1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079AA55 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079AAC5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079AB2D | `NoContent_Screen` | Known | Screen layout |
| 0x0079AB41 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079ABA5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0079AC0C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079AC26 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079AC94 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079AD06 | `NoContent_Screen` | Known | Screen layout |
| 0x0079AD1A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079AD84 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0079ADED | `No_Photos_Screen` | Known | Screen layout |
| 0x0079AE01 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079AE67 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079AED5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0079AF42 | `NoContent_Screen` | Known | Screen layout |
| 0x0079AF56 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079AFBE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0079B028 | `NoContent_Screen` | Known | Screen layout |
| 0x0079B03C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079B0A3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079B10D | `NoContent_Screen` | Known | Screen layout |
| 0x0079B121 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079B18E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079B200 | `NoContent_Screen` | Known | Screen layout |
| 0x0079B214 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079B27C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0079B2E5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0079B300 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079B366 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079B382 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079B461 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079B47A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0079B4DB | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079B4EF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0079B65D | `Radio_Screen` | Known | Screen layout |
| 0x0079B66D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079B6CE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079B751 | `LockediPod_Screen` | Known | Screen layout |
| 0x0079B7D9 | `Lock_Screen` | Known | Screen layout |
| 0x0079B7E8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079B84B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0079B8AD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079B8C9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0079B93B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079B95A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079B9C2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079B9DC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079BA44 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079BA61 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079BACD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079BB37 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0079BB51 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0079BBC1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079BC34 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079BCA5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079BD14 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079BD80 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079BD9B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0079BE10 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079BE77 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079BED9 | `Photos_Screen` | Known | Screen layout |
| 0x0079BF3D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079BF5B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0079BFCD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079BFEA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079C050 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079C06B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079C0D4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079C0F1 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079C168 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079C18C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079C1FA | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0079C215 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079C2D0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079C2EC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079C35A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079C377 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079C3E2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079C402 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079C479 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079C495 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079C505 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079C524 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079C590 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079C5A4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0079C61D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079C691 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079C701 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079C769 | `NoContent_Screen` | Known | Screen layout |
| 0x0079C77D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079C7E1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0079C848 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079C862 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079C8D0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079C942 | `NoContent_Screen` | Known | Screen layout |
| 0x0079C956 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079C9C0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0079CA29 | `No_Photos_Screen` | Known | Screen layout |
| 0x0079CA3D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079CAA3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079CB11 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0079CB7E | `NoContent_Screen` | Known | Screen layout |
| 0x0079CB92 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079CBFA | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0079CC64 | `NoContent_Screen` | Known | Screen layout |
| 0x0079CC78 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079CCDF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079CD49 | `NoContent_Screen` | Known | Screen layout |
| 0x0079CD5D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079CDCA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079CE3C | `NoContent_Screen` | Known | Screen layout |
| 0x0079CE50 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079CEB8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0079CF21 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0079CF3C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079CFA2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079CFBE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079D09D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079D0B6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0079D117 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079D12B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0079D299 | `Radio_Screen` | Known | Screen layout |
| 0x0079D2A9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079D30A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079D38D | `LockediPod_Screen` | Known | Screen layout |
| 0x0079D415 | `Lock_Screen` | Known | Screen layout |
| 0x0079D424 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079D487 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0079D4E9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079D505 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0079D577 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079D596 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079D5FE | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079D618 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079D680 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079D69D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079D709 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079D773 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0079D78D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0079D7FD | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079D870 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079D8E1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079D950 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079D9BC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079D9D7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0079DA4C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079DAB3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079DB15 | `Photos_Screen` | Known | Screen layout |
| 0x0079DB79 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079DB97 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0079DC09 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079DC26 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079DC8C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079DCA7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079DD10 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079DD2D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079DDA4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079DDC8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079DE36 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0079DE51 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079DF0C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079DF28 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079DF96 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079DFB3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079E01E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079E03E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079E0B5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079E0D1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079E141 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079E160 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079E1CC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079E1E0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0079E259 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079E2CD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079E33D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079E3A5 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E3B9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079E41D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0079E484 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079E49E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079E50C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079E57E | `NoContent_Screen` | Known | Screen layout |
| 0x0079E592 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079E5FC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0079E665 | `No_Photos_Screen` | Known | Screen layout |
| 0x0079E679 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079E6DF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079E74D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0079E7BA | `NoContent_Screen` | Known | Screen layout |
| 0x0079E7CE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079E836 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0079E8A0 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E8B4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079E91B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079E985 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E999 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079EA06 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079EA78 | `NoContent_Screen` | Known | Screen layout |
| 0x0079EA8C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079EAF4 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0079EB5D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0079EB78 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079EBDE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079EBFA | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079ECD9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079ECF2 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0079ED53 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079ED67 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0079EED5 | `Radio_Screen` | Known | Screen layout |
| 0x0079EEE5 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079EF46 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079EFC9 | `LockediPod_Screen` | Known | Screen layout |
| 0x0079F051 | `Lock_Screen` | Known | Screen layout |
| 0x0079F060 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079F0C3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0079F125 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079F141 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0079F1B3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079F1D2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079F23A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079F254 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079F2BC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079F2D9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079F345 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079F3AF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0079F3C9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0079F439 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079F4AC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079F51D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079F58C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079F5F8 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079F613 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0079F688 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079F6EF | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079F751 | `Photos_Screen` | Known | Screen layout |
| 0x0079F7B5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079F7D3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0079F845 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079F862 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079F8C8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079F8E3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079F94C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079F969 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079F9E0 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079FA04 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079FA72 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0079FA8D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079FB48 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079FB64 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079FBD2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079FBEF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079FC5A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079FC7A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079FCF1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079FD0D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079FD7D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079FD9C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079FE08 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079FE1C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0079FE95 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079FF09 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079FF79 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079FFE1 | `NoContent_Screen` | Known | Screen layout |
| 0x0079FFF5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A0059 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A00C0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A00DA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A0148 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A01BA | `NoContent_Screen` | Known | Screen layout |
| 0x007A01CE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A0238 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A02A1 | `No_Photos_Screen` | Known | Screen layout |
| 0x007A02B5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A031B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A0389 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A03F6 | `NoContent_Screen` | Known | Screen layout |
| 0x007A040A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A0472 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A04DC | `NoContent_Screen` | Known | Screen layout |
| 0x007A04F0 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A0557 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A05C1 | `NoContent_Screen` | Known | Screen layout |
| 0x007A05D5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A0642 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A06B4 | `NoContent_Screen` | Known | Screen layout |
| 0x007A06C8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A0730 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A0799 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A07B4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A081A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A0836 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A0915 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A092E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A098F | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A09A3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A0B11 | `Radio_Screen` | Known | Screen layout |
| 0x007A0B21 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A0B82 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A0C05 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A0C8D | `Lock_Screen` | Known | Screen layout |
| 0x007A0C9C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A0CFF | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A0D61 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A0D7D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A0DEF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A0E0E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A0E76 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A0E90 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A0EF8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A0F15 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A0F81 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A0FEB | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A1005 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A1075 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A10E8 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A1159 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A11C8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A1234 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A124F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A12C4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A132B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A138D | `Photos_Screen` | Known | Screen layout |
| 0x007A13F1 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A140F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A1481 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A149E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A1504 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A151F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A1588 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A15A5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A161C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A1640 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A16AE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A16C9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A1784 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A17A0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A180E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A182B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A1896 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A18B6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A192D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A1949 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A19B9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A19D8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A1A44 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A1A58 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A1AD1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A1B45 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A1BB5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A1C1D | `NoContent_Screen` | Known | Screen layout |
| 0x007A1C31 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A1C95 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A1CFC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A1D16 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A1D84 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A1DF6 | `NoContent_Screen` | Known | Screen layout |
| 0x007A1E0A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A1E74 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A1EDD | `No_Photos_Screen` | Known | Screen layout |
| 0x007A1EF1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A1F57 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A1FC5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A2032 | `NoContent_Screen` | Known | Screen layout |
| 0x007A2046 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A20AE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A2118 | `NoContent_Screen` | Known | Screen layout |
| 0x007A212C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A2193 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A21FD | `NoContent_Screen` | Known | Screen layout |
| 0x007A2211 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A227E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A22F0 | `NoContent_Screen` | Known | Screen layout |
| 0x007A2304 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A236C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A23D5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A23F0 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A2456 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A2472 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A2551 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A256A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A25CB | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A25DF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A274D | `Radio_Screen` | Known | Screen layout |
| 0x007A275D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A27BE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A2841 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A28C9 | `Lock_Screen` | Known | Screen layout |
| 0x007A28D8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A293B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A299D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A29B9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A2A2B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A2A4A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A2AB2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A2ACC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A2B34 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A2B51 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A2BBD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A2C27 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A2C41 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A2CB1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A2D24 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A2D95 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A2E04 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A2E70 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A2E8B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A2F00 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A2F67 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A2FC9 | `Photos_Screen` | Known | Screen layout |
| 0x007A302D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A304B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A30BD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A30DA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A3140 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A315B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A31C4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A31E1 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A3258 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A327C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A32EA | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A3305 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A33C0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A33DC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A344A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A3467 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A34D2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A34F2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A3569 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A3585 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A35F5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A3614 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A3680 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A3694 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A370D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A3781 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A37F1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A3859 | `NoContent_Screen` | Known | Screen layout |
| 0x007A386D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A38D1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A3938 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A3952 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A39C0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A3A32 | `NoContent_Screen` | Known | Screen layout |
| 0x007A3A46 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A3AB0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A3B19 | `No_Photos_Screen` | Known | Screen layout |
| 0x007A3B2D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A3B93 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A3C01 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A3C6E | `NoContent_Screen` | Known | Screen layout |
| 0x007A3C82 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A3CEA | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A3D54 | `NoContent_Screen` | Known | Screen layout |
| 0x007A3D68 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A3DCF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A3E39 | `NoContent_Screen` | Known | Screen layout |
| 0x007A3E4D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A3EBA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A3F2C | `NoContent_Screen` | Known | Screen layout |
| 0x007A3F40 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A3FA8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A4011 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A402C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A4092 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A40AE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A418D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A41A6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A4207 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A421B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A4389 | `Radio_Screen` | Known | Screen layout |
| 0x007A4399 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A43FA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A447D | `LockediPod_Screen` | Known | Screen layout |
| 0x007A4505 | `Lock_Screen` | Known | Screen layout |
| 0x007A4514 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A4577 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A45D9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A45F5 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A4667 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A4686 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A46EE | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A4708 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A4770 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A478D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A47F9 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A4863 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A487D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A48ED | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A4960 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A49D1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A4A40 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A4AAC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A4AC7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A4B3C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A4BA3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A4C05 | `Photos_Screen` | Known | Screen layout |
| 0x007A4C69 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A4C87 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A4CF9 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A4D16 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A4D7C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A4D97 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A4E00 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A4E1D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A4E94 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A4EB8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A4F26 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A4F41 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A4FFC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A5018 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A5086 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A50A3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A510E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A512E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A51A5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A51C1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A5231 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A5250 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A52BC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A52D0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A5349 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A53BD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A542D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A5495 | `NoContent_Screen` | Known | Screen layout |
| 0x007A54A9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A550D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A5574 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A558E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A55FC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A566E | `NoContent_Screen` | Known | Screen layout |
| 0x007A5682 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A56EC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A5755 | `No_Photos_Screen` | Known | Screen layout |
| 0x007A5769 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A57CF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A583D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A58AA | `NoContent_Screen` | Known | Screen layout |
| 0x007A58BE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A5926 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A5990 | `NoContent_Screen` | Known | Screen layout |
| 0x007A59A4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A5A0B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A5A75 | `NoContent_Screen` | Known | Screen layout |
| 0x007A5A89 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A5AF6 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A5B68 | `NoContent_Screen` | Known | Screen layout |
| 0x007A5B7C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A5BE4 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A5C4D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A5C68 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A5CCE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A5CEA | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A5DC9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A5DE2 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A5E43 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A5E57 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A5FC5 | `Radio_Screen` | Known | Screen layout |
| 0x007A5FD5 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A6036 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A60B9 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A6141 | `Lock_Screen` | Known | Screen layout |
| 0x007A6150 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A61B3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A6215 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A6231 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A62A3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A62C2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A632A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A6344 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A63AC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A63C9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A6435 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A649F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A64B9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A6529 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A659C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A660D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A667C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A66E8 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A6703 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A6778 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A67DF | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A6841 | `Photos_Screen` | Known | Screen layout |
| 0x007A68A5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A68C3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A6935 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A6952 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A69B8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A69D3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A6A3C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A6A59 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A6AD0 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A6AF4 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A6B62 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A6B7D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A6C38 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A6C54 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A6CC2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A6CDF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A6D4A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A6D6A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A6DE1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A6DFD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A6E6D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A6E8C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A6EF8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A6F0C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A6F85 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A6FF9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A7069 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A70D1 | `NoContent_Screen` | Known | Screen layout |
| 0x007A70E5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A7149 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A71B0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A71CA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A7238 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A72AA | `NoContent_Screen` | Known | Screen layout |
| 0x007A72BE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A7328 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A7391 | `No_Photos_Screen` | Known | Screen layout |
| 0x007A73A5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A740B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A7479 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A74E6 | `NoContent_Screen` | Known | Screen layout |
| 0x007A74FA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A7562 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A75CC | `NoContent_Screen` | Known | Screen layout |
| 0x007A75E0 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A7647 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A76B1 | `NoContent_Screen` | Known | Screen layout |
| 0x007A76C5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A7732 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A77A4 | `NoContent_Screen` | Known | Screen layout |
| 0x007A77B8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A7820 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A7889 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A78A4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A790A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A7926 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A7A05 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A7A1E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A7A7F | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A7A93 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A7C01 | `Radio_Screen` | Known | Screen layout |
| 0x007A7C11 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A7C72 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A7CF5 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A7D7D | `Lock_Screen` | Known | Screen layout |
| 0x007A7D8C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A7DEF | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A7E51 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A7E6D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A7EDF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A7EFE | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A7F66 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A7F80 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A7FE8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A8005 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A8071 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A80DB | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A80F5 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A8165 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A81D8 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A8249 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A82B8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A8324 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A833F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A83B4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A841B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A847D | `Photos_Screen` | Known | Screen layout |
| 0x007A84E1 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A84FF | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A8571 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A858E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A85F4 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A860F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A8678 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A8695 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A870C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A8730 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A879E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A87B9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A8874 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A8890 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A88FE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A891B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A8986 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A89A6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A8A1D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A8A39 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A8AA9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A8AC8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A8B34 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A8B48 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A8BC1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A8C35 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A8CA5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A8D0D | `NoContent_Screen` | Known | Screen layout |
| 0x007A8D21 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A8D85 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A8DEC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A8E06 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A8E74 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A8EE6 | `NoContent_Screen` | Known | Screen layout |
| 0x007A8EFA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A8F64 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A8FCD | `No_Photos_Screen` | Known | Screen layout |
| 0x007A8FE1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A9047 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A90B5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A9122 | `NoContent_Screen` | Known | Screen layout |
| 0x007A9136 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A919E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A9208 | `NoContent_Screen` | Known | Screen layout |
| 0x007A921C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A9283 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A92ED | `NoContent_Screen` | Known | Screen layout |
| 0x007A9301 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A936E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A93E0 | `NoContent_Screen` | Known | Screen layout |
| 0x007A93F4 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A945C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A94C5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A94E0 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A9546 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A9562 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A9641 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A965A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A96BB | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A96CF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A983D | `Radio_Screen` | Known | Screen layout |
| 0x007A984D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A98AE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A9931 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A99B9 | `Lock_Screen` | Known | Screen layout |
| 0x007A99C8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A9A2B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A9A8D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A9AA9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A9B1B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A9B3A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A9BA2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A9BBC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A9C24 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A9C41 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A9CAD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A9D17 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A9D31 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A9DA1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A9E14 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A9E85 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A9EF4 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A9F60 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A9F7B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A9FF0 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007AA057 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007AA0B9 | `Photos_Screen` | Known | Screen layout |
| 0x007AA11D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007AA13B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007AA1AD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007AA1CA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007AA230 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007AA24B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007AA2B4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007AA2D1 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007AA348 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007AA36C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007AA3DA | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007AA3F5 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007AA4B0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AA4CC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AA53A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AA557 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AA5C2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007AA5E2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007AA659 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AA675 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AA6E5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007AA704 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007AA770 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007AA784 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007AA7FD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007AA871 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007AA8E1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007AA949 | `NoContent_Screen` | Known | Screen layout |
| 0x007AA95D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007AA9C1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007AAA28 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AAA42 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007AAAB0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007AAB22 | `NoContent_Screen` | Known | Screen layout |
| 0x007AAB36 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007AABA0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007AAC09 | `No_Photos_Screen` | Known | Screen layout |
| 0x007AAC1D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007AAC83 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AACF1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007AAD5E | `NoContent_Screen` | Known | Screen layout |
| 0x007AAD72 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007AADDA | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007AAE44 | `NoContent_Screen` | Known | Screen layout |
| 0x007AAE58 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007AAEBF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007AAF29 | `NoContent_Screen` | Known | Screen layout |
| 0x007AAF3D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007AAFAA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007AB01C | `NoContent_Screen` | Known | Screen layout |
| 0x007AB030 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AB098 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007AB101 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007AB11C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007AB182 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007AB19E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007AB27D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007AB296 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007AB2F7 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007AB30B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007AB479 | `Radio_Screen` | Known | Screen layout |
| 0x007AB489 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007AB4EA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007AB56D | `LockediPod_Screen` | Known | Screen layout |
| 0x007AB5F5 | `Lock_Screen` | Known | Screen layout |
| 0x007AB604 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007AB667 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007AB6C9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007AB6E5 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007AB757 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007AB776 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007AB7DE | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AB7F8 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007AB860 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AB87D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AB8E9 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007AB953 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007AB96D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007AB9DD | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007ABA50 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007ABAC1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007ABB30 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007ABB9C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007ABBB7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007ABC2C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007ABC93 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007ABCF5 | `Photos_Screen` | Known | Screen layout |
| 0x007ABD59 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007ABD77 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007ABDE9 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007ABE06 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007ABE6C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007ABE87 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007ABEF0 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007ABF0D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007ABF84 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007ABFA8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007AC016 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007AC031 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007AC0EC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AC108 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AC176 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AC193 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AC1FE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007AC21E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007AC295 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AC2B1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AC321 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007AC340 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007AC3AC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007AC3C0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007AC439 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007AC4AD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007AC51D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007AC585 | `NoContent_Screen` | Known | Screen layout |
| 0x007AC599 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007AC5FD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007AC664 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AC67E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007AC6EC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007AC75E | `NoContent_Screen` | Known | Screen layout |
| 0x007AC772 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007AC7DC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007AC845 | `No_Photos_Screen` | Known | Screen layout |
| 0x007AC859 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007AC8BF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AC92D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007AC99A | `NoContent_Screen` | Known | Screen layout |
| 0x007AC9AE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007ACA16 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007ACA80 | `NoContent_Screen` | Known | Screen layout |
| 0x007ACA94 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007ACAFB | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007ACB65 | `NoContent_Screen` | Known | Screen layout |
| 0x007ACB79 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007ACBE6 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007ACC58 | `NoContent_Screen` | Known | Screen layout |
| 0x007ACC6C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007ACCD4 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007ACD3D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007ACD58 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007ACDBE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007ACDDA | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007ACEB9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007ACED2 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007ACF33 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007ACF47 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007AD0B5 | `Radio_Screen` | Known | Screen layout |
| 0x007AD0C5 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007AD126 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007AD1A9 | `LockediPod_Screen` | Known | Screen layout |
| 0x007AD231 | `Lock_Screen` | Known | Screen layout |
| 0x007AD240 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007AD2A3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007AD305 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007AD321 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007AD393 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007AD3B2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007AD41A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AD434 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007AD49C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AD4B9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AD525 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007AD58F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007AD5A9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007AD619 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007AD68C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007AD6FD | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007AD76C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007AD7D8 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007AD7F3 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007AD868 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007AD8CF | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007AD931 | `Photos_Screen` | Known | Screen layout |
| 0x007AD995 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007AD9B3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007ADA25 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007ADA42 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007ADAA8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007ADAC3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007ADB2C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007ADB49 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007ADBC0 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007ADBE4 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007ADC52 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007ADC6D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007ADD28 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007ADD44 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007ADDB2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007ADDCF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007ADE3A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007ADE5A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007ADED1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007ADEED | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007ADF5D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007ADF7C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007ADFE8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007ADFFC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007AE075 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007AE0E9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007AE159 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007AE1C1 | `NoContent_Screen` | Known | Screen layout |
| 0x007AE1D5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007AE239 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007AE2A0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AE2BA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007AE328 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007AE39A | `NoContent_Screen` | Known | Screen layout |
| 0x007AE3AE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007AE418 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007AE481 | `No_Photos_Screen` | Known | Screen layout |
| 0x007AE495 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007AE4FB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AE569 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007AE5D6 | `NoContent_Screen` | Known | Screen layout |
| 0x007AE5EA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007AE652 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007AE6BC | `NoContent_Screen` | Known | Screen layout |
| 0x007AE6D0 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007AE737 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007AE7A1 | `NoContent_Screen` | Known | Screen layout |
| 0x007AE7B5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007AE822 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007AE894 | `NoContent_Screen` | Known | Screen layout |
| 0x007AE8A8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AE910 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007AE979 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007AE994 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007AE9FA | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007AEA16 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007AEAF5 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007AEB0E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007AEB6F | `FirstBoot_Screen` | Known | Screen layout |
| 0x007AEB83 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007AECF1 | `Radio_Screen` | Known | Screen layout |
| 0x007AED01 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007AED62 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007AEDE5 | `LockediPod_Screen` | Known | Screen layout |
| 0x007AEE6D | `Lock_Screen` | Known | Screen layout |
| 0x007AEE7C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007AEEDF | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007AEF41 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007AEF5D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007AEFCF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007AEFEE | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007AF056 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AF070 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007AF0D8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AF0F5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AF161 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007AF1CB | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007AF1E5 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007AF255 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007AF2C8 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007AF339 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007AF3A8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007AF414 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007AF42F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007AF4A4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007AF50B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007AF56D | `Photos_Screen` | Known | Screen layout |
| 0x007AF5D1 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007AF5EF | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007AF661 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007AF67E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007AF6E4 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007AF6FF | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007AF768 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007AF785 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007AF7FC | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007AF820 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007AF88E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007AF8A9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007AF964 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AF980 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AF9EE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AFA0B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AFA76 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007AFA96 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007AFB0D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AFB29 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AFB99 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007AFBB8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007AFC24 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007AFC38 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007AFCB1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007AFD25 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007AFD95 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007AFDFD | `NoContent_Screen` | Known | Screen layout |
| 0x007AFE11 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007AFE75 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007AFEDC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AFEF6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007AFF64 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007AFFD6 | `NoContent_Screen` | Known | Screen layout |
| 0x007AFFEA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B0054 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B00BD | `No_Photos_Screen` | Known | Screen layout |
| 0x007B00D1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B0137 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B01A5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B0212 | `NoContent_Screen` | Known | Screen layout |
| 0x007B0226 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B028E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B02F8 | `NoContent_Screen` | Known | Screen layout |
| 0x007B030C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B0373 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B03DD | `NoContent_Screen` | Known | Screen layout |
| 0x007B03F1 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B045E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B04D0 | `NoContent_Screen` | Known | Screen layout |
| 0x007B04E4 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B054C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B05B5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B05D0 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B0636 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B0652 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B0731 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B074A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B07AB | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B07BF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B092D | `Radio_Screen` | Known | Screen layout |
| 0x007B093D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B099E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B0A21 | `LockediPod_Screen` | Known | Screen layout |
| 0x007B0AA9 | `Lock_Screen` | Known | Screen layout |
| 0x007B0AB8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B0B1B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B0B7D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B0B99 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B0C0B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B0C2A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B0C92 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B0CAC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B0D14 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B0D31 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B0D9D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B0E07 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B0E21 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B0E91 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B0F04 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B0F75 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B0FE4 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B1050 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B106B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B10E0 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B1147 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B11A9 | `Photos_Screen` | Known | Screen layout |
| 0x007B120D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B122B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B129D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B12BA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B1320 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B133B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B13A4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B13C1 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B1438 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B145C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B14CA | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B14E5 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B15A0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B15BC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B162A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B1647 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B16B2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B16D2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B1749 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B1765 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B17D5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B17F4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B1860 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B1874 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B18ED | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B1961 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B19D1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B1A39 | `NoContent_Screen` | Known | Screen layout |
| 0x007B1A4D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B1AB1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B1B18 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B1B32 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B1BA0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B1C12 | `NoContent_Screen` | Known | Screen layout |
| 0x007B1C26 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B1C90 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B1CF9 | `No_Photos_Screen` | Known | Screen layout |
| 0x007B1D0D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B1D73 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B1DE1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B1E4E | `NoContent_Screen` | Known | Screen layout |
| 0x007B1E62 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B1ECA | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B1F34 | `NoContent_Screen` | Known | Screen layout |
| 0x007B1F48 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B1FAF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B2019 | `NoContent_Screen` | Known | Screen layout |
| 0x007B202D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B209A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B210C | `NoContent_Screen` | Known | Screen layout |
| 0x007B2120 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B2188 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B21F1 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B220C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B2272 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B228E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B236D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B2386 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B23E7 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B23FB | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B2569 | `Radio_Screen` | Known | Screen layout |
| 0x007B2579 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B25DA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B265D | `LockediPod_Screen` | Known | Screen layout |
| 0x007B26E5 | `Lock_Screen` | Known | Screen layout |
| 0x007B26F4 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B2757 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B27B9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B27D5 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B2847 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B2866 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B28CE | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B28E8 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B2950 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B296D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B29D9 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B2A43 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B2A5D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B2ACD | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B2B40 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B2BB1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B2C20 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B2C8C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B2CA7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B2D1C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B2D83 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B2DE5 | `Photos_Screen` | Known | Screen layout |
| 0x007B2E49 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B2E67 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B2ED9 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B2EF6 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B2F5C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B2F77 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B2FE0 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B2FFD | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B3074 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B3098 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B3106 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B3121 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B31DC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B31F8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B3266 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B3283 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B32EE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B330E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B3385 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B33A1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B3411 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B3430 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B349C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B34B0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B3529 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B359D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B360D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B3675 | `NoContent_Screen` | Known | Screen layout |
| 0x007B3689 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B36ED | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B3754 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B376E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B37DC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B384E | `NoContent_Screen` | Known | Screen layout |
| 0x007B3862 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B38CC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B3935 | `No_Photos_Screen` | Known | Screen layout |
| 0x007B3949 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B39AF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B3A1D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B3A8A | `NoContent_Screen` | Known | Screen layout |
| 0x007B3A9E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B3B06 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B3B70 | `NoContent_Screen` | Known | Screen layout |
| 0x007B3B84 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B3BEB | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B3C55 | `NoContent_Screen` | Known | Screen layout |
| 0x007B3C69 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B3CD6 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B3D48 | `NoContent_Screen` | Known | Screen layout |
| 0x007B3D5C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B3DC4 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B3E2D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B3E48 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B3EAE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B3ECA | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B3FA9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B3FC2 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B4023 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B4037 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B41A5 | `Radio_Screen` | Known | Screen layout |
| 0x007B41B5 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B4216 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B4299 | `LockediPod_Screen` | Known | Screen layout |
| 0x007B4321 | `Lock_Screen` | Known | Screen layout |
| 0x007B4330 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B4393 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B43F5 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B4411 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B4483 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B44A2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B450A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B4524 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B458C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B45A9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B4615 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B467F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B4699 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B4709 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B477C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B47ED | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B485C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B48C8 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B48E3 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B4958 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B49BF | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B4A21 | `Photos_Screen` | Known | Screen layout |
| 0x007B4A85 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B4AA3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B4B15 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B4B32 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B4B98 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B4BB3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B4C1C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B4C39 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B4CB0 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B4CD4 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B4D42 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B4D5D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B4E18 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B4E34 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B4EA2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B4EBF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B4F2A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B4F4A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B4FC1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B4FDD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B504D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B506C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B50D8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B50EC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B5165 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B51D9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B5249 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B52B1 | `NoContent_Screen` | Known | Screen layout |
| 0x007B52C5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B5329 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B5390 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B53AA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B5418 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B548A | `NoContent_Screen` | Known | Screen layout |
| 0x007B549E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B5508 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B5571 | `No_Photos_Screen` | Known | Screen layout |
| 0x007B5585 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B55EB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B5659 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B56C6 | `NoContent_Screen` | Known | Screen layout |
| 0x007B56DA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B5742 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B57AC | `NoContent_Screen` | Known | Screen layout |
| 0x007B57C0 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B5827 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B5891 | `NoContent_Screen` | Known | Screen layout |
| 0x007B58A5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B5912 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B5984 | `NoContent_Screen` | Known | Screen layout |
| 0x007B5998 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B5A00 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B5A69 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B5A84 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B5AEA | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B5B06 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B5BE5 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B5BFE | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B5C5F | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B5C73 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B5DE1 | `Radio_Screen` | Known | Screen layout |
| 0x007B5DF1 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B5E52 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B5ED5 | `LockediPod_Screen` | Known | Screen layout |
| 0x007B5F5D | `Lock_Screen` | Known | Screen layout |
| 0x007B5F6C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B5FCF | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B6031 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B604D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B60BF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B60DE | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B6146 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B6160 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B61C8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B61E5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B6251 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B62BB | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B62D5 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B6345 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B63B8 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B6429 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B6498 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B6504 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B651F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B6594 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B65FB | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B665D | `Photos_Screen` | Known | Screen layout |
| 0x007B66C1 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B66DF | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B6751 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B676E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B67D4 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B67EF | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B6858 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B6875 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B68EC | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B6910 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B697E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B6999 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B6A54 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B6A70 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B6ADE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B6AFB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B6B66 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B6B86 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B6BFD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B6C19 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B6C89 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B6CA8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B6D14 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B6D28 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B6DA1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B6E15 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B6E85 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B6EED | `NoContent_Screen` | Known | Screen layout |
| 0x007B6F01 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B6F65 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B6FCC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B6FE6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B7054 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B70C6 | `NoContent_Screen` | Known | Screen layout |
| 0x007B70DA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B7144 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B71AD | `No_Photos_Screen` | Known | Screen layout |
| 0x007B71C1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B7227 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B7295 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B7302 | `NoContent_Screen` | Known | Screen layout |
| 0x007B7316 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B737E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B73E8 | `NoContent_Screen` | Known | Screen layout |
| 0x007B73FC | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B7463 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B74CD | `NoContent_Screen` | Known | Screen layout |
| 0x007B74E1 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B754E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B75C0 | `NoContent_Screen` | Known | Screen layout |
| 0x007B75D4 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B763C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B76A5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B76C0 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B7726 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B7742 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B7821 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B783A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B789B | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B78AF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B7A1D | `Radio_Screen` | Known | Screen layout |
| 0x007B7A2D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B7A8E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B7B11 | `LockediPod_Screen` | Known | Screen layout |
| 0x007B7B99 | `Lock_Screen` | Known | Screen layout |
| 0x007B7BA8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B7C0B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B7C6D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B7C89 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B7CFB | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B7D1A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B7D82 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B7D9C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B7E04 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B7E21 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B7E8D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B7EF7 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B7F11 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B7F81 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B7FF4 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B8065 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B80D4 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B8140 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B815B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B81D0 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B8237 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B8299 | `Photos_Screen` | Known | Screen layout |
| 0x007B82FD | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B831B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B838D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B83AA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B8410 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B842B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B8494 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B84B1 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B8528 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B854C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B85BA | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B85D5 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B8690 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B86AC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B871A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B8737 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B87A2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B87C2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B8839 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B8855 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B88C5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B88E4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B8950 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B8964 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B89DD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B8A51 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B8AC1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B8B29 | `NoContent_Screen` | Known | Screen layout |
| 0x007B8B3D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B8BA1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B8C08 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B8C22 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B8C90 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B8D02 | `NoContent_Screen` | Known | Screen layout |
| 0x007B8D16 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B8D80 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B8DE9 | `No_Photos_Screen` | Known | Screen layout |
| 0x007B8DFD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B8E63 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B8ED1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B8F3E | `NoContent_Screen` | Known | Screen layout |
| 0x007B8F52 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B8FBA | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B9024 | `NoContent_Screen` | Known | Screen layout |
| 0x007B9038 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B909F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B9109 | `NoContent_Screen` | Known | Screen layout |
| 0x007B911D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B918A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B91FC | `NoContent_Screen` | Known | Screen layout |
| 0x007B9210 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B9278 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B92E1 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B92FC | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B9362 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B937E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B945D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B9476 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B94D7 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B94EB | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B9659 | `Radio_Screen` | Known | Screen layout |
| 0x007B9669 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B96CA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B974D | `LockediPod_Screen` | Known | Screen layout |
| 0x007B97D5 | `Lock_Screen` | Known | Screen layout |
| 0x007B97E4 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B9847 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B98A9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B98C5 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B9937 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B9956 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B99BE | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B99D8 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B9A40 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B9A5D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B9AC9 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B9B33 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B9B4D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B9BBD | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B9C30 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B9CA1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B9D10 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B9D7C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B9D97 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B9E0C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B9E73 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B9ED5 | `Photos_Screen` | Known | Screen layout |
| 0x007B9F39 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B9F57 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B9FC9 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B9FE6 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007BA04C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BA067 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BA0D0 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007BA0ED | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007BA164 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007BA188 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007BA1F6 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007BA211 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007BA2CC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BA2E8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BA356 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BA373 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BA3DE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BA3FE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BA475 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BA491 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BA501 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BA520 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BA58C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BA5A0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007BA619 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BA68D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007BA6FD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007BA765 | `NoContent_Screen` | Known | Screen layout |
| 0x007BA779 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007BA7DD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007BA844 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007BA85E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007BA8CC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007BA93E | `NoContent_Screen` | Known | Screen layout |
| 0x007BA952 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007BA9BC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007BAA25 | `No_Photos_Screen` | Known | Screen layout |
| 0x007BAA39 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007BAA9F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007BAB0D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007BAB7A | `NoContent_Screen` | Known | Screen layout |
| 0x007BAB8E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007BABF6 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007BAC60 | `NoContent_Screen` | Known | Screen layout |
| 0x007BAC74 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007BACDB | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007BAD45 | `NoContent_Screen` | Known | Screen layout |
| 0x007BAD59 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007BADC6 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007BAE38 | `NoContent_Screen` | Known | Screen layout |
| 0x007BAE4C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007BAEB4 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007BAF1D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007BAF38 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007BAF9E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BAFBA | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BB099 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007BB0B2 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007BB113 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007BB127 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007BB295 | `Radio_Screen` | Known | Screen layout |
| 0x007BB2A5 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007BB306 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007BB389 | `LockediPod_Screen` | Known | Screen layout |
| 0x007BB411 | `Lock_Screen` | Known | Screen layout |
| 0x007BB420 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007BB483 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007BB4E5 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007BB501 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007BB573 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BB592 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BB5FA | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007BB614 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007BB67C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BB699 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BB705 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BB76F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007BB789 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007BB7F9 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007BB86C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007BB8DD | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007BB94C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007BB9B8 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007BB9D3 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007BBA48 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007BBAAF | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007BBB11 | `Photos_Screen` | Known | Screen layout |
| 0x007BBB75 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007BBB93 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007BBC05 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007BBC22 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007BBC88 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BBCA3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BBD0C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007BBD29 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007BBDA0 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007BBDC4 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007BBE32 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007BBE4D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007BBF08 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BBF24 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BBF92 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BBFAF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BC01A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BC03A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BC0B1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BC0CD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BC13D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BC15C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BC1C8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BC1DC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007BC255 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BC2C9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007BC339 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007BC3A1 | `NoContent_Screen` | Known | Screen layout |
| 0x007BC3B5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007BC419 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007BC480 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007BC49A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007BC508 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007BC57A | `NoContent_Screen` | Known | Screen layout |
| 0x007BC58E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007BC5F8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007BC661 | `No_Photos_Screen` | Known | Screen layout |
| 0x007BC675 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007BC6DB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007BC749 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007BC7B6 | `NoContent_Screen` | Known | Screen layout |
| 0x007BC7CA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007BC832 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007BC89C | `NoContent_Screen` | Known | Screen layout |
| 0x007BC8B0 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007BC917 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007BC981 | `NoContent_Screen` | Known | Screen layout |
| 0x007BC995 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007BCA02 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007BCA74 | `NoContent_Screen` | Known | Screen layout |
| 0x007BCA88 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007BCAF0 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007BCB59 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007BCB74 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007BCBDA | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BCBF6 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BCCD5 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007BCCEE | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007BCD4F | `FirstBoot_Screen` | Known | Screen layout |
| 0x007BCD63 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007BCED1 | `Radio_Screen` | Known | Screen layout |
| 0x007BCEE1 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007BCF42 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007BCFC5 | `LockediPod_Screen` | Known | Screen layout |
| 0x007BD04D | `Lock_Screen` | Known | Screen layout |
| 0x007BD05C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007BD0BF | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007BD121 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007BD13D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007BD1AF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BD1CE | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BD236 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007BD250 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007BD2B8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BD2D5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BD341 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BD3AB | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007BD3C5 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007BD435 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007BD4A8 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007BD519 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007BD588 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007BD5F4 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007BD60F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007BD684 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007BD6EB | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007BD74D | `Photos_Screen` | Known | Screen layout |
| 0x007BD7B1 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007BD7CF | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007BD841 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007BD85E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007BD8C4 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BD8DF | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BD948 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007BD965 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007BD9DC | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007BDA00 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007BDA6E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007BDA89 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007BDB29 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BDB45 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BDBB3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BDBD0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BDC3B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BDC5B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BDCD2 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BDCEE | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BDD5E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BDD7D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BDDE9 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BDDFD | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007BDE72 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007BDEDD | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007BDF4C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007BDFBD | `NoContent_Screen` | Known | Screen layout |
| 0x007BDFD1 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BE040 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007BE0B3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007BE120 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007BE189 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007BE1F9 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007BE269 | `NoContent_Screen` | Known | Screen layout |
| 0x007BE27D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007BE2E0 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007BE343 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BE35F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BE41F | `Radio_Screen` | Known | Screen layout |
| 0x007BE42F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007BE490 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007BE4FE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BE51D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BE58B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BE5F0 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BE60B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BE6B1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BE6CD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BE73B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BE758 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BE7C3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BE7E3 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BE85A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BE876 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BE8E6 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BE905 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BE971 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BE985 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007BE9FA | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007BEA65 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007BEAD4 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007BEB45 | `NoContent_Screen` | Known | Screen layout |
| 0x007BEB59 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BEBC8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007BEC3B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007BECA8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007BED11 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007BED81 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007BEDF1 | `NoContent_Screen` | Known | Screen layout |
| 0x007BEE05 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007BEE68 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007BEECB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BEEE7 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BEFA7 | `Radio_Screen` | Known | Screen layout |
| 0x007BEFB7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007BF018 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007BF086 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BF0A5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BF113 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BF178 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BF193 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BF239 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BF255 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BF2C3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BF2E0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BF34B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BF36B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BF3E2 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BF3FE | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BF46E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BF48D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BF4F9 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BF50D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007BF582 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007BF5ED | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007BF65C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007BF6CD | `NoContent_Screen` | Known | Screen layout |
| 0x007BF6E1 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BF750 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007BF7C3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007BF830 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007BF899 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007BF909 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007BF979 | `NoContent_Screen` | Known | Screen layout |
| 0x007BF98D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007BF9F0 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007BFA53 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BFA6F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BFB2F | `Radio_Screen` | Known | Screen layout |
| 0x007BFB3F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007BFBA0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007BFC0E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BFC2D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BFC9B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BFD00 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BFD1B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BFDC1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BFDDD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BFE4B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BFE68 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BFED3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BFEF3 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BFF6A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BFF86 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BFFF6 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C0015 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C0081 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C0095 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C010A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C0175 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C01E4 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C0255 | `NoContent_Screen` | Known | Screen layout |
| 0x007C0269 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C02D8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C034B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C03B8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C0421 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C0491 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C0501 | `NoContent_Screen` | Known | Screen layout |
| 0x007C0515 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C0578 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C05DB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C05F7 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C06B7 | `Radio_Screen` | Known | Screen layout |
| 0x007C06C7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C0728 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C0796 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C07B5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C0823 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C0888 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C08A3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C0949 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C0965 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C09D3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C09F0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C0A5B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C0A7B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C0AF2 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C0B0E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C0B7E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C0B9D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C0C09 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C0C1D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C0C92 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C0CFD | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C0D6C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C0DDD | `NoContent_Screen` | Known | Screen layout |
| 0x007C0DF1 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C0E60 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C0ED3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C0F40 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C0FA9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C1019 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C1089 | `NoContent_Screen` | Known | Screen layout |
| 0x007C109D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C1100 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C1163 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C117F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C123F | `Radio_Screen` | Known | Screen layout |
| 0x007C124F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C12B0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C131E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C133D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C13AB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C1410 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C142B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C14D1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C14ED | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C155B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C1578 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C15E3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C1603 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C167A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C1696 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C1706 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C1725 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C1791 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C17A5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C181A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C1885 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C18F4 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C1965 | `NoContent_Screen` | Known | Screen layout |
| 0x007C1979 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C19E8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C1A5B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C1AC8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C1B31 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C1BA1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C1C11 | `NoContent_Screen` | Known | Screen layout |
| 0x007C1C25 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C1C88 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C1CEB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C1D07 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C1DC7 | `Radio_Screen` | Known | Screen layout |
| 0x007C1DD7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C1E38 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C1EA6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C1EC5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C1F33 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C1F98 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C1FB3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C2059 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C2075 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C20E3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C2100 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C216B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C218B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C2202 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C221E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C228E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C22AD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C2319 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C232D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C23A2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C240D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C247C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C24ED | `NoContent_Screen` | Known | Screen layout |
| 0x007C2501 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C2570 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C25E3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C2650 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C26B9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C2729 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C2799 | `NoContent_Screen` | Known | Screen layout |
| 0x007C27AD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C2810 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C2873 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C288F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C294F | `Radio_Screen` | Known | Screen layout |
| 0x007C295F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C29C0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C2A2E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C2A4D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C2ABB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C2B20 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C2B3B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C2BE1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C2BFD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C2C6B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C2C88 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C2CF3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C2D13 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C2D8A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C2DA6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C2E16 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C2E35 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C2EA1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C2EB5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C2F2A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C2F95 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C3004 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C3075 | `NoContent_Screen` | Known | Screen layout |
| 0x007C3089 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C30F8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C316B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C31D8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C3241 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C32B1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C3321 | `NoContent_Screen` | Known | Screen layout |
| 0x007C3335 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C3398 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C33FB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C3417 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C34D7 | `Radio_Screen` | Known | Screen layout |
| 0x007C34E7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C3548 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C35B6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C35D5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C3643 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C36A8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C36C3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C3769 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C3785 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C37F3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C3810 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C387B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C389B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C3912 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C392E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C399E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C39BD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C3A29 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C3A3D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C3AB2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C3B1D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C3B8C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C3BFD | `NoContent_Screen` | Known | Screen layout |
| 0x007C3C11 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C3C80 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C3CF3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C3D60 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C3DC9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C3E39 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C3EA9 | `NoContent_Screen` | Known | Screen layout |
| 0x007C3EBD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C3F20 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C3F83 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C3F9F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C405F | `Radio_Screen` | Known | Screen layout |
| 0x007C406F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C40D0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C413E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C415D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C41CB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C4230 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C424B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C42F1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C430D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C437B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C4398 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C4403 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C4423 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C449A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C44B6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C4526 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C4545 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C45B1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C45C5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C463A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C46A5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C4714 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C4785 | `NoContent_Screen` | Known | Screen layout |
| 0x007C4799 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C4808 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C487B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C48E8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C4951 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C49C1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C4A31 | `NoContent_Screen` | Known | Screen layout |
| 0x007C4A45 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C4AA8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C4B0B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C4B27 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C4BE7 | `Radio_Screen` | Known | Screen layout |
| 0x007C4BF7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C4C58 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C4CC6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C4CE5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C4D53 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C4DB8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C4DD3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C4E79 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C4E95 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C4F03 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C4F20 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C4F8B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C4FAB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C5022 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C503E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C50AE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C50CD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C5139 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C514D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C51C2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C522D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C529C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C530D | `NoContent_Screen` | Known | Screen layout |
| 0x007C5321 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C5390 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C5403 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C5470 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C54D9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C5549 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C55B9 | `NoContent_Screen` | Known | Screen layout |
| 0x007C55CD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C5630 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C5693 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C56AF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C576F | `Radio_Screen` | Known | Screen layout |
| 0x007C577F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C57E0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C584E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C586D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C58DB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C5940 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C595B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C5A01 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C5A1D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C5A8B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C5AA8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C5B13 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007C5B33 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007C5BAA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C5BC6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007C5C36 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007C5C55 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007C5CC1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007C5CD5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007C5D4A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007C5DB5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007C5E24 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007C5E95 | `NoContent_Screen` | Known | Screen layout |
| 0x007C5EA9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007C5F18 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007C5F8B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007C5FF8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007C6061 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007C60D1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007C6141 | `NoContent_Screen` | Known | Screen layout |
| 0x007C6155 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007C61B8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007C621B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007C6237 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007C62F7 | `Radio_Screen` | Known | Screen layout |
| 0x007C6307 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007C6368 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007C63D6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007C63F5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007C6463 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007C64C8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C64E3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C65C4 | `MediaLists_AudiobookChapters_Screen)` | Known | Screen layout |
| 0x007C65EB | `MediaLists_AudiobookChapters_Screen_Plane!` | Known | Screen layout |
| 0x007C6D85 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007C6DA0 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x007C6E0B | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C6E26 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x007C6E99 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x007C6EB4 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x007C7071 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007C708C | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x007C70F7 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C7112 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x007C7185 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x007C71A0 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x007C7368 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C7384 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x007C73FF | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C741B | `MediaLists_Albums_Screen_WithTrackCount%` | Known | Screen layout |
| 0x007C7494 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007C74AF | `MediaLists_Songs_Screen_Plain1` | Known | Screen layout |
| 0x007C752A | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x007C7545 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x007C7767 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C7784 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C7863 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C787F | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x007C78FA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007C7915 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007C7AFB | `MediaLists_PodcastEpisodes_Screen'` | Known | Screen layout |
| 0x007C7B20 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007C7DF2 | `MediaLists_TVSeasons_Screen#` | Known | Screen layout |
| 0x007C7E11 | `MediaLists_TVSeasons_Screen_Default%` | Known | Screen layout |
| 0x007C7E86 | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x007C7EA6 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007C802E | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x007C804E | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007C8447 | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x007C846C | `MediaLists_NestedPlaylists_Screen_Default,` | Known | Screen layout |
| 0x007C84EE | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x007C850D | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x007C869D | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x007C86C2 | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x007C873A | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x007C8759 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x007C87BD | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C886A | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C88DC | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x007C89D2 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x007C8C94 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x007C8D94 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x007C8E00 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007C8E6A | `NoContent_Screen` | Known | Screen layout |
| 0x007C8E7E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007C8EE8 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007C8F5C | `NoContent_Screen` | Known | Screen layout |
| 0x007C8F70 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007C8FDB | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007C9047 | `NoContent_Screen` | Known | Screen layout |
| 0x007C905B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007C90C2 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007C912E | `NoContent_Screen` | Known | Screen layout |
| 0x007C9142 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007C91AF | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007C9223 | `NoContent_Screen` | Known | Screen layout |
| 0x007C9237 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007C929F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007C930C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007C9370 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007C938C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007C93F8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C9415 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C9482 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007C9549 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007C9566 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007C95DD | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007C9601 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C96B8 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007C9722 | `NoContent_Screen` | Known | Screen layout |
| 0x007C9736 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007C97A0 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007C9814 | `NoContent_Screen` | Known | Screen layout |
| 0x007C9828 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007C9893 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007C98FF | `NoContent_Screen` | Known | Screen layout |
| 0x007C9913 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007C997A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007C99E6 | `NoContent_Screen` | Known | Screen layout |
| 0x007C99FA | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007C9A67 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007C9ADB | `NoContent_Screen` | Known | Screen layout |
| 0x007C9AEF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007C9B57 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007C9BC4 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007C9C28 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007C9C44 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007C9CB0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C9CCD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C9D3A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007C9E01 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007C9E1E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007C9E95 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007C9EB9 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C9F70 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007C9FDA | `NoContent_Screen` | Known | Screen layout |
| 0x007C9FEE | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007CA058 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007CA0CC | `NoContent_Screen` | Known | Screen layout |
| 0x007CA0E0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007CA14B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007CA1B7 | `NoContent_Screen` | Known | Screen layout |
| 0x007CA1CB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007CA232 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007CA29E | `NoContent_Screen` | Known | Screen layout |
| 0x007CA2B2 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007CA31F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CA393 | `NoContent_Screen` | Known | Screen layout |
| 0x007CA3A7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CA40F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CA47C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CA4E0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CA4FC | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CA568 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CA585 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CA5F2 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CA6B9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CA6D6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CA74D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CA771 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CA828 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007CA892 | `NoContent_Screen` | Known | Screen layout |
| 0x007CA8A6 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007CA910 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007CA984 | `NoContent_Screen` | Known | Screen layout |
| 0x007CA998 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007CAA03 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007CAA6F | `NoContent_Screen` | Known | Screen layout |
| 0x007CAA83 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007CAAEA | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007CAB56 | `NoContent_Screen` | Known | Screen layout |
| 0x007CAB6A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007CABD7 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CAC4B | `NoContent_Screen` | Known | Screen layout |
| 0x007CAC5F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CACC7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CAD34 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CAD98 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CADB4 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CAE20 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CAE3D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CAEAA | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CAF71 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CAF8E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CB005 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CB029 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CB0E0 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007CB14A | `NoContent_Screen` | Known | Screen layout |
| 0x007CB15E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007CB1C8 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007CB23C | `NoContent_Screen` | Known | Screen layout |
| 0x007CB250 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007CB2BB | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007CB327 | `NoContent_Screen` | Known | Screen layout |
| 0x007CB33B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007CB3A2 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007CB40E | `NoContent_Screen` | Known | Screen layout |
| 0x007CB422 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007CB48F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CB503 | `NoContent_Screen` | Known | Screen layout |
| 0x007CB517 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CB57F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CB5EC | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CB650 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CB66C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CB6D8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CB6F5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CB762 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CB829 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CB846 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CB8BD | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CB8E1 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CB998 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007CBA02 | `NoContent_Screen` | Known | Screen layout |
| 0x007CBA16 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007CBA80 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007CBAF4 | `NoContent_Screen` | Known | Screen layout |
| 0x007CBB08 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007CBB73 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007CBBDF | `NoContent_Screen` | Known | Screen layout |
| 0x007CBBF3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007CBC5A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007CBCC6 | `NoContent_Screen` | Known | Screen layout |
| 0x007CBCDA | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007CBD47 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CBDBB | `NoContent_Screen` | Known | Screen layout |
| 0x007CBDCF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CBE37 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CBEA4 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CBF08 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CBF24 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CBF90 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CBFAD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CC01A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CC0E1 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CC0FE | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CC175 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CC199 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CC250 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007CC2BA | `NoContent_Screen` | Known | Screen layout |
| 0x007CC2CE | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007CC338 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007CC3AC | `NoContent_Screen` | Known | Screen layout |
| 0x007CC3C0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007CC42B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007CC497 | `NoContent_Screen` | Known | Screen layout |
| 0x007CC4AB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007CC512 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007CC57E | `NoContent_Screen` | Known | Screen layout |
| 0x007CC592 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007CC5FF | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CC673 | `NoContent_Screen` | Known | Screen layout |
| 0x007CC687 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CC6EF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CC75C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CC7C0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CC7DC | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CC848 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CC865 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CC8D2 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CC999 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CC9B6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CCA2D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CCA51 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CCB08 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007CCB72 | `NoContent_Screen` | Known | Screen layout |
| 0x007CCB86 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007CCBF0 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007CCC64 | `NoContent_Screen` | Known | Screen layout |
| 0x007CCC78 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007CCCE3 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007CCD4F | `NoContent_Screen` | Known | Screen layout |
| 0x007CCD63 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007CCDCA | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007CCE36 | `NoContent_Screen` | Known | Screen layout |
| 0x007CCE4A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007CCEB7 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007CCF2B | `NoContent_Screen` | Known | Screen layout |
| 0x007CCF3F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007CCFA7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007CD014 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007CD078 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007CD094 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007CD100 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007CD11D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007CD18A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007CD251 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007CD26E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007CD2E5 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007CD309 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007CD76C | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CD7DE | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CD849 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CD8AE | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CD918 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CD982 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CD9F2 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CDA69 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CDAD7 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CDB42 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CDBAC | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CDC13 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CDC82 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CDCF0 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CDD55 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CDDBD | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CDE28 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CDE93 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CDEFA | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CE268 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CE2DA | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CE345 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CE3AA | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CE414 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CE47E | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CE4EE | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CE565 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CE5D3 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CE63E | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CE6A8 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CE70F | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CE77E | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CE7EC | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CE851 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CE8B9 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CE924 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CE98F | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CE9F6 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CED62 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CEDD4 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CEE3F | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CEEA4 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CEF0E | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CEF78 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CEFE8 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CF05F | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CF0CD | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CF138 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CF1A2 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CF209 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CF278 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CF2E6 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CF34B | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CF3B3 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CF41E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CF489 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CF4F0 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CF85A | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CF8CC | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CF937 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CF99C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CFA06 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CFA70 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CFAE0 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CFB57 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CFBC5 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CFC30 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CFC9A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CFD01 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CFD70 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CFDDE | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CFE43 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CFEAB | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CFF16 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CFF81 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CFFE8 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D033A | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D03AC | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D0417 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D047C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D04E6 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D0550 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D05C0 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D0637 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D06A5 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D0710 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D077A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D07E1 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D0850 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D08BE | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D0923 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D098B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D09F6 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D0A61 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D0AC8 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D0E3F | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D0EB1 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D0F1C | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D0F81 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D0FEB | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D1055 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D10C5 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D113C | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D11AA | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D1215 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D127F | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D12E6 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D1355 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D13C3 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D1428 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D1490 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D14FB | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D1566 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D15CD | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D1941 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D19B3 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D1A1E | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D1A83 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D1AED | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D1B57 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D1BC7 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D1C3E | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D1CAC | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D1D17 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D1D81 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D1DE8 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D1E57 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D1EC5 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D1F2A | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D1F92 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D1FFD | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D2068 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D20CF | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D2429 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D249B | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D2506 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D256B | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D25D5 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D263F | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D26AF | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D2726 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D2794 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D27FF | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D2869 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D28D0 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D293F | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D29AD | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D2A12 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D2A7A | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D2AE5 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D2B50 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D2BB7 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D2F11 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D2F83 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D2FEE | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D3053 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D30BD | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D3127 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D3197 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D320E | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D327C | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D32E7 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D3351 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D33B8 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D3427 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D3495 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D34FA | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D3562 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D35CD | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D3638 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D369F | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D39FA | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D3A6C | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D3AD7 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D3B3C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D3BA6 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D3C10 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D3C80 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D3CF7 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D3D65 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D3DD0 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D3E3A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D3EA1 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D3F10 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D3F7E | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D3FE3 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D404B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D40B6 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D4121 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D4188 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D450C | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D457E | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D45E9 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D464E | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D46B8 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D4722 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D4792 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D4809 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D4877 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D48E2 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D494C | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D49B3 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D4A22 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D4A90 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D4AF5 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D4B5D | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D4BC8 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D4C33 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D4C9A | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D5028 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D509A | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D5105 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D516A | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D51D4 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D523E | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D52AE | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D5325 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D5393 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D53FE | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D5468 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D54CF | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D553E | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D55AC | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D5611 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D5679 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D56E4 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D574F | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D57B6 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D5B24 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D5B96 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D5C01 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D5C66 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D5CD0 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D5D3A | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D5DAA | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D5E21 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D5E8F | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D5EFA | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D5F64 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D5FCB | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D603A | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D60A8 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D610D | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D6175 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D61E0 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D624B | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D62B2 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D6618 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D668A | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D66F5 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D675A | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D67C4 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D682E | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D689E | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D6915 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D6983 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D69EE | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D6A58 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D6ABF | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D6B2E | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D6B9C | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D6C01 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D6C69 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D6CD4 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D6D3F | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D6DA6 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D70FA | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D716C | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D71D7 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D723C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D72A6 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D7310 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D7380 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D73F7 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D7465 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D74D0 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D753A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D75A1 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D7610 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D767E | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D76E3 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D774B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D77B6 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D7821 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D7888 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D7BD3 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D7C45 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D7CB0 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D7D15 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D7D7F | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D7DE9 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D7E59 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D7ED0 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D7F3E | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D7FA9 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D8013 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D807A | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D80E9 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D8157 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D81BC | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D8224 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D828F | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D82FA | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D8361 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D86C3 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D8735 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D87A0 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D8805 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D886F | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D88D9 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D8949 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D89C0 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D8A2E | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D8A99 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D8B03 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D8B6A | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D8BD9 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D8C47 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D8CAC | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D8D14 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D8D7F | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D8DEA | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D8E51 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D9169 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007D91DB | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007D9246 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007D92AB | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007D9315 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D937F | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D93EF | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D9466 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D94D4 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D953F | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D95A9 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D9610 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D967F | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D96ED | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D9752 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D97BA | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D9825 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D9890 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D98F7 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D9C0E | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007D9C85 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007D9D02 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007D9D74 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007D9DE4 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007D9E5A | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007D9EC8 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007D9F35 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DA27A | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DA2F1 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DA36E | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DA3E0 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DA450 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DA4C6 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DA534 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DA5A1 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DA90A | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DA981 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DA9FE | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DAA70 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DAAE0 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DAB56 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DABC4 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DAC31 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DAF9A | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DB011 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DB08C | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DB0FC | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DB172 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DB1E0 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DB24D | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DB586 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DB5FD | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DB678 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DB6E8 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DB75E | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DB7CC | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DB839 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DBB70 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DBBE7 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DBC62 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DBCD2 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DBD48 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DBDB6 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DBE23 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DC133 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007DC1AA | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007DC225 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007DC295 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007DC30B | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007DC379 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007DC3E6 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007DC9EA | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007DCA07 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007DCA82 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007DCA9B | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007DCB13 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007DCB2C | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007DCBA1 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DCBB7 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007DCC2E | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DCC44 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007DCCBB | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007DCCD8 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007DCD50 | `Notes_List_Screen` | Known | Screen layout |
| 0x007DCD65 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007DCF16 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007DCF33 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007DCFAE | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007DCFC7 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007DD03F | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007DD058 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007DD0CD | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DD0E3 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007DD15A | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DD170 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007DD1E7 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007DD204 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007DD27C | `Notes_List_Screen` | Known | Screen layout |
| 0x007DD291 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007DD472 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007DD48F | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007DD50A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007DD523 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007DD59B | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007DD5B4 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007DD629 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DD63F | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007DD6B6 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DD6CC | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007DD743 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007DD760 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007DD7D8 | `Notes_List_Screen` | Known | Screen layout |
| 0x007DD7ED | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007DD9A2 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007DD9BF | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007DDA3A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007DDA53 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007DDACB | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007DDAE4 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007DDB59 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DDB6F | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007DDBE6 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007DDBFC | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007DDC73 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007DDC90 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007DDD08 | `Notes_List_Screen` | Known | Screen layout |
| 0x007DDD1D | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007DE035 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007DE0DB | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007DE15E | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007DE216 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x007DE298 | `PhotosSettingsSlideshowMusic_Screen+` | Known | Screen layout |
| 0x007DE2BF | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x007DE3A5 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x007DE55D | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007DE5BD | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007DE61A | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007DE641 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007DE6E1 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007DE741 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007DE79E | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007DE7C5 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007DEA60 | `Photos_Screen` | Known | Screen layout |
| 0x007DEBAC | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007DEC10 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007DEC71 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007DECCE | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007DED2B | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007DED99 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007DEDF6 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007DEF6C | `Photos_Screen` | Known | Screen layout |
| 0x007DF0B8 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007DF11C | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007DF17D | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007DF1DA | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007DF237 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007DF2A5 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007DF302 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007DF478 | `Photos_Screen` | Known | Screen layout |
| 0x007DF5C4 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007DF628 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007DF689 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007DF6E6 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007DF743 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007DF7B1 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007DF80E | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007DF984 | `Photos_Screen` | Known | Screen layout |
| 0x007DFAD0 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007DFB34 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007DFB95 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007DFBF2 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007DFC4F | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007DFCBD | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007DFD1A | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007DFE90 | `Photos_Screen` | Known | Screen layout |
| 0x007DFFDC | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E0040 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007E00A1 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007E00FE | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007E015B | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E01C9 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007E0226 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007E039C | `Photos_Screen` | Known | Screen layout |
| 0x007E04E8 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E054C | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007E05AD | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007E060A | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007E0667 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E06D5 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007E0732 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007E08A8 | `Photos_Screen` | Known | Screen layout |
| 0x007E09F4 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E0A5A | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007E0ABC | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007E0B1E | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E0BB4 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007E0CD5 | `Photos_Screen` | Known | Screen layout |
| 0x007E0D40 | `Photos_Screen` | Known | Screen layout |
| 0x007E0E8C | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E0EF2 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007E0F54 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007E0FB6 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E104C | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007E116D | `Photos_Screen` | Known | Screen layout |
| 0x007E11D8 | `Photos_Screen` | Known | Screen layout |
| 0x007E1324 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E138A | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007E13EC | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007E144E | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E14E4 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007E1605 | `Photos_Screen` | Known | Screen layout |
| 0x007E1670 | `Photos_Screen` | Known | Screen layout |
| 0x007E17BC | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E1822 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007E1884 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007E18E6 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E197C | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007E1A9D | `Photos_Screen` | Known | Screen layout |
| 0x007E1B08 | `Photos_Screen` | Known | Screen layout |
| 0x007E1C54 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007E1CBA | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007E1D1C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007E1D7E | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007E1E14 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007E1F35 | `Photos_Screen` | Known | Screen layout |
| 0x007E2129 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007E218B | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007E21F9 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007E225F | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007E22C4 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007E2592 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007E25F4 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007E2662 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007E26C8 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007E29CE | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007E2A30 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007E2A9E | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007E2B04 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007E2DAD | `Radio_Screen_Default` | Known | Screen layout |
| 0x007E2E0A | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007E2E6C | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007E2EDA | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007E2F40 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007E323A | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007E32A4 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007E3512 | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007E357C | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007E3739 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007E379C | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007E3801 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007E3869 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007E38CC | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E3934 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007E399D | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E3A03 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007E3A68 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E3AD5 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007E3B45 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007E3BBB | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007E3C31 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007E3CA1 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E3D16 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007E3D8D | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007E3E01 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007E3E73 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007E3EED | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E3F60 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007E3FD2 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E4056 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E4080 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E4107 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E4194 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E4233 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E424D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E42C5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E42DF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E4349 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E4366 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E43DE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E4408 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E448F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E451C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E45BB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E45D5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E464D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E4667 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E46D1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E46EE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E4766 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E4790 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E4817 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E48A4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E4943 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E495D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E49D5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E49EF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E4A59 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E4A76 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E4AEE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E4B18 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E4B9F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E4C2C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E4CCB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E4CE5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E4D5D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E4D77 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E4DE1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E4DFE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E4E76 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E4EA0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E4F27 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E4FB4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E5053 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E506D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E50E5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E50FF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E5169 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E5186 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E51FE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E5228 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E52AF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E533C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E53DB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E53F5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E546D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E5487 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E54F1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E550E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E5586 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E55B0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E5637 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E56C4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E5763 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E577D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E57F5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E580F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E5879 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E5896 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E590E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E5938 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E59BF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E5A4C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E5AEB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E5B05 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E5B7D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E5B97 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E5C01 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E5C1E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E5C96 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E5CC0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E5D47 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E5DD4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E5E73 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E5E8D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E5F05 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E5F1F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E5F89 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E5FA6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E601E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E6048 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E60CF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E615C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E61FB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E6215 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E628D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E62A7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E6311 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E632E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E63A6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E63D0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E6457 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E64E4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E6583 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E659D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E6615 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E662F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E6699 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E66B6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E672E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E6758 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E67DF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E686C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E690B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E6925 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E699D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E69B7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E6A21 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E6A3E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E6AB6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E6AE0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E6B67 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E6BF4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E6C93 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E6CAD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E6D25 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E6D3F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E6DA9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E6DC6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E6E3E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E6E68 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E6EEF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E6F7C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E701B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E7035 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E70AD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E70C7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E7131 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E714E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E71C6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E71F0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E7277 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E7304 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E73A3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E73BD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E7435 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E744F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E74B9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E74D6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E754E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E7578 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E75FF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E768C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E772B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E7745 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E77BD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E77D7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E7841 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E785E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E78D6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E7900 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E7987 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E7A14 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E7AB3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E7ACD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E7B45 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E7B5F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E7BC9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E7BE6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E7C5E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E7C88 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E7D0F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E7D9C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E7E3B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E7E55 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E7ECD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E7EE7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E7F51 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E7F6E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E7FE6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007E8010 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007E8097 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007E8124 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007E81C3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E81DD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E8255 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E826F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E82D9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007E82F6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007E837D | `SettingsMenus_ResetMainMenu_Screen9` | Known | Screen layout |
| 0x007E844D | `SettingsMenus_ResetMusicMenu_Screen:` | Known | Screen layout |
| 0x007E8501 | `SettingsMenus_Main_Screen,` | Known | Screen layout |
| 0x007E8573 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E858D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007E8605 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007E861F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007E895A | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007E89C0 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007E8A1D | `Extras_Screen` | Known | Screen layout |
| 0x007E8A71 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x007E8B4F | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x007E8BBD | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007E8C5B | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x007E8C74 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x007E8CDC | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007E8D4F | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E8DD1 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007E8E32 | `NikePlus_Alert_Screen$` | Known | Screen layout |
| 0x007E8EB2 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E8F2B | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E8FA5 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E902A | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E904B | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E90BA | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x007E9142 | `NikePlus_Remote_Unlinking_Screen(` | Known | Screen layout |
| 0x007E9166 | `NikePlus_Remote_Unlinking_Screen_Default` | Known | Screen layout |
| 0x007E91DA | `NikePlus_Alert_Screen(` | Known | Screen layout |
| 0x007E9266 | `NikePlus_Dynamic_Workout_Screen,` | Known | Screen layout |
| 0x007E9289 | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x007E9302 | `NikePlus_Dynamic_Workout_Screen/` | Known | Screen layout |
| 0x007E9325 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x007E939E | `NikePlus_Dynamic_Workout_Screen.` | Known | Screen layout |
| 0x007E93C1 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x007E943A | `NikePlus_Alert_Screen(` | Known | Screen layout |
| 0x007E94BD | `NikePlus_Custom_Screen,` | Known | Screen layout |
| 0x007E94D7 | `NikePlus_Custom_Screen_Weight_ToWorkoutMusic` | Known | Screen layout |
| 0x007E9551 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E95D3 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007E964B | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007E9669 | `NikePlus_EndWorkout_Screen_Default` | Known | Screen layout |
| 0x007E9701 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007E977D | `NikePlus_EquipmentAlert_Screen<` | Known | Screen layout |
| 0x007E984A | `NikePlus_EquipmentAlert_Screen<` | Known | Screen layout |
| 0x007E9914 | `NikePlus_EquipmentAlert_Screen8` | Known | Screen layout |
| 0x007E99E1 | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007E9AAE | `NikePlus_EquipmentAlert_Screen1` | Known | Screen layout |
| 0x007E9B60 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E9B81 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E9C18 | `NikePlus_Dynamic_Workout_Screen.` | Known | Screen layout |
| 0x007E9C3B | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x007E9CDB | `NikePlus_Dynamic_Workout_Screen/` | Known | Screen layout |
| 0x007E9CFE | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x007E9D9C | `NikePlus_Dynamic_Workout_Screen,` | Known | Screen layout |
| 0x007E9DBF | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x007E9E55 | `NikePlus_EndPausedWorkout_Screen1` | Known | Screen layout |
| 0x007E9E79 | `NikePlus_EndPausedWorkout_Screen_BasicSave_Layout` | Known | Screen layout |
| 0x007E9F17 | `NikePlus_EndPausedWorkout_Screen4` | Known | Screen layout |
| 0x007E9F3B | `NikePlus_EndPausedWorkout_Screen_CaloriesSave_Layout` | Known | Screen layout |
| 0x007E9FDC | `NikePlus_EndPausedWorkout_Screen4` | Known | Screen layout |
| 0x007EA000 | `NikePlus_EndPausedWorkout_Screen_DistanceSave_Layout` | Known | Screen layout |
| 0x007EA09E | `NikePlus_EndPausedWorkout_Screen0` | Known | Screen layout |
| 0x007EA0C2 | `NikePlus_EndPausedWorkout_Screen_TimeSave_Layout` | Known | Screen layout |
| 0x007EA157 | `NikePlus_Songs_Screen` | Known | Screen layout |
| 0x007EA170 | `NikePlus_Songs_Screen_Default` | Known | Screen layout |
| 0x007EA282 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007EA29C | `NikePlus_Custom_Screen_Weight` | Known | Screen layout |
| 0x007EA2FF | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007EA373 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007EA3F1 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x007EA45B | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x007EA4BB | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007EA538 | `NikePlus_SensorSearching_Screen'` | Known | Screen layout |
| 0x007EA55B | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x007EA5CA | `NikePlus_Playlists_Screen ` | Known | Screen layout |
| 0x007EA5E7 | `NikePlus_Playlists_Screen_Nested` | Known | Screen layout |
| 0x007EA67B | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x007EA6DB | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007EA758 | `NikePlus_SensorSearching_Screen'` | Known | Screen layout |
| 0x007EA77B | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x007EA7EA | `NikePlus_Playlists_Screen ` | Known | Screen layout |
| 0x007EA807 | `NikePlus_Playlists_Screen_Nested` | Known | Screen layout |
| 0x007EA8D0 | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x007EA930 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007EA9AD | `NikePlus_Playlists_Screen!` | Known | Screen layout |
| 0x007EA9CA | `NikePlus_Playlists_Screen_Default` | Known | Screen layout |
| 0x007EAA36 | `NikePlus_SensorSearching_Screen'` | Known | Screen layout |
| 0x007EAA59 | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x007EABF5 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007EAC13 | `NikePlus_NowRunning_Screen_Basic'` | Known | Screen layout |
| 0x007EAC87 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EACA5 | `NikePlus_NowRunning_Screen_Calories'` | Known | Screen layout |
| 0x007EAD1C | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EAD3A | `NikePlus_NowRunning_Screen_Distance#` | Known | Screen layout |
| 0x007EADAD | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EADCB | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007EAE33 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007EAF0F | `NikePlus_NowRunning_Screen$` | Known | Screen layout |
| 0x007EAF2D | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x007EAFF7 | `NikePlus_NowRunning_Screen$` | Known | Screen layout |
| 0x007EB015 | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x007EB0DF | `NikePlus_NowRunning_Screen$` | Known | Screen layout |
| 0x007EB0FD | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x007EB3C2 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EB3EE | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EB470 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EB49E | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EB520 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EB542 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EB5B4 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EB5D7 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EB647 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EB665 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EB6D5 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EB6FB | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EB76F | `NikePlus_NowRunning_Screen_Basic_PowerSong` | Known | Screen layout |
| 0x007EBAFC | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EBB28 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EBBAA | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EBBD8 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EBC5A | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EBC7C | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EBCEE | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EBD11 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EBD81 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EBD9F | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EBE0F | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EBE35 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EBEA6 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007EC226 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EC252 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EC2D4 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EC302 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EC384 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EC3A6 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EC418 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EC43B | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EC4AB | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EC4C9 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EC539 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EC55F | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EC5D3 | `NikePlus_NowRunning_Screen_Time_PowerSong` | Known | Screen layout |
| 0x007EC960 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EC98C | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007ECA0E | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007ECA3C | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007ECABE | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007ECAE0 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007ECB52 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007ECB75 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007ECBE5 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007ECC03 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007ECC73 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007ECC99 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007ECD0A | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007ED08A | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007ED0B6 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007ED138 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007ED166 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007ED1E8 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007ED20A | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007ED27C | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007ED29F | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007ED30F | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007ED32D | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007ED39D | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007ED3C3 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007ED437 | `NikePlus_NowRunning_Screen_Distance_PowerSong` | Known | Screen layout |
| 0x007ED7C8 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007ED7F4 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007ED876 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007ED8A4 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007ED926 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007ED948 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007ED9BA | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007ED9DD | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EDA4D | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EDA6B | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EDADB | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EDB01 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EDB72 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007EDEF6 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EDF22 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EDFA4 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EDFD2 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EE054 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EE076 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EE0E8 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EE10B | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EE17B | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EE199 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EE209 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EE22F | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EE2A3 | `NikePlus_NowRunning_Screen_Calories_PowerSong` | Known | Screen layout |
| 0x007EE634 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EE660 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EE6E2 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EE710 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EE792 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EE7B4 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EE826 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EE849 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EE8B9 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EE8D7 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EE947 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EE96D | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EE9DE | `NikePlus_NowRunning_Screen_Calories` | Known | Screen layout |
| 0x007EED2C | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EED58 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007EEDDA | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EEE08 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007EEE8A | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007EEEAC | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007EEF1E | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007EEF41 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007EEFB1 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007EEFCF | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007EF03F | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EF065 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EF200 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EF267 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EF2DB | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EF34E | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007EF3BA | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007EF3DB | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007EF456 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EF47C | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EF539 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007EF565 | `NikePlus_CalibrationCompleteError_Screen_Default'` | Known | Screen layout |
| 0x007EF5E9 | `NikePlus_CalibrationCompleteError_Screen*` | Known | Screen layout |
| 0x007EF615 | `NikePlus_CalibrationComplete_Screen_Pacing%` | Known | Screen layout |
| 0x007EF691 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007EF6BF | `NikePlus_CalibrationCompleteSuccess_Screen_Default` | Known | Screen layout |
| 0x007EF738 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EF7C8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007EF81B | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x007EF888 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007EF8DB | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x007EF948 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007EF99C | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007EFA02 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007EFA20 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007EFA8C | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EFAAA | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007EFB1A | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EFB38 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007EFBA4 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EFBC2 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007EFC6F | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007EFC95 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007EFD28 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007EFD42 | `NikePlus_Custom_Screen_Time` | Known | Screen layout |
| 0x007EFDC3 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007EFDE4 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007EFE77 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007EFE91 | `NikePlus_Custom_Screen_Distance` | Known | Screen layout |
| 0x007EFF19 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007EFF3A | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007EFFB7 | `NikePlus_EquipmentAlert_Screen)` | Known | Screen layout |
| 0x007F0050 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F0071 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F00FC | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007F0116 | `NikePlus_Custom_Screen_Calories` | Known | Screen layout |
| 0x007F01C9 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F0250 | `NikePlus_Alert_Screen(` | Known | Screen layout |
| 0x007F02F9 | `NikePlus_EquipmentAlert_ScreenK` | Known | Screen layout |
| 0x007F03BC | `NikePlus_EquipmentAlert_Screen8` | Known | Screen layout |
| 0x007F0470 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007F0560 | `NikePlus_EquipmentAlert_Screen>` | Known | Screen layout |
| 0x007F061E | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007F06DD | `NikePlus_Alert_Screen$` | Known | Screen layout |
| 0x007F075E | `NikePlus_Remote_Unlinking_Screen(` | Known | Screen layout |
| 0x007F0782 | `NikePlus_Remote_Unlinking_Screen_Default!` | Known | Screen layout |
| 0x007F07F8 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x007F0936 | `NikePlus_Calibration_CalibrateWalk_Screen1` | Known | Screen layout |
| 0x007F09DA | `NikePlus_Calibration_CalibrateRun_Screen0` | Known | Screen layout |
| 0x007F0A9D | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007F0B5E | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F0B7F | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F0C06 | `NikePlus_Custom_Screen*` | Known | Screen layout |
| 0x007F0C20 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x007F0D11 | `NikePlus_Calibrate_ResetToDefault_Screen0` | Known | Screen layout |
| 0x007F0D3D | `NikePlus_Calibrate_ResetToDefault_Screen_Walking` | Known | Screen layout |
| 0x007F0DF3 | `NikePlus_Calibration_CalibrateRun_Screen"` | Known | Screen layout |
| 0x007F0E6B | `NikePlus_Calibration_CalibrateWalk_Screen` | Known | Screen layout |
| 0x007F0F17 | `NikePlus_Calibration_CalibrateRun_Screen"` | Known | Screen layout |
| 0x007F0F8F | `NikePlus_Calibration_CalibrateWalk_Screen` | Known | Screen layout |
| 0x007F101D | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007F10DE | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F10FF | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F1186 | `NikePlus_Custom_Screen*` | Known | Screen layout |
| 0x007F11A0 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x007F128F | `NikePlus_Calibrate_ResetToDefault_Screen0` | Known | Screen layout |
| 0x007F12BB | `NikePlus_Calibrate_ResetToDefault_Screen_Running` | Known | Screen layout |
| 0x007F132D | `NikePlus_StartWorkout_Screen$` | Known | Screen layout |
| 0x007F134D | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x007F13B4 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007F1407 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007F147C | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007F14D6 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007F158D | `NikePlus_Custom_Screen!` | Known | Screen layout |
| 0x007F1609 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007F1680 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007F1712 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x007F1790 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007F17EA | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007F1889 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007F190A | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007F1961 | `NikePlus_Calibration_ChooseCalibration_Screen5` | Known | Screen layout |
| 0x007F1A0E | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F1A8D | `NikePlus_StartCalibration_Screen(` | Known | Screen layout |
| 0x007F1AB1 | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x007F1B17 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F1B93 | `NikePlus_StartWorkout_Screen$` | Known | Screen layout |
| 0x007F1BB3 | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x007F1C1E | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F1C97 | `NikePlus_Workout_Music_Screen` | Known | Screen layout |
| 0x007F1CFE | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007F1D59 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F1E05 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F1E97 | `NikePlus_StartWorkout_Screen$` | Known | Screen layout |
| 0x007F1EB7 | `NikePlus_StartWorkout_Screen_Default#` | Known | Screen layout |
| 0x007F1F2B | `NikePlus_StartCalibration_Screen(` | Known | Screen layout |
| 0x007F1F4F | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x007F1FBC | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007F2044 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x007F20D3 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F20F4 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F2191 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F21B2 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F2251 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F2272 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F230D | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F232E | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F23FC | `NikePlus_EquipmentAlert_Screen)` | Known | Screen layout |
| 0x007F2495 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F24B6 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F2552 | `NikePlus_History_BestWorkouts_Screen,` | Known | Screen layout |
| 0x007F257A | `NikePlus_History_BestWorkouts_Screen_Default#` | Known | Screen layout |
| 0x007F25F6 | `NikePlus_History_RecentWorkouts_Screen.` | Known | Screen layout |
| 0x007F2620 | `NikePlus_History_RecentWorkouts_Screen_Default'` | Known | Screen layout |
| 0x007F26A2 | `NikePlus_History_WorkoutSummary_Screen+` | Known | Screen layout |
| 0x007F26CC | `NikePlus_History_WorkoutSummary_Screen_Last1` | Known | Screen layout |
| 0x007F2755 | `NikePlus_NoData_Screen%` | Known | Screen layout |
| 0x007F276F | `NikePlus_NoData_Screen_NoBestWorkouts2` | Known | Screen layout |
| 0x007F27F3 | `NikePlus_NoData_Screen&` | Known | Screen layout |
| 0x007F280D | `NikePlus_NoData_Screen_NoSavedWorkouts` | Known | Screen layout |
| 0x007F291D | `NikePlus_History_Totals_Screen&` | Known | Screen layout |
| 0x007F293F | `NikePlus_History_Totals_Screen_Default` | Known | Screen layout |
| 0x007F29B1 | `NikePlus_History_DeleteActiveWorkout_Screen2` | Known | Screen layout |
| 0x007F29E0 | `NikePlus_History_DeleteActiveWorkout_Screen_Cancel` | Known | Screen layout |
| 0x007F2A55 | `NikePlus_History_DeleteActiveWorkout_Screen7` | Known | Screen layout |
| 0x007F2A84 | `NikePlus_History_DeleteActiveWorkout_Screen_LastWorkout` | Known | Screen layout |
| 0x007F2AFC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007F2B4F | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007F2BA5 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007F2C60 | `NikePlus_History_RecentWorkouts_ScreenB` | Known | Screen layout |
| 0x007F2CF6 | `NikePlus_History_WorkoutSummary_ScreenA` | Known | Screen layout |
| 0x007F2D8B | `NikePlus_History_Screen` | Known | Screen layout |
| 0x007F2E57 | `NikePlus_History_RecentWorkouts_ScreenB` | Known | Screen layout |
| 0x007F2EED | `NikePlus_History_WorkoutSummary_ScreenA` | Known | Screen layout |
| 0x007F2F82 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x007F303F | `NikePlus_History_ScreenG` | Known | Screen layout |
| 0x007F30CB | `NikePlus_History_RecentWorkouts_Screen` | Known | Screen layout |
| 0x007F314B | `NikePlus_History_DeleteAllWorkouts_Screen0` | Known | Screen layout |
| 0x007F3178 | `NikePlus_History_DeleteAllWorkouts_Screen_Cancel#` | Known | Screen layout |
| 0x007F31F8 | `NikePlus_History_WorkoutSummary_Screen.` | Known | Screen layout |
| 0x007F3222 | `NikePlus_History_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007F32D5 | `NikePlus_History_ClearTotals_Screen+` | Known | Screen layout |
| 0x007F32FC | `NikePlus_History_ClearTotals_Screen_Default` | Known | Screen layout |
| 0x007F339E | `NikePlus_History_Totals_Screen` | Known | Screen layout |
| 0x007F3431 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007F3452 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007F34C1 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007F34DF | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007F354B | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F3569 | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007F35D9 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F35F7 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007F3663 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007F3681 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007F3717 | `NikePlus_Dynamic_Workout_Screen/` | Known | Screen layout |
| 0x007F373A | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x007F37B3 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007F37D1 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007F383D | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F385B | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007F38CB | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F38E9 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007F3955 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007F3973 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007F3A0B | `NikePlus_Dynamic_Workout_Screen,` | Known | Screen layout |
| 0x007F3A2E | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x007F3AA4 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007F3AC2 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007F3B2E | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F3B4C | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007F3BBC | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F3BDA | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007F3C46 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007F3C64 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007F3CFB | `NikePlus_Dynamic_Workout_Screen.` | Known | Screen layout |
| 0x007F3D1E | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x007F3D96 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007F3DB4 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007F3E20 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F3E3E | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007F3EAE | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007F3ECC | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007F3F38 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007F3F56 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007F3FC6 | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x007F3FDF | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x007F4042 | `DemoMode_Screen` | Known | Screen layout |
| 0x007F4055 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x007F40C2 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x007F40DB | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x007F414E | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x007F4169 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x007F4279 | `VoiceMemos_Context_PlayDelete_Screen)` | Known | Screen layout |
| 0x007F42A1 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x007F4318 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007F43E4 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007F4453 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007F4541 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007F45AA | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007F45CC | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007F4638 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007F465A | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007F47D6 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007F47F2 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007F48B9 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007F48D4 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007F4937 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007F499A | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007F4A31 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007F4A4D | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007F4B14 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007F4B2F | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007F4B92 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007F4BF5 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007F4C8D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007F4CA9 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007F4D70 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007F4D8B | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007F4DEE | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007F4E51 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007F4ECE | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007F4F39 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007F4FA5 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007F5017 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007F5084 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007F50EF | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x007F515B | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007F51C3 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007F522F | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007F52A3 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007F5311 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x007F538A | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x00817F24 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x00817FA9 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x00818296 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x009CCDF6 | `VoiceMemos_Screen_InsufficientDiskSpace2` | Known | Screen layout |
| 0x009CE796 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x009CE7AE | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x009CE7CC | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x009CE893 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x009CE911 | `NowPlaying_Screen_Music` | Known | Screen layout |
| 0x009CE952 | `MainMenu_Music_Screen_NoMusic` | Known | Screen layout |
| 0x009CE970 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x009CE98E | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x009CE9A7 | `NikePlus_Custom_Screen_Weight_ToWorkoutMusic` | Known | Screen layout |
| 0x009CEABC | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x009CEB70 | `DeleteRental_Confirmation_Screen_Unwatched` | Known | Screen layout |
| 0x009CEBC6 | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x009CEC12 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x009CEE75 | `NowPlaying_Screen_ExtraInfoLoadFailed` | Known | Screen layout |
| 0x009CEED0 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x009CEEE9 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x009CEF07 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x009CEF36 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x009CEF6E | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x009CEFFE | `NikePlus_Playlists_Screen_Nested` | Known | Screen layout |
| 0x009CF3DE | `Video_Settings_TV_Screen_Standard` | Known | Screen layout |
| 0x009CF410 | `Search_Main_Screen_WithKeyboard` | Known | Screen layout |
| 0x009CF430 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x009CF475 | `VoiceMemos_Screen_InsufficientDiskSpace` | Known | Screen layout |
| 0x009CF4DA | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x009CF4FE | `NikePlus_Custom_Screen_Distance` | Known | Screen layout |
| 0x009CF559 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x009CF5E0 | `Video_Settings_TV_Screen_Wide` | Known | Screen layout |
| 0x009CF628 | `CoverFlow_Screen_Backside` | Known | Screen layout |
| 0x009D2C6A | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x009D2E6F | `NowPlaying_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x009D2E94 | `Slideshow_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x009D2F64 | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x009D2F7E | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x009D3076 | `RentalDeleted_Screen_Title` | Known | Screen layout |
| 0x009D3091 | `SingleRentalExpiring_Screen_Title` | Known | Screen layout |
| 0x009D30B3 | `MultipleRentalsExpiring_Screen_Title` | Known | Screen layout |
| 0x009D30D8 | `DeleteRental_Screen_Title` | Known | Screen layout |
| 0x009D317B | `VoiceMemos_DeleteAll_ScreenTitle` | Known | Screen layout |
| 0x009D3218 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x009D325B | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x009D334D | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x009D336D | `NikePlus_Custom_Screen_Time` | Known | Screen layout |
| 0x009D34B8 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x009D35A1 | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x009D35BA | `Radio_Screen_Volume` | Known | Screen layout |
| 0x009D35CE | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x009D35EB | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x009D360A | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x009D3715 | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x009D3881 | `CoverFlow_Screen_BacksideForCapture` | Known | Screen layout |
| 0x009D4A56 | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x009D4B4E | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x009D4B69 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x009D4E0C | `NikePlus_CalibrationComplete_Screen_Pacing` | Known | Screen layout |
| 0x009D4EA4 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x009D4ED8 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x009D4F15 | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x009D5027 | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x009D5155 | `NikePlus_Calibrate_ResetToDefault_Screen_Walking` | Known | Screen layout |
| 0x009D5287 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x009D52A0 | `NikePlus_Calibrate_ResetToDefault_Screen_Running` | Known | Screen layout |
| 0x009D52F0 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x009D5316 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x009DAFA9 | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x009DB018 | `TVOutSlideshow_Screen_Playing` | Known | Screen layout |
| 0x009DB036 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x009DB08C | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x009DB0F6 | `NikePlus_NowRunning_Screen_Basic_PowerSong` | Known | Screen layout |
| 0x009DB121 | `NikePlus_NowRunning_Screen_Distance_PowerSong` | Known | Screen layout |
| 0x009DB14F | `NikePlus_NowRunning_Screen_Time_PowerSong` | Known | Screen layout |
| 0x009DB19C | `NikePlus_NowRunning_Screen_Calories_PowerSong` | Known | Screen layout |
| 0x009DB219 | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x009DB284 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x009DB304 | `Extras_Screen_Debug` | Known | Screen layout |
| 0x009DB40E | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x009DB42E | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x009DB9A0 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x009DB9BB | `Extras_Screen_Lock` | Known | Screen layout |
| 0x009DB9CE | `MainMenu_List_ScreenLock` | Known | Screen layout |
| 0x009DB9E7 | `ExtrasMenu_ScreenLock` | Known | Screen layout |
| 0x009DBA6A | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x009DBA8B | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x009DBB36 | `NikePlus_StartCalibration_Screen_Walk` | Known | Screen layout |
| 0x009DBBBE | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x009DBBE0 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x009DBCE7 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x009DBD27 | `VoiceMemos_Screen_DeletAllAsk` | Known | Screen layout |
| 0x009DBD45 | `NowPlaying_Screen_TVOutAsk` | Known | Screen layout |
| 0x009DBEA1 | `NowPlaying_Screen_Initial` | Known | Screen layout |
| 0x009DBEBB | `NowPlaying_Screen_Video_TVOut_Initial` | Known | Screen layout |
| 0x009DC18D | `NikePlus_History_DeleteAllWorkouts_Screen_Cancel` | Known | Screen layout |
| 0x009DC1BE | `NikePlus_History_DeleteActiveWorkout_Screen_Cancel` | Known | Screen layout |
| 0x009DCFE0 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x009DD061 | `RemoteUI_Screen` | Known | Screen layout |
| 0x009DD071 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x009DD089 | `Video_Settings_TV_Screen` | Known | Screen layout |
| 0x009DD0A2 | `NikePlus_NoData_Screen` | Known | Screen layout |
| 0x009DD0B9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x009DD0D0 | `NikePlus_Workout_Music_Screen` | Known | Screen layout |
| 0x009DD0EE | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x009DD112 | `AddressViewer_PartialLoad_Screen` | Known | Screen layout |
| 0x009DD133 | `NikePlus_ActivityStopped_Screen` | Known | Screen layout |
| 0x009DD153 | `Alarms_Alarm_Clock_Triggered_Screen` | Known | Screen layout |
| 0x009DD177 | `Alarms_Alarm_Triggered_Screen` | Known | Screen layout |
| 0x009DD195 | `Unsupported_Screen` | Known | Screen layout |
| 0x009DD1A8 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x009DD1C6 | `LockediPod_Screen` | Known | Screen layout |
| 0x009DD1D8 | `DiskMode_Screen` | Known | Screen layout |
| 0x009DD1E8 | `DemoMode_Screen` | Known | Screen layout |
| 0x009DD1F8 | `Notes_Image_Screen` | Known | Screen layout |
| 0x009DD20B | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x009DD229 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x009DD23F | `NowPlaying_Idle_Screen` | Known | Screen layout |
| 0x009DD256 | `Game_Screen` | Known | Screen layout |
| 0x009DD262 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x009DD27F | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x009DD298 | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x009DD2B9 | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x009DD2DE | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x009DD2F1 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x009DD30E | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x009DD32F | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x009DD354 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x009DD36B | `Notes_Loading_Screen` | Known | Screen layout |
| 0x009DD380 | `NikePlus_SensorSearching_Screen` | Known | Screen layout |
| 0x009DD3A0 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x009DD3BF | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x009DD3D7 | `NikePlus_Remote_Unlinking_Screen` | Known | Screen layout |
| 0x009DD3F8 | `SettingsMenus_AdjustScrolling_Screen` | Known | Screen layout |
| 0x009DD41D | `Game_Running_Screen` | Known | Screen layout |
| 0x009DD431 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x009DD44C | `Stopwatch_Screen` | Known | Screen layout |
| 0x009DD45D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x009DD474 | `Clock_Screen` | Known | Screen layout |
| 0x009DD481 | `NikePlus_Calibration_CalibrateWalk_Screen` | Known | Screen layout |
| 0x009DD4AB | `TVOutSlideshowAsk_Screen` | Known | Screen layout |
| 0x009DD4C4 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x009DD4DA | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x009DD4F8 | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x009DD514 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x009DD525 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x009DD53C | `DemoMode_Main_Screen` | Known | Screen layout |
| 0x009DD551 | `Search_Main_Screen` | Known | Screen layout |
| 0x009DD564 | `AddressViewer_Main_Screen` | Known | Screen layout |
| 0x009DD57E | `Speakers_Main_Screen` | Known | Screen layout |
| 0x009DD593 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x009DD5A9 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x009DD5C3 | `Clock_Region_Screen` | Known | Screen layout |
| 0x009DD5D7 | `RentalDeleted_Notification_Screen` | Known | Screen layout |
| 0x009DD5F9 | `SingleRentalExpiring_Notification_Screen` | Known | Screen layout |
| 0x009DD622 | `MultipleRentalsExpiring_Notification_Screen` | Known | Screen layout |
| 0x009DD64E | `RentalError_Notification_Screen` | Known | Screen layout |
| 0x009DD66E | `DeleteRental_Confirmation_Screen` | Known | Screen layout |
| 0x009DD68F | `LockConfirmation_Screen` | Known | Screen layout |
| 0x009DD6A7 | `NikePlus_EndCalibration_Screen` | Known | Screen layout |
| 0x009DD6C6 | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x009DD6F4 | `NikePlus_StartCalibration_Screen` | Known | Screen layout |
| 0x009DD715 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x009DD733 | `NikePlus_Calibration_CalibrateRun_Screen` | Known | Screen layout |
| 0x009DD75C | `Video_Settings_Fit_To_Screen` | Known | Screen layout |
| 0x009DD779 | `RentalInfo_Screen` | Known | Screen layout |
| 0x009DD78B | `Radio_Screen` | Known | Screen layout |
| 0x009DD798 | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x009DD7B2 | `ChargingMode_LowPower_Screen` | Known | Screen layout |
| 0x009DD7CF | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x009DD7E9 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x009DD803 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x009DD81D | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x009DD836 | `NikePlus_CalibrationCompleteError_Screen` | Known | Screen layout |
| 0x009DD85F | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x009DD876 | `Extras_Screen` | Known | Screen layout |
| 0x009DD884 | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x009DD8A1 | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x009DD8C3 | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x009DD8DC | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x009DD8FA | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x009DD913 | `Video_Settings_Screen` | Known | Screen layout |
| 0x009DD929 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x009DD942 | `SettingsMenus_AudiobookSettings_Screen` | Known | Screen layout |
| 0x009DD969 | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x009DD98F | `PhotosSettings_Screen` | Known | Screen layout |
| 0x009DD9A5 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x009DD9BD | `NikePlus_Songs_Screen` | Known | Screen layout |
| 0x009DD9D3 | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x009DD9F6 | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x009DDA13 | `MediaLists_Rentals_Screen` | Known | Screen layout |
| 0x009DDA2D | `NikePlus_History_Totals_Screen` | Known | Screen layout |
| 0x009DDA4C | `NikePlus_History_ClearTotals_Screen` | Known | Screen layout |
| 0x009DDA70 | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x009DDA94 | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x009DDAAD | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x009DDACF | `ToDo_Instructions_Screen` | Known | Screen layout |
| 0x009DDAE8 | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x009DDB04 | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x009DDB1E | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x009DDB3F | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x009DDB5B | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x009DDB73 | `VoiceMemos_Screen` | Known | Screen layout |
| 0x009DDB85 | `No_Photos_Screen` | Known | Screen layout |
| 0x009DDB96 | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x009DDBB0 | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x009DDBCC | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x009DDBF0 | `NikePlus_CalibrationCompleteSuccess_Screen` | Known | Screen layout |
| 0x009DDC1B | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x009DDC3B | `ChargingMode_Progress_Screen` | Known | Screen layout |
| 0x009DDC58 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x009DDC6E | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x009DDC89 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x009DDCA5 | `NikePlus_Playlists_Screen` | Known | Screen layout |
| 0x009DDCBF | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x009DDCE1 | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x009DDD02 | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x009DDD1C | `NikePlus_History_DeleteAllWorkouts_Screen` | Known | Screen layout |
| 0x009DDD46 | `NikePlus_History_RecentWorkouts_Screen` | Known | Screen layout |
| 0x009DDD6D | `NikePlus_History_BestWorkouts_Screen` | Known | Screen layout |
| 0x009DDD92 | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x009DDDAC | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x009DDDCB | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x009DDDEC | `NikePlus_Calibrate_ResetToDefault_Screen` | Known | Screen layout |
| 0x009DDE15 | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x009DDE2D | `NoContent_Screen` | Known | Screen layout |
| 0x009DDE3E | `Calendar_Event_Screen` | Known | Screen layout |
| 0x009DDE54 | `FirstBoot_Screen` | Known | Screen layout |
| 0x009DDE65 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x009DDE7B | `NikePlus_EquipmentAlert_Screen` | Known | Screen layout |
| 0x009DDE9A | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x009DDEB0 | `Notes_List_Screen` | Known | Screen layout |
| 0x009DDEC2 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x009DDED8 | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x009DDEF9 | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x009DDF13 | `NikePlus_Dynamic_Workout_Screen` | Known | Screen layout |
| 0x009DDF33 | `NikePlus_EndPausedWorkout_Screen` | Known | Screen layout |
| 0x009DDF54 | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x009DDF6F | `NikePlus_ResumeWorkout_Screen` | Known | Screen layout |
| 0x009DDF8D | `NikePlus_History_DeleteActiveWorkout_Screen` | Known | Screen layout |
| 0x009DDFB9 | `NikePlus_StartWorkout_Screen` | Known | Screen layout |
| 0x009DDFD6 | `TODOS_Menu_Screen` | Known | Screen layout |
| 0x009DDFE8 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x009DDFFE | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x009DE01A | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x009DE02F | `Games_Menu_Screen` | Known | Screen layout |
| 0x009DE041 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x009DE054 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x009DE073 | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x009DE092 | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x009DE0B6 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x009DE0CC | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x009DE0EA | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x009DE10D | `TVOutSlideshow_Screen` | Known | Screen layout |
| 0x009DE123 | `CoverFlow_Screen` | Known | Screen layout |
| 0x009DE134 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x009DE148 | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x009DE16A | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x009DE182 | `Stopwatch_SessionSummary_Screen` | Known | Screen layout |
| 0x009DE1A2 | `NikePlus_End_WorkoutSummary_Screen` | Known | Screen layout |
| 0x009DE1C5 | `NikePlus_History_WorkoutSummary_Screen` | Known | Screen layout |
| 0x009DE1EC | `SettingsMenu_AboutWithAccessory_Screen` | Known | Screen layout |
| 0x009DE213 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x009DE22B | `Clock_NorthAmerica_City_Screen` | Known | Screen layout |
| 0x009DE24A | `Clock_SouthAmerica_City_Screen` | Known | Screen layout |
| 0x009DE269 | `Clock_Africa_City_Screen` | Known | Screen layout |
| 0x009DE282 | `Clock_Australia_City_Screen` | Known | Screen layout |
| 0x009DE29E | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x009DE2B5 | `Clock_Pacific_City_Screen` | Known | Screen layout |
| 0x009DE2CF | `Clock_Atlantic_City_Screen` | Known | Screen layout |
| 0x009DE2EA | `Clock_Europe_City_Screen` | Known | Screen layout |
| 0x009DE3DE | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x009DE42F | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x009DE452 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x009DE47A | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x009DE806 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x009DE909 | `RentalInfo_Screen_ExpiringSoon` | Known | Screen layout |
| 0x009DE95F | `RentalInfo_Screen_NoAlbumArt_ExpiringSoon` | Known | Screen layout |
| 0x009DEA69 | `NikePlus_StartCalibration_Screen_Run` | Known | Screen layout |
| 0x009DED53 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x009DEDA9 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x009DEEFA | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x009DEF17 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x009DF348 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x009DF46A | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x009DF48C | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x009DF5C4 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x009DF5E3 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x009DFCD3 | `ChargingMode_ScreenLayout_LowPower` | Known | Screen layout |
| 0x009E0693 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x009E07F1 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x009E08A8 | `NikePlus_NowRunning_Screen_Calories` | Known | Screen layout |
| 0x009E08CC | `NikePlus_Custom_Screen_Calories` | Known | Screen layout |
| 0x009E0965 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x009E0983 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x009E09A3 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x009E0AAE | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x009E0ACA | `Extras_Screen_Games` | Known | Screen layout |
| 0x009E0BD0 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x009E0BEF | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x009E0C0B | `Extras_Screen_Notes` | Known | Screen layout |
| 0x009E0CF6 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x009E0E12 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x009E0FE0 | `MainMenu_Music_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009E1003 | `MainMenus_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009E1026 | `NoContent_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009E1060 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x009E107F | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x009E10A0 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x009E1197 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x009E11B4 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x009E1233 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x009E1317 | `MainMenu_Music_Screen_NoCompilations` | Known | Screen layout |
| 0x009E133C | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x009E14E5 | `MainMenu_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x009E1508 | `MainMenu_Videos_Screen_NoMusicVideos` | Known | Screen layout |
| 0x009E152D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x009E154C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x009E156B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x009E158C | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x009E15CA | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x009E15EB | `Extras_Screen_VoiceMemos` | Known | Screen layout |
| 0x009E1656 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x009E1688 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x009E16A7 | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x009E1754 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x009E17C0 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x009E18B9 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x009E18D5 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x009E1958 | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x009E1973 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x009E1994 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x009E1A43 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x009E1A77 | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x009E1A98 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x009E1B56 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x009E1B77 | `MainMenu_Videos_Screen_NoPlaylists` | Known | Screen layout |
| 0x009E1B9A | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x009E1BE9 | `MainMenu_Music_Screen_NoArtists` | Known | Screen layout |
| 0x009E1C59 | `NikePlus_NoData_Screen_NoSavedWorkouts` | Known | Screen layout |
| 0x009E1D0B | `NikePlus_NoData_Screen_NoBestWorkouts` | Known | Screen layout |
| 0x009E1DB8 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x009E1DD7 | `Extras_Screen_NikePlus` | Known | Screen layout |
| 0x009E1F27 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x009E1F46 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x009E1F67 | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x009E243E | `NikePlus_Custom_Screen_Weight` | Known | Screen layout |
| 0x009E24B3 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x009E2566 | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x009E25E0 | `Notes_Contents_Screen_Alt` | Known | Screen layout |
| 0x009E25FA | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x009E26A6 | `VoiceMemos_Screen_InsufficientDiskSpace2_Default` | Known | Screen layout |
| 0x009E2758 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x009E27FD | `VoiceMemos_Screen_InsufficientDiskSpace_Default` | Known | Screen layout |
| 0x009E282D | `NowPlaying_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x009E285A | `Slideshow_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x009E3734 | `Radio_Screen_Tuning_Default` | Known | Screen layout |
| 0x009E37C0 | `VoiceMemos_Screen_DeletAllAsk_Default` | Known | Screen layout |
| 0x009E37E6 | `NowPlaying_Screen_TVOutAsk_Default` | Known | Screen layout |
| 0x009E381D | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x009E3843 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x009E3861 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x009E388D | `AddressViewer_PartialLoad_Screen_Default` | Known | Screen layout |
| 0x009E38B6 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x009E38DE | `Alarms_Alarm_Clock_Triggered_Screen_Default` | Known | Screen layout |
| 0x009E390A | `Alarms_Alarm_Triggered_Screen_Default` | Known | Screen layout |
| 0x009E3930 | `Unsupported_Screen_Default` | Known | Screen layout |
| 0x009E394B | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x009E3971 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x009E3989 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x009E39A4 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x009E39C1 | `Game_Screen_Default` | Known | Screen layout |
| 0x009E39D5 | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x009E39FB | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x009E3A1C | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x009E3A45 | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x009E3A6F | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x009E3A9C | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x009E3AC5 | `Notes_Loading_Screen_Default` | Known | Screen layout |
| 0x009E3AE2 | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x009E3B0A | `NikePlus_Remote_Unlinking_Screen_Default` | Known | Screen layout |
| 0x009E3B33 | `Clock_Screen_Default` | Known | Screen layout |
| 0x009E3B48 | `TVOutSlideshowAsk_Screen_Default` | Known | Screen layout |
| 0x009E3B69 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x009E3B87 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x009E3BAD | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x009E3BD1 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x009E3BEA | `AddressViewer_Main_Screen_Default` | Known | Screen layout |
| 0x009E3C0C | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x009E3C29 | `MainMenus_Main_Screen_Default` | Known | Screen layout |
| 0x009E3C47 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x009E3C64 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x009E3C80 | `RentalDeleted_Notification_Screen_Default` | Known | Screen layout |
| 0x009E3CAA | `SingleRentalExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x009E3CDB | `MultipleRentalsExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x009E3D0F | `RentalError_Notification_Screen_Default` | Known | Screen layout |
| 0x009E3D37 | `DeleteRental_Confirmation_Screen_Default` | Known | Screen layout |
| 0x009E3D60 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x009E3D8C | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x009E3DB3 | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x009E3DDC | `RentalInfo_Screen_Default` | Known | Screen layout |
| 0x009E3DF6 | `Radio_Screen_Default` | Known | Screen layout |
| 0x009E3E0B | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x009E3E2D | `NikePlus_CalibrationCompleteError_Screen_Default` | Known | Screen layout |
| 0x009E3E5E | `Extras_Screen_Default` | Known | Screen layout |
| 0x009E3E74 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x009E3E9A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x009E3EBB | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x009E3ED9 | `NikePlus_Settings_Screen_Default` | Known | Screen layout |
| 0x009E3EFA | `NikePlus_Songs_Screen_Default` | Known | Screen layout |
| 0x009E3F18 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x009E3F3A | `NikePlus_History_Totals_Screen_Default` | Known | Screen layout |
| 0x009E3F61 | `NikePlus_History_ClearTotals_Screen_Default` | Known | Screen layout |
| 0x009E3F8D | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x009E3FB9 | `ToDo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009E3FDA | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009E3FFE | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x009E4020 | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x009E4044 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x009E4063 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x009E407C | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x009E409E | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x009E40C2 | `NikePlus_CalibrationCompleteSuccess_Screen_Default` | Known | Screen layout |
| 0x009E40F5 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x009E4113 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x009E4137 | `NikePlus_Playlists_Screen_Default` | Known | Screen layout |
| 0x009E4159 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x009E4183 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x009E41AC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x009E41CE | `NikePlus_History_RecentWorkouts_Screen_Default` | Known | Screen layout |
| 0x009E41FD | `NikePlus_History_BestWorkouts_Screen_Default` | Known | Screen layout |
| 0x009E422A | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x009E424A | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x009E4268 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x009E4281 | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x009E429F | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x009E42B9 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x009E42D7 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x009E4300 | `NikePlus_EndWorkout_Screen_Default` | Known | Screen layout |
| 0x009E4323 | `NikePlus_ResumeWorkout_Screen_Default` | Known | Screen layout |
| 0x009E4349 | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x009E436E | `TODOS_Menu_Screen_Default` | Known | Screen layout |
| 0x009E4388 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x009E43A6 | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x009E43C3 | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x009E43DD | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x009E43F8 | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x009E4417 | `Debug_MainMenu_Screen_Default` | Known | Screen layout |
| 0x009E4435 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x009E4453 | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x009E446C | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x009E4488 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x009E44B2 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x009E44D2 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x009E44FA | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x009E4525 | `NikePlus_History_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x009E4554 | `NikePlus_History_Screen_Default` | Known | Screen layout |
| 0x009E4574 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009E459B | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009E45C2 | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x009E45E3 | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x009E4607 | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x009E4626 | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x009E4648 | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x009E466B | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x009E46A8 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x009E4736 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x009E4766 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x009E4788 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x009E47F9 | `RentalInfo_Screen_NoAlbumArt_Default` | Known | Screen layout |
| 0x009E481E | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x009E4F4A | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009E4F76 | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009E4FBB | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x009E4FE3 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x009E5004 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x009E5025 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x009E504B | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x009E5068 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x009E508A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x009E50AE | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x009E50D2 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x009E5137 | `NikePlus_History_WorkoutSummary_Screen_Last` | Known | Screen layout |
| 0x009E52D8 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x009E5348 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x009E5399 | `NowPlaying_Screen_Video_TVOut` | Known | Screen layout |
| 0x009E54AC | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x009E5509 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x009E5558 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x009E561F | `NikePlus_History_DeleteActiveWorkout_Screen_LastWorkout` | Known | Screen layout |
| 0x009E5768 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x009E578F | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x009E5BE7 | `NikePlus_EndPausedWorkout_Screen_BasicSave_Layout` | Known | Screen layout |
| 0x009E5C19 | `NikePlus_EndPausedWorkout_Screen_DistanceSave_Layout` | Known | Screen layout |
| 0x009E5C4E | `NikePlus_EndPausedWorkout_Screen_TimeSave_Layout` | Known | Screen layout |
| 0x009E5C7F | `NikePlus_EndPausedWorkout_Screen_CaloriesSave_Layout` | Known | Screen layout |
| 0x009E5F37 | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x009E6134 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x009E63D8 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x009E66F5 | `Lock_Screen_Default_Layout` | Known | Screen layout |
| 0x009E678B | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x009E67B2 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x009E69CE | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x009E6AA8 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x009E6B0F | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009E6B39 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009E99E1 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x009E9A2D | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x009E9B0B | `MainMenu_List_ScreenLock_x` | Known | Screen layout |
| 0x009E9DD9 | `RentalInfo_Screen_ExpiresToday` | Known | Screen layout |
| 0x009E9E2F | `RentalInfo_Screen_NoAlbumArt_ExpiresToday` | Known | Screen layout |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000904B | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | Known | RTOS |
| 0x002C2C0C | `  K - RTXC` | Known | RTOS |
| 0x002C3BF4 | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | Known | RTOS |
| 0x009CB7B0 | `Returning from RTXCBug` | Known | RTOS |

---

## 9. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000DDA20 | `HostOSTask` | Known | RTOS task thread |
| 0x0013E89C | `MP3ExampleTask` | Known | RTOS task thread |
| 0x00146E68 | `USBDeviceTask` | Known | RTOS task thread |
| 0x00151370 | `DiskReaderTask` | Known | RTOS task thread |
| 0x00160F2C | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x00160F40 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x001B86E8 | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x001FA11C | `MeCCAIOTask` | Known | RTOS task thread |
| 0x0023127C | `StreamCacheMassStorageManagerTimeOutTask` | Known | RTOS task thread |
| 0x002313F8 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x002B5CD4 | `FirewireTask` | Known | RTOS task thread |
| 0x002B5CE8 | `TouchwheelTask` | Known | RTOS task thread |
| 0x002B5CFC | `AudioOutStateTask` | Known | RTOS task thread |
| 0x002B5D28 | `DiskMgrTask` | Known | RTOS task thread |
| 0x002B5D38 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x002B5D4C | `TopPlugTask` | Known | RTOS task thread |
| 0x002B5D5C | `HPhoneDetTask` | Known | RTOS task thread |
| 0x002B5DD4 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x002B5DFC | `AlarmTask` | Known | RTOS task thread |
| 0x002B5E1B | `"USBAudioTask` | Known | RTOS task thread |
| 0x002C32AC | `Undefined Task` | Known | RTOS task thread |
| 0x003CD9A0 | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x003D1EA4 | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x003DA5F0 | `MeCCARecordingTask` | Known | RTOS task thread |
| 0x0091F290 | `WaveFileDebugTask` | Known | RTOS task thread |

---

## 10. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0026C51C | `Channel Reserved` | Known | Logging channel |
| 0x0026C530 | `Channel AppBoot` | Known | Logging channel |
| 0x0026C540 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x0026C55C | `Channel PrefsWriting` | Known | Logging channel |
| 0x0026C574 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x0026C594 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x0026C5AC | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x0026C5C8 | `Channel TestLogging` | Known | Logging channel |
| 0x0026C5DC | `Channel AppFileLoading` | Known | Logging channel |
| 0x0026C5F4 | `Channel VCardReading` | Known | Logging channel |
| 0x0026C60C | `Channel LongSongScanning` | Known | Logging channel |
| 0x0026C680 | `Channel VoiceRecording` | Known | Logging channel |
| 0x0026C698 | `Channel PhotoImporting` | Known | Logging channel |
| 0x0026C6B0 | `Channel Notes` | Known | Logging channel |
| 0x0026C6C0 | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x0026C6DC | `Channel DiskMode` | Known | Logging channel |
| 0x0026C6F0 | `Channel Firewire` | Known | Logging channel |
| 0x0026C704 | `Channel USB` | Known | Logging channel |
| 0x0026C724 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x0026C73C | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0008CC3C | `gamedata_RW` | Known | Game system |
| 0x0008CC58 | `gamedata_ShareRW` | Known | Game system |
| 0x0008CC6C | `games_RO` | Known | Game system |
| 0x009CB80A | `iPod_Control/games_RO/` | Known | Game system |
| 0x009CB821 | `Resources/Games/games_RO/` | Known | Game system |
| 0x009D887C | `GamesMenu_StatusBar_String` | Known | Game system |
| 0x009D8FB6 | `AboutScreen_Games_String` | Known | Game system |
| 0x009E0ADE | `MainMenu_List_Games` | Known | Game system |
| 0x009E0AF2 | `ExtrasMenu_Games` | Known | Game system |
| 0x009E9B7A | `MainMenu_List_Games_x` | Known | Game system |

---

## 12. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009DDC0 | `adrmmp4a` | Known | DRM system |
| 0x0014E95C | `AppleDRMVersion` | Known | DRM system |
| 0x0014E9FC | `AppleDRM` | Known | DRM system |
| 0x0014FBCC | `AppleVideoDRM` | Known | DRM system |
| 0x001530E0 | `tx3gdrmsp608aavdmp4aesds` | Known | DRM system |
| 0x0020920C | `drmttx3g` | Known | DRM system |
| 0x009CBBEF | `DRMLevel` | Known | DRM system |

---

## 13. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00035ED4 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00035EEC | `iPod_Control\iTunes\firsttime` | Known | iTunes database |
| 0x00058498 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x000584C0 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00061190 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00088C9C | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x0008CBCC | `iPod_Control/iTunes/iTunesDB.p7b` | Known | iTunes database |
| 0x000A9F24 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x000AA10C | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000B2B44 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000B3FE8 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000B40E8 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00137028 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x003C6D14 | `iTunesDB` | Known | iTunes database |
| 0x003C6D20 | `iPod_Control\iTunes\` | Known | iTunes database |

---

## 14. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0005F644 | `NANDDRIVERSIGN` | Known | Hardware |
| 0x00060180 | `NANDDRIVERSIGN` | Known | Hardware |
| 0x00060B38 | `[FTL:MSG] Apple NAND Driver (AND) 0x%08x` | Known | Hardware |
| 0x00060C50 | `[FTL:MSG] Valid Signature not found! Re-initializing NAND!` | Known | Hardware |
| 0x0013688C | `NAND FLASH DRIVE` | Known | Hardware |
| 0x0014EEA4 | `FireWireGUID` | Known | FireWire |
| 0x0014EEB4 | `FireWireVersion` | Known | FireWire |
| 0x0014F590 | `FireWire` | Known | FireWire |
| 0x002D3B4C | `[FIL:ERR] No recognized NAND found (0x%X, 0x%X) (line:%d)!` | Known | Hardware |
| 0x00925E8C | `[FTL:WRN] Recovering NAND Data Structures - this will take some time!` | Known | Hardware |
| 0x009273D0 | `[FIL:WRN]  FNAND_GetStruct 0x%X is not identified is FIL data struct identifier!` | Known | Hardware |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0075FC7A | `Settings_RadioRegionsMenu_Layout"` | Known | FM Radio |
| 0x0075FD03 | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x00816154 | `Radio Regions` | Known | FM Radio |
| 0x00870FC0 | `Radio-Regionen` | Known | FM Radio |
| 0x009D59D9 | `Settings_RadioRegions_Australia_String` | Known | FM Radio |
| 0x009D5A00 | `Settings_RadioRegions_Asia_String` | Known | FM Radio |
| 0x009D6C2E | `Settings_RadioRegions_Europe_String` | Known | FM Radio |
| 0x009D819E | `Settings_RadioRegions_Japan_String` | Known | FM Radio |
| 0x009D8DD3 | `Settings_RadioRegions_Americas_String` | Known | FM Radio |
| 0x009D94B5 | `SettingsMenu_RadioRegions_String` | Known | FM Radio |
| 0x009DCE9F | `RadioRegionsMenuItem` | Known | FM Radio |
| 0x009E12A0 | `SettingsMenu_ListItem_RadioRegions` | Known | FM Radio |
| 0x009E623D | `SettingsInfo_Template_RadioRegions_Layout` | Known | FM Radio |
| 0x009E6267 | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x009E698F | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 16. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x008AFD50 | `Fotocamera` | Known | Camera |
| 0x008AFEE0 | `Collega fotocamera o inserisci scheda` | Known | Camera |
| 0x008AFF58 | `Fotocamera non supportata` | Known | Camera |
| 0x008CE8E8 | `Camera` | Known | Camera |
| 0x008CEA74 | `Sluit camera of kaart aan` | Known | Camera |
| 0x008CEAE0 | `Camera niet ondersteund` | Known | Camera |
| 0x009D5A22 | `Photos_DevName_Camera_String` | Known | Camera |

---

## 17. Pedometer

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00817100 | `Step away from all other sensors.` | Known | Pedometer |
| 0x008172E4 | `Step away from all other remotes.` | Known | Pedometer |
| 0x009E9EE2 | `NikePlus_Remote_Step_Away` | Known | Pedometer |
| 0x009E9EFC | `NikePlus_Step_Away` | Known | Pedometer |

---

## 18. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00035EC0 | `iPod_Control` | Filesystem Path |  |
| 0x00035F2C | `iPod_Control\Device` | Filesystem Path |  |
| 0x00045940 | `iPod_Control\Device` | Filesystem Path |  |
| 0x000479CC | `iPod_Control` | Filesystem Path |  |
| 0x00048034 | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x00058478 | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path |  |
| 0x0005AFDC | `iPod_Control\Music\` | Filesystem Path |  |
| 0x00061010 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x00096E78 | `iPod_Control` | Filesystem Path |  |
| 0x00096E88 | `Resources/Games` | Filesystem Path |  |
| 0x00096E98 | `iPod_Control/%s%s%s` | Filesystem Path |  |
| 0x001011AC | `iPod_Control/Device/alarms` | Filesystem Path |  |
| 0x001114E8 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x001129EC | `iPod_Control/Device` | Filesystem Path |  |
| 0x00112A00 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00131CFC | `iPod_Control/Device/Users` | Filesystem Path |  |
| 0x00162B00 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x00162D5C | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x0016EAD8 | `Resources/UI/active.bin` | Filesystem Path |  |
| 0x0016EAF0 | `Resources/UI/` | Filesystem Path |  |
| 0x001911B8 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001928B0 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path |  |
| 0x001928D8 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001BC0A0 | `iPod_Control/Device/PlayCounts` | Filesystem Path |  |
| 0x001D344C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D34FC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3678 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3810 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D38B8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3A68 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3B0C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3BB0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3C54 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3CF8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3DA8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3E4C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3EF0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3FA0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4050 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4100 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D426C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D431C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D43CC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4470 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4520 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4614 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D46B8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D476C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4828 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D48D8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D49FC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4AB8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4B68 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4D24 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4DE8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4E98 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4F54 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5090 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D515C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5218 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D52BC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5360 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D541C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D54D8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D55A0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5644 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D570C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D57D4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5884 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D594C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5A14 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5AC4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5B74 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5C38 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5CE8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5D98 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5E48 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5F1C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D5FF0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D60F0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D61D0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D62D8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D63C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003C6D92 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003CD240 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x003CFD94 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003D0142 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003D0200 | `iPod_Control/Device/Trainer/Workouts/Empeds/` | Filesystem Path |  |
| 0x003D2010 | `Resources/Fonts` | Filesystem Path |  |
| 0x003DA5BC | `iPod_Control:Device:voices:` | Filesystem Path |  |
| 0x003DC57E | `Resources/TrainerTemplates` | Filesystem Path |  |
| 0x003DC599 | `iPod_Control/Device/Trainer/TrainerTemplates` | Filesystem Path |  |
| 0x003DCBEC | `iPod_Control/Device/Trainer/Workouts/Empeds/` | Filesystem Path |  |
| 0x003DCC73 | `/iPod_Control/Device/Trainer/Workouts/Empeds` | Filesystem Path |  |
| 0x009CB6E5 | `Resources/Games/` | Filesystem Path |  |
| 0x009CBAD1 | `iPod_Control/Device` | Filesystem Path |  |
| 0x009CBAE5 | `iPod_Control/Device/accessories.reg` | Filesystem Path |  |
| 0x009CBB66 | `iPod_Control/Tones` | Filesystem Path |  |

---

## 19. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0091CECC | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x00921C14 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftu` | Build Path |  |
| 0x00921C6C | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fts` | Build Path |  |
| 0x00921CC4 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fto` | Build Path |  |
| 0x00925824 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x00925898 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x009264B4 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x00926BE8 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x00927034 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x0092DEF8 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\` | Build Path |  |
| 0x0092EA74 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfd` | Build Path |  |
| 0x0092FC70 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrg` | Build Path |  |
| 0x0092FCC8 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrc` | Build Path |  |
| 0x0092FD20 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfro` | Build Path |  |
| 0x00930064 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1` | Build Path |  |
| 0x0093F40C | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttc` | Build Path |  |
| 0x0093F688 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype` | Build Path |  |
| 0x0093FBF4 | `c:\bwa\N46FirmwareWin-435\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1` | Build Path |  |

---

## 20. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00094E44 | `Acoustic` | EQ Preset |  |
| 0x00094E50 | `Bass Booster` | EQ Preset |  |
| 0x00094E70 | `Classical` | EQ Preset |  |
| 0x00094E7C | `Dance` | EQ Preset |  |
| 0x00094E8C | `Electronic` | EQ Preset |  |
| 0x00094EA0 | `Hip Hop` | EQ Preset |  |
| 0x00094EA8 | `Jazz` | EQ Preset |  |
| 0x00094EB0 | `Latin` | EQ Preset |  |
| 0x00094EB8 | `Loudness` | EQ Preset |  |
| 0x00094EC4 | `Lounge` | EQ Preset |  |
| 0x00094ECC | `Piano` | EQ Preset |  |
| 0x00094EE0 | `Rock` | EQ Preset |  |
| 0x00094EE8 | `Small Speakers` | EQ Preset |  |
| 0x00094EF8 | `Spoken Word` | EQ Preset |  |
| 0x00094F04 | `Treble Booster` | EQ Preset |  |
| 0x00094F50 | `Vocal Booster` | EQ Preset |  |
| 0x00816444 | `Acoustic` | EQ Preset |  |
| 0x00816450 | `Bass Booster` | EQ Preset |  |
| 0x00816470 | `Classical` | EQ Preset |  |
| 0x0081647C | `Dance` | EQ Preset |  |
| 0x0081648C | `Electronic` | EQ Preset |  |
| 0x008164A0 | `Hip Hop` | EQ Preset |  |
| 0x008164A8 | `Jazz` | EQ Preset |  |
| 0x008164B0 | `Latin` | EQ Preset |  |
| 0x008164B8 | `Loudness` | EQ Preset |  |
| 0x008164C4 | `Lounge` | EQ Preset |  |
| 0x008164CC | `Piano` | EQ Preset |  |
| 0x008164DC | `Rock` | EQ Preset |  |
| 0x008164E4 | `Small Speakers` | EQ Preset |  |
| 0x008164F4 | `Spoken Word` | EQ Preset |  |
| 0x00816500 | `Treble Booster` | EQ Preset |  |
| 0x00816520 | `Vocal Booster` | EQ Preset |  |
| 0x0085EC08 | `Acoustic` | EQ Preset |  |
| 0x0085EC14 | `Bass Booster` | EQ Preset |  |
| 0x0085EC34 | `Classical` | EQ Preset |  |
| 0x0085EC40 | `Dance` | EQ Preset |  |
| 0x0085EC50 | `Electronic` | EQ Preset |  |
| 0x0085EC64 | `Hip Hop` | EQ Preset |  |
| 0x0085EC6C | `Jazz` | EQ Preset |  |
| 0x0085EC74 | `Latin` | EQ Preset |  |
| 0x0085EC7C | `Loudness` | EQ Preset |  |
| 0x0085EC88 | `Lounge` | EQ Preset |  |
| 0x0085EC90 | `Piano` | EQ Preset |  |
| 0x0085ECA0 | `Rock` | EQ Preset |  |
| 0x0085ECA8 | `Small Speakers` | EQ Preset |  |
| 0x0085ECB8 | `Spoken Word` | EQ Preset |  |
| 0x0085ECC4 | `Treble Booster` | EQ Preset |  |
| 0x0085ECE4 | `Vocal Booster` | EQ Preset |  |
| 0x00867D48 | `Acoustic` | EQ Preset |  |
| 0x00867D54 | `Bass Booster` | EQ Preset |  |
| 0x00867D74 | `Classical` | EQ Preset |  |
| 0x00867D80 | `Dance` | EQ Preset |  |
| 0x00867D90 | `Electronic` | EQ Preset |  |
| 0x00867DA4 | `Hip Hop` | EQ Preset |  |
| 0x00867DAC | `Jazz` | EQ Preset |  |
| 0x00867DB4 | `Latin` | EQ Preset |  |
| 0x00867DBC | `Loudness` | EQ Preset |  |
| 0x00867DC8 | `Lounge` | EQ Preset |  |
| 0x00867DD0 | `Piano` | EQ Preset |  |
| 0x00867DE0 | `Rock` | EQ Preset |  |
| 0x00867DE8 | `Small Speakers` | EQ Preset |  |
| 0x00867DF8 | `Spoken Word` | EQ Preset |  |
| 0x00867E04 | `Treble Booster` | EQ Preset |  |
| 0x00867E24 | `Vocal Booster` | EQ Preset |  |
| 0x00871368 | `Acoustic` | EQ Preset |  |
| 0x00871398 | `Dance` | EQ Preset |  |
| 0x008713A8 | `Electronic` | EQ Preset |  |
| 0x008713C4 | `Jazz` | EQ Preset |  |
| 0x008713CC | `Latin` | EQ Preset |  |
| 0x008713D4 | `Loudness` | EQ Preset |  |
| 0x008713E8 | `Piano` | EQ Preset |  |
| 0x008713F8 | `Rock` | EQ Preset |  |
| 0x00888B18 | `Dance` | EQ Preset |  |
| 0x00888B40 | `Hip Hop` | EQ Preset |  |
| 0x00888B48 | `Jazz` | EQ Preset |  |
| 0x00888B58 | `Loudness` | EQ Preset |  |
| 0x00888B64 | `Lounge` | EQ Preset |  |
| 0x00888B6C | `Piano` | EQ Preset |  |
| 0x00888B7C | `Rock` | EQ Preset |  |
| 0x00891CEC | `Jazz` | EQ Preset |  |
| 0x00891CF4 | `Latin` | EQ Preset |  |
| 0x00891D08 | `Lounge` | EQ Preset |  |
| 0x00891D10 | `Piano` | EQ Preset |  |
| 0x00891D20 | `Rock` | EQ Preset |  |
| 0x0089ADC4 | `Hip Hop` | EQ Preset |  |
| 0x0089ADCC | `Jazz` | EQ Preset |  |
| 0x0089ADE8 | `Lounge` | EQ Preset |  |
| 0x0089ADF0 | `Piano` | EQ Preset |  |
| 0x0089AE08 | `Rock` | EQ Preset |  |
| 0x008A4AB0 | `Latin` | EQ Preset |  |
| 0x008A4ADC | `Rock` | EQ Preset |  |
| 0x008AE108 | `Dance` | EQ Preset |  |
| 0x008AE12C | `Hip Hop` | EQ Preset |  |
| 0x008AE134 | `Jazz` | EQ Preset |  |
| 0x008AE144 | `Loudness` | EQ Preset |  |
| 0x008AE150 | `Lounge` | EQ Preset |  |
| 0x008AE158 | `Piano` | EQ Preset |  |
| 0x008AE168 | `Rock` | EQ Preset |  |
| 0x008B8B18 | `Acoustic` | EQ Preset |  |
| 0x008B8B24 | `Bass Booster` | EQ Preset |  |
| 0x008B8B44 | `Classical` | EQ Preset |  |
| 0x008B8B50 | `Dance` | EQ Preset |  |
| 0x008B8B60 | `Electronic` | EQ Preset |  |
| 0x008B8B74 | `Hip Hop` | EQ Preset |  |
| 0x008B8B7C | `Jazz` | EQ Preset |  |
| 0x008B8B84 | `Latin` | EQ Preset |  |
| 0x008B8B8C | `Loudness` | EQ Preset |  |
| 0x008B8B98 | `Lounge` | EQ Preset |  |
| 0x008B8BA0 | `Piano` | EQ Preset |  |
| 0x008B8BB0 | `Rock` | EQ Preset |  |
| 0x008B8BB8 | `Small Speakers` | EQ Preset |  |
| 0x008B8BC8 | `Spoken Word` | EQ Preset |  |
| 0x008B8BD4 | `Treble Booster` | EQ Preset |  |
| 0x008B8BF4 | `Vocal Booster` | EQ Preset |  |
| 0x008C33CC | `Acoustic` | EQ Preset |  |
| 0x008C33D8 | `Bass Booster` | EQ Preset |  |
| 0x008C33F8 | `Classical` | EQ Preset |  |
| 0x008C3404 | `Dance` | EQ Preset |  |
| 0x008C3414 | `Electronic` | EQ Preset |  |
| 0x008C3428 | `Hip Hop` | EQ Preset |  |
| 0x008C3430 | `Jazz` | EQ Preset |  |
| 0x008C3438 | `Latin` | EQ Preset |  |
| 0x008C3440 | `Loudness` | EQ Preset |  |
| 0x008C344C | `Lounge` | EQ Preset |  |
| 0x008C3454 | `Piano` | EQ Preset |  |
| 0x008C3464 | `Rock` | EQ Preset |  |
| 0x008C346C | `Small Speakers` | EQ Preset |  |
| 0x008C347C | `Spoken Word` | EQ Preset |  |
| 0x008C3488 | `Treble Booster` | EQ Preset |  |
| 0x008C34A8 | `Vocal Booster` | EQ Preset |  |
| 0x008CCC5C | `Dance` | EQ Preset |  |
| 0x008CCC90 | `Jazz` | EQ Preset |  |
| 0x008CCC98 | `Latin` | EQ Preset |  |
| 0x008CCCA0 | `Loudness` | EQ Preset |  |
| 0x008CCCAC | `Lounge` | EQ Preset |  |
| 0x008CCCB4 | `Piano` | EQ Preset |  |
| 0x008CCCC4 | `Rock` | EQ Preset |  |
| 0x008D5D78 | `Dance` | EQ Preset |  |
| 0x008D5DA4 | `Jazz` | EQ Preset |  |
| 0x008D5DB4 | `Loudness` | EQ Preset |  |
| 0x008D5DC0 | `Lounge` | EQ Preset |  |
| 0x008D5DC8 | `Piano` | EQ Preset |  |
| 0x008D5DD8 | `Rock` | EQ Preset |  |
| 0x008DF134 | `Hip Hop` | EQ Preset |  |
| 0x008DF13C | `Jazz` | EQ Preset |  |
| 0x008DF160 | `Lounge` | EQ Preset |  |
| 0x008DF178 | `Rock` | EQ Preset |  |
| 0x008E88D0 | `Hip Hop` | EQ Preset |  |
| 0x008E88D8 | `Jazz` | EQ Preset |  |
| 0x008E88F4 | `Lounge` | EQ Preset |  |
| 0x008E88FC | `Piano` | EQ Preset |  |
| 0x008E890C | `Rock` | EQ Preset |  |
| 0x008FEBE8 | `Acoustic` | EQ Preset |  |
| 0x008FEBF4 | `Bass Booster` | EQ Preset |  |
| 0x008FEC14 | `Classical` | EQ Preset |  |
| 0x008FEC20 | `Dance` | EQ Preset |  |
| 0x008FEC30 | `Electronic` | EQ Preset |  |
| 0x008FEC44 | `Hip Hop` | EQ Preset |  |
| 0x008FEC4C | `Jazz` | EQ Preset |  |
| 0x008FEC54 | `Latin` | EQ Preset |  |
| 0x008FEC5C | `Loudness` | EQ Preset |  |
| 0x008FEC68 | `Lounge` | EQ Preset |  |
| 0x008FEC70 | `Piano` | EQ Preset |  |
| 0x008FEC80 | `Rock` | EQ Preset |  |
| 0x008FEC88 | `Small Speakers` | EQ Preset |  |
| 0x008FEC98 | `Spoken Word` | EQ Preset |  |
| 0x008FECA4 | `Treble Booster` | EQ Preset |  |
| 0x008FECC4 | `Vocal Booster` | EQ Preset |  |
| 0x00907F94 | `Hip Hop` | EQ Preset |  |
| 0x00907FA0 | `Latin` | EQ Preset |  |
| 0x00907FA8 | `Loudness` | EQ Preset |  |
| 0x00907FB4 | `Lounge` | EQ Preset |  |
| 0x00907FCC | `Rock` | EQ Preset |  |
| 0x00911428 | `Acoustic` | EQ Preset |  |
| 0x00911434 | `Bass Booster` | EQ Preset |  |
| 0x00911454 | `Classical` | EQ Preset |  |
| 0x00911460 | `Dance` | EQ Preset |  |
| 0x00911470 | `Electronic` | EQ Preset |  |
| 0x00911484 | `Hip Hop` | EQ Preset |  |
| 0x0091148C | `Jazz` | EQ Preset |  |
| 0x00911494 | `Latin` | EQ Preset |  |
| 0x0091149C | `Loudness` | EQ Preset |  |
| 0x009114A8 | `Lounge` | EQ Preset |  |
| 0x009114B0 | `Piano` | EQ Preset |  |
| 0x009114C0 | `Rock` | EQ Preset |  |
| 0x009114C8 | `Small Speakers` | EQ Preset |  |
| 0x009114D8 | `Spoken Word` | EQ Preset |  |
| 0x009114E4 | `Treble Booster` | EQ Preset |  |
| 0x00911504 | `Vocal Booster` | EQ Preset |  |
| 0x0091A794 | `Acoustic` | EQ Preset |  |
| 0x0091A7A0 | `Bass Booster` | EQ Preset |  |
| 0x0091A7C0 | `Classical` | EQ Preset |  |
| 0x0091A7CC | `Dance` | EQ Preset |  |
| 0x0091A7DC | `Electronic` | EQ Preset |  |
| 0x0091A7F0 | `Hip Hop` | EQ Preset |  |
| 0x0091A7F8 | `Jazz` | EQ Preset |  |
| 0x0091A800 | `Latin` | EQ Preset |  |
| 0x0091A808 | `Loudness` | EQ Preset |  |
| 0x0091A814 | `Lounge` | EQ Preset |  |
| 0x0091A81C | `Piano` | EQ Preset |  |
| 0x0091A82C | `Rock` | EQ Preset |  |
| 0x0091A834 | `Small Speakers` | EQ Preset |  |
| 0x0091A844 | `Spoken Word` | EQ Preset |  |
| 0x0091A850 | `Treble Booster` | EQ Preset |  |
| 0x0091A870 | `Vocal Booster` | EQ Preset |  |

---
