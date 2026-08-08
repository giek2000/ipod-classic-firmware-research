# iPod Nano 3rd Gen - RetailOS 1.0.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.0.3 |
| **IPSW** | iPod_26.1.0.3.ipsw |
| **Device** | iPod Nano 3rd Gen (2007, 4/8GB NAND, Click Wheel, Cover Flow, Video) |
| **UpdaterFamilyID** | 26 |
| **Binary Size** | 10,359,824 bytes (9.88 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,357,776 bytes |
| **Total Strings (>=4)** | 67,190 |
| **Function Prologues** | 21,992 (ARM: 17,112, Thumb: 4,880) |
| **DRAM References** | 84,148 |
| **Peripheral Refs** | 6,000 |
| **Build** | N46FirmwareWin-313 |
| **SoC** | S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Codename** | N46 |
| **DFU PID** | 0x1229 |
| **SHA-256** | `e28cb0fc16afbe90bcbb4f03bb55a3991ae21db6729be207ae94c06861931318` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000A1F8C | `TSilverCntlr` | Known | Controller |
| 0x000A1FA4 | `TCExtrasMenu` | Known | Controller |
| 0x000A1FBC | `TCGameScreen` | Known | Controller |
| 0x000A1FD4 | `TCGamesMenu` | Known | Controller |
| 0x000A1FE8 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x000A2010 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x000A2038 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x000A2064 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x000A2088 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x000A20B0 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x000A20D8 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x000A2100 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x000A2128 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x000A2150 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x000A2180 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x000A21AC | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x000A21DC | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x000A2204 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x000A222C | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x000A2258 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x000A2280 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x000A22A8 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x000A22D8 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x000A2308 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x000A2384 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x000A23B4 | `TSilverGlobalCntlr` | Known | Controller |
| 0x000A23D0 | `TSilverTrainerCntlr` | Known | Controller |
| 0x000FB098 | `TCSlideshowLCD` | Known | Controller |
| 0x000FB0B0 | `TCSlideshowTVOut` | Known | Controller |
| 0x000FB0CC | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x000FB0EC | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x001231D4 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x00123200 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x0012322C | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x00123254 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x00123280 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x001232A8 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x001232D4 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x0012A438 | `TCRemoteUI` | Known | Controller |
| 0x0012A44C | `TCUnsupported` | Known | Controller |
| 0x0012F904 | `TCSpeakers` | Known | Controller |
| 0x0012F918 | `TCEQSetting` | Known | Controller |
| 0x0015A914 | `TCSportTimer` | Known | Controller |
| 0x0015A92C | `TCSportTimerMenu` | Known | Controller |
| 0x0015A948 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x0015A96C | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x0015BCDC | `TCVoiceMemos` | Known | Controller |
| 0x0015BCF4 | `TCVoiceMemosMenu` | Known | Controller |
| 0x0015BD10 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x0015BD30 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x0015BD50 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x0016CDB4 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x0016CDDC | `TCSettings_MainMenu` | Known | Controller |
| 0x0016CDF8 | `TCSettings_MusicMenu` | Known | Controller |
| 0x0016CE18 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x0016CE38 | `TCSettings_Brightness` | Known | Controller |
| 0x0016CE58 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0016CE7C | `TCSettings_EQ` | Known | Controller |
| 0x0016CE94 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x0016CEBC | `TCSettings_RadioRegions` | Known | Controller |
| 0x0016CEDC | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0016CF00 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x0016CF24 | `TCDateTimeScreen` | Known | Controller |
| 0x0016CF40 | `TCTimeZoneScreen` | Known | Controller |
| 0x0016CF5C | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0016CF84 | `TCFirstBoot` | Known | Controller |
| 0x00181558 | `TCDemoMode` | Known | Controller |
| 0x001ABDB0 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x001ABDD0 | `TCAddressViewerDetails` | Known | Controller |
| 0x001ABDF0 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x001ABE14 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x001DA7A4 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x001DA7C8 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x0027ABD4 | `TC_LockDialog` | Known | Controller |
| 0x0027ABEC | `TC_LockScreen` | Known | Controller |
| 0x0027AC04 | `TC_LockediPod` | Known | Controller |
| 0x0027AC1C | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x0027AC40 | `TCLockChosenDispatcher` | Known | Controller |
| 0x002805D4 | `TCClock` | Known | Controller |
| 0x002805E4 | `TCClockCityMenu` | Known | Controller |
| 0x002805FC | `TCClockRegionMenu` | Known | Controller |
| 0x00280618 | `TCAlarmMenu` | Known | Controller |
| 0x0028062C | `TCSleepTimerMenu` | Known | Controller |
| 0x00280648 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x00280668 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x00280690 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x002806B4 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x002806D8 | `TCAlarmDatePicker` | Known | Controller |
| 0x002806F4 | `TCAlarmTriggered` | Known | Controller |
| 0x00287224 | `TCNotesDispatcher` | Known | Controller |
| 0x00287240 | `TCNotesLoading` | Known | Controller |
| 0x00287258 | `TCNotesList` | Known | Controller |
| 0x0028726C | `TCNotesContents` | Known | Controller |
| 0x003A8710 | `TCAlarmTriggered` | Known | Controller |
| 0x003A8724 | `TSilverCntlr` | Known | Controller |
| 0x003A8744 | `TCClock` | Known | Controller |
| 0x003A874C | `TCClockRegionMenu` | Known | Controller |
| 0x003A8760 | `TCClockCityMenu` | Known | Controller |
| 0x003A8770 | `TCAlarmMenu` | Known | Controller |
| 0x003A877C | `TCSleepTimerMenu` | Known | Controller |
| 0x003A8790 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003A87A8 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003A87C8 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003A87E4 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003A8800 | `TCAlarmDatePicker` | Known | Controller |
| 0x003A8838 | `TSilverCntlr` | Known | Controller |
| 0x003A8858 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003A89E8 | `TSilverCntlr` | Known | Controller |
| 0x003A8A08 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x003A8A28 | `TCSettings_Brightness` | Known | Controller |
| 0x003A8A40 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x003A8A5C | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x003A8A7C | `TCSettings_RadioRegions` | Known | Controller |
| 0x003A8A94 | `TCSettings_EQ` | Known | Controller |
| 0x003A8AA4 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x003A8AC0 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x003A8AE0 | `TCFirstBoot` | Known | Controller |
| 0x003A8AEC | `TCSettings_MainMenu` | Known | Controller |
| 0x003A8B00 | `TCSettings_MusicMenu` | Known | Controller |
| 0x003A8B18 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003A8B30 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x003A8B4C | `TCDateTimeScreen` | Known | Controller |
| 0x003A8B60 | `TCTimeZoneScreen` | Known | Controller |
| 0x003AFB58 | `TSilverCntlr` | Known | Controller |
| 0x003AFB78 | `TCClock` | Known | Controller |
| 0x003AFB80 | `TCClockRegionMenu` | Known | Controller |
| 0x003AFB94 | `TCClockCityMenu` | Known | Controller |
| 0x003AFBA4 | `TCAlarmMenu` | Known | Controller |
| 0x003AFBB0 | `TCSleepTimerMenu` | Known | Controller |
| 0x003AFBC4 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003AFC3C | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003AFC5C | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003AFC78 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003AFCAC | `TCAlarmDatePicker` | Known | Controller |
| 0x003AFCC0 | `TCAlarmTriggered` | Known | Controller |
| 0x003B173C | `TSilverCntlr` | Known | Controller |
| 0x003B175C | `TC_LockDialog` | Known | Controller |
| 0x003B176C | `TC_LockScreen` | Known | Controller |
| 0x003B177C | `TC_LockediPod` | Known | Controller |
| 0x003B178C | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003B17A8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003B17C0 | `TSilverCntlr` | Known | Controller |
| 0x003B191C | `TSilverCntlr` | Known | Controller |
| 0x003B193C | `TCRemoteUI` | Known | Controller |
| 0x003B1948 | `TCUnsupported` | Known | Controller |
| 0x003B1958 | `TSilverCntlr` | Known | Controller |
| 0x003B19BC | `TSilverCntlr` | Known | Controller |
| 0x003B19DC | `TCSportTimer` | Known | Controller |
| 0x003B19EC | `TCSportTimerMenu` | Known | Controller |
| 0x003B1A00 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x003B1A1C | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x003B1B80 | `TSilverCntlr` | Known | Controller |
| 0x003B1BA0 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003B1BBC | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003B1BDC | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003B1BFC | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003B1C24 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003B1C48 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003B1C70 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003B1C90 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003B1CB0 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003B1CD0 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003B1CF0 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003B1D18 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003B2074 | `TSilverCntlr` | Known | Controller |
| 0x003B219C | `TSilverCntlr` | Known | Controller |
| 0x003B21BC | `TCDemoMode` | Known | Controller |
| 0x003B21D4 | `TSilverCntlr` | Known | Controller |
| 0x003B21F0 | `TSilverCntlr` | Known | Controller |
| 0x003B2200 | `TSilverCntlr` | Known | Controller |
| 0x003B2220 | `TCVoiceMemos` | Known | Controller |
| 0x003B2230 | `TCVoiceMemosMenu` | Known | Controller |
| 0x003B2244 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x003B225C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x003B2274 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x003B2294 | `TSilverCntlr` | Known | Controller |
| 0x003B22F4 | `TSilverCntlr` | Known | Controller |
| 0x003B2360 | `TSilverCntlr` | Known | Controller |
| 0x003B3198 | `TSilverCntlr` | Known | Controller |
| 0x003B32A4 | `TSilverCntlr` | Known | Controller |
| 0x003BB7CC | `TSilverCntlr` | Known | Controller |
| 0x003BB7EC | `TCAddressViewerMainMenu` | Known | Controller |
| 0x003BB804 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x003BB820 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x003BB840 | `TCAddressViewerDetails` | Known | Controller |
| 0x003BB858 | `TSilverCntlr` | Known | Controller |
| 0x003BB878 | `TSilverCalendarCntlr_Base` | Known | Controller |
| 0x003BB894 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x003BB8B8 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x003BB8DC | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x003BB8FC | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x003BB920 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x003BB940 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x003BB964 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x003BBB3C | `TSilverCntlr` | Known | Controller |
| 0x003BBB5C | `TC_LockDialog` | Known | Controller |
| 0x003BBB6C | `TC_LockScreen` | Known | Controller |
| 0x003BBB7C | `TC_LockediPod` | Known | Controller |
| 0x003BBB8C | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003BBBB0 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003BBC64 | `TSilverCntlr` | Known | Controller |
| 0x003BBC84 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003BBCA0 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003BBCC0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003BBCE0 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003BBD08 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003BBD2C | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003BBD54 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003BBD74 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003BBD94 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003BBDB4 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003BBDD4 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003BBDFC | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003BBE30 | `TSilverCntlr` | Known | Controller |
| 0x003BBE40 | `TSilverCntlr` | Known | Controller |
| 0x003BBFB4 | `TSilverCntlr` | Known | Controller |
| 0x003BBFD4 | `TCNotesDispatcher` | Known | Controller |
| 0x003BBFE8 | `TCNotesLoading` | Known | Controller |
| 0x003BBFF8 | `TCNotesBase` | Known | Controller |
| 0x003BC004 | `TCNotesList` | Known | Controller |
| 0x003BC010 | `TCNotesContents` | Known | Controller |
| 0x003BC020 | `TSilverCntlr` | Known | Controller |
| 0x003BC0E4 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003BC100 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003BC120 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003BC140 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003BC168 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003BC18C | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003BC1B4 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003BC1D4 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003BC1F4 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003BC214 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003BC234 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003BC25C | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003BC2AC | `TCSlideshowTVOut` | Known | Controller |
| 0x003BC2C0 | `TCSlideshowLCD` | Known | Controller |
| 0x003BC2D0 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x003BC2E8 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x003BC308 | `TSilverCntlr` | Known | Controller |
| 0x003BC334 | `TSilverCntlr` | Known | Controller |
| 0x003BC354 | `TCUnsupported` | Known | Controller |
| 0x003BC374 | `TSilverCntlr` | Known | Controller |
| 0x003BC3B4 | `TSilverCntlr` | Known | Controller |
| 0x003BC3D4 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x003BC3F0 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x003BC408 | `TSilverCntlr` | Known | Controller |
| 0x003BC428 | `TCSpeakers` | Known | Controller |
| 0x003BC434 | `TCEQSetting` | Known | Controller |
| 0x003BC4DC | `TSilverCntlr` | Known | Controller |
| 0x003BC4EC | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003BC508 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003BC528 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003BC548 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003BC570 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003BC594 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003BC5BC | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003BC5DC | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003BC5FC | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003BC61C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003BC63C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003BC664 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003BCC0C | `TSilverCntlr` | Known | Controller |
| 0x003BCC30 | `TSilverCntlr` | Known | Controller |
| 0x003BCC9C | `TSilverCntlr` | Known | Controller |
| 0x003BCCBC | `TCExtrasMenu` | Known | Controller |
| 0x003BCCCC | `TCGamesMenu` | Known | Controller |
| 0x003BCCD8 | `TCGameScreen` | Known | Controller |
| 0x003BCCE8 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x003BCD08 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x003BCD28 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x003BCD48 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x003BCD6C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003BCD88 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003BCDA8 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003BCDC8 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003BCDF0 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003BCE14 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003BCE3C | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003BCE5C | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003BCE7C | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003BCE9C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003BCEBC | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003BCEE4 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003BCF0C | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003BCF2C | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003BCF4C | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003BCF70 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003BCF90 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003BCFB4 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003BCFDC | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003BD008 | `TSilverGlobalCntlr` | Known | Controller |
| 0x003BD01C | `TSilverTrainerCntlr` | Known | Controller |
| 0x00445BC8 | `TCCCCCCCCCCCCCCCCCCCCCCT` | Known | Controller |
| 0x0072EA66 | `TCNotesDispatcher"` | Known | Controller |
| 0x0072EB25 | `TCLockChosenDispatcher"` | Known | Controller |
| 0x0072EBE8 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x00738AEF | `TCNotesDispatcher"` | Known | Controller |
| 0x00738C51 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x0074F9E4 | `TCExtrasMenuTCAddressViewerMainMenu` | Known | Controller |
| 0x0074FA08 | `TCAddressViewerDetails` | Known | Controller |
| 0x0074FA20 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x0074FA3C | `TCAlarmMenu` | Known | Controller |
| 0x0074FA48 | `TCSleepTimerMenuTCAlarmPropertiesMenu` | Known | Controller |
| 0x0074FA70 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x0074FA90 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x0074FAAC | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0074FAC8 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0074FAE4 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0074FB00 | `TCAlarmDatePicker` | Known | Controller |
| 0x0074FB14 | `TCAlarmDatePicker` | Known | Controller |
| 0x0074FB28 | `TCAlarmTriggeredTSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x0074FB54 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x0074FB78 | `TSilverCalendarCntlr_MonthViewerTSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x0074FBB8 | `TSilverCalendarCntlr_EventViewerTSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x0074FBF8 | `TSilverCalendarCntlr_EventViewerTSilverCntlrTCClockRegionMenu` | Known | Controller |
| 0x0074FC38 | `TCClockCityMenu` | Known | Controller |
| 0x0074FC48 | `TCClockCityMenu` | Known | Controller |
| 0x0074FC58 | `TCClockCityMenu` | Known | Controller |
| 0x0074FC68 | `TCClockCityMenu` | Known | Controller |
| 0x0074FC78 | `TCClockCityMenu` | Known | Controller |
| 0x0074FC88 | `TCClockCityMenu` | Known | Controller |
| 0x0074FC98 | `TCClockCityMenu` | Known | Controller |
| 0x0074FCA8 | `TCClockCityMenu` | Known | Controller |
| 0x0074FCB8 | `TCClock` | Known | Controller |
| 0x0074FCD0 | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDe` | Known | Controller |
| 0x0074FD28 | `TCGamesMenu` | Known | Controller |
| 0x0074FD34 | `TCGameScreenTC_LockediPod` | Known | Controller |
| 0x0074FD50 | `TC_LockDialog` | Known | Controller |
| 0x0074FD60 | `TC_LockScreen` | Known | Controller |
| 0x0074FD70 | `TC_VolumeLimitLockScreenTSilverCntlrTSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x0074FDB4 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x0074FDD4 | `TSilverMediaListCntlr_AudiobooksTSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x0074FE1C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x0074FE38 | `TSilverMediaListCntlr_AlbumsTSilverMediaListCntlr_Artists` | Known | Controller |
| 0x0074FE74 | `TSilverMediaListCntlr_GenresTSilverMediaListCntlr_Composers` | Known | Controller |
| 0x0074FEB0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0074FED0 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0074FEF8 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x0074FF18 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x0074FF38 | `TSilverMediaListCntlr_TVEpisodesTSilverMediaListCntlr_MoviesTSilverMediaListCntl` | Known | Controller |
| 0x0074FF94 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0074FFBC | `TSilverMediaListCntlr_VideoPlaylistsTSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x0075000C | `TSilverCntlrTSilverMainMediaListCntlr_VideosTSilverSettingsVideoCntlr` | Known | Controller |
| 0x00750054 | `TCFirstBoot` | Known | Controller |
| 0x007500FC | `TCNotesLoading` | Known | Controller |
| 0x0075010C | `TCNotesList` | Known | Controller |
| 0x00750118 | `TCNotesList` | Known | Controller |
| 0x00750124 | `TCNotesContents` | Known | Controller |
| 0x00750134 | `TCNotesContents` | Known | Controller |
| 0x00750144 | `TCNotesContents` | Known | Controller |
| 0x00750154 | `TCNotesContents` | Known | Controller |
| 0x00750210 | `TCSlideshowLCD` | Known | Controller |
| 0x00750220 | `TCSlideshowTVOutTCSlideshow_TVOutAskTCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x00750270 | `TCRemoteUI` | Known | Controller |
| 0x0075027C | `TCUnsupported` | Known | Controller |
| 0x0075028C | `TSilverSettingsMenuListCntlrTSilverSettingsMenuListCntlrTCSettings_MainMenu` | Known | Controller |
| 0x007502D8 | `TCSettings_MusicMenuTCSettings_VolumeLimit` | Known | Controller |
| 0x00750304 | `TCSettings_Brightness` | Known | Controller |
| 0x0075031C | `TCSettings_BacklightTimer` | Known | Controller |
| 0x00750338 | `TCSettings_AudiobookSettingsTCSettings_RadioRegions` | Known | Controller |
| 0x0075036C | `TCSettings_EQ` | Known | Controller |
| 0x0075037C | `TSilverSettingsMenuListCntlrTSilverCntlrTCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x007503C4 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x007503E0 | `TCSettings_MainMenu` | Known | Controller |
| 0x007503F4 | `TCSettings_MusicMenuTCSportTimerTCSportTimerMenuTCSportTimerSessionScreen` | Known | Controller |
| 0x00750558 | `TSilverCntlrTTrainerEndSessionCntlr` | Known | Controller |
| 0x007505D0 | `TSilverCntlrTTrainerCalibrateWalkMenuCntlr` | Known | Controller |
| 0x00750820 | `TCVoiceMemosTCVoiceMemosMainMenuTCVoiceMemosMainMenuTCVoiceMemosMainMenuTSearchC` | Known | Controller |
| 0x00750880 | `TCEQSetting` | Known | Controller |
| 0x0075092E | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x00751C2D | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x0075774D | `TCLockChosenDispatcher` | Known | Controller |
| 0x007577AB | `TCNotesDispatcher` | Known | Controller |
| 0x00759221 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075927F | `TCNotesDispatcher` | Known | Controller |
| 0x0075ACF5 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075AD53 | `TCNotesDispatcher` | Known | Controller |
| 0x0075C7C9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075C827 | `TCNotesDispatcher` | Known | Controller |
| 0x0075E29D | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075E2FB | `TCNotesDispatcher` | Known | Controller |
| 0x0075FD71 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075FDCF | `TCNotesDispatcher` | Known | Controller |
| 0x00761845 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007618A3 | `TCNotesDispatcher` | Known | Controller |
| 0x00763319 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00763377 | `TCNotesDispatcher` | Known | Controller |
| 0x00764DED | `TCLockChosenDispatcher` | Known | Controller |
| 0x00764E4B | `TCNotesDispatcher` | Known | Controller |
| 0x007668C1 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076691F | `TCNotesDispatcher` | Known | Controller |
| 0x00768395 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007683F3 | `TCNotesDispatcher` | Known | Controller |
| 0x00769E69 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00769EC7 | `TCNotesDispatcher` | Known | Controller |
| 0x0076B93D | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076B99B | `TCNotesDispatcher` | Known | Controller |
| 0x0076D411 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076D46F | `TCNotesDispatcher` | Known | Controller |
| 0x0076EEE5 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076EF43 | `TCNotesDispatcher` | Known | Controller |
| 0x007709B9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00770A17 | `TCNotesDispatcher` | Known | Controller |
| 0x0077248D | `TCLockChosenDispatcher` | Known | Controller |
| 0x007724EB | `TCNotesDispatcher` | Known | Controller |
| 0x00773F61 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00773FBF | `TCNotesDispatcher` | Known | Controller |
| 0x00775A35 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00775A93 | `TCNotesDispatcher` | Known | Controller |
| 0x00777509 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00777567 | `TCNotesDispatcher` | Known | Controller |
| 0x00778FDD | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077903B | `TCNotesDispatcher` | Known | Controller |
| 0x0077AAB1 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077AB0F | `TCNotesDispatcher` | Known | Controller |
| 0x0077C585 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077C5E3 | `TCNotesDispatcher` | Known | Controller |
| 0x0077E059 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077E0B7 | `TCNotesDispatcher` | Known | Controller |
| 0x0077FB2D | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077FB8B | `TCNotesDispatcher` | Known | Controller |
| 0x00781601 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078165F | `TCNotesDispatcher` | Known | Controller |
| 0x007830D5 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00783133 | `TCNotesDispatcher` | Known | Controller |
| 0x00784BA9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00784C07 | `TCNotesDispatcher` | Known | Controller |
| 0x0078667D | `TCLockChosenDispatcher` | Known | Controller |
| 0x007866DB | `TCNotesDispatcher` | Known | Controller |
| 0x00788151 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007881AF | `TCNotesDispatcher` | Known | Controller |
| 0x00789C25 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00789C83 | `TCNotesDispatcher` | Known | Controller |
| 0x0078B6F9 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078B757 | `TCNotesDispatcher` | Known | Controller |
| 0x0078D1CD | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078D22B | `TCNotesDispatcher` | Known | Controller |
| 0x0078ECA1 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0078ECFF | `TCNotesDispatcher` | Known | Controller |
| 0x00790775 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007907D3 | `TCNotesDispatcher` | Known | Controller |
| 0x00792249 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007922A7 | `TCNotesDispatcher` | Known | Controller |
| 0x0079DDFC | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x0079DF9E | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x008D8280 | `TCMockupModeNavScreen` | Known | Controller |
| 0x008D8298 | `TSilverCntlr` | Known | Controller |
| 0x008D82B8 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x008D82F0 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x008D8310 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x008D8330 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x008D8354 | `TCExtrasMenu` | Known | Controller |
| 0x008D8D40 | `TSilverCntlr` | Known | Controller |
| 0x008D8D60 | `TCSlideshowTVOut` | Known | Controller |
| 0x008D8D74 | `TCSlideshowLCD` | Known | Controller |
| 0x008D8D84 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x008D8D9C | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x008D8DD8 | `TSilverCntlr` | Known | Controller |
| 0x008D8E54 | `TCSlideshowTVOut` | Known | Controller |
| 0x008D8E68 | `TCSlideshowLCD` | Known | Controller |
| 0x008D8E78 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x008D8E90 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x008D8EB0 | `TSilverCntlr` | Known | Controller |
| 0x008D9A4C | `TSilverCntlr` | Known | Controller |
| 0x008D9A6C | `TCGamesMenu` | Known | Controller |
| 0x008D9A78 | `TCGameScreen` | Known | Controller |
| 0x00995587 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x009CA13D | `TCL$]` | Known | Controller |

---

## 2. Settings (Toggle/Show/TC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001389F4 | `ShowSetting_EQ` | Known | User setting |
| 0x001E3B28 | `ToggleSetting_Repeat` | Known | User setting |
| 0x001E3B44 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001E3B5C | `ToggleSetting_TVOut` | Known | User setting |
| 0x001E3B70 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x00212978 | `ShowSetting_Backlight` | Known | User setting |
| 0x00226BB0 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x00226BCC | `ToggleSetting_Repeat` | Known | User setting |
| 0x00226BE4 | `ToggleSetting_SortBy` | Known | User setting |
| 0x00226BFC | `ToggleSetting_ClassicUI` | Known | User setting |
| 0x00226C14 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x00226C30 | `ToggleSetting_Clicker` | Known | User setting |
| 0x00226C48 | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x00226C68 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x00226C84 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x00226CA0 | `ShowSetting_Shuffle` | Known | User setting |
| 0x00226E38 | `ShowSetting_Repeat` | Known | User setting |
| 0x00226E4C | `ShowSetting_About` | Known | User setting |
| 0x00226E60 | `ShowSetting_MainMenu` | Known | User setting |
| 0x00226E78 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x00226E90 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x00226EA8 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x00226EC4 | `ShowSetting_Brightness` | Known | User setting |
| 0x00226EDC | `ShowSetting_Audiobooks` | Known | User setting |
| 0x00226EF4 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x00226F10 | `ShowSetting_EQ` | Known | User setting |
| 0x00226F20 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x002270BC | `ShowSetting_Clicker` | Known | User setting |
| 0x002270D0 | `ShowSetting_DateAndTime` | Known | User setting |
| 0x002270E8 | `ShowSetting_SortBy` | Known | User setting |
| 0x002270FC | `ShowSetting_ClassicUI` | Known | User setting |
| 0x00227114 | `ShowSetting_Language` | Known | User setting |
| 0x0022712C | `ShowSetting_Legal` | Known | User setting |
| 0x00227140 | `ShowSetting_ResetAll` | Known | User setting |
| 0x007379F5 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x00737AA5 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x0073A09A | `ShowSetting_About` | Known | User setting |
| 0x0073A1A2 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x0073A1E6 | `ShowSetting_Shuffle` | Known | User setting |
| 0x0073A25D | `ToggleSetting_Repeat` | Known | User setting |
| 0x0073A2A0 | `ShowSetting_Repeat` | Known | User setting |
| 0x0073A3AA | `ShowSetting_MainMenu` | Known | User setting |
| 0x0073A4BA | `ShowSetting_MusicMenu` | Known | User setting |
| 0x0073A582 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x0073A64C | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x0073A764 | `ShowSetting_Brightness` | Known | User setting |
| 0x0073A89A | `ShowSetting_Audiobooks` | Known | User setting |
| 0x0073A9AB | `ShowSetting_RadioRegions` | Known | User setting |
| 0x0073AAAC | `ShowSetting_EQ` | Known | User setting |
| 0x0073AB19 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x0073AB60 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x0073ABDD | `ToggleSetting_Clicker` | Known | User setting |
| 0x0073AC21 | `ShowSetting_Clicker` | Known | User setting |
| 0x0073AD88 | `ToggleSetting_SortBy` | Known | User setting |
| 0x0073ADCB | `ShowSetting_SortBy` | Known | User setting |
| 0x0073AECC | `ShowSetting_Language` | Known | User setting |
| 0x0073AFDC | `ShowSetting_Legal` | Known | User setting |
| 0x0073B10D | `ShowSetting_ResetAll` | Known | User setting |
| 0x0073B280 | `ShowSetting_Backlight` | Known | User setting |
| 0x0073B330 | `ShowSetting_Backlight` | Known | User setting |
| 0x0073B3E0 | `ShowSetting_Backlight` | Known | User setting |
| 0x0073B491 | `ShowSetting_Backlight` | Known | User setting |
| 0x0073B542 | `ShowSetting_Backlight` | Known | User setting |
| 0x0073B5F3 | `ShowSetting_Backlight` | Known | User setting |
| 0x0073B6A7 | `ShowSetting_Backlight` | Known | User setting |
| 0x0073B756 | `ShowSetting_EQ` | Known | User setting |
| 0x0073B7CB | `ShowSetting_Language` | Known | User setting |
| 0x007ACE88 | `ToggleSetting_Repeat` | Known | User setting |
| 0x007ACEC2 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x007ACF84 | `ToggleSetting_TVOut` | Known | User setting |
| 0x007ACFBD | `ToggleSetting_TVSignal` | Known | User setting |

---

## 3. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0015678C | `MockupMode/MockupMode.xml` | Hidden | Developer Tool |
| 0x00156C8C | `MockupMode/` | Hidden | Developer Tool |
| 0x00260620 | `Channel UnitTests` | Hidden | Developer Tool |
| 0x002B4169 | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x002B41AC | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x002B41C1 | `RTXCbug> ` | Hidden | Developer Tool |
| 0x002B4B9D | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x002C632C | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x003525E9 | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x003526B1 | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x003ADB11 | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x007DFA5C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00823620 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0083488C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0084A5C4 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0085B2D4 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00864478 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0086D2F8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00880B24 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00889B78 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008AD63C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008C9694 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008D2028 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00986A69 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00987632 | `DemoMode_Background_Image` | Hidden | Demo/Retail Mode |
| 0x00988645 | `DemoMode_Icon_Image` | Hidden | Demo/Retail Mode |
| 0x0098A457 | `UnitTest_Menu_Title` | Hidden | Developer Tool |
| 0x00992D11 | `UnitTestModel` | Hidden | Developer Tool |
| 0x00994D29 | `DemoMode_View_Main` | Hidden | Demo/Retail Mode |
| 0x00994F11 | `DemoMode_Icon` | Hidden | Demo/Retail Mode |
| 0x00996A34 | `DemoMode_Text_Color` | Hidden | Demo/Retail Mode |

---

## 4. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000671B | `"MeCCADecode` | Known | Audio system |
| 0x0014D0E4 | `AudioCodecs` | Known | Audio system |
| 0x0018F23C | `MeCCA_PCM_Output.wav` | Known | Audio system |
| 0x001AAFDC | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x001B5890 | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x001B5A98 | `MeCCAVideoDecode` | Known | Audio system |
| 0x008E6DA4 | `MeCCA_StreamCache` | Known | Audio system |

---

## 5. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F6E4C | `HandleRemoteUIRemotePlayPause` | Known | Event handler |
| 0x000F6E70 | `HandleWheel` | Known | Event handler |
| 0x000F6E7C | `HandlePlayPause` | Known | Event handler |
| 0x000F6E8C | `HandleSelectDown` | Known | Event handler |
| 0x000F6EA0 | `HandleNext` | Known | Event handler |
| 0x000F6EAC | `HandlePrevious` | Known | Event handler |
| 0x000F6EBC | `HandleNextPushAndHold` | Known | Event handler |
| 0x000F6ED4 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000F71C0 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000F71E0 | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x00102E38 | `HandleSelect` | Known | Event handler |
| 0x00102E4C | `HandleHilite` | Known | Event handler |
| 0x00103170 | `HandleEQSettingSelected` | Known | Event handler |
| 0x001035A0 | `HandleSelect` | Known | Event handler |
| 0x001035B4 | `HandleGameHilited` | Known | Event handler |
| 0x00103864 | `HandleNotesSelected` | Known | Event handler |
| 0x0010387C | `HandleNotesPop` | Known | Event handler |
| 0x0010388C | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x001117C8 | `HandleVolumeWheel` | Known | Event handler |
| 0x001117DC | `HandleVolumeChange` | Known | Event handler |
| 0x001117F0 | `HandleTimerDone` | Known | Event handler |
| 0x00111800 | `HandleFrequencyChange` | Known | Event handler |
| 0x00111848 | `HandleTuning` | Known | Event handler |
| 0x00121578 | `HandleLock` | Known | Event handler |
| 0x00121588 | `HandleAddressBook` | Known | Event handler |
| 0x00121D88 | `HandleExit` | Known | Event handler |
| 0x00121D98 | `HandleLap` | Known | Event handler |
| 0x00121DA4 | `HandleResume` | Known | Event handler |
| 0x00121DB4 | `HandleStartStop` | Known | Event handler |
| 0x0012203C | `HandleWheel` | Known | Event handler |
| 0x0012204C | `HandlePlayPause` | Known | Event handler |
| 0x0012205C | `HandleSelectDown` | Known | Event handler |
| 0x00122070 | `HandleHilite` | Known | Event handler |
| 0x0012ACFC | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x00138C28 | `HandleExitUnsupported` | Known | Event handler |
| 0x00143E08 | `HandleBasicSelected` | Known | Event handler |
| 0x00143E20 | `HandleTimedWorkoutSelected` | Known | Event handler |
| 0x00143E3C | `HandleDistanceWorkoutSelected` | Known | Event handler |
| 0x00143E5C | `HandleCaloriesWorkoutSelected` | Known | Event handler |
| 0x00143E7C | `HandleSelectWorkout` | Known | Event handler |
| 0x001525D0 | `HandleNotesPop` | Known | Event handler |
| 0x001525E4 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x001534C8 | `HandleSelect` | Known | Event handler |
| 0x001534DC | `HandleWheel` | Known | Event handler |
| 0x001534E8 | `HandleImageNext` | Known | Event handler |
| 0x001534F8 | `HandleImagePrev` | Known | Event handler |
| 0x00153508 | `HandleImageLast` | Known | Event handler |
| 0x00153518 | `HandleImageFirst` | Known | Event handler |
| 0x0015352C | `HandlePlayPause` | Known | Event handler |
| 0x0015353C | `HandlePlay` | Known | Event handler |
| 0x00153548 | `HandlePause` | Known | Event handler |
| 0x00167A44 | `HandleSelectCity` | Known | Event handler |
| 0x00167A5C | `HandleHighlightCity` | Known | Event handler |
| 0x00168984 | `HandleWantPopFlow` | Known | Event handler |
| 0x0016899C | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x001689B8 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x001689D4 | `HandleFlowNext` | Known | Event handler |
| 0x001689E4 | `HandleFlowPrev` | Known | Event handler |
| 0x001689F4 | `HandleFlowWheel` | Known | Event handler |
| 0x00168A04 | `HandleAlbumSelected` | Known | Event handler |
| 0x00168A18 | `HandlePlayPause` | Known | Event handler |
| 0x00168A28 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x001910D8 | `HandleLeaveAlarm` | Known | Event handler |
| 0x001914C0 | `HandleSelect` | Known | Event handler |
| 0x00192380 | `HandleSelect` | Known | Event handler |
| 0x00192394 | `HandleWheel` | Known | Event handler |
| 0x001923A0 | `HandleImageNext` | Known | Event handler |
| 0x001923B0 | `HandleImagePrev` | Known | Event handler |
| 0x001923C0 | `HandleImageLast` | Known | Event handler |
| 0x001923D0 | `HandleImageFirst` | Known | Event handler |
| 0x001923E4 | `HandlePlayPause` | Known | Event handler |
| 0x001923F4 | `HandlePlay` | Known | Event handler |
| 0x00192400 | `HandlePause` | Known | Event handler |
| 0x001928A0 | `HandleNew` | Known | Event handler |
| 0x001928B0 | `HandleClear` | Known | Event handler |
| 0x001928BC | `HandleSelectCurrentSession` | Known | Event handler |
| 0x001928D8 | `HandleSelectIndexedSession` | Known | Event handler |
| 0x00192BE8 | `HandleWheel` | Known | Event handler |
| 0x00192BF8 | `HandleArrowUp` | Known | Event handler |
| 0x00192C08 | `HandleArrowDown` | Known | Event handler |
| 0x00194E2C | `HandleHiliteAlbum` | Known | Event handler |
| 0x00194E44 | `HandleBrowseAlbum` | Known | Event handler |
| 0x00194E58 | `HandlePlayPause` | Known | Event handler |
| 0x001AF680 | `HandleSelect` | Known | Event handler |
| 0x001AF810 | `HandleSelectRegion` | Known | Event handler |
| 0x001B4234 | `HandleChooseLink` | Known | Event handler |
| 0x001B424C | `HandleChooseCalibrate` | Known | Event handler |
| 0x001B4264 | `HandleUnlink` | Known | Event handler |
| 0x001C4178 | `HandleImageWheel` | Known | Event handler |
| 0x001C4190 | `HandlePlayPause` | Known | Event handler |
| 0x001C41A0 | `HandleBrowseLarge` | Known | Event handler |
| 0x001C41B4 | `HandleBrowseSmall` | Known | Event handler |
| 0x001C41C8 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x001C41E0 | `HandleImageNext` | Known | Event handler |
| 0x001C41F0 | `HandleImagePrev` | Known | Event handler |
| 0x001C4200 | `HandleHilite` | Known | Event handler |
| 0x001C4210 | `HandleImageLast` | Known | Event handler |
| 0x001C4220 | `HandleImageFirst` | Known | Event handler |
| 0x001C4234 | `HandleScreenNext` | Known | Event handler |
| 0x001C4248 | `HandleScreenPrev` | Known | Event handler |
| 0x001C6AE4 | `HandlePlayPause` | Known | Event handler |
| 0x001C6AF8 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x001C6B14 | `HandleNext` | Known | Event handler |
| 0x001C6B20 | `HandleNextPressAndHold` | Known | Event handler |
| 0x001C6B38 | `HandlePrevious` | Known | Event handler |
| 0x001C6B48 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x001C6B64 | `HandleRemotePlayPause` | Known | Event handler |
| 0x001C6B7C | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x001C6BA0 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x001C6BB8 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x001C6BD0 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x001C6DA0 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x001C6DB8 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x001C6DD0 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x001C6DEC | `HandleRemoteStop` | Known | Event handler |
| 0x001C6E00 | `HandleRemotePlay` | Known | Event handler |
| 0x001C6E14 | `HandleRemotePause` | Known | Event handler |
| 0x001C6E28 | `HandleRemoteMute` | Known | Event handler |
| 0x001C6E3C | `HandleRemoteNextChapter` | Known | Event handler |
| 0x001C6E54 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x001C6E6C | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x001C6E88 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x001C70AC | `HandleRemoteShuffle` | Known | Event handler |
| 0x001C70C0 | `HandleRemoteRepeat` | Known | Event handler |
| 0x001C70D4 | `HandleRemoteOn` | Known | Event handler |
| 0x001C70E4 | `HandleRemoteOff` | Known | Event handler |
| 0x001C70F4 | `HandleRemoteBacklight` | Known | Event handler |
| 0x001C710C | `HandleRemoteFFDown` | Known | Event handler |
| 0x001C7120 | `HandleRemoteFFUp` | Known | Event handler |
| 0x001C7134 | `HandleRemoteRewDown` | Known | Event handler |
| 0x001C7148 | `HandleRemoteRewUp` | Known | Event handler |
| 0x001C715C | `HandleRemoteMenuDown` | Known | Event handler |
| 0x001C7174 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x001C7188 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x001C71A0 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x001C7370 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x001C7388 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x001C73A0 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x001C73BC | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x001C73D4 | `HandleRemoteEvent` | Known | Event handler |
| 0x001C73E8 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x001C7404 | `HandleAudioPlayPause` | Known | Event handler |
| 0x001C741C | `HandleAudioNext` | Known | Event handler |
| 0x001C742C | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x001C7448 | `HandleAudioPrevious` | Known | Event handler |
| 0x001C745C | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x001C765C | `HandleAudioNextAlbum` | Known | Event handler |
| 0x001C7674 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x001C768C | `HandleAudioVolumeDown` | Known | Event handler |
| 0x001C76A4 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x001C76B8 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x001C76D0 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x001C76E8 | `HandleAudioStop` | Known | Event handler |
| 0x001C76F8 | `HandleAudioPlay` | Known | Event handler |
| 0x001C7708 | `HandleAudioPause` | Known | Event handler |
| 0x001C771C | `HandleAudioMute` | Known | Event handler |
| 0x001C772C | `HandleAudioNextChapter` | Known | Event handler |
| 0x001C7744 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x001C7964 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x001C797C | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x001C7994 | `HandleAudioShuffle` | Known | Event handler |
| 0x001C79A8 | `HandleAudioRepeat` | Known | Event handler |
| 0x001C79BC | `HandleAudioFFDown` | Known | Event handler |
| 0x001C79D0 | `HandleAudioFFUp` | Known | Event handler |
| 0x001C79E0 | `HandleAudioRewDown` | Known | Event handler |
| 0x001C79F4 | `HandleAudioRewUp` | Known | Event handler |
| 0x001C7A08 | `HandleVideoPlayPause` | Known | Event handler |
| 0x001C7A20 | `HandleVideoNext` | Known | Event handler |
| 0x001C7A30 | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x001C7A4C | `HandleVideoPrevious` | Known | Event handler |
| 0x001C7A60 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x001C7C68 | `HandleVideoStop` | Known | Event handler |
| 0x001C7C78 | `HandleVideoPlay` | Known | Event handler |
| 0x001C7C88 | `HandleVideoPause` | Known | Event handler |
| 0x001C7C9C | `HandleVideoFFDown` | Known | Event handler |
| 0x001C7CB0 | `HandleVideoFFUp` | Known | Event handler |
| 0x001C7CC0 | `HandleVideoRewDown` | Known | Event handler |
| 0x001C7CD4 | `HandleVideoRewUp` | Known | Event handler |
| 0x001C7CE8 | `HandleVideoNextChapter` | Known | Event handler |
| 0x001C7D00 | `HandleVideoPrevChapter` | Known | Event handler |
| 0x001C7D18 | `HandleVideoNextFrame` | Known | Event handler |
| 0x001C7D30 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x001C7D48 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x001CB254 | `HandleSelect` | Known | Event handler |
| 0x001CB268 | `HandleMenu` | Known | Event handler |
| 0x001CB274 | `HandleLinkCancelOption` | Known | Event handler |
| 0x001CB28C | `HandleLinkNewRemote` | Known | Event handler |
| 0x001CB5E4 | `HandleChoosePlaylistsPlay` | Known | Event handler |
| 0x001CB604 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x001CB620 | `HandleNoneSelected` | Known | Event handler |
| 0x001CB634 | `HandleNowPlayingSelected` | Known | Event handler |
| 0x001CB650 | `HandleMenuKeyNop` | Known | Event handler |
| 0x001CB664 | `HandlePlaylistSelected` | Known | Event handler |
| 0x001CBE30 | `HandlePauseWorkout` | Known | Event handler |
| 0x001CBE48 | `HandleEndWorkout` | Known | Event handler |
| 0x001CBE5C | `HandleResumeWorkout` | Known | Event handler |
| 0x001CBE70 | `HandleChooseMusic` | Known | Event handler |
| 0x001CBE84 | `HandleMenuKeyNop` | Known | Event handler |
| 0x001D73CC | `HandleMainMenu` | Known | Event handler |
| 0x001DB780 | `HandlePowerSongSelected` | Known | Event handler |
| 0x001DB79C | `HandlePowerSongChosen` | Known | Event handler |
| 0x001DB7B4 | `HandleSelectTimedWorkout` | Known | Event handler |
| 0x001DC024 | `HandleChooseRemoteSetting` | Known | Event handler |
| 0x001DC044 | `HandleChooseRemoteLink` | Known | Event handler |
| 0x001DC05C | `HandleChooseRemoteUnlink` | Known | Event handler |
| 0x001DC398 | `HandleSelectResume` | Known | Event handler |
| 0x001DC3B0 | `HandleEndWorkout` | Known | Event handler |
| 0x001E2534 | `HandleMusicMenu` | Known | Event handler |
| 0x001E27F4 | `HandleSelect` | Known | Event handler |
| 0x001E2B78 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x001E2B90 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x001E2BB0 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x001E2BD4 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x001E2BF0 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x001E308C | `HandleWheel` | Known | Event handler |
| 0x001E309C | `HandlePlayPause` | Known | Event handler |
| 0x001E30AC | `HandleSelectDown` | Known | Event handler |
| 0x001E30C0 | `HandleNext` | Known | Event handler |
| 0x001E30CC | `HandlePrevious` | Known | Event handler |
| 0x001E30DC | `HandleNextPushAndHold` | Known | Event handler |
| 0x001E30F4 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x001E8644 | `HandleChooseLast` | Known | Event handler |
| 0x001E865C | `HandleChooseRecent` | Known | Event handler |
| 0x001E8670 | `HandleChooseWorkout` | Known | Event handler |
| 0x001E8684 | `HandleChooseBest` | Known | Event handler |
| 0x001E8698 | `HandleChooseRecentWorkout` | Known | Event handler |
| 0x001EAD20 | `HandleSelect` | Known | Event handler |
| 0x001EAD34 | `HandleMenu` | Known | Event handler |
| 0x001F2AC8 | `HandleFrequencyChosen` | Known | Event handler |
| 0x001F2AE0 | `HandleDateChosen` | Known | Event handler |
| 0x001F2AF4 | `HandleTimeChosen` | Known | Event handler |
| 0x001F2B08 | `HandleSoundChosen` | Known | Event handler |
| 0x001F2B1C | `HandleLabelChosen` | Known | Event handler |
| 0x001F2B30 | `HandleDeleteChosen` | Known | Event handler |
| 0x001F8014 | `HandlePrev` | Known | Event handler |
| 0x001F8024 | `HandleNext` | Known | Event handler |
| 0x001F8030 | `HandlePlayPause` | Known | Event handler |
| 0x001F880C | `HandleChoosePowerPlay` | Known | Event handler |
| 0x001F8828 | `HandleChooseUnit` | Known | Event handler |
| 0x001F883C | `HandleChooseVoicePrompts` | Known | Event handler |
| 0x00200DF4 | `HandleNextContact` | Known | Event handler |
| 0x00200E0C | `HandlePreviousContact` | Known | Event handler |
| 0x002041D0 | `HandleSelect` | Known | Event handler |
| 0x002044AC | `HandleListChoose` | Known | Event handler |
| 0x00208EEC | `HandleItemSelected` | Known | Event handler |
| 0x002090E4 | `HandleRadioRegion` | Known | Event handler |
| 0x002092CC | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x00209B18 | `HandleSelect` | Known | Event handler |
| 0x00209F60 | `HandlePauseKey` | Known | Event handler |
| 0x00209F74 | `HandlePauseHold` | Known | Event handler |
| 0x00209F84 | `HandlePauseKeyNop` | Known | Event handler |
| 0x00209F98 | `HandleMenuKey` | Known | Event handler |
| 0x00209FA8 | `HandleMenuKeyNop` | Known | Event handler |
| 0x00209FBC | `HandleWheel` | Known | Event handler |
| 0x0020A00C | `HandleSelectKeyDown` | Known | Event handler |
| 0x0020A020 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x0020A038 | `HandlePowerPlay` | Known | Event handler |
| 0x0020EC30 | `HandlePlayPause` | Known | Event handler |
| 0x0020FE9C | `HandleSelect` | Known | Event handler |
| 0x0021012C | `HandleChooseCustomTimedWorkout` | Known | Event handler |
| 0x00210150 | `HandleChooseTimedWorkoutFromList` | Known | Event handler |
| 0x00210174 | `HandleChooseCustomDistanceWorkout` | Known | Event handler |
| 0x00210198 | `HandleChooseDistanceWorkoutFromList` | Known | Event handler |
| 0x002101BC | `HandleChooseCustomCalorieWorkout` | Known | Event handler |
| 0x002101E0 | `HandleChooseCalorieWorkoutFromList` | Known | Event handler |
| 0x00212C54 | `HandleDelete` | Known | Event handler |
| 0x00212C68 | `HandleSelectLozinch` | Known | Event handler |
| 0x00212F10 | `HandleSelect` | Known | Event handler |
| 0x0021318C | `HandleTVOutChanged` | Known | Event handler |
| 0x002131A4 | `HandleTVSignalChanged` | Known | Event handler |
| 0x002131BC | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x002131DC | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x002131FC | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x00213AE8 | `HandleBegin` | Known | Event handler |
| 0x00216C7C | `HandleSelectKey` | Known | Event handler |
| 0x00216E24 | `HandleSelect` | Known | Event handler |
| 0x00217BA4 | `HandlePlayPause` | Known | Event handler |
| 0x00217BB8 | `HandleWheel` | Known | Event handler |
| 0x00217BC4 | `HandleWheelRating` | Known | Event handler |
| 0x00217BD8 | `HandleWheelScrub` | Known | Event handler |
| 0x00217BEC | `HandleWheelVolume` | Known | Event handler |
| 0x00218774 | `HandleSelect` | Known | Event handler |
| 0x00218F08 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00219D3C | `HandleSelect` | Known | Event handler |
| 0x00219D50 | `HandleHilite` | Known | Event handler |
| 0x00219D60 | `HandlePlayPause` | Known | Event handler |
| 0x00219D70 | `HandleAddToOTG` | Known | Event handler |
| 0x00219D80 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0021AD8C | `HandleWeightWheel` | Known | Event handler |
| 0x0021ADA4 | `HandleWeightSelect` | Known | Event handler |
| 0x0021ADB8 | `HandleDistanceWheel` | Known | Event handler |
| 0x0021ADCC | `HandleDistanceSelect` | Known | Event handler |
| 0x0021ADE4 | `HandleTimeWheel` | Known | Event handler |
| 0x0021ADF4 | `HandleTimeSelect` | Known | Event handler |
| 0x0021AE08 | `HandleCaloriesWheel` | Known | Event handler |
| 0x0021AE1C | `HandleCaloriesSelect` | Known | Event handler |
| 0x0021B3E8 | `HandleSelect` | Known | Event handler |
| 0x0021B3FC | `HandleSelectFromLastWorkout` | Known | Event handler |
| 0x0021DCB8 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x0021E4C4 | `HandleSelect` | Known | Event handler |
| 0x0021E4D8 | `HandleWheel` | Known | Event handler |
| 0x0021E4E4 | `HandleWheelProgress` | Known | Event handler |
| 0x0021E4F8 | `HandleSelectProgress` | Known | Event handler |
| 0x0021E510 | `HandleSelectVolume` | Known | Event handler |
| 0x0021E524 | `HandleSelectScrub` | Known | Event handler |
| 0x0021E538 | `HandleSelectRating` | Known | Event handler |
| 0x0021E54C | `HandleSelectExtraInfo` | Known | Event handler |
| 0x0021E564 | `HandleSelectChapterArt` | Known | Event handler |
| 0x0021E57C | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x0021E598 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0021E5B4 | `HandleWheelBrightness` | Known | Event handler |
| 0x0021E6FC | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x0022003C | `HandleSelect` | Known | Event handler |
| 0x0022004C | `HandleSelectRating` | Known | Event handler |
| 0x00220060 | `HandleSelectProgress` | Known | Event handler |
| 0x00220078 | `HandleWheelProgress` | Known | Event handler |
| 0x0022008C | `HandleSelectScrub` | Known | Event handler |
| 0x002200A0 | `HandleWheelBrightness` | Known | Event handler |
| 0x002200B8 | `HandleNextTrackGotoScrub` | Known | Event handler |
| 0x002200D4 | `HandlePrevTrackGotoScrub` | Known | Event handler |
| 0x002200F0 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x00223420 | `HandleSelectWalking` | Known | Event handler |
| 0x00223438 | `HandleSelectRunning` | Known | Event handler |
| 0x00227178 | `HandleLanguage` | Known | Event handler |
| 0x00227188 | `HandleResetAllSettings` | Known | Event handler |
| 0x002271A0 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x002274E0 | `HandleUnlinkRemote` | Known | Event handler |
| 0x00227FD0 | `HandleSelect` | Known | Event handler |
| 0x00228200 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x00229278 | `Handle400MetersRun` | Known | Event handler |
| 0x00229290 | `HandleCustomRun` | Known | Event handler |
| 0x002292A0 | `HandleResetToDefault` | Known | Event handler |
| 0x00229700 | `HandleSelect_Basic` | Known | Event handler |
| 0x00229718 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x0022B7D8 | `HandleSelect` | Known | Event handler |
| 0x0022B974 | `HandleSelect` | Known | Event handler |
| 0x0022BC14 | `HandleNextDay` | Known | Event handler |
| 0x0022BC28 | `HandlePreviousDay` | Known | Event handler |
| 0x0022C42C | `HandleMusicHilited` | Known | Event handler |
| 0x0022C444 | `HandleVideosHilited` | Known | Event handler |
| 0x0022C458 | `HandlePodcastsHilited` | Known | Event handler |
| 0x0022C470 | `HandleGenericHilited` | Known | Event handler |
| 0x0022C488 | `HandlePhotosHilited` | Known | Event handler |
| 0x0022C49C | `HandleNowPlayingHilited` | Known | Event handler |
| 0x0022C4B4 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x0022C4D0 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x0022C4E8 | `HandleArtistsHilited` | Known | Event handler |
| 0x0022C500 | `HandleGenresHilited` | Known | Event handler |
| 0x0022C514 | `HandleAlbumsHilited` | Known | Event handler |
| 0x0022C528 | `HandleCompilationsHilited` | Known | Event handler |
| 0x0022C6FC | `HandleComposersHilited` | Known | Event handler |
| 0x0022C714 | `HandleSongsHilited` | Known | Event handler |
| 0x0022C728 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x0022C740 | `HandleTVShowsHilited` | Known | Event handler |
| 0x0022C758 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x0022C774 | `HandleMoviesHilited` | Known | Event handler |
| 0x0022C788 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x0022C7A4 | `HandleMusicSelected` | Known | Event handler |
| 0x0022C7B8 | `HandleVideosSelected` | Known | Event handler |
| 0x0022C7D0 | `HandlePodcastsSelected` | Known | Event handler |
| 0x0022C7E8 | `HandlePhotosSelected` | Known | Event handler |
| 0x0022C9B8 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x0022C9D0 | `HandleSongsSelected` | Known | Event handler |
| 0x0022C9E4 | `HandleAlbumsSelected` | Known | Event handler |
| 0x0022C9FC | `HandleCompilationsSelected` | Known | Event handler |
| 0x0022CA18 | `HandleArtistsSelected` | Known | Event handler |
| 0x0022CA30 | `HandleGenresSelected` | Known | Event handler |
| 0x0022CA48 | `HandleComposersSelected` | Known | Event handler |
| 0x0022CA60 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x0022CA7C | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x0022CA98 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x0022CAB0 | `HandleNowPlaying` | Known | Event handler |
| 0x0022CC5C | `HandleTVShowsSelected` | Known | Event handler |
| 0x0022CC74 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x0022CC90 | `HandleMoviesSelected` | Known | Event handler |
| 0x0022CCA8 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x0022CCC8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0022CCE0 | `HandleRadioPlayPause` | Known | Event handler |
| 0x0022CCF8 | `HandleLock` | Known | Event handler |
| 0x0022CD04 | `HandleBacklightSelected` | Known | Event handler |
| 0x0022CD1C | `HandleSleepSelected` | Known | Event handler |
| 0x0022CD30 | `HandleNikePlusSelected` | Known | Event handler |
| 0x0022F490 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0022FAA4 | `Handle400MetersWalk` | Known | Event handler |
| 0x0022FABC | `HandleCustomWalk` | Known | Event handler |
| 0x0022FAD0 | `HandleResetToDefault` | Known | Event handler |
| 0x0022FDBC | `HandleSelect` | Known | Event handler |
| 0x0023006C | `HandleWheel` | Known | Event handler |
| 0x0023146C | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x002316C4 | `HandleNextDay` | Known | Event handler |
| 0x002316D8 | `HandlePreviousDay` | Known | Event handler |
| 0x00231920 | `HandleSelect` | Known | Event handler |
| 0x00231BBC | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00234494 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x002344B0 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x00235418 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00235AF8 | `HandleSelect` | Known | Event handler |
| 0x002361C4 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x00271154 | `HandleDeleteClock` | Known | Event handler |
| 0x0027116C | `HandleSelectClock` | Known | Event handler |
| 0x00271180 | `HandleHilited` | Known | Event handler |
| 0x00271190 | `HandleWheel` | Known | Event handler |
| 0x0027119C | `HandleSelectLozinch` | Known | Event handler |
| 0x003E0E36 | `HandleAudioFFDown` | Known | Event handler |
| 0x003E0E5F | `HandleAudioFFUp` | Known | Event handler |
| 0x003E0E8A | `HandleAudioMute` | Known | Event handler |
| 0x003E0EBD | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x003E0EF2 | `HandleAudioNext` | Known | Event handler |
| 0x003E0F22 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x003E0F59 | `HandleAudioNextChapter` | Known | Event handler |
| 0x003E0F93 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x003E0FC7 | `HandleAudioPause` | Known | Event handler |
| 0x003E0FF3 | `HandleAudioPlay` | Known | Event handler |
| 0x003E1021 | `HandleAudioPlayPause` | Known | Event handler |
| 0x003E1059 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x003E1092 | `HandleAudioPrevious` | Known | Event handler |
| 0x003E10C6 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x003E10FD | `HandleAudioPrevChapter` | Known | Event handler |
| 0x003E1137 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x003E116C | `HandleAudioRepeat` | Known | Event handler |
| 0x003E1198 | `HandleAudioRewDown` | Known | Event handler |
| 0x003E11C3 | `HandleAudioRewUp` | Known | Event handler |
| 0x003E11F2 | `HandleAudioShuffle` | Known | Event handler |
| 0x003E1220 | `HandleAudioStop` | Known | Event handler |
| 0x003E1251 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x003E1286 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x003E12BD | `HandleAudioVolumeUp` | Known | Event handler |
| 0x003E12EE | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x003E13A7 | `HandleNextPressAndHold` | Known | Event handler |
| 0x003E13D8 | `HandleNext` | Known | Event handler |
| 0x003E1410 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x003E144B | `HandlePlayPause` | Known | Event handler |
| 0x003E147F | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x003E14B4 | `HandlePrevious` | Known | Event handler |
| 0x003E1541 | `HandleRemoteBacklight` | Known | Event handler |
| 0x003E1579 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x003E15B3 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x003E15EC | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x003E1621 | `HandleRemoteEvent` | Known | Event handler |
| 0x003E164D | `HandleRemoteFFDown` | Known | Event handler |
| 0x003E1678 | `HandleRemoteFFUp` | Known | Event handler |
| 0x003E16A5 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x003E16D4 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x003E1703 | `HandleRemoteMute` | Known | Event handler |
| 0x003E1735 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x003E176E | `HandleRemoteNextChapter` | Known | Event handler |
| 0x003E17AA | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x003E17EA | `HandleRemoteOff` | Known | Event handler |
| 0x003E1813 | `HandleRemoteOff` | Known | Event handler |
| 0x003E183D | `HandleRemoteOn` | Known | Event handler |
| 0x003E1869 | `HandleRemotePause` | Known | Event handler |
| 0x003E1897 | `HandleRemotePlay` | Known | Event handler |
| 0x003E18D5 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x003E1916 | `HandleRemotePlayPause` | Known | Event handler |
| 0x003E194D | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x003E1986 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x003E19C2 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x003E19F9 | `HandleRemoteRepeat` | Known | Event handler |
| 0x003E1A27 | `HandleRemoteRewDown` | Known | Event handler |
| 0x003E1A54 | `HandleRemoteRewUp` | Known | Event handler |
| 0x003E1A84 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x003E1AB7 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x003E1AEB | `HandleRemoteShuffle` | Known | Event handler |
| 0x003E1B1B | `HandleRemoteStop` | Known | Event handler |
| 0x003E1B4B | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x003E1B80 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x003E1BB8 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x003E1BEF | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x003E1C28 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x003E1C5B | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x003E1C90 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x003E1CC3 | `HandleVideoFFDown` | Known | Event handler |
| 0x003E1CEC | `HandleVideoFFUp` | Known | Event handler |
| 0x003E1D1F | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x003E1D54 | `HandleVideoNext` | Known | Event handler |
| 0x003E1D86 | `HandleVideoNextChapter` | Known | Event handler |
| 0x003E1DBD | `HandleVideoNextFrame` | Known | Event handler |
| 0x003E1DEE | `HandleVideoPause` | Known | Event handler |
| 0x003E1E1A | `HandleVideoPlay` | Known | Event handler |
| 0x003E1E48 | `HandleVideoPlayPause` | Known | Event handler |
| 0x003E1E80 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x003E1EB9 | `HandleVideoPrevious` | Known | Event handler |
| 0x003E1EEF | `HandleVideoPrevChapter` | Known | Event handler |
| 0x003E1F26 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x003E1F55 | `HandleVideoRewDown` | Known | Event handler |
| 0x003E1F80 | `HandleVideoRewUp` | Known | Event handler |
| 0x003E1FAC | `HandleVideoStop` | Known | Event handler |
| 0x0072E7EA | `HandleAddressBook` | Known | Event handler |
| 0x0072EC9E | `HandleSelect` | Known | Event handler |
| 0x0072ECD9 | `HandleHilite` | Known | Event handler |
| 0x0072ED5A | `HandleSelectRegion` | Known | Event handler |
| 0x0072EDFA | `HandleSelectRegion` | Known | Event handler |
| 0x0072EE96 | `HandleSelectRegion` | Known | Event handler |
| 0x0072EF3A | `HandleSelectRegion` | Known | Event handler |
| 0x0072EFE0 | `HandleSelectRegion` | Known | Event handler |
| 0x0072F080 | `HandleSelectRegion` | Known | Event handler |
| 0x0072F12C | `HandleSelectRegion` | Known | Event handler |
| 0x0072F1CE | `HandleSelectRegion` | Known | Event handler |
| 0x0072F27E | `HandleSelectCity` | Known | Event handler |
| 0x0072F2EA | `HandleHighlightCity` | Known | Event handler |
| 0x0072F323 | `HandleSelectCity` | Known | Event handler |
| 0x0072F38F | `HandleHighlightCity` | Known | Event handler |
| 0x0072F3C8 | `HandleSelectCity` | Known | Event handler |
| 0x0072F434 | `HandleHighlightCity` | Known | Event handler |
| 0x0072F46D | `HandleSelectCity` | Known | Event handler |
| 0x0072F4D9 | `HandleHighlightCity` | Known | Event handler |
| 0x0072F512 | `HandleSelectCity` | Known | Event handler |
| 0x0072F57E | `HandleHighlightCity` | Known | Event handler |
| 0x0072F5B7 | `HandleSelectCity` | Known | Event handler |
| 0x0072F623 | `HandleHighlightCity` | Known | Event handler |
| 0x0072F65C | `HandleSelectCity` | Known | Event handler |
| 0x0072F6C8 | `HandleHighlightCity` | Known | Event handler |
| 0x0072F701 | `HandleSelectCity` | Known | Event handler |
| 0x0072F76D | `HandleHighlightCity` | Known | Event handler |
| 0x0072F7A6 | `HandleSelectCity` | Known | Event handler |
| 0x0072F812 | `HandleHighlightCity` | Known | Event handler |
| 0x0072F84B | `HandleSelectCity` | Known | Event handler |
| 0x0072F8B7 | `HandleHighlightCity` | Known | Event handler |
| 0x0072F8F0 | `HandleSelectCity` | Known | Event handler |
| 0x0072F95C | `HandleHighlightCity` | Known | Event handler |
| 0x0072F995 | `HandleSelectCity` | Known | Event handler |
| 0x0072FA01 | `HandleHighlightCity` | Known | Event handler |
| 0x0072FA3A | `HandleSelectCity` | Known | Event handler |
| 0x0072FAA6 | `HandleHighlightCity` | Known | Event handler |
| 0x0072FADF | `HandleSelectCity` | Known | Event handler |
| 0x0072FB4B | `HandleHighlightCity` | Known | Event handler |
| 0x0072FB84 | `HandleSelectCity` | Known | Event handler |
| 0x0072FBF0 | `HandleHighlightCity` | Known | Event handler |
| 0x0072FC29 | `HandleSelectCity` | Known | Event handler |
| 0x0072FC95 | `HandleHighlightCity` | Known | Event handler |
| 0x0072FCCE | `HandleSelectCity` | Known | Event handler |
| 0x0072FD3A | `HandleHighlightCity` | Known | Event handler |
| 0x0072FD73 | `HandleSelectCity` | Known | Event handler |
| 0x0072FDDF | `HandleHighlightCity` | Known | Event handler |
| 0x0072FE18 | `HandleSelectCity` | Known | Event handler |
| 0x0072FE84 | `HandleHighlightCity` | Known | Event handler |
| 0x0072FEBD | `HandleSelectCity` | Known | Event handler |
| 0x0072FF29 | `HandleHighlightCity` | Known | Event handler |
| 0x0072FF62 | `HandleSelectCity` | Known | Event handler |
| 0x0072FFCE | `HandleHighlightCity` | Known | Event handler |
| 0x00730007 | `HandleSelectCity` | Known | Event handler |
| 0x00730073 | `HandleHighlightCity` | Known | Event handler |
| 0x007300AC | `HandleSelectCity` | Known | Event handler |
| 0x00730118 | `HandleHighlightCity` | Known | Event handler |
| 0x00730151 | `HandleSelectCity` | Known | Event handler |
| 0x007301BD | `HandleHighlightCity` | Known | Event handler |
| 0x007301F6 | `HandleSelectCity` | Known | Event handler |
| 0x00730262 | `HandleHighlightCity` | Known | Event handler |
| 0x0073029B | `HandleSelectCity` | Known | Event handler |
| 0x00730307 | `HandleHighlightCity` | Known | Event handler |
| 0x00730340 | `HandleSelectCity` | Known | Event handler |
| 0x007303AC | `HandleHighlightCity` | Known | Event handler |
| 0x007303E5 | `HandleSelectCity` | Known | Event handler |
| 0x00730451 | `HandleHighlightCity` | Known | Event handler |
| 0x0073048A | `HandleSelectCity` | Known | Event handler |
| 0x007304F6 | `HandleHighlightCity` | Known | Event handler |
| 0x0073052F | `HandleSelectCity` | Known | Event handler |
| 0x0073059B | `HandleHighlightCity` | Known | Event handler |
| 0x007305D4 | `HandleSelectCity` | Known | Event handler |
| 0x00730640 | `HandleHighlightCity` | Known | Event handler |
| 0x0073067E | `HandleSelectCity` | Known | Event handler |
| 0x007306EA | `HandleHighlightCity` | Known | Event handler |
| 0x00730723 | `HandleSelectCity` | Known | Event handler |
| 0x0073078F | `HandleHighlightCity` | Known | Event handler |
| 0x007307C8 | `HandleSelectCity` | Known | Event handler |
| 0x00730834 | `HandleHighlightCity` | Known | Event handler |
| 0x0073086D | `HandleSelectCity` | Known | Event handler |
| 0x007308D9 | `HandleHighlightCity` | Known | Event handler |
| 0x00730912 | `HandleSelectCity` | Known | Event handler |
| 0x0073097E | `HandleHighlightCity` | Known | Event handler |
| 0x007309B7 | `HandleSelectCity` | Known | Event handler |
| 0x00730A23 | `HandleHighlightCity` | Known | Event handler |
| 0x00730A5C | `HandleSelectCity` | Known | Event handler |
| 0x00730AC8 | `HandleHighlightCity` | Known | Event handler |
| 0x00730B01 | `HandleSelectCity` | Known | Event handler |
| 0x00730B6D | `HandleHighlightCity` | Known | Event handler |
| 0x00730BA6 | `HandleSelectCity` | Known | Event handler |
| 0x00730C12 | `HandleHighlightCity` | Known | Event handler |
| 0x00730C4B | `HandleSelectCity` | Known | Event handler |
| 0x00730CB7 | `HandleHighlightCity` | Known | Event handler |
| 0x00730CF0 | `HandleSelectCity` | Known | Event handler |
| 0x00730D5C | `HandleHighlightCity` | Known | Event handler |
| 0x00730D95 | `HandleSelectCity` | Known | Event handler |
| 0x00730E01 | `HandleHighlightCity` | Known | Event handler |
| 0x00730E3A | `HandleSelectCity` | Known | Event handler |
| 0x00730EA6 | `HandleHighlightCity` | Known | Event handler |
| 0x00730EDF | `HandleSelectCity` | Known | Event handler |
| 0x00730F4B | `HandleHighlightCity` | Known | Event handler |
| 0x00730F84 | `HandleSelectCity` | Known | Event handler |
| 0x00730FF0 | `HandleHighlightCity` | Known | Event handler |
| 0x00731029 | `HandleSelectCity` | Known | Event handler |
| 0x00731095 | `HandleHighlightCity` | Known | Event handler |
| 0x007310CE | `HandleSelectCity` | Known | Event handler |
| 0x0073113A | `HandleHighlightCity` | Known | Event handler |
| 0x00731173 | `HandleSelectCity` | Known | Event handler |
| 0x007311DF | `HandleHighlightCity` | Known | Event handler |
| 0x00731218 | `HandleSelectCity` | Known | Event handler |
| 0x00731284 | `HandleHighlightCity` | Known | Event handler |
| 0x007312BD | `HandleSelectCity` | Known | Event handler |
| 0x00731329 | `HandleHighlightCity` | Known | Event handler |
| 0x00731362 | `HandleSelectCity` | Known | Event handler |
| 0x007313CE | `HandleHighlightCity` | Known | Event handler |
| 0x00731407 | `HandleSelectCity` | Known | Event handler |
| 0x00731473 | `HandleHighlightCity` | Known | Event handler |
| 0x007314AC | `HandleSelectCity` | Known | Event handler |
| 0x00731518 | `HandleHighlightCity` | Known | Event handler |
| 0x00731551 | `HandleSelectCity` | Known | Event handler |
| 0x007315BD | `HandleHighlightCity` | Known | Event handler |
| 0x007315F6 | `HandleSelectCity` | Known | Event handler |
| 0x00731662 | `HandleHighlightCity` | Known | Event handler |
| 0x0073169B | `HandleSelectCity` | Known | Event handler |
| 0x00731707 | `HandleHighlightCity` | Known | Event handler |
| 0x00731740 | `HandleSelectCity` | Known | Event handler |
| 0x007317AC | `HandleHighlightCity` | Known | Event handler |
| 0x007317E5 | `HandleSelectCity` | Known | Event handler |
| 0x00731851 | `HandleHighlightCity` | Known | Event handler |
| 0x0073188A | `HandleSelectCity` | Known | Event handler |
| 0x007318F6 | `HandleHighlightCity` | Known | Event handler |
| 0x0073192F | `HandleSelectCity` | Known | Event handler |
| 0x0073199B | `HandleHighlightCity` | Known | Event handler |
| 0x007319D4 | `HandleSelectCity` | Known | Event handler |
| 0x00731A40 | `HandleHighlightCity` | Known | Event handler |
| 0x00731A79 | `HandleSelectCity` | Known | Event handler |
| 0x00731AE5 | `HandleHighlightCity` | Known | Event handler |
| 0x00731B1E | `HandleSelectCity` | Known | Event handler |
| 0x00731B8A | `HandleHighlightCity` | Known | Event handler |
| 0x00731BC3 | `HandleSelectCity` | Known | Event handler |
| 0x00731C2F | `HandleHighlightCity` | Known | Event handler |
| 0x00731C68 | `HandleSelectCity` | Known | Event handler |
| 0x00731CD4 | `HandleHighlightCity` | Known | Event handler |
| 0x00731D0D | `HandleSelectCity` | Known | Event handler |
| 0x00731D79 | `HandleHighlightCity` | Known | Event handler |
| 0x00731DB2 | `HandleSelectCity` | Known | Event handler |
| 0x00731E1E | `HandleHighlightCity` | Known | Event handler |
| 0x00731E57 | `HandleSelectCity` | Known | Event handler |
| 0x00731EC3 | `HandleHighlightCity` | Known | Event handler |
| 0x00731EFC | `HandleSelectCity` | Known | Event handler |
| 0x00731F68 | `HandleHighlightCity` | Known | Event handler |
| 0x00731FA1 | `HandleSelectCity` | Known | Event handler |
| 0x0073200D | `HandleHighlightCity` | Known | Event handler |
| 0x00732046 | `HandleSelectCity` | Known | Event handler |
| 0x007320B2 | `HandleHighlightCity` | Known | Event handler |
| 0x007320EB | `HandleSelectCity` | Known | Event handler |
| 0x00732157 | `HandleHighlightCity` | Known | Event handler |
| 0x00732190 | `HandleSelectCity` | Known | Event handler |
| 0x007321FC | `HandleHighlightCity` | Known | Event handler |
| 0x00732235 | `HandleSelectCity` | Known | Event handler |
| 0x007322A1 | `HandleHighlightCity` | Known | Event handler |
| 0x007322DA | `HandleSelectCity` | Known | Event handler |
| 0x00732346 | `HandleHighlightCity` | Known | Event handler |
| 0x0073237F | `HandleSelectCity` | Known | Event handler |
| 0x007323EB | `HandleHighlightCity` | Known | Event handler |
| 0x00732424 | `HandleSelectCity` | Known | Event handler |
| 0x00732490 | `HandleHighlightCity` | Known | Event handler |
| 0x007324C9 | `HandleSelectCity` | Known | Event handler |
| 0x00732535 | `HandleHighlightCity` | Known | Event handler |
| 0x0073256E | `HandleSelectCity` | Known | Event handler |
| 0x007325DA | `HandleHighlightCity` | Known | Event handler |
| 0x00732613 | `HandleSelectCity` | Known | Event handler |
| 0x0073267F | `HandleHighlightCity` | Known | Event handler |
| 0x007326B8 | `HandleSelectCity` | Known | Event handler |
| 0x00732724 | `HandleHighlightCity` | Known | Event handler |
| 0x0073275D | `HandleSelectCity` | Known | Event handler |
| 0x007327C9 | `HandleHighlightCity` | Known | Event handler |
| 0x00732802 | `HandleSelectCity` | Known | Event handler |
| 0x0073286E | `HandleHighlightCity` | Known | Event handler |
| 0x007328A7 | `HandleSelectCity` | Known | Event handler |
| 0x00732913 | `HandleHighlightCity` | Known | Event handler |
| 0x0073294C | `HandleSelectCity` | Known | Event handler |
| 0x007329B8 | `HandleHighlightCity` | Known | Event handler |
| 0x007329F1 | `HandleSelectCity` | Known | Event handler |
| 0x00732A5D | `HandleHighlightCity` | Known | Event handler |
| 0x00732A96 | `HandleSelectCity` | Known | Event handler |
| 0x00732B02 | `HandleHighlightCity` | Known | Event handler |
| 0x00732B42 | `HandleSelectCity` | Known | Event handler |
| 0x00732BAE | `HandleHighlightCity` | Known | Event handler |
| 0x00732BE7 | `HandleSelectCity` | Known | Event handler |
| 0x00732C53 | `HandleHighlightCity` | Known | Event handler |
| 0x00732C8C | `HandleSelectCity` | Known | Event handler |
| 0x00732CF8 | `HandleHighlightCity` | Known | Event handler |
| 0x00732D36 | `HandleSelectCity` | Known | Event handler |
| 0x00732DA2 | `HandleHighlightCity` | Known | Event handler |
| 0x00732DDB | `HandleSelectCity` | Known | Event handler |
| 0x00732E47 | `HandleHighlightCity` | Known | Event handler |
| 0x00732E80 | `HandleSelectCity` | Known | Event handler |
| 0x00732EEC | `HandleHighlightCity` | Known | Event handler |
| 0x00732F25 | `HandleSelectCity` | Known | Event handler |
| 0x00732F91 | `HandleHighlightCity` | Known | Event handler |
| 0x00732FCA | `HandleSelectCity` | Known | Event handler |
| 0x00733036 | `HandleHighlightCity` | Known | Event handler |
| 0x0073306F | `HandleSelectCity` | Known | Event handler |
| 0x007330DB | `HandleHighlightCity` | Known | Event handler |
| 0x00733114 | `HandleSelectCity` | Known | Event handler |
| 0x00733180 | `HandleHighlightCity` | Known | Event handler |
| 0x007331B9 | `HandleSelectCity` | Known | Event handler |
| 0x00733225 | `HandleHighlightCity` | Known | Event handler |
| 0x00733262 | `HandleSelectCity` | Known | Event handler |
| 0x007332CE | `HandleHighlightCity` | Known | Event handler |
| 0x00733307 | `HandleSelectCity` | Known | Event handler |
| 0x00733373 | `HandleHighlightCity` | Known | Event handler |
| 0x007333AC | `HandleSelectCity` | Known | Event handler |
| 0x00733418 | `HandleHighlightCity` | Known | Event handler |
| 0x00733451 | `HandleSelectCity` | Known | Event handler |
| 0x007334BD | `HandleHighlightCity` | Known | Event handler |
| 0x007334F6 | `HandleSelectCity` | Known | Event handler |
| 0x00733562 | `HandleHighlightCity` | Known | Event handler |
| 0x0073359B | `HandleSelectCity` | Known | Event handler |
| 0x00733607 | `HandleHighlightCity` | Known | Event handler |
| 0x00733640 | `HandleSelectCity` | Known | Event handler |
| 0x007336AC | `HandleHighlightCity` | Known | Event handler |
| 0x007336E5 | `HandleSelectCity` | Known | Event handler |
| 0x00733751 | `HandleHighlightCity` | Known | Event handler |
| 0x0073378A | `HandleSelectCity` | Known | Event handler |
| 0x007337F6 | `HandleHighlightCity` | Known | Event handler |
| 0x0073382F | `HandleSelectCity` | Known | Event handler |
| 0x0073389B | `HandleHighlightCity` | Known | Event handler |
| 0x007338D4 | `HandleSelectCity` | Known | Event handler |
| 0x00733940 | `HandleHighlightCity` | Known | Event handler |
| 0x00733979 | `HandleSelectCity` | Known | Event handler |
| 0x007339E5 | `HandleHighlightCity` | Known | Event handler |
| 0x00733A1E | `HandleSelectCity` | Known | Event handler |
| 0x00733A8A | `HandleHighlightCity` | Known | Event handler |
| 0x00733AC3 | `HandleSelectCity` | Known | Event handler |
| 0x00733B2F | `HandleHighlightCity` | Known | Event handler |
| 0x00733B68 | `HandleSelectCity` | Known | Event handler |
| 0x00733BD4 | `HandleHighlightCity` | Known | Event handler |
| 0x00733C0D | `HandleSelectCity` | Known | Event handler |
| 0x00733C79 | `HandleHighlightCity` | Known | Event handler |
| 0x00733CB2 | `HandleSelectCity` | Known | Event handler |
| 0x00733D1E | `HandleHighlightCity` | Known | Event handler |
| 0x00733D57 | `HandleSelectCity` | Known | Event handler |
| 0x00733DC3 | `HandleHighlightCity` | Known | Event handler |
| 0x00733DFC | `HandleSelectCity` | Known | Event handler |
| 0x00733E68 | `HandleHighlightCity` | Known | Event handler |
| 0x00733EA1 | `HandleSelectCity` | Known | Event handler |
| 0x00733F0D | `HandleHighlightCity` | Known | Event handler |
| 0x00733F46 | `HandleSelectCity` | Known | Event handler |
| 0x00733FB2 | `HandleHighlightCity` | Known | Event handler |
| 0x00733FEB | `HandleSelectCity` | Known | Event handler |
| 0x00734057 | `HandleHighlightCity` | Known | Event handler |
| 0x00734090 | `HandleSelectCity` | Known | Event handler |
| 0x007340FC | `HandleHighlightCity` | Known | Event handler |
| 0x00734135 | `HandleSelectCity` | Known | Event handler |
| 0x007341A1 | `HandleHighlightCity` | Known | Event handler |
| 0x007341DA | `HandleSelectCity` | Known | Event handler |
| 0x00734246 | `HandleHighlightCity` | Known | Event handler |
| 0x0073427F | `HandleSelectCity` | Known | Event handler |
| 0x007342EB | `HandleHighlightCity` | Known | Event handler |
| 0x00734324 | `HandleSelectCity` | Known | Event handler |
| 0x00734390 | `HandleHighlightCity` | Known | Event handler |
| 0x007343C9 | `HandleSelectCity` | Known | Event handler |
| 0x00734435 | `HandleHighlightCity` | Known | Event handler |
| 0x0073446E | `HandleSelectCity` | Known | Event handler |
| 0x007344DA | `HandleHighlightCity` | Known | Event handler |
| 0x00734513 | `HandleSelectCity` | Known | Event handler |
| 0x0073457F | `HandleHighlightCity` | Known | Event handler |
| 0x007345B8 | `HandleSelectCity` | Known | Event handler |
| 0x00734624 | `HandleHighlightCity` | Known | Event handler |
| 0x0073465D | `HandleSelectCity` | Known | Event handler |
| 0x007346C9 | `HandleHighlightCity` | Known | Event handler |
| 0x00734702 | `HandleSelectCity` | Known | Event handler |
| 0x0073476E | `HandleHighlightCity` | Known | Event handler |
| 0x007347A7 | `HandleSelectCity` | Known | Event handler |
| 0x00734813 | `HandleHighlightCity` | Known | Event handler |
| 0x00734852 | `HandleSelectCity` | Known | Event handler |
| 0x007348BE | `HandleHighlightCity` | Known | Event handler |
| 0x007348F7 | `HandleSelectCity` | Known | Event handler |
| 0x00734963 | `HandleHighlightCity` | Known | Event handler |
| 0x0073499C | `HandleSelectCity` | Known | Event handler |
| 0x00734A08 | `HandleHighlightCity` | Known | Event handler |
| 0x00734A41 | `HandleSelectCity` | Known | Event handler |
| 0x00734AAD | `HandleHighlightCity` | Known | Event handler |
| 0x00734AE6 | `HandleSelectCity` | Known | Event handler |
| 0x00734B52 | `HandleHighlightCity` | Known | Event handler |
| 0x00734B8B | `HandleSelectCity` | Known | Event handler |
| 0x00734BF7 | `HandleHighlightCity` | Known | Event handler |
| 0x00734C30 | `HandleSelectCity` | Known | Event handler |
| 0x00734C9C | `HandleHighlightCity` | Known | Event handler |
| 0x00734CD5 | `HandleSelectCity` | Known | Event handler |
| 0x00734D41 | `HandleHighlightCity` | Known | Event handler |
| 0x00734D7A | `HandleSelectCity` | Known | Event handler |
| 0x00734DE6 | `HandleHighlightCity` | Known | Event handler |
| 0x00734E1F | `HandleSelectCity` | Known | Event handler |
| 0x00734E8B | `HandleHighlightCity` | Known | Event handler |
| 0x00734EC4 | `HandleSelectCity` | Known | Event handler |
| 0x00734F30 | `HandleHighlightCity` | Known | Event handler |
| 0x00734F69 | `HandleSelectCity` | Known | Event handler |
| 0x00734FD5 | `HandleHighlightCity` | Known | Event handler |
| 0x0073500E | `HandleSelectCity` | Known | Event handler |
| 0x0073507A | `HandleHighlightCity` | Known | Event handler |
| 0x007350B3 | `HandleSelectCity` | Known | Event handler |
| 0x0073511F | `HandleHighlightCity` | Known | Event handler |
| 0x00735158 | `HandleSelectCity` | Known | Event handler |
| 0x007351C4 | `HandleHighlightCity` | Known | Event handler |
| 0x007351FD | `HandleSelectCity` | Known | Event handler |
| 0x00735269 | `HandleHighlightCity` | Known | Event handler |
| 0x007352A2 | `HandleSelectCity` | Known | Event handler |
| 0x0073530E | `HandleHighlightCity` | Known | Event handler |
| 0x00735347 | `HandleSelectCity` | Known | Event handler |
| 0x007353B3 | `HandleHighlightCity` | Known | Event handler |
| 0x007353EC | `HandleSelectCity` | Known | Event handler |
| 0x00735458 | `HandleHighlightCity` | Known | Event handler |
| 0x00735491 | `HandleSelectCity` | Known | Event handler |
| 0x007354FD | `HandleHighlightCity` | Known | Event handler |
| 0x00735536 | `HandleSelectCity` | Known | Event handler |
| 0x007355A2 | `HandleHighlightCity` | Known | Event handler |
| 0x007355DB | `HandleSelectCity` | Known | Event handler |
| 0x00735647 | `HandleHighlightCity` | Known | Event handler |
| 0x00735680 | `HandleSelectCity` | Known | Event handler |
| 0x007356EC | `HandleHighlightCity` | Known | Event handler |
| 0x00735725 | `HandleSelectCity` | Known | Event handler |
| 0x00735791 | `HandleHighlightCity` | Known | Event handler |
| 0x007357CA | `HandleSelectCity` | Known | Event handler |
| 0x00735836 | `HandleHighlightCity` | Known | Event handler |
| 0x0073586F | `HandleSelectCity` | Known | Event handler |
| 0x007358DB | `HandleHighlightCity` | Known | Event handler |
| 0x00735914 | `HandleSelectCity` | Known | Event handler |
| 0x00735980 | `HandleHighlightCity` | Known | Event handler |
| 0x007359B9 | `HandleSelectCity` | Known | Event handler |
| 0x00735A25 | `HandleHighlightCity` | Known | Event handler |
| 0x00735A5E | `HandleSelectCity` | Known | Event handler |
| 0x00735ACA | `HandleHighlightCity` | Known | Event handler |
| 0x00735B03 | `HandleSelectCity` | Known | Event handler |
| 0x00735B6F | `HandleHighlightCity` | Known | Event handler |
| 0x00735BA8 | `HandleSelectCity` | Known | Event handler |
| 0x00735C14 | `HandleHighlightCity` | Known | Event handler |
| 0x00735C4D | `HandleSelectCity` | Known | Event handler |
| 0x00735CB9 | `HandleHighlightCity` | Known | Event handler |
| 0x00735CF2 | `HandleSelectCity` | Known | Event handler |
| 0x00735D5E | `HandleHighlightCity` | Known | Event handler |
| 0x00735D97 | `HandleSelectCity` | Known | Event handler |
| 0x00735E03 | `HandleHighlightCity` | Known | Event handler |
| 0x00735E3C | `HandleSelectCity` | Known | Event handler |
| 0x00735EA8 | `HandleHighlightCity` | Known | Event handler |
| 0x00735EE1 | `HandleSelectCity` | Known | Event handler |
| 0x00735F4D | `HandleHighlightCity` | Known | Event handler |
| 0x00735F86 | `HandleSelectCity` | Known | Event handler |
| 0x00735FF2 | `HandleHighlightCity` | Known | Event handler |
| 0x0073602B | `HandleSelectCity` | Known | Event handler |
| 0x00736097 | `HandleHighlightCity` | Known | Event handler |
| 0x007360D0 | `HandleSelectCity` | Known | Event handler |
| 0x0073613C | `HandleHighlightCity` | Known | Event handler |
| 0x00736175 | `HandleSelectCity` | Known | Event handler |
| 0x007361E1 | `HandleHighlightCity` | Known | Event handler |
| 0x0073621A | `HandleSelectCity` | Known | Event handler |
| 0x00736286 | `HandleHighlightCity` | Known | Event handler |
| 0x007362BF | `HandleSelectCity` | Known | Event handler |
| 0x0073632B | `HandleHighlightCity` | Known | Event handler |
| 0x00736364 | `HandleSelectCity` | Known | Event handler |
| 0x007363D0 | `HandleHighlightCity` | Known | Event handler |
| 0x00736409 | `HandleSelectCity` | Known | Event handler |
| 0x00736475 | `HandleHighlightCity` | Known | Event handler |
| 0x007364AE | `HandleSelectCity` | Known | Event handler |
| 0x0073651A | `HandleHighlightCity` | Known | Event handler |
| 0x00736553 | `HandleSelectCity` | Known | Event handler |
| 0x007365BF | `HandleHighlightCity` | Known | Event handler |
| 0x007365F8 | `HandleSelectCity` | Known | Event handler |
| 0x00736664 | `HandleHighlightCity` | Known | Event handler |
| 0x0073669D | `HandleSelectCity` | Known | Event handler |
| 0x00736709 | `HandleHighlightCity` | Known | Event handler |
| 0x00736742 | `HandleSelectCity` | Known | Event handler |
| 0x007367AE | `HandleHighlightCity` | Known | Event handler |
| 0x007367E7 | `HandleSelectCity` | Known | Event handler |
| 0x00736853 | `HandleHighlightCity` | Known | Event handler |
| 0x00736892 | `HandleSelectCity` | Known | Event handler |
| 0x007368FE | `HandleHighlightCity` | Known | Event handler |
| 0x00736937 | `HandleSelectCity` | Known | Event handler |
| 0x007369A3 | `HandleHighlightCity` | Known | Event handler |
| 0x007369DC | `HandleSelectCity` | Known | Event handler |
| 0x00736A48 | `HandleHighlightCity` | Known | Event handler |
| 0x00736A81 | `HandleSelectCity` | Known | Event handler |
| 0x00736AED | `HandleHighlightCity` | Known | Event handler |
| 0x00736B26 | `HandleSelectCity` | Known | Event handler |
| 0x00736B92 | `HandleHighlightCity` | Known | Event handler |
| 0x00736BD2 | `HandleSelectCity` | Known | Event handler |
| 0x00736C3E | `HandleHighlightCity` | Known | Event handler |
| 0x00736C77 | `HandleSelectCity` | Known | Event handler |
| 0x00736CE3 | `HandleHighlightCity` | Known | Event handler |
| 0x00736D1C | `HandleSelectCity` | Known | Event handler |
| 0x00736D88 | `HandleHighlightCity` | Known | Event handler |
| 0x00736DC1 | `HandleSelectCity` | Known | Event handler |
| 0x00736E2D | `HandleHighlightCity` | Known | Event handler |
| 0x00736E66 | `HandleSelectCity` | Known | Event handler |
| 0x00736ED2 | `HandleHighlightCity` | Known | Event handler |
| 0x00736F0B | `HandleSelectCity` | Known | Event handler |
| 0x00736F77 | `HandleHighlightCity` | Known | Event handler |
| 0x00736FB0 | `HandleSelectCity` | Known | Event handler |
| 0x0073701C | `HandleHighlightCity` | Known | Event handler |
| 0x00737055 | `HandleSelectCity` | Known | Event handler |
| 0x007370C1 | `HandleHighlightCity` | Known | Event handler |
| 0x007370FA | `HandleSelectCity` | Known | Event handler |
| 0x00737166 | `HandleHighlightCity` | Known | Event handler |
| 0x0073719F | `HandleSelectCity` | Known | Event handler |
| 0x0073720B | `HandleHighlightCity` | Known | Event handler |
| 0x00737244 | `HandleSelectCity` | Known | Event handler |
| 0x007372B0 | `HandleHighlightCity` | Known | Event handler |
| 0x007372E9 | `HandleSelectCity` | Known | Event handler |
| 0x00737355 | `HandleHighlightCity` | Known | Event handler |
| 0x0073738E | `HandleSelectCity` | Known | Event handler |
| 0x007373FA | `HandleHighlightCity` | Known | Event handler |
| 0x00737433 | `HandleSelectCity` | Known | Event handler |
| 0x0073749F | `HandleHighlightCity` | Known | Event handler |
| 0x007374D8 | `HandleSelectCity` | Known | Event handler |
| 0x00737544 | `HandleHighlightCity` | Known | Event handler |
| 0x0073757D | `HandleSelectCity` | Known | Event handler |
| 0x007375E9 | `HandleHighlightCity` | Known | Event handler |
| 0x00737622 | `HandleSelectCity` | Known | Event handler |
| 0x0073768E | `HandleHighlightCity` | Known | Event handler |
| 0x00737B86 | `HandleMusicSelected` | Known | Event handler |
| 0x00737BC8 | `HandleMusicHilited` | Known | Event handler |
| 0x00737C00 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00737C46 | `HandleMusicHilited` | Known | Event handler |
| 0x00737C7E | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00737CC4 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x00737D00 | `HandleArtistsSelected` | Known | Event handler |
| 0x00737D44 | `HandleArtistsHilited` | Known | Event handler |
| 0x00737D7E | `HandleAlbumsSelected` | Known | Event handler |
| 0x00737DC1 | `HandleAlbumsHilited` | Known | Event handler |
| 0x00737DFA | `HandleCompilationsSelected` | Known | Event handler |
| 0x00737E43 | `HandleCompilationsHilited` | Known | Event handler |
| 0x00737E82 | `HandleSongsSelected` | Known | Event handler |
| 0x00737EC4 | `HandleSongsHilited` | Known | Event handler |
| 0x00737EFC | `HandleGenresSelected` | Known | Event handler |
| 0x00737F3F | `HandleGenresHilited` | Known | Event handler |
| 0x00737F78 | `HandleComposersSelected` | Known | Event handler |
| 0x00737FBE | `HandleComposersHilited` | Known | Event handler |
| 0x00737FFA | `HandleAudiobooksSelected` | Known | Event handler |
| 0x00738041 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00738100 | `HandleMusicHilited` | Known | Event handler |
| 0x00738138 | `HandleVideosSelected` | Known | Event handler |
| 0x0073817B | `HandleVideosHilited` | Known | Event handler |
| 0x007381B4 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x007381FF | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00738240 | `HandleMoviesSelected` | Known | Event handler |
| 0x00738283 | `HandleMoviesHilited` | Known | Event handler |
| 0x007382BC | `HandleTVShowsSelected` | Known | Event handler |
| 0x00738300 | `HandleTVShowsHilited` | Known | Event handler |
| 0x0073833A | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00738382 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x007383C0 | `HandlePhotosSelected` | Known | Event handler |
| 0x00738403 | `HandlePhotosHilited` | Known | Event handler |
| 0x0073843C | `HandlePhotosSelected` | Known | Event handler |
| 0x0073847F | `HandlePhotosHilited` | Known | Event handler |
| 0x007384B8 | `HandlePodcastsSelected` | Known | Event handler |
| 0x007384FD | `HandlePodcastsHilited` | Known | Event handler |
| 0x007385B0 | `HandleGenericHilited` | Known | Event handler |
| 0x007386A9 | `HandleGenericHilited` | Known | Event handler |
| 0x00738B8E | `HandleLock` | Known | Event handler |
| 0x00738CFF | `HandleNikePlusSelected` | Known | Event handler |
| 0x00738D44 | `HandleGenericHilited` | Known | Event handler |
| 0x00738E4A | `HandleGenericHilited` | Known | Event handler |
| 0x00738F49 | `HandleGenericHilited` | Known | Event handler |
| 0x00739036 | `HandleGenericHilited` | Known | Event handler |
| 0x007390B0 | `HandleRadioPlayPause` | Known | Event handler |
| 0x00739170 | `HandleGenericHilited` | Known | Event handler |
| 0x007391EA | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00739233 | `HandleGenericHilited` | Known | Event handler |
| 0x007392AC | `HandleBacklightSelected` | Known | Event handler |
| 0x007392F2 | `HandleGenericHilited` | Known | Event handler |
| 0x0073936D | `HandleSleepSelected` | Known | Event handler |
| 0x007393AF | `HandleGenericHilited` | Known | Event handler |
| 0x00739426 | `HandleNowPlaying` | Known | Event handler |
| 0x0073949E | `HandleNowPlayingHilited` | Known | Event handler |
| 0x007394E2 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00739528 | `HandleMusicHilited` | Known | Event handler |
| 0x00739560 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x007395A6 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x007395E4 | `HandleArtistsSelected` | Known | Event handler |
| 0x00739628 | `HandleArtistsHilited` | Known | Event handler |
| 0x00739662 | `HandleAlbumsSelected` | Known | Event handler |
| 0x007396A5 | `HandleAlbumsHilited` | Known | Event handler |
| 0x007396DE | `HandleCompilationsSelected` | Known | Event handler |
| 0x00739727 | `HandleCompilationsHilited` | Known | Event handler |
| 0x00739766 | `HandleSongsSelected` | Known | Event handler |
| 0x007397A8 | `HandleSongsHilited` | Known | Event handler |
| 0x00739853 | `HandleGenericHilited` | Known | Event handler |
| 0x007398CE | `HandleRadioPlayPause` | Known | Event handler |
| 0x00739908 | `HandleGenresSelected` | Known | Event handler |
| 0x0073994B | `HandleGenresHilited` | Known | Event handler |
| 0x00739984 | `HandleComposersSelected` | Known | Event handler |
| 0x007399CA | `HandleComposersHilited` | Known | Event handler |
| 0x00739A06 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x00739A4D | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00739B0C | `HandleMusicHilited` | Known | Event handler |
| 0x00739B81 | `HandlePlayPause` | Known | Event handler |
| 0x00739BB6 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x00739CA0 | `HandleSelect` | Known | Event handler |
| 0x00739CE2 | `HandleMoviesSelected` | Known | Event handler |
| 0x00739D25 | `HandleMoviesHilited` | Known | Event handler |
| 0x00739D5E | `HandleTVShowsSelected` | Known | Event handler |
| 0x00739DA2 | `HandleTVShowsHilited` | Known | Event handler |
| 0x00739DDC | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00739E24 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00739E62 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00739EAD | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00739F73 | `HandleVideosHilited` | Known | Event handler |
| 0x0073A601 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x0073B18A | `HandleMainMenu` | Known | Event handler |
| 0x0073B1C2 | `HandleMusicMenu` | Known | Event handler |
| 0x0073B6EA | `HandleRadioRegion` | Known | Event handler |
| 0x0073B78E | `HandleLanguage` | Known | Event handler |
| 0x0073B894 | `HandleNew` | Known | Event handler |
| 0x0073B90F | `HandleClear` | Known | Event handler |
| 0x0073B940 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x0073B9FC | `HandleSelectIndexedSession` | Known | Event handler |
| 0x0073BBC3 | `HandleBasicSelected` | Known | Event handler |
| 0x0073BC69 | `HandleTimedWorkoutSelected` | Known | Event handler |
| 0x0073BD16 | `HandleDistanceWorkoutSelected` | Known | Event handler |
| 0x0073BDC6 | `HandleCaloriesWorkoutSelected` | Known | Event handler |
| 0x0073C1A1 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x0073C1F4 | `HandleSelect` | Known | Event handler |
| 0x0073C31E | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x0073C358 | `HandleEQSettingSelected` | Known | Event handler |
| 0x0073C390 | `HandleEQSettingSelected` | Known | Event handler |
| 0x00750B00 | `HandleItemSelected` | Known | Event handler |
| 0x00750C4B | `HandleNextContact` | Known | Event handler |
| 0x00750C77 | `HandlePreviousContact` | Known | Event handler |
| 0x00750CAD | `HandleSelectKey` | Known | Event handler |
| 0x007512BE | `HandleSelect` | Known | Event handler |
| 0x007515E5 | `HandleDateChosen` | Known | Event handler |
| 0x0075161B | `HandleTimeChosen` | Known | Event handler |
| 0x00751651 | `HandleFrequencyChosen` | Known | Event handler |
| 0x0075168C | `HandleSoundChosen` | Known | Event handler |
| 0x007516C3 | `HandleLabelChosen` | Known | Event handler |
| 0x007516FA | `HandleDeleteChosen` | Known | Event handler |
| 0x00751736 | `HandleSelect` | Known | Event handler |
| 0x0075176E | `HandleSelect` | Known | Event handler |
| 0x00751AAF | `HandleLeaveAlarm` | Known | Event handler |
| 0x00751ADC | `HandleLeaveAlarm` | Known | Event handler |
| 0x00751B0B | `HandleLeaveAlarm` | Known | Event handler |
| 0x00751B38 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00751C72 | `HandleSelect` | Known | Event handler |
| 0x00751CA0 | `HandleSelect` | Known | Event handler |
| 0x00751DFF | `HandleNextDay` | Known | Event handler |
| 0x00751E27 | `HandlePreviousDay` | Known | Event handler |
| 0x00751FD6 | `HandleSelect` | Known | Event handler |
| 0x00752003 | `HandleNextDay` | Known | Event handler |
| 0x0075202B | `HandlePreviousDay` | Known | Event handler |
| 0x007521D3 | `HandleNextDay` | Known | Event handler |
| 0x007521FB | `HandlePreviousDay` | Known | Event handler |
| 0x007522BC | `HandleSelect` | Known | Event handler |
| 0x007522E7 | `HandleNextDay` | Known | Event handler |
| 0x0075230F | `HandlePreviousDay` | Known | Event handler |
| 0x00752486 | `HandleSelectLozinch` | Known | Event handler |
| 0x007525FE | `HandleSelectLozinch` | Known | Event handler |
| 0x0075271D | `HandleFlowNext` | Known | Event handler |
| 0x0075274B | `HandlePlayPause` | Known | Event handler |
| 0x0075279A | `HandleFlowPrev` | Known | Event handler |
| 0x007527C5 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x007528B9 | `HandleAlbumSelected` | Known | Event handler |
| 0x00752A54 | `HandleFlowNext` | Known | Event handler |
| 0x00752AA2 | `HandleFlowNext` | Known | Event handler |
| 0x00752AD0 | `HandlePlayPause` | Known | Event handler |
| 0x00752B1F | `HandleFlowPrev` | Known | Event handler |
| 0x00752B4B | `HandleFlowPrev` | Known | Event handler |
| 0x00752B6B | `HandleFlowWheel` | Known | Event handler |
| 0x00752EFB | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x00753326 | `HandleArrowDown` | Known | Event handler |
| 0x00753390 | `HandleArrowUp` | Known | Event handler |
| 0x007533AF | `HandleWheel` | Known | Event handler |
| 0x00753438 | `HandleSelect` | Known | Event handler |
| 0x007534B5 | `HandleGameHilited` | Known | Event handler |
| 0x00756917 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007583EB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00759EBF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075B993 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075D467 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075EF3B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00760A0F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007624E3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00763FB7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00765A8B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076755F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00769033 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076AB07 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076C5DB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076E0AF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076FB83 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00771657 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077312B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00774BFF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007766D3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007781A7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00779C7B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077B74F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077D223 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077ECF7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007807CB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078229F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00783D73 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00785847 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078731B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00788DEF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078A8C3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078C397 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078DE6B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078F93F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00791413 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00792ECC | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00793A54 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007945DC | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00795164 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00795CEC | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00796874 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007973FC | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00797F84 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00798B0C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00799694 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079A21C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079ADA4 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0079B92C | `HandlePlayPause` | Known | Event handler |
| 0x0079B962 | `HandleAddToOTG` | Known | Event handler |
| 0x0079BAFF | `HandlePlayPause` | Known | Event handler |
| 0x0079BB26 | `HandleSelect` | Known | Event handler |
| 0x0079BB53 | `HandleHilite` | Known | Event handler |
| 0x0079BB84 | `HandlePlayPause` | Known | Event handler |
| 0x0079BC17 | `HandlePlayPause` | Known | Event handler |
| 0x0079BC3E | `HandleSelect` | Known | Event handler |
| 0x0079BCA4 | `HandleHilite` | Known | Event handler |
| 0x0079BCD6 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x0079BD20 | `HandlePlayPause` | Known | Event handler |
| 0x0079BD56 | `HandleAddToOTG` | Known | Event handler |
| 0x0079BDE8 | `HandlePlayPause` | Known | Event handler |
| 0x0079BE0F | `HandleSelect` | Known | Event handler |
| 0x0079BE78 | `HandlePlayPause` | Known | Event handler |
| 0x0079BEAE | `HandleAddToOTG` | Known | Event handler |
| 0x0079BF40 | `HandlePlayPause` | Known | Event handler |
| 0x0079BF67 | `HandleSelect` | Known | Event handler |
| 0x0079BFD0 | `HandlePlayPause` | Known | Event handler |
| 0x0079C056 | `HandleSelect` | Known | Event handler |
| 0x0079C0BB | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079C0FC | `HandlePlayPause` | Known | Event handler |
| 0x0079C132 | `HandleAddToOTG` | Known | Event handler |
| 0x0079C364 | `HandlePlayPause` | Known | Event handler |
| 0x0079C38B | `HandleSelect` | Known | Event handler |
| 0x0079C3B8 | `HandleHilite` | Known | Event handler |
| 0x0079C3E8 | `HandlePlayPause` | Known | Event handler |
| 0x0079C41E | `HandleAddToOTG` | Known | Event handler |
| 0x0079C650 | `HandlePlayPause` | Known | Event handler |
| 0x0079C677 | `HandleSelect` | Known | Event handler |
| 0x0079C6A4 | `HandleHilite` | Known | Event handler |
| 0x0079C6D4 | `HandlePlayPause` | Known | Event handler |
| 0x0079C70A | `HandleAddToOTG` | Known | Event handler |
| 0x0079C9F5 | `HandlePlayPause` | Known | Event handler |
| 0x0079CA1C | `HandleSelect` | Known | Event handler |
| 0x0079CA4C | `HandlePlayPause` | Known | Event handler |
| 0x0079CA82 | `HandleAddToOTG` | Known | Event handler |
| 0x0079CB14 | `HandlePlayPause` | Known | Event handler |
| 0x0079CB3B | `HandleSelect` | Known | Event handler |
| 0x0079CBCC | `HandlePlayPause` | Known | Event handler |
| 0x0079CC02 | `HandleAddToOTG` | Known | Event handler |
| 0x0079CDBB | `HandlePlayPause` | Known | Event handler |
| 0x0079CDE2 | `HandleSelect` | Known | Event handler |
| 0x0079CE14 | `HandlePlayPause` | Known | Event handler |
| 0x0079CE4A | `HandleAddToOTG` | Known | Event handler |
| 0x0079CECF | `HandleSelect` | Known | Event handler |
| 0x0079CF68 | `HandleHilite` | Known | Event handler |
| 0x0079CF94 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079CFD8 | `HandlePlayPause` | Known | Event handler |
| 0x0079D00E | `HandleAddToOTG` | Known | Event handler |
| 0x0079D093 | `HandleSelect` | Known | Event handler |
| 0x0079D0F8 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079D13C | `HandlePlayPause` | Known | Event handler |
| 0x0079D2E0 | `HandleSelect` | Known | Event handler |
| 0x0079D30D | `HandleHilite` | Known | Event handler |
| 0x0079D339 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079D37C | `HandlePlayPause` | Known | Event handler |
| 0x0079D402 | `HandleSelect` | Known | Event handler |
| 0x0079D490 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079D4D4 | `HandlePlayPause` | Known | Event handler |
| 0x0079D55A | `HandleSelect` | Known | Event handler |
| 0x0079D5BF | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079D600 | `HandlePlayPause` | Known | Event handler |
| 0x0079D686 | `HandleSelect` | Known | Event handler |
| 0x0079D6EC | `HandleHilite` | Known | Event handler |
| 0x0079D718 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079D75C | `HandlePlayPause` | Known | Event handler |
| 0x0079D792 | `HandleAddToOTG` | Known | Event handler |
| 0x0079D955 | `HandlePlayPause` | Known | Event handler |
| 0x0079D97C | `HandleSelect` | Known | Event handler |
| 0x0079D9AC | `HandlePlayPause` | Known | Event handler |
| 0x0079D9E2 | `HandleAddToOTG` | Known | Event handler |
| 0x0079DC03 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x0079DD1C | `HandlePlayPause` | Known | Event handler |
| 0x0079DE49 | `HandleSelect` | Known | Event handler |
| 0x0079DE75 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079DEB8 | `HandlePlayPause` | Known | Event handler |
| 0x0079DFEB | `HandleSelect` | Known | Event handler |
| 0x0079E017 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079E7E5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079EF59 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079F6CD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079FE41 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007A05B5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007A0D29 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007A149D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007A14E6 | `HandleTVOutChanged` | Known | Event handler |
| 0x007A151E | `HandleTVSignalChanged` | Known | Event handler |
| 0x007A1559 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x007A15AA | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x007A15EF | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x007A1631 | `HandleSelect` | Known | Event handler |
| 0x007A1661 | `HandleSelect` | Known | Event handler |
| 0x007A16ED | `HandlePlayPause` | Known | Event handler |
| 0x007A176D | `HandleSelect` | Known | Event handler |
| 0x007A1ED0 | `HandlePlayPause` | Known | Event handler |
| 0x007A1F45 | `HandleWheelProgress` | Known | Event handler |
| 0x007A1FD5 | `HandlePlayPause` | Known | Event handler |
| 0x007A2055 | `HandleSelectProgress` | Known | Event handler |
| 0x007A27C0 | `HandlePlayPause` | Known | Event handler |
| 0x007A2835 | `HandleWheelProgress` | Known | Event handler |
| 0x007A28C5 | `HandlePlayPause` | Known | Event handler |
| 0x007A2945 | `HandleSelectVolume` | Known | Event handler |
| 0x007A30AE | `HandlePlayPause` | Known | Event handler |
| 0x007A3123 | `HandleWheelVolume` | Known | Event handler |
| 0x007A31B1 | `HandlePlayPause` | Known | Event handler |
| 0x007A3231 | `HandleSelectRating` | Known | Event handler |
| 0x007A399A | `HandlePlayPause` | Known | Event handler |
| 0x007A3A0F | `HandleWheelRating` | Known | Event handler |
| 0x007A3A8F | `HandlePlayPause` | Known | Event handler |
| 0x007A3B06 | `HandleSelectScrub` | Known | Event handler |
| 0x007A4260 | `HandlePlayPause` | Known | Event handler |
| 0x007A42CC | `HandleWheelScrub` | Known | Event handler |
| 0x007A4330 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007A4368 | `HandlePlayPause` | Known | Event handler |
| 0x007A43C2 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007A43F7 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x007A4B67 | `HandlePlayPause` | Known | Event handler |
| 0x007A4BDC | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x007A4C71 | `HandlePlayPause` | Known | Event handler |
| 0x007A4CF1 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007A545D | `HandlePlayPause` | Known | Event handler |
| 0x007A5551 | `HandlePlayPause` | Known | Event handler |
| 0x007A55D1 | `HandleSelectChapterArt` | Known | Event handler |
| 0x007A5D3E | `HandlePlayPause` | Known | Event handler |
| 0x007A5DB3 | `HandleWheelVolume` | Known | Event handler |
| 0x007A5E4A | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007A5EE1 | `HandleSelect` | Known | Event handler |
| 0x007A664D | `HandlePlayPause` | Known | Event handler |
| 0x007A66CB | `HandleWheel` | Known | Event handler |
| 0x007A675E | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007A67F5 | `HandleSelect` | Known | Event handler |
| 0x007A6F61 | `HandlePlayPause` | Known | Event handler |
| 0x007A6FDF | `HandleWheel` | Known | Event handler |
| 0x007A7069 | `HandlePlayPause` | Known | Event handler |
| 0x007A70E9 | `HandleSelect` | Known | Event handler |
| 0x007A784C | `HandlePlayPause` | Known | Event handler |
| 0x007A78C1 | `HandleWheel` | Known | Event handler |
| 0x007A7949 | `HandlePlayPause` | Known | Event handler |
| 0x007A79C9 | `HandleSelectProgress` | Known | Event handler |
| 0x007A8134 | `HandlePlayPause` | Known | Event handler |
| 0x007A81A9 | `HandleWheelProgress` | Known | Event handler |
| 0x007A822B | `HandlePlayPause` | Known | Event handler |
| 0x007A82A2 | `HandleSelectScrub` | Known | Event handler |
| 0x007A89FC | `HandlePlayPause` | Known | Event handler |
| 0x007A8A68 | `HandleWheelScrub` | Known | Event handler |
| 0x007A8AF5 | `HandlePlayPause` | Known | Event handler |
| 0x007A92D7 | `HandlePlayPause` | Known | Event handler |
| 0x007A934C | `HandleWheelVolume` | Known | Event handler |
| 0x007A93DD | `HandlePlayPause` | Known | Event handler |
| 0x007A9BBF | `HandlePlayPause` | Known | Event handler |
| 0x007A9C34 | `HandleWheelBrightness` | Known | Event handler |
| 0x007A9CC9 | `HandlePlayPause` | Known | Event handler |
| 0x007A9D49 | `HandleSelect` | Known | Event handler |
| 0x007AA0AE | `HandlePlayPause` | Known | Event handler |
| 0x007AA191 | `HandlePlayPause` | Known | Event handler |
| 0x007AA211 | `HandleSelectProgress` | Known | Event handler |
| 0x007AA57E | `HandlePlayPause` | Known | Event handler |
| 0x007AA5F3 | `HandleWheelProgress` | Known | Event handler |
| 0x007AA685 | `HandlePlayPause` | Known | Event handler |
| 0x007AA705 | `HandleSelectProgress` | Known | Event handler |
| 0x007AA9FE | `HandlePlayPause` | Known | Event handler |
| 0x007AAA73 | `HandleWheelProgress` | Known | Event handler |
| 0x007AAAEC | `HandlePlayPause` | Known | Event handler |
| 0x007AAB58 | `HandleSelectScrub` | Known | Event handler |
| 0x007AAE35 | `HandlePlayPause` | Known | Event handler |
| 0x007AAE96 | `HandleWheelScrub` | Known | Event handler |
| 0x007AAF25 | `HandlePlayPause` | Known | Event handler |
| 0x007AAFA5 | `HandleSelectVolume` | Known | Event handler |
| 0x007AB29C | `HandlePlayPause` | Known | Event handler |
| 0x007AB311 | `HandleWheelVolume` | Known | Event handler |
| 0x007AB345 | `HandleSelect` | Known | Event handler |
| 0x007AB37D | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AB3B0 | `HandleNotesPop` | Known | Event handler |
| 0x007AB42D | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AB460 | `HandleNotesPop` | Known | Event handler |
| 0x007AB91C | `HandleNotesSelected` | Known | Event handler |
| 0x007AB959 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AB98C | `HandleNotesPop` | Known | Event handler |
| 0x007ABE48 | `HandleNotesSelected` | Known | Event handler |
| 0x007ABE85 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007ABEB8 | `HandleNotesPop` | Known | Event handler |
| 0x007ABEE3 | `HandleNotesSelected` | Known | Event handler |
| 0x007AC3B5 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AC3E8 | `HandleNotesPop` | Known | Event handler |
| 0x007AC413 | `HandleNotesSelected` | Known | Event handler |
| 0x007AC8E5 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AC918 | `HandleNotesPop` | Known | Event handler |
| 0x007AC995 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AC9C8 | `HandleNotesPop` | Known | Event handler |
| 0x007ACA45 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007ACA78 | `HandleNotesPop` | Known | Event handler |
| 0x007ACAF0 | `HandlePlayPause` | Known | Event handler |
| 0x007ACB19 | `HandlePlayPause` | Known | Event handler |
| 0x007ACB47 | `HandlePlayPause` | Known | Event handler |
| 0x007ACB7C | `HandleBrowseAlbum` | Known | Event handler |
| 0x007ACBFC | `HandleHiliteAlbum` | Known | Event handler |
| 0x007ACCA5 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007ACD2C | `HandleHiliteAlbum` | Known | Event handler |
| 0x007ACFF0 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x007AD04C | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x007AD203 | `HandleSelect` | Known | Event handler |
| 0x007AD387 | `HandleSelect` | Known | Event handler |
| 0x007AD3C1 | `HandleImageLast` | Known | Event handler |
| 0x007AD3EB | `HandleImageNext` | Known | Event handler |
| 0x007AD41A | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AD454 | `HandleImageFirst` | Known | Event handler |
| 0x007AD47F | `HandleImagePrev` | Known | Event handler |
| 0x007AD4AB | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AD4DA | `HandleImageNext` | Known | Event handler |
| 0x007AD503 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AD537 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AD566 | `HandleImagePrev` | Known | Event handler |
| 0x007AD587 | `HandleImageWheel` | Known | Event handler |
| 0x007AD625 | `HandleImageNext` | Known | Event handler |
| 0x007AD654 | `HandlePlayPause` | Known | Event handler |
| 0x007AD6A3 | `HandleImagePrev` | Known | Event handler |
| 0x007AD6CF | `HandleSelect` | Known | Event handler |
| 0x007AD99F | `HandleImageNext` | Known | Event handler |
| 0x007AD9C9 | `HandlePause` | Known | Event handler |
| 0x007AD9EE | `HandlePlay` | Known | Event handler |
| 0x007ADA17 | `HandlePlayPause` | Known | Event handler |
| 0x007ADA40 | `HandleImagePrev` | Known | Event handler |
| 0x007ADA99 | `HandleWheel` | Known | Event handler |
| 0x007ADB31 | `HandleImageNext` | Known | Event handler |
| 0x007ADB60 | `HandlePlayPause` | Known | Event handler |
| 0x007ADBAF | `HandleImagePrev` | Known | Event handler |
| 0x007ADBDB | `HandleSelect` | Known | Event handler |
| 0x007ADEAB | `HandleImageNext` | Known | Event handler |
| 0x007ADED5 | `HandlePause` | Known | Event handler |
| 0x007ADEFA | `HandlePlay` | Known | Event handler |
| 0x007ADF23 | `HandlePlayPause` | Known | Event handler |
| 0x007ADF4C | `HandleImagePrev` | Known | Event handler |
| 0x007ADFA5 | `HandleWheel` | Known | Event handler |
| 0x007AE03D | `HandleImageNext` | Known | Event handler |
| 0x007AE06C | `HandlePlayPause` | Known | Event handler |
| 0x007AE0BB | `HandleImagePrev` | Known | Event handler |
| 0x007AE0E7 | `HandleSelect` | Known | Event handler |
| 0x007AE3B7 | `HandleImageNext` | Known | Event handler |
| 0x007AE3E1 | `HandlePause` | Known | Event handler |
| 0x007AE406 | `HandlePlay` | Known | Event handler |
| 0x007AE42F | `HandlePlayPause` | Known | Event handler |
| 0x007AE458 | `HandleImagePrev` | Known | Event handler |
| 0x007AE4B1 | `HandleWheel` | Known | Event handler |
| 0x007AE549 | `HandleImageNext` | Known | Event handler |
| 0x007AE578 | `HandlePlayPause` | Known | Event handler |
| 0x007AE5C7 | `HandleImagePrev` | Known | Event handler |
| 0x007AE5F3 | `HandleSelect` | Known | Event handler |
| 0x007AE8C3 | `HandleImageNext` | Known | Event handler |
| 0x007AE8ED | `HandlePause` | Known | Event handler |
| 0x007AE912 | `HandlePlay` | Known | Event handler |
| 0x007AE93B | `HandlePlayPause` | Known | Event handler |
| 0x007AE964 | `HandleImagePrev` | Known | Event handler |
| 0x007AE9BD | `HandleWheel` | Known | Event handler |
| 0x007AEA55 | `HandleImageNext` | Known | Event handler |
| 0x007AEA84 | `HandlePlayPause` | Known | Event handler |
| 0x007AEAD3 | `HandleImagePrev` | Known | Event handler |
| 0x007AEAFF | `HandleSelect` | Known | Event handler |
| 0x007AEDCF | `HandleImageNext` | Known | Event handler |
| 0x007AEDF9 | `HandlePause` | Known | Event handler |
| 0x007AEE1E | `HandlePlay` | Known | Event handler |
| 0x007AEE47 | `HandlePlayPause` | Known | Event handler |
| 0x007AEE70 | `HandleImagePrev` | Known | Event handler |
| 0x007AEEC9 | `HandleWheel` | Known | Event handler |
| 0x007AEF61 | `HandleImageNext` | Known | Event handler |
| 0x007AEF90 | `HandlePlayPause` | Known | Event handler |
| 0x007AEFDF | `HandleImagePrev` | Known | Event handler |
| 0x007AF00B | `HandleSelect` | Known | Event handler |
| 0x007AF2DB | `HandleImageNext` | Known | Event handler |
| 0x007AF305 | `HandlePause` | Known | Event handler |
| 0x007AF32A | `HandlePlay` | Known | Event handler |
| 0x007AF353 | `HandlePlayPause` | Known | Event handler |
| 0x007AF37C | `HandleImagePrev` | Known | Event handler |
| 0x007AF3D5 | `HandleWheel` | Known | Event handler |
| 0x007AF46D | `HandleImageNext` | Known | Event handler |
| 0x007AF49C | `HandlePlayPause` | Known | Event handler |
| 0x007AF4EB | `HandleImagePrev` | Known | Event handler |
| 0x007AF517 | `HandleSelect` | Known | Event handler |
| 0x007AF762 | `HandleImageNext` | Known | Event handler |
| 0x007AF78C | `HandlePause` | Known | Event handler |
| 0x007AF7B1 | `HandlePlay` | Known | Event handler |
| 0x007AF7DA | `HandlePlayPause` | Known | Event handler |
| 0x007AF803 | `HandleImagePrev` | Known | Event handler |
| 0x007AF86C | `HandleWheel` | Known | Event handler |
| 0x007AF905 | `HandleImageNext` | Known | Event handler |
| 0x007AF934 | `HandlePlayPause` | Known | Event handler |
| 0x007AF983 | `HandleImagePrev` | Known | Event handler |
| 0x007AF9AF | `HandleSelect` | Known | Event handler |
| 0x007AFBFA | `HandleImageNext` | Known | Event handler |
| 0x007AFC24 | `HandlePause` | Known | Event handler |
| 0x007AFC49 | `HandlePlay` | Known | Event handler |
| 0x007AFC72 | `HandlePlayPause` | Known | Event handler |
| 0x007AFC9B | `HandleImagePrev` | Known | Event handler |
| 0x007AFD04 | `HandleWheel` | Known | Event handler |
| 0x007AFD9D | `HandleImageNext` | Known | Event handler |
| 0x007AFDCC | `HandlePlayPause` | Known | Event handler |
| 0x007AFE1B | `HandleImagePrev` | Known | Event handler |
| 0x007AFE47 | `HandleSelect` | Known | Event handler |
| 0x007B0092 | `HandleImageNext` | Known | Event handler |
| 0x007B00BC | `HandlePause` | Known | Event handler |
| 0x007B00E1 | `HandlePlay` | Known | Event handler |
| 0x007B010A | `HandlePlayPause` | Known | Event handler |
| 0x007B0133 | `HandleImagePrev` | Known | Event handler |
| 0x007B019C | `HandleWheel` | Known | Event handler |
| 0x007B0235 | `HandleImageNext` | Known | Event handler |
| 0x007B0264 | `HandlePlayPause` | Known | Event handler |
| 0x007B02B3 | `HandleImagePrev` | Known | Event handler |
| 0x007B02DF | `HandleSelect` | Known | Event handler |
| 0x007B052A | `HandleImageNext` | Known | Event handler |
| 0x007B0554 | `HandlePause` | Known | Event handler |
| 0x007B0579 | `HandlePlay` | Known | Event handler |
| 0x007B05A2 | `HandlePlayPause` | Known | Event handler |
| 0x007B05CB | `HandleImagePrev` | Known | Event handler |
| 0x007B0634 | `HandleWheel` | Known | Event handler |
| 0x007B06CD | `HandleImageNext` | Known | Event handler |
| 0x007B06FC | `HandlePlayPause` | Known | Event handler |
| 0x007B074B | `HandleImagePrev` | Known | Event handler |
| 0x007B0777 | `HandleSelect` | Known | Event handler |
| 0x007B09C2 | `HandleImageNext` | Known | Event handler |
| 0x007B09EC | `HandlePause` | Known | Event handler |
| 0x007B0A11 | `HandlePlay` | Known | Event handler |
| 0x007B0A3A | `HandlePlayPause` | Known | Event handler |
| 0x007B0A63 | `HandleImagePrev` | Known | Event handler |
| 0x007B0ACC | `HandleWheel` | Known | Event handler |
| 0x007B0AF9 | `HandleSelect` | Known | Event handler |
| 0x007B0B29 | `HandleSelect` | Known | Event handler |
| 0x007B0C39 | `HandleTuning` | Known | Event handler |
| 0x007B0DF5 | `HandleVolumeChange` | Known | Event handler |
| 0x007B0F2E | `HandleVolumeWheel` | Known | Event handler |
| 0x007B1098 | `HandleTimerDone` | Known | Event handler |
| 0x007B12F5 | `HandleFrequencyChange` | Known | Event handler |
| 0x007B1464 | `HandleTimerDone` | Known | Event handler |
| 0x007B16C1 | `HandleFrequencyChange` | Known | Event handler |
| 0x007B17D9 | `HandleTimerDone` | Known | Event handler |
| 0x007B1908 | `HandleVolumeChange` | Known | Event handler |
| 0x007B1A0C | `HandleVolumeWheel` | Known | Event handler |
| 0x007B1C64 | `HandleRemoteUIRemotePlayPause` | Known | Event handler |
| 0x007B1ED8 | `HandleRemoteUIRemotePlayPause` | Known | Event handler |
| 0x007B1F27 | `HandleExitUnsupported` | Known | Event handler |
| 0x007B1F59 | `HandleExitUnsupported` | Known | Event handler |
| 0x007B52A9 | `HandleSelectKey` | Known | Event handler |
| 0x007B52DE | `HandleWheel` | Known | Event handler |
| 0x007B542C | `HandleLanguageAfterReset` | Known | Event handler |
| 0x007B547F | `HandleSelectKey` | Known | Event handler |
| 0x007B54A7 | `HandleSelectKey` | Known | Event handler |
| 0x007B54D7 | `HandleExit` | Known | Event handler |
| 0x007B5501 | `HandleStartStop` | Known | Event handler |
| 0x007B5567 | `HandleStartStop` | Known | Event handler |
| 0x007B567F | `HandleExit` | Known | Event handler |
| 0x007B56A9 | `HandleStartStop` | Known | Event handler |
| 0x007B56D5 | `HandleLap` | Known | Event handler |
| 0x007B57D9 | `HandleSelectLozinch` | Known | Event handler |
| 0x007B627C | `HandleSelect` | Known | Event handler |
| 0x007B6CCB | `HandleChoosePowerPlay` | Known | Event handler |
| 0x007B6D06 | `HandleChooseVoicePrompts` | Known | Event handler |
| 0x007B6D44 | `HandleChooseUnit` | Known | Event handler |
| 0x007B6ED8 | `HandleListChoose` | Known | Event handler |
| 0x007B7137 | `HandleSelect` | Known | Event handler |
| 0x007B7357 | `HandleSelect` | Known | Event handler |
| 0x007B738D | `HandleMenuKeyNop` | Known | Event handler |
| 0x007B75BE | `HandleNowPlayingSelected` | Known | Event handler |
| 0x007B75FC | `HandleChoosePlaylistsPlay` | Known | Event handler |
| 0x007B763B | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x007B767B | `HandleNoneSelected` | Known | Event handler |
| 0x007B76B1 | `HandleBegin` | Known | Event handler |
| 0x007B7922 | `HandleBegin` | Known | Event handler |
| 0x007B7951 | `HandleBegin` | Known | Event handler |
| 0x007B7A0D | `HandleBegin` | Known | Event handler |
| 0x007B7A39 | `HandleBegin` | Known | Event handler |
| 0x007B7AF5 | `HandleBegin` | Known | Event handler |
| 0x007B7B21 | `HandleBegin` | Known | Event handler |
| 0x007B7BDD | `HandleBegin` | Known | Event handler |
| 0x007B7C11 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007B7C3C | `HandleMenuKey` | Known | Event handler |
| 0x007B7CD3 | `HandlePauseHold` | Known | Event handler |
| 0x007B7D02 | `HandlePauseKey` | Known | Event handler |
| 0x007B7D8C | `HandleSelectKeyDown` | Known | Event handler |
| 0x007B7DC6 | `HandlePowerPlay` | Known | Event handler |
| 0x007B7DF2 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B825F | `HandlePauseHold` | Known | Event handler |
| 0x007B828E | `HandlePauseKey` | Known | Event handler |
| 0x007B82B9 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B82F7 | `HandlePowerPlay` | Known | Event handler |
| 0x007B8326 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B834C | `HandleWheel` | Known | Event handler |
| 0x007B8381 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007B83AC | `HandleMenuKey` | Known | Event handler |
| 0x007B8443 | `HandlePauseHold` | Known | Event handler |
| 0x007B8472 | `HandlePauseKey` | Known | Event handler |
| 0x007B84FC | `HandleSelectKeyDown` | Known | Event handler |
| 0x007B852C | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B898C | `HandlePauseHold` | Known | Event handler |
| 0x007B89BB | `HandlePauseKey` | Known | Event handler |
| 0x007B89E6 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B8A1A | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B8A40 | `HandleWheel` | Known | Event handler |
| 0x007B8A75 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007B8AA0 | `HandleMenuKey` | Known | Event handler |
| 0x007B8B37 | `HandlePauseHold` | Known | Event handler |
| 0x007B8B66 | `HandlePauseKey` | Known | Event handler |
| 0x007B8BF0 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007B8C2A | `HandlePowerPlay` | Known | Event handler |
| 0x007B8C56 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B90C2 | `HandlePauseHold` | Known | Event handler |
| 0x007B90F1 | `HandlePauseKey` | Known | Event handler |
| 0x007B911C | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B915A | `HandlePowerPlay` | Known | Event handler |
| 0x007B9189 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B91AF | `HandleWheel` | Known | Event handler |
| 0x007B91E5 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007B9210 | `HandleMenuKey` | Known | Event handler |
| 0x007B92A7 | `HandlePauseHold` | Known | Event handler |
| 0x007B92D6 | `HandlePauseKey` | Known | Event handler |
| 0x007B9360 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007B9390 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B97EF | `HandlePauseHold` | Known | Event handler |
| 0x007B981E | `HandlePauseKey` | Known | Event handler |
| 0x007B9849 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B987D | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B98A3 | `HandleWheel` | Known | Event handler |
| 0x007B98D9 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007B9904 | `HandleMenuKey` | Known | Event handler |
| 0x007B999B | `HandlePauseHold` | Known | Event handler |
| 0x007B99CA | `HandlePauseKey` | Known | Event handler |
| 0x007B9A54 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007B9A8E | `HandlePowerPlay` | Known | Event handler |
| 0x007B9ABA | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B9F2A | `HandlePauseHold` | Known | Event handler |
| 0x007B9F59 | `HandlePauseKey` | Known | Event handler |
| 0x007B9F84 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007B9FC2 | `HandlePowerPlay` | Known | Event handler |
| 0x007B9FF1 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BA017 | `HandleWheel` | Known | Event handler |
| 0x007BA04D | `HandleMenuKeyNop` | Known | Event handler |
| 0x007BA078 | `HandleMenuKey` | Known | Event handler |
| 0x007BA10F | `HandlePauseHold` | Known | Event handler |
| 0x007BA13E | `HandlePauseKey` | Known | Event handler |
| 0x007BA1C8 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007BA1F8 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BA65B | `HandlePauseHold` | Known | Event handler |
| 0x007BA68A | `HandlePauseKey` | Known | Event handler |
| 0x007BA6B5 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BA6E9 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BA70F | `HandleWheel` | Known | Event handler |
| 0x007BA745 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007BA770 | `HandleMenuKey` | Known | Event handler |
| 0x007BA807 | `HandlePauseHold` | Known | Event handler |
| 0x007BA836 | `HandlePauseKey` | Known | Event handler |
| 0x007BA8C0 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007BA8FA | `HandlePowerPlay` | Known | Event handler |
| 0x007BA926 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BAD96 | `HandlePauseHold` | Known | Event handler |
| 0x007BADC5 | `HandlePauseKey` | Known | Event handler |
| 0x007BADF0 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BAE2E | `HandlePowerPlay` | Known | Event handler |
| 0x007BAE5D | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BAE83 | `HandleWheel` | Known | Event handler |
| 0x007BAEB9 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007BAEE4 | `HandleMenuKey` | Known | Event handler |
| 0x007BAF7B | `HandlePauseHold` | Known | Event handler |
| 0x007BAFAA | `HandlePauseKey` | Known | Event handler |
| 0x007BB034 | `HandleSelectKeyDown` | Known | Event handler |
| 0x007BB064 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BB4C7 | `HandlePauseHold` | Known | Event handler |
| 0x007BB4F6 | `HandlePauseKey` | Known | Event handler |
| 0x007BB521 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BB555 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BB57B | `HandleWheel` | Known | Event handler |
| 0x007BB5B1 | `HandleMenuKeyNop` | Known | Event handler |
| 0x007BB5DC | `HandleMenuKey` | Known | Event handler |
| 0x007BB673 | `HandlePauseHold` | Known | Event handler |
| 0x007BB6A2 | `HandlePauseKey` | Known | Event handler |
| 0x007BB72C | `HandleSelectKeyDown` | Known | Event handler |
| 0x007BB75C | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BBB55 | `HandlePauseHold` | Known | Event handler |
| 0x007BBB84 | `HandlePauseKey` | Known | Event handler |
| 0x007BBBAF | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BBBE3 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x007BBC09 | `HandleWheel` | Known | Event handler |
| 0x007BBC3D | `HandleMenuKeyNop` | Known | Event handler |
| 0x007BBC68 | `HandleResumeWorkout` | Known | Event handler |
| 0x007BBD43 | `HandleResumeWorkout` | Known | Event handler |
| 0x007BBDB7 | `HandlePauseWorkout` | Known | Event handler |
| 0x007BBE25 | `HandleChooseMusic` | Known | Event handler |
| 0x007BBEC2 | `HandleEndWorkout` | Known | Event handler |
| 0x007BBF6D | `HandleMenuKeyNop` | Known | Event handler |
| 0x007BC214 | `HandleEndWorkout` | Known | Event handler |
| 0x007BC6A3 | `HandleSelectResume` | Known | Event handler |
| 0x007BC6DB | `HandleEndWorkout` | Known | Event handler |
| 0x007BC786 | `HandleChooseCustomTimedWorkout` | Known | Event handler |
| 0x007BC81F | `HandleChooseTimedWorkoutFromList` | Known | Event handler |
| 0x007BC8D2 | `HandleChooseCustomDistanceWorkout` | Known | Event handler |
| 0x007BC972 | `HandleChooseDistanceWorkoutFromList` | Known | Event handler |
| 0x007BCB58 | `HandleChooseCustomCalorieWorkout` | Known | Event handler |
| 0x007BCBF7 | `HandleChooseCalorieWorkoutFromList` | Known | Event handler |
| 0x007BCF66 | `HandleChooseLink` | Known | Event handler |
| 0x007BCF9C | `HandleChooseCalibrate` | Known | Event handler |
| 0x007BD2F5 | `HandleChooseRemoteSetting` | Known | Event handler |
| 0x007BD334 | `HandleChooseRemoteLink` | Known | Event handler |
| 0x007BD370 | `HandleChooseRemoteUnlink` | Known | Event handler |
| 0x007BD71E | `Handle400MetersWalk` | Known | Event handler |
| 0x007BD757 | `HandleCustomWalk` | Known | Event handler |
| 0x007BD82D | `HandleSelectWalking` | Known | Event handler |
| 0x007BD951 | `HandleSelectRunning` | Known | Event handler |
| 0x007BDC9E | `Handle400MetersRun` | Known | Event handler |
| 0x007BDCD6 | `HandleCustomRun` | Known | Event handler |
| 0x007BDF21 | `HandleSelect` | Known | Event handler |
| 0x007BDF51 | `HandleSelect` | Known | Event handler |
| 0x007BE0C7 | `HandleLinkNewRemote` | Known | Event handler |
| 0x007BE235 | `HandleSelect` | Known | Event handler |
| 0x007BE265 | `HandleSelect` | Known | Event handler |
| 0x007BE759 | `HandleUnlinkRemote` | Known | Event handler |
| 0x007BE9BD | `HandleWeightSelect` | Known | Event handler |
| 0x007BEA1A | `HandleWeightWheel` | Known | Event handler |
| 0x007BEA4D | `HandleWeightSelect` | Known | Event handler |
| 0x007BEAD7 | `HandleWeightWheel` | Known | Event handler |
| 0x007BEB09 | `HandleDistanceSelect` | Known | Event handler |
| 0x007BEB95 | `HandleDistanceWheel` | Known | Event handler |
| 0x007BEBC9 | `HandleDistanceSelect` | Known | Event handler |
| 0x007BEC55 | `HandleDistanceWheel` | Known | Event handler |
| 0x007BEC89 | `HandleTimeSelect` | Known | Event handler |
| 0x007BED11 | `HandleTimeWheel` | Known | Event handler |
| 0x007BED41 | `HandleCaloriesSelect` | Known | Event handler |
| 0x007BEE99 | `HandleCaloriesWheel` | Known | Event handler |
| 0x007BF205 | `HandleChooseLast` | Known | Event handler |
| 0x007BF23B | `HandleChooseRecent` | Known | Event handler |
| 0x007BF273 | `HandleChooseBest` | Known | Event handler |
| 0x007BF589 | `HandleSelect` | Known | Event handler |
| 0x007BF771 | `HandleSelectFromLastWorkout` | Known | Event handler |
| 0x007BF969 | `HandleSelect` | Known | Event handler |
| 0x007BFC22 | `HandleChooseRecentWorkout` | Known | Event handler |
| 0x007BFCF5 | `HandleSelect` | Known | Event handler |
| 0x007BFD89 | `HandleSelect_Basic` | Known | Event handler |
| 0x007C006D | `HandleSelect_Dynamic` | Known | Event handler |
| 0x007C0361 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x007C0651 | `HandleSelect_Dynamic` | Known | Event handler |
| 0x007C094C | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x007C0C36 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x007C0D17 | `HandlePlayPause` | Known | Event handler |
| 0x007C0DA5 | `HandlePlayPause` | Known | Event handler |
| 0x007C0E35 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x007C0E6D | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x007C0EA9 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x007C0EEC | `HandlePlayPause` | Known | Event handler |
| 0x007C0F22 | `HandleAddToOTG` | Known | Event handler |
| 0x007C1177 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007C13D3 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007E22EE | `HandleSelectClock` | Known | Event handler |
| 0x007E2327 | `HandleHilited` | Known | Event handler |
| 0x007E2359 | `HandleWheel` | Known | Event handler |
| 0x007E23A0 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x007E2425 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x007E25E5 | `HandleImageLast` | Known | Event handler |
| 0x007E260F | `HandleScreenNext` | Known | Event handler |
| 0x007E263F | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007E2679 | `HandleImageFirst` | Known | Event handler |
| 0x007E26A4 | `HandleScreenPrev` | Known | Event handler |
| 0x007E26D1 | `HandleBrowseLarge` | Known | Event handler |
| 0x007E2751 | `HandleImageNext` | Known | Event handler |
| 0x007E277A | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007E27AE | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007E27DD | `HandleImagePrev` | Known | Event handler |
| 0x007E280B | `HandleBrowseSmall` | Known | Event handler |

---

## 6. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00103AB4 | `GotoNowPlaying` | Known | Navigation |
| 0x00103B2C | `GotoMainMenu` | Known | Navigation |
| 0x001214B8 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x001214D0 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x00121648 | `GotoScreen_AddressBook` | Known | Navigation |
| 0x0012B6E0 | `GotoNowPlaying` | Known | Navigation |
| 0x0012B6F4 | `GotoAlbums` | Known | Navigation |
| 0x0012B700 | `GotoSongs` | Known | Navigation |
| 0x00138FE0 | `GotoScreen_EnterPassKey` | Known | Navigation |
| 0x00138FF8 | `GotoScreen_LockediPod` | Known | Navigation |
| 0x0013999C | `GotoScreen_MainMenu` | Known | Navigation |
| 0x00152814 | `GotoMainMenu` | Known | Navigation |
| 0x001D74B0 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001E2618 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001E2E68 | `GotoPlayDeleteMenu` | Known | Navigation |
| 0x001E2EEC | `GotoNowPlaying` | Known | Navigation |
| 0x002012D0 | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation |
| 0x0020E92C | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x0020E9E4 | `GotoScreen_SettingsMenu` | Known | Navigation |
| 0x00217724 | `GotoDefaultLayout` | Known | Navigation |
| 0x002177A8 | `GotoVolumeLayout` | Known | Navigation |
| 0x002178F0 | `GotoProgressLayout` | Known | Navigation |
| 0x00217C00 | `GotoDefault` | Known | Navigation |
| 0x00217E14 | `GotoPausedIconLayout` | Known | Navigation |
| 0x00217F88 | `GotoProgressLayout` | Known | Navigation |
| 0x002180DC | `GotoDefaultLayout` | Known | Navigation |
| 0x0021819C | `GotoDefaultLayout` | Known | Navigation |
| 0x00218270 | `GotoProgressLayout` | Known | Navigation |
| 0x002183FC | `GotoProgressLayout` | Known | Navigation |
| 0x00219C34 | `GotoNowPlaying` | Known | Navigation |
| 0x0021A4EC | `GotoNowPlaying` | Known | Navigation |
| 0x0021DDA8 | `GotoScreen_Language` | Known | Navigation |
| 0x0021E120 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x0021E13C | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x0021E154 | `GotoDefaultLayout` | Known | Navigation |
| 0x0021E1E4 | `GotoVolumeLayout` | Known | Navigation |
| 0x0021E1F8 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x0021E298 | `GotoProgressLayout` | Known | Navigation |
| 0x0021E2AC | `GotoProgressVideoLayout` | Known | Navigation |
| 0x0021E774 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x0021E9E0 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x0021EB74 | `GotoProgressLayout` | Known | Navigation |
| 0x0021EB88 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x0021EC4C | `GotoBrightnessVideoLayout` | Known | Navigation |
| 0x0021EC68 | `GotoRatingLayout` | Known | Navigation |
| 0x0021EF14 | `GotoChapterArtLayout` | Known | Navigation |
| 0x0021EF2C | `GotoExtraInfoLayout` | Known | Navigation |
| 0x0021EF40 | `GotoShuffleLayout` | Known | Navigation |
| 0x0021F2A4 | `GotoVolumeLayout` | Known | Navigation |
| 0x0021F2BC | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x0021F348 | `GotoVolumeLayout` | Known | Navigation |
| 0x0021F35C | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x0021F574 | `GotoScrubLayout` | Known | Navigation |
| 0x0021F584 | `GotoScrubVideoLayout` | Known | Navigation |
| 0x0021F614 | `GotoProgressLayout` | Known | Navigation |
| 0x0021F628 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x0021F78C | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x0021F7A8 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x0021F7C0 | `GotoDefaultLayout` | Known | Navigation |
| 0x0021F8A4 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x0021FA3C | `GotoChapterArtLayout` | Known | Navigation |
| 0x0021FA54 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x0021FB44 | `GotoProgressLayout` | Known | Navigation |
| 0x0021FBD0 | `GotoProgressLayout` | Known | Navigation |
| 0x0021FBE4 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x0021FE40 | `GotoStatusBarLayout` | Known | Navigation |
| 0x0021FE54 | `GotoDefaultLayout` | Known | Navigation |
| 0x0022002C | `GotoDefault` | Known | Navigation |
| 0x00220160 | `GotoProgressLayout` | Known | Navigation |
| 0x00220320 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x00220470 | `GotoBrightnessLayout` | Known | Navigation |
| 0x002204F4 | `GotoBrightnessLayout` | Known | Navigation |
| 0x00220574 | `GotoVolumeLayout` | Known | Navigation |
| 0x002205C0 | `GotoScrubLayout` | Known | Navigation |
| 0x00220694 | `GotoStatusBarLayout` | Known | Navigation |
| 0x002206A8 | `GotoDefaultLayout` | Known | Navigation |
| 0x00220780 | `GotoScrubLayout` | Known | Navigation |
| 0x002207D0 | `GotoScrubLayout` | Known | Navigation |
| 0x00227158 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x00227360 | `GotoScreen_BacklightTimer` | Known | Navigation |
| 0x002273F0 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x00227408 | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation |
| 0x0022BFAC | `GotoScreen_LockDialog` | Known | Navigation |
| 0x0022BFC4 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x0022D9CC | `GotoRadio` | Known | Navigation |
| 0x0022E414 | `GotoNowPlaying` | Known | Navigation |
| 0x0022EB6C | `GotoNowPlaying` | Known | Navigation |
| 0x0022F1EC | `GotoFirstBoot` | Known | Navigation |
| 0x0022F1FC | `GotoNotesApp` | Known | Navigation |
| 0x0022F210 | `GotoLockApp` | Known | Navigation |
| 0x002349A4 | `GotoNowPlaying` | Known | Navigation |
| 0x003BBF14 | `GotoProgressLayout` | Known | Navigation |
| 0x0073A535 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x007A8B75 | `GotoDefault` | Known | Navigation |
| 0x007A945D | `GotoDefault` | Known | Navigation |
| 0x0089D10C | `Gotowe?` | Known | Navigation |

---

## 7. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00169440 | `CoverFlow_Screen` | Known | Screen layout |
| 0x00193168 | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x00193188 | `VoiceMemos_Menu_Screen_Recording` | Known | Screen layout |
| 0x001931AC | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x0072E65A | `Clock_Screen` | Known | Screen layout |
| 0x0072E66A | `Clock_Screen_Default"` | Known | Screen layout |
| 0x0072E6CF | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x0072E72D | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0072E745 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x0072E7B2 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x0072E850 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x0072E8AF | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0072E8C5 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x0072E930 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x0072E98A | `Games_Menu_Screen` | Known | Screen layout |
| 0x0072E99F | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x0072EA09 | `Extras_Screen_Games` | Known | Screen layout |
| 0x0072EAC8 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x0072EB8C | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0072EC55 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x0072ED90 | `Clock_Africa_City_Screen ` | Known | Screen layout |
| 0x0072EDAC | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x0072EE30 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x0072EE4A | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x0072EECC | `Clock_Atlantic_City_Screen"` | Known | Screen layout |
| 0x0072EEEA | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x0072EF70 | `Clock_Australia_City_Screen#` | Known | Screen layout |
| 0x0072EF8F | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x0072F016 | `Clock_Europe_City_Screen ` | Known | Screen layout |
| 0x0072F032 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x0072F0B6 | `Clock_NorthAmerica_City_Screen&` | Known | Screen layout |
| 0x0072F0D8 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0072F162 | `Clock_Pacific_City_Screen!` | Known | Screen layout |
| 0x0072F17F | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x0072F204 | `Clock_SouthAmerica_City_Screen&` | Known | Screen layout |
| 0x0072F226 | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0072F2B3 | `Clock_Screen"` | Known | Screen layout |
| 0x0072F358 | `Clock_Screen"` | Known | Screen layout |
| 0x0072F3FD | `Clock_Screen"` | Known | Screen layout |
| 0x0072F4A2 | `Clock_Screen"` | Known | Screen layout |
| 0x0072F547 | `Clock_Screen"` | Known | Screen layout |
| 0x0072F5EC | `Clock_Screen"` | Known | Screen layout |
| 0x0072F691 | `Clock_Screen"` | Known | Screen layout |
| 0x0072F736 | `Clock_Screen"` | Known | Screen layout |
| 0x0072F7DB | `Clock_Screen"` | Known | Screen layout |
| 0x0072F880 | `Clock_Screen"` | Known | Screen layout |
| 0x0072F925 | `Clock_Screen"` | Known | Screen layout |
| 0x0072F9CA | `Clock_Screen"` | Known | Screen layout |
| 0x0072FA6F | `Clock_Screen"` | Known | Screen layout |
| 0x0072FB14 | `Clock_Screen"` | Known | Screen layout |
| 0x0072FBB9 | `Clock_Screen"` | Known | Screen layout |
| 0x0072FC5E | `Clock_Screen"` | Known | Screen layout |
| 0x0072FD03 | `Clock_Screen"` | Known | Screen layout |
| 0x0072FDA8 | `Clock_Screen"` | Known | Screen layout |
| 0x0072FE4D | `Clock_Screen"` | Known | Screen layout |
| 0x0072FEF2 | `Clock_Screen"` | Known | Screen layout |
| 0x0072FF97 | `Clock_Screen"` | Known | Screen layout |
| 0x0073003C | `Clock_Screen"` | Known | Screen layout |
| 0x007300E1 | `Clock_Screen"` | Known | Screen layout |
| 0x00730186 | `Clock_Screen"` | Known | Screen layout |
| 0x0073022B | `Clock_Screen"` | Known | Screen layout |
| 0x007302D0 | `Clock_Screen"` | Known | Screen layout |
| 0x00730375 | `Clock_Screen"` | Known | Screen layout |
| 0x0073041A | `Clock_Screen"` | Known | Screen layout |
| 0x007304BF | `Clock_Screen"` | Known | Screen layout |
| 0x00730564 | `Clock_Screen"` | Known | Screen layout |
| 0x00730609 | `Clock_Screen"` | Known | Screen layout |
| 0x007306B3 | `Clock_Screen"` | Known | Screen layout |
| 0x00730758 | `Clock_Screen"` | Known | Screen layout |
| 0x007307FD | `Clock_Screen"` | Known | Screen layout |
| 0x007308A2 | `Clock_Screen"` | Known | Screen layout |
| 0x00730947 | `Clock_Screen"` | Known | Screen layout |
| 0x007309EC | `Clock_Screen"` | Known | Screen layout |
| 0x00730A91 | `Clock_Screen"` | Known | Screen layout |
| 0x00730B36 | `Clock_Screen"` | Known | Screen layout |
| 0x00730BDB | `Clock_Screen"` | Known | Screen layout |
| 0x00730C80 | `Clock_Screen"` | Known | Screen layout |
| 0x00730D25 | `Clock_Screen"` | Known | Screen layout |
| 0x00730DCA | `Clock_Screen"` | Known | Screen layout |
| 0x00730E6F | `Clock_Screen"` | Known | Screen layout |
| 0x00730F14 | `Clock_Screen"` | Known | Screen layout |
| 0x00730FB9 | `Clock_Screen"` | Known | Screen layout |
| 0x0073105E | `Clock_Screen"` | Known | Screen layout |
| 0x00731103 | `Clock_Screen"` | Known | Screen layout |
| 0x007311A8 | `Clock_Screen"` | Known | Screen layout |
| 0x0073124D | `Clock_Screen"` | Known | Screen layout |
| 0x007312F2 | `Clock_Screen"` | Known | Screen layout |
| 0x00731397 | `Clock_Screen"` | Known | Screen layout |
| 0x0073143C | `Clock_Screen"` | Known | Screen layout |
| 0x007314E1 | `Clock_Screen"` | Known | Screen layout |
| 0x00731586 | `Clock_Screen"` | Known | Screen layout |
| 0x0073162B | `Clock_Screen"` | Known | Screen layout |
| 0x007316D0 | `Clock_Screen"` | Known | Screen layout |
| 0x00731775 | `Clock_Screen"` | Known | Screen layout |
| 0x0073181A | `Clock_Screen"` | Known | Screen layout |
| 0x007318BF | `Clock_Screen"` | Known | Screen layout |
| 0x00731964 | `Clock_Screen"` | Known | Screen layout |
| 0x00731A09 | `Clock_Screen"` | Known | Screen layout |
| 0x00731AAE | `Clock_Screen"` | Known | Screen layout |
| 0x00731B53 | `Clock_Screen"` | Known | Screen layout |
| 0x00731BF8 | `Clock_Screen"` | Known | Screen layout |
| 0x00731C9D | `Clock_Screen"` | Known | Screen layout |
| 0x00731D42 | `Clock_Screen"` | Known | Screen layout |
| 0x00731DE7 | `Clock_Screen"` | Known | Screen layout |
| 0x00731E8C | `Clock_Screen"` | Known | Screen layout |
| 0x00731F31 | `Clock_Screen"` | Known | Screen layout |
| 0x00731FD6 | `Clock_Screen"` | Known | Screen layout |
| 0x0073207B | `Clock_Screen"` | Known | Screen layout |
| 0x00732120 | `Clock_Screen"` | Known | Screen layout |
| 0x007321C5 | `Clock_Screen"` | Known | Screen layout |
| 0x0073226A | `Clock_Screen"` | Known | Screen layout |
| 0x0073230F | `Clock_Screen"` | Known | Screen layout |
| 0x007323B4 | `Clock_Screen"` | Known | Screen layout |
| 0x00732459 | `Clock_Screen"` | Known | Screen layout |
| 0x007324FE | `Clock_Screen"` | Known | Screen layout |
| 0x007325A3 | `Clock_Screen"` | Known | Screen layout |
| 0x00732648 | `Clock_Screen"` | Known | Screen layout |
| 0x007326ED | `Clock_Screen"` | Known | Screen layout |
| 0x00732792 | `Clock_Screen"` | Known | Screen layout |
| 0x00732837 | `Clock_Screen"` | Known | Screen layout |
| 0x007328DC | `Clock_Screen"` | Known | Screen layout |
| 0x00732981 | `Clock_Screen"` | Known | Screen layout |
| 0x00732A26 | `Clock_Screen"` | Known | Screen layout |
| 0x00732ACB | `Clock_Screen"` | Known | Screen layout |
| 0x00732B77 | `Clock_Screen"` | Known | Screen layout |
| 0x00732C1C | `Clock_Screen"` | Known | Screen layout |
| 0x00732CC1 | `Clock_Screen"` | Known | Screen layout |
| 0x00732D6B | `Clock_Screen"` | Known | Screen layout |
| 0x00732E10 | `Clock_Screen"` | Known | Screen layout |
| 0x00732EB5 | `Clock_Screen"` | Known | Screen layout |
| 0x00732F5A | `Clock_Screen"` | Known | Screen layout |
| 0x00732FFF | `Clock_Screen"` | Known | Screen layout |
| 0x007330A4 | `Clock_Screen"` | Known | Screen layout |
| 0x00733149 | `Clock_Screen"` | Known | Screen layout |
| 0x007331EE | `Clock_Screen"` | Known | Screen layout |
| 0x00733297 | `Clock_Screen"` | Known | Screen layout |
| 0x0073333C | `Clock_Screen"` | Known | Screen layout |
| 0x007333E1 | `Clock_Screen"` | Known | Screen layout |
| 0x00733486 | `Clock_Screen"` | Known | Screen layout |
| 0x0073352B | `Clock_Screen"` | Known | Screen layout |
| 0x007335D0 | `Clock_Screen"` | Known | Screen layout |
| 0x00733675 | `Clock_Screen"` | Known | Screen layout |
| 0x0073371A | `Clock_Screen"` | Known | Screen layout |
| 0x007337BF | `Clock_Screen"` | Known | Screen layout |
| 0x00733864 | `Clock_Screen"` | Known | Screen layout |
| 0x00733909 | `Clock_Screen"` | Known | Screen layout |
| 0x007339AE | `Clock_Screen"` | Known | Screen layout |
| 0x00733A53 | `Clock_Screen"` | Known | Screen layout |
| 0x00733AF8 | `Clock_Screen"` | Known | Screen layout |
| 0x00733B9D | `Clock_Screen"` | Known | Screen layout |
| 0x00733C42 | `Clock_Screen"` | Known | Screen layout |
| 0x00733CE7 | `Clock_Screen"` | Known | Screen layout |
| 0x00733D8C | `Clock_Screen"` | Known | Screen layout |
| 0x00733E31 | `Clock_Screen"` | Known | Screen layout |
| 0x00733ED6 | `Clock_Screen"` | Known | Screen layout |
| 0x00733F7B | `Clock_Screen"` | Known | Screen layout |
| 0x00734020 | `Clock_Screen"` | Known | Screen layout |
| 0x007340C5 | `Clock_Screen"` | Known | Screen layout |
| 0x0073416A | `Clock_Screen"` | Known | Screen layout |
| 0x0073420F | `Clock_Screen"` | Known | Screen layout |
| 0x007342B4 | `Clock_Screen"` | Known | Screen layout |
| 0x00734359 | `Clock_Screen"` | Known | Screen layout |
| 0x007343FE | `Clock_Screen"` | Known | Screen layout |
| 0x007344A3 | `Clock_Screen"` | Known | Screen layout |
| 0x00734548 | `Clock_Screen"` | Known | Screen layout |
| 0x007345ED | `Clock_Screen"` | Known | Screen layout |
| 0x00734692 | `Clock_Screen"` | Known | Screen layout |
| 0x00734737 | `Clock_Screen"` | Known | Screen layout |
| 0x007347DC | `Clock_Screen"` | Known | Screen layout |
| 0x00734887 | `Clock_Screen"` | Known | Screen layout |
| 0x0073492C | `Clock_Screen"` | Known | Screen layout |
| 0x007349D1 | `Clock_Screen"` | Known | Screen layout |
| 0x00734A76 | `Clock_Screen"` | Known | Screen layout |
| 0x00734B1B | `Clock_Screen"` | Known | Screen layout |
| 0x00734BC0 | `Clock_Screen"` | Known | Screen layout |
| 0x00734C65 | `Clock_Screen"` | Known | Screen layout |
| 0x00734D0A | `Clock_Screen"` | Known | Screen layout |
| 0x00734DAF | `Clock_Screen"` | Known | Screen layout |
| 0x00734E54 | `Clock_Screen"` | Known | Screen layout |
| 0x00734EF9 | `Clock_Screen"` | Known | Screen layout |
| 0x00734F9E | `Clock_Screen"` | Known | Screen layout |
| 0x00735043 | `Clock_Screen"` | Known | Screen layout |
| 0x007350E8 | `Clock_Screen"` | Known | Screen layout |
| 0x0073518D | `Clock_Screen"` | Known | Screen layout |
| 0x00735232 | `Clock_Screen"` | Known | Screen layout |
| 0x007352D7 | `Clock_Screen"` | Known | Screen layout |
| 0x0073537C | `Clock_Screen"` | Known | Screen layout |
| 0x00735421 | `Clock_Screen"` | Known | Screen layout |
| 0x007354C6 | `Clock_Screen"` | Known | Screen layout |
| 0x0073556B | `Clock_Screen"` | Known | Screen layout |
| 0x00735610 | `Clock_Screen"` | Known | Screen layout |
| 0x007356B5 | `Clock_Screen"` | Known | Screen layout |
| 0x0073575A | `Clock_Screen"` | Known | Screen layout |
| 0x007357FF | `Clock_Screen"` | Known | Screen layout |
| 0x007358A4 | `Clock_Screen"` | Known | Screen layout |
| 0x00735949 | `Clock_Screen"` | Known | Screen layout |
| 0x007359EE | `Clock_Screen"` | Known | Screen layout |
| 0x00735A93 | `Clock_Screen"` | Known | Screen layout |
| 0x00735B38 | `Clock_Screen"` | Known | Screen layout |
| 0x00735BDD | `Clock_Screen"` | Known | Screen layout |
| 0x00735C82 | `Clock_Screen"` | Known | Screen layout |
| 0x00735D27 | `Clock_Screen"` | Known | Screen layout |
| 0x00735DCC | `Clock_Screen"` | Known | Screen layout |
| 0x00735E71 | `Clock_Screen"` | Known | Screen layout |
| 0x00735F16 | `Clock_Screen"` | Known | Screen layout |
| 0x00735FBB | `Clock_Screen"` | Known | Screen layout |
| 0x00736060 | `Clock_Screen"` | Known | Screen layout |
| 0x00736105 | `Clock_Screen"` | Known | Screen layout |
| 0x007361AA | `Clock_Screen"` | Known | Screen layout |
| 0x0073624F | `Clock_Screen"` | Known | Screen layout |
| 0x007362F4 | `Clock_Screen"` | Known | Screen layout |
| 0x00736399 | `Clock_Screen"` | Known | Screen layout |
| 0x0073643E | `Clock_Screen"` | Known | Screen layout |
| 0x007364E3 | `Clock_Screen"` | Known | Screen layout |
| 0x00736588 | `Clock_Screen"` | Known | Screen layout |
| 0x0073662D | `Clock_Screen"` | Known | Screen layout |
| 0x007366D2 | `Clock_Screen"` | Known | Screen layout |
| 0x00736777 | `Clock_Screen"` | Known | Screen layout |
| 0x0073681C | `Clock_Screen"` | Known | Screen layout |
| 0x007368C7 | `Clock_Screen"` | Known | Screen layout |
| 0x0073696C | `Clock_Screen"` | Known | Screen layout |
| 0x00736A11 | `Clock_Screen"` | Known | Screen layout |
| 0x00736AB6 | `Clock_Screen"` | Known | Screen layout |
| 0x00736B5B | `Clock_Screen"` | Known | Screen layout |
| 0x00736C07 | `Clock_Screen"` | Known | Screen layout |
| 0x00736CAC | `Clock_Screen"` | Known | Screen layout |
| 0x00736D51 | `Clock_Screen"` | Known | Screen layout |
| 0x00736DF6 | `Clock_Screen"` | Known | Screen layout |
| 0x00736E9B | `Clock_Screen"` | Known | Screen layout |
| 0x00736F40 | `Clock_Screen"` | Known | Screen layout |
| 0x00736FE5 | `Clock_Screen"` | Known | Screen layout |
| 0x0073708A | `Clock_Screen"` | Known | Screen layout |
| 0x0073712F | `Clock_Screen"` | Known | Screen layout |
| 0x007371D4 | `Clock_Screen"` | Known | Screen layout |
| 0x00737279 | `Clock_Screen"` | Known | Screen layout |
| 0x0073731E | `Clock_Screen"` | Known | Screen layout |
| 0x007373C3 | `Clock_Screen"` | Known | Screen layout |
| 0x00737468 | `Clock_Screen"` | Known | Screen layout |
| 0x0073750D | `Clock_Screen"` | Known | Screen layout |
| 0x007375B2 | `Clock_Screen"` | Known | Screen layout |
| 0x00737657 | `Clock_Screen"` | Known | Screen layout |
| 0x007376FA | `Settings_DateTime_SetDate_Screen(` | Known | Screen layout |
| 0x0073771E | `Settings_DateTime_SetDate_Screen_Default"` | Known | Screen layout |
| 0x00737797 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x007377FD | `Settings_DateTime_SetTime_Screen(` | Known | Screen layout |
| 0x00737821 | `Settings_DateTime_SetTime_Screen_Default"` | Known | Screen layout |
| 0x0073789A | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x00737905 | `Settings_DateTime_SetTimeZone_Screen,` | Known | Screen layout |
| 0x0073792D | `Settings_DateTime_SetTimeZone_Screen_Default"` | Known | Screen layout |
| 0x007379AA | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x00737A63 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x00737B13 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x007380A2 | `Search_Main_Screen` | Known | Screen layout |
| 0x007380B8 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x0073855C | `Extras_Screen` | Known | Screen layout |
| 0x0073856D | `Extras_Screen_WorldClock"` | Known | Screen layout |
| 0x007385EA | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x0073864C | `Clock_Screen` | Known | Screen layout |
| 0x0073865C | `Clock_Screen_Default` | Known | Screen layout |
| 0x007386E3 | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x00738749 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0073875F | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x007387CA | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x0073882C | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x00738844 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x007388B1 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x00738915 | `AddressViewer_Main_Screen!` | Known | Screen layout |
| 0x00738932 | `AddressViewer_Main_Screen_Default"` | Known | Screen layout |
| 0x007389A4 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x00738A0B | `Games_Menu_Screen` | Known | Screen layout |
| 0x00738A20 | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x00738A8A | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x00738B51 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x00738BED | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x00738CBE | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x00738D7E | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x00738DE2 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x00738E01 | `VoiceMemos_Menu_Screen_Default"` | Known | Screen layout |
| 0x00738E84 | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x00738EEA | `Speakers_Main_Screen` | Known | Screen layout |
| 0x00738F02 | `Speakers_Main_Screen_Default"` | Known | Screen layout |
| 0x00738F83 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x00738FE7 | `Radio_Screen` | Known | Screen layout |
| 0x00738FF7 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x00739070 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x0073910E | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x007391AA | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x0073926D | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x0073932C | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x007393E9 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x00739804 | `Radio_Screen` | Known | Screen layout |
| 0x00739814 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x0073988D | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x00739AAE | `Search_Main_Screen` | Known | Screen layout |
| 0x00739AC4 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x00739BF0 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x00739C53 | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x00739F12 | `Video_Settings_Screen` | Known | Screen layout |
| 0x00739F2B | `Video_Settings_Screen_Default"` | Known | Screen layout |
| 0x0073A012 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0073A0CF | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x0073A0EC | `SettingsMenu_About_Screen_Capacity_Layout"` | Known | Screen layout |
| 0x0073A339 | `SettingsMenus_MainMenu_Screen%` | Known | Screen layout |
| 0x0073A447 | `SettingsMenus_MusicMenu_Screen&` | Known | Screen layout |
| 0x0073A6F0 | `SettingsMenus_Brightness_Screen&` | Known | Screen layout |
| 0x0073A805 | `SettingsMenus_AudiobookSettings_Screen.` | Known | Screen layout |
| 0x0073A93B | `SettingsMenus_RadioRegions_Screen ` | Known | Screen layout |
| 0x0073AA50 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x0073ACBC | `Settings_DateTime_Screen ` | Known | Screen layout |
| 0x0073ACD8 | `Settings_DateTime_Screen_Default"` | Known | Screen layout |
| 0x0073AE64 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x0073AF69 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x0073AF82 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x0073B073 | `SettingsMenus_ResetAllSettings_Screen4` | Known | Screen layout |
| 0x0073B846 | `Stopwatch_Screen` | Known | Screen layout |
| 0x0073B85A | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x0073B8C1 | `Stopwatch_Screen` | Known | Screen layout |
| 0x0073B8D5 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x0073B97E | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x0073B9A1 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x0073BA3A | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x0073BA5D | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x0073BAEE | `NikePlus_ResumeWorkout_Screen%` | Known | Screen layout |
| 0x0073BB0F | `NikePlus_ResumeWorkout_Screen_Default"` | Known | Screen layout |
| 0x0073BB85 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0073BC2B | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0073BCD8 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0073BD88 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0073BE38 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0073BE9A | `NikePlus_Settings_Screen ` | Known | Screen layout |
| 0x0073BEB6 | `NikePlus_Settings_Screen_Default` | Known | Screen layout |
| 0x0073BF39 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0073BF9B | `NikePlus_History_Screen` | Known | Screen layout |
| 0x0073BFB6 | `NikePlus_History_Screen_Default` | Known | Screen layout |
| 0x0073C038 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0073C24C | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x0073C2BA | `Speakers_ToneControl_Screen#` | Known | Screen layout |
| 0x0073C2D9 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x007509A5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00750A28 | `LockediPod_Screen` | Known | Screen layout |
| 0x00750AB0 | `Lock_Screen` | Known | Screen layout |
| 0x00750ABF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00750B36 | `AddressViewer_ContactDetails_Screen+` | Known | Screen layout |
| 0x00750B5D | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x00750BD8 | `Extras_Screen` | Known | Screen layout |
| 0x00750C23 | `Extras_Screen` | Known | Screen layout |
| 0x00750D0A | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x00750D68 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x00750D85 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x00750DF3 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00750E0C | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x00750E83 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x00750EA0 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x00750F0B | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x00750F28 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x00750F8F | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x00750FF6 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x00751054 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x00751071 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x007510DF | `Calendar_Event_Screen` | Known | Screen layout |
| 0x007510F8 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0075116F | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0075118C | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x007511F7 | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x00751214 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0075127B | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x0075131B | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x007513A4 | `Alarms_Set_Alarm_Frequency_Screen)` | Known | Screen layout |
| 0x007513C9 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x0075143A | `Alarms_Set_Alarm_Label_Screen%` | Known | Screen layout |
| 0x0075145B | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x007514C8 | `Alarms_Set_Alarm_Sound_Screen%` | Known | Screen layout |
| 0x007514E9 | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x00751555 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x007517D0 | `Alarms_Set_Alarm_Playlist_Screen(` | Known | Screen layout |
| 0x007517F4 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x00751864 | `Alarms_Set_Alarm_Tones_Screen%` | Known | Screen layout |
| 0x00751885 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x00751B98 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x00751BB3 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x00751D04 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x00751D1B | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x00751D9C | `Calendar_Day_Screen` | Known | Screen layout |
| 0x00751DB3 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x00751E89 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00751EA2 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x00751F27 | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x00751F98 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0075208D | `Calendar_Event_Screen` | Known | Screen layout |
| 0x007520A6 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0075212B | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x0075219C | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0075225C | `ToDo_Item_Screen` | Known | Screen layout |
| 0x00752270 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x0075239F | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x00752402 | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x00752459 | `Clock_Screen_Default` | Known | Screen layout |
| 0x007524EA | `Clock_Region_Screen` | Known | Screen layout |
| 0x00752501 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x0075257A | `Clock_Screen_Default` | Known | Screen layout |
| 0x007525D1 | `Clock_Screen_Default` | Known | Screen layout |
| 0x00752662 | `Clock_Region_Screen` | Known | Screen layout |
| 0x00752679 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x00752804 | `CoverFlow_Screen_DefaultNoAlbumInfo&` | Known | Screen layout |
| 0x007528F2 | `CoverFlow_Screen_DefaultNoAlbumInfo#` | Known | Screen layout |
| 0x00752967 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00752C5D | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00752E0D | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x00752F3B | `CoverFlow_Screen_Default&` | Known | Screen layout |
| 0x00753011 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x007531A6 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0075340B | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x00753468 | `Game_Screen` | Known | Screen layout |
| 0x00753477 | `Game_Screen_Default` | Known | Screen layout |
| 0x00753519 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0075357B | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x007535DE | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00753641 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0075369D | `Game_Running_Screen` | Known | Screen layout |
| 0x007536FD | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0075375F | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x007537C2 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00753825 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00753881 | `Game_Running_Screen` | Known | Screen layout |
| 0x007538E1 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00753943 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x007539A6 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00753A09 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00753A65 | `Game_Running_Screen` | Known | Screen layout |
| 0x00753AC5 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00753B27 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00753B8A | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00753BED | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00753C49 | `Game_Running_Screen` | Known | Screen layout |
| 0x00753CA9 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00753D0B | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00753D6E | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00753DD1 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00753E2D | `Game_Running_Screen` | Known | Screen layout |
| 0x00754073 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x007540D5 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00754138 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0075419B | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x007541F7 | `Game_Running_Screen` | Known | Screen layout |
| 0x007542AE | `Extras_Screen` | Known | Screen layout |
| 0x007542BF | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0075431D | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x007544BA | `Extras_Screen` | Known | Screen layout |
| 0x007544CB | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00754529 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x007546C6 | `Extras_Screen` | Known | Screen layout |
| 0x007546D7 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00754735 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x007548D2 | `Extras_Screen` | Known | Screen layout |
| 0x007548E3 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00754941 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x00754AE3 | `Lock_Screen` | Known | Screen layout |
| 0x00754AF2 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x00754B54 | `Extras_Screen` | Known | Screen layout |
| 0x00754B65 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x00754BC4 | `LockediPod_Screen` | Known | Screen layout |
| 0x00754C3E | `MainMenus_Main_Screen)` | Known | Screen layout |
| 0x00754E0F | `Lock_Screen` | Known | Screen layout |
| 0x00754E1E | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x00754E80 | `Extras_Screen` | Known | Screen layout |
| 0x00754E91 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x00754EF0 | `LockediPod_Screen` | Known | Screen layout |
| 0x00754F6A | `MainMenus_Main_Screen%` | Known | Screen layout |
| 0x00754FD1 | `LockediPod_Screen` | Known | Screen layout |
| 0x00754FE6 | `LockediPod_Screen_Default,` | Known | Screen layout |
| 0x00755135 | `Lock_Screen` | Known | Screen layout |
| 0x00755144 | `Lock_Screen_Default_Layout"` | Known | Screen layout |
| 0x007551AD | `Lock_Screen` | Known | Screen layout |
| 0x007551BC | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0075521E | `Extras_Screen` | Known | Screen layout |
| 0x0075522F | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0075528E | `LockediPod_Screen` | Known | Screen layout |
| 0x00755308 | `MainMenus_Main_Screen,` | Known | Screen layout |
| 0x00755464 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x007554CA | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0075552E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007555BD | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0075562A | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x00755697 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x00755704 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075576C | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x007557D2 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x00755836 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007558C5 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x00755932 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0075599F | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x00755A0C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00755A74 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00755ADA | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x00755B3E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00755BCD | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x00755C3A | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x00755CA7 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x00755D14 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00755D7C | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00755DE2 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x00755E46 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00755ED5 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x00755F42 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x00755FAF | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0075601C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00756084 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x007560EA | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0075614E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007561DD | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0075624A | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x007562B7 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x00756324 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075637D | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x007563E6 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x0075644D | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007564E8 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00756551 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x007565BA | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00756621 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007566BC | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00756725 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x0075678E | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x007567F5 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00756890 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00756978 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00756994 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00756A02 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00756A1F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00756A8A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00756AAA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00756B21 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00756B3D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00756BAD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00756BCC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00756C38 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00756C4C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00756CC5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00756D39 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00756DA9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00756E11 | `NoContent_Screen` | Known | Screen layout |
| 0x00756E25 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00756E89 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00756EF0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00756F0A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00756F78 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00756FEA | `NoContent_Screen` | Known | Screen layout |
| 0x00756FFE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00757068 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007570D1 | `No_Photos_Screen` | Known | Screen layout |
| 0x007570E5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075714B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007571B9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00757226 | `NoContent_Screen` | Known | Screen layout |
| 0x0075723A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007572A2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075730C | `NoContent_Screen` | Known | Screen layout |
| 0x00757320 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075738D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007573FF | `NoContent_Screen` | Known | Screen layout |
| 0x00757413 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075747B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007574E4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007574FF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00757565 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00757581 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00757660 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00757679 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007576DA | `FirstBoot_Screen` | Known | Screen layout |
| 0x007576EE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075785C | `Radio_Screen` | Known | Screen layout |
| 0x0075786C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007578CD | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00757950 | `LockediPod_Screen` | Known | Screen layout |
| 0x007579D8 | `Lock_Screen` | Known | Screen layout |
| 0x007579E7 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00757A4A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00757AAC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00757AC8 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00757B3A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00757B59 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00757BC1 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00757BDB | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00757C43 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00757C60 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00757CCC | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00757D36 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00757D50 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00757DC0 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00757E33 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00757EA4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00757F13 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00757F7F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00757F9A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075800F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00758076 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007580D8 | `Photos_Screen` | Known | Screen layout |
| 0x0075813C | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075815A | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007581CA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007581E5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075824E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075826B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007582E2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00758306 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00758374 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075838F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075844C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00758468 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007584D6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007584F3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075855E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075857E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007585F5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00758611 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00758681 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007586A0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075870C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00758720 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00758799 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075880D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075887D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007588E5 | `NoContent_Screen` | Known | Screen layout |
| 0x007588F9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075895D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007589C4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007589DE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00758A4C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00758ABE | `NoContent_Screen` | Known | Screen layout |
| 0x00758AD2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00758B3C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00758BA5 | `No_Photos_Screen` | Known | Screen layout |
| 0x00758BB9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00758C1F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00758C8D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00758CFA | `NoContent_Screen` | Known | Screen layout |
| 0x00758D0E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00758D76 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00758DE0 | `NoContent_Screen` | Known | Screen layout |
| 0x00758DF4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00758E61 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00758ED3 | `NoContent_Screen` | Known | Screen layout |
| 0x00758EE7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00758F4F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00758FB8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00758FD3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00759039 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00759055 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00759134 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075914D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007591AE | `FirstBoot_Screen` | Known | Screen layout |
| 0x007591C2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00759330 | `Radio_Screen` | Known | Screen layout |
| 0x00759340 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007593A1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00759424 | `LockediPod_Screen` | Known | Screen layout |
| 0x007594AC | `Lock_Screen` | Known | Screen layout |
| 0x007594BB | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075951E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00759580 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075959C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075960E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075962D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00759695 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007596AF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00759717 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00759734 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007597A0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075980A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00759824 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00759894 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00759907 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00759978 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007599E7 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00759A53 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00759A6E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00759AE3 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00759B4A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00759BAC | `Photos_Screen` | Known | Screen layout |
| 0x00759C10 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00759C2E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00759C9E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00759CB9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00759D22 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00759D3F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00759DB6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00759DDA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00759E48 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00759E63 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00759F20 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00759F3C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00759FAA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00759FC7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075A032 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075A052 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075A0C9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075A0E5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075A155 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075A174 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075A1E0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075A1F4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075A26D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075A2E1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075A351 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075A3B9 | `NoContent_Screen` | Known | Screen layout |
| 0x0075A3CD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075A431 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075A498 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075A4B2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075A520 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075A592 | `NoContent_Screen` | Known | Screen layout |
| 0x0075A5A6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075A610 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075A679 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075A68D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075A6F3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075A761 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075A7CE | `NoContent_Screen` | Known | Screen layout |
| 0x0075A7E2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075A84A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075A8B4 | `NoContent_Screen` | Known | Screen layout |
| 0x0075A8C8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075A935 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075A9A7 | `NoContent_Screen` | Known | Screen layout |
| 0x0075A9BB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075AA23 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075AA8C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075AAA7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075AB0D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075AB29 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075AC08 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075AC21 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075AC82 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075AC96 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075AE04 | `Radio_Screen` | Known | Screen layout |
| 0x0075AE14 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075AE75 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075AEF8 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075AF80 | `Lock_Screen` | Known | Screen layout |
| 0x0075AF8F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075AFF2 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075B054 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075B070 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075B0E2 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075B101 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075B169 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075B183 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075B1EB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075B208 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075B274 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075B2DE | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075B2F8 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075B368 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075B3DB | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075B44C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075B4BB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075B527 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075B542 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075B5B7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075B61E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075B680 | `Photos_Screen` | Known | Screen layout |
| 0x0075B6E4 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075B702 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075B772 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075B78D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075B7F6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075B813 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075B88A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075B8AE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075B91C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075B937 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075B9F4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075BA10 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075BA7E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075BA9B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075BB06 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075BB26 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075BB9D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075BBB9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075BC29 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075BC48 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075BCB4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075BCC8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075BD41 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075BDB5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075BE25 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075BE8D | `NoContent_Screen` | Known | Screen layout |
| 0x0075BEA1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075BF05 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075BF6C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075BF86 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075BFF4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075C066 | `NoContent_Screen` | Known | Screen layout |
| 0x0075C07A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075C0E4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075C14D | `No_Photos_Screen` | Known | Screen layout |
| 0x0075C161 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075C1C7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075C235 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075C2A2 | `NoContent_Screen` | Known | Screen layout |
| 0x0075C2B6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075C31E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075C388 | `NoContent_Screen` | Known | Screen layout |
| 0x0075C39C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075C409 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075C47B | `NoContent_Screen` | Known | Screen layout |
| 0x0075C48F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075C4F7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075C560 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075C57B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075C5E1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075C5FD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075C6DC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075C6F5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075C756 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075C76A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075C8D8 | `Radio_Screen` | Known | Screen layout |
| 0x0075C8E8 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075C949 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075C9CC | `LockediPod_Screen` | Known | Screen layout |
| 0x0075CA54 | `Lock_Screen` | Known | Screen layout |
| 0x0075CA63 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075CAC6 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075CB28 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075CB44 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075CBB6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075CBD5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075CC3D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075CC57 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075CCBF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075CCDC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075CD48 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075CDB2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075CDCC | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075CE3C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075CEAF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075CF20 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075CF8F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075CFFB | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075D016 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075D08B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075D0F2 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075D154 | `Photos_Screen` | Known | Screen layout |
| 0x0075D1B8 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075D1D6 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075D246 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075D261 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075D2CA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075D2E7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075D35E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075D382 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075D3F0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075D40B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075D4C8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075D4E4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075D552 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075D56F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075D5DA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075D5FA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075D671 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075D68D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075D6FD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075D71C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075D788 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075D79C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075D815 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075D889 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075D8F9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075D961 | `NoContent_Screen` | Known | Screen layout |
| 0x0075D975 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075D9D9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075DA40 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075DA5A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075DAC8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075DB3A | `NoContent_Screen` | Known | Screen layout |
| 0x0075DB4E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075DBB8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075DC21 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075DC35 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075DC9B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075DD09 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075DD76 | `NoContent_Screen` | Known | Screen layout |
| 0x0075DD8A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075DDF2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075DE5C | `NoContent_Screen` | Known | Screen layout |
| 0x0075DE70 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075DEDD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075DF4F | `NoContent_Screen` | Known | Screen layout |
| 0x0075DF63 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075DFCB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075E034 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075E04F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075E0B5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075E0D1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075E1B0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075E1C9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075E22A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075E23E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075E3AC | `Radio_Screen` | Known | Screen layout |
| 0x0075E3BC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075E41D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075E4A0 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075E528 | `Lock_Screen` | Known | Screen layout |
| 0x0075E537 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075E59A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075E5FC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075E618 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075E68A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075E6A9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075E711 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075E72B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075E793 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075E7B0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075E81C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075E886 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075E8A0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075E910 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075E983 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075E9F4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075EA63 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075EACF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075EAEA | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075EB5F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075EBC6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075EC28 | `Photos_Screen` | Known | Screen layout |
| 0x0075EC8C | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075ECAA | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075ED1A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075ED35 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075ED9E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075EDBB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075EE32 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075EE56 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075EEC4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075EEDF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075EF9C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075EFB8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075F026 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075F043 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075F0AE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075F0CE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075F145 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075F161 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075F1D1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075F1F0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075F25C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075F270 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075F2E9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075F35D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075F3CD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075F435 | `NoContent_Screen` | Known | Screen layout |
| 0x0075F449 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075F4AD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075F514 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075F52E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075F59C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075F60E | `NoContent_Screen` | Known | Screen layout |
| 0x0075F622 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075F68C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075F6F5 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075F709 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075F76F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075F7DD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075F84A | `NoContent_Screen` | Known | Screen layout |
| 0x0075F85E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075F8C6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075F930 | `NoContent_Screen` | Known | Screen layout |
| 0x0075F944 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075F9B1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075FA23 | `NoContent_Screen` | Known | Screen layout |
| 0x0075FA37 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075FA9F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075FB08 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075FB23 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075FB89 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075FBA5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075FC84 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075FC9D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075FCFE | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075FD12 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075FE80 | `Radio_Screen` | Known | Screen layout |
| 0x0075FE90 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075FEF1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075FF74 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075FFFC | `Lock_Screen` | Known | Screen layout |
| 0x0076000B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076006E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007600D0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007600EC | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076015E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076017D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007601E5 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007601FF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00760267 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00760284 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007602F0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076035A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00760374 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007603E4 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00760457 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007604C8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00760537 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007605A3 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007605BE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00760633 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076069A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007606FC | `Photos_Screen` | Known | Screen layout |
| 0x00760760 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076077E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007607EE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00760809 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00760872 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076088F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00760906 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076092A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00760998 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007609B3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00760A70 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00760A8C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00760AFA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00760B17 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00760B82 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00760BA2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00760C19 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00760C35 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00760CA5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00760CC4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00760D30 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00760D44 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00760DBD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00760E31 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00760EA1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00760F09 | `NoContent_Screen` | Known | Screen layout |
| 0x00760F1D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00760F81 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00760FE8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00761002 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00761070 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007610E2 | `NoContent_Screen` | Known | Screen layout |
| 0x007610F6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00761160 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007611C9 | `No_Photos_Screen` | Known | Screen layout |
| 0x007611DD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00761243 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007612B1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076131E | `NoContent_Screen` | Known | Screen layout |
| 0x00761332 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076139A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00761404 | `NoContent_Screen` | Known | Screen layout |
| 0x00761418 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00761485 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007614F7 | `NoContent_Screen` | Known | Screen layout |
| 0x0076150B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00761573 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007615DC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007615F7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076165D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00761679 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00761758 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00761771 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007617D2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007617E6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00761954 | `Radio_Screen` | Known | Screen layout |
| 0x00761964 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007619C5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00761A48 | `LockediPod_Screen` | Known | Screen layout |
| 0x00761AD0 | `Lock_Screen` | Known | Screen layout |
| 0x00761ADF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00761B42 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00761BA4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00761BC0 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00761C32 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00761C51 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00761CB9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00761CD3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00761D3B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00761D58 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00761DC4 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00761E2E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00761E48 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00761EB8 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00761F2B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00761F9C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076200B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00762077 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00762092 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00762107 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076216E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007621D0 | `Photos_Screen` | Known | Screen layout |
| 0x00762234 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00762252 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007622C2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007622DD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00762346 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00762363 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007623DA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007623FE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076246C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00762487 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00762544 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00762560 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007625CE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007625EB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00762656 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00762676 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007626ED | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00762709 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00762779 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00762798 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00762804 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00762818 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00762891 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00762905 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00762975 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007629DD | `NoContent_Screen` | Known | Screen layout |
| 0x007629F1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00762A55 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00762ABC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00762AD6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00762B44 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00762BB6 | `NoContent_Screen` | Known | Screen layout |
| 0x00762BCA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00762C34 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00762C9D | `No_Photos_Screen` | Known | Screen layout |
| 0x00762CB1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00762D17 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00762D85 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00762DF2 | `NoContent_Screen` | Known | Screen layout |
| 0x00762E06 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00762E6E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00762ED8 | `NoContent_Screen` | Known | Screen layout |
| 0x00762EEC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00762F59 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00762FCB | `NoContent_Screen` | Known | Screen layout |
| 0x00762FDF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00763047 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007630B0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007630CB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00763131 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076314D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076322C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00763245 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007632A6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007632BA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00763428 | `Radio_Screen` | Known | Screen layout |
| 0x00763438 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00763499 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076351C | `LockediPod_Screen` | Known | Screen layout |
| 0x007635A4 | `Lock_Screen` | Known | Screen layout |
| 0x007635B3 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00763616 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00763678 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00763694 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00763706 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00763725 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076378D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007637A7 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076380F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076382C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00763898 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00763902 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076391C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076398C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007639FF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00763A70 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00763ADF | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00763B4B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00763B66 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00763BDB | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00763C42 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00763CA4 | `Photos_Screen` | Known | Screen layout |
| 0x00763D08 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00763D26 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00763D96 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00763DB1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00763E1A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00763E37 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00763EAE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00763ED2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00763F40 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00763F5B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00764018 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00764034 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007640A2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007640BF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076412A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076414A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007641C1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007641DD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076424D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076426C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007642D8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007642EC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00764365 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007643D9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00764449 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007644B1 | `NoContent_Screen` | Known | Screen layout |
| 0x007644C5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00764529 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00764590 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007645AA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00764618 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076468A | `NoContent_Screen` | Known | Screen layout |
| 0x0076469E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00764708 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00764771 | `No_Photos_Screen` | Known | Screen layout |
| 0x00764785 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007647EB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00764859 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007648C6 | `NoContent_Screen` | Known | Screen layout |
| 0x007648DA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00764942 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007649AC | `NoContent_Screen` | Known | Screen layout |
| 0x007649C0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00764A2D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00764A9F | `NoContent_Screen` | Known | Screen layout |
| 0x00764AB3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00764B1B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00764B84 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00764B9F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00764C05 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00764C21 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00764D00 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00764D19 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00764D7A | `FirstBoot_Screen` | Known | Screen layout |
| 0x00764D8E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00764EFC | `Radio_Screen` | Known | Screen layout |
| 0x00764F0C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00764F6D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00764FF0 | `LockediPod_Screen` | Known | Screen layout |
| 0x00765078 | `Lock_Screen` | Known | Screen layout |
| 0x00765087 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007650EA | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076514C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00765168 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007651DA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007651F9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00765261 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076527B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007652E3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00765300 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076536C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007653D6 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007653F0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00765460 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007654D3 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00765544 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007655B3 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076561F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076563A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007656AF | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00765716 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00765778 | `Photos_Screen` | Known | Screen layout |
| 0x007657DC | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007657FA | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076586A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00765885 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007658EE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076590B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00765982 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007659A6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00765A14 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00765A2F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00765AEC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00765B08 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00765B76 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00765B93 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00765BFE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00765C1E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00765C95 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00765CB1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00765D21 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00765D40 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00765DAC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00765DC0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00765E39 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00765EAD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00765F1D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00765F85 | `NoContent_Screen` | Known | Screen layout |
| 0x00765F99 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00765FFD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00766064 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076607E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007660EC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076615E | `NoContent_Screen` | Known | Screen layout |
| 0x00766172 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007661DC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00766245 | `No_Photos_Screen` | Known | Screen layout |
| 0x00766259 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007662BF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076632D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076639A | `NoContent_Screen` | Known | Screen layout |
| 0x007663AE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00766416 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00766480 | `NoContent_Screen` | Known | Screen layout |
| 0x00766494 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00766501 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00766573 | `NoContent_Screen` | Known | Screen layout |
| 0x00766587 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007665EF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00766658 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00766673 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007666D9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007666F5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007667D4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007667ED | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076684E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00766862 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007669D0 | `Radio_Screen` | Known | Screen layout |
| 0x007669E0 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00766A41 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00766AC4 | `LockediPod_Screen` | Known | Screen layout |
| 0x00766B4C | `Lock_Screen` | Known | Screen layout |
| 0x00766B5B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00766BBE | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00766C20 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00766C3C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00766CAE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00766CCD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00766D35 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00766D4F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00766DB7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00766DD4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00766E40 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00766EAA | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00766EC4 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00766F34 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00766FA7 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00767018 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00767087 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007670F3 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076710E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00767183 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007671EA | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076724C | `Photos_Screen` | Known | Screen layout |
| 0x007672B0 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007672CE | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076733E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00767359 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007673C2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007673DF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00767456 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076747A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007674E8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00767503 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007675C0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007675DC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076764A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00767667 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007676D2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007676F2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00767769 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00767785 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007677F5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00767814 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00767880 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00767894 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076790D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00767981 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007679F1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00767A59 | `NoContent_Screen` | Known | Screen layout |
| 0x00767A6D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00767AD1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00767B38 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00767B52 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00767BC0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00767C32 | `NoContent_Screen` | Known | Screen layout |
| 0x00767C46 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00767CB0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00767D19 | `No_Photos_Screen` | Known | Screen layout |
| 0x00767D2D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00767D93 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00767E01 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00767E6E | `NoContent_Screen` | Known | Screen layout |
| 0x00767E82 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00767EEA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00767F54 | `NoContent_Screen` | Known | Screen layout |
| 0x00767F68 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00767FD5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00768047 | `NoContent_Screen` | Known | Screen layout |
| 0x0076805B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007680C3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076812C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00768147 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007681AD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007681C9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007682A8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007682C1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00768322 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00768336 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007684A4 | `Radio_Screen` | Known | Screen layout |
| 0x007684B4 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00768515 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00768598 | `LockediPod_Screen` | Known | Screen layout |
| 0x00768620 | `Lock_Screen` | Known | Screen layout |
| 0x0076862F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00768692 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007686F4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00768710 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00768782 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007687A1 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00768809 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00768823 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076888B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007688A8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00768914 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076897E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00768998 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00768A08 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00768A7B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00768AEC | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00768B5B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00768BC7 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00768BE2 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00768C57 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00768CBE | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00768D20 | `Photos_Screen` | Known | Screen layout |
| 0x00768D84 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00768DA2 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00768E12 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00768E2D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00768E96 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00768EB3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00768F2A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00768F4E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00768FBC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00768FD7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00769094 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007690B0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076911E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076913B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007691A6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007691C6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076923D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00769259 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007692C9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007692E8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00769354 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00769368 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007693E1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00769455 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007694C5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076952D | `NoContent_Screen` | Known | Screen layout |
| 0x00769541 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007695A5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076960C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00769626 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00769694 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00769706 | `NoContent_Screen` | Known | Screen layout |
| 0x0076971A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00769784 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007697ED | `No_Photos_Screen` | Known | Screen layout |
| 0x00769801 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00769867 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007698D5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00769942 | `NoContent_Screen` | Known | Screen layout |
| 0x00769956 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007699BE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00769A28 | `NoContent_Screen` | Known | Screen layout |
| 0x00769A3C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00769AA9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00769B1B | `NoContent_Screen` | Known | Screen layout |
| 0x00769B2F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00769B97 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00769C00 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00769C1B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00769C81 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00769C9D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00769D7C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00769D95 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00769DF6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00769E0A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00769F78 | `Radio_Screen` | Known | Screen layout |
| 0x00769F88 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00769FE9 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076A06C | `LockediPod_Screen` | Known | Screen layout |
| 0x0076A0F4 | `Lock_Screen` | Known | Screen layout |
| 0x0076A103 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076A166 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076A1C8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076A1E4 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076A256 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076A275 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076A2DD | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076A2F7 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076A35F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076A37C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076A3E8 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076A452 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076A46C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076A4DC | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076A54F | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076A5C0 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076A62F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076A69B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076A6B6 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076A72B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076A792 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076A7F4 | `Photos_Screen` | Known | Screen layout |
| 0x0076A858 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076A876 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076A8E6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076A901 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076A96A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076A987 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076A9FE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076AA22 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076AA90 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076AAAB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076AB68 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076AB84 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076ABF2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076AC0F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076AC7A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076AC9A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076AD11 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076AD2D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076AD9D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076ADBC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076AE28 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076AE3C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076AEB5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076AF29 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076AF99 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076B001 | `NoContent_Screen` | Known | Screen layout |
| 0x0076B015 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076B079 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076B0E0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076B0FA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076B168 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076B1DA | `NoContent_Screen` | Known | Screen layout |
| 0x0076B1EE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076B258 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076B2C1 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076B2D5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076B33B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076B3A9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076B416 | `NoContent_Screen` | Known | Screen layout |
| 0x0076B42A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076B492 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076B4FC | `NoContent_Screen` | Known | Screen layout |
| 0x0076B510 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076B57D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076B5EF | `NoContent_Screen` | Known | Screen layout |
| 0x0076B603 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076B66B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076B6D4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076B6EF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076B755 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076B771 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076B850 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076B869 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076B8CA | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076B8DE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076BA4C | `Radio_Screen` | Known | Screen layout |
| 0x0076BA5C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076BABD | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076BB40 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076BBC8 | `Lock_Screen` | Known | Screen layout |
| 0x0076BBD7 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076BC3A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076BC9C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076BCB8 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076BD2A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076BD49 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076BDB1 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076BDCB | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076BE33 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076BE50 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076BEBC | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076BF26 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076BF40 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076BFB0 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076C023 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076C094 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076C103 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076C16F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076C18A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076C1FF | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076C266 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076C2C8 | `Photos_Screen` | Known | Screen layout |
| 0x0076C32C | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076C34A | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076C3BA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076C3D5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076C43E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076C45B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076C4D2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076C4F6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076C564 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076C57F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076C63C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076C658 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076C6C6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076C6E3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076C74E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076C76E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076C7E5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076C801 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076C871 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076C890 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076C8FC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076C910 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076C989 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076C9FD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076CA6D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076CAD5 | `NoContent_Screen` | Known | Screen layout |
| 0x0076CAE9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076CB4D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076CBB4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076CBCE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076CC3C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076CCAE | `NoContent_Screen` | Known | Screen layout |
| 0x0076CCC2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076CD2C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076CD95 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076CDA9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076CE0F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076CE7D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076CEEA | `NoContent_Screen` | Known | Screen layout |
| 0x0076CEFE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076CF66 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076CFD0 | `NoContent_Screen` | Known | Screen layout |
| 0x0076CFE4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076D051 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076D0C3 | `NoContent_Screen` | Known | Screen layout |
| 0x0076D0D7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076D13F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076D1A8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076D1C3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076D229 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076D245 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076D324 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076D33D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076D39E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076D3B2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076D520 | `Radio_Screen` | Known | Screen layout |
| 0x0076D530 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076D591 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076D614 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076D69C | `Lock_Screen` | Known | Screen layout |
| 0x0076D6AB | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076D70E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076D770 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076D78C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076D7FE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076D81D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076D885 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076D89F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076D907 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076D924 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076D990 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076D9FA | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076DA14 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076DA84 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076DAF7 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076DB68 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076DBD7 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076DC43 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076DC5E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076DCD3 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076DD3A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076DD9C | `Photos_Screen` | Known | Screen layout |
| 0x0076DE00 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076DE1E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076DE8E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076DEA9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076DF12 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076DF2F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076DFA6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076DFCA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076E038 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076E053 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076E110 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076E12C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076E19A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076E1B7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076E222 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076E242 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076E2B9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076E2D5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076E345 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076E364 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076E3D0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076E3E4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076E45D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076E4D1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076E541 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076E5A9 | `NoContent_Screen` | Known | Screen layout |
| 0x0076E5BD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076E621 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076E688 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076E6A2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076E710 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076E782 | `NoContent_Screen` | Known | Screen layout |
| 0x0076E796 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076E800 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076E869 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076E87D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076E8E3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076E951 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076E9BE | `NoContent_Screen` | Known | Screen layout |
| 0x0076E9D2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076EA3A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076EAA4 | `NoContent_Screen` | Known | Screen layout |
| 0x0076EAB8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076EB25 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076EB97 | `NoContent_Screen` | Known | Screen layout |
| 0x0076EBAB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076EC13 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076EC7C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076EC97 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076ECFD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076ED19 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076EDF8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076EE11 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076EE72 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076EE86 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076EFF4 | `Radio_Screen` | Known | Screen layout |
| 0x0076F004 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076F065 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076F0E8 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076F170 | `Lock_Screen` | Known | Screen layout |
| 0x0076F17F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076F1E2 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076F244 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076F260 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076F2D2 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076F2F1 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076F359 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076F373 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076F3DB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076F3F8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076F464 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076F4CE | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076F4E8 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076F558 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076F5CB | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076F63C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076F6AB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076F717 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076F732 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076F7A7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076F80E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076F870 | `Photos_Screen` | Known | Screen layout |
| 0x0076F8D4 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076F8F2 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076F962 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076F97D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076F9E6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076FA03 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076FA7A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076FA9E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076FB0C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076FB27 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076FBE4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076FC00 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076FC6E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076FC8B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076FCF6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076FD16 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076FD8D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076FDA9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076FE19 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076FE38 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076FEA4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076FEB8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076FF31 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076FFA5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00770015 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077007D | `NoContent_Screen` | Known | Screen layout |
| 0x00770091 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007700F5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077015C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00770176 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007701E4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00770256 | `NoContent_Screen` | Known | Screen layout |
| 0x0077026A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007702D4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077033D | `No_Photos_Screen` | Known | Screen layout |
| 0x00770351 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007703B7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00770425 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00770492 | `NoContent_Screen` | Known | Screen layout |
| 0x007704A6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077050E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00770578 | `NoContent_Screen` | Known | Screen layout |
| 0x0077058C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007705F9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077066B | `NoContent_Screen` | Known | Screen layout |
| 0x0077067F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007706E7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00770750 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077076B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007707D1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007707ED | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007708CC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007708E5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00770946 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077095A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00770AC8 | `Radio_Screen` | Known | Screen layout |
| 0x00770AD8 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00770B39 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00770BBC | `LockediPod_Screen` | Known | Screen layout |
| 0x00770C44 | `Lock_Screen` | Known | Screen layout |
| 0x00770C53 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00770CB6 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00770D18 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00770D34 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00770DA6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00770DC5 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00770E2D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00770E47 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00770EAF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00770ECC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00770F38 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00770FA2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00770FBC | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077102C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077109F | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00771110 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077117F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007711EB | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00771206 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077127B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007712E2 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00771344 | `Photos_Screen` | Known | Screen layout |
| 0x007713A8 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007713C6 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00771436 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00771451 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007714BA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007714D7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077154E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00771572 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007715E0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007715FB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007716B8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007716D4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00771742 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077175F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007717CA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007717EA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00771861 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077187D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007718ED | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077190C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00771978 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077198C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00771A05 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00771A79 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00771AE9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00771B51 | `NoContent_Screen` | Known | Screen layout |
| 0x00771B65 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00771BC9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00771C30 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00771C4A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00771CB8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00771D2A | `NoContent_Screen` | Known | Screen layout |
| 0x00771D3E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00771DA8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00771E11 | `No_Photos_Screen` | Known | Screen layout |
| 0x00771E25 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00771E8B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00771EF9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00771F66 | `NoContent_Screen` | Known | Screen layout |
| 0x00771F7A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00771FE2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077204C | `NoContent_Screen` | Known | Screen layout |
| 0x00772060 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007720CD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077213F | `NoContent_Screen` | Known | Screen layout |
| 0x00772153 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007721BB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00772224 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077223F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007722A5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007722C1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007723A0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007723B9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077241A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077242E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077259C | `Radio_Screen` | Known | Screen layout |
| 0x007725AC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077260D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00772690 | `LockediPod_Screen` | Known | Screen layout |
| 0x00772718 | `Lock_Screen` | Known | Screen layout |
| 0x00772727 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077278A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007727EC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00772808 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077287A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00772899 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00772901 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077291B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00772983 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007729A0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00772A0C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00772A76 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00772A90 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00772B00 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00772B73 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00772BE4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00772C53 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00772CBF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00772CDA | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00772D4F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00772DB6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00772E18 | `Photos_Screen` | Known | Screen layout |
| 0x00772E7C | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00772E9A | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00772F0A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00772F25 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00772F8E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00772FAB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00773022 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00773046 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007730B4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007730CF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077318C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007731A8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00773216 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00773233 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077329E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007732BE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00773335 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00773351 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007733C1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007733E0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077344C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00773460 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007734D9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077354D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007735BD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00773625 | `NoContent_Screen` | Known | Screen layout |
| 0x00773639 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077369D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00773704 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077371E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077378C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007737FE | `NoContent_Screen` | Known | Screen layout |
| 0x00773812 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077387C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007738E5 | `No_Photos_Screen` | Known | Screen layout |
| 0x007738F9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077395F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007739CD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00773A3A | `NoContent_Screen` | Known | Screen layout |
| 0x00773A4E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00773AB6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00773B20 | `NoContent_Screen` | Known | Screen layout |
| 0x00773B34 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00773BA1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00773C13 | `NoContent_Screen` | Known | Screen layout |
| 0x00773C27 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00773C8F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00773CF8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00773D13 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00773D79 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00773D95 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00773E74 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00773E8D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00773EEE | `FirstBoot_Screen` | Known | Screen layout |
| 0x00773F02 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00774070 | `Radio_Screen` | Known | Screen layout |
| 0x00774080 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007740E1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00774164 | `LockediPod_Screen` | Known | Screen layout |
| 0x007741EC | `Lock_Screen` | Known | Screen layout |
| 0x007741FB | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077425E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007742C0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007742DC | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077434E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077436D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007743D5 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007743EF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00774457 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00774474 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007744E0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077454A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00774564 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007745D4 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00774647 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007746B8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00774727 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00774793 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007747AE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00774823 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077488A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007748EC | `Photos_Screen` | Known | Screen layout |
| 0x00774950 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077496E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007749DE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007749F9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00774A62 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00774A7F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00774AF6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00774B1A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00774B88 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00774BA3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00774C60 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00774C7C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00774CEA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00774D07 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00774D72 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00774D92 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00774E09 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00774E25 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00774E95 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00774EB4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00774F20 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00774F34 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00774FAD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00775021 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00775091 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007750F9 | `NoContent_Screen` | Known | Screen layout |
| 0x0077510D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00775171 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007751D8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007751F2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00775260 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007752D2 | `NoContent_Screen` | Known | Screen layout |
| 0x007752E6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00775350 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007753B9 | `No_Photos_Screen` | Known | Screen layout |
| 0x007753CD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00775433 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007754A1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077550E | `NoContent_Screen` | Known | Screen layout |
| 0x00775522 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077558A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007755F4 | `NoContent_Screen` | Known | Screen layout |
| 0x00775608 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00775675 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007756E7 | `NoContent_Screen` | Known | Screen layout |
| 0x007756FB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00775763 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007757CC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007757E7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077584D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00775869 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00775948 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00775961 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007759C2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007759D6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00775B44 | `Radio_Screen` | Known | Screen layout |
| 0x00775B54 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00775BB5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00775C38 | `LockediPod_Screen` | Known | Screen layout |
| 0x00775CC0 | `Lock_Screen` | Known | Screen layout |
| 0x00775CCF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00775D32 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00775D94 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00775DB0 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00775E22 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00775E41 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00775EA9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00775EC3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00775F2B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00775F48 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00775FB4 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077601E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00776038 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007760A8 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077611B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077618C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007761FB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00776267 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00776282 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007762F7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077635E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007763C0 | `Photos_Screen` | Known | Screen layout |
| 0x00776424 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00776442 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007764B2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007764CD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00776536 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00776553 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007765CA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007765EE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077665C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00776677 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00776734 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00776750 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007767BE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007767DB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00776846 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00776866 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007768DD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007768F9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00776969 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00776988 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007769F4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00776A08 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00776A81 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00776AF5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00776B65 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00776BCD | `NoContent_Screen` | Known | Screen layout |
| 0x00776BE1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00776C45 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00776CAC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00776CC6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00776D34 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00776DA6 | `NoContent_Screen` | Known | Screen layout |
| 0x00776DBA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00776E24 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00776E8D | `No_Photos_Screen` | Known | Screen layout |
| 0x00776EA1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00776F07 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00776F75 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00776FE2 | `NoContent_Screen` | Known | Screen layout |
| 0x00776FF6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077705E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007770C8 | `NoContent_Screen` | Known | Screen layout |
| 0x007770DC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00777149 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007771BB | `NoContent_Screen` | Known | Screen layout |
| 0x007771CF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00777237 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007772A0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007772BB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00777321 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077733D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077741C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00777435 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00777496 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007774AA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00777618 | `Radio_Screen` | Known | Screen layout |
| 0x00777628 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00777689 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077770C | `LockediPod_Screen` | Known | Screen layout |
| 0x00777794 | `Lock_Screen` | Known | Screen layout |
| 0x007777A3 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00777806 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00777868 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00777884 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007778F6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00777915 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077797D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00777997 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007779FF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00777A1C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00777A88 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00777AF2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00777B0C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00777B7C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00777BEF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00777C60 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00777CCF | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00777D3B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00777D56 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00777DCB | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00777E32 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00777E94 | `Photos_Screen` | Known | Screen layout |
| 0x00777EF8 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00777F16 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00777F86 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00777FA1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077800A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00778027 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077809E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007780C2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00778130 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077814B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00778208 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00778224 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00778292 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007782AF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077831A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077833A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007783B1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007783CD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077843D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077845C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007784C8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007784DC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00778555 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007785C9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00778639 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007786A1 | `NoContent_Screen` | Known | Screen layout |
| 0x007786B5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00778719 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00778780 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077879A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00778808 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077887A | `NoContent_Screen` | Known | Screen layout |
| 0x0077888E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007788F8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00778961 | `No_Photos_Screen` | Known | Screen layout |
| 0x00778975 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007789DB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00778A49 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00778AB6 | `NoContent_Screen` | Known | Screen layout |
| 0x00778ACA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00778B32 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00778B9C | `NoContent_Screen` | Known | Screen layout |
| 0x00778BB0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00778C1D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00778C8F | `NoContent_Screen` | Known | Screen layout |
| 0x00778CA3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00778D0B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00778D74 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00778D8F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00778DF5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00778E11 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00778EF0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00778F09 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00778F6A | `FirstBoot_Screen` | Known | Screen layout |
| 0x00778F7E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007790EC | `Radio_Screen` | Known | Screen layout |
| 0x007790FC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077915D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007791E0 | `LockediPod_Screen` | Known | Screen layout |
| 0x00779268 | `Lock_Screen` | Known | Screen layout |
| 0x00779277 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007792DA | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077933C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00779358 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007793CA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007793E9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00779451 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077946B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007794D3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007794F0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077955C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007795C6 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007795E0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00779650 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007796C3 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00779734 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007797A3 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077980F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077982A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077989F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00779906 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00779968 | `Photos_Screen` | Known | Screen layout |
| 0x007799CC | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007799EA | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00779A5A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00779A75 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00779ADE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00779AFB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00779B72 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00779B96 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00779C04 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00779C1F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00779CDC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00779CF8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00779D66 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00779D83 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00779DEE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00779E0E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00779E85 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00779EA1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00779F11 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00779F30 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00779F9C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00779FB0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077A029 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077A09D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077A10D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077A175 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A189 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077A1ED | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077A254 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077A26E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077A2DC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077A34E | `NoContent_Screen` | Known | Screen layout |
| 0x0077A362 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077A3CC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077A435 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077A449 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077A4AF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077A51D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077A58A | `NoContent_Screen` | Known | Screen layout |
| 0x0077A59E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077A606 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077A670 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A684 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077A6F1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077A763 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A777 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077A7DF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077A848 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077A863 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077A8C9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077A8E5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077A9C4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077A9DD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077AA3E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077AA52 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077ABC0 | `Radio_Screen` | Known | Screen layout |
| 0x0077ABD0 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077AC31 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077ACB4 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077AD3C | `Lock_Screen` | Known | Screen layout |
| 0x0077AD4B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077ADAE | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077AE10 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077AE2C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077AE9E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077AEBD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077AF25 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077AF3F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077AFA7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077AFC4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077B030 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077B09A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077B0B4 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077B124 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077B197 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077B208 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077B277 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077B2E3 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077B2FE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077B373 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077B3DA | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077B43C | `Photos_Screen` | Known | Screen layout |
| 0x0077B4A0 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077B4BE | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077B52E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077B549 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077B5B2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077B5CF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077B646 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077B66A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077B6D8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077B6F3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077B7B0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077B7CC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077B83A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077B857 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077B8C2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077B8E2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077B959 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077B975 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077B9E5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077BA04 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077BA70 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077BA84 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077BAFD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077BB71 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077BBE1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077BC49 | `NoContent_Screen` | Known | Screen layout |
| 0x0077BC5D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077BCC1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077BD28 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077BD42 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077BDB0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077BE22 | `NoContent_Screen` | Known | Screen layout |
| 0x0077BE36 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077BEA0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077BF09 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077BF1D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077BF83 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077BFF1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077C05E | `NoContent_Screen` | Known | Screen layout |
| 0x0077C072 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077C0DA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077C144 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C158 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077C1C5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077C237 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C24B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077C2B3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077C31C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077C337 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077C39D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077C3B9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077C498 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077C4B1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077C512 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077C526 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077C694 | `Radio_Screen` | Known | Screen layout |
| 0x0077C6A4 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077C705 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077C788 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077C810 | `Lock_Screen` | Known | Screen layout |
| 0x0077C81F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077C882 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077C8E4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077C900 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077C972 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077C991 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077C9F9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077CA13 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077CA7B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077CA98 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077CB04 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077CB6E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077CB88 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077CBF8 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077CC6B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077CCDC | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077CD4B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077CDB7 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077CDD2 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077CE47 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077CEAE | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077CF10 | `Photos_Screen` | Known | Screen layout |
| 0x0077CF74 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077CF92 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077D002 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077D01D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077D086 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077D0A3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077D11A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077D13E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077D1AC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077D1C7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077D284 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077D2A0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077D30E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077D32B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077D396 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077D3B6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077D42D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077D449 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077D4B9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077D4D8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077D544 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077D558 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077D5D1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077D645 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077D6B5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077D71D | `NoContent_Screen` | Known | Screen layout |
| 0x0077D731 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077D795 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077D7FC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077D816 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077D884 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077D8F6 | `NoContent_Screen` | Known | Screen layout |
| 0x0077D90A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077D974 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077D9DD | `No_Photos_Screen` | Known | Screen layout |
| 0x0077D9F1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077DA57 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077DAC5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077DB32 | `NoContent_Screen` | Known | Screen layout |
| 0x0077DB46 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077DBAE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077DC18 | `NoContent_Screen` | Known | Screen layout |
| 0x0077DC2C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077DC99 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077DD0B | `NoContent_Screen` | Known | Screen layout |
| 0x0077DD1F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077DD87 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077DDF0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077DE0B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077DE71 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077DE8D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077DF6C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077DF85 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077DFE6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077DFFA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077E168 | `Radio_Screen` | Known | Screen layout |
| 0x0077E178 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077E1D9 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077E25C | `LockediPod_Screen` | Known | Screen layout |
| 0x0077E2E4 | `Lock_Screen` | Known | Screen layout |
| 0x0077E2F3 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077E356 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077E3B8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077E3D4 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077E446 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077E465 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077E4CD | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077E4E7 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077E54F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077E56C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077E5D8 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077E642 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077E65C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077E6CC | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077E73F | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077E7B0 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077E81F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077E88B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077E8A6 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077E91B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077E982 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077E9E4 | `Photos_Screen` | Known | Screen layout |
| 0x0077EA48 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077EA66 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077EAD6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077EAF1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077EB5A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077EB77 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077EBEE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077EC12 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077EC80 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077EC9B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077ED58 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077ED74 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077EDE2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077EDFF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077EE6A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077EE8A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077EF01 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077EF1D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077EF8D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077EFAC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077F018 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077F02C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077F0A5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077F119 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077F189 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077F1F1 | `NoContent_Screen` | Known | Screen layout |
| 0x0077F205 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077F269 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077F2D0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077F2EA | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077F358 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077F3CA | `NoContent_Screen` | Known | Screen layout |
| 0x0077F3DE | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077F448 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077F4B1 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077F4C5 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077F52B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077F599 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077F606 | `NoContent_Screen` | Known | Screen layout |
| 0x0077F61A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077F682 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077F6EC | `NoContent_Screen` | Known | Screen layout |
| 0x0077F700 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077F76D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077F7DF | `NoContent_Screen` | Known | Screen layout |
| 0x0077F7F3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077F85B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077F8C4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077F8DF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077F945 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077F961 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077FA40 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077FA59 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077FABA | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077FACE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077FC3C | `Radio_Screen` | Known | Screen layout |
| 0x0077FC4C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077FCAD | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077FD30 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077FDB8 | `Lock_Screen` | Known | Screen layout |
| 0x0077FDC7 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077FE2A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077FE8C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077FEA8 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077FF1A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077FF39 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077FFA1 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077FFBB | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00780023 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00780040 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007800AC | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00780116 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00780130 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007801A0 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00780213 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00780284 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007802F3 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078035F | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078037A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007803EF | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00780456 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007804B8 | `Photos_Screen` | Known | Screen layout |
| 0x0078051C | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078053A | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007805AA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007805C5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078062E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078064B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007806C2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007806E6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00780754 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078076F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078082C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00780848 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007808B6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007808D3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078093E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078095E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007809D5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007809F1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00780A61 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00780A80 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00780AEC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00780B00 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00780B79 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00780BED | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00780C5D | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00780CC5 | `NoContent_Screen` | Known | Screen layout |
| 0x00780CD9 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00780D3D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00780DA4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00780DBE | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00780E2C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00780E9E | `NoContent_Screen` | Known | Screen layout |
| 0x00780EB2 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00780F1C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00780F85 | `No_Photos_Screen` | Known | Screen layout |
| 0x00780F99 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00780FFF | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078106D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007810DA | `NoContent_Screen` | Known | Screen layout |
| 0x007810EE | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00781156 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007811C0 | `NoContent_Screen` | Known | Screen layout |
| 0x007811D4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00781241 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007812B3 | `NoContent_Screen` | Known | Screen layout |
| 0x007812C7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078132F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00781398 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007813B3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00781419 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00781435 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00781514 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078152D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078158E | `FirstBoot_Screen` | Known | Screen layout |
| 0x007815A2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00781710 | `Radio_Screen` | Known | Screen layout |
| 0x00781720 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00781781 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00781804 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078188C | `Lock_Screen` | Known | Screen layout |
| 0x0078189B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007818FE | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00781960 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078197C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007819EE | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00781A0D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00781A75 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00781A8F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00781AF7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00781B14 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00781B80 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00781BEA | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00781C04 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00781C74 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00781CE7 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00781D58 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00781DC7 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00781E33 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00781E4E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00781EC3 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00781F2A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00781F8C | `Photos_Screen` | Known | Screen layout |
| 0x00781FF0 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078200E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078207E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00782099 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00782102 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078211F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00782196 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007821BA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00782228 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00782243 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00782300 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078231C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078238A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007823A7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00782412 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00782432 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007824A9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007824C5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00782535 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00782554 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007825C0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007825D4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078264D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007826C1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00782731 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00782799 | `NoContent_Screen` | Known | Screen layout |
| 0x007827AD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00782811 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00782878 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00782892 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00782900 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00782972 | `NoContent_Screen` | Known | Screen layout |
| 0x00782986 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007829F0 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00782A59 | `No_Photos_Screen` | Known | Screen layout |
| 0x00782A6D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00782AD3 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00782B41 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00782BAE | `NoContent_Screen` | Known | Screen layout |
| 0x00782BC2 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00782C2A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00782C94 | `NoContent_Screen` | Known | Screen layout |
| 0x00782CA8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00782D15 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00782D87 | `NoContent_Screen` | Known | Screen layout |
| 0x00782D9B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00782E03 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00782E6C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00782E87 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00782EED | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00782F09 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00782FE8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00783001 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00783062 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00783076 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007831E4 | `Radio_Screen` | Known | Screen layout |
| 0x007831F4 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00783255 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007832D8 | `LockediPod_Screen` | Known | Screen layout |
| 0x00783360 | `Lock_Screen` | Known | Screen layout |
| 0x0078336F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007833D2 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00783434 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00783450 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007834C2 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007834E1 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00783549 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00783563 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007835CB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007835E8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00783654 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007836BE | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007836D8 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00783748 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007837BB | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078382C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078389B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00783907 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00783922 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00783997 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007839FE | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00783A60 | `Photos_Screen` | Known | Screen layout |
| 0x00783AC4 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00783AE2 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00783B52 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00783B6D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00783BD6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00783BF3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00783C6A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00783C8E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00783CFC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00783D17 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00783DD4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00783DF0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00783E5E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00783E7B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00783EE6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00783F06 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00783F7D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00783F99 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00784009 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00784028 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00784094 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007840A8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00784121 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00784195 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00784205 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078426D | `NoContent_Screen` | Known | Screen layout |
| 0x00784281 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007842E5 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078434C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00784366 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007843D4 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00784446 | `NoContent_Screen` | Known | Screen layout |
| 0x0078445A | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007844C4 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078452D | `No_Photos_Screen` | Known | Screen layout |
| 0x00784541 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007845A7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00784615 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00784682 | `NoContent_Screen` | Known | Screen layout |
| 0x00784696 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007846FE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00784768 | `NoContent_Screen` | Known | Screen layout |
| 0x0078477C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007847E9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078485B | `NoContent_Screen` | Known | Screen layout |
| 0x0078486F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007848D7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00784940 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078495B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007849C1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007849DD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00784ABC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00784AD5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00784B36 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00784B4A | `FirstBoot_Screen_Default` | Known | Screen layout |
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
| 0x00785598 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007855B6 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00785626 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00785641 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007856AA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007856C7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078573E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00785762 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007857D0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007857EB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007858A8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007858C4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00785932 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078594F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007859BA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007859DA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00785A51 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00785A6D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00785ADD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00785AFC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00785B68 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00785B7C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00785BF5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00785C69 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00785CD9 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00785D41 | `NoContent_Screen` | Known | Screen layout |
| 0x00785D55 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00785DB9 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00785E20 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00785E3A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00785EA8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00785F1A | `NoContent_Screen` | Known | Screen layout |
| 0x00785F2E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00785F98 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00786001 | `No_Photos_Screen` | Known | Screen layout |
| 0x00786015 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078607B | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007860E9 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00786156 | `NoContent_Screen` | Known | Screen layout |
| 0x0078616A | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007861D2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078623C | `NoContent_Screen` | Known | Screen layout |
| 0x00786250 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007862BD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078632F | `NoContent_Screen` | Known | Screen layout |
| 0x00786343 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007863AB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00786414 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078642F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00786495 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007864B1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00786590 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007865A9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078660A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078661E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078678C | `Radio_Screen` | Known | Screen layout |
| 0x0078679C | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007867FD | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00786880 | `LockediPod_Screen` | Known | Screen layout |
| 0x00786908 | `Lock_Screen` | Known | Screen layout |
| 0x00786917 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078697A | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007869DC | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007869F8 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00786A6A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00786A89 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00786AF1 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00786B0B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00786B73 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00786B90 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00786BFC | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00786C66 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00786C80 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00786CF0 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00786D63 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00786DD4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00786E43 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00786EAF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00786ECA | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00786F3F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00786FA6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00787008 | `Photos_Screen` | Known | Screen layout |
| 0x0078706C | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078708A | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007870FA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00787115 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078717E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078719B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00787212 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00787236 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007872A4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007872BF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078737C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00787398 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00787406 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00787423 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078748E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007874AE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00787525 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00787541 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007875B1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007875D0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078763C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00787650 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007876C9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078773D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007877AD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00787815 | `NoContent_Screen` | Known | Screen layout |
| 0x00787829 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078788D | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007878F4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078790E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078797C | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007879EE | `NoContent_Screen` | Known | Screen layout |
| 0x00787A02 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00787A6C | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00787AD5 | `No_Photos_Screen` | Known | Screen layout |
| 0x00787AE9 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00787B4F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00787BBD | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00787C2A | `NoContent_Screen` | Known | Screen layout |
| 0x00787C3E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00787CA6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00787D10 | `NoContent_Screen` | Known | Screen layout |
| 0x00787D24 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00787D91 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00787E03 | `NoContent_Screen` | Known | Screen layout |
| 0x00787E17 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00787E7F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00787EE8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00787F03 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00787F69 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00787F85 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00788064 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078807D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007880DE | `FirstBoot_Screen` | Known | Screen layout |
| 0x007880F2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00788260 | `Radio_Screen` | Known | Screen layout |
| 0x00788270 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007882D1 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00788354 | `LockediPod_Screen` | Known | Screen layout |
| 0x007883DC | `Lock_Screen` | Known | Screen layout |
| 0x007883EB | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078844E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007884B0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007884CC | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078853E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078855D | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007885C5 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007885DF | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00788647 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00788664 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007886D0 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078873A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00788754 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007887C4 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00788837 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007888A8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00788917 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00788983 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078899E | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00788A13 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00788A7A | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00788ADC | `Photos_Screen` | Known | Screen layout |
| 0x00788B40 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00788B5E | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00788BCE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00788BE9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00788C52 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00788C6F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00788CE6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00788D0A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00788D78 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00788D93 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00788E50 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00788E6C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00788EDA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00788EF7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00788F62 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00788F82 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00788FF9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00789015 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00789085 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007890A4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00789110 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00789124 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078919D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00789211 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00789281 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007892E9 | `NoContent_Screen` | Known | Screen layout |
| 0x007892FD | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00789361 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007893C8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007893E2 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00789450 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007894C2 | `NoContent_Screen` | Known | Screen layout |
| 0x007894D6 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00789540 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007895A9 | `No_Photos_Screen` | Known | Screen layout |
| 0x007895BD | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00789623 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00789691 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007896FE | `NoContent_Screen` | Known | Screen layout |
| 0x00789712 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078977A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007897E4 | `NoContent_Screen` | Known | Screen layout |
| 0x007897F8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00789865 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007898D7 | `NoContent_Screen` | Known | Screen layout |
| 0x007898EB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00789953 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007899BC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007899D7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00789A3D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00789A59 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00789B38 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00789B51 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00789BB2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00789BC6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00789D34 | `Radio_Screen` | Known | Screen layout |
| 0x00789D44 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00789DA5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00789E28 | `LockediPod_Screen` | Known | Screen layout |
| 0x00789EB0 | `Lock_Screen` | Known | Screen layout |
| 0x00789EBF | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00789F22 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00789F84 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00789FA0 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078A012 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078A031 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078A099 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078A0B3 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078A11B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078A138 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078A1A4 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078A20E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078A228 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078A298 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078A30B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078A37C | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078A3EB | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078A457 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078A472 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078A4E7 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078A54E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078A5B0 | `Photos_Screen` | Known | Screen layout |
| 0x0078A614 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078A632 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078A6A2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078A6BD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078A726 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078A743 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078A7BA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078A7DE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078A84C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078A867 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078A924 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078A940 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078A9AE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078A9CB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078AA36 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078AA56 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078AACD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078AAE9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078AB59 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078AB78 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078ABE4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078ABF8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078AC71 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078ACE5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078AD55 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078ADBD | `NoContent_Screen` | Known | Screen layout |
| 0x0078ADD1 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078AE35 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078AE9C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078AEB6 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078AF24 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078AF96 | `NoContent_Screen` | Known | Screen layout |
| 0x0078AFAA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078B014 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078B07D | `No_Photos_Screen` | Known | Screen layout |
| 0x0078B091 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078B0F7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078B165 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078B1D2 | `NoContent_Screen` | Known | Screen layout |
| 0x0078B1E6 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078B24E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078B2B8 | `NoContent_Screen` | Known | Screen layout |
| 0x0078B2CC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078B339 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078B3AB | `NoContent_Screen` | Known | Screen layout |
| 0x0078B3BF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078B427 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078B490 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078B4AB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078B511 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078B52D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078B60C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078B625 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078B686 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078B69A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078B808 | `Radio_Screen` | Known | Screen layout |
| 0x0078B818 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078B879 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078B8FC | `LockediPod_Screen` | Known | Screen layout |
| 0x0078B984 | `Lock_Screen` | Known | Screen layout |
| 0x0078B993 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078B9F6 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078BA58 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078BA74 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078BAE6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078BB05 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078BB6D | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078BB87 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078BBEF | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078BC0C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078BC78 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078BCE2 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078BCFC | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078BD6C | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078BDDF | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078BE50 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078BEBF | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078BF2B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078BF46 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078BFBB | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078C022 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078C084 | `Photos_Screen` | Known | Screen layout |
| 0x0078C0E8 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078C106 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078C176 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078C191 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078C1FA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078C217 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078C28E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078C2B2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078C320 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078C33B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078C3F8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078C414 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078C482 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078C49F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078C50A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078C52A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078C5A1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078C5BD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078C62D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078C64C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078C6B8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078C6CC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078C745 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078C7B9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078C829 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078C891 | `NoContent_Screen` | Known | Screen layout |
| 0x0078C8A5 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078C909 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078C970 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078C98A | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078C9F8 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078CA6A | `NoContent_Screen` | Known | Screen layout |
| 0x0078CA7E | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078CAE8 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078CB51 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078CB65 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078CBCB | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078CC39 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078CCA6 | `NoContent_Screen` | Known | Screen layout |
| 0x0078CCBA | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078CD22 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078CD8C | `NoContent_Screen` | Known | Screen layout |
| 0x0078CDA0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078CE0D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078CE7F | `NoContent_Screen` | Known | Screen layout |
| 0x0078CE93 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078CEFB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078CF64 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078CF7F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078CFE5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078D001 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078D0E0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078D0F9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078D15A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078D16E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078D2DC | `Radio_Screen` | Known | Screen layout |
| 0x0078D2EC | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078D34D | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078D3D0 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078D458 | `Lock_Screen` | Known | Screen layout |
| 0x0078D467 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078D4CA | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078D52C | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078D548 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078D5BA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078D5D9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078D641 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078D65B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078D6C3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078D6E0 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078D74C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078D7B6 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078D7D0 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078D840 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078D8B3 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078D924 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078D993 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078D9FF | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078DA1A | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078DA8F | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078DAF6 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078DB58 | `Photos_Screen` | Known | Screen layout |
| 0x0078DBBC | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078DBDA | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078DC4A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078DC65 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078DCCE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078DCEB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078DD62 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078DD86 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078DDF4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078DE0F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078DECC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078DEE8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078DF56 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078DF73 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078DFDE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078DFFE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078E075 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078E091 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078E101 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078E120 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078E18C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078E1A0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078E219 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078E28D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078E2FD | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078E365 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E379 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078E3DD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078E444 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078E45E | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078E4CC | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078E53E | `NoContent_Screen` | Known | Screen layout |
| 0x0078E552 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078E5BC | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0078E625 | `No_Photos_Screen` | Known | Screen layout |
| 0x0078E639 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078E69F | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078E70D | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0078E77A | `NoContent_Screen` | Known | Screen layout |
| 0x0078E78E | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0078E7F6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078E860 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E874 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0078E8E1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078E953 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E967 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078E9CF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0078EA38 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078EA53 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078EAB9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078EAD5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078EBB4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078EBCD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078EC2E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078EC42 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0078EDB0 | `Radio_Screen` | Known | Screen layout |
| 0x0078EDC0 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0078EE21 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0078EEA4 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078EF2C | `Lock_Screen` | Known | Screen layout |
| 0x0078EF3B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0078EF9E | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0078F000 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078F01C | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0078F08E | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078F0AD | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078F115 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078F12F | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0078F197 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078F1B4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078F220 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078F28A | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078F2A4 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0078F314 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0078F387 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0078F3F8 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0078F467 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078F4D3 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0078F4EE | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078F563 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0078F5CA | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0078F62C | `Photos_Screen` | Known | Screen layout |
| 0x0078F690 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0078F6AE | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0078F71E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078F739 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078F7A2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078F7BF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0078F836 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078F85A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0078F8C8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078F8E3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078F9A0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078F9BC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078FA2A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078FA47 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078FAB2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078FAD2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078FB49 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078FB65 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078FBD5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078FBF4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078FC60 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078FC74 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0078FCED | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078FD61 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0078FDD1 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0078FE39 | `NoContent_Screen` | Known | Screen layout |
| 0x0078FE4D | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0078FEB1 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0078FF18 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078FF32 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0078FFA0 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00790012 | `NoContent_Screen` | Known | Screen layout |
| 0x00790026 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00790090 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007900F9 | `No_Photos_Screen` | Known | Screen layout |
| 0x0079010D | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00790173 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007901E1 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0079024E | `NoContent_Screen` | Known | Screen layout |
| 0x00790262 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007902CA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00790334 | `NoContent_Screen` | Known | Screen layout |
| 0x00790348 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007903B5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00790427 | `NoContent_Screen` | Known | Screen layout |
| 0x0079043B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007904A3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0079050C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00790527 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079058D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007905A9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00790688 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007906A1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00790702 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00790716 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00790884 | `Radio_Screen` | Known | Screen layout |
| 0x00790894 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007908F5 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00790978 | `LockediPod_Screen` | Known | Screen layout |
| 0x00790A00 | `Lock_Screen` | Known | Screen layout |
| 0x00790A0F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00790A72 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00790AD4 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00790AF0 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00790B62 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00790B81 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00790BE9 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00790C03 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00790C6B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00790C88 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00790CF4 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00790D5E | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00790D78 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00790DE8 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00790E5B | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00790ECC | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00790F3B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00790FA7 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00790FC2 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00791037 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0079109E | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00791100 | `Photos_Screen` | Known | Screen layout |
| 0x00791164 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00791182 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007911F2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079120D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00791276 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00791293 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079130A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079132E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079139C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007913B7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00791474 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00791490 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007914FE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079151B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00791586 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007915A6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079161D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00791639 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007916A9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007916C8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00791734 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00791748 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007917C1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00791835 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007918A5 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0079190D | `NoContent_Screen` | Known | Screen layout |
| 0x00791921 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00791985 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007919EC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00791A06 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00791A74 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00791AE6 | `NoContent_Screen` | Known | Screen layout |
| 0x00791AFA | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00791B64 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00791BCD | `No_Photos_Screen` | Known | Screen layout |
| 0x00791BE1 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00791C47 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00791CB5 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00791D22 | `NoContent_Screen` | Known | Screen layout |
| 0x00791D36 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00791D9E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00791E08 | `NoContent_Screen` | Known | Screen layout |
| 0x00791E1C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00791E89 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00791EFB | `NoContent_Screen` | Known | Screen layout |
| 0x00791F0F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00791F77 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00791FE0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00791FFB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00792061 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079207D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079215C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00792175 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007921D6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007921EA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00792358 | `Radio_Screen` | Known | Screen layout |
| 0x00792368 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007923C9 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0079244C | `LockediPod_Screen` | Known | Screen layout |
| 0x007924D4 | `Lock_Screen` | Known | Screen layout |
| 0x007924E3 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00792546 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007925A8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007925C4 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00792636 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00792655 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007926BD | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007926D7 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0079273F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079275C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007927C8 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00792832 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0079284C | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007928BC | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0079292F | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007929A0 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00792A0F | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00792A7B | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00792A96 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00792B0B | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00792B72 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00792BD4 | `Photos_Screen` | Known | Screen layout |
| 0x00792C38 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00792C56 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00792CC6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00792CE1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00792D4A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00792D67 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00792DDE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00792E02 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00792E70 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00792E8B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00792F2D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00792F49 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00792FB7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00792FD4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079303F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079305F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007930D6 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007930F2 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00793162 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00793181 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007931ED | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00793201 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00793276 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007932E1 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00793350 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007933C1 | `NoContent_Screen` | Known | Screen layout |
| 0x007933D5 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00793444 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007934B7 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00793524 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0079358D | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007935FD | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0079366D | `NoContent_Screen` | Known | Screen layout |
| 0x00793681 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007936E4 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00793747 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00793763 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00793823 | `Radio_Screen` | Known | Screen layout |
| 0x00793833 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00793894 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00793902 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00793921 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079398F | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007939F4 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00793A0F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00793AB5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00793AD1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00793B3F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00793B5C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00793BC7 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00793BE7 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00793C5E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00793C7A | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00793CEA | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00793D09 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00793D75 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00793D89 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00793DFE | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00793E69 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00793ED8 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00793F49 | `NoContent_Screen` | Known | Screen layout |
| 0x00793F5D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00793FCC | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0079403F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007940AC | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00794115 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00794185 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007941F5 | `NoContent_Screen` | Known | Screen layout |
| 0x00794209 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0079426C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007942CF | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007942EB | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007943AB | `Radio_Screen` | Known | Screen layout |
| 0x007943BB | `Radio_Screen_Default` | Known | Screen layout |
| 0x0079441C | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0079448A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007944A9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00794517 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079457C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00794597 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079463D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00794659 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007946C7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007946E4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079474F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079476F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007947E6 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00794802 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00794872 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00794891 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007948FD | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00794911 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00794986 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007949F1 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00794A60 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00794AD1 | `NoContent_Screen` | Known | Screen layout |
| 0x00794AE5 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00794B54 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00794BC7 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00794C34 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00794C9D | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00794D0D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00794D7D | `NoContent_Screen` | Known | Screen layout |
| 0x00794D91 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00794DF4 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00794E57 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00794E73 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00794F33 | `Radio_Screen` | Known | Screen layout |
| 0x00794F43 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00794FA4 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00795012 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00795031 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079509F | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00795104 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079511F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007951C5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007951E1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079524F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079526C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007952D7 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007952F7 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079536E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079538A | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007953FA | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00795419 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00795485 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00795499 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0079550E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00795579 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x007955E8 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00795659 | `NoContent_Screen` | Known | Screen layout |
| 0x0079566D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007956DC | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0079574F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007957BC | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00795825 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00795895 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00795905 | `NoContent_Screen` | Known | Screen layout |
| 0x00795919 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0079597C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007959DF | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007959FB | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00795ABB | `Radio_Screen` | Known | Screen layout |
| 0x00795ACB | `Radio_Screen_Default` | Known | Screen layout |
| 0x00795B2C | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00795B9A | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00795BB9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00795C27 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00795C8C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00795CA7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00795D4D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00795D69 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00795DD7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00795DF4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00795E5F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00795E7F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00795EF6 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00795F12 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00795F82 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00795FA1 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079600D | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00796021 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00796096 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00796101 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00796170 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007961E1 | `NoContent_Screen` | Known | Screen layout |
| 0x007961F5 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00796264 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007962D7 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00796344 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007963AD | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0079641D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0079648D | `NoContent_Screen` | Known | Screen layout |
| 0x007964A1 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00796504 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00796567 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00796583 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00796643 | `Radio_Screen` | Known | Screen layout |
| 0x00796653 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007966B4 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00796722 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00796741 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007967AF | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00796814 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079682F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007968D5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007968F1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079695F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079697C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007969E7 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00796A07 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00796A7E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00796A9A | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00796B0A | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00796B29 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00796B95 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00796BA9 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00796C1E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00796C89 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00796CF8 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00796D69 | `NoContent_Screen` | Known | Screen layout |
| 0x00796D7D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00796DEC | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00796E5F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00796ECC | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00796F35 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00796FA5 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00797015 | `NoContent_Screen` | Known | Screen layout |
| 0x00797029 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0079708C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007970EF | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079710B | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007971CB | `Radio_Screen` | Known | Screen layout |
| 0x007971DB | `Radio_Screen_Default` | Known | Screen layout |
| 0x0079723C | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007972AA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007972C9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00797337 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079739C | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007973B7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079745D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00797479 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007974E7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00797504 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079756F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079758F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00797606 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00797622 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00797692 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007976B1 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079771D | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00797731 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007977A6 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00797811 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00797880 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007978F1 | `NoContent_Screen` | Known | Screen layout |
| 0x00797905 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00797974 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007979E7 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00797A54 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00797ABD | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00797B2D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00797B9D | `NoContent_Screen` | Known | Screen layout |
| 0x00797BB1 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00797C14 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00797C77 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00797C93 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00797D53 | `Radio_Screen` | Known | Screen layout |
| 0x00797D63 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00797DC4 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00797E32 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00797E51 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00797EBF | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00797F24 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00797F3F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00797FE5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00798001 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079806F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079808C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007980F7 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00798117 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079818E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007981AA | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079821A | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00798239 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007982A5 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007982B9 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0079832E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00798399 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00798408 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00798479 | `NoContent_Screen` | Known | Screen layout |
| 0x0079848D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007984FC | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0079856F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007985DC | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00798645 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007986B5 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00798725 | `NoContent_Screen` | Known | Screen layout |
| 0x00798739 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0079879C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x007987FF | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079881B | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007988DB | `Radio_Screen` | Known | Screen layout |
| 0x007988EB | `Radio_Screen_Default` | Known | Screen layout |
| 0x0079894C | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007989BA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007989D9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00798A47 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00798AAC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00798AC7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00798B6D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00798B89 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00798BF7 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00798C14 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00798C7F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00798C9F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00798D16 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00798D32 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00798DA2 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00798DC1 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00798E2D | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00798E41 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00798EB6 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00798F21 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00798F90 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00799001 | `NoContent_Screen` | Known | Screen layout |
| 0x00799015 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00799084 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007990F7 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00799164 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007991CD | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0079923D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007992AD | `NoContent_Screen` | Known | Screen layout |
| 0x007992C1 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00799324 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00799387 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007993A3 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00799463 | `Radio_Screen` | Known | Screen layout |
| 0x00799473 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007994D4 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00799542 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00799561 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007995CF | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00799634 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079964F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007996F5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00799711 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079977F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079979C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00799807 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00799827 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079989E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007998BA | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079992A | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00799949 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007999B5 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007999C9 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00799A3E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00799AA9 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00799B18 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00799B89 | `NoContent_Screen` | Known | Screen layout |
| 0x00799B9D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00799C0C | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00799C7F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00799CEC | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00799D55 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00799DC5 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00799E35 | `NoContent_Screen` | Known | Screen layout |
| 0x00799E49 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00799EAC | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00799F0F | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00799F2B | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00799FEB | `Radio_Screen` | Known | Screen layout |
| 0x00799FFB | `Radio_Screen_Default` | Known | Screen layout |
| 0x0079A05C | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0079A0CA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079A0E9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079A157 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079A1BC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079A1D7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079A27D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079A299 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079A307 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079A324 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079A38F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079A3AF | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079A426 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079A442 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079A4B2 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079A4D1 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079A53D | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079A551 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0079A5C6 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0079A631 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0079A6A0 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0079A711 | `NoContent_Screen` | Known | Screen layout |
| 0x0079A725 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079A794 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0079A807 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0079A874 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0079A8DD | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0079A94D | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0079A9BD | `NoContent_Screen` | Known | Screen layout |
| 0x0079A9D1 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0079AA34 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0079AA97 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079AAB3 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079AB73 | `Radio_Screen` | Known | Screen layout |
| 0x0079AB83 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0079ABE4 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0079AC52 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079AC71 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079ACDF | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079AD44 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079AD5F | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079AE05 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079AE21 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079AE8F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079AEAC | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079AF17 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0079AF37 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0079AFAE | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079AFCA | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0079B03A | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0079B059 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0079B0C5 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0079B0D9 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0079B14E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0079B1B9 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0079B228 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0079B299 | `NoContent_Screen` | Known | Screen layout |
| 0x0079B2AD | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0079B31C | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0079B38F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0079B3FC | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0079B465 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0079B4D5 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0079B545 | `NoContent_Screen` | Known | Screen layout |
| 0x0079B559 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0079B5BC | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0079B61F | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0079B63B | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0079B6FB | `Radio_Screen` | Known | Screen layout |
| 0x0079B70B | `Radio_Screen_Default` | Known | Screen layout |
| 0x0079B76C | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0079B7DA | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079B7F9 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0079B867 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0079B8CC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079B8E7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079B9C8 | `MediaLists_AudiobookChapters_Screen)` | Known | Screen layout |
| 0x0079B9EF | `MediaLists_AudiobookChapters_Screen_Plane!` | Known | Screen layout |
| 0x0079C189 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0079C1A4 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x0079C20F | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079C22A | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x0079C29D | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0079C2B8 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0079C475 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0079C490 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x0079C4FB | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079C516 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x0079C589 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0079C5A4 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0079C76C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079C788 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x0079C803 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079C81F | `MediaLists_Albums_Screen_WithTrackCount%` | Known | Screen layout |
| 0x0079C898 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0079C8B3 | `MediaLists_Songs_Screen_Plain1` | Known | Screen layout |
| 0x0079C92E | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0079C949 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0079CB6B | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079CB88 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079CC67 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0079CC83 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x0079CCFE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0079CD19 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0079CEFF | `MediaLists_PodcastEpisodes_Screen'` | Known | Screen layout |
| 0x0079CF24 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0079D1F6 | `MediaLists_TVSeasons_Screen#` | Known | Screen layout |
| 0x0079D215 | `MediaLists_TVSeasons_Screen_Default%` | Known | Screen layout |
| 0x0079D28A | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x0079D2AA | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0079D432 | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x0079D452 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0079D84B | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x0079D870 | `MediaLists_NestedPlaylists_Screen_Default,` | Known | Screen layout |
| 0x0079D8F2 | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x0079D911 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x0079DAA1 | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x0079DAC6 | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x0079DB3E | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x0079DB5D | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x0079DBC1 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079DC6E | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079DCE0 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x0079DDD6 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x0079DF78 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x0079E078 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x0079E0E4 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079E14E | `NoContent_Screen` | Known | Screen layout |
| 0x0079E162 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079E1CC | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079E240 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E254 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079E2BF | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079E32B | `NoContent_Screen` | Known | Screen layout |
| 0x0079E33F | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079E3AC | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079E420 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E434 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079E49C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079E509 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079E56D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079E589 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0079E5F5 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079E612 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079E67F | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079E6E9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079E706 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079E77D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079E7A1 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079E858 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079E8C2 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E8D6 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079E940 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079E9B4 | `NoContent_Screen` | Known | Screen layout |
| 0x0079E9C8 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079EA33 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079EA9F | `NoContent_Screen` | Known | Screen layout |
| 0x0079EAB3 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079EB20 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079EB94 | `NoContent_Screen` | Known | Screen layout |
| 0x0079EBA8 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079EC10 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079EC7D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079ECE1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079ECFD | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0079ED69 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079ED86 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079EDF3 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079EE5D | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079EE7A | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079EEF1 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079EF15 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079EFCC | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079F036 | `NoContent_Screen` | Known | Screen layout |
| 0x0079F04A | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079F0B4 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079F128 | `NoContent_Screen` | Known | Screen layout |
| 0x0079F13C | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079F1A7 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079F213 | `NoContent_Screen` | Known | Screen layout |
| 0x0079F227 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079F294 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079F308 | `NoContent_Screen` | Known | Screen layout |
| 0x0079F31C | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079F384 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079F3F1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079F455 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079F471 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0079F4DD | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079F4FA | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079F567 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079F5D1 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079F5EE | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079F665 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079F689 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079F740 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079F7AA | `NoContent_Screen` | Known | Screen layout |
| 0x0079F7BE | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079F828 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079F89C | `NoContent_Screen` | Known | Screen layout |
| 0x0079F8B0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079F91B | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079F987 | `NoContent_Screen` | Known | Screen layout |
| 0x0079F99B | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0079FA08 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079FA7C | `NoContent_Screen` | Known | Screen layout |
| 0x0079FA90 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0079FAF8 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0079FB65 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0079FBC9 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0079FBE5 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0079FC51 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079FC6E | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079FCDB | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x0079FD45 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079FD62 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079FDD9 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0079FDFD | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0079FEB4 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079FF1E | `NoContent_Screen` | Known | Screen layout |
| 0x0079FF32 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x0079FF9C | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A0010 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0024 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A008F | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007A00FB | `NoContent_Screen` | Known | Screen layout |
| 0x007A010F | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A017C | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A01F0 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0204 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A026C | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A02D9 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A033D | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A0359 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007A03C5 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A03E2 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A044F | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A04B9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A04D6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A054D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A0571 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A0628 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007A0692 | `NoContent_Screen` | Known | Screen layout |
| 0x007A06A6 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007A0710 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A0784 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0798 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A0803 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007A086F | `NoContent_Screen` | Known | Screen layout |
| 0x007A0883 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A08F0 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A0964 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0978 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A09E0 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A0A4D | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A0AB1 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A0ACD | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007A0B39 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A0B56 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A0BC3 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A0C2D | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A0C4A | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A0CC1 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A0CE5 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A0D9C | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007A0E06 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0E1A | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007A0E84 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007A0EF8 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0F0C | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007A0F77 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007A0FE3 | `NoContent_Screen` | Known | Screen layout |
| 0x007A0FF7 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007A1064 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007A10D8 | `NoContent_Screen` | Known | Screen layout |
| 0x007A10EC | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007A1154 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007A11C1 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007A1225 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007A1241 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007A12AD | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007A12CA | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007A1337 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007A13A1 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007A13BE | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007A1435 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007A1459 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007A17D0 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A1842 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A18AD | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A1912 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A197C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A19E6 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A1A4D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A1AB8 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A1B22 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A1B89 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A1BF0 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A1C55 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A1CBD | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A1D28 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A1D93 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A1DFA | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A20C0 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A2132 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A219D | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A2202 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A226C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A22D6 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A233D | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A23A8 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A2412 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A2479 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A24E0 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A2545 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A25AD | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A2618 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A2683 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A26EA | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A29AE | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A2A20 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A2A8B | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A2AF0 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A2B5A | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A2BC4 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A2C2B | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A2C96 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A2D00 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A2D67 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A2DCE | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A2E33 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A2E9B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A2F06 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A2F71 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A2FD8 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A329A | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A330C | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A3377 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A33DC | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A3446 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A34B0 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A3517 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A3582 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A35EC | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A3653 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A36BA | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A371F | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A3787 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A37F2 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A385D | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A38C4 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A3B6E | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A3BE0 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A3C4B | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A3CB0 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A3D1A | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A3D84 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A3DEB | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A3E56 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A3EC0 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A3F27 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A3F8E | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A3FF3 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A405B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A40C6 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A4131 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A4198 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A4467 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A44D9 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A4544 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A45A9 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A4613 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A467D | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A46E4 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A474F | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A47B9 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A4820 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A4887 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A48EC | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A4954 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A49BF | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A4A2A | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A4A91 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A4D5D | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A4DCF | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A4E3A | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A4E9F | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A4F09 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A4F73 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A4FDA | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A5045 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A50AF | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A5116 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A517D | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A51E2 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A524A | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A52B5 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A5320 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A5387 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A563E | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A56B0 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A571B | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A5780 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A57EA | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A5854 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A58BB | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A5926 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A5990 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A59F7 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A5A5E | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A5AC3 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A5B2B | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A5B96 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A5C01 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A5C68 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A5F44 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A5FB6 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A6021 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A6086 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A60F0 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A615A | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A61C1 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A622C | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A6296 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A62FD | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A6364 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A63C9 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A6431 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A649C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A6507 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A656E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A6858 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A68CA | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A6935 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A699A | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A6A04 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A6A6E | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A6AD5 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A6B40 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A6BAA | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A6C11 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A6C78 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A6CDD | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A6D45 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A6DB0 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A6E1B | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A6E82 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A714C | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A71BE | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A7229 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A728E | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A72F8 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A7362 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A73C9 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A7434 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A749E | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A7505 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A756C | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A75D1 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A7639 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A76A4 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A770F | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A7776 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A7A34 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A7AA6 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A7B11 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A7B76 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A7BE0 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A7C4A | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A7CB1 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A7D1C | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A7D86 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A7DED | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A7E54 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A7EB9 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A7F21 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A7F8C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A7FF7 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A805E | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A830A | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A837C | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A83E7 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A844C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A84B6 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A8520 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A8587 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A85F2 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A865C | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A86C3 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A872A | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A878F | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A87F7 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A8862 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A88CD | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A8934 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A8BD7 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A8C49 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A8CB4 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A8D19 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A8D83 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A8DED | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A8E54 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A8EBF | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A8F29 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A8F90 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A8FF7 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A905C | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A90C4 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A912F | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A919A | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A9201 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A94BF | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A9531 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A959C | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A9601 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A966B | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A96D5 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x007A973C | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A97A7 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A9811 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A9878 | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x007A98DF | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A9944 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A99AC | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A9A17 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A9A82 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A9AE9 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A9DA2 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x007A9E15 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007A9E87 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007A9EF7 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x007A9F65 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007A9FD2 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AA272 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x007AA2E5 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AA357 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AA3C7 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x007AA435 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AA4A2 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AA766 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x007AA7D7 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AA847 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x007AA8B5 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AA922 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AABB6 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x007AAC27 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AAC97 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x007AAD05 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AAD72 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AB004 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x007AB075 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AB0E5 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x007AB153 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AB1C0 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AB566 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007AB583 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007AB5FE | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007AB617 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007AB68F | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007AB6A8 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007AB71D | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AB733 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007AB7AA | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AB7C0 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007AB837 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007AB854 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007AB8CC | `Notes_List_Screen` | Known | Screen layout |
| 0x007AB8E1 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007ABA92 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007ABAAF | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007ABB2A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007ABB43 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007ABBBB | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007ABBD4 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007ABC49 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007ABC5F | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007ABCD6 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007ABCEC | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007ABD63 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007ABD80 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007ABDF8 | `Notes_List_Screen` | Known | Screen layout |
| 0x007ABE0D | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007ABFEE | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007AC00B | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007AC086 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007AC09F | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007AC117 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007AC130 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007AC1A5 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AC1BB | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007AC232 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AC248 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007AC2BF | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007AC2DC | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007AC354 | `Notes_List_Screen` | Known | Screen layout |
| 0x007AC369 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007AC51E | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007AC53B | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007AC5B6 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007AC5CF | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007AC647 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007AC660 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007AC6D5 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AC6EB | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007AC762 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AC778 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007AC7EF | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007AC80C | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007AC884 | `Notes_List_Screen` | Known | Screen layout |
| 0x007AC899 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007ACBB1 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007ACC57 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007ACCDA | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007ACD92 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x007ACE14 | `PhotosSettingsSlideshowMusic_Screen+` | Known | Screen layout |
| 0x007ACE3B | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x007ACF21 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x007AD0D9 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AD139 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AD196 | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007AD1BD | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007AD25D | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AD2BD | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AD31A | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007AD341 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007AD5DC | `Photos_Screen` | Known | Screen layout |
| 0x007AD728 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AD78C | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007AD7ED | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007AD84A | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007AD8A7 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007AD915 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007AD972 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007ADAE8 | `Photos_Screen` | Known | Screen layout |
| 0x007ADC34 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007ADC98 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007ADCF9 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007ADD56 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007ADDB3 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007ADE21 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007ADE7E | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007ADFF4 | `Photos_Screen` | Known | Screen layout |
| 0x007AE140 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AE1A4 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007AE205 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007AE262 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007AE2BF | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007AE32D | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007AE38A | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007AE500 | `Photos_Screen` | Known | Screen layout |
| 0x007AE64C | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AE6B0 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007AE711 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007AE76E | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007AE7CB | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007AE839 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007AE896 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007AEA0C | `Photos_Screen` | Known | Screen layout |
| 0x007AEB58 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AEBBC | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007AEC1D | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007AEC7A | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007AECD7 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007AED45 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007AEDA2 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007AEF18 | `Photos_Screen` | Known | Screen layout |
| 0x007AF064 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AF0C8 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007AF129 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007AF186 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007AF1E3 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007AF251 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007AF2AE | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007AF424 | `Photos_Screen` | Known | Screen layout |
| 0x007AF570 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AF5D6 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007AF638 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007AF69A | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007AF730 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007AF851 | `Photos_Screen` | Known | Screen layout |
| 0x007AF8BC | `Photos_Screen` | Known | Screen layout |
| 0x007AFA08 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AFA6E | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007AFAD0 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007AFB32 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007AFBC8 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007AFCE9 | `Photos_Screen` | Known | Screen layout |
| 0x007AFD54 | `Photos_Screen` | Known | Screen layout |
| 0x007AFEA0 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AFF06 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007AFF68 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007AFFCA | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B0060 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B0181 | `Photos_Screen` | Known | Screen layout |
| 0x007B01EC | `Photos_Screen` | Known | Screen layout |
| 0x007B0338 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B039E | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B0400 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B0462 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B04F8 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B0619 | `Photos_Screen` | Known | Screen layout |
| 0x007B0684 | `Photos_Screen` | Known | Screen layout |
| 0x007B07D0 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B0836 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B0898 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B08FA | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B0990 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B0AB1 | `Photos_Screen` | Known | Screen layout |
| 0x007B0C92 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B0CF4 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B0D62 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B0DC8 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B0E2D | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B10CD | `Radio_Screen_Default$` | Known | Screen layout |
| 0x007B1134 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B119A | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B1499 | `Radio_Screen_Default$` | Known | Screen layout |
| 0x007B1500 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B1566 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B180E | `Radio_Screen_Default$` | Known | Screen layout |
| 0x007B1875 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B18DB | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B1BC2 | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007B1C2C | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007B1E36 | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007B1EA0 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007B1FF9 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007B205C | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B20C1 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007B2129 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B218C | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B21F4 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B225D | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B22C3 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007B2328 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B23AD | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B243A | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B24D9 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B24F3 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B256B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B2585 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B25F9 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B2686 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B2725 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B273F | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B27B7 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B27D1 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B2845 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B28D2 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B2971 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B298B | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B2A03 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B2A1D | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B2A91 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B2B1E | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B2BBD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B2BD7 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B2C4F | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B2C69 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B2CDD | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B2D6A | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B2E09 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B2E23 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B2E9B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B2EB5 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B2F29 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B2FB6 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B3055 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B306F | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B30E7 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B3101 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B3175 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B3202 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B32A1 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B32BB | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B3333 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B334D | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B33C1 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B344E | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B34ED | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B3507 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B357F | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B3599 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B360D | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B369A | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B3739 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B3753 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B37CB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B37E5 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B3859 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B38E6 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B3985 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B399F | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B3A17 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B3A31 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B3AA5 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B3B32 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B3BD1 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B3BEB | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B3C63 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B3C7D | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B3CF1 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B3D7E | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B3E1D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B3E37 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B3EAF | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B3EC9 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B3F3D | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B3FCA | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B4069 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B4083 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B40FB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B4115 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B4189 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B4216 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B42B5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B42CF | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B4347 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B4361 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B43D5 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B4462 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B4501 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B451B | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B4593 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B45AD | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B4621 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B46AE | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B474D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B4767 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B47DF | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B47F9 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B486D | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B48FA | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B4999 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B49B3 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B4A2B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B4A45 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B4AB9 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B4B46 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B4BE5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B4BFF | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B4C77 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B4C91 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B4D05 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B4D92 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B4E31 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B4E4B | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B4EC3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B4EDD | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B4F59 | `SettingsMenus_ResetMainMenu_Screen9` | Known | Screen layout |
| 0x007B5029 | `SettingsMenus_ResetMusicMenu_Screen:` | Known | Screen layout |
| 0x007B50DD | `SettingsMenus_Main_Screen,` | Known | Screen layout |
| 0x007B514F | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5169 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B51E1 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B51FB | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B5536 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007B559C | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007B55F9 | `Extras_Screen` | Known | Screen layout |
| 0x007B564D | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x007B572B | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x007B5799 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007B5837 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x007B5850 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x007B58B8 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007B592B | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007B59AD | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007B5A0E | `NikePlus_Alert_Screen$` | Known | Screen layout |
| 0x007B5A8E | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007B5B07 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007B5B81 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007B5C06 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007B5C27 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007B5C96 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x007B5D1E | `NikePlus_Remote_Unlinking_Screen(` | Known | Screen layout |
| 0x007B5D42 | `NikePlus_Remote_Unlinking_Screen_Default` | Known | Screen layout |
| 0x007B5DB6 | `NikePlus_Alert_Screen(` | Known | Screen layout |
| 0x007B5E42 | `NikePlus_Dynamic_Workout_Screen,` | Known | Screen layout |
| 0x007B5E65 | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x007B5EDE | `NikePlus_Dynamic_Workout_Screen/` | Known | Screen layout |
| 0x007B5F01 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x007B5F7A | `NikePlus_Dynamic_Workout_Screen.` | Known | Screen layout |
| 0x007B5F9D | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x007B6016 | `NikePlus_Alert_Screen(` | Known | Screen layout |
| 0x007B6099 | `NikePlus_Custom_Screen,` | Known | Screen layout |
| 0x007B60B3 | `NikePlus_Custom_Screen_Weight_ToWorkoutMusic` | Known | Screen layout |
| 0x007B612D | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007B61AF | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007B6223 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007B6241 | `NikePlus_EndWorkout_Screen_Default` | Known | Screen layout |
| 0x007B62D9 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007B6355 | `NikePlus_EquipmentAlert_Screen<` | Known | Screen layout |
| 0x007B6422 | `NikePlus_EquipmentAlert_Screen<` | Known | Screen layout |
| 0x007B64EC | `NikePlus_EquipmentAlert_Screen8` | Known | Screen layout |
| 0x007B65B9 | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007B667A | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007B669B | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007B6732 | `NikePlus_Dynamic_Workout_Screen.` | Known | Screen layout |
| 0x007B6755 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x007B67F5 | `NikePlus_Dynamic_Workout_Screen/` | Known | Screen layout |
| 0x007B6818 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x007B68B6 | `NikePlus_Dynamic_Workout_Screen,` | Known | Screen layout |
| 0x007B68D9 | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x007B696F | `NikePlus_EndPausedWorkout_Screen1` | Known | Screen layout |
| 0x007B6993 | `NikePlus_EndPausedWorkout_Screen_BasicSave_Layout` | Known | Screen layout |
| 0x007B6A31 | `NikePlus_EndPausedWorkout_Screen4` | Known | Screen layout |
| 0x007B6A55 | `NikePlus_EndPausedWorkout_Screen_CaloriesSave_Layout` | Known | Screen layout |
| 0x007B6AF6 | `NikePlus_EndPausedWorkout_Screen4` | Known | Screen layout |
| 0x007B6B1A | `NikePlus_EndPausedWorkout_Screen_DistanceSave_Layout` | Known | Screen layout |
| 0x007B6BB8 | `NikePlus_EndPausedWorkout_Screen0` | Known | Screen layout |
| 0x007B6BDC | `NikePlus_EndPausedWorkout_Screen_TimeSave_Layout` | Known | Screen layout |
| 0x007B6C73 | `NikePlus_Songs_Screen` | Known | Screen layout |
| 0x007B6C8C | `NikePlus_Songs_Screen_Default` | Known | Screen layout |
| 0x007B6D9E | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007B6DB8 | `NikePlus_Custom_Screen_Weight` | Known | Screen layout |
| 0x007B6E1B | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007B6E8F | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007B6F0D | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x007B6F77 | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x007B6FD7 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007B7054 | `NikePlus_SensorSearching_Screen'` | Known | Screen layout |
| 0x007B7077 | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x007B70E6 | `NikePlus_Playlists_Screen ` | Known | Screen layout |
| 0x007B7103 | `NikePlus_Playlists_Screen_Nested` | Known | Screen layout |
| 0x007B7197 | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x007B71F7 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007B7274 | `NikePlus_SensorSearching_Screen'` | Known | Screen layout |
| 0x007B7297 | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x007B7306 | `NikePlus_Playlists_Screen ` | Known | Screen layout |
| 0x007B7323 | `NikePlus_Playlists_Screen_Nested` | Known | Screen layout |
| 0x007B73EC | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x007B744C | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007B74C9 | `NikePlus_Playlists_Screen!` | Known | Screen layout |
| 0x007B74E6 | `NikePlus_Playlists_Screen_Default` | Known | Screen layout |
| 0x007B7552 | `NikePlus_SensorSearching_Screen'` | Known | Screen layout |
| 0x007B7575 | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x007B7711 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007B772F | `NikePlus_NowRunning_Screen_Basic'` | Known | Screen layout |
| 0x007B77A3 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007B77C1 | `NikePlus_NowRunning_Screen_Calories'` | Known | Screen layout |
| 0x007B7838 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007B7856 | `NikePlus_NowRunning_Screen_Distance#` | Known | Screen layout |
| 0x007B78C9 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007B78E7 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007B79AF | `NikePlus_NowRunning_Screen$` | Known | Screen layout |
| 0x007B79CD | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x007B7A97 | `NikePlus_NowRunning_Screen$` | Known | Screen layout |
| 0x007B7AB5 | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x007B7B7F | `NikePlus_NowRunning_Screen$` | Known | Screen layout |
| 0x007B7B9D | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x007B7E62 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007B7E8E | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007B7F10 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007B7F3E | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007B7FC0 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007B7FE2 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007B8054 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007B8077 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007B80E7 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007B8105 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007B8175 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007B819B | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007B820F | `NikePlus_NowRunning_Screen_Basic_PowerSong` | Known | Screen layout |
| 0x007B859C | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007B85C8 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007B864A | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007B8678 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007B86FA | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007B871C | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007B878E | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007B87B1 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007B8821 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007B883F | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007B88AF | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007B88D5 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007B8946 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007B8CC6 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007B8CF2 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007B8D74 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007B8DA2 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007B8E24 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007B8E46 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007B8EB8 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007B8EDB | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007B8F4B | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007B8F69 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007B8FD9 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007B8FFF | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007B9073 | `NikePlus_NowRunning_Screen_Time_PowerSong` | Known | Screen layout |
| 0x007B9400 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007B942C | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007B94AE | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007B94DC | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007B955E | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007B9580 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007B95F2 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007B9615 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007B9685 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007B96A3 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007B9713 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007B9739 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007B97AA | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007B9B2A | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007B9B56 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007B9BD8 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007B9C06 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007B9C88 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007B9CAA | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007B9D1C | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007B9D3F | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007B9DAF | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007B9DCD | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007B9E3D | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007B9E63 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007B9ED7 | `NikePlus_NowRunning_Screen_Distance_PowerSong` | Known | Screen layout |
| 0x007BA268 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007BA294 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007BA316 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007BA344 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007BA3C6 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007BA3E8 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007BA45A | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007BA47D | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007BA4ED | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007BA50B | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007BA57B | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007BA5A1 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007BA612 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007BA996 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007BA9C2 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007BAA44 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007BAA72 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007BAAF4 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007BAB16 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007BAB88 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007BABAB | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007BAC1B | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007BAC39 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007BACA9 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007BACCF | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007BAD43 | `NikePlus_NowRunning_Screen_Calories_PowerSong` | Known | Screen layout |
| 0x007BB0D4 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007BB100 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007BB182 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007BB1B0 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007BB232 | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007BB254 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007BB2C6 | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007BB2E9 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007BB359 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007BB377 | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007BB3E7 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007BB40D | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007BB47E | `NikePlus_NowRunning_Screen_Calories` | Known | Screen layout |
| 0x007BB7CC | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007BB7F8 | `NikePlus_CalibrationCompleteError_Screen_Default%` | Known | Screen layout |
| 0x007BB87A | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007BB8A8 | `NikePlus_CalibrationCompleteSuccess_Screen_Default#` | Known | Screen layout |
| 0x007BB92A | `NikePlus_EndCalibration_Screen&` | Known | Screen layout |
| 0x007BB94C | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x007BB9BE | `NikePlus_ActivityStopped_Screen'` | Known | Screen layout |
| 0x007BB9E1 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x007BBA51 | `NikePlus_EndWorkout_Screen"` | Known | Screen layout |
| 0x007BBA6F | `NikePlus_EndWorkout_Screen_Default!` | Known | Screen layout |
| 0x007BBADF | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007BBB05 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007BBCA0 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007BBD07 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007BBD7B | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007BBDEE | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007BBE5A | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007BBE7B | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007BBEF6 | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007BBF1C | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007BBFD9 | `NikePlus_CalibrationCompleteError_Screen0` | Known | Screen layout |
| 0x007BC005 | `NikePlus_CalibrationCompleteError_Screen_Default'` | Known | Screen layout |
| 0x007BC089 | `NikePlus_CalibrationCompleteError_Screen*` | Known | Screen layout |
| 0x007BC0B5 | `NikePlus_CalibrationComplete_Screen_Pacing%` | Known | Screen layout |
| 0x007BC131 | `NikePlus_CalibrationCompleteSuccess_Screen2` | Known | Screen layout |
| 0x007BC15F | `NikePlus_CalibrationCompleteSuccess_Screen_Default` | Known | Screen layout |
| 0x007BC1D8 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007BC268 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007BC2BB | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x007BC328 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007BC37B | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x007BC3E8 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007BC43C | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007BC4A2 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007BC4C0 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007BC52C | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007BC54A | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007BC5BA | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007BC5D8 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007BC644 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007BC662 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007BC70F | `NikePlus_End_WorkoutSummary_Screen*` | Known | Screen layout |
| 0x007BC735 | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007BC7C8 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007BC7E2 | `NikePlus_Custom_Screen_Time` | Known | Screen layout |
| 0x007BC863 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007BC884 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007BC917 | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007BC931 | `NikePlus_Custom_Screen_Distance` | Known | Screen layout |
| 0x007BC9B9 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007BC9DA | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007BCA57 | `NikePlus_EquipmentAlert_Screen)` | Known | Screen layout |
| 0x007BCAF0 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007BCB11 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007BCB9C | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x007BCBB6 | `NikePlus_Custom_Screen_Calories` | Known | Screen layout |
| 0x007BCC69 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007BCCF0 | `NikePlus_Alert_Screen(` | Known | Screen layout |
| 0x007BCD99 | `NikePlus_EquipmentAlert_ScreenK` | Known | Screen layout |
| 0x007BCE5C | `NikePlus_EquipmentAlert_Screen8` | Known | Screen layout |
| 0x007BCF10 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007BD000 | `NikePlus_EquipmentAlert_Screen>` | Known | Screen layout |
| 0x007BD0BE | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007BD17D | `NikePlus_Alert_Screen$` | Known | Screen layout |
| 0x007BD1FE | `NikePlus_Remote_Unlinking_Screen(` | Known | Screen layout |
| 0x007BD222 | `NikePlus_Remote_Unlinking_Screen_Default!` | Known | Screen layout |
| 0x007BD298 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x007BD3D6 | `NikePlus_Calibration_CalibrateWalk_Screen1` | Known | Screen layout |
| 0x007BD47A | `NikePlus_Calibration_CalibrateRun_Screen0` | Known | Screen layout |
| 0x007BD53D | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007BD5FE | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007BD61F | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007BD6A6 | `NikePlus_Custom_Screen*` | Known | Screen layout |
| 0x007BD6C0 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x007BD7B1 | `NikePlus_Calibrate_ResetToDefault_Screen0` | Known | Screen layout |
| 0x007BD7DD | `NikePlus_Calibrate_ResetToDefault_Screen_Walking` | Known | Screen layout |
| 0x007BD893 | `NikePlus_Calibration_CalibrateRun_Screen"` | Known | Screen layout |
| 0x007BD90B | `NikePlus_Calibration_CalibrateWalk_Screen` | Known | Screen layout |
| 0x007BD9B7 | `NikePlus_Calibration_CalibrateRun_Screen"` | Known | Screen layout |
| 0x007BDA2F | `NikePlus_Calibration_CalibrateWalk_Screen` | Known | Screen layout |
| 0x007BDABD | `NikePlus_EquipmentAlert_Screen@` | Known | Screen layout |
| 0x007BDB7E | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007BDB9F | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007BDC26 | `NikePlus_Custom_Screen*` | Known | Screen layout |
| 0x007BDC40 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x007BDD2F | `NikePlus_Calibrate_ResetToDefault_Screen0` | Known | Screen layout |
| 0x007BDD5B | `NikePlus_Calibrate_ResetToDefault_Screen_Running` | Known | Screen layout |
| 0x007BDDCC | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007BDE1F | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007BDE94 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007BDEEE | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007BDFA5 | `NikePlus_Custom_Screen!` | Known | Screen layout |
| 0x007BE021 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x007BE098 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007BE12A | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x007BE1A8 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007BE202 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007BE2BA | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x007BE311 | `NikePlus_Calibration_ChooseCalibration_Screen5` | Known | Screen layout |
| 0x007BE3BE | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007BE43D | `NikePlus_StartCalibration_Screen(` | Known | Screen layout |
| 0x007BE461 | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x007BE4C7 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007BE543 | `NikePlus_StartWorkout_Screen$` | Known | Screen layout |
| 0x007BE563 | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x007BE5CE | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007BE647 | `NikePlus_Workout_Music_Screen` | Known | Screen layout |
| 0x007BE6AE | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x007BE709 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007BE7B5 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007BE847 | `NikePlus_StartWorkout_Screen$` | Known | Screen layout |
| 0x007BE867 | `NikePlus_StartWorkout_Screen_Default#` | Known | Screen layout |
| 0x007BE8DB | `NikePlus_StartCalibration_Screen(` | Known | Screen layout |
| 0x007BE8FF | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x007BE96C | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x007BE9F4 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x007BEA83 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007BEAA4 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007BEB41 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007BEB62 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007BEC01 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007BEC22 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007BECBD | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007BECDE | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007BEDAC | `NikePlus_EquipmentAlert_Screen)` | Known | Screen layout |
| 0x007BEE45 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007BEE66 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007BEF02 | `NikePlus_History_BestWorkouts_Screen,` | Known | Screen layout |
| 0x007BEF2A | `NikePlus_History_BestWorkouts_Screen_Default#` | Known | Screen layout |
| 0x007BEFA6 | `NikePlus_History_RecentWorkouts_Screen.` | Known | Screen layout |
| 0x007BEFD0 | `NikePlus_History_RecentWorkouts_Screen_Default'` | Known | Screen layout |
| 0x007BF052 | `NikePlus_History_WorkoutSummary_Screen+` | Known | Screen layout |
| 0x007BF07C | `NikePlus_History_WorkoutSummary_Screen_Last1` | Known | Screen layout |
| 0x007BF105 | `NikePlus_NoData_Screen%` | Known | Screen layout |
| 0x007BF11F | `NikePlus_NoData_Screen_NoBestWorkouts2` | Known | Screen layout |
| 0x007BF1A3 | `NikePlus_NoData_Screen&` | Known | Screen layout |
| 0x007BF1BD | `NikePlus_NoData_Screen_NoSavedWorkouts` | Known | Screen layout |
| 0x007BF2CD | `NikePlus_History_Totals_Screen&` | Known | Screen layout |
| 0x007BF2EF | `NikePlus_History_Totals_Screen_Default` | Known | Screen layout |
| 0x007BF361 | `NikePlus_History_DeleteActiveWorkout_Screen2` | Known | Screen layout |
| 0x007BF390 | `NikePlus_History_DeleteActiveWorkout_Screen_Cancel` | Known | Screen layout |
| 0x007BF405 | `NikePlus_History_DeleteActiveWorkout_Screen7` | Known | Screen layout |
| 0x007BF434 | `NikePlus_History_DeleteActiveWorkout_Screen_LastWorkout` | Known | Screen layout |
| 0x007BF4AC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007BF4FF | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007BF555 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007BF610 | `NikePlus_History_RecentWorkouts_ScreenB` | Known | Screen layout |
| 0x007BF6A6 | `NikePlus_History_WorkoutSummary_ScreenA` | Known | Screen layout |
| 0x007BF73B | `NikePlus_History_Screen` | Known | Screen layout |
| 0x007BF807 | `NikePlus_History_RecentWorkouts_ScreenB` | Known | Screen layout |
| 0x007BF89D | `NikePlus_History_WorkoutSummary_ScreenA` | Known | Screen layout |
| 0x007BF932 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x007BF9EF | `NikePlus_History_ScreenG` | Known | Screen layout |
| 0x007BFA7B | `NikePlus_History_RecentWorkouts_Screen` | Known | Screen layout |
| 0x007BFAFB | `NikePlus_History_DeleteAllWorkouts_Screen0` | Known | Screen layout |
| 0x007BFB28 | `NikePlus_History_DeleteAllWorkouts_Screen_Cancel#` | Known | Screen layout |
| 0x007BFBA8 | `NikePlus_History_WorkoutSummary_Screen.` | Known | Screen layout |
| 0x007BFBD2 | `NikePlus_History_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x007BFC85 | `NikePlus_History_ClearTotals_Screen+` | Known | Screen layout |
| 0x007BFCAC | `NikePlus_History_ClearTotals_Screen_Default` | Known | Screen layout |
| 0x007BFD4E | `NikePlus_History_Totals_Screen` | Known | Screen layout |
| 0x007BFDE1 | `NikePlus_Workout_Music_Screen%` | Known | Screen layout |
| 0x007BFE02 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x007BFE71 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007BFE8F | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007BFEFB | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007BFF19 | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007BFF89 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007BFFA7 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007C0013 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007C0031 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007C00C7 | `NikePlus_Dynamic_Workout_Screen/` | Known | Screen layout |
| 0x007C00EA | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x007C0163 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007C0181 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007C01ED | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007C020B | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007C027B | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007C0299 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007C0305 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007C0323 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007C03BB | `NikePlus_Dynamic_Workout_Screen,` | Known | Screen layout |
| 0x007C03DE | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x007C0454 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007C0472 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007C04DE | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007C04FC | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007C056C | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007C058A | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007C05F6 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007C0614 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007C06AB | `NikePlus_Dynamic_Workout_Screen.` | Known | Screen layout |
| 0x007C06CE | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x007C0746 | `NikePlus_NowRunning_Screen ` | Known | Screen layout |
| 0x007C0764 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x007C07D0 | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007C07EE | `NikePlus_NowRunning_Screen_Calories ` | Known | Screen layout |
| 0x007C085E | `NikePlus_NowRunning_Screen#` | Known | Screen layout |
| 0x007C087C | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x007C08E8 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x007C0906 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x007C0A19 | `VoiceMemos_Context_PlayDelete_Screen)` | Known | Screen layout |
| 0x007C0A41 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x007C0AB8 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007C0B84 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007C0BF3 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007C0CE1 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007C0D4A | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007C0D6C | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007C0DD8 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007C0DFA | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007C0F76 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C0F92 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007C1059 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007C1074 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007C10D7 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007C113A | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007C11D1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C11ED | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007C12B4 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007C12CF | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007C1332 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007C1395 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007C142D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007C1449 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007C1510 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007C152B | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007C158E | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007C15F1 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007C166E | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007C16D9 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007C1745 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007C17B7 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007C1824 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007C188F | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x007C18FB | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007C1963 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007C19CF | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007C1A43 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007C1AB1 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x007C1B2A | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x007E23E0 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x007E2465 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x007E2706 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x0098431B | `VoiceMemos_Screen_InsufficientDiskSpace2` | Known | Screen layout |
| 0x00985C67 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x00985C7F | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x00985C9D | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x00985D64 | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x00985DE2 | `NowPlaying_Screen_Music` | Known | Screen layout |
| 0x00985E23 | `MainMenu_Music_Screen_NoMusic` | Known | Screen layout |
| 0x00985E41 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00985E5F | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00985E78 | `NikePlus_Custom_Screen_Weight_ToWorkoutMusic` | Known | Screen layout |
| 0x00985F8D | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x00986016 | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x00986062 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x009862CB | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x009862E4 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x00986302 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x00986331 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x00986369 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x009863F9 | `NikePlus_Playlists_Screen_Nested` | Known | Screen layout |
| 0x009867D9 | `Video_Settings_TV_Screen_Standard` | Known | Screen layout |
| 0x0098680B | `Search_Main_Screen_WithKeyboard` | Known | Screen layout |
| 0x0098682B | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x00986870 | `VoiceMemos_Screen_InsufficientDiskSpace` | Known | Screen layout |
| 0x009868D5 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x009868F9 | `NikePlus_Custom_Screen_Distance` | Known | Screen layout |
| 0x00986954 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x009869DB | `Video_Settings_TV_Screen_Wide` | Known | Screen layout |
| 0x00986A23 | `CoverFlow_Screen_Backside` | Known | Screen layout |
| 0x0098A001 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x0098A206 | `NowPlaying_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x0098A22B | `Slideshow_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x0098A2FB | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x0098A315 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x0098A485 | `VoiceMemos_DeleteAll_ScreenTitle` | Known | Screen layout |
| 0x0098A522 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0098A565 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0098A657 | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x0098A677 | `NikePlus_Custom_Screen_Time` | Known | Screen layout |
| 0x0098A7C2 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0098A8AB | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x0098A8C4 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x0098A8D8 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x0098A8F5 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0098A914 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x0098AA1F | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x0098AB8B | `CoverFlow_Screen_BacksideForCapture` | Known | Screen layout |
| 0x0098BC89 | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x0098BD81 | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x0098BD9C | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x0098C03F | `NikePlus_CalibrationComplete_Screen_Pacing` | Known | Screen layout |
| 0x0098C0AB | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x0098C0F9 | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x0098C20B | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x0098C339 | `NikePlus_Calibrate_ResetToDefault_Screen_Walking` | Known | Screen layout |
| 0x0098C46B | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x0098C484 | `NikePlus_Calibrate_ResetToDefault_Screen_Running` | Known | Screen layout |
| 0x00991AA6 | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x00991B15 | `TVOutSlideshow_Screen_Playing` | Known | Screen layout |
| 0x00991B33 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00991B89 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x00991BF3 | `NikePlus_NowRunning_Screen_Basic_PowerSong` | Known | Screen layout |
| 0x00991C1E | `NikePlus_NowRunning_Screen_Distance_PowerSong` | Known | Screen layout |
| 0x00991C4C | `NikePlus_NowRunning_Screen_Time_PowerSong` | Known | Screen layout |
| 0x00991C99 | `NikePlus_NowRunning_Screen_Calories_PowerSong` | Known | Screen layout |
| 0x00991D16 | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x00991D81 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x00991EE3 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x00991F03 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x00992475 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x00992490 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x009924A3 | `MainMenu_List_ScreenLock` | Known | Screen layout |
| 0x009924BC | `ExtrasMenu_ScreenLock` | Known | Screen layout |
| 0x0099253F | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x00992560 | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x0099260B | `NikePlus_StartCalibration_Screen_Walk` | Known | Screen layout |
| 0x00992693 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x009926B5 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x009927BC | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x009927FC | `VoiceMemos_Screen_DeletAllAsk` | Known | Screen layout |
| 0x0099281A | `NowPlaying_Screen_TVOutAsk` | Known | Screen layout |
| 0x00992976 | `NowPlaying_Screen_Initial` | Known | Screen layout |
| 0x00992990 | `NowPlaying_Screen_Video_TVOut_Initial` | Known | Screen layout |
| 0x00992C62 | `NikePlus_History_DeleteAllWorkouts_Screen_Cancel` | Known | Screen layout |
| 0x00992C93 | `NikePlus_History_DeleteActiveWorkout_Screen_Cancel` | Known | Screen layout |
| 0x00993A68 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x00993AE9 | `RemoteUI_Screen` | Known | Screen layout |
| 0x00993AF9 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x00993B11 | `Video_Settings_TV_Screen` | Known | Screen layout |
| 0x00993B2A | `NikePlus_NoData_Screen` | Known | Screen layout |
| 0x00993B41 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00993B58 | `NikePlus_Workout_Music_Screen` | Known | Screen layout |
| 0x00993B76 | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x00993B9A | `AddressViewer_PartialLoad_Screen` | Known | Screen layout |
| 0x00993BBB | `NikePlus_ActivityStopped_Screen` | Known | Screen layout |
| 0x00993BDB | `Alarms_Alarm_Clock_Triggered_Screen` | Known | Screen layout |
| 0x00993BFF | `Alarms_Alarm_Triggered_Screen` | Known | Screen layout |
| 0x00993C1D | `Unsupported_Screen` | Known | Screen layout |
| 0x00993C30 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x00993C4E | `LockediPod_Screen` | Known | Screen layout |
| 0x00993C60 | `DiskMode_Screen` | Known | Screen layout |
| 0x00993C70 | `DemoMode_Screen` | Known | Screen layout |
| 0x00993C80 | `Notes_Image_Screen` | Known | Screen layout |
| 0x00993C93 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x00993CB1 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00993CC7 | `NowPlaying_Idle_Screen` | Known | Screen layout |
| 0x00993CDE | `Game_Screen` | Known | Screen layout |
| 0x00993CEA | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x00993D07 | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x00993D20 | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x00993D41 | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x00993D66 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x00993D79 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x00993D96 | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x00993DB7 | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x00993DDC | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x00993DF3 | `Notes_Loading_Screen` | Known | Screen layout |
| 0x00993E08 | `NikePlus_SensorSearching_Screen` | Known | Screen layout |
| 0x00993E28 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x00993E47 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x00993E5F | `NikePlus_Remote_Unlinking_Screen` | Known | Screen layout |
| 0x00993E80 | `SettingsMenus_AdjustScrolling_Screen` | Known | Screen layout |
| 0x00993EA5 | `Game_Running_Screen` | Known | Screen layout |
| 0x00993EB9 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x00993ED4 | `Stopwatch_Screen` | Known | Screen layout |
| 0x00993EE5 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x00993EFC | `Clock_Screen` | Known | Screen layout |
| 0x00993F09 | `NikePlus_Calibration_CalibrateWalk_Screen` | Known | Screen layout |
| 0x00993F33 | `TVOutSlideshowAsk_Screen` | Known | Screen layout |
| 0x00993F4C | `Settings_Legal_Screen` | Known | Screen layout |
| 0x00993F62 | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x00993F80 | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x00993F9C | `ToDo_Item_Screen` | Known | Screen layout |
| 0x00993FAD | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x00993FC4 | `DemoMode_Main_Screen` | Known | Screen layout |
| 0x00993FD9 | `Search_Main_Screen` | Known | Screen layout |
| 0x00993FEC | `AddressViewer_Main_Screen` | Known | Screen layout |
| 0x00994006 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x0099401B | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x00994031 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0099404B | `Clock_Region_Screen` | Known | Screen layout |
| 0x0099405F | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00994077 | `NikePlus_EndCalibration_Screen` | Known | Screen layout |
| 0x00994096 | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x009940C4 | `NikePlus_StartCalibration_Screen` | Known | Screen layout |
| 0x009940E5 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x00994103 | `NikePlus_Calibration_CalibrateRun_Screen` | Known | Screen layout |
| 0x0099412C | `Video_Settings_Fit_To_Screen` | Known | Screen layout |
| 0x00994149 | `Radio_Screen` | Known | Screen layout |
| 0x00994156 | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x00994170 | `ChargingMode_LowPower_Screen` | Known | Screen layout |
| 0x0099418D | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x009941A7 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x009941C1 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x009941DB | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x009941F4 | `NikePlus_CalibrationCompleteError_Screen` | Known | Screen layout |
| 0x0099421D | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x00994234 | `Extras_Screen` | Known | Screen layout |
| 0x00994242 | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x0099425F | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x00994281 | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x0099429A | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x009942B8 | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x009942D1 | `Video_Settings_Screen` | Known | Screen layout |
| 0x009942E7 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x00994300 | `SettingsMenus_AudiobookSettings_Screen` | Known | Screen layout |
| 0x00994327 | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x0099434D | `PhotosSettings_Screen` | Known | Screen layout |
| 0x00994363 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0099437B | `NikePlus_Songs_Screen` | Known | Screen layout |
| 0x00994391 | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x009943B4 | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x009943D1 | `NikePlus_History_Totals_Screen` | Known | Screen layout |
| 0x009943F0 | `NikePlus_History_ClearTotals_Screen` | Known | Screen layout |
| 0x00994414 | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x00994438 | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x00994451 | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x00994473 | `ToDo_Instructions_Screen` | Known | Screen layout |
| 0x0099448C | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x009944A8 | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x009944C2 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x009944E3 | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x009944FF | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00994517 | `VoiceMemos_Screen` | Known | Screen layout |
| 0x00994529 | `No_Photos_Screen` | Known | Screen layout |
| 0x0099453A | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x00994554 | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x00994570 | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x00994594 | `NikePlus_CalibrationCompleteSuccess_Screen` | Known | Screen layout |
| 0x009945BF | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x009945DF | `ChargingMode_Progress_Screen` | Known | Screen layout |
| 0x009945FC | `Notes_Contents_Screen` | Known | Screen layout |
| 0x00994612 | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x0099462D | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x00994649 | `NikePlus_Playlists_Screen` | Known | Screen layout |
| 0x00994663 | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x00994685 | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x009946A6 | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x009946C0 | `NikePlus_History_DeleteAllWorkouts_Screen` | Known | Screen layout |
| 0x009946EA | `NikePlus_History_RecentWorkouts_Screen` | Known | Screen layout |
| 0x00994711 | `NikePlus_History_BestWorkouts_Screen` | Known | Screen layout |
| 0x00994736 | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x00994750 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0099476F | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x00994790 | `NikePlus_Calibrate_ResetToDefault_Screen` | Known | Screen layout |
| 0x009947B9 | `NoContent_Screen` | Known | Screen layout |
| 0x009947CA | `Calendar_Event_Screen` | Known | Screen layout |
| 0x009947E0 | `FirstBoot_Screen` | Known | Screen layout |
| 0x009947F1 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x00994807 | `NikePlus_EquipmentAlert_Screen` | Known | Screen layout |
| 0x00994826 | `Notes_List_Screen` | Known | Screen layout |
| 0x00994838 | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x00994859 | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x00994873 | `NikePlus_Dynamic_Workout_Screen` | Known | Screen layout |
| 0x00994893 | `NikePlus_EndPausedWorkout_Screen` | Known | Screen layout |
| 0x009948B4 | `NikePlus_EndWorkout_Screen` | Known | Screen layout |
| 0x009948CF | `NikePlus_ResumeWorkout_Screen` | Known | Screen layout |
| 0x009948ED | `NikePlus_History_DeleteActiveWorkout_Screen` | Known | Screen layout |
| 0x00994919 | `NikePlus_StartWorkout_Screen` | Known | Screen layout |
| 0x00994936 | `TODOS_Menu_Screen` | Known | Screen layout |
| 0x00994948 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x0099495E | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0099497A | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0099498F | `Games_Menu_Screen` | Known | Screen layout |
| 0x009949A1 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x009949B4 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x009949D3 | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x009949F2 | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x00994A16 | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x00994A34 | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x00994A57 | `TVOutSlideshow_Screen` | Known | Screen layout |
| 0x00994A6D | `CoverFlow_Screen` | Known | Screen layout |
| 0x00994A7E | `Calendar_Day_Screen` | Known | Screen layout |
| 0x00994A92 | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x00994AB4 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x00994ACC | `Stopwatch_SessionSummary_Screen` | Known | Screen layout |
| 0x00994AEC | `NikePlus_End_WorkoutSummary_Screen` | Known | Screen layout |
| 0x00994B0F | `NikePlus_History_WorkoutSummary_Screen` | Known | Screen layout |
| 0x00994B36 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x00994B4E | `Clock_NorthAmerica_City_Screen` | Known | Screen layout |
| 0x00994B6D | `Clock_SouthAmerica_City_Screen` | Known | Screen layout |
| 0x00994B8C | `Clock_Africa_City_Screen` | Known | Screen layout |
| 0x00994BA5 | `Clock_Australia_City_Screen` | Known | Screen layout |
| 0x00994BC1 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x00994BD8 | `Clock_Pacific_City_Screen` | Known | Screen layout |
| 0x00994BF2 | `Clock_Atlantic_City_Screen` | Known | Screen layout |
| 0x00994C0D | `Clock_Europe_City_Screen` | Known | Screen layout |
| 0x00994D01 | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x00994D52 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00994D75 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00994D9D | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x00995129 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x009952E0 | `NikePlus_StartCalibration_Screen_Run` | Known | Screen layout |
| 0x0099556B | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x009955C1 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x009956F5 | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x00995712 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x00995B02 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x00995C18 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x00995C3A | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00995D72 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x00995D91 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x0099644B | `ChargingMode_ScreenLayout_LowPower` | Known | Screen layout |
| 0x00996DAD | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x00996EDD | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x00996F94 | `NikePlus_NowRunning_Screen_Calories` | Known | Screen layout |
| 0x00996FB8 | `NikePlus_Custom_Screen_Calories` | Known | Screen layout |
| 0x00997051 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0099706F | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0099708F | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00997163 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x0099717F | `Extras_Screen_Games` | Known | Screen layout |
| 0x00997285 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x009972A4 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x009972C0 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x009973AB | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x009974C7 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00997695 | `MainMenu_Music_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009976B8 | `MainMenus_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009976DB | `NoContent_Screen_NoAudiobooks` | Known | Screen layout |
| 0x009977D5 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x009977F2 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x00997871 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00997955 | `MainMenu_Music_Screen_NoCompilations` | Known | Screen layout |
| 0x0099797A | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00997AE4 | `MainMenu_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00997B07 | `MainMenu_Videos_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00997B2C | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00997B4B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00997B6A | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00997B8B | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x00997BC9 | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x00997BEA | `Extras_Screen_VoiceMemos` | Known | Screen layout |
| 0x00997C55 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00997C87 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00997CA6 | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x00997D53 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x00997DBF | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00997EB8 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x00997ED4 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x00997F57 | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x00997F72 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x00997F93 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x00998042 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x00998076 | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x00998097 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00998155 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00998176 | `MainMenu_Videos_Screen_NoPlaylists` | Known | Screen layout |
| 0x00998199 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x009981E8 | `MainMenu_Music_Screen_NoArtists` | Known | Screen layout |
| 0x00998258 | `NikePlus_NoData_Screen_NoSavedWorkouts` | Known | Screen layout |
| 0x0099830A | `NikePlus_NoData_Screen_NoBestWorkouts` | Known | Screen layout |
| 0x009983B7 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x009983D6 | `Extras_Screen_NikePlus` | Known | Screen layout |
| 0x00998526 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00998545 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00998566 | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x009989EE | `NikePlus_Custom_Screen_Weight` | Known | Screen layout |
| 0x00998A63 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x00998B16 | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x00998B90 | `Notes_Contents_Screen_Alt` | Known | Screen layout |
| 0x00998BAA | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x00998C2A | `VoiceMemos_Screen_InsufficientDiskSpace2_Default` | Known | Screen layout |
| 0x00998CDC | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x00998D81 | `VoiceMemos_Screen_InsufficientDiskSpace_Default` | Known | Screen layout |
| 0x00998DB1 | `NowPlaying_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x00998DDE | `Slideshow_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x00999C6C | `Radio_Screen_Tuning_Default` | Known | Screen layout |
| 0x00999CF8 | `VoiceMemos_Screen_DeletAllAsk_Default` | Known | Screen layout |
| 0x00999D1E | `NowPlaying_Screen_TVOutAsk_Default` | Known | Screen layout |
| 0x00999D55 | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x00999D7B | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00999D99 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x00999DC5 | `AddressViewer_PartialLoad_Screen_Default` | Known | Screen layout |
| 0x00999DEE | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x00999E16 | `Alarms_Alarm_Clock_Triggered_Screen_Default` | Known | Screen layout |
| 0x00999E42 | `Alarms_Alarm_Triggered_Screen_Default` | Known | Screen layout |
| 0x00999E68 | `Unsupported_Screen_Default` | Known | Screen layout |
| 0x00999E83 | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x00999EA9 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x00999EC1 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x00999EDC | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00999EF9 | `Game_Screen_Default` | Known | Screen layout |
| 0x00999F0D | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x00999F33 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x00999F54 | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x00999F7D | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x00999FA7 | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x00999FD4 | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x00999FFD | `Notes_Loading_Screen_Default` | Known | Screen layout |
| 0x0099A01A | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x0099A042 | `NikePlus_Remote_Unlinking_Screen_Default` | Known | Screen layout |
| 0x0099A06B | `Clock_Screen_Default` | Known | Screen layout |
| 0x0099A080 | `TVOutSlideshowAsk_Screen_Default` | Known | Screen layout |
| 0x0099A0A1 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x0099A0BF | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x0099A0E5 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x0099A109 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x0099A122 | `AddressViewer_Main_Screen_Default` | Known | Screen layout |
| 0x0099A144 | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x0099A161 | `MainMenus_Main_Screen_Default` | Known | Screen layout |
| 0x0099A17F | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0099A19C | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x0099A1B8 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x0099A1E4 | `NikePlus_EndCalibration_Screen_Default` | Known | Screen layout |
| 0x0099A20B | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x0099A234 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0099A249 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0099A26B | `NikePlus_CalibrationCompleteError_Screen_Default` | Known | Screen layout |
| 0x0099A29C | `Extras_Screen_Default` | Known | Screen layout |
| 0x0099A2B2 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x0099A2D8 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0099A2F9 | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x0099A317 | `NikePlus_Settings_Screen_Default` | Known | Screen layout |
| 0x0099A338 | `NikePlus_Songs_Screen_Default` | Known | Screen layout |
| 0x0099A356 | `NikePlus_History_Totals_Screen_Default` | Known | Screen layout |
| 0x0099A37D | `NikePlus_History_ClearTotals_Screen_Default` | Known | Screen layout |
| 0x0099A3A9 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x0099A3D5 | `ToDo_Instructions_Screen_Default` | Known | Screen layout |
| 0x0099A3F6 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x0099A41A | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x0099A43C | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x0099A460 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0099A47F | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0099A498 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x0099A4BA | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0099A4DE | `NikePlus_CalibrationCompleteSuccess_Screen_Default` | Known | Screen layout |
| 0x0099A511 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x0099A52F | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0099A553 | `NikePlus_Playlists_Screen_Default` | Known | Screen layout |
| 0x0099A575 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x0099A59F | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0099A5C8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0099A5EA | `NikePlus_History_RecentWorkouts_Screen_Default` | Known | Screen layout |
| 0x0099A619 | `NikePlus_History_BestWorkouts_Screen_Default` | Known | Screen layout |
| 0x0099A646 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0099A664 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0099A67D | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x0099A697 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x0099A6C0 | `NikePlus_EndWorkout_Screen_Default` | Known | Screen layout |
| 0x0099A6E3 | `NikePlus_ResumeWorkout_Screen_Default` | Known | Screen layout |
| 0x0099A709 | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x0099A72E | `TODOS_Menu_Screen_Default` | Known | Screen layout |
| 0x0099A748 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x0099A766 | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x0099A783 | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x0099A79D | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x0099A7B8 | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x0099A7D7 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x0099A7F5 | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x0099A80E | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x0099A82A | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x0099A854 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x0099A874 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x0099A89C | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x0099A8C7 | `NikePlus_History_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x0099A8F6 | `NikePlus_History_Screen_Default` | Known | Screen layout |
| 0x0099A916 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0099A93D | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0099A964 | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x0099A985 | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x0099A9A9 | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x0099A9C8 | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x0099A9EA | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x0099AA0D | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x0099AA4A | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0099AAD8 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x0099AAFA | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x0099AB39 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x0099B256 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0099B282 | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0099B2C7 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x0099B2EF | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0099B310 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0099B331 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0099B357 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0099B374 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0099B396 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x0099B3BA | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0099B3DE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0099B443 | `NikePlus_History_WorkoutSummary_Screen_Last` | Known | Screen layout |
| 0x0099B5A2 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0099B612 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0099B663 | `NowPlaying_Screen_Video_TVOut` | Known | Screen layout |
| 0x0099B776 | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x0099B7D3 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x0099B822 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x0099B8E9 | `NikePlus_History_DeleteActiveWorkout_Screen_LastWorkout` | Known | Screen layout |
| 0x0099BA06 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x0099BE1F | `NikePlus_EndPausedWorkout_Screen_BasicSave_Layout` | Known | Screen layout |
| 0x0099BE51 | `NikePlus_EndPausedWorkout_Screen_DistanceSave_Layout` | Known | Screen layout |
| 0x0099BE86 | `NikePlus_EndPausedWorkout_Screen_TimeSave_Layout` | Known | Screen layout |
| 0x0099BEB7 | `NikePlus_EndPausedWorkout_Screen_CaloriesSave_Layout` | Known | Screen layout |
| 0x0099C16F | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x0099C28C | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0099C530 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x0099C84D | `Lock_Screen_Default_Layout` | Known | Screen layout |
| 0x0099C8B7 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x0099CAC6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0099CB70 | `SettingsMenu_About_Screen_Accessory_Layout` | Known | Screen layout |
| 0x0099CBC3 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0099F6C5 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x0099F711 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x0099F7EF | `MainMenu_List_ScreenLock_x` | Known | Screen layout |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00008FD7 | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | Known | RTOS |
| 0x002B4180 | `  K - RTXC` | Known | RTOS |
| 0x002B5168 | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | Known | RTOS |
| 0x00983210 | `Returning from RTXCBug` | Known | RTOS |

---

## 9. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000DEF28 | `HostOSTask` | Known | RTOS task thread |
| 0x0013D530 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x00145600 | `USBDeviceTask` | Known | RTOS task thread |
| 0x0014F94C | `DiskReaderTask` | Known | RTOS task thread |
| 0x0015F2F0 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0015F304 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x001B2B50 | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x001F21AC | `MeCCAIOTask` | Known | RTOS task thread |
| 0x00227C5C | `StreamCacheMassStorageManagerTimeOutTask` | Known | RTOS task thread |
| 0x00227DD8 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x002A7BD4 | `FirewireTask` | Known | RTOS task thread |
| 0x002A7BE8 | `TouchwheelTask` | Known | RTOS task thread |
| 0x002A7BFC | `AudioOutStateTask` | Known | RTOS task thread |
| 0x002A7C28 | `DiskMgrTask` | Known | RTOS task thread |
| 0x002A7C38 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x002A7C4C | `TopPlugTask` | Known | RTOS task thread |
| 0x002A7C5C | `HPhoneDetTask` | Known | RTOS task thread |
| 0x002A7CD4 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x002A7CFC | `AlarmTask` | Known | RTOS task thread |
| 0x002A7D1B | `"USBAudioTask` | Known | RTOS task thread |
| 0x002B4820 | `Undefined Task` | Known | RTOS task thread |
| 0x003AF2FC | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x003B31CC | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x003BB5B8 | `MeCCARecordingTask` | Known | RTOS task thread |
| 0x008D70FC | `WaveFileDebugTask` | Known | RTOS task thread |

---

## 10. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0026042C | `Channel Reserved` | Known | Logging channel |
| 0x00260440 | `Channel AppBoot` | Known | Logging channel |
| 0x00260450 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x0026046C | `Channel PrefsWriting` | Known | Logging channel |
| 0x00260484 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x002604A4 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x002604BC | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x002604D8 | `Channel TestLogging` | Known | Logging channel |
| 0x002604EC | `Channel AppFileLoading` | Known | Logging channel |
| 0x00260504 | `Channel VCardReading` | Known | Logging channel |
| 0x0026051C | `Channel LongSongScanning` | Known | Logging channel |
| 0x00260590 | `Channel VoiceRecording` | Known | Logging channel |
| 0x002605A8 | `Channel PhotoImporting` | Known | Logging channel |
| 0x002605C0 | `Channel Notes` | Known | Logging channel |
| 0x002605D0 | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x002605EC | `Channel DiskMode` | Known | Logging channel |
| 0x00260600 | `Channel Firewire` | Known | Logging channel |
| 0x00260614 | `Channel USB` | Known | Logging channel |
| 0x00260634 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x0026064C | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0008E4E0 | `gamedata_RW` | Known | Game system |
| 0x0008E4FC | `gamedata_ShareRW` | Known | Game system |
| 0x0008E510 | `games_RO` | Known | Game system |
| 0x0098326A | `iPod_Control/games_RO/` | Known | Game system |
| 0x00983281 | `Resources/Games/games_RO/` | Known | Game system |
| 0x0098F591 | `GamesMenu_StatusBar_String` | Known | Game system |
| 0x0098FC2C | `AboutScreen_Games_String` | Known | Game system |
| 0x00997193 | `MainMenu_List_Games` | Known | Game system |
| 0x009971A7 | `ExtrasMenu_Games` | Known | Game system |
| 0x0099F85E | `MainMenu_List_Games_x` | Known | Game system |

---

## 12. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0009F5D0 | `adrmmp4a` | Known | DRM system |
| 0x0014D0B8 | `AppleDRMVersion` | Known | DRM system |
| 0x0014D158 | `AppleDRM` | Known | DRM system |
| 0x0014E274 | `AppleVideoDRM` | Known | DRM system |
| 0x00151670 | `drmsp608aavdmp4aesds4K` | Known | DRM system |
| 0x0098364F | `DRMLevel` | Known | DRM system |

---

## 13. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00035A64 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00035A7C | `iPod_Control\iTunes\firsttime` | Known | iTunes database |
| 0x00058F24 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x00058F4C | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00062328 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0008A6BC | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x0008E474 | `iPod_Control/iTunes/iTunesDB.p7b` | Known | iTunes database |
| 0x000AB564 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x000AB738 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000B4170 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000B55B4 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000B56B4 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00135D90 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x003A867C | `iTunesDB` | Known | iTunes database |
| 0x003A8688 | `iPod_Control\iTunes\` | Known | iTunes database |

---

## 14. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0005FC74 | `NANDDRIVERSIGN` | Known | Hardware |
| 0x000607B0 | `NANDDRIVERSIGN` | Known | Hardware |
| 0x00061D68 | `[FTL:MSG] Apple NAND Driver (AND) 0x%08x` | Known | Hardware |
| 0x00061E80 | `[FTL:MSG] Valid Signature not found! Re-initializing NAND!` | Known | Hardware |
| 0x001355F4 | `NAND FLASH DRIVE` | Known | Hardware |
| 0x0014D5F8 | `FireWireGUID` | Known | FireWire |
| 0x0014D608 | `FireWireVersion` | Known | FireWire |
| 0x0014DC3C | `FireWire` | Known | FireWire |
| 0x002C5010 | `[FIL:ERR] No recognized NAND found (0x%X, 0x%X) (line:%d)!` | Known | Hardware |
| 0x008DEEB8 | `[FTL:WRN] Recovering NAND Data Structures - this will take some time!` | Known | Hardware |
| 0x008E03D0 | `[FIL:WRN]  FNAND_GetStruct 0x%X is not identified is FIL data struct identifier!` | Known | Hardware |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0073A960 | `Settings_RadioRegionsMenu_Layout"` | Known | FM Radio |
| 0x0073A9E9 | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x007E06CC | `Radio Regions` | Known | FM Radio |
| 0x00835760 | `Radio-Regionen` | Known | FM Radio |
| 0x0098CB71 | `Settings_RadioRegions_Australia_String` | Known | FM Radio |
| 0x0098CB98 | `Settings_RadioRegions_Asia_String` | Known | FM Radio |
| 0x0098DCD9 | `Settings_RadioRegions_Europe_String` | Known | FM Radio |
| 0x0098EFDB | `Settings_RadioRegions_Japan_String` | Known | FM Radio |
| 0x0098FA49 | `Settings_RadioRegions_Americas_String` | Known | FM Radio |
| 0x0099008B | `SettingsMenu_RadioRegions_String` | Known | FM Radio |
| 0x00993927 | `RadioRegionsMenuItem` | Known | FM Radio |
| 0x009978DE | `SettingsMenu_ListItem_RadioRegions` | Known | FM Radio |
| 0x0099C395 | `SettingsInfo_Template_RadioRegions_Layout` | Known | FM Radio |
| 0x0099C3BF | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x0099CA87 | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 16. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00870024 | `Fotocamera` | Known | Camera |
| 0x008701C4 | `Collega fotocamera o inserisci scheda` | Known | Camera |
| 0x0087023C | `Fotocamera non supportata` | Known | Camera |
| 0x0088C92C | `Camera` | Known | Camera |
| 0x0088CAC8 | `Sluit camera of kaart aan` | Known | Camera |
| 0x0088CB34 | `Camera niet ondersteund` | Known | Camera |
| 0x0098CBBA | `Photos_DevName_Camera_String` | Known | Camera |

---

## 17. Pedometer

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x007E1678 | `Step away from all other sensors.` | Known | Pedometer |
| 0x007E185C | `Step away from all other remotes.` | Known | Pedometer |
| 0x0099FAFC | `NikePlus_Remote_Step_Away` | Known | Pedometer |
| 0x0099FB16 | `NikePlus_Step_Away` | Known | Pedometer |

---

## 18. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00035A50 | `iPod_Control` | Filesystem Path |  |
| 0x00035ABC | `iPod_Control\Device` | Filesystem Path |  |
| 0x00045810 | `iPod_Control\Device` | Filesystem Path |  |
| 0x00047914 | `iPod_Control` | Filesystem Path |  |
| 0x00047F9C | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x00058F04 | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path |  |
| 0x0005BE9C | `iPod_Control\Music\` | Filesystem Path |  |
| 0x000621A8 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x000986B4 | `iPod_Control` | Filesystem Path |  |
| 0x000986C4 | `Resources/Games` | Filesystem Path |  |
| 0x000986D4 | `iPod_Control/%s%s%s` | Filesystem Path |  |
| 0x00102128 | `iPod_Control/Device/alarms` | Filesystem Path |  |
| 0x001120F4 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00113454 | `iPod_Control/Device` | Filesystem Path |  |
| 0x00113468 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x00130B34 | `iPod_Control/Device/Users` | Filesystem Path |  |
| 0x001609D4 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x00160C30 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x0016C910 | `Resources/UI/active.bin` | Filesystem Path |  |
| 0x0016C928 | `Resources/UI/` | Filesystem Path |  |
| 0x0018D494 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x0018E778 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path |  |
| 0x0018E7A0 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001B6508 | `iPod_Control/Device/PlayCounts` | Filesystem Path |  |
| 0x001CC72C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CC7DC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CC958 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CCAF0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CCB98 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CCD40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CCDE4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CCE88 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CCF2C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CCFD0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD074 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD118 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD1C8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD278 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD328 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD494 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD544 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD5F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD698 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD748 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD83C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD8E0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CD994 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CDA50 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CDB00 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CDC24 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CDCE0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CDD90 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CDF4C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE010 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE0C0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE17C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE2B8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE384 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE440 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE4E4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE588 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE644 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE700 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE7C8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE86C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE934 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CE9E4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CEAAC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CEB74 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CEC24 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CECD4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CED98 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CEE48 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CEEF8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CEFA8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CF07C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CF150 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CF250 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CF330 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CF438 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001CF524 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0027192C | `Resources/Fonts` | Filesystem Path |  |
| 0x0028AB90 | `Resources/Fonts` | Filesystem Path |  |
| 0x003A86FA | `iPod_Control/Device` | Filesystem Path |  |
| 0x003AEB9C | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x003B16F0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003B19A6 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003B1A64 | `iPod_Control/Device/Trainer/Workouts/Empeds/` | Filesystem Path |  |
| 0x003BB584 | `iPod_Control:Device:voices:` | Filesystem Path |  |
| 0x003BCA36 | `Resources/TrainerTemplates` | Filesystem Path |  |
| 0x003BCA51 | `iPod_Control/Device/Trainer/TrainerTemplates` | Filesystem Path |  |
| 0x003BD030 | `iPod_Control/Device/Trainer/Workouts/Empeds/` | Filesystem Path |  |
| 0x00983145 | `Resources/Games/` | Filesystem Path |  |
| 0x00983531 | `iPod_Control/Device` | Filesystem Path |  |
| 0x00983545 | `iPod_Control/Device/accessories.reg` | Filesystem Path |  |
| 0x009835C6 | `iPod_Control/Tones` | Filesystem Path |  |

---

## 19. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x008D56A4 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x008DAC40 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftu` | Build Path |  |
| 0x008DAC98 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fts` | Build Path |  |
| 0x008DACF0 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fto` | Build Path |  |
| 0x008DE850 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x008DE8C4 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x008DF4E0 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x008DFC14 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x008E0034 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Shared\Services\hwapi\soc\samsung\nan` | Build Path |  |
| 0x008E6EF4 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\` | Build Path |  |
| 0x008E7A70 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfd` | Build Path |  |
| 0x008E8C6C | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrg` | Build Path |  |
| 0x008E8CC4 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrc` | Build Path |  |
| 0x008E8D1C | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfro` | Build Path |  |
| 0x008E9060 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1` | Build Path |  |
| 0x008F8408 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttc` | Build Path |  |
| 0x008F8684 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype` | Build Path |  |
| 0x008F8BF0 | `c:\bwa\N46FirmwareWin-313\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1` | Build Path |  |

---

## 20. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00096680 | `Acoustic` | EQ Preset |  |
| 0x0009668C | `Bass Booster` | EQ Preset |  |
| 0x000966AC | `Classical` | EQ Preset |  |
| 0x000966B8 | `Dance` | EQ Preset |  |
| 0x000966C8 | `Electronic` | EQ Preset |  |
| 0x000966DC | `Hip Hop` | EQ Preset |  |
| 0x000966E4 | `Jazz` | EQ Preset |  |
| 0x000966EC | `Latin` | EQ Preset |  |
| 0x000966F4 | `Loudness` | EQ Preset |  |
| 0x00096700 | `Lounge` | EQ Preset |  |
| 0x00096708 | `Piano` | EQ Preset |  |
| 0x0009671C | `Rock` | EQ Preset |  |
| 0x00096724 | `Small Speakers` | EQ Preset |  |
| 0x00096734 | `Spoken Word` | EQ Preset |  |
| 0x00096740 | `Treble Booster` | EQ Preset |  |
| 0x0009678C | `Vocal Booster` | EQ Preset |  |
| 0x007E09BC | `Acoustic` | EQ Preset |  |
| 0x007E09C8 | `Bass Booster` | EQ Preset |  |
| 0x007E09E8 | `Classical` | EQ Preset |  |
| 0x007E09F4 | `Dance` | EQ Preset |  |
| 0x007E0A04 | `Electronic` | EQ Preset |  |
| 0x007E0A18 | `Hip Hop` | EQ Preset |  |
| 0x007E0A20 | `Jazz` | EQ Preset |  |
| 0x007E0A28 | `Latin` | EQ Preset |  |
| 0x007E0A30 | `Loudness` | EQ Preset |  |
| 0x007E0A3C | `Lounge` | EQ Preset |  |
| 0x007E0A44 | `Piano` | EQ Preset |  |
| 0x007E0A54 | `Rock` | EQ Preset |  |
| 0x007E0A5C | `Small Speakers` | EQ Preset |  |
| 0x007E0A6C | `Spoken Word` | EQ Preset |  |
| 0x007E0A78 | `Treble Booster` | EQ Preset |  |
| 0x007E0A98 | `Vocal Booster` | EQ Preset |  |
| 0x00824828 | `Acoustic` | EQ Preset |  |
| 0x00824834 | `Bass Booster` | EQ Preset |  |
| 0x00824854 | `Classical` | EQ Preset |  |
| 0x00824860 | `Dance` | EQ Preset |  |
| 0x00824870 | `Electronic` | EQ Preset |  |
| 0x00824884 | `Hip Hop` | EQ Preset |  |
| 0x0082488C | `Jazz` | EQ Preset |  |
| 0x00824894 | `Latin` | EQ Preset |  |
| 0x0082489C | `Loudness` | EQ Preset |  |
| 0x008248A8 | `Lounge` | EQ Preset |  |
| 0x008248B0 | `Piano` | EQ Preset |  |
| 0x008248C0 | `Rock` | EQ Preset |  |
| 0x008248C8 | `Small Speakers` | EQ Preset |  |
| 0x008248D8 | `Spoken Word` | EQ Preset |  |
| 0x008248E4 | `Treble Booster` | EQ Preset |  |
| 0x00824904 | `Vocal Booster` | EQ Preset |  |
| 0x0082CFF8 | `Acoustic` | EQ Preset |  |
| 0x0082D004 | `Bass Booster` | EQ Preset |  |
| 0x0082D024 | `Classical` | EQ Preset |  |
| 0x0082D030 | `Dance` | EQ Preset |  |
| 0x0082D040 | `Electronic` | EQ Preset |  |
| 0x0082D054 | `Hip Hop` | EQ Preset |  |
| 0x0082D05C | `Jazz` | EQ Preset |  |
| 0x0082D064 | `Latin` | EQ Preset |  |
| 0x0082D06C | `Loudness` | EQ Preset |  |
| 0x0082D078 | `Lounge` | EQ Preset |  |
| 0x0082D080 | `Piano` | EQ Preset |  |
| 0x0082D090 | `Rock` | EQ Preset |  |
| 0x0082D098 | `Small Speakers` | EQ Preset |  |
| 0x0082D0A8 | `Spoken Word` | EQ Preset |  |
| 0x0082D0B4 | `Treble Booster` | EQ Preset |  |
| 0x0082D0D4 | `Vocal Booster` | EQ Preset |  |
| 0x00835B08 | `Acoustic` | EQ Preset |  |
| 0x00835B38 | `Dance` | EQ Preset |  |
| 0x00835B48 | `Electronic` | EQ Preset |  |
| 0x00835B64 | `Jazz` | EQ Preset |  |
| 0x00835B6C | `Latin` | EQ Preset |  |
| 0x00835B74 | `Loudness` | EQ Preset |  |
| 0x00835B88 | `Piano` | EQ Preset |  |
| 0x00835B98 | `Rock` | EQ Preset |  |
| 0x0084B874 | `Dance` | EQ Preset |  |
| 0x0084B89C | `Hip Hop` | EQ Preset |  |
| 0x0084B8A4 | `Jazz` | EQ Preset |  |
| 0x0084B8B4 | `Loudness` | EQ Preset |  |
| 0x0084B8C0 | `Lounge` | EQ Preset |  |
| 0x0084B8C8 | `Piano` | EQ Preset |  |
| 0x0084B8D8 | `Rock` | EQ Preset |  |
| 0x00853FC8 | `Jazz` | EQ Preset |  |
| 0x00853FD0 | `Latin` | EQ Preset |  |
| 0x00853FE4 | `Lounge` | EQ Preset |  |
| 0x00853FEC | `Piano` | EQ Preset |  |
| 0x00853FFC | `Rock` | EQ Preset |  |
| 0x0085C668 | `Hip Hop` | EQ Preset |  |
| 0x0085C670 | `Jazz` | EQ Preset |  |
| 0x0085C68C | `Lounge` | EQ Preset |  |
| 0x0085C694 | `Piano` | EQ Preset |  |
| 0x0085C6AC | `Rock` | EQ Preset |  |
| 0x008658A4 | `Latin` | EQ Preset |  |
| 0x008658D0 | `Rock` | EQ Preset |  |
| 0x0086E474 | `Dance` | EQ Preset |  |
| 0x0086E498 | `Hip Hop` | EQ Preset |  |
| 0x0086E4A0 | `Jazz` | EQ Preset |  |
| 0x0086E4B0 | `Loudness` | EQ Preset |  |
| 0x0086E4BC | `Lounge` | EQ Preset |  |
| 0x0086E4C4 | `Piano` | EQ Preset |  |
| 0x0086E4D4 | `Rock` | EQ Preset |  |
| 0x00878188 | `Acoustic` | EQ Preset |  |
| 0x00878194 | `Bass Booster` | EQ Preset |  |
| 0x008781B4 | `Classical` | EQ Preset |  |
| 0x008781C0 | `Dance` | EQ Preset |  |
| 0x008781D0 | `Electronic` | EQ Preset |  |
| 0x008781E4 | `Hip Hop` | EQ Preset |  |
| 0x008781EC | `Jazz` | EQ Preset |  |
| 0x008781F4 | `Latin` | EQ Preset |  |
| 0x008781FC | `Loudness` | EQ Preset |  |
| 0x00878208 | `Lounge` | EQ Preset |  |
| 0x00878210 | `Piano` | EQ Preset |  |
| 0x00878220 | `Rock` | EQ Preset |  |
| 0x00878228 | `Small Speakers` | EQ Preset |  |
| 0x00878238 | `Spoken Word` | EQ Preset |  |
| 0x00878244 | `Treble Booster` | EQ Preset |  |
| 0x00878264 | `Vocal Booster` | EQ Preset |  |
| 0x00881EA8 | `Acoustic` | EQ Preset |  |
| 0x00881EB4 | `Bass Booster` | EQ Preset |  |
| 0x00881ED4 | `Classical` | EQ Preset |  |
| 0x00881EE0 | `Dance` | EQ Preset |  |
| 0x00881EF0 | `Electronic` | EQ Preset |  |
| 0x00881F04 | `Hip Hop` | EQ Preset |  |
| 0x00881F0C | `Jazz` | EQ Preset |  |
| 0x00881F14 | `Latin` | EQ Preset |  |
| 0x00881F1C | `Loudness` | EQ Preset |  |
| 0x00881F28 | `Lounge` | EQ Preset |  |
| 0x00881F30 | `Piano` | EQ Preset |  |
| 0x00881F40 | `Rock` | EQ Preset |  |
| 0x00881F48 | `Small Speakers` | EQ Preset |  |
| 0x00881F58 | `Spoken Word` | EQ Preset |  |
| 0x00881F64 | `Treble Booster` | EQ Preset |  |
| 0x00881F84 | `Vocal Booster` | EQ Preset |  |
| 0x0088AD24 | `Dance` | EQ Preset |  |
| 0x0088AD58 | `Jazz` | EQ Preset |  |
| 0x0088AD60 | `Latin` | EQ Preset |  |
| 0x0088AD68 | `Loudness` | EQ Preset |  |
| 0x0088AD74 | `Lounge` | EQ Preset |  |
| 0x0088AD7C | `Piano` | EQ Preset |  |
| 0x0088AD8C | `Rock` | EQ Preset |  |
| 0x008934F8 | `Dance` | EQ Preset |  |
| 0x00893524 | `Jazz` | EQ Preset |  |
| 0x00893534 | `Loudness` | EQ Preset |  |
| 0x00893540 | `Lounge` | EQ Preset |  |
| 0x00893548 | `Piano` | EQ Preset |  |
| 0x00893558 | `Rock` | EQ Preset |  |
| 0x0089BDFC | `Hip Hop` | EQ Preset |  |
| 0x0089BE04 | `Jazz` | EQ Preset |  |
| 0x0089BE28 | `Lounge` | EQ Preset |  |
| 0x0089BE30 | `Piano` | EQ Preset |  |
| 0x0089BE40 | `Rock` | EQ Preset |  |
| 0x008A4B10 | `Hip Hop` | EQ Preset |  |
| 0x008A4B18 | `Jazz` | EQ Preset |  |
| 0x008A4B34 | `Lounge` | EQ Preset |  |
| 0x008A4B3C | `Piano` | EQ Preset |  |
| 0x008A4B4C | `Rock` | EQ Preset |  |
| 0x008B92E0 | `Acoustic` | EQ Preset |  |
| 0x008B92EC | `Bass Booster` | EQ Preset |  |
| 0x008B930C | `Classical` | EQ Preset |  |
| 0x008B9318 | `Dance` | EQ Preset |  |
| 0x008B9328 | `Electronic` | EQ Preset |  |
| 0x008B933C | `Hip Hop` | EQ Preset |  |
| 0x008B9344 | `Jazz` | EQ Preset |  |
| 0x008B934C | `Latin` | EQ Preset |  |
| 0x008B9354 | `Loudness` | EQ Preset |  |
| 0x008B9360 | `Lounge` | EQ Preset |  |
| 0x008B9368 | `Piano` | EQ Preset |  |
| 0x008B9378 | `Rock` | EQ Preset |  |
| 0x008B9380 | `Small Speakers` | EQ Preset |  |
| 0x008B9390 | `Spoken Word` | EQ Preset |  |
| 0x008B939C | `Treble Booster` | EQ Preset |  |
| 0x008B93BC | `Vocal Booster` | EQ Preset |  |
| 0x008C1C4C | `Hip Hop` | EQ Preset |  |
| 0x008C1C58 | `Latin` | EQ Preset |  |
| 0x008C1C60 | `Loudness` | EQ Preset |  |
| 0x008C1C6C | `Lounge` | EQ Preset |  |
| 0x008C1C84 | `Rock` | EQ Preset |  |
| 0x008CA784 | `Acoustic` | EQ Preset |  |
| 0x008CA790 | `Bass Booster` | EQ Preset |  |
| 0x008CA7B0 | `Classical` | EQ Preset |  |
| 0x008CA7BC | `Dance` | EQ Preset |  |
| 0x008CA7CC | `Electronic` | EQ Preset |  |
| 0x008CA7E0 | `Hip Hop` | EQ Preset |  |
| 0x008CA7E8 | `Jazz` | EQ Preset |  |
| 0x008CA7F0 | `Latin` | EQ Preset |  |
| 0x008CA7F8 | `Loudness` | EQ Preset |  |
| 0x008CA804 | `Lounge` | EQ Preset |  |
| 0x008CA80C | `Piano` | EQ Preset |  |
| 0x008CA81C | `Rock` | EQ Preset |  |
| 0x008CA824 | `Small Speakers` | EQ Preset |  |
| 0x008CA834 | `Spoken Word` | EQ Preset |  |
| 0x008CA840 | `Treble Booster` | EQ Preset |  |
| 0x008CA860 | `Vocal Booster` | EQ Preset |  |
| 0x008D312C | `Acoustic` | EQ Preset |  |
| 0x008D3138 | `Bass Booster` | EQ Preset |  |
| 0x008D3158 | `Classical` | EQ Preset |  |
| 0x008D3164 | `Dance` | EQ Preset |  |
| 0x008D3174 | `Electronic` | EQ Preset |  |
| 0x008D3188 | `Hip Hop` | EQ Preset |  |
| 0x008D3190 | `Jazz` | EQ Preset |  |
| 0x008D3198 | `Latin` | EQ Preset |  |
| 0x008D31A0 | `Loudness` | EQ Preset |  |
| 0x008D31AC | `Lounge` | EQ Preset |  |
| 0x008D31B4 | `Piano` | EQ Preset |  |
| 0x008D31C4 | `Rock` | EQ Preset |  |
| 0x008D31CC | `Small Speakers` | EQ Preset |  |
| 0x008D31DC | `Spoken Word` | EQ Preset |  |
| 0x008D31E8 | `Treble Booster` | EQ Preset |  |
| 0x008D3208 | `Vocal Booster` | EQ Preset |  |

---
