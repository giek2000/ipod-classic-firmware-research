# iPod Nano 3rd Gen - RetailOS 1.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.1 |
| **IPSW** | iPod_26.1.1.ipsw |
| **Device** | iPod Nano 3rd Gen (2008, 4/8GB NAND, Click Wheel, Cover Flow, Video) |
| **UpdaterFamilyID** | 26 |
| **Binary Size** | 10,736,240 bytes (10.24 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,734,192 bytes |
| **Total Strings (>=4)** | 71,664 |
| **Function Prologues** | 22,511 (ARM: 17,336, Thumb: 5,175) |
| **DRAM References** | 105,440 |
| **Peripheral Refs** | 7,380 |
| **Build** | N46FirmwareWin-359 |
| **SoC** | S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Codename** | N46 |
| **DFU PID** | 0x1229 |
| **SHA-256** | `34976b46ac3e6cd6e2818ac7e86b8f07966e91e970f7f4cf988732e9e1996a98` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000A124C | `TSilverCntlr` | Known | Controller |
| 0x000A1264 | `TCExtrasMenu` | Known | Controller |
| 0x000A127C | `TCGameScreen` | Known | Controller |
| 0x000A1294 | `TCGamesMenu` | Known | Controller |
| 0x000A12A8 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x000A12D0 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x000A12F8 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x000A1324 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x000A1348 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x000A1370 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x000A1398 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x000A13C0 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x000A13E8 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x000A1410 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x000A1440 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x000A146C | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x000A149C | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x000A14C4 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x000A14EC | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x000A1518 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x000A1540 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x000A1568 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x000A1598 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x000A15C8 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x000A16D0 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x000A1700 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x000A1728 | `TCRentalNotification` | Known | Controller |
| 0x000A1748 | `TCRentalInfo` | Known | Controller |
| 0x000A1760 | `TCRentalConfirmDelete` | Known | Controller |
| 0x000A1780 | `TCRentalDispatcher` | Known | Controller |
| 0x000A179C | `TSilverGlobalCntlr` | Known | Controller |
| 0x000A17B8 | `TSilverTrainerCntlr` | Known | Controller |
| 0x000FA740 | `TCSlideshowLCD` | Known | Controller |
| 0x000FA758 | `TCSlideshowTVOut` | Known | Controller |
| 0x000FA774 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x000FA794 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x001236A0 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x001236CC | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x001236F8 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x00123720 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x0012374C | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x00123774 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x001237A0 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x0012A9A8 | `TCRemoteUI` | Known | Controller |
| 0x0012A9BC | `TCUnsupported` | Known | Controller |
| 0x00130D68 | `TCSpeakers` | Known | Controller |
| 0x00130D7C | `TCEQSetting` | Known | Controller |
| 0x0015C634 | `TCSportTimer` | Known | Controller |
| 0x0015C64C | `TCSportTimerMenu` | Known | Controller |
| 0x0015C668 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x0015C68C | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x0015DA0C | `TCVoiceMemos` | Known | Controller |
| 0x0015DA24 | `TCVoiceMemosMenu` | Known | Controller |
| 0x0015DA40 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x0015DA60 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x0015DA80 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x0016EC58 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x0016EC80 | `TCSettings_MainMenu` | Known | Controller |
| 0x0016EC9C | `TCSettings_MusicMenu` | Known | Controller |
| 0x0016ECBC | `TCSettings_VolumeLimit` | Known | Controller |
| 0x0016ECDC | `TCSettings_Brightness` | Known | Controller |
| 0x0016ECFC | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0016ED20 | `TCSettings_EQ` | Known | Controller |
| 0x0016ED38 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x0016ED60 | `TCSettings_RadioRegions` | Known | Controller |
| 0x0016ED80 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0016EDA4 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x0016EDC8 | `TCDateTimeScreen` | Known | Controller |
| 0x0016EDE4 | `TCTimeZoneScreen` | Known | Controller |
| 0x0016EE00 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0016EE28 | `TCFirstBoot` | Known | Controller |
| 0x00184648 | `TCDemoMode` | Known | Controller |
| 0x001B0000 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x001B0020 | `TCAddressViewerDetails` | Known | Controller |
| 0x001B0040 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x001B0064 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x001E00BC | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x001E00E0 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x001E80B4 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x00283350 | `TC_LockDialog` | Known | Controller |
| 0x00283368 | `TC_LockScreen` | Known | Controller |
| 0x00283380 | `TC_LockediPod` | Known | Controller |
| 0x00283398 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x002833BC | `TCLockChosenDispatcher` | Known | Controller |
| 0x00289258 | `TCClock` | Known | Controller |
| 0x00289268 | `TCClockCityMenu` | Known | Controller |
| 0x00289280 | `TCClockRegionMenu` | Known | Controller |
| 0x0028929C | `TCAlarmMenu` | Known | Controller |
| 0x002892B0 | `TCSleepTimerMenu` | Known | Controller |
| 0x002892CC | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x002892EC | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x00289314 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x00289338 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0028935C | `TCAlarmDatePicker` | Known | Controller |
| 0x00289378 | `TCAlarmTriggered` | Known | Controller |
| 0x002901DC | `TCNotesDispatcher` | Known | Controller |
| 0x002901F8 | `TCNotesLoading` | Known | Controller |
| 0x00290210 | `TCNotesList` | Known | Controller |
| 0x00290224 | `TCNotesContents` | Known | Controller |
| 0x003BDFA4 | `TCAlarmTriggered` | Known | Controller |
| 0x003BDFB8 | `TSilverCntlr` | Known | Controller |
| 0x003BDFD8 | `TCClock` | Known | Controller |
| 0x003BDFE0 | `TCClockRegionMenu` | Known | Controller |
| 0x003BDFF4 | `TCClockCityMenu` | Known | Controller |
| 0x003BE004 | `TCAlarmMenu` | Known | Controller |
| 0x003BE010 | `TCSleepTimerMenu` | Known | Controller |
| 0x003BE024 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003BE03C | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003BE05C | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003BE078 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003BE094 | `TCAlarmDatePicker` | Known | Controller |
| 0x003BE0CC | `TSilverCntlr` | Known | Controller |
| 0x003BE0EC | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003BE27C | `TSilverCntlr` | Known | Controller |
| 0x003BE29C | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x003BE2BC | `TCSettings_Brightness` | Known | Controller |
| 0x003BE2D4 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x003BE2F0 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x003BE310 | `TCSettings_RadioRegions` | Known | Controller |
| 0x003BE328 | `TCSettings_EQ` | Known | Controller |
| 0x003BE338 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x003BE354 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x003BE374 | `TCFirstBoot` | Known | Controller |
| 0x003BE380 | `TCSettings_MainMenu` | Known | Controller |
| 0x003BE394 | `TCSettings_MusicMenu` | Known | Controller |
| 0x003BE3AC | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003BE3C4 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x003BE3E0 | `TCDateTimeScreen` | Known | Controller |
| 0x003BE3F4 | `TCTimeZoneScreen` | Known | Controller |
| 0x003C53EC | `TSilverCntlr` | Known | Controller |
| 0x003C540C | `TCClock` | Known | Controller |
| 0x003C5414 | `TCClockRegionMenu` | Known | Controller |
| 0x003C5428 | `TCClockCityMenu` | Known | Controller |
| 0x003C5438 | `TCAlarmMenu` | Known | Controller |
| 0x003C5444 | `TCSleepTimerMenu` | Known | Controller |
| 0x003C5458 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003C54D0 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003C54F0 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003C550C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003C5540 | `TCAlarmDatePicker` | Known | Controller |
| 0x003C5554 | `TCAlarmTriggered` | Known | Controller |
| 0x003C6FD0 | `TSilverCntlr` | Known | Controller |
| 0x003C6FF0 | `TC_LockDialog` | Known | Controller |
| 0x003C7000 | `TC_LockScreen` | Known | Controller |
| 0x003C7010 | `TC_LockediPod` | Known | Controller |
| 0x003C7020 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003C703C | `TCLockChosenDispatcher` | Known | Controller |
| 0x003C7054 | `TSilverCntlr` | Known | Controller |
| 0x003C71BC | `TSilverCntlr` | Known | Controller |
| 0x003C71D8 | `TSilverCntlr` | Known | Controller |
| 0x003C723C | `TSilverCntlr` | Known | Controller |
| 0x003C725C | `TCNotesDispatcher` | Known | Controller |
| 0x003C7270 | `TCNotesLoading` | Known | Controller |
| 0x003C7280 | `TCNotesBase` | Known | Controller |
| 0x003C728C | `TCNotesList` | Known | Controller |
| 0x003C7298 | `TCNotesContents` | Known | Controller |
| 0x003C72A8 | `TSilverCntlr` | Known | Controller |
| 0x003C72C8 | `TCRemoteUI` | Known | Controller |
| 0x003C72D4 | `TCUnsupported` | Known | Controller |
| 0x003C72E4 | `TSilverCntlr` | Known | Controller |
| 0x003C7348 | `TSilverCntlr` | Known | Controller |
| 0x003C7368 | `TCSportTimer` | Known | Controller |
| 0x003C7378 | `TCSportTimerMenu` | Known | Controller |
| 0x003C738C | `TCSportTimerSessionScreen` | Known | Controller |
| 0x003C73A8 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x003C7500 | `TSilverCntlr` | Known | Controller |
| 0x003C7520 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003C753C | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003C755C | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003C757C | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003C75A4 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003C75C8 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003C75F0 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003C7610 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003C7630 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003C7650 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003C7670 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003C7698 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003C79F4 | `TSilverCntlr` | Known | Controller |
| 0x003C7B1C | `TSilverCntlr` | Known | Controller |
| 0x003C7B3C | `TCDemoMode` | Known | Controller |
| 0x003C7B48 | `TCClock` | Known | Controller |
| 0x003C7B50 | `TCClockRegionMenu` | Known | Controller |
| 0x003C7B64 | `TCClockCityMenu` | Known | Controller |
| 0x003C7B74 | `TCAlarmMenu` | Known | Controller |
| 0x003C7B80 | `TCSleepTimerMenu` | Known | Controller |
| 0x003C7B94 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003C7BAC | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003C7BCC | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003C7BE8 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003C7C04 | `TCAlarmDatePicker` | Known | Controller |
| 0x003C7C18 | `TCAlarmTriggered` | Known | Controller |
| 0x003C7C38 | `TSilverCntlr` | Known | Controller |
| 0x003C7C54 | `TSilverCntlr` | Known | Controller |
| 0x003C7C64 | `TSilverCntlr` | Known | Controller |
| 0x003C7C84 | `TCVoiceMemos` | Known | Controller |
| 0x003C7C94 | `TCVoiceMemosMenu` | Known | Controller |
| 0x003C7CA8 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x003C7CC0 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x003C7CD8 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x003C7CF8 | `TSilverCntlr` | Known | Controller |
| 0x003C7D58 | `TSilverCntlr` | Known | Controller |
| 0x003C7DC4 | `TSilverCntlr` | Known | Controller |
| 0x003C8DB4 | `TSilverCntlr` | Known | Controller |
| 0x003C8EC0 | `TSilverCntlr` | Known | Controller |
| 0x003D1654 | `TSilverCntlr` | Known | Controller |
| 0x003D1674 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x003D168C | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x003D16A8 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x003D16C8 | `TCAddressViewerDetails` | Known | Controller |
| 0x003D16E0 | `TSilverCntlr` | Known | Controller |
| 0x003D1700 | `TSilverCalendarCntlr_Base` | Known | Controller |
| 0x003D171C | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x003D1740 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x003D1764 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x003D1784 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x003D17A8 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x003D17C8 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x003D17EC | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x003D19C4 | `TSilverCntlr` | Known | Controller |
| 0x003D19E4 | `TC_LockDialog` | Known | Controller |
| 0x003D19F4 | `TC_LockScreen` | Known | Controller |
| 0x003D1A04 | `TC_LockediPod` | Known | Controller |
| 0x003D1A14 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003D1A38 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003D1AEC | `TSilverCntlr` | Known | Controller |
| 0x003D1B0C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003D1B28 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003D1B48 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003D1B68 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003D1B90 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003D1BB4 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003D1BDC | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003D1BFC | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003D1C1C | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003D1C3C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003D1C5C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003D1C84 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003D1CAC | `TSilverCntlr` | Known | Controller |
| 0x003D1DCC | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003D1DE8 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003D1E08 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003D1E28 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003D1E50 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003D1E74 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003D1E9C | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003D1EBC | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003D1EDC | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003D1EFC | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003D1F1C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003D1F44 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003D1F6C | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003D1F8C | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003D1FAC | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003D1FD0 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003D1FF0 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003D2014 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003D203C | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003D2068 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003D2088 | `TCRentalNotification` | Known | Controller |
| 0x003D20A0 | `TCRentalInfo` | Known | Controller |
| 0x003D20B0 | `TCRentalConfirmDelete` | Known | Controller |
| 0x003D20C8 | `TCRentalDispatcher` | Known | Controller |
| 0x003D29B8 | `TSilverCntlr` | Known | Controller |
| 0x003D2A7C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003D2A98 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003D2AB8 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003D2AD8 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003D2B00 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003D2B24 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003D2B4C | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003D2B6C | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003D2B8C | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003D2BAC | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003D2BCC | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003D2BF4 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003D2C44 | `TCSlideshowTVOut` | Known | Controller |
| 0x003D2C58 | `TCSlideshowLCD` | Known | Controller |
| 0x003D2C68 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x003D2C80 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x003D2CA0 | `TSilverCntlr` | Known | Controller |
| 0x003D2CCC | `TSilverCntlr` | Known | Controller |
| 0x003D2CEC | `TCUnsupported` | Known | Controller |
| 0x003D2D0C | `TSilverCntlr` | Known | Controller |
| 0x003D2D4C | `TSilverCntlr` | Known | Controller |
| 0x003D2D6C | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x003D2D88 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x003D2DA0 | `TSilverCntlr` | Known | Controller |
| 0x003D2DC0 | `TCSpeakers` | Known | Controller |
| 0x003D2DCC | `TCEQSetting` | Known | Controller |
| 0x003D2E74 | `TSilverCntlr` | Known | Controller |
| 0x003D2E84 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003D2EA0 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003D2EC0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003D2EE0 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003D2F08 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003D2F2C | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003D2F54 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003D2F74 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003D2F94 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003D2FB4 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003D2FD4 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003D2FFC | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003D35A4 | `TSilverCntlr` | Known | Controller |
| 0x003D35C8 | `TSilverCntlr` | Known | Controller |
| 0x003D3634 | `TSilverCntlr` | Known | Controller |
| 0x003D3654 | `TCExtrasMenu` | Known | Controller |
| 0x003D3664 | `TCGamesMenu` | Known | Controller |
| 0x003D3670 | `TCGameScreen` | Known | Controller |
| 0x003D3680 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x003D36A0 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x003D36C0 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x003D36E0 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x003D3704 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003D3720 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003D3740 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003D3760 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003D3788 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003D37AC | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003D37D4 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003D37F4 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003D3814 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003D3834 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003D3854 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003D387C | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003D38A4 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003D38C4 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003D38E4 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003D3908 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003D3928 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003D394C | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003D3974 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003D39A0 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003D39C0 | `TCRentalNotification` | Known | Controller |
| 0x003D39D8 | `TCRentalInfo` | Known | Controller |
| 0x003D39E8 | `TCRentalConfirmDelete` | Known | Controller |
| 0x003D3A00 | `TCRentalDispatcher` | Known | Controller |
| 0x003D3A14 | `TSilverGlobalCntlr` | Known | Controller |
| 0x003D3A28 | `TSilverTrainerCntlr` | Known | Controller |
| 0x0045E3C8 | `TCCCCCCCCCCCCCCCCCCCCCCT` | Known | Controller |
| 0x0074A8EE | `TCNotesDispatcher"` | Known | Controller |
| 0x0074A9AD | `TCLockChosenDispatcher"` | Known | Controller |
| 0x0074AA70 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x00754AD5 | `TCNotesDispatcher"` | Known | Controller |
| 0x00754C37 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x0076C298 | `TCExtrasMenuTCAddressViewerMainMenu` | Known | Controller |
| 0x0076C2BC | `TCAddressViewerDetails` | Known | Controller |
| 0x0076C2D4 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x0076C2F0 | `TCAlarmMenu` | Known | Controller |
| 0x0076C2FC | `TCSleepTimerMenuTCAlarmPropertiesMenu` | Known | Controller |
| 0x0076C324 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x0076C344 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x0076C360 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0076C37C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0076C398 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0076C3B4 | `TCAlarmDatePicker` | Known | Controller |
| 0x0076C3C8 | `TCAlarmDatePicker` | Known | Controller |
| 0x0076C3DC | `TCAlarmTriggeredTSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x0076C408 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x0076C42C | `TSilverCalendarCntlr_MonthViewerTSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x0076C46C | `TSilverCalendarCntlr_EventViewerTSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x0076C4AC | `TSilverCalendarCntlr_EventViewerTSilverCntlrTCClockRegionMenu` | Known | Controller |
| 0x0076C4EC | `TCClockCityMenu` | Known | Controller |
| 0x0076C4FC | `TCClockCityMenu` | Known | Controller |
| 0x0076C50C | `TCClockCityMenu` | Known | Controller |
| 0x0076C51C | `TCClockCityMenu` | Known | Controller |
| 0x0076C52C | `TCClockCityMenu` | Known | Controller |
| 0x0076C53C | `TCClockCityMenu` | Known | Controller |
| 0x0076C54C | `TCClockCityMenu` | Known | Controller |
| 0x0076C55C | `TCClockCityMenu` | Known | Controller |
| 0x0076C56C | `TCClock` | Known | Controller |
| 0x0076C584 | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDe` | Known | Controller |
| 0x0076C5DC | `TCGamesMenu` | Known | Controller |
| 0x0076C5E8 | `TCGameScreenTC_LockediPod` | Known | Controller |
| 0x0076C604 | `TC_LockDialog` | Known | Controller |
| 0x0076C614 | `TC_LockScreen` | Known | Controller |
| 0x0076C624 | `TC_VolumeLimitLockScreenTSilverCntlrTSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x0076C668 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x0076C688 | `TSilverMediaListCntlr_AudiobooksTSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x0076C6D0 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x0076C6EC | `TSilverMediaListCntlr_AlbumsTSilverMediaListCntlr_Artists` | Known | Controller |
| 0x0076C728 | `TSilverMediaListCntlr_GenresTSilverMediaListCntlr_Composers` | Known | Controller |
| 0x0076C764 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0076C784 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0076C7AC | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x0076C7CC | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x0076C7EC | `TSilverMediaListCntlr_TVEpisodesTSilverMediaListCntlr_MoviesTSilverMediaListCntl` | Known | Controller |
| 0x0076C848 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0076C870 | `TSilverMediaListCntlr_VideoPlaylistsTSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x0076C8B4 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x0076C8E0 | `TSilverCntlrTSilverMainMediaListCntlr_VideosTSilverSettingsVideoCntlr` | Known | Controller |
| 0x0076C928 | `TCFirstBoot` | Known | Controller |
| 0x0076C9D0 | `TCRentalInfoTCRentalConfirmDelete` | Known | Controller |
| 0x0076C9F4 | `TSilverCntlrTCRentalNotificationTCRentalNotificationTCRentalNotificationTCNotesL` | Known | Controller |
| 0x0076CA4C | `TCNotesList` | Known | Controller |
| 0x0076CA58 | `TCNotesList` | Known | Controller |
| 0x0076CA64 | `TCNotesContents` | Known | Controller |
| 0x0076CA74 | `TCNotesContents` | Known | Controller |
| 0x0076CA84 | `TCNotesContents` | Known | Controller |
| 0x0076CA94 | `TCNotesContents` | Known | Controller |
| 0x0076CB50 | `TCSlideshowLCD` | Known | Controller |
| 0x0076CB60 | `TCSlideshowTVOutTCSlideshow_TVOutAskTCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x0076CBB0 | `TCRemoteUI` | Known | Controller |
| 0x0076CBBC | `TCUnsupported` | Known | Controller |
| 0x0076CBCC | `TSilverSettingsMenuListCntlrTSilverSettingsMenuListCntlrTSilverSettingsMenuListC` | Known | Controller |
| 0x0076CC34 | `TCSettings_MusicMenuTCSettings_VolumeLimit` | Known | Controller |
| 0x0076CC60 | `TCSettings_Brightness` | Known | Controller |
| 0x0076CC78 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0076CC94 | `TCSettings_AudiobookSettingsTCSettings_RadioRegions` | Known | Controller |
| 0x0076CCC8 | `TCSettings_EQ` | Known | Controller |
| 0x0076CCD8 | `TSilverSettingsMenuListCntlrTSilverCntlrTCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0076CD20 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0076CD3C | `TCSettings_MainMenu` | Known | Controller |
| 0x0076CD50 | `TCSettings_MusicMenuTCSportTimerTCSportTimerMenuTCSportTimerSessionScreen` | Known | Controller |
| 0x0076CEB4 | `TSilverCntlrTTrainerEndSessionCntlr` | Known | Controller |
| 0x0076CF2C | `TSilverCntlrTTrainerCalibrateWalkMenuCntlr` | Known | Controller |
| 0x0076D1C0 | `TCVoiceMemosTCVoiceMemosMainMenuTCVoiceMemosMainMenuTCVoiceMemosMainMenuTSearchC` | Known | Controller |
| 0x0076D220 | `TCEQSetting` | Known | Controller |
| 0x0076D2CE | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x0076E5D1 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x007741DA | `TCLockChosenDispatcher` | Known | Controller |
| 0x00774238 | `TCNotesDispatcher` | Known | Controller |
| 0x00775E16 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00775E74 | `TCNotesDispatcher` | Known | Controller |
| 0x00777A52 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00777AB0 | `TCNotesDispatcher` | Known | Controller |
| 0x0077968E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007796EC | `TCNotesDispatcher` | Known | Controller |
| 0x0077B2CA | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077B328 | `TCNotesDispatcher` | Known | Controller |
| 0x0077CF06 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077CF64 | `TCNotesDispatcher` | Known | Controller |
| 0x0077EB42 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077EBA0 | `TCNotesDispatcher` | Known | Controller |
| 0x0078077E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007807DC | `TCNotesDispatcher` | Known | Controller |
| 0x007823BA | `TCLockChosenDispatcher` | Known | Controller |
| 0x00782418 | `TCNotesDispatcher` | Known | Controller |
| 0x00783FF6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00784054 | `TCNotesDispatcher` | Known | Controller |
| 0x00785C32 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00785C90 | `TCNotesDispatcher` | Known | Controller |
| 0x0078786E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007878CC | `TCNotesDispatcher` | Known | Controller |
| 0x007894AA | `TCLockChosenDispatcher` | Known | Controller |
| 0x00789508 | `TCNotesDispatcher` | Known | Controller |
| 0x0078B0E6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078B144 | `TCNotesDispatcher` | Known | Controller |
| 0x0078CD22 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078CD80 | `TCNotesDispatcher` | Known | Controller |
| 0x0078E95E | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078E9BC | `TCNotesDispatcher` | Known | Controller |
| 0x0079059A | `TCLockChosenDispatcher` | Known | Controller |
| 0x007905F8 | `TCNotesDispatcher` | Known | Controller |
| 0x007921D6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00792234 | `TCNotesDispatcher` | Known | Controller |
| 0x00793E12 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00793E70 | `TCNotesDispatcher` | Known | Controller |
| 0x00795A4E | `TCLockChosenDispatcher` | Known | Controller |
| 0x00795AAC | `TCNotesDispatcher` | Known | Controller |
| 0x0079768A | `TCLockChosenDispatcher` | Known | Controller |
| 0x007976E8 | `TCNotesDispatcher` | Known | Controller |
| 0x007992C6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00799324 | `TCNotesDispatcher` | Known | Controller |
| 0x0079AF02 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0079AF60 | `TCNotesDispatcher` | Known | Controller |
| 0x0079CB3E | `TCLockChosenDispatcher` | Known | Controller |
| 0x0079CB9C | `TCNotesDispatcher` | Known | Controller |
| 0x0079E77A | `TCLockChosenDispatcher` | Known | Controller |
| 0x0079E7D8 | `TCNotesDispatcher` | Known | Controller |
| 0x007A03B6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A0414 | `TCNotesDispatcher` | Known | Controller |
| 0x007A1FF2 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A2050 | `TCNotesDispatcher` | Known | Controller |
| 0x007A3C2E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A3C8C | `TCNotesDispatcher` | Known | Controller |
| 0x007A586A | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A58C8 | `TCNotesDispatcher` | Known | Controller |
| 0x007A74A6 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A7504 | `TCNotesDispatcher` | Known | Controller |
| 0x007A90E2 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007A9140 | `TCNotesDispatcher` | Known | Controller |
| 0x007AAD1E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007AAD7C | `TCNotesDispatcher` | Known | Controller |
| 0x007AC95A | `TCLockChosenDispatcher` | Known | Controller |
| 0x007AC9B8 | `TCNotesDispatcher` | Known | Controller |
| 0x007AE596 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007AE5F4 | `TCNotesDispatcher` | Known | Controller |
| 0x007B01D2 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B0230 | `TCNotesDispatcher` | Known | Controller |
| 0x007B1E0E | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B1E6C | `TCNotesDispatcher` | Known | Controller |
| 0x007B3A4A | `TCLockChosenDispatcher` | Known | Controller |
| 0x007B3AA8 | `TCNotesDispatcher` | Known | Controller |
| 0x007BF680 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x007BF942 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x007C0178 | `TCRentalDispatcher` | Known | Controller |
| 0x007C0A30 | `TCRentalDispatcher` | Known | Controller |
| 0x007C12E8 | `TCRentalDispatcher` | Known | Controller |
| 0x007C1BA0 | `TCRentalDispatcher` | Known | Controller |
| 0x007C2458 | `TCRentalDispatcher` | Known | Controller |
| 0x007C2D10 | `TCRentalDispatcher` | Known | Controller |
| 0x007C35C8 | `TCRentalDispatcher` | Known | Controller |
| 0x007C3E80 | `TCRentalDispatcher` | Known | Controller |
| 0x009155FC | `TCMockupModeNavScreen` | Known | Controller |
| 0x00915614 | `TSilverCntlr` | Known | Controller |
| 0x00915634 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x0091566C | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x0091568C | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x009156AC | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x009156D0 | `TCExtrasMenu` | Known | Controller |
| 0x009157E0 | `TSilverCntlr` | Known | Controller |
| 0x00915800 | `TCSlideshowTVOut` | Known | Controller |
| 0x00915814 | `TCSlideshowLCD` | Known | Controller |
| 0x00915824 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x0091583C | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00915878 | `TSilverCntlr` | Known | Controller |
| 0x009158F4 | `TCSlideshowTVOut` | Known | Controller |
| 0x00915908 | `TCSlideshowLCD` | Known | Controller |
| 0x00915918 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x00915930 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00915950 | `TSilverCntlr` | Known | Controller |
| 0x00916574 | `TSilverCntlr` | Known | Controller |
| 0x00916594 | `TCGamesMenu` | Known | Controller |
| 0x009165A0 | `TCGameScreen` | Known | Controller |
| 0x009D4469 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |

---

## 2. Settings (Toggle/Show/TC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00139F50 | `ShowSetting_EQ` | Known | User setting |
| 0x001E9868 | `ToggleSetting_Repeat` | Known | User setting |
| 0x001E9884 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001E989C | `ToggleSetting_TVOut` | Known | User setting |
| 0x001E98B0 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x002192A4 | `ShowSetting_Backlight` | Known | User setting |
| 0x0022DD70 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x0022DD8C | `ToggleSetting_Repeat` | Known | User setting |
| 0x0022DDA4 | `ToggleSetting_SortBy` | Known | User setting |
| 0x0022DDBC | `ToggleSetting_ClassicUI` | Known | User setting |
| 0x0022DDD4 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x0022DDF0 | `ToggleSetting_Clicker` | Known | User setting |
| 0x0022DE08 | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x0022DE28 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x0022DE44 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x0022DE60 | `ShowSetting_Shuffle` | Known | User setting |
| 0x0022E00C | `ShowSetting_Repeat` | Known | User setting |
| 0x0022E020 | `ShowSetting_About` | Known | User setting |
| 0x0022E034 | `ShowSetting_MainMenu` | Known | User setting |
| 0x0022E04C | `ShowSetting_MusicMenu` | Known | User setting |
| 0x0022E064 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x0022E07C | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x0022E098 | `ShowSetting_Brightness` | Known | User setting |
| 0x0022E0B0 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x0022E0C8 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x0022E0E4 | `ShowSetting_EQ` | Known | User setting |
| 0x0022E0F4 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x0022E290 | `ShowSetting_Clicker` | Known | User setting |
| 0x0022E2A4 | `ShowSetting_DateAndTime` | Known | User setting |
| 0x0022E2BC | `ShowSetting_SortBy` | Known | User setting |
| 0x0022E2D0 | `ShowSetting_ClassicUI` | Known | User setting |
| 0x0022E2E8 | `ShowSetting_Language` | Known | User setting |
| 0x0022E300 | `ShowSetting_Legal` | Known | User setting |
| 0x0022E314 | `ShowSetting_ResetAll` | Known | User setting |
| 0x0075395D | `ToggleSetting_24HourClock` | Known | User setting |
| 0x00753A0D | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x007560A2 | `ShowSetting_About` | Known | User setting |
| 0x00756144 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x00756188 | `ShowSetting_Shuffle` | Known | User setting |
| 0x007561FF | `ToggleSetting_Repeat` | Known | User setting |
| 0x00756242 | `ShowSetting_Repeat` | Known | User setting |
| 0x0075634C | `ShowSetting_MainMenu` | Known | User setting |
| 0x0075645C | `ShowSetting_MusicMenu` | Known | User setting |
| 0x00756524 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x007565EE | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x00756706 | `ShowSetting_Brightness` | Known | User setting |
| 0x0075683C | `ShowSetting_Audiobooks` | Known | User setting |
| 0x0075694D | `ShowSetting_RadioRegions` | Known | User setting |
| 0x00756A4E | `ShowSetting_EQ` | Known | User setting |
| 0x00756ABB | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x00756B02 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x00756B7F | `ToggleSetting_Clicker` | Known | User setting |
| 0x00756BC3 | `ShowSetting_Clicker` | Known | User setting |
| 0x00756D2A | `ToggleSetting_SortBy` | Known | User setting |
| 0x00756D6D | `ShowSetting_SortBy` | Known | User setting |
| 0x00756E6E | `ShowSetting_Language` | Known | User setting |
| 0x00756F7E | `ShowSetting_Legal` | Known | User setting |
| 0x007570AF | `ShowSetting_ResetAll` | Known | User setting |
| 0x00757220 | `ShowSetting_Backlight` | Known | User setting |
| 0x007572D0 | `ShowSetting_Backlight` | Known | User setting |
| 0x00757380 | `ShowSetting_Backlight` | Known | User setting |
| 0x00757431 | `ShowSetting_Backlight` | Known | User setting |
| 0x007574E2 | `ShowSetting_Backlight` | Known | User setting |
| 0x00757593 | `ShowSetting_Backlight` | Known | User setting |
| 0x00757647 | `ShowSetting_Backlight` | Known | User setting |
| 0x007576F6 | `ShowSetting_EQ` | Known | User setting |
| 0x0075776B | `ShowSetting_Language` | Known | User setting |
| 0x007D4F94 | `ToggleSetting_Repeat` | Known | User setting |
| 0x007D4FCE | `ToggleSetting_Shuffle` | Known | User setting |
| 0x007D5090 | `ToggleSetting_TVOut` | Known | User setting |
| 0x007D50C9 | `ToggleSetting_TVSignal` | Known | User setting |

---

## 3. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001584AC | `MockupMode/MockupMode.xml` | Hidden | Developer Tool |
| 0x001589AC | `MockupMode/` | Hidden | Developer Tool |
| 0x002686FC | `Channel UnitTests` | Hidden | Developer Tool |
| 0x002BE3F9 | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x002BE43C | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x002BE451 | `RTXCbug> ` | Hidden | Developer Tool |
| 0x002BEE2D | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x002D04E0 | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x00366975 | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x00366A3D | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x003C33A5 | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x0076D12C | `TTrainerLoadingCntlrTSilverCntlrTUnitTestSuiteCntlr` | Hidden | Developer Tool |
| 0x0076D160 | `TUnitTestSuiteTestsCntlrTSilverCntlrTCVoiceMemosMainMenuTCVoiceMemosMenuTCVoiceM` | Hidden | Developer Tool |
| 0x0080BAA0 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00853DC4 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00866290 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0087D830 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0088F940 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00899498 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008A2D0C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008B7B80 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008C1640 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008E7830 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00905AC0 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0090ED18 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x009C4DE2 | `Debug_ListItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x009C4DFA | `Debug_MenuItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x009C5A0F | `DemoMode_Background_Image` | Hidden | Demo/Retail Mode |
| 0x009C6A22 | `DemoMode_Icon_Image` | Hidden | Demo/Retail Mode |
| 0x009C889C | `Debug_Menu_Title` | Hidden | Debug/Diagnostic |
| 0x009C88C1 | `UnitTest_Menu_Title` | Hidden | Developer Tool |
| 0x009D1936 | `UnitTestModel` | Hidden | Developer Tool |
| 0x009D256E | `UnitTest_ListItem` | Hidden | Developer Tool |
| 0x009D3B00 | `DemoMode_View_Main` | Hidden | Demo/Retail Mode |
| 0x009D3CE8 | `DemoMode_Icon` | Hidden | Demo/Retail Mode |
| 0x009D4B2E | `UnitTestApp` | Hidden | Developer Tool |
| 0x009D5190 | `Debug_ListItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009D51AB | `Debug_MenuItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009D59B6 | `DemoMode_Text_Color` | Hidden | Demo/Retail Mode |
| 0x009D5DBF | `Debug_ListItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009D5DD6 | `Debug_MenuItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009DA87B | `Debug_ListItem_UnitTest` | Hidden | Developer Tool |
| 0x009DA893 | `Debug_MenuItem_UnitTest` | Hidden | Developer Tool |
| 0x009DF78F | `Debug_ListItem_Memory` | Hidden | Debug/Diagnostic |
| 0x009DF7A5 | `Debug_MenuItem_Memory` | Hidden | Debug/Diagnostic |

---

## 4. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000677B | `"MeCCADecode` | Known | Audio system |
| 0x0014EAC0 | `AudioCodecs` | Known | Audio system |
| 0x00192788 | `MeCCA_PCM_Output.wav` | Known | Audio system |
| 0x001AF23C | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x001B9AF0 | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x001B9CF8 | `MeCCAVideoDecode` | Known | Audio system |
| 0x009238F8 | `MeCCA_StreamCache` | Known | Audio system |

---

## 5. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F653C | `HandleWheel` | Known | Event handler |
| 0x000F6548 | `HandlePlayPause` | Known | Event handler |
| 0x000F6558 | `HandleSelectDown` | Known | Event handler |
| 0x000F656C | `HandleNext` | Known | Event handler |
| 0x000F6578 | `HandlePrevious` | Known | Event handler |
| 0x000F6588 | `HandleNextPushAndHold` | Known | Event handler |
| 0x000F65A0 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000F6838 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000F6858 | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x00102528 | `HandleSelect` | Known | Event handler |
| 0x0010253C | `HandleHilite` | Known | Event handler |
| 0x001028D4 | `HandleEQSettingSelected` | Known | Event handler |
| 0x00102D04 | `HandleSelect` | Known | Event handler |
| 0x00102D18 | `HandleGameHilited` | Known | Event handler |
| 0x00102FC8 | `HandleNotesSelected` | Known | Event handler |
| 0x00102FE0 | `HandleNotesPop` | Known | Event handler |
| 0x00102FF0 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00111044 | `HandleVolumeWheel` | Known | Event handler |
| 0x00111058 | `HandleVolumeChange` | Known | Event handler |
| 0x0011106C | `HandleTimerDone` | Known | Event handler |
| 0x0011107C | `HandleFrequencyChange` | Known | Event handler |
| 0x001110F4 | `HandleTuning` | Known | Event handler |
| 0x00111104 | `HandleTuningSelect` | Known | Event handler |
| 0x00121608 | `HandleLock` | Known | Event handler |
| 0x00121618 | `HandleAddressBook` | Known | Event handler |
| 0x00121D00 | `HandleSelect` | Known | Event handler |
| 0x00122238 | `HandleExit` | Known | Event handler |
| 0x00122248 | `HandleLap` | Known | Event handler |
| 0x00122254 | `HandleResume` | Known | Event handler |
| 0x00122264 | `HandleStartStop` | Known | Event handler |
| 0x001224EC | `HandleWheel` | Known | Event handler |
| 0x001224FC | `HandlePlayPause` | Known | Event handler |
| 0x0012250C | `HandleSelectDown` | Known | Event handler |
| 0x00122520 | `HandleHilite` | Known | Event handler |
| 0x0012C160 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x0013A184 | `HandleExitUnsupported` | Known | Event handler |
| 0x00145798 | `HandleBasicSelected` | Known | Event handler |
| 0x001457B0 | `HandleTimedWorkoutSelected` | Known | Event handler |
| 0x001457CC | `HandleDistanceWorkoutSelected` | Known | Event handler |
| 0x001457EC | `HandleCaloriesWorkoutSelected` | Known | Event handler |
| 0x0014580C | `HandleSelectWorkout` | Known | Event handler |
| 0x00154218 | `HandleNotesPop` | Known | Event handler |
| 0x0015422C | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00155110 | `HandleSelect` | Known | Event handler |
| 0x00155124 | `HandleWheel` | Known | Event handler |
| 0x00155130 | `HandleImageNext` | Known | Event handler |
| 0x00155140 | `HandleImagePrev` | Known | Event handler |
| 0x00155150 | `HandleImageLast` | Known | Event handler |
| 0x00155160 | `HandleImageFirst` | Known | Event handler |
| 0x00155174 | `HandlePlayPause` | Known | Event handler |
| 0x00155184 | `HandlePlay` | Known | Event handler |
| 0x00155190 | `HandlePause` | Known | Event handler |
| 0x0016975C | `HandleSelectCity` | Known | Event handler |
| 0x00169774 | `HandleHighlightCity` | Known | Event handler |
| 0x0016A69C | `HandleWantPopFlow` | Known | Event handler |
| 0x0016A6B4 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x0016A6D0 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0016A6EC | `HandleFlowNext` | Known | Event handler |
| 0x0016A6FC | `HandleFlowPrev` | Known | Event handler |
| 0x0016A70C | `HandleFlowWheel` | Known | Event handler |
| 0x0016A71C | `HandleAlbumSelected` | Known | Event handler |
| 0x0016A730 | `HandlePlayPause` | Known | Event handler |
| 0x0016A740 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00194624 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00194A14 | `HandleSelect` | Known | Event handler |
| 0x001958D4 | `HandleSelect` | Known | Event handler |
| 0x001958E8 | `HandleWheel` | Known | Event handler |
| 0x001958F4 | `HandleImageNext` | Known | Event handler |
| 0x00195904 | `HandleImagePrev` | Known | Event handler |
| 0x00195914 | `HandleImageLast` | Known | Event handler |
| 0x00195924 | `HandleImageFirst` | Known | Event handler |
| 0x00195938 | `HandlePlayPause` | Known | Event handler |
| 0x00195948 | `HandlePlay` | Known | Event handler |
| 0x00195954 | `HandlePause` | Known | Event handler |
| 0x00195DF4 | `HandleNew` | Known | Event handler |
| 0x00195E04 | `HandleClear` | Known | Event handler |
| 0x00195E10 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x00195E2C | `HandleSelectIndexedSession` | Known | Event handler |
| 0x0019613C | `HandleWheel` | Known | Event handler |
| 0x0019614C | `HandleArrowUp` | Known | Event handler |
| 0x0019615C | `HandleArrowDown` | Known | Event handler |
| 0x00198380 | `HandleHiliteAlbum` | Known | Event handler |
| 0x00198398 | `HandleBrowseAlbum` | Known | Event handler |
| 0x001983AC | `HandlePlayPause` | Known | Event handler |
| 0x001B38D0 | `HandleSelect` | Known | Event handler |
| 0x001B3A60 | `HandleSelectRegion` | Known | Event handler |
| 0x001B8494 | `HandleChooseLink` | Known | Event handler |
| 0x001B84AC | `HandleChooseCalibrate` | Known | Event handler |
| 0x001B84C4 | `HandleUnlink` | Known | Event handler |
| 0x001C942C | `HandleImageWheel` | Known | Event handler |
| 0x001C9444 | `HandlePlayPause` | Known | Event handler |
| 0x001C9454 | `HandleBrowseLarge` | Known | Event handler |
| 0x001C9468 | `HandleBrowseSmall` | Known | Event handler |
| 0x001C947C | `HandleBrowseSlideshow` | Known | Event handler |
| 0x001C9494 | `HandleImageNext` | Known | Event handler |
| 0x001C94A4 | `HandleImagePrev` | Known | Event handler |
| 0x001C94B4 | `HandleHilite` | Known | Event handler |
| 0x001C94C4 | `HandleImageLast` | Known | Event handler |
| 0x001C94D4 | `HandleImageFirst` | Known | Event handler |
| 0x001C94E8 | `HandleScreenNext` | Known | Event handler |
| 0x001C94FC | `HandleScreenPrev` | Known | Event handler |
| 0x001CBDE0 | `HandlePlayPause` | Known | Event handler |
| 0x001CBDF4 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x001CBE10 | `HandleNext` | Known | Event handler |
| 0x001CBE1C | `HandleNextPressAndHold` | Known | Event handler |
| 0x001CBE34 | `HandlePrevious` | Known | Event handler |
| 0x001CBE44 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x001CBE60 | `HandleRemotePlayPause` | Known | Event handler |
| 0x001CBE78 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x001CBE9C | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x001CBEB4 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x001CBECC | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x001CC09C | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x001CC0B4 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x001CC0CC | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x001CC0E8 | `HandleRemoteStop` | Known | Event handler |
| 0x001CC0FC | `HandleRemotePlay` | Known | Event handler |
| 0x001CC110 | `HandleRemotePause` | Known | Event handler |
| 0x001CC124 | `HandleRemoteMute` | Known | Event handler |
| 0x001CC138 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x001CC150 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x001CC168 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x001CC184 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x001CC3A8 | `HandleRemoteShuffle` | Known | Event handler |
| 0x001CC3BC | `HandleRemoteRepeat` | Known | Event handler |
| 0x001CC3D0 | `HandleRemoteOn` | Known | Event handler |
| 0x001CC3E0 | `HandleRemoteOff` | Known | Event handler |
| 0x001CC3F0 | `HandleRemoteBacklight` | Known | Event handler |
| 0x001CC408 | `HandleRemoteFFDown` | Known | Event handler |
| 0x001CC41C | `HandleRemoteFFUp` | Known | Event handler |
| 0x001CC430 | `HandleRemoteRewDown` | Known | Event handler |
| 0x001CC444 | `HandleRemoteRewUp` | Known | Event handler |
| 0x001CC458 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x001CC470 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x001CC484 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x001CC49C | `HandleRemoteSelectUp` | Known | Event handler |
| 0x001CC66C | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x001CC684 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x001CC69C | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x001CC6B8 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x001CC6D0 | `HandleRemoteEvent` | Known | Event handler |
| 0x001CC6E4 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x001CC700 | `HandleAudioPlayPause` | Known | Event handler |
| 0x001CC718 | `HandleAudioNext` | Known | Event handler |
| 0x001CC728 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x001CC744 | `HandleAudioPrevious` | Known | Event handler |
| 0x001CC758 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x001CC958 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x001CC970 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x001CC988 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x001CC9A0 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x001CC9B4 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x001CC9CC | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x001CC9E4 | `HandleAudioStop` | Known | Event handler |
| 0x001CC9F4 | `HandleAudioPlay` | Known | Event handler |
| 0x001CCA04 | `HandleAudioPause` | Known | Event handler |
| 0x001CCA18 | `HandleAudioMute` | Known | Event handler |
| 0x001CCA28 | `HandleAudioNextChapter` | Known | Event handler |
| 0x001CCA40 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x001CCC60 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x001CCC78 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x001CCC90 | `HandleAudioShuffle` | Known | Event handler |
| 0x001CCCA4 | `HandleAudioRepeat` | Known | Event handler |
| 0x001CCCB8 | `HandleAudioFFDown` | Known | Event handler |
| 0x001CCCCC | `HandleAudioFFUp` | Known | Event handler |
| 0x001CCCDC | `HandleAudioRewDown` | Known | Event handler |
| 0x001CCCF0 | `HandleAudioRewUp` | Known | Event handler |
| 0x001CCD04 | `HandleVideoPlayPause` | Known | Event handler |
| 0x001CCD1C | `HandleVideoNext` | Known | Event handler |
| 0x001CCD2C | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x001CCD48 | `HandleVideoPrevious` | Known | Event handler |
| 0x001CCD5C | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x001CCF64 | `HandleVideoStop` | Known | Event handler |
| 0x001CCF74 | `HandleVideoPlay` | Known | Event handler |
| 0x001CCF84 | `HandleVideoPause` | Known | Event handler |
| 0x001CCF98 | `HandleVideoFFDown` | Known | Event handler |
| 0x001CCFAC | `HandleVideoFFUp` | Known | Event handler |
| 0x001CCFBC | `HandleVideoRewDown` | Known | Event handler |
| 0x001CCFD0 | `HandleVideoRewUp` | Known | Event handler |
| 0x001CCFE4 | `HandleVideoNextChapter` | Known | Event handler |
| 0x001CCFFC | `HandleVideoPrevChapter` | Known | Event handler |
| 0x001CD014 | `HandleVideoNextFrame` | Known | Event handler |
| 0x001CD02C | `HandleVideoPrevFrame` | Known | Event handler |
| 0x001CD044 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x001D05C8 | `HandleSelect` | Known | Event handler |
| 0x001D05DC | `HandleMenu` | Known | Event handler |
| 0x001D05E8 | `HandleLinkCancelOption` | Known | Event handler |
| 0x001D0600 | `HandleLinkNewRemote` | Known | Event handler |
| 0x001D0614 | `HandleCancelRemoteLinking` | Known | Event handler |
| 0x001D0974 | `HandleChoosePlaylistsPlay` | Known | Event handler |
| 0x001D0994 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x001D09B0 | `HandleNoneSelected` | Known | Event handler |
| 0x001D09C4 | `HandleNowPlayingSelected` | Known | Event handler |
| 0x001D09E0 | `HandleMenuKeyNop` | Known | Event handler |
| 0x001D09F4 | `HandlePlaylistSelected` | Known | Event handler |
| 0x001D11C0 | `HandlePauseWorkout` | Known | Event handler |
| 0x001D11D8 | `HandleEndWorkout` | Known | Event handler |
| 0x001D11EC | `HandleResumeWorkout` | Known | Event handler |
| 0x001D1200 | `HandleChooseMusic` | Known | Event handler |
| 0x001D1214 | `HandleMenuKeyNop` | Known | Event handler |
| 0x001DCCE4 | `HandleMainMenu` | Known | Event handler |
| 0x001E1228 | `HandlePowerSongSelected` | Known | Event handler |
| 0x001E1244 | `HandlePowerSongChosen` | Known | Event handler |
| 0x001E125C | `HandleSelectTimedWorkout` | Known | Event handler |
| 0x001E1AE0 | `HandleChooseRemoteSetting` | Known | Event handler |
| 0x001E1B00 | `HandleChooseRemoteLink` | Known | Event handler |
| 0x001E1B18 | `HandleChooseRemoteUnlink` | Known | Event handler |
| 0x001E1E54 | `HandleSelectResume` | Known | Event handler |
| 0x001E1E6C | `HandleEndWorkout` | Known | Event handler |
| 0x001E7FCC | `HandleSelect` | Known | Event handler |
| 0x001E8274 | `HandleMusicMenu` | Known | Event handler |
| 0x001E8534 | `HandleSelect` | Known | Event handler |
| 0x001E88B8 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x001E88D0 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x001E88F0 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x001E8914 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x001E8930 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x001E8DCC | `HandleWheel` | Known | Event handler |
| 0x001E8DDC | `HandlePlayPause` | Known | Event handler |
| 0x001E8DEC | `HandleSelectDown` | Known | Event handler |
| 0x001E8E00 | `HandleNext` | Known | Event handler |
| 0x001E8E0C | `HandlePrevious` | Known | Event handler |
| 0x001E8E1C | `HandleNextPushAndHold` | Known | Event handler |
| 0x001E8E34 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x001EE648 | `HandleChooseLast` | Known | Event handler |
| 0x001EE660 | `HandleChooseRecent` | Known | Event handler |
| 0x001EE674 | `HandleChooseWorkout` | Known | Event handler |
| 0x001EE688 | `HandleChooseBest` | Known | Event handler |
| 0x001EE69C | `HandleChooseRecentWorkout` | Known | Event handler |
| 0x001F0D24 | `HandleSelect` | Known | Event handler |
| 0x001F0D38 | `HandleMenu` | Known | Event handler |
| 0x001F8ACC | `HandleFrequencyChosen` | Known | Event handler |
| 0x001F8AE4 | `HandleDateChosen` | Known | Event handler |
| 0x001F8AF8 | `HandleTimeChosen` | Known | Event handler |
| 0x001F8B0C | `HandleSoundChosen` | Known | Event handler |
| 0x001F8B20 | `HandleLabelChosen` | Known | Event handler |
| 0x001F8B34 | `HandleDeleteChosen` | Known | Event handler |
| 0x001F9C14 | `HandleSelect` | Known | Event handler |
| 0x001FE530 | `HandlePrev` | Known | Event handler |
| 0x001FE540 | `HandleNext` | Known | Event handler |
| 0x001FE54C | `HandlePlayPause` | Known | Event handler |
| 0x001FED28 | `HandleChoosePowerPlay` | Known | Event handler |
| 0x001FED44 | `HandleChooseUnit` | Known | Event handler |
| 0x001FED58 | `HandleChooseVoicePrompts` | Known | Event handler |
| 0x00207680 | `HandleNextContact` | Known | Event handler |
| 0x00207698 | `HandlePreviousContact` | Known | Event handler |
| 0x0020AA5C | `HandleSelect` | Known | Event handler |
| 0x0020AD38 | `HandleListChoose` | Known | Event handler |
| 0x0020F76C | `HandleItemSelected` | Known | Event handler |
| 0x0020F964 | `HandleRadioRegion` | Known | Event handler |
| 0x0020FB4C | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x00210398 | `HandleSelect` | Known | Event handler |
| 0x002107E0 | `HandlePauseKey` | Known | Event handler |
| 0x002107F4 | `HandlePauseHold` | Known | Event handler |
| 0x00210804 | `HandlePauseKeyNop` | Known | Event handler |
| 0x00210818 | `HandleMenuKey` | Known | Event handler |
| 0x00210828 | `HandleMenuKeyNop` | Known | Event handler |
| 0x0021083C | `HandleWheel` | Known | Event handler |
| 0x0021088C | `HandleSelectKeyDown` | Known | Event handler |
| 0x002108A0 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x002108B8 | `HandlePowerPlay` | Known | Event handler |
| 0x0021555C | `HandlePlayPause` | Known | Event handler |
| 0x002167C8 | `HandleSelect` | Known | Event handler |
| 0x00216A58 | `HandleChooseCustomTimedWorkout` | Known | Event handler |
| 0x00216A7C | `HandleChooseTimedWorkoutFromList` | Known | Event handler |
| 0x00216AA0 | `HandleChooseCustomDistanceWorkout` | Known | Event handler |
| 0x00216AC4 | `HandleChooseDistanceWorkoutFromList` | Known | Event handler |
| 0x00216AE8 | `HandleChooseCustomCalorieWorkout` | Known | Event handler |
| 0x00216B0C | `HandleChooseCalorieWorkoutFromList` | Known | Event handler |
| 0x00219580 | `HandleDelete` | Known | Event handler |
| 0x00219594 | `HandleSelectLozinch` | Known | Event handler |
| 0x0021983C | `HandleSelect` | Known | Event handler |
| 0x00219B08 | `HandleTVOutChanged` | Known | Event handler |
| 0x00219B20 | `HandleTVSignalChanged` | Known | Event handler |
| 0x00219B38 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x00219B58 | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x00219B78 | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x00219B9C | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x00219BBC | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x0021A4A8 | `HandleBegin` | Known | Event handler |
| 0x0021D6D0 | `HandleSelectKey` | Known | Event handler |
| 0x0021D878 | `HandleSelect` | Known | Event handler |
| 0x0021E5F4 | `HandlePlayPause` | Known | Event handler |
| 0x0021E608 | `HandleWheel` | Known | Event handler |
| 0x0021E614 | `HandleWheelRating` | Known | Event handler |
| 0x0021E628 | `HandleWheelScrub` | Known | Event handler |
| 0x0021E63C | `HandleWheelVolume` | Known | Event handler |
| 0x0021E6FC | `HandleMenuKey` | Known | Event handler |
| 0x0021E768 | `HandleMenuLongpress` | Known | Event handler |
| 0x0021E77C | `HandleRentalWarningChoice` | Known | Event handler |
| 0x0021F384 | `HandleSelect` | Known | Event handler |
| 0x0021FC54 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00220B44 | `HandleSelect` | Known | Event handler |
| 0x00220B58 | `HandleHilite` | Known | Event handler |
| 0x00220B68 | `HandlePlayPause` | Known | Event handler |
| 0x00220B78 | `HandleAddToOTG` | Known | Event handler |
| 0x00220B88 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00221BC4 | `HandleWeightWheel` | Known | Event handler |
| 0x00221BDC | `HandleWeightSelect` | Known | Event handler |
| 0x00221BF0 | `HandleDistanceWheel` | Known | Event handler |
| 0x00221C04 | `HandleDistanceSelect` | Known | Event handler |
| 0x00221C1C | `HandleTimeWheel` | Known | Event handler |
| 0x00221C2C | `HandleTimeSelect` | Known | Event handler |
| 0x00221C40 | `HandleCaloriesWheel` | Known | Event handler |
| 0x00221C54 | `HandleCaloriesSelect` | Known | Event handler |
| 0x00222220 | `HandleSelect` | Known | Event handler |
| 0x00222234 | `HandleSelectFromLastWorkout` | Known | Event handler |
| 0x00224AE8 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x002252F8 | `HandleSelect` | Known | Event handler |
| 0x0022530C | `HandleWheel` | Known | Event handler |
| 0x00225318 | `HandleWheelProgress` | Known | Event handler |
| 0x0022532C | `HandleSelectProgress` | Known | Event handler |
| 0x00225344 | `HandleSelectVolume` | Known | Event handler |
| 0x00225358 | `HandleSelectScrub` | Known | Event handler |
| 0x0022536C | `HandleSelectRating` | Known | Event handler |
| 0x00225380 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x00225398 | `HandleSelectChapterArt` | Known | Event handler |
| 0x002253B0 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x002253CC | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x002253E8 | `HandleWheelBrightness` | Known | Event handler |
| 0x00225530 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x002270CC | `HandleSelect` | Known | Event handler |
| 0x002270DC | `HandleSelectRating` | Known | Event handler |
| 0x002270F0 | `HandleSelectProgress` | Known | Event handler |
| 0x00227108 | `HandleWheelProgress` | Known | Event handler |
| 0x0022711C | `HandleSelectScrub` | Known | Event handler |
| 0x00227130 | `HandleWheelBrightness` | Known | Event handler |
| 0x00227148 | `HandleNextTrackGotoScrub` | Known | Event handler |
| 0x00227164 | `HandlePrevTrackGotoScrub` | Known | Event handler |
| 0x00227180 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x0022A5E0 | `HandleSelectWalking` | Known | Event handler |
| 0x0022A5F8 | `HandleSelectRunning` | Known | Event handler |
| 0x0022E34C | `HandleLanguage` | Known | Event handler |
| 0x0022E35C | `HandleResetAllSettings` | Known | Event handler |
| 0x0022E374 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x0022E73C | `HandleUnlinkRemote` | Known | Event handler |
| 0x0022F22C | `HandleSelect` | Known | Event handler |
| 0x0022F45C | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x00230BA0 | `Handle400MetersRun` | Known | Event handler |
| 0x00230BB8 | `HandleCustomRun` | Known | Event handler |
| 0x00230BC8 | `HandleResetToDefault` | Known | Event handler |
| 0x00231028 | `HandleSelect_Basic` | Known | Event handler |
| 0x00231040 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x00233100 | `HandleSelect` | Known | Event handler |
| 0x0023329C | `HandleSelect` | Known | Event handler |
| 0x0023353C | `HandleNextDay` | Known | Event handler |
| 0x00233550 | `HandlePreviousDay` | Known | Event handler |
| 0x00233D54 | `HandleMusicHilited` | Known | Event handler |
| 0x00233D6C | `HandleVideosHilited` | Known | Event handler |
| 0x00233D80 | `HandlePodcastsHilited` | Known | Event handler |
| 0x00233D98 | `HandleGenericHilited` | Known | Event handler |
| 0x00233DB0 | `HandlePhotosHilited` | Known | Event handler |
| 0x00233DC4 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x00233DDC | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x00233DF8 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00233E10 | `HandleArtistsHilited` | Known | Event handler |
| 0x00233E28 | `HandleGenresHilited` | Known | Event handler |
| 0x00233E3C | `HandleAlbumsHilited` | Known | Event handler |
| 0x00233E50 | `HandleCompilationsHilited` | Known | Event handler |
| 0x00234024 | `HandleComposersHilited` | Known | Event handler |
| 0x0023403C | `HandleSongsHilited` | Known | Event handler |
| 0x00234050 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x00234068 | `HandleTVShowsHilited` | Known | Event handler |
| 0x00234080 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x0023409C | `HandleMoviesHilited` | Known | Event handler |
| 0x002340B0 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x002340CC | `HandleRentalsHilited` | Known | Event handler |
| 0x002340E4 | `HandleMusicSelected` | Known | Event handler |
| 0x002340F8 | `HandleVideosSelected` | Known | Event handler |
| 0x00234110 | `HandlePodcastsSelected` | Known | Event handler |
| 0x002342E0 | `HandlePhotosSelected` | Known | Event handler |
| 0x002342F8 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00234310 | `HandleSongsSelected` | Known | Event handler |
| 0x00234324 | `HandleAlbumsSelected` | Known | Event handler |
| 0x0023433C | `HandleCompilationsSelected` | Known | Event handler |
| 0x00234358 | `HandleArtistsSelected` | Known | Event handler |
| 0x00234370 | `HandleGenresSelected` | Known | Event handler |
| 0x00234388 | `HandleComposersSelected` | Known | Event handler |
| 0x002343A0 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x002343BC | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x002343D8 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x002345B0 | `HandleNowPlaying` | Known | Event handler |
| 0x002345C4 | `HandleTVShowsSelected` | Known | Event handler |
| 0x002345DC | `HandleMusicVideosSelected` | Known | Event handler |
| 0x002345F8 | `HandleMoviesSelected` | Known | Event handler |
| 0x00234610 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00234630 | `HandleRentalsSelected` | Known | Event handler |
| 0x00234648 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00234660 | `HandleLock` | Known | Event handler |
| 0x0023466C | `HandleBacklightSelected` | Known | Event handler |
| 0x00234684 | `HandleSleepSelected` | Known | Event handler |
| 0x00234698 | `HandleNikePlusSelected` | Known | Event handler |
| 0x00236F5C | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00237570 | `Handle400MetersWalk` | Known | Event handler |
| 0x00237588 | `HandleCustomWalk` | Known | Event handler |
| 0x0023759C | `HandleResetToDefault` | Known | Event handler |
| 0x00237888 | `HandleSelect` | Known | Event handler |
| 0x00237B38 | `HandleWheel` | Known | Event handler |
| 0x00239008 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x00239260 | `HandleNextDay` | Known | Event handler |
| 0x00239274 | `HandlePreviousDay` | Known | Event handler |
| 0x002394BC | `HandleSelect` | Known | Event handler |
| 0x00239758 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0023C084 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x0023C0A0 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x0023D008 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0023D6E8 | `HandleSelect` | Known | Event handler |
| 0x0023DDB4 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x002798D8 | `HandleDeleteClock` | Known | Event handler |
| 0x002798F0 | `HandleSelectClock` | Known | Event handler |
| 0x00279904 | `HandleHilited` | Known | Event handler |
| 0x00279914 | `HandleWheel` | Known | Event handler |
| 0x00279920 | `HandleSelectLozinch` | Known | Event handler |
| 0x003F9636 | `HandleAudioFFDown` | Known | Event handler |
| 0x003F965F | `HandleAudioFFUp` | Known | Event handler |
| 0x003F968A | `HandleAudioMute` | Known | Event handler |
| 0x003F96BD | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x003F96F2 | `HandleAudioNext` | Known | Event handler |
| 0x003F9722 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x003F9759 | `HandleAudioNextChapter` | Known | Event handler |
| 0x003F9793 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x003F97C7 | `HandleAudioPause` | Known | Event handler |
| 0x003F97F3 | `HandleAudioPlay` | Known | Event handler |
| 0x003F9821 | `HandleAudioPlayPause` | Known | Event handler |
| 0x003F9859 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x003F9892 | `HandleAudioPrevious` | Known | Event handler |
| 0x003F98C6 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x003F98FD | `HandleAudioPrevChapter` | Known | Event handler |
| 0x003F9937 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x003F996C | `HandleAudioRepeat` | Known | Event handler |
| 0x003F9998 | `HandleAudioRewDown` | Known | Event handler |
| 0x003F99C3 | `HandleAudioRewUp` | Known | Event handler |
| 0x003F99F2 | `HandleAudioShuffle` | Known | Event handler |
| 0x003F9A20 | `HandleAudioStop` | Known | Event handler |
| 0x003F9A51 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x003F9A86 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x003F9ABD | `HandleAudioVolumeUp` | Known | Event handler |
| 0x003F9AEE | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x003F9BA7 | `HandleNextPressAndHold` | Known | Event handler |
| 0x003F9BD8 | `HandleNext` | Known | Event handler |
| 0x003F9C10 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x003F9C4B | `HandlePlayPause` | Known | Event handler |
| 0x003F9C7F | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x003F9CB4 | `HandlePrevious` | Known | Event handler |
| 0x003F9D41 | `HandleRemoteBacklight` | Known | Event handler |
| 0x003F9D79 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x003F9DB3 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x003F9DEC | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x003F9E21 | `HandleRemoteEvent` | Known | Event handler |
| 0x003F9E4D | `HandleRemoteFFDown` | Known | Event handler |
| 0x003F9E78 | `HandleRemoteFFUp` | Known | Event handler |
| 0x003F9EA5 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x003F9ED4 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x003F9F03 | `HandleRemoteMute` | Known | Event handler |
| 0x003F9F35 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x003F9F6E | `HandleRemoteNextChapter` | Known | Event handler |
| 0x003F9FAA | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x003F9FEA | `HandleRemoteOff` | Known | Event handler |
| 0x003FA013 | `HandleRemoteOff` | Known | Event handler |
| 0x003FA03D | `HandleRemoteOn` | Known | Event handler |
| 0x003FA069 | `HandleRemotePause` | Known | Event handler |
| 0x003FA097 | `HandleRemotePlay` | Known | Event handler |
| 0x003FA0D5 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x003FA116 | `HandleRemotePlayPause` | Known | Event handler |
| 0x003FA14D | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x003FA186 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x003FA1C2 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x003FA1F9 | `HandleRemoteRepeat` | Known | Event handler |
| 0x003FA227 | `HandleRemoteRewDown` | Known | Event handler |
| 0x003FA254 | `HandleRemoteRewUp` | Known | Event handler |
| 0x003FA284 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x003FA2B7 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x003FA2EB | `HandleRemoteShuffle` | Known | Event handler |
| 0x003FA31B | `HandleRemoteStop` | Known | Event handler |
| 0x003FA34B | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x003FA380 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x003FA3B8 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x003FA3EF | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x003FA428 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x003FA45B | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x003FA490 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x003FA4C3 | `HandleVideoFFDown` | Known | Event handler |
| 0x003FA4EC | `HandleVideoFFUp` | Known | Event handler |
| 0x003FA51F | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x003FA554 | `HandleVideoNext` | Known | Event handler |
| 0x003FA586 | `HandleVideoNextChapter` | Known | Event handler |
| 0x003FA5BD | `HandleVideoNextFrame` | Known | Event handler |
| 0x003FA5EE | `HandleVideoPause` | Known | Event handler |
| 0x003FA61A | `HandleVideoPlay` | Known | Event handler |
| 0x003FA648 | `HandleVideoPlayPause` | Known | Event handler |
| 0x003FA680 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x003FA6B9 | `HandleVideoPrevious` | Known | Event handler |
| 0x003FA6EF | `HandleVideoPrevChapter` | Known | Event handler |
| 0x003FA726 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x003FA755 | `HandleVideoRewDown` | Known | Event handler |
| 0x003FA780 | `HandleVideoRewUp` | Known | Event handler |
| 0x003FA7AC | `HandleVideoStop` | Known | Event handler |
| 0x0074A672 | `HandleAddressBook` | Known | Event handler |
| 0x0074AC06 | `HandleSelect` | Known | Event handler |
| 0x0074AC41 | `HandleHilite` | Known | Event handler |
| 0x0074ACC2 | `HandleSelectRegion` | Known | Event handler |
| 0x0074AD62 | `HandleSelectRegion` | Known | Event handler |
| 0x0074ADFE | `HandleSelectRegion` | Known | Event handler |
| 0x0074AEA2 | `HandleSelectRegion` | Known | Event handler |
| 0x0074AF48 | `HandleSelectRegion` | Known | Event handler |
| 0x0074AFE8 | `HandleSelectRegion` | Known | Event handler |
| 0x0074B094 | `HandleSelectRegion` | Known | Event handler |
| 0x0074B136 | `HandleSelectRegion` | Known | Event handler |
| 0x0074B1E6 | `HandleSelectCity` | Known | Event handler |
| 0x0074B252 | `HandleHighlightCity` | Known | Event handler |
| 0x0074B28B | `HandleSelectCity` | Known | Event handler |
| 0x0074B2F7 | `HandleHighlightCity` | Known | Event handler |
| 0x0074B330 | `HandleSelectCity` | Known | Event handler |
| 0x0074B39C | `HandleHighlightCity` | Known | Event handler |
| 0x0074B3D5 | `HandleSelectCity` | Known | Event handler |
| 0x0074B441 | `HandleHighlightCity` | Known | Event handler |
| 0x0074B47A | `HandleSelectCity` | Known | Event handler |
| 0x0074B4E6 | `HandleHighlightCity` | Known | Event handler |
| 0x0074B51F | `HandleSelectCity` | Known | Event handler |
| 0x0074B58B | `HandleHighlightCity` | Known | Event handler |
| 0x0074B5C4 | `HandleSelectCity` | Known | Event handler |
| 0x0074B630 | `HandleHighlightCity` | Known | Event handler |
| 0x0074B669 | `HandleSelectCity` | Known | Event handler |
| 0x0074B6D5 | `HandleHighlightCity` | Known | Event handler |
| 0x0074B70E | `HandleSelectCity` | Known | Event handler |
| 0x0074B77A | `HandleHighlightCity` | Known | Event handler |
| 0x0074B7B3 | `HandleSelectCity` | Known | Event handler |
| 0x0074B81F | `HandleHighlightCity` | Known | Event handler |
| 0x0074B858 | `HandleSelectCity` | Known | Event handler |
| 0x0074B8C4 | `HandleHighlightCity` | Known | Event handler |
| 0x0074B8FD | `HandleSelectCity` | Known | Event handler |
| 0x0074B969 | `HandleHighlightCity` | Known | Event handler |
| 0x0074B9A2 | `HandleSelectCity` | Known | Event handler |
| 0x0074BA0E | `HandleHighlightCity` | Known | Event handler |
| 0x0074BA47 | `HandleSelectCity` | Known | Event handler |
| 0x0074BAB3 | `HandleHighlightCity` | Known | Event handler |
| 0x0074BAEC | `HandleSelectCity` | Known | Event handler |
| 0x0074BB58 | `HandleHighlightCity` | Known | Event handler |
| 0x0074BB91 | `HandleSelectCity` | Known | Event handler |
| 0x0074BBFD | `HandleHighlightCity` | Known | Event handler |
| 0x0074BC36 | `HandleSelectCity` | Known | Event handler |
| 0x0074BCA2 | `HandleHighlightCity` | Known | Event handler |
| 0x0074BCDB | `HandleSelectCity` | Known | Event handler |
| 0x0074BD47 | `HandleHighlightCity` | Known | Event handler |
| 0x0074BD80 | `HandleSelectCity` | Known | Event handler |
| 0x0074BDEC | `HandleHighlightCity` | Known | Event handler |
| 0x0074BE25 | `HandleSelectCity` | Known | Event handler |
| 0x0074BE91 | `HandleHighlightCity` | Known | Event handler |
| 0x0074BECA | `HandleSelectCity` | Known | Event handler |
| 0x0074BF36 | `HandleHighlightCity` | Known | Event handler |
| 0x0074BF6F | `HandleSelectCity` | Known | Event handler |
| 0x0074BFDB | `HandleHighlightCity` | Known | Event handler |
| 0x0074C014 | `HandleSelectCity` | Known | Event handler |
| 0x0074C080 | `HandleHighlightCity` | Known | Event handler |
| 0x0074C0B9 | `HandleSelectCity` | Known | Event handler |
| 0x0074C125 | `HandleHighlightCity` | Known | Event handler |
| 0x0074C15E | `HandleSelectCity` | Known | Event handler |
| 0x0074C1CA | `HandleHighlightCity` | Known | Event handler |
| 0x0074C203 | `HandleSelectCity` | Known | Event handler |
| 0x0074C26F | `HandleHighlightCity` | Known | Event handler |
| 0x0074C2A8 | `HandleSelectCity` | Known | Event handler |
| 0x0074C314 | `HandleHighlightCity` | Known | Event handler |
| 0x0074C34D | `HandleSelectCity` | Known | Event handler |
| 0x0074C3B9 | `HandleHighlightCity` | Known | Event handler |
| 0x0074C3F2 | `HandleSelectCity` | Known | Event handler |
| 0x0074C45E | `HandleHighlightCity` | Known | Event handler |
| 0x0074C497 | `HandleSelectCity` | Known | Event handler |
| 0x0074C503 | `HandleHighlightCity` | Known | Event handler |
| 0x0074C53C | `HandleSelectCity` | Known | Event handler |
| 0x0074C5A8 | `HandleHighlightCity` | Known | Event handler |
| 0x0074C5E6 | `HandleSelectCity` | Known | Event handler |
| 0x0074C652 | `HandleHighlightCity` | Known | Event handler |
| 0x0074C68B | `HandleSelectCity` | Known | Event handler |
| 0x0074C6F7 | `HandleHighlightCity` | Known | Event handler |
| 0x0074C730 | `HandleSelectCity` | Known | Event handler |
| 0x0074C79C | `HandleHighlightCity` | Known | Event handler |
| 0x0074C7D5 | `HandleSelectCity` | Known | Event handler |
| 0x0074C841 | `HandleHighlightCity` | Known | Event handler |
| 0x0074C87A | `HandleSelectCity` | Known | Event handler |
| 0x0074C8E6 | `HandleHighlightCity` | Known | Event handler |
| 0x0074C91F | `HandleSelectCity` | Known | Event handler |
| 0x0074C98B | `HandleHighlightCity` | Known | Event handler |
| 0x0074C9C4 | `HandleSelectCity` | Known | Event handler |
| 0x0074CA30 | `HandleHighlightCity` | Known | Event handler |
| 0x0074CA69 | `HandleSelectCity` | Known | Event handler |
| 0x0074CAD5 | `HandleHighlightCity` | Known | Event handler |
| 0x0074CB0E | `HandleSelectCity` | Known | Event handler |
| 0x0074CB7A | `HandleHighlightCity` | Known | Event handler |
| 0x0074CBB3 | `HandleSelectCity` | Known | Event handler |
| 0x0074CC1F | `HandleHighlightCity` | Known | Event handler |
| 0x0074CC58 | `HandleSelectCity` | Known | Event handler |
| 0x0074CCC4 | `HandleHighlightCity` | Known | Event handler |
| 0x0074CCFD | `HandleSelectCity` | Known | Event handler |
| 0x0074CD69 | `HandleHighlightCity` | Known | Event handler |
| 0x0074CDA2 | `HandleSelectCity` | Known | Event handler |
| 0x0074CE0E | `HandleHighlightCity` | Known | Event handler |
| 0x0074CE47 | `HandleSelectCity` | Known | Event handler |
| 0x0074CEB3 | `HandleHighlightCity` | Known | Event handler |
| 0x0074CEEC | `HandleSelectCity` | Known | Event handler |
| 0x0074CF58 | `HandleHighlightCity` | Known | Event handler |
| 0x0074CF91 | `HandleSelectCity` | Known | Event handler |
| 0x0074CFFD | `HandleHighlightCity` | Known | Event handler |
| 0x0074D036 | `HandleSelectCity` | Known | Event handler |
| 0x0074D0A2 | `HandleHighlightCity` | Known | Event handler |
| 0x0074D0DB | `HandleSelectCity` | Known | Event handler |
| 0x0074D147 | `HandleHighlightCity` | Known | Event handler |
| 0x0074D180 | `HandleSelectCity` | Known | Event handler |
| 0x0074D1EC | `HandleHighlightCity` | Known | Event handler |
| 0x0074D225 | `HandleSelectCity` | Known | Event handler |
| 0x0074D291 | `HandleHighlightCity` | Known | Event handler |
| 0x0074D2CA | `HandleSelectCity` | Known | Event handler |
| 0x0074D336 | `HandleHighlightCity` | Known | Event handler |
| 0x0074D36F | `HandleSelectCity` | Known | Event handler |
| 0x0074D3DB | `HandleHighlightCity` | Known | Event handler |
| 0x0074D414 | `HandleSelectCity` | Known | Event handler |
| 0x0074D480 | `HandleHighlightCity` | Known | Event handler |
| 0x0074D4B9 | `HandleSelectCity` | Known | Event handler |
| 0x0074D525 | `HandleHighlightCity` | Known | Event handler |
| 0x0074D55E | `HandleSelectCity` | Known | Event handler |
| 0x0074D5CA | `HandleHighlightCity` | Known | Event handler |
| 0x0074D603 | `HandleSelectCity` | Known | Event handler |
| 0x0074D66F | `HandleHighlightCity` | Known | Event handler |
| 0x0074D6A8 | `HandleSelectCity` | Known | Event handler |
| 0x0074D714 | `HandleHighlightCity` | Known | Event handler |
| 0x0074D74D | `HandleSelectCity` | Known | Event handler |
| 0x0074D7B9 | `HandleHighlightCity` | Known | Event handler |
| 0x0074D7F2 | `HandleSelectCity` | Known | Event handler |
| 0x0074D85E | `HandleHighlightCity` | Known | Event handler |
| 0x0074D897 | `HandleSelectCity` | Known | Event handler |
| 0x0074D903 | `HandleHighlightCity` | Known | Event handler |
| 0x0074D93C | `HandleSelectCity` | Known | Event handler |
| 0x0074D9A8 | `HandleHighlightCity` | Known | Event handler |
| 0x0074D9E1 | `HandleSelectCity` | Known | Event handler |
| 0x0074DA4D | `HandleHighlightCity` | Known | Event handler |
| 0x0074DA86 | `HandleSelectCity` | Known | Event handler |
| 0x0074DAF2 | `HandleHighlightCity` | Known | Event handler |
| 0x0074DB2B | `HandleSelectCity` | Known | Event handler |
| 0x0074DB97 | `HandleHighlightCity` | Known | Event handler |
| 0x0074DBD0 | `HandleSelectCity` | Known | Event handler |
| 0x0074DC3C | `HandleHighlightCity` | Known | Event handler |
| 0x0074DC75 | `HandleSelectCity` | Known | Event handler |
| 0x0074DCE1 | `HandleHighlightCity` | Known | Event handler |
| 0x0074DD1A | `HandleSelectCity` | Known | Event handler |
| 0x0074DD86 | `HandleHighlightCity` | Known | Event handler |
| 0x0074DDBF | `HandleSelectCity` | Known | Event handler |
| 0x0074DE2B | `HandleHighlightCity` | Known | Event handler |
| 0x0074DE64 | `HandleSelectCity` | Known | Event handler |
| 0x0074DED0 | `HandleHighlightCity` | Known | Event handler |
| 0x0074DF09 | `HandleSelectCity` | Known | Event handler |
| 0x0074DF75 | `HandleHighlightCity` | Known | Event handler |
| 0x0074DFAE | `HandleSelectCity` | Known | Event handler |
| 0x0074E01A | `HandleHighlightCity` | Known | Event handler |
| 0x0074E053 | `HandleSelectCity` | Known | Event handler |
| 0x0074E0BF | `HandleHighlightCity` | Known | Event handler |
| 0x0074E0F8 | `HandleSelectCity` | Known | Event handler |
| 0x0074E164 | `HandleHighlightCity` | Known | Event handler |
| 0x0074E19D | `HandleSelectCity` | Known | Event handler |
| 0x0074E209 | `HandleHighlightCity` | Known | Event handler |
| 0x0074E242 | `HandleSelectCity` | Known | Event handler |
| 0x0074E2AE | `HandleHighlightCity` | Known | Event handler |
| 0x0074E2E7 | `HandleSelectCity` | Known | Event handler |
| 0x0074E353 | `HandleHighlightCity` | Known | Event handler |
| 0x0074E38C | `HandleSelectCity` | Known | Event handler |
| 0x0074E3F8 | `HandleHighlightCity` | Known | Event handler |
| 0x0074E431 | `HandleSelectCity` | Known | Event handler |
| 0x0074E49D | `HandleHighlightCity` | Known | Event handler |
| 0x0074E4D6 | `HandleSelectCity` | Known | Event handler |
| 0x0074E542 | `HandleHighlightCity` | Known | Event handler |
| 0x0074E57B | `HandleSelectCity` | Known | Event handler |
| 0x0074E5E7 | `HandleHighlightCity` | Known | Event handler |
| 0x0074E620 | `HandleSelectCity` | Known | Event handler |
| 0x0074E68C | `HandleHighlightCity` | Known | Event handler |
| 0x0074E6C5 | `HandleSelectCity` | Known | Event handler |
| 0x0074E731 | `HandleHighlightCity` | Known | Event handler |
| 0x0074E76A | `HandleSelectCity` | Known | Event handler |
| 0x0074E7D6 | `HandleHighlightCity` | Known | Event handler |
| 0x0074E80F | `HandleSelectCity` | Known | Event handler |
| 0x0074E87B | `HandleHighlightCity` | Known | Event handler |
| 0x0074E8B4 | `HandleSelectCity` | Known | Event handler |
| 0x0074E920 | `HandleHighlightCity` | Known | Event handler |
| 0x0074E959 | `HandleSelectCity` | Known | Event handler |
| 0x0074E9C5 | `HandleHighlightCity` | Known | Event handler |
| 0x0074E9FE | `HandleSelectCity` | Known | Event handler |
| 0x0074EA6A | `HandleHighlightCity` | Known | Event handler |
| 0x0074EAAA | `HandleSelectCity` | Known | Event handler |
| 0x0074EB16 | `HandleHighlightCity` | Known | Event handler |
| 0x0074EB4F | `HandleSelectCity` | Known | Event handler |
| 0x0074EBBB | `HandleHighlightCity` | Known | Event handler |
| 0x0074EBF4 | `HandleSelectCity` | Known | Event handler |
| 0x0074EC60 | `HandleHighlightCity` | Known | Event handler |
| 0x0074EC9E | `HandleSelectCity` | Known | Event handler |
| 0x0074ED0A | `HandleHighlightCity` | Known | Event handler |
| 0x0074ED43 | `HandleSelectCity` | Known | Event handler |
| 0x0074EDAF | `HandleHighlightCity` | Known | Event handler |
| 0x0074EDE8 | `HandleSelectCity` | Known | Event handler |
| 0x0074EE54 | `HandleHighlightCity` | Known | Event handler |
| 0x0074EE8D | `HandleSelectCity` | Known | Event handler |
| 0x0074EEF9 | `HandleHighlightCity` | Known | Event handler |
| 0x0074EF32 | `HandleSelectCity` | Known | Event handler |
| 0x0074EF9E | `HandleHighlightCity` | Known | Event handler |
| 0x0074EFD7 | `HandleSelectCity` | Known | Event handler |
| 0x0074F043 | `HandleHighlightCity` | Known | Event handler |
| 0x0074F07C | `HandleSelectCity` | Known | Event handler |
| 0x0074F0E8 | `HandleHighlightCity` | Known | Event handler |
| 0x0074F121 | `HandleSelectCity` | Known | Event handler |
| 0x0074F18D | `HandleHighlightCity` | Known | Event handler |
| 0x0074F1CA | `HandleSelectCity` | Known | Event handler |
| 0x0074F236 | `HandleHighlightCity` | Known | Event handler |
| 0x0074F26F | `HandleSelectCity` | Known | Event handler |
| 0x0074F2DB | `HandleHighlightCity` | Known | Event handler |
| 0x0074F314 | `HandleSelectCity` | Known | Event handler |
| 0x0074F380 | `HandleHighlightCity` | Known | Event handler |
| 0x0074F3B9 | `HandleSelectCity` | Known | Event handler |
| 0x0074F425 | `HandleHighlightCity` | Known | Event handler |
| 0x0074F45E | `HandleSelectCity` | Known | Event handler |
| 0x0074F4CA | `HandleHighlightCity` | Known | Event handler |
| 0x0074F503 | `HandleSelectCity` | Known | Event handler |
| 0x0074F56F | `HandleHighlightCity` | Known | Event handler |
| 0x0074F5A8 | `HandleSelectCity` | Known | Event handler |
| 0x0074F614 | `HandleHighlightCity` | Known | Event handler |
| 0x0074F64D | `HandleSelectCity` | Known | Event handler |
| 0x0074F6B9 | `HandleHighlightCity` | Known | Event handler |
| 0x0074F6F2 | `HandleSelectCity` | Known | Event handler |
| 0x0074F75E | `HandleHighlightCity` | Known | Event handler |
| 0x0074F797 | `HandleSelectCity` | Known | Event handler |
| 0x0074F803 | `HandleHighlightCity` | Known | Event handler |
| 0x0074F83C | `HandleSelectCity` | Known | Event handler |
| 0x0074F8A8 | `HandleHighlightCity` | Known | Event handler |
| 0x0074F8E1 | `HandleSelectCity` | Known | Event handler |
| 0x0074F94D | `HandleHighlightCity` | Known | Event handler |
| 0x0074F986 | `HandleSelectCity` | Known | Event handler |
| 0x0074F9F2 | `HandleHighlightCity` | Known | Event handler |
| 0x0074FA2B | `HandleSelectCity` | Known | Event handler |
| 0x0074FA97 | `HandleHighlightCity` | Known | Event handler |
| 0x0074FAD0 | `HandleSelectCity` | Known | Event handler |
| 0x0074FB3C | `HandleHighlightCity` | Known | Event handler |
| 0x0074FB75 | `HandleSelectCity` | Known | Event handler |
| 0x0074FBE1 | `HandleHighlightCity` | Known | Event handler |
| 0x0074FC1A | `HandleSelectCity` | Known | Event handler |
| 0x0074FC86 | `HandleHighlightCity` | Known | Event handler |
| 0x0074FCBF | `HandleSelectCity` | Known | Event handler |
| 0x0074FD2B | `HandleHighlightCity` | Known | Event handler |
| 0x0074FD64 | `HandleSelectCity` | Known | Event handler |
| 0x0074FDD0 | `HandleHighlightCity` | Known | Event handler |
| 0x0074FE09 | `HandleSelectCity` | Known | Event handler |
| 0x0074FE75 | `HandleHighlightCity` | Known | Event handler |
| 0x0074FEAE | `HandleSelectCity` | Known | Event handler |
| 0x0074FF1A | `HandleHighlightCity` | Known | Event handler |
| 0x0074FF53 | `HandleSelectCity` | Known | Event handler |
| 0x0074FFBF | `HandleHighlightCity` | Known | Event handler |
| 0x0074FFF8 | `HandleSelectCity` | Known | Event handler |
| 0x00750064 | `HandleHighlightCity` | Known | Event handler |
| 0x0075009D | `HandleSelectCity` | Known | Event handler |
| 0x00750109 | `HandleHighlightCity` | Known | Event handler |
| 0x00750142 | `HandleSelectCity` | Known | Event handler |
| 0x007501AE | `HandleHighlightCity` | Known | Event handler |
| 0x007501E7 | `HandleSelectCity` | Known | Event handler |
| 0x00750253 | `HandleHighlightCity` | Known | Event handler |
| 0x0075028C | `HandleSelectCity` | Known | Event handler |
| 0x007502F8 | `HandleHighlightCity` | Known | Event handler |
| 0x00750331 | `HandleSelectCity` | Known | Event handler |
| 0x0075039D | `HandleHighlightCity` | Known | Event handler |
| 0x007503D6 | `HandleSelectCity` | Known | Event handler |
| 0x00750442 | `HandleHighlightCity` | Known | Event handler |
| 0x0075047B | `HandleSelectCity` | Known | Event handler |
| 0x007504E7 | `HandleHighlightCity` | Known | Event handler |
| 0x00750520 | `HandleSelectCity` | Known | Event handler |
| 0x0075058C | `HandleHighlightCity` | Known | Event handler |
| 0x007505C5 | `HandleSelectCity` | Known | Event handler |
| 0x00750631 | `HandleHighlightCity` | Known | Event handler |
| 0x0075066A | `HandleSelectCity` | Known | Event handler |
| 0x007506D6 | `HandleHighlightCity` | Known | Event handler |
| 0x0075070F | `HandleSelectCity` | Known | Event handler |
| 0x0075077B | `HandleHighlightCity` | Known | Event handler |
| 0x007507BA | `HandleSelectCity` | Known | Event handler |
| 0x00750826 | `HandleHighlightCity` | Known | Event handler |
| 0x0075085F | `HandleSelectCity` | Known | Event handler |
| 0x007508CB | `HandleHighlightCity` | Known | Event handler |
| 0x00750904 | `HandleSelectCity` | Known | Event handler |
| 0x00750970 | `HandleHighlightCity` | Known | Event handler |
| 0x007509A9 | `HandleSelectCity` | Known | Event handler |
| 0x00750A15 | `HandleHighlightCity` | Known | Event handler |
| 0x00750A4E | `HandleSelectCity` | Known | Event handler |
| 0x00750ABA | `HandleHighlightCity` | Known | Event handler |
| 0x00750AF3 | `HandleSelectCity` | Known | Event handler |
| 0x00750B5F | `HandleHighlightCity` | Known | Event handler |
| 0x00750B98 | `HandleSelectCity` | Known | Event handler |
| 0x00750C04 | `HandleHighlightCity` | Known | Event handler |
| 0x00750C3D | `HandleSelectCity` | Known | Event handler |
| 0x00750CA9 | `HandleHighlightCity` | Known | Event handler |
| 0x00750CE2 | `HandleSelectCity` | Known | Event handler |
| 0x00750D4E | `HandleHighlightCity` | Known | Event handler |
| 0x00750D87 | `HandleSelectCity` | Known | Event handler |
| 0x00750DF3 | `HandleHighlightCity` | Known | Event handler |
| 0x00750E2C | `HandleSelectCity` | Known | Event handler |
| 0x00750E98 | `HandleHighlightCity` | Known | Event handler |
| 0x00750ED1 | `HandleSelectCity` | Known | Event handler |
| 0x00750F3D | `HandleHighlightCity` | Known | Event handler |
| 0x00750F76 | `HandleSelectCity` | Known | Event handler |
| 0x00750FE2 | `HandleHighlightCity` | Known | Event handler |
| 0x0075101B | `HandleSelectCity` | Known | Event handler |
| 0x00751087 | `HandleHighlightCity` | Known | Event handler |
| 0x007510C0 | `HandleSelectCity` | Known | Event handler |
| 0x0075112C | `HandleHighlightCity` | Known | Event handler |
| 0x00751165 | `HandleSelectCity` | Known | Event handler |
| 0x007511D1 | `HandleHighlightCity` | Known | Event handler |
| 0x0075120A | `HandleSelectCity` | Known | Event handler |
| 0x00751276 | `HandleHighlightCity` | Known | Event handler |
| 0x007512AF | `HandleSelectCity` | Known | Event handler |
| 0x0075131B | `HandleHighlightCity` | Known | Event handler |
| 0x00751354 | `HandleSelectCity` | Known | Event handler |
| 0x007513C0 | `HandleHighlightCity` | Known | Event handler |
| 0x007513F9 | `HandleSelectCity` | Known | Event handler |
| 0x00751465 | `HandleHighlightCity` | Known | Event handler |
| 0x0075149E | `HandleSelectCity` | Known | Event handler |
| 0x0075150A | `HandleHighlightCity` | Known | Event handler |
| 0x00751543 | `HandleSelectCity` | Known | Event handler |
| 0x007515AF | `HandleHighlightCity` | Known | Event handler |
| 0x007515E8 | `HandleSelectCity` | Known | Event handler |
| 0x00751654 | `HandleHighlightCity` | Known | Event handler |
| 0x0075168D | `HandleSelectCity` | Known | Event handler |
| 0x007516F9 | `HandleHighlightCity` | Known | Event handler |
| 0x00751732 | `HandleSelectCity` | Known | Event handler |
| 0x0075179E | `HandleHighlightCity` | Known | Event handler |
| 0x007517D7 | `HandleSelectCity` | Known | Event handler |
| 0x00751843 | `HandleHighlightCity` | Known | Event handler |
| 0x0075187C | `HandleSelectCity` | Known | Event handler |
| 0x007518E8 | `HandleHighlightCity` | Known | Event handler |
| 0x00751921 | `HandleSelectCity` | Known | Event handler |
| 0x0075198D | `HandleHighlightCity` | Known | Event handler |
| 0x007519C6 | `HandleSelectCity` | Known | Event handler |
| 0x00751A32 | `HandleHighlightCity` | Known | Event handler |
| 0x00751A6B | `HandleSelectCity` | Known | Event handler |
| 0x00751AD7 | `HandleHighlightCity` | Known | Event handler |
| 0x00751B10 | `HandleSelectCity` | Known | Event handler |
| 0x00751B7C | `HandleHighlightCity` | Known | Event handler |
| 0x00751BB5 | `HandleSelectCity` | Known | Event handler |
| 0x00751C21 | `HandleHighlightCity` | Known | Event handler |
| 0x00751C5A | `HandleSelectCity` | Known | Event handler |
| 0x00751CC6 | `HandleHighlightCity` | Known | Event handler |
| 0x00751CFF | `HandleSelectCity` | Known | Event handler |
| 0x00751D6B | `HandleHighlightCity` | Known | Event handler |
| 0x00751DA4 | `HandleSelectCity` | Known | Event handler |
| 0x00751E10 | `HandleHighlightCity` | Known | Event handler |
| 0x00751E49 | `HandleSelectCity` | Known | Event handler |
| 0x00751EB5 | `HandleHighlightCity` | Known | Event handler |
| 0x00751EEE | `HandleSelectCity` | Known | Event handler |
| 0x00751F5A | `HandleHighlightCity` | Known | Event handler |
| 0x00751F93 | `HandleSelectCity` | Known | Event handler |
| 0x00751FFF | `HandleHighlightCity` | Known | Event handler |
| 0x00752038 | `HandleSelectCity` | Known | Event handler |
| 0x007520A4 | `HandleHighlightCity` | Known | Event handler |
| 0x007520DD | `HandleSelectCity` | Known | Event handler |
| 0x00752149 | `HandleHighlightCity` | Known | Event handler |
| 0x00752182 | `HandleSelectCity` | Known | Event handler |
| 0x007521EE | `HandleHighlightCity` | Known | Event handler |
| 0x00752227 | `HandleSelectCity` | Known | Event handler |
| 0x00752293 | `HandleHighlightCity` | Known | Event handler |
| 0x007522CC | `HandleSelectCity` | Known | Event handler |
| 0x00752338 | `HandleHighlightCity` | Known | Event handler |
| 0x00752371 | `HandleSelectCity` | Known | Event handler |
| 0x007523DD | `HandleHighlightCity` | Known | Event handler |
| 0x00752416 | `HandleSelectCity` | Known | Event handler |
| 0x00752482 | `HandleHighlightCity` | Known | Event handler |
| 0x007524BB | `HandleSelectCity` | Known | Event handler |
| 0x00752527 | `HandleHighlightCity` | Known | Event handler |
| 0x00752560 | `HandleSelectCity` | Known | Event handler |
| 0x007525CC | `HandleHighlightCity` | Known | Event handler |
| 0x00752605 | `HandleSelectCity` | Known | Event handler |
| 0x00752671 | `HandleHighlightCity` | Known | Event handler |
| 0x007526AA | `HandleSelectCity` | Known | Event handler |
| 0x00752716 | `HandleHighlightCity` | Known | Event handler |
| 0x0075274F | `HandleSelectCity` | Known | Event handler |
| 0x007527BB | `HandleHighlightCity` | Known | Event handler |
| 0x007527FA | `HandleSelectCity` | Known | Event handler |
| 0x00752866 | `HandleHighlightCity` | Known | Event handler |
| 0x0075289F | `HandleSelectCity` | Known | Event handler |
| 0x0075290B | `HandleHighlightCity` | Known | Event handler |
| 0x00752944 | `HandleSelectCity` | Known | Event handler |
| 0x007529B0 | `HandleHighlightCity` | Known | Event handler |
| 0x007529E9 | `HandleSelectCity` | Known | Event handler |
| 0x00752A55 | `HandleHighlightCity` | Known | Event handler |
| 0x00752A8E | `HandleSelectCity` | Known | Event handler |
| 0x00752AFA | `HandleHighlightCity` | Known | Event handler |
| 0x00752B3A | `HandleSelectCity` | Known | Event handler |
| 0x00752BA6 | `HandleHighlightCity` | Known | Event handler |
| 0x00752BDF | `HandleSelectCity` | Known | Event handler |
| 0x00752C4B | `HandleHighlightCity` | Known | Event handler |
| 0x00752C84 | `HandleSelectCity` | Known | Event handler |
| 0x00752CF0 | `HandleHighlightCity` | Known | Event handler |
| 0x00752D29 | `HandleSelectCity` | Known | Event handler |
| 0x00752D95 | `HandleHighlightCity` | Known | Event handler |
| 0x00752DCE | `HandleSelectCity` | Known | Event handler |
| 0x00752E3A | `HandleHighlightCity` | Known | Event handler |
| 0x00752E73 | `HandleSelectCity` | Known | Event handler |
| 0x00752EDF | `HandleHighlightCity` | Known | Event handler |
| 0x00752F18 | `HandleSelectCity` | Known | Event handler |
| 0x00752F84 | `HandleHighlightCity` | Known | Event handler |
| 0x00752FBD | `HandleSelectCity` | Known | Event handler |
| 0x00753029 | `HandleHighlightCity` | Known | Event handler |
| 0x00753062 | `HandleSelectCity` | Known | Event handler |
| 0x007530CE | `HandleHighlightCity` | Known | Event handler |
| 0x00753107 | `HandleSelectCity` | Known | Event handler |
| 0x00753173 | `HandleHighlightCity` | Known | Event handler |
| 0x007531AC | `HandleSelectCity` | Known | Event handler |
| 0x00753218 | `HandleHighlightCity` | Known | Event handler |
| 0x00753251 | `HandleSelectCity` | Known | Event handler |
| 0x007532BD | `HandleHighlightCity` | Known | Event handler |
| 0x007532F6 | `HandleSelectCity` | Known | Event handler |
| 0x00753362 | `HandleHighlightCity` | Known | Event handler |
| 0x0075339B | `HandleSelectCity` | Known | Event handler |
| 0x00753407 | `HandleHighlightCity` | Known | Event handler |
| 0x00753440 | `HandleSelectCity` | Known | Event handler |
| 0x007534AC | `HandleHighlightCity` | Known | Event handler |
| 0x007534E5 | `HandleSelectCity` | Known | Event handler |
| 0x00753551 | `HandleHighlightCity` | Known | Event handler |
| 0x0075358A | `HandleSelectCity` | Known | Event handler |
| 0x007535F6 | `HandleHighlightCity` | Known | Event handler |
| 0x00753AEE | `HandleMusicSelected` | Known | Event handler |
| 0x00753B30 | `HandleMusicHilited` | Known | Event handler |
| 0x00753B68 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00753BAE | `HandleMusicHilited` | Known | Event handler |
| 0x00753BE6 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00753C2C | `HandlePlaylistsHilited` | Known | Event handler |
| 0x00753C68 | `HandleArtistsSelected` | Known | Event handler |
| 0x00753CAC | `HandleArtistsHilited` | Known | Event handler |
| 0x00753CE6 | `HandleAlbumsSelected` | Known | Event handler |
| 0x00753D29 | `HandleAlbumsHilited` | Known | Event handler |
| 0x00753D62 | `HandleCompilationsSelected` | Known | Event handler |
| 0x00753DAB | `HandleCompilationsHilited` | Known | Event handler |
| 0x00753DEA | `HandleSongsSelected` | Known | Event handler |
| 0x00753E2C | `HandleSongsHilited` | Known | Event handler |
| 0x00753E64 | `HandleGenresSelected` | Known | Event handler |
| 0x00753EA7 | `HandleGenresHilited` | Known | Event handler |
| 0x00753EE0 | `HandleComposersSelected` | Known | Event handler |
| 0x00753F26 | `HandleComposersHilited` | Known | Event handler |
| 0x00753F62 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x00753FA9 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00754068 | `HandleMusicHilited` | Known | Event handler |
| 0x007540A0 | `HandleVideosSelected` | Known | Event handler |
| 0x007540E3 | `HandleVideosHilited` | Known | Event handler |
| 0x0075411C | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00754167 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x007541A8 | `HandleMoviesSelected` | Known | Event handler |
| 0x007541EB | `HandleMoviesHilited` | Known | Event handler |
| 0x00754224 | `HandleTVShowsSelected` | Known | Event handler |
| 0x00754268 | `HandleTVShowsHilited` | Known | Event handler |
| 0x007542A2 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x007542EA | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00754328 | `HandleRentalsSelected` | Known | Event handler |
| 0x0075436C | `HandleRentalsHilited` | Known | Event handler |
| 0x007543A6 | `HandlePhotosSelected` | Known | Event handler |
| 0x007543E9 | `HandlePhotosHilited` | Known | Event handler |
| 0x00754422 | `HandlePhotosSelected` | Known | Event handler |
| 0x00754465 | `HandlePhotosHilited` | Known | Event handler |
| 0x0075449E | `HandlePodcastsSelected` | Known | Event handler |
| 0x007544E3 | `HandlePodcastsHilited` | Known | Event handler |
| 0x00754596 | `HandleGenericHilited` | Known | Event handler |
| 0x0075468F | `HandleGenericHilited` | Known | Event handler |
| 0x00754B74 | `HandleLock` | Known | Event handler |
| 0x00754CE5 | `HandleNikePlusSelected` | Known | Event handler |
| 0x00754D2A | `HandleGenericHilited` | Known | Event handler |
| 0x00754E30 | `HandleGenericHilited` | Known | Event handler |
| 0x00754F2F | `HandleGenericHilited` | Known | Event handler |
| 0x0075501C | `HandleGenericHilited` | Known | Event handler |
| 0x00755119 | `HandleGenericHilited` | Known | Event handler |
| 0x00755193 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x007551DC | `HandleGenericHilited` | Known | Event handler |
| 0x00755255 | `HandleBacklightSelected` | Known | Event handler |
| 0x0075529B | `HandleGenericHilited` | Known | Event handler |
| 0x00755316 | `HandleSleepSelected` | Known | Event handler |
| 0x00755358 | `HandleGenericHilited` | Known | Event handler |
| 0x007553CF | `HandleNowPlaying` | Known | Event handler |
| 0x00755447 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x0075548A | `HandleCoverFlowSelected` | Known | Event handler |
| 0x007554D0 | `HandleMusicHilited` | Known | Event handler |
| 0x00755508 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x0075554E | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x0075558C | `HandleArtistsSelected` | Known | Event handler |
| 0x007555D0 | `HandleArtistsHilited` | Known | Event handler |
| 0x0075560A | `HandleAlbumsSelected` | Known | Event handler |
| 0x0075564D | `HandleAlbumsHilited` | Known | Event handler |
| 0x00755686 | `HandleCompilationsSelected` | Known | Event handler |
| 0x007556CF | `HandleCompilationsHilited` | Known | Event handler |
| 0x0075570E | `HandleSongsSelected` | Known | Event handler |
| 0x00755750 | `HandleSongsHilited` | Known | Event handler |
| 0x007557FB | `HandleGenericHilited` | Known | Event handler |
| 0x00755873 | `HandleGenresSelected` | Known | Event handler |
| 0x007558B6 | `HandleGenresHilited` | Known | Event handler |
| 0x007558EF | `HandleComposersSelected` | Known | Event handler |
| 0x00755935 | `HandleComposersHilited` | Known | Event handler |
| 0x00755971 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x007559B8 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00755A77 | `HandleMusicHilited` | Known | Event handler |
| 0x00755AED | `HandlePlayPause` | Known | Event handler |
| 0x00755B22 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x00755C0C | `HandleSelect` | Known | Event handler |
| 0x00755C52 | `HandleMoviesSelected` | Known | Event handler |
| 0x00755C95 | `HandleMoviesHilited` | Known | Event handler |
| 0x00755CCE | `HandleRentalsSelected` | Known | Event handler |
| 0x00755D12 | `HandleRentalsHilited` | Known | Event handler |
| 0x00755D4C | `HandleTVShowsSelected` | Known | Event handler |
| 0x00755D90 | `HandleTVShowsHilited` | Known | Event handler |
| 0x00755DCA | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00755E12 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00755E50 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00755E9B | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00755F61 | `HandleVideosHilited` | Known | Event handler |
| 0x007565A3 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x0075712A | `HandleMainMenu` | Known | Event handler |
| 0x00757162 | `HandleMusicMenu` | Known | Event handler |
| 0x0075768A | `HandleRadioRegion` | Known | Event handler |
| 0x0075772E | `HandleLanguage` | Known | Event handler |
| 0x00757834 | `HandleNew` | Known | Event handler |
| 0x007578AF | `HandleClear` | Known | Event handler |
| 0x007578E0 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x0075799C | `HandleSelectIndexedSession` | Known | Event handler |
| 0x00757B63 | `HandleBasicSelected` | Known | Event handler |
| 0x00757C09 | `HandleTimedWorkoutSelected` | Known | Event handler |
| 0x00757CB6 | `HandleDistanceWorkoutSelected` | Known | Event handler |
| 0x00757D66 | `HandleCaloriesWorkoutSelected` | Known | Event handler |
| 0x00758151 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x007581A4 | `HandleSelect` | Known | Event handler |
| 0x007582CE | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x00758308 | `HandleEQSettingSelected` | Known | Event handler |
| 0x00758340 | `HandleEQSettingSelected` | Known | Event handler |
| 0x0076D4A4 | `HandleItemSelected` | Known | Event handler |
| 0x0076D5EF | `HandleNextContact` | Known | Event handler |
| 0x0076D61B | `HandlePreviousContact` | Known | Event handler |
| 0x0076D651 | `HandleSelectKey` | Known | Event handler |
| 0x0076DC62 | `HandleSelect` | Known | Event handler |
| 0x0076DF89 | `HandleDateChosen` | Known | Event handler |
| 0x0076DFBF | `HandleTimeChosen` | Known | Event handler |
| 0x0076DFF5 | `HandleFrequencyChosen` | Known | Event handler |
| 0x0076E030 | `HandleSoundChosen` | Known | Event handler |
| 0x0076E067 | `HandleLabelChosen` | Known | Event handler |
| 0x0076E09E | `HandleDeleteChosen` | Known | Event handler |
| 0x0076E0DA | `HandleSelect` | Known | Event handler |
| 0x0076E112 | `HandleSelect` | Known | Event handler |
| 0x0076E453 | `HandleLeaveAlarm` | Known | Event handler |
| 0x0076E480 | `HandleLeaveAlarm` | Known | Event handler |
| 0x0076E4AF | `HandleLeaveAlarm` | Known | Event handler |
| 0x0076E4DC | `HandleLeaveAlarm` | Known | Event handler |
| 0x0076E616 | `HandleSelect` | Known | Event handler |
| 0x0076E644 | `HandleSelect` | Known | Event handler |
| 0x0076E7A3 | `HandleNextDay` | Known | Event handler |
| 0x0076E7CB | `HandlePreviousDay` | Known | Event handler |
| 0x0076E97A | `HandleSelect` | Known | Event handler |
| 0x0076E9A7 | `HandleNextDay` | Known | Event handler |
| 0x0076E9CF | `HandlePreviousDay` | Known | Event handler |
| 0x0076EB77 | `HandleNextDay` | Known | Event handler |
| 0x0076EB9F | `HandlePreviousDay` | Known | Event handler |
| 0x0076EC60 | `HandleSelect` | Known | Event handler |
| 0x0076EC8B | `HandleNextDay` | Known | Event handler |
| 0x0076ECB3 | `HandlePreviousDay` | Known | Event handler |
| 0x0076EE2A | `HandleSelectLozinch` | Known | Event handler |
| 0x0076EFA2 | `HandleSelectLozinch` | Known | Event handler |
| 0x0076F0C1 | `HandleFlowNext` | Known | Event handler |
| 0x0076F0EF | `HandlePlayPause` | Known | Event handler |
| 0x0076F13E | `HandleFlowPrev` | Known | Event handler |
| 0x0076F169 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x0076F25D | `HandleAlbumSelected` | Known | Event handler |
| 0x0076F3F8 | `HandleFlowNext` | Known | Event handler |
| 0x0076F446 | `HandleFlowNext` | Known | Event handler |
| 0x0076F474 | `HandlePlayPause` | Known | Event handler |
| 0x0076F4C3 | `HandleFlowPrev` | Known | Event handler |
| 0x0076F4EF | `HandleFlowPrev` | Known | Event handler |
| 0x0076F50F | `HandleFlowWheel` | Known | Event handler |
| 0x0076F89F | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0076FCCA | `HandleArrowDown` | Known | Event handler |
| 0x0076FD34 | `HandleArrowUp` | Known | Event handler |
| 0x0076FD53 | `HandleWheel` | Known | Event handler |
| 0x0076FDDC | `HandleSelect` | Known | Event handler |
| 0x0076FE59 | `HandleGameHilited` | Known | Event handler |
| 0x007732BF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00774EFB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00776B37 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00778773 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077A3AF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077BFEB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077DC27 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077F863 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078149F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007830DB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00784D17 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00786953 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078858F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078A1CB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078BE07 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078DA43 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078F67F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007912BB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00792EF7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00794B33 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079676F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007983AB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00799FE7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079BC23 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079D85F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079F49B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A10D7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A2D13 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A494F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A658B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A81C7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007A9E03 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007ABA3F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007AD67B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007AF2B7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B0EF3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B2B2F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B4750 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B52D8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B5E60 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B69E8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B7570 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B80F8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B8C80 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007B9808 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BA390 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BAF18 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BBAA0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BC628 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007BD1B0 | `HandlePlayPause` | Known | Event handler |
| 0x007BD1E6 | `HandleAddToOTG` | Known | Event handler |
| 0x007BD383 | `HandlePlayPause` | Known | Event handler |
| 0x007BD3AA | `HandleSelect` | Known | Event handler |
| 0x007BD3D7 | `HandleHilite` | Known | Event handler |
| 0x007BD408 | `HandlePlayPause` | Known | Event handler |
| 0x007BD49B | `HandlePlayPause` | Known | Event handler |
| 0x007BD4C2 | `HandleSelect` | Known | Event handler |
| 0x007BD528 | `HandleHilite` | Known | Event handler |
| 0x007BD55A | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x007BD5A4 | `HandlePlayPause` | Known | Event handler |
| 0x007BD5DA | `HandleAddToOTG` | Known | Event handler |
| 0x007BD66C | `HandlePlayPause` | Known | Event handler |
| 0x007BD693 | `HandleSelect` | Known | Event handler |
| 0x007BD6FC | `HandlePlayPause` | Known | Event handler |
| 0x007BD732 | `HandleAddToOTG` | Known | Event handler |
| 0x007BD7C4 | `HandlePlayPause` | Known | Event handler |
| 0x007BD7EB | `HandleSelect` | Known | Event handler |
| 0x007BD854 | `HandlePlayPause` | Known | Event handler |
| 0x007BD8DA | `HandleSelect` | Known | Event handler |
| 0x007BD93F | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007BD980 | `HandlePlayPause` | Known | Event handler |
| 0x007BD9B6 | `HandleAddToOTG` | Known | Event handler |
| 0x007BDBE8 | `HandlePlayPause` | Known | Event handler |
| 0x007BDC0F | `HandleSelect` | Known | Event handler |
| 0x007BDC3C | `HandleHilite` | Known | Event handler |
| 0x007BDC6C | `HandlePlayPause` | Known | Event handler |
| 0x007BDCA2 | `HandleAddToOTG` | Known | Event handler |
| 0x007BDED4 | `HandlePlayPause` | Known | Event handler |
| 0x007BDEFB | `HandleSelect` | Known | Event handler |
| 0x007BDF28 | `HandleHilite` | Known | Event handler |
| 0x007BDF58 | `HandlePlayPause` | Known | Event handler |
| 0x007BDF8E | `HandleAddToOTG` | Known | Event handler |
| 0x007BE279 | `HandlePlayPause` | Known | Event handler |
| 0x007BE2A0 | `HandleSelect` | Known | Event handler |
| 0x007BE2D0 | `HandlePlayPause` | Known | Event handler |
| 0x007BE306 | `HandleAddToOTG` | Known | Event handler |
| 0x007BE398 | `HandlePlayPause` | Known | Event handler |
| 0x007BE3BF | `HandleSelect` | Known | Event handler |
| 0x007BE450 | `HandlePlayPause` | Known | Event handler |
| 0x007BE486 | `HandleAddToOTG` | Known | Event handler |
| 0x007BE63F | `HandlePlayPause` | Known | Event handler |
| 0x007BE666 | `HandleSelect` | Known | Event handler |
| 0x007BE698 | `HandlePlayPause` | Known | Event handler |
| 0x007BE6CE | `HandleAddToOTG` | Known | Event handler |
| 0x007BE753 | `HandleSelect` | Known | Event handler |
| 0x007BE7EC | `HandleHilite` | Known | Event handler |
| 0x007BE818 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007BE85C | `HandlePlayPause` | Known | Event handler |
| 0x007BE892 | `HandleAddToOTG` | Known | Event handler |
| 0x007BE917 | `HandleSelect` | Known | Event handler |
| 0x007BE97C | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007BE9C0 | `HandlePlayPause` | Known | Event handler |
| 0x007BEB64 | `HandleSelect` | Known | Event handler |
| 0x007BEB91 | `HandleHilite` | Known | Event handler |
| 0x007BEBBD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007BEC00 | `HandlePlayPause` | Known | Event handler |
| 0x007BEC86 | `HandleSelect` | Known | Event handler |
| 0x007BED14 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007BED58 | `HandlePlayPause` | Known | Event handler |
| 0x007BEDDE | `HandleSelect` | Known | Event handler |
| 0x007BEE43 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007BEE84 | `HandlePlayPause` | Known | Event handler |
| 0x007BEF0A | `HandleSelect` | Known | Event handler |
| 0x007BEF70 | `HandleHilite` | Known | Event handler |
| 0x007BEF9C | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007BEFE0 | `HandlePlayPause` | Known | Event handler |
| 0x007BF016 | `HandleAddToOTG` | Known | Event handler |
| 0x007BF1D9 | `HandlePlayPause` | Known | Event handler |
| 0x007BF200 | `HandleSelect` | Known | Event handler |
| 0x007BF230 | `HandlePlayPause` | Known | Event handler |
| 0x007BF266 | `HandleAddToOTG` | Known | Event handler |
| 0x007BF487 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x007BF5A0 | `HandlePlayPause` | Known | Event handler |
| 0x007BF6CD | `HandleSelect` | Known | Event handler |
| 0x007BF6F9 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007BF73C | `HandlePlayPause` | Known | Event handler |
| 0x007BF7C2 | `HandleSelect` | Known | Event handler |
| 0x007BF7EF | `HandleHilite` | Known | Event handler |
| 0x007BF81B | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007BF85C | `HandlePlayPause` | Known | Event handler |
| 0x007BF98F | `HandleSelect` | Known | Event handler |
| 0x007BF9BB | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C02CD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C0B85 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C143D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C1CF5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C25AD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C2E65 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C371D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C3FD5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007C401E | `HandleTVOutChanged` | Known | Event handler |
| 0x007C4056 | `HandleTVSignalChanged` | Known | Event handler |
| 0x007C4091 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x007C40E2 | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x007C4127 | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x007C4170 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x007C41B2 | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x007C41F5 | `HandleSelect` | Known | Event handler |
| 0x007C4225 | `HandleSelect` | Known | Event handler |
| 0x007C425D | `HandleMenuLongpress` | Known | Event handler |
| 0x007C428B | `HandleMenuKey` | Known | Event handler |
| 0x007C4311 | `HandlePlayPause` | Known | Event handler |
| 0x007C4391 | `HandleSelect` | Known | Event handler |
| 0x007C4C9E | `HandlePlayPause` | Known | Event handler |
| 0x007C4D13 | `HandleWheelProgress` | Known | Event handler |
| 0x007C4D51 | `HandleMenuLongpress` | Known | Event handler |
| 0x007C4D7F | `HandleMenuKey` | Known | Event handler |
| 0x007C4E05 | `HandlePlayPause` | Known | Event handler |
| 0x007C4E85 | `HandleSelectProgress` | Known | Event handler |
| 0x007C579A | `HandlePlayPause` | Known | Event handler |
| 0x007C580F | `HandleWheelProgress` | Known | Event handler |
| 0x007C584D | `HandleMenuLongpress` | Known | Event handler |
| 0x007C587B | `HandleMenuKey` | Known | Event handler |
| 0x007C5901 | `HandlePlayPause` | Known | Event handler |
| 0x007C5981 | `HandleSelectVolume` | Known | Event handler |
| 0x007C6294 | `HandlePlayPause` | Known | Event handler |
| 0x007C6309 | `HandleWheelVolume` | Known | Event handler |
| 0x007C6345 | `HandleMenuLongpress` | Known | Event handler |
| 0x007C6373 | `HandleMenuKey` | Known | Event handler |
| 0x007C63F9 | `HandlePlayPause` | Known | Event handler |
| 0x007C6479 | `HandleSelectRating` | Known | Event handler |
| 0x007C6D8C | `HandlePlayPause` | Known | Event handler |
| 0x007C6E01 | `HandleWheelRating` | Known | Event handler |
| 0x007C6E3D | `HandleMenuLongpress` | Known | Event handler |
| 0x007C6E6B | `HandleMenuKey` | Known | Event handler |
| 0x007C6EE3 | `HandlePlayPause` | Known | Event handler |
| 0x007C6F5A | `HandleSelectScrub` | Known | Event handler |
| 0x007C785E | `HandlePlayPause` | Known | Event handler |
| 0x007C78CA | `HandleWheelScrub` | Known | Event handler |
| 0x007C7905 | `HandleMenuLongpress` | Known | Event handler |
| 0x007C7933 | `HandleMenuKey` | Known | Event handler |
| 0x007C7990 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007C79C8 | `HandlePlayPause` | Known | Event handler |
| 0x007C7A22 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007C7A57 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x007C8371 | `HandlePlayPause` | Known | Event handler |
| 0x007C83E6 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007C8429 | `HandleMenuLongpress` | Known | Event handler |
| 0x007C8457 | `HandleMenuKey` | Known | Event handler |
| 0x007C84DD | `HandlePlayPause` | Known | Event handler |
| 0x007C855D | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007C8E73 | `HandlePlayPause` | Known | Event handler |
| 0x007C8F11 | `HandleMenuLongpress` | Known | Event handler |
| 0x007C8F3F | `HandleMenuKey` | Known | Event handler |
| 0x007C8FC5 | `HandlePlayPause` | Known | Event handler |
| 0x007C9045 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007C995B | `HandlePlayPause` | Known | Event handler |
| 0x007C99F9 | `HandleMenuLongpress` | Known | Event handler |
| 0x007C9A27 | `HandleMenuKey` | Known | Event handler |
| 0x007C9AAD | `HandlePlayPause` | Known | Event handler |
| 0x007C9B2D | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007CA443 | `HandlePlayPause` | Known | Event handler |
| 0x007CA4E1 | `HandleMenuLongpress` | Known | Event handler |
| 0x007CA50F | `HandleMenuKey` | Known | Event handler |
| 0x007CA595 | `HandlePlayPause` | Known | Event handler |
| 0x007CA615 | `HandleSelectChapterArt` | Known | Event handler |
| 0x007CAF2C | `HandlePlayPause` | Known | Event handler |
| 0x007CAFA1 | `HandleWheelVolume` | Known | Event handler |
| 0x007CAFDD | `HandleMenuLongpress` | Known | Event handler |
| 0x007CB00B | `HandleMenuKey` | Known | Event handler |
| 0x007CB09A | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007CB131 | `HandleSelect` | Known | Event handler |
| 0x007CBA47 | `HandlePlayPause` | Known | Event handler |
| 0x007CBAC5 | `HandleWheel` | Known | Event handler |
| 0x007CBAF9 | `HandleMenuLongpress` | Known | Event handler |
| 0x007CBB27 | `HandleMenuKey` | Known | Event handler |
| 0x007CBBB6 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007CBC4D | `HandleSelect` | Known | Event handler |
| 0x007CC563 | `HandlePlayPause` | Known | Event handler |
| 0x007CC5E1 | `HandleWheel` | Known | Event handler |
| 0x007CC615 | `HandleMenuLongpress` | Known | Event handler |
| 0x007CC643 | `HandleMenuKey` | Known | Event handler |
| 0x007CC6C9 | `HandlePlayPause` | Known | Event handler |
| 0x007CC749 | `HandleSelect` | Known | Event handler |
| 0x007CD056 | `HandlePlayPause` | Known | Event handler |
| 0x007CD0CB | `HandleWheel` | Known | Event handler |
| 0x007CD101 | `HandleMenuLongpress` | Known | Event handler |
| 0x007CD12F | `HandleMenuKey` | Known | Event handler |
| 0x007CD1B5 | `HandlePlayPause` | Known | Event handler |
| 0x007CD235 | `HandleSelectProgress` | Known | Event handler |
| 0x007CDB4A | `HandlePlayPause` | Known | Event handler |
| 0x007CDBBF | `HandleWheelProgress` | Known | Event handler |
| 0x007CDBFD | `HandleMenuLongpress` | Known | Event handler |
| 0x007CDC2B | `HandleMenuKey` | Known | Event handler |
| 0x007CDCA3 | `HandlePlayPause` | Known | Event handler |
| 0x007CDD1A | `HandleSelectScrub` | Known | Event handler |
| 0x007CE61E | `HandlePlayPause` | Known | Event handler |
| 0x007CE68A | `HandleWheelScrub` | Known | Event handler |
| 0x007CE6C5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007CE6F3 | `HandleMenuKey` | Known | Event handler |
| 0x007CE779 | `HandlePlayPause` | Known | Event handler |
| 0x007CF105 | `HandlePlayPause` | Known | Event handler |
| 0x007CF17A | `HandleWheelVolume` | Known | Event handler |
| 0x007CF1B5 | `HandleMenuLongpress` | Known | Event handler |
| 0x007CF1E3 | `HandleMenuKey` | Known | Event handler |
| 0x007CF269 | `HandlePlayPause` | Known | Event handler |
| 0x007CFBF5 | `HandlePlayPause` | Known | Event handler |
| 0x007CFC6A | `HandleWheelBrightness` | Known | Event handler |
| 0x007CFD81 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007D06D4 | `HandleWheel` | Known | Event handler |
| 0x007D0709 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D0737 | `HandleMenuKey` | Known | Event handler |
| 0x007D07BD | `HandlePlayPause` | Known | Event handler |
| 0x007D083D | `HandleSelect` | Known | Event handler |
| 0x007D0CDF | `HandlePlayPause` | Known | Event handler |
| 0x007D0D6D | `HandleMenuLongpress` | Known | Event handler |
| 0x007D0D9B | `HandleMenuKey` | Known | Event handler |
| 0x007D0E21 | `HandlePlayPause` | Known | Event handler |
| 0x007D0EA1 | `HandleSelectProgress` | Known | Event handler |
| 0x007D134B | `HandlePlayPause` | Known | Event handler |
| 0x007D13C0 | `HandleWheelProgress` | Known | Event handler |
| 0x007D13FD | `HandleMenuLongpress` | Known | Event handler |
| 0x007D142B | `HandleMenuKey` | Known | Event handler |
| 0x007D14B1 | `HandlePlayPause` | Known | Event handler |
| 0x007D1531 | `HandleSelectProgress` | Known | Event handler |
| 0x007D19DB | `HandlePlayPause` | Known | Event handler |
| 0x007D1A50 | `HandleWheelProgress` | Known | Event handler |
| 0x007D1A8D | `HandleMenuLongpress` | Known | Event handler |
| 0x007D1ABB | `HandleMenuKey` | Known | Event handler |
| 0x007D1B41 | `HandlePlayPause` | Known | Event handler |
| 0x007D1BC1 | `HandleSelectProgress` | Known | Event handler |
| 0x007D1FF7 | `HandlePlayPause` | Known | Event handler |
| 0x007D206C | `HandleWheelProgress` | Known | Event handler |
| 0x007D20A9 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D20D7 | `HandleMenuKey` | Known | Event handler |
| 0x007D2144 | `HandlePlayPause` | Known | Event handler |
| 0x007D21B0 | `HandleSelectScrub` | Known | Event handler |
| 0x007D25CA | `HandlePlayPause` | Known | Event handler |
| 0x007D262B | `HandleWheelScrub` | Known | Event handler |
| 0x007D2665 | `HandleMenuLongpress` | Known | Event handler |
| 0x007D2693 | `HandleMenuKey` | Known | Event handler |
| 0x007D2719 | `HandlePlayPause` | Known | Event handler |
| 0x007D2799 | `HandleSelectVolume` | Known | Event handler |
| 0x007D2BCD | `HandlePlayPause` | Known | Event handler |
| 0x007D2C42 | `HandleWheelVolume` | Known | Event handler |
| 0x007D2D55 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007D31F4 | `HandleSelect` | Known | Event handler |
| 0x007D3221 | `HandleSelect` | Known | Event handler |
| 0x007D3251 | `HandleSelect` | Known | Event handler |
| 0x007D3281 | `HandleSelect` | Known | Event handler |
| 0x007D32B1 | `HandleSelect` | Known | Event handler |
| 0x007D32E1 | `HandleSelect` | Known | Event handler |
| 0x007D3311 | `HandleSelect` | Known | Event handler |
| 0x007D3341 | `HandleSelect` | Known | Event handler |
| 0x007D3371 | `HandleSelect` | Known | Event handler |
| 0x007D33E1 | `HandleSelect` | Known | Event handler |
| 0x007D3411 | `HandleSelect` | Known | Event handler |
| 0x007D3489 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007D34BC | `HandleNotesPop` | Known | Event handler |
| 0x007D3539 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007D356C | `HandleNotesPop` | Known | Event handler |
| 0x007D3A28 | `HandleNotesSelected` | Known | Event handler |
| 0x007D3A65 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007D3A98 | `HandleNotesPop` | Known | Event handler |
| 0x007D3F54 | `HandleNotesSelected` | Known | Event handler |
| 0x007D3F91 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007D3FC4 | `HandleNotesPop` | Known | Event handler |
| 0x007D3FEF | `HandleNotesSelected` | Known | Event handler |
| 0x007D44C1 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007D44F4 | `HandleNotesPop` | Known | Event handler |
| 0x007D451F | `HandleNotesSelected` | Known | Event handler |
| 0x007D49F1 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007D4A24 | `HandleNotesPop` | Known | Event handler |
| 0x007D4AA1 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007D4AD4 | `HandleNotesPop` | Known | Event handler |
| 0x007D4B51 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007D4B84 | `HandleNotesPop` | Known | Event handler |
| 0x007D4BFC | `HandlePlayPause` | Known | Event handler |
| 0x007D4C25 | `HandlePlayPause` | Known | Event handler |
| 0x007D4C53 | `HandlePlayPause` | Known | Event handler |
| 0x007D4C88 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007D4D08 | `HandleHiliteAlbum` | Known | Event handler |
| 0x007D4DB1 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007D4E38 | `HandleHiliteAlbum` | Known | Event handler |
| 0x007D50FC | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x007D5158 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x007D530F | `HandleSelect` | Known | Event handler |
| 0x007D5493 | `HandleSelect` | Known | Event handler |
| 0x007D54CD | `HandleImageLast` | Known | Event handler |
| 0x007D54F7 | `HandleImageNext` | Known | Event handler |
| 0x007D5526 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D5560 | `HandleImageFirst` | Known | Event handler |
| 0x007D558B | `HandleImagePrev` | Known | Event handler |
| 0x007D55B7 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D55E6 | `HandleImageNext` | Known | Event handler |
| 0x007D560F | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D5643 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D5672 | `HandleImagePrev` | Known | Event handler |
| 0x007D5693 | `HandleImageWheel` | Known | Event handler |
| 0x007D5731 | `HandleImageNext` | Known | Event handler |
| 0x007D5760 | `HandlePlayPause` | Known | Event handler |
| 0x007D57AF | `HandleImagePrev` | Known | Event handler |
| 0x007D57DB | `HandleSelect` | Known | Event handler |
| 0x007D5AAB | `HandleImageNext` | Known | Event handler |
| 0x007D5AD5 | `HandlePause` | Known | Event handler |
| 0x007D5AFA | `HandlePlay` | Known | Event handler |
| 0x007D5B23 | `HandlePlayPause` | Known | Event handler |
| 0x007D5B4C | `HandleImagePrev` | Known | Event handler |
| 0x007D5BA5 | `HandleWheel` | Known | Event handler |
| 0x007D5C3D | `HandleImageNext` | Known | Event handler |
| 0x007D5C6C | `HandlePlayPause` | Known | Event handler |
| 0x007D5CBB | `HandleImagePrev` | Known | Event handler |
| 0x007D5CE7 | `HandleSelect` | Known | Event handler |
| 0x007D5FB7 | `HandleImageNext` | Known | Event handler |
| 0x007D5FE1 | `HandlePause` | Known | Event handler |
| 0x007D6006 | `HandlePlay` | Known | Event handler |
| 0x007D602F | `HandlePlayPause` | Known | Event handler |
| 0x007D6058 | `HandleImagePrev` | Known | Event handler |
| 0x007D60B1 | `HandleWheel` | Known | Event handler |
| 0x007D6149 | `HandleImageNext` | Known | Event handler |
| 0x007D6178 | `HandlePlayPause` | Known | Event handler |
| 0x007D61C7 | `HandleImagePrev` | Known | Event handler |
| 0x007D61F3 | `HandleSelect` | Known | Event handler |
| 0x007D64C3 | `HandleImageNext` | Known | Event handler |
| 0x007D64ED | `HandlePause` | Known | Event handler |
| 0x007D6512 | `HandlePlay` | Known | Event handler |
| 0x007D653B | `HandlePlayPause` | Known | Event handler |
| 0x007D6564 | `HandleImagePrev` | Known | Event handler |
| 0x007D65BD | `HandleWheel` | Known | Event handler |
| 0x007D6655 | `HandleImageNext` | Known | Event handler |
| 0x007D6684 | `HandlePlayPause` | Known | Event handler |
| 0x007D66D3 | `HandleImagePrev` | Known | Event handler |
| 0x007D66FF | `HandleSelect` | Known | Event handler |
| 0x007D69CF | `HandleImageNext` | Known | Event handler |
| 0x007D69F9 | `HandlePause` | Known | Event handler |
| 0x007D6A1E | `HandlePlay` | Known | Event handler |
| 0x007D6A47 | `HandlePlayPause` | Known | Event handler |
| 0x007D6A70 | `HandleImagePrev` | Known | Event handler |
| 0x007D6AC9 | `HandleWheel` | Known | Event handler |
| 0x007D6B61 | `HandleImageNext` | Known | Event handler |
| 0x007D6B90 | `HandlePlayPause` | Known | Event handler |
| 0x007D6BDF | `HandleImagePrev` | Known | Event handler |
| 0x007D6C0B | `HandleSelect` | Known | Event handler |
| 0x007D6EDB | `HandleImageNext` | Known | Event handler |
| 0x007D6F05 | `HandlePause` | Known | Event handler |
| 0x007D6F2A | `HandlePlay` | Known | Event handler |
| 0x007D6F53 | `HandlePlayPause` | Known | Event handler |
| 0x007D6F7C | `HandleImagePrev` | Known | Event handler |
| 0x007D6FD5 | `HandleWheel` | Known | Event handler |
| 0x007D706D | `HandleImageNext` | Known | Event handler |
| 0x007D709C | `HandlePlayPause` | Known | Event handler |
| 0x007D70EB | `HandleImagePrev` | Known | Event handler |
| 0x007D7117 | `HandleSelect` | Known | Event handler |
| 0x007D73E7 | `HandleImageNext` | Known | Event handler |
| 0x007D7411 | `HandlePause` | Known | Event handler |
| 0x007D7436 | `HandlePlay` | Known | Event handler |
| 0x007D745F | `HandlePlayPause` | Known | Event handler |
| 0x007D7488 | `HandleImagePrev` | Known | Event handler |
| 0x007D74E1 | `HandleWheel` | Known | Event handler |
| 0x007D7579 | `HandleImageNext` | Known | Event handler |
| 0x007D75A8 | `HandlePlayPause` | Known | Event handler |
| 0x007D75F7 | `HandleImagePrev` | Known | Event handler |
| 0x007D7623 | `HandleSelect` | Known | Event handler |
| 0x007D786E | `HandleImageNext` | Known | Event handler |
| 0x007D7898 | `HandlePause` | Known | Event handler |
| 0x007D78BD | `HandlePlay` | Known | Event handler |
| 0x007D78E6 | `HandlePlayPause` | Known | Event handler |
| 0x007D790F | `HandleImagePrev` | Known | Event handler |
| 0x007D7978 | `HandleWheel` | Known | Event handler |
| 0x007D7A11 | `HandleImageNext` | Known | Event handler |
| 0x007D7A40 | `HandlePlayPause` | Known | Event handler |
| 0x007D7A8F | `HandleImagePrev` | Known | Event handler |
| 0x007D7ABB | `HandleSelect` | Known | Event handler |
| 0x007D7D06 | `HandleImageNext` | Known | Event handler |
| 0x007D7D30 | `HandlePause` | Known | Event handler |
| 0x007D7D55 | `HandlePlay` | Known | Event handler |
| 0x007D7D7E | `HandlePlayPause` | Known | Event handler |
| 0x007D7DA7 | `HandleImagePrev` | Known | Event handler |
| 0x007D7E10 | `HandleWheel` | Known | Event handler |
| 0x007D7EA9 | `HandleImageNext` | Known | Event handler |
| 0x007D7ED8 | `HandlePlayPause` | Known | Event handler |
| 0x007D7F27 | `HandleImagePrev` | Known | Event handler |
| 0x007D7F53 | `HandleSelect` | Known | Event handler |
| 0x007D819E | `HandleImageNext` | Known | Event handler |
| 0x007D81C8 | `HandlePause` | Known | Event handler |
| 0x007D81ED | `HandlePlay` | Known | Event handler |
| 0x007D8216 | `HandlePlayPause` | Known | Event handler |
| 0x007D823F | `HandleImagePrev` | Known | Event handler |
| 0x007D82A8 | `HandleWheel` | Known | Event handler |
| 0x007D8341 | `HandleImageNext` | Known | Event handler |
| 0x007D8370 | `HandlePlayPause` | Known | Event handler |
| 0x007D83BF | `HandleImagePrev` | Known | Event handler |
| 0x007D83EB | `HandleSelect` | Known | Event handler |
| 0x007D8636 | `HandleImageNext` | Known | Event handler |
| 0x007D8660 | `HandlePause` | Known | Event handler |
| 0x007D8685 | `HandlePlay` | Known | Event handler |
| 0x007D86AE | `HandlePlayPause` | Known | Event handler |
| 0x007D86D7 | `HandleImagePrev` | Known | Event handler |
| 0x007D8740 | `HandleWheel` | Known | Event handler |
| 0x007D87D9 | `HandleImageNext` | Known | Event handler |
| 0x007D8808 | `HandlePlayPause` | Known | Event handler |
| 0x007D8857 | `HandleImagePrev` | Known | Event handler |
| 0x007D8883 | `HandleSelect` | Known | Event handler |
| 0x007D8ACE | `HandleImageNext` | Known | Event handler |
| 0x007D8AF8 | `HandlePause` | Known | Event handler |
| 0x007D8B1D | `HandlePlay` | Known | Event handler |
| 0x007D8B46 | `HandlePlayPause` | Known | Event handler |
| 0x007D8B6F | `HandleImagePrev` | Known | Event handler |
| 0x007D8BD8 | `HandleWheel` | Known | Event handler |
| 0x007D8C05 | `HandleSelect` | Known | Event handler |
| 0x007D8C35 | `HandleSelect` | Known | Event handler |
| 0x007D8D58 | `HandleTuning` | Known | Event handler |
| 0x007D8F14 | `HandleVolumeChange` | Known | Event handler |
| 0x007D9060 | `HandleVolumeWheel` | Known | Event handler |
| 0x007D91BB | `HandleTuningSelect` | Known | Event handler |
| 0x007D949A | `HandleFrequencyChange` | Known | Event handler |
| 0x007D95F7 | `HandleTuningSelect` | Known | Event handler |
| 0x007D98D6 | `HandleFrequencyChange` | Known | Event handler |
| 0x007D9A00 | `HandleTimerDone` | Known | Event handler |
| 0x007D9BF5 | `HandleVolumeChange` | Known | Event handler |
| 0x007D9D0C | `HandleVolumeWheel` | Known | Event handler |
| 0x007DA2EF | `HandleExitUnsupported` | Known | Event handler |
| 0x007DA321 | `HandleExitUnsupported` | Known | Event handler |
| 0x007DF355 | `HandleSelectKey` | Known | Event handler |
| 0x007DF38A | `HandleWheel` | Known | Event handler |
| 0x007DF4D8 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x007DF52B | `HandleSelectKey` | Known | Event handler |
| 0x007DF553 | `HandleSelectKey` | Known | Event handler |
| 0x007DF583 | `HandleExit` | Known | Event handler |
| 0x007DF5AD | `HandleStartStop` | Known | Event handler |
| 0x007DF613 | `HandleStartStop` | Known | Event handler |
| 0x007DF72B | `HandleExit` | Known | Event handler |
| 0x007DF755 | `HandleStartStop` | Known | Event handler |
| 0x007DF781 | `HandleLap` | Known | Event handler |
| 0x007DF885 | `HandleSelectLozinch` | Known | Event handler |
| 0x007E0328 | `HandleSelect` | Known | Event handler |
| 0x007E0D77 | `HandleChoosePowerPlay` | Known | Event handler |
| 0x007E0DB2 | `HandleChooseVoicePrompts` | Known | Event handler |
| 0x007E0DF0 | `HandleChooseUnit` | Known | Event handler |
| 0x007E0F84 | `HandleListChoose` | Known | Event handler |
| 0x007E11E3 | `HandleSelect` | Known | Event handler |
| 0x007E1403 | `HandleSelect` | Known | Event handler |
| 0x007E1439 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007E166A | `HandleNowPlayingSelected` | Known | Event handler |
| 0x007E16A8 | `HandleChoosePlaylistsPlay` | Known | Event handler |
| 0x007E16E7 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x007E1727 | `HandleNoneSelected` | Known | Event handler |
| 0x007E175D | `HandleBegin` | Known | Event handler |
| 0x007E1A4A | `HandleBegin` | Known | Event handler |
| 0x007E1A79 | `HandleBegin` | Known | Event handler |
| 0x007E1B35 | `HandleBegin` | Known | Event handler |
| 0x007E1B61 | `HandleBegin` | Known | Event handler |
| 0x007E1C1D | `HandleBegin` | Known | Event handler |
| 0x007E1C49 | `HandleBegin` | Known | Event handler |
| 0x007E1D05 | `HandleBegin` | Known | Event handler |
| 0x007E1D39 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007E1D64 | `HandleMenuKey` | Known | Event handler |
| 0x007E1DFB | `HandlePauseHold` | Known | Event handler |
| 0x007E1E2A | `HandlePauseKey` | Known | Event handler |
| 0x007E1EB4 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007E1EEE | `HandlePowerPlay` | Known | Event handler |
| 0x007E1F1A | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E2387 | `HandlePauseHold` | Known | Event handler |
| 0x007E23B6 | `HandlePauseKey` | Known | Event handler |
| 0x007E23E1 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E241F | `HandlePowerPlay` | Known | Event handler |
| 0x007E244E | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E2474 | `HandleWheel` | Known | Event handler |
| 0x007E24A9 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007E24D4 | `HandleMenuKey` | Known | Event handler |
| 0x007E256B | `HandlePauseHold` | Known | Event handler |
| 0x007E259A | `HandlePauseKey` | Known | Event handler |
| 0x007E2624 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007E2654 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E2AB4 | `HandlePauseHold` | Known | Event handler |
| 0x007E2AE3 | `HandlePauseKey` | Known | Event handler |
| 0x007E2B0E | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E2B42 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E2B68 | `HandleWheel` | Known | Event handler |
| 0x007E2B9D | `HandleMenuKeyNop` | Known | Event handler |
| 0x007E2BC8 | `HandleMenuKey` | Known | Event handler |
| 0x007E2C5F | `HandlePauseHold` | Known | Event handler |
| 0x007E2C8E | `HandlePauseKey` | Known | Event handler |
| 0x007E2D18 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007E2D52 | `HandlePowerPlay` | Known | Event handler |
| 0x007E2D7E | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E31EA | `HandlePauseHold` | Known | Event handler |
| 0x007E3219 | `HandlePauseKey` | Known | Event handler |
| 0x007E3244 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E3282 | `HandlePowerPlay` | Known | Event handler |
| 0x007E32B1 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E32D7 | `HandleWheel` | Known | Event handler |
| 0x007E330D | `HandleMenuKeyNop` | Known | Event handler |
| 0x007E3338 | `HandleMenuKey` | Known | Event handler |
| 0x007E33CF | `HandlePauseHold` | Known | Event handler |
| 0x007E33FE | `HandlePauseKey` | Known | Event handler |
| 0x007E3488 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007E34B8 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E3917 | `HandlePauseHold` | Known | Event handler |
| 0x007E3946 | `HandlePauseKey` | Known | Event handler |
| 0x007E3971 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E39A5 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E39CB | `HandleWheel` | Known | Event handler |
| 0x007E3A01 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007E3A2C | `HandleMenuKey` | Known | Event handler |
| 0x007E3AC3 | `HandlePauseHold` | Known | Event handler |
| 0x007E3AF2 | `HandlePauseKey` | Known | Event handler |
| 0x007E3B7C | `HandleSelectKeyDown` | Known | Event handler |
| 0x007E3BB6 | `HandlePowerPlay` | Known | Event handler |
| 0x007E3BE2 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E4052 | `HandlePauseHold` | Known | Event handler |
| 0x007E4081 | `HandlePauseKey` | Known | Event handler |
| 0x007E40AC | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E40EA | `HandlePowerPlay` | Known | Event handler |
| 0x007E4119 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E413F | `HandleWheel` | Known | Event handler |
| 0x007E4175 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007E41A0 | `HandleMenuKey` | Known | Event handler |
| 0x007E4237 | `HandlePauseHold` | Known | Event handler |
| 0x007E4266 | `HandlePauseKey` | Known | Event handler |
| 0x007E42F0 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007E4320 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E4783 | `HandlePauseHold` | Known | Event handler |
| 0x007E47B2 | `HandlePauseKey` | Known | Event handler |
| 0x007E47DD | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E4811 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E4837 | `HandleWheel` | Known | Event handler |
| 0x007E486D | `HandleMenuKeyNop` | Known | Event handler |
| 0x007E4898 | `HandleMenuKey` | Known | Event handler |
| 0x007E492F | `HandlePauseHold` | Known | Event handler |
| 0x007E495E | `HandlePauseKey` | Known | Event handler |
| 0x007E49E8 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007E4A22 | `HandlePowerPlay` | Known | Event handler |
| 0x007E4A4E | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E4EBE | `HandlePauseHold` | Known | Event handler |
| 0x007E4EED | `HandlePauseKey` | Known | Event handler |
| 0x007E4F18 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E4F56 | `HandlePowerPlay` | Known | Event handler |
| 0x007E4F85 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E4FAB | `HandleWheel` | Known | Event handler |
| 0x007E4FE1 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007E500C | `HandleMenuKey` | Known | Event handler |
| 0x007E50A3 | `HandlePauseHold` | Known | Event handler |
| 0x007E50D2 | `HandlePauseKey` | Known | Event handler |
| 0x007E515C | `HandleSelectKeyDown` | Known | Event handler |
| 0x007E518C | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E55EF | `HandlePauseHold` | Known | Event handler |
| 0x007E561E | `HandlePauseKey` | Known | Event handler |
| 0x007E5649 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E567D | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E56A3 | `HandleWheel` | Known | Event handler |
| 0x007E56D9 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007E5704 | `HandleMenuKey` | Known | Event handler |
| 0x007E579B | `HandlePauseHold` | Known | Event handler |
| 0x007E57CA | `HandlePauseKey` | Known | Event handler |
| 0x007E5854 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007E5884 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E5C7D | `HandlePauseHold` | Known | Event handler |
| 0x007E5CAC | `HandlePauseKey` | Known | Event handler |
| 0x007E5CD7 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E5D0B | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007E5D31 | `HandleWheel` | Known | Event handler |
| 0x007E5D65 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007E5D90 | `HandleResumeWorkout` | Known | Event handler |
| 0x007E5E6B | `HandleResumeWorkout` | Known | Event handler |
| 0x007E5EDF | `HandlePauseWorkout` | Known | Event handler |
| 0x007E5F4D | `HandleChooseMusic` | Known | Event handler |
| 0x007E5FEA | `HandleEndWorkout` | Known | Event handler |
| 0x007E6095 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007E633C | `HandleEndWorkout` | Known | Event handler |
| 0x007E67CB | `HandleSelectResume` | Known | Event handler |
| 0x007E6803 | `HandleEndWorkout` | Known | Event handler |
| 0x007E68AE | `HandleChooseCustomTimedWorkout` | Known | Event handler |
| 0x007E6947 | `HandleChooseTimedWorkoutFromList` | Known | Event handler |
| 0x007E69FA | `HandleChooseCustomDistanceWorkout` | Known | Event handler |
| 0x007E6A9A | `HandleChooseDistanceWorkoutFromList` | Known | Event handler |
| 0x007E6C80 | `HandleChooseCustomCalorieWorkout` | Known | Event handler |
| 0x007E6D1F | `HandleChooseCalorieWorkoutFromList` | Known | Event handler |
| 0x007E708E | `HandleChooseLink` | Known | Event handler |
| 0x007E70C4 | `HandleChooseCalibrate` | Known | Event handler |
| 0x007E741D | `HandleChooseRemoteSetting` | Known | Event handler |
| 0x007E745C | `HandleChooseRemoteLink` | Known | Event handler |
| 0x007E7498 | `HandleChooseRemoteUnlink` | Known | Event handler |
| 0x007E7846 | `Handle400MetersWalk` | Known | Event handler |
| 0x007E787F | `HandleCustomWalk` | Known | Event handler |
| 0x007E7955 | `HandleSelectWalking` | Known | Event handler |
| 0x007E7A79 | `HandleSelectRunning` | Known | Event handler |
| 0x007E7DC6 | `Handle400MetersRun` | Known | Event handler |
| 0x007E7DFE | `HandleCustomRun` | Known | Event handler |
| 0x007E80D1 | `HandleSelect` | Known | Event handler |
| 0x007E8101 | `HandleSelect` | Known | Event handler |
| 0x007E8277 | `HandleLinkNewRemote` | Known | Event handler |
| 0x007E83E5 | `HandleSelect` | Known | Event handler |
| 0x007E8413 | `HandleCancelRemoteLinking` | Known | Event handler |
| 0x007E8480 | `HandleSelect` | Known | Event handler |
| 0x007E8971 | `HandleUnlinkRemote` | Known | Event handler |
| 0x007E8BD5 | `HandleWeightSelect` | Known | Event handler |
| 0x007E8C32 | `HandleWeightWheel` | Known | Event handler |
| 0x007E8C65 | `HandleWeightSelect` | Known | Event handler |
| 0x007E8CEF | `HandleWeightWheel` | Known | Event handler |
| 0x007E8D21 | `HandleDistanceSelect` | Known | Event handler |
| 0x007E8DAD | `HandleDistanceWheel` | Known | Event handler |
| 0x007E8DE1 | `HandleDistanceSelect` | Known | Event handler |
| 0x007E8E6D | `HandleDistanceWheel` | Known | Event handler |
| 0x007E8EA1 | `HandleTimeSelect` | Known | Event handler |
| 0x007E8F29 | `HandleTimeWheel` | Known | Event handler |
| 0x007E8F59 | `HandleCaloriesSelect` | Known | Event handler |
| 0x007E90B1 | `HandleCaloriesWheel` | Known | Event handler |
| 0x007E941D | `HandleChooseLast` | Known | Event handler |
| 0x007E9453 | `HandleChooseRecent` | Known | Event handler |
| 0x007E948B | `HandleChooseBest` | Known | Event handler |
| 0x007E97A1 | `HandleSelect` | Known | Event handler |
| 0x007E9989 | `HandleSelectFromLastWorkout` | Known | Event handler |
| 0x007E9B81 | `HandleSelect` | Known | Event handler |
| 0x007E9E3A | `HandleChooseRecentWorkout` | Known | Event handler |
| 0x007E9F0D | `HandleSelect` | Known | Event handler |
| 0x007E9FA1 | `HandleSelect_Basic` | Known | Event handler |
| 0x007EA285 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x007EA579 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x007EA869 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x007EAC5A | `HandleSelect` | Known | Event handler |
| 0x007EACE6 | `HandleSelect` | Known | Event handler |
| 0x007EAD74 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x007EB05E | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x007EB13F | `HandlePlayPause` | Known | Event handler |
| 0x007EB1CD | `HandlePlayPause` | Known | Event handler |
| 0x007EB25D | `HandleDeleteAllSelect` | Known | Event handler |
| 0x007EB295 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x007EB2D1 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x007EB314 | `HandlePlayPause` | Known | Event handler |
| 0x007EB34A | `HandleAddToOTG` | Known | Event handler |
| 0x007EB59F | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007EB7FB | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x0080E8AE | `HandleSelectClock` | Known | Event handler |
| 0x0080E8E7 | `HandleHilited` | Known | Event handler |
| 0x0080E919 | `HandleWheel` | Known | Event handler |
| 0x0080E960 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x0080E9E5 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x0080EBF1 | `HandleImageLast` | Known | Event handler |
| 0x0080EC1B | `HandleScreenNext` | Known | Event handler |
| 0x0080EC4B | `HandleBrowseSlideshow` | Known | Event handler |
| 0x0080EC85 | `HandleImageFirst` | Known | Event handler |
| 0x0080ECB0 | `HandleScreenPrev` | Known | Event handler |
| 0x0080ECDD | `HandleBrowseLarge` | Known | Event handler |
| 0x0080ED5D | `HandleImageNext` | Known | Event handler |
| 0x0080ED86 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x0080EDBA | `HandleBrowseSlideshow` | Known | Event handler |
| 0x0080EDE9 | `HandleImagePrev` | Known | Event handler |
| 0x0080EE17 | `HandleBrowseSmall` | Known | Event handler |

---

## 6. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001032E4 | `GotoNowPlaying` | Known | Navigation |
| 0x0010335C | `GotoMainMenu` | Known | Navigation |
| 0x00121548 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x00121560 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x001216D8 | `GotoScreen_AddressBook` | Known | Navigation |
| 0x0012CB44 | `GotoNowPlaying` | Known | Navigation |
| 0x0012CB58 | `GotoAlbums` | Known | Navigation |
| 0x0012CB64 | `GotoSongs` | Known | Navigation |
| 0x0013A53C | `GotoScreen_EnterPassKey` | Known | Navigation |
| 0x0013A554 | `GotoScreen_LockediPod` | Known | Navigation |
| 0x0013AF58 | `GotoScreen_MainMenu` | Known | Navigation |
| 0x0015445C | `GotoMainMenu` | Known | Navigation |
| 0x001DCDC8 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001E8358 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001E8BA8 | `GotoPlayDeleteMenu` | Known | Navigation |
| 0x001E8C2C | `GotoNowPlaying` | Known | Navigation |
| 0x00207B5C | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation |
| 0x00215218 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x00215310 | `GotoScreen_SettingsMenu` | Known | Navigation |
| 0x0021E178 | `GotoDefaultLayout` | Known | Navigation |
| 0x0021E1FC | `GotoVolumeLayout` | Known | Navigation |
| 0x0021E334 | `GotoProgressLayout` | Known | Navigation |
| 0x0021E650 | `GotoDefault` | Known | Navigation |
| 0x0021E984 | `GotoProgressLayout` | Known | Navigation |
| 0x0021EB44 | `GotoRentalWarningLayout` | Known | Navigation |
| 0x0021EBC8 | `GotoProgressLayout` | Known | Navigation |
| 0x0021EED8 | `GotoProgressLayout` | Known | Navigation |
| 0x00220A3C | `GotoNowPlaying` | Known | Navigation |
| 0x00221308 | `GotoNowPlaying` | Known | Navigation |
| 0x00224BD8 | `GotoScreen_Language` | Known | Navigation |
| 0x00224F38 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x00224F54 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00224F6C | `GotoDefaultLayout` | Known | Navigation |
| 0x00224F80 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00225018 | `GotoVolumeLayout` | Known | Navigation |
| 0x0022502C | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x002250CC | `GotoProgressLayout` | Known | Navigation |
| 0x002250E0 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x002255A8 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x0022586C | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x00225A08 | `GotoProgressLayout` | Known | Navigation |
| 0x00225A1C | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00225AE0 | `GotoBrightnessVideoLayout` | Known | Navigation |
| 0x00225AFC | `GotoRatingLayout` | Known | Navigation |
| 0x00225DA0 | `GotoChapterArtLayout` | Known | Navigation |
| 0x00225DB8 | `GotoShuffleLayout` | Known | Navigation |
| 0x00226110 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x00226124 | `GotoExtraInfoLoadingLayout` | Known | Navigation |
| 0x002261F4 | `GotoVolumeLayout` | Known | Navigation |
| 0x0022620C | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x00226298 | `GotoVolumeLayout` | Known | Navigation |
| 0x002262AC | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x002264BC | `GotoScrubLayout` | Known | Navigation |
| 0x002264CC | `GotoScrubVideoLayout` | Known | Navigation |
| 0x0022655C | `GotoProgressLayout` | Known | Navigation |
| 0x00226570 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00226710 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x0022672C | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00226744 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00226760 | `GotoDefaultLayout` | Known | Navigation |
| 0x002269A4 | `GotoChapterArtLayout` | Known | Navigation |
| 0x00226A9C | `GotoProgressLayout` | Known | Navigation |
| 0x00226B28 | `GotoProgressLayout` | Known | Navigation |
| 0x00226B3C | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00226C18 | `GotoExtraInfoLoadFailedLayout` | Known | Navigation |
| 0x00226C38 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x00226ED0 | `GotoStatusBarLayout` | Known | Navigation |
| 0x00226EE4 | `GotoDefaultLayout` | Known | Navigation |
| 0x002270BC | `GotoDefault` | Known | Navigation |
| 0x002271F0 | `GotoProgressLayout` | Known | Navigation |
| 0x002273B0 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x00227500 | `GotoBrightnessLayout` | Known | Navigation |
| 0x00227584 | `GotoBrightnessLayout` | Known | Navigation |
| 0x00227604 | `GotoVolumeLayout` | Known | Navigation |
| 0x00227650 | `GotoScrubLayout` | Known | Navigation |
| 0x00227718 | `GotoStatusBarLayout` | Known | Navigation |
| 0x0022772C | `GotoDefaultLayout` | Known | Navigation |
| 0x00227804 | `GotoScrubLayout` | Known | Navigation |
| 0x00227854 | `GotoScrubLayout` | Known | Navigation |
| 0x0022E32C | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x0022E4BC | `GotoFourCard_About` | Known | Navigation |
| 0x0022E4D0 | `GotoThreeCard_About` | Known | Navigation |
| 0x0022E5BC | `GotoScreen_BacklightTimer` | Known | Navigation |
| 0x0022E64C | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x0022E664 | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation |
| 0x002338D4 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x002338EC | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x00235EE0 | `GotoNowPlaying` | Known | Navigation |
| 0x00236638 | `GotoNowPlaying` | Known | Navigation |
| 0x00236CB8 | `GotoFirstBoot` | Known | Navigation |
| 0x00236CC8 | `GotoNotesApp` | Known | Navigation |
| 0x00236CDC | `GotoLockApp` | Known | Navigation |
| 0x0023C594 | `GotoNowPlaying` | Known | Navigation |
| 0x003D1D80 | `GotoProgressLayout` | Known | Navigation |
| 0x007564D7 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x007CE7F9 | `GotoDefault` | Known | Navigation |
| 0x007CF2E9 | `GotoDefault` | Known | Navigation |
| 0x008D6428 | `Gotowe?` | Known | Navigation |

---

## 7. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016B158 | `CoverFlow_Screen` | Known | Screen layout |
| 0x001966BC | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x001966DC | `VoiceMemos_Menu_Screen_Recording` | Known | Screen layout |
| 0x00196700 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x0074A4E2 | `Clock_Screen` | Known | Screen layout |
| 0x0074A4F2 | `Clock_Screen_Default"` | Known | Screen layout |
| 0x0074A557 | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x0074A5B5 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0074A5CD | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x0074A63A | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x0074A6D8 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x0074A737 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0074A74D | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x0074A7B8 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x0074A812 | `Games_Menu_Screen` | Known | Screen layout |
| 0x0074A827 | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x0074A891 | `Extras_Screen_Games` | Known | Screen layout |
| 0x0074A950 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x0074AA14 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0074AADD | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x0074AB3A | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x0074AB53 | `Debug_MainMenu_Screen_Default"` | Known | Screen layout |
| 0x0074ABC1 | `Extras_Screen_Debug` | Known | Screen layout |
| 0x0074ACF8 | `Clock_Africa_City_Screen ` | Known | Screen layout |
| 0x0074AD14 | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x0074AD98 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x0074ADB2 | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x0074AE34 | `Clock_Atlantic_City_Screen"` | Known | Screen layout |
| 0x0074AE52 | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x0074AED8 | `Clock_Australia_City_Screen#` | Known | Screen layout |
| 0x0074AEF7 | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x0074AF7E | `Clock_Europe_City_Screen ` | Known | Screen layout |
| 0x0074AF9A | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x0074B01E | `Clock_NorthAmerica_City_Screen&` | Known | Screen layout |
| 0x0074B040 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0074B0CA | `Clock_Pacific_City_Screen!` | Known | Screen layout |
| 0x0074B0E7 | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x0074B16C | `Clock_SouthAmerica_City_Screen&` | Known | Screen layout |
| 0x0074B18E | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0074B21B | `Clock_Screen"` | Known | Screen layout |
| 0x0074B2C0 | `Clock_Screen"` | Known | Screen layout |
| 0x0074B365 | `Clock_Screen"` | Known | Screen layout |
| 0x0074B40A | `Clock_Screen"` | Known | Screen layout |
| 0x0074B4AF | `Clock_Screen"` | Known | Screen layout |
| 0x0074B554 | `Clock_Screen"` | Known | Screen layout |
| 0x0074B5F9 | `Clock_Screen"` | Known | Screen layout |
| 0x0074B69E | `Clock_Screen"` | Known | Screen layout |
| 0x0074B743 | `Clock_Screen"` | Known | Screen layout |
| 0x0074B7E8 | `Clock_Screen"` | Known | Screen layout |
| 0x0074B88D | `Clock_Screen"` | Known | Screen layout |
| 0x0074B932 | `Clock_Screen"` | Known | Screen layout |
| 0x0074B9D7 | `Clock_Screen"` | Known | Screen layout |
| 0x0074BA7C | `Clock_Screen"` | Known | Screen layout |
| 0x0074BB21 | `Clock_Screen"` | Known | Screen layout |
| 0x0074BBC6 | `Clock_Screen"` | Known | Screen layout |
| 0x0074BC6B | `Clock_Screen"` | Known | Screen layout |
| 0x0074BD10 | `Clock_Screen"` | Known | Screen layout |
| 0x0074BDB5 | `Clock_Screen"` | Known | Screen layout |
| 0x0074BE5A | `Clock_Screen"` | Known | Screen layout |
| 0x0074BEFF | `Clock_Screen"` | Known | Screen layout |
| 0x0074BFA4 | `Clock_Screen"` | Known | Screen layout |
| 0x0074C049 | `Clock_Screen"` | Known | Screen layout |
| 0x0074C0EE | `Clock_Screen"` | Known | Screen layout |
| 0x0074C193 | `Clock_Screen"` | Known | Screen layout |
| 0x0074C238 | `Clock_Screen"` | Known | Screen layout |
| 0x0074C2DD | `Clock_Screen"` | Known | Screen layout |
| 0x0074C382 | `Clock_Screen"` | Known | Screen layout |
| 0x0074C427 | `Clock_Screen"` | Known | Screen layout |
| 0x0074C4CC | `Clock_Screen"` | Known | Screen layout |
| 0x0074C571 | `Clock_Screen"` | Known | Screen layout |
| 0x0074C61B | `Clock_Screen"` | Known | Screen layout |
| 0x0074C6C0 | `Clock_Screen"` | Known | Screen layout |
| 0x0074C765 | `Clock_Screen"` | Known | Screen layout |
| 0x0074C80A | `Clock_Screen"` | Known | Screen layout |
| 0x0074C8AF | `Clock_Screen"` | Known | Screen layout |
| 0x0074C954 | `Clock_Screen"` | Known | Screen layout |
| 0x0074C9F9 | `Clock_Screen"` | Known | Screen layout |
| 0x0074CA9E | `Clock_Screen"` | Known | Screen layout |
| 0x0074CB43 | `Clock_Screen"` | Known | Screen layout |
| 0x0074CBE8 | `Clock_Screen"` | Known | Screen layout |
| 0x0074CC8D | `Clock_Screen"` | Known | Screen layout |
| 0x0074CD32 | `Clock_Screen"` | Known | Screen layout |
| 0x0074CDD7 | `Clock_Screen"` | Known | Screen layout |
| 0x0074CE7C | `Clock_Screen"` | Known | Screen layout |
| 0x0074CF21 | `Clock_Screen"` | Known | Screen layout |
| 0x0074CFC6 | `Clock_Screen"` | Known | Screen layout |
| 0x0074D06B | `Clock_Screen"` | Known | Screen layout |
| 0x0074D110 | `Clock_Screen"` | Known | Screen layout |
| 0x0074D1B5 | `Clock_Screen"` | Known | Screen layout |
| 0x0074D25A | `Clock_Screen"` | Known | Screen layout |
| 0x0074D2FF | `Clock_Screen"` | Known | Screen layout |
| 0x0074D3A4 | `Clock_Screen"` | Known | Screen layout |
| 0x0074D449 | `Clock_Screen"` | Known | Screen layout |
| 0x0074D4EE | `Clock_Screen"` | Known | Screen layout |
| 0x0074D593 | `Clock_Screen"` | Known | Screen layout |
| 0x0074D638 | `Clock_Screen"` | Known | Screen layout |
| 0x0074D6DD | `Clock_Screen"` | Known | Screen layout |
| 0x0074D782 | `Clock_Screen"` | Known | Screen layout |
| 0x0074D827 | `Clock_Screen"` | Known | Screen layout |
| 0x0074D8CC | `Clock_Screen"` | Known | Screen layout |
| 0x0074D971 | `Clock_Screen"` | Known | Screen layout |
| 0x0074DA16 | `Clock_Screen"` | Known | Screen layout |
| 0x0074DABB | `Clock_Screen"` | Known | Screen layout |
| 0x0074DB60 | `Clock_Screen"` | Known | Screen layout |
| 0x0074DC05 | `Clock_Screen"` | Known | Screen layout |
| 0x0074DCAA | `Clock_Screen"` | Known | Screen layout |
| 0x0074DD4F | `Clock_Screen"` | Known | Screen layout |
| 0x0074DDF4 | `Clock_Screen"` | Known | Screen layout |
| 0x0074DE99 | `Clock_Screen"` | Known | Screen layout |
| 0x0074DF3E | `Clock_Screen"` | Known | Screen layout |
| 0x0074DFE3 | `Clock_Screen"` | Known | Screen layout |
| 0x0074E088 | `Clock_Screen"` | Known | Screen layout |
| 0x0074E12D | `Clock_Screen"` | Known | Screen layout |
| 0x0074E1D2 | `Clock_Screen"` | Known | Screen layout |
| 0x0074E277 | `Clock_Screen"` | Known | Screen layout |
| 0x0074E31C | `Clock_Screen"` | Known | Screen layout |
| 0x0074E3C1 | `Clock_Screen"` | Known | Screen layout |
| 0x0074E466 | `Clock_Screen"` | Known | Screen layout |
| 0x0074E50B | `Clock_Screen"` | Known | Screen layout |
| 0x0074E5B0 | `Clock_Screen"` | Known | Screen layout |
| 0x0074E655 | `Clock_Screen"` | Known | Screen layout |
| 0x0074E6FA | `Clock_Screen"` | Known | Screen layout |
| 0x0074E79F | `Clock_Screen"` | Known | Screen layout |
| 0x0074E844 | `Clock_Screen"` | Known | Screen layout |
| 0x0074E8E9 | `Clock_Screen"` | Known | Screen layout |
| 0x0074E98E | `Clock_Screen"` | Known | Screen layout |
| 0x0074EA33 | `Clock_Screen"` | Known | Screen layout |
| 0x0074EADF | `Clock_Screen"` | Known | Screen layout |
| 0x0074EB84 | `Clock_Screen"` | Known | Screen layout |
| 0x0074EC29 | `Clock_Screen"` | Known | Screen layout |
| 0x0074ECD3 | `Clock_Screen"` | Known | Screen layout |
| 0x0074ED78 | `Clock_Screen"` | Known | Screen layout |
| 0x0074EE1D | `Clock_Screen"` | Known | Screen layout |
| 0x0074EEC2 | `Clock_Screen"` | Known | Screen layout |
| 0x0074EF67 | `Clock_Screen"` | Known | Screen layout |
| 0x0074F00C | `Clock_Screen"` | Known | Screen layout |
| 0x0074F0B1 | `Clock_Screen"` | Known | Screen layout |
| 0x0074F156 | `Clock_Screen"` | Known | Screen layout |
| 0x0074F1FF | `Clock_Screen"` | Known | Screen layout |
| 0x0074F2A4 | `Clock_Screen"` | Known | Screen layout |
| 0x0074F349 | `Clock_Screen"` | Known | Screen layout |
| 0x0074F3EE | `Clock_Screen"` | Known | Screen layout |
| 0x0074F493 | `Clock_Screen"` | Known | Screen layout |
| 0x0074F538 | `Clock_Screen"` | Known | Screen layout |
| 0x0074F5DD | `Clock_Screen"` | Known | Screen layout |
| 0x0074F682 | `Clock_Screen"` | Known | Screen layout |
| 0x0074F727 | `Clock_Screen"` | Known | Screen layout |
| 0x0074F7CC | `Clock_Screen"` | Known | Screen layout |
| 0x0074F871 | `Clock_Screen"` | Known | Screen layout |
| 0x0074F916 | `Clock_Screen"` | Known | Screen layout |
| 0x0074F9BB | `Clock_Screen"` | Known | Screen layout |
| 0x0074FA60 | `Clock_Screen"` | Known | Screen layout |
| 0x0074FB05 | `Clock_Screen"` | Known | Screen layout |
| 0x0074FBAA | `Clock_Screen"` | Known | Screen layout |
| 0x0074FC4F | `Clock_Screen"` | Known | Screen layout |
| 0x0074FCF4 | `Clock_Screen"` | Known | Screen layout |
| 0x0074FD99 | `Clock_Screen"` | Known | Screen layout |
| 0x0074FE3E | `Clock_Screen"` | Known | Screen layout |
| 0x0074FEE3 | `Clock_Screen"` | Known | Screen layout |
| 0x0074FF88 | `Clock_Screen"` | Known | Screen layout |
| 0x0075002D | `Clock_Screen"` | Known | Screen layout |
| 0x007500D2 | `Clock_Screen"` | Known | Screen layout |
| 0x00750177 | `Clock_Screen"` | Known | Screen layout |
| 0x0075021C | `Clock_Screen"` | Known | Screen layout |
| 0x007502C1 | `Clock_Screen"` | Known | Screen layout |
| 0x00750366 | `Clock_Screen"` | Known | Screen layout |
| 0x0075040B | `Clock_Screen"` | Known | Screen layout |
| 0x007504B0 | `Clock_Screen"` | Known | Screen layout |
| 0x00750555 | `Clock_Screen"` | Known | Screen layout |
| 0x007505FA | `Clock_Screen"` | Known | Screen layout |
| 0x0075069F | `Clock_Screen"` | Known | Screen layout |
| 0x00750744 | `Clock_Screen"` | Known | Screen layout |
| 0x007507EF | `Clock_Screen"` | Known | Screen layout |
| 0x00750894 | `Clock_Screen"` | Known | Screen layout |
| 0x00750939 | `Clock_Screen"` | Known | Screen layout |
| 0x007509DE | `Clock_Screen"` | Known | Screen layout |
| 0x00750A83 | `Clock_Screen"` | Known | Screen layout |
| 0x00750B28 | `Clock_Screen"` | Known | Screen layout |
| 0x00750BCD | `Clock_Screen"` | Known | Screen layout |
| 0x00750C72 | `Clock_Screen"` | Known | Screen layout |
| 0x00750D17 | `Clock_Screen"` | Known | Screen layout |
| 0x00750DBC | `Clock_Screen"` | Known | Screen layout |
| 0x00750E61 | `Clock_Screen"` | Known | Screen layout |
| 0x00750F06 | `Clock_Screen"` | Known | Screen layout |
| 0x00750FAB | `Clock_Screen"` | Known | Screen layout |
| 0x00751050 | `Clock_Screen"` | Known | Screen layout |
| 0x007510F5 | `Clock_Screen"` | Known | Screen layout |
| 0x0075119A | `Clock_Screen"` | Known | Screen layout |
| 0x0075123F | `Clock_Screen"` | Known | Screen layout |
| 0x007512E4 | `Clock_Screen"` | Known | Screen layout |
| 0x00751389 | `Clock_Screen"` | Known | Screen layout |
| 0x0075142E | `Clock_Screen"` | Known | Screen layout |
| 0x007514D3 | `Clock_Screen"` | Known | Screen layout |
| 0x00751578 | `Clock_Screen"` | Known | Screen layout |
| 0x0075161D | `Clock_Screen"` | Known | Screen layout |
| 0x007516C2 | `Clock_Screen"` | Known | Screen layout |
| 0x00751767 | `Clock_Screen"` | Known | Screen layout |
| 0x0075180C | `Clock_Screen"` | Known | Screen layout |
| 0x007518B1 | `Clock_Screen"` | Known | Screen layout |
| 0x00751956 | `Clock_Screen"` | Known | Screen layout |
| 0x007519FB | `Clock_Screen"` | Known | Screen layout |
| 0x00751AA0 | `Clock_Screen"` | Known | Screen layout |
| 0x00751B45 | `Clock_Screen"` | Known | Screen layout |
| 0x00751BEA | `Clock_Screen"` | Known | Screen layout |
| 0x00751C8F | `Clock_Screen"` | Known | Screen layout |
| 0x00751D34 | `Clock_Screen"` | Known | Screen layout |
| 0x00751DD9 | `Clock_Screen"` | Known | Screen layout |
| 0x00751E7E | `Clock_Screen"` | Known | Screen layout |
| 0x00751F23 | `Clock_Screen"` | Known | Screen layout |
| 0x00751FC8 | `Clock_Screen"` | Known | Screen layout |
| 0x0075206D | `Clock_Screen"` | Known | Screen layout |
| 0x00752112 | `Clock_Screen"` | Known | Screen layout |
| 0x007521B7 | `Clock_Screen"` | Known | Screen layout |
| 0x0075225C | `Clock_Screen"` | Known | Screen layout |
| 0x00752301 | `Clock_Screen"` | Known | Screen layout |
| 0x007523A6 | `Clock_Screen"` | Known | Screen layout |
| 0x0075244B | `Clock_Screen"` | Known | Screen layout |
| 0x007524F0 | `Clock_Screen"` | Known | Screen layout |
| 0x00752595 | `Clock_Screen"` | Known | Screen layout |
| 0x0075263A | `Clock_Screen"` | Known | Screen layout |
| 0x007526DF | `Clock_Screen"` | Known | Screen layout |
| 0x00752784 | `Clock_Screen"` | Known | Screen layout |
| 0x0075282F | `Clock_Screen"` | Known | Screen layout |
| 0x007528D4 | `Clock_Screen"` | Known | Screen layout |
| 0x00752979 | `Clock_Screen"` | Known | Screen layout |
| 0x00752A1E | `Clock_Screen"` | Known | Screen layout |
| 0x00752AC3 | `Clock_Screen"` | Known | Screen layout |
| 0x00752B6F | `Clock_Screen"` | Known | Screen layout |
| 0x00752C14 | `Clock_Screen"` | Known | Screen layout |
| 0x00752CB9 | `Clock_Screen"` | Known | Screen layout |
| 0x00752D5E | `Clock_Screen"` | Known | Screen layout |
| 0x00752E03 | `Clock_Screen"` | Known | Screen layout |
| 0x00752EA8 | `Clock_Screen"` | Known | Screen layout |
| 0x00752F4D | `Clock_Screen"` | Known | Screen layout |
| 0x00752FF2 | `Clock_Screen"` | Known | Screen layout |
| 0x00753097 | `Clock_Screen"` | Known | Screen layout |
| 0x0075313C | `Clock_Screen"` | Known | Screen layout |
| 0x007531E1 | `Clock_Screen"` | Known | Screen layout |
| 0x00753286 | `Clock_Screen"` | Known | Screen layout |
| 0x0075332B | `Clock_Screen"` | Known | Screen layout |
| 0x007533D0 | `Clock_Screen"` | Known | Screen layout |
| 0x00753475 | `Clock_Screen"` | Known | Screen layout |
| 0x0075351A | `Clock_Screen"` | Known | Screen layout |
| 0x007535BF | `Clock_Screen"` | Known | Screen layout |
| 0x00753662 | `Settings_DateTime_SetDate_Screen(` | Known | Screen layout |
| 0x00753686 | `Settings_DateTime_SetDate_Screen_Default"` | Known | Screen layout |
| 0x007536FF | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x00753765 | `Settings_DateTime_SetTime_Screen(` | Known | Screen layout |
| 0x00753789 | `Settings_DateTime_SetTime_Screen_Default"` | Known | Screen layout |
| 0x00753802 | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x0075386D | `Settings_DateTime_SetTimeZone_Screen,` | Known | Screen layout |
| 0x00753895 | `Settings_DateTime_SetTimeZone_Screen_Default"` | Known | Screen layout |
| 0x00753912 | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x007539CB | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x00753A7B | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0075400A | `Search_Main_Screen` | Known | Screen layout |
| 0x00754020 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x00754542 | `Extras_Screen` | Known | Screen layout |
| 0x00754553 | `Extras_Screen_WorldClock"` | Known | Screen layout |
| 0x007545D0 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x00754632 | `Clock_Screen` | Known | Screen layout |
| 0x00754642 | `Clock_Screen_Default` | Known | Screen layout |
| 0x007546C9 | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x0075472F | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x00754745 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x007547B0 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x00754812 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0075482A | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x00754897 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x007548FB | `AddressViewer_Main_Screen!` | Known | Screen layout |
| 0x00754918 | `AddressViewer_Main_Screen_Default"` | Known | Screen layout |
| 0x0075498A | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x007549F1 | `Games_Menu_Screen` | Known | Screen layout |
| 0x00754A06 | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x00754A70 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x00754B37 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x00754BD3 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x00754CA4 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x00754D64 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x00754DC8 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x00754DE7 | `VoiceMemos_Menu_Screen_Default"` | Known | Screen layout |
| 0x00754E6A | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x00754ED0 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x00754EE8 | `Speakers_Main_Screen_Default"` | Known | Screen layout |
| 0x00754F69 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x00754FCD | `Radio_Screen` | Known | Screen layout |
| 0x00754FDD | `Radio_Screen_Default"` | Known | Screen layout |
| 0x00755056 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x007550B7 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00755153 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x00755216 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x007552D5 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x00755392 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x007557AC | `Radio_Screen` | Known | Screen layout |
| 0x007557BC | `Radio_Screen_Default"` | Known | Screen layout |
| 0x00755835 | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x00755A19 | `Search_Main_Screen` | Known | Screen layout |
| 0x00755A2F | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x00755B5C | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x00755BBF | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x00755F00 | `Video_Settings_Screen` | Known | Screen layout |
| 0x00755F19 | `Video_Settings_Screen_Default"` | Known | Screen layout |
| 0x00756016 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007562DB | `SettingsMenus_MainMenu_Screen%` | Known | Screen layout |
| 0x007563E9 | `SettingsMenus_MusicMenu_Screen&` | Known | Screen layout |
| 0x00756692 | `SettingsMenus_Brightness_Screen&` | Known | Screen layout |
| 0x007567A7 | `SettingsMenus_AudiobookSettings_Screen.` | Known | Screen layout |
| 0x007568DD | `SettingsMenus_RadioRegions_Screen ` | Known | Screen layout |
| 0x007569F2 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x00756C5E | `Settings_DateTime_Screen ` | Known | Screen layout |
| 0x00756C7A | `Settings_DateTime_Screen_Default"` | Known | Screen layout |
| 0x00756E06 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x00756F0B | `Settings_Legal_Screen` | Known | Screen layout |
| 0x00756F24 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x00757015 | `SettingsMenus_ResetAllSettings_Screen4` | Known | Screen layout |
| 0x007577E6 | `Stopwatch_Screen` | Known | Screen layout |
| 0x007577FA | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x00757861 | `Stopwatch_Screen` | Known | Screen layout |
| 0x00757875 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x0075791E | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x00757941 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007579DA | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x007579FD | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x00757A8E | `NikePlus_ResumeWorkout_Screen%` | Known | Screen layout |
| 0x00757AAF | `NikePlus_ResumeWorkout_Screen_Default"` | Known | Screen layout |
| 0x00757B25 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00757BCB | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00757C78 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00757D28 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00757DD8 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00757E3A | `NikePlus_Settings_Screen ` | Known | Screen layout |
| 0x00757E56 | `NikePlus_Settings_Screen_Default` | Known | Screen layout |
| 0x00757ED9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00757F3B | `NikePlus_History_Screen` | Known | Screen layout |
| 0x00757F56 | `NikePlus_History_Screen_Default` | Known | Screen layout |
| 0x00757FD8 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007581FC | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x0075826A | `Speakers_ToneControl_Screen#` | Known | Screen layout |
| 0x00758289 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x0076D345 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076D3C8 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076D450 | `Lock_Screen` | Known | Screen layout |
| 0x0076D45F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076D4DA | `AddressViewer_ContactDetails_Screen+` | Known | Screen layout |
| 0x0076D501 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x0076D57C | `Extras_Screen` | Known | Screen layout |
| 0x0076D5C7 | `Extras_Screen` | Known | Screen layout |
| 0x0076D6AE | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x0076D70C | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0076D729 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x0076D797 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0076D7B0 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0076D827 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0076D844 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x0076D8AF | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x0076D8CC | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0076D933 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x0076D99A | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x0076D9F8 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0076DA15 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x0076DA83 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0076DA9C | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0076DB13 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0076DB30 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x0076DB9B | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x0076DBB8 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0076DC1F | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x0076DCBF | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x0076DD48 | `Alarms_Set_Alarm_Frequency_Screen)` | Known | Screen layout |
| 0x0076DD6D | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x0076DDDE | `Alarms_Set_Alarm_Label_Screen%` | Known | Screen layout |
| 0x0076DDFF | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x0076DE6C | `Alarms_Set_Alarm_Sound_Screen%` | Known | Screen layout |
| 0x0076DE8D | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x0076DEF9 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x0076E174 | `Alarms_Set_Alarm_Playlist_Screen(` | Known | Screen layout |
| 0x0076E198 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x0076E208 | `Alarms_Set_Alarm_Tones_Screen%` | Known | Screen layout |
| 0x0076E229 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x0076E53C | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x0076E557 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x0076E6A8 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0076E6BF | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x0076E740 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0076E757 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0076E82D | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0076E846 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0076E8CB | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x0076E93C | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0076EA31 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0076EA4A | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0076EACF | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x0076EB40 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0076EC00 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x0076EC14 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x0076ED43 | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x0076EDA6 | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x0076EDFD | `Clock_Screen_Default` | Known | Screen layout |
| 0x0076EE8E | `Clock_Region_Screen` | Known | Screen layout |
| 0x0076EEA5 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x0076EF1E | `Clock_Screen_Default` | Known | Screen layout |
| 0x0076EF75 | `Clock_Screen_Default` | Known | Screen layout |
| 0x0076F006 | `Clock_Region_Screen` | Known | Screen layout |
| 0x0076F01D | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x0076F1A8 | `CoverFlow_Screen_DefaultNoAlbumInfo&` | Known | Screen layout |
| 0x0076F296 | `CoverFlow_Screen_DefaultNoAlbumInfo#` | Known | Screen layout |
| 0x0076F30B | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0076F601 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0076F7B1 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0076F8DF | `CoverFlow_Screen_Default&` | Known | Screen layout |
| 0x0076F9B5 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0076FB4A | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0076FDAF | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x0076FE0C | `Game_Screen` | Known | Screen layout |
| 0x0076FE1B | `Game_Screen_Default` | Known | Screen layout |
| 0x0076FEBD | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0076FF1F | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0076FF82 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0076FFE5 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00770041 | `Game_Running_Screen` | Known | Screen layout |
| 0x007700A1 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00770103 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00770166 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x007701C9 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00770225 | `Game_Running_Screen` | Known | Screen layout |
| 0x00770285 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x007702E7 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0077034A | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x007703AD | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00770409 | `Game_Running_Screen` | Known | Screen layout |
| 0x00770469 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x007704CB | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0077052E | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00770591 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x007705ED | `Game_Running_Screen` | Known | Screen layout |
| 0x0077064D | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x007706AF | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00770712 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00770775 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x007707D1 | `Game_Running_Screen` | Known | Screen layout |
| 0x00770A17 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00770A79 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00770ADC | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00770B3F | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00770B9B | `Game_Running_Screen` | Known | Screen layout |
| 0x00770C52 | `Extras_Screen` | Known | Screen layout |
| 0x00770C63 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00770CC1 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x00770E5E | `Extras_Screen` | Known | Screen layout |
| 0x00770E6F | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00770ECD | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0077106A | `Extras_Screen` | Known | Screen layout |
| 0x0077107B | `Extras_Screen_Lock` | Known | Screen layout |
| 0x007710D9 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x00771276 | `Extras_Screen` | Known | Screen layout |
| 0x00771287 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x007712E5 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x00771487 | `Lock_Screen` | Known | Screen layout |
| 0x00771496 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x007714F8 | `Extras_Screen` | Known | Screen layout |
| 0x00771509 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x00771568 | `LockediPod_Screen` | Known | Screen layout |
| 0x007715E2 | `MainMenus_Main_Screen)` | Known | Screen layout |
| 0x007717B3 | `Lock_Screen` | Known | Screen layout |
| 0x007717C2 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x00771824 | `Extras_Screen` | Known | Screen layout |
| 0x00771835 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x00771894 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077190E | `MainMenus_Main_Screen%` | Known | Screen layout |
| 0x00771975 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077198A | `LockediPod_Screen_Default,` | Known | Screen layout |
| 0x00771AD9 | `Lock_Screen` | Known | Screen layout |
| 0x00771AE8 | `Lock_Screen_Default_Layout"` | Known | Screen layout |
| 0x00771B51 | `Lock_Screen` | Known | Screen layout |
| 0x00771B60 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x00771BC2 | `Extras_Screen` | Known | Screen layout |
| 0x00771BD3 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x00771C32 | `LockediPod_Screen` | Known | Screen layout |
| 0x00771CAC | `MainMenus_Main_Screen,` | Known | Screen layout |
| 0x00771E08 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00771E6E | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x00771ED2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00771F61 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x00771FCE | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0077203B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x007720A8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00772110 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00772176 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x007721DA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00772269 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x007722D6 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x00772343 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x007723B0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00772418 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0077247E | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x007724E2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00772571 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x007725DE | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0077264B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x007726B8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00772720 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00772786 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x007727EA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00772879 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x007728E6 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x00772953 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x007729C0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00772A28 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00772A8E | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x00772AF2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00772B81 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x00772BEE | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x00772C5B | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x00772CC8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00772D21 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x00772D8A | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00772DF1 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00772E8C | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00772EF5 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x00772F5E | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00772FC5 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00773060 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x007730C9 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x00773132 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00773199 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00773234 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00773320 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077333C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007733AA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007733C7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00773432 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00773452 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007734C9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007734E5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00773555 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00773574 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007735E0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007735F4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077366D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007736E1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00773751 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007737B9 | `NoContent_Screen` | Known | Screen layout |
| 0x007737CD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00773831 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00773898 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007738B2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00773920 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00773992 | `NoContent_Screen` | Known | Screen layout |
| 0x007739A6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00773A10 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00773A79 | `No_Photos_Screen` | Known | Screen layout |
| 0x00773A8D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00773AF3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00773B61 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00773BCE | `NoContent_Screen` | Known | Screen layout |
| 0x00773BE2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00773C4A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00773CB4 | `NoContent_Screen` | Known | Screen layout |
| 0x00773CC8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00773D2F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00773D99 | `NoContent_Screen` | Known | Screen layout |
| 0x00773DAD | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00773E1A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00773E8C | `NoContent_Screen` | Known | Screen layout |
| 0x00773EA0 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00773F08 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00773F71 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00773F8C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00773FF2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077400E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007740ED | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00774106 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00774167 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077417B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007742E9 | `Radio_Screen` | Known | Screen layout |
| 0x007742F9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077435A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007743DD | `LockediPod_Screen` | Known | Screen layout |
| 0x00774465 | `Lock_Screen` | Known | Screen layout |
| 0x00774474 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007744D7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00774539 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00774555 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007745C7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007745E6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077464E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00774668 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007746D0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007746ED | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00774759 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007747C3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007747DD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077484D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007748C0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00774931 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007749A0 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00774A0C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00774A27 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00774A9C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00774B03 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00774B65 | `Photos_Screen` | Known | Screen layout |
| 0x00774BC9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00774BE7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00774C59 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00774C76 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00774CDC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00774CF7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00774D60 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00774D7D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00774DF4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00774E18 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00774E86 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00774EA1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00774F5C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00774F78 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00774FE6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00775003 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077506E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077508E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00775105 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00775121 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00775191 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007751B0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077521C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00775230 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007752A9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077531D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077538D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007753F5 | `NoContent_Screen` | Known | Screen layout |
| 0x00775409 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077546D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007754D4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007754EE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077555C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007755CE | `NoContent_Screen` | Known | Screen layout |
| 0x007755E2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077564C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007756B5 | `No_Photos_Screen` | Known | Screen layout |
| 0x007756C9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077572F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077579D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077580A | `NoContent_Screen` | Known | Screen layout |
| 0x0077581E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00775886 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007758F0 | `NoContent_Screen` | Known | Screen layout |
| 0x00775904 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077596B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007759D5 | `NoContent_Screen` | Known | Screen layout |
| 0x007759E9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00775A56 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00775AC8 | `NoContent_Screen` | Known | Screen layout |
| 0x00775ADC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00775B44 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00775BAD | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00775BC8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00775C2E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00775C4A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00775D29 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00775D42 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00775DA3 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00775DB7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00775F25 | `Radio_Screen` | Known | Screen layout |
| 0x00775F35 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00775F96 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00776019 | `LockediPod_Screen` | Known | Screen layout |
| 0x007760A1 | `Lock_Screen` | Known | Screen layout |
| 0x007760B0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00776113 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00776175 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00776191 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00776203 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00776222 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077628A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007762A4 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077630C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00776329 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00776395 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007763FF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00776419 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00776489 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007764FC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077656D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007765DC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00776648 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00776663 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007766D8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077673F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007767A1 | `Photos_Screen` | Known | Screen layout |
| 0x00776805 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00776823 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00776895 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007768B2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00776918 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00776933 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077699C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007769B9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00776A30 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00776A54 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00776AC2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00776ADD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00776B98 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00776BB4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00776C22 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00776C3F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00776CAA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00776CCA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00776D41 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00776D5D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00776DCD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00776DEC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00776E58 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00776E6C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00776EE5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00776F59 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00776FC9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00777031 | `NoContent_Screen` | Known | Screen layout |
| 0x00777045 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007770A9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00777110 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077712A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00777198 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077720A | `NoContent_Screen` | Known | Screen layout |
| 0x0077721E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00777288 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007772F1 | `No_Photos_Screen` | Known | Screen layout |
| 0x00777305 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077736B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007773D9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00777446 | `NoContent_Screen` | Known | Screen layout |
| 0x0077745A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007774C2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077752C | `NoContent_Screen` | Known | Screen layout |
| 0x00777540 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007775A7 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00777611 | `NoContent_Screen` | Known | Screen layout |
| 0x00777625 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00777692 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00777704 | `NoContent_Screen` | Known | Screen layout |
| 0x00777718 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00777780 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007777E9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00777804 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077786A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00777886 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00777965 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077797E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007779DF | `FirstBoot_Screen` | Known | Screen layout |
| 0x007779F3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00777B61 | `Radio_Screen` | Known | Screen layout |
| 0x00777B71 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00777BD2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00777C55 | `LockediPod_Screen` | Known | Screen layout |
| 0x00777CDD | `Lock_Screen` | Known | Screen layout |
| 0x00777CEC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00777D4F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00777DB1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00777DCD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00777E3F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00777E5E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00777EC6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00777EE0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00777F48 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00777F65 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00777FD1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077803B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00778055 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007780C5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00778138 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007781A9 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00778218 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00778284 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077829F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00778314 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077837B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007783DD | `Photos_Screen` | Known | Screen layout |
| 0x00778441 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077845F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007784D1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007784EE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00778554 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077856F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007785D8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007785F5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077866C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00778690 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007786FE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00778719 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007787D4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007787F0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077885E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077887B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007788E6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00778906 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077897D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00778999 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00778A09 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00778A28 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00778A94 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00778AA8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00778B21 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00778B95 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00778C05 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00778C6D | `NoContent_Screen` | Known | Screen layout |
| 0x00778C81 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00778CE5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00778D4C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00778D66 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00778DD4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00778E46 | `NoContent_Screen` | Known | Screen layout |
| 0x00778E5A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00778EC4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00778F2D | `No_Photos_Screen` | Known | Screen layout |
| 0x00778F41 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00778FA7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00779015 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00779082 | `NoContent_Screen` | Known | Screen layout |
| 0x00779096 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007790FE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00779168 | `NoContent_Screen` | Known | Screen layout |
| 0x0077917C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007791E3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077924D | `NoContent_Screen` | Known | Screen layout |
| 0x00779261 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007792CE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00779340 | `NoContent_Screen` | Known | Screen layout |
| 0x00779354 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007793BC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00779425 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00779440 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007794A6 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007794C2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007795A1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007795BA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077961B | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077962F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077979D | `Radio_Screen` | Known | Screen layout |
| 0x007797AD | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077980E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00779891 | `LockediPod_Screen` | Known | Screen layout |
| 0x00779919 | `Lock_Screen` | Known | Screen layout |
| 0x00779928 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077998B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007799ED | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00779A09 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00779A7B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00779A9A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00779B02 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00779B1C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00779B84 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00779BA1 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00779C0D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00779C77 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00779C91 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00779D01 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00779D74 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00779DE5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00779E54 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00779EC0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00779EDB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00779F50 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00779FB7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077A019 | `Photos_Screen` | Known | Screen layout |
| 0x0077A07D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077A09B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077A10D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077A12A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077A190 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077A1AB | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077A214 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077A231 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077A2A8 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077A2CC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077A33A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077A355 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077A410 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077A42C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077A49A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077A4B7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077A522 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077A542 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077A5B9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077A5D5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077A645 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077A664 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077A6D0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077A6E4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077A75D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077A7D1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077A841 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077A8A9 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A8BD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077A921 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077A988 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077A9A2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077AA10 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077AA82 | `NoContent_Screen` | Known | Screen layout |
| 0x0077AA96 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077AB00 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077AB69 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077AB7D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077ABE3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077AC51 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077ACBE | `NoContent_Screen` | Known | Screen layout |
| 0x0077ACD2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077AD3A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077ADA4 | `NoContent_Screen` | Known | Screen layout |
| 0x0077ADB8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077AE1F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077AE89 | `NoContent_Screen` | Known | Screen layout |
| 0x0077AE9D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077AF0A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077AF7C | `NoContent_Screen` | Known | Screen layout |
| 0x0077AF90 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077AFF8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077B061 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077B07C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077B0E2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077B0FE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077B1DD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077B1F6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077B257 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077B26B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077B3D9 | `Radio_Screen` | Known | Screen layout |
| 0x0077B3E9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077B44A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077B4CD | `LockediPod_Screen` | Known | Screen layout |
| 0x0077B555 | `Lock_Screen` | Known | Screen layout |
| 0x0077B564 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077B5C7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077B629 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077B645 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077B6B7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077B6D6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077B73E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077B758 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077B7C0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077B7DD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077B849 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077B8B3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077B8CD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077B93D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077B9B0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077BA21 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077BA90 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077BAFC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077BB17 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077BB8C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077BBF3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077BC55 | `Photos_Screen` | Known | Screen layout |
| 0x0077BCB9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077BCD7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077BD49 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077BD66 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077BDCC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077BDE7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077BE50 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077BE6D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077BEE4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077BF08 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077BF76 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077BF91 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077C04C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077C068 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077C0D6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077C0F3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077C15E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077C17E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077C1F5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077C211 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077C281 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077C2A0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077C30C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077C320 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077C399 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077C40D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077C47D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077C4E5 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C4F9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077C55D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077C5C4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077C5DE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077C64C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077C6BE | `NoContent_Screen` | Known | Screen layout |
| 0x0077C6D2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077C73C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077C7A5 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077C7B9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077C81F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077C88D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077C8FA | `NoContent_Screen` | Known | Screen layout |
| 0x0077C90E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077C976 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077C9E0 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C9F4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077CA5B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077CAC5 | `NoContent_Screen` | Known | Screen layout |
| 0x0077CAD9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077CB46 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077CBB8 | `NoContent_Screen` | Known | Screen layout |
| 0x0077CBCC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077CC34 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077CC9D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077CCB8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077CD1E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077CD3A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077CE19 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077CE32 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077CE93 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077CEA7 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077D015 | `Radio_Screen` | Known | Screen layout |
| 0x0077D025 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077D086 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077D109 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077D191 | `Lock_Screen` | Known | Screen layout |
| 0x0077D1A0 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077D203 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077D265 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077D281 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077D2F3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077D312 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077D37A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077D394 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077D3FC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077D419 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077D485 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077D4EF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077D509 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077D579 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077D5EC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077D65D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077D6CC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077D738 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077D753 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077D7C8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077D82F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077D891 | `Photos_Screen` | Known | Screen layout |
| 0x0077D8F5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077D913 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077D985 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077D9A2 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077DA08 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077DA23 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077DA8C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077DAA9 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077DB20 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077DB44 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077DBB2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077DBCD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077DC88 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077DCA4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077DD12 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077DD2F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077DD9A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077DDBA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077DE31 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077DE4D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077DEBD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077DEDC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077DF48 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077DF5C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077DFD5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077E049 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077E0B9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077E121 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E135 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077E199 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077E200 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077E21A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077E288 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077E2FA | `NoContent_Screen` | Known | Screen layout |
| 0x0077E30E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077E378 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077E3E1 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077E3F5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077E45B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077E4C9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077E536 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E54A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077E5B2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077E61C | `NoContent_Screen` | Known | Screen layout |
| 0x0077E630 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077E697 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077E701 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E715 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077E782 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077E7F4 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E808 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077E870 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077E8D9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077E8F4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077E95A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077E976 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077EA55 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077EA6E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077EACF | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077EAE3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077EC51 | `Radio_Screen` | Known | Screen layout |
| 0x0077EC61 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077ECC2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077ED45 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077EDCD | `Lock_Screen` | Known | Screen layout |
| 0x0077EDDC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077EE3F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077EEA1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077EEBD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077EF2F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077EF4E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077EFB6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077EFD0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077F038 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077F055 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077F0C1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077F12B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077F145 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077F1B5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077F228 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077F299 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077F308 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077F374 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077F38F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077F404 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077F46B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077F4CD | `Photos_Screen` | Known | Screen layout |
| 0x0077F531 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077F54F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077F5C1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077F5DE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077F644 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077F65F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077F6C8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077F6E5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077F75C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077F780 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077F7EE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077F809 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077F8C4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077F8E0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077F94E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077F96B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077F9D6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077F9F6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077FA6D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077FA89 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077FAF9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077FB18 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077FB84 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077FB98 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077FC11 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077FC85 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077FCF5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077FD5D | `NoContent_Screen` | Known | Screen layout |
| 0x0077FD71 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077FDD5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077FE3C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077FE56 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077FEC4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077FF36 | `NoContent_Screen` | Known | Screen layout |
| 0x0077FF4A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077FFB4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078001D | `No_Photos_Screen` | Known | Screen layout |
| 0x00780031 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00780097 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00780105 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00780172 | `NoContent_Screen` | Known | Screen layout |
| 0x00780186 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007801EE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00780258 | `NoContent_Screen` | Known | Screen layout |
| 0x0078026C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007802D3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078033D | `NoContent_Screen` | Known | Screen layout |
| 0x00780351 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007803BE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00780430 | `NoContent_Screen` | Known | Screen layout |
| 0x00780444 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007804AC | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00780515 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00780530 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00780596 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007805B2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00780691 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007806AA | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078070B | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078071F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078088D | `Radio_Screen` | Known | Screen layout |
| 0x0078089D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007808FE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00780981 | `LockediPod_Screen` | Known | Screen layout |
| 0x00780A09 | `Lock_Screen` | Known | Screen layout |
| 0x00780A18 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00780A7B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00780ADD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00780AF9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00780B6B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00780B8A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00780BF2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00780C0C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00780C74 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00780C91 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00780CFD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00780D67 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00780D81 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00780DF1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00780E64 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00780ED5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00780F44 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00780FB0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00780FCB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00781040 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007810A7 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00781109 | `Photos_Screen` | Known | Screen layout |
| 0x0078116D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078118B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007811FD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078121A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00781280 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078129B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00781304 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00781321 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00781398 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007813BC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078142A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00781445 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00781500 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078151C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078158A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007815A7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00781612 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00781632 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007816A9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007816C5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00781735 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00781754 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007817C0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007817D4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078184D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007818C1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00781931 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00781999 | `NoContent_Screen` | Known | Screen layout |
| 0x007819AD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00781A11 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00781A78 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00781A92 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00781B00 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00781B72 | `NoContent_Screen` | Known | Screen layout |
| 0x00781B86 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00781BF0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00781C59 | `No_Photos_Screen` | Known | Screen layout |
| 0x00781C6D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00781CD3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00781D41 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00781DAE | `NoContent_Screen` | Known | Screen layout |
| 0x00781DC2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00781E2A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00781E94 | `NoContent_Screen` | Known | Screen layout |
| 0x00781EA8 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00781F0F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00781F79 | `NoContent_Screen` | Known | Screen layout |
| 0x00781F8D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00781FFA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078206C | `NoContent_Screen` | Known | Screen layout |
| 0x00782080 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007820E8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00782151 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078216C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007821D2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007821EE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007822CD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007822E6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00782347 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078235B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007824C9 | `Radio_Screen` | Known | Screen layout |
| 0x007824D9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078253A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007825BD | `LockediPod_Screen` | Known | Screen layout |
| 0x00782645 | `Lock_Screen` | Known | Screen layout |
| 0x00782654 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007826B7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00782719 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00782735 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007827A7 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007827C6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078282E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00782848 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007828B0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007828CD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00782939 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007829A3 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007829BD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00782A2D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00782AA0 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00782B11 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00782B80 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00782BEC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00782C07 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00782C7C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00782CE3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00782D45 | `Photos_Screen` | Known | Screen layout |
| 0x00782DA9 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00782DC7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00782E39 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00782E56 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00782EBC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00782ED7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00782F40 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00782F5D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00782FD4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00782FF8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00783066 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00783081 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078313C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00783158 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007831C6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007831E3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078324E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078326E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007832E5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00783301 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00783371 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00783390 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007833FC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00783410 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00783489 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007834FD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078356D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007835D5 | `NoContent_Screen` | Known | Screen layout |
| 0x007835E9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078364D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007836B4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007836CE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078373C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007837AE | `NoContent_Screen` | Known | Screen layout |
| 0x007837C2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078382C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00783895 | `No_Photos_Screen` | Known | Screen layout |
| 0x007838A9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078390F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078397D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007839EA | `NoContent_Screen` | Known | Screen layout |
| 0x007839FE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00783A66 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00783AD0 | `NoContent_Screen` | Known | Screen layout |
| 0x00783AE4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00783B4B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00783BB5 | `NoContent_Screen` | Known | Screen layout |
| 0x00783BC9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00783C36 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00783CA8 | `NoContent_Screen` | Known | Screen layout |
| 0x00783CBC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00783D24 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00783D8D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00783DA8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00783E0E | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00783E2A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00783F09 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00783F22 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00783F83 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00783F97 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00784105 | `Radio_Screen` | Known | Screen layout |
| 0x00784115 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00784176 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007841F9 | `LockediPod_Screen` | Known | Screen layout |
| 0x00784281 | `Lock_Screen` | Known | Screen layout |
| 0x00784290 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007842F3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00784355 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00784371 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007843E3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00784402 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078446A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00784484 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007844EC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00784509 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00784575 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007845DF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007845F9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00784669 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007846DC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078474D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007847BC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00784828 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00784843 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007848B8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078491F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00784981 | `Photos_Screen` | Known | Screen layout |
| 0x007849E5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00784A03 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00784A75 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00784A92 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00784AF8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00784B13 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00784B7C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00784B99 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00784C10 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00784C34 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00784CA2 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00784CBD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00784D78 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00784D94 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00784E02 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00784E1F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00784E8A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00784EAA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00784F21 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00784F3D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00784FAD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00784FCC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00785038 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078504C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007850C5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00785139 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007851A9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00785211 | `NoContent_Screen` | Known | Screen layout |
| 0x00785225 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00785289 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007852F0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078530A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00785378 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007853EA | `NoContent_Screen` | Known | Screen layout |
| 0x007853FE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00785468 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007854D1 | `No_Photos_Screen` | Known | Screen layout |
| 0x007854E5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078554B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007855B9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00785626 | `NoContent_Screen` | Known | Screen layout |
| 0x0078563A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007856A2 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078570C | `NoContent_Screen` | Known | Screen layout |
| 0x00785720 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00785787 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007857F1 | `NoContent_Screen` | Known | Screen layout |
| 0x00785805 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00785872 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007858E4 | `NoContent_Screen` | Known | Screen layout |
| 0x007858F8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00785960 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007859C9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007859E4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00785A4A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00785A66 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00785B45 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00785B5E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00785BBF | `FirstBoot_Screen` | Known | Screen layout |
| 0x00785BD3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00785D41 | `Radio_Screen` | Known | Screen layout |
| 0x00785D51 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00785DB2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00785E35 | `LockediPod_Screen` | Known | Screen layout |
| 0x00785EBD | `Lock_Screen` | Known | Screen layout |
| 0x00785ECC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00785F2F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00785F91 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00785FAD | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078601F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078603E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007860A6 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007860C0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00786128 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00786145 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007861B1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078621B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00786235 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007862A5 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00786318 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00786389 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007863F8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00786464 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078647F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007864F4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078655B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007865BD | `Photos_Screen` | Known | Screen layout |
| 0x00786621 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078663F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007866B1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007866CE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00786734 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078674F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007867B8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007867D5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078684C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00786870 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007868DE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007868F9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007869B4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007869D0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00786A3E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00786A5B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00786AC6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00786AE6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00786B5D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00786B79 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00786BE9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00786C08 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00786C74 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00786C88 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00786D01 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00786D75 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00786DE5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00786E4D | `NoContent_Screen` | Known | Screen layout |
| 0x00786E61 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00786EC5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00786F2C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00786F46 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00786FB4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00787026 | `NoContent_Screen` | Known | Screen layout |
| 0x0078703A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007870A4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078710D | `No_Photos_Screen` | Known | Screen layout |
| 0x00787121 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00787187 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007871F5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00787262 | `NoContent_Screen` | Known | Screen layout |
| 0x00787276 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007872DE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00787348 | `NoContent_Screen` | Known | Screen layout |
| 0x0078735C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007873C3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078742D | `NoContent_Screen` | Known | Screen layout |
| 0x00787441 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007874AE | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00787520 | `NoContent_Screen` | Known | Screen layout |
| 0x00787534 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078759C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00787605 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00787620 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00787686 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007876A2 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00787781 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078779A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007877FB | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078780F | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078797D | `Radio_Screen` | Known | Screen layout |
| 0x0078798D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007879EE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00787A71 | `LockediPod_Screen` | Known | Screen layout |
| 0x00787AF9 | `Lock_Screen` | Known | Screen layout |
| 0x00787B08 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00787B6B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00787BCD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00787BE9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00787C5B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00787C7A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00787CE2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00787CFC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00787D64 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00787D81 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00787DED | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00787E57 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00787E71 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00787EE1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00787F54 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00787FC5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00788034 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007880A0 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007880BB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00788130 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00788197 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007881F9 | `Photos_Screen` | Known | Screen layout |
| 0x0078825D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078827B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007882ED | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078830A | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00788370 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078838B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007883F4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00788411 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00788488 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007884AC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078851A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00788535 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007885F0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078860C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078867A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00788697 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00788702 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00788722 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00788799 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007887B5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00788825 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00788844 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007888B0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007888C4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078893D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007889B1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00788A21 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00788A89 | `NoContent_Screen` | Known | Screen layout |
| 0x00788A9D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00788B01 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00788B68 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00788B82 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00788BF0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00788C62 | `NoContent_Screen` | Known | Screen layout |
| 0x00788C76 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00788CE0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00788D49 | `No_Photos_Screen` | Known | Screen layout |
| 0x00788D5D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00788DC3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00788E31 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00788E9E | `NoContent_Screen` | Known | Screen layout |
| 0x00788EB2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00788F1A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00788F84 | `NoContent_Screen` | Known | Screen layout |
| 0x00788F98 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00788FFF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00789069 | `NoContent_Screen` | Known | Screen layout |
| 0x0078907D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007890EA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078915C | `NoContent_Screen` | Known | Screen layout |
| 0x00789170 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007891D8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00789241 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078925C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007892C2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007892DE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007893BD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007893D6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00789437 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078944B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007895B9 | `Radio_Screen` | Known | Screen layout |
| 0x007895C9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078962A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007896AD | `LockediPod_Screen` | Known | Screen layout |
| 0x00789735 | `Lock_Screen` | Known | Screen layout |
| 0x00789744 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007897A7 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00789809 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00789825 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00789897 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007898B6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078991E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00789938 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007899A0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007899BD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00789A29 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00789A93 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00789AAD | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00789B1D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00789B90 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00789C01 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00789C70 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00789CDC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00789CF7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00789D6C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00789DD3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00789E35 | `Photos_Screen` | Known | Screen layout |
| 0x00789E99 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00789EB7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00789F29 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00789F46 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00789FAC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00789FC7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078A030 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078A04D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078A0C4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078A0E8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078A156 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078A171 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078A22C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078A248 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078A2B6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078A2D3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078A33E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078A35E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078A3D5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078A3F1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078A461 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078A480 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078A4EC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078A500 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078A579 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078A5ED | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078A65D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078A6C5 | `NoContent_Screen` | Known | Screen layout |
| 0x0078A6D9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078A73D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078A7A4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078A7BE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078A82C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078A89E | `NoContent_Screen` | Known | Screen layout |
| 0x0078A8B2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078A91C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078A985 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078A999 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078A9FF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078AA6D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078AADA | `NoContent_Screen` | Known | Screen layout |
| 0x0078AAEE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078AB56 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078ABC0 | `NoContent_Screen` | Known | Screen layout |
| 0x0078ABD4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078AC3B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078ACA5 | `NoContent_Screen` | Known | Screen layout |
| 0x0078ACB9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078AD26 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078AD98 | `NoContent_Screen` | Known | Screen layout |
| 0x0078ADAC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078AE14 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078AE7D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078AE98 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078AEFE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078AF1A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078AFF9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078B012 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078B073 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078B087 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078B1F5 | `Radio_Screen` | Known | Screen layout |
| 0x0078B205 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078B266 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078B2E9 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078B371 | `Lock_Screen` | Known | Screen layout |
| 0x0078B380 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078B3E3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078B445 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078B461 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078B4D3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078B4F2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078B55A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078B574 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078B5DC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078B5F9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078B665 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078B6CF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078B6E9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078B759 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078B7CC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078B83D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078B8AC | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078B918 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078B933 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078B9A8 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078BA0F | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078BA71 | `Photos_Screen` | Known | Screen layout |
| 0x0078BAD5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078BAF3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078BB65 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078BB82 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078BBE8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078BC03 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078BC6C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078BC89 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078BD00 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078BD24 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078BD92 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078BDAD | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078BE68 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078BE84 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078BEF2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078BF0F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078BF7A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078BF9A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078C011 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078C02D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078C09D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078C0BC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078C128 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078C13C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078C1B5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078C229 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078C299 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078C301 | `NoContent_Screen` | Known | Screen layout |
| 0x0078C315 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078C379 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078C3E0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078C3FA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078C468 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078C4DA | `NoContent_Screen` | Known | Screen layout |
| 0x0078C4EE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078C558 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078C5C1 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078C5D5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078C63B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078C6A9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078C716 | `NoContent_Screen` | Known | Screen layout |
| 0x0078C72A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078C792 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078C7FC | `NoContent_Screen` | Known | Screen layout |
| 0x0078C810 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078C877 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078C8E1 | `NoContent_Screen` | Known | Screen layout |
| 0x0078C8F5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078C962 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078C9D4 | `NoContent_Screen` | Known | Screen layout |
| 0x0078C9E8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078CA50 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078CAB9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078CAD4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078CB3A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078CB56 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078CC35 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078CC4E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078CCAF | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078CCC3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078CE31 | `Radio_Screen` | Known | Screen layout |
| 0x0078CE41 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078CEA2 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078CF25 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078CFAD | `Lock_Screen` | Known | Screen layout |
| 0x0078CFBC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078D01F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078D081 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078D09D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078D10F | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078D12E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078D196 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078D1B0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078D218 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078D235 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078D2A1 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078D30B | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078D325 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078D395 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078D408 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078D479 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078D4E8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078D554 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078D56F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078D5E4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078D64B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078D6AD | `Photos_Screen` | Known | Screen layout |
| 0x0078D711 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078D72F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078D7A1 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078D7BE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078D824 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078D83F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078D8A8 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078D8C5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078D93C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078D960 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078D9CE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078D9E9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078DAA4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078DAC0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078DB2E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078DB4B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078DBB6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078DBD6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078DC4D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078DC69 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078DCD9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078DCF8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078DD64 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078DD78 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078DDF1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078DE65 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078DED5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078DF3D | `NoContent_Screen` | Known | Screen layout |
| 0x0078DF51 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078DFB5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078E01C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078E036 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078E0A4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078E116 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E12A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078E194 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078E1FD | `No_Photos_Screen` | Known | Screen layout |
| 0x0078E211 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078E277 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078E2E5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078E352 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E366 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078E3CE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078E438 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E44C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0078E4B3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078E51D | `NoContent_Screen` | Known | Screen layout |
| 0x0078E531 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078E59E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078E610 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E624 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078E68C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078E6F5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078E710 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078E776 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078E792 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078E871 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078E88A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078E8EB | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078E8FF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078EA6D | `Radio_Screen` | Known | Screen layout |
| 0x0078EA7D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078EADE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078EB61 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078EBE9 | `Lock_Screen` | Known | Screen layout |
| 0x0078EBF8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078EC5B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078ECBD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078ECD9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078ED4B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078ED6A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078EDD2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078EDEC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078EE54 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078EE71 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078EEDD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078EF47 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078EF61 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078EFD1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078F044 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078F0B5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078F124 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078F190 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078F1AB | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078F220 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078F287 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078F2E9 | `Photos_Screen` | Known | Screen layout |
| 0x0078F34D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078F36B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078F3DD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0078F3FA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078F460 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078F47B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078F4E4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078F501 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078F578 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078F59C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078F60A | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078F625 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078F6E0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078F6FC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078F76A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078F787 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078F7F2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078F812 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078F889 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078F8A5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078F915 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078F934 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078F9A0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078F9B4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078FA2D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078FAA1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078FB11 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078FB79 | `NoContent_Screen` | Known | Screen layout |
| 0x0078FB8D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078FBF1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078FC58 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078FC72 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078FCE0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078FD52 | `NoContent_Screen` | Known | Screen layout |
| 0x0078FD66 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078FDD0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078FE39 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078FE4D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078FEB3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078FF21 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078FF8E | `NoContent_Screen` | Known | Screen layout |
| 0x0078FFA2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079000A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00790074 | `NoContent_Screen` | Known | Screen layout |
| 0x00790088 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007900EF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00790159 | `NoContent_Screen` | Known | Screen layout |
| 0x0079016D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007901DA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079024C | `NoContent_Screen` | Known | Screen layout |
| 0x00790260 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007902C8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00790331 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0079034C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007903B2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007903CE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007904AD | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007904C6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00790527 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079053B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007906A9 | `Radio_Screen` | Known | Screen layout |
| 0x007906B9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079071A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079079D | `LockediPod_Screen` | Known | Screen layout |
| 0x00790825 | `Lock_Screen` | Known | Screen layout |
| 0x00790834 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00790897 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007908F9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00790915 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00790987 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007909A6 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00790A0E | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00790A28 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00790A90 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00790AAD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00790B19 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00790B83 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00790B9D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00790C0D | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00790C80 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00790CF1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00790D60 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00790DCC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00790DE7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00790E5C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00790EC3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00790F25 | `Photos_Screen` | Known | Screen layout |
| 0x00790F89 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00790FA7 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00791019 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00791036 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079109C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007910B7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00791120 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079113D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007911B4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007911D8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00791246 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00791261 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079131C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00791338 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007913A6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007913C3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079142E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079144E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007914C5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007914E1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00791551 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00791570 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007915DC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007915F0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00791669 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007916DD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079174D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007917B5 | `NoContent_Screen` | Known | Screen layout |
| 0x007917C9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079182D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00791894 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007918AE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079191C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079198E | `NoContent_Screen` | Known | Screen layout |
| 0x007919A2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00791A0C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00791A75 | `No_Photos_Screen` | Known | Screen layout |
| 0x00791A89 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00791AEF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00791B5D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00791BCA | `NoContent_Screen` | Known | Screen layout |
| 0x00791BDE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00791C46 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00791CB0 | `NoContent_Screen` | Known | Screen layout |
| 0x00791CC4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00791D2B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00791D95 | `NoContent_Screen` | Known | Screen layout |
| 0x00791DA9 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00791E16 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00791E88 | `NoContent_Screen` | Known | Screen layout |
| 0x00791E9C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00791F04 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00791F6D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00791F88 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00791FEE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079200A | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007920E9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00792102 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00792163 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00792177 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007922E5 | `Radio_Screen` | Known | Screen layout |
| 0x007922F5 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00792356 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007923D9 | `LockediPod_Screen` | Known | Screen layout |
| 0x00792461 | `Lock_Screen` | Known | Screen layout |
| 0x00792470 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007924D3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00792535 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00792551 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007925C3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007925E2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079264A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00792664 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007926CC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007926E9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00792755 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007927BF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007927D9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00792849 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007928BC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079292D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079299C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00792A08 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00792A23 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00792A98 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00792AFF | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00792B61 | `Photos_Screen` | Known | Screen layout |
| 0x00792BC5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00792BE3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00792C55 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00792C72 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00792CD8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00792CF3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00792D5C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00792D79 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00792DF0 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00792E14 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00792E82 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00792E9D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00792F58 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00792F74 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00792FE2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00792FFF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079306A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079308A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00793101 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079311D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079318D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007931AC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00793218 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079322C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007932A5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00793319 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00793389 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007933F1 | `NoContent_Screen` | Known | Screen layout |
| 0x00793405 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00793469 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007934D0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007934EA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00793558 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007935CA | `NoContent_Screen` | Known | Screen layout |
| 0x007935DE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00793648 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007936B1 | `No_Photos_Screen` | Known | Screen layout |
| 0x007936C5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079372B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00793799 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00793806 | `NoContent_Screen` | Known | Screen layout |
| 0x0079381A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00793882 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007938EC | `NoContent_Screen` | Known | Screen layout |
| 0x00793900 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00793967 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007939D1 | `NoContent_Screen` | Known | Screen layout |
| 0x007939E5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00793A52 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00793AC4 | `NoContent_Screen` | Known | Screen layout |
| 0x00793AD8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00793B40 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00793BA9 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00793BC4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00793C2A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00793C46 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00793D25 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00793D3E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00793D9F | `FirstBoot_Screen` | Known | Screen layout |
| 0x00793DB3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00793F21 | `Radio_Screen` | Known | Screen layout |
| 0x00793F31 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00793F92 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00794015 | `LockediPod_Screen` | Known | Screen layout |
| 0x0079409D | `Lock_Screen` | Known | Screen layout |
| 0x007940AC | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079410F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00794171 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079418D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007941FF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079421E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00794286 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007942A0 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00794308 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00794325 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00794391 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007943FB | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00794415 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00794485 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007944F8 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00794569 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007945D8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00794644 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079465F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007946D4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079473B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079479D | `Photos_Screen` | Known | Screen layout |
| 0x00794801 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079481F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00794891 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007948AE | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00794914 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079492F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00794998 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007949B5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00794A2C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00794A50 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00794ABE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00794AD9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00794B94 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00794BB0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00794C1E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00794C3B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00794CA6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00794CC6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00794D3D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00794D59 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00794DC9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00794DE8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00794E54 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00794E68 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00794EE1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00794F55 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00794FC5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079502D | `NoContent_Screen` | Known | Screen layout |
| 0x00795041 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007950A5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0079510C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00795126 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00795194 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00795206 | `NoContent_Screen` | Known | Screen layout |
| 0x0079521A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00795284 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007952ED | `No_Photos_Screen` | Known | Screen layout |
| 0x00795301 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00795367 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007953D5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00795442 | `NoContent_Screen` | Known | Screen layout |
| 0x00795456 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007954BE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00795528 | `NoContent_Screen` | Known | Screen layout |
| 0x0079553C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007955A3 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079560D | `NoContent_Screen` | Known | Screen layout |
| 0x00795621 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079568E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00795700 | `NoContent_Screen` | Known | Screen layout |
| 0x00795714 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079577C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007957E5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00795800 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00795866 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00795882 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00795961 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079597A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007959DB | `FirstBoot_Screen` | Known | Screen layout |
| 0x007959EF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00795B5D | `Radio_Screen` | Known | Screen layout |
| 0x00795B6D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00795BCE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00795C51 | `LockediPod_Screen` | Known | Screen layout |
| 0x00795CD9 | `Lock_Screen` | Known | Screen layout |
| 0x00795CE8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00795D4B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00795DAD | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00795DC9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00795E3B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00795E5A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00795EC2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00795EDC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00795F44 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00795F61 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00795FCD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00796037 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00796051 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007960C1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00796134 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007961A5 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00796214 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00796280 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079629B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00796310 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00796377 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007963D9 | `Photos_Screen` | Known | Screen layout |
| 0x0079643D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079645B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007964CD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007964EA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00796550 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079656B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007965D4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007965F1 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00796668 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079668C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007966FA | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00796715 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007967D0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007967EC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079685A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00796877 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007968E2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00796902 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00796979 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00796995 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00796A05 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00796A24 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00796A90 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00796AA4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00796B1D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00796B91 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00796C01 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00796C69 | `NoContent_Screen` | Known | Screen layout |
| 0x00796C7D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00796CE1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00796D48 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00796D62 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00796DD0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00796E42 | `NoContent_Screen` | Known | Screen layout |
| 0x00796E56 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00796EC0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00796F29 | `No_Photos_Screen` | Known | Screen layout |
| 0x00796F3D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00796FA3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00797011 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0079707E | `NoContent_Screen` | Known | Screen layout |
| 0x00797092 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007970FA | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00797164 | `NoContent_Screen` | Known | Screen layout |
| 0x00797178 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007971DF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00797249 | `NoContent_Screen` | Known | Screen layout |
| 0x0079725D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007972CA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079733C | `NoContent_Screen` | Known | Screen layout |
| 0x00797350 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007973B8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00797421 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0079743C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007974A2 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007974BE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079759D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007975B6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00797617 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079762B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00797799 | `Radio_Screen` | Known | Screen layout |
| 0x007977A9 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079780A | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079788D | `LockediPod_Screen` | Known | Screen layout |
| 0x00797915 | `Lock_Screen` | Known | Screen layout |
| 0x00797924 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00797987 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007979E9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00797A05 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00797A77 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00797A96 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00797AFE | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00797B18 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00797B80 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00797B9D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00797C09 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00797C73 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00797C8D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00797CFD | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00797D70 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00797DE1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00797E50 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00797EBC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00797ED7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00797F4C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00797FB3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00798015 | `Photos_Screen` | Known | Screen layout |
| 0x00798079 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00798097 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00798109 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00798126 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079818C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007981A7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00798210 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079822D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007982A4 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007982C8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00798336 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00798351 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079840C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00798428 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00798496 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007984B3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079851E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079853E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007985B5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007985D1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00798641 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00798660 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007986CC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007986E0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00798759 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007987CD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079883D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007988A5 | `NoContent_Screen` | Known | Screen layout |
| 0x007988B9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079891D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00798984 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079899E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00798A0C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00798A7E | `NoContent_Screen` | Known | Screen layout |
| 0x00798A92 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00798AFC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00798B65 | `No_Photos_Screen` | Known | Screen layout |
| 0x00798B79 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00798BDF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00798C4D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00798CBA | `NoContent_Screen` | Known | Screen layout |
| 0x00798CCE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00798D36 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00798DA0 | `NoContent_Screen` | Known | Screen layout |
| 0x00798DB4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00798E1B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00798E85 | `NoContent_Screen` | Known | Screen layout |
| 0x00798E99 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00798F06 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00798F78 | `NoContent_Screen` | Known | Screen layout |
| 0x00798F8C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00798FF4 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0079905D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00799078 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007990DE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007990FA | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007991D9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007991F2 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00799253 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00799267 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007993D5 | `Radio_Screen` | Known | Screen layout |
| 0x007993E5 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00799446 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007994C9 | `LockediPod_Screen` | Known | Screen layout |
| 0x00799551 | `Lock_Screen` | Known | Screen layout |
| 0x00799560 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007995C3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00799625 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00799641 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007996B3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007996D2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079973A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00799754 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007997BC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007997D9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00799845 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007998AF | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007998C9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00799939 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007999AC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00799A1D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00799A8C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00799AF8 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00799B13 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00799B88 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00799BEF | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00799C51 | `Photos_Screen` | Known | Screen layout |
| 0x00799CB5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00799CD3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00799D45 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00799D62 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00799DC8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00799DE3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00799E4C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00799E69 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00799EE0 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00799F04 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00799F72 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00799F8D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079A048 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079A064 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079A0D2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079A0EF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079A15A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079A17A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079A1F1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079A20D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079A27D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079A29C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079A308 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079A31C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0079A395 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079A409 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079A479 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079A4E1 | `NoContent_Screen` | Known | Screen layout |
| 0x0079A4F5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079A559 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0079A5C0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079A5DA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079A648 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079A6BA | `NoContent_Screen` | Known | Screen layout |
| 0x0079A6CE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079A738 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0079A7A1 | `No_Photos_Screen` | Known | Screen layout |
| 0x0079A7B5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079A81B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079A889 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0079A8F6 | `NoContent_Screen` | Known | Screen layout |
| 0x0079A90A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079A972 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0079A9DC | `NoContent_Screen` | Known | Screen layout |
| 0x0079A9F0 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079AA57 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079AAC1 | `NoContent_Screen` | Known | Screen layout |
| 0x0079AAD5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079AB42 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079ABB4 | `NoContent_Screen` | Known | Screen layout |
| 0x0079ABC8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079AC30 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0079AC99 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0079ACB4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079AD1A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079AD36 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079AE15 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079AE2E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0079AE8F | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079AEA3 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0079B011 | `Radio_Screen` | Known | Screen layout |
| 0x0079B021 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079B082 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079B105 | `LockediPod_Screen` | Known | Screen layout |
| 0x0079B18D | `Lock_Screen` | Known | Screen layout |
| 0x0079B19C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079B1FF | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0079B261 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079B27D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0079B2EF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079B30E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079B376 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079B390 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079B3F8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079B415 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079B481 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079B4EB | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0079B505 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0079B575 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079B5E8 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079B659 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079B6C8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079B734 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079B74F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0079B7C4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079B82B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079B88D | `Photos_Screen` | Known | Screen layout |
| 0x0079B8F1 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079B90F | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0079B981 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079B99E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079BA04 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079BA1F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079BA88 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079BAA5 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079BB1C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079BB40 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079BBAE | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0079BBC9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079BC84 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079BCA0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079BD0E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079BD2B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079BD96 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079BDB6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079BE2D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079BE49 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079BEB9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079BED8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079BF44 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079BF58 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0079BFD1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079C045 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079C0B5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079C11D | `NoContent_Screen` | Known | Screen layout |
| 0x0079C131 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079C195 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0079C1FC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079C216 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079C284 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079C2F6 | `NoContent_Screen` | Known | Screen layout |
| 0x0079C30A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079C374 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0079C3DD | `No_Photos_Screen` | Known | Screen layout |
| 0x0079C3F1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079C457 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079C4C5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0079C532 | `NoContent_Screen` | Known | Screen layout |
| 0x0079C546 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079C5AE | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0079C618 | `NoContent_Screen` | Known | Screen layout |
| 0x0079C62C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079C693 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079C6FD | `NoContent_Screen` | Known | Screen layout |
| 0x0079C711 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079C77E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079C7F0 | `NoContent_Screen` | Known | Screen layout |
| 0x0079C804 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079C86C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0079C8D5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0079C8F0 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079C956 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079C972 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079CA51 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079CA6A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0079CACB | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079CADF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0079CC4D | `Radio_Screen` | Known | Screen layout |
| 0x0079CC5D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079CCBE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079CD41 | `LockediPod_Screen` | Known | Screen layout |
| 0x0079CDC9 | `Lock_Screen` | Known | Screen layout |
| 0x0079CDD8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079CE3B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0079CE9D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079CEB9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0079CF2B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079CF4A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079CFB2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079CFCC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079D034 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079D051 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079D0BD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079D127 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0079D141 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0079D1B1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079D224 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079D295 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079D304 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079D370 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079D38B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0079D400 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079D467 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079D4C9 | `Photos_Screen` | Known | Screen layout |
| 0x0079D52D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079D54B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0079D5BD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079D5DA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079D640 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079D65B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079D6C4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079D6E1 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079D758 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079D77C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079D7EA | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0079D805 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079D8C0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079D8DC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079D94A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079D967 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079D9D2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079D9F2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079DA69 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079DA85 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079DAF5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079DB14 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079DB80 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079DB94 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0079DC0D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079DC81 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079DCF1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079DD59 | `NoContent_Screen` | Known | Screen layout |
| 0x0079DD6D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079DDD1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0079DE38 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079DE52 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079DEC0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079DF32 | `NoContent_Screen` | Known | Screen layout |
| 0x0079DF46 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079DFB0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0079E019 | `No_Photos_Screen` | Known | Screen layout |
| 0x0079E02D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079E093 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079E101 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0079E16E | `NoContent_Screen` | Known | Screen layout |
| 0x0079E182 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079E1EA | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0079E254 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E268 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079E2CF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079E339 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E34D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079E3BA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079E42C | `NoContent_Screen` | Known | Screen layout |
| 0x0079E440 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079E4A8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0079E511 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0079E52C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079E592 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079E5AE | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079E68D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0079E6A6 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0079E707 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0079E71B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0079E889 | `Radio_Screen` | Known | Screen layout |
| 0x0079E899 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0079E8FA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079E97D | `LockediPod_Screen` | Known | Screen layout |
| 0x0079EA05 | `Lock_Screen` | Known | Screen layout |
| 0x0079EA14 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0079EA77 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0079EAD9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079EAF5 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0079EB67 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079EB86 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079EBEE | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079EC08 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079EC70 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079EC8D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079ECF9 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079ED63 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0079ED7D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0079EDED | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079EE60 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0079EED1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0079EF40 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079EFAC | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0079EFC7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0079F03C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079F0A3 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0079F105 | `Photos_Screen` | Known | Screen layout |
| 0x0079F169 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0079F187 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0079F1F9 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0079F216 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0079F27C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079F297 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079F300 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079F31D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079F394 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079F3B8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079F426 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0079F441 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079F4FC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079F518 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079F586 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079F5A3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079F60E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079F62E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079F6A5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079F6C1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079F731 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079F750 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079F7BC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079F7D0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0079F849 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079F8BD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0079F92D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079F995 | `NoContent_Screen` | Known | Screen layout |
| 0x0079F9A9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0079FA0D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0079FA74 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0079FA8E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0079FAFC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079FB6E | `NoContent_Screen` | Known | Screen layout |
| 0x0079FB82 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079FBEC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0079FC55 | `No_Photos_Screen` | Known | Screen layout |
| 0x0079FC69 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0079FCCF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079FD3D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0079FDAA | `NoContent_Screen` | Known | Screen layout |
| 0x0079FDBE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0079FE26 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0079FE90 | `NoContent_Screen` | Known | Screen layout |
| 0x0079FEA4 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079FF0B | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0079FF75 | `NoContent_Screen` | Known | Screen layout |
| 0x0079FF89 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079FFF6 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A0068 | `NoContent_Screen` | Known | Screen layout |
| 0x007A007C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A00E4 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A014D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A0168 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A01CE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A01EA | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A02C9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A02E2 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A0343 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A0357 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A04C5 | `Radio_Screen` | Known | Screen layout |
| 0x007A04D5 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A0536 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A05B9 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A0641 | `Lock_Screen` | Known | Screen layout |
| 0x007A0650 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A06B3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A0715 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A0731 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A07A3 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A07C2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A082A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A0844 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A08AC | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A08C9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A0935 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A099F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A09B9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A0A29 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A0A9C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A0B0D | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A0B7C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A0BE8 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A0C03 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A0C78 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A0CDF | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A0D41 | `Photos_Screen` | Known | Screen layout |
| 0x007A0DA5 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A0DC3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A0E35 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A0E52 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A0EB8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A0ED3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A0F3C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A0F59 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A0FD0 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A0FF4 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A1062 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A107D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A1138 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A1154 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A11C2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A11DF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A124A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A126A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A12E1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A12FD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A136D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A138C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A13F8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A140C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A1485 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A14F9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A1569 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A15D1 | `NoContent_Screen` | Known | Screen layout |
| 0x007A15E5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A1649 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A16B0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A16CA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A1738 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A17AA | `NoContent_Screen` | Known | Screen layout |
| 0x007A17BE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A1828 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A1891 | `No_Photos_Screen` | Known | Screen layout |
| 0x007A18A5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A190B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A1979 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A19E6 | `NoContent_Screen` | Known | Screen layout |
| 0x007A19FA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A1A62 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A1ACC | `NoContent_Screen` | Known | Screen layout |
| 0x007A1AE0 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A1B47 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A1BB1 | `NoContent_Screen` | Known | Screen layout |
| 0x007A1BC5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A1C32 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A1CA4 | `NoContent_Screen` | Known | Screen layout |
| 0x007A1CB8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A1D20 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A1D89 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A1DA4 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A1E0A | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A1E26 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A1F05 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A1F1E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A1F7F | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A1F93 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A2101 | `Radio_Screen` | Known | Screen layout |
| 0x007A2111 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A2172 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A21F5 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A227D | `Lock_Screen` | Known | Screen layout |
| 0x007A228C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A22EF | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A2351 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A236D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A23DF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A23FE | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A2466 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A2480 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A24E8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A2505 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A2571 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A25DB | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A25F5 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A2665 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A26D8 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A2749 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A27B8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A2824 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A283F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A28B4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A291B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A297D | `Photos_Screen` | Known | Screen layout |
| 0x007A29E1 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A29FF | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A2A71 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A2A8E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A2AF4 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A2B0F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A2B78 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A2B95 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A2C0C | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A2C30 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A2C9E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A2CB9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A2D74 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A2D90 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A2DFE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A2E1B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A2E86 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A2EA6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A2F1D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A2F39 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A2FA9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A2FC8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A3034 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A3048 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A30C1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A3135 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A31A5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A320D | `NoContent_Screen` | Known | Screen layout |
| 0x007A3221 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A3285 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A32EC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A3306 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A3374 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A33E6 | `NoContent_Screen` | Known | Screen layout |
| 0x007A33FA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A3464 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A34CD | `No_Photos_Screen` | Known | Screen layout |
| 0x007A34E1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A3547 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A35B5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A3622 | `NoContent_Screen` | Known | Screen layout |
| 0x007A3636 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A369E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A3708 | `NoContent_Screen` | Known | Screen layout |
| 0x007A371C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A3783 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A37ED | `NoContent_Screen` | Known | Screen layout |
| 0x007A3801 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A386E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A38E0 | `NoContent_Screen` | Known | Screen layout |
| 0x007A38F4 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A395C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A39C5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A39E0 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A3A46 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A3A62 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A3B41 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A3B5A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A3BBB | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A3BCF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A3D3D | `Radio_Screen` | Known | Screen layout |
| 0x007A3D4D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A3DAE | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A3E31 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A3EB9 | `Lock_Screen` | Known | Screen layout |
| 0x007A3EC8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A3F2B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A3F8D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A3FA9 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A401B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A403A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A40A2 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A40BC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A4124 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A4141 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A41AD | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A4217 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A4231 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A42A1 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A4314 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A4385 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A43F4 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A4460 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A447B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A44F0 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A4557 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A45B9 | `Photos_Screen` | Known | Screen layout |
| 0x007A461D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A463B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A46AD | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A46CA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A4730 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A474B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A47B4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A47D1 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A4848 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A486C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A48DA | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A48F5 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A49B0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A49CC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A4A3A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A4A57 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A4AC2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A4AE2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A4B59 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A4B75 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A4BE5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A4C04 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A4C70 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A4C84 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A4CFD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A4D71 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A4DE1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A4E49 | `NoContent_Screen` | Known | Screen layout |
| 0x007A4E5D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A4EC1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A4F28 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A4F42 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A4FB0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A5022 | `NoContent_Screen` | Known | Screen layout |
| 0x007A5036 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A50A0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A5109 | `No_Photos_Screen` | Known | Screen layout |
| 0x007A511D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A5183 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A51F1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A525E | `NoContent_Screen` | Known | Screen layout |
| 0x007A5272 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A52DA | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A5344 | `NoContent_Screen` | Known | Screen layout |
| 0x007A5358 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A53BF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A5429 | `NoContent_Screen` | Known | Screen layout |
| 0x007A543D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A54AA | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A551C | `NoContent_Screen` | Known | Screen layout |
| 0x007A5530 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A5598 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A5601 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A561C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A5682 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A569E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A577D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A5796 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A57F7 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A580B | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A5979 | `Radio_Screen` | Known | Screen layout |
| 0x007A5989 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A59EA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A5A6D | `LockediPod_Screen` | Known | Screen layout |
| 0x007A5AF5 | `Lock_Screen` | Known | Screen layout |
| 0x007A5B04 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A5B67 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A5BC9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A5BE5 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A5C57 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A5C76 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A5CDE | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A5CF8 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A5D60 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A5D7D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A5DE9 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A5E53 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A5E6D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A5EDD | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A5F50 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A5FC1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A6030 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A609C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A60B7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A612C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A6193 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A61F5 | `Photos_Screen` | Known | Screen layout |
| 0x007A6259 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A6277 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A62E9 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A6306 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A636C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A6387 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A63F0 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A640D | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A6484 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A64A8 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A6516 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A6531 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A65EC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A6608 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A6676 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A6693 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A66FE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A671E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A6795 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A67B1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A6821 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A6840 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A68AC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A68C0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A6939 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A69AD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A6A1D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A6A85 | `NoContent_Screen` | Known | Screen layout |
| 0x007A6A99 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A6AFD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A6B64 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A6B7E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A6BEC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A6C5E | `NoContent_Screen` | Known | Screen layout |
| 0x007A6C72 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A6CDC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A6D45 | `No_Photos_Screen` | Known | Screen layout |
| 0x007A6D59 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A6DBF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A6E2D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A6E9A | `NoContent_Screen` | Known | Screen layout |
| 0x007A6EAE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A6F16 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A6F80 | `NoContent_Screen` | Known | Screen layout |
| 0x007A6F94 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A6FFB | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A7065 | `NoContent_Screen` | Known | Screen layout |
| 0x007A7079 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A70E6 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A7158 | `NoContent_Screen` | Known | Screen layout |
| 0x007A716C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A71D4 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A723D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A7258 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A72BE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A72DA | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A73B9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A73D2 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A7433 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A7447 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A75B5 | `Radio_Screen` | Known | Screen layout |
| 0x007A75C5 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A7626 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A76A9 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A7731 | `Lock_Screen` | Known | Screen layout |
| 0x007A7740 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A77A3 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A7805 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A7821 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A7893 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A78B2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A791A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A7934 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A799C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A79B9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A7A25 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A7A8F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A7AA9 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A7B19 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A7B8C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A7BFD | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A7C6C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A7CD8 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A7CF3 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A7D68 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A7DCF | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A7E31 | `Photos_Screen` | Known | Screen layout |
| 0x007A7E95 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A7EB3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A7F25 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A7F42 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A7FA8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A7FC3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A802C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A8049 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A80C0 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A80E4 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A8152 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A816D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A8228 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A8244 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A82B2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A82CF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A833A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A835A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007A83D1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A83ED | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A845D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007A847C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007A84E8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007A84FC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007A8575 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007A85E9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007A8659 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007A86C1 | `NoContent_Screen` | Known | Screen layout |
| 0x007A86D5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007A8739 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007A87A0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A87BA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007A8828 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A889A | `NoContent_Screen` | Known | Screen layout |
| 0x007A88AE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A8918 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007A8981 | `No_Photos_Screen` | Known | Screen layout |
| 0x007A8995 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007A89FB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A8A69 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007A8AD6 | `NoContent_Screen` | Known | Screen layout |
| 0x007A8AEA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007A8B52 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007A8BBC | `NoContent_Screen` | Known | Screen layout |
| 0x007A8BD0 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007A8C37 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007A8CA1 | `NoContent_Screen` | Known | Screen layout |
| 0x007A8CB5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A8D22 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A8D94 | `NoContent_Screen` | Known | Screen layout |
| 0x007A8DA8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A8E10 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007A8E79 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007A8E94 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A8EFA | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007A8F16 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007A8FF5 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007A900E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007A906F | `FirstBoot_Screen` | Known | Screen layout |
| 0x007A9083 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007A91F1 | `Radio_Screen` | Known | Screen layout |
| 0x007A9201 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007A9262 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007A92E5 | `LockediPod_Screen` | Known | Screen layout |
| 0x007A936D | `Lock_Screen` | Known | Screen layout |
| 0x007A937C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007A93DF | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007A9441 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A945D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007A94CF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007A94EE | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007A9556 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007A9570 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007A95D8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A95F5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A9661 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007A96CB | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007A96E5 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007A9755 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007A97C8 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007A9839 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007A98A8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A9914 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007A992F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007A99A4 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007A9A0B | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007A9A6D | `Photos_Screen` | Known | Screen layout |
| 0x007A9AD1 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007A9AEF | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007A9B61 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007A9B7E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007A9BE4 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007A9BFF | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007A9C68 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A9C85 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A9CFC | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A9D20 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A9D8E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007A9DA9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A9E64 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007A9E80 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007A9EEE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A9F0B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A9F76 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007A9F96 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007AA00D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AA029 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AA099 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007AA0B8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007AA124 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007AA138 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007AA1B1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007AA225 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007AA295 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007AA2FD | `NoContent_Screen` | Known | Screen layout |
| 0x007AA311 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007AA375 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007AA3DC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AA3F6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007AA464 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007AA4D6 | `NoContent_Screen` | Known | Screen layout |
| 0x007AA4EA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007AA554 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007AA5BD | `No_Photos_Screen` | Known | Screen layout |
| 0x007AA5D1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007AA637 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AA6A5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007AA712 | `NoContent_Screen` | Known | Screen layout |
| 0x007AA726 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007AA78E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007AA7F8 | `NoContent_Screen` | Known | Screen layout |
| 0x007AA80C | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007AA873 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007AA8DD | `NoContent_Screen` | Known | Screen layout |
| 0x007AA8F1 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007AA95E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007AA9D0 | `NoContent_Screen` | Known | Screen layout |
| 0x007AA9E4 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AAA4C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007AAAB5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007AAAD0 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007AAB36 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007AAB52 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007AAC31 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007AAC4A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007AACAB | `FirstBoot_Screen` | Known | Screen layout |
| 0x007AACBF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007AAE2D | `Radio_Screen` | Known | Screen layout |
| 0x007AAE3D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007AAE9E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007AAF21 | `LockediPod_Screen` | Known | Screen layout |
| 0x007AAFA9 | `Lock_Screen` | Known | Screen layout |
| 0x007AAFB8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007AB01B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007AB07D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007AB099 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007AB10B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007AB12A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007AB192 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AB1AC | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007AB214 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AB231 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AB29D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007AB307 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007AB321 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007AB391 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007AB404 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007AB475 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007AB4E4 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007AB550 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007AB56B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007AB5E0 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007AB647 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007AB6A9 | `Photos_Screen` | Known | Screen layout |
| 0x007AB70D | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007AB72B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007AB79D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007AB7BA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007AB820 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007AB83B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007AB8A4 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007AB8C1 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007AB938 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007AB95C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007AB9CA | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007AB9E5 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007ABAA0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007ABABC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007ABB2A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007ABB47 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007ABBB2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007ABBD2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007ABC49 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007ABC65 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007ABCD5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007ABCF4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007ABD60 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007ABD74 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007ABDED | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007ABE61 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007ABED1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007ABF39 | `NoContent_Screen` | Known | Screen layout |
| 0x007ABF4D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007ABFB1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007AC018 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AC032 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007AC0A0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007AC112 | `NoContent_Screen` | Known | Screen layout |
| 0x007AC126 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007AC190 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007AC1F9 | `No_Photos_Screen` | Known | Screen layout |
| 0x007AC20D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007AC273 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AC2E1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007AC34E | `NoContent_Screen` | Known | Screen layout |
| 0x007AC362 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007AC3CA | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007AC434 | `NoContent_Screen` | Known | Screen layout |
| 0x007AC448 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007AC4AF | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007AC519 | `NoContent_Screen` | Known | Screen layout |
| 0x007AC52D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007AC59A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007AC60C | `NoContent_Screen` | Known | Screen layout |
| 0x007AC620 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AC688 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007AC6F1 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007AC70C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007AC772 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007AC78E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007AC86D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007AC886 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007AC8E7 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007AC8FB | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007ACA69 | `Radio_Screen` | Known | Screen layout |
| 0x007ACA79 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007ACADA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007ACB5D | `LockediPod_Screen` | Known | Screen layout |
| 0x007ACBE5 | `Lock_Screen` | Known | Screen layout |
| 0x007ACBF4 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007ACC57 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007ACCB9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007ACCD5 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007ACD47 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007ACD66 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007ACDCE | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007ACDE8 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007ACE50 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007ACE6D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007ACED9 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007ACF43 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007ACF5D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007ACFCD | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007AD040 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007AD0B1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007AD120 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007AD18C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007AD1A7 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007AD21C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007AD283 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007AD2E5 | `Photos_Screen` | Known | Screen layout |
| 0x007AD349 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007AD367 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007AD3D9 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007AD3F6 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007AD45C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007AD477 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007AD4E0 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007AD4FD | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007AD574 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007AD598 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007AD606 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007AD621 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007AD6DC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AD6F8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AD766 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AD783 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AD7EE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007AD80E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007AD885 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AD8A1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AD911 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007AD930 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007AD99C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007AD9B0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007ADA29 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007ADA9D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007ADB0D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007ADB75 | `NoContent_Screen` | Known | Screen layout |
| 0x007ADB89 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007ADBED | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007ADC54 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007ADC6E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007ADCDC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007ADD4E | `NoContent_Screen` | Known | Screen layout |
| 0x007ADD62 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007ADDCC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007ADE35 | `No_Photos_Screen` | Known | Screen layout |
| 0x007ADE49 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007ADEAF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007ADF1D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007ADF8A | `NoContent_Screen` | Known | Screen layout |
| 0x007ADF9E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007AE006 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007AE070 | `NoContent_Screen` | Known | Screen layout |
| 0x007AE084 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007AE0EB | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007AE155 | `NoContent_Screen` | Known | Screen layout |
| 0x007AE169 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007AE1D6 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007AE248 | `NoContent_Screen` | Known | Screen layout |
| 0x007AE25C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AE2C4 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007AE32D | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007AE348 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007AE3AE | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007AE3CA | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007AE4A9 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007AE4C2 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007AE523 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007AE537 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007AE6A5 | `Radio_Screen` | Known | Screen layout |
| 0x007AE6B5 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007AE716 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007AE799 | `LockediPod_Screen` | Known | Screen layout |
| 0x007AE821 | `Lock_Screen` | Known | Screen layout |
| 0x007AE830 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007AE893 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007AE8F5 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007AE911 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007AE983 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007AE9A2 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007AEA0A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AEA24 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007AEA8C | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AEAA9 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AEB15 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007AEB7F | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007AEB99 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007AEC09 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007AEC7C | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007AECED | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007AED5C | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007AEDC8 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007AEDE3 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007AEE58 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007AEEBF | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007AEF21 | `Photos_Screen` | Known | Screen layout |
| 0x007AEF85 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007AEFA3 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007AF015 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007AF032 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007AF098 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007AF0B3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007AF11C | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007AF139 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007AF1B0 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007AF1D4 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007AF242 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007AF25D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007AF318 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AF334 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AF3A2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007AF3BF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007AF42A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007AF44A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007AF4C1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007AF4DD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007AF54D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007AF56C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007AF5D8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007AF5EC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007AF665 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007AF6D9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007AF749 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007AF7B1 | `NoContent_Screen` | Known | Screen layout |
| 0x007AF7C5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007AF829 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007AF890 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007AF8AA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007AF918 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007AF98A | `NoContent_Screen` | Known | Screen layout |
| 0x007AF99E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007AFA08 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007AFA71 | `No_Photos_Screen` | Known | Screen layout |
| 0x007AFA85 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007AFAEB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AFB59 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007AFBC6 | `NoContent_Screen` | Known | Screen layout |
| 0x007AFBDA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007AFC42 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007AFCAC | `NoContent_Screen` | Known | Screen layout |
| 0x007AFCC0 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007AFD27 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007AFD91 | `NoContent_Screen` | Known | Screen layout |
| 0x007AFDA5 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007AFE12 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007AFE84 | `NoContent_Screen` | Known | Screen layout |
| 0x007AFE98 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007AFF00 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007AFF69 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007AFF84 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007AFFEA | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B0006 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B00E5 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B00FE | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B015F | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B0173 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B02E1 | `Radio_Screen` | Known | Screen layout |
| 0x007B02F1 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B0352 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B03D5 | `LockediPod_Screen` | Known | Screen layout |
| 0x007B045D | `Lock_Screen` | Known | Screen layout |
| 0x007B046C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B04CF | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B0531 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B054D | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B05BF | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B05DE | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B0646 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B0660 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B06C8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B06E5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B0751 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B07BB | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B07D5 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B0845 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B08B8 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B0929 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B0998 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B0A04 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B0A1F | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B0A94 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B0AFB | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B0B5D | `Photos_Screen` | Known | Screen layout |
| 0x007B0BC1 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B0BDF | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B0C51 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B0C6E | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B0CD4 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B0CEF | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B0D58 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B0D75 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B0DEC | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B0E10 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B0E7E | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B0E99 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B0F54 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B0F70 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B0FDE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B0FFB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B1066 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B1086 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B10FD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B1119 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B1189 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B11A8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B1214 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B1228 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B12A1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B1315 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B1385 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B13ED | `NoContent_Screen` | Known | Screen layout |
| 0x007B1401 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B1465 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B14CC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B14E6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B1554 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B15C6 | `NoContent_Screen` | Known | Screen layout |
| 0x007B15DA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B1644 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B16AD | `No_Photos_Screen` | Known | Screen layout |
| 0x007B16C1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B1727 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B1795 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B1802 | `NoContent_Screen` | Known | Screen layout |
| 0x007B1816 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B187E | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B18E8 | `NoContent_Screen` | Known | Screen layout |
| 0x007B18FC | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B1963 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B19CD | `NoContent_Screen` | Known | Screen layout |
| 0x007B19E1 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B1A4E | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B1AC0 | `NoContent_Screen` | Known | Screen layout |
| 0x007B1AD4 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B1B3C | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B1BA5 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B1BC0 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B1C26 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B1C42 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B1D21 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B1D3A | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B1D9B | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B1DAF | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B1F1D | `Radio_Screen` | Known | Screen layout |
| 0x007B1F2D | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B1F8E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B2011 | `LockediPod_Screen` | Known | Screen layout |
| 0x007B2099 | `Lock_Screen` | Known | Screen layout |
| 0x007B20A8 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B210B | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B216D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B2189 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B21FB | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B221A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B2282 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B229C | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B2304 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B2321 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B238D | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B23F7 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B2411 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B2481 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B24F4 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B2565 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B25D4 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B2640 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B265B | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B26D0 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B2737 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B2799 | `Photos_Screen` | Known | Screen layout |
| 0x007B27FD | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B281B | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B288D | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B28AA | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B2910 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B292B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B2994 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B29B1 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B2A28 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B2A4C | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B2ABA | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B2AD5 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B2B90 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B2BAC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B2C1A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B2C37 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B2CA2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B2CC2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B2D39 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B2D55 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B2DC5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B2DE4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B2E50 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B2E64 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007B2EDD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B2F51 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007B2FC1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007B3029 | `NoContent_Screen` | Known | Screen layout |
| 0x007B303D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007B30A1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007B3108 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B3122 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007B3190 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007B3202 | `NoContent_Screen` | Known | Screen layout |
| 0x007B3216 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007B3280 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007B32E9 | `No_Photos_Screen` | Known | Screen layout |
| 0x007B32FD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007B3363 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B33D1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007B343E | `NoContent_Screen` | Known | Screen layout |
| 0x007B3452 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007B34BA | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007B3524 | `NoContent_Screen` | Known | Screen layout |
| 0x007B3538 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007B359F | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007B3609 | `NoContent_Screen` | Known | Screen layout |
| 0x007B361D | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007B368A | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007B36FC | `NoContent_Screen` | Known | Screen layout |
| 0x007B3710 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007B3778 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007B37E1 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007B37FC | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007B3862 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B387E | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B395D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007B3976 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007B39D7 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007B39EB | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007B3B59 | `Radio_Screen` | Known | Screen layout |
| 0x007B3B69 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007B3BCA | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007B3C4D | `LockediPod_Screen` | Known | Screen layout |
| 0x007B3CD5 | `Lock_Screen` | Known | Screen layout |
| 0x007B3CE4 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007B3D47 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007B3DA9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007B3DC5 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007B3E37 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B3E56 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B3EBE | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007B3ED8 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007B3F40 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B3F5D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B3FC9 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B4033 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007B404D | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007B40BD | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007B4130 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007B41A1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007B4210 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007B427C | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007B4297 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007B430C | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007B4373 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007B43D5 | `Photos_Screen` | Known | Screen layout |
| 0x007B4439 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007B4457 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007B44C9 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007B44E6 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007B454C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B4567 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B45D0 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007B45ED | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007B4664 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007B4688 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007B46F6 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007B4711 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007B47B1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B47CD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B483B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B4858 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B48C3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B48E3 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B495A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B4976 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B49E6 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B4A05 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B4A71 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B4A85 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007B4AFA | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007B4B65 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007B4BD4 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007B4C45 | `NoContent_Screen` | Known | Screen layout |
| 0x007B4C59 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B4CC8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007B4D3B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007B4DA8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007B4E11 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007B4E81 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007B4EF1 | `NoContent_Screen` | Known | Screen layout |
| 0x007B4F05 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007B4F68 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007B4FCB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B4FE7 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B50A7 | `Radio_Screen` | Known | Screen layout |
| 0x007B50B7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007B5118 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007B5186 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B51A5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B5213 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B5278 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B5293 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B5339 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B5355 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B53C3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B53E0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B544B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B546B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B54E2 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B54FE | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B556E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B558D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B55F9 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B560D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007B5682 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007B56ED | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007B575C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007B57CD | `NoContent_Screen` | Known | Screen layout |
| 0x007B57E1 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B5850 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007B58C3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007B5930 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007B5999 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007B5A09 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007B5A79 | `NoContent_Screen` | Known | Screen layout |
| 0x007B5A8D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007B5AF0 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007B5B53 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B5B6F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B5C2F | `Radio_Screen` | Known | Screen layout |
| 0x007B5C3F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007B5CA0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007B5D0E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B5D2D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B5D9B | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B5E00 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B5E1B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B5EC1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B5EDD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B5F4B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B5F68 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B5FD3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B5FF3 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B606A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B6086 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B60F6 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B6115 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B6181 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B6195 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007B620A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007B6275 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007B62E4 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007B6355 | `NoContent_Screen` | Known | Screen layout |
| 0x007B6369 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B63D8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007B644B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007B64B8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007B6521 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007B6591 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007B6601 | `NoContent_Screen` | Known | Screen layout |
| 0x007B6615 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007B6678 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007B66DB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B66F7 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B67B7 | `Radio_Screen` | Known | Screen layout |
| 0x007B67C7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007B6828 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007B6896 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B68B5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B6923 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B6988 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B69A3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B6A49 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B6A65 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B6AD3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B6AF0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B6B5B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B6B7B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B6BF2 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B6C0E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B6C7E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B6C9D | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B6D09 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B6D1D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007B6D92 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007B6DFD | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007B6E6C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007B6EDD | `NoContent_Screen` | Known | Screen layout |
| 0x007B6EF1 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B6F60 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007B6FD3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007B7040 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007B70A9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007B7119 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007B7189 | `NoContent_Screen` | Known | Screen layout |
| 0x007B719D | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007B7200 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007B7263 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B727F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B733F | `Radio_Screen` | Known | Screen layout |
| 0x007B734F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007B73B0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007B741E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B743D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B74AB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B7510 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B752B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B75D1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B75ED | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B765B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B7678 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B76E3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B7703 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B777A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B7796 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B7806 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B7825 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B7891 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B78A5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007B791A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007B7985 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007B79F4 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007B7A65 | `NoContent_Screen` | Known | Screen layout |
| 0x007B7A79 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B7AE8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007B7B5B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007B7BC8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007B7C31 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007B7CA1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007B7D11 | `NoContent_Screen` | Known | Screen layout |
| 0x007B7D25 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007B7D88 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007B7DEB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B7E07 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B7EC7 | `Radio_Screen` | Known | Screen layout |
| 0x007B7ED7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007B7F38 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007B7FA6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B7FC5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B8033 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B8098 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B80B3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B8159 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B8175 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B81E3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B8200 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B826B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B828B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B8302 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B831E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B838E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B83AD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B8419 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B842D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007B84A2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007B850D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007B857C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007B85ED | `NoContent_Screen` | Known | Screen layout |
| 0x007B8601 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B8670 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007B86E3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007B8750 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007B87B9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007B8829 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007B8899 | `NoContent_Screen` | Known | Screen layout |
| 0x007B88AD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007B8910 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007B8973 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B898F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B8A4F | `Radio_Screen` | Known | Screen layout |
| 0x007B8A5F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007B8AC0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007B8B2E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B8B4D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B8BBB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B8C20 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B8C3B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B8CE1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B8CFD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B8D6B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B8D88 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B8DF3 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B8E13 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B8E8A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B8EA6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B8F16 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B8F35 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B8FA1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B8FB5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007B902A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007B9095 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007B9104 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007B9175 | `NoContent_Screen` | Known | Screen layout |
| 0x007B9189 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B91F8 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007B926B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007B92D8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007B9341 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007B93B1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007B9421 | `NoContent_Screen` | Known | Screen layout |
| 0x007B9435 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007B9498 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007B94FB | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007B9517 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007B95D7 | `Radio_Screen` | Known | Screen layout |
| 0x007B95E7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007B9648 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007B96B6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007B96D5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B9743 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007B97A8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007B97C3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007B9869 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B9885 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B98F3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007B9910 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007B997B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007B999B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007B9A12 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007B9A2E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007B9A9E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007B9ABD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007B9B29 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007B9B3D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007B9BB2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007B9C1D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007B9C8C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007B9CFD | `NoContent_Screen` | Known | Screen layout |
| 0x007B9D11 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007B9D80 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007B9DF3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007B9E60 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007B9EC9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007B9F39 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007B9FA9 | `NoContent_Screen` | Known | Screen layout |
| 0x007B9FBD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007BA020 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007BA083 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BA09F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BA15F | `Radio_Screen` | Known | Screen layout |
| 0x007BA16F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007BA1D0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007BA23E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BA25D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BA2CB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BA330 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BA34B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BA3F1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BA40D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BA47B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BA498 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BA503 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BA523 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BA59A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BA5B6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BA626 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BA645 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BA6B1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BA6C5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007BA73A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007BA7A5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007BA814 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007BA885 | `NoContent_Screen` | Known | Screen layout |
| 0x007BA899 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BA908 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007BA97B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007BA9E8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007BAA51 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007BAAC1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007BAB31 | `NoContent_Screen` | Known | Screen layout |
| 0x007BAB45 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007BABA8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007BAC0B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BAC27 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BACE7 | `Radio_Screen` | Known | Screen layout |
| 0x007BACF7 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007BAD58 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007BADC6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BADE5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BAE53 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BAEB8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BAED3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BAF79 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BAF95 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BB003 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BB020 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BB08B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BB0AB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BB122 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BB13E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BB1AE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BB1CD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BB239 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BB24D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007BB2C2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007BB32D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007BB39C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007BB40D | `NoContent_Screen` | Known | Screen layout |
| 0x007BB421 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BB490 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007BB503 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007BB570 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007BB5D9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007BB649 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007BB6B9 | `NoContent_Screen` | Known | Screen layout |
| 0x007BB6CD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007BB730 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007BB793 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BB7AF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BB86F | `Radio_Screen` | Known | Screen layout |
| 0x007BB87F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007BB8E0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007BB94E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BB96D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BB9DB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BBA40 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BBA5B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BBB01 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BBB1D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BBB8B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BBBA8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BBC13 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BBC33 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BBCAA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BBCC6 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BBD36 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BBD55 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BBDC1 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BBDD5 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007BBE4A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007BBEB5 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007BBF24 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007BBF95 | `NoContent_Screen` | Known | Screen layout |
| 0x007BBFA9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BC018 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007BC08B | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007BC0F8 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007BC161 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007BC1D1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007BC241 | `NoContent_Screen` | Known | Screen layout |
| 0x007BC255 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007BC2B8 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007BC31B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BC337 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BC3F7 | `Radio_Screen` | Known | Screen layout |
| 0x007BC407 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007BC468 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007BC4D6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BC4F5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BC563 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BC5C8 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BC5E3 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BC689 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BC6A5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BC713 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BC730 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BC79B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007BC7BB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007BC832 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BC84E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007BC8BE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007BC8DD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007BC949 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007BC95D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007BC9D2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007BCA3D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007BCAAC | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007BCB1D | `NoContent_Screen` | Known | Screen layout |
| 0x007BCB31 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007BCBA0 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007BCC13 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007BCC80 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007BCCE9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007BCD59 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007BCDC9 | `NoContent_Screen` | Known | Screen layout |
| 0x007BCDDD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007BCE40 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007BCEA3 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007BCEBF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007BCF7F | `Radio_Screen` | Known | Screen layout |
| 0x007BCF8F | `Radio_Screen_Default` | Known | Screen layout |
| 0x007BCFF0 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007BD05E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007BD07D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007BD0EB | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007BD150 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BD16B | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BD24C | `MediaLists_AudiobookChapters_Screen)` | Known | Screen layout |
| 0x007BD273 | `MediaLists_AudiobookChapters_Screen_Plane!` | Known | Screen layout |
| 0x007BDA0D | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007BDA28 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x007BDA93 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BDAAE | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x007BDB21 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x007BDB3C | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x007BDCF9 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007BDD14 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x007BDD7F | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BDD9A | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x007BDE0D | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x007BDE28 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x007BDFF0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BE00C | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x007BE087 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BE0A3 | `MediaLists_Albums_Screen_WithTrackCount%` | Known | Screen layout |
| 0x007BE11C | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007BE137 | `MediaLists_Songs_Screen_Plain1` | Known | Screen layout |
| 0x007BE1B2 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x007BE1CD | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x007BE3EF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007BE40C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007BE4EB | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BE507 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x007BE582 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007BE59D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007BE783 | `MediaLists_PodcastEpisodes_Screen'` | Known | Screen layout |
| 0x007BE7A8 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007BEA7A | `MediaLists_TVSeasons_Screen#` | Known | Screen layout |
| 0x007BEA99 | `MediaLists_TVSeasons_Screen_Default%` | Known | Screen layout |
| 0x007BEB0E | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x007BEB2E | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007BECB6 | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x007BECD6 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007BF0CF | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x007BF0F4 | `MediaLists_NestedPlaylists_Screen_Default,` | Known | Screen layout |
| 0x007BF176 | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x007BF195 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x007BF325 | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x007BF34A | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x007BF3C2 | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x007BF3E1 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x007BF445 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x007BF4F2 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x007BF564 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x007BF65A | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x007BF91C | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x007BFA1C | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x007BFA88 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007BFAF2 | `NoContent_Screen` | Known | Screen layout |
| 0x007BFB06 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007BFB70 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007BFBE4 | `NoContent_Screen` | Known | Screen layout |
| 0x007BFBF8 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007BFC63 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007BFCCF | `NoContent_Screen` | Known | Screen layout |
| 0x007BFCE3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007BFD4A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007BFDB6 | `NoContent_Screen` | Known | Screen layout |
| 0x007BFDCA | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007BFE37 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007BFEAB | `NoContent_Screen` | Known | Screen layout |
| 0x007BFEBF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007BFF27 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007BFF94 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007BFFF8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007C0014 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007C0080 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C009D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C010A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007C01D1 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007C01EE | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007C0265 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007C0289 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C0340 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007C03AA | `NoContent_Screen` | Known | Screen layout |
| 0x007C03BE | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007C0428 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007C049C | `NoContent_Screen` | Known | Screen layout |
| 0x007C04B0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007C051B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007C0587 | `NoContent_Screen` | Known | Screen layout |
| 0x007C059B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007C0602 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007C066E | `NoContent_Screen` | Known | Screen layout |
| 0x007C0682 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007C06EF | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007C0763 | `NoContent_Screen` | Known | Screen layout |
| 0x007C0777 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007C07DF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007C084C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007C08B0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007C08CC | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007C0938 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C0955 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C09C2 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007C0A89 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007C0AA6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007C0B1D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007C0B41 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C0BF8 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007C0C62 | `NoContent_Screen` | Known | Screen layout |
| 0x007C0C76 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007C0CE0 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007C0D54 | `NoContent_Screen` | Known | Screen layout |
| 0x007C0D68 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007C0DD3 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007C0E3F | `NoContent_Screen` | Known | Screen layout |
| 0x007C0E53 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007C0EBA | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007C0F26 | `NoContent_Screen` | Known | Screen layout |
| 0x007C0F3A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007C0FA7 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007C101B | `NoContent_Screen` | Known | Screen layout |
| 0x007C102F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007C1097 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007C1104 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007C1168 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007C1184 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007C11F0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C120D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C127A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007C1341 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007C135E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007C13D5 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007C13F9 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C14B0 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007C151A | `NoContent_Screen` | Known | Screen layout |
| 0x007C152E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007C1598 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007C160C | `NoContent_Screen` | Known | Screen layout |
| 0x007C1620 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007C168B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007C16F7 | `NoContent_Screen` | Known | Screen layout |
| 0x007C170B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007C1772 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007C17DE | `NoContent_Screen` | Known | Screen layout |
| 0x007C17F2 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007C185F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007C18D3 | `NoContent_Screen` | Known | Screen layout |
| 0x007C18E7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007C194F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007C19BC | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007C1A20 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007C1A3C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007C1AA8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C1AC5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C1B32 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007C1BF9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007C1C16 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007C1C8D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007C1CB1 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C1D68 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007C1DD2 | `NoContent_Screen` | Known | Screen layout |
| 0x007C1DE6 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007C1E50 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007C1EC4 | `NoContent_Screen` | Known | Screen layout |
| 0x007C1ED8 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007C1F43 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007C1FAF | `NoContent_Screen` | Known | Screen layout |
| 0x007C1FC3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007C202A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007C2096 | `NoContent_Screen` | Known | Screen layout |
| 0x007C20AA | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007C2117 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007C218B | `NoContent_Screen` | Known | Screen layout |
| 0x007C219F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007C2207 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007C2274 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007C22D8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007C22F4 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007C2360 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C237D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C23EA | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007C24B1 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007C24CE | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007C2545 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007C2569 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C2620 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007C268A | `NoContent_Screen` | Known | Screen layout |
| 0x007C269E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007C2708 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007C277C | `NoContent_Screen` | Known | Screen layout |
| 0x007C2790 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007C27FB | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007C2867 | `NoContent_Screen` | Known | Screen layout |
| 0x007C287B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007C28E2 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007C294E | `NoContent_Screen` | Known | Screen layout |
| 0x007C2962 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007C29CF | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007C2A43 | `NoContent_Screen` | Known | Screen layout |
| 0x007C2A57 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007C2ABF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007C2B2C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007C2B90 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007C2BAC | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007C2C18 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C2C35 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C2CA2 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007C2D69 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007C2D86 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007C2DFD | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007C2E21 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C2ED8 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007C2F42 | `NoContent_Screen` | Known | Screen layout |
| 0x007C2F56 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007C2FC0 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007C3034 | `NoContent_Screen` | Known | Screen layout |
| 0x007C3048 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007C30B3 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007C311F | `NoContent_Screen` | Known | Screen layout |
| 0x007C3133 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007C319A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007C3206 | `NoContent_Screen` | Known | Screen layout |
| 0x007C321A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007C3287 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007C32FB | `NoContent_Screen` | Known | Screen layout |
| 0x007C330F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007C3377 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007C33E4 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007C3448 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007C3464 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007C34D0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C34ED | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C355A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007C3621 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007C363E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007C36B5 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007C36D9 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C3790 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007C37FA | `NoContent_Screen` | Known | Screen layout |
| 0x007C380E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007C3878 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007C38EC | `NoContent_Screen` | Known | Screen layout |
| 0x007C3900 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007C396B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007C39D7 | `NoContent_Screen` | Known | Screen layout |
| 0x007C39EB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007C3A52 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007C3ABE | `NoContent_Screen` | Known | Screen layout |
| 0x007C3AD2 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007C3B3F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007C3BB3 | `NoContent_Screen` | Known | Screen layout |
| 0x007C3BC7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007C3C2F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007C3C9C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007C3D00 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007C3D1C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007C3D88 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007C3DA5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007C3E12 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007C3ED9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007C3EF6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007C3F6D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007C3F91 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007C43F4 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007C4466 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007C44D1 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007C4536 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007C45A0 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007C460A | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007C467A | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007C46F1 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007C475F | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007C47CA | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007C4834 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007C489B | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007C490A | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007C4978 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007C49DD | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007C4A45 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007C4AB0 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007C4B1B | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007C4B82 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007C4EF0 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007C4F62 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007C4FCD | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007C5032 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007C509C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007C5106 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007C5176 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007C51ED | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007C525B | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007C52C6 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007C5330 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007C5397 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007C5406 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007C5474 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007C54D9 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007C5541 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007C55AC | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007C5617 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007C567E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007C59EA | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007C5A5C | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007C5AC7 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007C5B2C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007C5B96 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007C5C00 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007C5C70 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007C5CE7 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007C5D55 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007C5DC0 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007C5E2A | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007C5E91 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007C5F00 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007C5F6E | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007C5FD3 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007C603B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007C60A6 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007C6111 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007C6178 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007C64E2 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007C6554 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007C65BF | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007C6624 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007C668E | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007C66F8 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007C6768 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007C67DF | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007C684D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007C68B8 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007C6922 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007C6989 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007C69F8 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007C6A66 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007C6ACB | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007C6B33 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007C6B9E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007C6C09 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007C6C70 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007C6FC2 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007C7034 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007C709F | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007C7104 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007C716E | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007C71D8 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007C7248 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007C72BF | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007C732D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007C7398 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007C7402 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007C7469 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007C74D8 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007C7546 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007C75AB | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007C7613 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007C767E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007C76E9 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007C7750 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007C7AC7 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007C7B39 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007C7BA4 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007C7C09 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007C7C73 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007C7CDD | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007C7D4D | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007C7DC4 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007C7E32 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007C7E9D | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007C7F07 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007C7F6E | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007C7FDD | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007C804B | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007C80B0 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007C8118 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007C8183 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007C81EE | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007C8255 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007C85C9 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007C863B | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007C86A6 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007C870B | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007C8775 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007C87DF | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007C884F | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007C88C6 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007C8934 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007C899F | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007C8A09 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007C8A70 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007C8ADF | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007C8B4D | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007C8BB2 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007C8C1A | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007C8C85 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007C8CF0 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007C8D57 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007C90B1 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007C9123 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007C918E | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007C91F3 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007C925D | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007C92C7 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007C9337 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007C93AE | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007C941C | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007C9487 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007C94F1 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007C9558 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007C95C7 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007C9635 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007C969A | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007C9702 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007C976D | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007C97D8 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007C983F | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007C9B99 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007C9C0B | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007C9C76 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007C9CDB | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007C9D45 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007C9DAF | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007C9E1F | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007C9E96 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007C9F04 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007C9F6F | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007C9FD9 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CA040 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CA0AF | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CA11D | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CA182 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CA1EA | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CA255 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CA2C0 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CA327 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CA682 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CA6F4 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CA75F | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CA7C4 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CA82E | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CA898 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CA908 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CA97F | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CA9ED | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CAA58 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CAAC2 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CAB29 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CAB98 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CAC06 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CAC6B | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CACD3 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CAD3E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CADA9 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CAE10 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CB194 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CB206 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CB271 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CB2D6 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CB340 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CB3AA | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CB41A | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CB491 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CB4FF | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CB56A | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CB5D4 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CB63B | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CB6AA | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CB718 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CB77D | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CB7E5 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CB850 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CB8BB | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CB922 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CBCB0 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CBD22 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CBD8D | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CBDF2 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CBE5C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CBEC6 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CBF36 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CBFAD | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CC01B | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CC086 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CC0F0 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CC157 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CC1C6 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CC234 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CC299 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CC301 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CC36C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CC3D7 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CC43E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CC7AC | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CC81E | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CC889 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CC8EE | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CC958 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CC9C2 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CCA32 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CCAA9 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CCB17 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CCB82 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CCBEC | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CCC53 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CCCC2 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CCD30 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CCD95 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CCDFD | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CCE68 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CCED3 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CCF3A | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CD2A0 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CD312 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CD37D | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CD3E2 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CD44C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CD4B6 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CD526 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CD59D | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CD60B | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CD676 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CD6E0 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CD747 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CD7B6 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CD824 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CD889 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CD8F1 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CD95C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CD9C7 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CDA2E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CDD82 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CDDF4 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CDE5F | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CDEC4 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CDF2E | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CDF98 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CE008 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CE07F | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CE0ED | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CE158 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CE1C2 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CE229 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CE298 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CE306 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CE36B | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CE3D3 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CE43E | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CE4A9 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CE510 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CE85B | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CE8CD | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CE938 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CE99D | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CEA07 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CEA71 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CEAE1 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CEB58 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CEBC6 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CEC31 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CEC9B | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CED02 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CED71 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CEDDF | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CEE44 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CEEAC | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CEF17 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CEF82 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CEFE9 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CF34B | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CF3BD | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CF428 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CF48D | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CF4F7 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007CF561 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007CF5D1 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007CF648 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007CF6B6 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007CF721 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007CF78B | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007CF7F2 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007CF861 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007CF8CF | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007CF934 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007CF99C | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007CFA07 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007CFA72 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007CFAD9 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007CFDF1 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007CFE63 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007CFECE | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007CFF33 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007CFF9D | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007D0007 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007D0077 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007D00EE | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x007D015C | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007D01C7 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007D0231 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007D0298 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007D0307 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007D0375 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007D03DA | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007D0442 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007D04AD | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007D0518 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007D057F | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007D0896 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007D090D | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007D098A | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007D09FC | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007D0A6C | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007D0AE2 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007D0B50 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007D0BBD | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007D0F02 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007D0F79 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007D0FF6 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007D1068 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007D10D8 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007D114E | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007D11BC | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007D1229 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007D1592 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007D1609 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007D1686 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007D16F8 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007D1768 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007D17DE | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007D184C | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007D18B9 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007D1C22 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007D1C99 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007D1D14 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007D1D84 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007D1DFA | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007D1E68 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007D1ED5 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007D220E | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007D2285 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007D2300 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007D2370 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007D23E6 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007D2454 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007D24C1 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007D27F8 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007D286F | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007D28EA | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007D295A | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007D29D0 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007D2A3E | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007D2AAB | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007D2DBB | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007D2E32 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007D2EAD | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007D2F1D | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007D2F93 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007D3001 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007D306E | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007D3672 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007D368F | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007D370A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007D3723 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007D379B | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007D37B4 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007D3829 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007D383F | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007D38B6 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007D38CC | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007D3943 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007D3960 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007D39D8 | `Notes_List_Screen` | Known | Screen layout |
| 0x007D39ED | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007D3B9E | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007D3BBB | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007D3C36 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007D3C4F | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007D3CC7 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007D3CE0 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007D3D55 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007D3D6B | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007D3DE2 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007D3DF8 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007D3E6F | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007D3E8C | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007D3F04 | `Notes_List_Screen` | Known | Screen layout |
| 0x007D3F19 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007D40FA | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007D4117 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007D4192 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007D41AB | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007D4223 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007D423C | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007D42B1 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007D42C7 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007D433E | `Notes_Image_Screen` | Known | Screen layout |
| 0x007D4354 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007D43CB | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007D43E8 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007D4460 | `Notes_List_Screen` | Known | Screen layout |
| 0x007D4475 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007D462A | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007D4647 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007D46C2 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007D46DB | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007D4753 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007D476C | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007D47E1 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007D47F7 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007D486E | `Notes_Image_Screen` | Known | Screen layout |
| 0x007D4884 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007D48FB | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007D4918 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007D4990 | `Notes_List_Screen` | Known | Screen layout |
| 0x007D49A5 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007D4CBD | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007D4D63 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007D4DE6 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007D4E9E | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x007D4F20 | `PhotosSettingsSlideshowMusic_Screen+` | Known | Screen layout |
| 0x007D4F47 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x007D502D | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x007D51E5 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007D5245 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007D52A2 | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007D52C9 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007D5369 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007D53C9 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007D5426 | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007D544D | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007D56E8 | `Photos_Screen` | Known | Screen layout |
| 0x007D5834 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007D5898 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007D58F9 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007D5956 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007D59B3 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007D5A21 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007D5A7E | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007D5BF4 | `Photos_Screen` | Known | Screen layout |
| 0x007D5D40 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007D5DA4 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007D5E05 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007D5E62 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007D5EBF | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007D5F2D | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007D5F8A | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007D6100 | `Photos_Screen` | Known | Screen layout |
| 0x007D624C | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007D62B0 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007D6311 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007D636E | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007D63CB | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007D6439 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007D6496 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007D660C | `Photos_Screen` | Known | Screen layout |
| 0x007D6758 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007D67BC | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007D681D | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007D687A | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007D68D7 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007D6945 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007D69A2 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007D6B18 | `Photos_Screen` | Known | Screen layout |
| 0x007D6C64 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007D6CC8 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007D6D29 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007D6D86 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007D6DE3 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007D6E51 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007D6EAE | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007D7024 | `Photos_Screen` | Known | Screen layout |
| 0x007D7170 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007D71D4 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007D7235 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007D7292 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007D72EF | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007D735D | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007D73BA | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007D7530 | `Photos_Screen` | Known | Screen layout |
| 0x007D767C | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007D76E2 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007D7744 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007D77A6 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007D783C | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007D795D | `Photos_Screen` | Known | Screen layout |
| 0x007D79C8 | `Photos_Screen` | Known | Screen layout |
| 0x007D7B14 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007D7B7A | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007D7BDC | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007D7C3E | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007D7CD4 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007D7DF5 | `Photos_Screen` | Known | Screen layout |
| 0x007D7E60 | `Photos_Screen` | Known | Screen layout |
| 0x007D7FAC | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007D8012 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007D8074 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007D80D6 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007D816C | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007D828D | `Photos_Screen` | Known | Screen layout |
| 0x007D82F8 | `Photos_Screen` | Known | Screen layout |
| 0x007D8444 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007D84AA | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007D850C | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007D856E | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007D8604 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007D8725 | `Photos_Screen` | Known | Screen layout |
| 0x007D8790 | `Photos_Screen` | Known | Screen layout |
| 0x007D88DC | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007D8942 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007D89A4 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007D8A06 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007D8A9C | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007D8BBD | `Photos_Screen` | Known | Screen layout |
| 0x007D8DB1 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007D8E13 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007D8E81 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007D8EE7 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007D8F4C | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007D921A | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007D927C | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007D92EA | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007D9350 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007D9656 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007D96B8 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007D9726 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007D978C | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007D9A35 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007D9A92 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007D9AF4 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007D9B62 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007D9BC8 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007D9EC2 | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007D9F2C | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007DA19A | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007DA204 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007DA3C1 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007DA424 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007DA489 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007DA4F1 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007DA554 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DA5BC | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007DA625 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DA68B | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007DA6F0 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DA75D | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007DA7CD | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007DA843 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007DA8B9 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007DA929 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DA99E | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007DAA15 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007DAA89 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007DAAFB | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007DAB75 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DABE8 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007DAC5A | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DACDE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DAD08 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DAD8F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DAE1C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DAEBB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DAED5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DAF4D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DAF67 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DAFD1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DAFEE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DB066 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DB090 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DB117 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DB1A4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DB243 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DB25D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DB2D5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DB2EF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DB359 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DB376 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DB3EE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DB418 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DB49F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DB52C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DB5CB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DB5E5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DB65D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DB677 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DB6E1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DB6FE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DB776 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DB7A0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DB827 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DB8B4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DB953 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DB96D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DB9E5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DB9FF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DBA69 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DBA86 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DBAFE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DBB28 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DBBAF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DBC3C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DBCDB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DBCF5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DBD6D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DBD87 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DBDF1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DBE0E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DBE86 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DBEB0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DBF37 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DBFC4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DC063 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DC07D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DC0F5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DC10F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DC179 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DC196 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DC20E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DC238 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DC2BF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DC34C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DC3EB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DC405 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DC47D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DC497 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DC501 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DC51E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DC596 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DC5C0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DC647 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DC6D4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DC773 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DC78D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DC805 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DC81F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DC889 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DC8A6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DC91E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DC948 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DC9CF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DCA5C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DCAFB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DCB15 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DCB8D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DCBA7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DCC11 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DCC2E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DCCA6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DCCD0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DCD57 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DCDE4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DCE83 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DCE9D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DCF15 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DCF2F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DCF99 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DCFB6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DD02E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DD058 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DD0DF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DD16C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DD20B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DD225 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DD29D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DD2B7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DD321 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DD33E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DD3B6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DD3E0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DD467 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DD4F4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DD593 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DD5AD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DD625 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DD63F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DD6A9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DD6C6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DD73E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DD768 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DD7EF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DD87C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DD91B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DD935 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DD9AD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DD9C7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DDA31 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DDA4E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DDAC6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DDAF0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DDB77 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DDC04 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DDCA3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DDCBD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DDD35 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DDD4F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DDDB9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DDDD6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DDE4E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DDE78 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DDEFF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DDF8C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DE02B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DE045 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DE0BD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DE0D7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DE141 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DE15E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DE1D6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DE200 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DE287 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DE314 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DE3B3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DE3CD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DE445 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DE45F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DE4C9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DE4E6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DE55E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DE588 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DE60F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DE69C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DE73B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DE755 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DE7CD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DE7E7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DE851 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DE86E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DE8E6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DE910 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DE997 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DEA24 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DEAC3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DEADD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DEB55 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DEB6F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DEBD9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DEBF6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DEC6E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007DEC98 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007DED1F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007DEDAC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007DEE4B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DEE65 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DEEDD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DEEF7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DEF61 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007DEF7E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007DF005 | `SettingsMenus_ResetMainMenu_Screen9` | Known | Screen layout |
| 0x007DF0D5 | `SettingsMenus_ResetMusicMenu_Screen:` | Known | Screen layout |
| 0x007DF189 | `SettingsMenus_Main_Screen,` | Known | Screen layout |
| 0x007DF1FB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DF215 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007DF28D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007DF2A7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007DF5E2 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007DF648 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007DF6A5 | `Extras_Screen` | Known | Screen layout |
| 0x007DF6F9 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x007DF7D7 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x007DF845 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007DF8E3 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x007DF8FC | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x007DF964 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007DF9D7 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007DFA59 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007DFABA | `NikePlus_Alert_Screen$` | Known | Screen layout |
| 0x007DFB3A | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007DFBB3 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007DFC2D | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007DFCB2 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007DFCD3 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007DFD42 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x007DFDCA | `NikePlus_Remote_Unlinking_Screen(` | Known | Screen layout |
| 0x007DFDEE | `NikePlus_Remote_Unlinking_Screen_Default` | Known | Screen layout |
| 0x007DFE62 | `NikePlus_Alert_Screen(` | Known | Screen layout |
| 0x007DFEEE | `NikePlus_Dynamic_Workout_Screen,` | Known | Screen layout |
| 0x007DFF11 | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x007DFF8A | `NikePlus_Dynamic_Workout_Screen/` | Known | Screen layout |
| 0x007DFFAD | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x007E0026 | `NikePlus_Dynamic_Workout_Screen.` | Known | Screen layout |
| 0x007E0049 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x007E00C2 | `NikePlus_Alert_Screen(` | Known | Screen layout |
| 0x007E0145 | `NikePlus_Custom_Screen,` | Known | Screen layout |
| 0x007E015F | `NikePlus_Custom_Screen_Weight_ToWorkoutMusic` | Known | Screen layout |
| 0x007E01D9 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E025B | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007E02CF | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007E02ED | `NikePlus_EndWorkout_Screen_Default` | Known | Screen layout |
| 0x007E0385 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007E0401 | `NikePlus_EquipmentAlert_Screen<` | Known | Screen layout |
| 0x007E04CE | `NikePlus_EquipmentAlert_Screen<` | Known | Screen layout |
| 0x007E0598 | `NikePlus_EquipmentAlert_Screen8` | Known | Screen layout |
| 0x007E0665 | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007E0726 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E0747 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E07DE | `NikePlus_Dynamic_Workout_Screen.` | Known | Screen layout |
| 0x007E0801 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x007E08A1 | `NikePlus_Dynamic_Workout_Screen/` | Known | Screen layout |
| 0x007E08C4 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x007E0962 | `NikePlus_Dynamic_Workout_Screen,` | Known | Screen layout |
| 0x007E0985 | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x007E0A1B | `NikePlus_EndPausedWorkout_Screen1` | Known | Screen layout |
| 0x007E0A3F | `NikePlus_EndPausedWorkout_Screen_BasicSave_Layout` | Known | Screen layout |
| 0x007E0ADD | `NikePlus_EndPausedWorkout_Screen4` | Known | Screen layout |
| 0x007E0B01 | `NikePlus_EndPausedWorkout_Screen_CaloriesSave_Layout` | Known | Screen layout |
| 0x007E0BA2 | `NikePlus_EndPausedWorkout_Screen4` | Known | Screen layout |
| 0x007E0BC6 | `NikePlus_EndPausedWorkout_Screen_DistanceSave_Layout` | Known | Screen layout |
| 0x007E0C64 | `NikePlus_EndPausedWorkout_Screen0` | Known | Screen layout |
| 0x007E0C88 | `NikePlus_EndPausedWorkout_Screen_TimeSave_Layout` | Known | Screen layout |
| 0x007E0D1F | `NikePlus_Songs_Screen` | Known | Screen layout |
| 0x007E0D38 | `NikePlus_Songs_Screen_Default` | Known | Screen layout |
| 0x007E0E4A | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007E0E64 | `NikePlus_Custom_Screen_Weight` | Known | Screen layout |
| 0x007E0EC7 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007E0F3B | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007E0FB9 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x007E1023 | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x007E1083 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E1100 | `NikePlus_SensorSearching_Screen'` | Known | Screen layout |
| 0x007E1123 | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x007E1192 | `NikePlus_Playlists_Screen ` | Known | Screen layout |
| 0x007E11AF | `NikePlus_Playlists_Screen_Nested` | Known | Screen layout |
| 0x007E1243 | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x007E12A3 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E1320 | `NikePlus_SensorSearching_Screen'` | Known | Screen layout |
| 0x007E1343 | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x007E13B2 | `NikePlus_Playlists_Screen ` | Known | Screen layout |
| 0x007E13CF | `NikePlus_Playlists_Screen_Nested` | Known | Screen layout |
| 0x007E1498 | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x007E14F8 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E1575 | `NikePlus_Playlists_Screen!` | Known | Screen layout |
| 0x007E1592 | `NikePlus_Playlists_Screen_Default` | Known | Screen layout |
| 0x007E15FE | `NikePlus_SensorSearching_Screen'` | Known | Screen layout |
| 0x007E1621 | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x007E17BD | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007E17DB | `NikePlus_NowRunning_Screen_Basic'` | Known | Screen layout |
| 0x007E184F | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007E186D | `NikePlus_NowRunning_Screen_Calories'` | Known | Screen layout |
| 0x007E18E4 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007E1902 | `NikePlus_NowRunning_Screen_Distance#` | Known | Screen layout |
| 0x007E1975 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007E1993 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007E19FB | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E1AD7 | `NikePlus_NowRunning_Screen$` | Known | Screen layout |
| 0x007E1AF5 | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x007E1BBF | `NikePlus_NowRunning_Screen$` | Known | Screen layout |
| 0x007E1BDD | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x007E1CA7 | `NikePlus_NowRunning_Screen$` | Known | Screen layout |
| 0x007E1CC5 | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x007E1F8A | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007E1FB6 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007E2038 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007E2066 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007E20E8 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007E210A | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007E217C | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007E219F | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007E220F | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007E222D | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007E229D | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007E22C3 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007E2337 | `NikePlus_NowRunning_Screen_Basic_PowerSong` | Known | Screen layout |
| 0x007E26C4 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007E26F0 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007E2772 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007E27A0 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007E2822 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007E2844 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007E28B6 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007E28D9 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007E2949 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007E2967 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007E29D7 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007E29FD | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007E2A6E | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007E2DEE | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007E2E1A | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007E2E9C | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007E2ECA | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007E2F4C | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007E2F6E | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007E2FE0 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007E3003 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007E3073 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007E3091 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007E3101 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007E3127 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007E319B | `NikePlus_NowRunning_Screen_Time_PowerSong` | Known | Screen layout |
| 0x007E3528 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007E3554 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007E35D6 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007E3604 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007E3686 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007E36A8 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007E371A | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007E373D | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007E37AD | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007E37CB | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007E383B | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007E3861 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007E38D2 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007E3C52 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007E3C7E | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007E3D00 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007E3D2E | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007E3DB0 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007E3DD2 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007E3E44 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007E3E67 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007E3ED7 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007E3EF5 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007E3F65 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007E3F8B | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007E3FFF | `NikePlus_NowRunning_Screen_Distance_PowerSong` | Known | Screen layout |
| 0x007E4390 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007E43BC | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007E443E | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007E446C | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007E44EE | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007E4510 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007E4582 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007E45A5 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007E4615 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007E4633 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007E46A3 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007E46C9 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007E473A | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007E4ABE | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007E4AEA | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007E4B6C | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007E4B9A | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007E4C1C | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007E4C3E | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007E4CB0 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007E4CD3 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007E4D43 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007E4D61 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007E4DD1 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007E4DF7 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007E4E6B | `NikePlus_NowRunning_Screen_Calories_PowerSong` | Known | Screen layout |
| 0x007E51FC | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007E5228 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007E52AA | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007E52D8 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007E535A | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007E537C | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007E53EE | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007E5411 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007E5481 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007E549F | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007E550F | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007E5535 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007E55A6 | `NikePlus_NowRunning_Screen_Calories` | Known | Screen layout |
| 0x007E58F4 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007E5920 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007E59A2 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007E59D0 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007E5A52 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007E5A74 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007E5AE6 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007E5B09 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007E5B79 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007E5B97 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007E5C07 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007E5C2D | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007E5DC8 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007E5E2F | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007E5EA3 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007E5F16 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007E5F82 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E5FA3 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E601E | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007E6044 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007E6101 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007E612D | `NikePlus_CalibrationCompleteError_Screen_Default'` | Known | Screen layout |
| 0x007E61B1 | `NikePlus_CalibrationCompleteError_Screen*` | Known | Screen layout |
| 0x007E61DD | `NikePlus_CalibrationComplete_Screen_Pacing%` | Known | Screen layout |
| 0x007E6259 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007E6287 | `NikePlus_CalibrationCompleteSuccess_Screen_Default` | Known | Screen layout |
| 0x007E6300 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007E6390 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007E63E3 | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x007E6450 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007E64A3 | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x007E6510 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007E6564 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007E65CA | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007E65E8 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007E6654 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007E6672 | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007E66E2 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007E6700 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007E676C | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007E678A | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007E6837 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007E685D | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007E68F0 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007E690A | `NikePlus_Custom_Screen_Time` | Known | Screen layout |
| 0x007E698B | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E69AC | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E6A3F | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007E6A59 | `NikePlus_Custom_Screen_Distance` | Known | Screen layout |
| 0x007E6AE1 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E6B02 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E6B7F | `NikePlus_EquipmentAlert_Screen)` | Known | Screen layout |
| 0x007E6C18 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E6C39 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E6CC4 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007E6CDE | `NikePlus_Custom_Screen_Calories` | Known | Screen layout |
| 0x007E6D91 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E6E18 | `NikePlus_Alert_Screen(` | Known | Screen layout |
| 0x007E6EC1 | `NikePlus_EquipmentAlert_ScreenK` | Known | Screen layout |
| 0x007E6F84 | `NikePlus_EquipmentAlert_Screen8` | Known | Screen layout |
| 0x007E7038 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007E7128 | `NikePlus_EquipmentAlert_Screen>` | Known | Screen layout |
| 0x007E71E6 | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007E72A5 | `NikePlus_Alert_Screen$` | Known | Screen layout |
| 0x007E7326 | `NikePlus_Remote_Unlinking_Screen(` | Known | Screen layout |
| 0x007E734A | `NikePlus_Remote_Unlinking_Screen_Default!` | Known | Screen layout |
| 0x007E73C0 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x007E74FE | `NikePlus_Calibration_CalibrateWalk_Screen1` | Known | Screen layout |
| 0x007E75A2 | `NikePlus_Calibration_CalibrateRun_Screen0` | Known | Screen layout |
| 0x007E7665 | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007E7726 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E7747 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E77CE | `NikePlus_Custom_Screen*` | Known | Screen layout |
| 0x007E77E8 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x007E78D9 | `NikePlus_Calibrate_ResetToDefault_Screen0` | Known | Screen layout |
| 0x007E7905 | `NikePlus_Calibrate_ResetToDefault_Screen_Walking` | Known | Screen layout |
| 0x007E79BB | `NikePlus_Calibration_CalibrateRun_Screen"` | Known | Screen layout |
| 0x007E7A33 | `NikePlus_Calibration_CalibrateWalk_Screen` | Known | Screen layout |
| 0x007E7ADF | `NikePlus_Calibration_CalibrateRun_Screen"` | Known | Screen layout |
| 0x007E7B57 | `NikePlus_Calibration_CalibrateWalk_Screen` | Known | Screen layout |
| 0x007E7BE5 | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007E7CA6 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E7CC7 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E7D4E | `NikePlus_Custom_Screen*` | Known | Screen layout |
| 0x007E7D68 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x007E7E57 | `NikePlus_Calibrate_ResetToDefault_Screen0` | Known | Screen layout |
| 0x007E7E83 | `NikePlus_Calibrate_ResetToDefault_Screen_Running` | Known | Screen layout |
| 0x007E7EF5 | `NikePlus_StartWorkout_Screen$` | Known | Screen layout |
| 0x007E7F15 | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x007E7F7C | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007E7FCF | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007E8044 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007E809E | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007E8155 | `NikePlus_Custom_Screen!` | Known | Screen layout |
| 0x007E81D1 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007E8248 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007E82DA | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x007E8358 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007E83B2 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007E8451 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007E84D2 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007E8529 | `NikePlus_Calibration_ChooseCalibration_Screen5` | Known | Screen layout |
| 0x007E85D6 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E8655 | `NikePlus_StartCalibration_Screen(` | Known | Screen layout |
| 0x007E8679 | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x007E86DF | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E875B | `NikePlus_StartWorkout_Screen$` | Known | Screen layout |
| 0x007E877B | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x007E87E6 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E885F | `NikePlus_Workout_Music_Screen` | Known | Screen layout |
| 0x007E88C6 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007E8921 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E89CD | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E8A5F | `NikePlus_StartWorkout_Screen$` | Known | Screen layout |
| 0x007E8A7F | `NikePlus_StartWorkout_Screen_Default#` | Known | Screen layout |
| 0x007E8AF3 | `NikePlus_StartCalibration_Screen(` | Known | Screen layout |
| 0x007E8B17 | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x007E8B84 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007E8C0C | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x007E8C9B | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E8CBC | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E8D59 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E8D7A | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E8E19 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E8E3A | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E8ED5 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E8EF6 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E8FC4 | `NikePlus_EquipmentAlert_Screen)` | Known | Screen layout |
| 0x007E905D | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007E907E | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007E911A | `NikePlus_History_BestWorkouts_Screen,` | Known | Screen layout |
| 0x007E9142 | `NikePlus_History_BestWorkouts_Screen_Default#` | Known | Screen layout |
| 0x007E91BE | `NikePlus_History_RecentWorkouts_Screen.` | Known | Screen layout |
| 0x007E91E8 | `NikePlus_History_RecentWorkouts_Screen_Default'` | Known | Screen layout |
| 0x007E926A | `NikePlus_History_WorkoutSummary_Screen+` | Known | Screen layout |
| 0x007E9294 | `NikePlus_History_WorkoutSummary_Screen_Last1` | Known | Screen layout |
| 0x007E931D | `NikePlus_NoData_Screen%` | Known | Screen layout |
| 0x007E9337 | `NikePlus_NoData_Screen_NoBestWorkouts2` | Known | Screen layout |
| 0x007E93BB | `NikePlus_NoData_Screen&` | Known | Screen layout |
| 0x007E93D5 | `NikePlus_NoData_Screen_NoSavedWorkouts` | Known | Screen layout |
| 0x007E94E5 | `NikePlus_History_Totals_Screen&` | Known | Screen layout |
| 0x007E9507 | `NikePlus_History_Totals_Screen_Default` | Known | Screen layout |
| 0x007E9579 | `NikePlus_History_DeleteActiveWorkout_Screen2` | Known | Screen layout |
| 0x007E95A8 | `NikePlus_History_DeleteActiveWorkout_Screen_Cancel` | Known | Screen layout |
| 0x007E961D | `NikePlus_History_DeleteActiveWorkout_Screen7` | Known | Screen layout |
| 0x007E964C | `NikePlus_History_DeleteActiveWorkout_Screen_LastWorkout` | Known | Screen layout |
| 0x007E96C4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007E9717 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007E976D | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007E9828 | `NikePlus_History_RecentWorkouts_ScreenB` | Known | Screen layout |
| 0x007E98BE | `NikePlus_History_WorkoutSummary_ScreenA` | Known | Screen layout |
| 0x007E9953 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x007E9A1F | `NikePlus_History_RecentWorkouts_ScreenB` | Known | Screen layout |
| 0x007E9AB5 | `NikePlus_History_WorkoutSummary_ScreenA` | Known | Screen layout |
| 0x007E9B4A | `NikePlus_History_Screen` | Known | Screen layout |
| 0x007E9C07 | `NikePlus_History_ScreenG` | Known | Screen layout |
| 0x007E9C93 | `NikePlus_History_RecentWorkouts_Screen` | Known | Screen layout |
| 0x007E9D13 | `NikePlus_History_DeleteAllWorkouts_Screen0` | Known | Screen layout |
| 0x007E9D40 | `NikePlus_History_DeleteAllWorkouts_Screen_Cancel#` | Known | Screen layout |
| 0x007E9DC0 | `NikePlus_History_WorkoutSummary_Screen.` | Known | Screen layout |
| 0x007E9DEA | `NikePlus_History_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007E9E9D | `NikePlus_History_ClearTotals_Screen+` | Known | Screen layout |
| 0x007E9EC4 | `NikePlus_History_ClearTotals_Screen_Default` | Known | Screen layout |
| 0x007E9F66 | `NikePlus_History_Totals_Screen` | Known | Screen layout |
| 0x007E9FF9 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007EA01A | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007EA089 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007EA0A7 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007EA113 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EA131 | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007EA1A1 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EA1BF | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007EA22B | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EA249 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007EA2DF | `NikePlus_Dynamic_Workout_Screen/` | Known | Screen layout |
| 0x007EA302 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x007EA37B | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007EA399 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007EA405 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EA423 | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007EA493 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EA4B1 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007EA51D | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EA53B | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007EA5D3 | `NikePlus_Dynamic_Workout_Screen,` | Known | Screen layout |
| 0x007EA5F6 | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x007EA66C | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007EA68A | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007EA6F6 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EA714 | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007EA784 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EA7A2 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007EA80E | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EA82C | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007EA8C3 | `NikePlus_Dynamic_Workout_Screen.` | Known | Screen layout |
| 0x007EA8E6 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x007EA95E | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007EA97C | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007EA9E8 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EAA06 | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007EAA76 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007EAA94 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007EAB00 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007EAB1E | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007EAB8E | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x007EABA7 | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x007EAC0A | `DemoMode_Screen` | Known | Screen layout |
| 0x007EAC1D | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x007EAC8A | `Debug_TestList_Screen` | Known | Screen layout |
| 0x007EACA3 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x007EAD16 | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x007EAD31 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x007EAE41 | `VoiceMemos_Context_PlayDelete_Screen)` | Known | Screen layout |
| 0x007EAE69 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x007EAEE0 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007EAFAC | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007EB01B | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007EB109 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007EB172 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007EB194 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007EB200 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007EB222 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007EB39E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007EB3BA | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007EB481 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007EB49C | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007EB4FF | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007EB562 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007EB5F9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007EB615 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007EB6DC | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007EB6F7 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007EB75A | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007EB7BD | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007EB855 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007EB871 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007EB938 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007EB953 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007EB9B6 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007EBA19 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007EBA96 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007EBB01 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007EBB6D | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007EBBDF | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007EBC4C | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007EBCB7 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x007EBD23 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007EBD8B | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007EBDF7 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007EBE6B | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007EBED9 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x007EBF52 | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x0080E9A0 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x0080EA25 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x0080ED12 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x009C256A | `VoiceMemos_Screen_InsufficientDiskSpace2` | Known | Screen layout |
| 0x009C3F0A | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x009C3F22 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x009C3F40 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x009C4007 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x009C4085 | `NowPlaying_Screen_Music` | Known | Screen layout |
| 0x009C40C6 | `MainMenu_Music_Screen_NoMusic` | Known | Screen layout |
| 0x009C40E4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x009C4102 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x009C411B | `NikePlus_Custom_Screen_Weight_ToWorkoutMusic` | Known | Screen layout |
| 0x009C4230 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x009C42E4 | `DeleteRental_Confirmation_Screen_Unwatched` | Known | Screen layout |
| 0x009C433A | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x009C4386 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x009C45E9 | `NowPlaying_Screen_ExtraInfoLoadFailed` | Known | Screen layout |
| 0x009C4644 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x009C465D | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x009C467B | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x009C46AA | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x009C46E2 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x009C4772 | `NikePlus_Playlists_Screen_Nested` | Known | Screen layout |
| 0x009C4B52 | `Video_Settings_TV_Screen_Standard` | Known | Screen layout |
| 0x009C4B84 | `Search_Main_Screen_WithKeyboard` | Known | Screen layout |
| 0x009C4BA4 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x009C4BE9 | `VoiceMemos_Screen_InsufficientDiskSpace` | Known | Screen layout |
| 0x009C4C4E | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x009C4C72 | `NikePlus_Custom_Screen_Distance` | Known | Screen layout |
| 0x009C4CCD | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x009C4D54 | `Video_Settings_TV_Screen_Wide` | Known | Screen layout |
| 0x009C4D9C | `CoverFlow_Screen_Backside` | Known | Screen layout |
| 0x009C83DE | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x009C85E3 | `NowPlaying_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x009C8608 | `Slideshow_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x009C86D8 | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x009C86F2 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x009C87EA | `RentalDeleted_Screen_Title` | Known | Screen layout |
| 0x009C8805 | `SingleRentalExpiring_Screen_Title` | Known | Screen layout |
| 0x009C8827 | `MultipleRentalsExpiring_Screen_Title` | Known | Screen layout |
| 0x009C884C | `DeleteRental_Screen_Title` | Known | Screen layout |
| 0x009C88EF | `VoiceMemos_DeleteAll_ScreenTitle` | Known | Screen layout |
| 0x009C898C | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x009C89CF | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x009C8AC1 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x009C8AE1 | `NikePlus_Custom_Screen_Time` | Known | Screen layout |
| 0x009C8C2C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x009C8D15 | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x009C8D2E | `Radio_Screen_Volume` | Known | Screen layout |
| 0x009C8D42 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x009C8D5F | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x009C8D7E | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x009C8E89 | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x009C8FF5 | `CoverFlow_Screen_BacksideForCapture` | Known | Screen layout |
| 0x009CA1CA | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x009CA2C2 | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x009CA2DD | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x009CA580 | `NikePlus_CalibrationComplete_Screen_Pacing` | Known | Screen layout |
| 0x009CA618 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x009CA64C | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x009CA689 | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x009CA79B | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x009CA8C9 | `NikePlus_Calibrate_ResetToDefault_Screen_Walking` | Known | Screen layout |
| 0x009CA9FB | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x009CAA14 | `NikePlus_Calibrate_ResetToDefault_Screen_Running` | Known | Screen layout |
| 0x009CAA64 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x009CAA8A | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x009D06A3 | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x009D0712 | `TVOutSlideshow_Screen_Playing` | Known | Screen layout |
| 0x009D0730 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x009D0786 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x009D07F0 | `NikePlus_NowRunning_Screen_Basic_PowerSong` | Known | Screen layout |
| 0x009D081B | `NikePlus_NowRunning_Screen_Distance_PowerSong` | Known | Screen layout |
| 0x009D0849 | `NikePlus_NowRunning_Screen_Time_PowerSong` | Known | Screen layout |
| 0x009D0896 | `NikePlus_NowRunning_Screen_Calories_PowerSong` | Known | Screen layout |
| 0x009D0913 | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x009D097E | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x009D09FE | `Extras_Screen_Debug` | Known | Screen layout |
| 0x009D0B08 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x009D0B28 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x009D109A | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x009D10B5 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x009D10C8 | `MainMenu_List_ScreenLock` | Known | Screen layout |
| 0x009D10E1 | `ExtrasMenu_ScreenLock` | Known | Screen layout |
| 0x009D1164 | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x009D1185 | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x009D1230 | `NikePlus_StartCalibration_Screen_Walk` | Known | Screen layout |
| 0x009D12B8 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x009D12DA | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x009D13E1 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x009D1421 | `VoiceMemos_Screen_DeletAllAsk` | Known | Screen layout |
| 0x009D143F | `NowPlaying_Screen_TVOutAsk` | Known | Screen layout |
| 0x009D159B | `NowPlaying_Screen_Initial` | Known | Screen layout |
| 0x009D15B5 | `NowPlaying_Screen_Video_TVOut_Initial` | Known | Screen layout |
| 0x009D1887 | `NikePlus_History_DeleteAllWorkouts_Screen_Cancel` | Known | Screen layout |
| 0x009D18B8 | `NikePlus_History_DeleteActiveWorkout_Screen_Cancel` | Known | Screen layout |
| 0x009D26DA | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x009D275B | `RemoteUI_Screen` | Known | Screen layout |
| 0x009D276B | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x009D2783 | `Video_Settings_TV_Screen` | Known | Screen layout |
| 0x009D279C | `NikePlus_NoData_Screen` | Known | Screen layout |
| 0x009D27B3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x009D27CA | `NikePlus_Workout_Music_Screen` | Known | Screen layout |
| 0x009D27E8 | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x009D280C | `AddressViewer_PartialLoad_Screen` | Known | Screen layout |
| 0x009D282D | `NikePlus_ActivityStopped_Screen` | Known | Screen layout |
| 0x009D284D | `Alarms_Alarm_Clock_Triggered_Screen` | Known | Screen layout |
| 0x009D2871 | `Alarms_Alarm_Triggered_Screen` | Known | Screen layout |
| 0x009D288F | `Unsupported_Screen` | Known | Screen layout |
| 0x009D28A2 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x009D28C0 | `LockediPod_Screen` | Known | Screen layout |
| 0x009D28D2 | `DiskMode_Screen` | Known | Screen layout |
| 0x009D28E2 | `DemoMode_Screen` | Known | Screen layout |
| 0x009D28F2 | `Notes_Image_Screen` | Known | Screen layout |
| 0x009D2905 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x009D2923 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x009D2939 | `NowPlaying_Idle_Screen` | Known | Screen layout |
| 0x009D2950 | `Game_Screen` | Known | Screen layout |
| 0x009D295C | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x009D2979 | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x009D2992 | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x009D29B3 | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x009D29D8 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x009D29EB | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x009D2A08 | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x009D2A29 | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x009D2A4E | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x009D2A65 | `Notes_Loading_Screen` | Known | Screen layout |
| 0x009D2A7A | `NikePlus_SensorSearching_Screen` | Known | Screen layout |
| 0x009D2A9A | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x009D2AB9 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x009D2AD1 | `NikePlus_Remote_Unlinking_Screen` | Known | Screen layout |
| 0x009D2AF2 | `SettingsMenus_AdjustScrolling_Screen` | Known | Screen layout |
| 0x009D2B17 | `Game_Running_Screen` | Known | Screen layout |
| 0x009D2B2B | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x009D2B46 | `Stopwatch_Screen` | Known | Screen layout |
| 0x009D2B57 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x009D2B6E | `Clock_Screen` | Known | Screen layout |
| 0x009D2B7B | `NikePlus_Calibration_CalibrateWalk_Screen` | Known | Screen layout |
| 0x009D2BA5 | `TVOutSlideshowAsk_Screen` | Known | Screen layout |
| 0x009D2BBE | `Settings_Legal_Screen` | Known | Screen layout |
| 0x009D2BD4 | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x009D2BF2 | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x009D2C0E | `ToDo_Item_Screen` | Known | Screen layout |
| 0x009D2C1F | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x009D2C36 | `DemoMode_Main_Screen` | Known | Screen layout |
| 0x009D2C4B | `Search_Main_Screen` | Known | Screen layout |
| 0x009D2C5E | `AddressViewer_Main_Screen` | Known | Screen layout |
| 0x009D2C78 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x009D2C8D | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x009D2CA3 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x009D2CBD | `Clock_Region_Screen` | Known | Screen layout |
| 0x009D2CD1 | `RentalDeleted_Notification_Screen` | Known | Screen layout |
| 0x009D2CF3 | `SingleRentalExpiring_Notification_Screen` | Known | Screen layout |
| 0x009D2D1C | `MultipleRentalsExpiring_Notification_Screen` | Known | Screen layout |
| 0x009D2D48 | `RentalError_Notification_Screen` | Known | Screen layout |
| 0x009D2D68 | `DeleteRental_Confirmation_Screen` | Known | Screen layout |
| 0x009D2D89 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x009D2DA1 | `NikePlus_EndCalibration_Screen` | Known | Screen layout |
| 0x009D2DC0 | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x009D2DEE | `NikePlus_StartCalibration_Screen` | Known | Screen layout |
| 0x009D2E0F | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x009D2E2D | `NikePlus_Calibration_CalibrateRun_Screen` | Known | Screen layout |
| 0x009D2E56 | `Video_Settings_Fit_To_Screen` | Known | Screen layout |
| 0x009D2E73 | `RentalInfo_Screen` | Known | Screen layout |
| 0x009D2E85 | `Radio_Screen` | Known | Screen layout |
| 0x009D2E92 | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x009D2EAC | `ChargingMode_LowPower_Screen` | Known | Screen layout |
| 0x009D2EC9 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x009D2EE3 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x009D2EFD | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x009D2F17 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x009D2F30 | `NikePlus_CalibrationCompleteError_Screen` | Known | Screen layout |
| 0x009D2F59 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x009D2F70 | `Extras_Screen` | Known | Screen layout |
| 0x009D2F7E | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x009D2F9B | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x009D2FBD | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x009D2FD6 | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x009D2FF4 | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x009D300D | `Video_Settings_Screen` | Known | Screen layout |
| 0x009D3023 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x009D303C | `SettingsMenus_AudiobookSettings_Screen` | Known | Screen layout |
| 0x009D3063 | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x009D3089 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x009D309F | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x009D30B7 | `NikePlus_Songs_Screen` | Known | Screen layout |
| 0x009D30CD | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x009D30F0 | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x009D310D | `MediaLists_Rentals_Screen` | Known | Screen layout |
| 0x009D3127 | `NikePlus_History_Totals_Screen` | Known | Screen layout |
| 0x009D3146 | `NikePlus_History_ClearTotals_Screen` | Known | Screen layout |
| 0x009D316A | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x009D318E | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x009D31A7 | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x009D31C9 | `ToDo_Instructions_Screen` | Known | Screen layout |
| 0x009D31E2 | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x009D31FE | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x009D3218 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x009D3239 | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x009D3255 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x009D326D | `VoiceMemos_Screen` | Known | Screen layout |
| 0x009D327F | `No_Photos_Screen` | Known | Screen layout |
| 0x009D3290 | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x009D32AA | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x009D32C6 | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x009D32EA | `NikePlus_CalibrationCompleteSuccess_Screen` | Known | Screen layout |
| 0x009D3315 | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x009D3335 | `ChargingMode_Progress_Screen` | Known | Screen layout |
| 0x009D3352 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x009D3368 | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x009D3383 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x009D339F | `NikePlus_Playlists_Screen` | Known | Screen layout |
| 0x009D33B9 | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x009D33DB | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x009D33FC | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x009D3416 | `NikePlus_History_DeleteAllWorkouts_Screen` | Known | Screen layout |
| 0x009D3440 | `NikePlus_History_RecentWorkouts_Screen` | Known | Screen layout |
| 0x009D3467 | `NikePlus_History_BestWorkouts_Screen` | Known | Screen layout |
| 0x009D348C | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x009D34A6 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x009D34C5 | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x009D34E6 | `NikePlus_Calibrate_ResetToDefault_Screen` | Known | Screen layout |
| 0x009D350F | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x009D3527 | `NoContent_Screen` | Known | Screen layout |
| 0x009D3538 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x009D354E | `FirstBoot_Screen` | Known | Screen layout |
| 0x009D355F | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x009D3575 | `NikePlus_EquipmentAlert_Screen` | Known | Screen layout |
| 0x009D3594 | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x009D35AA | `Notes_List_Screen` | Known | Screen layout |
| 0x009D35BC | `Debug_TestList_Screen` | Known | Screen layout |
| 0x009D35D2 | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x009D35F3 | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x009D360D | `NikePlus_Dynamic_Workout_Screen` | Known | Screen layout |
| 0x009D362D | `NikePlus_EndPausedWorkout_Screen` | Known | Screen layout |
| 0x009D364E | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x009D3669 | `NikePlus_ResumeWorkout_Screen` | Known | Screen layout |
| 0x009D3687 | `NikePlus_History_DeleteActiveWorkout_Screen` | Known | Screen layout |
| 0x009D36B3 | `NikePlus_StartWorkout_Screen` | Known | Screen layout |
| 0x009D36D0 | `TODOS_Menu_Screen` | Known | Screen layout |
| 0x009D36E2 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x009D36F8 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x009D3714 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x009D3729 | `Games_Menu_Screen` | Known | Screen layout |
| 0x009D373B | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x009D374E | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x009D376D | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x009D378C | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x009D37B0 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x009D37C6 | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x009D37E4 | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x009D3807 | `TVOutSlideshow_Screen` | Known | Screen layout |
| 0x009D381D | `CoverFlow_Screen` | Known | Screen layout |
| 0x009D382E | `Calendar_Day_Screen` | Known | Screen layout |
| 0x009D3842 | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x009D3864 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x009D387C | `Stopwatch_SessionSummary_Screen` | Known | Screen layout |
| 0x009D389C | `NikePlus_End_WorkoutSummary_Screen` | Known | Screen layout |
| 0x009D38BF | `NikePlus_History_WorkoutSummary_Screen` | Known | Screen layout |
| 0x009D38E6 | `SettingsMenu_AboutWithAccessory_Screen` | Known | Screen layout |
| 0x009D390D | `NikePlus_History_Screen` | Known | Screen layout |
| 0x009D3925 | `Clock_NorthAmerica_City_Screen` | Known | Screen layout |
| 0x009D3944 | `Clock_SouthAmerica_City_Screen` | Known | Screen layout |
| 0x009D3963 | `Clock_Africa_City_Screen` | Known | Screen layout |
| 0x009D397C | `Clock_Australia_City_Screen` | Known | Screen layout |
| 0x009D3998 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x009D39AF | `Clock_Pacific_City_Screen` | Known | Screen layout |
| 0x009D39C9 | `Clock_Atlantic_City_Screen` | Known | Screen layout |
| 0x009D39E4 | `Clock_Europe_City_Screen` | Known | Screen layout |
| 0x009D3AD8 | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x009D3B29 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x009D3B4C | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x009D3B74 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x009D3F00 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x009D4003 | `RentalInfo_Screen_ExpiringSoon` | Known | Screen layout |
| 0x009D4059 | `RentalInfo_Screen_NoAlbumArt_ExpiringSoon` | Known | Screen layout |
| 0x009D4163 | `NikePlus_StartCalibration_Screen_Run` | Known | Screen layout |
| 0x009D444D | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x009D44A3 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x009D45F4 | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x009D4611 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x009D4A42 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x009D4B64 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x009D4B86 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x009D4CBE | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x009D4CDD | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x009D53CD | `ChargingMode_ScreenLayout_LowPower` | Known | Screen layout |
| 0x009D5D8D | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x009D5EEB | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x009D5FA2 | `NikePlus_NowRunning_Screen_Calories` | Known | Screen layout |
| 0x009D5FC6 | `NikePlus_Custom_Screen_Calories` | Known | Screen layout |
| 0x009D605F | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x009D607D | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x009D609D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x009D61A8 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x009D61C4 | `Extras_Screen_Games` | Known | Screen layout |
| 0x009D62CA | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x009D62E9 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x009D6305 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x009D63F0 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x009D650C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x009D66DA | `MainMenu_Music_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009D66FD | `MainMenus_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009D6720 | `NoContent_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009D675A | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x009D6779 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x009D679A | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x009D6891 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x009D68AE | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x009D692D | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x009D6A11 | `MainMenu_Music_Screen_NoCompilations` | Known | Screen layout |
| 0x009D6A36 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x009D6BDF | `MainMenu_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x009D6C02 | `MainMenu_Videos_Screen_NoMusicVideos` | Known | Screen layout |
| 0x009D6C27 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x009D6C46 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x009D6C65 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x009D6C86 | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x009D6CC4 | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x009D6CE5 | `Extras_Screen_VoiceMemos` | Known | Screen layout |
| 0x009D6D50 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x009D6D82 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x009D6DA1 | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x009D6E4E | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x009D6EBA | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x009D6FB3 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x009D6FCF | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x009D7052 | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x009D706D | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x009D708E | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x009D713D | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x009D7171 | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x009D7192 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x009D7250 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x009D7271 | `MainMenu_Videos_Screen_NoPlaylists` | Known | Screen layout |
| 0x009D7294 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x009D72E3 | `MainMenu_Music_Screen_NoArtists` | Known | Screen layout |
| 0x009D7353 | `NikePlus_NoData_Screen_NoSavedWorkouts` | Known | Screen layout |
| 0x009D7405 | `NikePlus_NoData_Screen_NoBestWorkouts` | Known | Screen layout |
| 0x009D74B2 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x009D74D1 | `Extras_Screen_NikePlus` | Known | Screen layout |
| 0x009D7621 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x009D7640 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x009D7661 | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x009D7B38 | `NikePlus_Custom_Screen_Weight` | Known | Screen layout |
| 0x009D7BAD | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x009D7C60 | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x009D7CDA | `Notes_Contents_Screen_Alt` | Known | Screen layout |
| 0x009D7CF4 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x009D7DA0 | `VoiceMemos_Screen_InsufficientDiskSpace2_Default` | Known | Screen layout |
| 0x009D7E52 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x009D7EF7 | `VoiceMemos_Screen_InsufficientDiskSpace_Default` | Known | Screen layout |
| 0x009D7F27 | `NowPlaying_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x009D7F54 | `Slideshow_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x009D8E2E | `Radio_Screen_Tuning_Default` | Known | Screen layout |
| 0x009D8EBA | `VoiceMemos_Screen_DeletAllAsk_Default` | Known | Screen layout |
| 0x009D8EE0 | `NowPlaying_Screen_TVOutAsk_Default` | Known | Screen layout |
| 0x009D8F17 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x009D8F3D | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x009D8F5B | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x009D8F87 | `AddressViewer_PartialLoad_Screen_Default` | Known | Screen layout |
| 0x009D8FB0 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x009D8FD8 | `Alarms_Alarm_Clock_Triggered_Screen_Default` | Known | Screen layout |
| 0x009D9004 | `Alarms_Alarm_Triggered_Screen_Default` | Known | Screen layout |
| 0x009D902A | `Unsupported_Screen_Default` | Known | Screen layout |
| 0x009D9045 | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x009D906B | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x009D9083 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x009D909E | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x009D90BB | `Game_Screen_Default` | Known | Screen layout |
| 0x009D90CF | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x009D90F5 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x009D9116 | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x009D913F | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x009D9169 | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x009D9196 | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x009D91BF | `Notes_Loading_Screen_Default` | Known | Screen layout |
| 0x009D91DC | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x009D9204 | `NikePlus_Remote_Unlinking_Screen_Default` | Known | Screen layout |
| 0x009D922D | `Clock_Screen_Default` | Known | Screen layout |
| 0x009D9242 | `TVOutSlideshowAsk_Screen_Default` | Known | Screen layout |
| 0x009D9263 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x009D9281 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x009D92A7 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x009D92CB | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x009D92E4 | `AddressViewer_Main_Screen_Default` | Known | Screen layout |
| 0x009D9306 | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x009D9323 | `MainMenus_Main_Screen_Default` | Known | Screen layout |
| 0x009D9341 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x009D935E | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x009D937A | `RentalDeleted_Notification_Screen_Default` | Known | Screen layout |
| 0x009D93A4 | `SingleRentalExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x009D93D5 | `MultipleRentalsExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x009D9409 | `RentalError_Notification_Screen_Default` | Known | Screen layout |
| 0x009D9431 | `DeleteRental_Confirmation_Screen_Default` | Known | Screen layout |
| 0x009D945A | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x009D9486 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x009D94AD | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x009D94D6 | `RentalInfo_Screen_Default` | Known | Screen layout |
| 0x009D94F0 | `Radio_Screen_Default` | Known | Screen layout |
| 0x009D9505 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x009D9527 | `NikePlus_CalibrationCompleteError_Screen_Default` | Known | Screen layout |
| 0x009D9558 | `Extras_Screen_Default` | Known | Screen layout |
| 0x009D956E | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x009D9594 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x009D95B5 | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x009D95D3 | `NikePlus_Settings_Screen_Default` | Known | Screen layout |
| 0x009D95F4 | `NikePlus_Songs_Screen_Default` | Known | Screen layout |
| 0x009D9612 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x009D9634 | `NikePlus_History_Totals_Screen_Default` | Known | Screen layout |
| 0x009D965B | `NikePlus_History_ClearTotals_Screen_Default` | Known | Screen layout |
| 0x009D9687 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x009D96B3 | `ToDo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009D96D4 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x009D96F8 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x009D971A | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x009D973E | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x009D975D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x009D9776 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x009D9798 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x009D97BC | `NikePlus_CalibrationCompleteSuccess_Screen_Default` | Known | Screen layout |
| 0x009D97EF | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x009D980D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x009D9831 | `NikePlus_Playlists_Screen_Default` | Known | Screen layout |
| 0x009D9853 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x009D987D | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x009D98A6 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x009D98C8 | `NikePlus_History_RecentWorkouts_Screen_Default` | Known | Screen layout |
| 0x009D98F7 | `NikePlus_History_BestWorkouts_Screen_Default` | Known | Screen layout |
| 0x009D9924 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x009D9944 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x009D9962 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x009D997B | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x009D9999 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x009D99B3 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x009D99D1 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x009D99FA | `NikePlus_EndWorkout_Screen_Default` | Known | Screen layout |
| 0x009D9A1D | `NikePlus_ResumeWorkout_Screen_Default` | Known | Screen layout |
| 0x009D9A43 | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x009D9A68 | `TODOS_Menu_Screen_Default` | Known | Screen layout |
| 0x009D9A82 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x009D9AA0 | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x009D9ABD | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x009D9AD7 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x009D9AF2 | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x009D9B11 | `Debug_MainMenu_Screen_Default` | Known | Screen layout |
| 0x009D9B2F | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x009D9B4D | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x009D9B66 | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x009D9B82 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x009D9BAC | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x009D9BCC | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x009D9BF4 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x009D9C1F | `NikePlus_History_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x009D9C4E | `NikePlus_History_Screen_Default` | Known | Screen layout |
| 0x009D9C6E | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009D9C95 | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009D9CBC | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x009D9CDD | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x009D9D01 | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x009D9D20 | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x009D9D42 | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x009D9D65 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x009D9DA2 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x009D9E30 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x009D9E60 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x009D9E82 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x009D9EF3 | `RentalInfo_Screen_NoAlbumArt_Default` | Known | Screen layout |
| 0x009D9F18 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x009DA644 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009DA670 | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009DA6B5 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x009DA6DD | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x009DA6FE | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x009DA71F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x009DA745 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x009DA762 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x009DA784 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x009DA7A8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x009DA7CC | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x009DA831 | `NikePlus_History_WorkoutSummary_Screen_Last` | Known | Screen layout |
| 0x009DA9D2 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x009DAA42 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x009DAA93 | `NowPlaying_Screen_Video_TVOut` | Known | Screen layout |
| 0x009DABA6 | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x009DAC03 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x009DAC52 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x009DAD19 | `NikePlus_History_DeleteActiveWorkout_Screen_LastWorkout` | Known | Screen layout |
| 0x009DAE62 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x009DAE89 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x009DB2AF | `NikePlus_EndPausedWorkout_Screen_BasicSave_Layout` | Known | Screen layout |
| 0x009DB2E1 | `NikePlus_EndPausedWorkout_Screen_DistanceSave_Layout` | Known | Screen layout |
| 0x009DB316 | `NikePlus_EndPausedWorkout_Screen_TimeSave_Layout` | Known | Screen layout |
| 0x009DB347 | `NikePlus_EndPausedWorkout_Screen_CaloriesSave_Layout` | Known | Screen layout |
| 0x009DB5FF | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x009DB7FC | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x009DBAA0 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x009DBDBD | `Lock_Screen_Default_Layout` | Known | Screen layout |
| 0x009DBE53 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x009DBE7A | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x009DC096 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x009DC170 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x009DC1D7 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009DC201 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009DF0A9 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x009DF0F5 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x009DF1D3 | `MainMenu_List_ScreenLock_x` | Known | Screen layout |
| 0x009DF4A1 | `RentalInfo_Screen_ExpiresToday` | Known | Screen layout |
| 0x009DF4F7 | `RentalInfo_Screen_NoAlbumArt_ExpiresToday` | Known | Screen layout |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000904B | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | Known | RTOS |
| 0x002BE410 | `  K - RTXC` | Known | RTOS |
| 0x002BF3F8 | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | Known | RTOS |
| 0x009C0F24 | `Returning from RTXCBug` | Known | RTOS |

---

## 9. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000DE4E4 | `HostOSTask` | Known | RTOS task thread |
| 0x0013EAFC | `MP3ExampleTask` | Known | RTOS task thread |
| 0x00146FA0 | `USBDeviceTask` | Known | RTOS task thread |
| 0x0015147C | `DiskReaderTask` | Known | RTOS task thread |
| 0x00161038 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0016104C | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x001B6DB0 | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x001F81B0 | `MeCCAIOTask` | Known | RTOS task thread |
| 0x0022EEB8 | `StreamCacheMassStorageManagerTimeOutTask` | Known | RTOS task thread |
| 0x0022F034 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x002B14C8 | `FirewireTask` | Known | RTOS task thread |
| 0x002B14DC | `TouchwheelTask` | Known | RTOS task thread |
| 0x002B14F0 | `AudioOutStateTask` | Known | RTOS task thread |
| 0x002B151C | `DiskMgrTask` | Known | RTOS task thread |
| 0x002B152C | `HoldSwitchTask` | Known | RTOS task thread |
| 0x002B1540 | `TopPlugTask` | Known | RTOS task thread |
| 0x002B1550 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x002B15C8 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x002B15F0 | `AlarmTask` | Known | RTOS task thread |
| 0x002B160F | `"USBAudioTask` | Known | RTOS task thread |
| 0x002BEAB0 | `Undefined Task` | Known | RTOS task thread |
| 0x003C4B90 | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x003C8DE8 | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x003D1440 | `MeCCARecordingTask` | Known | RTOS task thread |
| 0x00914478 | `WaveFileDebugTask` | Known | RTOS task thread |

---

## 10. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00268508 | `Channel Reserved` | Known | Logging channel |
| 0x0026851C | `Channel AppBoot` | Known | Logging channel |
| 0x0026852C | `Channel BufferedSongReading` | Known | Logging channel |
| 0x00268548 | `Channel PrefsWriting` | Known | Logging channel |
| 0x00268560 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x00268580 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x00268598 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x002685B4 | `Channel TestLogging` | Known | Logging channel |
| 0x002685C8 | `Channel AppFileLoading` | Known | Logging channel |
| 0x002685E0 | `Channel VCardReading` | Known | Logging channel |
| 0x002685F8 | `Channel LongSongScanning` | Known | Logging channel |
| 0x0026866C | `Channel VoiceRecording` | Known | Logging channel |
| 0x00268684 | `Channel PhotoImporting` | Known | Logging channel |
| 0x0026869C | `Channel Notes` | Known | Logging channel |
| 0x002686AC | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x002686C8 | `Channel DiskMode` | Known | Logging channel |
| 0x002686DC | `Channel Firewire` | Known | Logging channel |
| 0x002686F0 | `Channel USB` | Known | Logging channel |
| 0x00268710 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x00268728 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0008D700 | `gamedata_RW` | Known | Game system |
| 0x0008D71C | `gamedata_ShareRW` | Known | Game system |
| 0x0008D730 | `games_RO` | Known | Game system |
| 0x009C0F7E | `iPod_Control/games_RO/` | Known | Game system |
| 0x009C0F95 | `Resources/Games/games_RO/` | Known | Game system |
| 0x009CDF76 | `GamesMenu_StatusBar_String` | Known | Game system |
| 0x009CE6B0 | `AboutScreen_Games_String` | Known | Game system |
| 0x009D61D8 | `MainMenu_List_Games` | Known | Game system |
| 0x009D61EC | `ExtrasMenu_Games` | Known | Game system |
| 0x009DF242 | `MainMenu_List_Games_x` | Known | Game system |

---

## 12. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009E884 | `adrmmp4a` | Known | DRM system |
| 0x0014EA94 | `AppleDRMVersion` | Known | DRM system |
| 0x0014EB34 | `AppleDRM` | Known | DRM system |
| 0x0014FD30 | `AppleVideoDRM` | Known | DRM system |
| 0x001531EC | `tx3gdrmsp608aavdmp4aesds` | Known | DRM system |
| 0x00207250 | `drmttx3g` | Known | DRM system |
| 0x009C1363 | `DRMLevel` | Known | DRM system |

---

## 13. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00035ED8 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00035EF0 | `iPod_Control\iTunes\firsttime` | Known | iTunes database |
| 0x00058350 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x00058378 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000614B0 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00089760 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x0008D690 | `iPod_Control/iTunes/iTunesDB.p7b` | Known | iTunes database |
| 0x000AA9E8 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x000AABD0 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000B3608 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000B4AAC | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000B4BAC | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00137294 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x003BDF10 | `iTunesDB` | Known | iTunes database |
| 0x003BDF1C | `iPod_Control\iTunes\` | Known | iTunes database |

---

## 14. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0005F4FC | `NANDDRIVERSIGN` | Known | Hardware |
| 0x00060038 | `NANDDRIVERSIGN` | Known | Hardware |
| 0x00060E58 | `[FTL:MSG] Apple NAND Driver (AND) 0x%08x` | Known | Hardware |
| 0x00060F70 | `[FTL:MSG] Valid Signature not found! Re-initializing NAND!` | Known | Hardware |
| 0x00136AF8 | `NAND FLASH DRIVE` | Known | Hardware |
| 0x0014F004 | `FireWireGUID` | Known | FireWire |
| 0x0014F014 | `FireWireVersion` | Known | FireWire |
| 0x0014F6F8 | `FireWire` | Known | FireWire |
| 0x002CF21C | `[FIL:ERR] No recognized NAND found (0x%X, 0x%X) (line:%d)!` | Known | Hardware |
| 0x0091B9E0 | `[FTL:WRN] Recovering NAND Data Structures - this will take some time!` | Known | Hardware |
| 0x0091CF24 | `[FIL:WRN]  FNAND_GetStruct 0x%X is not identified is FIL data struct identifier!` | Known | Hardware |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00756902 | `Settings_RadioRegionsMenu_Layout"` | Known | FM Radio |
| 0x0075698B | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x0080CC3C | `Radio Regions` | Known | FM Radio |
| 0x008677BC | `Radio-Regionen` | Known | FM Radio |
| 0x009CB14D | `Settings_RadioRegions_Australia_String` | Known | FM Radio |
| 0x009CB174 | `Settings_RadioRegions_Asia_String` | Known | FM Radio |
| 0x009CC356 | `Settings_RadioRegions_Europe_String` | Known | FM Radio |
| 0x009CD8C6 | `Settings_RadioRegions_Japan_String` | Known | FM Radio |
| 0x009CE4CD | `Settings_RadioRegions_Americas_String` | Known | FM Radio |
| 0x009CEBAF | `SettingsMenu_RadioRegions_String` | Known | FM Radio |
| 0x009D2599 | `RadioRegionsMenuItem` | Known | FM Radio |
| 0x009D699A | `SettingsMenu_ListItem_RadioRegions` | Known | FM Radio |
| 0x009DB905 | `SettingsInfo_Template_RadioRegions_Layout` | Known | FM Radio |
| 0x009DB92F | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x009DC057 | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 16. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x008A6038 | `Fotocamera` | Known | Camera |
| 0x008A61C8 | `Collega fotocamera o inserisci scheda` | Known | Camera |
| 0x008A6240 | `Fotocamera non supportata` | Known | Camera |
| 0x008C497C | `Camera` | Known | Camera |
| 0x008C4B08 | `Sluit camera of kaart aan` | Known | Camera |
| 0x008C4B74 | `Camera niet ondersteund` | Known | Camera |
| 0x009CB196 | `Photos_DevName_Camera_String` | Known | Camera |

---

## 17. Pedometer

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0080DBE8 | `Step away from all other sensors.` | Known | Pedometer |
| 0x0080DDCC | `Step away from all other remotes.` | Known | Pedometer |
| 0x009DF5AA | `NikePlus_Remote_Step_Away` | Known | Pedometer |
| 0x009DF5C4 | `NikePlus_Step_Away` | Known | Pedometer |

---

## 18. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00035EC4 | `iPod_Control` | Filesystem Path |  |
| 0x00035F30 | `iPod_Control\Device` | Filesystem Path |  |
| 0x00045944 | `iPod_Control\Device` | Filesystem Path |  |
| 0x000479CC | `iPod_Control` | Filesystem Path |  |
| 0x00048034 | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x00058330 | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path |  |
| 0x0005AE94 | `iPod_Control\Music\` | Filesystem Path |  |
| 0x00061330 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x0009793C | `iPod_Control` | Filesystem Path |  |
| 0x0009794C | `Resources/Games` | Filesystem Path |  |
| 0x0009795C | `iPod_Control/%s%s%s` | Filesystem Path |  |
| 0x00101814 | `iPod_Control/Device/alarms` | Filesystem Path |  |
| 0x00111B14 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00113018 | `iPod_Control/Device` | Filesystem Path |  |
| 0x0011302C | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00132074 | `iPod_Control/Device/Users` | Filesystem Path |  |
| 0x0016271C | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x00162978 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x0016E6F0 | `Resources/UI/active.bin` | Filesystem Path |  |
| 0x0016E708 | `Resources/UI/` | Filesystem Path |  |
| 0x001909E0 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x00191CC4 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path |  |
| 0x00191CEC | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001BA768 | `iPod_Control/Device/PlayCounts` | Filesystem Path |  |
| 0x001D1AC0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1B70 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1CEC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1E84 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D1F2C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D20DC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2180 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2224 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D22C8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D236C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D241C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D24C0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2564 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2614 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D26C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2774 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D28E0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2990 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2A40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2AE4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2B94 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2C88 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2D2C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2DE0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2E9C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D2F4C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3070 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D312C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D31DC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3398 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D345C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D350C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D35C8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3704 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D37D0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D388C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3930 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D39D4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3A90 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3B4C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3C14 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3CB8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3D80 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3E48 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3EF8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D3FC0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4088 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4138 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D41E8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D42AC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D435C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D440C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D44BC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4590 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4664 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4764 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4844 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D494C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001D4A38 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003BDF8E | `iPod_Control/Device` | Filesystem Path |  |
| 0x003C4430 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x003C6F84 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003C7332 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003C73F0 | `iPod_Control/Device/Trainer/Workouts/Empeds/` | Filesystem Path |  |
| 0x003C8F54 | `Resources/Fonts` | Filesystem Path |  |
| 0x003D140C | `iPod_Control:Device:voices:` | Filesystem Path |  |
| 0x003D33CE | `Resources/TrainerTemplates` | Filesystem Path |  |
| 0x003D33E9 | `iPod_Control/Device/Trainer/TrainerTemplates` | Filesystem Path |  |
| 0x003D3A3C | `iPod_Control/Device/Trainer/Workouts/Empeds/` | Filesystem Path |  |
| 0x009C0E59 | `Resources/Games/` | Filesystem Path |  |
| 0x009C1245 | `iPod_Control/Device` | Filesystem Path |  |
| 0x009C1259 | `iPod_Control/Device/accessories.reg` | Filesystem Path |  |
| 0x009C12DA | `iPod_Control/Tones` | Filesystem Path |  |

---

## 19. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00912A20 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x00917768 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftu` | Build Path |  |
| 0x009177C0 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fts` | Build Path |  |
| 0x00917818 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fto` | Build Path |  |
| 0x0091B378 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x0091B3EC | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x0091C008 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x0091C73C | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x0091CB88 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x00923A48 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\` | Build Path |  |
| 0x009245C4 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfd` | Build Path |  |
| 0x009257C0 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrg` | Build Path |  |
| 0x00925818 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrc` | Build Path |  |
| 0x00925870 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfro` | Build Path |  |
| 0x00925BB4 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1` | Build Path |  |
| 0x00934F5C | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttc` | Build Path |  |
| 0x009351D8 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype` | Build Path |  |
| 0x00935744 | `c:\bwa\N46FirmwareWin-359\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1` | Build Path |  |

---

## 20. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00095908 | `Acoustic` | EQ Preset |  |
| 0x00095914 | `Bass Booster` | EQ Preset |  |
| 0x00095934 | `Classical` | EQ Preset |  |
| 0x00095940 | `Dance` | EQ Preset |  |
| 0x00095950 | `Electronic` | EQ Preset |  |
| 0x00095964 | `Hip Hop` | EQ Preset |  |
| 0x0009596C | `Jazz` | EQ Preset |  |
| 0x00095974 | `Latin` | EQ Preset |  |
| 0x0009597C | `Loudness` | EQ Preset |  |
| 0x00095988 | `Lounge` | EQ Preset |  |
| 0x00095990 | `Piano` | EQ Preset |  |
| 0x000959A4 | `Rock` | EQ Preset |  |
| 0x000959AC | `Small Speakers` | EQ Preset |  |
| 0x000959BC | `Spoken Word` | EQ Preset |  |
| 0x000959C8 | `Treble Booster` | EQ Preset |  |
| 0x00095A14 | `Vocal Booster` | EQ Preset |  |
| 0x0080CF2C | `Acoustic` | EQ Preset |  |
| 0x0080CF38 | `Bass Booster` | EQ Preset |  |
| 0x0080CF58 | `Classical` | EQ Preset |  |
| 0x0080CF64 | `Dance` | EQ Preset |  |
| 0x0080CF74 | `Electronic` | EQ Preset |  |
| 0x0080CF88 | `Hip Hop` | EQ Preset |  |
| 0x0080CF90 | `Jazz` | EQ Preset |  |
| 0x0080CF98 | `Latin` | EQ Preset |  |
| 0x0080CFA0 | `Loudness` | EQ Preset |  |
| 0x0080CFAC | `Lounge` | EQ Preset |  |
| 0x0080CFB4 | `Piano` | EQ Preset |  |
| 0x0080CFC4 | `Rock` | EQ Preset |  |
| 0x0080CFCC | `Small Speakers` | EQ Preset |  |
| 0x0080CFDC | `Spoken Word` | EQ Preset |  |
| 0x0080CFE8 | `Treble Booster` | EQ Preset |  |
| 0x0080D008 | `Vocal Booster` | EQ Preset |  |
| 0x00855544 | `Acoustic` | EQ Preset |  |
| 0x00855550 | `Bass Booster` | EQ Preset |  |
| 0x00855570 | `Classical` | EQ Preset |  |
| 0x0085557C | `Dance` | EQ Preset |  |
| 0x0085558C | `Electronic` | EQ Preset |  |
| 0x008555A0 | `Hip Hop` | EQ Preset |  |
| 0x008555A8 | `Jazz` | EQ Preset |  |
| 0x008555B0 | `Latin` | EQ Preset |  |
| 0x008555B8 | `Loudness` | EQ Preset |  |
| 0x008555C4 | `Lounge` | EQ Preset |  |
| 0x008555CC | `Piano` | EQ Preset |  |
| 0x008555DC | `Rock` | EQ Preset |  |
| 0x008555E4 | `Small Speakers` | EQ Preset |  |
| 0x008555F4 | `Spoken Word` | EQ Preset |  |
| 0x00855600 | `Treble Booster` | EQ Preset |  |
| 0x00855620 | `Vocal Booster` | EQ Preset |  |
| 0x0085E5F0 | `Acoustic` | EQ Preset |  |
| 0x0085E5FC | `Bass Booster` | EQ Preset |  |
| 0x0085E61C | `Classical` | EQ Preset |  |
| 0x0085E628 | `Dance` | EQ Preset |  |
| 0x0085E638 | `Electronic` | EQ Preset |  |
| 0x0085E64C | `Hip Hop` | EQ Preset |  |
| 0x0085E654 | `Jazz` | EQ Preset |  |
| 0x0085E65C | `Latin` | EQ Preset |  |
| 0x0085E664 | `Loudness` | EQ Preset |  |
| 0x0085E670 | `Lounge` | EQ Preset |  |
| 0x0085E678 | `Piano` | EQ Preset |  |
| 0x0085E688 | `Rock` | EQ Preset |  |
| 0x0085E690 | `Small Speakers` | EQ Preset |  |
| 0x0085E6A0 | `Spoken Word` | EQ Preset |  |
| 0x0085E6AC | `Treble Booster` | EQ Preset |  |
| 0x0085E6CC | `Vocal Booster` | EQ Preset |  |
| 0x00867B64 | `Acoustic` | EQ Preset |  |
| 0x00867B94 | `Dance` | EQ Preset |  |
| 0x00867BA4 | `Electronic` | EQ Preset |  |
| 0x00867BC0 | `Jazz` | EQ Preset |  |
| 0x00867BC8 | `Latin` | EQ Preset |  |
| 0x00867BD0 | `Loudness` | EQ Preset |  |
| 0x00867BE4 | `Piano` | EQ Preset |  |
| 0x00867BF4 | `Rock` | EQ Preset |  |
| 0x0087F154 | `Dance` | EQ Preset |  |
| 0x0087F17C | `Hip Hop` | EQ Preset |  |
| 0x0087F184 | `Jazz` | EQ Preset |  |
| 0x0087F194 | `Loudness` | EQ Preset |  |
| 0x0087F1A0 | `Lounge` | EQ Preset |  |
| 0x0087F1A8 | `Piano` | EQ Preset |  |
| 0x0087F1B8 | `Rock` | EQ Preset |  |
| 0x00888268 | `Jazz` | EQ Preset |  |
| 0x00888270 | `Latin` | EQ Preset |  |
| 0x00888284 | `Lounge` | EQ Preset |  |
| 0x0088828C | `Piano` | EQ Preset |  |
| 0x0088829C | `Rock` | EQ Preset |  |
| 0x008912A4 | `Hip Hop` | EQ Preset |  |
| 0x008912AC | `Jazz` | EQ Preset |  |
| 0x008912C8 | `Lounge` | EQ Preset |  |
| 0x008912D0 | `Piano` | EQ Preset |  |
| 0x008912E8 | `Rock` | EQ Preset |  |
| 0x0089AEDC | `Latin` | EQ Preset |  |
| 0x0089AF08 | `Rock` | EQ Preset |  |
| 0x008A4478 | `Dance` | EQ Preset |  |
| 0x008A449C | `Hip Hop` | EQ Preset |  |
| 0x008A44A4 | `Jazz` | EQ Preset |  |
| 0x008A44B4 | `Loudness` | EQ Preset |  |
| 0x008A44C0 | `Lounge` | EQ Preset |  |
| 0x008A44C8 | `Piano` | EQ Preset |  |
| 0x008A44D8 | `Rock` | EQ Preset |  |
| 0x008AEDDC | `Acoustic` | EQ Preset |  |
| 0x008AEDE8 | `Bass Booster` | EQ Preset |  |
| 0x008AEE08 | `Classical` | EQ Preset |  |
| 0x008AEE14 | `Dance` | EQ Preset |  |
| 0x008AEE24 | `Electronic` | EQ Preset |  |
| 0x008AEE38 | `Hip Hop` | EQ Preset |  |
| 0x008AEE40 | `Jazz` | EQ Preset |  |
| 0x008AEE48 | `Latin` | EQ Preset |  |
| 0x008AEE50 | `Loudness` | EQ Preset |  |
| 0x008AEE5C | `Lounge` | EQ Preset |  |
| 0x008AEE64 | `Piano` | EQ Preset |  |
| 0x008AEE74 | `Rock` | EQ Preset |  |
| 0x008AEE7C | `Small Speakers` | EQ Preset |  |
| 0x008AEE8C | `Spoken Word` | EQ Preset |  |
| 0x008AEE98 | `Treble Booster` | EQ Preset |  |
| 0x008AEEB8 | `Vocal Booster` | EQ Preset |  |
| 0x008B95A0 | `Acoustic` | EQ Preset |  |
| 0x008B95AC | `Bass Booster` | EQ Preset |  |
| 0x008B95CC | `Classical` | EQ Preset |  |
| 0x008B95D8 | `Dance` | EQ Preset |  |
| 0x008B95E8 | `Electronic` | EQ Preset |  |
| 0x008B95FC | `Hip Hop` | EQ Preset |  |
| 0x008B9604 | `Jazz` | EQ Preset |  |
| 0x008B960C | `Latin` | EQ Preset |  |
| 0x008B9614 | `Loudness` | EQ Preset |  |
| 0x008B9620 | `Lounge` | EQ Preset |  |
| 0x008B9628 | `Piano` | EQ Preset |  |
| 0x008B9638 | `Rock` | EQ Preset |  |
| 0x008B9640 | `Small Speakers` | EQ Preset |  |
| 0x008B9650 | `Spoken Word` | EQ Preset |  |
| 0x008B965C | `Treble Booster` | EQ Preset |  |
| 0x008B967C | `Vocal Booster` | EQ Preset |  |
| 0x008C2D64 | `Dance` | EQ Preset |  |
| 0x008C2D98 | `Jazz` | EQ Preset |  |
| 0x008C2DA0 | `Latin` | EQ Preset |  |
| 0x008C2DA8 | `Loudness` | EQ Preset |  |
| 0x008C2DB4 | `Lounge` | EQ Preset |  |
| 0x008C2DBC | `Piano` | EQ Preset |  |
| 0x008C2DCC | `Rock` | EQ Preset |  |
| 0x008CBDE8 | `Dance` | EQ Preset |  |
| 0x008CBE14 | `Jazz` | EQ Preset |  |
| 0x008CBE24 | `Loudness` | EQ Preset |  |
| 0x008CBE30 | `Lounge` | EQ Preset |  |
| 0x008CBE38 | `Piano` | EQ Preset |  |
| 0x008CBE48 | `Rock` | EQ Preset |  |
| 0x008D5118 | `Hip Hop` | EQ Preset |  |
| 0x008D5120 | `Jazz` | EQ Preset |  |
| 0x008D5144 | `Lounge` | EQ Preset |  |
| 0x008D515C | `Rock` | EQ Preset |  |
| 0x008DE818 | `Hip Hop` | EQ Preset |  |
| 0x008DE820 | `Jazz` | EQ Preset |  |
| 0x008DE83C | `Lounge` | EQ Preset |  |
| 0x008DE844 | `Piano` | EQ Preset |  |
| 0x008DE854 | `Rock` | EQ Preset |  |
| 0x008F49AC | `Acoustic` | EQ Preset |  |
| 0x008F49B8 | `Bass Booster` | EQ Preset |  |
| 0x008F49D8 | `Classical` | EQ Preset |  |
| 0x008F49E4 | `Dance` | EQ Preset |  |
| 0x008F49F4 | `Electronic` | EQ Preset |  |
| 0x008F4A08 | `Hip Hop` | EQ Preset |  |
| 0x008F4A10 | `Jazz` | EQ Preset |  |
| 0x008F4A18 | `Latin` | EQ Preset |  |
| 0x008F4A20 | `Loudness` | EQ Preset |  |
| 0x008F4A2C | `Lounge` | EQ Preset |  |
| 0x008F4A34 | `Piano` | EQ Preset |  |
| 0x008F4A44 | `Rock` | EQ Preset |  |
| 0x008F4A4C | `Small Speakers` | EQ Preset |  |
| 0x008F4A5C | `Spoken Word` | EQ Preset |  |
| 0x008F4A68 | `Treble Booster` | EQ Preset |  |
| 0x008F4A88 | `Vocal Booster` | EQ Preset |  |
| 0x008FDCC0 | `Hip Hop` | EQ Preset |  |
| 0x008FDCCC | `Latin` | EQ Preset |  |
| 0x008FDCD4 | `Loudness` | EQ Preset |  |
| 0x008FDCE0 | `Lounge` | EQ Preset |  |
| 0x008FDCF8 | `Rock` | EQ Preset |  |
| 0x009070BC | `Acoustic` | EQ Preset |  |
| 0x009070C8 | `Bass Booster` | EQ Preset |  |
| 0x009070E8 | `Classical` | EQ Preset |  |
| 0x009070F4 | `Dance` | EQ Preset |  |
| 0x00907104 | `Electronic` | EQ Preset |  |
| 0x00907118 | `Hip Hop` | EQ Preset |  |
| 0x00907120 | `Jazz` | EQ Preset |  |
| 0x00907128 | `Latin` | EQ Preset |  |
| 0x00907130 | `Loudness` | EQ Preset |  |
| 0x0090713C | `Lounge` | EQ Preset |  |
| 0x00907144 | `Piano` | EQ Preset |  |
| 0x00907154 | `Rock` | EQ Preset |  |
| 0x0090715C | `Small Speakers` | EQ Preset |  |
| 0x0090716C | `Spoken Word` | EQ Preset |  |
| 0x00907178 | `Treble Booster` | EQ Preset |  |
| 0x00907198 | `Vocal Booster` | EQ Preset |  |
| 0x00910378 | `Acoustic` | EQ Preset |  |
| 0x00910384 | `Bass Booster` | EQ Preset |  |
| 0x009103A4 | `Classical` | EQ Preset |  |
| 0x009103B0 | `Dance` | EQ Preset |  |
| 0x009103C0 | `Electronic` | EQ Preset |  |
| 0x009103D4 | `Hip Hop` | EQ Preset |  |
| 0x009103DC | `Jazz` | EQ Preset |  |
| 0x009103E4 | `Latin` | EQ Preset |  |
| 0x009103EC | `Loudness` | EQ Preset |  |
| 0x009103F8 | `Lounge` | EQ Preset |  |
| 0x00910400 | `Piano` | EQ Preset |  |
| 0x00910410 | `Rock` | EQ Preset |  |
| 0x00910418 | `Small Speakers` | EQ Preset |  |
| 0x00910428 | `Spoken Word` | EQ Preset |  |
| 0x00910434 | `Treble Booster` | EQ Preset |  |
| 0x00910454 | `Vocal Booster` | EQ Preset |  |

---
